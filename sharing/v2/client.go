// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ProvidersClient struct {
	ProvidersInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
		Config:             cfg,
		apiClient:          apiClient,
		ProvidersInterface: NewProviders(databricksClient),
	}, nil
}

type RecipientActivationClient struct {
	RecipientActivationInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
		Config:                       cfg,
		apiClient:                    apiClient,
		RecipientActivationInterface: NewRecipientActivation(databricksClient),
	}, nil
}

type RecipientsClient struct {
	RecipientsInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
		Config:              cfg,
		apiClient:           apiClient,
		RecipientsInterface: NewRecipients(databricksClient),
	}, nil
}

type SharesClient struct {
	SharesInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
		Config:          cfg,
		apiClient:       apiClient,
		SharesInterface: NewShares(databricksClient),
	}, nil
}
