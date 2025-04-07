package config

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

func TestDatabricksGithubWIFCredentials(t *testing.T) {
	testCases := []struct {
		desc          string
		cfg           *Config
		wantHeaders   map[string]string
		wantErrPrefix *string
	}{
		{
			desc: "not an databricks config",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud: "foo-bar-cloud",
				},
			},
			wantErrPrefix: errPrefix("github-oidc auth: not configured"),
		},
		{
			desc: "missing host",
			cfg: &Config{
				ClientID:                   "client-id",
				TokenAudience:              "token-audience",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
			},
			wantErrPrefix: errPrefix("github-oidc auth: not configured"),
		},
		{
			desc: "missing client ID",
			cfg: &Config{
				Host:                       "http://host.com/test",
				TokenAudience:              "token-audience",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
			},
			wantErrPrefix: errPrefix("github-oidc auth: not configured"),
		},
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
			wantErrPrefix: errPrefix("github-oidc auth: not configured"),
		},
		{
			desc: "missing env ACTIONS_ID_TOKEN_REQUEST_URL",
			cfg: &Config{
				ClientID:                   "client-id",
				Host:                       "http://host.com/test",
				TokenAudience:              "token-audience",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
				},
			},
			wantErrPrefix: errPrefix("github-oidc auth: not configured"),
		},
		{
			desc: "databricks token exchange server error",
			cfg: &Config{
				ClientID:                   "client-id",
				Host:                       "http://host.com/test",
				TokenAudience:              "token-audience",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"GET /test?version=1&audience=token-audience": {
						Status: http.StatusInternalServerError,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
					},
				},
			},
			wantErrPrefix: errPrefix("github-oidc"),
		},
		{
			desc: "databricks workspace server error",
			cfg: &Config{
				ClientID:                   "client-id",
				Host:                       "http://host.com/test",
				TokenAudience:              "token-audience",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"GET /test?version=1&audience=token-audience": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
					"POST /oidc/v1/token": {
						Status: http.StatusInternalServerError,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
					},
				},
			},
			wantErrPrefix: errPrefix("inner token: Post \"https://host.com/oidc/v1/token\""),
		},
		{
			desc: "invalid auth token",
			cfg: &Config{
				ClientID:                   "client-id",
				Host:                       "http://host.com/test",
				TokenAudience:              "token-audience",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"GET /test?version=1&audience=token-audience": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
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
			wantErrPrefix: errPrefix("inner token: oauth2: server response missing access_token"),
		},
		{
			desc: "success workspace",
			cfg: &Config{
				ClientID:                   "client-id",
				Host:                       "http://host.com/test",
				TokenAudience:              "token-audience",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"GET /test?version=1&audience=token-audience": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
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
			wantHeaders: map[string]string{
				"Authorization": "access-token test-auth-token",
			},
		},
		{
			desc: "success account",
			cfg: &Config{
				ClientID:                   "client-id",
				AccountID:                  "ac123",
				Host:                       "https://accounts.databricks.com",
				TokenAudience:              "token-audience",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?version=1&audience=token-audience": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
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
			wantHeaders: map[string]string{
				"Authorization": "access-token test-auth-token",
			},
		},
		{
			desc: "default token audience account",
			cfg: &Config{
				ClientID:                   "client-id",
				AccountID:                  "ac123",
				Host:                       "https://accounts.databricks.com",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?version=1&audience=ac123": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
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
			wantHeaders: map[string]string{
				"Authorization": "access-token test-auth-token",
			},
		},
		{
			desc: "default token audience workspace",
			cfg: &Config{
				ClientID:                   "client-id",
				Host:                       "https://host.com",
				ActionsIDTokenRequestURL:   "http://endpoint.com/test?version=1",
				ActionsIDTokenRequestToken: "token-1337",
				HTTPTransport: fixtures.MappingTransport{
					"GET /oidc/.well-known/oauth-authorization-server": {
						Response: u2m.OAuthAuthorizationServer{
							AuthorizationEndpoint: "https://host.com/auth",
							TokenEndpoint:         "https://host.com/oidc/v1/token",
						},
					},
					"GET /test?version=1&audience=https://host.com/oidc/v1/token": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
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
			wantHeaders: map[string]string{
				"Authorization": "access-token test-auth-token",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.cfg.Credentials = &DatabricksWIFCredentials{} // only test this credential strategy
			tc.cfg.DebugHeaders = true
			if tc.wantHeaders == nil {
				tc.wantHeaders = map[string]string{}
			}

			req, _ := http.NewRequest("GET", "http://localhost", nil)
			gotErr := tc.cfg.Authenticate(req)
			gotHeaders := map[string]string{}
			for h := range req.Header {
				gotHeaders[h] = req.Header.Get(h)
			}

			if tc.wantErrPrefix == nil && gotErr != nil {
				t.Errorf("Authenticate(): got error %q, want none", gotErr)
			}
			if tc.wantErrPrefix != nil && !hasPrefix(gotErr, *tc.wantErrPrefix) {
				t.Errorf("Authenticate(): got error %q, want error with prefix %q", gotErr, *tc.wantErrPrefix)
			}
			if diff := cmp.Diff(tc.wantHeaders, gotHeaders); diff != "" {
				t.Errorf("Authenticate(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDatabricksWIFCredentials_Name(t *testing.T) {
	c := DatabricksWIFCredentials{}
	want := "github-oidc"

	if got := c.Name(); got != want {
		t.Errorf("Name(): got %s, want %s", got, want)
	}
}
