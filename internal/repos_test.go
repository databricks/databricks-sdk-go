package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/repos"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccRepos(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	wsc := workspaces.New()

	root := "/Repos/acc-test"
	err := wsc.Workspace.MkdirsByPath(ctx, root)
	require.NoError(t, err)

	ri, err := wsc.Repos.Create(ctx, repos.CreateRepo{
		Url:      "https://github.com/databricks/notebook-best-practices.git",
		Provider: "github",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = wsc.Repos.DeleteByRepoId(ctx, ri.Id)
		assert.NoError(t, err)
	})

	assert.Equal(t, "dais-2022", ri.Branch)
	err = wsc.Repos.Update(ctx, repos.UpdateRepo{
		RepoId: ri.Id,
		Branch: "main",
	})
	require.NoError(t, err)

	all, err := wsc.Repos.ListAll(ctx, repos.ListRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) > 1)
}
