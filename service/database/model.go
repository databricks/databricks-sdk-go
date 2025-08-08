// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/database/databasepb"
)

type CreateDatabaseCatalogRequest struct {

	// Wire name: 'catalog'
	Catalog DatabaseCatalog `json:"catalog"`
}

func (st CreateDatabaseCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.CreateDatabaseCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateDatabaseCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	DatabaseInstance DatabaseInstance `json:"database_instance"`
}

func (st CreateDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.CreateDatabaseInstanceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateDatabaseInstanceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	DatabaseInstanceRole DatabaseInstanceRole `json:"database_instance_role"`

	InstanceName string `json:"-" tf:"-"`
}

func (st CreateDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.CreateDatabaseInstanceRoleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateDatabaseInstanceRoleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Table DatabaseTable `json:"table"`
}

func (st CreateDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.CreateDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	SyncedTable SyncedDatabaseTable `json:"synced_table"`
}

func (st CreateSyncedDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.CreateSyncedDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateSyncedDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Uid             string   `json:"uid,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseCatalog) MarshalJSON() ([]byte, error) {
	pb, err := DatabaseCatalogToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DatabaseCatalog) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DatabaseCatalogPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatabaseCatalogFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DatabaseCredential struct {

	// Wire name: 'expiration_time'
	ExpirationTime string `json:"expiration_time,omitempty"` //legacy

	// Wire name: 'token'
	Token           string   `json:"token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseCredential) MarshalJSON() ([]byte, error) {
	pb, err := DatabaseCredentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DatabaseCredential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DatabaseCredentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatabaseCredentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DatabaseCredentialToPb(st *DatabaseCredential) (*databasepb.DatabaseCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseCredentialPb{}
	pb.ExpirationTime = st.ExpirationTime
	pb.Token = st.Token

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DatabaseCredentialFromPb(pb *databasepb.DatabaseCredentialPb) (*DatabaseCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseCredential{}
	st.ExpirationTime = pb.ExpirationTime
	st.Token = pb.Token

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	CreationTime string `json:"creation_time,omitempty"` //legacy
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
	Uid             string   `json:"uid,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseInstance) MarshalJSON() ([]byte, error) {
	pb, err := DatabaseInstanceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DatabaseInstance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DatabaseInstancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatabaseInstanceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	pb.CreationTime = st.CreationTime
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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
	st.CreationTime = pb.CreationTime
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	BranchTime string `json:"branch_time,omitempty"` //legacy
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
	Uid             string   `json:"uid,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseInstanceRef) MarshalJSON() ([]byte, error) {
	pb, err := DatabaseInstanceRefToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DatabaseInstanceRef) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DatabaseInstanceRefPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatabaseInstanceRefFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DatabaseInstanceRefToPb(st *DatabaseInstanceRef) (*databasepb.DatabaseInstanceRefPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseInstanceRefPb{}
	pb.BranchTime = st.BranchTime
	pb.EffectiveLsn = st.EffectiveLsn
	pb.Lsn = st.Lsn
	pb.Name = st.Name
	pb.Uid = st.Uid

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DatabaseInstanceRefFromPb(pb *databasepb.DatabaseInstanceRefPb) (*DatabaseInstanceRef, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseInstanceRef{}
	st.BranchTime = pb.BranchTime
	st.EffectiveLsn = pb.EffectiveLsn
	st.Lsn = pb.Lsn
	st.Name = pb.Name
	st.Uid = pb.Uid

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseInstanceRole) MarshalJSON() ([]byte, error) {
	pb, err := DatabaseInstanceRoleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DatabaseInstanceRole) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DatabaseInstanceRolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatabaseInstanceRoleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	Bypassrls bool `json:"bypassrls,omitempty"`

	// Wire name: 'createdb'
	Createdb bool `json:"createdb,omitempty"`

	// Wire name: 'createrole'
	Createrole      bool     `json:"createrole,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseInstanceRoleAttributes) MarshalJSON() ([]byte, error) {
	pb, err := DatabaseInstanceRoleAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DatabaseInstanceRoleAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DatabaseInstanceRoleAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatabaseInstanceRoleAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DatabaseInstanceRoleAttributesToPb(st *DatabaseInstanceRoleAttributes) (*databasepb.DatabaseInstanceRoleAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseInstanceRoleAttributesPb{}
	pb.Bypassrls = st.Bypassrls
	pb.Createdb = st.Createdb
	pb.Createrole = st.Createrole

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	Name            string   `json:"name"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabaseTable) MarshalJSON() ([]byte, error) {
	pb, err := DatabaseTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DatabaseTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DatabaseTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatabaseTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DatabaseTableToPb(st *DatabaseTable) (*databasepb.DatabaseTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DatabaseTablePb{}
	pb.DatabaseInstanceName = st.DatabaseInstanceName
	pb.LogicalDatabaseName = st.LogicalDatabaseName
	pb.Name = st.Name

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteDatabaseCatalogRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st DeleteDatabaseCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DeleteDatabaseCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDatabaseCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Purge           bool     `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DeleteDatabaseInstanceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDatabaseInstanceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDatabaseInstanceRequestToPb(st *DeleteDatabaseInstanceRequest) (*databasepb.DeleteDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DeleteDatabaseInstanceRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name
	pb.Purge = st.Purge

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteDatabaseInstanceRoleRequest struct {
	// This is the AIP standard name for the equivalent of Postgres' `IF EXISTS`
	// option
	AllowMissing bool `json:"-" tf:"-"`

	InstanceName string `json:"-" tf:"-"`

	Name string `json:"-" tf:"-"`

	ReassignOwnedTo string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DeleteDatabaseInstanceRoleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDatabaseInstanceRoleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteDatabaseTableRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st DeleteDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DeleteDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st DeleteSyncedDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DeleteSyncedDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteSyncedDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	DeltaCommitTimestamp string `json:"delta_commit_timestamp,omitempty"` //legacy
	// The Delta Lake commit version that was last successfully synced.
	// Wire name: 'delta_commit_version'
	DeltaCommitVersion int64    `json:"delta_commit_version,omitempty"`
	ForceSendFields    []string `json:"-" tf:"-"`
}

func (st DeltaTableSyncInfo) MarshalJSON() ([]byte, error) {
	pb, err := DeltaTableSyncInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeltaTableSyncInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.DeltaTableSyncInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeltaTableSyncInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeltaTableSyncInfoToPb(st *DeltaTableSyncInfo) (*databasepb.DeltaTableSyncInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.DeltaTableSyncInfoPb{}
	pb.DeltaCommitTimestamp = st.DeltaCommitTimestamp
	pb.DeltaCommitVersion = st.DeltaCommitVersion

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeltaTableSyncInfoFromPb(pb *databasepb.DeltaTableSyncInfoPb) (*DeltaTableSyncInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaTableSyncInfo{}
	st.DeltaCommitTimestamp = pb.DeltaCommitTimestamp
	st.DeltaCommitVersion = pb.DeltaCommitVersion

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FindDatabaseInstanceByUidRequest struct {
	// UID of the cluster to get.
	Uid             string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FindDatabaseInstanceByUidRequest) MarshalJSON() ([]byte, error) {
	pb, err := FindDatabaseInstanceByUidRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FindDatabaseInstanceByUidRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.FindDatabaseInstanceByUidRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FindDatabaseInstanceByUidRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FindDatabaseInstanceByUidRequestToPb(st *FindDatabaseInstanceByUidRequest) (*databasepb.FindDatabaseInstanceByUidRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.FindDatabaseInstanceByUidRequestPb{}
	pb.Uid = st.Uid

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FindDatabaseInstanceByUidRequestFromPb(pb *databasepb.FindDatabaseInstanceByUidRequestPb) (*FindDatabaseInstanceByUidRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FindDatabaseInstanceByUidRequest{}
	st.Uid = pb.Uid

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	RequestId       string   `json:"request_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GenerateDatabaseCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := GenerateDatabaseCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GenerateDatabaseCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.GenerateDatabaseCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GenerateDatabaseCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetDatabaseCatalogRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st GetDatabaseCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDatabaseCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.GetDatabaseCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDatabaseCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st GetDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.GetDatabaseInstanceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDatabaseInstanceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	InstanceName string `json:"-" tf:"-"`

	Name string `json:"-" tf:"-"`
}

func (st GetDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDatabaseInstanceRoleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.GetDatabaseInstanceRoleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDatabaseInstanceRoleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st GetDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.GetDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st GetSyncedDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetSyncedDatabaseTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.GetSyncedDatabaseTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetSyncedDatabaseTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	InstanceName string `json:"-" tf:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDatabaseInstanceRolesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListDatabaseInstanceRolesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListDatabaseInstanceRolesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.ListDatabaseInstanceRolesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListDatabaseInstanceRolesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListDatabaseInstanceRolesRequestToPb(st *ListDatabaseInstanceRolesRequest) (*databasepb.ListDatabaseInstanceRolesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.ListDatabaseInstanceRolesRequestPb{}
	pb.InstanceName = st.InstanceName
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListDatabaseInstanceRolesResponse struct {
	// List of database instance roles.
	// Wire name: 'database_instance_roles'
	DatabaseInstanceRoles []DatabaseInstanceRole `json:"database_instance_roles,omitempty"`
	// Pagination token to request the next page of instances.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDatabaseInstanceRolesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListDatabaseInstanceRolesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListDatabaseInstanceRolesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.ListDatabaseInstanceRolesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListDatabaseInstanceRolesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListDatabaseInstancesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDatabaseInstancesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListDatabaseInstancesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListDatabaseInstancesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.ListDatabaseInstancesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListDatabaseInstancesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListDatabaseInstancesRequestToPb(st *ListDatabaseInstancesRequest) (*databasepb.ListDatabaseInstancesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.ListDatabaseInstancesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListDatabaseInstancesRequestFromPb(pb *databasepb.ListDatabaseInstancesRequestPb) (*ListDatabaseInstancesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDatabaseInstancesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListDatabaseInstancesResponse struct {
	// List of instances.
	// Wire name: 'database_instances'
	DatabaseInstances []DatabaseInstance `json:"database_instances,omitempty"`
	// Pagination token to request the next page of instances.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDatabaseInstancesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListDatabaseInstancesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListDatabaseInstancesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.ListDatabaseInstancesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListDatabaseInstancesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	StorageCatalog string `json:"storage_catalog,omitempty"`
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	// Wire name: 'storage_schema'
	StorageSchema   string   `json:"storage_schema,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NewPipelineSpec) MarshalJSON() ([]byte, error) {
	pb, err := NewPipelineSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NewPipelineSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.NewPipelineSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NewPipelineSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NewPipelineSpecToPb(st *NewPipelineSpec) (*databasepb.NewPipelineSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.NewPipelineSpecPb{}
	pb.StorageCatalog = st.StorageCatalog
	pb.StorageSchema = st.StorageSchema

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NewPipelineSpecFromPb(pb *databasepb.NewPipelineSpecPb) (*NewPipelineSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NewPipelineSpec{}
	st.StorageCatalog = pb.StorageCatalog
	st.StorageSchema = pb.StorageSchema

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	PermissionSet RequestedClaimsPermissionSet `json:"permission_set,omitempty"`

	// Wire name: 'resources'
	Resources []RequestedResource `json:"resources,omitempty"`
}

func (st RequestedClaims) MarshalJSON() ([]byte, error) {
	pb, err := RequestedClaimsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RequestedClaims) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.RequestedClaimsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RequestedClaimsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	TableName string `json:"table_name,omitempty"`

	// Wire name: 'unspecified_resource_name'
	UnspecifiedResourceName string   `json:"unspecified_resource_name,omitempty"`
	ForceSendFields         []string `json:"-" tf:"-"`
}

func (st RequestedResource) MarshalJSON() ([]byte, error) {
	pb, err := RequestedResourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RequestedResource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.RequestedResourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RequestedResourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RequestedResourceToPb(st *RequestedResource) (*databasepb.RequestedResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.RequestedResourcePb{}
	pb.TableName = st.TableName
	pb.UnspecifiedResourceName = st.UnspecifiedResourceName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RequestedResourceFromPb(pb *databasepb.RequestedResourcePb) (*RequestedResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RequestedResource{}
	st.TableName = pb.TableName
	st.UnspecifiedResourceName = pb.UnspecifiedResourceName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	ForceSendFields               []string              `json:"-" tf:"-"`
}

func (st SyncedDatabaseTable) MarshalJSON() ([]byte, error) {
	pb, err := SyncedDatabaseTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedDatabaseTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedDatabaseTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedDatabaseTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	Timestamp       string   `json:"timestamp,omitempty"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTableContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	pb, err := SyncedTableContinuousUpdateStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedTableContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedTableContinuousUpdateStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedTableContinuousUpdateStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	pb.Timestamp = st.Timestamp

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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
	st.Timestamp = pb.Timestamp

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. Only populated if the table is still
	// synced and available for serving.
	// Wire name: 'timestamp'
	Timestamp       string   `json:"timestamp,omitempty"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTableFailedStatus) MarshalJSON() ([]byte, error) {
	pb, err := SyncedTableFailedStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedTableFailedStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedTableFailedStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedTableFailedStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SyncedTableFailedStatusToPb(st *SyncedTableFailedStatus) (*databasepb.SyncedTableFailedStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTableFailedStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	pb.Timestamp = st.Timestamp

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SyncedTableFailedStatusFromPb(pb *databasepb.SyncedTableFailedStatusPb) (*SyncedTableFailedStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableFailedStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	// The current phase of the data synchronization pipeline.
	// Wire name: 'provisioning_phase'
	ProvisioningPhase ProvisioningPhase `json:"provisioning_phase,omitempty"`
	// The completion ratio of this update. This is a number between 0 and 1.
	// Wire name: 'sync_progress_completion'
	SyncProgressCompletion float64 `json:"sync_progress_completion,omitempty"`
	// The number of rows that have been synced in this update.
	// Wire name: 'synced_row_count'
	SyncedRowCount int64 `json:"synced_row_count,omitempty"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	// Wire name: 'total_row_count'
	TotalRowCount   int64    `json:"total_row_count,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTablePipelineProgress) MarshalJSON() ([]byte, error) {
	pb, err := SyncedTablePipelineProgressToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedTablePipelineProgress) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedTablePipelineProgressPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedTablePipelineProgressFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SyncedTablePosition struct {

	// Wire name: 'delta_table_sync_info'
	DeltaTableSyncInfo *DeltaTableSyncInfo `json:"delta_table_sync_info,omitempty"`
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	// Wire name: 'sync_end_timestamp'
	SyncEndTimestamp string `json:"sync_end_timestamp,omitempty"` //legacy
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	// Wire name: 'sync_start_timestamp'
	SyncStartTimestamp string   `json:"sync_start_timestamp,omitempty"` //legacy
	ForceSendFields    []string `json:"-" tf:"-"`
}

func (st SyncedTablePosition) MarshalJSON() ([]byte, error) {
	pb, err := SyncedTablePositionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedTablePosition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedTablePositionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedTablePositionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	pb.SyncEndTimestamp = st.SyncEndTimestamp
	pb.SyncStartTimestamp = st.SyncStartTimestamp

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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
	st.SyncEndTimestamp = pb.SyncEndTimestamp
	st.SyncStartTimestamp = pb.SyncStartTimestamp

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Detailed status of a synced table. Shown if the synced table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type SyncedTableProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

func (st SyncedTableProvisioningStatus) MarshalJSON() ([]byte, error) {
	pb, err := SyncedTableProvisioningStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedTableProvisioningStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedTableProvisioningStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedTableProvisioningStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	CreateDatabaseObjectsIfMissing bool `json:"create_database_objects_if_missing,omitempty"`
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If existing_pipeline_id is defined, the synced table will be bin packed
	// into the existing pipeline referenced. This avoids creating a new
	// pipeline and allows sharing existing compute. In this case, the
	// scheduling_policy of this synced table must match the scheduling policy
	// of the existing pipeline.
	// Wire name: 'existing_pipeline_id'
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
	TimeseriesKey   string   `json:"timeseries_key,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SyncedTableSpec) MarshalJSON() ([]byte, error) {
	pb, err := SyncedTableSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedTableSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedTableSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedTableSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	ForceSendFields       []string                          `json:"-" tf:"-"`
}

func (st SyncedTableStatus) MarshalJSON() ([]byte, error) {
	pb, err := SyncedTableStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedTableStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedTableStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedTableStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	Timestamp string `json:"timestamp,omitempty"` //legacy
	// Progress of the active data synchronization pipeline.
	// Wire name: 'triggered_update_progress'
	TriggeredUpdateProgress *SyncedTablePipelineProgress `json:"triggered_update_progress,omitempty"`
	ForceSendFields         []string                     `json:"-" tf:"-"`
}

func (st SyncedTableTriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	pb, err := SyncedTableTriggeredUpdateStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SyncedTableTriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.SyncedTableTriggeredUpdateStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SyncedTableTriggeredUpdateStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SyncedTableTriggeredUpdateStatusToPb(st *SyncedTableTriggeredUpdateStatus) (*databasepb.SyncedTableTriggeredUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databasepb.SyncedTableTriggeredUpdateStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	pb.Timestamp = st.Timestamp
	triggeredUpdateProgressPb, err := SyncedTablePipelineProgressToPb(st.TriggeredUpdateProgress)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateProgressPb != nil {
		pb.TriggeredUpdateProgress = triggeredUpdateProgressPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SyncedTableTriggeredUpdateStatusFromPb(pb *databasepb.SyncedTableTriggeredUpdateStatusPb) (*SyncedTableTriggeredUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableTriggeredUpdateStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp
	triggeredUpdateProgressField, err := SyncedTablePipelineProgressFromPb(pb.TriggeredUpdateProgress)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateProgressField != nil {
		st.TriggeredUpdateProgress = triggeredUpdateProgressField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateDatabaseInstanceRequest struct {

	// Wire name: 'database_instance'
	DatabaseInstance DatabaseInstance `json:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	Name string `json:"-" tf:"-"`
	// The list of fields to update. This field is not yet supported, and is
	// ignored by the server.
	UpdateMask string `json:"-" tf:"-"` //legacy

}

func (st UpdateDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateDatabaseInstanceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databasepb.UpdateDatabaseInstanceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateDatabaseInstanceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	pb.UpdateMask = st.UpdateMask

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
	st.UpdateMask = pb.UpdateMask

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
