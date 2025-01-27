// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ConsumerFulfillmentsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ConsumerFulfillments ConsumerFulfillmentsInterface
}

func NewConsumerFulfillmentsClient(cfg *config.Config) (*ConsumerFulfillmentsClient, error) {
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

	return &ConsumerFulfillmentsClient{
		cfg:                  cfg,
		apiClient:            apiClient,
		ConsumerFulfillments: NewConsumerFulfillments(databricksClient),
	}, nil
}

type ConsumerInstallationsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ConsumerInstallations ConsumerInstallationsInterface
}

func NewConsumerInstallationsClient(cfg *config.Config) (*ConsumerInstallationsClient, error) {
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

	return &ConsumerInstallationsClient{
		cfg:                   cfg,
		apiClient:             apiClient,
		ConsumerInstallations: NewConsumerInstallations(databricksClient),
	}, nil
}

type ConsumerListingsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ConsumerListings ConsumerListingsInterface
}

func NewConsumerListingsClient(cfg *config.Config) (*ConsumerListingsClient, error) {
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

	return &ConsumerListingsClient{
		cfg:              cfg,
		apiClient:        apiClient,
		ConsumerListings: NewConsumerListings(databricksClient),
	}, nil
}

type ConsumerPersonalizationRequestsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ConsumerPersonalizationRequests ConsumerPersonalizationRequestsInterface
}

func NewConsumerPersonalizationRequestsClient(cfg *config.Config) (*ConsumerPersonalizationRequestsClient, error) {
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

	return &ConsumerPersonalizationRequestsClient{
		cfg:                             cfg,
		apiClient:                       apiClient,
		ConsumerPersonalizationRequests: NewConsumerPersonalizationRequests(databricksClient),
	}, nil
}

type ConsumerProvidersClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ConsumerProviders ConsumerProvidersInterface
}

func NewConsumerProvidersClient(cfg *config.Config) (*ConsumerProvidersClient, error) {
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

	return &ConsumerProvidersClient{
		cfg:               cfg,
		apiClient:         apiClient,
		ConsumerProviders: NewConsumerProviders(databricksClient),
	}, nil
}

type ProviderExchangeFiltersClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ProviderExchangeFilters ProviderExchangeFiltersInterface
}

func NewProviderExchangeFiltersClient(cfg *config.Config) (*ProviderExchangeFiltersClient, error) {
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

	return &ProviderExchangeFiltersClient{
		cfg:                     cfg,
		apiClient:               apiClient,
		ProviderExchangeFilters: NewProviderExchangeFilters(databricksClient),
	}, nil
}

type ProviderExchangesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ProviderExchanges ProviderExchangesInterface
}

func NewProviderExchangesClient(cfg *config.Config) (*ProviderExchangesClient, error) {
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

	return &ProviderExchangesClient{
		cfg:               cfg,
		apiClient:         apiClient,
		ProviderExchanges: NewProviderExchanges(databricksClient),
	}, nil
}

type ProviderFilesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ProviderFiles ProviderFilesInterface
}

func NewProviderFilesClient(cfg *config.Config) (*ProviderFilesClient, error) {
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

	return &ProviderFilesClient{
		cfg:           cfg,
		apiClient:     apiClient,
		ProviderFiles: NewProviderFiles(databricksClient),
	}, nil
}

type ProviderListingsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ProviderListings ProviderListingsInterface
}

func NewProviderListingsClient(cfg *config.Config) (*ProviderListingsClient, error) {
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

	return &ProviderListingsClient{
		cfg:              cfg,
		apiClient:        apiClient,
		ProviderListings: NewProviderListings(databricksClient),
	}, nil
}

type ProviderPersonalizationRequestsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ProviderPersonalizationRequests ProviderPersonalizationRequestsInterface
}

func NewProviderPersonalizationRequestsClient(cfg *config.Config) (*ProviderPersonalizationRequestsClient, error) {
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

	return &ProviderPersonalizationRequestsClient{
		cfg:                             cfg,
		apiClient:                       apiClient,
		ProviderPersonalizationRequests: NewProviderPersonalizationRequests(databricksClient),
	}, nil
}

type ProviderProviderAnalyticsDashboardsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ProviderProviderAnalyticsDashboards ProviderProviderAnalyticsDashboardsInterface
}

func NewProviderProviderAnalyticsDashboardsClient(cfg *config.Config) (*ProviderProviderAnalyticsDashboardsClient, error) {
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

	return &ProviderProviderAnalyticsDashboardsClient{
		cfg:                                 cfg,
		apiClient:                           apiClient,
		ProviderProviderAnalyticsDashboards: NewProviderProviderAnalyticsDashboards(databricksClient),
	}, nil
}

type ProviderProvidersClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ProviderProviders ProviderProvidersInterface
}

func NewProviderProvidersClient(cfg *config.Config) (*ProviderProvidersClient, error) {
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

	return &ProviderProvidersClient{
		cfg:               cfg,
		apiClient:         apiClient,
		ProviderProviders: NewProviderProviders(databricksClient),
	}, nil
}
