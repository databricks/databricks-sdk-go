// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/marketplace/marketplacepb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just ConsumerFulfillments API methods
type consumerFulfillmentsImpl struct {
	client *client.DatabricksClient
}

// Get a high level preview of the metadata of listing installable content.
func (a *consumerFulfillmentsImpl) Get(ctx context.Context, request GetListingContentMetadataRequest) listing.Iterator[SharedDataObject] {

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

// Get a high level preview of the metadata of listing installable content.
func (a *consumerFulfillmentsImpl) GetAll(ctx context.Context, request GetListingContentMetadataRequest) ([]SharedDataObject, error) {
	iterator := a.Get(ctx, request)
	return listing.ToSlice[SharedDataObject](ctx, iterator)
}

func (a *consumerFulfillmentsImpl) internalGet(ctx context.Context, request GetListingContentMetadataRequest) (*GetListingContentMetadataResponse, error) {
	requestPb, pbErr := GetListingContentMetadataRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getListingContentMetadataResponsePb marketplacepb.GetListingContentMetadataResponsePb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/content", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getListingContentMetadataResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := GetListingContentMetadataResponseFromPb(&getListingContentMetadataResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Get all listings fulfillments associated with a listing. A _fulfillment_ is a
// potential installation. Standard installations contain metadata about the
// attached share or git repo. Only one of these fields will be present.
// Personalized installations contain metadata about the attached share or git
// repo, as well as the Delta Sharing recipient type.
func (a *consumerFulfillmentsImpl) List(ctx context.Context, request ListFulfillmentsRequest) listing.Iterator[ListingFulfillment] {

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

// Get all listings fulfillments associated with a listing. A _fulfillment_ is a
// potential installation. Standard installations contain metadata about the
// attached share or git repo. Only one of these fields will be present.
// Personalized installations contain metadata about the attached share or git
// repo, as well as the Delta Sharing recipient type.
func (a *consumerFulfillmentsImpl) ListAll(ctx context.Context, request ListFulfillmentsRequest) ([]ListingFulfillment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListingFulfillment](ctx, iterator)
}

func (a *consumerFulfillmentsImpl) internalList(ctx context.Context, request ListFulfillmentsRequest) (*ListFulfillmentsResponse, error) {
	requestPb, pbErr := ListFulfillmentsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listFulfillmentsResponsePb marketplacepb.ListFulfillmentsResponsePb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/fulfillments", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listFulfillmentsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListFulfillmentsResponseFromPb(&listFulfillmentsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ConsumerInstallations API methods
type consumerInstallationsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerInstallationsImpl) Create(ctx context.Context, request CreateInstallationRequest) (*Installation, error) {
	requestPb, pbErr := CreateInstallationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var installationPb marketplacepb.InstallationPb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&installationPb,
	)
	resp, err := InstallationFromPb(&installationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *consumerInstallationsImpl) Delete(ctx context.Context, request DeleteInstallationRequest) error {
	requestPb, pbErr := DeleteInstallationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations/%v", request.ListingId, request.InstallationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// List all installations across all listings.
func (a *consumerInstallationsImpl) List(ctx context.Context, request ListAllInstallationsRequest) listing.Iterator[InstallationDetail] {

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

// List all installations across all listings.
func (a *consumerInstallationsImpl) ListAll(ctx context.Context, request ListAllInstallationsRequest) ([]InstallationDetail, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[InstallationDetail](ctx, iterator)
}

func (a *consumerInstallationsImpl) internalList(ctx context.Context, request ListAllInstallationsRequest) (*ListAllInstallationsResponse, error) {
	requestPb, pbErr := ListAllInstallationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAllInstallationsResponsePb marketplacepb.ListAllInstallationsResponsePb
	path := "/api/2.1/marketplace-consumer/installations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listAllInstallationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListAllInstallationsResponseFromPb(&listAllInstallationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List all installations for a particular listing.
func (a *consumerInstallationsImpl) ListListingInstallations(ctx context.Context, request ListInstallationsRequest) listing.Iterator[InstallationDetail] {

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

// List all installations for a particular listing.
func (a *consumerInstallationsImpl) ListListingInstallationsAll(ctx context.Context, request ListInstallationsRequest) ([]InstallationDetail, error) {
	iterator := a.ListListingInstallations(ctx, request)
	return listing.ToSlice[InstallationDetail](ctx, iterator)
}

func (a *consumerInstallationsImpl) internalListListingInstallations(ctx context.Context, request ListInstallationsRequest) (*ListInstallationsResponse, error) {
	requestPb, pbErr := ListInstallationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listInstallationsResponsePb marketplacepb.ListInstallationsResponsePb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listInstallationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListInstallationsResponseFromPb(&listInstallationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *consumerInstallationsImpl) Update(ctx context.Context, request UpdateInstallationRequest) (*UpdateInstallationResponse, error) {
	requestPb, pbErr := UpdateInstallationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateInstallationResponsePb marketplacepb.UpdateInstallationResponsePb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations/%v", request.ListingId, request.InstallationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&updateInstallationResponsePb,
	)
	resp, err := UpdateInstallationResponseFromPb(&updateInstallationResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ConsumerListings API methods
type consumerListingsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerListingsImpl) BatchGet(ctx context.Context, request BatchGetListingsRequest) (*BatchGetListingsResponse, error) {
	requestPb, pbErr := BatchGetListingsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var batchGetListingsResponsePb marketplacepb.BatchGetListingsResponsePb
	path := "/api/2.1/marketplace-consumer/listings:batchGet"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&batchGetListingsResponsePb,
	)
	resp, err := BatchGetListingsResponseFromPb(&batchGetListingsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *consumerListingsImpl) Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error) {
	requestPb, pbErr := GetListingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getListingResponsePb marketplacepb.GetListingResponsePb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getListingResponsePb,
	)
	resp, err := GetListingResponseFromPb(&getListingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List all published listings in the Databricks Marketplace that the consumer
// has access to.
func (a *consumerListingsImpl) List(ctx context.Context, request ListListingsRequest) listing.Iterator[Listing] {

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

// List all published listings in the Databricks Marketplace that the consumer
// has access to.
func (a *consumerListingsImpl) ListAll(ctx context.Context, request ListListingsRequest) ([]Listing, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}

func (a *consumerListingsImpl) internalList(ctx context.Context, request ListListingsRequest) (*ListListingsResponse, error) {
	requestPb, pbErr := ListListingsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listListingsResponsePb marketplacepb.ListListingsResponsePb
	path := "/api/2.1/marketplace-consumer/listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listListingsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListListingsResponseFromPb(&listListingsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Search published listings in the Databricks Marketplace that the consumer has
// access to. This query supports a variety of different search parameters and
// performs fuzzy matching.
func (a *consumerListingsImpl) Search(ctx context.Context, request SearchListingsRequest) listing.Iterator[Listing] {

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

// Search published listings in the Databricks Marketplace that the consumer has
// access to. This query supports a variety of different search parameters and
// performs fuzzy matching.
func (a *consumerListingsImpl) SearchAll(ctx context.Context, request SearchListingsRequest) ([]Listing, error) {
	iterator := a.Search(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}

func (a *consumerListingsImpl) internalSearch(ctx context.Context, request SearchListingsRequest) (*SearchListingsResponse, error) {
	requestPb, pbErr := SearchListingsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchListingsResponsePb marketplacepb.SearchListingsResponsePb
	path := "/api/2.1/marketplace-consumer/search-listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&searchListingsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := SearchListingsResponseFromPb(&searchListingsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ConsumerPersonalizationRequests API methods
type consumerPersonalizationRequestsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerPersonalizationRequestsImpl) Create(ctx context.Context, request CreatePersonalizationRequest) (*CreatePersonalizationRequestResponse, error) {
	requestPb, pbErr := CreatePersonalizationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createPersonalizationRequestResponsePb marketplacepb.CreatePersonalizationRequestResponsePb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/personalization-requests", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createPersonalizationRequestResponsePb,
	)
	resp, err := CreatePersonalizationRequestResponseFromPb(&createPersonalizationRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *consumerPersonalizationRequestsImpl) Get(ctx context.Context, request GetPersonalizationRequestRequest) (*GetPersonalizationRequestResponse, error) {
	requestPb, pbErr := GetPersonalizationRequestRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPersonalizationRequestResponsePb marketplacepb.GetPersonalizationRequestResponsePb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/personalization-requests", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getPersonalizationRequestResponsePb,
	)
	resp, err := GetPersonalizationRequestResponseFromPb(&getPersonalizationRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List personalization requests for a consumer across all listings.
func (a *consumerPersonalizationRequestsImpl) List(ctx context.Context, request ListAllPersonalizationRequestsRequest) listing.Iterator[PersonalizationRequest] {

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

// List personalization requests for a consumer across all listings.
func (a *consumerPersonalizationRequestsImpl) ListAll(ctx context.Context, request ListAllPersonalizationRequestsRequest) ([]PersonalizationRequest, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PersonalizationRequest](ctx, iterator)
}

func (a *consumerPersonalizationRequestsImpl) internalList(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
	requestPb, pbErr := ListAllPersonalizationRequestsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAllPersonalizationRequestsResponsePb marketplacepb.ListAllPersonalizationRequestsResponsePb
	path := "/api/2.1/marketplace-consumer/personalization-requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listAllPersonalizationRequestsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListAllPersonalizationRequestsResponseFromPb(&listAllPersonalizationRequestsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ConsumerProviders API methods
type consumerProvidersImpl struct {
	client *client.DatabricksClient
}

func (a *consumerProvidersImpl) BatchGet(ctx context.Context, request BatchGetProvidersRequest) (*BatchGetProvidersResponse, error) {
	requestPb, pbErr := BatchGetProvidersRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var batchGetProvidersResponsePb marketplacepb.BatchGetProvidersResponsePb
	path := "/api/2.1/marketplace-consumer/providers:batchGet"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&batchGetProvidersResponsePb,
	)
	resp, err := BatchGetProvidersResponseFromPb(&batchGetProvidersResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *consumerProvidersImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	requestPb, pbErr := GetProviderRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getProviderResponsePb marketplacepb.GetProviderResponsePb
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getProviderResponsePb,
	)
	resp, err := GetProviderResponseFromPb(&getProviderResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List all providers in the Databricks Marketplace with at least one visible
// listing.
func (a *consumerProvidersImpl) List(ctx context.Context, request ListProvidersRequest) listing.Iterator[ProviderInfo] {

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

// List all providers in the Databricks Marketplace with at least one visible
// listing.
func (a *consumerProvidersImpl) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}

func (a *consumerProvidersImpl) internalList(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	requestPb, pbErr := ListProvidersRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listProvidersResponsePb marketplacepb.ListProvidersResponsePb
	path := "/api/2.1/marketplace-consumer/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listProvidersResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListProvidersResponseFromPb(&listProvidersResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ProviderExchangeFilters API methods
type providerExchangeFiltersImpl struct {
	client *client.DatabricksClient
}

func (a *providerExchangeFiltersImpl) Create(ctx context.Context, request CreateExchangeFilterRequest) (*CreateExchangeFilterResponse, error) {
	requestPb, pbErr := CreateExchangeFilterRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createExchangeFilterResponsePb marketplacepb.CreateExchangeFilterResponsePb
	path := "/api/2.0/marketplace-exchange/filters"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createExchangeFilterResponsePb,
	)
	resp, err := CreateExchangeFilterResponseFromPb(&createExchangeFilterResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerExchangeFiltersImpl) Delete(ctx context.Context, request DeleteExchangeFilterRequest) error {
	requestPb, pbErr := DeleteExchangeFilterRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/marketplace-exchange/filters/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// List exchange filter
func (a *providerExchangeFiltersImpl) List(ctx context.Context, request ListExchangeFiltersRequest) listing.Iterator[ExchangeFilter] {

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

// List exchange filter
func (a *providerExchangeFiltersImpl) ListAll(ctx context.Context, request ListExchangeFiltersRequest) ([]ExchangeFilter, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ExchangeFilter](ctx, iterator)
}

func (a *providerExchangeFiltersImpl) internalList(ctx context.Context, request ListExchangeFiltersRequest) (*ListExchangeFiltersResponse, error) {
	requestPb, pbErr := ListExchangeFiltersRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listExchangeFiltersResponsePb marketplacepb.ListExchangeFiltersResponsePb
	path := "/api/2.0/marketplace-exchange/filters"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listExchangeFiltersResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListExchangeFiltersResponseFromPb(&listExchangeFiltersResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerExchangeFiltersImpl) Update(ctx context.Context, request UpdateExchangeFilterRequest) (*UpdateExchangeFilterResponse, error) {
	requestPb, pbErr := UpdateExchangeFilterRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateExchangeFilterResponsePb marketplacepb.UpdateExchangeFilterResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/filters/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&updateExchangeFilterResponsePb,
	)
	resp, err := UpdateExchangeFilterResponseFromPb(&updateExchangeFilterResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ProviderExchanges API methods
type providerExchangesImpl struct {
	client *client.DatabricksClient
}

func (a *providerExchangesImpl) AddListingToExchange(ctx context.Context, request AddExchangeForListingRequest) (*AddExchangeForListingResponse, error) {
	requestPb, pbErr := AddExchangeForListingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var addExchangeForListingResponsePb marketplacepb.AddExchangeForListingResponsePb
	path := "/api/2.0/marketplace-exchange/exchanges-for-listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&addExchangeForListingResponsePb,
	)
	resp, err := AddExchangeForListingResponseFromPb(&addExchangeForListingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerExchangesImpl) Create(ctx context.Context, request CreateExchangeRequest) (*CreateExchangeResponse, error) {
	requestPb, pbErr := CreateExchangeRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createExchangeResponsePb marketplacepb.CreateExchangeResponsePb
	path := "/api/2.0/marketplace-exchange/exchanges"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createExchangeResponsePb,
	)
	resp, err := CreateExchangeResponseFromPb(&createExchangeResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerExchangesImpl) Delete(ctx context.Context, request DeleteExchangeRequest) error {
	requestPb, pbErr := DeleteExchangeRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *providerExchangesImpl) DeleteListingFromExchange(ctx context.Context, request RemoveExchangeForListingRequest) error {
	requestPb, pbErr := RemoveExchangeForListingRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges-for-listing/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *providerExchangesImpl) Get(ctx context.Context, request GetExchangeRequest) (*GetExchangeResponse, error) {
	requestPb, pbErr := GetExchangeRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getExchangeResponsePb marketplacepb.GetExchangeResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getExchangeResponsePb,
	)
	resp, err := GetExchangeResponseFromPb(&getExchangeResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List exchanges visible to provider
func (a *providerExchangesImpl) List(ctx context.Context, request ListExchangesRequest) listing.Iterator[Exchange] {

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

// List exchanges visible to provider
func (a *providerExchangesImpl) ListAll(ctx context.Context, request ListExchangesRequest) ([]Exchange, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Exchange](ctx, iterator)
}

func (a *providerExchangesImpl) internalList(ctx context.Context, request ListExchangesRequest) (*ListExchangesResponse, error) {
	requestPb, pbErr := ListExchangesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listExchangesResponsePb marketplacepb.ListExchangesResponsePb
	path := "/api/2.0/marketplace-exchange/exchanges"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listExchangesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListExchangesResponseFromPb(&listExchangesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List exchanges associated with a listing
func (a *providerExchangesImpl) ListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) listing.Iterator[ExchangeListing] {

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

// List exchanges associated with a listing
func (a *providerExchangesImpl) ListExchangesForListingAll(ctx context.Context, request ListExchangesForListingRequest) ([]ExchangeListing, error) {
	iterator := a.ListExchangesForListing(ctx, request)
	return listing.ToSlice[ExchangeListing](ctx, iterator)
}

func (a *providerExchangesImpl) internalListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) (*ListExchangesForListingResponse, error) {
	requestPb, pbErr := ListExchangesForListingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listExchangesForListingResponsePb marketplacepb.ListExchangesForListingResponsePb
	path := "/api/2.0/marketplace-exchange/exchanges-for-listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listExchangesForListingResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListExchangesForListingResponseFromPb(&listExchangesForListingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List listings associated with an exchange
func (a *providerExchangesImpl) ListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) listing.Iterator[ExchangeListing] {

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

// List listings associated with an exchange
func (a *providerExchangesImpl) ListListingsForExchangeAll(ctx context.Context, request ListListingsForExchangeRequest) ([]ExchangeListing, error) {
	iterator := a.ListListingsForExchange(ctx, request)
	return listing.ToSlice[ExchangeListing](ctx, iterator)
}

func (a *providerExchangesImpl) internalListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) (*ListListingsForExchangeResponse, error) {
	requestPb, pbErr := ListListingsForExchangeRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listListingsForExchangeResponsePb marketplacepb.ListListingsForExchangeResponsePb
	path := "/api/2.0/marketplace-exchange/listings-for-exchange"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listListingsForExchangeResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListListingsForExchangeResponseFromPb(&listListingsForExchangeResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerExchangesImpl) Update(ctx context.Context, request UpdateExchangeRequest) (*UpdateExchangeResponse, error) {
	requestPb, pbErr := UpdateExchangeRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateExchangeResponsePb marketplacepb.UpdateExchangeResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&updateExchangeResponsePb,
	)
	resp, err := UpdateExchangeResponseFromPb(&updateExchangeResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ProviderFiles API methods
type providerFilesImpl struct {
	client *client.DatabricksClient
}

func (a *providerFilesImpl) Create(ctx context.Context, request CreateFileRequest) (*CreateFileResponse, error) {
	requestPb, pbErr := CreateFileRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createFileResponsePb marketplacepb.CreateFileResponsePb
	path := "/api/2.0/marketplace-provider/files"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createFileResponsePb,
	)
	resp, err := CreateFileResponseFromPb(&createFileResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerFilesImpl) Delete(ctx context.Context, request DeleteFileRequest) error {
	requestPb, pbErr := DeleteFileRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/marketplace-provider/files/%v", request.FileId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *providerFilesImpl) Get(ctx context.Context, request GetFileRequest) (*GetFileResponse, error) {
	requestPb, pbErr := GetFileRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getFileResponsePb marketplacepb.GetFileResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-provider/files/%v", request.FileId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getFileResponsePb,
	)
	resp, err := GetFileResponseFromPb(&getFileResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List files attached to a parent entity.
func (a *providerFilesImpl) List(ctx context.Context, request ListFilesRequest) listing.Iterator[FileInfo] {

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

// List files attached to a parent entity.
func (a *providerFilesImpl) ListAll(ctx context.Context, request ListFilesRequest) ([]FileInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FileInfo](ctx, iterator)
}

func (a *providerFilesImpl) internalList(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error) {
	requestPb, pbErr := ListFilesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listFilesResponsePb marketplacepb.ListFilesResponsePb
	path := "/api/2.0/marketplace-provider/files"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listFilesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListFilesResponseFromPb(&listFilesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ProviderListings API methods
type providerListingsImpl struct {
	client *client.DatabricksClient
}

func (a *providerListingsImpl) Create(ctx context.Context, request CreateListingRequest) (*CreateListingResponse, error) {
	requestPb, pbErr := CreateListingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createListingResponsePb marketplacepb.CreateListingResponsePb
	path := "/api/2.0/marketplace-provider/listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createListingResponsePb,
	)
	resp, err := CreateListingResponseFromPb(&createListingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerListingsImpl) Delete(ctx context.Context, request DeleteListingRequest) error {
	requestPb, pbErr := DeleteListingRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *providerListingsImpl) Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error) {
	requestPb, pbErr := GetListingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getListingResponsePb marketplacepb.GetListingResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getListingResponsePb,
	)
	resp, err := GetListingResponseFromPb(&getListingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List listings owned by this provider
func (a *providerListingsImpl) List(ctx context.Context, request GetListingsRequest) listing.Iterator[Listing] {

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

// List listings owned by this provider
func (a *providerListingsImpl) ListAll(ctx context.Context, request GetListingsRequest) ([]Listing, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}

func (a *providerListingsImpl) internalList(ctx context.Context, request GetListingsRequest) (*GetListingsResponse, error) {
	requestPb, pbErr := GetListingsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getListingsResponsePb marketplacepb.GetListingsResponsePb
	path := "/api/2.0/marketplace-provider/listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getListingsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := GetListingsResponseFromPb(&getListingsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerListingsImpl) Update(ctx context.Context, request UpdateListingRequest) (*UpdateListingResponse, error) {
	requestPb, pbErr := UpdateListingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateListingResponsePb marketplacepb.UpdateListingResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&updateListingResponsePb,
	)
	resp, err := UpdateListingResponseFromPb(&updateListingResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ProviderPersonalizationRequests API methods
type providerPersonalizationRequestsImpl struct {
	client *client.DatabricksClient
}

// List personalization requests to this provider. This will return all
// personalization requests, regardless of which listing they are for.
func (a *providerPersonalizationRequestsImpl) List(ctx context.Context, request ListAllPersonalizationRequestsRequest) listing.Iterator[PersonalizationRequest] {

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

// List personalization requests to this provider. This will return all
// personalization requests, regardless of which listing they are for.
func (a *providerPersonalizationRequestsImpl) ListAll(ctx context.Context, request ListAllPersonalizationRequestsRequest) ([]PersonalizationRequest, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PersonalizationRequest](ctx, iterator)
}

func (a *providerPersonalizationRequestsImpl) internalList(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
	requestPb, pbErr := ListAllPersonalizationRequestsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAllPersonalizationRequestsResponsePb marketplacepb.ListAllPersonalizationRequestsResponsePb
	path := "/api/2.0/marketplace-provider/personalization-requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listAllPersonalizationRequestsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListAllPersonalizationRequestsResponseFromPb(&listAllPersonalizationRequestsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerPersonalizationRequestsImpl) Update(ctx context.Context, request UpdatePersonalizationRequestRequest) (*UpdatePersonalizationRequestResponse, error) {
	requestPb, pbErr := UpdatePersonalizationRequestRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updatePersonalizationRequestResponsePb marketplacepb.UpdatePersonalizationRequestResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v/personalization-requests/%v/request-status", request.ListingId, request.RequestId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&updatePersonalizationRequestResponsePb,
	)
	resp, err := UpdatePersonalizationRequestResponseFromPb(&updatePersonalizationRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ProviderProviderAnalyticsDashboards API methods
type providerProviderAnalyticsDashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *providerProviderAnalyticsDashboardsImpl) Create(ctx context.Context) (*ProviderAnalyticsDashboard, error) {

	var providerAnalyticsDashboardPb marketplacepb.ProviderAnalyticsDashboardPb
	path := "/api/2.0/marketplace-provider/analytics_dashboard"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		nil,
		nil,
		&providerAnalyticsDashboardPb,
	)
	resp, err := ProviderAnalyticsDashboardFromPb(&providerAnalyticsDashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerProviderAnalyticsDashboardsImpl) Get(ctx context.Context) (*ListProviderAnalyticsDashboardResponse, error) {

	var listProviderAnalyticsDashboardResponsePb marketplacepb.ListProviderAnalyticsDashboardResponsePb
	path := "/api/2.0/marketplace-provider/analytics_dashboard"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&listProviderAnalyticsDashboardResponsePb,
	)
	resp, err := ListProviderAnalyticsDashboardResponseFromPb(&listProviderAnalyticsDashboardResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerProviderAnalyticsDashboardsImpl) GetLatestVersion(ctx context.Context) (*GetLatestVersionProviderAnalyticsDashboardResponse, error) {

	var getLatestVersionProviderAnalyticsDashboardResponsePb marketplacepb.GetLatestVersionProviderAnalyticsDashboardResponsePb
	path := "/api/2.0/marketplace-provider/analytics_dashboard/latest"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&getLatestVersionProviderAnalyticsDashboardResponsePb,
	)
	resp, err := GetLatestVersionProviderAnalyticsDashboardResponseFromPb(&getLatestVersionProviderAnalyticsDashboardResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerProviderAnalyticsDashboardsImpl) Update(ctx context.Context, request UpdateProviderAnalyticsDashboardRequest) (*UpdateProviderAnalyticsDashboardResponse, error) {
	requestPb, pbErr := UpdateProviderAnalyticsDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateProviderAnalyticsDashboardResponsePb marketplacepb.UpdateProviderAnalyticsDashboardResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-provider/analytics_dashboard/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&updateProviderAnalyticsDashboardResponsePb,
	)
	resp, err := UpdateProviderAnalyticsDashboardResponseFromPb(&updateProviderAnalyticsDashboardResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ProviderProviders API methods
type providerProvidersImpl struct {
	client *client.DatabricksClient
}

func (a *providerProvidersImpl) Create(ctx context.Context, request CreateProviderRequest) (*CreateProviderResponse, error) {
	requestPb, pbErr := CreateProviderRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createProviderResponsePb marketplacepb.CreateProviderResponsePb
	path := "/api/2.0/marketplace-provider/provider"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createProviderResponsePb,
	)
	resp, err := CreateProviderResponseFromPb(&createProviderResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerProvidersImpl) Delete(ctx context.Context, request DeleteProviderRequest) error {
	requestPb, pbErr := DeleteProviderRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *providerProvidersImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	requestPb, pbErr := GetProviderRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getProviderResponsePb marketplacepb.GetProviderResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getProviderResponsePb,
	)
	resp, err := GetProviderResponseFromPb(&getProviderResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List provider profiles for account.
func (a *providerProvidersImpl) List(ctx context.Context, request ListProvidersRequest) listing.Iterator[ProviderInfo] {

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

// List provider profiles for account.
func (a *providerProvidersImpl) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}

func (a *providerProvidersImpl) internalList(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	requestPb, pbErr := ListProvidersRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listProvidersResponsePb marketplacepb.ListProvidersResponsePb
	path := "/api/2.0/marketplace-provider/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listProvidersResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListProvidersResponseFromPb(&listProvidersResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providerProvidersImpl) Update(ctx context.Context, request UpdateProviderRequest) (*UpdateProviderResponse, error) {
	requestPb, pbErr := UpdateProviderRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateProviderResponsePb marketplacepb.UpdateProviderResponsePb
	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&updateProviderResponsePb,
	)
	resp, err := UpdateProviderResponseFromPb(&updateProviderResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
