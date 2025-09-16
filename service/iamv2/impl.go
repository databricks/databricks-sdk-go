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

func (a *accountIamV2Impl) CreateGroup(ctx context.Context, request CreateGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *accountIamV2Impl) CreateServicePrincipal(ctx context.Context, request CreateServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/servicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountIamV2Impl) CreateUser(ctx context.Context, request CreateUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *accountIamV2Impl) CreateWorkspaceAccessDetail(ctx context.Context, request CreateWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAccessDetails", a.client.ConfiguredAccountID(), request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAccessDetail, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *accountIamV2Impl) DeleteGroup(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/servicePrincipals/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteUser(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteWorkspaceAccessDetail(ctx context.Context, request DeleteWorkspaceAccessDetailRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAccessDetails/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *accountIamV2Impl) GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/servicePrincipals/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountIamV2Impl) GetUser(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
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

func (a *accountIamV2Impl) ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *accountIamV2Impl) ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error) {
	var listServicePrincipalsResponse ListServicePrincipalsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/servicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalsResponse)
	return &listServicePrincipalsResponse, err
}

func (a *accountIamV2Impl) ListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *accountIamV2Impl) ListWorkspaceAccessDetails(ctx context.Context, request ListWorkspaceAccessDetailsRequest) (*ListWorkspaceAccessDetailsResponse, error) {
	var listWorkspaceAccessDetailsResponse ListWorkspaceAccessDetailsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAccessDetails", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAccessDetailsResponse)
	return &listWorkspaceAccessDetailsResponse, err
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

func (a *accountIamV2Impl) UpdateGroup(ctx context.Context, request UpdateGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *accountIamV2Impl) UpdateServicePrincipal(ctx context.Context, request UpdateServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/servicePrincipals/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountIamV2Impl) UpdateUser(ctx context.Context, request UpdateUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/%v", a.client.ConfiguredAccountID(), request.InternalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *accountIamV2Impl) UpdateWorkspaceAccessDetail(ctx context.Context, request UpdateWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspaceAccessDetails/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAccessDetail, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

// unexported type that holds implementations of just WorkspaceIamV2 API methods
type workspaceIamV2Impl struct {
	client *client.DatabricksClient
}

func (a *workspaceIamV2Impl) CreateGroupProxy(ctx context.Context, request CreateGroupProxyRequest) (*Group, error) {
	var group Group
	path := "/api/2.0/identity/groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *workspaceIamV2Impl) CreateServicePrincipalProxy(ctx context.Context, request CreateServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/identity/servicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *workspaceIamV2Impl) CreateUserProxy(ctx context.Context, request CreateUserProxyRequest) (*User, error) {
	var user User
	path := "/api/2.0/identity/users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *workspaceIamV2Impl) CreateWorkspaceAccessDetailLocal(ctx context.Context, request CreateWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := "/api/2.0/identity/workspaceAccessDetails"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAccessDetail, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *workspaceIamV2Impl) DeleteGroupProxy(ctx context.Context, request DeleteGroupProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/groups/%v", request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) DeleteServicePrincipalProxy(ctx context.Context, request DeleteServicePrincipalProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/servicePrincipals/%v", request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) DeleteUserProxy(ctx context.Context, request DeleteUserProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/users/%v", request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) DeleteWorkspaceAccessDetailLocal(ctx context.Context, request DeleteWorkspaceAccessDetailLocalRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/workspaceAccessDetails/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) GetGroupProxy(ctx context.Context, request GetGroupProxyRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/groups/%v", request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *workspaceIamV2Impl) GetServicePrincipalProxy(ctx context.Context, request GetServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/servicePrincipals/%v", request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *workspaceIamV2Impl) GetUserProxy(ctx context.Context, request GetUserProxyRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/users/%v", request.InternalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
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

func (a *workspaceIamV2Impl) ListGroupsProxy(ctx context.Context, request ListGroupsProxyRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/identity/groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *workspaceIamV2Impl) ListServicePrincipalsProxy(ctx context.Context, request ListServicePrincipalsProxyRequest) (*ListServicePrincipalsResponse, error) {
	var listServicePrincipalsResponse ListServicePrincipalsResponse
	path := "/api/2.0/identity/servicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalsResponse)
	return &listServicePrincipalsResponse, err
}

func (a *workspaceIamV2Impl) ListUsersProxy(ctx context.Context, request ListUsersProxyRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/identity/users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *workspaceIamV2Impl) ListWorkspaceAccessDetailsLocal(ctx context.Context, request ListWorkspaceAccessDetailsLocalRequest) (*ListWorkspaceAccessDetailsResponse, error) {
	var listWorkspaceAccessDetailsResponse ListWorkspaceAccessDetailsResponse
	path := "/api/2.0/identity/workspaceAccessDetails"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAccessDetailsResponse)
	return &listWorkspaceAccessDetailsResponse, err
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

func (a *workspaceIamV2Impl) UpdateGroupProxy(ctx context.Context, request UpdateGroupProxyRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/groups/%v", request.InternalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *workspaceIamV2Impl) UpdateServicePrincipalProxy(ctx context.Context, request UpdateServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/servicePrincipals/%v", request.InternalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *workspaceIamV2Impl) UpdateUserProxy(ctx context.Context, request UpdateUserProxyRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/users/%v", request.InternalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *workspaceIamV2Impl) UpdateWorkspaceAccessDetailLocal(ctx context.Context, request UpdateWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/workspaceAccessDetails/%v", request.PrincipalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAccessDetail, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}
