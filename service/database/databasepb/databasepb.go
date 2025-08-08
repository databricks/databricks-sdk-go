// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package databasepb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateDatabaseCatalogRequestPb struct {
	Catalog DatabaseCatalogPb `json:"catalog"`
}

type CreateDatabaseInstanceRequestPb struct {
	DatabaseInstance DatabaseInstancePb `json:"database_instance"`
}

type CreateDatabaseInstanceRoleRequestPb struct {
	DatabaseInstanceRole DatabaseInstanceRolePb `json:"database_instance_role"`
	InstanceName         string                 `json:"-" url:"-"`
}

type CreateDatabaseTableRequestPb struct {
	Table DatabaseTablePb `json:"table"`
}

type CreateSyncedDatabaseTableRequestPb struct {
	SyncedTable SyncedDatabaseTablePb `json:"synced_table"`
}

type DatabaseCatalogPb struct {
	CreateDatabaseIfNotExists bool   `json:"create_database_if_not_exists,omitempty"`
	DatabaseInstanceName      string `json:"database_instance_name"`
	DatabaseName              string `json:"database_name"`
	Name                      string `json:"name"`
	Uid                       string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabaseCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabaseCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabaseCredentialPb struct {
	ExpirationTime string `json:"expiration_time,omitempty"`
	Token          string `json:"token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabaseCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabaseCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabaseInstancePb struct {
	Capacity                           string                  `json:"capacity,omitempty"`
	ChildInstanceRefs                  []DatabaseInstanceRefPb `json:"child_instance_refs,omitempty"`
	CreationTime                       string                  `json:"creation_time,omitempty"`
	Creator                            string                  `json:"creator,omitempty"`
	EffectiveEnableReadableSecondaries bool                    `json:"effective_enable_readable_secondaries,omitempty"`
	EffectiveNodeCount                 int                     `json:"effective_node_count,omitempty"`
	EffectiveRetentionWindowInDays     int                     `json:"effective_retention_window_in_days,omitempty"`
	EffectiveStopped                   bool                    `json:"effective_stopped,omitempty"`
	EnableReadableSecondaries          bool                    `json:"enable_readable_secondaries,omitempty"`
	Name                               string                  `json:"name"`
	NodeCount                          int                     `json:"node_count,omitempty"`
	ParentInstanceRef                  *DatabaseInstanceRefPb  `json:"parent_instance_ref,omitempty"`
	PgVersion                          string                  `json:"pg_version,omitempty"`
	ReadOnlyDns                        string                  `json:"read_only_dns,omitempty"`
	ReadWriteDns                       string                  `json:"read_write_dns,omitempty"`
	RetentionWindowInDays              int                     `json:"retention_window_in_days,omitempty"`
	State                              DatabaseInstanceStatePb `json:"state,omitempty"`
	Stopped                            bool                    `json:"stopped,omitempty"`
	Uid                                string                  `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabaseInstancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabaseInstancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabaseInstanceRefPb struct {
	BranchTime   string `json:"branch_time,omitempty"`
	EffectiveLsn string `json:"effective_lsn,omitempty"`
	Lsn          string `json:"lsn,omitempty"`
	Name         string `json:"name,omitempty"`
	Uid          string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabaseInstanceRefPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabaseInstanceRefPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabaseInstanceRolePb struct {
	Attributes     *DatabaseInstanceRoleAttributesPb    `json:"attributes,omitempty"`
	IdentityType   DatabaseInstanceRoleIdentityTypePb   `json:"identity_type,omitempty"`
	MembershipRole DatabaseInstanceRoleMembershipRolePb `json:"membership_role,omitempty"`
	Name           string                               `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabaseInstanceRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabaseInstanceRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabaseInstanceRoleAttributesPb struct {
	Bypassrls  bool `json:"bypassrls,omitempty"`
	Createdb   bool `json:"createdb,omitempty"`
	Createrole bool `json:"createrole,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabaseInstanceRoleAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabaseInstanceRoleAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabaseInstanceRoleIdentityTypePb string

const DatabaseInstanceRoleIdentityTypeGroup DatabaseInstanceRoleIdentityTypePb = `GROUP`

const DatabaseInstanceRoleIdentityTypePgOnly DatabaseInstanceRoleIdentityTypePb = `PG_ONLY`

const DatabaseInstanceRoleIdentityTypeServicePrincipal DatabaseInstanceRoleIdentityTypePb = `SERVICE_PRINCIPAL`

const DatabaseInstanceRoleIdentityTypeUser DatabaseInstanceRoleIdentityTypePb = `USER`

type DatabaseInstanceRoleMembershipRolePb string

const DatabaseInstanceRoleMembershipRoleDatabricksSuperuser DatabaseInstanceRoleMembershipRolePb = `DATABRICKS_SUPERUSER`

type DatabaseInstanceStatePb string

const DatabaseInstanceStateAvailable DatabaseInstanceStatePb = `AVAILABLE`

const DatabaseInstanceStateDeleting DatabaseInstanceStatePb = `DELETING`

const DatabaseInstanceStateFailingOver DatabaseInstanceStatePb = `FAILING_OVER`

const DatabaseInstanceStateStarting DatabaseInstanceStatePb = `STARTING`

const DatabaseInstanceStateStopped DatabaseInstanceStatePb = `STOPPED`

const DatabaseInstanceStateUpdating DatabaseInstanceStatePb = `UPDATING`

type DatabaseTablePb struct {
	DatabaseInstanceName string `json:"database_instance_name,omitempty"`
	LogicalDatabaseName  string `json:"logical_database_name,omitempty"`
	Name                 string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabaseTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabaseTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDatabaseCatalogRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeleteDatabaseInstanceRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`
	Purge bool   `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDatabaseInstanceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDatabaseInstanceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDatabaseInstanceRoleRequestPb struct {
	AllowMissing    bool   `json:"-" url:"allow_missing,omitempty"`
	InstanceName    string `json:"-" url:"-"`
	Name            string `json:"-" url:"-"`
	ReassignOwnedTo string `json:"-" url:"reassign_owned_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDatabaseInstanceRoleRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDatabaseInstanceRoleRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDatabaseTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeleteSyncedDatabaseTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeltaTableSyncInfoPb struct {
	DeltaCommitTimestamp string `json:"delta_commit_timestamp,omitempty"`
	DeltaCommitVersion   int64  `json:"delta_commit_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeltaTableSyncInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeltaTableSyncInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FindDatabaseInstanceByUidRequestPb struct {
	Uid string `json:"-" url:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FindDatabaseInstanceByUidRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FindDatabaseInstanceByUidRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenerateDatabaseCredentialRequestPb struct {
	Claims        []RequestedClaimsPb `json:"claims,omitempty"`
	InstanceNames []string            `json:"instance_names,omitempty"`
	RequestId     string              `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenerateDatabaseCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenerateDatabaseCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetDatabaseCatalogRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetDatabaseInstanceRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetDatabaseInstanceRoleRequestPb struct {
	InstanceName string `json:"-" url:"-"`
	Name         string `json:"-" url:"-"`
}

type GetDatabaseTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetSyncedDatabaseTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

type ListDatabaseInstanceRolesRequestPb struct {
	InstanceName string `json:"-" url:"-"`
	PageSize     int    `json:"-" url:"page_size,omitempty"`
	PageToken    string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDatabaseInstanceRolesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDatabaseInstanceRolesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListDatabaseInstanceRolesResponsePb struct {
	DatabaseInstanceRoles []DatabaseInstanceRolePb `json:"database_instance_roles,omitempty"`
	NextPageToken         string                   `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDatabaseInstanceRolesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDatabaseInstanceRolesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListDatabaseInstancesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDatabaseInstancesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDatabaseInstancesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListDatabaseInstancesResponsePb struct {
	DatabaseInstances []DatabaseInstancePb `json:"database_instances,omitempty"`
	NextPageToken     string               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDatabaseInstancesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDatabaseInstancesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NewPipelineSpecPb struct {
	StorageCatalog string `json:"storage_catalog,omitempty"`
	StorageSchema  string `json:"storage_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NewPipelineSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NewPipelineSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ProvisioningInfoStatePb string

const ProvisioningInfoStateActive ProvisioningInfoStatePb = `ACTIVE`

const ProvisioningInfoStateDegraded ProvisioningInfoStatePb = `DEGRADED`

const ProvisioningInfoStateDeleting ProvisioningInfoStatePb = `DELETING`

const ProvisioningInfoStateFailed ProvisioningInfoStatePb = `FAILED`

const ProvisioningInfoStateProvisioning ProvisioningInfoStatePb = `PROVISIONING`

const ProvisioningInfoStateUpdating ProvisioningInfoStatePb = `UPDATING`

type ProvisioningPhasePb string

const ProvisioningPhaseProvisioningPhaseIndexScan ProvisioningPhasePb = `PROVISIONING_PHASE_INDEX_SCAN`

const ProvisioningPhaseProvisioningPhaseIndexSort ProvisioningPhasePb = `PROVISIONING_PHASE_INDEX_SORT`

const ProvisioningPhaseProvisioningPhaseMain ProvisioningPhasePb = `PROVISIONING_PHASE_MAIN`

type RequestedClaimsPb struct {
	PermissionSet RequestedClaimsPermissionSetPb `json:"permission_set,omitempty"`
	Resources     []RequestedResourcePb          `json:"resources,omitempty"`
}

type RequestedClaimsPermissionSetPb string

const RequestedClaimsPermissionSetReadOnly RequestedClaimsPermissionSetPb = `READ_ONLY`

type RequestedResourcePb struct {
	TableName               string `json:"table_name,omitempty"`
	UnspecifiedResourceName string `json:"unspecified_resource_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RequestedResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RequestedResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SyncedDatabaseTablePb struct {
	DataSynchronizationStatus     *SyncedTableStatusPb    `json:"data_synchronization_status,omitempty"`
	DatabaseInstanceName          string                  `json:"database_instance_name,omitempty"`
	LogicalDatabaseName           string                  `json:"logical_database_name,omitempty"`
	Name                          string                  `json:"name"`
	Spec                          *SyncedTableSpecPb      `json:"spec,omitempty"`
	UnityCatalogProvisioningState ProvisioningInfoStatePb `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SyncedDatabaseTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SyncedDatabaseTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SyncedTableContinuousUpdateStatusPb struct {
	InitialPipelineSyncProgress *SyncedTablePipelineProgressPb `json:"initial_pipeline_sync_progress,omitempty"`
	LastProcessedCommitVersion  int64                          `json:"last_processed_commit_version,omitempty"`
	Timestamp                   string                         `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SyncedTableContinuousUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SyncedTableContinuousUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SyncedTableFailedStatusPb struct {
	LastProcessedCommitVersion int64  `json:"last_processed_commit_version,omitempty"`
	Timestamp                  string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SyncedTableFailedStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SyncedTableFailedStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SyncedTablePipelineProgressPb struct {
	EstimatedCompletionTimeSeconds   float64             `json:"estimated_completion_time_seconds,omitempty"`
	LatestVersionCurrentlyProcessing int64               `json:"latest_version_currently_processing,omitempty"`
	ProvisioningPhase                ProvisioningPhasePb `json:"provisioning_phase,omitempty"`
	SyncProgressCompletion           float64             `json:"sync_progress_completion,omitempty"`
	SyncedRowCount                   int64               `json:"synced_row_count,omitempty"`
	TotalRowCount                    int64               `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SyncedTablePipelineProgressPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SyncedTablePipelineProgressPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SyncedTablePositionPb struct {
	DeltaTableSyncInfo *DeltaTableSyncInfoPb `json:"delta_table_sync_info,omitempty"`
	SyncEndTimestamp   string                `json:"sync_end_timestamp,omitempty"`
	SyncStartTimestamp string                `json:"sync_start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SyncedTablePositionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SyncedTablePositionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SyncedTableProvisioningStatusPb struct {
	InitialPipelineSyncProgress *SyncedTablePipelineProgressPb `json:"initial_pipeline_sync_progress,omitempty"`
}

type SyncedTableSchedulingPolicyPb string

const SyncedTableSchedulingPolicyContinuous SyncedTableSchedulingPolicyPb = `CONTINUOUS`

const SyncedTableSchedulingPolicySnapshot SyncedTableSchedulingPolicyPb = `SNAPSHOT`

const SyncedTableSchedulingPolicyTriggered SyncedTableSchedulingPolicyPb = `TRIGGERED`

type SyncedTableSpecPb struct {
	CreateDatabaseObjectsIfMissing bool                          `json:"create_database_objects_if_missing,omitempty"`
	ExistingPipelineId             string                        `json:"existing_pipeline_id,omitempty"`
	NewPipelineSpec                *NewPipelineSpecPb            `json:"new_pipeline_spec,omitempty"`
	PrimaryKeyColumns              []string                      `json:"primary_key_columns,omitempty"`
	SchedulingPolicy               SyncedTableSchedulingPolicyPb `json:"scheduling_policy,omitempty"`
	SourceTableFullName            string                        `json:"source_table_full_name,omitempty"`
	TimeseriesKey                  string                        `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SyncedTableSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SyncedTableSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SyncedTableStatePb string

const SyncedTableStateSyncedTabledOffline SyncedTableStatePb = `SYNCED_TABLED_OFFLINE`

const SyncedTableStateSyncedTableOfflineFailed SyncedTableStatePb = `SYNCED_TABLE_OFFLINE_FAILED`

const SyncedTableStateSyncedTableOnline SyncedTableStatePb = `SYNCED_TABLE_ONLINE`

const SyncedTableStateSyncedTableOnlineContinuousUpdate SyncedTableStatePb = `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`

const SyncedTableStateSyncedTableOnlineNoPendingUpdate SyncedTableStatePb = `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`

const SyncedTableStateSyncedTableOnlinePipelineFailed SyncedTableStatePb = `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`

const SyncedTableStateSyncedTableOnlineTriggeredUpdate SyncedTableStatePb = `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`

const SyncedTableStateSyncedTableOnlineUpdatingPipelineResources SyncedTableStatePb = `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`

const SyncedTableStateSyncedTableProvisioning SyncedTableStatePb = `SYNCED_TABLE_PROVISIONING`

const SyncedTableStateSyncedTableProvisioningInitialSnapshot SyncedTableStatePb = `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`

const SyncedTableStateSyncedTableProvisioningPipelineResources SyncedTableStatePb = `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`

type SyncedTableStatusPb struct {
	ContinuousUpdateStatus *SyncedTableContinuousUpdateStatusPb `json:"continuous_update_status,omitempty"`
	DetailedState          SyncedTableStatePb                   `json:"detailed_state,omitempty"`
	FailedStatus           *SyncedTableFailedStatusPb           `json:"failed_status,omitempty"`
	LastSync               *SyncedTablePositionPb               `json:"last_sync,omitempty"`
	Message                string                               `json:"message,omitempty"`
	PipelineId             string                               `json:"pipeline_id,omitempty"`
	ProvisioningStatus     *SyncedTableProvisioningStatusPb     `json:"provisioning_status,omitempty"`
	TriggeredUpdateStatus  *SyncedTableTriggeredUpdateStatusPb  `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SyncedTableStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SyncedTableStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SyncedTableTriggeredUpdateStatusPb struct {
	LastProcessedCommitVersion int64                          `json:"last_processed_commit_version,omitempty"`
	Timestamp                  string                         `json:"timestamp,omitempty"`
	TriggeredUpdateProgress    *SyncedTablePipelineProgressPb `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SyncedTableTriggeredUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SyncedTableTriggeredUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateDatabaseInstanceRequestPb struct {
	DatabaseInstance DatabaseInstancePb `json:"database_instance"`
	Name             string             `json:"-" url:"-"`
	UpdateMask       string             `json:"-" url:"update_mask"`
}
