// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

import (
	"context"
)

// The IP Access List API enables Databricks admins to configure IP access lists
// for a workspace.
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

	// Create access list
	//
	// Creates an IP access list for this workspace. A list can be an allow list
	// or a block list. See the top of this file for a description of how the
	// server treats allow lists and block lists at runtime.
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
	// [`/workspace-conf`](#operation/set-status).
	Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error)

	// Delete access list
	//
	// Deletes an IP access list, specified by its list ID.
	Delete(ctx context.Context, request Delete) error

	// Get access list
	//
	// Gets an IP access list, specified by its list ID.
	Get(ctx context.Context, request Get) (*FetchIpAccessListResponse, error)

	// Get access lists
	//
	// Gets all IP access lists for the specified workspace.
	//
	// Use ListAll() to get all IpAccessListInfo instances
	List(ctx context.Context) (*GetIpAccessListResponse, error)

	// Replace access list
	//
	// Replaces an IP access list, specified by its ID. A list can include allow
	// lists and block lists. See the top of this file for a description of how
	// the server treats allow lists and block lists at run time. When replacing
	// an IP access list: * For all allow lists and block lists combined, the
	// API supports a maximum of 1000 IP/CIDR values, where one CIDR counts as a
	// single value. Attempts to exceed that number return error 400 with
	// `error_code` value `QUOTA_EXCEEDED`. * If the resulting list would block
	// the calling user's current IP, error 400 is returned with `error_code`
	// value `INVALID_STATE`. It can take a few minutes for the changes to take
	// effect. Note that your resulting IP access list has no effect until you
	// enable the feature. See :method:workspaceconf/setStatus.
	Replace(ctx context.Context, request ReplaceIpAccessList) error

	// Update access list
	//
	// Updates an existing IP access list, specified by its ID. A list can
	// include allow lists and block lists. See the top of this file for a
	// description of how the server treats allow lists and block lists at run
	// time.
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
