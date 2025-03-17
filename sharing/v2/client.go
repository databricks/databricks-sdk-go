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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProvidersClient{
		apiClient:          databricksClient.ApiClient(),
		ProvidersInterface: NewProviders(databricksClient),
	}, nil
}

type RecipientActivationClient struct {
	RecipientActivationInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RecipientActivationClient{
		apiClient:                    databricksClient.ApiClient(),
		RecipientActivationInterface: NewRecipientActivation(databricksClient),
	}, nil
}

type RecipientsClient struct {
	RecipientsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RecipientsClient{
		apiClient:           databricksClient.ApiClient(),
		RecipientsInterface: NewRecipients(databricksClient),
	}, nil
}

type SharesClient struct {
	SharesInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &SharesClient{
		apiClient:       databricksClient.ApiClient(),
		SharesInterface: NewShares(databricksClient),
	}, nil
}
