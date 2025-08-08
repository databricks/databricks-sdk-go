// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/sharing/sharingpb"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Providers API methods
type providersImpl struct {
	client *client.DatabricksClient
}

func (a *providersImpl) Create(ctx context.Context, request CreateProvider) (*ProviderInfo, error) {
	requestPb, pbErr := CreateProviderToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var providerInfoPb sharingpb.ProviderInfoPb
	path := "/api/2.1/unity-catalog/providers"
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
		&providerInfoPb,
	)
	resp, err := ProviderInfoFromPb(&providerInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providersImpl) Delete(ctx context.Context, request DeleteProviderRequest) error {
	requestPb, pbErr := DeleteProviderRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *providersImpl) Get(ctx context.Context, request GetProviderRequest) (*ProviderInfo, error) {
	requestPb, pbErr := GetProviderRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var providerInfoPb sharingpb.ProviderInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&providerInfoPb,
	)
	resp, err := ProviderInfoFromPb(&providerInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets an array of available authentication providers. The caller must either
// be a metastore admin or the owner of the providers. Providers not owned by
// the caller are not included in the response. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *providersImpl) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}

func (a *providersImpl) internalList(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	requestPb, pbErr := ListProvidersRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listProvidersResponsePb sharingpb.ListProvidersResponsePb
	path := "/api/2.1/unity-catalog/providers"
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

func (a *providersImpl) ListProviderShareAssets(ctx context.Context, request ListProviderShareAssetsRequest) (*ListProviderShareAssetsResponse, error) {
	requestPb, pbErr := ListProviderShareAssetsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listProviderShareAssetsResponsePb sharingpb.ListProviderShareAssetsResponsePb
	path := fmt.Sprintf("/api/2.1/data-sharing/providers/%v/shares/%v", request.ProviderName, request.ShareName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listProviderShareAssetsResponsePb,
	)
	resp, err := ListProviderShareAssetsResponseFromPb(&listProviderShareAssetsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *providersImpl) ListSharesAll(ctx context.Context, request ListSharesRequest) ([]ProviderShare, error) {
	iterator := a.ListShares(ctx, request)
	return listing.ToSlice[ProviderShare](ctx, iterator)
}

func (a *providersImpl) internalListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error) {
	requestPb, pbErr := ListSharesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listProviderSharesResponsePb sharingpb.ListProviderSharesResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v/shares", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listProviderSharesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListProviderSharesResponseFromPb(&listProviderSharesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *providersImpl) Update(ctx context.Context, request UpdateProvider) (*ProviderInfo, error) {
	requestPb, pbErr := UpdateProviderToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var providerInfoPb sharingpb.ProviderInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&providerInfoPb,
	)
	resp, err := ProviderInfoFromPb(&providerInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just RecipientActivation API methods
type recipientActivationImpl struct {
	client *client.DatabricksClient
}

func (a *recipientActivationImpl) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	requestPb, pbErr := GetActivationUrlInfoRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation_info/%v", request.ActivationUrl)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *recipientActivationImpl) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	requestPb, pbErr := RetrieveTokenRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var retrieveTokenResponsePb sharingpb.RetrieveTokenResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation/%v", request.ActivationUrl)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&retrieveTokenResponsePb,
	)
	resp, err := RetrieveTokenResponseFromPb(&retrieveTokenResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just RecipientFederationPolicies API methods
type recipientFederationPoliciesImpl struct {
	client *client.DatabricksClient
}

func (a *recipientFederationPoliciesImpl) Create(ctx context.Context, request CreateFederationPolicyRequest) (*FederationPolicy, error) {
	requestPb, pbErr := CreateFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb sharingpb.FederationPolicyPb
	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies", request.RecipientName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.Policy,
		&federationPolicyPb,
	)
	resp, err := FederationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *recipientFederationPoliciesImpl) Delete(ctx context.Context, request DeleteFederationPolicyRequest) error {
	requestPb, pbErr := DeleteFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies/%v", request.RecipientName, request.Name)
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

func (a *recipientFederationPoliciesImpl) GetFederationPolicy(ctx context.Context, request GetFederationPolicyRequest) (*FederationPolicy, error) {
	requestPb, pbErr := GetFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb sharingpb.FederationPolicyPb
	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies/%v", request.RecipientName, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&federationPolicyPb,
	)
	resp, err := FederationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists federation policies for an OIDC_FEDERATION recipient for sharing data
// from Databricks to non-Databricks recipients. The caller must have read
// access to the recipient.
func (a *recipientFederationPoliciesImpl) List(ctx context.Context, request ListFederationPoliciesRequest) listing.Iterator[FederationPolicy] {

	getNextPage := func(ctx context.Context, req ListFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFederationPoliciesResponse) []FederationPolicy {
		return resp.Policies
	}
	getNextReq := func(resp *ListFederationPoliciesResponse) *ListFederationPoliciesRequest {
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

// Lists federation policies for an OIDC_FEDERATION recipient for sharing data
// from Databricks to non-Databricks recipients. The caller must have read
// access to the recipient.
func (a *recipientFederationPoliciesImpl) ListAll(ctx context.Context, request ListFederationPoliciesRequest) ([]FederationPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FederationPolicy](ctx, iterator)
}

func (a *recipientFederationPoliciesImpl) internalList(ctx context.Context, request ListFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
	requestPb, pbErr := ListFederationPoliciesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listFederationPoliciesResponsePb sharingpb.ListFederationPoliciesResponsePb
	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies", request.RecipientName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listFederationPoliciesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListFederationPoliciesResponseFromPb(&listFederationPoliciesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *recipientFederationPoliciesImpl) Update(ctx context.Context, request UpdateFederationPolicyRequest) (*FederationPolicy, error) {
	requestPb, pbErr := UpdateFederationPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var federationPolicyPb sharingpb.FederationPolicyPb
	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies/%v", request.RecipientName, request.Name)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" || slices.Contains(requestPb.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.Policy,
		&federationPolicyPb,
	)
	resp, err := FederationPolicyFromPb(&federationPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Recipients API methods
type recipientsImpl struct {
	client *client.DatabricksClient
}

func (a *recipientsImpl) Create(ctx context.Context, request CreateRecipient) (*RecipientInfo, error) {
	requestPb, pbErr := CreateRecipientToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var recipientInfoPb sharingpb.RecipientInfoPb
	path := "/api/2.1/unity-catalog/recipients"
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
		&recipientInfoPb,
	)
	resp, err := RecipientInfoFromPb(&recipientInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *recipientsImpl) Delete(ctx context.Context, request DeleteRecipientRequest) error {
	requestPb, pbErr := DeleteRecipientRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *recipientsImpl) Get(ctx context.Context, request GetRecipientRequest) (*RecipientInfo, error) {
	requestPb, pbErr := GetRecipientRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var recipientInfoPb sharingpb.RecipientInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&recipientInfoPb,
	)
	resp, err := RecipientInfoFromPb(&recipientInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets an array of all share recipients within the current metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner. There is no
// guarantee of a specific ordering of the elements in the array.
func (a *recipientsImpl) ListAll(ctx context.Context, request ListRecipientsRequest) ([]RecipientInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[RecipientInfo](ctx, iterator)
}

func (a *recipientsImpl) internalList(ctx context.Context, request ListRecipientsRequest) (*ListRecipientsResponse, error) {
	requestPb, pbErr := ListRecipientsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listRecipientsResponsePb sharingpb.ListRecipientsResponsePb
	path := "/api/2.1/unity-catalog/recipients"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listRecipientsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListRecipientsResponseFromPb(&listRecipientsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *recipientsImpl) RotateToken(ctx context.Context, request RotateRecipientToken) (*RecipientInfo, error) {
	requestPb, pbErr := RotateRecipientTokenToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var recipientInfoPb sharingpb.RecipientInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/rotate-token", request.Name)
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
		&recipientInfoPb,
	)
	resp, err := RecipientInfoFromPb(&recipientInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *recipientsImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	requestPb, pbErr := SharePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRecipientSharePermissionsResponsePb sharingpb.GetRecipientSharePermissionsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/share-permissions", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getRecipientSharePermissionsResponsePb,
	)
	resp, err := GetRecipientSharePermissionsResponseFromPb(&getRecipientSharePermissionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *recipientsImpl) Update(ctx context.Context, request UpdateRecipient) (*RecipientInfo, error) {
	requestPb, pbErr := UpdateRecipientToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var recipientInfoPb sharingpb.RecipientInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&recipientInfoPb,
	)
	resp, err := RecipientInfoFromPb(&recipientInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Shares API methods
type sharesImpl struct {
	client *client.DatabricksClient
}

func (a *sharesImpl) Create(ctx context.Context, request CreateShare) (*ShareInfo, error) {
	requestPb, pbErr := CreateShareToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var shareInfoPb sharingpb.ShareInfoPb
	path := "/api/2.1/unity-catalog/shares"
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
		&shareInfoPb,
	)
	resp, err := ShareInfoFromPb(&shareInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *sharesImpl) Delete(ctx context.Context, request DeleteShareRequest) error {
	requestPb, pbErr := DeleteShareRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *sharesImpl) Get(ctx context.Context, request GetShareRequest) (*ShareInfo, error) {
	requestPb, pbErr := GetShareRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var shareInfoPb sharingpb.ShareInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&shareInfoPb,
	)
	resp, err := ShareInfoFromPb(&shareInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets an array of data object shares from the metastore. The caller must be a
// metastore admin or the owner of the share. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *sharesImpl) ListAll(ctx context.Context, request ListSharesRequest) ([]ShareInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ShareInfo](ctx, iterator)
}

func (a *sharesImpl) internalList(ctx context.Context, request ListSharesRequest) (*ListSharesResponse, error) {
	requestPb, pbErr := ListSharesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSharesResponsePb sharingpb.ListSharesResponsePb
	path := "/api/2.1/unity-catalog/shares"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listSharesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListSharesResponseFromPb(&listSharesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *sharesImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetSharePermissionsResponse, error) {
	requestPb, pbErr := SharePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getSharePermissionsResponsePb sharingpb.GetSharePermissionsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getSharePermissionsResponsePb,
	)
	resp, err := GetSharePermissionsResponseFromPb(&getSharePermissionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *sharesImpl) Update(ctx context.Context, request UpdateShare) (*ShareInfo, error) {
	requestPb, pbErr := UpdateShareToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var shareInfoPb sharingpb.ShareInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&shareInfoPb,
	)
	resp, err := ShareInfoFromPb(&shareInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *sharesImpl) UpdatePermissions(ctx context.Context, request UpdateSharePermissions) (*UpdateSharePermissionsResponse, error) {
	requestPb, pbErr := UpdateSharePermissionsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateSharePermissionsResponsePb sharingpb.UpdateSharePermissionsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updateSharePermissionsResponsePb,
	)
	resp, err := UpdateSharePermissionsResponseFromPb(&updateSharePermissionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
