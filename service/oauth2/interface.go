// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"context"
)

// These APIs enable administrators to manage custom oauth app integrations,
// which is required for adding/using Custom OAuth App Integration like Tableau
// Cloud for Databricks in AWS cloud.
//
// **Note:** You can only add/use the OAuth custom application integrations when
// OAuth enrollment status is enabled. For more details see
// :method:OAuthEnrollment/create
type CustomAppIntegrationService interface {

	// Create Custom OAuth App Integration.
	//
	// Create Custom OAuth App Integration.
	//
	// You can retrieve the custom oauth app integration via :method:get.
	Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error)

	// Delete Custom OAuth App Integration.
	//
	// Delete an existing Custom OAuth App Integration. You can retrieve the
	// custom oauth app integration via :method:get.
	Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error

	// Get OAuth Custom App Integration.
	//
	// Gets the Custom OAuth App Integration for the given integration id.
	Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error)

	// Updates Custom OAuth App Integration.
	//
	// Updates an existing custom OAuth App Integration. You can retrieve the
	// custom oauth app integration via :method:get.
	Update(ctx context.Context, request UpdateCustomAppIntegration) error
}

// These APIs enable administrators to manage published oauth app integrations,
// which is required for adding/using Published OAuth App Integration like
// Tableau Cloud for Databricks in AWS cloud.
//
// **Note:** You can only add/use the OAuth published application integrations
// when OAuth enrollment status is enabled. For more details see
// :method:OAuthEnrollment/create
type PublishedAppIntegrationService interface {

	// Create Published OAuth App Integration.
	//
	// Create Published OAuth App Integration.
	//
	// You can retrieve the published oauth app integration via :method:get.
	Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error)

	// Delete Published OAuth App Integration.
	//
	// Delete an existing Published OAuth App Integration. You can retrieve the
	// published oauth app integration via :method:get.
	Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error

	// Get OAuth Published App Integration.
	//
	// Gets the Published OAuth App Integration for the given integration id.
	Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error)

	// Updates Published OAuth App Integration.
	//
	// Updates an existing published OAuth App Integration. You can retrieve the
	// published oauth app integration via :method:get.
	Update(ctx context.Context, request UpdatePublishedAppIntegration) error
}
