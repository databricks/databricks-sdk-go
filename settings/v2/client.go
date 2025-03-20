// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
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
type AccountIpAccessListsClient struct {
	AccountIpAccessListsAPI
}

func NewAccountIpAccessListsClient(cfg *config.Config) (*AccountIpAccessListsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AccountIpAccessListsClient{
		AccountIpAccessListsAPI: AccountIpAccessListsAPI{
			accountIpAccessListsImpl: accountIpAccessListsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Accounts Settings API allows users to manage settings at the account level.
type AccountSettingsClient struct {
	AccountSettingsAPI
}

func NewAccountSettingsClient(cfg *config.Config) (*AccountSettingsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AccountSettingsClient{
		AccountSettingsAPI: AccountSettingsAPI{
			accountSettingsImpl: accountSettingsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Controls whether AI/BI published dashboard embedding is enabled,
// conditionally enabled, or disabled at the workspace level. By default, this
// setting is conditionally enabled (ALLOW_APPROVED_DOMAINS).
type AibiDashboardEmbeddingAccessPolicyClient struct {
	AibiDashboardEmbeddingAccessPolicyAPI
}

func NewAibiDashboardEmbeddingAccessPolicyClient(cfg *config.Config) (*AibiDashboardEmbeddingAccessPolicyClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AibiDashboardEmbeddingAccessPolicyClient{
		AibiDashboardEmbeddingAccessPolicyAPI: AibiDashboardEmbeddingAccessPolicyAPI{
			aibiDashboardEmbeddingAccessPolicyImpl: aibiDashboardEmbeddingAccessPolicyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Controls the list of domains approved to host the embedded AI/BI dashboards.
// The approved domains list can't be mutated when the current access policy is
// not set to ALLOW_APPROVED_DOMAINS.
type AibiDashboardEmbeddingApprovedDomainsClient struct {
	AibiDashboardEmbeddingApprovedDomainsAPI
}

func NewAibiDashboardEmbeddingApprovedDomainsClient(cfg *config.Config) (*AibiDashboardEmbeddingApprovedDomainsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AibiDashboardEmbeddingApprovedDomainsClient{
		AibiDashboardEmbeddingApprovedDomainsAPI: AibiDashboardEmbeddingApprovedDomainsAPI{
			aibiDashboardEmbeddingApprovedDomainsImpl: aibiDashboardEmbeddingApprovedDomainsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Controls whether automatic cluster update is enabled for the current
// workspace. By default, it is turned off.
type AutomaticClusterUpdateClient struct {
	AutomaticClusterUpdateAPI
}

func NewAutomaticClusterUpdateClient(cfg *config.Config) (*AutomaticClusterUpdateClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AutomaticClusterUpdateClient{
		AutomaticClusterUpdateAPI: AutomaticClusterUpdateAPI{
			automaticClusterUpdateImpl: automaticClusterUpdateImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Controls whether to enable the compliance security profile for the current
// workspace. Enabling it on a workspace is permanent. By default, it is turned
// off.
//
// This settings can NOT be disabled once it is enabled.
type ComplianceSecurityProfileClient struct {
	ComplianceSecurityProfileAPI
}

func NewComplianceSecurityProfileClient(cfg *config.Config) (*ComplianceSecurityProfileClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ComplianceSecurityProfileClient{
		ComplianceSecurityProfileAPI: ComplianceSecurityProfileAPI{
			complianceSecurityProfileImpl: complianceSecurityProfileImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Credentials manager interacts with with Identity Providers to to perform
// token exchanges using stored credentials and refresh tokens.
type CredentialsManagerClient struct {
	CredentialsManagerAPI
}

func NewCredentialsManagerClient(cfg *config.Config) (*CredentialsManagerClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CredentialsManagerClient{
		CredentialsManagerAPI: CredentialsManagerAPI{
			credentialsManagerImpl: credentialsManagerImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The compliance security profile settings at the account level control whether
// to enable it for new workspaces. By default, this account-level setting is
// disabled for new workspaces. After workspace creation, account admins can
// enable the compliance security profile individually for each workspace.
//
// This settings can be disabled so that new workspaces do not have compliance
// security profile enabled by default.
type CspEnablementAccountClient struct {
	CspEnablementAccountAPI
}

func NewCspEnablementAccountClient(cfg *config.Config) (*CspEnablementAccountClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CspEnablementAccountClient{
		CspEnablementAccountAPI: CspEnablementAccountAPI{
			cspEnablementAccountImpl: cspEnablementAccountImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
type DefaultNamespaceClient struct {
	DefaultNamespaceAPI
}

func NewDefaultNamespaceClient(cfg *config.Config) (*DefaultNamespaceClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DefaultNamespaceClient{
		DefaultNamespaceAPI: DefaultNamespaceAPI{
			defaultNamespaceImpl: defaultNamespaceImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// 'Disabling legacy access' has the following impacts:
//
// 1. Disables direct access to the Hive Metastore. However, you can still
// access Hive Metastore through HMS Federation. 2. Disables Fallback Mode (docs
// link) on any External Location access from the workspace. 3. Alters DBFS path
// access to use External Location permissions in place of legacy credentials.
// 4. Enforces Unity Catalog access on all path based access.
type DisableLegacyAccessClient struct {
	DisableLegacyAccessAPI
}

func NewDisableLegacyAccessClient(cfg *config.Config) (*DisableLegacyAccessClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DisableLegacyAccessClient{
		DisableLegacyAccessAPI: DisableLegacyAccessAPI{
			disableLegacyAccessImpl: disableLegacyAccessImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// When this setting is on, access to DBFS root and DBFS mounts is disallowed
// (as well as creation of new mounts). When the setting is off, all DBFS
// functionality is enabled
type DisableLegacyDbfsClient struct {
	DisableLegacyDbfsAPI
}

func NewDisableLegacyDbfsClient(cfg *config.Config) (*DisableLegacyDbfsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DisableLegacyDbfsClient{
		DisableLegacyDbfsAPI: DisableLegacyDbfsAPI{
			disableLegacyDbfsImpl: disableLegacyDbfsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Disable legacy features for new Databricks workspaces.
//
// For newly created workspaces: 1. Disables the use of DBFS root and mounts. 2.
// Hive Metastore will not be provisioned. 3. Disables the use of
// ‘No-isolation clusters’. 4. Disables Databricks Runtime versions prior to
// 13.3LTS.
type DisableLegacyFeaturesClient struct {
	DisableLegacyFeaturesAPI
}

func NewDisableLegacyFeaturesClient(cfg *config.Config) (*DisableLegacyFeaturesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DisableLegacyFeaturesClient{
		DisableLegacyFeaturesAPI: DisableLegacyFeaturesAPI{
			disableLegacyFeaturesImpl: disableLegacyFeaturesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Controls the enforcement of IP access lists for accessing the account
// console. Allowing you to enable or disable restricted access based on IP
// addresses.
type EnableIpAccessListsClient struct {
	EnableIpAccessListsAPI
}

func NewEnableIpAccessListsClient(cfg *config.Config) (*EnableIpAccessListsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &EnableIpAccessListsClient{
		EnableIpAccessListsAPI: EnableIpAccessListsAPI{
			enableIpAccessListsImpl: enableIpAccessListsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Controls whether enhanced security monitoring is enabled for the current
// workspace. If the compliance security profile is enabled, this is
// automatically enabled. By default, it is disabled. However, if the compliance
// security profile is enabled, this is automatically enabled.
//
// If the compliance security profile is disabled, you can enable or disable
// this setting and it is not permanent.
type EnhancedSecurityMonitoringClient struct {
	EnhancedSecurityMonitoringAPI
}

func NewEnhancedSecurityMonitoringClient(cfg *config.Config) (*EnhancedSecurityMonitoringClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &EnhancedSecurityMonitoringClient{
		EnhancedSecurityMonitoringAPI: EnhancedSecurityMonitoringAPI{
			enhancedSecurityMonitoringImpl: enhancedSecurityMonitoringImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The enhanced security monitoring setting at the account level controls
// whether to enable the feature on new workspaces. By default, this
// account-level setting is disabled for new workspaces. After workspace
// creation, account admins can enable enhanced security monitoring individually
// for each workspace.
type EsmEnablementAccountClient struct {
	EsmEnablementAccountAPI
}

func NewEsmEnablementAccountClient(cfg *config.Config) (*EsmEnablementAccountClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &EsmEnablementAccountClient{
		EsmEnablementAccountAPI: EsmEnablementAccountAPI{
			esmEnablementAccountImpl: esmEnablementAccountImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
type IpAccessListsClient struct {
	IpAccessListsAPI
}

func NewIpAccessListsClient(cfg *config.Config) (*IpAccessListsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &IpAccessListsClient{
		IpAccessListsAPI: IpAccessListsAPI{
			ipAccessListsImpl: ipAccessListsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs provide configurations for the network connectivity of your
// workspaces for serverless compute resources.
type NetworkConnectivityClient struct {
	NetworkConnectivityAPI
}

func NewNetworkConnectivityClient(cfg *config.Config) (*NetworkConnectivityClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &NetworkConnectivityClient{
		NetworkConnectivityAPI: NetworkConnectivityAPI{
			networkConnectivityImpl: networkConnectivityImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The notification destinations API lets you programmatically manage a
// workspace's notification destinations. Notification destinations are used to
// send notifications for query alerts and jobs to destinations outside of
// Databricks. Only workspace admins can create, update, and delete notification
// destinations.
type NotificationDestinationsClient struct {
	NotificationDestinationsAPI
}

func NewNotificationDestinationsClient(cfg *config.Config) (*NotificationDestinationsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &NotificationDestinationsClient{
		NotificationDestinationsAPI: NotificationDestinationsAPI{
			notificationDestinationsImpl: notificationDestinationsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
type PersonalComputeClient struct {
	PersonalComputeAPI
}

func NewPersonalComputeClient(cfg *config.Config) (*PersonalComputeClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PersonalComputeClient{
		PersonalComputeAPI: PersonalComputeAPI{
			personalComputeImpl: personalComputeImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
type RestrictWorkspaceAdminsClient struct {
	RestrictWorkspaceAdminsAPI
}

func NewRestrictWorkspaceAdminsClient(cfg *config.Config) (*RestrictWorkspaceAdminsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RestrictWorkspaceAdminsClient{
		RestrictWorkspaceAdminsAPI: RestrictWorkspaceAdminsAPI{
			restrictWorkspaceAdminsImpl: restrictWorkspaceAdminsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Workspace Settings API allows users to manage settings at the workspace
// level.
type SettingsClient struct {
	SettingsAPI
}

func NewSettingsClient(cfg *config.Config) (*SettingsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &SettingsClient{
		SettingsAPI: SettingsAPI{
			settingsImpl: settingsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
type TokenManagementClient struct {
	TokenManagementAPI
}

func NewTokenManagementClient(cfg *config.Config) (*TokenManagementClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &TokenManagementClient{
		TokenManagementAPI: TokenManagementAPI{
			tokenManagementImpl: tokenManagementImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
type TokensClient struct {
	TokensAPI
}

func NewTokensClient(cfg *config.Config) (*TokensClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &TokensClient{
		TokensAPI: TokensAPI{
			tokensImpl: tokensImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// This API allows updating known workspace settings for advanced users.
type WorkspaceConfClient struct {
	WorkspaceConfAPI
}

func NewWorkspaceConfClient(cfg *config.Config) (*WorkspaceConfClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WorkspaceConfClient{
		WorkspaceConfAPI: WorkspaceConfAPI{
			workspaceConfImpl: workspaceConfImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
