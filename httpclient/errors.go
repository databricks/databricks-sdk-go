package httpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
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
func DefaultErrorMapper(ctx context.Context, resp common.ResponseWrapper) error {
	if resp.Response.StatusCode < 400 {
		return nil
	}
	raw, err := io.ReadAll(resp.ReadCloser)
	if err != nil {
		return &HttpError{
			Response: resp.Response,
			Message:  "failed to read response",
			err:      err,
		}
	}
	return &HttpError{
		Response: resp.Response,
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

var urlErrorTransientErrorMessages = []string{
	"connection reset by peer",
	"TLS handshake timeout",
	"connection refused",
	"Unexpected error",
	"i/o timeout",
}

func isRetriableUrlError(err error) bool {
	var urlError *url.Error
	if !errors.As(err, &urlError) {
		return false
	}
	for _, msg := range urlErrorTransientErrorMessages {
		if strings.Contains(err.Error(), msg) {
			return true
		}
	}
	return false
}
