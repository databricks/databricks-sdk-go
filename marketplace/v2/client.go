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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerFulfillmentsClient{
		apiClient:                     databricksClient.ApiClient(),
		ConsumerFulfillmentsInterface: NewConsumerFulfillments(databricksClient),
	}, nil
}

type ConsumerInstallationsClient struct {
	ConsumerInstallationsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerInstallationsClient{
		apiClient:                      databricksClient.ApiClient(),
		ConsumerInstallationsInterface: NewConsumerInstallations(databricksClient),
	}, nil
}

type ConsumerListingsClient struct {
	ConsumerListingsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerListingsClient{
		apiClient:                 databricksClient.ApiClient(),
		ConsumerListingsInterface: NewConsumerListings(databricksClient),
	}, nil
}

type ConsumerPersonalizationRequestsClient struct {
	ConsumerPersonalizationRequestsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerPersonalizationRequestsClient{
		apiClient:                                databricksClient.ApiClient(),
		ConsumerPersonalizationRequestsInterface: NewConsumerPersonalizationRequests(databricksClient),
	}, nil
}

type ConsumerProvidersClient struct {
	ConsumerProvidersInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerProvidersClient{
		apiClient:                  databricksClient.ApiClient(),
		ConsumerProvidersInterface: NewConsumerProviders(databricksClient),
	}, nil
}

type ProviderExchangeFiltersClient struct {
	ProviderExchangeFiltersInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderExchangeFiltersClient{
		apiClient:                        databricksClient.ApiClient(),
		ProviderExchangeFiltersInterface: NewProviderExchangeFilters(databricksClient),
	}, nil
}

type ProviderExchangesClient struct {
	ProviderExchangesInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderExchangesClient{
		apiClient:                  databricksClient.ApiClient(),
		ProviderExchangesInterface: NewProviderExchanges(databricksClient),
	}, nil
}

type ProviderFilesClient struct {
	ProviderFilesInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderFilesClient{
		apiClient:              databricksClient.ApiClient(),
		ProviderFilesInterface: NewProviderFiles(databricksClient),
	}, nil
}

type ProviderListingsClient struct {
	ProviderListingsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderListingsClient{
		apiClient:                 databricksClient.ApiClient(),
		ProviderListingsInterface: NewProviderListings(databricksClient),
	}, nil
}

type ProviderPersonalizationRequestsClient struct {
	ProviderPersonalizationRequestsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderPersonalizationRequestsClient{
		apiClient:                                databricksClient.ApiClient(),
		ProviderPersonalizationRequestsInterface: NewProviderPersonalizationRequests(databricksClient),
	}, nil
}

type ProviderProviderAnalyticsDashboardsClient struct {
	ProviderProviderAnalyticsDashboardsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderProviderAnalyticsDashboardsClient{
		apiClient: databricksClient.ApiClient(),
		ProviderProviderAnalyticsDashboardsInterface: NewProviderProviderAnalyticsDashboards(databricksClient),
	}, nil
}

type ProviderProvidersClient struct {
	ProviderProvidersInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderProvidersClient{
		apiClient:                  databricksClient.ApiClient(),
		ProviderProvidersInterface: NewProviderProviders(databricksClient),
	}, nil
}
