// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

import (
	"context"
)

type IpAccessListsService interface {
	// Add an IP access list for this workspace. A list can be an allow list or
	// a block list. See the top of this file for a description of how the
	// server treats allow lists and block lists at run time. When
	// creating/updating an IP access list: * For all allow lists and block
	// lists combined, the API supports a maximum of 1000 IP/CIDR values, where
	// one CIDR counts as a single value. Attempts to exceed that number return
	// error 400 with `error_code` value `QUOTA_EXCEEDED`. * If the new list
	// would block the calling user?s current IP, error 400 is returned with
	// `error_code` value `INVALID_STATE`. It can take a few minutes for the
	// changes to take effect. Note that your new IP access list has no effect
	// until you enable the feature. See
	// [`/workspace-conf`](#operation/set-status)
	CreateIpAccessList(ctx context.Context, createIPAccessListRequest CreateIPAccessListRequest) (*CreateIPAccessListResponse, error)
	// Delete an IP access list, specified by its list ID.
	DeleteIpAccessList(ctx context.Context, deleteIpAccessListRequest DeleteIpAccessListRequest) error
	DeleteIpAccessListByIpAccessListId(ctx context.Context, ipAccessListId string) error
	// Get an IP access list, specified by its list ID.
	FetchIpAccessList(ctx context.Context, fetchIpAccessListRequest FetchIpAccessListRequest) (*CreateIPAccessListResponse, error)
	FetchIpAccessListByIpAccessListId(ctx context.Context, ipAccessListId string) (*CreateIPAccessListResponse, error)

	GetAllIpAccessLists(ctx context.Context) (*GetIPAccessListResponse, error)
	// Replace an IP access list, specified by its ID. A list can include allow
	// lists and block lists. See the top of this file for a description of how
	// the server treats allow lists and block lists at run time. When replacing
	// an IP access list: * For all allow lists and block lists combined, the
	// API supports a maximum of 1000 IP/CIDR values, where one CIDR counts as a
	// single value. Attempts to exceed that number return error 400 with
	// `error_code` value `QUOTA_EXCEEDED`. * If the resulting list would block
	// the calling user?s current IP, error 400 is returned with `error_code`
	// value `INVALID_STATE`. It can take a few minutes for the changes to take
	// effect. Note that your resulting IP access list has no effect until you
	// enable the feature. See [`/workspace-conf`](#operation/set-status)
	ReplaceIpAccessList(ctx context.Context, replaceIPAccessListRequest ReplaceIPAccessListRequest) error
	// Modify an existing IP access list, specified by its ID. A list can
	// include allow lists and block lists. See the top of this file for a
	// description of how the server treats allow lists and block lists at run
	// time. When updating an IP access list: * For all allow lists and block
	// lists combined, the API supports a maximum of 1000?IP/CIDR values, where
	// one CIDR counts as a single value. Attempts to exceed that number return
	// error 400 with `error_code` value `QUOTA_EXCEEDED`. * If the updated list
	// would block the calling user?s current IP, error 400 is returned with
	// `error_code` value `INVALID_STATE`. It can take a few minutes for the
	// changes to take effect. Note that your resulting IP access list has no
	// effect until you enable the feature. See
	// [`/workspace-conf`](#operation/set-status)
	UpdateIpAccessList(ctx context.Context, updateIPAccessListRequest UpdateIPAccessListRequest) error
}
