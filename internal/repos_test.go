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
		Url:      "https://github.com/shreyas-goenka/empty-repo.git",
		Provider: "github",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = wsc.Repos.DeleteByRepoId(ctx, ri.Id)
		assert.NoError(t, err)
	})

	assert.Equal(t, "main", ri.Branch)
	err = wsc.Repos.Update(ctx, repos.UpdateRepo{
		RepoId: ri.Id,
		Branch: "foo",
	})
	require.NoError(t, err)

	_, err = wsc.Repos.List(ctx, repos.ListRequest{})
	require.NoError(t, err)
}
