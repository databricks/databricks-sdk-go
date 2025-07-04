// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccountIpAccessLists API methods
type accountIpAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *accountIpAccessListsImpl) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {

	requestPb, pbErr := createIpAccessListToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createIpAccessListResponsePb createIpAccessListResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists", a.client.ConfiguredAccountID())
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
		&createIpAccessListResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createIpAccessListResponseFromPb(&createIpAccessListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountIpAccessListsImpl) Delete(ctx context.Context, request DeleteAccountIpAccessListRequest) error {

	requestPb, pbErr := deleteAccountIpAccessListRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), requestPb.IpAccessListId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *accountIpAccessListsImpl) Get(ctx context.Context, request GetAccountIpAccessListRequest) (*GetIpAccessListResponse, error) {

	requestPb, pbErr := getAccountIpAccessListRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getIpAccessListResponsePb getIpAccessListResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), requestPb.IpAccessListId)
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
		&getIpAccessListResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getIpAccessListResponseFromPb(&getIpAccessListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets all IP access lists for the specified account.
func (a *accountIpAccessListsImpl) List(ctx context.Context) listing.Iterator[IpAccessListInfo] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*GetIpAccessListsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *GetIpAccessListsResponse) []IpAccessListInfo {
		return resp.IpAccessLists
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Gets all IP access lists for the specified account.
func (a *accountIpAccessListsImpl) ListAll(ctx context.Context) ([]IpAccessListInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[IpAccessListInfo](ctx, iterator)
}

func (a *accountIpAccessListsImpl) internalList(ctx context.Context) (*GetIpAccessListsResponse, error) {

	var getIpAccessListsResponsePb getIpAccessListsResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&getIpAccessListsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getIpAccessListsResponseFromPb(&getIpAccessListsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountIpAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {

	requestPb, pbErr := replaceIpAccessListToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), requestPb.IpAccessListId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *accountIpAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {

	requestPb, pbErr := updateIpAccessListToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), requestPb.IpAccessListId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just AccountSettings API methods
type accountSettingsImpl struct {
	client *client.DatabricksClient
}

// unexported type that holds implementations of just AibiDashboardEmbeddingAccessPolicy API methods
type aibiDashboardEmbeddingAccessPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *aibiDashboardEmbeddingAccessPolicyImpl) Delete(ctx context.Context, request DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error) {

	requestPb, pbErr := deleteAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
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
		&deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteAibiDashboardEmbeddingAccessPolicySettingResponseFromPb(&deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aibiDashboardEmbeddingAccessPolicyImpl) Get(ctx context.Context, request GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error) {

	requestPb, pbErr := getAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aibiDashboardEmbeddingAccessPolicySettingPb aibiDashboardEmbeddingAccessPolicySettingPb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
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
		&aibiDashboardEmbeddingAccessPolicySettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := aibiDashboardEmbeddingAccessPolicySettingFromPb(&aibiDashboardEmbeddingAccessPolicySettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aibiDashboardEmbeddingAccessPolicyImpl) Update(ctx context.Context, request UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error) {

	requestPb, pbErr := updateAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aibiDashboardEmbeddingAccessPolicySettingPb aibiDashboardEmbeddingAccessPolicySettingPb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
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
		&aibiDashboardEmbeddingAccessPolicySettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := aibiDashboardEmbeddingAccessPolicySettingFromPb(&aibiDashboardEmbeddingAccessPolicySettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just AibiDashboardEmbeddingApprovedDomains API methods
type aibiDashboardEmbeddingApprovedDomainsImpl struct {
	client *client.DatabricksClient
}

func (a *aibiDashboardEmbeddingApprovedDomainsImpl) Delete(ctx context.Context, request DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error) {

	requestPb, pbErr := deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
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
		&deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteAibiDashboardEmbeddingApprovedDomainsSettingResponseFromPb(&deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aibiDashboardEmbeddingApprovedDomainsImpl) Get(ctx context.Context, request GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {

	requestPb, pbErr := getAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aibiDashboardEmbeddingApprovedDomainsSettingPb aibiDashboardEmbeddingApprovedDomainsSettingPb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
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
		&aibiDashboardEmbeddingApprovedDomainsSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := aibiDashboardEmbeddingApprovedDomainsSettingFromPb(&aibiDashboardEmbeddingApprovedDomainsSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aibiDashboardEmbeddingApprovedDomainsImpl) Update(ctx context.Context, request UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {

	requestPb, pbErr := updateAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aibiDashboardEmbeddingApprovedDomainsSettingPb aibiDashboardEmbeddingApprovedDomainsSettingPb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
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
		&aibiDashboardEmbeddingApprovedDomainsSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := aibiDashboardEmbeddingApprovedDomainsSettingFromPb(&aibiDashboardEmbeddingApprovedDomainsSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just AutomaticClusterUpdate API methods
type automaticClusterUpdateImpl struct {
	client *client.DatabricksClient
}

func (a *automaticClusterUpdateImpl) Get(ctx context.Context, request GetAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error) {

	requestPb, pbErr := getAutomaticClusterUpdateSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var automaticClusterUpdateSettingPb automaticClusterUpdateSettingPb
	path := "/api/2.0/settings/types/automatic_cluster_update/names/default"
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
		&automaticClusterUpdateSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := automaticClusterUpdateSettingFromPb(&automaticClusterUpdateSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *automaticClusterUpdateImpl) Update(ctx context.Context, request UpdateAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error) {

	requestPb, pbErr := updateAutomaticClusterUpdateSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var automaticClusterUpdateSettingPb automaticClusterUpdateSettingPb
	path := "/api/2.0/settings/types/automatic_cluster_update/names/default"
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
		&automaticClusterUpdateSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := automaticClusterUpdateSettingFromPb(&automaticClusterUpdateSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ComplianceSecurityProfile API methods
type complianceSecurityProfileImpl struct {
	client *client.DatabricksClient
}

func (a *complianceSecurityProfileImpl) Get(ctx context.Context, request GetComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error) {

	requestPb, pbErr := getComplianceSecurityProfileSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var complianceSecurityProfileSettingPb complianceSecurityProfileSettingPb
	path := "/api/2.0/settings/types/shield_csp_enablement_ws_db/names/default"
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
		&complianceSecurityProfileSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := complianceSecurityProfileSettingFromPb(&complianceSecurityProfileSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *complianceSecurityProfileImpl) Update(ctx context.Context, request UpdateComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error) {

	requestPb, pbErr := updateComplianceSecurityProfileSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var complianceSecurityProfileSettingPb complianceSecurityProfileSettingPb
	path := "/api/2.0/settings/types/shield_csp_enablement_ws_db/names/default"
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
		&complianceSecurityProfileSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := complianceSecurityProfileSettingFromPb(&complianceSecurityProfileSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CredentialsManager API methods
type credentialsManagerImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsManagerImpl) ExchangeToken(ctx context.Context, request ExchangeTokenRequest) (*ExchangeTokenResponse, error) {

	requestPb, pbErr := exchangeTokenRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var exchangeTokenResponsePb exchangeTokenResponsePb
	path := "/api/2.0/credentials-manager/exchange-tokens/token"
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
		&exchangeTokenResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := exchangeTokenResponseFromPb(&exchangeTokenResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CspEnablementAccount API methods
type cspEnablementAccountImpl struct {
	client *client.DatabricksClient
}

func (a *cspEnablementAccountImpl) Get(ctx context.Context, request GetCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error) {

	requestPb, pbErr := getCspEnablementAccountSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cspEnablementAccountSettingPb cspEnablementAccountSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_csp_enablement_ac/names/default", a.client.ConfiguredAccountID())
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
		&cspEnablementAccountSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := cspEnablementAccountSettingFromPb(&cspEnablementAccountSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cspEnablementAccountImpl) Update(ctx context.Context, request UpdateCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error) {

	requestPb, pbErr := updateCspEnablementAccountSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cspEnablementAccountSettingPb cspEnablementAccountSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_csp_enablement_ac/names/default", a.client.ConfiguredAccountID())
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
		&cspEnablementAccountSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := cspEnablementAccountSettingFromPb(&cspEnablementAccountSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just DashboardEmailSubscriptions API methods
type dashboardEmailSubscriptionsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardEmailSubscriptionsImpl) Delete(ctx context.Context, request DeleteDashboardEmailSubscriptionsRequest) (*DeleteDashboardEmailSubscriptionsResponse, error) {

	requestPb, pbErr := deleteDashboardEmailSubscriptionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDashboardEmailSubscriptionsResponsePb deleteDashboardEmailSubscriptionsResponsePb
	path := "/api/2.0/settings/types/dashboard_email_subscriptions/names/default"
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
		&deleteDashboardEmailSubscriptionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteDashboardEmailSubscriptionsResponseFromPb(&deleteDashboardEmailSubscriptionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardEmailSubscriptionsImpl) Get(ctx context.Context, request GetDashboardEmailSubscriptionsRequest) (*DashboardEmailSubscriptions, error) {

	requestPb, pbErr := getDashboardEmailSubscriptionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardEmailSubscriptionsPb dashboardEmailSubscriptionsPb
	path := "/api/2.0/settings/types/dashboard_email_subscriptions/names/default"
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
		&dashboardEmailSubscriptionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := dashboardEmailSubscriptionsFromPb(&dashboardEmailSubscriptionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardEmailSubscriptionsImpl) Update(ctx context.Context, request UpdateDashboardEmailSubscriptionsRequest) (*DashboardEmailSubscriptions, error) {

	requestPb, pbErr := updateDashboardEmailSubscriptionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardEmailSubscriptionsPb dashboardEmailSubscriptionsPb
	path := "/api/2.0/settings/types/dashboard_email_subscriptions/names/default"
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
		&dashboardEmailSubscriptionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := dashboardEmailSubscriptionsFromPb(&dashboardEmailSubscriptionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just DefaultNamespace API methods
type defaultNamespaceImpl struct {
	client *client.DatabricksClient
}

func (a *defaultNamespaceImpl) Delete(ctx context.Context, request DeleteDefaultNamespaceSettingRequest) (*DeleteDefaultNamespaceSettingResponse, error) {

	requestPb, pbErr := deleteDefaultNamespaceSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDefaultNamespaceSettingResponsePb deleteDefaultNamespaceSettingResponsePb
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
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
		&deleteDefaultNamespaceSettingResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteDefaultNamespaceSettingResponseFromPb(&deleteDefaultNamespaceSettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *defaultNamespaceImpl) Get(ctx context.Context, request GetDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error) {

	requestPb, pbErr := getDefaultNamespaceSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var defaultNamespaceSettingPb defaultNamespaceSettingPb
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
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
		&defaultNamespaceSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := defaultNamespaceSettingFromPb(&defaultNamespaceSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *defaultNamespaceImpl) Update(ctx context.Context, request UpdateDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error) {

	requestPb, pbErr := updateDefaultNamespaceSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var defaultNamespaceSettingPb defaultNamespaceSettingPb
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
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
		&defaultNamespaceSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := defaultNamespaceSettingFromPb(&defaultNamespaceSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just DisableLegacyAccess API methods
type disableLegacyAccessImpl struct {
	client *client.DatabricksClient
}

func (a *disableLegacyAccessImpl) Delete(ctx context.Context, request DeleteDisableLegacyAccessRequest) (*DeleteDisableLegacyAccessResponse, error) {

	requestPb, pbErr := deleteDisableLegacyAccessRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDisableLegacyAccessResponsePb deleteDisableLegacyAccessResponsePb
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
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
		&deleteDisableLegacyAccessResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteDisableLegacyAccessResponseFromPb(&deleteDisableLegacyAccessResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyAccessImpl) Get(ctx context.Context, request GetDisableLegacyAccessRequest) (*DisableLegacyAccess, error) {

	requestPb, pbErr := getDisableLegacyAccessRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyAccessPb disableLegacyAccessPb
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
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
		&disableLegacyAccessPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := disableLegacyAccessFromPb(&disableLegacyAccessPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyAccessImpl) Update(ctx context.Context, request UpdateDisableLegacyAccessRequest) (*DisableLegacyAccess, error) {

	requestPb, pbErr := updateDisableLegacyAccessRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyAccessPb disableLegacyAccessPb
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
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
		&disableLegacyAccessPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := disableLegacyAccessFromPb(&disableLegacyAccessPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just DisableLegacyDbfs API methods
type disableLegacyDbfsImpl struct {
	client *client.DatabricksClient
}

func (a *disableLegacyDbfsImpl) Delete(ctx context.Context, request DeleteDisableLegacyDbfsRequest) (*DeleteDisableLegacyDbfsResponse, error) {

	requestPb, pbErr := deleteDisableLegacyDbfsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDisableLegacyDbfsResponsePb deleteDisableLegacyDbfsResponsePb
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
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
		&deleteDisableLegacyDbfsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteDisableLegacyDbfsResponseFromPb(&deleteDisableLegacyDbfsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyDbfsImpl) Get(ctx context.Context, request GetDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error) {

	requestPb, pbErr := getDisableLegacyDbfsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyDbfsPb disableLegacyDbfsPb
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
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
		&disableLegacyDbfsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := disableLegacyDbfsFromPb(&disableLegacyDbfsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyDbfsImpl) Update(ctx context.Context, request UpdateDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error) {

	requestPb, pbErr := updateDisableLegacyDbfsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyDbfsPb disableLegacyDbfsPb
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
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
		&disableLegacyDbfsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := disableLegacyDbfsFromPb(&disableLegacyDbfsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just DisableLegacyFeatures API methods
type disableLegacyFeaturesImpl struct {
	client *client.DatabricksClient
}

func (a *disableLegacyFeaturesImpl) Delete(ctx context.Context, request DeleteDisableLegacyFeaturesRequest) (*DeleteDisableLegacyFeaturesResponse, error) {

	requestPb, pbErr := deleteDisableLegacyFeaturesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDisableLegacyFeaturesResponsePb deleteDisableLegacyFeaturesResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
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
		&deleteDisableLegacyFeaturesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteDisableLegacyFeaturesResponseFromPb(&deleteDisableLegacyFeaturesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyFeaturesImpl) Get(ctx context.Context, request GetDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error) {

	requestPb, pbErr := getDisableLegacyFeaturesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyFeaturesPb disableLegacyFeaturesPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
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
		&disableLegacyFeaturesPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := disableLegacyFeaturesFromPb(&disableLegacyFeaturesPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyFeaturesImpl) Update(ctx context.Context, request UpdateDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error) {

	requestPb, pbErr := updateDisableLegacyFeaturesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyFeaturesPb disableLegacyFeaturesPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
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
		&disableLegacyFeaturesPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := disableLegacyFeaturesFromPb(&disableLegacyFeaturesPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just EnableExportNotebook API methods
type enableExportNotebookImpl struct {
	client *client.DatabricksClient
}

func (a *enableExportNotebookImpl) GetEnableExportNotebook(ctx context.Context) (*EnableExportNotebook, error) {

	var enableExportNotebookPb enableExportNotebookPb
	path := "/api/2.0/settings/types/enable-export-notebook/names/default"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&enableExportNotebookPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enableExportNotebookFromPb(&enableExportNotebookPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableExportNotebookImpl) PatchEnableExportNotebook(ctx context.Context, request UpdateEnableExportNotebookRequest) (*EnableExportNotebook, error) {

	requestPb, pbErr := updateEnableExportNotebookRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enableExportNotebookPb enableExportNotebookPb
	path := "/api/2.0/settings/types/enable-export-notebook/names/default"
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
		&enableExportNotebookPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enableExportNotebookFromPb(&enableExportNotebookPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just EnableIpAccessLists API methods
type enableIpAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *enableIpAccessListsImpl) Delete(ctx context.Context, request DeleteAccountIpAccessEnableRequest) (*DeleteAccountIpAccessEnableResponse, error) {

	requestPb, pbErr := deleteAccountIpAccessEnableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteAccountIpAccessEnableResponsePb deleteAccountIpAccessEnableResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/acct_ip_acl_enable/names/default", a.client.ConfiguredAccountID())
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
		&deleteAccountIpAccessEnableResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteAccountIpAccessEnableResponseFromPb(&deleteAccountIpAccessEnableResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableIpAccessListsImpl) Get(ctx context.Context, request GetAccountIpAccessEnableRequest) (*AccountIpAccessEnable, error) {

	requestPb, pbErr := getAccountIpAccessEnableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountIpAccessEnablePb accountIpAccessEnablePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/acct_ip_acl_enable/names/default", a.client.ConfiguredAccountID())
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
		&accountIpAccessEnablePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := accountIpAccessEnableFromPb(&accountIpAccessEnablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableIpAccessListsImpl) Update(ctx context.Context, request UpdateAccountIpAccessEnableRequest) (*AccountIpAccessEnable, error) {

	requestPb, pbErr := updateAccountIpAccessEnableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountIpAccessEnablePb accountIpAccessEnablePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/acct_ip_acl_enable/names/default", a.client.ConfiguredAccountID())
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
		&accountIpAccessEnablePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := accountIpAccessEnableFromPb(&accountIpAccessEnablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just EnableNotebookTableClipboard API methods
type enableNotebookTableClipboardImpl struct {
	client *client.DatabricksClient
}

func (a *enableNotebookTableClipboardImpl) GetEnableNotebookTableClipboard(ctx context.Context) (*EnableNotebookTableClipboard, error) {

	var enableNotebookTableClipboardPb enableNotebookTableClipboardPb
	path := "/api/2.0/settings/types/enable-notebook-table-clipboard/names/default"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&enableNotebookTableClipboardPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enableNotebookTableClipboardFromPb(&enableNotebookTableClipboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableNotebookTableClipboardImpl) PatchEnableNotebookTableClipboard(ctx context.Context, request UpdateEnableNotebookTableClipboardRequest) (*EnableNotebookTableClipboard, error) {

	requestPb, pbErr := updateEnableNotebookTableClipboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enableNotebookTableClipboardPb enableNotebookTableClipboardPb
	path := "/api/2.0/settings/types/enable-notebook-table-clipboard/names/default"
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
		&enableNotebookTableClipboardPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enableNotebookTableClipboardFromPb(&enableNotebookTableClipboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just EnableResultsDownloading API methods
type enableResultsDownloadingImpl struct {
	client *client.DatabricksClient
}

func (a *enableResultsDownloadingImpl) GetEnableResultsDownloading(ctx context.Context) (*EnableResultsDownloading, error) {

	var enableResultsDownloadingPb enableResultsDownloadingPb
	path := "/api/2.0/settings/types/enable-results-downloading/names/default"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&enableResultsDownloadingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enableResultsDownloadingFromPb(&enableResultsDownloadingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableResultsDownloadingImpl) PatchEnableResultsDownloading(ctx context.Context, request UpdateEnableResultsDownloadingRequest) (*EnableResultsDownloading, error) {

	requestPb, pbErr := updateEnableResultsDownloadingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enableResultsDownloadingPb enableResultsDownloadingPb
	path := "/api/2.0/settings/types/enable-results-downloading/names/default"
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
		&enableResultsDownloadingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enableResultsDownloadingFromPb(&enableResultsDownloadingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just EnhancedSecurityMonitoring API methods
type enhancedSecurityMonitoringImpl struct {
	client *client.DatabricksClient
}

func (a *enhancedSecurityMonitoringImpl) Get(ctx context.Context, request GetEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error) {

	requestPb, pbErr := getEnhancedSecurityMonitoringSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enhancedSecurityMonitoringSettingPb enhancedSecurityMonitoringSettingPb
	path := "/api/2.0/settings/types/shield_esm_enablement_ws_db/names/default"
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
		&enhancedSecurityMonitoringSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enhancedSecurityMonitoringSettingFromPb(&enhancedSecurityMonitoringSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enhancedSecurityMonitoringImpl) Update(ctx context.Context, request UpdateEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error) {

	requestPb, pbErr := updateEnhancedSecurityMonitoringSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enhancedSecurityMonitoringSettingPb enhancedSecurityMonitoringSettingPb
	path := "/api/2.0/settings/types/shield_esm_enablement_ws_db/names/default"
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
		&enhancedSecurityMonitoringSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enhancedSecurityMonitoringSettingFromPb(&enhancedSecurityMonitoringSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just EsmEnablementAccount API methods
type esmEnablementAccountImpl struct {
	client *client.DatabricksClient
}

func (a *esmEnablementAccountImpl) Get(ctx context.Context, request GetEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error) {

	requestPb, pbErr := getEsmEnablementAccountSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var esmEnablementAccountSettingPb esmEnablementAccountSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_esm_enablement_ac/names/default", a.client.ConfiguredAccountID())
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
		&esmEnablementAccountSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := esmEnablementAccountSettingFromPb(&esmEnablementAccountSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *esmEnablementAccountImpl) Update(ctx context.Context, request UpdateEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error) {

	requestPb, pbErr := updateEsmEnablementAccountSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var esmEnablementAccountSettingPb esmEnablementAccountSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_esm_enablement_ac/names/default", a.client.ConfiguredAccountID())
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
		&esmEnablementAccountSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := esmEnablementAccountSettingFromPb(&esmEnablementAccountSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just IpAccessLists API methods
type ipAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *ipAccessListsImpl) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {

	requestPb, pbErr := createIpAccessListToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createIpAccessListResponsePb createIpAccessListResponsePb
	path := "/api/2.0/ip-access-lists"
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
		&createIpAccessListResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createIpAccessListResponseFromPb(&createIpAccessListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *ipAccessListsImpl) Delete(ctx context.Context, request DeleteIpAccessListRequest) error {

	requestPb, pbErr := deleteIpAccessListRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", requestPb.IpAccessListId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *ipAccessListsImpl) Get(ctx context.Context, request GetIpAccessListRequest) (*FetchIpAccessListResponse, error) {

	requestPb, pbErr := getIpAccessListRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var fetchIpAccessListResponsePb fetchIpAccessListResponsePb
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", requestPb.IpAccessListId)
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
		&fetchIpAccessListResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := fetchIpAccessListResponseFromPb(&fetchIpAccessListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets all IP access lists for the specified workspace.
func (a *ipAccessListsImpl) List(ctx context.Context) listing.Iterator[IpAccessListInfo] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListIpAccessListResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListIpAccessListResponse) []IpAccessListInfo {
		return resp.IpAccessLists
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Gets all IP access lists for the specified workspace.
func (a *ipAccessListsImpl) ListAll(ctx context.Context) ([]IpAccessListInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[IpAccessListInfo](ctx, iterator)
}

func (a *ipAccessListsImpl) internalList(ctx context.Context) (*ListIpAccessListResponse, error) {

	var listIpAccessListResponsePb listIpAccessListResponsePb
	path := "/api/2.0/ip-access-lists"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&listIpAccessListResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listIpAccessListResponseFromPb(&listIpAccessListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *ipAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {

	requestPb, pbErr := replaceIpAccessListToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", requestPb.IpAccessListId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *ipAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {

	requestPb, pbErr := updateIpAccessListToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", requestPb.IpAccessListId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just LlmProxyPartnerPoweredAccount API methods
type llmProxyPartnerPoweredAccountImpl struct {
	client *client.DatabricksClient
}

func (a *llmProxyPartnerPoweredAccountImpl) Get(ctx context.Context, request GetLlmProxyPartnerPoweredAccountRequest) (*LlmProxyPartnerPoweredAccount, error) {

	requestPb, pbErr := getLlmProxyPartnerPoweredAccountRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredAccountPb llmProxyPartnerPoweredAccountPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/llm_proxy_partner_powered/names/default", a.client.ConfiguredAccountID())
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
		&llmProxyPartnerPoweredAccountPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := llmProxyPartnerPoweredAccountFromPb(&llmProxyPartnerPoweredAccountPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *llmProxyPartnerPoweredAccountImpl) Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredAccountRequest) (*LlmProxyPartnerPoweredAccount, error) {

	requestPb, pbErr := updateLlmProxyPartnerPoweredAccountRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredAccountPb llmProxyPartnerPoweredAccountPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/llm_proxy_partner_powered/names/default", a.client.ConfiguredAccountID())
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
		&llmProxyPartnerPoweredAccountPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := llmProxyPartnerPoweredAccountFromPb(&llmProxyPartnerPoweredAccountPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just LlmProxyPartnerPoweredEnforce API methods
type llmProxyPartnerPoweredEnforceImpl struct {
	client *client.DatabricksClient
}

func (a *llmProxyPartnerPoweredEnforceImpl) Get(ctx context.Context, request GetLlmProxyPartnerPoweredEnforceRequest) (*LlmProxyPartnerPoweredEnforce, error) {

	requestPb, pbErr := getLlmProxyPartnerPoweredEnforceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredEnforcePb llmProxyPartnerPoweredEnforcePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/llm_proxy_partner_powered_enforce/names/default", a.client.ConfiguredAccountID())
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
		&llmProxyPartnerPoweredEnforcePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := llmProxyPartnerPoweredEnforceFromPb(&llmProxyPartnerPoweredEnforcePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *llmProxyPartnerPoweredEnforceImpl) Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredEnforceRequest) (*LlmProxyPartnerPoweredEnforce, error) {

	requestPb, pbErr := updateLlmProxyPartnerPoweredEnforceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredEnforcePb llmProxyPartnerPoweredEnforcePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/llm_proxy_partner_powered_enforce/names/default", a.client.ConfiguredAccountID())
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
		&llmProxyPartnerPoweredEnforcePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := llmProxyPartnerPoweredEnforceFromPb(&llmProxyPartnerPoweredEnforcePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just LlmProxyPartnerPoweredWorkspace API methods
type llmProxyPartnerPoweredWorkspaceImpl struct {
	client *client.DatabricksClient
}

func (a *llmProxyPartnerPoweredWorkspaceImpl) Delete(ctx context.Context, request DeleteLlmProxyPartnerPoweredWorkspaceRequest) (*DeleteLlmProxyPartnerPoweredWorkspaceResponse, error) {

	requestPb, pbErr := deleteLlmProxyPartnerPoweredWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteLlmProxyPartnerPoweredWorkspaceResponsePb deleteLlmProxyPartnerPoweredWorkspaceResponsePb
	path := "/api/2.0/settings/types/llm_proxy_partner_powered/names/default"
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
		&deleteLlmProxyPartnerPoweredWorkspaceResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteLlmProxyPartnerPoweredWorkspaceResponseFromPb(&deleteLlmProxyPartnerPoweredWorkspaceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *llmProxyPartnerPoweredWorkspaceImpl) Get(ctx context.Context, request GetLlmProxyPartnerPoweredWorkspaceRequest) (*LlmProxyPartnerPoweredWorkspace, error) {

	requestPb, pbErr := getLlmProxyPartnerPoweredWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredWorkspacePb llmProxyPartnerPoweredWorkspacePb
	path := "/api/2.0/settings/types/llm_proxy_partner_powered/names/default"
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
		&llmProxyPartnerPoweredWorkspacePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := llmProxyPartnerPoweredWorkspaceFromPb(&llmProxyPartnerPoweredWorkspacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *llmProxyPartnerPoweredWorkspaceImpl) Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredWorkspaceRequest) (*LlmProxyPartnerPoweredWorkspace, error) {

	requestPb, pbErr := updateLlmProxyPartnerPoweredWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredWorkspacePb llmProxyPartnerPoweredWorkspacePb
	path := "/api/2.0/settings/types/llm_proxy_partner_powered/names/default"
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
		&llmProxyPartnerPoweredWorkspacePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := llmProxyPartnerPoweredWorkspaceFromPb(&llmProxyPartnerPoweredWorkspacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just NetworkConnectivity API methods
type networkConnectivityImpl struct {
	client *client.DatabricksClient
}

func (a *networkConnectivityImpl) CreateNetworkConnectivityConfiguration(ctx context.Context, request CreateNetworkConnectivityConfigRequest) (*NetworkConnectivityConfiguration, error) {

	requestPb, pbErr := createNetworkConnectivityConfigRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var networkConnectivityConfigurationPb networkConnectivityConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs", a.client.ConfiguredAccountID())
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
		(*requestPb).NetworkConnectivityConfig,
		&networkConnectivityConfigurationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := networkConnectivityConfigurationFromPb(&networkConnectivityConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) CreatePrivateEndpointRule(ctx context.Context, request CreatePrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error) {

	requestPb, pbErr := createPrivateEndpointRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var nccPrivateEndpointRulePb nccPrivateEndpointRulePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules", a.client.ConfiguredAccountID(), requestPb.NetworkConnectivityConfigId)
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
		(*requestPb).PrivateEndpointRule,
		&nccPrivateEndpointRulePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := nccPrivateEndpointRuleFromPb(&nccPrivateEndpointRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) DeleteNetworkConnectivityConfiguration(ctx context.Context, request DeleteNetworkConnectivityConfigurationRequest) error {

	requestPb, pbErr := deleteNetworkConnectivityConfigurationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v", a.client.ConfiguredAccountID(), requestPb.NetworkConnectivityConfigId)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *networkConnectivityImpl) DeletePrivateEndpointRule(ctx context.Context, request DeletePrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error) {

	requestPb, pbErr := deletePrivateEndpointRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var nccPrivateEndpointRulePb nccPrivateEndpointRulePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules/%v", a.client.ConfiguredAccountID(), requestPb.NetworkConnectivityConfigId, requestPb.PrivateEndpointRuleId)
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
		&nccPrivateEndpointRulePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := nccPrivateEndpointRuleFromPb(&nccPrivateEndpointRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) GetNetworkConnectivityConfiguration(ctx context.Context, request GetNetworkConnectivityConfigurationRequest) (*NetworkConnectivityConfiguration, error) {

	requestPb, pbErr := getNetworkConnectivityConfigurationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var networkConnectivityConfigurationPb networkConnectivityConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v", a.client.ConfiguredAccountID(), requestPb.NetworkConnectivityConfigId)
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
		&networkConnectivityConfigurationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := networkConnectivityConfigurationFromPb(&networkConnectivityConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) GetPrivateEndpointRule(ctx context.Context, request GetPrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error) {

	requestPb, pbErr := getPrivateEndpointRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var nccPrivateEndpointRulePb nccPrivateEndpointRulePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules/%v", a.client.ConfiguredAccountID(), requestPb.NetworkConnectivityConfigId, requestPb.PrivateEndpointRuleId)
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
		&nccPrivateEndpointRulePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := nccPrivateEndpointRuleFromPb(&nccPrivateEndpointRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of network connectivity configurations.
func (a *networkConnectivityImpl) ListNetworkConnectivityConfigurations(ctx context.Context, request ListNetworkConnectivityConfigurationsRequest) listing.Iterator[NetworkConnectivityConfiguration] {

	getNextPage := func(ctx context.Context, req ListNetworkConnectivityConfigurationsRequest) (*ListNetworkConnectivityConfigurationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListNetworkConnectivityConfigurations(ctx, req)
	}
	getItems := func(resp *ListNetworkConnectivityConfigurationsResponse) []NetworkConnectivityConfiguration {
		return resp.Items
	}
	getNextReq := func(resp *ListNetworkConnectivityConfigurationsResponse) *ListNetworkConnectivityConfigurationsRequest {
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

// Gets an array of network connectivity configurations.
func (a *networkConnectivityImpl) ListNetworkConnectivityConfigurationsAll(ctx context.Context, request ListNetworkConnectivityConfigurationsRequest) ([]NetworkConnectivityConfiguration, error) {
	iterator := a.ListNetworkConnectivityConfigurations(ctx, request)
	return listing.ToSlice[NetworkConnectivityConfiguration](ctx, iterator)
}

func (a *networkConnectivityImpl) internalListNetworkConnectivityConfigurations(ctx context.Context, request ListNetworkConnectivityConfigurationsRequest) (*ListNetworkConnectivityConfigurationsResponse, error) {

	requestPb, pbErr := listNetworkConnectivityConfigurationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listNetworkConnectivityConfigurationsResponsePb listNetworkConnectivityConfigurationsResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs", a.client.ConfiguredAccountID())
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
		&listNetworkConnectivityConfigurationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listNetworkConnectivityConfigurationsResponseFromPb(&listNetworkConnectivityConfigurationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of private endpoint rules.
func (a *networkConnectivityImpl) ListPrivateEndpointRules(ctx context.Context, request ListPrivateEndpointRulesRequest) listing.Iterator[NccPrivateEndpointRule] {

	getNextPage := func(ctx context.Context, req ListPrivateEndpointRulesRequest) (*ListPrivateEndpointRulesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListPrivateEndpointRules(ctx, req)
	}
	getItems := func(resp *ListPrivateEndpointRulesResponse) []NccPrivateEndpointRule {
		return resp.Items
	}
	getNextReq := func(resp *ListPrivateEndpointRulesResponse) *ListPrivateEndpointRulesRequest {
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

// Gets an array of private endpoint rules.
func (a *networkConnectivityImpl) ListPrivateEndpointRulesAll(ctx context.Context, request ListPrivateEndpointRulesRequest) ([]NccPrivateEndpointRule, error) {
	iterator := a.ListPrivateEndpointRules(ctx, request)
	return listing.ToSlice[NccPrivateEndpointRule](ctx, iterator)
}

func (a *networkConnectivityImpl) internalListPrivateEndpointRules(ctx context.Context, request ListPrivateEndpointRulesRequest) (*ListPrivateEndpointRulesResponse, error) {

	requestPb, pbErr := listPrivateEndpointRulesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listPrivateEndpointRulesResponsePb listPrivateEndpointRulesResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules", a.client.ConfiguredAccountID(), requestPb.NetworkConnectivityConfigId)
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
		&listPrivateEndpointRulesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listPrivateEndpointRulesResponseFromPb(&listPrivateEndpointRulesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) UpdatePrivateEndpointRule(ctx context.Context, request UpdateNccPrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error) {

	requestPb, pbErr := updateNccPrivateEndpointRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var nccPrivateEndpointRulePb nccPrivateEndpointRulePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules/%v", a.client.ConfiguredAccountID(), requestPb.NetworkConnectivityConfigId, requestPb.PrivateEndpointRuleId)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" {
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
		(*requestPb).PrivateEndpointRule,
		&nccPrivateEndpointRulePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := nccPrivateEndpointRuleFromPb(&nccPrivateEndpointRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just NetworkPolicies API methods
type networkPoliciesImpl struct {
	client *client.DatabricksClient
}

func (a *networkPoliciesImpl) CreateNetworkPolicyRpc(ctx context.Context, request CreateNetworkPolicyRequest) (*AccountNetworkPolicy, error) {

	requestPb, pbErr := createNetworkPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountNetworkPolicyPb accountNetworkPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies", a.client.ConfiguredAccountID())
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
		(*requestPb).NetworkPolicy,
		&accountNetworkPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := accountNetworkPolicyFromPb(&accountNetworkPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkPoliciesImpl) DeleteNetworkPolicyRpc(ctx context.Context, request DeleteNetworkPolicyRequest) error {

	requestPb, pbErr := deleteNetworkPolicyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies/%v", a.client.ConfiguredAccountID(), requestPb.NetworkPolicyId)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *networkPoliciesImpl) GetNetworkPolicyRpc(ctx context.Context, request GetNetworkPolicyRequest) (*AccountNetworkPolicy, error) {

	requestPb, pbErr := getNetworkPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountNetworkPolicyPb accountNetworkPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies/%v", a.client.ConfiguredAccountID(), requestPb.NetworkPolicyId)
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
		&accountNetworkPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := accountNetworkPolicyFromPb(&accountNetworkPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of network policies.
func (a *networkPoliciesImpl) ListNetworkPoliciesRpc(ctx context.Context, request ListNetworkPoliciesRequest) listing.Iterator[AccountNetworkPolicy] {

	getNextPage := func(ctx context.Context, req ListNetworkPoliciesRequest) (*ListNetworkPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListNetworkPoliciesRpc(ctx, req)
	}
	getItems := func(resp *ListNetworkPoliciesResponse) []AccountNetworkPolicy {
		return resp.Items
	}
	getNextReq := func(resp *ListNetworkPoliciesResponse) *ListNetworkPoliciesRequest {
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

// Gets an array of network policies.
func (a *networkPoliciesImpl) ListNetworkPoliciesRpcAll(ctx context.Context, request ListNetworkPoliciesRequest) ([]AccountNetworkPolicy, error) {
	iterator := a.ListNetworkPoliciesRpc(ctx, request)
	return listing.ToSlice[AccountNetworkPolicy](ctx, iterator)
}

func (a *networkPoliciesImpl) internalListNetworkPoliciesRpc(ctx context.Context, request ListNetworkPoliciesRequest) (*ListNetworkPoliciesResponse, error) {

	requestPb, pbErr := listNetworkPoliciesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listNetworkPoliciesResponsePb listNetworkPoliciesResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies", a.client.ConfiguredAccountID())
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
		&listNetworkPoliciesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listNetworkPoliciesResponseFromPb(&listNetworkPoliciesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkPoliciesImpl) UpdateNetworkPolicyRpc(ctx context.Context, request UpdateNetworkPolicyRequest) (*AccountNetworkPolicy, error) {

	requestPb, pbErr := updateNetworkPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountNetworkPolicyPb accountNetworkPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies/%v", a.client.ConfiguredAccountID(), requestPb.NetworkPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb).NetworkPolicy,
		&accountNetworkPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := accountNetworkPolicyFromPb(&accountNetworkPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just NotificationDestinations API methods
type notificationDestinationsImpl struct {
	client *client.DatabricksClient
}

func (a *notificationDestinationsImpl) Create(ctx context.Context, request CreateNotificationDestinationRequest) (*NotificationDestination, error) {

	requestPb, pbErr := createNotificationDestinationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var notificationDestinationPb notificationDestinationPb
	path := "/api/2.0/notification-destinations"
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
		&notificationDestinationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := notificationDestinationFromPb(&notificationDestinationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *notificationDestinationsImpl) Delete(ctx context.Context, request DeleteNotificationDestinationRequest) error {

	requestPb, pbErr := deleteNotificationDestinationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", requestPb.Id)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *notificationDestinationsImpl) Get(ctx context.Context, request GetNotificationDestinationRequest) (*NotificationDestination, error) {

	requestPb, pbErr := getNotificationDestinationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var notificationDestinationPb notificationDestinationPb
	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", requestPb.Id)
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
		&notificationDestinationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := notificationDestinationFromPb(&notificationDestinationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists notification destinations.
func (a *notificationDestinationsImpl) List(ctx context.Context, request ListNotificationDestinationsRequest) listing.Iterator[ListNotificationDestinationsResult] {

	getNextPage := func(ctx context.Context, req ListNotificationDestinationsRequest) (*ListNotificationDestinationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListNotificationDestinationsResponse) []ListNotificationDestinationsResult {
		return resp.Results
	}
	getNextReq := func(resp *ListNotificationDestinationsResponse) *ListNotificationDestinationsRequest {
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

// Lists notification destinations.
func (a *notificationDestinationsImpl) ListAll(ctx context.Context, request ListNotificationDestinationsRequest) ([]ListNotificationDestinationsResult, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListNotificationDestinationsResult](ctx, iterator)
}

func (a *notificationDestinationsImpl) internalList(ctx context.Context, request ListNotificationDestinationsRequest) (*ListNotificationDestinationsResponse, error) {

	requestPb, pbErr := listNotificationDestinationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listNotificationDestinationsResponsePb listNotificationDestinationsResponsePb
	path := "/api/2.0/notification-destinations"
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
		&listNotificationDestinationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listNotificationDestinationsResponseFromPb(&listNotificationDestinationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *notificationDestinationsImpl) Update(ctx context.Context, request UpdateNotificationDestinationRequest) (*NotificationDestination, error) {

	requestPb, pbErr := updateNotificationDestinationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var notificationDestinationPb notificationDestinationPb
	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", requestPb.Id)
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
		&notificationDestinationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := notificationDestinationFromPb(&notificationDestinationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just PersonalCompute API methods
type personalComputeImpl struct {
	client *client.DatabricksClient
}

func (a *personalComputeImpl) Delete(ctx context.Context, request DeletePersonalComputeSettingRequest) (*DeletePersonalComputeSettingResponse, error) {

	requestPb, pbErr := deletePersonalComputeSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deletePersonalComputeSettingResponsePb deletePersonalComputeSettingResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
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
		&deletePersonalComputeSettingResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deletePersonalComputeSettingResponseFromPb(&deletePersonalComputeSettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *personalComputeImpl) Get(ctx context.Context, request GetPersonalComputeSettingRequest) (*PersonalComputeSetting, error) {

	requestPb, pbErr := getPersonalComputeSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var personalComputeSettingPb personalComputeSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
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
		&personalComputeSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := personalComputeSettingFromPb(&personalComputeSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *personalComputeImpl) Update(ctx context.Context, request UpdatePersonalComputeSettingRequest) (*PersonalComputeSetting, error) {

	requestPb, pbErr := updatePersonalComputeSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var personalComputeSettingPb personalComputeSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
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
		&personalComputeSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := personalComputeSettingFromPb(&personalComputeSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just RestrictWorkspaceAdmins API methods
type restrictWorkspaceAdminsImpl struct {
	client *client.DatabricksClient
}

func (a *restrictWorkspaceAdminsImpl) Delete(ctx context.Context, request DeleteRestrictWorkspaceAdminsSettingRequest) (*DeleteRestrictWorkspaceAdminsSettingResponse, error) {

	requestPb, pbErr := deleteRestrictWorkspaceAdminsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteRestrictWorkspaceAdminsSettingResponsePb deleteRestrictWorkspaceAdminsSettingResponsePb
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
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
		&deleteRestrictWorkspaceAdminsSettingResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteRestrictWorkspaceAdminsSettingResponseFromPb(&deleteRestrictWorkspaceAdminsSettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *restrictWorkspaceAdminsImpl) Get(ctx context.Context, request GetRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error) {

	requestPb, pbErr := getRestrictWorkspaceAdminsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var restrictWorkspaceAdminsSettingPb restrictWorkspaceAdminsSettingPb
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
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
		&restrictWorkspaceAdminsSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := restrictWorkspaceAdminsSettingFromPb(&restrictWorkspaceAdminsSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *restrictWorkspaceAdminsImpl) Update(ctx context.Context, request UpdateRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error) {

	requestPb, pbErr := updateRestrictWorkspaceAdminsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var restrictWorkspaceAdminsSettingPb restrictWorkspaceAdminsSettingPb
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
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
		&restrictWorkspaceAdminsSettingPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := restrictWorkspaceAdminsSettingFromPb(&restrictWorkspaceAdminsSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Settings API methods
type settingsImpl struct {
	client *client.DatabricksClient
}

// unexported type that holds implementations of just SqlResultsDownload API methods
type sqlResultsDownloadImpl struct {
	client *client.DatabricksClient
}

func (a *sqlResultsDownloadImpl) Delete(ctx context.Context, request DeleteSqlResultsDownloadRequest) (*DeleteSqlResultsDownloadResponse, error) {

	requestPb, pbErr := deleteSqlResultsDownloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteSqlResultsDownloadResponsePb deleteSqlResultsDownloadResponsePb
	path := "/api/2.0/settings/types/sql_results_download/names/default"
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
		&deleteSqlResultsDownloadResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteSqlResultsDownloadResponseFromPb(&deleteSqlResultsDownloadResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *sqlResultsDownloadImpl) Get(ctx context.Context, request GetSqlResultsDownloadRequest) (*SqlResultsDownload, error) {

	requestPb, pbErr := getSqlResultsDownloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var sqlResultsDownloadPb sqlResultsDownloadPb
	path := "/api/2.0/settings/types/sql_results_download/names/default"
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
		&sqlResultsDownloadPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := sqlResultsDownloadFromPb(&sqlResultsDownloadPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *sqlResultsDownloadImpl) Update(ctx context.Context, request UpdateSqlResultsDownloadRequest) (*SqlResultsDownload, error) {

	requestPb, pbErr := updateSqlResultsDownloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var sqlResultsDownloadPb sqlResultsDownloadPb
	path := "/api/2.0/settings/types/sql_results_download/names/default"
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
		&sqlResultsDownloadPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := sqlResultsDownloadFromPb(&sqlResultsDownloadPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just TokenManagement API methods
type tokenManagementImpl struct {
	client *client.DatabricksClient
}

func (a *tokenManagementImpl) CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error) {

	requestPb, pbErr := createOboTokenRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createOboTokenResponsePb createOboTokenResponsePb
	path := "/api/2.0/token-management/on-behalf-of/tokens"
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
		&createOboTokenResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createOboTokenResponseFromPb(&createOboTokenResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) Delete(ctx context.Context, request DeleteTokenManagementRequest) error {

	requestPb, pbErr := deleteTokenManagementRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", requestPb.TokenId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *tokenManagementImpl) Get(ctx context.Context, request GetTokenManagementRequest) (*GetTokenResponse, error) {

	requestPb, pbErr := getTokenManagementRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getTokenResponsePb getTokenResponsePb
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", requestPb.TokenId)
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
		&getTokenResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getTokenResponseFromPb(&getTokenResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) GetPermissionLevels(ctx context.Context) (*GetTokenPermissionLevelsResponse, error) {

	var getTokenPermissionLevelsResponsePb getTokenPermissionLevelsResponsePb
	path := "/api/2.0/permissions/authorization/tokens/permissionLevels"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&getTokenPermissionLevelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getTokenPermissionLevelsResponseFromPb(&getTokenPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) GetPermissions(ctx context.Context) (*TokenPermissions, error) {

	var tokenPermissionsPb tokenPermissionsPb
	path := "/api/2.0/permissions/authorization/tokens"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&tokenPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := tokenPermissionsFromPb(&tokenPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists all tokens associated with the specified workspace or user.
func (a *tokenManagementImpl) List(ctx context.Context, request ListTokenManagementRequest) listing.Iterator[TokenInfo] {

	getNextPage := func(ctx context.Context, req ListTokenManagementRequest) (*ListTokensResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListTokensResponse) []TokenInfo {
		return resp.TokenInfos
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Lists all tokens associated with the specified workspace or user.
func (a *tokenManagementImpl) ListAll(ctx context.Context, request ListTokenManagementRequest) ([]TokenInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[TokenInfo](ctx, iterator)
}

func (a *tokenManagementImpl) internalList(ctx context.Context, request ListTokenManagementRequest) (*ListTokensResponse, error) {

	requestPb, pbErr := listTokenManagementRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listTokensResponsePb listTokensResponsePb
	path := "/api/2.0/token-management/tokens"
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
		&listTokensResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listTokensResponseFromPb(&listTokensResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) SetPermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error) {

	requestPb, pbErr := tokenPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var tokenPermissionsPb tokenPermissionsPb
	path := "/api/2.0/permissions/authorization/tokens"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&tokenPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := tokenPermissionsFromPb(&tokenPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) UpdatePermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error) {

	requestPb, pbErr := tokenPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var tokenPermissionsPb tokenPermissionsPb
	path := "/api/2.0/permissions/authorization/tokens"
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
		&tokenPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := tokenPermissionsFromPb(&tokenPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Tokens API methods
type tokensImpl struct {
	client *client.DatabricksClient
}

func (a *tokensImpl) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {

	requestPb, pbErr := createTokenRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createTokenResponsePb createTokenResponsePb
	path := "/api/2.0/token/create"
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
		&createTokenResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createTokenResponseFromPb(&createTokenResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokensImpl) Delete(ctx context.Context, request RevokeTokenRequest) error {

	requestPb, pbErr := revokeTokenRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/token/delete"
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

// Lists all the valid tokens for a user-workspace pair.
func (a *tokensImpl) List(ctx context.Context) listing.Iterator[PublicTokenInfo] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListPublicTokensResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListPublicTokensResponse) []PublicTokenInfo {
		return resp.TokenInfos
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Lists all the valid tokens for a user-workspace pair.
func (a *tokensImpl) ListAll(ctx context.Context) ([]PublicTokenInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[PublicTokenInfo](ctx, iterator)
}

func (a *tokensImpl) internalList(ctx context.Context) (*ListPublicTokensResponse, error) {

	var listPublicTokensResponsePb listPublicTokensResponsePb
	path := "/api/2.0/token/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&listPublicTokensResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listPublicTokensResponseFromPb(&listPublicTokensResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just WorkspaceConf API methods
type workspaceConfImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceConfImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*map[string]string, error) {

	requestPb, pbErr := getStatusRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceConfPb map[string]string
	path := "/api/2.0/workspace-conf"
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
		&workspaceConfPb,
	)
	if err != nil {
		return nil, err
	}
	resp := &workspaceConfPb

	return resp, err
}

func (a *workspaceConfImpl) SetStatus(ctx context.Context, request WorkspaceConf) error {

	requestPb, pbErr := workspaceConfToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/workspace-conf"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just WorkspaceNetworkConfiguration API methods
type workspaceNetworkConfigurationImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceNetworkConfigurationImpl) GetWorkspaceNetworkOptionRpc(ctx context.Context, request GetWorkspaceNetworkOptionRequest) (*WorkspaceNetworkOption, error) {

	requestPb, pbErr := getWorkspaceNetworkOptionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceNetworkOptionPb workspaceNetworkOptionPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/network", a.client.ConfiguredAccountID(), requestPb.WorkspaceId)
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
		&workspaceNetworkOptionPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := workspaceNetworkOptionFromPb(&workspaceNetworkOptionPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceNetworkConfigurationImpl) UpdateWorkspaceNetworkOptionRpc(ctx context.Context, request UpdateWorkspaceNetworkOptionRequest) (*WorkspaceNetworkOption, error) {

	requestPb, pbErr := updateWorkspaceNetworkOptionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceNetworkOptionPb workspaceNetworkOptionPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/network", a.client.ConfiguredAccountID(), requestPb.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb).WorkspaceNetworkOption,
		&workspaceNetworkOptionPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := workspaceNetworkOptionFromPb(&workspaceNetworkOptionPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
