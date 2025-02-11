// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingspreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccountIpAccessListsPreviewClient struct {
	AccountIpAccessListsPreviewInterface

	Config *config.Config
}

func NewAccountIpAccessListsPreviewClient(cfg *config.Config) (*AccountIpAccessListsPreviewClient, error) {
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

	return &AccountIpAccessListsPreviewClient{
		Config:                               cfg,
		AccountIpAccessListsPreviewInterface: NewAccountIpAccessListsPreview(apiClient),
	}, nil
}

type AccountSettingsClient struct {
	AccountSettingsInterface

	Config *config.Config
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
		Config:                   cfg,
		AccountSettingsInterface: NewAccountSettings(apiClient),
	}, nil
}

type AccountSettingsPreviewClient struct {
	AccountSettingsPreviewInterface

	Config *config.Config
}

func NewAccountSettingsPreviewClient(cfg *config.Config) (*AccountSettingsPreviewClient, error) {
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

	return &AccountSettingsPreviewClient{
		Config:                          cfg,
		AccountSettingsPreviewInterface: NewAccountSettingsPreview(apiClient),
	}, nil
}

type AibiDashboardEmbeddingAccessPolicyPreviewClient struct {
	AibiDashboardEmbeddingAccessPolicyPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewAibiDashboardEmbeddingAccessPolicyPreviewClient(cfg *config.Config) (*AibiDashboardEmbeddingAccessPolicyPreviewClient, error) {
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

	return &AibiDashboardEmbeddingAccessPolicyPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		AibiDashboardEmbeddingAccessPolicyPreviewInterface: NewAibiDashboardEmbeddingAccessPolicyPreview(databricksClient),
	}, nil
}

type AibiDashboardEmbeddingApprovedDomainsPreviewClient struct {
	AibiDashboardEmbeddingApprovedDomainsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewAibiDashboardEmbeddingApprovedDomainsPreviewClient(cfg *config.Config) (*AibiDashboardEmbeddingApprovedDomainsPreviewClient, error) {
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

	return &AibiDashboardEmbeddingApprovedDomainsPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		AibiDashboardEmbeddingApprovedDomainsPreviewInterface: NewAibiDashboardEmbeddingApprovedDomainsPreview(databricksClient),
	}, nil
}

type AutomaticClusterUpdatePreviewClient struct {
	AutomaticClusterUpdatePreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewAutomaticClusterUpdatePreviewClient(cfg *config.Config) (*AutomaticClusterUpdatePreviewClient, error) {
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

	return &AutomaticClusterUpdatePreviewClient{
		Config:                                 cfg,
		apiClient:                              apiClient,
		AutomaticClusterUpdatePreviewInterface: NewAutomaticClusterUpdatePreview(databricksClient),
	}, nil
}

type ComplianceSecurityProfilePreviewClient struct {
	ComplianceSecurityProfilePreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewComplianceSecurityProfilePreviewClient(cfg *config.Config) (*ComplianceSecurityProfilePreviewClient, error) {
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

	return &ComplianceSecurityProfilePreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		ComplianceSecurityProfilePreviewInterface: NewComplianceSecurityProfilePreview(databricksClient),
	}, nil
}

type CredentialsManagerPreviewClient struct {
	CredentialsManagerPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewCredentialsManagerPreviewClient(cfg *config.Config) (*CredentialsManagerPreviewClient, error) {
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

	return &CredentialsManagerPreviewClient{
		Config:                             cfg,
		apiClient:                          apiClient,
		CredentialsManagerPreviewInterface: NewCredentialsManagerPreview(databricksClient),
	}, nil
}

type CspEnablementAccountPreviewClient struct {
	CspEnablementAccountPreviewInterface

	Config *config.Config
}

func NewCspEnablementAccountPreviewClient(cfg *config.Config) (*CspEnablementAccountPreviewClient, error) {
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

	return &CspEnablementAccountPreviewClient{
		Config:                               cfg,
		CspEnablementAccountPreviewInterface: NewCspEnablementAccountPreview(apiClient),
	}, nil
}

type DefaultNamespacePreviewClient struct {
	DefaultNamespacePreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewDefaultNamespacePreviewClient(cfg *config.Config) (*DefaultNamespacePreviewClient, error) {
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

	return &DefaultNamespacePreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		DefaultNamespacePreviewInterface: NewDefaultNamespacePreview(databricksClient),
	}, nil
}

type DisableLegacyAccessPreviewClient struct {
	DisableLegacyAccessPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewDisableLegacyAccessPreviewClient(cfg *config.Config) (*DisableLegacyAccessPreviewClient, error) {
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

	return &DisableLegacyAccessPreviewClient{
		Config:                              cfg,
		apiClient:                           apiClient,
		DisableLegacyAccessPreviewInterface: NewDisableLegacyAccessPreview(databricksClient),
	}, nil
}

type DisableLegacyDbfsPreviewClient struct {
	DisableLegacyDbfsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewDisableLegacyDbfsPreviewClient(cfg *config.Config) (*DisableLegacyDbfsPreviewClient, error) {
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

	return &DisableLegacyDbfsPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		DisableLegacyDbfsPreviewInterface: NewDisableLegacyDbfsPreview(databricksClient),
	}, nil
}

type DisableLegacyFeaturesPreviewClient struct {
	DisableLegacyFeaturesPreviewInterface

	Config *config.Config
}

func NewDisableLegacyFeaturesPreviewClient(cfg *config.Config) (*DisableLegacyFeaturesPreviewClient, error) {
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

	return &DisableLegacyFeaturesPreviewClient{
		Config:                                cfg,
		DisableLegacyFeaturesPreviewInterface: NewDisableLegacyFeaturesPreview(apiClient),
	}, nil
}

type EnableIpAccessListsPreviewClient struct {
	EnableIpAccessListsPreviewInterface

	Config *config.Config
}

func NewEnableIpAccessListsPreviewClient(cfg *config.Config) (*EnableIpAccessListsPreviewClient, error) {
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

	return &EnableIpAccessListsPreviewClient{
		Config:                              cfg,
		EnableIpAccessListsPreviewInterface: NewEnableIpAccessListsPreview(apiClient),
	}, nil
}

type EnhancedSecurityMonitoringPreviewClient struct {
	EnhancedSecurityMonitoringPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewEnhancedSecurityMonitoringPreviewClient(cfg *config.Config) (*EnhancedSecurityMonitoringPreviewClient, error) {
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

	return &EnhancedSecurityMonitoringPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		EnhancedSecurityMonitoringPreviewInterface: NewEnhancedSecurityMonitoringPreview(databricksClient),
	}, nil
}

type EsmEnablementAccountPreviewClient struct {
	EsmEnablementAccountPreviewInterface

	Config *config.Config
}

func NewEsmEnablementAccountPreviewClient(cfg *config.Config) (*EsmEnablementAccountPreviewClient, error) {
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

	return &EsmEnablementAccountPreviewClient{
		Config:                               cfg,
		EsmEnablementAccountPreviewInterface: NewEsmEnablementAccountPreview(apiClient),
	}, nil
}

type IpAccessListsPreviewClient struct {
	IpAccessListsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewIpAccessListsPreviewClient(cfg *config.Config) (*IpAccessListsPreviewClient, error) {
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

	return &IpAccessListsPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		IpAccessListsPreviewInterface: NewIpAccessListsPreview(databricksClient),
	}, nil
}

type NetworkConnectivityPreviewClient struct {
	NetworkConnectivityPreviewInterface

	Config *config.Config
}

func NewNetworkConnectivityPreviewClient(cfg *config.Config) (*NetworkConnectivityPreviewClient, error) {
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

	return &NetworkConnectivityPreviewClient{
		Config:                              cfg,
		NetworkConnectivityPreviewInterface: NewNetworkConnectivityPreview(apiClient),
	}, nil
}

type NotificationDestinationsPreviewClient struct {
	NotificationDestinationsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewNotificationDestinationsPreviewClient(cfg *config.Config) (*NotificationDestinationsPreviewClient, error) {
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

	return &NotificationDestinationsPreviewClient{
		Config:                                   cfg,
		apiClient:                                apiClient,
		NotificationDestinationsPreviewInterface: NewNotificationDestinationsPreview(databricksClient),
	}, nil
}

type PersonalComputePreviewClient struct {
	PersonalComputePreviewInterface

	Config *config.Config
}

func NewPersonalComputePreviewClient(cfg *config.Config) (*PersonalComputePreviewClient, error) {
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

	return &PersonalComputePreviewClient{
		Config:                          cfg,
		PersonalComputePreviewInterface: NewPersonalComputePreview(apiClient),
	}, nil
}

type RestrictWorkspaceAdminsPreviewClient struct {
	RestrictWorkspaceAdminsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewRestrictWorkspaceAdminsPreviewClient(cfg *config.Config) (*RestrictWorkspaceAdminsPreviewClient, error) {
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

	return &RestrictWorkspaceAdminsPreviewClient{
		Config:                                  cfg,
		apiClient:                               apiClient,
		RestrictWorkspaceAdminsPreviewInterface: NewRestrictWorkspaceAdminsPreview(databricksClient),
	}, nil
}

type SettingsClient struct {
	SettingsInterface
	Config    *config.Config
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
		Config:            cfg,
		apiClient:         apiClient,
		SettingsInterface: NewSettings(databricksClient),
	}, nil
}

type SettingsPreviewClient struct {
	SettingsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewSettingsPreviewClient(cfg *config.Config) (*SettingsPreviewClient, error) {
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

	return &SettingsPreviewClient{
		Config:                   cfg,
		apiClient:                apiClient,
		SettingsPreviewInterface: NewSettingsPreview(databricksClient),
	}, nil
}

type TokenManagementPreviewClient struct {
	TokenManagementPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewTokenManagementPreviewClient(cfg *config.Config) (*TokenManagementPreviewClient, error) {
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

	return &TokenManagementPreviewClient{
		Config:                          cfg,
		apiClient:                       apiClient,
		TokenManagementPreviewInterface: NewTokenManagementPreview(databricksClient),
	}, nil
}

type TokensPreviewClient struct {
	TokensPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewTokensPreviewClient(cfg *config.Config) (*TokensPreviewClient, error) {
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

	return &TokensPreviewClient{
		Config:                 cfg,
		apiClient:              apiClient,
		TokensPreviewInterface: NewTokensPreview(databricksClient),
	}, nil
}

type WorkspaceConfPreviewClient struct {
	WorkspaceConfPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewWorkspaceConfPreviewClient(cfg *config.Config) (*WorkspaceConfPreviewClient, error) {
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

	return &WorkspaceConfPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		WorkspaceConfPreviewInterface: NewWorkspaceConfPreview(databricksClient),
	}, nil
}
