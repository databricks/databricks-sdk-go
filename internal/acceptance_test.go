package internal

import (
	"context"
	"os/exec"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccDefaultCredentials(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ws := workspaces.New()

	ctx := context.Background()
	versions, err := ws.Clusters.SparkVersions(ctx)
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
	versions, err := ws.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(clusters.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

func TestAccExplicitAzureCliAuth(t *testing.T) {
	env.CleanupEnvironment(t)

	t.Setenv("AZURE_CONFIG_DIR", t.TempDir())

	// Login with Azure CLI
	cmd := exec.Command(
		"az",
		"login",
		"--service-principal",
		"--user", GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		"--password", GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		"--tenant", GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
	)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("error running az: %s (%s)", err, out)
	}

	ws := workspaces.New(&databricks.Config{
		AzureResourceID: GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		Credentials:     databricks.AzureCliCredentials{},
	})
	ctx := context.Background()
	versions, err := ws.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(clusters.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
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
	versions, err := ws.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(clusters.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

func TestAccDatabricksOauth(t *testing.T) {
	ws := workspaces.New(&databricks.Config{
		Host:         GetEnvOrSkipTest(t, "DATABRICKS_HOST"),
		ClientID:     GetEnvOrSkipTest(t, "DATABRICKS_CLIENT_ID"),
		ClientSecret: GetEnvOrSkipTest(t, "DATABRICKS_CLIENT_SECRET"),
	})

	ctx := context.Background()
	versions, err := ws.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(clusters.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}
