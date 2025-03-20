// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Federation Policy, Custom App Integration, O Auth Published Apps, Published App Integration, Service Principal Federation Policy, Service Principal Secrets, etc.
package oauth2

import (
	"context"
)

type AccountFederationPolicyAPI struct {
	accountFederationPolicyImpl
}

// Delete account federation policy.
func (a *AccountFederationPolicyAPI) DeleteByPolicyId(ctx context.Context, policyId string) (*DeleteResponse, error) {
	return a.accountFederationPolicyImpl.Delete(ctx, DeleteAccountFederationPolicyRequest{
		PolicyId: policyId,
	})
}

// Get account federation policy.
func (a *AccountFederationPolicyAPI) GetByPolicyId(ctx context.Context, policyId string) (*FederationPolicy, error) {
	return a.accountFederationPolicyImpl.Get(ctx, GetAccountFederationPolicyRequest{
		PolicyId: policyId,
	})
}

type CustomAppIntegrationAPI struct {
	customAppIntegrationImpl
}

// Delete Custom OAuth App Integration.
//
// Delete an existing Custom OAuth App Integration. You can retrieve the custom
// OAuth app integration via :method:CustomAppIntegration/get.
func (a *CustomAppIntegrationAPI) DeleteByIntegrationId(ctx context.Context, integrationId string) (*DeleteCustomAppIntegrationOutput, error) {
	return a.customAppIntegrationImpl.Delete(ctx, DeleteCustomAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get OAuth Custom App Integration.
//
// Gets the Custom OAuth App Integration for the given integration id.
func (a *CustomAppIntegrationAPI) GetByIntegrationId(ctx context.Context, integrationId string) (*GetCustomAppIntegrationOutput, error) {
	return a.customAppIntegrationImpl.Get(ctx, GetCustomAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

type OAuthPublishedAppsAPI struct {
	oAuthPublishedAppsImpl
}

type PublishedAppIntegrationAPI struct {
	publishedAppIntegrationImpl
}

// Delete Published OAuth App Integration.
//
// Delete an existing Published OAuth App Integration. You can retrieve the
// published OAuth app integration via :method:PublishedAppIntegration/get.
func (a *PublishedAppIntegrationAPI) DeleteByIntegrationId(ctx context.Context, integrationId string) (*DeletePublishedAppIntegrationOutput, error) {
	return a.publishedAppIntegrationImpl.Delete(ctx, DeletePublishedAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get OAuth Published App Integration.
//
// Gets the Published OAuth App Integration for the given integration id.
func (a *PublishedAppIntegrationAPI) GetByIntegrationId(ctx context.Context, integrationId string) (*GetPublishedAppIntegrationOutput, error) {
	return a.publishedAppIntegrationImpl.Get(ctx, GetPublishedAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

type ServicePrincipalFederationPolicyAPI struct {
	servicePrincipalFederationPolicyImpl
}

// Delete service principal federation policy.
func (a *ServicePrincipalFederationPolicyAPI) DeleteByServicePrincipalIdAndPolicyId(ctx context.Context, servicePrincipalId int64, policyId string) (*DeleteResponse, error) {
	return a.servicePrincipalFederationPolicyImpl.Delete(ctx, DeleteServicePrincipalFederationPolicyRequest{
		ServicePrincipalId: servicePrincipalId,
		PolicyId:           policyId,
	})
}

// Get service principal federation policy.
func (a *ServicePrincipalFederationPolicyAPI) GetByServicePrincipalIdAndPolicyId(ctx context.Context, servicePrincipalId int64, policyId string) (*FederationPolicy, error) {
	return a.servicePrincipalFederationPolicyImpl.Get(ctx, GetServicePrincipalFederationPolicyRequest{
		ServicePrincipalId: servicePrincipalId,
		PolicyId:           policyId,
	})
}

// List service principal federation policies.
func (a *ServicePrincipalFederationPolicyAPI) ListByServicePrincipalId(ctx context.Context, servicePrincipalId int64) (*ListFederationPoliciesResponse, error) {
	return a.servicePrincipalFederationPolicyImpl.internalList(ctx, ListServicePrincipalFederationPoliciesRequest{
		ServicePrincipalId: servicePrincipalId,
	})
}

type ServicePrincipalSecretsAPI struct {
	servicePrincipalSecretsImpl
}

// Delete service principal secret.
//
// Delete a secret from the given service principal.
func (a *ServicePrincipalSecretsAPI) DeleteByServicePrincipalIdAndSecretId(ctx context.Context, servicePrincipalId int64, secretId string) (*DeleteResponse, error) {
	return a.servicePrincipalSecretsImpl.Delete(ctx, DeleteServicePrincipalSecretRequest{
		ServicePrincipalId: servicePrincipalId,
		SecretId:           secretId,
	})
}

// List service principal secrets.
//
// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
func (a *ServicePrincipalSecretsAPI) ListByServicePrincipalId(ctx context.Context, servicePrincipalId int64) (*ListServicePrincipalSecretsResponse, error) {
	return a.servicePrincipalSecretsImpl.internalList(ctx, ListServicePrincipalSecretsRequest{
		ServicePrincipalId: servicePrincipalId,
	})
}
