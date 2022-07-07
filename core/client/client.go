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
	"log"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/databricks/sdk-go/core/config"
	"github.com/databricks/sdk-go/retries"
	"github.com/google/go-querystring/query"
	"golang.org/x/time/rate"
)

type Credentials interface {
	Name() string
	Configure(context.Context, *ApiClient) (func(*http.Request) error, error)
}

// Default settings
const (
	defaultDebugTruncateBytes      = 96
	defaultRateLimitPerSecond = 15
	defaultHTTPTimeoutSeconds = 60
)

type ApiClient struct {
	Config      *config.Config
	Credentials Credentials

	// Mutex used by Authenticate method to guard `authVisitor`, which
	// has to be lazily created on the first request to Databricks API.
	// It is done because databricks host and token may often be available
	// only in the middle of Terraform DAG execution.
	authMutex sync.Mutex

	// HTTP request interceptor, that assigns Authorization header
	authVisitor func(r *http.Request) error

	// Databricks REST API rate limiter
	rateLimiter *rate.Limiter

	// // retryalble HTTP client
	// httpClient *retryablehttp.Client
	httpClient *http.Client

	init sync.Once
}

func (c *ApiClient) GetHost() string {
	return c.Config.Host
}

// Authenticate lazily authenticates across authorizers or returns error
func (c *ApiClient) Authenticate(ctx context.Context) error {
	if c.authVisitor != nil {
		return nil
	}
	c.authMutex.Lock()
	defer c.authMutex.Unlock()
	if c.authVisitor != nil {
		return nil
	}
	visitor, err := c.Credentials.Configure(ctx, c)
	if err != nil {
		return fmt.Errorf("cannot configure %s auth: %w", c.Credentials.Name(), err)
	}
	c.authVisitor = visitor
	c.Config.AuthType = c.Credentials.Name() // move to implicit
	c.fixHost()
	// TODO: error customization
	return nil
}

func (c *ApiClient) authenticatedQuery(ctx context.Context, method, requestURL string,
	data interface{}, visitors ...func(*http.Request) error) (body []byte, err error) {
	err = c.Authenticate(ctx)
	if err != nil {
		return
	}
	visitors = append([]func(*http.Request) error{c.authVisitor}, visitors...)
	return c.perform(ctx, method, requestURL, data, visitors...)
}

// Get on path
func (c *ApiClient) Get(ctx context.Context, path string, request interface{}, response interface{}) error {
	body, err := c.authenticatedQuery(ctx, http.MethodGet, path, request, c.completeUrl)
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

// Post on path
func (c *ApiClient) Post(ctx context.Context, path string, request interface{}, response interface{}) error {
	body, err := c.authenticatedQuery(ctx, http.MethodPost, path, request, c.completeUrl)
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

// Delete on path
func (c *ApiClient) Delete(ctx context.Context, path string, request interface{}) error {
	_, err := c.authenticatedQuery(ctx, http.MethodDelete, path, request, c.completeUrl)
	return err
}

// Patch on path
func (c *ApiClient) Patch(ctx context.Context, path string, request interface{}) error {
	_, err := c.authenticatedQuery(ctx, http.MethodPatch, path, request, c.completeUrl)
	return err
}

// Put on path
func (c *ApiClient) Put(ctx context.Context, path string, request interface{}) error {
	_, err := c.authenticatedQuery(ctx, http.MethodPut, path, request, c.completeUrl)
	return err
}

func (c *ApiClient) unmarshall(path string, body []byte, response interface{}) error {
	if response == nil {
		return nil
	}
	if len(body) == 0 {
		return nil
	}
	err := json.Unmarshal(body, &response)
	if err == nil {
		return nil
	}
	return APIError{
		ErrorCode:  "UNKNOWN",
		StatusCode: 200,
		Resource:   "..." + path,
		Message: fmt.Sprintf("Invalid JSON received (%d bytes): %v",
			len(body), string(body)),
	}
}

type contextKey int

const Api contextKey = 5

type ApiVersion string

const (
	API_1_2 ApiVersion = "1.2"
	API_2_0 ApiVersion = "2.0"
	API_2_1 ApiVersion = "2.1"
)

func (c *ApiClient) completeUrl(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("no URL found in request")
	}

	ctx := r.Context()
	av, ok := ctx.Value(Api).(ApiVersion)
	if !ok {
		av = API_2_0
	}

	r.URL.Path = fmt.Sprintf("/api/%s%s", av, r.URL.Path)
	r.Header.Set("Content-Type", "application/json")

	url, err := url.Parse(c.Config.Host)
	if err != nil {
		return err
	}
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme

	return nil
}

func (c *ApiClient) configure() {
	if c.httpClient != nil {
		return
	}
	if c.Config.DebugTruncateBytes == 0 {
		c.Config.DebugTruncateBytes = defaultDebugTruncateBytes
	}
	if c.Config.HTTPTimeoutSeconds == 0 {
		c.Config.HTTPTimeoutSeconds = defaultHTTPTimeoutSeconds
	}
	if c.Config.RateLimitPerSecond == 0 {
		c.Config.RateLimitPerSecond = defaultRateLimitPerSecond
	}
	c.rateLimiter = rate.NewLimiter(rate.Limit(c.Config.RateLimitPerSecond), 1)
	// Set up a retryable HTTP Client to handle cases where the service returns
	// a transient error on initial creation
	retryDelayDuration := 10 * time.Second
	retryMaximumDuration := 5 * time.Minute

	c.Config.RetryWaitMin = retryDelayDuration
	c.Config.RetryWaitMax = retryDelayDuration
	c.Config.MaxRetryAttempts = int(retryMaximumDuration / retryDelayDuration)

	c.httpClient = &http.Client{
		Timeout: time.Duration(c.Config.HTTPTimeoutSeconds) * time.Second,
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
				InsecureSkipVerify: c.Config.InsecureSkipVerify,
			},
		},
	}
}

func (c *ApiClient) userAgent(ctx context.Context) string {
	// TODO: add user-agent framework to get things propagated dynamically
	terraformVersion := "unknown"
	resource := "unknown"
	version := "0.0.1"
	return fmt.Sprintf("databricks-tf-provider/%s (+%s) terraform/%s",
		version, resource, terraformVersion)
}

// CWE-117 prevention
func escapeNewLines(in string) string {
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	return in
}

func (c *ApiClient) createDebugHeaders(header http.Header, host string) string {
	headers := ""
	if c.Config.DebugHeaders {
		if host != "" {
			headers += fmt.Sprintf("\n * Host: %s", escapeNewLines(host))
		}
		for k, v := range header {
			trunc := onlyNBytes(strings.Join(v, ""), c.Config.DebugTruncateBytes)
			headers += fmt.Sprintf("\n * %s: %s", k, escapeNewLines(trunc))
		}
		if len(headers) > 0 {
			headers += "\n"
		}
	}
	return headers
}

func (c *ApiClient) recursiveMask(requestMap map[string]interface{}) interface{} {
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
		if m, ok := v.(map[string]interface{}); ok {
			requestMap[k] = c.recursiveMask(m)
			continue
		}
		// todo: dapi...
		// TODO: just redact any dapiXXX & "secret": "...."...
		if s, ok := v.(string); ok {
			requestMap[k] = onlyNBytes(s, c.Config.DebugTruncateBytes)
		}
	}
	return requestMap
}

func (c *ApiClient) redactedDump(body []byte) (res string) {
	if len(body) == 0 {
		return
	}
	if body[0] != '{' {
		return fmt.Sprintf("[non-JSON document of %d bytes]", len(body))
	}
	var requestMap map[string]interface{}
	err := json.Unmarshal(body, &requestMap)
	if err != nil {
		// error in this case is not much relevant
		return
	}
	rePacked, err := json.MarshalIndent(c.recursiveMask(requestMap), "", "  ")
	if err != nil {
		// error in this case is not much relevant
		return
	}
	maxBytes := 1024
	if c.Config.DebugTruncateBytes > maxBytes {
		maxBytes = c.Config.DebugTruncateBytes
	}
	return onlyNBytes(string(rePacked), maxBytes)
}

func (c *ApiClient) drain(r *http.Response) {
	defer r.Body.Close()
	_, err := io.Copy(ioutil.Discard, io.LimitReader(r.Body, 4096))
	if err != nil {
		log.Printf("[ERROR] failed to drain body: %s", err)
	}
}

func (c *ApiClient) retried(ctx context.Context, method, requestURL string,
	requestBody []byte, visitors ...func(*http.Request) error) (resp *http.Response, err error) {
	for attempt := 0; ; attempt++ {
		request, err := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewBuffer(requestBody))
		if err != nil {
			return nil, fmt.Errorf("new request: %w", err)
		}
		request.Header.Set("User-Agent", c.userAgent(ctx))
		for _, requestVisitor := range visitors {
			err = requestVisitor(request)
			if err != nil {
				return nil, fmt.Errorf("failed visitor: %w", err)
			}
		}
		headers := c.createDebugHeaders(request.Header, c.Config.Host)
		log.Printf("[DEBUG] %s %s %s%v", method, escapeNewLines(request.URL.Path),
			headers, c.redactedDump(requestBody)) // lgtm [go/log-injection] lgtm [go/clear-text-logging]
		// attempt the actual request
		resp, err = c.httpClient.Do(request)
		retry, err := c.checkHTTPRetry(ctx, resp, err)
		if !retry {
			break
		}
		if (c.Config.MaxRetryAttempts - attempt) <= 0 {
			break
		}
		if err == nil {
			c.drain(resp)
		}
		wait := retries.Backoff(attempt)
		timer := time.NewTimer(wait)
		select {
		case <-ctx.Done():
			timer.Stop()
			c.httpClient.CloseIdleConnections()
			return nil, fmt.Errorf("retry timed out: %w", ctx.Err())
		case <-timer.C:
		}
	}
	if err == nil {
		return resp, nil
	}
	defer c.httpClient.CloseIdleConnections()
	if resp != nil {
		c.drain(resp)
	}
	return nil, fmt.Errorf("exceeded %d retries: %s %s: %w", c.Config.MaxRetryAttempts, method, requestURL, err)
}

func (c *ApiClient) perform(ctx context.Context, method, requestURL string, data interface{},
	visitors ...func(*http.Request) error) (body []byte, err error) {
	c.init.Do(c.configure)
	if c.authVisitor == nil {
		return nil, fmt.Errorf("no credentials")
	}
	if err = c.rateLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limited: %w", err)
	}
	requestBody, err := makeRequestBody(method, &requestURL, data)
	if err != nil {
		return nil, fmt.Errorf("request marshal: %w", err)
	}
	resp, err := c.retried(ctx, method, requestURL, requestBody, visitors...)
	// retryablehttp library now returns only wrapped errors
	var ae APIError
	if errors.As(err, &ae) {
		// don't re-wrap, as upper layers may depend on handling common.APIError
		return nil, ae
	}
	if err != nil {
		// i don't even know which errors in the real world would end up here.
		// `retryablehttp` package nicely wraps _everything_ to `url.Error`.
		return nil, fmt.Errorf("failed request: %w", err)
	}
	defer func() {
		if ferr := resp.Body.Close(); ferr != nil {
			err = fmt.Errorf("failed to close: %w", ferr)
		}
	}()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response body: %w", err)
	}
	headers := c.createDebugHeaders(resp.Header, "")
	log.Printf("[DEBUG] %s %s %s <- %s %s", resp.Status, headers, c.redactedDump(body), method, strings.ReplaceAll(requestURL, "\n", ""))
	return body, nil
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

func (c *ApiClient) fixHost() {
	if c.Config.Host != "" && !(strings.HasPrefix(c.Config.Host, "https://") || strings.HasPrefix(c.Config.Host, "http://")) {
		// azurerm_databricks_workspace.*.workspace_url is giving URL without scheme
		// so that is why this line is here
		c.Config.Host = "https://" + c.Config.Host
	}
}

func onlyNBytes(j string, numBytes int) string {
	diff := len([]byte(j)) - numBytes
	if diff > 0 {
		return fmt.Sprintf("%s... (%d more bytes)", j[:numBytes], diff)
	}
	return j
}

func (c *ApiClient) parseUnknownError(status string, body []byte, err error) (errorBody APIErrorBody) {
	// this is most likely HTML... since un-marshalling JSON failed
	// Status parts first in case html message is not as expected
	statusParts := strings.SplitN(status, " ", 2)
	if len(statusParts) < 2 {
		errorBody.ErrorCode = "UNKNOWN"
	} else {
		errorBody.ErrorCode = strings.ReplaceAll(
			strings.ToUpper(strings.Trim(statusParts[1], " .")),
			" ", "_")
	}
	stringBody := string(body)
	messageRE := regexp.MustCompile(`<pre>(.*)</pre>`)
	messageMatches := messageRE.FindStringSubmatch(stringBody)
	// No messages with <pre> </pre> format found so return a APIError
	if len(messageMatches) < 2 {
		errorBody.Message = fmt.Sprintf("Response from server (%s) %s: %v",
			status, stringBody, err)
		return
	}
	errorBody.Message = strings.Trim(messageMatches[1], " .")
	return
}

func (c *ApiClient) parseError(resp *http.Response) APIError {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIError{
			Message:    err.Error(),
			ErrorCode:  "IO_READ",
			StatusCode: resp.StatusCode,
			Resource:   resp.Request.URL.Path,
		}
	}
	log.Printf("[DEBUG] %s %v", resp.Status, c.redactedDump(body))
	// try to read in nicely formatted API error response
	var errorBody APIErrorBody
	err = json.Unmarshal(body, &errorBody)
	if err != nil {
		errorBody = c.parseUnknownError(resp.Status, body, err)
	}
	if errorBody.API12Error != "" {
		// API 1.2 has different response format, let's adapt
		errorBody.Message = errorBody.API12Error
	}
	// Handle SCIM error message details
	if errorBody.Message == "" && errorBody.ScimDetail != "" {
		if errorBody.ScimDetail == "null" {
			errorBody.Message = "SCIM API Internal Error"
		} else {
			errorBody.Message = errorBody.ScimDetail
		}
		// add more context from SCIM responses
		errorBody.Message = fmt.Sprintf("%s %s", errorBody.ScimType, errorBody.Message)
		errorBody.Message = strings.Trim(errorBody.Message, " ")
		errorBody.ErrorCode = fmt.Sprintf("SCIM_%s", errorBody.ScimStatus)
	}
	// if resp.StatusCode == 403 {
	// 	errorBody.Message = fmt.Sprintf("%s. Using %s auth: %s",
	// 		strings.Trim(errorBody.Message, "."), c.AuthType,
	// 		c.configDebugString())
	// }
	return APIError{
		Message:    errorBody.Message,
		ErrorCode:  errorBody.ErrorCode,
		StatusCode: resp.StatusCode,
		Resource:   resp.Request.URL.Path,
	}
}

// checkHTTPRetry inspects HTTP errors from the Databricks API for known transient errors on Workspace creation
func (c *ApiClient) checkHTTPRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// TODO: port baseRetryPolicy from vendor/github.com/hashicorp/go-retryablehttp/client.go
	if ue, ok := err.(*url.Error); ok {
		apiError := APIError{
			ErrorCode:  "IO_ERROR",
			StatusCode: 523,
			Message:    ue.Error(),
		}
		return apiError.IsRetriable(), apiError
	}
	if resp == nil {
		// If response is nil we can't make retry choices.
		// In this case don't retry and return the original error from httpclient
		return false, err
	}
	if resp.StatusCode == 429 {
		return true, APIError{
			ErrorCode:  "TOO_MANY_REQUESTS",
			Message:    "Current request has to be retried",
			StatusCode: 429,
		}
	}
	if resp.StatusCode >= 400 {
		apiError := c.parseError(resp)
		return apiError.IsRetriable(), apiError
	}
	return false, nil
}
