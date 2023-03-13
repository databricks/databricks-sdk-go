// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Custom App Integration, Published App Integration, etc.
package oauth2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
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
//
// **Note:** You can only add/use the OAuth custom application integrations when
// OAuth enrollment status is enabled. For more details see
// :method:OAuthEnrollment/create
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
// You can retrieve the custom oauth app integration via :method:get.
func (a *CustomAppIntegrationAPI) Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error) {
	return a.impl.Create(ctx, request)
}

// Delete Custom OAuth App Integration.
//
// Delete an existing Custom OAuth App Integration. You can retrieve the custom
// oauth app integration via :method:get.
func (a *CustomAppIntegrationAPI) Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete Custom OAuth App Integration.
//
// Delete an existing Custom OAuth App Integration. You can retrieve the custom
// oauth app integration via :method:get.
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

// Updates Custom OAuth App Integration.
//
// Updates an existing custom OAuth App Integration. You can retrieve the custom
// oauth app integration via :method:get.
func (a *CustomAppIntegrationAPI) Update(ctx context.Context, request UpdateCustomAppIntegration) error {
	return a.impl.Update(ctx, request)
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
// Tableau Cloud for Databricks in AWS cloud.
//
// **Note:** You can only add/use the OAuth published application integrations
// when OAuth enrollment status is enabled. For more details see
// :method:OAuthEnrollment/create
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
// You can retrieve the published oauth app integration via :method:get.
func (a *PublishedAppIntegrationAPI) Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error) {
	return a.impl.Create(ctx, request)
}

// Delete Published OAuth App Integration.
//
// Delete an existing Published OAuth App Integration. You can retrieve the
// published oauth app integration via :method:get.
func (a *PublishedAppIntegrationAPI) Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete Published OAuth App Integration.
//
// Delete an existing Published OAuth App Integration. You can retrieve the
// published oauth app integration via :method:get.
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

// Updates Published OAuth App Integration.
//
// Updates an existing published OAuth App Integration. You can retrieve the
// published oauth app integration via :method:get.
func (a *PublishedAppIntegrationAPI) Update(ctx context.Context, request UpdatePublishedAppIntegration) error {
	return a.impl.Update(ctx, request)
}
