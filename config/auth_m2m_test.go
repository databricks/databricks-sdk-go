package config

import (
	"errors"
	"net/http"
	"net/url"
	"sort"
	"strings"
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
		name              string
		host              string
		accountID         string
		scopes            []string
		wellKnownEndpoint string
		authServer        *u2m.OAuthAuthorizationServer
		tokenEndpoint     string
		expectedToken     string
	}{
		{
			name:              "default scopes when not configured",
			host:              "a",
			scopes:            nil,
			wellKnownEndpoint: "GET /oidc/.well-known/oauth-authorization-server",
			authServer: &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://localhost:1234/dummy/auth",
				TokenEndpoint:         "https://localhost:1234/dummy/token",
			},
			tokenEndpoint: "POST /dummy/token",
			expectedToken: "test-token",
		},
		{
			name:              "single scope",
			host:              "a",
			scopes:            []string{"dashboards"},
			wellKnownEndpoint: "GET /oidc/.well-known/oauth-authorization-server",
			authServer: &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://localhost:1234/dummy/auth",
				TokenEndpoint:         "https://localhost:1234/dummy/token",
			},
			tokenEndpoint: "POST /dummy/token",
			expectedToken: "test-token",
		},
		{
			name:              "multiple scopes sorted",
			host:              "a",
			scopes:            []string{"jobs", "files", "mlflow"},
			wellKnownEndpoint: "GET /oidc/.well-known/oauth-authorization-server",
			authServer: &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://localhost:1234/dummy/auth",
				TokenEndpoint:         "https://localhost:1234/dummy/token",
			},
			tokenEndpoint: "POST /dummy/token",
			expectedToken: "test-token",
		},
		{
			name:              "empty scopes uses default",
			host:              "a",
			scopes:            []string{},
			wellKnownEndpoint: "GET /oidc/.well-known/oauth-authorization-server",
			authServer: &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://localhost:1234/dummy/auth",
				TokenEndpoint:         "https://localhost:1234/dummy/token",
			},
			tokenEndpoint: "POST /dummy/token",
			expectedToken: "test-token",
		},
		{
			name:              "workspace host",
			host:              "https://my-workspace.cloud.databricks.com",
			scopes:            []string{"mlflow:read"},
			wellKnownEndpoint: "GET /oidc/.well-known/oauth-authorization-server",
			authServer: &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://my-workspace.cloud.databricks.com/oidc/v1/authorize",
				TokenEndpoint:         "https://my-workspace.cloud.databricks.com/oidc/v1/token",
			},
			tokenEndpoint: "POST /oidc/v1/token",
			expectedToken: "workspace-token",
		},
		{
			name:          "account host",
			host:          "accounts.cloud.databricks.com",
			accountID:     "my-account",
			scopes:        []string{"iam", "jobs", "files"},
			tokenEndpoint: "POST /oidc/accounts/my-account/v1/token",
			expectedToken: "account-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare expected scope to match SDK's normalization (sorted, space-separated).
			expectedScope := "all-apis"
			if len(tt.scopes) > 0 {
				sortedScopes := make([]string, len(tt.scopes))
				copy(sortedScopes, tt.scopes)
				sort.Strings(sortedScopes)
				expectedScope = strings.Join(sortedScopes, " ")
			}

			transport := fixtures.MappingTransport{}

			if tt.wellKnownEndpoint != "" {
				transport[tt.wellKnownEndpoint] = fixtures.HTTPFixture{
					Response: tt.authServer,
				}
			}

			transport[tt.tokenEndpoint] = fixtures.HTTPFixture{
				ExpectedHeaders: map[string]string{
					"Authorization": "Basic Yjpj",
					"Content-Type":  "application/x-www-form-urlencoded",
				},
				ExpectedRequest: url.Values{
					"grant_type": {"client_credentials"},
					"scope":      {expectedScope},
				},
				Response: oauth2.Token{
					TokenType:   "Bearer",
					AccessToken: tt.expectedToken,
				},
			}

			cfg := &Config{
				Host:          tt.host,
				ClientID:      "b",
				ClientSecret:  "c",
				AuthType:      "oauth-m2m",
				HTTPTransport: transport,
			}

			if tt.accountID != "" {
				cfg.AccountID = tt.accountID
			}

			if tt.scopes != nil {
				cfg.Scopes = tt.scopes
			}

			cfg.ConfigFile = "/dev/null"
			req, _ := http.NewRequest("GET", "http://localhost", nil)
			err := cfg.Authenticate(req)
			if err != nil {
				t.Fatalf("Authenticate(): unexpected error: %v", err)
			}
			if got, want := req.Header.Get("Authorization"), "Bearer "+tt.expectedToken; got != want {
				t.Errorf("Authorization header: got %q, want %q", got, want)
			}
		})
	}
}
