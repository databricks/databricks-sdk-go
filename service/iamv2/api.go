// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Iam V2, Workspace Iam V2, etc.
package iamv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type AccountIamV2Interface interface {

	// Returns the access details for a principal in a workspace. Allows for
	// checking access details for any provisioned principal (user, service
	// principal, or group) in a workspace. * Provisioned principal here refers to
	// one that has been synced into Databricks from the customer's IdP or added
	// explicitly to Databricks via SCIM/UI. Allows for passing in a "view"
	// parameter to control what fields are returned (BASIC by default or FULL).
	GetWorkspaceAccessDetail(ctx context.Context, request GetWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error)

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

	// Returns the access details for a principal in the current workspace. Allows
	// for checking access details for any provisioned principal (user, service
	// principal, or group) in the current workspace. * Provisioned principal here
	// refers to one that has been synced into Databricks from the customer's IdP or
	// added explicitly to Databricks via SCIM/UI. Allows for passing in a "view"
	// parameter to control what fields are returned (BASIC by default or FULL).
	GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error)

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
