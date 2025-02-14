// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

// unexported type that holds implementations of just Providers API methods
type providersImpl struct {
	client *client.DatabricksClient
}

func (a *providersImpl) Create(ctx context.Context, request CreateProvider) (*ProviderInfo, error) {
	var providerInfo ProviderInfo
	path := "/api/2.1/unity-catalog/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &providerInfo)
	return &providerInfo, err
}

func (a *providersImpl) Delete(ctx context.Context, request DeleteProviderRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *providersImpl) Get(ctx context.Context, request GetProviderRequest) (*ProviderInfo, error) {
	var providerInfo ProviderInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &providerInfo)
	return &providerInfo, err
}

// List providers.
//
// Gets an array of available authentication providers. The caller must either
// be a metastore admin or the owner of the providers. Providers not owned by
// the caller are not included in the response. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *providersImpl) List(ctx context.Context, request ListProvidersRequest) listing.Iterator[ProviderInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

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
// Gets an array of available authentication providers. The caller must either
// be a metastore admin or the owner of the providers. Providers not owned by
// the caller are not included in the response. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *providersImpl) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}

func (a *providersImpl) internalList(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/unity-catalog/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProvidersResponse)
	return &listProvidersResponse, err
}

// List shares by Provider.
//
// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *providersImpl) ListShares(ctx context.Context, request ListSharesRequest) listing.Iterator[ProviderShare] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSharesRequest) (*ListProviderSharesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListShares(ctx, req)
	}
	getItems := func(resp *ListProviderSharesResponse) []ProviderShare {
		return resp.Shares
	}
	getNextReq := func(resp *ListProviderSharesResponse) *ListSharesRequest {
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

// List shares by Provider.
//
// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *providersImpl) ListSharesAll(ctx context.Context, request ListSharesRequest) ([]ProviderShare, error) {
	iterator := a.ListShares(ctx, request)
	return listing.ToSlice[ProviderShare](ctx, iterator)
}

func (a *providersImpl) internalListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error) {
	var listProviderSharesResponse ListProviderSharesResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v/shares", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProviderSharesResponse)
	return &listProviderSharesResponse, err
}

func (a *providersImpl) Update(ctx context.Context, request UpdateProvider) (*ProviderInfo, error) {
	var providerInfo ProviderInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &providerInfo)
	return &providerInfo, err
}

// unexported type that holds implementations of just RecipientActivation API methods
type recipientActivationImpl struct {
	client *client.DatabricksClient
}

func (a *recipientActivationImpl) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	var getActivationUrlInfoResponse GetActivationUrlInfoResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation_info/%v", request.ActivationUrl)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getActivationUrlInfoResponse)
	return err
}

func (a *recipientActivationImpl) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	var retrieveTokenResponse RetrieveTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation/%v", request.ActivationUrl)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &retrieveTokenResponse)
	return &retrieveTokenResponse, err
}

// unexported type that holds implementations of just Recipients API methods
type recipientsImpl struct {
	client *client.DatabricksClient
}

func (a *recipientsImpl) Create(ctx context.Context, request CreateRecipient) (*RecipientInfo, error) {
	var recipientInfo RecipientInfo
	path := "/api/2.1/unity-catalog/recipients"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &recipientInfo)
	return &recipientInfo, err
}

func (a *recipientsImpl) Delete(ctx context.Context, request DeleteRecipientRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *recipientsImpl) Get(ctx context.Context, request GetRecipientRequest) (*RecipientInfo, error) {
	var recipientInfo RecipientInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &recipientInfo)
	return &recipientInfo, err
}

// List share recipients.
//
// Gets an array of all share recipients within the current metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner. There is no
// guarantee of a specific ordering of the elements in the array.
func (a *recipientsImpl) List(ctx context.Context, request ListRecipientsRequest) listing.Iterator[RecipientInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListRecipientsRequest) (*ListRecipientsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListRecipientsResponse) []RecipientInfo {
		return resp.Recipients
	}
	getNextReq := func(resp *ListRecipientsResponse) *ListRecipientsRequest {
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

// List share recipients.
//
// Gets an array of all share recipients within the current metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner. There is no
// guarantee of a specific ordering of the elements in the array.
func (a *recipientsImpl) ListAll(ctx context.Context, request ListRecipientsRequest) ([]RecipientInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[RecipientInfo](ctx, iterator)
}

func (a *recipientsImpl) internalList(ctx context.Context, request ListRecipientsRequest) (*ListRecipientsResponse, error) {
	var listRecipientsResponse ListRecipientsResponse
	path := "/api/2.1/unity-catalog/recipients"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRecipientsResponse)
	return &listRecipientsResponse, err
}

func (a *recipientsImpl) RotateToken(ctx context.Context, request RotateRecipientToken) (*RecipientInfo, error) {
	var recipientInfo RecipientInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/rotate-token", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &recipientInfo)
	return &recipientInfo, err
}

func (a *recipientsImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	var getRecipientSharePermissionsResponse GetRecipientSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/share-permissions", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRecipientSharePermissionsResponse)
	return &getRecipientSharePermissionsResponse, err
}

func (a *recipientsImpl) Update(ctx context.Context, request UpdateRecipient) (*RecipientInfo, error) {
	var recipientInfo RecipientInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &recipientInfo)
	return &recipientInfo, err
}

// unexported type that holds implementations of just Shares API methods
type sharesImpl struct {
	client *client.DatabricksClient
}

func (a *sharesImpl) Create(ctx context.Context, request CreateShare) (*ShareInfo, error) {
	var shareInfo ShareInfo
	path := "/api/2.1/unity-catalog/shares"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &shareInfo)
	return &shareInfo, err
}

func (a *sharesImpl) Delete(ctx context.Context, request DeleteShareRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *sharesImpl) Get(ctx context.Context, request GetShareRequest) (*ShareInfo, error) {
	var shareInfo ShareInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &shareInfo)
	return &shareInfo, err
}

// List shares.
//
// Gets an array of data object shares from the metastore. The caller must be a
// metastore admin or the owner of the share. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *sharesImpl) List(ctx context.Context, request ListSharesRequest) listing.Iterator[ShareInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSharesRequest) (*ListSharesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListSharesResponse) []ShareInfo {
		return resp.Shares
	}
	getNextReq := func(resp *ListSharesResponse) *ListSharesRequest {
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

// List shares.
//
// Gets an array of data object shares from the metastore. The caller must be a
// metastore admin or the owner of the share. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *sharesImpl) ListAll(ctx context.Context, request ListSharesRequest) ([]ShareInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ShareInfo](ctx, iterator)
}

func (a *sharesImpl) internalList(ctx context.Context, request ListSharesRequest) (*ListSharesResponse, error) {
	var listSharesResponse ListSharesResponse
	path := "/api/2.1/unity-catalog/shares"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSharesResponse)
	return &listSharesResponse, err
}

func (a *sharesImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*catalog.PermissionsList, error) {
	var permissionsList catalog.PermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &permissionsList)
	return &permissionsList, err
}

func (a *sharesImpl) Update(ctx context.Context, request UpdateShare) (*ShareInfo, error) {
	var shareInfo ShareInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &shareInfo)
	return &shareInfo, err
}

func (a *sharesImpl) UpdatePermissions(ctx context.Context, request UpdateSharePermissions) error {
	var updatePermissionsResponse UpdatePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	queryParams := make(map[string]any)
	if request.MaxResults != 0 || slices.Contains(request.ForceSendFields, "MaxResults") {
		queryParams["max_results"] = request.MaxResults
	}
	if request.PageToken != "" || slices.Contains(request.ForceSendFields, "PageToken") {
		queryParams["page_token"] = request.PageToken
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updatePermissionsResponse)
	return err
}
