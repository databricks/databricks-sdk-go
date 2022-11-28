package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccCurrentUser(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.TODO()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	me, err := w.CurrentUser.Me(ctx)
	assert.NoError(t, err)

	assert.NotEmpty(t, me.UserName)
}

func TestAccUsers(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.TODO()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

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
	namesToIds := map[string]string{}
	for _, u := range allUsers {
		namesToIds[u.UserName] = u.Id
	}
	assert.Equal(t, user.Id, namesToIds[user.UserName])

	// remove user by ID
	err = w.Users.DeleteById(ctx, user.Id)
	require.NoError(t, err)

	// and verify that user is missing
	_, err = w.Users.GetById(ctx, user.Id)
	assert.True(t, apierr.IsMissing(err))
}

func TestAccGroups(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.TODO()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

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
	allGroups, err := w.Groups.ListAll(ctx, scim.ListGroupsRequest{
		SortOrder: scim.ListSortOrderDescending,
		Filter:    "displayName sw 'go-sdk-'",
	})
	require.NoError(t, err)

	// verify that the group we've creates is in the list
	namesToIds := map[string]string{}
	for _, u := range allGroups {
		namesToIds[u.DisplayName] = u.Id
	}
	assert.Equal(t, group.Id, namesToIds[group.DisplayName])

	// remove group by ID
	err = w.Groups.DeleteById(ctx, group.Id)
	require.NoError(t, err)

	// and verify the group is missing
	_, err = w.Groups.GetById(ctx, group.Id)
	assert.True(t, apierr.IsMissing(err))
}
