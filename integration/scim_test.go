package integration

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/iam/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func me(t *testing.T, cfg *config.Config) *iam.User {
	ctx := context.Background()
	CurrentUserAPI, err := iam.NewCurrentUserClient(cfg)
	require.NoError(t, err)
	me, err := CurrentUserAPI.Me(ctx)
	require.NoError(t, err)
	return me
}

func TestAccCurrentUser(t *testing.T) {
	ctx := workspaceTest(t)

	CurrentUserAPI, err := iam.NewCurrentUserClient(nil)
	require.NoError(t, err)
	me, err := CurrentUserAPI.Me(ctx)
	require.NoError(t, err)

	assert.NotEmpty(t, me.UserName)
}

func TestAccWorkspaceUsers(t *testing.T) {
	ctx := workspaceTest(t)

	// create new user
	UsersAPI, err := iam.NewUsersClient(nil)
	require.NoError(t, err)
	user, err := UsersAPI.Create(ctx, iam.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	require.NoError(t, err)

	// fetch the user by the newly created ID
	fetch, err := UsersAPI.GetById(ctx, user.Id)
	require.NoError(t, err)
	assert.Equal(t, user.DisplayName, fetch.DisplayName)

	err = UsersAPI.Patch(ctx, iam.PartialUpdate{
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
	require.NoError(t, err)

	err = UsersAPI.Update(ctx, iam.User{
		Id:       user.Id,
		UserName: user.UserName,
		Active:   true,
	})
	require.NoError(t, err)

	byName, err := UsersAPI.GetByUserName(ctx, fetch.UserName)
	require.NoError(t, err)
	assert.Equal(t, fetch.Id, byName.Id)

	// list all users
	allUsers, err := UsersAPI.ListAll(ctx, iam.ListUsersRequest{
		Attributes: "id,userName",
		SortBy:     "userName",
		SortOrder:  iam.ListSortOrderDescending,
	})
	require.NoError(t, err)

	// verify that the user we've creates is in the list
	namesToIds, err := UsersAPI.UserUserNameToIdMap(ctx, iam.ListUsersRequest{
		Attributes: "id,userName",
	})
	require.NoError(t, err)
	assert.Equal(t, len(namesToIds), len(allUsers))
	assert.Equal(t, user.Id, namesToIds[user.UserName])

	// remove user by ID
	err = UsersAPI.DeleteById(ctx, user.Id)
	require.NoError(t, err)

	// and verify that user is missing
	_, err = UsersAPI.GetById(ctx, user.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccAccountUsers(t *testing.T) {
	ctx, cfg := accountTest(t)

	UsersAPI, err := iam.NewAccountUsersClient(cfg)
	require.NoError(t, err)
	user, err := UsersAPI.Create(ctx, iam.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := UsersAPI.DeleteById(ctx, user.Id)
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	assert.Equal(t, 0, len(user.Roles))
	err = UsersAPI.Patch(ctx, iam.PartialUpdate{
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
	require.NoError(t, err)

	byId, err := UsersAPI.GetById(ctx, user.Id)
	require.NoError(t, err)
	assert.Equal(t, user.Id, byId.Id)
	assert.Equal(t, 1, len(byId.Roles))
}

func TestAccGroups(t *testing.T) {
	ctx := workspaceTest(t)

	// create new group
	GroupsAPI, err := iam.NewGroupsClient(nil)
	require.NoError(t, err)
	group, err := GroupsAPI.Create(ctx, iam.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := GroupsAPI.DeleteById(ctx, group.Id)
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	// fetch the group we've just created
	fetch, err := GroupsAPI.GetById(ctx, group.Id)
	require.NoError(t, err)
	assert.Equal(t, group.DisplayName, fetch.DisplayName)

	byName, err := GroupsAPI.GetByDisplayName(ctx, fetch.DisplayName)
	require.NoError(t, err)
	assert.Equal(t, fetch.Id, byName.Id)

	// list all groups that start with `go-sdk-`
	namesToIds, err := GroupsAPI.GroupDisplayNameToIdMap(ctx, iam.ListGroupsRequest{
		SortOrder:          iam.ListSortOrderDescending,
		ExcludedAttributes: "roles",
		Filter:             "displayName sw 'go-sdk-'",
	})
	require.NoError(t, err)
	assert.Equal(t, group.Id, namesToIds[group.DisplayName])

	// remove group by ID
	err = GroupsAPI.DeleteById(ctx, group.Id)
	require.NoError(t, err)

	// and verify the group is missing
	_, err = GroupsAPI.GetById(ctx, group.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccServicePrincipalsOnAWS(t *testing.T) {
	ctx := workspaceTest(t)

	ServicePrincipalsAPI, err := iam.NewServicePrincipalsClient(nil)
	require.NoError(t, err)
	if !ServicePrincipalsAPI.Config.IsAws() {
		t.Skip("test only for aws")
	}

	created, err := ServicePrincipalsAPI.Create(ctx, iam.ServicePrincipal{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := ServicePrincipalsAPI.DeleteById(ctx, created.Id)
		require.NoError(t, err)
	})
	err = ServicePrincipalsAPI.Update(ctx, iam.ServicePrincipal{
		Id:          created.Id,
		DisplayName: RandomName("go-sdk-updated-"),
		Roles: []iam.ComplexValue{
			{
				Value: "xyz",
			},
		},
	})
	require.NoError(t, err)

	byId, err := ServicePrincipalsAPI.GetById(ctx, created.Id)
	require.NoError(t, err)

	byName, err := ServicePrincipalsAPI.GetByDisplayName(ctx, byId.DisplayName)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	err = ServicePrincipalsAPI.Patch(ctx, iam.PartialUpdate{
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
	require.NoError(t, err)

	all, err := ServicePrincipalsAPI.ListAll(ctx, iam.ListServicePrincipalsRequest{})
	require.NoError(t, err)

	names, err := ServicePrincipalsAPI.ServicePrincipalDisplayNameToIdMap(ctx, iam.ListServicePrincipalsRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.Id, names[byId.DisplayName])
}
