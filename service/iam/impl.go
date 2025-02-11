// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccessControl API methods
type accessControlImpl struct {
	client *client.DatabricksClient
}

func (a *accessControlImpl) CheckPolicy(ctx context.Context, request CheckPolicyRequest) (*CheckPolicyResponse, error) {
	var checkPolicyResponse CheckPolicyResponse
	path := "/api/2.0/access-control/check-policy-v2"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &checkPolicyResponse)
	return &checkPolicyResponse, err
}

// unexported type that holds implementations of just AccountAccessControl API methods
type accountAccessControlImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/assignable-roles", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountAccessControlProxy API methods
type accountAccessControlProxyImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlProxyImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := "/api/2.0/preview/accounts/access-control/assignable-roles"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlProxyImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := "/api/2.0/preview/accounts/access-control/rule-sets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlProxyImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := "/api/2.0/preview/accounts/access-control/rule-sets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountGroups API methods
type accountGroupsImpl struct {
	client *client.DatabricksClient
}

func (a *accountGroupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *accountGroupsImpl) Delete(ctx context.Context, request DeleteAccountGroupRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountGroupsImpl) Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

// List group details.
//
// Gets all details of the groups associated with the Databricks account.
func (a *accountGroupsImpl) List(ctx context.Context, request ListAccountGroupsRequest) listing.Iterator[Group] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListAccountGroupsRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Resources
	}
	getNextReq := func(resp *ListGroupsResponse) *ListAccountGroupsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[Group, string](
		iterator,
		func(item Group) string {
			return item.Id
		})
	return dedupedIterator
}

// List group details.
//
// Gets all details of the groups associated with the Databricks account.
func (a *accountGroupsImpl) ListAll(ctx context.Context, request ListAccountGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}

func (a *accountGroupsImpl) internalList(ctx context.Context, request ListAccountGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *accountGroupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *accountGroupsImpl) Update(ctx context.Context, request Group) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just AccountServicePrincipals API methods
type accountServicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *accountServicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountServicePrincipalsImpl) Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountServicePrincipalsImpl) Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsImpl) List(ctx context.Context, request ListAccountServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalResponse) []ServicePrincipal {
		return resp.Resources
	}
	getNextReq := func(resp *ListServicePrincipalResponse) *ListAccountServicePrincipalsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[ServicePrincipal, string](
		iterator,
		func(item ServicePrincipal) string {
			return item.Id
		})
	return dedupedIterator
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsImpl) ListAll(ctx context.Context, request ListAccountServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}

func (a *accountServicePrincipalsImpl) internalList(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *accountServicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *accountServicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just AccountUsers API methods
type accountUsersImpl struct {
	client *client.DatabricksClient
}

func (a *accountUsersImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *accountUsersImpl) Delete(ctx context.Context, request DeleteAccountUserRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountUsersImpl) Get(ctx context.Context, request GetAccountUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

// List users.
//
// Gets details for all the users associated with a Databricks account.
func (a *accountUsersImpl) List(ctx context.Context, request ListAccountUsersRequest) listing.Iterator[User] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListAccountUsersRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Resources
	}
	getNextReq := func(resp *ListUsersResponse) *ListAccountUsersRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[User, string](
		iterator,
		func(item User) string {
			return item.Id
		})
	return dedupedIterator
}

// List users.
//
// Gets details for all the users associated with a Databricks account.
func (a *accountUsersImpl) ListAll(ctx context.Context, request ListAccountUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}

func (a *accountUsersImpl) internalList(ctx context.Context, request ListAccountUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *accountUsersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *accountUsersImpl) Update(ctx context.Context, request User) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just CurrentUser API methods
type currentUserImpl struct {
	client *client.DatabricksClient
}

func (a *currentUserImpl) Me(ctx context.Context) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Me"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &user)
	return &user, err
}

// unexported type that holds implementations of just Groups API methods
type groupsImpl struct {
	client *client.DatabricksClient
}

func (a *groupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *groupsImpl) Delete(ctx context.Context, request DeleteGroupRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *groupsImpl) Get(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

// List group details.
//
// Gets all details of the groups associated with the Databricks workspace.
func (a *groupsImpl) List(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListGroupsRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Resources
	}
	getNextReq := func(resp *ListGroupsResponse) *ListGroupsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[Group, string](
		iterator,
		func(item Group) string {
			return item.Id
		})
	return dedupedIterator
}

// List group details.
//
// Gets all details of the groups associated with the Databricks workspace.
func (a *groupsImpl) ListAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}

func (a *groupsImpl) internalList(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/preview/scim/v2/Groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *groupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *groupsImpl) Update(ctx context.Context, request Group) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just PermissionMigration API methods
type permissionMigrationImpl struct {
	client *client.DatabricksClient
}

func (a *permissionMigrationImpl) MigratePermissions(ctx context.Context, request MigratePermissionsRequest) (*MigratePermissionsResponse, error) {
	var migratePermissionsResponse MigratePermissionsResponse
	path := "/api/2.0/permissionmigration"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &migratePermissionsResponse)
	return &migratePermissionsResponse, err
}

// unexported type that holds implementations of just Permissions API methods
type permissionsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsImpl) Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	var getPermissionLevelsResponse GetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPermissionLevelsResponse)
	return &getPermissionLevelsResponse, err
}

func (a *permissionsImpl) Set(ctx context.Context, request PermissionsRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsImpl) Update(ctx context.Context, request PermissionsRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

// unexported type that holds implementations of just ServicePrincipals API methods
type servicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *servicePrincipalsImpl) Delete(ctx context.Context, request DeleteServicePrincipalRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *servicePrincipalsImpl) Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks workspace.
func (a *servicePrincipalsImpl) List(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalResponse) []ServicePrincipal {
		return resp.Resources
	}
	getNextReq := func(resp *ListServicePrincipalResponse) *ListServicePrincipalsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[ServicePrincipal, string](
		iterator,
		func(item ServicePrincipal) string {
			return item.Id
		})
	return dedupedIterator
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks workspace.
func (a *servicePrincipalsImpl) ListAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}

func (a *servicePrincipalsImpl) internalList(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *servicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *servicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just Users API methods
type usersImpl struct {
	client *client.DatabricksClient
}

func (a *usersImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *usersImpl) Delete(ctx context.Context, request DeleteUserRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *usersImpl) Get(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *usersImpl) GetPermissionLevels(ctx context.Context) (*GetPasswordPermissionLevelsResponse, error) {
	var getPasswordPermissionLevelsResponse GetPasswordPermissionLevelsResponse
	path := "/api/2.0/permissions/authorization/passwords/permissionLevels"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getPasswordPermissionLevelsResponse)
	return &getPasswordPermissionLevelsResponse, err
}

func (a *usersImpl) GetPermissions(ctx context.Context) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0/permissions/authorization/passwords"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &passwordPermissions)
	return &passwordPermissions, err
}

// List users.
//
// Gets details for all the users associated with a Databricks workspace.
func (a *usersImpl) List(ctx context.Context, request ListUsersRequest) listing.Iterator[User] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListUsersRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Resources
	}
	getNextReq := func(resp *ListUsersResponse) *ListUsersRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[User, string](
		iterator,
		func(item User) string {
			return item.Id
		})
	return dedupedIterator
}

// List users.
//
// Gets details for all the users associated with a Databricks workspace.
func (a *usersImpl) ListAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}

func (a *usersImpl) internalList(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/preview/scim/v2/Users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *usersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *usersImpl) SetPermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &passwordPermissions)
	return &passwordPermissions, err
}

func (a *usersImpl) Update(ctx context.Context, request User) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

func (a *usersImpl) UpdatePermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &passwordPermissions)
	return &passwordPermissions, err
}

// unexported type that holds implementations of just WorkspaceAssignment API methods
type workspaceAssignmentImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceAssignmentImpl) Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	var deleteWorkspacePermissionAssignmentResponse DeleteWorkspacePermissionAssignmentResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteWorkspacePermissionAssignmentResponse)
	return err
}

func (a *workspaceAssignmentImpl) Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error) {
	var workspacePermissions WorkspacePermissions
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspacePermissions)
	return &workspacePermissions, err
}

// Get permission assignments.
//
// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
func (a *workspaceAssignmentImpl) List(ctx context.Context, request ListWorkspaceAssignmentRequest) listing.Iterator[PermissionAssignment] {

	getNextPage := func(ctx context.Context, req ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *PermissionAssignments) []PermissionAssignment {
		return resp.PermissionAssignments
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get permission assignments.
//
// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
func (a *workspaceAssignmentImpl) ListAll(ctx context.Context, request ListWorkspaceAssignmentRequest) ([]PermissionAssignment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PermissionAssignment](ctx, iterator)
}

func (a *workspaceAssignmentImpl) internalList(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
	var permissionAssignments PermissionAssignments
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &permissionAssignments)
	return &permissionAssignments, err
}

func (a *workspaceAssignmentImpl) Update(ctx context.Context, request UpdateWorkspaceAssignments) (*PermissionAssignment, error) {
	var permissionAssignment PermissionAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &permissionAssignment)
	return &permissionAssignment, err
}
