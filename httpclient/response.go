package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Represents a response body.
//
// Responses must always be closed. For non-streaming responses, they are closed
// during deserialization in the client (see unmarshall()). For streaming
// responses, they are returned to the caller, who is responsible for closing
// them.
type responseBody struct {
	ReadCloser io.ReadCloser
	DebugBytes []byte
	Header     http.Header
	Status     string
	StatusCode int
}

func newResponseBody(data any, header http.Header, statusCode int, status string) (responseBody, error) {
	switch v := data.(type) {
	case io.ReadCloser:
		return responseBody{
			ReadCloser: v,
			DebugBytes: []byte("<io.ReadCloser>"),
			Header:     header,
			StatusCode: statusCode,
			Status:     status,
		}, nil
	case []byte:
		return responseBody{
			ReadCloser: io.NopCloser(bytes.NewReader(v)),
			DebugBytes: v,
			Header:     header,
			StatusCode: statusCode,
			Status:     status,
		}, nil
	default:
		return responseBody{}, errors.New("newResponseBody can only be called with io.ReadCloser or []byte")
	}
}

func WithResponseHeader(key string, value *string) DoOption {
	return DoOption{
		out: func(body *responseBody) error {
			*value = body.Header.Get(key)
			return nil
		},
	}
}

func WithResponseUnmarshal(response any) DoOption {
	return DoOption{
		in: func(r *http.Request) error {
			if r.Header.Get("Accept") != "" {
				return nil
			}
			switch response.(type) {
			case *bytes.Buffer, *io.ReadCloser, *[]byte:
				// don't send Accept header for raw types, even though we have
				// openapi/code/method.go:440#IsResponseByteStream() setting the
				// Accept header explicitly.
				return nil
			default:
				r.Header.Set("Accept", "application/json")
				return nil
			}
		},
		out: func(body *responseBody) error {
			if response == nil {
				return nil
			}
			// If the destination is bytes.Buffer, write the body over there
			if raw, ok := response.(*io.ReadCloser); ok {
				*raw = body.ReadCloser
				return nil
			}
			defer body.ReadCloser.Close()
			bs, err := io.ReadAll(body.ReadCloser)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}
			if len(bs) == 0 {
				return nil
			}
			// If the destination is a byte slice or buffer, pass the body verbatim.
			if raw, ok := response.(*[]byte); ok {
				*raw = bs
				return nil
			}
			if raw, ok := response.(*bytes.Buffer); ok {
				_, err := raw.Write(bs)
				return err
			}
			err = json.Unmarshal(bs, &response)
			if err != nil {
				return fmt.Errorf("failed to unmarshal response body: %w (original: %s)", err, string(bs))
			}
			return nil
		},
	}
}
