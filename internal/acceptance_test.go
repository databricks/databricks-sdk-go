package internal

import (
	"context"
	"os/exec"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccDefaultCredentials(t *testing.T) {
	t.Log(qa.GetEnvOrSkipTest(t, "CLOUD_ENV"))
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
		Profile: qa.GetEnvOrSkipTest(t, "DATABRICKS_CONFIG_PROFILE"),
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
		"--user", qa.GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		"--password", qa.GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		"--tenant", qa.GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
	)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("error running az: %s (%s)", err, out)
	}

	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		AzureResourceID: qa.GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
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
		AzureTenantID:     qa.GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AzureClientID:     qa.GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret: qa.GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		AzureResourceID:   qa.GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
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
