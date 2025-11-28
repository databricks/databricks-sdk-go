package config

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHostType_AwsAccount(t *testing.T) {
	c := &Config{
		Host:      "https://accounts.cloud.databricks.com",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.Equal(t, AccountHost, c.HostType())
}

func TestHostType_AwsDodAccount(t *testing.T) {
	c := &Config{
		Host:      "https://accounts-dod.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.Equal(t, AccountHost, c.HostType())
}

func TestHostType_AwsWorkspace(t *testing.T) {
	c := &Config{
		Host:      "https://my-workspace.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.Equal(t, WorkspaceHost, c.HostType())
}

func TestHostType_Unified(t *testing.T) {
	c := &Config{
		Host:                       "https://unified.cloud.databricks.com",
		AccountID:                  "123e4567-e89b-12d3-a456-426614174000",
		Experimental_IsUnifiedHost: true,
	}
	assert.Equal(t, UnifiedHost, c.HostType())
}

func TestIsAccountClient_PanicsOnUnifiedHost(t *testing.T) {
	c := &Config{
		Host:                       "https://unified.cloud.databricks.com",
		AccountID:                  "test-account",
		Experimental_IsUnifiedHost: true,
	}
	assert.Panics(t, func() { c.IsAccountClient() })
}

func TestNewWithWorkspaceHost(t *testing.T) {
	c := &Config{
		Host:               "https://accounts.cloud.databricks.com",
		AccountID:          "123e4567-e89b-12d3-a456-426614174000",
		AzureResourceID:    "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Databricks/workspaces/test",
		ClientID:           "client-id",
		MetadataServiceURL: "http://",
		resolved:           true,
	}
	c2, err := c.NewWithWorkspaceHost("https://my-workspace.cloud.databricks.us")
	assert.NoError(t, err)
	// Host should be updated
	assert.Equal(t, "https://my-workspace.cloud.databricks.us", c2.Host)
	// Account ID and Azure Resource ID should be cleared
	assert.Equal(t, "", c2.AccountID)
	assert.Equal(t, "", c2.AzureResourceID)
	// Other fields should be preserved
	assert.Equal(t, "client-id", c2.ClientID)
	assert.Equal(t, "http://", c2.MetadataServiceURL)
	// The new config will not be automatically resolved.
	assert.False(t, c2.resolved)
}

func TestAuthenticate_InvalidHostSet(t *testing.T) {
	c := &Config{
		Host:  "https://:443",
		Token: "abcdefg",
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", c.Host, nil)
	require.NoError(t, err)
	err = c.Authenticate(req)
	assert.ErrorIs(t, err, ErrNoHostConfigured)
}

func TestConfig_getOidcEndpoints_account(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		accountID string
	}{
		{
			name:      "without trailing slash",
			host:      "https://accounts.cloud.databricks.com",
			accountID: "abc",
		},
		{
			name:      "with trailing slash",
			host:      "https://accounts.cloud.databricks.com/",
			accountID: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host:      tt.host,
				AccountID: tt.accountID,
			}
			got, err := c.getOidcEndpoints(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://accounts.cloud.databricks.com/oidc/accounts/abc/v1/authorize",
				TokenEndpoint:         "https://accounts.cloud.databricks.com/oidc/accounts/abc/v1/token",
			}, got)
		})
	}
}

func TestConfig_getOidcEndpoints_workspace(t *testing.T) {
	tests := []struct {
		name string
		host string
	}{
		{
			name: "without trailing slash",
			host: "https://myworkspace.cloud.databricks.com",
		},
		{
			name: "with trailing slash",
			host: "https://myworkspace.cloud.databricks.com/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host: tt.host,
				HTTPTransport: fixtures.SliceTransport{
					{
						Method:   "GET",
						Resource: "/oidc/.well-known/oauth-authorization-server",
						Status:   200,
						Response: `{"authorization_endpoint": "https://myworkspace.cloud.databricks.com/oidc/v1/authorize", "token_endpoint": "https://myworkspace.cloud.databricks.com/oidc/v1/token"}`,
					},
				},
			}
			got, err := c.getOidcEndpoints(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://myworkspace.cloud.databricks.com/oidc/v1/authorize",
				TokenEndpoint:         "https://myworkspace.cloud.databricks.com/oidc/v1/token",
			}, got)
		})
	}
}

func TestConfig_getOidcEndpoints_unified(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		accountID string
	}{
		{
			name:      "without trailing slash",
			host:      "https://unified.cloud.databricks.com",
			accountID: "abc",
		},
		{
			name:      "with trailing slash",
			host:      "https://unified.cloud.databricks.com/",
			accountID: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host:                       tt.host,
				AccountID:                  tt.accountID,
				Experimental_IsUnifiedHost: true,
				HTTPTransport: fixtures.SliceTransport{
					{
						Method:   "GET",
						Resource: "/oidc/accounts/abc/.well-known/oauth-authorization-server",
						Status:   200,
						Response: `{"authorization_endpoint": "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/authorize", "token_endpoint": "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/token"}`,
					},
				},
			}
			got, err := c.getOidcEndpoints(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/authorize",
				TokenEndpoint:         "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/token",
			}, got)
		})
	}
}

func TestConfig_getOAuthArgument_account(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		accountID string
	}{
		{
			name:      "without trailing slash",
			host:      "https://accounts.cloud.databricks.com",
			accountID: "abc",
		},
		{
			name:      "with trailing slash",
			host:      "https://accounts.cloud.databricks.com/",
			accountID: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host:      tt.host,
				AccountID: tt.accountID,
			}
			rawGot, err := c.getOAuthArgument()
			assert.NoError(t, err)
			got, ok := rawGot.(u2m.BasicAccountOAuthArgument)
			assert.True(t, ok)
			assert.Equal(t, "https://accounts.cloud.databricks.com", got.GetAccountHost())
			assert.Equal(t, "abc", got.GetAccountId())
		})
	}
}

func TestConfig_getOAuthArgument_workspace(t *testing.T) {
	tests := []struct {
		name string
		host string
	}{
		{
			name: "without trailing slash",
			host: "https://myworkspace.cloud.databricks.com",
		},
		{
			name: "with trailing slash",
			host: "https://myworkspace.cloud.databricks.com/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host: tt.host,
			}
			rawGot, err := c.getOAuthArgument()
			assert.NoError(t, err)
			got, ok := rawGot.(u2m.BasicWorkspaceOAuthArgument)
			assert.True(t, ok)
			assert.Equal(t, "https://myworkspace.cloud.databricks.com", got.GetWorkspaceHost())
		})
	}
}

func TestConfig_getOAuthArgument_Unified(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		accountID string
	}{
		{
			name:      "without trailing slash",
			host:      "https://unified.cloud.databricks.com",
			accountID: "account-123",
		},
		{
			name:      "with trailing slash",
			host:      "https://unified.cloud.databricks.com/",
			accountID: "account-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host:                       tt.host,
				AccountID:                  tt.accountID,
				Experimental_IsUnifiedHost: true,
			}
			rawGot, err := c.getOAuthArgument()
			assert.NoError(t, err)
			got, ok := rawGot.(u2m.UnifiedOAuthArgument)
			assert.True(t, ok, "Expected UnifiedOAuthArgument")
			assert.Equal(t, "https://unified.cloud.databricks.com", got.GetHost())
			assert.Equal(t, "account-123", got.GetAccountId())
		})
	}
}

func TestConfig_HTTPHeaders(t *testing.T) {
	c := &Config{
		Host:  "https://my-workspace.cloud.databricks.com",
		Token: "test-token",
		HTTPHeaders: map[string]string{
			"X-Custom-Header-1": "value1",
			"X-Custom-Header-2": "value2",
		},
	}

	cfg, err := HTTPClientConfigFromConfig(c)
	require.NoError(t, err)

	// Create a test request and apply only the first two visitors (host/path and headers)
	// Skip other visitors that require context to be set
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/api/2.0/clusters/list", nil)
	require.NoError(t, err)
	req.URL = &url.URL{Path: "/api/2.0/clusters/list"}

	// Apply only the first two visitors (host/path rewrite and custom headers)
	for i := 0; i < 2 && i < len(cfg.Visitors); i++ {
		err = cfg.Visitors[i](req)
		require.NoError(t, err)
	}

	// Verify custom headers are set
	assert.Equal(t, "value1", req.Header.Get("X-Custom-Header-1"))
	assert.Equal(t, "value2", req.Header.Get("X-Custom-Header-2"))
}

func TestConfig_HTTPPathPrefix(t *testing.T) {
	c := &Config{
		Host:           "https://proxy.example.com",
		Token:          "test-token",
		HTTPPathPrefix: "/prefix/path",
	}

	cfg, err := HTTPClientConfigFromConfig(c)
	require.NoError(t, err)

	// Create a test request
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/api/2.0/clusters/list", nil)
	require.NoError(t, err)
	req.URL = &url.URL{Path: "/api/2.0/clusters/list"}

	// Apply only the first visitor (host/path rewrite)
	err = cfg.Visitors[0](req)
	require.NoError(t, err)

	// Verify path prefix is prepended
	assert.Equal(t, "/prefix/path/api/2.0/clusters/list", req.URL.Path)
}

func TestConfig_HTTPHeadersAndPathPrefix(t *testing.T) {
	c := &Config{
		Host:  "https://proxy.example.com",
		Token: "test-token",
		HTTPHeaders: map[string]string{
			"X-Custom-Header": "custom-value",
		},
		HTTPPathPrefix: "/prefix/path",
	}

	cfg, err := HTTPClientConfigFromConfig(c)
	require.NoError(t, err)

	// Create a test request
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/api/2.0/clusters/list", nil)
	require.NoError(t, err)
	req.URL = &url.URL{Path: "/api/2.0/clusters/list"}

	// Apply only the first two visitors (host/path rewrite and custom headers)
	for i := 0; i < 2 && i < len(cfg.Visitors); i++ {
		err = cfg.Visitors[i](req)
		require.NoError(t, err)
	}

	// Verify both custom header and path prefix
	assert.Equal(t, "custom-value", req.Header.Get("X-Custom-Header"))
	assert.Equal(t, "/prefix/path/api/2.0/clusters/list", req.URL.Path)
}

func TestConfig_HTTPHeadersFromEnvVar(t *testing.T) {
	t.Setenv("DATABRICKS_HTTP_HEADERS", "X-Custom-Header-1=value1;X-Custom-Header-2=value2")
	t.Setenv("DATABRICKS_HOST", "https://my-workspace.cloud.databricks.com")
	t.Setenv("DATABRICKS_TOKEN", "test-token")

	c := &Config{}
	err := c.EnsureResolved()
	require.NoError(t, err)

	assert.Equal(t, map[string]string{
		"X-Custom-Header-1": "value1",
		"X-Custom-Header-2": "value2",
	}, c.HTTPHeaders)
}

func TestConfig_HTTPPathPrefixFromEnvVar(t *testing.T) {
	t.Setenv("DATABRICKS_HTTP_PATH_PREFIX", "/prefix/path")
	t.Setenv("DATABRICKS_HOST", "https://my-workspace.cloud.databricks.com")
	t.Setenv("DATABRICKS_TOKEN", "test-token")

	c := &Config{}
	err := c.EnsureResolved()
	require.NoError(t, err)

	assert.Equal(t, "/prefix/path", c.HTTPPathPrefix)
}

func TestParseMapFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
		wantErr  bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: map[string]string{},
		},
		{
			name:  "single key-value pair",
			input: "key1=value1",
			expected: map[string]string{
				"key1": "value1",
			},
		},
		{
			name:  "multiple key-value pairs",
			input: "key1=value1;key2=value2",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name:  "value with equals sign",
			input: "key1=value=with=equals",
			expected: map[string]string{
				"key1": "value=with=equals",
			},
		},
		{
			name:  "empty value",
			input: "key1=",
			expected: map[string]string{
				"key1": "",
			},
		},
		{
			name:  "whitespace around pairs",
			input: " key1=value1 ; key2=value2 ",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name:  "trailing semicolon",
			input: "key1=value1;",
			expected: map[string]string{
				"key1": "value1",
			},
		},
		{
			name:    "missing equals sign",
			input:   "key1value1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMapFromString(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}
