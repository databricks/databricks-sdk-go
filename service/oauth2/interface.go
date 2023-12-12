// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

// These APIs enable administrators to manage custom oauth app integrations,
// which is required for adding/using Custom OAuth App Integration like Tableau
// Cloud for Databricks in AWS cloud.
type CustomAppIntegrationService interface {

	// Create Custom OAuth App Integration.
	//
	// Create Custom OAuth App Integration.
	//
	// You can retrieve the custom oauth app integration via
	// :method:CustomAppIntegration/get.
	Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error)

	// Delete Custom OAuth App Integration.
	//
	// Delete an existing Custom OAuth App Integration. You can retrieve the
	// custom oauth app integration via :method:CustomAppIntegration/get.
	Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error

	// Get OAuth Custom App Integration.
	//
	// Gets the Custom OAuth App Integration for the given integration id.
	Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error)

	// Get custom oauth app integrations.
	//
	// Get the list of custom oauth app integrations for the specified
	// Databricks account
	//
	// Use ListAll() to get all GetCustomAppIntegrationOutput instances
	List(ctx context.Context) (*GetCustomAppIntegrationsOutput, error)

	// Updates Custom OAuth App Integration.
	//
	// Updates an existing custom OAuth App Integration. You can retrieve the
	// custom oauth app integration via :method:CustomAppIntegration/get.
	Update(ctx context.Context, request UpdateCustomAppIntegration) error
	// Returns an instance of DatabricksClient
	Client() client.DatabricksClientInterface
}

// These APIs enable administrators to view all the available published OAuth
// applications in Databricks. Administrators can add the published OAuth
// applications to their account through the OAuth Published App Integration
// APIs.
type OAuthPublishedAppsService interface {

	// Get all the published OAuth apps.
	//
	// Get all the available published OAuth apps in Databricks.
	//
	// Use ListAll() to get all PublishedAppOutput instances, which will iterate over every result page.
	List(ctx context.Context, request ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error)
	// Returns an instance of DatabricksClient
	Client() client.DatabricksClientInterface
}

// These APIs enable administrators to manage published oauth app integrations,
// which is required for adding/using Published OAuth App Integration like
// Tableau Desktop for Databricks in AWS cloud.
type PublishedAppIntegrationService interface {

	// Create Published OAuth App Integration.
	//
	// Create Published OAuth App Integration.
	//
	// You can retrieve the published oauth app integration via
	// :method:PublishedAppIntegration/get.
	Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error)

	// Delete Published OAuth App Integration.
	//
	// Delete an existing Published OAuth App Integration. You can retrieve the
	// published oauth app integration via :method:PublishedAppIntegration/get.
	Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error

	// Get OAuth Published App Integration.
	//
	// Gets the Published OAuth App Integration for the given integration id.
	Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error)

	// Get published oauth app integrations.
	//
	// Get the list of published oauth app integrations for the specified
	// Databricks account
	//
	// Use ListAll() to get all GetPublishedAppIntegrationOutput instances
	List(ctx context.Context) (*GetPublishedAppIntegrationsOutput, error)

	// Updates Published OAuth App Integration.
	//
	// Updates an existing published OAuth App Integration. You can retrieve the
	// published oauth app integration via :method:PublishedAppIntegration/get.
	Update(ctx context.Context, request UpdatePublishedAppIntegration) error
	// Returns an instance of DatabricksClient
	Client() client.DatabricksClientInterface
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
type ServicePrincipalSecretsService interface {

	// Create service principal secret.
	//
	// Create a secret for the given service principal.
	Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error)

	// Delete service principal secret.
	//
	// Delete a secret from the given service principal.
	Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error

	// List service principal secrets.
	//
	// List all secrets associated with the given service principal. This
	// operation only returns information about the secrets themselves and does
	// not include the secret values.
	//
	// Use ListAll() to get all SecretInfo instances
	List(ctx context.Context, request ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error)
	// Returns an instance of DatabricksClient
	Client() client.DatabricksClientInterface
}
