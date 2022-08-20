// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package users

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func New(cfg *databricks.Config) UsersService {
	return &UsersAPI{
		client: client.New(cfg),
	}
}


type UsersService interface {
    // Delete one user 
    DeleteUser(ctx context.Context, deleteUserRequest DeleteUserRequest) error
    // Fetch information of one user 
    FetchUser(ctx context.Context, fetchUserRequest FetchUserRequest) (*User, error)
    // Get multiple users associated with a &lt;Workspace&gt;. 
    ListUsers(ctx context.Context, listUsersRequest ListUsersRequest) (*ListUsersResponse, error)
    // Create one user in the &lt;Workspace&gt;. 
    NewUser(ctx context.Context, createUserRequest CreateUserRequest) (*User, error)
    // Update details of one user. 
    UpdateUser(ctx context.Context, updateUserRequest UpdateUserRequest) error
	DeleteUserByUserId(ctx context.Context, userId string) error
	FetchUserByUserId(ctx context.Context, userId string) (*User, error)
}

type UsersAPI struct {
	client *client.DatabricksClient
}

// Delete one user 
func (a *UsersAPI) DeleteUser(ctx context.Context, in DeleteUserRequest) error {
	
	err := a.client.Delete(ctx, "/scim/v2/Users/{user_id}", in)
	return err
}

// Fetch information of one user 
func (a *UsersAPI) FetchUser(ctx context.Context, in FetchUserRequest) (*User, error) {
	var user User
	err := a.client.Get(ctx, "/scim/v2/Users/{user_id}", in, &user)
	return &user, err
}

// Get multiple users associated with a &lt;Workspace&gt;. 
func (a *UsersAPI) ListUsers(ctx context.Context, in ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	err := a.client.Get(ctx, "/scim/v2/Users", in, &listUsersResponse)
	return &listUsersResponse, err
}

// Create one user in the &lt;Workspace&gt;. 
func (a *UsersAPI) NewUser(ctx context.Context, in CreateUserRequest) (*User, error) {
	var user User
	err := a.client.Post(ctx, "/scim/v2/Users", in, &user)
	return &user, err
}

// Update details of one user. 
func (a *UsersAPI) UpdateUser(ctx context.Context, in UpdateUserRequest) error {
	
	err := a.client.Patch(ctx, "/scim/v2/Users/{user_id}", in)
	return err
}


func (a *UsersAPI) DeleteUserByUserId(ctx context.Context, userId string) error {
	return a.DeleteUser(ctx, DeleteUserRequest{
		UserId: userId,
	})
}

func (a *UsersAPI) FetchUserByUserId(ctx context.Context, userId string) (*User, error) {
	return a.FetchUser(ctx, FetchUserRequest{
		UserId: userId,
	})
}
