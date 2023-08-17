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

// Represents a request or response Body.
//
// If the provided request data is an io.ReadCloser, DebugBytes is set to
// "<io.ReadCloser>". Otherwise, DebugBytes is set to the marshaled JSON
// representation of the request data, and ReadCloser is set to a new
// io.ReadCloser that reads from DebugBytes.
type Body struct {
	ReadCloser io.ReadCloser
	DebugBytes []byte
}

func fromBytes(bs []byte) Body {
	return Body{
		ReadCloser: io.NopCloser(bytes.NewReader(bs)),
		DebugBytes: bs,
	}
}

func fromReader(r io.Reader) Body {
	return Body{
		ReadCloser: io.NopCloser(r),
		DebugBytes: []byte("<io.Reader>"),
	}
}

func fromJsonData(data any) (Body, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return Body{}, err
	}
	return fromBytes(bs), nil
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
	headers map[string]string, request, response any,
	visitors ...func(*http.Request) error) error {
	body, err := c.perform(ctx, method, path, headers, request, visitors...)
	if err != nil {
		return err
	}
	return c.unmarshal(body, response)
}

func (c *DatabricksClient) unmarshal(body *Body, response any) error {
	if response == nil {
		return nil
	}
	// If the destination is bytes.Buffer, write the body over there
	if raw, ok := response.(**io.ReadCloser); ok {
		*raw = &body.ReadCloser
		return nil
	}
	defer func() { body.ReadCloser.Close() }()
	bs := body.DebugBytes
	if len(bs) == 0 {
		return nil
	}
	// If the destination is a byte slice, pass the body verbatim.
	if raw, ok := response.(*[]byte); ok {
		*raw = bs
		return nil
	}
	return json.Unmarshal(bs, &response)
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

func (c *DatabricksClient) fromResponse(r *http.Response) (Body, error) {
	safeToReadFullResponse := r.Header.Get("Content-Type") == "application/json"
	if safeToReadFullResponse {
		defer r.Body.Close()
		bs, err := io.ReadAll(r.Body)
		if err != nil {
			return Body{}, err
		}
		return fromBytes(bs), nil
	}
	return Body{
		ReadCloser: r.Body,
		DebugBytes: []byte("<io.ReadCloser>"),
	}, nil
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
	headers map[string]string,
	requestBody Body,
	visitors ...func(*http.Request) error,
) func() (*Body, *retries.Err) {
	return func() (*Body, *retries.Err) {
		err := c.rateLimiter.Wait(ctx)
		if err != nil {
			return nil, retries.Halt(err)
		}
		request, err := http.NewRequestWithContext(ctx, method, requestURL, requestBody.ReadCloser)
		for k, v := range headers {
			request.Header.Set(k, v)
		}
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
		responseBody, responseBodyErr := c.fromResponse(response)

		//
		retry, err := apierr.CheckForRetry(ctx, response, err, responseBody.DebugBytes, responseBodyErr)
		var apiErr *apierr.APIError
		if err != nil && !errors.As(err, &apiErr) {
			err = fmt.Errorf("failed request: %w", err)
		}
		if err == nil && response == nil {
			err = fmt.Errorf("no response: %s %s", method, requestURL)
		}

		defer c.recordRequestLog(ctx, request, response, err, requestBody.DebugBytes, responseBody.DebugBytes)

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
	requestBody, responseBody []byte,
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
	headers map[string]string,
	data interface{},
	visitors ...func(*http.Request) error,
) (*Body, error) {
	requestBody, err := makeRequestBody(method, &requestURL, data)
	if err != nil {
		return nil, fmt.Errorf("request marshal: %w", err)
	}
	visitors = append([]func(*http.Request) error{
		c.Config.Authenticate,
		c.addHostToRequestUrl,
		c.addAuthHeaderToUserAgent,
	}, visitors...)
	resp, err := retries.Poll(ctx, c.retryTimeout,
		c.attempt(ctx, method, requestURL, headers, requestBody, visitors...))
	if err != nil {
		// Don't re-wrap, as upper layers may depend on handling apierr.APIError.
		return nil, err
	}
	return resp, nil
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
		// Query parameters may be nested, but the keys generated by
		// query.Values use the "[" and "]" characters to represent nesting.
		// Replace all instances of "[" with "." and "]" with empty string
		// to make the query string compatible with the proto API.
		// See the following for more information:
		// https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule
		protoCompatibleParams := make(url.Values)
		for k, vs := range params {
			newK := strings.Replace(k, "[", ".", -1)
			newK = strings.Replace(newK, "]", "", -1)
			for _, v := range vs {
				protoCompatibleParams.Add(newK, v)
			}
		}
		return "?" + protoCompatibleParams.Encode(), nil
	}
	return "", fmt.Errorf("unsupported query string data: %#v", data)
}

// Remove all custom request serializer logic once APP-1331 is rolled out.
type serializer func(any) (Body, error)

func serializeQueryParamsToRequestBody(data any) (Body, error) {
	m := map[string]any{}
	rv := reflect.ValueOf(data)
	rt := reflect.TypeOf(data)
	// If data is a map, just serialize it to JSON.
	if rv.Kind() == reflect.Map {
		return fromJsonData(data)
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

	return fromJsonData(data)
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

func makeRequestBody(method string, requestURL *string, data interface{}) (Body, error) {
	if data == nil && (method == "DELETE" || method == "GET") {
		return Body{}, nil
	}
	if customSerializer := getRequestCustomSerializer(method, requestURL); customSerializer != nil {
		return customSerializer(data)
	}
	if method == "GET" || method == "DELETE" {
		qs, err := makeQueryString(data)
		if err != nil {
			return Body{}, err
		}
		*requestURL += qs
		return fromBytes([]byte{}), nil
	}
	if bytes, ok := data.([]byte); ok {
		return fromBytes(bytes), nil
	}
	if reader, ok := data.(io.Reader); ok {
		return fromReader(reader), nil
	}
	if str, ok := data.(string); ok {
		return fromBytes([]byte(str)), nil
	}
	return fromJsonData(data)
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
