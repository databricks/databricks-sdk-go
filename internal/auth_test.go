package internal

import (
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/stretchr/testify/require"
)

func TestUcAccWifAuth(t *testing.T) {
	// This test cannot be run locally. It can only be run from GitHub Workflows.
	_ = GetEnvOrSkipTest(t, "ACTIONS_ID_TOKEN_REQUEST_URL")
	ctx, a := ucacctTest(t)

	// Create SP with access to the workspace
	sp, err := a.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: RandomName("go-sdk-sp-"),
		Roles: []iam.ComplexValue{
			{Value: "account_admin"}, // Assigning account-level admin role
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.ServicePrincipals.Delete(ctx, iam.DeleteAccountServicePrincipalRequest{Id: sp.Id})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	spId, err := strconv.ParseInt(sp.Id, 10, 64)
	require.NoError(t, err)

	// Setup Federation Policy
	p, err := a.ServicePrincipalFederationPolicy.Create(ctx, oauth2.CreateServicePrincipalFederationPolicyRequest{
		Policy: oauth2.FederationPolicy{
			OidcPolicy: &oauth2.OidcFederationPolicy{
				Issuer: "https://token.actions.githubusercontent.com",
				Audiences: []string{
					"https://github.com/databricks-eng",
				},
				Subject: "repo:databricks-eng/eng-dev-ecosystem:environment:integration-tests",
			},
		},
		ServicePrincipalId: spId,
	})

	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.ServicePrincipalFederationPolicy.Delete(ctx, oauth2.DeleteServicePrincipalFederationPolicyRequest{
			ServicePrincipalId: spId,
			PolicyId:           p.Uid,
		})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	// Test Workspace Identity Federation at Account Level

	accCfg := &databricks.Config{
		Host:          a.Config.Host,
		AccountID:     a.Config.AccountID,
		ClientID:      sp.ApplicationId,
		AuthType:      "github-oidc",
		TokenAudience: "https://github.com/databricks-eng",
	}

	wifAccClient, err := databricks.NewAccountClient(accCfg)

	require.NoError(t, err)
	it := wifAccClient.Groups.List(ctx, iam.ListAccountGroupsRequest{})
	_, err = it.Next(ctx)
	require.NoError(t, err)
}

func TestUcAccWifAuthWorkspace(t *testing.T) {
	// This test cannot be run locally. It can only be run from GitHub Workflows.
	_ = GetEnvOrSkipTest(t, "ACTIONS_ID_TOKEN_REQUEST_URL")
	ctx, a := ucacctTest(t)

	workspaceIdEnvVar := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	workspaceId, err := strconv.ParseInt(workspaceIdEnvVar, 10, 64)
	require.NoError(t, err)

	workspaceUrl := GetEnvOrSkipTest(t, "TEST_WORKSPACE_URL")

	// Create SP with access to the workspace
	sp, err := a.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: RandomName("go-sdk-sp-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.ServicePrincipals.Delete(ctx, iam.DeleteAccountServicePrincipalRequest{Id: sp.Id})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	spId, err := strconv.ParseInt(sp.Id, 10, 64)
	require.NoError(t, err)

	_, err = a.WorkspaceAssignment.Update(ctx, iam.UpdateWorkspaceAssignments{
		WorkspaceId: workspaceId,
		PrincipalId: spId,
		Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionAdmin},
	})

	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.WorkspaceAssignment.Delete(ctx, iam.DeleteWorkspaceAssignmentRequest{
			PrincipalId: spId,
			WorkspaceId: workspaceId,
		})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	// Setup Federation Policy
	p, err := a.ServicePrincipalFederationPolicy.Create(ctx, oauth2.CreateServicePrincipalFederationPolicyRequest{
		Policy: oauth2.FederationPolicy{
			OidcPolicy: &oauth2.OidcFederationPolicy{
				Issuer: "https://token.actions.githubusercontent.com",
				Audiences: []string{
					"https://github.com/databricks-eng",
				},
				Subject: "repo:databricks-eng/eng-dev-ecosystem:environment:integration-tests",
			},
		},
		ServicePrincipalId: spId,
	})

	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.ServicePrincipalFederationPolicy.Delete(ctx, oauth2.DeleteServicePrincipalFederationPolicyRequest{
			ServicePrincipalId: spId,
			PolicyId:           p.Uid,
		})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	wsCfg := &databricks.Config{
		Host:          workspaceUrl,
		ClientID:      sp.ApplicationId,
		AuthType:      "github-oidc",
		TokenAudience: "https://github.com/databricks-eng",
	}

	wifWsClient, err := databricks.NewWorkspaceClient(wsCfg)

	require.NoError(t, err)
	_, err = wifWsClient.CurrentUser.Me(ctx)
	require.NoError(t, err)
}

func TestUcAccWorkspaceOAuthM2MAuth(t *testing.T) {
	ctx, _ := ucwsTest(t)
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))

	// Get environment variables
	host := GetEnvOrSkipTest(t, "DATABRICKS_HOST")
	clientID := GetEnvOrSkipTest(t, "TEST_DATABRICKS_CLIENT_ID")
	clientSecret := GetEnvOrSkipTest(t, "TEST_DATABRICKS_CLIENT_SECRET")
	// Create workspace client with OAuth M2M authentication
	wsCfg := &databricks.Config{
		Host:         host,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthType:     "oauth-m2m",
	}

	wsClient, err := databricks.NewWorkspaceClient(wsCfg)
	if err != nil {
		t.Fatalf("failed to create workspace client: %v", err)
	}

	// Call the "me" API
	me, err := wsClient.CurrentUser.Me(ctx)
	if err != nil {
		t.Fatalf("failed to call CurrentUser.Me(): %v", err)
	}

	// Verify we got a valid response
	if me.UserName == "" {
		t.Errorf("expected non-empty UserName, got empty string")
	}
}

func TestUcAccWorkspaceAzureClientSecretAuth(t *testing.T) {
	ctx, _ := ucwsTest(t)
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))

	host := GetEnvOrSkipTest(t, "DATABRICKS_HOST")
	azureClientID := GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	azureClientSecret := GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET")
	azureTenantID := GetEnvOrSkipTest(t, "ARM_TENANT_ID")

	// Create workspace client with Azure client secret authentication
	wsCfg := &databricks.Config{
		Host:              host,
		AzureClientID:     azureClientID,
		AzureClientSecret: azureClientSecret,
		AzureTenantID:     azureTenantID,
		AuthType:          "azure-client-secret",
	}

	wsClient, err := databricks.NewWorkspaceClient(wsCfg)
	if err != nil {
		t.Fatalf("failed to create workspace client: %v", err)
	}

	// Call the "me" API
	me, err := wsClient.CurrentUser.Me(ctx)
	if err != nil {
		t.Fatalf("failed to call CurrentUser.Me(): %v", err)
	}

	// Verify we got a valid response
	if me.UserName == "" {
		t.Errorf("expected non-empty UserName, got empty string")
	}
}

func TestMwsAccAccountOAuthM2MAuth(t *testing.T) {
	ctx, _ := accountTest(t)
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))

	// Get environment variables
	host := GetEnvOrSkipTest(t, "DATABRICKS_HOST")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	clientID := GetEnvOrSkipTest(t, "TEST_DATABRICKS_CLIENT_ID")
	clientSecret := GetEnvOrSkipTest(t, "TEST_DATABRICKS_CLIENT_SECRET")

	// Create account client with OAuth M2M authentication
	accCfg := &databricks.Config{
		Host:         host,
		AccountID:    accountID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthType:     "oauth-m2m",
	}

	accClient, err := databricks.NewAccountClient(accCfg)
	if err != nil {
		t.Fatalf("failed to create account client: %v", err)
	}

	// List service principals to verify authentication works
	it := accClient.ServicePrincipals.List(ctx, iam.ListAccountServicePrincipalsRequest{})
	_, err = it.Next(ctx)
	if err != nil {
		t.Fatalf("failed to list service principals: %v", err)
	}
}

func TestMwsAccAccountAzureClientSecretAuth(t *testing.T) {
	ctx, _ := accountTest(t)
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))

	// Get environment variables
	host := GetEnvOrSkipTest(t, "DATABRICKS_HOST")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	azureClientID := GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	azureClientSecret := GetEnvOrSkipTest(t, "ARM_CLIENT_SECRET")
	azureTenantID := GetEnvOrSkipTest(t, "ARM_TENANT_ID")

	// Create account client with Azure client secret authentication
	accCfg := &databricks.Config{
		Host:              host,
		AccountID:         accountID,
		AzureClientID:     azureClientID,
		AzureClientSecret: azureClientSecret,
		AzureTenantID:     azureTenantID,
		AuthType:          "azure-client-secret",
	}

	accClient, err := databricks.NewAccountClient(accCfg)
	if err != nil {
		t.Fatalf("failed to create account client: %v", err)
	}

	// List service principals to verify authentication works
	it := accClient.ServicePrincipals.List(ctx, iam.ListAccountServicePrincipalsRequest{})
	_, err = it.Next(ctx)
	if err != nil {
		t.Fatalf("failed to list service principals: %v", err)
	}
}
