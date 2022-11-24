package internal

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/repos"
	"github.com/databricks/databricks-sdk-go/service/workspaceconf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccRepos(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	// Skip this test if "Files in Repos" is not enabled.
	conf, err := w.WorkspaceConf.GetStatus(ctx, workspaceconf.GetStatusRequest{
		Keys: "enableWorkspaceFilesystem",
	})
	require.NoError(t, err)
	if (*conf)["enableWorkspaceFilesystem"] == "false" {
		t.Skipf("Files in Repos not enabled in this workspace")
	}

	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)

	// Synthesize unique path for this checkout in this user's home.
	root := RandomName(fmt.Sprintf("/Repos/%s/tf-", me.UserName))
	ri, err := w.Repos.Create(ctx, repos.CreateRepo{
		Path:     root,
		Url:      "https://github.com/shreyas-goenka/empty-repo.git",
		Provider: "github",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Repos.DeleteByRepoId(ctx, ri.Id)
		assert.NoError(t, err)
	})

	assert.Equal(t, "main", ri.Branch)
	err = w.Repos.Update(ctx, repos.UpdateRepo{
		RepoId: ri.Id,
		Branch: "foo",
	})
	require.NoError(t, err)

	all, err := w.Repos.ListAll(ctx, repos.ListRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}
