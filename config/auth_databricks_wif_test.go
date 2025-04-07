package config

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
)

type mockIdTokenProvider struct {
	idToken *IDToken
	err     error
}

func (m *mockIdTokenProvider) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	return m.idToken, m.err
}

func TestDatabricksGithubWIFCredentials(t *testing.T) {
	testCases := []struct {
		desc               string
		cfg                *Config
		idToken            string
		tokenProviderError error
		wantToken          string
		wantErrPrefix      *string
	}{
		// {
		// 	desc: "missing host",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		TokenAudience:              "token-audience",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 	},
		// 	wantErrPrefix: errPrefix("missing Host"),
		// },
		// {
		// 	desc: "missing client ID",
		// 	cfg: &Config{
		// 		Host:                       "http://host.com/test",
		// 		TokenAudience:              "token-audience",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 	},
		// 	wantErrPrefix: errPrefix("missing ClientID"),
		// },
		{
			desc: "missing env ACTIONS_ID_TOKEN_REQUEST_TOKEN",
			cfg: &Config{
				ClientID:                 "client-id",
				Host:                     "http://host.com/test",
				TokenAudience:            "token-audience",
				ActionsIDTokenRequestURL: "http://endpoint.com/test?version=1",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
				},
			},
			wantErrPrefix: errPrefix("missing ActionsIdTokenRequestToken"),
		},
		// {
		// 	desc: "missing env ACTIONS_ID_TOKEN_REQUEST_URL",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		Host:                       "http://host.com/test",
		// 		TokenAudience:              "token-audience",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 		HTTPTransport: fixtures.MappingTransport{
		// 			"GET /oidc/.well-known/oauth-authorization-server": {
		// 				Response: u2m.OAuthAuthorizationServer{
		// 					AuthorizationEndpoint: "https://host.com/auth",
		// 					TokenEndpoint:         "https://host.com/oidc/v1/token",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantErrPrefix: errPrefix("federated-oidc-github auth: not configured"),
		// },
		// {
		// 	desc: "databricks token exchange server error",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		Host:                       "http://host.com/test",
		// 		TokenAudience:              "token-audience",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 		HTTPTransport: fixtures.MappingTransport{
		// 			"GET /oidc/.well-known/oauth-authorization-server": {
		// 				Response: u2m.OAuthAuthorizationServer{
		// 					AuthorizationEndpoint: "https://host.com/auth",
		// 					TokenEndpoint:         "https://host.com/oidc/v1/token",
		// 				},
		// 			},
		// 			"GET /test?version=1&audience=token-audience": {
		// 				Status: http.StatusInternalServerError,
		// 				ExpectedHeaders: map[string]string{
		// 					"Authorization": "Bearer token-1337",
		// 					"Accept":        "application/json",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantErrPrefix: errPrefix("federated-oidc-github"),
		// },
		// {
		// 	desc: "databricks workspace server error",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		Host:                       "http://host.com/test",
		// 		TokenAudience:              "token-audience",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 		HTTPTransport: fixtures.MappingTransport{
		// 			"GET /oidc/.well-known/oauth-authorization-server": {
		// 				Response: u2m.OAuthAuthorizationServer{
		// 					AuthorizationEndpoint: "https://host.com/auth",
		// 					TokenEndpoint:         "https://host.com/oidc/v1/token",
		// 				},
		// 			},
		// 			"GET /test?version=1&audience=token-audience": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Authorization": "Bearer token-1337",
		// 					"Accept":        "application/json",
		// 				},
		// 				Response: `{"value": "id-token-42"}`,
		// 			},
		// 			"POST /oidc/v1/token": {
		// 				Status: http.StatusInternalServerError,
		// 				ExpectedHeaders: map[string]string{
		// 					"Content-Type": "application/x-www-form-urlencoded",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantErrPrefix: errPrefix("federated-oidc-github auth: not configured"),
		// },
		// {
		// 	desc: "invalid auth token",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		Host:                       "http://host.com/test",
		// 		TokenAudience:              "token-audience",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 		HTTPTransport: fixtures.MappingTransport{
		// 			"GET /oidc/.well-known/oauth-authorization-server": {
		// 				Response: u2m.OAuthAuthorizationServer{
		// 					AuthorizationEndpoint: "https://host.com/auth",
		// 					TokenEndpoint:         "https://host.com/oidc/v1/token",
		// 				},
		// 			},
		// 			"GET /test?version=1&audience=token-audience": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Authorization": "Bearer token-1337",
		// 					"Accept":        "application/json",
		// 				},
		// 				Response: `{"value": "id-token-42"}`,
		// 			},
		// 			"POST /oidc/v1/token": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Content-Type": "application/x-www-form-urlencoded",
		// 				},
		// 				Response: map[string]string{
		// 					"foo": "bar",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantErrPrefix: errPrefix("federated-oidc-github auth: not configured"),
		// },
		// {
		// 	desc: "success workspace",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		Host:                       "http://host.com/test",
		// 		TokenAudience:              "token-audience",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 		HTTPTransport: fixtures.MappingTransport{
		// 			"GET /oidc/.well-known/oauth-authorization-server": {
		// 				Response: u2m.OAuthAuthorizationServer{
		// 					AuthorizationEndpoint: "https://host.com/auth",
		// 					TokenEndpoint:         "https://host.com/oidc/v1/token",
		// 				},
		// 			},
		// 			"GET /test?version=1&audience=token-audience": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Authorization": "Bearer token-1337",
		// 					"Accept":        "application/json",
		// 				},
		// 				Response: `{"value": "id-token-42"}`,
		// 			},
		// 			"POST /oidc/v1/token": {

		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Content-Type": "application/x-www-form-urlencoded",
		// 				},
		// 				Response: map[string]string{
		// 					"token_type":    "access-token",
		// 					"access_token":  "test-auth-token",
		// 					"refresh_token": "refresh",
		// 					"expires_on":    "0",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantHeaders: map[string]string{
		// 		"Authorization": "access-token test-auth-token",
		// 	},
		// },
		// {
		// 	desc: "success account",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		AccountID:                  "ac123",
		// 		Host:                       "https://accounts.databricks.com",
		// 		TokenAudience:              "token-audience",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 		HTTPTransport: fixtures.MappingTransport{
		// 			"GET /test?version=1&audience=token-audience": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Authorization": "Bearer token-1337",
		// 					"Accept":        "application/json",
		// 				},
		// 				Response: `{"value": "id-token-42"}`,
		// 			},
		// 			"POST /oidc/accounts/ac123/v1/token": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Content-Type": "application/x-www-form-urlencoded",
		// 				},
		// 				Response: map[string]string{
		// 					"token_type":    "access-token",
		// 					"access_token":  "test-auth-token",
		// 					"refresh_token": "refresh",
		// 					"expires_on":    "0",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantHeaders: map[string]string{
		// 		"Authorization": "access-token test-auth-token",
		// 	},
		// },
		// {
		// 	desc: "default token audience account",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		AccountID:                  "ac123",
		// 		Host:                       "https://accounts.databricks.com",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 		HTTPTransport: fixtures.MappingTransport{
		// 			"GET /test?version=1&audience=ac123": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Authorization": "Bearer token-1337",
		// 					"Accept":        "application/json",
		// 				},
		// 				Response: `{"value": "id-token-42"}`,
		// 			},
		// 			"POST /oidc/accounts/ac123/v1/token": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Content-Type": "application/x-www-form-urlencoded",
		// 				},
		// 				Response: map[string]string{
		// 					"token_type":    "access-token",
		// 					"access_token":  "test-auth-token",
		// 					"refresh_token": "refresh",
		// 					"expires_on":    "0",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantHeaders: map[string]string{
		// 		"Authorization": "access-token test-auth-token",
		// 	},
		// },
		// {
		// 	desc: "default token audience workspace",
		// 	cfg: &Config{
		// 		ClientID:                   "client-id",
		// 		Host:                       "https://host.com",
		// 		ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
		// 		ActionsIDTokenRequestToken: "token-1337",
		// 		HTTPTransport: fixtures.MappingTransport{
		// 			"GET /oidc/.well-known/oauth-authorization-server": {
		// 				Response: u2m.OAuthAuthorizationServer{
		// 					AuthorizationEndpoint: "https://host.com/auth",
		// 					TokenEndpoint:         "https://host.com/oidc/v1/token",
		// 				},
		// 			},
		// 			"GET /test?version=1&audience=https://host.com/oidc/v1/token": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Authorization": "Bearer token-1337",
		// 					"Accept":        "application/json",
		// 				},
		// 				Response: `{"value": "id-token-42"}`,
		// 			},
		// 			"POST /oidc/v1/token": {
		// 				Status: http.StatusOK,
		// 				ExpectedHeaders: map[string]string{
		// 					"Content-Type": "application/x-www-form-urlencoded",
		// 				},
		// 				Response: map[string]string{
		// 					"token_type":    "access-token",
		// 					"access_token":  "test-auth-token",
		// 					"refresh_token": "refresh",
		// 					"expires_on":    "0",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	wantHeaders: map[string]string{
		// 		"Authorization": "access-token test-auth-token",
		// 	},
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// if tc.desc != "" {
			// 	t.Skip()
			// }
			p := &mockIdTokenProvider{
				idToken: &IDToken{
					Value: tc.idToken,
				},
				err: tc.tokenProviderError,
			}
			resolveErr := tc.cfg.EnsureResolved()
			if resolveErr != nil {
				t.Errorf("EnsureResolved(): got error %q, want none", resolveErr)
			}
			ex := &wifTokenExchange{
				clientID:              tc.cfg.ClientID,
				account:               tc.cfg.IsAccountClient(),
				accountID:             tc.cfg.AccountID,
				host:                  tc.cfg.Host,
				tokenEndpointProvider: tc.cfg.getOidcEndpoints,
				audience:              tc.cfg.TokenAudience,
				tokenProvider:         p,
			}
			ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tc.cfg.HTTPTransport)
			token, gotErr := ex.Token(ctx)
			if tc.wantErrPrefix == nil && gotErr != nil {
				t.Errorf("Token(ctx): got error %q, want none", gotErr)
			}
			if tc.wantErrPrefix != nil && !hasPrefix(gotErr, *tc.wantErrPrefix) {
				t.Errorf("Token(ctx): got error %q, want error with prefix %q", gotErr, *tc.wantErrPrefix)
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
