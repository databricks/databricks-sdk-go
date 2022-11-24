package internal

import (
	"context"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/permissions"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/stretchr/testify/require"
)

func TestMwsAccWorkspaceAssignment(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	a := databricks.Must(databricks.NewAccountClient(&databricks.Config{
		AccountID: GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
	}))
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

	spnId, err := strconv.ParseInt(spn.Id, 10, 64)
	require.NoError(t, err)

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
		err = a.WorkspaceAssignment.DeleteByWorkspaceIdAndPrincipalId(ctx, workspaceId, spnId)
	}()

	all, err := a.WorkspaceAssignment.ListByWorkspaceId(ctx, workspaceId)
	require.NoError(t, err)

	var found bool
	for _, v := range all.PermissionAssignments {
		if v.Principal.PrincipalId == spnId {
			found = true
		}
	}
	require.True(t, found)
}
