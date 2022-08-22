package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/service/groups"
	"github.com/databricks/databricks-sdk-go/service/users"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccUsers(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ctx := context.TODO()

	ws := workspaces.New()

	// create new user
	user, err := ws.Users.NewUser(ctx, users.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	require.NoError(t, err)

	// fetch the user by the newly created ID
	fetch, err := ws.Users.FetchUserById(ctx, user.Id)
	require.NoError(t, err)
	assert.Equal(t, user.DisplayName, fetch.DisplayName)

	// list all users
	allUsers, err := ws.Users.ListUsers(ctx, users.ListUsersRequest{
		Attributes: "id,userName",
		SortBy:     "userName",
		SortOrder:  users.ListUsersSortOrderDescending,
	})
	require.NoError(t, err)

	// verify that the user we've creates is in the list
	namesToIds := map[string]string{}
	for _, u := range allUsers.Resources {
		namesToIds[u.UserName] = u.Id
	}
	assert.Equal(t, user.Id, namesToIds[user.UserName])

	// remove user by ID
	err = ws.Users.DeleteUserById(ctx, user.Id)
	require.NoError(t, err)

	// and verify that user is missing
	_, err = ws.Users.FetchUserById(ctx, user.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccGroups(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ctx := context.TODO()

	ws := workspaces.New()

	// create new group
	group, err := ws.Groups.NewGroup(ctx, groups.Group{
		DisplayName: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	// fetch the group we've just created
	fetch, err := ws.Groups.FetchGroupById(ctx, group.Id)
	require.NoError(t, err)
	assert.Equal(t, group.DisplayName, fetch.DisplayName)

	// list all groups that start with `go-sdk-`
	allGroups, err := ws.Groups.ListGroups(ctx, groups.ListGroupsRequest{
		SortOrder: groups.ListGroupsSortOrderDescending,
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
	err = ws.Groups.DeleteGroupById(ctx, group.Id)
	require.NoError(t, err)

	// and verify the group is missing
	_, err = ws.Groups.FetchGroupById(ctx, group.Id)
	assert.True(t, apierr.IsMissing(err))
}
