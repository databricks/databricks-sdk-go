// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type ProvidersClient struct {
	ProvidersInterface
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
		ProvidersInterface: NewProviders(apiClient),
	}, nil
}

type RecipientActivationClient struct {
	RecipientActivationInterface
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
		RecipientActivationInterface: NewRecipientActivation(apiClient),
	}, nil
}

type RecipientsClient struct {
	RecipientsInterface
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
		RecipientsInterface: NewRecipients(apiClient),
	}, nil
}

type SharesClient struct {
	SharesInterface
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
		SharesInterface: NewShares(apiClient),
	}, nil
}
