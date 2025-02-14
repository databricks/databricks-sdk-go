// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

// A data provider is an object representing the organization in the real world
// who shares the data. A provider contains shares which further contain the
// shared data.
type ProvidersService interface {

	// Create an auth provider.
	//
	// Creates a new authentication provider minimally based on a name and
	// authentication type. The caller must be an admin on the metastore.
	Create(ctx context.Context, request CreateProvider) (*ProviderInfo, error)

	// Delete a provider.
	//
	// Deletes an authentication provider, if the caller is a metastore admin or
	// is the owner of the provider.
	Delete(ctx context.Context, request DeleteProviderRequest) error

	// Get a provider.
	//
	// Gets a specific authentication provider. The caller must supply the name
	// of the provider, and must either be a metastore admin or the owner of the
	// provider.
	Get(ctx context.Context, request GetProviderRequest) (*ProviderInfo, error)

	// List providers.
	//
	// Gets an array of available authentication providers. The caller must
	// either be a metastore admin or the owner of the providers. Providers not
	// owned by the caller are not included in the response. There is no
	// guarantee of a specific ordering of the elements in the array.
	//
	// Use ListAll() to get all ProviderInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error)

	// List shares by Provider.
	//
	// Gets an array of a specified provider's shares within the metastore
	// where:
	//
	// * the caller is a metastore admin, or * the caller is the owner.
	//
	// Use ListSharesAll() to get all ProviderShare instances, which will iterate over every result page.
	ListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error)

	// Update a provider.
	//
	// Updates the information for an authentication provider, if the caller is
	// a metastore admin or is the owner of the provider. If the update changes
	// the provider name, the caller must be both a metastore admin and the
	// owner of the provider.
	Update(ctx context.Context, request UpdateProvider) (*ProviderInfo, error)
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
type RecipientActivationService interface {

	// Get a share activation URL.
	//
	// Gets an activation URL for a share.
	GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error

	// Get an access token.
	//
	// Retrieve access token with an activation url. This is a public API
	// without any authentication.
	RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error)
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
type RecipientsService interface {

	// Create a share recipient.
	//
	// Creates a new recipient with the delta sharing authentication type in the
	// metastore. The caller must be a metastore admin or have the
	// **CREATE_RECIPIENT** privilege on the metastore.
	Create(ctx context.Context, request CreateRecipient) (*RecipientInfo, error)

	// Delete a share recipient.
	//
	// Deletes the specified recipient from the metastore. The caller must be
	// the owner of the recipient.
	Delete(ctx context.Context, request DeleteRecipientRequest) error

	// Get a share recipient.
	//
	// Gets a share recipient from the metastore if:
	//
	// * the caller is the owner of the share recipient, or: * is a metastore
	// admin
	Get(ctx context.Context, request GetRecipientRequest) (*RecipientInfo, error)

	// List share recipients.
	//
	// Gets an array of all share recipients within the current metastore where:
	//
	// * the caller is a metastore admin, or * the caller is the owner. There is
	// no guarantee of a specific ordering of the elements in the array.
	//
	// Use ListAll() to get all RecipientInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListRecipientsRequest) (*ListRecipientsResponse, error)

	// Rotate a token.
	//
	// Refreshes the specified recipient's delta sharing authentication token
	// with the provided token info. The caller must be the owner of the
	// recipient.
	RotateToken(ctx context.Context, request RotateRecipientToken) (*RecipientInfo, error)

	// Get recipient share permissions.
	//
	// Gets the share permissions for the specified Recipient. The caller must
	// be a metastore admin or the owner of the Recipient.
	SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error)

	// Update a share recipient.
	//
	// Updates an existing recipient in the metastore. The caller must be a
	// metastore admin or the owner of the recipient. If the recipient name will
	// be updated, the user must be both a metastore admin and the owner of the
	// recipient.
	Update(ctx context.Context, request UpdateRecipient) (*RecipientInfo, error)
}

// A share is a container instantiated with :method:shares/create. Once created
// you can iteratively register a collection of existing data assets defined
// within the metastore using :method:shares/update. You can register data
// assets under their original name, qualified by their original schema, or
// provide alternate exposed names.
type SharesService interface {

	// Create a share.
	//
	// Creates a new share for data objects. Data objects can be added after
	// creation with **update**. The caller must be a metastore admin or have
	// the **CREATE_SHARE** privilege on the metastore.
	Create(ctx context.Context, request CreateShare) (*ShareInfo, error)

	// Delete a share.
	//
	// Deletes a data object share from the metastore. The caller must be an
	// owner of the share.
	Delete(ctx context.Context, request DeleteShareRequest) error

	// Get a share.
	//
	// Gets a data object share from the metastore. The caller must be a
	// metastore admin or the owner of the share.
	Get(ctx context.Context, request GetShareRequest) (*ShareInfo, error)

	// List shares.
	//
	// Gets an array of data object shares from the metastore. The caller must
	// be a metastore admin or the owner of the share. There is no guarantee of
	// a specific ordering of the elements in the array.
	//
	// Use ListAll() to get all ShareInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListSharesRequest) (*ListSharesResponse, error)

	// Get permissions.
	//
	// Gets the permissions for a data share from the metastore. The caller must
	// be a metastore admin or the owner of the share.
	SharePermissions(ctx context.Context, request SharePermissionsRequest) (*catalog.PermissionsList, error)

	// Update a share.
	//
	// Updates the share with the changes and data objects in the request. The
	// caller must be the owner of the share or a metastore admin.
	//
	// When the caller is a metastore admin, only the __owner__ field can be
	// updated.
	//
	// In the case that the share name is changed, **updateShare** requires that
	// the caller is both the share owner and a metastore admin.
	//
	// If there are notebook files in the share, the __storage_root__ field
	// cannot be updated.
	//
	// For each table that is added through this method, the share owner must
	// also have **SELECT** privilege on the table. This privilege must be
	// maintained indefinitely for recipients to be able to access the table.
	// Typically, you should use a group as the share owner.
	//
	// Table removals through **update** do not require additional privileges.
	Update(ctx context.Context, request UpdateShare) (*ShareInfo, error)

	// Update permissions.
	//
	// Updates the permissions for a data share in the metastore. The caller
	// must be a metastore admin or an owner of the share.
	//
	// For new recipient grants, the user must also be the owner of the
	// recipients. recipient revocations do not require additional privileges.
	UpdatePermissions(ctx context.Context, request UpdateSharePermissions) error
}
