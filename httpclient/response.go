package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common"
)

func WithResponseHeader(key string, value *string) DoOption {
	return DoOption{
		out: func(body *common.ResponseWrapper) error {
			*value = body.Response.Header.Get(key)
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
		out: func(body *common.ResponseWrapper) error {
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
				return apierr.MakeUnexpectedError(body.Response, err, body.RequestBody.DebugBytes, bs)
			}
			return nil
		},
	}
}
