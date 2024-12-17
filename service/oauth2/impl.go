// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AccountFederationPolicy API methods
type accountFederationPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *accountFederationPolicyImpl) Create(ctx context.Context, request CreateAccountFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

func (a *accountFederationPolicyImpl) Delete(ctx context.Context, request DeleteAccountFederationPolicyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *accountFederationPolicyImpl) Get(ctx context.Context, request GetAccountFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &federationPolicy)
	return &federationPolicy, err
}

func (a *accountFederationPolicyImpl) List(ctx context.Context, request ListAccountFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
	var listFederationPoliciesResponse ListFederationPoliciesResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listFederationPoliciesResponse)
	return &listFederationPoliciesResponse, err
}

func (a *accountFederationPolicyImpl) Update(ctx context.Context, request UpdateAccountFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

// unexported type that holds implementations of just CustomAppIntegration API methods
type customAppIntegrationImpl struct {
	client *client.DatabricksClient
}

func (a *customAppIntegrationImpl) Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error) {
	var createCustomAppIntegrationOutput CreateCustomAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createCustomAppIntegrationOutput)
	return &createCustomAppIntegrationOutput, err
}

func (a *customAppIntegrationImpl) Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error {
	var deleteCustomAppIntegrationOutput DeleteCustomAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteCustomAppIntegrationOutput)
	return err
}

func (a *customAppIntegrationImpl) Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error) {
	var getCustomAppIntegrationOutput GetCustomAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getCustomAppIntegrationOutput)
	return &getCustomAppIntegrationOutput, err
}

func (a *customAppIntegrationImpl) List(ctx context.Context, request ListCustomAppIntegrationsRequest) (*GetCustomAppIntegrationsOutput, error) {
	var getCustomAppIntegrationsOutput GetCustomAppIntegrationsOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getCustomAppIntegrationsOutput)
	return &getCustomAppIntegrationsOutput, err
}

func (a *customAppIntegrationImpl) Update(ctx context.Context, request UpdateCustomAppIntegration) error {
	var updateCustomAppIntegrationOutput UpdateCustomAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateCustomAppIntegrationOutput)
	return err
}

// unexported type that holds implementations of just OAuthPublishedApps API methods
type oAuthPublishedAppsImpl struct {
	client *client.DatabricksClient
}

func (a *oAuthPublishedAppsImpl) List(ctx context.Context, request ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error) {
	var getPublishedAppsOutput GetPublishedAppsOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-apps", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getPublishedAppsOutput)
	return &getPublishedAppsOutput, err
}

// unexported type that holds implementations of just PublishedAppIntegration API methods
type publishedAppIntegrationImpl struct {
	client *client.DatabricksClient
}

func (a *publishedAppIntegrationImpl) Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error) {
	var createPublishedAppIntegrationOutput CreatePublishedAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createPublishedAppIntegrationOutput)
	return &createPublishedAppIntegrationOutput, err
}

func (a *publishedAppIntegrationImpl) Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error {
	var deletePublishedAppIntegrationOutput DeletePublishedAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deletePublishedAppIntegrationOutput)
	return err
}

func (a *publishedAppIntegrationImpl) Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error) {
	var getPublishedAppIntegrationOutput GetPublishedAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getPublishedAppIntegrationOutput)
	return &getPublishedAppIntegrationOutput, err
}

func (a *publishedAppIntegrationImpl) List(ctx context.Context, request ListPublishedAppIntegrationsRequest) (*GetPublishedAppIntegrationsOutput, error) {
	var getPublishedAppIntegrationsOutput GetPublishedAppIntegrationsOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getPublishedAppIntegrationsOutput)
	return &getPublishedAppIntegrationsOutput, err
}

func (a *publishedAppIntegrationImpl) Update(ctx context.Context, request UpdatePublishedAppIntegration) error {
	var updatePublishedAppIntegrationOutput UpdatePublishedAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updatePublishedAppIntegrationOutput)
	return err
}

// unexported type that holds implementations of just ServicePrincipalFederationPolicy API methods
type servicePrincipalFederationPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalFederationPolicyImpl) Create(ctx context.Context, request CreateServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

func (a *servicePrincipalFederationPolicyImpl) Delete(ctx context.Context, request DeleteServicePrincipalFederationPolicyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId, request.PolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *servicePrincipalFederationPolicyImpl) Get(ctx context.Context, request GetServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId, request.PolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &federationPolicy)
	return &federationPolicy, err
}

func (a *servicePrincipalFederationPolicyImpl) List(ctx context.Context, request ListServicePrincipalFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
	var listFederationPoliciesResponse ListFederationPoliciesResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listFederationPoliciesResponse)
	return &listFederationPoliciesResponse, err
}

func (a *servicePrincipalFederationPolicyImpl) Update(ctx context.Context, request UpdateServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId, request.PolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

// unexported type that holds implementations of just ServicePrincipalSecrets API methods
type servicePrincipalSecretsImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalSecretsImpl) Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error) {
	var createServicePrincipalSecretResponse CreateServicePrincipalSecretResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &createServicePrincipalSecretResponse)
	return &createServicePrincipalSecretResponse, err
}

func (a *servicePrincipalSecretsImpl) Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId, request.SecretId)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *servicePrincipalSecretsImpl) List(ctx context.Context, request ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {
	var listServicePrincipalSecretsResponse ListServicePrincipalSecretsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listServicePrincipalSecretsResponse)
	return &listServicePrincipalSecretsResponse, err
}
