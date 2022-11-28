// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

// all definitions in this file are in alphabetical order

type AwsIamRole struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem.
	//
	// [Create:IGN].
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	// [Create:REQ].
	RoleArn string `json:"role_arn,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	//
	// [Create:IGN].
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`
}

type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant. [Create:REQ]
	ApplicationId string `json:"application_id,omitempty"`
	// The client secret generated for the above app ID in AAD. [Create:REQ]
	ClientSecret string `json:"client_secret,omitempty"`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application. [Create:REQ].
	DirectoryId string `json:"directory_id,omitempty"`
}

type CatalogInfo struct {
	// [Create,Update:IGN] The type of the catalog.
	CatalogType CatalogInfoCatalogType `json:"catalog_type,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Catalog was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Catalog creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN,Update:OPT] Username of current owner of Catalog.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Catalog.
	Privileges []CatalogInfoPrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// Delta Sharing Catalog specific fields. A Delta Sharing Catalog is a
	// catalog that is based on a Delta share on a remote sharing server.
	// [Create:OPT,Update:IGN] The name of delta sharing provider.
	ProviderName string `json:"provider_name,omitempty"`
	// [Create:OPT,Update: IGN] The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// [Create,Update:IGN] Time at which this Catalog was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Catalog.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create,Update:IGN] The type of the catalog.
type CatalogInfoCatalogType string

const CatalogInfoCatalogTypeDeltasharingCatalog CatalogInfoCatalogType = `DELTASHARING_CATALOG`

const CatalogInfoCatalogTypeManagedCatalog CatalogInfoCatalogType = `MANAGED_CATALOG`

const CatalogInfoCatalogTypeSystemCatalog CatalogInfoCatalogType = `SYSTEM_CATALOG`

const CatalogInfoCatalogTypeUnknownCatalogType CatalogInfoCatalogType = `UNKNOWN_CATALOG_TYPE`

type CatalogInfoPrivilegesItem string

const CatalogInfoPrivilegesItemCreate CatalogInfoPrivilegesItem = `CREATE`

const CatalogInfoPrivilegesItemCreateMount CatalogInfoPrivilegesItem = `CREATE_MOUNT`

const CatalogInfoPrivilegesItemCreateTable CatalogInfoPrivilegesItem = `CREATE_TABLE`

const CatalogInfoPrivilegesItemModify CatalogInfoPrivilegesItem = `MODIFY`

const CatalogInfoPrivilegesItemReadFiles CatalogInfoPrivilegesItem = `READ_FILES`

const CatalogInfoPrivilegesItemSelect CatalogInfoPrivilegesItem = `SELECT`

const CatalogInfoPrivilegesItemUnknownPrivilege CatalogInfoPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CatalogInfoPrivilegesItemUsage CatalogInfoPrivilegesItem = `USAGE`

const CatalogInfoPrivilegesItemWriteFiles CatalogInfoPrivilegesItem = `WRITE_FILES`

type ColumnInfo struct {
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create:REQ Update:OPT] Name of Column.
	Name string `json:"name,omitempty"`
	// [Create,Update:OPT] Whether field may be Null (default: True).
	Nullable bool `json:"nullable,omitempty"`
	// [Create,Update:OPT] Partition index for column.
	PartitionIndex int `json:"partition_index,omitempty"`
	// [Create:REQ Update:OPT] Ordinal position of column (starting at position
	// 0).
	Position int `json:"position,omitempty"`
	// [Create: OPT, Update: OPT] Format of IntervalType.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// [Create:OPT Update:OPT] Full data type spec, JSON-serialized.
	TypeJson string `json:"type_json,omitempty"`
	// [Create: REQ Update: OPT] Name of type (INT, STRUCT, MAP, etc.)
	TypeName ColumnInfoTypeName `json:"type_name,omitempty"`
	// [Create: OPT, Update: OPT] Digits of precision; required on Create for
	// DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// [Create: OPT, Update: OPT] Digits to right of decimal; Required on Create
	// for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// [Create:REQ Update:OPT] Full data type spec, SQL/catalogString text.
	TypeText string `json:"type_text,omitempty"`
}

// [Create: REQ Update: OPT] Name of type (INT, STRUCT, MAP, etc.)
type ColumnInfoTypeName string

const ColumnInfoTypeNameArray ColumnInfoTypeName = `ARRAY`

const ColumnInfoTypeNameBinary ColumnInfoTypeName = `BINARY`

const ColumnInfoTypeNameBoolean ColumnInfoTypeName = `BOOLEAN`

const ColumnInfoTypeNameByte ColumnInfoTypeName = `BYTE`

const ColumnInfoTypeNameChar ColumnInfoTypeName = `CHAR`

const ColumnInfoTypeNameDate ColumnInfoTypeName = `DATE`

const ColumnInfoTypeNameDecimal ColumnInfoTypeName = `DECIMAL`

const ColumnInfoTypeNameDouble ColumnInfoTypeName = `DOUBLE`

const ColumnInfoTypeNameFloat ColumnInfoTypeName = `FLOAT`

const ColumnInfoTypeNameInt ColumnInfoTypeName = `INT`

const ColumnInfoTypeNameInterval ColumnInfoTypeName = `INTERVAL`

const ColumnInfoTypeNameLong ColumnInfoTypeName = `LONG`

const ColumnInfoTypeNameMap ColumnInfoTypeName = `MAP`

const ColumnInfoTypeNameNull ColumnInfoTypeName = `NULL`

const ColumnInfoTypeNameShort ColumnInfoTypeName = `SHORT`

const ColumnInfoTypeNameString ColumnInfoTypeName = `STRING`

const ColumnInfoTypeNameStruct ColumnInfoTypeName = `STRUCT`

const ColumnInfoTypeNameTimestamp ColumnInfoTypeName = `TIMESTAMP`

const ColumnInfoTypeNameUnknownColumnTypeName ColumnInfoTypeName = `UNKNOWN_COLUMN_TYPE_NAME`

type CreateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of Catalog
	Name string `json:"name"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// Delta Sharing Catalog specific fields. A Delta Sharing Catalog is a
	// catalog that is based on a Delta share on a remote sharing server.
	// [Create:OPT,Update:IGN] The name of delta sharing provider.
	ProviderName string `json:"provider_name,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
}

type CreateCatalogResponse struct {
	// [Create,Update:IGN] The type of the catalog.
	CatalogType CreateCatalogResponseCatalogType `json:"catalog_type,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Catalog was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Catalog creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN,Update:OPT] Username of current owner of Catalog.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Catalog.
	Privileges []CreateCatalogResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// Delta Sharing Catalog specific fields. A Delta Sharing Catalog is a
	// catalog that is based on a Delta share on a remote sharing server.
	// [Create:OPT,Update:IGN] The name of delta sharing provider.
	ProviderName string `json:"provider_name,omitempty"`
	// [Create:OPT,Update: IGN] The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// [Create,Update:IGN] Time at which this Catalog was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Catalog.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create,Update:IGN] The type of the catalog.
type CreateCatalogResponseCatalogType string

const CreateCatalogResponseCatalogTypeDeltasharingCatalog CreateCatalogResponseCatalogType = `DELTASHARING_CATALOG`

const CreateCatalogResponseCatalogTypeManagedCatalog CreateCatalogResponseCatalogType = `MANAGED_CATALOG`

const CreateCatalogResponseCatalogTypeSystemCatalog CreateCatalogResponseCatalogType = `SYSTEM_CATALOG`

const CreateCatalogResponseCatalogTypeUnknownCatalogType CreateCatalogResponseCatalogType = `UNKNOWN_CATALOG_TYPE`

type CreateCatalogResponsePrivilegesItem string

const CreateCatalogResponsePrivilegesItemCreate CreateCatalogResponsePrivilegesItem = `CREATE`

const CreateCatalogResponsePrivilegesItemCreateMount CreateCatalogResponsePrivilegesItem = `CREATE_MOUNT`

const CreateCatalogResponsePrivilegesItemCreateTable CreateCatalogResponsePrivilegesItem = `CREATE_TABLE`

const CreateCatalogResponsePrivilegesItemModify CreateCatalogResponsePrivilegesItem = `MODIFY`

const CreateCatalogResponsePrivilegesItemReadFiles CreateCatalogResponsePrivilegesItem = `READ_FILES`

const CreateCatalogResponsePrivilegesItemSelect CreateCatalogResponsePrivilegesItem = `SELECT`

const CreateCatalogResponsePrivilegesItemUnknownPrivilege CreateCatalogResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateCatalogResponsePrivilegesItemUsage CreateCatalogResponsePrivilegesItem = `USAGE`

const CreateCatalogResponsePrivilegesItemWriteFiles CreateCatalogResponsePrivilegesItem = `WRITE_FILES`

type CreateExternalLocation struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location's Storage Credential.
	CredentialId string `json:"credential_id,omitempty"`
	// [Create:REQ Update:OPT] Current name of the Storage Credential this
	// location uses.
	CredentialName string `json:"credential_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore hosting the External
	// Location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of the External Location.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the External Location.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Time at which this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the External
	// Location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create:REQ Update:OPT] Path URL of the External Location.
	Url string `json:"url,omitempty"`
}

type CreateExternalLocationResponse struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location's Storage Credential.
	CredentialId string `json:"credential_id,omitempty"`
	// [Create:REQ Update:OPT] Current name of the Storage Credential this
	// location uses.
	CredentialName string `json:"credential_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore hosting the External
	// Location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of the External Location.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the External Location.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Time at which this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the External
	// Location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create:REQ Update:OPT] Path URL of the External Location.
	Url string `json:"url,omitempty"`
}

type CreateMetastore struct {
	// [Create,Update:IGN] Time at which this Metastore was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:IGN Update:OPT] Unique identifier of (Default) Data Access
	// Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// [Create:IGN Update:OPT] Whether Delta Sharing is enabled on this
	// metastore.
	DeltaSharingEnabled bool `json:"delta_sharing_enabled,omitempty"`
	// [Create:IGN Update:OPT] The lifetime of delta sharing recipient token in
	// seconds
	DeltaSharingRecipientTokenLifetimeInSeconds int `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Metastore.
	Privileges []CreateMetastorePrivilegesItem `json:"privileges,omitempty"`
	// The region this metastore has an afinity to. This is used by
	// accounts-manager. Ignored by Unity Catalog.
	Region string `json:"region,omitempty"`
	// [Create:REQ Update:ERR] Storage root URL for Metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// [Create:IGN Update:OPT] UUID of storage credential to access storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// [Create,Update:IGN] Time at which the Metastore was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Metastore.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CreateMetastoreAssignment struct {
	// THe name of the default catalog in the Metastore.
	DefaultCatalogName string `json:"default_catalog_name"`
	// The ID of the Metastore.
	MetastoreId string `json:"metastore_id"`
	// A workspace ID.
	WorkspaceId int `json:"-" url:"-"`
}

type CreateMetastorePrivilegesItem string

const CreateMetastorePrivilegesItemCreate CreateMetastorePrivilegesItem = `CREATE`

const CreateMetastorePrivilegesItemCreateMount CreateMetastorePrivilegesItem = `CREATE_MOUNT`

const CreateMetastorePrivilegesItemCreateTable CreateMetastorePrivilegesItem = `CREATE_TABLE`

const CreateMetastorePrivilegesItemModify CreateMetastorePrivilegesItem = `MODIFY`

const CreateMetastorePrivilegesItemReadFiles CreateMetastorePrivilegesItem = `READ_FILES`

const CreateMetastorePrivilegesItemSelect CreateMetastorePrivilegesItem = `SELECT`

const CreateMetastorePrivilegesItemUnknownPrivilege CreateMetastorePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateMetastorePrivilegesItemUsage CreateMetastorePrivilegesItem = `USAGE`

const CreateMetastorePrivilegesItemWriteFiles CreateMetastorePrivilegesItem = `WRITE_FILES`

type CreateMetastoreResponse struct {
	// [Create,Update:IGN] Time at which this Metastore was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:IGN Update:OPT] Unique identifier of (Default) Data Access
	// Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// [Create:IGN Update:OPT] Whether Delta Sharing is enabled on this
	// metastore.
	DeltaSharingEnabled bool `json:"delta_sharing_enabled,omitempty"`
	// [Create:IGN Update:OPT] The lifetime of delta sharing recipient token in
	// seconds
	DeltaSharingRecipientTokenLifetimeInSeconds int `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Metastore.
	Privileges []CreateMetastoreResponsePrivilegesItem `json:"privileges,omitempty"`
	// The region this metastore has an afinity to. This is used by
	// accounts-manager. Ignored by Unity Catalog.
	Region string `json:"region,omitempty"`
	// [Create:REQ Update:ERR] Storage root URL for Metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// [Create:IGN Update:OPT] UUID of storage credential to access storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// [Create,Update:IGN] Time at which the Metastore was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Metastore.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CreateMetastoreResponsePrivilegesItem string

const CreateMetastoreResponsePrivilegesItemCreate CreateMetastoreResponsePrivilegesItem = `CREATE`

const CreateMetastoreResponsePrivilegesItemCreateMount CreateMetastoreResponsePrivilegesItem = `CREATE_MOUNT`

const CreateMetastoreResponsePrivilegesItemCreateTable CreateMetastoreResponsePrivilegesItem = `CREATE_TABLE`

const CreateMetastoreResponsePrivilegesItemModify CreateMetastoreResponsePrivilegesItem = `MODIFY`

const CreateMetastoreResponsePrivilegesItemReadFiles CreateMetastoreResponsePrivilegesItem = `READ_FILES`

const CreateMetastoreResponsePrivilegesItemSelect CreateMetastoreResponsePrivilegesItem = `SELECT`

const CreateMetastoreResponsePrivilegesItemUnknownPrivilege CreateMetastoreResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateMetastoreResponsePrivilegesItemUsage CreateMetastoreResponsePrivilegesItem = `USAGE`

const CreateMetastoreResponsePrivilegesItemWriteFiles CreateMetastoreResponsePrivilegesItem = `WRITE_FILES`

type CreateProvider struct {
	// [Create,Update:IGN] Whether this provider is successfully activated by
	// the data provider. This field is only present when the authentication
	// type is DATABRICKS.
	ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType CreateProviderAuthenticationType `json:"authentication_type,omitempty"`
	// [Create,Update:OPT] Description about the provider.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:REQ] The name of the Provider.
	Name string `json:"name,omitempty"`
	// [Create,Update:IGN] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// [Create,Update:OPT] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// [Create,Update:IGN] The server-generated one-time sharing code. This
	// field is only present when the authentication type is DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Share.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type CreateProviderAuthenticationType string

const CreateProviderAuthenticationTypeDatabricks CreateProviderAuthenticationType = `DATABRICKS`

const CreateProviderAuthenticationTypeToken CreateProviderAuthenticationType = `TOKEN`

const CreateProviderAuthenticationTypeUnknown CreateProviderAuthenticationType = `UNKNOWN`

type CreateProviderResponse struct {
	// [Create,Update:IGN] Whether this provider is successfully activated by
	// the data provider. This field is only present when the authentication
	// type is DATABRICKS.
	ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType CreateProviderResponseAuthenticationType `json:"authentication_type,omitempty"`
	// [Create,Update:OPT] Description about the provider.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:REQ] The name of the Provider.
	Name string `json:"name,omitempty"`
	// [Create,Update:IGN] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// [Create,Update:OPT] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// [Create,Update:IGN] The server-generated one-time sharing code. This
	// field is only present when the authentication type is DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Share.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type CreateProviderResponseAuthenticationType string

const CreateProviderResponseAuthenticationTypeDatabricks CreateProviderResponseAuthenticationType = `DATABRICKS`

const CreateProviderResponseAuthenticationTypeToken CreateProviderResponseAuthenticationType = `TOKEN`

const CreateProviderResponseAuthenticationTypeUnknown CreateProviderResponseAuthenticationType = `UNKNOWN`

type CreateRecipient struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient's activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType CreateRecipientAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which the recipient was updated, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type CreateRecipientAuthenticationType string

const CreateRecipientAuthenticationTypeDatabricks CreateRecipientAuthenticationType = `DATABRICKS`

const CreateRecipientAuthenticationTypeToken CreateRecipientAuthenticationType = `TOKEN`

const CreateRecipientAuthenticationTypeUnknown CreateRecipientAuthenticationType = `UNKNOWN`

type CreateRecipientResponse struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient's activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType CreateRecipientResponseAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which the recipient was updated, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type CreateRecipientResponseAuthenticationType string

const CreateRecipientResponseAuthenticationTypeDatabricks CreateRecipientResponseAuthenticationType = `DATABRICKS`

const CreateRecipientResponseAuthenticationTypeToken CreateRecipientResponseAuthenticationType = `TOKEN`

const CreateRecipientResponseAuthenticationTypeUnknown CreateRecipientResponseAuthenticationType = `UNKNOWN`

type CreateSchema struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Schema creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Full name of Schema, in form of
	// <catalog_name>.<schema_name>.
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []CreateSchemaPrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CreateSchemaPrivilegesItem string

const CreateSchemaPrivilegesItemCreate CreateSchemaPrivilegesItem = `CREATE`

const CreateSchemaPrivilegesItemCreateMount CreateSchemaPrivilegesItem = `CREATE_MOUNT`

const CreateSchemaPrivilegesItemCreateTable CreateSchemaPrivilegesItem = `CREATE_TABLE`

const CreateSchemaPrivilegesItemModify CreateSchemaPrivilegesItem = `MODIFY`

const CreateSchemaPrivilegesItemReadFiles CreateSchemaPrivilegesItem = `READ_FILES`

const CreateSchemaPrivilegesItemSelect CreateSchemaPrivilegesItem = `SELECT`

const CreateSchemaPrivilegesItemUnknownPrivilege CreateSchemaPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateSchemaPrivilegesItemUsage CreateSchemaPrivilegesItem = `USAGE`

const CreateSchemaPrivilegesItemWriteFiles CreateSchemaPrivilegesItem = `WRITE_FILES`

type CreateSchemaResponse struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Schema creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Full name of Schema, in form of
	// <catalog_name>.<schema_name>.
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []CreateSchemaResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CreateSchemaResponsePrivilegesItem string

const CreateSchemaResponsePrivilegesItemCreate CreateSchemaResponsePrivilegesItem = `CREATE`

const CreateSchemaResponsePrivilegesItemCreateMount CreateSchemaResponsePrivilegesItem = `CREATE_MOUNT`

const CreateSchemaResponsePrivilegesItemCreateTable CreateSchemaResponsePrivilegesItem = `CREATE_TABLE`

const CreateSchemaResponsePrivilegesItemModify CreateSchemaResponsePrivilegesItem = `MODIFY`

const CreateSchemaResponsePrivilegesItemReadFiles CreateSchemaResponsePrivilegesItem = `READ_FILES`

const CreateSchemaResponsePrivilegesItemSelect CreateSchemaResponsePrivilegesItem = `SELECT`

const CreateSchemaResponsePrivilegesItemUnknownPrivilege CreateSchemaResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateSchemaResponsePrivilegesItemUsage CreateSchemaResponsePrivilegesItem = `USAGE`

const CreateSchemaResponsePrivilegesItemWriteFiles CreateSchemaResponsePrivilegesItem = `WRITE_FILES`

type CreateShare struct {
	// [Create: OPT] comment when creating the share
	Comment string `json:"comment,omitempty"`
	// [Create:IGN] Time at which this Share was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN] Username of Share creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:REQ] Name of the Share.
	Name string `json:"name,omitempty"`
	// [Create: IGN] A list of shared data objects within the Share.
	Objects []SharedDataObject `json:"objects,omitempty"`
}

type CreateShareResponse struct {
	// [Create: OPT] comment when creating the share
	Comment string `json:"comment,omitempty"`
	// [Create:IGN] Time at which this Share was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN] Username of Share creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:REQ] Name of the Share.
	Name string `json:"name,omitempty"`
	// [Create: IGN] A list of shared data objects within the Share.
	Objects []SharedDataObject `json:"objects,omitempty"`
}

type CreateStagingTable struct {
	// [Create:REQ] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// [Create:IGN] Unique id generated for the staging table
	Id string `json:"id,omitempty"`
	// [Create:REQ] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create:REQ] Name of parent Schema relative to its parent Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// [Create:IGN] URI generated for the staging table
	StagingLocation string `json:"staging_location,omitempty"`
}

type CreateStagingTableResponse struct {
	// [Create:REQ] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// [Create:IGN] Unique id generated for the staging table
	Id string `json:"id,omitempty"`
	// [Create:REQ] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create:REQ] Name of parent Schema relative to its parent Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// [Create:IGN] URI generated for the staging table
	StagingLocation string `json:"staging_location,omitempty"`
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// [Create:IGN] The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ, Update:OPT] The credential name. The name MUST be unique
	// within the Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// [Create,Update:IGN] Time at which this credential was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CreateStorageCredentialResponse struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// [Create:IGN] The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ, Update:OPT] The credential name. The name MUST be unique
	// within the Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Time at which this credential was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CreateTable struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name ('columns') is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Columns []ColumnInfo `json:"columns,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Table was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the data_access_configuration to use.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
	DataSourceFormat CreateTableDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// <catalog_name>.<schema_name>.<table_name>
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []CreateTablePrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create:REQ Update:IGN] Name of parent Schema relative to its parent
	// Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// [Create,Update:OPT] List of schemes whose objects can be referenced
	// without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// [Create:OPT Update:IGN] Name of the storage credential this table used
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, EXTERNAL
	// tables)
	StorageLocation string `json:"storage_location,omitempty"`
	// [Create:IGN Update:IGN] Name of Table, relative to parent Schema.
	TableId string `json:"table_id,omitempty"`
	// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
	TableType CreateTableTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == "VIEW")
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
type CreateTableDataSourceFormat string

const CreateTableDataSourceFormatAvro CreateTableDataSourceFormat = `AVRO`

const CreateTableDataSourceFormatCsv CreateTableDataSourceFormat = `CSV`

const CreateTableDataSourceFormatDelta CreateTableDataSourceFormat = `DELTA`

const CreateTableDataSourceFormatDeltasharing CreateTableDataSourceFormat = `DELTASHARING`

const CreateTableDataSourceFormatJson CreateTableDataSourceFormat = `JSON`

const CreateTableDataSourceFormatOrc CreateTableDataSourceFormat = `ORC`

const CreateTableDataSourceFormatParquet CreateTableDataSourceFormat = `PARQUET`

const CreateTableDataSourceFormatText CreateTableDataSourceFormat = `TEXT`

const CreateTableDataSourceFormatUnityCatalog CreateTableDataSourceFormat = `UNITY_CATALOG`

const CreateTableDataSourceFormatUnknownDataSourceFormat CreateTableDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

type CreateTablePrivilegesItem string

const CreateTablePrivilegesItemCreate CreateTablePrivilegesItem = `CREATE`

const CreateTablePrivilegesItemCreateMount CreateTablePrivilegesItem = `CREATE_MOUNT`

const CreateTablePrivilegesItemCreateTable CreateTablePrivilegesItem = `CREATE_TABLE`

const CreateTablePrivilegesItemModify CreateTablePrivilegesItem = `MODIFY`

const CreateTablePrivilegesItemReadFiles CreateTablePrivilegesItem = `READ_FILES`

const CreateTablePrivilegesItemSelect CreateTablePrivilegesItem = `SELECT`

const CreateTablePrivilegesItemUnknownPrivilege CreateTablePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateTablePrivilegesItemUsage CreateTablePrivilegesItem = `USAGE`

const CreateTablePrivilegesItemWriteFiles CreateTablePrivilegesItem = `WRITE_FILES`

type CreateTableResponse struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name ('columns') is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Columns []ColumnInfo `json:"columns,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Table was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the data_access_configuration to use.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
	DataSourceFormat CreateTableResponseDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// <catalog_name>.<schema_name>.<table_name>
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []CreateTableResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create:REQ Update:IGN] Name of parent Schema relative to its parent
	// Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// [Create,Update:OPT] List of schemes whose objects can be referenced
	// without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// [Create:OPT Update:IGN] Name of the storage credential this table used
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, EXTERNAL
	// tables)
	StorageLocation string `json:"storage_location,omitempty"`
	// [Create:IGN Update:IGN] Name of Table, relative to parent Schema.
	TableId string `json:"table_id,omitempty"`
	// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
	TableType CreateTableResponseTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == "VIEW")
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
type CreateTableResponseDataSourceFormat string

const CreateTableResponseDataSourceFormatAvro CreateTableResponseDataSourceFormat = `AVRO`

const CreateTableResponseDataSourceFormatCsv CreateTableResponseDataSourceFormat = `CSV`

const CreateTableResponseDataSourceFormatDelta CreateTableResponseDataSourceFormat = `DELTA`

const CreateTableResponseDataSourceFormatDeltasharing CreateTableResponseDataSourceFormat = `DELTASHARING`

const CreateTableResponseDataSourceFormatJson CreateTableResponseDataSourceFormat = `JSON`

const CreateTableResponseDataSourceFormatOrc CreateTableResponseDataSourceFormat = `ORC`

const CreateTableResponseDataSourceFormatParquet CreateTableResponseDataSourceFormat = `PARQUET`

const CreateTableResponseDataSourceFormatText CreateTableResponseDataSourceFormat = `TEXT`

const CreateTableResponseDataSourceFormatUnityCatalog CreateTableResponseDataSourceFormat = `UNITY_CATALOG`

const CreateTableResponseDataSourceFormatUnknownDataSourceFormat CreateTableResponseDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

type CreateTableResponsePrivilegesItem string

const CreateTableResponsePrivilegesItemCreate CreateTableResponsePrivilegesItem = `CREATE`

const CreateTableResponsePrivilegesItemCreateMount CreateTableResponsePrivilegesItem = `CREATE_MOUNT`

const CreateTableResponsePrivilegesItemCreateTable CreateTableResponsePrivilegesItem = `CREATE_TABLE`

const CreateTableResponsePrivilegesItemModify CreateTableResponsePrivilegesItem = `MODIFY`

const CreateTableResponsePrivilegesItemReadFiles CreateTableResponsePrivilegesItem = `READ_FILES`

const CreateTableResponsePrivilegesItemSelect CreateTableResponsePrivilegesItem = `SELECT`

const CreateTableResponsePrivilegesItemUnknownPrivilege CreateTableResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateTableResponsePrivilegesItemUsage CreateTableResponsePrivilegesItem = `USAGE`

const CreateTableResponsePrivilegesItemWriteFiles CreateTableResponsePrivilegesItem = `WRITE_FILES`

// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
type CreateTableResponseTableType string

const CreateTableResponseTableTypeExternal CreateTableResponseTableType = `EXTERNAL`

const CreateTableResponseTableTypeManaged CreateTableResponseTableType = `MANAGED`

const CreateTableResponseTableTypeUnknownTableType CreateTableResponseTableType = `UNKNOWN_TABLE_TYPE`

const CreateTableResponseTableTypeView CreateTableResponseTableType = `VIEW`

// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
type CreateTableTableType string

const CreateTableTableTypeExternal CreateTableTableType = `EXTERNAL`

const CreateTableTableTypeManaged CreateTableTableType = `MANAGED`

const CreateTableTableTypeUnknownTableType CreateTableTableType = `UNKNOWN_TABLE_TYPE`

const CreateTableTableTypeView CreateTableTableType = `VIEW`

// Delete a catalog
type DeleteCatalogRequest struct {
	// Required. The name of the catalog.
	Name string `json:"-" url:"-"`
}

// Delete an external location
type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force bool `json:"-" url:"force,omitempty"`
	// Required. Name of the storage credential.
	Name string `json:"-" url:"-"`
}

// Delete a Metastore
type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Required. Unique ID of the Metastore (from URL).
	Id string `json:"-" url:"-"`
}

// Delete a provider
type DeleteProviderRequest struct {
	// Required. Name of the provider.
	Name string `json:"-" url:"-"`
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Required. Name of the recipient.
	Name string `json:"-" url:"-"`
}

// Delete a schema
type DeleteSchemaRequest struct {
	// Required. Full name of the schema (from URL).
	FullName string `json:"-" url:"-"`
}

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name string `json:"-" url:"-"`
}

// Delete a credential
type DeleteStorageCredentialRequest struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force bool `json:"-" url:"force,omitempty"`
	// Required. Name of the storage credential.
	Name string `json:"-" url:"-"`
}

// Delete a table
type DeleteTableRequest struct {
	// Required. Full name of the Table (from URL).
	FullName string `json:"-" url:"-"`
}

type ExternalLocationInfo struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location's Storage Credential.
	CredentialId string `json:"credential_id,omitempty"`
	// [Create:REQ Update:OPT] Current name of the Storage Credential this
	// location uses.
	CredentialName string `json:"credential_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore hosting the External
	// Location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of the External Location.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the External Location.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Time at which this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the External
	// Location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create:REQ Update:OPT] Path URL of the External Location.
	Url string `json:"url,omitempty"`
}

type GcpServiceAccountKey struct {
	// The email of the service account. [Create:REQ].
	Email string `json:"email,omitempty"`
	// The service account's RSA private key. [Create:REQ]
	PrivateKey string `json:"private_key,omitempty"`
	// The ID of the service account's private key. [Create:REQ]
	PrivateKeyId string `json:"private_key_id,omitempty"`
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// Required. The one time activation url. It also accepts activation token.
	ActivationUrl string `json:"-" url:"-"`
}

// Get a catalog
type GetCatalogRequest struct {
	// Required. The name of the catalog.
	Name string `json:"-" url:"-"`
}

type GetCatalogResponse struct {
	// [Create,Update:IGN] The type of the catalog.
	CatalogType GetCatalogResponseCatalogType `json:"catalog_type,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Catalog was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Catalog creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN,Update:OPT] Username of current owner of Catalog.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Catalog.
	Privileges []GetCatalogResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// Delta Sharing Catalog specific fields. A Delta Sharing Catalog is a
	// catalog that is based on a Delta share on a remote sharing server.
	// [Create:OPT,Update:IGN] The name of delta sharing provider.
	ProviderName string `json:"provider_name,omitempty"`
	// [Create:OPT,Update: IGN] The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// [Create,Update:IGN] Time at which this Catalog was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Catalog.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create,Update:IGN] The type of the catalog.
type GetCatalogResponseCatalogType string

const GetCatalogResponseCatalogTypeDeltasharingCatalog GetCatalogResponseCatalogType = `DELTASHARING_CATALOG`

const GetCatalogResponseCatalogTypeManagedCatalog GetCatalogResponseCatalogType = `MANAGED_CATALOG`

const GetCatalogResponseCatalogTypeSystemCatalog GetCatalogResponseCatalogType = `SYSTEM_CATALOG`

const GetCatalogResponseCatalogTypeUnknownCatalogType GetCatalogResponseCatalogType = `UNKNOWN_CATALOG_TYPE`

type GetCatalogResponsePrivilegesItem string

const GetCatalogResponsePrivilegesItemCreate GetCatalogResponsePrivilegesItem = `CREATE`

const GetCatalogResponsePrivilegesItemCreateMount GetCatalogResponsePrivilegesItem = `CREATE_MOUNT`

const GetCatalogResponsePrivilegesItemCreateTable GetCatalogResponsePrivilegesItem = `CREATE_TABLE`

const GetCatalogResponsePrivilegesItemModify GetCatalogResponsePrivilegesItem = `MODIFY`

const GetCatalogResponsePrivilegesItemReadFiles GetCatalogResponsePrivilegesItem = `READ_FILES`

const GetCatalogResponsePrivilegesItemSelect GetCatalogResponsePrivilegesItem = `SELECT`

const GetCatalogResponsePrivilegesItemUnknownPrivilege GetCatalogResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const GetCatalogResponsePrivilegesItemUsage GetCatalogResponsePrivilegesItem = `USAGE`

const GetCatalogResponsePrivilegesItemWriteFiles GetCatalogResponsePrivilegesItem = `WRITE_FILES`

// Get an external location
type GetExternalLocationRequest struct {
	// Required. Name of the storage credential.
	Name string `json:"-" url:"-"`
}

type GetExternalLocationResponse struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location's Storage Credential.
	CredentialId string `json:"credential_id,omitempty"`
	// [Create:REQ Update:OPT] Current name of the Storage Credential this
	// location uses.
	CredentialName string `json:"credential_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore hosting the External
	// Location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of the External Location.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the External Location.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Time at which this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the External
	// Location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create:REQ Update:OPT] Path URL of the External Location.
	Url string `json:"url,omitempty"`
}

// Get permissions
type GetGrantRequest struct {
	// Required. Unique identifier (full name) of Securable (from URL).
	FullName string `json:"-" url:"-"`
	// Optional. List permissions granted to this principal.
	Principal string `json:"-" url:"principal,omitempty"`
	// Required. Type of Securable (from URL).
	SecurableType string `json:"-" url:"-"`
}

// Get a Metastore
type GetMetastoreRequest struct {
	// Required. Unique ID of the Metastore (from URL).
	Id string `json:"-" url:"-"`
}

type GetMetastoreResponse struct {
	// [Create,Update:IGN] Time at which this Metastore was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:IGN Update:OPT] Unique identifier of (Default) Data Access
	// Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// [Create:IGN Update:OPT] Whether Delta Sharing is enabled on this
	// metastore.
	DeltaSharingEnabled bool `json:"delta_sharing_enabled,omitempty"`
	// [Create:IGN Update:OPT] The lifetime of delta sharing recipient token in
	// seconds
	DeltaSharingRecipientTokenLifetimeInSeconds int `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Metastore.
	Privileges []GetMetastoreResponsePrivilegesItem `json:"privileges,omitempty"`
	// The region this metastore has an afinity to. This is used by
	// accounts-manager. Ignored by Unity Catalog.
	Region string `json:"region,omitempty"`
	// [Create:REQ Update:ERR] Storage root URL for Metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// [Create:IGN Update:OPT] UUID of storage credential to access storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// [Create,Update:IGN] Time at which the Metastore was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Metastore.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type GetMetastoreResponsePrivilegesItem string

const GetMetastoreResponsePrivilegesItemCreate GetMetastoreResponsePrivilegesItem = `CREATE`

const GetMetastoreResponsePrivilegesItemCreateMount GetMetastoreResponsePrivilegesItem = `CREATE_MOUNT`

const GetMetastoreResponsePrivilegesItemCreateTable GetMetastoreResponsePrivilegesItem = `CREATE_TABLE`

const GetMetastoreResponsePrivilegesItemModify GetMetastoreResponsePrivilegesItem = `MODIFY`

const GetMetastoreResponsePrivilegesItemReadFiles GetMetastoreResponsePrivilegesItem = `READ_FILES`

const GetMetastoreResponsePrivilegesItemSelect GetMetastoreResponsePrivilegesItem = `SELECT`

const GetMetastoreResponsePrivilegesItemUnknownPrivilege GetMetastoreResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const GetMetastoreResponsePrivilegesItemUsage GetMetastoreResponsePrivilegesItem = `USAGE`

const GetMetastoreResponsePrivilegesItemWriteFiles GetMetastoreResponsePrivilegesItem = `WRITE_FILES`

type GetMetastoreSummaryResponse struct {
	// Unique identifier of the Metastore's (Default) Data Access Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The unique ID (UUID) of the Metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the Metastore
	Name string `json:"name,omitempty"`
	// UUID of storage credential to access the metastore storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
}

type GetPermissionsResponse struct {
	// Note to self (acain): Unfortunately, neither json_inline nor json_map
	// work here.
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

// Get a provider
type GetProviderRequest struct {
	// Required. Name of the provider.
	Name string `json:"-" url:"-"`
}

type GetProviderResponse struct {
	// [Create,Update:IGN] Whether this provider is successfully activated by
	// the data provider. This field is only present when the authentication
	// type is DATABRICKS.
	ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType GetProviderResponseAuthenticationType `json:"authentication_type,omitempty"`
	// [Create,Update:OPT] Description about the provider.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:REQ] The name of the Provider.
	Name string `json:"name,omitempty"`
	// [Create,Update:IGN] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// [Create,Update:OPT] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// [Create,Update:IGN] The server-generated one-time sharing code. This
	// field is only present when the authentication type is DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Share.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type GetProviderResponseAuthenticationType string

const GetProviderResponseAuthenticationTypeDatabricks GetProviderResponseAuthenticationType = `DATABRICKS`

const GetProviderResponseAuthenticationTypeToken GetProviderResponseAuthenticationType = `TOKEN`

const GetProviderResponseAuthenticationTypeUnknown GetProviderResponseAuthenticationType = `UNKNOWN`

// Get a share recipient
type GetRecipientRequest struct {
	// Required. Name of the recipient.
	Name string `json:"-" url:"-"`
}

type GetRecipientResponse struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient's activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType GetRecipientResponseAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which the recipient was updated, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type GetRecipientResponseAuthenticationType string

const GetRecipientResponseAuthenticationTypeDatabricks GetRecipientResponseAuthenticationType = `DATABRICKS`

const GetRecipientResponseAuthenticationTypeToken GetRecipientResponseAuthenticationType = `TOKEN`

const GetRecipientResponseAuthenticationTypeUnknown GetRecipientResponseAuthenticationType = `UNKNOWN`

type GetRecipientSharePermissionsResponse struct {
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`
}

// Get a schema
type GetSchemaRequest struct {
	// Required. Full name of the schema (from URL).
	FullName string `json:"-" url:"-"`
}

type GetSchemaResponse struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Schema creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Full name of Schema, in form of
	// <catalog_name>.<schema_name>.
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []GetSchemaResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type GetSchemaResponsePrivilegesItem string

const GetSchemaResponsePrivilegesItemCreate GetSchemaResponsePrivilegesItem = `CREATE`

const GetSchemaResponsePrivilegesItemCreateMount GetSchemaResponsePrivilegesItem = `CREATE_MOUNT`

const GetSchemaResponsePrivilegesItemCreateTable GetSchemaResponsePrivilegesItem = `CREATE_TABLE`

const GetSchemaResponsePrivilegesItemModify GetSchemaResponsePrivilegesItem = `MODIFY`

const GetSchemaResponsePrivilegesItemReadFiles GetSchemaResponsePrivilegesItem = `READ_FILES`

const GetSchemaResponsePrivilegesItemSelect GetSchemaResponsePrivilegesItem = `SELECT`

const GetSchemaResponsePrivilegesItemUnknownPrivilege GetSchemaResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const GetSchemaResponsePrivilegesItemUsage GetSchemaResponsePrivilegesItem = `USAGE`

const GetSchemaResponsePrivilegesItemWriteFiles GetSchemaResponsePrivilegesItem = `WRITE_FILES`

type GetSharePermissionsResponse struct {
	// Note to self (acain): Unfortunately, neither json_inline nor json_map
	// work here.
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData bool `json:"-" url:"include_shared_data,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`
}

type GetShareResponse struct {
	// [Create: OPT] comment when creating the share
	Comment string `json:"comment,omitempty"`
	// [Create:IGN] Time at which this Share was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN] Username of Share creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:REQ] Name of the Share.
	Name string `json:"name,omitempty"`
	// [Create: IGN] A list of shared data objects within the Share.
	Objects []SharedDataObject `json:"objects,omitempty"`
}

// Get a credential
type GetStorageCredentialRequest struct {
	// Required. Name of the storage credential.
	Name string `json:"-" url:"-"`
}

type GetStorageCredentialResponse struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// [Create:IGN] The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ, Update:OPT] The credential name. The name MUST be unique
	// within the Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Time at which this credential was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Get a table
type GetTableRequest struct {
	// Required. Full name of the Table (from URL).
	FullName string `json:"-" url:"-"`
}

type GetTableResponse struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name ('columns') is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Columns []ColumnInfo `json:"columns,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Table was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the data_access_configuration to use.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
	DataSourceFormat GetTableResponseDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// <catalog_name>.<schema_name>.<table_name>
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []GetTableResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create:REQ Update:IGN] Name of parent Schema relative to its parent
	// Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// [Create,Update:OPT] List of schemes whose objects can be referenced
	// without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// [Create:OPT Update:IGN] Name of the storage credential this table used
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, EXTERNAL
	// tables)
	StorageLocation string `json:"storage_location,omitempty"`
	// [Create:IGN Update:IGN] Name of Table, relative to parent Schema.
	TableId string `json:"table_id,omitempty"`
	// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
	TableType GetTableResponseTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == "VIEW")
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
type GetTableResponseDataSourceFormat string

const GetTableResponseDataSourceFormatAvro GetTableResponseDataSourceFormat = `AVRO`

const GetTableResponseDataSourceFormatCsv GetTableResponseDataSourceFormat = `CSV`

const GetTableResponseDataSourceFormatDelta GetTableResponseDataSourceFormat = `DELTA`

const GetTableResponseDataSourceFormatDeltasharing GetTableResponseDataSourceFormat = `DELTASHARING`

const GetTableResponseDataSourceFormatJson GetTableResponseDataSourceFormat = `JSON`

const GetTableResponseDataSourceFormatOrc GetTableResponseDataSourceFormat = `ORC`

const GetTableResponseDataSourceFormatParquet GetTableResponseDataSourceFormat = `PARQUET`

const GetTableResponseDataSourceFormatText GetTableResponseDataSourceFormat = `TEXT`

const GetTableResponseDataSourceFormatUnityCatalog GetTableResponseDataSourceFormat = `UNITY_CATALOG`

const GetTableResponseDataSourceFormatUnknownDataSourceFormat GetTableResponseDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

type GetTableResponsePrivilegesItem string

const GetTableResponsePrivilegesItemCreate GetTableResponsePrivilegesItem = `CREATE`

const GetTableResponsePrivilegesItemCreateMount GetTableResponsePrivilegesItem = `CREATE_MOUNT`

const GetTableResponsePrivilegesItemCreateTable GetTableResponsePrivilegesItem = `CREATE_TABLE`

const GetTableResponsePrivilegesItemModify GetTableResponsePrivilegesItem = `MODIFY`

const GetTableResponsePrivilegesItemReadFiles GetTableResponsePrivilegesItem = `READ_FILES`

const GetTableResponsePrivilegesItemSelect GetTableResponsePrivilegesItem = `SELECT`

const GetTableResponsePrivilegesItemUnknownPrivilege GetTableResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const GetTableResponsePrivilegesItemUsage GetTableResponsePrivilegesItem = `USAGE`

const GetTableResponsePrivilegesItemWriteFiles GetTableResponsePrivilegesItem = `WRITE_FILES`

// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
type GetTableResponseTableType string

const GetTableResponseTableTypeExternal GetTableResponseTableType = `EXTERNAL`

const GetTableResponseTableTypeManaged GetTableResponseTableType = `MANAGED`

const GetTableResponseTableTypeUnknownTableType GetTableResponseTableType = `UNKNOWN_TABLE_TYPE`

const GetTableResponseTableTypeView GetTableResponseTableType = `VIEW`

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}

type ListCatalogsResponse struct {
	// AN array of catalog information objects.
	Catalogs []CatalogInfo `json:"catalogs,omitempty"`
}

type ListExternalLocationsResponse struct {
	// AN array of external locations.
	ExternalLocations []ExternalLocationInfo `json:"external_locations,omitempty"`
}

type ListMetastoresResponse struct {
	// An array of Metastore information objects.
	Metastores []MetastoreInfo `json:"metastores,omitempty"`
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

// List schemas
type ListSchemasRequest struct {
	// Optional. Parent catalog for schemas of interest.
	CatalogName string `json:"-" url:"catalog_name,omitempty"`
}

type ListSchemasResponse struct {
	// An array of schema information objects.
	Schemas []SchemaInfo `json:"schemas,omitempty"`
}

// List shares
type ListSharesRequest struct {
	// Required. Name of the provider in which to list shares.
	Name string `json:"-" url:"-"`
}

type ListSharesResponse struct {
	// An array of data share information objects.
	Shares []ShareInfo `json:"shares,omitempty"`
}

type ListStorageCredentialsResponse struct {
	// TODO: add pagination to UC list APIs.
	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`
}

type ListTableSummariesResponse struct {
	// Optional. Opaque token for pagination. Empty if there's no more page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Only name, catalog_name, schema_name, full_name and table_type will be
	// set.
	Tables []TableSummary `json:"tables,omitempty"`
}

// List tables
type ListTablesRequest struct {
	// Required. Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name,omitempty"`
	// Required (for now -- may be optional for wildcard search in future).
	// Parent schema of tables.
	SchemaName string `json:"-" url:"schema_name,omitempty"`
}

type ListTablesResponse struct {
	// An array of table information objects.
	Tables []TableInfo `json:"tables,omitempty"`
}

type MetastoreInfo struct {
	// [Create,Update:IGN] Time at which this Metastore was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:IGN Update:OPT] Unique identifier of (Default) Data Access
	// Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// [Create:IGN Update:OPT] Whether Delta Sharing is enabled on this
	// metastore.
	DeltaSharingEnabled bool `json:"delta_sharing_enabled,omitempty"`
	// [Create:IGN Update:OPT] The lifetime of delta sharing recipient token in
	// seconds
	DeltaSharingRecipientTokenLifetimeInSeconds int `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Metastore.
	Privileges []MetastoreInfoPrivilegesItem `json:"privileges,omitempty"`
	// The region this metastore has an afinity to. This is used by
	// accounts-manager. Ignored by Unity Catalog.
	Region string `json:"region,omitempty"`
	// [Create:REQ Update:ERR] Storage root URL for Metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// [Create:IGN Update:OPT] UUID of storage credential to access storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// [Create,Update:IGN] Time at which the Metastore was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Metastore.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type MetastoreInfoPrivilegesItem string

const MetastoreInfoPrivilegesItemCreate MetastoreInfoPrivilegesItem = `CREATE`

const MetastoreInfoPrivilegesItemCreateMount MetastoreInfoPrivilegesItem = `CREATE_MOUNT`

const MetastoreInfoPrivilegesItemCreateTable MetastoreInfoPrivilegesItem = `CREATE_TABLE`

const MetastoreInfoPrivilegesItemModify MetastoreInfoPrivilegesItem = `MODIFY`

const MetastoreInfoPrivilegesItemReadFiles MetastoreInfoPrivilegesItem = `READ_FILES`

const MetastoreInfoPrivilegesItemSelect MetastoreInfoPrivilegesItem = `SELECT`

const MetastoreInfoPrivilegesItemUnknownPrivilege MetastoreInfoPrivilegesItem = `UNKNOWN_PRIVILEGE`

const MetastoreInfoPrivilegesItemUsage MetastoreInfoPrivilegesItem = `USAGE`

const MetastoreInfoPrivilegesItemWriteFiles MetastoreInfoPrivilegesItem = `WRITE_FILES`

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
	// "databricks-account-id". When this field is set, field `value` can not be
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

type PermissionsChange struct {
	// The set of privileges to add.
	Add []PermissionsChangeAddItem `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	Remove []PermissionsChangeRemoveItem `json:"remove,omitempty"`
}

type PermissionsChangeAddItem string

const PermissionsChangeAddItemCreate PermissionsChangeAddItem = `CREATE`

const PermissionsChangeAddItemCreateMount PermissionsChangeAddItem = `CREATE_MOUNT`

const PermissionsChangeAddItemCreateTable PermissionsChangeAddItem = `CREATE_TABLE`

const PermissionsChangeAddItemModify PermissionsChangeAddItem = `MODIFY`

const PermissionsChangeAddItemReadFiles PermissionsChangeAddItem = `READ_FILES`

const PermissionsChangeAddItemSelect PermissionsChangeAddItem = `SELECT`

const PermissionsChangeAddItemUnknownPrivilege PermissionsChangeAddItem = `UNKNOWN_PRIVILEGE`

const PermissionsChangeAddItemUsage PermissionsChangeAddItem = `USAGE`

const PermissionsChangeAddItemWriteFiles PermissionsChangeAddItem = `WRITE_FILES`

type PermissionsChangeRemoveItem string

const PermissionsChangeRemoveItemCreate PermissionsChangeRemoveItem = `CREATE`

const PermissionsChangeRemoveItemCreateMount PermissionsChangeRemoveItem = `CREATE_MOUNT`

const PermissionsChangeRemoveItemCreateTable PermissionsChangeRemoveItem = `CREATE_TABLE`

const PermissionsChangeRemoveItemModify PermissionsChangeRemoveItem = `MODIFY`

const PermissionsChangeRemoveItemReadFiles PermissionsChangeRemoveItem = `READ_FILES`

const PermissionsChangeRemoveItemSelect PermissionsChangeRemoveItem = `SELECT`

const PermissionsChangeRemoveItemUnknownPrivilege PermissionsChangeRemoveItem = `UNKNOWN_PRIVILEGE`

const PermissionsChangeRemoveItemUsage PermissionsChangeRemoveItem = `USAGE`

const PermissionsChangeRemoveItemWriteFiles PermissionsChangeRemoveItem = `WRITE_FILES`

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	Privileges []PrivilegeAssignmentPrivilegesItem `json:"privileges,omitempty"`
}

type PrivilegeAssignmentPrivilegesItem string

const PrivilegeAssignmentPrivilegesItemCreate PrivilegeAssignmentPrivilegesItem = `CREATE`

const PrivilegeAssignmentPrivilegesItemCreateMount PrivilegeAssignmentPrivilegesItem = `CREATE_MOUNT`

const PrivilegeAssignmentPrivilegesItemCreateTable PrivilegeAssignmentPrivilegesItem = `CREATE_TABLE`

const PrivilegeAssignmentPrivilegesItemModify PrivilegeAssignmentPrivilegesItem = `MODIFY`

const PrivilegeAssignmentPrivilegesItemReadFiles PrivilegeAssignmentPrivilegesItem = `READ_FILES`

const PrivilegeAssignmentPrivilegesItemSelect PrivilegeAssignmentPrivilegesItem = `SELECT`

const PrivilegeAssignmentPrivilegesItemUnknownPrivilege PrivilegeAssignmentPrivilegesItem = `UNKNOWN_PRIVILEGE`

const PrivilegeAssignmentPrivilegesItemUsage PrivilegeAssignmentPrivilegesItem = `USAGE`

const PrivilegeAssignmentPrivilegesItemWriteFiles PrivilegeAssignmentPrivilegesItem = `WRITE_FILES`

type ProviderInfo struct {
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType ProviderInfoAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:IGN,Update:IGN] Cloud vendor of the provider's UC Metastore. This
	// field is only present when the authentication type is DATABRICKS.
	Cloud string `json:"cloud,omitempty"`
	// [Create,Update:OPT] Description about the provider.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:IGN,Update:IGN] The global UC metastore id of the data provider
	// This field is only present when the authentication type is DATABRICKS.
	// The identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId string `json:"data_provider_global_metastore_id,omitempty"`
	// [Create:IGN,Update:IGN] UUID of the provider's UC Metastore. This field
	// is only present when the authentication type is DATABRICKS.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create,Update:REQ] The name of the Provider.
	Name string `json:"name,omitempty"`
	// [Create,Update:OPT] Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// [Create,Update:OPT] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// [Create:IGN,Update:IGN] Cloud region of the provider's UC Metastore. This
	// field is only present when the authentication type is DATABRICKS.
	Region string `json:"region,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Share.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type ProviderInfoAuthenticationType string

const ProviderInfoAuthenticationTypeDatabricks ProviderInfoAuthenticationType = `DATABRICKS`

const ProviderInfoAuthenticationTypeToken ProviderInfoAuthenticationType = `TOKEN`

const ProviderInfoAuthenticationTypeUnknown ProviderInfoAuthenticationType = `UNKNOWN`

type ProviderShare struct {
	// The name of the Provider Share.
	Name string `json:"name,omitempty"`
}

type RecipientInfo struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient's activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType RecipientInfoAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:IGN,Update:IGN] Cloud vendor of the recipient's UC Metastore.
	// This field is only present when the authentication type is DATABRICKS.
	Cloud string `json:"cloud,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:IGN] The global UC metastore id provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS. The identifier is of format
	// <cloud>:<region>:<metastore-uuid>.
	DataRecipientGlobalMetastoreId string `json:"data_recipient_global_metastore_id,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:IGN,Update:IGN] UUID of the recipient's UC Metastore. This field
	// is only present when the authentication type is DATABRICKS.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:IGN,Update:OPT] Username of Recipient owner.
	Owner string `json:"owner,omitempty"`
	// [Create:IGN,Update:IGN] Cloud region of the recipient's UC Metastore.
	// This field is only present when the authentication type is DATABRICKS.
	Region string `json:"region,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] Recipient Tokens. This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this recipient was updated, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type RecipientInfoAuthenticationType string

const RecipientInfoAuthenticationTypeDatabricks RecipientInfoAuthenticationType = `DATABRICKS`

const RecipientInfoAuthenticationTypeToken RecipientInfoAuthenticationType = `TOKEN`

const RecipientInfoAuthenticationTypeUnknown RecipientInfoAuthenticationType = `UNKNOWN`

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
	// Required. The one time activation url. It also accepts activation token.
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
	// Required. This will set the expiration_time of existing token only to a
	// smaller timestamp, it cannot extend the expiration_time. Use 0 to expire
	// the existing token immediately, negative number will return an error.
	ExistingTokenExpireInSeconds int64 `json:"existing_token_expire_in_seconds,omitempty"`
	// Required. The name of the recipient.
	Name string `json:"-" url:"-"`
}

type RotateRecipientTokenResponse struct {
	// [Create:IGN, Update:IGN] A boolean status field showing whether the
	// Recipient's activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN, Update:IGN] Full activation url to retrieve the access
	// token. It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ, Update:IGN] The delta sharing authentication type.
	AuthenticationType RotateRecipientTokenResponseAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN, Update:IGN] Time at which this recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN, Update:IGN] Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ, Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] recipient Tokens. This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which the recipient was updated, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ, Update:IGN] The delta sharing authentication type.
type RotateRecipientTokenResponseAuthenticationType string

const RotateRecipientTokenResponseAuthenticationTypeDatabricks RotateRecipientTokenResponseAuthenticationType = `DATABRICKS`

const RotateRecipientTokenResponseAuthenticationTypeToken RotateRecipientTokenResponseAuthenticationType = `TOKEN`

const RotateRecipientTokenResponseAuthenticationTypeUnknown RotateRecipientTokenResponseAuthenticationType = `UNKNOWN`

type SchemaInfo struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Schema creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Full name of Schema, in form of
	// <catalog_name>.<schema_name>.
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []SchemaInfoPrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type SchemaInfoPrivilegesItem string

const SchemaInfoPrivilegesItemCreate SchemaInfoPrivilegesItem = `CREATE`

const SchemaInfoPrivilegesItemCreateMount SchemaInfoPrivilegesItem = `CREATE_MOUNT`

const SchemaInfoPrivilegesItemCreateTable SchemaInfoPrivilegesItem = `CREATE_TABLE`

const SchemaInfoPrivilegesItemModify SchemaInfoPrivilegesItem = `MODIFY`

const SchemaInfoPrivilegesItemReadFiles SchemaInfoPrivilegesItem = `READ_FILES`

const SchemaInfoPrivilegesItemSelect SchemaInfoPrivilegesItem = `SELECT`

const SchemaInfoPrivilegesItemUnknownPrivilege SchemaInfoPrivilegesItem = `UNKNOWN_PRIVILEGE`

const SchemaInfoPrivilegesItemUsage SchemaInfoPrivilegesItem = `USAGE`

const SchemaInfoPrivilegesItemWriteFiles SchemaInfoPrivilegesItem = `WRITE_FILES`

type ShareInfo struct {
	// [Create: OPT] comment when creating the share
	Comment string `json:"comment,omitempty"`
	// [Create:IGN] Time at which this Share was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN] Username of Share creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:REQ] Name of the Share.
	Name string `json:"name,omitempty"`
	// [Create: IGN] A list of shared data objects within the Share.
	Objects []SharedDataObject `json:"objects,omitempty"`
	// [Create:IGN,Update:OPT] Username of Share owner.
	Owner string `json:"owner,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this Share was updated, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Share updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Get share permissions
type SharePermissionsRequest struct {
	// Required. The name of the Recipient.
	Name string `json:"-" url:"-"`
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
	// The share name.
	ShareName string `json:"share_name,omitempty"`
}

type SharedDataObject struct {
	// The time when this data object is added to the Share, in epoch
	// milliseconds. Output only field. [Update:IGN]
	AddedAt int64 `json:"added_at,omitempty"`
	// Username of the sharer. Output only field. [Update:IGN]
	AddedBy string `json:"added_by,omitempty"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	// [Update: OPT]
	CdfEnabled bool `json:"cdf_enabled,omitempty"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment string `json:"comment,omitempty"`
	// The type of the data object. Output only field. [Update:IGN]
	DataObjectType string `json:"data_object_type,omitempty"`
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`. [Update:REQ]
	Name string `json:"name,omitempty"`
	// Array of partitions for the shared data.
	Partitions []Partition `json:"partitions,omitempty"`
	// A user-provided new name for the data object within the share. If this
	// new name is not not provided, the object's original name will be used as
	// the `shared_as` name. The `shared_as` name must be unique within a Share.
	//
	// For tables, the new name must follow the format of `<schema>.<table>`.
	// [Update:OPT]
	SharedAs string `json:"shared_as,omitempty"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the \"current\" version of the
	// object. [Update: OPT]
	StartVersion int64 `json:"start_version,omitempty"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status SharedDataObjectStatus `json:"status,omitempty"`
}

// One of: **ACTIVE**, **PERMISSION_DENIED**.
type SharedDataObjectStatus string

const SharedDataObjectStatusActive SharedDataObjectStatus = `ACTIVE`

const SharedDataObjectStatusPermissionDenied SharedDataObjectStatus = `PERMISSION_DENIED`

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

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// [Create:IGN] The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ, Update:OPT] The credential name. The name MUST be unique
	// within the Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Time at which this credential was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type StringKeyValuePair struct {
	// The key for the tuple.
	Key string `json:"key"`
	// The value for the tuple.
	Value string `json:"value"`
}

type TableInfo struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name ('columns') is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Columns []ColumnInfo `json:"columns,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Table was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the data_access_configuration to use.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
	DataSourceFormat TableInfoDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// <catalog_name>.<schema_name>.<table_name>
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []TableInfoPrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create:REQ Update:IGN] Name of parent Schema relative to its parent
	// Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// [Create,Update:OPT] List of schemes whose objects can be referenced
	// without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// [Create:OPT Update:IGN] Name of the storage credential this table used
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, EXTERNAL
	// tables)
	StorageLocation string `json:"storage_location,omitempty"`
	// [Create:IGN Update:IGN] Name of Table, relative to parent Schema.
	TableId string `json:"table_id,omitempty"`
	// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
	TableType TableInfoTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == "VIEW")
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
type TableInfoDataSourceFormat string

const TableInfoDataSourceFormatAvro TableInfoDataSourceFormat = `AVRO`

const TableInfoDataSourceFormatCsv TableInfoDataSourceFormat = `CSV`

const TableInfoDataSourceFormatDelta TableInfoDataSourceFormat = `DELTA`

const TableInfoDataSourceFormatDeltasharing TableInfoDataSourceFormat = `DELTASHARING`

const TableInfoDataSourceFormatJson TableInfoDataSourceFormat = `JSON`

const TableInfoDataSourceFormatOrc TableInfoDataSourceFormat = `ORC`

const TableInfoDataSourceFormatParquet TableInfoDataSourceFormat = `PARQUET`

const TableInfoDataSourceFormatText TableInfoDataSourceFormat = `TEXT`

const TableInfoDataSourceFormatUnityCatalog TableInfoDataSourceFormat = `UNITY_CATALOG`

const TableInfoDataSourceFormatUnknownDataSourceFormat TableInfoDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

type TableInfoPrivilegesItem string

const TableInfoPrivilegesItemCreate TableInfoPrivilegesItem = `CREATE`

const TableInfoPrivilegesItemCreateMount TableInfoPrivilegesItem = `CREATE_MOUNT`

const TableInfoPrivilegesItemCreateTable TableInfoPrivilegesItem = `CREATE_TABLE`

const TableInfoPrivilegesItemModify TableInfoPrivilegesItem = `MODIFY`

const TableInfoPrivilegesItemReadFiles TableInfoPrivilegesItem = `READ_FILES`

const TableInfoPrivilegesItemSelect TableInfoPrivilegesItem = `SELECT`

const TableInfoPrivilegesItemUnknownPrivilege TableInfoPrivilegesItem = `UNKNOWN_PRIVILEGE`

const TableInfoPrivilegesItemUsage TableInfoPrivilegesItem = `USAGE`

const TableInfoPrivilegesItemWriteFiles TableInfoPrivilegesItem = `WRITE_FILES`

// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
type TableInfoTableType string

const TableInfoTableTypeExternal TableInfoTableType = `EXTERNAL`

const TableInfoTableTypeManaged TableInfoTableType = `MANAGED`

const TableInfoTableTypeUnknownTableType TableInfoTableType = `UNKNOWN_TABLE_TYPE`

const TableInfoTableTypeView TableInfoTableType = `VIEW`

// List table summaries
type TableSummariesRequest struct {
	// Required. Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name,omitempty"`
	// Optional. Maximum number of tables to return (page length). Defaults to
	// 10000.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Optional. Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Optional. A sql LIKE pattern (% and _) for schema names. All schemas will
	// be returned if not set or empty.
	SchemaNamePattern string `json:"-" url:"schema_name_pattern,omitempty"`
	// Optional. A sql LIKE pattern (% and _) for table names. All tables will
	// be returned if not set or empty.
	TableNamePattern string `json:"-" url:"table_name_pattern,omitempty"`
}

type TableSummary struct {
	// The full name of the table.
	FullName string `json:"full_name,omitempty"`
	// One of: **UNKNOWN_TABLE_TYPE**, **MANAGED**, **EXTERNAL**, **VIEW**.
	TableType TableSummaryTableType `json:"table_type,omitempty"`
}

// One of: **UNKNOWN_TABLE_TYPE**, **MANAGED**, **EXTERNAL**, **VIEW**.
type TableSummaryTableType string

const TableSummaryTableTypeExternal TableSummaryTableType = `EXTERNAL`

const TableSummaryTableTypeManaged TableSummaryTableType = `MANAGED`

const TableSummaryTableTypeUnknownTableType TableSummaryTableType = `UNKNOWN_TABLE_TYPE`

const TableSummaryTableTypeView TableSummaryTableType = `VIEW`

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the Metastore to delete.
	MetastoreId string `json:"-" url:"metastore_id"`
	// A workspace ID.
	WorkspaceId int `json:"-" url:"-"`
}

type UpdateCatalog struct {
	// [Create,Update:IGN] The type of the catalog.
	CatalogType UpdateCatalogCatalogType `json:"catalog_type,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Catalog was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Catalog creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Catalog.
	Name string `json:"name,omitempty" url:"-"`
	// [Create:IGN,Update:OPT] Username of current owner of Catalog.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Catalog.
	Privileges []UpdateCatalogPrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// Delta Sharing Catalog specific fields. A Delta Sharing Catalog is a
	// catalog that is based on a Delta share on a remote sharing server.
	// [Create:OPT,Update:IGN] The name of delta sharing provider.
	ProviderName string `json:"provider_name,omitempty"`
	// [Create:OPT,Update: IGN] The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// [Create,Update:IGN] Time at which this Catalog was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Catalog.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create,Update:IGN] The type of the catalog.
type UpdateCatalogCatalogType string

const UpdateCatalogCatalogTypeDeltasharingCatalog UpdateCatalogCatalogType = `DELTASHARING_CATALOG`

const UpdateCatalogCatalogTypeManagedCatalog UpdateCatalogCatalogType = `MANAGED_CATALOG`

const UpdateCatalogCatalogTypeSystemCatalog UpdateCatalogCatalogType = `SYSTEM_CATALOG`

const UpdateCatalogCatalogTypeUnknownCatalogType UpdateCatalogCatalogType = `UNKNOWN_CATALOG_TYPE`

type UpdateCatalogPrivilegesItem string

const UpdateCatalogPrivilegesItemCreate UpdateCatalogPrivilegesItem = `CREATE`

const UpdateCatalogPrivilegesItemCreateMount UpdateCatalogPrivilegesItem = `CREATE_MOUNT`

const UpdateCatalogPrivilegesItemCreateTable UpdateCatalogPrivilegesItem = `CREATE_TABLE`

const UpdateCatalogPrivilegesItemModify UpdateCatalogPrivilegesItem = `MODIFY`

const UpdateCatalogPrivilegesItemReadFiles UpdateCatalogPrivilegesItem = `READ_FILES`

const UpdateCatalogPrivilegesItemSelect UpdateCatalogPrivilegesItem = `SELECT`

const UpdateCatalogPrivilegesItemUnknownPrivilege UpdateCatalogPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateCatalogPrivilegesItemUsage UpdateCatalogPrivilegesItem = `USAGE`

const UpdateCatalogPrivilegesItemWriteFiles UpdateCatalogPrivilegesItem = `WRITE_FILES`

type UpdateExternalLocation struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location's Storage Credential.
	CredentialId string `json:"credential_id,omitempty"`
	// [Create:REQ Update:OPT] Current name of the Storage Credential this
	// location uses.
	CredentialName string `json:"credential_name,omitempty"`
	// TODO: SC-90063 re-add 'force' parameter in backward-compatible way for
	// DBR (not removed below as it still works with CLI) Optional. Force update
	// even if changing url invalidates dependent external tables or mounts.
	Force bool `json:"force,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore hosting the External
	// Location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of the External Location.
	Name string `json:"name,omitempty" url:"-"`
	// [Create:IGN Update:OPT] The owner of the External Location.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Time at which this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the External
	// Location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create:REQ Update:OPT] Path URL of the External Location.
	Url string `json:"url,omitempty"`
}

type UpdateMetastore struct {
	// [Create,Update:IGN] Time at which this Metastore was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:IGN Update:OPT] Unique identifier of (Default) Data Access
	// Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// [Create:IGN Update:OPT] Whether Delta Sharing is enabled on this
	// metastore.
	DeltaSharingEnabled bool `json:"delta_sharing_enabled,omitempty"`
	// [Create:IGN Update:OPT] The lifetime of delta sharing recipient token in
	// seconds
	DeltaSharingRecipientTokenLifetimeInSeconds int `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// Required. Unique ID of the Metastore (from URL).
	Id string `json:"-" url:"-"`
	// [Create,Update:IGN] Unique identifier of Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Metastore.
	Privileges []UpdateMetastorePrivilegesItem `json:"privileges,omitempty"`
	// The region this metastore has an afinity to. This is used by
	// accounts-manager. Ignored by Unity Catalog.
	Region string `json:"region,omitempty"`
	// [Create:REQ Update:ERR] Storage root URL for Metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// [Create:IGN Update:OPT] UUID of storage credential to access storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// [Create,Update:IGN] Time at which the Metastore was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Metastore.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog for the Metastore.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// A workspace ID.
	WorkspaceId int `json:"-" url:"-"`
}

type UpdateMetastorePrivilegesItem string

const UpdateMetastorePrivilegesItemCreate UpdateMetastorePrivilegesItem = `CREATE`

const UpdateMetastorePrivilegesItemCreateMount UpdateMetastorePrivilegesItem = `CREATE_MOUNT`

const UpdateMetastorePrivilegesItemCreateTable UpdateMetastorePrivilegesItem = `CREATE_TABLE`

const UpdateMetastorePrivilegesItemModify UpdateMetastorePrivilegesItem = `MODIFY`

const UpdateMetastorePrivilegesItemReadFiles UpdateMetastorePrivilegesItem = `READ_FILES`

const UpdateMetastorePrivilegesItemSelect UpdateMetastorePrivilegesItem = `SELECT`

const UpdateMetastorePrivilegesItemUnknownPrivilege UpdateMetastorePrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateMetastorePrivilegesItemUsage UpdateMetastorePrivilegesItem = `USAGE`

const UpdateMetastorePrivilegesItemWriteFiles UpdateMetastorePrivilegesItem = `WRITE_FILES`

type UpdatePermissions struct {
	// Array of permissions change objects.
	Changes []PermissionsChange `json:"changes,omitempty"`
	// Required. Unique identifier (full name) of Securable (from URL).
	FullName string `json:"-" url:"-"`
	// Optional. List permissions granted to this principal.
	Principal string `json:"-" url:"principal,omitempty"`
	// Required. Type of Securable (from URL).
	SecurableType string `json:"-" url:"-"`
}

type UpdateProvider struct {
	// [Create,Update:IGN] Whether this provider is successfully activated by
	// the data provider. This field is only present when the authentication
	// type is DATABRICKS.
	ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType UpdateProviderAuthenticationType `json:"authentication_type,omitempty"`
	// [Create,Update:OPT] Description about the provider.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create, Update:REQ] The name of the Provider.
	Name string `json:"name,omitempty" url:"-"`
	// [Create,Update:IGN] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// [Create,Update:OPT] This field is only present when the authentication
	// type is TOKEN.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// [Create,Update:IGN] The server-generated one-time sharing code. This
	// field is only present when the authentication type is DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Share.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type UpdateProviderAuthenticationType string

const UpdateProviderAuthenticationTypeDatabricks UpdateProviderAuthenticationType = `DATABRICKS`

const UpdateProviderAuthenticationTypeToken UpdateProviderAuthenticationType = `TOKEN`

const UpdateProviderAuthenticationTypeUnknown UpdateProviderAuthenticationType = `UNKNOWN`

type UpdateRecipient struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient's activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType UpdateRecipientAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty" url:"-"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which the recipient was updated, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type UpdateRecipientAuthenticationType string

const UpdateRecipientAuthenticationTypeDatabricks UpdateRecipientAuthenticationType = `DATABRICKS`

const UpdateRecipientAuthenticationTypeToken UpdateRecipientAuthenticationType = `TOKEN`

const UpdateRecipientAuthenticationTypeUnknown UpdateRecipientAuthenticationType = `UNKNOWN`

type UpdateSchema struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Schema creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Full name of Schema, in form of
	// <catalog_name>.<schema_name>.
	FullName string `json:"full_name,omitempty" url:"-"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []UpdateSchemaPrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type UpdateSchemaPrivilegesItem string

const UpdateSchemaPrivilegesItemCreate UpdateSchemaPrivilegesItem = `CREATE`

const UpdateSchemaPrivilegesItemCreateMount UpdateSchemaPrivilegesItem = `CREATE_MOUNT`

const UpdateSchemaPrivilegesItemCreateTable UpdateSchemaPrivilegesItem = `CREATE_TABLE`

const UpdateSchemaPrivilegesItemModify UpdateSchemaPrivilegesItem = `MODIFY`

const UpdateSchemaPrivilegesItemReadFiles UpdateSchemaPrivilegesItem = `READ_FILES`

const UpdateSchemaPrivilegesItemSelect UpdateSchemaPrivilegesItem = `SELECT`

const UpdateSchemaPrivilegesItemUnknownPrivilege UpdateSchemaPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateSchemaPrivilegesItemUsage UpdateSchemaPrivilegesItem = `USAGE`

const UpdateSchemaPrivilegesItemWriteFiles UpdateSchemaPrivilegesItem = `WRITE_FILES`

type UpdateShare struct {
	// The name of the share.
	Name string `json:"-" url:"-"`
	// Array of shared data object updates.
	Updates []SharedDataObjectUpdate `json:"updates,omitempty"`
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	Changes []PermissionsChange `json:"changes,omitempty"`
	// Required. The name of the share.
	Name string `json:"-" url:"-"`
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// [Create:IGN] The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ, Update:OPT] The credential name. The name MUST be unique
	// within the Metastore.
	Name string `json:"name,omitempty" url:"-"`
	// [Create:IGN Update:OPT] Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Optional. Supplying true to this argument skips validation of the updated
	// set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// [Create,Update:IGN] Time at which this credential was last modified, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type UpdateTable struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name ('columns') is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Columns []ColumnInfo `json:"columns,omitempty"`
	// [Create,Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Table was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the data_access_configuration to use.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
	DataSourceFormat UpdateTableDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// <catalog_name>.<schema_name>.<table_name>
	FullName string `json:"full_name,omitempty" url:"-"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []UpdateTablePrivilegesItem `json:"privileges,omitempty"`
	// This name ('properties') is what the client sees as the field name in
	// messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create:REQ Update:IGN] Name of parent Schema relative to its parent
	// Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// [Create,Update:OPT] List of schemes whose objects can be referenced
	// without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// [Create:OPT Update:IGN] Name of the storage credential this table used
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, EXTERNAL
	// tables)
	StorageLocation string `json:"storage_location,omitempty"`
	// [Create:IGN Update:IGN] Name of Table, relative to parent Schema.
	TableId string `json:"table_id,omitempty"`
	// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
	TableType UpdateTableTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == "VIEW")
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format ("DELTA", "CSV", etc.).
type UpdateTableDataSourceFormat string

const UpdateTableDataSourceFormatAvro UpdateTableDataSourceFormat = `AVRO`

const UpdateTableDataSourceFormatCsv UpdateTableDataSourceFormat = `CSV`

const UpdateTableDataSourceFormatDelta UpdateTableDataSourceFormat = `DELTA`

const UpdateTableDataSourceFormatDeltasharing UpdateTableDataSourceFormat = `DELTASHARING`

const UpdateTableDataSourceFormatJson UpdateTableDataSourceFormat = `JSON`

const UpdateTableDataSourceFormatOrc UpdateTableDataSourceFormat = `ORC`

const UpdateTableDataSourceFormatParquet UpdateTableDataSourceFormat = `PARQUET`

const UpdateTableDataSourceFormatText UpdateTableDataSourceFormat = `TEXT`

const UpdateTableDataSourceFormatUnityCatalog UpdateTableDataSourceFormat = `UNITY_CATALOG`

const UpdateTableDataSourceFormatUnknownDataSourceFormat UpdateTableDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

type UpdateTablePrivilegesItem string

const UpdateTablePrivilegesItemCreate UpdateTablePrivilegesItem = `CREATE`

const UpdateTablePrivilegesItemCreateMount UpdateTablePrivilegesItem = `CREATE_MOUNT`

const UpdateTablePrivilegesItemCreateTable UpdateTablePrivilegesItem = `CREATE_TABLE`

const UpdateTablePrivilegesItemModify UpdateTablePrivilegesItem = `MODIFY`

const UpdateTablePrivilegesItemReadFiles UpdateTablePrivilegesItem = `READ_FILES`

const UpdateTablePrivilegesItemSelect UpdateTablePrivilegesItem = `SELECT`

const UpdateTablePrivilegesItemUnknownPrivilege UpdateTablePrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateTablePrivilegesItemUsage UpdateTablePrivilegesItem = `USAGE`

const UpdateTablePrivilegesItemWriteFiles UpdateTablePrivilegesItem = `WRITE_FILES`

// [Create:REQ Update:OPT] Table type ("MANAGED", "EXTERNAL", "VIEW").
type UpdateTableTableType string

const UpdateTableTableTypeExternal UpdateTableTableType = `EXTERNAL`

const UpdateTableTableTypeManaged UpdateTableTableType = `MANAGED`

const UpdateTableTableTypeUnknownTableType UpdateTableTableType = `UNKNOWN_TABLE_TYPE`

const UpdateTableTableTypeView UpdateTableTableType = `VIEW`
