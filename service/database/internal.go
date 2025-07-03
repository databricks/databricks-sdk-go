// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func createDatabaseCatalogRequestToPb(st *CreateDatabaseCatalogRequest) (*createDatabaseCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDatabaseCatalogRequestPb{}
	pb.Catalog = st.Catalog

	return pb, nil
}

type createDatabaseCatalogRequestPb struct {
	Catalog DatabaseCatalog `json:"catalog"`
}

func createDatabaseCatalogRequestFromPb(pb *createDatabaseCatalogRequestPb) (*CreateDatabaseCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDatabaseCatalogRequest{}
	st.Catalog = pb.Catalog

	return st, nil
}

func createDatabaseInstanceRequestToPb(st *CreateDatabaseInstanceRequest) (*createDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDatabaseInstanceRequestPb{}
	pb.DatabaseInstance = st.DatabaseInstance

	return pb, nil
}

type createDatabaseInstanceRequestPb struct {
	DatabaseInstance DatabaseInstance `json:"database_instance"`
}

func createDatabaseInstanceRequestFromPb(pb *createDatabaseInstanceRequestPb) (*CreateDatabaseInstanceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDatabaseInstanceRequest{}
	st.DatabaseInstance = pb.DatabaseInstance

	return st, nil
}

func createDatabaseInstanceRoleRequestToPb(st *CreateDatabaseInstanceRoleRequest) (*createDatabaseInstanceRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDatabaseInstanceRoleRequestPb{}
	pb.DatabaseInstanceRole = st.DatabaseInstanceRole
	pb.InstanceName = st.InstanceName

	return pb, nil
}

type createDatabaseInstanceRoleRequestPb struct {
	DatabaseInstanceRole DatabaseInstanceRole `json:"database_instance_role"`
	InstanceName         string               `json:"-" url:"-"`
}

func createDatabaseInstanceRoleRequestFromPb(pb *createDatabaseInstanceRoleRequestPb) (*CreateDatabaseInstanceRoleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDatabaseInstanceRoleRequest{}
	st.DatabaseInstanceRole = pb.DatabaseInstanceRole
	st.InstanceName = pb.InstanceName

	return st, nil
}

func createDatabaseTableRequestToPb(st *CreateDatabaseTableRequest) (*createDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDatabaseTableRequestPb{}
	pb.Table = st.Table

	return pb, nil
}

type createDatabaseTableRequestPb struct {
	Table DatabaseTable `json:"table"`
}

func createDatabaseTableRequestFromPb(pb *createDatabaseTableRequestPb) (*CreateDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDatabaseTableRequest{}
	st.Table = pb.Table

	return st, nil
}

func createSyncedDatabaseTableRequestToPb(st *CreateSyncedDatabaseTableRequest) (*createSyncedDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createSyncedDatabaseTableRequestPb{}
	pb.SyncedTable = st.SyncedTable

	return pb, nil
}

type createSyncedDatabaseTableRequestPb struct {
	SyncedTable SyncedDatabaseTable `json:"synced_table"`
}

func createSyncedDatabaseTableRequestFromPb(pb *createSyncedDatabaseTableRequestPb) (*CreateSyncedDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateSyncedDatabaseTableRequest{}
	st.SyncedTable = pb.SyncedTable

	return st, nil
}

func databaseCatalogToPb(st *DatabaseCatalog) (*databaseCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databaseCatalogPb{}
	pb.CreateDatabaseIfNotExists = st.CreateDatabaseIfNotExists
	pb.DatabaseInstanceName = st.DatabaseInstanceName
	pb.DatabaseName = st.DatabaseName
	pb.Name = st.Name
	pb.Uid = st.Uid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databaseCatalogPb struct {
	CreateDatabaseIfNotExists bool   `json:"create_database_if_not_exists,omitempty"`
	DatabaseInstanceName      string `json:"database_instance_name"`
	DatabaseName              string `json:"database_name"`
	Name                      string `json:"name"`
	Uid                       string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databaseCatalogFromPb(pb *databaseCatalogPb) (*DatabaseCatalog, error) {
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

func (st *databaseCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databaseCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databaseCredentialToPb(st *DatabaseCredential) (*databaseCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databaseCredentialPb{}
	pb.ExpirationTime = st.ExpirationTime
	pb.Token = st.Token

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databaseCredentialPb struct {
	ExpirationTime string `json:"expiration_time,omitempty"`
	Token          string `json:"token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databaseCredentialFromPb(pb *databaseCredentialPb) (*DatabaseCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseCredential{}
	st.ExpirationTime = pb.ExpirationTime
	st.Token = pb.Token

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databaseCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databaseCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databaseInstanceToPb(st *DatabaseInstance) (*databaseInstancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databaseInstancePb{}
	pb.Capacity = st.Capacity
	pb.ChildInstanceRefs = st.ChildInstanceRefs
	pb.CreationTime = st.CreationTime
	pb.Creator = st.Creator
	pb.EffectiveEnableReadableSecondaries = st.EffectiveEnableReadableSecondaries
	pb.EffectiveNodeCount = st.EffectiveNodeCount
	pb.EffectiveRetentionWindowInDays = st.EffectiveRetentionWindowInDays
	pb.EffectiveStopped = st.EffectiveStopped
	pb.EnableReadableSecondaries = st.EnableReadableSecondaries
	pb.Name = st.Name
	pb.NodeCount = st.NodeCount
	pb.ParentInstanceRef = st.ParentInstanceRef
	pb.PgVersion = st.PgVersion
	pb.ReadOnlyDns = st.ReadOnlyDns
	pb.ReadWriteDns = st.ReadWriteDns
	pb.RetentionWindowInDays = st.RetentionWindowInDays
	pb.State = st.State
	pb.Stopped = st.Stopped
	pb.Uid = st.Uid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databaseInstancePb struct {
	Capacity                           string                `json:"capacity,omitempty"`
	ChildInstanceRefs                  []DatabaseInstanceRef `json:"child_instance_refs,omitempty"`
	CreationTime                       string                `json:"creation_time,omitempty"`
	Creator                            string                `json:"creator,omitempty"`
	EffectiveEnableReadableSecondaries bool                  `json:"effective_enable_readable_secondaries,omitempty"`
	EffectiveNodeCount                 int                   `json:"effective_node_count,omitempty"`
	EffectiveRetentionWindowInDays     int                   `json:"effective_retention_window_in_days,omitempty"`
	EffectiveStopped                   bool                  `json:"effective_stopped,omitempty"`
	EnableReadableSecondaries          bool                  `json:"enable_readable_secondaries,omitempty"`
	Name                               string                `json:"name"`
	NodeCount                          int                   `json:"node_count,omitempty"`
	ParentInstanceRef                  *DatabaseInstanceRef  `json:"parent_instance_ref,omitempty"`
	PgVersion                          string                `json:"pg_version,omitempty"`
	ReadOnlyDns                        string                `json:"read_only_dns,omitempty"`
	ReadWriteDns                       string                `json:"read_write_dns,omitempty"`
	RetentionWindowInDays              int                   `json:"retention_window_in_days,omitempty"`
	State                              DatabaseInstanceState `json:"state,omitempty"`
	Stopped                            bool                  `json:"stopped,omitempty"`
	Uid                                string                `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databaseInstanceFromPb(pb *databaseInstancePb) (*DatabaseInstance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseInstance{}
	st.Capacity = pb.Capacity
	st.ChildInstanceRefs = pb.ChildInstanceRefs
	st.CreationTime = pb.CreationTime
	st.Creator = pb.Creator
	st.EffectiveEnableReadableSecondaries = pb.EffectiveEnableReadableSecondaries
	st.EffectiveNodeCount = pb.EffectiveNodeCount
	st.EffectiveRetentionWindowInDays = pb.EffectiveRetentionWindowInDays
	st.EffectiveStopped = pb.EffectiveStopped
	st.EnableReadableSecondaries = pb.EnableReadableSecondaries
	st.Name = pb.Name
	st.NodeCount = pb.NodeCount
	st.ParentInstanceRef = pb.ParentInstanceRef
	st.PgVersion = pb.PgVersion
	st.ReadOnlyDns = pb.ReadOnlyDns
	st.ReadWriteDns = pb.ReadWriteDns
	st.RetentionWindowInDays = pb.RetentionWindowInDays
	st.State = pb.State
	st.Stopped = pb.Stopped
	st.Uid = pb.Uid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databaseInstancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databaseInstancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databaseInstanceRefToPb(st *DatabaseInstanceRef) (*databaseInstanceRefPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databaseInstanceRefPb{}
	pb.BranchTime = st.BranchTime
	pb.EffectiveLsn = st.EffectiveLsn
	pb.Lsn = st.Lsn
	pb.Name = st.Name
	pb.Uid = st.Uid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databaseInstanceRefPb struct {
	BranchTime   string `json:"branch_time,omitempty"`
	EffectiveLsn string `json:"effective_lsn,omitempty"`
	Lsn          string `json:"lsn,omitempty"`
	Name         string `json:"name,omitempty"`
	Uid          string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databaseInstanceRefFromPb(pb *databaseInstanceRefPb) (*DatabaseInstanceRef, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseInstanceRef{}
	st.BranchTime = pb.BranchTime
	st.EffectiveLsn = pb.EffectiveLsn
	st.Lsn = pb.Lsn
	st.Name = pb.Name
	st.Uid = pb.Uid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databaseInstanceRefPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databaseInstanceRefPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databaseInstanceRoleToPb(st *DatabaseInstanceRole) (*databaseInstanceRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databaseInstanceRolePb{}
	pb.Attributes = st.Attributes
	pb.IdentityType = st.IdentityType
	pb.MembershipRole = st.MembershipRole
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databaseInstanceRolePb struct {
	Attributes     *DatabaseInstanceRoleAttributes    `json:"attributes,omitempty"`
	IdentityType   DatabaseInstanceRoleIdentityType   `json:"identity_type,omitempty"`
	MembershipRole DatabaseInstanceRoleMembershipRole `json:"membership_role,omitempty"`
	Name           string                             `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databaseInstanceRoleFromPb(pb *databaseInstanceRolePb) (*DatabaseInstanceRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabaseInstanceRole{}
	st.Attributes = pb.Attributes
	st.IdentityType = pb.IdentityType
	st.MembershipRole = pb.MembershipRole
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databaseInstanceRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databaseInstanceRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databaseInstanceRoleAttributesToPb(st *DatabaseInstanceRoleAttributes) (*databaseInstanceRoleAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databaseInstanceRoleAttributesPb{}
	pb.Bypassrls = st.Bypassrls
	pb.Createdb = st.Createdb
	pb.Createrole = st.Createrole

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databaseInstanceRoleAttributesPb struct {
	Bypassrls  bool `json:"bypassrls,omitempty"`
	Createdb   bool `json:"createdb,omitempty"`
	Createrole bool `json:"createrole,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databaseInstanceRoleAttributesFromPb(pb *databaseInstanceRoleAttributesPb) (*DatabaseInstanceRoleAttributes, error) {
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

func (st *databaseInstanceRoleAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databaseInstanceRoleAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databaseTableToPb(st *DatabaseTable) (*databaseTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databaseTablePb{}
	pb.DatabaseInstanceName = st.DatabaseInstanceName
	pb.LogicalDatabaseName = st.LogicalDatabaseName
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databaseTablePb struct {
	DatabaseInstanceName string `json:"database_instance_name,omitempty"`
	LogicalDatabaseName  string `json:"logical_database_name,omitempty"`
	Name                 string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databaseTableFromPb(pb *databaseTablePb) (*DatabaseTable, error) {
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

func (st *databaseTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databaseTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDatabaseCatalogRequestToPb(st *DeleteDatabaseCatalogRequest) (*deleteDatabaseCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDatabaseCatalogRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteDatabaseCatalogRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteDatabaseCatalogRequestFromPb(pb *deleteDatabaseCatalogRequestPb) (*DeleteDatabaseCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseCatalogRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteDatabaseCatalogResponseToPb(st *DeleteDatabaseCatalogResponse) (*deleteDatabaseCatalogResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDatabaseCatalogResponsePb{}

	return pb, nil
}

type deleteDatabaseCatalogResponsePb struct {
}

func deleteDatabaseCatalogResponseFromPb(pb *deleteDatabaseCatalogResponsePb) (*DeleteDatabaseCatalogResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseCatalogResponse{}

	return st, nil
}

func deleteDatabaseInstanceRequestToPb(st *DeleteDatabaseInstanceRequest) (*deleteDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDatabaseInstanceRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name
	pb.Purge = st.Purge

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteDatabaseInstanceRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`
	Purge bool   `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteDatabaseInstanceRequestFromPb(pb *deleteDatabaseInstanceRequestPb) (*DeleteDatabaseInstanceRequest, error) {
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

func (st *deleteDatabaseInstanceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDatabaseInstanceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDatabaseInstanceResponseToPb(st *DeleteDatabaseInstanceResponse) (*deleteDatabaseInstanceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDatabaseInstanceResponsePb{}

	return pb, nil
}

type deleteDatabaseInstanceResponsePb struct {
}

func deleteDatabaseInstanceResponseFromPb(pb *deleteDatabaseInstanceResponsePb) (*DeleteDatabaseInstanceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseInstanceResponse{}

	return st, nil
}

func deleteDatabaseInstanceRoleRequestToPb(st *DeleteDatabaseInstanceRoleRequest) (*deleteDatabaseInstanceRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDatabaseInstanceRoleRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.InstanceName = st.InstanceName
	pb.Name = st.Name
	pb.ReassignOwnedTo = st.ReassignOwnedTo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteDatabaseInstanceRoleRequestPb struct {
	AllowMissing    bool   `json:"-" url:"allow_missing,omitempty"`
	InstanceName    string `json:"-" url:"-"`
	Name            string `json:"-" url:"-"`
	ReassignOwnedTo string `json:"-" url:"reassign_owned_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteDatabaseInstanceRoleRequestFromPb(pb *deleteDatabaseInstanceRoleRequestPb) (*DeleteDatabaseInstanceRoleRequest, error) {
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

func (st *deleteDatabaseInstanceRoleRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDatabaseInstanceRoleRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDatabaseInstanceRoleResponseToPb(st *DeleteDatabaseInstanceRoleResponse) (*deleteDatabaseInstanceRoleResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDatabaseInstanceRoleResponsePb{}

	return pb, nil
}

type deleteDatabaseInstanceRoleResponsePb struct {
}

func deleteDatabaseInstanceRoleResponseFromPb(pb *deleteDatabaseInstanceRoleResponsePb) (*DeleteDatabaseInstanceRoleResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseInstanceRoleResponse{}

	return st, nil
}

func deleteDatabaseTableRequestToPb(st *DeleteDatabaseTableRequest) (*deleteDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDatabaseTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteDatabaseTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteDatabaseTableRequestFromPb(pb *deleteDatabaseTableRequestPb) (*DeleteDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseTableRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteDatabaseTableResponseToPb(st *DeleteDatabaseTableResponse) (*deleteDatabaseTableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDatabaseTableResponsePb{}

	return pb, nil
}

type deleteDatabaseTableResponsePb struct {
}

func deleteDatabaseTableResponseFromPb(pb *deleteDatabaseTableResponsePb) (*DeleteDatabaseTableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDatabaseTableResponse{}

	return st, nil
}

func deleteSyncedDatabaseTableRequestToPb(st *DeleteSyncedDatabaseTableRequest) (*deleteSyncedDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSyncedDatabaseTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteSyncedDatabaseTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteSyncedDatabaseTableRequestFromPb(pb *deleteSyncedDatabaseTableRequestPb) (*DeleteSyncedDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSyncedDatabaseTableRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteSyncedDatabaseTableResponseToPb(st *DeleteSyncedDatabaseTableResponse) (*deleteSyncedDatabaseTableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSyncedDatabaseTableResponsePb{}

	return pb, nil
}

type deleteSyncedDatabaseTableResponsePb struct {
}

func deleteSyncedDatabaseTableResponseFromPb(pb *deleteSyncedDatabaseTableResponsePb) (*DeleteSyncedDatabaseTableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSyncedDatabaseTableResponse{}

	return st, nil
}

func deltaTableSyncInfoToPb(st *DeltaTableSyncInfo) (*deltaTableSyncInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaTableSyncInfoPb{}
	pb.DeltaCommitTimestamp = st.DeltaCommitTimestamp
	pb.DeltaCommitVersion = st.DeltaCommitVersion

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deltaTableSyncInfoPb struct {
	DeltaCommitTimestamp string `json:"delta_commit_timestamp,omitempty"`
	DeltaCommitVersion   int64  `json:"delta_commit_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaTableSyncInfoFromPb(pb *deltaTableSyncInfoPb) (*DeltaTableSyncInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaTableSyncInfo{}
	st.DeltaCommitTimestamp = pb.DeltaCommitTimestamp
	st.DeltaCommitVersion = pb.DeltaCommitVersion

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deltaTableSyncInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deltaTableSyncInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func findDatabaseInstanceByUidRequestToPb(st *FindDatabaseInstanceByUidRequest) (*findDatabaseInstanceByUidRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &findDatabaseInstanceByUidRequestPb{}
	pb.Uid = st.Uid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type findDatabaseInstanceByUidRequestPb struct {
	Uid string `json:"-" url:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func findDatabaseInstanceByUidRequestFromPb(pb *findDatabaseInstanceByUidRequestPb) (*FindDatabaseInstanceByUidRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FindDatabaseInstanceByUidRequest{}
	st.Uid = pb.Uid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *findDatabaseInstanceByUidRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st findDatabaseInstanceByUidRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func generateDatabaseCredentialRequestToPb(st *GenerateDatabaseCredentialRequest) (*generateDatabaseCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateDatabaseCredentialRequestPb{}
	pb.Claims = st.Claims
	pb.InstanceNames = st.InstanceNames
	pb.RequestId = st.RequestId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type generateDatabaseCredentialRequestPb struct {
	Claims        []RequestedClaims `json:"claims,omitempty"`
	InstanceNames []string          `json:"instance_names,omitempty"`
	RequestId     string            `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func generateDatabaseCredentialRequestFromPb(pb *generateDatabaseCredentialRequestPb) (*GenerateDatabaseCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateDatabaseCredentialRequest{}
	st.Claims = pb.Claims
	st.InstanceNames = pb.InstanceNames
	st.RequestId = pb.RequestId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *generateDatabaseCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st generateDatabaseCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getDatabaseCatalogRequestToPb(st *GetDatabaseCatalogRequest) (*getDatabaseCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDatabaseCatalogRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getDatabaseCatalogRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getDatabaseCatalogRequestFromPb(pb *getDatabaseCatalogRequestPb) (*GetDatabaseCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDatabaseCatalogRequest{}
	st.Name = pb.Name

	return st, nil
}

func getDatabaseInstanceRequestToPb(st *GetDatabaseInstanceRequest) (*getDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDatabaseInstanceRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getDatabaseInstanceRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getDatabaseInstanceRequestFromPb(pb *getDatabaseInstanceRequestPb) (*GetDatabaseInstanceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDatabaseInstanceRequest{}
	st.Name = pb.Name

	return st, nil
}

func getDatabaseInstanceRoleRequestToPb(st *GetDatabaseInstanceRoleRequest) (*getDatabaseInstanceRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDatabaseInstanceRoleRequestPb{}
	pb.InstanceName = st.InstanceName
	pb.Name = st.Name

	return pb, nil
}

type getDatabaseInstanceRoleRequestPb struct {
	InstanceName string `json:"-" url:"-"`
	Name         string `json:"-" url:"-"`
}

func getDatabaseInstanceRoleRequestFromPb(pb *getDatabaseInstanceRoleRequestPb) (*GetDatabaseInstanceRoleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDatabaseInstanceRoleRequest{}
	st.InstanceName = pb.InstanceName
	st.Name = pb.Name

	return st, nil
}

func getDatabaseTableRequestToPb(st *GetDatabaseTableRequest) (*getDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDatabaseTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getDatabaseTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getDatabaseTableRequestFromPb(pb *getDatabaseTableRequestPb) (*GetDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDatabaseTableRequest{}
	st.Name = pb.Name

	return st, nil
}

func getSyncedDatabaseTableRequestToPb(st *GetSyncedDatabaseTableRequest) (*getSyncedDatabaseTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSyncedDatabaseTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getSyncedDatabaseTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getSyncedDatabaseTableRequestFromPb(pb *getSyncedDatabaseTableRequestPb) (*GetSyncedDatabaseTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSyncedDatabaseTableRequest{}
	st.Name = pb.Name

	return st, nil
}

func listDatabaseInstanceRolesRequestToPb(st *ListDatabaseInstanceRolesRequest) (*listDatabaseInstanceRolesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDatabaseInstanceRolesRequestPb{}
	pb.InstanceName = st.InstanceName
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listDatabaseInstanceRolesRequestPb struct {
	InstanceName string `json:"-" url:"-"`
	PageSize     int    `json:"-" url:"page_size,omitempty"`
	PageToken    string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDatabaseInstanceRolesRequestFromPb(pb *listDatabaseInstanceRolesRequestPb) (*ListDatabaseInstanceRolesRequest, error) {
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

func (st *listDatabaseInstanceRolesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDatabaseInstanceRolesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listDatabaseInstanceRolesResponseToPb(st *ListDatabaseInstanceRolesResponse) (*listDatabaseInstanceRolesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDatabaseInstanceRolesResponsePb{}
	pb.DatabaseInstanceRoles = st.DatabaseInstanceRoles
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listDatabaseInstanceRolesResponsePb struct {
	DatabaseInstanceRoles []DatabaseInstanceRole `json:"database_instance_roles,omitempty"`
	NextPageToken         string                 `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDatabaseInstanceRolesResponseFromPb(pb *listDatabaseInstanceRolesResponsePb) (*ListDatabaseInstanceRolesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDatabaseInstanceRolesResponse{}
	st.DatabaseInstanceRoles = pb.DatabaseInstanceRoles
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDatabaseInstanceRolesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDatabaseInstanceRolesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listDatabaseInstancesRequestToPb(st *ListDatabaseInstancesRequest) (*listDatabaseInstancesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDatabaseInstancesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listDatabaseInstancesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDatabaseInstancesRequestFromPb(pb *listDatabaseInstancesRequestPb) (*ListDatabaseInstancesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDatabaseInstancesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDatabaseInstancesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDatabaseInstancesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listDatabaseInstancesResponseToPb(st *ListDatabaseInstancesResponse) (*listDatabaseInstancesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDatabaseInstancesResponsePb{}
	pb.DatabaseInstances = st.DatabaseInstances
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listDatabaseInstancesResponsePb struct {
	DatabaseInstances []DatabaseInstance `json:"database_instances,omitempty"`
	NextPageToken     string             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDatabaseInstancesResponseFromPb(pb *listDatabaseInstancesResponsePb) (*ListDatabaseInstancesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDatabaseInstancesResponse{}
	st.DatabaseInstances = pb.DatabaseInstances
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDatabaseInstancesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDatabaseInstancesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func newPipelineSpecToPb(st *NewPipelineSpec) (*newPipelineSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &newPipelineSpecPb{}
	pb.StorageCatalog = st.StorageCatalog
	pb.StorageSchema = st.StorageSchema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type newPipelineSpecPb struct {
	StorageCatalog string `json:"storage_catalog,omitempty"`
	StorageSchema  string `json:"storage_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func newPipelineSpecFromPb(pb *newPipelineSpecPb) (*NewPipelineSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NewPipelineSpec{}
	st.StorageCatalog = pb.StorageCatalog
	st.StorageSchema = pb.StorageSchema

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *newPipelineSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st newPipelineSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func requestedClaimsToPb(st *RequestedClaims) (*requestedClaimsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &requestedClaimsPb{}
	pb.PermissionSet = st.PermissionSet
	pb.Resources = st.Resources

	return pb, nil
}

type requestedClaimsPb struct {
	PermissionSet RequestedClaimsPermissionSet `json:"permission_set,omitempty"`
	Resources     []RequestedResource          `json:"resources,omitempty"`
}

func requestedClaimsFromPb(pb *requestedClaimsPb) (*RequestedClaims, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RequestedClaims{}
	st.PermissionSet = pb.PermissionSet
	st.Resources = pb.Resources

	return st, nil
}

func requestedResourceToPb(st *RequestedResource) (*requestedResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &requestedResourcePb{}
	pb.TableName = st.TableName
	pb.UnspecifiedResourceName = st.UnspecifiedResourceName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type requestedResourcePb struct {
	TableName               string `json:"table_name,omitempty"`
	UnspecifiedResourceName string `json:"unspecified_resource_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func requestedResourceFromPb(pb *requestedResourcePb) (*RequestedResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RequestedResource{}
	st.TableName = pb.TableName
	st.UnspecifiedResourceName = pb.UnspecifiedResourceName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *requestedResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st requestedResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func syncedDatabaseTableToPb(st *SyncedDatabaseTable) (*syncedDatabaseTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedDatabaseTablePb{}
	pb.DataSynchronizationStatus = st.DataSynchronizationStatus
	pb.DatabaseInstanceName = st.DatabaseInstanceName
	pb.LogicalDatabaseName = st.LogicalDatabaseName
	pb.Name = st.Name
	pb.Spec = st.Spec
	pb.UnityCatalogProvisioningState = st.UnityCatalogProvisioningState

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type syncedDatabaseTablePb struct {
	DataSynchronizationStatus     *SyncedTableStatus    `json:"data_synchronization_status,omitempty"`
	DatabaseInstanceName          string                `json:"database_instance_name,omitempty"`
	LogicalDatabaseName           string                `json:"logical_database_name,omitempty"`
	Name                          string                `json:"name"`
	Spec                          *SyncedTableSpec      `json:"spec,omitempty"`
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func syncedDatabaseTableFromPb(pb *syncedDatabaseTablePb) (*SyncedDatabaseTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedDatabaseTable{}
	st.DataSynchronizationStatus = pb.DataSynchronizationStatus
	st.DatabaseInstanceName = pb.DatabaseInstanceName
	st.LogicalDatabaseName = pb.LogicalDatabaseName
	st.Name = pb.Name
	st.Spec = pb.Spec
	st.UnityCatalogProvisioningState = pb.UnityCatalogProvisioningState

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *syncedDatabaseTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st syncedDatabaseTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func syncedTableContinuousUpdateStatusToPb(st *SyncedTableContinuousUpdateStatus) (*syncedTableContinuousUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedTableContinuousUpdateStatusPb{}
	pb.InitialPipelineSyncProgress = st.InitialPipelineSyncProgress
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type syncedTableContinuousUpdateStatusPb struct {
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
	LastProcessedCommitVersion  int64                        `json:"last_processed_commit_version,omitempty"`
	Timestamp                   string                       `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func syncedTableContinuousUpdateStatusFromPb(pb *syncedTableContinuousUpdateStatusPb) (*SyncedTableContinuousUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableContinuousUpdateStatus{}
	st.InitialPipelineSyncProgress = pb.InitialPipelineSyncProgress
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *syncedTableContinuousUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st syncedTableContinuousUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func syncedTableFailedStatusToPb(st *SyncedTableFailedStatus) (*syncedTableFailedStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedTableFailedStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type syncedTableFailedStatusPb struct {
	LastProcessedCommitVersion int64  `json:"last_processed_commit_version,omitempty"`
	Timestamp                  string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func syncedTableFailedStatusFromPb(pb *syncedTableFailedStatusPb) (*SyncedTableFailedStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableFailedStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *syncedTableFailedStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st syncedTableFailedStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func syncedTablePipelineProgressToPb(st *SyncedTablePipelineProgress) (*syncedTablePipelineProgressPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedTablePipelineProgressPb{}
	pb.EstimatedCompletionTimeSeconds = st.EstimatedCompletionTimeSeconds
	pb.LatestVersionCurrentlyProcessing = st.LatestVersionCurrentlyProcessing
	pb.SyncProgressCompletion = st.SyncProgressCompletion
	pb.SyncedRowCount = st.SyncedRowCount
	pb.TotalRowCount = st.TotalRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type syncedTablePipelineProgressPb struct {
	EstimatedCompletionTimeSeconds   float64 `json:"estimated_completion_time_seconds,omitempty"`
	LatestVersionCurrentlyProcessing int64   `json:"latest_version_currently_processing,omitempty"`
	SyncProgressCompletion           float64 `json:"sync_progress_completion,omitempty"`
	SyncedRowCount                   int64   `json:"synced_row_count,omitempty"`
	TotalRowCount                    int64   `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func syncedTablePipelineProgressFromPb(pb *syncedTablePipelineProgressPb) (*SyncedTablePipelineProgress, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTablePipelineProgress{}
	st.EstimatedCompletionTimeSeconds = pb.EstimatedCompletionTimeSeconds
	st.LatestVersionCurrentlyProcessing = pb.LatestVersionCurrentlyProcessing
	st.SyncProgressCompletion = pb.SyncProgressCompletion
	st.SyncedRowCount = pb.SyncedRowCount
	st.TotalRowCount = pb.TotalRowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *syncedTablePipelineProgressPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st syncedTablePipelineProgressPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func syncedTablePositionToPb(st *SyncedTablePosition) (*syncedTablePositionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedTablePositionPb{}
	pb.DeltaTableSyncInfo = st.DeltaTableSyncInfo
	pb.SyncEndTimestamp = st.SyncEndTimestamp
	pb.SyncStartTimestamp = st.SyncStartTimestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type syncedTablePositionPb struct {
	DeltaTableSyncInfo *DeltaTableSyncInfo `json:"delta_table_sync_info,omitempty"`
	SyncEndTimestamp   string              `json:"sync_end_timestamp,omitempty"`
	SyncStartTimestamp string              `json:"sync_start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func syncedTablePositionFromPb(pb *syncedTablePositionPb) (*SyncedTablePosition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTablePosition{}
	st.DeltaTableSyncInfo = pb.DeltaTableSyncInfo
	st.SyncEndTimestamp = pb.SyncEndTimestamp
	st.SyncStartTimestamp = pb.SyncStartTimestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *syncedTablePositionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st syncedTablePositionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func syncedTableProvisioningStatusToPb(st *SyncedTableProvisioningStatus) (*syncedTableProvisioningStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedTableProvisioningStatusPb{}
	pb.InitialPipelineSyncProgress = st.InitialPipelineSyncProgress

	return pb, nil
}

type syncedTableProvisioningStatusPb struct {
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

func syncedTableProvisioningStatusFromPb(pb *syncedTableProvisioningStatusPb) (*SyncedTableProvisioningStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableProvisioningStatus{}
	st.InitialPipelineSyncProgress = pb.InitialPipelineSyncProgress

	return st, nil
}

func syncedTableSpecToPb(st *SyncedTableSpec) (*syncedTableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedTableSpecPb{}
	pb.CreateDatabaseObjectsIfMissing = st.CreateDatabaseObjectsIfMissing
	pb.ExistingPipelineId = st.ExistingPipelineId
	pb.NewPipelineSpec = st.NewPipelineSpec
	pb.PrimaryKeyColumns = st.PrimaryKeyColumns
	pb.SchedulingPolicy = st.SchedulingPolicy
	pb.SourceTableFullName = st.SourceTableFullName
	pb.TimeseriesKey = st.TimeseriesKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type syncedTableSpecPb struct {
	CreateDatabaseObjectsIfMissing bool                        `json:"create_database_objects_if_missing,omitempty"`
	ExistingPipelineId             string                      `json:"existing_pipeline_id,omitempty"`
	NewPipelineSpec                *NewPipelineSpec            `json:"new_pipeline_spec,omitempty"`
	PrimaryKeyColumns              []string                    `json:"primary_key_columns,omitempty"`
	SchedulingPolicy               SyncedTableSchedulingPolicy `json:"scheduling_policy,omitempty"`
	SourceTableFullName            string                      `json:"source_table_full_name,omitempty"`
	TimeseriesKey                  string                      `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func syncedTableSpecFromPb(pb *syncedTableSpecPb) (*SyncedTableSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableSpec{}
	st.CreateDatabaseObjectsIfMissing = pb.CreateDatabaseObjectsIfMissing
	st.ExistingPipelineId = pb.ExistingPipelineId
	st.NewPipelineSpec = pb.NewPipelineSpec
	st.PrimaryKeyColumns = pb.PrimaryKeyColumns
	st.SchedulingPolicy = pb.SchedulingPolicy
	st.SourceTableFullName = pb.SourceTableFullName
	st.TimeseriesKey = pb.TimeseriesKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *syncedTableSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st syncedTableSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func syncedTableStatusToPb(st *SyncedTableStatus) (*syncedTableStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedTableStatusPb{}
	pb.ContinuousUpdateStatus = st.ContinuousUpdateStatus
	pb.DetailedState = st.DetailedState
	pb.FailedStatus = st.FailedStatus
	pb.LastSync = st.LastSync
	pb.Message = st.Message
	pb.PipelineId = st.PipelineId
	pb.ProvisioningStatus = st.ProvisioningStatus
	pb.TriggeredUpdateStatus = st.TriggeredUpdateStatus

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type syncedTableStatusPb struct {
	ContinuousUpdateStatus *SyncedTableContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	DetailedState          SyncedTableState                   `json:"detailed_state,omitempty"`
	FailedStatus           *SyncedTableFailedStatus           `json:"failed_status,omitempty"`
	LastSync               *SyncedTablePosition               `json:"last_sync,omitempty"`
	Message                string                             `json:"message,omitempty"`
	PipelineId             string                             `json:"pipeline_id,omitempty"`
	ProvisioningStatus     *SyncedTableProvisioningStatus     `json:"provisioning_status,omitempty"`
	TriggeredUpdateStatus  *SyncedTableTriggeredUpdateStatus  `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func syncedTableStatusFromPb(pb *syncedTableStatusPb) (*SyncedTableStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableStatus{}
	st.ContinuousUpdateStatus = pb.ContinuousUpdateStatus
	st.DetailedState = pb.DetailedState
	st.FailedStatus = pb.FailedStatus
	st.LastSync = pb.LastSync
	st.Message = pb.Message
	st.PipelineId = pb.PipelineId
	st.ProvisioningStatus = pb.ProvisioningStatus
	st.TriggeredUpdateStatus = pb.TriggeredUpdateStatus

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *syncedTableStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st syncedTableStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func syncedTableTriggeredUpdateStatusToPb(st *SyncedTableTriggeredUpdateStatus) (*syncedTableTriggeredUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncedTableTriggeredUpdateStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	pb.Timestamp = st.Timestamp
	pb.TriggeredUpdateProgress = st.TriggeredUpdateProgress

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type syncedTableTriggeredUpdateStatusPb struct {
	LastProcessedCommitVersion int64                        `json:"last_processed_commit_version,omitempty"`
	Timestamp                  string                       `json:"timestamp,omitempty"`
	TriggeredUpdateProgress    *SyncedTablePipelineProgress `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func syncedTableTriggeredUpdateStatusFromPb(pb *syncedTableTriggeredUpdateStatusPb) (*SyncedTableTriggeredUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncedTableTriggeredUpdateStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp
	st.TriggeredUpdateProgress = pb.TriggeredUpdateProgress

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *syncedTableTriggeredUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st syncedTableTriggeredUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateDatabaseInstanceRequestToPb(st *UpdateDatabaseInstanceRequest) (*updateDatabaseInstanceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDatabaseInstanceRequestPb{}
	pb.DatabaseInstance = st.DatabaseInstance
	pb.Name = st.Name
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

type updateDatabaseInstanceRequestPb struct {
	DatabaseInstance DatabaseInstance `json:"database_instance"`
	Name             string           `json:"-" url:"-"`
	UpdateMask       string           `json:"-" url:"update_mask"`
}

func updateDatabaseInstanceRequestFromPb(pb *updateDatabaseInstanceRequestPb) (*UpdateDatabaseInstanceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDatabaseInstanceRequest{}
	st.DatabaseInstance = pb.DatabaseInstance
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
