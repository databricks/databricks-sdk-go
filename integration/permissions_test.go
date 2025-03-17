package integration

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/iam/v2"
	"github.com/databricks/databricks-sdk-go/workspace/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccGenericPermissions(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}
	WorkspaceAPI, err := workspace.NewWorkspaceClient(nil)
	require.NoError(t, err)
	notebookPath := myNotebookPath(t, cfg)

	err = WorkspaceAPI.Import(ctx, workspace.Import{
		Path:      notebookPath,
		Overwrite: true,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content:   base64.StdEncoding.EncodeToString([]byte(`print(1)`)),
	})
	require.NoError(t, err)

	obj, err := WorkspaceAPI.GetStatusByPath(ctx, notebookPath)
	require.NoError(t, err)

	PermissionsAPI, err := iam.NewPermissionsClient(nil)
	require.NoError(t, err)
	levels, err := PermissionsAPI.GetPermissionLevels(ctx, iam.GetPermissionLevelsRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
	})
	require.NoError(t, err)
	assert.True(t, len(levels.PermissionLevels) > 1)

	GroupsAPI, err := iam.NewGroupsClient(nil)
	require.NoError(t, err)
	group, err := GroupsAPI.Create(ctx, iam.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := GroupsAPI.DeleteById(ctx, group.Id)
		require.NoError(t, err)
	})

	_, err = PermissionsAPI.Set(ctx, iam.PermissionsRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
		AccessControlList: []iam.AccessControlRequest{
			{
				GroupName:       group.DisplayName,
				PermissionLevel: iam.PermissionLevelCanRun,
			},
		},
	})
	require.NoError(t, err)

	_, err = PermissionsAPI.Get(ctx, iam.GetPermissionRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
	})
	require.NoError(t, err)
}

func TestUcAccWorkspaceAssignmentOnAws(t *testing.T) {
	ctx, cfg := ucacctTest(t)
	if !cfg.IsAws() {
		t.SkipNow()
	}
	workspaceId := MustParseInt64(GetEnvOrSkipTest(t, "DUMMY_WORKSPACE_ID"))

	ServicePrincipalsAPI, err := iam.NewAccountServicePrincipalsClient(cfg)
	require.NoError(t, err)
	spn, err := ServicePrincipalsAPI.Create(ctx, iam.ServicePrincipal{
		DisplayName: RandomName("sdk-go-"),
	})
	require.NoError(t, err)
	defer func() {
		err = ServicePrincipalsAPI.DeleteById(ctx, spn.Id)
		require.NoError(t, err)
	}()

	spnId := MustParseInt64(spn.Id)

	WorkspaceAssignmentAPI, err := iam.NewWorkspaceAssignmentClient(cfg)
	require.NoError(t, err)
	_, err = WorkspaceAssignmentAPI.Update(ctx, iam.UpdateWorkspaceAssignments{
		WorkspaceId: workspaceId,
		PrincipalId: spnId,
		Permissions: []iam.WorkspacePermission{
			iam.WorkspacePermissionUser,
		},
	})
	require.NoError(t, err)
	defer func() {
		err = WorkspaceAssignmentAPI.DeleteByWorkspaceIdAndPrincipalId(ctx, workspaceId, spnId)
	}()

	all, err := WorkspaceAssignmentAPI.ListByWorkspaceId(ctx, workspaceId)
	require.NoError(t, err)

	var found bool
	for _, v := range all.PermissionAssignments {
		if v.Principal.PrincipalId == spnId {
			found = true
		}
	}
	require.True(t, found)
}
