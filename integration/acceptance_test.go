package integration

import (
	"context"
	"os/exec"
	"testing"

	"github.com/databricks/databricks-sdk-go/compute/v2"
	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/integration/env"
	"github.com/databricks/databricks-sdk-go/settings/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccDefaultCredentials(t *testing.T) {
	ctx := workspaceTest(t)
	ClustersAPI, err := compute.NewClustersClient(nil)
	require.NoError(t, err)
	versions, err := ClustersAPI.SparkVersions(ctx)
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
	ctx := workspaceTest(t)
	ClustersAPI, err := compute.NewClustersClient(&config.Config{
		Profile: GetEnvOrSkipTest(t, "DATABRICKS_CONFIG_PROFILE"),
	})
	require.NoError(t, err)
	versions, err := ClustersAPI.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

func TestAccExplicitAzureCliAuth(t *testing.T) {
	// if we don't cache these, test will get skipped with the following error
	// > Environment variable ARM_CLIENT_ID is missing
	clientID := GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	clientSecret := GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET")
	tenantID := GetEnvOrSkipTest(t, "ARM_TENANT_ID")
	resourceID := GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID")
	env.CleanupEnvironment(t)

	t.Setenv("AZURE_CONFIG_DIR", t.TempDir())

	// Login with Azure CLI
	cmd := exec.Command(
		"az",
		"login",
		"--service-principal",
		"--user", clientID,
		"--password", clientSecret,
		"--tenant", tenantID,
	)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("error running az: %s (%s)", err, out)
	}

	ClustersAPI, err := compute.NewClustersClient(&config.Config{
		AzureResourceID: resourceID,
		Credentials:     config.AzureCliCredentials{},
	})
	require.NoError(t, err)
	ctx := context.Background()
	versions, err := ClustersAPI.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

func TestAccAzureErrorMappingForUnauthenticated(t *testing.T) {
	ctx := workspaceTest(t)
	ClustersAPI, err := compute.NewClustersClient(&config.Config{
		DebugHeaders:      true,
		AzureTenantID:     GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AzureClientID:     GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret: "invalid-for-integration-tests",
		AzureResourceID:   "/a/b/c",
	})
	require.NoError(t, err)
	_, err = ClustersAPI.SparkVersions(ctx)
	require.ErrorIs(t, err, apierr.ErrUnauthenticated)
}

func TestAccExplicitAzureSpnAuth(t *testing.T) {
	ctx := workspaceTest(t)
	ClustersAPI, err := compute.NewClustersClient(&config.Config{
		AzureTenantID:     GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AzureClientID:     GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret: GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		AzureResourceID:   GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		Credentials:       config.AzureClientSecretCredentials{},
	})
	require.NoError(t, err)
	versions, err := ClustersAPI.SparkVersions(ctx)
	require.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

// Confirm that we can access the account IP access lists using MSI credentials.
func TestAccAccountsMsiCredentials(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "ARM_USE_MSI"))
	ctx, cfg := accountTest(t)
	IpAccessListsAPI, err := settings.NewIpAccessListsClient(cfg)
	require.NoError(t, err)
	_, err = IpAccessListsAPI.ListAll(ctx)
	require.NoError(t, err)
}
