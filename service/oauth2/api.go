// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Custom App Integration, O Auth Published Apps, Published App Integration, Service Principal Secrets, etc.
package oauth2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type CustomAppIntegrationInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockCustomAppIntegrationInterface instead.
	WithImpl(impl CustomAppIntegrationService) CustomAppIntegrationInterface

	// Impl returns low-level CustomAppIntegration API implementation
	// Deprecated: use MockCustomAppIntegrationInterface instead.
	Impl() CustomAppIntegrationService

	// Create Custom OAuth App Integration.
	//
	// Create Custom OAuth App Integration.
	//
	// You can retrieve the custom OAuth app integration via
	// :method:CustomAppIntegration/get.
	Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error)

	// Delete Custom OAuth App Integration.
	//
	// Delete an existing Custom OAuth App Integration. You can retrieve the custom
	// OAuth app integration via :method:CustomAppIntegration/get.
	Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error

	// Delete Custom OAuth App Integration.
	//
	// Delete an existing Custom OAuth App Integration. You can retrieve the custom
	// OAuth app integration via :method:CustomAppIntegration/get.
	DeleteByIntegrationId(ctx context.Context, integrationId string) error

	// Get OAuth Custom App Integration.
	//
	// Gets the Custom OAuth App Integration for the given integration id.
	Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error)

	// Get OAuth Custom App Integration.
	//
	// Gets the Custom OAuth App Integration for the given integration id.
	GetByIntegrationId(ctx context.Context, integrationId string) (*GetCustomAppIntegrationOutput, error)

	// Get custom oauth app integrations.
	//
	// Get the list of custom OAuth app integrations for the specified Databricks
	// account
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListCustomAppIntegrationsRequest) listing.Iterator[GetCustomAppIntegrationOutput]

	// Get custom oauth app integrations.
	//
	// Get the list of custom OAuth app integrations for the specified Databricks
	// account
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListCustomAppIntegrationsRequest) ([]GetCustomAppIntegrationOutput, error)

	// Updates Custom OAuth App Integration.
	//
	// Updates an existing custom OAuth App Integration. You can retrieve the custom
	// OAuth app integration via :method:CustomAppIntegration/get.
	Update(ctx context.Context, request UpdateCustomAppIntegration) error
}

func NewCustomAppIntegration(client *client.DatabricksClient) *CustomAppIntegrationAPI {
	return &CustomAppIntegrationAPI{
		CustomAppIntegrationService: &customAppIntegrationImpl{
			client: client,
		},
	}
}

// These APIs enable administrators to manage custom OAuth app integrations,
// which is required for adding/using Custom OAuth App Integration like Tableau
// Cloud for Databricks in AWS cloud.
type CustomAppIntegrationAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CustomAppIntegrationService)
	CustomAppIntegrationService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockCustomAppIntegrationInterface instead.
func (a *CustomAppIntegrationAPI) WithImpl(impl CustomAppIntegrationService) CustomAppIntegrationInterface {
	return &CustomAppIntegrationAPI{
		CustomAppIntegrationService: impl,
	}
}

// Impl returns low-level CustomAppIntegration API implementation
// Deprecated: use MockCustomAppIntegrationInterface instead.
func (a *CustomAppIntegrationAPI) Impl() CustomAppIntegrationService {
	return a.CustomAppIntegrationService
}

// Delete Custom OAuth App Integration.
//
// Delete an existing Custom OAuth App Integration. You can retrieve the custom
// OAuth app integration via :method:CustomAppIntegration/get.
func (a *CustomAppIntegrationAPI) DeleteByIntegrationId(ctx context.Context, integrationId string) error {
	return a.CustomAppIntegrationService.Delete(ctx, DeleteCustomAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get OAuth Custom App Integration.
//
// Gets the Custom OAuth App Integration for the given integration id.
func (a *CustomAppIntegrationAPI) GetByIntegrationId(ctx context.Context, integrationId string) (*GetCustomAppIntegrationOutput, error) {
	return a.CustomAppIntegrationService.Get(ctx, GetCustomAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get custom oauth app integrations.
//
// Get the list of custom OAuth app integrations for the specified Databricks
// account
//
// This method is generated by Databricks SDK Code Generator.
func (a *CustomAppIntegrationAPI) List(ctx context.Context, request ListCustomAppIntegrationsRequest) listing.Iterator[GetCustomAppIntegrationOutput] {

	getNextPage := func(ctx context.Context, req ListCustomAppIntegrationsRequest) (*GetCustomAppIntegrationsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.CustomAppIntegrationService.List(ctx, req)
	}
	getItems := func(resp *GetCustomAppIntegrationsOutput) []GetCustomAppIntegrationOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetCustomAppIntegrationsOutput) *ListCustomAppIntegrationsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Get custom oauth app integrations.
//
// Get the list of custom OAuth app integrations for the specified Databricks
// account
//
// This method is generated by Databricks SDK Code Generator.
func (a *CustomAppIntegrationAPI) ListAll(ctx context.Context, request ListCustomAppIntegrationsRequest) ([]GetCustomAppIntegrationOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[GetCustomAppIntegrationOutput](ctx, iterator)
}

type OAuthPublishedAppsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockOAuthPublishedAppsInterface instead.
	WithImpl(impl OAuthPublishedAppsService) OAuthPublishedAppsInterface

	// Impl returns low-level OAuthPublishedApps API implementation
	// Deprecated: use MockOAuthPublishedAppsInterface instead.
	Impl() OAuthPublishedAppsService

	// Get all the published OAuth apps.
	//
	// Get all the available published OAuth apps in Databricks.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListOAuthPublishedAppsRequest) listing.Iterator[PublishedAppOutput]

	// Get all the published OAuth apps.
	//
	// Get all the available published OAuth apps in Databricks.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListOAuthPublishedAppsRequest) ([]PublishedAppOutput, error)
}

func NewOAuthPublishedApps(client *client.DatabricksClient) *OAuthPublishedAppsAPI {
	return &OAuthPublishedAppsAPI{
		OAuthPublishedAppsService: &oAuthPublishedAppsImpl{
			client: client,
		},
	}
}

// These APIs enable administrators to view all the available published OAuth
// applications in Databricks. Administrators can add the published OAuth
// applications to their account through the OAuth Published App Integration
// APIs.
type OAuthPublishedAppsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(OAuthPublishedAppsService)
	OAuthPublishedAppsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockOAuthPublishedAppsInterface instead.
func (a *OAuthPublishedAppsAPI) WithImpl(impl OAuthPublishedAppsService) OAuthPublishedAppsInterface {
	return &OAuthPublishedAppsAPI{
		OAuthPublishedAppsService: impl,
	}
}

// Impl returns low-level OAuthPublishedApps API implementation
// Deprecated: use MockOAuthPublishedAppsInterface instead.
func (a *OAuthPublishedAppsAPI) Impl() OAuthPublishedAppsService {
	return a.OAuthPublishedAppsService
}

// Get all the published OAuth apps.
//
// Get all the available published OAuth apps in Databricks.
//
// This method is generated by Databricks SDK Code Generator.
func (a *OAuthPublishedAppsAPI) List(ctx context.Context, request ListOAuthPublishedAppsRequest) listing.Iterator[PublishedAppOutput] {

	getNextPage := func(ctx context.Context, req ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.OAuthPublishedAppsService.List(ctx, req)
	}
	getItems := func(resp *GetPublishedAppsOutput) []PublishedAppOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetPublishedAppsOutput) *ListOAuthPublishedAppsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Get all the published OAuth apps.
//
// Get all the available published OAuth apps in Databricks.
//
// This method is generated by Databricks SDK Code Generator.
func (a *OAuthPublishedAppsAPI) ListAll(ctx context.Context, request ListOAuthPublishedAppsRequest) ([]PublishedAppOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PublishedAppOutput](ctx, iterator)
}

type PublishedAppIntegrationInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockPublishedAppIntegrationInterface instead.
	WithImpl(impl PublishedAppIntegrationService) PublishedAppIntegrationInterface

	// Impl returns low-level PublishedAppIntegration API implementation
	// Deprecated: use MockPublishedAppIntegrationInterface instead.
	Impl() PublishedAppIntegrationService

	// Create Published OAuth App Integration.
	//
	// Create Published OAuth App Integration.
	//
	// You can retrieve the published OAuth app integration via
	// :method:PublishedAppIntegration/get.
	Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error)

	// Delete Published OAuth App Integration.
	//
	// Delete an existing Published OAuth App Integration. You can retrieve the
	// published OAuth app integration via :method:PublishedAppIntegration/get.
	Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error

	// Delete Published OAuth App Integration.
	//
	// Delete an existing Published OAuth App Integration. You can retrieve the
	// published OAuth app integration via :method:PublishedAppIntegration/get.
	DeleteByIntegrationId(ctx context.Context, integrationId string) error

	// Get OAuth Published App Integration.
	//
	// Gets the Published OAuth App Integration for the given integration id.
	Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error)

	// Get OAuth Published App Integration.
	//
	// Gets the Published OAuth App Integration for the given integration id.
	GetByIntegrationId(ctx context.Context, integrationId string) (*GetPublishedAppIntegrationOutput, error)

	// Get published oauth app integrations.
	//
	// Get the list of published OAuth app integrations for the specified Databricks
	// account
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListPublishedAppIntegrationsRequest) listing.Iterator[GetPublishedAppIntegrationOutput]

	// Get published oauth app integrations.
	//
	// Get the list of published OAuth app integrations for the specified Databricks
	// account
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListPublishedAppIntegrationsRequest) ([]GetPublishedAppIntegrationOutput, error)

	// Updates Published OAuth App Integration.
	//
	// Updates an existing published OAuth App Integration. You can retrieve the
	// published OAuth app integration via :method:PublishedAppIntegration/get.
	Update(ctx context.Context, request UpdatePublishedAppIntegration) error
}

func NewPublishedAppIntegration(client *client.DatabricksClient) *PublishedAppIntegrationAPI {
	return &PublishedAppIntegrationAPI{
		PublishedAppIntegrationService: &publishedAppIntegrationImpl{
			client: client,
		},
	}
}

// These APIs enable administrators to manage published OAuth app integrations,
// which is required for adding/using Published OAuth App Integration like
// Tableau Desktop for Databricks in AWS cloud.
type PublishedAppIntegrationAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PublishedAppIntegrationService)
	PublishedAppIntegrationService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockPublishedAppIntegrationInterface instead.
func (a *PublishedAppIntegrationAPI) WithImpl(impl PublishedAppIntegrationService) PublishedAppIntegrationInterface {
	return &PublishedAppIntegrationAPI{
		PublishedAppIntegrationService: impl,
	}
}

// Impl returns low-level PublishedAppIntegration API implementation
// Deprecated: use MockPublishedAppIntegrationInterface instead.
func (a *PublishedAppIntegrationAPI) Impl() PublishedAppIntegrationService {
	return a.PublishedAppIntegrationService
}

// Delete Published OAuth App Integration.
//
// Delete an existing Published OAuth App Integration. You can retrieve the
// published OAuth app integration via :method:PublishedAppIntegration/get.
func (a *PublishedAppIntegrationAPI) DeleteByIntegrationId(ctx context.Context, integrationId string) error {
	return a.PublishedAppIntegrationService.Delete(ctx, DeletePublishedAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get OAuth Published App Integration.
//
// Gets the Published OAuth App Integration for the given integration id.
func (a *PublishedAppIntegrationAPI) GetByIntegrationId(ctx context.Context, integrationId string) (*GetPublishedAppIntegrationOutput, error) {
	return a.PublishedAppIntegrationService.Get(ctx, GetPublishedAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get published oauth app integrations.
//
// Get the list of published OAuth app integrations for the specified Databricks
// account
//
// This method is generated by Databricks SDK Code Generator.
func (a *PublishedAppIntegrationAPI) List(ctx context.Context, request ListPublishedAppIntegrationsRequest) listing.Iterator[GetPublishedAppIntegrationOutput] {

	getNextPage := func(ctx context.Context, req ListPublishedAppIntegrationsRequest) (*GetPublishedAppIntegrationsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.PublishedAppIntegrationService.List(ctx, req)
	}
	getItems := func(resp *GetPublishedAppIntegrationsOutput) []GetPublishedAppIntegrationOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetPublishedAppIntegrationsOutput) *ListPublishedAppIntegrationsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Get published oauth app integrations.
//
// Get the list of published OAuth app integrations for the specified Databricks
// account
//
// This method is generated by Databricks SDK Code Generator.
func (a *PublishedAppIntegrationAPI) ListAll(ctx context.Context, request ListPublishedAppIntegrationsRequest) ([]GetPublishedAppIntegrationOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[GetPublishedAppIntegrationOutput](ctx, iterator)
}

type ServicePrincipalSecretsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockServicePrincipalSecretsInterface instead.
	WithImpl(impl ServicePrincipalSecretsService) ServicePrincipalSecretsInterface

	// Impl returns low-level ServicePrincipalSecrets API implementation
	// Deprecated: use MockServicePrincipalSecretsInterface instead.
	Impl() ServicePrincipalSecretsService

	// Create service principal secret.
	//
	// Create a secret for the given service principal.
	Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error)

	// Delete service principal secret.
	//
	// Delete a secret from the given service principal.
	Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error

	// Delete service principal secret.
	//
	// Delete a secret from the given service principal.
	DeleteByServicePrincipalIdAndSecretId(ctx context.Context, servicePrincipalId int64, secretId string) error

	// List service principal secrets.
	//
	// List all secrets associated with the given service principal. This operation
	// only returns information about the secrets themselves and does not include
	// the secret values.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListServicePrincipalSecretsRequest) listing.Iterator[SecretInfo]

	// List service principal secrets.
	//
	// List all secrets associated with the given service principal. This operation
	// only returns information about the secrets themselves and does not include
	// the secret values.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListServicePrincipalSecretsRequest) ([]SecretInfo, error)

	// List service principal secrets.
	//
	// List all secrets associated with the given service principal. This operation
	// only returns information about the secrets themselves and does not include
	// the secret values.
	ListByServicePrincipalId(ctx context.Context, servicePrincipalId int64) (*ListServicePrincipalSecretsResponse, error)
}

func NewServicePrincipalSecrets(client *client.DatabricksClient) *ServicePrincipalSecretsAPI {
	return &ServicePrincipalSecretsAPI{
		ServicePrincipalSecretsService: &servicePrincipalSecretsImpl{
			client: client,
		},
	}
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
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ServicePrincipalSecretsService)
	ServicePrincipalSecretsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockServicePrincipalSecretsInterface instead.
func (a *ServicePrincipalSecretsAPI) WithImpl(impl ServicePrincipalSecretsService) ServicePrincipalSecretsInterface {
	return &ServicePrincipalSecretsAPI{
		ServicePrincipalSecretsService: impl,
	}
}

// Impl returns low-level ServicePrincipalSecrets API implementation
// Deprecated: use MockServicePrincipalSecretsInterface instead.
func (a *ServicePrincipalSecretsAPI) Impl() ServicePrincipalSecretsService {
	return a.ServicePrincipalSecretsService
}

// Delete service principal secret.
//
// Delete a secret from the given service principal.
func (a *ServicePrincipalSecretsAPI) DeleteByServicePrincipalIdAndSecretId(ctx context.Context, servicePrincipalId int64, secretId string) error {
	return a.ServicePrincipalSecretsService.Delete(ctx, DeleteServicePrincipalSecretRequest{
		ServicePrincipalId: servicePrincipalId,
		SecretId:           secretId,
	})
}

// List service principal secrets.
//
// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalSecretsAPI) List(ctx context.Context, request ListServicePrincipalSecretsRequest) listing.Iterator[SecretInfo] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.ServicePrincipalSecretsService.List(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalSecretsResponse) []SecretInfo {
		return resp.Secrets
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List service principal secrets.
//
// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalSecretsAPI) ListAll(ctx context.Context, request ListServicePrincipalSecretsRequest) ([]SecretInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SecretInfo](ctx, iterator)
}

// List service principal secrets.
//
// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
func (a *ServicePrincipalSecretsAPI) ListByServicePrincipalId(ctx context.Context, servicePrincipalId int64) (*ListServicePrincipalSecretsResponse, error) {
	return a.ServicePrincipalSecretsService.List(ctx, ListServicePrincipalSecretsRequest{
		ServicePrincipalId: servicePrincipalId,
	})
}
