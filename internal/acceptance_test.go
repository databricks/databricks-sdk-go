package internal

import (
	"context"
	"testing"

	"github.com/databricks/sdk-go/databricks"
	"github.com/databricks/sdk-go/databricks/client"
	"github.com/stretchr/testify/assert"
)

func TestAccDefaultCredentials(t *testing.T) {
	var resp map[string]interface{}
	client := client.New()
	err := client.Get(context.Background(), "/token/list", nil, &resp)
	assert.NoError(t, err)
	_, ok := resp["token_infos"]
	assert.True(t, ok)
}

// TODO: add CredentialProviderChain

func TestAccExplicitDatabricksCfg(t *testing.T) {
	var resp map[string]interface{}
	client := client.New(&databricks.Config{
		Profile: GetEnvOrSkipTest(t, "DATABRICKS_CONFIG_PROFILE"),
	})
	err := client.Get(context.Background(), "/token/list", nil, &resp)
	assert.NoError(t, err)
	_, ok := resp["token_infos"]
	assert.True(t, ok)
}

func TestAccExplicitAzureCliAuth(t *testing.T) {
	var resp map[string]interface{}
	client := client.New(&databricks.Config{
		Host:        GetEnvOrSkipTest(t, "DATABRICKS_HOST"),
		Credentials: databricks.AzureCliCredentials{},
	})
	err := client.Get(context.Background(), "/token/list", nil, &resp)
	assert.NoError(t, err)
	_, ok := resp["token_infos"]
	assert.True(t, ok)
}

func TestAccExplicitAzureSpnAuth(t *testing.T) {
	var resp map[string]interface{}
	client := client.New(&databricks.Config{
		AzureTenantID:     GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AzureClientID:     GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret: GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		AzureResourceID:   GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		Credentials:       databricks.AzureClientSecretCredentials{},
		DebugHeaders:      true,
	})
	err := client.Get(context.Background(), "/preview/scim/Me", nil, &resp)
	assert.NoError(t, err)
	_, ok := resp["userName"]
	assert.True(t, ok)
}
