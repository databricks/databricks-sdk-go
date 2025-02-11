// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2preview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccountFederationPolicyPreviewClient struct {
	AccountFederationPolicyPreviewInterface

	Config *config.Config
}

func NewAccountFederationPolicyPreviewClient(cfg *config.Config) (*AccountFederationPolicyPreviewClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AccountFederationPolicyPreviewClient{
		Config:                                  cfg,
		AccountFederationPolicyPreviewInterface: NewAccountFederationPolicyPreview(apiClient),
	}, nil
}

type CustomAppIntegrationPreviewClient struct {
	CustomAppIntegrationPreviewInterface

	Config *config.Config
}

func NewCustomAppIntegrationPreviewClient(cfg *config.Config) (*CustomAppIntegrationPreviewClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CustomAppIntegrationPreviewClient{
		Config:                               cfg,
		CustomAppIntegrationPreviewInterface: NewCustomAppIntegrationPreview(apiClient),
	}, nil
}

type OAuthPublishedAppsPreviewClient struct {
	OAuthPublishedAppsPreviewInterface

	Config *config.Config
}

func NewOAuthPublishedAppsPreviewClient(cfg *config.Config) (*OAuthPublishedAppsPreviewClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &OAuthPublishedAppsPreviewClient{
		Config:                             cfg,
		OAuthPublishedAppsPreviewInterface: NewOAuthPublishedAppsPreview(apiClient),
	}, nil
}

type PublishedAppIntegrationPreviewClient struct {
	PublishedAppIntegrationPreviewInterface

	Config *config.Config
}

func NewPublishedAppIntegrationPreviewClient(cfg *config.Config) (*PublishedAppIntegrationPreviewClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PublishedAppIntegrationPreviewClient{
		Config:                                  cfg,
		PublishedAppIntegrationPreviewInterface: NewPublishedAppIntegrationPreview(apiClient),
	}, nil
}

type ServicePrincipalFederationPolicyPreviewClient struct {
	ServicePrincipalFederationPolicyPreviewInterface

	Config *config.Config
}

func NewServicePrincipalFederationPolicyPreviewClient(cfg *config.Config) (*ServicePrincipalFederationPolicyPreviewClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipalFederationPolicyPreviewClient{
		Config: cfg,
		ServicePrincipalFederationPolicyPreviewInterface: NewServicePrincipalFederationPolicyPreview(apiClient),
	}, nil
}

type ServicePrincipalSecretsPreviewClient struct {
	ServicePrincipalSecretsPreviewInterface

	Config *config.Config
}

func NewServicePrincipalSecretsPreviewClient(cfg *config.Config) (*ServicePrincipalSecretsPreviewClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipalSecretsPreviewClient{
		Config:                                  cfg,
		ServicePrincipalSecretsPreviewInterface: NewServicePrincipalSecretsPreview(apiClient),
	}, nil
}
