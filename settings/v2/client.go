// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccountIpAccessListsClient struct {
	cfg *config.Config

	AccountIpAccessLists AccountIpAccessListsInterface
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
		cfg:                  cfg,
		AccountIpAccessLists: NewAccountIpAccessLists(apiClient),
	}, nil
}

type AccountSettingsClient struct {
	cfg *config.Config

	AccountSettings AccountSettingsInterface
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
		cfg:             cfg,
		AccountSettings: NewAccountSettings(apiClient),
	}, nil
}

type AibiDashboardEmbeddingAccessPolicyClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	AibiDashboardEmbeddingAccessPolicy AibiDashboardEmbeddingAccessPolicyInterface
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
		cfg:                                cfg,
		apiClient:                          apiClient,
		AibiDashboardEmbeddingAccessPolicy: NewAibiDashboardEmbeddingAccessPolicy(databricksClient),
	}, nil
}

type AibiDashboardEmbeddingApprovedDomainsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	AibiDashboardEmbeddingApprovedDomains AibiDashboardEmbeddingApprovedDomainsInterface
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
		cfg:                                   cfg,
		apiClient:                             apiClient,
		AibiDashboardEmbeddingApprovedDomains: NewAibiDashboardEmbeddingApprovedDomains(databricksClient),
	}, nil
}

type AutomaticClusterUpdateClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	AutomaticClusterUpdate AutomaticClusterUpdateInterface
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
		cfg:                    cfg,
		apiClient:              apiClient,
		AutomaticClusterUpdate: NewAutomaticClusterUpdate(databricksClient),
	}, nil
}

type ComplianceSecurityProfileClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ComplianceSecurityProfile ComplianceSecurityProfileInterface
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		ComplianceSecurityProfile: NewComplianceSecurityProfile(databricksClient),
	}, nil
}

type CredentialsManagerClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	CredentialsManager CredentialsManagerInterface
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
		cfg:                cfg,
		apiClient:          apiClient,
		CredentialsManager: NewCredentialsManager(databricksClient),
	}, nil
}

type CspEnablementAccountClient struct {
	cfg *config.Config

	CspEnablementAccount CspEnablementAccountInterface
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
		cfg:                  cfg,
		CspEnablementAccount: NewCspEnablementAccount(apiClient),
	}, nil
}

type DefaultNamespaceClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	DefaultNamespace DefaultNamespaceInterface
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
		cfg:              cfg,
		apiClient:        apiClient,
		DefaultNamespace: NewDefaultNamespace(databricksClient),
	}, nil
}

type DisableLegacyAccessClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	DisableLegacyAccess DisableLegacyAccessInterface
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
		cfg:                 cfg,
		apiClient:           apiClient,
		DisableLegacyAccess: NewDisableLegacyAccess(databricksClient),
	}, nil
}

type DisableLegacyDbfsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	DisableLegacyDbfs DisableLegacyDbfsInterface
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
		cfg:               cfg,
		apiClient:         apiClient,
		DisableLegacyDbfs: NewDisableLegacyDbfs(databricksClient),
	}, nil
}

type DisableLegacyFeaturesClient struct {
	cfg *config.Config

	DisableLegacyFeatures DisableLegacyFeaturesInterface
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
		cfg:                   cfg,
		DisableLegacyFeatures: NewDisableLegacyFeatures(apiClient),
	}, nil
}

type EnableIpAccessListsClient struct {
	cfg *config.Config

	EnableIpAccessLists EnableIpAccessListsInterface
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
		cfg:                 cfg,
		EnableIpAccessLists: NewEnableIpAccessLists(apiClient),
	}, nil
}

type EnhancedSecurityMonitoringClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	EnhancedSecurityMonitoring EnhancedSecurityMonitoringInterface
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
		cfg:                        cfg,
		apiClient:                  apiClient,
		EnhancedSecurityMonitoring: NewEnhancedSecurityMonitoring(databricksClient),
	}, nil
}

type EsmEnablementAccountClient struct {
	cfg *config.Config

	EsmEnablementAccount EsmEnablementAccountInterface
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
		cfg:                  cfg,
		EsmEnablementAccount: NewEsmEnablementAccount(apiClient),
	}, nil
}

type IpAccessListsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	IpAccessLists IpAccessListsInterface
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
		cfg:           cfg,
		apiClient:     apiClient,
		IpAccessLists: NewIpAccessLists(databricksClient),
	}, nil
}

type NetworkConnectivityClient struct {
	cfg *config.Config

	NetworkConnectivity NetworkConnectivityInterface
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
		cfg:                 cfg,
		NetworkConnectivity: NewNetworkConnectivity(apiClient),
	}, nil
}

type NotificationDestinationsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	NotificationDestinations NotificationDestinationsInterface
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
		cfg:                      cfg,
		apiClient:                apiClient,
		NotificationDestinations: NewNotificationDestinations(databricksClient),
	}, nil
}

type PersonalComputeClient struct {
	cfg *config.Config

	PersonalCompute PersonalComputeInterface
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
		cfg:             cfg,
		PersonalCompute: NewPersonalCompute(apiClient),
	}, nil
}

type RestrictWorkspaceAdminsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	RestrictWorkspaceAdmins RestrictWorkspaceAdminsInterface
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
		cfg:                     cfg,
		apiClient:               apiClient,
		RestrictWorkspaceAdmins: NewRestrictWorkspaceAdmins(databricksClient),
	}, nil
}

type SettingsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Settings SettingsInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Settings:  NewSettings(databricksClient),
	}, nil
}

type TokenManagementClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	TokenManagement TokenManagementInterface
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
		cfg:             cfg,
		apiClient:       apiClient,
		TokenManagement: NewTokenManagement(databricksClient),
	}, nil
}

type TokensClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Tokens TokensInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Tokens:    NewTokens(databricksClient),
	}, nil
}

type WorkspaceConfClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	WorkspaceConf WorkspaceConfInterface
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
		cfg:           cfg,
		apiClient:     apiClient,
		WorkspaceConf: NewWorkspaceConf(databricksClient),
	}, nil
}
