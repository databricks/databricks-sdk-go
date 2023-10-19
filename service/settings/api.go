// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Ip Access Lists, Account Settings, Credentials Manager, Ip Access Lists, Settings, Token Management, Tokens, Workspace Conf, etc.
package settings

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewAccountIpAccessLists(client *client.DatabricksClient) *AccountIpAccessListsAPI {
	return &AccountIpAccessListsAPI{
		impl: &accountIpAccessListsImpl{
			client: client,
		},
	}
}

// The Accounts IP Access List API enables account admins to configure IP access
// lists for access to the account console.
//
// Account IP Access Lists affect web application access and REST API access to
// the account console and account APIs. If the feature is disabled for the
// account, all access is allowed for this account. There is support for allow
// lists (inclusion) and block lists (exclusion).
//
// When a connection is attempted: 1. **First, all block lists are checked.** If
// the connection IP address matches any block list, the connection is rejected.
// 2. **If the connection was not rejected by block lists**, the IP address is
// compared with the allow lists.
//
// If there is at least one allow list for the account, the connection is
// allowed only if the IP address matches an allow list. If there are no allow
// lists for the account, all IP addresses are allowed.
//
// For all allow lists and block lists combined, the account supports a maximum
// of 1000 IP/CIDR values, where one CIDR counts as a single value.
//
// After changes to the account-level IP access lists, it can take a few minutes
// for changes to take effect.
type AccountIpAccessListsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountIpAccessListsService)
	impl AccountIpAccessListsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountIpAccessListsAPI) WithImpl(impl AccountIpAccessListsService) *AccountIpAccessListsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountIpAccessLists API implementation
func (a *AccountIpAccessListsAPI) Impl() AccountIpAccessListsService {
	return a.impl
}

// Create access list.
//
// Creates an IP access list for the account.
//
// A list can be an allow list or a block list. See the top of this file for a
// description of how the server treats allow lists and block lists at runtime.
//
// When creating or updating an IP access list:
//
// * For all allow lists and block lists combined, the API supports a maximum of
// 1000 IP/CIDR values, where one CIDR counts as a single value. Attempts to
// exceed that number return error 400 with `error_code` value `QUOTA_EXCEEDED`.
// * If the new list would block the calling user's current IP, error 400 is
// returned with `error_code` value `INVALID_STATE`.
//
// It can take a few minutes for the changes to take effect.
func (a *AccountIpAccessListsAPI) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete access list.
//
// Deletes an IP access list, specified by its list ID.
func (a *AccountIpAccessListsAPI) Delete(ctx context.Context, request DeleteAccountIpAccessListRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete access list.
//
// Deletes an IP access list, specified by its list ID.
func (a *AccountIpAccessListsAPI) DeleteByIpAccessListId(ctx context.Context, ipAccessListId string) error {
	return a.impl.Delete(ctx, DeleteAccountIpAccessListRequest{
		IpAccessListId: ipAccessListId,
	})
}

// Get IP access list.
//
// Gets an IP access list, specified by its list ID.
func (a *AccountIpAccessListsAPI) Get(ctx context.Context, request GetAccountIpAccessListRequest) (*GetIpAccessListResponse, error) {
	return a.impl.Get(ctx, request)
}

// Get IP access list.
//
// Gets an IP access list, specified by its list ID.
func (a *AccountIpAccessListsAPI) GetByIpAccessListId(ctx context.Context, ipAccessListId string) (*GetIpAccessListResponse, error) {
	return a.impl.Get(ctx, GetAccountIpAccessListRequest{
		IpAccessListId: ipAccessListId,
	})
}

// Get access lists.
//
// Gets all IP access lists for the specified account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountIpAccessListsAPI) List(ctx context.Context) *listing.Iterator[struct{}, *GetIpAccessListsResponse, IpAccessListInfo] {

	getNextPage := func(ctx context.Context, req struct{}) (*GetIpAccessListsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx)
	}
	getItems := func(resp *GetIpAccessListsResponse) []IpAccessListInfo {
		return resp.IpAccessLists
	}

	return listing.NewIterator(
		struct{}{},
		getNextPage,
		getItems,
		nil)
}

func (a *AccountIpAccessListsAPI) ListAll(ctx context.Context) ([]IpAccessListInfo, error) {
	var results []IpAccessListInfo
	iter := a.List(ctx)
	var err error
	var next IpAccessListInfo
	for next, err = iter.Next(ctx); err != nil; next, err = iter.Next(ctx) {

		results = append(results, next)
	}
	if err != listing.ErrNoMoreItems {
		return nil, err
	}
	return results, nil
}

// IpAccessListInfoLabelToListIdMap calls [AccountIpAccessListsAPI.ListAll] and creates a map of results with [IpAccessListInfo].Label as key and [IpAccessListInfo].ListId as value.
//
// Returns an error if there's more than one [IpAccessListInfo] with the same .Label.
//
// Note: All [IpAccessListInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountIpAccessListsAPI) IpAccessListInfoLabelToListIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Label
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Label: %s", key)
		}
		mapping[key] = v.ListId
	}
	return mapping, nil
}

// GetByLabel calls [AccountIpAccessListsAPI.IpAccessListInfoLabelToListIdMap] and returns a single [IpAccessListInfo].
//
// Returns an error if there's more than one [IpAccessListInfo] with the same .Label.
//
// Note: All [IpAccessListInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountIpAccessListsAPI) GetByLabel(ctx context.Context, name string) (*IpAccessListInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]IpAccessListInfo{}
	for _, v := range result {
		key := v.Label
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("IpAccessListInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of IpAccessListInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Replace access list.
//
// Replaces an IP access list, specified by its ID.
//
// A list can include allow lists and block lists. See the top of this file for
// a description of how the server treats allow lists and block lists at run
// time. When replacing an IP access list: * For all allow lists and block lists
// combined, the API supports a maximum of 1000 IP/CIDR values, where one CIDR
// counts as a single value. Attempts to exceed that number return error 400
// with `error_code` value `QUOTA_EXCEEDED`. * If the resulting list would block
// the calling user's current IP, error 400 is returned with `error_code` value
// `INVALID_STATE`. It can take a few minutes for the changes to take effect.
func (a *AccountIpAccessListsAPI) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	return a.impl.Replace(ctx, request)
}

// Update access list.
//
// Updates an existing IP access list, specified by its ID.
//
// A list can include allow lists and block lists. See the top of this file for
// a description of how the server treats allow lists and block lists at run
// time.
//
// When updating an IP access list:
//
// * For all allow lists and block lists combined, the API supports a maximum of
// 1000 IP/CIDR values, where one CIDR counts as a single value. Attempts to
// exceed that number return error 400 with `error_code` value `QUOTA_EXCEEDED`.
// * If the updated list would block the calling user's current IP, error 400 is
// returned with `error_code` value `INVALID_STATE`.
//
// It can take a few minutes for the changes to take effect.
func (a *AccountIpAccessListsAPI) Update(ctx context.Context, request UpdateIpAccessList) error {
	return a.impl.Update(ctx, request)
}

func NewAccountSettings(client *client.DatabricksClient) *AccountSettingsAPI {
	return &AccountSettingsAPI{
		impl: &accountSettingsImpl{
			client: client,
		},
	}
}

// The Personal Compute enablement setting lets you control which users can use
// the Personal Compute default policy to create compute resources. By default
// all users in all workspaces have access (ON), but you can change the setting
// to instead let individual workspaces configure access control (DELEGATE).
//
// There is only one instance of this setting per account. Since this setting
// has a default value, this setting is present on all accounts even though it's
// never set on a given account. Deletion reverts the value of the setting back
// to the default value.
type AccountSettingsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountSettingsService)
	impl AccountSettingsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountSettingsAPI) WithImpl(impl AccountSettingsService) *AccountSettingsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountSettings API implementation
func (a *AccountSettingsAPI) Impl() AccountSettingsService {
	return a.impl
}

// Delete Personal Compute setting.
//
// Reverts back the Personal Compute setting value to default (ON)
func (a *AccountSettingsAPI) DeletePersonalComputeSetting(ctx context.Context, request DeletePersonalComputeSettingRequest) (*DeletePersonalComputeSettingResponse, error) {
	return a.impl.DeletePersonalComputeSetting(ctx, request)
}

// Get Personal Compute setting.
//
// Gets the value of the Personal Compute setting.
func (a *AccountSettingsAPI) ReadPersonalComputeSetting(ctx context.Context, request ReadPersonalComputeSettingRequest) (*PersonalComputeSetting, error) {
	return a.impl.ReadPersonalComputeSetting(ctx, request)
}

// Update Personal Compute setting.
//
// Updates the value of the Personal Compute setting.
func (a *AccountSettingsAPI) UpdatePersonalComputeSetting(ctx context.Context, request UpdatePersonalComputeSettingRequest) (*PersonalComputeSetting, error) {
	return a.impl.UpdatePersonalComputeSetting(ctx, request)
}

func NewCredentialsManager(client *client.DatabricksClient) *CredentialsManagerAPI {
	return &CredentialsManagerAPI{
		impl: &credentialsManagerImpl{
			client: client,
		},
	}
}

// Credentials manager interacts with with Identity Providers to to perform
// token exchanges using stored credentials and refresh tokens.
type CredentialsManagerAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CredentialsManagerService)
	impl CredentialsManagerService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *CredentialsManagerAPI) WithImpl(impl CredentialsManagerService) *CredentialsManagerAPI {
	a.impl = impl
	return a
}

// Impl returns low-level CredentialsManager API implementation
func (a *CredentialsManagerAPI) Impl() CredentialsManagerService {
	return a.impl
}

// Exchange token.
//
// Exchange tokens with an Identity Provider to get a new access token. It
// allowes specifying scopes to determine token permissions.
func (a *CredentialsManagerAPI) ExchangeToken(ctx context.Context, request ExchangeTokenRequest) (*ExchangeTokenResponse, error) {
	return a.impl.ExchangeToken(ctx, request)
}

func NewIpAccessLists(client *client.DatabricksClient) *IpAccessListsAPI {
	return &IpAccessListsAPI{
		impl: &ipAccessListsImpl{
			client: client,
		},
	}
}

// IP Access List enables admins to configure IP access lists.
//
// IP access lists affect web application access and REST API access to this
// workspace only. If the feature is disabled for a workspace, all access is
// allowed for this workspace. There is support for allow lists (inclusion) and
// block lists (exclusion).
//
// When a connection is attempted: 1. **First, all block lists are checked.** If
// the connection IP address matches any block list, the connection is rejected.
// 2. **If the connection was not rejected by block lists**, the IP address is
// compared with the allow lists.
//
// If there is at least one allow list for the workspace, the connection is
// allowed only if the IP address matches an allow list. If there are no allow
// lists for the workspace, all IP addresses are allowed.
//
// For all allow lists and block lists combined, the workspace supports a
// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
//
// After changes to the IP access list feature, it can take a few minutes for
// changes to take effect.
type IpAccessListsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(IpAccessListsService)
	impl IpAccessListsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *IpAccessListsAPI) WithImpl(impl IpAccessListsService) *IpAccessListsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level IpAccessLists API implementation
func (a *IpAccessListsAPI) Impl() IpAccessListsService {
	return a.impl
}

// Create access list.
//
// Creates an IP access list for this workspace.
//
// A list can be an allow list or a block list. See the top of this file for a
// description of how the server treats allow lists and block lists at runtime.
//
// When creating or updating an IP access list:
//
// * For all allow lists and block lists combined, the API supports a maximum of
// 1000 IP/CIDR values, where one CIDR counts as a single value. Attempts to
// exceed that number return error 400 with `error_code` value `QUOTA_EXCEEDED`.
// * If the new list would block the calling user's current IP, error 400 is
// returned with `error_code` value `INVALID_STATE`.
//
// It can take a few minutes for the changes to take effect. **Note**: Your new
// IP access list has no effect until you enable the feature. See
// :method:workspaceconf/setStatus
func (a *IpAccessListsAPI) Create(ctx context.Context, request CreateIpAccessList) (*CreateIpAccessListResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete access list.
//
// Deletes an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) Delete(ctx context.Context, request DeleteIpAccessListRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete access list.
//
// Deletes an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) DeleteByIpAccessListId(ctx context.Context, ipAccessListId string) error {
	return a.impl.Delete(ctx, DeleteIpAccessListRequest{
		IpAccessListId: ipAccessListId,
	})
}

// Get access list.
//
// Gets an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) Get(ctx context.Context, request GetIpAccessListRequest) (*FetchIpAccessListResponse, error) {
	return a.impl.Get(ctx, request)
}

// Get access list.
//
// Gets an IP access list, specified by its list ID.
func (a *IpAccessListsAPI) GetByIpAccessListId(ctx context.Context, ipAccessListId string) (*FetchIpAccessListResponse, error) {
	return a.impl.Get(ctx, GetIpAccessListRequest{
		IpAccessListId: ipAccessListId,
	})
}

// Get access lists.
//
// Gets all IP access lists for the specified workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *IpAccessListsAPI) List(ctx context.Context) *listing.Iterator[struct{}, *ListIpAccessListResponse, IpAccessListInfo] {

	getNextPage := func(ctx context.Context, req struct{}) (*ListIpAccessListResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx)
	}
	getItems := func(resp *ListIpAccessListResponse) []IpAccessListInfo {
		return resp.IpAccessLists
	}

	return listing.NewIterator(
		struct{}{},
		getNextPage,
		getItems,
		nil)
}

func (a *IpAccessListsAPI) ListAll(ctx context.Context) ([]IpAccessListInfo, error) {
	var results []IpAccessListInfo
	iter := a.List(ctx)
	var err error
	var next IpAccessListInfo
	for next, err = iter.Next(ctx); err != nil; next, err = iter.Next(ctx) {

		results = append(results, next)
	}
	if err != listing.ErrNoMoreItems {
		return nil, err
	}
	return results, nil
}

// IpAccessListInfoLabelToListIdMap calls [IpAccessListsAPI.ListAll] and creates a map of results with [IpAccessListInfo].Label as key and [IpAccessListInfo].ListId as value.
//
// Returns an error if there's more than one [IpAccessListInfo] with the same .Label.
//
// Note: All [IpAccessListInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *IpAccessListsAPI) IpAccessListInfoLabelToListIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Label
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Label: %s", key)
		}
		mapping[key] = v.ListId
	}
	return mapping, nil
}

// GetByLabel calls [IpAccessListsAPI.IpAccessListInfoLabelToListIdMap] and returns a single [IpAccessListInfo].
//
// Returns an error if there's more than one [IpAccessListInfo] with the same .Label.
//
// Note: All [IpAccessListInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *IpAccessListsAPI) GetByLabel(ctx context.Context, name string) (*IpAccessListInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]IpAccessListInfo{}
	for _, v := range result {
		key := v.Label
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("IpAccessListInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of IpAccessListInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Replace access list.
//
// Replaces an IP access list, specified by its ID.
//
// A list can include allow lists and block lists. See the top of this file for
// a description of how the server treats allow lists and block lists at run
// time. When replacing an IP access list: * For all allow lists and block lists
// combined, the API supports a maximum of 1000 IP/CIDR values, where one CIDR
// counts as a single value. Attempts to exceed that number return error 400
// with `error_code` value `QUOTA_EXCEEDED`. * If the resulting list would block
// the calling user's current IP, error 400 is returned with `error_code` value
// `INVALID_STATE`. It can take a few minutes for the changes to take effect.
// Note that your resulting IP access list has no effect until you enable the
// feature. See :method:workspaceconf/setStatus.
func (a *IpAccessListsAPI) Replace(ctx context.Context, request ReplaceIpAccessList) error {
	return a.impl.Replace(ctx, request)
}

// Update access list.
//
// Updates an existing IP access list, specified by its ID.
//
// A list can include allow lists and block lists. See the top of this file for
// a description of how the server treats allow lists and block lists at run
// time.
//
// When updating an IP access list:
//
// * For all allow lists and block lists combined, the API supports a maximum of
// 1000 IP/CIDR values, where one CIDR counts as a single value. Attempts to
// exceed that number return error 400 with `error_code` value `QUOTA_EXCEEDED`.
// * If the updated list would block the calling user's current IP, error 400 is
// returned with `error_code` value `INVALID_STATE`.
//
// It can take a few minutes for the changes to take effect. Note that your
// resulting IP access list has no effect until you enable the feature. See
// :method:workspaceconf/setStatus.
func (a *IpAccessListsAPI) Update(ctx context.Context, request UpdateIpAccessList) error {
	return a.impl.Update(ctx, request)
}

func NewSettings(client *client.DatabricksClient) *SettingsAPI {
	return &SettingsAPI{
		impl: &settingsImpl{
			client: client,
		},
	}
}

// // TODO(yuyuan.tang) to add the description for the setting
type SettingsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(SettingsService)
	impl SettingsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *SettingsAPI) WithImpl(impl SettingsService) *SettingsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Settings API implementation
func (a *SettingsAPI) Impl() SettingsService {
	return a.impl
}

// Delete the default namespace.
//
// Deletes the default namespace.
func (a *SettingsAPI) DeleteDefaultWorkspaceNamespace(ctx context.Context, request DeleteDefaultWorkspaceNamespaceRequest) (*DeleteDefaultWorkspaceNamespaceResponse, error) {
	return a.impl.DeleteDefaultWorkspaceNamespace(ctx, request)
}

// Get the default namespace.
//
// Gets the default namespace.
func (a *SettingsAPI) ReadDefaultWorkspaceNamespace(ctx context.Context, request ReadDefaultWorkspaceNamespaceRequest) (*DefaultNamespaceSetting, error) {
	return a.impl.ReadDefaultWorkspaceNamespace(ctx, request)
}

// Updates the default namespace setting.
//
// Updates the default namespace setting for the workspace. A fresh etag needs
// to be provided in PATCH requests (as part the setting field). The etag can be
// retrieved by making a GET request before the PATCH request. Note that if the
// setting does not exist, GET will return a NOT_FOUND error and the etag will
// be present in the error response, which should be set in the PATCH request.
func (a *SettingsAPI) UpdateDefaultWorkspaceNamespace(ctx context.Context, request UpdateDefaultWorkspaceNamespaceRequest) (*DefaultNamespaceSetting, error) {
	return a.impl.UpdateDefaultWorkspaceNamespace(ctx, request)
}

func NewTokenManagement(client *client.DatabricksClient) *TokenManagementAPI {
	return &TokenManagementAPI{
		impl: &tokenManagementImpl{
			client: client,
		},
	}
}

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
type TokenManagementAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(TokenManagementService)
	impl TokenManagementService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *TokenManagementAPI) WithImpl(impl TokenManagementService) *TokenManagementAPI {
	a.impl = impl
	return a
}

// Impl returns low-level TokenManagement API implementation
func (a *TokenManagementAPI) Impl() TokenManagementService {
	return a.impl
}

// Create on-behalf token.
//
// Creates a token on behalf of a service principal.
func (a *TokenManagementAPI) CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error) {
	return a.impl.CreateOboToken(ctx, request)
}

// Delete a token.
//
// Deletes a token, specified by its ID.
func (a *TokenManagementAPI) Delete(ctx context.Context, request DeleteTokenManagementRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a token.
//
// Deletes a token, specified by its ID.
func (a *TokenManagementAPI) DeleteByTokenId(ctx context.Context, tokenId string) error {
	return a.impl.Delete(ctx, DeleteTokenManagementRequest{
		TokenId: tokenId,
	})
}

// Get token info.
//
// Gets information about a token, specified by its ID.
func (a *TokenManagementAPI) Get(ctx context.Context, request GetTokenManagementRequest) (*TokenInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get token info.
//
// Gets information about a token, specified by its ID.
func (a *TokenManagementAPI) GetByTokenId(ctx context.Context, tokenId string) (*TokenInfo, error) {
	return a.impl.Get(ctx, GetTokenManagementRequest{
		TokenId: tokenId,
	})
}

// Get token permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *TokenManagementAPI) GetPermissionLevels(ctx context.Context) (*GetTokenPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx)
}

// Get token permissions.
//
// Gets the permissions of all tokens. Tokens can inherit permissions from their
// root object.
func (a *TokenManagementAPI) GetPermissions(ctx context.Context) (*TokenPermissions, error) {
	return a.impl.GetPermissions(ctx)
}

// List all tokens.
//
// Lists all tokens associated with the specified workspace or user.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TokenManagementAPI) List(ctx context.Context, request ListTokenManagementRequest) *listing.Iterator[ListTokenManagementRequest, *ListTokensResponse, TokenInfo] {

	getNextPage := func(ctx context.Context, req ListTokenManagementRequest) (*ListTokensResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx, req)
	}
	getItems := func(resp *ListTokensResponse) []TokenInfo {
		return resp.TokenInfos
	}

	return listing.NewIterator(
		request,
		getNextPage,
		getItems,
		nil)
}

func (a *TokenManagementAPI) ListAll(ctx context.Context, request ListTokenManagementRequest) ([]TokenInfo, error) {
	var results []TokenInfo
	iter := a.List(ctx, request)
	var err error
	var next TokenInfo
	for next, err = iter.Next(ctx); err != nil; next, err = iter.Next(ctx) {

		results = append(results, next)
	}
	if err != listing.ErrNoMoreItems {
		return nil, err
	}
	return results, nil
}

// TokenInfoCommentToTokenIdMap calls [TokenManagementAPI.ListAll] and creates a map of results with [TokenInfo].Comment as key and [TokenInfo].TokenId as value.
//
// Returns an error if there's more than one [TokenInfo] with the same .Comment.
//
// Note: All [TokenInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TokenManagementAPI) TokenInfoCommentToTokenIdMap(ctx context.Context, request ListTokenManagementRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Comment
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Comment: %s", key)
		}
		mapping[key] = v.TokenId
	}
	return mapping, nil
}

// GetByComment calls [TokenManagementAPI.TokenInfoCommentToTokenIdMap] and returns a single [TokenInfo].
//
// Returns an error if there's more than one [TokenInfo] with the same .Comment.
//
// Note: All [TokenInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TokenManagementAPI) GetByComment(ctx context.Context, name string) (*TokenInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListTokenManagementRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]TokenInfo{}
	for _, v := range result {
		key := v.Comment
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("TokenInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of TokenInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Set token permissions.
//
// Sets permissions on all tokens. Tokens can inherit permissions from their
// root object.
func (a *TokenManagementAPI) SetPermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error) {
	return a.impl.SetPermissions(ctx, request)
}

// Update token permissions.
//
// Updates the permissions on all tokens. Tokens can inherit permissions from
// their root object.
func (a *TokenManagementAPI) UpdatePermissions(ctx context.Context, request TokenPermissionsRequest) (*TokenPermissions, error) {
	return a.impl.UpdatePermissions(ctx, request)
}

func NewTokens(client *client.DatabricksClient) *TokensAPI {
	return &TokensAPI{
		impl: &tokensImpl{
			client: client,
		},
	}
}

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
type TokensAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(TokensService)
	impl TokensService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *TokensAPI) WithImpl(impl TokensService) *TokensAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Tokens API implementation
func (a *TokensAPI) Impl() TokensService {
	return a.impl
}

// Create a user token.
//
// Creates and returns a token for a user. If this call is made through token
// authentication, it creates a token with the same client ID as the
// authenticated token. If the user's token quota is exceeded, this call returns
// an error **QUOTA_EXCEEDED**.
func (a *TokensAPI) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {
	return a.impl.Create(ctx, request)
}

// Revoke token.
//
// Revokes an access token.
//
// If a token with the specified ID is not valid, this call returns an error
// **RESOURCE_DOES_NOT_EXIST**.
func (a *TokensAPI) Delete(ctx context.Context, request RevokeTokenRequest) error {
	return a.impl.Delete(ctx, request)
}

// Revoke token.
//
// Revokes an access token.
//
// If a token with the specified ID is not valid, this call returns an error
// **RESOURCE_DOES_NOT_EXIST**.
func (a *TokensAPI) DeleteByTokenId(ctx context.Context, tokenId string) error {
	return a.impl.Delete(ctx, RevokeTokenRequest{
		TokenId: tokenId,
	})
}

// List tokens.
//
// Lists all the valid tokens for a user-workspace pair.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TokensAPI) List(ctx context.Context) *listing.Iterator[struct{}, *ListTokensResponse, TokenInfo] {

	getNextPage := func(ctx context.Context, req struct{}) (*ListTokensResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx)
	}
	getItems := func(resp *ListTokensResponse) []TokenInfo {
		return resp.TokenInfos
	}

	return listing.NewIterator(
		struct{}{},
		getNextPage,
		getItems,
		nil)
}

func (a *TokensAPI) ListAll(ctx context.Context) ([]TokenInfo, error) {
	var results []TokenInfo
	iter := a.List(ctx)
	var err error
	var next TokenInfo
	for next, err = iter.Next(ctx); err != nil; next, err = iter.Next(ctx) {

		results = append(results, next)
	}
	if err != listing.ErrNoMoreItems {
		return nil, err
	}
	return results, nil
}

// TokenInfoCommentToTokenIdMap calls [TokensAPI.ListAll] and creates a map of results with [TokenInfo].Comment as key and [TokenInfo].TokenId as value.
//
// Returns an error if there's more than one [TokenInfo] with the same .Comment.
//
// Note: All [TokenInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TokensAPI) TokenInfoCommentToTokenIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Comment
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Comment: %s", key)
		}
		mapping[key] = v.TokenId
	}
	return mapping, nil
}

// GetByComment calls [TokensAPI.TokenInfoCommentToTokenIdMap] and returns a single [TokenInfo].
//
// Returns an error if there's more than one [TokenInfo] with the same .Comment.
//
// Note: All [TokenInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TokensAPI) GetByComment(ctx context.Context, name string) (*TokenInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]TokenInfo{}
	for _, v := range result {
		key := v.Comment
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("TokenInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of TokenInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

func NewWorkspaceConf(client *client.DatabricksClient) *WorkspaceConfAPI {
	return &WorkspaceConfAPI{
		impl: &workspaceConfImpl{
			client: client,
		},
	}
}

// This API allows updating known workspace settings for advanced users.
type WorkspaceConfAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WorkspaceConfService)
	impl WorkspaceConfService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *WorkspaceConfAPI) WithImpl(impl WorkspaceConfService) *WorkspaceConfAPI {
	a.impl = impl
	return a
}

// Impl returns low-level WorkspaceConf API implementation
func (a *WorkspaceConfAPI) Impl() WorkspaceConfService {
	return a.impl
}

// Check configuration status.
//
// Gets the configuration status for a workspace.
func (a *WorkspaceConfAPI) GetStatus(ctx context.Context, request GetStatusRequest) (*map[string]string, error) {
	return a.impl.GetStatus(ctx, request)
}

// Enable/disable features.
//
// Sets the configuration status for a workspace, including enabling or
// disabling it.
func (a *WorkspaceConfAPI) SetStatus(ctx context.Context, request WorkspaceConf) error {
	return a.impl.SetStatus(ctx, request)
}
