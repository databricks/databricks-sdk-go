// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AccountIamV2 API methods
type accountIamV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountIamV2Impl) GetWorkspaceAccessDetail(ctx context.Context, request GetWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAccessDetails/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *accountIamV2Impl) ResolveGroup(ctx context.Context, request ResolveGroupRequest) (*ResolveGroupResponse, error) {
	var resolveGroupResponse ResolveGroupResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/resolveByExternalId", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveGroupResponse)
	return &resolveGroupResponse, err
}

func (a *accountIamV2Impl) ResolveServicePrincipal(ctx context.Context, request ResolveServicePrincipalRequest) (*ResolveServicePrincipalResponse, error) {
	var resolveServicePrincipalResponse ResolveServicePrincipalResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/servicePrincipals/resolveByExternalId", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveServicePrincipalResponse)
	return &resolveServicePrincipalResponse, err
}

func (a *accountIamV2Impl) ResolveUser(ctx context.Context, request ResolveUserRequest) (*ResolveUserResponse, error) {
	var resolveUserResponse ResolveUserResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/resolveByExternalId", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveUserResponse)
	return &resolveUserResponse, err
}

// unexported type that holds implementations of just WorkspaceIamV2 API methods
type workspaceIamV2Impl struct {
	client *client.DatabricksClient
}

func (a *workspaceIamV2Impl) GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/workspaceAccessDetails/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *workspaceIamV2Impl) ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error) {
	var resolveGroupResponse ResolveGroupResponse
	path := "/api/2.0/identity/groups/resolveByExternalId"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveGroupResponse)
	return &resolveGroupResponse, err
}

func (a *workspaceIamV2Impl) ResolveServicePrincipalProxy(ctx context.Context, request ResolveServicePrincipalProxyRequest) (*ResolveServicePrincipalResponse, error) {
	var resolveServicePrincipalResponse ResolveServicePrincipalResponse
	path := "/api/2.0/identity/servicePrincipals/resolveByExternalId"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveServicePrincipalResponse)
	return &resolveServicePrincipalResponse, err
}

func (a *workspaceIamV2Impl) ResolveUserProxy(ctx context.Context, request ResolveUserProxyRequest) (*ResolveUserResponse, error) {
	var resolveUserResponse ResolveUserResponse
	path := "/api/2.0/identity/users/resolveByExternalId"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveUserResponse)
	return &resolveUserResponse, err
}
