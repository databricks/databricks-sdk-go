// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewIpAccessLists(client *client.DatabricksClient) *IpAccessListsAPI {
	return &IpAccessListsAPI{
		impl: &ipAccessListsImpl{
			client: client,
		},
	}
}

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
type IpAccessListsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(IpAccessListsService)
	impl IpAccessListsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *IpAccessListsAPI) WithImpl(impl IpAccessListsService) *IpAccessListsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level IpAccessLists API implementation
func (a *IpAccessListsAPI) Impl() IpAccessListsService {
	return a.impl
}

// Create access list
//
// Creates an IP access list for this workspace. A list can be an allow list or
// a block list. See the top of this file for a description of how the server
// treats allow lists and block lists at runtime.
//
// When creating or updating an IP access list:
//
// * For all allow lists and block lists combined, the API supports a maximum of
// 1000 IP/CIDR values, where one CIDR counts as a single value. Attempts to
// exceed that number return error 400 with `error_code` value `QUOTA_EXCEEDED`.
// * If the new list would block the calling user's current IP, error 400 is
// returned with `error_code` value `INVALID_STATE`.
//
// It can take a few minutes for the changes to take effect. **Note**: Your new
// IP access list has no effect until you enable the feature. See
// [`/workspace-conf`](#operation/set-status).
func (a *IpAccessListsAPI) CreateIpAccessList(ctx context.Context, request CreateIpAccessListRequest) (*IpAccessListInfo, error) {
	return a.impl.CreateIpAccessList(ctx, request)
}

// Delete access list
//
// Deletes an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) DeleteIpAccessList(ctx context.Context, request DeleteIpAccessListRequest) error {
	return a.impl.DeleteIpAccessList(ctx, request)
}

// Delete access list
//
// Deletes an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) DeleteIpAccessListByIpAccessListId(ctx context.Context, ipAccessListId string) error {
	return a.impl.DeleteIpAccessList(ctx, DeleteIpAccessListRequest{
		IpAccessListId: ipAccessListId,
	})
}

// Get access list
//
// Gets an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) FetchIpAccessList(ctx context.Context, request FetchIpAccessListRequest) (*IpAccessListInfo, error) {
	return a.impl.FetchIpAccessList(ctx, request)
}

// Get access list
//
// Gets an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) FetchIpAccessListByIpAccessListId(ctx context.Context, ipAccessListId string) (*IpAccessListInfo, error) {
	return a.impl.FetchIpAccessList(ctx, FetchIpAccessListRequest{
		IpAccessListId: ipAccessListId,
	})
}

// Get access lists
//
// Gets all IP access lists for the specified workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *IpAccessListsAPI) ListIpAccessListsAll(ctx context.Context) ([]IpAccessListInfo, error) {
	response, err := a.impl.ListIpAccessLists(ctx)
	if err != nil {
		return nil, err
	}
	return response.IpAccessLists, nil
}

// IpAccessListInfoLabelToListIdMap calls [IpAccessListsAPI.ListIpAccessListsAll] and creates a map of results with [IpAccessListInfo].Label as key and [IpAccessListInfo].ListId as value.
//
// Returns an error if there's more than one [IpAccessListInfo] with the same .Label.
//
// Note: All [IpAccessListInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *IpAccessListsAPI) IpAccessListInfoLabelToListIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListIpAccessListsAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Label
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Label: %s", key)
		}
		mapping[key] = v.ListId
	}
	return mapping, nil
}

// GetIpAccessListInfoByLabel calls [IpAccessListsAPI.IpAccessListInfoLabelToListIdMap] and returns a single [IpAccessListInfo].
//
// Returns an error if there's more than one [IpAccessListInfo] with the same .Label.
//
// Note: All [IpAccessListInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *IpAccessListsAPI) GetIpAccessListInfoByLabel(ctx context.Context, name string) (*IpAccessListInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListIpAccessListsAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.Label != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("IpAccessListInfo named '%s' does not exist", name)
}

// Replace access list
//
// Replaces an IP access list, specified by its ID. A list can include allow
// lists and block lists. See the top of this file for a description of how the
// server treats allow lists and block lists at run time. When replacing an IP
// access list: * For all allow lists and block lists combined, the API supports
// a maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
// Attempts to exceed that number return error 400 with `error_code` value
// `QUOTA_EXCEEDED`. * If the resulting list would block the calling user's
// current IP, error 400 is returned with `error_code` value `INVALID_STATE`. It
// can take a few minutes for the changes to take effect. Note that your
// resulting IP access list has no effect until you enable the feature. See
// :method:workspaceconf/setStatus.
func (a *IpAccessListsAPI) ReplaceIpAccessList(ctx context.Context, request ReplaceIpAccessListRequest) error {
	return a.impl.ReplaceIpAccessList(ctx, request)
}

// Update access list
//
// Updates an existing IP access list, specified by its ID. A list can include
// allow lists and block lists. See the top of this file for a description of
// how the server treats allow lists and block lists at run time.
//
// When updating an IP access list:
//
// * For all allow lists and block lists combined, the API supports a maximum of
// 1000 IP/CIDR values, where one CIDR counts as a single value. Attempts to
// exceed that number return error 400 with `error_code` value `QUOTA_EXCEEDED`.
// * If the updated list would block the calling user's current IP, error 400 is
// returned with `error_code` value `INVALID_STATE`.
//
// It can take a few minutes for the changes to take effect. Note that your
// resulting IP access list has no effect until you enable the feature. See
// [`/workspace-conf`](#operation/set-status).
func (a *IpAccessListsAPI) UpdateIpAccessList(ctx context.Context, request UpdateIpAccessListRequest) error {
	return a.impl.UpdateIpAccessList(ctx, request)
}
