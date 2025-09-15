// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type Activity struct {
	ActivityType ActivityType `json:"activity_type,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Source stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	FromStage string `json:"from_stage,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	SystemComment string `json:"system_comment,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage string `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Activity) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Activity) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An action that a user (with sufficient permissions) could take on an activity
// or comment.
//
// For activities, valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
//
// For comments, valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type ActivityAction string

// Approve a transition request
const ActivityActionApproveTransitionRequest ActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityAction = `CANCEL_TRANSITION_REQUEST`

// Delete the comment
const ActivityActionDeleteComment ActivityAction = `DELETE_COMMENT`

// Edit the comment
const ActivityActionEditComment ActivityAction = `EDIT_COMMENT`

// Reject a transition request
const ActivityActionRejectTransitionRequest ActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *ActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `DELETE_COMMENT`, `EDIT_COMMENT`, `REJECT_TRANSITION_REQUEST`:
		*f = ActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "DELETE_COMMENT", "EDIT_COMMENT", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Values returns all possible values for ActivityAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *ActivityAction) Values() []ActivityAction {
	return []ActivityAction{
		ActivityActionApproveTransitionRequest,
		ActivityActionCancelTransitionRequest,
		ActivityActionDeleteComment,
		ActivityActionEditComment,
		ActivityActionRejectTransitionRequest,
	}
}

// Type always returns ActivityAction to satisfy [pflag.Value] interface
func (f *ActivityAction) Type() string {
	return "ActivityAction"
}

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition.
//
// * `REQUESTED_TRANSITION`: User requested the corresponding stage transition.
//
// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
//
// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
//
// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
//
// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
// archiving existing model versions in a stage.
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

// String representation for [fmt.Print]
func (f *ActivityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityType) Set(v string) error {
	switch v {
	case `APPLIED_TRANSITION`, `APPROVED_REQUEST`, `CANCELLED_REQUEST`, `NEW_COMMENT`, `REJECTED_REQUEST`, `REQUESTED_TRANSITION`, `SYSTEM_TRANSITION`:
		*f = ActivityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPLIED_TRANSITION", "APPROVED_REQUEST", "CANCELLED_REQUEST", "NEW_COMMENT", "REJECTED_REQUEST", "REQUESTED_TRANSITION", "SYSTEM_TRANSITION"`, v)
	}
}

// Values returns all possible values for ActivityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ActivityType) Values() []ActivityType {
	return []ActivityType{
		ActivityTypeAppliedTransition,
		ActivityTypeApprovedRequest,
		ActivityTypeCancelledRequest,
		ActivityTypeNewComment,
		ActivityTypeRejectedRequest,
		ActivityTypeRequestedTransition,
		ActivityTypeSystemTransition,
	}
}

// Type always returns ActivityType to satisfy [pflag.Value] interface
func (f *ActivityType) Type() string {
	return "ActivityType"
}

// Details required to identify and approve a model version stage transition
// request.
type ApproveTransitionRequest struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ApproveTransitionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApproveTransitionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ApproveTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity *Activity `json:"activity,omitempty"`
}

// An action that a user (with sufficient permissions) could take on an activity
// or comment.
//
// For activities, valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
//
// For comments, valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type CommentActivityAction string

// Approve a transition request
const CommentActivityActionApproveTransitionRequest CommentActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const CommentActivityActionCancelTransitionRequest CommentActivityAction = `CANCEL_TRANSITION_REQUEST`

// Delete the comment
const CommentActivityActionDeleteComment CommentActivityAction = `DELETE_COMMENT`

// Edit the comment
const CommentActivityActionEditComment CommentActivityAction = `EDIT_COMMENT`

// Reject a transition request
const CommentActivityActionRejectTransitionRequest CommentActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *CommentActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommentActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `DELETE_COMMENT`, `EDIT_COMMENT`, `REJECT_TRANSITION_REQUEST`:
		*f = CommentActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "DELETE_COMMENT", "EDIT_COMMENT", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Values returns all possible values for CommentActivityAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *CommentActivityAction) Values() []CommentActivityAction {
	return []CommentActivityAction{
		CommentActivityActionApproveTransitionRequest,
		CommentActivityActionCancelTransitionRequest,
		CommentActivityActionDeleteComment,
		CommentActivityActionEditComment,
		CommentActivityActionRejectTransitionRequest,
	}
}

// Type always returns CommentActivityAction to satisfy [pflag.Value] interface
func (f *CommentActivityAction) Type() string {
	return "CommentActivityAction"
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []CommentActivityAction `json:"available_actions,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CommentObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CommentObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to create a comment on a model version.
type CreateComment struct {
	// User-provided comment on the action.
	Comment string `json:"comment"`
	// Name of the model.
	Name string `json:"name"`
	// Version of the model.
	Version string `json:"version"`
}

type CreateCommentResponse struct {
	// New comment object
	Comment *CommentObject `json:"comment,omitempty"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExperiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExperiment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExperimentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExperimentResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFeatureRequest struct {
	// Feature to create.
	Feature Feature `json:"feature"`
}

type CreateFeatureTagRequest struct {
	FeatureName string `json:"-" url:"-"`

	FeatureTag FeatureTag `json:"feature_tag"`

	TableName string `json:"-" url:"-"`
}

type CreateForecastingExperimentRequest struct {
	// The column in the training table used to customize weights for each time
	// series.
	CustomWeightsColumn string `json:"custom_weights_column,omitempty"`
	// The path in the workspace to store the created experiment.
	ExperimentPath string `json:"experiment_path,omitempty"`
	// The time interval between consecutive rows in the time series data.
	// Possible values include: '1 second', '1 minute', '5 minutes', '10
	// minutes', '15 minutes', '30 minutes', 'Hourly', 'Daily', 'Weekly',
	// 'Monthly', 'Quarterly', 'Yearly'.
	ForecastGranularity string `json:"forecast_granularity"`
	// The number of time steps into the future to make predictions, calculated
	// as a multiple of forecast_granularity. This value represents how far
	// ahead the model should forecast.
	ForecastHorizon int64 `json:"forecast_horizon"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store future feature data
	// for predictions.
	FutureFeatureDataPath string `json:"future_feature_data_path,omitempty"`
	// The region code(s) to automatically add holiday features. Currently
	// supports only one region.
	HolidayRegions []string `json:"holiday_regions,omitempty"`
	// Specifies the list of feature columns to include in model training. These
	// columns must exist in the training data and be of type string, numerical,
	// or boolean. If not specified, no additional features will be included.
	// Note: Certain columns are automatically handled: - Automatically
	// excluded: split_column, target_column, custom_weights_column. -
	// Automatically included: time_column.
	IncludeFeatures []string `json:"include_features,omitempty"`
	// The maximum duration for the experiment in minutes. The experiment stops
	// automatically if it exceeds this limit.
	MaxRuntime int64 `json:"max_runtime,omitempty"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store predictions.
	PredictionDataPath string `json:"prediction_data_path,omitempty"`
	// The evaluation metric used to optimize the forecasting model.
	PrimaryMetric string `json:"primary_metric,omitempty"`
	// The fully qualified path of a Unity Catalog model, formatted as
	// catalog_name.schema_name.model_name, used to store the best model.
	RegisterTo string `json:"register_to,omitempty"`
	// // The column in the training table used for custom data splits. Values
	// must be 'train', 'validate', or 'test'.
	SplitColumn string `json:"split_column,omitempty"`
	// The column in the input training table used as the prediction target for
	// model training. The values in this column are used as the ground truth
	// for model training.
	TargetColumn string `json:"target_column"`
	// The column in the input training table that represents each row's
	// timestamp.
	TimeColumn string `json:"time_column"`
	// The column in the training table used to group the dataset for predicting
	// individual time series.
	TimeseriesIdentifierColumns []string `json:"timeseries_identifier_columns,omitempty"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used as training data for the
	// forecasting model.
	TrainDataPath string `json:"train_data_path"`
	// List of frameworks to include for model tuning. Possible values are
	// 'Prophet', 'ARIMA', 'DeepAR'. An empty list includes all supported
	// frameworks.
	TrainingFrameworks []string `json:"training_frameworks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateForecastingExperimentRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateForecastingExperimentResponse struct {
	// The unique ID of the created forecasting experiment
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateForecastingExperimentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateForecastingExperimentResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateLoggedModelRequest struct {
	// The ID of the experiment that owns the model.
	ExperimentId string `json:"experiment_id"`
	// The type of the model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	ModelType string `json:"model_type,omitempty"`
	// The name of the model (optional). If not specified one will be generated.
	Name string `json:"name,omitempty"`
	// Parameters attached to the model.
	Params []LoggedModelParameter `json:"params,omitempty"`
	// The ID of the run that created the model.
	SourceRunId string `json:"source_run_id,omitempty"`
	// Tags attached to the model.
	Tags []LoggedModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateLoggedModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateLoggedModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateLoggedModelResponse struct {
	// The newly created logged model.
	Model *LoggedModel `json:"model,omitempty"`
}

type CreateModelRequest struct {
	// Optional description for registered model.
	Description string `json:"description,omitempty"`
	// Register models under this name
	Name string `json:"name"`
	// Additional metadata for registered model.
	Tags []ModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateModelResponse struct {
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

type CreateModelVersionRequest struct {
	// Optional description for model version.
	Description string `json:"description,omitempty"`
	// Register model under this name
	Name string `json:"name"`
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	RunId string `json:"run_id,omitempty"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	RunLink string `json:"run_link,omitempty"`
	// URI indicating the location of the model artifacts.
	Source string `json:"source"`
	// Additional metadata for model version.
	Tags []ModelVersionTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

type CreateOnlineStoreRequest struct {
	// Online store to create.
	OnlineStore OnlineStore `json:"online_store"`
}

// Details required to create a registry webhook.
type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events"`
	// External HTTPS URL called on event trigger (by using a POST request).
	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`
	// ID of the job that the webhook runs.
	JobSpec *JobSpec `json:"job_spec,omitempty"`
	// If model name is not specified, a registry-wide webhook is created that
	// listens for the specified events across all versions of all registered
	// models.
	ModelName string `json:"model_name,omitempty"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRegistryWebhook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRegistryWebhook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRun struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The name of the run.
	RunName string `json:"run_name,omitempty"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime int64 `json:"start_time,omitempty"`
	// Additional metadata for run.
	Tags []RunTag `json:"tags,omitempty"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRunResponse struct {
	// The newly created run.
	Run *Run `json:"run,omitempty"`
}

// Details required to create a model version stage transition request.
type CreateTransitionRequest struct {
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateTransitionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTransitionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTransitionRequestResponse struct {
	// New activity generated for stage transition request.
	Request *TransitionRequest `json:"request,omitempty"`
}

type CreateWebhookResponse struct {
	Webhook *RegistryWebhook `json:"webhook,omitempty"`
}

type DataSource struct {
	DeltaTableSource *DeltaTableSource `json:"delta_table_source,omitempty"`
}

// Dataset. Represents a reference to data used for training, testing, or
// evaluation during the model development process.
type Dataset struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	Digest string `json:"digest"`
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	Name string `json:"name"`
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	Profile string `json:"profile,omitempty"`
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	Schema string `json:"schema,omitempty"`
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	Source string `json:"source"`
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	SourceType string `json:"source_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Dataset) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dataset) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// DatasetInput. Represents a dataset and input tags.
type DatasetInput struct {
	// The dataset being used as a Run input.
	Dataset Dataset `json:"dataset"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags []InputTag `json:"tags,omitempty"`
}

type DeleteCommentRequest struct {
	// Unique identifier of an activity
	Id string `json:"-" url:"id"`
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

type DeleteFeatureRequest struct {
	// Name of the feature to delete.
	FullName string `json:"-" url:"-"`
}

type DeleteFeatureTagRequest struct {
	// The name of the feature within the feature table.
	FeatureName string `json:"-" url:"-"`
	// The key of the tag to delete.
	Key string `json:"-" url:"-"`
	// The name of the feature table.
	TableName string `json:"-" url:"-"`
}

type DeleteLoggedModelRequest struct {
	// The ID of the logged model to delete.
	ModelId string `json:"-" url:"-"`
}

type DeleteLoggedModelTagRequest struct {
	// The ID of the logged model to delete the tag from.
	ModelId string `json:"-" url:"-"`
	// The tag key.
	TagKey string `json:"-" url:"-"`
}

type DeleteModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"-" url:"name"`
}

type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"-" url:"key"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"-" url:"name"`
}

type DeleteModelVersionRequest struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"-" url:"key"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"-" url:"name"`
	// Model version number that the tag was logged under.
	Version string `json:"-" url:"version"`
}

type DeleteOnlineStoreRequest struct {
	// Name of the online store to delete.
	Name string `json:"-" url:"-"`
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId string `json:"run_id"`
}

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	ExperimentId string `json:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	MaxRuns int `json:"max_runs,omitempty"`
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	MaxTimestampMillis int64 `json:"max_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRuns) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRuns) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteRunsResponse struct {
	// The number of runs deleted.
	RunsDeleted int `json:"runs_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key string `json:"key"`
	// ID of the run that the tag was logged under. Must be provided.
	RunId string `json:"run_id"`
}

type DeleteTransitionRequestRequest struct {
	// User-provided comment on the action.
	Comment string `json:"-" url:"comment,omitempty"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	Creator string `json:"-" url:"creator"`
	// Name of the model.
	Name string `json:"-" url:"name"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"-" url:"stage"`
	// Version of the model.
	Version string `json:"-" url:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteTransitionRequestRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteTransitionRequestRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity *Activity `json:"activity,omitempty"`
}

type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	Id string `json:"-" url:"id"`
}

type DeltaTableSource struct {
	// The entity columns of the Delta table.
	EntityColumns []string `json:"entity_columns"`
	// The full three-part (catalog, schema, table) name of the Delta table.
	FullName string `json:"full_name"`
	// The timeseries column of the Delta table.
	TimeseriesColumn string `json:"timeseries_column"`
}

// An experiment and its metadata.
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Experiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Experiment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentAccessControlResponse struct {
	// All permissions.
	AllPermissions []ExperimentPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type ExperimentPermissionLevel string

const ExperimentPermissionLevelCanEdit ExperimentPermissionLevel = `CAN_EDIT`

const ExperimentPermissionLevelCanManage ExperimentPermissionLevel = `CAN_MANAGE`

const ExperimentPermissionLevelCanRead ExperimentPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *ExperimentPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExperimentPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`:
		*f = ExperimentPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ"`, v)
	}
}

// Values returns all possible values for ExperimentPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExperimentPermissionLevel) Values() []ExperimentPermissionLevel {
	return []ExperimentPermissionLevel{
		ExperimentPermissionLevelCanEdit,
		ExperimentPermissionLevelCanManage,
		ExperimentPermissionLevelCanRead,
	}
}

// Type always returns ExperimentPermissionLevel to satisfy [pflag.Value] interface
func (f *ExperimentPermissionLevel) Type() string {
	return "ExperimentPermissionLevel"
}

type ExperimentPermissions struct {
	AccessControlList []ExperimentAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentPermissionsRequest struct {
	AccessControlList []ExperimentAccessControlRequest `json:"access_control_list,omitempty"`
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

// A tag for an experiment.
type ExperimentTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Feature struct {
	// The description of the feature.
	Description string `json:"description,omitempty"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName string `json:"full_name"`
	// The function by which the feature is computed.
	Function Function `json:"function"`
	// The input columns from which the feature is computed.
	Inputs []string `json:"inputs"`
	// The data source of the feature.
	Source DataSource `json:"source"`
	// The time window in which the feature is computed.
	TimeWindow TimeWindow `json:"time_window"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Feature) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Feature) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FeatureLineage struct {
	// List of feature specs that contain this feature.
	FeatureSpecs []FeatureLineageFeatureSpec `json:"feature_specs,omitempty"`
	// List of Unity Catalog models that were trained on this feature.
	Models []FeatureLineageModel `json:"models,omitempty"`
	// List of online features that use this feature as source.
	OnlineFeatures []FeatureLineageOnlineFeature `json:"online_features,omitempty"`
}

type FeatureLineageFeatureSpec struct {
	// The full name of the feature spec in Unity Catalog.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureLineageFeatureSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureLineageFeatureSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FeatureLineageModel struct {
	// The full name of the model in Unity Catalog.
	Name string `json:"name,omitempty"`
	// The version of the model.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureLineageModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureLineageModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FeatureLineageOnlineFeature struct {
	// The name of the online feature (column name).
	FeatureName string `json:"feature_name,omitempty"`
	// The full name of the online table in Unity Catalog.
	TableName string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureLineageOnlineFeature) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureLineageOnlineFeature) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Feature list wrap all the features for a model version
type FeatureList struct {
	Features []LinkedFeature `json:"features,omitempty"`
}

// Represents a tag on a feature in a feature table.
type FeatureTag struct {
	Key string `json:"key"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata of a single artifact file or directory.
type FileInfo struct {
	// The size in bytes of the file. Unset for directories.
	FileSize int64 `json:"file_size,omitempty"`
	// Whether the path is a directory.
	IsDir bool `json:"is_dir,omitempty"`
	// The path relative to the root artifact directory run.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FinalizeLoggedModelRequest struct {
	// The ID of the logged model to finalize.
	ModelId string `json:"-" url:"-"`
	// Whether or not the model is ready for use.
	// ``"LOGGED_MODEL_UPLOAD_FAILED"`` indicates that something went wrong when
	// logging the model weights / agent code.
	Status LoggedModelStatus `json:"status"`
}

type FinalizeLoggedModelResponse struct {
	// The updated logged model.
	Model *LoggedModel `json:"model,omitempty"`
}

// Represents a forecasting experiment with its unique identifier, URL, and
// state.
type ForecastingExperiment struct {
	// The unique ID for the forecasting experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The URL to the forecasting experiment page.
	ExperimentPageUrl string `json:"experiment_page_url,omitempty"`
	// The current state of the forecasting experiment.
	State ForecastingExperimentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ForecastingExperiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForecastingExperiment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForecastingExperimentState string

const ForecastingExperimentStateCancelled ForecastingExperimentState = `CANCELLED`

const ForecastingExperimentStateFailed ForecastingExperimentState = `FAILED`

const ForecastingExperimentStatePending ForecastingExperimentState = `PENDING`

const ForecastingExperimentStateRunning ForecastingExperimentState = `RUNNING`

const ForecastingExperimentStateSucceeded ForecastingExperimentState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *ForecastingExperimentState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ForecastingExperimentState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = ForecastingExperimentState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for ForecastingExperimentState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ForecastingExperimentState) Values() []ForecastingExperimentState {
	return []ForecastingExperimentState{
		ForecastingExperimentStateCancelled,
		ForecastingExperimentStateFailed,
		ForecastingExperimentStatePending,
		ForecastingExperimentStateRunning,
		ForecastingExperimentStateSucceeded,
	}
}

// Type always returns ForecastingExperimentState to satisfy [pflag.Value] interface
func (f *ForecastingExperimentState) Type() string {
	return "ForecastingExperimentState"
}

type Function struct {
	// Extra parameters for parameterized functions.
	ExtraParameters []FunctionExtraParameter `json:"extra_parameters,omitempty"`
	// The type of the function.
	FunctionType FunctionFunctionType `json:"function_type"`
}

type FunctionExtraParameter struct {
	// The name of the parameter.
	Key string `json:"key"`
	// The value of the parameter.
	Value string `json:"value"`
}

type FunctionFunctionType string

const FunctionFunctionTypeApproxCountDistinct FunctionFunctionType = `APPROX_COUNT_DISTINCT`

const FunctionFunctionTypeApproxPercentile FunctionFunctionType = `APPROX_PERCENTILE`

const FunctionFunctionTypeAvg FunctionFunctionType = `AVG`

const FunctionFunctionTypeCount FunctionFunctionType = `COUNT`

const FunctionFunctionTypeFirst FunctionFunctionType = `FIRST`

const FunctionFunctionTypeLast FunctionFunctionType = `LAST`

const FunctionFunctionTypeMax FunctionFunctionType = `MAX`

const FunctionFunctionTypeMin FunctionFunctionType = `MIN`

const FunctionFunctionTypeStddevPop FunctionFunctionType = `STDDEV_POP`

const FunctionFunctionTypeStddevSamp FunctionFunctionType = `STDDEV_SAMP`

const FunctionFunctionTypeSum FunctionFunctionType = `SUM`

const FunctionFunctionTypeVarPop FunctionFunctionType = `VAR_POP`

const FunctionFunctionTypeVarSamp FunctionFunctionType = `VAR_SAMP`

// String representation for [fmt.Print]
func (f *FunctionFunctionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionFunctionType) Set(v string) error {
	switch v {
	case `APPROX_COUNT_DISTINCT`, `APPROX_PERCENTILE`, `AVG`, `COUNT`, `FIRST`, `LAST`, `MAX`, `MIN`, `STDDEV_POP`, `STDDEV_SAMP`, `SUM`, `VAR_POP`, `VAR_SAMP`:
		*f = FunctionFunctionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROX_COUNT_DISTINCT", "APPROX_PERCENTILE", "AVG", "COUNT", "FIRST", "LAST", "MAX", "MIN", "STDDEV_POP", "STDDEV_SAMP", "SUM", "VAR_POP", "VAR_SAMP"`, v)
	}
}

// Values returns all possible values for FunctionFunctionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionFunctionType) Values() []FunctionFunctionType {
	return []FunctionFunctionType{
		FunctionFunctionTypeApproxCountDistinct,
		FunctionFunctionTypeApproxPercentile,
		FunctionFunctionTypeAvg,
		FunctionFunctionTypeCount,
		FunctionFunctionTypeFirst,
		FunctionFunctionTypeLast,
		FunctionFunctionTypeMax,
		FunctionFunctionTypeMin,
		FunctionFunctionTypeStddevPop,
		FunctionFunctionTypeStddevSamp,
		FunctionFunctionTypeSum,
		FunctionFunctionTypeVarPop,
		FunctionFunctionTypeVarSamp,
	}
}

// Type always returns FunctionFunctionType to satisfy [pflag.Value] interface
func (f *FunctionFunctionType) Type() string {
	return "FunctionFunctionType"
}

type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName string `json:"-" url:"experiment_name"`
}

type GetExperimentByNameResponse struct {
	// Experiment details.
	Experiment *Experiment `json:"experiment,omitempty"`
}

type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

type GetExperimentPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ExperimentPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId string `json:"-" url:"experiment_id"`
}

type GetExperimentResponse struct {
	// Experiment details.
	Experiment *Experiment `json:"experiment,omitempty"`
}

type GetFeatureLineageRequest struct {
	// The name of the feature.
	FeatureName string `json:"-" url:"-"`
	// The full name of the feature table in Unity Catalog.
	TableName string `json:"-" url:"-"`
}

type GetFeatureRequest struct {
	// Name of the feature to get.
	FullName string `json:"-" url:"-"`
}

type GetFeatureTagRequest struct {
	FeatureName string `json:"-" url:"-"`

	Key string `json:"-" url:"-"`

	TableName string `json:"-" url:"-"`
}

type GetForecastingExperimentRequest struct {
	// The unique ID of a forecasting experiment
	ExperimentId string `json:"-" url:"-"`
}

type GetHistoryRequest struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Name of the metric.
	MetricKey string `json:"-" url:"metric_key"`
	// Token indicating the page of metric histories to fetch.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run from which to fetch
	// metric values. This field will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetHistoryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetHistoryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
	// List of stages.
	Stages []string `json:"stages,omitempty"`
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
}

type GetLoggedModelRequest struct {
	// The ID of the logged model to retrieve.
	ModelId string `json:"-" url:"-"`
}

type GetLoggedModelResponse struct {
	// The retrieved logged model.
	Model *LoggedModel `json:"model,omitempty"`
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric if `max_results` is not specified in
	// the request or if the total count of metrics returned is less than the
	// service level pagination threshold. Otherwise, this is one page of
	// results.
	Metrics []Metric `json:"metrics,omitempty"`
	// A token that can be used to issue a query for the next page of metric
	// history values. A missing token indicates that no additional metrics are
	// available to fetch.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetMetricHistoryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetMetricHistoryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"-" url:"name"`
}

type GetModelResponse struct {
	RegisteredModelDatabricks *ModelDatabricks `json:"registered_model_databricks,omitempty"`
}

type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

type GetModelVersionDownloadUriResponse struct {
	// URI corresponding to where artifacts for this model version are stored.
	ArtifactUri string `json:"artifact_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetModelVersionDownloadUriResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetModelVersionDownloadUriResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetModelVersionRequest struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

type GetModelVersionResponse struct {
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

type GetOnlineStoreRequest struct {
	// Name of the online store to get.
	Name string `json:"-" url:"-"`
}

type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

type GetRegisteredModelPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []RegisteredModelPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId string `json:"-" url:"run_id"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRunRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *HttpUrlSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s HttpUrlSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type HttpUrlSpecWithoutSecret struct {
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *HttpUrlSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s HttpUrlSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Tag for a dataset input.
type InputTag struct {
	// The tag key.
	Key string `json:"key"`
	// The tag value.
	Value string `json:"value"`
}

type JobSpec struct {
	// The personal access token used to authorize webhook's job runs.
	AccessToken string `json:"access_token"`
	// ID of the job that the webhook runs.
	JobId string `json:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	JobId string `json:"job_id,omitempty"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Feature for model version. ([ML-57150] Renamed from Feature to LinkedFeature)
type LinkedFeature struct {
	// Feature name
	FeatureName string `json:"feature_name,omitempty"`
	// Feature table id
	FeatureTableId string `json:"feature_table_id,omitempty"`
	// Feature table name
	FeatureTableName string `json:"feature_table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LinkedFeature) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LinkedFeature) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListArtifactsRequest struct {
	// The token indicating the page of artifact results to fetch. `page_token`
	// is not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path string `json:"-" url:"path,omitempty"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListArtifactsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListArtifactsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListArtifactsResponse struct {
	// The file location and metadata for artifacts.
	Files []FileInfo `json:"files,omitempty"`
	// The token that can be used to retrieve the next page of artifact results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The root artifact directory for the run.
	RootUri string `json:"root_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListArtifactsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListArtifactsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// Token indicating the page of experiments to fetch
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType ViewType `json:"-" url:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExperimentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExperimentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExperimentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExperimentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFeatureTagsRequest struct {
	FeatureName string `json:"-" url:"-"`
	// The maximum number of results to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	TableName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFeatureTagsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFeatureTagsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for ListFeatureTag.
type ListFeatureTagsResponse struct {
	FeatureTags []FeatureTag `json:"feature_tags,omitempty"`
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFeatureTagsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFeatureTagsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFeaturesRequest struct {
	// The maximum number of results to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFeaturesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFeaturesResponse struct {
	// List of features.
	Features []Feature `json:"features,omitempty"`
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFeaturesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFeaturesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	NextPageToken string `json:"next_page_token,omitempty"`

	RegisteredModels []Model `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListOnlineStoresRequest struct {
	// The maximum number of results to return. Defaults to 100 if not
	// specified.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListOnlineStoresRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListOnlineStoresRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListOnlineStoresResponse struct {
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of online stores.
	OnlineStores []OnlineStore `json:"online_stores,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListOnlineStoresResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListOnlineStoresResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken string `json:"next_page_token,omitempty"`
	// Array of registry webhooks.
	Webhooks []RegistryWebhook `json:"webhooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRegistryWebhooks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegistryWebhooks) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTransitionRequestsRequest struct {
	// Name of the registered model.
	Name string `json:"-" url:"name"`
	// Version of the model.
	Version string `json:"-" url:"version"`
}

type ListTransitionRequestsResponse struct {
	// Array of open transition requests.
	Requests []Activity `json:"requests,omitempty"`
}

type ListWebhooksRequest struct {
	// Events that trigger the webhook. * `MODEL_VERSION_CREATED`: A new model
	// version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	//
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events []RegistryWebhookEvent `json:"-" url:"events,omitempty"`

	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// Registered model name If not specified, all webhooks associated with the
	// specified events are listed, regardless of their associated model.
	ModelName string `json:"-" url:"model_name,omitempty"`
	// Token indicating the page of artifact results to fetch
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWebhooksRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWebhooksRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogBatch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogBatch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LogInputs struct {
	// Dataset inputs
	Datasets []DatasetInput `json:"datasets,omitempty"`
	// Model inputs
	Models []ModelInput `json:"models,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id"`
}

type LogLoggedModelParamsRequest struct {
	// The ID of the logged model to log params for.
	ModelId string `json:"-" url:"-"`
	// Parameters to attach to the model.
	Params []LoggedModelParameter `json:"params,omitempty"`
}

type LogMetric struct {
	// Dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	DatasetName string `json:"dataset_name,omitempty"`
	// Name of the metric.
	Key string `json:"key"`
	// ID of the logged model associated with the metric, if applicable
	ModelId string `json:"model_id,omitempty"`
	// ID of the run under which to log the metric. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Step at which to log the metric
	Step int64 `json:"step,omitempty"`
	// Unix timestamp in milliseconds at the time metric was logged.
	Timestamp int64 `json:"timestamp"`
	// Double value of the metric being logged.
	Value float64 `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogMetric) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogMetric) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LogModel struct {
	// MLmodel file in json format.
	ModelJson string `json:"model_json,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LogOutputsRequest struct {
	// The model outputs from the Run.
	Models []ModelOutput `json:"models,omitempty"`
	// The ID of the Run from which to log outputs.
	RunId string `json:"run_id"`
}

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	Key string `json:"key"`
	// ID of the run under which to log the param. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the param being logged. Maximum size is 500 bytes.
	Value string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogParam) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogParam) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A logged model message includes logged model attributes, tags, registration
// info, params, and linked run metrics.
type LoggedModel struct {
	// The params and metrics attached to the logged model.
	Data *LoggedModelData `json:"data,omitempty"`
	// The logged model attributes such as model ID, status, tags, etc.
	Info *LoggedModelInfo `json:"info,omitempty"`
}

// A LoggedModelData message includes logged model params and linked metrics.
type LoggedModelData struct {
	// Performance metrics linked to the model.
	Metrics []Metric `json:"metrics,omitempty"`
	// Immutable string key-value pairs of the model.
	Params []LoggedModelParameter `json:"params,omitempty"`
}

// A LoggedModelInfo includes logged model attributes, tags, and registration
// info.
type LoggedModelInfo struct {
	// The URI of the directory where model artifacts are stored.
	ArtifactUri string `json:"artifact_uri,omitempty"`
	// The timestamp when the model was created in milliseconds since the UNIX
	// epoch.
	CreationTimestampMs int64 `json:"creation_timestamp_ms,omitempty"`
	// The ID of the user or principal that created the model.
	CreatorId int64 `json:"creator_id,omitempty"`
	// The ID of the experiment that owns the model.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The timestamp when the model was last updated in milliseconds since the
	// UNIX epoch.
	LastUpdatedTimestampMs int64 `json:"last_updated_timestamp_ms,omitempty"`
	// The unique identifier for the logged model.
	ModelId string `json:"model_id,omitempty"`
	// The type of model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	ModelType string `json:"model_type,omitempty"`
	// The name of the model.
	Name string `json:"name,omitempty"`
	// The ID of the run that created the model.
	SourceRunId string `json:"source_run_id,omitempty"`
	// The status of whether or not the model is ready for use.
	Status LoggedModelStatus `json:"status,omitempty"`
	// Details on the current model status.
	StatusMessage string `json:"status_message,omitempty"`
	// Mutable string key-value pairs set on the model.
	Tags []LoggedModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LoggedModelInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LoggedModelInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Parameter associated with a LoggedModel.
type LoggedModelParameter struct {
	// The key identifying this param.
	Key string `json:"key,omitempty"`
	// The value of this param.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LoggedModelParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LoggedModelParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A LoggedModelStatus enum value represents the status of a logged model.
type LoggedModelStatus string

const LoggedModelStatusLoggedModelPending LoggedModelStatus = `LOGGED_MODEL_PENDING`

const LoggedModelStatusLoggedModelReady LoggedModelStatus = `LOGGED_MODEL_READY`

const LoggedModelStatusLoggedModelUploadFailed LoggedModelStatus = `LOGGED_MODEL_UPLOAD_FAILED`

// String representation for [fmt.Print]
func (f *LoggedModelStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LoggedModelStatus) Set(v string) error {
	switch v {
	case `LOGGED_MODEL_PENDING`, `LOGGED_MODEL_READY`, `LOGGED_MODEL_UPLOAD_FAILED`:
		*f = LoggedModelStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LOGGED_MODEL_PENDING", "LOGGED_MODEL_READY", "LOGGED_MODEL_UPLOAD_FAILED"`, v)
	}
}

// Values returns all possible values for LoggedModelStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *LoggedModelStatus) Values() []LoggedModelStatus {
	return []LoggedModelStatus{
		LoggedModelStatusLoggedModelPending,
		LoggedModelStatusLoggedModelReady,
		LoggedModelStatusLoggedModelUploadFailed,
	}
}

// Type always returns LoggedModelStatus to satisfy [pflag.Value] interface
func (f *LoggedModelStatus) Type() string {
	return "LoggedModelStatus"
}

// Tag for a LoggedModel.
type LoggedModelTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LoggedModelTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LoggedModelTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metric associated with a run, represented as a key-value pair.
type Metric struct {
	// The dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	DatasetName string `json:"dataset_name,omitempty"`
	// The key identifying the metric.
	Key string `json:"key,omitempty"`
	// The ID of the logged model or registered model version associated with
	// the metric, if applicable.
	ModelId string `json:"model_id,omitempty"`
	// The ID of the run containing the metric.
	RunId string `json:"run_id,omitempty"`
	// The step at which the metric was logged.
	Step int64 `json:"step,omitempty"`
	// The timestamp at which the metric was recorded.
	Timestamp int64 `json:"timestamp,omitempty"`
	// The value of the metric.
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Metric) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Metric) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Model struct {
	// Timestamp recorded when this `registered_model` was created.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Description of this `registered_model`.
	Description string `json:"description,omitempty"`
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	LatestVersions []ModelVersion `json:"latest_versions,omitempty"`
	// Unique name for the model.
	Name string `json:"name,omitempty"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	Tags []ModelTag `json:"tags,omitempty"`
	// User that created this `registered_model`
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Model) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Model) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ModelDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User-specified description for the object.
	Description string `json:"description,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Last update time of the object, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Array of model versions, each the latest version for its stage.
	LatestVersions []ModelVersion `json:"latest_versions,omitempty"`
	// Name of the model.
	Name string `json:"name,omitempty"`
	// Permission level granted for the requesting user on this registered model
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Array of tags associated with the model.
	Tags []ModelTag `json:"tags,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelDatabricks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelDatabricks) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput struct {
	// The unique identifier of the model.
	ModelId string `json:"model_id"`
}

// Represents a LoggedModel output of a Run.
type ModelOutput struct {
	// The unique identifier of the model.
	ModelId string `json:"model_id"`
	// The step at which the model was produced.
	Step int64 `json:"step"`
}

// Tag for a registered model
type ModelTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ModelVersion struct {
	// Timestamp recorded when this `model_version` was created.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Current stage for this `model_version`.
	CurrentStage string `json:"current_stage,omitempty"`
	// Description of this `model_version`.
	Description string `json:"description,omitempty"`
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Unique name of the model
	Name string `json:"name,omitempty"`
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	RunId string `json:"run_id,omitempty"`
	// Run Link: Direct link to the run that generated this version
	RunLink string `json:"run_link,omitempty"`
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	Source string `json:"source,omitempty"`
	// Current status of `model_version`
	Status ModelVersionStatus `json:"status,omitempty"`
	// Details on current `status`, if it is pending or failed.
	StatusMessage string `json:"status_message,omitempty"`
	// Tags: Additional metadata key-value pairs for this `model_version`.
	Tags []ModelVersionTag `json:"tags,omitempty"`
	// User that created this `model_version`.
	UserId string `json:"user_id,omitempty"`
	// Model's version number.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelVersion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ModelVersionDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	CurrentStage string `json:"current_stage,omitempty"`
	// User-specified description for the object.
	Description string `json:"description,omitempty"`
	// Email Subscription Status: This is the subscription status of the user to
	// the model version Users get subscribed by interacting with the model
	// version.
	EmailSubscriptionStatus RegistryEmailSubscriptionType `json:"email_subscription_status,omitempty"`
	// Feature lineage of `model_version`.
	FeatureList *FeatureList `json:"feature_list,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the model.
	Name string `json:"name,omitempty"`
	// Open requests for this `model_versions`. Gap in sequence number is
	// intentional and is done in order to match field sequence numbers of
	// `ModelVersion` proto message
	OpenRequests []Activity `json:"open_requests,omitempty"`

	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	RunId string `json:"run_id,omitempty"`
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	RunLink string `json:"run_link,omitempty"`
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	Source string `json:"source,omitempty"`

	Status Status `json:"status,omitempty"`
	// Details on the current status, for example why registration failed.
	StatusMessage string `json:"status_message,omitempty"`
	// Array of tags that are associated with the model version.
	Tags []ModelVersionTag `json:"tags,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`
	// Version of the model.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelVersionDatabricks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersionDatabricks) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type ModelVersionStatus string

// Request to register a new model version has failed.
const ModelVersionStatusFailedRegistration ModelVersionStatus = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const ModelVersionStatusPendingRegistration ModelVersionStatus = `PENDING_REGISTRATION`

// Model version is ready for use.
const ModelVersionStatusReady ModelVersionStatus = `READY`

// String representation for [fmt.Print]
func (f *ModelVersionStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelVersionStatus) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = ModelVersionStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Values returns all possible values for ModelVersionStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ModelVersionStatus) Values() []ModelVersionStatus {
	return []ModelVersionStatus{
		ModelVersionStatusFailedRegistration,
		ModelVersionStatusPendingRegistration,
		ModelVersionStatusReady,
	}
}

// Type always returns ModelVersionStatus to satisfy [pflag.Value] interface
func (f *ModelVersionStatus) Type() string {
	return "ModelVersionStatus"
}

type ModelVersionTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelVersionTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersionTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An OnlineStore is a logical database instance that stores and serves features
// online.
type OnlineStore struct {
	// The capacity of the online store. Valid values are "CU_1", "CU_2",
	// "CU_4", "CU_8".
	Capacity string `json:"capacity"`
	// The timestamp when the online store was created.
	CreationTime string `json:"creation_time,omitempty"`
	// The email of the creator of the online store.
	Creator string `json:"creator,omitempty"`
	// The name of the online store. This is the unique identifier for the
	// online store.
	Name string `json:"name"`
	// The number of read replicas for the online store. Defaults to 0.
	ReadReplicaCount int `json:"read_replica_count,omitempty"`
	// The current state of the online store.
	State OnlineStoreState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OnlineStore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineStore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type OnlineStoreState string

const OnlineStoreStateAvailable OnlineStoreState = `AVAILABLE`

const OnlineStoreStateDeleting OnlineStoreState = `DELETING`

const OnlineStoreStateFailingOver OnlineStoreState = `FAILING_OVER`

const OnlineStoreStateStarting OnlineStoreState = `STARTING`

const OnlineStoreStateStopped OnlineStoreState = `STOPPED`

const OnlineStoreStateUpdating OnlineStoreState = `UPDATING`

// String representation for [fmt.Print]
func (f *OnlineStoreState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OnlineStoreState) Set(v string) error {
	switch v {
	case `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`:
		*f = OnlineStoreState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVAILABLE", "DELETING", "FAILING_OVER", "STARTING", "STOPPED", "UPDATING"`, v)
	}
}

// Values returns all possible values for OnlineStoreState.
//
// There is no guarantee on the order of the values in the slice.
func (f *OnlineStoreState) Values() []OnlineStoreState {
	return []OnlineStoreState{
		OnlineStoreStateAvailable,
		OnlineStoreStateDeleting,
		OnlineStoreStateFailingOver,
		OnlineStoreStateStarting,
		OnlineStoreStateStopped,
		OnlineStoreStateUpdating,
	}
}

// Type always returns OnlineStoreState to satisfy [pflag.Value] interface
func (f *OnlineStoreState) Type() string {
	return "OnlineStoreState"
}

// Param associated with a run.
type Param struct {
	// Key identifying this param.
	Key string `json:"key,omitempty"`
	// Value associated with this param.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Param) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Param) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
type PermissionLevel string

const PermissionLevelCanCreateRegisteredModel PermissionLevel = `CAN_CREATE_REGISTERED_MODEL`

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanRead PermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_CREATE_REGISTERED_MODEL`, `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_CREATE_REGISTERED_MODEL", "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Values returns all possible values for PermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PermissionLevel) Values() []PermissionLevel {
	return []PermissionLevel{
		PermissionLevelCanCreateRegisteredModel,
		PermissionLevelCanEdit,
		PermissionLevelCanManage,
		PermissionLevelCanManageProductionVersions,
		PermissionLevelCanManageStagingVersions,
		PermissionLevelCanRead,
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

type PublishSpec struct {
	// The name of the target online store.
	OnlineStore string `json:"online_store"`
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName string `json:"online_table_name"`
	// The publish mode of the pipeline that syncs the online table with the
	// source table.
	PublishMode PublishSpecPublishMode `json:"publish_mode"`
}

type PublishSpecPublishMode string

const PublishSpecPublishModeContinuous PublishSpecPublishMode = `CONTINUOUS`

const PublishSpecPublishModeSnapshot PublishSpecPublishMode = `SNAPSHOT`

const PublishSpecPublishModeTriggered PublishSpecPublishMode = `TRIGGERED`

// String representation for [fmt.Print]
func (f *PublishSpecPublishMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PublishSpecPublishMode) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`:
		*f = PublishSpecPublishMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "SNAPSHOT", "TRIGGERED"`, v)
	}
}

// Values returns all possible values for PublishSpecPublishMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *PublishSpecPublishMode) Values() []PublishSpecPublishMode {
	return []PublishSpecPublishMode{
		PublishSpecPublishModeContinuous,
		PublishSpecPublishModeSnapshot,
		PublishSpecPublishModeTriggered,
	}
}

// Type always returns PublishSpecPublishMode to satisfy [pflag.Value] interface
func (f *PublishSpecPublishMode) Type() string {
	return "PublishSpecPublishMode"
}

type PublishTableRequest struct {
	// The specification for publishing the online table from the source table.
	PublishSpec PublishSpec `json:"publish_spec"`
	// The full three-part (catalog, schema, table) name of the source table.
	SourceTableName string `json:"-" url:"-"`
}

type PublishTableResponse struct {
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName string `json:"online_table_name,omitempty"`
	// The ID of the pipeline that syncs the online table with the source table.
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublishTableResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishTableResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelAccessControlResponse struct {
	// All permissions.
	AllPermissions []RegisteredModelPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type RegisteredModelPermissionLevel string

const RegisteredModelPermissionLevelCanEdit RegisteredModelPermissionLevel = `CAN_EDIT`

const RegisteredModelPermissionLevelCanManage RegisteredModelPermissionLevel = `CAN_MANAGE`

const RegisteredModelPermissionLevelCanManageProductionVersions RegisteredModelPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const RegisteredModelPermissionLevelCanManageStagingVersions RegisteredModelPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const RegisteredModelPermissionLevelCanRead RegisteredModelPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *RegisteredModelPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegisteredModelPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = RegisteredModelPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Values returns all possible values for RegisteredModelPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegisteredModelPermissionLevel) Values() []RegisteredModelPermissionLevel {
	return []RegisteredModelPermissionLevel{
		RegisteredModelPermissionLevelCanEdit,
		RegisteredModelPermissionLevelCanManage,
		RegisteredModelPermissionLevelCanManageProductionVersions,
		RegisteredModelPermissionLevelCanManageStagingVersions,
		RegisteredModelPermissionLevelCanRead,
	}
}

// Type always returns RegisteredModelPermissionLevel to satisfy [pflag.Value] interface
func (f *RegisteredModelPermissionLevel) Type() string {
	return "RegisteredModelPermissionLevel"
}

type RegisteredModelPermissions struct {
	AccessControlList []RegisteredModelAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelPermissionsRequest struct {
	AccessControlList []RegisteredModelAccessControlRequest `json:"access_control_list,omitempty"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

// .. note:: Experimental: This entity may change or be removed in a future
// release without warning. Email subscription types for registry notifications:
// - `ALL_EVENTS`: Subscribed to all events. - `DEFAULT`: Default subscription
// type. - `SUBSCRIBED`: Subscribed to notifications. - `UNSUBSCRIBED`: Not
// subscribed to notifications.
type RegistryEmailSubscriptionType string

// Subscribed to all events.
const RegistryEmailSubscriptionTypeAllEvents RegistryEmailSubscriptionType = `ALL_EVENTS`

// Default subscription type.
const RegistryEmailSubscriptionTypeDefault RegistryEmailSubscriptionType = `DEFAULT`

// Subscribed to notifications.
const RegistryEmailSubscriptionTypeSubscribed RegistryEmailSubscriptionType = `SUBSCRIBED`

// Not subscribed to notifications.
const RegistryEmailSubscriptionTypeUnsubscribed RegistryEmailSubscriptionType = `UNSUBSCRIBED`

// String representation for [fmt.Print]
func (f *RegistryEmailSubscriptionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryEmailSubscriptionType) Set(v string) error {
	switch v {
	case `ALL_EVENTS`, `DEFAULT`, `SUBSCRIBED`, `UNSUBSCRIBED`:
		*f = RegistryEmailSubscriptionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_EVENTS", "DEFAULT", "SUBSCRIBED", "UNSUBSCRIBED"`, v)
	}
}

// Values returns all possible values for RegistryEmailSubscriptionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryEmailSubscriptionType) Values() []RegistryEmailSubscriptionType {
	return []RegistryEmailSubscriptionType{
		RegistryEmailSubscriptionTypeAllEvents,
		RegistryEmailSubscriptionTypeDefault,
		RegistryEmailSubscriptionTypeSubscribed,
		RegistryEmailSubscriptionTypeUnsubscribed,
	}
}

// Type always returns RegistryEmailSubscriptionType to satisfy [pflag.Value] interface
func (f *RegistryEmailSubscriptionType) Type() string {
	return "RegistryEmailSubscriptionType"
}

type RegistryWebhook struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events,omitempty"`

	HttpUrlSpec *HttpUrlSpecWithoutSecret `json:"http_url_spec,omitempty"`
	// Webhook ID
	Id string `json:"id,omitempty"`

	JobSpec *JobSpecWithoutSecret `json:"job_spec,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the model whose events would trigger this webhook.
	ModelName string `json:"model_name,omitempty"`

	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegistryWebhook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegistryWebhook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// String representation for [fmt.Print]
func (f *RegistryWebhookEvent) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryWebhookEvent) Set(v string) error {
	switch v {
	case `COMMENT_CREATED`, `MODEL_VERSION_CREATED`, `MODEL_VERSION_TAG_SET`, `MODEL_VERSION_TRANSITIONED_STAGE`, `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`, `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`, `MODEL_VERSION_TRANSITIONED_TO_STAGING`, `REGISTERED_MODEL_CREATED`, `TRANSITION_REQUEST_CREATED`, `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`, `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`, `TRANSITION_REQUEST_TO_STAGING_CREATED`:
		*f = RegistryWebhookEvent(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMMENT_CREATED", "MODEL_VERSION_CREATED", "MODEL_VERSION_TAG_SET", "MODEL_VERSION_TRANSITIONED_STAGE", "MODEL_VERSION_TRANSITIONED_TO_ARCHIVED", "MODEL_VERSION_TRANSITIONED_TO_PRODUCTION", "MODEL_VERSION_TRANSITIONED_TO_STAGING", "REGISTERED_MODEL_CREATED", "TRANSITION_REQUEST_CREATED", "TRANSITION_REQUEST_TO_ARCHIVED_CREATED", "TRANSITION_REQUEST_TO_PRODUCTION_CREATED", "TRANSITION_REQUEST_TO_STAGING_CREATED"`, v)
	}
}

// Values returns all possible values for RegistryWebhookEvent.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryWebhookEvent) Values() []RegistryWebhookEvent {
	return []RegistryWebhookEvent{
		RegistryWebhookEventCommentCreated,
		RegistryWebhookEventModelVersionCreated,
		RegistryWebhookEventModelVersionTagSet,
		RegistryWebhookEventModelVersionTransitionedStage,
		RegistryWebhookEventModelVersionTransitionedToArchived,
		RegistryWebhookEventModelVersionTransitionedToProduction,
		RegistryWebhookEventModelVersionTransitionedToStaging,
		RegistryWebhookEventRegisteredModelCreated,
		RegistryWebhookEventTransitionRequestCreated,
		RegistryWebhookEventTransitionRequestToArchivedCreated,
		RegistryWebhookEventTransitionRequestToProductionCreated,
		RegistryWebhookEventTransitionRequestToStagingCreated,
	}
}

// Type always returns RegistryWebhookEvent to satisfy [pflag.Value] interface
func (f *RegistryWebhookEvent) Type() string {
	return "RegistryWebhookEvent"
}

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens.
//
// * `DISABLED`: Webhook is not triggered.
//
// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is not
// triggered on a real event.
type RegistryWebhookStatus string

// Webhook is triggered when an associated event happens.
const RegistryWebhookStatusActive RegistryWebhookStatus = `ACTIVE`

// Webhook is not triggered.
const RegistryWebhookStatusDisabled RegistryWebhookStatus = `DISABLED`

// Webhook can be triggered through the test endpoint, but is not triggered on a
// real event.
const RegistryWebhookStatusTestMode RegistryWebhookStatus = `TEST_MODE`

// String representation for [fmt.Print]
func (f *RegistryWebhookStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryWebhookStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `DISABLED`, `TEST_MODE`:
		*f = RegistryWebhookStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DISABLED", "TEST_MODE"`, v)
	}
}

// Values returns all possible values for RegistryWebhookStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryWebhookStatus) Values() []RegistryWebhookStatus {
	return []RegistryWebhookStatus{
		RegistryWebhookStatusActive,
		RegistryWebhookStatusDisabled,
		RegistryWebhookStatusTestMode,
	}
}

// Type always returns RegistryWebhookStatus to satisfy [pflag.Value] interface
func (f *RegistryWebhookStatus) Type() string {
	return "RegistryWebhookStatus"
}

// Details required to identify and reject a model version stage transition
// request.
type RejectTransitionRequest struct {
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RejectTransitionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RejectTransitionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RejectTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity *Activity `json:"activity,omitempty"`
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
	// If provided, updates the name for this `registered_model`.
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RenameModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RenameModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RenameModelResponse struct {
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId string `json:"run_id"`
}

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	ExperimentId string `json:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	MaxRuns int `json:"max_runs,omitempty"`
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	MinTimestampMillis int64 `json:"min_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestoreRuns) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestoreRuns) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestoreRunsResponse struct {
	// The number of runs restored.
	RunsRestored int `json:"runs_restored,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestoreRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestoreRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A single run.
type Run struct {
	// Run data.
	Data *RunData `json:"data,omitempty"`
	// Run metadata.
	Info *RunInfo `json:"info,omitempty"`
	// Run inputs.
	Inputs *RunInputs `json:"inputs,omitempty"`
}

// Run data (metrics, params, and tags).
type RunData struct {
	// Run metrics.
	Metrics []Metric `json:"metrics,omitempty"`
	// Run parameters.
	Params []Param `json:"params,omitempty"`
	// Additional metadata key-value pairs.
	Tags []RunTag `json:"tags,omitempty"`
}

// Metadata of a single run.
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
	// The name of the run.
	RunName string `json:"run_name,omitempty"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of a run.
type RunInfoStatus string

const RunInfoStatusFailed RunInfoStatus = `FAILED`

const RunInfoStatusFinished RunInfoStatus = `FINISHED`

const RunInfoStatusKilled RunInfoStatus = `KILLED`

const RunInfoStatusRunning RunInfoStatus = `RUNNING`

const RunInfoStatusScheduled RunInfoStatus = `SCHEDULED`

// String representation for [fmt.Print]
func (f *RunInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunInfoStatus) Set(v string) error {
	switch v {
	case `FAILED`, `FINISHED`, `KILLED`, `RUNNING`, `SCHEDULED`:
		*f = RunInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "FINISHED", "KILLED", "RUNNING", "SCHEDULED"`, v)
	}
}

// Values returns all possible values for RunInfoStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunInfoStatus) Values() []RunInfoStatus {
	return []RunInfoStatus{
		RunInfoStatusFailed,
		RunInfoStatusFinished,
		RunInfoStatusKilled,
		RunInfoStatusRunning,
		RunInfoStatusScheduled,
	}
}

// Type always returns RunInfoStatus to satisfy [pflag.Value] interface
func (f *RunInfoStatus) Type() string {
	return "RunInfoStatus"
}

// Run inputs.
type RunInputs struct {
	// Run metrics.
	DatasetInputs []DatasetInput `json:"dataset_inputs,omitempty"`
	// Model inputs to the Run.
	ModelInputs []ModelInput `json:"model_inputs,omitempty"`
}

// Tag for a run.
type RunTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	ViewType ViewType `json:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchExperiments) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchExperiments) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchExperimentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchExperimentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchLoggedModelsDataset struct {
	// The digest of the dataset.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset.
	DatasetName string `json:"dataset_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchLoggedModelsDataset) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchLoggedModelsDataset) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchLoggedModelsOrderBy struct {
	// Whether the search results order is ascending or not.
	Ascending bool `json:"ascending,omitempty"`
	// If ``field_name`` refers to a metric, this field specifies the digest of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name and digest will be considered for ordering. This
	// field may only be set if ``dataset_name`` is also set.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// If ``field_name`` refers to a metric, this field specifies the name of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name will be considered for ordering. This field may
	// only be set if ``field_name`` refers to a metric.
	DatasetName string `json:"dataset_name,omitempty"`
	// The name of the field to order by, e.g. "metrics.accuracy".
	FieldName string `json:"field_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchLoggedModelsOrderBy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchLoggedModelsOrderBy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchLoggedModelsRequest struct {
	// List of datasets on which to apply the metrics filter clauses. For
	// example, a filter with `metrics.accuracy > 0.9` and dataset info with
	// name "test_dataset" means we will return all logged models with accuracy
	// > 0.9 on the test_dataset. Metric values from ANY dataset matching the
	// criteria are considered. If no datasets are specified, then metrics
	// across all datasets are considered in the filter.
	Datasets []SearchLoggedModelsDataset `json:"datasets,omitempty"`
	// The IDs of the experiments in which to search for logged models.
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	// A filter expression over logged model info and data that allows returning
	// a subset of logged models. The syntax is a subset of SQL that supports
	// AND'ing together binary operations.
	//
	// Example: ``params.alpha < 0.3 AND metrics.accuracy > 0.9``.
	Filter string `json:"filter,omitempty"`
	// The maximum number of Logged Models to return. The maximum limit is 50.
	MaxResults int `json:"max_results,omitempty"`
	// The list of columns for ordering the results, with additional fields for
	// sorting criteria.
	OrderBy []SearchLoggedModelsOrderBy `json:"order_by,omitempty"`
	// The token indicating the page of logged models to fetch.
	PageToken string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchLoggedModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchLoggedModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchLoggedModelsResponse struct {
	// Logged models that match the search criteria.
	Models []LoggedModel `json:"models,omitempty"`
	// The token that can be used to retrieve the next page of logged models.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchLoggedModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchLoggedModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchModelVersionsRequest struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Max threshold is 10K.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Pagination token to go to next page based on previous search query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchModelVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchModelVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchModelVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchModelVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Registered Models that match the search criteria.
	RegisteredModels []Model `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	// A filter expression over params, metrics, and tags, that allows returning
	// a subset of runs. The syntax is a subset of SQL that supports ANDing
	// together binary operations between a param, metric, or tag and a
	// constant.
	//
	// Example: `metrics.rmse < 1 and params.model_class = 'LogisticRegression'`
	//
	// You can select columns with special characters (hyphen, space, period,
	// etc.) by using double quotes: `metrics."model class" = 'LinearRegression'
	// and tags."user-name" = 'Tomas'`
	//
	// Supported operators are `=`, `!=`, `>`, `>=`, `<`, and `<=`.
	Filter string `json:"filter,omitempty"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults int `json:"max_results,omitempty"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional `"DESC"` or `"ASC"` annotation, where `"ASC"`
	// is the default. Example: `["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"]`. Tiebreaks are done by start_time `DESC` followed by
	// `run_id` for runs with the same start time (and this is the default
	// ordering criterion if order_by is not provided).
	OrderBy []string `json:"order_by,omitempty"`
	// Token for the current page of runs.
	PageToken string `json:"page_token,omitempty"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType ViewType `json:"run_view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchRuns) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchRuns) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Runs that match the search criteria.
	Runs []Run `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	ExperimentId string `json:"experiment_id"`
	// Name of the tag. Keys up to 250 bytes in size are supported.
	Key string `json:"key"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	Value string `json:"value"`
}

type SetLoggedModelTagsRequest struct {
	// The ID of the logged model to set the tags on.
	ModelId string `json:"-" url:"-"`
	// The tags to set on the logged model.
	Tags []LoggedModelTag `json:"tags,omitempty"`
}

type SetModelTagRequest struct {
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

type SetTag struct {
	// Name of the tag. Keys up to 250 bytes in size are supported.
	Key string `json:"key"`
	// ID of the run under which to log the tag. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	Value string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SetTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SetTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type Status string

// Request to register a new model version has failed.
const StatusFailedRegistration Status = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const StatusPendingRegistration Status = `PENDING_REGISTRATION`

// Model version is ready for use.
const StatusReady Status = `READY`

// String representation for [fmt.Print]
func (f *Status) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Status) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = Status(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Values returns all possible values for Status.
//
// There is no guarantee on the order of the values in the slice.
func (f *Status) Values() []Status {
	return []Status{
		StatusFailedRegistration,
		StatusPendingRegistration,
		StatusReady,
	}
}

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

// Details required to test a registry webhook.
type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event RegistryWebhookEvent `json:"event,omitempty"`
	// Webhook ID
	Id string `json:"id"`
}

type TestRegistryWebhookResponse struct {
	// Body of the response from the webhook URL
	Body string `json:"body,omitempty"`
	// Status code returned by the webhook URL
	StatusCode int `json:"status_code,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TestRegistryWebhookResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TestRegistryWebhookResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TimeWindow struct {
	// The duration of the time window.
	Duration string `json:"duration"`
	// The offset of the time window.
	Offset string `json:"offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TimeWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TimeWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to transition a model version's stage.
type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TransitionModelVersionStageDatabricks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransitionModelVersionStageDatabricks) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []ActivityAction `json:"available_actions,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage string `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TransitionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransitionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TransitionStageResponse struct {
	// Updated model version
	ModelVersionDatabricks *ModelVersionDatabricks `json:"model_version_databricks,omitempty"`
}

// Details required to edit a comment on a model version.
type UpdateComment struct {
	// User-provided comment on the action.
	Comment string `json:"comment"`
	// Unique identifier of an activity
	Id string `json:"id"`
}

type UpdateCommentResponse struct {
	// Updated comment object
	Comment *CommentObject `json:"comment,omitempty"`
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateExperiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateExperiment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateFeatureRequest struct {
	// Feature to update.
	Feature Feature `json:"feature"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName string `json:"-" url:"-"`
	// The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateFeatureTagRequest struct {
	FeatureName string `json:"-" url:"-"`

	FeatureTag FeatureTag `json:"feature_tag"`

	Key string `json:"-" url:"-"`

	TableName string `json:"-" url:"-"`
	// The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateFeatureTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateFeatureTagRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description string `json:"description,omitempty"`
	// Registered model unique name identifier.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateModelResponse struct {
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description string `json:"description,omitempty"`
	// Name of the registered model
	Name string `json:"name"`
	// Model version number
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

type UpdateOnlineStoreRequest struct {
	// The name of the online store. This is the unique identifier for the
	// online store.
	Name string `json:"-" url:"-"`
	// Online store to update.
	OnlineStore OnlineStore `json:"online_store"`
	// The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

// Details required to update a registry webhook. Only the fields that need to
// be updated should be specified, and both `http_url_spec` and `job_spec`
// should not be specified in the same request.
type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events,omitempty"`

	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`
	// Webhook ID
	Id string `json:"id"`

	JobSpec *JobSpec `json:"job_spec,omitempty"`

	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRegistryWebhook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRegistryWebhook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime int64 `json:"end_time,omitempty"`
	// ID of the run to update. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// Updated name of the run.
	RunName string `json:"run_name,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run to update. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Updated status of the run.
	Status UpdateRunStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo *RunInfo `json:"run_info,omitempty"`
}

// Status of a run.
type UpdateRunStatus string

const UpdateRunStatusFailed UpdateRunStatus = `FAILED`

const UpdateRunStatusFinished UpdateRunStatus = `FINISHED`

const UpdateRunStatusKilled UpdateRunStatus = `KILLED`

const UpdateRunStatusRunning UpdateRunStatus = `RUNNING`

const UpdateRunStatusScheduled UpdateRunStatus = `SCHEDULED`

// String representation for [fmt.Print]
func (f *UpdateRunStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateRunStatus) Set(v string) error {
	switch v {
	case `FAILED`, `FINISHED`, `KILLED`, `RUNNING`, `SCHEDULED`:
		*f = UpdateRunStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "FINISHED", "KILLED", "RUNNING", "SCHEDULED"`, v)
	}
}

// Values returns all possible values for UpdateRunStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateRunStatus) Values() []UpdateRunStatus {
	return []UpdateRunStatus{
		UpdateRunStatusFailed,
		UpdateRunStatusFinished,
		UpdateRunStatusKilled,
		UpdateRunStatusRunning,
		UpdateRunStatusScheduled,
	}
}

// Type always returns UpdateRunStatus to satisfy [pflag.Value] interface
func (f *UpdateRunStatus) Type() string {
	return "UpdateRunStatus"
}

type UpdateWebhookResponse struct {
	Webhook *RegistryWebhook `json:"webhook,omitempty"`
}

// Qualifier for the view type.
type ViewType string

const ViewTypeActiveOnly ViewType = `ACTIVE_ONLY`

const ViewTypeAll ViewType = `ALL`

const ViewTypeDeletedOnly ViewType = `DELETED_ONLY`

// String representation for [fmt.Print]
func (f *ViewType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewType) Set(v string) error {
	switch v {
	case `ACTIVE_ONLY`, `ALL`, `DELETED_ONLY`:
		*f = ViewType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE_ONLY", "ALL", "DELETED_ONLY"`, v)
	}
}

// Values returns all possible values for ViewType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ViewType) Values() []ViewType {
	return []ViewType{
		ViewTypeActiveOnly,
		ViewTypeAll,
		ViewTypeDeletedOnly,
	}
}

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
}
