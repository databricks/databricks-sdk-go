package internal

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
)

func TestAccRepos(t *testing.T) {
	ctx, w := workspaceTest(t)

	// Skip this test if "Files in Repos" is not enabled.
	conf, err := w.WorkspaceConf.GetStatus(ctx, settings.GetStatusRequest{
		Keys: "enableWorkspaceFilesystem",
	})
	assert.NoError(t, err)
	if (*conf)["enableWorkspaceFilesystem"] == "false" {
		t.Skipf("Files in Repos not enabled in this workspace")
	}

	// Synthesize unique path for this checkout in this user's home.
	root := RandomName(fmt.Sprintf("/Repos/%s/tf-", me(t, w).UserName))
	ri, err := w.Repos.Create(ctx, workspace.CreateRepo{
		Path:     root,
		Url:      "https://github.com/shreyas-goenka/empty-repo.git",
		Provider: "github",
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Repos.DeleteByRepoId(ctx, ri.Id)
		assert.NoError(t, err)
	})

	assert.Equal(t, "main", ri.Branch)
	err = w.Repos.Update(ctx, workspace.UpdateRepo{
		RepoId: ri.Id,
		Branch: "foo",
	})
	assert.NoError(t, err)

	byId, err := w.Repos.GetByRepoId(ctx, ri.Id)
	assert.NoError(t, err)

	byName, err := w.Repos.GetByPath(ctx, byId.Path)
	assert.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := w.Repos.ListAll(ctx, workspace.ListReposRequest{})
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)

	paths, err := w.Repos.RepoInfoPathToIdMap(ctx, workspace.ListReposRequest{
		PathPrefix: "/",
	})
	assert.NoError(t, err)
	assert.Equal(t, len(paths), len(all))
}
