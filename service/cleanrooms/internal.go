// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
)

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
	pb.OutputCatalog = st.OutputCatalog
	pb.Owner = st.Owner
	pb.RemoteDetailedInfo = st.RemoteDetailedInfo
	pb.Status = st.Status
	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomPb struct {
	AccessRestricted       CleanRoomAccessRestricted `json:"access_restricted,omitempty"`
	Comment                string                    `json:"comment,omitempty"`
	CreatedAt              int64                     `json:"created_at,omitempty"`
	LocalCollaboratorAlias string                    `json:"local_collaborator_alias,omitempty"`
	Name                   string                    `json:"name,omitempty"`
	OutputCatalog          *CleanRoomOutputCatalog   `json:"output_catalog,omitempty"`
	Owner                  string                    `json:"owner,omitempty"`
	RemoteDetailedInfo     *CleanRoomRemoteDetail    `json:"remote_detailed_info,omitempty"`
	Status                 CleanRoomStatusEnum       `json:"status,omitempty"`
	UpdatedAt              int64                     `json:"updated_at,omitempty"`

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
	st.OutputCatalog = pb.OutputCatalog
	st.Owner = pb.Owner
	st.RemoteDetailedInfo = pb.RemoteDetailedInfo
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

func cleanRoomAssetToPb(st *CleanRoomAsset) (*cleanRoomAssetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetPb{}
	pb.AddedAt = st.AddedAt
	pb.AssetType = st.AssetType
	pb.CleanRoomName = st.CleanRoomName
	pb.ForeignTable = st.ForeignTable
	pb.ForeignTableLocalDetails = st.ForeignTableLocalDetails
	pb.Name = st.Name
	pb.Notebook = st.Notebook
	pb.OwnerCollaboratorAlias = st.OwnerCollaboratorAlias
	pb.Status = st.Status
	pb.Table = st.Table
	pb.TableLocalDetails = st.TableLocalDetails
	pb.View = st.View
	pb.ViewLocalDetails = st.ViewLocalDetails
	pb.VolumeLocalDetails = st.VolumeLocalDetails

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomAssetPb struct {
	AddedAt                  int64                                   `json:"added_at,omitempty"`
	AssetType                CleanRoomAssetAssetType                 `json:"asset_type,omitempty"`
	CleanRoomName            string                                  `json:"clean_room_name,omitempty"`
	ForeignTable             *CleanRoomAssetForeignTable             `json:"foreign_table,omitempty"`
	ForeignTableLocalDetails *CleanRoomAssetForeignTableLocalDetails `json:"foreign_table_local_details,omitempty"`
	Name                     string                                  `json:"name,omitempty"`
	Notebook                 *CleanRoomAssetNotebook                 `json:"notebook,omitempty"`
	OwnerCollaboratorAlias   string                                  `json:"owner_collaborator_alias,omitempty"`
	Status                   CleanRoomAssetStatusEnum                `json:"status,omitempty"`
	Table                    *CleanRoomAssetTable                    `json:"table,omitempty"`
	TableLocalDetails        *CleanRoomAssetTableLocalDetails        `json:"table_local_details,omitempty"`
	View                     *CleanRoomAssetView                     `json:"view,omitempty"`
	ViewLocalDetails         *CleanRoomAssetViewLocalDetails         `json:"view_local_details,omitempty"`
	VolumeLocalDetails       *CleanRoomAssetVolumeLocalDetails       `json:"volume_local_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetFromPb(pb *cleanRoomAssetPb) (*CleanRoomAsset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAsset{}
	st.AddedAt = pb.AddedAt
	st.AssetType = pb.AssetType
	st.CleanRoomName = pb.CleanRoomName
	st.ForeignTable = pb.ForeignTable
	st.ForeignTableLocalDetails = pb.ForeignTableLocalDetails
	st.Name = pb.Name
	st.Notebook = pb.Notebook
	st.OwnerCollaboratorAlias = pb.OwnerCollaboratorAlias
	st.Status = pb.Status
	st.Table = pb.Table
	st.TableLocalDetails = pb.TableLocalDetails
	st.View = pb.View
	st.ViewLocalDetails = pb.ViewLocalDetails
	st.VolumeLocalDetails = pb.VolumeLocalDetails

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cleanRoomAssetForeignTableToPb(st *CleanRoomAssetForeignTable) (*cleanRoomAssetForeignTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetForeignTablePb{}
	pb.Columns = st.Columns

	return pb, nil
}

type cleanRoomAssetForeignTablePb struct {
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func cleanRoomAssetForeignTableFromPb(pb *cleanRoomAssetForeignTablePb) (*CleanRoomAssetForeignTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetForeignTable{}
	st.Columns = pb.Columns

	return st, nil
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

type cleanRoomAssetForeignTableLocalDetailsPb struct {
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

func cleanRoomAssetNotebookToPb(st *CleanRoomAssetNotebook) (*cleanRoomAssetNotebookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetNotebookPb{}
	pb.Etag = st.Etag
	pb.NotebookContent = st.NotebookContent
	pb.ReviewState = st.ReviewState
	pb.Reviews = st.Reviews
	pb.RunnerCollaboratorAliases = st.RunnerCollaboratorAliases

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomAssetNotebookPb struct {
	Etag                      string                                     `json:"etag,omitempty"`
	NotebookContent           string                                     `json:"notebook_content,omitempty"`
	ReviewState               CleanRoomNotebookReviewNotebookReviewState `json:"review_state,omitempty"`
	Reviews                   []CleanRoomNotebookReview                  `json:"reviews,omitempty"`
	RunnerCollaboratorAliases []string                                   `json:"runner_collaborator_aliases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetNotebookFromPb(pb *cleanRoomAssetNotebookPb) (*CleanRoomAssetNotebook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetNotebook{}
	st.Etag = pb.Etag
	st.NotebookContent = pb.NotebookContent
	st.ReviewState = pb.ReviewState
	st.Reviews = pb.Reviews
	st.RunnerCollaboratorAliases = pb.RunnerCollaboratorAliases

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetNotebookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetNotebookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cleanRoomAssetTableToPb(st *CleanRoomAssetTable) (*cleanRoomAssetTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetTablePb{}
	pb.Columns = st.Columns

	return pb, nil
}

type cleanRoomAssetTablePb struct {
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func cleanRoomAssetTableFromPb(pb *cleanRoomAssetTablePb) (*CleanRoomAssetTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetTable{}
	st.Columns = pb.Columns

	return st, nil
}

func cleanRoomAssetTableLocalDetailsToPb(st *CleanRoomAssetTableLocalDetails) (*cleanRoomAssetTableLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetTableLocalDetailsPb{}
	pb.LocalName = st.LocalName
	pb.Partitions = st.Partitions

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomAssetTableLocalDetailsPb struct {
	LocalName  string              `json:"local_name,omitempty"`
	Partitions []sharing.Partition `json:"partitions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomAssetTableLocalDetailsFromPb(pb *cleanRoomAssetTableLocalDetailsPb) (*CleanRoomAssetTableLocalDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetTableLocalDetails{}
	st.LocalName = pb.LocalName
	st.Partitions = pb.Partitions

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomAssetTableLocalDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomAssetTableLocalDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cleanRoomAssetViewToPb(st *CleanRoomAssetView) (*cleanRoomAssetViewPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetViewPb{}
	pb.Columns = st.Columns

	return pb, nil
}

type cleanRoomAssetViewPb struct {
	Columns []catalog.ColumnInfo `json:"columns,omitempty"`
}

func cleanRoomAssetViewFromPb(pb *cleanRoomAssetViewPb) (*CleanRoomAssetView, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomAssetView{}
	st.Columns = pb.Columns

	return st, nil
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

type cleanRoomAssetViewLocalDetailsPb struct {
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

func cleanRoomAssetVolumeLocalDetailsToPb(st *CleanRoomAssetVolumeLocalDetails) (*cleanRoomAssetVolumeLocalDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomAssetVolumeLocalDetailsPb{}
	pb.LocalName = st.LocalName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomAssetVolumeLocalDetailsPb struct {
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

type cleanRoomCollaboratorPb struct {
	CollaboratorAlias          string `json:"collaborator_alias"`
	DisplayName                string `json:"display_name,omitempty"`
	GlobalMetastoreId          string `json:"global_metastore_id,omitempty"`
	InviteRecipientEmail       string `json:"invite_recipient_email,omitempty"`
	InviteRecipientWorkspaceId int64  `json:"invite_recipient_workspace_id,omitempty"`
	OrganizationName           string `json:"organization_name,omitempty"`

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

func cleanRoomNotebookReviewToPb(st *CleanRoomNotebookReview) (*cleanRoomNotebookReviewPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomNotebookReviewPb{}
	pb.Comment = st.Comment
	pb.CreatedAtMillis = st.CreatedAtMillis
	pb.ReviewState = st.ReviewState
	pb.ReviewSubReason = st.ReviewSubReason
	pb.ReviewerCollaboratorAlias = st.ReviewerCollaboratorAlias

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomNotebookReviewPb struct {
	Comment                   string                                         `json:"comment,omitempty"`
	CreatedAtMillis           int64                                          `json:"created_at_millis,omitempty"`
	ReviewState               CleanRoomNotebookReviewNotebookReviewState     `json:"review_state,omitempty"`
	ReviewSubReason           CleanRoomNotebookReviewNotebookReviewSubReason `json:"review_sub_reason,omitempty"`
	ReviewerCollaboratorAlias string                                         `json:"reviewer_collaborator_alias,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomNotebookReviewFromPb(pb *cleanRoomNotebookReviewPb) (*CleanRoomNotebookReview, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomNotebookReview{}
	st.Comment = pb.Comment
	st.CreatedAtMillis = pb.CreatedAtMillis
	st.ReviewState = pb.ReviewState
	st.ReviewSubReason = pb.ReviewSubReason
	st.ReviewerCollaboratorAlias = pb.ReviewerCollaboratorAlias

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomNotebookReviewPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomNotebookReviewPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cleanRoomNotebookTaskRunToPb(st *CleanRoomNotebookTaskRun) (*cleanRoomNotebookTaskRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomNotebookTaskRunPb{}
	pb.CollaboratorJobRunInfo = st.CollaboratorJobRunInfo
	pb.NotebookEtag = st.NotebookEtag
	pb.NotebookJobRunState = st.NotebookJobRunState
	pb.NotebookName = st.NotebookName
	pb.NotebookUpdatedAt = st.NotebookUpdatedAt
	pb.OutputSchemaExpirationTime = st.OutputSchemaExpirationTime
	pb.OutputSchemaName = st.OutputSchemaName
	pb.RunDuration = st.RunDuration
	pb.StartTime = st.StartTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomNotebookTaskRunPb struct {
	CollaboratorJobRunInfo     *CollaboratorJobRunInfo     `json:"collaborator_job_run_info,omitempty"`
	NotebookEtag               string                      `json:"notebook_etag,omitempty"`
	NotebookJobRunState        *jobs.CleanRoomTaskRunState `json:"notebook_job_run_state,omitempty"`
	NotebookName               string                      `json:"notebook_name,omitempty"`
	NotebookUpdatedAt          int64                       `json:"notebook_updated_at,omitempty"`
	OutputSchemaExpirationTime int64                       `json:"output_schema_expiration_time,omitempty"`
	OutputSchemaName           string                      `json:"output_schema_name,omitempty"`
	RunDuration                int64                       `json:"run_duration,omitempty"`
	StartTime                  int64                       `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomNotebookTaskRunFromPb(pb *cleanRoomNotebookTaskRunPb) (*CleanRoomNotebookTaskRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomNotebookTaskRun{}
	st.CollaboratorJobRunInfo = pb.CollaboratorJobRunInfo
	st.NotebookEtag = pb.NotebookEtag
	st.NotebookJobRunState = pb.NotebookJobRunState
	st.NotebookName = pb.NotebookName
	st.NotebookUpdatedAt = pb.NotebookUpdatedAt
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

type cleanRoomOutputCatalogPb struct {
	CatalogName string                                    `json:"catalog_name,omitempty"`
	Status      CleanRoomOutputCatalogOutputCatalogStatus `json:"status,omitempty"`

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

func cleanRoomRemoteDetailToPb(st *CleanRoomRemoteDetail) (*cleanRoomRemoteDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomRemoteDetailPb{}
	pb.CentralCleanRoomId = st.CentralCleanRoomId
	pb.CloudVendor = st.CloudVendor
	pb.Collaborators = st.Collaborators
	pb.ComplianceSecurityProfile = st.ComplianceSecurityProfile
	pb.Creator = st.Creator
	pb.EgressNetworkPolicy = st.EgressNetworkPolicy
	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomRemoteDetailPb struct {
	CentralCleanRoomId        string                        `json:"central_clean_room_id,omitempty"`
	CloudVendor               string                        `json:"cloud_vendor,omitempty"`
	Collaborators             []CleanRoomCollaborator       `json:"collaborators,omitempty"`
	ComplianceSecurityProfile *ComplianceSecurityProfile    `json:"compliance_security_profile,omitempty"`
	Creator                   *CleanRoomCollaborator        `json:"creator,omitempty"`
	EgressNetworkPolicy       *settings.EgressNetworkPolicy `json:"egress_network_policy,omitempty"`
	Region                    string                        `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomRemoteDetailFromPb(pb *cleanRoomRemoteDetailPb) (*CleanRoomRemoteDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomRemoteDetail{}
	st.CentralCleanRoomId = pb.CentralCleanRoomId
	st.CloudVendor = pb.CloudVendor
	st.Collaborators = pb.Collaborators
	st.ComplianceSecurityProfile = pb.ComplianceSecurityProfile
	st.Creator = pb.Creator
	st.EgressNetworkPolicy = pb.EgressNetworkPolicy
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

type collaboratorJobRunInfoPb struct {
	CollaboratorAlias       string `json:"collaborator_alias,omitempty"`
	CollaboratorJobId       int64  `json:"collaborator_job_id,omitempty"`
	CollaboratorJobRunId    int64  `json:"collaborator_job_run_id,omitempty"`
	CollaboratorTaskRunId   int64  `json:"collaborator_task_run_id,omitempty"`
	CollaboratorWorkspaceId int64  `json:"collaborator_workspace_id,omitempty"`

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

func complianceSecurityProfileToPb(st *ComplianceSecurityProfile) (*complianceSecurityProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &complianceSecurityProfilePb{}
	pb.ComplianceStandards = st.ComplianceStandards
	pb.IsEnabled = st.IsEnabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type complianceSecurityProfilePb struct {
	ComplianceStandards []settings.ComplianceStandard `json:"compliance_standards,omitempty"`
	IsEnabled           bool                          `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func complianceSecurityProfileFromPb(pb *complianceSecurityProfilePb) (*ComplianceSecurityProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplianceSecurityProfile{}
	st.ComplianceStandards = pb.ComplianceStandards
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

func createCleanRoomAssetRequestToPb(st *CreateCleanRoomAssetRequest) (*createCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomAssetRequestPb{}
	pb.Asset = st.Asset
	pb.CleanRoomName = st.CleanRoomName

	return pb, nil
}

type createCleanRoomAssetRequestPb struct {
	Asset         CleanRoomAsset `json:"asset"`
	CleanRoomName string         `json:"-" url:"-"`
}

func createCleanRoomAssetRequestFromPb(pb *createCleanRoomAssetRequestPb) (*CreateCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomAssetRequest{}
	st.Asset = pb.Asset
	st.CleanRoomName = pb.CleanRoomName

	return st, nil
}

func createCleanRoomOutputCatalogRequestToPb(st *CreateCleanRoomOutputCatalogRequest) (*createCleanRoomOutputCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomOutputCatalogRequestPb{}
	pb.CleanRoomName = st.CleanRoomName
	pb.OutputCatalog = st.OutputCatalog

	return pb, nil
}

type createCleanRoomOutputCatalogRequestPb struct {
	CleanRoomName string                 `json:"-" url:"-"`
	OutputCatalog CleanRoomOutputCatalog `json:"output_catalog"`
}

func createCleanRoomOutputCatalogRequestFromPb(pb *createCleanRoomOutputCatalogRequestPb) (*CreateCleanRoomOutputCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomOutputCatalogRequest{}
	st.CleanRoomName = pb.CleanRoomName
	st.OutputCatalog = pb.OutputCatalog

	return st, nil
}

func createCleanRoomOutputCatalogResponseToPb(st *CreateCleanRoomOutputCatalogResponse) (*createCleanRoomOutputCatalogResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomOutputCatalogResponsePb{}
	pb.OutputCatalog = st.OutputCatalog

	return pb, nil
}

type createCleanRoomOutputCatalogResponsePb struct {
	OutputCatalog *CleanRoomOutputCatalog `json:"output_catalog,omitempty"`
}

func createCleanRoomOutputCatalogResponseFromPb(pb *createCleanRoomOutputCatalogResponsePb) (*CreateCleanRoomOutputCatalogResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomOutputCatalogResponse{}
	st.OutputCatalog = pb.OutputCatalog

	return st, nil
}

func createCleanRoomRequestToPb(st *CreateCleanRoomRequest) (*createCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCleanRoomRequestPb{}
	pb.CleanRoom = st.CleanRoom

	return pb, nil
}

type createCleanRoomRequestPb struct {
	CleanRoom CleanRoom `json:"clean_room"`
}

func createCleanRoomRequestFromPb(pb *createCleanRoomRequestPb) (*CreateCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCleanRoomRequest{}
	st.CleanRoom = pb.CleanRoom

	return st, nil
}

func deleteCleanRoomAssetRequestToPb(st *DeleteCleanRoomAssetRequest) (*deleteCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCleanRoomAssetRequestPb{}
	pb.AssetType = st.AssetType
	pb.CleanRoomName = st.CleanRoomName
	pb.Name = st.Name

	return pb, nil
}

type deleteCleanRoomAssetRequestPb struct {
	AssetType     CleanRoomAssetAssetType `json:"-" url:"-"`
	CleanRoomName string                  `json:"-" url:"-"`
	Name          string                  `json:"-" url:"-"`
}

func deleteCleanRoomAssetRequestFromPb(pb *deleteCleanRoomAssetRequestPb) (*DeleteCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCleanRoomAssetRequest{}
	st.AssetType = pb.AssetType
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name

	return st, nil
}

func deleteCleanRoomAssetResponseToPb(st *DeleteCleanRoomAssetResponse) (*deleteCleanRoomAssetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCleanRoomAssetResponsePb{}

	return pb, nil
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

func deleteCleanRoomRequestToPb(st *DeleteCleanRoomRequest) (*deleteCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCleanRoomRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteCleanRoomRequestPb struct {
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

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
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

func getCleanRoomAssetRequestToPb(st *GetCleanRoomAssetRequest) (*getCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCleanRoomAssetRequestPb{}
	pb.AssetType = st.AssetType
	pb.CleanRoomName = st.CleanRoomName
	pb.Name = st.Name

	return pb, nil
}

type getCleanRoomAssetRequestPb struct {
	AssetType     CleanRoomAssetAssetType `json:"-" url:"-"`
	CleanRoomName string                  `json:"-" url:"-"`
	Name          string                  `json:"-" url:"-"`
}

func getCleanRoomAssetRequestFromPb(pb *getCleanRoomAssetRequestPb) (*GetCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCleanRoomAssetRequest{}
	st.AssetType = pb.AssetType
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name

	return st, nil
}

func getCleanRoomRequestToPb(st *GetCleanRoomRequest) (*getCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCleanRoomRequestPb{}
	pb.Name = st.Name

	return pb, nil
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

type listCleanRoomAssetsRequestPb struct {
	CleanRoomName string `json:"-" url:"-"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

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

func listCleanRoomAssetsResponseToPb(st *ListCleanRoomAssetsResponse) (*listCleanRoomAssetsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomAssetsResponsePb{}
	pb.Assets = st.Assets
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listCleanRoomAssetsResponsePb struct {
	Assets        []CleanRoomAsset `json:"assets,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomAssetsResponseFromPb(pb *listCleanRoomAssetsResponsePb) (*ListCleanRoomAssetsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomAssetsResponse{}
	st.Assets = pb.Assets
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

type listCleanRoomNotebookTaskRunsRequestPb struct {
	CleanRoomName string `json:"-" url:"-"`
	NotebookName  string `json:"-" url:"notebook_name,omitempty"`
	PageSize      int    `json:"-" url:"page_size,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

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

func listCleanRoomNotebookTaskRunsResponseToPb(st *ListCleanRoomNotebookTaskRunsResponse) (*listCleanRoomNotebookTaskRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomNotebookTaskRunsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Runs = st.Runs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listCleanRoomNotebookTaskRunsResponsePb struct {
	NextPageToken string                     `json:"next_page_token,omitempty"`
	Runs          []CleanRoomNotebookTaskRun `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomNotebookTaskRunsResponseFromPb(pb *listCleanRoomNotebookTaskRunsResponsePb) (*ListCleanRoomNotebookTaskRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomNotebookTaskRunsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Runs = pb.Runs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCleanRoomNotebookTaskRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCleanRoomNotebookTaskRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type listCleanRoomsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
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

func listCleanRoomsResponseToPb(st *ListCleanRoomsResponse) (*listCleanRoomsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCleanRoomsResponsePb{}
	pb.CleanRooms = st.CleanRooms
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listCleanRoomsResponsePb struct {
	CleanRooms    []CleanRoom `json:"clean_rooms,omitempty"`
	NextPageToken string      `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCleanRoomsResponseFromPb(pb *listCleanRoomsResponsePb) (*ListCleanRoomsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCleanRoomsResponse{}
	st.CleanRooms = pb.CleanRooms
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

func updateCleanRoomAssetRequestToPb(st *UpdateCleanRoomAssetRequest) (*updateCleanRoomAssetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCleanRoomAssetRequestPb{}
	pb.Asset = st.Asset
	pb.AssetType = st.AssetType
	pb.CleanRoomName = st.CleanRoomName
	pb.Name = st.Name

	return pb, nil
}

type updateCleanRoomAssetRequestPb struct {
	Asset         CleanRoomAsset          `json:"asset"`
	AssetType     CleanRoomAssetAssetType `json:"-" url:"-"`
	CleanRoomName string                  `json:"-" url:"-"`
	Name          string                  `json:"-" url:"-"`
}

func updateCleanRoomAssetRequestFromPb(pb *updateCleanRoomAssetRequestPb) (*UpdateCleanRoomAssetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCleanRoomAssetRequest{}
	st.Asset = pb.Asset
	st.AssetType = pb.AssetType
	st.CleanRoomName = pb.CleanRoomName
	st.Name = pb.Name

	return st, nil
}

func updateCleanRoomRequestToPb(st *UpdateCleanRoomRequest) (*updateCleanRoomRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCleanRoomRequestPb{}
	pb.CleanRoom = st.CleanRoom
	pb.Name = st.Name

	return pb, nil
}

type updateCleanRoomRequestPb struct {
	CleanRoom *CleanRoom `json:"clean_room,omitempty"`
	Name      string     `json:"-" url:"-"`
}

func updateCleanRoomRequestFromPb(pb *updateCleanRoomRequestPb) (*UpdateCleanRoomRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCleanRoomRequest{}
	st.CleanRoom = pb.CleanRoom
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
