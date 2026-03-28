package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

// testHTTPError is a test error that implements the httpStatusCoder interface,
// mirroring httpclient.HttpError without importing it.
type testHTTPError struct {
	code    int
	message string
}

func (e *testHTTPError) Error() string       { return fmt.Sprintf("http %d: %s", e.code, e.message) }
func (e *testHTTPError) HTTPStatusCode() int { return e.code }

func TestRetryingTokenSource(t *testing.T) {
	validToken := &oauth2.Token{
		AccessToken: "test-token",
		Expiry:      time.Now().Add(time.Hour),
	}

	testCases := []struct {
		name       string
		callErrors []error
		wantToken  bool
		wantErr    bool
		wantCalls  int
	}{
		{
			name:       "success on first call",
			callErrors: []error{nil},
			wantToken:  true,
			wantCalls:  1,
		},
		{
			name: "retry on http 429 then succeed",
			callErrors: []error{
				&testHTTPError{code: 429, message: "rate limited"},
				nil,
			},
			wantToken: true,
			wantCalls: 2,
		},
		{
			name: "retry on http 503 then succeed",
			callErrors: []error{
				&testHTTPError{code: 503, message: "service unavailable"},
				nil,
			},
			wantToken: true,
			wantCalls: 2,
		},
		{
			name: "retry on oauth2 retrieve error 429",
			callErrors: []error{
				&oauth2.RetrieveError{Response: &http.Response{StatusCode: 429}},
				nil,
			},
			wantToken: true,
			wantCalls: 2,
		},
		{
			name: "retry on transient network error",
			callErrors: []error{
				&url.Error{Op: "Post", URL: "https://host/token", Err: fmt.Errorf("connection reset by peer")},
				nil,
			},
			wantToken: true,
			wantCalls: 2,
		},
		{
			name: "no retry on http 401",
			callErrors: []error{
				&testHTTPError{code: 401, message: "unauthorized"},
			},
			wantErr:   true,
			wantCalls: 1,
		},
		{
			name: "no retry on http 400",
			callErrors: []error{
				&testHTTPError{code: 400, message: "bad request"},
			},
			wantErr:   true,
			wantCalls: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			callCount := 0
			inner := TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
				err := tc.callErrors[callCount]
				callCount++
				if err != nil {
					return nil, err
				}
				return validToken, nil
			})

			ts := NewRetryingTokenSource(inner)
			tok, err := ts.Token(context.Background())

			if callCount != tc.wantCalls {
				t.Errorf("got %d calls, want %d", callCount, tc.wantCalls)
			}
			if tc.wantErr && err == nil {
				t.Error("got nil error, want error")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("got error %v, want nil", err)
			}
			if tc.wantToken && tok == nil {
				t.Error("got nil token, want token")
			}
			if !tc.wantToken && tok != nil {
				t.Errorf("got token %v, want nil", tok)
			}
		})
	}
}

func TestIsRetriableTokenError(t *testing.T) {
	testCases := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "http 429",
			err:  &testHTTPError{code: 429},
			want: true,
		},
		{
			name: "http 500",
			err:  &testHTTPError{code: 500},
			want: false,
		},
		{
			name: "http 503",
			err:  &testHTTPError{code: 503},
			want: true,
		},
		{
			name: "http 401",
			err:  &testHTTPError{code: 401},
			want: false,
		},
		{
			name: "http 403",
			err:  &testHTTPError{code: 403},
			want: false,
		},
		{
			name: "oauth2 retrieve error 429",
			err:  &oauth2.RetrieveError{Response: &http.Response{StatusCode: 429}},
			want: true,
		},
		{
			name: "oauth2 retrieve error 500",
			err:  &oauth2.RetrieveError{Response: &http.Response{StatusCode: 500}},
			want: false,
		},
		{
			name: "oauth2 retrieve error 400",
			err:  &oauth2.RetrieveError{Response: &http.Response{StatusCode: 400}},
			want: false,
		},
		{
			name: "connection reset",
			err:  &url.Error{Op: "Post", URL: "https://host/token", Err: fmt.Errorf("connection reset by peer")},
			want: true,
		},
		{
			name: "tls handshake timeout",
			err:  &url.Error{Op: "Post", URL: "https://host/token", Err: fmt.Errorf("TLS handshake timeout")},
			want: true,
		},
		{
			name: "connection refused",
			err:  &url.Error{Op: "Post", URL: "https://host/token", Err: fmt.Errorf("connection refused")},
			want: true,
		},
		{
			name: "i/o timeout",
			err:  &url.Error{Op: "Post", URL: "https://host/token", Err: fmt.Errorf("i/o timeout")},
			want: true,
		},
		{
			name: "url error non-transient",
			err:  &url.Error{Op: "Post", URL: "https://host/token", Err: fmt.Errorf("no such host")},
			want: false,
		},
		{
			name: "generic error",
			err:  errors.New("something went wrong"),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isRetriableTokenError(tc.err)
			if got != tc.want {
				t.Errorf("isRetriableTokenError(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
