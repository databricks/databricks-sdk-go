// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

// all definitions in this file are in alphabetical order

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
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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

type CreateCatalogRequest struct {
	// [Create,Update:IGN] The type of the catalog.
	CatalogType CreateCatalogRequestCatalogType `json:"catalog_type,omitempty"`
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
	Privileges []CreateCatalogRequestPrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
type CreateCatalogRequestCatalogType string

const CreateCatalogRequestCatalogTypeDeltasharingCatalog CreateCatalogRequestCatalogType = `DELTASHARING_CATALOG`

const CreateCatalogRequestCatalogTypeManagedCatalog CreateCatalogRequestCatalogType = `MANAGED_CATALOG`

const CreateCatalogRequestCatalogTypeSystemCatalog CreateCatalogRequestCatalogType = `SYSTEM_CATALOG`

const CreateCatalogRequestCatalogTypeUnknownCatalogType CreateCatalogRequestCatalogType = `UNKNOWN_CATALOG_TYPE`

type CreateCatalogRequestPrivilegesItem string

const CreateCatalogRequestPrivilegesItemCreate CreateCatalogRequestPrivilegesItem = `CREATE`

const CreateCatalogRequestPrivilegesItemCreateMount CreateCatalogRequestPrivilegesItem = `CREATE_MOUNT`

const CreateCatalogRequestPrivilegesItemCreateTable CreateCatalogRequestPrivilegesItem = `CREATE_TABLE`

const CreateCatalogRequestPrivilegesItemModify CreateCatalogRequestPrivilegesItem = `MODIFY`

const CreateCatalogRequestPrivilegesItemReadFiles CreateCatalogRequestPrivilegesItem = `READ_FILES`

const CreateCatalogRequestPrivilegesItemSelect CreateCatalogRequestPrivilegesItem = `SELECT`

const CreateCatalogRequestPrivilegesItemUnknownPrivilege CreateCatalogRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateCatalogRequestPrivilegesItemUsage CreateCatalogRequestPrivilegesItem = `USAGE`

const CreateCatalogRequestPrivilegesItemWriteFiles CreateCatalogRequestPrivilegesItem = `WRITE_FILES`

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
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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

type CreateExternalLocationRequest struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location&#39;s Storage Credential.
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
	// [Create,Update:IGN] Unique ID of the location&#39;s Storage Credential.
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

type CreateMetastoreAssignmentRequest struct {
	DefaultCatalogName string `json:"default_catalog_name"`

	MetastoreId string `json:"metastore_id"`

	WorkspaceId int ` path:"workspace_id"`
}

type CreateMetastoreRequest struct {
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
	Privileges []CreateMetastoreRequestPrivilegesItem `json:"privileges,omitempty"`
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

type CreateMetastoreRequestPrivilegesItem string

const CreateMetastoreRequestPrivilegesItemCreate CreateMetastoreRequestPrivilegesItem = `CREATE`

const CreateMetastoreRequestPrivilegesItemCreateMount CreateMetastoreRequestPrivilegesItem = `CREATE_MOUNT`

const CreateMetastoreRequestPrivilegesItemCreateTable CreateMetastoreRequestPrivilegesItem = `CREATE_TABLE`

const CreateMetastoreRequestPrivilegesItemModify CreateMetastoreRequestPrivilegesItem = `MODIFY`

const CreateMetastoreRequestPrivilegesItemReadFiles CreateMetastoreRequestPrivilegesItem = `READ_FILES`

const CreateMetastoreRequestPrivilegesItemSelect CreateMetastoreRequestPrivilegesItem = `SELECT`

const CreateMetastoreRequestPrivilegesItemUnknownPrivilege CreateMetastoreRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateMetastoreRequestPrivilegesItemUsage CreateMetastoreRequestPrivilegesItem = `USAGE`

const CreateMetastoreRequestPrivilegesItemWriteFiles CreateMetastoreRequestPrivilegesItem = `WRITE_FILES`

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

type CreateProviderRequest struct {
	// [Create,Update:IGN] Whether this provider is successfully activated by
	// the data provider. This field is only present when the authentication
	// type is DATABRICKS.
	ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType CreateProviderRequestAuthenticationType `json:"authentication_type,omitempty"`
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
type CreateProviderRequestAuthenticationType string

const CreateProviderRequestAuthenticationTypeDatabricks CreateProviderRequestAuthenticationType = `DATABRICKS`

const CreateProviderRequestAuthenticationTypeToken CreateProviderRequestAuthenticationType = `TOKEN`

const CreateProviderRequestAuthenticationTypeUnknown CreateProviderRequestAuthenticationType = `UNKNOWN`

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

type CreateRecipientRequest struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient&#39;s activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType CreateRecipientRequestAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this Recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] Recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type CreateRecipientRequestAuthenticationType string

const CreateRecipientRequestAuthenticationTypeDatabricks CreateRecipientRequestAuthenticationType = `DATABRICKS`

const CreateRecipientRequestAuthenticationTypeToken CreateRecipientRequestAuthenticationType = `TOKEN`

const CreateRecipientRequestAuthenticationTypeUnknown CreateRecipientRequestAuthenticationType = `UNKNOWN`

type CreateRecipientResponse struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient&#39;s activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType CreateRecipientResponseAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this Recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] Recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type CreateRecipientResponseAuthenticationType string

const CreateRecipientResponseAuthenticationTypeDatabricks CreateRecipientResponseAuthenticationType = `DATABRICKS`

const CreateRecipientResponseAuthenticationTypeToken CreateRecipientResponseAuthenticationType = `TOKEN`

const CreateRecipientResponseAuthenticationTypeUnknown CreateRecipientResponseAuthenticationType = `UNKNOWN`

type CreateSchemaRequest struct {
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
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []CreateSchemaRequestPrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CreateSchemaRequestPrivilegesItem string

const CreateSchemaRequestPrivilegesItemCreate CreateSchemaRequestPrivilegesItem = `CREATE`

const CreateSchemaRequestPrivilegesItemCreateMount CreateSchemaRequestPrivilegesItem = `CREATE_MOUNT`

const CreateSchemaRequestPrivilegesItemCreateTable CreateSchemaRequestPrivilegesItem = `CREATE_TABLE`

const CreateSchemaRequestPrivilegesItemModify CreateSchemaRequestPrivilegesItem = `MODIFY`

const CreateSchemaRequestPrivilegesItemReadFiles CreateSchemaRequestPrivilegesItem = `READ_FILES`

const CreateSchemaRequestPrivilegesItemSelect CreateSchemaRequestPrivilegesItem = `SELECT`

const CreateSchemaRequestPrivilegesItemUnknownPrivilege CreateSchemaRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateSchemaRequestPrivilegesItemUsage CreateSchemaRequestPrivilegesItem = `USAGE`

const CreateSchemaRequestPrivilegesItemWriteFiles CreateSchemaRequestPrivilegesItem = `WRITE_FILES`

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
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []CreateSchemaResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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

type CreateShareRequest struct {
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

type CreateStagingTableRequest struct {
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

type CreateStorageCredentialRequest struct {
	AwsIamRole any/* ERROR */ `json:"aws_iam_role,omitempty"`

	AzureServicePrincipal any/* ERROR */ `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`

	GcpServiceAccountKey any/* ERROR */ `json:"gcp_service_account_key,omitempty"`
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
	AwsIamRole any/* ERROR */ `json:"aws_iam_role,omitempty"`

	AzureServicePrincipal any/* ERROR */ `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`

	GcpServiceAccountKey any/* ERROR */ `json:"gcp_service_account_key,omitempty"`
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

type CreateTableRequest struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name (&#39;columns&#39;) is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
	DataSourceFormat CreateTableRequestDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.&lt;table_name&gt;
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []CreateTableRequestPrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
	TableType CreateTableRequestTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == &#34;VIEW&#34;)
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
type CreateTableRequestDataSourceFormat string

const CreateTableRequestDataSourceFormatAvro CreateTableRequestDataSourceFormat = `AVRO`

const CreateTableRequestDataSourceFormatCsv CreateTableRequestDataSourceFormat = `CSV`

const CreateTableRequestDataSourceFormatDelta CreateTableRequestDataSourceFormat = `DELTA`

const CreateTableRequestDataSourceFormatDeltasharing CreateTableRequestDataSourceFormat = `DELTASHARING`

const CreateTableRequestDataSourceFormatJson CreateTableRequestDataSourceFormat = `JSON`

const CreateTableRequestDataSourceFormatOrc CreateTableRequestDataSourceFormat = `ORC`

const CreateTableRequestDataSourceFormatParquet CreateTableRequestDataSourceFormat = `PARQUET`

const CreateTableRequestDataSourceFormatText CreateTableRequestDataSourceFormat = `TEXT`

const CreateTableRequestDataSourceFormatUnityCatalog CreateTableRequestDataSourceFormat = `UNITY_CATALOG`

const CreateTableRequestDataSourceFormatUnknownDataSourceFormat CreateTableRequestDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

type CreateTableRequestPrivilegesItem string

const CreateTableRequestPrivilegesItemCreate CreateTableRequestPrivilegesItem = `CREATE`

const CreateTableRequestPrivilegesItemCreateMount CreateTableRequestPrivilegesItem = `CREATE_MOUNT`

const CreateTableRequestPrivilegesItemCreateTable CreateTableRequestPrivilegesItem = `CREATE_TABLE`

const CreateTableRequestPrivilegesItemModify CreateTableRequestPrivilegesItem = `MODIFY`

const CreateTableRequestPrivilegesItemReadFiles CreateTableRequestPrivilegesItem = `READ_FILES`

const CreateTableRequestPrivilegesItemSelect CreateTableRequestPrivilegesItem = `SELECT`

const CreateTableRequestPrivilegesItemUnknownPrivilege CreateTableRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateTableRequestPrivilegesItemUsage CreateTableRequestPrivilegesItem = `USAGE`

const CreateTableRequestPrivilegesItemWriteFiles CreateTableRequestPrivilegesItem = `WRITE_FILES`

// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type CreateTableRequestTableType string

const CreateTableRequestTableTypeExternal CreateTableRequestTableType = `EXTERNAL`

const CreateTableRequestTableTypeManaged CreateTableRequestTableType = `MANAGED`

const CreateTableRequestTableTypeUnknownTableType CreateTableRequestTableType = `UNKNOWN_TABLE_TYPE`

const CreateTableRequestTableTypeView CreateTableRequestTableType = `VIEW`

type CreateTableResponse struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name (&#39;columns&#39;) is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
	DataSourceFormat CreateTableResponseDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.&lt;table_name&gt;
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []CreateTableResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
	TableType CreateTableResponseTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == &#34;VIEW&#34;)
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
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

// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type CreateTableResponseTableType string

const CreateTableResponseTableTypeExternal CreateTableResponseTableType = `EXTERNAL`

const CreateTableResponseTableTypeManaged CreateTableResponseTableType = `MANAGED`

const CreateTableResponseTableTypeUnknownTableType CreateTableResponseTableType = `UNKNOWN_TABLE_TYPE`

const CreateTableResponseTableTypeView CreateTableResponseTableType = `VIEW`

type DeleteCatalogRequest struct {
	// Required. Name of the Catalog (from URL).
	NameArg string ` path:"name_arg"`
}

type DeleteExternalLocationRequest struct {
	// Optional. Force deletion even if there are dependent external tables or
	// mounts.
	Force bool `json:"force,omitempty"`
	// Required. Name of the storage credential.
	NameArg string ` path:"name_arg"`
}

type DeleteMetastoreAssignmentRequest struct {
	MetastoreId string `json:"metastore_id"`

	WorkspaceId int ` path:"workspace_id"`
}

type DeleteMetastoreRequest struct {
	// Optional. Force deletion even if the metastore is not empty. Default is
	// false.
	Force bool `json:"force,omitempty"`
	// Required. Unique ID of the Metastore (from URL).
	Id string ` path:"id"`
}

type DeleteProviderRequest struct {
	// Required. Name of the provider.
	NameArg string ` path:"name_arg"`
}

type DeleteRecipientRequest struct {
	// Required. Name of the recipient.
	Name string ` path:"name"`
}

type DeleteSchemaRequest struct {
	// Required. Full name of the Schema (from URL).
	FullNameArg string ` path:"full_name_arg"`
}

type DeleteShareRequest struct {
	Name string ` path:"name"`
}

type DeleteStorageCredentialRequest struct {
	// Optional. Force deletion even if there are dependent external locations
	// or external tables.
	Force bool `json:"force,omitempty"`
	// Required. Name of the storage credential.
	NameArg string ` path:"name_arg"`
}

type DeleteTableRequest struct {
	// Required. Full name of the Table (from URL).
	FullNameArg string ` path:"full_name_arg"`
}

type ExternalLocationInfo struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location&#39;s Storage Credential.
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

type FileInfo struct {
	// Whether the object represents a directory or a file.
	IsDir bool `json:"is_dir,omitempty"`
	// Modification time, unix epoch.
	Mtime int64 `json:"mtime,omitempty"`
	// Name of the object.
	Name string `json:"name,omitempty"`
	// Path URI of the storage object.
	Path string `json:"path,omitempty"`
	// Size in bytes.
	Size int64 `json:"size,omitempty"`
}

type GetActivationUrlInfoRequest struct {
	// Required. The one time activation url. It also accepts activation token.
	ActivationUrl string ` path:"activation_url"`
}

type GetCatalogRequest struct {
	// Required. Name of the Catalog (from URL).
	NameArg string ` path:"name_arg"`
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
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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

type GetExternalLocationRequest struct {
	// Required. Name of the storage credential.
	NameArg string ` path:"name_arg"`
}

type GetExternalLocationResponse struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location&#39;s Storage Credential.
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

type GetMetastoreRequest struct {
	// Required. Unique ID of the Metastore (from URL).
	Id string ` path:"id"`
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
	// Unique identifier of the Metastore&#39;s (Default) Data Access Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The unique ID (UUID) of the Metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the Metastore
	Name string `json:"name,omitempty"`
	// UUID of storage credential to access the metastore storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
}

type GetPermissionsRequest struct {
	// Optional. List permissions granted to this principal.
	Principal string ` url:"principal,omitempty"`
	// Required. Unique identifier (full name) of Securable (from URL).
	SecurableFullName string ` path:"securable_full_name"`
	// Required. Type of Securable (from URL).
	SecurableType string ` path:"securable_type"`
}

type GetPermissionsResponse struct {
	// Note to self (acain): Unfortunately, neither json_inline nor json_map
	// work here.
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

type GetProviderRequest struct {
	// Required. Name of the provider.
	NameArg string ` path:"name_arg"`
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

type GetRecipientRequest struct {
	// Required. Name of the recipient.
	Name string ` path:"name"`
}

type GetRecipientResponse struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient&#39;s activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType GetRecipientResponseAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this Recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] Recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type GetRecipientResponseAuthenticationType string

const GetRecipientResponseAuthenticationTypeDatabricks GetRecipientResponseAuthenticationType = `DATABRICKS`

const GetRecipientResponseAuthenticationTypeToken GetRecipientResponseAuthenticationType = `TOKEN`

const GetRecipientResponseAuthenticationTypeUnknown GetRecipientResponseAuthenticationType = `UNKNOWN`

type GetRecipientSharePermissionsRequest struct {
	// Required. The name of the Recipient.
	Name string ` path:"name"`
}

type GetRecipientSharePermissionsResponse struct {
	PermissionsOut []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`
}

type GetSchemaRequest struct {
	// Required. Full name of the Schema (from URL).
	FullNameArg string ` path:"full_name_arg"`
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
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []GetSchemaResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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

type GetSharePermissionsRequest struct {
	// Required. The name of the Share.
	Name string ` path:"name"`
}

type GetSharePermissionsResponse struct {
	// Note to self (acain): Unfortunately, neither json_inline nor json_map
	// work here.
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

type GetShareRequest struct {
	IncludeSharedData bool ` url:"include_shared_data,omitempty"`

	Name string ` path:"name"`
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

type GetStorageCredentialRequest struct {
	// Required. Name of the storage credential.
	NameArg string ` path:"name_arg"`
}

type GetStorageCredentialResponse struct {
	AwsIamRole any/* ERROR */ `json:"aws_iam_role,omitempty"`

	AzureServicePrincipal any/* ERROR */ `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`

	GcpServiceAccountKey any/* ERROR */ `json:"gcp_service_account_key,omitempty"`
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

type GetTableRequest struct {
	// Required. Full name of the Table (from URL).
	FullNameArg string ` path:"full_name_arg"`
}

type GetTableResponse struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name (&#39;columns&#39;) is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
	DataSourceFormat GetTableResponseDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.&lt;table_name&gt;
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []GetTableResponsePrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
	TableType GetTableResponseTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == &#34;VIEW&#34;)
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
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

// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
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
	Catalogs []CatalogInfo `json:"catalogs,omitempty"`
}

type ListExternalLocationsResponse struct {
	ExternalLocations []ExternalLocationInfo `json:"external_locations,omitempty"`
}

type ListFilesRequest struct {
	// Optional. Name of a Storage Credential to use for accessing the URL.
	CredentialName string ` url:"credential_name,omitempty"`
	// Optional. Limit on number of results to return.
	MaxResults int ` url:"max_results,omitempty"`
	// Required. Path URL to list files from.
	Url string ` url:"url,omitempty"`
}

type ListFilesResponse struct {
	Files []FileInfo `json:"files,omitempty"`
}

type ListMetastoresResponse struct {
	Metastores []MetastoreInfo `json:"metastores,omitempty"`
}

type ListProviderSharesRequest struct {
	// Required. Name of the provider in which to list shares.
	ProviderNameArg string ` path:"provider_name_arg"`
}

type ListProviderSharesResponse struct {
	Shares []ProviderShare `json:"shares,omitempty"`
}

type ListProvidersResponse struct {
	Providers []ProviderInfo `json:"providers,omitempty"`
}

type ListRecipientsResponse struct {
	Recipients []RecipientInfo `json:"recipients,omitempty"`
}

type ListSchemasRequest struct {
	// Optional. Parent catalog for schemas of interest.
	CatalogName string ` url:"catalog_name,omitempty"`
}

type ListSchemasResponse struct {
	Schemas []SchemaInfo `json:"schemas,omitempty"`
}

type ListSharesResponse struct {
	Shares []ShareInfo `json:"shares,omitempty"`
}

type ListStorageCredentialsResponse struct {
	// TODO: add pagination to UC list APIs.
	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`
}

type ListTableSummariesRequest struct {
	// Required. Name of parent catalog for tables of interest.
	CatalogName string ` url:"catalog_name,omitempty"`
	// Optional. Maximum number of tables to return (page length). Defaults to
	// 10000.
	MaxResults int ` url:"max_results,omitempty"`
	// Optional. Opaque token to send for the next page of results (pagination).
	PageToken string ` url:"page_token,omitempty"`
	// Optional. A sql LIKE pattern (% and _) for schema names. All schemas will
	// be returned if not set or empty.
	SchemaNamePattern string ` url:"schema_name_pattern,omitempty"`
	// Optional. A sql LIKE pattern (% and _) for table names. All tables will
	// be returned if not set or empty.
	TableNamePattern string ` url:"table_name_pattern,omitempty"`
}

type ListTableSummariesResponse struct {
	// Optional. Opaque token for pagination. Empty if there&#39;s no more page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Only name, catalog_name, schema_name, full_name and table_type will be
	// set.
	Tables []TableSummary `json:"tables,omitempty"`
}

type ListTablesRequest struct {
	// Required. Name of parent catalog for tables of interest.
	CatalogName string ` url:"catalog_name,omitempty"`
	// Required (for now -- may be optional for wildcard search in future).
	// Parent schema of tables.
	SchemaName string ` url:"schema_name,omitempty"`
}

type ListTablesResponse struct {
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
	Values []PartitionValue `json:"values,omitempty"`
}

type PartitionValue struct {
	// The name of the partition column.
	Name string `json:"name,omitempty"`
	// The operator to apply for the value.
	Op PartitionValueOp `json:"op,omitempty"`
	// The value of the partition column. When this value is not set, it means
	// `null` value.
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
	// [Create,Update:IGN] Whether this provider is successfully activated by
	// the data provider. This field is only present when the authentication
	// type is DATABRICKS.
	ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType ProviderInfoAuthenticationType `json:"authentication_type,omitempty"`
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
	// Recipient&#39;s activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType RecipientInfoAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this Recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] Recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type RecipientInfoAuthenticationType string

const RecipientInfoAuthenticationTypeDatabricks RecipientInfoAuthenticationType = `DATABRICKS`

const RecipientInfoAuthenticationTypeToken RecipientInfoAuthenticationType = `TOKEN`

const RecipientInfoAuthenticationTypeUnknown RecipientInfoAuthenticationType = `UNKNOWN`

type RecipientProfile struct {
	BearerToken string `json:"bearer_token,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	ShareCredentialsVersion int `json:"share_credentials_version,omitempty"`
}

type RecipientTokenInfo struct {
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// Time at which this Recipient Token was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Recipient Token creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Unique id of the Recipient Token.
	Id string `json:"id,omitempty"`
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of Recipient Token updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type RetrieveTokenRequest struct {
	// Required. The one time activation url. It also accepts activation token.
	ActivationUrl string ` path:"activation_url"`
}

type RetrieveTokenResponse struct {
	BearerToken string `json:"bearerToken,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`
}

type RotateRecipientTokenRequest struct {
	// Required. This will set the expiration_time of existing token only to a
	// smaller timestamp, it cannot extend the expiration_time. Use 0 to expire
	// the existing token immediately, negative number will return an error.
	ExistingTokenExpireInSeconds int64 `json:"existing_token_expire_in_seconds,omitempty"`
	// Required. The name of the Recipient.
	Name string ` path:"name"`
}

type RotateRecipientTokenResponse struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient&#39;s activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType RotateRecipientTokenResponseAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this Recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] Recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
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
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []SchemaInfoPrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment string `json:"comment,omitempty"`
	// The type of the data object. Output only field. [Update:IGN]
	DataObjectType string `json:"data_object_type,omitempty"`
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table&#39;s fully qualified name is in the format of
	// `&lt;catalog&gt;.&lt;schema&gt;.&lt;table&gt;`. [Update:REQ]
	Name string `json:"name,omitempty"`

	Partitions []Partition `json:"partitions,omitempty"`
	// A user-provided new name for the data object within the share. If this
	// new name is not not provided, the object&#39;s original name will be used as
	// the `shared_as` name. The `shared_as` name must be unique within a Share.
	// For tables, the new name must follow the format of `&lt;schema&gt;.&lt;table&gt;`.
	// [Update:OPT]
	SharedAs string `json:"shared_as,omitempty"`
}

type SharedDataObjectUpdate struct {
	Action SharedDataObjectUpdateAction `json:"action,omitempty"`
	// The data object that is being updated (added / removed).
	DataObject *SharedDataObject `json:"data_object,omitempty"`
}

type SharedDataObjectUpdateAction string

const SharedDataObjectUpdateActionAdd SharedDataObjectUpdateAction = `ADD`

const SharedDataObjectUpdateActionRemove SharedDataObjectUpdateAction = `REMOVE`

type StorageCredentialInfo struct {
	AwsIamRole any/* ERROR */ `json:"aws_iam_role,omitempty"`

	AzureServicePrincipal any/* ERROR */ `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`

	GcpServiceAccountKey any/* ERROR */ `json:"gcp_service_account_key,omitempty"`
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
	Key string `json:"key"`

	Value string `json:"value"`
}

type TableInfo struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name (&#39;columns&#39;) is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
	DataSourceFormat TableInfoDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.&lt;table_name&gt;
	FullName string `json:"full_name,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []TableInfoPrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
	TableType TableInfoTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == &#34;VIEW&#34;)
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
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

// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type TableInfoTableType string

const TableInfoTableTypeExternal TableInfoTableType = `EXTERNAL`

const TableInfoTableTypeManaged TableInfoTableType = `MANAGED`

const TableInfoTableTypeUnknownTableType TableInfoTableType = `UNKNOWN_TABLE_TYPE`

const TableInfoTableTypeView TableInfoTableType = `VIEW`

type TableSummary struct {
	FullName string `json:"full_name,omitempty"`

	TableType TableSummaryTableType `json:"table_type,omitempty"`
}

type TableSummaryTableType string

const TableSummaryTableTypeExternal TableSummaryTableType = `EXTERNAL`

const TableSummaryTableTypeManaged TableSummaryTableType = `MANAGED`

const TableSummaryTableTypeUnknownTableType TableSummaryTableType = `UNKNOWN_TABLE_TYPE`

const TableSummaryTableTypeView TableSummaryTableType = `VIEW`

type UpdateCatalogRequest struct {
	// [Create,Update:IGN] The type of the catalog.
	CatalogType UpdateCatalogRequestCatalogType `json:"catalog_type,omitempty"`
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
	// Required. Name of the Catalog (from URL).
	NameArg string ` path:"name_arg"`
	// [Create:IGN,Update:OPT] Username of current owner of Catalog.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Catalog.
	Privileges []UpdateCatalogRequestPrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
type UpdateCatalogRequestCatalogType string

const UpdateCatalogRequestCatalogTypeDeltasharingCatalog UpdateCatalogRequestCatalogType = `DELTASHARING_CATALOG`

const UpdateCatalogRequestCatalogTypeManagedCatalog UpdateCatalogRequestCatalogType = `MANAGED_CATALOG`

const UpdateCatalogRequestCatalogTypeSystemCatalog UpdateCatalogRequestCatalogType = `SYSTEM_CATALOG`

const UpdateCatalogRequestCatalogTypeUnknownCatalogType UpdateCatalogRequestCatalogType = `UNKNOWN_CATALOG_TYPE`

type UpdateCatalogRequestPrivilegesItem string

const UpdateCatalogRequestPrivilegesItemCreate UpdateCatalogRequestPrivilegesItem = `CREATE`

const UpdateCatalogRequestPrivilegesItemCreateMount UpdateCatalogRequestPrivilegesItem = `CREATE_MOUNT`

const UpdateCatalogRequestPrivilegesItemCreateTable UpdateCatalogRequestPrivilegesItem = `CREATE_TABLE`

const UpdateCatalogRequestPrivilegesItemModify UpdateCatalogRequestPrivilegesItem = `MODIFY`

const UpdateCatalogRequestPrivilegesItemReadFiles UpdateCatalogRequestPrivilegesItem = `READ_FILES`

const UpdateCatalogRequestPrivilegesItemSelect UpdateCatalogRequestPrivilegesItem = `SELECT`

const UpdateCatalogRequestPrivilegesItemUnknownPrivilege UpdateCatalogRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateCatalogRequestPrivilegesItemUsage UpdateCatalogRequestPrivilegesItem = `USAGE`

const UpdateCatalogRequestPrivilegesItemWriteFiles UpdateCatalogRequestPrivilegesItem = `WRITE_FILES`

type UpdateExternalLocationRequest struct {
	// [Create:OPT Update:OPT] User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this External Location was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:IGN] Unique ID of the location&#39;s Storage Credential.
	CredentialId string `json:"credential_id,omitempty"`
	// [Create:REQ Update:OPT] Current name of the Storage Credential this
	// location uses.
	CredentialName string `json:"credential_name,omitempty"`
	// TODO: SC-90063 re-add &#39;force&#39; parameter in backward-compatible way for
	// DBR (not removed below as it still works with CLI) Optional. Force update
	// even if changing url invalidates dependent external tables or mounts.
	Force bool `json:"force,omitempty"`
	// [Create,Update:IGN] Unique identifier of Metastore hosting the External
	// Location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of the External Location.
	Name string `json:"name,omitempty"`
	// Required. Name of the storage credential.
	NameArg string ` path:"name_arg"`
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

type UpdateMetastoreAssignmentRequest struct {
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`

	MetastoreId string `json:"metastore_id,omitempty"`

	WorkspaceId int ` path:"workspace_id"`
}

type UpdateMetastoreRequest struct {
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
	Id string ` path:"id"`
	// [Create,Update:IGN] Unique identifier of Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Metastore.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Metastore.
	Privileges []UpdateMetastoreRequestPrivilegesItem `json:"privileges,omitempty"`
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

type UpdateMetastoreRequestPrivilegesItem string

const UpdateMetastoreRequestPrivilegesItemCreate UpdateMetastoreRequestPrivilegesItem = `CREATE`

const UpdateMetastoreRequestPrivilegesItemCreateMount UpdateMetastoreRequestPrivilegesItem = `CREATE_MOUNT`

const UpdateMetastoreRequestPrivilegesItemCreateTable UpdateMetastoreRequestPrivilegesItem = `CREATE_TABLE`

const UpdateMetastoreRequestPrivilegesItemModify UpdateMetastoreRequestPrivilegesItem = `MODIFY`

const UpdateMetastoreRequestPrivilegesItemReadFiles UpdateMetastoreRequestPrivilegesItem = `READ_FILES`

const UpdateMetastoreRequestPrivilegesItemSelect UpdateMetastoreRequestPrivilegesItem = `SELECT`

const UpdateMetastoreRequestPrivilegesItemUnknownPrivilege UpdateMetastoreRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateMetastoreRequestPrivilegesItemUsage UpdateMetastoreRequestPrivilegesItem = `USAGE`

const UpdateMetastoreRequestPrivilegesItemWriteFiles UpdateMetastoreRequestPrivilegesItem = `WRITE_FILES`

type UpdatePermissionsRequest struct {
	Changes []PermissionsChange `json:"changes,omitempty"`
	// Optional. List permissions granted to this principal.
	Principal string ` url:"principal,omitempty"`
	// Required. Unique identifier (full name) of Securable (from URL).
	SecurableFullName string ` path:"securable_full_name"`
	// Required. Type of Securable (from URL).
	SecurableType string ` path:"securable_type"`
}

type UpdateProviderRequest struct {
	// [Create,Update:IGN] Whether this provider is successfully activated by
	// the data provider. This field is only present when the authentication
	// type is DATABRICKS.
	ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType UpdateProviderRequestAuthenticationType `json:"authentication_type,omitempty"`
	// [Create,Update:OPT] Description about the provider.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Provider was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create,Update:REQ] The name of the Provider.
	Name string `json:"name,omitempty"`
	// Required. Name of the provider.
	NameArg string ` path:"name_arg"`
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
type UpdateProviderRequestAuthenticationType string

const UpdateProviderRequestAuthenticationTypeDatabricks UpdateProviderRequestAuthenticationType = `DATABRICKS`

const UpdateProviderRequestAuthenticationTypeToken UpdateProviderRequestAuthenticationType = `TOKEN`

const UpdateProviderRequestAuthenticationTypeUnknown UpdateProviderRequestAuthenticationType = `UNKNOWN`

type UpdateRecipientRequest struct {
	// [Create:IGN,Update:IGN] A boolean status field showing whether the
	// Recipient&#39;s activation URL has been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// [Create:IGN,Update:IGN] Full activation url to retrieve the access token.
	// It will be empty if the token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// [Create:REQ,Update:IGN] The delta sharing authentication type.
	AuthenticationType UpdateRecipientRequestAuthenticationType `json:"authentication_type,omitempty"`
	// [Create:OPT,Update:OPT] Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// [Create:IGN,Update:IGN] Time at which this Recipient was created, in
	// epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// [Create:OPT,Update:OPT] IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// [Create:REQ,Update:OPT] Name of Recipient.
	Name string `json:"name,omitempty"`
	// Required. Name of the recipient.
	NameArg string ` path:"name_arg"`
	// [Create:OPT,Update:IGN] The one-time sharing code provided by the data
	// recipient. This field is only present when the authentication type is
	// DATABRICKS.
	SharingCode string `json:"sharing_code,omitempty"`
	// [Create:IGN,Update:IGN] Recipient Tokens This field is only present when
	// the authentication type is TOKEN.
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in
	// epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create:IGN,Update:IGN] Username of Recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type.
type UpdateRecipientRequestAuthenticationType string

const UpdateRecipientRequestAuthenticationTypeDatabricks UpdateRecipientRequestAuthenticationType = `DATABRICKS`

const UpdateRecipientRequestAuthenticationTypeToken UpdateRecipientRequestAuthenticationType = `TOKEN`

const UpdateRecipientRequestAuthenticationTypeUnknown UpdateRecipientRequestAuthenticationType = `UNKNOWN`

type UpdateSchemaRequest struct {
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
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.
	FullName string `json:"full_name,omitempty"`
	// Required. Full name of the Schema (from URL).
	FullNameArg string ` path:"full_name_arg"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// [Create:IGN Update:OPT] Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Schema.
	Privileges []UpdateSchemaRequestPrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
	// TableInfo).
	Properties []StringKeyValuePair `json:"properties,omitempty"`
	// [Create,Update:IGN] Time at which this Schema was created, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified Schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type UpdateSchemaRequestPrivilegesItem string

const UpdateSchemaRequestPrivilegesItemCreate UpdateSchemaRequestPrivilegesItem = `CREATE`

const UpdateSchemaRequestPrivilegesItemCreateMount UpdateSchemaRequestPrivilegesItem = `CREATE_MOUNT`

const UpdateSchemaRequestPrivilegesItemCreateTable UpdateSchemaRequestPrivilegesItem = `CREATE_TABLE`

const UpdateSchemaRequestPrivilegesItemModify UpdateSchemaRequestPrivilegesItem = `MODIFY`

const UpdateSchemaRequestPrivilegesItemReadFiles UpdateSchemaRequestPrivilegesItem = `READ_FILES`

const UpdateSchemaRequestPrivilegesItemSelect UpdateSchemaRequestPrivilegesItem = `SELECT`

const UpdateSchemaRequestPrivilegesItemUnknownPrivilege UpdateSchemaRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateSchemaRequestPrivilegesItemUsage UpdateSchemaRequestPrivilegesItem = `USAGE`

const UpdateSchemaRequestPrivilegesItemWriteFiles UpdateSchemaRequestPrivilegesItem = `WRITE_FILES`

type UpdateSharePermissionsRequest struct {
	Changes []PermissionsChange `json:"changes,omitempty"`
	// Required. The name of the Share.
	Name string ` path:"name"`
}

type UpdateShareRequest struct {
	Name string ` path:"name"`

	Updates []SharedDataObjectUpdate `json:"updates,omitempty"`
}

type UpdateStorageCredentialRequest struct {
	AwsIamRole any/* ERROR */ `json:"aws_iam_role,omitempty"`

	AzureServicePrincipal any/* ERROR */ `json:"azure_service_principal,omitempty"`
	// [Create,Update:OPT] Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// [Create,Update:IGN] Time at which this Credential was created, in epoch
	// milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// [Create,Update:IGN] Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`

	GcpServiceAccountKey any/* ERROR */ `json:"gcp_service_account_key,omitempty"`
	// [Create:IGN] The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ, Update:OPT] The credential name. The name MUST be unique
	// within the Metastore.
	Name string `json:"name,omitempty"`
	// Required. Name of the storage credential.
	NameArg string ` path:"name_arg"`
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

type UpdateTableRequest struct {
	// [Create:REQ Update:IGN] Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name (&#39;columns&#39;) is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
	DataSourceFormat UpdateTableRequestDataSourceFormat `json:"data_source_format,omitempty"`
	// [Create,Update:IGN] Full name of Table, in form of
	// &lt;catalog_name&gt;.&lt;schema_name&gt;.&lt;table_name&gt;
	FullName string `json:"full_name,omitempty"`
	// Required. Full name of the Table (from URL).
	FullNameArg string ` path:"full_name_arg"`
	// [Create,Update:IGN] Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// [Create:REQ Update:OPT] Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// [Create: IGN Update:OPT] Username of current owner of Table.
	Owner string `json:"owner,omitempty"`
	// [Create,Update:IGN] Privileges the user has on the Table.
	Privileges []UpdateTableRequestPrivilegesItem `json:"privileges,omitempty"`
	// This name (&#39;properties&#39;) is what the client sees as the field name in
	// messages that include PropertiesKVPairs using &#39;json_inline&#39; (e.g.,
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
	// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
	TableType UpdateTableRequestTableType `json:"table_type,omitempty"`
	// [Create,Update:IGN] Time at which this Table was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// [Create,Update:IGN] Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// [Create,Update:OPT] View definition SQL (when table_type == &#34;VIEW&#34;)
	ViewDefinition string `json:"view_definition,omitempty"`
}

// [Create:REQ Update:OPT] Data source format (&#34;DELTA&#34;, &#34;CSV&#34;, etc.).
type UpdateTableRequestDataSourceFormat string

const UpdateTableRequestDataSourceFormatAvro UpdateTableRequestDataSourceFormat = `AVRO`

const UpdateTableRequestDataSourceFormatCsv UpdateTableRequestDataSourceFormat = `CSV`

const UpdateTableRequestDataSourceFormatDelta UpdateTableRequestDataSourceFormat = `DELTA`

const UpdateTableRequestDataSourceFormatDeltasharing UpdateTableRequestDataSourceFormat = `DELTASHARING`

const UpdateTableRequestDataSourceFormatJson UpdateTableRequestDataSourceFormat = `JSON`

const UpdateTableRequestDataSourceFormatOrc UpdateTableRequestDataSourceFormat = `ORC`

const UpdateTableRequestDataSourceFormatParquet UpdateTableRequestDataSourceFormat = `PARQUET`

const UpdateTableRequestDataSourceFormatText UpdateTableRequestDataSourceFormat = `TEXT`

const UpdateTableRequestDataSourceFormatUnityCatalog UpdateTableRequestDataSourceFormat = `UNITY_CATALOG`

const UpdateTableRequestDataSourceFormatUnknownDataSourceFormat UpdateTableRequestDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

type UpdateTableRequestPrivilegesItem string

const UpdateTableRequestPrivilegesItemCreate UpdateTableRequestPrivilegesItem = `CREATE`

const UpdateTableRequestPrivilegesItemCreateMount UpdateTableRequestPrivilegesItem = `CREATE_MOUNT`

const UpdateTableRequestPrivilegesItemCreateTable UpdateTableRequestPrivilegesItem = `CREATE_TABLE`

const UpdateTableRequestPrivilegesItemModify UpdateTableRequestPrivilegesItem = `MODIFY`

const UpdateTableRequestPrivilegesItemReadFiles UpdateTableRequestPrivilegesItem = `READ_FILES`

const UpdateTableRequestPrivilegesItemSelect UpdateTableRequestPrivilegesItem = `SELECT`

const UpdateTableRequestPrivilegesItemUnknownPrivilege UpdateTableRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateTableRequestPrivilegesItemUsage UpdateTableRequestPrivilegesItem = `USAGE`

const UpdateTableRequestPrivilegesItemWriteFiles UpdateTableRequestPrivilegesItem = `WRITE_FILES`

// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type UpdateTableRequestTableType string

const UpdateTableRequestTableTypeExternal UpdateTableRequestTableType = `EXTERNAL`

const UpdateTableRequestTableTypeManaged UpdateTableRequestTableType = `MANAGED`

const UpdateTableRequestTableTypeUnknownTableType UpdateTableRequestTableType = `UNKNOWN_TABLE_TYPE`

const UpdateTableRequestTableTypeView UpdateTableRequestTableType = `VIEW`
