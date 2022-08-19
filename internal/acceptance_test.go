package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccDefaultCredentials(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ws := workspaces.New()

	ctx := context.Background()
	versions, err := ws.Clusters.ListSparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(clusters.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

// TODO: add CredentialProviderChain

func TestAccExplicitDatabricksCfg(t *testing.T) {
	ws := workspaces.New(&databricks.Config{
		Profile: GetEnvOrSkipTest(t, "DATABRICKS_CONFIG_PROFILE"),
	})
	ctx := context.Background()
	versions, err := ws.Clusters.ListSparkVersions(ctx)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(versions.SparkVersions), 10)
}

func TestAccExplicitAzureCliAuth(t *testing.T) {
	ws := workspaces.New(&databricks.Config{
		AzureResourceID: GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		Credentials:     databricks.AzureCliCredentials{},
	})
	ctx := context.Background()
	versions, err := ws.Clusters.ListSparkVersions(ctx)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(versions.SparkVersions), 10)
}

func TestAccExplicitAzureSpnAuth(t *testing.T) {
	ws := workspaces.New(&databricks.Config{
		AzureTenantID:     GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AzureClientID:     GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret: GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		AzureResourceID:   GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		Credentials:       databricks.AzureClientSecretCredentials{},
	})
	ctx := context.Background()
	versions, err := ws.Clusters.ListSparkVersions(ctx)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(versions.SparkVersions), 10)
}
