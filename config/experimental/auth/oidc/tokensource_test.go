package oidc

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
)

func errPrefix(s string) *string {
	return &s
}

func hasPrefix(err error, prefix string) bool {
	return strings.HasPrefix(err.Error(), prefix)
}

func TestDatabricksOidcTokenSource(t *testing.T) {
	testCases := []struct {
		desc                 string
		clientID             string
		accountID            string
		host                 string
		tokenAudience        string
		httpTransport        http.RoundTripper
		oidcEndpointProvider func(context.Context) (*u2m.OAuthAuthorizationServer, error)
		idToken              string
		wantAudience         string
		tokenProviderError   error
		wantToken            string
		wantErrPrefix        *string
	}{
		{
			desc:          "missing host",
			clientID:      "client-id",
			tokenAudience: "token-audience",
			wantErrPrefix: errPrefix("missing Host"),
		},
		{
			desc: "token provider error",

			clientID:      "client-id",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			wantAudience:       "token-audience",
			tokenProviderError: errors.New("error getting id token"),
			wantErrPrefix:      errPrefix("error getting id token"),
		},
		{
			desc:          "databricks workspace server error",
			clientID:      "client-id",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {
					Status: http.StatusInternalServerError,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
				},
			},
			wantAudience:  "token-audience",
			idToken:       "id-token-42",
			wantErrPrefix: errPrefix("oauth2: cannot fetch token: Internal Server Error"),
		},
		{
			desc:          "invalid auth token",
			clientID:      "client-id",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
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
			wantAudience:  "token-audience",
			idToken:       "id-token-42",
			wantErrPrefix: errPrefix("oauth2: server response missing access_token"),
		},
		{
			desc:          "success WIF workspace",
			clientID:      "client-id",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
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
			wantAudience: "token-audience",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
		{
			desc:          "success WIF account",
			clientID:      "client-id",
			accountID:     "ac123",
			host:          "https://accounts.databricks.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
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
			wantAudience: "token-audience",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
		{
			desc:      "default token audience account",
			clientID:  "client-id",
			accountID: "ac123",
			host:      "https://accounts.databricks.com",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
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
			wantAudience: "ac123",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
		{
			desc:     "default token audience workspace",
			clientID: "client-id",
			host:     "https://host.com",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
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
			wantAudience: "https://host.com/oidc/v1/token",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
		{
			desc:          "success account-wide",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {

					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
					ExpectedRequest: url.Values{
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
			wantAudience: "token-audience",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			var gotAudience string // set when IDTokenSource is called
			cfg := DatabricksOIDCTokenSourceConfig{
				ClientID:              tc.clientID,
				AccountID:             tc.accountID,
				Host:                  tc.host,
				Scopes:                []string{"all-apis"},
				TokenEndpointProvider: tc.oidcEndpointProvider,
				Audience:              tc.tokenAudience,
				IDTokenSource: IDTokenSourceFn(func(ctx context.Context, aud string) (*IDToken, error) {
					gotAudience = aud
					return &IDToken{Value: tc.idToken}, tc.tokenProviderError
				}),
			}

			ts := NewDatabricksOIDCTokenSource(cfg)
			if tc.httpTransport != nil {
				ts.(*databricksOIDCTokenSource).cfg.TokenEndpointProvider = func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
					return &u2m.OAuthAuthorizationServer{
						TokenEndpoint: "https://host.com/oidc/v1/token",
					}, nil
				}
			}

			ctx := context.Background()
			if tc.httpTransport != nil {
				ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
					Transport: tc.httpTransport,
				})
			}

			token, err := ts.Token(ctx)
			if tc.wantErrPrefix == nil && err != nil {
				t.Errorf("Token(ctx): got error %q, want none", err)
			}
			if tc.wantErrPrefix != nil && !hasPrefix(err, *tc.wantErrPrefix) {
				t.Errorf("Token(ctx): got error %q, want error with prefix %q", err, *tc.wantErrPrefix)
			}
			if tc.wantAudience != gotAudience {
				t.Errorf("mockTokenProvider: got audience %s, want %s", gotAudience, tc.wantAudience)
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

func TestWIF_Scopes(t *testing.T) {
	tests := []struct {
		name                string
		clientID            string
		accountID           string
		host                string
		audience            string
		scopes              []string
		tokenEndpoint       string
		expectedClientID    string
		expectedScope       string
		expectedAccessToken string
	}{
		{
			name:                "single scope",
			clientID:            "client-id",
			host:                "http://host.com",
			audience:            "token-audience",
			scopes:              []string{"dashboards"},
			tokenEndpoint:       "https://host.com/oidc/v1/token",
			expectedClientID:    "client-id",
			expectedScope:       "dashboards",
			expectedAccessToken: "test-token",
		},
		{
			name:                "multiple scopes sorted",
			clientID:            "client-id",
			host:                "http://host.com",
			audience:            "token-audience",
			scopes:              []string{"files", "jobs", "mlflow"},
			tokenEndpoint:       "https://host.com/oidc/v1/token",
			expectedClientID:    "client-id",
			expectedScope:       "files jobs mlflow",
			expectedAccessToken: "test-token",
		},
		{
			name:                "workspace-level WIF",
			clientID:            "client-id",
			host:                "https://my-workspace.cloud.databricks.com",
			audience:            "workspace-audience",
			scopes:              []string{"genie"},
			tokenEndpoint:       "https://my-workspace.cloud.databricks.com/oidc/v1/token",
			expectedClientID:    "client-id",
			expectedScope:       "genie",
			expectedAccessToken: "workspace-token",
		},
		{
			name:                "account-level WIF",
			clientID:            "client-id",
			accountID:           "my-account",
			host:                "https://accounts.cloud.databricks.com",
			audience:            "account-audience",
			scopes:              []string{"files", "iam"},
			tokenEndpoint:       "https://accounts.cloud.databricks.com/oidc/accounts/my-account/v1/token",
			expectedClientID:    "client-id",
			expectedScope:       "files iam",
			expectedAccessToken: "account-token",
		},
		{
			name:                "account-wide token federation (no ClientID)",
			clientID:            "",
			host:                "http://host.com",
			audience:            "token-audience",
			scopes:              []string{"workspaces"},
			tokenEndpoint:       "https://host.com/oidc/v1/token",
			expectedClientID:    "",
			expectedScope:       "workspaces",
			expectedAccessToken: "account-wide-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := DatabricksOIDCTokenSourceConfig{
				ClientID:  tt.clientID,
				AccountID: tt.accountID,
				Host:      tt.host,
				TokenEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
					return &u2m.OAuthAuthorizationServer{
						TokenEndpoint: tt.tokenEndpoint,
					}, nil
				},
				Audience: tt.audience,
				IDTokenSource: IDTokenSourceFn(func(ctx context.Context, aud string) (*IDToken, error) {
					return &IDToken{Value: "id-token"}, nil
				}),
				Scopes: tt.scopes,
			}

			ts := NewDatabricksOIDCTokenSource(cfg)

			expectedRequest := url.Values{
				"scope":              {tt.expectedScope},
				"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
				"subject_token":      {"id-token"},
				"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
			}
			if tt.expectedClientID != "" {
				expectedRequest["client_id"] = []string{tt.expectedClientID}
			}

			endpointURL, _ := url.Parse(tt.tokenEndpoint)
			endpointPath := "POST " + endpointURL.Path

			ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
				Transport: fixtures.MappingTransport{
					endpointPath: {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
						ExpectedRequest: expectedRequest,
						Response: map[string]string{
							"token_type":   "Bearer",
							"access_token": tt.expectedAccessToken,
						},
					},
				},
			})

			token, err := ts.Token(ctx)
			if err != nil {
				t.Fatalf("Token(ctx): got error %q, want none", err)
			}
			if token.AccessToken != tt.expectedAccessToken {
				t.Errorf("Token(ctx): got access token %q, want %q", token.AccessToken, tt.expectedAccessToken)
			}
		})
	}
}
