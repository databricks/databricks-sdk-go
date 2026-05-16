// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AccountIamV2 API methods
type accountIamV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountIamV2Impl) CreateWorkspaceAssignmentDetail(ctx context.Context, request CreateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAssignmentDetails", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *accountIamV2Impl) DeleteWorkspaceAssignmentDetail(ctx context.Context, request DeleteWorkspaceAssignmentDetailRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAssignmentDetails/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
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

func (a *accountIamV2Impl) GetWorkspaceAssignmentDetail(ctx context.Context, request GetWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAssignmentDetails/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *accountIamV2Impl) ListWorkspaceAssignmentDetails(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) (*ListWorkspaceAssignmentDetailsResponse, error) {
	var listWorkspaceAssignmentDetailsResponse ListWorkspaceAssignmentDetailsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAssignmentDetails", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentDetailsResponse)
	return &listWorkspaceAssignmentDetailsResponse, err
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

func (a *accountIamV2Impl) UpdateWorkspaceAssignmentDetail(ctx context.Context, request UpdateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAssignmentDetails/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

// unexported type that holds implementations of just WorkspaceIamV2 API methods
type workspaceIamV2Impl struct {
	client *client.DatabricksClient
}

func (a *workspaceIamV2Impl) CreateWorkspaceAssignmentDetailProxy(ctx context.Context, request CreateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := "/api/2.0/identity/workspaceAssignmentDetails"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *workspaceIamV2Impl) DeleteWorkspaceAssignmentDetailProxy(ctx context.Context, request DeleteWorkspaceAssignmentDetailProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/workspaceAssignmentDetails/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/workspaceAccessDetails/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *workspaceIamV2Impl) GetWorkspaceAssignmentDetailProxy(ctx context.Context, request GetWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/workspaceAssignmentDetails/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *workspaceIamV2Impl) ListWorkspaceAssignmentDetailsProxy(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) (*ListWorkspaceAssignmentDetailsResponse, error) {
	var listWorkspaceAssignmentDetailsResponse ListWorkspaceAssignmentDetailsResponse
	path := "/api/2.0/identity/workspaceAssignmentDetails"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentDetailsResponse)
	return &listWorkspaceAssignmentDetailsResponse, err
}

func (a *workspaceIamV2Impl) ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error) {
	var resolveGroupResponse ResolveGroupResponse
	path := "/api/2.0/identity/groups/resolveByExternalId"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
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
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
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
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveUserResponse)
	return &resolveUserResponse, err
}

func (a *workspaceIamV2Impl) UpdateWorkspaceAssignmentDetailProxy(ctx context.Context, request UpdateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/workspaceAssignmentDetails/%v", request.PrincipalId)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}
