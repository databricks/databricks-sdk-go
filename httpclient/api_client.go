package httpclient // has to be a separate package than client, otherwise a circular dependency

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
)

type RequestVisitor func(*http.Request) error

type ClientConfig struct {
	Visitors []RequestVisitor

	RetryTimeout       time.Duration
	HTTPTimeout        time.Duration
	InsecureSkipVerify bool
	DebugHeaders       bool
	DebugTruncateBytes int
	RateLimitPerSecond int

	ErrorMapper     func(ctx context.Context, resp *http.Response, body io.ReadCloser) error
	ErrorRetriable  func(ctx context.Context, err error) bool
	TransientErrors []string

	Transport http.RoundTripper
}

func (cfg ClientConfig) httpTransport() http.RoundTripper {
	if cfg.Transport != nil {
		return cfg.Transport
	}
	return &http.Transport{
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
	}
}

var DefaultClient = NewApiClient(ClientConfig{
	ErrorRetriable:     DefaultErrorRetriable,
	ErrorMapper:        DefaultErrorMapper,
	HTTPTimeout:        30 * time.Second,
	RetryTimeout:       5 * time.Minute,
	RateLimitPerSecond: 30,
})

func NewApiClient(cfg ClientConfig) *ApiClient {
	cfg.HTTPTimeout = time.Duration(orDefault(int(cfg.HTTPTimeout), int(30*time.Second)))
	cfg.DebugTruncateBytes = orDefault(cfg.DebugTruncateBytes, 96)
	cfg.RetryTimeout = time.Duration(orDefault(int(cfg.RetryTimeout), int(5*time.Minute)))
	cfg.HTTPTimeout = time.Duration(orDefault(int(cfg.HTTPTimeout), int(30*time.Second)))
	rateLimiter := rate.NewLimiter(rate.Limit(orDefault(cfg.RateLimitPerSecond, 15)), 1)
	if cfg.ErrorMapper == nil {
		// default generic error mapper
		cfg.ErrorMapper = DefaultErrorMapper
	}
	if cfg.ErrorRetriable == nil {
		// by default, we just retry on HTTP 429/504
		cfg.ErrorRetriable = DefaultErrorRetriable
	}
	return &ApiClient{
		config:      cfg,
		rateLimiter: rateLimiter,
		httpClient: &http.Client{
			Timeout:   cfg.HTTPTimeout,
			Transport: cfg.httpTransport(),
		},
	}
}

type ApiClient struct {
	config      ClientConfig
	rateLimiter *rate.Limiter
	httpClient  *http.Client
}

type DoOption struct {
	in   RequestVisitor
	out  func(body *responseBody) error
	body any
}

// Do sends an HTTP request against path.
func (c *ApiClient) Do(ctx context.Context, method, path string, opts ...DoOption) error {
	visitors := c.config.Visitors[:]
	for _, o := range opts {
		if o.in == nil {
			continue
		}
		// merge client-wide and request-specific visitors
		visitors = append(visitors, o.in)
	}
	var requestBody any
	for _, o := range opts {
		if o.body == nil {
			continue
		}
		requestBody = o.body
	}
	responseBody, err := c.perform(ctx, method, path, requestBody, visitors...)
	if err != nil {
		return err
	}
	for _, o := range opts {
		if o.out == nil {
			continue
		}
		err = o.out(responseBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ApiClient) fromResponse(r *http.Response) (responseBody, error) {
	if r == nil {
		return responseBody{}, fmt.Errorf("nil response")
	}
	if r.Request == nil {
		return responseBody{}, fmt.Errorf("nil request")
	}
	streamResponse := r.Request.Header.Get("Accept") != "application/json" && r.Header.Get("Content-Type") != "application/json"
	if streamResponse {
		return newResponseBody(r.Body, r.Header)
	}
	defer r.Body.Close()
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		return responseBody{}, fmt.Errorf("response body: %w", err)
	}
	return newResponseBody(bs, r.Header)
}

func (c *ApiClient) redactedDump(prefix string, body []byte) (res string) {
	return bodyLogger{
		debugTruncateBytes: c.config.DebugTruncateBytes,
	}.redactedDump(prefix, body)
}

func (c *ApiClient) isRetriable(ctx context.Context, err error) bool {
	if c.config.ErrorRetriable(ctx, err) {
		return true
	}
	_, isIO := err.(*url.Error)
	if isIO {
		// all IO errors are retriable
		logger.Debugf(ctx, "Attempting retry because of IO error: %s", err)
		return true
	}
	message := err.Error()
	// Handle transient errors for retries
	for _, substring := range c.config.TransientErrors {
		if strings.Contains(message, substring) {
			logger.Debugf(ctx, "Attempting retry because of %#v", substring)
			return true
		}
	}
	// some API's recommend retries on HTTP 500, but we'll add that later
	return false
}

// Common error-handling logic for all responses that may need to be retried.
//
// If the error is retriable, return a retries.Err to retry the request. However, as the request body will have been consumed
// by the first attempt, the body must be reset before retrying. If the body cannot be reset, return a retries.Err to halt.
//
// Always returns nil for the first parameter as there is no meaningful response body to return in the error case.
//
// If it is certain that an error should not be retried, use failRequest() instead.
func (c *ApiClient) handleError(ctx context.Context, err error, body requestBody) (*responseBody, *retries.Err) {
	if !c.isRetriable(ctx, err) {
		return c.failRequest(ctx, "non-retriable error", err)
	}
	if resetErr := body.reset(); resetErr != nil {
		return nil, retries.Halt(resetErr)
	}
	return nil, retries.Continue(err)
}

// Fails the request with a retries.Err to halt future retries.
func (c *ApiClient) failRequest(ctx context.Context, msg string, err error) (*responseBody, *retries.Err) {
	logger.Debugf(ctx, "%s: %s", msg, err)
	return nil, retries.Halt(err)
}

func (c *ApiClient) attempt(
	ctx context.Context,
	method string,
	requestURL string,
	requestBody requestBody,
	visitors ...RequestVisitor,
) func() (*responseBody, *retries.Err) {
	return func() (*responseBody, *retries.Err) {
		err := c.rateLimiter.Wait(ctx)
		if err != nil {
			return c.failRequest(ctx, "failed in rate limiter", err)
		}
		request, err := http.NewRequestWithContext(ctx, method, requestURL, requestBody.Reader)
		if err != nil {
			return c.failRequest(ctx, "failed creating new request", err)
		}
		for _, requestVisitor := range visitors {
			err = requestVisitor(request)
			if err != nil {
				return c.failRequest(ctx, "failed during request visitor", err)
			}
		}
		// request.Context() holds context potentially enhanced by visitors
		request.Header.Set("User-Agent", useragent.FromContext(request.Context()))

		// attempt the actual request
		response, err := c.httpClient.Do(request)

		// After this point, the request body has (probably) been consumed. handleError() must be called to reset it if
		// possible.
		if _, ok := err.(*url.Error); ok {
			return c.handleError(ctx, err, requestBody)
		}

		// By this point, the request body has certainly been consumed.
		responseBody, responseBodyErr := c.fromResponse(response)
		if responseBodyErr != nil {
			return c.failRequest(ctx, "failed while reading response", responseBodyErr)
		}

		mappedError := c.config.ErrorMapper(ctx, response, responseBody.ReadCloser)
		defer c.recordRequestLog(ctx, request, response, mappedError, requestBody.DebugBytes, responseBody.DebugBytes)

		if mappedError == nil {
			return &responseBody, nil
		}

		// proactively release the connections in HTTP connection pool
		c.httpClient.CloseIdleConnections()
		return c.handleError(ctx, mappedError, requestBody)
	}
}

func (c *ApiClient) recordRequestLog(
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
	if c.config.DebugHeaders {
		sb.WriteString("> * Host: ")
		sb.WriteString(escapeNewLines(request.Host))
		sb.WriteString("\n")
		for k, v := range request.Header {
			trunc := onlyNBytes(strings.Join(v, ""), c.config.DebugTruncateBytes)
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

// RoundTrip implements http.RoundTripper to integrate with golang.org/x/oauth2
func (c *ApiClient) RoundTrip(request *http.Request) (*http.Response, error) {
	ctx := request.Context()
	requestURL := request.URL.String()
	resp, err := retries.Poll(ctx, c.config.RetryTimeout,
		c.attempt(ctx, request.Method, requestURL, requestBody{
			Reader: request.Body,
			// DO NOT DECODE BODY, because it may contain sensitive payload,
			// like Azure Service Principal in a multipart/form-data body.
			DebugBytes: []byte("<http.RoundTripper>"),
		}))
	if err != nil {
		return nil, err
	}
	// here we assume only successful responses, as HTTP 4XX and 5XX are mapped
	// to Go's error implementations.
	return &http.Response{
		Status:     "OK",
		StatusCode: 200,
		Request:    request,
		Header:     resp.Header,
		Body:       resp.ReadCloser,
	}, nil
}

// InContextForOAuth2 returns a context with a custom *http.Client to be used
// for only for token acquisition through golang.org/x/oauth2 package
func (c *ApiClient) InContextForOAuth2(ctx context.Context) context.Context {
	return context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
		Timeout:   c.config.HTTPTimeout,
		Transport: c,
	})
}

func (c *ApiClient) perform(
	ctx context.Context,
	method,
	requestURL string,
	data interface{},
	visitors ...RequestVisitor,
) (*responseBody, error) {
	requestBody, err := makeRequestBody(method, &requestURL, data)
	if err != nil {
		return nil, fmt.Errorf("request marshal: %w", err)
	}
	resp, err := retries.Poll(ctx, c.config.RetryTimeout,
		c.attempt(ctx, method, requestURL, requestBody, visitors...))
	var timedOut *retries.ErrTimedOut
	if errors.As(err, &timedOut) {
		// TODO: check if we want to unwrap this error here
		return nil, timedOut.Unwrap()
	} else if err != nil {
		// Don't re-wrap, as upper layers may depend on handling apierr.APIError.
		return nil, err
	}
	return resp, nil
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
