// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/marshal"
)

type CleanRoom struct {
	// Whether clean room access is restricted due to [CSP]
	//
	// [CSP]: https://docs.databricks.com/en/security/privacy/security-profile.html
	AccessRestricted CleanRoomAccessRestricted `json:"access_restricted,omitempty"`

	Comment string `json:"comment,omitempty"`
	// When the clean room was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// The alias of the collaborator tied to the local clean room.
	LocalCollaboratorAlias string `json:"local_collaborator_alias,omitempty"`
	// The name of the clean room. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	Name string `json:"name,omitempty"`
	// Output catalog of the clean room. It is an output only field. Output
	// catalog is manipulated using the separate CreateCleanRoomOutputCatalog
	// API.
	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
	// This is Databricks username of the owner of the local clean room
	// securable for permission management.
	Owner string `json:"owner,omitempty"`
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	RemoteDetailedInfo *CleanRoomRemoteDetail `json:"remote_detailed_info,omitempty"`
	// Clean room status.
	Status CleanRoomStatusEnum `json:"status,omitempty"`
	// When the clean room was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoom) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoom) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomAccessRestricted string

const CleanRoomAccessRestrictedCspMismatch CleanRoomAccessRestricted = `CSP_MISMATCH`

const CleanRoomAccessRestrictedNoRestriction CleanRoomAccessRestricted = `NO_RESTRICTION`

// String representation for [fmt.Print]
func (f *CleanRoomAccessRestricted) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomAccessRestricted) Set(v string) error {
	switch v {
	case `CSP_MISMATCH`, `NO_RESTRICTION`:
		*f = CleanRoomAccessRestricted(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CSP_MISMATCH", "NO_RESTRICTION"`, v)
	}
}

// Type always returns CleanRoomAccessRestricted to satisfy [pflag.Value] interface
func (f *CleanRoomAccessRestricted) Type() string {
	return "CleanRoomAccessRestricted"
}

// Metadata of the clean room asset
type CleanRoomAsset struct {
	// When the asset is added to the clean room, in epoch milliseconds.
	AddedAt int64 `json:"added_at,omitempty"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"asset_type,omitempty"`
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTable *CleanRoomAssetForeignTable `json:"foreign_table,omitempty"`
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTableLocalDetails *CleanRoomAssetForeignTableLocalDetails `json:"foreign_table_local_details,omitempty"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name string `json:"name,omitempty"`
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	Notebook *CleanRoomAssetNotebook `json:"notebook,omitempty"`
	// The alias of the collaborator who owns this asset
	OwnerCollaboratorAlias string `json:"owner_collaborator_alias,omitempty"`
	// Status of the asset
	Status CleanRoomAssetStatusEnum `json:"status,omitempty"`
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	Table *CleanRoomAssetTable `json:"table,omitempty"`
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	TableLocalDetails *CleanRoomAssetTableLocalDetails `json:"table_local_details,omitempty"`
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	View *CleanRoomAssetView `json:"view,omitempty"`
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	ViewLocalDetails *CleanRoomAssetViewLocalDetails `json:"view_local_details,omitempty"`
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	VolumeLocalDetails *CleanRoomAssetVolumeLocalDetails `json:"volume_local_details,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomAsset) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAsset) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomAssetAssetType string

const CleanRoomAssetAssetTypeForeignTable CleanRoomAssetAssetType = `FOREIGN_TABLE`

const CleanRoomAssetAssetTypeNotebookFile CleanRoomAssetAssetType = `NOTEBOOK_FILE`

const CleanRoomAssetAssetTypeTable CleanRoomAssetAssetType = `TABLE`

const CleanRoomAssetAssetTypeView CleanRoomAssetAssetType = `VIEW`

const CleanRoomAssetAssetTypeVolume CleanRoomAssetAssetType = `VOLUME`

// String representation for [fmt.Print]
func (f *CleanRoomAssetAssetType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomAssetAssetType) Set(v string) error {
	switch v {
	case `FOREIGN_TABLE`, `NOTEBOOK_FILE`, `TABLE`, `VIEW`, `VOLUME`:
		*f = CleanRoomAssetAssetType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FOREIGN_TABLE", "NOTEBOOK_FILE", "TABLE", "VIEW", "VOLUME"`, v)
	}
}

// Type always returns CleanRoomAssetAssetType to satisfy [pflag.Value] interface
func (f *CleanRoomAssetAssetType) Type() string {
	return "CleanRoomAssetAssetType"
}

type CleanRoomAssetForeignTable struct {
	// The metadata information of the columns in the foreign table
	Columns []ColumnInfo `json:"columns,omitempty"`
}

type CleanRoomAssetForeignTableLocalDetails struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomAssetForeignTableLocalDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAssetForeignTableLocalDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomAssetNotebook struct {
	// Server generated checksum that represents the notebook version.
	Etag string `json:"etag,omitempty"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	NotebookContent string `json:"notebook_content,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomAssetNotebook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAssetNotebook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomAssetStatusEnum string

const CleanRoomAssetStatusEnumActive CleanRoomAssetStatusEnum = `ACTIVE`

const CleanRoomAssetStatusEnumPending CleanRoomAssetStatusEnum = `PENDING`

const CleanRoomAssetStatusEnumPermissionDenied CleanRoomAssetStatusEnum = `PERMISSION_DENIED`

// String representation for [fmt.Print]
func (f *CleanRoomAssetStatusEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomAssetStatusEnum) Set(v string) error {
	switch v {
	case `ACTIVE`, `PENDING`, `PERMISSION_DENIED`:
		*f = CleanRoomAssetStatusEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "PENDING", "PERMISSION_DENIED"`, v)
	}
}

// Type always returns CleanRoomAssetStatusEnum to satisfy [pflag.Value] interface
func (f *CleanRoomAssetStatusEnum) Type() string {
	return "CleanRoomAssetStatusEnum"
}

type CleanRoomAssetTable struct {
	// The metadata information of the columns in the table
	Columns []ColumnInfo `json:"columns,omitempty"`
}

type CleanRoomAssetTableLocalDetails struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	LocalName string `json:"local_name,omitempty"`
	// Partition filtering specification for a shared table.
	Partitions []PartitionSpecificationPartition `json:"partitions,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomAssetTableLocalDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAssetTableLocalDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomAssetView struct {
	// The metadata information of the columns in the view
	Columns []ColumnInfo `json:"columns,omitempty"`
}

type CleanRoomAssetViewLocalDetails struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomAssetViewLocalDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAssetViewLocalDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomAssetVolumeLocalDetails struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomAssetVolumeLocalDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAssetVolumeLocalDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Publicly visible clean room collaborator.
type CleanRoomCollaborator struct {
	// Collaborator alias specified by the clean room creator. It is unique
	// across all collaborators of this clean room, and used to derive multiple
	// values internally such as catalog alias and clean room name for single
	// metastore clean rooms. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CollaboratorAlias string `json:"collaborator_alias"`
	// Generated display name for the collaborator. In the case of a single
	// metastore clean room, it is the clean room name. For x-metastore clean
	// rooms, it is the organization name of the metastore. It is not restricted
	// to these values and could change in the future
	DisplayName string `json:"display_name,omitempty"`
	// The global Unity Catalog metastore id of the collaborator. The identifier
	// is of format cloud:region:metastore-uuid.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Email of the user who is receiving the clean room "invitation". It should
	// be empty for the creator of the clean room, and non-empty for the
	// invitees of the clean room. It is only returned in the output when clean
	// room creator calls GET
	InviteRecipientEmail string `json:"invite_recipient_email,omitempty"`
	// Workspace ID of the user who is receiving the clean room "invitation".
	// Must be specified if invite_recipient_email is specified. It should be
	// empty when the collaborator is the creator of the clean room.
	InviteRecipientWorkspaceId int64 `json:"invite_recipient_workspace_id,omitempty"`
	// [Organization
	// name](:method:metastores/list#metastores-delta_sharing_organization_name)
	// configured in the metastore
	OrganizationName string `json:"organization_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomCollaborator) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomCollaborator) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Stores information about a single task run.
type CleanRoomNotebookTaskRun struct {
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API. if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	CollaboratorJobRunInfo *CollaboratorJobRunInfo `json:"collaborator_job_run_info,omitempty"`
	// State of the task run.
	NotebookJobRunState *CleanRoomTaskRunState `json:"notebook_job_run_state,omitempty"`
	// Asset name of the notebook executed in this task run.
	NotebookName string `json:"notebook_name,omitempty"`
	// Expiration time of the output schema of the task run (if any), in epoch
	// milliseconds.
	OutputSchemaExpirationTime int64 `json:"output_schema_expiration_time,omitempty"`
	// Name of the output schema associated with the clean rooms notebook task
	// run.
	OutputSchemaName string `json:"output_schema_name,omitempty"`
	// Duration of the task run, in milliseconds.
	RunDuration int64 `json:"run_duration,omitempty"`
	// When the task run started, in epoch milliseconds.
	StartTime int64 `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomNotebookTaskRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomNotebookTaskRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName string `json:"catalog_name,omitempty"`

	Status CleanRoomOutputCatalogOutputCatalogStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomOutputCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomOutputCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomOutputCatalogOutputCatalogStatus string

const CleanRoomOutputCatalogOutputCatalogStatusCreated CleanRoomOutputCatalogOutputCatalogStatus = `CREATED`

const CleanRoomOutputCatalogOutputCatalogStatusNotCreated CleanRoomOutputCatalogOutputCatalogStatus = `NOT_CREATED`

const CleanRoomOutputCatalogOutputCatalogStatusNotEligible CleanRoomOutputCatalogOutputCatalogStatus = `NOT_ELIGIBLE`

// String representation for [fmt.Print]
func (f *CleanRoomOutputCatalogOutputCatalogStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomOutputCatalogOutputCatalogStatus) Set(v string) error {
	switch v {
	case `CREATED`, `NOT_CREATED`, `NOT_ELIGIBLE`:
		*f = CleanRoomOutputCatalogOutputCatalogStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATED", "NOT_CREATED", "NOT_ELIGIBLE"`, v)
	}
}

// Type always returns CleanRoomOutputCatalogOutputCatalogStatus to satisfy [pflag.Value] interface
func (f *CleanRoomOutputCatalogOutputCatalogStatus) Type() string {
	return "CleanRoomOutputCatalogOutputCatalogStatus"
}

// Publicly visible central clean room details.
type CleanRoomRemoteDetail struct {
	// Central clean room ID.
	CentralCleanRoomId string `json:"central_clean_room_id,omitempty"`
	// Cloud vendor (aws,azure,gcp) of the central clean room.
	CloudVendor string `json:"cloud_vendor,omitempty"`
	// Collaborators in the central clean room. There should one and only one
	// collaborator in the list that satisfies the owner condition:
	//
	// 1. It has the creator's global_metastore_id (determined by caller of
	// CreateCleanRoom).
	//
	// 2. Its invite_recipient_email is empty.
	Collaborators []CleanRoomCollaborator `json:"collaborators,omitempty"`
	// The compliance security profile used to process regulated data following
	// compliance standards.
	ComplianceSecurityProfile *ComplianceSecurityProfile `json:"compliance_security_profile,omitempty"`
	// Collaborator who creates the clean room.
	Creator *CleanRoomCollaborator `json:"creator,omitempty"`
	// Egress network policy to apply to the central clean room workspace.
	EgressNetworkPolicy *EgressNetworkPolicy `json:"egress_network_policy,omitempty"`
	// Region of the central clean room.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CleanRoomRemoteDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomRemoteDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomStatusEnum string

const CleanRoomStatusEnumActive CleanRoomStatusEnum = `ACTIVE`

const CleanRoomStatusEnumDeleted CleanRoomStatusEnum = `DELETED`

const CleanRoomStatusEnumFailed CleanRoomStatusEnum = `FAILED`

const CleanRoomStatusEnumProvisioning CleanRoomStatusEnum = `PROVISIONING`

// String representation for [fmt.Print]
func (f *CleanRoomStatusEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomStatusEnum) Set(v string) error {
	switch v {
	case `ACTIVE`, `DELETED`, `FAILED`, `PROVISIONING`:
		*f = CleanRoomStatusEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DELETED", "FAILED", "PROVISIONING"`, v)
	}
}

// Type always returns CleanRoomStatusEnum to satisfy [pflag.Value] interface
func (f *CleanRoomStatusEnum) Type() string {
	return "CleanRoomStatusEnum"
}

// Copied from elastic-spark-common/api/messages/runs.proto. Using the original
// definition to remove coupling with jobs API definition
type CleanRoomTaskRunLifeCycleState string

const CleanRoomTaskRunLifeCycleStateBlocked CleanRoomTaskRunLifeCycleState = `BLOCKED`

const CleanRoomTaskRunLifeCycleStateInternalError CleanRoomTaskRunLifeCycleState = `INTERNAL_ERROR`

const CleanRoomTaskRunLifeCycleStatePending CleanRoomTaskRunLifeCycleState = `PENDING`

const CleanRoomTaskRunLifeCycleStateQueued CleanRoomTaskRunLifeCycleState = `QUEUED`

const CleanRoomTaskRunLifeCycleStateRunning CleanRoomTaskRunLifeCycleState = `RUNNING`

const CleanRoomTaskRunLifeCycleStateRunLifeCycleStateUnspecified CleanRoomTaskRunLifeCycleState = `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`

const CleanRoomTaskRunLifeCycleStateSkipped CleanRoomTaskRunLifeCycleState = `SKIPPED`

const CleanRoomTaskRunLifeCycleStateTerminated CleanRoomTaskRunLifeCycleState = `TERMINATED`

const CleanRoomTaskRunLifeCycleStateTerminating CleanRoomTaskRunLifeCycleState = `TERMINATING`

const CleanRoomTaskRunLifeCycleStateWaitingForRetry CleanRoomTaskRunLifeCycleState = `WAITING_FOR_RETRY`

// String representation for [fmt.Print]
func (f *CleanRoomTaskRunLifeCycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomTaskRunLifeCycleState) Set(v string) error {
	switch v {
	case `BLOCKED`, `INTERNAL_ERROR`, `PENDING`, `QUEUED`, `RUNNING`, `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`, `SKIPPED`, `TERMINATED`, `TERMINATING`, `WAITING_FOR_RETRY`:
		*f = CleanRoomTaskRunLifeCycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "INTERNAL_ERROR", "PENDING", "QUEUED", "RUNNING", "RUN_LIFE_CYCLE_STATE_UNSPECIFIED", "SKIPPED", "TERMINATED", "TERMINATING", "WAITING_FOR_RETRY"`, v)
	}
}

// Type always returns CleanRoomTaskRunLifeCycleState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunLifeCycleState) Type() string {
	return "CleanRoomTaskRunLifeCycleState"
}

// Copied from elastic-spark-common/api/messages/runs.proto. Using the original
// definition to avoid cyclic dependency.
type CleanRoomTaskRunResultState string

const CleanRoomTaskRunResultStateCanceled CleanRoomTaskRunResultState = `CANCELED`

const CleanRoomTaskRunResultStateDisabled CleanRoomTaskRunResultState = `DISABLED`

const CleanRoomTaskRunResultStateEvicted CleanRoomTaskRunResultState = `EVICTED`

const CleanRoomTaskRunResultStateExcluded CleanRoomTaskRunResultState = `EXCLUDED`

const CleanRoomTaskRunResultStateFailed CleanRoomTaskRunResultState = `FAILED`

const CleanRoomTaskRunResultStateMaximumConcurrentRunsReached CleanRoomTaskRunResultState = `MAXIMUM_CONCURRENT_RUNS_REACHED`

const CleanRoomTaskRunResultStateRunResultStateUnspecified CleanRoomTaskRunResultState = `RUN_RESULT_STATE_UNSPECIFIED`

const CleanRoomTaskRunResultStateSuccess CleanRoomTaskRunResultState = `SUCCESS`

const CleanRoomTaskRunResultStateSuccessWithFailures CleanRoomTaskRunResultState = `SUCCESS_WITH_FAILURES`

const CleanRoomTaskRunResultStateTimedout CleanRoomTaskRunResultState = `TIMEDOUT`

const CleanRoomTaskRunResultStateUpstreamCanceled CleanRoomTaskRunResultState = `UPSTREAM_CANCELED`

const CleanRoomTaskRunResultStateUpstreamEvicted CleanRoomTaskRunResultState = `UPSTREAM_EVICTED`

const CleanRoomTaskRunResultStateUpstreamFailed CleanRoomTaskRunResultState = `UPSTREAM_FAILED`

// String representation for [fmt.Print]
func (f *CleanRoomTaskRunResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomTaskRunResultState) Set(v string) error {
	switch v {
	case `CANCELED`, `DISABLED`, `EVICTED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `RUN_RESULT_STATE_UNSPECIFIED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_EVICTED`, `UPSTREAM_FAILED`:
		*f = CleanRoomTaskRunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "DISABLED", "EVICTED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "RUN_RESULT_STATE_UNSPECIFIED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_EVICTED", "UPSTREAM_FAILED"`, v)
	}
}

// Type always returns CleanRoomTaskRunResultState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunResultState) Type() string {
	return "CleanRoomTaskRunResultState"
}

// Stores the run state of the clean rooms notebook task.
type CleanRoomTaskRunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response.
	LifeCycleState CleanRoomTaskRunLifeCycleState `json:"life_cycle_state,omitempty"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states.
	ResultState CleanRoomTaskRunResultState `json:"result_state,omitempty"`
}

type CollaboratorJobRunInfo struct {
	// Alias of the collaborator that triggered the task run.
	CollaboratorAlias string `json:"collaborator_alias,omitempty"`
	// Job ID of the task run in the collaborator's workspace.
	CollaboratorJobId int64 `json:"collaborator_job_id,omitempty"`
	// Job run ID of the task run in the collaborator's workspace.
	CollaboratorJobRunId int64 `json:"collaborator_job_run_id,omitempty"`
	// Task run ID of the task run in the collaborator's workspace.
	CollaboratorTaskRunId int64 `json:"collaborator_task_run_id,omitempty"`
	// ID of the collaborator's workspace that triggered the task run.
	CollaboratorWorkspaceId int64 `json:"collaborator_workspace_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CollaboratorJobRunInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CollaboratorJobRunInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
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

// The compliance security profile used to process regulated data following
// compliance standards.
type ComplianceSecurityProfile struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	// Whether the compliance security profile is enabled.
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Compliance stardard for SHIELD customers
type ComplianceStandard string

const ComplianceStandardCanadaProtectedB ComplianceStandard = `CANADA_PROTECTED_B`

const ComplianceStandardCyberEssentialPlus ComplianceStandard = `CYBER_ESSENTIAL_PLUS`

const ComplianceStandardFedrampHigh ComplianceStandard = `FEDRAMP_HIGH`

const ComplianceStandardFedrampIl5 ComplianceStandard = `FEDRAMP_IL5`

const ComplianceStandardFedrampModerate ComplianceStandard = `FEDRAMP_MODERATE`

const ComplianceStandardHipaa ComplianceStandard = `HIPAA`

const ComplianceStandardHitrust ComplianceStandard = `HITRUST`

const ComplianceStandardIrapProtected ComplianceStandard = `IRAP_PROTECTED`

const ComplianceStandardIsmap ComplianceStandard = `ISMAP`

const ComplianceStandardItarEar ComplianceStandard = `ITAR_EAR`

const ComplianceStandardNone ComplianceStandard = `NONE`

const ComplianceStandardPciDss ComplianceStandard = `PCI_DSS`

// String representation for [fmt.Print]
func (f *ComplianceStandard) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComplianceStandard) Set(v string) error {
	switch v {
	case `CANADA_PROTECTED_B`, `CYBER_ESSENTIAL_PLUS`, `FEDRAMP_HIGH`, `FEDRAMP_IL5`, `FEDRAMP_MODERATE`, `HIPAA`, `HITRUST`, `IRAP_PROTECTED`, `ISMAP`, `ITAR_EAR`, `NONE`, `PCI_DSS`:
		*f = ComplianceStandard(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANADA_PROTECTED_B", "CYBER_ESSENTIAL_PLUS", "FEDRAMP_HIGH", "FEDRAMP_IL5", "FEDRAMP_MODERATE", "HIPAA", "HITRUST", "IRAP_PROTECTED", "ISMAP", "ITAR_EAR", "NONE", "PCI_DSS"`, v)
	}
}

// Type always returns ComplianceStandard to satisfy [pflag.Value] interface
func (f *ComplianceStandard) Type() string {
	return "ComplianceStandard"
}

// Create an asset
type CreateCleanRoomAssetRequest struct {
	// Metadata of the clean room asset
	Asset *CleanRoomAsset `json:"asset,omitempty"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
}

// Create an output catalog
type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`

	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
}

type CreateCleanRoomOutputCatalogResponse struct {
	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
}

// Create a clean room
type CreateCleanRoomRequest struct {
	CleanRoom *CleanRoom `json:"clean_room,omitempty"`
}

// Delete an asset
type DeleteCleanRoomAssetRequest struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	AssetFullName string `json:"-" url:"-"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
}

// Response for delete clean room request. Using an empty message since the
// generic Empty proto does not externd UnshadedMessageMarker.
type DeleteCleanRoomAssetResponse struct {
}

// Delete a clean room
type DeleteCleanRoomRequest struct {
	// Name of the clean room.
	Name string `json:"-" url:"-"`
}

type DeleteResponse struct {
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess *EgressNetworkPolicyInternetAccessPolicy `json:"internet_access,omitempty"`
}

type EgressNetworkPolicyInternetAccessPolicy struct {
	AllowedInternetDestinations []EgressNetworkPolicyInternetAccessPolicyInternetDestination `json:"allowed_internet_destinations,omitempty"`

	AllowedStorageDestinations []EgressNetworkPolicyInternetAccessPolicyStorageDestination `json:"allowed_storage_destinations,omitempty"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode `json:"log_only_mode,omitempty"`
	// At which level can Databricks and Databricks managed compute access
	// Internet. FULL_ACCESS: Databricks can access Internet. No blocking rules
	// will apply. RESTRICTED_ACCESS: Databricks can only access explicitly
	// allowed internet and storage destinations, as well as UC connections and
	// external locations. PRIVATE_ACCESS_ONLY (not used): Databricks can only
	// access destinations via private link.
	RestrictionMode EgressNetworkPolicyInternetAccessPolicyRestrictionMode `json:"restriction_mode,omitempty"`
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support domain name (FQDN) destinations for the time
// being, though going forwards we want to support host names and IP addresses.
type EgressNetworkPolicyInternetAccessPolicyInternetDestination struct {
	Destination string `json:"destination,omitempty"`
	// The filtering protocol used by the DP. For private and public preview,
	// SEG will only support TCP filtering (i.e. DNS based filtering, filtering
	// by destination IP address), so protocol will be set to TCP by default and
	// hidden from the user. In the future, users may be able to select HTTP
	// filtering (i.e. SNI based filtering, filtering by FQDN).
	Protocol EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol `json:"protocol,omitempty"`

	Type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EgressNetworkPolicyInternetAccessPolicyInternetDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EgressNetworkPolicyInternetAccessPolicyInternetDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The filtering protocol used by the DP. For private and public preview, SEG
// will only support TCP filtering (i.e. DNS based filtering, filtering by
// destination IP address), so protocol will be set to TCP by default and hidden
// from the user. In the future, users may be able to select HTTP filtering
// (i.e. SNI based filtering, filtering by FQDN).
type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol string

const EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolTcp EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol = `TCP`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) Set(v string) error {
	switch v {
	case `TCP`:
		*f = EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TCP"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol"
}

type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType string

const EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFqdn EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType = `FQDN`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Set(v string) error {
	switch v {
	case `FQDN`:
		*f = EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FQDN"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType"
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyMode struct {
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType `json:"log_only_mode_type,omitempty"`

	Workloads []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType `json:"workloads,omitempty"`
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType string

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeAllServices EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType = `ALL_SERVICES`

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeSelectedServices EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType = `SELECTED_SERVICES`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) Set(v string) error {
	switch v {
	case `ALL_SERVICES`, `SELECTED_SERVICES`:
		*f = EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_SERVICES", "SELECTED_SERVICES"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType"
}

// The values should match the list of workloads used in networkconfig.proto
type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType string

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeDbsql EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType = `DBSQL`

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeMlServing EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType = `ML_SERVING`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) Set(v string) error {
	switch v {
	case `DBSQL`, `ML_SERVING`:
		*f = EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DBSQL", "ML_SERVING"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType"
}

// At which level can Databricks and Databricks managed compute access Internet.
// FULL_ACCESS: Databricks can access Internet. No blocking rules will apply.
// RESTRICTED_ACCESS: Databricks can only access explicitly allowed internet and
// storage destinations, as well as UC connections and external locations.
// PRIVATE_ACCESS_ONLY (not used): Databricks can only access destinations via
// private link.
type EgressNetworkPolicyInternetAccessPolicyRestrictionMode string

const EgressNetworkPolicyInternetAccessPolicyRestrictionModeFullAccess EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `FULL_ACCESS`

const EgressNetworkPolicyInternetAccessPolicyRestrictionModePrivateAccessOnly EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `PRIVATE_ACCESS_ONLY`

const EgressNetworkPolicyInternetAccessPolicyRestrictionModeRestrictedAccess EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `RESTRICTED_ACCESS`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Set(v string) error {
	switch v {
	case `FULL_ACCESS`, `PRIVATE_ACCESS_ONLY`, `RESTRICTED_ACCESS`:
		*f = EgressNetworkPolicyInternetAccessPolicyRestrictionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL_ACCESS", "PRIVATE_ACCESS_ONLY", "RESTRICTED_ACCESS"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyRestrictionMode to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyRestrictionMode"
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyInternetAccessPolicyStorageDestination struct {
	AllowedPaths []string `json:"allowed_paths,omitempty"`

	AzureContainer string `json:"azure_container,omitempty"`

	AzureDnsZone string `json:"azure_dns_zone,omitempty"`

	AzureStorageAccount string `json:"azure_storage_account,omitempty"`

	AzureStorageService string `json:"azure_storage_service,omitempty"`

	BucketName string `json:"bucket_name,omitempty"`

	Region string `json:"region,omitempty"`

	Type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EgressNetworkPolicyInternetAccessPolicyStorageDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EgressNetworkPolicyInternetAccessPolicyStorageDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType string

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAwsS3 EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `AWS_S3`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `AZURE_STORAGE`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeCloudflareR2 EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `CLOUDFLARE_R2`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `GOOGLE_CLOUD_STORAGE`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) Set(v string) error {
	switch v {
	case `AWS_S3`, `AZURE_STORAGE`, `CLOUDFLARE_R2`, `GOOGLE_CLOUD_STORAGE`:
		*f = EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_S3", "AZURE_STORAGE", "CLOUDFLARE_R2", "GOOGLE_CLOUD_STORAGE"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType"
}

// Get an asset
type GetCleanRoomAssetRequest struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	AssetFullName string `json:"-" url:"-"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
}

// Get a clean room
type GetCleanRoomRequest struct {
	Name string `json:"-" url:"-"`
}

// List assets
type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListCleanRoomAssetsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomAssetsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomAssetsResponse struct {
	// Assets in the clean room.
	Assets []CleanRoomAsset `json:"assets,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListCleanRoomAssetsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomAssetsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List notebook task runs
type ListCleanRoomNotebookTaskRunsRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// Notebook name
	NotebookName string `json:"-" url:"notebook_name,omitempty"`
	// The maximum number of task runs to return
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListCleanRoomNotebookTaskRunsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomNotebookTaskRunsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomNotebookTaskRunsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// Name of the clean room.
	Runs []CleanRoomNotebookTaskRun `json:"runs,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListCleanRoomNotebookTaskRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomNotebookTaskRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List clean rooms
type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListCleanRoomsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomsResponse struct {
	CleanRooms []CleanRoom `json:"clean_rooms,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListCleanRoomsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PartitionSpecificationPartition struct {
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

	ForceSendFields []string `json:"-"`
}

func (s *PartitionValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Update an asset
type UpdateCleanRoomAssetRequest struct {
	// Metadata of the clean room asset
	Asset *CleanRoomAsset `json:"asset,omitempty"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name string `json:"-" url:"-"`
}

type UpdateCleanRoomRequest struct {
	CleanRoom *CleanRoom `json:"clean_room,omitempty"`
	// Name of the clean room.
	Name string `json:"-" url:"-"`
}
