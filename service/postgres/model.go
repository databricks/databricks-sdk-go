// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/duration"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type Branch struct {
	// A timestamp indicating when the branch was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"name,omitempty"`
	// The project containing this branch. Format: projects/{project_id}
	Parent string `json:"parent,omitempty"`
	// The desired state of a Branch.
	Spec *BranchSpec `json:"spec,omitempty"`
	// The current status of a Branch.
	Status *BranchStatus `json:"status,omitempty"`
	// System generated unique ID for the branch.
	Uid string `json:"uid,omitempty"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Branch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Branch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BranchOperationMetadata struct {
}

type BranchSpec struct {
	// Whether the branch is the project's default branch.
	Default bool `json:"default,omitempty"`
	// Whether the branch is protected.
	IsProtected bool `json:"is_protected,omitempty"`
	// The name of the source branch from which this branch was created. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch string `json:"source_branch,omitempty"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	SourceBranchLsn string `json:"source_branch_lsn,omitempty"`
	// The point in time on the source branch from which this branch was
	// created.
	SourceBranchTime *time.Time `json:"source_branch_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BranchSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BranchSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BranchStatus struct {
	// The branch's state, indicating if it is initializing, ready for use, or
	// archived.
	CurrentState BranchStatusState `json:"current_state,omitempty"`
	// Whether the branch is the project's default branch.
	Default bool `json:"default,omitempty"`
	// Whether the branch is protected.
	IsProtected bool `json:"is_protected,omitempty"`
	// The logical size of the branch.
	LogicalSizeBytes int64 `json:"logical_size_bytes,omitempty"`
	// The pending state of the branch, if a state transition is in progress.
	PendingState BranchStatusState `json:"pending_state,omitempty"`
	// The name of the source branch from which this branch was created. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch string `json:"source_branch,omitempty"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	SourceBranchLsn string `json:"source_branch_lsn,omitempty"`
	// The point in time on the source branch from which this branch was
	// created.
	SourceBranchTime *time.Time `json:"source_branch_time,omitempty"`
	// A timestamp indicating when the `current_state` began.
	StateChangeTime *time.Time `json:"state_change_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BranchStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BranchStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the database branch.
type BranchStatusState string

const BranchStatusStateArchived BranchStatusState = `ARCHIVED`

const BranchStatusStateImporting BranchStatusState = `IMPORTING`

const BranchStatusStateInit BranchStatusState = `INIT`

const BranchStatusStateReady BranchStatusState = `READY`

const BranchStatusStateResetting BranchStatusState = `RESETTING`

// String representation for [fmt.Print]
func (f *BranchStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *BranchStatusState) Set(v string) error {
	switch v {
	case `ARCHIVED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`:
		*f = BranchStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARCHIVED", "IMPORTING", "INIT", "READY", "RESETTING"`, v)
	}
}

// Values returns all possible values for BranchStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *BranchStatusState) Values() []BranchStatusState {
	return []BranchStatusState{
		BranchStatusStateArchived,
		BranchStatusStateImporting,
		BranchStatusStateInit,
		BranchStatusStateReady,
		BranchStatusStateResetting,
	}
}

// Type always returns BranchStatusState to satisfy [pflag.Value] interface
func (f *BranchStatusState) Type() string {
	return "BranchStatusState"
}

type CreateBranchRequest struct {
	// The Branch to create.
	Branch Branch `json:"branch"`
	// The ID to use for the Branch, which will become the final component of
	// the branch's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	BranchId string `json:"-" url:"branch_id,omitempty"`
	// The Project where this Branch will be created. Format:
	// projects/{project_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBranchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBranchRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpointRequest struct {
	// The Endpoint to create.
	Endpoint Endpoint `json:"endpoint"`
	// The ID to use for the Endpoint, which will become the final component of
	// the endpoint's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	EndpointId string `json:"-" url:"endpoint_id,omitempty"`
	// The Branch where this Endpoint will be created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateProjectRequest struct {
	// The Project to create.
	Project Project `json:"project"`
	// The ID to use for the Project, which will become the final component of
	// the project's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	ProjectId string `json:"-" url:"project_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateProjectRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRoleRequest struct {
	// The Branch where this Role is created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`
	// The desired specification of a Role.
	Role Role `json:"role"`
	// The ID to use for the Role, which will become the final component of the
	// branch's resource name. This ID becomes the role in postgres.
	//
	// This value should be 4-63 characters, and only use characters available
	// in DNS names, as defined by RFC-1123
	RoleId string `json:"-" url:"role_id"`
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
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

type DeleteBranchRequest struct {
	// The name of the Branch to delete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
}

type DeleteEndpointRequest struct {
	// The name of the Endpoint to delete. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

type DeleteProjectRequest struct {
	// The name of the Project to delete. Format: projects/{project_id}
	Name string `json:"-" url:"-"`
}

type DeleteRoleRequest struct {
	// The resource name of the postgres role. Format:
	// projects/{project_id}/branch/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
	// Reassign objects. If this is set, all objects owned by the role are
	// reassigned to the role specified in this parameter.
	//
	// NOTE: setting this requires spinning up a compute to succeed, since it
	// involves running SQL queries.
	//
	// TODO: #LKB-7187 implement reassign_owned_to on LBM side. This might
	// end-up being a synchronous query when this parameter is used.
	ReassignOwnedTo string `json:"-" url:"reassign_owned_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Endpoint struct {
	// A timestamp indicating when the compute endpoint was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The resource name of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"name,omitempty"`
	// The branch containing this endpoint. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`
	// The desired state of an Endpoint.
	Spec *EndpointSpec `json:"spec,omitempty"`
	// The current status of an Endpoint.
	Status *EndpointStatus `json:"status,omitempty"`
	// System generated unique ID for the endpoint.
	Uid string `json:"uid,omitempty"`
	// A timestamp indicating when the compute endpoint was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Endpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Endpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointOperationMetadata struct {
}

// A collection of settings for a compute endpoint.
type EndpointSettings struct {
	// A raw representation of Postgres settings.
	PgSettings map[string]string `json:"pg_settings,omitempty"`
}

type EndpointSpec struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled bool `json:"disabled,omitempty"`
	// The endpoint type. A branch can only have one READ_WRITE endpoint.
	EndpointType EndpointType `json:"endpoint_type"`

	Settings *EndpointSettings `json:"settings,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration *duration.Duration `json:"suspend_timeout_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointStatus struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`

	CurrentState EndpointStatusState `json:"current_state,omitempty"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled bool `json:"disabled,omitempty"`
	// The endpoint type. A branch can only have one READ_WRITE endpoint.
	EndpointType EndpointType `json:"endpoint_type,omitempty"`
	// The hostname of the compute endpoint. This is the hostname specified when
	// connecting to a database.
	Host string `json:"host,omitempty"`
	// A timestamp indicating when the compute endpoint was last active.
	LastActiveTime *time.Time `json:"last_active_time,omitempty"`

	PendingState EndpointStatusState `json:"pending_state,omitempty"`

	Settings *EndpointSettings `json:"settings,omitempty"`
	// A timestamp indicating when the compute endpoint was last started.
	StartTime *time.Time `json:"start_time,omitempty"`
	// A timestamp indicating when the compute endpoint was last suspended.
	SuspendTime *time.Time `json:"suspend_time,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration *duration.Duration `json:"suspend_timeout_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the compute endpoint.
type EndpointStatusState string

const EndpointStatusStateActive EndpointStatusState = `ACTIVE`

const EndpointStatusStateIdle EndpointStatusState = `IDLE`

const EndpointStatusStateInit EndpointStatusState = `INIT`

// String representation for [fmt.Print]
func (f *EndpointStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStatusState) Set(v string) error {
	switch v {
	case `ACTIVE`, `IDLE`, `INIT`:
		*f = EndpointStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "IDLE", "INIT"`, v)
	}
}

// Values returns all possible values for EndpointStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStatusState) Values() []EndpointStatusState {
	return []EndpointStatusState{
		EndpointStatusStateActive,
		EndpointStatusStateIdle,
		EndpointStatusStateInit,
	}
}

// Type always returns EndpointStatusState to satisfy [pflag.Value] interface
func (f *EndpointStatusState) Type() string {
	return "EndpointStatusState"
}

// The compute endpoint type. Either `read_write` or `read_only`.
type EndpointType string

const EndpointTypeReadOnly EndpointType = `READ_ONLY`

const EndpointTypeReadWrite EndpointType = `READ_WRITE`

// String representation for [fmt.Print]
func (f *EndpointType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointType) Set(v string) error {
	switch v {
	case `READ_ONLY`, `READ_WRITE`:
		*f = EndpointType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ_ONLY", "READ_WRITE"`, v)
	}
}

// Values returns all possible values for EndpointType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointType) Values() []EndpointType {
	return []EndpointType{
		EndpointTypeReadOnly,
		EndpointTypeReadWrite,
	}
}

// Type always returns EndpointType to satisfy [pflag.Value] interface
func (f *EndpointType) Type() string {
	return "EndpointType"
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

type GetBranchRequest struct {
	// The name of the Branch to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
}

type GetEndpointRequest struct {
	// The name of the Endpoint to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

type GetProjectRequest struct {
	// The name of the Project to retrieve. Format: projects/{project_id}
	Name string `json:"-" url:"-"`
}

type GetRoleRequest struct {
	// The name of the Role to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
}

type ListBranchesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Branches. Requests first page
	// if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Project that owns this collection of branches. Format:
	// projects/{project_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBranchesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBranchesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBranchesResponse struct {
	// List of branches.
	Branches []Branch `json:"branches,omitempty"`
	// Pagination token to request the next page of branches.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBranchesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBranchesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Endpoints. Requests first page
	// if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Branch that owns this collection of endpoints. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsResponse struct {
	// List of endpoints.
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// Pagination token to request the next page of endpoints.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProjectsRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Projects. Requests first page
	// if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProjectsResponse struct {
	// Pagination token to request the next page of projects.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of projects.
	Projects []Project `json:"projects,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRolesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Roles. Requests first page if
	// absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Branch that owns this collection of roles. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRolesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRolesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRolesResponse struct {
	// Pagination token to request the next page of roles.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of roles.
	Roles []Role `json:"roles,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRolesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRolesResponse) MarshalJSON() ([]byte, error) {
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

type Project struct {
	// A timestamp indicating when the project was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The resource name of the project. Format: projects/{project_id}
	Name string `json:"name,omitempty"`
	// The desired state of a Project.
	Spec *ProjectSpec `json:"spec,omitempty"`
	// The current status of a Project.
	Status *ProjectStatus `json:"status,omitempty"`
	// System generated unique ID for the project.
	Uid string `json:"uid,omitempty"`
	// A timestamp indicating when the project was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Project) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Project) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A collection of settings for a compute endpoint.
type ProjectDefaultEndpointSettings struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// A raw representation of Postgres settings.
	PgSettings map[string]string `json:"pg_settings,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration *duration.Duration `json:"suspend_timeout_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectDefaultEndpointSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectDefaultEndpointSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProjectOperationMetadata struct {
}

type ProjectSettings struct {
	// Sets wal_level=logical for all compute endpoints in this project. All
	// active endpoints will be suspended. Once enabled, logical replication
	// cannot be disabled.
	EnableLogicalReplication bool `json:"enable_logical_replication,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProjectSpec struct {
	DefaultEndpointSettings *ProjectDefaultEndpointSettings `json:"default_endpoint_settings,omitempty"`
	// Human-readable project name.
	DisplayName string `json:"display_name,omitempty"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project.
	HistoryRetentionDuration *duration.Duration `json:"history_retention_duration,omitempty"`
	// The major Postgres version number.
	PgVersion int `json:"pg_version,omitempty"`

	Settings *ProjectSettings `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProjectStatus struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes int64 `json:"branch_logical_size_limit_bytes,omitempty"`
	// The most recent time when any endpoint of this project was active.
	ComputeLastActiveTime *time.Time `json:"compute_last_active_time,omitempty"`
	// The effective default endpoint settings.
	DefaultEndpointSettings *ProjectDefaultEndpointSettings `json:"default_endpoint_settings,omitempty"`
	// The effective human-readable project name.
	DisplayName string `json:"display_name,omitempty"`
	// The effective number of seconds to retain the shared history for point in
	// time recovery.
	HistoryRetentionDuration *duration.Duration `json:"history_retention_duration,omitempty"`
	// The effective major Postgres version number.
	PgVersion int `json:"pg_version,omitempty"`
	// The effective project settings.
	Settings *ProjectSettings `json:"settings,omitempty"`
	// The current space occupied by the project in storage.
	SyntheticStorageSizeBytes int64 `json:"synthetic_storage_size_bytes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Role represents a Postgres role within a Branch.
type Role struct {
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The resource name of the role. Format:
	// projects/{project_id}/branch/{branch_id}/roles/{role_id}
	Name string `json:"name,omitempty"`
	// The Branch where this Role exists. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`
	// The desired state of the Role.
	Spec *RoleRoleSpec `json:"spec,omitempty"`
	// The observed state of the Role.
	Status *RoleRoleStatus `json:"status,omitempty"`

	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Role) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Role) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// How the role is authenticated when connecting to Postgres.
type RoleAuthMethod string

const RoleAuthMethodLakebaseOauthV1 RoleAuthMethod = `LAKEBASE_OAUTH_V1`

const RoleAuthMethodNoLogin RoleAuthMethod = `NO_LOGIN`

const RoleAuthMethodPgPasswordScramSha256 RoleAuthMethod = `PG_PASSWORD_SCRAM_SHA_256`

// String representation for [fmt.Print]
func (f *RoleAuthMethod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RoleAuthMethod) Set(v string) error {
	switch v {
	case `LAKEBASE_OAUTH_V1`, `NO_LOGIN`, `PG_PASSWORD_SCRAM_SHA_256`:
		*f = RoleAuthMethod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LAKEBASE_OAUTH_V1", "NO_LOGIN", "PG_PASSWORD_SCRAM_SHA_256"`, v)
	}
}

// Values returns all possible values for RoleAuthMethod.
//
// There is no guarantee on the order of the values in the slice.
func (f *RoleAuthMethod) Values() []RoleAuthMethod {
	return []RoleAuthMethod{
		RoleAuthMethodLakebaseOauthV1,
		RoleAuthMethodNoLogin,
		RoleAuthMethodPgPasswordScramSha256,
	}
}

// Type always returns RoleAuthMethod to satisfy [pflag.Value] interface
func (f *RoleAuthMethod) Type() string {
	return "RoleAuthMethod"
}

// The type of the Databricks managed identity that this Role represents. Leave
// empty if you wish to create a regular Postgres role not associated with a
// Databricks identity.
type RoleIdentityType string

const RoleIdentityTypeGroup RoleIdentityType = `GROUP`

const RoleIdentityTypeServicePrincipal RoleIdentityType = `SERVICE_PRINCIPAL`

const RoleIdentityTypeUser RoleIdentityType = `USER`

// String representation for [fmt.Print]
func (f *RoleIdentityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RoleIdentityType) Set(v string) error {
	switch v {
	case `GROUP`, `SERVICE_PRINCIPAL`, `USER`:
		*f = RoleIdentityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GROUP", "SERVICE_PRINCIPAL", "USER"`, v)
	}
}

// Values returns all possible values for RoleIdentityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RoleIdentityType) Values() []RoleIdentityType {
	return []RoleIdentityType{
		RoleIdentityTypeGroup,
		RoleIdentityTypeServicePrincipal,
		RoleIdentityTypeUser,
	}
}

// Type always returns RoleIdentityType to satisfy [pflag.Value] interface
func (f *RoleIdentityType) Type() string {
	return "RoleIdentityType"
}

type RoleOperationMetadata struct {
}

type RoleRoleSpec struct {
	// If auth_method is left unspecified, a meaningful authentication method is
	// derived from the identity_type: * For the managed identities, OAUTH is
	// used. * For the regular postgres roles, authentication based on postgres
	// passwords is used.
	//
	// NOTE: this is ignored for the Databricks identity type GROUP, and
	// NO_LOGIN is implicitly assumed instead for the GROUP identity type.
	AuthMethod RoleAuthMethod `json:"auth_method,omitempty"`
	// The type of the role. When specifying a managed-identity, the chosen
	// role_id must be a valid:
	//
	// * application ID for SERVICE_PRINCIPAL * user email for USER * group name
	// for GROUP
	IdentityType RoleIdentityType `json:"identity_type,omitempty"`
}

type RoleRoleStatus struct {
	AuthMethod RoleAuthMethod `json:"auth_method,omitempty"`
	// The type of the role.
	IdentityType RoleIdentityType `json:"identity_type,omitempty"`
}

type UpdateBranchRequest struct {
	// The Branch to update.
	//
	// The branch's `name` field is used to identify the branch to update.
	// Format: projects/{project_id}/branches/{branch_id}
	Branch Branch `json:"branch"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

type UpdateEndpointRequest struct {
	// The Endpoint to update.
	//
	// The endpoint's `name` field is used to identify the endpoint to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint Endpoint `json:"endpoint"`
	// The resource name of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

type UpdateProjectRequest struct {
	// The resource name of the project. Format: projects/{project_id}
	Name string `json:"-" url:"-"`
	// The Project to update.
	//
	// The project's `name` field is used to identify the project to update.
	// Format: projects/{project_id}
	Project Project `json:"project"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}
