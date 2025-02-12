// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccountsCreateMetastore struct {
	MetastoreInfo *CreateMetastore `json:"metastore_info,omitempty"`
}

type AccountsCreateMetastoreAssignment struct {
	MetastoreAssignment *CreateMetastoreAssignment `json:"metastore_assignment,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type AccountsCreateStorageCredential struct {
	CredentialInfo *CreateStorageCredential `json:"credential_info,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

type AccountsMetastoreAssignment struct {
	MetastoreAssignment *MetastoreAssignment `json:"metastore_assignment,omitempty"`
}

type AccountsMetastoreInfo struct {
	MetastoreInfo *MetastoreInfo `json:"metastore_info,omitempty"`
}

type AccountsStorageCredentialInfo struct {
	CredentialInfo *StorageCredentialInfo `json:"credential_info,omitempty"`
}

type AccountsUpdateMetastore struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`

	MetastoreInfo *UpdateMetastore `json:"metastore_info,omitempty"`
}

type AccountsUpdateMetastoreAssignment struct {
	MetastoreAssignment *UpdateMetastoreAssignment `json:"metastore_assignment,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type AccountsUpdateStorageCredential struct {
	CredentialInfo *UpdateStorageCredential `json:"credential_info,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`
}

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher `json:"artifact_matchers,omitempty"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of the user who set the artifact allowlist.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ArtifactAllowlistInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ArtifactAllowlistInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	Artifact string `json:"artifact"`
	// The pattern matching type of the artifact
	MatchType MatchType `json:"match_type"`
}

// The artifact type
type ArtifactType string

const ArtifactTypeInitScript ArtifactType = `INIT_SCRIPT`

const ArtifactTypeLibraryJar ArtifactType = `LIBRARY_JAR`

const ArtifactTypeLibraryMaven ArtifactType = `LIBRARY_MAVEN`

// String representation for [fmt.Print]
func (f *ArtifactType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ArtifactType) Set(v string) error {
	switch v {
	case `INIT_SCRIPT`, `LIBRARY_JAR`, `LIBRARY_MAVEN`:
		*f = ArtifactType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INIT_SCRIPT", "LIBRARY_JAR", "LIBRARY_MAVEN"`, v)
	}
}

// Type always returns ArtifactType to satisfy [pflag.Value] interface
func (f *ArtifactType) Type() string {
	return "ArtifactType"
}

type AssignResponse struct {
}

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId string `json:"access_key_id,omitempty"`
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	AccessPoint string `json:"access_point,omitempty"`
	// The secret access key that can be used to sign AWS API requests.
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	SessionToken string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsCredentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The AWS IAM role configuration
type AwsIamRole struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn string `json:"role_arn,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsIamRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsIamRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AwsIamRoleRequest struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string `json:"role_arn"`
}

type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string `json:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsIamRoleResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsIamRoleResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Azure Active Directory token, essentially the Oauth token for Azure Service
// Principal or Managed Identity. Read more at
// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
type AzureActiveDirectoryToken struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	AadToken string `json:"aad_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureActiveDirectoryToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureActiveDirectoryToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Azure managed identity configuration.
type AzureManagedIdentity struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	AccessConnectorId string `json:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database. .
	CredentialId string `json:"credential_id,omitempty"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureManagedIdentity) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentity) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureManagedIdentityRequest struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId string `json:"access_connector_id"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureManagedIdentityRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentityRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureManagedIdentityResponse struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId string `json:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity.
	CredentialId string `json:"credential_id,omitempty"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureManagedIdentityResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentityResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Azure service principal configuration. Only applicable when purpose is
// **STORAGE**.
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

// Azure temporary credentials for API authentication. Read more at
// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
type AzureUserDelegationSas struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	SasToken string `json:"sas_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureUserDelegationSas) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureUserDelegationSas) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Cancel refresh
type CancelRefreshRequest struct {
	// ID of the refresh.
	RefreshId string `json:"-" url:"-"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

type CancelRefreshResponse struct {
}

type CatalogInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The type of the catalog.
	CatalogType CatalogType `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The name of the connection to an external data source.
	ConnectionName string `json:"connection_name,omitempty"`
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of catalog creator.
	CreatedBy string `json:"created_by,omitempty"`

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// The full name of the catalog. Corresponds with the name field.
	FullName string `json:"full_name,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode CatalogIsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of catalog.
	Name string `json:"name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of catalog.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `json:"provider_name,omitempty"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo *ProvisioningInfo `json:"provisioning_info,omitempty"`

	SecurableType string `json:"securable_type,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// Storage Location URL (full path) for managed tables within catalog.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within catalog.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified catalog.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CatalogInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Whether the current securable is accessible from all workspaces or a specific
// set of workspaces.
type CatalogIsolationMode string

const CatalogIsolationModeIsolated CatalogIsolationMode = `ISOLATED`

const CatalogIsolationModeOpen CatalogIsolationMode = `OPEN`

// String representation for [fmt.Print]
func (f *CatalogIsolationMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogIsolationMode) Set(v string) error {
	switch v {
	case `ISOLATED`, `OPEN`:
		*f = CatalogIsolationMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ISOLATED", "OPEN"`, v)
	}
}

// Type always returns CatalogIsolationMode to satisfy [pflag.Value] interface
func (f *CatalogIsolationMode) Type() string {
	return "CatalogIsolationMode"
}

// The type of the catalog.
type CatalogType string

const CatalogTypeDeltasharingCatalog CatalogType = `DELTASHARING_CATALOG`

const CatalogTypeManagedCatalog CatalogType = `MANAGED_CATALOG`

const CatalogTypeSystemCatalog CatalogType = `SYSTEM_CATALOG`

// String representation for [fmt.Print]
func (f *CatalogType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogType) Set(v string) error {
	switch v {
	case `DELTASHARING_CATALOG`, `MANAGED_CATALOG`, `SYSTEM_CATALOG`:
		*f = CatalogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTASHARING_CATALOG", "MANAGED_CATALOG", "SYSTEM_CATALOG"`, v)
	}
}

// Type always returns CatalogType to satisfy [pflag.Value] interface
func (f *CatalogType) Type() string {
	return "CatalogType"
}

type CloudflareApiToken struct {
	// The Cloudflare access key id of the token.
	AccessKeyId string `json:"access_key_id"`
	// The account id associated with the API token.
	AccountId string `json:"account_id"`
	// The secret access token generated for the access key id
	SecretAccessKey string `json:"secret_access_key"`
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`

	Mask *ColumnMask `json:"mask,omitempty"`
	// Name of Column.
	Name string `json:"name,omitempty"`
	// Whether field may be Null (default: true).
	Nullable bool `json:"nullable,omitempty"`
	// Partition index for column.
	PartitionIndex int `json:"partition_index,omitempty"`
	// Ordinal position of column (starting at position 0).
	Position int `json:"position,omitempty"`
	// Format of IntervalType.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type specification, JSON-serialized.
	TypeJson string `json:"type_json,omitempty"`

	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type specification as SQL/catalogString text.
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	FunctionName string `json:"function_name,omitempty"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames []string `json:"using_column_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnMask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnMask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ColumnTypeName string

const ColumnTypeNameArray ColumnTypeName = `ARRAY`

const ColumnTypeNameBinary ColumnTypeName = `BINARY`

const ColumnTypeNameBoolean ColumnTypeName = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeName = `BYTE`

const ColumnTypeNameChar ColumnTypeName = `CHAR`

const ColumnTypeNameDate ColumnTypeName = `DATE`

const ColumnTypeNameDecimal ColumnTypeName = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeName = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeName = `FLOAT`

const ColumnTypeNameInt ColumnTypeName = `INT`

const ColumnTypeNameInterval ColumnTypeName = `INTERVAL`

const ColumnTypeNameLong ColumnTypeName = `LONG`

const ColumnTypeNameMap ColumnTypeName = `MAP`

const ColumnTypeNameNull ColumnTypeName = `NULL`

const ColumnTypeNameShort ColumnTypeName = `SHORT`

const ColumnTypeNameString ColumnTypeName = `STRING`

const ColumnTypeNameStruct ColumnTypeName = `STRUCT`

const ColumnTypeNameTableType ColumnTypeName = `TABLE_TYPE`

const ColumnTypeNameTimestamp ColumnTypeName = `TIMESTAMP`

const ColumnTypeNameTimestampNtz ColumnTypeName = `TIMESTAMP_NTZ`

const ColumnTypeNameUserDefinedType ColumnTypeName = `USER_DEFINED_TYPE`

const ColumnTypeNameVariant ColumnTypeName = `VARIANT`

// String representation for [fmt.Print]
func (f *ColumnTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`, `VARIANT`:
		*f = ColumnTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TABLE_TYPE", "TIMESTAMP", "TIMESTAMP_NTZ", "USER_DEFINED_TYPE", "VARIANT"`, v)
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

type ConnectionInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Unique identifier of the Connection.
	ConnectionId string `json:"connection_id,omitempty"`
	// The type of connection.
	ConnectionType ConnectionType `json:"connection_type,omitempty"`
	// Time at which this connection was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of connection creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The type of credential.
	CredentialType CredentialType `json:"credential_type,omitempty"`
	// Full name of connection.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the connection.
	Name string `json:"name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of the connection.
	Owner string `json:"owner,omitempty"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string `json:"properties,omitempty"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo *ProvisioningInfo `json:"provisioning_info,omitempty"`
	// If the connection is read only.
	ReadOnly bool `json:"read_only,omitempty"`

	SecurableType string `json:"securable_type,omitempty"`
	// Time at which this connection was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified connection.
	UpdatedBy string `json:"updated_by,omitempty"`
	// URL of the remote data source, extracted from options.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ConnectionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConnectionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of connection.
type ConnectionType string

const ConnectionTypeBigquery ConnectionType = `BIGQUERY`

const ConnectionTypeDatabricks ConnectionType = `DATABRICKS`

const ConnectionTypeGlue ConnectionType = `GLUE`

const ConnectionTypeHiveMetastore ConnectionType = `HIVE_METASTORE`

const ConnectionTypeHttp ConnectionType = `HTTP`

const ConnectionTypeMysql ConnectionType = `MYSQL`

const ConnectionTypePostgresql ConnectionType = `POSTGRESQL`

const ConnectionTypeRedshift ConnectionType = `REDSHIFT`

const ConnectionTypeSnowflake ConnectionType = `SNOWFLAKE`

const ConnectionTypeSqldw ConnectionType = `SQLDW`

const ConnectionTypeSqlserver ConnectionType = `SQLSERVER`

// String representation for [fmt.Print]
func (f *ConnectionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConnectionType) Set(v string) error {
	switch v {
	case `BIGQUERY`, `DATABRICKS`, `GLUE`, `HIVE_METASTORE`, `HTTP`, `MYSQL`, `POSTGRESQL`, `REDSHIFT`, `SNOWFLAKE`, `SQLDW`, `SQLSERVER`:
		*f = ConnectionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BIGQUERY", "DATABRICKS", "GLUE", "HIVE_METASTORE", "HTTP", "MYSQL", "POSTGRESQL", "REDSHIFT", "SNOWFLAKE", "SQLDW", "SQLSERVER"`, v)
	}
}

// Type always returns ConnectionType to satisfy [pflag.Value] interface
func (f *ConnectionType) Type() string {
	return "ConnectionType"
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
type ContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress *PipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The name of the connection to an external data source.
	ConnectionName string `json:"connection_name,omitempty"`
	// Name of catalog.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `json:"provider_name,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// Storage root URL for managed tables within catalog.
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateConnection struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The type of connection.
	ConnectionType ConnectionType `json:"connection_type"`
	// Name of the connection.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string `json:"properties,omitempty"`
	// If the connection is read only.
	ReadOnly bool `json:"read_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name"`
	// Indicates the purpose of the credential.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateExternalLocation struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`
	// Name of the external location.
	Name string `json:"name"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the external location.
	Url string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExternalLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExternalLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFunction struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Scalar function return data type.
	DataType ColumnTypeName `json:"data_type"`
	// External function language.
	ExternalLanguage string `json:"external_language,omitempty"`
	// External function name.
	ExternalName string `json:"external_name,omitempty"`
	// Pretty printed function data type.
	FullDataType string `json:"full_data_type"`

	InputParams FunctionParameterInfos `json:"input_params"`
	// Whether the function is deterministic.
	IsDeterministic bool `json:"is_deterministic"`
	// Function null call.
	IsNullCall bool `json:"is_null_call"`
	// Name of function, relative to parent schema.
	Name string `json:"name"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle CreateFunctionParameterStyle `json:"parameter_style"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string `json:"properties,omitempty"`
	// Table function return parameters.
	ReturnParams *FunctionParameterInfos `json:"return_params,omitempty"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody CreateFunctionRoutineBody `json:"routine_body"`
	// Function body.
	RoutineDefinition string `json:"routine_definition"`
	// Function dependencies.
	RoutineDependencies *DependencyList `json:"routine_dependencies,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name"`
	// Function security type.
	SecurityType CreateFunctionSecurityType `json:"security_type"`
	// Specific name of the function; Reserved for future use.
	SpecificName string `json:"specific_name"`
	// Function SQL data access.
	SqlDataAccess CreateFunctionSqlDataAccess `json:"sql_data_access"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Function parameter style. **S** is the value for SQL.
type CreateFunctionParameterStyle string

const CreateFunctionParameterStyleS CreateFunctionParameterStyle = `S`

// String representation for [fmt.Print]
func (f *CreateFunctionParameterStyle) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionParameterStyle) Set(v string) error {
	switch v {
	case `S`:
		*f = CreateFunctionParameterStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "S"`, v)
	}
}

// Type always returns CreateFunctionParameterStyle to satisfy [pflag.Value] interface
func (f *CreateFunctionParameterStyle) Type() string {
	return "CreateFunctionParameterStyle"
}

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	FunctionInfo CreateFunction `json:"function_info"`
}

// Function language. When **EXTERNAL** is used, the language of the routine
// function should be specified in the __external_language__ field, and the
// __return_params__ of the function cannot be used (as **TABLE** return type is
// not supported), and the __sql_data_access__ field must be **NO_SQL**.
type CreateFunctionRoutineBody string

const CreateFunctionRoutineBodyExternal CreateFunctionRoutineBody = `EXTERNAL`

const CreateFunctionRoutineBodySql CreateFunctionRoutineBody = `SQL`

// String representation for [fmt.Print]
func (f *CreateFunctionRoutineBody) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionRoutineBody) Set(v string) error {
	switch v {
	case `EXTERNAL`, `SQL`:
		*f = CreateFunctionRoutineBody(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "SQL"`, v)
	}
}

// Type always returns CreateFunctionRoutineBody to satisfy [pflag.Value] interface
func (f *CreateFunctionRoutineBody) Type() string {
	return "CreateFunctionRoutineBody"
}

// The security type of the function.
type CreateFunctionSecurityType string

const CreateFunctionSecurityTypeDefiner CreateFunctionSecurityType = `DEFINER`

// String representation for [fmt.Print]
func (f *CreateFunctionSecurityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionSecurityType) Set(v string) error {
	switch v {
	case `DEFINER`:
		*f = CreateFunctionSecurityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFINER"`, v)
	}
}

// Type always returns CreateFunctionSecurityType to satisfy [pflag.Value] interface
func (f *CreateFunctionSecurityType) Type() string {
	return "CreateFunctionSecurityType"
}

// Function SQL data access.
type CreateFunctionSqlDataAccess string

const CreateFunctionSqlDataAccessContainsSql CreateFunctionSqlDataAccess = `CONTAINS_SQL`

const CreateFunctionSqlDataAccessNoSql CreateFunctionSqlDataAccess = `NO_SQL`

const CreateFunctionSqlDataAccessReadsSqlData CreateFunctionSqlDataAccess = `READS_SQL_DATA`

// String representation for [fmt.Print]
func (f *CreateFunctionSqlDataAccess) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionSqlDataAccess) Set(v string) error {
	switch v {
	case `CONTAINS_SQL`, `NO_SQL`, `READS_SQL_DATA`:
		*f = CreateFunctionSqlDataAccess(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTAINS_SQL", "NO_SQL", "READS_SQL_DATA"`, v)
	}
}

// Type always returns CreateFunctionSqlDataAccess to satisfy [pflag.Value] interface
func (f *CreateFunctionSqlDataAccess) Type() string {
	return "CreateFunctionSqlDataAccess"
}

type CreateMetastore struct {
	// The user-specified name of the metastore.
	Name string `json:"name"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// The field can be omitted in the __workspace-level__ __API__ but not in
	// the __account-level__ __API__. If this field is omitted, the region of
	// the workspace receiving the request will be used.
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName string `json:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type CreateMonitor struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir string `json:"assets_dir"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// The notification settings for the monitor.
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	OutputSchemaName string `json:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard bool `json:"skip_builtin_dashboard,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateMonitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMonitor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Create an Online Table
type CreateOnlineTableRequest struct {
	// Online Table information.
	Table *OnlineTable `json:"table,omitempty"`
}

type CreateRegisteredModelRequest struct {
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `json:"catalog_name"`
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// The name of the registered model
	Name string `json:"name"`
	// The name of the schema where the registered model resides
	SchemaName string `json:"schema_name"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateResponse struct {
}

type CreateSchema struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique within the metastore.
	Name string `json:"name"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTableConstraint struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	Constraint TableConstraint `json:"constraint"`
	// The full name of the table referenced by the constraint.
	FullNameArg string `json:"full_name_arg"`
}

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The name of the volume
	Name string `json:"name"`
	// The name of the schema where the volume is
	SchemaName string `json:"schema_name"`
	// The storage location on the cloud
	StorageLocation string `json:"storage_location,omitempty"`

	VolumeType VolumeType `json:"volume_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CredentialInfo struct {
	// The AWS IAM role configuration
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Time at which this credential was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// The full name of the credential.
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of the parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Indicates the purpose of the credential.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CredentialPurpose string

const CredentialPurposeService CredentialPurpose = `SERVICE`

const CredentialPurposeStorage CredentialPurpose = `STORAGE`

// String representation for [fmt.Print]
func (f *CredentialPurpose) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CredentialPurpose) Set(v string) error {
	switch v {
	case `SERVICE`, `STORAGE`:
		*f = CredentialPurpose(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SERVICE", "STORAGE"`, v)
	}
}

// Type always returns CredentialPurpose to satisfy [pflag.Value] interface
func (f *CredentialPurpose) Type() string {
	return "CredentialPurpose"
}

// The type of credential.
type CredentialType string

const CredentialTypeBearerToken CredentialType = `BEARER_TOKEN`

const CredentialTypeUsernamePassword CredentialType = `USERNAME_PASSWORD`

// String representation for [fmt.Print]
func (f *CredentialType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CredentialType) Set(v string) error {
	switch v {
	case `BEARER_TOKEN`, `USERNAME_PASSWORD`:
		*f = CredentialType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BEARER_TOKEN", "USERNAME_PASSWORD"`, v)
	}
}

// Type always returns CredentialType to satisfy [pflag.Value] interface
func (f *CredentialType) Type() string {
	return "CredentialType"
}

type CredentialValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message string `json:"message,omitempty"`
	// The results of the tested operation.
	Result ValidateCredentialResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CredentialValidationResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialValidationResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Currently assigned workspaces
type CurrentWorkspaceBindings struct {
	// A list of workspace IDs.
	Workspaces []int64 `json:"workspaces,omitempty"`
}

// Data source format
type DataSourceFormat string

const DataSourceFormatAvro DataSourceFormat = `AVRO`

const DataSourceFormatBigqueryFormat DataSourceFormat = `BIGQUERY_FORMAT`

const DataSourceFormatCsv DataSourceFormat = `CSV`

const DataSourceFormatDatabricksFormat DataSourceFormat = `DATABRICKS_FORMAT`

const DataSourceFormatDelta DataSourceFormat = `DELTA`

const DataSourceFormatDeltasharing DataSourceFormat = `DELTASHARING`

const DataSourceFormatHiveCustom DataSourceFormat = `HIVE_CUSTOM`

const DataSourceFormatHiveSerde DataSourceFormat = `HIVE_SERDE`

const DataSourceFormatJson DataSourceFormat = `JSON`

const DataSourceFormatMysqlFormat DataSourceFormat = `MYSQL_FORMAT`

const DataSourceFormatNetsuiteFormat DataSourceFormat = `NETSUITE_FORMAT`

const DataSourceFormatOrc DataSourceFormat = `ORC`

const DataSourceFormatParquet DataSourceFormat = `PARQUET`

const DataSourceFormatPostgresqlFormat DataSourceFormat = `POSTGRESQL_FORMAT`

const DataSourceFormatRedshiftFormat DataSourceFormat = `REDSHIFT_FORMAT`

const DataSourceFormatSalesforceFormat DataSourceFormat = `SALESFORCE_FORMAT`

const DataSourceFormatSnowflakeFormat DataSourceFormat = `SNOWFLAKE_FORMAT`

const DataSourceFormatSqldwFormat DataSourceFormat = `SQLDW_FORMAT`

const DataSourceFormatSqlserverFormat DataSourceFormat = `SQLSERVER_FORMAT`

const DataSourceFormatText DataSourceFormat = `TEXT`

const DataSourceFormatUnityCatalog DataSourceFormat = `UNITY_CATALOG`

const DataSourceFormatVectorIndexFormat DataSourceFormat = `VECTOR_INDEX_FORMAT`

const DataSourceFormatWorkdayRaasFormat DataSourceFormat = `WORKDAY_RAAS_FORMAT`

// String representation for [fmt.Print]
func (f *DataSourceFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataSourceFormat) Set(v string) error {
	switch v {
	case `AVRO`, `BIGQUERY_FORMAT`, `CSV`, `DATABRICKS_FORMAT`, `DELTA`, `DELTASHARING`, `HIVE_CUSTOM`, `HIVE_SERDE`, `JSON`, `MYSQL_FORMAT`, `NETSUITE_FORMAT`, `ORC`, `PARQUET`, `POSTGRESQL_FORMAT`, `REDSHIFT_FORMAT`, `SALESFORCE_FORMAT`, `SNOWFLAKE_FORMAT`, `SQLDW_FORMAT`, `SQLSERVER_FORMAT`, `TEXT`, `UNITY_CATALOG`, `VECTOR_INDEX_FORMAT`, `WORKDAY_RAAS_FORMAT`:
		*f = DataSourceFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "BIGQUERY_FORMAT", "CSV", "DATABRICKS_FORMAT", "DELTA", "DELTASHARING", "HIVE_CUSTOM", "HIVE_SERDE", "JSON", "MYSQL_FORMAT", "NETSUITE_FORMAT", "ORC", "PARQUET", "POSTGRESQL_FORMAT", "REDSHIFT_FORMAT", "SALESFORCE_FORMAT", "SNOWFLAKE_FORMAT", "SQLDW_FORMAT", "SQLSERVER_FORMAT", "TEXT", "UNITY_CATALOG", "VECTOR_INDEX_FORMAT", "WORKDAY_RAAS_FORMAT"`, v)
	}
}

// Type always returns DataSourceFormat to satisfy [pflag.Value] interface
func (f *DataSourceFormat) Type() string {
	return "DataSourceFormat"
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccount struct {
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database
	CredentialId string `json:"credential_id,omitempty"`
	// The email of the service account.
	Email string `json:"email,omitempty"`
	// The ID that represents the private key for this Service Account
	PrivateKeyId string `json:"private_key_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksGcpServiceAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksGcpServiceAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabricksGcpServiceAccountRequest struct {
}

type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this service account. This is
	// an output-only field.
	CredentialId string `json:"credential_id,omitempty"`
	// The email of the service account. This is an output-only field.
	Email string `json:"email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksGcpServiceAccountResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksGcpServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a metastore assignment
type DeleteAccountMetastoreAssignmentRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

// Delete a metastore
type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAccountMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountMetastoreRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a storage credential
type DeleteAccountStorageCredentialRequest struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAccountStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a Registered Model Alias
type DeleteAliasRequest struct {
	// The name of the alias
	Alias string `json:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
}

type DeleteAliasResponse struct {
}

// Delete a catalog
type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	Force bool `json:"-" url:"force,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCatalogRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a connection
type DeleteConnectionRequest struct {
	// The name of the connection to be deleted.
	Name string `json:"-" url:"-"`
}

// Delete a credential
type DeleteCredentialRequest struct {
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the credential.
	NameArg string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteCredentialResponse struct {
}

// Delete an external location
type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteExternalLocationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteExternalLocationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a function
type DeleteFunctionRequest struct {
	// Force deletion even if the function is notempty.
	Force bool `json:"-" url:"force,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a metastore
type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteMetastoreRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a Model Version
type DeleteModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" url:"-"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`
}

// Delete an Online Table
type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" url:"-"`
}

// Delete a table monitor
type DeleteQualityMonitorRequest struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

// Delete a Registered Model
type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete a schema
type DeleteSchemaRequest struct {
	// Force deletion even if the schema is not empty.
	Force bool `json:"-" url:"force,omitempty"`
	// Full name of the schema.
	FullName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteSchemaRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSchemaRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a credential
type DeleteStorageCredentialRequest struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the storage credential.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a table constraint
type DeleteTableConstraintRequest struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	Cascade bool `json:"-" url:"cascade"`
	// The name of the constraint to delete.
	ConstraintName string `json:"-" url:"constraint_name"`
	// Full name of the table referenced by the constraint.
	FullName string `json:"-" url:"-"`
}

// Delete a table
type DeleteTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
}

// Delete a Volume
type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" url:"-"`
}

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	DeltaRuntimeProperties map[string]string `json:"delta_runtime_properties"`
}

// A dependency of a SQL object. Either the __table__ field or the __function__
// field must be defined.
type Dependency struct {
	// A function that is dependent on a SQL object.
	Function *FunctionDependency `json:"function,omitempty"`
	// A table that is dependent on a SQL object.
	Table *TableDependency `json:"table,omitempty"`
}

// A list of dependencies.
type DependencyList struct {
	// Array of dependencies.
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

// Disable a system schema
type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string `json:"-" url:"-"`
	// Full name of the system schema.
	SchemaName string `json:"-" url:"-"`
}

type DisableResponse struct {
}

type EffectivePermissionsList struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments []EffectivePrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType `json:"inherited_from_type,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	Value EnablePredictiveOptimization `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EffectivePredictiveOptimizationFlag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePredictiveOptimizationFlag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of the object from which the flag was inherited. If there was no
// inheritance, this field is left blank.
type EffectivePredictiveOptimizationFlagInheritedFromType string

const EffectivePredictiveOptimizationFlagInheritedFromTypeCatalog EffectivePredictiveOptimizationFlagInheritedFromType = `CATALOG`

const EffectivePredictiveOptimizationFlagInheritedFromTypeSchema EffectivePredictiveOptimizationFlagInheritedFromType = `SCHEMA`

// String representation for [fmt.Print]
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Set(v string) error {
	switch v {
	case `CATALOG`, `SCHEMA`:
		*f = EffectivePredictiveOptimizationFlagInheritedFromType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG", "SCHEMA"`, v)
	}
}

// Type always returns EffectivePredictiveOptimizationFlagInheritedFromType to satisfy [pflag.Value] interface
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Type() string {
	return "EffectivePredictiveOptimizationFlagInheritedFromType"
}

type EffectivePrivilege struct {
	// The full name of the object that conveys this privilege via inheritance.
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType SecurableType `json:"inherited_from_type,omitempty"`
	// The privilege assigned to the principal.
	Privilege Privilege `json:"privilege,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EffectivePrivilege) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePrivilege) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges []EffectivePrivilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EffectivePrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Whether predictive optimization should be enabled for this object and objects
// under it.
type EnablePredictiveOptimization string

const EnablePredictiveOptimizationDisable EnablePredictiveOptimization = `DISABLE`

const EnablePredictiveOptimizationEnable EnablePredictiveOptimization = `ENABLE`

const EnablePredictiveOptimizationInherit EnablePredictiveOptimization = `INHERIT`

// String representation for [fmt.Print]
func (f *EnablePredictiveOptimization) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EnablePredictiveOptimization) Set(v string) error {
	switch v {
	case `DISABLE`, `ENABLE`, `INHERIT`:
		*f = EnablePredictiveOptimization(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLE", "ENABLE", "INHERIT"`, v)
	}
}

// Type always returns EnablePredictiveOptimization to satisfy [pflag.Value] interface
func (f *EnablePredictiveOptimization) Type() string {
	return "EnablePredictiveOptimization"
}

// Enable a system schema
type EnableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string `json:"-" url:"-"`
	// Full name of the system schema.
	SchemaName string `json:"-" url:"-"`
}

type EnableResponse struct {
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	SseEncryptionDetails *SseEncryptionDetails `json:"sse_encryption_details,omitempty"`
}

// Get boolean reflecting if table exists
type ExistsRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
}

type ExternalLocationInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of external location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the location's storage credential.
	CredentialId string `json:"credential_id,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`

	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of metastore hosting the external location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the external location.
	Name string `json:"name,omitempty"`
	// The owner of the external location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the external location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Path URL of the external location.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLocationInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLocationInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FailedStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailedStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	Name string `json:"name"`
	// Column names for this constraint.
	ParentColumns []string `json:"parent_columns"`
	// The full name of the parent constraint.
	ParentTable string `json:"parent_table"`
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName string `json:"function_full_name"`
}

type FunctionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of function creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Scalar function return data type.
	DataType ColumnTypeName `json:"data_type,omitempty"`
	// External function language.
	ExternalLanguage string `json:"external_language,omitempty"`
	// External function name.
	ExternalName string `json:"external_name,omitempty"`
	// Pretty printed function data type.
	FullDataType string `json:"full_data_type,omitempty"`
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	FullName string `json:"full_name,omitempty"`
	// Id of Function, relative to parent schema.
	FunctionId string `json:"function_id,omitempty"`

	InputParams *FunctionParameterInfos `json:"input_params,omitempty"`
	// Whether the function is deterministic.
	IsDeterministic bool `json:"is_deterministic,omitempty"`
	// Function null call.
	IsNullCall bool `json:"is_null_call,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of function, relative to parent schema.
	Name string `json:"name,omitempty"`
	// Username of current owner of function.
	Owner string `json:"owner,omitempty"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle FunctionInfoParameterStyle `json:"parameter_style,omitempty"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string `json:"properties,omitempty"`
	// Table function return parameters.
	ReturnParams *FunctionParameterInfos `json:"return_params,omitempty"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody FunctionInfoRoutineBody `json:"routine_body,omitempty"`
	// Function body.
	RoutineDefinition string `json:"routine_definition,omitempty"`
	// Function dependencies.
	RoutineDependencies *DependencyList `json:"routine_dependencies,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// Function security type.
	SecurityType FunctionInfoSecurityType `json:"security_type,omitempty"`
	// Specific name of the function; Reserved for future use.
	SpecificName string `json:"specific_name,omitempty"`
	// Function SQL data access.
	SqlDataAccess FunctionInfoSqlDataAccess `json:"sql_data_access,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified function.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FunctionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Function parameter style. **S** is the value for SQL.
type FunctionInfoParameterStyle string

const FunctionInfoParameterStyleS FunctionInfoParameterStyle = `S`

// String representation for [fmt.Print]
func (f *FunctionInfoParameterStyle) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoParameterStyle) Set(v string) error {
	switch v {
	case `S`:
		*f = FunctionInfoParameterStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "S"`, v)
	}
}

// Type always returns FunctionInfoParameterStyle to satisfy [pflag.Value] interface
func (f *FunctionInfoParameterStyle) Type() string {
	return "FunctionInfoParameterStyle"
}

// Function language. When **EXTERNAL** is used, the language of the routine
// function should be specified in the __external_language__ field, and the
// __return_params__ of the function cannot be used (as **TABLE** return type is
// not supported), and the __sql_data_access__ field must be **NO_SQL**.
type FunctionInfoRoutineBody string

const FunctionInfoRoutineBodyExternal FunctionInfoRoutineBody = `EXTERNAL`

const FunctionInfoRoutineBodySql FunctionInfoRoutineBody = `SQL`

// String representation for [fmt.Print]
func (f *FunctionInfoRoutineBody) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoRoutineBody) Set(v string) error {
	switch v {
	case `EXTERNAL`, `SQL`:
		*f = FunctionInfoRoutineBody(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "SQL"`, v)
	}
}

// Type always returns FunctionInfoRoutineBody to satisfy [pflag.Value] interface
func (f *FunctionInfoRoutineBody) Type() string {
	return "FunctionInfoRoutineBody"
}

// The security type of the function.
type FunctionInfoSecurityType string

const FunctionInfoSecurityTypeDefiner FunctionInfoSecurityType = `DEFINER`

// String representation for [fmt.Print]
func (f *FunctionInfoSecurityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoSecurityType) Set(v string) error {
	switch v {
	case `DEFINER`:
		*f = FunctionInfoSecurityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFINER"`, v)
	}
}

// Type always returns FunctionInfoSecurityType to satisfy [pflag.Value] interface
func (f *FunctionInfoSecurityType) Type() string {
	return "FunctionInfoSecurityType"
}

// Function SQL data access.
type FunctionInfoSqlDataAccess string

const FunctionInfoSqlDataAccessContainsSql FunctionInfoSqlDataAccess = `CONTAINS_SQL`

const FunctionInfoSqlDataAccessNoSql FunctionInfoSqlDataAccess = `NO_SQL`

const FunctionInfoSqlDataAccessReadsSqlData FunctionInfoSqlDataAccess = `READS_SQL_DATA`

// String representation for [fmt.Print]
func (f *FunctionInfoSqlDataAccess) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoSqlDataAccess) Set(v string) error {
	switch v {
	case `CONTAINS_SQL`, `NO_SQL`, `READS_SQL_DATA`:
		*f = FunctionInfoSqlDataAccess(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTAINS_SQL", "NO_SQL", "READS_SQL_DATA"`, v)
	}
}

// Type always returns FunctionInfoSqlDataAccess to satisfy [pflag.Value] interface
func (f *FunctionInfoSqlDataAccess) Type() string {
	return "FunctionInfoSqlDataAccess"
}

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of parameter.
	Name string `json:"name"`
	// Default value of the parameter.
	ParameterDefault string `json:"parameter_default,omitempty"`
	// The mode of the function parameter.
	ParameterMode FunctionParameterMode `json:"parameter_mode,omitempty"`
	// The type of function parameter.
	ParameterType FunctionParameterType `json:"parameter_type,omitempty"`
	// Ordinal position of column (starting at position 0).
	Position int `json:"position"`
	// Format of IntervalType.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type spec, JSON-serialized.
	TypeJson string `json:"type_json,omitempty"`

	TypeName ColumnTypeName `json:"type_name"`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type spec, SQL/catalogString text.
	TypeText string `json:"type_text"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FunctionParameterInfos struct {
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	Parameters []FunctionParameterInfo `json:"parameters,omitempty"`
}

// The mode of the function parameter.
type FunctionParameterMode string

const FunctionParameterModeIn FunctionParameterMode = `IN`

// String representation for [fmt.Print]
func (f *FunctionParameterMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterMode) Set(v string) error {
	switch v {
	case `IN`:
		*f = FunctionParameterMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN"`, v)
	}
}

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (f *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

// The type of function parameter.
type FunctionParameterType string

const FunctionParameterTypeColumn FunctionParameterType = `COLUMN`

const FunctionParameterTypeParam FunctionParameterType = `PARAM`

// String representation for [fmt.Print]
func (f *FunctionParameterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterType) Set(v string) error {
	switch v {
	case `COLUMN`, `PARAM`:
		*f = FunctionParameterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COLUMN", "PARAM"`, v)
	}
}

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (f *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

// GCP temporary credentials for API authentication. Read more at
// https://developers.google.com/identity/protocols/oauth2/service-account
type GcpOauthToken struct {
	OauthToken string `json:"oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpOauthToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpOauthToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Azure cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialAzureOptions struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python)
	Resources []string `json:"resources,omitempty"`
}

// The GCP cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialGcpOptions struct {
	// The scopes to which the temporary GCP credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://google-auth.readthedocs.io/en/latest/reference/google.auth.html#google.auth.credentials.Credentials)
	Scopes []string `json:"scopes,omitempty"`
}

type GenerateTemporaryServiceCredentialRequest struct {
	// The Azure cloud options to customize the requested temporary credential
	AzureOptions *GenerateTemporaryServiceCredentialAzureOptions `json:"azure_options,omitempty"`
	// The name of the service credential used to generate a temporary
	// credential
	CredentialName string `json:"credential_name"`
	// The GCP cloud options to customize the requested temporary credential
	GcpOptions *GenerateTemporaryServiceCredentialGcpOptions `json:"gcp_options,omitempty"`
}

type GenerateTemporaryTableCredentialRequest struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	Operation TableOperation `json:"operation,omitempty"`
	// UUID of the table to read or write.
	TableId string `json:"table_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateTemporaryTableCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryTableCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenerateTemporaryTableCredentialResponse struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials *AwsCredentials `json:"aws_temp_credentials,omitempty"`
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	// Azure temporary credentials for API authentication. Read more at
	// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
	AzureUserDelegationSas *AzureUserDelegationSas `json:"azure_user_delegation_sas,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken *GcpOauthToken `json:"gcp_oauth_token,omitempty"`
	// R2 temporary credentials for API authentication. Read more at
	// https://developers.cloudflare.com/r2/api/s3/tokens/.
	R2TempCredentials *R2Credentials `json:"r2_temp_credentials,omitempty"`
	// The URL of the storage path accessible by the temporary credential.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateTemporaryTableCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryTableCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Gets the metastore assignment for a workspace
type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

// Get a metastore
type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

// Gets the named storage credential
type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`
}

// Get an artifact allowlist
type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `json:"-" url:"-"`
}

// Get securable workspace bindings
type GetBindingsRequest struct {
	// Maximum number of workspace bindings to return. - When set to 0, the page
	// length is set to a server configured value (recommended); - When set to a
	// value greater than 0, the page length is the minimum of this value and a
	// server configured value; - When set to a value less than 0, an invalid
	// parameter error is returned; - If not set, all the workspace bindings are
	// returned (not recommended).
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The name of the securable.
	SecurableName string `json:"-" url:"-"`
	// The type of the securable to bind to a workspace.
	SecurableType GetBindingsSecurableType `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetBindingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetBindingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetBindingsSecurableType string

const GetBindingsSecurableTypeCatalog GetBindingsSecurableType = `catalog`

const GetBindingsSecurableTypeCredential GetBindingsSecurableType = `credential`

const GetBindingsSecurableTypeExternalLocation GetBindingsSecurableType = `external_location`

const GetBindingsSecurableTypeStorageCredential GetBindingsSecurableType = `storage_credential`

// String representation for [fmt.Print]
func (f *GetBindingsSecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetBindingsSecurableType) Set(v string) error {
	switch v {
	case `catalog`, `credential`, `external_location`, `storage_credential`:
		*f = GetBindingsSecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "catalog", "credential", "external_location", "storage_credential"`, v)
	}
}

// Type always returns GetBindingsSecurableType to satisfy [pflag.Value] interface
func (f *GetBindingsSecurableType) Type() string {
	return "GetBindingsSecurableType"
}

// Get Model Version By Alias
type GetByAliasRequest struct {
	// The name of the alias
	Alias string `json:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetByAliasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetByAliasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a catalog
type GetCatalogRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCatalogRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a connection
type GetConnectionRequest struct {
	// Name of the connection.
	Name string `json:"-" url:"-"`
}

// Get a credential
type GetCredentialRequest struct {
	// Name of the credential.
	NameArg string `json:"-" url:"-"`
}

// Get effective permissions
type GetEffectiveRequest struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetEffectiveRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEffectiveRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get an external location
type GetExternalLocationRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetExternalLocationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetExternalLocationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a function
type GetFunctionRequest struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get permissions
type GetGrantRequest struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetGrantRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetGrantRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a metastore
type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
}

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope GetMetastoreSummaryResponseDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetMetastoreSummaryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetMetastoreSummaryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The scope of Delta Sharing enabled for the metastore.
type GetMetastoreSummaryResponseDeltaSharingScope string

const GetMetastoreSummaryResponseDeltaSharingScopeInternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL`

const GetMetastoreSummaryResponseDeltaSharingScopeInternalAndExternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *GetMetastoreSummaryResponseDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetMetastoreSummaryResponseDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = GetMetastoreSummaryResponseDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns GetMetastoreSummaryResponseDeltaSharingScope to satisfy [pflag.Value] interface
func (f *GetMetastoreSummaryResponseDeltaSharingScope) Type() string {
	return "GetMetastoreSummaryResponseDeltaSharingScope"
}

// Get a Model Version
type GetModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" url:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get an Online Table
type GetOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" url:"-"`
}

// Get a table monitor
type GetQualityMonitorRequest struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

// Get information for a single resource quota.
type GetQuotaRequest struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	ParentFullName string `json:"-" url:"-"`
	// Securable type of the quota parent.
	ParentSecurableType string `json:"-" url:"-"`
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	QuotaName string `json:"-" url:"-"`
}

type GetQuotaResponse struct {
	// The returned QuotaInfo.
	QuotaInfo *QuotaInfo `json:"quota_info,omitempty"`
}

// Get refresh
type GetRefreshRequest struct {
	// ID of the refresh.
	RefreshId string `json:"-" url:"-"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

// Get a Registered Model
type GetRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// Whether to include registered model aliases in the response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a schema
type GetSchemaRequest struct {
	// Full name of the schema.
	FullName string `json:"-" url:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetSchemaRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSchemaRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a credential
type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
}

// Get a table
type GetTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `json:"-" url:"include_delta_metadata,omitempty"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get catalog workspace bindings
type GetWorkspaceBindingRequest struct {
	// The name of the catalog.
	Name string `json:"-" url:"-"`
}

type IsolationMode string

const IsolationModeIsolationModeIsolated IsolationMode = `ISOLATION_MODE_ISOLATED`

const IsolationModeIsolationModeOpen IsolationMode = `ISOLATION_MODE_OPEN`

// String representation for [fmt.Print]
func (f *IsolationMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IsolationMode) Set(v string) error {
	switch v {
	case `ISOLATION_MODE_ISOLATED`, `ISOLATION_MODE_OPEN`:
		*f = IsolationMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ISOLATION_MODE_ISOLATED", "ISOLATION_MODE_OPEN"`, v)
	}
}

// Type always returns IsolationMode to satisfy [pflag.Value] interface
func (f *IsolationMode) Type() string {
	return "IsolationMode"
}

// Get all workspaces assigned to a metastore
type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse struct {
	WorkspaceIds []int64 `json:"workspace_ids,omitempty"`
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`
}

// List catalogs
type ListCatalogsRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of catalogs to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid catalogs are returned (not
	// recommended). - Note: The number of returned catalogs might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further catalogs can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCatalogsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCatalogsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	Catalogs []CatalogInfo `json:"catalogs,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCatalogsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCatalogsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List connections
type ListConnectionsRequest struct {
	// Maximum number of connections to return. - If not set, all connections
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListConnectionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConnectionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListConnectionsResponse struct {
	// An array of connection information objects.
	Connections []ConnectionInfo `json:"connections,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListConnectionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConnectionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List credentials
type ListCredentialsRequest struct {
	// Maximum number of credentials to return. - If not set, the default max
	// page size is used. - When set to a value greater than 0, the page length
	// is the minimum of this value and a server-configured value. - When set to
	// 0, the page length is set to a server-configured value (recommended). -
	// When set to a value less than 0, an invalid parameter error is returned.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token to retrieve the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Return only credentials for the specified purpose.
	Purpose CredentialPurpose `json:"-" url:"purpose,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCredentialsResponse struct {
	Credentials []CredentialInfo `json:"credentials,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List external locations
type ListExternalLocationsRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExternalLocationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLocationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExternalLocationsResponse struct {
	// An array of external locations.
	ExternalLocations []ExternalLocationInfo `json:"external_locations,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExternalLocationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLocationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List functions
type ListFunctionsRequest struct {
	// Name of parent catalog for functions of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema of functions.
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFunctionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFunctionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFunctionsResponse struct {
	// An array of function information objects.
	Functions []FunctionInfo `json:"functions,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFunctionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFunctionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	Metastores []MetastoreInfo `json:"metastores,omitempty"`
}

// List Model Versions
type ListModelVersionsRequest struct {
	// The full three-level name of the registered model under which to list
	// model versions
	FullName string `json:"-" url:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelVersionsResponse struct {
	ModelVersions []ModelVersionInfo `json:"model_versions,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List all resource quotas under a metastore.
type ListQuotasRequest struct {
	// The number of quotas to return.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token for the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQuotasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQuotasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListQuotasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of returned QuotaInfos.
	Quotas []QuotaInfo `json:"quotas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQuotasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQuotasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List refreshes
type ListRefreshesRequest struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

// List Registered Models
type ListRegisteredModelsRequest struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	CatalogName string `json:"-" url:"catalog_name,omitempty"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Max number of registered models to return.
	//
	// If both catalog and schema are specified: - when max_results is not
	// specified, the page length is set to a server configured value (10000, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (10000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (10000, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	//
	// If neither schema nor catalog is specified: - when max_results is not
	// specified, the page length is set to a server configured value (100, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (1000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (100, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	SchemaName string `json:"-" url:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRegisteredModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegisteredModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRegisteredModelsResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	NextPageToken string `json:"next_page_token,omitempty"`

	RegisteredModels []RegisteredModelInfo `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRegisteredModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegisteredModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List schemas
type ListSchemasRequest struct {
	// Parent catalog for schemas of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSchemasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchemasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of schema information objects.
	Schemas []SchemaInfo `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSchemasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchemasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List credentials
type ListStorageCredentialsRequest struct {
	// Maximum number of storage credentials to return. If not set, all the
	// storage credentials are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStorageCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStorageCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListStorageCredentialsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStorageCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStorageCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List table summaries
type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern string `json:"-" url:"schema_name_pattern,omitempty"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern string `json:"-" url:"table_name_pattern,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSummariesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSummariesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List system schemas
type ListSystemSchemasRequest struct {
	// Maximum number of schemas to return. - When set to 0, the page length is
	// set to a server configured value (recommended); - When set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - When set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all the schemas are returned (not
	// recommended).
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// The ID for the metastore in which the system schema resides.
	MetastoreId string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSystemSchemasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSystemSchemasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSystemSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of system schema information objects.
	Schemas []SystemSchemaInfo `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSystemSchemasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSystemSchemasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of table summaries.
	Tables []TableSummary `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTableSummariesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTableSummariesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List tables
type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `json:"-" url:"include_delta_metadata,omitempty"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Whether to omit the columns of the table from the response or not.
	OmitColumns bool `json:"-" url:"omit_columns,omitempty"`
	// Whether to omit the properties of the table from the response or not.
	OmitProperties bool `json:"-" url:"omit_properties,omitempty"`
	// Whether to omit the username of the table (e.g. owner, updated_by,
	// created_by) from the response or not.
	OmitUsername bool `json:"-" url:"omit_username,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema of tables.
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTablesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTablesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTablesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of table information objects.
	Tables []TableInfo `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTablesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTablesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List Volumes
type ListVolumesRequest struct {
	// The identifier of the catalog
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of volumes to return (page length).
	//
	// If not set, the page length is set to a server configured value (10000,
	// as of 1/29/2024). - when set to a value greater than 0, the page length
	// is the minimum of this value and a server configured value (10000, as of
	// 1/29/2024); - when set to 0, the page length is set to a server
	// configured value (10000, as of 1/29/2024) (recommended); - when set to a
	// value less than 0, an invalid parameter error is returned;
	//
	// Note: this parameter controls only the maximum number of volumes to
	// return. The actual number of volumes returned in a page may be smaller
	// than this value, including 0, even if there are more pages.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The identifier of the schema
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListVolumesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVolumesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListVolumesResponseContent struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Volumes []VolumeInfo `json:"volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListVolumesResponseContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVolumesResponseContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The artifact pattern matching type
type MatchType string

const MatchTypePrefixMatch MatchType = `PREFIX_MATCH`

// String representation for [fmt.Print]
func (f *MatchType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MatchType) Set(v string) error {
	switch v {
	case `PREFIX_MATCH`:
		*f = MatchType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PREFIX_MATCH"`, v)
	}
}

// Type always returns MatchType to satisfy [pflag.Value] interface
func (f *MatchType) Type() string {
	return "MatchType"
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id"`
	// The unique ID of the Databricks workspace.
	WorkspaceId int64 `json:"workspace_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetastoreAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope MetastoreInfoDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetastoreInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetastoreInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The scope of Delta Sharing enabled for the metastore.
type MetastoreInfoDeltaSharingScope string

const MetastoreInfoDeltaSharingScopeInternal MetastoreInfoDeltaSharingScope = `INTERNAL`

const MetastoreInfoDeltaSharingScopeInternalAndExternal MetastoreInfoDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *MetastoreInfoDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MetastoreInfoDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = MetastoreInfoDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns MetastoreInfoDeltaSharingScope to satisfy [pflag.Value] interface
func (f *MetastoreInfoDeltaSharingScope) Type() string {
	return "MetastoreInfoDeltaSharingScope"
}

type ModelVersionInfo struct {
	// List of aliases associated with the model version
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog containing the model version
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the model version
	Comment string `json:"comment,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the model version
	CreatedBy string `json:"created_by,omitempty"`
	// The unique identifier of the model version
	Id string `json:"id,omitempty"`
	// The unique identifier of the metastore containing the model version
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName string `json:"model_name,omitempty"`
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies *DependencyList `json:"model_version_dependencies,omitempty"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId string `json:"run_id,omitempty"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId int `json:"run_workspace_id,omitempty"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName string `json:"schema_name,omitempty"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source string `json:"source,omitempty"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status ModelVersionInfoStatus `json:"status,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the model version last time
	UpdatedBy string `json:"updated_by,omitempty"`
	// Integer model version number, used to reference the model version in API
	// requests.
	Version int `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelVersionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Current status of the model version. Newly created model versions start in
// PENDING_REGISTRATION status, then move to READY status once the model version
// files are uploaded and the model version is finalized. Only model versions in
// READY status can be loaded for inference or served.
type ModelVersionInfoStatus string

const ModelVersionInfoStatusFailedRegistration ModelVersionInfoStatus = `FAILED_REGISTRATION`

const ModelVersionInfoStatusPendingRegistration ModelVersionInfoStatus = `PENDING_REGISTRATION`

const ModelVersionInfoStatusReady ModelVersionInfoStatus = `READY`

// String representation for [fmt.Print]
func (f *ModelVersionInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelVersionInfoStatus) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = ModelVersionInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Type always returns ModelVersionInfoStatus to satisfy [pflag.Value] interface
func (f *ModelVersionInfoStatus) Type() string {
	return "ModelVersionInfoStatus"
}

type MonitorCronSchedule struct {
	// Read only field that indicates whether a schedule is paused or not.
	PauseStatus MonitorCronSchedulePauseStatus `json:"pause_status,omitempty"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	TimezoneId string `json:"timezone_id"`
}

// Read only field that indicates whether a schedule is paused or not.
type MonitorCronSchedulePauseStatus string

const MonitorCronSchedulePauseStatusPaused MonitorCronSchedulePauseStatus = `PAUSED`

const MonitorCronSchedulePauseStatusUnpaused MonitorCronSchedulePauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *MonitorCronSchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorCronSchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = MonitorCronSchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Type always returns MonitorCronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *MonitorCronSchedulePauseStatus) Type() string {
	return "MonitorCronSchedulePauseStatus"
}

type MonitorDataClassificationConfig struct {
	// Whether data classification is enabled.
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MonitorDataClassificationConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorDataClassificationConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MonitorDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

type MonitorInferenceLog struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []string `json:"granularities"`
	// Optional column that contains the ground truth for the prediction.
	LabelCol string `json:"label_col,omitempty"`
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	ModelIdCol string `json:"model_id_col"`
	// Column that contains the output/prediction from the model.
	PredictionCol string `json:"prediction_col"`
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	PredictionProbaCol string `json:"prediction_proba_col,omitempty"`
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	ProblemType MonitorInferenceLogProblemType `json:"problem_type"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol string `json:"timestamp_col"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MonitorInferenceLog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorInferenceLog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Problem type the model aims to solve. Determines the type of model-quality
// metrics that will be computed.
type MonitorInferenceLogProblemType string

const MonitorInferenceLogProblemTypeProblemTypeClassification MonitorInferenceLogProblemType = `PROBLEM_TYPE_CLASSIFICATION`

const MonitorInferenceLogProblemTypeProblemTypeRegression MonitorInferenceLogProblemType = `PROBLEM_TYPE_REGRESSION`

// String representation for [fmt.Print]
func (f *MonitorInferenceLogProblemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorInferenceLogProblemType) Set(v string) error {
	switch v {
	case `PROBLEM_TYPE_CLASSIFICATION`, `PROBLEM_TYPE_REGRESSION`:
		*f = MonitorInferenceLogProblemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PROBLEM_TYPE_CLASSIFICATION", "PROBLEM_TYPE_REGRESSION"`, v)
	}
}

// Type always returns MonitorInferenceLogProblemType to satisfy [pflag.Value] interface
func (f *MonitorInferenceLogProblemType) Type() string {
	return "MonitorInferenceLogProblemType"
}

type MonitorInfo struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir string `json:"assets_dir,omitempty"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	DriftMetricsTableName string `json:"drift_metrics_table_name"`
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// The latest failure message of the monitor (if any).
	LatestMonitorFailureMsg string `json:"latest_monitor_failure_msg,omitempty"`
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	MonitorVersion string `json:"monitor_version"`
	// The notification settings for the monitor.
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	OutputSchemaName string `json:"output_schema_name,omitempty"`
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	ProfileMetricsTableName string `json:"profile_metrics_table_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`
	// The status of the monitor.
	Status MonitorInfoStatus `json:"status"`
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	TableName string `json:"table_name"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MonitorInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of the monitor.
type MonitorInfoStatus string

const MonitorInfoStatusMonitorStatusActive MonitorInfoStatus = `MONITOR_STATUS_ACTIVE`

const MonitorInfoStatusMonitorStatusDeletePending MonitorInfoStatus = `MONITOR_STATUS_DELETE_PENDING`

const MonitorInfoStatusMonitorStatusError MonitorInfoStatus = `MONITOR_STATUS_ERROR`

const MonitorInfoStatusMonitorStatusFailed MonitorInfoStatus = `MONITOR_STATUS_FAILED`

const MonitorInfoStatusMonitorStatusPending MonitorInfoStatus = `MONITOR_STATUS_PENDING`

// String representation for [fmt.Print]
func (f *MonitorInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorInfoStatus) Set(v string) error {
	switch v {
	case `MONITOR_STATUS_ACTIVE`, `MONITOR_STATUS_DELETE_PENDING`, `MONITOR_STATUS_ERROR`, `MONITOR_STATUS_FAILED`, `MONITOR_STATUS_PENDING`:
		*f = MonitorInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MONITOR_STATUS_ACTIVE", "MONITOR_STATUS_DELETE_PENDING", "MONITOR_STATUS_ERROR", "MONITOR_STATUS_FAILED", "MONITOR_STATUS_PENDING"`, v)
	}
}

// Type always returns MonitorInfoStatus to satisfy [pflag.Value] interface
func (f *MonitorInfoStatus) Type() string {
	return "MonitorInfoStatus"
}

type MonitorMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition string `json:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns []string `json:"input_columns"`
	// Name of the metric in the output tables.
	Name string `json:"name"`
	// The output type of the custom metric.
	OutputDataType string `json:"output_data_type"`
	// Can only be one of ``"CUSTOM_METRIC_TYPE_AGGREGATE"``,
	// ``"CUSTOM_METRIC_TYPE_DERIVED"``, or ``"CUSTOM_METRIC_TYPE_DRIFT"``. The
	// ``"CUSTOM_METRIC_TYPE_AGGREGATE"`` and ``"CUSTOM_METRIC_TYPE_DERIVED"``
	// metrics are computed on a single table, whereas the
	// ``"CUSTOM_METRIC_TYPE_DRIFT"`` compare metrics across baseline and input
	// table, or across the two consecutive time windows. -
	// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
	// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed
	// aggregate metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously
	// computed aggregate or derived metrics
	Type MonitorMetricType `json:"type"`
}

// Can only be one of “"CUSTOM_METRIC_TYPE_AGGREGATE"“,
// “"CUSTOM_METRIC_TYPE_DERIVED"“, or “"CUSTOM_METRIC_TYPE_DRIFT"“. The
// “"CUSTOM_METRIC_TYPE_AGGREGATE"“ and “"CUSTOM_METRIC_TYPE_DERIVED"“
// metrics are computed on a single table, whereas the
// “"CUSTOM_METRIC_TYPE_DRIFT"“ compare metrics across baseline and input
// table, or across the two consecutive time windows. -
// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed aggregate
// metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously computed aggregate
// or derived metrics
type MonitorMetricType string

const MonitorMetricTypeCustomMetricTypeAggregate MonitorMetricType = `CUSTOM_METRIC_TYPE_AGGREGATE`

const MonitorMetricTypeCustomMetricTypeDerived MonitorMetricType = `CUSTOM_METRIC_TYPE_DERIVED`

const MonitorMetricTypeCustomMetricTypeDrift MonitorMetricType = `CUSTOM_METRIC_TYPE_DRIFT`

// String representation for [fmt.Print]
func (f *MonitorMetricType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorMetricType) Set(v string) error {
	switch v {
	case `CUSTOM_METRIC_TYPE_AGGREGATE`, `CUSTOM_METRIC_TYPE_DERIVED`, `CUSTOM_METRIC_TYPE_DRIFT`:
		*f = MonitorMetricType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CUSTOM_METRIC_TYPE_AGGREGATE", "CUSTOM_METRIC_TYPE_DERIVED", "CUSTOM_METRIC_TYPE_DRIFT"`, v)
	}
}

// Type always returns MonitorMetricType to satisfy [pflag.Value] interface
func (f *MonitorMetricType) Type() string {
	return "MonitorMetricType"
}

type MonitorNotifications struct {
	// Who to send notifications to on monitor failure.
	OnFailure *MonitorDestination `json:"on_failure,omitempty"`
	// Who to send notifications to when new data classification tags are
	// detected.
	OnNewClassificationTagDetected *MonitorDestination `json:"on_new_classification_tag_detected,omitempty"`
}

type MonitorRefreshInfo struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	EndTimeMs int64 `json:"end_time_ms,omitempty"`
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	Message string `json:"message,omitempty"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"refresh_id"`
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	StartTimeMs int64 `json:"start_time_ms"`
	// The current state of the refresh.
	State MonitorRefreshInfoState `json:"state"`
	// The method by which the refresh was triggered.
	Trigger MonitorRefreshInfoTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MonitorRefreshInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorRefreshInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The current state of the refresh.
type MonitorRefreshInfoState string

const MonitorRefreshInfoStateCanceled MonitorRefreshInfoState = `CANCELED`

const MonitorRefreshInfoStateFailed MonitorRefreshInfoState = `FAILED`

const MonitorRefreshInfoStatePending MonitorRefreshInfoState = `PENDING`

const MonitorRefreshInfoStateRunning MonitorRefreshInfoState = `RUNNING`

const MonitorRefreshInfoStateSuccess MonitorRefreshInfoState = `SUCCESS`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCESS`:
		*f = MonitorRefreshInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "PENDING", "RUNNING", "SUCCESS"`, v)
	}
}

// Type always returns MonitorRefreshInfoState to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoState) Type() string {
	return "MonitorRefreshInfoState"
}

// The method by which the refresh was triggered.
type MonitorRefreshInfoTrigger string

const MonitorRefreshInfoTriggerManual MonitorRefreshInfoTrigger = `MANUAL`

const MonitorRefreshInfoTriggerSchedule MonitorRefreshInfoTrigger = `SCHEDULE`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoTrigger) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoTrigger) Set(v string) error {
	switch v {
	case `MANUAL`, `SCHEDULE`:
		*f = MonitorRefreshInfoTrigger(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANUAL", "SCHEDULE"`, v)
	}
}

// Type always returns MonitorRefreshInfoTrigger to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoTrigger) Type() string {
	return "MonitorRefreshInfoTrigger"
}

type MonitorRefreshListResponse struct {
	// List of refreshes.
	Refreshes []MonitorRefreshInfo `json:"refreshes,omitempty"`
}

type MonitorSnapshot struct {
}

type MonitorTimeSeries struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []string `json:"granularities"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol string `json:"timestamp_col"`
}

type NamedTableConstraint struct {
	// The name of the constraint.
	Name string `json:"name"`
}

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"name,omitempty"`
	// Specification of the online table.
	Spec *OnlineTableSpec `json:"spec,omitempty"`
	// Online Table data synchronization status
	Status *OnlineTableStatus `json:"status,omitempty"`
	// Data serving REST API URL for this table
	TableServingUrl string `json:"table_serving_url,omitempty"`
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OnlineTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Specification of an online table.
type OnlineTableSpec struct {
	// Whether to create a full-copy pipeline -- a pipeline that stops after
	// creates a full copy of the source table upon initialization and does not
	// process any change data feeds (CDFs) afterwards. The pipeline can still
	// be manually triggered afterwards, but it always perform a full copy of
	// the source table and there are no incremental updates. This mode is
	// useful for syncing views or tables without CDFs to online tables. Note
	// that the full-copy pipeline only supports "triggered" scheduling policy.
	PerformFullCopy bool `json:"perform_full_copy,omitempty"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Pipeline runs continuously after generating the initial data.
	RunContinuously *OnlineTableSpecContinuousSchedulingPolicy `json:"run_continuously,omitempty"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered *OnlineTableSpecTriggeredSchedulingPolicy `json:"run_triggered,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey string `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OnlineTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

type OnlineTableSpecTriggeredSchedulingPolicy struct {
}

// The state of an online table.
type OnlineTableState string

const OnlineTableStateOffline OnlineTableState = `OFFLINE`

const OnlineTableStateOfflineFailed OnlineTableState = `OFFLINE_FAILED`

const OnlineTableStateOnline OnlineTableState = `ONLINE`

const OnlineTableStateOnlineContinuousUpdate OnlineTableState = `ONLINE_CONTINUOUS_UPDATE`

const OnlineTableStateOnlineNoPendingUpdate OnlineTableState = `ONLINE_NO_PENDING_UPDATE`

const OnlineTableStateOnlinePipelineFailed OnlineTableState = `ONLINE_PIPELINE_FAILED`

const OnlineTableStateOnlineTriggeredUpdate OnlineTableState = `ONLINE_TRIGGERED_UPDATE`

const OnlineTableStateOnlineUpdatingPipelineResources OnlineTableState = `ONLINE_UPDATING_PIPELINE_RESOURCES`

const OnlineTableStateProvisioning OnlineTableState = `PROVISIONING`

const OnlineTableStateProvisioningInitialSnapshot OnlineTableState = `PROVISIONING_INITIAL_SNAPSHOT`

const OnlineTableStateProvisioningPipelineResources OnlineTableState = `PROVISIONING_PIPELINE_RESOURCES`

// String representation for [fmt.Print]
func (f *OnlineTableState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OnlineTableState) Set(v string) error {
	switch v {
	case `OFFLINE`, `OFFLINE_FAILED`, `ONLINE`, `ONLINE_CONTINUOUS_UPDATE`, `ONLINE_NO_PENDING_UPDATE`, `ONLINE_PIPELINE_FAILED`, `ONLINE_TRIGGERED_UPDATE`, `ONLINE_UPDATING_PIPELINE_RESOURCES`, `PROVISIONING`, `PROVISIONING_INITIAL_SNAPSHOT`, `PROVISIONING_PIPELINE_RESOURCES`:
		*f = OnlineTableState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OFFLINE", "OFFLINE_FAILED", "ONLINE", "ONLINE_CONTINUOUS_UPDATE", "ONLINE_NO_PENDING_UPDATE", "ONLINE_PIPELINE_FAILED", "ONLINE_TRIGGERED_UPDATE", "ONLINE_UPDATING_PIPELINE_RESOURCES", "PROVISIONING", "PROVISIONING_INITIAL_SNAPSHOT", "PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Type always returns OnlineTableState to satisfy [pflag.Value] interface
func (f *OnlineTableState) Type() string {
	return "OnlineTableState"
}

// Status of an online table.
type OnlineTableStatus struct {
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus *ContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	// The state of the online table.
	DetailedState OnlineTableState `json:"detailed_state,omitempty"`
	// Detailed status of an online table. Shown if the online table is in the
	// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
	FailedStatus *FailedStatus `json:"failed_status,omitempty"`
	// A text description of the current state of the online table.
	Message string `json:"message,omitempty"`
	// Detailed status of an online table. Shown if the online table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus *ProvisioningStatus `json:"provisioning_status,omitempty"`
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus *TriggeredUpdateStatus `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OnlineTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionsChange struct {
	// The set of privileges to add.
	Add []Privilege `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	Remove []Privilege `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PermissionsChange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionsChange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionsList struct {
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds float64 `json:"estimated_completion_time_seconds,omitempty"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing int64 `json:"latest_version_currently_processing,omitempty"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion float64 `json:"sync_progress_completion,omitempty"`
	// The number of rows that have been synced in this update.
	SyncedRowCount int64 `json:"synced_row_count,omitempty"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount int64 `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineProgress) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	Name string `json:"name"`
}

type Privilege string

const PrivilegeAccess Privilege = `ACCESS`

const PrivilegeAllPrivileges Privilege = `ALL_PRIVILEGES`

const PrivilegeApplyTag Privilege = `APPLY_TAG`

const PrivilegeCreate Privilege = `CREATE`

const PrivilegeCreateCatalog Privilege = `CREATE_CATALOG`

const PrivilegeCreateConnection Privilege = `CREATE_CONNECTION`

const PrivilegeCreateExternalLocation Privilege = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable Privilege = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateExternalVolume Privilege = `CREATE_EXTERNAL_VOLUME`

const PrivilegeCreateForeignCatalog Privilege = `CREATE_FOREIGN_CATALOG`

const PrivilegeCreateForeignSecurable Privilege = `CREATE_FOREIGN_SECURABLE`

const PrivilegeCreateFunction Privilege = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage Privilege = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView Privilege = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateModel Privilege = `CREATE_MODEL`

const PrivilegeCreateProvider Privilege = `CREATE_PROVIDER`

const PrivilegeCreateRecipient Privilege = `CREATE_RECIPIENT`

const PrivilegeCreateSchema Privilege = `CREATE_SCHEMA`

const PrivilegeCreateServiceCredential Privilege = `CREATE_SERVICE_CREDENTIAL`

const PrivilegeCreateShare Privilege = `CREATE_SHARE`

const PrivilegeCreateStorageCredential Privilege = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable Privilege = `CREATE_TABLE`

const PrivilegeCreateView Privilege = `CREATE_VIEW`

const PrivilegeCreateVolume Privilege = `CREATE_VOLUME`

const PrivilegeExecute Privilege = `EXECUTE`

const PrivilegeManage Privilege = `MANAGE`

const PrivilegeManageAllowlist Privilege = `MANAGE_ALLOWLIST`

const PrivilegeModify Privilege = `MODIFY`

const PrivilegeReadFiles Privilege = `READ_FILES`

const PrivilegeReadPrivateFiles Privilege = `READ_PRIVATE_FILES`

const PrivilegeReadVolume Privilege = `READ_VOLUME`

const PrivilegeRefresh Privilege = `REFRESH`

const PrivilegeSelect Privilege = `SELECT`

const PrivilegeSetSharePermission Privilege = `SET_SHARE_PERMISSION`

const PrivilegeUsage Privilege = `USAGE`

const PrivilegeUseCatalog Privilege = `USE_CATALOG`

const PrivilegeUseConnection Privilege = `USE_CONNECTION`

const PrivilegeUseMarketplaceAssets Privilege = `USE_MARKETPLACE_ASSETS`

const PrivilegeUseProvider Privilege = `USE_PROVIDER`

const PrivilegeUseRecipient Privilege = `USE_RECIPIENT`

const PrivilegeUseSchema Privilege = `USE_SCHEMA`

const PrivilegeUseShare Privilege = `USE_SHARE`

const PrivilegeWriteFiles Privilege = `WRITE_FILES`

const PrivilegeWritePrivateFiles Privilege = `WRITE_PRIVATE_FILES`

const PrivilegeWriteVolume Privilege = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *Privilege) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Privilege) Set(v string) error {
	switch v {
	case `ACCESS`, `ALL_PRIVILEGES`, `APPLY_TAG`, `CREATE`, `CREATE_CATALOG`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_EXTERNAL_VOLUME`, `CREATE_FOREIGN_CATALOG`, `CREATE_FOREIGN_SECURABLE`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_MODEL`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SERVICE_CREDENTIAL`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `CREATE_VOLUME`, `EXECUTE`, `MANAGE`, `MANAGE_ALLOWLIST`, `MODIFY`, `READ_FILES`, `READ_PRIVATE_FILES`, `READ_VOLUME`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_CONNECTION`, `USE_MARKETPLACE_ASSETS`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`, `WRITE_VOLUME`:
		*f = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCESS", "ALL_PRIVILEGES", "APPLY_TAG", "CREATE", "CREATE_CATALOG", "CREATE_CONNECTION", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_EXTERNAL_VOLUME", "CREATE_FOREIGN_CATALOG", "CREATE_FOREIGN_SECURABLE", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_MODEL", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SERVICE_CREDENTIAL", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "CREATE_VOLUME", "EXECUTE", "MANAGE", "MANAGE_ALLOWLIST", "MODIFY", "READ_FILES", "READ_PRIVATE_FILES", "READ_VOLUME", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_CONNECTION", "USE_MARKETPLACE_ASSETS", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES", "WRITE_VOLUME"`, v)
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An object containing map of key-value properties attached to the connection.
type PropertiesKvPairs map[string]string

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {
	State ProvisioningInfoState `json:"state,omitempty"`
}

type ProvisioningInfoState string

const ProvisioningInfoStateActive ProvisioningInfoState = `ACTIVE`

const ProvisioningInfoStateDegraded ProvisioningInfoState = `DEGRADED`

const ProvisioningInfoStateDeleting ProvisioningInfoState = `DELETING`

const ProvisioningInfoStateFailed ProvisioningInfoState = `FAILED`

const ProvisioningInfoStateProvisioning ProvisioningInfoState = `PROVISIONING`

const ProvisioningInfoStateUpdating ProvisioningInfoState = `UPDATING`

// String representation for [fmt.Print]
func (f *ProvisioningInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ProvisioningInfoState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DEGRADED`, `DELETING`, `FAILED`, `PROVISIONING`, `UPDATING`:
		*f = ProvisioningInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DEGRADED", "DELETING", "FAILED", "PROVISIONING", "UPDATING"`, v)
	}
}

// Type always returns ProvisioningInfoState to satisfy [pflag.Value] interface
func (f *ProvisioningInfoState) Type() string {
	return "ProvisioningInfoState"
}

// Detailed status of an online table. Shown if the online table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type ProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress *PipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

type QuotaInfo struct {
	// The timestamp that indicates when the quota count was last updated.
	LastRefreshedAt int64 `json:"last_refreshed_at,omitempty"`
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	ParentFullName string `json:"parent_full_name,omitempty"`
	// The quota parent securable type.
	ParentSecurableType SecurableType `json:"parent_securable_type,omitempty"`
	// The current usage of the resource quota.
	QuotaCount int `json:"quota_count,omitempty"`
	// The current limit of the resource quota.
	QuotaLimit int `json:"quota_limit,omitempty"`
	// The name of the quota.
	QuotaName string `json:"quota_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QuotaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QuotaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// R2 temporary credentials for API authentication. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/.
type R2Credentials struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId string `json:"access_key_id,omitempty"`
	// The secret access key associated with the access key.
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	// The generated JWT that users must pass to use the temporary credentials.
	SessionToken string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *R2Credentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s R2Credentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a Volume
type ReadVolumeRequest struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ReadVolumeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReadVolumeRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegenerateDashboardRequest struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegenerateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegenerateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegenerateDashboardResponse struct {
	// Id of the regenerated monitoring dashboard.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The directory where the regenerated dashboard is stored.
	ParentFolder string `json:"parent_folder,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegenerateDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegenerateDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Registered model alias.
type RegisteredModelAlias struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	AliasName string `json:"alias_name,omitempty"`
	// Integer version number of the model version to which this alias points.
	VersionNum int `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelInfo struct {
	// List of aliases associated with the registered model
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the registered model
	CreatedBy string `json:"created_by,omitempty"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the registered model
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the registered model
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the registered model resides
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Queue a metric refresh for a monitor
type RunRefreshRequest struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// The type of the parent catalog.
	CatalogType string `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of schema creator.
	CreatedBy string `json:"created_by,omitempty"`

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name,omitempty"`
	// Username of current owner of schema.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The unique identifier of the schema.
	SchemaId string `json:"schema_id,omitempty"`
	// Storage location for managed tables within schema.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified schema.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SchemaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A map of key-value properties attached to the securable.
type SecurableOptionsMap map[string]string

// A map of key-value properties attached to the securable.
type SecurablePropertiesMap map[string]string

// The type of Unity Catalog securable
type SecurableType string

const SecurableTypeCatalog SecurableType = `CATALOG`

const SecurableTypeCleanRoom SecurableType = `CLEAN_ROOM`

const SecurableTypeConnection SecurableType = `CONNECTION`

const SecurableTypeCredential SecurableType = `CREDENTIAL`

const SecurableTypeExternalLocation SecurableType = `EXTERNAL_LOCATION`

const SecurableTypeFunction SecurableType = `FUNCTION`

const SecurableTypeMetastore SecurableType = `METASTORE`

const SecurableTypePipeline SecurableType = `PIPELINE`

const SecurableTypeProvider SecurableType = `PROVIDER`

const SecurableTypeRecipient SecurableType = `RECIPIENT`

const SecurableTypeSchema SecurableType = `SCHEMA`

const SecurableTypeShare SecurableType = `SHARE`

const SecurableTypeStorageCredential SecurableType = `STORAGE_CREDENTIAL`

const SecurableTypeTable SecurableType = `TABLE`

const SecurableTypeVolume SecurableType = `VOLUME`

// String representation for [fmt.Print]
func (f *SecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SecurableType) Set(v string) error {
	switch v {
	case `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `FUNCTION`, `METASTORE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`:
		*f = SecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG", "CLEAN_ROOM", "CONNECTION", "CREDENTIAL", "EXTERNAL_LOCATION", "FUNCTION", "METASTORE", "PIPELINE", "PROVIDER", "RECIPIENT", "SCHEMA", "SHARE", "STORAGE_CREDENTIAL", "TABLE", "VOLUME"`, v)
	}
}

// Type always returns SecurableType to satisfy [pflag.Value] interface
func (f *SecurableType) Type() string {
	return "SecurableType"
}

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher `json:"artifact_matchers"`
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `json:"-" url:"-"`
}

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	Alias string `json:"alias" url:"-"`
	// Full name of the registered model
	FullName string `json:"full_name" url:"-"`
	// The version number of the model version to which the alias points
	VersionNum int `json:"version_num"`
}

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails struct {
	// The type of key encryption to use (affects headers from s3 client).
	Algorithm SseEncryptionDetailsAlgorithm `json:"algorithm,omitempty"`
	// When algorithm is **AWS_SSE_KMS** this field specifies the ARN of the SSE
	// key to use.
	AwsKmsKeyArn string `json:"aws_kms_key_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SseEncryptionDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SseEncryptionDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of key encryption to use (affects headers from s3 client).
type SseEncryptionDetailsAlgorithm string

const SseEncryptionDetailsAlgorithmAwsSseKms SseEncryptionDetailsAlgorithm = `AWS_SSE_KMS`

const SseEncryptionDetailsAlgorithmAwsSseS3 SseEncryptionDetailsAlgorithm = `AWS_SSE_S3`

// String representation for [fmt.Print]
func (f *SseEncryptionDetailsAlgorithm) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SseEncryptionDetailsAlgorithm) Set(v string) error {
	switch v {
	case `AWS_SSE_KMS`, `AWS_SSE_S3`:
		*f = SseEncryptionDetailsAlgorithm(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_SSE_KMS", "AWS_SSE_S3"`, v)
	}
}

// Type always returns SseEncryptionDetailsAlgorithm to satisfy [pflag.Value] interface
func (f *SseEncryptionDetailsAlgorithm) Type() string {
	return "SseEncryptionDetailsAlgorithm"
}

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleResponse `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponse `json:"databricks_gcp_service_account,omitempty"`
	// The full name of the credential.
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the credential.
	Id string `json:"id,omitempty"`

	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique within the metastore.
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential.
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StorageCredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageCredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SystemSchemaInfo struct {
	// Name of the system schema.
	Schema string `json:"schema,omitempty"`
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in.
	State SystemSchemaInfoState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SystemSchemaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SystemSchemaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The current state of enablement for the system schema. An empty string means
// the system schema is available and ready for opt-in.
type SystemSchemaInfoState string

const SystemSchemaInfoStateAvailable SystemSchemaInfoState = `AVAILABLE`

const SystemSchemaInfoStateDisableInitialized SystemSchemaInfoState = `DISABLE_INITIALIZED`

const SystemSchemaInfoStateEnableCompleted SystemSchemaInfoState = `ENABLE_COMPLETED`

const SystemSchemaInfoStateEnableInitialized SystemSchemaInfoState = `ENABLE_INITIALIZED`

const SystemSchemaInfoStateUnavailable SystemSchemaInfoState = `UNAVAILABLE`

// String representation for [fmt.Print]
func (f *SystemSchemaInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SystemSchemaInfoState) Set(v string) error {
	switch v {
	case `AVAILABLE`, `DISABLE_INITIALIZED`, `ENABLE_COMPLETED`, `ENABLE_INITIALIZED`, `UNAVAILABLE`:
		*f = SystemSchemaInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVAILABLE", "DISABLE_INITIALIZED", "ENABLE_COMPLETED", "ENABLE_INITIALIZED", "UNAVAILABLE"`, v)
	}
}

// Type always returns SystemSchemaInfoState to satisfy [pflag.Value] interface
func (f *SystemSchemaInfoState) Type() string {
	return "SystemSchemaInfoState"
}

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {
	ForeignKeyConstraint *ForeignKeyConstraint `json:"foreign_key_constraint,omitempty"`

	NamedTableConstraint *NamedTableConstraint `json:"named_table_constraint,omitempty"`

	PrimaryKeyConstraint *PrimaryKeyConstraint `json:"primary_key_constraint,omitempty"`
}

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName string `json:"table_full_name"`
}

type TableExistsResponse struct {
	// Whether the table exists or not.
	TableExists bool `json:"table_exists,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableExistsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableExistsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `json:"columns,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// Data source format
	DataSourceFormat DataSourceFormat `json:"data_source_format,omitempty"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt int64 `json:"deleted_at,omitempty"`
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs *DeltaRuntimePropertiesKvPairs `json:"delta_runtime_properties_kvpairs,omitempty"`

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of table, relative to parent schema.
	Name string `json:"name,omitempty"`
	// Username of current owner of table.
	Owner string `json:"owner,omitempty"`
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	PipelineId string `json:"pipeline_id,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	RowFilter *TableRowFilter `json:"row_filter,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	StorageLocation string `json:"storage_location,omitempty"`
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	TableConstraints []TableConstraint `json:"table_constraints,omitempty"`
	// The unique identifier of the table.
	TableId string `json:"table_id,omitempty"`

	TableType TableType `json:"table_type,omitempty"`
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition string `json:"view_definition,omitempty"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	ViewDependencies *DependencyList `json:"view_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableOperation string

const TableOperationRead TableOperation = `READ`

const TableOperationReadWrite TableOperation = `READ_WRITE`

// String representation for [fmt.Print]
func (f *TableOperation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableOperation) Set(v string) error {
	switch v {
	case `READ`, `READ_WRITE`:
		*f = TableOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ", "READ_WRITE"`, v)
	}
}

// Type always returns TableOperation to satisfy [pflag.Value] interface
func (f *TableOperation) Type() string {
	return "TableOperation"
}

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	FunctionName string `json:"function_name"`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames []string `json:"input_column_names"`
}

type TableSummary struct {
	// The full name of the table.
	FullName string `json:"full_name,omitempty"`

	TableType TableType `json:"table_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableType string

const TableTypeExternal TableType = `EXTERNAL`

const TableTypeExternalShallowClone TableType = `EXTERNAL_SHALLOW_CLONE`

const TableTypeForeign TableType = `FOREIGN`

const TableTypeManaged TableType = `MANAGED`

const TableTypeManagedShallowClone TableType = `MANAGED_SHALLOW_CLONE`

const TableTypeMaterializedView TableType = `MATERIALIZED_VIEW`

const TableTypeStreamingTable TableType = `STREAMING_TABLE`

const TableTypeView TableType = `VIEW`

// String representation for [fmt.Print]
func (f *TableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `EXTERNAL_SHALLOW_CLONE`, `FOREIGN`, `MANAGED`, `MANAGED_SHALLOW_CLONE`, `MATERIALIZED_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*f = TableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "EXTERNAL_SHALLOW_CLONE", "FOREIGN", "MANAGED", "MANAGED_SHALLOW_CLONE", "MATERIALIZED_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Type always returns TableType to satisfy [pflag.Value] interface
func (f *TableType) Type() string {
	return "TableType"
}

type TemporaryCredentials struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials *AwsCredentials `json:"aws_temp_credentials,omitempty"`
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken *GcpOauthToken `json:"gcp_oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TemporaryCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TemporaryCredentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string `json:"timestamp,omitempty"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress *PipelineProgress `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	MetastoreId string `json:"-" url:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type UnassignResponse struct {
}

type UpdateAssignmentResponse struct {
}

type UpdateBindingsSecurableType string

const UpdateBindingsSecurableTypeCatalog UpdateBindingsSecurableType = `catalog`

const UpdateBindingsSecurableTypeCredential UpdateBindingsSecurableType = `credential`

const UpdateBindingsSecurableTypeExternalLocation UpdateBindingsSecurableType = `external_location`

const UpdateBindingsSecurableTypeStorageCredential UpdateBindingsSecurableType = `storage_credential`

// String representation for [fmt.Print]
func (f *UpdateBindingsSecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateBindingsSecurableType) Set(v string) error {
	switch v {
	case `catalog`, `credential`, `external_location`, `storage_credential`:
		*f = UpdateBindingsSecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "catalog", "credential", "external_location", "storage_credential"`, v)
	}
}

// Type always returns UpdateBindingsSecurableType to satisfy [pflag.Value] interface
func (f *UpdateBindingsSecurableType) Type() string {
	return "UpdateBindingsSecurableType"
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode CatalogIsolationMode `json:"isolation_mode,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`
	// New name for the catalog.
	NewName string `json:"new_name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of catalog.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateConnection struct {
	// Name of the connection.
	Name string `json:"-" url:"-"`
	// New name for the connection.
	NewName string `json:"new_name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options"`
	// Username of current owner of the connection.
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool `json:"force,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the credential.
	NameArg string `json:"-" url:"-"`
	// New name of credential.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supply true to this argument to skip validation of the updated
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateExternalLocation struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force bool `json:"force,omitempty"`

	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`
	// New name for the external location.
	NewName string `json:"new_name,omitempty"`
	// The owner of the external location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the external location.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateExternalLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateExternalLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`
	// Username of current owner of function.
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope UpdateMetastoreDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
	// New name for the metastore.
	NewName string `json:"new_name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMetastoreAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The scope of Delta Sharing enabled for the metastore.
type UpdateMetastoreDeltaSharingScope string

const UpdateMetastoreDeltaSharingScopeInternal UpdateMetastoreDeltaSharingScope = `INTERNAL`

const UpdateMetastoreDeltaSharingScopeInternalAndExternal UpdateMetastoreDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *UpdateMetastoreDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateMetastoreDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = UpdateMetastoreDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns UpdateMetastoreDeltaSharingScope to satisfy [pflag.Value] interface
func (f *UpdateMetastoreDeltaSharingScope) Type() string {
	return "UpdateMetastoreDeltaSharingScope"
}

type UpdateModelVersionRequest struct {
	// The comment attached to the model version
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" url:"-"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMonitor struct {
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// The notification settings for the monitor.
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	OutputSchemaName string `json:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateMonitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMonitor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	Changes []PermissionsChange `json:"changes,omitempty"`
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`
}

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// New name for the registered model.
	NewName string `json:"new_name,omitempty"`
	// The identifier of the user who owns the registered model
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateResponse struct {
}

type UpdateSchema struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Full name of the schema.
	FullName string `json:"-" url:"-"`
	// New name for the schema.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of schema.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// Force update even if there are dependent external locations or external
	// tables.
	Force bool `json:"force,omitempty"`

	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
	// New name for the storage credential.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Update a table owner.
type UpdateTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`

	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" url:"-"`
	// New name for the volume.
	NewName string `json:"new_name,omitempty"`
	// The identifier of the user who owns the volume
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateWorkspaceBindings struct {
	// A list of workspace IDs.
	AssignWorkspaces []int64 `json:"assign_workspaces,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`
	// A list of workspace IDs.
	UnassignWorkspaces []int64 `json:"unassign_workspaces,omitempty"`
}

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings
	Add []WorkspaceBinding `json:"add,omitempty"`
	// List of workspace bindings
	Remove []WorkspaceBinding `json:"remove,omitempty"`
	// The name of the securable.
	SecurableName string `json:"-" url:"-"`
	// The type of the securable to bind to a workspace.
	SecurableType UpdateBindingsSecurableType `json:"-" url:"-"`
}

type ValidateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	CredentialName string `json:"credential_name,omitempty"`
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	ReadOnly bool `json:"read_only,omitempty"`
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidateCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	Results []CredentialValidationResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidateCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A enum represents the result of the file operation
type ValidateCredentialResult string

const ValidateCredentialResultFail ValidateCredentialResult = `FAIL`

const ValidateCredentialResultPass ValidateCredentialResult = `PASS`

const ValidateCredentialResultSkip ValidateCredentialResult = `SKIP`

// String representation for [fmt.Print]
func (f *ValidateCredentialResult) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidateCredentialResult) Set(v string) error {
	switch v {
	case `FAIL`, `PASS`, `SKIP`:
		*f = ValidateCredentialResult(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAIL", "PASS", "SKIP"`, v)
	}
}

// Type always returns ValidateCredentialResult to satisfy [pflag.Value] interface
func (f *ValidateCredentialResult) Type() string {
	return "ValidateCredentialResult"
}

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// The Databricks created GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// The name of an existing external location to validate.
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// The name of the storage credential to validate.
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// The external location url to validate.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	Results []ValidationResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidateStorageCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateStorageCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message string `json:"message,omitempty"`
	// The operation tested.
	Operation ValidationResultOperation `json:"operation,omitempty"`
	// The results of the tested operation.
	Result ValidationResultResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidationResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidationResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The operation tested.
type ValidationResultOperation string

const ValidationResultOperationDelete ValidationResultOperation = `DELETE`

const ValidationResultOperationList ValidationResultOperation = `LIST`

const ValidationResultOperationPathExists ValidationResultOperation = `PATH_EXISTS`

const ValidationResultOperationRead ValidationResultOperation = `READ`

const ValidationResultOperationWrite ValidationResultOperation = `WRITE`

// String representation for [fmt.Print]
func (f *ValidationResultOperation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidationResultOperation) Set(v string) error {
	switch v {
	case `DELETE`, `LIST`, `PATH_EXISTS`, `READ`, `WRITE`:
		*f = ValidationResultOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE", "LIST", "PATH_EXISTS", "READ", "WRITE"`, v)
	}
}

// Type always returns ValidationResultOperation to satisfy [pflag.Value] interface
func (f *ValidationResultOperation) Type() string {
	return "ValidationResultOperation"
}

// The results of the tested operation.
type ValidationResultResult string

const ValidationResultResultFail ValidationResultResult = `FAIL`

const ValidationResultResultPass ValidationResultResult = `PASS`

const ValidationResultResultSkip ValidationResultResult = `SKIP`

// String representation for [fmt.Print]
func (f *ValidationResultResult) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidationResultResult) Set(v string) error {
	switch v {
	case `FAIL`, `PASS`, `SKIP`:
		*f = ValidationResultResult(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAIL", "PASS", "SKIP"`, v)
	}
}

// Type always returns ValidationResultResult to satisfy [pflag.Value] interface
func (f *ValidationResultResult) Type() string {
	return "ValidationResultResult"
}

type VolumeInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the volume
	CreatedBy string `json:"created_by,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// The three-level (fully qualified) name of the volume
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the volume
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the volume
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the volume is
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud
	StorageLocation string `json:"storage_location,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the volume last time
	UpdatedBy string `json:"updated_by,omitempty"`
	// The unique identifier of the volume
	VolumeId string `json:"volume_id,omitempty"`

	VolumeType VolumeType `json:"volume_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VolumeInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VolumeInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VolumeType string

const VolumeTypeExternal VolumeType = `EXTERNAL`

const VolumeTypeManaged VolumeType = `MANAGED`

// String representation for [fmt.Print]
func (f *VolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VolumeType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `MANAGED`:
		*f = VolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "MANAGED"`, v)
	}
}

// Type always returns VolumeType to satisfy [pflag.Value] interface
func (f *VolumeType) Type() string {
	return "VolumeType"
}

type WorkspaceBinding struct {
	BindingType WorkspaceBindingBindingType `json:"binding_type,omitempty"`

	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceBinding) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceBinding) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceBindingBindingType string

const WorkspaceBindingBindingTypeBindingTypeReadOnly WorkspaceBindingBindingType = `BINDING_TYPE_READ_ONLY`

const WorkspaceBindingBindingTypeBindingTypeReadWrite WorkspaceBindingBindingType = `BINDING_TYPE_READ_WRITE`

// String representation for [fmt.Print]
func (f *WorkspaceBindingBindingType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceBindingBindingType) Set(v string) error {
	switch v {
	case `BINDING_TYPE_READ_ONLY`, `BINDING_TYPE_READ_WRITE`:
		*f = WorkspaceBindingBindingType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BINDING_TYPE_READ_ONLY", "BINDING_TYPE_READ_WRITE"`, v)
	}
}

// Type always returns WorkspaceBindingBindingType to satisfy [pflag.Value] interface
func (f *WorkspaceBindingBindingType) Type() string {
	return "WorkspaceBindingBindingType"
}

// Currently assigned workspace bindings
type WorkspaceBindingsResponse struct {
	// List of workspace bindings
	Bindings []WorkspaceBinding `json:"bindings,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceBindingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
