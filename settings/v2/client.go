// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccountIpAccessListsClient struct {
	AccountIpAccessListsInterface

	cfg *config.Config
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
		cfg:                           cfg,
		AccountIpAccessListsInterface: NewAccountIpAccessLists(apiClient),
	}, nil
}

type AccountSettingsClient struct {
	AccountSettingsInterface

	cfg *config.Config
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
		cfg:                      cfg,
		AccountSettingsInterface: NewAccountSettings(apiClient),
	}, nil
}

type AibiDashboardEmbeddingAccessPolicyClient struct {
	AibiDashboardEmbeddingAccessPolicyInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &AibiDashboardEmbeddingAccessPolicyClient{
		cfg:       cfg,
		apiClient: apiClient,
		AibiDashboardEmbeddingAccessPolicyInterface: NewAibiDashboardEmbeddingAccessPolicy(databricksClient),
	}, nil
}

type AibiDashboardEmbeddingApprovedDomainsClient struct {
	AibiDashboardEmbeddingApprovedDomainsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &AibiDashboardEmbeddingApprovedDomainsClient{
		cfg:       cfg,
		apiClient: apiClient,
		AibiDashboardEmbeddingApprovedDomainsInterface: NewAibiDashboardEmbeddingApprovedDomains(databricksClient),
	}, nil
}

type AutomaticClusterUpdateClient struct {
	AutomaticClusterUpdateInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &AutomaticClusterUpdateClient{
		cfg:                             cfg,
		apiClient:                       apiClient,
		AutomaticClusterUpdateInterface: NewAutomaticClusterUpdate(databricksClient),
	}, nil
}

type ComplianceSecurityProfileClient struct {
	ComplianceSecurityProfileInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ComplianceSecurityProfileClient{
		cfg:                                cfg,
		apiClient:                          apiClient,
		ComplianceSecurityProfileInterface: NewComplianceSecurityProfile(databricksClient),
	}, nil
}

type CredentialsManagerClient struct {
	CredentialsManagerInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &CredentialsManagerClient{
		cfg:                         cfg,
		apiClient:                   apiClient,
		CredentialsManagerInterface: NewCredentialsManager(databricksClient),
	}, nil
}

type CspEnablementAccountClient struct {
	CspEnablementAccountInterface

	cfg *config.Config
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
		cfg:                           cfg,
		CspEnablementAccountInterface: NewCspEnablementAccount(apiClient),
	}, nil
}

type DefaultNamespaceClient struct {
	DefaultNamespaceInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &DefaultNamespaceClient{
		cfg:                       cfg,
		apiClient:                 apiClient,
		DefaultNamespaceInterface: NewDefaultNamespace(databricksClient),
	}, nil
}

type DisableLegacyAccessClient struct {
	DisableLegacyAccessInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &DisableLegacyAccessClient{
		cfg:                          cfg,
		apiClient:                    apiClient,
		DisableLegacyAccessInterface: NewDisableLegacyAccess(databricksClient),
	}, nil
}

type DisableLegacyDbfsClient struct {
	DisableLegacyDbfsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &DisableLegacyDbfsClient{
		cfg:                        cfg,
		apiClient:                  apiClient,
		DisableLegacyDbfsInterface: NewDisableLegacyDbfs(databricksClient),
	}, nil
}

type DisableLegacyFeaturesClient struct {
	DisableLegacyFeaturesInterface

	cfg *config.Config
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
		cfg:                            cfg,
		DisableLegacyFeaturesInterface: NewDisableLegacyFeatures(apiClient),
	}, nil
}

type EnableIpAccessListsClient struct {
	EnableIpAccessListsInterface

	cfg *config.Config
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
		cfg:                          cfg,
		EnableIpAccessListsInterface: NewEnableIpAccessLists(apiClient),
	}, nil
}

type EnhancedSecurityMonitoringClient struct {
	EnhancedSecurityMonitoringInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &EnhancedSecurityMonitoringClient{
		cfg:                                 cfg,
		apiClient:                           apiClient,
		EnhancedSecurityMonitoringInterface: NewEnhancedSecurityMonitoring(databricksClient),
	}, nil
}

type EsmEnablementAccountClient struct {
	EsmEnablementAccountInterface

	cfg *config.Config
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
		cfg:                           cfg,
		EsmEnablementAccountInterface: NewEsmEnablementAccount(apiClient),
	}, nil
}

type IpAccessListsClient struct {
	IpAccessListsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &IpAccessListsClient{
		cfg:                    cfg,
		apiClient:              apiClient,
		IpAccessListsInterface: NewIpAccessLists(databricksClient),
	}, nil
}

type NetworkConnectivityClient struct {
	NetworkConnectivityInterface

	cfg *config.Config
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
		cfg:                          cfg,
		NetworkConnectivityInterface: NewNetworkConnectivity(apiClient),
	}, nil
}

type NotificationDestinationsClient struct {
	NotificationDestinationsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &NotificationDestinationsClient{
		cfg:                               cfg,
		apiClient:                         apiClient,
		NotificationDestinationsInterface: NewNotificationDestinations(databricksClient),
	}, nil
}

type PersonalComputeClient struct {
	PersonalComputeInterface

	cfg *config.Config
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
		cfg:                      cfg,
		PersonalComputeInterface: NewPersonalCompute(apiClient),
	}, nil
}

type RestrictWorkspaceAdminsClient struct {
	RestrictWorkspaceAdminsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &RestrictWorkspaceAdminsClient{
		cfg:                              cfg,
		apiClient:                        apiClient,
		RestrictWorkspaceAdminsInterface: NewRestrictWorkspaceAdmins(databricksClient),
	}, nil
}

type SettingsClient struct {
	SettingsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &SettingsClient{
		cfg:               cfg,
		apiClient:         apiClient,
		SettingsInterface: NewSettings(databricksClient),
	}, nil
}

type TokenManagementClient struct {
	TokenManagementInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &TokenManagementClient{
		cfg:                      cfg,
		apiClient:                apiClient,
		TokenManagementInterface: NewTokenManagement(databricksClient),
	}, nil
}

type TokensClient struct {
	TokensInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &TokensClient{
		cfg:             cfg,
		apiClient:       apiClient,
		TokensInterface: NewTokens(databricksClient),
	}, nil
}

type WorkspaceConfClient struct {
	WorkspaceConfInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &WorkspaceConfClient{
		cfg:                    cfg,
		apiClient:              apiClient,
		WorkspaceConfInterface: NewWorkspaceConf(databricksClient),
	}, nil
}
