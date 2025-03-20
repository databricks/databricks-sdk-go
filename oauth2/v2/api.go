// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Federation Policy, Custom App Integration, O Auth Published Apps, Published App Integration, Service Principal Federation Policy, Service Principal Secrets, etc.
package oauth2

import (
	"context"
)

type accountFederationPolicyBaseClient struct {
	accountFederationPolicyImpl
}

// Delete account federation policy.
func (a *accountFederationPolicyBaseClient) DeleteByPolicyId(ctx context.Context, policyId string) (*DeleteResponse, error) {
	return a.accountFederationPolicyImpl.Delete(ctx, DeleteAccountFederationPolicyRequest{
		PolicyId: policyId,
	})
}

// Get account federation policy.
func (a *accountFederationPolicyBaseClient) GetByPolicyId(ctx context.Context, policyId string) (*FederationPolicy, error) {
	return a.accountFederationPolicyImpl.Get(ctx, GetAccountFederationPolicyRequest{
		PolicyId: policyId,
	})
}

type customAppIntegrationBaseClient struct {
	customAppIntegrationImpl
}

// Delete Custom OAuth App Integration.
//
// Delete an existing Custom OAuth App Integration. You can retrieve the custom
// OAuth app integration via :method:CustomAppIntegration/get.
func (a *customAppIntegrationBaseClient) DeleteByIntegrationId(ctx context.Context, integrationId string) (*DeleteCustomAppIntegrationOutput, error) {
	return a.customAppIntegrationImpl.Delete(ctx, DeleteCustomAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get OAuth Custom App Integration.
//
// Gets the Custom OAuth App Integration for the given integration id.
func (a *customAppIntegrationBaseClient) GetByIntegrationId(ctx context.Context, integrationId string) (*GetCustomAppIntegrationOutput, error) {
	return a.customAppIntegrationImpl.Get(ctx, GetCustomAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

type oAuthPublishedAppsBaseClient struct {
	oAuthPublishedAppsImpl
}

type publishedAppIntegrationBaseClient struct {
	publishedAppIntegrationImpl
}

// Delete Published OAuth App Integration.
//
// Delete an existing Published OAuth App Integration. You can retrieve the
// published OAuth app integration via :method:PublishedAppIntegration/get.
func (a *publishedAppIntegrationBaseClient) DeleteByIntegrationId(ctx context.Context, integrationId string) (*DeletePublishedAppIntegrationOutput, error) {
	return a.publishedAppIntegrationImpl.Delete(ctx, DeletePublishedAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get OAuth Published App Integration.
//
// Gets the Published OAuth App Integration for the given integration id.
func (a *publishedAppIntegrationBaseClient) GetByIntegrationId(ctx context.Context, integrationId string) (*GetPublishedAppIntegrationOutput, error) {
	return a.publishedAppIntegrationImpl.Get(ctx, GetPublishedAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

type servicePrincipalFederationPolicyBaseClient struct {
	servicePrincipalFederationPolicyImpl
}

// Delete service principal federation policy.
func (a *servicePrincipalFederationPolicyBaseClient) DeleteByServicePrincipalIdAndPolicyId(ctx context.Context, servicePrincipalId int64, policyId string) (*DeleteResponse, error) {
	return a.servicePrincipalFederationPolicyImpl.Delete(ctx, DeleteServicePrincipalFederationPolicyRequest{
		ServicePrincipalId: servicePrincipalId,
		PolicyId:           policyId,
	})
}

// Get service principal federation policy.
func (a *servicePrincipalFederationPolicyBaseClient) GetByServicePrincipalIdAndPolicyId(ctx context.Context, servicePrincipalId int64, policyId string) (*FederationPolicy, error) {
	return a.servicePrincipalFederationPolicyImpl.Get(ctx, GetServicePrincipalFederationPolicyRequest{
		ServicePrincipalId: servicePrincipalId,
		PolicyId:           policyId,
	})
}

// List service principal federation policies.
func (a *servicePrincipalFederationPolicyBaseClient) ListByServicePrincipalId(ctx context.Context, servicePrincipalId int64) (*ListFederationPoliciesResponse, error) {
	return a.servicePrincipalFederationPolicyImpl.internalList(ctx, ListServicePrincipalFederationPoliciesRequest{
		ServicePrincipalId: servicePrincipalId,
	})
}

type servicePrincipalSecretsBaseClient struct {
	servicePrincipalSecretsImpl
}

// Delete service principal secret.
//
// Delete a secret from the given service principal.
func (a *servicePrincipalSecretsBaseClient) DeleteByServicePrincipalIdAndSecretId(ctx context.Context, servicePrincipalId int64, secretId string) (*DeleteResponse, error) {
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
func (a *servicePrincipalSecretsBaseClient) ListByServicePrincipalId(ctx context.Context, servicePrincipalId int64) (*ListServicePrincipalSecretsResponse, error) {
	return a.servicePrincipalSecretsImpl.internalList(ctx, ListServicePrincipalSecretsRequest{
		ServicePrincipalId: servicePrincipalId,
	})
}
