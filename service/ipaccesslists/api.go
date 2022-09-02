// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewIpAccessLists(client *client.DatabricksClient) IpAccessListsService {
	return &IpAccessListsAPI{
		client: client,
	}
}

type IpAccessListsAPI struct {
	client *client.DatabricksClient
}

// Add an IP access list for this workspace. A list can be an allow list or a
// block list. See the top of this file for a description of how the server
// treats allow lists and block lists at run time. When creating/updating an IP
// access list: * For all allow lists and block lists combined, the API supports
// a maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
// Attempts to exceed that number return error 400 with `error_code` value
// `QUOTA_EXCEEDED`. * If the new list would block the calling user?s current
// IP, error 400 is returned with `error_code` value `INVALID_STATE`. It can
// take a few minutes for the changes to take effect. Note that your new IP
// access list has no effect until you enable the feature. See
// [`/workspace-conf`](#operation/set-status)
func (a *IpAccessListsAPI) CreateIpAccessList(ctx context.Context, request CreateIPAccessListRequest) (*CreateIPAccessListResponse, error) {
	var createIPAccessListResponse CreateIPAccessListResponse
	path := "/api/2.0/ip-access-lists"
	err := a.client.Post(ctx, path, request, &createIPAccessListResponse)
	return &createIPAccessListResponse, err
}

// Delete an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) DeleteIpAccessList(ctx context.Context, request DeleteIpAccessListRequest) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) FetchIpAccessList(ctx context.Context, request FetchIpAccessListRequest) (*CreateIPAccessListResponse, error) {
	var createIPAccessListResponse CreateIPAccessListResponse
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Get(ctx, path, request, &createIPAccessListResponse)
	return &createIPAccessListResponse, err
}


func (a *IpAccessListsAPI) GetAllIpAccessLists(ctx context.Context) (*GetIPAccessListResponse, error) {
	var getIPAccessListResponse GetIPAccessListResponse
	path := "/api/2.0/ip-access-lists"
	err := a.client.Get(ctx, path, nil, &getIPAccessListResponse)
	return &getIPAccessListResponse, err
}

// Replace an IP access list, specified by its ID. A list can include allow
// lists and block lists. See the top of this file for a description of how the
// server treats allow lists and block lists at run time. When replacing an IP
// access list: * For all allow lists and block lists combined, the API supports
// a maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
// Attempts to exceed that number return error 400 with `error_code` value
// `QUOTA_EXCEEDED`. * If the resulting list would block the calling user?s
// current IP, error 400 is returned with `error_code` value `INVALID_STATE`. It
// can take a few minutes for the changes to take effect. Note that your
// resulting IP access list has no effect until you enable the feature. See
// [`/workspace-conf`](#operation/set-status)
func (a *IpAccessListsAPI) ReplaceIpAccessList(ctx context.Context, request ReplaceIPAccessListRequest) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Put(ctx, path, request)
	return err
}

// Modify an existing IP access list, specified by its ID. A list can include
// allow lists and block lists. See the top of this file for a description of
// how the server treats allow lists and block lists at run time. When updating
// an IP access list: * For all allow lists and block lists combined, the API
// supports a maximum of 1000?IP/CIDR values, where one CIDR counts as a single
// value. Attempts to exceed that number return error 400 with `error_code`
// value `QUOTA_EXCEEDED`. * If the updated list would block the calling user?s
// current IP, error 400 is returned with `error_code` value `INVALID_STATE`. It
// can take a few minutes for the changes to take effect. Note that your
// resulting IP access list has no effect until you enable the feature. See
// [`/workspace-conf`](#operation/set-status)
func (a *IpAccessListsAPI) UpdateIpAccessList(ctx context.Context, request UpdateIPAccessListRequest) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Patch(ctx, path, request)
	return err
}


func (a *IpAccessListsAPI) DeleteIpAccessListByIpAccessListId(ctx context.Context, ipAccessListId string) error {
	return a.DeleteIpAccessList(ctx, DeleteIpAccessListRequest{
		IpAccessListId: ipAccessListId,
	})
}

func (a *IpAccessListsAPI) FetchIpAccessListByIpAccessListId(ctx context.Context, ipAccessListId string) (*CreateIPAccessListResponse, error) {
	return a.FetchIpAccessList(ctx, FetchIpAccessListRequest{
		IpAccessListId: ipAccessListId,
	})
}
