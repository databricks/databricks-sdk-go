// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just IpAccessLists API methods
type ipAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *ipAccessListsImpl) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {
	var createIpAccessListResponse CreateIpAccessListResponse
	path := "/api/2.0/ip-access-lists"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createIpAccessListResponse)
	return &createIpAccessListResponse, err
}

func (a *ipAccessListsImpl) Delete(ctx context.Context, request DeleteRequest) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *ipAccessListsImpl) Get(ctx context.Context, request GetRequest) (*FetchIpAccessListResponse, error) {
	var fetchIpAccessListResponse FetchIpAccessListResponse
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &fetchIpAccessListResponse)
	return &fetchIpAccessListResponse, err
}

func (a *ipAccessListsImpl) List(ctx context.Context) (*GetIpAccessListResponse, error) {
	var getIpAccessListResponse GetIpAccessListResponse
	path := "/api/2.0/ip-access-lists"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getIpAccessListResponse)
	return &getIpAccessListResponse, err
}

func (a *ipAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *ipAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}
