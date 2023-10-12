// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"context"
)

// The Accounts IP Access List API enables account admins to configure IP access
// lists for access to the account console.
//
// Account IP Access Lists affect web application access and REST API access to
// the account console and account APIs. If the feature is disabled for the
// account, all access is allowed for this account. There is support for allow
// lists (inclusion) and block lists (exclusion).
//
// When a connection is attempted: 1. **First, all block lists are checked.** If
// the connection IP address matches any block list, the connection is rejected.
// 2. **If the connection was not rejected by block lists**, the IP address is
// compared with the allow lists.
//
// If there is at least one allow list for the account, the connection is
// allowed only if the IP address matches an allow list. If there are no allow
// lists for the account, all IP addresses are allowed.
//
// For all allow lists and block lists combined, the account supports a maximum
// of 1000 IP/CIDR values, where one CIDR counts as a single value.
//
// After changes to the account-level IP access lists, it can take a few minutes
// for changes to take effect.
type AccountIpAccessListsService interface {

	// Create access list.
	//
	// Creates an IP access list for the account.
	//
	// A list can be an allow list or a block list. See the top of this file for
	// a description of how the server treats allow lists and block lists at
	// runtime.
	//
	// When creating or updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the new list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect.
	Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error)

	// Delete access list.
	//
	// Deletes an IP access list, specified by its list ID.
	Delete(ctx context.Context, request DeleteAccountIpAccessListRequest) error

	// Get IP access list.
	//
	// Gets an IP access list, specified by its list ID.
	Get(ctx context.Context, request GetAccountIpAccessListRequest) (*GetIpAccessListResponse, error)

	// Get access lists.
	//
	// Gets all IP access lists for the specified account.
	//
	// Use ListAll() to get all IpAccessListInfo instances
	List(ctx context.Context) (*GetIpAccessListsResponse, error)

	// Replace access list.
	//
	// Replaces an IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time. When replacing an IP access list: * For all allow lists and
	// block lists combined, the API supports a maximum of 1000 IP/CIDR values,
	// where one CIDR counts as a single value. Attempts to exceed that number
	// return error 400 with `error_code` value `QUOTA_EXCEEDED`. * If the
	// resulting list would block the calling user's current IP, error 400 is
	// returned with `error_code` value `INVALID_STATE`. It can take a few
	// minutes for the changes to take effect.
	Replace(ctx context.Context, request ReplaceIpAccessList) error

	// Update access list.
	//
	// Updates an existing IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time.
	//
	// When updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the updated list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect.
	Update(ctx context.Context, request UpdateIpAccessList) error
}

// The Personal Compute enablement setting lets you control which users can use
// the Personal Compute default policy to create compute resources. By default
// all users in all workspaces have access (ON), but you can change the setting
// to instead let individual workspaces configure access control (DELEGATE).
//
// There is only one instance of this setting per account. Since this setting
// has a default value, this setting is present on all accounts even though it's
// never set on a given account. Deletion reverts the value of the setting back
// to the default value.
type AccountSettingsService interface {

	// Delete Personal Compute setting.
	//
	// Reverts back the Personal Compute setting value to default (ON)
	DeletePersonalComputeSetting(ctx context.Context, request DeletePersonalComputeSettingRequest) (*DeletePersonalComputeSettingResponse, error)

	// Get Personal Compute setting.
	//
	// Gets the value of the Personal Compute setting.
	ReadPersonalComputeSetting(ctx context.Context, request ReadPersonalComputeSettingRequest) (*PersonalComputeSetting, error)

	// Update Personal Compute setting.
	//
	// Updates the value of the Personal Compute setting.
	UpdatePersonalComputeSetting(ctx context.Context, request UpdatePersonalComputeSettingRequest) (*PersonalComputeSetting, error)
}

// Credentials manager interacts with with Identity Providers to to perform
// token exchanges using stored credentials and refresh tokens.
type CredentialsManagerService interface {

	// Exchange token.
	//
	// Exchange tokens with an Identity Provider to get a new access token. It
	// allowes specifying scopes to determine token permissions.
	ExchangeToken(ctx context.Context, request ExchangeTokenRequest) (*ExchangeTokenResponse, error)
}

// IP Access List enables admins to configure IP access lists.
//
// IP access lists affect web application access and REST API access to this
// workspace only. If the feature is disabled for a workspace, all access is
// allowed for this workspace. There is support for allow lists (inclusion) and
// block lists (exclusion).
//
// When a connection is attempted: 1. **First, all block lists are checked.** If
// the connection IP address matches any block list, the connection is rejected.
// 2. **If the connection was not rejected by block lists**, the IP address is
// compared with the allow lists.
//
// If there is at least one allow list for the workspace, the connection is
// allowed only if the IP address matches an allow list. If there are no allow
// lists for the workspace, all IP addresses are allowed.
//
// For all allow lists and block lists combined, the workspace supports a
// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
//
// After changes to the IP access list feature, it can take a few minutes for
// changes to take effect.
type IpAccessListsService interface {

	// Create access list.
	//
	// Creates an IP access list for this workspace.
	//
	// A list can be an allow list or a block list. See the top of this file for
	// a description of how the server treats allow lists and block lists at
	// runtime.
	//
	// When creating or updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the new list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect. **Note**: Your
	// new IP access list has no effect until you enable the feature. See
	// :method:workspaceconf/setStatus
	Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error)

	// Delete access list.
	//
	// Deletes an IP access list, specified by its list ID.
	Delete(ctx context.Context, request DeleteIpAccessListRequest) error

	// Get access list.
	//
	// Gets an IP access list, specified by its list ID.
	Get(ctx context.Context, request GetIpAccessListRequest) (*FetchIpAccessListResponse, error)

	// Get access lists.
	//
	// Gets all IP access lists for the specified workspace.
	//
	// Use ListAll() to get all IpAccessListInfo instances
	List(ctx context.Context) (*ListIpAccessListResponse, error)

	// Replace access list.
	//
	// Replaces an IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time. When replacing an IP access list: * For all allow lists and
	// block lists combined, the API supports a maximum of 1000 IP/CIDR values,
	// where one CIDR counts as a single value. Attempts to exceed that number
	// return error 400 with `error_code` value `QUOTA_EXCEEDED`. * If the
	// resulting list would block the calling user's current IP, error 400 is
	// returned with `error_code` value `INVALID_STATE`. It can take a few
	// minutes for the changes to take effect. Note that your resulting IP
	// access list has no effect until you enable the feature. See
	// :method:workspaceconf/setStatus.
	Replace(ctx context.Context, request ReplaceIpAccessList) error

	// Update access list.
	//
	// Updates an existing IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time.
	//
	// When updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the updated list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect. Note that your
	// resulting IP access list has no effect until you enable the feature. See
	// :method:workspaceconf/setStatus.
	Update(ctx context.Context, request UpdateIpAccessList) error
}

// // TODO(yuyuan.tang) to add the description for the setting
type SettingsService interface {

	// Delete the default namespace.
	//
	// Deletes the default namespace.
	DeleteDefaultWorkspaceNamespace(ctx context.Context, request DeleteDefaultWorkspaceNamespaceRequest) (*DeleteDefaultWorkspaceNamespaceResponse, error)

	// Get the default namespace.
	//
	// Gets the default namespace.
	ReadDefaultWorkspaceNamespace(ctx context.Context, request ReadDefaultWorkspaceNamespaceRequest) (*DefaultNamespaceSetting, error)

	// Updates the default namespace setting.
	//
	// Updates the default namespace setting for the workspace. A fresh etag
	// needs to be provided in PATCH requests (as part the setting field). The
	// etag can be retrieved by making a GET request before the PATCH request.
	// Note that if the setting does not exist, GET will return a NOT_FOUND
	// error and the etag will be present in the error response, which should be
	// set in the PATCH request.
	UpdateDefaultWorkspaceNamespace(ctx context.Context, request UpdateDefaultWorkspaceNamespaceRequest) (*DefaultNamespaceSetting, error)
}

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
type TokenManagementService interface {

	// Create on-behalf token.
	//
	// Creates a token on behalf of a service principal.
	CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error)

	// Delete a token.
	//
	// Deletes a token, specified by its ID.
	Delete(ctx context.Context, request DeleteTokenManagementRequest) error

	// Get token info.
	//
	// Gets information about a token, specified by its ID.
	Get(ctx context.Context, request GetTokenManagementRequest) (*TokenInfo, error)

	// Get token permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context) (*GetTokenPermissionLevelsResponse, error)

	// Get token permissions.
	//
	// Gets the permissions of all tokens. Tokens can inherit permissions from
	// their root object.
	GetPermissions(ctx context.Context) (*TokenPermissions, error)

	// List all tokens.
	//
	// Lists all tokens associated with the specified workspace or user.
	//
	// Use ListAll() to get all TokenInfo instances
	List(ctx context.Context, request ListTokenManagementRequest) (*ListTokensResponse, error)

	// Set token permissions.
	//
	// Sets permissions on all tokens. Tokens can inherit permissions from their
	// root object.
	SetPermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error)

	// Update token permissions.
	//
	// Updates the permissions on all tokens. Tokens can inherit permissions
	// from their root object.
	UpdatePermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error)
}

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
type TokensService interface {

	// Create a user token.
	//
	// Creates and returns a token for a user. If this call is made through
	// token authentication, it creates a token with the same client ID as the
	// authenticated token. If the user's token quota is exceeded, this call
	// returns an error **QUOTA_EXCEEDED**.
	Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error)

	// Revoke token.
	//
	// Revokes an access token.
	//
	// If a token with the specified ID is not valid, this call returns an error
	// **RESOURCE_DOES_NOT_EXIST**.
	Delete(ctx context.Context, request RevokeTokenRequest) error

	// List tokens.
	//
	// Lists all the valid tokens for a user-workspace pair.
	//
	// Use ListAll() to get all TokenInfo instances
	List(ctx context.Context) (*ListTokensResponse, error)
}

// This API allows updating known workspace settings for advanced users.
type WorkspaceConfService interface {

	// Check configuration status.
	//
	// Gets the configuration status for a workspace.
	GetStatus(ctx context.Context, request GetStatusRequest) (*map[string]string, error)

	// Enable/disable features.
	//
	// Sets the configuration status for a workspace, including enabling or
	// disabling it.
	SetStatus(ctx context.Context, request WorkspaceConf) error
}
