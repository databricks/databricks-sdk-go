package httpclient

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleRequestRawResponse(t *testing.T) {
	c, err := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader("Hello, world!")),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)
	var raw []byte
	err = c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&raw))
	require.NoError(t, err)
	require.Equal(t, "Hello, world!", string(raw))
}

func TestWithResponseHeader(t *testing.T) {
	client, err := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				Request:    r,
				StatusCode: 204,
				Header: http.Header{
					"Foo": []string{"some"},
				},
				Body: io.NopCloser(strings.NewReader("")),
			}, nil
		}),
	})
	require.NoError(t, err)

	var out string
	ctx := context.Background()
	err = client.Do(ctx, "GET", "abc",
		WithResponseHeader("Foo", &out))
	require.NoError(t, err)
	require.Equal(t, "some", out)
}
