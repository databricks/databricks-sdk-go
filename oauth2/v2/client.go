// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccountFederationPolicyClient struct {
	AccountFederationPolicyInterface
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
		AccountFederationPolicyInterface: NewAccountFederationPolicy(apiClient),
	}, nil
}

type CustomAppIntegrationClient struct {
	CustomAppIntegrationInterface
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
		CustomAppIntegrationInterface: NewCustomAppIntegration(apiClient),
	}, nil
}

type OAuthPublishedAppsClient struct {
	OAuthPublishedAppsInterface
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
		OAuthPublishedAppsInterface: NewOAuthPublishedApps(apiClient),
	}, nil
}

type PublishedAppIntegrationClient struct {
	PublishedAppIntegrationInterface
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
		PublishedAppIntegrationInterface: NewPublishedAppIntegration(apiClient),
	}, nil
}

type ServicePrincipalFederationPolicyClient struct {
	ServicePrincipalFederationPolicyInterface
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
		ServicePrincipalFederationPolicyInterface: NewServicePrincipalFederationPolicy(apiClient),
	}, nil
}

type ServicePrincipalSecretsClient struct {
	ServicePrincipalSecretsInterface
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
		ServicePrincipalSecretsInterface: NewServicePrincipalSecrets(apiClient),
	}, nil
}
