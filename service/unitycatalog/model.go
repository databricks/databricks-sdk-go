// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

// all definitions in this file are in alphabetical order

type AwsIamRole struct {
    // The external ID used in role assumption to prevent confused deputy 
    // problem. [Create:IGN]. 
    ExternalId string `json:"external_id,omitempty"`
    // The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access. 
    // [Create:REQ]. 
    RoleArn string `json:"role_arn,omitempty"`
    // The Amazon Resource Name (ARN) of the AWS IAM user managed by 
    // Databricks. This is the identity that is going to assume the AWS IAM 
    // role. [Create:IGN]. 
    UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`
}


type AzureServicePrincipal struct {
    // The application ID of the application registration within the referenced 
    // AAD tenant. [Create:REQ] 
    ApplicationId string `json:"application_id,omitempty"`
    // The client secret generated for the above app ID in AAD. [Create:REQ] 
    ClientSecret string `json:"client_secret,omitempty"`
    // The directory ID corresponding to the Azure Active Directory (AAD) 
    // tenant of the application. [Create:REQ]. 
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
    // [Create: OPT, Update: OPT] Digits to right of decimal; Required on 
    // Create for DecimalTypes. 
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
    // [Create:IGN Update:OPT] UUID of storage credential to access 
    // storage_root 
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
    // [Create:IGN Update:OPT] UUID of storage credential to access 
    // storage_root 
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
    
    AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator. 
    CreatedBy string `json:"created_by,omitempty"`
    
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
    // Optional. Supplying true to this argument skips validation of the 
    // created set of credentials. 
    SkipValidation bool `json:"skip_validation,omitempty"`
    // [Create,Update:IGN] Time at which this credential was last modified, in 
    // epoch milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create,Update:IGN] Username of user who last modified the credential. 
    UpdatedBy string `json:"updated_by,omitempty"`
}


type CreateStorageCredentialResponse struct {
    
    AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator. 
    CreatedBy string `json:"created_by,omitempty"`
    
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
    // [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, 
    // EXTERNAL tables) 
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
    // [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, 
    // EXTERNAL tables) 
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
    // Required. Name of the catalog. 
    NameArg string ` path:"name_arg"`
}


type DeleteExternalLocationRequest struct {
    // Optional. Force deletion even if there are dependent external tables or 
    // mounts. 
    Force bool `json:"force,omitempty"`
    // Required. Name of the external location. 
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


type DeleteSchemaRequest struct {
    // Required. Full name of the schema to delete. 
    FullNameArg string ` path:"full_name_arg"`
}


type DeleteStorageCredentialRequest struct {
    // Optional. Force deletion even if there are dependent external locations 
    // or external tables. 
    Force bool `json:"force,omitempty"`
    // Required. Name of the storage credential. 
    NameArg string ` path:"name_arg"`
}


type DeleteTableRequest struct {
    // Required. Full name of the table. 
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


type GcpServiceAccountKey struct {
    // The email of the service account. [Create:REQ]. 
    Email string `json:"email,omitempty"`
    // The service account&#39;s RSA private key. [Create:REQ] 
    PrivateKey string `json:"private_key,omitempty"`
    // The ID of the service account&#39;s private key. [Create:REQ] 
    PrivateKeyId string `json:"private_key_id,omitempty"`
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
    // [Create:IGN Update:OPT] UUID of storage credential to access 
    // storage_root 
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
    Principal string ` query:"principal,omitempty"`
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

type GetStorageCredentialRequest struct {
    // Required. Name of the storage credential. 
    NameArg string ` path:"name_arg"`
}


type GetStorageCredentialResponse struct {
    
    AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator. 
    CreatedBy string `json:"created_by,omitempty"`
    
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
    // [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, 
    // EXTERNAL tables) 
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

type ListCatalogsResponse struct {
    
    Catalogs []CatalogInfo `json:"catalogs,omitempty"`
}


type ListExternalLocationsResponse struct {
    
    ExternalLocations []ExternalLocationInfo `json:"external_locations,omitempty"`
}


type ListMetastoresResponse struct {
    
    Metastores []MetastoreInfo `json:"metastores,omitempty"`
}


type ListSchemasRequest struct {
    // Optional. Parent catalog for schemas of interest. 
    CatalogName string ` query:"catalog_name,omitempty"`
}


type ListSchemasResponse struct {
    
    Schemas []SchemaInfo `json:"schemas,omitempty"`
}


type ListStorageCredentialsResponse struct {
    // TODO: add pagination to UC list APIs. 
    StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`
}


type ListTablesRequest struct {
    // Required. Name of parent catalog for tables of interest. 
    CatalogName string ` query:"catalog_name,omitempty"`
    // Required (for now -- may be optional for wildcard search in future). 
    // Parent schema of tables. 
    SchemaName string ` query:"schema_name,omitempty"`
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
    // [Create:IGN Update:OPT] UUID of storage credential to access 
    // storage_root 
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

type StorageCredentialInfo struct {
    
    AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator. 
    CreatedBy string `json:"created_by,omitempty"`
    
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
    // [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, 
    // EXTERNAL tables) 
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
    // DBR (not removed below as it still works with CLI) Optional. Force 
    // update even if changing url invalidates dependent external tables or 
    // mounts. 
    Force bool `json:"force,omitempty"`
    // [Create,Update:IGN] Unique identifier of Metastore hosting the External 
    // Location. 
    MetastoreId string `json:"metastore_id,omitempty"`
    // [Create:REQ Update:OPT] Name of the External Location. 
    Name string `json:"name,omitempty"`
    // Required. Name of the external location. 
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
    // [Create:IGN Update:OPT] UUID of storage credential to access 
    // storage_root 
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
    // Required. Unique identifier (full name) of Securable (from URL). 
    SecurableFullName string ` path:"securable_full_name"`
    // Required. Type of Securable (from URL). 
    SecurableType string ` path:"securable_type"`
}


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

type UpdateStorageCredentialRequest struct {
    
    AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
    
    AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
    // [Create,Update:OPT] Comment associated with the credential. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Credential was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of credential creator. 
    CreatedBy string `json:"created_by,omitempty"`
    
    GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
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
    // Optional. Supplying true to this argument skips validation of the 
    // updated set of credentials. 
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
    // [Create:REQ Update:OPT] Storage root URL for table (for MANAGED, 
    // EXTERNAL tables) 
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
