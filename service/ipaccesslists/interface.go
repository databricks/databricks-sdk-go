// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

import (
	"context"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
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
	CreateIpAccessList(ctx context.Context, request CreateIPAccessListRequest) (*CreateIPAccessListResponse, error)

	// Delete access list
	//
	// Deletes an IP access list, specified by its list ID.
	DeleteIpAccessList(ctx context.Context, request DeleteIpAccessListRequest) error

	// DeleteIpAccessListByIpAccessListId calls DeleteIpAccessList, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteIpAccessListByIpAccessListId(ctx context.Context, ipAccessListId string) error

	// Get access list
	//
	// Gets an IP access list, specified by its list ID.
	FetchIpAccessList(ctx context.Context, request FetchIpAccessListRequest) (*CreateIPAccessListResponse, error)

	// FetchIpAccessListByIpAccessListId calls FetchIpAccessList, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	FetchIpAccessListByIpAccessListId(ctx context.Context, ipAccessListId string) (*CreateIPAccessListResponse, error)

	// Get access lists
	//
	// Gets all IP access lists for the specified workspace.
	GetAllIpAccessLists(ctx context.Context) (*GetIPAccessListResponse, error)

	// GetAllIpAccessListsAll calls GetAllIpAccessLists() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetAllIpAccessListsAll(ctx context.Context) ([]CreateIPAccessListResponse, error)

	// Replace access list
	//
	// Replaces an IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time.
	//
	// When replacing an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the resulting list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect.
	//
	// Note that your resulting IP access list has no effect until you enable
	// the feature. See [`/workspace-conf`](#operation/set-status).
	ReplaceIpAccessList(ctx context.Context, request ReplaceIPAccessListRequest) error

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
	// [`/workspace-conf`](#operation/set-status).
	UpdateIpAccessList(ctx context.Context, request UpdateIPAccessListRequest) error
}
