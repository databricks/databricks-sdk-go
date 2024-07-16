package httpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func make200Response(body string) *http.Response {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func mockClient(resp *http.Response) *ApiClient {
	return NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			resp.Request = r
			return resp, nil
		}),
	})
}

func TestWithResponseUnmarshal_structWithContent(t *testing.T) {
	c := mockClient(make200Response("foo bar"))
	type structWithContents = struct {
		Contents io.ReadCloser
	}
	want := structWithContents{
		Contents: io.NopCloser(strings.NewReader("foo bar")),
	}

	var got structWithContents
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	require.NoError(t, gotErr)
	wantBytes, _ := io.ReadAll(want.Contents)
	gotBytes, _ := io.ReadAll(got.Contents)
	require.Equal(t, wantBytes, gotBytes)
}

func TestWithResponseUnmarshal_readCloser(t *testing.T) {
	c := mockClient(make200Response("foo bar"))
	want := io.NopCloser(strings.NewReader("foo bar"))

	var got io.ReadCloser
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	require.NoError(t, gotErr)
	wantBytes, _ := io.ReadAll(want)
	gotBytes, _ := io.ReadAll(got)
	require.Equal(t, wantBytes, gotBytes)
}

func TestWithResponseUnmarshal_byteBuffer(t *testing.T) {
	c := mockClient(make200Response("foo bar"))
	want := bytes.NewBuffer([]byte("foo bar"))

	var got bytes.Buffer
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	require.NoError(t, gotErr)
	require.Equal(t, want.Bytes(), got.Bytes())
}

func TestWithResponseUnmarshal_bytes(t *testing.T) {
	c := mockClient(make200Response("foo bar"))
	want := []byte("foo bar")

	var got []byte
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	require.NoError(t, gotErr)
	require.Equal(t, want, got)
}

func TestWithResponseUnmarshal_json(t *testing.T) {
	c := mockClient(make200Response(`{"foo": "bar"}`))
	type jsonStruct struct {
		Foo string `json:"foo"`
	}
	want := jsonStruct{Foo: "bar"}

	var got jsonStruct
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	require.NoError(t, gotErr)
	require.Equal(t, want, got)
}

func TestWithResponseHeader(t *testing.T) {
	client := NewApiClient(ClientConfig{
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

	var out string
	ctx := context.Background()
	err := client.Do(ctx, "GET", "abc",
		WithResponseHeader("Foo", &out))
	require.NoError(t, err)
	require.Equal(t, "some", out)
}
