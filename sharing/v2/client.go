// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ProvidersClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Providers ProvidersInterface
}

func NewProvidersClient(cfg *config.Config) (*ProvidersClient, error) {
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

	return &ProvidersClient{
		cfg:       cfg,
		apiClient: apiClient,
		Providers: NewProviders(databricksClient),
	}, nil
}

type RecipientActivationClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	RecipientActivation RecipientActivationInterface
}

func NewRecipientActivationClient(cfg *config.Config) (*RecipientActivationClient, error) {
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

	return &RecipientActivationClient{
		cfg:                 cfg,
		apiClient:           apiClient,
		RecipientActivation: NewRecipientActivation(databricksClient),
	}, nil
}

type RecipientsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Recipients RecipientsInterface
}

func NewRecipientsClient(cfg *config.Config) (*RecipientsClient, error) {
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

	return &RecipientsClient{
		cfg:        cfg,
		apiClient:  apiClient,
		Recipients: NewRecipients(databricksClient),
	}, nil
}

type SharesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Shares SharesInterface
}

func NewSharesClient(cfg *config.Config) (*SharesClient, error) {
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

	return &SharesClient{
		cfg:       cfg,
		apiClient: apiClient,
		Shares:    NewShares(databricksClient),
	}, nil
}
