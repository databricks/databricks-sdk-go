// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// Create a Database Catalog
type CreateDatabaseCatalogRequest struct {
	Catalog DatabaseCatalog `json:"catalog"`
}

// Create a Database Instance
type CreateDatabaseInstanceRequest struct {
	// A DatabaseInstance represents a logical Postgres instance, comprised of
	// both compute and storage.
	DatabaseInstance DatabaseInstance `json:"database_instance"`
}

// Create a Database Table
type CreateDatabaseTableRequest struct {
	// Next field marker: 13
	Table DatabaseTable `json:"table"`
}

// Create a Synced Database Table
type CreateSyncedDatabaseTableRequest struct {
	// Next field marker: 12
	SyncedTable SyncedDatabaseTable `json:"synced_table"`
}

type DatabaseCatalog struct {
	CreateDatabaseIfNotExists bool `json:"create_database_if_not_exists,omitempty"`
	// The name of the DatabaseInstance housing the database.
	DatabaseInstanceName string `json:"database_instance_name"`
	// The name of the database (in a instance) associated with the catalog.
	DatabaseName string `json:"database_name"`
	// The name of the catalog in UC.
	Name string `json:"name"`

	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseCredential struct {
	ExpirationTime string `json:"expiration_time,omitempty"`

	Token string `json:"token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	Capacity string `json:"capacity,omitempty"`
	// The timestamp when the instance was created.
	CreationTime string `json:"creation_time,omitempty"`
	// The email of the creator of the instance.
	Creator string `json:"creator,omitempty"`
	// xref AIP-129. `stopped` is owned by the client, while `effective_stopped`
	// is owned by the server. `stopped` will only be set in Create/Update
	// response messages if and only if the user provides the field via the
	// request. `effective_stopped` on the other hand will always bet set in all
	// response messages (Create/Update/Get/List).
	EffectiveStopped bool `json:"effective_stopped,omitempty"`
	// The name of the instance. This is the unique identifier for the instance.
	Name string `json:"name"`
	// The version of Postgres running on the instance.
	PgVersion string `json:"pg_version,omitempty"`
	// The DNS endpoint to connect to the instance for read+write access.
	ReadWriteDns string `json:"read_write_dns,omitempty"`
	// The current state of the instance.
	State DatabaseInstanceState `json:"state,omitempty"`
	// Whether the instance is stopped.
	Stopped bool `json:"stopped,omitempty"`
	// An immutable UUID identifier for the instance.
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstance) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstance) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseInstanceState string

const DatabaseInstanceStateAvailable DatabaseInstanceState = `AVAILABLE`

const DatabaseInstanceStateDeleting DatabaseInstanceState = `DELETING`

const DatabaseInstanceStateFailingOver DatabaseInstanceState = `FAILING_OVER`

const DatabaseInstanceStateStarting DatabaseInstanceState = `STARTING`

const DatabaseInstanceStateStopped DatabaseInstanceState = `STOPPED`

const DatabaseInstanceStateUpdating DatabaseInstanceState = `UPDATING`

// String representation for [fmt.Print]
func (f *DatabaseInstanceState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseInstanceState) Set(v string) error {
	switch v {
	case `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`:
		*f = DatabaseInstanceState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVAILABLE", "DELETING", "FAILING_OVER", "STARTING", "STOPPED", "UPDATING"`, v)
	}
}

// Values returns all possible values for DatabaseInstanceState.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseInstanceState) Values() []DatabaseInstanceState {
	return []DatabaseInstanceState{
		DatabaseInstanceStateAvailable,
		DatabaseInstanceStateDeleting,
		DatabaseInstanceStateFailingOver,
		DatabaseInstanceStateStarting,
		DatabaseInstanceStateStopped,
		DatabaseInstanceStateUpdating,
	}
}

// Type always returns DatabaseInstanceState to satisfy [pflag.Value] interface
func (f *DatabaseInstanceState) Type() string {
	return "DatabaseInstanceState"
}

// Next field marker: 13
type DatabaseTable struct {
	// Name of the target database instance. This is required when creating
	// database tables in standard catalogs. This is optional when creating
	// database tables in registered catalogs. If this field is specified when
	// creating database tables in registered catalogs, the database instance
	// name MUST match that of the registered catalog (or the request will be
	// rejected).
	DatabaseInstanceName string `json:"database_instance_name,omitempty"`
	// Target Postgres database object (logical database) name for this table.
	// This field is optional in all scenarios.
	//
	// When creating a table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a table in a standard catalog, the target database name is
	// inferred to be that of the standard catalog. In this scenario, specifying
	// this field will allow targeting an arbitrary postgres database. Note that
	// this has implications for the `create_database_objects_is_missing` field
	// in `spec`.
	LogicalDatabaseName string `json:"logical_database_name,omitempty"`
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a Database Catalog
type DeleteDatabaseCatalogRequest struct {
	Name string `json:"-" url:"-"`
}

type DeleteDatabaseCatalogResponse struct {
}

// Delete a Database Instance
type DeleteDatabaseInstanceRequest struct {
	// By default, a instance cannot be deleted if it has descendant instances
	// created via PITR. If this flag is specified as true, all descendent
	// instances will be deleted as well.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the instance to delete.
	Name string `json:"-" url:"-"`
	// Note purge=false is in development. If false, the database instance is
	// soft deleted (implementation pending). Soft deleted instances behave as
	// if they are deleted, and cannot be used for CRUD operations nor connected
	// to. However they can be undeleted by calling the undelete API for a
	// limited time (implementation pending). If true, the database instance is
	// hard deleted and cannot be undeleted. For the time being, setting this
	// value to true is required to delete an instance (soft delete is not yet
	// supported).
	Purge bool `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDatabaseInstanceResponse struct {
}

// Delete a Database Table
type DeleteDatabaseTableRequest struct {
	Name string `json:"-" url:"-"`
}

type DeleteDatabaseTableResponse struct {
}

// Delete a Synced Database Table
type DeleteSyncedDatabaseTableRequest struct {
	Name string `json:"-" url:"-"`
}

type DeleteSyncedDatabaseTableResponse struct {
}

// Find a Database Instance by uid
type FindDatabaseInstanceByUidRequest struct {
	// UID of the cluster to get.
	Uid string `json:"-" url:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FindDatabaseInstanceByUidRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FindDatabaseInstanceByUidRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Generates a credential that can be used to access database instances
type GenerateDatabaseCredentialRequest struct {
	// Instances to which the token will be scoped.
	InstanceNames []string `json:"instance_names,omitempty"`

	RequestId string `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateDatabaseCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateDatabaseCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a Database Catalog
type GetDatabaseCatalogRequest struct {
	Name string `json:"-" url:"-"`
}

// Get a Database Instance
type GetDatabaseInstanceRequest struct {
	// Name of the cluster to get.
	Name string `json:"-" url:"-"`
}

// Get a Database Table
type GetDatabaseTableRequest struct {
	Name string `json:"-" url:"-"`
}

// Get a Synced Database Table
type GetSyncedDatabaseTableRequest struct {
	Name string `json:"-" url:"-"`
}

// List Database Instances
type ListDatabaseInstancesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseInstancesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstancesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseInstancesResponse struct {
	// List of instances.
	DatabaseInstances []DatabaseInstance `json:"database_instances,omitempty"`
	// Pagination token to request the next page of instances.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseInstancesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstancesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec struct {
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	StorageCatalog string `json:"storage_catalog,omitempty"`
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	StorageSchema string `json:"storage_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NewPipelineSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NewPipelineSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for ProvisioningInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ProvisioningInfoState) Values() []ProvisioningInfoState {
	return []ProvisioningInfoState{
		ProvisioningInfoStateActive,
		ProvisioningInfoStateDegraded,
		ProvisioningInfoStateDeleting,
		ProvisioningInfoStateFailed,
		ProvisioningInfoStateProvisioning,
		ProvisioningInfoStateUpdating,
	}
}

// Type always returns ProvisioningInfoState to satisfy [pflag.Value] interface
func (f *ProvisioningInfoState) Type() string {
	return "ProvisioningInfoState"
}

// Next field marker: 12
type SyncedDatabaseTable struct {
	// Synced Table data synchronization status
	DataSynchronizationStatus *SyncedTableStatus `json:"data_synchronization_status,omitempty"`
	// Name of the target database instance. This is required when creating
	// synced database tables in standard catalogs. This is optional when
	// creating synced database tables in registered catalogs. If this field is
	// specified when creating synced database tables in registered catalogs,
	// the database instance name MUST match that of the registered catalog (or
	// the request will be rejected).
	DatabaseInstanceName string `json:"database_instance_name,omitempty"`
	// Target Postgres database object (logical database) name for this table.
	// This field is optional in all scenarios.
	//
	// When creating a synced table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a synced table in a standard catalog, the target database
	// name is inferred to be that of the standard catalog. In this scenario,
	// specifying this field will allow targeting an arbitrary postgres
	// database.
	LogicalDatabaseName string `json:"logical_database_name,omitempty"`
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"name"`
	// Specification of a synced database table.
	Spec *SyncedTableSpec `json:"spec,omitempty"`
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedDatabaseTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedDatabaseTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
type SyncedTableContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
	// The last source table Delta version that was synced to the synced table.
	// Note that this Delta version may not be completely synced to the synced
	// table yet.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the synced table.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of a synced table. Shown if the synced table is in the
// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
type SyncedTableFailedStatus struct {
	// The last source table Delta version that was synced to the synced table.
	// Note that this Delta version may only be partially synced to the synced
	// table. Only populated if the table is still synced and available for
	// serving.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the synced table. Only populated if the table is still synced
	// and available for serving.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableFailedStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableFailedStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Progress information of the Synced Table data synchronization pipeline.
type SyncedTablePipelineProgress struct {
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

func (s *SyncedTablePipelineProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTablePipelineProgress) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of a synced table. Shown if the synced table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type SyncedTableProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

type SyncedTableSchedulingPolicy string

const SyncedTableSchedulingPolicyContinuous SyncedTableSchedulingPolicy = `CONTINUOUS`

const SyncedTableSchedulingPolicySnapshot SyncedTableSchedulingPolicy = `SNAPSHOT`

const SyncedTableSchedulingPolicyTriggered SyncedTableSchedulingPolicy = `TRIGGERED`

// String representation for [fmt.Print]
func (f *SyncedTableSchedulingPolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableSchedulingPolicy) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`:
		*f = SyncedTableSchedulingPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "SNAPSHOT", "TRIGGERED"`, v)
	}
}

// Values returns all possible values for SyncedTableSchedulingPolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableSchedulingPolicy) Values() []SyncedTableSchedulingPolicy {
	return []SyncedTableSchedulingPolicy{
		SyncedTableSchedulingPolicyContinuous,
		SyncedTableSchedulingPolicySnapshot,
		SyncedTableSchedulingPolicyTriggered,
	}
}

// Type always returns SyncedTableSchedulingPolicy to satisfy [pflag.Value] interface
func (f *SyncedTableSchedulingPolicy) Type() string {
	return "SyncedTableSchedulingPolicy"
}

// Specification of a synced database table.
type SyncedTableSpec struct {
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist.
	CreateDatabaseObjectsIfMissing bool `json:"create_database_objects_if_missing,omitempty"`
	// User-specified ID of a pre-existing pipeline to bin pack. This field is
	// optional, and should be empty if new_pipeline_spec is set. This field
	// will only be set by the server in response messages if it is specified in
	// the request. The SyncedTableStatus message will always contain the
	// effective pipeline ID (either client provided or server generated),
	// however.
	ExistingPipelineId string `json:"existing_pipeline_id,omitempty"`
	// Spec of new pipeline. Should be empty if pipeline_id /
	// existing_pipeline_id is set
	NewPipelineSpec *NewPipelineSpec `json:"new_pipeline_spec,omitempty"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Scheduling policy of the underlying pipeline.
	SchedulingPolicy SyncedTableSchedulingPolicy `json:"scheduling_policy,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey string `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of a synced table.
type SyncedTableState string

const SyncedTableStateSyncedTabledOffline SyncedTableState = `SYNCED_TABLED_OFFLINE`

const SyncedTableStateSyncedTableOfflineFailed SyncedTableState = `SYNCED_TABLE_OFFLINE_FAILED`

const SyncedTableStateSyncedTableOnline SyncedTableState = `SYNCED_TABLE_ONLINE`

const SyncedTableStateSyncedTableOnlineContinuousUpdate SyncedTableState = `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`

const SyncedTableStateSyncedTableOnlineNoPendingUpdate SyncedTableState = `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`

const SyncedTableStateSyncedTableOnlinePipelineFailed SyncedTableState = `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`

const SyncedTableStateSyncedTableOnlineTriggeredUpdate SyncedTableState = `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`

const SyncedTableStateSyncedTableOnlineUpdatingPipelineResources SyncedTableState = `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`

const SyncedTableStateSyncedTableProvisioning SyncedTableState = `SYNCED_TABLE_PROVISIONING`

const SyncedTableStateSyncedTableProvisioningInitialSnapshot SyncedTableState = `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`

const SyncedTableStateSyncedTableProvisioningPipelineResources SyncedTableState = `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`

// String representation for [fmt.Print]
func (f *SyncedTableState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableState) Set(v string) error {
	switch v {
	case `SYNCED_TABLED_OFFLINE`, `SYNCED_TABLE_OFFLINE_FAILED`, `SYNCED_TABLE_ONLINE`, `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`, `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`, `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`, `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`, `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`, `SYNCED_TABLE_PROVISIONING`, `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`, `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`:
		*f = SyncedTableState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SYNCED_TABLED_OFFLINE", "SYNCED_TABLE_OFFLINE_FAILED", "SYNCED_TABLE_ONLINE", "SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE", "SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE", "SYNCED_TABLE_ONLINE_PIPELINE_FAILED", "SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE", "SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES", "SYNCED_TABLE_PROVISIONING", "SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT", "SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Values returns all possible values for SyncedTableState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableState) Values() []SyncedTableState {
	return []SyncedTableState{
		SyncedTableStateSyncedTabledOffline,
		SyncedTableStateSyncedTableOfflineFailed,
		SyncedTableStateSyncedTableOnline,
		SyncedTableStateSyncedTableOnlineContinuousUpdate,
		SyncedTableStateSyncedTableOnlineNoPendingUpdate,
		SyncedTableStateSyncedTableOnlinePipelineFailed,
		SyncedTableStateSyncedTableOnlineTriggeredUpdate,
		SyncedTableStateSyncedTableOnlineUpdatingPipelineResources,
		SyncedTableStateSyncedTableProvisioning,
		SyncedTableStateSyncedTableProvisioningInitialSnapshot,
		SyncedTableStateSyncedTableProvisioningPipelineResources,
	}
}

// Type always returns SyncedTableState to satisfy [pflag.Value] interface
func (f *SyncedTableState) Type() string {
	return "SyncedTableState"
}

// Status of a synced table.
type SyncedTableStatus struct {
	// Detailed status of a synced table. Shown if the synced table is in the
	// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus *SyncedTableContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	// The state of the synced table.
	DetailedState SyncedTableState `json:"detailed_state,omitempty"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
	FailedStatus *SyncedTableFailedStatus `json:"failed_status,omitempty"`
	// A text description of the current state of the synced table.
	Message string `json:"message,omitempty"`
	// ID of the associated pipeline. The pipeline ID may have been provided by
	// the client (in the case of bin packing), or generated by the server (when
	// creating a new pipeline).
	PipelineId string `json:"pipeline_id,omitempty"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus *SyncedTableProvisioningStatus `json:"provisioning_status,omitempty"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus *SyncedTableTriggeredUpdateStatus `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
type SyncedTableTriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the synced table.
	// Note that this Delta version may not be completely synced to the synced
	// table yet.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the synced table.
	Timestamp string `json:"timestamp,omitempty"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress *SyncedTablePipelineProgress `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableTriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableTriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Update a Database Instance
type UpdateDatabaseInstanceRequest struct {
	// A DatabaseInstance represents a logical Postgres instance, comprised of
	// both compute and storage.
	DatabaseInstance DatabaseInstance `json:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	Name string `json:"-" url:"-"`
	// The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}
