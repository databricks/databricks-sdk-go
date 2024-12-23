package httpclient

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mock struct {
	MaxFails     int            // number of times the failed Response is returned
	FailResponse *http.Response // response to return in case of fail
	FailError    error          // error to return in case of fail
	NumCalls     int            // total number of calls
}

func (m *mock) RoundTrip(r *http.Request) (*http.Response, error) {
	m.NumCalls++
	if m.NumCalls <= m.MaxFails {
		resp := *m.FailResponse
		resp.Request = r
		return &resp, m.FailError
	}
	return &http.Response{
		Request:    r,
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{}`)),
	}, nil
}

func (m *mock) SkipRetryOnIO() bool {
	return true
}

func TestApiClient_Do_retries(t *testing.T) {
	testCases := []struct {
		desc         string
		mock         *mock
		errorRetrier ErrorRetrier
		wantErrorMsg string
		wantNumCalls int
	}{
		{
			desc: "default retrier retries on 429",
			mock: &mock{
				MaxFails:     1,
				FailResponse: &http.Response{StatusCode: http.StatusGatewayTimeout},
			},
			wantNumCalls: 2,
		},
		{
			desc: "default retrier retries on 504",
			mock: &mock{
				MaxFails:     1,
				FailResponse: &http.Response{StatusCode: http.StatusGatewayTimeout},
			},
			wantNumCalls: 2,
		},
		{
			desc: "default retrier does not retry on 503",
			mock: &mock{
				MaxFails:     1,
				FailResponse: &http.Response{StatusCode: http.StatusServiceUnavailable},
			},
			wantErrorMsg: "http 503: ",
			wantNumCalls: 1,
		},
		{
			desc: "no retry when ErrorRetriable returns false",
			mock: &mock{
				MaxFails:     1,
				FailResponse: &http.Response{StatusCode: http.StatusGatewayTimeout},
			},
			errorRetrier: func(context.Context, *http.Request, *common.ResponseWrapper, error) bool {
				return false
			},
			wantErrorMsg: "http 504: ",
			wantNumCalls: 1,
		},
		{
			desc: "retry 1 time",
			mock: &mock{
				MaxFails:     1,
				FailResponse: &http.Response{StatusCode: http.StatusGatewayTimeout},
			},
			errorRetrier: func(context.Context, *http.Request, *common.ResponseWrapper, error) bool {
				return true
			},
			wantNumCalls: 2,
		},
		{
			desc: "retry multiple times",
			mock: &mock{
				MaxFails:     3,
				FailResponse: &http.Response{StatusCode: http.StatusGatewayTimeout},
			},
			errorRetrier: func(_ context.Context, _ *http.Request, _ *common.ResponseWrapper, _ error) bool {
				return true
			},
			wantNumCalls: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			clientConfig := ClientConfig{
				Transport:      tc.mock,
				ErrorRetriable: tc.errorRetrier,
				RetryBackoff: func(_ int) time.Duration {
					// Do not wait to retry in tests
					return 0
				},
			}
			client := NewApiClient(clientConfig)

			err := client.Do(context.Background(), "GET", "test-path")
			if tc.wantErrorMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.wantErrorMsg)
			}
			gotNumCalls := tc.mock.NumCalls

			if gotNumCalls != tc.wantNumCalls {
				t.Errorf("got %d calls, want %d", gotNumCalls, tc.wantNumCalls)
			}
		})
	}
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
	err := c.Do(context.Background(), "PATCH", "/a", WithRequestData(map[string]any{}))
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
	err := c.Do(context.Background(), "PATCH", "/a", WithRequestHeaders(headers), WithRequestData(map[string]any{}))
	require.EqualError(t, err, "response body: test error")
}
