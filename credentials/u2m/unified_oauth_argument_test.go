package u2m

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBasicUnifiedOAuthArgument(t *testing.T) {
	arg, err := NewBasicUnifiedOAuthArgument("https://unified.databricks.com", "account-123")
	assert.NoError(t, err)
	assert.Equal(t, "https://unified.databricks.com", arg.GetHost())
	assert.Equal(t, "account-123", arg.GetAccountId())
	assert.Equal(t, "https://unified.databricks.com/oidc/accounts/account-123", arg.GetCacheKey())
}

func TestNewBasicUnifiedOAuthArgument_ValidatesHost(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		accountID string
		wantErr   string
	}{
		{
			name:      "invalid http protocol",
			host:      "http://insecure.com",
			accountID: "account-123",
			wantErr:   "host must start with 'https://': http://insecure.com",
		},
		{
			name:      "trailing slash",
			host:      "https://unified.databricks.com/",
			accountID: "account-123",
			wantErr:   "host must not have a trailing slash: https://unified.databricks.com/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBasicUnifiedOAuthArgument(tt.host, tt.accountID)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestBasicUnifiedOAuthArgument_GetHost(t *testing.T) {
	arg, _ := NewBasicUnifiedOAuthArgument("https://unified.databricks.com", "account-123")
	assert.Equal(t, "https://unified.databricks.com", arg.GetHost())
}

func TestBasicUnifiedOAuthArgument_GetAccountId(t *testing.T) {
	arg, _ := NewBasicUnifiedOAuthArgument("https://unified.databricks.com", "account-123")
	assert.Equal(t, "account-123", arg.GetAccountId())
}

func TestBasicUnifiedOAuthArgument_GetCacheKey(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		accountID string
		wantKey   string
	}{
		{
			name:      "standard case",
			host:      "https://unified.databricks.com",
			accountID: "account-123",
			wantKey:   "https://unified.databricks.com/oidc/accounts/account-123",
		},
		{
			name:      "different account",
			host:      "https://unified.databricks.com",
			accountID: "account-456",
			wantKey:   "https://unified.databricks.com/oidc/accounts/account-456",
		},
		{
			name:      "different host",
			host:      "https://other-unified.databricks.com",
			accountID: "account-123",
			wantKey:   "https://other-unified.databricks.com/oidc/accounts/account-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg, err := NewBasicUnifiedOAuthArgument(tt.host, tt.accountID)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantKey, arg.GetCacheKey())
		})
	}
}

func TestNewProfileUnifiedOAuthArgument(t *testing.T) {
	arg, err := NewProfileUnifiedOAuthArgument("https://unified.databricks.com", "account-123", "my-profile")
	assert.NoError(t, err)
	assert.Equal(t, "https://unified.databricks.com", arg.GetHost())
	assert.Equal(t, "account-123", arg.GetAccountId())
	assert.Equal(t, "my-profile", arg.GetCacheKey())
	assert.Equal(t, "https://unified.databricks.com/oidc/accounts/account-123", arg.GetHostCacheKey())
}

func TestNewProfileUnifiedOAuthArgument_ValidatesHost(t *testing.T) {
	_, err := NewProfileUnifiedOAuthArgument("http://insecure.com", "account-123", "my-profile")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "host must start with 'https://'")
}

func TestBasicUnifiedOAuthArgument_ProfileCacheKeys(t *testing.T) {
	tests := []struct {
		name        string
		host        string
		accountID   string
		profile     string
		wantKey     string
		wantHostKey string
	}{
		{
			name:        "without profile returns host-based key",
			host:        "https://unified.databricks.com",
			accountID:   "account-123",
			wantKey:     "https://unified.databricks.com/oidc/accounts/account-123",
			wantHostKey: "https://unified.databricks.com/oidc/accounts/account-123",
		},
		{
			name:        "with profile returns profile name",
			host:        "https://unified.databricks.com",
			accountID:   "account-123",
			profile:     "my-profile",
			wantKey:     "my-profile",
			wantHostKey: "https://unified.databricks.com/oidc/accounts/account-123",
		},
		{
			name:        "empty profile returns host-based key",
			host:        "https://unified.databricks.com",
			accountID:   "account-123",
			profile:     "",
			wantKey:     "https://unified.databricks.com/oidc/accounts/account-123",
			wantHostKey: "https://unified.databricks.com/oidc/accounts/account-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var arg BasicUnifiedOAuthArgument
			var err error
			if tt.profile != "" {
				arg, err = NewProfileUnifiedOAuthArgument(tt.host, tt.accountID, tt.profile)
			} else {
				arg, err = NewBasicUnifiedOAuthArgument(tt.host, tt.accountID)
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantKey, arg.GetCacheKey())
			assert.Equal(t, tt.wantHostKey, arg.GetHostCacheKey())
		})
	}
}

func TestBasicUnifiedOAuthArgument_ImplementsHostCacheKeyProvider(t *testing.T) {
	arg, err := NewBasicUnifiedOAuthArgument("https://unified.databricks.com", "account-123")
	assert.NoError(t, err)
	var _ HostCacheKeyProvider = arg
}
