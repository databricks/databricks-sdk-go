package config

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockLoader is a test helper that implements the Loader interface.
type mockLoader func(cfg *Config) error

func (m mockLoader) Name() string {
	return "mockLoader"
}

func (m mockLoader) Configure(cfg *Config) error {
	return m(cfg)
}

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

func TestConfig_getOAuthArgument_profileCacheKeys(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	tests := []struct {
		name        string
		config      *Config
		wantKey     string
		wantHostKey string
	}{
		{
			name: "workspace without profile",
			config: &Config{
				Host:    "https://myworkspace.cloud.databricks.com",
				Loaders: []Loader{noopLoader},
			},
			wantKey:     "https://myworkspace.cloud.databricks.com",
			wantHostKey: "https://myworkspace.cloud.databricks.com",
		},
		{
			name: "workspace with profile",
			config: &Config{
				Host:    "https://myworkspace.cloud.databricks.com",
				Profile: "ws-profile",
				Loaders: []Loader{noopLoader},
			},
			wantKey:     "ws-profile",
			wantHostKey: "https://myworkspace.cloud.databricks.com",
		},
		{
			name: "account without profile",
			config: &Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "abc",
				Loaders:   []Loader{noopLoader},
			},
			wantKey:     "https://accounts.cloud.databricks.com/oidc/accounts/abc",
			wantHostKey: "https://accounts.cloud.databricks.com/oidc/accounts/abc",
		},
		{
			name: "account with profile",
			config: &Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "abc",
				Profile:   "my-profile",
				Loaders:   []Loader{noopLoader},
			},
			wantKey:     "my-profile",
			wantHostKey: "https://accounts.cloud.databricks.com/oidc/accounts/abc",
		},
		{
			name: "unified without profile",
			config: &Config{
				Host:                       "https://unified.cloud.databricks.com",
				AccountID:                  "account-123",
				Experimental_IsUnifiedHost: true,
				Loaders:                    []Loader{noopLoader},
			},
			wantKey:     "https://unified.cloud.databricks.com/oidc/accounts/account-123",
			wantHostKey: "https://unified.cloud.databricks.com/oidc/accounts/account-123",
		},
		{
			name: "unified with profile",
			config: &Config{
				Host:                       "https://unified.cloud.databricks.com",
				AccountID:                  "account-123",
				Profile:                    "unified-profile",
				Experimental_IsUnifiedHost: true,
				Loaders:                    []Loader{noopLoader},
			},
			wantKey:     "unified-profile",
			wantHostKey: "https://unified.cloud.databricks.com/oidc/accounts/account-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawGot, err := tt.config.getOAuthArgument()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantKey, rawGot.GetCacheKey())
			if hcp, ok := rawGot.(u2m.HostCacheKeyProvider); ok {
				assert.Equal(t, tt.wantHostKey, hcp.GetHostCacheKey())
			}
		})
	}
}

func TestConfig_EnsureResolved_scopeNormalization(t *testing.T) {
	testCases := []struct {
		desc   string
		scopes []string
		want   []string
	}{
		{
			desc:   "nil scopes",
			scopes: nil,
			want:   nil,
		},
		{
			desc:   "empty scopes",
			scopes: []string{},
			want:   []string{},
		},
		{
			desc:   "single scope",
			scopes: []string{"clusters"},
			want:   []string{"clusters"},
		},
		{
			desc:   "already sorted no duplicates",
			scopes: []string{"a", "b", "c"},
			want:   []string{"a", "b", "c"},
		},
		{
			desc:   "unsorted scopes are sorted",
			scopes: []string{"jobs", "clusters", "pipelines"},
			want:   []string{"clusters", "jobs", "pipelines"},
		},
		{
			desc:   "duplicate scopes are removed",
			scopes: []string{"clusters", "jobs", "clusters", "pipelines:read", "jobs"},
			want:   []string{"clusters", "jobs", "pipelines:read"},
		},
		{
			desc:   "all duplicates reduced to one",
			scopes: []string{"all-apis", "all-apis", "all-apis"},
			want:   []string{"all-apis"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cfg := &Config{
				Host: "https://example.cloud.databricks.com",
				Loaders: []Loader{mockLoader(func(cfg *Config) error {
					cfg.Scopes = tc.scopes
					return nil
				})},
			}

			err := cfg.EnsureResolved()
			if err != nil {
				t.Fatalf("EnsureResolved() error: %v", err)
			}

			if diff := cmp.Diff(tc.want, cfg.Scopes); diff != "" {
				t.Errorf("EnsureResolved() scopes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestConfig_DiscoveryURL_FromEnv(t *testing.T) {
	t.Setenv("DATABRICKS_DISCOVERY_URL", "https://custom.idp.example.com/oidc")
	cfg := &Config{Host: testHMHost, Token: "t"}
	if err := cfg.EnsureResolved(); err != nil {
		t.Fatal(err)
	}
	if cfg.DiscoveryURL != "https://custom.idp.example.com/oidc" {
		t.Errorf("unexpected DiscoveryURL: %q", cfg.DiscoveryURL)
	}
}

func TestConfig_getOidcEndpoints_UsesDiscoveryURL(t *testing.T) {
	discoveryURL := testHMHost + "/oidc"
	cfg := &Config{
		Host:         testHMHost,
		Token:        "t",
		DiscoveryURL: discoveryURL,
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/oidc",
				Status:   200,
				Response: `{"authorization_endpoint": "` + testHMHost + `/oidc/v1/authorize", "token_endpoint": "` + testHMHost + `/oidc/v1/token"}`,
			},
		},
	}
	endpoints, err := cfg.getOidcEndpoints(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	want := &u2m.OAuthAuthorizationServer{
		AuthorizationEndpoint: testHMHost + "/oidc/v1/authorize",
		TokenEndpoint:         testHMHost + "/oidc/v1/token",
	}
	if diff := cmp.Diff(want, endpoints); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestConfig_ResolveHostMetadata_WorkspacePopulatesAllFields(t *testing.T) {
	cfg := &Config{
		Host:  testHMHost,
		Token: "t",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   200,
				Response: `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "workspace_id": "` + testHMWorkspaceID + `"}`,
			},
		},
	}
	if err := cfg.resolveHostMetadata(context.Background()); err != nil {
		t.Fatal(err)
	}
	if cfg.AccountID != testHMAccountID {
		t.Errorf("unexpected AccountID: %q", cfg.AccountID)
	}
	if cfg.WorkspaceID != testHMWorkspaceID {
		t.Errorf("unexpected WorkspaceID: %q", cfg.WorkspaceID)
	}
	if cfg.DiscoveryURL != testHMHost+"/oidc" {
		t.Errorf("unexpected DiscoveryURL: %q", cfg.DiscoveryURL)
	}
}

func TestConfig_ResolveHostMetadata_AccountSubstitutesAccountID(t *testing.T) {
	cfg := &Config{
		Host:      testHMAccHost,
		Token:     "t",
		AccountID: testHMAccountID,
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   200,
				Response: `{"oidc_endpoint": "` + testHMAccHost + `/oidc/accounts/{account_id}"}`,
			},
		},
	}
	if err := cfg.resolveHostMetadata(context.Background()); err != nil {
		t.Fatal(err)
	}
	want := testHMAccHost + "/oidc/accounts/" + testHMAccountID
	if cfg.DiscoveryURL != want {
		t.Errorf("unexpected DiscoveryURL: %q", cfg.DiscoveryURL)
	}
}

func TestConfig_ResolveHostMetadata_DoesNotOverwriteExistingFields(t *testing.T) {
	cfg := &Config{
		Host:        testHMHost,
		Token:       "t",
		AccountID:   testHMAccountID,
		WorkspaceID: testHMWorkspaceID,
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   200,
				Response: `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "other-account", "workspace_id": "other-ws"}`,
			},
		},
	}
	if err := cfg.resolveHostMetadata(context.Background()); err != nil {
		t.Fatal(err)
	}
	if cfg.AccountID != testHMAccountID {
		t.Errorf("AccountID was overwritten: %q", cfg.AccountID)
	}
	if cfg.WorkspaceID != testHMWorkspaceID {
		t.Errorf("WorkspaceID was overwritten: %q", cfg.WorkspaceID)
	}
}

func TestConfig_ResolveHostMetadata_MissingAccountID(t *testing.T) {
	cfg := &Config{
		Host:  testHMAccHost,
		Token: "t",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   200,
				Response: `{"oidc_endpoint": "` + testHMAccHost + `/oidc/accounts/{account_id}"}`,
			},
		},
	}
	err := cfg.resolveHostMetadata(context.Background())
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "account_id is not configured") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestConfig_ResolveHostMetadata_MissingOIDCEndpoint(t *testing.T) {
	cfg := &Config{
		Host:  testHMHost,
		Token: "t",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   200,
				Response: `{"account_id": "` + testHMAccountID + `"}`,
			},
		},
	}
	err := cfg.resolveHostMetadata(context.Background())
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "discovery_url is not configured") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestConfig_ResolveHostMetadata_HTTPError(t *testing.T) {
	cfg := &Config{
		Host:  testHMHost,
		Token: "t",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   500,
				Response: `{"error": "internal error"}`,
			},
		},
	}
	err := cfg.resolveHostMetadata(context.Background())
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "fetching host metadata from") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestConfig_ResolveHostMetadata_NoHost(t *testing.T) {
	cfg := &Config{}
	if err := cfg.resolveHostMetadata(context.Background()); err != nil {
		t.Fatal(err)
	}
}
