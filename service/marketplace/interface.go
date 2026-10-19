// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"context"
)

// Fulfillments are entities that allow consumers to preview installations.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ConsumerFulfillmentsService interface {

	// Get a high level preview of the metadata of listing installable content.
	Get(ctx context.Context, request GetListingContentMetadataRequest) (*GetListingContentMetadataResponse, error)

	// Get all listings fulfillments associated with a listing. A _fulfillment_
	// is a potential installation. Standard installations contain metadata
	// about the attached share or git repo. Only one of these fields will be
	// present. Personalized installations contain metadata about the attached
	// share or git repo, as well as the Delta Sharing recipient type.
	List(ctx context.Context, request ListFulfillmentsRequest) (*ListFulfillmentsResponse, error)
}

// Installations are entities that allow consumers to interact with Databricks
// Marketplace listings.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ConsumerInstallationsService interface {

	// Install payload associated with a Databricks Marketplace listing.
	Create(ctx context.Context, request CreateInstallationRequest) (*Installation, error)

	// Uninstall an installation associated with a Databricks Marketplace
	// listing.
	Delete(ctx context.Context, request DeleteInstallationRequest) error

	// List all installations across all listings.
	List(ctx context.Context, request ListAllInstallationsRequest) (*ListAllInstallationsResponse, error)

	// List all installations for a particular listing.
	ListListingInstallations(ctx context.Context, request ListInstallationsRequest) (*ListInstallationsResponse, error)

	// This is a update API that will update the part of the fields defined in
	// the installation table as well as interact with external services
	// according to the fields not included in the installation table 1. the
	// token will be rotate if the rotateToken flag is true 2. the token will be
	// forcibly rotate if the rotateToken flag is true and the tokenInfo field
	// is empty
	Update(ctx context.Context, request UpdateInstallationRequest) (*UpdateInstallationResponse, error)
}

// Listings are the core entities in the Marketplace. They represent the
// products that are available for consumption.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ConsumerListingsService interface {

	// Batch get a published listing in the Databricks Marketplace that the
	// consumer has access to.
	BatchGet(ctx context.Context, request BatchGetListingsRequest) (*BatchGetListingsResponse, error)

	// Get a published listing in the Databricks Marketplace that the consumer
	// has access to.
	Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error)

	// List all published listings in the Databricks Marketplace that the
	// consumer has access to.
	List(ctx context.Context, request ListListingsRequest) (*ListListingsResponse, error)

	// Search published listings in the Databricks Marketplace that the consumer
	// has access to. This query supports a variety of different search
	// parameters and performs fuzzy matching.
	Search(ctx context.Context, request SearchListingsRequest) (*SearchListingsResponse, error)
}

// Personalization Requests allow customers to interact with the individualized
// Marketplace listing flow.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ConsumerPersonalizationRequestsService interface {

	// Create a personalization request for a listing.
	Create(ctx context.Context, request CreatePersonalizationRequest) (*CreatePersonalizationRequestResponse, error)

	// Get the personalization request for a listing. Each consumer can make at
	// *most* one personalization request for a listing.
	Get(ctx context.Context, request GetPersonalizationRequestRequest) (*GetPersonalizationRequestResponse, error)

	// List personalization requests for a consumer across all listings.
	List(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error)
}

// Providers are the entities that publish listings to the Marketplace.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ConsumerProvidersService interface {

	// Batch get a provider in the Databricks Marketplace with at least one
	// visible listing.
	BatchGet(ctx context.Context, request BatchGetProvidersRequest) (*BatchGetProvidersResponse, error)

	// Get a provider in the Databricks Marketplace with at least one visible
	// listing.
	Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error)

	// List all providers in the Databricks Marketplace with at least one
	// visible listing.
	List(ctx context.Context, request ListConsumerProvidersRequest) (*ListProvidersResponse, error)
}

// Marketplace exchanges filters curate which groups can access an exchange.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ProviderExchangeFiltersService interface {

	// Add an exchange filter.
	Create(ctx context.Context, request CreateExchangeFilterRequest) (*CreateExchangeFilterResponse, error)

	// Delete an exchange filter
	Delete(ctx context.Context, request DeleteExchangeFilterRequest) error

	// List exchange filter
	List(ctx context.Context, request ListExchangeFiltersRequest) (*ListExchangeFiltersResponse, error)

	// Update an exchange filter.
	Update(ctx context.Context, request UpdateExchangeFilterRequest) (*UpdateExchangeFilterResponse, error)
}

// Marketplace exchanges allow providers to share their listings with a curated
// set of customers.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ProviderExchangesService interface {

	// Associate an exchange with a listing
	AddListingToExchange(ctx context.Context, request AddExchangeForListingRequest) (*AddExchangeForListingResponse, error)

	// Create an exchange
	Create(ctx context.Context, request CreateExchangeRequest) (*CreateExchangeResponse, error)

	// This removes a listing from marketplace.
	Delete(ctx context.Context, request DeleteExchangeRequest) error

	// Disassociate an exchange with a listing
	DeleteListingFromExchange(ctx context.Context, request RemoveExchangeForListingRequest) error

	// Get an exchange.
	Get(ctx context.Context, request GetExchangeRequest) (*GetExchangeResponse, error)

	// List exchanges visible to provider
	List(ctx context.Context, request ListExchangesRequest) (*ListExchangesResponse, error)

	// List exchanges associated with a listing
	ListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) (*ListExchangesForListingResponse, error)

	// List listings associated with an exchange
	ListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) (*ListListingsForExchangeResponse, error)

	// Update an exchange
	Update(ctx context.Context, request UpdateExchangeRequest) (*UpdateExchangeResponse, error)
}

// Marketplace offers a set of file APIs for various purposes such as preview
// notebooks and provider icons.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ProviderFilesService interface {

	// Create a file. Currently, only provider icons and attached notebooks are
	// supported.
	Create(ctx context.Context, request CreateFileRequest) (*CreateFileResponse, error)

	// Delete a file
	Delete(ctx context.Context, request DeleteFileRequest) error

	// Get a file
	Get(ctx context.Context, request GetFileRequest) (*GetFileResponse, error)

	// List files attached to a parent entity.
	List(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error)
}

// Listings are the core entities in the Marketplace. They represent the
// products that are available for consumption.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ProviderListingsService interface {

	// Create a new listing
	Create(ctx context.Context, request CreateListingRequest) (*CreateListingResponse, error)

	// Delete a listing
	Delete(ctx context.Context, request DeleteListingRequest) error

	// Get a listing
	Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error)

	// List listings owned by this provider
	List(ctx context.Context, request GetListingsRequest) (*GetListingsResponse, error)

	// Update a listing
	Update(ctx context.Context, request UpdateListingRequest) (*UpdateListingResponse, error)
}

// Personalization requests are an alternate to instantly available listings.
// Control the lifecycle of personalized solutions.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ProviderPersonalizationRequestsService interface {

	// List personalization requests to this provider. This will return all
	// personalization requests, regardless of which listing they are for.
	List(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error)

	// Update personalization request. This method only permits updating the
	// status of the request.
	Update(ctx context.Context, request UpdatePersonalizationRequestRequest) (*UpdatePersonalizationRequestResponse, error)
}

// Manage templated analytics solution for providers.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ProviderProviderAnalyticsDashboardsService interface {

	// Create provider analytics dashboard. Returns Marketplace specific `id`.
	// Not to be confused with the Lakeview dashboard id.
	Create(ctx context.Context) (*ProviderAnalyticsDashboard, error)

	// Get provider analytics dashboard.
	Get(ctx context.Context) (*ListProviderAnalyticsDashboardResponse, error)

	// Get latest version of provider analytics dashboard.
	GetLatestVersion(ctx context.Context) (*GetLatestVersionProviderAnalyticsDashboardResponse, error)

	// Update provider analytics dashboard.
	Update(ctx context.Context, request UpdateProviderAnalyticsDashboardRequest) (*UpdateProviderAnalyticsDashboardResponse, error)
}

// Providers are entities that manage assets in Marketplace.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ProviderProvidersService interface {

	// Create a provider
	Create(ctx context.Context, request CreateProviderRequest) (*CreateProviderResponse, error)

	// Delete provider
	Delete(ctx context.Context, request DeleteProviderRequest) error

	// Get provider profile
	Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error)

	// List provider profiles for account.
	List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error)

	// Update provider profile
	Update(ctx context.Context, request UpdateProviderRequest) (*UpdateProviderResponse, error)
}
