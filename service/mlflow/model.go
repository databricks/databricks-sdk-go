// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflow

// all definitions in this file are in alphabetical order
// Activity recorded for the action.
type Activity struct {
	ActivityType ActivityType `json:"activity_type,omitempty"`

	Comment string `json:"comment,omitempty"`

	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Source stage of the transition (if the activity is stage transition
	// related). Valid values are: * `None`: The initial stage of a model
	// version. * `Staging`: Staging or pre-production stage. * `Production`:
	// Production stage. * `Archived`: Archived stage.
	FromStage Stage `json:"from_stage,omitempty"`

	Id string `json:"id,omitempty"`

	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	SystemComment string `json:"system_comment,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are: * `None`: The initial stage of a model
	// version. * `Staging`: Staging or pre-production stage. * `Production`:
	// Production stage. * `Archived`: Archived stage.
	ToStage Stage `json:"to_stage,omitempty"`
	// Username of the user who created the activity, e.g. the user who made the
	// stage transition request.
	UserId string `json:"user_id,omitempty"`
}

// This describes an enum
type ActivityAction string

// Approve a transition request
const ActivityActionApproveTransitionRequest ActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityAction = `CANCEL_TRANSITION_REQUEST`

// Reject a transition request
const ActivityActionRejectTransitionRequest ActivityAction = `REJECT_TRANSITION_REQUEST`

// This describes an enum
type ActivityType string

// User applied the corresponding stage transition.
const ActivityTypeAppliedTransition ActivityType = `APPLIED_TRANSITION`

// User approved the corresponding stage transition.
const ActivityTypeApprovedRequest ActivityType = `APPROVED_REQUEST`

// User cancelled an existing transition request.
const ActivityTypeCancelledRequest ActivityType = `CANCELLED_REQUEST`

const ActivityTypeNewComment ActivityType = `NEW_COMMENT`

// User rejected the coressponding stage transition.
const ActivityTypeRejectedRequest ActivityType = `REJECTED_REQUEST`

// User requested the corresponding stage transition.
const ActivityTypeRequestedTransition ActivityType = `REQUESTED_TRANSITION`

// For events performed as a side effect, such as archiving existing model
// versions in a stage.
const ActivityTypeSystemTransition ActivityType = `SYSTEM_TRANSITION`

type ApproveTransitionRequest struct {
	ArchiveExistingVersions bool `json:"archive_existing_versions"`

	Comment string `json:"comment,omitempty"`

	Name string `json:"name"`

	Stage Stage `json:"stage"`

	Version string `json:"version"`
}

// This describes an enum
type CommentActivityAction string

// Delete the comment
const CommentActivityActionDeleteComment CommentActivityAction = `DELETE_COMMENT`

// Edit the comment
const CommentActivityActionEditComment CommentActivityAction = `EDIT_COMMENT`

// Comment details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []CommentActivityAction `json:"available_actions,omitempty"`

	Comment string `json:"comment,omitempty"`

	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Username of the user who made the comment.
	UserId string `json:"user_id,omitempty"`
}

type CreateComment struct {
	Comment string `json:"comment"`

	Name string `json:"name"`

	Version string `json:"version"`
}

type CreateExperiment struct {
	// Location where all artifacts for the experiment are stored. If not
	// provided, the remote server will select an appropriate default.
	ArtifactLocation string `json:"artifact_location,omitempty"`
	// Experiment name.
	Name string `json:"name"`
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	Tags []ExperimentTag `json:"tags,omitempty"`
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
}

type CreateModelVersionRequest struct {
	// Optional description for model version.
	Description string `json:"description,omitempty"`
	// Register model under this name
	Name string `json:"name"`
	// MLflow run ID for correlation, if ``source`` was generated by an
	// experiment run in MLflow tracking server
	RunId string `json:"run_id,omitempty"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	RunLink string `json:"run_link,omitempty"`
	// URI indicating the location of the model artifacts.
	Source string `json:"source"`
	// Additional metadata for model version.
	Tags []ModelVersionTag `json:"tags,omitempty"`
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

type CreateRegisteredModelRequest struct {
	// Optional description for registered model.
	Description string `json:"description,omitempty"`
	// Register models under this name
	Name string `json:"name"`
	// Additional metadata for registered model.
	Tags []RegisteredModelTag `json:"tags,omitempty"`
}

type CreateRegisteredModelResponse struct {
	RegisteredModel *RegisteredModel `json:"registered_model,omitempty"`
}

type CreateRegistryWebhook struct {
	Description string `json:"description,omitempty"`

	Events []RegistryWebhookEvent `json:"events"`

	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`

	JobSpec *JobSpec `json:"job_spec,omitempty"`
	// If model name is not specified, a registry-wide webhook is created that
	// listens for the specified events across all versions of all registered
	// models.
	ModelName string `json:"model_name,omitempty"`

	Status RegistryWebhookStatus `json:"status,omitempty"`
}

type CreateRun struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime int64 `json:"start_time,omitempty"`
	// Additional metadata for run.
	Tags []RunTag `json:"tags,omitempty"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId string `json:"user_id,omitempty"`
}

type CreateRunResponse struct {
	// The newly created run.
	Run *Run `json:"run,omitempty"`
}

type CreateTransitionRequest struct {
	Comment string `json:"comment,omitempty"`

	Name string `json:"name"`

	Stage Stage `json:"stage"`

	Version string `json:"version"`
}

type DeleteComment struct {
	Id string `json:"id"`
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

type DeleteModelVersionRequest struct {
	// Name of the registered model
	Name string `json:"name"`
	// Model version number
	Version string `json:"version"`
}

type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"key"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"name"`
	// Model version number that the tag was logged under.
	Version string `json:"version"`
}

type DeleteRegisteredModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
}

type DeleteRegisteredModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"key"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"name"`
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId string `json:"run_id"`
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key string `json:"key"`
	// ID of the run that the tag was logged under. Must be provided.
	RunId string `json:"run_id"`
}

type Experiment struct {
	// Location where artifacts for the experiment are stored.
	ArtifactLocation string `json:"artifact_location,omitempty"`
	// Creation time
	CreationTime int64 `json:"creation_time,omitempty"`
	// Unique identifier for the experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Last update time
	LastUpdateTime int64 `json:"last_update_time,omitempty"`
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	LifecycleStage string `json:"lifecycle_stage,omitempty"`
	// Human readable name that identifies the experiment.
	Name string `json:"name,omitempty"`
	// Tags: Additional metadata key-value pairs.
	Tags []ExperimentTag `json:"tags,omitempty"`
}

type ExperimentTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`
}

type FileInfo struct {
	// Size in bytes. Unset for directories.
	FileSize int64 `json:"file_size,omitempty"`
	// Whether the path is a directory.
	IsDir bool `json:"is_dir,omitempty"`
	// Path relative to the root artifact directory run.
	Path string `json:"path,omitempty"`
}

type GetExperimentByNameResponse struct {
	// Experiment details.
	Experiment *Experiment `json:"experiment,omitempty"`
}

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
	// List of stages.
	Stages []string `json:"stages,omitempty"`
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current ``READY`` status. If no ``stages`` provided, returns the latest
	// version for each stage, including ``"None"``.
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric.
	Metrics []Metric `json:"metrics,omitempty"`
}

type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name string `json:"-" url:"name,omitempty"`
	// Model version number
	Version string `json:"-" url:"version,omitempty"`
}

type GetModelVersionDownloadUriResponse struct {
	// URI corresponding to where artifacts for this model version are stored.
	ArtifactUri string `json:"artifact_uri,omitempty"`
}

type GetModelVersionRequest struct {
	// Name of the registered model
	Name string `json:"-" url:"name,omitempty"`
	// Model version number
	Version string `json:"-" url:"version,omitempty"`
}

type GetModelVersionResponse struct {
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

type GetRegisteredModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"-" url:"name,omitempty"`
}

type GetRegisteredModelResponse struct {
	RegisteredModel *RegisteredModel `json:"registered_model,omitempty"`
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	Run *Run `json:"run,omitempty"`
}

type HttpUrlSpec struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	Authorization string `json:"authorization,omitempty"`
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	Secret string `json:"secret,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url string `json:"url"`
}

type JobSpec struct {
	// The personal access token used to authorize webhook's job runs.
	AccessToken string `json:"access_token"`
	// ID of the job that the webhook runs.
	JobId string `json:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job?s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl string `json:"workspace_url,omitempty"`
}

type ListArtifactsResponse struct {
	// File location and metadata for artifacts.
	Files []FileInfo `json:"files,omitempty"`
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken string `json:"next_page_token,omitempty"`
	// Root artifact directory for the run.
	RootUri string `json:"root_uri,omitempty"`
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`
}

type ListRegisteredModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
}

type ListRegisteredModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	NextPageToken string `json:"next_page_token,omitempty"`

	RegisteredModels []RegisteredModel `json:"registered_models,omitempty"`
}

type ListRegistryWebhooks struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events []RegistryWebhookEvent `json:"events,omitempty"`
	// If `model_name` is not specified, all webhooks associated with the
	// specified events are listed, regardless of their associated model.
	ModelName string `json:"model_name,omitempty"`
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	Metrics []Metric `json:"metrics,omitempty"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	Params []Param `json:"params,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id,omitempty"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	Tags []RunTag `json:"tags,omitempty"`
}

type LogMetric struct {
	// Name of the metric.
	Key string `json:"key"`
	// ID of the run under which to log the metric. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Step at which to log the metric
	Step int64 `json:"step,omitempty"`
	// Unix timestamp in milliseconds at the time metric was logged.
	Timestamp int64 `json:"timestamp"`
	// Double value of the metric being logged.
	Value float64 `json:"value"`
}

type LogModel struct {
	// MLmodel file in json format.
	ModelJson string `json:"model_json,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id,omitempty"`
}

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	Key string `json:"key"`
	// ID of the run under which to log the param. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the param being logged. Maximum size is 500 bytes.
	Value string `json:"value"`
}

type Metric struct {
	// Key identifying this metric.
	Key string `json:"key,omitempty"`
	// Step at which to log the metric.
	Step int64 `json:"step,omitempty"`
	// The timestamp at which this metric was recorded.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Value associated with this metric.
	Value float64 `json:"value,omitempty"`
}

type ModelVersion struct {
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	CurrentStage Stage `json:"current_stage,omitempty"`

	Description string `json:"description,omitempty"`

	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	Name string `json:"name,omitempty"`

	RunId string `json:"run_id,omitempty"`

	RunLink string `json:"run_link,omitempty"`

	Source string `json:"source,omitempty"`

	Status Status `json:"status,omitempty"`

	StatusMessage string `json:"status_message,omitempty"`

	Tags []ModelVersionTag `json:"tags,omitempty"`

	UserId string `json:"user_id,omitempty"`

	Version string `json:"version,omitempty"`
}

type ModelVersionDatabricks struct {
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	CurrentStage Stage `json:"current_stage,omitempty"`

	Description string `json:"description,omitempty"`

	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	Name string `json:"name,omitempty"`

	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	RunId string `json:"run_id,omitempty"`

	RunLink string `json:"run_link,omitempty"`

	Source string `json:"source,omitempty"`

	Status Status `json:"status,omitempty"`

	StatusMessage string `json:"status_message,omitempty"`

	Tags []ModelVersionTag `json:"tags,omitempty"`

	UserId string `json:"user_id,omitempty"`

	Version string `json:"version,omitempty"`
}

type ModelVersionTag struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type Param struct {
	// Key identifying this param.
	Key string `json:"key,omitempty"`
	// Value associated with this param.
	Value string `json:"value,omitempty"`
}

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
type PermissionLevel string

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanRead PermissionLevel = `CAN_READ`

type RegisteredModel struct {
	// Timestamp recorded when this ``registered_model`` was created.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Description of this ``registered_model``.
	Description string `json:"description,omitempty"`
	// Timestamp recorded when metadata for this ``registered_model`` was last
	// updated.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Collection of latest model versions for each stage. Only contains models
	// with current ``READY`` status.
	LatestVersions []ModelVersion `json:"latest_versions,omitempty"`
	// Unique name for the model.
	Name string `json:"name,omitempty"`
	// Tags: Additional metadata key-value pairs for this ``registered_model``.
	Tags []RegisteredModelTag `json:"tags,omitempty"`
	// User that created this ``registered_model``
	UserId string `json:"user_id,omitempty"`
}

type RegisteredModelDatabricks struct {
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Array of model versions, each the latest version for its stage.
	LatestVersions []ModelVersion `json:"latest_versions,omitempty"`

	Name string `json:"name,omitempty"`

	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Array of tags associated with the model.
	Tags []RegisteredModelTag `json:"tags,omitempty"`

	UserId string `json:"user_id,omitempty"`
}

type RegisteredModelTag struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type RegistryWebhookEvent string

const RegistryWebhookEventCommentCreated RegistryWebhookEvent = `COMMENT_CREATED`

const RegistryWebhookEventModelVersionCreated RegistryWebhookEvent = `MODEL_VERSION_CREATED`

const RegistryWebhookEventModelVersionTagSet RegistryWebhookEvent = `MODEL_VERSION_TAG_SET`

const RegistryWebhookEventModelVersionTransitionedStage RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_STAGE`

const RegistryWebhookEventModelVersionTransitionedToArchived RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`

const RegistryWebhookEventModelVersionTransitionedToProduction RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`

const RegistryWebhookEventModelVersionTransitionedToStaging RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_STAGING`

const RegistryWebhookEventRegisteredModelCreated RegistryWebhookEvent = `REGISTERED_MODEL_CREATED`

const RegistryWebhookEventTransitionRequestCreated RegistryWebhookEvent = `TRANSITION_REQUEST_CREATED`

const RegistryWebhookEventTransitionRequestToArchivedCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`

const RegistryWebhookEventTransitionRequestToProductionCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`

const RegistryWebhookEventTransitionRequestToStagingCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_STAGING_CREATED`

// This describes an enum
type RegistryWebhookStatus string

// Webhook is triggered when an associated event happens.
const RegistryWebhookStatusActive RegistryWebhookStatus = `ACTIVE`

// Webhook is not triggered.
const RegistryWebhookStatusDisabled RegistryWebhookStatus = `DISABLED`

// Webhook can be triggered through the test endpoint, but is not triggered on a
// real event.
const RegistryWebhookStatusTestMode RegistryWebhookStatus = `TEST_MODE`

type RejectTransitionRequest struct {
	Comment string `json:"comment,omitempty"`

	Name string `json:"name"`

	Stage Stage `json:"stage"`

	Version string `json:"version"`
}

type RenameRegisteredModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
	// If provided, updates the name for this ``registered_model``.
	NewName string `json:"new_name,omitempty"`
}

type RenameRegisteredModelResponse struct {
	RegisteredModel *RegisteredModel `json:"registered_model,omitempty"`
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId string `json:"run_id"`
}

type Run struct {
	// Run data.
	Data *RunData `json:"data,omitempty"`
	// Run metadata.
	Info *RunInfo `json:"info,omitempty"`
}

type RunData struct {
	// Run metrics.
	Metrics []Metric `json:"metrics,omitempty"`
	// Run parameters.
	Params []Param `json:"params,omitempty"`
	// Additional metadata key-value pairs.
	Tags []RunTag `json:"tags,omitempty"`
}

type RunInfo struct {
	// URI of the directory where artifacts should be uploaded. This can be a
	// local path (starting with "/"), or a distributed file system (DFS) path,
	// like ``s3://bucket/directory`` or ``dbfs:/my/directory``. If not set, the
	// local ``./mlruns`` directory is chosen.
	ArtifactUri string `json:"artifact_uri,omitempty"`
	// Unix timestamp of when the run ended in milliseconds.
	EndTime int64 `json:"end_time,omitempty"`
	// The experiment ID.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	LifecycleStage string `json:"lifecycle_stage,omitempty"`
	// Unique identifier for the run.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Unix timestamp of when the run started in milliseconds.
	StartTime int64 `json:"start_time,omitempty"`
	// Current status of the run.
	Status RunInfoStatus `json:"status,omitempty"`
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	UserId string `json:"user_id,omitempty"`
}

// Current status of the run.
type RunInfoStatus string

const RunInfoStatusFailed RunInfoStatus = `FAILED`

const RunInfoStatusFinished RunInfoStatus = `FINISHED`

const RunInfoStatusKilled RunInfoStatus = `KILLED`

const RunInfoStatusRunning RunInfoStatus = `RUNNING`

const RunInfoStatusScheduled RunInfoStatus = `SCHEDULED`

type RunTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`
}

type SearchExperiments struct {
	// String representing a SQL filter condition (e.g. "name ILIKE
	// 'my-experiment%'")
	Filter string `json:"filter,omitempty"`
	// Maximum number of experiments desired. Max threshold is 3000.
	MaxResults int64 `json:"max_results,omitempty"`
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	OrderBy []string `json:"order_by,omitempty"`
	// Token indicating the page of experiments to fetch
	PageToken string `json:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType SearchExperimentsViewType `json:"view_type,omitempty"`
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`
}

// Qualifier for type of experiments to be returned. If unspecified, return only
// active experiments.
type SearchExperimentsViewType string

const SearchExperimentsViewTypeActiveOnly SearchExperimentsViewType = `ACTIVE_ONLY`

const SearchExperimentsViewTypeAll SearchExperimentsViewType = `ALL`

const SearchExperimentsViewTypeDeletedOnly SearchExperimentsViewType = `DELETED_ONLY`

type SearchModelVersionsRequest struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Max threshold is 10K.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	OrderBy any/* MISSING TYPE */ `json:"-" url:"order_by,omitempty"`
	// Pagination token to go to next page based on previous search query.
	PageToken string `json:"-" url:"page_token,omitempty"`
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken string `json:"next_page_token,omitempty"`
}

type SearchRegisteredModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	OrderBy any/* MISSING TYPE */ `json:"-" url:"order_by,omitempty"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken string `json:"-" url:"page_token,omitempty"`
}

type SearchRegisteredModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Registered Models that match the search criteria.
	RegisteredModels []RegisteredModel `json:"registered_models,omitempty"`
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	// A filter expression over params, metrics, and tags, that allows returning
	// a subset of runs. The syntax is a subset of SQL that supports ANDing
	// together binary operations between a param, metric, or tag and a
	// constant. Example: ``metrics.rmse < 1 and params.model_class =
	// 'LogisticRegression'`` You can select columns with special characters
	// (hyphen, space, period, etc.) by using double quotes: ``metrics."model
	// class" = 'LinearRegression' and tags."user-name" = 'Tomas'`` Supported
	// operators are ``=``, ``!=``, ``>``, ``>=``, ``<``, and ``<=``.
	Filter string `json:"filter,omitempty"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults int `json:"max_results,omitempty"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional "DESC" or "ASC" annotation, where "ASC" is the
	// default. Example: ["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"] Tiebreaks are done by start_time DESC followed by run_id
	// for runs with the same start time (and this is the default ordering
	// criterion if order_by is not provided).
	OrderBy []string `json:"order_by,omitempty"`

	PageToken string `json:"page_token,omitempty"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType SearchRunsRunViewType `json:"run_view_type,omitempty"`
}

type SearchRunsResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`
	// Runs that match the search criteria.
	Runs []Run `json:"runs,omitempty"`
}

// Whether to display only active, only deleted, or all runs. Defaults to only
// active runs.
type SearchRunsRunViewType string

const SearchRunsRunViewTypeActiveOnly SearchRunsRunViewType = `ACTIVE_ONLY`

const SearchRunsRunViewTypeAll SearchRunsRunViewType = `ALL`

const SearchRunsRunViewTypeDeletedOnly SearchRunsRunViewType = `DELETED_ONLY`

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	ExperimentId string `json:"experiment_id"`
	// Name of the tag. Maximum size depends on storage backend. All storage
	// backends are guaranteed to support key values up to 250 bytes in size.
	Key string `json:"key"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value string `json:"value"`
}

type SetModelVersionTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key string `json:"key"`
	// Unique name of the model.
	Name string `json:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value string `json:"value"`
	// Model version number.
	Version string `json:"version"`
}

type SetRegisteredModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key string `json:"key"`
	// Unique name of the model.
	Name string `json:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value string `json:"value"`
}

type SetTag struct {
	// Name of the tag. Maximum size depends on storage backend. All storage
	// backends are guaranteed to support key values up to 250 bytes in size.
	Key string `json:"key"`
	// ID of the run under which to log the tag. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value string `json:"value"`
}

// Test webhook response object.
type TestRegistryWebhook struct {
	// Body of the response from the webhook URL
	Body string `json:"body,omitempty"`
	// Status code returned by the webhook URL
	StatusCode int `json:"status_code,omitempty"`
}

type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event RegistryWebhookEvent `json:"event,omitempty"`

	Id string `json:"id"`
}

type TransitionModelVersionStage struct {
	// When transitioning a model version to a particular stage, this flag
	// dictates whether all existing model versions in that stage should be
	// atomically moved to the "archived" stage. This ensures that at-most-one
	// model version exists in the target stage. This field is *required* when
	// transitioning a model versions's stage
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// Name of the registered model
	Name string `json:"name"`
	// Transition `model_version` to new stage.
	Stage string `json:"stage"`
	// Model version number
	Version string `json:"version"`
}

type TransitionModelVersionStageDatabricks struct {
	ArchiveExistingVersions bool `json:"archive_existing_versions"`

	Comment string `json:"comment,omitempty"`

	Name string `json:"name"`

	Stage Stage `json:"stage"`

	Version string `json:"version"`
}

type TransitionModelVersionStageResponse struct {
	// Updated model version
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

// Transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []ActivityAction `json:"available_actions,omitempty"`

	Comment string `json:"comment,omitempty"`

	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are: * `None`: The initial stage of a model
	// version. * `Staging`: Staging or pre-production stage. * `Production`:
	// Production stage. * `Archived`: Archived stage.
	ToStage Stage `json:"to_stage,omitempty"`
	// Username of the user who made the stage transition request.
	UserId string `json:"user_id,omitempty"`
}

type UpdateComment struct {
	Comment string `json:"comment"`

	Id string `json:"id"`
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName string `json:"new_name,omitempty"`
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this ``registered_model``.
	Description string `json:"description,omitempty"`
	// Name of the registered model
	Name string `json:"name"`
	// Model version number
	Version string `json:"version"`
}

type UpdateRegisteredModelRequest struct {
	// If provided, updates the description for this ``registered_model``.
	Description string `json:"description,omitempty"`
	// Registered model unique name identifier.
	Name string `json:"name"`
}

type UpdateRegistryWebhook struct {
	Description string `json:"description,omitempty"`

	Events []RegistryWebhookEvent `json:"events,omitempty"`

	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`

	Id string `json:"id"`

	JobSpec *JobSpec `json:"job_spec,omitempty"`

	Status RegistryWebhookStatus `json:"status,omitempty"`
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime int64 `json:"end_time,omitempty"`
	// ID of the run to update. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run to update.. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Updated status of the run.
	Status UpdateRunStatus `json:"status,omitempty"`
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo *RunInfo `json:"run_info,omitempty"`
}

// Updated status of the run.
type UpdateRunStatus string

const UpdateRunStatusFailed UpdateRunStatus = `FAILED`

const UpdateRunStatusFinished UpdateRunStatus = `FINISHED`

const UpdateRunStatusKilled UpdateRunStatus = `KILLED`

const UpdateRunStatusRunning UpdateRunStatus = `RUNNING`

const UpdateRunStatusScheduled UpdateRunStatus = `SCHEDULED`

// User-provided comment associated with the activity.

// Unique identifier of an activity

type ApproveResponse struct {
	Activity *Activity `json:"activity,omitempty"`
}

// Specifies whether to archive all current model versions in the target stage.

// User-provided comment on the action.

type CreateResponse struct {
	Comment *CommentObject `json:"comment,omitempty"`
}

// Creation time of the object, as a Unix timestamp in milliseconds.

type DeleteRequest struct {
	Id string `json:"id"`
}

// User-specified description for the object.

type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName string `json:"-" url:"experiment_name,omitempty"`
}

type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId string `json:"-" url:"experiment_id,omitempty"`
}

type GetHistoryRequest struct {
	// Name of the metric.
	MetricKey string `json:"-" url:"metric_key,omitempty"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run from which to fetch metric
	// values. This field will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`
}

type GetRequest struct {
	// Name of the model.
	Name string `json:"-" url:"name,omitempty"`
}

type GetResponse struct {
	RegisteredModel *RegisteredModelDatabricks `json:"registered_model,omitempty"`
}

type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run to fetch. This field will
	// be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`
}

// Unique identifier for the object.

// Key for the tag.

// Time of the object at last update, as a Unix timestamp in milliseconds.

type ListArtifactsRequest struct {
	// Token indicating the page of artifact results to fetch
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path string `json:"-" url:"path,omitempty"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`
}

type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Token indicating the page of experiments to fetch
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType string `json:"-" url:"view_type,omitempty"`
}

type ListRequest struct {
	// Name of the model.
	Name string `json:"-" url:"name,omitempty"`
	// Version of the model.
	Version string `json:"-" url:"version,omitempty"`
}

type ListResponse struct {
	// Array of registry webhooks.
	Webhooks any /* MISSING TYPE */ `json:"webhooks,omitempty"`
}

// Name of the model whose events would trigger this webhook.

// Name of the model.

type RejectResponse struct {
	Activity *Activity `json:"activity,omitempty"`
}

// User-provided comment associated with the transition request.

// Unique identifier for the MLflow tracking run associated with the source
// model artifacts.

// URL of the run associated with the model artifacts, potentially in another
// workspace.

// URI that indicates the location of the source model artifacts. This is used
// when creating the model version.

// This describes an enum
type Stage string

// Archived stage.
const StageArchived Stage = `Archived`

// The initial stage of a model version.
const StageNone Stage = `None`

// Production stage.
const StageProduction Stage = `Production`

// Staging or pre-production stage.
const StageStaging Stage = `Staging`

// This describes an enum
type Status string

// Request to register a new model version has failed.
const StatusFailedRegistration Status = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const StatusPendingRegistration Status = `PENDING_REGISTRATION`

// Model version is ready for use.
const StatusReady Status = `READY`

// Details on the current status, for example why registration failed.

type TestRegistryWebhookResponse struct {
	Webhook *TestRegistryWebhook `json:"webhook,omitempty"`
}

type TransitionStageResponse struct {
	ModelVersion *ModelVersionDatabricks `json:"model_version,omitempty"`
}

type UpdateResponse struct {
	Comment *CommentObject `json:"comment,omitempty"`
}

// The username of the user that created the object.

// Value for the tag.

// Version of the model.

// User-specified description for the webhook.

// Webhook ID
