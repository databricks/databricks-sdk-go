package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
)

func TestAccSecrets(t *testing.T) {
	ctx, w := workspaceTest(t)

	scopeName := RandomName("scope-")
	keyName := RandomName("key-")

	// creates scopeName
	err := w.Secrets.CreateScope(ctx, workspace.CreateScope{
		Scope: scopeName,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Secrets.DeleteScopeByScope(ctx, scopeName)
		assert.NoError(t, err)
	})

	scopes, err := w.Secrets.ListScopesAll(ctx)
	assert.NoError(t, err)
	assert.True(t, len(scopes) >= 1)

	// creates keyName
	err = w.Secrets.PutSecret(ctx, workspace.PutSecret{
		Scope:       scopeName,
		Key:         keyName,
		StringValue: RandomName("dummy"),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Secrets.DeleteSecret(ctx, workspace.DeleteSecret{
			Scope: scopeName,
			Key:   keyName,
		})
		assert.NoError(t, err)
	})

	scrts, err := w.Secrets.ListSecretsByScope(ctx, scopeName)
	assert.NoError(t, err)
	assert.True(t, len(scrts.Secrets) == 1)

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Groups.DeleteById(ctx, group.Id)
		assert.NoError(t, err)
	})

	err = w.Secrets.PutAcl(ctx, workspace.PutAcl{
		Scope:      scopeName,
		Permission: workspace.AclPermissionManage,
		Principal:  group.DisplayName,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		w.Secrets.DeleteAcl(ctx, workspace.DeleteAcl{
			Scope:     scopeName,
			Principal: group.DisplayName,
		})
	})

	acls, err := w.Secrets.ListAclsByScope(ctx, scopeName)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(acls.Items))
}
