// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"context"
)

// The Accounts IP Access List API enables account admins to configure IP access
// lists for access to the account console.
//
// Account IP Access Lists affect web application access and REST API access to
// the account console and account APIs. If the feature is disabled for the
// account, all access is allowed for this account. There is support for allow
// lists (inclusion) and block lists (exclusion).
//
// When a connection is attempted: 1. **First, all block lists are checked.** If
// the connection IP address matches any block list, the connection is rejected.
// 2. **If the connection was not rejected by block lists**, the IP address is
// compared with the allow lists.
//
// If there is at least one allow list for the account, the connection is
// allowed only if the IP address matches an allow list. If there are no allow
// lists for the account, all IP addresses are allowed.
//
// For all allow lists and block lists combined, the account supports a maximum
// of 1000 IP/CIDR values, where one CIDR counts as a single value.
//
// After changes to the account-level IP access lists, it can take a few minutes
// for changes to take effect.
type AccountIpAccessListsService interface {

	// Create access list.
	//
	// Creates an IP access list for the account.
	//
	// A list can be an allow list or a block list. See the top of this file for
	// a description of how the server treats allow lists and block lists at
	// runtime.
	//
	// When creating or updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the new list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect.
	Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error)

	// Delete access list.
	//
	// Deletes an IP access list, specified by its list ID.
	Delete(ctx context.Context, request DeleteAccountIpAccessListRequest) error

	// Get IP access list.
	//
	// Gets an IP access list, specified by its list ID.
	Get(ctx context.Context, request GetAccountIpAccessListRequest) (*GetIpAccessListResponse, error)

	// Get access lists.
	//
	// Gets all IP access lists for the specified account.
	//
	// Use ListAll() to get all IpAccessListInfo instances
	List(ctx context.Context) (*GetIpAccessListsResponse, error)

	// Replace access list.
	//
	// Replaces an IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time. When replacing an IP access list: * For all allow lists and
	// block lists combined, the API supports a maximum of 1000 IP/CIDR values,
	// where one CIDR counts as a single value. Attempts to exceed that number
	// return error 400 with `error_code` value `QUOTA_EXCEEDED`. * If the
	// resulting list would block the calling user's current IP, error 400 is
	// returned with `error_code` value `INVALID_STATE`. It can take a few
	// minutes for the changes to take effect.
	Replace(ctx context.Context, request ReplaceIpAccessList) error

	// Update access list.
	//
	// Updates an existing IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time.
	//
	// When updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the updated list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect.
	Update(ctx context.Context, request UpdateIpAccessList) error
}

// Accounts Settings API allows users to manage settings at the account level.
type AccountSettingsService interface {
}

// Controls whether AI/BI published dashboard embedding is enabled,
// conditionally enabled, or disabled at the workspace level. By default, this
// setting is conditionally enabled (ALLOW_APPROVED_DOMAINS).
type AibiDashboardEmbeddingAccessPolicyService interface {

	// Delete the AI/BI dashboard embedding access policy.
	//
	// Delete the AI/BI dashboard embedding access policy, reverting back to the
	// default.
	Delete(ctx context.Context, request DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error)

	// Retrieve the AI/BI dashboard embedding access policy.
	//
	// Retrieves the AI/BI dashboard embedding access policy. The default
	// setting is ALLOW_APPROVED_DOMAINS, permitting AI/BI dashboards to be
	// embedded on approved domains.
	Get(ctx context.Context, request GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error)

	// Update the AI/BI dashboard embedding access policy.
	//
	// Updates the AI/BI dashboard embedding access policy at the workspace
	// level.
	Update(ctx context.Context, request UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*AibiDashboardEmbeddingAccessPolicySetting, error)
}

// Controls the list of domains approved to host the embedded AI/BI dashboards.
// The approved domains list can't be mutated when the current access policy is
// not set to ALLOW_APPROVED_DOMAINS.
type AibiDashboardEmbeddingApprovedDomainsService interface {

	// Delete AI/BI dashboard embedding approved domains.
	//
	// Delete the list of domains approved to host embedded AI/BI dashboards,
	// reverting back to the default empty list.
	Delete(ctx context.Context, request DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error)

	// Retrieve the list of domains approved to host embedded AI/BI dashboards.
	//
	// Retrieves the list of domains approved to host embedded AI/BI dashboards.
	Get(ctx context.Context, request GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error)

	// Update the list of domains approved to host embedded AI/BI dashboards.
	//
	// Updates the list of domains approved to host embedded AI/BI dashboards.
	// This update will fail if the current workspace access policy is not
	// ALLOW_APPROVED_DOMAINS.
	Update(ctx context.Context, request UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*AibiDashboardEmbeddingApprovedDomainsSetting, error)
}

// Controls whether automatic cluster update is enabled for the current
// workspace. By default, it is turned off.
type AutomaticClusterUpdateService interface {

	// Get the automatic cluster update setting.
	//
	// Gets the automatic cluster update setting.
	Get(ctx context.Context, request GetAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error)

	// Update the automatic cluster update setting.
	//
	// Updates the automatic cluster update setting for the workspace. A fresh
	// etag needs to be provided in `PATCH` requests (as part of the setting
	// field). The etag can be retrieved by making a `GET` request before the
	// `PATCH` request. If the setting is updated concurrently, `PATCH` fails
	// with 409 and the request must be retried by using the fresh etag in the
	// 409 response.
	Update(ctx context.Context, request UpdateAutomaticClusterUpdateSettingRequest) (*AutomaticClusterUpdateSetting, error)
}

// Controls whether to enable the compliance security profile for the current
// workspace. Enabling it on a workspace is permanent. By default, it is turned
// off.
//
// This settings can NOT be disabled once it is enabled.
type ComplianceSecurityProfileService interface {

	// Get the compliance security profile setting.
	//
	// Gets the compliance security profile setting.
	Get(ctx context.Context, request GetComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error)

	// Update the compliance security profile setting.
	//
	// Updates the compliance security profile setting for the workspace. A
	// fresh etag needs to be provided in `PATCH` requests (as part of the
	// setting field). The etag can be retrieved by making a `GET` request
	// before the `PATCH` request. If the setting is updated concurrently,
	// `PATCH` fails with 409 and the request must be retried by using the fresh
	// etag in the 409 response.
	Update(ctx context.Context, request UpdateComplianceSecurityProfileSettingRequest) (*ComplianceSecurityProfileSetting, error)
}

// Credentials manager interacts with with Identity Providers to to perform
// token exchanges using stored credentials and refresh tokens.
type CredentialsManagerService interface {

	// Exchange token.
	//
	// Exchange tokens with an Identity Provider to get a new access token. It
	// allows specifying scopes to determine token permissions.
	ExchangeToken(ctx context.Context, request ExchangeTokenRequest) (*ExchangeTokenResponse, error)
}

// The compliance security profile settings at the account level control whether
// to enable it for new workspaces. By default, this account-level setting is
// disabled for new workspaces. After workspace creation, account admins can
// enable the compliance security profile individually for each workspace.
//
// This settings can be disabled so that new workspaces do not have compliance
// security profile enabled by default.
type CspEnablementAccountService interface {

	// Get the compliance security profile setting for new workspaces.
	//
	// Gets the compliance security profile setting for new workspaces.
	Get(ctx context.Context, request GetCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error)

	// Update the compliance security profile setting for new workspaces.
	//
	// Updates the value of the compliance security profile setting for new
	// workspaces.
	Update(ctx context.Context, request UpdateCspEnablementAccountSettingRequest) (*CspEnablementAccountSetting, error)
}

// The default namespace setting API allows users to configure the default
// namespace for a Databricks workspace.
//
// Through this API, users can retrieve, set, or modify the default namespace
// used when queries do not reference a fully qualified three-level name. For
// example, if you use the API to set 'retail_prod' as the default catalog, then
// a query 'SELECT * FROM myTable' would reference the object
// 'retail_prod.default.myTable' (the schema 'default' is always assumed).
//
// This setting requires a restart of clusters and SQL warehouses to take
// effect. Additionally, the default namespace only applies when using Unity
// Catalog-enabled compute.
type DefaultNamespaceService interface {

	// Delete the default namespace setting.
	//
	// Deletes the default namespace setting for the workspace. A fresh etag
	// needs to be provided in `DELETE` requests (as a query parameter). The
	// etag can be retrieved by making a `GET` request before the `DELETE`
	// request. If the setting is updated/deleted concurrently, `DELETE` fails
	// with 409 and the request must be retried by using the fresh etag in the
	// 409 response.
	Delete(ctx context.Context, request DeleteDefaultNamespaceSettingRequest) (*DeleteDefaultNamespaceSettingResponse, error)

	// Get the default namespace setting.
	//
	// Gets the default namespace setting.
	Get(ctx context.Context, request GetDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error)

	// Update the default namespace setting.
	//
	// Updates the default namespace setting for the workspace. A fresh etag
	// needs to be provided in `PATCH` requests (as part of the setting field).
	// The etag can be retrieved by making a `GET` request before the `PATCH`
	// request. Note that if the setting does not exist, `GET` returns a
	// NOT_FOUND error and the etag is present in the error response, which
	// should be set in the `PATCH` request. If the setting is updated
	// concurrently, `PATCH` fails with 409 and the request must be retried by
	// using the fresh etag in the 409 response.
	Update(ctx context.Context, request UpdateDefaultNamespaceSettingRequest) (*DefaultNamespaceSetting, error)
}

// 'Disabling legacy access' has the following impacts:
//
// 1. Disables direct access to Hive Metastores from the workspace. However, you
// can still access a Hive Metastore through Hive Metastore federation. 2.
// Disables fallback mode on external location access from the workspace. 3.
// Disables Databricks Runtime versions prior to 13.3LTS.
type DisableLegacyAccessService interface {

	// Delete Legacy Access Disablement Status.
	//
	// Deletes legacy access disablement status.
	Delete(ctx context.Context, request DeleteDisableLegacyAccessRequest) (*DeleteDisableLegacyAccessResponse, error)

	// Retrieve Legacy Access Disablement Status.
	//
	// Retrieves legacy access disablement Status.
	Get(ctx context.Context, request GetDisableLegacyAccessRequest) (*DisableLegacyAccess, error)

	// Update Legacy Access Disablement Status.
	//
	// Updates legacy access disablement status.
	Update(ctx context.Context, request UpdateDisableLegacyAccessRequest) (*DisableLegacyAccess, error)
}

// Disabling legacy DBFS has the following implications:
//
// 1. Access to DBFS root and DBFS mounts is disallowed (as well as the creation
// of new mounts). 2. Disables Databricks Runtime versions prior to 13.3LTS.
//
// When the setting is off, all DBFS functionality is enabled and no
// restrictions are imposed on Databricks Runtime versions. This setting can
// take up to 20 minutes to take effect and requires a manual restart of
// all-purpose compute clusters and SQL warehouses.
type DisableLegacyDbfsService interface {

	// Delete the disable legacy DBFS setting.
	//
	// Deletes the disable legacy DBFS setting for a workspace, reverting back
	// to the default.
	Delete(ctx context.Context, request DeleteDisableLegacyDbfsRequest) (*DeleteDisableLegacyDbfsResponse, error)

	// Get the disable legacy DBFS setting.
	//
	// Gets the disable legacy DBFS setting.
	Get(ctx context.Context, request GetDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error)

	// Update the disable legacy DBFS setting.
	//
	// Updates the disable legacy DBFS setting for the workspace.
	Update(ctx context.Context, request UpdateDisableLegacyDbfsRequest) (*DisableLegacyDbfs, error)
}

// Disable legacy features for new Databricks workspaces.
//
// For newly created workspaces: 1. Disables the use of DBFS root and mounts. 2.
// Hive Metastore will not be provisioned. 3. Disables the use of
// ‘No-isolation clusters’. 4. Disables Databricks Runtime versions prior to
// 13.3LTS.
type DisableLegacyFeaturesService interface {

	// Delete the disable legacy features setting.
	//
	// Deletes the disable legacy features setting.
	Delete(ctx context.Context, request DeleteDisableLegacyFeaturesRequest) (*DeleteDisableLegacyFeaturesResponse, error)

	// Get the disable legacy features setting.
	//
	// Gets the value of the disable legacy features setting.
	Get(ctx context.Context, request GetDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error)

	// Update the disable legacy features setting.
	//
	// Updates the value of the disable legacy features setting.
	Update(ctx context.Context, request UpdateDisableLegacyFeaturesRequest) (*DisableLegacyFeatures, error)
}

// Controls whether users can export notebooks and files from the Workspace UI.
// By default, this setting is enabled.
type EnableExportNotebookService interface {

	// Get the Notebook and File exporting setting.
	//
	// Gets the Notebook and File exporting setting.
	GetEnableExportNotebook(ctx context.Context) (*EnableExportNotebook, error)

	// Update the Notebook and File exporting setting.
	//
	// Updates the Notebook and File exporting setting. The model follows
	// eventual consistency, which means the get after the update operation
	// might receive stale values for some time.
	PatchEnableExportNotebook(ctx context.Context, request UpdateEnableExportNotebookRequest) (*EnableExportNotebook, error)
}

// Controls the enforcement of IP access lists for accessing the account
// console. Allowing you to enable or disable restricted access based on IP
// addresses.
type EnableIpAccessListsService interface {

	// Delete the account IP access toggle setting.
	//
	// Reverts the value of the account IP access toggle setting to default (ON)
	Delete(ctx context.Context, request DeleteAccountIpAccessEnableRequest) (*DeleteAccountIpAccessEnableResponse, error)

	// Get the account IP access toggle setting.
	//
	// Gets the value of the account IP access toggle setting.
	Get(ctx context.Context, request GetAccountIpAccessEnableRequest) (*AccountIpAccessEnable, error)

	// Update the account IP access toggle setting.
	//
	// Updates the value of the account IP access toggle setting.
	Update(ctx context.Context, request UpdateAccountIpAccessEnableRequest) (*AccountIpAccessEnable, error)
}

// Controls whether users can copy tabular data to the clipboard via the UI. By
// default, this setting is enabled.
type EnableNotebookTableClipboardService interface {

	// Get the Results Table Clipboard features setting.
	//
	// Gets the Results Table Clipboard features setting.
	GetEnableNotebookTableClipboard(ctx context.Context) (*EnableNotebookTableClipboard, error)

	// Update the Results Table Clipboard features setting.
	//
	// Updates the Results Table Clipboard features setting. The model follows
	// eventual consistency, which means the get after the update operation
	// might receive stale values for some time.
	PatchEnableNotebookTableClipboard(ctx context.Context, request UpdateEnableNotebookTableClipboardRequest) (*EnableNotebookTableClipboard, error)
}

// Controls whether users can download notebook results. By default, this
// setting is enabled.
type EnableResultsDownloadingService interface {

	// Get the Notebook results download setting.
	//
	// Gets the Notebook results download setting.
	GetEnableResultsDownloading(ctx context.Context) (*EnableResultsDownloading, error)

	// Update the Notebook results download setting.
	//
	// Updates the Notebook results download setting. The model follows eventual
	// consistency, which means the get after the update operation might receive
	// stale values for some time.
	PatchEnableResultsDownloading(ctx context.Context, request UpdateEnableResultsDownloadingRequest) (*EnableResultsDownloading, error)
}

// Controls whether enhanced security monitoring is enabled for the current
// workspace. If the compliance security profile is enabled, this is
// automatically enabled. By default, it is disabled. However, if the compliance
// security profile is enabled, this is automatically enabled.
//
// If the compliance security profile is disabled, you can enable or disable
// this setting and it is not permanent.
type EnhancedSecurityMonitoringService interface {

	// Get the enhanced security monitoring setting.
	//
	// Gets the enhanced security monitoring setting.
	Get(ctx context.Context, request GetEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error)

	// Update the enhanced security monitoring setting.
	//
	// Updates the enhanced security monitoring setting for the workspace. A
	// fresh etag needs to be provided in `PATCH` requests (as part of the
	// setting field). The etag can be retrieved by making a `GET` request
	// before the `PATCH` request. If the setting is updated concurrently,
	// `PATCH` fails with 409 and the request must be retried by using the fresh
	// etag in the 409 response.
	Update(ctx context.Context, request UpdateEnhancedSecurityMonitoringSettingRequest) (*EnhancedSecurityMonitoringSetting, error)
}

// The enhanced security monitoring setting at the account level controls
// whether to enable the feature on new workspaces. By default, this
// account-level setting is disabled for new workspaces. After workspace
// creation, account admins can enable enhanced security monitoring individually
// for each workspace.
type EsmEnablementAccountService interface {

	// Get the enhanced security monitoring setting for new workspaces.
	//
	// Gets the enhanced security monitoring setting for new workspaces.
	Get(ctx context.Context, request GetEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error)

	// Update the enhanced security monitoring setting for new workspaces.
	//
	// Updates the value of the enhanced security monitoring setting for new
	// workspaces.
	Update(ctx context.Context, request UpdateEsmEnablementAccountSettingRequest) (*EsmEnablementAccountSetting, error)
}

// IP Access List enables admins to configure IP access lists.
//
// IP access lists affect web application access and REST API access to this
// workspace only. If the feature is disabled for a workspace, all access is
// allowed for this workspace. There is support for allow lists (inclusion) and
// block lists (exclusion).
//
// When a connection is attempted: 1. **First, all block lists are checked.** If
// the connection IP address matches any block list, the connection is rejected.
// 2. **If the connection was not rejected by block lists**, the IP address is
// compared with the allow lists.
//
// If there is at least one allow list for the workspace, the connection is
// allowed only if the IP address matches an allow list. If there are no allow
// lists for the workspace, all IP addresses are allowed.
//
// For all allow lists and block lists combined, the workspace supports a
// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
//
// After changes to the IP access list feature, it can take a few minutes for
// changes to take effect.
type IpAccessListsService interface {

	// Create access list.
	//
	// Creates an IP access list for this workspace.
	//
	// A list can be an allow list or a block list. See the top of this file for
	// a description of how the server treats allow lists and block lists at
	// runtime.
	//
	// When creating or updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the new list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect. **Note**: Your
	// new IP access list has no effect until you enable the feature. See
	// :method:workspaceconf/setStatus
	Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error)

	// Delete access list.
	//
	// Deletes an IP access list, specified by its list ID.
	Delete(ctx context.Context, request DeleteIpAccessListRequest) error

	// Get access list.
	//
	// Gets an IP access list, specified by its list ID.
	Get(ctx context.Context, request GetIpAccessListRequest) (*FetchIpAccessListResponse, error)

	// Get access lists.
	//
	// Gets all IP access lists for the specified workspace.
	//
	// Use ListAll() to get all IpAccessListInfo instances
	List(ctx context.Context) (*ListIpAccessListResponse, error)

	// Replace access list.
	//
	// Replaces an IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time. When replacing an IP access list: * For all allow lists and
	// block lists combined, the API supports a maximum of 1000 IP/CIDR values,
	// where one CIDR counts as a single value. Attempts to exceed that number
	// return error 400 with `error_code` value `QUOTA_EXCEEDED`. * If the
	// resulting list would block the calling user's current IP, error 400 is
	// returned with `error_code` value `INVALID_STATE`. It can take a few
	// minutes for the changes to take effect. Note that your resulting IP
	// access list has no effect until you enable the feature. See
	// :method:workspaceconf/setStatus.
	Replace(ctx context.Context, request ReplaceIpAccessList) error

	// Update access list.
	//
	// Updates an existing IP access list, specified by its ID.
	//
	// A list can include allow lists and block lists. See the top of this file
	// for a description of how the server treats allow lists and block lists at
	// run time.
	//
	// When updating an IP access list:
	//
	// * For all allow lists and block lists combined, the API supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	// Attempts to exceed that number return error 400 with `error_code` value
	// `QUOTA_EXCEEDED`. * If the updated list would block the calling user's
	// current IP, error 400 is returned with `error_code` value
	// `INVALID_STATE`.
	//
	// It can take a few minutes for the changes to take effect. Note that your
	// resulting IP access list has no effect until you enable the feature. See
	// :method:workspaceconf/setStatus.
	Update(ctx context.Context, request UpdateIpAccessList) error
}

// Determines if partner powered models are enabled or not for a specific
// account
type LlmProxyPartnerPoweredAccountService interface {

	// Get the enable partner powered AI features account setting.
	//
	// Gets the enable partner powered AI features account setting.
	Get(ctx context.Context, request GetLlmProxyPartnerPoweredAccountRequest) (*LlmProxyPartnerPoweredAccount, error)

	// Update the enable partner powered AI features account setting.
	//
	// Updates the enable partner powered AI features account setting.
	Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredAccountRequest) (*LlmProxyPartnerPoweredAccount, error)
}

// Determines if the account-level partner-powered setting value is enforced
// upon the workspace-level partner-powered setting
type LlmProxyPartnerPoweredEnforceService interface {

	// Get the enforcement status of partner powered AI features account
	// setting.
	//
	// Gets the enforcement status of partner powered AI features account
	// setting.
	Get(ctx context.Context, request GetLlmProxyPartnerPoweredEnforceRequest) (*LlmProxyPartnerPoweredEnforce, error)

	// Update the enforcement status of partner powered AI features account
	// setting.
	//
	// Updates the enable enforcement status of partner powered AI features
	// account setting.
	Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredEnforceRequest) (*LlmProxyPartnerPoweredEnforce, error)
}

// Determines if partner powered models are enabled or not for a specific
// workspace
type LlmProxyPartnerPoweredWorkspaceService interface {

	// Delete the enable partner powered AI features workspace setting.
	//
	// Reverts the enable partner powered AI features workspace setting to its
	// default value.
	Delete(ctx context.Context, request DeleteLlmProxyPartnerPoweredWorkspaceRequest) (*DeleteLlmProxyPartnerPoweredWorkspaceResponse, error)

	// Get the enable partner powered AI features workspace setting.
	//
	// Gets the enable partner powered AI features workspace setting.
	Get(ctx context.Context, request GetLlmProxyPartnerPoweredWorkspaceRequest) (*LlmProxyPartnerPoweredWorkspace, error)

	// Update the enable partner powered AI features workspace setting.
	//
	// Updates the enable partner powered AI features workspace setting.
	Update(ctx context.Context, request UpdateLlmProxyPartnerPoweredWorkspaceRequest) (*LlmProxyPartnerPoweredWorkspace, error)
}

// These APIs provide configurations for the network connectivity of your
// workspaces for serverless compute resources. This API provides stable subnets
// for your workspace so that you can configure your firewalls on your Azure
// Storage accounts to allow access from Databricks. You can also use the API to
// provision private endpoints for Databricks to privately connect serverless
// compute resources to your Azure resources using Azure Private Link. See
// [configure serverless secure connectivity].
//
// [configure serverless secure connectivity]: https://learn.microsoft.com/azure/databricks/security/network/serverless-network-security
type NetworkConnectivityService interface {

	// Create a network connectivity configuration.
	//
	// Creates a network connectivity configuration (NCC), which provides stable
	// Azure service subnets when accessing your Azure Storage accounts. You can
	// also use a network connectivity configuration to create Databricks
	// managed private endpoints so that Databricks serverless compute resources
	// privately access your resources.
	//
	// **IMPORTANT**: After you create the network connectivity configuration,
	// you must assign one or more workspaces to the new network connectivity
	// configuration. You can share one network connectivity configuration with
	// multiple workspaces from the same Azure region within the same Databricks
	// account. See [configure serverless secure connectivity].
	//
	// [configure serverless secure connectivity]: https://learn.microsoft.com/azure/databricks/security/network/serverless-network-security
	CreateNetworkConnectivityConfiguration(ctx context.Context, request CreateNetworkConnectivityConfigRequest) (*NetworkConnectivityConfiguration, error)

	// Create a private endpoint rule.
	//
	// Create a private endpoint rule for the specified network connectivity
	// config object. Once the object is created, Databricks asynchronously
	// provisions a new Azure private endpoint to your specified Azure resource.
	//
	// **IMPORTANT**: You must use Azure portal or other Azure tools to approve
	// the private endpoint to complete the connection. To get the information
	// of the private endpoint created, make a `GET` request on the new private
	// endpoint rule. See [serverless private link].
	//
	// [serverless private link]: https://learn.microsoft.com/azure/databricks/security/network/serverless-network-security/serverless-private-link
	CreatePrivateEndpointRule(ctx context.Context, request CreatePrivateEndpointRuleRequest) (*NccAzurePrivateEndpointRule, error)

	// Delete a network connectivity configuration.
	//
	// Deletes a network connectivity configuration.
	DeleteNetworkConnectivityConfiguration(ctx context.Context, request DeleteNetworkConnectivityConfigurationRequest) error

	// Delete a private endpoint rule.
	//
	// Initiates deleting a private endpoint rule. If the connection state is
	// PENDING or EXPIRED, the private endpoint is immediately deleted.
	// Otherwise, the private endpoint is deactivated and will be deleted after
	// seven days of deactivation. When a private endpoint is deactivated, the
	// `deactivated` field is set to `true` and the private endpoint is not
	// available to your serverless compute resources.
	DeletePrivateEndpointRule(ctx context.Context, request DeletePrivateEndpointRuleRequest) (*NccAzurePrivateEndpointRule, error)

	// Get a network connectivity configuration.
	//
	// Gets a network connectivity configuration.
	GetNetworkConnectivityConfiguration(ctx context.Context, request GetNetworkConnectivityConfigurationRequest) (*NetworkConnectivityConfiguration, error)

	// Gets a private endpoint rule.
	//
	// Gets the private endpoint rule.
	GetPrivateEndpointRule(ctx context.Context, request GetPrivateEndpointRuleRequest) (*NccAzurePrivateEndpointRule, error)

	// List network connectivity configurations.
	//
	// Gets an array of network connectivity configurations.
	//
	// Use ListNetworkConnectivityConfigurationsAll() to get all NetworkConnectivityConfiguration instances, which will iterate over every result page.
	ListNetworkConnectivityConfigurations(ctx context.Context, request ListNetworkConnectivityConfigurationsRequest) (*ListNetworkConnectivityConfigurationsResponse, error)

	// List private endpoint rules.
	//
	// Gets an array of private endpoint rules.
	//
	// Use ListPrivateEndpointRulesAll() to get all NccAzurePrivateEndpointRule instances, which will iterate over every result page.
	ListPrivateEndpointRules(ctx context.Context, request ListPrivateEndpointRulesRequest) (*ListNccAzurePrivateEndpointRulesResponse, error)

	// Update a private endpoint rule.
	//
	// Updates a private endpoint rule. Currently only a private endpoint rule
	// to customer-managed resources is allowed to be updated.
	UpdateNccAzurePrivateEndpointRulePublic(ctx context.Context, request UpdateNccAzurePrivateEndpointRulePublicRequest) (*NccAzurePrivateEndpointRule, error)
}

// These APIs manage network policies for this account. Network policies control
// which network destinations can be accessed from the Databricks environment.
// Each Databricks account includes a default policy named 'default-policy'.
// 'default-policy' is associated with any workspace lacking an explicit network
// policy assignment, and is automatically associated with each newly created
// workspace. 'default-policy' is reserved and cannot be deleted, but it can be
// updated to customize the default network access rules for your account.
type NetworkPoliciesService interface {

	// Create a network policy.
	//
	// Creates a new network policy to manage which network destinations can be
	// accessed from the Databricks environment.
	CreateNetworkPolicyRpc(ctx context.Context, request CreateNetworkPolicyRequest) (*AccountNetworkPolicy, error)

	// Delete a network policy.
	//
	// Deletes a network policy. Cannot be called on 'default-policy'.
	DeleteNetworkPolicyRpc(ctx context.Context, request DeleteNetworkPolicyRequest) error

	// Get a network policy.
	//
	// Gets a network policy.
	GetNetworkPolicyRpc(ctx context.Context, request GetNetworkPolicyRequest) (*AccountNetworkPolicy, error)

	// List network policies.
	//
	// Gets an array of network policies.
	//
	// Use ListNetworkPoliciesRpcAll() to get all AccountNetworkPolicy instances, which will iterate over every result page.
	ListNetworkPoliciesRpc(ctx context.Context, request ListNetworkPoliciesRequest) (*ListNetworkPoliciesResponse, error)

	// Update a network policy.
	//
	// Updates a network policy. This allows you to modify the configuration of
	// a network policy.
	UpdateNetworkPolicyRpc(ctx context.Context, request UpdateNetworkPolicyRequest) (*AccountNetworkPolicy, error)
}

// The notification destinations API lets you programmatically manage a
// workspace's notification destinations. Notification destinations are used to
// send notifications for query alerts and jobs to destinations outside of
// Databricks. Only workspace admins can create, update, and delete notification
// destinations.
type NotificationDestinationsService interface {

	// Create a notification destination.
	//
	// Creates a notification destination. Requires workspace admin permissions.
	Create(ctx context.Context, request CreateNotificationDestinationRequest) (*NotificationDestination, error)

	// Delete a notification destination.
	//
	// Deletes a notification destination. Requires workspace admin permissions.
	Delete(ctx context.Context, request DeleteNotificationDestinationRequest) error

	// Get a notification destination.
	//
	// Gets a notification destination.
	Get(ctx context.Context, request GetNotificationDestinationRequest) (*NotificationDestination, error)

	// List notification destinations.
	//
	// Lists notification destinations.
	//
	// Use ListAll() to get all ListNotificationDestinationsResult instances, which will iterate over every result page.
	List(ctx context.Context, request ListNotificationDestinationsRequest) (*ListNotificationDestinationsResponse, error)

	// Update a notification destination.
	//
	// Updates a notification destination. Requires workspace admin permissions.
	// At least one field is required in the request body.
	Update(ctx context.Context, request UpdateNotificationDestinationRequest) (*NotificationDestination, error)
}

// The Personal Compute enablement setting lets you control which users can use
// the Personal Compute default policy to create compute resources. By default
// all users in all workspaces have access (ON), but you can change the setting
// to instead let individual workspaces configure access control (DELEGATE).
//
// There is only one instance of this setting per account. Since this setting
// has a default value, this setting is present on all accounts even though it's
// never set on a given account. Deletion reverts the value of the setting back
// to the default value.
type PersonalComputeService interface {

	// Delete Personal Compute setting.
	//
	// Reverts back the Personal Compute setting value to default (ON)
	Delete(ctx context.Context, request DeletePersonalComputeSettingRequest) (*DeletePersonalComputeSettingResponse, error)

	// Get Personal Compute setting.
	//
	// Gets the value of the Personal Compute setting.
	Get(ctx context.Context, request GetPersonalComputeSettingRequest) (*PersonalComputeSetting, error)

	// Update Personal Compute setting.
	//
	// Updates the value of the Personal Compute setting.
	Update(ctx context.Context, request UpdatePersonalComputeSettingRequest) (*PersonalComputeSetting, error)
}

// The Restrict Workspace Admins setting lets you control the capabilities of
// workspace admins. With the setting status set to ALLOW_ALL, workspace admins
// can create service principal personal access tokens on behalf of any service
// principal in their workspace. Workspace admins can also change a job owner to
// any user in their workspace. And they can change the job run_as setting to
// any user in their workspace or to a service principal on which they have the
// Service Principal User role. With the setting status set to
// RESTRICT_TOKENS_AND_JOB_RUN_AS, workspace admins can only create personal
// access tokens on behalf of service principals they have the Service Principal
// User role on. They can also only change a job owner to themselves. And they
// can change the job run_as setting to themselves or to a service principal on
// which they have the Service Principal User role.
type RestrictWorkspaceAdminsService interface {

	// Delete the restrict workspace admins setting.
	//
	// Reverts the restrict workspace admins setting status for the workspace. A
	// fresh etag needs to be provided in `DELETE` requests (as a query
	// parameter). The etag can be retrieved by making a `GET` request before
	// the DELETE request. If the setting is updated/deleted concurrently,
	// `DELETE` fails with 409 and the request must be retried by using the
	// fresh etag in the 409 response.
	Delete(ctx context.Context, request DeleteRestrictWorkspaceAdminsSettingRequest) (*DeleteRestrictWorkspaceAdminsSettingResponse, error)

	// Get the restrict workspace admins setting.
	//
	// Gets the restrict workspace admins setting.
	Get(ctx context.Context, request GetRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error)

	// Update the restrict workspace admins setting.
	//
	// Updates the restrict workspace admins setting for the workspace. A fresh
	// etag needs to be provided in `PATCH` requests (as part of the setting
	// field). The etag can be retrieved by making a GET request before the
	// `PATCH` request. If the setting is updated concurrently, `PATCH` fails
	// with 409 and the request must be retried by using the fresh etag in the
	// 409 response.
	Update(ctx context.Context, request UpdateRestrictWorkspaceAdminsSettingRequest) (*RestrictWorkspaceAdminsSetting, error)
}

// Workspace Settings API allows users to manage settings at the workspace
// level.
type SettingsService interface {
}

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
type TokenManagementService interface {

	// Create on-behalf token.
	//
	// Creates a token on behalf of a service principal.
	CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error)

	// Delete a token.
	//
	// Deletes a token, specified by its ID.
	Delete(ctx context.Context, request DeleteTokenManagementRequest) error

	// Get token info.
	//
	// Gets information about a token, specified by its ID.
	Get(ctx context.Context, request GetTokenManagementRequest) (*GetTokenResponse, error)

	// Get token permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context) (*GetTokenPermissionLevelsResponse, error)

	// Get token permissions.
	//
	// Gets the permissions of all tokens. Tokens can inherit permissions from
	// their root object.
	GetPermissions(ctx context.Context) (*TokenPermissions, error)

	// List all tokens.
	//
	// Lists all tokens associated with the specified workspace or user.
	//
	// Use ListAll() to get all TokenInfo instances
	List(ctx context.Context, request ListTokenManagementRequest) (*ListTokensResponse, error)

	// Set token permissions.
	//
	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error)

	// Update token permissions.
	//
	// Updates the permissions on all tokens. Tokens can inherit permissions
	// from their root object.
	UpdatePermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error)
}

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
type TokensService interface {

	// Create a user token.
	//
	// Creates and returns a token for a user. If this call is made through
	// token authentication, it creates a token with the same client ID as the
	// authenticated token. If the user's token quota is exceeded, this call
	// returns an error **QUOTA_EXCEEDED**.
	Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error)

	// Revoke token.
	//
	// Revokes an access token.
	//
	// If a token with the specified ID is not valid, this call returns an error
	// **RESOURCE_DOES_NOT_EXIST**.
	Delete(ctx context.Context, request RevokeTokenRequest) error

	// List tokens.
	//
	// Lists all the valid tokens for a user-workspace pair.
	//
	// Use ListAll() to get all PublicTokenInfo instances
	List(ctx context.Context) (*ListPublicTokensResponse, error)
}

// This API allows updating known workspace settings for advanced users.
type WorkspaceConfService interface {

	// Check configuration status.
	//
	// Gets the configuration status for a workspace.
	GetStatus(ctx context.Context, request GetStatusRequest) (*map[string]string, error)

	// Enable/disable features.
	//
	// Sets the configuration status for a workspace, including enabling or
	// disabling it.
	SetStatus(ctx context.Context, request WorkspaceConf) error
}

// These APIs allow configuration of network settings for Databricks workspaces.
// Each workspace is always associated with exactly one network policy that
// controls which network destinations can be accessed from the Databricks
// environment. By default, workspaces are associated with the 'default-policy'
// network policy. You cannot create or delete a workspace's network
// configuration, only update it to associate the workspace with a different
// policy.
type WorkspaceNetworkConfigurationService interface {

	// Get workspace network configuration.
	//
	// Gets the network configuration for a workspace. Every workspace has
	// exactly one network policy binding, with 'default-policy' used if no
	// explicit assignment exists.
	GetWorkspaceNetworkOptionRpc(ctx context.Context, request GetWorkspaceNetworkOptionRequest) (*WorkspaceNetworkOption, error)

	// Update workspace network configuration.
	//
	// Updates the network configuration for a workspace. This operation
	// associates the workspace with the specified network policy. To revert to
	// the default policy, specify 'default-policy' as the network_policy_id.
	UpdateWorkspaceNetworkOptionRpc(ctx context.Context, request UpdateWorkspaceNetworkOptionRequest) (*WorkspaceNetworkOption, error)
}
