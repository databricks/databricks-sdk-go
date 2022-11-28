package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccCurrentUser(t *testing.T) {
	ctx, w := workspaceTest(t)

	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)

	assert.NotEmpty(t, me.UserName)
}

func TestAccUsers(t *testing.T) {
	ctx, w := workspaceTest(t)

	// create new user
	user, err := w.Users.Create(ctx, scim.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	require.NoError(t, err)

	// fetch the user by the newly created ID
	fetch, err := w.Users.GetById(ctx, user.Id)
	require.NoError(t, err)
	assert.Equal(t, user.DisplayName, fetch.DisplayName)

	// list all users
	allUsers, err := w.Users.ListAll(ctx, scim.ListUsersRequest{
		Attributes: "id,userName",
		SortBy:     "userName",
		SortOrder:  scim.ListSortOrderDescending,
	})
	require.NoError(t, err)

	// verify that the user we've creates is in the list
	namesToIds, err := w.Users.UserUserNameToIdMap(ctx, scim.ListUsersRequest{
		Attributes: "id,userName",
	})
	require.NoError(t, err)
	assert.Equal(t, len(namesToIds), len(allUsers))
	assert.Equal(t, user.Id, namesToIds[user.UserName])

	// remove user by ID
	err = w.Users.DeleteById(ctx, user.Id)
	require.NoError(t, err)

	// and verify that user is missing
	_, err = w.Users.GetById(ctx, user.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccGroups(t *testing.T) {
	ctx, w := workspaceTest(t)

	// create new group
	group, err := w.Groups.Create(ctx, scim.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	// fetch the group we've just created
	fetch, err := w.Groups.GetById(ctx, group.Id)
	require.NoError(t, err)
	assert.Equal(t, group.DisplayName, fetch.DisplayName)

	// list all groups that start with `go-sdk-`
	namesToIds, err := w.Groups.GroupDisplayNameToIdMap(ctx, scim.ListGroupsRequest{
		SortOrder: scim.ListSortOrderDescending,
		Filter:    "displayName sw 'go-sdk-'",
	})
	require.NoError(t, err)
	assert.Equal(t, group.Id, namesToIds[group.DisplayName])

	// remove group by ID
	err = w.Groups.DeleteById(ctx, group.Id)
	require.NoError(t, err)

	// and verify the group is missing
	_, err = w.Groups.GetById(ctx, group.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccServicePrincipalsOnAWS(t *testing.T) {
	ctx, w := workspaceTest(t)
	if !w.Config.IsAws() {
		t.Skip("test only for aws")
	}

	created, err := w.ServicePrincipals.Create(ctx, scim.ServicePrincipal{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.ServicePrincipals.DeleteById(ctx, created.Id)
		require.NoError(t, err)
	})
	err = w.ServicePrincipals.Update(ctx, scim.ServicePrincipal{
		Id:          created.Id,
		DisplayName: RandomName("go-sdk-updated-"),
		Roles: []scim.ComplexValue{
			{
				Value: "xyz",
			},
		},
	})
	require.NoError(t, err)

	byId, err := w.ServicePrincipals.GetById(ctx, created.Id)
	require.NoError(t, err)

	all, err := w.ServicePrincipals.ListAll(ctx, scim.ListServicePrincipalsRequest{})
	require.NoError(t, err)

	names, err := w.ServicePrincipals.ServicePrincipalDisplayNameToIdMap(ctx, scim.ListServicePrincipalsRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.Id, names[byId.DisplayName])
}
