// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanroomspb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog/catalogpb"
	"github.com/databricks/databricks-sdk-go/service/jobs/jobspb"
	"github.com/databricks/databricks-sdk-go/service/settings/settingspb"
	"github.com/databricks/databricks-sdk-go/service/sharing/sharingpb"
)

type CleanRoomPb struct {
	AccessRestricted       CleanRoomAccessRestrictedPb `json:"access_restricted,omitempty"`
	Comment                string                      `json:"comment,omitempty"`
	CreatedAt              int64                       `json:"created_at,omitempty"`
	LocalCollaboratorAlias string                      `json:"local_collaborator_alias,omitempty"`
	Name                   string                      `json:"name,omitempty"`
	OutputCatalog          *CleanRoomOutputCatalogPb   `json:"output_catalog,omitempty"`
	Owner                  string                      `json:"owner,omitempty"`
	RemoteDetailedInfo     *CleanRoomRemoteDetailPb    `json:"remote_detailed_info,omitempty"`
	Status                 CleanRoomStatusEnumPb       `json:"status,omitempty"`
	UpdatedAt              int64                       `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAccessRestrictedPb string

const CleanRoomAccessRestrictedCspMismatch CleanRoomAccessRestrictedPb = `CSP_MISMATCH`

const CleanRoomAccessRestrictedNoRestriction CleanRoomAccessRestrictedPb = `NO_RESTRICTION`

type CleanRoomAssetPb struct {
	AddedAt                  int64                                     `json:"added_at,omitempty"`
	AssetType                CleanRoomAssetAssetTypePb                 `json:"asset_type"`
	CleanRoomName            string                                    `json:"clean_room_name,omitempty"`
	ForeignTable             *CleanRoomAssetForeignTablePb             `json:"foreign_table,omitempty"`
	ForeignTableLocalDetails *CleanRoomAssetForeignTableLocalDetailsPb `json:"foreign_table_local_details,omitempty"`
	Name                     string                                    `json:"name"`
	Notebook                 *CleanRoomAssetNotebookPb                 `json:"notebook,omitempty"`
	OwnerCollaboratorAlias   string                                    `json:"owner_collaborator_alias,omitempty"`
	Status                   CleanRoomAssetStatusEnumPb                `json:"status,omitempty"`
	Table                    *CleanRoomAssetTablePb                    `json:"table,omitempty"`
	TableLocalDetails        *CleanRoomAssetTableLocalDetailsPb        `json:"table_local_details,omitempty"`
	View                     *CleanRoomAssetViewPb                     `json:"view,omitempty"`
	ViewLocalDetails         *CleanRoomAssetViewLocalDetailsPb         `json:"view_local_details,omitempty"`
	VolumeLocalDetails       *CleanRoomAssetVolumeLocalDetailsPb       `json:"volume_local_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomAssetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomAssetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAssetAssetTypePb string

const CleanRoomAssetAssetTypeForeignTable CleanRoomAssetAssetTypePb = `FOREIGN_TABLE`

const CleanRoomAssetAssetTypeNotebookFile CleanRoomAssetAssetTypePb = `NOTEBOOK_FILE`

const CleanRoomAssetAssetTypeTable CleanRoomAssetAssetTypePb = `TABLE`

const CleanRoomAssetAssetTypeView CleanRoomAssetAssetTypePb = `VIEW`

const CleanRoomAssetAssetTypeVolume CleanRoomAssetAssetTypePb = `VOLUME`

type CleanRoomAssetForeignTablePb struct {
	Columns []catalogpb.ColumnInfoPb `json:"columns,omitempty"`
}

type CleanRoomAssetForeignTableLocalDetailsPb struct {
	LocalName string `json:"local_name"`
}

type CleanRoomAssetNotebookPb struct {
	Etag                      string                                       `json:"etag,omitempty"`
	NotebookContent           string                                       `json:"notebook_content"`
	ReviewState               CleanRoomNotebookReviewNotebookReviewStatePb `json:"review_state,omitempty"`
	Reviews                   []CleanRoomNotebookReviewPb                  `json:"reviews,omitempty"`
	RunnerCollaboratorAliases []string                                     `json:"runner_collaborator_aliases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomAssetNotebookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomAssetNotebookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAssetStatusEnumPb string

const CleanRoomAssetStatusEnumActive CleanRoomAssetStatusEnumPb = `ACTIVE`

const CleanRoomAssetStatusEnumPending CleanRoomAssetStatusEnumPb = `PENDING`

const CleanRoomAssetStatusEnumPermissionDenied CleanRoomAssetStatusEnumPb = `PERMISSION_DENIED`

type CleanRoomAssetTablePb struct {
	Columns []catalogpb.ColumnInfoPb `json:"columns,omitempty"`
}

type CleanRoomAssetTableLocalDetailsPb struct {
	LocalName  string                  `json:"local_name"`
	Partitions []sharingpb.PartitionPb `json:"partitions,omitempty"`
}

type CleanRoomAssetViewPb struct {
	Columns []catalogpb.ColumnInfoPb `json:"columns,omitempty"`
}

type CleanRoomAssetViewLocalDetailsPb struct {
	LocalName string `json:"local_name"`
}

type CleanRoomAssetVolumeLocalDetailsPb struct {
	LocalName string `json:"local_name"`
}

type CleanRoomAutoApprovalRulePb struct {
	AuthorCollaboratorAlias    string                                 `json:"author_collaborator_alias,omitempty"`
	AuthorScope                CleanRoomAutoApprovalRuleAuthorScopePb `json:"author_scope,omitempty"`
	CleanRoomName              string                                 `json:"clean_room_name,omitempty"`
	CreatedAt                  int64                                  `json:"created_at,omitempty"`
	RuleId                     string                                 `json:"rule_id,omitempty"`
	RuleOwnerCollaboratorAlias string                                 `json:"rule_owner_collaborator_alias,omitempty"`
	RunnerCollaboratorAlias    string                                 `json:"runner_collaborator_alias,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomAutoApprovalRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomAutoApprovalRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomAutoApprovalRuleAuthorScopePb string

const CleanRoomAutoApprovalRuleAuthorScopeAnyAuthor CleanRoomAutoApprovalRuleAuthorScopePb = `ANY_AUTHOR`

type CleanRoomCollaboratorPb struct {
	CollaboratorAlias          string `json:"collaborator_alias"`
	DisplayName                string `json:"display_name,omitempty"`
	GlobalMetastoreId          string `json:"global_metastore_id,omitempty"`
	InviteRecipientEmail       string `json:"invite_recipient_email,omitempty"`
	InviteRecipientWorkspaceId int64  `json:"invite_recipient_workspace_id,omitempty"`
	OrganizationName           string `json:"organization_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomCollaboratorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomCollaboratorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomNotebookReviewPb struct {
	Comment                   string                                           `json:"comment,omitempty"`
	CreatedAtMillis           int64                                            `json:"created_at_millis,omitempty"`
	ReviewState               CleanRoomNotebookReviewNotebookReviewStatePb     `json:"review_state,omitempty"`
	ReviewSubReason           CleanRoomNotebookReviewNotebookReviewSubReasonPb `json:"review_sub_reason,omitempty"`
	ReviewerCollaboratorAlias string                                           `json:"reviewer_collaborator_alias,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomNotebookReviewPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomNotebookReviewPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomNotebookReviewNotebookReviewStatePb string

const CleanRoomNotebookReviewNotebookReviewStateApproved CleanRoomNotebookReviewNotebookReviewStatePb = `APPROVED`

const CleanRoomNotebookReviewNotebookReviewStatePending CleanRoomNotebookReviewNotebookReviewStatePb = `PENDING`

const CleanRoomNotebookReviewNotebookReviewStateRejected CleanRoomNotebookReviewNotebookReviewStatePb = `REJECTED`

type CleanRoomNotebookReviewNotebookReviewSubReasonPb string

const CleanRoomNotebookReviewNotebookReviewSubReasonAutoApproved CleanRoomNotebookReviewNotebookReviewSubReasonPb = `AUTO_APPROVED`

const CleanRoomNotebookReviewNotebookReviewSubReasonBackfilled CleanRoomNotebookReviewNotebookReviewSubReasonPb = `BACKFILLED`

type CleanRoomNotebookTaskRunPb struct {
	CollaboratorJobRunInfo     *CollaboratorJobRunInfoPb       `json:"collaborator_job_run_info,omitempty"`
	NotebookEtag               string                          `json:"notebook_etag,omitempty"`
	NotebookJobRunState        *jobspb.CleanRoomTaskRunStatePb `json:"notebook_job_run_state,omitempty"`
	NotebookName               string                          `json:"notebook_name,omitempty"`
	NotebookUpdatedAt          int64                           `json:"notebook_updated_at,omitempty"`
	OutputSchemaExpirationTime int64                           `json:"output_schema_expiration_time,omitempty"`
	OutputSchemaName           string                          `json:"output_schema_name,omitempty"`
	RunDuration                int64                           `json:"run_duration,omitempty"`
	StartTime                  int64                           `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomNotebookTaskRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomNotebookTaskRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomOutputCatalogPb struct {
	CatalogName string                                      `json:"catalog_name,omitempty"`
	Status      CleanRoomOutputCatalogOutputCatalogStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomOutputCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomOutputCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomOutputCatalogOutputCatalogStatusPb string

const CleanRoomOutputCatalogOutputCatalogStatusCreated CleanRoomOutputCatalogOutputCatalogStatusPb = `CREATED`

const CleanRoomOutputCatalogOutputCatalogStatusNotCreated CleanRoomOutputCatalogOutputCatalogStatusPb = `NOT_CREATED`

const CleanRoomOutputCatalogOutputCatalogStatusNotEligible CleanRoomOutputCatalogOutputCatalogStatusPb = `NOT_ELIGIBLE`

type CleanRoomRemoteDetailPb struct {
	CentralCleanRoomId        string                            `json:"central_clean_room_id,omitempty"`
	CloudVendor               string                            `json:"cloud_vendor,omitempty"`
	Collaborators             []CleanRoomCollaboratorPb         `json:"collaborators,omitempty"`
	ComplianceSecurityProfile *ComplianceSecurityProfilePb      `json:"compliance_security_profile,omitempty"`
	Creator                   *CleanRoomCollaboratorPb          `json:"creator,omitempty"`
	EgressNetworkPolicy       *settingspb.EgressNetworkPolicyPb `json:"egress_network_policy,omitempty"`
	Region                    string                            `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomRemoteDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomRemoteDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomStatusEnumPb string

const CleanRoomStatusEnumActive CleanRoomStatusEnumPb = `ACTIVE`

const CleanRoomStatusEnumDeleted CleanRoomStatusEnumPb = `DELETED`

const CleanRoomStatusEnumFailed CleanRoomStatusEnumPb = `FAILED`

const CleanRoomStatusEnumProvisioning CleanRoomStatusEnumPb = `PROVISIONING`

type CollaboratorJobRunInfoPb struct {
	CollaboratorAlias       string `json:"collaborator_alias,omitempty"`
	CollaboratorJobId       int64  `json:"collaborator_job_id,omitempty"`
	CollaboratorJobRunId    int64  `json:"collaborator_job_run_id,omitempty"`
	CollaboratorTaskRunId   int64  `json:"collaborator_task_run_id,omitempty"`
	CollaboratorWorkspaceId int64  `json:"collaborator_workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CollaboratorJobRunInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CollaboratorJobRunInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComplianceSecurityProfilePb struct {
	ComplianceStandards []settingspb.ComplianceStandardPb `json:"compliance_standards,omitempty"`
	IsEnabled           bool                              `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ComplianceSecurityProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ComplianceSecurityProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCleanRoomAssetRequestPb struct {
	Asset         CleanRoomAssetPb `json:"asset"`
	CleanRoomName string           `json:"-" url:"-"`
}

type CreateCleanRoomAssetReviewRequestPb struct {
	AssetType      CleanRoomAssetAssetTypePb `json:"-" url:"-"`
	CleanRoomName  string                    `json:"-" url:"-"`
	Name           string                    `json:"-" url:"-"`
	NotebookReview NotebookVersionReviewPb   `json:"notebook_review"`
}

type CreateCleanRoomAssetReviewResponsePb struct {
	NotebookReviewState CleanRoomNotebookReviewNotebookReviewStatePb `json:"notebook_review_state,omitempty"`
	NotebookReviews     []CleanRoomNotebookReviewPb                  `json:"notebook_reviews,omitempty"`
}

type CreateCleanRoomAutoApprovalRuleRequestPb struct {
	AutoApprovalRule CleanRoomAutoApprovalRulePb `json:"auto_approval_rule"`
	CleanRoomName    string                      `json:"-" url:"-"`
}

type CreateCleanRoomOutputCatalogRequestPb struct {
	CleanRoomName string                   `json:"-" url:"-"`
	OutputCatalog CleanRoomOutputCatalogPb `json:"output_catalog"`
}

type CreateCleanRoomOutputCatalogResponsePb struct {
	OutputCatalog *CleanRoomOutputCatalogPb `json:"output_catalog,omitempty"`
}

type CreateCleanRoomRequestPb struct {
	CleanRoom CleanRoomPb `json:"clean_room"`
}

type DeleteCleanRoomAssetRequestPb struct {
	AssetType     CleanRoomAssetAssetTypePb `json:"-" url:"-"`
	CleanRoomName string                    `json:"-" url:"-"`
	Name          string                    `json:"-" url:"-"`
}

type DeleteCleanRoomAssetResponsePb struct {
}

type DeleteCleanRoomAutoApprovalRuleRequestPb struct {
	CleanRoomName string `json:"-" url:"-"`
	RuleId        string `json:"-" url:"-"`
}

type DeleteCleanRoomRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetCleanRoomAssetRequestPb struct {
	AssetType     CleanRoomAssetAssetTypePb `json:"-" url:"-"`
	CleanRoomName string                    `json:"-" url:"-"`
	Name          string                    `json:"-" url:"-"`
}

type GetCleanRoomAssetRevisionRequestPb struct {
	AssetType     CleanRoomAssetAssetTypePb `json:"-" url:"-"`
	CleanRoomName string                    `json:"-" url:"-"`
	Etag          string                    `json:"-" url:"-"`
	Name          string                    `json:"-" url:"-"`
}

type GetCleanRoomAutoApprovalRuleRequestPb struct {
	CleanRoomName string `json:"-" url:"-"`
	RuleId        string `json:"-" url:"-"`
}

type GetCleanRoomRequestPb struct {
	Name string `json:"-" url:"-"`
}

type ListCleanRoomAssetRevisionsRequestPb struct {
	AssetType     CleanRoomAssetAssetTypePb `json:"-" url:"-"`
	CleanRoomName string                    `json:"-" url:"-"`
	Name          string                    `json:"-" url:"-"`
	PageSize      int                       `json:"-" url:"page_size,omitempty"`
	PageToken     string                    `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomAssetRevisionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomAssetRevisionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomAssetRevisionsResponsePb struct {
	NextPageToken string             `json:"next_page_token,omitempty"`
	Revisions     []CleanRoomAssetPb `json:"revisions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomAssetRevisionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomAssetRevisionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomAssetsRequestPb struct {
	CleanRoomName string `json:"-" url:"-"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomAssetsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomAssetsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomAssetsResponsePb struct {
	Assets        []CleanRoomAssetPb `json:"assets,omitempty"`
	NextPageToken string             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomAssetsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomAssetsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomAutoApprovalRulesRequestPb struct {
	CleanRoomName string `json:"-" url:"-"`
	PageSize      int    `json:"-" url:"page_size,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomAutoApprovalRulesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomAutoApprovalRulesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomAutoApprovalRulesResponsePb struct {
	NextPageToken string                        `json:"next_page_token,omitempty"`
	Rules         []CleanRoomAutoApprovalRulePb `json:"rules,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomAutoApprovalRulesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomAutoApprovalRulesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomNotebookTaskRunsRequestPb struct {
	CleanRoomName string `json:"-" url:"-"`
	NotebookName  string `json:"-" url:"notebook_name,omitempty"`
	PageSize      int    `json:"-" url:"page_size,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomNotebookTaskRunsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomNotebookTaskRunsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomNotebookTaskRunsResponsePb struct {
	NextPageToken string                       `json:"next_page_token,omitempty"`
	Runs          []CleanRoomNotebookTaskRunPb `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomNotebookTaskRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomNotebookTaskRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCleanRoomsResponsePb struct {
	CleanRooms    []CleanRoomPb `json:"clean_rooms,omitempty"`
	NextPageToken string        `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCleanRoomsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCleanRoomsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotebookVersionReviewPb struct {
	Comment     string                                       `json:"comment,omitempty"`
	Etag        string                                       `json:"etag"`
	ReviewState CleanRoomNotebookReviewNotebookReviewStatePb `json:"review_state"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NotebookVersionReviewPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NotebookVersionReviewPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCleanRoomAssetRequestPb struct {
	Asset         CleanRoomAssetPb          `json:"asset"`
	AssetType     CleanRoomAssetAssetTypePb `json:"-" url:"-"`
	CleanRoomName string                    `json:"-" url:"-"`
	Name          string                    `json:"-" url:"-"`
}

type UpdateCleanRoomAutoApprovalRuleRequestPb struct {
	AutoApprovalRule CleanRoomAutoApprovalRulePb `json:"auto_approval_rule"`
	CleanRoomName    string                      `json:"-" url:"-"`
	RuleId           string                      `json:"-" url:"-"`
}

type UpdateCleanRoomRequestPb struct {
	CleanRoom *CleanRoomPb `json:"clean_room,omitempty"`
	Name      string       `json:"-" url:"-"`
}
