// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Providers, Recipient Activation, Recipient Federation Policies, Recipients, Shares, etc.
package sharing

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type ProvidersInterface interface {

	// Creates a new authentication provider minimally based on a name and
	// authentication type. The caller must be an admin on the metastore.
	Create(ctx context.Context, request CreateProvider) (*ProviderInfo, error)

	// Deletes an authentication provider, if the caller is a metastore admin or is
	// the owner of the provider.
	Delete(ctx context.Context, request DeleteProviderRequest) error

	// Deletes an authentication provider, if the caller is a metastore admin or is
	// the owner of the provider.
	DeleteByName(ctx context.Context, name string) error

	// Gets a specific authentication provider. The caller must supply the name of
	// the provider, and must either be a metastore admin or the owner of the
	// provider.
	Get(ctx context.Context, request GetProviderRequest) (*ProviderInfo, error)

	// Gets a specific authentication provider. The caller must supply the name of
	// the provider, and must either be a metastore admin or the owner of the
	// provider.
	GetByName(ctx context.Context, name string) (*ProviderInfo, error)

	// Gets an array of available authentication providers. The caller must either
	// be a metastore admin or the owner of the providers. Providers not owned by
	// the caller are not included in the response. There is no guarantee of a
	// specific ordering of the elements in the array.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListProvidersRequest) listing.Iterator[ProviderInfo]

	// Gets an array of available authentication providers. The caller must either
	// be a metastore admin or the owner of the providers. Providers not owned by
	// the caller are not included in the response. There is no guarantee of a
	// specific ordering of the elements in the array.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error)

	// ProviderInfoNameToMetastoreIdMap calls [ProvidersAPI.ListAll] and creates a map of results with [ProviderInfo].Name as key and [ProviderInfo].MetastoreId as value.
	//
	// Returns an error if there's more than one [ProviderInfo] with the same .Name.
	//
	// Note: All [ProviderInfo] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	ProviderInfoNameToMetastoreIdMap(ctx context.Context, request ListProvidersRequest) (map[string]string, error)

	// Get arrays of assets associated with a specified provider's share. The caller
	// is the recipient of the share.
	ListProviderShareAssets(ctx context.Context, request ListProviderShareAssetsRequest) (*ListProviderShareAssetsResponse, error)

	// Get arrays of assets associated with a specified provider's share. The caller
	// is the recipient of the share.
	ListProviderShareAssetsByProviderNameAndShareName(ctx context.Context, providerName string, shareName string) (*ListProviderShareAssetsResponse, error)

	// Gets an array of a specified provider's shares within the metastore where:
	//
	// * the caller is a metastore admin, or * the caller is the owner.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListShares(ctx context.Context, request ListSharesRequest) listing.Iterator[ProviderShare]

	// Gets an array of a specified provider's shares within the metastore where:
	//
	// * the caller is a metastore admin, or * the caller is the owner.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSharesAll(ctx context.Context, request ListSharesRequest) ([]ProviderShare, error)

	// Gets an array of a specified provider's shares within the metastore where:
	//
	// * the caller is a metastore admin, or * the caller is the owner.
	ListSharesByName(ctx context.Context, name string) (*ListProviderSharesResponse, error)

	// Updates the information for an authentication provider, if the caller is a
	// metastore admin or is the owner of the provider. If the update changes the
	// provider name, the caller must be both a metastore admin and the owner of the
	// provider.
	Update(ctx context.Context, request UpdateProvider) (*ProviderInfo, error)
}

func NewProviders(client *client.DatabricksClient) *ProvidersAPI {
	return &ProvidersAPI{
		providersImpl: providersImpl{
			client: client,
		},
	}
}

// A data provider is an object representing the organization in the real world
// who shares the data. A provider contains shares which further contain the
// shared data.
type ProvidersAPI struct {
	providersImpl
}

// Deletes an authentication provider, if the caller is a metastore admin or is
// the owner of the provider.
func (a *ProvidersAPI) DeleteByName(ctx context.Context, name string) error {
	return a.providersImpl.Delete(ctx, DeleteProviderRequest{
		Name: name,
	})
}

// Gets a specific authentication provider. The caller must supply the name of
// the provider, and must either be a metastore admin or the owner of the
// provider.
func (a *ProvidersAPI) GetByName(ctx context.Context, name string) (*ProviderInfo, error) {
	return a.providersImpl.Get(ctx, GetProviderRequest{
		Name: name,
	})
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

// Get arrays of assets associated with a specified provider's share. The caller
// is the recipient of the share.
func (a *ProvidersAPI) ListProviderShareAssetsByProviderNameAndShareName(ctx context.Context, providerName string, shareName string) (*ListProviderShareAssetsResponse, error) {
	return a.providersImpl.ListProviderShareAssets(ctx, ListProviderShareAssetsRequest{
		ProviderName: providerName,
		ShareName:    shareName,
	})
}

// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *ProvidersAPI) ListSharesByName(ctx context.Context, name string) (*ListProviderSharesResponse, error) {
	return a.providersImpl.internalListShares(ctx, ListSharesRequest{
		Name: name,
	})
}

type RecipientActivationInterface interface {

	// Gets an activation URL for a share.
	GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error

	// Gets an activation URL for a share.
	GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error

	// Retrieve access token with an activation url. This is a public API without
	// any authentication.
	RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error)

	// Retrieve access token with an activation url. This is a public API without
	// any authentication.
	RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error)
}

func NewRecipientActivation(client *client.DatabricksClient) *RecipientActivationAPI {
	return &RecipientActivationAPI{
		recipientActivationImpl: recipientActivationImpl{
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
	recipientActivationImpl
}

// Gets an activation URL for a share.
func (a *RecipientActivationAPI) GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error {
	return a.recipientActivationImpl.GetActivationUrlInfo(ctx, GetActivationUrlInfoRequest{
		ActivationUrl: activationUrl,
	})
}

// Retrieve access token with an activation url. This is a public API without
// any authentication.
func (a *RecipientActivationAPI) RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error) {
	return a.recipientActivationImpl.RetrieveToken(ctx, RetrieveTokenRequest{
		ActivationUrl: activationUrl,
	})
}

type RecipientFederationPoliciesInterface interface {

	// Create a federation policy for an OIDC_FEDERATION recipient for sharing data
	// from Databricks to non-Databricks recipients. The caller must be the owner of
	// the recipient. When sharing data from Databricks to non-Databricks clients,
	// you can define a federation policy to authenticate non-Databricks recipients.
	// The federation policy validates OIDC claims in federated tokens and is
	// defined at the recipient level. This enables secretless sharing clients to
	// authenticate using OIDC tokens.
	//
	// Supported scenarios for federation policies: 1. **User-to-Machine (U2M)
	// flow** (e.g., PowerBI): A user accesses a resource using their own identity.
	// 2. **Machine-to-Machine (M2M) flow** (e.g., OAuth App): An OAuth App accesses
	// a resource using its own identity, typically for tasks like running nightly
	// jobs.
	//
	// For an overview, refer to: - Blog post: Overview of feature:
	// https://www.databricks.com/blog/announcing-oidc-token-federation-enhanced-delta-sharing-security
	//
	// For detailed configuration guides based on your use case: - Creating a
	// Federation Policy as a provider:
	// https://docs.databricks.com/en/delta-sharing/create-recipient-oidc-fed -
	// Configuration and usage for Machine-to-Machine (M2M) applications (e.g.,
	// Python Delta Sharing Client):
	// https://docs.databricks.com/aws/en/delta-sharing/sharing-over-oidc-m2m -
	// Configuration and usage for User-to-Machine (U2M) applications (e.g.,
	// PowerBI):
	// https://docs.databricks.com/aws/en/delta-sharing/sharing-over-oidc-u2m
	Create(ctx context.Context, request CreateFederationPolicyRequest) (*FederationPolicy, error)

	// Deletes an existing federation policy for an OIDC_FEDERATION recipient. The
	// caller must be the owner of the recipient.
	Delete(ctx context.Context, request DeleteFederationPolicyRequest) error

	// Deletes an existing federation policy for an OIDC_FEDERATION recipient. The
	// caller must be the owner of the recipient.
	DeleteByRecipientNameAndName(ctx context.Context, recipientName string, name string) error

	// Reads an existing federation policy for an OIDC_FEDERATION recipient for
	// sharing data from Databricks to non-Databricks recipients. The caller must
	// have read access to the recipient.
	GetFederationPolicy(ctx context.Context, request GetFederationPolicyRequest) (*FederationPolicy, error)

	// Reads an existing federation policy for an OIDC_FEDERATION recipient for
	// sharing data from Databricks to non-Databricks recipients. The caller must
	// have read access to the recipient.
	GetFederationPolicyByRecipientNameAndName(ctx context.Context, recipientName string, name string) (*FederationPolicy, error)

	// Lists federation policies for an OIDC_FEDERATION recipient for sharing data
	// from Databricks to non-Databricks recipients. The caller must have read
	// access to the recipient.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListFederationPoliciesRequest) listing.Iterator[FederationPolicy]

	// Lists federation policies for an OIDC_FEDERATION recipient for sharing data
	// from Databricks to non-Databricks recipients. The caller must have read
	// access to the recipient.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListFederationPoliciesRequest) ([]FederationPolicy, error)

	// Lists federation policies for an OIDC_FEDERATION recipient for sharing data
	// from Databricks to non-Databricks recipients. The caller must have read
	// access to the recipient.
	ListByRecipientName(ctx context.Context, recipientName string) (*ListFederationPoliciesResponse, error)

	// Updates an existing federation policy for an OIDC_RECIPIENT. The caller must
	// be the owner of the recipient.
	Update(ctx context.Context, request UpdateFederationPolicyRequest) (*FederationPolicy, error)
}

func NewRecipientFederationPolicies(client *client.DatabricksClient) *RecipientFederationPoliciesAPI {
	return &RecipientFederationPoliciesAPI{
		recipientFederationPoliciesImpl: recipientFederationPoliciesImpl{
			client: client,
		},
	}
}

// The Recipient Federation Policies APIs are only applicable in the open
// sharing model where the recipient object has the authentication type of
// `OIDC_RECIPIENT`, enabling data sharing from Databricks to non-Databricks
// recipients. OIDC Token Federation enables secure, secret-less authentication
// for accessing Delta Sharing servers. Users and applications authenticate
// using short-lived OIDC tokens issued by their own Identity Provider (IdP),
// such as Azure Entra ID or Okta, without the need for managing static
// credentials or client secrets. A federation policy defines how non-Databricks
// recipients authenticate using OIDC tokens. It validates the OIDC claims in
// federated tokens and is set at the recipient level. The caller must be the
// owner of the recipient to create or manage a federation policy. Federation
// policies support the following scenarios: - User-to-Machine (U2M) flow: A
// user accesses Delta Shares using their own identity, such as connecting
// through PowerBI Delta Sharing Client. - Machine-to-Machine (M2M) flow: An
// application accesses Delta Shares using its own identity, typically for
// automation tasks like nightly jobs through Python Delta Sharing Client. OIDC
// Token Federation enables fine-grained access control, supports Multi-Factor
// Authentication (MFA), and enhances security by minimizing the risk of
// credential leakage through the use of short-lived, expiring tokens. It is
// designed for strong identity governance, secure cross-platform data sharing,
// and reduced operational overhead for credential management.
//
// For more information, see
// https://www.databricks.com/blog/announcing-oidc-token-federation-enhanced-delta-sharing-security
// and https://docs.databricks.com/en/delta-sharing/create-recipient-oidc-fed
type RecipientFederationPoliciesAPI struct {
	recipientFederationPoliciesImpl
}

// Deletes an existing federation policy for an OIDC_FEDERATION recipient. The
// caller must be the owner of the recipient.
func (a *RecipientFederationPoliciesAPI) DeleteByRecipientNameAndName(ctx context.Context, recipientName string, name string) error {
	return a.recipientFederationPoliciesImpl.Delete(ctx, DeleteFederationPolicyRequest{
		RecipientName: recipientName,
		Name:          name,
	})
}

// Reads an existing federation policy for an OIDC_FEDERATION recipient for
// sharing data from Databricks to non-Databricks recipients. The caller must
// have read access to the recipient.
func (a *RecipientFederationPoliciesAPI) GetFederationPolicyByRecipientNameAndName(ctx context.Context, recipientName string, name string) (*FederationPolicy, error) {
	return a.recipientFederationPoliciesImpl.GetFederationPolicy(ctx, GetFederationPolicyRequest{
		RecipientName: recipientName,
		Name:          name,
	})
}

// Lists federation policies for an OIDC_FEDERATION recipient for sharing data
// from Databricks to non-Databricks recipients. The caller must have read
// access to the recipient.
func (a *RecipientFederationPoliciesAPI) ListByRecipientName(ctx context.Context, recipientName string) (*ListFederationPoliciesResponse, error) {
	return a.recipientFederationPoliciesImpl.internalList(ctx, ListFederationPoliciesRequest{
		RecipientName: recipientName,
	})
}

type RecipientsInterface interface {

	// Creates a new recipient with the delta sharing authentication type in the
	// metastore. The caller must be a metastore admin or have the
	// **CREATE_RECIPIENT** privilege on the metastore.
	Create(ctx context.Context, request CreateRecipient) (*RecipientInfo, error)

	// Deletes the specified recipient from the metastore. The caller must be the
	// owner of the recipient.
	Delete(ctx context.Context, request DeleteRecipientRequest) error

	// Deletes the specified recipient from the metastore. The caller must be the
	// owner of the recipient.
	DeleteByName(ctx context.Context, name string) error

	// Gets a share recipient from the metastore if:
	//
	// * the caller is the owner of the share recipient, or: * is a metastore admin
	Get(ctx context.Context, request GetRecipientRequest) (*RecipientInfo, error)

	// Gets a share recipient from the metastore if:
	//
	// * the caller is the owner of the share recipient, or: * is a metastore admin
	GetByName(ctx context.Context, name string) (*RecipientInfo, error)

	// Gets an array of all share recipients within the current metastore where:
	//
	// * the caller is a metastore admin, or * the caller is the owner. There is no
	// guarantee of a specific ordering of the elements in the array.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListRecipientsRequest) listing.Iterator[RecipientInfo]

	// Gets an array of all share recipients within the current metastore where:
	//
	// * the caller is a metastore admin, or * the caller is the owner. There is no
	// guarantee of a specific ordering of the elements in the array.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListRecipientsRequest) ([]RecipientInfo, error)

	// Refreshes the specified recipient's delta sharing authentication token with
	// the provided token info. The caller must be the owner of the recipient.
	RotateToken(ctx context.Context, request RotateRecipientToken) (*RecipientInfo, error)

	// Gets the share permissions for the specified Recipient. The caller must be a
	// metastore admin or the owner of the Recipient.
	SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error)

	// Gets the share permissions for the specified Recipient. The caller must be a
	// metastore admin or the owner of the Recipient.
	SharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error)

	// Updates an existing recipient in the metastore. The caller must be a
	// metastore admin or the owner of the recipient. If the recipient name will be
	// updated, the user must be both a metastore admin and the owner of the
	// recipient.
	Update(ctx context.Context, request UpdateRecipient) (*RecipientInfo, error)
}

func NewRecipients(client *client.DatabricksClient) *RecipientsAPI {
	return &RecipientsAPI{
		recipientsImpl: recipientsImpl{
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
	recipientsImpl
}

// Deletes the specified recipient from the metastore. The caller must be the
// owner of the recipient.
func (a *RecipientsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.recipientsImpl.Delete(ctx, DeleteRecipientRequest{
		Name: name,
	})
}

// Gets a share recipient from the metastore if:
//
// * the caller is the owner of the share recipient, or: * is a metastore admin
func (a *RecipientsAPI) GetByName(ctx context.Context, name string) (*RecipientInfo, error) {
	return a.recipientsImpl.Get(ctx, GetRecipientRequest{
		Name: name,
	})
}

// Gets the share permissions for the specified Recipient. The caller must be a
// metastore admin or the owner of the Recipient.
func (a *RecipientsAPI) SharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error) {
	return a.recipientsImpl.SharePermissions(ctx, SharePermissionsRequest{
		Name: name,
	})
}

type SharesInterface interface {

	// Creates a new share for data objects. Data objects can be added after
	// creation with **update**. The caller must be a metastore admin or have the
	// **CREATE_SHARE** privilege on the metastore.
	Create(ctx context.Context, request CreateShare) (*ShareInfo, error)

	// Deletes a data object share from the metastore. The caller must be an owner
	// of the share.
	Delete(ctx context.Context, request DeleteShareRequest) error

	// Deletes a data object share from the metastore. The caller must be an owner
	// of the share.
	DeleteByName(ctx context.Context, name string) error

	// Gets a data object share from the metastore. The caller must be a metastore
	// admin or the owner of the share.
	Get(ctx context.Context, request GetShareRequest) (*ShareInfo, error)

	// Gets a data object share from the metastore. The caller must be a metastore
	// admin or the owner of the share.
	GetByName(ctx context.Context, name string) (*ShareInfo, error)

	// Gets an array of data object shares from the metastore. The caller must be a
	// metastore admin or the owner of the share. There is no guarantee of a
	// specific ordering of the elements in the array.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListSharesRequest) listing.Iterator[ShareInfo]

	// Gets an array of data object shares from the metastore. The caller must be a
	// metastore admin or the owner of the share. There is no guarantee of a
	// specific ordering of the elements in the array.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListSharesRequest) ([]ShareInfo, error)

	// Gets the permissions for a data share from the metastore. The caller must be
	// a metastore admin or the owner of the share.
	SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetSharePermissionsResponse, error)

	// Gets the permissions for a data share from the metastore. The caller must be
	// a metastore admin or the owner of the share.
	SharePermissionsByName(ctx context.Context, name string) (*GetSharePermissionsResponse, error)

	// Updates the share with the changes and data objects in the request. The
	// caller must be the owner of the share or a metastore admin.
	//
	// When the caller is a metastore admin, only the __owner__ field can be
	// updated.
	//
	// In the case the share name is changed, **updateShare** requires that the
	// caller is the owner of the share and has the CREATE_SHARE privilege.
	//
	// If there are notebook files in the share, the __storage_root__ field cannot
	// be updated.
	//
	// For each table that is added through this method, the share owner must also
	// have **SELECT** privilege on the table. This privilege must be maintained
	// indefinitely for recipients to be able to access the table. Typically, you
	// should use a group as the share owner.
	//
	// Table removals through **update** do not require additional privileges.
	Update(ctx context.Context, request UpdateShare) (*ShareInfo, error)

	// Updates the permissions for a data share in the metastore. The caller must be
	// a metastore admin or an owner of the share.
	//
	// For new recipient grants, the user must also be the recipient owner or
	// metastore admin. recipient revocations do not require additional privileges.
	UpdatePermissions(ctx context.Context, request UpdateSharePermissions) (*UpdateSharePermissionsResponse, error)
}

func NewShares(client *client.DatabricksClient) *SharesAPI {
	return &SharesAPI{
		sharesImpl: sharesImpl{
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
	sharesImpl
}

// Deletes a data object share from the metastore. The caller must be an owner
// of the share.
func (a *SharesAPI) DeleteByName(ctx context.Context, name string) error {
	return a.sharesImpl.Delete(ctx, DeleteShareRequest{
		Name: name,
	})
}

// Gets a data object share from the metastore. The caller must be a metastore
// admin or the owner of the share.
func (a *SharesAPI) GetByName(ctx context.Context, name string) (*ShareInfo, error) {
	return a.sharesImpl.Get(ctx, GetShareRequest{
		Name: name,
	})
}

// Gets the permissions for a data share from the metastore. The caller must be
// a metastore admin or the owner of the share.
func (a *SharesAPI) SharePermissionsByName(ctx context.Context, name string) (*GetSharePermissionsResponse, error) {
	return a.sharesImpl.SharePermissions(ctx, SharePermissionsRequest{
		Name: name,
	})
}
