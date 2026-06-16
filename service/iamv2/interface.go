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

	// Creates a workspace assignment detail for a principal. Entitlement grants
	// are applied individually and non-atomically — if a failure occurs
	// partway through, the principal will be assigned to the workspace but with
	// only a subset of the requested entitlements. Use
	// GetWorkspaceAssignmentDetail to confirm which entitlements were
	// successfully granted.
	CreateWorkspaceAssignmentDetail(ctx context.Context, request CreateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)

	// Deletes a workspace assignment detail for a principal, revoking all
	// associated entitlements. Entitlement revocations are applied individually
	// and non-atomically — if a failure occurs partway through, the principal
	// remains assigned with a subset of its original entitlements, and the
	// operation is safe to retry.
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

	// Lists workspace assignment details for a workspace. For scalability, the
	// response omits the per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`); call GetWorkspaceAssignmentDetail to read
	// entitlements for a single principal.
	ListWorkspaceAssignmentDetails(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) (*ListWorkspaceAssignmentDetailsResponse, error)

	// Resolves a group with the given external ID from the customer's IdP. If
	// the group does not exist, it will be created in the account. If the
	// customer is not onboarded onto Automatic Identity Management (AIM), this
	// will return an error.
	ResolveGroup(ctx context.Context, request ResolveGroupRequest) (*ResolveGroupResponse, error)

	// Resolves an SP with the given external ID from the customer's IdP. If the
	// SP does not exist, it will be created. If the customer is not onboarded
	// onto Automatic Identity Management (AIM), this will return an error.
	ResolveServicePrincipal(ctx context.Context, request ResolveServicePrincipalRequest) (*ResolveServicePrincipalResponse, error)

	// Resolves a user with the given external ID from the customer's IdP. If
	// the user does not exist, it will be created. If the customer is not
	// onboarded onto Automatic Identity Management (AIM), this will return an
	// error.
	ResolveUser(ctx context.Context, request ResolveUserRequest) (*ResolveUserResponse, error)

	// Updates the entitlements of a directly assigned principal in a workspace.
	// Entitlement changes are applied individually and non-atomically — if a
	// failure occurs partway through, only a subset of the requested changes
	// may have been applied. Use GetWorkspaceAssignmentDetail to confirm the
	// final state.
	UpdateWorkspaceAssignmentDetail(ctx context.Context, request UpdateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)
}

// These APIs are used to manage identities and the workspace access of these
// identities in <Databricks>.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceIamV2Service interface {

	// Creates a workspace assignment detail for a principal (workspace-level
	// proxy). Entitlement grants are applied individually and non-atomically
	// — if a failure occurs partway through, the principal will be assigned
	// to the workspace but with only a subset of the requested entitlements.
	// Use GetWorkspaceAssignmentDetail to confirm which entitlements were
	// successfully granted.
	CreateWorkspaceAssignmentDetailProxy(ctx context.Context, request CreateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Deletes a workspace assignment detail for a principal (workspace-level
	// proxy), revoking all associated entitlements. Entitlement revocations are
	// applied individually and non-atomically — if a failure occurs partway
	// through, the principal remains assigned with a subset of its original
	// entitlements, and the operation is safe to retry.
	DeleteWorkspaceAssignmentDetailProxy(ctx context.Context, request DeleteWorkspaceAssignmentDetailProxyRequest) error

	// Returns the access details for a principal in the current workspace.
	// Allows for checking access details for any provisioned principal (user,
	// service principal, or group) in the current workspace. * Provisioned
	// principal here refers to one that has been synced into Databricks from
	// the customer's IdP or added explicitly to Databricks via SCIM/UI. Allows
	// for passing in a "view" parameter to control what fields are returned
	// (BASIC by default or FULL).
	GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error)

	// Returns the assignment details for a principal in a workspace
	// (workspace-level proxy).
	GetWorkspaceAssignmentDetailProxy(ctx context.Context, request GetWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Lists workspace assignment details for a workspace (workspace-level
	// proxy). For scalability, the response omits the per-principal entitlement
	// fields (`entitlements` and `effective_entitlements`); call
	// GetWorkspaceAssignmentDetailProxy to read entitlements for a single
	// principal.
	ListWorkspaceAssignmentDetailsProxy(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) (*ListWorkspaceAssignmentDetailsResponse, error)

	// Resolves a group with the given external ID from the customer's IdP. If
	// the group does not exist, it will be created in the account. If the
	// customer is not onboarded onto Automatic Identity Management (AIM), this
	// will return an error.
	ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error)

	// Resolves an SP with the given external ID from the customer's IdP. If the
	// SP does not exist, it will be created. If the customer is not onboarded
	// onto Automatic Identity Management (AIM), this will return an error.
	ResolveServicePrincipalProxy(ctx context.Context, request ResolveServicePrincipalProxyRequest) (*ResolveServicePrincipalResponse, error)

	// Resolves a user with the given external ID from the customer's IdP. If
	// the user does not exist, it will be created. If the customer is not
	// onboarded onto Automatic Identity Management (AIM), this will return an
	// error.
	ResolveUserProxy(ctx context.Context, request ResolveUserProxyRequest) (*ResolveUserResponse, error)

	// Updates the entitlements of a directly assigned principal in a workspace
	// (workspace-level proxy). Entitlement changes are applied individually and
	// non-atomically — if a failure occurs partway through, only a subset of
	// the requested changes may have been applied. Use
	// GetWorkspaceAssignmentDetail to confirm the final state.
	UpdateWorkspaceAssignmentDetailProxy(ctx context.Context, request UpdateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)
}
