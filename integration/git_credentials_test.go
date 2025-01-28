package integration

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/qa/lock"
	"github.com/databricks/databricks-sdk-go/iam/v2"
	"github.com/databricks/databricks-sdk-go/workspace/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func acquireGitCredentialsLock(ctx context.Context, t *testing.T, CurrentUserAPI *iam.CurrentUserClient) {
	me, err := CurrentUserAPI.Me(ctx)
	require.NoError(t, err)
	lockable := lock.GitCredentials{
		WorkspaceHost: CurrentUserAPI.Config.Host,
		Username:      me.UserName,
	}
	_, err = lock.Acquire(ctx, lockable, lock.InTest(t))
	require.NoError(t, err)
}

func TestAccGitCredentials(t *testing.T) {
	ctx := workspaceTest(t)

	CurrentUserAPI, err := iam.NewCurrentUserClient(nil)
	require.NoError(t, err)
	// skip-next-line-roll
	acquireGitCredentialsLock(ctx, t, CurrentUserAPI)

	GitCredentialsAPI, err := workspace.NewGitCredentialsClient(nil)
	list, err := GitCredentialsAPI.ListAll(ctx)
	require.NoError(t, err)
	for _, v := range list {
		err = GitCredentialsAPI.DeleteByCredentialId(ctx, v.CredentialId)
		require.NoError(t, err)
	}

	cr, err := GitCredentialsAPI.Create(ctx, workspace.CreateCredentialsRequest{
		GitProvider:         "gitHub",
		GitUsername:         "test",
		PersonalAccessToken: "test",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = GitCredentialsAPI.DeleteByCredentialId(ctx, cr.CredentialId)
		require.NoError(t, err)
	})

	err = GitCredentialsAPI.Update(ctx, workspace.UpdateCredentialsRequest{
		CredentialId:        cr.CredentialId,
		GitProvider:         "gitHub",
		GitUsername:         RandomEmail(),
		PersonalAccessToken: RandomName(),
	})
	require.NoError(t, err)

	byId, err := GitCredentialsAPI.GetByCredentialId(ctx, cr.CredentialId)
	require.NoError(t, err)
	assert.NotEqual(t, cr.GitUsername, byId.GitUsername)

	byName, err := GitCredentialsAPI.GetByGitProvider(ctx, byId.GitProvider)
	require.NoError(t, err)
	assert.Equal(t, byId.GitUsername, byName.GitUsername)

	names, err := GitCredentialsAPI.CredentialInfoGitProviderToCredentialIdMap(ctx)
	require.NoError(t, err)
	assert.Contains(t, names, byId.GitProvider)
}
