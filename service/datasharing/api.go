// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package datasharing

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type DatasharingService interface {
	CreateProvider(ctx context.Context, createProviderRequest CreateProviderRequest) (*CreateProviderResponse, error)

	CreateRecipient(ctx context.Context, createRecipientRequest CreateRecipientRequest) (*CreateRecipientResponse, error)

	CreateShare(ctx context.Context, createShareRequest CreateShareRequest) (*CreateShareResponse, error)

	DeleteProvider(ctx context.Context, deleteProviderRequest DeleteProviderRequest) error

	DeleteRecipient(ctx context.Context, deleteRecipientRequest DeleteRecipientRequest) error

	DeleteShare(ctx context.Context, deleteShareRequest DeleteShareRequest) error

	GetActivationUrlInfo(ctx context.Context, getActivationUrlInfoRequest GetActivationUrlInfoRequest) error

	GetProvider(ctx context.Context, getProviderRequest GetProviderRequest) (*GetProviderResponse, error)

	GetRecipient(ctx context.Context, getRecipientRequest GetRecipientRequest) (*GetRecipientResponse, error)

	GetRecipientSharePermissions(ctx context.Context, getRecipientSharePermissionsRequest GetRecipientSharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error)

	GetShare(ctx context.Context, getShareRequest GetShareRequest) (*GetShareResponse, error)
	// Permissions API for data sharing.
	GetSharePermissions(ctx context.Context, getSharePermissionsRequest GetSharePermissionsRequest) (*GetSharePermissionsResponse, error)

	ListProviderShares(ctx context.Context, listProviderSharesRequest ListProviderSharesRequest) (*ListProviderSharesResponse, error)

	ListProviders(ctx context.Context) (*ListProvidersResponse, error)

	ListRecipients(ctx context.Context) (*ListRecipientsResponse, error)

	ListShares(ctx context.Context) (*ListSharesResponse, error)
	// Rpc to retrieve access token with an activation token. This is an public
	// API without any authentication.
	RetrieveToken(ctx context.Context, retrieveTokenRequest RetrieveTokenRequest) (*RetrieveTokenResponse, error)

	RotateRecipientToken(ctx context.Context, rotateRecipientTokenRequest RotateRecipientTokenRequest) (*RotateRecipientTokenResponse, error)

	UpdateProvider(ctx context.Context, updateProviderRequest UpdateProviderRequest) error

	UpdateRecipient(ctx context.Context, updateRecipientRequest UpdateRecipientRequest) error

	UpdateShare(ctx context.Context, updateShareRequest UpdateShareRequest) error

	UpdateSharePermissions(ctx context.Context, updateSharePermissionsRequest UpdateSharePermissionsRequest) error
	ListProviderSharesByProviderNameArg(ctx context.Context, providerNameArg string) (*ListProviderSharesResponse, error)
	DeleteProviderByNameArg(ctx context.Context, nameArg string) error
	GetProviderByNameArg(ctx context.Context, nameArg string) (*GetProviderResponse, error)
	RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error)
	GetRecipientByName(ctx context.Context, name string) (*GetRecipientResponse, error)
	DeleteRecipientByName(ctx context.Context, name string) error
	GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error
	GetRecipientSharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error)
	GetShareByName(ctx context.Context, name string) (*GetShareResponse, error)
	DeleteShareByName(ctx context.Context, name string) error
	GetSharePermissionsByName(ctx context.Context, name string) (*GetSharePermissionsResponse, error)
}

func New(client *client.DatabricksClient) DatasharingService {
	return &DatasharingAPI{
		client: client,
	}
}

type DatasharingAPI struct {
	client *client.DatabricksClient
}

func (a *DatasharingAPI) CreateProvider(ctx context.Context, request CreateProviderRequest) (*CreateProviderResponse, error) {
	var createProviderResponse CreateProviderResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Post(ctx, path, request, &createProviderResponse)
	return &createProviderResponse, err
}

func (a *DatasharingAPI) CreateRecipient(ctx context.Context, request CreateRecipientRequest) (*CreateRecipientResponse, error) {
	var createRecipientResponse CreateRecipientResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Post(ctx, path, request, &createRecipientResponse)
	return &createRecipientResponse, err
}

func (a *DatasharingAPI) CreateShare(ctx context.Context, request CreateShareRequest) (*CreateShareResponse, error) {
	var createShareResponse CreateShareResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Post(ctx, path, request, &createShareResponse)
	return &createShareResponse, err
}

func (a *DatasharingAPI) DeleteProvider(ctx context.Context, request DeleteProviderRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.NameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *DatasharingAPI) DeleteRecipient(ctx context.Context, request DeleteRecipientRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *DatasharingAPI) DeleteShare(ctx context.Context, request DeleteShareRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *DatasharingAPI) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation_info/%v", request.ActivationUrl)
	err := a.client.Get(ctx, path, request, nil)
	return err
}

func (a *DatasharingAPI) GetProvider(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.NameArg)
	err := a.client.Get(ctx, path, request, &getProviderResponse)
	return &getProviderResponse, err
}

func (a *DatasharingAPI) GetRecipient(ctx context.Context, request GetRecipientRequest) (*GetRecipientResponse, error) {
	var getRecipientResponse GetRecipientResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getRecipientResponse)
	return &getRecipientResponse, err
}

func (a *DatasharingAPI) GetRecipientSharePermissions(ctx context.Context, request GetRecipientSharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	var getRecipientSharePermissionsResponse GetRecipientSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/share-permissions", request.Name)
	err := a.client.Get(ctx, path, request, &getRecipientSharePermissionsResponse)
	return &getRecipientSharePermissionsResponse, err
}

func (a *DatasharingAPI) GetShare(ctx context.Context, request GetShareRequest) (*GetShareResponse, error) {
	var getShareResponse GetShareResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getShareResponse)
	return &getShareResponse, err
}

// Permissions API for data sharing.
func (a *DatasharingAPI) GetSharePermissions(ctx context.Context, request GetSharePermissionsRequest) (*GetSharePermissionsResponse, error) {
	var getSharePermissionsResponse GetSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	err := a.client.Get(ctx, path, request, &getSharePermissionsResponse)
	return &getSharePermissionsResponse, err
}

func (a *DatasharingAPI) ListProviderShares(ctx context.Context, request ListProviderSharesRequest) (*ListProviderSharesResponse, error) {
	var listProviderSharesResponse ListProviderSharesResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v/shares", request.ProviderNameArg)
	err := a.client.Get(ctx, path, request, &listProviderSharesResponse)
	return &listProviderSharesResponse, err
}

func (a *DatasharingAPI) ListProviders(ctx context.Context) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Get(ctx, path, nil, &listProvidersResponse)
	return &listProvidersResponse, err
}

func (a *DatasharingAPI) ListRecipients(ctx context.Context) (*ListRecipientsResponse, error) {
	var listRecipientsResponse ListRecipientsResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Get(ctx, path, nil, &listRecipientsResponse)
	return &listRecipientsResponse, err
}

func (a *DatasharingAPI) ListShares(ctx context.Context) (*ListSharesResponse, error) {
	var listSharesResponse ListSharesResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Get(ctx, path, nil, &listSharesResponse)
	return &listSharesResponse, err
}

// Rpc to retrieve access token with an activation token. This is an public API
// without any authentication.
func (a *DatasharingAPI) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	var retrieveTokenResponse RetrieveTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation/%v", request.ActivationUrl)
	err := a.client.Get(ctx, path, request, &retrieveTokenResponse)
	return &retrieveTokenResponse, err
}

func (a *DatasharingAPI) RotateRecipientToken(ctx context.Context, request RotateRecipientTokenRequest) (*RotateRecipientTokenResponse, error) {
	var rotateRecipientTokenResponse RotateRecipientTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/rotate-token", request.Name)
	err := a.client.Post(ctx, path, request, &rotateRecipientTokenResponse)
	return &rotateRecipientTokenResponse, err
}

func (a *DatasharingAPI) UpdateProvider(ctx context.Context, request UpdateProviderRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.NameArg)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *DatasharingAPI) UpdateRecipient(ctx context.Context, request UpdateRecipientRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.NameArg)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *DatasharingAPI) UpdateShare(ctx context.Context, request UpdateShareRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *DatasharingAPI) UpdateSharePermissions(ctx context.Context, request UpdateSharePermissionsRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *DatasharingAPI) ListProviderSharesByProviderNameArg(ctx context.Context, providerNameArg string) (*ListProviderSharesResponse, error) {
	return a.ListProviderShares(ctx, ListProviderSharesRequest{
		ProviderNameArg: providerNameArg,
	})
}

func (a *DatasharingAPI) DeleteProviderByNameArg(ctx context.Context, nameArg string) error {
	return a.DeleteProvider(ctx, DeleteProviderRequest{
		NameArg: nameArg,
	})
}

func (a *DatasharingAPI) GetProviderByNameArg(ctx context.Context, nameArg string) (*GetProviderResponse, error) {
	return a.GetProvider(ctx, GetProviderRequest{
		NameArg: nameArg,
	})
}

func (a *DatasharingAPI) RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error) {
	return a.RetrieveToken(ctx, RetrieveTokenRequest{
		ActivationUrl: activationUrl,
	})
}

func (a *DatasharingAPI) GetRecipientByName(ctx context.Context, name string) (*GetRecipientResponse, error) {
	return a.GetRecipient(ctx, GetRecipientRequest{
		Name: name,
	})
}

func (a *DatasharingAPI) DeleteRecipientByName(ctx context.Context, name string) error {
	return a.DeleteRecipient(ctx, DeleteRecipientRequest{
		Name: name,
	})
}

func (a *DatasharingAPI) GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error {
	return a.GetActivationUrlInfo(ctx, GetActivationUrlInfoRequest{
		ActivationUrl: activationUrl,
	})
}

func (a *DatasharingAPI) GetRecipientSharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error) {
	return a.GetRecipientSharePermissions(ctx, GetRecipientSharePermissionsRequest{
		Name: name,
	})
}

func (a *DatasharingAPI) GetShareByName(ctx context.Context, name string) (*GetShareResponse, error) {
	return a.GetShare(ctx, GetShareRequest{
		Name: name,
	})
}

func (a *DatasharingAPI) DeleteShareByName(ctx context.Context, name string) error {
	return a.DeleteShare(ctx, DeleteShareRequest{
		Name: name,
	})
}

func (a *DatasharingAPI) GetSharePermissionsByName(ctx context.Context, name string) (*GetSharePermissionsResponse, error) {
	return a.GetSharePermissions(ctx, GetSharePermissionsRequest{
		Name: name,
	})
}
