// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package users

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type UsersService interface {
    // Delete one user 
    DeleteUser(ctx context.Context, deleteUserRequest DeleteUserRequest) error
    // Fetch information of one user 
    FetchUser(ctx context.Context, fetchUserRequest FetchUserRequest) (*User, error)
    // Get multiple users associated with a &lt;Workspace&gt;. 
    ListUsers(ctx context.Context, listUsersRequest ListUsersRequest) (*ListUsersResponse, error)
    // Create one user in the &lt;Workspace&gt;. 
    NewUser(ctx context.Context, user User) (*User, error)
    // Partially update details of one user. 
    PatchUser(ctx context.Context, partialUpdate PartialUpdate) error
    // Replaces user with the data supplied in request 
    ReplaceUser(ctx context.Context, user User) error
	FetchUserById(ctx context.Context, id string) (*User, error)
	DeleteUserById(ctx context.Context, id string) error
}

func New(client *client.DatabricksClient) UsersService {
	return &UsersAPI{
		client: client,
	}
}

type UsersAPI struct {
	client *client.DatabricksClient
}

// Delete one user 
func (a *UsersAPI) DeleteUser(ctx context.Context, request DeleteUserRequest) error {
	path := "/api/2.0/preview/scim/v2/Users/"+request.Id
	err := a.client.Delete(ctx, path, request)
	return err
}

// Fetch information of one user 
func (a *UsersAPI) FetchUser(ctx context.Context, request FetchUserRequest) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Users/"+request.Id
	err := a.client.Get(ctx, path, request, &user)
	return &user, err
}

// Get multiple users associated with a &lt;Workspace&gt;. 
func (a *UsersAPI) ListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Get(ctx, path, request, &listUsersResponse)
	return &listUsersResponse, err
}

// Create one user in the &lt;Workspace&gt;. 
func (a *UsersAPI) NewUser(ctx context.Context, request User) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Post(ctx, path, request, &user)
	return &user, err
}

// Partially update details of one user. 
func (a *UsersAPI) PatchUser(ctx context.Context, request PartialUpdate) error {
	path := "/api/2.0/preview/scim/v2/Users/"+request.Id
	err := a.client.Patch(ctx, path, request)
	return err
}

// Replaces user with the data supplied in request 
func (a *UsersAPI) ReplaceUser(ctx context.Context, request User) error {
	path := "/api/2.0/preview/scim/v2/Users/"+request.Id
	err := a.client.Put(ctx, path, request)
	return err
}


func (a *UsersAPI) FetchUserById(ctx context.Context, id string) (*User, error) {
	return a.FetchUser(ctx, FetchUserRequest{
		Id: id,
	})
}

func (a *UsersAPI) DeleteUserById(ctx context.Context, id string) error {
	return a.DeleteUser(ctx, DeleteUserRequest{
		Id: id,
	})
}
