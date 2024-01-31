package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.True(t, c2.resolved)
}
