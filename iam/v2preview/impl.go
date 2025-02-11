// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iampreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just AccessControlPreview API methods
type accessControlPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *accessControlPreviewImpl) CheckPolicy(ctx context.Context, request CheckPolicyRequest) (*CheckPolicyResponse, error) {
	var checkPolicyResponse CheckPolicyResponse
	path := "/api/2.0preview/access-control/check-policy-v2"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &checkPolicyResponse)
	return &checkPolicyResponse, err
}

// unexported type that holds implementations of just AccountAccessControlPreview API methods
type accountAccessControlPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlPreviewImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := fmt.Sprintf("/api/2.0preview/preview/accounts/%v/access-control/assignable-roles", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlPreviewImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0preview/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlPreviewImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0preview/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountAccessControlProxyPreview API methods
type accountAccessControlProxyPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlProxyPreviewImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := "/api/2.0preview/preview/accounts/access-control/assignable-roles"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlProxyPreviewImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := "/api/2.0preview/preview/accounts/access-control/rule-sets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlProxyPreviewImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := "/api/2.0preview/preview/accounts/access-control/rule-sets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountGroupsPreview API methods
type accountGroupsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *accountGroupsPreviewImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *accountGroupsPreviewImpl) Delete(ctx context.Context, request DeleteAccountGroupRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountGroupsPreviewImpl) Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

// List group details.
//
// Gets all details of the groups associated with the Databricks account.
func (a *accountGroupsPreviewImpl) List(ctx context.Context, request ListAccountGroupsRequest) listing.Iterator[Group] {

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
func (a *accountGroupsPreviewImpl) ListAll(ctx context.Context, request ListAccountGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}
func (a *accountGroupsPreviewImpl) internalList(ctx context.Context, request ListAccountGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *accountGroupsPreviewImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *accountGroupsPreviewImpl) Update(ctx context.Context, request Group) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just AccountServicePrincipalsPreview API methods
type accountServicePrincipalsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *accountServicePrincipalsPreviewImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountServicePrincipalsPreviewImpl) Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountServicePrincipalsPreviewImpl) Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsPreviewImpl) List(ctx context.Context, request ListAccountServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

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
func (a *accountServicePrincipalsPreviewImpl) ListAll(ctx context.Context, request ListAccountServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}
func (a *accountServicePrincipalsPreviewImpl) internalList(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *accountServicePrincipalsPreviewImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *accountServicePrincipalsPreviewImpl) Update(ctx context.Context, request ServicePrincipal) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just AccountUsersPreview API methods
type accountUsersPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *accountUsersPreviewImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *accountUsersPreviewImpl) Delete(ctx context.Context, request DeleteAccountUserRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountUsersPreviewImpl) Get(ctx context.Context, request GetAccountUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

// List users.
//
// Gets details for all the users associated with a Databricks account.
func (a *accountUsersPreviewImpl) List(ctx context.Context, request ListAccountUsersRequest) listing.Iterator[User] {

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
func (a *accountUsersPreviewImpl) ListAll(ctx context.Context, request ListAccountUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}
func (a *accountUsersPreviewImpl) internalList(ctx context.Context, request ListAccountUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *accountUsersPreviewImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *accountUsersPreviewImpl) Update(ctx context.Context, request User) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just CurrentUserPreview API methods
type currentUserPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *currentUserPreviewImpl) Me(ctx context.Context) (*User, error) {
	var user User
	path := "/api/2.0preview/preview/scim/v2/Me"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &user)
	return &user, err
}

// unexported type that holds implementations of just GroupsPreview API methods
type groupsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *groupsPreviewImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0preview/preview/scim/v2/Groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *groupsPreviewImpl) Delete(ctx context.Context, request DeleteGroupRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *groupsPreviewImpl) Get(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

// List group details.
//
// Gets all details of the groups associated with the Databricks workspace.
func (a *groupsPreviewImpl) List(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group] {

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
func (a *groupsPreviewImpl) ListAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}
func (a *groupsPreviewImpl) internalList(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0preview/preview/scim/v2/Groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *groupsPreviewImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *groupsPreviewImpl) Update(ctx context.Context, request Group) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just PermissionMigrationPreview API methods
type permissionMigrationPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *permissionMigrationPreviewImpl) MigratePermissions(ctx context.Context, request MigratePermissionsRequest) (*MigratePermissionsResponse, error) {
	var migratePermissionsResponse MigratePermissionsResponse
	path := "/api/2.0preview/permissionmigration"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &migratePermissionsResponse)
	return &migratePermissionsResponse, err
}

// unexported type that holds implementations of just PermissionsPreview API methods
type permissionsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsPreviewImpl) Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsPreviewImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	var getPermissionLevelsResponse GetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0preview/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPermissionLevelsResponse)
	return &getPermissionLevelsResponse, err
}

func (a *permissionsPreviewImpl) Set(ctx context.Context, request PermissionsRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsPreviewImpl) Update(ctx context.Context, request PermissionsRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

// unexported type that holds implementations of just ServicePrincipalsPreview API methods
type servicePrincipalsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalsPreviewImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0preview/preview/scim/v2/ServicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *servicePrincipalsPreviewImpl) Delete(ctx context.Context, request DeleteServicePrincipalRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *servicePrincipalsPreviewImpl) Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks workspace.
func (a *servicePrincipalsPreviewImpl) List(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

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
func (a *servicePrincipalsPreviewImpl) ListAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}
func (a *servicePrincipalsPreviewImpl) internalList(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0preview/preview/scim/v2/ServicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *servicePrincipalsPreviewImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *servicePrincipalsPreviewImpl) Update(ctx context.Context, request ServicePrincipal) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just UsersPreview API methods
type usersPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *usersPreviewImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := "/api/2.0preview/preview/scim/v2/Users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *usersPreviewImpl) Delete(ctx context.Context, request DeleteUserRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *usersPreviewImpl) Get(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *usersPreviewImpl) GetPermissionLevels(ctx context.Context) (*GetPasswordPermissionLevelsResponse, error) {
	var getPasswordPermissionLevelsResponse GetPasswordPermissionLevelsResponse
	path := "/api/2.0preview/permissions/authorization/passwords/permissionLevels"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getPasswordPermissionLevelsResponse)
	return &getPasswordPermissionLevelsResponse, err
}

func (a *usersPreviewImpl) GetPermissions(ctx context.Context) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0preview/permissions/authorization/passwords"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &passwordPermissions)
	return &passwordPermissions, err
}

// List users.
//
// Gets details for all the users associated with a Databricks workspace.
func (a *usersPreviewImpl) List(ctx context.Context, request ListUsersRequest) listing.Iterator[User] {

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
func (a *usersPreviewImpl) ListAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}
func (a *usersPreviewImpl) internalList(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0preview/preview/scim/v2/Users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *usersPreviewImpl) Patch(ctx context.Context, request PartialUpdate) error {
	var patchResponse PatchResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchResponse)
	return err
}

func (a *usersPreviewImpl) SetPermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0preview/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &passwordPermissions)
	return &passwordPermissions, err
}

func (a *usersPreviewImpl) Update(ctx context.Context, request User) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0preview/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

func (a *usersPreviewImpl) UpdatePermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0preview/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &passwordPermissions)
	return &passwordPermissions, err
}

// unexported type that holds implementations of just WorkspaceAssignmentPreview API methods
type workspaceAssignmentPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceAssignmentPreviewImpl) Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	var deleteWorkspacePermissionAssignmentResponse DeleteWorkspacePermissionAssignmentResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteWorkspacePermissionAssignmentResponse)
	return err
}

func (a *workspaceAssignmentPreviewImpl) Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error) {
	var workspacePermissions WorkspacePermissions
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.ConfiguredAccountID(), request.WorkspaceId)
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
func (a *workspaceAssignmentPreviewImpl) List(ctx context.Context, request ListWorkspaceAssignmentRequest) listing.Iterator[PermissionAssignment] {

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
func (a *workspaceAssignmentPreviewImpl) ListAll(ctx context.Context, request ListWorkspaceAssignmentRequest) ([]PermissionAssignment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PermissionAssignment](ctx, iterator)
}
func (a *workspaceAssignmentPreviewImpl) internalList(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
	var permissionAssignments PermissionAssignments
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces/%v/permissionassignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &permissionAssignments)
	return &permissionAssignments, err
}

func (a *workspaceAssignmentPreviewImpl) Update(ctx context.Context, request UpdateWorkspaceAssignments) (*PermissionAssignment, error) {
	var permissionAssignment PermissionAssignment
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &permissionAssignment)
	return &permissionAssignment, err
}
