// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
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
type AccountFederationPolicyClient struct {
	accountFederationPolicyBaseClient
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
		accountFederationPolicyBaseClient: accountFederationPolicyBaseClient{
			accountFederationPolicyImpl: accountFederationPolicyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs enable administrators to manage custom OAuth app integrations,
// which is required for adding/using Custom OAuth App Integration like Tableau
// Cloud for Databricks in AWS cloud.
type CustomAppIntegrationClient struct {
	customAppIntegrationBaseClient
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
		customAppIntegrationBaseClient: customAppIntegrationBaseClient{
			customAppIntegrationImpl: customAppIntegrationImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs enable administrators to view all the available published OAuth
// applications in Databricks. Administrators can add the published OAuth
// applications to their account through the OAuth Published App Integration
// APIs.
type OAuthPublishedAppsClient struct {
	oAuthPublishedAppsBaseClient
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
		oAuthPublishedAppsBaseClient: oAuthPublishedAppsBaseClient{
			oAuthPublishedAppsImpl: oAuthPublishedAppsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs enable administrators to manage published OAuth app integrations,
// which is required for adding/using Published OAuth App Integration like
// Tableau Desktop for Databricks in AWS cloud.
type PublishedAppIntegrationClient struct {
	publishedAppIntegrationBaseClient
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
		publishedAppIntegrationBaseClient: publishedAppIntegrationBaseClient{
			publishedAppIntegrationImpl: publishedAppIntegrationImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
type ServicePrincipalFederationPolicyClient struct {
	servicePrincipalFederationPolicyBaseClient
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
		servicePrincipalFederationPolicyBaseClient: servicePrincipalFederationPolicyBaseClient{
			servicePrincipalFederationPolicyImpl: servicePrincipalFederationPolicyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
type ServicePrincipalSecretsClient struct {
	servicePrincipalSecretsBaseClient
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
		servicePrincipalSecretsBaseClient: servicePrincipalSecretsBaseClient{
			servicePrincipalSecretsImpl: servicePrincipalSecretsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
