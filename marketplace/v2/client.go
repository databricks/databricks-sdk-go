// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ConsumerFulfillmentsClient struct {
	ConsumerFulfillmentsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                           cfg,
		apiClient:                     apiClient,
		ConsumerFulfillmentsInterface: NewConsumerFulfillments(databricksClient),
	}, nil
}

type ConsumerInstallationsClient struct {
	ConsumerInstallationsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                            cfg,
		apiClient:                      apiClient,
		ConsumerInstallationsInterface: NewConsumerInstallations(databricksClient),
	}, nil
}

type ConsumerListingsClient struct {
	ConsumerListingsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		ConsumerListingsInterface: NewConsumerListings(databricksClient),
	}, nil
}

type ConsumerPersonalizationRequestsClient struct {
	ConsumerPersonalizationRequestsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                                      cfg,
		apiClient:                                apiClient,
		ConsumerPersonalizationRequestsInterface: NewConsumerPersonalizationRequests(databricksClient),
	}, nil
}

type ConsumerProvidersClient struct {
	ConsumerProvidersInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                        cfg,
		apiClient:                  apiClient,
		ConsumerProvidersInterface: NewConsumerProviders(databricksClient),
	}, nil
}

type ProviderExchangeFiltersClient struct {
	ProviderExchangeFiltersInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                              cfg,
		apiClient:                        apiClient,
		ProviderExchangeFiltersInterface: NewProviderExchangeFilters(databricksClient),
	}, nil
}

type ProviderExchangesClient struct {
	ProviderExchangesInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                        cfg,
		apiClient:                  apiClient,
		ProviderExchangesInterface: NewProviderExchanges(databricksClient),
	}, nil
}

type ProviderFilesClient struct {
	ProviderFilesInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                    cfg,
		apiClient:              apiClient,
		ProviderFilesInterface: NewProviderFiles(databricksClient),
	}, nil
}

type ProviderListingsClient struct {
	ProviderListingsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		ProviderListingsInterface: NewProviderListings(databricksClient),
	}, nil
}

type ProviderPersonalizationRequestsClient struct {
	ProviderPersonalizationRequestsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                                      cfg,
		apiClient:                                apiClient,
		ProviderPersonalizationRequestsInterface: NewProviderPersonalizationRequests(databricksClient),
	}, nil
}

type ProviderProviderAnalyticsDashboardsClient struct {
	ProviderProviderAnalyticsDashboardsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:       cfg,
		apiClient: apiClient,
		ProviderProviderAnalyticsDashboardsInterface: NewProviderProviderAnalyticsDashboards(databricksClient),
	}, nil
}

type ProviderProvidersClient struct {
	ProviderProvidersInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                        cfg,
		apiClient:                  apiClient,
		ProviderProvidersInterface: NewProviderProviders(databricksClient),
	}, nil
}
