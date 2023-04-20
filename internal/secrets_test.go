package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSecrets(t *testing.T) {
	ctx, w := workspaceTest(t)

	scope := workspace.CreateScope{
		Scope: RandomEmail(),
	}
	err := w.Secrets.CreateScope(ctx, scope)
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Secrets.DeleteScopeByScope(ctx, scope.Scope)
		require.NoError(t, err)
	})
	scopes, err := w.Secrets.ListScopesAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(scopes) >= 1)

	put := workspace.PutSecret{
		Scope:       scope.Scope,
		Key:         RandomName("sdk-go"),
		StringValue: RandomName("dummy"),
	}
	err = w.Secrets.PutSecret(ctx, put)
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Secrets.DeleteSecret(ctx, workspace.DeleteSecret{
			Scope: scope.Scope,
			Key:   put.Key,
		})
		require.NoError(t, err)
	})

	scrts, err := w.Secrets.ListSecretsByScope(ctx, scope.Scope)
	require.NoError(t, err)
	assert.True(t, len(scrts.Secrets) == 1)

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Groups.DeleteById(ctx, group.Id)
		require.NoError(t, err)
	})

	err = w.Secrets.PutAcl(ctx, workspace.PutAcl{
		Scope:      scope.Scope,
		Permission: workspace.AclPermissionManage,
		Principal:  group.DisplayName,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		w.Secrets.DeleteAcl(ctx, workspace.DeleteAcl{
			Scope:     scope.Scope,
			Principal: group.DisplayName,
		})
	})

	acls, err := w.Secrets.ListAclsByScope(ctx, scope.Scope)
	require.NoError(t, err)

	assert.Equal(t, 2, len(acls.Items))
}
