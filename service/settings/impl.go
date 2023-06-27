// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AccountIpAccessLists API methods
type accountIpAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *accountIpAccessListsImpl) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {
	var createIpAccessListResponse CreateIpAccessListResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/ip-access-lists", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &createIpAccessListResponse)
	return &createIpAccessListResponse, err
}

func (a *accountIpAccessListsImpl) Delete(ctx context.Context, request DeleteAccountIpAccessListRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountIpAccessListsImpl) Get(ctx context.Context, request GetAccountIpAccessListRequest) (*GetIpAccessListResponse, error) {
	var getIpAccessListResponse GetIpAccessListResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getIpAccessListResponse)
	return &getIpAccessListResponse, err
}

func (a *accountIpAccessListsImpl) List(ctx context.Context) (*GetIpAccessListsResponse, error) {
	var getIpAccessListsResponse GetIpAccessListsResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/ip-access-lists", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getIpAccessListsResponse)
	return &getIpAccessListsResponse, err
}

func (a *accountIpAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *accountIpAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/ip-access-lists/%v", a.client.ConfiguredAccountID(), request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just AccountSettings API methods
type accountSettingsImpl struct {
	client *client.DatabricksClient
}

func (a *accountSettingsImpl) DeletePersonalComputeSetting(ctx context.Context, request DeletePersonalComputeSettingRequest) (*DeletePersonalComputeSettingResponse, error) {
	var deletePersonalComputeSettingResponse DeletePersonalComputeSettingResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodDelete, path, request, &deletePersonalComputeSettingResponse)
	return &deletePersonalComputeSettingResponse, err
}

func (a *accountSettingsImpl) ReadPersonalComputeSetting(ctx context.Context, request ReadPersonalComputeSettingRequest) (*PersonalComputeSetting, error) {
	var personalComputeSetting PersonalComputeSetting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &personalComputeSetting)
	return &personalComputeSetting, err
}

func (a *accountSettingsImpl) UpdatePersonalComputeSetting(ctx context.Context, request UpdatePersonalComputeSettingRequest) (*PersonalComputeSetting, error) {
	var personalComputeSetting PersonalComputeSetting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/types/dcp_acct_enable/names/default", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPatch, path, request, &personalComputeSetting)
	return &personalComputeSetting, err
}

// unexported type that holds implementations of just IpAccessLists API methods
type ipAccessListsImpl struct {
	client *client.DatabricksClient
}

func (a *ipAccessListsImpl) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {
	var createIpAccessListResponse CreateIpAccessListResponse
	path := "/api/2.0/ip-access-lists"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createIpAccessListResponse)
	return &createIpAccessListResponse, err
}

func (a *ipAccessListsImpl) Delete(ctx context.Context, request DeleteIpAccessListRequest) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *ipAccessListsImpl) Get(ctx context.Context, request GetIpAccessListRequest) (*FetchIpAccessListResponse, error) {
	var fetchIpAccessListResponse FetchIpAccessListResponse
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &fetchIpAccessListResponse)
	return &fetchIpAccessListResponse, err
}

func (a *ipAccessListsImpl) List(ctx context.Context) (*GetIpAccessListResponse, error) {
	var getIpAccessListResponse GetIpAccessListResponse
	path := "/api/2.0/ip-access-lists"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getIpAccessListResponse)
	return &getIpAccessListResponse, err
}

func (a *ipAccessListsImpl) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *ipAccessListsImpl) Update(ctx context.Context, request UpdateIpAccessList) error {
	path := fmt.Sprintf("/api/2.0/ip-access-lists/%v", request.IpAccessListId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just TokenManagement API methods
type tokenManagementImpl struct {
	client *client.DatabricksClient
}

func (a *tokenManagementImpl) CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error) {
	var createOboTokenResponse CreateOboTokenResponse
	path := "/api/2.0/token-management/on-behalf-of/tokens"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createOboTokenResponse)
	return &createOboTokenResponse, err
}

func (a *tokenManagementImpl) Delete(ctx context.Context, request DeleteTokenManagementRequest) error {
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *tokenManagementImpl) Get(ctx context.Context, request GetTokenManagementRequest) (*TokenInfo, error) {
	var tokenInfo TokenInfo
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &tokenInfo)
	return &tokenInfo, err
}

func (a *tokenManagementImpl) List(ctx context.Context, request ListTokenManagementRequest) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token-management/tokens"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listTokensResponse)
	return &listTokensResponse, err
}

// unexported type that holds implementations of just Tokens API methods
type tokensImpl struct {
	client *client.DatabricksClient
}

func (a *tokensImpl) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {
	var createTokenResponse CreateTokenResponse
	path := "/api/2.0/token/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createTokenResponse)
	return &createTokenResponse, err
}

func (a *tokensImpl) Delete(ctx context.Context, request RevokeTokenRequest) error {
	path := "/api/2.0/token/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *tokensImpl) List(ctx context.Context) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token/list"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listTokensResponse)
	return &listTokensResponse, err
}

// unexported type that holds implementations of just WorkspaceConf API methods
type workspaceConfImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceConfImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*map[string]string, error) {
	var workspaceConf map[string]string
	path := "/api/2.0/workspace-conf"
	err := a.client.Do(ctx, http.MethodGet, path, request, &workspaceConf)
	return &workspaceConf, err
}

func (a *workspaceConfImpl) SetStatus(ctx context.Context, request WorkspaceConf) error {
	path := "/api/2.0/workspace-conf"
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}
