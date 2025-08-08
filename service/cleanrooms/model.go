// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/catalog/catalogpb"
	"github.com/databricks/databricks-sdk-go/service/cleanrooms/cleanroomspb"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/settings/settingspb"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/databricks-sdk-go/service/sharing/sharingpb"
)

type CleanRoom struct {
	// Whether clean room access is restricted due to [CSP]
	//
	// [CSP]: https://docs.databricks.com/en/security/privacy/security-profile.html
	// Wire name: 'access_restricted'
	AccessRestricted CleanRoomAccessRestricted `json:"access_restricted,omitempty"`

	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// When the clean room was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// The alias of the collaborator tied to the local clean room.
	// Wire name: 'local_collaborator_alias'
	LocalCollaboratorAlias string `json:"local_collaborator_alias,omitempty"`
	// The name of the clean room. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Output catalog of the clean room. It is an output only field. Output
	// catalog is manipulated using the separate CreateCleanRoomOutputCatalog
	// API.
	// Wire name: 'output_catalog'
	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
	// This is Databricks username of the owner of the local clean room
	// securable for permission management.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	// Wire name: 'remote_detailed_info'
	RemoteDetailedInfo *CleanRoomRemoteDetail `json:"remote_detailed_info,omitempty"`
	// Clean room status.
	// Wire name: 'status'
	Status CleanRoomStatusEnum `json:"status,omitempty"`
	// When the clean room was last updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt       int64    `json:"updated_at,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoom) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoom) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomToPb(st *CleanRoom) (*cleanroomspb.CleanRoomPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomPb{}
	accessRestrictedPb, err := CleanRoomAccessRestrictedToPb(&st.AccessRestricted)
	if err != nil {
		return nil, err
	}
	if accessRestrictedPb != nil {
		pb.AccessRestricted = *accessRestrictedPb
	}
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.LocalCollaboratorAlias = st.LocalCollaboratorAlias
	pb.Name = st.Name
	outputCatalogPb, err := CleanRoomOutputCatalogToPb(st.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogPb != nil {
		pb.OutputCatalog = outputCatalogPb
	}
	pb.Owner = st.Owner
	remoteDetailedInfoPb, err := CleanRoomRemoteDetailToPb(st.RemoteDetailedInfo)
	if err != nil {
		return nil, err
	}
	if remoteDetailedInfoPb != nil {
		pb.RemoteDetailedInfo = remoteDetailedInfoPb
	}
	statusPb, err := CleanRoomStatusEnumToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.UpdatedAt = st.UpdatedAt

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomFromPb(pb *cleanroomspb.CleanRoomPb) (*CleanRoom, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoom{}
	accessRestrictedField, err := CleanRoomAccessRestrictedFromPb(&pb.AccessRestricted)
	if err != nil {
		return nil, err
	}
	if accessRestrictedField != nil {
		st.AccessRestricted = *accessRestrictedField
	}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.LocalCollaboratorAlias = pb.LocalCollaboratorAlias
	st.Name = pb.Name
	outputCatalogField, err := CleanRoomOutputCatalogFromPb(pb.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogField != nil {
		st.OutputCatalog = outputCatalogField
	}
	st.Owner = pb.Owner
	remoteDetailedInfoField, err := CleanRoomRemoteDetailFromPb(pb.RemoteDetailedInfo)
	if err != nil {
		return nil, err
	}
	if remoteDetailedInfoField != nil {
		st.RemoteDetailedInfo = remoteDetailedInfoField
	}
	statusField, err := CleanRoomStatusEnumFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.UpdatedAt = pb.UpdatedAt

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

// Values returns all possible values for CleanRoomAccessRestricted.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomAccessRestricted) Values() []CleanRoomAccessRestricted {
	return []CleanRoomAccessRestricted{
		CleanRoomAccessRestrictedCspMismatch,
		CleanRoomAccessRestrictedNoRestriction,
	}
}

// Type always returns CleanRoomAccessRestricted to satisfy [pflag.Value] interface
func (f *CleanRoomAccessRestricted) Type() string {
	return "CleanRoomAccessRestricted"
}

func CleanRoomAccessRestrictedToPb(st *CleanRoomAccessRestricted) (*cleanroomspb.CleanRoomAccessRestrictedPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanroomspb.CleanRoomAccessRestrictedPb(*st)
	return &pb, nil
}

func CleanRoomAccessRestrictedFromPb(pb *cleanroomspb.CleanRoomAccessRestrictedPb) (*CleanRoomAccessRestricted, error) {
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
	AddedAt int64 `json:"added_at,omitempty"`
	// The type of the asset.
	// Wire name: 'asset_type'
	AssetType CleanRoomAssetAssetType `json:"asset_type"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	// Wire name: 'clean_room_name'
	CleanRoomName string `json:"clean_room_name,omitempty"`
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	// Wire name: 'foreign_table'
	ForeignTable *CleanRoomAssetForeignTable `json:"foreign_table,omitempty"`
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	// Wire name: 'foreign_table_local_details'
	ForeignTableLocalDetails *CleanRoomAssetForeignTableLocalDetails `json:"foreign_table_local_details,omitempty"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	// Wire name: 'name'
	Name string `json:"name"`
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	// Wire name: 'notebook'
	Notebook *CleanRoomAssetNotebook `json:"notebook,omitempty"`
	// The alias of the collaborator who owns this asset
	// Wire name: 'owner_collaborator_alias'
	OwnerCollaboratorAlias string `json:"owner_collaborator_alias,omitempty"`
	// Status of the asset
	// Wire name: 'status'
	Status CleanRoomAssetStatusEnum `json:"status,omitempty"`
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	// Wire name: 'table'
	Table *CleanRoomAssetTable `json:"table,omitempty"`
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	// Wire name: 'table_local_details'
	TableLocalDetails *CleanRoomAssetTableLocalDetails `json:"table_local_details,omitempty"`
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	// Wire name: 'view'
	View *CleanRoomAssetView `json:"view,omitempty"`
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	// Wire name: 'view_local_details'
	ViewLocalDetails *CleanRoomAssetViewLocalDetails `json:"view_local_details,omitempty"`
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	// Wire name: 'volume_local_details'
	VolumeLocalDetails *CleanRoomAssetVolumeLocalDetails `json:"volume_local_details,omitempty"`
	ForceSendFields    []string                          `json:"-" tf:"-"`
}

func (st CleanRoomAsset) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAsset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetToPb(st *CleanRoomAsset) (*cleanroomspb.CleanRoomAssetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetPb{}
	pb.AddedAt = st.AddedAt
	assetTypePb, err := CleanRoomAssetAssetTypeToPb(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}
	pb.CleanRoomName = st.CleanRoomName
	foreignTablePb, err := CleanRoomAssetForeignTableToPb(st.ForeignTable)
	if err != nil {
		return nil, err
	}
	if foreignTablePb != nil {
		pb.ForeignTable = foreignTablePb
	}
	foreignTableLocalDetailsPb, err := CleanRoomAssetForeignTableLocalDetailsToPb(st.ForeignTableLocalDetails)
	if err != nil {
		return nil, err
	}
	if foreignTableLocalDetailsPb != nil {
		pb.ForeignTableLocalDetails = foreignTableLocalDetailsPb
	}
	pb.Name = st.Name
	notebookPb, err := CleanRoomAssetNotebookToPb(st.Notebook)
	if err != nil {
		return nil, err
	}
	if notebookPb != nil {
		pb.Notebook = notebookPb
	}
	pb.OwnerCollaboratorAlias = st.OwnerCollaboratorAlias
	statusPb, err := CleanRoomAssetStatusEnumToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	tablePb, err := CleanRoomAssetTableToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}
	tableLocalDetailsPb, err := CleanRoomAssetTableLocalDetailsToPb(st.TableLocalDetails)
	if err != nil {
		return nil, err
	}
	if tableLocalDetailsPb != nil {
		pb.TableLocalDetails = tableLocalDetailsPb
	}
	viewPb, err := CleanRoomAssetViewToPb(st.View)
	if err != nil {
		return nil, err
	}
	if viewPb != nil {
		pb.View = viewPb
	}
	viewLocalDetailsPb, err := CleanRoomAssetViewLocalDetailsToPb(st.ViewLocalDetails)
	if err != nil {
		return nil, err
	}
	if viewLocalDetailsPb != nil {
		pb.ViewLocalDetails = viewLocalDetailsPb
	}
	volumeLocalDetailsPb, err := CleanRoomAssetVolumeLocalDetailsToPb(st.VolumeLocalDetails)
	if err != nil {
		return nil, err
	}
	if volumeLocalDetailsPb != nil {
		pb.VolumeLocalDetails = volumeLocalDetailsPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomAssetFromPb(pb *cleanroomspb.CleanRoomAssetPb) (*CleanRoomAsset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAsset{}
	st.AddedAt = pb.AddedAt
	assetTypeField, err := CleanRoomAssetAssetTypeFromPb(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	st.CleanRoomName = pb.CleanRoomName
	foreignTableField, err := CleanRoomAssetForeignTableFromPb(pb.ForeignTable)
	if err != nil {
		return nil, err
	}
	if foreignTableField != nil {
		st.ForeignTable = foreignTableField
	}
	foreignTableLocalDetailsField, err := CleanRoomAssetForeignTableLocalDetailsFromPb(pb.ForeignTableLocalDetails)
	if err != nil {
		return nil, err
	}
	if foreignTableLocalDetailsField != nil {
		st.ForeignTableLocalDetails = foreignTableLocalDetailsField
	}
	st.Name = pb.Name
	notebookField, err := CleanRoomAssetNotebookFromPb(pb.Notebook)
	if err != nil {
		return nil, err
	}
	if notebookField != nil {
		st.Notebook = notebookField
	}
	st.OwnerCollaboratorAlias = pb.OwnerCollaboratorAlias
	statusField, err := CleanRoomAssetStatusEnumFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	tableField, err := CleanRoomAssetTableFromPb(pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = tableField
	}
	tableLocalDetailsField, err := CleanRoomAssetTableLocalDetailsFromPb(pb.TableLocalDetails)
	if err != nil {
		return nil, err
	}
	if tableLocalDetailsField != nil {
		st.TableLocalDetails = tableLocalDetailsField
	}
	viewField, err := CleanRoomAssetViewFromPb(pb.View)
	if err != nil {
		return nil, err
	}
	if viewField != nil {
		st.View = viewField
	}
	viewLocalDetailsField, err := CleanRoomAssetViewLocalDetailsFromPb(pb.ViewLocalDetails)
	if err != nil {
		return nil, err
	}
	if viewLocalDetailsField != nil {
		st.ViewLocalDetails = viewLocalDetailsField
	}
	volumeLocalDetailsField, err := CleanRoomAssetVolumeLocalDetailsFromPb(pb.VolumeLocalDetails)
	if err != nil {
		return nil, err
	}
	if volumeLocalDetailsField != nil {
		st.VolumeLocalDetails = volumeLocalDetailsField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

// Values returns all possible values for CleanRoomAssetAssetType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomAssetAssetType) Values() []CleanRoomAssetAssetType {
	return []CleanRoomAssetAssetType{
		CleanRoomAssetAssetTypeForeignTable,
		CleanRoomAssetAssetTypeNotebookFile,
		CleanRoomAssetAssetTypeTable,
		CleanRoomAssetAssetTypeView,
		CleanRoomAssetAssetTypeVolume,
	}
}

// Type always returns CleanRoomAssetAssetType to satisfy [pflag.Value] interface
func (f *CleanRoomAssetAssetType) Type() string {
	return "CleanRoomAssetAssetType"
}

func CleanRoomAssetAssetTypeToPb(st *CleanRoomAssetAssetType) (*cleanroomspb.CleanRoomAssetAssetTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanroomspb.CleanRoomAssetAssetTypePb(*st)
	return &pb, nil
}

func CleanRoomAssetAssetTypeFromPb(pb *cleanroomspb.CleanRoomAssetAssetTypePb) (*CleanRoomAssetAssetType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomAssetAssetType(*pb)
	return &st, nil
}

type CleanRoomAssetForeignTable struct {
	// The metadata information of the columns in the foreign table
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func (st CleanRoomAssetForeignTable) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetForeignTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAssetForeignTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetForeignTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetForeignTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetForeignTableToPb(st *CleanRoomAssetForeignTable) (*cleanroomspb.CleanRoomAssetForeignTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetForeignTablePb{}

	var columnsPb []catalogpb.ColumnInfoPb
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

func CleanRoomAssetForeignTableFromPb(pb *cleanroomspb.CleanRoomAssetForeignTablePb) (*CleanRoomAssetForeignTable, error) {
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
	LocalName string `json:"local_name"`
}

func (st CleanRoomAssetForeignTableLocalDetails) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetForeignTableLocalDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAssetForeignTableLocalDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetForeignTableLocalDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetForeignTableLocalDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetForeignTableLocalDetailsToPb(st *CleanRoomAssetForeignTableLocalDetails) (*cleanroomspb.CleanRoomAssetForeignTableLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetForeignTableLocalDetailsPb{}
	pb.LocalName = st.LocalName

	return pb, nil
}

func CleanRoomAssetForeignTableLocalDetailsFromPb(pb *cleanroomspb.CleanRoomAssetForeignTableLocalDetailsPb) (*CleanRoomAssetForeignTableLocalDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetForeignTableLocalDetails{}
	st.LocalName = pb.LocalName

	return st, nil
}

type CleanRoomAssetNotebook struct {
	// Server generated etag that represents the notebook version.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	// Wire name: 'notebook_content'
	NotebookContent string `json:"notebook_content"`
	// top-level status derived from all reviews
	// Wire name: 'review_state'
	ReviewState CleanRoomNotebookReviewNotebookReviewState `json:"review_state,omitempty"`
	// All existing approvals or rejections
	// Wire name: 'reviews'
	Reviews []CleanRoomNotebookReview `json:"reviews,omitempty"`
	// collaborators that can run the notebook
	// Wire name: 'runner_collaborator_aliases'
	RunnerCollaboratorAliases []string `json:"runner_collaborator_aliases,omitempty"`
	ForceSendFields           []string `json:"-" tf:"-"`
}

func (st CleanRoomAssetNotebook) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetNotebookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAssetNotebook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetNotebookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetNotebookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetNotebookToPb(st *CleanRoomAssetNotebook) (*cleanroomspb.CleanRoomAssetNotebookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetNotebookPb{}
	pb.Etag = st.Etag
	pb.NotebookContent = st.NotebookContent
	reviewStatePb, err := CleanRoomNotebookReviewNotebookReviewStateToPb(&st.ReviewState)
	if err != nil {
		return nil, err
	}
	if reviewStatePb != nil {
		pb.ReviewState = *reviewStatePb
	}

	var reviewsPb []cleanroomspb.CleanRoomNotebookReviewPb
	for _, item := range st.Reviews {
		itemPb, err := CleanRoomNotebookReviewToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			reviewsPb = append(reviewsPb, *itemPb)
		}
	}
	pb.Reviews = reviewsPb
	pb.RunnerCollaboratorAliases = st.RunnerCollaboratorAliases

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomAssetNotebookFromPb(pb *cleanroomspb.CleanRoomAssetNotebookPb) (*CleanRoomAssetNotebook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetNotebook{}
	st.Etag = pb.Etag
	st.NotebookContent = pb.NotebookContent
	reviewStateField, err := CleanRoomNotebookReviewNotebookReviewStateFromPb(&pb.ReviewState)
	if err != nil {
		return nil, err
	}
	if reviewStateField != nil {
		st.ReviewState = *reviewStateField
	}

	var reviewsField []CleanRoomNotebookReview
	for _, itemPb := range pb.Reviews {
		item, err := CleanRoomNotebookReviewFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			reviewsField = append(reviewsField, *item)
		}
	}
	st.Reviews = reviewsField
	st.RunnerCollaboratorAliases = pb.RunnerCollaboratorAliases

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

// Values returns all possible values for CleanRoomAssetStatusEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomAssetStatusEnum) Values() []CleanRoomAssetStatusEnum {
	return []CleanRoomAssetStatusEnum{
		CleanRoomAssetStatusEnumActive,
		CleanRoomAssetStatusEnumPending,
		CleanRoomAssetStatusEnumPermissionDenied,
	}
}

// Type always returns CleanRoomAssetStatusEnum to satisfy [pflag.Value] interface
func (f *CleanRoomAssetStatusEnum) Type() string {
	return "CleanRoomAssetStatusEnum"
}

func CleanRoomAssetStatusEnumToPb(st *CleanRoomAssetStatusEnum) (*cleanroomspb.CleanRoomAssetStatusEnumPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanroomspb.CleanRoomAssetStatusEnumPb(*st)
	return &pb, nil
}

func CleanRoomAssetStatusEnumFromPb(pb *cleanroomspb.CleanRoomAssetStatusEnumPb) (*CleanRoomAssetStatusEnum, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomAssetStatusEnum(*pb)
	return &st, nil
}

type CleanRoomAssetTable struct {
	// The metadata information of the columns in the table
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func (st CleanRoomAssetTable) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAssetTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetTableToPb(st *CleanRoomAssetTable) (*cleanroomspb.CleanRoomAssetTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetTablePb{}

	var columnsPb []catalogpb.ColumnInfoPb
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

func CleanRoomAssetTableFromPb(pb *cleanroomspb.CleanRoomAssetTablePb) (*CleanRoomAssetTable, error) {
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
	LocalName string `json:"local_name"`
	// Partition filtering specification for a shared table.
	// Wire name: 'partitions'
	Partitions []sharing.Partition `json:"partitions,omitempty"`
}

func (st CleanRoomAssetTableLocalDetails) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetTableLocalDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAssetTableLocalDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetTableLocalDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetTableLocalDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetTableLocalDetailsToPb(st *CleanRoomAssetTableLocalDetails) (*cleanroomspb.CleanRoomAssetTableLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetTableLocalDetailsPb{}
	pb.LocalName = st.LocalName

	var partitionsPb []sharingpb.PartitionPb
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

	return pb, nil
}

func CleanRoomAssetTableLocalDetailsFromPb(pb *cleanroomspb.CleanRoomAssetTableLocalDetailsPb) (*CleanRoomAssetTableLocalDetails, error) {
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

	return st, nil
}

type CleanRoomAssetView struct {
	// The metadata information of the columns in the view
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func (st CleanRoomAssetView) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetViewToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAssetView) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetViewPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetViewFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetViewToPb(st *CleanRoomAssetView) (*cleanroomspb.CleanRoomAssetViewPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetViewPb{}

	var columnsPb []catalogpb.ColumnInfoPb
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

func CleanRoomAssetViewFromPb(pb *cleanroomspb.CleanRoomAssetViewPb) (*CleanRoomAssetView, error) {
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
	LocalName string `json:"local_name"`
}

func (st CleanRoomAssetViewLocalDetails) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetViewLocalDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAssetViewLocalDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetViewLocalDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetViewLocalDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetViewLocalDetailsToPb(st *CleanRoomAssetViewLocalDetails) (*cleanroomspb.CleanRoomAssetViewLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetViewLocalDetailsPb{}
	pb.LocalName = st.LocalName

	return pb, nil
}

func CleanRoomAssetViewLocalDetailsFromPb(pb *cleanroomspb.CleanRoomAssetViewLocalDetailsPb) (*CleanRoomAssetViewLocalDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetViewLocalDetails{}
	st.LocalName = pb.LocalName

	return st, nil
}

type CleanRoomAssetVolumeLocalDetails struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	// Wire name: 'local_name'
	LocalName string `json:"local_name"`
}

func (st CleanRoomAssetVolumeLocalDetails) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAssetVolumeLocalDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAssetVolumeLocalDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAssetVolumeLocalDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAssetVolumeLocalDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAssetVolumeLocalDetailsToPb(st *CleanRoomAssetVolumeLocalDetails) (*cleanroomspb.CleanRoomAssetVolumeLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAssetVolumeLocalDetailsPb{}
	pb.LocalName = st.LocalName

	return pb, nil
}

func CleanRoomAssetVolumeLocalDetailsFromPb(pb *cleanroomspb.CleanRoomAssetVolumeLocalDetailsPb) (*CleanRoomAssetVolumeLocalDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetVolumeLocalDetails{}
	st.LocalName = pb.LocalName

	return st, nil
}

type CleanRoomAutoApprovalRule struct {

	// Wire name: 'author_collaborator_alias'
	AuthorCollaboratorAlias string `json:"author_collaborator_alias,omitempty"`

	// Wire name: 'author_scope'
	AuthorScope CleanRoomAutoApprovalRuleAuthorScope `json:"author_scope,omitempty"`
	// The name of the clean room this auto-approval rule belongs to.
	// Wire name: 'clean_room_name'
	CleanRoomName string `json:"clean_room_name,omitempty"`
	// Timestamp of when the rule was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// A generated UUID identifying the rule.
	// Wire name: 'rule_id'
	RuleId string `json:"rule_id,omitempty"`
	// The owner of the rule to whom the rule applies.
	// Wire name: 'rule_owner_collaborator_alias'
	RuleOwnerCollaboratorAlias string `json:"rule_owner_collaborator_alias,omitempty"`

	// Wire name: 'runner_collaborator_alias'
	RunnerCollaboratorAlias string   `json:"runner_collaborator_alias,omitempty"`
	ForceSendFields         []string `json:"-" tf:"-"`
}

func (st CleanRoomAutoApprovalRule) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomAutoApprovalRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomAutoApprovalRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomAutoApprovalRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomAutoApprovalRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomAutoApprovalRuleToPb(st *CleanRoomAutoApprovalRule) (*cleanroomspb.CleanRoomAutoApprovalRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomAutoApprovalRulePb{}
	pb.AuthorCollaboratorAlias = st.AuthorCollaboratorAlias
	authorScopePb, err := CleanRoomAutoApprovalRuleAuthorScopeToPb(&st.AuthorScope)
	if err != nil {
		return nil, err
	}
	if authorScopePb != nil {
		pb.AuthorScope = *authorScopePb
	}
	pb.CleanRoomName = st.CleanRoomName
	pb.CreatedAt = st.CreatedAt
	pb.RuleId = st.RuleId
	pb.RuleOwnerCollaboratorAlias = st.RuleOwnerCollaboratorAlias
	pb.RunnerCollaboratorAlias = st.RunnerCollaboratorAlias

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomAutoApprovalRuleFromPb(pb *cleanroomspb.CleanRoomAutoApprovalRulePb) (*CleanRoomAutoApprovalRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAutoApprovalRule{}
	st.AuthorCollaboratorAlias = pb.AuthorCollaboratorAlias
	authorScopeField, err := CleanRoomAutoApprovalRuleAuthorScopeFromPb(&pb.AuthorScope)
	if err != nil {
		return nil, err
	}
	if authorScopeField != nil {
		st.AuthorScope = *authorScopeField
	}
	st.CleanRoomName = pb.CleanRoomName
	st.CreatedAt = pb.CreatedAt
	st.RuleId = pb.RuleId
	st.RuleOwnerCollaboratorAlias = pb.RuleOwnerCollaboratorAlias
	st.RunnerCollaboratorAlias = pb.RunnerCollaboratorAlias

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CleanRoomAutoApprovalRuleAuthorScope string

const CleanRoomAutoApprovalRuleAuthorScopeAnyAuthor CleanRoomAutoApprovalRuleAuthorScope = `ANY_AUTHOR`

// String representation for [fmt.Print]
func (f *CleanRoomAutoApprovalRuleAuthorScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomAutoApprovalRuleAuthorScope) Set(v string) error {
	switch v {
	case `ANY_AUTHOR`:
		*f = CleanRoomAutoApprovalRuleAuthorScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ANY_AUTHOR"`, v)
	}
}

// Values returns all possible values for CleanRoomAutoApprovalRuleAuthorScope.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomAutoApprovalRuleAuthorScope) Values() []CleanRoomAutoApprovalRuleAuthorScope {
	return []CleanRoomAutoApprovalRuleAuthorScope{
		CleanRoomAutoApprovalRuleAuthorScopeAnyAuthor,
	}
}

// Type always returns CleanRoomAutoApprovalRuleAuthorScope to satisfy [pflag.Value] interface
func (f *CleanRoomAutoApprovalRuleAuthorScope) Type() string {
	return "CleanRoomAutoApprovalRuleAuthorScope"
}

func CleanRoomAutoApprovalRuleAuthorScopeToPb(st *CleanRoomAutoApprovalRuleAuthorScope) (*cleanroomspb.CleanRoomAutoApprovalRuleAuthorScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanroomspb.CleanRoomAutoApprovalRuleAuthorScopePb(*st)
	return &pb, nil
}

func CleanRoomAutoApprovalRuleAuthorScopeFromPb(pb *cleanroomspb.CleanRoomAutoApprovalRuleAuthorScopePb) (*CleanRoomAutoApprovalRuleAuthorScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomAutoApprovalRuleAuthorScope(*pb)
	return &st, nil
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
	CollaboratorAlias string `json:"collaborator_alias"`
	// Generated display name for the collaborator. In the case of a single
	// metastore clean room, it is the clean room name. For x-metastore clean
	// rooms, it is the organization name of the metastore. It is not restricted
	// to these values and could change in the future
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// The global Unity Catalog metastore id of the collaborator. The identifier
	// is of format cloud:region:metastore-uuid.
	// Wire name: 'global_metastore_id'
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Email of the user who is receiving the clean room "invitation". It should
	// be empty for the creator of the clean room, and non-empty for the
	// invitees of the clean room. It is only returned in the output when clean
	// room creator calls GET
	// Wire name: 'invite_recipient_email'
	InviteRecipientEmail string `json:"invite_recipient_email,omitempty"`
	// Workspace ID of the user who is receiving the clean room "invitation".
	// Must be specified if invite_recipient_email is specified. It should be
	// empty when the collaborator is the creator of the clean room.
	// Wire name: 'invite_recipient_workspace_id'
	InviteRecipientWorkspaceId int64 `json:"invite_recipient_workspace_id,omitempty"`
	// [Organization
	// name](:method:metastores/list#metastores-delta_sharing_organization_name)
	// configured in the metastore
	// Wire name: 'organization_name'
	OrganizationName string   `json:"organization_name,omitempty"`
	ForceSendFields  []string `json:"-" tf:"-"`
}

func (st CleanRoomCollaborator) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomCollaboratorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomCollaborator) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomCollaboratorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomCollaboratorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomCollaboratorToPb(st *CleanRoomCollaborator) (*cleanroomspb.CleanRoomCollaboratorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomCollaboratorPb{}
	pb.CollaboratorAlias = st.CollaboratorAlias
	pb.DisplayName = st.DisplayName
	pb.GlobalMetastoreId = st.GlobalMetastoreId
	pb.InviteRecipientEmail = st.InviteRecipientEmail
	pb.InviteRecipientWorkspaceId = st.InviteRecipientWorkspaceId
	pb.OrganizationName = st.OrganizationName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomCollaboratorFromPb(pb *cleanroomspb.CleanRoomCollaboratorPb) (*CleanRoomCollaborator, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CleanRoomNotebookReview struct {
	// review comment
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// timestamp of when the review was submitted
	// Wire name: 'created_at_millis'
	CreatedAtMillis int64 `json:"created_at_millis,omitempty"`
	// review outcome
	// Wire name: 'review_state'
	ReviewState CleanRoomNotebookReviewNotebookReviewState `json:"review_state,omitempty"`
	// specified when the review was not explicitly made by a user
	// Wire name: 'review_sub_reason'
	ReviewSubReason CleanRoomNotebookReviewNotebookReviewSubReason `json:"review_sub_reason,omitempty"`
	// collaborator alias of the reviewer
	// Wire name: 'reviewer_collaborator_alias'
	ReviewerCollaboratorAlias string   `json:"reviewer_collaborator_alias,omitempty"`
	ForceSendFields           []string `json:"-" tf:"-"`
}

func (st CleanRoomNotebookReview) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomNotebookReviewToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomNotebookReview) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomNotebookReviewPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomNotebookReviewFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomNotebookReviewToPb(st *CleanRoomNotebookReview) (*cleanroomspb.CleanRoomNotebookReviewPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomNotebookReviewPb{}
	pb.Comment = st.Comment
	pb.CreatedAtMillis = st.CreatedAtMillis
	reviewStatePb, err := CleanRoomNotebookReviewNotebookReviewStateToPb(&st.ReviewState)
	if err != nil {
		return nil, err
	}
	if reviewStatePb != nil {
		pb.ReviewState = *reviewStatePb
	}
	reviewSubReasonPb, err := CleanRoomNotebookReviewNotebookReviewSubReasonToPb(&st.ReviewSubReason)
	if err != nil {
		return nil, err
	}
	if reviewSubReasonPb != nil {
		pb.ReviewSubReason = *reviewSubReasonPb
	}
	pb.ReviewerCollaboratorAlias = st.ReviewerCollaboratorAlias

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomNotebookReviewFromPb(pb *cleanroomspb.CleanRoomNotebookReviewPb) (*CleanRoomNotebookReview, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomNotebookReview{}
	st.Comment = pb.Comment
	st.CreatedAtMillis = pb.CreatedAtMillis
	reviewStateField, err := CleanRoomNotebookReviewNotebookReviewStateFromPb(&pb.ReviewState)
	if err != nil {
		return nil, err
	}
	if reviewStateField != nil {
		st.ReviewState = *reviewStateField
	}
	reviewSubReasonField, err := CleanRoomNotebookReviewNotebookReviewSubReasonFromPb(&pb.ReviewSubReason)
	if err != nil {
		return nil, err
	}
	if reviewSubReasonField != nil {
		st.ReviewSubReason = *reviewSubReasonField
	}
	st.ReviewerCollaboratorAlias = pb.ReviewerCollaboratorAlias

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CleanRoomNotebookReviewNotebookReviewState string

const CleanRoomNotebookReviewNotebookReviewStateApproved CleanRoomNotebookReviewNotebookReviewState = `APPROVED`

const CleanRoomNotebookReviewNotebookReviewStatePending CleanRoomNotebookReviewNotebookReviewState = `PENDING`

const CleanRoomNotebookReviewNotebookReviewStateRejected CleanRoomNotebookReviewNotebookReviewState = `REJECTED`

// String representation for [fmt.Print]
func (f *CleanRoomNotebookReviewNotebookReviewState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomNotebookReviewNotebookReviewState) Set(v string) error {
	switch v {
	case `APPROVED`, `PENDING`, `REJECTED`:
		*f = CleanRoomNotebookReviewNotebookReviewState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVED", "PENDING", "REJECTED"`, v)
	}
}

// Values returns all possible values for CleanRoomNotebookReviewNotebookReviewState.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomNotebookReviewNotebookReviewState) Values() []CleanRoomNotebookReviewNotebookReviewState {
	return []CleanRoomNotebookReviewNotebookReviewState{
		CleanRoomNotebookReviewNotebookReviewStateApproved,
		CleanRoomNotebookReviewNotebookReviewStatePending,
		CleanRoomNotebookReviewNotebookReviewStateRejected,
	}
}

// Type always returns CleanRoomNotebookReviewNotebookReviewState to satisfy [pflag.Value] interface
func (f *CleanRoomNotebookReviewNotebookReviewState) Type() string {
	return "CleanRoomNotebookReviewNotebookReviewState"
}

func CleanRoomNotebookReviewNotebookReviewStateToPb(st *CleanRoomNotebookReviewNotebookReviewState) (*cleanroomspb.CleanRoomNotebookReviewNotebookReviewStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanroomspb.CleanRoomNotebookReviewNotebookReviewStatePb(*st)
	return &pb, nil
}

func CleanRoomNotebookReviewNotebookReviewStateFromPb(pb *cleanroomspb.CleanRoomNotebookReviewNotebookReviewStatePb) (*CleanRoomNotebookReviewNotebookReviewState, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomNotebookReviewNotebookReviewState(*pb)
	return &st, nil
}

type CleanRoomNotebookReviewNotebookReviewSubReason string

const CleanRoomNotebookReviewNotebookReviewSubReasonAutoApproved CleanRoomNotebookReviewNotebookReviewSubReason = `AUTO_APPROVED`

const CleanRoomNotebookReviewNotebookReviewSubReasonBackfilled CleanRoomNotebookReviewNotebookReviewSubReason = `BACKFILLED`

// String representation for [fmt.Print]
func (f *CleanRoomNotebookReviewNotebookReviewSubReason) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomNotebookReviewNotebookReviewSubReason) Set(v string) error {
	switch v {
	case `AUTO_APPROVED`, `BACKFILLED`:
		*f = CleanRoomNotebookReviewNotebookReviewSubReason(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO_APPROVED", "BACKFILLED"`, v)
	}
}

// Values returns all possible values for CleanRoomNotebookReviewNotebookReviewSubReason.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomNotebookReviewNotebookReviewSubReason) Values() []CleanRoomNotebookReviewNotebookReviewSubReason {
	return []CleanRoomNotebookReviewNotebookReviewSubReason{
		CleanRoomNotebookReviewNotebookReviewSubReasonAutoApproved,
		CleanRoomNotebookReviewNotebookReviewSubReasonBackfilled,
	}
}

// Type always returns CleanRoomNotebookReviewNotebookReviewSubReason to satisfy [pflag.Value] interface
func (f *CleanRoomNotebookReviewNotebookReviewSubReason) Type() string {
	return "CleanRoomNotebookReviewNotebookReviewSubReason"
}

func CleanRoomNotebookReviewNotebookReviewSubReasonToPb(st *CleanRoomNotebookReviewNotebookReviewSubReason) (*cleanroomspb.CleanRoomNotebookReviewNotebookReviewSubReasonPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanroomspb.CleanRoomNotebookReviewNotebookReviewSubReasonPb(*st)
	return &pb, nil
}

func CleanRoomNotebookReviewNotebookReviewSubReasonFromPb(pb *cleanroomspb.CleanRoomNotebookReviewNotebookReviewSubReasonPb) (*CleanRoomNotebookReviewNotebookReviewSubReason, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomNotebookReviewNotebookReviewSubReason(*pb)
	return &st, nil
}

// Stores information about a single task run.
type CleanRoomNotebookTaskRun struct {
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API. if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	// Wire name: 'collaborator_job_run_info'
	CollaboratorJobRunInfo *CollaboratorJobRunInfo `json:"collaborator_job_run_info,omitempty"`
	// Etag of the notebook executed in this task run, used to identify the
	// notebook version.
	// Wire name: 'notebook_etag'
	NotebookEtag string `json:"notebook_etag,omitempty"`
	// State of the task run.
	// Wire name: 'notebook_job_run_state'
	NotebookJobRunState *jobs.CleanRoomTaskRunState `json:"notebook_job_run_state,omitempty"`
	// Asset name of the notebook executed in this task run.
	// Wire name: 'notebook_name'
	NotebookName string `json:"notebook_name,omitempty"`
	// The timestamp of when the notebook was last updated.
	// Wire name: 'notebook_updated_at'
	NotebookUpdatedAt int64 `json:"notebook_updated_at,omitempty"`
	// Expiration time of the output schema of the task run (if any), in epoch
	// milliseconds.
	// Wire name: 'output_schema_expiration_time'
	OutputSchemaExpirationTime int64 `json:"output_schema_expiration_time,omitempty"`
	// Name of the output schema associated with the clean rooms notebook task
	// run.
	// Wire name: 'output_schema_name'
	OutputSchemaName string `json:"output_schema_name,omitempty"`
	// Duration of the task run, in milliseconds.
	// Wire name: 'run_duration'
	RunDuration int64 `json:"run_duration,omitempty"`
	// When the task run started, in epoch milliseconds.
	// Wire name: 'start_time'
	StartTime       int64    `json:"start_time,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomNotebookTaskRun) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomNotebookTaskRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomNotebookTaskRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomNotebookTaskRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomNotebookTaskRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomNotebookTaskRunToPb(st *CleanRoomNotebookTaskRun) (*cleanroomspb.CleanRoomNotebookTaskRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomNotebookTaskRunPb{}
	collaboratorJobRunInfoPb, err := CollaboratorJobRunInfoToPb(st.CollaboratorJobRunInfo)
	if err != nil {
		return nil, err
	}
	if collaboratorJobRunInfoPb != nil {
		pb.CollaboratorJobRunInfo = collaboratorJobRunInfoPb
	}
	pb.NotebookEtag = st.NotebookEtag
	notebookJobRunStatePb, err := jobs.CleanRoomTaskRunStateToPb(st.NotebookJobRunState)
	if err != nil {
		return nil, err
	}
	if notebookJobRunStatePb != nil {
		pb.NotebookJobRunState = notebookJobRunStatePb
	}
	pb.NotebookName = st.NotebookName
	pb.NotebookUpdatedAt = st.NotebookUpdatedAt
	pb.OutputSchemaExpirationTime = st.OutputSchemaExpirationTime
	pb.OutputSchemaName = st.OutputSchemaName
	pb.RunDuration = st.RunDuration
	pb.StartTime = st.StartTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomNotebookTaskRunFromPb(pb *cleanroomspb.CleanRoomNotebookTaskRunPb) (*CleanRoomNotebookTaskRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomNotebookTaskRun{}
	collaboratorJobRunInfoField, err := CollaboratorJobRunInfoFromPb(pb.CollaboratorJobRunInfo)
	if err != nil {
		return nil, err
	}
	if collaboratorJobRunInfoField != nil {
		st.CollaboratorJobRunInfo = collaboratorJobRunInfoField
	}
	st.NotebookEtag = pb.NotebookEtag
	notebookJobRunStateField, err := jobs.CleanRoomTaskRunStateFromPb(pb.NotebookJobRunState)
	if err != nil {
		return nil, err
	}
	if notebookJobRunStateField != nil {
		st.NotebookJobRunState = notebookJobRunStateField
	}
	st.NotebookName = pb.NotebookName
	st.NotebookUpdatedAt = pb.NotebookUpdatedAt
	st.OutputSchemaExpirationTime = pb.OutputSchemaExpirationTime
	st.OutputSchemaName = pb.OutputSchemaName
	st.RunDuration = pb.RunDuration
	st.StartTime = pb.StartTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`

	// Wire name: 'status'
	Status          CleanRoomOutputCatalogOutputCatalogStatus `json:"status,omitempty"`
	ForceSendFields []string                                  `json:"-" tf:"-"`
}

func (st CleanRoomOutputCatalog) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomOutputCatalogToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomOutputCatalog) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomOutputCatalogPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomOutputCatalogFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomOutputCatalogToPb(st *CleanRoomOutputCatalog) (*cleanroomspb.CleanRoomOutputCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomOutputCatalogPb{}
	pb.CatalogName = st.CatalogName
	statusPb, err := CleanRoomOutputCatalogOutputCatalogStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomOutputCatalogFromPb(pb *cleanroomspb.CleanRoomOutputCatalogPb) (*CleanRoomOutputCatalog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomOutputCatalog{}
	st.CatalogName = pb.CatalogName
	statusField, err := CleanRoomOutputCatalogOutputCatalogStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

// Values returns all possible values for CleanRoomOutputCatalogOutputCatalogStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomOutputCatalogOutputCatalogStatus) Values() []CleanRoomOutputCatalogOutputCatalogStatus {
	return []CleanRoomOutputCatalogOutputCatalogStatus{
		CleanRoomOutputCatalogOutputCatalogStatusCreated,
		CleanRoomOutputCatalogOutputCatalogStatusNotCreated,
		CleanRoomOutputCatalogOutputCatalogStatusNotEligible,
	}
}

// Type always returns CleanRoomOutputCatalogOutputCatalogStatus to satisfy [pflag.Value] interface
func (f *CleanRoomOutputCatalogOutputCatalogStatus) Type() string {
	return "CleanRoomOutputCatalogOutputCatalogStatus"
}

func CleanRoomOutputCatalogOutputCatalogStatusToPb(st *CleanRoomOutputCatalogOutputCatalogStatus) (*cleanroomspb.CleanRoomOutputCatalogOutputCatalogStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanroomspb.CleanRoomOutputCatalogOutputCatalogStatusPb(*st)
	return &pb, nil
}

func CleanRoomOutputCatalogOutputCatalogStatusFromPb(pb *cleanroomspb.CleanRoomOutputCatalogOutputCatalogStatusPb) (*CleanRoomOutputCatalogOutputCatalogStatus, error) {
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
	CentralCleanRoomId string `json:"central_clean_room_id,omitempty"`
	// Cloud vendor (aws,azure,gcp) of the central clean room.
	// Wire name: 'cloud_vendor'
	CloudVendor string `json:"cloud_vendor,omitempty"`
	// Collaborators in the central clean room. There should one and only one
	// collaborator in the list that satisfies the owner condition:
	//
	// 1. It has the creator's global_metastore_id (determined by caller of
	// CreateCleanRoom).
	//
	// 2. Its invite_recipient_email is empty.
	// Wire name: 'collaborators'
	Collaborators []CleanRoomCollaborator `json:"collaborators,omitempty"`

	// Wire name: 'compliance_security_profile'
	ComplianceSecurityProfile *ComplianceSecurityProfile `json:"compliance_security_profile,omitempty"`
	// Collaborator who creates the clean room.
	// Wire name: 'creator'
	Creator *CleanRoomCollaborator `json:"creator,omitempty"`
	// Egress network policy to apply to the central clean room workspace.
	// Wire name: 'egress_network_policy'
	EgressNetworkPolicy *settings.EgressNetworkPolicy `json:"egress_network_policy,omitempty"`
	// Region of the central clean room.
	// Wire name: 'region'
	Region          string   `json:"region,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomRemoteDetail) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomRemoteDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomRemoteDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CleanRoomRemoteDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomRemoteDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomRemoteDetailToPb(st *CleanRoomRemoteDetail) (*cleanroomspb.CleanRoomRemoteDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CleanRoomRemoteDetailPb{}
	pb.CentralCleanRoomId = st.CentralCleanRoomId
	pb.CloudVendor = st.CloudVendor

	var collaboratorsPb []cleanroomspb.CleanRoomCollaboratorPb
	for _, item := range st.Collaborators {
		itemPb, err := CleanRoomCollaboratorToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			collaboratorsPb = append(collaboratorsPb, *itemPb)
		}
	}
	pb.Collaborators = collaboratorsPb
	complianceSecurityProfilePb, err := ComplianceSecurityProfileToPb(st.ComplianceSecurityProfile)
	if err != nil {
		return nil, err
	}
	if complianceSecurityProfilePb != nil {
		pb.ComplianceSecurityProfile = complianceSecurityProfilePb
	}
	creatorPb, err := CleanRoomCollaboratorToPb(st.Creator)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomRemoteDetailFromPb(pb *cleanroomspb.CleanRoomRemoteDetailPb) (*CleanRoomRemoteDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomRemoteDetail{}
	st.CentralCleanRoomId = pb.CentralCleanRoomId
	st.CloudVendor = pb.CloudVendor

	var collaboratorsField []CleanRoomCollaborator
	for _, itemPb := range pb.Collaborators {
		item, err := CleanRoomCollaboratorFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			collaboratorsField = append(collaboratorsField, *item)
		}
	}
	st.Collaborators = collaboratorsField
	complianceSecurityProfileField, err := ComplianceSecurityProfileFromPb(pb.ComplianceSecurityProfile)
	if err != nil {
		return nil, err
	}
	if complianceSecurityProfileField != nil {
		st.ComplianceSecurityProfile = complianceSecurityProfileField
	}
	creatorField, err := CleanRoomCollaboratorFromPb(pb.Creator)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

// Values returns all possible values for CleanRoomStatusEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomStatusEnum) Values() []CleanRoomStatusEnum {
	return []CleanRoomStatusEnum{
		CleanRoomStatusEnumActive,
		CleanRoomStatusEnumDeleted,
		CleanRoomStatusEnumFailed,
		CleanRoomStatusEnumProvisioning,
	}
}

// Type always returns CleanRoomStatusEnum to satisfy [pflag.Value] interface
func (f *CleanRoomStatusEnum) Type() string {
	return "CleanRoomStatusEnum"
}

func CleanRoomStatusEnumToPb(st *CleanRoomStatusEnum) (*cleanroomspb.CleanRoomStatusEnumPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanroomspb.CleanRoomStatusEnumPb(*st)
	return &pb, nil
}

func CleanRoomStatusEnumFromPb(pb *cleanroomspb.CleanRoomStatusEnumPb) (*CleanRoomStatusEnum, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomStatusEnum(*pb)
	return &st, nil
}

type CollaboratorJobRunInfo struct {
	// Alias of the collaborator that triggered the task run.
	// Wire name: 'collaborator_alias'
	CollaboratorAlias string `json:"collaborator_alias,omitempty"`
	// Job ID of the task run in the collaborator's workspace.
	// Wire name: 'collaborator_job_id'
	CollaboratorJobId int64 `json:"collaborator_job_id,omitempty"`
	// Job run ID of the task run in the collaborator's workspace.
	// Wire name: 'collaborator_job_run_id'
	CollaboratorJobRunId int64 `json:"collaborator_job_run_id,omitempty"`
	// Task run ID of the task run in the collaborator's workspace.
	// Wire name: 'collaborator_task_run_id'
	CollaboratorTaskRunId int64 `json:"collaborator_task_run_id,omitempty"`
	// ID of the collaborator's workspace that triggered the task run.
	// Wire name: 'collaborator_workspace_id'
	CollaboratorWorkspaceId int64    `json:"collaborator_workspace_id,omitempty"`
	ForceSendFields         []string `json:"-" tf:"-"`
}

func (st CollaboratorJobRunInfo) MarshalJSON() ([]byte, error) {
	pb, err := CollaboratorJobRunInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CollaboratorJobRunInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CollaboratorJobRunInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CollaboratorJobRunInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CollaboratorJobRunInfoToPb(st *CollaboratorJobRunInfo) (*cleanroomspb.CollaboratorJobRunInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CollaboratorJobRunInfoPb{}
	pb.CollaboratorAlias = st.CollaboratorAlias
	pb.CollaboratorJobId = st.CollaboratorJobId
	pb.CollaboratorJobRunId = st.CollaboratorJobRunId
	pb.CollaboratorTaskRunId = st.CollaboratorTaskRunId
	pb.CollaboratorWorkspaceId = st.CollaboratorWorkspaceId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CollaboratorJobRunInfoFromPb(pb *cleanroomspb.CollaboratorJobRunInfoPb) (*CollaboratorJobRunInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CollaboratorJobRunInfo{}
	st.CollaboratorAlias = pb.CollaboratorAlias
	st.CollaboratorJobId = pb.CollaboratorJobId
	st.CollaboratorJobRunId = pb.CollaboratorJobRunId
	st.CollaboratorTaskRunId = pb.CollaboratorTaskRunId
	st.CollaboratorWorkspaceId = pb.CollaboratorWorkspaceId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The compliance security profile used to process regulated data following
// compliance standards.
type ComplianceSecurityProfile struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	// Wire name: 'compliance_standards'
	ComplianceStandards []settings.ComplianceStandard `json:"compliance_standards,omitempty"`
	// Whether the compliance security profile is enabled.
	// Wire name: 'is_enabled'
	IsEnabled       bool     `json:"is_enabled,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	pb, err := ComplianceSecurityProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ComplianceSecurityProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ComplianceSecurityProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ComplianceSecurityProfileToPb(st *ComplianceSecurityProfile) (*cleanroomspb.ComplianceSecurityProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ComplianceSecurityProfilePb{}

	var complianceStandardsPb []settingspb.ComplianceStandardPb
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ComplianceSecurityProfileFromPb(pb *cleanroomspb.ComplianceSecurityProfilePb) (*ComplianceSecurityProfile, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateCleanRoomAssetRequest struct {

	// Wire name: 'asset'
	Asset CleanRoomAsset `json:"asset"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName string `json:"-" tf:"-"`
}

func (st CreateCleanRoomAssetRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateCleanRoomAssetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCleanRoomAssetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CreateCleanRoomAssetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCleanRoomAssetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCleanRoomAssetRequestToPb(st *CreateCleanRoomAssetRequest) (*cleanroomspb.CreateCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CreateCleanRoomAssetRequestPb{}
	assetPb, err := CleanRoomAssetToPb(&st.Asset)
	if err != nil {
		return nil, err
	}
	if assetPb != nil {
		pb.Asset = *assetPb
	}
	pb.CleanRoomName = st.CleanRoomName

	return pb, nil
}

func CreateCleanRoomAssetRequestFromPb(pb *cleanroomspb.CreateCleanRoomAssetRequestPb) (*CreateCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomAssetRequest{}
	assetField, err := CleanRoomAssetFromPb(&pb.Asset)
	if err != nil {
		return nil, err
	}
	if assetField != nil {
		st.Asset = *assetField
	}
	st.CleanRoomName = pb.CleanRoomName

	return st, nil
}

type CreateCleanRoomAssetReviewRequest struct {
	// can only be NOTEBOOK_FILE for now
	AssetType CleanRoomAssetAssetType `json:"-" tf:"-"`
	// Name of the clean room
	CleanRoomName string `json:"-" tf:"-"`
	// Name of the asset
	Name string `json:"-" tf:"-"`

	// Wire name: 'notebook_review'
	NotebookReview NotebookVersionReview `json:"notebook_review"`
}

func (st CreateCleanRoomAssetReviewRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateCleanRoomAssetReviewRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCleanRoomAssetReviewRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CreateCleanRoomAssetReviewRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCleanRoomAssetReviewRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCleanRoomAssetReviewRequestToPb(st *CreateCleanRoomAssetReviewRequest) (*cleanroomspb.CreateCleanRoomAssetReviewRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CreateCleanRoomAssetReviewRequestPb{}
	assetTypePb, err := CleanRoomAssetAssetTypeToPb(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}
	pb.CleanRoomName = st.CleanRoomName
	pb.Name = st.Name
	notebookReviewPb, err := NotebookVersionReviewToPb(&st.NotebookReview)
	if err != nil {
		return nil, err
	}
	if notebookReviewPb != nil {
		pb.NotebookReview = *notebookReviewPb
	}

	return pb, nil
}

func CreateCleanRoomAssetReviewRequestFromPb(pb *cleanroomspb.CreateCleanRoomAssetReviewRequestPb) (*CreateCleanRoomAssetReviewRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomAssetReviewRequest{}
	assetTypeField, err := CleanRoomAssetAssetTypeFromPb(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name
	notebookReviewField, err := NotebookVersionReviewFromPb(&pb.NotebookReview)
	if err != nil {
		return nil, err
	}
	if notebookReviewField != nil {
		st.NotebookReview = *notebookReviewField
	}

	return st, nil
}

type CreateCleanRoomAssetReviewResponse struct {
	// top-level status derived from all reviews
	// Wire name: 'notebook_review_state'
	NotebookReviewState CleanRoomNotebookReviewNotebookReviewState `json:"notebook_review_state,omitempty"`
	// All existing notebook approvals or rejections
	// Wire name: 'notebook_reviews'
	NotebookReviews []CleanRoomNotebookReview `json:"notebook_reviews,omitempty"`
}

func (st CreateCleanRoomAssetReviewResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateCleanRoomAssetReviewResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCleanRoomAssetReviewResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CreateCleanRoomAssetReviewResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCleanRoomAssetReviewResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCleanRoomAssetReviewResponseToPb(st *CreateCleanRoomAssetReviewResponse) (*cleanroomspb.CreateCleanRoomAssetReviewResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CreateCleanRoomAssetReviewResponsePb{}
	notebookReviewStatePb, err := CleanRoomNotebookReviewNotebookReviewStateToPb(&st.NotebookReviewState)
	if err != nil {
		return nil, err
	}
	if notebookReviewStatePb != nil {
		pb.NotebookReviewState = *notebookReviewStatePb
	}

	var notebookReviewsPb []cleanroomspb.CleanRoomNotebookReviewPb
	for _, item := range st.NotebookReviews {
		itemPb, err := CleanRoomNotebookReviewToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notebookReviewsPb = append(notebookReviewsPb, *itemPb)
		}
	}
	pb.NotebookReviews = notebookReviewsPb

	return pb, nil
}

func CreateCleanRoomAssetReviewResponseFromPb(pb *cleanroomspb.CreateCleanRoomAssetReviewResponsePb) (*CreateCleanRoomAssetReviewResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomAssetReviewResponse{}
	notebookReviewStateField, err := CleanRoomNotebookReviewNotebookReviewStateFromPb(&pb.NotebookReviewState)
	if err != nil {
		return nil, err
	}
	if notebookReviewStateField != nil {
		st.NotebookReviewState = *notebookReviewStateField
	}

	var notebookReviewsField []CleanRoomNotebookReview
	for _, itemPb := range pb.NotebookReviews {
		item, err := CleanRoomNotebookReviewFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notebookReviewsField = append(notebookReviewsField, *item)
		}
	}
	st.NotebookReviews = notebookReviewsField

	return st, nil
}

type CreateCleanRoomAutoApprovalRuleRequest struct {

	// Wire name: 'auto_approval_rule'
	AutoApprovalRule CleanRoomAutoApprovalRule `json:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName string `json:"-" tf:"-"`
}

func (st CreateCleanRoomAutoApprovalRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateCleanRoomAutoApprovalRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCleanRoomAutoApprovalRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CreateCleanRoomAutoApprovalRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCleanRoomAutoApprovalRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCleanRoomAutoApprovalRuleRequestToPb(st *CreateCleanRoomAutoApprovalRuleRequest) (*cleanroomspb.CreateCleanRoomAutoApprovalRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CreateCleanRoomAutoApprovalRuleRequestPb{}
	autoApprovalRulePb, err := CleanRoomAutoApprovalRuleToPb(&st.AutoApprovalRule)
	if err != nil {
		return nil, err
	}
	if autoApprovalRulePb != nil {
		pb.AutoApprovalRule = *autoApprovalRulePb
	}
	pb.CleanRoomName = st.CleanRoomName

	return pb, nil
}

func CreateCleanRoomAutoApprovalRuleRequestFromPb(pb *cleanroomspb.CreateCleanRoomAutoApprovalRuleRequestPb) (*CreateCleanRoomAutoApprovalRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomAutoApprovalRuleRequest{}
	autoApprovalRuleField, err := CleanRoomAutoApprovalRuleFromPb(&pb.AutoApprovalRule)
	if err != nil {
		return nil, err
	}
	if autoApprovalRuleField != nil {
		st.AutoApprovalRule = *autoApprovalRuleField
	}
	st.CleanRoomName = pb.CleanRoomName

	return st, nil
}

type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`

	// Wire name: 'output_catalog'
	OutputCatalog CleanRoomOutputCatalog `json:"output_catalog"`
}

func (st CreateCleanRoomOutputCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateCleanRoomOutputCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCleanRoomOutputCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CreateCleanRoomOutputCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCleanRoomOutputCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCleanRoomOutputCatalogRequestToPb(st *CreateCleanRoomOutputCatalogRequest) (*cleanroomspb.CreateCleanRoomOutputCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CreateCleanRoomOutputCatalogRequestPb{}
	pb.CleanRoomName = st.CleanRoomName
	outputCatalogPb, err := CleanRoomOutputCatalogToPb(&st.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogPb != nil {
		pb.OutputCatalog = *outputCatalogPb
	}

	return pb, nil
}

func CreateCleanRoomOutputCatalogRequestFromPb(pb *cleanroomspb.CreateCleanRoomOutputCatalogRequestPb) (*CreateCleanRoomOutputCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomOutputCatalogRequest{}
	st.CleanRoomName = pb.CleanRoomName
	outputCatalogField, err := CleanRoomOutputCatalogFromPb(&pb.OutputCatalog)
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
	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
}

func (st CreateCleanRoomOutputCatalogResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateCleanRoomOutputCatalogResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCleanRoomOutputCatalogResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CreateCleanRoomOutputCatalogResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCleanRoomOutputCatalogResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCleanRoomOutputCatalogResponseToPb(st *CreateCleanRoomOutputCatalogResponse) (*cleanroomspb.CreateCleanRoomOutputCatalogResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CreateCleanRoomOutputCatalogResponsePb{}
	outputCatalogPb, err := CleanRoomOutputCatalogToPb(st.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogPb != nil {
		pb.OutputCatalog = outputCatalogPb
	}

	return pb, nil
}

func CreateCleanRoomOutputCatalogResponseFromPb(pb *cleanroomspb.CreateCleanRoomOutputCatalogResponsePb) (*CreateCleanRoomOutputCatalogResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomOutputCatalogResponse{}
	outputCatalogField, err := CleanRoomOutputCatalogFromPb(pb.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogField != nil {
		st.OutputCatalog = outputCatalogField
	}

	return st, nil
}

type CreateCleanRoomRequest struct {

	// Wire name: 'clean_room'
	CleanRoom CleanRoom `json:"clean_room"`
}

func (st CreateCleanRoomRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateCleanRoomRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCleanRoomRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.CreateCleanRoomRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCleanRoomRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCleanRoomRequestToPb(st *CreateCleanRoomRequest) (*cleanroomspb.CreateCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.CreateCleanRoomRequestPb{}
	cleanRoomPb, err := CleanRoomToPb(&st.CleanRoom)
	if err != nil {
		return nil, err
	}
	if cleanRoomPb != nil {
		pb.CleanRoom = *cleanRoomPb
	}

	return pb, nil
}

func CreateCleanRoomRequestFromPb(pb *cleanroomspb.CreateCleanRoomRequestPb) (*CreateCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomRequest{}
	cleanRoomField, err := CleanRoomFromPb(&pb.CleanRoom)
	if err != nil {
		return nil, err
	}
	if cleanRoomField != nil {
		st.CleanRoom = *cleanRoomField
	}

	return st, nil
}

type DeleteCleanRoomAssetRequest struct {
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" tf:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name string `json:"-" tf:"-"`
}

func (st DeleteCleanRoomAssetRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteCleanRoomAssetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteCleanRoomAssetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.DeleteCleanRoomAssetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteCleanRoomAssetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteCleanRoomAssetRequestToPb(st *DeleteCleanRoomAssetRequest) (*cleanroomspb.DeleteCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.DeleteCleanRoomAssetRequestPb{}
	assetTypePb, err := CleanRoomAssetAssetTypeToPb(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}
	pb.CleanRoomName = st.CleanRoomName
	pb.Name = st.Name

	return pb, nil
}

func DeleteCleanRoomAssetRequestFromPb(pb *cleanroomspb.DeleteCleanRoomAssetRequestPb) (*DeleteCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCleanRoomAssetRequest{}
	assetTypeField, err := CleanRoomAssetAssetTypeFromPb(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name

	return st, nil
}

type DeleteCleanRoomAutoApprovalRuleRequest struct {
	CleanRoomName string `json:"-" tf:"-"`

	RuleId string `json:"-" tf:"-"`
}

func (st DeleteCleanRoomAutoApprovalRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteCleanRoomAutoApprovalRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteCleanRoomAutoApprovalRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.DeleteCleanRoomAutoApprovalRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteCleanRoomAutoApprovalRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteCleanRoomAutoApprovalRuleRequestToPb(st *DeleteCleanRoomAutoApprovalRuleRequest) (*cleanroomspb.DeleteCleanRoomAutoApprovalRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.DeleteCleanRoomAutoApprovalRuleRequestPb{}
	pb.CleanRoomName = st.CleanRoomName
	pb.RuleId = st.RuleId

	return pb, nil
}

func DeleteCleanRoomAutoApprovalRuleRequestFromPb(pb *cleanroomspb.DeleteCleanRoomAutoApprovalRuleRequestPb) (*DeleteCleanRoomAutoApprovalRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCleanRoomAutoApprovalRuleRequest{}
	st.CleanRoomName = pb.CleanRoomName
	st.RuleId = pb.RuleId

	return st, nil
}

type DeleteCleanRoomRequest struct {
	// Name of the clean room.
	Name string `json:"-" tf:"-"`
}

func (st DeleteCleanRoomRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteCleanRoomRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteCleanRoomRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.DeleteCleanRoomRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteCleanRoomRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteCleanRoomRequestToPb(st *DeleteCleanRoomRequest) (*cleanroomspb.DeleteCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.DeleteCleanRoomRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteCleanRoomRequestFromPb(pb *cleanroomspb.DeleteCleanRoomRequestPb) (*DeleteCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCleanRoomRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetCleanRoomAssetRequest struct {
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" tf:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name string `json:"-" tf:"-"`
}

func (st GetCleanRoomAssetRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetCleanRoomAssetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCleanRoomAssetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.GetCleanRoomAssetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCleanRoomAssetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCleanRoomAssetRequestToPb(st *GetCleanRoomAssetRequest) (*cleanroomspb.GetCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.GetCleanRoomAssetRequestPb{}
	assetTypePb, err := CleanRoomAssetAssetTypeToPb(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}
	pb.CleanRoomName = st.CleanRoomName
	pb.Name = st.Name

	return pb, nil
}

func GetCleanRoomAssetRequestFromPb(pb *cleanroomspb.GetCleanRoomAssetRequestPb) (*GetCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCleanRoomAssetRequest{}
	assetTypeField, err := CleanRoomAssetAssetTypeFromPb(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name

	return st, nil
}

type GetCleanRoomAssetRevisionRequest struct {
	// Asset type. Only NOTEBOOK_FILE is supported.
	AssetType CleanRoomAssetAssetType `json:"-" tf:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// Revision etag to fetch. If not provided, the latest revision will be
	// returned.
	Etag string `json:"-" tf:"-"`
	// Name of the asset.
	Name string `json:"-" tf:"-"`
}

func (st GetCleanRoomAssetRevisionRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetCleanRoomAssetRevisionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCleanRoomAssetRevisionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.GetCleanRoomAssetRevisionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCleanRoomAssetRevisionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCleanRoomAssetRevisionRequestToPb(st *GetCleanRoomAssetRevisionRequest) (*cleanroomspb.GetCleanRoomAssetRevisionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.GetCleanRoomAssetRevisionRequestPb{}
	assetTypePb, err := CleanRoomAssetAssetTypeToPb(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}
	pb.CleanRoomName = st.CleanRoomName
	pb.Etag = st.Etag
	pb.Name = st.Name

	return pb, nil
}

func GetCleanRoomAssetRevisionRequestFromPb(pb *cleanroomspb.GetCleanRoomAssetRevisionRequestPb) (*GetCleanRoomAssetRevisionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCleanRoomAssetRevisionRequest{}
	assetTypeField, err := CleanRoomAssetAssetTypeFromPb(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	st.CleanRoomName = pb.CleanRoomName
	st.Etag = pb.Etag
	st.Name = pb.Name

	return st, nil
}

type GetCleanRoomAutoApprovalRuleRequest struct {
	CleanRoomName string `json:"-" tf:"-"`

	RuleId string `json:"-" tf:"-"`
}

func (st GetCleanRoomAutoApprovalRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetCleanRoomAutoApprovalRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCleanRoomAutoApprovalRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.GetCleanRoomAutoApprovalRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCleanRoomAutoApprovalRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCleanRoomAutoApprovalRuleRequestToPb(st *GetCleanRoomAutoApprovalRuleRequest) (*cleanroomspb.GetCleanRoomAutoApprovalRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.GetCleanRoomAutoApprovalRuleRequestPb{}
	pb.CleanRoomName = st.CleanRoomName
	pb.RuleId = st.RuleId

	return pb, nil
}

func GetCleanRoomAutoApprovalRuleRequestFromPb(pb *cleanroomspb.GetCleanRoomAutoApprovalRuleRequestPb) (*GetCleanRoomAutoApprovalRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCleanRoomAutoApprovalRuleRequest{}
	st.CleanRoomName = pb.CleanRoomName
	st.RuleId = pb.RuleId

	return st, nil
}

type GetCleanRoomRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st GetCleanRoomRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetCleanRoomRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCleanRoomRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.GetCleanRoomRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCleanRoomRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCleanRoomRequestToPb(st *GetCleanRoomRequest) (*cleanroomspb.GetCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.GetCleanRoomRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetCleanRoomRequestFromPb(pb *cleanroomspb.GetCleanRoomRequestPb) (*GetCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCleanRoomRequest{}
	st.Name = pb.Name

	return st, nil
}

type ListCleanRoomAssetRevisionsRequest struct {
	// Asset type. Only NOTEBOOK_FILE is supported.
	AssetType CleanRoomAssetAssetType `json:"-" tf:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// Name of the asset.
	Name string `json:"-" tf:"-"`
	// Maximum number of asset revisions to return. Defaults to 10.
	PageSize int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on the previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomAssetRevisionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomAssetRevisionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomAssetRevisionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomAssetRevisionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomAssetRevisionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomAssetRevisionsRequestToPb(st *ListCleanRoomAssetRevisionsRequest) (*cleanroomspb.ListCleanRoomAssetRevisionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomAssetRevisionsRequestPb{}
	assetTypePb, err := CleanRoomAssetAssetTypeToPb(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}
	pb.CleanRoomName = st.CleanRoomName
	pb.Name = st.Name
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomAssetRevisionsRequestFromPb(pb *cleanroomspb.ListCleanRoomAssetRevisionsRequestPb) (*ListCleanRoomAssetRevisionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAssetRevisionsRequest{}
	assetTypeField, err := CleanRoomAssetAssetTypeFromPb(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomAssetRevisionsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'revisions'
	Revisions       []CleanRoomAsset `json:"revisions,omitempty"`
	ForceSendFields []string         `json:"-" tf:"-"`
}

func (st ListCleanRoomAssetRevisionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomAssetRevisionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomAssetRevisionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomAssetRevisionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomAssetRevisionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomAssetRevisionsResponseToPb(st *ListCleanRoomAssetRevisionsResponse) (*cleanroomspb.ListCleanRoomAssetRevisionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomAssetRevisionsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var revisionsPb []cleanroomspb.CleanRoomAssetPb
	for _, item := range st.Revisions {
		itemPb, err := CleanRoomAssetToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			revisionsPb = append(revisionsPb, *itemPb)
		}
	}
	pb.Revisions = revisionsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomAssetRevisionsResponseFromPb(pb *cleanroomspb.ListCleanRoomAssetRevisionsResponsePb) (*ListCleanRoomAssetRevisionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAssetRevisionsResponse{}
	st.NextPageToken = pb.NextPageToken

	var revisionsField []CleanRoomAsset
	for _, itemPb := range pb.Revisions {
		item, err := CleanRoomAssetFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			revisionsField = append(revisionsField, *item)
		}
	}
	st.Revisions = revisionsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomAssetsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomAssetsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomAssetsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomAssetsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomAssetsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomAssetsRequestToPb(st *ListCleanRoomAssetsRequest) (*cleanroomspb.ListCleanRoomAssetsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomAssetsRequestPb{}
	pb.CleanRoomName = st.CleanRoomName
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomAssetsRequestFromPb(pb *cleanroomspb.ListCleanRoomAssetsRequestPb) (*ListCleanRoomAssetsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAssetsRequest{}
	st.CleanRoomName = pb.CleanRoomName
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomAssetsResponse struct {
	// Assets in the clean room.
	// Wire name: 'assets'
	Assets []CleanRoomAsset `json:"assets,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomAssetsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomAssetsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomAssetsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomAssetsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomAssetsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomAssetsResponseToPb(st *ListCleanRoomAssetsResponse) (*cleanroomspb.ListCleanRoomAssetsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomAssetsResponsePb{}

	var assetsPb []cleanroomspb.CleanRoomAssetPb
	for _, item := range st.Assets {
		itemPb, err := CleanRoomAssetToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			assetsPb = append(assetsPb, *itemPb)
		}
	}
	pb.Assets = assetsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomAssetsResponseFromPb(pb *cleanroomspb.ListCleanRoomAssetsResponsePb) (*ListCleanRoomAssetsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAssetsResponse{}

	var assetsField []CleanRoomAsset
	for _, itemPb := range pb.Assets {
		item, err := CleanRoomAssetFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			assetsField = append(assetsField, *item)
		}
	}
	st.Assets = assetsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomAutoApprovalRulesRequest struct {
	CleanRoomName string `json:"-" tf:"-"`
	// Maximum number of auto-approval rules to return. Defaults to 100.
	PageSize int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomAutoApprovalRulesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomAutoApprovalRulesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomAutoApprovalRulesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomAutoApprovalRulesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomAutoApprovalRulesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomAutoApprovalRulesRequestToPb(st *ListCleanRoomAutoApprovalRulesRequest) (*cleanroomspb.ListCleanRoomAutoApprovalRulesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomAutoApprovalRulesRequestPb{}
	pb.CleanRoomName = st.CleanRoomName
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomAutoApprovalRulesRequestFromPb(pb *cleanroomspb.ListCleanRoomAutoApprovalRulesRequestPb) (*ListCleanRoomAutoApprovalRulesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAutoApprovalRulesRequest{}
	st.CleanRoomName = pb.CleanRoomName
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomAutoApprovalRulesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'rules'
	Rules           []CleanRoomAutoApprovalRule `json:"rules,omitempty"`
	ForceSendFields []string                    `json:"-" tf:"-"`
}

func (st ListCleanRoomAutoApprovalRulesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomAutoApprovalRulesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomAutoApprovalRulesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomAutoApprovalRulesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomAutoApprovalRulesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomAutoApprovalRulesResponseToPb(st *ListCleanRoomAutoApprovalRulesResponse) (*cleanroomspb.ListCleanRoomAutoApprovalRulesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomAutoApprovalRulesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var rulesPb []cleanroomspb.CleanRoomAutoApprovalRulePb
	for _, item := range st.Rules {
		itemPb, err := CleanRoomAutoApprovalRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rulesPb = append(rulesPb, *itemPb)
		}
	}
	pb.Rules = rulesPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomAutoApprovalRulesResponseFromPb(pb *cleanroomspb.ListCleanRoomAutoApprovalRulesResponsePb) (*ListCleanRoomAutoApprovalRulesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAutoApprovalRulesResponse{}
	st.NextPageToken = pb.NextPageToken

	var rulesField []CleanRoomAutoApprovalRule
	for _, itemPb := range pb.Rules {
		item, err := CleanRoomAutoApprovalRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rulesField = append(rulesField, *item)
		}
	}
	st.Rules = rulesField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomNotebookTaskRunsRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// Notebook name
	NotebookName string `json:"-" tf:"-"`
	// The maximum number of task runs to return. Currently ignored - all runs
	// will be returned.
	PageSize int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomNotebookTaskRunsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomNotebookTaskRunsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomNotebookTaskRunsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomNotebookTaskRunsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomNotebookTaskRunsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomNotebookTaskRunsRequestToPb(st *ListCleanRoomNotebookTaskRunsRequest) (*cleanroomspb.ListCleanRoomNotebookTaskRunsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomNotebookTaskRunsRequestPb{}
	pb.CleanRoomName = st.CleanRoomName
	pb.NotebookName = st.NotebookName
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomNotebookTaskRunsRequestFromPb(pb *cleanroomspb.ListCleanRoomNotebookTaskRunsRequestPb) (*ListCleanRoomNotebookTaskRunsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomNotebookTaskRunsRequest{}
	st.CleanRoomName = pb.CleanRoomName
	st.NotebookName = pb.NotebookName
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomNotebookTaskRunsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// Name of the clean room.
	// Wire name: 'runs'
	Runs            []CleanRoomNotebookTaskRun `json:"runs,omitempty"`
	ForceSendFields []string                   `json:"-" tf:"-"`
}

func (st ListCleanRoomNotebookTaskRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomNotebookTaskRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomNotebookTaskRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomNotebookTaskRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomNotebookTaskRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomNotebookTaskRunsResponseToPb(st *ListCleanRoomNotebookTaskRunsResponse) (*cleanroomspb.ListCleanRoomNotebookTaskRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomNotebookTaskRunsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var runsPb []cleanroomspb.CleanRoomNotebookTaskRunPb
	for _, item := range st.Runs {
		itemPb, err := CleanRoomNotebookTaskRunToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			runsPb = append(runsPb, *itemPb)
		}
	}
	pb.Runs = runsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomNotebookTaskRunsResponseFromPb(pb *cleanroomspb.ListCleanRoomNotebookTaskRunsResponsePb) (*ListCleanRoomNotebookTaskRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomNotebookTaskRunsResponse{}
	st.NextPageToken = pb.NextPageToken

	var runsField []CleanRoomNotebookTaskRun
	for _, itemPb := range pb.Runs {
		item, err := CleanRoomNotebookTaskRunFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			runsField = append(runsField, *item)
		}
	}
	st.Runs = runsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomsRequestToPb(st *ListCleanRoomsRequest) (*cleanroomspb.ListCleanRoomsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomsRequestFromPb(pb *cleanroomspb.ListCleanRoomsRequestPb) (*ListCleanRoomsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCleanRoomsResponse struct {

	// Wire name: 'clean_rooms'
	CleanRooms []CleanRoom `json:"clean_rooms,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListCleanRoomsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCleanRoomsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.ListCleanRoomsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCleanRoomsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCleanRoomsResponseToPb(st *ListCleanRoomsResponse) (*cleanroomspb.ListCleanRoomsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.ListCleanRoomsResponsePb{}

	var cleanRoomsPb []cleanroomspb.CleanRoomPb
	for _, item := range st.CleanRooms {
		itemPb, err := CleanRoomToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			cleanRoomsPb = append(cleanRoomsPb, *itemPb)
		}
	}
	pb.CleanRooms = cleanRoomsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCleanRoomsResponseFromPb(pb *cleanroomspb.ListCleanRoomsResponsePb) (*ListCleanRoomsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomsResponse{}

	var cleanRoomsField []CleanRoom
	for _, itemPb := range pb.CleanRooms {
		item, err := CleanRoomFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			cleanRoomsField = append(cleanRoomsField, *item)
		}
	}
	st.CleanRooms = cleanRoomsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type NotebookVersionReview struct {
	// review comment
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// etag that identifies the notebook version
	// Wire name: 'etag'
	Etag string `json:"etag"`
	// review outcome
	// Wire name: 'review_state'
	ReviewState     CleanRoomNotebookReviewNotebookReviewState `json:"review_state"`
	ForceSendFields []string                                   `json:"-" tf:"-"`
}

func (st NotebookVersionReview) MarshalJSON() ([]byte, error) {
	pb, err := NotebookVersionReviewToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NotebookVersionReview) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.NotebookVersionReviewPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NotebookVersionReviewFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NotebookVersionReviewToPb(st *NotebookVersionReview) (*cleanroomspb.NotebookVersionReviewPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.NotebookVersionReviewPb{}
	pb.Comment = st.Comment
	pb.Etag = st.Etag
	reviewStatePb, err := CleanRoomNotebookReviewNotebookReviewStateToPb(&st.ReviewState)
	if err != nil {
		return nil, err
	}
	if reviewStatePb != nil {
		pb.ReviewState = *reviewStatePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NotebookVersionReviewFromPb(pb *cleanroomspb.NotebookVersionReviewPb) (*NotebookVersionReview, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookVersionReview{}
	st.Comment = pb.Comment
	st.Etag = pb.Etag
	reviewStateField, err := CleanRoomNotebookReviewNotebookReviewStateFromPb(&pb.ReviewState)
	if err != nil {
		return nil, err
	}
	if reviewStateField != nil {
		st.ReviewState = *reviewStateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateCleanRoomAssetRequest struct {
	// The asset to update. The asset's `name` and `asset_type` fields are used
	// to identify the asset to update.
	// Wire name: 'asset'
	Asset CleanRoomAsset `json:"asset"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" tf:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name string `json:"-" tf:"-"`
}

func (st UpdateCleanRoomAssetRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCleanRoomAssetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCleanRoomAssetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.UpdateCleanRoomAssetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCleanRoomAssetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateCleanRoomAssetRequestToPb(st *UpdateCleanRoomAssetRequest) (*cleanroomspb.UpdateCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.UpdateCleanRoomAssetRequestPb{}
	assetPb, err := CleanRoomAssetToPb(&st.Asset)
	if err != nil {
		return nil, err
	}
	if assetPb != nil {
		pb.Asset = *assetPb
	}
	assetTypePb, err := CleanRoomAssetAssetTypeToPb(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}
	pb.CleanRoomName = st.CleanRoomName
	pb.Name = st.Name

	return pb, nil
}

func UpdateCleanRoomAssetRequestFromPb(pb *cleanroomspb.UpdateCleanRoomAssetRequestPb) (*UpdateCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCleanRoomAssetRequest{}
	assetField, err := CleanRoomAssetFromPb(&pb.Asset)
	if err != nil {
		return nil, err
	}
	if assetField != nil {
		st.Asset = *assetField
	}
	assetTypeField, err := CleanRoomAssetAssetTypeFromPb(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name

	return st, nil
}

type UpdateCleanRoomAutoApprovalRuleRequest struct {
	// The auto-approval rule to update. The rule_id field is used to identify
	// the rule to update.
	// Wire name: 'auto_approval_rule'
	AutoApprovalRule CleanRoomAutoApprovalRule `json:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName string `json:"-" tf:"-"`
	// A generated UUID identifying the rule.
	RuleId string `json:"-" tf:"-"`
}

func (st UpdateCleanRoomAutoApprovalRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCleanRoomAutoApprovalRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCleanRoomAutoApprovalRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.UpdateCleanRoomAutoApprovalRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCleanRoomAutoApprovalRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateCleanRoomAutoApprovalRuleRequestToPb(st *UpdateCleanRoomAutoApprovalRuleRequest) (*cleanroomspb.UpdateCleanRoomAutoApprovalRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.UpdateCleanRoomAutoApprovalRuleRequestPb{}
	autoApprovalRulePb, err := CleanRoomAutoApprovalRuleToPb(&st.AutoApprovalRule)
	if err != nil {
		return nil, err
	}
	if autoApprovalRulePb != nil {
		pb.AutoApprovalRule = *autoApprovalRulePb
	}
	pb.CleanRoomName = st.CleanRoomName
	pb.RuleId = st.RuleId

	return pb, nil
}

func UpdateCleanRoomAutoApprovalRuleRequestFromPb(pb *cleanroomspb.UpdateCleanRoomAutoApprovalRuleRequestPb) (*UpdateCleanRoomAutoApprovalRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCleanRoomAutoApprovalRuleRequest{}
	autoApprovalRuleField, err := CleanRoomAutoApprovalRuleFromPb(&pb.AutoApprovalRule)
	if err != nil {
		return nil, err
	}
	if autoApprovalRuleField != nil {
		st.AutoApprovalRule = *autoApprovalRuleField
	}
	st.CleanRoomName = pb.CleanRoomName
	st.RuleId = pb.RuleId

	return st, nil
}

type UpdateCleanRoomRequest struct {

	// Wire name: 'clean_room'
	CleanRoom *CleanRoom `json:"clean_room,omitempty"`
	// Name of the clean room.
	Name string `json:"-" tf:"-"`
}

func (st UpdateCleanRoomRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCleanRoomRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCleanRoomRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanroomspb.UpdateCleanRoomRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCleanRoomRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateCleanRoomRequestToPb(st *UpdateCleanRoomRequest) (*cleanroomspb.UpdateCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanroomspb.UpdateCleanRoomRequestPb{}
	cleanRoomPb, err := CleanRoomToPb(st.CleanRoom)
	if err != nil {
		return nil, err
	}
	if cleanRoomPb != nil {
		pb.CleanRoom = cleanRoomPb
	}
	pb.Name = st.Name

	return pb, nil
}

func UpdateCleanRoomRequestFromPb(pb *cleanroomspb.UpdateCleanRoomRequestPb) (*UpdateCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCleanRoomRequest{}
	cleanRoomField, err := CleanRoomFromPb(pb.CleanRoom)
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
