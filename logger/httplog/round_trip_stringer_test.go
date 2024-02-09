package httplog

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoHeadersNoBody(t *testing.T) {
	res := RoundTripStringer{
		Response: &http.Response{
			Request: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Path:     "/",
					RawQuery: "foo=bar&baz=qux",
				},
			},
			Status: "200 OK",
			Proto:  "HTTP/1.1",
		},
	}.String()
	assert.Equal(t, `GET /?foo=bar&baz=qux
< HTTP/1.1 200 OK`, res)
}

func TestRequestAndResponseHaveHeadersAndBody(t *testing.T) {
	res := RoundTripStringer{
		Response: &http.Response{
			Request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/"},
				Header: http.Header{
					"Foo": []string{"bar"},
					"Bar": []string{"baz"},
				},
			},
			Status: "200 OK",
			Proto:  "HTTP/1.1",
			Header: http.Header{
				"Response-Foo": []string{"bar"},
				"Response-Bar": []string{"baz"},
			},
		},
		RequestBody:        []byte("request-hello"),
		ResponseBody:       []byte("response-hello"),
		DebugHeaders:       true,
		DebugTruncateBytes: 100,
	}.String()
	assert.Equal(t, `GET /
> * Host: 
> * Bar: baz
> * Foo: bar
> request-hello
< HTTP/1.1 200 OK
< * Response-Bar: baz
< * Response-Foo: bar
< response-hello`, res)
}

func TestDoNotPrintHeadersWhenNotConfigured(t *testing.T) {
	res := RoundTripStringer{
		Response: &http.Response{
			Request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/"},
				Header: http.Header{
					"Foo": []string{"bar"},
					"Bar": []string{"baz"},
				},
			},
			Status: "200 OK",
			Proto:  "HTTP/1.1",
			Header: http.Header{
				"Response-Foo": []string{"bar"},
				"Response-Bar": []string{"baz"},
			},
		},
		RequestBody:        []byte("request-hello"),
		ResponseBody:       []byte("response-hello"),
		DebugHeaders:       false,
		DebugTruncateBytes: 100,
	}.String()
	assert.Equal(t, `GET /
> request-hello
< HTTP/1.1 200 OK
< response-hello`, res)
}

func TestHideAuthorizationHeaderWhenConfigured(t *testing.T) {
	res := RoundTripStringer{
		Response: &http.Response{
			Request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/"},
				Header: http.Header{
					"Foo":           []string{"bar"},
					"Authorization": []string{"baz"},
				},
			},
			Status: "200 OK",
			Proto:  "HTTP/1.1",
		},
		RequestBody:           []byte("request-hello"),
		ResponseBody:          []byte("response-hello"),
		DebugHeaders:          true,
		DebugTruncateBytes:    100,
		DebugSensitiveHeaders: false,
	}.String()
	assert.Equal(t, `GET /
> * Host: 
> * Authorization: REDACTED
> * Foo: bar
> request-hello
< HTTP/1.1 200 OK
< response-hello`, res)
}
