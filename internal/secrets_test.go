package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSecrets(t *testing.T) {
	ctx, w := qa.WorkspaceTest(t)

	scopeName := qa.RandomName("scope-")
	keyName := qa.RandomName("key-")

	// creates scopeName
	err := w.Secrets.CreateScope(ctx, workspace.CreateScope{
		Scope: scopeName,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Secrets.DeleteScopeByScope(ctx, scopeName)
		require.NoError(t, err)
	})

	scopes, err := w.Secrets.ListScopesAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(scopes) >= 1)

	// creates keyName
	err = w.Secrets.PutSecret(ctx, workspace.PutSecret{
		Scope:       scopeName,
		Key:         keyName,
		StringValue: qa.RandomName("dummy"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Secrets.DeleteSecret(ctx, workspace.DeleteSecret{
			Scope: scopeName,
			Key:   keyName,
		})
		require.NoError(t, err)
	})

	scrts, err := w.Secrets.ListSecretsByScope(ctx, scopeName)
	require.NoError(t, err)
	assert.True(t, len(scrts.Secrets) == 1)

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: qa.RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Groups.DeleteById(ctx, group.Id)
		require.NoError(t, err)
	})

	err = w.Secrets.PutAcl(ctx, workspace.PutAcl{
		Scope:      scopeName,
		Permission: workspace.AclPermissionManage,
		Principal:  group.DisplayName,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		w.Secrets.DeleteAcl(ctx, workspace.DeleteAcl{
			Scope:     scopeName,
			Principal: group.DisplayName,
		})
	})

	acls, err := w.Secrets.ListAclsByScope(ctx, scopeName)
	require.NoError(t, err)

	assert.Equal(t, 2, len(acls.Items))
}
