package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccUnifiedHostWorkspaceGroups(t *testing.T) {
	ctx, _, w := unifiedHostAccountTest(t)

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Groups.DeleteById(ctx, group.Id)
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	fetch, err := w.Groups.GetById(ctx, group.Id)
	require.NoError(t, err)
	assert.Equal(t, group.DisplayName, fetch.DisplayName)
}

func TestAccUnifiedHost(t *testing.T) {
	ctx, a, w := unifiedHostAccountTest(t)

	// Account level operations
	user, err := a.Users.Create(ctx, iam.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.Users.DeleteById(ctx, user.Id)
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	assert.Equal(t, 0, len(user.Roles))
	err = a.Users.Patch(ctx, iam.PartialUpdate{
		Id:      user.Id,
		Schemas: []iam.PatchSchema{iam.PatchSchemaUrnIetfParamsScimApiMessages20PatchOp},
		Operations: []iam.Patch{
			{Op: iam.PatchOpAdd, Value: iam.User{
				Roles: []iam.ComplexValue{
					{Value: "account_admin"},
				},
			}},
		},
	})
	require.NoError(t, err)

	byId, err := a.Users.GetById(ctx, user.Id)
	require.NoError(t, err)
	assert.Equal(t, user.Id, byId.Id)
	assert.Equal(t, 1, len(byId.Roles))

	// Workspace level operations
	_, err = w.CurrentUser.Me(ctx)
	require.NoError(t, err)
}

// SPOG/W — Workspace operations on unified host with explicit auth

func TestAccSpogWorkspacePAT(t *testing.T) {
	cfg := &config.Config{
		Host:               GetEnvOrSkipTest(t, "UNIFIED_HOST"),
		Token:              GetEnvOrSkipTest(t, "DATABRICKS_TOKEN"),
		AccountID:          GetEnvOrSkipTest(t, "TEST_ACCOUNT_ID"),
		WorkspaceID:        GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID"),
		AuthType:           "pat",
		HTTPTimeoutSeconds: 300,
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	t.Parallel()
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient((*databricks.Config)(cfg)))
	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, me.UserName)
}

func TestAccSpogWorkspaceOAuthM2M(t *testing.T) {
	cfg := &config.Config{
		Host:               GetEnvOrSkipTest(t, "UNIFIED_HOST"),
		ClientID:           GetEnvOrSkipTest(t, "DATABRICKS_CLIENT_ID"),
		ClientSecret:       GetEnvOrSkipTest(t, "DATABRICKS_CLIENT_SECRET"),
		AccountID:          GetEnvOrSkipTest(t, "TEST_ACCOUNT_ID"),
		WorkspaceID:        GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID"),
		AuthType:           "oauth-m2m",
		HTTPTimeoutSeconds: 300,
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	t.Parallel()
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient((*databricks.Config)(cfg)))
	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, me.UserName)
}

func TestAccSpogWorkspaceAzureClientSecret(t *testing.T) {
	cfg := &config.Config{
		Host:               GetEnvOrSkipTest(t, "UNIFIED_HOST"),
		AccountID:          GetEnvOrSkipTest(t, "TEST_ACCOUNT_ID"),
		WorkspaceID:        GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID"),
		AzureClientID:      GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret:  GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		AzureTenantID:      GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AuthType:           "azure-client-secret",
		HTTPTimeoutSeconds: 300,
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	t.Parallel()
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient((*databricks.Config)(cfg)))
	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, me.UserName)
}

func TestAccSpogWorkspaceGoogleCredentials(t *testing.T) {
	// google-credentials uses a GCP ID token with target_audience=cfg.host.
	// On the unified host this produces the same token for both account and
	// workspace requests. Account-level APIs accept this token, but
	// workspace-level APIs return 401.
	t.Skip("google-credentials ID token is rejected for workspace operations on unified hosts")
	cfg := &config.Config{
		Host:                 GetEnvOrSkipTest(t, "UNIFIED_HOST"),
		AccountID:            GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
		WorkspaceID:          GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID"),
		GoogleCredentials:    GetEnvOrSkipTest(t, "GOOGLE_CREDENTIALS"),
		GoogleServiceAccount: GetEnvOrSkipTest(t, "DATABRICKS_GOOGLE_SERVICE_ACCOUNT"),
		AuthType:             "google-credentials",
		HTTPTimeoutSeconds:   300,
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	t.Parallel()
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient((*databricks.Config)(cfg)))
	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, me.UserName)
}

// SPOG/A — Account operations on unified host with explicit auth

func TestAccSpogAccountOAuthM2M(t *testing.T) {
	cfg := &config.Config{
		Host:               GetEnvOrSkipTest(t, "UNIFIED_HOST"),
		AccountID:          GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
		ClientID:           GetEnvOrSkipTest(t, "DATABRICKS_CLIENT_ID"),
		ClientSecret:       GetEnvOrSkipTest(t, "DATABRICKS_CLIENT_SECRET"),
		AuthType:           "oauth-m2m",
		HTTPTimeoutSeconds: 300,
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	t.Parallel()
	ctx := context.Background()
	a := databricks.Must(databricks.NewAccountClient((*databricks.Config)(cfg)))
	it := a.ServicePrincipals.List(ctx, iam.ListAccountServicePrincipalsRequest{})
	sp, err := it.Next(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, sp.DisplayName)
}

func TestAccSpogAccountAzureClientSecret(t *testing.T) {
	cfg := &config.Config{
		Host:               GetEnvOrSkipTest(t, "UNIFIED_HOST"),
		AccountID:          GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
		AzureClientID:      GetEnvOrSkipTest(t, "ARM_CLIENT_ID"),
		AzureClientSecret:  GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET"),
		AzureTenantID:      GetEnvOrSkipTest(t, "ARM_TENANT_ID"),
		AuthType:           "azure-client-secret",
		HTTPTimeoutSeconds: 300,
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	t.Parallel()
	ctx := context.Background()
	a := databricks.Must(databricks.NewAccountClient((*databricks.Config)(cfg)))
	it := a.ServicePrincipals.List(ctx, iam.ListAccountServicePrincipalsRequest{})
	sp, err := it.Next(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, sp.DisplayName)
}

func TestAccSpogAccountGoogleCredentials(t *testing.T) {
	cfg := &config.Config{
		Host:                 GetEnvOrSkipTest(t, "UNIFIED_HOST"),
		AccountID:            GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
		GoogleCredentials:    GetEnvOrSkipTest(t, "GOOGLE_CREDENTIALS"),
		GoogleServiceAccount: GetEnvOrSkipTest(t, "DATABRICKS_GOOGLE_SERVICE_ACCOUNT"),
		AuthType:             "google-credentials",
		HTTPTimeoutSeconds:   300,
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	t.Parallel()
	ctx := context.Background()
	a := databricks.Must(databricks.NewAccountClient((*databricks.Config)(cfg)))
	it := a.ServicePrincipals.List(ctx, iam.ListAccountServicePrincipalsRequest{})
	sp, err := it.Next(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, sp.DisplayName)
}
