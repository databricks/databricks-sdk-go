// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

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
