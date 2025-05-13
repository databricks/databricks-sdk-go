// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
)

type CleanRoom struct {
	// Whether clean room access is restricted due to [CSP]
	//
	// [CSP]: https://docs.databricks.com/en/security/privacy/security-profile.html
	// Wire name: 'access_restricted'
	AccessRestricted CleanRoomAccessRestricted

	// Wire name: 'comment'
	Comment string
	// When the clean room was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// The alias of the collaborator tied to the local clean room.
	// Wire name: 'local_collaborator_alias'
	LocalCollaboratorAlias string
	// The name of the clean room. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	// Wire name: 'name'
	Name string
	// Output catalog of the clean room. It is an output only field. Output
	// catalog is manipulated using the separate CreateCleanRoomOutputCatalog
	// API.
	// Wire name: 'output_catalog'
	OutputCatalog *CleanRoomOutputCatalog
	// This is Databricks username of the owner of the local clean room
	// securable for permission management.
	// Wire name: 'owner'
	Owner string
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	// Wire name: 'remote_detailed_info'
	RemoteDetailedInfo *CleanRoomRemoteDetail
	// Clean room status.
	// Wire name: 'status'
	Status CleanRoomStatusEnum
	// When the clean room was last updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64

	ForceSendFields []string `tf:"-"`
}

func cleanRoomToPb(st *CleanRoom) (*cleanRoomPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomPb{}
	pb.AccessRestricted = st.AccessRestricted

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.LocalCollaboratorAlias = st.LocalCollaboratorAlias

	pb.Name = st.Name

	outputCatalogPb, err := cleanRoomOutputCatalogToPb(st.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogPb != nil {
		pb.OutputCatalog = outputCatalogPb
	}

	pb.Owner = st.Owner

	remoteDetailedInfoPb, err := cleanRoomRemoteDetailToPb(st.RemoteDetailedInfo)
	if err != nil {
		return nil, err
	}
	if remoteDetailedInfoPb != nil {
		pb.RemoteDetailedInfo = remoteDetailedInfoPb
	}

	pb.Status = st.Status

	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoom) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoom) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomPb struct {
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
	OutputCatalog *cleanRoomOutputCatalogPb `json:"output_catalog,omitempty"`
	// This is Databricks username of the owner of the local clean room
	// securable for permission management.
	Owner string `json:"owner,omitempty"`
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	RemoteDetailedInfo *cleanRoomRemoteDetailPb `json:"remote_detailed_info,omitempty"`
	// Clean room status.
	Status CleanRoomStatusEnum `json:"status,omitempty"`
	// When the clean room was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomFromPb(pb *cleanRoomPb) (*CleanRoom, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoom{}
	st.AccessRestricted = pb.AccessRestricted
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.LocalCollaboratorAlias = pb.LocalCollaboratorAlias
	st.Name = pb.Name
	outputCatalogField, err := cleanRoomOutputCatalogFromPb(pb.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogField != nil {
		st.OutputCatalog = outputCatalogField
	}
	st.Owner = pb.Owner
	remoteDetailedInfoField, err := cleanRoomRemoteDetailFromPb(pb.RemoteDetailedInfo)
	if err != nil {
		return nil, err
	}
	if remoteDetailedInfoField != nil {
		st.RemoteDetailedInfo = remoteDetailedInfoField
	}
	st.Status = pb.Status
	st.UpdatedAt = pb.UpdatedAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAccessRestricted string
type cleanRoomAccessRestrictedPb string

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

func cleanRoomAccessRestrictedToPb(st *CleanRoomAccessRestricted) (*cleanRoomAccessRestrictedPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanRoomAccessRestrictedPb(*st)
	return &pb, nil
}

func cleanRoomAccessRestrictedFromPb(pb *cleanRoomAccessRestrictedPb) (*CleanRoomAccessRestricted, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomAccessRestricted(*pb)
	return &st, nil
}

// Metadata of the clean room asset
type CleanRoomAsset struct {
	// When the asset is added to the clean room, in epoch milliseconds.
	// Wire name: 'added_at'
	AddedAt int64
	// The type of the asset.
	// Wire name: 'asset_type'
	AssetType CleanRoomAssetAssetType
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	// Wire name: 'foreign_table'
	ForeignTable *CleanRoomAssetForeignTable
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	// Wire name: 'foreign_table_local_details'
	ForeignTableLocalDetails *CleanRoomAssetForeignTableLocalDetails
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	// Wire name: 'name'
	Name string
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	// Wire name: 'notebook'
	Notebook *CleanRoomAssetNotebook
	// The alias of the collaborator who owns this asset
	// Wire name: 'owner_collaborator_alias'
	OwnerCollaboratorAlias string
	// Status of the asset
	// Wire name: 'status'
	Status CleanRoomAssetStatusEnum
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	// Wire name: 'table'
	Table *CleanRoomAssetTable
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	// Wire name: 'table_local_details'
	TableLocalDetails *CleanRoomAssetTableLocalDetails
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	// Wire name: 'view'
	View *CleanRoomAssetView
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	// Wire name: 'view_local_details'
	ViewLocalDetails *CleanRoomAssetViewLocalDetails
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	// Wire name: 'volume_local_details'
	VolumeLocalDetails *CleanRoomAssetVolumeLocalDetails

	ForceSendFields []string `tf:"-"`
}

func cleanRoomAssetToPb(st *CleanRoomAsset) (*cleanRoomAssetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetPb{}
	pb.AddedAt = st.AddedAt

	pb.AssetType = st.AssetType

	foreignTablePb, err := cleanRoomAssetForeignTableToPb(st.ForeignTable)
	if err != nil {
		return nil, err
	}
	if foreignTablePb != nil {
		pb.ForeignTable = foreignTablePb
	}

	foreignTableLocalDetailsPb, err := cleanRoomAssetForeignTableLocalDetailsToPb(st.ForeignTableLocalDetails)
	if err != nil {
		return nil, err
	}
	if foreignTableLocalDetailsPb != nil {
		pb.ForeignTableLocalDetails = foreignTableLocalDetailsPb
	}

	pb.Name = st.Name

	notebookPb, err := cleanRoomAssetNotebookToPb(st.Notebook)
	if err != nil {
		return nil, err
	}
	if notebookPb != nil {
		pb.Notebook = notebookPb
	}

	pb.OwnerCollaboratorAlias = st.OwnerCollaboratorAlias

	pb.Status = st.Status

	tablePb, err := cleanRoomAssetTableToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	tableLocalDetailsPb, err := cleanRoomAssetTableLocalDetailsToPb(st.TableLocalDetails)
	if err != nil {
		return nil, err
	}
	if tableLocalDetailsPb != nil {
		pb.TableLocalDetails = tableLocalDetailsPb
	}

	viewPb, err := cleanRoomAssetViewToPb(st.View)
	if err != nil {
		return nil, err
	}
	if viewPb != nil {
		pb.View = viewPb
	}

	viewLocalDetailsPb, err := cleanRoomAssetViewLocalDetailsToPb(st.ViewLocalDetails)
	if err != nil {
		return nil, err
	}
	if viewLocalDetailsPb != nil {
		pb.ViewLocalDetails = viewLocalDetailsPb
	}

	volumeLocalDetailsPb, err := cleanRoomAssetVolumeLocalDetailsToPb(st.VolumeLocalDetails)
	if err != nil {
		return nil, err
	}
	if volumeLocalDetailsPb != nil {
		pb.VolumeLocalDetails = volumeLocalDetailsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomAsset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAsset) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetPb struct {
	// When the asset is added to the clean room, in epoch milliseconds.
	AddedAt int64 `json:"added_at,omitempty"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"asset_type,omitempty"`
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTable *cleanRoomAssetForeignTablePb `json:"foreign_table,omitempty"`
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTableLocalDetails *cleanRoomAssetForeignTableLocalDetailsPb `json:"foreign_table_local_details,omitempty"`
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
	Notebook *cleanRoomAssetNotebookPb `json:"notebook,omitempty"`
	// The alias of the collaborator who owns this asset
	OwnerCollaboratorAlias string `json:"owner_collaborator_alias,omitempty"`
	// Status of the asset
	Status CleanRoomAssetStatusEnum `json:"status,omitempty"`
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	Table *cleanRoomAssetTablePb `json:"table,omitempty"`
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	TableLocalDetails *cleanRoomAssetTableLocalDetailsPb `json:"table_local_details,omitempty"`
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	View *cleanRoomAssetViewPb `json:"view,omitempty"`
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	ViewLocalDetails *cleanRoomAssetViewLocalDetailsPb `json:"view_local_details,omitempty"`
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	VolumeLocalDetails *cleanRoomAssetVolumeLocalDetailsPb `json:"volume_local_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetFromPb(pb *cleanRoomAssetPb) (*CleanRoomAsset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAsset{}
	st.AddedAt = pb.AddedAt
	st.AssetType = pb.AssetType
	foreignTableField, err := cleanRoomAssetForeignTableFromPb(pb.ForeignTable)
	if err != nil {
		return nil, err
	}
	if foreignTableField != nil {
		st.ForeignTable = foreignTableField
	}
	foreignTableLocalDetailsField, err := cleanRoomAssetForeignTableLocalDetailsFromPb(pb.ForeignTableLocalDetails)
	if err != nil {
		return nil, err
	}
	if foreignTableLocalDetailsField != nil {
		st.ForeignTableLocalDetails = foreignTableLocalDetailsField
	}
	st.Name = pb.Name
	notebookField, err := cleanRoomAssetNotebookFromPb(pb.Notebook)
	if err != nil {
		return nil, err
	}
	if notebookField != nil {
		st.Notebook = notebookField
	}
	st.OwnerCollaboratorAlias = pb.OwnerCollaboratorAlias
	st.Status = pb.Status
	tableField, err := cleanRoomAssetTableFromPb(pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = tableField
	}
	tableLocalDetailsField, err := cleanRoomAssetTableLocalDetailsFromPb(pb.TableLocalDetails)
	if err != nil {
		return nil, err
	}
	if tableLocalDetailsField != nil {
		st.TableLocalDetails = tableLocalDetailsField
	}
	viewField, err := cleanRoomAssetViewFromPb(pb.View)
	if err != nil {
		return nil, err
	}
	if viewField != nil {
		st.View = viewField
	}
	viewLocalDetailsField, err := cleanRoomAssetViewLocalDetailsFromPb(pb.ViewLocalDetails)
	if err != nil {
		return nil, err
	}
	if viewLocalDetailsField != nil {
		st.ViewLocalDetails = viewLocalDetailsField
	}
	volumeLocalDetailsField, err := cleanRoomAssetVolumeLocalDetailsFromPb(pb.VolumeLocalDetails)
	if err != nil {
		return nil, err
	}
	if volumeLocalDetailsField != nil {
		st.VolumeLocalDetails = volumeLocalDetailsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAssetAssetType string
type cleanRoomAssetAssetTypePb string

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

func cleanRoomAssetAssetTypeToPb(st *CleanRoomAssetAssetType) (*cleanRoomAssetAssetTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanRoomAssetAssetTypePb(*st)
	return &pb, nil
}

func cleanRoomAssetAssetTypeFromPb(pb *cleanRoomAssetAssetTypePb) (*CleanRoomAssetAssetType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomAssetAssetType(*pb)
	return &st, nil
}

type CleanRoomAssetForeignTable struct {
	// The metadata information of the columns in the foreign table
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo
}

func cleanRoomAssetForeignTableToPb(st *CleanRoomAssetForeignTable) (*cleanRoomAssetForeignTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetForeignTablePb{}

	var columnsPb []catalog.ColumnInfoPb
	for _, item := range st.Columns {
		itemPb, err := catalog.ColumnInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb

	return pb, nil
}

func (st *CleanRoomAssetForeignTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetForeignTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetForeignTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAssetForeignTable) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetForeignTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetForeignTablePb struct {
	// The metadata information of the columns in the foreign table
	Columns []catalog.ColumnInfoPb `json:"columns,omitempty"`
}

func cleanRoomAssetForeignTableFromPb(pb *cleanRoomAssetForeignTablePb) (*CleanRoomAssetForeignTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetForeignTable{}

	var columnsField []catalog.ColumnInfo
	for _, itemPb := range pb.Columns {
		item, err := catalog.ColumnInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField

	return st, nil
}

type CleanRoomAssetForeignTableLocalDetails struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	// Wire name: 'local_name'
	LocalName string

	ForceSendFields []string `tf:"-"`
}

func cleanRoomAssetForeignTableLocalDetailsToPb(st *CleanRoomAssetForeignTableLocalDetails) (*cleanRoomAssetForeignTableLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetForeignTableLocalDetailsPb{}
	pb.LocalName = st.LocalName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomAssetForeignTableLocalDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetForeignTableLocalDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetForeignTableLocalDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAssetForeignTableLocalDetails) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetForeignTableLocalDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetForeignTableLocalDetailsPb struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetForeignTableLocalDetailsFromPb(pb *cleanRoomAssetForeignTableLocalDetailsPb) (*CleanRoomAssetForeignTableLocalDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetForeignTableLocalDetails{}
	st.LocalName = pb.LocalName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetForeignTableLocalDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetForeignTableLocalDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAssetNotebook struct {
	// Server generated etag that represents the notebook version.
	// Wire name: 'etag'
	Etag string
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	// Wire name: 'notebook_content'
	NotebookContent string

	ForceSendFields []string `tf:"-"`
}

func cleanRoomAssetNotebookToPb(st *CleanRoomAssetNotebook) (*cleanRoomAssetNotebookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetNotebookPb{}
	pb.Etag = st.Etag

	pb.NotebookContent = st.NotebookContent

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomAssetNotebook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetNotebookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetNotebookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAssetNotebook) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetNotebookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetNotebookPb struct {
	// Server generated etag that represents the notebook version.
	Etag string `json:"etag,omitempty"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	NotebookContent string `json:"notebook_content,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetNotebookFromPb(pb *cleanRoomAssetNotebookPb) (*CleanRoomAssetNotebook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetNotebook{}
	st.Etag = pb.Etag
	st.NotebookContent = pb.NotebookContent

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetNotebookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetNotebookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAssetStatusEnum string
type cleanRoomAssetStatusEnumPb string

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

func cleanRoomAssetStatusEnumToPb(st *CleanRoomAssetStatusEnum) (*cleanRoomAssetStatusEnumPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanRoomAssetStatusEnumPb(*st)
	return &pb, nil
}

func cleanRoomAssetStatusEnumFromPb(pb *cleanRoomAssetStatusEnumPb) (*CleanRoomAssetStatusEnum, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomAssetStatusEnum(*pb)
	return &st, nil
}

type CleanRoomAssetTable struct {
	// The metadata information of the columns in the table
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo
}

func cleanRoomAssetTableToPb(st *CleanRoomAssetTable) (*cleanRoomAssetTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetTablePb{}

	var columnsPb []catalog.ColumnInfoPb
	for _, item := range st.Columns {
		itemPb, err := catalog.ColumnInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb

	return pb, nil
}

func (st *CleanRoomAssetTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAssetTable) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetTablePb struct {
	// The metadata information of the columns in the table
	Columns []catalog.ColumnInfoPb `json:"columns,omitempty"`
}

func cleanRoomAssetTableFromPb(pb *cleanRoomAssetTablePb) (*CleanRoomAssetTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetTable{}

	var columnsField []catalog.ColumnInfo
	for _, itemPb := range pb.Columns {
		item, err := catalog.ColumnInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField

	return st, nil
}

type CleanRoomAssetTableLocalDetails struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	// Wire name: 'local_name'
	LocalName string
	// Partition filtering specification for a shared table.
	// Wire name: 'partitions'
	Partitions []sharing.Partition

	ForceSendFields []string `tf:"-"`
}

func cleanRoomAssetTableLocalDetailsToPb(st *CleanRoomAssetTableLocalDetails) (*cleanRoomAssetTableLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetTableLocalDetailsPb{}
	pb.LocalName = st.LocalName

	var partitionsPb []sharing.PartitionPb
	for _, item := range st.Partitions {
		itemPb, err := sharing.PartitionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			partitionsPb = append(partitionsPb, *itemPb)
		}
	}
	pb.Partitions = partitionsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomAssetTableLocalDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetTableLocalDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetTableLocalDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAssetTableLocalDetails) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetTableLocalDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetTableLocalDetailsPb struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	LocalName string `json:"local_name,omitempty"`
	// Partition filtering specification for a shared table.
	Partitions []sharing.PartitionPb `json:"partitions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetTableLocalDetailsFromPb(pb *cleanRoomAssetTableLocalDetailsPb) (*CleanRoomAssetTableLocalDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetTableLocalDetails{}
	st.LocalName = pb.LocalName

	var partitionsField []sharing.Partition
	for _, itemPb := range pb.Partitions {
		item, err := sharing.PartitionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			partitionsField = append(partitionsField, *item)
		}
	}
	st.Partitions = partitionsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetTableLocalDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetTableLocalDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAssetView struct {
	// The metadata information of the columns in the view
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo
}

func cleanRoomAssetViewToPb(st *CleanRoomAssetView) (*cleanRoomAssetViewPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetViewPb{}

	var columnsPb []catalog.ColumnInfoPb
	for _, item := range st.Columns {
		itemPb, err := catalog.ColumnInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb

	return pb, nil
}

func (st *CleanRoomAssetView) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetViewPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetViewFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAssetView) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetViewToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetViewPb struct {
	// The metadata information of the columns in the view
	Columns []catalog.ColumnInfoPb `json:"columns,omitempty"`
}

func cleanRoomAssetViewFromPb(pb *cleanRoomAssetViewPb) (*CleanRoomAssetView, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetView{}

	var columnsField []catalog.ColumnInfo
	for _, itemPb := range pb.Columns {
		item, err := catalog.ColumnInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField

	return st, nil
}

type CleanRoomAssetViewLocalDetails struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	// Wire name: 'local_name'
	LocalName string

	ForceSendFields []string `tf:"-"`
}

func cleanRoomAssetViewLocalDetailsToPb(st *CleanRoomAssetViewLocalDetails) (*cleanRoomAssetViewLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetViewLocalDetailsPb{}
	pb.LocalName = st.LocalName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomAssetViewLocalDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetViewLocalDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetViewLocalDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAssetViewLocalDetails) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetViewLocalDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetViewLocalDetailsPb struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetViewLocalDetailsFromPb(pb *cleanRoomAssetViewLocalDetailsPb) (*CleanRoomAssetViewLocalDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetViewLocalDetails{}
	st.LocalName = pb.LocalName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetViewLocalDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetViewLocalDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAssetVolumeLocalDetails struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	// Wire name: 'local_name'
	LocalName string

	ForceSendFields []string `tf:"-"`
}

func cleanRoomAssetVolumeLocalDetailsToPb(st *CleanRoomAssetVolumeLocalDetails) (*cleanRoomAssetVolumeLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetVolumeLocalDetailsPb{}
	pb.LocalName = st.LocalName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomAssetVolumeLocalDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomAssetVolumeLocalDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomAssetVolumeLocalDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomAssetVolumeLocalDetails) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomAssetVolumeLocalDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomAssetVolumeLocalDetailsPb struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetVolumeLocalDetailsFromPb(pb *cleanRoomAssetVolumeLocalDetailsPb) (*CleanRoomAssetVolumeLocalDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetVolumeLocalDetails{}
	st.LocalName = pb.LocalName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetVolumeLocalDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetVolumeLocalDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	// Wire name: 'collaborator_alias'
	CollaboratorAlias string
	// Generated display name for the collaborator. In the case of a single
	// metastore clean room, it is the clean room name. For x-metastore clean
	// rooms, it is the organization name of the metastore. It is not restricted
	// to these values and could change in the future
	// Wire name: 'display_name'
	DisplayName string
	// The global Unity Catalog metastore id of the collaborator. The identifier
	// is of format cloud:region:metastore-uuid.
	// Wire name: 'global_metastore_id'
	GlobalMetastoreId string
	// Email of the user who is receiving the clean room "invitation". It should
	// be empty for the creator of the clean room, and non-empty for the
	// invitees of the clean room. It is only returned in the output when clean
	// room creator calls GET
	// Wire name: 'invite_recipient_email'
	InviteRecipientEmail string
	// Workspace ID of the user who is receiving the clean room "invitation".
	// Must be specified if invite_recipient_email is specified. It should be
	// empty when the collaborator is the creator of the clean room.
	// Wire name: 'invite_recipient_workspace_id'
	InviteRecipientWorkspaceId int64
	// [Organization
	// name](:method:metastores/list#metastores-delta_sharing_organization_name)
	// configured in the metastore
	// Wire name: 'organization_name'
	OrganizationName string

	ForceSendFields []string `tf:"-"`
}

func cleanRoomCollaboratorToPb(st *CleanRoomCollaborator) (*cleanRoomCollaboratorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomCollaboratorPb{}
	pb.CollaboratorAlias = st.CollaboratorAlias

	pb.DisplayName = st.DisplayName

	pb.GlobalMetastoreId = st.GlobalMetastoreId

	pb.InviteRecipientEmail = st.InviteRecipientEmail

	pb.InviteRecipientWorkspaceId = st.InviteRecipientWorkspaceId

	pb.OrganizationName = st.OrganizationName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomCollaborator) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomCollaboratorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomCollaboratorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomCollaborator) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomCollaboratorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomCollaboratorPb struct {
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

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomCollaboratorFromPb(pb *cleanRoomCollaboratorPb) (*CleanRoomCollaborator, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomCollaborator{}
	st.CollaboratorAlias = pb.CollaboratorAlias
	st.DisplayName = pb.DisplayName
	st.GlobalMetastoreId = pb.GlobalMetastoreId
	st.InviteRecipientEmail = pb.InviteRecipientEmail
	st.InviteRecipientWorkspaceId = pb.InviteRecipientWorkspaceId
	st.OrganizationName = pb.OrganizationName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomCollaboratorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomCollaboratorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Stores information about a single task run.
type CleanRoomNotebookTaskRun struct {
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API. if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	// Wire name: 'collaborator_job_run_info'
	CollaboratorJobRunInfo *CollaboratorJobRunInfo
	// State of the task run.
	// Wire name: 'notebook_job_run_state'
	NotebookJobRunState *jobs.CleanRoomTaskRunState
	// Asset name of the notebook executed in this task run.
	// Wire name: 'notebook_name'
	NotebookName string
	// Expiration time of the output schema of the task run (if any), in epoch
	// milliseconds.
	// Wire name: 'output_schema_expiration_time'
	OutputSchemaExpirationTime int64
	// Name of the output schema associated with the clean rooms notebook task
	// run.
	// Wire name: 'output_schema_name'
	OutputSchemaName string
	// Duration of the task run, in milliseconds.
	// Wire name: 'run_duration'
	RunDuration int64
	// When the task run started, in epoch milliseconds.
	// Wire name: 'start_time'
	StartTime int64

	ForceSendFields []string `tf:"-"`
}

func cleanRoomNotebookTaskRunToPb(st *CleanRoomNotebookTaskRun) (*cleanRoomNotebookTaskRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomNotebookTaskRunPb{}
	collaboratorJobRunInfoPb, err := collaboratorJobRunInfoToPb(st.CollaboratorJobRunInfo)
	if err != nil {
		return nil, err
	}
	if collaboratorJobRunInfoPb != nil {
		pb.CollaboratorJobRunInfo = collaboratorJobRunInfoPb
	}

	notebookJobRunStatePb, err := jobs.CleanRoomTaskRunStateToPb(st.NotebookJobRunState)
	if err != nil {
		return nil, err
	}
	if notebookJobRunStatePb != nil {
		pb.NotebookJobRunState = notebookJobRunStatePb
	}

	pb.NotebookName = st.NotebookName

	pb.OutputSchemaExpirationTime = st.OutputSchemaExpirationTime

	pb.OutputSchemaName = st.OutputSchemaName

	pb.RunDuration = st.RunDuration

	pb.StartTime = st.StartTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomNotebookTaskRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomNotebookTaskRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomNotebookTaskRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomNotebookTaskRun) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomNotebookTaskRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomNotebookTaskRunPb struct {
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API. if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	CollaboratorJobRunInfo *collaboratorJobRunInfoPb `json:"collaborator_job_run_info,omitempty"`
	// State of the task run.
	NotebookJobRunState *jobs.CleanRoomTaskRunStatePb `json:"notebook_job_run_state,omitempty"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomNotebookTaskRunFromPb(pb *cleanRoomNotebookTaskRunPb) (*CleanRoomNotebookTaskRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomNotebookTaskRun{}
	collaboratorJobRunInfoField, err := collaboratorJobRunInfoFromPb(pb.CollaboratorJobRunInfo)
	if err != nil {
		return nil, err
	}
	if collaboratorJobRunInfoField != nil {
		st.CollaboratorJobRunInfo = collaboratorJobRunInfoField
	}
	notebookJobRunStateField, err := jobs.CleanRoomTaskRunStateFromPb(pb.NotebookJobRunState)
	if err != nil {
		return nil, err
	}
	if notebookJobRunStateField != nil {
		st.NotebookJobRunState = notebookJobRunStateField
	}
	st.NotebookName = pb.NotebookName
	st.OutputSchemaExpirationTime = pb.OutputSchemaExpirationTime
	st.OutputSchemaName = pb.OutputSchemaName
	st.RunDuration = pb.RunDuration
	st.StartTime = pb.StartTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomNotebookTaskRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomNotebookTaskRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	// Wire name: 'catalog_name'
	CatalogName string

	// Wire name: 'status'
	Status CleanRoomOutputCatalogOutputCatalogStatus

	ForceSendFields []string `tf:"-"`
}

func cleanRoomOutputCatalogToPb(st *CleanRoomOutputCatalog) (*cleanRoomOutputCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomOutputCatalogPb{}
	pb.CatalogName = st.CatalogName

	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomOutputCatalog) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomOutputCatalogPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomOutputCatalogFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomOutputCatalog) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomOutputCatalogToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomOutputCatalogPb struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName string `json:"catalog_name,omitempty"`

	Status CleanRoomOutputCatalogOutputCatalogStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomOutputCatalogFromPb(pb *cleanRoomOutputCatalogPb) (*CleanRoomOutputCatalog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomOutputCatalog{}
	st.CatalogName = pb.CatalogName
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomOutputCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomOutputCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomOutputCatalogOutputCatalogStatus string
type cleanRoomOutputCatalogOutputCatalogStatusPb string

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

func cleanRoomOutputCatalogOutputCatalogStatusToPb(st *CleanRoomOutputCatalogOutputCatalogStatus) (*cleanRoomOutputCatalogOutputCatalogStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanRoomOutputCatalogOutputCatalogStatusPb(*st)
	return &pb, nil
}

func cleanRoomOutputCatalogOutputCatalogStatusFromPb(pb *cleanRoomOutputCatalogOutputCatalogStatusPb) (*CleanRoomOutputCatalogOutputCatalogStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomOutputCatalogOutputCatalogStatus(*pb)
	return &st, nil
}

// Publicly visible central clean room details.
type CleanRoomRemoteDetail struct {
	// Central clean room ID.
	// Wire name: 'central_clean_room_id'
	CentralCleanRoomId string
	// Cloud vendor (aws,azure,gcp) of the central clean room.
	// Wire name: 'cloud_vendor'
	CloudVendor string
	// Collaborators in the central clean room. There should one and only one
	// collaborator in the list that satisfies the owner condition:
	//
	// 1. It has the creator's global_metastore_id (determined by caller of
	// CreateCleanRoom).
	//
	// 2. Its invite_recipient_email is empty.
	// Wire name: 'collaborators'
	Collaborators []CleanRoomCollaborator
	// The compliance security profile used to process regulated data following
	// compliance standards.
	// Wire name: 'compliance_security_profile'
	ComplianceSecurityProfile *ComplianceSecurityProfile
	// Collaborator who creates the clean room.
	// Wire name: 'creator'
	Creator *CleanRoomCollaborator
	// Egress network policy to apply to the central clean room workspace.
	// Wire name: 'egress_network_policy'
	EgressNetworkPolicy *settings.EgressNetworkPolicy
	// Region of the central clean room.
	// Wire name: 'region'
	Region string

	ForceSendFields []string `tf:"-"`
}

func cleanRoomRemoteDetailToPb(st *CleanRoomRemoteDetail) (*cleanRoomRemoteDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomRemoteDetailPb{}
	pb.CentralCleanRoomId = st.CentralCleanRoomId

	pb.CloudVendor = st.CloudVendor

	var collaboratorsPb []cleanRoomCollaboratorPb
	for _, item := range st.Collaborators {
		itemPb, err := cleanRoomCollaboratorToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			collaboratorsPb = append(collaboratorsPb, *itemPb)
		}
	}
	pb.Collaborators = collaboratorsPb

	complianceSecurityProfilePb, err := complianceSecurityProfileToPb(st.ComplianceSecurityProfile)
	if err != nil {
		return nil, err
	}
	if complianceSecurityProfilePb != nil {
		pb.ComplianceSecurityProfile = complianceSecurityProfilePb
	}

	creatorPb, err := cleanRoomCollaboratorToPb(st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = creatorPb
	}

	egressNetworkPolicyPb, err := settings.EgressNetworkPolicyToPb(st.EgressNetworkPolicy)
	if err != nil {
		return nil, err
	}
	if egressNetworkPolicyPb != nil {
		pb.EgressNetworkPolicy = egressNetworkPolicyPb
	}

	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomRemoteDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomRemoteDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomRemoteDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomRemoteDetail) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomRemoteDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomRemoteDetailPb struct {
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
	Collaborators []cleanRoomCollaboratorPb `json:"collaborators,omitempty"`
	// The compliance security profile used to process regulated data following
	// compliance standards.
	ComplianceSecurityProfile *complianceSecurityProfilePb `json:"compliance_security_profile,omitempty"`
	// Collaborator who creates the clean room.
	Creator *cleanRoomCollaboratorPb `json:"creator,omitempty"`
	// Egress network policy to apply to the central clean room workspace.
	EgressNetworkPolicy *settings.EgressNetworkPolicyPb `json:"egress_network_policy,omitempty"`
	// Region of the central clean room.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomRemoteDetailFromPb(pb *cleanRoomRemoteDetailPb) (*CleanRoomRemoteDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomRemoteDetail{}
	st.CentralCleanRoomId = pb.CentralCleanRoomId
	st.CloudVendor = pb.CloudVendor

	var collaboratorsField []CleanRoomCollaborator
	for _, itemPb := range pb.Collaborators {
		item, err := cleanRoomCollaboratorFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			collaboratorsField = append(collaboratorsField, *item)
		}
	}
	st.Collaborators = collaboratorsField
	complianceSecurityProfileField, err := complianceSecurityProfileFromPb(pb.ComplianceSecurityProfile)
	if err != nil {
		return nil, err
	}
	if complianceSecurityProfileField != nil {
		st.ComplianceSecurityProfile = complianceSecurityProfileField
	}
	creatorField, err := cleanRoomCollaboratorFromPb(pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = creatorField
	}
	egressNetworkPolicyField, err := settings.EgressNetworkPolicyFromPb(pb.EgressNetworkPolicy)
	if err != nil {
		return nil, err
	}
	if egressNetworkPolicyField != nil {
		st.EgressNetworkPolicy = egressNetworkPolicyField
	}
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomRemoteDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomRemoteDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomStatusEnum string
type cleanRoomStatusEnumPb string

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

func cleanRoomStatusEnumToPb(st *CleanRoomStatusEnum) (*cleanRoomStatusEnumPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanRoomStatusEnumPb(*st)
	return &pb, nil
}

func cleanRoomStatusEnumFromPb(pb *cleanRoomStatusEnumPb) (*CleanRoomStatusEnum, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomStatusEnum(*pb)
	return &st, nil
}

type CollaboratorJobRunInfo struct {
	// Alias of the collaborator that triggered the task run.
	// Wire name: 'collaborator_alias'
	CollaboratorAlias string
	// Job ID of the task run in the collaborator's workspace.
	// Wire name: 'collaborator_job_id'
	CollaboratorJobId int64
	// Job run ID of the task run in the collaborator's workspace.
	// Wire name: 'collaborator_job_run_id'
	CollaboratorJobRunId int64
	// Task run ID of the task run in the collaborator's workspace.
	// Wire name: 'collaborator_task_run_id'
	CollaboratorTaskRunId int64
	// ID of the collaborator's workspace that triggered the task run.
	// Wire name: 'collaborator_workspace_id'
	CollaboratorWorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func collaboratorJobRunInfoToPb(st *CollaboratorJobRunInfo) (*collaboratorJobRunInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &collaboratorJobRunInfoPb{}
	pb.CollaboratorAlias = st.CollaboratorAlias

	pb.CollaboratorJobId = st.CollaboratorJobId

	pb.CollaboratorJobRunId = st.CollaboratorJobRunId

	pb.CollaboratorTaskRunId = st.CollaboratorTaskRunId

	pb.CollaboratorWorkspaceId = st.CollaboratorWorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CollaboratorJobRunInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &collaboratorJobRunInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := collaboratorJobRunInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CollaboratorJobRunInfo) MarshalJSON() ([]byte, error) {
	pb, err := collaboratorJobRunInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type collaboratorJobRunInfoPb struct {
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

	ForceSendFields []string `json:"-" url:"-"`
}

func collaboratorJobRunInfoFromPb(pb *collaboratorJobRunInfoPb) (*CollaboratorJobRunInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CollaboratorJobRunInfo{}
	st.CollaboratorAlias = pb.CollaboratorAlias
	st.CollaboratorJobId = pb.CollaboratorJobId
	st.CollaboratorJobRunId = pb.CollaboratorJobRunId
	st.CollaboratorTaskRunId = pb.CollaboratorTaskRunId
	st.CollaboratorWorkspaceId = pb.CollaboratorWorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *collaboratorJobRunInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st collaboratorJobRunInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The compliance security profile used to process regulated data following
// compliance standards.
type ComplianceSecurityProfile struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	// Wire name: 'compliance_standards'
	ComplianceStandards []settings.ComplianceStandard
	// Whether the compliance security profile is enabled.
	// Wire name: 'is_enabled'
	IsEnabled bool

	ForceSendFields []string `tf:"-"`
}

func complianceSecurityProfileToPb(st *ComplianceSecurityProfile) (*complianceSecurityProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &complianceSecurityProfilePb{}

	var complianceStandardsPb []settings.ComplianceStandardPb
	for _, item := range st.ComplianceStandards {
		itemPb, err := settings.ComplianceStandardToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			complianceStandardsPb = append(complianceStandardsPb, *itemPb)
		}
	}
	pb.ComplianceStandards = complianceStandardsPb

	pb.IsEnabled = st.IsEnabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &complianceSecurityProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := complianceSecurityProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	pb, err := complianceSecurityProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type complianceSecurityProfilePb struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	ComplianceStandards []settings.ComplianceStandardPb `json:"compliance_standards,omitempty"`
	// Whether the compliance security profile is enabled.
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func complianceSecurityProfileFromPb(pb *complianceSecurityProfilePb) (*ComplianceSecurityProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplianceSecurityProfile{}

	var complianceStandardsField []settings.ComplianceStandard
	for _, itemPb := range pb.ComplianceStandards {
		item, err := settings.ComplianceStandardFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			complianceStandardsField = append(complianceStandardsField, *item)
		}
	}
	st.ComplianceStandards = complianceStandardsField
	st.IsEnabled = pb.IsEnabled

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *complianceSecurityProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st complianceSecurityProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Create an asset
type CreateCleanRoomAssetRequest struct {
	// Metadata of the clean room asset
	// Wire name: 'asset'
	Asset CleanRoomAsset
	// Name of the clean room.
	// Wire name: 'clean_room_name'
	CleanRoomName string `tf:"-"`
}

func createCleanRoomAssetRequestToPb(st *CreateCleanRoomAssetRequest) (*createCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomAssetRequestPb{}
	assetPb, err := cleanRoomAssetToPb(&st.Asset)
	if err != nil {
		return nil, err
	}
	if assetPb != nil {
		pb.Asset = *assetPb
	}

	pb.CleanRoomName = st.CleanRoomName

	return pb, nil
}

func (st *CreateCleanRoomAssetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCleanRoomAssetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCleanRoomAssetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCleanRoomAssetRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCleanRoomAssetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCleanRoomAssetRequestPb struct {
	// Metadata of the clean room asset
	Asset cleanRoomAssetPb `json:"asset"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
}

func createCleanRoomAssetRequestFromPb(pb *createCleanRoomAssetRequestPb) (*CreateCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomAssetRequest{}
	assetField, err := cleanRoomAssetFromPb(&pb.Asset)
	if err != nil {
		return nil, err
	}
	if assetField != nil {
		st.Asset = *assetField
	}
	st.CleanRoomName = pb.CleanRoomName

	return st, nil
}

// Create an output catalog
type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	// Wire name: 'clean_room_name'
	CleanRoomName string `tf:"-"`

	// Wire name: 'output_catalog'
	OutputCatalog CleanRoomOutputCatalog
}

func createCleanRoomOutputCatalogRequestToPb(st *CreateCleanRoomOutputCatalogRequest) (*createCleanRoomOutputCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomOutputCatalogRequestPb{}
	pb.CleanRoomName = st.CleanRoomName

	outputCatalogPb, err := cleanRoomOutputCatalogToPb(&st.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogPb != nil {
		pb.OutputCatalog = *outputCatalogPb
	}

	return pb, nil
}

func (st *CreateCleanRoomOutputCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCleanRoomOutputCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCleanRoomOutputCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCleanRoomOutputCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCleanRoomOutputCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCleanRoomOutputCatalogRequestPb struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`

	OutputCatalog cleanRoomOutputCatalogPb `json:"output_catalog"`
}

func createCleanRoomOutputCatalogRequestFromPb(pb *createCleanRoomOutputCatalogRequestPb) (*CreateCleanRoomOutputCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomOutputCatalogRequest{}
	st.CleanRoomName = pb.CleanRoomName
	outputCatalogField, err := cleanRoomOutputCatalogFromPb(&pb.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogField != nil {
		st.OutputCatalog = *outputCatalogField
	}

	return st, nil
}

type CreateCleanRoomOutputCatalogResponse struct {

	// Wire name: 'output_catalog'
	OutputCatalog *CleanRoomOutputCatalog
}

func createCleanRoomOutputCatalogResponseToPb(st *CreateCleanRoomOutputCatalogResponse) (*createCleanRoomOutputCatalogResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomOutputCatalogResponsePb{}
	outputCatalogPb, err := cleanRoomOutputCatalogToPb(st.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogPb != nil {
		pb.OutputCatalog = outputCatalogPb
	}

	return pb, nil
}

func (st *CreateCleanRoomOutputCatalogResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCleanRoomOutputCatalogResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCleanRoomOutputCatalogResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCleanRoomOutputCatalogResponse) MarshalJSON() ([]byte, error) {
	pb, err := createCleanRoomOutputCatalogResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCleanRoomOutputCatalogResponsePb struct {
	OutputCatalog *cleanRoomOutputCatalogPb `json:"output_catalog,omitempty"`
}

func createCleanRoomOutputCatalogResponseFromPb(pb *createCleanRoomOutputCatalogResponsePb) (*CreateCleanRoomOutputCatalogResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomOutputCatalogResponse{}
	outputCatalogField, err := cleanRoomOutputCatalogFromPb(pb.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogField != nil {
		st.OutputCatalog = outputCatalogField
	}

	return st, nil
}

// Create a clean room
type CreateCleanRoomRequest struct {

	// Wire name: 'clean_room'
	CleanRoom CleanRoom
}

func createCleanRoomRequestToPb(st *CreateCleanRoomRequest) (*createCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomRequestPb{}
	cleanRoomPb, err := cleanRoomToPb(&st.CleanRoom)
	if err != nil {
		return nil, err
	}
	if cleanRoomPb != nil {
		pb.CleanRoom = *cleanRoomPb
	}

	return pb, nil
}

func (st *CreateCleanRoomRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCleanRoomRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCleanRoomRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCleanRoomRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCleanRoomRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCleanRoomRequestPb struct {
	CleanRoom cleanRoomPb `json:"clean_room"`
}

func createCleanRoomRequestFromPb(pb *createCleanRoomRequestPb) (*CreateCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomRequest{}
	cleanRoomField, err := cleanRoomFromPb(&pb.CleanRoom)
	if err != nil {
		return nil, err
	}
	if cleanRoomField != nil {
		st.CleanRoom = *cleanRoomField
	}

	return st, nil
}

// Delete an asset
type DeleteCleanRoomAssetRequest struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	// Wire name: 'asset_full_name'
	AssetFullName string `tf:"-"`
	// The type of the asset.
	// Wire name: 'asset_type'
	AssetType CleanRoomAssetAssetType `tf:"-"`
	// Name of the clean room.
	// Wire name: 'clean_room_name'
	CleanRoomName string `tf:"-"`
}

func deleteCleanRoomAssetRequestToPb(st *DeleteCleanRoomAssetRequest) (*deleteCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCleanRoomAssetRequestPb{}
	pb.AssetFullName = st.AssetFullName

	pb.AssetType = st.AssetType

	pb.CleanRoomName = st.CleanRoomName

	return pb, nil
}

func (st *DeleteCleanRoomAssetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCleanRoomAssetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCleanRoomAssetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCleanRoomAssetRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCleanRoomAssetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCleanRoomAssetRequestPb struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	AssetFullName string `json:"-" url:"-"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
}

func deleteCleanRoomAssetRequestFromPb(pb *deleteCleanRoomAssetRequestPb) (*DeleteCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCleanRoomAssetRequest{}
	st.AssetFullName = pb.AssetFullName
	st.AssetType = pb.AssetType
	st.CleanRoomName = pb.CleanRoomName

	return st, nil
}

// Response for delete clean room request. Using an empty message since the
// generic Empty proto does not externd UnshadedMessageMarker.
type DeleteCleanRoomAssetResponse struct {
}

func deleteCleanRoomAssetResponseToPb(st *DeleteCleanRoomAssetResponse) (*deleteCleanRoomAssetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCleanRoomAssetResponsePb{}

	return pb, nil
}

func (st *DeleteCleanRoomAssetResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCleanRoomAssetResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCleanRoomAssetResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCleanRoomAssetResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteCleanRoomAssetResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCleanRoomAssetResponsePb struct {
}

func deleteCleanRoomAssetResponseFromPb(pb *deleteCleanRoomAssetResponsePb) (*DeleteCleanRoomAssetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCleanRoomAssetResponse{}

	return st, nil
}

// Delete a clean room
type DeleteCleanRoomRequest struct {
	// Name of the clean room.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteCleanRoomRequestToPb(st *DeleteCleanRoomRequest) (*deleteCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCleanRoomRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteCleanRoomRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCleanRoomRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCleanRoomRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCleanRoomRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCleanRoomRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCleanRoomRequestPb struct {
	// Name of the clean room.
	Name string `json:"-" url:"-"`
}

func deleteCleanRoomRequestFromPb(pb *deleteCleanRoomRequestPb) (*DeleteCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCleanRoomRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteResponse struct {
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
}

func (st *DeleteResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteResponsePb struct {
}

func deleteResponseFromPb(pb *deleteResponsePb) (*DeleteResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteResponse{}

	return st, nil
}

// Get an asset
type GetCleanRoomAssetRequest struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	// Wire name: 'asset_full_name'
	AssetFullName string `tf:"-"`
	// The type of the asset.
	// Wire name: 'asset_type'
	AssetType CleanRoomAssetAssetType `tf:"-"`
	// Name of the clean room.
	// Wire name: 'clean_room_name'
	CleanRoomName string `tf:"-"`
}

func getCleanRoomAssetRequestToPb(st *GetCleanRoomAssetRequest) (*getCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCleanRoomAssetRequestPb{}
	pb.AssetFullName = st.AssetFullName

	pb.AssetType = st.AssetType

	pb.CleanRoomName = st.CleanRoomName

	return pb, nil
}

func (st *GetCleanRoomAssetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCleanRoomAssetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCleanRoomAssetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCleanRoomAssetRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCleanRoomAssetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCleanRoomAssetRequestPb struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	AssetFullName string `json:"-" url:"-"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
}

func getCleanRoomAssetRequestFromPb(pb *getCleanRoomAssetRequestPb) (*GetCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCleanRoomAssetRequest{}
	st.AssetFullName = pb.AssetFullName
	st.AssetType = pb.AssetType
	st.CleanRoomName = pb.CleanRoomName

	return st, nil
}

// Get a clean room
type GetCleanRoomRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func getCleanRoomRequestToPb(st *GetCleanRoomRequest) (*getCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCleanRoomRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *GetCleanRoomRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCleanRoomRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCleanRoomRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCleanRoomRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCleanRoomRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCleanRoomRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getCleanRoomRequestFromPb(pb *getCleanRoomRequestPb) (*GetCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCleanRoomRequest{}
	st.Name = pb.Name

	return st, nil
}

// List assets
type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	// Wire name: 'clean_room_name'
	CleanRoomName string `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listCleanRoomAssetsRequestToPb(st *ListCleanRoomAssetsRequest) (*listCleanRoomAssetsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomAssetsRequestPb{}
	pb.CleanRoomName = st.CleanRoomName

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCleanRoomAssetsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCleanRoomAssetsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCleanRoomAssetsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCleanRoomAssetsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listCleanRoomAssetsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCleanRoomAssetsRequestPb struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomAssetsRequestFromPb(pb *listCleanRoomAssetsRequestPb) (*ListCleanRoomAssetsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAssetsRequest{}
	st.CleanRoomName = pb.CleanRoomName
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCleanRoomAssetsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCleanRoomAssetsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomAssetsResponse struct {
	// Assets in the clean room.
	// Wire name: 'assets'
	Assets []CleanRoomAsset
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listCleanRoomAssetsResponseToPb(st *ListCleanRoomAssetsResponse) (*listCleanRoomAssetsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomAssetsResponsePb{}

	var assetsPb []cleanRoomAssetPb
	for _, item := range st.Assets {
		itemPb, err := cleanRoomAssetToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			assetsPb = append(assetsPb, *itemPb)
		}
	}
	pb.Assets = assetsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCleanRoomAssetsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCleanRoomAssetsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCleanRoomAssetsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCleanRoomAssetsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listCleanRoomAssetsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCleanRoomAssetsResponsePb struct {
	// Assets in the clean room.
	Assets []cleanRoomAssetPb `json:"assets,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomAssetsResponseFromPb(pb *listCleanRoomAssetsResponsePb) (*ListCleanRoomAssetsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAssetsResponse{}

	var assetsField []CleanRoomAsset
	for _, itemPb := range pb.Assets {
		item, err := cleanRoomAssetFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			assetsField = append(assetsField, *item)
		}
	}
	st.Assets = assetsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCleanRoomAssetsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCleanRoomAssetsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List notebook task runs
type ListCleanRoomNotebookTaskRunsRequest struct {
	// Name of the clean room.
	// Wire name: 'clean_room_name'
	CleanRoomName string `tf:"-"`
	// Notebook name
	// Wire name: 'notebook_name'
	NotebookName string `tf:"-"`
	// The maximum number of task runs to return. Currently ignored - all runs
	// will be returned.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listCleanRoomNotebookTaskRunsRequestToPb(st *ListCleanRoomNotebookTaskRunsRequest) (*listCleanRoomNotebookTaskRunsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomNotebookTaskRunsRequestPb{}
	pb.CleanRoomName = st.CleanRoomName

	pb.NotebookName = st.NotebookName

	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCleanRoomNotebookTaskRunsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCleanRoomNotebookTaskRunsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCleanRoomNotebookTaskRunsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCleanRoomNotebookTaskRunsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listCleanRoomNotebookTaskRunsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCleanRoomNotebookTaskRunsRequestPb struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// Notebook name
	NotebookName string `json:"-" url:"notebook_name,omitempty"`
	// The maximum number of task runs to return. Currently ignored - all runs
	// will be returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomNotebookTaskRunsRequestFromPb(pb *listCleanRoomNotebookTaskRunsRequestPb) (*ListCleanRoomNotebookTaskRunsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomNotebookTaskRunsRequest{}
	st.CleanRoomName = pb.CleanRoomName
	st.NotebookName = pb.NotebookName
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCleanRoomNotebookTaskRunsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCleanRoomNotebookTaskRunsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomNotebookTaskRunsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// Name of the clean room.
	// Wire name: 'runs'
	Runs []CleanRoomNotebookTaskRun

	ForceSendFields []string `tf:"-"`
}

func listCleanRoomNotebookTaskRunsResponseToPb(st *ListCleanRoomNotebookTaskRunsResponse) (*listCleanRoomNotebookTaskRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomNotebookTaskRunsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var runsPb []cleanRoomNotebookTaskRunPb
	for _, item := range st.Runs {
		itemPb, err := cleanRoomNotebookTaskRunToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			runsPb = append(runsPb, *itemPb)
		}
	}
	pb.Runs = runsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCleanRoomNotebookTaskRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCleanRoomNotebookTaskRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCleanRoomNotebookTaskRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCleanRoomNotebookTaskRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listCleanRoomNotebookTaskRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCleanRoomNotebookTaskRunsResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// Name of the clean room.
	Runs []cleanRoomNotebookTaskRunPb `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomNotebookTaskRunsResponseFromPb(pb *listCleanRoomNotebookTaskRunsResponsePb) (*ListCleanRoomNotebookTaskRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomNotebookTaskRunsResponse{}
	st.NextPageToken = pb.NextPageToken

	var runsField []CleanRoomNotebookTaskRun
	for _, itemPb := range pb.Runs {
		item, err := cleanRoomNotebookTaskRunFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			runsField = append(runsField, *item)
		}
	}
	st.Runs = runsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCleanRoomNotebookTaskRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCleanRoomNotebookTaskRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List clean rooms
type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listCleanRoomsRequestToPb(st *ListCleanRoomsRequest) (*listCleanRoomsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomsRequestPb{}
	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCleanRoomsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCleanRoomsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCleanRoomsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCleanRoomsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listCleanRoomsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCleanRoomsRequestPb struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomsRequestFromPb(pb *listCleanRoomsRequestPb) (*ListCleanRoomsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCleanRoomsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCleanRoomsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomsResponse struct {

	// Wire name: 'clean_rooms'
	CleanRooms []CleanRoom
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listCleanRoomsResponseToPb(st *ListCleanRoomsResponse) (*listCleanRoomsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomsResponsePb{}

	var cleanRoomsPb []cleanRoomPb
	for _, item := range st.CleanRooms {
		itemPb, err := cleanRoomToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			cleanRoomsPb = append(cleanRoomsPb, *itemPb)
		}
	}
	pb.CleanRooms = cleanRoomsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCleanRoomsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCleanRoomsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCleanRoomsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCleanRoomsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listCleanRoomsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCleanRoomsResponsePb struct {
	CleanRooms []cleanRoomPb `json:"clean_rooms,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomsResponseFromPb(pb *listCleanRoomsResponsePb) (*ListCleanRoomsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomsResponse{}

	var cleanRoomsField []CleanRoom
	for _, itemPb := range pb.CleanRooms {
		item, err := cleanRoomFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			cleanRoomsField = append(cleanRoomsField, *item)
		}
	}
	st.CleanRooms = cleanRoomsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCleanRoomsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCleanRoomsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Update an asset
type UpdateCleanRoomAssetRequest struct {
	// Metadata of the clean room asset
	// Wire name: 'asset'
	Asset CleanRoomAsset
	// The type of the asset.
	// Wire name: 'asset_type'
	AssetType CleanRoomAssetAssetType `tf:"-"`
	// Name of the clean room.
	// Wire name: 'clean_room_name'
	CleanRoomName string `tf:"-"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func updateCleanRoomAssetRequestToPb(st *UpdateCleanRoomAssetRequest) (*updateCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCleanRoomAssetRequestPb{}
	assetPb, err := cleanRoomAssetToPb(&st.Asset)
	if err != nil {
		return nil, err
	}
	if assetPb != nil {
		pb.Asset = *assetPb
	}

	pb.AssetType = st.AssetType

	pb.CleanRoomName = st.CleanRoomName

	pb.Name = st.Name

	return pb, nil
}

func (st *UpdateCleanRoomAssetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCleanRoomAssetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCleanRoomAssetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCleanRoomAssetRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateCleanRoomAssetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCleanRoomAssetRequestPb struct {
	// Metadata of the clean room asset
	Asset cleanRoomAssetPb `json:"asset"`
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

func updateCleanRoomAssetRequestFromPb(pb *updateCleanRoomAssetRequestPb) (*UpdateCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCleanRoomAssetRequest{}
	assetField, err := cleanRoomAssetFromPb(&pb.Asset)
	if err != nil {
		return nil, err
	}
	if assetField != nil {
		st.Asset = *assetField
	}
	st.AssetType = pb.AssetType
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name

	return st, nil
}

type UpdateCleanRoomRequest struct {

	// Wire name: 'clean_room'
	CleanRoom *CleanRoom
	// Name of the clean room.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func updateCleanRoomRequestToPb(st *UpdateCleanRoomRequest) (*updateCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCleanRoomRequestPb{}
	cleanRoomPb, err := cleanRoomToPb(st.CleanRoom)
	if err != nil {
		return nil, err
	}
	if cleanRoomPb != nil {
		pb.CleanRoom = cleanRoomPb
	}

	pb.Name = st.Name

	return pb, nil
}

func (st *UpdateCleanRoomRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCleanRoomRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCleanRoomRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCleanRoomRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateCleanRoomRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCleanRoomRequestPb struct {
	CleanRoom *cleanRoomPb `json:"clean_room,omitempty"`
	// Name of the clean room.
	Name string `json:"-" url:"-"`
}

func updateCleanRoomRequestFromPb(pb *updateCleanRoomRequestPb) (*UpdateCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCleanRoomRequest{}
	cleanRoomField, err := cleanRoomFromPb(pb.CleanRoom)
	if err != nil {
		return nil, err
	}
	if cleanRoomField != nil {
		st.CleanRoom = cleanRoomField
	}
	st.Name = pb.Name

	return st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
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
