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

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
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
	err := c.Do(context.Background(), "GET", "/a/b",
		WithRequestHeaders(map[string]string{
			"e": "f",
		}), WithRequestData(map[string]string{
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
	err := c.Do(context.Background(), "GET", "/a/b",
		WithRequestHeaders(map[string]string{
			"e": "f",
		}), WithRequestData(map[string]string{
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
	err := c.Do(context.Background(), "POST", "/c",
		WithRequestData(Dummy{1}),
		WithResponseUnmarshal(&resp))
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
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestData(Dummy{1}),
		WithResponseUnmarshal(&resp))
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
	require.True(t, retried[0], "request was not retried")
}

func TestSimpleRequestNotRetried(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	var tried, retried bool
	transportErr := &url.Error{
		Op:  "open",
		URL: "/a/b",
		Err: fmt.Errorf("some other reason"),
	}

	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			if !tried {
				tried = true
				return nil, transportErr
			}
			retried = true
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	var resp Dummy
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestData(Dummy{1}),
		WithResponseUnmarshal(&resp))
	require.Error(t, err)
	require.ErrorIs(t, err, transportErr)
	require.True(t, tried, "request was not tried")
	require.False(t, retried, "request was retried")
}

func TestHaltAttemptForLimit(t *testing.T) {
	ctx := context.Background()
	c := &ApiClient{
		config:      ClientConfig{},
		rateLimiter: &rate.Limiter{},
	}
	req, err := common.NewRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "GET", "foo", req)()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, "rate: Wait(n=1) exceeds limiter's burst 0")
}

func TestHaltAttemptForNewRequest(t *testing.T) {
	ctx := context.Background()
	c := NewApiClient(ClientConfig{})
	req, err := common.NewRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "🥱", "/", req)()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, `net/http: invalid method "🥱"`)
}

func TestHaltAttemptForVisitor(t *testing.T) {
	ctx := context.Background()
	c := NewApiClient(ClientConfig{})
	req, err := common.NewRequestBody([]byte{})
	require.NoError(t, err)
	_, rerr := c.attempt(ctx, "GET", "/", req,
		func(r *http.Request) error {
			return fmt.Errorf("🥱")
		})()
	require.NotNil(t, rerr)
	require.Equal(t, true, rerr.Halt)
	require.EqualError(t, rerr.Err, "🥱")
}

func TestFailPerformChannel(t *testing.T) {
	ctx := context.Background()
	c := &ApiClient{
		rateLimiter: rate.NewLimiter(rate.Inf, 1),
	}
	_, err := c.perform(ctx, "GET", "/", true)
	require.EqualError(t, err, "request marshal: unsupported query string data: true")
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
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{}))
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
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{}))
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
	err := c.Do(context.Background(), "PATCH", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{}))
	require.NoError(t, err, "response body should not be closed for streaming responses")
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
	err := c.Do(context.Background(), "GET", "/a",
		WithRequestData(map[string]any{
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
	err := c.Do(context.Background(), "GET", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{
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
	err := c.Do(context.Background(), "GET", "/a",
		WithRequestHeaders(headers),
		WithRequestData(map[string]any{
			"b": 0,
			"a": 3,
			"c": 23,
		}))
	require.NoError(t, err)

	require.Equal(t, `[DEBUG] GET /a?a=3&b=0&c=23
<  
< <Streaming response>`, bufLogger.String())
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
	err = client.Do(context.Background(), "POST", "/a",
		WithRequestData(r),
		WithResponseUnmarshal(&respBytes))
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
	err := client.Do(context.Background(), "POST", "/a",
		WithRequestData(customReader{}))
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
	err := client.Do(context.Background(), "GET", "/a",
		WithResponseUnmarshal(&respBytes))
	require.NoError(t, err)
	require.Equal(t, "succeeded", respBytes.String())
	require.True(t, succeed)
}

func TestOAuth2Integration(t *testing.T) {
	inner := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 204,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		}),
	})

	ctx := context.Background()
	ctx = inner.InContextForOAuth2(ctx)

	outer := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: "abc",
	}))
	res, err := outer.Get("abc")
	require.NoError(t, err)
	require.Equal(t, 204, res.StatusCode)
}

type logEntry struct {
	String string
	Args   []interface{}
}

type inMemoryLogger struct {
	logs map[logger.Level][]logEntry
}

func newInMemoryLogger() *inMemoryLogger {
	return &inMemoryLogger{
		logs: map[logger.Level][]logEntry{},
	}
}

func (l *inMemoryLogger) Enabled(_ context.Context, level logger.Level) bool {
	return true
}

func (l *inMemoryLogger) Tracef(_ context.Context, format string, v ...interface{}) {
	l.logs[logger.LevelTrace] = append(l.logs[logger.LevelTrace], logEntry{format, v})
}

func (l *inMemoryLogger) Debugf(_ context.Context, format string, v ...interface{}) {
	l.logs[logger.LevelDebug] = append(l.logs[logger.LevelDebug], logEntry{format, v})
}

func (l *inMemoryLogger) Infof(_ context.Context, format string, v ...interface{}) {
	l.logs[logger.LevelInfo] = append(l.logs[logger.LevelInfo], logEntry{format, v})
}

func (l *inMemoryLogger) Warnf(_ context.Context, format string, v ...interface{}) {
	l.logs[logger.LevelWarn] = append(l.logs[logger.LevelWarn], logEntry{format, v})
}

func (l *inMemoryLogger) Errorf(_ context.Context, format string, v ...interface{}) {
	l.logs[logger.LevelError] = append(l.logs[logger.LevelError], logEntry{format, v})
}

func TestAuthorizationHeaderRedactedInLog(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
		DebugHeaders:          true,
		DebugSensitiveHeaders: false,
	})
	ctx := context.Background()
	log := newInMemoryLogger()
	ctx = logger.NewContext(ctx, log)
	err := c.Do(ctx, "POST", "/c",
		WithRequestHeader("Authorization", "Bearer secret-token"))
	assert.NoError(t, err)
	for _, logs := range log.logs {
		for _, logMessage := range logs {
			require.NotContains(t, logMessage.String, "Bearer secret-token")
		}
	}
}

func TestAuthorizationHeaderPresentInLog(t *testing.T) {
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
		DebugHeaders:          true,
		DebugSensitiveHeaders: true,
	})
	ctx := context.Background()
	log := newInMemoryLogger()
	ctx = logger.NewContext(ctx, log)
	err := c.Do(ctx, "POST", "/c",
		WithRequestHeader("Authorization", "Bearer secret-token"))
	assert.NoError(t, err)
	containsToken := false
out:
	for _, logs := range log.logs {
		for _, logMessage := range logs {
			if strings.Contains(logMessage.String, "Bearer secret-token") {
				containsToken = true
				break out
			}
		}
	}
	assert.True(t, containsToken)
}
