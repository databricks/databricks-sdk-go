package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

// Represents a request body.
//
// If the provided request data is an io.Reader, DebugBytes is set to
// "<io.Reader>". Otherwise, DebugBytes is set to the marshaled JSON
// representation of the request data, and ReadCloser is set to a new
// io.ReadCloser that reads from DebugBytes.
//
// Request bodies are never closed by the client, hence only accepting
// io.Reader.
type RequestBody struct {
	Reader      io.Reader
	ContentType string
	DebugBytes  []byte
}

func newRequestBody(data any) (RequestBody, error) {
	switch v := data.(type) {
	case io.Reader:
		return RequestBody{
			Reader:     v,
			DebugBytes: []byte("<io.Reader>"),
		}, nil
	case string:
		return RequestBody{
			Reader:     strings.NewReader(v),
			DebugBytes: []byte(v),
		}, nil
	case []byte:
		return RequestBody{
			Reader:     bytes.NewReader(v),
			DebugBytes: v,
		}, nil
	default:
		bs, err := json.Marshal(data)
		if err != nil {
			return RequestBody{}, fmt.Errorf("request marshal failure: %w", err)
		}
		return RequestBody{
			Reader:      bytes.NewReader(bs),
			ContentType: "application/json",
			DebugBytes:  bs,
		}, nil
	}
}

// Reset a request body to its initial state.
//
// This is used to retry requests with a body that has already been read.
// If the request body is not resettable (i.e. not nil and of type other than
// strings.Reader or bytes.Reader), this will return an error.
func (r RequestBody) reset() error {
	if r.Reader == nil {
		return nil
	}
	if v, ok := r.Reader.(io.Seeker); ok {
		_, err := v.Seek(0, io.SeekStart)
		return err
	} else {
		return fmt.Errorf("cannot reset reader of type %T", r.Reader)
	}
}

// WithRequestHeader adds a request visitor, that sets a header on a request
func WithRequestHeader(k, v string) DoOption {
	return WithRequestVisitor(func(r *http.Request) error {
		r.Header.Set(k, v)
		return nil
	})
}

// WithRequestHeaders adds a request visitor, that set all headers from a map
func WithRequestHeaders(headers map[string]string) DoOption {
	return WithRequestVisitor(func(r *http.Request) error {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
		return nil
	})
}

// WithTokenSource uses the specified golang.org/x/oauth2 token source on a request
func WithTokenSource(ts oauth2.TokenSource) DoOption {
	return WithRequestVisitor(func(r *http.Request) error {
		token, err := ts.Token()
		if err != nil {
			return fmt.Errorf("token: %w", err)
		}
		auth := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
		r.Header.Set("Authorization", auth)
		return nil
	})
}

// WithRequestVisitor applies given function on a request
func WithRequestVisitor(visitor func(r *http.Request) error) DoOption {
	return DoOption{
		in: visitor,
	}
}

// WithRequestData takes either a struct instance, map, string, bytes, or io.Reader
// and sends it either as query string for GET and DELETE calls, or as request body
// for POST, PUT, and PATCH calls.
//
// Experimental: this method may eventually be split into more granular options.
func WithRequestData(body any) DoOption {
	// refactor this, so that we split JSON/query string serialization and make
	// separate request visitors internally.
	return DoOption{
		body: body,
	}
}

func makeQueryString(data interface{}) (string, error) {
	inputVal := reflect.ValueOf(data)
	inputType := reflect.TypeOf(data)
	if inputType.Kind() == reflect.Map {
		s := []string{}
		keys := inputVal.MapKeys()
		// sort map keys by their string repr, so that tests can be deterministic
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})
		for _, k := range keys {
			v := inputVal.MapIndex(k)
			if v.IsZero() {
				continue
			}
			s = append(s, fmt.Sprintf("%s=%s",
				strings.Replace(url.QueryEscape(fmt.Sprintf("%v", k.Interface())), "+", "%20", -1),
				strings.Replace(url.QueryEscape(fmt.Sprintf("%v", v.Interface())), "+", "%20", -1)))
		}
		return "?" + strings.Join(s, "&"), nil
	}
	if inputType.Kind() == reflect.Struct {
		params, err := query.Values(data)
		if err != nil {
			return "", fmt.Errorf("cannot create query string: %w", err)
		}
		// Query parameters may be nested, but the keys generated by
		// query.Values use the "[" and "]" characters to represent nesting.
		// Replace all instances of "[" with "." and "]" with empty string
		// to make the query string compatible with the proto API.
		// See the following for more information:
		// https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule
		protoCompatibleParams := make(url.Values)
		for k, vs := range params {
			newK := strings.Replace(k, "[", ".", -1)
			newK = strings.Replace(newK, "]", "", -1)
			for _, v := range vs {
				protoCompatibleParams.Add(newK, v)
			}
		}
		return "?" + protoCompatibleParams.Encode(), nil
	}
	return "", fmt.Errorf("unsupported query string data: %#v", data)
}

func makeRequestBody(method string, requestURL *string, data interface{}) (common.RequestBody, error) {
	if data == nil {
		return common.RequestBody{}, nil
	}
	if method == "GET" || method == "DELETE" {
		qs, err := makeQueryString(data)
		if err != nil {
			return common.RequestBody{}, err
		}
		*requestURL += qs
		return common.NewRequestBody([]byte{})
	}
	return common.NewRequestBody(data)
}
