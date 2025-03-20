// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type ProvidersClient struct {
	ProvidersAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ProvidersClient{
		ProvidersAPI: ProvidersAPI{
			providersImpl: providersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type RecipientActivationClient struct {
	RecipientActivationAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RecipientActivationClient{
		RecipientActivationAPI: RecipientActivationAPI{
			recipientActivationImpl: recipientActivationImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type RecipientsClient struct {
	RecipientsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RecipientsClient{
		RecipientsAPI: RecipientsAPI{
			recipientsImpl: recipientsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type SharesClient struct {
	SharesAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &SharesClient{
		SharesAPI: SharesAPI{
			sharesImpl: sharesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
