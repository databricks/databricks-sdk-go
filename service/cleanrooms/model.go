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

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type CleanRoom struct {
	// Whether clean room access is restricted due to [CSP]
	//
	// [CSP]: https://docs.databricks.com/en/security/privacy/security-profile.html
	AccessRestricted CleanRoomAccessRestricted

	Comment string
	// When the clean room was created, in epoch milliseconds.
	CreatedAt int64
	// The alias of the collaborator tied to the local clean room.
	LocalCollaboratorAlias string
	// The name of the clean room. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	Name string
	// Output catalog of the clean room. It is an output only field. Output
	// catalog is manipulated using the separate CreateCleanRoomOutputCatalog
	// API.
	OutputCatalog *CleanRoomOutputCatalog
	// This is Databricks username of the owner of the local clean room
	// securable for permission management.
	Owner string
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	RemoteDetailedInfo *CleanRoomRemoteDetail
	// Clean room status.
	Status CleanRoomStatusEnum
	// When the clean room was last updated, in epoch milliseconds.
	UpdatedAt int64

	ForceSendFields []string
}

func cleanRoomToPb(st *CleanRoom) (*cleanRoomPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomPb{}
	accessRestrictedPb, err := identity(&st.AccessRestricted)
	if err != nil {
		return nil, err
	}
	if accessRestrictedPb != nil {
		pb.AccessRestricted = *accessRestrictedPb
	}

	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	localCollaboratorAliasPb, err := identity(&st.LocalCollaboratorAlias)
	if err != nil {
		return nil, err
	}
	if localCollaboratorAliasPb != nil {
		pb.LocalCollaboratorAlias = *localCollaboratorAliasPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	outputCatalogPb, err := cleanRoomOutputCatalogToPb(st.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogPb != nil {
		pb.OutputCatalog = outputCatalogPb
	}

	ownerPb, err := identity(&st.Owner)
	if err != nil {
		return nil, err
	}
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	remoteDetailedInfoPb, err := cleanRoomRemoteDetailToPb(st.RemoteDetailedInfo)
	if err != nil {
		return nil, err
	}
	if remoteDetailedInfoPb != nil {
		pb.RemoteDetailedInfo = remoteDetailedInfoPb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

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
	accessRestrictedField, err := identity(&pb.AccessRestricted)
	if err != nil {
		return nil, err
	}
	if accessRestrictedField != nil {
		st.AccessRestricted = *accessRestrictedField
	}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	localCollaboratorAliasField, err := identity(&pb.LocalCollaboratorAlias)
	if err != nil {
		return nil, err
	}
	if localCollaboratorAliasField != nil {
		st.LocalCollaboratorAlias = *localCollaboratorAliasField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	outputCatalogField, err := cleanRoomOutputCatalogFromPb(pb.OutputCatalog)
	if err != nil {
		return nil, err
	}
	if outputCatalogField != nil {
		st.OutputCatalog = outputCatalogField
	}
	ownerField, err := identity(&pb.Owner)
	if err != nil {
		return nil, err
	}
	if ownerField != nil {
		st.Owner = *ownerField
	}
	remoteDetailedInfoField, err := cleanRoomRemoteDetailFromPb(pb.RemoteDetailedInfo)
	if err != nil {
		return nil, err
	}
	if remoteDetailedInfoField != nil {
		st.RemoteDetailedInfo = remoteDetailedInfoField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}

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
	AddedAt int64
	// The type of the asset.
	AssetType CleanRoomAssetAssetType
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTable *CleanRoomAssetForeignTable
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTableLocalDetails *CleanRoomAssetForeignTableLocalDetails
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name string
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	Notebook *CleanRoomAssetNotebook
	// The alias of the collaborator who owns this asset
	OwnerCollaboratorAlias string
	// Status of the asset
	Status CleanRoomAssetStatusEnum
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	Table *CleanRoomAssetTable
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	TableLocalDetails *CleanRoomAssetTableLocalDetails
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	View *CleanRoomAssetView
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	ViewLocalDetails *CleanRoomAssetViewLocalDetails
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	VolumeLocalDetails *CleanRoomAssetVolumeLocalDetails

	ForceSendFields []string
}

func cleanRoomAssetToPb(st *CleanRoomAsset) (*cleanRoomAssetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetPb{}
	addedAtPb, err := identity(&st.AddedAt)
	if err != nil {
		return nil, err
	}
	if addedAtPb != nil {
		pb.AddedAt = *addedAtPb
	}

	assetTypePb, err := identity(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}

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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	notebookPb, err := cleanRoomAssetNotebookToPb(st.Notebook)
	if err != nil {
		return nil, err
	}
	if notebookPb != nil {
		pb.Notebook = notebookPb
	}

	ownerCollaboratorAliasPb, err := identity(&st.OwnerCollaboratorAlias)
	if err != nil {
		return nil, err
	}
	if ownerCollaboratorAliasPb != nil {
		pb.OwnerCollaboratorAlias = *ownerCollaboratorAliasPb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

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
	addedAtField, err := identity(&pb.AddedAt)
	if err != nil {
		return nil, err
	}
	if addedAtField != nil {
		st.AddedAt = *addedAtField
	}
	assetTypeField, err := identity(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	notebookField, err := cleanRoomAssetNotebookFromPb(pb.Notebook)
	if err != nil {
		return nil, err
	}
	if notebookField != nil {
		st.Notebook = notebookField
	}
	ownerCollaboratorAliasField, err := identity(&pb.OwnerCollaboratorAlias)
	if err != nil {
		return nil, err
	}
	if ownerCollaboratorAliasField != nil {
		st.OwnerCollaboratorAlias = *ownerCollaboratorAliasField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
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
	LocalName string

	ForceSendFields []string
}

func cleanRoomAssetForeignTableLocalDetailsToPb(st *CleanRoomAssetForeignTableLocalDetails) (*cleanRoomAssetForeignTableLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetForeignTableLocalDetailsPb{}
	localNamePb, err := identity(&st.LocalName)
	if err != nil {
		return nil, err
	}
	if localNamePb != nil {
		pb.LocalName = *localNamePb
	}

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
	localNameField, err := identity(&pb.LocalName)
	if err != nil {
		return nil, err
	}
	if localNameField != nil {
		st.LocalName = *localNameField
	}

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
	Etag string
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	NotebookContent string

	ForceSendFields []string
}

func cleanRoomAssetNotebookToPb(st *CleanRoomAssetNotebook) (*cleanRoomAssetNotebookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetNotebookPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	notebookContentPb, err := identity(&st.NotebookContent)
	if err != nil {
		return nil, err
	}
	if notebookContentPb != nil {
		pb.NotebookContent = *notebookContentPb
	}

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
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	notebookContentField, err := identity(&pb.NotebookContent)
	if err != nil {
		return nil, err
	}
	if notebookContentField != nil {
		st.NotebookContent = *notebookContentField
	}

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
	LocalName string
	// Partition filtering specification for a shared table.
	Partitions []sharing.Partition

	ForceSendFields []string
}

func cleanRoomAssetTableLocalDetailsToPb(st *CleanRoomAssetTableLocalDetails) (*cleanRoomAssetTableLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetTableLocalDetailsPb{}
	localNamePb, err := identity(&st.LocalName)
	if err != nil {
		return nil, err
	}
	if localNamePb != nil {
		pb.LocalName = *localNamePb
	}

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
	localNameField, err := identity(&pb.LocalName)
	if err != nil {
		return nil, err
	}
	if localNameField != nil {
		st.LocalName = *localNameField
	}

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
	LocalName string

	ForceSendFields []string
}

func cleanRoomAssetViewLocalDetailsToPb(st *CleanRoomAssetViewLocalDetails) (*cleanRoomAssetViewLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetViewLocalDetailsPb{}
	localNamePb, err := identity(&st.LocalName)
	if err != nil {
		return nil, err
	}
	if localNamePb != nil {
		pb.LocalName = *localNamePb
	}

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
	localNameField, err := identity(&pb.LocalName)
	if err != nil {
		return nil, err
	}
	if localNameField != nil {
		st.LocalName = *localNameField
	}

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
	LocalName string

	ForceSendFields []string
}

func cleanRoomAssetVolumeLocalDetailsToPb(st *CleanRoomAssetVolumeLocalDetails) (*cleanRoomAssetVolumeLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetVolumeLocalDetailsPb{}
	localNamePb, err := identity(&st.LocalName)
	if err != nil {
		return nil, err
	}
	if localNamePb != nil {
		pb.LocalName = *localNamePb
	}

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
	localNameField, err := identity(&pb.LocalName)
	if err != nil {
		return nil, err
	}
	if localNameField != nil {
		st.LocalName = *localNameField
	}

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
	CollaboratorAlias string
	// Generated display name for the collaborator. In the case of a single
	// metastore clean room, it is the clean room name. For x-metastore clean
	// rooms, it is the organization name of the metastore. It is not restricted
	// to these values and could change in the future
	DisplayName string
	// The global Unity Catalog metastore id of the collaborator. The identifier
	// is of format cloud:region:metastore-uuid.
	GlobalMetastoreId string
	// Email of the user who is receiving the clean room "invitation". It should
	// be empty for the creator of the clean room, and non-empty for the
	// invitees of the clean room. It is only returned in the output when clean
	// room creator calls GET
	InviteRecipientEmail string
	// Workspace ID of the user who is receiving the clean room "invitation".
	// Must be specified if invite_recipient_email is specified. It should be
	// empty when the collaborator is the creator of the clean room.
	InviteRecipientWorkspaceId int64
	// [Organization
	// name](:method:metastores/list#metastores-delta_sharing_organization_name)
	// configured in the metastore
	OrganizationName string

	ForceSendFields []string
}

func cleanRoomCollaboratorToPb(st *CleanRoomCollaborator) (*cleanRoomCollaboratorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomCollaboratorPb{}
	collaboratorAliasPb, err := identity(&st.CollaboratorAlias)
	if err != nil {
		return nil, err
	}
	if collaboratorAliasPb != nil {
		pb.CollaboratorAlias = *collaboratorAliasPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	globalMetastoreIdPb, err := identity(&st.GlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if globalMetastoreIdPb != nil {
		pb.GlobalMetastoreId = *globalMetastoreIdPb
	}

	inviteRecipientEmailPb, err := identity(&st.InviteRecipientEmail)
	if err != nil {
		return nil, err
	}
	if inviteRecipientEmailPb != nil {
		pb.InviteRecipientEmail = *inviteRecipientEmailPb
	}

	inviteRecipientWorkspaceIdPb, err := identity(&st.InviteRecipientWorkspaceId)
	if err != nil {
		return nil, err
	}
	if inviteRecipientWorkspaceIdPb != nil {
		pb.InviteRecipientWorkspaceId = *inviteRecipientWorkspaceIdPb
	}

	organizationNamePb, err := identity(&st.OrganizationName)
	if err != nil {
		return nil, err
	}
	if organizationNamePb != nil {
		pb.OrganizationName = *organizationNamePb
	}

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
	collaboratorAliasField, err := identity(&pb.CollaboratorAlias)
	if err != nil {
		return nil, err
	}
	if collaboratorAliasField != nil {
		st.CollaboratorAlias = *collaboratorAliasField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	globalMetastoreIdField, err := identity(&pb.GlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if globalMetastoreIdField != nil {
		st.GlobalMetastoreId = *globalMetastoreIdField
	}
	inviteRecipientEmailField, err := identity(&pb.InviteRecipientEmail)
	if err != nil {
		return nil, err
	}
	if inviteRecipientEmailField != nil {
		st.InviteRecipientEmail = *inviteRecipientEmailField
	}
	inviteRecipientWorkspaceIdField, err := identity(&pb.InviteRecipientWorkspaceId)
	if err != nil {
		return nil, err
	}
	if inviteRecipientWorkspaceIdField != nil {
		st.InviteRecipientWorkspaceId = *inviteRecipientWorkspaceIdField
	}
	organizationNameField, err := identity(&pb.OrganizationName)
	if err != nil {
		return nil, err
	}
	if organizationNameField != nil {
		st.OrganizationName = *organizationNameField
	}

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
	CollaboratorJobRunInfo *CollaboratorJobRunInfo
	// State of the task run.
	NotebookJobRunState *jobs.CleanRoomTaskRunState
	// Asset name of the notebook executed in this task run.
	NotebookName string
	// Expiration time of the output schema of the task run (if any), in epoch
	// milliseconds.
	OutputSchemaExpirationTime int64
	// Name of the output schema associated with the clean rooms notebook task
	// run.
	OutputSchemaName string
	// Duration of the task run, in milliseconds.
	RunDuration int64
	// When the task run started, in epoch milliseconds.
	StartTime int64

	ForceSendFields []string
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

	notebookNamePb, err := identity(&st.NotebookName)
	if err != nil {
		return nil, err
	}
	if notebookNamePb != nil {
		pb.NotebookName = *notebookNamePb
	}

	outputSchemaExpirationTimePb, err := identity(&st.OutputSchemaExpirationTime)
	if err != nil {
		return nil, err
	}
	if outputSchemaExpirationTimePb != nil {
		pb.OutputSchemaExpirationTime = *outputSchemaExpirationTimePb
	}

	outputSchemaNamePb, err := identity(&st.OutputSchemaName)
	if err != nil {
		return nil, err
	}
	if outputSchemaNamePb != nil {
		pb.OutputSchemaName = *outputSchemaNamePb
	}

	runDurationPb, err := identity(&st.RunDuration)
	if err != nil {
		return nil, err
	}
	if runDurationPb != nil {
		pb.RunDuration = *runDurationPb
	}

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

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
	notebookNameField, err := identity(&pb.NotebookName)
	if err != nil {
		return nil, err
	}
	if notebookNameField != nil {
		st.NotebookName = *notebookNameField
	}
	outputSchemaExpirationTimeField, err := identity(&pb.OutputSchemaExpirationTime)
	if err != nil {
		return nil, err
	}
	if outputSchemaExpirationTimeField != nil {
		st.OutputSchemaExpirationTime = *outputSchemaExpirationTimeField
	}
	outputSchemaNameField, err := identity(&pb.OutputSchemaName)
	if err != nil {
		return nil, err
	}
	if outputSchemaNameField != nil {
		st.OutputSchemaName = *outputSchemaNameField
	}
	runDurationField, err := identity(&pb.RunDuration)
	if err != nil {
		return nil, err
	}
	if runDurationField != nil {
		st.RunDuration = *runDurationField
	}
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}

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
	CatalogName string

	Status CleanRoomOutputCatalogOutputCatalogStatus

	ForceSendFields []string
}

func cleanRoomOutputCatalogToPb(st *CleanRoomOutputCatalog) (*cleanRoomOutputCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomOutputCatalogPb{}
	catalogNamePb, err := identity(&st.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

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
	catalogNameField, err := identity(&pb.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

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
	CentralCleanRoomId string
	// Cloud vendor (aws,azure,gcp) of the central clean room.
	CloudVendor string
	// Collaborators in the central clean room. There should one and only one
	// collaborator in the list that satisfies the owner condition:
	//
	// 1. It has the creator's global_metastore_id (determined by caller of
	// CreateCleanRoom).
	//
	// 2. Its invite_recipient_email is empty.
	Collaborators []CleanRoomCollaborator
	// The compliance security profile used to process regulated data following
	// compliance standards.
	ComplianceSecurityProfile *ComplianceSecurityProfile
	// Collaborator who creates the clean room.
	Creator *CleanRoomCollaborator
	// Egress network policy to apply to the central clean room workspace.
	EgressNetworkPolicy *settings.EgressNetworkPolicy
	// Region of the central clean room.
	Region string

	ForceSendFields []string
}

func cleanRoomRemoteDetailToPb(st *CleanRoomRemoteDetail) (*cleanRoomRemoteDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomRemoteDetailPb{}
	centralCleanRoomIdPb, err := identity(&st.CentralCleanRoomId)
	if err != nil {
		return nil, err
	}
	if centralCleanRoomIdPb != nil {
		pb.CentralCleanRoomId = *centralCleanRoomIdPb
	}

	cloudVendorPb, err := identity(&st.CloudVendor)
	if err != nil {
		return nil, err
	}
	if cloudVendorPb != nil {
		pb.CloudVendor = *cloudVendorPb
	}

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

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

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
	centralCleanRoomIdField, err := identity(&pb.CentralCleanRoomId)
	if err != nil {
		return nil, err
	}
	if centralCleanRoomIdField != nil {
		st.CentralCleanRoomId = *centralCleanRoomIdField
	}
	cloudVendorField, err := identity(&pb.CloudVendor)
	if err != nil {
		return nil, err
	}
	if cloudVendorField != nil {
		st.CloudVendor = *cloudVendorField
	}

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
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}

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
	CollaboratorAlias string
	// Job ID of the task run in the collaborator's workspace.
	CollaboratorJobId int64
	// Job run ID of the task run in the collaborator's workspace.
	CollaboratorJobRunId int64
	// Task run ID of the task run in the collaborator's workspace.
	CollaboratorTaskRunId int64
	// ID of the collaborator's workspace that triggered the task run.
	CollaboratorWorkspaceId int64

	ForceSendFields []string
}

func collaboratorJobRunInfoToPb(st *CollaboratorJobRunInfo) (*collaboratorJobRunInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &collaboratorJobRunInfoPb{}
	collaboratorAliasPb, err := identity(&st.CollaboratorAlias)
	if err != nil {
		return nil, err
	}
	if collaboratorAliasPb != nil {
		pb.CollaboratorAlias = *collaboratorAliasPb
	}

	collaboratorJobIdPb, err := identity(&st.CollaboratorJobId)
	if err != nil {
		return nil, err
	}
	if collaboratorJobIdPb != nil {
		pb.CollaboratorJobId = *collaboratorJobIdPb
	}

	collaboratorJobRunIdPb, err := identity(&st.CollaboratorJobRunId)
	if err != nil {
		return nil, err
	}
	if collaboratorJobRunIdPb != nil {
		pb.CollaboratorJobRunId = *collaboratorJobRunIdPb
	}

	collaboratorTaskRunIdPb, err := identity(&st.CollaboratorTaskRunId)
	if err != nil {
		return nil, err
	}
	if collaboratorTaskRunIdPb != nil {
		pb.CollaboratorTaskRunId = *collaboratorTaskRunIdPb
	}

	collaboratorWorkspaceIdPb, err := identity(&st.CollaboratorWorkspaceId)
	if err != nil {
		return nil, err
	}
	if collaboratorWorkspaceIdPb != nil {
		pb.CollaboratorWorkspaceId = *collaboratorWorkspaceIdPb
	}

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
	collaboratorAliasField, err := identity(&pb.CollaboratorAlias)
	if err != nil {
		return nil, err
	}
	if collaboratorAliasField != nil {
		st.CollaboratorAlias = *collaboratorAliasField
	}
	collaboratorJobIdField, err := identity(&pb.CollaboratorJobId)
	if err != nil {
		return nil, err
	}
	if collaboratorJobIdField != nil {
		st.CollaboratorJobId = *collaboratorJobIdField
	}
	collaboratorJobRunIdField, err := identity(&pb.CollaboratorJobRunId)
	if err != nil {
		return nil, err
	}
	if collaboratorJobRunIdField != nil {
		st.CollaboratorJobRunId = *collaboratorJobRunIdField
	}
	collaboratorTaskRunIdField, err := identity(&pb.CollaboratorTaskRunId)
	if err != nil {
		return nil, err
	}
	if collaboratorTaskRunIdField != nil {
		st.CollaboratorTaskRunId = *collaboratorTaskRunIdField
	}
	collaboratorWorkspaceIdField, err := identity(&pb.CollaboratorWorkspaceId)
	if err != nil {
		return nil, err
	}
	if collaboratorWorkspaceIdField != nil {
		st.CollaboratorWorkspaceId = *collaboratorWorkspaceIdField
	}

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
	ComplianceStandards []settings.ComplianceStandard
	// Whether the compliance security profile is enabled.
	IsEnabled bool

	ForceSendFields []string
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

	isEnabledPb, err := identity(&st.IsEnabled)
	if err != nil {
		return nil, err
	}
	if isEnabledPb != nil {
		pb.IsEnabled = *isEnabledPb
	}

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
	isEnabledField, err := identity(&pb.IsEnabled)
	if err != nil {
		return nil, err
	}
	if isEnabledField != nil {
		st.IsEnabled = *isEnabledField
	}

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
	Asset CleanRoomAsset
	// Name of the clean room.
	CleanRoomName string
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

	cleanRoomNamePb, err := identity(&st.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNamePb != nil {
		pb.CleanRoomName = *cleanRoomNamePb
	}

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
	cleanRoomNameField, err := identity(&pb.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNameField != nil {
		st.CleanRoomName = *cleanRoomNameField
	}

	return st, nil
}

// Create an output catalog
type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName string

	OutputCatalog CleanRoomOutputCatalog
}

func createCleanRoomOutputCatalogRequestToPb(st *CreateCleanRoomOutputCatalogRequest) (*createCleanRoomOutputCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomOutputCatalogRequestPb{}
	cleanRoomNamePb, err := identity(&st.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNamePb != nil {
		pb.CleanRoomName = *cleanRoomNamePb
	}

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
	cleanRoomNameField, err := identity(&pb.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNameField != nil {
		st.CleanRoomName = *cleanRoomNameField
	}
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
	AssetFullName string
	// The type of the asset.
	AssetType CleanRoomAssetAssetType
	// Name of the clean room.
	CleanRoomName string
}

func deleteCleanRoomAssetRequestToPb(st *DeleteCleanRoomAssetRequest) (*deleteCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCleanRoomAssetRequestPb{}
	assetFullNamePb, err := identity(&st.AssetFullName)
	if err != nil {
		return nil, err
	}
	if assetFullNamePb != nil {
		pb.AssetFullName = *assetFullNamePb
	}

	assetTypePb, err := identity(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}

	cleanRoomNamePb, err := identity(&st.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNamePb != nil {
		pb.CleanRoomName = *cleanRoomNamePb
	}

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
	assetFullNameField, err := identity(&pb.AssetFullName)
	if err != nil {
		return nil, err
	}
	if assetFullNameField != nil {
		st.AssetFullName = *assetFullNameField
	}
	assetTypeField, err := identity(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	cleanRoomNameField, err := identity(&pb.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNameField != nil {
		st.CleanRoomName = *cleanRoomNameField
	}

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
	Name string
}

func deleteCleanRoomRequestToPb(st *DeleteCleanRoomRequest) (*deleteCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCleanRoomRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	AssetFullName string
	// The type of the asset.
	AssetType CleanRoomAssetAssetType
	// Name of the clean room.
	CleanRoomName string
}

func getCleanRoomAssetRequestToPb(st *GetCleanRoomAssetRequest) (*getCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCleanRoomAssetRequestPb{}
	assetFullNamePb, err := identity(&st.AssetFullName)
	if err != nil {
		return nil, err
	}
	if assetFullNamePb != nil {
		pb.AssetFullName = *assetFullNamePb
	}

	assetTypePb, err := identity(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}

	cleanRoomNamePb, err := identity(&st.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNamePb != nil {
		pb.CleanRoomName = *cleanRoomNamePb
	}

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
	assetFullNameField, err := identity(&pb.AssetFullName)
	if err != nil {
		return nil, err
	}
	if assetFullNameField != nil {
		st.AssetFullName = *assetFullNameField
	}
	assetTypeField, err := identity(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	cleanRoomNameField, err := identity(&pb.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNameField != nil {
		st.CleanRoomName = *cleanRoomNameField
	}

	return st, nil
}

// Get a clean room
type GetCleanRoomRequest struct {
	Name string
}

func getCleanRoomRequestToPb(st *GetCleanRoomRequest) (*getCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCleanRoomRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// List assets
type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName string
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listCleanRoomAssetsRequestToPb(st *ListCleanRoomAssetsRequest) (*listCleanRoomAssetsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomAssetsRequestPb{}
	cleanRoomNamePb, err := identity(&st.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNamePb != nil {
		pb.CleanRoomName = *cleanRoomNamePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	cleanRoomNameField, err := identity(&pb.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNameField != nil {
		st.CleanRoomName = *cleanRoomNameField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	Assets []CleanRoomAsset
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	CleanRoomName string
	// Notebook name
	NotebookName string
	// The maximum number of task runs to return. Currently ignored - all runs
	// will be returned.
	PageSize int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listCleanRoomNotebookTaskRunsRequestToPb(st *ListCleanRoomNotebookTaskRunsRequest) (*listCleanRoomNotebookTaskRunsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomNotebookTaskRunsRequestPb{}
	cleanRoomNamePb, err := identity(&st.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNamePb != nil {
		pb.CleanRoomName = *cleanRoomNamePb
	}

	notebookNamePb, err := identity(&st.NotebookName)
	if err != nil {
		return nil, err
	}
	if notebookNamePb != nil {
		pb.NotebookName = *notebookNamePb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	cleanRoomNameField, err := identity(&pb.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNameField != nil {
		st.CleanRoomName = *cleanRoomNameField
	}
	notebookNameField, err := identity(&pb.NotebookName)
	if err != nil {
		return nil, err
	}
	if notebookNameField != nil {
		st.NotebookName = *notebookNameField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string
	// Name of the clean room.
	Runs []CleanRoomNotebookTaskRun

	ForceSendFields []string
}

func listCleanRoomNotebookTaskRunsResponseToPb(st *ListCleanRoomNotebookTaskRunsResponse) (*listCleanRoomNotebookTaskRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomNotebookTaskRunsResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	PageSize int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listCleanRoomsRequestToPb(st *ListCleanRoomsRequest) (*listCleanRoomsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomsRequestPb{}
	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	CleanRooms []CleanRoom
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	Asset CleanRoomAsset
	// The type of the asset.
	AssetType CleanRoomAssetAssetType
	// Name of the clean room.
	CleanRoomName string
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name string
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

	assetTypePb, err := identity(&st.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypePb != nil {
		pb.AssetType = *assetTypePb
	}

	cleanRoomNamePb, err := identity(&st.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNamePb != nil {
		pb.CleanRoomName = *cleanRoomNamePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	assetTypeField, err := identity(&pb.AssetType)
	if err != nil {
		return nil, err
	}
	if assetTypeField != nil {
		st.AssetType = *assetTypeField
	}
	cleanRoomNameField, err := identity(&pb.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNameField != nil {
		st.CleanRoomName = *cleanRoomNameField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type UpdateCleanRoomRequest struct {
	CleanRoom *CleanRoom
	// Name of the clean room.
	Name string
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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}
