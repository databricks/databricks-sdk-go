package httpclient

import (
	"context"
	"errors"
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

// DefaultErrorMapper returns *HttpError
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
	var httpError *HttpError
	if errors.As(err, &httpError) {
		if httpError.StatusCode == 429 {
			return true
		}
		if httpError.StatusCode == 504 {
			return true
		}
	}
	return false
}
