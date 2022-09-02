// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package scim

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewCurrentUser(client *client.DatabricksClient) CurrentUserService {
	return &CurrentUserAPI{
		client: client,
	}
}

type CurrentUserAPI struct {
	client *client.DatabricksClient
}


func (a *CurrentUserAPI) Me(ctx context.Context) (*User, error) {
	var user User
	path := "/preview/scim/v2/Me"
	err := a.client.Get(ctx, path, nil, &user)
	return &user, err
}


func NewGroups(client *client.DatabricksClient) GroupsService {
	return &GroupsAPI{
		client: client,
	}
}

type GroupsAPI struct {
	client *client.DatabricksClient
}

// Delete one group
func (a *GroupsAPI) DeleteGroup(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Fetch information of one group
func (a *GroupsAPI) FetchGroup(ctx context.Context, request FetchGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Get(ctx, path, request, &group)
	return &group, err
}

// Get multiple groups associated with the &lt;Workspace&gt;.
func (a *GroupsAPI) ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Get(ctx, path, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

// Create one group in the &lt;Workspace&gt;.
func (a *GroupsAPI) NewGroup(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Post(ctx, path, request, &group)
	return &group, err
}

// Partially update details of a group
func (a *GroupsAPI) PatchGroup(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Update details of a group by replacing the entire entity
func (a *GroupsAPI) ReplaceGroup(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Put(ctx, path, request)
	return err
}


func (a *GroupsAPI) DeleteGroupById(ctx context.Context, id string) error {
	return a.DeleteGroup(ctx, DeleteGroupRequest{
		Id: id,
	})
}

func (a *GroupsAPI) FetchGroupById(ctx context.Context, id string) (*Group, error) {
	return a.FetchGroup(ctx, FetchGroupRequest{
		Id: id,
	})
}

func NewServicePrincipals(client *client.DatabricksClient) ServicePrincipalsService {
	return &ServicePrincipalsAPI{
		client: client,
	}
}

type ServicePrincipalsAPI struct {
	client *client.DatabricksClient
}

// Delete one service principal
func (a *ServicePrincipalsAPI) DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Fetch information of one service principal
func (a *ServicePrincipalsAPI) FetchServicePrincipal(ctx context.Context, request FetchServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Get(ctx, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Get multiple service principals associated with a &lt;Workspace&gt;.
func (a *ServicePrincipalsAPI) ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Get(ctx, path, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

// Create one service principal in the &lt;Workspace&gt;.
func (a *ServicePrincipalsAPI) NewServicePrincipal(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Post(ctx, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Partially update details of one service principal.
func (a *ServicePrincipalsAPI) PatchServicePrincipal(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Update details of one service principal.
func (a *ServicePrincipalsAPI) ReplaceServicePrincipal(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Put(ctx, path, request)
	return err
}


func (a *ServicePrincipalsAPI) FetchServicePrincipalById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.FetchServicePrincipal(ctx, FetchServicePrincipalRequest{
		Id: id,
	})
}

func (a *ServicePrincipalsAPI) DeleteServicePrincipalById(ctx context.Context, id string) error {
	return a.DeleteServicePrincipal(ctx, DeleteServicePrincipalRequest{
		Id: id,
	})
}

func NewUsers(client *client.DatabricksClient) UsersService {
	return &UsersAPI{
		client: client,
	}
}

type UsersAPI struct {
	client *client.DatabricksClient
}

// Delete one user
func (a *UsersAPI) DeleteUser(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Fetch information of one user
func (a *UsersAPI) FetchUser(ctx context.Context, request FetchUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
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
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Replaces user with the data supplied in request
func (a *UsersAPI) ReplaceUser(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
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
