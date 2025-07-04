// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type CreateDatabaseCatalogRequest struct {

	// Wire name: 'catalog'
	Catalog DatabaseCatalog `json:"catalog"`
}

func (st CreateDatabaseCatalogRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createDatabaseCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createDatabaseCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateDatabaseCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := createDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateDatabaseInstanceRequest struct {
	// Instance to create.
	// Wire name: 'database_instance'
	DatabaseInstance DatabaseInstance `json:"database_instance"`
}

func (st CreateDatabaseInstanceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createDatabaseInstanceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createDatabaseInstanceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	pb, err := createDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateDatabaseInstanceRoleRequest struct {

	// Wire name: 'database_instance_role'
	DatabaseInstanceRole DatabaseInstanceRole `json:"database_instance_role"`

	InstanceName string `json:"-" tf:"-"`
}

func (st CreateDatabaseInstanceRoleRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createDatabaseInstanceRoleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createDatabaseInstanceRoleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	pb, err := createDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateDatabaseTableRequest struct {

	// Wire name: 'table'
	Table DatabaseTable `json:"table"`
}

func (st CreateDatabaseTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createDatabaseTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := createDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateSyncedDatabaseTableRequest struct {

	// Wire name: 'synced_table'
	SyncedTable SyncedDatabaseTable `json:"synced_table"`
}

func (st CreateSyncedDatabaseTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createSyncedDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createSyncedDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateSyncedDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := createSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DatabaseCatalog struct {

	// Wire name: 'create_database_if_not_exists'
	CreateDatabaseIfNotExists bool `json:"create_database_if_not_exists,omitempty"`
	// The name of the DatabaseInstance housing the database.
	// Wire name: 'database_instance_name'
	DatabaseInstanceName string `json:"database_instance_name"`
	// The name of the database (in a instance) associated with the catalog.
	// Wire name: 'database_name'
	DatabaseName string `json:"database_name"`
	// The name of the catalog in UC.
	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'uid'
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseCatalog) EncodeValues(key string, v *url.Values) error {
	pb, err := databaseCatalogToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DatabaseCatalog) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databaseCatalogPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databaseCatalogFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabaseCatalog) MarshalJSON() ([]byte, error) {
	pb, err := databaseCatalogToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DatabaseCredential struct {

	// Wire name: 'expiration_time'
	ExpirationTime string `json:"expiration_time,omitempty"`

	// Wire name: 'token'
	Token string `json:"token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseCredential) EncodeValues(key string, v *url.Values) error {
	pb, err := databaseCredentialToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DatabaseCredential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databaseCredentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databaseCredentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabaseCredential) MarshalJSON() ([]byte, error) {
	pb, err := databaseCredentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	// Wire name: 'capacity'
	Capacity string `json:"capacity,omitempty"`
	// The refs of the child instances. This is only available if the instance
	// is parent instance.
	// Wire name: 'child_instance_refs'
	ChildInstanceRefs []DatabaseInstanceRef `json:"child_instance_refs,omitempty"`
	// The timestamp when the instance was created.
	// Wire name: 'creation_time'
	CreationTime string `json:"creation_time,omitempty"`
	// The email of the creator of the instance.
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// xref AIP-129. `enable_readable_secondaries` is owned by the client, while
	// `effective_enable_readable_secondaries` is owned by the server.
	// `enable_readable_secondaries` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_enable_readable_secondaries` on the other hand will always bet
	// set in all response messages (Create/Update/Get/List).
	// Wire name: 'effective_enable_readable_secondaries'
	EffectiveEnableReadableSecondaries bool `json:"effective_enable_readable_secondaries,omitempty"`
	// xref AIP-129. `node_count` is owned by the client, while
	// `effective_node_count` is owned by the server. `node_count` will only be
	// set in Create/Update response messages if and only if the user provides
	// the field via the request. `effective_node_count` on the other hand will
	// always bet set in all response messages (Create/Update/Get/List).
	// Wire name: 'effective_node_count'
	EffectiveNodeCount int `json:"effective_node_count,omitempty"`
	// xref AIP-129. `retention_window_in_days` is owned by the client, while
	// `effective_retention_window_in_days` is owned by the server.
	// `retention_window_in_days` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_retention_window_in_days` on the other hand will always bet
	// set in all response messages (Create/Update/Get/List).
	// Wire name: 'effective_retention_window_in_days'
	EffectiveRetentionWindowInDays int `json:"effective_retention_window_in_days,omitempty"`
	// xref AIP-129. `stopped` is owned by the client, while `effective_stopped`
	// is owned by the server. `stopped` will only be set in Create/Update
	// response messages if and only if the user provides the field via the
	// request. `effective_stopped` on the other hand will always bet set in all
	// response messages (Create/Update/Get/List).
	// Wire name: 'effective_stopped'
	EffectiveStopped bool `json:"effective_stopped,omitempty"`
	// Whether to enable secondaries to serve read-only traffic. Defaults to
	// false.
	// Wire name: 'enable_readable_secondaries'
	EnableReadableSecondaries bool `json:"enable_readable_secondaries,omitempty"`
	// The name of the instance. This is the unique identifier for the instance.
	// Wire name: 'name'
	Name string `json:"name"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries.
	// Wire name: 'node_count'
	NodeCount int `json:"node_count,omitempty"`
	// The ref of the parent instance. This is only available if the instance is
	// child instance. Input: For specifying the parent instance to create a
	// child instance. Optional. Output: Only populated if provided as input to
	// create a child instance.
	// Wire name: 'parent_instance_ref'
	ParentInstanceRef *DatabaseInstanceRef `json:"parent_instance_ref,omitempty"`
	// The version of Postgres running on the instance.
	// Wire name: 'pg_version'
	PgVersion string `json:"pg_version,omitempty"`
	// The DNS endpoint to connect to the instance for read only access. This is
	// only available if enable_readable_secondaries is true.
	// Wire name: 'read_only_dns'
	ReadOnlyDns string `json:"read_only_dns,omitempty"`
	// The DNS endpoint to connect to the instance for read+write access.
	// Wire name: 'read_write_dns'
	ReadWriteDns string `json:"read_write_dns,omitempty"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained. The default value is 7 days.
	// Valid values are 2 to 35 days.
	// Wire name: 'retention_window_in_days'
	RetentionWindowInDays int `json:"retention_window_in_days,omitempty"`
	// The current state of the instance.
	// Wire name: 'state'
	State DatabaseInstanceState `json:"state,omitempty"`
	// Whether the instance is stopped.
	// Wire name: 'stopped'
	Stopped bool `json:"stopped,omitempty"`
	// An immutable UUID identifier for the instance.
	// Wire name: 'uid'
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseInstance) EncodeValues(key string, v *url.Values) error {
	pb, err := databaseInstanceToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DatabaseInstance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databaseInstancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databaseInstanceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabaseInstance) MarshalJSON() ([]byte, error) {
	pb, err := databaseInstanceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'branch_time'
	BranchTime string `json:"branch_time,omitempty"`
	// xref AIP-129. `lsn` is owned by the client, while `effective_lsn` is
	// owned by the server. `lsn` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_lsn` on the other hand will always bet set in all response
	// messages (Create/Update/Get/List). For a parent ref instance, this is the
	// LSN on the parent instance from which the instance was created. For a
	// child ref instance, this is the LSN on the instance from which the child
	// instance was created.
	// Wire name: 'effective_lsn'
	EffectiveLsn string `json:"effective_lsn,omitempty"`
	// User-specified WAL LSN of the ref database instance.
	//
	// Input: For specifying the WAL LSN to create a child instance. Optional.
	// Output: Only populated if provided as input to create a child instance.
	// Wire name: 'lsn'
	Lsn string `json:"lsn,omitempty"`
	// Name of the ref database instance.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Id of the ref database instance.
	// Wire name: 'uid'
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseInstanceRef) EncodeValues(key string, v *url.Values) error {
	pb, err := databaseInstanceRefToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DatabaseInstanceRef) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databaseInstanceRefPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databaseInstanceRefFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabaseInstanceRef) MarshalJSON() ([]byte, error) {
	pb, err := databaseInstanceRefToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A DatabaseInstanceRole represents a Postgres role in a database instance.
type DatabaseInstanceRole struct {
	// API-exposed Postgres role attributes
	// Wire name: 'attributes'
	Attributes *DatabaseInstanceRoleAttributes `json:"attributes,omitempty"`
	// The type of the role.
	// Wire name: 'identity_type'
	IdentityType DatabaseInstanceRoleIdentityType `json:"identity_type,omitempty"`
	// An enum value for a standard role that this role is a member of.
	// Wire name: 'membership_role'
	MembershipRole DatabaseInstanceRoleMembershipRole `json:"membership_role,omitempty"`
	// The name of the role. This is the unique identifier for the role in an
	// instance.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseInstanceRole) EncodeValues(key string, v *url.Values) error {
	pb, err := databaseInstanceRoleToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DatabaseInstanceRole) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databaseInstanceRolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databaseInstanceRoleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabaseInstanceRole) MarshalJSON() ([]byte, error) {
	pb, err := databaseInstanceRoleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Attributes that can be granted to a Postgres role. We are only implementing a
// subset for now, see xref:
// https://www.postgresql.org/docs/16/sql-createrole.html The values follow
// Postgres keyword naming e.g. CREATEDB, BYPASSRLS, etc. which is why they
// don't include typical underscores between words. We were requested to make
// this a nested object/struct representation since these are knobs from an
// external spec.
type DatabaseInstanceRoleAttributes struct {

	// Wire name: 'bypassrls'
	Bypassrls bool `json:"bypassrls,omitempty"`

	// Wire name: 'createdb'
	Createdb bool `json:"createdb,omitempty"`

	// Wire name: 'createrole'
	Createrole bool `json:"createrole,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseInstanceRoleAttributes) EncodeValues(key string, v *url.Values) error {
	pb, err := databaseInstanceRoleAttributesToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DatabaseInstanceRoleAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databaseInstanceRoleAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databaseInstanceRoleAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabaseInstanceRoleAttributes) MarshalJSON() ([]byte, error) {
	pb, err := databaseInstanceRoleAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Next field marker: 13
type DatabaseTable struct {
	// Name of the target database instance. This is required when creating
	// database tables in standard catalogs. This is optional when creating
	// database tables in registered catalogs. If this field is specified when
	// creating database tables in registered catalogs, the database instance
	// name MUST match that of the registered catalog (or the request will be
	// rejected).
	// Wire name: 'database_instance_name'
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
	// Wire name: 'logical_database_name'
	LogicalDatabaseName string `json:"logical_database_name,omitempty"`
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string `json:"name"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseTable) EncodeValues(key string, v *url.Values) error {
	pb, err := databaseTableToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DatabaseTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databaseTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databaseTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabaseTable) MarshalJSON() ([]byte, error) {
	pb, err := databaseTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDatabaseCatalogRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st DeleteDatabaseCatalogRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDatabaseCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDatabaseCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDatabaseCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDatabaseCatalogResponse struct {
}

func (st DeleteDatabaseCatalogResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDatabaseCatalogResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDatabaseCatalogResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDatabaseCatalogResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDatabaseCatalogResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDatabaseCatalogResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDatabaseCatalogResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDatabaseInstanceRequest struct {
	// By default, a instance cannot be deleted if it has descendant instances
	// created via PITR. If this flag is specified as true, all descendent
	// instances will be deleted as well.
	Force bool `json:"-" tf:"-"`
	// Name of the instance to delete.
	Name string `json:"-" tf:"-"`
	// Note purge=false is in development. If false, the database instance is
	// soft deleted (implementation pending). Soft deleted instances behave as
	// if they are deleted, and cannot be used for CRUD operations nor connected
	// to. However they can be undeleted by calling the undelete API for a
	// limited time (implementation pending). If true, the database instance is
	// hard deleted and cannot be undeleted. For the time being, setting this
	// value to true is required to delete an instance (soft delete is not yet
	// supported).
	Purge bool `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDatabaseInstanceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDatabaseInstanceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDatabaseInstanceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDatabaseInstanceResponse struct {
}

func (st DeleteDatabaseInstanceResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDatabaseInstanceResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDatabaseInstanceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDatabaseInstanceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDatabaseInstanceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDatabaseInstanceResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDatabaseInstanceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDatabaseInstanceRoleRequest struct {
	// This is the AIP standard name for the equivalent of Postgres' `IF EXISTS`
	// option
	AllowMissing bool `json:"-" tf:"-"`

	InstanceName string `json:"-" tf:"-"`

	Name string `json:"-" tf:"-"`

	ReassignOwnedTo string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDatabaseInstanceRoleRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDatabaseInstanceRoleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDatabaseInstanceRoleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDatabaseInstanceRoleResponse struct {
}

func (st DeleteDatabaseInstanceRoleResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDatabaseInstanceRoleResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDatabaseInstanceRoleResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDatabaseInstanceRoleResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDatabaseInstanceRoleResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDatabaseInstanceRoleResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDatabaseInstanceRoleResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDatabaseTableRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st DeleteDatabaseTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDatabaseTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDatabaseTableResponse struct {
}

func (st DeleteDatabaseTableResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDatabaseTableResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDatabaseTableResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDatabaseTableResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDatabaseTableResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDatabaseTableResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDatabaseTableResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteSyncedDatabaseTableRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st DeleteSyncedDatabaseTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSyncedDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSyncedDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSyncedDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteSyncedDatabaseTableResponse struct {
}

func (st DeleteSyncedDatabaseTableResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteSyncedDatabaseTableResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteSyncedDatabaseTableResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSyncedDatabaseTableResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSyncedDatabaseTableResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSyncedDatabaseTableResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteSyncedDatabaseTableResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeltaTableSyncInfo struct {
	// The timestamp when the above Delta version was committed in the source
	// Delta table. Note: This is the Delta commit time, not the time the data
	// was written to the synced table.
	// Wire name: 'delta_commit_timestamp'
	DeltaCommitTimestamp string `json:"delta_commit_timestamp,omitempty"`
	// The Delta Lake commit version that was last successfully synced.
	// Wire name: 'delta_commit_version'
	DeltaCommitVersion int64 `json:"delta_commit_version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeltaTableSyncInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := deltaTableSyncInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeltaTableSyncInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaTableSyncInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaTableSyncInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaTableSyncInfo) MarshalJSON() ([]byte, error) {
	pb, err := deltaTableSyncInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FindDatabaseInstanceByUidRequest struct {
	// UID of the cluster to get.
	Uid string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FindDatabaseInstanceByUidRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := findDatabaseInstanceByUidRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *FindDatabaseInstanceByUidRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &findDatabaseInstanceByUidRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := findDatabaseInstanceByUidRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FindDatabaseInstanceByUidRequest) MarshalJSON() ([]byte, error) {
	pb, err := findDatabaseInstanceByUidRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Generates a credential that can be used to access database instances
type GenerateDatabaseCredentialRequest struct {
	// The returned token will be scoped to the union of instance_names and
	// instances containing the specified UC tables, so instance_names is
	// allowed to be empty.
	// Wire name: 'claims'
	Claims []RequestedClaims `json:"claims,omitempty"`
	// Instances to which the token will be scoped.
	// Wire name: 'instance_names'
	InstanceNames []string `json:"instance_names,omitempty"`

	// Wire name: 'request_id'
	RequestId string `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GenerateDatabaseCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := generateDatabaseCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GenerateDatabaseCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &generateDatabaseCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := generateDatabaseCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenerateDatabaseCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := generateDatabaseCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetDatabaseCatalogRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st GetDatabaseCatalogRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDatabaseCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDatabaseCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDatabaseCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetDatabaseInstanceRequest struct {
	// Name of the cluster to get.
	Name string `json:"-" tf:"-"`
}

func (st GetDatabaseInstanceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDatabaseInstanceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDatabaseInstanceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetDatabaseInstanceRoleRequest struct {
	InstanceName string `json:"-" tf:"-"`

	Name string `json:"-" tf:"-"`
}

func (st GetDatabaseInstanceRoleRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDatabaseInstanceRoleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDatabaseInstanceRoleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetDatabaseTableRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st GetDatabaseTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDatabaseTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetSyncedDatabaseTableRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st GetSyncedDatabaseTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getSyncedDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getSyncedDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetSyncedDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := getSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListDatabaseInstanceRolesRequest struct {
	InstanceName string `json:"-" tf:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDatabaseInstanceRolesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listDatabaseInstanceRolesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListDatabaseInstanceRolesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDatabaseInstanceRolesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDatabaseInstanceRolesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDatabaseInstanceRolesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listDatabaseInstanceRolesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListDatabaseInstanceRolesResponse struct {
	// List of database instance roles.
	// Wire name: 'database_instance_roles'
	DatabaseInstanceRoles []DatabaseInstanceRole `json:"database_instance_roles,omitempty"`
	// Pagination token to request the next page of instances.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDatabaseInstanceRolesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listDatabaseInstanceRolesResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListDatabaseInstanceRolesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDatabaseInstanceRolesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDatabaseInstanceRolesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDatabaseInstanceRolesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listDatabaseInstanceRolesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListDatabaseInstancesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDatabaseInstancesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listDatabaseInstancesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListDatabaseInstancesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDatabaseInstancesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDatabaseInstancesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDatabaseInstancesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listDatabaseInstancesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListDatabaseInstancesResponse struct {
	// List of instances.
	// Wire name: 'database_instances'
	DatabaseInstances []DatabaseInstance `json:"database_instances,omitempty"`
	// Pagination token to request the next page of instances.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDatabaseInstancesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listDatabaseInstancesResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListDatabaseInstancesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDatabaseInstancesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDatabaseInstancesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDatabaseInstancesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listDatabaseInstancesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec struct {
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	// Wire name: 'storage_catalog'
	StorageCatalog string `json:"storage_catalog,omitempty"`
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	// Wire name: 'storage_schema'
	StorageSchema string `json:"storage_schema,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NewPipelineSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := newPipelineSpecToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *NewPipelineSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &newPipelineSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := newPipelineSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NewPipelineSpec) MarshalJSON() ([]byte, error) {
	pb, err := newPipelineSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type RequestedClaims struct {

	// Wire name: 'permission_set'
	PermissionSet RequestedClaimsPermissionSet `json:"permission_set,omitempty"`

	// Wire name: 'resources'
	Resources []RequestedResource `json:"resources,omitempty"`
}

func (st RequestedClaims) EncodeValues(key string, v *url.Values) error {
	pb, err := requestedClaimsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RequestedClaims) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &requestedClaimsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := requestedClaimsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RequestedClaims) MarshalJSON() ([]byte, error) {
	pb, err := requestedClaimsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'table_name'
	TableName string `json:"table_name,omitempty"`

	// Wire name: 'unspecified_resource_name'
	UnspecifiedResourceName string `json:"unspecified_resource_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RequestedResource) EncodeValues(key string, v *url.Values) error {
	pb, err := requestedResourceToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RequestedResource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &requestedResourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := requestedResourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RequestedResource) MarshalJSON() ([]byte, error) {
	pb, err := requestedResourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Next field marker: 12
type SyncedDatabaseTable struct {
	// Synced Table data synchronization status
	// Wire name: 'data_synchronization_status'
	DataSynchronizationStatus *SyncedTableStatus `json:"data_synchronization_status,omitempty"`
	// Name of the target database instance. This is required when creating
	// synced database tables in standard catalogs. This is optional when
	// creating synced database tables in registered catalogs. If this field is
	// specified when creating synced database tables in registered catalogs,
	// the database instance name MUST match that of the registered catalog (or
	// the request will be rejected).
	// Wire name: 'database_instance_name'
	DatabaseInstanceName string `json:"database_instance_name,omitempty"`
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
	// Wire name: 'logical_database_name'
	LogicalDatabaseName string `json:"logical_database_name,omitempty"`
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'spec'
	Spec *SyncedTableSpec `json:"spec,omitempty"`
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	// Wire name: 'unity_catalog_provisioning_state'
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedDatabaseTable) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedDatabaseTableToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedDatabaseTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedDatabaseTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedDatabaseTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedDatabaseTable) MarshalJSON() ([]byte, error) {
	pb, err := syncedDatabaseTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
type SyncedTableContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
	// The last source table Delta version that was successfully synced to the
	// synced table.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	// Wire name: 'timestamp'
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTableContinuousUpdateStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedTableContinuousUpdateStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedTableContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedTableContinuousUpdateStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedTableContinuousUpdateStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedTableContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	pb, err := syncedTableContinuousUpdateStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Detailed status of a synced table. Shown if the synced table is in the
// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
type SyncedTableFailedStatus struct {
	// The last source table Delta version that was successfully synced to the
	// synced table. The last source table Delta version that was synced to the
	// synced table. Only populated if the table is still synced and available
	// for serving.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. Only populated if the table is still
	// synced and available for serving.
	// Wire name: 'timestamp'
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTableFailedStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedTableFailedStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedTableFailedStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedTableFailedStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedTableFailedStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedTableFailedStatus) MarshalJSON() ([]byte, error) {
	pb, err := syncedTableFailedStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Progress information of the Synced Table data synchronization pipeline.
type SyncedTablePipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	// Wire name: 'estimated_completion_time_seconds'
	EstimatedCompletionTimeSeconds float64 `json:"estimated_completion_time_seconds,omitempty"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	// Wire name: 'latest_version_currently_processing'
	LatestVersionCurrentlyProcessing int64 `json:"latest_version_currently_processing,omitempty"`
	// The completion ratio of this update. This is a number between 0 and 1.
	// Wire name: 'sync_progress_completion'
	SyncProgressCompletion float64 `json:"sync_progress_completion,omitempty"`
	// The number of rows that have been synced in this update.
	// Wire name: 'synced_row_count'
	SyncedRowCount int64 `json:"synced_row_count,omitempty"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	// Wire name: 'total_row_count'
	TotalRowCount int64 `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTablePipelineProgress) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedTablePipelineProgressToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedTablePipelineProgress) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedTablePipelineProgressPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedTablePipelineProgressFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedTablePipelineProgress) MarshalJSON() ([]byte, error) {
	pb, err := syncedTablePipelineProgressToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SyncedTablePosition struct {

	// Wire name: 'delta_table_sync_info'
	DeltaTableSyncInfo *DeltaTableSyncInfo `json:"delta_table_sync_info,omitempty"`
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	// Wire name: 'sync_end_timestamp'
	SyncEndTimestamp string `json:"sync_end_timestamp,omitempty"`
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	// Wire name: 'sync_start_timestamp'
	SyncStartTimestamp string `json:"sync_start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTablePosition) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedTablePositionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedTablePosition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedTablePositionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedTablePositionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedTablePosition) MarshalJSON() ([]byte, error) {
	pb, err := syncedTablePositionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Detailed status of a synced table. Shown if the synced table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type SyncedTableProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

func (st SyncedTableProvisioningStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedTableProvisioningStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedTableProvisioningStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedTableProvisioningStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedTableProvisioningStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedTableProvisioningStatus) MarshalJSON() ([]byte, error) {
	pb, err := syncedTableProvisioningStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'create_database_objects_if_missing'
	CreateDatabaseObjectsIfMissing bool `json:"create_database_objects_if_missing,omitempty"`
	// User-specified ID of a pre-existing pipeline to bin pack. This field is
	// optional, and should be empty if new_pipeline_spec is set. This field
	// will only be set by the server in response messages if it is specified in
	// the request. The SyncedTableStatus message will always contain the
	// effective pipeline ID (either client provided or server generated),
	// however.
	// Wire name: 'existing_pipeline_id'
	ExistingPipelineId string `json:"existing_pipeline_id,omitempty"`
	// Spec of new pipeline. Should be empty if pipeline_id /
	// existing_pipeline_id is set
	// Wire name: 'new_pipeline_spec'
	NewPipelineSpec *NewPipelineSpec `json:"new_pipeline_spec,omitempty"`
	// Primary Key columns to be used for data insert/update in the destination.
	// Wire name: 'primary_key_columns'
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Scheduling policy of the underlying pipeline.
	// Wire name: 'scheduling_policy'
	SchedulingPolicy SyncedTableSchedulingPolicy `json:"scheduling_policy,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	// Wire name: 'source_table_full_name'
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	// Wire name: 'timeseries_key'
	TimeseriesKey string `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTableSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedTableSpecToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedTableSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedTableSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedTableSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedTableSpec) MarshalJSON() ([]byte, error) {
	pb, err := syncedTableSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'continuous_update_status'
	ContinuousUpdateStatus *SyncedTableContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	// The state of the synced table.
	// Wire name: 'detailed_state'
	DetailedState SyncedTableState `json:"detailed_state,omitempty"`

	// Wire name: 'failed_status'
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
	// Wire name: 'last_sync'
	LastSync *SyncedTablePosition `json:"last_sync,omitempty"`
	// A text description of the current state of the synced table.
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// ID of the associated pipeline. The pipeline ID may have been provided by
	// the client (in the case of bin packing), or generated by the server (when
	// creating a new pipeline).
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`

	// Wire name: 'provisioning_status'
	ProvisioningStatus *SyncedTableProvisioningStatus `json:"provisioning_status,omitempty"`

	// Wire name: 'triggered_update_status'
	TriggeredUpdateStatus *SyncedTableTriggeredUpdateStatus `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTableStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedTableStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedTableStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedTableStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedTableStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedTableStatus) MarshalJSON() ([]byte, error) {
	pb, err := syncedTableStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
type SyncedTableTriggeredUpdateStatus struct {
	// The last source table Delta version that was successfully synced to the
	// synced table.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	// Wire name: 'timestamp'
	Timestamp string `json:"timestamp,omitempty"`
	// Progress of the active data synchronization pipeline.
	// Wire name: 'triggered_update_progress'
	TriggeredUpdateProgress *SyncedTablePipelineProgress `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTableTriggeredUpdateStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := syncedTableTriggeredUpdateStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SyncedTableTriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncedTableTriggeredUpdateStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncedTableTriggeredUpdateStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncedTableTriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	pb, err := syncedTableTriggeredUpdateStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateDatabaseInstanceRequest struct {

	// Wire name: 'database_instance'
	DatabaseInstance DatabaseInstance `json:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	Name string `json:"-" tf:"-"`
	// The list of fields to update.
	UpdateMask string `json:"-" tf:"-"`
}

func (st UpdateDatabaseInstanceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateDatabaseInstanceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateDatabaseInstanceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
