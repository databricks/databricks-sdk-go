// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// Fulfillments are entities that allow consumers to preview installations.
type ConsumerFulfillmentsClient struct {
	consumerFulfillmentsBaseClient
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
		consumerFulfillmentsBaseClient: consumerFulfillmentsBaseClient{
			consumerFulfillmentsImpl: consumerFulfillmentsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Installations are entities that allow consumers to interact with Databricks
// Marketplace listings.
type ConsumerInstallationsClient struct {
	consumerInstallationsBaseClient
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
		consumerInstallationsBaseClient: consumerInstallationsBaseClient{
			consumerInstallationsImpl: consumerInstallationsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Listings are the core entities in the Marketplace. They represent the
// products that are available for consumption.
type ConsumerListingsClient struct {
	consumerListingsBaseClient
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
		consumerListingsBaseClient: consumerListingsBaseClient{
			consumerListingsImpl: consumerListingsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Personalization Requests allow customers to interact with the individualized
// Marketplace listing flow.
type ConsumerPersonalizationRequestsClient struct {
	consumerPersonalizationRequestsBaseClient
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
		consumerPersonalizationRequestsBaseClient: consumerPersonalizationRequestsBaseClient{
			consumerPersonalizationRequestsImpl: consumerPersonalizationRequestsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Providers are the entities that publish listings to the Marketplace.
type ConsumerProvidersClient struct {
	consumerProvidersBaseClient
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
		consumerProvidersBaseClient: consumerProvidersBaseClient{
			consumerProvidersImpl: consumerProvidersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Marketplace exchanges filters curate which groups can access an exchange.
type ProviderExchangeFiltersClient struct {
	providerExchangeFiltersBaseClient
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
		providerExchangeFiltersBaseClient: providerExchangeFiltersBaseClient{
			providerExchangeFiltersImpl: providerExchangeFiltersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Marketplace exchanges allow providers to share their listings with a curated
// set of customers.
type ProviderExchangesClient struct {
	providerExchangesBaseClient
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
		providerExchangesBaseClient: providerExchangesBaseClient{
			providerExchangesImpl: providerExchangesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Marketplace offers a set of file APIs for various purposes such as preview
// notebooks and provider icons.
type ProviderFilesClient struct {
	providerFilesBaseClient
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
		providerFilesBaseClient: providerFilesBaseClient{
			providerFilesImpl: providerFilesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Listings are the core entities in the Marketplace. They represent the
// products that are available for consumption.
type ProviderListingsClient struct {
	providerListingsBaseClient
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
		providerListingsBaseClient: providerListingsBaseClient{
			providerListingsImpl: providerListingsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Personalization requests are an alternate to instantly available listings.
// Control the lifecycle of personalized solutions.
type ProviderPersonalizationRequestsClient struct {
	providerPersonalizationRequestsBaseClient
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
		providerPersonalizationRequestsBaseClient: providerPersonalizationRequestsBaseClient{
			providerPersonalizationRequestsImpl: providerPersonalizationRequestsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Manage templated analytics solution for providers.
type ProviderProviderAnalyticsDashboardsClient struct {
	providerProviderAnalyticsDashboardsBaseClient
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
		providerProviderAnalyticsDashboardsBaseClient: providerProviderAnalyticsDashboardsBaseClient{
			providerProviderAnalyticsDashboardsImpl: providerProviderAnalyticsDashboardsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Providers are entities that manage assets in Marketplace.
type ProviderProvidersClient struct {
	providerProvidersBaseClient
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
		providerProvidersBaseClient: providerProvidersBaseClient{
			providerProvidersImpl: providerProvidersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
