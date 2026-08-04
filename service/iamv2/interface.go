// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"context"
)

// These APIs are used to manage identities and the workspace access of these
// identities in <Databricks>.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountIamV2Service interface {

	// Creates a workspace assignment detail for a principal. Entitlements are
	// granted one at a time rather than atomically. If the request fails
	// partway through, the principal stays assigned to the workspace with only
	// some of the requested entitlements. Get the assignment detail afterwards
	// to confirm which entitlements were granted.
	CreateWorkspaceAssignmentDetail(ctx context.Context, request CreateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)

	// Deletes a workspace assignment detail for a principal, revoking all of
	// its entitlements. Entitlements are revoked one at a time rather than
	// atomically. If the request fails partway through, the principal stays
	// assigned with some of its original entitlements. Retrying is safe.
	DeleteWorkspaceAssignmentDetail(ctx context.Context, request DeleteWorkspaceAssignmentDetailRequest) error

	// Returns the access details for a principal in a workspace. Allows for
	// checking access details for any provisioned principal (user, service
	// principal, or group) in a workspace. * Provisioned principal here refers
	// to one that has been synced into Databricks from the customer's IdP or
	// added explicitly to Databricks via SCIM/UI. Allows for passing in a
	// "view" parameter to control what fields are returned (BASIC by default or
	// FULL).
	GetWorkspaceAccessDetail(ctx context.Context, request GetWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error)

	// Returns the assignment details for a principal in a workspace.
	GetWorkspaceAssignmentDetail(ctx context.Context, request GetWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)

	// Lists workspace assignment details for a workspace. The response omits
	// the per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`). To read the entitlements for a single
	// principal, get that principal's assignment detail.
	ListWorkspaceAssignmentDetails(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) (*ListWorkspaceAssignmentDetailsResponse, error)

	// Resolves a group with the given external ID from the customer's IdP. If
	// the group does not exist, it will be created in the account. If the
	// customer is not onboarded onto Automatic Identity Management (AIM), this
	// will return an error.
	ResolveGroup(ctx context.Context, request ResolveGroupRequest) (*ResolveGroupResponse, error)

	// Resolves a service principal with the given external ID from the
	// customer's IdP. If the service principal does not exist, it will be
	// created. If the customer is not onboarded onto Automatic Identity
	// Management (AIM), this will return an error.
	ResolveServicePrincipal(ctx context.Context, request ResolveServicePrincipalRequest) (*ResolveServicePrincipalResponse, error)

	// Resolves a user with the given external ID from the customer's IdP. If
	// the user does not exist, it will be created. If the customer is not
	// onboarded onto Automatic Identity Management (AIM), this will return an
	// error.
	ResolveUser(ctx context.Context, request ResolveUserRequest) (*ResolveUserResponse, error)

	// Updates the entitlements of a directly assigned principal in a workspace.
	// Changes are applied one at a time rather than atomically. If the request
	// fails partway through, only some of the requested changes take effect.
	// Get the assignment detail afterwards to confirm the final state.
	UpdateWorkspaceAssignmentDetail(ctx context.Context, request UpdateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)
}

// These APIs are used to manage identities and the workspace access of these
// identities in <Databricks>.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceIamV2Service interface {

	// Creates a workspace assignment detail for a principal in the calling
	// workspace. Entitlements are granted one at a time rather than atomically.
	// If the request fails partway through, the principal stays assigned to the
	// workspace with only some of the requested entitlements. Get the
	// assignment detail afterwards to confirm which entitlements were granted.
	CreateWorkspaceAssignmentDetailProxy(ctx context.Context, request CreateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Deletes a workspace assignment detail for a principal in the calling
	// workspace, revoking all of its entitlements. Entitlements are revoked one
	// at a time rather than atomically. If the request fails partway through,
	// the principal stays assigned with some of its original entitlements.
	// Retrying is safe.
	DeleteWorkspaceAssignmentDetailProxy(ctx context.Context, request DeleteWorkspaceAssignmentDetailProxyRequest) error

	// Returns the access details for a principal in the current workspace.
	// Allows for checking access details for any provisioned principal (user,
	// service principal, or group) in the current workspace. * Provisioned
	// principal here refers to one that has been synced into Databricks from
	// the customer's IdP or added explicitly to Databricks via SCIM/UI. Allows
	// for passing in a "view" parameter to control what fields are returned
	// (BASIC by default or FULL).
	GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error)

	// Returns the assignment details for a principal in the calling workspace.
	GetWorkspaceAssignmentDetailProxy(ctx context.Context, request GetWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Lists workspace assignment details for the calling workspace. The
	// response omits the per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`). To read the entitlements for a single
	// principal, get that principal's assignment detail.
	ListWorkspaceAssignmentDetailsProxy(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) (*ListWorkspaceAssignmentDetailsResponse, error)

	// Resolves a group with the given external ID from the customer's IdP. If
	// the group does not exist, it will be created in the account. If the
	// customer is not onboarded onto Automatic Identity Management (AIM), this
	// will return an error.
	ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error)

	// Resolves a service principal with the given external ID from the
	// customer's IdP. If the service principal does not exist, it will be
	// created. If the customer is not onboarded onto Automatic Identity
	// Management (AIM), this will return an error.
	ResolveServicePrincipalProxy(ctx context.Context, request ResolveServicePrincipalProxyRequest) (*ResolveServicePrincipalResponse, error)

	// Resolves a user with the given external ID from the customer's IdP. If
	// the user does not exist, it will be created. If the customer is not
	// onboarded onto Automatic Identity Management (AIM), this will return an
	// error.
	ResolveUserProxy(ctx context.Context, request ResolveUserProxyRequest) (*ResolveUserResponse, error)

	// Updates the entitlements of a directly assigned principal in the calling
	// workspace. Changes are applied one at a time rather than atomically. If
	// the request fails partway through, only some of the requested changes
	// take effect. Get the assignment detail afterwards to confirm the final
	// state.
	UpdateWorkspaceAssignmentDetailProxy(ctx context.Context, request UpdateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)
}
