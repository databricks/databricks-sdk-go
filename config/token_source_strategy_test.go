package config

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
)

type staticTokenSource struct {
	token *oauth2.Token
	err   error
}

func (s *staticTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	return s.token, s.err
}

func TestDatabricksTokenSourceStrategy(t *testing.T) {
	testCases := []struct {
		desc             string
		token            *oauth2.Token
		tokenSourceError error
		wantHeaders      http.Header
	}{
		{
			desc:             "token source error skips",
			tokenSourceError: errors.New("random error"),
		},
		{
			desc: "token source error skips",
			token: &oauth2.Token{
				AccessToken: "token-123",
			},
			wantHeaders: http.Header{"Authorization": {"Bearer token-123"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ts := &staticTokenSource{
				token: tc.token,
				err:   tc.tokenSourceError,
			}
			strat := &tokenSourceStrategy{
				name:        "github-oidc",
				tokenSource: ts,
			}
			provider, err := strat.Configure(context.Background(), &Config{})
			if tc.tokenSourceError == nil && provider == nil {
				t.Error("Provider expected to not be nil, but it is")
			}
			if tc.tokenSourceError != nil && provider != nil {
				t.Error("A failure in the TokenSource should cause the provider to be nil, but it's not")
			}
			if err != nil {
				t.Errorf("Configure() got error %q, want none", err)
			}

			if provider != nil {
				req, _ := http.NewRequest("GET", "http://localhost", nil)

				gotErr := provider.SetHeaders(req)

				if gotErr != nil {
					t.Errorf("SetHeaders(): got error %q, want none", gotErr)
				}
				if diff := cmp.Diff(tc.wantHeaders, req.Header); diff != "" {
					t.Errorf("Authenticate(): mismatch (-want +got):\n%s", diff)
				}

			}

		})
	}
}
