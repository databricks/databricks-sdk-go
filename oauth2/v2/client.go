// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccountFederationPolicyClient struct {
	cfg *config.Config

	AccountFederationPolicy AccountFederationPolicyInterface
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
		cfg:                     cfg,
		AccountFederationPolicy: NewAccountFederationPolicy(apiClient),
	}, nil
}

type CustomAppIntegrationClient struct {
	cfg *config.Config

	CustomAppIntegration CustomAppIntegrationInterface
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
		cfg:                  cfg,
		CustomAppIntegration: NewCustomAppIntegration(apiClient),
	}, nil
}

type OAuthPublishedAppsClient struct {
	cfg *config.Config

	OAuthPublishedApps OAuthPublishedAppsInterface
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
		cfg:                cfg,
		OAuthPublishedApps: NewOAuthPublishedApps(apiClient),
	}, nil
}

type PublishedAppIntegrationClient struct {
	cfg *config.Config

	PublishedAppIntegration PublishedAppIntegrationInterface
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
		cfg:                     cfg,
		PublishedAppIntegration: NewPublishedAppIntegration(apiClient),
	}, nil
}

type ServicePrincipalFederationPolicyClient struct {
	cfg *config.Config

	ServicePrincipalFederationPolicy ServicePrincipalFederationPolicyInterface
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
		cfg:                              cfg,
		ServicePrincipalFederationPolicy: NewServicePrincipalFederationPolicy(apiClient),
	}, nil
}

type ServicePrincipalSecretsClient struct {
	cfg *config.Config

	ServicePrincipalSecrets ServicePrincipalSecretsInterface
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
		cfg:                     cfg,
		ServicePrincipalSecrets: NewServicePrincipalSecrets(apiClient),
	}, nil
}
