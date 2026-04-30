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
	// Output only. The full resource path of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"name,omitempty"`
	// The project containing this branch (API resource hierarchy). Format:
	// projects/{project_id}
	//
	// Note: This field indicates where the branch exists in the resource
	// hierarchy. For point-in-time branching from another branch, see
	// `status.source_branch`.
	Parent string `json:"parent,omitempty"`
	// The spec contains the branch configuration.
	Spec *BranchSpec `json:"spec,omitempty"`
	// The current status of a Branch.
	Status *BranchStatus `json:"status,omitempty"`
	// System-generated unique ID for the branch.
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
	// Absolute expiration timestamp. When set, the branch will expire at this
	// time. Mutually exclusive with `ttl` and `no_expiry`. When updating, use
	// `spec.expiration` in the update_mask.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// When set to true, protects the branch from deletion and reset. Associated
	// compute endpoints and the project cannot be deleted while the branch is
	// protected.
	IsProtected bool `json:"is_protected,omitempty"`
	// Explicitly disable expiration. When set to true, the branch will not
	// expire. If set to false, the request is invalid; provide either ttl or
	// expire_time instead. Mutually exclusive with `expire_time` and `ttl`.
	// When updating, use `spec.expiration` in the update_mask.
	NoExpiry bool `json:"no_expiry,omitempty"`
	// The name of the source branch from which this branch was created (data
	// lineage for point-in-time recovery). If not specified, defaults to the
	// project's default branch. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch string `json:"source_branch,omitempty"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	SourceBranchLsn string `json:"source_branch_lsn,omitempty"`
	// The point in time on the source branch from which this branch was
	// created.
	SourceBranchTime *time.Time `json:"source_branch_time,omitempty"`
	// Relative time-to-live duration. When set, the branch will expire at
	// creation_time + ttl. Mutually exclusive with `expire_time` and
	// `no_expiry`. When updating, use `spec.expiration` in the update_mask.
	Ttl *duration.Duration `json:"ttl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BranchSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BranchSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BranchStatus struct {
	// The short identifier of the branch, suitable for showing to the users.
	// For a branch with name `projects/my-project/branches/my-branch`, the
	// branch_id is `my-branch`.
	//
	// Use this field when building UI components that display branches to users
	// (e.g., a drop-down selector). Prefer showing `branch_id` instead of the
	// full resource name from `Branch.name`, which follows the
	// `projects/{project_id}/branches/{branch_id}` format and is not
	// user-friendly.
	BranchId string `json:"branch_id,omitempty"`
	// The branch's state, indicating if it is initializing, ready for use, or
	// archived.
	CurrentState BranchStatusState `json:"current_state,omitempty"`
	// Whether the branch is the project's default branch.
	Default bool `json:"default,omitempty"`
	// Absolute expiration time for the branch. Empty if expiration is disabled.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
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

// The state of the branch.
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

type Catalog struct {
	// A timestamp indicating when the catalog was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. The full resource path of the catalog.
	//
	// Format: "catalogs/{catalog_id}".
	Name string `json:"name,omitempty"`
	// The desired state of the Catalog.
	Spec *CatalogCatalogSpec `json:"spec,omitempty"`
	// The observed state of the Catalog.
	Status *CatalogCatalogStatus `json:"status,omitempty"`
	// System-generated unique identifier for the catalog.
	Uid string `json:"uid,omitempty"`
	// A timestamp indicating when the catalog was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Catalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Catalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The desired state of the Catalog.
type CatalogCatalogSpec struct {
	// The resource path of the branch associated with the catalog.
	//
	// Format: projects/{project_id}/branches/{branch_id}.
	Branch string `json:"branch,omitempty"`
	// If set to true, the specified postgres_database is created on behalf of
	// the calling user if it does not already exist. In this case, the calling
	// user has a role created for them in Postgres if they do not already have
	// one.
	//
	// Defaults to false, meaning that the request fails if the specified
	// postgres_database does not already exist.
	CreateDatabaseIfMissing bool `json:"create_database_if_missing,omitempty"`
	// The name of the Postgres database inside the specified Lakebase project
	// and branch to be associated with the UC catalog. This database must
	// already exist, unless create_database_if_missing is set to true on
	// creation.
	//
	// A database can only be registered with one UC catalog at a time. To
	// re-register a database with a different catalog, the existing catalog
	// must be deleted first.
	//
	// A child branch inherits the fact of parent's registration. This means the
	// same-named database in a child branch cannot be registered with a second
	// catalog while the parent's registration exists. To allow registering the
	// database of a child branch, drop and recreate the database on the child
	// branch. This removes the fact of parent's registration from this branch
	// only.
	//
	// Doing Point In Time Restore (PITR) prior to the moment before the
	// Postgres DB was registered in the Catalog drops the fact of registration
	// of the database. So the user should avoid doing so.
	PostgresDatabase string `json:"postgres_database"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CatalogCatalogSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogCatalogSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The observed state of the Catalog.
type CatalogCatalogStatus struct {
	// The resource path of the branch associated with the catalog.
	//
	// Format: projects/{project_id}/branches/{branch_id}.
	Branch string `json:"branch,omitempty"`
	// The short identifier of the catalog, suitable for showing to the users.
	// For a catalog with name `catalogs/my-catalog`, the catalog_id is
	// `my-catalog`.
	//
	// Use this field when building UI components that display catalogs to users
	// (e.g., a drop-down selector). Prefer showing `catalog_id` instead of the
	// full resource name from `Catalog.name`, which follows the
	// `catalogs/{catalog_id}` format and is not user-friendly.
	CatalogId string `json:"catalog_id,omitempty"`
	// The name of the Postgres database associated with the catalog.
	PostgresDatabase string `json:"postgres_database,omitempty"`
	// The resource path of the project associated with the catalog.
	//
	// Format: projects/{project_id}.
	Project string `json:"project,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CatalogCatalogStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogCatalogStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CatalogOperationMetadata struct {
}

type CreateBranchRequest struct {
	// The Branch to create.
	Branch Branch `json:"branch"`
	// The ID to use for the Branch. This becomes the final component of the
	// branch's resource name. The ID is required and must be 1-63 characters
	// long, start with a lowercase letter, and contain only lowercase letters,
	// numbers, and hyphens. For example, `development` becomes
	// `projects/my-app/branches/development`.
	BranchId string `json:"-" url:"branch_id"`
	// The Project where this Branch will be created. Format:
	// projects/{project_id}
	Parent string `json:"-" url:"-"`
	// If true, update the branch if it already exists instead of returning an
	// error.
	ReplaceExisting bool `json:"-" url:"replace_existing,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBranchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBranchRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCatalogRequest struct {
	Catalog Catalog `json:"catalog"`
	// The ID in the Unity Catalog. It becomes the full resource name, for
	// example "my_catalog" becomes "catalogs/my_catalog".
	CatalogId string `json:"-" url:"catalog_id"`
}

type CreateDatabaseRequest struct {
	// The desired specification of a Database.
	Database Database `json:"database"`
	// The ID to use for the Database, which will become the final component of
	// the database's resource name. This ID becomes the database name in
	// postgres.
	//
	// This value should be 4-63 characters, and only use characters available
	// in DNS names, as defined by RFC-1123
	//
	// If database_id is not specified in the request, it is generated
	// automatically.
	DatabaseId string `json:"-" url:"database_id,omitempty"`
	// The Branch where this Database will be created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDatabaseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDatabaseRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpointRequest struct {
	// The Endpoint to create.
	Endpoint Endpoint `json:"endpoint"`
	// The ID to use for the Endpoint. This becomes the final component of the
	// endpoint's resource name. The ID is required and must be 1-63 characters
	// long, start with a lowercase letter, and contain only lowercase letters,
	// numbers, and hyphens. For example, `primary` becomes
	// `projects/my-app/branches/development/endpoints/primary`.
	EndpointId string `json:"-" url:"endpoint_id"`
	// The Branch where this Endpoint will be created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`
	// If true, update the endpoint if it already exists instead of returning an
	// error.
	ReplaceExisting bool `json:"-" url:"replace_existing,omitempty"`

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
	// The ID to use for the Project. This becomes the final component of the
	// project's resource name. The ID is required and must be 1-63 characters
	// long, start with a lowercase letter, and contain only lowercase letters,
	// numbers, and hyphens. For example, `my-app` becomes `projects/my-app`.
	ProjectId string `json:"-" url:"project_id"`
}

type CreateRoleRequest struct {
	// The Branch where this Role is created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`
	// The desired specification of a Role.
	Role Role `json:"role"`
	// The ID to use for the Role, which will become the final component of the
	// role's resource name. This ID becomes the role in Postgres.
	//
	// This value should be 4-63 characters, and valid characters are lowercase
	// letters, numbers, and hyphens, as defined by RFC 1123.
	//
	// If role_id is not specified in the request, it is generated
	// automatically.
	RoleId string `json:"-" url:"role_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateSyncedTableRequest struct {
	SyncedTable SyncedTable `json:"synced_table"`
	// The ID to use for the Synced Table. This becomes the final component of
	// the SyncedTable's resource name. ID is required and is the synced table
	// name, containing (catalog, schema, table) tuple. Elements of the tuple
	// are the UC entity names.
	//
	// Example: "{catalog}.{schema}.{table}"
	//
	// synced_table_id represents both of the following:
	//
	// 1. An online VIEW virtual table in the Unity Catalog accessible via the
	// Lakehouse Federation. 2. Postgres table named "{table}" in schema
	// "{schema}" in the connected Postgres database
	SyncedTableId string `json:"-" url:"synced_table_id"`
}

// Database represents a Postgres database within a Branch.
type Database struct {
	// A timestamp indicating when the database was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The resource name of the database. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Name string `json:"name,omitempty"`
	// The branch containing this database. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`
	// The desired state of the Database.
	Spec *DatabaseDatabaseSpec `json:"spec,omitempty"`
	// The observed state of the Database.
	Status *DatabaseDatabaseStatus `json:"status,omitempty"`
	// A timestamp indicating when the database was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Database) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Database) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseCredential struct {
	// Timestamp in UTC of when this credential expires.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// The OAuth token that can be used as a password when connecting to a
	// database.
	Token string `json:"token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseDatabaseSpec struct {
	// The name of the Postgres database.
	//
	// This expects a valid Postgres identifier as specified in the link below.
	// https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
	// Required when creating the Database.
	//
	// To rename, pass a valid postgres identifier when updating the Database.
	PostgresDatabase string `json:"postgres_database,omitempty"`
	// The name of the role that owns the database. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	//
	// To change the owner, pass valid existing Role name when updating the
	// Database
	//
	// A database always has an owner.
	Role string `json:"role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseDatabaseSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseDatabaseSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseDatabaseStatus struct {
	// The short identifier of the database, suitable for showing to the users.
	// For a database with name
	// `projects/my-project/branches/my-branch/databases/my-db`, the database_id
	// is `my-db`.
	//
	// Use this field when building UI components that display databases to
	// users (e.g., a drop-down selector). Prefer showing `database_id` instead
	// of the full resource name from `Database.name`, which follows the
	// `projects/{project_id}/branches/{branch_id}/databases/{database_id}`
	// format and is not user-friendly.
	DatabaseId string `json:"database_id,omitempty"`
	// The name of the Postgres database.
	PostgresDatabase string `json:"postgres_database,omitempty"`
	// The name of the role that owns the database. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Role string `json:"role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseDatabaseStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseDatabaseStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseOperationMetadata struct {
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
	// The full resource path of the branch to delete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
}

type DeleteCatalogRequest struct {
	// The full resource path of the catalog to delete.
	//
	// Format: "catalogs/{catalog_id}".
	Name string `json:"-" url:"-"`
}

type DeleteDatabaseRequest struct {
	// The resource name of the postgres database. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Name string `json:"-" url:"-"`
}

type DeleteEndpointRequest struct {
	// The full resource path of the endpoint to delete. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

type DeleteProjectRequest struct {
	// The full resource path of the project to delete. Format:
	// projects/{project_id}
	Name string `json:"-" url:"-"`
	// If true, permanently deletes the project (hard delete). If false or
	// unset, performs a soft delete.
	Purge bool `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteProjectRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteRoleRequest struct {
	// The full resource path of the role to delete. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
	// Reassign objects. If this is set, all objects owned by the role are
	// reassigned to the role specified in this parameter.
	//
	// NOTE: setting this requires spinning up a compute to succeed, since it
	// involves running SQL queries.
	ReassignOwnedTo string `json:"-" url:"reassign_owned_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteSyncedTableRequest struct {
	// The Full resource name of the synced table, of the format
	// "synced_tables/{catalog}.{schema}.{table}", where (catalog, schema,
	// table) are the UC entity names.
	Name string `json:"-" url:"-"`
}

type DeltaTableSyncInfo struct {
	// The timestamp when the above Delta version was committed in the source
	// Delta table. Note: This is the Delta commit time, not the time the data
	// was written to the synced table.
	DeltaCommitTime *time.Time `json:"delta_commit_time,omitempty"`
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

type Endpoint struct {
	// A timestamp indicating when the compute endpoint was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. The full resource path of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"name,omitempty"`
	// The branch containing this endpoint (API resource hierarchy). Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`
	// The spec contains the compute endpoint configuration, including
	// autoscaling limits, suspend timeout, and disabled state.
	Spec *EndpointSpec `json:"spec,omitempty"`
	// Current operational status of the compute endpoint.
	Status *EndpointStatus `json:"status,omitempty"`
	// System-generated unique ID for the endpoint.
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

type EndpointGroupSpec struct {
	// Whether to allow read-only connections to read-write endpoints. Only
	// relevant for read-write endpoints where size.max > 1.
	EnableReadableSecondaries bool `json:"enable_readable_secondaries,omitempty"`
	// The maximum number of computes in the endpoint group. Currently, this
	// must be equal to min. Set to 1 for single compute endpoints, to disable
	// HA. To manually suspend all computes in an endpoint group, set disabled
	// to true on the EndpointSpec.
	Max int `json:"max"`
	// The minimum number of computes in the endpoint group. Currently, this
	// must be equal to max. This must be greater than or equal to 1.
	Min int `json:"min"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointGroupSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointGroupSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointGroupStatus struct {
	// Whether read-only connections to read-write endpoints are allowed. Only
	// relevant if read replicas are configured by specifying size.max > 1.
	EnableReadableSecondaries bool `json:"enable_readable_secondaries,omitempty"`
	// The maximum number of computes in the endpoint group. Currently, this
	// must be equal to min. Set to 1 for single compute endpoints, to disable
	// HA. To manually suspend all computes in an endpoint group, set disabled
	// to true on the EndpointSpec.
	Max int `json:"max"`
	// The minimum number of computes in the endpoint group. Currently, this
	// must be equal to max. This must be greater than or equal to 1.
	Min int `json:"min"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointGroupStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointGroupStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Encapsulates various hostnames (r/w or r/o, pooled or not) for an endpoint.
type EndpointHosts struct {
	// The hostname to connect to this endpoint. For read-write endpoints, this
	// is a read-write hostname which connects to the primary compute. For
	// read-only endpoints, this is a read-only hostname which allows read-only
	// operations.
	Host string `json:"host,omitempty"`
	// An optionally defined read-only host for the endpoint, without pooling.
	// For read-only endpoints, this attribute is always defined and is
	// equivalent to host. For read-write endpoints, this attribute is defined
	// if the enclosing endpoint is a group with greater than 1 computes
	// configured, and has readable secondaries enabled.
	ReadOnlyHost string `json:"read_only_host,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointHosts) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointHosts) MarshalJSON() ([]byte, error) {
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
	// The maximum number of Compute Units. The maximum value is 64. The
	// difference between the minimum and maximum Compute Units (max - min) must
	// not exceed 16.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled bool `json:"disabled,omitempty"`
	// The endpoint type. A branch can only have one READ_WRITE endpoint.
	EndpointType EndpointType `json:"endpoint_type"`
	// Settings for optional HA configuration of the endpoint. If unspecified,
	// the endpoint defaults to non HA settings, with a single compute backing
	// the endpoint (and no readable secondaries for Read/Write endpoints).
	Group *EndpointGroupSpec `json:"group,omitempty"`
	// When set to true, explicitly disables automatic suspension (never
	// suspend). Should be set to true when provided. Mutually exclusive with
	// `suspend_timeout_duration`. When updating, use `spec.suspension` in the
	// update_mask.
	NoSuspension bool `json:"no_suspension,omitempty"`

	Settings *EndpointSettings `json:"settings,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended. If specified should be between 60s and 604800s (1 minute to 1
	// week). Mutually exclusive with `no_suspension`. When updating, use
	// `spec.suspension` in the update_mask.
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
	// The maximum number of Compute Units. The maximum value is 64. The
	// difference between the minimum and maximum Compute Units (max - min) must
	// not exceed 16.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`

	CurrentState EndpointStatusState `json:"current_state,omitempty"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled bool `json:"disabled,omitempty"`
	// The short identifier of the endpoint, suitable for showing to the users.
	// For an endpoint with name
	// `projects/my-project/branches/my-branch/endpoints/my-endpoint`, the
	// endpoint_id is `my-endpoint`.
	//
	// Use this field when building UI components that display endpoints to
	// users (e.g., a drop-down selector). Prefer showing `endpoint_id` instead
	// of the full resource name from `Endpoint.name`, which follows the
	// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`
	// format and is not user-friendly.
	EndpointId string `json:"endpoint_id,omitempty"`
	// The endpoint type. A branch can only have one READ_WRITE endpoint.
	EndpointType EndpointType `json:"endpoint_type,omitempty"`
	// Details on the HA configuration of the endpoint.
	Group *EndpointGroupStatus `json:"group,omitempty"`
	// Contains host information for connecting to the endpoint.
	Hosts *EndpointHosts `json:"hosts,omitempty"`

	PendingState EndpointStatusState `json:"pending_state,omitempty"`

	Settings *EndpointSettings `json:"settings,omitempty"`
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

const EndpointStatusStateDegraded EndpointStatusState = `DEGRADED`

const EndpointStatusStateIdle EndpointStatusState = `IDLE`

const EndpointStatusStateInit EndpointStatusState = `INIT`

// String representation for [fmt.Print]
func (f *EndpointStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStatusState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DEGRADED`, `IDLE`, `INIT`:
		*f = EndpointStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DEGRADED", "IDLE", "INIT"`, v)
	}
}

// Values returns all possible values for EndpointStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStatusState) Values() []EndpointStatusState {
	return []EndpointStatusState{
		EndpointStatusStateActive,
		EndpointStatusStateDegraded,
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

const EndpointTypeEndpointTypeReadOnly EndpointType = `ENDPOINT_TYPE_READ_ONLY`

const EndpointTypeEndpointTypeReadWrite EndpointType = `ENDPOINT_TYPE_READ_WRITE`

// String representation for [fmt.Print]
func (f *EndpointType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointType) Set(v string) error {
	switch v {
	case `ENDPOINT_TYPE_READ_ONLY`, `ENDPOINT_TYPE_READ_WRITE`:
		*f = EndpointType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ENDPOINT_TYPE_READ_ONLY", "ENDPOINT_TYPE_READ_WRITE"`, v)
	}
}

// Values returns all possible values for EndpointType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointType) Values() []EndpointType {
	return []EndpointType{
		EndpointTypeEndpointTypeReadOnly,
		EndpointTypeEndpointTypeReadWrite,
	}
}

// Type always returns EndpointType to satisfy [pflag.Value] interface
func (f *EndpointType) Type() string {
	return "EndpointType"
}

// Error codes returned by Databricks APIs to indicate specific failure
// conditions.
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

type GenerateDatabaseCredentialRequest struct {
	// The returned token will be scoped to UC tables with the specified
	// permissions.
	Claims []RequestedClaims `json:"claims,omitempty"`
	// This field is not yet supported. The endpoint for which this credential
	// will be generated. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint string `json:"endpoint"`
}

type GetBranchRequest struct {
	// The full resource path of the branch to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
}

type GetCatalogRequest struct {
	// The full resource path of the catalog to retrieve.
	//
	// Format: "catalogs/{catalog_id}".
	Name string `json:"-" url:"-"`
}

type GetDatabaseRequest struct {
	// The name of the Database to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Name string `json:"-" url:"-"`
}

type GetEndpointRequest struct {
	// The full resource path of the endpoint to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

type GetProjectRequest struct {
	// The full resource path of the project to retrieve. Format:
	// projects/{project_id}
	Name string `json:"-" url:"-"`
}

type GetRoleRequest struct {
	// The full resource path of the role to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
}

type GetSyncedTableRequest struct {
	// The Full resource name of the synced table. Format:
	// "synced_tables/{catalog}.{schema}.{table}", where (catalog, schema,
	// table) are the entity names in the Unity Catalog.
	Name string `json:"-" url:"-"`
}

// Configuration for the initial Read/Write endpoint created during project
// creation.
type InitialEndpointSpec struct {
	// Settings for HA configuration of the endpoint.
	Group *EndpointGroupSpec `json:"group,omitempty"`
}

type ListBranchesRequest struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
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
	// List of branches in the project.
	Branches []Branch `json:"branches,omitempty"`
	// Token to request the next page of branches.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBranchesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBranchesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabasesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Databases. Requests first page
	// if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Branch that owns this collection of databases. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabasesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabasesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabasesResponse struct {
	// List of databases.
	Databases []Database `json:"databases,omitempty"`
	// Pagination token to request the next page of databases.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabasesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabasesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsRequest struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
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
	// List of compute endpoints in the branch.
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// Token to request the next page of compute endpoints.
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
	// Upper bound for items returned. Cannot be negative. The maximum value is
	// 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Whether to include soft-deleted projects in the response. When true,
	// soft-deleted projects are included alongside active projects.
	// Hard-deleted and already-purged projects are never returned.
	ShowDeleted bool `json:"-" url:"show_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProjectsResponse struct {
	// Token to request the next page of projects.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all projects in the workspace that the user has permission to
	// access.
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
	// Upper bound for items returned. Cannot be negative.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
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
	// Token to request the next page of Postgres roles.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Postgres roles in the branch.
	Roles []Role `json:"roles,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRolesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRolesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NewPipelineSpec struct {
	// Budget policy to set on the newly created pipeline.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
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
	// A timestamp indicating when the project was soft-deleted. Empty if the
	// project is not deleted, otherwise set to a timestamp in the past.
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// Configuration settings for the initial Read/Write endpoint created inside
	// the initial branch for a newly created project. If omitted, the initial
	// endpoint created will have default settings, without high availability
	// configured. This field does not apply to any endpoints created after
	// project creation. Use spec.default_endpoint_settings to configure default
	// settings for endpoints created after project creation.
	InitialEndpointSpec *InitialEndpointSpec `json:"initial_endpoint_spec,omitempty"`
	// Output only. The full resource path of the project. Format:
	// projects/{project_id}
	Name string `json:"name,omitempty"`
	// A timestamp indicating when the project is scheduled for permanent
	// deletion. Empty if the project is not deleted, otherwise set to a
	// timestamp in the future.
	PurgeTime *time.Time `json:"purge_time,omitempty"`
	// The spec contains the project configuration, including display_name,
	// pg_version (Postgres version), history_retention_duration, and
	// default_endpoint_settings.
	Spec *ProjectSpec `json:"spec,omitempty"`
	// The current status of a Project.
	Status *ProjectStatus `json:"status,omitempty"`
	// System-generated unique ID for the project.
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

type ProjectCustomTag struct {
	// The key of the custom tag.
	Key string `json:"key,omitempty"`
	// The value of the custom tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectCustomTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectCustomTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A collection of settings for a compute endpoint.
type ProjectDefaultEndpointSettings struct {
	// The maximum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// When set to true, explicitly disables automatic suspension (never
	// suspend). Should be set to true when provided. Mutually exclusive with
	// `suspend_timeout_duration`. When updating, use
	// `spec.project_default_settings.suspension` in the update_mask.
	NoSuspension bool `json:"no_suspension,omitempty"`
	// A raw representation of Postgres settings.
	PgSettings map[string]string `json:"pg_settings,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended. If specified should be between 60s and 604800s (1 minute to 1
	// week). Mutually exclusive with `no_suspension`. When updating, use
	// `spec.project_default_settings.suspension` in the update_mask.
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

type ProjectSpec struct {
	// The desired budget policy to associate with the project. See
	// status.budget_policy_id for the policy that is actually applied to the
	// project.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Custom tags to associate with the project. Forwarded to LBM for billing
	// and cost tracking. To update tags, provide the new tag list and include
	// "spec.custom_tags" in the update_mask. To clear all tags, provide an
	// empty list and include "spec.custom_tags" in the update_mask. To preserve
	// existing tags, omit this field from the update_mask (or use wildcard "*"
	// which auto-excludes empty tags).
	CustomTags []ProjectCustomTag `json:"custom_tags,omitempty"`
	// The full resource path for the default branch of the project Format:
	// projects/{project_id}/branches/{branch_id}
	DefaultBranch string `json:"default_branch,omitempty"`

	DefaultEndpointSettings *ProjectDefaultEndpointSettings `json:"default_endpoint_settings,omitempty"`
	// Human-readable project name. Length should be between 1 and 256
	// characters.
	DisplayName string `json:"display_name,omitempty"`
	// Whether to enable PG native password login on all endpoints in this
	// project. Defaults to true.
	EnablePgNativeLogin bool `json:"enable_pg_native_login,omitempty"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project. Value should be between
	// 172800s (2 days) and 3024000s (35 days).
	HistoryRetentionDuration *duration.Duration `json:"history_retention_duration,omitempty"`
	// The major Postgres version number. Supported versions are 16 and 17.
	PgVersion int `json:"pg_version,omitempty"`

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
	// The budget policy that is applied to the project.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The effective custom tags associated with the project.
	CustomTags []ProjectCustomTag `json:"custom_tags,omitempty"`
	// The full resource path of the default branch of the project
	DefaultBranch string `json:"default_branch,omitempty"`
	// The effective default endpoint settings.
	DefaultEndpointSettings *ProjectDefaultEndpointSettings `json:"default_endpoint_settings,omitempty"`
	// The effective human-readable project name.
	DisplayName string `json:"display_name,omitempty"`
	// Whether to enable PG native password login on all endpoints in this
	// project.
	EnablePgNativeLogin bool `json:"enable_pg_native_login,omitempty"`
	// The effective number of seconds to retain the shared history for point in
	// time recovery.
	HistoryRetentionDuration *duration.Duration `json:"history_retention_duration,omitempty"`
	// The email of the project owner.
	Owner string `json:"owner,omitempty"`
	// The effective major Postgres version number.
	PgVersion int `json:"pg_version,omitempty"`
	// The short identifier of the project, suitable for showing to the users.
	// For a project with name `projects/my-project`, the project_id is
	// `my-project`.
	//
	// Use this field when building UI components that display projects to users
	// (e.g., a drop-down selector). Prefer showing `project_id` instead of the
	// full resource name from `Project.name`, which follows the
	// `projects/{project_id}` format and is not user-friendly.
	ProjectId string `json:"project_id,omitempty"`
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

// The current phase of the data synchronization pipeline.
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
	// The full Unity Catalog table name.
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

// Role represents a Postgres role within a Branch.
type Role struct {
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. The full resource path of the role. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"name,omitempty"`
	// The Branch where this Role exists. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`
	// The spec contains the role configuration, including identity type,
	// authentication method, and role attributes.
	Spec *RoleRoleSpec `json:"spec,omitempty"`
	// Current status of the role, including its identity type, authentication
	// method, and role attributes.
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

// Attributes that can be granted to a Postgres role. We are only implementing a
// subset for now, see xref:
// https://www.postgresql.org/docs/16/sql-createrole.html The values follow
// Postgres keyword naming e.g. CREATEDB, BYPASSRLS, etc. which is why they
// don't include typical underscores between words.
type RoleAttributes struct {
	Bypassrls bool `json:"bypassrls,omitempty"`

	Createdb bool `json:"createdb,omitempty"`

	Createrole bool `json:"createrole,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RoleAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RoleAttributes) MarshalJSON() ([]byte, error) {
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

// Roles that the DatabaseInstanceRole can be a member of.
type RoleMembershipRole string

const RoleMembershipRoleDatabricksSuperuser RoleMembershipRole = `DATABRICKS_SUPERUSER`

// String representation for [fmt.Print]
func (f *RoleMembershipRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RoleMembershipRole) Set(v string) error {
	switch v {
	case `DATABRICKS_SUPERUSER`:
		*f = RoleMembershipRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS_SUPERUSER"`, v)
	}
}

// Values returns all possible values for RoleMembershipRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *RoleMembershipRole) Values() []RoleMembershipRole {
	return []RoleMembershipRole{
		RoleMembershipRoleDatabricksSuperuser,
	}
}

// Type always returns RoleMembershipRole to satisfy [pflag.Value] interface
func (f *RoleMembershipRole) Type() string {
	return "RoleMembershipRole"
}

type RoleOperationMetadata struct {
}

type RoleRoleSpec struct {
	// The desired API-exposed Postgres role attribute to associate with the
	// role. Optional.
	Attributes *RoleAttributes `json:"attributes,omitempty"`
	// Controls how the Postgres role authenticates when a client opens a
	// database connection. Supported values:
	//
	// * LAKEBASE_OAUTH_V1: the role authenticates by presenting a Databricks
	// OAuth access token derived from the backing managed identity (the
	// Databricks user, service principal, or group named by the role's
	// `postgres_role`). No static password exists for roles using this method.
	// * PG_PASSWORD_SCRAM_SHA_256: the role authenticates with a Postgres
	// password verified server-side using the SCRAM-SHA-256 mechanism. Lakebase
	// generates a password for the role. * NO_LOGIN: the role cannot open a
	// Postgres session at all. Useful for roles that exist only to own objects
	// or to aggregate privileges that are then granted to other, loginable
	// roles.
	//
	// If auth_method is left unspecified, a meaningful authentication method is
	// derived from the identity_type: * For the managed identities, OAUTH is
	// used. * For the regular postgres roles, authentication based on postgres
	// passwords is used.
	//
	// NOTE: for the Databricks identity type GROUP, LAKEBASE_OAUTH_V1 is the
	// default auth method (group can login as well).
	AuthMethod RoleAuthMethod `json:"auth_method,omitempty"`
	// The type of role. When specifying a managed-identity, the chosen role_id
	// must be a valid:
	//
	// * application ID for SERVICE_PRINCIPAL * user email for USER * group name
	// for GROUP
	IdentityType RoleIdentityType `json:"identity_type,omitempty"`
	// An enum value for a standard role that this role is a member of.
	MembershipRoles []RoleMembershipRole `json:"membership_roles,omitempty"`
	// The name of the Postgres role.
	//
	// This expects a valid Postgres identifier as specified in the link below.
	// https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
	//
	// Required when creating the Role.
	//
	// If you wish to create a Postgres Role backed by a managed Databricks
	// identity, then postgres_role must be one of the following:
	//
	// 1. user email for IdentityType.USER 2. app ID for
	// IdentityType.SERVICE_PRINCIPAL 2. group name for IdentityType.GROUP
	PostgresRole string `json:"postgres_role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RoleRoleSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RoleRoleSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RoleRoleStatus struct {
	// The PG role attributes associated with the role.
	Attributes *RoleAttributes `json:"attributes,omitempty"`

	AuthMethod RoleAuthMethod `json:"auth_method,omitempty"`
	// The type of the role.
	IdentityType RoleIdentityType `json:"identity_type,omitempty"`
	// An enum value for a standard role that this role is a member of.
	MembershipRoles []RoleMembershipRole `json:"membership_roles,omitempty"`
	// The name of the Postgres role.
	PostgresRole string `json:"postgres_role,omitempty"`
	// The short identifier of the role, suitable for showing to the users. For
	// a role with name `projects/my-project/branches/my-branch/roles/my-role`,
	// the role_id is `my-role`.
	//
	// Use this field when building UI components that display roles to users
	// (e.g., a drop-down selector). Prefer showing `role_id` instead of the
	// full resource name from `Role.name`, which follows the
	// `projects/{project_id}/branches/{branch_id}/roles/{role_id}` format and
	// is not user-friendly.
	RoleId string `json:"role_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RoleRoleStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RoleRoleStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SyncedTable struct {
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. The Full resource name of the synced table in Postgres where
	// (catalog, schema, table) are the UC entity names.
	//
	// Format "synced_tables/{catalog}.{schema}.{table}"
	//
	// For the corresponding source table in the Unity catalog look for the
	// "source_table_full_name" attribute.
	Name string `json:"name,omitempty"`
	// Configuration details of the synced table, such as the source table,
	// scheduling policy, etc. This attribute is specified at creation time and
	// most fields are returned as is on subsequent queries.
	Spec *SyncedTableSyncedTableSpec `json:"spec,omitempty"`
	// Synced Table data synchronization status.
	Status *SyncedTableSyncedTableStatus `json:"status,omitempty"`
	// The Unity Catalog table ID for this synced table.
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata for SyncedTable long-running operations.
type SyncedTableOperationMetadata struct {
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

type SyncedTablePosition struct {
	DeltaTableSyncInfo *DeltaTableSyncInfo `json:"delta_table_sync_info,omitempty"`
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	SyncEndTime *time.Time `json:"sync_end_time,omitempty"`
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	SyncStartTime *time.Time `json:"sync_start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTablePosition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTablePosition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of a synced table.
type SyncedTableState string

const SyncedTableStateSyncedTableOffline SyncedTableState = `SYNCED_TABLE_OFFLINE`

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
	case `SYNCED_TABLE_OFFLINE`, `SYNCED_TABLE_OFFLINE_FAILED`, `SYNCED_TABLE_ONLINE`, `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`, `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`, `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`, `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`, `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`, `SYNCED_TABLE_PROVISIONING`, `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`, `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`:
		*f = SyncedTableState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SYNCED_TABLE_OFFLINE", "SYNCED_TABLE_OFFLINE_FAILED", "SYNCED_TABLE_ONLINE", "SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE", "SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE", "SYNCED_TABLE_ONLINE_PIPELINE_FAILED", "SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE", "SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES", "SYNCED_TABLE_PROVISIONING", "SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT", "SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Values returns all possible values for SyncedTableState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableState) Values() []SyncedTableState {
	return []SyncedTableState{
		SyncedTableStateSyncedTableOffline,
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

type SyncedTableSyncedTableSpec struct {
	// The full resource name the branch associated with the table.
	//
	// Format: "projects/{project_id}/branches/{branch_id}".
	Branch string `json:"branch,omitempty"`
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist. The request will fail if
	// this is false and the database/schema do not exist.
	//
	// Defaults to true if omitted.
	CreateDatabaseObjectsIfMissing bool `json:"create_database_objects_if_missing,omitempty"`
	// ID of an existing pipeline to bin-pack this synced table into. At most
	// one of existing_pipeline_id and new_pipeline_spec should be defined.
	//
	// The pipeline used for the synced table is returned via the top level
	// pipeline_id attribute.
	ExistingPipelineId string `json:"existing_pipeline_id,omitempty"`
	// Specification for creating a new pipeline. At most one of
	// existing_pipeline_id and new_pipeline_spec should be defined.
	//
	// The pipeline used for the synced table is returned via the top level
	// pipeline_id attribute.
	NewPipelineSpec *NewPipelineSpec `json:"new_pipeline_spec,omitempty"`
	// The Postgres database name where the synced table will be created in.
	//
	// If this synced table is created inside a Lakebase Catalog, this attribute
	// can be omitted on creation and is inferred from the postgres_database
	// associated with the Lakebase Catalog. If specified when inside a Lakebase
	// Catalog, the value must match.
	//
	// A value must be specified when creating a synced table inside a Standard
	// Catalog.
	PostgresDatabase string `json:"postgres_database,omitempty"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Scheduling policy of the underlying pipeline.
	SchedulingPolicy SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy `json:"scheduling_policy,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	//
	// For the corresponding destination table, use any of the two:
	//
	// * synced_table_id used at the creation of the SyncedTable * "name"
	// consisting of "synced_tables/" prefix and the full name of the
	// destination table.
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey string `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSyncedTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSyncedTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Scheduling policy of the synced table's underlying pipeline.
type SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy string

const SyncedTableSyncedTableSpecSyncedTableSchedulingPolicyContinuous SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy = `CONTINUOUS`

const SyncedTableSyncedTableSpecSyncedTableSchedulingPolicySnapshot SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy = `SNAPSHOT`

const SyncedTableSyncedTableSpecSyncedTableSchedulingPolicyTriggered SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy = `TRIGGERED`

// String representation for [fmt.Print]
func (f *SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`:
		*f = SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "SNAPSHOT", "TRIGGERED"`, v)
	}
}

// Values returns all possible values for SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy) Values() []SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy {
	return []SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy{
		SyncedTableSyncedTableSpecSyncedTableSchedulingPolicyContinuous,
		SyncedTableSyncedTableSpecSyncedTableSchedulingPolicySnapshot,
		SyncedTableSyncedTableSpecSyncedTableSchedulingPolicyTriggered,
	}
}

// Type always returns SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy to satisfy [pflag.Value] interface
func (f *SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy) Type() string {
	return "SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy"
}

type SyncedTableSyncedTableStatus struct {
	// The state of the synced table.
	DetailedState SyncedTableState `json:"detailed_state,omitempty"`
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// Summary of the last successful synchronization from source to
	// destination.
	LastSync *SyncedTablePosition `json:"last_sync,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	LastSyncTime *time.Time `json:"last_sync_time,omitempty"`
	// A text description of the current state of the synced table.
	Message string `json:"message,omitempty"`

	OngoingSyncProgress *SyncedTablePipelineProgress `json:"ongoing_sync_progress,omitempty"`
	// ID of the associated pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// The full resource name of the project associated with the table.
	//
	// Format: "projects/{project_id}".
	Project string `json:"project,omitempty"`
	// The current phase of the data synchronization pipeline.
	ProvisioningPhase ProvisioningPhase `json:"provisioning_phase,omitempty"`
	// The provisioning state of the synced table entity in Unity Catalog.
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSyncedTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSyncedTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to restore a soft-deleted project within its retention period.
type UndeleteProjectRequest struct {
	// The full resource path of the project to undelete. Format:
	// projects/{project_id}
	Name string `json:"-" url:"-"`
}

type UpdateBranchRequest struct {
	// The Branch to update.
	//
	// The branch's `name` field is used to identify the branch to update.
	// Format: projects/{project_id}/branches/{branch_id}
	Branch Branch `json:"branch"`
	// Output only. The full resource path of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

type UpdateDatabaseRequest struct {
	// The Database to update.
	//
	// The database's `name` field is used to identify the database to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Database Database `json:"database"`
	// The resource name of the database. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
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
	// Output only. The full resource path of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

type UpdateProjectRequest struct {
	// Output only. The full resource path of the project. Format:
	// projects/{project_id}
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

type UpdateRoleRequest struct {
	// Output only. The full resource path of the role. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
	// The Postgres Role to update.
	//
	// The role's `name` field is used to identify the role to update. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Role Role `json:"role"`
	// The list of fields to update in Postgres Role. If unspecified, all fields
	// will be updated when possible.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}
