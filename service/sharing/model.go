// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

// all definitions in this file are in alphabetical order

// The delta sharing authentication type.
type AuthenticationType string

const AuthenticationTypeDatabricks AuthenticationType = `DATABRICKS`

const AuthenticationTypeToken AuthenticationType = `TOKEN`

// String representation for [fmt.Print]
func (f *AuthenticationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AuthenticationType) Set(v string) error {
	switch v {
	case `DATABRICKS`, `TOKEN`:
		*f = AuthenticationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS", "TOKEN"`, v)
	}
}

// Type always returns AuthenticationType to satisfy [pflag.Value] interface
func (f *AuthenticationType) Type() string {
	return "AuthenticationType"
}

type CreateProvider struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// The name of the Provider.
	Name string `json:"name"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
}

type CreateRecipient struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is required when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId any `json:"data_recipient_global_metastore_id,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Name of Recipient.
	Name string `json:"name"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// required when the __authentication_type__ is **DATABRICKS**.
	SharingCode string `json:"sharing_code,omitempty"`
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the share.
	Name string `json:"name"`
}

// Delete a provider
type DeleteProviderRequest struct {
	// Name of the provider.
	Name string `json:"-" url:"-"`
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name string `json:"-" url:"-"`
}

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name string `json:"-" url:"-"`
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `json:"-" url:"-"`
}

// Get a provider
type GetProviderRequest struct {
	// Name of the provider.
	Name string `json:"-" url:"-"`
}

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	Name string `json:"-" url:"-"`
}

type GetRecipientSharePermissionsResponse struct {
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData bool `json:"-" url:"include_shared_data,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}

type ListProviderSharesResponse struct {
	// An array of provider shares.
	Shares []ProviderShare `json:"shares,omitempty"`
}

// List providers
type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId string `json:"-" url:"data_provider_global_metastore_id,omitempty"`
}

type ListProvidersResponse struct {
	// An array of provider information objects.
	Providers []ProviderInfo `json:"providers,omitempty"`
}

// List share recipients
type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId string `json:"-" url:"data_recipient_global_metastore_id,omitempty"`
}

type ListRecipientsResponse struct {
	// An array of recipient information objects.
	Recipients []RecipientInfo `json:"recipients,omitempty"`
}

// List shares by Provider
type ListSharesRequest struct {
	// Name of the provider in which to list shares.
	Name string `json:"-" url:"-"`
}

type ListSharesResponse struct {
	// An array of data share information objects.
	Shares []ShareInfo `json:"shares,omitempty"`
}

type Partition struct {
	// An array of partition values.
	Values []PartitionValue `json:"values,omitempty"`
}

type PartitionValue struct {
	// The name of the partition column.
	Name string `json:"name,omitempty"`
	// The operator to apply for the value.
	Op PartitionValueOp `json:"op,omitempty"`
	// The key of a Delta Sharing recipient's property. For example
	// `databricks-account-id`. When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey string `json:"recipient_property_key,omitempty"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value string `json:"value,omitempty"`
}

// The operator to apply for the value.
type PartitionValueOp string

const PartitionValueOpEqual PartitionValueOp = `EQUAL`

const PartitionValueOpLike PartitionValueOp = `LIKE`

// String representation for [fmt.Print]
func (f *PartitionValueOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PartitionValueOp) Set(v string) error {
	switch v {
	case `EQUAL`, `LIKE`:
		*f = PartitionValueOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "LIKE"`, v)
	}
}

// Type always returns PartitionValueOp to satisfy [pflag.Value] interface
func (f *PartitionValueOp) Type() string {
	return "PartitionValueOp"
}

type Privilege string

const PrivilegeAllPrivileges Privilege = `ALL_PRIVILEGES`

const PrivilegeCreate Privilege = `CREATE`

const PrivilegeCreateCatalog Privilege = `CREATE_CATALOG`

const PrivilegeCreateExternalLocation Privilege = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable Privilege = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateFunction Privilege = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage Privilege = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView Privilege = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateProvider Privilege = `CREATE_PROVIDER`

const PrivilegeCreateRecipient Privilege = `CREATE_RECIPIENT`

const PrivilegeCreateSchema Privilege = `CREATE_SCHEMA`

const PrivilegeCreateShare Privilege = `CREATE_SHARE`

const PrivilegeCreateStorageCredential Privilege = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable Privilege = `CREATE_TABLE`

const PrivilegeCreateView Privilege = `CREATE_VIEW`

const PrivilegeExecute Privilege = `EXECUTE`

const PrivilegeModify Privilege = `MODIFY`

const PrivilegeReadFiles Privilege = `READ_FILES`

const PrivilegeReadPrivateFiles Privilege = `READ_PRIVATE_FILES`

const PrivilegeRefresh Privilege = `REFRESH`

const PrivilegeSelect Privilege = `SELECT`

const PrivilegeSetSharePermission Privilege = `SET_SHARE_PERMISSION`

const PrivilegeUsage Privilege = `USAGE`

const PrivilegeUseCatalog Privilege = `USE_CATALOG`

const PrivilegeUseProvider Privilege = `USE_PROVIDER`

const PrivilegeUseRecipient Privilege = `USE_RECIPIENT`

const PrivilegeUseSchema Privilege = `USE_SCHEMA`

const PrivilegeUseShare Privilege = `USE_SHARE`

const PrivilegeWriteFiles Privilege = `WRITE_FILES`

const PrivilegeWritePrivateFiles Privilege = `WRITE_PRIVATE_FILES`

// String representation for [fmt.Print]
func (f *Privilege) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Privilege) Set(v string) error {
	switch v {
	case `ALL_PRIVILEGES`, `CREATE`, `CREATE_CATALOG`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `EXECUTE`, `MODIFY`, `READ_FILES`, `READ_PRIVATE_FILES`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`:
		*f = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_PRIVILEGES", "CREATE", "CREATE_CATALOG", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "EXECUTE", "MODIFY", "READ_FILES", "READ_PRIVATE_FILES", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES"`, v)
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	Privileges []Privilege `json:"privileges,omitempty"`
}

type ProviderInfo struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type,omitempty"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud string `json:"cloud,omitempty"`
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId string `json:"data_provider_global_metastore_id,omitempty"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the Provider.
	Name string `json:"name,omitempty"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// This field is only present when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region string `json:"region,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified Share.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type ProviderShare struct {
	// The name of the Provider Share.
	Name string `json:"name,omitempty"`
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type,omitempty"`
	// Cloud vendor of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**`.
	Cloud string `json:"cloud,omitempty"`
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId any `json:"data_recipient_global_metastore_id,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Unique identifier of recipient's Unity Catalog metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Recipient.
	Name string `json:"name,omitempty"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region string `json:"region,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode string `json:"sharing_code,omitempty"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken string `json:"bearer_token,omitempty"`
	// The endpoint for the share to be used by the recipient.
	Endpoint string `json:"endpoint,omitempty"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion int `json:"share_credentials_version,omitempty"`
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// Time at which this recipient Token was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient token creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Unique ID of the recipient token.
	Id string `json:"id,omitempty"`
	// Time at which this recipient Token was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of recipient Token updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Get an access token
type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `json:"-" url:"-"`
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken string `json:"bearerToken,omitempty"`
	// The endpoint for the share to be used by the recipient.
	Endpoint string `json:"endpoint,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds int64 `json:"existing_token_expire_in_seconds"`
	// The name of the recipient.
	Name string `json:"-" url:"-"`
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties"`
}

// A map of key-value properties attached to the securable.
type SecurablePropertiesMap map[string]string

type ShareInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of share creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Name of the share.
	Name string `json:"name,omitempty"`
	// A list of shared data objects within the share.
	Objects []SharedDataObject `json:"objects,omitempty"`
	// Username of current owner of share.
	Owner string `json:"owner,omitempty"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of share updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Get recipient share permissions
type SharePermissionsRequest struct {
	// The name of the Recipient.
	Name string `json:"-" url:"-"`
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
	// The share name.
	ShareName string `json:"share_name,omitempty"`
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt int64 `json:"added_at,omitempty"`
	// Username of the sharer.
	AddedBy string `json:"added_by,omitempty"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled bool `json:"cdf_enabled,omitempty"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment string `json:"comment,omitempty"`
	// The type of the data object.
	DataObjectType string `json:"data_object_type,omitempty"`
	// A fully qualified name that uniquely identifies a data object.
	//
	// For example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`.
	Name string `json:"name"`
	// Array of partitions for the shared data.
	Partitions []Partition `json:"partitions,omitempty"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs string `json:"shared_as,omitempty"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion int64 `json:"start_version,omitempty"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status SharedDataObjectStatus `json:"status,omitempty"`
}

// One of: **ACTIVE**, **PERMISSION_DENIED**.
type SharedDataObjectStatus string

const SharedDataObjectStatusActive SharedDataObjectStatus = `ACTIVE`

const SharedDataObjectStatusPermissionDenied SharedDataObjectStatus = `PERMISSION_DENIED`

// String representation for [fmt.Print]
func (f *SharedDataObjectStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `PERMISSION_DENIED`:
		*f = SharedDataObjectStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "PERMISSION_DENIED"`, v)
	}
}

// Type always returns SharedDataObjectStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectStatus) Type() string {
	return "SharedDataObjectStatus"
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action SharedDataObjectUpdateAction `json:"action,omitempty"`
	// The data object that is being added, removed, or updated.
	DataObject *SharedDataObject `json:"data_object,omitempty"`
}

// One of: **ADD**, **REMOVE**, **UPDATE**.
type SharedDataObjectUpdateAction string

const SharedDataObjectUpdateActionAdd SharedDataObjectUpdateAction = `ADD`

const SharedDataObjectUpdateActionRemove SharedDataObjectUpdateAction = `REMOVE`

const SharedDataObjectUpdateActionUpdate SharedDataObjectUpdateAction = `UPDATE`

// String representation for [fmt.Print]
func (f *SharedDataObjectUpdateAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectUpdateAction) Set(v string) error {
	switch v {
	case `ADD`, `REMOVE`, `UPDATE`:
		*f = SharedDataObjectUpdateAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADD", "REMOVE", "UPDATE"`, v)
	}
}

// Type always returns SharedDataObjectUpdateAction to satisfy [pflag.Value] interface
func (f *SharedDataObjectUpdateAction) Type() string {
	return "SharedDataObjectUpdateAction"
}

type UpdateProvider struct {
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// The name of the Provider.
	Name string `json:"name,omitempty" url:"-"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
}

type UpdateRecipient struct {
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Name of Recipient.
	Name string `json:"name,omitempty" url:"-"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the share.
	Name string `json:"name,omitempty" url:"-"`
	// Username of current owner of share.
	Owner string `json:"owner,omitempty"`
	// Array of shared data object updates.
	Updates []SharedDataObjectUpdate `json:"updates,omitempty"`
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	Changes []catalog.PermissionsChange `json:"changes,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`
}
