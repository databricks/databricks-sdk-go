package internal

import (
	"context"
	"os"
	"os/exec"
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
	versions, err := ws.Clusters.GetSparkVersions(ctx)
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
	versions, err := ws.Clusters.GetSparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(clusters.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

func temporaryDirectory(t *testing.T) string {
	dir, err := os.MkdirTemp(os.TempDir(), "tmp.*")
	if err != nil {
		t.Fatal(err)
	}

	// Remove temporary directory upon test completion.
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	return dir
}

func setEnv(t *testing.T, key, value string) {
	cur, ok := os.LookupEnv(key)
	os.Setenv(key, value)

	// Revert environment variable upon test completion.
	t.Cleanup(func() {
		if ok {
			os.Setenv(key, cur)
		} else {
			os.Unsetenv(key)
		}
	})
}

func TestAccExplicitAzureCliAuth(t *testing.T) {
	setEnv(t, "AZURE_CONFIG_DIR", temporaryDirectory(t))

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
	versions, err := ws.Clusters.GetSparkVersions(ctx)
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
	versions, err := ws.Clusters.GetSparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(clusters.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}
