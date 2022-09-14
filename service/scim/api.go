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

// Fetch details about caller identity
//
// Get details about caller identity
func (a *CurrentUserAPI) Me(ctx context.Context) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Me"
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

// Delete a group in <Workspace>
//
// Remove a group in the <Workspace>.
func (a *GroupsAPI) DeleteGroup(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a group in <Workspace>
//
// Remove a group in the <Workspace>.
func (a *GroupsAPI) DeleteGroupById(ctx context.Context, id string) error {
	return a.DeleteGroup(ctx, DeleteGroupRequest{
		Id: id,
	})
}

// Fetch details of a group in <Workspace>
//
// Fetch information of one group in the <Workspace>
func (a *GroupsAPI) FetchGroup(ctx context.Context, request FetchGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Get(ctx, path, request, &group)
	return &group, err
}

// Fetch details of a group in <Workspace>
//
// Fetch information of one group in the <Workspace>
func (a *GroupsAPI) FetchGroupById(ctx context.Context, id string) (*Group, error) {
	return a.FetchGroup(ctx, FetchGroupRequest{
		Id: id,
	})
}

// Fetch details of multiple groups in <Workspace>
//
// Get all details of the groups associated with the <Workspace>.
func (a *GroupsAPI) ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Get(ctx, path, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

<<<<<<< HEAD
// Create a new group in <Workspace>
//
// Create one group in the <Workspace> with a unique name
=======
func (a *GroupsAPI) ListGroupsAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	response, err := a.ListGroups(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

func (a *GroupsAPI) GroupDisplayNameToIdMap(ctx context.Context, request ListGroupsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListGroupsAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.DisplayName] = v.Id
	}
	return mapping, nil
}

func (a *GroupsAPI) GetGroupByDisplayName(ctx context.Context, name string) (*Group, error) {
	result, err := a.ListGroupsAll(ctx, ListGroupsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.DisplayName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Group named '%s' does not exist", name)
}

// Create one group in the <Workspace>.
>>>>>>> e767f34 (Unify list return types)
func (a *GroupsAPI) NewGroup(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Post(ctx, path, request, &group)
	return &group, err
}

// Update details of a group
//
// Partially update details of a group
func (a *GroupsAPI) PatchGroup(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Update details of a group
//
// Update details of a group by replacing the entire entity
func (a *GroupsAPI) ReplaceGroup(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Put(ctx, path, request)
	return err
}

func NewServicePrincipals(client *client.DatabricksClient) ServicePrincipalsService {
	return &ServicePrincipalsAPI{
		client: client,
	}
}

type ServicePrincipalsAPI struct {
	client *client.DatabricksClient
}

// Delete a service principal in <Workspace>
//
// Delete one service principal
func (a *ServicePrincipalsAPI) DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a service principal in <Workspace>
//
// Delete one service principal
func (a *ServicePrincipalsAPI) DeleteServicePrincipalById(ctx context.Context, id string) error {
	return a.DeleteServicePrincipal(ctx, DeleteServicePrincipalRequest{
		Id: id,
	})
}

// Fetch details of a service principal in <Workspace>
//
// Fetch information of one service principal
func (a *ServicePrincipalsAPI) FetchServicePrincipal(ctx context.Context, request FetchServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Get(ctx, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Fetch details of a service principal in <Workspace>
//
// Fetch information of one service principal
func (a *ServicePrincipalsAPI) FetchServicePrincipalById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.FetchServicePrincipal(ctx, FetchServicePrincipalRequest{
		Id: id,
	})
}

// Fetch details of multiple service principals in <Workspace>
//
// Get multiple service principals associated with a <Workspace>.
func (a *ServicePrincipalsAPI) ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Get(ctx, path, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

<<<<<<< HEAD
// Create a new service principal in <Workspace>
//
=======
func (a *ServicePrincipalsAPI) ListServicePrincipalsAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	response, err := a.ListServicePrincipals(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

func (a *ServicePrincipalsAPI) ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListServicePrincipalsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListServicePrincipalsAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.DisplayName] = v.Id
	}
	return mapping, nil
}

func (a *ServicePrincipalsAPI) GetServicePrincipalByDisplayName(ctx context.Context, name string) (*ServicePrincipal, error) {
	result, err := a.ListServicePrincipalsAll(ctx, ListServicePrincipalsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.DisplayName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("ServicePrincipal named '%s' does not exist", name)
}

>>>>>>> e767f34 (Unify list return types)
// Create one service principal in the <Workspace>.
func (a *ServicePrincipalsAPI) NewServicePrincipal(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Post(ctx, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Update details of a service principal in <Workspace>
//
// Partially update details of one service principal.
func (a *ServicePrincipalsAPI) PatchServicePrincipal(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Replace service principal in <Workspace>
//
// Update details of one service principal.
func (a *ServicePrincipalsAPI) ReplaceServicePrincipal(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Put(ctx, path, request)
	return err
}

func NewUsers(client *client.DatabricksClient) UsersService {
	return &UsersAPI{
		client: client,
	}
}

type UsersAPI struct {
	client *client.DatabricksClient
}

// Delete a user in <Workspace>
//
// Delete one user. Deleting a user from a workspace also removes objects
// associated with the user.
func (a *UsersAPI) DeleteUser(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a user in <Workspace>
//
// Delete one user. Deleting a user from a workspace also removes objects
// associated with the user.
func (a *UsersAPI) DeleteUserById(ctx context.Context, id string) error {
	return a.DeleteUser(ctx, DeleteUserRequest{
		Id: id,
	})
}

// Get details of a user in <Workspace>
//
// Fetch information of one user in <Workspace>
func (a *UsersAPI) FetchUser(ctx context.Context, request FetchUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Get(ctx, path, request, &user)
	return &user, err
}

// Get details of a user in <Workspace>
//
// Fetch information of one user in <Workspace>
func (a *UsersAPI) FetchUserById(ctx context.Context, id string) (*User, error) {
	return a.FetchUser(ctx, FetchUserRequest{
		Id: id,
	})
}

// Fetch details of multiple users in <Workspace>
//
// Get all the users associated with a <Workspace>.
func (a *UsersAPI) ListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Get(ctx, path, request, &listUsersResponse)
	return &listUsersResponse, err
}

<<<<<<< HEAD
// Create a new user in <Workspace>
//
// Create a user in the <Workspace> who will automatically added to the account.
=======
func (a *UsersAPI) ListUsersAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	response, err := a.ListUsers(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

func (a *UsersAPI) UserUserNameToIdMap(ctx context.Context, request ListUsersRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListUsersAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.UserName] = v.Id
	}
	return mapping, nil
}

func (a *UsersAPI) GetUserByUserName(ctx context.Context, name string) (*User, error) {
	result, err := a.ListUsersAll(ctx, ListUsersRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.UserName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("User named '%s' does not exist", name)
}

// Create one user in the <Workspace>.
>>>>>>> e767f34 (Unify list return types)
func (a *UsersAPI) NewUser(ctx context.Context, request User) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Post(ctx, path, request, &user)
	return &user, err
}

// Update details of a user in <Workspace>
//
// Partially update a user resource with operations on specific attributes
func (a *UsersAPI) PatchUser(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Update details of a user in <Workspace>
//
// Replaces user with the data supplied in request
func (a *UsersAPI) ReplaceUser(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Put(ctx, path, request)
	return err
}
