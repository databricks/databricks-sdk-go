package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/databricks/logger"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"github.com/google/go-querystring/query"
	"golang.org/x/time/rate"
)

func New(cfg *databricks.Config) (*DatabricksClient, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	retryTimeout := time.Duration(orDefault(cfg.RetryTimeoutSeconds, 300)) * time.Second
	httpTimeout := time.Duration(orDefault(cfg.HTTPTimeoutSeconds, 60)) * time.Second
	rateLimiter := rate.NewLimiter(rate.Limit(orDefault(cfg.RateLimitPerSecond, 15)), 1)
	debugTruncateBytes := orDefault(cfg.DebugTruncateBytes, 96)
	return &DatabricksClient{
		Config:             cfg,
		debugHeaders:       cfg.DebugHeaders,
		debugTruncateBytes: debugTruncateBytes,
		retryTimeout:       retryTimeout,
		rateLimiter:        rateLimiter,
		httpClient: &http.Client{
			Timeout: httpTimeout,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
				IdleConnTimeout:       180 * time.Second,
				TLSHandshakeTimeout:   30 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: cfg.InsecureSkipVerify,
				},
			},
		},
	}, nil
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
	CloseIdleConnections()
}

type DatabricksClient struct {
	Config             *databricks.Config
	rateLimiter        *rate.Limiter
	retryTimeout       time.Duration
	httpClient         httpClient
	debugHeaders       bool
	debugTruncateBytes int
}

// Do sends an HTTP request against path.
func (c *DatabricksClient) Do(ctx context.Context, method string, path string, request interface{}, response interface{}) error {
	body, err := c.perform(ctx, method, path, request, c.completeUrl)
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

// Get on path
func (c *DatabricksClient) Get(ctx context.Context, path string, request interface{}, response interface{}) error {
	return c.Do(ctx, http.MethodGet, path, request, response)
}

// Post on path
func (c *DatabricksClient) Post(ctx context.Context, path string, request interface{}, response interface{}) error {
	return c.Do(ctx, http.MethodPost, path, request, response)
}

// Delete on path
func (c *DatabricksClient) Delete(ctx context.Context, path string, request interface{}) error {
	return c.Do(ctx, http.MethodDelete, path, request, nil)
}

// Patch on path
func (c *DatabricksClient) Patch(ctx context.Context, path string, request interface{}) error {
	return c.Do(ctx, http.MethodPatch, path, request, nil)
}

// Put on path
func (c *DatabricksClient) Put(ctx context.Context, path string, request interface{}) error {
	return c.Do(ctx, http.MethodPut, path, request, nil)
}

func (c *DatabricksClient) unmarshall(path string, body []byte, response interface{}) error {
	if response == nil {
		return nil
	}
	if len(body) == 0 {
		return nil
	}
	return json.Unmarshal(body, &response)
}

type contextKey int

const Api contextKey = 5

type ApiVersion string

const (
	API_1_2 ApiVersion = "1.2"
	API_2_0 ApiVersion = "2.0"
	API_2_1 ApiVersion = "2.1"
)

func (c *DatabricksClient) completeUrl(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("no URL found in request")
	}
	// TODO: accounts client
	// ctx := r.Context()
	// av, ok := ctx.Value(Api).(ApiVersion)
	// if !ok {
	// 	av = API_2_0
	// }
	//r.URL.Path = fmt.Sprintf("/api/%s%s", av, r.URL.Path)
	r.Header.Set("Content-Type", "application/json")

	url, err := url.Parse(c.Config.Host)
	if err != nil {
		return err
	}
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme

	return nil
}

func (c *DatabricksClient) recursiveMask(requestMap map[string]any) any {
	for k, v := range requestMap {
		if k == "string_value" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if k == "token_value" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if k == "content" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if m, ok := v.(map[string]any); ok {
			requestMap[k] = c.recursiveMask(m)
			continue
		}
		if s, ok := v.(string); ok {
			requestMap[k] = onlyNBytes(s, c.debugTruncateBytes)
		}
	}
	return requestMap
}

func (c *DatabricksClient) redactedDump(prefix string, body []byte) (res string) {
	if len(body) == 0 {
		return
	}
	if body[0] == '[' {
		var requestSlice, tmp []any
		err := json.Unmarshal(body, &requestSlice)
		if err != nil {
			// error in this case is not much relevant
			return
		}
		for _, v := range requestSlice {
			x, ok := v.(map[string]any)
			if !ok {
				continue
			}
			tmp = append(tmp, c.recursiveMask(x))
		}
		return c.rePack(prefix, tmp)
	}
	if body[0] != '{' {
		return fmt.Sprintf("[non-JSON document of %d bytes]", len(body))
	}
	var requestMap map[string]any
	err := json.Unmarshal(body, &requestMap)
	if err != nil {
		// error in this case is not much relevant
		return
	}
	return c.rePack(prefix, c.recursiveMask(requestMap))
}

func (c *DatabricksClient) rePack(prefix string, v any) string {
	rePacked, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		// error in this case is not much relevant
		return ""
	}
	maxBytes := 1024
	if c.debugTruncateBytes > maxBytes {
		maxBytes = c.debugTruncateBytes
	}
	lines := strings.Split(onlyNBytes(string(rePacked), maxBytes), "\n")
	return prefix + strings.Join(lines, "\n"+prefix)
}

// Fully read and close HTTP response body to release HTTP connections.
//
// HTTP client's Transport may not reuse HTTP/1.x "keep-alive" TCP connections
// if the Body is not read to completion and closed.
//
// See: https://groups.google.com/g/golang-nuts/c/IoSvPz-rpfc
func (c *DatabricksClient) drainToEnsureConnectionRelease(r *http.Response) {
	if r == nil || r.Body == nil {
		return
	}
	defer r.Body.Close()
	_, err := io.Copy(ioutil.Discard, r.Body)
	if err != nil {
		logger.Errorf("failed to drain body: %s", err)
	}
}

func (c *DatabricksClient) attempt(ctx context.Context,
	method, requestURL string, requestBody []byte,
	visitors ...func(*http.Request) error) func() (*http.Response, *retries.Err) {
	return func() (*http.Response, *retries.Err) {
		err := c.rateLimiter.Wait(ctx)
		if err != nil {
			return nil, retries.Halt(err)
		}
		request, err := http.NewRequestWithContext(ctx, method, requestURL,
			bytes.NewBuffer(requestBody))
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, requestVisitor := range visitors {
			err = requestVisitor(request)
			if err != nil {
				return nil, retries.Halt(err)
			}
		}
		// request.Context() holds context potentially enhanced by visitors
		request.Header.Set("User-Agent", useragent.FromContext(request.Context()))

		// attempt the actual request
		resp, err := c.httpClient.Do(request)
		retry, err := apierr.CheckForRetry(ctx, resp, err)
		if err == nil {
			return resp, nil
		}
		c.drainToEnsureConnectionRelease(resp)
		// proactively release the connections in HTTP connection pool
		c.httpClient.CloseIdleConnections()
		if retry {
			return nil, retries.Continue(err)
		}
		return nil, retries.Halt(err)
	}
}

func (c *DatabricksClient) recordRequestLog(response *http.Response,
	requestBody, responseBody []byte) {
	sb := strings.Builder{}
	request := response.Request
	sb.WriteString(fmt.Sprintf("%s %s", request.Method,
		escapeNewLines(request.URL.Path)))
	if request.URL.RawQuery != "" {
		sb.WriteString("?")
		q, _ := url.QueryUnescape(request.URL.RawQuery)
		sb.WriteString(q)
	}
	sb.WriteString("\n")
	if c.debugHeaders {
		if c.Config.Host != "" {
			sb.WriteString("> * Host: ")
			sb.WriteString(escapeNewLines(c.Config.Host))
			sb.WriteString("\n")
		}
		for k, v := range request.Header {
			trunc := onlyNBytes(strings.Join(v, ""), c.debugTruncateBytes)
			sb.WriteString(fmt.Sprintf("> * %s: %s\n", k, escapeNewLines(trunc)))
		}
	}
	if len(requestBody) > 0 {
		sb.WriteString(c.redactedDump("> ", requestBody))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("< %s %s\n", response.Proto, response.Status))
	if len(responseBody) > 0 {
		sb.WriteString(c.redactedDump("< ", responseBody))
	}
	logger.Debugf(sb.String()) // lgtm [go/log-injection] lgtm [go/clear-text-logging]
}

func (c *DatabricksClient) addAuthHeaderToUserAgent(r *http.Request) error {
	ctx := useragent.InContext(r.Context(), "auth", c.Config.AuthType)
	*r = *r.WithContext(ctx) // replace request
	return nil
}

func (c *DatabricksClient) perform(ctx context.Context, method, requestURL string, data interface{},
	visitors ...func(*http.Request) error) (responseBody []byte, err error) {
	requestBody, err := makeRequestBody(method, &requestURL, data)
	if err != nil {
		return nil, fmt.Errorf("request marshal: %w", err)
	}
	visitors = append([]func(*http.Request) error{
		c.Config.Authenticate,
		c.addAuthHeaderToUserAgent,
	}, visitors...)
	resp, err := retries.Poll(ctx, c.retryTimeout,
		c.attempt(ctx, method, requestURL, requestBody, visitors...))
	var ae apierr.APIError
	if errors.As(err, &ae) {
		// don't re-wrap, as upper layers may depend on handling common.APIError
		return nil, ae
	}
	if err != nil {
		// i don't even know which errors in the real world would end up here.
		// `retryablehttp` package nicely wraps _everything_ to `url.Error`.
		return nil, fmt.Errorf("failed request: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("no response: %s %s", method, requestURL)
	}
	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response body: %w", err)
	}
	c.recordRequestLog(resp, requestBody, responseBody)
	c.drainToEnsureConnectionRelease(resp)
	return responseBody, nil
}

func makeQueryString(data interface{}) (string, error) {
	inputVal := reflect.ValueOf(data)
	inputType := reflect.TypeOf(data)
	if inputType.Kind() == reflect.Map {
		s := []string{}
		keys := inputVal.MapKeys()
		// sort map keys by their string repr, so that tests can be deterministic
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})
		for _, k := range keys {
			v := inputVal.MapIndex(k)
			if v.IsZero() {
				continue
			}
			s = append(s, fmt.Sprintf("%s=%s",
				strings.Replace(url.QueryEscape(fmt.Sprintf("%v", k.Interface())), "+", "%20", -1),
				strings.Replace(url.QueryEscape(fmt.Sprintf("%v", v.Interface())), "+", "%20", -1)))
		}
		return "?" + strings.Join(s, "&"), nil
	}
	if inputType.Kind() == reflect.Struct {
		params, err := query.Values(data)
		if err != nil {
			return "", fmt.Errorf("cannot create query string: %w", err)
		}
		return "?" + params.Encode(), nil
	}
	return "", fmt.Errorf("unsupported query string data: %#v", data)
}

func makeRequestBody(method string, requestURL *string, data interface{}) ([]byte, error) {
	var requestBody []byte
	if data == nil && (method == "DELETE" || method == "GET") {
		return requestBody, nil
	}
	if method == "GET" {
		qs, err := makeQueryString(data)
		if err != nil {
			return nil, err
		}
		*requestURL += qs
		return requestBody, nil
	}
	if bytes, ok := data.([]byte); ok {
		return bytes, nil
	}
	if reader, ok := data.(io.Reader); ok {
		raw, err := io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("failed to read from reader: %w", err)
		}
		return raw, nil
	}
	if str, ok := data.(string); ok {
		return []byte(str), nil
	}
	bodyBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("request marshal failure: %w", err)
	}
	return bodyBytes, nil
}

func orDefault(configured, _default int) int {
	if configured == 0 {
		return _default
	}
	return configured
}

func onlyNBytes(j string, numBytes int) string {
	diff := len([]byte(j)) - numBytes
	if diff > 0 {
		return fmt.Sprintf("%s... (%d more bytes)", j[:numBytes], diff)
	}
	return j
}

// CWE-117 prevention
func escapeNewLines(in string) string {
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	return in
}
