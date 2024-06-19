package config

import (
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

func errPrefix(prefix string) *string {
	return &prefix
}

func hasPrefix(err error, prefix string) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), prefix)
}

func TestAzureGithubOIDCCredentials(t *testing.T) {
	testCases := []struct {
		desc         string
		envs         map[string]string
		cfg          *Config
		wantHeaders  map[string]string
		wantErrPefix *string
	}{
		{
			desc: "not an azure config",
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud: "foo-bar-cloud",
				},
			},
			wantErrPefix: errPrefix("default auth: cannot configure default credentials"),
		},
		{
			desc: "missing azure client ID",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_URL":   "http://endpoint.com/test",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN": "token-1337",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				Host:          "http://host.com/test",
				AzureTenantID: "test-tenant-id",
			},
			wantErrPefix: errPrefix("default auth: cannot configure default credentials"),
		},
		{
			desc: "missing host",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_URL":   "http://endpoint.com/test",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN": "token-1337",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID: "test-client-id",
				AzureTenantID: "test-tenant-id",
			},
			wantErrPefix: errPrefix("default auth: cannot configure default credentials"),
		},
		{
			desc: "missing tenant ID",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_URL":   "http://endpoint.com/test",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN": "token-1337",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				Host:          "http://host.com/test",
				AzureClientID: "test-client-id",
			},
			wantErrPefix: errPrefix("default auth: cannot configure default credentials"),
		},
		{
			desc: "missing env ACTIONS_ID_TOKEN_REQUEST_TOKEN",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_URL": "url",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID: "test-client-id",
				AzureTenantID: "test-tenant-id",
				Host:          "http://host.com/test",
			},
			wantErrPefix: errPrefix("default auth: cannot configure default credentials"),
		},
		{
			desc: "missing env ACTIONS_ID_TOKEN_REQUEST_URL",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN": "token",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID: "test-client-id",
				AzureTenantID: "test-tenant-id",
				Host:          "http://host.com/test",
			},
			wantErrPefix: errPrefix("default auth: cannot configure default credentials"),
		},
		{
			desc: "azure aad token exchange server error",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_URL":   "http://endpoint.com/test",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN": "token-1337",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID: "test-client-id",
				AzureTenantID: "test-tenant-id",
				Host:          "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?audience=api%3A%2F%2FAzureADTokenExchange": {
						Status: http.StatusInternalServerError,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
					},
				},
			},
			wantErrPefix: errPrefix("default auth: github-oidc-azure"),
		},
		{
			desc: "azure auth server error",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_URL":   "http://endpoint.com/test",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN": "token-1337",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID: "test-client-id",
				AzureTenantID: "test-tenant-id",
				Host:          "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?audience=api%3A%2F%2FAzureADTokenExchange": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
					"POST /test-tenant-id": {
						Status: http.StatusInternalServerError,
						ExpectedHeaders: map[string]string{
							"Accept":       "application/json",
							"Content-Type": "application/json",
						},
						ExpectedRequest: map[string]any{
							"client_assertion":      "id-token-42",
							"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
							"client_id":             "test-client-id",
							"grant_type":            "client_credentials",
							"resource":              "test-azure-app-id",
						},
					},
				},
			},
			wantErrPefix: errPrefix("inner token: http 500"),
		},
		{
			desc: "invalid auth token",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_URL":   "http://endpoint.com/test",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN": "token-1337",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID: "test-client-id",
				AzureTenantID: "test-tenant-id",
				Host:          "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?audience=api%3A%2F%2FAzureADTokenExchange": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
					"POST /test-tenant-id": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Accept":       "application/json",
							"Content-Type": "application/json",
						},
						ExpectedRequest: map[string]any{
							"client_assertion":      "id-token-42",
							"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
							"client_id":             "test-client-id",
							"grant_type":            "client_credentials",
							"resource":              "test-azure-app-id",
						},
						Response: map[string]string{
							"foo": "bar",
						},
					},
				},
			},
			wantErrPefix: errPrefix("inner token: invalid token"),
		},
		{
			desc: "success",
			envs: map[string]string{
				"ACTIONS_ID_TOKEN_REQUEST_URL":   "http://endpoint.com/test",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN": "token-1337",
			},
			cfg: &Config{
				DatabricksEnvironment: &environment.DatabricksEnvironment{
					Cloud:              environment.CloudAzure,
					AzureApplicationID: "test-azure-app-id",
					AzureEnvironment:   &environment.AzurePublicCloud,
				},
				AzureClientID: "test-client-id",
				AzureTenantID: "test-tenant-id",
				Host:          "http://host.com/test",
				HTTPTransport: fixtures.MappingTransport{
					"GET /test?audience=api%3A%2F%2FAzureADTokenExchange": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Authorization": "Bearer token-1337",
							"Accept":        "application/json",
						},
						Response: `{"value": "id-token-42"}`,
					},
					"POST /test-tenant-id": {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Accept":       "application/json",
							"Content-Type": "application/json",
						},
						ExpectedRequest: map[string]any{
							"client_assertion":      "id-token-42",
							"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
							"client_id":             "test-client-id",
							"grant_type":            "client_credentials",
							"resource":              "test-azure-app-id",
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
			for k, v := range tc.envs {
				t.Setenv(k, v)
			}
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

			if tc.wantErrPefix == nil && gotErr != nil {
				t.Errorf("Authenticate(): got error %q, want none", gotErr)
			}
			if tc.wantErrPefix != nil && !hasPrefix(gotErr, *tc.wantErrPefix) {
				t.Errorf("Authenticate(): got error %q, want error with prefix %q", gotErr, *tc.wantErrPefix)
			}
			if diff := cmp.Diff(tc.wantHeaders, gotHeaders); diff != "" {
				t.Errorf("Authenticate(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAzureGithubOIDCCredentials_Name(t *testing.T) {
	c := AzureGithubOIDCCredentials{}
	want := "github-oidc-azure"

	if got := c.Name(); got != want {
		t.Errorf("Name(): want %s, got %s", want, got)
	}
}
