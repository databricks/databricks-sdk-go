// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/google/go-querystring/query"
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
	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoom) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomToPb(&st)
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

// Metadata of the clean room asset
type CleanRoomAsset struct {
	// When the asset is added to the clean room, in epoch milliseconds.
	// Wire name: 'added_at'
	AddedAt int64 `json:"added_at,omitempty"`
	// The type of the asset.
	// Wire name: 'asset_type'
	AssetType CleanRoomAssetAssetType `json:"asset_type,omitempty"`
	// The name of the clean room this asset belongs to. This is an output-only
	// field to ensure proper resource identification.
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
	Name string `json:"name,omitempty"`
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

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomAsset) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetToPb(&st)
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

type CleanRoomAssetForeignTable struct {
	// The metadata information of the columns in the foreign table
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func (st CleanRoomAssetForeignTable) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetForeignTableToPb(&st)
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

type CleanRoomAssetForeignTableLocalDetails struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	// Wire name: 'local_name'
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomAssetForeignTableLocalDetails) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetForeignTableLocalDetailsToPb(&st)
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

type CleanRoomAssetNotebook struct {
	// Server generated etag that represents the notebook version.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	// Wire name: 'notebook_content'
	NotebookContent string `json:"notebook_content,omitempty"`
	// top-level status derived from all reviews
	// Wire name: 'review_state'
	ReviewState CleanRoomNotebookReviewNotebookReviewState `json:"review_state,omitempty"`
	// All existing approvals or rejections
	// Wire name: 'reviews'
	Reviews []CleanRoomNotebookReview `json:"reviews,omitempty"`
	// collaborators that can run the notebook
	// Wire name: 'runner_collaborator_aliases'
	RunnerCollaboratorAliases []string `json:"runner_collaborator_aliases,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomAssetNotebook) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetNotebookToPb(&st)
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

type CleanRoomAssetTable struct {
	// The metadata information of the columns in the table
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func (st CleanRoomAssetTable) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetTableToPb(&st)
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

type CleanRoomAssetTableLocalDetails struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	// Wire name: 'local_name'
	LocalName string `json:"local_name,omitempty"`
	// Partition filtering specification for a shared table.
	// Wire name: 'partitions'
	Partitions []sharing.Partition `json:"partitions,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomAssetTableLocalDetails) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetTableLocalDetailsToPb(&st)
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

type CleanRoomAssetView struct {
	// The metadata information of the columns in the view
	// Wire name: 'columns'
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func (st CleanRoomAssetView) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetViewToPb(&st)
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

type CleanRoomAssetViewLocalDetails struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	// Wire name: 'local_name'
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomAssetViewLocalDetails) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetViewLocalDetailsToPb(&st)
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

type CleanRoomAssetVolumeLocalDetails struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	// Wire name: 'local_name'
	LocalName string `json:"local_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomAssetVolumeLocalDetails) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomAssetVolumeLocalDetailsToPb(&st)
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
	OrganizationName string `json:"organization_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomCollaborator) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomCollaboratorToPb(&st)
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
	ReviewerCollaboratorAlias string `json:"reviewer_collaborator_alias,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomNotebookReview) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomNotebookReviewToPb(&st)
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

func (st *CleanRoomNotebookReview) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomNotebookReviewPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomNotebookReviewFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomNotebookReview) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomNotebookReviewToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	StartTime int64 `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomNotebookTaskRun) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomNotebookTaskRunToPb(&st)
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

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`

	// Wire name: 'status'
	Status CleanRoomOutputCatalogOutputCatalogStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomOutputCatalog) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomOutputCatalogToPb(&st)
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
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CleanRoomRemoteDetail) EncodeValues(key string, v *url.Values) error {
	pb, err := cleanRoomRemoteDetailToPb(&st)
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
	CollaboratorWorkspaceId int64 `json:"collaborator_workspace_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CollaboratorJobRunInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := collaboratorJobRunInfoToPb(&st)
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

// The compliance security profile used to process regulated data following
// compliance standards.
type ComplianceSecurityProfile struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	// Wire name: 'compliance_standards'
	ComplianceStandards []settings.ComplianceStandard `json:"compliance_standards,omitempty"`
	// Whether the compliance security profile is enabled.
	// Wire name: 'is_enabled'
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ComplianceSecurityProfile) EncodeValues(key string, v *url.Values) error {
	pb, err := complianceSecurityProfileToPb(&st)
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

type CreateCleanRoomAssetRequest struct {

	// Wire name: 'asset'
	Asset CleanRoomAsset `json:"asset"`
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
}

func (st CreateCleanRoomAssetRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createCleanRoomAssetRequestToPb(&st)
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

type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`

	// Wire name: 'output_catalog'
	OutputCatalog CleanRoomOutputCatalog `json:"output_catalog"`
}

func (st CreateCleanRoomOutputCatalogRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createCleanRoomOutputCatalogRequestToPb(&st)
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

type CreateCleanRoomOutputCatalogResponse struct {

	// Wire name: 'output_catalog'
	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
}

func (st CreateCleanRoomOutputCatalogResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createCleanRoomOutputCatalogResponseToPb(&st)
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

type CreateCleanRoomRequest struct {

	// Wire name: 'clean_room'
	CleanRoom CleanRoom `json:"clean_room"`
}

func (st CreateCleanRoomRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createCleanRoomRequestToPb(&st)
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

type DeleteCleanRoomAssetRequest struct {
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" tf:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name string `json:"-" tf:"-"`
}

func (st DeleteCleanRoomAssetRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteCleanRoomAssetRequestToPb(&st)
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

type DeleteCleanRoomRequest struct {
	// Name of the clean room.
	Name string `json:"-" tf:"-"`
}

func (st DeleteCleanRoomRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteCleanRoomRequestToPb(&st)
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

type DeleteResponse struct {
}

func (st DeleteResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteResponseToPb(&st)
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

type GetCleanRoomAssetRequest struct {
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" tf:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name string `json:"-" tf:"-"`
}

func (st GetCleanRoomAssetRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getCleanRoomAssetRequestToPb(&st)
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

type GetCleanRoomRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st GetCleanRoomRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getCleanRoomRequestToPb(&st)
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

type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomAssetsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listCleanRoomAssetsRequestToPb(&st)
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

type ListCleanRoomAssetsResponse struct {
	// Assets in the clean room.
	// Wire name: 'assets'
	Assets []CleanRoomAsset `json:"assets,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomAssetsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listCleanRoomAssetsResponseToPb(&st)
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

type ListCleanRoomNotebookTaskRunsRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" tf:"-"`
	// Notebook name
	NotebookName string `json:"-" tf:"-"`
	// The maximum number of task runs to return. Currently ignored - all runs
	// will be returned.
	PageSize int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomNotebookTaskRunsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listCleanRoomNotebookTaskRunsRequestToPb(&st)
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

type ListCleanRoomNotebookTaskRunsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// Name of the clean room.
	// Wire name: 'runs'
	Runs []CleanRoomNotebookTaskRun `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomNotebookTaskRunsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listCleanRoomNotebookTaskRunsResponseToPb(&st)
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

type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listCleanRoomsRequestToPb(&st)
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

type ListCleanRoomsResponse struct {

	// Wire name: 'clean_rooms'
	CleanRooms []CleanRoom `json:"clean_rooms,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCleanRoomsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listCleanRoomsResponseToPb(&st)
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

func (st UpdateCleanRoomAssetRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCleanRoomAssetRequestToPb(&st)
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

type UpdateCleanRoomRequest struct {

	// Wire name: 'clean_room'
	CleanRoom *CleanRoom `json:"clean_room,omitempty"`
	// Name of the clean room.
	Name string `json:"-" tf:"-"`
}

func (st UpdateCleanRoomRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCleanRoomRequestToPb(&st)
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
