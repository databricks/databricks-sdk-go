package auth

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"syscall"
	"testing"

	"golang.org/x/oauth2"
)

// timeoutError is a test error that implements net.Error with Timeout() == true.
type timeoutError struct{}

func (e timeoutError) Error() string   { return "i/o timeout" }
func (e timeoutError) Timeout() bool   { return true }
func (e timeoutError) Temporary() bool { return false }

// testHTTPError is a test error that implements the httpStatusCoder interface,
// mirroring httpclient.HttpError without importing it.
type testHTTPError struct {
	code    int
	message string
}

func (e *testHTTPError) Error() string       { return fmt.Sprintf("http %d: %s", e.code, e.message) }
func (e *testHTTPError) HTTPStatusCode() int { return e.code }

func TestRetryingTokenSource(t *testing.T) {
	err401 := &testHTTPError{code: 401, message: "unauthorized"}
	err400 := &testHTTPError{code: 400, message: "bad request"}
	token := &oauth2.Token{AccessToken: "test-token"}

	testCases := []struct {
		name         string
		callErrors   []error
		wantToken    *oauth2.Token
		wantErr      error
		wantNumCalls int
	}{
		{
			name:         "success on first call",
			callErrors:   []error{nil},
			wantToken:    token,
			wantNumCalls: 1,
		},
		{
			name: "retry on http 429 then succeed",
			callErrors: []error{
				&testHTTPError{code: 429, message: "rate limited"},
				nil,
			},
			wantToken:    token,
			wantNumCalls: 2,
		},
		{
			name: "retry on http 503 then succeed",
			callErrors: []error{
				&testHTTPError{code: 503, message: "service unavailable"},
				nil,
			},
			wantToken:    token,
			wantNumCalls: 2,
		},
		{
			name: "retry on oauth2 retrieve error 429",
			callErrors: []error{
				&oauth2.RetrieveError{Response: &http.Response{StatusCode: 429}},
				nil,
			},
			wantToken:    token,
			wantNumCalls: 2,
		},
		{
			name: "retry on transient network error",
			callErrors: []error{
				&url.Error{Op: "Post", URL: "https://host/token", Err: &net.OpError{
					Op: "read", Net: "tcp",
					Err: &os.SyscallError{Syscall: "read", Err: syscall.ECONNRESET},
				}},
				nil,
			},
			wantToken:    token,
			wantNumCalls: 2,
		},
		{
			name: "no retry on http 401",
			callErrors: []error{
				err401,
			},
			wantErr:      err401,
			wantNumCalls: 1,
		},
		{
			name: "no retry on http 400",
			callErrors: []error{
				err400,
			},
			wantErr:      err400,
			wantNumCalls: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotNumCalls := 0
			inner := TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
				err := tc.callErrors[gotNumCalls]
				gotNumCalls++
				if err != nil {
					return nil, err
				}
				return token, nil
			})

			ts := NewRetryingTokenSource(inner)
			gotToken, gotErr := ts.Token(context.Background())

			if gotNumCalls != tc.wantNumCalls {
				t.Errorf("got %d calls, want %d", gotNumCalls, tc.wantNumCalls)
			}
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("got error %v, want %v", gotErr, tc.wantErr)
			}
			if gotToken != tc.wantToken {
				t.Errorf("got token %v, want %v", gotToken, tc.wantToken)
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
			name: "http 502",
			err:  &testHTTPError{code: 502},
			want: true,
		},
		{
			name: "http 503",
			err:  &testHTTPError{code: 503},
			want: true,
		},
		{
			name: "http 504",
			err:  &testHTTPError{code: 504},
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
			err: &url.Error{Op: "Post", URL: "https://host/token", Err: &net.OpError{
				Op: "read", Net: "tcp",
				Err: &os.SyscallError{Syscall: "read", Err: syscall.ECONNRESET},
			}},
			want: true,
		},
		{
			name: "connection refused",
			err: &url.Error{Op: "Post", URL: "https://host/token", Err: &net.OpError{
				Op: "dial", Net: "tcp",
				Err: &os.SyscallError{Syscall: "connect", Err: syscall.ECONNREFUSED},
			}},
			want: true,
		},
		{
			name: "i/o timeout",
			err: &url.Error{Op: "Post", URL: "https://host/token", Err: &net.OpError{
				Op: "read", Net: "tcp",
				Err: timeoutError{},
			}},
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
