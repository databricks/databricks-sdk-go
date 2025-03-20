// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Federation Policy, Custom App Integration, O Auth Published Apps, Published App Integration, Service Principal Federation Policy, Service Principal Secrets, etc.
package oauth2

import (
	"context"
)

// These APIs manage account federation policies.
//
// Account federation policies allow users and service principals in your
// Databricks account to securely access Databricks APIs using tokens from your
// trusted identity providers (IdPs).
//
// With token federation, your users and service principals can exchange tokens
// from your IdP for Databricks OAuth tokens, which can be used to access
// Databricks APIs. Token federation eliminates the need to manage Databricks
// secrets, and allows you to centralize management of token issuance policies
// in your IdP. Databricks token federation is typically used in combination
// with [SCIM], so users in your IdP are synchronized into your Databricks
// account.
//
// Token federation is configured in your Databricks account using an account
// federation policy. An account federation policy specifies: * which IdP, or
// issuer, your Databricks account should accept tokens from * how to determine
// which Databricks user, or subject, a token is issued for
//
// To configure a federation policy, you provide the following: * The required
// token __issuer__, as specified in the “iss” claim of your tokens. The
// issuer is an https URL that identifies your IdP. * The allowed token
// __audiences__, as specified in the “aud” claim of your tokens. This
// identifier is intended to represent the recipient of the token. As long as
// the audience in the token matches at least one audience in the policy, the
// token is considered a match. If unspecified, the default value is your
// Databricks account id. * The __subject claim__, which indicates which token
// claim contains the Databricks username of the user the token was issued for.
// If unspecified, the default value is “sub”. * Optionally, the public keys
// used to validate the signature of your tokens, in JWKS format. If unspecified
// (recommended), Databricks automatically fetches the public keys from your
// issuer’s well known endpoint. Databricks strongly recommends relying on
// your issuer’s well known endpoint for discovering public keys.
//
// An example federation policy is: ``` issuer: "https://idp.mycompany.com/oidc"
// audiences: ["databricks"] subject_claim: "sub" ```
//
// An example JWT token body that matches this policy and could be used to
// authenticate to Databricks as user `username@mycompany.com` is: ``` { "iss":
// "https://idp.mycompany.com/oidc", "aud": "databricks", "sub":
// "username@mycompany.com" } ```
//
// You may also need to configure your IdP to generate tokens for your users to
// exchange with Databricks, if your users do not already have the ability to
// generate tokens that are compatible with your federation policy.
//
// You do not need to configure an OAuth application in Databricks to use token
// federation.
//
// [SCIM]: https://docs.databricks.com/admin/users-groups/scim/index.html
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

// These APIs enable administrators to manage custom OAuth app integrations,
// which is required for adding/using Custom OAuth App Integration like Tableau
// Cloud for Databricks in AWS cloud.
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

// These APIs enable administrators to view all the available published OAuth
// applications in Databricks. Administrators can add the published OAuth
// applications to their account through the OAuth Published App Integration
// APIs.
type OAuthPublishedAppsAPI struct {
	oAuthPublishedAppsImpl
}

// These APIs enable administrators to manage published OAuth app integrations,
// which is required for adding/using Published OAuth App Integration like
// Tableau Desktop for Databricks in AWS cloud.
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

// These APIs manage service principal federation policies.
//
// Service principal federation, also known as Workload Identity Federation,
// allows your automated workloads running outside of Databricks to securely
// access Databricks APIs without the need for Databricks secrets. With Workload
// Identity Federation, your application (or workload) authenticates to
// Databricks as a Databricks service principal, using tokens provided by the
// workload runtime.
//
// Databricks strongly recommends using Workload Identity Federation to
// authenticate to Databricks from automated workloads, over alternatives such
// as OAuth client secrets or Personal Access Tokens, whenever possible.
// Workload Identity Federation is supported by many popular services, including
// Github Actions, Azure DevOps, GitLab, Terraform Cloud, and Kubernetes
// clusters, among others.
//
// Workload identity federation is configured in your Databricks account using a
// service principal federation policy. A service principal federation policy
// specifies: * which IdP, or issuer, the service principal is allowed to
// authenticate from * which workload identity, or subject, is allowed to
// authenticate as the Databricks service principal
//
// To configure a federation policy, you provide the following: * The required
// token __issuer__, as specified in the “iss” claim of workload identity
// tokens. The issuer is an https URL that identifies the workload identity
// provider. * The required token __subject__, as specified in the “sub”
// claim of workload identity tokens. The subject uniquely identifies the
// workload in the workload runtime environment. * The allowed token
// __audiences__, as specified in the “aud” claim of workload identity
// tokens. The audience is intended to represent the recipient of the token. As
// long as the audience in the token matches at least one audience in the
// policy, the token is considered a match. If unspecified, the default value is
// your Databricks account id. * Optionally, the public keys used to validate
// the signature of the workload identity tokens, in JWKS format. If unspecified
// (recommended), Databricks automatically fetches the public keys from the
// issuer’s well known endpoint. Databricks strongly recommends relying on the
// issuer’s well known endpoint for discovering public keys.
//
// An example service principal federation policy, for a Github Actions
// workload, is: ``` issuer: "https://token.actions.githubusercontent.com"
// audiences: ["https://github.com/my-github-org"] subject:
// "repo:my-github-org/my-repo:environment:prod" ```
//
// An example JWT token body that matches this policy and could be used to
// authenticate to Databricks is: ``` { "iss":
// "https://token.actions.githubusercontent.com", "aud":
// "https://github.com/my-github-org", "sub":
// "repo:my-github-org/my-repo:environment:prod" } ```
//
// You may also need to configure the workload runtime to generate tokens for
// your workloads.
//
// You do not need to configure an OAuth application in Databricks to use token
// federation.
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

// These APIs enable administrators to manage service principal secrets.
//
// You can use the generated secrets to obtain OAuth access tokens for a service
// principal, which can then be used to access Databricks Accounts and Workspace
// APIs. For more information, see [Authentication using OAuth tokens for
// service principals],
//
// In addition, the generated secrets can be used to configure the Databricks
// Terraform Provider to authenticate with the service principal. For more
// information, see [Databricks Terraform Provider].
//
// [Authentication using OAuth tokens for service principals]: https://docs.databricks.com/dev-tools/authentication-oauth.html
// [Databricks Terraform Provider]: https://github.com/databricks/terraform-provider-databricks/blob/master/docs/index.md#authenticating-with-service-principal
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
