// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/iam/iampb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccessControl API methods
type accessControlImpl struct {
	client *client.DatabricksClient
}

func (a *accessControlImpl) CheckPolicy(ctx context.Context, request CheckPolicyRequest) (*CheckPolicyResponse, error) {
	requestPb, pbErr := CheckPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var checkPolicyResponsePb iampb.CheckPolicyResponsePb
	path := "/api/2.0/access-control/check-policy-v2"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&checkPolicyResponsePb,
	)
	resp, err := CheckPolicyResponseFromPb(&checkPolicyResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just AccountAccessControl API methods
type accountAccessControlImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	requestPb, pbErr := GetAssignableRolesForResourceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getAssignableRolesForResourceResponsePb iampb.GetAssignableRolesForResourceResponsePb
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/assignable-roles", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getAssignableRolesForResourceResponsePb,
	)
	resp, err := GetAssignableRolesForResourceResponseFromPb(&getAssignableRolesForResourceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountAccessControlImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	requestPb, pbErr := GetRuleSetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var ruleSetResponsePb iampb.RuleSetResponsePb
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&ruleSetResponsePb,
	)
	resp, err := RuleSetResponseFromPb(&ruleSetResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountAccessControlImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	requestPb, pbErr := UpdateRuleSetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var ruleSetResponsePb iampb.RuleSetResponsePb
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&ruleSetResponsePb,
	)
	resp, err := RuleSetResponseFromPb(&ruleSetResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just AccountAccessControlProxy API methods
type accountAccessControlProxyImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlProxyImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	requestPb, pbErr := GetAssignableRolesForResourceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getAssignableRolesForResourceResponsePb iampb.GetAssignableRolesForResourceResponsePb
	path := "/api/2.0/preview/accounts/access-control/assignable-roles"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getAssignableRolesForResourceResponsePb,
	)
	resp, err := GetAssignableRolesForResourceResponseFromPb(&getAssignableRolesForResourceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountAccessControlProxyImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	requestPb, pbErr := GetRuleSetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var ruleSetResponsePb iampb.RuleSetResponsePb
	path := "/api/2.0/preview/accounts/access-control/rule-sets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&ruleSetResponsePb,
	)
	resp, err := RuleSetResponseFromPb(&ruleSetResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountAccessControlProxyImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	requestPb, pbErr := UpdateRuleSetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var ruleSetResponsePb iampb.RuleSetResponsePb
	path := "/api/2.0/preview/accounts/access-control/rule-sets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&ruleSetResponsePb,
	)
	resp, err := RuleSetResponseFromPb(&ruleSetResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just AccountGroups API methods
type accountGroupsImpl struct {
	client *client.DatabricksClient
}

func (a *accountGroupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	requestPb, pbErr := GroupToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var groupPb iampb.GroupPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&groupPb,
	)
	resp, err := GroupFromPb(&groupPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountGroupsImpl) Delete(ctx context.Context, request DeleteAccountGroupRequest) error {
	requestPb, pbErr := DeleteAccountGroupRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountGroupsImpl) Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error) {
	requestPb, pbErr := GetAccountGroupRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var groupPb iampb.GroupPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&groupPb,
	)
	resp, err := GroupFromPb(&groupPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets all details of the groups associated with the Databricks account.
func (a *accountGroupsImpl) List(ctx context.Context, request ListAccountGroupsRequest) listing.Iterator[Group] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
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

// Gets all details of the groups associated with the Databricks account.
func (a *accountGroupsImpl) ListAll(ctx context.Context, request ListAccountGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}

func (a *accountGroupsImpl) internalList(ctx context.Context, request ListAccountGroupsRequest) (*ListGroupsResponse, error) {
	requestPb, pbErr := ListAccountGroupsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listGroupsResponsePb iampb.ListGroupsResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listGroupsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListGroupsResponseFromPb(&listGroupsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountGroupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	requestPb, pbErr := PartialUpdateToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountGroupsImpl) Update(ctx context.Context, request Group) error {
	requestPb, pbErr := GroupToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just AccountServicePrincipals API methods
type accountServicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *accountServicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	requestPb, pbErr := ServicePrincipalToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servicePrincipalPb iampb.ServicePrincipalPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&servicePrincipalPb,
	)
	resp, err := ServicePrincipalFromPb(&servicePrincipalPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountServicePrincipalsImpl) Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error {
	requestPb, pbErr := DeleteAccountServicePrincipalRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountServicePrincipalsImpl) Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error) {
	requestPb, pbErr := GetAccountServicePrincipalRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servicePrincipalPb iampb.ServicePrincipalPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&servicePrincipalPb,
	)
	resp, err := ServicePrincipalFromPb(&servicePrincipalPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsImpl) List(ctx context.Context, request ListAccountServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
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

// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsImpl) ListAll(ctx context.Context, request ListAccountServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}

func (a *accountServicePrincipalsImpl) internalList(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	requestPb, pbErr := ListAccountServicePrincipalsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listServicePrincipalResponsePb iampb.ListServicePrincipalResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listServicePrincipalResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListServicePrincipalResponseFromPb(&listServicePrincipalResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountServicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	requestPb, pbErr := PartialUpdateToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountServicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	requestPb, pbErr := ServicePrincipalToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just AccountUsers API methods
type accountUsersImpl struct {
	client *client.DatabricksClient
}

func (a *accountUsersImpl) Create(ctx context.Context, request User) (*User, error) {
	requestPb, pbErr := UserToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var userPb iampb.UserPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&userPb,
	)
	resp, err := UserFromPb(&userPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountUsersImpl) Delete(ctx context.Context, request DeleteAccountUserRequest) error {
	requestPb, pbErr := DeleteAccountUserRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountUsersImpl) Get(ctx context.Context, request GetAccountUserRequest) (*User, error) {
	requestPb, pbErr := GetAccountUserRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var userPb iampb.UserPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&userPb,
	)
	resp, err := UserFromPb(&userPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets details for all the users associated with a Databricks account.
func (a *accountUsersImpl) List(ctx context.Context, request ListAccountUsersRequest) listing.Iterator[User] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
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

// Gets details for all the users associated with a Databricks account.
func (a *accountUsersImpl) ListAll(ctx context.Context, request ListAccountUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}

func (a *accountUsersImpl) internalList(ctx context.Context, request ListAccountUsersRequest) (*ListUsersResponse, error) {
	requestPb, pbErr := ListAccountUsersRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listUsersResponsePb iampb.ListUsersResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listUsersResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListUsersResponseFromPb(&listUsersResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountUsersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	requestPb, pbErr := PartialUpdateToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountUsersImpl) Update(ctx context.Context, request User) error {
	requestPb, pbErr := UserToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just CurrentUser API methods
type currentUserImpl struct {
	client *client.DatabricksClient
}

func (a *currentUserImpl) Me(ctx context.Context) (*User, error) {

	var userPb iampb.UserPb
	path := "/api/2.0/preview/scim/v2/Me"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&userPb,
	)
	resp, err := UserFromPb(&userPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Groups API methods
type groupsImpl struct {
	client *client.DatabricksClient
}

func (a *groupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	requestPb, pbErr := GroupToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var groupPb iampb.GroupPb
	path := "/api/2.0/preview/scim/v2/Groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&groupPb,
	)
	resp, err := GroupFromPb(&groupPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *groupsImpl) Delete(ctx context.Context, request DeleteGroupRequest) error {
	requestPb, pbErr := DeleteGroupRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *groupsImpl) Get(ctx context.Context, request GetGroupRequest) (*Group, error) {
	requestPb, pbErr := GetGroupRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var groupPb iampb.GroupPb
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&groupPb,
	)
	resp, err := GroupFromPb(&groupPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets all details of the groups associated with the Databricks workspace.
func (a *groupsImpl) List(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
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

// Gets all details of the groups associated with the Databricks workspace.
func (a *groupsImpl) ListAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}

func (a *groupsImpl) internalList(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	requestPb, pbErr := ListGroupsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listGroupsResponsePb iampb.ListGroupsResponsePb
	path := "/api/2.0/preview/scim/v2/Groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listGroupsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListGroupsResponseFromPb(&listGroupsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *groupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	requestPb, pbErr := PartialUpdateToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *groupsImpl) Update(ctx context.Context, request Group) error {
	requestPb, pbErr := GroupToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just PermissionMigration API methods
type permissionMigrationImpl struct {
	client *client.DatabricksClient
}

func (a *permissionMigrationImpl) MigratePermissions(ctx context.Context, request MigratePermissionsRequest) (*MigratePermissionsResponse, error) {
	requestPb, pbErr := MigratePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var migratePermissionsResponsePb iampb.MigratePermissionsResponsePb
	path := "/api/2.0/permissionmigration"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&migratePermissionsResponsePb,
	)
	resp, err := MigratePermissionsResponseFromPb(&migratePermissionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Permissions API methods
type permissionsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsImpl) Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error) {
	requestPb, pbErr := GetPermissionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var objectPermissionsPb iampb.ObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&objectPermissionsPb,
	)
	resp, err := ObjectPermissionsFromPb(&objectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *permissionsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	requestPb, pbErr := GetPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPermissionLevelsResponsePb iampb.GetPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getPermissionLevelsResponsePb,
	)
	resp, err := GetPermissionLevelsResponseFromPb(&getPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *permissionsImpl) Set(ctx context.Context, request SetObjectPermissions) (*ObjectPermissions, error) {
	requestPb, pbErr := SetObjectPermissionsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var objectPermissionsPb iampb.ObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&objectPermissionsPb,
	)
	resp, err := ObjectPermissionsFromPb(&objectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *permissionsImpl) Update(ctx context.Context, request UpdateObjectPermissions) (*ObjectPermissions, error) {
	requestPb, pbErr := UpdateObjectPermissionsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var objectPermissionsPb iampb.ObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&objectPermissionsPb,
	)
	resp, err := ObjectPermissionsFromPb(&objectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ServicePrincipals API methods
type servicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	requestPb, pbErr := ServicePrincipalToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servicePrincipalPb iampb.ServicePrincipalPb
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&servicePrincipalPb,
	)
	resp, err := ServicePrincipalFromPb(&servicePrincipalPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servicePrincipalsImpl) Delete(ctx context.Context, request DeleteServicePrincipalRequest) error {
	requestPb, pbErr := DeleteServicePrincipalRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *servicePrincipalsImpl) Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	requestPb, pbErr := GetServicePrincipalRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servicePrincipalPb iampb.ServicePrincipalPb
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&servicePrincipalPb,
	)
	resp, err := ServicePrincipalFromPb(&servicePrincipalPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets the set of service principals associated with a Databricks workspace.
func (a *servicePrincipalsImpl) List(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
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

// Gets the set of service principals associated with a Databricks workspace.
func (a *servicePrincipalsImpl) ListAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}

func (a *servicePrincipalsImpl) internalList(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	requestPb, pbErr := ListServicePrincipalsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listServicePrincipalResponsePb iampb.ListServicePrincipalResponsePb
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listServicePrincipalResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListServicePrincipalResponseFromPb(&listServicePrincipalResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	requestPb, pbErr := PartialUpdateToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *servicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	requestPb, pbErr := ServicePrincipalToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just Users API methods
type usersImpl struct {
	client *client.DatabricksClient
}

func (a *usersImpl) Create(ctx context.Context, request User) (*User, error) {
	requestPb, pbErr := UserToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var userPb iampb.UserPb
	path := "/api/2.0/preview/scim/v2/Users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&userPb,
	)
	resp, err := UserFromPb(&userPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *usersImpl) Delete(ctx context.Context, request DeleteUserRequest) error {
	requestPb, pbErr := DeleteUserRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *usersImpl) Get(ctx context.Context, request GetUserRequest) (*User, error) {
	requestPb, pbErr := GetUserRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var userPb iampb.UserPb
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&userPb,
	)
	resp, err := UserFromPb(&userPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *usersImpl) GetPermissionLevels(ctx context.Context) (*GetPasswordPermissionLevelsResponse, error) {

	var getPasswordPermissionLevelsResponsePb iampb.GetPasswordPermissionLevelsResponsePb
	path := "/api/2.0/permissions/authorization/passwords/permissionLevels"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&getPasswordPermissionLevelsResponsePb,
	)
	resp, err := GetPasswordPermissionLevelsResponseFromPb(&getPasswordPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *usersImpl) GetPermissions(ctx context.Context) (*PasswordPermissions, error) {

	var passwordPermissionsPb iampb.PasswordPermissionsPb
	path := "/api/2.0/permissions/authorization/passwords"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&passwordPermissionsPb,
	)
	resp, err := PasswordPermissionsFromPb(&passwordPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets details for all the users associated with a Databricks workspace.
func (a *usersImpl) List(ctx context.Context, request ListUsersRequest) listing.Iterator[User] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
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

// Gets details for all the users associated with a Databricks workspace.
func (a *usersImpl) ListAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}

func (a *usersImpl) internalList(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	requestPb, pbErr := ListUsersRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listUsersResponsePb iampb.ListUsersResponsePb
	path := "/api/2.0/preview/scim/v2/Users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listUsersResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListUsersResponseFromPb(&listUsersResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *usersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	requestPb, pbErr := PartialUpdateToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *usersImpl) SetPermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
	requestPb, pbErr := PasswordPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var passwordPermissionsPb iampb.PasswordPermissionsPb
	path := "/api/2.0/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&passwordPermissionsPb,
	)
	resp, err := PasswordPermissionsFromPb(&passwordPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *usersImpl) Update(ctx context.Context, request User) error {
	requestPb, pbErr := UserToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *usersImpl) UpdatePermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
	requestPb, pbErr := PasswordPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var passwordPermissionsPb iampb.PasswordPermissionsPb
	path := "/api/2.0/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&passwordPermissionsPb,
	)
	resp, err := PasswordPermissionsFromPb(&passwordPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just WorkspaceAssignment API methods
type workspaceAssignmentImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceAssignmentImpl) Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	requestPb, pbErr := DeleteWorkspaceAssignmentRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *workspaceAssignmentImpl) Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error) {
	requestPb, pbErr := GetWorkspaceAssignmentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspacePermissionsPb iampb.WorkspacePermissionsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&workspacePermissionsPb,
	)
	resp, err := WorkspacePermissionsFromPb(&workspacePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
func (a *workspaceAssignmentImpl) ListAll(ctx context.Context, request ListWorkspaceAssignmentRequest) ([]PermissionAssignment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PermissionAssignment](ctx, iterator)
}

func (a *workspaceAssignmentImpl) internalList(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
	requestPb, pbErr := ListWorkspaceAssignmentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var permissionAssignmentsPb iampb.PermissionAssignmentsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&permissionAssignmentsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := PermissionAssignmentsFromPb(&permissionAssignmentsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceAssignmentImpl) Update(ctx context.Context, request UpdateWorkspaceAssignments) (*PermissionAssignment, error) {
	requestPb, pbErr := UpdateWorkspaceAssignmentsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var permissionAssignmentPb iampb.PermissionAssignmentPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&permissionAssignmentPb,
	)
	resp, err := PermissionAssignmentFromPb(&permissionAssignmentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
