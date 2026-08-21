package internal

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/databricks-sdk-go/service/workspace"
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
	_, err = wifWsClient.CurrentUser.Me(ctx, iam.MeRequest{})
	require.NoError(t, err)
}

// TestUcAccWifAuthWorkspaceGroupRole creates a WIF service principal, an
// assumable group role, and a notebook that only the role can read. It verifies
// that role access succeeds while normal service principal access is denied.
func TestUcAccWifAuthWorkspaceGroupRole(t *testing.T) {
	// Use the GitHub Actions OIDC environment and an account administrator to arrange the test.
	_ = GetEnvOrSkipTest(t, "ACTIONS_ID_TOKEN_REQUEST_URL")
	ctx, a := ucacctTest(t)

	workspaceID := MustParseInt64(GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID"))
	workspaceURL := GetEnvOrSkipTest(t, "TEST_WORKSPACE_URL")

	// Create an administrator client for workspace resources and permissions.
	workspaceAdmin, err := databricks.NewWorkspaceClient(&databricks.Config{
		Host: workspaceURL,
	})
	require.NoError(t, err)

	// Create the service principal whose normal and role-based WIF access will be tested.
	sp, err := a.ServicePrincipalsV2.Create(ctx, iam.CreateAccountServicePrincipalRequest{
		Active:      true,
		DisplayName: RandomName("go-sdk-wif-role-sp-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupErr := a.ServicePrincipalsV2.Delete(ctx, iam.DeleteAccountServicePrincipalRequest{Id: sp.Id})
		require.True(t, cleanupErr == nil || apierr.IsMissing(cleanupErr))
	})

	spID, err := strconv.ParseInt(sp.Id, 10, 64)
	require.NoError(t, err)

	// Give the service principal basic workspace access without notebook access.
	_, err = a.WorkspaceAssignment.Update(ctx, iam.UpdateWorkspaceAssignments{
		WorkspaceId: workspaceID,
		PrincipalId: spID,
		Permissions: []iam.WorkspacePermission{
			iam.WorkspacePermissionUser,
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupErr := a.WorkspaceAssignment.Delete(ctx, iam.DeleteWorkspaceAssignmentRequest{
			PrincipalId: spID,
			WorkspaceId: workspaceID,
		})
		require.True(t, cleanupErr == nil || apierr.IsMissing(cleanupErr))
	})

	// Create the group that represents the temporary workspace role.
	group, err := a.GroupsV2.Create(ctx, iam.CreateAccountGroupRequest{
		DisplayName: RandomName("go-sdk-wif-role-group-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupErr := a.GroupsV2.Delete(ctx, iam.DeleteAccountGroupRequest{Id: group.Id})
		require.True(t, cleanupErr == nil || apierr.IsMissing(cleanupErr))
	})

	groupID, err := strconv.ParseInt(group.Id, 10, 64)
	require.NoError(t, err)

	// Assign the group to the workspace so it can receive workspace permissions.
	_, err = a.WorkspaceAssignment.Update(ctx, iam.UpdateWorkspaceAssignments{
		WorkspaceId: workspaceID,
		PrincipalId: groupID,
		Permissions: []iam.WorkspacePermission{
			iam.WorkspacePermissionUser,
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupErr := a.WorkspaceAssignment.Delete(ctx, iam.DeleteWorkspaceAssignmentRequest{
			PrincipalId: groupID,
			WorkspaceId: workspaceID,
		})
		require.True(t, cleanupErr == nil || apierr.IsMissing(cleanupErr))
	})

	// Grant the service principal permission to assume the group role.
	ruleSetName := fmt.Sprintf("accounts/%s/groups/%s/ruleSets/default", a.Config.AccountID, group.Id)
	ruleSet, err := a.AccessControl.GetRuleSet(ctx, iam.GetRuleSetRequest{
		Name: ruleSetName,
	})
	require.NoError(t, err)

	grantRules := append([]iam.GrantRule(nil), ruleSet.GrantRules...)
	grantRules = append(grantRules, iam.GrantRule{
		Principals: []string{
			"servicePrincipals/" + sp.ApplicationId,
		},
		Role: "roles/group.assumer",
	})

	_, err = a.AccessControl.UpdateRuleSet(ctx, iam.UpdateRuleSetRequest{
		Name: ruleSetName,
		RuleSet: iam.RuleSetUpdateRequest{
			Etag:       ruleSet.Etag,
			GrantRules: grantRules,
			Name:       ruleSetName,
		},
	})
	require.NoError(t, err)

	// Trust this repository's GitHub OIDC identity to authenticate as the service principal.
	policy, err := a.ServicePrincipalFederationPolicy.Create(ctx, oauth2.CreateServicePrincipalFederationPolicyRequest{
		Policy: oauth2.FederationPolicy{
			OidcPolicy: &oauth2.OidcFederationPolicy{
				Issuer: "https://token.actions.githubusercontent.com",
				Audiences: []string{
					"https://github.com/databricks-eng",
				},
				Subject: "repo:databricks-eng/eng-dev-ecosystem:environment:integration-tests",
			},
		},
		ServicePrincipalId: spID,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupErr := a.ServicePrincipalFederationPolicy.Delete(ctx, oauth2.DeleteServicePrincipalFederationPolicyRequest{
			ServicePrincipalId: spID,
			PolicyId:           policy.Uid,
		})
		require.True(t, cleanupErr == nil || apierr.IsMissing(cleanupErr))
	})

	// Create a temporary notebook that will distinguish normal access from role access.
	notebookPath := myNotebookPath(t, workspaceAdmin)
	err = workspaceAdmin.Workspace.Import(ctx, workspace.Import{
		Path:      notebookPath,
		Overwrite: true,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content:   base64.StdEncoding.EncodeToString([]byte("print(1)")),
	})
	require.NoError(t, err)

	notebook, err := workspaceAdmin.Workspace.GetStatusByPath(ctx, notebookPath)
	require.NoError(t, err)

	// Give only the group role permission to read the notebook.
	_, err = workspaceAdmin.Permissions.Update(ctx, iam.UpdateObjectPermissions{
		RequestObjectType: "notebooks",
		RequestObjectId:   strconv.FormatInt(notebook.ObjectId, 10),
		AccessControlList: []iam.AccessControlRequest{
			{
				GroupName:       group.DisplayName,
				PermissionLevel: iam.PermissionLevelCanRead,
			},
		},
	})
	require.NoError(t, err)

	// Authenticate with the group role and verify that its notebook permission is usable.
	roleClient, err := databricks.NewWorkspaceClient(&databricks.Config{
		Host:          workspaceURL,
		ClientID:      sp.ApplicationId,
		GroupID:       group.Id,
		AuthType:      "github-oidc",
		TokenAudience: "https://github.com/databricks-eng",
	})
	require.NoError(t, err)

	_, err = roleClient.Workspace.GetStatusByPath(ctx, notebookPath)
	require.NoError(t, err)

	// Authenticate normally as the same service principal and verify that access is denied.
	normalClient, err := databricks.NewWorkspaceClient(&databricks.Config{
		Host:          workspaceURL,
		ClientID:      sp.ApplicationId,
		AuthType:      "github-oidc",
		TokenAudience: "https://github.com/databricks-eng",
	})
	require.NoError(t, err)

	_, err = normalClient.Workspace.GetStatusByPath(ctx, notebookPath)
	if err == nil {
		t.Fatal("normal WIF credentials accessed a notebook whose permission belongs only to the group role")
	}
	if !errors.Is(err, apierr.ErrPermissionDenied) && !apierr.IsMissing(err) {
		t.Fatalf("normal WIF credentials failed with %v, want a notebook permission denial", err)
	}
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
	me, err := wsClient.CurrentUser.Me(ctx, iam.MeRequest{})
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
	me, err := wsClient.CurrentUser.Me(ctx, iam.MeRequest{})
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
