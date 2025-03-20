// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccountFederationPolicyClient struct {
	AccountFederationPolicyAPI
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
		AccountFederationPolicyAPI: AccountFederationPolicyAPI{
			accountFederationPolicyImpl: accountFederationPolicyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type CustomAppIntegrationClient struct {
	CustomAppIntegrationAPI
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
		CustomAppIntegrationAPI: CustomAppIntegrationAPI{
			customAppIntegrationImpl: customAppIntegrationImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type OAuthPublishedAppsClient struct {
	OAuthPublishedAppsAPI
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
		OAuthPublishedAppsAPI: OAuthPublishedAppsAPI{
			oAuthPublishedAppsImpl: oAuthPublishedAppsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type PublishedAppIntegrationClient struct {
	PublishedAppIntegrationAPI
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
		PublishedAppIntegrationAPI: PublishedAppIntegrationAPI{
			publishedAppIntegrationImpl: publishedAppIntegrationImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ServicePrincipalFederationPolicyClient struct {
	ServicePrincipalFederationPolicyAPI
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
		ServicePrincipalFederationPolicyAPI: ServicePrincipalFederationPolicyAPI{
			servicePrincipalFederationPolicyImpl: servicePrincipalFederationPolicyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ServicePrincipalSecretsClient struct {
	ServicePrincipalSecretsAPI
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
		ServicePrincipalSecretsAPI: ServicePrincipalSecretsAPI{
			servicePrincipalSecretsImpl: servicePrincipalSecretsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
