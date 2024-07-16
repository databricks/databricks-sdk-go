package internal

import (
	"context"
	"os/exec"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
)

func init() {
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
}

func TestAccDefaultCredentials(t *testing.T) {
	workspaceTest(t)
	w := databricks.Must(databricks.NewWorkspaceClient())
	ctx := context.Background()
	versions, err := w.Clusters.SparkVersions(ctx)
	assert.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

// TODO: add CredentialProviderChain

func TestAccExplicitDatabricksCfg(t *testing.T) {
	workspaceTest(t)
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Profile: GetEnvOrSkipTest(t, "DATABRICKS_CONFIG_PROFILE"),
	}))
	ctx := context.Background()
	versions, err := w.Clusters.SparkVersions(ctx)
	assert.NoError(t, err)

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

	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		AzureResourceID: resourceID,
		Credentials:     config.AzureCliCredentials{},
	}))
	ctx := context.Background()
	versions, err := w.Clusters.SparkVersions(ctx)
	assert.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

func TestAccAzureErrorMappingForUnauthenticated(t *testing.T) {
	workspaceTest(t)
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		DebugHeaders:      true,
		AzureTenantID:     GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AzureClientID:     GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret: "invalid-for-integration-tests",
		AzureResourceID:   "/a/b/c",
	}))
	ctx := context.Background()
	_, err := w.Clusters.SparkVersions(ctx)
	assert.ErrorIs(t, err, apierr.ErrUnauthenticated)
}

func TestAccExplicitAzureSpnAuth(t *testing.T) {
	workspaceTest(t)
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		AzureTenantID:     GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AzureClientID:     GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret: GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		AzureResourceID:   GetEnvOrSkipTest(t, "DATABRICKS_AZURE_RESOURCE_ID"),
		Credentials:       config.AzureClientSecretCredentials{},
	}))
	ctx := context.Background()
	versions, err := w.Clusters.SparkVersions(ctx)
	assert.NoError(t, err)

	v, err := versions.Select(compute.SparkVersionRequest{
		LongTermSupport: true,
		Latest:          true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, v)
}

func TestAccErrNotAccountClient(t *testing.T) {
	workspaceTest(t)

	// Confirm that we get an error when trying to create an account client
	// if the configuration indicates a workspace.
	_, err := databricks.NewAccountClient()
	assert.ErrorIs(t, err, databricks.ErrNotAccountClient)
}

func TestAccErrNotWorkspaceClient(t *testing.T) {
	accountTest(t)

	// Confirm that we get an error when trying to create a workspace client
	// if the configuration indicates an account.
	_, err := databricks.NewWorkspaceClient()
	assert.ErrorIs(t, err, databricks.ErrNotWorkspaceClient)
}

// Confirm that we can access the account IP access lists using MSI credentials.
func TestAccAccountsMsiCredentials(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "ARM_USE_MSI"))
	ctx, a := accountTest(t)
	_, err := a.IpAccessLists.ListAll(ctx)
	assert.NoError(t, err)
}
