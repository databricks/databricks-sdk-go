package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	*http.Response
	Message string
	err     error
}

func (r *HttpError) Unwrap() error {
	return r.err
}

func (r *HttpError) Error() string {
	return fmt.Sprintf("http %d: %s", r.StatusCode, r.Message)
}

func DefaultErrorMapper(ctx context.Context, resp *http.Response, body io.ReadCloser) error {
	if resp.StatusCode < 400 {
		return nil
	}
	raw, err := io.ReadAll(body)
	if err != nil {
		return &HttpError{
			Response: resp,
			Message:  "failed to read response",
			err:      err,
		}
	}
	return &HttpError{
		Response: resp,
		Message:  string(raw),
	}
}

func DefaultErrorRetriable(ctx context.Context, err error) bool {
	switch some := err.(type) {
	case *HttpError:
		if some.StatusCode == 429 {
			return true
		}
		if some.StatusCode == 504 {
			return true
		}
	}
	return false
}
