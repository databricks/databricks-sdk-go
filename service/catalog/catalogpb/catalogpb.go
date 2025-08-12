// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalogpb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccountsCreateMetastorePb struct {
	MetastoreInfo *CreateMetastorePb `json:"metastore_info,omitempty"`
}

type AccountsCreateMetastoreAssignmentPb struct {
	MetastoreAssignment *CreateMetastoreAssignmentPb `json:"metastore_assignment,omitempty"`
	MetastoreId         string                       `json:"-" url:"-"`
	WorkspaceId         int64                        `json:"-" url:"-"`
}

type AccountsCreateStorageCredentialPb struct {
	CredentialInfo *CreateStorageCredentialPb `json:"credential_info,omitempty"`
	MetastoreId    string                     `json:"-" url:"-"`
}

type AccountsMetastoreAssignmentPb struct {
	MetastoreAssignment *MetastoreAssignmentPb `json:"metastore_assignment,omitempty"`
}

type AccountsMetastoreInfoPb struct {
	MetastoreInfo *MetastoreInfoPb `json:"metastore_info,omitempty"`
}

type AccountsStorageCredentialInfoPb struct {
	CredentialInfo *StorageCredentialInfoPb `json:"credential_info,omitempty"`
}

type AccountsUpdateMetastorePb struct {
	MetastoreId   string             `json:"-" url:"-"`
	MetastoreInfo *UpdateMetastorePb `json:"metastore_info,omitempty"`
}

type AccountsUpdateMetastoreAssignmentPb struct {
	MetastoreAssignment *UpdateMetastoreAssignmentPb `json:"metastore_assignment,omitempty"`
	MetastoreId         string                       `json:"-" url:"-"`
	WorkspaceId         int64                        `json:"-" url:"-"`
}

type AccountsUpdateStorageCredentialPb struct {
	CredentialInfo        *UpdateStorageCredentialPb `json:"credential_info,omitempty"`
	MetastoreId           string                     `json:"-" url:"-"`
	StorageCredentialName string                     `json:"-" url:"-"`
}

type ArtifactAllowlistInfoPb struct {
	ArtifactMatchers []ArtifactMatcherPb `json:"artifact_matchers,omitempty"`
	CreatedAt        int64               `json:"created_at,omitempty"`
	CreatedBy        string              `json:"created_by,omitempty"`
	MetastoreId      string              `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ArtifactAllowlistInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ArtifactAllowlistInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ArtifactMatcherPb struct {
	Artifact  string      `json:"artifact"`
	MatchType MatchTypePb `json:"match_type"`
}

type ArtifactTypePb string

const ArtifactTypeInitScript ArtifactTypePb = `INIT_SCRIPT`

const ArtifactTypeLibraryJar ArtifactTypePb = `LIBRARY_JAR`

const ArtifactTypeLibraryMaven ArtifactTypePb = `LIBRARY_MAVEN`

type AssignResponsePb struct {
}

type AwsCredentialsPb struct {
	AccessKeyId     string `json:"access_key_id,omitempty"`
	AccessPoint     string `json:"access_point,omitempty"`
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	SessionToken    string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AwsCredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AwsCredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AwsIamRolePb struct {
	ExternalId         string `json:"external_id,omitempty"`
	RoleArn            string `json:"role_arn,omitempty"`
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AwsIamRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AwsIamRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AwsIamRoleRequestPb struct {
	RoleArn string `json:"role_arn"`
}

type AwsIamRoleResponsePb struct {
	ExternalId         string `json:"external_id,omitempty"`
	RoleArn            string `json:"role_arn"`
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AwsIamRoleResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AwsIamRoleResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AwsSqsQueuePb struct {
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	QueueUrl          string `json:"queue_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AwsSqsQueuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AwsSqsQueuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureActiveDirectoryTokenPb struct {
	AadToken string `json:"aad_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AzureActiveDirectoryTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureActiveDirectoryTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureManagedIdentityPb struct {
	AccessConnectorId string `json:"access_connector_id"`
	CredentialId      string `json:"credential_id,omitempty"`
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AzureManagedIdentityPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureManagedIdentityPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureManagedIdentityRequestPb struct {
	AccessConnectorId string `json:"access_connector_id"`
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AzureManagedIdentityRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureManagedIdentityRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureManagedIdentityResponsePb struct {
	AccessConnectorId string `json:"access_connector_id"`
	CredentialId      string `json:"credential_id,omitempty"`
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AzureManagedIdentityResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureManagedIdentityResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureQueueStoragePb struct {
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	QueueUrl          string `json:"queue_url,omitempty"`
	ResourceGroup     string `json:"resource_group,omitempty"`
	SubscriptionId    string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AzureQueueStoragePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureQueueStoragePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureServicePrincipalPb struct {
	ApplicationId string `json:"application_id"`
	ClientSecret  string `json:"client_secret"`
	DirectoryId   string `json:"directory_id"`
}

type AzureUserDelegationSasPb struct {
	SasToken string `json:"sas_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AzureUserDelegationSasPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureUserDelegationSasPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CancelRefreshRequestPb struct {
	RefreshId int64  `json:"-" url:"-"`
	TableName string `json:"-" url:"-"`
}

type CancelRefreshResponsePb struct {
}

type CatalogInfoPb struct {
	BrowseOnly                          bool                                   `json:"browse_only,omitempty"`
	CatalogType                         CatalogTypePb                          `json:"catalog_type,omitempty"`
	Comment                             string                                 `json:"comment,omitempty"`
	ConnectionName                      string                                 `json:"connection_name,omitempty"`
	CreatedAt                           int64                                  `json:"created_at,omitempty"`
	CreatedBy                           string                                 `json:"created_by,omitempty"`
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlagPb `json:"effective_predictive_optimization_flag,omitempty"`
	EnablePredictiveOptimization        EnablePredictiveOptimizationPb         `json:"enable_predictive_optimization,omitempty"`
	FullName                            string                                 `json:"full_name,omitempty"`
	IsolationMode                       CatalogIsolationModePb                 `json:"isolation_mode,omitempty"`
	MetastoreId                         string                                 `json:"metastore_id,omitempty"`
	Name                                string                                 `json:"name,omitempty"`
	Options                             map[string]string                      `json:"options,omitempty"`
	Owner                               string                                 `json:"owner,omitempty"`
	Properties                          map[string]string                      `json:"properties,omitempty"`
	ProviderName                        string                                 `json:"provider_name,omitempty"`
	ProvisioningInfo                    *ProvisioningInfoPb                    `json:"provisioning_info,omitempty"`
	SecurableType                       SecurableTypePb                        `json:"securable_type,omitempty"`
	ShareName                           string                                 `json:"share_name,omitempty"`
	StorageLocation                     string                                 `json:"storage_location,omitempty"`
	StorageRoot                         string                                 `json:"storage_root,omitempty"`
	UpdatedAt                           int64                                  `json:"updated_at,omitempty"`
	UpdatedBy                           string                                 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CatalogInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CatalogInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CatalogIsolationModePb string

const CatalogIsolationModeIsolated CatalogIsolationModePb = `ISOLATED`

const CatalogIsolationModeOpen CatalogIsolationModePb = `OPEN`

type CatalogTypePb string

const CatalogTypeDeltasharingCatalog CatalogTypePb = `DELTASHARING_CATALOG`

const CatalogTypeForeignCatalog CatalogTypePb = `FOREIGN_CATALOG`

const CatalogTypeInternalCatalog CatalogTypePb = `INTERNAL_CATALOG`

const CatalogTypeManagedCatalog CatalogTypePb = `MANAGED_CATALOG`

const CatalogTypeManagedOnlineCatalog CatalogTypePb = `MANAGED_ONLINE_CATALOG`

const CatalogTypeSystemCatalog CatalogTypePb = `SYSTEM_CATALOG`

type CloudflareApiTokenPb struct {
	AccessKeyId     string `json:"access_key_id"`
	AccountId       string `json:"account_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

type ColumnInfoPb struct {
	Comment          string           `json:"comment,omitempty"`
	Mask             *ColumnMaskPb    `json:"mask,omitempty"`
	Name             string           `json:"name,omitempty"`
	Nullable         bool             `json:"nullable,omitempty"`
	PartitionIndex   int              `json:"partition_index,omitempty"`
	Position         int              `json:"position,omitempty"`
	TypeIntervalType string           `json:"type_interval_type,omitempty"`
	TypeJson         string           `json:"type_json,omitempty"`
	TypeName         ColumnTypeNamePb `json:"type_name,omitempty"`
	TypePrecision    int              `json:"type_precision,omitempty"`
	TypeScale        int              `json:"type_scale,omitempty"`
	TypeText         string           `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ColumnInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ColumnInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ColumnMaskPb struct {
	FunctionName     string   `json:"function_name,omitempty"`
	UsingColumnNames []string `json:"using_column_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ColumnMaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ColumnMaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ColumnRelationshipPb struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ColumnRelationshipPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ColumnRelationshipPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ColumnTypeNamePb string

const ColumnTypeNameArray ColumnTypeNamePb = `ARRAY`

const ColumnTypeNameBinary ColumnTypeNamePb = `BINARY`

const ColumnTypeNameBoolean ColumnTypeNamePb = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeNamePb = `BYTE`

const ColumnTypeNameChar ColumnTypeNamePb = `CHAR`

const ColumnTypeNameDate ColumnTypeNamePb = `DATE`

const ColumnTypeNameDecimal ColumnTypeNamePb = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeNamePb = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeNamePb = `FLOAT`

const ColumnTypeNameGeography ColumnTypeNamePb = `GEOGRAPHY`

const ColumnTypeNameGeometry ColumnTypeNamePb = `GEOMETRY`

const ColumnTypeNameInt ColumnTypeNamePb = `INT`

const ColumnTypeNameInterval ColumnTypeNamePb = `INTERVAL`

const ColumnTypeNameLong ColumnTypeNamePb = `LONG`

const ColumnTypeNameMap ColumnTypeNamePb = `MAP`

const ColumnTypeNameNull ColumnTypeNamePb = `NULL`

const ColumnTypeNameShort ColumnTypeNamePb = `SHORT`

const ColumnTypeNameString ColumnTypeNamePb = `STRING`

const ColumnTypeNameStruct ColumnTypeNamePb = `STRUCT`

const ColumnTypeNameTableType ColumnTypeNamePb = `TABLE_TYPE`

const ColumnTypeNameTimestamp ColumnTypeNamePb = `TIMESTAMP`

const ColumnTypeNameTimestampNtz ColumnTypeNamePb = `TIMESTAMP_NTZ`

const ColumnTypeNameUserDefinedType ColumnTypeNamePb = `USER_DEFINED_TYPE`

const ColumnTypeNameVariant ColumnTypeNamePb = `VARIANT`

type ConnectionDependencyPb struct {
	ConnectionName string `json:"connection_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ConnectionDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ConnectionDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ConnectionInfoPb struct {
	Comment             string                 `json:"comment,omitempty"`
	ConnectionId        string                 `json:"connection_id,omitempty"`
	ConnectionType      ConnectionTypePb       `json:"connection_type,omitempty"`
	CreatedAt           int64                  `json:"created_at,omitempty"`
	CreatedBy           string                 `json:"created_by,omitempty"`
	CredentialType      CredentialTypePb       `json:"credential_type,omitempty"`
	EnvironmentSettings *EnvironmentSettingsPb `json:"environment_settings,omitempty"`
	FullName            string                 `json:"full_name,omitempty"`
	MetastoreId         string                 `json:"metastore_id,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Options             map[string]string      `json:"options,omitempty"`
	Owner               string                 `json:"owner,omitempty"`
	Properties          map[string]string      `json:"properties,omitempty"`
	ProvisioningInfo    *ProvisioningInfoPb    `json:"provisioning_info,omitempty"`
	ReadOnly            bool                   `json:"read_only,omitempty"`
	SecurableType       SecurableTypePb        `json:"securable_type,omitempty"`
	UpdatedAt           int64                  `json:"updated_at,omitempty"`
	UpdatedBy           string                 `json:"updated_by,omitempty"`
	Url                 string                 `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ConnectionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ConnectionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ConnectionTypePb string

const ConnectionTypeBigquery ConnectionTypePb = `BIGQUERY`

const ConnectionTypeDatabricks ConnectionTypePb = `DATABRICKS`

const ConnectionTypeGa4RawData ConnectionTypePb = `GA4_RAW_DATA`

const ConnectionTypeGlue ConnectionTypePb = `GLUE`

const ConnectionTypeHiveMetastore ConnectionTypePb = `HIVE_METASTORE`

const ConnectionTypeHttp ConnectionTypePb = `HTTP`

const ConnectionTypeMysql ConnectionTypePb = `MYSQL`

const ConnectionTypeOracle ConnectionTypePb = `ORACLE`

const ConnectionTypePostgresql ConnectionTypePb = `POSTGRESQL`

const ConnectionTypePowerBi ConnectionTypePb = `POWER_BI`

const ConnectionTypeRedshift ConnectionTypePb = `REDSHIFT`

const ConnectionTypeSalesforce ConnectionTypePb = `SALESFORCE`

const ConnectionTypeSalesforceDataCloud ConnectionTypePb = `SALESFORCE_DATA_CLOUD`

const ConnectionTypeServicenow ConnectionTypePb = `SERVICENOW`

const ConnectionTypeSnowflake ConnectionTypePb = `SNOWFLAKE`

const ConnectionTypeSqldw ConnectionTypePb = `SQLDW`

const ConnectionTypeSqlserver ConnectionTypePb = `SQLSERVER`

const ConnectionTypeTeradata ConnectionTypePb = `TERADATA`

const ConnectionTypeUnknownConnectionType ConnectionTypePb = `UNKNOWN_CONNECTION_TYPE`

const ConnectionTypeWorkdayRaas ConnectionTypePb = `WORKDAY_RAAS`

type ContinuousUpdateStatusPb struct {
	InitialPipelineSyncProgress *PipelineProgressPb `json:"initial_pipeline_sync_progress,omitempty"`
	LastProcessedCommitVersion  int64               `json:"last_processed_commit_version,omitempty"`
	Timestamp                   string              `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ContinuousUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ContinuousUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCatalogPb struct {
	Comment        string            `json:"comment,omitempty"`
	ConnectionName string            `json:"connection_name,omitempty"`
	Name           string            `json:"name"`
	Options        map[string]string `json:"options,omitempty"`
	Properties     map[string]string `json:"properties,omitempty"`
	ProviderName   string            `json:"provider_name,omitempty"`
	ShareName      string            `json:"share_name,omitempty"`
	StorageRoot    string            `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateConnectionPb struct {
	Comment             string                 `json:"comment,omitempty"`
	ConnectionType      ConnectionTypePb       `json:"connection_type"`
	EnvironmentSettings *EnvironmentSettingsPb `json:"environment_settings,omitempty"`
	Name                string                 `json:"name"`
	Options             map[string]string      `json:"options"`
	Properties          map[string]string      `json:"properties,omitempty"`
	ReadOnly            bool                   `json:"read_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateConnectionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateConnectionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCredentialRequestPb struct {
	AwsIamRole                  *AwsIamRolePb                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityPb        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipalPb       `json:"azure_service_principal,omitempty"`
	Comment                     string                         `json:"comment,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountPb `json:"databricks_gcp_service_account,omitempty"`
	Name                        string                         `json:"name"`
	Purpose                     CredentialPurposePb            `json:"purpose,omitempty"`
	ReadOnly                    bool                           `json:"read_only,omitempty"`
	SkipValidation              bool                           `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateExternalLineageRelationshipRequestPb struct {
	ExternalLineageRelationship CreateRequestExternalLineagePb `json:"external_lineage_relationship"`
}

type CreateExternalLocationPb struct {
	Comment           string               `json:"comment,omitempty"`
	CredentialName    string               `json:"credential_name"`
	EnableFileEvents  bool                 `json:"enable_file_events,omitempty"`
	EncryptionDetails *EncryptionDetailsPb `json:"encryption_details,omitempty"`
	Fallback          bool                 `json:"fallback,omitempty"`
	FileEventQueue    *FileEventQueuePb    `json:"file_event_queue,omitempty"`
	Name              string               `json:"name"`
	ReadOnly          bool                 `json:"read_only,omitempty"`
	SkipValidation    bool                 `json:"skip_validation,omitempty"`
	Url               string               `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateExternalLocationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateExternalLocationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateExternalMetadataRequestPb struct {
	ExternalMetadata ExternalMetadataPb `json:"external_metadata"`
}

type CreateFunctionPb struct {
	CatalogName         string                         `json:"catalog_name"`
	Comment             string                         `json:"comment,omitempty"`
	DataType            ColumnTypeNamePb               `json:"data_type"`
	ExternalLanguage    string                         `json:"external_language,omitempty"`
	ExternalName        string                         `json:"external_name,omitempty"`
	FullDataType        string                         `json:"full_data_type"`
	InputParams         FunctionParameterInfosPb       `json:"input_params"`
	IsDeterministic     bool                           `json:"is_deterministic"`
	IsNullCall          bool                           `json:"is_null_call"`
	Name                string                         `json:"name"`
	ParameterStyle      CreateFunctionParameterStylePb `json:"parameter_style"`
	Properties          string                         `json:"properties,omitempty"`
	ReturnParams        *FunctionParameterInfosPb      `json:"return_params,omitempty"`
	RoutineBody         CreateFunctionRoutineBodyPb    `json:"routine_body"`
	RoutineDefinition   string                         `json:"routine_definition"`
	RoutineDependencies *DependencyListPb              `json:"routine_dependencies,omitempty"`
	SchemaName          string                         `json:"schema_name"`
	SecurityType        CreateFunctionSecurityTypePb   `json:"security_type"`
	SpecificName        string                         `json:"specific_name"`
	SqlDataAccess       CreateFunctionSqlDataAccessPb  `json:"sql_data_access"`
	SqlPath             string                         `json:"sql_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateFunctionParameterStylePb string

const CreateFunctionParameterStyleS CreateFunctionParameterStylePb = `S`

type CreateFunctionRequestPb struct {
	FunctionInfo CreateFunctionPb `json:"function_info"`
}

type CreateFunctionRoutineBodyPb string

const CreateFunctionRoutineBodyExternal CreateFunctionRoutineBodyPb = `EXTERNAL`

const CreateFunctionRoutineBodySql CreateFunctionRoutineBodyPb = `SQL`

type CreateFunctionSecurityTypePb string

const CreateFunctionSecurityTypeDefiner CreateFunctionSecurityTypePb = `DEFINER`

type CreateFunctionSqlDataAccessPb string

const CreateFunctionSqlDataAccessContainsSql CreateFunctionSqlDataAccessPb = `CONTAINS_SQL`

const CreateFunctionSqlDataAccessNoSql CreateFunctionSqlDataAccessPb = `NO_SQL`

const CreateFunctionSqlDataAccessReadsSqlData CreateFunctionSqlDataAccessPb = `READS_SQL_DATA`

type CreateMetastorePb struct {
	Name        string `json:"name"`
	Region      string `json:"region,omitempty"`
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateMetastorePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateMetastorePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateMetastoreAssignmentPb struct {
	DefaultCatalogName string `json:"default_catalog_name"`
	MetastoreId        string `json:"metastore_id"`
	WorkspaceId        int64  `json:"-" url:"-"`
}

type CreateMonitorPb struct {
	AssetsDir                string                             `json:"assets_dir"`
	BaselineTableName        string                             `json:"baseline_table_name,omitempty"`
	CustomMetrics            []MonitorMetricPb                  `json:"custom_metrics,omitempty"`
	DataClassificationConfig *MonitorDataClassificationConfigPb `json:"data_classification_config,omitempty"`
	InferenceLog             *MonitorInferenceLogPb             `json:"inference_log,omitempty"`
	LatestMonitorFailureMsg  string                             `json:"latest_monitor_failure_msg,omitempty"`
	Notifications            *MonitorNotificationsPb            `json:"notifications,omitempty"`
	OutputSchemaName         string                             `json:"output_schema_name"`
	Schedule                 *MonitorCronSchedulePb             `json:"schedule,omitempty"`
	SkipBuiltinDashboard     bool                               `json:"skip_builtin_dashboard,omitempty"`
	SlicingExprs             []string                           `json:"slicing_exprs,omitempty"`
	Snapshot                 *MonitorSnapshotPb                 `json:"snapshot,omitempty"`
	TableName                string                             `json:"-" url:"-"`
	TimeSeries               *MonitorTimeSeriesPb               `json:"time_series,omitempty"`
	WarehouseId              string                             `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateMonitorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateMonitorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateOnlineTableRequestPb struct {
	Table OnlineTablePb `json:"table"`
}

type CreateRegisteredModelRequestPb struct {
	CatalogName     string `json:"catalog_name"`
	Comment         string `json:"comment,omitempty"`
	Name            string `json:"name"`
	SchemaName      string `json:"schema_name"`
	StorageLocation string `json:"storage_location,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRequestExternalLineagePb struct {
	Columns    []ColumnRelationshipPb  `json:"columns,omitempty"`
	Id         string                  `json:"id,omitempty"`
	Properties map[string]string       `json:"properties,omitempty"`
	Source     ExternalLineageObjectPb `json:"source"`
	Target     ExternalLineageObjectPb `json:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateRequestExternalLineagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateRequestExternalLineagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateSchemaPb struct {
	CatalogName string            `json:"catalog_name"`
	Comment     string            `json:"comment,omitempty"`
	Name        string            `json:"name"`
	Properties  map[string]string `json:"properties,omitempty"`
	StorageRoot string            `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateStorageCredentialPb struct {
	AwsIamRole                  *AwsIamRoleRequestPb                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityRequestPb        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipalPb              `json:"azure_service_principal,omitempty"`
	CloudflareApiToken          *CloudflareApiTokenPb                 `json:"cloudflare_api_token,omitempty"`
	Comment                     string                                `json:"comment,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequestPb `json:"databricks_gcp_service_account,omitempty"`
	Name                        string                                `json:"name"`
	ReadOnly                    bool                                  `json:"read_only,omitempty"`
	SkipValidation              bool                                  `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateTableConstraintPb struct {
	Constraint  TableConstraintPb `json:"constraint"`
	FullNameArg string            `json:"full_name_arg"`
}

type CreateVolumeRequestContentPb struct {
	CatalogName     string       `json:"catalog_name"`
	Comment         string       `json:"comment,omitempty"`
	Name            string       `json:"name"`
	SchemaName      string       `json:"schema_name"`
	StorageLocation string       `json:"storage_location,omitempty"`
	VolumeType      VolumeTypePb `json:"volume_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateVolumeRequestContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateVolumeRequestContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CredentialDependencyPb struct {
	CredentialName string `json:"credential_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CredentialDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CredentialDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CredentialInfoPb struct {
	AwsIamRole                  *AwsIamRolePb                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityPb        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipalPb       `json:"azure_service_principal,omitempty"`
	Comment                     string                         `json:"comment,omitempty"`
	CreatedAt                   int64                          `json:"created_at,omitempty"`
	CreatedBy                   string                         `json:"created_by,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountPb `json:"databricks_gcp_service_account,omitempty"`
	FullName                    string                         `json:"full_name,omitempty"`
	Id                          string                         `json:"id,omitempty"`
	IsolationMode               IsolationModePb                `json:"isolation_mode,omitempty"`
	MetastoreId                 string                         `json:"metastore_id,omitempty"`
	Name                        string                         `json:"name,omitempty"`
	Owner                       string                         `json:"owner,omitempty"`
	Purpose                     CredentialPurposePb            `json:"purpose,omitempty"`
	ReadOnly                    bool                           `json:"read_only,omitempty"`
	UpdatedAt                   int64                          `json:"updated_at,omitempty"`
	UpdatedBy                   string                         `json:"updated_by,omitempty"`
	UsedForManagedStorage       bool                           `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CredentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CredentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CredentialPurposePb string

const CredentialPurposeService CredentialPurposePb = `SERVICE`

const CredentialPurposeStorage CredentialPurposePb = `STORAGE`

type CredentialTypePb string

const CredentialTypeAnyStaticCredential CredentialTypePb = `ANY_STATIC_CREDENTIAL`

const CredentialTypeBearerToken CredentialTypePb = `BEARER_TOKEN`

const CredentialTypeOauthAccessToken CredentialTypePb = `OAUTH_ACCESS_TOKEN`

const CredentialTypeOauthM2m CredentialTypePb = `OAUTH_M2M`

const CredentialTypeOauthRefreshToken CredentialTypePb = `OAUTH_REFRESH_TOKEN`

const CredentialTypeOauthResourceOwnerPassword CredentialTypePb = `OAUTH_RESOURCE_OWNER_PASSWORD`

const CredentialTypeOauthU2m CredentialTypePb = `OAUTH_U2M`

const CredentialTypeOauthU2mMapping CredentialTypePb = `OAUTH_U2M_MAPPING`

const CredentialTypeOidcToken CredentialTypePb = `OIDC_TOKEN`

const CredentialTypePemPrivateKey CredentialTypePb = `PEM_PRIVATE_KEY`

const CredentialTypeServiceCredential CredentialTypePb = `SERVICE_CREDENTIAL`

const CredentialTypeUnknownCredentialType CredentialTypePb = `UNKNOWN_CREDENTIAL_TYPE`

const CredentialTypeUsernamePassword CredentialTypePb = `USERNAME_PASSWORD`

type CredentialValidationResultPb struct {
	Message string                     `json:"message,omitempty"`
	Result  ValidateCredentialResultPb `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CredentialValidationResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CredentialValidationResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CurrentRequestPb struct {
}

type DataSourceFormatPb string

const DataSourceFormatAvro DataSourceFormatPb = `AVRO`

const DataSourceFormatBigqueryFormat DataSourceFormatPb = `BIGQUERY_FORMAT`

const DataSourceFormatCsv DataSourceFormatPb = `CSV`

const DataSourceFormatDatabricksFormat DataSourceFormatPb = `DATABRICKS_FORMAT`

const DataSourceFormatDatabricksRowStoreFormat DataSourceFormatPb = `DATABRICKS_ROW_STORE_FORMAT`

const DataSourceFormatDelta DataSourceFormatPb = `DELTA`

const DataSourceFormatDeltasharing DataSourceFormatPb = `DELTASHARING`

const DataSourceFormatDeltaUniformHudi DataSourceFormatPb = `DELTA_UNIFORM_HUDI`

const DataSourceFormatDeltaUniformIceberg DataSourceFormatPb = `DELTA_UNIFORM_ICEBERG`

const DataSourceFormatHive DataSourceFormatPb = `HIVE`

const DataSourceFormatIceberg DataSourceFormatPb = `ICEBERG`

const DataSourceFormatJson DataSourceFormatPb = `JSON`

const DataSourceFormatMongodbFormat DataSourceFormatPb = `MONGODB_FORMAT`

const DataSourceFormatMysqlFormat DataSourceFormatPb = `MYSQL_FORMAT`

const DataSourceFormatNetsuiteFormat DataSourceFormatPb = `NETSUITE_FORMAT`

const DataSourceFormatOracleFormat DataSourceFormatPb = `ORACLE_FORMAT`

const DataSourceFormatOrc DataSourceFormatPb = `ORC`

const DataSourceFormatParquet DataSourceFormatPb = `PARQUET`

const DataSourceFormatPostgresqlFormat DataSourceFormatPb = `POSTGRESQL_FORMAT`

const DataSourceFormatRedshiftFormat DataSourceFormatPb = `REDSHIFT_FORMAT`

const DataSourceFormatSalesforceDataCloudFormat DataSourceFormatPb = `SALESFORCE_DATA_CLOUD_FORMAT`

const DataSourceFormatSalesforceFormat DataSourceFormatPb = `SALESFORCE_FORMAT`

const DataSourceFormatSnowflakeFormat DataSourceFormatPb = `SNOWFLAKE_FORMAT`

const DataSourceFormatSqldwFormat DataSourceFormatPb = `SQLDW_FORMAT`

const DataSourceFormatSqlserverFormat DataSourceFormatPb = `SQLSERVER_FORMAT`

const DataSourceFormatTeradataFormat DataSourceFormatPb = `TERADATA_FORMAT`

const DataSourceFormatText DataSourceFormatPb = `TEXT`

const DataSourceFormatUnityCatalog DataSourceFormatPb = `UNITY_CATALOG`

const DataSourceFormatVectorIndexFormat DataSourceFormatPb = `VECTOR_INDEX_FORMAT`

const DataSourceFormatWorkdayRaasFormat DataSourceFormatPb = `WORKDAY_RAAS_FORMAT`

type DatabricksGcpServiceAccountPb struct {
	CredentialId string `json:"credential_id,omitempty"`
	Email        string `json:"email,omitempty"`
	PrivateKeyId string `json:"private_key_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabricksGcpServiceAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabricksGcpServiceAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabricksGcpServiceAccountRequestPb struct {
}

type DatabricksGcpServiceAccountResponsePb struct {
	CredentialId string `json:"credential_id,omitempty"`
	Email        string `json:"email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabricksGcpServiceAccountResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabricksGcpServiceAccountResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAccountMetastoreAssignmentRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
	WorkspaceId int64  `json:"-" url:"-"`
}

type DeleteAccountMetastoreRequestPb struct {
	Force       bool   `json:"-" url:"force,omitempty"`
	MetastoreId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteAccountMetastoreRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteAccountMetastoreRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAccountStorageCredentialRequestPb struct {
	Force                 bool   `json:"-" url:"force,omitempty"`
	MetastoreId           string `json:"-" url:"-"`
	StorageCredentialName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteAccountStorageCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteAccountStorageCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAliasRequestPb struct {
	Alias    string `json:"-" url:"-"`
	FullName string `json:"-" url:"-"`
}

type DeleteCatalogRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteCatalogRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteCatalogRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteConnectionRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeleteCredentialRequestPb struct {
	Force   bool   `json:"-" url:"force,omitempty"`
	NameArg string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteCredentialResponsePb struct {
}

type DeleteExternalLineageRelationshipRequestPb struct {
	ExternalLineageRelationship DeleteRequestExternalLineagePb `json:"-" url:"external_lineage_relationship"`
}

type DeleteExternalLocationRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteExternalLocationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteExternalLocationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteExternalMetadataRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeleteFunctionRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteMetastoreRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Id    string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteMetastoreRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteMetastoreRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteModelVersionRequestPb struct {
	FullName string `json:"-" url:"-"`
	Version  int    `json:"-" url:"-"`
}

type DeleteMonitorResponsePb struct {
}

type DeleteOnlineTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeleteQualityMonitorRequestPb struct {
	TableName string `json:"-" url:"-"`
}

type DeleteRegisteredModelRequestPb struct {
	FullName string `json:"-" url:"-"`
}

type DeleteRequestExternalLineagePb struct {
	Id     string                  `json:"id,omitempty" url:"id,omitempty"`
	Source ExternalLineageObjectPb `json:"source" url:"source"`
	Target ExternalLineageObjectPb `json:"target" url:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteRequestExternalLineagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteRequestExternalLineagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteResponsePb struct {
}

type DeleteSchemaRequestPb struct {
	Force    bool   `json:"-" url:"force,omitempty"`
	FullName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteSchemaRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteSchemaRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteStorageCredentialRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteStorageCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteStorageCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteTableConstraintRequestPb struct {
	Cascade        bool   `json:"-" url:"cascade"`
	ConstraintName string `json:"-" url:"constraint_name"`
	FullName       string `json:"-" url:"-"`
}

type DeleteTableConstraintResponsePb struct {
}

type DeleteTableRequestPb struct {
	FullName string `json:"-" url:"-"`
}

type DeleteVolumeRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeltaRuntimePropertiesKvPairsPb struct {
	DeltaRuntimeProperties map[string]string `json:"delta_runtime_properties"`
}

type DeltaSharingScopeEnumPb string

const DeltaSharingScopeEnumInternal DeltaSharingScopeEnumPb = `INTERNAL`

const DeltaSharingScopeEnumInternalAndExternal DeltaSharingScopeEnumPb = `INTERNAL_AND_EXTERNAL`

type DependencyPb struct {
	Connection *ConnectionDependencyPb `json:"connection,omitempty"`
	Credential *CredentialDependencyPb `json:"credential,omitempty"`
	Function   *FunctionDependencyPb   `json:"function,omitempty"`
	Table      *TableDependencyPb      `json:"table,omitempty"`
}

type DependencyListPb struct {
	Dependencies []DependencyPb `json:"dependencies,omitempty"`
}

type DisableRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
	SchemaName  string `json:"-" url:"-"`
}

type DisableResponsePb struct {
}

type EffectivePermissionsListPb struct {
	NextPageToken        string                           `json:"next_page_token,omitempty"`
	PrivilegeAssignments []EffectivePrivilegeAssignmentPb `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EffectivePermissionsListPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EffectivePermissionsListPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EffectivePredictiveOptimizationFlagPb struct {
	InheritedFromName string                                                 `json:"inherited_from_name,omitempty"`
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromTypePb `json:"inherited_from_type,omitempty"`
	Value             EnablePredictiveOptimizationPb                         `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EffectivePredictiveOptimizationFlagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EffectivePredictiveOptimizationFlagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EffectivePredictiveOptimizationFlagInheritedFromTypePb string

const EffectivePredictiveOptimizationFlagInheritedFromTypeCatalog EffectivePredictiveOptimizationFlagInheritedFromTypePb = `CATALOG`

const EffectivePredictiveOptimizationFlagInheritedFromTypeSchema EffectivePredictiveOptimizationFlagInheritedFromTypePb = `SCHEMA`

type EffectivePrivilegePb struct {
	InheritedFromName string          `json:"inherited_from_name,omitempty"`
	InheritedFromType SecurableTypePb `json:"inherited_from_type,omitempty"`
	Privilege         PrivilegePb     `json:"privilege,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EffectivePrivilegePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EffectivePrivilegePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EffectivePrivilegeAssignmentPb struct {
	Principal  string                 `json:"principal,omitempty"`
	Privileges []EffectivePrivilegePb `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EffectivePrivilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EffectivePrivilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnablePredictiveOptimizationPb string

const EnablePredictiveOptimizationDisable EnablePredictiveOptimizationPb = `DISABLE`

const EnablePredictiveOptimizationEnable EnablePredictiveOptimizationPb = `ENABLE`

const EnablePredictiveOptimizationInherit EnablePredictiveOptimizationPb = `INHERIT`

type EnableRequestPb struct {
	CatalogName string `json:"catalog_name,omitempty"`
	MetastoreId string `json:"-" url:"-"`
	SchemaName  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnableResponsePb struct {
}

type EncryptionDetailsPb struct {
	SseEncryptionDetails *SseEncryptionDetailsPb `json:"sse_encryption_details,omitempty"`
}

type EnvironmentSettingsPb struct {
	EnvironmentVersion string   `json:"environment_version,omitempty"`
	JavaDependencies   []string `json:"java_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnvironmentSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnvironmentSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExistsRequestPb struct {
	FullName string `json:"-" url:"-"`
}

type ExternalLineageExternalMetadataPb struct {
	Name string `json:"name,omitempty" url:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageExternalMetadataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageExternalMetadataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageExternalMetadataInfoPb struct {
	EntityType string       `json:"entity_type,omitempty"`
	EventTime  string       `json:"event_time,omitempty"`
	Name       string       `json:"name,omitempty"`
	SystemType SystemTypePb `json:"system_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageExternalMetadataInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageExternalMetadataInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageFileInfoPb struct {
	EventTime       string `json:"event_time,omitempty"`
	Path            string `json:"path,omitempty"`
	SecurableName   string `json:"securable_name,omitempty"`
	SecurableType   string `json:"securable_type,omitempty"`
	StorageLocation string `json:"storage_location,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageFileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageFileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageInfoPb struct {
	ExternalLineageInfo  *ExternalLineageRelationshipInfoPb     `json:"external_lineage_info,omitempty"`
	ExternalMetadataInfo *ExternalLineageExternalMetadataInfoPb `json:"external_metadata_info,omitempty"`
	FileInfo             *ExternalLineageFileInfoPb             `json:"file_info,omitempty"`
	ModelInfo            *ExternalLineageModelVersionInfoPb     `json:"model_info,omitempty"`
	TableInfo            *ExternalLineageTableInfoPb            `json:"table_info,omitempty"`
}

type ExternalLineageModelVersionPb struct {
	Name    string `json:"name,omitempty" url:"name,omitempty"`
	Version string `json:"version,omitempty" url:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageModelVersionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageModelVersionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageModelVersionInfoPb struct {
	EventTime string `json:"event_time,omitempty"`
	ModelName string `json:"model_name,omitempty"`
	Version   int64  `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageModelVersionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageModelVersionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageObjectPb struct {
	ExternalMetadata *ExternalLineageExternalMetadataPb `json:"external_metadata,omitempty" url:"external_metadata,omitempty"`
	ModelVersion     *ExternalLineageModelVersionPb     `json:"model_version,omitempty" url:"model_version,omitempty"`
	Path             *ExternalLineagePathPb             `json:"path,omitempty" url:"path,omitempty"`
	Table            *ExternalLineageTablePb            `json:"table,omitempty" url:"table,omitempty"`
}

type ExternalLineagePathPb struct {
	Url string `json:"url,omitempty" url:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineagePathPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineagePathPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageRelationshipPb struct {
	Columns    []ColumnRelationshipPb  `json:"columns,omitempty"`
	Id         string                  `json:"id,omitempty"`
	Properties map[string]string       `json:"properties,omitempty"`
	Source     ExternalLineageObjectPb `json:"source"`
	Target     ExternalLineageObjectPb `json:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageRelationshipPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageRelationshipPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageRelationshipInfoPb struct {
	Columns    []ColumnRelationshipPb  `json:"columns,omitempty"`
	Id         string                  `json:"id,omitempty"`
	Properties map[string]string       `json:"properties,omitempty"`
	Source     ExternalLineageObjectPb `json:"source"`
	Target     ExternalLineageObjectPb `json:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageRelationshipInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageRelationshipInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageTablePb struct {
	Name string `json:"name,omitempty" url:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLineageTableInfoPb struct {
	CatalogName string `json:"catalog_name,omitempty"`
	EventTime   string `json:"event_time,omitempty"`
	Name        string `json:"name,omitempty"`
	SchemaName  string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLineageTableInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLineageTableInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalLocationInfoPb struct {
	BrowseOnly        bool                 `json:"browse_only,omitempty"`
	Comment           string               `json:"comment,omitempty"`
	CreatedAt         int64                `json:"created_at,omitempty"`
	CreatedBy         string               `json:"created_by,omitempty"`
	CredentialId      string               `json:"credential_id,omitempty"`
	CredentialName    string               `json:"credential_name,omitempty"`
	EnableFileEvents  bool                 `json:"enable_file_events,omitempty"`
	EncryptionDetails *EncryptionDetailsPb `json:"encryption_details,omitempty"`
	Fallback          bool                 `json:"fallback,omitempty"`
	FileEventQueue    *FileEventQueuePb    `json:"file_event_queue,omitempty"`
	IsolationMode     IsolationModePb      `json:"isolation_mode,omitempty"`
	MetastoreId       string               `json:"metastore_id,omitempty"`
	Name              string               `json:"name,omitempty"`
	Owner             string               `json:"owner,omitempty"`
	ReadOnly          bool                 `json:"read_only,omitempty"`
	UpdatedAt         int64                `json:"updated_at,omitempty"`
	UpdatedBy         string               `json:"updated_by,omitempty"`
	Url               string               `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalLocationInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLocationInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalMetadataPb struct {
	Columns     []string          `json:"columns,omitempty"`
	CreateTime  string            `json:"create_time,omitempty"`
	CreatedBy   string            `json:"created_by,omitempty"`
	Description string            `json:"description,omitempty"`
	EntityType  string            `json:"entity_type"`
	Id          string            `json:"id,omitempty"`
	MetastoreId string            `json:"metastore_id,omitempty"`
	Name        string            `json:"name"`
	Owner       string            `json:"owner,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	SystemType  SystemTypePb      `json:"system_type"`
	UpdateTime  string            `json:"update_time,omitempty"`
	UpdatedBy   string            `json:"updated_by,omitempty"`
	Url         string            `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalMetadataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalMetadataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FailedStatusPb struct {
	LastProcessedCommitVersion int64  `json:"last_processed_commit_version,omitempty"`
	Timestamp                  string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FailedStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FailedStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileEventQueuePb struct {
	ManagedAqs     *AzureQueueStoragePb `json:"managed_aqs,omitempty"`
	ManagedPubsub  *GcpPubsubPb         `json:"managed_pubsub,omitempty"`
	ManagedSqs     *AwsSqsQueuePb       `json:"managed_sqs,omitempty"`
	ProvidedAqs    *AzureQueueStoragePb `json:"provided_aqs,omitempty"`
	ProvidedPubsub *GcpPubsubPb         `json:"provided_pubsub,omitempty"`
	ProvidedSqs    *AwsSqsQueuePb       `json:"provided_sqs,omitempty"`
}

type ForeignKeyConstraintPb struct {
	ChildColumns  []string `json:"child_columns"`
	Name          string   `json:"name"`
	ParentColumns []string `json:"parent_columns"`
	ParentTable   string   `json:"parent_table"`
	Rely          bool     `json:"rely,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ForeignKeyConstraintPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ForeignKeyConstraintPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FunctionDependencyPb struct {
	FunctionFullName string `json:"function_full_name"`
}

type FunctionInfoPb struct {
	BrowseOnly          bool                         `json:"browse_only,omitempty"`
	CatalogName         string                       `json:"catalog_name,omitempty"`
	Comment             string                       `json:"comment,omitempty"`
	CreatedAt           int64                        `json:"created_at,omitempty"`
	CreatedBy           string                       `json:"created_by,omitempty"`
	DataType            ColumnTypeNamePb             `json:"data_type,omitempty"`
	ExternalLanguage    string                       `json:"external_language,omitempty"`
	ExternalName        string                       `json:"external_name,omitempty"`
	FullDataType        string                       `json:"full_data_type,omitempty"`
	FullName            string                       `json:"full_name,omitempty"`
	FunctionId          string                       `json:"function_id,omitempty"`
	InputParams         *FunctionParameterInfosPb    `json:"input_params,omitempty"`
	IsDeterministic     bool                         `json:"is_deterministic,omitempty"`
	IsNullCall          bool                         `json:"is_null_call,omitempty"`
	MetastoreId         string                       `json:"metastore_id,omitempty"`
	Name                string                       `json:"name,omitempty"`
	Owner               string                       `json:"owner,omitempty"`
	ParameterStyle      FunctionInfoParameterStylePb `json:"parameter_style,omitempty"`
	Properties          string                       `json:"properties,omitempty"`
	ReturnParams        *FunctionParameterInfosPb    `json:"return_params,omitempty"`
	RoutineBody         FunctionInfoRoutineBodyPb    `json:"routine_body,omitempty"`
	RoutineDefinition   string                       `json:"routine_definition,omitempty"`
	RoutineDependencies *DependencyListPb            `json:"routine_dependencies,omitempty"`
	SchemaName          string                       `json:"schema_name,omitempty"`
	SecurityType        FunctionInfoSecurityTypePb   `json:"security_type,omitempty"`
	SpecificName        string                       `json:"specific_name,omitempty"`
	SqlDataAccess       FunctionInfoSqlDataAccessPb  `json:"sql_data_access,omitempty"`
	SqlPath             string                       `json:"sql_path,omitempty"`
	UpdatedAt           int64                        `json:"updated_at,omitempty"`
	UpdatedBy           string                       `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FunctionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FunctionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FunctionInfoParameterStylePb string

const FunctionInfoParameterStyleS FunctionInfoParameterStylePb = `S`

type FunctionInfoRoutineBodyPb string

const FunctionInfoRoutineBodyExternal FunctionInfoRoutineBodyPb = `EXTERNAL`

const FunctionInfoRoutineBodySql FunctionInfoRoutineBodyPb = `SQL`

type FunctionInfoSecurityTypePb string

const FunctionInfoSecurityTypeDefiner FunctionInfoSecurityTypePb = `DEFINER`

type FunctionInfoSqlDataAccessPb string

const FunctionInfoSqlDataAccessContainsSql FunctionInfoSqlDataAccessPb = `CONTAINS_SQL`

const FunctionInfoSqlDataAccessNoSql FunctionInfoSqlDataAccessPb = `NO_SQL`

const FunctionInfoSqlDataAccessReadsSqlData FunctionInfoSqlDataAccessPb = `READS_SQL_DATA`

type FunctionParameterInfoPb struct {
	Comment          string                  `json:"comment,omitempty"`
	Name             string                  `json:"name"`
	ParameterDefault string                  `json:"parameter_default,omitempty"`
	ParameterMode    FunctionParameterModePb `json:"parameter_mode,omitempty"`
	ParameterType    FunctionParameterTypePb `json:"parameter_type,omitempty"`
	Position         int                     `json:"position"`
	TypeIntervalType string                  `json:"type_interval_type,omitempty"`
	TypeJson         string                  `json:"type_json,omitempty"`
	TypeName         ColumnTypeNamePb        `json:"type_name"`
	TypePrecision    int                     `json:"type_precision,omitempty"`
	TypeScale        int                     `json:"type_scale,omitempty"`
	TypeText         string                  `json:"type_text"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FunctionParameterInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FunctionParameterInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FunctionParameterInfosPb struct {
	Parameters []FunctionParameterInfoPb `json:"parameters,omitempty"`
}

type FunctionParameterModePb string

const FunctionParameterModeIn FunctionParameterModePb = `IN`

type FunctionParameterTypePb string

const FunctionParameterTypeColumn FunctionParameterTypePb = `COLUMN`

const FunctionParameterTypeParam FunctionParameterTypePb = `PARAM`

type GcpOauthTokenPb struct {
	OauthToken string `json:"oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GcpOauthTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GcpOauthTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GcpPubsubPb struct {
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	SubscriptionName  string `json:"subscription_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GcpPubsubPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GcpPubsubPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenerateTemporaryServiceCredentialAzureOptionsPb struct {
	Resources []string `json:"resources,omitempty"`
}

type GenerateTemporaryServiceCredentialGcpOptionsPb struct {
	Scopes []string `json:"scopes,omitempty"`
}

type GenerateTemporaryServiceCredentialRequestPb struct {
	AzureOptions   *GenerateTemporaryServiceCredentialAzureOptionsPb `json:"azure_options,omitempty"`
	CredentialName string                                            `json:"credential_name"`
	GcpOptions     *GenerateTemporaryServiceCredentialGcpOptionsPb   `json:"gcp_options,omitempty"`
}

type GenerateTemporaryTableCredentialRequestPb struct {
	Operation TableOperationPb `json:"operation,omitempty"`
	TableId   string           `json:"table_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenerateTemporaryTableCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenerateTemporaryTableCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenerateTemporaryTableCredentialResponsePb struct {
	AwsTempCredentials     *AwsCredentialsPb            `json:"aws_temp_credentials,omitempty"`
	AzureAad               *AzureActiveDirectoryTokenPb `json:"azure_aad,omitempty"`
	AzureUserDelegationSas *AzureUserDelegationSasPb    `json:"azure_user_delegation_sas,omitempty"`
	ExpirationTime         int64                        `json:"expiration_time,omitempty"`
	GcpOauthToken          *GcpOauthTokenPb             `json:"gcp_oauth_token,omitempty"`
	R2TempCredentials      *R2CredentialsPb             `json:"r2_temp_credentials,omitempty"`
	Url                    string                       `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenerateTemporaryTableCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenerateTemporaryTableCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetAccountMetastoreAssignmentRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

type GetAccountMetastoreRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
}

type GetAccountStorageCredentialRequestPb struct {
	MetastoreId           string `json:"-" url:"-"`
	StorageCredentialName string `json:"-" url:"-"`
}

type GetArtifactAllowlistRequestPb struct {
	ArtifactType ArtifactTypePb `json:"-" url:"-"`
}

type GetBindingsRequestPb struct {
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	SecurableName string `json:"-" url:"-"`
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetBindingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetBindingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetByAliasRequestPb struct {
	Alias          string `json:"-" url:"-"`
	FullName       string `json:"-" url:"-"`
	IncludeAliases bool   `json:"-" url:"include_aliases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetByAliasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetByAliasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetCatalogRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	Name          string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetCatalogRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetCatalogRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetCatalogWorkspaceBindingsResponsePb struct {
	Workspaces []int64 `json:"workspaces,omitempty"`
}

type GetConnectionRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetCredentialRequestPb struct {
	NameArg string `json:"-" url:"-"`
}

type GetEffectiveRequestPb struct {
	FullName      string `json:"-" url:"-"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	Principal     string `json:"-" url:"principal,omitempty"`
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetEffectiveRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetEffectiveRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetExternalLocationRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	Name          string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetExternalLocationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetExternalLocationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetExternalMetadataRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetFunctionRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	Name          string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetGrantRequestPb struct {
	FullName      string `json:"-" url:"-"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	Principal     string `json:"-" url:"principal,omitempty"`
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetGrantRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetGrantRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetMetastoreRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetMetastoreSummaryResponsePb struct {
	Cloud                                       string                  `json:"cloud,omitempty"`
	CreatedAt                                   int64                   `json:"created_at,omitempty"`
	CreatedBy                                   string                  `json:"created_by,omitempty"`
	DefaultDataAccessConfigId                   string                  `json:"default_data_access_config_id,omitempty"`
	DeltaSharingOrganizationName                string                  `json:"delta_sharing_organization_name,omitempty"`
	DeltaSharingRecipientTokenLifetimeInSeconds int64                   `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	DeltaSharingScope                           DeltaSharingScopeEnumPb `json:"delta_sharing_scope,omitempty"`
	ExternalAccessEnabled                       bool                    `json:"external_access_enabled,omitempty"`
	GlobalMetastoreId                           string                  `json:"global_metastore_id,omitempty"`
	MetastoreId                                 string                  `json:"metastore_id,omitempty"`
	Name                                        string                  `json:"name,omitempty"`
	Owner                                       string                  `json:"owner,omitempty"`
	PrivilegeModelVersion                       string                  `json:"privilege_model_version,omitempty"`
	Region                                      string                  `json:"region,omitempty"`
	StorageRoot                                 string                  `json:"storage_root,omitempty"`
	StorageRootCredentialId                     string                  `json:"storage_root_credential_id,omitempty"`
	StorageRootCredentialName                   string                  `json:"storage_root_credential_name,omitempty"`
	UpdatedAt                                   int64                   `json:"updated_at,omitempty"`
	UpdatedBy                                   string                  `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetMetastoreSummaryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetMetastoreSummaryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetModelVersionRequestPb struct {
	FullName       string `json:"-" url:"-"`
	IncludeAliases bool   `json:"-" url:"include_aliases,omitempty"`
	IncludeBrowse  bool   `json:"-" url:"include_browse,omitempty"`
	Version        int    `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetOnlineTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetPermissionsResponsePb struct {
	NextPageToken        string                  `json:"next_page_token,omitempty"`
	PrivilegeAssignments []PrivilegeAssignmentPb `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetQualityMonitorRequestPb struct {
	TableName string `json:"-" url:"-"`
}

type GetQuotaRequestPb struct {
	ParentFullName      string `json:"-" url:"-"`
	ParentSecurableType string `json:"-" url:"-"`
	QuotaName           string `json:"-" url:"-"`
}

type GetQuotaResponsePb struct {
	QuotaInfo *QuotaInfoPb `json:"quota_info,omitempty"`
}

type GetRefreshRequestPb struct {
	RefreshId int64  `json:"-" url:"-"`
	TableName string `json:"-" url:"-"`
}

type GetRegisteredModelRequestPb struct {
	FullName       string `json:"-" url:"-"`
	IncludeAliases bool   `json:"-" url:"include_aliases,omitempty"`
	IncludeBrowse  bool   `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetSchemaRequestPb struct {
	FullName      string `json:"-" url:"-"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetSchemaRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetSchemaRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetStorageCredentialRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetTableRequestPb struct {
	FullName                    string `json:"-" url:"-"`
	IncludeBrowse               bool   `json:"-" url:"include_browse,omitempty"`
	IncludeDeltaMetadata        bool   `json:"-" url:"include_delta_metadata,omitempty"`
	IncludeManifestCapabilities bool   `json:"-" url:"include_manifest_capabilities,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetTableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetTableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetWorkspaceBindingRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetWorkspaceBindingsResponsePb struct {
	Bindings      []WorkspaceBindingPb `json:"bindings,omitempty"`
	NextPageToken string               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetWorkspaceBindingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetWorkspaceBindingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type IsolationModePb string

const IsolationModeIsolationModeIsolated IsolationModePb = `ISOLATION_MODE_ISOLATED`

const IsolationModeIsolationModeOpen IsolationModePb = `ISOLATION_MODE_OPEN`

type LineageDirectionPb string

const LineageDirectionDownstream LineageDirectionPb = `DOWNSTREAM`

const LineageDirectionUpstream LineageDirectionPb = `UPSTREAM`

type ListAccountMetastoreAssignmentsRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
}

type ListAccountMetastoreAssignmentsResponsePb struct {
	WorkspaceIds []int64 `json:"workspace_ids,omitempty"`
}

type ListAccountMetastoresRequestPb struct {
}

type ListAccountStorageCredentialsRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
}

type ListAccountStorageCredentialsResponsePb struct {
	StorageCredentials []StorageCredentialInfoPb `json:"storage_credentials,omitempty"`
}

type ListCatalogsRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCatalogsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCatalogsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCatalogsResponsePb struct {
	Catalogs      []CatalogInfoPb `json:"catalogs,omitempty"`
	NextPageToken string          `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCatalogsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCatalogsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListConnectionsRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListConnectionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListConnectionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListConnectionsResponsePb struct {
	Connections   []ConnectionInfoPb `json:"connections,omitempty"`
	NextPageToken string             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListConnectionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListConnectionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCredentialsRequestPb struct {
	MaxResults int                 `json:"-" url:"max_results,omitempty"`
	PageToken  string              `json:"-" url:"page_token,omitempty"`
	Purpose    CredentialPurposePb `json:"-" url:"purpose,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCredentialsResponsePb struct {
	Credentials   []CredentialInfoPb `json:"credentials,omitempty"`
	NextPageToken string             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExternalLineageRelationshipsRequestPb struct {
	LineageDirection LineageDirectionPb      `json:"-" url:"lineage_direction"`
	ObjectInfo       ExternalLineageObjectPb `json:"-" url:"object_info"`
	PageSize         int                     `json:"-" url:"page_size,omitempty"`
	PageToken        string                  `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExternalLineageRelationshipsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExternalLineageRelationshipsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExternalLineageRelationshipsResponsePb struct {
	ExternalLineageRelationships []ExternalLineageInfoPb `json:"external_lineage_relationships,omitempty"`
	NextPageToken                string                  `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExternalLineageRelationshipsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExternalLineageRelationshipsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExternalLocationsRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExternalLocationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExternalLocationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExternalLocationsResponsePb struct {
	ExternalLocations []ExternalLocationInfoPb `json:"external_locations,omitempty"`
	NextPageToken     string                   `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExternalLocationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExternalLocationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExternalMetadataRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExternalMetadataRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExternalMetadataRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExternalMetadataResponsePb struct {
	ExternalMetadata []ExternalMetadataPb `json:"external_metadata,omitempty"`
	NextPageToken    string               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExternalMetadataResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExternalMetadataResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFunctionsRequestPb struct {
	CatalogName   string `json:"-" url:"catalog_name"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	SchemaName    string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFunctionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFunctionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFunctionsResponsePb struct {
	Functions     []FunctionInfoPb `json:"functions,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFunctionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFunctionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListMetastoresRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListMetastoresRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListMetastoresRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListMetastoresResponsePb struct {
	Metastores    []MetastoreInfoPb `json:"metastores,omitempty"`
	NextPageToken string            `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListMetastoresResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListMetastoresResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListModelVersionsRequestPb struct {
	FullName      string `json:"-" url:"-"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListModelVersionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListModelVersionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListModelVersionsResponsePb struct {
	ModelVersions []ModelVersionInfoPb `json:"model_versions,omitempty"`
	NextPageToken string               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListModelVersionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListModelVersionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQuotasRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQuotasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQuotasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQuotasResponsePb struct {
	NextPageToken string        `json:"next_page_token,omitempty"`
	Quotas        []QuotaInfoPb `json:"quotas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQuotasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQuotasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRefreshesRequestPb struct {
	TableName string `json:"-" url:"-"`
}

type ListRegisteredModelsRequestPb struct {
	CatalogName   string `json:"-" url:"catalog_name,omitempty"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	SchemaName    string `json:"-" url:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListRegisteredModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListRegisteredModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRegisteredModelsResponsePb struct {
	NextPageToken    string                  `json:"next_page_token,omitempty"`
	RegisteredModels []RegisteredModelInfoPb `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListRegisteredModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListRegisteredModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSchemasRequestPb struct {
	CatalogName   string `json:"-" url:"catalog_name"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSchemasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSchemasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSchemasResponsePb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	Schemas       []SchemaInfoPb `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSchemasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSchemasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListStorageCredentialsRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListStorageCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListStorageCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListStorageCredentialsResponsePb struct {
	NextPageToken      string                    `json:"next_page_token,omitempty"`
	StorageCredentials []StorageCredentialInfoPb `json:"storage_credentials,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListStorageCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListStorageCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSummariesRequestPb struct {
	CatalogName                 string `json:"-" url:"catalog_name"`
	IncludeManifestCapabilities bool   `json:"-" url:"include_manifest_capabilities,omitempty"`
	MaxResults                  int    `json:"-" url:"max_results,omitempty"`
	PageToken                   string `json:"-" url:"page_token,omitempty"`
	SchemaNamePattern           string `json:"-" url:"schema_name_pattern,omitempty"`
	TableNamePattern            string `json:"-" url:"table_name_pattern,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSummariesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSummariesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSystemSchemasRequestPb struct {
	MaxResults  int    `json:"-" url:"max_results,omitempty"`
	MetastoreId string `json:"-" url:"-"`
	PageToken   string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSystemSchemasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSystemSchemasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSystemSchemasResponsePb struct {
	NextPageToken string               `json:"next_page_token,omitempty"`
	Schemas       []SystemSchemaInfoPb `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSystemSchemasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSystemSchemasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListTableSummariesResponsePb struct {
	NextPageToken string           `json:"next_page_token,omitempty"`
	Tables        []TableSummaryPb `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListTableSummariesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListTableSummariesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListTablesRequestPb struct {
	CatalogName                 string `json:"-" url:"catalog_name"`
	IncludeBrowse               bool   `json:"-" url:"include_browse,omitempty"`
	IncludeManifestCapabilities bool   `json:"-" url:"include_manifest_capabilities,omitempty"`
	MaxResults                  int    `json:"-" url:"max_results,omitempty"`
	OmitColumns                 bool   `json:"-" url:"omit_columns,omitempty"`
	OmitProperties              bool   `json:"-" url:"omit_properties,omitempty"`
	OmitUsername                bool   `json:"-" url:"omit_username,omitempty"`
	PageToken                   string `json:"-" url:"page_token,omitempty"`
	SchemaName                  string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListTablesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListTablesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListTablesResponsePb struct {
	NextPageToken string        `json:"next_page_token,omitempty"`
	Tables        []TableInfoPb `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListTablesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListTablesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListVolumesRequestPb struct {
	CatalogName   string `json:"-" url:"catalog_name"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	SchemaName    string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListVolumesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListVolumesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListVolumesResponseContentPb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	Volumes       []VolumeInfoPb `json:"volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListVolumesResponseContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListVolumesResponseContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MatchTypePb string

const MatchTypePrefixMatch MatchTypePb = `PREFIX_MATCH`

type MetastoreAssignmentPb struct {
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	MetastoreId        string `json:"metastore_id"`
	WorkspaceId        int64  `json:"workspace_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MetastoreAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MetastoreAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MetastoreInfoPb struct {
	Cloud                                       string                  `json:"cloud,omitempty"`
	CreatedAt                                   int64                   `json:"created_at,omitempty"`
	CreatedBy                                   string                  `json:"created_by,omitempty"`
	DefaultDataAccessConfigId                   string                  `json:"default_data_access_config_id,omitempty"`
	DeltaSharingOrganizationName                string                  `json:"delta_sharing_organization_name,omitempty"`
	DeltaSharingRecipientTokenLifetimeInSeconds int64                   `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	DeltaSharingScope                           DeltaSharingScopeEnumPb `json:"delta_sharing_scope,omitempty"`
	ExternalAccessEnabled                       bool                    `json:"external_access_enabled,omitempty"`
	GlobalMetastoreId                           string                  `json:"global_metastore_id,omitempty"`
	MetastoreId                                 string                  `json:"metastore_id,omitempty"`
	Name                                        string                  `json:"name,omitempty"`
	Owner                                       string                  `json:"owner,omitempty"`
	PrivilegeModelVersion                       string                  `json:"privilege_model_version,omitempty"`
	Region                                      string                  `json:"region,omitempty"`
	StorageRoot                                 string                  `json:"storage_root,omitempty"`
	StorageRootCredentialId                     string                  `json:"storage_root_credential_id,omitempty"`
	StorageRootCredentialName                   string                  `json:"storage_root_credential_name,omitempty"`
	UpdatedAt                                   int64                   `json:"updated_at,omitempty"`
	UpdatedBy                                   string                  `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MetastoreInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MetastoreInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelVersionInfoPb struct {
	Aliases                  []RegisteredModelAliasPb `json:"aliases,omitempty"`
	BrowseOnly               bool                     `json:"browse_only,omitempty"`
	CatalogName              string                   `json:"catalog_name,omitempty"`
	Comment                  string                   `json:"comment,omitempty"`
	CreatedAt                int64                    `json:"created_at,omitempty"`
	CreatedBy                string                   `json:"created_by,omitempty"`
	Id                       string                   `json:"id,omitempty"`
	MetastoreId              string                   `json:"metastore_id,omitempty"`
	ModelName                string                   `json:"model_name,omitempty"`
	ModelVersionDependencies *DependencyListPb        `json:"model_version_dependencies,omitempty"`
	RunId                    string                   `json:"run_id,omitempty"`
	RunWorkspaceId           int                      `json:"run_workspace_id,omitempty"`
	SchemaName               string                   `json:"schema_name,omitempty"`
	Source                   string                   `json:"source,omitempty"`
	Status                   ModelVersionInfoStatusPb `json:"status,omitempty"`
	StorageLocation          string                   `json:"storage_location,omitempty"`
	UpdatedAt                int64                    `json:"updated_at,omitempty"`
	UpdatedBy                string                   `json:"updated_by,omitempty"`
	Version                  int                      `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ModelVersionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ModelVersionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelVersionInfoStatusPb string

const ModelVersionInfoStatusFailedRegistration ModelVersionInfoStatusPb = `FAILED_REGISTRATION`

const ModelVersionInfoStatusPendingRegistration ModelVersionInfoStatusPb = `PENDING_REGISTRATION`

const ModelVersionInfoStatusReady ModelVersionInfoStatusPb = `READY`

type MonitorCronSchedulePb struct {
	PauseStatus          MonitorCronSchedulePauseStatusPb `json:"pause_status,omitempty"`
	QuartzCronExpression string                           `json:"quartz_cron_expression"`
	TimezoneId           string                           `json:"timezone_id"`
}

type MonitorCronSchedulePauseStatusPb string

const MonitorCronSchedulePauseStatusPaused MonitorCronSchedulePauseStatusPb = `PAUSED`

const MonitorCronSchedulePauseStatusUnpaused MonitorCronSchedulePauseStatusPb = `UNPAUSED`

const MonitorCronSchedulePauseStatusUnspecified MonitorCronSchedulePauseStatusPb = `UNSPECIFIED`

type MonitorDataClassificationConfigPb struct {
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MonitorDataClassificationConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MonitorDataClassificationConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MonitorDestinationPb struct {
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

type MonitorInferenceLogPb struct {
	Granularities      []string                         `json:"granularities"`
	LabelCol           string                           `json:"label_col,omitempty"`
	ModelIdCol         string                           `json:"model_id_col"`
	PredictionCol      string                           `json:"prediction_col"`
	PredictionProbaCol string                           `json:"prediction_proba_col,omitempty"`
	ProblemType        MonitorInferenceLogProblemTypePb `json:"problem_type"`
	TimestampCol       string                           `json:"timestamp_col"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MonitorInferenceLogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MonitorInferenceLogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MonitorInferenceLogProblemTypePb string

const MonitorInferenceLogProblemTypeProblemTypeClassification MonitorInferenceLogProblemTypePb = `PROBLEM_TYPE_CLASSIFICATION`

const MonitorInferenceLogProblemTypeProblemTypeRegression MonitorInferenceLogProblemTypePb = `PROBLEM_TYPE_REGRESSION`

type MonitorInfoPb struct {
	AssetsDir                string                             `json:"assets_dir,omitempty"`
	BaselineTableName        string                             `json:"baseline_table_name,omitempty"`
	CustomMetrics            []MonitorMetricPb                  `json:"custom_metrics,omitempty"`
	DashboardId              string                             `json:"dashboard_id,omitempty"`
	DataClassificationConfig *MonitorDataClassificationConfigPb `json:"data_classification_config,omitempty"`
	DriftMetricsTableName    string                             `json:"drift_metrics_table_name"`
	InferenceLog             *MonitorInferenceLogPb             `json:"inference_log,omitempty"`
	LatestMonitorFailureMsg  string                             `json:"latest_monitor_failure_msg,omitempty"`
	MonitorVersion           int64                              `json:"monitor_version"`
	Notifications            *MonitorNotificationsPb            `json:"notifications,omitempty"`
	OutputSchemaName         string                             `json:"output_schema_name"`
	ProfileMetricsTableName  string                             `json:"profile_metrics_table_name"`
	Schedule                 *MonitorCronSchedulePb             `json:"schedule,omitempty"`
	SlicingExprs             []string                           `json:"slicing_exprs,omitempty"`
	Snapshot                 *MonitorSnapshotPb                 `json:"snapshot,omitempty"`
	Status                   MonitorInfoStatusPb                `json:"status"`
	TableName                string                             `json:"table_name"`
	TimeSeries               *MonitorTimeSeriesPb               `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MonitorInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MonitorInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MonitorInfoStatusPb string

const MonitorInfoStatusMonitorStatusActive MonitorInfoStatusPb = `MONITOR_STATUS_ACTIVE`

const MonitorInfoStatusMonitorStatusDeletePending MonitorInfoStatusPb = `MONITOR_STATUS_DELETE_PENDING`

const MonitorInfoStatusMonitorStatusError MonitorInfoStatusPb = `MONITOR_STATUS_ERROR`

const MonitorInfoStatusMonitorStatusFailed MonitorInfoStatusPb = `MONITOR_STATUS_FAILED`

const MonitorInfoStatusMonitorStatusPending MonitorInfoStatusPb = `MONITOR_STATUS_PENDING`

type MonitorMetricPb struct {
	Definition     string              `json:"definition"`
	InputColumns   []string            `json:"input_columns"`
	Name           string              `json:"name"`
	OutputDataType string              `json:"output_data_type"`
	Type           MonitorMetricTypePb `json:"type"`
}

type MonitorMetricTypePb string

const MonitorMetricTypeCustomMetricTypeAggregate MonitorMetricTypePb = `CUSTOM_METRIC_TYPE_AGGREGATE`

const MonitorMetricTypeCustomMetricTypeDerived MonitorMetricTypePb = `CUSTOM_METRIC_TYPE_DERIVED`

const MonitorMetricTypeCustomMetricTypeDrift MonitorMetricTypePb = `CUSTOM_METRIC_TYPE_DRIFT`

type MonitorNotificationsPb struct {
	OnFailure                      *MonitorDestinationPb `json:"on_failure,omitempty"`
	OnNewClassificationTagDetected *MonitorDestinationPb `json:"on_new_classification_tag_detected,omitempty"`
}

type MonitorRefreshInfoPb struct {
	EndTimeMs   int64                       `json:"end_time_ms,omitempty"`
	Message     string                      `json:"message,omitempty"`
	RefreshId   int64                       `json:"refresh_id"`
	StartTimeMs int64                       `json:"start_time_ms"`
	State       MonitorRefreshInfoStatePb   `json:"state"`
	Trigger     MonitorRefreshInfoTriggerPb `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MonitorRefreshInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MonitorRefreshInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MonitorRefreshInfoStatePb string

const MonitorRefreshInfoStateCanceled MonitorRefreshInfoStatePb = `CANCELED`

const MonitorRefreshInfoStateFailed MonitorRefreshInfoStatePb = `FAILED`

const MonitorRefreshInfoStatePending MonitorRefreshInfoStatePb = `PENDING`

const MonitorRefreshInfoStateRunning MonitorRefreshInfoStatePb = `RUNNING`

const MonitorRefreshInfoStateSuccess MonitorRefreshInfoStatePb = `SUCCESS`

const MonitorRefreshInfoStateUnknown MonitorRefreshInfoStatePb = `UNKNOWN`

type MonitorRefreshInfoTriggerPb string

const MonitorRefreshInfoTriggerManual MonitorRefreshInfoTriggerPb = `MANUAL`

const MonitorRefreshInfoTriggerSchedule MonitorRefreshInfoTriggerPb = `SCHEDULE`

const MonitorRefreshInfoTriggerUnknownTrigger MonitorRefreshInfoTriggerPb = `UNKNOWN_TRIGGER`

type MonitorRefreshListResponsePb struct {
	Refreshes []MonitorRefreshInfoPb `json:"refreshes,omitempty"`
}

type MonitorSnapshotPb struct {
}

type MonitorTimeSeriesPb struct {
	Granularities []string `json:"granularities"`
	TimestampCol  string   `json:"timestamp_col"`
}

type NamedTableConstraintPb struct {
	Name string `json:"name"`
}

type OnlineTablePb struct {
	Name                          string                  `json:"name,omitempty"`
	Spec                          *OnlineTableSpecPb      `json:"spec,omitempty"`
	Status                        *OnlineTableStatusPb    `json:"status,omitempty"`
	TableServingUrl               string                  `json:"table_serving_url,omitempty"`
	UnityCatalogProvisioningState ProvisioningInfoStatePb `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OnlineTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OnlineTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OnlineTableSpecPb struct {
	PerformFullCopy     bool                                         `json:"perform_full_copy,omitempty"`
	PipelineId          string                                       `json:"pipeline_id,omitempty"`
	PrimaryKeyColumns   []string                                     `json:"primary_key_columns,omitempty"`
	RunContinuously     *OnlineTableSpecContinuousSchedulingPolicyPb `json:"run_continuously,omitempty"`
	RunTriggered        *OnlineTableSpecTriggeredSchedulingPolicyPb  `json:"run_triggered,omitempty"`
	SourceTableFullName string                                       `json:"source_table_full_name,omitempty"`
	TimeseriesKey       string                                       `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OnlineTableSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OnlineTableSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OnlineTableSpecContinuousSchedulingPolicyPb struct {
}

type OnlineTableSpecTriggeredSchedulingPolicyPb struct {
}

type OnlineTableStatePb string

const OnlineTableStateOffline OnlineTableStatePb = `OFFLINE`

const OnlineTableStateOfflineFailed OnlineTableStatePb = `OFFLINE_FAILED`

const OnlineTableStateOnline OnlineTableStatePb = `ONLINE`

const OnlineTableStateOnlineContinuousUpdate OnlineTableStatePb = `ONLINE_CONTINUOUS_UPDATE`

const OnlineTableStateOnlineNoPendingUpdate OnlineTableStatePb = `ONLINE_NO_PENDING_UPDATE`

const OnlineTableStateOnlinePipelineFailed OnlineTableStatePb = `ONLINE_PIPELINE_FAILED`

const OnlineTableStateOnlineTriggeredUpdate OnlineTableStatePb = `ONLINE_TRIGGERED_UPDATE`

const OnlineTableStateOnlineUpdatingPipelineResources OnlineTableStatePb = `ONLINE_UPDATING_PIPELINE_RESOURCES`

const OnlineTableStateProvisioning OnlineTableStatePb = `PROVISIONING`

const OnlineTableStateProvisioningInitialSnapshot OnlineTableStatePb = `PROVISIONING_INITIAL_SNAPSHOT`

const OnlineTableStateProvisioningPipelineResources OnlineTableStatePb = `PROVISIONING_PIPELINE_RESOURCES`

type OnlineTableStatusPb struct {
	ContinuousUpdateStatus *ContinuousUpdateStatusPb `json:"continuous_update_status,omitempty"`
	DetailedState          OnlineTableStatePb        `json:"detailed_state,omitempty"`
	FailedStatus           *FailedStatusPb           `json:"failed_status,omitempty"`
	Message                string                    `json:"message,omitempty"`
	ProvisioningStatus     *ProvisioningStatusPb     `json:"provisioning_status,omitempty"`
	TriggeredUpdateStatus  *TriggeredUpdateStatusPb  `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OnlineTableStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OnlineTableStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OptionSpecPb struct {
	AllowedValues []string               `json:"allowed_values,omitempty"`
	DefaultValue  string                 `json:"default_value,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Hint          string                 `json:"hint,omitempty"`
	IsCopiable    bool                   `json:"is_copiable,omitempty"`
	IsCreatable   bool                   `json:"is_creatable,omitempty"`
	IsHidden      bool                   `json:"is_hidden,omitempty"`
	IsLoggable    bool                   `json:"is_loggable,omitempty"`
	IsRequired    bool                   `json:"is_required,omitempty"`
	IsSecret      bool                   `json:"is_secret,omitempty"`
	IsUpdatable   bool                   `json:"is_updatable,omitempty"`
	Name          string                 `json:"name,omitempty"`
	OauthStage    OptionSpecOauthStagePb `json:"oauth_stage,omitempty"`
	Type          OptionSpecOptionTypePb `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OptionSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OptionSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OptionSpecOauthStagePb string

const OptionSpecOauthStageBeforeAccessToken OptionSpecOauthStagePb = `BEFORE_ACCESS_TOKEN`

const OptionSpecOauthStageBeforeAuthorizationCode OptionSpecOauthStagePb = `BEFORE_AUTHORIZATION_CODE`

type OptionSpecOptionTypePb string

const OptionSpecOptionTypeOptionBigint OptionSpecOptionTypePb = `OPTION_BIGINT`

const OptionSpecOptionTypeOptionBoolean OptionSpecOptionTypePb = `OPTION_BOOLEAN`

const OptionSpecOptionTypeOptionEnum OptionSpecOptionTypePb = `OPTION_ENUM`

const OptionSpecOptionTypeOptionMultilineString OptionSpecOptionTypePb = `OPTION_MULTILINE_STRING`

const OptionSpecOptionTypeOptionNumber OptionSpecOptionTypePb = `OPTION_NUMBER`

const OptionSpecOptionTypeOptionServiceCredential OptionSpecOptionTypePb = `OPTION_SERVICE_CREDENTIAL`

const OptionSpecOptionTypeOptionString OptionSpecOptionTypePb = `OPTION_STRING`

type PermissionsChangePb struct {
	Add       []PrivilegePb `json:"add,omitempty"`
	Principal string        `json:"principal,omitempty"`
	Remove    []PrivilegePb `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PermissionsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PermissionsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineProgressPb struct {
	EstimatedCompletionTimeSeconds   float64 `json:"estimated_completion_time_seconds,omitempty"`
	LatestVersionCurrentlyProcessing int64   `json:"latest_version_currently_processing,omitempty"`
	SyncProgressCompletion           float64 `json:"sync_progress_completion,omitempty"`
	SyncedRowCount                   int64   `json:"synced_row_count,omitempty"`
	TotalRowCount                    int64   `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineProgressPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineProgressPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PrimaryKeyConstraintPb struct {
	ChildColumns      []string `json:"child_columns"`
	Name              string   `json:"name"`
	Rely              bool     `json:"rely,omitempty"`
	TimeseriesColumns []string `json:"timeseries_columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PrimaryKeyConstraintPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PrimaryKeyConstraintPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PrivilegePb string

const PrivilegeAccess PrivilegePb = `ACCESS`

const PrivilegeAllPrivileges PrivilegePb = `ALL_PRIVILEGES`

const PrivilegeApplyTag PrivilegePb = `APPLY_TAG`

const PrivilegeBrowse PrivilegePb = `BROWSE`

const PrivilegeCreate PrivilegePb = `CREATE`

const PrivilegeCreateCatalog PrivilegePb = `CREATE_CATALOG`

const PrivilegeCreateCleanRoom PrivilegePb = `CREATE_CLEAN_ROOM`

const PrivilegeCreateConnection PrivilegePb = `CREATE_CONNECTION`

const PrivilegeCreateExternalLocation PrivilegePb = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable PrivilegePb = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateExternalVolume PrivilegePb = `CREATE_EXTERNAL_VOLUME`

const PrivilegeCreateForeignCatalog PrivilegePb = `CREATE_FOREIGN_CATALOG`

const PrivilegeCreateForeignSecurable PrivilegePb = `CREATE_FOREIGN_SECURABLE`

const PrivilegeCreateFunction PrivilegePb = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage PrivilegePb = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView PrivilegePb = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateModel PrivilegePb = `CREATE_MODEL`

const PrivilegeCreateProvider PrivilegePb = `CREATE_PROVIDER`

const PrivilegeCreateRecipient PrivilegePb = `CREATE_RECIPIENT`

const PrivilegeCreateSchema PrivilegePb = `CREATE_SCHEMA`

const PrivilegeCreateServiceCredential PrivilegePb = `CREATE_SERVICE_CREDENTIAL`

const PrivilegeCreateShare PrivilegePb = `CREATE_SHARE`

const PrivilegeCreateStorageCredential PrivilegePb = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable PrivilegePb = `CREATE_TABLE`

const PrivilegeCreateView PrivilegePb = `CREATE_VIEW`

const PrivilegeCreateVolume PrivilegePb = `CREATE_VOLUME`

const PrivilegeExecute PrivilegePb = `EXECUTE`

const PrivilegeExecuteCleanRoomTask PrivilegePb = `EXECUTE_CLEAN_ROOM_TASK`

const PrivilegeManage PrivilegePb = `MANAGE`

const PrivilegeManageAllowlist PrivilegePb = `MANAGE_ALLOWLIST`

const PrivilegeModify PrivilegePb = `MODIFY`

const PrivilegeModifyCleanRoom PrivilegePb = `MODIFY_CLEAN_ROOM`

const PrivilegeReadFiles PrivilegePb = `READ_FILES`

const PrivilegeReadPrivateFiles PrivilegePb = `READ_PRIVATE_FILES`

const PrivilegeReadVolume PrivilegePb = `READ_VOLUME`

const PrivilegeRefresh PrivilegePb = `REFRESH`

const PrivilegeSelect PrivilegePb = `SELECT`

const PrivilegeSetSharePermission PrivilegePb = `SET_SHARE_PERMISSION`

const PrivilegeUsage PrivilegePb = `USAGE`

const PrivilegeUseCatalog PrivilegePb = `USE_CATALOG`

const PrivilegeUseConnection PrivilegePb = `USE_CONNECTION`

const PrivilegeUseMarketplaceAssets PrivilegePb = `USE_MARKETPLACE_ASSETS`

const PrivilegeUseProvider PrivilegePb = `USE_PROVIDER`

const PrivilegeUseRecipient PrivilegePb = `USE_RECIPIENT`

const PrivilegeUseSchema PrivilegePb = `USE_SCHEMA`

const PrivilegeUseShare PrivilegePb = `USE_SHARE`

const PrivilegeWriteFiles PrivilegePb = `WRITE_FILES`

const PrivilegeWritePrivateFiles PrivilegePb = `WRITE_PRIVATE_FILES`

const PrivilegeWriteVolume PrivilegePb = `WRITE_VOLUME`

type PrivilegeAssignmentPb struct {
	Principal  string        `json:"principal,omitempty"`
	Privileges []PrivilegePb `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PrivilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PrivilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ProvisioningInfoPb struct {
	State ProvisioningInfoStatePb `json:"state,omitempty"`
}

type ProvisioningInfoStatePb string

const ProvisioningInfoStateActive ProvisioningInfoStatePb = `ACTIVE`

const ProvisioningInfoStateDegraded ProvisioningInfoStatePb = `DEGRADED`

const ProvisioningInfoStateDeleting ProvisioningInfoStatePb = `DELETING`

const ProvisioningInfoStateFailed ProvisioningInfoStatePb = `FAILED`

const ProvisioningInfoStateProvisioning ProvisioningInfoStatePb = `PROVISIONING`

const ProvisioningInfoStateUpdating ProvisioningInfoStatePb = `UPDATING`

type ProvisioningStatusPb struct {
	InitialPipelineSyncProgress *PipelineProgressPb `json:"initial_pipeline_sync_progress,omitempty"`
}

type QuotaInfoPb struct {
	LastRefreshedAt     int64           `json:"last_refreshed_at,omitempty"`
	ParentFullName      string          `json:"parent_full_name,omitempty"`
	ParentSecurableType SecurableTypePb `json:"parent_securable_type,omitempty"`
	QuotaCount          int             `json:"quota_count,omitempty"`
	QuotaLimit          int             `json:"quota_limit,omitempty"`
	QuotaName           string          `json:"quota_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QuotaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QuotaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type R2CredentialsPb struct {
	AccessKeyId     string `json:"access_key_id,omitempty"`
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	SessionToken    string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *R2CredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st R2CredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ReadVolumeRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	Name          string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ReadVolumeRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ReadVolumeRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegenerateDashboardRequestPb struct {
	TableName   string `json:"-" url:"-"`
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegenerateDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegenerateDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegenerateDashboardResponsePb struct {
	DashboardId  string `json:"dashboard_id,omitempty"`
	ParentFolder string `json:"parent_folder,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegenerateDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegenerateDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelAliasPb struct {
	AliasName  string `json:"alias_name,omitempty"`
	VersionNum int    `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegisteredModelAliasPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegisteredModelAliasPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelInfoPb struct {
	Aliases         []RegisteredModelAliasPb `json:"aliases,omitempty"`
	BrowseOnly      bool                     `json:"browse_only,omitempty"`
	CatalogName     string                   `json:"catalog_name,omitempty"`
	Comment         string                   `json:"comment,omitempty"`
	CreatedAt       int64                    `json:"created_at,omitempty"`
	CreatedBy       string                   `json:"created_by,omitempty"`
	FullName        string                   `json:"full_name,omitempty"`
	MetastoreId     string                   `json:"metastore_id,omitempty"`
	Name            string                   `json:"name,omitempty"`
	Owner           string                   `json:"owner,omitempty"`
	SchemaName      string                   `json:"schema_name,omitempty"`
	StorageLocation string                   `json:"storage_location,omitempty"`
	UpdatedAt       int64                    `json:"updated_at,omitempty"`
	UpdatedBy       string                   `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegisteredModelInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegisteredModelInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunRefreshRequestPb struct {
	TableName string `json:"-" url:"-"`
}

type SchemaInfoPb struct {
	BrowseOnly                          bool                                   `json:"browse_only,omitempty"`
	CatalogName                         string                                 `json:"catalog_name,omitempty"`
	CatalogType                         CatalogTypePb                          `json:"catalog_type,omitempty"`
	Comment                             string                                 `json:"comment,omitempty"`
	CreatedAt                           int64                                  `json:"created_at,omitempty"`
	CreatedBy                           string                                 `json:"created_by,omitempty"`
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlagPb `json:"effective_predictive_optimization_flag,omitempty"`
	EnablePredictiveOptimization        EnablePredictiveOptimizationPb         `json:"enable_predictive_optimization,omitempty"`
	FullName                            string                                 `json:"full_name,omitempty"`
	MetastoreId                         string                                 `json:"metastore_id,omitempty"`
	Name                                string                                 `json:"name,omitempty"`
	Owner                               string                                 `json:"owner,omitempty"`
	Properties                          map[string]string                      `json:"properties,omitempty"`
	SchemaId                            string                                 `json:"schema_id,omitempty"`
	StorageLocation                     string                                 `json:"storage_location,omitempty"`
	StorageRoot                         string                                 `json:"storage_root,omitempty"`
	UpdatedAt                           int64                                  `json:"updated_at,omitempty"`
	UpdatedBy                           string                                 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SchemaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SchemaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SecurableKindPb string

const SecurableKindTableDbStorage SecurableKindPb = `TABLE_DB_STORAGE`

const SecurableKindTableDelta SecurableKindPb = `TABLE_DELTA`

const SecurableKindTableDeltasharing SecurableKindPb = `TABLE_DELTASHARING`

const SecurableKindTableDeltasharingMutable SecurableKindPb = `TABLE_DELTASHARING_MUTABLE`

const SecurableKindTableDeltaExternal SecurableKindPb = `TABLE_DELTA_EXTERNAL`

const SecurableKindTableDeltaIcebergDeltasharing SecurableKindPb = `TABLE_DELTA_ICEBERG_DELTASHARING`

const SecurableKindTableDeltaIcebergManaged SecurableKindPb = `TABLE_DELTA_ICEBERG_MANAGED`

const SecurableKindTableDeltaUniformHudiExternal SecurableKindPb = `TABLE_DELTA_UNIFORM_HUDI_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergExternal SecurableKindPb = `TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreExternal SecurableKindPb = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreManaged SecurableKindPb = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_MANAGED`

const SecurableKindTableDeltaUniformIcebergForeignSnowflake SecurableKindPb = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_SNOWFLAKE`

const SecurableKindTableExternal SecurableKindPb = `TABLE_EXTERNAL`

const SecurableKindTableFeatureStore SecurableKindPb = `TABLE_FEATURE_STORE`

const SecurableKindTableFeatureStoreExternal SecurableKindPb = `TABLE_FEATURE_STORE_EXTERNAL`

const SecurableKindTableForeignBigquery SecurableKindPb = `TABLE_FOREIGN_BIGQUERY`

const SecurableKindTableForeignDatabricks SecurableKindPb = `TABLE_FOREIGN_DATABRICKS`

const SecurableKindTableForeignDeltasharing SecurableKindPb = `TABLE_FOREIGN_DELTASHARING`

const SecurableKindTableForeignHiveMetastore SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE`

const SecurableKindTableForeignHiveMetastoreDbfsExternal SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreDbfsManaged SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_MANAGED`

const SecurableKindTableForeignHiveMetastoreDbfsShallowCloneExternal SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreDbfsShallowCloneManaged SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_MANAGED`

const SecurableKindTableForeignHiveMetastoreDbfsView SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_VIEW`

const SecurableKindTableForeignHiveMetastoreExternal SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreManaged SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_MANAGED`

const SecurableKindTableForeignHiveMetastoreShallowCloneExternal SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreShallowCloneManaged SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_MANAGED`

const SecurableKindTableForeignHiveMetastoreView SecurableKindPb = `TABLE_FOREIGN_HIVE_METASTORE_VIEW`

const SecurableKindTableForeignMongodb SecurableKindPb = `TABLE_FOREIGN_MONGODB`

const SecurableKindTableForeignMysql SecurableKindPb = `TABLE_FOREIGN_MYSQL`

const SecurableKindTableForeignNetsuite SecurableKindPb = `TABLE_FOREIGN_NETSUITE`

const SecurableKindTableForeignOracle SecurableKindPb = `TABLE_FOREIGN_ORACLE`

const SecurableKindTableForeignPostgresql SecurableKindPb = `TABLE_FOREIGN_POSTGRESQL`

const SecurableKindTableForeignRedshift SecurableKindPb = `TABLE_FOREIGN_REDSHIFT`

const SecurableKindTableForeignSalesforce SecurableKindPb = `TABLE_FOREIGN_SALESFORCE`

const SecurableKindTableForeignSalesforceDataCloud SecurableKindPb = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD`

const SecurableKindTableForeignSalesforceDataCloudFileSharing SecurableKindPb = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING`

const SecurableKindTableForeignSalesforceDataCloudFileSharingView SecurableKindPb = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING_VIEW`

const SecurableKindTableForeignSnowflake SecurableKindPb = `TABLE_FOREIGN_SNOWFLAKE`

const SecurableKindTableForeignSqldw SecurableKindPb = `TABLE_FOREIGN_SQLDW`

const SecurableKindTableForeignSqlserver SecurableKindPb = `TABLE_FOREIGN_SQLSERVER`

const SecurableKindTableForeignTeradata SecurableKindPb = `TABLE_FOREIGN_TERADATA`

const SecurableKindTableForeignWorkdayRaas SecurableKindPb = `TABLE_FOREIGN_WORKDAY_RAAS`

const SecurableKindTableIcebergUniformManaged SecurableKindPb = `TABLE_ICEBERG_UNIFORM_MANAGED`

const SecurableKindTableInternal SecurableKindPb = `TABLE_INTERNAL`

const SecurableKindTableManagedPostgresql SecurableKindPb = `TABLE_MANAGED_POSTGRESQL`

const SecurableKindTableMaterializedView SecurableKindPb = `TABLE_MATERIALIZED_VIEW`

const SecurableKindTableMaterializedViewDeltasharing SecurableKindPb = `TABLE_MATERIALIZED_VIEW_DELTASHARING`

const SecurableKindTableMetricView SecurableKindPb = `TABLE_METRIC_VIEW`

const SecurableKindTableOnlineVectorIndexDirect SecurableKindPb = `TABLE_ONLINE_VECTOR_INDEX_DIRECT`

const SecurableKindTableOnlineVectorIndexReplica SecurableKindPb = `TABLE_ONLINE_VECTOR_INDEX_REPLICA`

const SecurableKindTableOnlineView SecurableKindPb = `TABLE_ONLINE_VIEW`

const SecurableKindTableStandard SecurableKindPb = `TABLE_STANDARD`

const SecurableKindTableStreamingLiveTable SecurableKindPb = `TABLE_STREAMING_LIVE_TABLE`

const SecurableKindTableStreamingLiveTableDeltasharing SecurableKindPb = `TABLE_STREAMING_LIVE_TABLE_DELTASHARING`

const SecurableKindTableSystem SecurableKindPb = `TABLE_SYSTEM`

const SecurableKindTableSystemDeltasharing SecurableKindPb = `TABLE_SYSTEM_DELTASHARING`

const SecurableKindTableView SecurableKindPb = `TABLE_VIEW`

const SecurableKindTableViewDeltasharing SecurableKindPb = `TABLE_VIEW_DELTASHARING`

type SecurableKindManifestPb struct {
	AssignablePrivileges []string        `json:"assignable_privileges,omitempty"`
	Capabilities         []string        `json:"capabilities,omitempty"`
	Options              []OptionSpecPb  `json:"options,omitempty"`
	SecurableKind        SecurableKindPb `json:"securable_kind,omitempty"`
	SecurableType        SecurableTypePb `json:"securable_type,omitempty"`
}

type SecurableTypePb string

const SecurableTypeCatalog SecurableTypePb = `CATALOG`

const SecurableTypeCleanRoom SecurableTypePb = `CLEAN_ROOM`

const SecurableTypeConnection SecurableTypePb = `CONNECTION`

const SecurableTypeCredential SecurableTypePb = `CREDENTIAL`

const SecurableTypeExternalLocation SecurableTypePb = `EXTERNAL_LOCATION`

const SecurableTypeExternalMetadata SecurableTypePb = `EXTERNAL_METADATA`

const SecurableTypeFunction SecurableTypePb = `FUNCTION`

const SecurableTypeMetastore SecurableTypePb = `METASTORE`

const SecurableTypePipeline SecurableTypePb = `PIPELINE`

const SecurableTypeProvider SecurableTypePb = `PROVIDER`

const SecurableTypeRecipient SecurableTypePb = `RECIPIENT`

const SecurableTypeSchema SecurableTypePb = `SCHEMA`

const SecurableTypeShare SecurableTypePb = `SHARE`

const SecurableTypeStagingTable SecurableTypePb = `STAGING_TABLE`

const SecurableTypeStorageCredential SecurableTypePb = `STORAGE_CREDENTIAL`

const SecurableTypeTable SecurableTypePb = `TABLE`

const SecurableTypeVolume SecurableTypePb = `VOLUME`

type SetArtifactAllowlistPb struct {
	ArtifactMatchers []ArtifactMatcherPb `json:"artifact_matchers"`
	ArtifactType     ArtifactTypePb      `json:"-" url:"-"`
	CreatedAt        int64               `json:"created_at,omitempty"`
	CreatedBy        string              `json:"created_by,omitempty"`
	MetastoreId      string              `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SetArtifactAllowlistPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SetArtifactAllowlistPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetRegisteredModelAliasRequestPb struct {
	Alias      string `json:"alias"`
	FullName   string `json:"full_name"`
	VersionNum int    `json:"version_num"`
}

type SseEncryptionDetailsPb struct {
	Algorithm    SseEncryptionDetailsAlgorithmPb `json:"algorithm,omitempty"`
	AwsKmsKeyArn string                          `json:"aws_kms_key_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SseEncryptionDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SseEncryptionDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SseEncryptionDetailsAlgorithmPb string

const SseEncryptionDetailsAlgorithmAwsSseKms SseEncryptionDetailsAlgorithmPb = `AWS_SSE_KMS`

const SseEncryptionDetailsAlgorithmAwsSseS3 SseEncryptionDetailsAlgorithmPb = `AWS_SSE_S3`

type StorageCredentialInfoPb struct {
	AwsIamRole                  *AwsIamRoleResponsePb                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityResponsePb        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipalPb               `json:"azure_service_principal,omitempty"`
	CloudflareApiToken          *CloudflareApiTokenPb                  `json:"cloudflare_api_token,omitempty"`
	Comment                     string                                 `json:"comment,omitempty"`
	CreatedAt                   int64                                  `json:"created_at,omitempty"`
	CreatedBy                   string                                 `json:"created_by,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponsePb `json:"databricks_gcp_service_account,omitempty"`
	FullName                    string                                 `json:"full_name,omitempty"`
	Id                          string                                 `json:"id,omitempty"`
	IsolationMode               IsolationModePb                        `json:"isolation_mode,omitempty"`
	MetastoreId                 string                                 `json:"metastore_id,omitempty"`
	Name                        string                                 `json:"name,omitempty"`
	Owner                       string                                 `json:"owner,omitempty"`
	ReadOnly                    bool                                   `json:"read_only,omitempty"`
	UpdatedAt                   int64                                  `json:"updated_at,omitempty"`
	UpdatedBy                   string                                 `json:"updated_by,omitempty"`
	UsedForManagedStorage       bool                                   `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StorageCredentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StorageCredentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SummaryRequestPb struct {
}

type SystemSchemaInfoPb struct {
	Schema string `json:"schema"`
	State  string `json:"state"`
}

type SystemTypePb string

const SystemTypeAmazonRedshift SystemTypePb = `AMAZON_REDSHIFT`

const SystemTypeAzureSynapse SystemTypePb = `AZURE_SYNAPSE`

const SystemTypeConfluent SystemTypePb = `CONFLUENT`

const SystemTypeDatabricks SystemTypePb = `DATABRICKS`

const SystemTypeGoogleBigquery SystemTypePb = `GOOGLE_BIGQUERY`

const SystemTypeKafka SystemTypePb = `KAFKA`

const SystemTypeLooker SystemTypePb = `LOOKER`

const SystemTypeMicrosoftFabric SystemTypePb = `MICROSOFT_FABRIC`

const SystemTypeMicrosoftSqlServer SystemTypePb = `MICROSOFT_SQL_SERVER`

const SystemTypeMongodb SystemTypePb = `MONGODB`

const SystemTypeMysql SystemTypePb = `MYSQL`

const SystemTypeOracle SystemTypePb = `ORACLE`

const SystemTypeOther SystemTypePb = `OTHER`

const SystemTypePostgresql SystemTypePb = `POSTGRESQL`

const SystemTypePowerBi SystemTypePb = `POWER_BI`

const SystemTypeSalesforce SystemTypePb = `SALESFORCE`

const SystemTypeSap SystemTypePb = `SAP`

const SystemTypeServicenow SystemTypePb = `SERVICENOW`

const SystemTypeSnowflake SystemTypePb = `SNOWFLAKE`

const SystemTypeTableau SystemTypePb = `TABLEAU`

const SystemTypeTeradata SystemTypePb = `TERADATA`

const SystemTypeWorkday SystemTypePb = `WORKDAY`

type TableConstraintPb struct {
	ForeignKeyConstraint *ForeignKeyConstraintPb `json:"foreign_key_constraint,omitempty"`
	NamedTableConstraint *NamedTableConstraintPb `json:"named_table_constraint,omitempty"`
	PrimaryKeyConstraint *PrimaryKeyConstraintPb `json:"primary_key_constraint,omitempty"`
}

type TableDependencyPb struct {
	TableFullName string `json:"table_full_name"`
}

type TableExistsResponsePb struct {
	TableExists bool `json:"table_exists,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TableExistsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TableExistsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableInfoPb struct {
	AccessPoint                         string                                 `json:"access_point,omitempty"`
	BrowseOnly                          bool                                   `json:"browse_only,omitempty"`
	CatalogName                         string                                 `json:"catalog_name,omitempty"`
	Columns                             []ColumnInfoPb                         `json:"columns,omitempty"`
	Comment                             string                                 `json:"comment,omitempty"`
	CreatedAt                           int64                                  `json:"created_at,omitempty"`
	CreatedBy                           string                                 `json:"created_by,omitempty"`
	DataAccessConfigurationId           string                                 `json:"data_access_configuration_id,omitempty"`
	DataSourceFormat                    DataSourceFormatPb                     `json:"data_source_format,omitempty"`
	DeletedAt                           int64                                  `json:"deleted_at,omitempty"`
	DeltaRuntimePropertiesKvpairs       *DeltaRuntimePropertiesKvPairsPb       `json:"delta_runtime_properties_kvpairs,omitempty"`
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlagPb `json:"effective_predictive_optimization_flag,omitempty"`
	EnablePredictiveOptimization        EnablePredictiveOptimizationPb         `json:"enable_predictive_optimization,omitempty"`
	EncryptionDetails                   *EncryptionDetailsPb                   `json:"encryption_details,omitempty"`
	FullName                            string                                 `json:"full_name,omitempty"`
	MetastoreId                         string                                 `json:"metastore_id,omitempty"`
	Name                                string                                 `json:"name,omitempty"`
	Owner                               string                                 `json:"owner,omitempty"`
	PipelineId                          string                                 `json:"pipeline_id,omitempty"`
	Properties                          map[string]string                      `json:"properties,omitempty"`
	RowFilter                           *TableRowFilterPb                      `json:"row_filter,omitempty"`
	SchemaName                          string                                 `json:"schema_name,omitempty"`
	SecurableKindManifest               *SecurableKindManifestPb               `json:"securable_kind_manifest,omitempty"`
	SqlPath                             string                                 `json:"sql_path,omitempty"`
	StorageCredentialName               string                                 `json:"storage_credential_name,omitempty"`
	StorageLocation                     string                                 `json:"storage_location,omitempty"`
	TableConstraints                    []TableConstraintPb                    `json:"table_constraints,omitempty"`
	TableId                             string                                 `json:"table_id,omitempty"`
	TableType                           TableTypePb                            `json:"table_type,omitempty"`
	UpdatedAt                           int64                                  `json:"updated_at,omitempty"`
	UpdatedBy                           string                                 `json:"updated_by,omitempty"`
	ViewDefinition                      string                                 `json:"view_definition,omitempty"`
	ViewDependencies                    *DependencyListPb                      `json:"view_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TableInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TableInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableOperationPb string

const TableOperationRead TableOperationPb = `READ`

const TableOperationReadWrite TableOperationPb = `READ_WRITE`

type TableRowFilterPb struct {
	FunctionName     string   `json:"function_name"`
	InputColumnNames []string `json:"input_column_names"`
}

type TableSummaryPb struct {
	FullName              string                   `json:"full_name,omitempty"`
	SecurableKindManifest *SecurableKindManifestPb `json:"securable_kind_manifest,omitempty"`
	TableType             TableTypePb              `json:"table_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TableSummaryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TableSummaryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableTypePb string

const TableTypeExternal TableTypePb = `EXTERNAL`

const TableTypeExternalShallowClone TableTypePb = `EXTERNAL_SHALLOW_CLONE`

const TableTypeForeign TableTypePb = `FOREIGN`

const TableTypeManaged TableTypePb = `MANAGED`

const TableTypeManagedShallowClone TableTypePb = `MANAGED_SHALLOW_CLONE`

const TableTypeMaterializedView TableTypePb = `MATERIALIZED_VIEW`

const TableTypeMetricView TableTypePb = `METRIC_VIEW`

const TableTypeStreamingTable TableTypePb = `STREAMING_TABLE`

const TableTypeView TableTypePb = `VIEW`

type TagKeyValuePb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TagKeyValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TagKeyValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TemporaryCredentialsPb struct {
	AwsTempCredentials *AwsCredentialsPb            `json:"aws_temp_credentials,omitempty"`
	AzureAad           *AzureActiveDirectoryTokenPb `json:"azure_aad,omitempty"`
	ExpirationTime     int64                        `json:"expiration_time,omitempty"`
	GcpOauthToken      *GcpOauthTokenPb             `json:"gcp_oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TemporaryCredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TemporaryCredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TriggeredUpdateStatusPb struct {
	LastProcessedCommitVersion int64               `json:"last_processed_commit_version,omitempty"`
	Timestamp                  string              `json:"timestamp,omitempty"`
	TriggeredUpdateProgress    *PipelineProgressPb `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TriggeredUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TriggeredUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UnassignRequestPb struct {
	MetastoreId string `json:"-" url:"metastore_id"`
	WorkspaceId int64  `json:"-" url:"-"`
}

type UnassignResponsePb struct {
}

type UpdateAssignmentResponsePb struct {
}

type UpdateCatalogPb struct {
	Comment                      string                         `json:"comment,omitempty"`
	EnablePredictiveOptimization EnablePredictiveOptimizationPb `json:"enable_predictive_optimization,omitempty"`
	IsolationMode                CatalogIsolationModePb         `json:"isolation_mode,omitempty"`
	Name                         string                         `json:"-" url:"-"`
	NewName                      string                         `json:"new_name,omitempty"`
	Options                      map[string]string              `json:"options,omitempty"`
	Owner                        string                         `json:"owner,omitempty"`
	Properties                   map[string]string              `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCatalogWorkspaceBindingsResponsePb struct {
	Workspaces []int64 `json:"workspaces,omitempty"`
}

type UpdateConnectionPb struct {
	EnvironmentSettings *EnvironmentSettingsPb `json:"environment_settings,omitempty"`
	Name                string                 `json:"-" url:"-"`
	NewName             string                 `json:"new_name,omitempty"`
	Options             map[string]string      `json:"options"`
	Owner               string                 `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateConnectionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateConnectionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCredentialRequestPb struct {
	AwsIamRole                  *AwsIamRolePb                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityPb        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipalPb       `json:"azure_service_principal,omitempty"`
	Comment                     string                         `json:"comment,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountPb `json:"databricks_gcp_service_account,omitempty"`
	Force                       bool                           `json:"force,omitempty"`
	IsolationMode               IsolationModePb                `json:"isolation_mode,omitempty"`
	NameArg                     string                         `json:"-" url:"-"`
	NewName                     string                         `json:"new_name,omitempty"`
	Owner                       string                         `json:"owner,omitempty"`
	ReadOnly                    bool                           `json:"read_only,omitempty"`
	SkipValidation              bool                           `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateExternalLineageRelationshipRequestPb struct {
	ExternalLineageRelationship UpdateRequestExternalLineagePb `json:"external_lineage_relationship"`
	UpdateMask                  string                         `json:"-" url:"update_mask"`
}

type UpdateExternalLocationPb struct {
	Comment           string               `json:"comment,omitempty"`
	CredentialName    string               `json:"credential_name,omitempty"`
	EnableFileEvents  bool                 `json:"enable_file_events,omitempty"`
	EncryptionDetails *EncryptionDetailsPb `json:"encryption_details,omitempty"`
	Fallback          bool                 `json:"fallback,omitempty"`
	FileEventQueue    *FileEventQueuePb    `json:"file_event_queue,omitempty"`
	Force             bool                 `json:"force,omitempty"`
	IsolationMode     IsolationModePb      `json:"isolation_mode,omitempty"`
	Name              string               `json:"-" url:"-"`
	NewName           string               `json:"new_name,omitempty"`
	Owner             string               `json:"owner,omitempty"`
	ReadOnly          bool                 `json:"read_only,omitempty"`
	SkipValidation    bool                 `json:"skip_validation,omitempty"`
	Url               string               `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateExternalLocationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateExternalLocationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateExternalMetadataRequestPb struct {
	ExternalMetadata ExternalMetadataPb `json:"external_metadata"`
	Name             string             `json:"-" url:"-"`
	UpdateMask       string             `json:"-" url:"update_mask"`
}

type UpdateFunctionPb struct {
	Name  string `json:"-" url:"-"`
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateMetastorePb struct {
	DeltaSharingOrganizationName                string                  `json:"delta_sharing_organization_name,omitempty"`
	DeltaSharingRecipientTokenLifetimeInSeconds int64                   `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	DeltaSharingScope                           DeltaSharingScopeEnumPb `json:"delta_sharing_scope,omitempty"`
	Id                                          string                  `json:"-" url:"-"`
	NewName                                     string                  `json:"new_name,omitempty"`
	Owner                                       string                  `json:"owner,omitempty"`
	PrivilegeModelVersion                       string                  `json:"privilege_model_version,omitempty"`
	StorageRootCredentialId                     string                  `json:"storage_root_credential_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateMetastorePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateMetastorePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateMetastoreAssignmentPb struct {
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	MetastoreId        string `json:"metastore_id,omitempty"`
	WorkspaceId        int64  `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateMetastoreAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateMetastoreAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateModelVersionRequestPb struct {
	Comment  string `json:"comment,omitempty"`
	FullName string `json:"-" url:"-"`
	Version  int    `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateMonitorPb struct {
	BaselineTableName        string                             `json:"baseline_table_name,omitempty"`
	CustomMetrics            []MonitorMetricPb                  `json:"custom_metrics,omitempty"`
	DashboardId              string                             `json:"dashboard_id,omitempty"`
	DataClassificationConfig *MonitorDataClassificationConfigPb `json:"data_classification_config,omitempty"`
	InferenceLog             *MonitorInferenceLogPb             `json:"inference_log,omitempty"`
	LatestMonitorFailureMsg  string                             `json:"latest_monitor_failure_msg,omitempty"`
	Notifications            *MonitorNotificationsPb            `json:"notifications,omitempty"`
	OutputSchemaName         string                             `json:"output_schema_name"`
	Schedule                 *MonitorCronSchedulePb             `json:"schedule,omitempty"`
	SlicingExprs             []string                           `json:"slicing_exprs,omitempty"`
	Snapshot                 *MonitorSnapshotPb                 `json:"snapshot,omitempty"`
	TableName                string                             `json:"-" url:"-"`
	TimeSeries               *MonitorTimeSeriesPb               `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateMonitorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateMonitorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdatePermissionsPb struct {
	Changes       []PermissionsChangePb `json:"changes,omitempty"`
	FullName      string                `json:"-" url:"-"`
	SecurableType string                `json:"-" url:"-"`
}

type UpdatePermissionsResponsePb struct {
	PrivilegeAssignments []PrivilegeAssignmentPb `json:"privilege_assignments,omitempty"`
}

type UpdateRegisteredModelRequestPb struct {
	Comment  string `json:"comment,omitempty"`
	FullName string `json:"-" url:"-"`
	NewName  string `json:"new_name,omitempty"`
	Owner    string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRequestExternalLineagePb struct {
	Columns    []ColumnRelationshipPb  `json:"columns,omitempty"`
	Id         string                  `json:"id,omitempty"`
	Properties map[string]string       `json:"properties,omitempty"`
	Source     ExternalLineageObjectPb `json:"source"`
	Target     ExternalLineageObjectPb `json:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateRequestExternalLineagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateRequestExternalLineagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateResponsePb struct {
}

type UpdateSchemaPb struct {
	Comment                      string                         `json:"comment,omitempty"`
	EnablePredictiveOptimization EnablePredictiveOptimizationPb `json:"enable_predictive_optimization,omitempty"`
	FullName                     string                         `json:"-" url:"-"`
	NewName                      string                         `json:"new_name,omitempty"`
	Owner                        string                         `json:"owner,omitempty"`
	Properties                   map[string]string              `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateStorageCredentialPb struct {
	AwsIamRole                  *AwsIamRoleRequestPb                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityResponsePb       `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipalPb              `json:"azure_service_principal,omitempty"`
	CloudflareApiToken          *CloudflareApiTokenPb                 `json:"cloudflare_api_token,omitempty"`
	Comment                     string                                `json:"comment,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequestPb `json:"databricks_gcp_service_account,omitempty"`
	Force                       bool                                  `json:"force,omitempty"`
	IsolationMode               IsolationModePb                       `json:"isolation_mode,omitempty"`
	Name                        string                                `json:"-" url:"-"`
	NewName                     string                                `json:"new_name,omitempty"`
	Owner                       string                                `json:"owner,omitempty"`
	ReadOnly                    bool                                  `json:"read_only,omitempty"`
	SkipValidation              bool                                  `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateTableRequestPb struct {
	FullName string `json:"-" url:"-"`
	Owner    string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateTableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateTableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateVolumeRequestContentPb struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"-" url:"-"`
	NewName string `json:"new_name,omitempty"`
	Owner   string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateVolumeRequestContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateVolumeRequestContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateWorkspaceBindingsPb struct {
	AssignWorkspaces   []int64 `json:"assign_workspaces,omitempty"`
	Name               string  `json:"-" url:"-"`
	UnassignWorkspaces []int64 `json:"unassign_workspaces,omitempty"`
}

type UpdateWorkspaceBindingsParametersPb struct {
	Add           []WorkspaceBindingPb `json:"add,omitempty"`
	Remove        []WorkspaceBindingPb `json:"remove,omitempty"`
	SecurableName string               `json:"-" url:"-"`
	SecurableType string               `json:"-" url:"-"`
}

type UpdateWorkspaceBindingsResponsePb struct {
	Bindings []WorkspaceBindingPb `json:"bindings,omitempty"`
}

type ValidateCredentialRequestPb struct {
	AwsIamRole                  *AwsIamRolePb                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityPb        `json:"azure_managed_identity,omitempty"`
	CredentialName              string                         `json:"credential_name,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountPb `json:"databricks_gcp_service_account,omitempty"`
	ExternalLocationName        string                         `json:"external_location_name,omitempty"`
	Purpose                     CredentialPurposePb            `json:"purpose,omitempty"`
	ReadOnly                    bool                           `json:"read_only,omitempty"`
	Url                         string                         `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ValidateCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ValidateCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ValidateCredentialResponsePb struct {
	IsDir   bool                           `json:"isDir,omitempty"`
	Results []CredentialValidationResultPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ValidateCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ValidateCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ValidateCredentialResultPb string

const ValidateCredentialResultFail ValidateCredentialResultPb = `FAIL`

const ValidateCredentialResultPass ValidateCredentialResultPb = `PASS`

const ValidateCredentialResultSkip ValidateCredentialResultPb = `SKIP`

type ValidateStorageCredentialPb struct {
	AwsIamRole                  *AwsIamRoleRequestPb                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityRequestPb        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipalPb              `json:"azure_service_principal,omitempty"`
	CloudflareApiToken          *CloudflareApiTokenPb                 `json:"cloudflare_api_token,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequestPb `json:"databricks_gcp_service_account,omitempty"`
	ExternalLocationName        string                                `json:"external_location_name,omitempty"`
	ReadOnly                    bool                                  `json:"read_only,omitempty"`
	StorageCredentialName       string                                `json:"storage_credential_name,omitempty"`
	Url                         string                                `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ValidateStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ValidateStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ValidateStorageCredentialResponsePb struct {
	IsDir   bool                 `json:"isDir,omitempty"`
	Results []ValidationResultPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ValidateStorageCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ValidateStorageCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ValidationResultPb struct {
	Message   string                      `json:"message,omitempty"`
	Operation ValidationResultOperationPb `json:"operation,omitempty"`
	Result    ValidationResultResultPb    `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ValidationResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ValidationResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ValidationResultOperationPb string

const ValidationResultOperationDelete ValidationResultOperationPb = `DELETE`

const ValidationResultOperationList ValidationResultOperationPb = `LIST`

const ValidationResultOperationPathExists ValidationResultOperationPb = `PATH_EXISTS`

const ValidationResultOperationRead ValidationResultOperationPb = `READ`

const ValidationResultOperationWrite ValidationResultOperationPb = `WRITE`

type ValidationResultResultPb string

const ValidationResultResultFail ValidationResultResultPb = `FAIL`

const ValidationResultResultPass ValidationResultResultPb = `PASS`

const ValidationResultResultSkip ValidationResultResultPb = `SKIP`

type VolumeInfoPb struct {
	AccessPoint       string               `json:"access_point,omitempty"`
	BrowseOnly        bool                 `json:"browse_only,omitempty"`
	CatalogName       string               `json:"catalog_name,omitempty"`
	Comment           string               `json:"comment,omitempty"`
	CreatedAt         int64                `json:"created_at,omitempty"`
	CreatedBy         string               `json:"created_by,omitempty"`
	EncryptionDetails *EncryptionDetailsPb `json:"encryption_details,omitempty"`
	FullName          string               `json:"full_name,omitempty"`
	MetastoreId       string               `json:"metastore_id,omitempty"`
	Name              string               `json:"name,omitempty"`
	Owner             string               `json:"owner,omitempty"`
	SchemaName        string               `json:"schema_name,omitempty"`
	StorageLocation   string               `json:"storage_location,omitempty"`
	UpdatedAt         int64                `json:"updated_at,omitempty"`
	UpdatedBy         string               `json:"updated_by,omitempty"`
	VolumeId          string               `json:"volume_id,omitempty"`
	VolumeType        VolumeTypePb         `json:"volume_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *VolumeInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st VolumeInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VolumeTypePb string

const VolumeTypeExternal VolumeTypePb = `EXTERNAL`

const VolumeTypeManaged VolumeTypePb = `MANAGED`

type WorkspaceBindingPb struct {
	BindingType WorkspaceBindingBindingTypePb `json:"binding_type,omitempty"`
	WorkspaceId int64                         `json:"workspace_id"`
}

type WorkspaceBindingBindingTypePb string

const WorkspaceBindingBindingTypeBindingTypeReadOnly WorkspaceBindingBindingTypePb = `BINDING_TYPE_READ_ONLY`

const WorkspaceBindingBindingTypeBindingTypeReadWrite WorkspaceBindingBindingTypePb = `BINDING_TYPE_READ_WRITE`
