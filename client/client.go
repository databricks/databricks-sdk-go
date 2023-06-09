package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/google/go-querystring/query"
	"golang.org/x/time/rate"
)

func New(cfg *config.Config) (*DatabricksClient, error) {
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
	Config             *config.Config
	rateLimiter        *rate.Limiter
	retryTimeout       time.Duration
	httpClient         httpClient
	debugHeaders       bool
	debugTruncateBytes int
}

// ConfiguredAccountID returns Databricks Account ID if it's provided in config,
// empty string otherwise
func (c *DatabricksClient) ConfiguredAccountID() string {
	return c.Config.AccountID
}

// Do sends an HTTP request against path.
func (c *DatabricksClient) Do(ctx context.Context, method, path string,
	request, response any, visitors ...func(*http.Request) error) error {
	body, err := c.perform(ctx, method, path, request, visitors...)
	if err != nil {
		return err
	}
	return c.unmarshal(body, response)
}

func (c *DatabricksClient) unmarshal(body []byte, response any) error {
	if response == nil {
		return nil
	}
	if len(body) == 0 {
		return nil
	}
	// If the destination is bytes.Buffer, write the body over there
	if raw, ok := response.(*bytes.Buffer); ok {
		_, err := raw.Write(body)
		return err
	}
	// If the destination is a byte slice, pass the body verbatim.
	if raw, ok := response.(*[]byte); ok {
		*raw = body
		return nil
	}
	return json.Unmarshal(body, &response)
}

func (c *DatabricksClient) addHostToRequestUrl(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("no URL found in request")
	}
	url, err := url.Parse(c.Config.Host)
	if err != nil {
		return err
	}
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	return nil
}

func (c *DatabricksClient) addApplicationJsonContentType(r *http.Request) error {
	r.Header.Set("Content-Type", "application/json")
	return nil
}

func (c *DatabricksClient) redactedDump(prefix string, body []byte) (res string) {
	return bodyLogger{
		debugTruncateBytes: c.debugTruncateBytes,
	}.redactedDump(prefix, body)
}

func (c *DatabricksClient) attempt(
	ctx context.Context,
	method string,
	requestURL string,
	requestBody []byte,
	visitors ...func(*http.Request) error,
) func() (*bytes.Buffer, *retries.Err) {
	return func() (*bytes.Buffer, *retries.Err) {
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
		response, err := c.httpClient.Do(request)

		// Read responseBody immediately because we need it regardless of success or failure.
		//
		// Fully read and close HTTP response responseBody to release HTTP connections.
		//
		// HTTP client's Transport may not reuse HTTP/1.x "keep-alive" TCP connections
		// if the Body is not read to completion and closed.
		//
		// See: https://groups.google.com/g/golang-nuts/c/IoSvPz-rpfc
		var responseBody bytes.Buffer
		var responseBodyErr error
		if response != nil {
			_, responseBodyErr = responseBody.ReadFrom(response.Body)
			response.Body.Close()
		}

		retry, err := apierr.CheckForRetry(ctx, response, err, responseBody.Bytes(), responseBodyErr)
		var apiErr *apierr.APIError
		if err != nil && !errors.As(err, &apiErr) {
			err = fmt.Errorf("failed request: %w", err)
		}
		if err == nil && response == nil {
			err = fmt.Errorf("no response: %s %s", method, requestURL)
		}
		if err == nil && responseBodyErr != nil {
			err = fmt.Errorf("response body: %w", responseBodyErr)
		}

		defer c.recordRequestLog(ctx, request, response, err, requestBody, responseBody.Bytes())

		if err == nil {
			return &responseBody, nil
		}

		// proactively release the connections in HTTP connection pool
		c.httpClient.CloseIdleConnections()
		if retry {
			return nil, retries.Continue(err)
		}

		return nil, retries.Halt(err)
	}
}

func (c *DatabricksClient) recordRequestLog(
	ctx context.Context,
	request *http.Request,
	response *http.Response,
	err error,
	requestBody []byte,
	responseBody []byte,
) {
	// Don't compute expensive debug message if debug logging is not enabled.
	if !logger.Get(ctx).Enabled(ctx, logger.LevelDebug) {
		return
	}
	sb := strings.Builder{}
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
	sb.WriteString("< ")
	if response != nil {
		sb.WriteString(fmt.Sprintf("%s %s", response.Proto, response.Status))
		// Only display error on this line if the response body is empty.
		// Otherwise the response body will include details about the error.
		if len(responseBody) == 0 && err != nil {
			sb.WriteString(fmt.Sprintf(" (Error: %s)", err))
		}
	} else {
		sb.WriteString(fmt.Sprintf("Error: %s", err))
	}
	sb.WriteString("\n")
	if len(responseBody) > 0 {
		sb.WriteString(c.redactedDump("< ", responseBody))
	}
	logger.Debugf(ctx, sb.String()) // lgtm [go/log-injection] lgtm [go/clear-text-logging]
}

func (c *DatabricksClient) addAuthHeaderToUserAgent(r *http.Request) error {
	ctx := useragent.InContext(r.Context(), "auth", c.Config.AuthType)
	*r = *r.WithContext(ctx) // replace request
	return nil
}

func (c *DatabricksClient) perform(
	ctx context.Context,
	method,
	requestURL string,
	data interface{},
	visitors ...func(*http.Request) error,
) ([]byte, error) {
	requestBody, err := makeRequestBody(method, &requestURL, data)
	if err != nil {
		return nil, fmt.Errorf("request marshal: %w", err)
	}
	visitors = append([]func(*http.Request) error{
		c.Config.Authenticate,
		c.addHostToRequestUrl,
		c.addApplicationJsonContentType,
		c.addAuthHeaderToUserAgent,
	}, visitors...)
	resp, err := retries.Poll(ctx, c.retryTimeout,
		c.attempt(ctx, method, requestURL, requestBody, visitors...))
	if err != nil {
		// Don't re-wrap, as upper layers may depend on handling apierr.APIError.
		return nil, err
	}
	return resp.Bytes(), nil
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

// Remove all custom request serializer logic once APP-1331 is rolled out.
type serializer func(any) ([]byte, error)

func serializeQueryParamsToRequestBody(data any) ([]byte, error) {
	m := map[string]any{}
	rv := reflect.ValueOf(data)
	rt := reflect.TypeOf(data)
	// If data is a map, just serialize it to JSON.
	if rv.Kind() == reflect.Map {
		return json.MarshalIndent(data, "", "  ")
	}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		tag := rt.Field(i).Tag.Get("url")
		if tag == "" {
			continue
		}
		value := field.Interface()
		if field.IsZero() {
			continue
		}
		key := strings.Split(tag, ",")[0]
		m[key] = value
	}

	return json.MarshalIndent(m, "", "  ")
}

// List of (method, URL) pairs whose request bodies need to be serialized in a
// custom way. These can be removed after APP-1331 is rolled out.
var requestInBodyOverrides = []struct {
	method     string
	urlRegexp  *regexp.Regexp
	serializer serializer
}{
	{"DELETE", regexp.MustCompile("(/api/2.1)?/unity-catalog/(metastores|catalogs)/[^/]+"), serializeQueryParamsToRequestBody},
	{"DELETE", regexp.MustCompile("(/api/2.1)?/unity-catalog/workspaces/[^/]+/(metastore|catalog)"), serializeQueryParamsToRequestBody},
}

func getRequestCustomSerializer(method string, requestURL *string) serializer {
	for _, override := range requestInBodyOverrides {
		// Return true if the request URL has the prefix and no trailing segments
		// which would indicate other APIs.
		matchRequestMethod := method == override.method
		matchRequestURL := override.urlRegexp.MatchString(*requestURL)
		if matchRequestMethod && matchRequestURL {
			return override.serializer
		}
	}
	return nil
}

func makeRequestBody(method string, requestURL *string, data interface{}) ([]byte, error) {
	var requestBody []byte
	if data == nil && (method == "DELETE" || method == "GET") {
		return requestBody, nil
	}
	if customSerializer := getRequestCustomSerializer(method, requestURL); customSerializer != nil {
		return customSerializer(data)
	}
	if method == "GET" || method == "DELETE" {
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

// CWE-117 prevention
func escapeNewLines(in string) string {
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	return in
}
