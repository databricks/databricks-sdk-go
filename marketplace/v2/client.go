// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type ConsumerFulfillmentsClient struct {
	ConsumerFulfillmentsAPI
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
		ConsumerFulfillmentsAPI: ConsumerFulfillmentsAPI{
			consumerFulfillmentsImpl: consumerFulfillmentsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ConsumerInstallationsClient struct {
	ConsumerInstallationsAPI
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
		ConsumerInstallationsAPI: ConsumerInstallationsAPI{
			consumerInstallationsImpl: consumerInstallationsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ConsumerListingsClient struct {
	ConsumerListingsAPI
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
		ConsumerListingsAPI: ConsumerListingsAPI{
			consumerListingsImpl: consumerListingsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ConsumerPersonalizationRequestsClient struct {
	ConsumerPersonalizationRequestsAPI
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
		ConsumerPersonalizationRequestsAPI: ConsumerPersonalizationRequestsAPI{
			consumerPersonalizationRequestsImpl: consumerPersonalizationRequestsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ConsumerProvidersClient struct {
	ConsumerProvidersAPI
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
		ConsumerProvidersAPI: ConsumerProvidersAPI{
			consumerProvidersImpl: consumerProvidersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ProviderExchangeFiltersClient struct {
	ProviderExchangeFiltersAPI
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
		ProviderExchangeFiltersAPI: ProviderExchangeFiltersAPI{
			providerExchangeFiltersImpl: providerExchangeFiltersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ProviderExchangesClient struct {
	ProviderExchangesAPI
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
		ProviderExchangesAPI: ProviderExchangesAPI{
			providerExchangesImpl: providerExchangesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ProviderFilesClient struct {
	ProviderFilesAPI
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
		ProviderFilesAPI: ProviderFilesAPI{
			providerFilesImpl: providerFilesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ProviderListingsClient struct {
	ProviderListingsAPI
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
		ProviderListingsAPI: ProviderListingsAPI{
			providerListingsImpl: providerListingsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ProviderPersonalizationRequestsClient struct {
	ProviderPersonalizationRequestsAPI
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
		ProviderPersonalizationRequestsAPI: ProviderPersonalizationRequestsAPI{
			providerPersonalizationRequestsImpl: providerPersonalizationRequestsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ProviderProviderAnalyticsDashboardsClient struct {
	ProviderProviderAnalyticsDashboardsAPI
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
		ProviderProviderAnalyticsDashboardsAPI: ProviderProviderAnalyticsDashboardsAPI{
			providerProviderAnalyticsDashboardsImpl: providerProviderAnalyticsDashboardsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ProviderProvidersClient struct {
	ProviderProvidersAPI
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
		ProviderProvidersAPI: ProviderProvidersAPI{
			providerProvidersImpl: providerProvidersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
