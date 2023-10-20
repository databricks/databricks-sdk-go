// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Clean Rooms, Providers, Recipient Activation, Recipients, Shares, etc.
package sharing

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewCleanRooms(client *client.DatabricksClient) *CleanRoomsAPI {
	return &CleanRoomsAPI{
		impl: &cleanRoomsImpl{
			client: client,
		},
	}
}

// A clean room is a secure, privacy-protecting environment where two or more
// parties can share sensitive enterprise data, including customer data, for
// measurements, insights, activation and other use cases.
//
// To create clean rooms, you must be a metastore admin or a user with the
// **CREATE_CLEAN_ROOM** privilege.
type CleanRoomsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CleanRoomsService)
	impl CleanRoomsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *CleanRoomsAPI) WithImpl(impl CleanRoomsService) *CleanRoomsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level CleanRooms API implementation
func (a *CleanRoomsAPI) Impl() CleanRoomsService {
	return a.impl
}

// Create a clean room.
//
// Creates a new clean room with specified colaborators. The caller must be a
// metastore admin or have the **CREATE_CLEAN_ROOM** privilege on the metastore.
func (a *CleanRoomsAPI) Create(ctx context.Context, request CreateCleanRoom) (*CleanRoomInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a clean room.
//
// Deletes a data object clean room from the metastore. The caller must be an
// owner of the clean room.
func (a *CleanRoomsAPI) Delete(ctx context.Context, request DeleteCleanRoomRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a clean room.
//
// Deletes a data object clean room from the metastore. The caller must be an
// owner of the clean room.
func (a *CleanRoomsAPI) DeleteByNameArg(ctx context.Context, nameArg string) error {
	return a.impl.Delete(ctx, DeleteCleanRoomRequest{
		NameArg: nameArg,
	})
}

// Get a clean room.
//
// Gets a data object clean room from the metastore. The caller must be a
// metastore admin or the owner of the clean room.
func (a *CleanRoomsAPI) Get(ctx context.Context, request GetCleanRoomRequest) (*CleanRoomInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a clean room.
//
// Gets a data object clean room from the metastore. The caller must be a
// metastore admin or the owner of the clean room.
func (a *CleanRoomsAPI) GetByNameArg(ctx context.Context, nameArg string) (*CleanRoomInfo, error) {
	return a.impl.Get(ctx, GetCleanRoomRequest{
		NameArg: nameArg,
	})
}

// List clean rooms.
//
// Gets an array of data object clean rooms from the metastore. The caller must
// be a metastore admin or the owner of the clean room. There is no guarantee of
// a specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *CleanRoomsAPI) List(ctx context.Context, request ListCleanRoomsRequest) (it listing.Iterator[CleanRoomInfo]) {

	getNextPage := func(ctx context.Context, req ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx, req)
	}
	getItems := func(resp *ListCleanRoomsResponse) []CleanRoomInfo {
		return resp.CleanRooms
	}
	getNextReq := func(resp *ListCleanRoomsResponse) *ListCleanRoomsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	it = listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)

	return it
}

func (a *CleanRoomsAPI) ListAll(ctx context.Context, request ListCleanRoomsRequest) ([]CleanRoomInfo, error) {
	iter := a.List(ctx, request)
	var results []CleanRoomInfo
	var totalCount int = 0
	limit := request.MaxResults
	for iter.HasNext(ctx) {
		next, err := iter.Next(ctx)
		if err != nil {
			return nil, err
		}
		results = append(results, next)
		totalCount++
		if limit > 0 && totalCount >= limit {
			break
		}
	}
	return results, nil

}

// Update a clean room.
//
// Updates the clean room with the changes and data objects in the request. The
// caller must be the owner of the clean room or a metastore admin.
//
// When the caller is a metastore admin, only the __owner__ field can be
// updated.
//
// In the case that the clean room name is changed **updateCleanRoom** requires
// that the caller is both the clean room owner and a metastore admin.
//
// For each table that is added through this method, the clean room owner must
// also have **SELECT** privilege on the table. The privilege must be maintained
// indefinitely for recipients to be able to access the table. Typically, you
// should use a group as the clean room owner.
//
// Table removals through **update** do not require additional privileges.
func (a *CleanRoomsAPI) Update(ctx context.Context, request UpdateCleanRoom) (*CleanRoomInfo, error) {
	return a.impl.Update(ctx, request)
}

func NewProviders(client *client.DatabricksClient) *ProvidersAPI {
	return &ProvidersAPI{
		impl: &providersImpl{
			client: client,
		},
	}
}

// A data provider is an object representing the organization in the real world
// who shares the data. A provider contains shares which further contain the
// shared data.
type ProvidersAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ProvidersService)
	impl ProvidersService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ProvidersAPI) WithImpl(impl ProvidersService) *ProvidersAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Providers API implementation
func (a *ProvidersAPI) Impl() ProvidersService {
	return a.impl
}

// Create an auth provider.
//
// Creates a new authentication provider minimally based on a name and
// authentication type. The caller must be an admin on the metastore.
func (a *ProvidersAPI) Create(ctx context.Context, request CreateProvider) (*ProviderInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a provider.
//
// Deletes an authentication provider, if the caller is a metastore admin or is
// the owner of the provider.
func (a *ProvidersAPI) Delete(ctx context.Context, request DeleteProviderRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a provider.
//
// Deletes an authentication provider, if the caller is a metastore admin or is
// the owner of the provider.
func (a *ProvidersAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteProviderRequest{
		Name: name,
	})
}

// Get a provider.
//
// Gets a specific authentication provider. The caller must supply the name of
// the provider, and must either be a metastore admin or the owner of the
// provider.
func (a *ProvidersAPI) Get(ctx context.Context, request GetProviderRequest) (*ProviderInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a provider.
//
// Gets a specific authentication provider. The caller must supply the name of
// the provider, and must either be a metastore admin or the owner of the
// provider.
func (a *ProvidersAPI) GetByName(ctx context.Context, name string) (*ProviderInfo, error) {
	return a.impl.Get(ctx, GetProviderRequest{
		Name: name,
	})
}

// List providers.
//
// Gets an array of available authentication providers. The caller must either
// be a metastore admin or the owner of the providers. Providers not owned by
// the caller are not included in the response. There is no guarantee of a
// specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ProvidersAPI) List(ctx context.Context, request ListProvidersRequest) (it listing.Iterator[ProviderInfo]) {

	getNextPage := func(ctx context.Context, req ListProvidersRequest) (*ListProvidersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx, req)
	}
	getItems := func(resp *ListProvidersResponse) []ProviderInfo {
		return resp.Providers
	}

	it = listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)

	return it
}

func (a *ProvidersAPI) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iter := a.List(ctx, request)
	return listing.ToSlice(ctx, iter)
}

// ProviderInfoNameToMetastoreIdMap calls [ProvidersAPI.ListAll] and creates a map of results with [ProviderInfo].Name as key and [ProviderInfo].MetastoreId as value.
//
// Returns an error if there's more than one [ProviderInfo] with the same .Name.
//
// Note: All [ProviderInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ProvidersAPI) ProviderInfoNameToMetastoreIdMap(ctx context.Context, request ListProvidersRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.MetastoreId
	}
	return mapping, nil
}

// List shares by Provider.
//
// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ProvidersAPI) ListShares(ctx context.Context, request ListSharesRequest) (it listing.Iterator[ProviderShare]) {

	getNextPage := func(ctx context.Context, req ListSharesRequest) (*ListProviderSharesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.ListShares(ctx, req)
	}
	getItems := func(resp *ListProviderSharesResponse) []ProviderShare {
		return resp.Shares
	}

	it = listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)

	return it
}

func (a *ProvidersAPI) ListSharesAll(ctx context.Context, request ListSharesRequest) ([]ProviderShare, error) {
	iter := a.ListShares(ctx, request)
	return listing.ToSlice(ctx, iter)
}

// List shares by Provider.
//
// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *ProvidersAPI) ListSharesByName(ctx context.Context, name string) (*ListProviderSharesResponse, error) {
	return a.impl.ListShares(ctx, ListSharesRequest{
		Name: name,
	})
}

// Update a provider.
//
// Updates the information for an authentication provider, if the caller is a
// metastore admin or is the owner of the provider. If the update changes the
// provider name, the caller must be both a metastore admin and the owner of the
// provider.
func (a *ProvidersAPI) Update(ctx context.Context, request UpdateProvider) (*ProviderInfo, error) {
	return a.impl.Update(ctx, request)
}

func NewRecipientActivation(client *client.DatabricksClient) *RecipientActivationAPI {
	return &RecipientActivationAPI{
		impl: &recipientActivationImpl{
			client: client,
		},
	}
}

// The Recipient Activation API is only applicable in the open sharing model
// where the recipient object has the authentication type of `TOKEN`. The data
// recipient follows the activation link shared by the data provider to download
// the credential file that includes the access token. The recipient will then
// use the credential file to establish a secure connection with the provider to
// receive the shared data.
//
// Note that you can download the credential file only once. Recipients should
// treat the downloaded credential as a secret and must not share it outside of
// their organization.
type RecipientActivationAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(RecipientActivationService)
	impl RecipientActivationService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *RecipientActivationAPI) WithImpl(impl RecipientActivationService) *RecipientActivationAPI {
	a.impl = impl
	return a
}

// Impl returns low-level RecipientActivation API implementation
func (a *RecipientActivationAPI) Impl() RecipientActivationService {
	return a.impl
}

// Get a share activation URL.
//
// Gets an activation URL for a share.
func (a *RecipientActivationAPI) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	return a.impl.GetActivationUrlInfo(ctx, request)
}

// Get a share activation URL.
//
// Gets an activation URL for a share.
func (a *RecipientActivationAPI) GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error {
	return a.impl.GetActivationUrlInfo(ctx, GetActivationUrlInfoRequest{
		ActivationUrl: activationUrl,
	})
}

// Get an access token.
//
// Retrieve access token with an activation url. This is a public API without
// any authentication.
func (a *RecipientActivationAPI) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	return a.impl.RetrieveToken(ctx, request)
}

// Get an access token.
//
// Retrieve access token with an activation url. This is a public API without
// any authentication.
func (a *RecipientActivationAPI) RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error) {
	return a.impl.RetrieveToken(ctx, RetrieveTokenRequest{
		ActivationUrl: activationUrl,
	})
}

func NewRecipients(client *client.DatabricksClient) *RecipientsAPI {
	return &RecipientsAPI{
		impl: &recipientsImpl{
			client: client,
		},
	}
}

// A recipient is an object you create using :method:recipients/create to
// represent an organization which you want to allow access shares. The way how
// sharing works differs depending on whether or not your recipient has access
// to a Databricks workspace that is enabled for Unity Catalog:
//
// - For recipients with access to a Databricks workspace that is enabled for
// Unity Catalog, you can create a recipient object along with a unique sharing
// identifier you get from the recipient. The sharing identifier is the key
// identifier that enables the secure connection. This sharing mode is called
// **Databricks-to-Databricks sharing**.
//
// - For recipients without access to a Databricks workspace that is enabled for
// Unity Catalog, when you create a recipient object, Databricks generates an
// activation link you can send to the recipient. The recipient follows the
// activation link to download the credential file, and then uses the credential
// file to establish a secure connection to receive the shared data. This
// sharing mode is called **open sharing**.
type RecipientsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(RecipientsService)
	impl RecipientsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *RecipientsAPI) WithImpl(impl RecipientsService) *RecipientsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Recipients API implementation
func (a *RecipientsAPI) Impl() RecipientsService {
	return a.impl
}

// Create a share recipient.
//
// Creates a new recipient with the delta sharing authentication type in the
// metastore. The caller must be a metastore admin or has the
// **CREATE_RECIPIENT** privilege on the metastore.
func (a *RecipientsAPI) Create(ctx context.Context, request CreateRecipient) (*RecipientInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a share recipient.
//
// Deletes the specified recipient from the metastore. The caller must be the
// owner of the recipient.
func (a *RecipientsAPI) Delete(ctx context.Context, request DeleteRecipientRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a share recipient.
//
// Deletes the specified recipient from the metastore. The caller must be the
// owner of the recipient.
func (a *RecipientsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteRecipientRequest{
		Name: name,
	})
}

// Get a share recipient.
//
// Gets a share recipient from the metastore if:
//
// * the caller is the owner of the share recipient, or: * is a metastore admin
func (a *RecipientsAPI) Get(ctx context.Context, request GetRecipientRequest) (*RecipientInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a share recipient.
//
// Gets a share recipient from the metastore if:
//
// * the caller is the owner of the share recipient, or: * is a metastore admin
func (a *RecipientsAPI) GetByName(ctx context.Context, name string) (*RecipientInfo, error) {
	return a.impl.Get(ctx, GetRecipientRequest{
		Name: name,
	})
}

// List share recipients.
//
// Gets an array of all share recipients within the current metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner. There is no
// guarantee of a specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RecipientsAPI) List(ctx context.Context, request ListRecipientsRequest) (it listing.Iterator[RecipientInfo]) {

	getNextPage := func(ctx context.Context, req ListRecipientsRequest) (*ListRecipientsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx, req)
	}
	getItems := func(resp *ListRecipientsResponse) []RecipientInfo {
		return resp.Recipients
	}

	it = listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)

	return it
}

func (a *RecipientsAPI) ListAll(ctx context.Context, request ListRecipientsRequest) ([]RecipientInfo, error) {
	iter := a.List(ctx, request)
	return listing.ToSlice(ctx, iter)
}

// RecipientInfoNameToMetastoreIdMap calls [RecipientsAPI.ListAll] and creates a map of results with [RecipientInfo].Name as key and [RecipientInfo].MetastoreId as value.
//
// Returns an error if there's more than one [RecipientInfo] with the same .Name.
//
// Note: All [RecipientInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RecipientsAPI) RecipientInfoNameToMetastoreIdMap(ctx context.Context, request ListRecipientsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.MetastoreId
	}
	return mapping, nil
}

// Rotate a token.
//
// Refreshes the specified recipient's delta sharing authentication token with
// the provided token info. The caller must be the owner of the recipient.
func (a *RecipientsAPI) RotateToken(ctx context.Context, request RotateRecipientToken) (*RecipientInfo, error) {
	return a.impl.RotateToken(ctx, request)
}

// Get recipient share permissions.
//
// Gets the share permissions for the specified Recipient. The caller must be a
// metastore admin or the owner of the Recipient.
func (a *RecipientsAPI) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	return a.impl.SharePermissions(ctx, request)
}

// Get recipient share permissions.
//
// Gets the share permissions for the specified Recipient. The caller must be a
// metastore admin or the owner of the Recipient.
func (a *RecipientsAPI) SharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error) {
	return a.impl.SharePermissions(ctx, SharePermissionsRequest{
		Name: name,
	})
}

// Update a share recipient.
//
// Updates an existing recipient in the metastore. The caller must be a
// metastore admin or the owner of the recipient. If the recipient name will be
// updated, the user must be both a metastore admin and the owner of the
// recipient.
func (a *RecipientsAPI) Update(ctx context.Context, request UpdateRecipient) error {
	return a.impl.Update(ctx, request)
}

func NewShares(client *client.DatabricksClient) *SharesAPI {
	return &SharesAPI{
		impl: &sharesImpl{
			client: client,
		},
	}
}

// A share is a container instantiated with :method:shares/create. Once created
// you can iteratively register a collection of existing data assets defined
// within the metastore using :method:shares/update. You can register data
// assets under their original name, qualified by their original schema, or
// provide alternate exposed names.
type SharesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(SharesService)
	impl SharesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *SharesAPI) WithImpl(impl SharesService) *SharesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Shares API implementation
func (a *SharesAPI) Impl() SharesService {
	return a.impl
}

// Create a share.
//
// Creates a new share for data objects. Data objects can be added after
// creation with **update**. The caller must be a metastore admin or have the
// **CREATE_SHARE** privilege on the metastore.
func (a *SharesAPI) Create(ctx context.Context, request CreateShare) (*ShareInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a share.
//
// Deletes a data object share from the metastore. The caller must be an owner
// of the share.
func (a *SharesAPI) Delete(ctx context.Context, request DeleteShareRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a share.
//
// Deletes a data object share from the metastore. The caller must be an owner
// of the share.
func (a *SharesAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteShareRequest{
		Name: name,
	})
}

// Get a share.
//
// Gets a data object share from the metastore. The caller must be a metastore
// admin or the owner of the share.
func (a *SharesAPI) Get(ctx context.Context, request GetShareRequest) (*ShareInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a share.
//
// Gets a data object share from the metastore. The caller must be a metastore
// admin or the owner of the share.
func (a *SharesAPI) GetByName(ctx context.Context, name string) (*ShareInfo, error) {
	return a.impl.Get(ctx, GetShareRequest{
		Name: name,
	})
}

// List shares.
//
// Gets an array of data object shares from the metastore. The caller must be a
// metastore admin or the owner of the share. There is no guarantee of a
// specific ordering of the elements in the array.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SharesAPI) List(ctx context.Context) (it listing.Iterator[ShareInfo]) {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListSharesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx)
	}
	getItems := func(resp *ListSharesResponse) []ShareInfo {
		return resp.Shares
	}

	it = listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)

	return it
}

func (a *SharesAPI) ListAll(ctx context.Context) ([]ShareInfo, error) {
	iter := a.List(ctx)
	return listing.ToSlice(ctx, iter)
}

// Get permissions.
//
// Gets the permissions for a data share from the metastore. The caller must be
// a metastore admin or the owner of the share.
func (a *SharesAPI) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*catalog.PermissionsList, error) {
	return a.impl.SharePermissions(ctx, request)
}

// Get permissions.
//
// Gets the permissions for a data share from the metastore. The caller must be
// a metastore admin or the owner of the share.
func (a *SharesAPI) SharePermissionsByName(ctx context.Context, name string) (*catalog.PermissionsList, error) {
	return a.impl.SharePermissions(ctx, SharePermissionsRequest{
		Name: name,
	})
}

// Update a share.
//
// Updates the share with the changes and data objects in the request. The
// caller must be the owner of the share or a metastore admin.
//
// When the caller is a metastore admin, only the __owner__ field can be
// updated.
//
// In the case that the share name is changed, **updateShare** requires that the
// caller is both the share owner and a metastore admin.
//
// For each table that is added through this method, the share owner must also
// have **SELECT** privilege on the table. This privilege must be maintained
// indefinitely for recipients to be able to access the table. Typically, you
// should use a group as the share owner.
//
// Table removals through **update** do not require additional privileges.
func (a *SharesAPI) Update(ctx context.Context, request UpdateShare) (*ShareInfo, error) {
	return a.impl.Update(ctx, request)
}

// Update permissions.
//
// Updates the permissions for a data share in the metastore. The caller must be
// a metastore admin or an owner of the share.
//
// For new recipient grants, the user must also be the owner of the recipients.
// recipient revocations do not require additional privileges.
func (a *SharesAPI) UpdatePermissions(ctx context.Context, request UpdateSharePermissions) error {
	return a.impl.UpdatePermissions(ctx, request)
}
