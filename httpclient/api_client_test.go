package httpclient

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

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func (cb hc) RoundTrip(r *http.Request) (*http.Response, error) {
	return cb(r)
}

func TestNew(t *testing.T) {
	c := NewApiClient(ClientConfig{})

	require.Equal(t, 96, c.config.DebugTruncateBytes)
	require.Equal(t, 5*time.Minute, c.config.RetryTimeout)
}

func TestSimpleRequestFailsURLError(t *testing.T) {
	c := NewApiClient(ClientConfig{
		RetryTimeout: 1 * time.Millisecond,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "GET", r.Method)
			require.Equal(t, "/a/b", r.URL.Path)
			require.Equal(t, "c=d", r.URL.RawQuery)
			require.Equal(t, "f", r.Header.Get("e"))
			return nil, fmt.Errorf("nope")
		}),
	})
	err := c.Do(context.Background(), "GET", "/a/b", WithHeaders(map[string]string{
		"e": "f",
	}), WithData(map[string]string{
		"c": "d",
	}))
	require.EqualError(t, err, `Get "/a/b?c=d": nope`)
}

func TestSimpleRequestFailsAPIError(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "GET", r.Method)
			require.Equal(t, "/a/b", r.URL.Path)
			require.Equal(t, "c=d", r.URL.RawQuery)
			require.Equal(t, "f", r.Header.Get("e"))
			return &http.Response{
				StatusCode: 400,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader(`nope`)),
			}, nil
		}),
	})
	err := c.Do(context.Background(), "GET", "/a/b", WithHeaders(map[string]string{
		"e": "f",
	}), WithData(map[string]string{
		"c": "d",
	}))
	require.EqualError(t, err, "http 400: nope")
}

func TestSimpleRequestSucceeds(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	var resp Dummy
	err := c.Do(context.Background(), "POST", "/c", WithData(Dummy{1}), WithUnmarshal(&resp))
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
}

func TestSimpleRequestRetried(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	var retried [1]bool
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
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
	})
	var resp Dummy
	err := c.Do(context.Background(), "PATCH", "/a", WithData(Dummy{1}), WithUnmarshal(&resp))
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
	require.True(t, retried[0], "request was not retried")
}

func TestHaltAttemptForLimit(t *testing.T) {
	ctx := context.Background()
	c := &ApiClient{
		config:      ClientConfig{},
		rateLimiter: &rate.Limiter{},
	}
	req, err := newRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "GET", "foo", req)()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, "rate: Wait(n=1) exceeds limiter's burst 0")
}

func TestHaltAttemptForNewRequest(t *testing.T) {
	ctx := context.Background()
	c := NewApiClient(ClientConfig{})
	req, err := newRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "🥱", "/", req)()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, `net/http: invalid method "🥱"`)
}

func TestHaltAttemptForVisitor(t *testing.T) {
	ctx := context.Background()
	c := NewApiClient(ClientConfig{})
	req, err := newRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "GET", "/", req,
		func(r *http.Request) error {
			return fmt.Errorf("🥱")
		})()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, "🥱")
}

func TestMakeRequestBody(t *testing.T) {
	type x struct {
		Scope string `json:"scope" url:"scope"`
	}
	requestURL := "/a/b/c"
	body, err := makeRequestBody("GET", &requestURL, x{"test"})
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))

	requestURL = "/a/b/c"
	body, err = makeRequestBody("POST", &requestURL, x{"test"})
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c", requestURL)
	x1 := `{"scope":"test"}`
	require.Equal(t, []byte(x1), bodyBytes)
}

func TestMakeRequestBodyFromReader(t *testing.T) {
	requestURL := "/a/b/c"
	body, err := makeRequestBody("PUT", &requestURL, strings.NewReader("abc"))
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, []byte("abc"), bodyBytes)
}

func TestMakeRequestBodyReaderError(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("POST", &requestURL, errReader(false))
	// The request body is only read once the request is sent, so no error
	// should be returned until then.
	require.NoError(t, err, "request body reader error should be ignored")
}

func TestMakeRequestBodyJsonError(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo chan string `json:"foo"`
	}
	_, err := makeRequestBody("POST", &requestURL, x{make(chan string)})
	require.EqualError(t, err, "request marshal failure: json: unsupported type: chan string")
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
	require.EqualError(t, err, "cannot create query string: always failing")
}

func TestMakeRequestBodyQueryUnsupported(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("GET", &requestURL, true)
	require.EqualError(t, err, "unsupported query string data: true")
}

func TestFailPerformChannel(t *testing.T) {
	ctx := context.Background()
	c := &ApiClient{
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	_, err := c.perform(ctx, "GET", "/", true)
	require.EqualError(t, err, "request marshal: unsupported query string data: true")
}

func TestSimpleRequestAPIError(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Body: io.NopCloser(strings.NewReader(`{
					"error_code": "NOT_FOUND",
					"message": "Something was not found"
				}`)),
				Request: r,
			}, nil
		}),
	})
	err := c.Do(context.Background(), "PATCH", "/a", WithData(map[string]any{}))
	var httpErr *HttpError
	if assert.ErrorAs(t, err, &httpErr) {
		require.Equal(t, 400, httpErr.StatusCode)
	}
}

func TestSimpleRequestErrReaderBody(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(false),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "PATCH", "/a", WithHeaders(headers), WithData(map[string]any{}))
	require.EqualError(t, err, "response body: test error")
}

func TestSimpleRequestErrReaderBodyStreamResponse(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(false),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "PATCH", "/a", WithHeaders(headers), WithData(map[string]any{}))
	require.NoError(t, err, "streaming response bodies are not read")
}

func TestSimpleRequestErrReaderCloseBody(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(true),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "PATCH", "/a", WithHeaders(headers), WithData(map[string]any{}))
	require.EqualError(t, err, "response body: test error")
}

func TestSimpleRequestErrReaderCloseBody_StreamResponse(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errReader(true),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "PATCH", "/a", WithHeaders(headers), WithData(map[string]any{}))
	require.NoError(t, err, "response body should not be closed for streaming responses")
}

func TestSimpleRequestRawResponse(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader("Hello, world!")),
				Request:    r,
			}, nil
		}),
	})
	var raw []byte
	err := c.Do(context.Background(), "GET", "/a", WithUnmarshal(&raw))
	require.NoError(t, err)
	require.Equal(t, "Hello, world!", string(raw))
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
	prevLogger := logger.DefaultLogger
	bufLogger := &BufferLogger{}
	logger.DefaultLogger = bufLogger
	defer func() {
		logger.DefaultLogger = prevLogger
	}()

	c := NewApiClient(ClientConfig{
		DebugTruncateBytes: 16,
		DebugHeaders:       true,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
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
	})
	err := c.Do(context.Background(), "GET", "/a", WithData(map[string]any{
		"b": 0,
		"a": 3,
		"c": 23,
	}))
	require.NoError(t, err)
	// not testing for exact logged lines, as header order is not deterministic
	require.NotContains(t, bufLogger.String(), "__SENSITIVE01__")
	require.NotContains(t, bufLogger.String(), "__SENSITIVE02__")
	require.NotContains(t, bufLogger.String(), "__SENSITIVE03__")
	require.NotContains(t, bufLogger.String(), "__SENSITIVE04__")
	require.NotContains(t, bufLogger.String(), "12345678901234567890qwerty")
}

func TestInlineArrayDebugging(t *testing.T) {
	prevLogger := logger.DefaultLogger
	bufLogger := &BufferLogger{}
	logger.DefaultLogger = bufLogger
	defer func() {
		logger.DefaultLogger = prevLogger
	}()

	c := NewApiClient(ClientConfig{
		DebugTruncateBytes: 2048,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`[
					{"foo": "bar"}
				]`)),
				Request: r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/json"}
	err := c.Do(context.Background(), "GET", "/a", WithHeaders(headers), WithData(map[string]any{
		"b": 0,
		"a": 3,
		"c": 23,
	}))
	require.NoError(t, err)

	require.Equal(t, `[DEBUG] GET /a?a=3&b=0&c=23
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

	c := NewApiClient(ClientConfig{
		DebugTruncateBytes: 2048,
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`lots of bytes`)),
				Request:    r,
			}, nil
		}),
	})
	headers := map[string]string{"Accept": "application/octet-stream"}
	err := c.Do(context.Background(), "GET", "/a", WithHeaders(headers), WithData(map[string]any{
		"b": 0,
		"a": 3,
		"c": 23,
	}))
	require.NoError(t, err)

	require.Equal(t, `[DEBUG] GET /a?a=3&b=0&c=23
<  
< [non-JSON document of 15 bytes]. <io.ReadCloser>`, bufLogger.String())
}

func TestStreamRequestFromFileWithReset(t *testing.T) {
	// make a temporary file with some content
	f, err := os.CreateTemp("", "databricks-client-test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	_, err = f.WriteString("hello world")
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// Make a reader that reads this file
	r, err := os.Open(f.Name())
	require.NoError(t, err)
	defer r.Close()

	succeed := false
	handler := func(req *http.Request) (*http.Response, error) {
		bytes, err := io.ReadAll(req.Body)
		require.NoError(t, err)
		require.Equal(t, "hello world", string(bytes))
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

	client := NewApiClient(ClientConfig{
		Transport: hc(handler),
	})

	respBytes := bytes.Buffer{}
	err = client.Do(context.Background(), "POST", "/a", WithData(r), WithUnmarshal(&respBytes))
	require.NoError(t, err)
	require.Equal(t, "succeeded", respBytes.String())
	require.True(t, succeed)
}

type customReader struct{}

func (c customReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func TestCannotRetryArbitraryReader(t *testing.T) {
	client := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 429,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		}),
	})
	err := client.Do(context.Background(), "POST", "/a", WithData(customReader{}))
	require.ErrorContains(t, err, "cannot reset reader of type httpclient.customReader")
}

func TestRetryGetRequest(t *testing.T) {
	// This test was added in response to https://github.com/databricks/terraform-provider-databricks/issues/2675.
	succeed := false
	handler := func(req *http.Request) (*http.Response, error) {
		require.Nil(t, req.Body)

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

	client := NewApiClient(ClientConfig{
		Transport: hc(handler),
	})

	respBytes := bytes.Buffer{}
	err := client.Do(context.Background(), "GET", "/a", WithUnmarshal(&respBytes))
	require.NoError(t, err)
	require.Equal(t, "succeeded", respBytes.String())
	require.True(t, succeed)
}
