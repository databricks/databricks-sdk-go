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

func TestAccWifAuth(t *testing.T) {
	ctx, a := ucacctTest(t)

	workspaceIdEnvVar := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	workspaceId, err := strconv.ParseInt(workspaceIdEnvVar, 10, 64)
	require.NoError(t, err)

	workspaceUrl := GetEnvOrSkipTest(t, "TEST_WORKSPACE_URL")

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
		Policy: &oauth2.FederationPolicy{
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

	// Test Workspace Identity Federation at Workspace Level

	wsCfg := &databricks.Config{
		Host:          workspaceUrl,
		ClientID:      sp.ApplicationId,
		AuthType:      "databricks-oidc",
		TokenAudience: "https://github.com/databricks-eng",
	}

	wifWsClient, err := databricks.NewWorkspaceClient(wsCfg)

	require.NoError(t, err)
	_, err = wifWsClient.CurrentUser.Me(ctx)
	require.NoError(t, err)

	// Test Workspace Identity Federation at Account Level

	accCfg := &databricks.Config{
		Host:          a.Config.Host,
		AccountID:     a.Config.AccountID,
		ClientID:      sp.ApplicationId,
		AuthType:      "databricks-oidc",
		TokenAudience: "https://github.com/databricks-eng",
	}

	wifAccClient, err := databricks.NewAccountClient(accCfg)

	require.NoError(t, err)
	it := wifAccClient.Groups.List(ctx, iam.ListAccountGroupsRequest{})
	_, err = it.Next(ctx)
	require.NoError(t, err)

}
