// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just AccountFederationPolicy API methods
type accountFederationPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *accountFederationPolicyImpl) Create(ctx context.Context, request CreateAccountFederationPolicyRequest) (*FederationPolicy, error) {

	requestPb, pbErr := createAccountFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb federationPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	if requestPb.PolicyId != "" || slices.Contains(requestPb.ForceSendFields, "PolicyId") {
		queryParams["policy_id"] = requestPb.PolicyId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).Policy,
		&federationPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := federationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountFederationPolicyImpl) Delete(ctx context.Context, request DeleteAccountFederationPolicyRequest) error {

	requestPb, pbErr := deleteAccountFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), requestPb.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *accountFederationPolicyImpl) Get(ctx context.Context, request GetAccountFederationPolicyRequest) (*FederationPolicy, error) {

	requestPb, pbErr := getAccountFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb federationPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), requestPb.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&federationPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := federationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List account federation policies.
func (a *accountFederationPolicyImpl) List(ctx context.Context, request ListAccountFederationPoliciesRequest) listing.Iterator[FederationPolicy] {

	getNextPage := func(ctx context.Context, req ListAccountFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFederationPoliciesResponse) []FederationPolicy {
		return resp.Policies
	}
	getNextReq := func(resp *ListFederationPoliciesResponse) *ListAccountFederationPoliciesRequest {
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

// List account federation policies.
func (a *accountFederationPolicyImpl) ListAll(ctx context.Context, request ListAccountFederationPoliciesRequest) ([]FederationPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FederationPolicy](ctx, iterator)
}

func (a *accountFederationPolicyImpl) internalList(ctx context.Context, request ListAccountFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {

	requestPb, pbErr := listAccountFederationPoliciesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listFederationPoliciesResponsePb listFederationPoliciesResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listFederationPoliciesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listFederationPoliciesResponseFromPb(&listFederationPoliciesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountFederationPolicyImpl) Update(ctx context.Context, request UpdateAccountFederationPolicyRequest) (*FederationPolicy, error) {

	requestPb, pbErr := updateAccountFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb federationPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), requestPb.PolicyId)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" || slices.Contains(requestPb.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb).Policy,
		&federationPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := federationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CustomAppIntegration API methods
type customAppIntegrationImpl struct {
	client *client.DatabricksClient
}

func (a *customAppIntegrationImpl) Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error) {

	requestPb, pbErr := createCustomAppIntegrationToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createCustomAppIntegrationOutputPb createCustomAppIntegrationOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createCustomAppIntegrationOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createCustomAppIntegrationOutputFromPb(&createCustomAppIntegrationOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *customAppIntegrationImpl) Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error {

	requestPb, pbErr := deleteCustomAppIntegrationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteCustomAppIntegrationOutputPb deleteCustomAppIntegrationOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), requestPb.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteCustomAppIntegrationOutputPb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *customAppIntegrationImpl) Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error) {

	requestPb, pbErr := getCustomAppIntegrationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getCustomAppIntegrationOutputPb getCustomAppIntegrationOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), requestPb.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getCustomAppIntegrationOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getCustomAppIntegrationOutputFromPb(&getCustomAppIntegrationOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Get the list of custom OAuth app integrations for the specified Databricks
// account
func (a *customAppIntegrationImpl) List(ctx context.Context, request ListCustomAppIntegrationsRequest) listing.Iterator[GetCustomAppIntegrationOutput] {

	getNextPage := func(ctx context.Context, req ListCustomAppIntegrationsRequest) (*GetCustomAppIntegrationsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *GetCustomAppIntegrationsOutput) []GetCustomAppIntegrationOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetCustomAppIntegrationsOutput) *ListCustomAppIntegrationsRequest {
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

// Get the list of custom OAuth app integrations for the specified Databricks
// account
func (a *customAppIntegrationImpl) ListAll(ctx context.Context, request ListCustomAppIntegrationsRequest) ([]GetCustomAppIntegrationOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[GetCustomAppIntegrationOutput](ctx, iterator)
}

func (a *customAppIntegrationImpl) internalList(ctx context.Context, request ListCustomAppIntegrationsRequest) (*GetCustomAppIntegrationsOutput, error) {

	requestPb, pbErr := listCustomAppIntegrationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getCustomAppIntegrationsOutputPb getCustomAppIntegrationsOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getCustomAppIntegrationsOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getCustomAppIntegrationsOutputFromPb(&getCustomAppIntegrationsOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *customAppIntegrationImpl) Update(ctx context.Context, request UpdateCustomAppIntegration) error {

	requestPb, pbErr := updateCustomAppIntegrationToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateCustomAppIntegrationOutputPb updateCustomAppIntegrationOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), requestPb.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateCustomAppIntegrationOutputPb,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just OAuthPublishedApps API methods
type oAuthPublishedAppsImpl struct {
	client *client.DatabricksClient
}

// Get all the available published OAuth apps in Databricks.
func (a *oAuthPublishedAppsImpl) List(ctx context.Context, request ListOAuthPublishedAppsRequest) listing.Iterator[PublishedAppOutput] {

	getNextPage := func(ctx context.Context, req ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *GetPublishedAppsOutput) []PublishedAppOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetPublishedAppsOutput) *ListOAuthPublishedAppsRequest {
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

// Get all the available published OAuth apps in Databricks.
func (a *oAuthPublishedAppsImpl) ListAll(ctx context.Context, request ListOAuthPublishedAppsRequest) ([]PublishedAppOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PublishedAppOutput](ctx, iterator)
}

func (a *oAuthPublishedAppsImpl) internalList(ctx context.Context, request ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error) {

	requestPb, pbErr := listOAuthPublishedAppsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPublishedAppsOutputPb getPublishedAppsOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-apps", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getPublishedAppsOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getPublishedAppsOutputFromPb(&getPublishedAppsOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just PublishedAppIntegration API methods
type publishedAppIntegrationImpl struct {
	client *client.DatabricksClient
}

func (a *publishedAppIntegrationImpl) Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error) {

	requestPb, pbErr := createPublishedAppIntegrationToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createPublishedAppIntegrationOutputPb createPublishedAppIntegrationOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createPublishedAppIntegrationOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createPublishedAppIntegrationOutputFromPb(&createPublishedAppIntegrationOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *publishedAppIntegrationImpl) Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error {

	requestPb, pbErr := deletePublishedAppIntegrationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deletePublishedAppIntegrationOutputPb deletePublishedAppIntegrationOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), requestPb.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deletePublishedAppIntegrationOutputPb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *publishedAppIntegrationImpl) Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error) {

	requestPb, pbErr := getPublishedAppIntegrationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPublishedAppIntegrationOutputPb getPublishedAppIntegrationOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), requestPb.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getPublishedAppIntegrationOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getPublishedAppIntegrationOutputFromPb(&getPublishedAppIntegrationOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Get the list of published OAuth app integrations for the specified Databricks
// account
func (a *publishedAppIntegrationImpl) List(ctx context.Context, request ListPublishedAppIntegrationsRequest) listing.Iterator[GetPublishedAppIntegrationOutput] {

	getNextPage := func(ctx context.Context, req ListPublishedAppIntegrationsRequest) (*GetPublishedAppIntegrationsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *GetPublishedAppIntegrationsOutput) []GetPublishedAppIntegrationOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetPublishedAppIntegrationsOutput) *ListPublishedAppIntegrationsRequest {
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

// Get the list of published OAuth app integrations for the specified Databricks
// account
func (a *publishedAppIntegrationImpl) ListAll(ctx context.Context, request ListPublishedAppIntegrationsRequest) ([]GetPublishedAppIntegrationOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[GetPublishedAppIntegrationOutput](ctx, iterator)
}

func (a *publishedAppIntegrationImpl) internalList(ctx context.Context, request ListPublishedAppIntegrationsRequest) (*GetPublishedAppIntegrationsOutput, error) {

	requestPb, pbErr := listPublishedAppIntegrationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPublishedAppIntegrationsOutputPb getPublishedAppIntegrationsOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getPublishedAppIntegrationsOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getPublishedAppIntegrationsOutputFromPb(&getPublishedAppIntegrationsOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *publishedAppIntegrationImpl) Update(ctx context.Context, request UpdatePublishedAppIntegration) error {

	requestPb, pbErr := updatePublishedAppIntegrationToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updatePublishedAppIntegrationOutputPb updatePublishedAppIntegrationOutputPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), requestPb.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updatePublishedAppIntegrationOutputPb,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just ServicePrincipalFederationPolicy API methods
type servicePrincipalFederationPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalFederationPolicyImpl) Create(ctx context.Context, request CreateServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {

	requestPb, pbErr := createServicePrincipalFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb federationPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies", a.client.ConfiguredAccountID(), requestPb.ServicePrincipalId)
	queryParams := make(map[string]any)
	if requestPb.PolicyId != "" || slices.Contains(requestPb.ForceSendFields, "PolicyId") {
		queryParams["policy_id"] = requestPb.PolicyId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).Policy,
		&federationPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := federationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servicePrincipalFederationPolicyImpl) Delete(ctx context.Context, request DeleteServicePrincipalFederationPolicyRequest) error {

	requestPb, pbErr := deleteServicePrincipalFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), requestPb.ServicePrincipalId, requestPb.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *servicePrincipalFederationPolicyImpl) Get(ctx context.Context, request GetServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {

	requestPb, pbErr := getServicePrincipalFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb federationPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), requestPb.ServicePrincipalId, requestPb.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&federationPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := federationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List account federation policies.
func (a *servicePrincipalFederationPolicyImpl) List(ctx context.Context, request ListServicePrincipalFederationPoliciesRequest) listing.Iterator[FederationPolicy] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFederationPoliciesResponse) []FederationPolicy {
		return resp.Policies
	}
	getNextReq := func(resp *ListFederationPoliciesResponse) *ListServicePrincipalFederationPoliciesRequest {
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

// List account federation policies.
func (a *servicePrincipalFederationPolicyImpl) ListAll(ctx context.Context, request ListServicePrincipalFederationPoliciesRequest) ([]FederationPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FederationPolicy](ctx, iterator)
}

func (a *servicePrincipalFederationPolicyImpl) internalList(ctx context.Context, request ListServicePrincipalFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {

	requestPb, pbErr := listServicePrincipalFederationPoliciesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listFederationPoliciesResponsePb listFederationPoliciesResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies", a.client.ConfiguredAccountID(), requestPb.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listFederationPoliciesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listFederationPoliciesResponseFromPb(&listFederationPoliciesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servicePrincipalFederationPolicyImpl) Update(ctx context.Context, request UpdateServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {

	requestPb, pbErr := updateServicePrincipalFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb federationPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), requestPb.ServicePrincipalId, requestPb.PolicyId)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" || slices.Contains(requestPb.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb).Policy,
		&federationPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := federationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ServicePrincipalSecrets API methods
type servicePrincipalSecretsImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalSecretsImpl) Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error) {

	requestPb, pbErr := createServicePrincipalSecretRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createServicePrincipalSecretResponsePb createServicePrincipalSecretResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets", a.client.ConfiguredAccountID(), requestPb.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createServicePrincipalSecretResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createServicePrincipalSecretResponseFromPb(&createServicePrincipalSecretResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servicePrincipalSecretsImpl) Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error {

	requestPb, pbErr := deleteServicePrincipalSecretRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets/%v", a.client.ConfiguredAccountID(), requestPb.ServicePrincipalId, requestPb.SecretId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
func (a *servicePrincipalSecretsImpl) List(ctx context.Context, request ListServicePrincipalSecretsRequest) listing.Iterator[SecretInfo] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalSecretsResponse) []SecretInfo {
		return resp.Secrets
	}
	getNextReq := func(resp *ListServicePrincipalSecretsResponse) *ListServicePrincipalSecretsRequest {
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

// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
func (a *servicePrincipalSecretsImpl) ListAll(ctx context.Context, request ListServicePrincipalSecretsRequest) ([]SecretInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SecretInfo](ctx, iterator)
}

func (a *servicePrincipalSecretsImpl) internalList(ctx context.Context, request ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {

	requestPb, pbErr := listServicePrincipalSecretsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listServicePrincipalSecretsResponsePb listServicePrincipalSecretsResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets", a.client.ConfiguredAccountID(), requestPb.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listServicePrincipalSecretsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listServicePrincipalSecretsResponseFromPb(&listServicePrincipalSecretsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
