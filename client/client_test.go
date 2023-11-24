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

func TestNew(t *testing.T) {
	_, err := New(&config.Config{
		ConfigFile: "/dev/null",
	})
	assert.NoError(t, err)
}

func TestSimpleRequestFailsURLError(t *testing.T) {
	cfg := config.NewMockConfig(func(r *http.Request) error {
		r.Header.Add("Authenticated", "yes")
		return nil
	})
	cfg.RetryTimeoutSeconds = 1
	c := newWithTransport(cfg, hc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/a/b", r.URL.Path)
		assert.Equal(t, "c=d", r.URL.RawQuery)
		assert.Equal(t, "f", r.Header.Get("e"))
		auth := r.Header.Get("Authenticated")
		assert.Equal(t, "yes", auth)
		return nil, fmt.Errorf("nope")
	}))
	err := c.Do(context.Background(), "GET", "/a/b", map[string]string{
		"e": "f",
	}, map[string]string{
		"c": "d",
	}, nil)
	assert.EqualError(t, err, `Get "/a/b?c=d": nope`)
}

func TestSimpleRequestFailsAPIError(t *testing.T) {
	c := *newWithTransport(config.NewMockConfig(func(r *http.Request) error {
		r.Header.Add("Authenticated", "yes")
		return nil
	}), hc(func(r *http.Request) (*http.Response, error) {
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
	}))
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
	c := newWithTransport(config.NewMockConfig(func(r *http.Request) error {
		return nil
	}), hc(func(r *http.Request) (*http.Response, error) {
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
	}))
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
	c := newWithTransport(config.NewMockConfig(func(r *http.Request) error {
		return nil
	}), hc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
			Request:    r,
		}, nil
	}))
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
	c := newWithTransport(config.NewMockConfig(func(r *http.Request) error {
		return nil
	}), hc(func(r *http.Request) (*http.Response, error) {
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
	}))
	var resp Dummy
	err := c.Do(context.Background(), "PATCH", "/a", nil, Dummy{1}, &resp)
	assert.NoError(t, err)
	assert.Equal(t, 2, resp.Foo)
	assert.True(t, retried[0], "request was not retried")
}

func TestSimpleRequestAPIError(t *testing.T) {
	c := newWithTransport(config.NewMockConfig(func(r *http.Request) error {
		return nil
	}), hc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
			Body: io.NopCloser(strings.NewReader(`{
					"error_code": "NOT_FOUND",
					"message": "Something was not found"
				}`)),
			Request: r,
		}, nil
	}))
	err := c.Do(context.Background(), "PATCH", "/a", nil, map[string]any{}, nil)
	var aerr *apierr.APIError
	if assert.ErrorAs(t, err, &aerr) {
		assert.Equal(t, "NOT_FOUND", aerr.ErrorCode)
	}
	assert.ErrorIs(t, err, apierr.ErrNotFound)
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

	assert.True(t, calledMock)
}
