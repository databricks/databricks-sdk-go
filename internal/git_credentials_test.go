package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/gitcredentials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccGitCredentials(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	list, err := w.GitCredentials.ListAll(ctx)
	require.NoError(t, err)
	for _, v := range list {
		err = w.GitCredentials.DeleteByCredentialId(ctx, v.CredentialId)
		require.NoError(t, err)
	}

	cr, err := w.GitCredentials.Create(ctx, gitcredentials.CreateCredentials{
		GitProvider:         "gitHub",
		GitUsername:         "test",
		PersonalAccessToken: "test",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.GitCredentials.DeleteByCredentialId(ctx, cr.CredentialId)
		require.NoError(t, err)
	})

	err = w.GitCredentials.Update(ctx, gitcredentials.UpdateCredentials{
		CredentialId:        cr.CredentialId,
		GitProvider:         "gitHub",
		GitUsername:         RandomEmail(),
		PersonalAccessToken: RandomName(),
	})
	require.NoError(t, err)

	load, err := w.GitCredentials.GetByCredentialId(ctx, cr.CredentialId)
	require.NoError(t, err)

	assert.NotEqual(t, cr.GitUsername, load.GitUsername)
}
