// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AccountAccessControl API methods
type accountAccessControlImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/assignable-roles", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPut, path, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountAccessControlProxy API methods
type accountAccessControlProxyImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlProxyImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := "/api/2.0/preview/accounts/access-control/assignable-roles"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlProxyImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := "/api/2.0/preview/accounts/access-control/rule-sets"
	err := a.client.Do(ctx, http.MethodGet, path, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlProxyImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := "/api/2.0/preview/accounts/access-control/rule-sets"
	err := a.client.Do(ctx, http.MethodPut, path, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountGroups API methods
type accountGroupsImpl struct {
	client *client.DatabricksClient
}

func (a *accountGroupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &group)
	return &group, err
}

func (a *accountGroupsImpl) Delete(ctx context.Context, request DeleteAccountGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountGroupsImpl) Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &group)
	return &group, err
}

func (a *accountGroupsImpl) List(ctx context.Context, request ListAccountGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *accountGroupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *accountGroupsImpl) Update(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just AccountServicePrincipals API methods
type accountServicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *accountServicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountServicePrincipalsImpl) Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountServicePrincipalsImpl) Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountServicePrincipalsImpl) List(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *accountServicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *accountServicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just AccountUsers API methods
type accountUsersImpl struct {
	client *client.DatabricksClient
}

func (a *accountUsersImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &user)
	return &user, err
}

func (a *accountUsersImpl) Delete(ctx context.Context, request DeleteAccountUserRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountUsersImpl) Get(ctx context.Context, request GetAccountUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &user)
	return &user, err
}

func (a *accountUsersImpl) List(ctx context.Context, request ListAccountUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *accountUsersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *accountUsersImpl) Update(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just CurrentUser API methods
type currentUserImpl struct {
	client *client.DatabricksClient
}

func (a *currentUserImpl) Me(ctx context.Context) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Me"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &user)
	return &user, err
}

// unexported type that holds implementations of just Groups API methods
type groupsImpl struct {
	client *client.DatabricksClient
}

func (a *groupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Do(ctx, http.MethodPost, path, request, &group)
	return &group, err
}

func (a *groupsImpl) Delete(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *groupsImpl) Get(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &group)
	return &group, err
}

func (a *groupsImpl) List(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *groupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *groupsImpl) Update(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just Permissions API methods
type permissionsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsImpl) Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	var getPermissionLevelsResponse GetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getPermissionLevelsResponse)
	return &getPermissionLevelsResponse, err
}

func (a *permissionsImpl) Set(ctx context.Context, request PermissionsRequest) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *permissionsImpl) Update(ctx context.Context, request PermissionsRequest) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just ServicePrincipals API methods
type servicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Do(ctx, http.MethodPost, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *servicePrincipalsImpl) Delete(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *servicePrincipalsImpl) Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *servicePrincipalsImpl) List(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *servicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *servicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just Users API methods
type usersImpl struct {
	client *client.DatabricksClient
}

func (a *usersImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Do(ctx, http.MethodPost, path, request, &user)
	return &user, err
}

func (a *usersImpl) Delete(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *usersImpl) Get(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &user)
	return &user, err
}

func (a *usersImpl) List(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *usersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *usersImpl) Update(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just WorkspaceAssignment API methods
type workspaceAssignmentImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceAssignmentImpl) Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *workspaceAssignmentImpl) Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error) {
	var workspacePermissions WorkspacePermissions
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &workspacePermissions)
	return &workspacePermissions, err
}

func (a *workspaceAssignmentImpl) List(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
	var permissionAssignments PermissionAssignments
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &permissionAssignments)
	return &permissionAssignments, err
}

func (a *workspaceAssignmentImpl) Update(ctx context.Context, request UpdateWorkspaceAssignments) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}
