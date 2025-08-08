// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/settings/settingspb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccountIpAccessLists API methods
type accountIpAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *accountIpAccessListsImpl) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {
	requestPb, pbErr := CreateIpAccessListToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createIpAccessListResponsePb settingspb.CreateIpAccessListResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists", a.client.ConfiguredAccountID())
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
		&createIpAccessListResponsePb,
	)
	resp, err := CreateIpAccessListResponseFromPb(&createIpAccessListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountIpAccessListsImpl) Delete(ctx context.Context, request DeleteAccountIpAccessListRequest) error {
	requestPb, pbErr := DeleteAccountIpAccessListRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
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

func (a *accountIpAccessListsImpl) Get(ctx context.Context, request GetAccountIpAccessListRequest) (*GetIpAccessListResponse, error) {
	requestPb, pbErr := GetAccountIpAccessListRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getIpAccessListResponsePb settingspb.GetIpAccessListResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getIpAccessListResponsePb,
	)
	resp, err := GetIpAccessListResponseFromPb(&getIpAccessListResponsePb)
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

	var getIpAccessListsResponsePb settingspb.GetIpAccessListsResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := GetIpAccessListsResponseFromPb(&getIpAccessListsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountIpAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	requestPb, pbErr := ReplaceIpAccessListToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
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

func (a *accountIpAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {
	requestPb, pbErr := UpdateIpAccessListToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
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

// unexported type that holds implementations of just AccountSettings API methods
type accountSettingsImpl struct {
	client *client.DatabricksClient
}

// unexported type that holds implementations of just AibiDashboardEmbeddingAccessPolicy API methods
type aibiDashboardEmbeddingAccessPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *aibiDashboardEmbeddingAccessPolicyImpl) Delete(ctx context.Context, request DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error) {
	requestPb, pbErr := DeleteAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingResponsePb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb,
	)
	resp, err := DeleteAibiDashboardEmbeddingAccessPolicySettingResponseFromPb(&deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aibiDashboardEmbeddingAccessPolicyImpl) Get(ctx context.Context, request GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error) {
	requestPb, pbErr := GetAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aibiDashboardEmbeddingAccessPolicySettingPb settingspb.AibiDashboardEmbeddingAccessPolicySettingPb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&aibiDashboardEmbeddingAccessPolicySettingPb,
	)
	resp, err := AibiDashboardEmbeddingAccessPolicySettingFromPb(&aibiDashboardEmbeddingAccessPolicySettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aibiDashboardEmbeddingAccessPolicyImpl) Update(ctx context.Context, request UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error) {
	requestPb, pbErr := UpdateAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aibiDashboardEmbeddingAccessPolicySettingPb settingspb.AibiDashboardEmbeddingAccessPolicySettingPb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
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
		&aibiDashboardEmbeddingAccessPolicySettingPb,
	)
	resp, err := AibiDashboardEmbeddingAccessPolicySettingFromPb(&aibiDashboardEmbeddingAccessPolicySettingPb)
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
	requestPb, pbErr := DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb,
	)
	resp, err := DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponseFromPb(&deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aibiDashboardEmbeddingApprovedDomainsImpl) Get(ctx context.Context, request GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	requestPb, pbErr := GetAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aibiDashboardEmbeddingApprovedDomainsSettingPb settingspb.AibiDashboardEmbeddingApprovedDomainsSettingPb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&aibiDashboardEmbeddingApprovedDomainsSettingPb,
	)
	resp, err := AibiDashboardEmbeddingApprovedDomainsSettingFromPb(&aibiDashboardEmbeddingApprovedDomainsSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aibiDashboardEmbeddingApprovedDomainsImpl) Update(ctx context.Context, request UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	requestPb, pbErr := UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aibiDashboardEmbeddingApprovedDomainsSettingPb settingspb.AibiDashboardEmbeddingApprovedDomainsSettingPb
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
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
		&aibiDashboardEmbeddingApprovedDomainsSettingPb,
	)
	resp, err := AibiDashboardEmbeddingApprovedDomainsSettingFromPb(&aibiDashboardEmbeddingApprovedDomainsSettingPb)
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
	requestPb, pbErr := GetAutomaticClusterUpdateSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var automaticClusterUpdateSettingPb settingspb.AutomaticClusterUpdateSettingPb
	path := "/api/2.0/settings/types/automatic_cluster_update/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&automaticClusterUpdateSettingPb,
	)
	resp, err := AutomaticClusterUpdateSettingFromPb(&automaticClusterUpdateSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *automaticClusterUpdateImpl) Update(ctx context.Context, request UpdateAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error) {
	requestPb, pbErr := UpdateAutomaticClusterUpdateSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var automaticClusterUpdateSettingPb settingspb.AutomaticClusterUpdateSettingPb
	path := "/api/2.0/settings/types/automatic_cluster_update/names/default"
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
		&automaticClusterUpdateSettingPb,
	)
	resp, err := AutomaticClusterUpdateSettingFromPb(&automaticClusterUpdateSettingPb)
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
	requestPb, pbErr := GetComplianceSecurityProfileSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var complianceSecurityProfileSettingPb settingspb.ComplianceSecurityProfileSettingPb
	path := "/api/2.0/settings/types/shield_csp_enablement_ws_db/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&complianceSecurityProfileSettingPb,
	)
	resp, err := ComplianceSecurityProfileSettingFromPb(&complianceSecurityProfileSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *complianceSecurityProfileImpl) Update(ctx context.Context, request UpdateComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error) {
	requestPb, pbErr := UpdateComplianceSecurityProfileSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var complianceSecurityProfileSettingPb settingspb.ComplianceSecurityProfileSettingPb
	path := "/api/2.0/settings/types/shield_csp_enablement_ws_db/names/default"
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
		&complianceSecurityProfileSettingPb,
	)
	resp, err := ComplianceSecurityProfileSettingFromPb(&complianceSecurityProfileSettingPb)
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
	requestPb, pbErr := ExchangeTokenRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var exchangeTokenResponsePb settingspb.ExchangeTokenResponsePb
	path := "/api/2.0/credentials-manager/exchange-tokens/token"
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
		&exchangeTokenResponsePb,
	)
	resp, err := ExchangeTokenResponseFromPb(&exchangeTokenResponsePb)
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
	requestPb, pbErr := GetCspEnablementAccountSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cspEnablementAccountSettingPb settingspb.CspEnablementAccountSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_csp_enablement_ac/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&cspEnablementAccountSettingPb,
	)
	resp, err := CspEnablementAccountSettingFromPb(&cspEnablementAccountSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cspEnablementAccountImpl) Update(ctx context.Context, request UpdateCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error) {
	requestPb, pbErr := UpdateCspEnablementAccountSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cspEnablementAccountSettingPb settingspb.CspEnablementAccountSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_csp_enablement_ac/names/default", a.client.ConfiguredAccountID())
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
		&cspEnablementAccountSettingPb,
	)
	resp, err := CspEnablementAccountSettingFromPb(&cspEnablementAccountSettingPb)
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
	requestPb, pbErr := DeleteDashboardEmailSubscriptionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDashboardEmailSubscriptionsResponsePb settingspb.DeleteDashboardEmailSubscriptionsResponsePb
	path := "/api/2.0/settings/types/dashboard_email_subscriptions/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteDashboardEmailSubscriptionsResponsePb,
	)
	resp, err := DeleteDashboardEmailSubscriptionsResponseFromPb(&deleteDashboardEmailSubscriptionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardEmailSubscriptionsImpl) Get(ctx context.Context, request GetDashboardEmailSubscriptionsRequest) (*DashboardEmailSubscriptions, error) {
	requestPb, pbErr := GetDashboardEmailSubscriptionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardEmailSubscriptionsPb settingspb.DashboardEmailSubscriptionsPb
	path := "/api/2.0/settings/types/dashboard_email_subscriptions/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&dashboardEmailSubscriptionsPb,
	)
	resp, err := DashboardEmailSubscriptionsFromPb(&dashboardEmailSubscriptionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardEmailSubscriptionsImpl) Update(ctx context.Context, request UpdateDashboardEmailSubscriptionsRequest) (*DashboardEmailSubscriptions, error) {
	requestPb, pbErr := UpdateDashboardEmailSubscriptionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardEmailSubscriptionsPb settingspb.DashboardEmailSubscriptionsPb
	path := "/api/2.0/settings/types/dashboard_email_subscriptions/names/default"
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
		&dashboardEmailSubscriptionsPb,
	)
	resp, err := DashboardEmailSubscriptionsFromPb(&dashboardEmailSubscriptionsPb)
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
	requestPb, pbErr := DeleteDefaultNamespaceSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDefaultNamespaceSettingResponsePb settingspb.DeleteDefaultNamespaceSettingResponsePb
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteDefaultNamespaceSettingResponsePb,
	)
	resp, err := DeleteDefaultNamespaceSettingResponseFromPb(&deleteDefaultNamespaceSettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *defaultNamespaceImpl) Get(ctx context.Context, request GetDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error) {
	requestPb, pbErr := GetDefaultNamespaceSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var defaultNamespaceSettingPb settingspb.DefaultNamespaceSettingPb
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&defaultNamespaceSettingPb,
	)
	resp, err := DefaultNamespaceSettingFromPb(&defaultNamespaceSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *defaultNamespaceImpl) Update(ctx context.Context, request UpdateDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error) {
	requestPb, pbErr := UpdateDefaultNamespaceSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var defaultNamespaceSettingPb settingspb.DefaultNamespaceSettingPb
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
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
		&defaultNamespaceSettingPb,
	)
	resp, err := DefaultNamespaceSettingFromPb(&defaultNamespaceSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just DefaultWarehouseId API methods
type defaultWarehouseIdImpl struct {
	client *client.DatabricksClient
}

func (a *defaultWarehouseIdImpl) Delete(ctx context.Context, request DeleteDefaultWarehouseIdRequest) (*DeleteDefaultWarehouseIdResponse, error) {
	requestPb, pbErr := DeleteDefaultWarehouseIdRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDefaultWarehouseIdResponsePb settingspb.DeleteDefaultWarehouseIdResponsePb
	path := "/api/2.0/settings/types/default_warehouse_id/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteDefaultWarehouseIdResponsePb,
	)
	resp, err := DeleteDefaultWarehouseIdResponseFromPb(&deleteDefaultWarehouseIdResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *defaultWarehouseIdImpl) Get(ctx context.Context, request GetDefaultWarehouseIdRequest) (*DefaultWarehouseId, error) {
	requestPb, pbErr := GetDefaultWarehouseIdRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var defaultWarehouseIdPb settingspb.DefaultWarehouseIdPb
	path := "/api/2.0/settings/types/default_warehouse_id/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&defaultWarehouseIdPb,
	)
	resp, err := DefaultWarehouseIdFromPb(&defaultWarehouseIdPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *defaultWarehouseIdImpl) Update(ctx context.Context, request UpdateDefaultWarehouseIdRequest) (*DefaultWarehouseId, error) {
	requestPb, pbErr := UpdateDefaultWarehouseIdRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var defaultWarehouseIdPb settingspb.DefaultWarehouseIdPb
	path := "/api/2.0/settings/types/default_warehouse_id/names/default"
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
		&defaultWarehouseIdPb,
	)
	resp, err := DefaultWarehouseIdFromPb(&defaultWarehouseIdPb)
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
	requestPb, pbErr := DeleteDisableLegacyAccessRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDisableLegacyAccessResponsePb settingspb.DeleteDisableLegacyAccessResponsePb
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteDisableLegacyAccessResponsePb,
	)
	resp, err := DeleteDisableLegacyAccessResponseFromPb(&deleteDisableLegacyAccessResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyAccessImpl) Get(ctx context.Context, request GetDisableLegacyAccessRequest) (*DisableLegacyAccess, error) {
	requestPb, pbErr := GetDisableLegacyAccessRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyAccessPb settingspb.DisableLegacyAccessPb
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&disableLegacyAccessPb,
	)
	resp, err := DisableLegacyAccessFromPb(&disableLegacyAccessPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyAccessImpl) Update(ctx context.Context, request UpdateDisableLegacyAccessRequest) (*DisableLegacyAccess, error) {
	requestPb, pbErr := UpdateDisableLegacyAccessRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyAccessPb settingspb.DisableLegacyAccessPb
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
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
		&disableLegacyAccessPb,
	)
	resp, err := DisableLegacyAccessFromPb(&disableLegacyAccessPb)
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
	requestPb, pbErr := DeleteDisableLegacyDbfsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDisableLegacyDbfsResponsePb settingspb.DeleteDisableLegacyDbfsResponsePb
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteDisableLegacyDbfsResponsePb,
	)
	resp, err := DeleteDisableLegacyDbfsResponseFromPb(&deleteDisableLegacyDbfsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyDbfsImpl) Get(ctx context.Context, request GetDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error) {
	requestPb, pbErr := GetDisableLegacyDbfsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyDbfsPb settingspb.DisableLegacyDbfsPb
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&disableLegacyDbfsPb,
	)
	resp, err := DisableLegacyDbfsFromPb(&disableLegacyDbfsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyDbfsImpl) Update(ctx context.Context, request UpdateDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error) {
	requestPb, pbErr := UpdateDisableLegacyDbfsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyDbfsPb settingspb.DisableLegacyDbfsPb
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
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
		&disableLegacyDbfsPb,
	)
	resp, err := DisableLegacyDbfsFromPb(&disableLegacyDbfsPb)
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
	requestPb, pbErr := DeleteDisableLegacyFeaturesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDisableLegacyFeaturesResponsePb settingspb.DeleteDisableLegacyFeaturesResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteDisableLegacyFeaturesResponsePb,
	)
	resp, err := DeleteDisableLegacyFeaturesResponseFromPb(&deleteDisableLegacyFeaturesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyFeaturesImpl) Get(ctx context.Context, request GetDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error) {
	requestPb, pbErr := GetDisableLegacyFeaturesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyFeaturesPb settingspb.DisableLegacyFeaturesPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&disableLegacyFeaturesPb,
	)
	resp, err := DisableLegacyFeaturesFromPb(&disableLegacyFeaturesPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *disableLegacyFeaturesImpl) Update(ctx context.Context, request UpdateDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error) {
	requestPb, pbErr := UpdateDisableLegacyFeaturesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var disableLegacyFeaturesPb settingspb.DisableLegacyFeaturesPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
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
		&disableLegacyFeaturesPb,
	)
	resp, err := DisableLegacyFeaturesFromPb(&disableLegacyFeaturesPb)
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

	var enableExportNotebookPb settingspb.EnableExportNotebookPb
	path := "/api/2.0/settings/types/enable-export-notebook/names/default"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&enableExportNotebookPb,
	)
	resp, err := EnableExportNotebookFromPb(&enableExportNotebookPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableExportNotebookImpl) PatchEnableExportNotebook(ctx context.Context, request UpdateEnableExportNotebookRequest) (*EnableExportNotebook, error) {
	requestPb, pbErr := UpdateEnableExportNotebookRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enableExportNotebookPb settingspb.EnableExportNotebookPb
	path := "/api/2.0/settings/types/enable-export-notebook/names/default"
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
		&enableExportNotebookPb,
	)
	resp, err := EnableExportNotebookFromPb(&enableExportNotebookPb)
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
	requestPb, pbErr := DeleteAccountIpAccessEnableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteAccountIpAccessEnableResponsePb settingspb.DeleteAccountIpAccessEnableResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/acct_ip_acl_enable/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteAccountIpAccessEnableResponsePb,
	)
	resp, err := DeleteAccountIpAccessEnableResponseFromPb(&deleteAccountIpAccessEnableResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableIpAccessListsImpl) Get(ctx context.Context, request GetAccountIpAccessEnableRequest) (*AccountIpAccessEnable, error) {
	requestPb, pbErr := GetAccountIpAccessEnableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountIpAccessEnablePb settingspb.AccountIpAccessEnablePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/acct_ip_acl_enable/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&accountIpAccessEnablePb,
	)
	resp, err := AccountIpAccessEnableFromPb(&accountIpAccessEnablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableIpAccessListsImpl) Update(ctx context.Context, request UpdateAccountIpAccessEnableRequest) (*AccountIpAccessEnable, error) {
	requestPb, pbErr := UpdateAccountIpAccessEnableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountIpAccessEnablePb settingspb.AccountIpAccessEnablePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/acct_ip_acl_enable/names/default", a.client.ConfiguredAccountID())
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
		&accountIpAccessEnablePb,
	)
	resp, err := AccountIpAccessEnableFromPb(&accountIpAccessEnablePb)
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

	var enableNotebookTableClipboardPb settingspb.EnableNotebookTableClipboardPb
	path := "/api/2.0/settings/types/enable-notebook-table-clipboard/names/default"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&enableNotebookTableClipboardPb,
	)
	resp, err := EnableNotebookTableClipboardFromPb(&enableNotebookTableClipboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableNotebookTableClipboardImpl) PatchEnableNotebookTableClipboard(ctx context.Context, request UpdateEnableNotebookTableClipboardRequest) (*EnableNotebookTableClipboard, error) {
	requestPb, pbErr := UpdateEnableNotebookTableClipboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enableNotebookTableClipboardPb settingspb.EnableNotebookTableClipboardPb
	path := "/api/2.0/settings/types/enable-notebook-table-clipboard/names/default"
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
		&enableNotebookTableClipboardPb,
	)
	resp, err := EnableNotebookTableClipboardFromPb(&enableNotebookTableClipboardPb)
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

	var enableResultsDownloadingPb settingspb.EnableResultsDownloadingPb
	path := "/api/2.0/settings/types/enable-results-downloading/names/default"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&enableResultsDownloadingPb,
	)
	resp, err := EnableResultsDownloadingFromPb(&enableResultsDownloadingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enableResultsDownloadingImpl) PatchEnableResultsDownloading(ctx context.Context, request UpdateEnableResultsDownloadingRequest) (*EnableResultsDownloading, error) {
	requestPb, pbErr := UpdateEnableResultsDownloadingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enableResultsDownloadingPb settingspb.EnableResultsDownloadingPb
	path := "/api/2.0/settings/types/enable-results-downloading/names/default"
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
		&enableResultsDownloadingPb,
	)
	resp, err := EnableResultsDownloadingFromPb(&enableResultsDownloadingPb)
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
	requestPb, pbErr := GetEnhancedSecurityMonitoringSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enhancedSecurityMonitoringSettingPb settingspb.EnhancedSecurityMonitoringSettingPb
	path := "/api/2.0/settings/types/shield_esm_enablement_ws_db/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&enhancedSecurityMonitoringSettingPb,
	)
	resp, err := EnhancedSecurityMonitoringSettingFromPb(&enhancedSecurityMonitoringSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *enhancedSecurityMonitoringImpl) Update(ctx context.Context, request UpdateEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error) {
	requestPb, pbErr := UpdateEnhancedSecurityMonitoringSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enhancedSecurityMonitoringSettingPb settingspb.EnhancedSecurityMonitoringSettingPb
	path := "/api/2.0/settings/types/shield_esm_enablement_ws_db/names/default"
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
		&enhancedSecurityMonitoringSettingPb,
	)
	resp, err := EnhancedSecurityMonitoringSettingFromPb(&enhancedSecurityMonitoringSettingPb)
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
	requestPb, pbErr := GetEsmEnablementAccountSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var esmEnablementAccountSettingPb settingspb.EsmEnablementAccountSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_esm_enablement_ac/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&esmEnablementAccountSettingPb,
	)
	resp, err := EsmEnablementAccountSettingFromPb(&esmEnablementAccountSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *esmEnablementAccountImpl) Update(ctx context.Context, request UpdateEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error) {
	requestPb, pbErr := UpdateEsmEnablementAccountSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var esmEnablementAccountSettingPb settingspb.EsmEnablementAccountSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_esm_enablement_ac/names/default", a.client.ConfiguredAccountID())
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
		&esmEnablementAccountSettingPb,
	)
	resp, err := EsmEnablementAccountSettingFromPb(&esmEnablementAccountSettingPb)
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
	requestPb, pbErr := CreateIpAccessListToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createIpAccessListResponsePb settingspb.CreateIpAccessListResponsePb
	path := "/api/2.0/ip-access-lists"
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
		&createIpAccessListResponsePb,
	)
	resp, err := CreateIpAccessListResponseFromPb(&createIpAccessListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *ipAccessListsImpl) Delete(ctx context.Context, request DeleteIpAccessListRequest) error {
	requestPb, pbErr := DeleteIpAccessListRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
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

func (a *ipAccessListsImpl) Get(ctx context.Context, request GetIpAccessListRequest) (*FetchIpAccessListResponse, error) {
	requestPb, pbErr := GetIpAccessListRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var fetchIpAccessListResponsePb settingspb.FetchIpAccessListResponsePb
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&fetchIpAccessListResponsePb,
	)
	resp, err := FetchIpAccessListResponseFromPb(&fetchIpAccessListResponsePb)
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

	var listIpAccessListResponsePb settingspb.ListIpAccessListResponsePb
	path := "/api/2.0/ip-access-lists"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListIpAccessListResponseFromPb(&listIpAccessListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *ipAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	requestPb, pbErr := ReplaceIpAccessListToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
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

func (a *ipAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {
	requestPb, pbErr := UpdateIpAccessListToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
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

// unexported type that holds implementations of just LlmProxyPartnerPoweredAccount API methods
type llmProxyPartnerPoweredAccountImpl struct {
	client *client.DatabricksClient
}

func (a *llmProxyPartnerPoweredAccountImpl) Get(ctx context.Context, request GetLlmProxyPartnerPoweredAccountRequest) (*LlmProxyPartnerPoweredAccount, error) {
	requestPb, pbErr := GetLlmProxyPartnerPoweredAccountRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredAccountPb settingspb.LlmProxyPartnerPoweredAccountPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/llm_proxy_partner_powered/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&llmProxyPartnerPoweredAccountPb,
	)
	resp, err := LlmProxyPartnerPoweredAccountFromPb(&llmProxyPartnerPoweredAccountPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *llmProxyPartnerPoweredAccountImpl) Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredAccountRequest) (*LlmProxyPartnerPoweredAccount, error) {
	requestPb, pbErr := UpdateLlmProxyPartnerPoweredAccountRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredAccountPb settingspb.LlmProxyPartnerPoweredAccountPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/llm_proxy_partner_powered/names/default", a.client.ConfiguredAccountID())
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
		&llmProxyPartnerPoweredAccountPb,
	)
	resp, err := LlmProxyPartnerPoweredAccountFromPb(&llmProxyPartnerPoweredAccountPb)
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
	requestPb, pbErr := GetLlmProxyPartnerPoweredEnforceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredEnforcePb settingspb.LlmProxyPartnerPoweredEnforcePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/llm_proxy_partner_powered_enforce/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&llmProxyPartnerPoweredEnforcePb,
	)
	resp, err := LlmProxyPartnerPoweredEnforceFromPb(&llmProxyPartnerPoweredEnforcePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *llmProxyPartnerPoweredEnforceImpl) Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredEnforceRequest) (*LlmProxyPartnerPoweredEnforce, error) {
	requestPb, pbErr := UpdateLlmProxyPartnerPoweredEnforceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredEnforcePb settingspb.LlmProxyPartnerPoweredEnforcePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/llm_proxy_partner_powered_enforce/names/default", a.client.ConfiguredAccountID())
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
		&llmProxyPartnerPoweredEnforcePb,
	)
	resp, err := LlmProxyPartnerPoweredEnforceFromPb(&llmProxyPartnerPoweredEnforcePb)
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
	requestPb, pbErr := DeleteLlmProxyPartnerPoweredWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteLlmProxyPartnerPoweredWorkspaceResponsePb settingspb.DeleteLlmProxyPartnerPoweredWorkspaceResponsePb
	path := "/api/2.0/settings/types/llm_proxy_partner_powered/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteLlmProxyPartnerPoweredWorkspaceResponsePb,
	)
	resp, err := DeleteLlmProxyPartnerPoweredWorkspaceResponseFromPb(&deleteLlmProxyPartnerPoweredWorkspaceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *llmProxyPartnerPoweredWorkspaceImpl) Get(ctx context.Context, request GetLlmProxyPartnerPoweredWorkspaceRequest) (*LlmProxyPartnerPoweredWorkspace, error) {
	requestPb, pbErr := GetLlmProxyPartnerPoweredWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredWorkspacePb settingspb.LlmProxyPartnerPoweredWorkspacePb
	path := "/api/2.0/settings/types/llm_proxy_partner_powered/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&llmProxyPartnerPoweredWorkspacePb,
	)
	resp, err := LlmProxyPartnerPoweredWorkspaceFromPb(&llmProxyPartnerPoweredWorkspacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *llmProxyPartnerPoweredWorkspaceImpl) Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredWorkspaceRequest) (*LlmProxyPartnerPoweredWorkspace, error) {
	requestPb, pbErr := UpdateLlmProxyPartnerPoweredWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var llmProxyPartnerPoweredWorkspacePb settingspb.LlmProxyPartnerPoweredWorkspacePb
	path := "/api/2.0/settings/types/llm_proxy_partner_powered/names/default"
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
		&llmProxyPartnerPoweredWorkspacePb,
	)
	resp, err := LlmProxyPartnerPoweredWorkspaceFromPb(&llmProxyPartnerPoweredWorkspacePb)
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
	requestPb, pbErr := CreateNetworkConnectivityConfigRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var networkConnectivityConfigurationPb settingspb.NetworkConnectivityConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.NetworkConnectivityConfig,
		&networkConnectivityConfigurationPb,
	)
	resp, err := NetworkConnectivityConfigurationFromPb(&networkConnectivityConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) CreatePrivateEndpointRule(ctx context.Context, request CreatePrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error) {
	requestPb, pbErr := CreatePrivateEndpointRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var nccPrivateEndpointRulePb settingspb.NccPrivateEndpointRulePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.PrivateEndpointRule,
		&nccPrivateEndpointRulePb,
	)
	resp, err := NccPrivateEndpointRuleFromPb(&nccPrivateEndpointRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) DeleteNetworkConnectivityConfiguration(ctx context.Context, request DeleteNetworkConnectivityConfigurationRequest) error {
	requestPb, pbErr := DeleteNetworkConnectivityConfigurationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId)
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

func (a *networkConnectivityImpl) DeletePrivateEndpointRule(ctx context.Context, request DeletePrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error) {
	requestPb, pbErr := DeletePrivateEndpointRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var nccPrivateEndpointRulePb settingspb.NccPrivateEndpointRulePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId, request.PrivateEndpointRuleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&nccPrivateEndpointRulePb,
	)
	resp, err := NccPrivateEndpointRuleFromPb(&nccPrivateEndpointRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) GetNetworkConnectivityConfiguration(ctx context.Context, request GetNetworkConnectivityConfigurationRequest) (*NetworkConnectivityConfiguration, error) {
	requestPb, pbErr := GetNetworkConnectivityConfigurationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var networkConnectivityConfigurationPb settingspb.NetworkConnectivityConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&networkConnectivityConfigurationPb,
	)
	resp, err := NetworkConnectivityConfigurationFromPb(&networkConnectivityConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) GetPrivateEndpointRule(ctx context.Context, request GetPrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error) {
	requestPb, pbErr := GetPrivateEndpointRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var nccPrivateEndpointRulePb settingspb.NccPrivateEndpointRulePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId, request.PrivateEndpointRuleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&nccPrivateEndpointRulePb,
	)
	resp, err := NccPrivateEndpointRuleFromPb(&nccPrivateEndpointRulePb)
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
	requestPb, pbErr := ListNetworkConnectivityConfigurationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listNetworkConnectivityConfigurationsResponsePb settingspb.ListNetworkConnectivityConfigurationsResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listNetworkConnectivityConfigurationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListNetworkConnectivityConfigurationsResponseFromPb(&listNetworkConnectivityConfigurationsResponsePb)
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
	requestPb, pbErr := ListPrivateEndpointRulesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listPrivateEndpointRulesResponsePb settingspb.ListPrivateEndpointRulesResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listPrivateEndpointRulesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListPrivateEndpointRulesResponseFromPb(&listPrivateEndpointRulesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkConnectivityImpl) UpdatePrivateEndpointRule(ctx context.Context, request UpdateNccPrivateEndpointRuleRequest) (*NccPrivateEndpointRule, error) {
	requestPb, pbErr := UpdateNccPrivateEndpointRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var nccPrivateEndpointRulePb settingspb.NccPrivateEndpointRulePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId, request.PrivateEndpointRuleId)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.PrivateEndpointRule,
		&nccPrivateEndpointRulePb,
	)
	resp, err := NccPrivateEndpointRuleFromPb(&nccPrivateEndpointRulePb)
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
	requestPb, pbErr := CreateNetworkPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountNetworkPolicyPb settingspb.AccountNetworkPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.NetworkPolicy,
		&accountNetworkPolicyPb,
	)
	resp, err := AccountNetworkPolicyFromPb(&accountNetworkPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkPoliciesImpl) DeleteNetworkPolicyRpc(ctx context.Context, request DeleteNetworkPolicyRequest) error {
	requestPb, pbErr := DeleteNetworkPolicyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies/%v", a.client.ConfiguredAccountID(), request.NetworkPolicyId)
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

func (a *networkPoliciesImpl) GetNetworkPolicyRpc(ctx context.Context, request GetNetworkPolicyRequest) (*AccountNetworkPolicy, error) {
	requestPb, pbErr := GetNetworkPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountNetworkPolicyPb settingspb.AccountNetworkPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies/%v", a.client.ConfiguredAccountID(), request.NetworkPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&accountNetworkPolicyPb,
	)
	resp, err := AccountNetworkPolicyFromPb(&accountNetworkPolicyPb)
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
	requestPb, pbErr := ListNetworkPoliciesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listNetworkPoliciesResponsePb settingspb.ListNetworkPoliciesResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listNetworkPoliciesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListNetworkPoliciesResponseFromPb(&listNetworkPoliciesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networkPoliciesImpl) UpdateNetworkPolicyRpc(ctx context.Context, request UpdateNetworkPolicyRequest) (*AccountNetworkPolicy, error) {
	requestPb, pbErr := UpdateNetworkPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountNetworkPolicyPb settingspb.AccountNetworkPolicyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-policies/%v", a.client.ConfiguredAccountID(), request.NetworkPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb.NetworkPolicy,
		&accountNetworkPolicyPb,
	)
	resp, err := AccountNetworkPolicyFromPb(&accountNetworkPolicyPb)
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
	requestPb, pbErr := CreateNotificationDestinationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var notificationDestinationPb settingspb.NotificationDestinationPb
	path := "/api/2.0/notification-destinations"
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
		&notificationDestinationPb,
	)
	resp, err := NotificationDestinationFromPb(&notificationDestinationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *notificationDestinationsImpl) Delete(ctx context.Context, request DeleteNotificationDestinationRequest) error {
	requestPb, pbErr := DeleteNotificationDestinationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", request.Id)
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

func (a *notificationDestinationsImpl) Get(ctx context.Context, request GetNotificationDestinationRequest) (*NotificationDestination, error) {
	requestPb, pbErr := GetNotificationDestinationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var notificationDestinationPb settingspb.NotificationDestinationPb
	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&notificationDestinationPb,
	)
	resp, err := NotificationDestinationFromPb(&notificationDestinationPb)
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
	requestPb, pbErr := ListNotificationDestinationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listNotificationDestinationsResponsePb settingspb.ListNotificationDestinationsResponsePb
	path := "/api/2.0/notification-destinations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listNotificationDestinationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListNotificationDestinationsResponseFromPb(&listNotificationDestinationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *notificationDestinationsImpl) Update(ctx context.Context, request UpdateNotificationDestinationRequest) (*NotificationDestination, error) {
	requestPb, pbErr := UpdateNotificationDestinationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var notificationDestinationPb settingspb.NotificationDestinationPb
	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", request.Id)
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
		&notificationDestinationPb,
	)
	resp, err := NotificationDestinationFromPb(&notificationDestinationPb)
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
	requestPb, pbErr := DeletePersonalComputeSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deletePersonalComputeSettingResponsePb settingspb.DeletePersonalComputeSettingResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deletePersonalComputeSettingResponsePb,
	)
	resp, err := DeletePersonalComputeSettingResponseFromPb(&deletePersonalComputeSettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *personalComputeImpl) Get(ctx context.Context, request GetPersonalComputeSettingRequest) (*PersonalComputeSetting, error) {
	requestPb, pbErr := GetPersonalComputeSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var personalComputeSettingPb settingspb.PersonalComputeSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&personalComputeSettingPb,
	)
	resp, err := PersonalComputeSettingFromPb(&personalComputeSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *personalComputeImpl) Update(ctx context.Context, request UpdatePersonalComputeSettingRequest) (*PersonalComputeSetting, error) {
	requestPb, pbErr := UpdatePersonalComputeSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var personalComputeSettingPb settingspb.PersonalComputeSettingPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
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
		&personalComputeSettingPb,
	)
	resp, err := PersonalComputeSettingFromPb(&personalComputeSettingPb)
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
	requestPb, pbErr := DeleteRestrictWorkspaceAdminsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteRestrictWorkspaceAdminsSettingResponsePb settingspb.DeleteRestrictWorkspaceAdminsSettingResponsePb
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteRestrictWorkspaceAdminsSettingResponsePb,
	)
	resp, err := DeleteRestrictWorkspaceAdminsSettingResponseFromPb(&deleteRestrictWorkspaceAdminsSettingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *restrictWorkspaceAdminsImpl) Get(ctx context.Context, request GetRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error) {
	requestPb, pbErr := GetRestrictWorkspaceAdminsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var restrictWorkspaceAdminsSettingPb settingspb.RestrictWorkspaceAdminsSettingPb
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&restrictWorkspaceAdminsSettingPb,
	)
	resp, err := RestrictWorkspaceAdminsSettingFromPb(&restrictWorkspaceAdminsSettingPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *restrictWorkspaceAdminsImpl) Update(ctx context.Context, request UpdateRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error) {
	requestPb, pbErr := UpdateRestrictWorkspaceAdminsSettingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var restrictWorkspaceAdminsSettingPb settingspb.RestrictWorkspaceAdminsSettingPb
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
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
		&restrictWorkspaceAdminsSettingPb,
	)
	resp, err := RestrictWorkspaceAdminsSettingFromPb(&restrictWorkspaceAdminsSettingPb)
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
	requestPb, pbErr := DeleteSqlResultsDownloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteSqlResultsDownloadResponsePb settingspb.DeleteSqlResultsDownloadResponsePb
	path := "/api/2.0/settings/types/sql_results_download/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteSqlResultsDownloadResponsePb,
	)
	resp, err := DeleteSqlResultsDownloadResponseFromPb(&deleteSqlResultsDownloadResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *sqlResultsDownloadImpl) Get(ctx context.Context, request GetSqlResultsDownloadRequest) (*SqlResultsDownload, error) {
	requestPb, pbErr := GetSqlResultsDownloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var sqlResultsDownloadPb settingspb.SqlResultsDownloadPb
	path := "/api/2.0/settings/types/sql_results_download/names/default"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&sqlResultsDownloadPb,
	)
	resp, err := SqlResultsDownloadFromPb(&sqlResultsDownloadPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *sqlResultsDownloadImpl) Update(ctx context.Context, request UpdateSqlResultsDownloadRequest) (*SqlResultsDownload, error) {
	requestPb, pbErr := UpdateSqlResultsDownloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var sqlResultsDownloadPb settingspb.SqlResultsDownloadPb
	path := "/api/2.0/settings/types/sql_results_download/names/default"
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
		&sqlResultsDownloadPb,
	)
	resp, err := SqlResultsDownloadFromPb(&sqlResultsDownloadPb)
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
	requestPb, pbErr := CreateOboTokenRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createOboTokenResponsePb settingspb.CreateOboTokenResponsePb
	path := "/api/2.0/token-management/on-behalf-of/tokens"
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
		&createOboTokenResponsePb,
	)
	resp, err := CreateOboTokenResponseFromPb(&createOboTokenResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) Delete(ctx context.Context, request DeleteTokenManagementRequest) error {
	requestPb, pbErr := DeleteTokenManagementRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
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

func (a *tokenManagementImpl) Get(ctx context.Context, request GetTokenManagementRequest) (*GetTokenResponse, error) {
	requestPb, pbErr := GetTokenManagementRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getTokenResponsePb settingspb.GetTokenResponsePb
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getTokenResponsePb,
	)
	resp, err := GetTokenResponseFromPb(&getTokenResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) GetPermissionLevels(ctx context.Context) (*GetTokenPermissionLevelsResponse, error) {

	var getTokenPermissionLevelsResponsePb settingspb.GetTokenPermissionLevelsResponsePb
	path := "/api/2.0/permissions/authorization/tokens/permissionLevels"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&getTokenPermissionLevelsResponsePb,
	)
	resp, err := GetTokenPermissionLevelsResponseFromPb(&getTokenPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) GetPermissions(ctx context.Context) (*TokenPermissions, error) {

	var tokenPermissionsPb settingspb.TokenPermissionsPb
	path := "/api/2.0/permissions/authorization/tokens"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&tokenPermissionsPb,
	)
	resp, err := TokenPermissionsFromPb(&tokenPermissionsPb)
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
	requestPb, pbErr := ListTokenManagementRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listTokensResponsePb settingspb.ListTokensResponsePb
	path := "/api/2.0/token-management/tokens"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listTokensResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListTokensResponseFromPb(&listTokensResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) SetPermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error) {
	requestPb, pbErr := TokenPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var tokenPermissionsPb settingspb.TokenPermissionsPb
	path := "/api/2.0/permissions/authorization/tokens"
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
		&tokenPermissionsPb,
	)
	resp, err := TokenPermissionsFromPb(&tokenPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokenManagementImpl) UpdatePermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error) {
	requestPb, pbErr := TokenPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var tokenPermissionsPb settingspb.TokenPermissionsPb
	path := "/api/2.0/permissions/authorization/tokens"
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
		&tokenPermissionsPb,
	)
	resp, err := TokenPermissionsFromPb(&tokenPermissionsPb)
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
	requestPb, pbErr := CreateTokenRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createTokenResponsePb settingspb.CreateTokenResponsePb
	path := "/api/2.0/token/create"
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
		&createTokenResponsePb,
	)
	resp, err := CreateTokenResponseFromPb(&createTokenResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tokensImpl) Delete(ctx context.Context, request RevokeTokenRequest) error {
	requestPb, pbErr := RevokeTokenRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/token/delete"
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
		nil,
	)

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

	var listPublicTokensResponsePb settingspb.ListPublicTokensResponsePb
	path := "/api/2.0/token/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListPublicTokensResponseFromPb(&listPublicTokensResponsePb)
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
	requestPb, pbErr := GetStatusRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceConfPb map[string]string
	path := "/api/2.0/workspace-conf"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&workspaceConfPb,
	)
	resp := &workspaceConfPb

	return resp, err
}

func (a *workspaceConfImpl) SetStatus(ctx context.Context, request WorkspaceConf) error {
	requestPb, pbErr := WorkspaceConfToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/workspace-conf"
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

// unexported type that holds implementations of just WorkspaceNetworkConfiguration API methods
type workspaceNetworkConfigurationImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceNetworkConfigurationImpl) GetWorkspaceNetworkOptionRpc(ctx context.Context, request GetWorkspaceNetworkOptionRequest) (*WorkspaceNetworkOption, error) {
	requestPb, pbErr := GetWorkspaceNetworkOptionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceNetworkOptionPb settingspb.WorkspaceNetworkOptionPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/network", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&workspaceNetworkOptionPb,
	)
	resp, err := WorkspaceNetworkOptionFromPb(&workspaceNetworkOptionPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceNetworkConfigurationImpl) UpdateWorkspaceNetworkOptionRpc(ctx context.Context, request UpdateWorkspaceNetworkOptionRequest) (*WorkspaceNetworkOption, error) {
	requestPb, pbErr := UpdateWorkspaceNetworkOptionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceNetworkOptionPb settingspb.WorkspaceNetworkOptionPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/network", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb.WorkspaceNetworkOption,
		&workspaceNetworkOptionPb,
	)
	resp, err := WorkspaceNetworkOptionFromPb(&workspaceNetworkOptionPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
