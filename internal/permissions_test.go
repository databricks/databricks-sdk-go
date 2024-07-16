package internal

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
)

func TestAccGenericPermissions(t *testing.T) {
	ctx, w := workspaceTest(t)
	notebookPath := myNotebookPath(t, w)

	err := w.Workspace.Import(ctx, workspace.Import{
		Path:      notebookPath,
		Overwrite: true,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content:   base64.StdEncoding.EncodeToString([]byte(`print(1)`)),
	})
	assert.NoError(t, err)

	obj, err := w.Workspace.GetStatusByPath(ctx, notebookPath)
	assert.NoError(t, err)

	levels, err := w.Permissions.GetPermissionLevels(ctx, iam.GetPermissionLevelsRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
	})
	assert.NoError(t, err)
	assert.True(t, len(levels.PermissionLevels) > 1)

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := w.Groups.DeleteById(ctx, group.Id)
		assert.NoError(t, err)
	})

	_, err = w.Permissions.Set(ctx, iam.PermissionsRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
		AccessControlList: []iam.AccessControlRequest{
			{
				GroupName:       group.DisplayName,
				PermissionLevel: iam.PermissionLevelCanRun,
			},
		},
	})
	assert.NoError(t, err)

	_, err = w.Permissions.Get(ctx, iam.GetPermissionRequest{
		RequestObjectType: "notebooks",
		RequestObjectId:   fmt.Sprintf("%d", obj.ObjectId),
	})
	assert.NoError(t, err)
}

func TestUcAccWorkspaceAssignmentOnAws(t *testing.T) {
	ctx, a := ucacctTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}
	workspaceId := MustParseInt64(GetEnvOrSkipTest(t, "DUMMY_WORKSPACE_ID"))

	spn, err := a.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: RandomName("sdk-go-"),
	})
	assert.NoError(t, err)
	defer func() {
		err = a.ServicePrincipals.DeleteById(ctx, spn.Id)
		assert.NoError(t, err)
	}()

	spnId := MustParseInt64(spn.Id)

	_, err = a.WorkspaceAssignment.Update(ctx, iam.UpdateWorkspaceAssignments{
		WorkspaceId: workspaceId,
		PrincipalId: spnId,
		Permissions: []iam.WorkspacePermission{
			iam.WorkspacePermissionUser,
		},
	})
	assert.NoError(t, err)
	defer func() {
		err = a.WorkspaceAssignment.DeleteByWorkspaceIdAndPrincipalId(ctx, workspaceId, spnId)
	}()

	all, err := a.WorkspaceAssignment.ListByWorkspaceId(ctx, workspaceId)
	assert.NoError(t, err)

	var found bool
	for _, v := range all.PermissionAssignments {
		if v.Principal.PrincipalId == spnId {
			found = true
		}
	}
	assert.True(t, found)
}
