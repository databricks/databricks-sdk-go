package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/databricks/qa/lock"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func acquireGitCredentialsLock(ctx context.Context, t *testing.T, w *databricks.WorkspaceClient) {
	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)
	lockable := lock.GitCredentials{
		WorkspaceHost: w.Config.Host,
		Username:      me.UserName,
	}
	_, err = lock.Acquire(ctx, lockable, lock.InTest(t))
	require.NoError(t, err)
}

func TestAccGitCredentials(t *testing.T) {
	ctx, w := workspaceTest(t)

	// skip-next-line-roll
	acquireGitCredentialsLock(ctx, t, w)
	list, err := w.GitCredentials.ListAll(ctx)
	require.NoError(t, err)
	for _, v := range list {
		err = w.GitCredentials.DeleteByCredentialId(ctx, v.CredentialId)
		require.NoError(t, err)
	}

	cr, err := w.GitCredentials.Create(ctx, workspace.CreateCredentialsRequest{
		GitProvider:         "gitHub",
		GitUsername:         "test",
		PersonalAccessToken: "test",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.GitCredentials.DeleteByCredentialId(ctx, cr.CredentialId)
		require.NoError(t, err)
	})

	err = w.GitCredentials.Update(ctx, workspace.UpdateCredentialsRequest{
		CredentialId:        cr.CredentialId,
		GitProvider:         "gitHub",
		GitUsername:         RandomEmail(),
		PersonalAccessToken: RandomName(),
	})
	require.NoError(t, err)

	byId, err := w.GitCredentials.GetByCredentialId(ctx, cr.CredentialId)
	require.NoError(t, err)
	assert.NotEqual(t, cr.GitUsername, byId.GitUsername)

	byName, err := w.GitCredentials.GetByGitProvider(ctx, byId.GitProvider)
	require.NoError(t, err)
	assert.Equal(t, byId.GitUsername, byName.GitUsername)

	names, err := w.GitCredentials.CredentialInfoGitProviderToCredentialIdMap(ctx)
	require.NoError(t, err)
	assert.Contains(t, names, byId.GitProvider)
}
