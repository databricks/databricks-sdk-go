// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplacepreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ConsumerFulfillmentsPreviewClient struct {
	ConsumerFulfillmentsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewConsumerFulfillmentsPreviewClient(cfg *config.Config) (*ConsumerFulfillmentsPreviewClient, error) {
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

	return &ConsumerFulfillmentsPreviewClient{
		Config:                               cfg,
		apiClient:                            apiClient,
		ConsumerFulfillmentsPreviewInterface: NewConsumerFulfillmentsPreview(databricksClient),
	}, nil
}

type ConsumerInstallationsPreviewClient struct {
	ConsumerInstallationsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewConsumerInstallationsPreviewClient(cfg *config.Config) (*ConsumerInstallationsPreviewClient, error) {
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

	return &ConsumerInstallationsPreviewClient{
		Config:                                cfg,
		apiClient:                             apiClient,
		ConsumerInstallationsPreviewInterface: NewConsumerInstallationsPreview(databricksClient),
	}, nil
}

type ConsumerListingsPreviewClient struct {
	ConsumerListingsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewConsumerListingsPreviewClient(cfg *config.Config) (*ConsumerListingsPreviewClient, error) {
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

	return &ConsumerListingsPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		ConsumerListingsPreviewInterface: NewConsumerListingsPreview(databricksClient),
	}, nil
}

type ConsumerPersonalizationRequestsPreviewClient struct {
	ConsumerPersonalizationRequestsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewConsumerPersonalizationRequestsPreviewClient(cfg *config.Config) (*ConsumerPersonalizationRequestsPreviewClient, error) {
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

	return &ConsumerPersonalizationRequestsPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		ConsumerPersonalizationRequestsPreviewInterface: NewConsumerPersonalizationRequestsPreview(databricksClient),
	}, nil
}

type ConsumerProvidersPreviewClient struct {
	ConsumerProvidersPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewConsumerProvidersPreviewClient(cfg *config.Config) (*ConsumerProvidersPreviewClient, error) {
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

	return &ConsumerProvidersPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		ConsumerProvidersPreviewInterface: NewConsumerProvidersPreview(databricksClient),
	}, nil
}

type ProviderExchangeFiltersPreviewClient struct {
	ProviderExchangeFiltersPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewProviderExchangeFiltersPreviewClient(cfg *config.Config) (*ProviderExchangeFiltersPreviewClient, error) {
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

	return &ProviderExchangeFiltersPreviewClient{
		Config:                                  cfg,
		apiClient:                               apiClient,
		ProviderExchangeFiltersPreviewInterface: NewProviderExchangeFiltersPreview(databricksClient),
	}, nil
}

type ProviderExchangesPreviewClient struct {
	ProviderExchangesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewProviderExchangesPreviewClient(cfg *config.Config) (*ProviderExchangesPreviewClient, error) {
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

	return &ProviderExchangesPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		ProviderExchangesPreviewInterface: NewProviderExchangesPreview(databricksClient),
	}, nil
}

type ProviderFilesPreviewClient struct {
	ProviderFilesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewProviderFilesPreviewClient(cfg *config.Config) (*ProviderFilesPreviewClient, error) {
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

	return &ProviderFilesPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		ProviderFilesPreviewInterface: NewProviderFilesPreview(databricksClient),
	}, nil
}

type ProviderListingsPreviewClient struct {
	ProviderListingsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewProviderListingsPreviewClient(cfg *config.Config) (*ProviderListingsPreviewClient, error) {
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

	return &ProviderListingsPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		ProviderListingsPreviewInterface: NewProviderListingsPreview(databricksClient),
	}, nil
}

type ProviderPersonalizationRequestsPreviewClient struct {
	ProviderPersonalizationRequestsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewProviderPersonalizationRequestsPreviewClient(cfg *config.Config) (*ProviderPersonalizationRequestsPreviewClient, error) {
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

	return &ProviderPersonalizationRequestsPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		ProviderPersonalizationRequestsPreviewInterface: NewProviderPersonalizationRequestsPreview(databricksClient),
	}, nil
}

type ProviderProviderAnalyticsDashboardsPreviewClient struct {
	ProviderProviderAnalyticsDashboardsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewProviderProviderAnalyticsDashboardsPreviewClient(cfg *config.Config) (*ProviderProviderAnalyticsDashboardsPreviewClient, error) {
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

	return &ProviderProviderAnalyticsDashboardsPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		ProviderProviderAnalyticsDashboardsPreviewInterface: NewProviderProviderAnalyticsDashboardsPreview(databricksClient),
	}, nil
}

type ProviderProvidersPreviewClient struct {
	ProviderProvidersPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewProviderProvidersPreviewClient(cfg *config.Config) (*ProviderProvidersPreviewClient, error) {
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

	return &ProviderProvidersPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		ProviderProvidersPreviewInterface: NewProviderProvidersPreview(databricksClient),
	}, nil
}
