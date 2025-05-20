package config

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"golang.org/x/oauth2"
)

func TestTokenSourceStrategy_Configure(t *testing.T) {
	testCases := []struct {
		desc       string
		ts         auth.TokenSource
		wantHeader string
		wantError  bool
	}{
		{
			desc: "token source return error",
			ts: auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
				return nil, errors.New("test error")
			}),
			wantError: true,
		},
		{
			desc: "token source return token",
			ts: auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
				return &oauth2.Token{
					AccessToken: "token-123",
				}, nil
			}),
			wantHeader: "Bearer token-123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			start := &tokenSourceStrategy{
				name: "test-strategy",
				ts:   tc.ts,
			}

			provider, err := start.Configure(context.Background(), &Config{})

			if tc.wantError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
				if provider != nil {
					t.Errorf("Expected provider to be nil, but it's not")
				}
				return
			}

			if err != nil {
				t.Errorf("Expected no error, but got %q", err)
			}
			req := &http.Request{Header: make(http.Header)}
			provider.SetHeaders(req)
			if req.Header.Get("Authorization") != tc.wantHeader {
				t.Errorf("Expected header %q, but got %q", tc.wantHeader, req.Header.Get("Authorization"))
			}
		})
	}
}
