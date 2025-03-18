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
	//testWorkspaceId := int64(470576644108500)
	//testWorkspaceUrl := "https://dbc-1232e87d-9384.cloud.databricks.com"

	// Create SP
	sp, err := a.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: RandomName("go-sdk-sp-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.ServicePrincipals.Delete(ctx, iam.DeleteAccountServicePrincipalRequest{Id: sp.Id})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	applicationId, err := strconv.ParseInt(sp.Id, 10, 64)
	require.NoError(t, err)

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
		ServicePrincipalId: applicationId,
	})

	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.ServicePrincipalFederationPolicy.Delete(ctx, oauth2.DeleteServicePrincipalFederationPolicyRequest{
			ServicePrincipalId: applicationId,
			PolicyId:           p.Uid,
		})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	// _, err = a.WorkspaceAssignment.Update(ctx, iam.UpdateWorkspaceAssignments{
	// 	WorkspaceId: testWorkspaceId,
	// 	PrincipalId: applicationId,
	// 	Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionAdmin},
	// })

	// require.NoError(t, err)
	// t.Cleanup(func() {
	// 	err := a.WorkspaceAssignment.Delete(ctx, iam.DeleteWorkspaceAssignmentRequest{
	// 		PrincipalId: applicationId,
	// 		WorkspaceId: testWorkspaceId,
	// 	})
	// 	require.True(t, err == nil || apierr.IsMissing(err))
	// })

	cfg := &databricks.Config{
		//Host:     testWorkspaceUrl,
		Host:     a.Config.Host,
		ClientID: sp.Id,
		AuthType: "databricks-wif",
		//Host:      testWorkspaceUrl,
	}

	ws, err := databricks.NewAccountClient(cfg)

	require.NoError(t, err)
	users := ws.Users.List(ctx, iam.ListAccountUsersRequest{})
	_, err = users.Next(ctx)
	require.NoError(t, err)
	// 	testWorkspaceId := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	// testWorkspaceUrl := GetEnvOrSkipTest(t, "TEST_WORKSPACE_HOST")

}
