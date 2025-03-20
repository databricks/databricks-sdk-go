// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type ConsumerFulfillmentsClient struct {
	ConsumerFulfillmentsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerFulfillmentsClient{
		ConsumerFulfillmentsInterface: NewConsumerFulfillments(apiClient),
	}, nil
}

type ConsumerInstallationsClient struct {
	ConsumerInstallationsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerInstallationsClient{
		ConsumerInstallationsInterface: NewConsumerInstallations(apiClient),
	}, nil
}

type ConsumerListingsClient struct {
	ConsumerListingsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerListingsClient{
		ConsumerListingsInterface: NewConsumerListings(apiClient),
	}, nil
}

type ConsumerPersonalizationRequestsClient struct {
	ConsumerPersonalizationRequestsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerPersonalizationRequestsClient{
		ConsumerPersonalizationRequestsInterface: NewConsumerPersonalizationRequests(apiClient),
	}, nil
}

type ConsumerProvidersClient struct {
	ConsumerProvidersInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsumerProvidersClient{
		ConsumerProvidersInterface: NewConsumerProviders(apiClient),
	}, nil
}

type ProviderExchangeFiltersClient struct {
	ProviderExchangeFiltersInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderExchangeFiltersClient{
		ProviderExchangeFiltersInterface: NewProviderExchangeFilters(apiClient),
	}, nil
}

type ProviderExchangesClient struct {
	ProviderExchangesInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderExchangesClient{
		ProviderExchangesInterface: NewProviderExchanges(apiClient),
	}, nil
}

type ProviderFilesClient struct {
	ProviderFilesInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderFilesClient{
		ProviderFilesInterface: NewProviderFiles(apiClient),
	}, nil
}

type ProviderListingsClient struct {
	ProviderListingsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderListingsClient{
		ProviderListingsInterface: NewProviderListings(apiClient),
	}, nil
}

type ProviderPersonalizationRequestsClient struct {
	ProviderPersonalizationRequestsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderPersonalizationRequestsClient{
		ProviderPersonalizationRequestsInterface: NewProviderPersonalizationRequests(apiClient),
	}, nil
}

type ProviderProviderAnalyticsDashboardsClient struct {
	ProviderProviderAnalyticsDashboardsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderProviderAnalyticsDashboardsClient{
		ProviderProviderAnalyticsDashboardsInterface: NewProviderProviderAnalyticsDashboards(apiClient),
	}, nil
}

type ProviderProvidersClient struct {
	ProviderProvidersInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProviderProvidersClient{
		ProviderProvidersInterface: NewProviderProviders(apiClient),
	}, nil
}
