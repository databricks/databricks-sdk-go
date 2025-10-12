// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateDatabaseBranchRequest struct {
	DatabaseBranch DatabaseBranch `json:"database_branch"`

	ProjectId string `json:"-" url:"-"`
}

type CreateDatabaseCatalogRequest struct {
	Catalog DatabaseCatalog `json:"catalog"`
}

type CreateDatabaseEndpointRequest struct {
	BranchId string `json:"-" url:"-"`

	DatabaseEndpoint DatabaseEndpoint `json:"database_endpoint"`

	ProjectId string `json:"-" url:"-"`
}

type CreateDatabaseInstanceRequest struct {
	// Instance to create.
	DatabaseInstance DatabaseInstance `json:"database_instance"`
}

type CreateDatabaseInstanceRoleRequest struct {
	DatabaseInstanceName string `json:"-" url:"database_instance_name,omitempty"`

	DatabaseInstanceRole DatabaseInstanceRole `json:"database_instance_role"`

	InstanceName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDatabaseProjectRequest struct {
	DatabaseProject DatabaseProject `json:"database_project"`
}

type CreateDatabaseTableRequest struct {
	Table DatabaseTable `json:"table"`
}

type CreateSyncedDatabaseTableRequest struct {
	SyncedTable SyncedDatabaseTable `json:"synced_table"`
}

type CustomTag struct {
	// The key of the custom tag.
	Key string `json:"key,omitempty"`
	// The value of the custom tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseBranch struct {
	BranchId string `json:"branch_id,omitempty"`
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
	// The id of the parent branch
	ParentId string `json:"parent_id,omitempty"`
	// The Log Sequence Number (LSN) on the parent branch from which this branch
	// was created. When restoring a branch using the Restore Database Branch
	// endpoint, this value isn’t finalized until all operations related to
	// the restore have completed successfully.
	ParentLsn string `json:"parent_lsn,omitempty"`
	// The point in time on the parent branch from which this branch was
	// created.
	ParentTime string `json:"parent_time,omitempty"`

	PendingState string `json:"pending_state,omitempty"`

	ProjectId string `json:"project_id"`
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

type DatabaseCatalog struct {
	CreateDatabaseIfNotExists bool `json:"create_database_if_not_exists,omitempty"`
	// The branch_id of the database branch associated with the catalog.
	DatabaseBranchId string `json:"database_branch_id,omitempty"`
	// The name of the DatabaseInstance housing the database.
	DatabaseInstanceName string `json:"database_instance_name"`
	// The name of the database (in a instance) associated with the catalog.
	DatabaseName string `json:"database_name"`
	// The project_id of the database project associated with the catalog.
	DatabaseProjectId string `json:"database_project_id,omitempty"`
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

type DatabaseEndpoint struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`

	BranchId string `json:"branch_id"`
	// A timestamp indicating when the compute endpoint was created.
	CreateTime string `json:"create_time,omitempty"`

	CurrentState DatabaseEndpointState `json:"current_state,omitempty"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled bool `json:"disabled,omitempty"`

	EndpointId string `json:"endpoint_id,omitempty"`
	// The hostname of the compute endpoint. This is the hostname specified when
	// connecting to a database.
	Host string `json:"host,omitempty"`
	// A timestamp indicating when the compute endpoint was last active.
	LastActiveTime string `json:"last_active_time,omitempty"`

	PendingState DatabaseEndpointState `json:"pending_state,omitempty"`

	PoolerMode DatabaseEndpointPoolerMode `json:"pooler_mode,omitempty"`

	ProjectId string `json:"project_id"`

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

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	Capacity string `json:"capacity,omitempty"`
	// The refs of the child instances. This is only available if the instance
	// is parent instance.
	ChildInstanceRefs []DatabaseInstanceRef `json:"child_instance_refs,omitempty"`
	// The timestamp when the instance was created.
	CreationTime string `json:"creation_time,omitempty"`
	// The email of the creator of the instance.
	Creator string `json:"creator,omitempty"`
	// Custom tags associated with the instance. This field is only included on
	// create and update responses.
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// Deprecated. The sku of the instance; this field will always match the
	// value of capacity.
	EffectiveCapacity string `json:"effective_capacity,omitempty"`
	// The recorded custom tags associated with the instance.
	EffectiveCustomTags []CustomTag `json:"effective_custom_tags,omitempty"`
	// Whether the instance has PG native password login enabled.
	EffectiveEnablePgNativeLogin bool `json:"effective_enable_pg_native_login,omitempty"`
	// Whether secondaries serving read-only traffic are enabled. Defaults to
	// false.
	EffectiveEnableReadableSecondaries bool `json:"effective_enable_readable_secondaries,omitempty"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries.
	EffectiveNodeCount int `json:"effective_node_count,omitempty"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained.
	EffectiveRetentionWindowInDays int `json:"effective_retention_window_in_days,omitempty"`
	// Whether the instance is stopped.
	EffectiveStopped bool `json:"effective_stopped,omitempty"`
	// The policy that is applied to the instance.
	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// Whether to enable PG native password login on the instance. Defaults to
	// false.
	EnablePgNativeLogin bool `json:"enable_pg_native_login,omitempty"`
	// Whether to enable secondaries to serve read-only traffic. Defaults to
	// false.
	EnableReadableSecondaries bool `json:"enable_readable_secondaries,omitempty"`
	// The name of the instance. This is the unique identifier for the instance.
	Name string `json:"name"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries. This field is input
	// only, see effective_node_count for the output.
	NodeCount int `json:"node_count,omitempty"`
	// The ref of the parent instance. This is only available if the instance is
	// child instance. Input: For specifying the parent instance to create a
	// child instance. Optional. Output: Only populated if provided as input to
	// create a child instance.
	ParentInstanceRef *DatabaseInstanceRef `json:"parent_instance_ref,omitempty"`
	// The version of Postgres running on the instance.
	PgVersion string `json:"pg_version,omitempty"`
	// The DNS endpoint to connect to the instance for read only access. This is
	// only available if enable_readable_secondaries is true.
	ReadOnlyDns string `json:"read_only_dns,omitempty"`
	// The DNS endpoint to connect to the instance for read+write access.
	ReadWriteDns string `json:"read_write_dns,omitempty"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained. The default value is 7 days.
	// Valid values are 2 to 35 days.
	RetentionWindowInDays int `json:"retention_window_in_days,omitempty"`
	// The current state of the instance.
	State DatabaseInstanceState `json:"state,omitempty"`
	// Whether to stop the instance. An input only param, see effective_stopped
	// for the output.
	Stopped bool `json:"stopped,omitempty"`
	// An immutable UUID identifier for the instance.
	Uid string `json:"uid,omitempty"`
	// The desired usage policy to associate with the instance.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstance) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstance) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// DatabaseInstanceRef is a reference to a database instance. It is used in the
// DatabaseInstance object to refer to the parent instance of an instance and to
// refer the child instances of an instance. To specify as a parent instance
// during creation of an instance, the lsn and branch_time fields are optional.
// If not specified, the child instance will be created from the latest lsn of
// the parent. If both lsn and branch_time are specified, the lsn will be used
// to create the child instance.
type DatabaseInstanceRef struct {
	// Branch time of the ref database instance. For a parent ref instance, this
	// is the point in time on the parent instance from which the instance was
	// created. For a child ref instance, this is the point in time on the
	// instance from which the child instance was created. Input: For specifying
	// the point in time to create a child instance. Optional. Output: Only
	// populated if provided as input to create a child instance.
	BranchTime string `json:"branch_time,omitempty"`
	// For a parent ref instance, this is the LSN on the parent instance from
	// which the instance was created. For a child ref instance, this is the LSN
	// on the instance from which the child instance was created.
	EffectiveLsn string `json:"effective_lsn,omitempty"`
	// User-specified WAL LSN of the ref database instance.
	//
	// Input: For specifying the WAL LSN to create a child instance. Optional.
	// Output: Only populated if provided as input to create a child instance.
	Lsn string `json:"lsn,omitempty"`
	// Name of the ref database instance.
	Name string `json:"name,omitempty"`
	// Id of the ref database instance.
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstanceRef) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRef) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A DatabaseInstanceRole represents a Postgres role in a database instance.
type DatabaseInstanceRole struct {
	// The desired API-exposed Postgres role attribute to associate with the
	// role. Optional.
	Attributes *DatabaseInstanceRoleAttributes `json:"attributes,omitempty"`
	// The attributes that are applied to the role.
	EffectiveAttributes *DatabaseInstanceRoleAttributes `json:"effective_attributes,omitempty"`
	// The type of the role.
	IdentityType DatabaseInstanceRoleIdentityType `json:"identity_type,omitempty"`

	InstanceName string `json:"instance_name,omitempty"`
	// An enum value for a standard role that this role is a member of.
	MembershipRole DatabaseInstanceRoleMembershipRole `json:"membership_role,omitempty"`
	// The name of the role. This is the unique identifier for the role in an
	// instance.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstanceRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Attributes that can be granted to a Postgres role. We are only implementing a
// subset for now, see xref:
// https://www.postgresql.org/docs/16/sql-createrole.html The values follow
// Postgres keyword naming e.g. CREATEDB, BYPASSRLS, etc. which is why they
// don't include typical underscores between words. We were requested to make
// this a nested object/struct representation since these are knobs from an
// external spec.
type DatabaseInstanceRoleAttributes struct {
	Bypassrls bool `json:"bypassrls,omitempty"`

	Createdb bool `json:"createdb,omitempty"`

	Createrole bool `json:"createrole,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstanceRoleAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRoleAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseInstanceRoleIdentityType string

const DatabaseInstanceRoleIdentityTypeGroup DatabaseInstanceRoleIdentityType = `GROUP`

const DatabaseInstanceRoleIdentityTypePgOnly DatabaseInstanceRoleIdentityType = `PG_ONLY`

const DatabaseInstanceRoleIdentityTypeServicePrincipal DatabaseInstanceRoleIdentityType = `SERVICE_PRINCIPAL`

const DatabaseInstanceRoleIdentityTypeUser DatabaseInstanceRoleIdentityType = `USER`

// String representation for [fmt.Print]
func (f *DatabaseInstanceRoleIdentityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseInstanceRoleIdentityType) Set(v string) error {
	switch v {
	case `GROUP`, `PG_ONLY`, `SERVICE_PRINCIPAL`, `USER`:
		*f = DatabaseInstanceRoleIdentityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GROUP", "PG_ONLY", "SERVICE_PRINCIPAL", "USER"`, v)
	}
}

// Values returns all possible values for DatabaseInstanceRoleIdentityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseInstanceRoleIdentityType) Values() []DatabaseInstanceRoleIdentityType {
	return []DatabaseInstanceRoleIdentityType{
		DatabaseInstanceRoleIdentityTypeGroup,
		DatabaseInstanceRoleIdentityTypePgOnly,
		DatabaseInstanceRoleIdentityTypeServicePrincipal,
		DatabaseInstanceRoleIdentityTypeUser,
	}
}

// Type always returns DatabaseInstanceRoleIdentityType to satisfy [pflag.Value] interface
func (f *DatabaseInstanceRoleIdentityType) Type() string {
	return "DatabaseInstanceRoleIdentityType"
}

// Roles that the DatabaseInstanceRole can be a member of.
type DatabaseInstanceRoleMembershipRole string

const DatabaseInstanceRoleMembershipRoleDatabricksSuperuser DatabaseInstanceRoleMembershipRole = `DATABRICKS_SUPERUSER`

// String representation for [fmt.Print]
func (f *DatabaseInstanceRoleMembershipRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseInstanceRoleMembershipRole) Set(v string) error {
	switch v {
	case `DATABRICKS_SUPERUSER`:
		*f = DatabaseInstanceRoleMembershipRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS_SUPERUSER"`, v)
	}
}

// Values returns all possible values for DatabaseInstanceRoleMembershipRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseInstanceRoleMembershipRole) Values() []DatabaseInstanceRoleMembershipRole {
	return []DatabaseInstanceRoleMembershipRole{
		DatabaseInstanceRoleMembershipRoleDatabricksSuperuser,
	}
}

// Type always returns DatabaseInstanceRoleMembershipRole to satisfy [pflag.Value] interface
func (f *DatabaseInstanceRoleMembershipRole) Type() string {
	return "DatabaseInstanceRoleMembershipRole"
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
	// The major Postgres version number.
	PgVersion int `json:"pg_version,omitempty"`

	ProjectId string `json:"project_id,omitempty"`

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
	//
	// When creating a table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a table in a standard catalog, this field is required. In
	// this scenario, specifying this field will allow targeting an arbitrary
	// postgres database.
	LogicalDatabaseName string `json:"logical_database_name,omitempty"`
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"name"`
	// Data serving REST API URL for this table
	TableServingUrl string `json:"table_serving_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDatabaseBranchRequest struct {
	BranchId string `json:"-" url:"-"`

	ProjectId string `json:"-" url:"-"`
}

type DeleteDatabaseCatalogRequest struct {
	Name string `json:"-" url:"-"`
}

type DeleteDatabaseEndpointRequest struct {
	BranchId string `json:"-" url:"-"`

	EndpointId string `json:"-" url:"-"`

	ProjectId string `json:"-" url:"-"`
}

type DeleteDatabaseInstanceRequest struct {
	// By default, a instance cannot be deleted if it has descendant instances
	// created via PITR. If this flag is specified as true, all descendent
	// instances will be deleted as well.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the instance to delete.
	Name string `json:"-" url:"-"`
	// Deprecated. Omitting the field or setting it to true will result in the
	// field being hard deleted. Setting a value of false will throw a bad
	// request.
	Purge bool `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDatabaseInstanceRoleRequest struct {
	// This is the AIP standard name for the equivalent of Postgres' `IF EXISTS`
	// option
	AllowMissing bool `json:"-" url:"allow_missing,omitempty"`

	InstanceName string `json:"-" url:"-"`

	Name string `json:"-" url:"-"`

	ReassignOwnedTo string `json:"-" url:"reassign_owned_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDatabaseProjectRequest struct {
	ProjectId string `json:"-" url:"-"`
}

type DeleteDatabaseTableRequest struct {
	Name string `json:"-" url:"-"`
}

type DeleteSyncedDatabaseTableRequest struct {
	Name string `json:"-" url:"-"`
}

type DeltaTableSyncInfo struct {
	// The timestamp when the above Delta version was committed in the source
	// Delta table. Note: This is the Delta commit time, not the time the data
	// was written to the synced table.
	DeltaCommitTimestamp string `json:"delta_commit_timestamp,omitempty"`
	// The Delta Lake commit version that was last successfully synced.
	DeltaCommitVersion int64 `json:"delta_commit_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaTableSyncInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaTableSyncInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FailoverDatabaseInstanceRequest struct {
	FailoverTargetDatabaseInstanceName string `json:"failover_target_database_instance_name,omitempty"`
	// Name of the instance to failover.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FailoverDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailoverDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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
	// The returned token will be scoped to the union of instance_names and
	// instances containing the specified UC tables, so instance_names is
	// allowed to be empty.
	Claims []RequestedClaims `json:"claims,omitempty"`
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

type GetDatabaseBranchRequest struct {
	BranchId string `json:"-" url:"-"`

	ProjectId string `json:"-" url:"-"`
}

type GetDatabaseCatalogRequest struct {
	Name string `json:"-" url:"-"`
}

type GetDatabaseEndpointRequest struct {
	BranchId string `json:"-" url:"-"`

	EndpointId string `json:"-" url:"-"`

	ProjectId string `json:"-" url:"-"`
}

type GetDatabaseInstanceRequest struct {
	// Name of the cluster to get.
	Name string `json:"-" url:"-"`
}

type GetDatabaseInstanceRoleRequest struct {
	InstanceName string `json:"-" url:"-"`

	Name string `json:"-" url:"-"`
}

type GetDatabaseProjectRequest struct {
	ProjectId string `json:"-" url:"-"`
}

type GetDatabaseTableRequest struct {
	Name string `json:"-" url:"-"`
}

type GetSyncedDatabaseTableRequest struct {
	Name string `json:"-" url:"-"`
}

type ListDatabaseBranchesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Branches. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ProjectId string `json:"-" url:"-"`

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

type ListDatabaseCatalogsRequest struct {
	// Name of the instance to get database catalogs for.
	InstanceName string `json:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseCatalogsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseCatalogsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseCatalogsResponse struct {
	DatabaseCatalogs []DatabaseCatalog `json:"database_catalogs,omitempty"`
	// Pagination token to request the next page of database catalogs.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseCatalogsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseCatalogsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseEndpointsRequest struct {
	BranchId string `json:"-" url:"-"`
	// Upper bound for items returned. If specified must be at least 10.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Endpoints. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ProjectId string `json:"-" url:"-"`

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

type ListDatabaseInstanceRolesRequest struct {
	InstanceName string `json:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseInstanceRolesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstanceRolesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseInstanceRolesResponse struct {
	// List of database instance roles.
	DatabaseInstanceRoles []DatabaseInstanceRole `json:"database_instance_roles,omitempty"`
	// Pagination token to request the next page of instances.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseInstanceRolesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstanceRolesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

type ListSyncedDatabaseTablesRequest struct {
	// Name of the instance to get synced tables for.
	InstanceName string `json:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSyncedDatabaseTablesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSyncedDatabaseTablesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSyncedDatabaseTablesResponse struct {
	// Pagination token to request the next page of synced tables.
	NextPageToken string `json:"next_page_token,omitempty"`

	SyncedTables []SyncedDatabaseTable `json:"synced_tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSyncedDatabaseTablesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSyncedDatabaseTablesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec struct {
	// Budget policy of this pipeline.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	StorageCatalog string `json:"storage_catalog,omitempty"`
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
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

type ProvisioningPhase string

const ProvisioningPhaseProvisioningPhaseIndexScan ProvisioningPhase = `PROVISIONING_PHASE_INDEX_SCAN`

const ProvisioningPhaseProvisioningPhaseIndexSort ProvisioningPhase = `PROVISIONING_PHASE_INDEX_SORT`

const ProvisioningPhaseProvisioningPhaseMain ProvisioningPhase = `PROVISIONING_PHASE_MAIN`

// String representation for [fmt.Print]
func (f *ProvisioningPhase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ProvisioningPhase) Set(v string) error {
	switch v {
	case `PROVISIONING_PHASE_INDEX_SCAN`, `PROVISIONING_PHASE_INDEX_SORT`, `PROVISIONING_PHASE_MAIN`:
		*f = ProvisioningPhase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PROVISIONING_PHASE_INDEX_SCAN", "PROVISIONING_PHASE_INDEX_SORT", "PROVISIONING_PHASE_MAIN"`, v)
	}
}

// Values returns all possible values for ProvisioningPhase.
//
// There is no guarantee on the order of the values in the slice.
func (f *ProvisioningPhase) Values() []ProvisioningPhase {
	return []ProvisioningPhase{
		ProvisioningPhaseProvisioningPhaseIndexScan,
		ProvisioningPhaseProvisioningPhaseIndexSort,
		ProvisioningPhaseProvisioningPhaseMain,
	}
}

// Type always returns ProvisioningPhase to satisfy [pflag.Value] interface
func (f *ProvisioningPhase) Type() string {
	return "ProvisioningPhase"
}

type RequestedClaims struct {
	PermissionSet RequestedClaimsPermissionSet `json:"permission_set,omitempty"`

	Resources []RequestedResource `json:"resources,omitempty"`
}

// Might add WRITE in the future
type RequestedClaimsPermissionSet string

const RequestedClaimsPermissionSetReadOnly RequestedClaimsPermissionSet = `READ_ONLY`

// String representation for [fmt.Print]
func (f *RequestedClaimsPermissionSet) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RequestedClaimsPermissionSet) Set(v string) error {
	switch v {
	case `READ_ONLY`:
		*f = RequestedClaimsPermissionSet(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ_ONLY"`, v)
	}
}

// Values returns all possible values for RequestedClaimsPermissionSet.
//
// There is no guarantee on the order of the values in the slice.
func (f *RequestedClaimsPermissionSet) Values() []RequestedClaimsPermissionSet {
	return []RequestedClaimsPermissionSet{
		RequestedClaimsPermissionSetReadOnly,
	}
}

// Type always returns RequestedClaimsPermissionSet to satisfy [pflag.Value] interface
func (f *RequestedClaimsPermissionSet) Type() string {
	return "RequestedClaimsPermissionSet"
}

type RequestedResource struct {
	TableName string `json:"table_name,omitempty"`

	UnspecifiedResourceName string `json:"unspecified_resource_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RequestedResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RequestedResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestartDatabaseEndpointRequest struct {
	BranchId string `json:"-" url:"-"`

	EndpointId string `json:"-" url:"-"`

	ProjectId string `json:"-" url:"-"`
}

// Next field marker: 18
type SyncedDatabaseTable struct {
	// Synced Table data synchronization status
	DataSynchronizationStatus *SyncedTableStatus `json:"data_synchronization_status,omitempty"`
	// The branch_id of the database branch associated with the table.
	DatabaseBranchId string `json:"database_branch_id,omitempty"`
	// Name of the target database instance. This is required when creating
	// synced database tables in standard catalogs. This is optional when
	// creating synced database tables in registered catalogs. If this field is
	// specified when creating synced database tables in registered catalogs,
	// the database instance name MUST match that of the registered catalog (or
	// the request will be rejected).
	DatabaseInstanceName string `json:"database_instance_name,omitempty"`
	// The project_id of the database project associated with the table.
	DatabaseProjectId string `json:"database_project_id,omitempty"`
	// The branch_id of the database branch associated with the table.
	EffectiveDatabaseBranchId string `json:"effective_database_branch_id,omitempty"`
	// The name of the database instance that this table is registered to. This
	// field is always returned, and for tables inside database catalogs is
	// inferred database instance associated with the catalog.
	EffectiveDatabaseInstanceName string `json:"effective_database_instance_name,omitempty"`
	// The project_id of the database project associated with the table.
	EffectiveDatabaseProjectId string `json:"effective_database_project_id,omitempty"`
	// The name of the logical database that this table is registered to.
	EffectiveLogicalDatabaseName string `json:"effective_logical_database_name,omitempty"`
	// Target Postgres database object (logical database) name for this table.
	//
	// When creating a synced table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a synced table in a standard catalog, this field is
	// required. In this scenario, specifying this field will allow targeting an
	// arbitrary postgres database. Note that this has implications for the
	// `create_database_objects_is_missing` field in `spec`.
	LogicalDatabaseName string `json:"logical_database_name,omitempty"`
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"name"`

	Spec *SyncedTableSpec `json:"spec,omitempty"`
	// Data serving REST API URL for this table
	TableServingUrl string `json:"table_serving_url,omitempty"`
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
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
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
	// The last source table Delta version that was successfully synced to the
	// synced table. The last source table Delta version that was synced to the
	// synced table. Only populated if the table is still synced and available
	// for serving.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. Only populated if the table is still
	// synced and available for serving.
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
	// The current phase of the data synchronization pipeline.
	ProvisioningPhase ProvisioningPhase `json:"provisioning_phase,omitempty"`
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

type SyncedTablePosition struct {
	DeltaTableSyncInfo *DeltaTableSyncInfo `json:"delta_table_sync_info,omitempty"`
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	SyncEndTimestamp string `json:"sync_end_timestamp,omitempty"`
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	SyncStartTimestamp string `json:"sync_start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTablePosition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTablePosition) MarshalJSON() ([]byte, error) {
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
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If existing_pipeline_id is defined, the synced table will be bin packed
	// into the existing pipeline referenced. This avoids creating a new
	// pipeline and allows sharing existing compute. In this case, the
	// scheduling_policy of this synced table must match the scheduling policy
	// of the existing pipeline.
	ExistingPipelineId string `json:"existing_pipeline_id,omitempty"`
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If new_pipeline_spec is defined, a new pipeline is created for this
	// synced table. The location pointed to is used to store intermediate files
	// (checkpoints, event logs etc). The caller must have write permissions to
	// create Delta tables in the specified catalog and schema. Again, note this
	// requires write permissions, whereas the source table only requires read
	// permissions.
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
	ContinuousUpdateStatus *SyncedTableContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	// The state of the synced table.
	DetailedState SyncedTableState `json:"detailed_state,omitempty"`

	FailedStatus *SyncedTableFailedStatus `json:"failed_status,omitempty"`
	// Summary of the last successful synchronization from source to
	// destination.
	//
	// Will always be present if there has been a successful sync. Even if the
	// most recent syncs have failed.
	//
	// Limitation: The only exception is if the synced table is doing a FULL
	// REFRESH, then the last sync information will not be available until the
	// full refresh is complete. This limitation will be addressed in a future
	// version.
	//
	// This top-level field is a convenience for consumers who want easy access
	// to last sync information without having to traverse detailed_status.
	LastSync *SyncedTablePosition `json:"last_sync,omitempty"`
	// A text description of the current state of the synced table.
	Message string `json:"message,omitempty"`
	// ID of the associated pipeline. The pipeline ID may have been provided by
	// the client (in the case of bin packing), or generated by the server (when
	// creating a new pipeline).
	PipelineId string `json:"pipeline_id,omitempty"`

	ProvisioningStatus *SyncedTableProvisioningStatus `json:"provisioning_status,omitempty"`

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
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
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

type UpdateDatabaseBranchRequest struct {
	BranchId string `json:"-" url:"-"`

	DatabaseBranch DatabaseBranch `json:"database_branch"`

	ProjectId string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateDatabaseCatalogRequest struct {
	// Note that updating a database catalog is not yet supported.
	DatabaseCatalog DatabaseCatalog `json:"database_catalog"`
	// The name of the catalog in UC.
	Name string `json:"-" url:"-"`
	// The list of fields to update. Setting this field is not yet supported.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateDatabaseEndpointRequest struct {
	BranchId string `json:"-" url:"-"`

	DatabaseEndpoint DatabaseEndpoint `json:"database_endpoint"`

	EndpointId string `json:"-" url:"-"`

	ProjectId string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateDatabaseInstanceRequest struct {
	DatabaseInstance DatabaseInstance `json:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible. To wipe out custom_tags, specify custom_tags in the
	// update_mask with an empty custom_tags map.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateDatabaseInstanceRoleRequest struct {
	DatabaseInstanceName string `json:"-" url:"database_instance_name,omitempty"`

	DatabaseInstanceRole DatabaseInstanceRole `json:"database_instance_role"`

	InstanceName string `json:"-" url:"-"`
	// The name of the role. This is the unique identifier for the role in an
	// instance.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateDatabaseProjectRequest struct {
	DatabaseProject DatabaseProject `json:"database_project"`

	ProjectId string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateSyncedDatabaseTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" url:"-"`
	// Note that updating a synced database table is not yet supported.
	SyncedTable SyncedDatabaseTable `json:"synced_table"`
	// The list of fields to update. Setting this field is not yet supported.
	UpdateMask string `json:"-" url:"update_mask"`
}
