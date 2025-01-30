// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just ConsumerFulfillments API methods
type consumerFulfillmentsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerFulfillmentsImpl) Get(ctx context.Context, request GetListingContentMetadataRequest) (*GetListingContentMetadataResponse, error) {
	var getListingContentMetadataResponse GetListingContentMetadataResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/content", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingContentMetadataResponse)
	return &getListingContentMetadataResponse, err
}

func (a *consumerFulfillmentsImpl) List(ctx context.Context, request ListFulfillmentsRequest) (*ListFulfillmentsResponse, error) {
	var listFulfillmentsResponse ListFulfillmentsResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/fulfillments", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFulfillmentsResponse)
	return &listFulfillmentsResponse, err
}

// unexported type that holds implementations of just ConsumerInstallations API methods
type consumerInstallationsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerInstallationsImpl) Create(ctx context.Context, request CreateInstallationRequest) (*Installation, error) {
	var installation Installation
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &installation)
	return &installation, err
}

func (a *consumerInstallationsImpl) Delete(ctx context.Context, request DeleteInstallationRequest) error {
	var deleteInstallationResponse DeleteInstallationResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations/%v", request.ListingId, request.InstallationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteInstallationResponse)
	return err
}

func (a *consumerInstallationsImpl) List(ctx context.Context, request ListAllInstallationsRequest) (*ListAllInstallationsResponse, error) {
	var listAllInstallationsResponse ListAllInstallationsResponse
	path := "/api/2.1/marketplace-consumer/installations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllInstallationsResponse)
	return &listAllInstallationsResponse, err
}

func (a *consumerInstallationsImpl) ListListingInstallations(ctx context.Context, request ListInstallationsRequest) (*ListInstallationsResponse, error) {
	var listInstallationsResponse ListInstallationsResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listInstallationsResponse)
	return &listInstallationsResponse, err
}

func (a *consumerInstallationsImpl) Update(ctx context.Context, request UpdateInstallationRequest) (*UpdateInstallationResponse, error) {
	var updateInstallationResponse UpdateInstallationResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations/%v", request.ListingId, request.InstallationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateInstallationResponse)
	return &updateInstallationResponse, err
}

// unexported type that holds implementations of just ConsumerListings API methods
type consumerListingsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerListingsImpl) BatchGet(ctx context.Context, request BatchGetListingsRequest) (*BatchGetListingsResponse, error) {
	var batchGetListingsResponse BatchGetListingsResponse
	path := "/api/2.1/marketplace-consumer/listings:batchGet"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &batchGetListingsResponse)
	return &batchGetListingsResponse, err
}

func (a *consumerListingsImpl) Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error) {
	var getListingResponse GetListingResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingResponse)
	return &getListingResponse, err
}

func (a *consumerListingsImpl) List(ctx context.Context, request ListListingsRequest) (*ListListingsResponse, error) {
	var listListingsResponse ListListingsResponse
	path := "/api/2.1/marketplace-consumer/listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listListingsResponse)
	return &listListingsResponse, err
}

func (a *consumerListingsImpl) Search(ctx context.Context, request SearchListingsRequest) (*SearchListingsResponse, error) {
	var searchListingsResponse SearchListingsResponse
	path := "/api/2.1/marketplace-consumer/search-listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &searchListingsResponse)
	return &searchListingsResponse, err
}

// unexported type that holds implementations of just ConsumerPersonalizationRequests API methods
type consumerPersonalizationRequestsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerPersonalizationRequestsImpl) Create(ctx context.Context, request CreatePersonalizationRequest) (*CreatePersonalizationRequestResponse, error) {
	var createPersonalizationRequestResponse CreatePersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/personalization-requests", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createPersonalizationRequestResponse)
	return &createPersonalizationRequestResponse, err
}

func (a *consumerPersonalizationRequestsImpl) Get(ctx context.Context, request GetPersonalizationRequestRequest) (*GetPersonalizationRequestResponse, error) {
	var getPersonalizationRequestResponse GetPersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/personalization-requests", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPersonalizationRequestResponse)
	return &getPersonalizationRequestResponse, err
}

func (a *consumerPersonalizationRequestsImpl) List(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
	var listAllPersonalizationRequestsResponse ListAllPersonalizationRequestsResponse
	path := "/api/2.1/marketplace-consumer/personalization-requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllPersonalizationRequestsResponse)
	return &listAllPersonalizationRequestsResponse, err
}

// unexported type that holds implementations of just ConsumerProviders API methods
type consumerProvidersImpl struct {
	client *client.DatabricksClient
}

func (a *consumerProvidersImpl) BatchGet(ctx context.Context, request BatchGetProvidersRequest) (*BatchGetProvidersResponse, error) {
	var batchGetProvidersResponse BatchGetProvidersResponse
	path := "/api/2.1/marketplace-consumer/providers:batchGet"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &batchGetProvidersResponse)
	return &batchGetProvidersResponse, err
}

func (a *consumerProvidersImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getProviderResponse)
	return &getProviderResponse, err
}

func (a *consumerProvidersImpl) List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/marketplace-consumer/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProvidersResponse)
	return &listProvidersResponse, err
}

// unexported type that holds implementations of just ProviderExchangeFilters API methods
type providerExchangeFiltersImpl struct {
	client *client.DatabricksClient
}

func (a *providerExchangeFiltersImpl) Create(ctx context.Context, request CreateExchangeFilterRequest) (*CreateExchangeFilterResponse, error) {
	var createExchangeFilterResponse CreateExchangeFilterResponse
	path := "/api/2.0/marketplace-exchange/filters"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createExchangeFilterResponse)
	return &createExchangeFilterResponse, err
}

func (a *providerExchangeFiltersImpl) Delete(ctx context.Context, request DeleteExchangeFilterRequest) error {
	var deleteExchangeFilterResponse DeleteExchangeFilterResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/filters/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteExchangeFilterResponse)
	return err
}

func (a *providerExchangeFiltersImpl) List(ctx context.Context, request ListExchangeFiltersRequest) (*ListExchangeFiltersResponse, error) {
	var listExchangeFiltersResponse ListExchangeFiltersResponse
	path := "/api/2.0/marketplace-exchange/filters"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangeFiltersResponse)
	return &listExchangeFiltersResponse, err
}

func (a *providerExchangeFiltersImpl) Update(ctx context.Context, request UpdateExchangeFilterRequest) (*UpdateExchangeFilterResponse, error) {
	var updateExchangeFilterResponse UpdateExchangeFilterResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/filters/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateExchangeFilterResponse)
	return &updateExchangeFilterResponse, err
}

// unexported type that holds implementations of just ProviderExchanges API methods
type providerExchangesImpl struct {
	client *client.DatabricksClient
}

func (a *providerExchangesImpl) AddListingToExchange(ctx context.Context, request AddExchangeForListingRequest) (*AddExchangeForListingResponse, error) {
	var addExchangeForListingResponse AddExchangeForListingResponse
	path := "/api/2.0/marketplace-exchange/exchanges-for-listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &addExchangeForListingResponse)
	return &addExchangeForListingResponse, err
}

func (a *providerExchangesImpl) Create(ctx context.Context, request CreateExchangeRequest) (*CreateExchangeResponse, error) {
	var createExchangeResponse CreateExchangeResponse
	path := "/api/2.0/marketplace-exchange/exchanges"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createExchangeResponse)
	return &createExchangeResponse, err
}

func (a *providerExchangesImpl) Delete(ctx context.Context, request DeleteExchangeRequest) error {
	var deleteExchangeResponse DeleteExchangeResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteExchangeResponse)
	return err
}

func (a *providerExchangesImpl) DeleteListingFromExchange(ctx context.Context, request RemoveExchangeForListingRequest) error {
	var removeExchangeForListingResponse RemoveExchangeForListingResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges-for-listing/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &removeExchangeForListingResponse)
	return err
}

func (a *providerExchangesImpl) Get(ctx context.Context, request GetExchangeRequest) (*GetExchangeResponse, error) {
	var getExchangeResponse GetExchangeResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExchangeResponse)
	return &getExchangeResponse, err
}

func (a *providerExchangesImpl) List(ctx context.Context, request ListExchangesRequest) (*ListExchangesResponse, error) {
	var listExchangesResponse ListExchangesResponse
	path := "/api/2.0/marketplace-exchange/exchanges"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangesResponse)
	return &listExchangesResponse, err
}

func (a *providerExchangesImpl) ListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) (*ListExchangesForListingResponse, error) {
	var listExchangesForListingResponse ListExchangesForListingResponse
	path := "/api/2.0/marketplace-exchange/exchanges-for-listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangesForListingResponse)
	return &listExchangesForListingResponse, err
}

func (a *providerExchangesImpl) ListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) (*ListListingsForExchangeResponse, error) {
	var listListingsForExchangeResponse ListListingsForExchangeResponse
	path := "/api/2.0/marketplace-exchange/listings-for-exchange"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listListingsForExchangeResponse)
	return &listListingsForExchangeResponse, err
}

func (a *providerExchangesImpl) Update(ctx context.Context, request UpdateExchangeRequest) (*UpdateExchangeResponse, error) {
	var updateExchangeResponse UpdateExchangeResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateExchangeResponse)
	return &updateExchangeResponse, err
}

// unexported type that holds implementations of just ProviderFiles API methods
type providerFilesImpl struct {
	client *client.DatabricksClient
}

func (a *providerFilesImpl) Create(ctx context.Context, request CreateFileRequest) (*CreateFileResponse, error) {
	var createFileResponse CreateFileResponse
	path := "/api/2.0/marketplace-provider/files"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createFileResponse)
	return &createFileResponse, err
}

func (a *providerFilesImpl) Delete(ctx context.Context, request DeleteFileRequest) error {
	var deleteFileResponse DeleteFileResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/files/%v", request.FileId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteFileResponse)
	return err
}

func (a *providerFilesImpl) Get(ctx context.Context, request GetFileRequest) (*GetFileResponse, error) {
	var getFileResponse GetFileResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/files/%v", request.FileId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getFileResponse)
	return &getFileResponse, err
}

func (a *providerFilesImpl) List(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error) {
	var listFilesResponse ListFilesResponse
	path := "/api/2.0/marketplace-provider/files"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFilesResponse)
	return &listFilesResponse, err
}

// unexported type that holds implementations of just ProviderListings API methods
type providerListingsImpl struct {
	client *client.DatabricksClient
}

func (a *providerListingsImpl) Create(ctx context.Context, request CreateListingRequest) (*CreateListingResponse, error) {
	var createListingResponse CreateListingResponse
	path := "/api/2.0/marketplace-provider/listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createListingResponse)
	return &createListingResponse, err
}

func (a *providerListingsImpl) Delete(ctx context.Context, request DeleteListingRequest) error {
	var deleteListingResponse DeleteListingResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteListingResponse)
	return err
}

func (a *providerListingsImpl) Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error) {
	var getListingResponse GetListingResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingResponse)
	return &getListingResponse, err
}

func (a *providerListingsImpl) List(ctx context.Context, request GetListingsRequest) (*GetListingsResponse, error) {
	var getListingsResponse GetListingsResponse
	path := "/api/2.0/marketplace-provider/listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingsResponse)
	return &getListingsResponse, err
}

func (a *providerListingsImpl) Update(ctx context.Context, request UpdateListingRequest) (*UpdateListingResponse, error) {
	var updateListingResponse UpdateListingResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateListingResponse)
	return &updateListingResponse, err
}

// unexported type that holds implementations of just ProviderPersonalizationRequests API methods
type providerPersonalizationRequestsImpl struct {
	client *client.DatabricksClient
}

func (a *providerPersonalizationRequestsImpl) List(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
	var listAllPersonalizationRequestsResponse ListAllPersonalizationRequestsResponse
	path := "/api/2.0/marketplace-provider/personalization-requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllPersonalizationRequestsResponse)
	return &listAllPersonalizationRequestsResponse, err
}

func (a *providerPersonalizationRequestsImpl) Update(ctx context.Context, request UpdatePersonalizationRequestRequest) (*UpdatePersonalizationRequestResponse, error) {
	var updatePersonalizationRequestResponse UpdatePersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v/personalization-requests/%v/request-status", request.ListingId, request.RequestId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updatePersonalizationRequestResponse)
	return &updatePersonalizationRequestResponse, err
}

// unexported type that holds implementations of just ProviderProviderAnalyticsDashboards API methods
type providerProviderAnalyticsDashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *providerProviderAnalyticsDashboardsImpl) Create(ctx context.Context) (*ProviderAnalyticsDashboard, error) {
	var providerAnalyticsDashboard ProviderAnalyticsDashboard
	path := "/api/2.0/marketplace-provider/analytics_dashboard"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, nil, &providerAnalyticsDashboard)
	return &providerAnalyticsDashboard, err
}

func (a *providerProviderAnalyticsDashboardsImpl) Get(ctx context.Context) (*ListProviderAnalyticsDashboardResponse, error) {
	var listProviderAnalyticsDashboardResponse ListProviderAnalyticsDashboardResponse
	path := "/api/2.0/marketplace-provider/analytics_dashboard"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listProviderAnalyticsDashboardResponse)
	return &listProviderAnalyticsDashboardResponse, err
}

func (a *providerProviderAnalyticsDashboardsImpl) GetLatestVersion(ctx context.Context) (*GetLatestVersionProviderAnalyticsDashboardResponse, error) {
	var getLatestVersionProviderAnalyticsDashboardResponse GetLatestVersionProviderAnalyticsDashboardResponse
	path := "/api/2.0/marketplace-provider/analytics_dashboard/latest"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getLatestVersionProviderAnalyticsDashboardResponse)
	return &getLatestVersionProviderAnalyticsDashboardResponse, err
}

func (a *providerProviderAnalyticsDashboardsImpl) Update(ctx context.Context, request UpdateProviderAnalyticsDashboardRequest) (*UpdateProviderAnalyticsDashboardResponse, error) {
	var updateProviderAnalyticsDashboardResponse UpdateProviderAnalyticsDashboardResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/analytics_dashboard/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateProviderAnalyticsDashboardResponse)
	return &updateProviderAnalyticsDashboardResponse, err
}

// unexported type that holds implementations of just ProviderProviders API methods
type providerProvidersImpl struct {
	client *client.DatabricksClient
}

func (a *providerProvidersImpl) Create(ctx context.Context, request CreateProviderRequest) (*CreateProviderResponse, error) {
	var createProviderResponse CreateProviderResponse
	path := "/api/2.0/marketplace-provider/provider"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createProviderResponse)
	return &createProviderResponse, err
}

func (a *providerProvidersImpl) Delete(ctx context.Context, request DeleteProviderRequest) error {
	var deleteProviderResponse DeleteProviderResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteProviderResponse)
	return err
}

func (a *providerProvidersImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getProviderResponse)
	return &getProviderResponse, err
}

func (a *providerProvidersImpl) List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.0/marketplace-provider/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProvidersResponse)
	return &listProvidersResponse, err
}

func (a *providerProvidersImpl) Update(ctx context.Context, request UpdateProviderRequest) (*UpdateProviderResponse, error) {
	var updateProviderResponse UpdateProviderResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateProviderResponse)
	return &updateProviderResponse, err
}
