package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

// errReader(true) will also fail on Close
type errReader bool

func (errReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("test error")
}

func (i errReader) Close() error {
	if i {
		return fmt.Errorf("test error")
	}
	return nil
}

type hc func(r *http.Request) (*http.Response, error)

func (cb hc) Do(r *http.Request) (*http.Response, error) {
	return cb(r)
}

func (cb hc) CloseIdleConnections() {}

func TestNew(t *testing.T) {
	c, err := New(&config.Config{
		ConfigFile: "/dev/null",
	})
	assert.NoError(t, err)

	assert.Equal(t, 96, c.debugTruncateBytes)
	assert.Equal(t, 5*time.Minute, c.retryTimeout)
}

func TestSimpleRequestFailsURLError(t *testing.T) {
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			r.Header.Add("Authenticated", "yes")
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "GET", r.Method)
			assert.Equal(t, "/a/b", r.URL.Path)
			assert.Equal(t, "c=d", r.URL.RawQuery)
			assert.Equal(t, "f", r.Header.Get("e"))
			auth := r.Header.Get("Authenticated")
			assert.Equal(t, "yes", auth)
			return &http.Response{
					Request: r,
				}, &url.Error{
					Op:  "GET",
					URL: "/a/b",
					Err: fmt.Errorf("nope"),
				}
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	err := c.Do(context.Background(), "GET", "/a/b", map[string]string{
		"e": "f",
	}, map[string]string{
		"c": "d",
	}, nil)
	assert.EqualError(t, err, "GET \"/a/b\": nope")
}

func TestSimpleRequestFailsAPIError(t *testing.T) {
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			r.Header.Add("Authenticated", "yes")
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "GET", r.Method)
			assert.Equal(t, "/a/b", r.URL.Path)
			assert.Equal(t, "c=d", r.URL.RawQuery)
			assert.Equal(t, "f", r.Header.Get("e"))
			auth := r.Header.Get("Authenticated")
			assert.Equal(t, "yes", auth)
			return &http.Response{
				StatusCode: 400,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader(`{"error_code": "INVALID_PARAMETER_VALUE", "message": "nope"}`)),
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	err := c.Do(context.Background(), "GET", "/a/b", map[string]string{
		"e": "f",
	}, map[string]string{
		"c": "d",
	}, nil)
	assert.EqualError(t, err, "nope")
}

func TestETag(t *testing.T) {
	reason := "some_reason"
	domain := "a_domain"
	eTag := "sample_etag"
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Request:    r,
				Body: io.NopCloser(strings.NewReader(fmt.Sprintf(`{
					"error_code": "RESOURCE_CONFLICT",
					"message": "test_public_workspace_setting",
					"stack_trace": "java.io.PrintWriter@329e4ed3",
					"details": [
					  {
						"@type": "%s",
						"reason": "%s",
						"domain": "%s",
						"metadata": {
						  "etag": "%s"
						}
					  },
					  {
						"@type": "anotherType",
						"reason": "",
						"domain": "",
						"metadata": {
						  "etag": "anotherTag"
						}
					  }
					]
				  }`, "type.googleapis.com/google.rpc.ErrorInfo", reason, domain, eTag))),
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	err := c.Do(context.Background(), "GET", "/a/b", map[string]string{
		"e": "f",
	}, map[string]string{
		"c": "d",
	}, nil)
	details := apierr.GetErrorInfo(err)
	assert.Equal(t, 1, len(details))
	errorDetails := details[0]
	assert.Equal(t, reason, errorDetails.Reason)
	assert.Equal(t, domain, errorDetails.Domain)
	assert.Equal(t, map[string]string{
		"etag": eTag,
	}, errorDetails.Metadata)
}

func TestSimpleRequestSucceeds(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	var resp Dummy
	err := c.Do(context.Background(), "POST", "/c", nil, Dummy{1}, &resp)
	assert.NoError(t, err)
	assert.Equal(t, 2, resp.Foo)
}

func TestSimpleRequestRetried(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	var retried [1]bool
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			if !retried[0] {
				retried[0] = true
				return nil, &url.Error{
					Op:  "open",
					URL: "/a/b",
					Err: fmt.Errorf("connection refused"),
				}
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
		retryTimeout: 1 * time.Minute,
		rateLimiter:  rate.NewLimiter(rate.Inf, 1),
	}
	var resp Dummy
	err := c.Do(context.Background(), "PATCH", "/a", nil, Dummy{1}, &resp)
	assert.NoError(t, err)
	assert.Equal(t, 2, resp.Foo)
	assert.True(t, retried[0], "request was not retried")
}

func TestHaltAttemptForLimit(t *testing.T) {
	ctx := context.Background()
	c := &DatabricksClient{
		rateLimiter: &rate.Limiter{},
	}
	req, err := newRequestBody([]byte{})
	assert.NoError(t, err)
	_, rerr := c.attempt(ctx, "GET", "foo", nil, req)()
	assert.NotNil(t, rerr)
	assert.Equal(t, true, rerr.Halt)
	assert.EqualError(t, rerr.Err, "rate: Wait(n=1) exceeds limiter's burst 0")
}

func TestHaltAttemptForNewRequest(t *testing.T) {
	ctx := context.Background()
	c := &DatabricksClient{
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	req, err := newRequestBody([]byte{})
	assert.NoError(t, err)
	_, rerr := c.attempt(ctx, "🥱", "/", nil, req)()
	assert.NotNil(t, rerr)
	assert.Equal(t, true, rerr.Halt)
	assert.EqualError(t, rerr.Err, `net/http: invalid method "🥱"`)
}

func TestHaltAttemptForVisitor(t *testing.T) {
	ctx := context.Background()
	c := &DatabricksClient{
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	req, err := newRequestBody([]byte{})
	assert.NoError(t, err)
	_, rerr := c.attempt(ctx, "GET", "/", nil, req,
		func(r *http.Request) error {
			return fmt.Errorf("🥱")
		})()
	assert.NotNil(t, rerr)
	assert.Equal(t, true, rerr.Halt)
	assert.EqualError(t, rerr.Err, "🥱")
}

func TestMakeRequestBody(t *testing.T) {
	type x struct {
		Scope string `json:"scope" url:"scope"`
	}
	requestURL := "/a/b/c"
	body, err := makeRequestBody("GET", &requestURL, x{"test"})
	assert.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	assert.NoError(t, err)
	assert.Equal(t, "/a/b/c?scope=test", requestURL)
	assert.Equal(t, 0, len(bodyBytes))

	requestURL = "/a/b/c"
	body, err = makeRequestBody("POST", &requestURL, x{"test"})
	assert.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	assert.NoError(t, err)
	assert.Equal(t, "/a/b/c", requestURL)
	x1 := `{"scope":"test"}`
	assert.Equal(t, []byte(x1), bodyBytes)
}

func TestMakeRequestBodyFromReader(t *testing.T) {
	requestURL := "/a/b/c"
	body, err := makeRequestBody("PUT", &requestURL, strings.NewReader("abc"))
	assert.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	assert.NoError(t, err)
	assert.Equal(t, []byte("abc"), bodyBytes)
}

func TestMakeRequestBodyReaderError(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("POST", &requestURL, errReader(false))
	// The request body is only read once the request is sent, so no error
	// should be returned until then.
	assert.NoError(t, err, "request body reader error should be ignored")
}

func TestMakeRequestBodyJsonError(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo chan string `json:"foo"`
	}
	_, err := makeRequestBody("POST", &requestURL, x{make(chan string)})
	assert.EqualError(t, err, "request marshal failure: json: unsupported type: chan string")
}

type failingUrlEncode string

func (fue failingUrlEncode) EncodeValues(key string, v *url.Values) error {
	return fmt.Errorf(string(fue))
}

func TestMakeRequestBodyQueryFailingEncode(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo failingUrlEncode `url:"foo"`
	}
	_, err := makeRequestBody("GET", &requestURL, x{failingUrlEncode("always failing")})
	assert.EqualError(t, err, "cannot create query string: always failing")
}

func TestMakeRequestBodyQueryUnsupported(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("GET", &requestURL, true)
	assert.EqualError(t, err, "unsupported query string data: true")
}

func TestFailPerformChannel(t *testing.T) {
	ctx := context.Background()
	c := &DatabricksClient{
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	_, err := c.perform(ctx, "GET", "/", nil, true)
	assert.EqualError(t, err, "request marshal: unsupported query string data: true")
}

func TestSimpleRequestAPIError(t *testing.T) {
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Body: io.NopCloser(strings.NewReader(`{
					"error_code": "NOT_FOUND",
					"message": "Something was not found"
				}`)),
				Request: r,
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	err := c.Do(context.Background(), "PATCH", "/a", nil, map[string]any{}, nil)
	var aerr *apierr.APIError
	if assert.ErrorAs(t, err, &aerr) {
		assert.Equal(t, "NOT_FOUND", aerr.ErrorCode)
	}
}

func TestSimpleRequestErrReaderBody(t *testing.T) {
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(false),
				Request:    r,
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "PATCH", "/a", headers, map[string]any{}, nil)
	assert.EqualError(t, err, "response body: test error")
}

func TestSimpleRequestErrReaderBodyStreamResponse(t *testing.T) {
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(false),
				Request:    r,
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "PATCH", "/a", headers, map[string]any{}, nil)
	assert.NoError(t, err, "streaming response bodies are not read")
}

func TestSimpleRequestErrReaderCloseBody(t *testing.T) {
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(true),
				Request:    r,
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "PATCH", "/a", headers, map[string]any{}, nil)
	assert.EqualError(t, err, "response body: test error")
}

func TestSimpleRequestErrReaderCloseBody_StreamResponse(t *testing.T) {
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(true),
				Request:    r,
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "PATCH", "/a", headers, map[string]any{}, nil)
	assert.NoError(t, err, "response body should not be closed for streaming responses")
}

func TestSimpleRequestRawResponse(t *testing.T) {
	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader("Hello, world!")),
				Request:    r,
			}, nil
		}),
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	var raw []byte
	err := c.Do(context.Background(), "GET", "/a", nil, nil, &raw)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, world!", string(raw))
}

type BufferLogger struct {
	strings.Builder
}

func (l *BufferLogger) Enabled(_ context.Context, level logger.Level) bool {
	return true
}

func (l *BufferLogger) Tracef(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[TRACE] "+format, v...))
}

func (l *BufferLogger) Debugf(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[DEBUG] "+format, v...))
}

func (l *BufferLogger) Infof(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[INFO] "+format, v...))
}

func (l *BufferLogger) Warnf(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[WARN] "+format, v...))
}

func (l *BufferLogger) Errorf(_ context.Context, format string, v ...interface{}) {
	l.WriteString(fmt.Sprintf("[ERROR] "+format, v...))
}

func TestSimpleResponseRedaction(t *testing.T) {
	cfg := config.NewMockConfig(func(r *http.Request) error {
		r.Header.Add("X-For-Logging", "yes")
		return nil
	})
	cfg.Host = "http://localhost:12345"

	prevLogger := logger.DefaultLogger
	bufLogger := &BufferLogger{}
	logger.DefaultLogger = bufLogger
	defer func() {
		logger.DefaultLogger = prevLogger
	}()

	c := &DatabricksClient{
		Config: cfg,
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Proto:      "HTTP/3.4",
				Status:     "200 Fine",
				Body: io.NopCloser(strings.NewReader(`{
					"string_value": "__SENSITIVE01__",
					"inner": {
					  "token_value": "__SENSITIVE02__",
					  "content": "__SENSITIVE03__"
					},
					"list": [
					  {
					    "token_value": "__SENSITIVE04__"
					  }
					],
					"longer": "12345678901234567890qwerty"
				}`)),
				Request: r,
			}, nil
		}),
		debugTruncateBytes: 16,
		debugHeaders:       true,
		rateLimiter:        rate.NewLimiter(rate.Inf, 1),
	}
	err := c.Do(context.Background(), "GET", "/a", nil, map[string]any{
		"b": 0,
		"a": 3,
		"c": 23,
	}, nil)
	assert.NoError(t, err)
	// not testing for exact logged lines, as header order is not deterministic
	assert.NotContains(t, bufLogger.String(), "__SENSITIVE01__")
	assert.NotContains(t, bufLogger.String(), "__SENSITIVE02__")
	assert.NotContains(t, bufLogger.String(), "__SENSITIVE03__")
	assert.NotContains(t, bufLogger.String(), "__SENSITIVE04__")
	assert.NotContains(t, bufLogger.String(), "12345678901234567890qwerty")
}

func TestInlineArrayDebugging(t *testing.T) {
	prevLogger := logger.DefaultLogger
	bufLogger := &BufferLogger{}
	logger.DefaultLogger = bufLogger
	defer func() {
		logger.DefaultLogger = prevLogger
	}()

	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`[
					{"foo": "bar"}
				]`)),
				Request: r,
			}, nil
		}),
		debugTruncateBytes: 2048,
		rateLimiter:        rate.NewLimiter(rate.Inf, 1),
	}
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "GET", "/a", headers, map[string]any{
		"b": 0,
		"a": 3,
		"c": 23,
	}, nil)
	assert.NoError(t, err)

	assert.Equal(t, `[DEBUG] GET /a?a=3&b=0&c=23
<  
< [
<   {
<     "foo": "bar"
<   }
< ]`, bufLogger.String())
}

func TestInlineArrayDebugging_StreamResponse(t *testing.T) {
	prevLogger := logger.DefaultLogger
	bufLogger := &BufferLogger{}
	logger.DefaultLogger = bufLogger
	defer func() {
		logger.DefaultLogger = prevLogger
	}()

	c := &DatabricksClient{
		Config: config.NewMockConfig(func(r *http.Request) error {
			return nil
		}),
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`lots of bytes`)),
				Request:    r,
			}, nil
		}),
		debugTruncateBytes: 2048,
		rateLimiter:        rate.NewLimiter(rate.Inf, 1),
	}
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "GET", "/a", headers, map[string]any{
		"b": 0,
		"a": 3,
		"c": 23,
	}, nil)
	assert.NoError(t, err)

	assert.Equal(t, `[DEBUG] GET /a?a=3&b=0&c=23
<  
< [non-JSON document of 15 bytes]. <io.ReadCloser>`, bufLogger.String())
}

func TestStreamRequestFromFileWithReset(t *testing.T) {
	// make a temporary file with some content
	f, err := os.CreateTemp("", "databricks-client-test")
	assert.NoError(t, err)
	defer os.Remove(f.Name())
	_, err = f.WriteString("hello world")
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	// Make a reader that reads this file
	r, err := os.Open(f.Name())
	assert.NoError(t, err)
	defer r.Close()

	succeed := false
	handler := func(req *http.Request) (*http.Response, error) {
		bytes, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Equal(t, "hello world", string(bytes))
		if succeed {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader("succeeded")),
				Request:    req,
			}, nil
		}
		succeed = true
		return &http.Response{
			StatusCode: 429,
			Body:       io.NopCloser(strings.NewReader("failed")),
			Request:    req,
		}, nil
	}

	client := &DatabricksClient{
		httpClient:   hc(handler),
		rateLimiter:  rate.NewLimiter(rate.Limit(1), 1),
		Config:       config.NewMockConfig(func(r *http.Request) error { return nil }),
		retryTimeout: time.Hour,
	}

	respBytes := bytes.Buffer{}
	err = client.Do(context.Background(), "POST", "/a", nil, r, &respBytes)
	assert.NoError(t, err)
	assert.Equal(t, "succeeded", respBytes.String())
	assert.True(t, succeed)
}

type customReader struct{}

func (c customReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func TestCannotRetryArbitraryReader(t *testing.T) {
	client := &DatabricksClient{
		httpClient: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 429,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		}),
		rateLimiter:  rate.NewLimiter(rate.Limit(1), 1),
		Config:       config.NewMockConfig(func(r *http.Request) error { return nil }),
		retryTimeout: time.Hour,
	}
	err := client.Do(context.Background(), "POST", "/a", nil, customReader{}, nil)
	assert.ErrorContains(t, err, "cannot reset reader of type client.customReader")
}

func TestRetryGetRequest(t *testing.T) {
	// This test was added in response to https://github.com/databricks/terraform-provider-databricks/issues/2675.
	succeed := false
	handler := func(req *http.Request) (*http.Response, error) {
		assert.Nil(t, req.Body)

		if succeed {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader("succeeded")),
				Request:    req,
			}, nil
		}

		succeed = true
		return &http.Response{
			StatusCode: 429,
			Body:       io.NopCloser(strings.NewReader("failed")),
			Request:    req,
		}, nil
	}

	client := &DatabricksClient{
		httpClient:   hc(handler),
		rateLimiter:  rate.NewLimiter(rate.Limit(1), 1),
		Config:       config.NewMockConfig(func(r *http.Request) error { return nil }),
		retryTimeout: time.Hour,
	}

	respBytes := bytes.Buffer{}
	err := client.Do(context.Background(), "GET", "/a", nil, nil, &respBytes)
	assert.NoError(t, err)
	assert.Equal(t, "succeeded", respBytes.String())
	assert.True(t, succeed)
}
