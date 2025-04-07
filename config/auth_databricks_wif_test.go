package config

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
)

type mockIdTokenProvider struct {
	// input
	audience string
	// output
	idToken string
	err     error
}

func (m *mockIdTokenProvider) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	m.audience = audience
	return &IDToken{
		m.idToken,
	}, m.err
}

func TestDatabricksGithubWIFCredentials(t *testing.T) {
	testCases := []struct {
		desc               string
		cfg                *Config
		idToken            string
		expectedAudience   string
		tokenProviderError error
		wantToken          string
		wantErrPrefix      *string
	}{
		{
			desc: "missing host",
			cfg: &Config{
				ClientID:      "client-id",
				TokenAudience: "token-audience",
			},
			wantErrPrefix: errPrefix("missing Host"),
		},
		{
			desc: "missing client ID",
			cfg: &Config{
				Host:          "http://host.com/test",
				TokenAudience: "token-audience",
			},
			wantErrPrefix: errPrefix("missing ClientID"),
		},
		{
			desc: "auth server error",
			cfg: &Config{
				ClientID: "client-id",
				Host:     "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Status: http.StatusNotFound,
					},
				},
			},
			wantErrPrefix: errPrefix("databricks OAuth is not supported for this host"),
		},
		{
			desc: "token provider error",
			cfg: &Config{
				ClientID:      "client-id",
				Host:          "http://host.com/test",
				TokenAudience: "token-audience",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
				},
			},
			expectedAudience:   "token-audience",
			tokenProviderError: errors.New("error getting id token"),
			wantErrPrefix:      errPrefix("error getting id token"),
		},
		{
			desc: "databricks workspace server error",
			cfg: &Config{
				ClientID:      "client-id",
				Host:          "http://host.com/test",
				TokenAudience: "token-audience",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"POST /oidc/v1/token": {
						Status: http.StatusInternalServerError,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
					},
				},
			},
			expectedAudience: "token-audience",
			idToken:          "id-token-42",
			wantErrPrefix:    errPrefix("oauth2: cannot fetch token: Internal Server Error"),
		},
		{
			desc: "invalid auth token",
			cfg: &Config{
				ClientID:      "client-id",
				Host:          "http://host.com/test",
				TokenAudience: "token-audience",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"POST /oidc/v1/token": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
						Response: map[string]string{
							"foo": "bar",
						},
					},
				},
			},
			expectedAudience: "token-audience",
			idToken:          "id-token-42",
			wantErrPrefix:    errPrefix("oauth2: server response missing access_token"),
		},
		{
			desc: "success workspace",
			cfg: &Config{
				ClientID:      "client-id",
				Host:          "http://host.com/test",
				TokenAudience: "token-audience",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"POST /oidc/v1/token": {

						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
						ExpectedRequest: url.Values{
							"client_id":          {"client-id"},
							"scope":              {"all-apis"},
							"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
							"subject_token":      {"id-token-42"},
							"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
						},
						Response: map[string]string{
							"token_type":    "access-token",
							"access_token":  "test-auth-token",
							"refresh_token": "refresh",
							"expires_on":    "0",
						},
					},
				},
			},
			expectedAudience: "token-audience",
			idToken:          "id-token-42",
			wantToken:        "test-auth-token",
		},
		{
			desc: "success account",
			cfg: &Config{
				ClientID:      "client-id",
				AccountID:     "ac123",
				Host:          "https://accounts.databricks.com",
				TokenAudience: "token-audience",
				HTTPTransport: fixtures.MappingTransport{
					"POST /oidc/accounts/ac123/v1/token": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
						ExpectedRequest: url.Values{
							"client_id":          {"client-id"},
							"scope":              {"all-apis"},
							"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
							"subject_token":      {"id-token-42"},
							"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
						},
						Response: map[string]string{
							"token_type":    "access-token",
							"access_token":  "test-auth-token",
							"refresh_token": "refresh",
							"expires_on":    "0",
						},
					},
				},
			},
			expectedAudience: "token-audience",
			idToken:          "id-token-42",
			wantToken:        "test-auth-token",
		},
		{
			desc: "default token audience account",
			cfg: &Config{
				ClientID:  "client-id",
				AccountID: "ac123",
				Host:      "https://accounts.databricks.com",
				HTTPTransport: fixtures.MappingTransport{
					"POST /oidc/accounts/ac123/v1/token": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
						Response: map[string]string{
							"token_type":    "access-token",
							"access_token":  "test-auth-token",
							"refresh_token": "refresh",
							"expires_on":    "0",
						},
					},
				},
			},
			expectedAudience: "ac123",
			idToken:          "id-token-42",
			wantToken:        "test-auth-token",
		},
		{
			desc: "default token audience workspace",
			cfg: &Config{
				ClientID: "client-id",
				Host:     "https://host.com",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"POST /oidc/v1/token": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
						Response: map[string]string{
							"token_type":    "access-token",
							"access_token":  "test-auth-token",
							"refresh_token": "refresh",
							"expires_on":    "0",
						},
					},
				},
			},
			expectedAudience: "https://host.com/oidc/v1/token",
			idToken:          "id-token-42",
			wantToken:        "test-auth-token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// if tc.desc != "" {
			// 	t.Skip()
			// }
			p := &mockIdTokenProvider{
				idToken: tc.idToken,
				err:     tc.tokenProviderError,
			}
			tc.cfg.EnsureResolved()
			c := tc.cfg.CanonicalHostName()
			ex := &wifTokenExchange{
				clientID:              tc.cfg.ClientID,
				account:               tc.cfg.IsAccountClient(),
				accountID:             tc.cfg.AccountID,
				host:                  c,
				tokenEndpointProvider: tc.cfg.getOidcEndpoints,
				audience:              tc.cfg.TokenAudience,
				tokenProvider:         p,
			}
			ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
				Transport: tc.cfg.HTTPTransport,
			})
			token, gotErr := ex.Token(ctx)
			if tc.wantErrPrefix == nil && gotErr != nil {
				t.Errorf("Token(ctx): got error %q, want none", gotErr)
			}
			if tc.wantErrPrefix != nil && !hasPrefix(gotErr, *tc.wantErrPrefix) {
				t.Errorf("Token(ctx): got error %q, want error with prefix %q", gotErr, *tc.wantErrPrefix)
			}
			if tc.expectedAudience != p.audience {
				t.Errorf("mockTokenProvider: got audience %s, want %s", p.audience, tc.expectedAudience)
			}
			tokenValue := ""
			if token != nil {
				tokenValue = token.AccessToken
			}
			if diff := cmp.Diff(tc.wantToken, tokenValue); diff != "" {
				t.Errorf("Authenticate(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDatabricksWIFCredentials_Name(t *testing.T) {
	strategies := WifTokenCredentialStrategies(&Config{})
	expected := []string{"github-oidc-federated-oidc-github"}
	found := []string{}
	for _, strategy := range strategies {
		found = append(found, strategy.Name())
	}
	if diff := cmp.Diff(expected, found); diff != "" {
		t.Errorf("Strategies mismatch (-want +got):\n%s\n(order must be the same))", diff)
	}
}
