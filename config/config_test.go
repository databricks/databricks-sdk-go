package config

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

// mockLoader is a test helper that implements the Loader interface.
type mockLoader func(cfg *Config) error

func (m mockLoader) Name() string {
	return "mockLoader"
}

func (m mockLoader) Configure(cfg *Config) error {
	return m(cfg)
}

func TestConfig_HostType(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
		want   HostType
	}{
		{
			name: "AWS account",
			config: &Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "123e4567-e89b-12d3-a456-426614174000",
			},
			want: AccountHost,
		},
		{
			name: "AWS DoD account",
			config: &Config{
				Host:      "https://accounts-dod.cloud.databricks.us",
				AccountID: "123e4567-e89b-12d3-a456-426614174000",
			},
			want: AccountHost,
		},
		{
			name: "AWS account without scheme",
			config: &Config{
				Host:      "accounts.cloud.databricks.com",
				AccountID: "123e4567-e89b-12d3-a456-426614174000",
			},
			want: AccountHost,
		},
		{
			name: "AWS DoD account without scheme",
			config: &Config{
				Host:      "accounts-dod.cloud.databricks.us",
				AccountID: "123e4567-e89b-12d3-a456-426614174000",
			},
			want: AccountHost,
		},
		{
			name: "AWS workspace",
			config: &Config{
				Host:      "https://my-workspace.cloud.databricks.us",
				AccountID: "123e4567-e89b-12d3-a456-426614174000",
			},
			want: WorkspaceHost,
		},
		{
			name: "unified",
			config: &Config{
				Host:                       "https://unified.cloud.databricks.com",
				AccountID:                  "123e4567-e89b-12d3-a456-426614174000",
				Experimental_IsUnifiedHost: true,
			},
			want: UnifiedHost,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.config.HostType()
			if got != tt.want {
				t.Errorf("HostType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAccountClient_PanicsOnUnifiedHost(t *testing.T) {
	c := &Config{
		Host:                       "https://unified.cloud.databricks.com",
		AccountID:                  "test-account",
		Experimental_IsUnifiedHost: true,
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected IsAccountClient() to panic, but it did not")
		}
	}()
	c.IsAccountClient()
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
	if err != nil {
		t.Fatalf("NewWithWorkspaceHost() error: %v", err)
	}
	if c2.Host != "https://my-workspace.cloud.databricks.us" {
		t.Errorf("Host = %q, want %q", c2.Host, "https://my-workspace.cloud.databricks.us")
	}
	if c2.AccountID != "" {
		t.Errorf("AccountID = %q, want empty", c2.AccountID)
	}
	if c2.AzureResourceID != "" {
		t.Errorf("AzureResourceID = %q, want empty", c2.AzureResourceID)
	}
	if c2.ClientID != "client-id" {
		t.Errorf("ClientID = %q, want %q", c2.ClientID, "client-id")
	}
	if c2.MetadataServiceURL != "http://" {
		t.Errorf("MetadataServiceURL = %q, want %q", c2.MetadataServiceURL, "http://")
	}
	if c2.resolved {
		t.Error("resolved = true, want false")
	}
}

func TestAuthenticate_InvalidHostSet(t *testing.T) {
	c := &Config{
		Host:  "https://:443",
		Token: "abcdefg",
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", c.Host, nil)
	if err != nil {
		t.Fatalf("NewRequestWithContext() error: %v", err)
	}
	err = c.Authenticate(req)
	if !errors.Is(err, ErrNoHostConfigured) {
		t.Fatalf("Authenticate() error = %v, want %v", err, ErrNoHostConfigured)
	}
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
			if err != nil {
				t.Fatalf("getOidcEndpoints() error: %v", err)
			}
			want := &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://accounts.cloud.databricks.com/oidc/accounts/abc/v1/authorize",
				TokenEndpoint:         "https://accounts.cloud.databricks.com/oidc/accounts/abc/v1/token",
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("getOidcEndpoints() mismatch (-want +got):\n%s", diff)
			}
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
			if err != nil {
				t.Fatalf("getOidcEndpoints() error: %v", err)
			}
			want := &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://myworkspace.cloud.databricks.com/oidc/v1/authorize",
				TokenEndpoint:         "https://myworkspace.cloud.databricks.com/oidc/v1/token",
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("getOidcEndpoints() mismatch (-want +got):\n%s", diff)
			}
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
			if err != nil {
				t.Fatalf("getOidcEndpoints() error: %v", err)
			}
			want := &u2m.OAuthAuthorizationServer{
				AuthorizationEndpoint: "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/authorize",
				TokenEndpoint:         "https://unified.cloud.databricks.com/oidc/accounts/abc/v1/token",
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("getOidcEndpoints() mismatch (-want +got):\n%s", diff)
			}
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
			if err != nil {
				t.Fatalf("getOAuthArgument() error: %v", err)
			}
			got, ok := rawGot.(u2m.BasicAccountOAuthArgument)
			if !ok {
				t.Fatalf("getOAuthArgument() returned %T, want BasicAccountOAuthArgument", rawGot)
			}
			if got.GetAccountHost() != "https://accounts.cloud.databricks.com" {
				t.Errorf("GetAccountHost() = %q, want %q", got.GetAccountHost(), "https://accounts.cloud.databricks.com")
			}
			if got.GetAccountId() != "abc" {
				t.Errorf("GetAccountId() = %q, want %q", got.GetAccountId(), "abc")
			}
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
			if err != nil {
				t.Fatalf("getOAuthArgument() error: %v", err)
			}
			got, ok := rawGot.(u2m.BasicWorkspaceOAuthArgument)
			if !ok {
				t.Fatalf("getOAuthArgument() returned %T, want BasicWorkspaceOAuthArgument", rawGot)
			}
			if got.GetWorkspaceHost() != "https://myworkspace.cloud.databricks.com" {
				t.Errorf("GetWorkspaceHost() = %q, want %q", got.GetWorkspaceHost(), "https://myworkspace.cloud.databricks.com")
			}
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
			if err != nil {
				t.Fatalf("getOAuthArgument() error: %v", err)
			}
			got, ok := rawGot.(u2m.UnifiedOAuthArgument)
			if !ok {
				t.Fatalf("getOAuthArgument() returned %T, want UnifiedOAuthArgument", rawGot)
			}
			if got.GetHost() != "https://unified.cloud.databricks.com" {
				t.Errorf("GetHost() = %q, want %q", got.GetHost(), "https://unified.cloud.databricks.com")
			}
			if got.GetAccountId() != "account-123" {
				t.Errorf("GetAccountId() = %q, want %q", got.GetAccountId(), "account-123")
			}
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
			if err != nil {
				t.Fatalf("getOAuthArgument() error: %v", err)
			}
			if rawGot.GetCacheKey() != tt.wantKey {
				t.Errorf("GetCacheKey() = %q, want %q", rawGot.GetCacheKey(), tt.wantKey)
			}
			if hcp, ok := rawGot.(u2m.HostCacheKeyProvider); ok {
				if hcp.GetHostCacheKey() != tt.wantHostKey {
					t.Errorf("GetHostCacheKey() = %q, want %q", hcp.GetHostCacheKey(), tt.wantHostKey)
				}
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

func TestConfig_ResolveHostMetadata_PopulatesCloudFromAPI(t *testing.T) {
	cfg := &Config{
		Host:  testHMHost,
		Token: "t",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   200,
				Response: `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "Azure"}`,
			},
		},
	}
	if err := cfg.resolveHostMetadata(context.Background()); err != nil {
		t.Fatal(err)
	}
	if cfg.Cloud != "Azure" {
		t.Errorf("unexpected Cloud: %q", cfg.Cloud)
	}
}

func TestConfig_ResolveHostMetadata_CloudFallbackToDNS(t *testing.T) {
	cfg := &Config{
		Host:  "https://my-workspace.azuredatabricks.net",
		Token: "t",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   200,
				Response: `{"oidc_endpoint": "https://my-workspace.azuredatabricks.net/oidc", "account_id": "` + testHMAccountID + `"}`,
			},
		},
	}
	if err := cfg.resolveHostMetadata(context.Background()); err != nil {
		t.Fatal(err)
	}
	if cfg.Cloud != "Azure" {
		t.Errorf("unexpected Cloud from DNS fallback: %q", cfg.Cloud)
	}
}

func TestConfig_ResolveHostMetadata_DoesNotOverwriteExistingCloud(t *testing.T) {
	cfg := &Config{
		Host:  testHMHost,
		Token: "t",
		Cloud: "GCP",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/.well-known/databricks-config",
				Status:   200,
				Response: `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "AWS"}`,
			},
		},
	}
	if err := cfg.resolveHostMetadata(context.Background()); err != nil {
		t.Fatal(err)
	}
	if cfg.Cloud != "GCP" {
		t.Errorf("Cloud was overwritten: got %q, want %q", cfg.Cloud, "GCP")
	}
}

func TestConfig_ResolveHostMetadata_Clouds(t *testing.T) {
	tests := []struct {
		name      string
		cloudJSON string
		wantCloud string
	}{
		{
			name:      "AWS",
			cloudJSON: "AWS",
			wantCloud: "AWS",
		},
		{
			name:      "Azure",
			cloudJSON: "Azure",
			wantCloud: "Azure",
		},
		{
			name:      "GCP",
			cloudJSON: "GCP",
			wantCloud: "GCP",
		},
		{
			name:      "aws lowercase",
			cloudJSON: "aws",
			wantCloud: "AWS",
		},
		{
			name:      "AWS uppercase",
			cloudJSON: "AWS",
			wantCloud: "AWS",
		},
		{
			name:      "azure lowercase",
			cloudJSON: "azure",
			wantCloud: "Azure",
		},
		{
			name:      "AZURE uppercase",
			cloudJSON: "AZURE",
			wantCloud: "Azure",
		},
		{
			name:      "Azure title case",
			cloudJSON: "Azure",
			wantCloud: "Azure",
		},
		{
			name:      "gcp lowercase",
			cloudJSON: "gcp",
			wantCloud: "GCP",
		},
		{
			name:      "GCP uppercase",
			cloudJSON: "GCP",
			wantCloud: "GCP",
		},
		{
			name:      "Another cloud is supported",
			cloudJSON: "Another",
			wantCloud: "Another",
		},
		{
			name:      "Unknown cloud falls back to DNS",
			cloudJSON: "",
			wantCloud: "AWS", // Falls back to DNS-based detection for testHMHost
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &Config{
				Host:  testHMHost,
				Token: "t",
				HTTPTransport: fixtures.SliceTransport{
					{
						Method:   "GET",
						Resource: "/.well-known/databricks-config",
						Status:   200,
						Response: `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "cloud": "` + tc.cloudJSON + `"}`,
					},
				},
			}
			if err := cfg.resolveHostMetadata(context.Background()); err != nil {
				t.Fatal(err)
			}
			if string(cfg.Cloud) != tc.wantCloud {
				t.Errorf("unexpected Cloud: got %q, want %q", cfg.Cloud, tc.wantCloud)
			}
		})
	}
}
