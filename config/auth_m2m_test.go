package config

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"golang.org/x/oauth2"
)

func TestM2mHappyFlow(t *testing.T) {
	assertHeaders(
		t,
		&Config{
			Host:         "a",
			ClientID:     "b",
			ClientSecret: "c",
			AuthType:     "oauth-m2m",
			HTTPTransport: fixtures.MappingTransport{
				"GET /oidc/.well-known/oauth-authorization-server": {
					Response: u2m.OAuthAuthorizationServer{
						AuthorizationEndpoint: "https://localhost:1234/dummy/auth",
						TokenEndpoint:         "https://localhost:1234/dummy/token",
					},
				},
				"POST /dummy/token": {
					ExpectedHeaders: map[string]string{
						"Authorization": "Basic Yjpj",
						"Content-Type":  "application/x-www-form-urlencoded",
					},
					ExpectedRequest: url.Values{
						"grant_type": {"client_credentials"},
						"scope":      {"all-apis"},
					},
					Response: oauth2.Token{
						TokenType:   "Some",
						AccessToken: "cde",
					},
				},
			},
		},
		map[string]string{
			"Authorization": "Some cde",
		},
	)
}

func TestM2mHappyFlowForAccount(t *testing.T) {
	assertHeaders(t,
		&Config{
			Host:         "accounts.cloud.databricks.com",
			AccountID:    "a",
			ClientID:     "b",
			ClientSecret: "c",
			HTTPTransport: fixtures.MappingTransport{
				"POST /oidc/accounts/a/v1/token": {
					ExpectedHeaders: map[string]string{
						"Authorization": "Basic Yjpj",
						"Content-Type":  "application/x-www-form-urlencoded",
					},
					ExpectedRequest: url.Values{
						"grant_type": {"client_credentials"},
						"scope":      {"all-apis"},
					},
					Response: oauth2.Token{
						TokenType:   "Some",
						AccessToken: "cde",
					},
				},
			},
		},
		map[string]string{
			"Authorization": "Some cde",
		},
	)
}

func TestM2mNotSupported(t *testing.T) {
	_, err := authenticateRequest(&Config{
		Host:         "a",
		ClientID:     "b",
		ClientSecret: "c",
		AuthType:     "oauth-m2m",
		HTTPTransport: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Status:   404,
				Response: `...`,
			},
		},
	})
	if !errors.Is(err, u2m.ErrOAuthNotSupported) {
		t.Errorf("got error %v, want %v", err, u2m.ErrOAuthNotSupported)
	}
}

// TestM2mCredentials_DirectTokenSource verifies that M2mCredentials.Configure
// returns a token source that calls the HTTP endpoint on every Token() call
// (no inner ReuseTokenSource caching). With the old code that passed
// ccfg.TokenSource(ctx) (a ReuseTokenSource) only the first of N calls would
// hit the endpoint; the fix passes ccfg.Token(ctx) directly so each call does.
// See https://github.com/databricks/databricks-sdk-go/issues/1549.
func TestM2mCredentials_DirectTokenSource(t *testing.T) {
	var tokenCalls int32
	transport := &tokenCountingTransport{
		calls: &tokenCalls,
		inner: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://localhost/token",
				},
			},
			"POST /token": {
				Response: oauth2.Token{
					TokenType:   "Bearer",
					AccessToken: "test-token",
				},
			},
		},
	}

	cfg := &Config{
		Host:                     "a",
		ClientID:                 "b",
		ClientSecret:             "c",
		AuthType:                 "oauth-m2m",
		ConfigFile:               "/dev/null",
		DisableAsyncTokenRefresh: true,
		HTTPTransport:            transport,
	}

	err := cfg.EnsureResolved()
	if err != nil {
		t.Fatalf("EnsureResolved(): %v", err)
	}

	ctx := cfg.refreshClient.InContextForOAuth2(cfg.refreshCtx)
	provider, err := M2mCredentials{}.Configure(ctx, cfg)
	if err != nil {
		t.Fatalf("Configure(): %v", err)
	}

	oauthProvider := provider.(credentials.OAuthCredentialsProvider)

	// Call Token() 3 times. With the fix, each call hits the endpoint.
	// Without the fix (ReuseTokenSource), only the first call does.
	const nCalls = 3
	for i := 0; i < nCalls; i++ {
		tok, err := oauthProvider.Token(context.Background())
		if err != nil {
			t.Fatalf("Token() call %d: %v", i+1, err)
		}
		if tok.AccessToken == "" {
			t.Fatalf("Token() call %d: empty access token", i+1)
		}
	}

	got := int(atomic.LoadInt32(&tokenCalls))
	if got != nCalls {
		t.Errorf("token endpoint calls = %d, want %d", got, nCalls)
	}
}

type tokenCountingTransport struct {
	calls *int32
	inner http.RoundTripper
}

func (t *tokenCountingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == "POST" {
		atomic.AddInt32(t.calls, 1)
	}
	return t.inner.RoundTrip(req)
}

func (t *tokenCountingTransport) SkipRetryOnIO() bool { return true }

func TestM2M_Scopes(t *testing.T) {
	tests := []struct {
		name   string
		scopes []string
		want   string
	}{
		{
			name:   "nil scopes uses default",
			scopes: nil,
			want:   "all-apis",
		},
		{
			name:   "empty scopes uses default",
			scopes: []string{},
			want:   "all-apis",
		},
		{
			name:   "single scope",
			scopes: []string{"dashboards"},
			want:   "dashboards",
		},
		{
			name:   "multiple scopes are sorted",
			scopes: []string{"jobs", "files:read", "mlflow"},
			want:   "files:read jobs mlflow",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := fixtures.MappingTransport{
				"GET /oidc/.well-known/oauth-authorization-server": {
					Response: u2m.OAuthAuthorizationServer{
						AuthorizationEndpoint: "https://localhost/auth",
						TokenEndpoint:         "https://localhost/token",
					},
				},
				"POST /token": {
					// The scope assertion: verifies the SDK sends the correct scope parameter.
					ExpectedRequest: url.Values{
						"grant_type": {"client_credentials"},
						"scope":      {tt.want},
					},
					Response: oauth2.Token{
						TokenType:   "Bearer",
						AccessToken: "test-token",
					},
				},
			}

			cfg := &Config{
				Host:          "a",
				ClientID:      "b",
				ClientSecret:  "c",
				AuthType:      "oauth-m2m",
				Scopes:        tt.scopes,
				HTTPTransport: transport,
				ConfigFile:    "/dev/null",
			}

			req, err := http.NewRequest("GET", "http://localhost", nil)
			if err != nil {
				t.Fatalf("http.NewRequest(): unexpected error: %v", err)
			}
			err = cfg.Authenticate(req)
			if err != nil {
				t.Fatalf("Authenticate(): unexpected error: %v", err)
			}
			if got, want := req.Header.Get("Authorization"), "Bearer test-token"; got != want {
				t.Errorf("Authorization header: got %q, want %q", got, want)
			}
		})
	}
}
