package config

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

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
