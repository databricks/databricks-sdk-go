package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/accounts"
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/permissions"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/stretchr/testify/require"
)

func TestMwsAccWorkspaceAssignment(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	a := accounts.New(&databricks.Config{
		AccountID: GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
	})
	if !a.Config.IsAccountsClient() || !a.Config.IsAws() {
		t.SkipNow()
	}
	spn, err := a.ServicePrincipals.CreateServicePrincipal(ctx, scim.ServicePrincipal{
		DisplayName: RandomName("sdk-go-"),
	})
	require.NoError(t, err)
	defer func() {
		err = a.ServicePrincipals.DeleteServicePrincipalById(ctx, spn.Id)
		require.NoError(t, err)
	}()

	workspaceId := GetEnvInt64OrSkipTest(t, "TEST_WORKSPACE_ID")
	_, err = a.WorkspaceAssignment.Create(ctx, permissions.CreateWorkspaceAssignments{
		WorkspaceId: workspaceId,
		PermissionAssignments: []permissions.PermissionAssignmentInput{
			{
				ServicePrincipalName: spn.ApplicationId,
				Permissions: []permissions.WorkspacePermission{
					permissions.WorkspacePermissionUser,
				},
			},
		},
	})
	require.NoError(t, err)
	defer func() {
		err = a.WorkspaceAssignment.DeleteByWorkspaceIdAndPrincipalId(ctx, workspaceId, spn.Id)
	}()

	all, err := a.WorkspaceAssignment.ListByWorkspaceId(ctx, workspaceId)
	require.NoError(t, err)

	var found bool
	for _, v := range all.PermissionAssignments {
		if v.Principal.PrincipalId == spn.Id {
			found = true
		}
	}
	require.True(t, found)
}
