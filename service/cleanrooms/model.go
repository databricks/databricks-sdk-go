// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"fmt"

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
	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
	// This is the Databricks username of the owner of the local clean room
	// securable for permission management.
	Owner string `json:"owner,omitempty"`
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	RemoteDetailedInfo *CleanRoomRemoteDetail `json:"remote_detailed_info,omitempty"`
	// Clean room status.
	Status CleanRoomStatusEnum `json:"status,omitempty"`
	// When the clean room was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CleanRoom) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoom) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	AddedAt int64 `json:"added_at,omitempty"`
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"asset_type"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName string `json:"clean_room_name,omitempty"`
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTable *CleanRoomAssetForeignTable `json:"foreign_table,omitempty"`
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTableLocalDetails *CleanRoomAssetForeignTableLocalDetails `json:"foreign_table_local_details,omitempty"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name. For jar analyses, the
	// name is the jar analysis name.
	Name string `json:"name"`
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	Notebook *CleanRoomAssetNotebook `json:"notebook,omitempty"`
	// The alias of the collaborator who owns this asset
	OwnerCollaboratorAlias string `json:"owner_collaborator_alias,omitempty"`
	// Status of the asset
	Status CleanRoomAssetStatusEnum `json:"status,omitempty"`
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	Table *CleanRoomAssetTable `json:"table,omitempty"`
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	TableLocalDetails *CleanRoomAssetTableLocalDetails `json:"table_local_details,omitempty"`
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	View *CleanRoomAssetView `json:"view,omitempty"`
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	ViewLocalDetails *CleanRoomAssetViewLocalDetails `json:"view_local_details,omitempty"`
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	VolumeLocalDetails *CleanRoomAssetVolumeLocalDetails `json:"volume_local_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CleanRoomAsset) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAsset) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

type CleanRoomAssetForeignTableLocalDetails struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	LocalName string `json:"local_name"`
}

type CleanRoomAssetNotebook struct {
	// Server generated etag that represents the notebook version.
	Etag string `json:"etag,omitempty"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	NotebookContent string `json:"notebook_content"`
	// Top-level status derived from all reviews
	ReviewState CleanRoomNotebookReviewNotebookReviewState `json:"review_state,omitempty"`
	// All existing approvals or rejections
	Reviews []CleanRoomNotebookReview `json:"reviews,omitempty"`
	// Aliases of collaborators that can run the notebook.
	RunnerCollaboratorAliases []string `json:"runner_collaborator_aliases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CleanRoomAssetNotebook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAssetNotebook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

type CleanRoomAssetTableLocalDetails struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	LocalName string `json:"local_name"`
	// Partition filtering specification for a shared table.
	Partitions []sharing.Partition `json:"partitions,omitempty"`
}

type CleanRoomAssetView struct {
	// The metadata information of the columns in the view
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

type CleanRoomAssetViewLocalDetails struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	LocalName string `json:"local_name"`
}

type CleanRoomAssetVolumeLocalDetails struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	LocalName string `json:"local_name"`
}

type CleanRoomAutoApprovalRule struct {
	// Collaborator alias of the author covered by the rule. Only one of
	// `author_collaborator_alias` and `author_scope` can be set.
	AuthorCollaboratorAlias string `json:"author_collaborator_alias,omitempty"`
	// Scope of authors covered by the rule. Only one of
	// `author_collaborator_alias` and `author_scope` can be set.
	AuthorScope CleanRoomAutoApprovalRuleAuthorScope `json:"author_scope,omitempty"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName string `json:"clean_room_name,omitempty"`
	// Timestamp of when the rule was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// A generated UUID identifying the rule.
	RuleId string `json:"rule_id,omitempty"`
	// The owner of the rule to whom the rule applies.
	RuleOwnerCollaboratorAlias string `json:"rule_owner_collaborator_alias,omitempty"`
	// Collaborator alias of the runner covered by the rule.
	RunnerCollaboratorAlias string `json:"runner_collaborator_alias,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CleanRoomAutoApprovalRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAutoApprovalRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Publicly visible clean room collaborator.
type CleanRoomCollaborator struct {
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
	// The global Unity Catalog metastore ID of the collaborator. The identifier
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

func (s *CleanRoomCollaborator) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomCollaborator) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomNotebookReview struct {
	// Review comment
	Comment string `json:"comment,omitempty"`
	// When the review was submitted, in epoch milliseconds
	CreatedAtMillis int64 `json:"created_at_millis,omitempty"`
	// Review outcome
	ReviewState CleanRoomNotebookReviewNotebookReviewState `json:"review_state,omitempty"`
	// Specified when the review was not explicitly made by a user
	ReviewSubReason CleanRoomNotebookReviewNotebookReviewSubReason `json:"review_sub_reason,omitempty"`
	// Collaborator alias of the reviewer
	ReviewerCollaboratorAlias string `json:"reviewer_collaborator_alias,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CleanRoomNotebookReview) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomNotebookReview) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	CollaboratorJobRunInfo *CollaboratorJobRunInfo `json:"collaborator_job_run_info,omitempty"`
	// Etag of the notebook executed in this task run, used to identify the
	// notebook version.
	NotebookEtag string `json:"notebook_etag,omitempty"`
	// State of the task run.
	NotebookJobRunState *jobs.CleanRoomTaskRunState `json:"notebook_job_run_state,omitempty"`
	// Asset name of the notebook executed in this task run.
	NotebookName string `json:"notebook_name,omitempty"`
	// The timestamp of when the notebook was last updated.
	NotebookUpdatedAt int64 `json:"notebook_updated_at,omitempty"`
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

func (s *CleanRoomNotebookTaskRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomNotebookTaskRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName string `json:"catalog_name,omitempty"`

	Status CleanRoomOutputCatalogOutputCatalogStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CleanRoomOutputCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomOutputCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Collaborators []CleanRoomCollaborator `json:"collaborators,omitempty"`

	ComplianceSecurityProfile *ComplianceSecurityProfile `json:"compliance_security_profile,omitempty"`
	// Collaborator who creates the clean room.
	Creator *CleanRoomCollaborator `json:"creator,omitempty"`
	// Egress network policy to apply to the central clean room workspace.
	EgressNetworkPolicy *settings.EgressNetworkPolicy `json:"egress_network_policy,omitempty"`
	// Region of the central clean room.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CleanRoomRemoteDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomRemoteDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

func (s *CollaboratorJobRunInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CollaboratorJobRunInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The compliance security profile used to process regulated data following
// compliance standards.
type ComplianceSecurityProfile struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	ComplianceStandards []settings.ComplianceStandard `json:"compliance_standards,omitempty"`
	// Whether the compliance security profile is enabled.
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCleanRoomAssetRequest struct {
	Asset CleanRoomAsset `json:"asset"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName string `json:"-" url:"-"`
}

type CreateCleanRoomAssetReviewRequest struct {
	// Asset type. Can either be NOTEBOOK_FILE or JAR_ANALYSIS.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room
	CleanRoomName string `json:"-" url:"-"`
	// Name of the asset
	Name string `json:"-" url:"-"`

	NotebookReview *NotebookVersionReview `json:"notebook_review,omitempty"`
}

type CreateCleanRoomAssetReviewResponse struct {
	// Top-level status derived from all reviews
	NotebookReviewState CleanRoomNotebookReviewNotebookReviewState `json:"notebook_review_state,omitempty"`
	// All existing notebook approvals or rejections
	NotebookReviews []CleanRoomNotebookReview `json:"notebook_reviews,omitempty"`
}

type CreateCleanRoomAutoApprovalRuleRequest struct {
	AutoApprovalRule CleanRoomAutoApprovalRule `json:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName string `json:"-" url:"-"`
}

type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`

	OutputCatalog CleanRoomOutputCatalog `json:"output_catalog"`
}

type CreateCleanRoomOutputCatalogResponse struct {
	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
}

type CreateCleanRoomRequest struct {
	CleanRoom CleanRoom `json:"clean_room"`
}

type DeleteCleanRoomAssetRequest struct {
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name string `json:"-" url:"-"`
}

type DeleteCleanRoomAutoApprovalRuleRequest struct {
	CleanRoomName string `json:"-" url:"-"`

	RuleId string `json:"-" url:"-"`
}

type DeleteCleanRoomRequest struct {
	// Name of the clean room.
	Name string `json:"-" url:"-"`
}

type GetCleanRoomAssetRequest struct {
	// The type of the asset.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name string `json:"-" url:"-"`
}

type GetCleanRoomAssetRevisionRequest struct {
	// Asset type. Only NOTEBOOK_FILE is supported.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// Revision etag to fetch. If not provided, the latest revision will be
	// returned.
	Etag string `json:"-" url:"-"`
	// Name of the asset.
	Name string `json:"-" url:"-"`
}

type GetCleanRoomAutoApprovalRuleRequest struct {
	CleanRoomName string `json:"-" url:"-"`

	RuleId string `json:"-" url:"-"`
}

type GetCleanRoomRequest struct {
	Name string `json:"-" url:"-"`
}

type ListCleanRoomAssetRevisionsRequest struct {
	// Asset type. Only NOTEBOOK_FILE is supported.
	AssetType CleanRoomAssetAssetType `json:"-" url:"-"`
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// Name of the asset.
	Name string `json:"-" url:"-"`
	// Maximum number of asset revisions to return. Defaults to 10.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on the previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomAssetRevisionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomAssetRevisionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomAssetRevisionsResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Revisions []CleanRoomAsset `json:"revisions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomAssetRevisionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomAssetRevisionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomAssetsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomAssetsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomAssetsResponse struct {
	// Assets in the clean room.
	Assets []CleanRoomAsset `json:"assets,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomAssetsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomAssetsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomAutoApprovalRulesRequest struct {
	CleanRoomName string `json:"-" url:"-"`
	// Maximum number of auto-approval rules to return. Defaults to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomAutoApprovalRulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomAutoApprovalRulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomAutoApprovalRulesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	Rules []CleanRoomAutoApprovalRule `json:"rules,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomAutoApprovalRulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomAutoApprovalRulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomNotebookTaskRunsRequest struct {
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

func (s *ListCleanRoomNotebookTaskRunsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomNotebookTaskRunsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomNotebookTaskRunsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// Name of the clean room.
	Runs []CleanRoomNotebookTaskRun `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomNotebookTaskRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomNotebookTaskRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomsResponse struct {
	CleanRooms []CleanRoom `json:"clean_rooms,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCleanRoomsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NotebookVersionReview struct {
	// Review comment
	Comment string `json:"comment,omitempty"`
	// Etag identifying the notebook version
	Etag string `json:"etag"`
	// Review outcome
	ReviewState CleanRoomNotebookReviewNotebookReviewState `json:"review_state"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotebookVersionReview) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookVersionReview) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCleanRoomAssetRequest struct {
	// The asset to update. The asset's `name` and `asset_type` fields are used
	// to identify the asset to update.
	Asset CleanRoomAsset `json:"asset"`
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
	// For notebooks, the name is the notebook file name. For jar analyses, the
	// name is the jar analysis name.
	Name string `json:"-" url:"-"`
}

type UpdateCleanRoomAutoApprovalRuleRequest struct {
	// The auto-approval rule to update. The rule_id field is used to identify
	// the rule to update.
	AutoApprovalRule CleanRoomAutoApprovalRule `json:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName string `json:"-" url:"-"`
	// A generated UUID identifying the rule.
	RuleId string `json:"-" url:"-"`
}

type UpdateCleanRoomRequest struct {
	CleanRoom *CleanRoom `json:"clean_room,omitempty"`
	// Name of the clean room.
	Name string `json:"-" url:"-"`
}
