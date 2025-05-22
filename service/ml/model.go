// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Activity recorded for the action.
type Activity struct {
	// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied
	// the corresponding stage transition.
	//
	// * `REQUESTED_TRANSITION`: User requested the corresponding stage
	// transition.
	//
	// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
	//
	// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
	//
	// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
	//
	// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
	// archiving existing model versions in a stage.
	// Wire name: 'activity_type'
	ActivityType ActivityType
	// User-provided comment associated with the activity.
	// Wire name: 'comment'
	Comment string
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
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
	// Wire name: 'from_stage'
	FromStage Stage
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	// Wire name: 'system_comment'
	SystemComment string
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
	// Wire name: 'to_stage'
	ToStage Stage
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func (st *Activity) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &activityPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := activityFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Activity) MarshalJSON() ([]byte, error) {
	pb, err := activityToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// An action that a user (with sufficient permissions) could take on an
// activity. Valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
type ActivityAction string
type activityActionPb string

// Approve a transition request
const ActivityActionApproveTransitionRequest ActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityAction = `CANCEL_TRANSITION_REQUEST`

// Reject a transition request
const ActivityActionRejectTransitionRequest ActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *ActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `REJECT_TRANSITION_REQUEST`:
		*f = ActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Type always returns ActivityAction to satisfy [pflag.Value] interface
func (f *ActivityAction) Type() string {
	return "ActivityAction"
}

func activityActionToPb(st *ActivityAction) (*activityActionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := activityActionPb(*st)
	return &pb, nil
}

func activityActionFromPb(pb *activityActionPb) (*ActivityAction, error) {
	if pb == nil {
		return nil, nil
	}
	st := ActivityAction(*pb)
	return &st, nil
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
type activityTypePb string

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

// Type always returns ActivityType to satisfy [pflag.Value] interface
func (f *ActivityType) Type() string {
	return "ActivityType"
}

func activityTypeToPb(st *ActivityType) (*activityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := activityTypePb(*st)
	return &pb, nil
}

func activityTypeFromPb(pb *activityTypePb) (*ActivityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ActivityType(*pb)
	return &st, nil
}

type ApproveTransitionRequest struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	// Wire name: 'archive_existing_versions'
	ArchiveExistingVersions bool
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage Stage
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func (st *ApproveTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &approveTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := approveTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ApproveTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := approveTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ApproveTransitionRequestResponse struct {
	// Activity recorded for the action.
	// Wire name: 'activity'
	Activity *Activity
}

func (st *ApproveTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &approveTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := approveTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ApproveTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := approveTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ArtifactCredentialInfo struct {
	// A collection of HTTP headers that should be specified when uploading to
	// or downloading from the specified `signed_uri`.
	// Wire name: 'headers'
	Headers []ArtifactCredentialInfoHttpHeader
	// The path, relative to the Run's artifact root location, of the artifact
	// that can be accessed with the credential.
	// Wire name: 'path'
	Path string
	// The ID of the MLflow Run containing the artifact that can be accessed
	// with the credential.
	// Wire name: 'run_id'
	RunId string
	// The signed URI credential that provides access to the artifact.
	// Wire name: 'signed_uri'
	SignedUri string
	// The type of the signed credential URI (e.g., an AWS presigned URL or an
	// Azure Shared Access Signature URI).
	// Wire name: 'type'
	Type ArtifactCredentialType

	ForceSendFields []string `tf:"-"`
}

func (st *ArtifactCredentialInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &artifactCredentialInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := artifactCredentialInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ArtifactCredentialInfo) MarshalJSON() ([]byte, error) {
	pb, err := artifactCredentialInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ArtifactCredentialInfoHttpHeader struct {
	// The HTTP header name.
	// Wire name: 'name'
	Name string
	// The HTTP header value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *ArtifactCredentialInfoHttpHeader) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &artifactCredentialInfoHttpHeaderPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := artifactCredentialInfoHttpHeaderFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ArtifactCredentialInfoHttpHeader) MarshalJSON() ([]byte, error) {
	pb, err := artifactCredentialInfoHttpHeaderToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The type of a given artifact access credential
type ArtifactCredentialType string
type artifactCredentialTypePb string

const ArtifactCredentialTypeAwsPresignedUrl ArtifactCredentialType = `AWS_PRESIGNED_URL`

const ArtifactCredentialTypeAzureAdlsGen2SasUri ArtifactCredentialType = `AZURE_ADLS_GEN2_SAS_URI`

const ArtifactCredentialTypeAzureSasUri ArtifactCredentialType = `AZURE_SAS_URI`

const ArtifactCredentialTypeGcpSignedUrl ArtifactCredentialType = `GCP_SIGNED_URL`

// String representation for [fmt.Print]
func (f *ArtifactCredentialType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ArtifactCredentialType) Set(v string) error {
	switch v {
	case `AWS_PRESIGNED_URL`, `AZURE_ADLS_GEN2_SAS_URI`, `AZURE_SAS_URI`, `GCP_SIGNED_URL`:
		*f = ArtifactCredentialType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_PRESIGNED_URL", "AZURE_ADLS_GEN2_SAS_URI", "AZURE_SAS_URI", "GCP_SIGNED_URL"`, v)
	}
}

// Type always returns ArtifactCredentialType to satisfy [pflag.Value] interface
func (f *ArtifactCredentialType) Type() string {
	return "ArtifactCredentialType"
}

func artifactCredentialTypeToPb(st *ArtifactCredentialType) (*artifactCredentialTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := artifactCredentialTypePb(*st)
	return &pb, nil
}

func artifactCredentialTypeFromPb(pb *artifactCredentialTypePb) (*ArtifactCredentialType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ArtifactCredentialType(*pb)
	return &st, nil
}

// An action that a user (with sufficient permissions) could take on a comment.
// Valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type CommentActivityAction string
type commentActivityActionPb string

// Delete the comment
const CommentActivityActionDeleteComment CommentActivityAction = `DELETE_COMMENT`

// Edit the comment
const CommentActivityActionEditComment CommentActivityAction = `EDIT_COMMENT`

// String representation for [fmt.Print]
func (f *CommentActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommentActivityAction) Set(v string) error {
	switch v {
	case `DELETE_COMMENT`, `EDIT_COMMENT`:
		*f = CommentActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE_COMMENT", "EDIT_COMMENT"`, v)
	}
}

// Type always returns CommentActivityAction to satisfy [pflag.Value] interface
func (f *CommentActivityAction) Type() string {
	return "CommentActivityAction"
}

func commentActivityActionToPb(st *CommentActivityAction) (*commentActivityActionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := commentActivityActionPb(*st)
	return &pb, nil
}

func commentActivityActionFromPb(pb *commentActivityActionPb) (*CommentActivityAction, error) {
	if pb == nil {
		return nil, nil
	}
	st := CommentActivityAction(*pb)
	return &st, nil
}

// Comment details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	// Wire name: 'available_actions'
	AvailableActions []CommentActivityAction
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Comment ID
	// Wire name: 'id'
	Id string
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func (st *CommentObject) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &commentObjectPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := commentObjectFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CommentObject) MarshalJSON() ([]byte, error) {
	pb, err := commentObjectToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateComment struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Version of the model.
	// Wire name: 'version'
	Version string
}

func (st *CreateComment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCommentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCommentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateComment) MarshalJSON() ([]byte, error) {
	pb, err := createCommentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCommentResponse struct {
	// Comment details.
	// Wire name: 'comment'
	Comment *CommentObject
}

func (st *CreateCommentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCommentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCommentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCommentResponse) MarshalJSON() ([]byte, error) {
	pb, err := createCommentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateExperiment struct {
	// Location where all artifacts for the experiment are stored. If not
	// provided, the remote server will select an appropriate default.
	// Wire name: 'artifact_location'
	ArtifactLocation string
	// Experiment name.
	// Wire name: 'name'
	Name string
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	// Wire name: 'tags'
	Tags []ExperimentTag

	ForceSendFields []string `tf:"-"`
}

func (st *CreateExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExperiment) MarshalJSON() ([]byte, error) {
	pb, err := createExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	// Wire name: 'experiment_id'
	ExperimentId string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := createExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateForecastingExperimentRequest struct {
	// The column in the training table used to customize weights for each time
	// series.
	// Wire name: 'custom_weights_column'
	CustomWeightsColumn string
	// The path in the workspace to store the created experiment.
	// Wire name: 'experiment_path'
	ExperimentPath string
	// The time interval between consecutive rows in the time series data.
	// Possible values include: '1 second', '1 minute', '5 minutes', '10
	// minutes', '15 minutes', '30 minutes', 'Hourly', 'Daily', 'Weekly',
	// 'Monthly', 'Quarterly', 'Yearly'.
	// Wire name: 'forecast_granularity'
	ForecastGranularity string
	// The number of time steps into the future to make predictions, calculated
	// as a multiple of forecast_granularity. This value represents how far
	// ahead the model should forecast.
	// Wire name: 'forecast_horizon'
	ForecastHorizon int64
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store future feature data
	// for predictions.
	// Wire name: 'future_feature_data_path'
	FutureFeatureDataPath string
	// The region code(s) to automatically add holiday features. Currently
	// supports only one region.
	// Wire name: 'holiday_regions'
	HolidayRegions []string
	// Specifies the list of feature columns to include in model training. These
	// columns must exist in the training data and be of type string, numerical,
	// or boolean. If not specified, no additional features will be included.
	// Note: Certain columns are automatically handled: - Automatically
	// excluded: split_column, target_column, custom_weights_column. -
	// Automatically included: time_column.
	// Wire name: 'include_features'
	IncludeFeatures []string
	// The maximum duration for the experiment in minutes. The experiment stops
	// automatically if it exceeds this limit.
	// Wire name: 'max_runtime'
	MaxRuntime int64
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store predictions.
	// Wire name: 'prediction_data_path'
	PredictionDataPath string
	// The evaluation metric used to optimize the forecasting model.
	// Wire name: 'primary_metric'
	PrimaryMetric string
	// The fully qualified path of a Unity Catalog model, formatted as
	// catalog_name.schema_name.model_name, used to store the best model.
	// Wire name: 'register_to'
	RegisterTo string
	// // The column in the training table used for custom data splits. Values
	// must be 'train', 'validate', or 'test'.
	// Wire name: 'split_column'
	SplitColumn string
	// The column in the input training table used as the prediction target for
	// model training. The values in this column are used as the ground truth
	// for model training.
	// Wire name: 'target_column'
	TargetColumn string
	// The column in the input training table that represents each row's
	// timestamp.
	// Wire name: 'time_column'
	TimeColumn string
	// The column in the training table used to group the dataset for predicting
	// individual time series.
	// Wire name: 'timeseries_identifier_columns'
	TimeseriesIdentifierColumns []string
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used as training data for the
	// forecasting model.
	// Wire name: 'train_data_path'
	TrainDataPath string
	// List of frameworks to include for model tuning. Possible values are
	// 'Prophet', 'ARIMA', 'DeepAR'. An empty list includes all supported
	// frameworks.
	// Wire name: 'training_frameworks'
	TrainingFrameworks []string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createForecastingExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createForecastingExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateForecastingExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := createForecastingExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateForecastingExperimentResponse struct {
	// The unique ID of the created forecasting experiment
	// Wire name: 'experiment_id'
	ExperimentId string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateForecastingExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createForecastingExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createForecastingExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateForecastingExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := createForecastingExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateLoggedModelRequest struct {
	// The ID of the experiment that owns the model.
	// Wire name: 'experiment_id'
	ExperimentId string
	// The type of the model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	// Wire name: 'model_type'
	ModelType string
	// The name of the model (optional). If not specified one will be generated.
	// Wire name: 'name'
	Name string
	// Parameters attached to the model.
	// Wire name: 'params'
	Params []LoggedModelParameter
	// The ID of the run that created the model.
	// Wire name: 'source_run_id'
	SourceRunId string
	// Tags attached to the model.
	// Wire name: 'tags'
	Tags []LoggedModelTag

	ForceSendFields []string `tf:"-"`
}

func (st *CreateLoggedModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createLoggedModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createLoggedModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateLoggedModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := createLoggedModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateLoggedModelResponse struct {
	// The newly created logged model.
	// Wire name: 'model'
	Model *LoggedModel
}

func (st *CreateLoggedModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createLoggedModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createLoggedModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateLoggedModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := createLoggedModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateModelRequest struct {
	// Optional description for registered model.
	// Wire name: 'description'
	Description string
	// Register models under this name
	// Wire name: 'name'
	Name string
	// Additional metadata for registered model.
	// Wire name: 'tags'
	Tags []ModelTag

	ForceSendFields []string `tf:"-"`
}

func (st *CreateModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := createModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateModelResponse struct {

	// Wire name: 'registered_model'
	RegisteredModel *Model
}

func (st *CreateModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := createModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateModelVersionRequest struct {
	// Optional description for model version.
	// Wire name: 'description'
	Description string
	// Register model under this name
	// Wire name: 'name'
	Name string
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	// Wire name: 'run_id'
	RunId string
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	// Wire name: 'run_link'
	RunLink string
	// URI indicating the location of the model artifacts.
	// Wire name: 'source'
	Source string
	// Additional metadata for model version.
	// Wire name: 'tags'
	Tags []ModelVersionTag

	ForceSendFields []string `tf:"-"`
}

func (st *CreateModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := createModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	// Wire name: 'model_version'
	ModelVersion *ModelVersion
}

func (st *CreateModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := createModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string
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
	// Wire name: 'events'
	Events []RegistryWebhookEvent

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpec

	// Wire name: 'job_spec'
	JobSpec *JobSpec
	// Name of the model whose events would trigger this webhook.
	// Wire name: 'model_name'
	ModelName string
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	// Wire name: 'status'
	Status RegistryWebhookStatus

	ForceSendFields []string `tf:"-"`
}

func (st *CreateRegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := createRegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateRun struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
	// The name of the run.
	// Wire name: 'run_name'
	RunName string
	// Unix timestamp in milliseconds of when the run started.
	// Wire name: 'start_time'
	StartTime int64
	// Additional metadata for run.
	// Wire name: 'tags'
	Tags []RunTag
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRun) MarshalJSON() ([]byte, error) {
	pb, err := createRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateRunResponse struct {
	// The newly created run.
	// Wire name: 'run'
	Run *Run
}

func (st *CreateRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := createRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateTransitionRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage Stage
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := createTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateTransitionRequestResponse struct {
	// Transition request details.
	// Wire name: 'request'
	Request *TransitionRequest
}

func (st *CreateTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := createTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateWebhookResponse struct {

	// Wire name: 'webhook'
	Webhook *RegistryWebhook
}

func (st *CreateWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := createWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Dataset. Represents a reference to data used for training, testing, or
// evaluation during the model development process.
type Dataset struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	// Wire name: 'digest'
	Digest string
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	// Wire name: 'name'
	Name string
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	// Wire name: 'profile'
	Profile string
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	// Wire name: 'schema'
	Schema string
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	// Wire name: 'source'
	Source string
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	// Wire name: 'source_type'
	SourceType string

	ForceSendFields []string `tf:"-"`
}

func (st *Dataset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &datasetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := datasetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Dataset) MarshalJSON() ([]byte, error) {
	pb, err := datasetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// DatasetInput. Represents a dataset and input tags.
type DatasetInput struct {
	// The dataset being used as a Run input.
	// Wire name: 'dataset'
	Dataset Dataset
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	// Wire name: 'tags'
	Tags []InputTag
}

func (st *DatasetInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &datasetInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := datasetInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatasetInput) MarshalJSON() ([]byte, error) {
	pb, err := datasetInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a comment
type DeleteCommentRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *DeleteCommentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCommentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCommentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCommentRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCommentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteCommentResponse struct {
}

func (st *DeleteCommentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCommentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCommentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCommentResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteCommentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
}

func (st *DeleteExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExperiment) MarshalJSON() ([]byte, error) {
	pb, err := deleteExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteExperimentResponse struct {
}

func (st *DeleteExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a logged model
type DeleteLoggedModelRequest struct {
	// The ID of the logged model to delete.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
}

func (st *DeleteLoggedModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteLoggedModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteLoggedModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteLoggedModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteLoggedModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteLoggedModelResponse struct {
}

func (st *DeleteLoggedModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteLoggedModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteLoggedModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteLoggedModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteLoggedModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a tag on a logged model
type DeleteLoggedModelTagRequest struct {
	// The ID of the logged model to delete the tag from.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// The tag key.
	// Wire name: 'tag_key'
	TagKey string `tf:"-"`
}

func (st *DeleteLoggedModelTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteLoggedModelTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteLoggedModelTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteLoggedModelTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteLoggedModelTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteLoggedModelTagResponse struct {
}

func (st *DeleteLoggedModelTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteLoggedModelTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteLoggedModelTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteLoggedModelTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteLoggedModelTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a model
type DeleteModelRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st *DeleteModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteModelResponse struct {
}

func (st *DeleteModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a model tag
type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	// Wire name: 'key'
	Key string `tf:"-"`
	// Name of the registered model that the tag was logged under.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st *DeleteModelTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteModelTagResponse struct {
}

func (st *DeleteModelTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a model version.
type DeleteModelVersionRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st *DeleteModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteModelVersionResponse struct {
}

func (st *DeleteModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a model version tag
type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	// Wire name: 'key'
	Key string `tf:"-"`
	// Name of the registered model that the tag was logged under.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number that the tag was logged under.
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st *DeleteModelVersionTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteModelVersionTagResponse struct {
}

func (st *DeleteModelVersionTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteRun struct {
	// ID of the run to delete.
	// Wire name: 'run_id'
	RunId string
}

func (st *DeleteRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRun) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteRunResponse struct {
}

func (st *DeleteRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	// Wire name: 'experiment_id'
	ExperimentId string
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	// Wire name: 'max_runs'
	MaxRuns int
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	// Wire name: 'max_timestamp_millis'
	MaxTimestampMillis int64

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRuns) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteRunsResponse struct {
	// The number of runs deleted.
	// Wire name: 'runs_deleted'
	RunsDeleted int

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	// Wire name: 'key'
	Key string
	// ID of the run that the tag was logged under. Must be provided.
	// Wire name: 'run_id'
	RunId string
}

func (st *DeleteTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTag) MarshalJSON() ([]byte, error) {
	pb, err := deleteTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteTagResponse struct {
}

func (st *DeleteTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a transition request
type DeleteTransitionRequestRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string `tf:"-"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	// Wire name: 'creator'
	Creator string `tf:"-"`
	// Name of the model.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage DeleteTransitionRequestStage `tf:"-"`
	// Version of the model.
	// Wire name: 'version'
	Version string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteTransitionRequestRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTransitionRequestRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTransitionRequestRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTransitionRequestRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteTransitionRequestRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteTransitionRequestResponse struct {
}

func (st *DeleteTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteTransitionRequestStage string
type deleteTransitionRequestStagePb string

const DeleteTransitionRequestStageArchived DeleteTransitionRequestStage = `Archived`

const DeleteTransitionRequestStageNone DeleteTransitionRequestStage = `None`

const DeleteTransitionRequestStageProduction DeleteTransitionRequestStage = `Production`

const DeleteTransitionRequestStageStaging DeleteTransitionRequestStage = `Staging`

// String representation for [fmt.Print]
func (f *DeleteTransitionRequestStage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeleteTransitionRequestStage) Set(v string) error {
	switch v {
	case `Archived`, `None`, `Production`, `Staging`:
		*f = DeleteTransitionRequestStage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Archived", "None", "Production", "Staging"`, v)
	}
}

// Type always returns DeleteTransitionRequestStage to satisfy [pflag.Value] interface
func (f *DeleteTransitionRequestStage) Type() string {
	return "DeleteTransitionRequestStage"
}

func deleteTransitionRequestStageToPb(st *DeleteTransitionRequestStage) (*deleteTransitionRequestStagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := deleteTransitionRequestStagePb(*st)
	return &pb, nil
}

func deleteTransitionRequestStageFromPb(pb *deleteTransitionRequestStagePb) (*DeleteTransitionRequestStage, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeleteTransitionRequestStage(*pb)
	return &st, nil
}

// Delete a webhook
type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	// Wire name: 'id'
	Id string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteWebhookRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWebhookRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWebhookRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWebhookRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteWebhookRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteWebhookResponse struct {
}

func (st *DeleteWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// An experiment and its metadata.
type Experiment struct {
	// Location where artifacts for the experiment are stored.
	// Wire name: 'artifact_location'
	ArtifactLocation string
	// Creation time
	// Wire name: 'creation_time'
	CreationTime int64
	// Unique identifier for the experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
	// Last update time
	// Wire name: 'last_update_time'
	LastUpdateTime int64
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	// Wire name: 'lifecycle_stage'
	LifecycleStage string
	// Human readable name that identifies the experiment.
	// Wire name: 'name'
	Name string
	// Tags: Additional metadata key-value pairs.
	// Wire name: 'tags'
	Tags []ExperimentTag

	ForceSendFields []string `tf:"-"`
}

func (st *Experiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Experiment) MarshalJSON() ([]byte, error) {
	pb, err := experimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExperimentAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *ExperimentAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := experimentAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExperimentAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ExperimentPermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *ExperimentAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := experimentAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExperimentPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *ExperimentPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentPermission) MarshalJSON() ([]byte, error) {
	pb, err := experimentPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Permission level
type ExperimentPermissionLevel string
type experimentPermissionLevelPb string

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

// Type always returns ExperimentPermissionLevel to satisfy [pflag.Value] interface
func (f *ExperimentPermissionLevel) Type() string {
	return "ExperimentPermissionLevel"
}

func experimentPermissionLevelToPb(st *ExperimentPermissionLevel) (*experimentPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := experimentPermissionLevelPb(*st)
	return &pb, nil
}

func experimentPermissionLevelFromPb(pb *experimentPermissionLevelPb) (*ExperimentPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExperimentPermissionLevel(*pb)
	return &st, nil
}

type ExperimentPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []ExperimentAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func (st *ExperimentPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentPermissions) MarshalJSON() ([]byte, error) {
	pb, err := experimentPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExperimentPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *ExperimentPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := experimentPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExperimentPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ExperimentAccessControlRequest
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st *ExperimentPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := experimentPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A tag for an experiment.
type ExperimentTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *ExperimentTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentTag) MarshalJSON() ([]byte, error) {
	pb, err := experimentTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Metadata of a single artifact file or directory.
type FileInfo struct {
	// The size in bytes of the file. Unset for directories.
	// Wire name: 'file_size'
	FileSize int64
	// Whether the path is a directory.
	// Wire name: 'is_dir'
	IsDir bool
	// The path relative to the root artifact directory run.
	// Wire name: 'path'
	Path string

	ForceSendFields []string `tf:"-"`
}

func (st *FileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileInfo) MarshalJSON() ([]byte, error) {
	pb, err := fileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FinalizeLoggedModelRequest struct {
	// The ID of the logged model to finalize.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// Whether or not the model is ready for use.
	// ``"LOGGED_MODEL_UPLOAD_FAILED"`` indicates that something went wrong when
	// logging the model weights / agent code).
	// Wire name: 'status'
	Status LoggedModelStatus
}

func (st *FinalizeLoggedModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &finalizeLoggedModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := finalizeLoggedModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FinalizeLoggedModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := finalizeLoggedModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FinalizeLoggedModelResponse struct {
	// The updated logged model.
	// Wire name: 'model'
	Model *LoggedModel
}

func (st *FinalizeLoggedModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &finalizeLoggedModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := finalizeLoggedModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FinalizeLoggedModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := finalizeLoggedModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents a forecasting experiment with its unique identifier, URL, and
// state.
type ForecastingExperiment struct {
	// The unique ID for the forecasting experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
	// The URL to the forecasting experiment page.
	// Wire name: 'experiment_page_url'
	ExperimentPageUrl string
	// The current state of the forecasting experiment.
	// Wire name: 'state'
	State ForecastingExperimentState

	ForceSendFields []string `tf:"-"`
}

func (st *ForecastingExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forecastingExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forecastingExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForecastingExperiment) MarshalJSON() ([]byte, error) {
	pb, err := forecastingExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ForecastingExperimentState string
type forecastingExperimentStatePb string

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

// Type always returns ForecastingExperimentState to satisfy [pflag.Value] interface
func (f *ForecastingExperimentState) Type() string {
	return "ForecastingExperimentState"
}

func forecastingExperimentStateToPb(st *ForecastingExperimentState) (*forecastingExperimentStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := forecastingExperimentStatePb(*st)
	return &pb, nil
}

func forecastingExperimentStateFromPb(pb *forecastingExperimentStatePb) (*ForecastingExperimentState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ForecastingExperimentState(*pb)
	return &st, nil
}

// Get an experiment by name
type GetByNameRequest struct {
	// Name of the associated experiment.
	// Wire name: 'experiment_name'
	ExperimentName string `tf:"-"`
}

func (st *GetByNameRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getByNameRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getByNameRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetByNameRequest) MarshalJSON() ([]byte, error) {
	pb, err := getByNameRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get credentials to download trace data
type GetCredentialsForTraceDataDownloadRequest struct {
	// The ID of the trace to fetch artifact download credentials for.
	// Wire name: 'request_id'
	RequestId string `tf:"-"`
}

func (st *GetCredentialsForTraceDataDownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsForTraceDataDownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsForTraceDataDownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsForTraceDataDownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsForTraceDataDownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetCredentialsForTraceDataDownloadResponse struct {
	// The artifact download credentials for the specified trace data.
	// Wire name: 'credential_info'
	CredentialInfo *ArtifactCredentialInfo
}

func (st *GetCredentialsForTraceDataDownloadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsForTraceDataDownloadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsForTraceDataDownloadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsForTraceDataDownloadResponse) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsForTraceDataDownloadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get credentials to upload trace data
type GetCredentialsForTraceDataUploadRequest struct {
	// The ID of the trace to fetch artifact upload credentials for.
	// Wire name: 'request_id'
	RequestId string `tf:"-"`
}

func (st *GetCredentialsForTraceDataUploadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsForTraceDataUploadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsForTraceDataUploadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsForTraceDataUploadRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsForTraceDataUploadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetCredentialsForTraceDataUploadResponse struct {
	// The artifact upload credentials for the specified trace data.
	// Wire name: 'credential_info'
	CredentialInfo *ArtifactCredentialInfo
}

func (st *GetCredentialsForTraceDataUploadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsForTraceDataUploadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsForTraceDataUploadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsForTraceDataUploadResponse) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsForTraceDataUploadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetExperimentByNameResponse struct {
	// Experiment details.
	// Wire name: 'experiment'
	Experiment *Experiment
}

func (st *GetExperimentByNameResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentByNameResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentByNameResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentByNameResponse) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentByNameResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get experiment permission levels
type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st *GetExperimentPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetExperimentPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []ExperimentPermissionsDescription
}

func (st *GetExperimentPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get experiment permissions
type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st *GetExperimentPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get an experiment
type GetExperimentRequest struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st *GetExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetExperimentResponse struct {
	// Experiment details.
	// Wire name: 'experiment'
	Experiment *Experiment
}

func (st *GetExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a forecasting experiment
type GetForecastingExperimentRequest struct {
	// The unique ID of a forecasting experiment
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st *GetForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getForecastingExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getForecastingExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetForecastingExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := getForecastingExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get metric history for a run
type GetHistoryRequest struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Name of the metric.
	// Wire name: 'metric_key'
	MetricKey string `tf:"-"`
	// Token indicating the page of metric histories to fetch.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// ID of the run from which to fetch metric values. Must be provided.
	// Wire name: 'run_id'
	RunId string `tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run from which to fetch
	// metric values. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetHistoryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getHistoryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getHistoryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetHistoryRequest) MarshalJSON() ([]byte, error) {
	pb, err := getHistoryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string
	// List of stages.
	// Wire name: 'stages'
	Stages []string
}

func (st *GetLatestVersionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLatestVersionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLatestVersionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLatestVersionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getLatestVersionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	// Wire name: 'model_versions'
	ModelVersions []ModelVersion
}

func (st *GetLatestVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLatestVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLatestVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLatestVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getLatestVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a logged model
type GetLoggedModelRequest struct {
	// The ID of the logged model to retrieve.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
}

func (st *GetLoggedModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLoggedModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLoggedModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLoggedModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := getLoggedModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetLoggedModelResponse struct {
	// The retrieved logged model.
	// Wire name: 'model'
	Model *LoggedModel
}

func (st *GetLoggedModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLoggedModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLoggedModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLoggedModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := getLoggedModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric if `max_results` is not specified in
	// the request or if the total count of metrics returned is less than the
	// service level pagination threshold. Otherwise, this is one page of
	// results.
	// Wire name: 'metrics'
	Metrics []Metric
	// A token that can be used to issue a query for the next page of metric
	// history values. A missing token indicates that no additional metrics are
	// available to fetch.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *GetMetricHistoryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getMetricHistoryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getMetricHistoryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetMetricHistoryResponse) MarshalJSON() ([]byte, error) {
	pb, err := getMetricHistoryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get model
type GetModelRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st *GetModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := getModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetModelResponse struct {

	// Wire name: 'registered_model_databricks'
	RegisteredModelDatabricks *ModelDatabricks
}

func (st *GetModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := getModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a model version URI
type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st *GetModelVersionDownloadUriRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionDownloadUriRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionDownloadUriRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionDownloadUriRequest) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionDownloadUriRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetModelVersionDownloadUriResponse struct {
	// URI corresponding to where artifacts for this model version are stored.
	// Wire name: 'artifact_uri'
	ArtifactUri string

	ForceSendFields []string `tf:"-"`
}

func (st *GetModelVersionDownloadUriResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionDownloadUriResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionDownloadUriResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionDownloadUriResponse) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionDownloadUriResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a model version
type GetModelVersionRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetModelVersionResponse struct {

	// Wire name: 'model_version'
	ModelVersion *ModelVersion
}

func (st *GetModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get registered model permission levels
type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func (st *GetRegisteredModelPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRegisteredModelPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRegisteredModelPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRegisteredModelPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRegisteredModelPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetRegisteredModelPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []RegisteredModelPermissionsDescription
}

func (st *GetRegisteredModelPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRegisteredModelPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRegisteredModelPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRegisteredModelPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getRegisteredModelPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get registered model permissions
type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func (st *GetRegisteredModelPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRegisteredModelPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRegisteredModelPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRegisteredModelPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRegisteredModelPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a run
type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	// Wire name: 'run_id'
	RunId string `tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	// Wire name: 'run'
	Run *Run
}

func (st *GetRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := getRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type HttpUrlSpec struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	// Wire name: 'authorization'
	Authorization string
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	// Wire name: 'enable_ssl_verification'
	EnableSslVerification bool
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	// Wire name: 'secret'
	Secret string
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func (st *HttpUrlSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpUrlSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := httpUrlSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st HttpUrlSpec) MarshalJSON() ([]byte, error) {
	pb, err := httpUrlSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type HttpUrlSpecWithoutSecret struct {
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	// Wire name: 'enable_ssl_verification'
	EnableSslVerification bool
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func (st *HttpUrlSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpUrlSpecWithoutSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := httpUrlSpecWithoutSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st HttpUrlSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	pb, err := httpUrlSpecWithoutSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Tag for a dataset input.
type InputTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string
}

func (st *InputTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &inputTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := inputTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InputTag) MarshalJSON() ([]byte, error) {
	pb, err := inputTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobSpec struct {
	// The personal access token used to authorize webhook's job runs.
	// Wire name: 'access_token'
	AccessToken string
	// ID of the job that the webhook runs.
	// Wire name: 'job_id'
	JobId string
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	// Wire name: 'workspace_url'
	WorkspaceUrl string

	ForceSendFields []string `tf:"-"`
}

func (st *JobSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobSpec) MarshalJSON() ([]byte, error) {
	pb, err := jobSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	// Wire name: 'job_id'
	JobId string
	// URL of the workspace containing the job that this webhook runs. Defaults
	// to the workspace URL in which the webhook is created. If not specified,
	// the job’s workspace is assumed to be the same as the webhook’s.
	// Wire name: 'workspace_url'
	WorkspaceUrl string

	ForceSendFields []string `tf:"-"`
}

func (st *JobSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobSpecWithoutSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobSpecWithoutSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	pb, err := jobSpecWithoutSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List artifacts
type ListArtifactsRequest struct {
	// The token indicating the page of artifact results to fetch. `page_token`
	// is not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	// Wire name: 'path'
	Path string `tf:"-"`
	// ID of the run whose artifacts to list. Must be provided.
	// Wire name: 'run_id'
	RunId string `tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListArtifactsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listArtifactsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listArtifactsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListArtifactsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listArtifactsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListArtifactsResponse struct {
	// The file location and metadata for artifacts.
	// Wire name: 'files'
	Files []FileInfo
	// The token that can be used to retrieve the next page of artifact results.
	// Wire name: 'next_page_token'
	NextPageToken string
	// The root artifact directory for the run.
	// Wire name: 'root_uri'
	RootUri string

	ForceSendFields []string `tf:"-"`
}

func (st *ListArtifactsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listArtifactsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listArtifactsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListArtifactsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listArtifactsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List experiments
type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	// Wire name: 'max_results'
	MaxResults int64 `tf:"-"`
	// Token indicating the page of experiments to fetch
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	// Wire name: 'view_type'
	ViewType ViewType `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListExperimentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExperimentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExperimentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExperimentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listExperimentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	// Wire name: 'experiments'
	Experiments []Experiment
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *ListExperimentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExperimentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExperimentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExperimentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listExperimentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List artifacts for a logged model
type ListLoggedModelArtifactsRequest struct {
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	// Wire name: 'artifact_directory_path'
	ArtifactDirectoryPath string `tf:"-"`
	// The ID of the logged model for which to list the artifacts.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// Token indicating the page of artifact results to fetch. `page_token` is
	// not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListLoggedModelArtifactsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listLoggedModelArtifactsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listLoggedModelArtifactsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListLoggedModelArtifactsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listLoggedModelArtifactsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListLoggedModelArtifactsResponse struct {
	// File location and metadata for artifacts.
	// Wire name: 'files'
	Files []FileInfo
	// Token that can be used to retrieve the next page of artifact results
	// Wire name: 'next_page_token'
	NextPageToken string
	// Root artifact directory for the logged model.
	// Wire name: 'root_uri'
	RootUri string

	ForceSendFields []string `tf:"-"`
}

func (st *ListLoggedModelArtifactsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listLoggedModelArtifactsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listLoggedModelArtifactsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListLoggedModelArtifactsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listLoggedModelArtifactsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List models
type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Pagination token to go to the next page based on a previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'registered_models'
	RegisteredModels []Model

	ForceSendFields []string `tf:"-"`
}

func (st *ListModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	// Wire name: 'next_page_token'
	NextPageToken string
	// Array of registry webhooks.
	// Wire name: 'webhooks'
	Webhooks []RegistryWebhook

	ForceSendFields []string `tf:"-"`
}

func (st *ListRegistryWebhooks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRegistryWebhooksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRegistryWebhooksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRegistryWebhooks) MarshalJSON() ([]byte, error) {
	pb, err := listRegistryWebhooksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List transition requests
type ListTransitionRequestsRequest struct {
	// Name of the model.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Version of the model.
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st *ListTransitionRequestsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTransitionRequestsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTransitionRequestsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTransitionRequestsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listTransitionRequestsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListTransitionRequestsResponse struct {
	// Array of open transition requests.
	// Wire name: 'requests'
	Requests []Activity
}

func (st *ListTransitionRequestsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTransitionRequestsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTransitionRequestsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTransitionRequestsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listTransitionRequestsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List registry webhooks
type ListWebhooksRequest struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	// Wire name: 'events'
	Events []RegistryWebhookEvent `tf:"-"`
	// If not specified, all webhooks associated with the specified events are
	// listed, regardless of their associated model.
	// Wire name: 'model_name'
	ModelName string `tf:"-"`
	// Token indicating the page of artifact results to fetch
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListWebhooksRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listWebhooksRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listWebhooksRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListWebhooksRequest) MarshalJSON() ([]byte, error) {
	pb, err := listWebhooksRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	// Wire name: 'metrics'
	Metrics []Metric
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	// Wire name: 'params'
	Params []Param
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	// Wire name: 'tags'
	Tags []RunTag

	ForceSendFields []string `tf:"-"`
}

func (st *LogBatch) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logBatchPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logBatchFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogBatch) MarshalJSON() ([]byte, error) {
	pb, err := logBatchToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogBatchResponse struct {
}

func (st *LogBatchResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logBatchResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logBatchResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogBatchResponse) MarshalJSON() ([]byte, error) {
	pb, err := logBatchResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogInputs struct {
	// Dataset inputs
	// Wire name: 'datasets'
	Datasets []DatasetInput
	// Model inputs
	// Wire name: 'models'
	Models []ModelInput
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string
}

func (st *LogInputs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logInputsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logInputsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogInputs) MarshalJSON() ([]byte, error) {
	pb, err := logInputsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogInputsResponse struct {
}

func (st *LogInputsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logInputsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logInputsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogInputsResponse) MarshalJSON() ([]byte, error) {
	pb, err := logInputsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogLoggedModelParamsRequest struct {
	// The ID of the logged model to log params for.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// Parameters to attach to the model.
	// Wire name: 'params'
	Params []LoggedModelParameter
}

func (st *LogLoggedModelParamsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logLoggedModelParamsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logLoggedModelParamsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogLoggedModelParamsRequest) MarshalJSON() ([]byte, error) {
	pb, err := logLoggedModelParamsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogLoggedModelParamsRequestResponse struct {
}

func (st *LogLoggedModelParamsRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logLoggedModelParamsRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logLoggedModelParamsRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogLoggedModelParamsRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := logLoggedModelParamsRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogMetric struct {
	// Dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	// Wire name: 'dataset_digest'
	DatasetDigest string
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	// Wire name: 'dataset_name'
	DatasetName string
	// Name of the metric.
	// Wire name: 'key'
	Key string
	// ID of the logged model associated with the metric, if applicable
	// Wire name: 'model_id'
	ModelId string
	// ID of the run under which to log the metric. Must be provided.
	// Wire name: 'run_id'
	RunId string
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// Step at which to log the metric
	// Wire name: 'step'
	Step int64
	// Unix timestamp in milliseconds at the time metric was logged.
	// Wire name: 'timestamp'
	Timestamp int64
	// Double value of the metric being logged.
	// Wire name: 'value'
	Value float64

	ForceSendFields []string `tf:"-"`
}

func (st *LogMetric) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logMetricPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logMetricFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogMetric) MarshalJSON() ([]byte, error) {
	pb, err := logMetricToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogMetricResponse struct {
}

func (st *LogMetricResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logMetricResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logMetricResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogMetricResponse) MarshalJSON() ([]byte, error) {
	pb, err := logMetricResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogModel struct {
	// MLmodel file in json format.
	// Wire name: 'model_json'
	ModelJson string
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string

	ForceSendFields []string `tf:"-"`
}

func (st *LogModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogModel) MarshalJSON() ([]byte, error) {
	pb, err := logModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogModelResponse struct {
}

func (st *LogModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := logModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogOutputsRequest struct {
	// The model outputs from the Run.
	// Wire name: 'models'
	Models []ModelOutput
	// The ID of the Run from which to log outputs.
	// Wire name: 'run_id'
	RunId string
}

func (st *LogOutputsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logOutputsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logOutputsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogOutputsRequest) MarshalJSON() ([]byte, error) {
	pb, err := logOutputsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogOutputsResponse struct {
}

func (st *LogOutputsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logOutputsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logOutputsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogOutputsResponse) MarshalJSON() ([]byte, error) {
	pb, err := logOutputsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	// Wire name: 'key'
	Key string
	// ID of the run under which to log the param. Must be provided.
	// Wire name: 'run_id'
	RunId string
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// String value of the param being logged. Maximum size is 500 bytes.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *LogParam) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logParamPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logParamFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogParam) MarshalJSON() ([]byte, error) {
	pb, err := logParamToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogParamResponse struct {
}

func (st *LogParamResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logParamResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logParamResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogParamResponse) MarshalJSON() ([]byte, error) {
	pb, err := logParamResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A logged model message includes logged model attributes, tags, registration
// info, params, and linked run metrics.
type LoggedModel struct {
	// The params and metrics attached to the logged model.
	// Wire name: 'data'
	Data *LoggedModelData
	// The logged model attributes such as model ID, status, tags, etc.
	// Wire name: 'info'
	Info *LoggedModelInfo
}

func (st *LoggedModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &loggedModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := loggedModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LoggedModel) MarshalJSON() ([]byte, error) {
	pb, err := loggedModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A LoggedModelData message includes logged model params and linked metrics.
type LoggedModelData struct {
	// Performance metrics linked to the model.
	// Wire name: 'metrics'
	Metrics []Metric
	// Immutable string key-value pairs of the model.
	// Wire name: 'params'
	Params []LoggedModelParameter
}

func (st *LoggedModelData) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &loggedModelDataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := loggedModelDataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LoggedModelData) MarshalJSON() ([]byte, error) {
	pb, err := loggedModelDataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A LoggedModelInfo includes logged model attributes, tags, and registration
// info.
type LoggedModelInfo struct {
	// The URI of the directory where model artifacts are stored.
	// Wire name: 'artifact_uri'
	ArtifactUri string
	// The timestamp when the model was created in milliseconds since the UNIX
	// epoch.
	// Wire name: 'creation_timestamp_ms'
	CreationTimestampMs int64
	// The ID of the user or principal that created the model.
	// Wire name: 'creator_id'
	CreatorId int64
	// The ID of the experiment that owns the model.
	// Wire name: 'experiment_id'
	ExperimentId string
	// The timestamp when the model was last updated in milliseconds since the
	// UNIX epoch.
	// Wire name: 'last_updated_timestamp_ms'
	LastUpdatedTimestampMs int64
	// The unique identifier for the logged model.
	// Wire name: 'model_id'
	ModelId string
	// The type of model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	// Wire name: 'model_type'
	ModelType string
	// The name of the model.
	// Wire name: 'name'
	Name string
	// The ID of the run that created the model.
	// Wire name: 'source_run_id'
	SourceRunId string
	// The status of whether or not the model is ready for use.
	// Wire name: 'status'
	Status LoggedModelStatus
	// Details on the current model status.
	// Wire name: 'status_message'
	StatusMessage string
	// Mutable string key-value pairs set on the model.
	// Wire name: 'tags'
	Tags []LoggedModelTag

	ForceSendFields []string `tf:"-"`
}

func (st *LoggedModelInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &loggedModelInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := loggedModelInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LoggedModelInfo) MarshalJSON() ([]byte, error) {
	pb, err := loggedModelInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Parameter associated with a LoggedModel.
type LoggedModelParameter struct {
	// The key identifying this param.
	// Wire name: 'key'
	Key string
	// The value of this param.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *LoggedModelParameter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &loggedModelParameterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := loggedModelParameterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LoggedModelParameter) MarshalJSON() ([]byte, error) {
	pb, err := loggedModelParameterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A LoggedModelStatus enum value represents the status of a logged model.
type LoggedModelStatus string
type loggedModelStatusPb string

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

// Type always returns LoggedModelStatus to satisfy [pflag.Value] interface
func (f *LoggedModelStatus) Type() string {
	return "LoggedModelStatus"
}

func loggedModelStatusToPb(st *LoggedModelStatus) (*loggedModelStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := loggedModelStatusPb(*st)
	return &pb, nil
}

func loggedModelStatusFromPb(pb *loggedModelStatusPb) (*LoggedModelStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := LoggedModelStatus(*pb)
	return &st, nil
}

// Tag for a LoggedModel.
type LoggedModelTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *LoggedModelTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &loggedModelTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := loggedModelTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LoggedModelTag) MarshalJSON() ([]byte, error) {
	pb, err := loggedModelTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Metric associated with a run, represented as a key-value pair.
type Metric struct {
	// The dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	// Wire name: 'dataset_digest'
	DatasetDigest string
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	// Wire name: 'dataset_name'
	DatasetName string
	// The key identifying the metric.
	// Wire name: 'key'
	Key string
	// The ID of the logged model or registered model version associated with
	// the metric, if applicable.
	// Wire name: 'model_id'
	ModelId string
	// The ID of the run containing the metric.
	// Wire name: 'run_id'
	RunId string
	// The step at which the metric was logged.
	// Wire name: 'step'
	Step int64
	// The timestamp at which the metric was recorded.
	// Wire name: 'timestamp'
	Timestamp int64
	// The value of the metric.
	// Wire name: 'value'
	Value float64

	ForceSendFields []string `tf:"-"`
}

func (st *Metric) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &metricPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := metricFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Metric) MarshalJSON() ([]byte, error) {
	pb, err := metricToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Model struct {
	// Timestamp recorded when this `registered_model` was created.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Description of this `registered_model`.
	// Wire name: 'description'
	Description string
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	// Wire name: 'latest_versions'
	LatestVersions []ModelVersion
	// Unique name for the model.
	// Wire name: 'name'
	Name string
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	// Wire name: 'tags'
	Tags []ModelTag
	// User that created this `registered_model`
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func (st *Model) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Model) MarshalJSON() ([]byte, error) {
	pb, err := modelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ModelDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// User-specified description for the object.
	// Wire name: 'description'
	Description string
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Array of model versions, each the latest version for its stage.
	// Wire name: 'latest_versions'
	LatestVersions []ModelVersion
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel
	// Array of tags associated with the model.
	// Wire name: 'tags'
	Tags []ModelTag
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func (st *ModelDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := modelDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput struct {
	// The unique identifier of the model.
	// Wire name: 'model_id'
	ModelId string
}

func (st *ModelInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelInput) MarshalJSON() ([]byte, error) {
	pb, err := modelInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents a LoggedModel output of a Run.
type ModelOutput struct {
	// The unique identifier of the model.
	// Wire name: 'model_id'
	ModelId string
	// The step at which the model was produced.
	// Wire name: 'step'
	Step int64
}

func (st *ModelOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelOutput) MarshalJSON() ([]byte, error) {
	pb, err := modelOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ModelTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *ModelTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelTag) MarshalJSON() ([]byte, error) {
	pb, err := modelTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ModelVersion struct {
	// Timestamp recorded when this `model_version` was created.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Current stage for this `model_version`.
	// Wire name: 'current_stage'
	CurrentStage string
	// Description of this `model_version`.
	// Wire name: 'description'
	Description string
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Unique name of the model
	// Wire name: 'name'
	Name string
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	// Wire name: 'run_id'
	RunId string
	// Run Link: Direct link to the run that generated this version
	// Wire name: 'run_link'
	RunLink string
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	// Wire name: 'source'
	Source string
	// Current status of `model_version`
	// Wire name: 'status'
	Status ModelVersionStatus
	// Details on current `status`, if it is pending or failed.
	// Wire name: 'status_message'
	StatusMessage string
	// Tags: Additional metadata key-value pairs for this `model_version`.
	// Wire name: 'tags'
	Tags []ModelVersionTag
	// User that created this `model_version`.
	// Wire name: 'user_id'
	UserId string
	// Model's version number.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func (st *ModelVersion) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelVersionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelVersionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelVersion) MarshalJSON() ([]byte, error) {
	pb, err := modelVersionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ModelVersionDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Stage of the model version. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'current_stage'
	CurrentStage Stage
	// User-specified description for the object.
	// Wire name: 'description'
	Description string
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	// Wire name: 'run_id'
	RunId string
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	// Wire name: 'run_link'
	RunLink string
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	// Wire name: 'source'
	Source string
	// The status of the model version. Valid values are: *
	// `PENDING_REGISTRATION`: Request to register a new model version is
	// pending as server performs background tasks.
	//
	// * `FAILED_REGISTRATION`: Request to register a new model version has
	// failed.
	//
	// * `READY`: Model version is ready for use.
	// Wire name: 'status'
	Status Status
	// Details on the current status, for example why registration failed.
	// Wire name: 'status_message'
	StatusMessage string
	// Array of tags that are associated with the model version.
	// Wire name: 'tags'
	Tags []ModelVersionTag
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func (st *ModelVersionDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelVersionDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelVersionDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelVersionDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := modelVersionDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Current status of `model_version`
type ModelVersionStatus string
type modelVersionStatusPb string

const ModelVersionStatusFailedRegistration ModelVersionStatus = `FAILED_REGISTRATION`

const ModelVersionStatusPendingRegistration ModelVersionStatus = `PENDING_REGISTRATION`

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

// Type always returns ModelVersionStatus to satisfy [pflag.Value] interface
func (f *ModelVersionStatus) Type() string {
	return "ModelVersionStatus"
}

func modelVersionStatusToPb(st *ModelVersionStatus) (*modelVersionStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := modelVersionStatusPb(*st)
	return &pb, nil
}

func modelVersionStatusFromPb(pb *modelVersionStatusPb) (*ModelVersionStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ModelVersionStatus(*pb)
	return &st, nil
}

type ModelVersionTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *ModelVersionTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelVersionTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelVersionTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelVersionTag) MarshalJSON() ([]byte, error) {
	pb, err := modelVersionTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Param associated with a run.
type Param struct {
	// Key identifying this param.
	// Wire name: 'key'
	Key string
	// Value associated with this param.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *Param) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &paramPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := paramFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Param) MarshalJSON() ([]byte, error) {
	pb, err := paramToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
type PermissionLevel string
type permissionLevelPb string

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
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

func permissionLevelToPb(st *PermissionLevel) (*permissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := permissionLevelPb(*st)
	return &pb, nil
}

func permissionLevelFromPb(pb *permissionLevelPb) (*PermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PermissionLevel(*pb)
	return &st, nil
}

type RegisteredModelAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *RegisteredModelAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RegisteredModelAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []RegisteredModelPermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *RegisteredModelAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RegisteredModelPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *RegisteredModelPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelPermission) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Permission level
type RegisteredModelPermissionLevel string
type registeredModelPermissionLevelPb string

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

// Type always returns RegisteredModelPermissionLevel to satisfy [pflag.Value] interface
func (f *RegisteredModelPermissionLevel) Type() string {
	return "RegisteredModelPermissionLevel"
}

func registeredModelPermissionLevelToPb(st *RegisteredModelPermissionLevel) (*registeredModelPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := registeredModelPermissionLevelPb(*st)
	return &pb, nil
}

func registeredModelPermissionLevelFromPb(pb *registeredModelPermissionLevelPb) (*RegisteredModelPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegisteredModelPermissionLevel(*pb)
	return &st, nil
}

type RegisteredModelPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []RegisteredModelAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func (st *RegisteredModelPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelPermissions) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RegisteredModelPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *RegisteredModelPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RegisteredModelPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []RegisteredModelAccessControlRequest
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func (st *RegisteredModelPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RegistryWebhook struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string
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
	// Wire name: 'events'
	Events []RegistryWebhookEvent

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpecWithoutSecret
	// Webhook ID
	// Wire name: 'id'
	Id string

	// Wire name: 'job_spec'
	JobSpec *JobSpecWithoutSecret
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Name of the model whose events would trigger this webhook.
	// Wire name: 'model_name'
	ModelName string
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	// Wire name: 'status'
	Status RegistryWebhookStatus

	ForceSendFields []string `tf:"-"`
}

func (st *RegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := registryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RegistryWebhookEvent string
type registryWebhookEventPb string

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

// Type always returns RegistryWebhookEvent to satisfy [pflag.Value] interface
func (f *RegistryWebhookEvent) Type() string {
	return "RegistryWebhookEvent"
}

func registryWebhookEventToPb(st *RegistryWebhookEvent) (*registryWebhookEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := registryWebhookEventPb(*st)
	return &pb, nil
}

func registryWebhookEventFromPb(pb *registryWebhookEventPb) (*RegistryWebhookEvent, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegistryWebhookEvent(*pb)
	return &st, nil
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
type registryWebhookStatusPb string

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

// Type always returns RegistryWebhookStatus to satisfy [pflag.Value] interface
func (f *RegistryWebhookStatus) Type() string {
	return "RegistryWebhookStatus"
}

func registryWebhookStatusToPb(st *RegistryWebhookStatus) (*registryWebhookStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := registryWebhookStatusPb(*st)
	return &pb, nil
}

func registryWebhookStatusFromPb(pb *registryWebhookStatusPb) (*RegistryWebhookStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegistryWebhookStatus(*pb)
	return &st, nil
}

type RejectTransitionRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage Stage
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func (st *RejectTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rejectTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rejectTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RejectTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := rejectTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RejectTransitionRequestResponse struct {
	// Activity recorded for the action.
	// Wire name: 'activity'
	Activity *Activity
}

func (st *RejectTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rejectTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rejectTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RejectTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := rejectTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string
	// If provided, updates the name for this `registered_model`.
	// Wire name: 'new_name'
	NewName string

	ForceSendFields []string `tf:"-"`
}

func (st *RenameModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &renameModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := renameModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RenameModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := renameModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RenameModelResponse struct {

	// Wire name: 'registered_model'
	RegisteredModel *Model
}

func (st *RenameModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &renameModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := renameModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RenameModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := renameModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
}

func (st *RestoreExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreExperiment) MarshalJSON() ([]byte, error) {
	pb, err := restoreExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestoreExperimentResponse struct {
}

func (st *RestoreExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := restoreExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestoreRun struct {
	// ID of the run to restore.
	// Wire name: 'run_id'
	RunId string
}

func (st *RestoreRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreRun) MarshalJSON() ([]byte, error) {
	pb, err := restoreRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestoreRunResponse struct {
}

func (st *RestoreRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := restoreRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	// Wire name: 'experiment_id'
	ExperimentId string
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	// Wire name: 'max_runs'
	MaxRuns int
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	// Wire name: 'min_timestamp_millis'
	MinTimestampMillis int64

	ForceSendFields []string `tf:"-"`
}

func (st *RestoreRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreRuns) MarshalJSON() ([]byte, error) {
	pb, err := restoreRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestoreRunsResponse struct {
	// The number of runs restored.
	// Wire name: 'runs_restored'
	RunsRestored int

	ForceSendFields []string `tf:"-"`
}

func (st *RestoreRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := restoreRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A single run.
type Run struct {
	// Run data.
	// Wire name: 'data'
	Data *RunData
	// Run metadata.
	// Wire name: 'info'
	Info *RunInfo
	// Run inputs.
	// Wire name: 'inputs'
	Inputs *RunInputs
}

func (st *Run) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Run) MarshalJSON() ([]byte, error) {
	pb, err := runToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Run data (metrics, params, and tags).
type RunData struct {
	// Run metrics.
	// Wire name: 'metrics'
	Metrics []Metric
	// Run parameters.
	// Wire name: 'params'
	Params []Param
	// Additional metadata key-value pairs.
	// Wire name: 'tags'
	Tags []RunTag
}

func (st *RunData) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runDataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runDataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunData) MarshalJSON() ([]byte, error) {
	pb, err := runDataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Metadata of a single run.
type RunInfo struct {
	// URI of the directory where artifacts should be uploaded. This can be a
	// local path (starting with "/"), or a distributed file system (DFS) path,
	// like ``s3://bucket/directory`` or ``dbfs:/my/directory``. If not set, the
	// local ``./mlruns`` directory is chosen.
	// Wire name: 'artifact_uri'
	ArtifactUri string
	// Unix timestamp of when the run ended in milliseconds.
	// Wire name: 'end_time'
	EndTime int64
	// The experiment ID.
	// Wire name: 'experiment_id'
	ExperimentId string
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	// Wire name: 'lifecycle_stage'
	LifecycleStage string
	// Unique identifier for the run.
	// Wire name: 'run_id'
	RunId string
	// The name of the run.
	// Wire name: 'run_name'
	RunName string
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// Unix timestamp of when the run started in milliseconds.
	// Wire name: 'start_time'
	StartTime int64
	// Current status of the run.
	// Wire name: 'status'
	Status RunInfoStatus
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func (st *RunInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunInfo) MarshalJSON() ([]byte, error) {
	pb, err := runInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Status of a run.
type RunInfoStatus string
type runInfoStatusPb string

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

// Type always returns RunInfoStatus to satisfy [pflag.Value] interface
func (f *RunInfoStatus) Type() string {
	return "RunInfoStatus"
}

func runInfoStatusToPb(st *RunInfoStatus) (*runInfoStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runInfoStatusPb(*st)
	return &pb, nil
}

func runInfoStatusFromPb(pb *runInfoStatusPb) (*RunInfoStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunInfoStatus(*pb)
	return &st, nil
}

// Run inputs.
type RunInputs struct {
	// Run metrics.
	// Wire name: 'dataset_inputs'
	DatasetInputs []DatasetInput
	// **NOTE**: Experimental: This API field may change or be removed in a
	// future release without warning.
	//
	// Model inputs to the Run.
	// Wire name: 'model_inputs'
	ModelInputs []ModelInput
}

func (st *RunInputs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runInputsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runInputsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunInputs) MarshalJSON() ([]byte, error) {
	pb, err := runInputsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Tag for a run.
type RunTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *RunTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunTag) MarshalJSON() ([]byte, error) {
	pb, err := runTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchExperiments struct {
	// String representing a SQL filter condition (e.g. "name ILIKE
	// 'my-experiment%'")
	// Wire name: 'filter'
	Filter string
	// Maximum number of experiments desired. Max threshold is 3000.
	// Wire name: 'max_results'
	MaxResults int64
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	// Wire name: 'order_by'
	OrderBy []string
	// Token indicating the page of experiments to fetch
	// Wire name: 'page_token'
	PageToken string
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	// Wire name: 'view_type'
	ViewType ViewType

	ForceSendFields []string `tf:"-"`
}

func (st *SearchExperiments) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchExperimentsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchExperimentsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchExperiments) MarshalJSON() ([]byte, error) {
	pb, err := searchExperimentsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	// Wire name: 'experiments'
	Experiments []Experiment
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *SearchExperimentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchExperimentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchExperimentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchExperimentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchExperimentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchLoggedModelsDataset struct {
	// The digest of the dataset.
	// Wire name: 'dataset_digest'
	DatasetDigest string
	// The name of the dataset.
	// Wire name: 'dataset_name'
	DatasetName string

	ForceSendFields []string `tf:"-"`
}

func (st *SearchLoggedModelsDataset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchLoggedModelsDatasetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchLoggedModelsDatasetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchLoggedModelsDataset) MarshalJSON() ([]byte, error) {
	pb, err := searchLoggedModelsDatasetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchLoggedModelsOrderBy struct {
	// Whether the search results order is ascending or not.
	// Wire name: 'ascending'
	Ascending bool
	// If ``field_name`` refers to a metric, this field specifies the digest of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name and digest will be considered for ordering. This
	// field may only be set if ``dataset_name`` is also set.
	// Wire name: 'dataset_digest'
	DatasetDigest string
	// If ``field_name`` refers to a metric, this field specifies the name of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name will be considered for ordering. This field may
	// only be set if ``field_name`` refers to a metric.
	// Wire name: 'dataset_name'
	DatasetName string
	// The name of the field to order by, e.g. "metrics.accuracy".
	// Wire name: 'field_name'
	FieldName string

	ForceSendFields []string `tf:"-"`
}

func (st *SearchLoggedModelsOrderBy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchLoggedModelsOrderByPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchLoggedModelsOrderByFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchLoggedModelsOrderBy) MarshalJSON() ([]byte, error) {
	pb, err := searchLoggedModelsOrderByToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchLoggedModelsRequest struct {
	// List of datasets on which to apply the metrics filter clauses. For
	// example, a filter with `metrics.accuracy > 0.9` and dataset info with
	// name "test_dataset" means we will return all logged models with accuracy
	// > 0.9 on the test_dataset. Metric values from ANY dataset matching the
	// criteria are considered. If no datasets are specified, then metrics
	// across all datasets are considered in the filter.
	// Wire name: 'datasets'
	Datasets []SearchLoggedModelsDataset
	// The IDs of the experiments in which to search for logged models.
	// Wire name: 'experiment_ids'
	ExperimentIds []string
	// A filter expression over logged model info and data that allows returning
	// a subset of logged models. The syntax is a subset of SQL that supports
	// AND'ing together binary operations.
	//
	// Example: ``params.alpha < 0.3 AND metrics.accuracy > 0.9``.
	// Wire name: 'filter'
	Filter string
	// The maximum number of Logged Models to return. The maximum limit is 50.
	// Wire name: 'max_results'
	MaxResults int
	// The list of columns for ordering the results, with additional fields for
	// sorting criteria.
	// Wire name: 'order_by'
	OrderBy []SearchLoggedModelsOrderBy
	// The token indicating the page of logged models to fetch.
	// Wire name: 'page_token'
	PageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *SearchLoggedModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchLoggedModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchLoggedModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchLoggedModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := searchLoggedModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchLoggedModelsResponse struct {
	// Logged models that match the search criteria.
	// Wire name: 'models'
	Models []LoggedModel
	// The token that can be used to retrieve the next page of logged models.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *SearchLoggedModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchLoggedModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchLoggedModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchLoggedModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchLoggedModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Searches model versions
type SearchModelVersionsRequest struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Maximum number of models desired. Max threshold is 10K.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	// Wire name: 'order_by'
	OrderBy []string `tf:"-"`
	// Pagination token to go to next page based on previous search query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *SearchModelVersionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchModelVersionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchModelVersionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchModelVersionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := searchModelVersionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	// Wire name: 'model_versions'
	ModelVersions []ModelVersion
	// Pagination token to request next page of models for the same search
	// query.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *SearchModelVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchModelVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchModelVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchModelVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchModelVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Search models
type SearchModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	// Wire name: 'order_by'
	OrderBy []string `tf:"-"`
	// Pagination token to go to the next page based on a previous search query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *SearchModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := searchModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	// Wire name: 'next_page_token'
	NextPageToken string
	// Registered Models that match the search criteria.
	// Wire name: 'registered_models'
	RegisteredModels []Model

	ForceSendFields []string `tf:"-"`
}

func (st *SearchModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	// Wire name: 'experiment_ids'
	ExperimentIds []string
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
	// Wire name: 'filter'
	Filter string
	// Maximum number of runs desired. Max threshold is 50000
	// Wire name: 'max_results'
	MaxResults int
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional `"DESC"` or `"ASC"` annotation, where `"ASC"`
	// is the default. Example: `["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"]`. Tiebreaks are done by start_time `DESC` followed by
	// `run_id` for runs with the same start time (and this is the default
	// ordering criterion if order_by is not provided).
	// Wire name: 'order_by'
	OrderBy []string
	// Token for the current page of runs.
	// Wire name: 'page_token'
	PageToken string
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	// Wire name: 'run_view_type'
	RunViewType ViewType

	ForceSendFields []string `tf:"-"`
}

func (st *SearchRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchRuns) MarshalJSON() ([]byte, error) {
	pb, err := searchRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	// Wire name: 'next_page_token'
	NextPageToken string
	// Runs that match the search criteria.
	// Wire name: 'runs'
	Runs []Run

	ForceSendFields []string `tf:"-"`
}

func (st *SearchRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	// Wire name: 'experiment_id'
	ExperimentId string
	// Name of the tag. Keys up to 250 bytes in size are supported.
	// Wire name: 'key'
	Key string
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	// Wire name: 'value'
	Value string
}

func (st *SetExperimentTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setExperimentTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setExperimentTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetExperimentTag) MarshalJSON() ([]byte, error) {
	pb, err := setExperimentTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetExperimentTagResponse struct {
}

func (st *SetExperimentTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setExperimentTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setExperimentTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetExperimentTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := setExperimentTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetLoggedModelTagsRequest struct {
	// The ID of the logged model to set the tags on.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// The tags to set on the logged model.
	// Wire name: 'tags'
	Tags []LoggedModelTag
}

func (st *SetLoggedModelTagsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setLoggedModelTagsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setLoggedModelTagsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetLoggedModelTagsRequest) MarshalJSON() ([]byte, error) {
	pb, err := setLoggedModelTagsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetLoggedModelTagsResponse struct {
}

func (st *SetLoggedModelTagsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setLoggedModelTagsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setLoggedModelTagsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetLoggedModelTagsResponse) MarshalJSON() ([]byte, error) {
	pb, err := setLoggedModelTagsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	// Wire name: 'key'
	Key string
	// Unique name of the model.
	// Wire name: 'name'
	Name string
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	// Wire name: 'value'
	Value string
}

func (st *SetModelTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setModelTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setModelTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetModelTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := setModelTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetModelTagResponse struct {
}

func (st *SetModelTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setModelTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setModelTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetModelTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := setModelTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetModelVersionTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	// Wire name: 'key'
	Key string
	// Unique name of the model.
	// Wire name: 'name'
	Name string
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	// Wire name: 'value'
	Value string
	// Model version number.
	// Wire name: 'version'
	Version string
}

func (st *SetModelVersionTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setModelVersionTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setModelVersionTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetModelVersionTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := setModelVersionTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetModelVersionTagResponse struct {
}

func (st *SetModelVersionTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setModelVersionTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setModelVersionTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetModelVersionTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := setModelVersionTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetTag struct {
	// Name of the tag. Keys up to 250 bytes in size are supported.
	// Wire name: 'key'
	Key string
	// ID of the run under which to log the tag. Must be provided.
	// Wire name: 'run_id'
	RunId string
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *SetTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetTag) MarshalJSON() ([]byte, error) {
	pb, err := setTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetTagResponse struct {
}

func (st *SetTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := setTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Stage of the model version. Valid values are:
//
// * `None`: The initial stage of a model version.
//
// * `Staging`: Staging or pre-production stage.
//
// * `Production`: Production stage.
//
// * `Archived`: Archived stage.
type Stage string
type stagePb string

// Archived stage.
const StageArchived Stage = `Archived`

// The initial stage of a model version.
const StageNone Stage = `None`

// Production stage.
const StageProduction Stage = `Production`

// Staging or pre-production stage.
const StageStaging Stage = `Staging`

// String representation for [fmt.Print]
func (f *Stage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Stage) Set(v string) error {
	switch v {
	case `Archived`, `None`, `Production`, `Staging`:
		*f = Stage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Archived", "None", "Production", "Staging"`, v)
	}
}

// Type always returns Stage to satisfy [pflag.Value] interface
func (f *Stage) Type() string {
	return "Stage"
}

func stageToPb(st *Stage) (*stagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := stagePb(*st)
	return &pb, nil
}

func stageFromPb(pb *stagePb) (*Stage, error) {
	if pb == nil {
		return nil, nil
	}
	st := Stage(*pb)
	return &st, nil
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type Status string
type statusPb string

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

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

func statusToPb(st *Status) (*statusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := statusPb(*st)
	return &pb, nil
}

func statusFromPb(pb *statusPb) (*Status, error) {
	if pb == nil {
		return nil, nil
	}
	st := Status(*pb)
	return &st, nil
}

// Test webhook response object.
type TestRegistryWebhook struct {
	// Body of the response from the webhook URL
	// Wire name: 'body'
	Body string
	// Status code returned by the webhook URL
	// Wire name: 'status_code'
	StatusCode int

	ForceSendFields []string `tf:"-"`
}

func (st *TestRegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &testRegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := testRegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TestRegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := testRegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	// Wire name: 'event'
	Event RegistryWebhookEvent
	// Webhook ID
	// Wire name: 'id'
	Id string
}

func (st *TestRegistryWebhookRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &testRegistryWebhookRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := testRegistryWebhookRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TestRegistryWebhookRequest) MarshalJSON() ([]byte, error) {
	pb, err := testRegistryWebhookRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TestRegistryWebhookResponse struct {
	// Test webhook response object.
	// Wire name: 'webhook'
	Webhook *TestRegistryWebhook
}

func (st *TestRegistryWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &testRegistryWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := testRegistryWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TestRegistryWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := testRegistryWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	// Wire name: 'archive_existing_versions'
	ArchiveExistingVersions bool
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage Stage
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func (st *TransitionModelVersionStageDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &transitionModelVersionStageDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := transitionModelVersionStageDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TransitionModelVersionStageDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := transitionModelVersionStageDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	// Wire name: 'available_actions'
	AvailableActions []ActivityAction
	// User-provided comment associated with the transition request.
	// Wire name: 'comment'
	Comment string
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
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
	// Wire name: 'to_stage'
	ToStage Stage
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func (st *TransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &transitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := transitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := transitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TransitionStageResponse struct {

	// Wire name: 'model_version'
	ModelVersion *ModelVersionDatabricks
}

func (st *TransitionStageResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &transitionStageResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := transitionStageResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TransitionStageResponse) MarshalJSON() ([]byte, error) {
	pb, err := transitionStageResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateComment struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Unique identifier of an activity
	// Wire name: 'id'
	Id string
}

func (st *UpdateComment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCommentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCommentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateComment) MarshalJSON() ([]byte, error) {
	pb, err := updateCommentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateCommentResponse struct {
	// Comment details.
	// Wire name: 'comment'
	Comment *CommentObject
}

func (st *UpdateCommentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCommentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCommentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCommentResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateCommentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	// Wire name: 'new_name'
	NewName string

	ForceSendFields []string `tf:"-"`
}

func (st *UpdateExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExperiment) MarshalJSON() ([]byte, error) {
	pb, err := updateExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateExperimentResponse struct {
}

func (st *UpdateExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	// Wire name: 'description'
	Description string
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func (st *UpdateModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateModelResponse struct {
}

func (st *UpdateModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	// Wire name: 'description'
	Description string
	// Name of the registered model
	// Wire name: 'name'
	Name string
	// Model version number
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func (st *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateModelVersionResponse struct {
}

func (st *UpdateModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string
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
	// Wire name: 'events'
	Events []RegistryWebhookEvent

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpec
	// Webhook ID
	// Wire name: 'id'
	Id string

	// Wire name: 'job_spec'
	JobSpec *JobSpec
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	// Wire name: 'status'
	Status RegistryWebhookStatus

	ForceSendFields []string `tf:"-"`
}

func (st *UpdateRegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := updateRegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	// Wire name: 'end_time'
	EndTime int64
	// ID of the run to update. Must be provided.
	// Wire name: 'run_id'
	RunId string
	// Updated name of the run.
	// Wire name: 'run_name'
	RunName string
	// [Deprecated, use `run_id` instead] ID of the run to update. This field
	// will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// Updated status of the run.
	// Wire name: 'status'
	Status UpdateRunStatus

	ForceSendFields []string `tf:"-"`
}

func (st *UpdateRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRun) MarshalJSON() ([]byte, error) {
	pb, err := updateRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	// Wire name: 'run_info'
	RunInfo *RunInfo
}

func (st *UpdateRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Status of a run.
type UpdateRunStatus string
type updateRunStatusPb string

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

// Type always returns UpdateRunStatus to satisfy [pflag.Value] interface
func (f *UpdateRunStatus) Type() string {
	return "UpdateRunStatus"
}

func updateRunStatusToPb(st *UpdateRunStatus) (*updateRunStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := updateRunStatusPb(*st)
	return &pb, nil
}

func updateRunStatusFromPb(pb *updateRunStatusPb) (*UpdateRunStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateRunStatus(*pb)
	return &st, nil
}

type UpdateWebhookResponse struct {
}

func (st *UpdateWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Qualifier for the view type.
type ViewType string
type viewTypePb string

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

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
}

func viewTypeToPb(st *ViewType) (*viewTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := viewTypePb(*st)
	return &pb, nil
}

func viewTypeFromPb(pb *viewTypePb) (*ViewType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ViewType(*pb)
	return &st, nil
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
