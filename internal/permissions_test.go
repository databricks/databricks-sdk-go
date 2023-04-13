package internal

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/permissions"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccGenericPermissions(t *testing.T) {
	ctx, w := workspaceTest(t)
	notebookPath := myNotebookPath(t, w)

	err := w.Workspace.Import(ctx, workspace.PythonNotebookOverwrite(
		notebookPath, `print(1)`))
	require.NoError(t, err)

	obj, err := w.Workspace.GetStatusByPath(ctx, notebookPath)
	require.NoError(t, err)

	levels, err := w.Permissions.GetPermissionLevels(ctx, permissions.GetPermissionLevels{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
	})
	require.NoError(t, err)
	assert.True(t, len(levels.PermissionLevels) > 1)

	group, err := w.Groups.Create(ctx, scim.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Groups.DeleteById(ctx, group.Id)
		require.NoError(t, err)
	})

	err = w.Permissions.Set(ctx, permissions.PermissionsRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
		AccessControlList: []permissions.AccessControlRequest{
			{
				GroupName:       group.DisplayName,
				PermissionLevel: permissions.PermissionLevelCanRun,
			},
		},
	})
	require.NoError(t, err)

	_, err = w.Permissions.Get(ctx, permissions.Get{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
	})
	require.NoError(t, err)
}

func TestMwsAccWorkspaceAssignment(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}
	spn, err := a.ServicePrincipals.Create(ctx, scim.ServicePrincipal{
		DisplayName: RandomName("sdk-go-"),
	})
	require.NoError(t, err)
	defer func() {
		err = a.ServicePrincipals.DeleteById(ctx, spn.Id)
		require.NoError(t, err)
	}()

	spnId, err := strconv.ParseInt(spn.Id, 10, 64)
	require.NoError(t, err)

	workspaceId := GetEnvInt64OrSkipTest(t, "TEST_WORKSPACE_ID")
	err = a.WorkspaceAssignment.Update(ctx, permissions.UpdateWorkspaceAssignments{
		WorkspaceId: workspaceId,
		PrincipalId: spnId,
		Permissions: []permissions.WorkspacePermission{
			permissions.WorkspacePermissionUser,
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
