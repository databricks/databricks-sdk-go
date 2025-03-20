package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/iam/v2"
	"github.com/databricks/databricks-sdk-go/workspace/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSecrets(t *testing.T) {
	ctx := workspaceTest(t)

	scopeName := RandomName("scope-")
	keyName := RandomName("key-")

	// creates scopeName
	SecretsAPI, err := workspace.NewSecretsClient(nil)
	require.NoError(t, err)
	_, err = SecretsAPI.CreateScope(ctx, workspace.CreateScope{
		Scope: scopeName,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err = SecretsAPI.DeleteScopeByScope(ctx, scopeName)
		require.NoError(t, err)
	})

	scopes, err := SecretsAPI.ListScopesAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(scopes) >= 1)

	// creates keyName
	_, err = SecretsAPI.PutSecret(ctx, workspace.PutSecret{
		Scope:       scopeName,
		Key:         keyName,
		StringValue: RandomName("dummy"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err = SecretsAPI.DeleteSecret(ctx, workspace.DeleteSecret{
			Scope: scopeName,
			Key:   keyName,
		})
		require.NoError(t, err)
	})

	scrts, err := SecretsAPI.ListSecretsByScope(ctx, scopeName)
	require.NoError(t, err)
	assert.True(t, len(scrts.Secrets) == 1)

	GroupsAPI, err := iam.NewGroupsClient(nil)
	require.NoError(t, err)
	group, err := GroupsAPI.Create(ctx, iam.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err = GroupsAPI.DeleteById(ctx, group.Id)
		require.NoError(t, err)
	})

	_, err = SecretsAPI.PutAcl(ctx, workspace.PutAcl{
		Scope:      scopeName,
		Permission: workspace.AclPermissionManage,
		Principal:  group.DisplayName,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		SecretsAPI.DeleteAcl(ctx, workspace.DeleteAcl{
			Scope:     scopeName,
			Principal: group.DisplayName,
		})
	})

	acls, err := SecretsAPI.ListAclsByScope(ctx, scopeName)
	require.NoError(t, err)

	assert.Equal(t, 2, len(acls.Items))
}
