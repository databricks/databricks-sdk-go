package internal

import (
	"context"
	"testing"

	"github.com/databricks/sdk-go/core/auth/azure"
	"github.com/databricks/sdk-go/core/client"
	"github.com/databricks/sdk-go/core/config"
	"github.com/stretchr/testify/assert"
)

func TestAccExplicitAzureCliAuth(t *testing.T) {
	var resp map[string]interface{}
	client := &client.ApiClient{
		Config: &config.Config{
			Host: GetEnvOrSkipTest(t, "DATABRICKS_HOST"),
		},
		Credentials: azure.AzureCliCredentials{},
	}
	err := client.Get(context.Background(), "/token/list", nil, &resp)
	assert.NoError(t, err)
	_, ok := resp["token_infos"]
	assert.True(t, ok)
}

func TestAccExplicitAzureSpnAuth(t *testing.T) {
	var resp map[string]interface{}
	client := &client.ApiClient{
		Config: &config.Config{
			DebugHeaders: true,
		},
		Credentials: azure.ClientSecretCredentials{
			TenantID:        GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
			ClientID:        GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
			ClientSecret:    GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
			AzureResourceID: GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		},
	}
	err := client.Get(context.Background(), "/preview/scim/Me", nil, &resp)
	assert.NoError(t, err)
	_, ok := resp["userName"]
	assert.True(t, ok)
}

