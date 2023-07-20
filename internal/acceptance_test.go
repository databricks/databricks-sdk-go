package internal

import (
	"context"
	"github.com/xuxiaoshuo/databricks-sdk-go"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuxiaoshuo/databricks-sdk-go/config"
	"github.com/xuxiaoshuo/databricks-sdk-go/internal/env"
	"github.com/xuxiaoshuo/databricks-sdk-go/service/compute"
)

func TestAccDefaultCredentials(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountClient() {
		t.SkipNow()
	}
	ctx := context.Background()
	versions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

// TODO: add CredentialProviderChain

func TestAccExplicitDatabricksCfg(t *testing.T) {
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Profile: GetEnvOrSkipTest(t, "DATABRICKS_CONFIG_PROFILE"),
	}))
	if w.Config.IsAccountClient() {
		t.SkipNow()
	}
	ctx := context.Background()
	versions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
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

	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		AzureResourceID: GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		Credentials:     config.AzureCliCredentials{},
	}))
	ctx := context.Background()
	versions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

func TestAccExplicitAzureSpnAuth(t *testing.T) {
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		AzureTenantID:     GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AzureClientID:     GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret: GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		AzureResourceID:   GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		Credentials:       config.AzureClientSecretCredentials{},
	}))
	ctx := context.Background()
	versions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}
