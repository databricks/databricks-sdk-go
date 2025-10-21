// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Iam V2, Workspace Iam V2, etc.
package iamv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type AccountIamV2Interface interface {

	// TODO: Write description later when this method is implemented
	CreateGroup(ctx context.Context, request CreateGroupRequest) (*Group, error)

	// TODO: Write description later when this method is implemented
	CreateServicePrincipal(ctx context.Context, request CreateServicePrincipalRequest) (*ServicePrincipal, error)

	// TODO: Write description later when this method is implemented
	CreateUser(ctx context.Context, request CreateUserRequest) (*User, error)

	// TODO: Write description later when this method is implemented
	CreateWorkspaceAccessDetail(ctx context.Context, request CreateWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error)

	// TODO: Write description later when this method is implemented
	DeleteGroup(ctx context.Context, request DeleteGroupRequest) error

	// TODO: Write description later when this method is implemented
	DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error

	// TODO: Write description later when this method is implemented
	DeleteUser(ctx context.Context, request DeleteUserRequest) error

	// TODO: Write description later when this method is implemented
	DeleteWorkspaceAccessDetail(ctx context.Context, request DeleteWorkspaceAccessDetailRequest) error

	// TODO: Write description later when this method is implemented
	GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error)

	// TODO: Write description later when this method is implemented
	GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error)

	// TODO: Write description later when this method is implemented
	GetUser(ctx context.Context, request GetUserRequest) (*User, error)

	// Returns the access details for a principal in a workspace. Allows for
	// checking access details for any provisioned principal (user, service
	// principal, or group) in a workspace. * Provisioned principal here refers to
	// one that has been synced into Databricks from the customer's IdP or added
	// explicitly to Databricks via SCIM/UI. Allows for passing in a "view"
	// parameter to control what fields are returned (BASIC by default or FULL).
	GetWorkspaceAccessDetail(ctx context.Context, request GetWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error)

	// TODO: Write description later when this method is implemented
	ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error)

	// TODO: Write description later when this method is implemented
	ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error)

	// TODO: Write description later when this method is implemented
	ListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error)

	// TODO: Write description later when this method is implemented
	ListWorkspaceAccessDetails(ctx context.Context, request ListWorkspaceAccessDetailsRequest) (*ListWorkspaceAccessDetailsResponse, error)

	// Resolves a group with the given external ID from the customer's IdP. If the
	// group does not exist, it will be created in the account. If the customer is
	// not onboarded onto Automatic Identity Management (AIM), this will return an
	// error.
	ResolveGroup(ctx context.Context, request ResolveGroupRequest) (*ResolveGroupResponse, error)

	// Resolves an SP with the given external ID from the customer's IdP. If the SP
	// does not exist, it will be created. If the customer is not onboarded onto
	// Automatic Identity Management (AIM), this will return an error.
	ResolveServicePrincipal(ctx context.Context, request ResolveServicePrincipalRequest) (*ResolveServicePrincipalResponse, error)

	// Resolves a user with the given external ID from the customer's IdP. If the
	// user does not exist, it will be created. If the customer is not onboarded
	// onto Automatic Identity Management (AIM), this will return an error.
	ResolveUser(ctx context.Context, request ResolveUserRequest) (*ResolveUserResponse, error)

	// TODO: Write description later when this method is implemented
	UpdateGroup(ctx context.Context, request UpdateGroupRequest) (*Group, error)

	// TODO: Write description later when this method is implemented
	UpdateServicePrincipal(ctx context.Context, request UpdateServicePrincipalRequest) (*ServicePrincipal, error)

	// TODO: Write description later when this method is implemented
	UpdateUser(ctx context.Context, request UpdateUserRequest) (*User, error)

	// TODO: Write description later when this method is implemented
	UpdateWorkspaceAccessDetail(ctx context.Context, request UpdateWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error)
}

func NewAccountIamV2(client *client.DatabricksClient) *AccountIamV2API {
	return &AccountIamV2API{
		accountIamV2Impl: accountIamV2Impl{
			client: client,
		},
	}
}

// These APIs are used to manage identities and the workspace access of these
// identities in <Databricks>.
type AccountIamV2API struct {
	accountIamV2Impl
}

type WorkspaceIamV2Interface interface {

	// TODO: Write description later when this method is implemented
	CreateGroupProxy(ctx context.Context, request CreateGroupProxyRequest) (*Group, error)

	// TODO: Write description later when this method is implemented
	CreateServicePrincipalProxy(ctx context.Context, request CreateServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// TODO: Write description later when this method is implemented
	CreateUserProxy(ctx context.Context, request CreateUserProxyRequest) (*User, error)

	// TODO: Write description later when this method is implemented
	CreateWorkspaceAccessDetailLocal(ctx context.Context, request CreateWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error)

	// TODO: Write description later when this method is implemented
	DeleteGroupProxy(ctx context.Context, request DeleteGroupProxyRequest) error

	// TODO: Write description later when this method is implemented
	DeleteServicePrincipalProxy(ctx context.Context, request DeleteServicePrincipalProxyRequest) error

	// TODO: Write description later when this method is implemented
	DeleteUserProxy(ctx context.Context, request DeleteUserProxyRequest) error

	// TODO: Write description later when this method is implemented
	DeleteWorkspaceAccessDetailLocal(ctx context.Context, request DeleteWorkspaceAccessDetailLocalRequest) error

	// TODO: Write description later when this method is implemented
	GetGroupProxy(ctx context.Context, request GetGroupProxyRequest) (*Group, error)

	// TODO: Write description later when this method is implemented
	GetServicePrincipalProxy(ctx context.Context, request GetServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// TODO: Write description later when this method is implemented
	GetUserProxy(ctx context.Context, request GetUserProxyRequest) (*User, error)

	// Returns the access details for a principal in the current workspace. Allows
	// for checking access details for any provisioned principal (user, service
	// principal, or group) in the current workspace. * Provisioned principal here
	// refers to one that has been synced into Databricks from the customer's IdP or
	// added explicitly to Databricks via SCIM/UI. Allows for passing in a "view"
	// parameter to control what fields are returned (BASIC by default or FULL).
	GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error)

	// TODO: Write description later when this method is implemented
	ListGroupsProxy(ctx context.Context, request ListGroupsProxyRequest) (*ListGroupsResponse, error)

	// TODO: Write description later when this method is implemented
	ListServicePrincipalsProxy(ctx context.Context, request ListServicePrincipalsProxyRequest) (*ListServicePrincipalsResponse, error)

	// TODO: Write description later when this method is implemented
	ListUsersProxy(ctx context.Context, request ListUsersProxyRequest) (*ListUsersResponse, error)

	// TODO: Write description later when this method is implemented
	ListWorkspaceAccessDetailsLocal(ctx context.Context, request ListWorkspaceAccessDetailsLocalRequest) (*ListWorkspaceAccessDetailsResponse, error)

	// Resolves a group with the given external ID from the customer's IdP. If the
	// group does not exist, it will be created in the account. If the customer is
	// not onboarded onto Automatic Identity Management (AIM), this will return an
	// error.
	ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error)

	// Resolves an SP with the given external ID from the customer's IdP. If the SP
	// does not exist, it will be created. If the customer is not onboarded onto
	// Automatic Identity Management (AIM), this will return an error.
	ResolveServicePrincipalProxy(ctx context.Context, request ResolveServicePrincipalProxyRequest) (*ResolveServicePrincipalResponse, error)

	// Resolves a user with the given external ID from the customer's IdP. If the
	// user does not exist, it will be created. If the customer is not onboarded
	// onto Automatic Identity Management (AIM), this will return an error.
	ResolveUserProxy(ctx context.Context, request ResolveUserProxyRequest) (*ResolveUserResponse, error)

	// TODO: Write description later when this method is implemented
	UpdateGroupProxy(ctx context.Context, request UpdateGroupProxyRequest) (*Group, error)

	// TODO: Write description later when this method is implemented
	UpdateServicePrincipalProxy(ctx context.Context, request UpdateServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// TODO: Write description later when this method is implemented
	UpdateUserProxy(ctx context.Context, request UpdateUserProxyRequest) (*User, error)

	// TODO: Write description later when this method is implemented
	UpdateWorkspaceAccessDetailLocal(ctx context.Context, request UpdateWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error)
}

func NewWorkspaceIamV2(client *client.DatabricksClient) *WorkspaceIamV2API {
	return &WorkspaceIamV2API{
		workspaceIamV2Impl: workspaceIamV2Impl{
			client: client,
		},
	}
}

// These APIs are used to manage identities and the workspace access of these
// identities in <Databricks>.
type WorkspaceIamV2API struct {
	workspaceIamV2Impl
}
