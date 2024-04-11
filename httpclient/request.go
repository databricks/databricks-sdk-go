package httpclient

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

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

func WithAuthVisitor(visitor func(*http.Request) error) DoOption {
	return DoOption{
		in:           visitor,
		isAuthOption: true,
	}
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

// WithUrlEncodedData takes either a struct instance, map, string, bytes, or io.Reader plus
// a content type, and sends it either as query string for GET and DELETE calls, or as request body
// for POST, PUT, and PATCH calls. The content type is set to "application/x-www-form-urlencoded"
// and the body is encoded as a query string.
//
// Experimental: this method may eventually be split into more granular options.
func WithUrlEncodedData(body any) DoOption {
	return DoOption{
		in: func(r *http.Request) error {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			return nil
		},
		body:        body,
		contentType: "application/x-www-form-urlencoded",
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
		return strings.Join(s, "&"), nil
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
		return protoCompatibleParams.Encode(), nil
	}
	return "", fmt.Errorf("unsupported query string data: %#v", data)
}

func EncodeMultiSegmentPathParameter(p string) string {
	segments := strings.Split(p, "/")
	b := strings.Builder{}
	b.Grow(len(p))
	for i, s := range segments {
		if i > 0 {
			b.WriteString("/")
		}
		b.WriteString(url.PathEscape(s))
	}
	return b.String()
}

func makeRequestBody(method string, requestURL *string, data interface{}, contentType string) (common.RequestBody, error) {
	if data == nil {
		return common.RequestBody{}, nil
	}
	if method == "GET" || method == "DELETE" || method == "HEAD" {
		qs, err := makeQueryString(data)
		if err != nil {
			return common.RequestBody{}, err
		}
		*requestURL += "?" + qs
		return common.NewRequestBody([]byte{})
	}
	if contentType == "application/x-www-form-urlencoded" {
		qs, err := makeQueryString(data)
		if err != nil {
			return common.RequestBody{}, err
		}
		return common.NewRequestBody(qs)
	}
	return common.NewRequestBody(data)
}
