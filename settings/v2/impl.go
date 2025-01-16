// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just AccountIpAccessLists API methods
type accountIpAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *accountIpAccessListsImpl) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {
	var createIpAccessListResponse CreateIpAccessListResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createIpAccessListResponse)
	return &createIpAccessListResponse, err
}

func (a *accountIpAccessListsImpl) Delete(ctx context.Context, request DeleteAccountIpAccessListRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *accountIpAccessListsImpl) Get(ctx context.Context, request GetAccountIpAccessListRequest) (*GetIpAccessListResponse, error) {
	var getIpAccessListResponse GetIpAccessListResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getIpAccessListResponse)
	return &getIpAccessListResponse, err
}

func (a *accountIpAccessListsImpl) List(ctx context.Context) (*GetIpAccessListsResponse, error) {
	var getIpAccessListsResponse GetIpAccessListsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &getIpAccessListsResponse)
	return &getIpAccessListsResponse, err
}

func (a *accountIpAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	var replaceResponse ReplaceResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &replaceResponse)
	return err
}

func (a *accountIpAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateResponse)
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
	var deleteAibiDashboardEmbeddingAccessPolicySettingResponse DeleteAibiDashboardEmbeddingAccessPolicySettingResponse
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteAibiDashboardEmbeddingAccessPolicySettingResponse)
	return &deleteAibiDashboardEmbeddingAccessPolicySettingResponse, err
}

func (a *aibiDashboardEmbeddingAccessPolicyImpl) Get(ctx context.Context, request GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error) {
	var aibiDashboardEmbeddingAccessPolicySetting AibiDashboardEmbeddingAccessPolicySetting
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &aibiDashboardEmbeddingAccessPolicySetting)
	return &aibiDashboardEmbeddingAccessPolicySetting, err
}

func (a *aibiDashboardEmbeddingAccessPolicyImpl) Update(ctx context.Context, request UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error) {
	var aibiDashboardEmbeddingAccessPolicySetting AibiDashboardEmbeddingAccessPolicySetting
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &aibiDashboardEmbeddingAccessPolicySetting)
	return &aibiDashboardEmbeddingAccessPolicySetting, err
}

// unexported type that holds implementations of just AibiDashboardEmbeddingApprovedDomains API methods
type aibiDashboardEmbeddingApprovedDomainsImpl struct {
	client *client.DatabricksClient
}

func (a *aibiDashboardEmbeddingApprovedDomainsImpl) Delete(ctx context.Context, request DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error) {
	var deleteAibiDashboardEmbeddingApprovedDomainsSettingResponse DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteAibiDashboardEmbeddingApprovedDomainsSettingResponse)
	return &deleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, err
}

func (a *aibiDashboardEmbeddingApprovedDomainsImpl) Get(ctx context.Context, request GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	var aibiDashboardEmbeddingApprovedDomainsSetting AibiDashboardEmbeddingApprovedDomainsSetting
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &aibiDashboardEmbeddingApprovedDomainsSetting)
	return &aibiDashboardEmbeddingApprovedDomainsSetting, err
}

func (a *aibiDashboardEmbeddingApprovedDomainsImpl) Update(ctx context.Context, request UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	var aibiDashboardEmbeddingApprovedDomainsSetting AibiDashboardEmbeddingApprovedDomainsSetting
	path := "/api/2.0/settings/types/aibi_dash_embed_ws_apprvd_domains/names/default"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &aibiDashboardEmbeddingApprovedDomainsSetting)
	return &aibiDashboardEmbeddingApprovedDomainsSetting, err
}

// unexported type that holds implementations of just AutomaticClusterUpdate API methods
type automaticClusterUpdateImpl struct {
	client *client.DatabricksClient
}

func (a *automaticClusterUpdateImpl) Get(ctx context.Context, request GetAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error) {
	var automaticClusterUpdateSetting AutomaticClusterUpdateSetting
	path := "/api/2.0/settings/types/automatic_cluster_update/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &automaticClusterUpdateSetting)
	return &automaticClusterUpdateSetting, err
}

func (a *automaticClusterUpdateImpl) Update(ctx context.Context, request UpdateAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error) {
	var automaticClusterUpdateSetting AutomaticClusterUpdateSetting
	path := "/api/2.0/settings/types/automatic_cluster_update/names/default"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &automaticClusterUpdateSetting)
	return &automaticClusterUpdateSetting, err
}

// unexported type that holds implementations of just ComplianceSecurityProfile API methods
type complianceSecurityProfileImpl struct {
	client *client.DatabricksClient
}

func (a *complianceSecurityProfileImpl) Get(ctx context.Context, request GetComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error) {
	var complianceSecurityProfileSetting ComplianceSecurityProfileSetting
	path := "/api/2.0/settings/types/shield_csp_enablement_ws_db/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &complianceSecurityProfileSetting)
	return &complianceSecurityProfileSetting, err
}

func (a *complianceSecurityProfileImpl) Update(ctx context.Context, request UpdateComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error) {
	var complianceSecurityProfileSetting ComplianceSecurityProfileSetting
	path := "/api/2.0/settings/types/shield_csp_enablement_ws_db/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &complianceSecurityProfileSetting)
	return &complianceSecurityProfileSetting, err
}

// unexported type that holds implementations of just CredentialsManager API methods
type credentialsManagerImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsManagerImpl) ExchangeToken(ctx context.Context, request ExchangeTokenRequest) (*ExchangeTokenResponse, error) {
	var exchangeTokenResponse ExchangeTokenResponse
	path := "/api/2.0/credentials-manager/exchange-tokens/token"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &exchangeTokenResponse)
	return &exchangeTokenResponse, err
}

// unexported type that holds implementations of just CspEnablementAccount API methods
type cspEnablementAccountImpl struct {
	client *client.DatabricksClient
}

func (a *cspEnablementAccountImpl) Get(ctx context.Context, request GetCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error) {
	var cspEnablementAccountSetting CspEnablementAccountSetting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_csp_enablement_ac/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &cspEnablementAccountSetting)
	return &cspEnablementAccountSetting, err
}

func (a *cspEnablementAccountImpl) Update(ctx context.Context, request UpdateCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error) {
	var cspEnablementAccountSetting CspEnablementAccountSetting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_csp_enablement_ac/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &cspEnablementAccountSetting)
	return &cspEnablementAccountSetting, err
}

// unexported type that holds implementations of just DefaultNamespace API methods
type defaultNamespaceImpl struct {
	client *client.DatabricksClient
}

func (a *defaultNamespaceImpl) Delete(ctx context.Context, request DeleteDefaultNamespaceSettingRequest) (*DeleteDefaultNamespaceSettingResponse, error) {
	var deleteDefaultNamespaceSettingResponse DeleteDefaultNamespaceSettingResponse
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteDefaultNamespaceSettingResponse)
	return &deleteDefaultNamespaceSettingResponse, err
}

func (a *defaultNamespaceImpl) Get(ctx context.Context, request GetDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error) {
	var defaultNamespaceSetting DefaultNamespaceSetting
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &defaultNamespaceSetting)
	return &defaultNamespaceSetting, err
}

func (a *defaultNamespaceImpl) Update(ctx context.Context, request UpdateDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error) {
	var defaultNamespaceSetting DefaultNamespaceSetting
	path := "/api/2.0/settings/types/default_namespace_ws/names/default"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &defaultNamespaceSetting)
	return &defaultNamespaceSetting, err
}

// unexported type that holds implementations of just DisableLegacyAccess API methods
type disableLegacyAccessImpl struct {
	client *client.DatabricksClient
}

func (a *disableLegacyAccessImpl) Delete(ctx context.Context, request DeleteDisableLegacyAccessRequest) (*DeleteDisableLegacyAccessResponse, error) {
	var deleteDisableLegacyAccessResponse DeleteDisableLegacyAccessResponse
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteDisableLegacyAccessResponse)
	return &deleteDisableLegacyAccessResponse, err
}

func (a *disableLegacyAccessImpl) Get(ctx context.Context, request GetDisableLegacyAccessRequest) (*DisableLegacyAccess, error) {
	var disableLegacyAccess DisableLegacyAccess
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &disableLegacyAccess)
	return &disableLegacyAccess, err
}

func (a *disableLegacyAccessImpl) Update(ctx context.Context, request UpdateDisableLegacyAccessRequest) (*DisableLegacyAccess, error) {
	var disableLegacyAccess DisableLegacyAccess
	path := "/api/2.0/settings/types/disable_legacy_access/names/default"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &disableLegacyAccess)
	return &disableLegacyAccess, err
}

// unexported type that holds implementations of just DisableLegacyDbfs API methods
type disableLegacyDbfsImpl struct {
	client *client.DatabricksClient
}

func (a *disableLegacyDbfsImpl) Delete(ctx context.Context, request DeleteDisableLegacyDbfsRequest) (*DeleteDisableLegacyDbfsResponse, error) {
	var deleteDisableLegacyDbfsResponse DeleteDisableLegacyDbfsResponse
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteDisableLegacyDbfsResponse)
	return &deleteDisableLegacyDbfsResponse, err
}

func (a *disableLegacyDbfsImpl) Get(ctx context.Context, request GetDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error) {
	var disableLegacyDbfs DisableLegacyDbfs
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &disableLegacyDbfs)
	return &disableLegacyDbfs, err
}

func (a *disableLegacyDbfsImpl) Update(ctx context.Context, request UpdateDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error) {
	var disableLegacyDbfs DisableLegacyDbfs
	path := "/api/2.0/settings/types/disable_legacy_dbfs/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &disableLegacyDbfs)
	return &disableLegacyDbfs, err
}

// unexported type that holds implementations of just DisableLegacyFeatures API methods
type disableLegacyFeaturesImpl struct {
	client *client.DatabricksClient
}

func (a *disableLegacyFeaturesImpl) Delete(ctx context.Context, request DeleteDisableLegacyFeaturesRequest) (*DeleteDisableLegacyFeaturesResponse, error) {
	var deleteDisableLegacyFeaturesResponse DeleteDisableLegacyFeaturesResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteDisableLegacyFeaturesResponse)
	return &deleteDisableLegacyFeaturesResponse, err
}

func (a *disableLegacyFeaturesImpl) Get(ctx context.Context, request GetDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error) {
	var disableLegacyFeatures DisableLegacyFeatures
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &disableLegacyFeatures)
	return &disableLegacyFeatures, err
}

func (a *disableLegacyFeaturesImpl) Update(ctx context.Context, request UpdateDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error) {
	var disableLegacyFeatures DisableLegacyFeatures
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/disable_legacy_features/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &disableLegacyFeatures)
	return &disableLegacyFeatures, err
}

// unexported type that holds implementations of just EnhancedSecurityMonitoring API methods
type enhancedSecurityMonitoringImpl struct {
	client *client.DatabricksClient
}

func (a *enhancedSecurityMonitoringImpl) Get(ctx context.Context, request GetEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error) {
	var enhancedSecurityMonitoringSetting EnhancedSecurityMonitoringSetting
	path := "/api/2.0/settings/types/shield_esm_enablement_ws_db/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &enhancedSecurityMonitoringSetting)
	return &enhancedSecurityMonitoringSetting, err
}

func (a *enhancedSecurityMonitoringImpl) Update(ctx context.Context, request UpdateEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error) {
	var enhancedSecurityMonitoringSetting EnhancedSecurityMonitoringSetting
	path := "/api/2.0/settings/types/shield_esm_enablement_ws_db/names/default"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &enhancedSecurityMonitoringSetting)
	return &enhancedSecurityMonitoringSetting, err
}

// unexported type that holds implementations of just EsmEnablementAccount API methods
type esmEnablementAccountImpl struct {
	client *client.DatabricksClient
}

func (a *esmEnablementAccountImpl) Get(ctx context.Context, request GetEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error) {
	var esmEnablementAccountSetting EsmEnablementAccountSetting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_esm_enablement_ac/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &esmEnablementAccountSetting)
	return &esmEnablementAccountSetting, err
}

func (a *esmEnablementAccountImpl) Update(ctx context.Context, request UpdateEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error) {
	var esmEnablementAccountSetting EsmEnablementAccountSetting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/shield_esm_enablement_ac/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &esmEnablementAccountSetting)
	return &esmEnablementAccountSetting, err
}

// unexported type that holds implementations of just IpAccessLists API methods
type ipAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *ipAccessListsImpl) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {
	var createIpAccessListResponse CreateIpAccessListResponse
	path := "/api/2.0/ip-access-lists"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createIpAccessListResponse)
	return &createIpAccessListResponse, err
}

func (a *ipAccessListsImpl) Delete(ctx context.Context, request DeleteIpAccessListRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *ipAccessListsImpl) Get(ctx context.Context, request GetIpAccessListRequest) (*FetchIpAccessListResponse, error) {
	var fetchIpAccessListResponse FetchIpAccessListResponse
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &fetchIpAccessListResponse)
	return &fetchIpAccessListResponse, err
}

func (a *ipAccessListsImpl) List(ctx context.Context) (*ListIpAccessListResponse, error) {
	var listIpAccessListResponse ListIpAccessListResponse
	path := "/api/2.0/ip-access-lists"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listIpAccessListResponse)
	return &listIpAccessListResponse, err
}

func (a *ipAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	var replaceResponse ReplaceResponse
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &replaceResponse)
	return err
}

func (a *ipAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just NetworkConnectivity API methods
type networkConnectivityImpl struct {
	client *client.DatabricksClient
}

func (a *networkConnectivityImpl) CreateNetworkConnectivityConfiguration(ctx context.Context, request CreateNetworkConnectivityConfigRequest) (*NetworkConnectivityConfiguration, error) {
	var networkConnectivityConfiguration NetworkConnectivityConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &networkConnectivityConfiguration)
	return &networkConnectivityConfiguration, err
}

func (a *networkConnectivityImpl) CreatePrivateEndpointRule(ctx context.Context, request CreatePrivateEndpointRuleRequest) (*NccAzurePrivateEndpointRule, error) {
	var nccAzurePrivateEndpointRule NccAzurePrivateEndpointRule
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &nccAzurePrivateEndpointRule)
	return &nccAzurePrivateEndpointRule, err
}

func (a *networkConnectivityImpl) DeleteNetworkConnectivityConfiguration(ctx context.Context, request DeleteNetworkConnectivityConfigurationRequest) error {
	var deleteNetworkConnectivityConfigurationResponse DeleteNetworkConnectivityConfigurationResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteNetworkConnectivityConfigurationResponse)
	return err
}

func (a *networkConnectivityImpl) DeletePrivateEndpointRule(ctx context.Context, request DeletePrivateEndpointRuleRequest) (*NccAzurePrivateEndpointRule, error) {
	var nccAzurePrivateEndpointRule NccAzurePrivateEndpointRule
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId, request.PrivateEndpointRuleId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &nccAzurePrivateEndpointRule)
	return &nccAzurePrivateEndpointRule, err
}

func (a *networkConnectivityImpl) GetNetworkConnectivityConfiguration(ctx context.Context, request GetNetworkConnectivityConfigurationRequest) (*NetworkConnectivityConfiguration, error) {
	var networkConnectivityConfiguration NetworkConnectivityConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &networkConnectivityConfiguration)
	return &networkConnectivityConfiguration, err
}

func (a *networkConnectivityImpl) GetPrivateEndpointRule(ctx context.Context, request GetPrivateEndpointRuleRequest) (*NccAzurePrivateEndpointRule, error) {
	var nccAzurePrivateEndpointRule NccAzurePrivateEndpointRule
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules/%v", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId, request.PrivateEndpointRuleId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &nccAzurePrivateEndpointRule)
	return &nccAzurePrivateEndpointRule, err
}

func (a *networkConnectivityImpl) ListNetworkConnectivityConfigurations(ctx context.Context, request ListNetworkConnectivityConfigurationsRequest) (*ListNetworkConnectivityConfigurationsResponse, error) {
	var listNetworkConnectivityConfigurationsResponse ListNetworkConnectivityConfigurationsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listNetworkConnectivityConfigurationsResponse)
	return &listNetworkConnectivityConfigurationsResponse, err
}

func (a *networkConnectivityImpl) ListPrivateEndpointRules(ctx context.Context, request ListPrivateEndpointRulesRequest) (*ListNccAzurePrivateEndpointRulesResponse, error) {
	var listNccAzurePrivateEndpointRulesResponse ListNccAzurePrivateEndpointRulesResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/network-connectivity-configs/%v/private-endpoint-rules", a.client.ConfiguredAccountID(), request.NetworkConnectivityConfigId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listNccAzurePrivateEndpointRulesResponse)
	return &listNccAzurePrivateEndpointRulesResponse, err
}

// unexported type that holds implementations of just NotificationDestinations API methods
type notificationDestinationsImpl struct {
	client *client.DatabricksClient
}

func (a *notificationDestinationsImpl) Create(ctx context.Context, request CreateNotificationDestinationRequest) (*NotificationDestination, error) {
	var notificationDestination NotificationDestination
	path := "/api/2.0/notification-destinations"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &notificationDestination)
	return &notificationDestination, err
}

func (a *notificationDestinationsImpl) Delete(ctx context.Context, request DeleteNotificationDestinationRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &empty)
	return err
}

func (a *notificationDestinationsImpl) Get(ctx context.Context, request GetNotificationDestinationRequest) (*NotificationDestination, error) {
	var notificationDestination NotificationDestination
	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &notificationDestination)
	return &notificationDestination, err
}

func (a *notificationDestinationsImpl) List(ctx context.Context, request ListNotificationDestinationsRequest) (*ListNotificationDestinationsResponse, error) {
	var listNotificationDestinationsResponse ListNotificationDestinationsResponse
	path := "/api/2.0/notification-destinations"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listNotificationDestinationsResponse)
	return &listNotificationDestinationsResponse, err
}

func (a *notificationDestinationsImpl) Update(ctx context.Context, request UpdateNotificationDestinationRequest) (*NotificationDestination, error) {
	var notificationDestination NotificationDestination
	path := fmt.Sprintf("/api/2.0/notification-destinations/%v", request.Id)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &notificationDestination)
	return &notificationDestination, err
}

// unexported type that holds implementations of just PersonalCompute API methods
type personalComputeImpl struct {
	client *client.DatabricksClient
}

func (a *personalComputeImpl) Delete(ctx context.Context, request DeletePersonalComputeSettingRequest) (*DeletePersonalComputeSettingResponse, error) {
	var deletePersonalComputeSettingResponse DeletePersonalComputeSettingResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deletePersonalComputeSettingResponse)
	return &deletePersonalComputeSettingResponse, err
}

func (a *personalComputeImpl) Get(ctx context.Context, request GetPersonalComputeSettingRequest) (*PersonalComputeSetting, error) {
	var personalComputeSetting PersonalComputeSetting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &personalComputeSetting)
	return &personalComputeSetting, err
}

func (a *personalComputeImpl) Update(ctx context.Context, request UpdatePersonalComputeSettingRequest) (*PersonalComputeSetting, error) {
	var personalComputeSetting PersonalComputeSetting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &personalComputeSetting)
	return &personalComputeSetting, err
}

// unexported type that holds implementations of just RestrictWorkspaceAdmins API methods
type restrictWorkspaceAdminsImpl struct {
	client *client.DatabricksClient
}

func (a *restrictWorkspaceAdminsImpl) Delete(ctx context.Context, request DeleteRestrictWorkspaceAdminsSettingRequest) (*DeleteRestrictWorkspaceAdminsSettingResponse, error) {
	var deleteRestrictWorkspaceAdminsSettingResponse DeleteRestrictWorkspaceAdminsSettingResponse
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteRestrictWorkspaceAdminsSettingResponse)
	return &deleteRestrictWorkspaceAdminsSettingResponse, err
}

func (a *restrictWorkspaceAdminsImpl) Get(ctx context.Context, request GetRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error) {
	var restrictWorkspaceAdminsSetting RestrictWorkspaceAdminsSetting
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &restrictWorkspaceAdminsSetting)
	return &restrictWorkspaceAdminsSetting, err
}

func (a *restrictWorkspaceAdminsImpl) Update(ctx context.Context, request UpdateRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error) {
	var restrictWorkspaceAdminsSetting RestrictWorkspaceAdminsSetting
	path := "/api/2.0/settings/types/restrict_workspace_admins/names/default"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &restrictWorkspaceAdminsSetting)
	return &restrictWorkspaceAdminsSetting, err
}

// unexported type that holds implementations of just Settings API methods
type settingsImpl struct {
	client *client.DatabricksClient
}

// unexported type that holds implementations of just TokenManagement API methods
type tokenManagementImpl struct {
	client *client.DatabricksClient
}

func (a *tokenManagementImpl) CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error) {
	var createOboTokenResponse CreateOboTokenResponse
	path := "/api/2.0/token-management/on-behalf-of/tokens"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createOboTokenResponse)
	return &createOboTokenResponse, err
}

func (a *tokenManagementImpl) Delete(ctx context.Context, request DeleteTokenManagementRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *tokenManagementImpl) Get(ctx context.Context, request GetTokenManagementRequest) (*GetTokenResponse, error) {
	var getTokenResponse GetTokenResponse
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getTokenResponse)
	return &getTokenResponse, err
}

func (a *tokenManagementImpl) GetPermissionLevels(ctx context.Context) (*GetTokenPermissionLevelsResponse, error) {
	var getTokenPermissionLevelsResponse GetTokenPermissionLevelsResponse
	path := "/api/2.0/permissions/authorization/tokens/permissionLevels"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &getTokenPermissionLevelsResponse)
	return &getTokenPermissionLevelsResponse, err
}

func (a *tokenManagementImpl) GetPermissions(ctx context.Context) (*TokenPermissions, error) {
	var tokenPermissions TokenPermissions
	path := "/api/2.0/permissions/authorization/tokens"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &tokenPermissions)
	return &tokenPermissions, err
}

func (a *tokenManagementImpl) List(ctx context.Context, request ListTokenManagementRequest) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token-management/tokens"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listTokensResponse)
	return &listTokensResponse, err
}

func (a *tokenManagementImpl) SetPermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error) {
	var tokenPermissions TokenPermissions
	path := "/api/2.0/permissions/authorization/tokens"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &tokenPermissions)
	return &tokenPermissions, err
}

func (a *tokenManagementImpl) UpdatePermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error) {
	var tokenPermissions TokenPermissions
	path := "/api/2.0/permissions/authorization/tokens"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &tokenPermissions)
	return &tokenPermissions, err
}

// unexported type that holds implementations of just Tokens API methods
type tokensImpl struct {
	client *client.DatabricksClient
}

func (a *tokensImpl) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {
	var createTokenResponse CreateTokenResponse
	path := "/api/2.0/token/create"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createTokenResponse)
	return &createTokenResponse, err
}

func (a *tokensImpl) Delete(ctx context.Context, request RevokeTokenRequest) error {
	var revokeTokenResponse RevokeTokenResponse
	path := "/api/2.0/token/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &revokeTokenResponse)
	return err
}

func (a *tokensImpl) List(ctx context.Context) (*ListPublicTokensResponse, error) {
	var listPublicTokensResponse ListPublicTokensResponse
	path := "/api/2.0/token/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listPublicTokensResponse)
	return &listPublicTokensResponse, err
}

// unexported type that holds implementations of just WorkspaceConf API methods
type workspaceConfImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceConfImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*map[string]string, error) {
	var workspaceConf map[string]string
	path := "/api/2.0/workspace-conf"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &workspaceConf)
	return &workspaceConf, err
}

func (a *workspaceConfImpl) SetStatus(ctx context.Context, request WorkspaceConf) error {
	var setStatusResponse SetStatusResponse
	path := "/api/2.0/workspace-conf"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &setStatusResponse)
	return err
}
