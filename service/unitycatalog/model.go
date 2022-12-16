// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

import "fmt"

// all definitions in this file are in alphabetical order

// The delta sharing authentication type.
type AuthenticationType string

const AuthenticationTypeDatabricks AuthenticationType = `DATABRICKS`

const AuthenticationTypeToken AuthenticationType = `TOKEN`

const AuthenticationTypeUnknown AuthenticationType = `UNKNOWN`

// String representation for [fmt.Print]
func (at *AuthenticationType) String() string {
	return string(*at)
}

// Set raw string value and validate it against allowed values
func (at *AuthenticationType) Set(v string) error {
	switch v {
	case `DATABRICKS`, `TOKEN`, `UNKNOWN`:
		*at = AuthenticationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS", "TOKEN", "UNKNOWN"`, v)
	}
}

// Type always returns AuthenticationType to satisfy [pflag.Value] interface
func (at *AuthenticationType) Type() string {
	return "AuthenticationType"
}

type AwsIamRole struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string `json:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`
}

type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId string `json:"application_id"`
	// The client secret generated for the above app ID in AAD.
	ClientSecret string `json:"client_secret"`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId string `json:"directory_id"`
}

type CatalogInfo struct {
	// The type of the catalog.
	CatalogType CatalogType `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this Catalog was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Catalog creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Catalog.
	Name string `json:"name,omitempty"`
	// Username of current owner of Catalog.
	Owner string `json:"owner,omitempty"`

	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing Catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `json:"provider_name,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// Storage Location URL (full path) for managed tables within Catalog.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within Catalog.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this Catalog was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified Catalog.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// The type of the catalog.
type CatalogType string

const CatalogTypeDeltasharingCatalog CatalogType = `DELTASHARING_CATALOG`

const CatalogTypeManagedCatalog CatalogType = `MANAGED_CATALOG`

const CatalogTypeSystemCatalog CatalogType = `SYSTEM_CATALOG`

// String representation for [fmt.Print]
func (ct *CatalogType) String() string {
	return string(*ct)
}

// Set raw string value and validate it against allowed values
func (ct *CatalogType) Set(v string) error {
	switch v {
	case `DELTASHARING_CATALOG`, `MANAGED_CATALOG`, `SYSTEM_CATALOG`:
		*ct = CatalogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTASHARING_CATALOG", "MANAGED_CATALOG", "SYSTEM_CATALOG"`, v)
	}
}

// Type always returns CatalogType to satisfy [pflag.Value] interface
func (ct *CatalogType) Type() string {
	return "CatalogType"
}

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

const ColumnInfoTypeNameUserDefinedType ColumnInfoTypeName = `USER_DEFINED_TYPE`

// String representation for [fmt.Print]
func (citn *ColumnInfoTypeName) String() string {
	return string(*citn)
}

// Set raw string value and validate it against allowed values
func (citn *ColumnInfoTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TIMESTAMP`, `USER_DEFINED_TYPE`:
		*citn = ColumnInfoTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TIMESTAMP", "USER_DEFINED_TYPE"`, v)
	}
}

// Type always returns ColumnInfoTypeName to satisfy [pflag.Value] interface
func (citn *ColumnInfoTypeName) Type() string {
	return "ColumnInfoTypeName"
}

type CreateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of Catalog.
	Name string `json:"name"`

	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing Catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `json:"provider_name,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// Storage root URL for managed tables within Catalog.
	StorageRoot string `json:"storage_root,omitempty"`
}

type CreateExternalLocation struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Current name of the Storage Credential this location uses.
	CredentialName string `json:"credential_name"`
	// Name of the External Location.
	Name string `json:"name"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the External Location.
	Url string `json:"url"`
}

type CreateMetastore struct {
	// Name of Metastore.
	Name string `json:"name"`
	// Storage root URL for Metastore
	StorageRoot string `json:"storage_root"`
}

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the Metastore.
	DefaultCatalogName string `json:"default_catalog_name"`
	// The ID of the Metastore.
	MetastoreId string `json:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type CreateProvider struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// The name of the Provider.
	Name string `json:"name"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// This field is required when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
}

type CreateRecipient struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.\n
	// This field is only present when the authentication type is
	// `DATABRICKS`.\n The identifier is of format
	// <cloud>:<region>:<metastore-uuid>.
	DataRecipientGlobalMetastoreId any `json:"data_recipient_global_metastore_id,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Name of Recipient.
	Name string `json:"name"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the authentication type is `DATABRICKS`.
	SharingCode string `json:"sharing_code,omitempty"`
}

type CreateSchema struct {
	// Name of parent Catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of Schema, relative to parent Catalog.
	Name string `json:"name"`

	Properties map[string]string `json:"properties,omitempty"`
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the Share.
	Name string `json:"name"`
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// The credential name. The name MUST be unique within the Metastore.
	Name string `json:"name"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`
}

// Data source format
type DataSourceFormat string

const DataSourceFormatAvro DataSourceFormat = `AVRO`

const DataSourceFormatCsv DataSourceFormat = `CSV`

const DataSourceFormatDelta DataSourceFormat = `DELTA`

const DataSourceFormatDeltasharing DataSourceFormat = `DELTASHARING`

const DataSourceFormatJson DataSourceFormat = `JSON`

const DataSourceFormatOrc DataSourceFormat = `ORC`

const DataSourceFormatParquet DataSourceFormat = `PARQUET`

const DataSourceFormatText DataSourceFormat = `TEXT`

const DataSourceFormatUnityCatalog DataSourceFormat = `UNITY_CATALOG`

// String representation for [fmt.Print]
func (dsf *DataSourceFormat) String() string {
	return string(*dsf)
}

// Set raw string value and validate it against allowed values
func (dsf *DataSourceFormat) Set(v string) error {
	switch v {
	case `AVRO`, `CSV`, `DELTA`, `DELTASHARING`, `JSON`, `ORC`, `PARQUET`, `TEXT`, `UNITY_CATALOG`:
		*dsf = DataSourceFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "CSV", "DELTA", "DELTASHARING", "JSON", "ORC", "PARQUET", "TEXT", "UNITY_CATALOG"`, v)
	}
}

// Type always returns DataSourceFormat to satisfy [pflag.Value] interface
func (dsf *DataSourceFormat) Type() string {
	return "DataSourceFormat"
}

// Delete a catalog
type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is notempty.
	Force bool `json:"-" url:"force,omitempty"`
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
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this External Location was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of External Location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the location's Storage Credential.
	CredentialId string `json:"credential_id,omitempty"`
	// Current name of the Storage Credential this location uses.
	CredentialName string `json:"credential_name,omitempty"`
	// Unique identifier of Metastore hosting the External Location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the External Location.
	Name string `json:"name,omitempty"`
	// The owner of the External Location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which External Location this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the External Location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Path URL of the External Location.
	Url string `json:"url,omitempty"`
}

type GcpServiceAccountKey struct {
	// The email of the service account.
	Email string `json:"email"`
	// The service account's RSA private key.
	PrivateKey string `json:"private_key"`
	// The ID of the service account's private key.
	PrivateKeyId string `json:"private_key_id"`
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

// Get an external location
type GetExternalLocationRequest struct {
	// Required. Name of the storage credential.
	Name string `json:"-" url:"-"`
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

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the Metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this Metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the Metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the Metastore
	DeltaSharingScope GetMetastoreSummaryResponseDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// The unique ID (UUID) of the Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the Metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`)
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region of the Metastore home shard (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for the Metastore.
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which this Metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the External Location.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// The scope of Delta Sharing enabled for the Metastore
type GetMetastoreSummaryResponseDeltaSharingScope string

const GetMetastoreSummaryResponseDeltaSharingScopeInternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL`

const GetMetastoreSummaryResponseDeltaSharingScopeInternalAndExternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (gmsrdss *GetMetastoreSummaryResponseDeltaSharingScope) String() string {
	return string(*gmsrdss)
}

// Set raw string value and validate it against allowed values
func (gmsrdss *GetMetastoreSummaryResponseDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*gmsrdss = GetMetastoreSummaryResponseDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns GetMetastoreSummaryResponseDeltaSharingScope to satisfy [pflag.Value] interface
func (gmsrdss *GetMetastoreSummaryResponseDeltaSharingScope) Type() string {
	return "GetMetastoreSummaryResponseDeltaSharingScope"
}

type GetPermissionsResponse struct {
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

// Get a provider
type GetProviderRequest struct {
	// Required. Name of the provider.
	Name string `json:"-" url:"-"`
}

// Get a share recipient
type GetRecipientRequest struct {
	// Required. Name of the recipient.
	Name string `json:"-" url:"-"`
}

type GetRecipientSharePermissionsResponse struct {
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`
}

// Get a schema
type GetSchemaRequest struct {
	// Required. Full name of the schema (from URL).
	FullName string `json:"-" url:"-"`
}

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

// Get a credential
type GetStorageCredentialRequest struct {
	// Required. Name of the storage credential.
	Name string `json:"-" url:"-"`
}

// Get a table
type GetTableRequest struct {
	// Required. Full name of the Table (from URL).
	FullName string `json:"-" url:"-"`
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	Catalogs []CatalogInfo `json:"catalogs,omitempty"`
}

type ListExternalLocationsResponse struct {
	// An array of external locations.
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
	// Time at which this Metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of (Default) Data Access Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// Whether Delta Sharing is enabled on this metastore.
	DeltaSharingEnabled bool `json:"delta_sharing_enabled,omitempty"`
	// The lifetime of delta sharing recipient token in seconds
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// Unique identifier of Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// The region this metastore has an afinity to. This is used by
	// accounts-manager. Ignored by Unity Catalog.
	Region string `json:"region,omitempty"`
	// Storage root URL for Metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Time at which the Metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the Metastore.
	UpdatedBy string `json:"updated_by,omitempty"`
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
func (pvo *PartitionValueOp) String() string {
	return string(*pvo)
}

// Set raw string value and validate it against allowed values
func (pvo *PartitionValueOp) Set(v string) error {
	switch v {
	case `EQUAL`, `LIKE`:
		*pvo = PartitionValueOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "LIKE"`, v)
	}
}

// Type always returns PartitionValueOp to satisfy [pflag.Value] interface
func (pvo *PartitionValueOp) Type() string {
	return "PartitionValueOp"
}

type PermissionsChange struct {
	// The set of privileges to add.
	Add []Privilege `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	Remove []Privilege `json:"remove,omitempty"`
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
func (p *Privilege) String() string {
	return string(*p)
}

// Set raw string value and validate it against allowed values
func (p *Privilege) Set(v string) error {
	switch v {
	case `ALL_PRIVILEGES`, `CREATE`, `CREATE_CATALOG`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `EXECUTE`, `MODIFY`, `READ_FILES`, `READ_PRIVATE_FILES`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`:
		*p = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_PRIVILEGES", "CREATE", "CREATE_CATALOG", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "EXECUTE", "MODIFY", "READ_FILES", "READ_PRIVATE_FILES", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES"`, v)
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (p *Privilege) Type() string {
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
	// Cloud vendor of the provider's UC Metastore. This field is only present
	// when the authentication_type is `DATABRICKS`.
	Cloud string `json:"cloud,omitempty"`
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The global UC metastore id of the data provider. This field is only
	// present when the authentication type is `DATABRICKS`. The identifier is
	// of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId string `json:"data_provider_global_metastore_id,omitempty"`
	// UUID of the provider's UC Metastore. This field is only present when the
	// authentication type is `DATABRICKS`.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the Provider.
	Name string `json:"name,omitempty"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// This field is required when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// Cloud region of the provider's UC Metastore. This field is only present
	// when the authentication type is `DATABRICKS`.
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
	// only present when the authentication type is `DATABRICKS`.
	Cloud string `json:"cloud,omitempty"`
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.\n
	// This field is only present when the authentication type is
	// `DATABRICKS`.\n The identifier is of format
	// <cloud>:<region>:<metastore-uuid>.
	DataRecipientGlobalMetastoreId any `json:"data_recipient_global_metastore_id,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Unique identifier of recipient's Unity Catalog Metastore. This field is
	// only present when the authentication type is `DATABRICKS`
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Recipient.
	Name string `json:"name,omitempty"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
	// only present when the authentication type is `DATABRICKS`.
	Region string `json:"region,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the authentication type is `DATABRICKS`.
	SharingCode string `json:"sharing_code,omitempty"`
	// This field is only present when the authentication type is `TOKEN`.
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

type SchemaInfo struct {
	// Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this Schema was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Schema creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Full name of Schema, in form of <catalog_name>.<schema_name>.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`

	Properties map[string]string `json:"properties,omitempty"`
	// Storage location for managed tables within schema.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this Schema was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified Schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type ShareInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this Share was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Share creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Name of the Share.
	Name string `json:"name,omitempty"`
	// A list of shared data objects within the Share.
	Objects []SharedDataObject `json:"objects,omitempty"`
	// Username of current owner of Share.
	Owner string `json:"owner,omitempty"`
	// Array of shared data object updates.
	Updates []SharedDataObjectUpdate `json:"updates,omitempty"`
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
	// new name is not not provided, the object's original name will be used as
	// the `shared_as` name. The `shared_as` name must be unique within a Share.
	// For tables, the new name must follow the format of `<schema>.<table>`.
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
func (sdos *SharedDataObjectStatus) String() string {
	return string(*sdos)
}

// Set raw string value and validate it against allowed values
func (sdos *SharedDataObjectStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `PERMISSION_DENIED`:
		*sdos = SharedDataObjectStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "PERMISSION_DENIED"`, v)
	}
}

// Type always returns SharedDataObjectStatus to satisfy [pflag.Value] interface
func (sdos *SharedDataObjectStatus) Type() string {
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
func (sdoua *SharedDataObjectUpdateAction) String() string {
	return string(*sdoua)
}

// Set raw string value and validate it against allowed values
func (sdoua *SharedDataObjectUpdateAction) Set(v string) error {
	switch v {
	case `ADD`, `REMOVE`, `UPDATE`:
		*sdoua = SharedDataObjectUpdateAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADD", "REMOVE", "UPDATE"`, v)
	}
}

// Type always returns SharedDataObjectUpdateAction to satisfy [pflag.Value] interface
func (sdoua *SharedDataObjectUpdateAction) Type() string {
	return "SharedDataObjectUpdateAction"
}

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name MUST be unique within the Metastore.
	Name string `json:"name,omitempty"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
}

type TableInfo struct {
	// Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// This name ('columns') is what the client actually sees as the field name
	// in messages that include PropertiesKVPairs using 'json_inline' (e.g.,
	// TableInfo).
	Columns []ColumnInfo `json:"columns,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this Table was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the data_access_configuration to use.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// Data source format
	DataSourceFormat DataSourceFormat `json:"data_source_format,omitempty"`
	// Full name of Table, in form of <catalog_name>.<schema_name>.<table_name>
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Table, relative to parent Schema.
	Name string `json:"name,omitempty"`
	// Username of current owner of Table.
	Owner string `json:"owner,omitempty"`

	Properties map[string]string `json:"properties,omitempty"`
	// Name of parent Schema relative to its parent Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Name of the storage credential this table used
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// Storage root URL for table (for MANAGED, EXTERNAL tables)
	StorageLocation string `json:"storage_location,omitempty"`
	// Name of Table, relative to parent Schema.
	TableId string `json:"table_id,omitempty"`

	TableType TableType `json:"table_type,omitempty"`
	// Time at which this Table was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the Table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// View definition SQL (when table_type == "VIEW")
	ViewDefinition string `json:"view_definition,omitempty"`
}

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

	TableType TableType `json:"table_type,omitempty"`
}

type TableType string

const TableTypeExternal TableType = `EXTERNAL`

const TableTypeManaged TableType = `MANAGED`

const TableTypeMaterializedView TableType = `MATERIALIZED_VIEW`

const TableTypeStreamingTable TableType = `STREAMING_TABLE`

const TableTypeView TableType = `VIEW`

// String representation for [fmt.Print]
func (tt *TableType) String() string {
	return string(*tt)
}

// Set raw string value and validate it against allowed values
func (tt *TableType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `MANAGED`, `MATERIALIZED_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*tt = TableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "MANAGED", "MATERIALIZED_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Type always returns TableType to satisfy [pflag.Value] interface
func (tt *TableType) Type() string {
	return "TableType"
}

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the Metastore to delete.
	MetastoreId string `json:"-" url:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of Catalog.
	Name string `json:"name,omitempty" url:"-"`
	// Username of current owner of Catalog.
	Owner string `json:"owner,omitempty"`

	Properties map[string]string `json:"properties,omitempty"`
}

type UpdateExternalLocation struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Current name of the Storage Credential this location uses.
	CredentialName string `json:"credential_name,omitempty"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force bool `json:"force,omitempty"`
	// Name of the External Location.
	Name string `json:"name,omitempty" url:"-"`
	// The owner of the External Location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the External Location.
	Url string `json:"url,omitempty"`
}

type UpdateMetastore struct {
	// Unique identifier of (Default) Data Access Configuration
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// Whether Delta Sharing is enabled on this metastore.
	DeltaSharingEnabled bool `json:"delta_sharing_enabled,omitempty"`
	// The lifetime of delta sharing recipient token in seconds
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// Required. Unique ID of the Metastore (from URL).
	Id string `json:"-" url:"-"`
	// Name of Metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// UUID of storage credential to access storage_root
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog for the Metastore.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the Metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

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
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// The name of the Provider.
	Name string `json:"name,omitempty" url:"-"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// This field is required when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
}

type UpdateRecipient struct {
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// IP Access List
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Name of Recipient.
	Name string `json:"name,omitempty" url:"-"`
}

type UpdateSchema struct {
	// Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Required. Full name of the schema (from URL).
	FullName string `json:"-" url:"-"`
	// Name of Schema, relative to parent Catalog.
	Name string `json:"name,omitempty"`
	// Username of current owner of Schema.
	Owner string `json:"owner,omitempty"`

	Properties map[string]string `json:"properties,omitempty"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `json:"storage_root,omitempty"`
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the Share.
	Name string `json:"name,omitempty" url:"-"`
	// Username of current owner of Share.
	Owner string `json:"owner,omitempty"`
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
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// The credential name. The name MUST be unique within the Metastore.
	Name string `json:"name,omitempty" url:"-"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
}
