// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccountIamV2 API methods
type accountIamV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountIamV2Impl) CreateDirectGroupMember(ctx context.Context, request CreateDirectGroupMemberRequest) (*DirectGroupMember, error) {
	var directGroupMember DirectGroupMember
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v/direct-members", a.client.ConfiguredAccountID(), request.GroupId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DirectGroupMember, &directGroupMember)
	return &directGroupMember, err
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
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/service-principals", a.client.ConfiguredAccountID())
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

func (a *accountIamV2Impl) CreateWorkspaceAssignment(ctx context.Context, request CreateWorkspaceAssignmentRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignment, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *accountIamV2Impl) CreateWorkspaceAssignmentDetail(ctx context.Context, request CreateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *accountIamV2Impl) DeleteDirectGroupMember(ctx context.Context, request DeleteDirectGroupMemberRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v/direct-members/%v", a.client.ConfiguredAccountID(), request.GroupId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteGroup(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v", a.client.ConfiguredAccountID(), request.GroupId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/service-principals/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteUser(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/%v", a.client.ConfiguredAccountID(), request.UserId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteWorkspaceAssignment(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignments/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteWorkspaceAssignmentDetail(ctx context.Context, request DeleteWorkspaceAssignmentDetailRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) GetDirectGroupMember(ctx context.Context, request GetDirectGroupMemberRequest) (*DirectGroupMember, error) {
	var directGroupMember DirectGroupMember
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v/direct-members/%v", a.client.ConfiguredAccountID(), request.GroupId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &directGroupMember)
	return &directGroupMember, err
}

func (a *accountIamV2Impl) GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v", a.client.ConfiguredAccountID(), request.GroupId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *accountIamV2Impl) GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/service-principals/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountIamV2Impl) GetUser(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/%v", a.client.ConfiguredAccountID(), request.UserId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *accountIamV2Impl) GetWorkspaceAccessDetail(ctx context.Context, request GetWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-access-details/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *accountIamV2Impl) GetWorkspaceAssignment(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignments/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *accountIamV2Impl) GetWorkspaceAssignmentDetail(ctx context.Context, request GetWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *accountIamV2Impl) ListDirectGroupMembers(ctx context.Context, request ListDirectGroupMembersRequest) (*ListDirectGroupMembersResponse, error) {
	var listDirectGroupMembersResponse ListDirectGroupMembersResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v/direct-members", a.client.ConfiguredAccountID(), request.GroupId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDirectGroupMembersResponse)
	return &listDirectGroupMembersResponse, err
}

// Lists the groups in the Databricks account, returning one page per call.
// Supports filtering by group name or external ID.
func (a *accountIamV2Impl) ListGroups(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group] {

	getNextPage := func(ctx context.Context, req ListGroupsRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListGroups(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Groups
	}
	getNextReq := func(resp *ListGroupsResponse) *ListGroupsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the groups in the Databricks account, returning one page per call.
// Supports filtering by group name or external ID.
func (a *accountIamV2Impl) ListGroupsAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	iterator := a.ListGroups(ctx, request)
	return listing.ToSlice[Group](ctx, iterator)
}

func (a *accountIamV2Impl) internalListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

// Lists the service principals in the Databricks account, returning one page
// per call. Supports filtering by application ID or external ID.
func (a *accountIamV2Impl) ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListServicePrincipals(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalsResponse) []ServicePrincipal {
		return resp.ServicePrincipals
	}
	getNextReq := func(resp *ListServicePrincipalsResponse) *ListServicePrincipalsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the service principals in the Databricks account, returning one page
// per call. Supports filtering by application ID or external ID.
func (a *accountIamV2Impl) ListServicePrincipalsAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.ListServicePrincipals(ctx, request)
	return listing.ToSlice[ServicePrincipal](ctx, iterator)
}

func (a *accountIamV2Impl) internalListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error) {
	var listServicePrincipalsResponse ListServicePrincipalsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/service-principals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalsResponse)
	return &listServicePrincipalsResponse, err
}

func (a *accountIamV2Impl) ListTransitiveParentGroups(ctx context.Context, request ListTransitiveParentGroupsRequest) (*ListTransitiveParentGroupsResponse, error) {
	var listTransitiveParentGroupsResponse ListTransitiveParentGroupsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/principals/%v/transitive-parent-groups", a.client.ConfiguredAccountID(), request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTransitiveParentGroupsResponse)
	return &listTransitiveParentGroupsResponse, err
}

// Lists the users in the Databricks account, returning one page per call.
// Supports filtering by username or external ID.
func (a *accountIamV2Impl) ListUsers(ctx context.Context, request ListUsersRequest) listing.Iterator[User] {

	getNextPage := func(ctx context.Context, req ListUsersRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListUsers(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Users
	}
	getNextReq := func(resp *ListUsersResponse) *ListUsersRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the users in the Databricks account, returning one page per call.
// Supports filtering by username or external ID.
func (a *accountIamV2Impl) ListUsersAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	iterator := a.ListUsers(ctx, request)
	return listing.ToSlice[User](ctx, iterator)
}

func (a *accountIamV2Impl) internalListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *accountIamV2Impl) ListWorkspaceAssignmentDetails(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) (*ListWorkspaceAssignmentDetailsResponse, error) {
	var listWorkspaceAssignmentDetailsResponse ListWorkspaceAssignmentDetailsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentDetailsResponse)
	return &listWorkspaceAssignmentDetailsResponse, err
}

func (a *accountIamV2Impl) ListWorkspaceAssignments(ctx context.Context, request ListWorkspaceAssignmentsRequest) (*ListWorkspaceAssignmentsResponse, error) {
	var listWorkspaceAssignmentsResponse ListWorkspaceAssignmentsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentsResponse)
	return &listWorkspaceAssignmentsResponse, err
}

func (a *accountIamV2Impl) ResolveGroup(ctx context.Context, request ResolveGroupRequest) (*ResolveGroupResponse, error) {
	var resolveGroupResponse ResolveGroupResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/resolve-by-external-id", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveGroupResponse)
	return &resolveGroupResponse, err
}

func (a *accountIamV2Impl) ResolveServicePrincipal(ctx context.Context, request ResolveServicePrincipalRequest) (*ResolveServicePrincipalResponse, error) {
	var resolveServicePrincipalResponse ResolveServicePrincipalResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/service-principals/resolve-by-external-id", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveServicePrincipalResponse)
	return &resolveServicePrincipalResponse, err
}

func (a *accountIamV2Impl) ResolveUser(ctx context.Context, request ResolveUserRequest) (*ResolveUserResponse, error) {
	var resolveUserResponse ResolveUserResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/resolve-by-external-id", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveUserResponse)
	return &resolveUserResponse, err
}

func (a *accountIamV2Impl) UpdateGroup(ctx context.Context, request UpdateGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/%v", a.client.ConfiguredAccountID(), request.GroupId)
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
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/service-principals/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
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
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/%v", a.client.ConfiguredAccountID(), request.UserId)
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

func (a *accountIamV2Impl) UpdateWorkspaceAssignment(ctx context.Context, request UpdateWorkspaceAssignmentRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignments/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignment, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *accountIamV2Impl) UpdateWorkspaceAssignmentDetail(ctx context.Context, request UpdateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
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

func (a *workspaceIamV2Impl) CreateDirectGroupMemberProxy(ctx context.Context, request CreateDirectGroupMemberProxyRequest) (*DirectGroupMember, error) {
	var directGroupMember DirectGroupMember
	path := fmt.Sprintf("/api/2.0/identity/groups/%v/direct-members", request.GroupId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DirectGroupMember, &directGroupMember)
	return &directGroupMember, err
}

func (a *workspaceIamV2Impl) CreateGroupProxy(ctx context.Context, request CreateGroupProxyRequest) (*Group, error) {
	var group Group
	path := "/api/2.0/identity/groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *workspaceIamV2Impl) CreateServicePrincipalProxy(ctx context.Context, request CreateServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/identity/service-principals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
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
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *workspaceIamV2Impl) CreateWorkspaceAssignmentDetailProxy(ctx context.Context, request CreateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := "/api/2.0/identity/workspace-assignment-details"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *workspaceIamV2Impl) CreateWorkspaceAssignmentProxy(ctx context.Context, request CreateWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := "/api/2.0/identity/workspace-assignments"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignment, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *workspaceIamV2Impl) DeleteDirectGroupMemberProxy(ctx context.Context, request DeleteDirectGroupMemberProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/groups/%v/direct-members/%v", request.GroupId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) DeleteGroupProxy(ctx context.Context, request DeleteGroupProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/groups/%v", request.GroupId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) DeleteServicePrincipalProxy(ctx context.Context, request DeleteServicePrincipalProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/service-principals/%v", request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) DeleteUserProxy(ctx context.Context, request DeleteUserProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/users/%v", request.UserId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) DeleteWorkspaceAssignmentDetailProxy(ctx context.Context, request DeleteWorkspaceAssignmentDetailProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignment-details/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) DeleteWorkspaceAssignmentProxy(ctx context.Context, request DeleteWorkspaceAssignmentProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignments/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceIamV2Impl) GetDirectGroupMemberProxy(ctx context.Context, request GetDirectGroupMemberProxyRequest) (*DirectGroupMember, error) {
	var directGroupMember DirectGroupMember
	path := fmt.Sprintf("/api/2.0/identity/groups/%v/direct-members/%v", request.GroupId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &directGroupMember)
	return &directGroupMember, err
}

func (a *workspaceIamV2Impl) GetGroupProxy(ctx context.Context, request GetGroupProxyRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/groups/%v", request.GroupId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *workspaceIamV2Impl) GetServicePrincipalProxy(ctx context.Context, request GetServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/service-principals/%v", request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *workspaceIamV2Impl) GetUserProxy(ctx context.Context, request GetUserProxyRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/users/%v", request.UserId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *workspaceIamV2Impl) GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/workspace-access-details/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *workspaceIamV2Impl) GetWorkspaceAssignmentDetailProxy(ctx context.Context, request GetWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignment-details/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *workspaceIamV2Impl) GetWorkspaceAssignmentProxy(ctx context.Context, request GetWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignments/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *workspaceIamV2Impl) GetWorkspaceIdentityDetail(ctx context.Context, request GetWorkspaceIdentityDetailRequest) (*WorkspaceIdentityDetail, error) {
	var workspaceIdentityDetail WorkspaceIdentityDetail
	path := fmt.Sprintf("/api/2.0/identity/workspace-identity-details/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceIdentityDetail)
	return &workspaceIdentityDetail, err
}

func (a *workspaceIamV2Impl) ListDirectGroupMembersProxy(ctx context.Context, request ListDirectGroupMembersProxyRequest) (*ListDirectGroupMembersResponse, error) {
	var listDirectGroupMembersResponse ListDirectGroupMembersResponse
	path := fmt.Sprintf("/api/2.0/identity/groups/%v/direct-members", request.GroupId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDirectGroupMembersResponse)
	return &listDirectGroupMembersResponse, err
}

// Lists the groups in the Databricks account that parents the calling
// workspace, returning one page per call. Supports filtering by group name or
// external ID.
func (a *workspaceIamV2Impl) ListGroupsProxy(ctx context.Context, request ListGroupsProxyRequest) listing.Iterator[Group] {

	getNextPage := func(ctx context.Context, req ListGroupsProxyRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListGroupsProxy(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Groups
	}
	getNextReq := func(resp *ListGroupsResponse) *ListGroupsProxyRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the groups in the Databricks account that parents the calling
// workspace, returning one page per call. Supports filtering by group name or
// external ID.
func (a *workspaceIamV2Impl) ListGroupsProxyAll(ctx context.Context, request ListGroupsProxyRequest) ([]Group, error) {
	iterator := a.ListGroupsProxy(ctx, request)
	return listing.ToSlice[Group](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListGroupsProxy(ctx context.Context, request ListGroupsProxyRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/identity/groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

// Lists the service principals in the Databricks account that parents the
// calling workspace, returning one page per call. Supports filtering by
// application ID or external ID.
func (a *workspaceIamV2Impl) ListServicePrincipalsProxy(ctx context.Context, request ListServicePrincipalsProxyRequest) listing.Iterator[ServicePrincipal] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalsProxyRequest) (*ListServicePrincipalsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListServicePrincipalsProxy(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalsResponse) []ServicePrincipal {
		return resp.ServicePrincipals
	}
	getNextReq := func(resp *ListServicePrincipalsResponse) *ListServicePrincipalsProxyRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the service principals in the Databricks account that parents the
// calling workspace, returning one page per call. Supports filtering by
// application ID or external ID.
func (a *workspaceIamV2Impl) ListServicePrincipalsProxyAll(ctx context.Context, request ListServicePrincipalsProxyRequest) ([]ServicePrincipal, error) {
	iterator := a.ListServicePrincipalsProxy(ctx, request)
	return listing.ToSlice[ServicePrincipal](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListServicePrincipalsProxy(ctx context.Context, request ListServicePrincipalsProxyRequest) (*ListServicePrincipalsResponse, error) {
	var listServicePrincipalsResponse ListServicePrincipalsResponse
	path := "/api/2.0/identity/service-principals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalsResponse)
	return &listServicePrincipalsResponse, err
}

func (a *workspaceIamV2Impl) ListTransitiveParentGroupsProxy(ctx context.Context, request ListTransitiveParentGroupsProxyRequest) (*ListTransitiveParentGroupsResponse, error) {
	var listTransitiveParentGroupsResponse ListTransitiveParentGroupsResponse
	path := fmt.Sprintf("/api/2.0/identity/principals/%v/transitive-parent-groups", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTransitiveParentGroupsResponse)
	return &listTransitiveParentGroupsResponse, err
}

// Lists the users in the Databricks account that parents the calling workspace,
// returning one page per call. Supports filtering by username or external ID.
func (a *workspaceIamV2Impl) ListUsersProxy(ctx context.Context, request ListUsersProxyRequest) listing.Iterator[User] {

	getNextPage := func(ctx context.Context, req ListUsersProxyRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListUsersProxy(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Users
	}
	getNextReq := func(resp *ListUsersResponse) *ListUsersProxyRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the users in the Databricks account that parents the calling workspace,
// returning one page per call. Supports filtering by username or external ID.
func (a *workspaceIamV2Impl) ListUsersProxyAll(ctx context.Context, request ListUsersProxyRequest) ([]User, error) {
	iterator := a.ListUsersProxy(ctx, request)
	return listing.ToSlice[User](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListUsersProxy(ctx context.Context, request ListUsersProxyRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/identity/users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *workspaceIamV2Impl) ListWorkspaceAssignmentDetailsProxy(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) (*ListWorkspaceAssignmentDetailsResponse, error) {
	var listWorkspaceAssignmentDetailsResponse ListWorkspaceAssignmentDetailsResponse
	path := "/api/2.0/identity/workspace-assignment-details"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentDetailsResponse)
	return &listWorkspaceAssignmentDetailsResponse, err
}

func (a *workspaceIamV2Impl) ListWorkspaceAssignmentsProxy(ctx context.Context, request ListWorkspaceAssignmentsProxyRequest) (*ListWorkspaceAssignmentsResponse, error) {
	var listWorkspaceAssignmentsResponse ListWorkspaceAssignmentsResponse
	path := "/api/2.0/identity/workspace-assignments"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentsResponse)
	return &listWorkspaceAssignmentsResponse, err
}

func (a *workspaceIamV2Impl) ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error) {
	var resolveGroupResponse ResolveGroupResponse
	path := "/api/2.0/identity/groups/resolve-by-external-id"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveGroupResponse)
	return &resolveGroupResponse, err
}

func (a *workspaceIamV2Impl) ResolveServicePrincipalProxy(ctx context.Context, request ResolveServicePrincipalProxyRequest) (*ResolveServicePrincipalResponse, error) {
	var resolveServicePrincipalResponse ResolveServicePrincipalResponse
	path := "/api/2.0/identity/service-principals/resolve-by-external-id"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveServicePrincipalResponse)
	return &resolveServicePrincipalResponse, err
}

func (a *workspaceIamV2Impl) ResolveUserProxy(ctx context.Context, request ResolveUserProxyRequest) (*ResolveUserResponse, error) {
	var resolveUserResponse ResolveUserResponse
	path := "/api/2.0/identity/users/resolve-by-external-id"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveUserResponse)
	return &resolveUserResponse, err
}

func (a *workspaceIamV2Impl) UpdateGroupProxy(ctx context.Context, request UpdateGroupProxyRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/groups/%v", request.GroupId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *workspaceIamV2Impl) UpdateServicePrincipalProxy(ctx context.Context, request UpdateServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/service-principals/%v", request.ServicePrincipalId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *workspaceIamV2Impl) UpdateUserProxy(ctx context.Context, request UpdateUserProxyRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/users/%v", request.UserId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *workspaceIamV2Impl) UpdateWorkspaceAssignmentDetailProxy(ctx context.Context, request UpdateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignment-details/%v", request.PrincipalId)
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
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *workspaceIamV2Impl) UpdateWorkspaceAssignmentProxy(ctx context.Context, request UpdateWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignments/%v", request.PrincipalId)
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
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignment, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *workspaceIamV2Impl) UpdateWorkspaceIdentityDetail(ctx context.Context, request UpdateWorkspaceIdentityDetailRequest) (*WorkspaceIdentityDetail, error) {
	var workspaceIdentityDetail WorkspaceIdentityDetail
	path := fmt.Sprintf("/api/2.0/identity/workspace-identity-details/%v", request.PrincipalId)
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
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceIdentityDetail, &workspaceIdentityDetail)
	return &workspaceIdentityDetail, err
}
