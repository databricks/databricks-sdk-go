package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
)

func me(t *testing.T, w *databricks.WorkspaceClient) *iam.User {
	ctx := context.Background()
	me, err := w.CurrentUser.Me(ctx)
	assert.NoError(t, err)
	return me
}

func TestAccCurrentUser(t *testing.T) {
	ctx, w := workspaceTest(t)

	me, err := w.CurrentUser.Me(ctx)
	assert.NoError(t, err)

	assert.NotEmpty(t, me.UserName)
}

func TestAccWorkspaceUsers(t *testing.T) {
	ctx, w := workspaceTest(t)

	// create new user
	user, err := w.Users.Create(ctx, iam.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	assert.NoError(t, err)

	// fetch the user by the newly created ID
	fetch, err := w.Users.GetById(ctx, user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.DisplayName, fetch.DisplayName)

	err = w.Users.Patch(ctx, iam.PartialUpdate{
		Id: user.Id,
		Operations: []iam.Patch{
			{
				Op:    iam.PatchOpReplace,
				Path:  "active",
				Value: "false",
			},
		},
		Schemas: []iam.PatchSchema{
			iam.PatchSchemaUrnIetfParamsScimApiMessages20PatchOp,
		},
	})
	assert.NoError(t, err)

	err = w.Users.Update(ctx, iam.User{
		Id:       user.Id,
		UserName: user.UserName,
		Active:   true,
	})
	assert.NoError(t, err)

	byName, err := w.Users.GetByUserName(ctx, fetch.UserName)
	assert.NoError(t, err)
	assert.Equal(t, fetch.Id, byName.Id)

	// list all users
	allUsers, err := w.Users.ListAll(ctx, iam.ListUsersRequest{
		Attributes: "id,userName",
		SortBy:     "userName",
		SortOrder:  iam.ListSortOrderDescending,
	})
	assert.NoError(t, err)

	// verify that the user we've creates is in the list
	namesToIds, err := w.Users.UserUserNameToIdMap(ctx, iam.ListUsersRequest{
		Attributes: "id,userName",
	})
	assert.NoError(t, err)
	assert.Equal(t, len(namesToIds), len(allUsers))
	assert.Equal(t, user.Id, namesToIds[user.UserName])

	// remove user by ID
	err = w.Users.DeleteById(ctx, user.Id)
	assert.NoError(t, err)

	// and verify that user is missing
	_, err = w.Users.GetById(ctx, user.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccAccountUsers(t *testing.T) {
	ctx, a := accountTest(t)

	user, err := a.Users.Create(ctx, iam.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := a.Users.DeleteById(ctx, user.Id)
		assert.True(t, err == nil || apierr.IsMissing(err))
	})

	assert.Equal(t, 0, len(user.Roles))
	err = a.Users.Patch(ctx, iam.PartialUpdate{
		Id:      user.Id,
		Schemas: []iam.PatchSchema{iam.PatchSchemaUrnIetfParamsScimApiMessages20PatchOp},
		Operations: []iam.Patch{
			{Op: iam.PatchOpAdd, Value: iam.User{
				Roles: []iam.ComplexValue{
					{Value: "account_admin"},
				},
			}},
		},
	})
	assert.NoError(t, err)

	byId, err := a.Users.GetById(ctx, user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, byId.Id)
	assert.Equal(t, 1, len(byId.Roles))
}

func TestAccGroups(t *testing.T) {
	ctx, w := workspaceTest(t)

	// create new group
	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := w.Groups.DeleteById(ctx, group.Id)
		assert.True(t, err == nil || apierr.IsMissing(err))
	})

	// fetch the group we've just created
	fetch, err := w.Groups.GetById(ctx, group.Id)
	assert.NoError(t, err)
	assert.Equal(t, group.DisplayName, fetch.DisplayName)

	byName, err := w.Groups.GetByDisplayName(ctx, fetch.DisplayName)
	assert.NoError(t, err)
	assert.Equal(t, fetch.Id, byName.Id)

	// list all groups that start with `go-sdk-`
	namesToIds, err := w.Groups.GroupDisplayNameToIdMap(ctx, iam.ListGroupsRequest{
		SortOrder:          iam.ListSortOrderDescending,
		ExcludedAttributes: "roles",
		Filter:             "displayName sw 'go-sdk-'",
	})
	assert.NoError(t, err)
	assert.Equal(t, group.Id, namesToIds[group.DisplayName])

	// remove group by ID
	err = w.Groups.DeleteById(ctx, group.Id)
	assert.NoError(t, err)

	// and verify the group is missing
	_, err = w.Groups.GetById(ctx, group.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccServicePrincipalsOnAWS(t *testing.T) {
	ctx, w := workspaceTest(t)
	if !w.Config.IsAws() {
		t.Skip("test only for aws")
	}

	created, err := w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: RandomName("go-sdk-"),
	})
	assert.NoError(t, err)

	t.Cleanup(func() {
		err := w.ServicePrincipals.DeleteById(ctx, created.Id)
		assert.NoError(t, err)
	})
	err = w.ServicePrincipals.Update(ctx, iam.ServicePrincipal{
		Id:          created.Id,
		DisplayName: RandomName("go-sdk-updated-"),
		Roles: []iam.ComplexValue{
			{
				Value: "xyz",
			},
		},
	})
	assert.NoError(t, err)

	byId, err := w.ServicePrincipals.GetById(ctx, created.Id)
	assert.NoError(t, err)

	byName, err := w.ServicePrincipals.GetByDisplayName(ctx, byId.DisplayName)
	assert.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	err = w.ServicePrincipals.Patch(ctx, iam.PartialUpdate{
		Id: byId.Id,
		Operations: []iam.Patch{
			{
				Op:    iam.PatchOpReplace,
				Path:  "active",
				Value: "false",
			},
		},
		Schemas: []iam.PatchSchema{
			iam.PatchSchemaUrnIetfParamsScimApiMessages20PatchOp,
		},
	})
	assert.NoError(t, err)

	all, err := w.ServicePrincipals.ListAll(ctx, iam.ListServicePrincipalsRequest{})
	assert.NoError(t, err)

	names, err := w.ServicePrincipals.ServicePrincipalDisplayNameToIdMap(ctx, iam.ListServicePrincipalsRequest{})
	assert.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.Id, names[byId.DisplayName])
}
