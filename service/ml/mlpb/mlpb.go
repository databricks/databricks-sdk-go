// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlpb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type ActivityPb struct {
	ActivityType         ActivityTypePb `json:"activity_type,omitempty"`
	Comment              string         `json:"comment,omitempty"`
	CreationTimestamp    int64          `json:"creation_timestamp,omitempty"`
	FromStage            string         `json:"from_stage,omitempty"`
	Id                   string         `json:"id,omitempty"`
	LastUpdatedTimestamp int64          `json:"last_updated_timestamp,omitempty"`
	SystemComment        string         `json:"system_comment,omitempty"`
	ToStage              string         `json:"to_stage,omitempty"`
	UserId               string         `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ActivityPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ActivityPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ActivityActionPb string

// Approve a transition request
const ActivityActionApproveTransitionRequest ActivityActionPb = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityActionPb = `CANCEL_TRANSITION_REQUEST`

// Delete the comment
const ActivityActionDeleteComment ActivityActionPb = `DELETE_COMMENT`

// Edit the comment
const ActivityActionEditComment ActivityActionPb = `EDIT_COMMENT`

// Reject a transition request
const ActivityActionRejectTransitionRequest ActivityActionPb = `REJECT_TRANSITION_REQUEST`

type ActivityTypePb string

// User applied the corresponding stage transition.
const ActivityTypeAppliedTransition ActivityTypePb = `APPLIED_TRANSITION`

// User approved the corresponding stage transition.
const ActivityTypeApprovedRequest ActivityTypePb = `APPROVED_REQUEST`

// User cancelled an existing transition request.
const ActivityTypeCancelledRequest ActivityTypePb = `CANCELLED_REQUEST`

const ActivityTypeNewComment ActivityTypePb = `NEW_COMMENT`

// User rejected the coressponding stage transition.
const ActivityTypeRejectedRequest ActivityTypePb = `REJECTED_REQUEST`

// User requested the corresponding stage transition.
const ActivityTypeRequestedTransition ActivityTypePb = `REQUESTED_TRANSITION`

// For events performed as a side effect, such as archiving existing model
// versions in a stage.
const ActivityTypeSystemTransition ActivityTypePb = `SYSTEM_TRANSITION`

type ApproveTransitionRequestPb struct {
	ArchiveExistingVersions bool   `json:"archive_existing_versions"`
	Comment                 string `json:"comment,omitempty"`
	Name                    string `json:"name"`
	Stage                   string `json:"stage"`
	Version                 string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ApproveTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ApproveTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ApproveTransitionRequestResponsePb struct {
	Activity *ActivityPb `json:"activity,omitempty"`
}

type CommentActivityActionPb string

// Approve a transition request
const CommentActivityActionApproveTransitionRequest CommentActivityActionPb = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const CommentActivityActionCancelTransitionRequest CommentActivityActionPb = `CANCEL_TRANSITION_REQUEST`

// Delete the comment
const CommentActivityActionDeleteComment CommentActivityActionPb = `DELETE_COMMENT`

// Edit the comment
const CommentActivityActionEditComment CommentActivityActionPb = `EDIT_COMMENT`

// Reject a transition request
const CommentActivityActionRejectTransitionRequest CommentActivityActionPb = `REJECT_TRANSITION_REQUEST`

type CommentObjectPb struct {
	AvailableActions     []CommentActivityActionPb `json:"available_actions,omitempty"`
	Comment              string                    `json:"comment,omitempty"`
	CreationTimestamp    int64                     `json:"creation_timestamp,omitempty"`
	Id                   string                    `json:"id,omitempty"`
	LastUpdatedTimestamp int64                     `json:"last_updated_timestamp,omitempty"`
	UserId               string                    `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CommentObjectPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CommentObjectPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCommentPb struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CreateCommentResponsePb struct {
	Comment *CommentObjectPb `json:"comment,omitempty"`
}

type CreateExperimentPb struct {
	ArtifactLocation string            `json:"artifact_location,omitempty"`
	Name             string            `json:"name"`
	Tags             []ExperimentTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateExperimentResponsePb struct {
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateExperimentResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateExperimentResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateFeatureTagRequestPb struct {
	FeatureName string       `json:"-" url:"-"`
	FeatureTag  FeatureTagPb `json:"feature_tag"`
	TableName   string       `json:"-" url:"-"`
}

type CreateForecastingExperimentRequestPb struct {
	CustomWeightsColumn         string   `json:"custom_weights_column,omitempty"`
	ExperimentPath              string   `json:"experiment_path,omitempty"`
	ForecastGranularity         string   `json:"forecast_granularity"`
	ForecastHorizon             int64    `json:"forecast_horizon"`
	FutureFeatureDataPath       string   `json:"future_feature_data_path,omitempty"`
	HolidayRegions              []string `json:"holiday_regions,omitempty"`
	IncludeFeatures             []string `json:"include_features,omitempty"`
	MaxRuntime                  int64    `json:"max_runtime,omitempty"`
	PredictionDataPath          string   `json:"prediction_data_path,omitempty"`
	PrimaryMetric               string   `json:"primary_metric,omitempty"`
	RegisterTo                  string   `json:"register_to,omitempty"`
	SplitColumn                 string   `json:"split_column,omitempty"`
	TargetColumn                string   `json:"target_column"`
	TimeColumn                  string   `json:"time_column"`
	TimeseriesIdentifierColumns []string `json:"timeseries_identifier_columns,omitempty"`
	TrainDataPath               string   `json:"train_data_path"`
	TrainingFrameworks          []string `json:"training_frameworks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateForecastingExperimentRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateForecastingExperimentRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateForecastingExperimentResponsePb struct {
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateForecastingExperimentResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateForecastingExperimentResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateLoggedModelRequestPb struct {
	ExperimentId string                   `json:"experiment_id"`
	ModelType    string                   `json:"model_type,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Params       []LoggedModelParameterPb `json:"params,omitempty"`
	SourceRunId  string                   `json:"source_run_id,omitempty"`
	Tags         []LoggedModelTagPb       `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateLoggedModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateLoggedModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateLoggedModelResponsePb struct {
	Model *LoggedModelPb `json:"model,omitempty"`
}

type CreateModelRequestPb struct {
	Description string       `json:"description,omitempty"`
	Name        string       `json:"name"`
	Tags        []ModelTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateModelResponsePb struct {
	RegisteredModel *ModelPb `json:"registered_model,omitempty"`
}

type CreateModelVersionRequestPb struct {
	Description string              `json:"description,omitempty"`
	Name        string              `json:"name"`
	RunId       string              `json:"run_id,omitempty"`
	RunLink     string              `json:"run_link,omitempty"`
	Source      string              `json:"source"`
	Tags        []ModelVersionTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateModelVersionResponsePb struct {
	ModelVersion *ModelVersionPb `json:"model_version,omitempty"`
}

type CreateOnlineStoreRequestPb struct {
	OnlineStore OnlineStorePb `json:"online_store"`
}

type CreateRegistryWebhookPb struct {
	Description string                   `json:"description,omitempty"`
	Events      []RegistryWebhookEventPb `json:"events"`
	HttpUrlSpec *HttpUrlSpecPb           `json:"http_url_spec,omitempty"`
	JobSpec     *JobSpecPb               `json:"job_spec,omitempty"`
	ModelName   string                   `json:"model_name,omitempty"`
	Status      RegistryWebhookStatusPb  `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateRegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateRegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRunPb struct {
	ExperimentId string     `json:"experiment_id,omitempty"`
	RunName      string     `json:"run_name,omitempty"`
	StartTime    int64      `json:"start_time,omitempty"`
	Tags         []RunTagPb `json:"tags,omitempty"`
	UserId       string     `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRunResponsePb struct {
	Run *RunPb `json:"run,omitempty"`
}

type CreateTransitionRequestPb struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name"`
	Stage   string `json:"stage"`
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateTransitionRequestResponsePb struct {
	Request *TransitionRequestPb `json:"request,omitempty"`
}

type CreateWebhookResponsePb struct {
	Webhook *RegistryWebhookPb `json:"webhook,omitempty"`
}

type DatasetPb struct {
	Digest     string `json:"digest"`
	Name       string `json:"name"`
	Profile    string `json:"profile,omitempty"`
	Schema     string `json:"schema,omitempty"`
	Source     string `json:"source"`
	SourceType string `json:"source_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatasetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatasetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatasetInputPb struct {
	Dataset DatasetPb    `json:"dataset"`
	Tags    []InputTagPb `json:"tags,omitempty"`
}

type DeleteCommentRequestPb struct {
	Id string `json:"-" url:"id"`
}

type DeleteCommentResponsePb struct {
}

type DeleteExperimentPb struct {
	ExperimentId string `json:"experiment_id"`
}

type DeleteExperimentResponsePb struct {
}

type DeleteFeatureTagRequestPb struct {
	FeatureName string `json:"-" url:"-"`
	Key         string `json:"-" url:"-"`
	TableName   string `json:"-" url:"-"`
}

type DeleteLoggedModelRequestPb struct {
	ModelId string `json:"-" url:"-"`
}

type DeleteLoggedModelResponsePb struct {
}

type DeleteLoggedModelTagRequestPb struct {
	ModelId string `json:"-" url:"-"`
	TagKey  string `json:"-" url:"-"`
}

type DeleteLoggedModelTagResponsePb struct {
}

type DeleteModelRequestPb struct {
	Name string `json:"-" url:"name"`
}

type DeleteModelResponsePb struct {
}

type DeleteModelTagRequestPb struct {
	Key  string `json:"-" url:"key"`
	Name string `json:"-" url:"name"`
}

type DeleteModelTagResponsePb struct {
}

type DeleteModelVersionRequestPb struct {
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

type DeleteModelVersionResponsePb struct {
}

type DeleteModelVersionTagRequestPb struct {
	Key     string `json:"-" url:"key"`
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

type DeleteModelVersionTagResponsePb struct {
}

type DeleteOnlineStoreRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeleteRunPb struct {
	RunId string `json:"run_id"`
}

type DeleteRunResponsePb struct {
}

type DeleteRunsPb struct {
	ExperimentId       string `json:"experiment_id"`
	MaxRuns            int    `json:"max_runs,omitempty"`
	MaxTimestampMillis int64  `json:"max_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteRunsResponsePb struct {
	RunsDeleted int `json:"runs_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteTagPb struct {
	Key   string `json:"key"`
	RunId string `json:"run_id"`
}

type DeleteTagResponsePb struct {
}

type DeleteTransitionRequestRequestPb struct {
	Comment string `json:"-" url:"comment,omitempty"`
	Creator string `json:"-" url:"creator"`
	Name    string `json:"-" url:"name"`
	Stage   string `json:"-" url:"stage"`
	Version string `json:"-" url:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteTransitionRequestRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteTransitionRequestRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteTransitionRequestResponsePb struct {
	Activity *ActivityPb `json:"activity,omitempty"`
}

type DeleteWebhookRequestPb struct {
	Id string `json:"-" url:"id"`
}

type DeleteWebhookResponsePb struct {
}

type ExperimentPb struct {
	ArtifactLocation string            `json:"artifact_location,omitempty"`
	CreationTime     int64             `json:"creation_time,omitempty"`
	ExperimentId     string            `json:"experiment_id,omitempty"`
	LastUpdateTime   int64             `json:"last_update_time,omitempty"`
	LifecycleStage   string            `json:"lifecycle_stage,omitempty"`
	Name             string            `json:"name,omitempty"`
	Tags             []ExperimentTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentAccessControlRequestPb struct {
	GroupName            string                      `json:"group_name,omitempty"`
	PermissionLevel      ExperimentPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                      `json:"service_principal_name,omitempty"`
	UserName             string                      `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExperimentAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExperimentAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentAccessControlResponsePb struct {
	AllPermissions       []ExperimentPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                   `json:"display_name,omitempty"`
	GroupName            string                   `json:"group_name,omitempty"`
	ServicePrincipalName string                   `json:"service_principal_name,omitempty"`
	UserName             string                   `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExperimentAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExperimentAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentPermissionPb struct {
	Inherited           bool                        `json:"inherited,omitempty"`
	InheritedFromObject []string                    `json:"inherited_from_object,omitempty"`
	PermissionLevel     ExperimentPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExperimentPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExperimentPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentPermissionLevelPb string

const ExperimentPermissionLevelCanEdit ExperimentPermissionLevelPb = `CAN_EDIT`

const ExperimentPermissionLevelCanManage ExperimentPermissionLevelPb = `CAN_MANAGE`

const ExperimentPermissionLevelCanRead ExperimentPermissionLevelPb = `CAN_READ`

type ExperimentPermissionsPb struct {
	AccessControlList []ExperimentAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                              `json:"object_id,omitempty"`
	ObjectType        string                              `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExperimentPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExperimentPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentPermissionsDescriptionPb struct {
	Description     string                      `json:"description,omitempty"`
	PermissionLevel ExperimentPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExperimentPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExperimentPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentPermissionsRequestPb struct {
	AccessControlList []ExperimentAccessControlRequestPb `json:"access_control_list,omitempty"`
	ExperimentId      string                             `json:"-" url:"-"`
}

type ExperimentTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExperimentTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExperimentTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FeaturePb struct {
	FeatureName      string `json:"feature_name,omitempty"`
	FeatureTableId   string `json:"feature_table_id,omitempty"`
	FeatureTableName string `json:"feature_table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FeaturePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FeaturePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FeatureLineagePb struct {
	FeatureSpecs   []FeatureLineageFeatureSpecPb   `json:"feature_specs,omitempty"`
	Models         []FeatureLineageModelPb         `json:"models,omitempty"`
	OnlineFeatures []FeatureLineageOnlineFeaturePb `json:"online_features,omitempty"`
}

type FeatureLineageFeatureSpecPb struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FeatureLineageFeatureSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FeatureLineageFeatureSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FeatureLineageModelPb struct {
	Name    string `json:"name,omitempty"`
	Version int64  `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FeatureLineageModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FeatureLineageModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FeatureLineageOnlineFeaturePb struct {
	FeatureName string `json:"feature_name,omitempty"`
	TableName   string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FeatureLineageOnlineFeaturePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FeatureLineageOnlineFeaturePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FeatureListPb struct {
	Features []FeaturePb `json:"features,omitempty"`
}

type FeatureTagPb struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FeatureTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FeatureTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileInfoPb struct {
	FileSize int64  `json:"file_size,omitempty"`
	IsDir    bool   `json:"is_dir,omitempty"`
	Path     string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FinalizeLoggedModelRequestPb struct {
	ModelId string              `json:"-" url:"-"`
	Status  LoggedModelStatusPb `json:"status"`
}

type FinalizeLoggedModelResponsePb struct {
	Model *LoggedModelPb `json:"model,omitempty"`
}

type ForecastingExperimentPb struct {
	ExperimentId      string                       `json:"experiment_id,omitempty"`
	ExperimentPageUrl string                       `json:"experiment_page_url,omitempty"`
	State             ForecastingExperimentStatePb `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ForecastingExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ForecastingExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForecastingExperimentStatePb string

const ForecastingExperimentStateCancelled ForecastingExperimentStatePb = `CANCELLED`

const ForecastingExperimentStateFailed ForecastingExperimentStatePb = `FAILED`

const ForecastingExperimentStatePending ForecastingExperimentStatePb = `PENDING`

const ForecastingExperimentStateRunning ForecastingExperimentStatePb = `RUNNING`

const ForecastingExperimentStateSucceeded ForecastingExperimentStatePb = `SUCCEEDED`

type GetByNameRequestPb struct {
	ExperimentName string `json:"-" url:"experiment_name"`
}

type GetExperimentByNameResponsePb struct {
	Experiment *ExperimentPb `json:"experiment,omitempty"`
}

type GetExperimentPermissionLevelsRequestPb struct {
	ExperimentId string `json:"-" url:"-"`
}

type GetExperimentPermissionLevelsResponsePb struct {
	PermissionLevels []ExperimentPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetExperimentPermissionsRequestPb struct {
	ExperimentId string `json:"-" url:"-"`
}

type GetExperimentRequestPb struct {
	ExperimentId string `json:"-" url:"experiment_id"`
}

type GetExperimentResponsePb struct {
	Experiment *ExperimentPb `json:"experiment,omitempty"`
}

type GetFeatureLineageRequestPb struct {
	FeatureName string `json:"-" url:"-"`
	TableName   string `json:"-" url:"-"`
}

type GetFeatureTagRequestPb struct {
	FeatureName string `json:"-" url:"-"`
	Key         string `json:"-" url:"-"`
	TableName   string `json:"-" url:"-"`
}

type GetForecastingExperimentRequestPb struct {
	ExperimentId string `json:"-" url:"-"`
}

type GetHistoryRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	MetricKey  string `json:"-" url:"metric_key"`
	PageToken  string `json:"-" url:"page_token,omitempty"`
	RunId      string `json:"-" url:"run_id,omitempty"`
	RunUuid    string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetHistoryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetHistoryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetLatestVersionsRequestPb struct {
	Name   string   `json:"name"`
	Stages []string `json:"stages,omitempty"`
}

type GetLatestVersionsResponsePb struct {
	ModelVersions []ModelVersionPb `json:"model_versions,omitempty"`
}

type GetLoggedModelRequestPb struct {
	ModelId string `json:"-" url:"-"`
}

type GetLoggedModelResponsePb struct {
	Model *LoggedModelPb `json:"model,omitempty"`
}

type GetMetricHistoryResponsePb struct {
	Metrics       []MetricPb `json:"metrics,omitempty"`
	NextPageToken string     `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetMetricHistoryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetMetricHistoryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetModelRequestPb struct {
	Name string `json:"-" url:"name"`
}

type GetModelResponsePb struct {
	RegisteredModelDatabricks *ModelDatabricksPb `json:"registered_model_databricks,omitempty"`
}

type GetModelVersionDownloadUriRequestPb struct {
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

type GetModelVersionDownloadUriResponsePb struct {
	ArtifactUri string `json:"artifact_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetModelVersionDownloadUriResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetModelVersionDownloadUriResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetModelVersionRequestPb struct {
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

type GetModelVersionResponsePb struct {
	ModelVersion *ModelVersionPb `json:"model_version,omitempty"`
}

type GetOnlineStoreRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetRegisteredModelPermissionLevelsRequestPb struct {
	RegisteredModelId string `json:"-" url:"-"`
}

type GetRegisteredModelPermissionLevelsResponsePb struct {
	PermissionLevels []RegisteredModelPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetRegisteredModelPermissionsRequestPb struct {
	RegisteredModelId string `json:"-" url:"-"`
}

type GetRunRequestPb struct {
	RunId   string `json:"-" url:"run_id"`
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetRunRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetRunRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetRunResponsePb struct {
	Run *RunPb `json:"run,omitempty"`
}

type HttpUrlSpecPb struct {
	Authorization         string `json:"authorization,omitempty"`
	EnableSslVerification bool   `json:"enable_ssl_verification,omitempty"`
	Secret                string `json:"secret,omitempty"`
	Url                   string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *HttpUrlSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st HttpUrlSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type HttpUrlSpecWithoutSecretPb struct {
	EnableSslVerification bool   `json:"enable_ssl_verification,omitempty"`
	Url                   string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *HttpUrlSpecWithoutSecretPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st HttpUrlSpecWithoutSecretPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InputTagPb struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type JobSpecPb struct {
	AccessToken  string `json:"access_token"`
	JobId        string `json:"job_id"`
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobSpecWithoutSecretPb struct {
	JobId        string `json:"job_id,omitempty"`
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobSpecWithoutSecretPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobSpecWithoutSecretPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListArtifactsRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`
	Path      string `json:"-" url:"path,omitempty"`
	RunId     string `json:"-" url:"run_id,omitempty"`
	RunUuid   string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListArtifactsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListArtifactsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListArtifactsResponsePb struct {
	Files         []FileInfoPb `json:"files,omitempty"`
	NextPageToken string       `json:"next_page_token,omitempty"`
	RootUri       string       `json:"root_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListArtifactsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListArtifactsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExperimentsRequestPb struct {
	MaxResults int64      `json:"-" url:"max_results,omitempty"`
	PageToken  string     `json:"-" url:"page_token,omitempty"`
	ViewType   ViewTypePb `json:"-" url:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExperimentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExperimentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExperimentsResponsePb struct {
	Experiments   []ExperimentPb `json:"experiments,omitempty"`
	NextPageToken string         `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExperimentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExperimentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFeatureTagsRequestPb struct {
	FeatureName string `json:"-" url:"-"`
	PageSize    int    `json:"-" url:"page_size,omitempty"`
	PageToken   string `json:"-" url:"page_token,omitempty"`
	TableName   string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFeatureTagsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFeatureTagsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFeatureTagsResponsePb struct {
	FeatureTags   []FeatureTagPb `json:"feature_tags,omitempty"`
	NextPageToken string         `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFeatureTagsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFeatureTagsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListModelsRequestPb struct {
	MaxResults int64  `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListModelsResponsePb struct {
	NextPageToken    string    `json:"next_page_token,omitempty"`
	RegisteredModels []ModelPb `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListOnlineStoresRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListOnlineStoresRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListOnlineStoresRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListOnlineStoresResponsePb struct {
	NextPageToken string          `json:"next_page_token,omitempty"`
	OnlineStores  []OnlineStorePb `json:"online_stores,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListOnlineStoresResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListOnlineStoresResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRegistryWebhooksPb struct {
	NextPageToken string              `json:"next_page_token,omitempty"`
	Webhooks      []RegistryWebhookPb `json:"webhooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListRegistryWebhooksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListRegistryWebhooksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListTransitionRequestsRequestPb struct {
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

type ListTransitionRequestsResponsePb struct {
	Requests []ActivityPb `json:"requests,omitempty"`
}

type ListWebhooksRequestPb struct {
	Events     []RegistryWebhookEventPb `json:"-" url:"events,omitempty"`
	MaxResults int64                    `json:"-" url:"max_results,omitempty"`
	ModelName  string                   `json:"-" url:"model_name,omitempty"`
	PageToken  string                   `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListWebhooksRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListWebhooksRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogBatchPb struct {
	Metrics []MetricPb `json:"metrics,omitempty"`
	Params  []ParamPb  `json:"params,omitempty"`
	RunId   string     `json:"run_id,omitempty"`
	Tags    []RunTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LogBatchPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LogBatchPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogBatchResponsePb struct {
}

type LogInputsPb struct {
	Datasets []DatasetInputPb `json:"datasets,omitempty"`
	Models   []ModelInputPb   `json:"models,omitempty"`
	RunId    string           `json:"run_id"`
}

type LogInputsResponsePb struct {
}

type LogLoggedModelParamsRequestPb struct {
	ModelId string                   `json:"-" url:"-"`
	Params  []LoggedModelParameterPb `json:"params,omitempty"`
}

type LogLoggedModelParamsRequestResponsePb struct {
}

type LogMetricPb struct {
	DatasetDigest string  `json:"dataset_digest,omitempty"`
	DatasetName   string  `json:"dataset_name,omitempty"`
	Key           string  `json:"key"`
	ModelId       string  `json:"model_id,omitempty"`
	RunId         string  `json:"run_id,omitempty"`
	RunUuid       string  `json:"run_uuid,omitempty"`
	Step          int64   `json:"step,omitempty"`
	Timestamp     int64   `json:"timestamp"`
	Value         float64 `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LogMetricPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LogMetricPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogMetricResponsePb struct {
}

type LogModelPb struct {
	ModelJson string `json:"model_json,omitempty"`
	RunId     string `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LogModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LogModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogModelResponsePb struct {
}

type LogOutputsRequestPb struct {
	Models []ModelOutputPb `json:"models,omitempty"`
	RunId  string          `json:"run_id"`
}

type LogOutputsResponsePb struct {
}

type LogParamPb struct {
	Key     string `json:"key"`
	RunId   string `json:"run_id,omitempty"`
	RunUuid string `json:"run_uuid,omitempty"`
	Value   string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LogParamPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LogParamPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogParamResponsePb struct {
}

type LoggedModelPb struct {
	Data *LoggedModelDataPb `json:"data,omitempty"`
	Info *LoggedModelInfoPb `json:"info,omitempty"`
}

type LoggedModelDataPb struct {
	Metrics []MetricPb               `json:"metrics,omitempty"`
	Params  []LoggedModelParameterPb `json:"params,omitempty"`
}

type LoggedModelInfoPb struct {
	ArtifactUri            string              `json:"artifact_uri,omitempty"`
	CreationTimestampMs    int64               `json:"creation_timestamp_ms,omitempty"`
	CreatorId              int64               `json:"creator_id,omitempty"`
	ExperimentId           string              `json:"experiment_id,omitempty"`
	LastUpdatedTimestampMs int64               `json:"last_updated_timestamp_ms,omitempty"`
	ModelId                string              `json:"model_id,omitempty"`
	ModelType              string              `json:"model_type,omitempty"`
	Name                   string              `json:"name,omitempty"`
	SourceRunId            string              `json:"source_run_id,omitempty"`
	Status                 LoggedModelStatusPb `json:"status,omitempty"`
	StatusMessage          string              `json:"status_message,omitempty"`
	Tags                   []LoggedModelTagPb  `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LoggedModelInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LoggedModelInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LoggedModelParameterPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LoggedModelParameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LoggedModelParameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LoggedModelStatusPb string

const LoggedModelStatusLoggedModelPending LoggedModelStatusPb = `LOGGED_MODEL_PENDING`

const LoggedModelStatusLoggedModelReady LoggedModelStatusPb = `LOGGED_MODEL_READY`

const LoggedModelStatusLoggedModelUploadFailed LoggedModelStatusPb = `LOGGED_MODEL_UPLOAD_FAILED`

type LoggedModelTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LoggedModelTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LoggedModelTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MetricPb struct {
	DatasetDigest string  `json:"dataset_digest,omitempty"`
	DatasetName   string  `json:"dataset_name,omitempty"`
	Key           string  `json:"key,omitempty"`
	ModelId       string  `json:"model_id,omitempty"`
	RunId         string  `json:"run_id,omitempty"`
	Step          int64   `json:"step,omitempty"`
	Timestamp     int64   `json:"timestamp,omitempty"`
	Value         float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MetricPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MetricPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelPb struct {
	CreationTimestamp    int64            `json:"creation_timestamp,omitempty"`
	Description          string           `json:"description,omitempty"`
	LastUpdatedTimestamp int64            `json:"last_updated_timestamp,omitempty"`
	LatestVersions       []ModelVersionPb `json:"latest_versions,omitempty"`
	Name                 string           `json:"name,omitempty"`
	Tags                 []ModelTagPb     `json:"tags,omitempty"`
	UserId               string           `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelDatabricksPb struct {
	CreationTimestamp    int64             `json:"creation_timestamp,omitempty"`
	Description          string            `json:"description,omitempty"`
	Id                   string            `json:"id,omitempty"`
	LastUpdatedTimestamp int64             `json:"last_updated_timestamp,omitempty"`
	LatestVersions       []ModelVersionPb  `json:"latest_versions,omitempty"`
	Name                 string            `json:"name,omitempty"`
	PermissionLevel      PermissionLevelPb `json:"permission_level,omitempty"`
	Tags                 []ModelTagPb      `json:"tags,omitempty"`
	UserId               string            `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ModelDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ModelDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelInputPb struct {
	ModelId string `json:"model_id"`
}

type ModelOutputPb struct {
	ModelId string `json:"model_id"`
	Step    int64  `json:"step"`
}

type ModelTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ModelTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ModelTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelVersionPb struct {
	CreationTimestamp    int64                `json:"creation_timestamp,omitempty"`
	CurrentStage         string               `json:"current_stage,omitempty"`
	Description          string               `json:"description,omitempty"`
	LastUpdatedTimestamp int64                `json:"last_updated_timestamp,omitempty"`
	Name                 string               `json:"name,omitempty"`
	RunId                string               `json:"run_id,omitempty"`
	RunLink              string               `json:"run_link,omitempty"`
	Source               string               `json:"source,omitempty"`
	Status               ModelVersionStatusPb `json:"status,omitempty"`
	StatusMessage        string               `json:"status_message,omitempty"`
	Tags                 []ModelVersionTagPb  `json:"tags,omitempty"`
	UserId               string               `json:"user_id,omitempty"`
	Version              string               `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ModelVersionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ModelVersionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelVersionDatabricksPb struct {
	CreationTimestamp       int64                           `json:"creation_timestamp,omitempty"`
	CurrentStage            string                          `json:"current_stage,omitempty"`
	Description             string                          `json:"description,omitempty"`
	EmailSubscriptionStatus RegistryEmailSubscriptionTypePb `json:"email_subscription_status,omitempty"`
	FeatureList             *FeatureListPb                  `json:"feature_list,omitempty"`
	LastUpdatedTimestamp    int64                           `json:"last_updated_timestamp,omitempty"`
	Name                    string                          `json:"name,omitempty"`
	OpenRequests            []ActivityPb                    `json:"open_requests,omitempty"`
	PermissionLevel         PermissionLevelPb               `json:"permission_level,omitempty"`
	RunId                   string                          `json:"run_id,omitempty"`
	RunLink                 string                          `json:"run_link,omitempty"`
	Source                  string                          `json:"source,omitempty"`
	Status                  StatusPb                        `json:"status,omitempty"`
	StatusMessage           string                          `json:"status_message,omitempty"`
	Tags                    []ModelVersionTagPb             `json:"tags,omitempty"`
	UserId                  string                          `json:"user_id,omitempty"`
	Version                 string                          `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ModelVersionDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ModelVersionDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelVersionStatusPb string

// Request to register a new model version has failed.
const ModelVersionStatusFailedRegistration ModelVersionStatusPb = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const ModelVersionStatusPendingRegistration ModelVersionStatusPb = `PENDING_REGISTRATION`

// Model version is ready for use.
const ModelVersionStatusReady ModelVersionStatusPb = `READY`

type ModelVersionTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ModelVersionTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ModelVersionTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OnlineStorePb struct {
	Capacity         string             `json:"capacity"`
	CreationTime     string             `json:"creation_time,omitempty"`
	Creator          string             `json:"creator,omitempty"`
	Name             string             `json:"name"`
	ReadReplicaCount int                `json:"read_replica_count,omitempty"`
	State            OnlineStoreStatePb `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OnlineStorePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OnlineStorePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OnlineStoreStatePb string

const OnlineStoreStateAvailable OnlineStoreStatePb = `AVAILABLE`

const OnlineStoreStateDeleting OnlineStoreStatePb = `DELETING`

const OnlineStoreStateFailingOver OnlineStoreStatePb = `FAILING_OVER`

const OnlineStoreStateStarting OnlineStoreStatePb = `STARTING`

const OnlineStoreStateStopped OnlineStoreStatePb = `STOPPED`

const OnlineStoreStateUpdating OnlineStoreStatePb = `UPDATING`

type ParamPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ParamPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ParamPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionLevelPb string

const PermissionLevelCanCreateRegisteredModel PermissionLevelPb = `CAN_CREATE_REGISTERED_MODEL`

const PermissionLevelCanEdit PermissionLevelPb = `CAN_EDIT`

const PermissionLevelCanManage PermissionLevelPb = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevelPb = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageStagingVersions PermissionLevelPb = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanRead PermissionLevelPb = `CAN_READ`

type PublishSpecPb struct {
	OnlineStore     string                   `json:"online_store"`
	OnlineTableName string                   `json:"online_table_name"`
	PublishMode     PublishSpecPublishModePb `json:"publish_mode,omitempty"`
}

type PublishSpecPublishModePb string

const PublishSpecPublishModeContinuous PublishSpecPublishModePb = `CONTINUOUS`

const PublishSpecPublishModeTriggered PublishSpecPublishModePb = `TRIGGERED`

type PublishTableRequestPb struct {
	PublishSpec     PublishSpecPb `json:"publish_spec"`
	SourceTableName string        `json:"-" url:"-"`
}

type PublishTableResponsePb struct {
	OnlineTableName string `json:"online_table_name,omitempty"`
	PipelineId      string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PublishTableResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PublishTableResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelAccessControlRequestPb struct {
	GroupName            string                           `json:"group_name,omitempty"`
	PermissionLevel      RegisteredModelPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                           `json:"service_principal_name,omitempty"`
	UserName             string                           `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegisteredModelAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegisteredModelAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelAccessControlResponsePb struct {
	AllPermissions       []RegisteredModelPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                        `json:"display_name,omitempty"`
	GroupName            string                        `json:"group_name,omitempty"`
	ServicePrincipalName string                        `json:"service_principal_name,omitempty"`
	UserName             string                        `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegisteredModelAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegisteredModelAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelPermissionPb struct {
	Inherited           bool                             `json:"inherited,omitempty"`
	InheritedFromObject []string                         `json:"inherited_from_object,omitempty"`
	PermissionLevel     RegisteredModelPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegisteredModelPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegisteredModelPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelPermissionLevelPb string

const RegisteredModelPermissionLevelCanEdit RegisteredModelPermissionLevelPb = `CAN_EDIT`

const RegisteredModelPermissionLevelCanManage RegisteredModelPermissionLevelPb = `CAN_MANAGE`

const RegisteredModelPermissionLevelCanManageProductionVersions RegisteredModelPermissionLevelPb = `CAN_MANAGE_PRODUCTION_VERSIONS`

const RegisteredModelPermissionLevelCanManageStagingVersions RegisteredModelPermissionLevelPb = `CAN_MANAGE_STAGING_VERSIONS`

const RegisteredModelPermissionLevelCanRead RegisteredModelPermissionLevelPb = `CAN_READ`

type RegisteredModelPermissionsPb struct {
	AccessControlList []RegisteredModelAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                                   `json:"object_id,omitempty"`
	ObjectType        string                                   `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegisteredModelPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegisteredModelPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelPermissionsDescriptionPb struct {
	Description     string                           `json:"description,omitempty"`
	PermissionLevel RegisteredModelPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegisteredModelPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegisteredModelPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelPermissionsRequestPb struct {
	AccessControlList []RegisteredModelAccessControlRequestPb `json:"access_control_list,omitempty"`
	RegisteredModelId string                                  `json:"-" url:"-"`
}

type RegistryEmailSubscriptionTypePb string

// Subscribed to all events.
const RegistryEmailSubscriptionTypeAllEvents RegistryEmailSubscriptionTypePb = `ALL_EVENTS`

// Default subscription type.
const RegistryEmailSubscriptionTypeDefault RegistryEmailSubscriptionTypePb = `DEFAULT`

// Subscribed to notifications.
const RegistryEmailSubscriptionTypeSubscribed RegistryEmailSubscriptionTypePb = `SUBSCRIBED`

// Not subscribed to notifications.
const RegistryEmailSubscriptionTypeUnsubscribed RegistryEmailSubscriptionTypePb = `UNSUBSCRIBED`

type RegistryWebhookPb struct {
	CreationTimestamp    int64                       `json:"creation_timestamp,omitempty"`
	Description          string                      `json:"description,omitempty"`
	Events               []RegistryWebhookEventPb    `json:"events,omitempty"`
	HttpUrlSpec          *HttpUrlSpecWithoutSecretPb `json:"http_url_spec,omitempty"`
	Id                   string                      `json:"id,omitempty"`
	JobSpec              *JobSpecWithoutSecretPb     `json:"job_spec,omitempty"`
	LastUpdatedTimestamp int64                       `json:"last_updated_timestamp,omitempty"`
	ModelName            string                      `json:"model_name,omitempty"`
	Status               RegistryWebhookStatusPb     `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegistryWebhookEventPb string

const RegistryWebhookEventCommentCreated RegistryWebhookEventPb = `COMMENT_CREATED`

const RegistryWebhookEventModelVersionCreated RegistryWebhookEventPb = `MODEL_VERSION_CREATED`

const RegistryWebhookEventModelVersionTagSet RegistryWebhookEventPb = `MODEL_VERSION_TAG_SET`

const RegistryWebhookEventModelVersionTransitionedStage RegistryWebhookEventPb = `MODEL_VERSION_TRANSITIONED_STAGE`

const RegistryWebhookEventModelVersionTransitionedToArchived RegistryWebhookEventPb = `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`

const RegistryWebhookEventModelVersionTransitionedToProduction RegistryWebhookEventPb = `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`

const RegistryWebhookEventModelVersionTransitionedToStaging RegistryWebhookEventPb = `MODEL_VERSION_TRANSITIONED_TO_STAGING`

const RegistryWebhookEventRegisteredModelCreated RegistryWebhookEventPb = `REGISTERED_MODEL_CREATED`

const RegistryWebhookEventTransitionRequestCreated RegistryWebhookEventPb = `TRANSITION_REQUEST_CREATED`

const RegistryWebhookEventTransitionRequestToArchivedCreated RegistryWebhookEventPb = `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`

const RegistryWebhookEventTransitionRequestToProductionCreated RegistryWebhookEventPb = `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`

const RegistryWebhookEventTransitionRequestToStagingCreated RegistryWebhookEventPb = `TRANSITION_REQUEST_TO_STAGING_CREATED`

type RegistryWebhookStatusPb string

// Webhook is triggered when an associated event happens.
const RegistryWebhookStatusActive RegistryWebhookStatusPb = `ACTIVE`

// Webhook is not triggered.
const RegistryWebhookStatusDisabled RegistryWebhookStatusPb = `DISABLED`

// Webhook can be triggered through the test endpoint, but is not triggered on a
// real event.
const RegistryWebhookStatusTestMode RegistryWebhookStatusPb = `TEST_MODE`

type RejectTransitionRequestPb struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name"`
	Stage   string `json:"stage"`
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RejectTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RejectTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RejectTransitionRequestResponsePb struct {
	Activity *ActivityPb `json:"activity,omitempty"`
}

type RenameModelRequestPb struct {
	Name    string `json:"name"`
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RenameModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RenameModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RenameModelResponsePb struct {
	RegisteredModel *ModelPb `json:"registered_model,omitempty"`
}

type RestoreExperimentPb struct {
	ExperimentId string `json:"experiment_id"`
}

type RestoreExperimentResponsePb struct {
}

type RestoreRunPb struct {
	RunId string `json:"run_id"`
}

type RestoreRunResponsePb struct {
}

type RestoreRunsPb struct {
	ExperimentId       string `json:"experiment_id"`
	MaxRuns            int    `json:"max_runs,omitempty"`
	MinTimestampMillis int64  `json:"min_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RestoreRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RestoreRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RestoreRunsResponsePb struct {
	RunsRestored int `json:"runs_restored,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RestoreRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RestoreRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunPb struct {
	Data   *RunDataPb   `json:"data,omitempty"`
	Info   *RunInfoPb   `json:"info,omitempty"`
	Inputs *RunInputsPb `json:"inputs,omitempty"`
}

type RunDataPb struct {
	Metrics []MetricPb `json:"metrics,omitempty"`
	Params  []ParamPb  `json:"params,omitempty"`
	Tags    []RunTagPb `json:"tags,omitempty"`
}

type RunInfoPb struct {
	ArtifactUri    string          `json:"artifact_uri,omitempty"`
	EndTime        int64           `json:"end_time,omitempty"`
	ExperimentId   string          `json:"experiment_id,omitempty"`
	LifecycleStage string          `json:"lifecycle_stage,omitempty"`
	RunId          string          `json:"run_id,omitempty"`
	RunName        string          `json:"run_name,omitempty"`
	RunUuid        string          `json:"run_uuid,omitempty"`
	StartTime      int64           `json:"start_time,omitempty"`
	Status         RunInfoStatusPb `json:"status,omitempty"`
	UserId         string          `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunInfoStatusPb string

const RunInfoStatusFailed RunInfoStatusPb = `FAILED`

const RunInfoStatusFinished RunInfoStatusPb = `FINISHED`

const RunInfoStatusKilled RunInfoStatusPb = `KILLED`

const RunInfoStatusRunning RunInfoStatusPb = `RUNNING`

const RunInfoStatusScheduled RunInfoStatusPb = `SCHEDULED`

type RunInputsPb struct {
	DatasetInputs []DatasetInputPb `json:"dataset_inputs,omitempty"`
	ModelInputs   []ModelInputPb   `json:"model_inputs,omitempty"`
}

type RunTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchExperimentsPb struct {
	Filter     string     `json:"filter,omitempty"`
	MaxResults int64      `json:"max_results,omitempty"`
	OrderBy    []string   `json:"order_by,omitempty"`
	PageToken  string     `json:"page_token,omitempty"`
	ViewType   ViewTypePb `json:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchExperimentsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchExperimentsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchExperimentsResponsePb struct {
	Experiments   []ExperimentPb `json:"experiments,omitempty"`
	NextPageToken string         `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchExperimentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchExperimentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchLoggedModelsDatasetPb struct {
	DatasetDigest string `json:"dataset_digest,omitempty"`
	DatasetName   string `json:"dataset_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchLoggedModelsDatasetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchLoggedModelsDatasetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchLoggedModelsOrderByPb struct {
	Ascending     bool   `json:"ascending,omitempty"`
	DatasetDigest string `json:"dataset_digest,omitempty"`
	DatasetName   string `json:"dataset_name,omitempty"`
	FieldName     string `json:"field_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchLoggedModelsOrderByPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchLoggedModelsOrderByPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchLoggedModelsRequestPb struct {
	Datasets      []SearchLoggedModelsDatasetPb `json:"datasets,omitempty"`
	ExperimentIds []string                      `json:"experiment_ids,omitempty"`
	Filter        string                        `json:"filter,omitempty"`
	MaxResults    int                           `json:"max_results,omitempty"`
	OrderBy       []SearchLoggedModelsOrderByPb `json:"order_by,omitempty"`
	PageToken     string                        `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchLoggedModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchLoggedModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchLoggedModelsResponsePb struct {
	Models        []LoggedModelPb `json:"models,omitempty"`
	NextPageToken string          `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchLoggedModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchLoggedModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchModelVersionsRequestPb struct {
	Filter     string   `json:"-" url:"filter,omitempty"`
	MaxResults int64    `json:"-" url:"max_results,omitempty"`
	OrderBy    []string `json:"-" url:"order_by,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchModelVersionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchModelVersionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchModelVersionsResponsePb struct {
	ModelVersions []ModelVersionPb `json:"model_versions,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchModelVersionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchModelVersionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchModelsRequestPb struct {
	Filter     string   `json:"-" url:"filter,omitempty"`
	MaxResults int64    `json:"-" url:"max_results,omitempty"`
	OrderBy    []string `json:"-" url:"order_by,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchModelsResponsePb struct {
	NextPageToken    string    `json:"next_page_token,omitempty"`
	RegisteredModels []ModelPb `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchRunsPb struct {
	ExperimentIds []string   `json:"experiment_ids,omitempty"`
	Filter        string     `json:"filter,omitempty"`
	MaxResults    int        `json:"max_results,omitempty"`
	OrderBy       []string   `json:"order_by,omitempty"`
	PageToken     string     `json:"page_token,omitempty"`
	RunViewType   ViewTypePb `json:"run_view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchRunsResponsePb struct {
	NextPageToken string  `json:"next_page_token,omitempty"`
	Runs          []RunPb `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetExperimentTagPb struct {
	ExperimentId string `json:"experiment_id"`
	Key          string `json:"key"`
	Value        string `json:"value"`
}

type SetExperimentTagResponsePb struct {
}

type SetLoggedModelTagsRequestPb struct {
	ModelId string             `json:"-" url:"-"`
	Tags    []LoggedModelTagPb `json:"tags,omitempty"`
}

type SetLoggedModelTagsResponsePb struct {
}

type SetModelTagRequestPb struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SetModelTagResponsePb struct {
}

type SetModelVersionTagRequestPb struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Version string `json:"version"`
}

type SetModelVersionTagResponsePb struct {
}

type SetTagPb struct {
	Key     string `json:"key"`
	RunId   string `json:"run_id,omitempty"`
	RunUuid string `json:"run_uuid,omitempty"`
	Value   string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SetTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SetTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetTagResponsePb struct {
}

type StatusPb string

// Request to register a new model version has failed.
const StatusFailedRegistration StatusPb = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const StatusPendingRegistration StatusPb = `PENDING_REGISTRATION`

// Model version is ready for use.
const StatusReady StatusPb = `READY`

type TestRegistryWebhookRequestPb struct {
	Event RegistryWebhookEventPb `json:"event,omitempty"`
	Id    string                 `json:"id"`
}

type TestRegistryWebhookResponsePb struct {
	Body       string `json:"body,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TestRegistryWebhookResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TestRegistryWebhookResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TransitionModelVersionStageDatabricksPb struct {
	ArchiveExistingVersions bool   `json:"archive_existing_versions"`
	Comment                 string `json:"comment,omitempty"`
	Name                    string `json:"name"`
	Stage                   string `json:"stage"`
	Version                 string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TransitionModelVersionStageDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TransitionModelVersionStageDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TransitionRequestPb struct {
	AvailableActions  []ActivityActionPb `json:"available_actions,omitempty"`
	Comment           string             `json:"comment,omitempty"`
	CreationTimestamp int64              `json:"creation_timestamp,omitempty"`
	ToStage           string             `json:"to_stage,omitempty"`
	UserId            string             `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TransitionStageResponsePb struct {
	ModelVersionDatabricks *ModelVersionDatabricksPb `json:"model_version_databricks,omitempty"`
}

type UpdateCommentPb struct {
	Comment string `json:"comment"`
	Id      string `json:"id"`
}

type UpdateCommentResponsePb struct {
	Comment *CommentObjectPb `json:"comment,omitempty"`
}

type UpdateExperimentPb struct {
	ExperimentId string `json:"experiment_id"`
	NewName      string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateExperimentResponsePb struct {
}

type UpdateFeatureTagRequestPb struct {
	FeatureName string       `json:"-" url:"-"`
	FeatureTag  FeatureTagPb `json:"feature_tag"`
	Key         string       `json:"-" url:"-"`
	TableName   string       `json:"-" url:"-"`
	UpdateMask  string       `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateFeatureTagRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateFeatureTagRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateModelRequestPb struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateModelResponsePb struct {
	RegisteredModel *ModelPb `json:"registered_model,omitempty"`
}

type UpdateModelVersionRequestPb struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Version     string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateModelVersionResponsePb struct {
	ModelVersion *ModelVersionPb `json:"model_version,omitempty"`
}

type UpdateOnlineStoreRequestPb struct {
	Name        string        `json:"-" url:"-"`
	OnlineStore OnlineStorePb `json:"online_store"`
	UpdateMask  string        `json:"-" url:"update_mask"`
}

type UpdateRegistryWebhookPb struct {
	Description string                   `json:"description,omitempty"`
	Events      []RegistryWebhookEventPb `json:"events,omitempty"`
	HttpUrlSpec *HttpUrlSpecPb           `json:"http_url_spec,omitempty"`
	Id          string                   `json:"id"`
	JobSpec     *JobSpecPb               `json:"job_spec,omitempty"`
	Status      RegistryWebhookStatusPb  `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateRegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateRegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRunPb struct {
	EndTime int64             `json:"end_time,omitempty"`
	RunId   string            `json:"run_id,omitempty"`
	RunName string            `json:"run_name,omitempty"`
	RunUuid string            `json:"run_uuid,omitempty"`
	Status  UpdateRunStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRunResponsePb struct {
	RunInfo *RunInfoPb `json:"run_info,omitempty"`
}

type UpdateRunStatusPb string

const UpdateRunStatusFailed UpdateRunStatusPb = `FAILED`

const UpdateRunStatusFinished UpdateRunStatusPb = `FINISHED`

const UpdateRunStatusKilled UpdateRunStatusPb = `KILLED`

const UpdateRunStatusRunning UpdateRunStatusPb = `RUNNING`

const UpdateRunStatusScheduled UpdateRunStatusPb = `SCHEDULED`

type UpdateWebhookResponsePb struct {
	Webhook *RegistryWebhookPb `json:"webhook,omitempty"`
}

type ViewTypePb string

const ViewTypeActiveOnly ViewTypePb = `ACTIVE_ONLY`

const ViewTypeAll ViewTypePb = `ALL`

const ViewTypeDeletedOnly ViewTypePb = `DELETED_ONLY`
