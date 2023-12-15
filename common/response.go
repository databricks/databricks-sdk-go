package common

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// Represents a response.
//
// Responses must always be closed. For non-streaming responses, they are closed
// during deserialization in the client (see unmarshall()). For streaming
// responses, they are returned to the caller, who is responsible for closing
// them. When debugging is enabled, the value of DebugBytes is printed when
// logging the response body.
type ResponseWrapper struct {
	// ReadCloser is the response body that was received from the server.
	ReadCloser io.ReadCloser

	// DebugBytes is the raw response body that was received from the server, if
	// the response was not streamed.
	DebugBytes []byte

	// Response is the original http.Response object.
	Response *http.Response

	// RequestBody is the request body that was sent to the server.
	RequestBody RequestBody
}

func NewResponseWrapper(data any, response *http.Response, req RequestBody) (ResponseWrapper, error) {
	switch v := data.(type) {
	case io.ReadCloser:
		return ResponseWrapper{
			ReadCloser:  v,
			DebugBytes:  []byte("<Streaming response>"),
			Response:    response,
			RequestBody: req,
		}, nil
	case []byte:
		return ResponseWrapper{
			ReadCloser:  io.NopCloser(bytes.NewReader(v)),
			DebugBytes:  v,
			Response:    response,
			RequestBody: req,
		}, nil
	default:
		return ResponseWrapper{}, errors.New("newResponseBody can only be called with io.ReadCloser or []byte")
	}
}
