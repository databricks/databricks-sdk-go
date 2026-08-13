// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Iam V2, Workspace Iam V2, etc.
package iamv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
)

type AccountIamV2Interface interface {

	// Creates a group membership (assigns a principal to a group).
	CreateDirectGroupMember(ctx context.Context, request CreateDirectGroupMemberRequest) (*DirectGroupMember, error)

	// Creates a local group in the Databricks account and returns the created
	// group. A local group is one that is not synced from the customer's identity
	// provider, and can be created whether or not Account Identity Management (AIM)
	// is enabled.
	//
	// When AIM is enabled, supplying an external ID returns an error. Use the
	// ExternalGroup resource to sync groups from the identity provider instead.
	CreateGroup(ctx context.Context, request CreateGroupRequest) (*Group, error)

	// Creates a local service principal in the Databricks account and returns the
	// created service principal. A local service principal is one that is not
	// synced from the customer's identity provider, and can be created whether or
	// not Account Identity Management (AIM) is enabled.
	//
	// When AIM is enabled, supplying an external ID returns an error. Use the
	// ExternalServicePrincipal resource to sync service principals from the
	// identity provider instead.
	CreateServicePrincipal(ctx context.Context, request CreateServicePrincipalRequest) (*ServicePrincipal, error)

	// Creates a local user in the Databricks account and returns the created user.
	// A local user is one that is not synced from the customer's identity provider,
	// and can be created whether or not Account Identity Management (AIM) is
	// enabled.
	//
	// When AIM is enabled, supplying an external ID returns an error. Use the
	// ExternalUser resource to sync users from the identity provider instead.
	CreateUser(ctx context.Context, request CreateUserRequest) (*User, error)

	// Creates a workspace assignment for a principal. Entitlements are granted one
	// at a time rather than atomically. If the request fails partway through, the
	// principal stays assigned to the workspace with only some of the requested
	// entitlements. Get the assignment afterwards to confirm which entitlements
	// were granted.
	CreateWorkspaceAssignment(ctx context.Context, request CreateWorkspaceAssignmentRequest) (*WorkspaceAssignment, error)

	// Creates a workspace assignment detail for a principal. Entitlements are
	// granted one at a time rather than atomically. If the request fails partway
	// through, the principal stays assigned to the workspace with only some of the
	// requested entitlements. Get the assignment detail afterwards to confirm which
	// entitlements were granted.
	CreateWorkspaceAssignmentDetail(ctx context.Context, request CreateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)

	// Deletes a group membership (unassigns a principal from a group).
	DeleteDirectGroupMember(ctx context.Context, request DeleteDirectGroupMemberRequest) error

	// Deletes a group from the Databricks account by its internal ID.
	DeleteGroup(ctx context.Context, request DeleteGroupRequest) error

	// Deletes a service principal from the Databricks account by its internal ID.
	DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error

	// Deletes a user from the Databricks account by its internal ID.
	DeleteUser(ctx context.Context, request DeleteUserRequest) error

	// Deletes a workspace assignment for a principal, revoking all of its
	// entitlements. Entitlements are revoked one at a time rather than atomically.
	// If the request fails partway through, the principal stays assigned with some
	// of its original entitlements. Retrying is safe.
	DeleteWorkspaceAssignment(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error

	// Deletes a workspace assignment detail for a principal, revoking all of its
	// entitlements. Entitlements are revoked one at a time rather than atomically.
	// If the request fails partway through, the principal stays assigned with some
	// of its original entitlements. Retrying is safe.
	DeleteWorkspaceAssignmentDetail(ctx context.Context, request DeleteWorkspaceAssignmentDetailRequest) error

	// Gets a provisioned direct member of a group.
	GetDirectGroupMember(ctx context.Context, request GetDirectGroupMemberRequest) (*DirectGroupMember, error)

	// Fetches a group from the Databricks account by its internal ID.
	GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error)

	// Fetches a service principal from the Databricks account by its internal ID.
	GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error)

	// Fetches a user from the Databricks account by its internal ID.
	GetUser(ctx context.Context, request GetUserRequest) (*User, error)

	// Returns the access details for a principal in a workspace. Allows for
	// checking access details for any provisioned principal (user, service
	// principal, or group) in a workspace. * Provisioned principal here refers to
	// one that has been synced into Databricks from the customer's IdP or added
	// explicitly to Databricks via SCIM/UI. Allows for passing in a "view"
	// parameter to control what fields are returned (BASIC by default or FULL).
	GetWorkspaceAccessDetail(ctx context.Context, request GetWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error)

	// Returns the assignment for a principal in a workspace.
	GetWorkspaceAssignment(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspaceAssignment, error)

	// Returns the assignment details for a principal in a workspace.
	GetWorkspaceAssignmentDetail(ctx context.Context, request GetWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)

	// Lists provisioned direct members of a group with their membership source
	// (internal or from identity provider).
	ListDirectGroupMembers(ctx context.Context, request ListDirectGroupMembersRequest) (*ListDirectGroupMembersResponse, error)

	// Lists the groups in the Databricks account, returning one page per call.
	// Supports filtering by group name or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListGroups(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group]

	// Lists the groups in the Databricks account, returning one page per call.
	// Supports filtering by group name or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListGroupsAll(ctx context.Context, request ListGroupsRequest) ([]Group, error)

	// Lists the service principals in the Databricks account, returning one page
	// per call. Supports filtering by application ID or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal]

	// Lists the service principals in the Databricks account, returning one page
	// per call. Supports filtering by application ID or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListServicePrincipalsAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error)

	// Lists all transitive parent groups of a principal.
	ListTransitiveParentGroups(ctx context.Context, request ListTransitiveParentGroupsRequest) (*ListTransitiveParentGroupsResponse, error)

	// Lists the users in the Databricks account, returning one page per call.
	// Supports filtering by username or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListUsers(ctx context.Context, request ListUsersRequest) listing.Iterator[User]

	// Lists the users in the Databricks account, returning one page per call.
	// Supports filtering by username or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListUsersAll(ctx context.Context, request ListUsersRequest) ([]User, error)

	// Lists workspace assignment details for a workspace. The response omits the
	// per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`). To read the entitlements for a single principal,
	// get that principal's assignment detail.
	ListWorkspaceAssignmentDetails(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) (*ListWorkspaceAssignmentDetailsResponse, error)

	// Lists workspace assignments for a workspace. The response omits the
	// per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`). To read the entitlements for a single principal,
	// get that principal's assignment.
	ListWorkspaceAssignments(ctx context.Context, request ListWorkspaceAssignmentsRequest) (*ListWorkspaceAssignmentsResponse, error)

	// Resolves a group with the given external ID from the customer's IdP. If the
	// group does not exist, it will be created in the account. If the customer is
	// not onboarded onto Automatic Identity Management (AIM), this will return an
	// error.
	ResolveGroup(ctx context.Context, request ResolveGroupRequest) (*ResolveGroupResponse, error)

	// Resolves a service principal with the given external ID from the customer's
	// IdP. If the service principal does not exist, it will be created. If the
	// customer is not onboarded onto Automatic Identity Management (AIM), this will
	// return an error.
	ResolveServicePrincipal(ctx context.Context, request ResolveServicePrincipalRequest) (*ResolveServicePrincipalResponse, error)

	// Resolves a user with the given external ID from the customer's IdP. If the
	// user does not exist, it will be created. If the customer is not onboarded
	// onto Automatic Identity Management (AIM), this will return an error.
	ResolveUser(ctx context.Context, request ResolveUserRequest) (*ResolveUserResponse, error)

	// Updates an existing group in the Databricks account. Only the fields named in
	// the update mask are modified. Returns the updated Group resource.
	UpdateGroup(ctx context.Context, request UpdateGroupRequest) (*Group, error)

	// Updates an existing service principal in the Databricks account. Only the
	// fields named in the update mask are modified. Returns the updated
	// ServicePrincipal resource.
	UpdateServicePrincipal(ctx context.Context, request UpdateServicePrincipalRequest) (*ServicePrincipal, error)

	// Updates an existing user in the Databricks account and returns the updated
	// user. Only the fields named in the update mask are modified. The updatable
	// fields are fullName.givenName, fullName.familyName, status, and externalId.
	// The behavior is the same whether or not Account Identity Management (AIM) is
	// enabled.
	UpdateUser(ctx context.Context, request UpdateUserRequest) (*User, error)

	// Updates the entitlements of a directly assigned principal in a workspace.
	// Changes are applied one at a time rather than atomically. If the request
	// fails partway through, only some of the requested changes take effect. Get
	// the assignment afterwards to confirm the final state.
	UpdateWorkspaceAssignment(ctx context.Context, request UpdateWorkspaceAssignmentRequest) (*WorkspaceAssignment, error)

	// Updates the entitlements of a directly assigned principal in a workspace.
	// Changes are applied one at a time rather than atomically. If the request
	// fails partway through, only some of the requested changes take effect. Get
	// the assignment detail afterwards to confirm the final state.
	UpdateWorkspaceAssignmentDetail(ctx context.Context, request UpdateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)
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

	// Creates a group membership (assigns a principal to a group).
	CreateDirectGroupMemberProxy(ctx context.Context, request CreateDirectGroupMemberProxyRequest) (*DirectGroupMember, error)

	// Creates a local group in the Databricks account that parents the calling
	// workspace and returns the created group. A local group is one that is not
	// synced from the customer's identity provider, and can be created whether or
	// not Account Identity Management (AIM) is enabled.
	//
	// When AIM is enabled, supplying an external ID returns an error. Use the
	// ExternalGroup resource to sync groups from the identity provider instead.
	CreateGroupProxy(ctx context.Context, request CreateGroupProxyRequest) (*Group, error)

	// Creates a local service principal in the Databricks account that parents the
	// calling workspace and returns the created service principal. A local service
	// principal is one that is not synced from the customer's identity provider,
	// and can be created whether or not Account Identity Management (AIM) is
	// enabled.
	//
	// When AIM is enabled, supplying an external ID returns an error. Use the
	// ExternalServicePrincipal resource to sync service principals from the
	// identity provider instead.
	CreateServicePrincipalProxy(ctx context.Context, request CreateServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// Creates a local user in the Databricks account that parents the calling
	// workspace and returns the created user. A local user is one that is not
	// synced from the customer's identity provider, and can be created whether or
	// not Account Identity Management (AIM) is enabled.
	//
	// When AIM is enabled, supplying an external ID returns an error. Use the
	// ExternalUser resource to sync users from the identity provider instead.
	CreateUserProxy(ctx context.Context, request CreateUserProxyRequest) (*User, error)

	// Creates a workspace assignment detail for a principal in the calling
	// workspace. Entitlements are granted one at a time rather than atomically. If
	// the request fails partway through, the principal stays assigned to the
	// workspace with only some of the requested entitlements. Get the assignment
	// detail afterwards to confirm which entitlements were granted.
	CreateWorkspaceAssignmentDetailProxy(ctx context.Context, request CreateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Creates a workspace assignment for a principal in the calling workspace.
	// Entitlements are granted one at a time rather than atomically. If the request
	// fails partway through, the principal stays assigned to the workspace with
	// only some of the requested entitlements. Get the assignment afterwards to
	// confirm which entitlements were granted.
	CreateWorkspaceAssignmentProxy(ctx context.Context, request CreateWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error)

	// Deletes a group membership (unassigns a principal from a group).
	DeleteDirectGroupMemberProxy(ctx context.Context, request DeleteDirectGroupMemberProxyRequest) error

	// Deletes a group by its internal ID from the Databricks account that parents
	// the calling workspace.
	DeleteGroupProxy(ctx context.Context, request DeleteGroupProxyRequest) error

	// Deletes a service principal by its internal ID from the Databricks account
	// that parents the calling workspace.
	DeleteServicePrincipalProxy(ctx context.Context, request DeleteServicePrincipalProxyRequest) error

	// Deletes a user by its internal ID from the Databricks account that parents
	// the calling workspace.
	DeleteUserProxy(ctx context.Context, request DeleteUserProxyRequest) error

	// Deletes a workspace assignment detail for a principal in the calling
	// workspace, revoking all of its entitlements. Entitlements are revoked one at
	// a time rather than atomically. If the request fails partway through, the
	// principal stays assigned with some of its original entitlements. Retrying is
	// safe.
	DeleteWorkspaceAssignmentDetailProxy(ctx context.Context, request DeleteWorkspaceAssignmentDetailProxyRequest) error

	// Deletes a workspace assignment for a principal in the calling workspace,
	// revoking all of its entitlements. Entitlements are revoked one at a time
	// rather than atomically. If the request fails partway through, the principal
	// stays assigned with some of its original entitlements. Retrying is safe.
	DeleteWorkspaceAssignmentProxy(ctx context.Context, request DeleteWorkspaceAssignmentProxyRequest) error

	// Gets a provisioned direct member of a group.
	GetDirectGroupMemberProxy(ctx context.Context, request GetDirectGroupMemberProxyRequest) (*DirectGroupMember, error)

	// Fetches a group by its internal ID from the Databricks account that parents
	// the calling workspace.
	GetGroupProxy(ctx context.Context, request GetGroupProxyRequest) (*Group, error)

	// Fetches a service principal by its internal ID from the Databricks account
	// that parents the calling workspace.
	GetServicePrincipalProxy(ctx context.Context, request GetServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// Fetches a user by its internal ID from the Databricks account that parents
	// the calling workspace.
	GetUserProxy(ctx context.Context, request GetUserProxyRequest) (*User, error)

	// Returns the access details for a principal in the current workspace. Allows
	// for checking access details for any provisioned principal (user, service
	// principal, or group) in the current workspace. * Provisioned principal here
	// refers to one that has been synced into Databricks from the customer's IdP or
	// added explicitly to Databricks via SCIM/UI. Allows for passing in a "view"
	// parameter to control what fields are returned (BASIC by default or FULL).
	GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error)

	// Returns the assignment details for a principal in the calling workspace.
	GetWorkspaceAssignmentDetailProxy(ctx context.Context, request GetWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Returns the assignment for a principal in the calling workspace.
	GetWorkspaceAssignmentProxy(ctx context.Context, request GetWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error)

	// Returns the identity details for a principal in a workspace.
	GetWorkspaceIdentityDetail(ctx context.Context, request GetWorkspaceIdentityDetailRequest) (*WorkspaceIdentityDetail, error)

	// Lists provisioned direct members of a group with their membership source
	// (internal or from identity provider).
	ListDirectGroupMembersProxy(ctx context.Context, request ListDirectGroupMembersProxyRequest) (*ListDirectGroupMembersResponse, error)

	// Lists the groups in the Databricks account that parents the calling
	// workspace, returning one page per call. Supports filtering by group name or
	// external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListGroupsProxy(ctx context.Context, request ListGroupsProxyRequest) listing.Iterator[Group]

	// Lists the groups in the Databricks account that parents the calling
	// workspace, returning one page per call. Supports filtering by group name or
	// external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListGroupsProxyAll(ctx context.Context, request ListGroupsProxyRequest) ([]Group, error)

	// Lists the service principals in the Databricks account that parents the
	// calling workspace, returning one page per call. Supports filtering by
	// application ID or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListServicePrincipalsProxy(ctx context.Context, request ListServicePrincipalsProxyRequest) listing.Iterator[ServicePrincipal]

	// Lists the service principals in the Databricks account that parents the
	// calling workspace, returning one page per call. Supports filtering by
	// application ID or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListServicePrincipalsProxyAll(ctx context.Context, request ListServicePrincipalsProxyRequest) ([]ServicePrincipal, error)

	// Lists all transitive parent groups of a principal.
	ListTransitiveParentGroupsProxy(ctx context.Context, request ListTransitiveParentGroupsProxyRequest) (*ListTransitiveParentGroupsResponse, error)

	// Lists the users in the Databricks account that parents the calling workspace,
	// returning one page per call. Supports filtering by username or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListUsersProxy(ctx context.Context, request ListUsersProxyRequest) listing.Iterator[User]

	// Lists the users in the Databricks account that parents the calling workspace,
	// returning one page per call. Supports filtering by username or external ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListUsersProxyAll(ctx context.Context, request ListUsersProxyRequest) ([]User, error)

	// Lists workspace assignment details for the calling workspace. The response
	// omits the per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`). To read the entitlements for a single principal,
	// get that principal's assignment detail.
	ListWorkspaceAssignmentDetailsProxy(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) (*ListWorkspaceAssignmentDetailsResponse, error)

	// Lists workspace assignments for the calling workspace. The response omits the
	// per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`). To read the entitlements for a single principal,
	// get that principal's assignment.
	ListWorkspaceAssignmentsProxy(ctx context.Context, request ListWorkspaceAssignmentsProxyRequest) (*ListWorkspaceAssignmentsResponse, error)

	// Resolves a group with the given external ID from the customer's IdP. If the
	// group does not exist, it will be created in the account. If the customer is
	// not onboarded onto Automatic Identity Management (AIM), this will return an
	// error.
	ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error)

	// Resolves a service principal with the given external ID from the customer's
	// IdP. If the service principal does not exist, it will be created. If the
	// customer is not onboarded onto Automatic Identity Management (AIM), this will
	// return an error.
	ResolveServicePrincipalProxy(ctx context.Context, request ResolveServicePrincipalProxyRequest) (*ResolveServicePrincipalResponse, error)

	// Resolves a user with the given external ID from the customer's IdP. If the
	// user does not exist, it will be created. If the customer is not onboarded
	// onto Automatic Identity Management (AIM), this will return an error.
	ResolveUserProxy(ctx context.Context, request ResolveUserProxyRequest) (*ResolveUserResponse, error)

	// Updates an existing group in the Databricks account that parents the calling
	// workspace. Only the fields named in the update mask are modified. Returns the
	// updated Group resource.
	UpdateGroupProxy(ctx context.Context, request UpdateGroupProxyRequest) (*Group, error)

	// Updates an existing service principal in the Databricks account that parents
	// the calling workspace. Only the fields named in the update mask are modified.
	// Returns the updated ServicePrincipal resource.
	UpdateServicePrincipalProxy(ctx context.Context, request UpdateServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// Updates an existing user in the Databricks account that parents the calling
	// workspace and returns the updated user. Only the fields named in the update
	// mask are modified. The updatable fields are fullName.givenName,
	// fullName.familyName, status, and externalId.
	UpdateUserProxy(ctx context.Context, request UpdateUserProxyRequest) (*User, error)

	// Updates the entitlements of a directly assigned principal in the calling
	// workspace. Changes are applied one at a time rather than atomically. If the
	// request fails partway through, only some of the requested changes take
	// effect. Get the assignment detail afterwards to confirm the final state.
	UpdateWorkspaceAssignmentDetailProxy(ctx context.Context, request UpdateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Updates the entitlements of a directly assigned principal in the calling
	// workspace. Changes are applied one at a time rather than atomically. If the
	// request fails partway through, only some of the requested changes take
	// effect. Get the assignment afterwards to confirm the final state.
	UpdateWorkspaceAssignmentProxy(ctx context.Context, request UpdateWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error)

	// Updates a workspace identity detail for a principal.
	UpdateWorkspaceIdentityDetail(ctx context.Context, request UpdateWorkspaceIdentityDetailRequest) (*WorkspaceIdentityDetail, error)
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
