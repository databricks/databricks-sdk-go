package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/databricks/databricks-sdk-go/service/secrets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSecrets(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountClient() {
		t.SkipNow()
	}

	scope := secrets.CreateScope{
		Scope: RandomEmail("sdk-go"),
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

	put := secrets.PutSecret{
		Scope:       scope.Scope,
		Key:         RandomName("sdk-go"),
		StringValue: RandomName("dummy"),
	}
	err = w.Secrets.PutSecret(ctx, put)
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Secrets.DeleteSecret(ctx, secrets.DeleteSecret{
			Scope: scope.Scope,
			Key:   put.Key,
		})
		require.NoError(t, err)
	})

	scrts, err := w.Secrets.ListSecretsByScope(ctx, scope.Scope)
	require.NoError(t, err)
	assert.True(t, len(scrts.Secrets) == 1)

	group, err := w.Groups.CreateGroup(ctx, scim.Group{
		DisplayName: RandomName("sdk-go-secret-managers"),
	})
	require.NoError(t, err)

	err = w.Secrets.PutAcl(ctx, secrets.PutAcl{
		Scope:      scope.Scope,
		Permission: secrets.AclPermissionManage,
		Principal:  group.DisplayName,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		w.Secrets.DeleteAcl(ctx, secrets.DeleteAcl{
			Scope:     scope.Scope,
			Principal: group.DisplayName,
		})
	})

	acls, err := w.Secrets.ListAclsByScope(ctx, scope.Scope)
	require.NoError(t, err)

	assert.Equal(t, 2, len(acls.Items))
}
