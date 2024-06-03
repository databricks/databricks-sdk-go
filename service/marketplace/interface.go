// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"context"
)

// Fulfillments are entities that allow consumers to preview installations.
type ConsumerFulfillmentsService interface {

	// Get listing content metadata.
	//
	// Get a high level preview of the metadata of listing installable content.
	//
	// Use GetAll() to get all SharedDataObject instances, which will iterate over every result page.
	Get(ctx context.Context, request GetListingContentMetadataRequest) (*GetListingContentMetadataResponse, error)

	// List all listing fulfillments.
	//
	// Get all listings fulfillments associated with a listing. A _fulfillment_
	// is a potential installation. Standard installations contain metadata
	// about the attached share or git repo. Only one of these fields will be
	// present. Personalized installations contain metadata about the attached
	// share or git repo, as well as the Delta Sharing recipient type.
	//
	// Use ListAll() to get all ListingFulfillment instances, which will iterate over every result page.
	List(ctx context.Context, request ListFulfillmentsRequest) (*ListFulfillmentsResponse, error)
}

// Installations are entities that allow consumers to interact with Databricks
// Marketplace listings.
type ConsumerInstallationsService interface {

	// Install from a listing.
	//
	// Install payload associated with a Databricks Marketplace listing.
	Create(ctx context.Context, request CreateInstallationRequest) (*Installation, error)

	// Uninstall from a listing.
	//
	// Uninstall an installation associated with a Databricks Marketplace
	// listing.
	Delete(ctx context.Context, request DeleteInstallationRequest) error

	// List all installations.
	//
	// List all installations across all listings.
	//
	// Use ListAll() to get all InstallationDetail instances, which will iterate over every result page.
	List(ctx context.Context, request ListAllInstallationsRequest) (*ListAllInstallationsResponse, error)

	// List installations for a listing.
	//
	// List all installations for a particular listing.
	//
	// Use ListListingInstallationsAll() to get all InstallationDetail instances, which will iterate over every result page.
	ListListingInstallations(ctx context.Context, request ListInstallationsRequest) (*ListInstallationsResponse, error)

	// Update an installation.
	//
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
type ConsumerListingsService interface {

	// Get one batch of listings. One may specify up to 50 IDs per request.
	//
	// Batch get a published listing in the Databricks Marketplace that the
	// consumer has access to.
	BatchGet(ctx context.Context, request BatchGetListingsRequest) (*BatchGetListingsResponse, error)

	// Get listing.
	//
	// Get a published listing in the Databricks Marketplace that the consumer
	// has access to.
	Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error)

	// List listings.
	//
	// List all published listings in the Databricks Marketplace that the
	// consumer has access to.
	//
	// Use ListAll() to get all Listing instances, which will iterate over every result page.
	List(ctx context.Context, request ListListingsRequest) (*ListListingsResponse, error)

	// Search listings.
	//
	// Search published listings in the Databricks Marketplace that the consumer
	// has access to. This query supports a variety of different search
	// parameters and performs fuzzy matching.
	//
	// Use SearchAll() to get all Listing instances, which will iterate over every result page.
	Search(ctx context.Context, request SearchListingsRequest) (*SearchListingsResponse, error)
}

// Personalization Requests allow customers to interact with the individualized
// Marketplace listing flow.
type ConsumerPersonalizationRequestsService interface {

	// Create a personalization request.
	//
	// Create a personalization request for a listing.
	Create(ctx context.Context, request CreatePersonalizationRequest) (*CreatePersonalizationRequestResponse, error)

	// Get the personalization request for a listing.
	//
	// Get the personalization request for a listing. Each consumer can make at
	// *most* one personalization request for a listing.
	Get(ctx context.Context, request GetPersonalizationRequestRequest) (*GetPersonalizationRequestResponse, error)

	// List all personalization requests.
	//
	// List personalization requests for a consumer across all listings.
	//
	// Use ListAll() to get all PersonalizationRequest instances, which will iterate over every result page.
	List(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error)
}

// Providers are the entities that publish listings to the Marketplace.
type ConsumerProvidersService interface {

	// Get one batch of providers. One may specify up to 50 IDs per request.
	//
	// Batch get a provider in the Databricks Marketplace with at least one
	// visible listing.
	BatchGet(ctx context.Context, request BatchGetProvidersRequest) (*BatchGetProvidersResponse, error)

	// Get a provider.
	//
	// Get a provider in the Databricks Marketplace with at least one visible
	// listing.
	Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error)

	// List providers.
	//
	// List all providers in the Databricks Marketplace with at least one
	// visible listing.
	//
	// Use ListAll() to get all ProviderInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error)
}

// Marketplace exchanges filters curate which groups can access an exchange.
type ProviderExchangeFiltersService interface {

	// Create a new exchange filter.
	//
	// Add an exchange filter.
	Create(ctx context.Context, request CreateExchangeFilterRequest) (*CreateExchangeFilterResponse, error)

	// Delete an exchange filter.
	//
	// Delete an exchange filter
	Delete(ctx context.Context, request DeleteExchangeFilterRequest) error

	// List exchange filters.
	//
	// List exchange filter
	//
	// Use ListAll() to get all ExchangeFilter instances, which will iterate over every result page.
	List(ctx context.Context, request ListExchangeFiltersRequest) (*ListExchangeFiltersResponse, error)

	// Update exchange filter.
	//
	// Update an exchange filter.
	Update(ctx context.Context, request UpdateExchangeFilterRequest) (*UpdateExchangeFilterResponse, error)
}

// Marketplace exchanges allow providers to share their listings with a curated
// set of customers.
type ProviderExchangesService interface {

	// Add an exchange for listing.
	//
	// Associate an exchange with a listing
	AddListingToExchange(ctx context.Context, request AddExchangeForListingRequest) (*AddExchangeForListingResponse, error)

	// Create an exchange.
	//
	// Create an exchange
	Create(ctx context.Context, request CreateExchangeRequest) (*CreateExchangeResponse, error)

	// Delete an exchange.
	//
	// This removes a listing from marketplace.
	Delete(ctx context.Context, request DeleteExchangeRequest) error

	// Remove an exchange for listing.
	//
	// Disassociate an exchange with a listing
	DeleteListingFromExchange(ctx context.Context, request RemoveExchangeForListingRequest) error

	// Get an exchange.
	//
	// Get an exchange.
	Get(ctx context.Context, request GetExchangeRequest) (*GetExchangeResponse, error)

	// List exchanges.
	//
	// List exchanges visible to provider
	//
	// Use ListAll() to get all Exchange instances, which will iterate over every result page.
	List(ctx context.Context, request ListExchangesRequest) (*ListExchangesResponse, error)

	// List exchanges for listing.
	//
	// List exchanges associated with a listing
	//
	// Use ListExchangesForListingAll() to get all ExchangeListing instances, which will iterate over every result page.
	ListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) (*ListExchangesForListingResponse, error)

	// List listings for exchange.
	//
	// List listings associated with an exchange
	//
	// Use ListListingsForExchangeAll() to get all ExchangeListing instances, which will iterate over every result page.
	ListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) (*ListListingsForExchangeResponse, error)

	// Update exchange.
	//
	// Update an exchange
	Update(ctx context.Context, request UpdateExchangeRequest) (*UpdateExchangeResponse, error)
}

// Marketplace offers a set of file APIs for various purposes such as preview
// notebooks and provider icons.
type ProviderFilesService interface {

	// Create a file.
	//
	// Create a file. Currently, only provider icons and attached notebooks are
	// supported.
	Create(ctx context.Context, request CreateFileRequest) (*CreateFileResponse, error)

	// Delete a file.
	//
	// Delete a file
	Delete(ctx context.Context, request DeleteFileRequest) error

	// Get a file.
	//
	// Get a file
	Get(ctx context.Context, request GetFileRequest) (*GetFileResponse, error)

	// List files.
	//
	// List files attached to a parent entity.
	//
	// Use ListAll() to get all FileInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error)
}

// Listings are the core entities in the Marketplace. They represent the
// products that are available for consumption.
type ProviderListingsService interface {

	// Create a listing.
	//
	// Create a new listing
	Create(ctx context.Context, request CreateListingRequest) (*CreateListingResponse, error)

	// Delete a listing.
	//
	// Delete a listing
	Delete(ctx context.Context, request DeleteListingRequest) error

	// Get a listing.
	//
	// Get a listing
	Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error)

	// List listings.
	//
	// List listings owned by this provider
	//
	// Use ListAll() to get all Listing instances, which will iterate over every result page.
	List(ctx context.Context, request GetListingsRequest) (*GetListingsResponse, error)

	// Update listing.
	//
	// Update a listing
	Update(ctx context.Context, request UpdateListingRequest) (*UpdateListingResponse, error)
}

// Personalization requests are an alternate to instantly available listings.
// Control the lifecycle of personalized solutions.
type ProviderPersonalizationRequestsService interface {

	// All personalization requests across all listings.
	//
	// List personalization requests to this provider. This will return all
	// personalization requests, regardless of which listing they are for.
	//
	// Use ListAll() to get all PersonalizationRequest instances, which will iterate over every result page.
	List(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error)

	// Update personalization request status.
	//
	// Update personalization request. This method only permits updating the
	// status of the request.
	Update(ctx context.Context, request UpdatePersonalizationRequestRequest) (*UpdatePersonalizationRequestResponse, error)
}

// Manage templated analytics solution for providers.
type ProviderProviderAnalyticsDashboardsService interface {

	// Create provider analytics dashboard.
	//
	// Create provider analytics dashboard. Returns Marketplace specific `id`.
	// Not to be confused with the Lakeview dashboard id.
	Create(ctx context.Context) (*ProviderAnalyticsDashboard, error)

	// Get provider analytics dashboard.
	//
	// Get provider analytics dashboard.
	Get(ctx context.Context) (*ListProviderAnalyticsDashboardResponse, error)

	// Get latest version of provider analytics dashboard.
	//
	// Get latest version of provider analytics dashboard.
	GetLatestVersion(ctx context.Context) (*GetLatestVersionProviderAnalyticsDashboardResponse, error)

	// Update provider analytics dashboard.
	//
	// Update provider analytics dashboard.
	Update(ctx context.Context, request UpdateProviderAnalyticsDashboardRequest) (*UpdateProviderAnalyticsDashboardResponse, error)
}

// Providers are entities that manage assets in Marketplace.
type ProviderProvidersService interface {

	// Create a provider.
	//
	// Create a provider
	Create(ctx context.Context, request CreateProviderRequest) (*CreateProviderResponse, error)

	// Delete provider.
	//
	// Delete provider
	Delete(ctx context.Context, request DeleteProviderRequest) error

	// Get provider.
	//
	// Get provider profile
	Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error)

	// List providers.
	//
	// List provider profiles for account.
	//
	// Use ListAll() to get all ProviderInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error)

	// Update provider.
	//
	// Update provider profile
	Update(ctx context.Context, request UpdateProviderRequest) (*UpdateProviderResponse, error)
}
