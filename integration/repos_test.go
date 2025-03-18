package integration

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/settings/v2"
	"github.com/databricks/databricks-sdk-go/workspace/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccRepos(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}

	// Skip this test if "Files in Repos" is not enabled.
	WorkspaceConfAPI, err := settings.NewWorkspaceConfClient(cfg)
	require.NoError(t, err)
	conf, err := WorkspaceConfAPI.GetStatus(ctx, settings.GetStatusRequest{
		Keys: "enableWorkspaceFilesystem",
	})
	require.NoError(t, err)
	if (*conf)["enableWorkspaceFilesystem"] == "false" {
		t.Skipf("Files in Repos not enabled in this workspace")
	}

	// Synthesize unique path for this checkout in this user's home.
	root := RandomName(fmt.Sprintf("/Repos/%s/tf-", me(t, cfg).UserName))
	ReposAPI, err := workspace.NewReposClient(nil)
	require.NoError(t, err)
	ri, err := ReposAPI.Create(ctx, workspace.CreateRepoRequest{
		Path:     root,
		Url:      "https://github.com/shreyas-goenka/empty-repo.git",
		Provider: "github",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err = ReposAPI.DeleteByRepoId(ctx, ri.Id)
		require.NoError(t, err)
	})

	assert.Equal(t, "main", ri.Branch)
	_, err = ReposAPI.Update(ctx, workspace.UpdateRepoRequest{
		RepoId: ri.Id,
		Branch: "foo",
	})
	require.NoError(t, err)

	byId, err := ReposAPI.GetByRepoId(ctx, ri.Id)
	require.NoError(t, err)

	byName, err := ReposAPI.GetByPath(ctx, byId.Path)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := ReposAPI.ListAll(ctx, workspace.ListReposRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	paths, err := ReposAPI.RepoInfoPathToIdMap(ctx, workspace.ListReposRequest{
		PathPrefix: "/",
	})
	require.NoError(t, err)
	assert.Equal(t, len(paths), len(all))
}
