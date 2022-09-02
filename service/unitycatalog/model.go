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


const CatalogInfoCatalogTypeUnknownCatalogType CatalogInfoCatalogType = `UNKNOWN_CATALOG_TYPE`

const CatalogInfoCatalogTypeManagedCatalog CatalogInfoCatalogType = `MANAGED_CATALOG`

const CatalogInfoCatalogTypeDeltasharingCatalog CatalogInfoCatalogType = `DELTASHARING_CATALOG`

const CatalogInfoCatalogTypeSystemCatalog CatalogInfoCatalogType = `SYSTEM_CATALOG`

type CatalogInfoPrivilegesItem string


const CatalogInfoPrivilegesItemUnknownPrivilege CatalogInfoPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CatalogInfoPrivilegesItemSelect CatalogInfoPrivilegesItem = `SELECT`

const CatalogInfoPrivilegesItemCreate CatalogInfoPrivilegesItem = `CREATE`

const CatalogInfoPrivilegesItemModify CatalogInfoPrivilegesItem = `MODIFY`

const CatalogInfoPrivilegesItemUsage CatalogInfoPrivilegesItem = `USAGE`

const CatalogInfoPrivilegesItemReadFiles CatalogInfoPrivilegesItem = `READ_FILES`

const CatalogInfoPrivilegesItemWriteFiles CatalogInfoPrivilegesItem = `WRITE_FILES`

const CatalogInfoPrivilegesItemCreateTable CatalogInfoPrivilegesItem = `CREATE_TABLE`

const CatalogInfoPrivilegesItemCreateMount CatalogInfoPrivilegesItem = `CREATE_MOUNT`

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


const ColumnInfoTypeNameUnknownColumnTypeName ColumnInfoTypeName = `UNKNOWN_COLUMN_TYPE_NAME`

const ColumnInfoTypeNameBoolean ColumnInfoTypeName = `BOOLEAN`

const ColumnInfoTypeNameByte ColumnInfoTypeName = `BYTE`

const ColumnInfoTypeNameShort ColumnInfoTypeName = `SHORT`

const ColumnInfoTypeNameInt ColumnInfoTypeName = `INT`

const ColumnInfoTypeNameLong ColumnInfoTypeName = `LONG`

const ColumnInfoTypeNameFloat ColumnInfoTypeName = `FLOAT`

const ColumnInfoTypeNameDouble ColumnInfoTypeName = `DOUBLE`

const ColumnInfoTypeNameDate ColumnInfoTypeName = `DATE`

const ColumnInfoTypeNameTimestamp ColumnInfoTypeName = `TIMESTAMP`

const ColumnInfoTypeNameString ColumnInfoTypeName = `STRING`

const ColumnInfoTypeNameBinary ColumnInfoTypeName = `BINARY`

const ColumnInfoTypeNameDecimal ColumnInfoTypeName = `DECIMAL`

const ColumnInfoTypeNameInterval ColumnInfoTypeName = `INTERVAL`

const ColumnInfoTypeNameArray ColumnInfoTypeName = `ARRAY`

const ColumnInfoTypeNameStruct ColumnInfoTypeName = `STRUCT`

const ColumnInfoTypeNameMap ColumnInfoTypeName = `MAP`

const ColumnInfoTypeNameChar ColumnInfoTypeName = `CHAR`

const ColumnInfoTypeNameNull ColumnInfoTypeName = `NULL`

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


const CreateCatalogRequestCatalogTypeUnknownCatalogType CreateCatalogRequestCatalogType = `UNKNOWN_CATALOG_TYPE`

const CreateCatalogRequestCatalogTypeManagedCatalog CreateCatalogRequestCatalogType = `MANAGED_CATALOG`

const CreateCatalogRequestCatalogTypeDeltasharingCatalog CreateCatalogRequestCatalogType = `DELTASHARING_CATALOG`

const CreateCatalogRequestCatalogTypeSystemCatalog CreateCatalogRequestCatalogType = `SYSTEM_CATALOG`

type CreateCatalogRequestPrivilegesItem string


const CreateCatalogRequestPrivilegesItemUnknownPrivilege CreateCatalogRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateCatalogRequestPrivilegesItemSelect CreateCatalogRequestPrivilegesItem = `SELECT`

const CreateCatalogRequestPrivilegesItemCreate CreateCatalogRequestPrivilegesItem = `CREATE`

const CreateCatalogRequestPrivilegesItemModify CreateCatalogRequestPrivilegesItem = `MODIFY`

const CreateCatalogRequestPrivilegesItemUsage CreateCatalogRequestPrivilegesItem = `USAGE`

const CreateCatalogRequestPrivilegesItemReadFiles CreateCatalogRequestPrivilegesItem = `READ_FILES`

const CreateCatalogRequestPrivilegesItemWriteFiles CreateCatalogRequestPrivilegesItem = `WRITE_FILES`

const CreateCatalogRequestPrivilegesItemCreateTable CreateCatalogRequestPrivilegesItem = `CREATE_TABLE`

const CreateCatalogRequestPrivilegesItemCreateMount CreateCatalogRequestPrivilegesItem = `CREATE_MOUNT`

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


const CreateCatalogResponseCatalogTypeUnknownCatalogType CreateCatalogResponseCatalogType = `UNKNOWN_CATALOG_TYPE`

const CreateCatalogResponseCatalogTypeManagedCatalog CreateCatalogResponseCatalogType = `MANAGED_CATALOG`

const CreateCatalogResponseCatalogTypeDeltasharingCatalog CreateCatalogResponseCatalogType = `DELTASHARING_CATALOG`

const CreateCatalogResponseCatalogTypeSystemCatalog CreateCatalogResponseCatalogType = `SYSTEM_CATALOG`

type CreateCatalogResponsePrivilegesItem string


const CreateCatalogResponsePrivilegesItemUnknownPrivilege CreateCatalogResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateCatalogResponsePrivilegesItemSelect CreateCatalogResponsePrivilegesItem = `SELECT`

const CreateCatalogResponsePrivilegesItemCreate CreateCatalogResponsePrivilegesItem = `CREATE`

const CreateCatalogResponsePrivilegesItemModify CreateCatalogResponsePrivilegesItem = `MODIFY`

const CreateCatalogResponsePrivilegesItemUsage CreateCatalogResponsePrivilegesItem = `USAGE`

const CreateCatalogResponsePrivilegesItemReadFiles CreateCatalogResponsePrivilegesItem = `READ_FILES`

const CreateCatalogResponsePrivilegesItemWriteFiles CreateCatalogResponsePrivilegesItem = `WRITE_FILES`

const CreateCatalogResponsePrivilegesItemCreateTable CreateCatalogResponsePrivilegesItem = `CREATE_TABLE`

const CreateCatalogResponsePrivilegesItemCreateMount CreateCatalogResponsePrivilegesItem = `CREATE_MOUNT`

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


const CreateMetastoreRequestPrivilegesItemUnknownPrivilege CreateMetastoreRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateMetastoreRequestPrivilegesItemSelect CreateMetastoreRequestPrivilegesItem = `SELECT`

const CreateMetastoreRequestPrivilegesItemCreate CreateMetastoreRequestPrivilegesItem = `CREATE`

const CreateMetastoreRequestPrivilegesItemModify CreateMetastoreRequestPrivilegesItem = `MODIFY`

const CreateMetastoreRequestPrivilegesItemUsage CreateMetastoreRequestPrivilegesItem = `USAGE`

const CreateMetastoreRequestPrivilegesItemReadFiles CreateMetastoreRequestPrivilegesItem = `READ_FILES`

const CreateMetastoreRequestPrivilegesItemWriteFiles CreateMetastoreRequestPrivilegesItem = `WRITE_FILES`

const CreateMetastoreRequestPrivilegesItemCreateTable CreateMetastoreRequestPrivilegesItem = `CREATE_TABLE`

const CreateMetastoreRequestPrivilegesItemCreateMount CreateMetastoreRequestPrivilegesItem = `CREATE_MOUNT`

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


const CreateMetastoreResponsePrivilegesItemUnknownPrivilege CreateMetastoreResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateMetastoreResponsePrivilegesItemSelect CreateMetastoreResponsePrivilegesItem = `SELECT`

const CreateMetastoreResponsePrivilegesItemCreate CreateMetastoreResponsePrivilegesItem = `CREATE`

const CreateMetastoreResponsePrivilegesItemModify CreateMetastoreResponsePrivilegesItem = `MODIFY`

const CreateMetastoreResponsePrivilegesItemUsage CreateMetastoreResponsePrivilegesItem = `USAGE`

const CreateMetastoreResponsePrivilegesItemReadFiles CreateMetastoreResponsePrivilegesItem = `READ_FILES`

const CreateMetastoreResponsePrivilegesItemWriteFiles CreateMetastoreResponsePrivilegesItem = `WRITE_FILES`

const CreateMetastoreResponsePrivilegesItemCreateTable CreateMetastoreResponsePrivilegesItem = `CREATE_TABLE`

const CreateMetastoreResponsePrivilegesItemCreateMount CreateMetastoreResponsePrivilegesItem = `CREATE_MOUNT`

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


const CreateProviderRequestAuthenticationTypeUnknown CreateProviderRequestAuthenticationType = `UNKNOWN`

const CreateProviderRequestAuthenticationTypeToken CreateProviderRequestAuthenticationType = `TOKEN`

const CreateProviderRequestAuthenticationTypeDatabricks CreateProviderRequestAuthenticationType = `DATABRICKS`

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


const CreateProviderResponseAuthenticationTypeUnknown CreateProviderResponseAuthenticationType = `UNKNOWN`

const CreateProviderResponseAuthenticationTypeToken CreateProviderResponseAuthenticationType = `TOKEN`

const CreateProviderResponseAuthenticationTypeDatabricks CreateProviderResponseAuthenticationType = `DATABRICKS`

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


const CreateRecipientRequestAuthenticationTypeUnknown CreateRecipientRequestAuthenticationType = `UNKNOWN`

const CreateRecipientRequestAuthenticationTypeToken CreateRecipientRequestAuthenticationType = `TOKEN`

const CreateRecipientRequestAuthenticationTypeDatabricks CreateRecipientRequestAuthenticationType = `DATABRICKS`

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


const CreateRecipientResponseAuthenticationTypeUnknown CreateRecipientResponseAuthenticationType = `UNKNOWN`

const CreateRecipientResponseAuthenticationTypeToken CreateRecipientResponseAuthenticationType = `TOKEN`

const CreateRecipientResponseAuthenticationTypeDatabricks CreateRecipientResponseAuthenticationType = `DATABRICKS`

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


const CreateSchemaRequestPrivilegesItemUnknownPrivilege CreateSchemaRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateSchemaRequestPrivilegesItemSelect CreateSchemaRequestPrivilegesItem = `SELECT`

const CreateSchemaRequestPrivilegesItemCreate CreateSchemaRequestPrivilegesItem = `CREATE`

const CreateSchemaRequestPrivilegesItemModify CreateSchemaRequestPrivilegesItem = `MODIFY`

const CreateSchemaRequestPrivilegesItemUsage CreateSchemaRequestPrivilegesItem = `USAGE`

const CreateSchemaRequestPrivilegesItemReadFiles CreateSchemaRequestPrivilegesItem = `READ_FILES`

const CreateSchemaRequestPrivilegesItemWriteFiles CreateSchemaRequestPrivilegesItem = `WRITE_FILES`

const CreateSchemaRequestPrivilegesItemCreateTable CreateSchemaRequestPrivilegesItem = `CREATE_TABLE`

const CreateSchemaRequestPrivilegesItemCreateMount CreateSchemaRequestPrivilegesItem = `CREATE_MOUNT`

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


const CreateSchemaResponsePrivilegesItemUnknownPrivilege CreateSchemaResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateSchemaResponsePrivilegesItemSelect CreateSchemaResponsePrivilegesItem = `SELECT`

const CreateSchemaResponsePrivilegesItemCreate CreateSchemaResponsePrivilegesItem = `CREATE`

const CreateSchemaResponsePrivilegesItemModify CreateSchemaResponsePrivilegesItem = `MODIFY`

const CreateSchemaResponsePrivilegesItemUsage CreateSchemaResponsePrivilegesItem = `USAGE`

const CreateSchemaResponsePrivilegesItemReadFiles CreateSchemaResponsePrivilegesItem = `READ_FILES`

const CreateSchemaResponsePrivilegesItemWriteFiles CreateSchemaResponsePrivilegesItem = `WRITE_FILES`

const CreateSchemaResponsePrivilegesItemCreateTable CreateSchemaResponsePrivilegesItem = `CREATE_TABLE`

const CreateSchemaResponsePrivilegesItemCreateMount CreateSchemaResponsePrivilegesItem = `CREATE_MOUNT`

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
    
    AwsIamRole any /* ERROR */ `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal any /* ERROR */ `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential.
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch
    // milliseconds.
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator.
    CreatedBy string `json:"created_by,omitempty"`
    
    GcpServiceAccountKey any /* ERROR */ `json:"gcp_service_account_key,omitempty"`
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
    
    AwsIamRole any /* ERROR */ `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal any /* ERROR */ `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential.
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch
    // milliseconds.
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator.
    CreatedBy string `json:"created_by,omitempty"`
    
    GcpServiceAccountKey any /* ERROR */ `json:"gcp_service_account_key,omitempty"`
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


const CreateTableRequestDataSourceFormatUnknownDataSourceFormat CreateTableRequestDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

const CreateTableRequestDataSourceFormatDelta CreateTableRequestDataSourceFormat = `DELTA`

const CreateTableRequestDataSourceFormatCsv CreateTableRequestDataSourceFormat = `CSV`

const CreateTableRequestDataSourceFormatJson CreateTableRequestDataSourceFormat = `JSON`

const CreateTableRequestDataSourceFormatAvro CreateTableRequestDataSourceFormat = `AVRO`

const CreateTableRequestDataSourceFormatParquet CreateTableRequestDataSourceFormat = `PARQUET`

const CreateTableRequestDataSourceFormatOrc CreateTableRequestDataSourceFormat = `ORC`

const CreateTableRequestDataSourceFormatText CreateTableRequestDataSourceFormat = `TEXT`

const CreateTableRequestDataSourceFormatUnityCatalog CreateTableRequestDataSourceFormat = `UNITY_CATALOG`

const CreateTableRequestDataSourceFormatDeltasharing CreateTableRequestDataSourceFormat = `DELTASHARING`

type CreateTableRequestPrivilegesItem string


const CreateTableRequestPrivilegesItemUnknownPrivilege CreateTableRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateTableRequestPrivilegesItemSelect CreateTableRequestPrivilegesItem = `SELECT`

const CreateTableRequestPrivilegesItemCreate CreateTableRequestPrivilegesItem = `CREATE`

const CreateTableRequestPrivilegesItemModify CreateTableRequestPrivilegesItem = `MODIFY`

const CreateTableRequestPrivilegesItemUsage CreateTableRequestPrivilegesItem = `USAGE`

const CreateTableRequestPrivilegesItemReadFiles CreateTableRequestPrivilegesItem = `READ_FILES`

const CreateTableRequestPrivilegesItemWriteFiles CreateTableRequestPrivilegesItem = `WRITE_FILES`

const CreateTableRequestPrivilegesItemCreateTable CreateTableRequestPrivilegesItem = `CREATE_TABLE`

const CreateTableRequestPrivilegesItemCreateMount CreateTableRequestPrivilegesItem = `CREATE_MOUNT`
// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type CreateTableRequestTableType string


const CreateTableRequestTableTypeUnknownTableType CreateTableRequestTableType = `UNKNOWN_TABLE_TYPE`

const CreateTableRequestTableTypeManaged CreateTableRequestTableType = `MANAGED`

const CreateTableRequestTableTypeExternal CreateTableRequestTableType = `EXTERNAL`

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


const CreateTableResponseDataSourceFormatUnknownDataSourceFormat CreateTableResponseDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

const CreateTableResponseDataSourceFormatDelta CreateTableResponseDataSourceFormat = `DELTA`

const CreateTableResponseDataSourceFormatCsv CreateTableResponseDataSourceFormat = `CSV`

const CreateTableResponseDataSourceFormatJson CreateTableResponseDataSourceFormat = `JSON`

const CreateTableResponseDataSourceFormatAvro CreateTableResponseDataSourceFormat = `AVRO`

const CreateTableResponseDataSourceFormatParquet CreateTableResponseDataSourceFormat = `PARQUET`

const CreateTableResponseDataSourceFormatOrc CreateTableResponseDataSourceFormat = `ORC`

const CreateTableResponseDataSourceFormatText CreateTableResponseDataSourceFormat = `TEXT`

const CreateTableResponseDataSourceFormatUnityCatalog CreateTableResponseDataSourceFormat = `UNITY_CATALOG`

const CreateTableResponseDataSourceFormatDeltasharing CreateTableResponseDataSourceFormat = `DELTASHARING`

type CreateTableResponsePrivilegesItem string


const CreateTableResponsePrivilegesItemUnknownPrivilege CreateTableResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const CreateTableResponsePrivilegesItemSelect CreateTableResponsePrivilegesItem = `SELECT`

const CreateTableResponsePrivilegesItemCreate CreateTableResponsePrivilegesItem = `CREATE`

const CreateTableResponsePrivilegesItemModify CreateTableResponsePrivilegesItem = `MODIFY`

const CreateTableResponsePrivilegesItemUsage CreateTableResponsePrivilegesItem = `USAGE`

const CreateTableResponsePrivilegesItemReadFiles CreateTableResponsePrivilegesItem = `READ_FILES`

const CreateTableResponsePrivilegesItemWriteFiles CreateTableResponsePrivilegesItem = `WRITE_FILES`

const CreateTableResponsePrivilegesItemCreateTable CreateTableResponsePrivilegesItem = `CREATE_TABLE`

const CreateTableResponsePrivilegesItemCreateMount CreateTableResponsePrivilegesItem = `CREATE_MOUNT`
// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type CreateTableResponseTableType string


const CreateTableResponseTableTypeUnknownTableType CreateTableResponseTableType = `UNKNOWN_TABLE_TYPE`

const CreateTableResponseTableTypeManaged CreateTableResponseTableType = `MANAGED`

const CreateTableResponseTableTypeExternal CreateTableResponseTableType = `EXTERNAL`

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


const GetCatalogResponseCatalogTypeUnknownCatalogType GetCatalogResponseCatalogType = `UNKNOWN_CATALOG_TYPE`

const GetCatalogResponseCatalogTypeManagedCatalog GetCatalogResponseCatalogType = `MANAGED_CATALOG`

const GetCatalogResponseCatalogTypeDeltasharingCatalog GetCatalogResponseCatalogType = `DELTASHARING_CATALOG`

const GetCatalogResponseCatalogTypeSystemCatalog GetCatalogResponseCatalogType = `SYSTEM_CATALOG`

type GetCatalogResponsePrivilegesItem string


const GetCatalogResponsePrivilegesItemUnknownPrivilege GetCatalogResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const GetCatalogResponsePrivilegesItemSelect GetCatalogResponsePrivilegesItem = `SELECT`

const GetCatalogResponsePrivilegesItemCreate GetCatalogResponsePrivilegesItem = `CREATE`

const GetCatalogResponsePrivilegesItemModify GetCatalogResponsePrivilegesItem = `MODIFY`

const GetCatalogResponsePrivilegesItemUsage GetCatalogResponsePrivilegesItem = `USAGE`

const GetCatalogResponsePrivilegesItemReadFiles GetCatalogResponsePrivilegesItem = `READ_FILES`

const GetCatalogResponsePrivilegesItemWriteFiles GetCatalogResponsePrivilegesItem = `WRITE_FILES`

const GetCatalogResponsePrivilegesItemCreateTable GetCatalogResponsePrivilegesItem = `CREATE_TABLE`

const GetCatalogResponsePrivilegesItemCreateMount GetCatalogResponsePrivilegesItem = `CREATE_MOUNT`

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


const GetMetastoreResponsePrivilegesItemUnknownPrivilege GetMetastoreResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const GetMetastoreResponsePrivilegesItemSelect GetMetastoreResponsePrivilegesItem = `SELECT`

const GetMetastoreResponsePrivilegesItemCreate GetMetastoreResponsePrivilegesItem = `CREATE`

const GetMetastoreResponsePrivilegesItemModify GetMetastoreResponsePrivilegesItem = `MODIFY`

const GetMetastoreResponsePrivilegesItemUsage GetMetastoreResponsePrivilegesItem = `USAGE`

const GetMetastoreResponsePrivilegesItemReadFiles GetMetastoreResponsePrivilegesItem = `READ_FILES`

const GetMetastoreResponsePrivilegesItemWriteFiles GetMetastoreResponsePrivilegesItem = `WRITE_FILES`

const GetMetastoreResponsePrivilegesItemCreateTable GetMetastoreResponsePrivilegesItem = `CREATE_TABLE`

const GetMetastoreResponsePrivilegesItemCreateMount GetMetastoreResponsePrivilegesItem = `CREATE_MOUNT`

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


const GetProviderResponseAuthenticationTypeUnknown GetProviderResponseAuthenticationType = `UNKNOWN`

const GetProviderResponseAuthenticationTypeToken GetProviderResponseAuthenticationType = `TOKEN`

const GetProviderResponseAuthenticationTypeDatabricks GetProviderResponseAuthenticationType = `DATABRICKS`

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


const GetRecipientResponseAuthenticationTypeUnknown GetRecipientResponseAuthenticationType = `UNKNOWN`

const GetRecipientResponseAuthenticationTypeToken GetRecipientResponseAuthenticationType = `TOKEN`

const GetRecipientResponseAuthenticationTypeDatabricks GetRecipientResponseAuthenticationType = `DATABRICKS`

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


const GetSchemaResponsePrivilegesItemUnknownPrivilege GetSchemaResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const GetSchemaResponsePrivilegesItemSelect GetSchemaResponsePrivilegesItem = `SELECT`

const GetSchemaResponsePrivilegesItemCreate GetSchemaResponsePrivilegesItem = `CREATE`

const GetSchemaResponsePrivilegesItemModify GetSchemaResponsePrivilegesItem = `MODIFY`

const GetSchemaResponsePrivilegesItemUsage GetSchemaResponsePrivilegesItem = `USAGE`

const GetSchemaResponsePrivilegesItemReadFiles GetSchemaResponsePrivilegesItem = `READ_FILES`

const GetSchemaResponsePrivilegesItemWriteFiles GetSchemaResponsePrivilegesItem = `WRITE_FILES`

const GetSchemaResponsePrivilegesItemCreateTable GetSchemaResponsePrivilegesItem = `CREATE_TABLE`

const GetSchemaResponsePrivilegesItemCreateMount GetSchemaResponsePrivilegesItem = `CREATE_MOUNT`

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
    
    AwsIamRole any /* ERROR */ `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal any /* ERROR */ `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential.
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch
    // milliseconds.
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator.
    CreatedBy string `json:"created_by,omitempty"`
    
    GcpServiceAccountKey any /* ERROR */ `json:"gcp_service_account_key,omitempty"`
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


const GetTableResponseDataSourceFormatUnknownDataSourceFormat GetTableResponseDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

const GetTableResponseDataSourceFormatDelta GetTableResponseDataSourceFormat = `DELTA`

const GetTableResponseDataSourceFormatCsv GetTableResponseDataSourceFormat = `CSV`

const GetTableResponseDataSourceFormatJson GetTableResponseDataSourceFormat = `JSON`

const GetTableResponseDataSourceFormatAvro GetTableResponseDataSourceFormat = `AVRO`

const GetTableResponseDataSourceFormatParquet GetTableResponseDataSourceFormat = `PARQUET`

const GetTableResponseDataSourceFormatOrc GetTableResponseDataSourceFormat = `ORC`

const GetTableResponseDataSourceFormatText GetTableResponseDataSourceFormat = `TEXT`

const GetTableResponseDataSourceFormatUnityCatalog GetTableResponseDataSourceFormat = `UNITY_CATALOG`

const GetTableResponseDataSourceFormatDeltasharing GetTableResponseDataSourceFormat = `DELTASHARING`

type GetTableResponsePrivilegesItem string


const GetTableResponsePrivilegesItemUnknownPrivilege GetTableResponsePrivilegesItem = `UNKNOWN_PRIVILEGE`

const GetTableResponsePrivilegesItemSelect GetTableResponsePrivilegesItem = `SELECT`

const GetTableResponsePrivilegesItemCreate GetTableResponsePrivilegesItem = `CREATE`

const GetTableResponsePrivilegesItemModify GetTableResponsePrivilegesItem = `MODIFY`

const GetTableResponsePrivilegesItemUsage GetTableResponsePrivilegesItem = `USAGE`

const GetTableResponsePrivilegesItemReadFiles GetTableResponsePrivilegesItem = `READ_FILES`

const GetTableResponsePrivilegesItemWriteFiles GetTableResponsePrivilegesItem = `WRITE_FILES`

const GetTableResponsePrivilegesItemCreateTable GetTableResponsePrivilegesItem = `CREATE_TABLE`

const GetTableResponsePrivilegesItemCreateMount GetTableResponsePrivilegesItem = `CREATE_MOUNT`
// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type GetTableResponseTableType string


const GetTableResponseTableTypeUnknownTableType GetTableResponseTableType = `UNKNOWN_TABLE_TYPE`

const GetTableResponseTableTypeManaged GetTableResponseTableType = `MANAGED`

const GetTableResponseTableTypeExternal GetTableResponseTableType = `EXTERNAL`

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


const MetastoreInfoPrivilegesItemUnknownPrivilege MetastoreInfoPrivilegesItem = `UNKNOWN_PRIVILEGE`

const MetastoreInfoPrivilegesItemSelect MetastoreInfoPrivilegesItem = `SELECT`

const MetastoreInfoPrivilegesItemCreate MetastoreInfoPrivilegesItem = `CREATE`

const MetastoreInfoPrivilegesItemModify MetastoreInfoPrivilegesItem = `MODIFY`

const MetastoreInfoPrivilegesItemUsage MetastoreInfoPrivilegesItem = `USAGE`

const MetastoreInfoPrivilegesItemReadFiles MetastoreInfoPrivilegesItem = `READ_FILES`

const MetastoreInfoPrivilegesItemWriteFiles MetastoreInfoPrivilegesItem = `WRITE_FILES`

const MetastoreInfoPrivilegesItemCreateTable MetastoreInfoPrivilegesItem = `CREATE_TABLE`

const MetastoreInfoPrivilegesItemCreateMount MetastoreInfoPrivilegesItem = `CREATE_MOUNT`

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


const PermissionsChangeAddItemUnknownPrivilege PermissionsChangeAddItem = `UNKNOWN_PRIVILEGE`

const PermissionsChangeAddItemSelect PermissionsChangeAddItem = `SELECT`

const PermissionsChangeAddItemCreate PermissionsChangeAddItem = `CREATE`

const PermissionsChangeAddItemModify PermissionsChangeAddItem = `MODIFY`

const PermissionsChangeAddItemUsage PermissionsChangeAddItem = `USAGE`

const PermissionsChangeAddItemReadFiles PermissionsChangeAddItem = `READ_FILES`

const PermissionsChangeAddItemWriteFiles PermissionsChangeAddItem = `WRITE_FILES`

const PermissionsChangeAddItemCreateTable PermissionsChangeAddItem = `CREATE_TABLE`

const PermissionsChangeAddItemCreateMount PermissionsChangeAddItem = `CREATE_MOUNT`

type PermissionsChangeRemoveItem string


const PermissionsChangeRemoveItemUnknownPrivilege PermissionsChangeRemoveItem = `UNKNOWN_PRIVILEGE`

const PermissionsChangeRemoveItemSelect PermissionsChangeRemoveItem = `SELECT`

const PermissionsChangeRemoveItemCreate PermissionsChangeRemoveItem = `CREATE`

const PermissionsChangeRemoveItemModify PermissionsChangeRemoveItem = `MODIFY`

const PermissionsChangeRemoveItemUsage PermissionsChangeRemoveItem = `USAGE`

const PermissionsChangeRemoveItemReadFiles PermissionsChangeRemoveItem = `READ_FILES`

const PermissionsChangeRemoveItemWriteFiles PermissionsChangeRemoveItem = `WRITE_FILES`

const PermissionsChangeRemoveItemCreateTable PermissionsChangeRemoveItem = `CREATE_TABLE`

const PermissionsChangeRemoveItemCreateMount PermissionsChangeRemoveItem = `CREATE_MOUNT`

type PrivilegeAssignment struct {
    // The principal (user email address or group name).
    Principal string `json:"principal,omitempty"`
    // The privileges assigned to the principal.
    Privileges []PrivilegeAssignmentPrivilegesItem `json:"privileges,omitempty"`
}


type PrivilegeAssignmentPrivilegesItem string


const PrivilegeAssignmentPrivilegesItemUnknownPrivilege PrivilegeAssignmentPrivilegesItem = `UNKNOWN_PRIVILEGE`

const PrivilegeAssignmentPrivilegesItemSelect PrivilegeAssignmentPrivilegesItem = `SELECT`

const PrivilegeAssignmentPrivilegesItemCreate PrivilegeAssignmentPrivilegesItem = `CREATE`

const PrivilegeAssignmentPrivilegesItemModify PrivilegeAssignmentPrivilegesItem = `MODIFY`

const PrivilegeAssignmentPrivilegesItemUsage PrivilegeAssignmentPrivilegesItem = `USAGE`

const PrivilegeAssignmentPrivilegesItemReadFiles PrivilegeAssignmentPrivilegesItem = `READ_FILES`

const PrivilegeAssignmentPrivilegesItemWriteFiles PrivilegeAssignmentPrivilegesItem = `WRITE_FILES`

const PrivilegeAssignmentPrivilegesItemCreateTable PrivilegeAssignmentPrivilegesItem = `CREATE_TABLE`

const PrivilegeAssignmentPrivilegesItemCreateMount PrivilegeAssignmentPrivilegesItem = `CREATE_MOUNT`

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


const ProviderInfoAuthenticationTypeUnknown ProviderInfoAuthenticationType = `UNKNOWN`

const ProviderInfoAuthenticationTypeToken ProviderInfoAuthenticationType = `TOKEN`

const ProviderInfoAuthenticationTypeDatabricks ProviderInfoAuthenticationType = `DATABRICKS`

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


const RecipientInfoAuthenticationTypeUnknown RecipientInfoAuthenticationType = `UNKNOWN`

const RecipientInfoAuthenticationTypeToken RecipientInfoAuthenticationType = `TOKEN`

const RecipientInfoAuthenticationTypeDatabricks RecipientInfoAuthenticationType = `DATABRICKS`

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


const RotateRecipientTokenResponseAuthenticationTypeUnknown RotateRecipientTokenResponseAuthenticationType = `UNKNOWN`

const RotateRecipientTokenResponseAuthenticationTypeToken RotateRecipientTokenResponseAuthenticationType = `TOKEN`

const RotateRecipientTokenResponseAuthenticationTypeDatabricks RotateRecipientTokenResponseAuthenticationType = `DATABRICKS`

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


const SchemaInfoPrivilegesItemUnknownPrivilege SchemaInfoPrivilegesItem = `UNKNOWN_PRIVILEGE`

const SchemaInfoPrivilegesItemSelect SchemaInfoPrivilegesItem = `SELECT`

const SchemaInfoPrivilegesItemCreate SchemaInfoPrivilegesItem = `CREATE`

const SchemaInfoPrivilegesItemModify SchemaInfoPrivilegesItem = `MODIFY`

const SchemaInfoPrivilegesItemUsage SchemaInfoPrivilegesItem = `USAGE`

const SchemaInfoPrivilegesItemReadFiles SchemaInfoPrivilegesItem = `READ_FILES`

const SchemaInfoPrivilegesItemWriteFiles SchemaInfoPrivilegesItem = `WRITE_FILES`

const SchemaInfoPrivilegesItemCreateTable SchemaInfoPrivilegesItem = `CREATE_TABLE`

const SchemaInfoPrivilegesItemCreateMount SchemaInfoPrivilegesItem = `CREATE_MOUNT`

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
    
    AwsIamRole any /* ERROR */ `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal any /* ERROR */ `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential.
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch
    // milliseconds.
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator.
    CreatedBy string `json:"created_by,omitempty"`
    
    GcpServiceAccountKey any /* ERROR */ `json:"gcp_service_account_key,omitempty"`
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


const TableInfoDataSourceFormatUnknownDataSourceFormat TableInfoDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

const TableInfoDataSourceFormatDelta TableInfoDataSourceFormat = `DELTA`

const TableInfoDataSourceFormatCsv TableInfoDataSourceFormat = `CSV`

const TableInfoDataSourceFormatJson TableInfoDataSourceFormat = `JSON`

const TableInfoDataSourceFormatAvro TableInfoDataSourceFormat = `AVRO`

const TableInfoDataSourceFormatParquet TableInfoDataSourceFormat = `PARQUET`

const TableInfoDataSourceFormatOrc TableInfoDataSourceFormat = `ORC`

const TableInfoDataSourceFormatText TableInfoDataSourceFormat = `TEXT`

const TableInfoDataSourceFormatUnityCatalog TableInfoDataSourceFormat = `UNITY_CATALOG`

const TableInfoDataSourceFormatDeltasharing TableInfoDataSourceFormat = `DELTASHARING`

type TableInfoPrivilegesItem string


const TableInfoPrivilegesItemUnknownPrivilege TableInfoPrivilegesItem = `UNKNOWN_PRIVILEGE`

const TableInfoPrivilegesItemSelect TableInfoPrivilegesItem = `SELECT`

const TableInfoPrivilegesItemCreate TableInfoPrivilegesItem = `CREATE`

const TableInfoPrivilegesItemModify TableInfoPrivilegesItem = `MODIFY`

const TableInfoPrivilegesItemUsage TableInfoPrivilegesItem = `USAGE`

const TableInfoPrivilegesItemReadFiles TableInfoPrivilegesItem = `READ_FILES`

const TableInfoPrivilegesItemWriteFiles TableInfoPrivilegesItem = `WRITE_FILES`

const TableInfoPrivilegesItemCreateTable TableInfoPrivilegesItem = `CREATE_TABLE`

const TableInfoPrivilegesItemCreateMount TableInfoPrivilegesItem = `CREATE_MOUNT`
// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type TableInfoTableType string


const TableInfoTableTypeUnknownTableType TableInfoTableType = `UNKNOWN_TABLE_TYPE`

const TableInfoTableTypeManaged TableInfoTableType = `MANAGED`

const TableInfoTableTypeExternal TableInfoTableType = `EXTERNAL`

const TableInfoTableTypeView TableInfoTableType = `VIEW`

type TableSummary struct {
    
    FullName string `json:"full_name,omitempty"`
    
    TableType TableSummaryTableType `json:"table_type,omitempty"`
}


type TableSummaryTableType string


const TableSummaryTableTypeUnknownTableType TableSummaryTableType = `UNKNOWN_TABLE_TYPE`

const TableSummaryTableTypeManaged TableSummaryTableType = `MANAGED`

const TableSummaryTableTypeExternal TableSummaryTableType = `EXTERNAL`

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


const UpdateCatalogRequestCatalogTypeUnknownCatalogType UpdateCatalogRequestCatalogType = `UNKNOWN_CATALOG_TYPE`

const UpdateCatalogRequestCatalogTypeManagedCatalog UpdateCatalogRequestCatalogType = `MANAGED_CATALOG`

const UpdateCatalogRequestCatalogTypeDeltasharingCatalog UpdateCatalogRequestCatalogType = `DELTASHARING_CATALOG`

const UpdateCatalogRequestCatalogTypeSystemCatalog UpdateCatalogRequestCatalogType = `SYSTEM_CATALOG`

type UpdateCatalogRequestPrivilegesItem string


const UpdateCatalogRequestPrivilegesItemUnknownPrivilege UpdateCatalogRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateCatalogRequestPrivilegesItemSelect UpdateCatalogRequestPrivilegesItem = `SELECT`

const UpdateCatalogRequestPrivilegesItemCreate UpdateCatalogRequestPrivilegesItem = `CREATE`

const UpdateCatalogRequestPrivilegesItemModify UpdateCatalogRequestPrivilegesItem = `MODIFY`

const UpdateCatalogRequestPrivilegesItemUsage UpdateCatalogRequestPrivilegesItem = `USAGE`

const UpdateCatalogRequestPrivilegesItemReadFiles UpdateCatalogRequestPrivilegesItem = `READ_FILES`

const UpdateCatalogRequestPrivilegesItemWriteFiles UpdateCatalogRequestPrivilegesItem = `WRITE_FILES`

const UpdateCatalogRequestPrivilegesItemCreateTable UpdateCatalogRequestPrivilegesItem = `CREATE_TABLE`

const UpdateCatalogRequestPrivilegesItemCreateMount UpdateCatalogRequestPrivilegesItem = `CREATE_MOUNT`

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


const UpdateMetastoreRequestPrivilegesItemUnknownPrivilege UpdateMetastoreRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateMetastoreRequestPrivilegesItemSelect UpdateMetastoreRequestPrivilegesItem = `SELECT`

const UpdateMetastoreRequestPrivilegesItemCreate UpdateMetastoreRequestPrivilegesItem = `CREATE`

const UpdateMetastoreRequestPrivilegesItemModify UpdateMetastoreRequestPrivilegesItem = `MODIFY`

const UpdateMetastoreRequestPrivilegesItemUsage UpdateMetastoreRequestPrivilegesItem = `USAGE`

const UpdateMetastoreRequestPrivilegesItemReadFiles UpdateMetastoreRequestPrivilegesItem = `READ_FILES`

const UpdateMetastoreRequestPrivilegesItemWriteFiles UpdateMetastoreRequestPrivilegesItem = `WRITE_FILES`

const UpdateMetastoreRequestPrivilegesItemCreateTable UpdateMetastoreRequestPrivilegesItem = `CREATE_TABLE`

const UpdateMetastoreRequestPrivilegesItemCreateMount UpdateMetastoreRequestPrivilegesItem = `CREATE_MOUNT`

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


const UpdateProviderRequestAuthenticationTypeUnknown UpdateProviderRequestAuthenticationType = `UNKNOWN`

const UpdateProviderRequestAuthenticationTypeToken UpdateProviderRequestAuthenticationType = `TOKEN`

const UpdateProviderRequestAuthenticationTypeDatabricks UpdateProviderRequestAuthenticationType = `DATABRICKS`

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


const UpdateRecipientRequestAuthenticationTypeUnknown UpdateRecipientRequestAuthenticationType = `UNKNOWN`

const UpdateRecipientRequestAuthenticationTypeToken UpdateRecipientRequestAuthenticationType = `TOKEN`

const UpdateRecipientRequestAuthenticationTypeDatabricks UpdateRecipientRequestAuthenticationType = `DATABRICKS`

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


const UpdateSchemaRequestPrivilegesItemUnknownPrivilege UpdateSchemaRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateSchemaRequestPrivilegesItemSelect UpdateSchemaRequestPrivilegesItem = `SELECT`

const UpdateSchemaRequestPrivilegesItemCreate UpdateSchemaRequestPrivilegesItem = `CREATE`

const UpdateSchemaRequestPrivilegesItemModify UpdateSchemaRequestPrivilegesItem = `MODIFY`

const UpdateSchemaRequestPrivilegesItemUsage UpdateSchemaRequestPrivilegesItem = `USAGE`

const UpdateSchemaRequestPrivilegesItemReadFiles UpdateSchemaRequestPrivilegesItem = `READ_FILES`

const UpdateSchemaRequestPrivilegesItemWriteFiles UpdateSchemaRequestPrivilegesItem = `WRITE_FILES`

const UpdateSchemaRequestPrivilegesItemCreateTable UpdateSchemaRequestPrivilegesItem = `CREATE_TABLE`

const UpdateSchemaRequestPrivilegesItemCreateMount UpdateSchemaRequestPrivilegesItem = `CREATE_MOUNT`

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
    
    AwsIamRole any /* ERROR */ `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal any /* ERROR */ `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential.
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch
    // milliseconds.
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator.
    CreatedBy string `json:"created_by,omitempty"`
    
    GcpServiceAccountKey any /* ERROR */ `json:"gcp_service_account_key,omitempty"`
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


const UpdateTableRequestDataSourceFormatUnknownDataSourceFormat UpdateTableRequestDataSourceFormat = `UNKNOWN_DATA_SOURCE_FORMAT`

const UpdateTableRequestDataSourceFormatDelta UpdateTableRequestDataSourceFormat = `DELTA`

const UpdateTableRequestDataSourceFormatCsv UpdateTableRequestDataSourceFormat = `CSV`

const UpdateTableRequestDataSourceFormatJson UpdateTableRequestDataSourceFormat = `JSON`

const UpdateTableRequestDataSourceFormatAvro UpdateTableRequestDataSourceFormat = `AVRO`

const UpdateTableRequestDataSourceFormatParquet UpdateTableRequestDataSourceFormat = `PARQUET`

const UpdateTableRequestDataSourceFormatOrc UpdateTableRequestDataSourceFormat = `ORC`

const UpdateTableRequestDataSourceFormatText UpdateTableRequestDataSourceFormat = `TEXT`

const UpdateTableRequestDataSourceFormatUnityCatalog UpdateTableRequestDataSourceFormat = `UNITY_CATALOG`

const UpdateTableRequestDataSourceFormatDeltasharing UpdateTableRequestDataSourceFormat = `DELTASHARING`

type UpdateTableRequestPrivilegesItem string


const UpdateTableRequestPrivilegesItemUnknownPrivilege UpdateTableRequestPrivilegesItem = `UNKNOWN_PRIVILEGE`

const UpdateTableRequestPrivilegesItemSelect UpdateTableRequestPrivilegesItem = `SELECT`

const UpdateTableRequestPrivilegesItemCreate UpdateTableRequestPrivilegesItem = `CREATE`

const UpdateTableRequestPrivilegesItemModify UpdateTableRequestPrivilegesItem = `MODIFY`

const UpdateTableRequestPrivilegesItemUsage UpdateTableRequestPrivilegesItem = `USAGE`

const UpdateTableRequestPrivilegesItemReadFiles UpdateTableRequestPrivilegesItem = `READ_FILES`

const UpdateTableRequestPrivilegesItemWriteFiles UpdateTableRequestPrivilegesItem = `WRITE_FILES`

const UpdateTableRequestPrivilegesItemCreateTable UpdateTableRequestPrivilegesItem = `CREATE_TABLE`

const UpdateTableRequestPrivilegesItemCreateMount UpdateTableRequestPrivilegesItem = `CREATE_MOUNT`
// [Create:REQ Update:OPT] Table type (&#34;MANAGED&#34;, &#34;EXTERNAL&#34;, &#34;VIEW&#34;).
type UpdateTableRequestTableType string


const UpdateTableRequestTableTypeUnknownTableType UpdateTableRequestTableType = `UNKNOWN_TABLE_TYPE`

const UpdateTableRequestTableTypeManaged UpdateTableRequestTableType = `MANAGED`

const UpdateTableRequestTableTypeExternal UpdateTableRequestTableType = `EXTERNAL`

const UpdateTableRequestTableTypeView UpdateTableRequestTableType = `VIEW`
