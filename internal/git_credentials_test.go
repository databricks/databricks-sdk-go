package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/qa/lock"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
)

func acquireGitCredentialsLock(ctx context.Context, t *testing.T, w *databricks.WorkspaceClient) {
	me, err := w.CurrentUser.Me(ctx)
	assert.NoError(t, err)
	lockable := lock.GitCredentials{
		WorkspaceHost: w.Config.Host,
		Username:      me.UserName,
	}
	_, err = lock.Acquire(ctx, lockable, lock.InTest(t))
	assert.NoError(t, err)
}

func TestAccGitCredentials(t *testing.T) {
	ctx, w := workspaceTest(t)

	// skip-next-line-roll
	acquireGitCredentialsLock(ctx, t, w)
	list, err := w.GitCredentials.ListAll(ctx)
	assert.NoError(t, err)
	for _, v := range list {
		err = w.GitCredentials.DeleteByCredentialId(ctx, v.CredentialId)
		assert.NoError(t, err)
	}

	cr, err := w.GitCredentials.Create(ctx, workspace.CreateCredentials{
		GitProvider:         "gitHub",
		GitUsername:         "test",
		PersonalAccessToken: "test",
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.GitCredentials.DeleteByCredentialId(ctx, cr.CredentialId)
		assert.NoError(t, err)
	})

	err = w.GitCredentials.Update(ctx, workspace.UpdateCredentials{
		CredentialId:        cr.CredentialId,
		GitProvider:         "gitHub",
		GitUsername:         RandomEmail(),
		PersonalAccessToken: RandomName(),
	})
	assert.NoError(t, err)

	byId, err := w.GitCredentials.GetByCredentialId(ctx, cr.CredentialId)
	assert.NoError(t, err)
	assert.NotEqual(t, cr.GitUsername, byId.GitUsername)

	byName, err := w.GitCredentials.GetByGitProvider(ctx, byId.GitProvider)
	assert.NoError(t, err)
	assert.Equal(t, byId.GitUsername, byName.GitUsername)

	names, err := w.GitCredentials.CredentialInfoGitProviderToCredentialIdMap(ctx)
	assert.NoError(t, err)
	assert.Contains(t, names, byId.GitProvider)
}
