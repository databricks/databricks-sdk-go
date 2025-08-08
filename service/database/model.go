// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/database/databasepb"
)

type CreateDatabaseCatalogRequest struct {

	// Wire name: 'catalog'
	Catalog DatabaseCatalog ``
}

func CreateDatabaseCatalogRequestToPb(st *CreateDatabaseCatalogRequest) (*databasepb.CreateDatabaseCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.CreateDatabaseCatalogRequestPb{}
	catalogPb, err := DatabaseCatalogToPb(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	return pb, nil
}

func CreateDatabaseCatalogRequestFromPb(pb *databasepb.CreateDatabaseCatalogRequestPb) (*CreateDatabaseCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDatabaseCatalogRequest{}
	catalogField, err := DatabaseCatalogFromPb(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}

	return st, nil
}

type CreateDatabaseInstanceRequest struct {
	// Instance to create.
	// Wire name: 'database_instance'
	DatabaseInstance DatabaseInstance ``
}

func CreateDatabaseInstanceRequestToPb(st *CreateDatabaseInstanceRequest) (*databasepb.CreateDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.CreateDatabaseInstanceRequestPb{}
	databaseInstancePb, err := DatabaseInstanceToPb(&st.DatabaseInstance)
	if err != nil {
		return nil, err
	}
	if databaseInstancePb != nil {
		pb.DatabaseInstance = *databaseInstancePb
	}

	return pb, nil
}

func CreateDatabaseInstanceRequestFromPb(pb *databasepb.CreateDatabaseInstanceRequestPb) (*CreateDatabaseInstanceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDatabaseInstanceRequest{}
	databaseInstanceField, err := DatabaseInstanceFromPb(&pb.DatabaseInstance)
	if err != nil {
		return nil, err
	}
	if databaseInstanceField != nil {
		st.DatabaseInstance = *databaseInstanceField
	}

	return st, nil
}

type CreateDatabaseInstanceRoleRequest struct {

	// Wire name: 'database_instance_role'
	DatabaseInstanceRole DatabaseInstanceRole ``

	// Wire name: 'instance_name'
	InstanceName string `tf:"-"`
}

func CreateDatabaseInstanceRoleRequestToPb(st *CreateDatabaseInstanceRoleRequest) (*databasepb.CreateDatabaseInstanceRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.CreateDatabaseInstanceRoleRequestPb{}
	databaseInstanceRolePb, err := DatabaseInstanceRoleToPb(&st.DatabaseInstanceRole)
	if err != nil {
		return nil, err
	}
	if databaseInstanceRolePb != nil {
		pb.DatabaseInstanceRole = *databaseInstanceRolePb
	}
	pb.InstanceName = st.InstanceName

	return pb, nil
}

func CreateDatabaseInstanceRoleRequestFromPb(pb *databasepb.CreateDatabaseInstanceRoleRequestPb) (*CreateDatabaseInstanceRoleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDatabaseInstanceRoleRequest{}
	databaseInstanceRoleField, err := DatabaseInstanceRoleFromPb(&pb.DatabaseInstanceRole)
	if err != nil {
		return nil, err
	}
	if databaseInstanceRoleField != nil {
		st.DatabaseInstanceRole = *databaseInstanceRoleField
	}
	st.InstanceName = pb.InstanceName

	return st, nil
}

type CreateDatabaseTableRequest struct {

	// Wire name: 'table'
	Table DatabaseTable ``
}

func CreateDatabaseTableRequestToPb(st *CreateDatabaseTableRequest) (*databasepb.CreateDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.CreateDatabaseTableRequestPb{}
	tablePb, err := DatabaseTableToPb(&st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = *tablePb
	}

	return pb, nil
}

func CreateDatabaseTableRequestFromPb(pb *databasepb.CreateDatabaseTableRequestPb) (*CreateDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDatabaseTableRequest{}
	tableField, err := DatabaseTableFromPb(&pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = *tableField
	}

	return st, nil
}

type CreateSyncedDatabaseTableRequest struct {

	// Wire name: 'synced_table'
	SyncedTable SyncedDatabaseTable ``
}

func CreateSyncedDatabaseTableRequestToPb(st *CreateSyncedDatabaseTableRequest) (*databasepb.CreateSyncedDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.CreateSyncedDatabaseTableRequestPb{}
	syncedTablePb, err := SyncedDatabaseTableToPb(&st.SyncedTable)
	if err != nil {
		return nil, err
	}
	if syncedTablePb != nil {
		pb.SyncedTable = *syncedTablePb
	}

	return pb, nil
}

func CreateSyncedDatabaseTableRequestFromPb(pb *databasepb.CreateSyncedDatabaseTableRequestPb) (*CreateSyncedDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateSyncedDatabaseTableRequest{}
	syncedTableField, err := SyncedDatabaseTableFromPb(&pb.SyncedTable)
	if err != nil {
		return nil, err
	}
	if syncedTableField != nil {
		st.SyncedTable = *syncedTableField
	}

	return st, nil
}

type DatabaseCatalog struct {

	// Wire name: 'create_database_if_not_exists'
	CreateDatabaseIfNotExists bool ``
	// The name of the DatabaseInstance housing the database.
	// Wire name: 'database_instance_name'
	DatabaseInstanceName string ``
	// The name of the database (in a instance) associated with the catalog.
	// Wire name: 'database_name'
	DatabaseName string ``
	// The name of the catalog in UC.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'uid'
	Uid             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabaseCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabaseCatalogToPb(st *DatabaseCatalog) (*databasepb.DatabaseCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseCatalogPb{}
	pb.CreateDatabaseIfNotExists = st.CreateDatabaseIfNotExists
	pb.DatabaseInstanceName = st.DatabaseInstanceName
	pb.DatabaseName = st.DatabaseName
	pb.Name = st.Name
	pb.Uid = st.Uid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabaseCatalogFromPb(pb *databasepb.DatabaseCatalogPb) (*DatabaseCatalog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseCatalog{}
	st.CreateDatabaseIfNotExists = pb.CreateDatabaseIfNotExists
	st.DatabaseInstanceName = pb.DatabaseInstanceName
	st.DatabaseName = pb.DatabaseName
	st.Name = pb.Name
	st.Uid = pb.Uid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DatabaseCredential struct {

	// Wire name: 'expiration_time'
	ExpirationTime *time.Time ``

	// Wire name: 'token'
	Token           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabaseCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabaseCredentialToPb(st *DatabaseCredential) (*databasepb.DatabaseCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseCredentialPb{}
	expirationTimePb, err := timestampToPb(st.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}
	pb.Token = st.Token

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabaseCredentialFromPb(pb *databasepb.DatabaseCredentialPb) (*DatabaseCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseCredential{}
	expirationTimeField, err := timestampFromPb(&pb.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimeField != nil {
		st.ExpirationTime = expirationTimeField
	}
	st.Token = pb.Token

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	// Wire name: 'capacity'
	Capacity string ``
	// The refs of the child instances. This is only available if the instance
	// is parent instance.
	// Wire name: 'child_instance_refs'
	ChildInstanceRefs []DatabaseInstanceRef ``
	// The timestamp when the instance was created.
	// Wire name: 'creation_time'
	CreationTime *time.Time ``
	// The email of the creator of the instance.
	// Wire name: 'creator'
	Creator string ``
	// xref AIP-129. `enable_readable_secondaries` is owned by the client, while
	// `effective_enable_readable_secondaries` is owned by the server.
	// `enable_readable_secondaries` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_enable_readable_secondaries` on the other hand will always bet
	// set in all response messages (Create/Update/Get/List).
	// Wire name: 'effective_enable_readable_secondaries'
	EffectiveEnableReadableSecondaries bool ``
	// xref AIP-129. `node_count` is owned by the client, while
	// `effective_node_count` is owned by the server. `node_count` will only be
	// set in Create/Update response messages if and only if the user provides
	// the field via the request. `effective_node_count` on the other hand will
	// always bet set in all response messages (Create/Update/Get/List).
	// Wire name: 'effective_node_count'
	EffectiveNodeCount int ``
	// xref AIP-129. `retention_window_in_days` is owned by the client, while
	// `effective_retention_window_in_days` is owned by the server.
	// `retention_window_in_days` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_retention_window_in_days` on the other hand will always bet
	// set in all response messages (Create/Update/Get/List).
	// Wire name: 'effective_retention_window_in_days'
	EffectiveRetentionWindowInDays int ``
	// xref AIP-129. `stopped` is owned by the client, while `effective_stopped`
	// is owned by the server. `stopped` will only be set in Create/Update
	// response messages if and only if the user provides the field via the
	// request. `effective_stopped` on the other hand will always bet set in all
	// response messages (Create/Update/Get/List).
	// Wire name: 'effective_stopped'
	EffectiveStopped bool ``
	// Whether to enable secondaries to serve read-only traffic. Defaults to
	// false.
	// Wire name: 'enable_readable_secondaries'
	EnableReadableSecondaries bool ``
	// The name of the instance. This is the unique identifier for the instance.
	// Wire name: 'name'
	Name string ``
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries.
	// Wire name: 'node_count'
	NodeCount int ``
	// The ref of the parent instance. This is only available if the instance is
	// child instance. Input: For specifying the parent instance to create a
	// child instance. Optional. Output: Only populated if provided as input to
	// create a child instance.
	// Wire name: 'parent_instance_ref'
	ParentInstanceRef *DatabaseInstanceRef ``
	// The version of Postgres running on the instance.
	// Wire name: 'pg_version'
	PgVersion string ``
	// The DNS endpoint to connect to the instance for read only access. This is
	// only available if enable_readable_secondaries is true.
	// Wire name: 'read_only_dns'
	ReadOnlyDns string ``
	// The DNS endpoint to connect to the instance for read+write access.
	// Wire name: 'read_write_dns'
	ReadWriteDns string ``
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained. The default value is 7 days.
	// Valid values are 2 to 35 days.
	// Wire name: 'retention_window_in_days'
	RetentionWindowInDays int ``
	// The current state of the instance.
	// Wire name: 'state'
	State DatabaseInstanceState ``
	// Whether the instance is stopped.
	// Wire name: 'stopped'
	Stopped bool ``
	// An immutable UUID identifier for the instance.
	// Wire name: 'uid'
	Uid             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabaseInstance) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstance) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabaseInstanceToPb(st *DatabaseInstance) (*databasepb.DatabaseInstancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseInstancePb{}
	pb.Capacity = st.Capacity

	var childInstanceRefsPb []databasepb.DatabaseInstanceRefPb
	for _, item := range st.ChildInstanceRefs {
		itemPb, err := DatabaseInstanceRefToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			childInstanceRefsPb = append(childInstanceRefsPb, *itemPb)
		}
	}
	pb.ChildInstanceRefs = childInstanceRefsPb
	creationTimePb, err := timestampToPb(st.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}
	pb.Creator = st.Creator
	pb.EffectiveEnableReadableSecondaries = st.EffectiveEnableReadableSecondaries
	pb.EffectiveNodeCount = st.EffectiveNodeCount
	pb.EffectiveRetentionWindowInDays = st.EffectiveRetentionWindowInDays
	pb.EffectiveStopped = st.EffectiveStopped
	pb.EnableReadableSecondaries = st.EnableReadableSecondaries
	pb.Name = st.Name
	pb.NodeCount = st.NodeCount
	parentInstanceRefPb, err := DatabaseInstanceRefToPb(st.ParentInstanceRef)
	if err != nil {
		return nil, err
	}
	if parentInstanceRefPb != nil {
		pb.ParentInstanceRef = parentInstanceRefPb
	}
	pb.PgVersion = st.PgVersion
	pb.ReadOnlyDns = st.ReadOnlyDns
	pb.ReadWriteDns = st.ReadWriteDns
	pb.RetentionWindowInDays = st.RetentionWindowInDays
	statePb, err := DatabaseInstanceStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	pb.Stopped = st.Stopped
	pb.Uid = st.Uid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabaseInstanceFromPb(pb *databasepb.DatabaseInstancePb) (*DatabaseInstance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseInstance{}
	st.Capacity = pb.Capacity

	var childInstanceRefsField []DatabaseInstanceRef
	for _, itemPb := range pb.ChildInstanceRefs {
		item, err := DatabaseInstanceRefFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			childInstanceRefsField = append(childInstanceRefsField, *item)
		}
	}
	st.ChildInstanceRefs = childInstanceRefsField
	creationTimeField, err := timestampFromPb(&pb.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimeField != nil {
		st.CreationTime = creationTimeField
	}
	st.Creator = pb.Creator
	st.EffectiveEnableReadableSecondaries = pb.EffectiveEnableReadableSecondaries
	st.EffectiveNodeCount = pb.EffectiveNodeCount
	st.EffectiveRetentionWindowInDays = pb.EffectiveRetentionWindowInDays
	st.EffectiveStopped = pb.EffectiveStopped
	st.EnableReadableSecondaries = pb.EnableReadableSecondaries
	st.Name = pb.Name
	st.NodeCount = pb.NodeCount
	parentInstanceRefField, err := DatabaseInstanceRefFromPb(pb.ParentInstanceRef)
	if err != nil {
		return nil, err
	}
	if parentInstanceRefField != nil {
		st.ParentInstanceRef = parentInstanceRefField
	}
	st.PgVersion = pb.PgVersion
	st.ReadOnlyDns = pb.ReadOnlyDns
	st.ReadWriteDns = pb.ReadWriteDns
	st.RetentionWindowInDays = pb.RetentionWindowInDays
	stateField, err := DatabaseInstanceStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	st.Stopped = pb.Stopped
	st.Uid = pb.Uid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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
	BranchTime *time.Time ``
	// xref AIP-129. `lsn` is owned by the client, while `effective_lsn` is
	// owned by the server. `lsn` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_lsn` on the other hand will always bet set in all response
	// messages (Create/Update/Get/List). For a parent ref instance, this is the
	// LSN on the parent instance from which the instance was created. For a
	// child ref instance, this is the LSN on the instance from which the child
	// instance was created.
	// Wire name: 'effective_lsn'
	EffectiveLsn string ``
	// User-specified WAL LSN of the ref database instance.
	//
	// Input: For specifying the WAL LSN to create a child instance. Optional.
	// Output: Only populated if provided as input to create a child instance.
	// Wire name: 'lsn'
	Lsn string ``
	// Name of the ref database instance.
	// Wire name: 'name'
	Name string ``
	// Id of the ref database instance.
	// Wire name: 'uid'
	Uid             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabaseInstanceRef) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRef) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabaseInstanceRefToPb(st *DatabaseInstanceRef) (*databasepb.DatabaseInstanceRefPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseInstanceRefPb{}
	branchTimePb, err := timestampToPb(st.BranchTime)
	if err != nil {
		return nil, err
	}
	if branchTimePb != nil {
		pb.BranchTime = *branchTimePb
	}
	pb.EffectiveLsn = st.EffectiveLsn
	pb.Lsn = st.Lsn
	pb.Name = st.Name
	pb.Uid = st.Uid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabaseInstanceRefFromPb(pb *databasepb.DatabaseInstanceRefPb) (*DatabaseInstanceRef, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseInstanceRef{}
	branchTimeField, err := timestampFromPb(&pb.BranchTime)
	if err != nil {
		return nil, err
	}
	if branchTimeField != nil {
		st.BranchTime = branchTimeField
	}
	st.EffectiveLsn = pb.EffectiveLsn
	st.Lsn = pb.Lsn
	st.Name = pb.Name
	st.Uid = pb.Uid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A DatabaseInstanceRole represents a Postgres role in a database instance.
type DatabaseInstanceRole struct {
	// API-exposed Postgres role attributes
	// Wire name: 'attributes'
	Attributes *DatabaseInstanceRoleAttributes ``
	// The type of the role.
	// Wire name: 'identity_type'
	IdentityType DatabaseInstanceRoleIdentityType ``
	// An enum value for a standard role that this role is a member of.
	// Wire name: 'membership_role'
	MembershipRole DatabaseInstanceRoleMembershipRole ``
	// The name of the role. This is the unique identifier for the role in an
	// instance.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabaseInstanceRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabaseInstanceRoleToPb(st *DatabaseInstanceRole) (*databasepb.DatabaseInstanceRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseInstanceRolePb{}
	attributesPb, err := DatabaseInstanceRoleAttributesToPb(st.Attributes)
	if err != nil {
		return nil, err
	}
	if attributesPb != nil {
		pb.Attributes = attributesPb
	}
	identityTypePb, err := DatabaseInstanceRoleIdentityTypeToPb(&st.IdentityType)
	if err != nil {
		return nil, err
	}
	if identityTypePb != nil {
		pb.IdentityType = *identityTypePb
	}
	membershipRolePb, err := DatabaseInstanceRoleMembershipRoleToPb(&st.MembershipRole)
	if err != nil {
		return nil, err
	}
	if membershipRolePb != nil {
		pb.MembershipRole = *membershipRolePb
	}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabaseInstanceRoleFromPb(pb *databasepb.DatabaseInstanceRolePb) (*DatabaseInstanceRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseInstanceRole{}
	attributesField, err := DatabaseInstanceRoleAttributesFromPb(pb.Attributes)
	if err != nil {
		return nil, err
	}
	if attributesField != nil {
		st.Attributes = attributesField
	}
	identityTypeField, err := DatabaseInstanceRoleIdentityTypeFromPb(&pb.IdentityType)
	if err != nil {
		return nil, err
	}
	if identityTypeField != nil {
		st.IdentityType = *identityTypeField
	}
	membershipRoleField, err := DatabaseInstanceRoleMembershipRoleFromPb(&pb.MembershipRole)
	if err != nil {
		return nil, err
	}
	if membershipRoleField != nil {
		st.MembershipRole = *membershipRoleField
	}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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
	Bypassrls bool ``

	// Wire name: 'createdb'
	Createdb bool ``

	// Wire name: 'createrole'
	Createrole      bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabaseInstanceRoleAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRoleAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabaseInstanceRoleAttributesToPb(st *DatabaseInstanceRoleAttributes) (*databasepb.DatabaseInstanceRoleAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseInstanceRoleAttributesPb{}
	pb.Bypassrls = st.Bypassrls
	pb.Createdb = st.Createdb
	pb.Createrole = st.Createrole

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabaseInstanceRoleAttributesFromPb(pb *databasepb.DatabaseInstanceRoleAttributesPb) (*DatabaseInstanceRoleAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseInstanceRoleAttributes{}
	st.Bypassrls = pb.Bypassrls
	st.Createdb = pb.Createdb
	st.Createrole = pb.Createrole

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func DatabaseInstanceRoleIdentityTypeToPb(st *DatabaseInstanceRoleIdentityType) (*databasepb.DatabaseInstanceRoleIdentityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := databasepb.DatabaseInstanceRoleIdentityTypePb(*st)
	return &pb, nil
}

func DatabaseInstanceRoleIdentityTypeFromPb(pb *databasepb.DatabaseInstanceRoleIdentityTypePb) (*DatabaseInstanceRoleIdentityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DatabaseInstanceRoleIdentityType(*pb)
	return &st, nil
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

func DatabaseInstanceRoleMembershipRoleToPb(st *DatabaseInstanceRoleMembershipRole) (*databasepb.DatabaseInstanceRoleMembershipRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := databasepb.DatabaseInstanceRoleMembershipRolePb(*st)
	return &pb, nil
}

func DatabaseInstanceRoleMembershipRoleFromPb(pb *databasepb.DatabaseInstanceRoleMembershipRolePb) (*DatabaseInstanceRoleMembershipRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := DatabaseInstanceRoleMembershipRole(*pb)
	return &st, nil
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

func DatabaseInstanceStateToPb(st *DatabaseInstanceState) (*databasepb.DatabaseInstanceStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := databasepb.DatabaseInstanceStatePb(*st)
	return &pb, nil
}

func DatabaseInstanceStateFromPb(pb *databasepb.DatabaseInstanceStatePb) (*DatabaseInstanceState, error) {
	if pb == nil {
		return nil, nil
	}
	st := DatabaseInstanceState(*pb)
	return &st, nil
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
	DatabaseInstanceName string ``
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
	LogicalDatabaseName string ``
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabaseTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabaseTableToPb(st *DatabaseTable) (*databasepb.DatabaseTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseTablePb{}
	pb.DatabaseInstanceName = st.DatabaseInstanceName
	pb.LogicalDatabaseName = st.LogicalDatabaseName
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabaseTableFromPb(pb *databasepb.DatabaseTablePb) (*DatabaseTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseTable{}
	st.DatabaseInstanceName = pb.DatabaseInstanceName
	st.LogicalDatabaseName = pb.LogicalDatabaseName
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteDatabaseCatalogRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteDatabaseCatalogRequestToPb(st *DeleteDatabaseCatalogRequest) (*databasepb.DeleteDatabaseCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DeleteDatabaseCatalogRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteDatabaseCatalogRequestFromPb(pb *databasepb.DeleteDatabaseCatalogRequestPb) (*DeleteDatabaseCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseCatalogRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteDatabaseInstanceRequest struct {
	// By default, a instance cannot be deleted if it has descendant instances
	// created via PITR. If this flag is specified as true, all descendent
	// instances will be deleted as well.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Name of the instance to delete.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Note purge=false is in development. If false, the database instance is
	// soft deleted (implementation pending). Soft deleted instances behave as
	// if they are deleted, and cannot be used for CRUD operations nor connected
	// to. However they can be undeleted by calling the undelete API for a
	// limited time (implementation pending). If true, the database instance is
	// hard deleted and cannot be undeleted. For the time being, setting this
	// value to true is required to delete an instance (soft delete is not yet
	// supported).
	// Wire name: 'purge'
	Purge           bool     `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteDatabaseInstanceRequestToPb(st *DeleteDatabaseInstanceRequest) (*databasepb.DeleteDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DeleteDatabaseInstanceRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name
	pb.Purge = st.Purge

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteDatabaseInstanceRequestFromPb(pb *databasepb.DeleteDatabaseInstanceRequestPb) (*DeleteDatabaseInstanceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseInstanceRequest{}
	st.Force = pb.Force
	st.Name = pb.Name
	st.Purge = pb.Purge

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteDatabaseInstanceRoleRequest struct {
	// This is the AIP standard name for the equivalent of Postgres' `IF EXISTS`
	// option
	// Wire name: 'allow_missing'
	AllowMissing bool `tf:"-"`

	// Wire name: 'instance_name'
	InstanceName string `tf:"-"`

	// Wire name: 'name'
	Name string `tf:"-"`

	// Wire name: 'reassign_owned_to'
	ReassignOwnedTo string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteDatabaseInstanceRoleRequestToPb(st *DeleteDatabaseInstanceRoleRequest) (*databasepb.DeleteDatabaseInstanceRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DeleteDatabaseInstanceRoleRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.InstanceName = st.InstanceName
	pb.Name = st.Name
	pb.ReassignOwnedTo = st.ReassignOwnedTo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteDatabaseInstanceRoleRequestFromPb(pb *databasepb.DeleteDatabaseInstanceRoleRequestPb) (*DeleteDatabaseInstanceRoleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseInstanceRoleRequest{}
	st.AllowMissing = pb.AllowMissing
	st.InstanceName = pb.InstanceName
	st.Name = pb.Name
	st.ReassignOwnedTo = pb.ReassignOwnedTo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteDatabaseTableRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteDatabaseTableRequestToPb(st *DeleteDatabaseTableRequest) (*databasepb.DeleteDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DeleteDatabaseTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteDatabaseTableRequestFromPb(pb *databasepb.DeleteDatabaseTableRequestPb) (*DeleteDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseTableRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteSyncedDatabaseTableRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteSyncedDatabaseTableRequestToPb(st *DeleteSyncedDatabaseTableRequest) (*databasepb.DeleteSyncedDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DeleteSyncedDatabaseTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteSyncedDatabaseTableRequestFromPb(pb *databasepb.DeleteSyncedDatabaseTableRequestPb) (*DeleteSyncedDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSyncedDatabaseTableRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeltaTableSyncInfo struct {
	// The timestamp when the above Delta version was committed in the source
	// Delta table. Note: This is the Delta commit time, not the time the data
	// was written to the synced table.
	// Wire name: 'delta_commit_timestamp'
	DeltaCommitTimestamp *time.Time ``
	// The Delta Lake commit version that was last successfully synced.
	// Wire name: 'delta_commit_version'
	DeltaCommitVersion int64    ``
	ForceSendFields    []string `tf:"-"`
}

func (s *DeltaTableSyncInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaTableSyncInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeltaTableSyncInfoToPb(st *DeltaTableSyncInfo) (*databasepb.DeltaTableSyncInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DeltaTableSyncInfoPb{}
	deltaCommitTimestampPb, err := timestampToPb(st.DeltaCommitTimestamp)
	if err != nil {
		return nil, err
	}
	if deltaCommitTimestampPb != nil {
		pb.DeltaCommitTimestamp = *deltaCommitTimestampPb
	}
	pb.DeltaCommitVersion = st.DeltaCommitVersion

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeltaTableSyncInfoFromPb(pb *databasepb.DeltaTableSyncInfoPb) (*DeltaTableSyncInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaTableSyncInfo{}
	deltaCommitTimestampField, err := timestampFromPb(&pb.DeltaCommitTimestamp)
	if err != nil {
		return nil, err
	}
	if deltaCommitTimestampField != nil {
		st.DeltaCommitTimestamp = deltaCommitTimestampField
	}
	st.DeltaCommitVersion = pb.DeltaCommitVersion

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type FindDatabaseInstanceByUidRequest struct {
	// UID of the cluster to get.
	// Wire name: 'uid'
	Uid             string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *FindDatabaseInstanceByUidRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FindDatabaseInstanceByUidRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FindDatabaseInstanceByUidRequestToPb(st *FindDatabaseInstanceByUidRequest) (*databasepb.FindDatabaseInstanceByUidRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.FindDatabaseInstanceByUidRequestPb{}
	pb.Uid = st.Uid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FindDatabaseInstanceByUidRequestFromPb(pb *databasepb.FindDatabaseInstanceByUidRequestPb) (*FindDatabaseInstanceByUidRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FindDatabaseInstanceByUidRequest{}
	st.Uid = pb.Uid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Generates a credential that can be used to access database instances
type GenerateDatabaseCredentialRequest struct {
	// The returned token will be scoped to the union of instance_names and
	// instances containing the specified UC tables, so instance_names is
	// allowed to be empty.
	// Wire name: 'claims'
	Claims []RequestedClaims ``
	// Instances to which the token will be scoped.
	// Wire name: 'instance_names'
	InstanceNames []string ``

	// Wire name: 'request_id'
	RequestId       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenerateDatabaseCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateDatabaseCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenerateDatabaseCredentialRequestToPb(st *GenerateDatabaseCredentialRequest) (*databasepb.GenerateDatabaseCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.GenerateDatabaseCredentialRequestPb{}

	var claimsPb []databasepb.RequestedClaimsPb
	for _, item := range st.Claims {
		itemPb, err := RequestedClaimsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			claimsPb = append(claimsPb, *itemPb)
		}
	}
	pb.Claims = claimsPb
	pb.InstanceNames = st.InstanceNames
	pb.RequestId = st.RequestId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenerateDatabaseCredentialRequestFromPb(pb *databasepb.GenerateDatabaseCredentialRequestPb) (*GenerateDatabaseCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateDatabaseCredentialRequest{}

	var claimsField []RequestedClaims
	for _, itemPb := range pb.Claims {
		item, err := RequestedClaimsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			claimsField = append(claimsField, *item)
		}
	}
	st.Claims = claimsField
	st.InstanceNames = pb.InstanceNames
	st.RequestId = pb.RequestId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetDatabaseCatalogRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetDatabaseCatalogRequestToPb(st *GetDatabaseCatalogRequest) (*databasepb.GetDatabaseCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.GetDatabaseCatalogRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetDatabaseCatalogRequestFromPb(pb *databasepb.GetDatabaseCatalogRequestPb) (*GetDatabaseCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDatabaseCatalogRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetDatabaseInstanceRequest struct {
	// Name of the cluster to get.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetDatabaseInstanceRequestToPb(st *GetDatabaseInstanceRequest) (*databasepb.GetDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.GetDatabaseInstanceRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetDatabaseInstanceRequestFromPb(pb *databasepb.GetDatabaseInstanceRequestPb) (*GetDatabaseInstanceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDatabaseInstanceRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetDatabaseInstanceRoleRequest struct {

	// Wire name: 'instance_name'
	InstanceName string `tf:"-"`

	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetDatabaseInstanceRoleRequestToPb(st *GetDatabaseInstanceRoleRequest) (*databasepb.GetDatabaseInstanceRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.GetDatabaseInstanceRoleRequestPb{}
	pb.InstanceName = st.InstanceName
	pb.Name = st.Name

	return pb, nil
}

func GetDatabaseInstanceRoleRequestFromPb(pb *databasepb.GetDatabaseInstanceRoleRequestPb) (*GetDatabaseInstanceRoleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDatabaseInstanceRoleRequest{}
	st.InstanceName = pb.InstanceName
	st.Name = pb.Name

	return st, nil
}

type GetDatabaseTableRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetDatabaseTableRequestToPb(st *GetDatabaseTableRequest) (*databasepb.GetDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.GetDatabaseTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetDatabaseTableRequestFromPb(pb *databasepb.GetDatabaseTableRequestPb) (*GetDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDatabaseTableRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetSyncedDatabaseTableRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetSyncedDatabaseTableRequestToPb(st *GetSyncedDatabaseTableRequest) (*databasepb.GetSyncedDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.GetSyncedDatabaseTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetSyncedDatabaseTableRequestFromPb(pb *databasepb.GetSyncedDatabaseTableRequestPb) (*GetSyncedDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSyncedDatabaseTableRequest{}
	st.Name = pb.Name

	return st, nil
}

type ListDatabaseInstanceRolesRequest struct {

	// Wire name: 'instance_name'
	InstanceName string `tf:"-"`
	// Upper bound for items returned.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListDatabaseInstanceRolesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstanceRolesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListDatabaseInstanceRolesRequestToPb(st *ListDatabaseInstanceRolesRequest) (*databasepb.ListDatabaseInstanceRolesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.ListDatabaseInstanceRolesRequestPb{}
	pb.InstanceName = st.InstanceName
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListDatabaseInstanceRolesRequestFromPb(pb *databasepb.ListDatabaseInstanceRolesRequestPb) (*ListDatabaseInstanceRolesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDatabaseInstanceRolesRequest{}
	st.InstanceName = pb.InstanceName
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListDatabaseInstanceRolesResponse struct {
	// List of database instance roles.
	// Wire name: 'database_instance_roles'
	DatabaseInstanceRoles []DatabaseInstanceRole ``
	// Pagination token to request the next page of instances.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListDatabaseInstanceRolesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstanceRolesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListDatabaseInstanceRolesResponseToPb(st *ListDatabaseInstanceRolesResponse) (*databasepb.ListDatabaseInstanceRolesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.ListDatabaseInstanceRolesResponsePb{}

	var databaseInstanceRolesPb []databasepb.DatabaseInstanceRolePb
	for _, item := range st.DatabaseInstanceRoles {
		itemPb, err := DatabaseInstanceRoleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			databaseInstanceRolesPb = append(databaseInstanceRolesPb, *itemPb)
		}
	}
	pb.DatabaseInstanceRoles = databaseInstanceRolesPb
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListDatabaseInstanceRolesResponseFromPb(pb *databasepb.ListDatabaseInstanceRolesResponsePb) (*ListDatabaseInstanceRolesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDatabaseInstanceRolesResponse{}

	var databaseInstanceRolesField []DatabaseInstanceRole
	for _, itemPb := range pb.DatabaseInstanceRoles {
		item, err := DatabaseInstanceRoleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			databaseInstanceRolesField = append(databaseInstanceRolesField, *item)
		}
	}
	st.DatabaseInstanceRoles = databaseInstanceRolesField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListDatabaseInstancesRequest struct {
	// Upper bound for items returned.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListDatabaseInstancesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstancesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListDatabaseInstancesRequestToPb(st *ListDatabaseInstancesRequest) (*databasepb.ListDatabaseInstancesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.ListDatabaseInstancesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListDatabaseInstancesRequestFromPb(pb *databasepb.ListDatabaseInstancesRequestPb) (*ListDatabaseInstancesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDatabaseInstancesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListDatabaseInstancesResponse struct {
	// List of instances.
	// Wire name: 'database_instances'
	DatabaseInstances []DatabaseInstance ``
	// Pagination token to request the next page of instances.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListDatabaseInstancesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstancesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListDatabaseInstancesResponseToPb(st *ListDatabaseInstancesResponse) (*databasepb.ListDatabaseInstancesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.ListDatabaseInstancesResponsePb{}

	var databaseInstancesPb []databasepb.DatabaseInstancePb
	for _, item := range st.DatabaseInstances {
		itemPb, err := DatabaseInstanceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			databaseInstancesPb = append(databaseInstancesPb, *itemPb)
		}
	}
	pb.DatabaseInstances = databaseInstancesPb
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListDatabaseInstancesResponseFromPb(pb *databasepb.ListDatabaseInstancesResponsePb) (*ListDatabaseInstancesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDatabaseInstancesResponse{}

	var databaseInstancesField []DatabaseInstance
	for _, itemPb := range pb.DatabaseInstances {
		item, err := DatabaseInstanceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			databaseInstancesField = append(databaseInstancesField, *item)
		}
	}
	st.DatabaseInstances = databaseInstancesField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec struct {
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	// Wire name: 'storage_catalog'
	StorageCatalog string ``
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	// Wire name: 'storage_schema'
	StorageSchema   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *NewPipelineSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NewPipelineSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func NewPipelineSpecToPb(st *NewPipelineSpec) (*databasepb.NewPipelineSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.NewPipelineSpecPb{}
	pb.StorageCatalog = st.StorageCatalog
	pb.StorageSchema = st.StorageSchema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func NewPipelineSpecFromPb(pb *databasepb.NewPipelineSpecPb) (*NewPipelineSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NewPipelineSpec{}
	st.StorageCatalog = pb.StorageCatalog
	st.StorageSchema = pb.StorageSchema

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ProvisioningInfoStateToPb(st *ProvisioningInfoState) (*databasepb.ProvisioningInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := databasepb.ProvisioningInfoStatePb(*st)
	return &pb, nil
}

func ProvisioningInfoStateFromPb(pb *databasepb.ProvisioningInfoStatePb) (*ProvisioningInfoState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ProvisioningInfoState(*pb)
	return &st, nil
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

func ProvisioningPhaseToPb(st *ProvisioningPhase) (*databasepb.ProvisioningPhasePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := databasepb.ProvisioningPhasePb(*st)
	return &pb, nil
}

func ProvisioningPhaseFromPb(pb *databasepb.ProvisioningPhasePb) (*ProvisioningPhase, error) {
	if pb == nil {
		return nil, nil
	}
	st := ProvisioningPhase(*pb)
	return &st, nil
}

type RequestedClaims struct {

	// Wire name: 'permission_set'
	PermissionSet RequestedClaimsPermissionSet ``

	// Wire name: 'resources'
	Resources []RequestedResource ``
}

func RequestedClaimsToPb(st *RequestedClaims) (*databasepb.RequestedClaimsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.RequestedClaimsPb{}
	permissionSetPb, err := RequestedClaimsPermissionSetToPb(&st.PermissionSet)
	if err != nil {
		return nil, err
	}
	if permissionSetPb != nil {
		pb.PermissionSet = *permissionSetPb
	}

	var resourcesPb []databasepb.RequestedResourcePb
	for _, item := range st.Resources {
		itemPb, err := RequestedResourceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

	return pb, nil
}

func RequestedClaimsFromPb(pb *databasepb.RequestedClaimsPb) (*RequestedClaims, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RequestedClaims{}
	permissionSetField, err := RequestedClaimsPermissionSetFromPb(&pb.PermissionSet)
	if err != nil {
		return nil, err
	}
	if permissionSetField != nil {
		st.PermissionSet = *permissionSetField
	}

	var resourcesField []RequestedResource
	for _, itemPb := range pb.Resources {
		item, err := RequestedResourceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField

	return st, nil
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

func RequestedClaimsPermissionSetToPb(st *RequestedClaimsPermissionSet) (*databasepb.RequestedClaimsPermissionSetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := databasepb.RequestedClaimsPermissionSetPb(*st)
	return &pb, nil
}

func RequestedClaimsPermissionSetFromPb(pb *databasepb.RequestedClaimsPermissionSetPb) (*RequestedClaimsPermissionSet, error) {
	if pb == nil {
		return nil, nil
	}
	st := RequestedClaimsPermissionSet(*pb)
	return &st, nil
}

type RequestedResource struct {

	// Wire name: 'table_name'
	TableName string ``

	// Wire name: 'unspecified_resource_name'
	UnspecifiedResourceName string   ``
	ForceSendFields         []string `tf:"-"`
}

func (s *RequestedResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RequestedResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RequestedResourceToPb(st *RequestedResource) (*databasepb.RequestedResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.RequestedResourcePb{}
	pb.TableName = st.TableName
	pb.UnspecifiedResourceName = st.UnspecifiedResourceName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RequestedResourceFromPb(pb *databasepb.RequestedResourcePb) (*RequestedResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RequestedResource{}
	st.TableName = pb.TableName
	st.UnspecifiedResourceName = pb.UnspecifiedResourceName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Next field marker: 12
type SyncedDatabaseTable struct {
	// Synced Table data synchronization status
	// Wire name: 'data_synchronization_status'
	DataSynchronizationStatus *SyncedTableStatus ``
	// Name of the target database instance. This is required when creating
	// synced database tables in standard catalogs. This is optional when
	// creating synced database tables in registered catalogs. If this field is
	// specified when creating synced database tables in registered catalogs,
	// the database instance name MUST match that of the registered catalog (or
	// the request will be rejected).
	// Wire name: 'database_instance_name'
	DatabaseInstanceName string ``
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
	LogicalDatabaseName string ``
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'spec'
	Spec *SyncedTableSpec ``
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	// Wire name: 'unity_catalog_provisioning_state'
	UnityCatalogProvisioningState ProvisioningInfoState ``
	ForceSendFields               []string              `tf:"-"`
}

func (s *SyncedDatabaseTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedDatabaseTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SyncedDatabaseTableToPb(st *SyncedDatabaseTable) (*databasepb.SyncedDatabaseTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedDatabaseTablePb{}
	dataSynchronizationStatusPb, err := SyncedTableStatusToPb(st.DataSynchronizationStatus)
	if err != nil {
		return nil, err
	}
	if dataSynchronizationStatusPb != nil {
		pb.DataSynchronizationStatus = dataSynchronizationStatusPb
	}
	pb.DatabaseInstanceName = st.DatabaseInstanceName
	pb.LogicalDatabaseName = st.LogicalDatabaseName
	pb.Name = st.Name
	specPb, err := SyncedTableSpecToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}
	unityCatalogProvisioningStatePb, err := ProvisioningInfoStateToPb(&st.UnityCatalogProvisioningState)
	if err != nil {
		return nil, err
	}
	if unityCatalogProvisioningStatePb != nil {
		pb.UnityCatalogProvisioningState = *unityCatalogProvisioningStatePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SyncedDatabaseTableFromPb(pb *databasepb.SyncedDatabaseTablePb) (*SyncedDatabaseTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedDatabaseTable{}
	dataSynchronizationStatusField, err := SyncedTableStatusFromPb(pb.DataSynchronizationStatus)
	if err != nil {
		return nil, err
	}
	if dataSynchronizationStatusField != nil {
		st.DataSynchronizationStatus = dataSynchronizationStatusField
	}
	st.DatabaseInstanceName = pb.DatabaseInstanceName
	st.LogicalDatabaseName = pb.LogicalDatabaseName
	st.Name = pb.Name
	specField, err := SyncedTableSpecFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}
	unityCatalogProvisioningStateField, err := ProvisioningInfoStateFromPb(&pb.UnityCatalogProvisioningState)
	if err != nil {
		return nil, err
	}
	if unityCatalogProvisioningStateField != nil {
		st.UnityCatalogProvisioningState = *unityCatalogProvisioningStateField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
type SyncedTableContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *SyncedTablePipelineProgress ``
	// The last source table Delta version that was successfully synced to the
	// synced table.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 ``
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	// Wire name: 'timestamp'
	Timestamp       *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *SyncedTableContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SyncedTableContinuousUpdateStatusToPb(st *SyncedTableContinuousUpdateStatus) (*databasepb.SyncedTableContinuousUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTableContinuousUpdateStatusPb{}
	initialPipelineSyncProgressPb, err := SyncedTablePipelineProgressToPb(st.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressPb != nil {
		pb.InitialPipelineSyncProgress = initialPipelineSyncProgressPb
	}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	timestampPb, err := timestampToPb(st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SyncedTableContinuousUpdateStatusFromPb(pb *databasepb.SyncedTableContinuousUpdateStatusPb) (*SyncedTableContinuousUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableContinuousUpdateStatus{}
	initialPipelineSyncProgressField, err := SyncedTablePipelineProgressFromPb(pb.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressField != nil {
		st.InitialPipelineSyncProgress = initialPipelineSyncProgressField
	}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	timestampField, err := timestampFromPb(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = timestampField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Detailed status of a synced table. Shown if the synced table is in the
// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
type SyncedTableFailedStatus struct {
	// The last source table Delta version that was successfully synced to the
	// synced table. The last source table Delta version that was synced to the
	// synced table. Only populated if the table is still synced and available
	// for serving.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 ``
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. Only populated if the table is still
	// synced and available for serving.
	// Wire name: 'timestamp'
	Timestamp       *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *SyncedTableFailedStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableFailedStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SyncedTableFailedStatusToPb(st *SyncedTableFailedStatus) (*databasepb.SyncedTableFailedStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTableFailedStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	timestampPb, err := timestampToPb(st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SyncedTableFailedStatusFromPb(pb *databasepb.SyncedTableFailedStatusPb) (*SyncedTableFailedStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableFailedStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	timestampField, err := timestampFromPb(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = timestampField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Progress information of the Synced Table data synchronization pipeline.
type SyncedTablePipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	// Wire name: 'estimated_completion_time_seconds'
	EstimatedCompletionTimeSeconds float64 ``
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	// Wire name: 'latest_version_currently_processing'
	LatestVersionCurrentlyProcessing int64 ``
	// The current phase of the data synchronization pipeline.
	// Wire name: 'provisioning_phase'
	ProvisioningPhase ProvisioningPhase ``
	// The completion ratio of this update. This is a number between 0 and 1.
	// Wire name: 'sync_progress_completion'
	SyncProgressCompletion float64 ``
	// The number of rows that have been synced in this update.
	// Wire name: 'synced_row_count'
	SyncedRowCount int64 ``
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	// Wire name: 'total_row_count'
	TotalRowCount   int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *SyncedTablePipelineProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTablePipelineProgress) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SyncedTablePipelineProgressToPb(st *SyncedTablePipelineProgress) (*databasepb.SyncedTablePipelineProgressPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTablePipelineProgressPb{}
	pb.EstimatedCompletionTimeSeconds = st.EstimatedCompletionTimeSeconds
	pb.LatestVersionCurrentlyProcessing = st.LatestVersionCurrentlyProcessing
	provisioningPhasePb, err := ProvisioningPhaseToPb(&st.ProvisioningPhase)
	if err != nil {
		return nil, err
	}
	if provisioningPhasePb != nil {
		pb.ProvisioningPhase = *provisioningPhasePb
	}
	pb.SyncProgressCompletion = st.SyncProgressCompletion
	pb.SyncedRowCount = st.SyncedRowCount
	pb.TotalRowCount = st.TotalRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SyncedTablePipelineProgressFromPb(pb *databasepb.SyncedTablePipelineProgressPb) (*SyncedTablePipelineProgress, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTablePipelineProgress{}
	st.EstimatedCompletionTimeSeconds = pb.EstimatedCompletionTimeSeconds
	st.LatestVersionCurrentlyProcessing = pb.LatestVersionCurrentlyProcessing
	provisioningPhaseField, err := ProvisioningPhaseFromPb(&pb.ProvisioningPhase)
	if err != nil {
		return nil, err
	}
	if provisioningPhaseField != nil {
		st.ProvisioningPhase = *provisioningPhaseField
	}
	st.SyncProgressCompletion = pb.SyncProgressCompletion
	st.SyncedRowCount = pb.SyncedRowCount
	st.TotalRowCount = pb.TotalRowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SyncedTablePosition struct {

	// Wire name: 'delta_table_sync_info'
	DeltaTableSyncInfo *DeltaTableSyncInfo ``
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	// Wire name: 'sync_end_timestamp'
	SyncEndTimestamp *time.Time ``
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	// Wire name: 'sync_start_timestamp'
	SyncStartTimestamp *time.Time ``
	ForceSendFields    []string   `tf:"-"`
}

func (s *SyncedTablePosition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTablePosition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SyncedTablePositionToPb(st *SyncedTablePosition) (*databasepb.SyncedTablePositionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTablePositionPb{}
	deltaTableSyncInfoPb, err := DeltaTableSyncInfoToPb(st.DeltaTableSyncInfo)
	if err != nil {
		return nil, err
	}
	if deltaTableSyncInfoPb != nil {
		pb.DeltaTableSyncInfo = deltaTableSyncInfoPb
	}
	syncEndTimestampPb, err := timestampToPb(st.SyncEndTimestamp)
	if err != nil {
		return nil, err
	}
	if syncEndTimestampPb != nil {
		pb.SyncEndTimestamp = *syncEndTimestampPb
	}
	syncStartTimestampPb, err := timestampToPb(st.SyncStartTimestamp)
	if err != nil {
		return nil, err
	}
	if syncStartTimestampPb != nil {
		pb.SyncStartTimestamp = *syncStartTimestampPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SyncedTablePositionFromPb(pb *databasepb.SyncedTablePositionPb) (*SyncedTablePosition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTablePosition{}
	deltaTableSyncInfoField, err := DeltaTableSyncInfoFromPb(pb.DeltaTableSyncInfo)
	if err != nil {
		return nil, err
	}
	if deltaTableSyncInfoField != nil {
		st.DeltaTableSyncInfo = deltaTableSyncInfoField
	}
	syncEndTimestampField, err := timestampFromPb(&pb.SyncEndTimestamp)
	if err != nil {
		return nil, err
	}
	if syncEndTimestampField != nil {
		st.SyncEndTimestamp = syncEndTimestampField
	}
	syncStartTimestampField, err := timestampFromPb(&pb.SyncStartTimestamp)
	if err != nil {
		return nil, err
	}
	if syncStartTimestampField != nil {
		st.SyncStartTimestamp = syncStartTimestampField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Detailed status of a synced table. Shown if the synced table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type SyncedTableProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *SyncedTablePipelineProgress ``
}

func SyncedTableProvisioningStatusToPb(st *SyncedTableProvisioningStatus) (*databasepb.SyncedTableProvisioningStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTableProvisioningStatusPb{}
	initialPipelineSyncProgressPb, err := SyncedTablePipelineProgressToPb(st.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressPb != nil {
		pb.InitialPipelineSyncProgress = initialPipelineSyncProgressPb
	}

	return pb, nil
}

func SyncedTableProvisioningStatusFromPb(pb *databasepb.SyncedTableProvisioningStatusPb) (*SyncedTableProvisioningStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableProvisioningStatus{}
	initialPipelineSyncProgressField, err := SyncedTablePipelineProgressFromPb(pb.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressField != nil {
		st.InitialPipelineSyncProgress = initialPipelineSyncProgressField
	}

	return st, nil
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

func SyncedTableSchedulingPolicyToPb(st *SyncedTableSchedulingPolicy) (*databasepb.SyncedTableSchedulingPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := databasepb.SyncedTableSchedulingPolicyPb(*st)
	return &pb, nil
}

func SyncedTableSchedulingPolicyFromPb(pb *databasepb.SyncedTableSchedulingPolicyPb) (*SyncedTableSchedulingPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := SyncedTableSchedulingPolicy(*pb)
	return &st, nil
}

// Specification of a synced database table.
type SyncedTableSpec struct {
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist.
	// Wire name: 'create_database_objects_if_missing'
	CreateDatabaseObjectsIfMissing bool ``
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If existing_pipeline_id is defined, the synced table will be bin packed
	// into the existing pipeline referenced. This avoids creating a new
	// pipeline and allows sharing existing compute. In this case, the
	// scheduling_policy of this synced table must match the scheduling policy
	// of the existing pipeline.
	// Wire name: 'existing_pipeline_id'
	ExistingPipelineId string ``
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If new_pipeline_spec is defined, a new pipeline is created for this
	// synced table. The location pointed to is used to store intermediate files
	// (checkpoints, event logs etc). The caller must have write permissions to
	// create Delta tables in the specified catalog and schema. Again, note this
	// requires write permissions, whereas the source table only requires read
	// permissions.
	// Wire name: 'new_pipeline_spec'
	NewPipelineSpec *NewPipelineSpec ``
	// Primary Key columns to be used for data insert/update in the destination.
	// Wire name: 'primary_key_columns'
	PrimaryKeyColumns []string ``
	// Scheduling policy of the underlying pipeline.
	// Wire name: 'scheduling_policy'
	SchedulingPolicy SyncedTableSchedulingPolicy ``
	// Three-part (catalog, schema, table) name of the source Delta table.
	// Wire name: 'source_table_full_name'
	SourceTableFullName string ``
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	// Wire name: 'timeseries_key'
	TimeseriesKey   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *SyncedTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SyncedTableSpecToPb(st *SyncedTableSpec) (*databasepb.SyncedTableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTableSpecPb{}
	pb.CreateDatabaseObjectsIfMissing = st.CreateDatabaseObjectsIfMissing
	pb.ExistingPipelineId = st.ExistingPipelineId
	newPipelineSpecPb, err := NewPipelineSpecToPb(st.NewPipelineSpec)
	if err != nil {
		return nil, err
	}
	if newPipelineSpecPb != nil {
		pb.NewPipelineSpec = newPipelineSpecPb
	}
	pb.PrimaryKeyColumns = st.PrimaryKeyColumns
	schedulingPolicyPb, err := SyncedTableSchedulingPolicyToPb(&st.SchedulingPolicy)
	if err != nil {
		return nil, err
	}
	if schedulingPolicyPb != nil {
		pb.SchedulingPolicy = *schedulingPolicyPb
	}
	pb.SourceTableFullName = st.SourceTableFullName
	pb.TimeseriesKey = st.TimeseriesKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SyncedTableSpecFromPb(pb *databasepb.SyncedTableSpecPb) (*SyncedTableSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableSpec{}
	st.CreateDatabaseObjectsIfMissing = pb.CreateDatabaseObjectsIfMissing
	st.ExistingPipelineId = pb.ExistingPipelineId
	newPipelineSpecField, err := NewPipelineSpecFromPb(pb.NewPipelineSpec)
	if err != nil {
		return nil, err
	}
	if newPipelineSpecField != nil {
		st.NewPipelineSpec = newPipelineSpecField
	}
	st.PrimaryKeyColumns = pb.PrimaryKeyColumns
	schedulingPolicyField, err := SyncedTableSchedulingPolicyFromPb(&pb.SchedulingPolicy)
	if err != nil {
		return nil, err
	}
	if schedulingPolicyField != nil {
		st.SchedulingPolicy = *schedulingPolicyField
	}
	st.SourceTableFullName = pb.SourceTableFullName
	st.TimeseriesKey = pb.TimeseriesKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func SyncedTableStateToPb(st *SyncedTableState) (*databasepb.SyncedTableStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := databasepb.SyncedTableStatePb(*st)
	return &pb, nil
}

func SyncedTableStateFromPb(pb *databasepb.SyncedTableStatePb) (*SyncedTableState, error) {
	if pb == nil {
		return nil, nil
	}
	st := SyncedTableState(*pb)
	return &st, nil
}

// Status of a synced table.
type SyncedTableStatus struct {

	// Wire name: 'continuous_update_status'
	ContinuousUpdateStatus *SyncedTableContinuousUpdateStatus ``
	// The state of the synced table.
	// Wire name: 'detailed_state'
	DetailedState SyncedTableState ``

	// Wire name: 'failed_status'
	FailedStatus *SyncedTableFailedStatus ``
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
	LastSync *SyncedTablePosition ``
	// A text description of the current state of the synced table.
	// Wire name: 'message'
	Message string ``
	// ID of the associated pipeline. The pipeline ID may have been provided by
	// the client (in the case of bin packing), or generated by the server (when
	// creating a new pipeline).
	// Wire name: 'pipeline_id'
	PipelineId string ``

	// Wire name: 'provisioning_status'
	ProvisioningStatus *SyncedTableProvisioningStatus ``

	// Wire name: 'triggered_update_status'
	TriggeredUpdateStatus *SyncedTableTriggeredUpdateStatus ``
	ForceSendFields       []string                          `tf:"-"`
}

func (s *SyncedTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SyncedTableStatusToPb(st *SyncedTableStatus) (*databasepb.SyncedTableStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTableStatusPb{}
	continuousUpdateStatusPb, err := SyncedTableContinuousUpdateStatusToPb(st.ContinuousUpdateStatus)
	if err != nil {
		return nil, err
	}
	if continuousUpdateStatusPb != nil {
		pb.ContinuousUpdateStatus = continuousUpdateStatusPb
	}
	detailedStatePb, err := SyncedTableStateToPb(&st.DetailedState)
	if err != nil {
		return nil, err
	}
	if detailedStatePb != nil {
		pb.DetailedState = *detailedStatePb
	}
	failedStatusPb, err := SyncedTableFailedStatusToPb(st.FailedStatus)
	if err != nil {
		return nil, err
	}
	if failedStatusPb != nil {
		pb.FailedStatus = failedStatusPb
	}
	lastSyncPb, err := SyncedTablePositionToPb(st.LastSync)
	if err != nil {
		return nil, err
	}
	if lastSyncPb != nil {
		pb.LastSync = lastSyncPb
	}
	pb.Message = st.Message
	pb.PipelineId = st.PipelineId
	provisioningStatusPb, err := SyncedTableProvisioningStatusToPb(st.ProvisioningStatus)
	if err != nil {
		return nil, err
	}
	if provisioningStatusPb != nil {
		pb.ProvisioningStatus = provisioningStatusPb
	}
	triggeredUpdateStatusPb, err := SyncedTableTriggeredUpdateStatusToPb(st.TriggeredUpdateStatus)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateStatusPb != nil {
		pb.TriggeredUpdateStatus = triggeredUpdateStatusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SyncedTableStatusFromPb(pb *databasepb.SyncedTableStatusPb) (*SyncedTableStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableStatus{}
	continuousUpdateStatusField, err := SyncedTableContinuousUpdateStatusFromPb(pb.ContinuousUpdateStatus)
	if err != nil {
		return nil, err
	}
	if continuousUpdateStatusField != nil {
		st.ContinuousUpdateStatus = continuousUpdateStatusField
	}
	detailedStateField, err := SyncedTableStateFromPb(&pb.DetailedState)
	if err != nil {
		return nil, err
	}
	if detailedStateField != nil {
		st.DetailedState = *detailedStateField
	}
	failedStatusField, err := SyncedTableFailedStatusFromPb(pb.FailedStatus)
	if err != nil {
		return nil, err
	}
	if failedStatusField != nil {
		st.FailedStatus = failedStatusField
	}
	lastSyncField, err := SyncedTablePositionFromPb(pb.LastSync)
	if err != nil {
		return nil, err
	}
	if lastSyncField != nil {
		st.LastSync = lastSyncField
	}
	st.Message = pb.Message
	st.PipelineId = pb.PipelineId
	provisioningStatusField, err := SyncedTableProvisioningStatusFromPb(pb.ProvisioningStatus)
	if err != nil {
		return nil, err
	}
	if provisioningStatusField != nil {
		st.ProvisioningStatus = provisioningStatusField
	}
	triggeredUpdateStatusField, err := SyncedTableTriggeredUpdateStatusFromPb(pb.TriggeredUpdateStatus)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateStatusField != nil {
		st.TriggeredUpdateStatus = triggeredUpdateStatusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
type SyncedTableTriggeredUpdateStatus struct {
	// The last source table Delta version that was successfully synced to the
	// synced table.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 ``
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	// Wire name: 'timestamp'
	Timestamp *time.Time ``
	// Progress of the active data synchronization pipeline.
	// Wire name: 'triggered_update_progress'
	TriggeredUpdateProgress *SyncedTablePipelineProgress ``
	ForceSendFields         []string                     `tf:"-"`
}

func (s *SyncedTableTriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableTriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SyncedTableTriggeredUpdateStatusToPb(st *SyncedTableTriggeredUpdateStatus) (*databasepb.SyncedTableTriggeredUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTableTriggeredUpdateStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	timestampPb, err := timestampToPb(st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}
	triggeredUpdateProgressPb, err := SyncedTablePipelineProgressToPb(st.TriggeredUpdateProgress)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateProgressPb != nil {
		pb.TriggeredUpdateProgress = triggeredUpdateProgressPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SyncedTableTriggeredUpdateStatusFromPb(pb *databasepb.SyncedTableTriggeredUpdateStatusPb) (*SyncedTableTriggeredUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableTriggeredUpdateStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	timestampField, err := timestampFromPb(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = timestampField
	}
	triggeredUpdateProgressField, err := SyncedTablePipelineProgressFromPb(pb.TriggeredUpdateProgress)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateProgressField != nil {
		st.TriggeredUpdateProgress = triggeredUpdateProgressField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateDatabaseInstanceRequest struct {

	// Wire name: 'database_instance'
	DatabaseInstance DatabaseInstance ``
	// The name of the instance. This is the unique identifier for the instance.
	// Wire name: 'name'
	Name string `tf:"-"`
	// The list of fields to update. This field is not yet supported, and is
	// ignored by the server.
	// Wire name: 'update_mask'
	UpdateMask []string `tf:"-"`
}

func UpdateDatabaseInstanceRequestToPb(st *UpdateDatabaseInstanceRequest) (*databasepb.UpdateDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.UpdateDatabaseInstanceRequestPb{}
	databaseInstancePb, err := DatabaseInstanceToPb(&st.DatabaseInstance)
	if err != nil {
		return nil, err
	}
	if databaseInstancePb != nil {
		pb.DatabaseInstance = *databaseInstancePb
	}
	pb.Name = st.Name
	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	return pb, nil
}

func UpdateDatabaseInstanceRequestFromPb(pb *databasepb.UpdateDatabaseInstanceRequestPb) (*UpdateDatabaseInstanceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDatabaseInstanceRequest{}
	databaseInstanceField, err := DatabaseInstanceFromPb(&pb.DatabaseInstance)
	if err != nil {
		return nil, err
	}
	if databaseInstanceField != nil {
		st.DatabaseInstance = *databaseInstanceField
	}
	st.Name = pb.Name
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	return st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%.9fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
