// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateDatabaseBranchRequest struct {
	// The Database Branch to create.
	DatabaseBranch DatabaseBranch `json:"database_branch"`
	// The ID to use for the Database Branch, which will become the final
	// component of the branch's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	DatabaseBranchId string `json:"-" url:"database_branch_id,omitempty"`
	// The Database Project where this Database Branch will be created. Format:
	// projects/{project_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDatabaseBranchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDatabaseBranchRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDatabaseEndpointRequest struct {
	// The Database Endpoint to create.
	DatabaseEndpoint DatabaseEndpoint `json:"database_endpoint"`
	// The ID to use for the Database Endpoint, which will become the final
	// component of the endpoint's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	DatabaseEndpointId string `json:"-" url:"database_endpoint_id,omitempty"`
	// The Database Branch where this Database Endpoint will be created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDatabaseEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDatabaseEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDatabaseProjectRequest struct {
	// The Database Project to create
	DatabaseProject DatabaseProject `json:"database_project"`
	// The ID to use for the Database Project, which will become the final
	// component of the project's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	DatabaseProjectId string `json:"-" url:"database_project_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDatabaseProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDatabaseProjectRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseBranch struct {
	// A timestamp indicating when the branch was created.
	CreateTime string `json:"create_time,omitempty"`
	// The branch’s state, indicating if it is initializing, ready for use, or
	// archived.
	CurrentState string `json:"current_state,omitempty"`
	// Whether the branch is the project's default branch. This field is only
	// returned on create/update responses. See effective_default for the value
	// that is actually applied to the database branch.
	Default bool `json:"default,omitempty"`
	// Whether the branch is the project's default branch.
	EffectiveDefault bool `json:"effective_default,omitempty"`
	// Whether the branch is protected.
	IsProtected bool `json:"is_protected,omitempty"`
	// The logical size of the branch.
	LogicalSizeBytes int64 `json:"logical_size_bytes,omitempty"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"name,omitempty"`
	// The parent to list branches from. Format: projects/{project_id}
	Parent string `json:"parent,omitempty"`
	// The parent branch Format: projects/{project_id}/branches/{branch_id}
	ParentBranch string `json:"parent_branch,omitempty"`
	// The Log Sequence Number (LSN) on the parent branch from which this branch
	// was created. When restoring a branch using the Restore Database Branch
	// endpoint, this value isn’t finalized until all operations related to
	// the restore have completed successfully.
	ParentBranchLsn string `json:"parent_branch_lsn,omitempty"`
	// The point in time on the parent branch from which this branch was
	// created.
	ParentBranchTime string `json:"parent_branch_time,omitempty"`

	PendingState string `json:"pending_state,omitempty"`
	// A timestamp indicating when the `current_state` began.
	StateChangeTime string `json:"state_change_time,omitempty"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseBranch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseBranch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseBranchOperationMetadata struct {
}

type DatabaseEndpoint struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// A timestamp indicating when the compute endpoint was created.
	CreateTime string `json:"create_time,omitempty"`

	CurrentState DatabaseEndpointState `json:"current_state,omitempty"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled bool `json:"disabled,omitempty"`
	// The hostname of the compute endpoint. This is the hostname specified when
	// connecting to a database.
	Host string `json:"host,omitempty"`
	// A timestamp indicating when the compute endpoint was last active.
	LastActiveTime string `json:"last_active_time,omitempty"`
	// The resource name of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"name,omitempty"`
	// The parent to list endpoints from. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`

	PendingState DatabaseEndpointState `json:"pending_state,omitempty"`

	PoolerMode DatabaseEndpointPoolerMode `json:"pooler_mode,omitempty"`

	Settings *DatabaseEndpointSettings `json:"settings,omitempty"`
	// A timestamp indicating when the compute endpoint was last started.
	StartTime string `json:"start_time,omitempty"`
	// A timestamp indicating when the compute endpoint was last suspended.
	SuspendTime string `json:"suspend_time,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration string `json:"suspend_timeout_duration,omitempty"`
	// NOTE: if want type to default to some value set the server then an
	// effective_type field OR make this field REQUIRED
	Type DatabaseEndpointType `json:"type,omitempty"`
	// A timestamp indicating when the compute endpoint was last updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseEndpointOperationMetadata struct {
}

// The connection pooler mode. Lakebase supports PgBouncer in `transaction` mode
// only.
type DatabaseEndpointPoolerMode string

const DatabaseEndpointPoolerModeTransaction DatabaseEndpointPoolerMode = `TRANSACTION`

// String representation for [fmt.Print]
func (f *DatabaseEndpointPoolerMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseEndpointPoolerMode) Set(v string) error {
	switch v {
	case `TRANSACTION`:
		*f = DatabaseEndpointPoolerMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TRANSACTION"`, v)
	}
}

// Values returns all possible values for DatabaseEndpointPoolerMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseEndpointPoolerMode) Values() []DatabaseEndpointPoolerMode {
	return []DatabaseEndpointPoolerMode{
		DatabaseEndpointPoolerModeTransaction,
	}
}

// Type always returns DatabaseEndpointPoolerMode to satisfy [pflag.Value] interface
func (f *DatabaseEndpointPoolerMode) Type() string {
	return "DatabaseEndpointPoolerMode"
}

// A collection of settings for a compute endpoint
type DatabaseEndpointSettings struct {
	// A raw representation of Postgres settings.
	PgSettings map[string]string `json:"pg_settings,omitempty"`
	// A raw representation of PgBouncer settings.
	PgbouncerSettings map[string]string `json:"pgbouncer_settings,omitempty"`
}

// The state of the compute endpoint
type DatabaseEndpointState string

const DatabaseEndpointStateActive DatabaseEndpointState = `ACTIVE`

const DatabaseEndpointStateIdle DatabaseEndpointState = `IDLE`

const DatabaseEndpointStateInit DatabaseEndpointState = `INIT`

// String representation for [fmt.Print]
func (f *DatabaseEndpointState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseEndpointState) Set(v string) error {
	switch v {
	case `ACTIVE`, `IDLE`, `INIT`:
		*f = DatabaseEndpointState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "IDLE", "INIT"`, v)
	}
}

// Values returns all possible values for DatabaseEndpointState.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseEndpointState) Values() []DatabaseEndpointState {
	return []DatabaseEndpointState{
		DatabaseEndpointStateActive,
		DatabaseEndpointStateIdle,
		DatabaseEndpointStateInit,
	}
}

// Type always returns DatabaseEndpointState to satisfy [pflag.Value] interface
func (f *DatabaseEndpointState) Type() string {
	return "DatabaseEndpointState"
}

// The compute endpoint type. Either `read_write` or `read_only`.
type DatabaseEndpointType string

const DatabaseEndpointTypeReadOnly DatabaseEndpointType = `READ_ONLY`

const DatabaseEndpointTypeReadWrite DatabaseEndpointType = `READ_WRITE`

// String representation for [fmt.Print]
func (f *DatabaseEndpointType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseEndpointType) Set(v string) error {
	switch v {
	case `READ_ONLY`, `READ_WRITE`:
		*f = DatabaseEndpointType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ_ONLY", "READ_WRITE"`, v)
	}
}

// Values returns all possible values for DatabaseEndpointType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseEndpointType) Values() []DatabaseEndpointType {
	return []DatabaseEndpointType{
		DatabaseEndpointTypeReadOnly,
		DatabaseEndpointTypeReadWrite,
	}
}

// Type always returns DatabaseEndpointType to satisfy [pflag.Value] interface
func (f *DatabaseEndpointType) Type() string {
	return "DatabaseEndpointType"
}

type DatabaseProject struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes int64 `json:"branch_logical_size_limit_bytes,omitempty"`
	// The desired budget policy to associate with the instance. This field is
	// only returned on create/update responses, and represents the customer
	// provided budget policy. See effective_budget_policy_id for the policy
	// that is actually applied to the instance.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The most recent time when any endpoint of this project was active.
	ComputeLastActiveTime string `json:"compute_last_active_time,omitempty"`
	// A timestamp indicating when the project was created.
	CreateTime string `json:"create_time,omitempty"`
	// Custom tags associated with the instance.
	CustomTags []DatabaseProjectCustomTag `json:"custom_tags,omitempty"`

	DefaultEndpointSettings *DatabaseProjectDefaultEndpointSettings `json:"default_endpoint_settings,omitempty"`
	// Human-readable project name.
	DisplayName string `json:"display_name,omitempty"`
	// The policy that is applied to the instance.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	EffectiveDefaultEndpointSettings *DatabaseProjectDefaultEndpointSettings `json:"effective_default_endpoint_settings,omitempty"`

	EffectiveDisplayName string `json:"effective_display_name,omitempty"`

	EffectiveHistoryRetentionDuration string `json:"effective_history_retention_duration,omitempty"`

	EffectivePgVersion int `json:"effective_pg_version,omitempty"`

	EffectiveSettings *DatabaseProjectSettings `json:"effective_settings,omitempty"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project.
	HistoryRetentionDuration string `json:"history_retention_duration,omitempty"`
	// The resource name of the project. Format: projects/{project_id}
	Name string `json:"name,omitempty"`
	// The major Postgres version number.
	PgVersion int `json:"pg_version,omitempty"`

	Settings *DatabaseProjectSettings `json:"settings,omitempty"`
	// The current space occupied by the project in storage. Synthetic storage
	// size combines the logical data size and Write-Ahead Log (WAL) size for
	// all branches in a project.
	SyntheticStorageSizeBytes int64 `json:"synthetic_storage_size_bytes,omitempty"`
	// A timestamp indicating when the project was last updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseProject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseProject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseProjectCustomTag struct {
	// The key of the custom tag.
	Key string `json:"key,omitempty"`
	// The value of the custom tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseProjectCustomTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseProjectCustomTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A collection of settings for a database endpoint.
type DatabaseProjectDefaultEndpointSettings struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// A raw representation of Postgres settings.
	PgSettings map[string]string `json:"pg_settings,omitempty"`
	// A raw representation of PgBouncer settings.
	PgbouncerSettings map[string]string `json:"pgbouncer_settings,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration string `json:"suspend_timeout_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseProjectDefaultEndpointSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseProjectDefaultEndpointSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseProjectOperationMetadata struct {
}

type DatabaseProjectSettings struct {
	// Sets wal_level=logical for all compute endpoints in this project. All
	// active endpoints will be suspended. Once enabled, logical replication
	// cannot be disabled.
	EnableLogicalReplication bool `json:"enable_logical_replication,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseProjectSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseProjectSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	// @pbjson-skip
	Details []json.RawMessage `json:"details,omitempty"`

	ErrorCode ErrorCode `json:"error_code,omitempty"`

	Message string `json:"message,omitempty"`

	StackTrace string `json:"stack_trace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksServiceExceptionWithDetailsProto) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksServiceExceptionWithDetailsProto) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDatabaseBranchRequest struct {
	// The name of the Database Branch to delete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
}

type DeleteDatabaseEndpointRequest struct {
	// The name of the Database Endpoint to delete. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

type DeleteDatabaseProjectRequest struct {
	// The name of the Database Project to delete. Format: projects/{project_id}
	Name string `json:"-" url:"-"`
}

// Legacy definition of the ErrorCode enum. Please keep in sync with
// api-base/proto/error_code.proto (except status code mapping annotations as
// this file doesn't have them). Will be removed eventually, pending the ScalaPB
// 0.4 cleanup.
type ErrorCode string

const ErrorCodeAborted ErrorCode = `ABORTED`

const ErrorCodeAlreadyExists ErrorCode = `ALREADY_EXISTS`

const ErrorCodeBadRequest ErrorCode = `BAD_REQUEST`

const ErrorCodeCancelled ErrorCode = `CANCELLED`

const ErrorCodeCatalogAlreadyExists ErrorCode = `CATALOG_ALREADY_EXISTS`

const ErrorCodeCatalogDoesNotExist ErrorCode = `CATALOG_DOES_NOT_EXIST`

const ErrorCodeCatalogNotEmpty ErrorCode = `CATALOG_NOT_EMPTY`

const ErrorCodeCouldNotAcquireLock ErrorCode = `COULD_NOT_ACQUIRE_LOCK`

const ErrorCodeCustomerUnauthorized ErrorCode = `CUSTOMER_UNAUTHORIZED`

const ErrorCodeDacAlreadyExists ErrorCode = `DAC_ALREADY_EXISTS`

const ErrorCodeDacDoesNotExist ErrorCode = `DAC_DOES_NOT_EXIST`

const ErrorCodeDataLoss ErrorCode = `DATA_LOSS`

const ErrorCodeDeadlineExceeded ErrorCode = `DEADLINE_EXCEEDED`

const ErrorCodeDeploymentTimeout ErrorCode = `DEPLOYMENT_TIMEOUT`

const ErrorCodeDirectoryNotEmpty ErrorCode = `DIRECTORY_NOT_EMPTY`

const ErrorCodeDirectoryProtected ErrorCode = `DIRECTORY_PROTECTED`

const ErrorCodeDryRunFailed ErrorCode = `DRY_RUN_FAILED`

const ErrorCodeEndpointNotFound ErrorCode = `ENDPOINT_NOT_FOUND`

const ErrorCodeExternalLocationAlreadyExists ErrorCode = `EXTERNAL_LOCATION_ALREADY_EXISTS`

const ErrorCodeExternalLocationDoesNotExist ErrorCode = `EXTERNAL_LOCATION_DOES_NOT_EXIST`

const ErrorCodeFeatureDisabled ErrorCode = `FEATURE_DISABLED`

const ErrorCodeGitConflict ErrorCode = `GIT_CONFLICT`

const ErrorCodeGitRemoteError ErrorCode = `GIT_REMOTE_ERROR`

const ErrorCodeGitSensitiveTokenDetected ErrorCode = `GIT_SENSITIVE_TOKEN_DETECTED`

const ErrorCodeGitUnknownRef ErrorCode = `GIT_UNKNOWN_REF`

const ErrorCodeGitUrlNotOnAllowList ErrorCode = `GIT_URL_NOT_ON_ALLOW_LIST`

const ErrorCodeInsecurePartnerResponse ErrorCode = `INSECURE_PARTNER_RESPONSE`

const ErrorCodeInternalError ErrorCode = `INTERNAL_ERROR`

const ErrorCodeInvalidParameterValue ErrorCode = `INVALID_PARAMETER_VALUE`

const ErrorCodeInvalidState ErrorCode = `INVALID_STATE`

const ErrorCodeInvalidStateTransition ErrorCode = `INVALID_STATE_TRANSITION`

const ErrorCodeIoError ErrorCode = `IO_ERROR`

const ErrorCodeIpynbFileInRepo ErrorCode = `IPYNB_FILE_IN_REPO`

const ErrorCodeMalformedPartnerResponse ErrorCode = `MALFORMED_PARTNER_RESPONSE`

const ErrorCodeMalformedRequest ErrorCode = `MALFORMED_REQUEST`

const ErrorCodeManagedResourceGroupDoesNotExist ErrorCode = `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`

const ErrorCodeMaxBlockSizeExceeded ErrorCode = `MAX_BLOCK_SIZE_EXCEEDED`

const ErrorCodeMaxChildNodeSizeExceeded ErrorCode = `MAX_CHILD_NODE_SIZE_EXCEEDED`

const ErrorCodeMaxListSizeExceeded ErrorCode = `MAX_LIST_SIZE_EXCEEDED`

const ErrorCodeMaxNotebookSizeExceeded ErrorCode = `MAX_NOTEBOOK_SIZE_EXCEEDED`

const ErrorCodeMaxReadSizeExceeded ErrorCode = `MAX_READ_SIZE_EXCEEDED`

const ErrorCodeMetastoreAlreadyExists ErrorCode = `METASTORE_ALREADY_EXISTS`

const ErrorCodeMetastoreDoesNotExist ErrorCode = `METASTORE_DOES_NOT_EXIST`

const ErrorCodeMetastoreNotEmpty ErrorCode = `METASTORE_NOT_EMPTY`

const ErrorCodeNotFound ErrorCode = `NOT_FOUND`

const ErrorCodeNotImplemented ErrorCode = `NOT_IMPLEMENTED`

const ErrorCodePartialDelete ErrorCode = `PARTIAL_DELETE`

const ErrorCodePermissionDenied ErrorCode = `PERMISSION_DENIED`

const ErrorCodePermissionNotPropagated ErrorCode = `PERMISSION_NOT_PROPAGATED`

const ErrorCodePrincipalDoesNotExist ErrorCode = `PRINCIPAL_DOES_NOT_EXIST`

const ErrorCodeProjectsOperationTimeout ErrorCode = `PROJECTS_OPERATION_TIMEOUT`

const ErrorCodeProviderAlreadyExists ErrorCode = `PROVIDER_ALREADY_EXISTS`

const ErrorCodeProviderDoesNotExist ErrorCode = `PROVIDER_DOES_NOT_EXIST`

const ErrorCodeProviderShareNotAccessible ErrorCode = `PROVIDER_SHARE_NOT_ACCESSIBLE`

const ErrorCodeQuotaExceeded ErrorCode = `QUOTA_EXCEEDED`

const ErrorCodeRecipientAlreadyExists ErrorCode = `RECIPIENT_ALREADY_EXISTS`

const ErrorCodeRecipientDoesNotExist ErrorCode = `RECIPIENT_DOES_NOT_EXIST`

const ErrorCodeRequestLimitExceeded ErrorCode = `REQUEST_LIMIT_EXCEEDED`

const ErrorCodeResourceAlreadyExists ErrorCode = `RESOURCE_ALREADY_EXISTS`

const ErrorCodeResourceConflict ErrorCode = `RESOURCE_CONFLICT`

const ErrorCodeResourceDoesNotExist ErrorCode = `RESOURCE_DOES_NOT_EXIST`

const ErrorCodeResourceExhausted ErrorCode = `RESOURCE_EXHAUSTED`

const ErrorCodeResourceLimitExceeded ErrorCode = `RESOURCE_LIMIT_EXCEEDED`

const ErrorCodeSchemaAlreadyExists ErrorCode = `SCHEMA_ALREADY_EXISTS`

const ErrorCodeSchemaDoesNotExist ErrorCode = `SCHEMA_DOES_NOT_EXIST`

const ErrorCodeSchemaNotEmpty ErrorCode = `SCHEMA_NOT_EMPTY`

const ErrorCodeSearchQueryTooLong ErrorCode = `SEARCH_QUERY_TOO_LONG`

const ErrorCodeSearchQueryTooShort ErrorCode = `SEARCH_QUERY_TOO_SHORT`

const ErrorCodeServiceUnderMaintenance ErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ErrorCodeShareAlreadyExists ErrorCode = `SHARE_ALREADY_EXISTS`

const ErrorCodeShareDoesNotExist ErrorCode = `SHARE_DOES_NOT_EXIST`

const ErrorCodeStorageCredentialAlreadyExists ErrorCode = `STORAGE_CREDENTIAL_ALREADY_EXISTS`

const ErrorCodeStorageCredentialDoesNotExist ErrorCode = `STORAGE_CREDENTIAL_DOES_NOT_EXIST`

const ErrorCodeTableAlreadyExists ErrorCode = `TABLE_ALREADY_EXISTS`

const ErrorCodeTableDoesNotExist ErrorCode = `TABLE_DOES_NOT_EXIST`

const ErrorCodeTemporarilyUnavailable ErrorCode = `TEMPORARILY_UNAVAILABLE`

const ErrorCodeUnauthenticated ErrorCode = `UNAUTHENTICATED`

const ErrorCodeUnavailable ErrorCode = `UNAVAILABLE`

const ErrorCodeUnknown ErrorCode = `UNKNOWN`

const ErrorCodeUnparseableHttpError ErrorCode = `UNPARSEABLE_HTTP_ERROR`

const ErrorCodeWorkspaceTemporarilyUnavailable ErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `CATALOG_ALREADY_EXISTS`, `CATALOG_DOES_NOT_EXIST`, `CATALOG_NOT_EMPTY`, `COULD_NOT_ACQUIRE_LOCK`, `CUSTOMER_UNAUTHORIZED`, `DAC_ALREADY_EXISTS`, `DAC_DOES_NOT_EXIST`, `DATA_LOSS`, `DEADLINE_EXCEEDED`, `DEPLOYMENT_TIMEOUT`, `DIRECTORY_NOT_EMPTY`, `DIRECTORY_PROTECTED`, `DRY_RUN_FAILED`, `ENDPOINT_NOT_FOUND`, `EXTERNAL_LOCATION_ALREADY_EXISTS`, `EXTERNAL_LOCATION_DOES_NOT_EXIST`, `FEATURE_DISABLED`, `GIT_CONFLICT`, `GIT_REMOTE_ERROR`, `GIT_SENSITIVE_TOKEN_DETECTED`, `GIT_UNKNOWN_REF`, `GIT_URL_NOT_ON_ALLOW_LIST`, `INSECURE_PARTNER_RESPONSE`, `INTERNAL_ERROR`, `INVALID_PARAMETER_VALUE`, `INVALID_STATE`, `INVALID_STATE_TRANSITION`, `IO_ERROR`, `IPYNB_FILE_IN_REPO`, `MALFORMED_PARTNER_RESPONSE`, `MALFORMED_REQUEST`, `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`, `MAX_BLOCK_SIZE_EXCEEDED`, `MAX_CHILD_NODE_SIZE_EXCEEDED`, `MAX_LIST_SIZE_EXCEEDED`, `MAX_NOTEBOOK_SIZE_EXCEEDED`, `MAX_READ_SIZE_EXCEEDED`, `METASTORE_ALREADY_EXISTS`, `METASTORE_DOES_NOT_EXIST`, `METASTORE_NOT_EMPTY`, `NOT_FOUND`, `NOT_IMPLEMENTED`, `PARTIAL_DELETE`, `PERMISSION_DENIED`, `PERMISSION_NOT_PROPAGATED`, `PRINCIPAL_DOES_NOT_EXIST`, `PROJECTS_OPERATION_TIMEOUT`, `PROVIDER_ALREADY_EXISTS`, `PROVIDER_DOES_NOT_EXIST`, `PROVIDER_SHARE_NOT_ACCESSIBLE`, `QUOTA_EXCEEDED`, `RECIPIENT_ALREADY_EXISTS`, `RECIPIENT_DOES_NOT_EXIST`, `REQUEST_LIMIT_EXCEEDED`, `RESOURCE_ALREADY_EXISTS`, `RESOURCE_CONFLICT`, `RESOURCE_DOES_NOT_EXIST`, `RESOURCE_EXHAUSTED`, `RESOURCE_LIMIT_EXCEEDED`, `SCHEMA_ALREADY_EXISTS`, `SCHEMA_DOES_NOT_EXIST`, `SCHEMA_NOT_EMPTY`, `SEARCH_QUERY_TOO_LONG`, `SEARCH_QUERY_TOO_SHORT`, `SERVICE_UNDER_MAINTENANCE`, `SHARE_ALREADY_EXISTS`, `SHARE_DOES_NOT_EXIST`, `STORAGE_CREDENTIAL_ALREADY_EXISTS`, `STORAGE_CREDENTIAL_DOES_NOT_EXIST`, `TABLE_ALREADY_EXISTS`, `TABLE_DOES_NOT_EXIST`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNAVAILABLE`, `UNKNOWN`, `UNPARSEABLE_HTTP_ERROR`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "CATALOG_ALREADY_EXISTS", "CATALOG_DOES_NOT_EXIST", "CATALOG_NOT_EMPTY", "COULD_NOT_ACQUIRE_LOCK", "CUSTOMER_UNAUTHORIZED", "DAC_ALREADY_EXISTS", "DAC_DOES_NOT_EXIST", "DATA_LOSS", "DEADLINE_EXCEEDED", "DEPLOYMENT_TIMEOUT", "DIRECTORY_NOT_EMPTY", "DIRECTORY_PROTECTED", "DRY_RUN_FAILED", "ENDPOINT_NOT_FOUND", "EXTERNAL_LOCATION_ALREADY_EXISTS", "EXTERNAL_LOCATION_DOES_NOT_EXIST", "FEATURE_DISABLED", "GIT_CONFLICT", "GIT_REMOTE_ERROR", "GIT_SENSITIVE_TOKEN_DETECTED", "GIT_UNKNOWN_REF", "GIT_URL_NOT_ON_ALLOW_LIST", "INSECURE_PARTNER_RESPONSE", "INTERNAL_ERROR", "INVALID_PARAMETER_VALUE", "INVALID_STATE", "INVALID_STATE_TRANSITION", "IO_ERROR", "IPYNB_FILE_IN_REPO", "MALFORMED_PARTNER_RESPONSE", "MALFORMED_REQUEST", "MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST", "MAX_BLOCK_SIZE_EXCEEDED", "MAX_CHILD_NODE_SIZE_EXCEEDED", "MAX_LIST_SIZE_EXCEEDED", "MAX_NOTEBOOK_SIZE_EXCEEDED", "MAX_READ_SIZE_EXCEEDED", "METASTORE_ALREADY_EXISTS", "METASTORE_DOES_NOT_EXIST", "METASTORE_NOT_EMPTY", "NOT_FOUND", "NOT_IMPLEMENTED", "PARTIAL_DELETE", "PERMISSION_DENIED", "PERMISSION_NOT_PROPAGATED", "PRINCIPAL_DOES_NOT_EXIST", "PROJECTS_OPERATION_TIMEOUT", "PROVIDER_ALREADY_EXISTS", "PROVIDER_DOES_NOT_EXIST", "PROVIDER_SHARE_NOT_ACCESSIBLE", "QUOTA_EXCEEDED", "RECIPIENT_ALREADY_EXISTS", "RECIPIENT_DOES_NOT_EXIST", "REQUEST_LIMIT_EXCEEDED", "RESOURCE_ALREADY_EXISTS", "RESOURCE_CONFLICT", "RESOURCE_DOES_NOT_EXIST", "RESOURCE_EXHAUSTED", "RESOURCE_LIMIT_EXCEEDED", "SCHEMA_ALREADY_EXISTS", "SCHEMA_DOES_NOT_EXIST", "SCHEMA_NOT_EMPTY", "SEARCH_QUERY_TOO_LONG", "SEARCH_QUERY_TOO_SHORT", "SERVICE_UNDER_MAINTENANCE", "SHARE_ALREADY_EXISTS", "SHARE_DOES_NOT_EXIST", "STORAGE_CREDENTIAL_ALREADY_EXISTS", "STORAGE_CREDENTIAL_DOES_NOT_EXIST", "TABLE_ALREADY_EXISTS", "TABLE_DOES_NOT_EXIST", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNAVAILABLE", "UNKNOWN", "UNPARSEABLE_HTTP_ERROR", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Values returns all possible values for ErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		ErrorCodeAborted,
		ErrorCodeAlreadyExists,
		ErrorCodeBadRequest,
		ErrorCodeCancelled,
		ErrorCodeCatalogAlreadyExists,
		ErrorCodeCatalogDoesNotExist,
		ErrorCodeCatalogNotEmpty,
		ErrorCodeCouldNotAcquireLock,
		ErrorCodeCustomerUnauthorized,
		ErrorCodeDacAlreadyExists,
		ErrorCodeDacDoesNotExist,
		ErrorCodeDataLoss,
		ErrorCodeDeadlineExceeded,
		ErrorCodeDeploymentTimeout,
		ErrorCodeDirectoryNotEmpty,
		ErrorCodeDirectoryProtected,
		ErrorCodeDryRunFailed,
		ErrorCodeEndpointNotFound,
		ErrorCodeExternalLocationAlreadyExists,
		ErrorCodeExternalLocationDoesNotExist,
		ErrorCodeFeatureDisabled,
		ErrorCodeGitConflict,
		ErrorCodeGitRemoteError,
		ErrorCodeGitSensitiveTokenDetected,
		ErrorCodeGitUnknownRef,
		ErrorCodeGitUrlNotOnAllowList,
		ErrorCodeInsecurePartnerResponse,
		ErrorCodeInternalError,
		ErrorCodeInvalidParameterValue,
		ErrorCodeInvalidState,
		ErrorCodeInvalidStateTransition,
		ErrorCodeIoError,
		ErrorCodeIpynbFileInRepo,
		ErrorCodeMalformedPartnerResponse,
		ErrorCodeMalformedRequest,
		ErrorCodeManagedResourceGroupDoesNotExist,
		ErrorCodeMaxBlockSizeExceeded,
		ErrorCodeMaxChildNodeSizeExceeded,
		ErrorCodeMaxListSizeExceeded,
		ErrorCodeMaxNotebookSizeExceeded,
		ErrorCodeMaxReadSizeExceeded,
		ErrorCodeMetastoreAlreadyExists,
		ErrorCodeMetastoreDoesNotExist,
		ErrorCodeMetastoreNotEmpty,
		ErrorCodeNotFound,
		ErrorCodeNotImplemented,
		ErrorCodePartialDelete,
		ErrorCodePermissionDenied,
		ErrorCodePermissionNotPropagated,
		ErrorCodePrincipalDoesNotExist,
		ErrorCodeProjectsOperationTimeout,
		ErrorCodeProviderAlreadyExists,
		ErrorCodeProviderDoesNotExist,
		ErrorCodeProviderShareNotAccessible,
		ErrorCodeQuotaExceeded,
		ErrorCodeRecipientAlreadyExists,
		ErrorCodeRecipientDoesNotExist,
		ErrorCodeRequestLimitExceeded,
		ErrorCodeResourceAlreadyExists,
		ErrorCodeResourceConflict,
		ErrorCodeResourceDoesNotExist,
		ErrorCodeResourceExhausted,
		ErrorCodeResourceLimitExceeded,
		ErrorCodeSchemaAlreadyExists,
		ErrorCodeSchemaDoesNotExist,
		ErrorCodeSchemaNotEmpty,
		ErrorCodeSearchQueryTooLong,
		ErrorCodeSearchQueryTooShort,
		ErrorCodeServiceUnderMaintenance,
		ErrorCodeShareAlreadyExists,
		ErrorCodeShareDoesNotExist,
		ErrorCodeStorageCredentialAlreadyExists,
		ErrorCodeStorageCredentialDoesNotExist,
		ErrorCodeTableAlreadyExists,
		ErrorCodeTableDoesNotExist,
		ErrorCodeTemporarilyUnavailable,
		ErrorCodeUnauthenticated,
		ErrorCodeUnavailable,
		ErrorCodeUnknown,
		ErrorCodeUnparseableHttpError,
		ErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ErrorCode to satisfy [pflag.Value] interface
func (f *ErrorCode) Type() string {
	return "ErrorCode"
}

type GetDatabaseBranchRequest struct {
	// The name of the Database Branch to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
}

type GetDatabaseEndpointRequest struct {
	// The name of the Database Endpoint to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

type GetDatabaseProjectRequest struct {
	// The name of the Database Project to retrieve. Format:
	// projects/{project_id}
	Name string `json:"-" url:"-"`
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

type ListDatabaseBranchesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Branches. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Database Project, which owns this collection of branches. Format:
	// projects/{project_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseBranchesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseBranchesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseBranchesResponse struct {
	// List of branches.
	DatabaseBranches []DatabaseBranch `json:"database_branches,omitempty"`
	// Pagination token to request the next page of instances.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseBranchesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseBranchesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseEndpointsRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Branches. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Database Branch, which owns this collection of endpoints. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseEndpointsResponse struct {
	// List of endpoints.
	DatabaseEndpoints []DatabaseEndpoint `json:"database_endpoints,omitempty"`
	// Pagination token to request the next page of instances.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseEndpointsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseEndpointsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseProjectsRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Projects. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseProjectsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseProjectsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseProjectsResponse struct {
	// List of projects.
	DatabaseProjects []DatabaseProject `json:"database_projects,omitempty"`
	// Pagination token to request the next page of instances.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseProjectsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseProjectsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done,omitempty"`
	// The error result of the operation in case of failure or cancellation.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name string `json:"name,omitempty"`
	// The normal, successful response of the operation.
	Response json.RawMessage `json:"response,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Operation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Operation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestartDatabaseEndpointRequest struct {
	// The name of the Database Endpoint to restart. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

type UpdateDatabaseBranchRequest struct {
	// The Database Branch to update.
	//
	// The branch's `name` field is used to identify the branch to update.
	// Format: projects/{project_id}/branches/{branch_id}
	DatabaseBranch DatabaseBranch `json:"database_branch"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateDatabaseEndpointRequest struct {
	// The Database Endpoint to update.
	//
	// The endpoints's `name` field is used to identify the endpoint to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	DatabaseEndpoint DatabaseEndpoint `json:"database_endpoint"`
	// The resource name of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateDatabaseProjectRequest struct {
	// The Database Project to update.
	//
	// The project's `name` field is used to identify the project to update.
	// Format: projects/{project_id}
	DatabaseProject DatabaseProject `json:"database_project"`
	// The resource name of the project. Format: projects/{project_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask string `json:"-" url:"update_mask"`
}
