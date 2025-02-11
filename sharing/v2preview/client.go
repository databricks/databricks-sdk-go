// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharingpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ProvidersPreviewClient struct {
	ProvidersPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewProvidersPreviewClient(cfg *config.Config) (*ProvidersPreviewClient, error) {
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

	return &ProvidersPreviewClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		ProvidersPreviewInterface: NewProvidersPreview(databricksClient),
	}, nil
}

type RecipientActivationPreviewClient struct {
	RecipientActivationPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewRecipientActivationPreviewClient(cfg *config.Config) (*RecipientActivationPreviewClient, error) {
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

	return &RecipientActivationPreviewClient{
		Config:                              cfg,
		apiClient:                           apiClient,
		RecipientActivationPreviewInterface: NewRecipientActivationPreview(databricksClient),
	}, nil
}

type RecipientsPreviewClient struct {
	RecipientsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewRecipientsPreviewClient(cfg *config.Config) (*RecipientsPreviewClient, error) {
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

	return &RecipientsPreviewClient{
		Config:                     cfg,
		apiClient:                  apiClient,
		RecipientsPreviewInterface: NewRecipientsPreview(databricksClient),
	}, nil
}

type SharesPreviewClient struct {
	SharesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewSharesPreviewClient(cfg *config.Config) (*SharesPreviewClient, error) {
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

	return &SharesPreviewClient{
		Config:                 cfg,
		apiClient:              apiClient,
		SharesPreviewInterface: NewSharesPreview(databricksClient),
	}, nil
}
