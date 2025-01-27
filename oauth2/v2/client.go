// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccountFederationPolicyClient struct {
	AccountFederationPolicyInterface

	cfg *config.Config
}

func NewAccountFederationPolicyClient(cfg *config.Config) (*AccountFederationPolicyClient, error) {
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

	return &AccountFederationPolicyClient{
		cfg:                              cfg,
		AccountFederationPolicyInterface: NewAccountFederationPolicy(apiClient),
	}, nil
}

type CustomAppIntegrationClient struct {
	CustomAppIntegrationInterface

	cfg *config.Config
}

func NewCustomAppIntegrationClient(cfg *config.Config) (*CustomAppIntegrationClient, error) {
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

	return &CustomAppIntegrationClient{
		cfg:                           cfg,
		CustomAppIntegrationInterface: NewCustomAppIntegration(apiClient),
	}, nil
}

type OAuthPublishedAppsClient struct {
	OAuthPublishedAppsInterface

	cfg *config.Config
}

func NewOAuthPublishedAppsClient(cfg *config.Config) (*OAuthPublishedAppsClient, error) {
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

	return &OAuthPublishedAppsClient{
		cfg:                         cfg,
		OAuthPublishedAppsInterface: NewOAuthPublishedApps(apiClient),
	}, nil
}

type PublishedAppIntegrationClient struct {
	PublishedAppIntegrationInterface

	cfg *config.Config
}

func NewPublishedAppIntegrationClient(cfg *config.Config) (*PublishedAppIntegrationClient, error) {
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

	return &PublishedAppIntegrationClient{
		cfg:                              cfg,
		PublishedAppIntegrationInterface: NewPublishedAppIntegration(apiClient),
	}, nil
}

type ServicePrincipalFederationPolicyClient struct {
	ServicePrincipalFederationPolicyInterface

	cfg *config.Config
}

func NewServicePrincipalFederationPolicyClient(cfg *config.Config) (*ServicePrincipalFederationPolicyClient, error) {
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

	return &ServicePrincipalFederationPolicyClient{
		cfg: cfg,
		ServicePrincipalFederationPolicyInterface: NewServicePrincipalFederationPolicy(apiClient),
	}, nil
}

type ServicePrincipalSecretsClient struct {
	ServicePrincipalSecretsInterface

	cfg *config.Config
}

func NewServicePrincipalSecretsClient(cfg *config.Config) (*ServicePrincipalSecretsClient, error) {
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

	return &ServicePrincipalSecretsClient{
		cfg:                              cfg,
		ServicePrincipalSecretsInterface: NewServicePrincipalSecrets(apiClient),
	}, nil
}
