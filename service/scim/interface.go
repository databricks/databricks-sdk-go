// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package scim

import (
	"context"
)



type CurrentUserService interface {
    
    Me(ctx context.Context) (*User, error)
}


type GroupsService interface {
    // Delete one group
    DeleteGroup(ctx context.Context, deleteGroupRequest DeleteGroupRequest) error
    // Fetch information of one group
    FetchGroup(ctx context.Context, fetchGroupRequest FetchGroupRequest) (*Group, error)
    // Get multiple groups associated with the &lt;Workspace&gt;.
    ListGroups(ctx context.Context, listGroupsRequest ListGroupsRequest) (*ListGroupsResponse, error)
    // Create one group in the &lt;Workspace&gt;.
    NewGroup(ctx context.Context, group Group) (*Group, error)
    // Partially update details of a group
    PatchGroup(ctx context.Context, partialUpdate PartialUpdate) error
    // Update details of a group by replacing the entire entity
    ReplaceGroup(ctx context.Context, group Group) error
	DeleteGroupById(ctx context.Context, id string) error
	FetchGroupById(ctx context.Context, id string) (*Group, error)
}


type ServicePrincipalsService interface {
    // Delete one service principal
    DeleteServicePrincipal(ctx context.Context, deleteServicePrincipalRequest DeleteServicePrincipalRequest) error
    // Fetch information of one service principal
    FetchServicePrincipal(ctx context.Context, fetchServicePrincipalRequest FetchServicePrincipalRequest) (*ServicePrincipal, error)
    // Get multiple service principals associated with a &lt;Workspace&gt;.
    ListServicePrincipals(ctx context.Context, listServicePrincipalsRequest ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error)
    // Create one service principal in the &lt;Workspace&gt;.
    NewServicePrincipal(ctx context.Context, servicePrincipal ServicePrincipal) (*ServicePrincipal, error)
    // Partially update details of one service principal.
    PatchServicePrincipal(ctx context.Context, partialUpdate PartialUpdate) error
    // Update details of one service principal.
    ReplaceServicePrincipal(ctx context.Context, servicePrincipal ServicePrincipal) error
	FetchServicePrincipalById(ctx context.Context, id string) (*ServicePrincipal, error)
	DeleteServicePrincipalById(ctx context.Context, id string) error
}


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
