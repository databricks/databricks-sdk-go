package httpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	type structWithContents = struct {
		Contents io.ReadCloser
	}
	want := structWithContents{
		Contents: io.NopCloser(strings.NewReader("foo bar")),
	}

	var got structWithContents
	c := mockClient(make200Response("foo bar"))
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	if gotErr != nil {
		t.Errorf("WithResponseUnmarshal(): want no error, got: %s", gotErr)
	}

	wantBytes, _ := io.ReadAll(want.Contents)
	gotBytes, _ := io.ReadAll(got.Contents)
	if diff := cmp.Diff(wantBytes, gotBytes); diff != "" {
		t.Errorf("WithResponseUnmarshal(): want != got: (-want +got):\n%s", diff)
	}
}

func TestWithResponseUnmarshal_readCloser(t *testing.T) {
	want := io.NopCloser(strings.NewReader("foo bar"))

	var got io.ReadCloser
	c := mockClient(make200Response("foo bar"))
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	if gotErr != nil {
		t.Errorf("WithResponseUnmarshal(): want no error, got: %s", gotErr)
	}

	wantBytes, _ := io.ReadAll(want)
	gotBytes, _ := io.ReadAll(got)
	if diff := cmp.Diff(wantBytes, gotBytes); diff != "" {
		t.Errorf("WithResponseUnmarshal(): want != got: (-want +got):\n%s", diff)
	}
}

func TestWithResponseUnmarshal_byteBuffer(t *testing.T) {
	want := bytes.NewBuffer([]byte("foo bar"))

	var got bytes.Buffer
	c := mockClient(make200Response("foo bar"))
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	if gotErr != nil {
		t.Errorf("WithResponseUnmarshal(): want no error, got: %s", gotErr)
	}
	if diff := cmp.Diff(want.Bytes(), got.Bytes()); diff != "" {
		t.Errorf("WithResponseUnmarshal(): want != got: (-want +got):\n%s", diff)
	}
}

func TestWithResponseUnmarshal_bytes(t *testing.T) {
	want := []byte("foo bar")

	var got []byte
	c := mockClient(make200Response("foo bar"))
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	if gotErr != nil {
		t.Errorf("WithResponseUnmarshal(): want no error, got: %s", gotErr)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("WithResponseUnmarshal(): want != got: (-want +got):\n%s", diff)
	}
}

func TestWithResponseUnmarshal_json(t *testing.T) {
	type jsonStruct struct {
		Foo string `json:"foo"`
	}
	want := jsonStruct{Foo: "bar"}

	var got jsonStruct
	c := mockClient(make200Response(`{"foo": "bar"}`))
	gotErr := c.Do(context.Background(), "GET", "/a", WithResponseUnmarshal(&got))

	if gotErr != nil {
		t.Errorf("WithResponseUnmarshal(): want no error, got: %s", gotErr)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("WithResponseUnmarshal(): want != got: (-want +got):\n%s", diff)
	}
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
