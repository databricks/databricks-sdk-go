package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccCurrentUser(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.TODO()
	wsc := workspaces.New()

	me, err := wsc.CurrentUser.Me(ctx)
	assert.NoError(t, err)

	assert.NotEmpty(t, me.UserName)
}

func TestAccUsers(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.TODO()
	wsc := workspaces.New()

	// create new user
	user, err := wsc.Users.CreateUser(ctx, scim.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	require.NoError(t, err)

	// fetch the user by the newly created ID
	fetch, err := wsc.Users.GetUserById(ctx, user.Id)
	require.NoError(t, err)
	assert.Equal(t, user.DisplayName, fetch.DisplayName)

	// list all users
	allUsers, err := wsc.Users.ListUsers(ctx, scim.ListUsersRequest{
		Attributes: "id,userName",
		SortBy:     "userName",
		SortOrder:  scim.ListUsersSortOrderDescending,
	})
	require.NoError(t, err)

	// verify that the user we've creates is in the list
	namesToIds := map[string]string{}
	for _, u := range allUsers.Resources {
		namesToIds[u.UserName] = u.Id
	}
	assert.Equal(t, user.Id, namesToIds[user.UserName])

	// remove user by ID
	err = wsc.Users.DeleteUserById(ctx, user.Id)
	require.NoError(t, err)

	// and verify that user is missing
	_, err = wsc.Users.GetUserById(ctx, user.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccGroups(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.TODO()
	wsc := workspaces.New()

	// create new group
	group, err := wsc.Groups.CreateGroup(ctx, scim.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	// fetch the group we've just created
	fetch, err := wsc.Groups.GetGroupById(ctx, group.Id)
	require.NoError(t, err)
	assert.Equal(t, group.DisplayName, fetch.DisplayName)

	// list all groups that start with `go-sdk-`
	allGroups, err := wsc.Groups.ListGroups(ctx, scim.ListGroupsRequest{
		SortOrder: scim.ListGroupsSortOrderDescending,
		Filter:    "displayName sw 'go-sdk-'",
	})
	require.NoError(t, err)

	// verify that the group we've creates is in the list
	namesToIds := map[string]string{}
	for _, u := range allGroups.Resources {
		namesToIds[u.DisplayName] = u.Id
	}
	assert.Equal(t, group.Id, namesToIds[group.DisplayName])

	// remove group by ID
	err = wsc.Groups.DeleteGroupById(ctx, group.Id)
	require.NoError(t, err)

	// and verify the group is missing
	_, err = wsc.Groups.GetGroupById(ctx, group.Id)
	assert.True(t, apierr.IsMissing(err))
}
