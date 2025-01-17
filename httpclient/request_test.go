package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestMakeRequestBody(t *testing.T) {
	type x struct {
		Scope string `json:"scope" url:"scope"`
	}
	requestURL := "/a/b/c"
	body, err := makeRequestBody("GET", &requestURL, x{"test"}, "", nil)
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))

	requestURL = "/a/b/c"
	body, err = makeRequestBody("POST", &requestURL, x{"test"}, "", nil)
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c", requestURL)
	x1 := `{"scope":"test"}`
	require.Equal(t, []byte(x1), bodyBytes)

	requestURL = "/a/b/c"
	body, err = makeRequestBody("HEAD", &requestURL, x{"test"}, "", nil)
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))
}

func TestMakeRequestBodyFromReader(t *testing.T) {
	requestURL := "/a/b/c"
	body, err := makeRequestBody("PUT", &requestURL, strings.NewReader("abc"), "", nil)
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, []byte("abc"), bodyBytes)
}

func TestUrlEncoding(t *testing.T) {
	data := GetOAuthTokenRequest{
		Assertion:            "assertion",
		AuthorizationDetails: "[{\"a\":\"b\"}]",
		GrantType:            "grant",
	}
	requestURL := "/a/b/c"
	body, err := makeRequestBody("POST", &requestURL, data, UrlEncodedContentType, nil)
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c", requestURL)
	require.Equal(t, string(bodyBytes), "assertion=assertion&authorization_details=%5B%7B%22a%22%3A%22b%22%7D%5D&grant_type=grant")
}

func TestMakeRequestBodyReaderError(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("POST", &requestURL, errReader(false), "", nil)
	// The request body is only read once the request is sent, so no error
	// should be returned until then.
	require.NoError(t, err, "request body reader error should be ignored")
}

func TestMakeRequestBodyJsonError(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo chan string `json:"foo"`
	}
	_, err := makeRequestBody("POST", &requestURL, x{make(chan string)}, "", nil)
	require.EqualError(t, err, "request marshal failure: json: unsupported type: chan string")
}

type failingUrlEncode string

func (fue failingUrlEncode) EncodeValues(key string, v *url.Values) error {
	return fmt.Errorf("%s", string(fue))
}

func TestMakeRequestBodyQueryFailingEncode(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo failingUrlEncode `url:"foo"`
	}
	_, err := makeRequestBody("GET", &requestURL, x{failingUrlEncode("always failing")}, "", nil)
	require.EqualError(t, err, "cannot create query string: always failing")
}

func TestMakeRequestBodyQueryUnsupported(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("GET", &requestURL, true, "", nil)
	require.EqualError(t, err, "unsupported query string data: true")
}

func TestWithTokenSource(t *testing.T) {
	client := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			foo := r.Header.Get("Foo")
			require.Equal(t, "bar", foo)
			token := r.Header.Get("Authorization")
			reader := strings.NewReader(token)
			return &http.Response{
				StatusCode: 204,
				Request:    r,
				Body:       io.NopCloser(reader),
			}, nil
		}),
	})

	var buf bytes.Buffer
	ctx := context.Background()
	err := client.Do(ctx, "GET", "abc",
		WithResponseUnmarshal(&buf),
		WithRequestHeader("Foo", "bar"),
		WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{
			TokenType:   "awesome",
			AccessToken: "token",
		})))
	require.NoError(t, err)
	require.Equal(t, "awesome token", buf.String())
}

func TestEncodeMultiSegmentPathParameter(t *testing.T) {
	// Slashes should not be encoded.
	assert.Equal(t, "a/b/c", EncodeMultiSegmentPathParameter("a/b/c"))
	// # and ? should be encoded.
	assert.Equal(t, "a%23b%3Fc", EncodeMultiSegmentPathParameter("a#b?c"))
}

func TestMakeRequestBodyExplicitQueryParams(t *testing.T) {
	type x struct {
		Scope string `json:"scope" url:"scope"`
	}
	requestURL := "/a/b/c"
	// For GET, it should be ignored.
	body, err := makeRequestBody("GET", &requestURL, x{"test"}, "", map[string]any{"foo": "bar"})
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))

	requestURL = "/a/b/c"
	body, err = makeRequestBody("POST", &requestURL, x{"test"}, "", map[string]any{"foo": "bar"})
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?foo=bar", requestURL)
	x1 := `{"scope":"test"}`
	require.Equal(t, []byte(x1), bodyBytes)
}
