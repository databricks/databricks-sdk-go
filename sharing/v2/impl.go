// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"golang.org/x/exp/slices"
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

func (a *providersImpl) List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/unity-catalog/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProvidersResponse)
	return &listProvidersResponse, err
}

func (a *providersImpl) ListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error) {
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

func (a *recipientsImpl) List(ctx context.Context, request ListRecipientsRequest) (*ListRecipientsResponse, error) {
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

func (a *sharesImpl) List(ctx context.Context, request ListSharesRequest) (*ListSharesResponse, error) {
	var listSharesResponse ListSharesResponse
	path := "/api/2.1/unity-catalog/shares"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSharesResponse)
	return &listSharesResponse, err
}

func (a *sharesImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*PermissionsList, error) {
	var permissionsList PermissionsList
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
