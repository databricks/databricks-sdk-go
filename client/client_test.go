package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type hc func(r *http.Request) (*http.Response, error)

func (cb hc) RoundTrip(r *http.Request) (*http.Response, error) {
	return cb(r)
}

func TestSimpleRequestFailsURLError(t *testing.T) {
	c, err := New(&config.Config{
		Host:                "some",
		Token:               "token",
		ConfigFile:          "/dev/null",
		RetryTimeoutSeconds: 1,
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "GET", r.Method)
			require.Equal(t, "/a/b", r.URL.Path)
			require.Equal(t, "c=d", r.URL.RawQuery)
			require.Equal(t, "f", r.Header.Get("e"))
			auth := r.Header.Get("Authorization")
			require.Equal(t, "Bearer token", auth)
			return nil, fmt.Errorf("nope")
		}),
	})
	require.NoError(t, err)
	err = c.Do(context.Background(), "GET", "/a/b", map[string]string{
		"e": "f",
	}, map[string]string{
		"c": "d",
	}, nil)
	require.EqualError(t, err, `Get "https://some/a/b?c=d": nope`)
}

func TestSimpleRequestFailsAPIError(t *testing.T) {
	c, err := New(&config.Config{
		Host:                "some",
		Token:               "token",
		ConfigFile:          "/dev/null",
		RetryTimeoutSeconds: 1,
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader(`{"error_code": "INVALID_PARAMETER_VALUE", "message": "nope"}`)),
			}, nil
		}),
	})
	require.NoError(t, err)
	err = c.Do(context.Background(), "GET", "/a/b", map[string]string{
		"e": "f",
	}, map[string]string{
		"c": "d",
	}, nil)
	require.EqualError(t, err, "nope")
	require.ErrorIs(t, err, apierr.ErrInvalidParameterValue)
}

func TestETag(t *testing.T) {
	reason := "some_reason"
	domain := "a_domain"
	eTag := "sample_etag"
	c, err := New(&config.Config{
		Host:                "some",
		Token:               "token",
		ConfigFile:          "/dev/null",
		RetryTimeoutSeconds: 1,
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
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
	})
	require.NoError(t, err)
	err = c.Do(context.Background(), "GET", "/a/b", map[string]string{
		"e": "f",
	}, map[string]string{
		"c": "d",
	}, nil)
	details := apierr.GetErrorInfo(err)
	require.Equal(t, 1, len(details))
	errorDetails := details[0]
	require.Equal(t, reason, errorDetails.Reason)
	require.Equal(t, domain, errorDetails.Domain)
	require.Equal(t, map[string]string{
		"etag": eTag,
	}, errorDetails.Metadata)
}

func TestSimpleRequestSucceeds(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)
	var resp Dummy
	err = c.Do(context.Background(), "POST", "/c", nil, Dummy{1}, &resp)
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
}

func TestSimpleRequestRetried(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	var retried [1]bool
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
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
	require.NoError(t, err)
	var resp Dummy
	err = c.Do(context.Background(), "PATCH", "/a", nil, Dummy{1}, &resp)
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
	require.True(t, retried[0], "request was not retried")
}

func TestSimpleRequestAPIError(t *testing.T) {
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body: io.NopCloser(strings.NewReader(`{
						"error_code": "NOT_FOUND",
						"message": "Something was not found"
					}`)),
				Request: r,
			}, nil
		}),
	})
	require.NoError(t, err)
	err = c.Do(context.Background(), "PATCH", "/a", nil, map[string]any{}, nil)
	var aerr *apierr.APIError
	require.ErrorAs(t, err, &aerr)
	require.Equal(t, "NOT_FOUND", aerr.ErrorCode)
	require.ErrorIs(t, err, apierr.ErrNotFound)
}

func TestHttpTransport(t *testing.T) {
	calledMock := false
	cfg := config.NewMockConfig(func(r *http.Request) error { return nil })
	cfg.HTTPTransport = hc(func(r *http.Request) (*http.Response, error) {
		calledMock = true
		return &http.Response{Request: r}, nil
	})
	client, err := New(cfg)
	require.NoError(t, err)

	err = client.Do(context.Background(), "GET", "/a", nil, nil, bytes.Buffer{})
	require.NoError(t, err)
	require.True(t, calledMock)
}

func TestDoRemovesDoubleSlashesFromFilesAPI(t *testing.T) {
	i := 0
	expectedPaths := []string{
		"/api/2.0/fs/files/Volumes/abc/def/ghi",
		"/api/2.0/anotherservice//test",
	}
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, expectedPaths[i], r.URL.Path)
			i++
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{}`)),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)
	err = c.Do(context.Background(), "GET", "/api/2.0/fs/files//Volumes/abc/def/ghi", nil, map[string]any{}, nil)
	require.NoError(t, err)
	err = c.Do(context.Background(), "GET", "/api/2.0/anotherservice//test", nil, map[string]any{}, nil)
	require.NoError(t, err)
}
