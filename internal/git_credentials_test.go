package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccGitCredentials(t *testing.T) {
	ctx, w := workspaceTest(t)

	cr, err := w.GitCredentials.Create(ctx, workspace.CreateCredentials{
		GitProvider:         "gitHub",
		GitUsername:         "test",
		PersonalAccessToken: "test",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.GitCredentials.DeleteByCredentialId(ctx, cr.CredentialId)
		require.NoError(t, err)
	})

	list, err := w.GitCredentials.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(list) >= 1)

	err = w.GitCredentials.Update(ctx, workspace.UpdateCredentials{
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
