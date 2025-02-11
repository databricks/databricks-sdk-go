// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplacepreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just ConsumerFulfillmentsPreview API methods
type consumerFulfillmentsPreviewImpl struct {
	client *client.DatabricksClient
}

// Get listing content metadata.
//
// Get a high level preview of the metadata of listing installable content.
func (a *consumerFulfillmentsPreviewImpl) Get(ctx context.Context, request GetListingContentMetadataRequest) listing.Iterator[SharedDataObject] {

	getNextPage := func(ctx context.Context, req GetListingContentMetadataRequest) (*GetListingContentMetadataResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGet(ctx, req)
	}
	getItems := func(resp *GetListingContentMetadataResponse) []SharedDataObject {
		return resp.SharedDataObjects
	}
	getNextReq := func(resp *GetListingContentMetadataResponse) *GetListingContentMetadataRequest {
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

// Get listing content metadata.
//
// Get a high level preview of the metadata of listing installable content.
func (a *consumerFulfillmentsPreviewImpl) GetAll(ctx context.Context, request GetListingContentMetadataRequest) ([]SharedDataObject, error) {
	iterator := a.Get(ctx, request)
	return listing.ToSlice[SharedDataObject](ctx, iterator)
}
func (a *consumerFulfillmentsPreviewImpl) internalGet(ctx context.Context, request GetListingContentMetadataRequest) (*GetListingContentMetadataResponse, error) {
	var getListingContentMetadataResponse GetListingContentMetadataResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v/content", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingContentMetadataResponse)
	return &getListingContentMetadataResponse, err
}

// List all listing fulfillments.
//
// Get all listings fulfillments associated with a listing. A _fulfillment_ is a
// potential installation. Standard installations contain metadata about the
// attached share or git repo. Only one of these fields will be present.
// Personalized installations contain metadata about the attached share or git
// repo, as well as the Delta Sharing recipient type.
func (a *consumerFulfillmentsPreviewImpl) List(ctx context.Context, request ListFulfillmentsRequest) listing.Iterator[ListingFulfillment] {

	getNextPage := func(ctx context.Context, req ListFulfillmentsRequest) (*ListFulfillmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFulfillmentsResponse) []ListingFulfillment {
		return resp.Fulfillments
	}
	getNextReq := func(resp *ListFulfillmentsResponse) *ListFulfillmentsRequest {
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

// List all listing fulfillments.
//
// Get all listings fulfillments associated with a listing. A _fulfillment_ is a
// potential installation. Standard installations contain metadata about the
// attached share or git repo. Only one of these fields will be present.
// Personalized installations contain metadata about the attached share or git
// repo, as well as the Delta Sharing recipient type.
func (a *consumerFulfillmentsPreviewImpl) ListAll(ctx context.Context, request ListFulfillmentsRequest) ([]ListingFulfillment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListingFulfillment](ctx, iterator)
}
func (a *consumerFulfillmentsPreviewImpl) internalList(ctx context.Context, request ListFulfillmentsRequest) (*ListFulfillmentsResponse, error) {
	var listFulfillmentsResponse ListFulfillmentsResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v/fulfillments", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFulfillmentsResponse)
	return &listFulfillmentsResponse, err
}

// unexported type that holds implementations of just ConsumerInstallationsPreview API methods
type consumerInstallationsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *consumerInstallationsPreviewImpl) Create(ctx context.Context, request CreateInstallationRequest) (*Installation, error) {
	var installation Installation
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v/installations", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &installation)
	return &installation, err
}

func (a *consumerInstallationsPreviewImpl) Delete(ctx context.Context, request DeleteInstallationRequest) error {
	var deleteInstallationResponse DeleteInstallationResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v/installations/%v", request.ListingId, request.InstallationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteInstallationResponse)
	return err
}

// List all installations.
//
// List all installations across all listings.
func (a *consumerInstallationsPreviewImpl) List(ctx context.Context, request ListAllInstallationsRequest) listing.Iterator[InstallationDetail] {

	getNextPage := func(ctx context.Context, req ListAllInstallationsRequest) (*ListAllInstallationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAllInstallationsResponse) []InstallationDetail {
		return resp.Installations
	}
	getNextReq := func(resp *ListAllInstallationsResponse) *ListAllInstallationsRequest {
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

// List all installations.
//
// List all installations across all listings.
func (a *consumerInstallationsPreviewImpl) ListAll(ctx context.Context, request ListAllInstallationsRequest) ([]InstallationDetail, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[InstallationDetail](ctx, iterator)
}
func (a *consumerInstallationsPreviewImpl) internalList(ctx context.Context, request ListAllInstallationsRequest) (*ListAllInstallationsResponse, error) {
	var listAllInstallationsResponse ListAllInstallationsResponse
	path := "/api/2.1preview/marketplace-consumer/installations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllInstallationsResponse)
	return &listAllInstallationsResponse, err
}

// List installations for a listing.
//
// List all installations for a particular listing.
func (a *consumerInstallationsPreviewImpl) ListListingInstallations(ctx context.Context, request ListInstallationsRequest) listing.Iterator[InstallationDetail] {

	getNextPage := func(ctx context.Context, req ListInstallationsRequest) (*ListInstallationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListListingInstallations(ctx, req)
	}
	getItems := func(resp *ListInstallationsResponse) []InstallationDetail {
		return resp.Installations
	}
	getNextReq := func(resp *ListInstallationsResponse) *ListInstallationsRequest {
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

// List installations for a listing.
//
// List all installations for a particular listing.
func (a *consumerInstallationsPreviewImpl) ListListingInstallationsAll(ctx context.Context, request ListInstallationsRequest) ([]InstallationDetail, error) {
	iterator := a.ListListingInstallations(ctx, request)
	return listing.ToSlice[InstallationDetail](ctx, iterator)
}
func (a *consumerInstallationsPreviewImpl) internalListListingInstallations(ctx context.Context, request ListInstallationsRequest) (*ListInstallationsResponse, error) {
	var listInstallationsResponse ListInstallationsResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v/installations", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listInstallationsResponse)
	return &listInstallationsResponse, err
}

func (a *consumerInstallationsPreviewImpl) Update(ctx context.Context, request UpdateInstallationRequest) (*UpdateInstallationResponse, error) {
	var updateInstallationResponse UpdateInstallationResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v/installations/%v", request.ListingId, request.InstallationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateInstallationResponse)
	return &updateInstallationResponse, err
}

// unexported type that holds implementations of just ConsumerListingsPreview API methods
type consumerListingsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *consumerListingsPreviewImpl) BatchGet(ctx context.Context, request BatchGetListingsRequest) (*BatchGetListingsResponse, error) {
	var batchGetListingsResponse BatchGetListingsResponse
	path := "/api/2.1preview/marketplace-consumer/listings:batchGet"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &batchGetListingsResponse)
	return &batchGetListingsResponse, err
}

func (a *consumerListingsPreviewImpl) Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error) {
	var getListingResponse GetListingResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingResponse)
	return &getListingResponse, err
}

// List listings.
//
// List all published listings in the Databricks Marketplace that the consumer
// has access to.
func (a *consumerListingsPreviewImpl) List(ctx context.Context, request ListListingsRequest) listing.Iterator[Listing] {

	getNextPage := func(ctx context.Context, req ListListingsRequest) (*ListListingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListListingsResponse) []Listing {
		return resp.Listings
	}
	getNextReq := func(resp *ListListingsResponse) *ListListingsRequest {
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

// List listings.
//
// List all published listings in the Databricks Marketplace that the consumer
// has access to.
func (a *consumerListingsPreviewImpl) ListAll(ctx context.Context, request ListListingsRequest) ([]Listing, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}
func (a *consumerListingsPreviewImpl) internalList(ctx context.Context, request ListListingsRequest) (*ListListingsResponse, error) {
	var listListingsResponse ListListingsResponse
	path := "/api/2.1preview/marketplace-consumer/listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listListingsResponse)
	return &listListingsResponse, err
}

// Search listings.
//
// Search published listings in the Databricks Marketplace that the consumer has
// access to. This query supports a variety of different search parameters and
// performs fuzzy matching.
func (a *consumerListingsPreviewImpl) Search(ctx context.Context, request SearchListingsRequest) listing.Iterator[Listing] {

	getNextPage := func(ctx context.Context, req SearchListingsRequest) (*SearchListingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearch(ctx, req)
	}
	getItems := func(resp *SearchListingsResponse) []Listing {
		return resp.Listings
	}
	getNextReq := func(resp *SearchListingsResponse) *SearchListingsRequest {
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

// Search listings.
//
// Search published listings in the Databricks Marketplace that the consumer has
// access to. This query supports a variety of different search parameters and
// performs fuzzy matching.
func (a *consumerListingsPreviewImpl) SearchAll(ctx context.Context, request SearchListingsRequest) ([]Listing, error) {
	iterator := a.Search(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}
func (a *consumerListingsPreviewImpl) internalSearch(ctx context.Context, request SearchListingsRequest) (*SearchListingsResponse, error) {
	var searchListingsResponse SearchListingsResponse
	path := "/api/2.1preview/marketplace-consumer/search-listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &searchListingsResponse)
	return &searchListingsResponse, err
}

// unexported type that holds implementations of just ConsumerPersonalizationRequestsPreview API methods
type consumerPersonalizationRequestsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *consumerPersonalizationRequestsPreviewImpl) Create(ctx context.Context, request CreatePersonalizationRequest) (*CreatePersonalizationRequestResponse, error) {
	var createPersonalizationRequestResponse CreatePersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v/personalization-requests", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createPersonalizationRequestResponse)
	return &createPersonalizationRequestResponse, err
}

func (a *consumerPersonalizationRequestsPreviewImpl) Get(ctx context.Context, request GetPersonalizationRequestRequest) (*GetPersonalizationRequestResponse, error) {
	var getPersonalizationRequestResponse GetPersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/listings/%v/personalization-requests", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPersonalizationRequestResponse)
	return &getPersonalizationRequestResponse, err
}

// List all personalization requests.
//
// List personalization requests for a consumer across all listings.
func (a *consumerPersonalizationRequestsPreviewImpl) List(ctx context.Context, request ListAllPersonalizationRequestsRequest) listing.Iterator[PersonalizationRequest] {

	getNextPage := func(ctx context.Context, req ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAllPersonalizationRequestsResponse) []PersonalizationRequest {
		return resp.PersonalizationRequests
	}
	getNextReq := func(resp *ListAllPersonalizationRequestsResponse) *ListAllPersonalizationRequestsRequest {
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

// List all personalization requests.
//
// List personalization requests for a consumer across all listings.
func (a *consumerPersonalizationRequestsPreviewImpl) ListAll(ctx context.Context, request ListAllPersonalizationRequestsRequest) ([]PersonalizationRequest, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PersonalizationRequest](ctx, iterator)
}
func (a *consumerPersonalizationRequestsPreviewImpl) internalList(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
	var listAllPersonalizationRequestsResponse ListAllPersonalizationRequestsResponse
	path := "/api/2.1preview/marketplace-consumer/personalization-requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllPersonalizationRequestsResponse)
	return &listAllPersonalizationRequestsResponse, err
}

// unexported type that holds implementations of just ConsumerProvidersPreview API methods
type consumerProvidersPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *consumerProvidersPreviewImpl) BatchGet(ctx context.Context, request BatchGetProvidersRequest) (*BatchGetProvidersResponse, error) {
	var batchGetProvidersResponse BatchGetProvidersResponse
	path := "/api/2.1preview/marketplace-consumer/providers:batchGet"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &batchGetProvidersResponse)
	return &batchGetProvidersResponse, err
}

func (a *consumerProvidersPreviewImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.1preview/marketplace-consumer/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getProviderResponse)
	return &getProviderResponse, err
}

// List providers.
//
// List all providers in the Databricks Marketplace with at least one visible
// listing.
func (a *consumerProvidersPreviewImpl) List(ctx context.Context, request ListProvidersRequest) listing.Iterator[ProviderInfo] {

	getNextPage := func(ctx context.Context, req ListProvidersRequest) (*ListProvidersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListProvidersResponse) []ProviderInfo {
		return resp.Providers
	}
	getNextReq := func(resp *ListProvidersResponse) *ListProvidersRequest {
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

// List providers.
//
// List all providers in the Databricks Marketplace with at least one visible
// listing.
func (a *consumerProvidersPreviewImpl) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}
func (a *consumerProvidersPreviewImpl) internalList(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1preview/marketplace-consumer/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProvidersResponse)
	return &listProvidersResponse, err
}

// unexported type that holds implementations of just ProviderExchangeFiltersPreview API methods
type providerExchangeFiltersPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *providerExchangeFiltersPreviewImpl) Create(ctx context.Context, request CreateExchangeFilterRequest) (*CreateExchangeFilterResponse, error) {
	var createExchangeFilterResponse CreateExchangeFilterResponse
	path := "/api/2.0preview/marketplace-exchange/filters"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createExchangeFilterResponse)
	return &createExchangeFilterResponse, err
}

func (a *providerExchangeFiltersPreviewImpl) Delete(ctx context.Context, request DeleteExchangeFilterRequest) error {
	var deleteExchangeFilterResponse DeleteExchangeFilterResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-exchange/filters/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteExchangeFilterResponse)
	return err
}

// List exchange filters.
//
// List exchange filter
func (a *providerExchangeFiltersPreviewImpl) List(ctx context.Context, request ListExchangeFiltersRequest) listing.Iterator[ExchangeFilter] {

	getNextPage := func(ctx context.Context, req ListExchangeFiltersRequest) (*ListExchangeFiltersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListExchangeFiltersResponse) []ExchangeFilter {
		return resp.Filters
	}
	getNextReq := func(resp *ListExchangeFiltersResponse) *ListExchangeFiltersRequest {
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

// List exchange filters.
//
// List exchange filter
func (a *providerExchangeFiltersPreviewImpl) ListAll(ctx context.Context, request ListExchangeFiltersRequest) ([]ExchangeFilter, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ExchangeFilter](ctx, iterator)
}
func (a *providerExchangeFiltersPreviewImpl) internalList(ctx context.Context, request ListExchangeFiltersRequest) (*ListExchangeFiltersResponse, error) {
	var listExchangeFiltersResponse ListExchangeFiltersResponse
	path := "/api/2.0preview/marketplace-exchange/filters"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangeFiltersResponse)
	return &listExchangeFiltersResponse, err
}

func (a *providerExchangeFiltersPreviewImpl) Update(ctx context.Context, request UpdateExchangeFilterRequest) (*UpdateExchangeFilterResponse, error) {
	var updateExchangeFilterResponse UpdateExchangeFilterResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-exchange/filters/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateExchangeFilterResponse)
	return &updateExchangeFilterResponse, err
}

// unexported type that holds implementations of just ProviderExchangesPreview API methods
type providerExchangesPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *providerExchangesPreviewImpl) AddListingToExchange(ctx context.Context, request AddExchangeForListingRequest) (*AddExchangeForListingResponse, error) {
	var addExchangeForListingResponse AddExchangeForListingResponse
	path := "/api/2.0preview/marketplace-exchange/exchanges-for-listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &addExchangeForListingResponse)
	return &addExchangeForListingResponse, err
}

func (a *providerExchangesPreviewImpl) Create(ctx context.Context, request CreateExchangeRequest) (*CreateExchangeResponse, error) {
	var createExchangeResponse CreateExchangeResponse
	path := "/api/2.0preview/marketplace-exchange/exchanges"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createExchangeResponse)
	return &createExchangeResponse, err
}

func (a *providerExchangesPreviewImpl) Delete(ctx context.Context, request DeleteExchangeRequest) error {
	var deleteExchangeResponse DeleteExchangeResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteExchangeResponse)
	return err
}

func (a *providerExchangesPreviewImpl) DeleteListingFromExchange(ctx context.Context, request RemoveExchangeForListingRequest) error {
	var removeExchangeForListingResponse RemoveExchangeForListingResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-exchange/exchanges-for-listing/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &removeExchangeForListingResponse)
	return err
}

func (a *providerExchangesPreviewImpl) Get(ctx context.Context, request GetExchangeRequest) (*GetExchangeResponse, error) {
	var getExchangeResponse GetExchangeResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExchangeResponse)
	return &getExchangeResponse, err
}

// List exchanges.
//
// List exchanges visible to provider
func (a *providerExchangesPreviewImpl) List(ctx context.Context, request ListExchangesRequest) listing.Iterator[Exchange] {

	getNextPage := func(ctx context.Context, req ListExchangesRequest) (*ListExchangesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListExchangesResponse) []Exchange {
		return resp.Exchanges
	}
	getNextReq := func(resp *ListExchangesResponse) *ListExchangesRequest {
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

// List exchanges.
//
// List exchanges visible to provider
func (a *providerExchangesPreviewImpl) ListAll(ctx context.Context, request ListExchangesRequest) ([]Exchange, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Exchange](ctx, iterator)
}
func (a *providerExchangesPreviewImpl) internalList(ctx context.Context, request ListExchangesRequest) (*ListExchangesResponse, error) {
	var listExchangesResponse ListExchangesResponse
	path := "/api/2.0preview/marketplace-exchange/exchanges"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangesResponse)
	return &listExchangesResponse, err
}

// List exchanges for listing.
//
// List exchanges associated with a listing
func (a *providerExchangesPreviewImpl) ListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) listing.Iterator[ExchangeListing] {

	getNextPage := func(ctx context.Context, req ListExchangesForListingRequest) (*ListExchangesForListingResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExchangesForListing(ctx, req)
	}
	getItems := func(resp *ListExchangesForListingResponse) []ExchangeListing {
		return resp.ExchangeListing
	}
	getNextReq := func(resp *ListExchangesForListingResponse) *ListExchangesForListingRequest {
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

// List exchanges for listing.
//
// List exchanges associated with a listing
func (a *providerExchangesPreviewImpl) ListExchangesForListingAll(ctx context.Context, request ListExchangesForListingRequest) ([]ExchangeListing, error) {
	iterator := a.ListExchangesForListing(ctx, request)
	return listing.ToSlice[ExchangeListing](ctx, iterator)
}
func (a *providerExchangesPreviewImpl) internalListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) (*ListExchangesForListingResponse, error) {
	var listExchangesForListingResponse ListExchangesForListingResponse
	path := "/api/2.0preview/marketplace-exchange/exchanges-for-listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangesForListingResponse)
	return &listExchangesForListingResponse, err
}

// List listings for exchange.
//
// List listings associated with an exchange
func (a *providerExchangesPreviewImpl) ListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) listing.Iterator[ExchangeListing] {

	getNextPage := func(ctx context.Context, req ListListingsForExchangeRequest) (*ListListingsForExchangeResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListListingsForExchange(ctx, req)
	}
	getItems := func(resp *ListListingsForExchangeResponse) []ExchangeListing {
		return resp.ExchangeListings
	}
	getNextReq := func(resp *ListListingsForExchangeResponse) *ListListingsForExchangeRequest {
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

// List listings for exchange.
//
// List listings associated with an exchange
func (a *providerExchangesPreviewImpl) ListListingsForExchangeAll(ctx context.Context, request ListListingsForExchangeRequest) ([]ExchangeListing, error) {
	iterator := a.ListListingsForExchange(ctx, request)
	return listing.ToSlice[ExchangeListing](ctx, iterator)
}
func (a *providerExchangesPreviewImpl) internalListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) (*ListListingsForExchangeResponse, error) {
	var listListingsForExchangeResponse ListListingsForExchangeResponse
	path := "/api/2.0preview/marketplace-exchange/listings-for-exchange"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listListingsForExchangeResponse)
	return &listListingsForExchangeResponse, err
}

func (a *providerExchangesPreviewImpl) Update(ctx context.Context, request UpdateExchangeRequest) (*UpdateExchangeResponse, error) {
	var updateExchangeResponse UpdateExchangeResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateExchangeResponse)
	return &updateExchangeResponse, err
}

// unexported type that holds implementations of just ProviderFilesPreview API methods
type providerFilesPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *providerFilesPreviewImpl) Create(ctx context.Context, request CreateFileRequest) (*CreateFileResponse, error) {
	var createFileResponse CreateFileResponse
	path := "/api/2.0preview/marketplace-provider/files"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createFileResponse)
	return &createFileResponse, err
}

func (a *providerFilesPreviewImpl) Delete(ctx context.Context, request DeleteFileRequest) error {
	var deleteFileResponse DeleteFileResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/files/%v", request.FileId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteFileResponse)
	return err
}

func (a *providerFilesPreviewImpl) Get(ctx context.Context, request GetFileRequest) (*GetFileResponse, error) {
	var getFileResponse GetFileResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/files/%v", request.FileId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getFileResponse)
	return &getFileResponse, err
}

// List files.
//
// List files attached to a parent entity.
func (a *providerFilesPreviewImpl) List(ctx context.Context, request ListFilesRequest) listing.Iterator[FileInfo] {

	getNextPage := func(ctx context.Context, req ListFilesRequest) (*ListFilesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFilesResponse) []FileInfo {
		return resp.FileInfos
	}
	getNextReq := func(resp *ListFilesResponse) *ListFilesRequest {
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

// List files.
//
// List files attached to a parent entity.
func (a *providerFilesPreviewImpl) ListAll(ctx context.Context, request ListFilesRequest) ([]FileInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FileInfo](ctx, iterator)
}
func (a *providerFilesPreviewImpl) internalList(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error) {
	var listFilesResponse ListFilesResponse
	path := "/api/2.0preview/marketplace-provider/files"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFilesResponse)
	return &listFilesResponse, err
}

// unexported type that holds implementations of just ProviderListingsPreview API methods
type providerListingsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *providerListingsPreviewImpl) Create(ctx context.Context, request CreateListingRequest) (*CreateListingResponse, error) {
	var createListingResponse CreateListingResponse
	path := "/api/2.0preview/marketplace-provider/listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createListingResponse)
	return &createListingResponse, err
}

func (a *providerListingsPreviewImpl) Delete(ctx context.Context, request DeleteListingRequest) error {
	var deleteListingResponse DeleteListingResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteListingResponse)
	return err
}

func (a *providerListingsPreviewImpl) Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error) {
	var getListingResponse GetListingResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingResponse)
	return &getListingResponse, err
}

// List listings.
//
// List listings owned by this provider
func (a *providerListingsPreviewImpl) List(ctx context.Context, request GetListingsRequest) listing.Iterator[Listing] {

	getNextPage := func(ctx context.Context, req GetListingsRequest) (*GetListingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *GetListingsResponse) []Listing {
		return resp.Listings
	}
	getNextReq := func(resp *GetListingsResponse) *GetListingsRequest {
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

// List listings.
//
// List listings owned by this provider
func (a *providerListingsPreviewImpl) ListAll(ctx context.Context, request GetListingsRequest) ([]Listing, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}
func (a *providerListingsPreviewImpl) internalList(ctx context.Context, request GetListingsRequest) (*GetListingsResponse, error) {
	var getListingsResponse GetListingsResponse
	path := "/api/2.0preview/marketplace-provider/listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingsResponse)
	return &getListingsResponse, err
}

func (a *providerListingsPreviewImpl) Update(ctx context.Context, request UpdateListingRequest) (*UpdateListingResponse, error) {
	var updateListingResponse UpdateListingResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateListingResponse)
	return &updateListingResponse, err
}

// unexported type that holds implementations of just ProviderPersonalizationRequestsPreview API methods
type providerPersonalizationRequestsPreviewImpl struct {
	client *client.DatabricksClient
}

// All personalization requests across all listings.
//
// List personalization requests to this provider. This will return all
// personalization requests, regardless of which listing they are for.
func (a *providerPersonalizationRequestsPreviewImpl) List(ctx context.Context, request ListAllPersonalizationRequestsRequest) listing.Iterator[PersonalizationRequest] {

	getNextPage := func(ctx context.Context, req ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAllPersonalizationRequestsResponse) []PersonalizationRequest {
		return resp.PersonalizationRequests
	}
	getNextReq := func(resp *ListAllPersonalizationRequestsResponse) *ListAllPersonalizationRequestsRequest {
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

// All personalization requests across all listings.
//
// List personalization requests to this provider. This will return all
// personalization requests, regardless of which listing they are for.
func (a *providerPersonalizationRequestsPreviewImpl) ListAll(ctx context.Context, request ListAllPersonalizationRequestsRequest) ([]PersonalizationRequest, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PersonalizationRequest](ctx, iterator)
}
func (a *providerPersonalizationRequestsPreviewImpl) internalList(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
	var listAllPersonalizationRequestsResponse ListAllPersonalizationRequestsResponse
	path := "/api/2.0preview/marketplace-provider/personalization-requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllPersonalizationRequestsResponse)
	return &listAllPersonalizationRequestsResponse, err
}

func (a *providerPersonalizationRequestsPreviewImpl) Update(ctx context.Context, request UpdatePersonalizationRequestRequest) (*UpdatePersonalizationRequestResponse, error) {
	var updatePersonalizationRequestResponse UpdatePersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/listings/%v/personalization-requests/%v/request-status", request.ListingId, request.RequestId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updatePersonalizationRequestResponse)
	return &updatePersonalizationRequestResponse, err
}

// unexported type that holds implementations of just ProviderProviderAnalyticsDashboardsPreview API methods
type providerProviderAnalyticsDashboardsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *providerProviderAnalyticsDashboardsPreviewImpl) Create(ctx context.Context) (*ProviderAnalyticsDashboard, error) {
	var providerAnalyticsDashboard ProviderAnalyticsDashboard
	path := "/api/2.0preview/marketplace-provider/analytics_dashboard"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, nil, &providerAnalyticsDashboard)
	return &providerAnalyticsDashboard, err
}

func (a *providerProviderAnalyticsDashboardsPreviewImpl) Get(ctx context.Context) (*ListProviderAnalyticsDashboardResponse, error) {
	var listProviderAnalyticsDashboardResponse ListProviderAnalyticsDashboardResponse
	path := "/api/2.0preview/marketplace-provider/analytics_dashboard"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listProviderAnalyticsDashboardResponse)
	return &listProviderAnalyticsDashboardResponse, err
}

func (a *providerProviderAnalyticsDashboardsPreviewImpl) GetLatestVersion(ctx context.Context) (*GetLatestVersionProviderAnalyticsDashboardResponse, error) {
	var getLatestVersionProviderAnalyticsDashboardResponse GetLatestVersionProviderAnalyticsDashboardResponse
	path := "/api/2.0preview/marketplace-provider/analytics_dashboard/latest"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getLatestVersionProviderAnalyticsDashboardResponse)
	return &getLatestVersionProviderAnalyticsDashboardResponse, err
}

func (a *providerProviderAnalyticsDashboardsPreviewImpl) Update(ctx context.Context, request UpdateProviderAnalyticsDashboardRequest) (*UpdateProviderAnalyticsDashboardResponse, error) {
	var updateProviderAnalyticsDashboardResponse UpdateProviderAnalyticsDashboardResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/analytics_dashboard/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateProviderAnalyticsDashboardResponse)
	return &updateProviderAnalyticsDashboardResponse, err
}

// unexported type that holds implementations of just ProviderProvidersPreview API methods
type providerProvidersPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *providerProvidersPreviewImpl) Create(ctx context.Context, request CreateProviderRequest) (*CreateProviderResponse, error) {
	var createProviderResponse CreateProviderResponse
	path := "/api/2.0preview/marketplace-provider/provider"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createProviderResponse)
	return &createProviderResponse, err
}

func (a *providerProvidersPreviewImpl) Delete(ctx context.Context, request DeleteProviderRequest) error {
	var deleteProviderResponse DeleteProviderResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteProviderResponse)
	return err
}

func (a *providerProvidersPreviewImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getProviderResponse)
	return &getProviderResponse, err
}

// List providers.
//
// List provider profiles for account.
func (a *providerProvidersPreviewImpl) List(ctx context.Context, request ListProvidersRequest) listing.Iterator[ProviderInfo] {

	getNextPage := func(ctx context.Context, req ListProvidersRequest) (*ListProvidersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListProvidersResponse) []ProviderInfo {
		return resp.Providers
	}
	getNextReq := func(resp *ListProvidersResponse) *ListProvidersRequest {
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

// List providers.
//
// List provider profiles for account.
func (a *providerProvidersPreviewImpl) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}
func (a *providerProvidersPreviewImpl) internalList(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.0preview/marketplace-provider/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProvidersResponse)
	return &listProvidersResponse, err
}

func (a *providerProvidersPreviewImpl) Update(ctx context.Context, request UpdateProviderRequest) (*UpdateProviderResponse, error) {
	var updateProviderResponse UpdateProviderResponse
	path := fmt.Sprintf("/api/2.0preview/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateProviderResponse)
	return &updateProviderResponse, err
}
