package config

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/internal/credentials/oauth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsAccountClient_AwsAccount(t *testing.T) {
	c := &Config{
		Host:      "https://accounts.cloud.databricks.com",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.True(t, c.IsAccountClient())
}

func TestIsAccountClient_AwsDodAccount(t *testing.T) {
	c := &Config{
		Host:      "https://accounts-dod.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.True(t, c.IsAccountClient())
}

func TestIsAccountClient_AwsWorkspace(t *testing.T) {
	c := &Config{
		Host:      "https://my-workspace.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.False(t, c.IsAccountClient())
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
	c := &Config{
		Host:      "https://accounts.cloud.databricks.com",
		AccountID: "abc",
	}
	got, err := c.getOidcEndpoints(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, &oauth.OAuthAuthorizationServer{
		AuthorizationEndpoint: "https://accounts.cloud.databricks.com/oidc/accounts/abc/v1/authorize",
		TokenEndpoint:         "https://accounts.cloud.databricks.com/oidc/accounts/abc/v1/token",
	}, got)
}

func TestConfig_getOidcEndpoints_workspace(t *testing.T) {
	c := &Config{
		Host: "https://myworkspace.cloud.databricks.com",
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
	assert.Equal(t, &oauth.OAuthAuthorizationServer{
		AuthorizationEndpoint: "https://myworkspace.cloud.databricks.com/oidc/v1/authorize",
		TokenEndpoint:         "https://myworkspace.cloud.databricks.com/oidc/v1/token",
	}, got)
}
