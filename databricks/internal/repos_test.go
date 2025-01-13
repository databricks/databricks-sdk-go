package internal

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccRepos(t *testing.T) {
	ctx, w := workspaceTest(t)

	// Skip this test if "Files in Repos" is not enabled.
	conf, err := w.WorkspaceConf.GetStatus(ctx, settings.GetStatusRequest{
		Keys: "enableWorkspaceFilesystem",
	})
	require.NoError(t, err)
	if (*conf)["enableWorkspaceFilesystem"] == "false" {
		t.Skipf("Files in Repos not enabled in this workspace")
	}

	// Synthesize unique path for this checkout in this user's home.
	root := RandomName(fmt.Sprintf("/Repos/%s/tf-", me(t, w).UserName))
	ri, err := w.Repos.Create(ctx, workspace.CreateRepoRequest{
		Path:     root,
		Url:      "https://github.com/shreyas-goenka/empty-repo.git",
		Provider: "github",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Repos.DeleteByRepoId(ctx, ri.Id)
		require.NoError(t, err)
	})

	assert.Equal(t, "main", ri.Branch)
	err = w.Repos.Update(ctx, workspace.UpdateRepoRequest{
		RepoId: ri.Id,
		Branch: "foo",
	})
	require.NoError(t, err)

	byId, err := w.Repos.GetByRepoId(ctx, ri.Id)
	require.NoError(t, err)

	byName, err := w.Repos.GetByPath(ctx, byId.Path)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := w.Repos.ListAll(ctx, workspace.ListReposRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	paths, err := w.Repos.RepoInfoPathToIdMap(ctx, workspace.ListReposRequest{
		PathPrefix: "/",
	})
	require.NoError(t, err)
	assert.Equal(t, len(paths), len(all))
}
