package config

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type hc func(r *http.Request) (*http.Response, error)

func (cb hc) RoundTrip(r *http.Request) (*http.Response, error) {
	return cb(r)
}

func (cb hc) SkipRetryOnIO() bool {
	return true
}

func TestApiClient_RetriesGetPermissionsOnGatewayTimeout(t *testing.T) {
	requestCount := 0
	c := &Config{
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			initialRequestCount := requestCount
			requestCount++
			if initialRequestCount == 0 {
				return &http.Response{
					Request:    r,
					StatusCode: http.StatusGatewayTimeout,
					Body: io.NopCloser(strings.NewReader(
						fmt.Sprintf(`{"error_code":"TEMPORARILY_UNAVAILABLE", "message":"The service at %s is taking too long to process your request. Please try again later or try a faster operation."}`, r.URL))),
				}, nil
			}
			return &http.Response{
				Request:    r,
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"permissions": ["can_run_queries"]}`)),
			}, nil
		}),
	}
	client, err := c.NewApiClient()
	require.NoError(t, err)
	ctx := context.Background()
	var res map[string][]string
	err = client.Do(ctx, "GET", "/api/2.0/permissions/object/id", httpclient.WithResponseUnmarshal(&res))
	assert.NoError(t, err)
	assert.Equal(t, map[string][]string{"permissions": {"can_run_queries"}}, res)
}
