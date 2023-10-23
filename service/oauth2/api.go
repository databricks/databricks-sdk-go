// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Custom App Integration, O Auth Enrollment, O Auth Published Apps, Published App Integration, Service Principal Secrets, etc.
package oauth2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewCustomAppIntegration(client *client.DatabricksClient) *CustomAppIntegrationAPI {
	return &CustomAppIntegrationAPI{
		impl: &customAppIntegrationImpl{
			client: client,
		},
	}
}

// These APIs enable administrators to manage custom oauth app integrations,
// which is required for adding/using Custom OAuth App Integration like Tableau
// Cloud for Databricks in AWS cloud.
type CustomAppIntegrationAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CustomAppIntegrationService)
	impl CustomAppIntegrationService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *CustomAppIntegrationAPI) WithImpl(impl CustomAppIntegrationService) *CustomAppIntegrationAPI {
	a.impl = impl
	return a
}

// Impl returns low-level CustomAppIntegration API implementation
func (a *CustomAppIntegrationAPI) Impl() CustomAppIntegrationService {
	return a.impl
}

// Create Custom OAuth App Integration.
//
// Create Custom OAuth App Integration.
//
// You can retrieve the custom oauth app integration via
// :method:CustomAppIntegration/get.
func (a *CustomAppIntegrationAPI) Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error) {
	return a.impl.Create(ctx, request)
}

// Delete Custom OAuth App Integration.
//
// Delete an existing Custom OAuth App Integration. You can retrieve the custom
// oauth app integration via :method:CustomAppIntegration/get.
func (a *CustomAppIntegrationAPI) Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete Custom OAuth App Integration.
//
// Delete an existing Custom OAuth App Integration. You can retrieve the custom
// oauth app integration via :method:CustomAppIntegration/get.
func (a *CustomAppIntegrationAPI) DeleteByIntegrationId(ctx context.Context, integrationId string) error {
	return a.impl.Delete(ctx, DeleteCustomAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get OAuth Custom App Integration.
//
// Gets the Custom OAuth App Integration for the given integration id.
func (a *CustomAppIntegrationAPI) Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error) {
	return a.impl.Get(ctx, request)
}

// Get OAuth Custom App Integration.
//
// Gets the Custom OAuth App Integration for the given integration id.
func (a *CustomAppIntegrationAPI) GetByIntegrationId(ctx context.Context, integrationId string) (*GetCustomAppIntegrationOutput, error) {
	return a.impl.Get(ctx, GetCustomAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get custom oauth app integrations.
//
// Get the list of custom oauth app integrations for the specified Databricks
// account
//
// This method is generated by Databricks SDK Code Generator.
func (a *CustomAppIntegrationAPI) List(ctx context.Context) *listing.PaginatingIterator[struct{}, *GetCustomAppIntegrationsOutput, GetCustomAppIntegrationOutput] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*GetCustomAppIntegrationsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx)
	}
	getItems := func(resp *GetCustomAppIntegrationsOutput) []GetCustomAppIntegrationOutput {
		return resp.Apps
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get custom oauth app integrations.
//
// Get the list of custom oauth app integrations for the specified Databricks
// account
//
// This method is generated by Databricks SDK Code Generator.
func (a *CustomAppIntegrationAPI) ListAll(ctx context.Context) ([]GetCustomAppIntegrationOutput, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[GetCustomAppIntegrationOutput](ctx, iterator)
}

// Updates Custom OAuth App Integration.
//
// Updates an existing custom OAuth App Integration. You can retrieve the custom
// oauth app integration via :method:CustomAppIntegration/get.
func (a *CustomAppIntegrationAPI) Update(ctx context.Context, request UpdateCustomAppIntegration) error {
	return a.impl.Update(ctx, request)
}

func NewOAuthEnrollment(client *client.DatabricksClient) *OAuthEnrollmentAPI {
	return &OAuthEnrollmentAPI{
		impl: &oAuthEnrollmentImpl{
			client: client,
		},
	}
}

// These APIs enable administrators to enroll OAuth for their accounts, which is
// required for adding/using any OAuth published/custom application integration.
//
// **Note:** Your account must be on the E2 version to use these APIs, this is
// because OAuth is only supported on the E2 version.
type OAuthEnrollmentAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(OAuthEnrollmentService)
	impl OAuthEnrollmentService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *OAuthEnrollmentAPI) WithImpl(impl OAuthEnrollmentService) *OAuthEnrollmentAPI {
	a.impl = impl
	return a
}

// Impl returns low-level OAuthEnrollment API implementation
func (a *OAuthEnrollmentAPI) Impl() OAuthEnrollmentService {
	return a.impl
}

// Create OAuth Enrollment request.
//
// Create an OAuth Enrollment request to enroll OAuth for this account and
// optionally enable the OAuth integration for all the partner applications in
// the account.
//
// The parter applications are: - Power BI - Tableau Desktop - Databricks CLI
//
// The enrollment is executed asynchronously, so the API will return 204
// immediately. The actual enrollment take a few minutes, you can check the
// status via API :method:OAuthEnrollment/get.
func (a *OAuthEnrollmentAPI) Create(ctx context.Context, request CreateOAuthEnrollment) error {
	return a.impl.Create(ctx, request)
}

// Get OAuth enrollment status.
//
// Gets the OAuth enrollment status for this Account.
//
// You can only add/use the OAuth published/custom application integrations when
// OAuth enrollment status is enabled.
func (a *OAuthEnrollmentAPI) Get(ctx context.Context) (*OAuthEnrollmentStatus, error) {
	return a.impl.Get(ctx)
}

func NewOAuthPublishedApps(client *client.DatabricksClient) *OAuthPublishedAppsAPI {
	return &OAuthPublishedAppsAPI{
		impl: &oAuthPublishedAppsImpl{
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
	impl OAuthPublishedAppsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *OAuthPublishedAppsAPI) WithImpl(impl OAuthPublishedAppsService) *OAuthPublishedAppsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level OAuthPublishedApps API implementation
func (a *OAuthPublishedAppsAPI) Impl() OAuthPublishedAppsService {
	return a.impl
}

// Get all the published OAuth apps.
//
// Get all the available published OAuth apps in Databricks.
//
// This method is generated by Databricks SDK Code Generator.
func (a *OAuthPublishedAppsAPI) List(ctx context.Context, request ListOAuthPublishedAppsRequest) *listing.PaginatingIterator[ListOAuthPublishedAppsRequest, *GetPublishedAppsOutput, PublishedAppOutput] {

	getNextPage := func(ctx context.Context, req ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx, req)
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
	return listing.ToSliceN[PublishedAppOutput, int64](ctx, iterator, request.PageSize)

}

func NewPublishedAppIntegration(client *client.DatabricksClient) *PublishedAppIntegrationAPI {
	return &PublishedAppIntegrationAPI{
		impl: &publishedAppIntegrationImpl{
			client: client,
		},
	}
}

// These APIs enable administrators to manage published oauth app integrations,
// which is required for adding/using Published OAuth App Integration like
// Tableau Desktop for Databricks in AWS cloud.
type PublishedAppIntegrationAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PublishedAppIntegrationService)
	impl PublishedAppIntegrationService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PublishedAppIntegrationAPI) WithImpl(impl PublishedAppIntegrationService) *PublishedAppIntegrationAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PublishedAppIntegration API implementation
func (a *PublishedAppIntegrationAPI) Impl() PublishedAppIntegrationService {
	return a.impl
}

// Create Published OAuth App Integration.
//
// Create Published OAuth App Integration.
//
// You can retrieve the published oauth app integration via
// :method:PublishedAppIntegration/get.
func (a *PublishedAppIntegrationAPI) Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error) {
	return a.impl.Create(ctx, request)
}

// Delete Published OAuth App Integration.
//
// Delete an existing Published OAuth App Integration. You can retrieve the
// published oauth app integration via :method:PublishedAppIntegration/get.
func (a *PublishedAppIntegrationAPI) Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete Published OAuth App Integration.
//
// Delete an existing Published OAuth App Integration. You can retrieve the
// published oauth app integration via :method:PublishedAppIntegration/get.
func (a *PublishedAppIntegrationAPI) DeleteByIntegrationId(ctx context.Context, integrationId string) error {
	return a.impl.Delete(ctx, DeletePublishedAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get OAuth Published App Integration.
//
// Gets the Published OAuth App Integration for the given integration id.
func (a *PublishedAppIntegrationAPI) Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error) {
	return a.impl.Get(ctx, request)
}

// Get OAuth Published App Integration.
//
// Gets the Published OAuth App Integration for the given integration id.
func (a *PublishedAppIntegrationAPI) GetByIntegrationId(ctx context.Context, integrationId string) (*GetPublishedAppIntegrationOutput, error) {
	return a.impl.Get(ctx, GetPublishedAppIntegrationRequest{
		IntegrationId: integrationId,
	})
}

// Get published oauth app integrations.
//
// Get the list of published oauth app integrations for the specified Databricks
// account
//
// This method is generated by Databricks SDK Code Generator.
func (a *PublishedAppIntegrationAPI) List(ctx context.Context) *listing.PaginatingIterator[struct{}, *GetPublishedAppIntegrationsOutput, GetPublishedAppIntegrationOutput] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*GetPublishedAppIntegrationsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx)
	}
	getItems := func(resp *GetPublishedAppIntegrationsOutput) []GetPublishedAppIntegrationOutput {
		return resp.Apps
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get published oauth app integrations.
//
// Get the list of published oauth app integrations for the specified Databricks
// account
//
// This method is generated by Databricks SDK Code Generator.
func (a *PublishedAppIntegrationAPI) ListAll(ctx context.Context) ([]GetPublishedAppIntegrationOutput, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[GetPublishedAppIntegrationOutput](ctx, iterator)
}

// Updates Published OAuth App Integration.
//
// Updates an existing published OAuth App Integration. You can retrieve the
// published oauth app integration via :method:PublishedAppIntegration/get.
func (a *PublishedAppIntegrationAPI) Update(ctx context.Context, request UpdatePublishedAppIntegration) error {
	return a.impl.Update(ctx, request)
}

func NewServicePrincipalSecrets(client *client.DatabricksClient) *ServicePrincipalSecretsAPI {
	return &ServicePrincipalSecretsAPI{
		impl: &servicePrincipalSecretsImpl{
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
	impl ServicePrincipalSecretsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ServicePrincipalSecretsAPI) WithImpl(impl ServicePrincipalSecretsService) *ServicePrincipalSecretsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level ServicePrincipalSecrets API implementation
func (a *ServicePrincipalSecretsAPI) Impl() ServicePrincipalSecretsService {
	return a.impl
}

// Create service principal secret.
//
// Create a secret for the given service principal.
func (a *ServicePrincipalSecretsAPI) Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete service principal secret.
//
// Delete a secret from the given service principal.
func (a *ServicePrincipalSecretsAPI) Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete service principal secret.
//
// Delete a secret from the given service principal.
func (a *ServicePrincipalSecretsAPI) DeleteByServicePrincipalIdAndSecretId(ctx context.Context, servicePrincipalId int64, secretId string) error {
	return a.impl.Delete(ctx, DeleteServicePrincipalSecretRequest{
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
func (a *ServicePrincipalSecretsAPI) List(ctx context.Context, request ListServicePrincipalSecretsRequest) *listing.PaginatingIterator[ListServicePrincipalSecretsRequest, *ListServicePrincipalSecretsResponse, SecretInfo] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx, req)
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
	return a.impl.List(ctx, ListServicePrincipalSecretsRequest{
		ServicePrincipalId: servicePrincipalId,
	})
}
