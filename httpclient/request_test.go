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
	body, err := makeRequestBody("GET", &requestURL, x{"test"}, encoding{gRpcEncoding: true})
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))

	requestURL = "/a/b/c"
	body, err = makeRequestBody("POST", &requestURL, x{"test"}, encoding{gRpcEncoding: true})
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c", requestURL)
	x1 := `{"scope":"test"}`
	require.Equal(t, []byte(x1), bodyBytes)

	requestURL = "/a/b/c"
	body, err = makeRequestBody("HEAD", &requestURL, x{"test"}, encoding{gRpcEncoding: true})
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))
}

func TestMakeRequestBodyFromReader(t *testing.T) {
	requestURL := "/a/b/c"
	body, err := makeRequestBody("PUT", &requestURL, strings.NewReader("abc"), encoding{gRpcEncoding: true})
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, []byte("abc"), bodyBytes)
}

func TestUrlEncoding(t *testing.T) {
	data := map[string]any{
		"mapping": map[string]string{
			"key": "value",
		},
	}
	requestURL := "/a/b/c"
	body, err := makeRequestBody("POST", &requestURL, data, encoding{contentType: "application/x-www-form-urlencoded", gRpcEncoding: false})
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c", requestURL)
	// Encoding of mapping={"key":"value"}
	require.Equal(t, "mapping=%7B%22key%22%3A%22value%22%7D", string(bodyBytes))
}

func TestMakeRequestBodyReaderError(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("POST", &requestURL, errReader(false), encoding{gRpcEncoding: true})
	// The request body is only read once the request is sent, so no error
	// should be returned until then.
	require.NoError(t, err, "request body reader error should be ignored")
}

func TestMakeRequestBodyJsonError(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo chan string `json:"foo"`
	}
	_, err := makeRequestBody("POST", &requestURL, x{make(chan string)}, encoding{gRpcEncoding: true})
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
	_, err := makeRequestBody("GET", &requestURL, x{failingUrlEncode("always failing")}, encoding{gRpcEncoding: true})
	require.EqualError(t, err, "cannot create query string: always failing")
}

func TestMakeRequestBodyQueryUnsupported(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("GET", &requestURL, true, encoding{gRpcEncoding: true})
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
