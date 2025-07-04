// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type Activity struct {

	// Wire name: 'activity_type'
	ActivityType ActivityType `json:"activity_type,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
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
	// Wire name: 'from_stage'
	FromStage string `json:"from_stage,omitempty"`
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	// Wire name: 'system_comment'
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
	// Wire name: 'to_stage'
	ToStage string `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Activity) EncodeValues(key string, v *url.Values) error {
	pb, err := activityToPb(&st)
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
	// Wire name: 'archive_existing_versions'
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	// Wire name: 'name'
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
	// Wire name: 'stage'
	Stage string `json:"stage"`
	// Version of the model.
	// Wire name: 'version'
	Version string `json:"version"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ApproveTransitionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := approveTransitionRequestToPb(&st)
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
	// New activity generated as a result of this operation.
	// Wire name: 'activity'
	Activity *Activity `json:"activity,omitempty"`
}

func (st ApproveTransitionRequestResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := approveTransitionRequestResponseToPb(&st)
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
	// Wire name: 'available_actions'
	AvailableActions []CommentActivityAction `json:"available_actions,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CommentObject) EncodeValues(key string, v *url.Values) error {
	pb, err := commentObjectToPb(&st)
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

// Details required to create a comment on a model version.
type CreateComment struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string `json:"comment"`
	// Name of the model.
	// Wire name: 'name'
	Name string `json:"name"`
	// Version of the model.
	// Wire name: 'version'
	Version string `json:"version"`
}

func (st CreateComment) EncodeValues(key string, v *url.Values) error {
	pb, err := createCommentToPb(&st)
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
	// New comment object
	// Wire name: 'comment'
	Comment *CommentObject `json:"comment,omitempty"`
}

func (st CreateCommentResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createCommentResponseToPb(&st)
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
	ArtifactLocation string `json:"artifact_location,omitempty"`
	// Experiment name.
	// Wire name: 'name'
	Name string `json:"name"`
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	// Wire name: 'tags'
	Tags []ExperimentTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateExperiment) EncodeValues(key string, v *url.Values) error {
	pb, err := createExperimentToPb(&st)
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
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateExperimentResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createExperimentResponseToPb(&st)
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

type CreateFeatureTagRequest struct {
	FeatureName string `json:"-" tf:"-"`

	// Wire name: 'feature_tag'
	FeatureTag FeatureTag `json:"feature_tag"`

	TableName string `json:"-" tf:"-"`
}

func (st CreateFeatureTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createFeatureTagRequestToPb(&st)
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

func (st *CreateFeatureTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createFeatureTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createFeatureTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateFeatureTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := createFeatureTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateForecastingExperimentRequest struct {
	// The column in the training table used to customize weights for each time
	// series.
	// Wire name: 'custom_weights_column'
	CustomWeightsColumn string `json:"custom_weights_column,omitempty"`
	// The path in the workspace to store the created experiment.
	// Wire name: 'experiment_path'
	ExperimentPath string `json:"experiment_path,omitempty"`
	// The time interval between consecutive rows in the time series data.
	// Possible values include: '1 second', '1 minute', '5 minutes', '10
	// minutes', '15 minutes', '30 minutes', 'Hourly', 'Daily', 'Weekly',
	// 'Monthly', 'Quarterly', 'Yearly'.
	// Wire name: 'forecast_granularity'
	ForecastGranularity string `json:"forecast_granularity"`
	// The number of time steps into the future to make predictions, calculated
	// as a multiple of forecast_granularity. This value represents how far
	// ahead the model should forecast.
	// Wire name: 'forecast_horizon'
	ForecastHorizon int64 `json:"forecast_horizon"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store future feature data
	// for predictions.
	// Wire name: 'future_feature_data_path'
	FutureFeatureDataPath string `json:"future_feature_data_path,omitempty"`
	// The region code(s) to automatically add holiday features. Currently
	// supports only one region.
	// Wire name: 'holiday_regions'
	HolidayRegions []string `json:"holiday_regions,omitempty"`
	// Specifies the list of feature columns to include in model training. These
	// columns must exist in the training data and be of type string, numerical,
	// or boolean. If not specified, no additional features will be included.
	// Note: Certain columns are automatically handled: - Automatically
	// excluded: split_column, target_column, custom_weights_column. -
	// Automatically included: time_column.
	// Wire name: 'include_features'
	IncludeFeatures []string `json:"include_features,omitempty"`
	// The maximum duration for the experiment in minutes. The experiment stops
	// automatically if it exceeds this limit.
	// Wire name: 'max_runtime'
	MaxRuntime int64 `json:"max_runtime,omitempty"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store predictions.
	// Wire name: 'prediction_data_path'
	PredictionDataPath string `json:"prediction_data_path,omitempty"`
	// The evaluation metric used to optimize the forecasting model.
	// Wire name: 'primary_metric'
	PrimaryMetric string `json:"primary_metric,omitempty"`
	// The fully qualified path of a Unity Catalog model, formatted as
	// catalog_name.schema_name.model_name, used to store the best model.
	// Wire name: 'register_to'
	RegisterTo string `json:"register_to,omitempty"`
	// // The column in the training table used for custom data splits. Values
	// must be 'train', 'validate', or 'test'.
	// Wire name: 'split_column'
	SplitColumn string `json:"split_column,omitempty"`
	// The column in the input training table used as the prediction target for
	// model training. The values in this column are used as the ground truth
	// for model training.
	// Wire name: 'target_column'
	TargetColumn string `json:"target_column"`
	// The column in the input training table that represents each row's
	// timestamp.
	// Wire name: 'time_column'
	TimeColumn string `json:"time_column"`
	// The column in the training table used to group the dataset for predicting
	// individual time series.
	// Wire name: 'timeseries_identifier_columns'
	TimeseriesIdentifierColumns []string `json:"timeseries_identifier_columns,omitempty"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used as training data for the
	// forecasting model.
	// Wire name: 'train_data_path'
	TrainDataPath string `json:"train_data_path"`
	// List of frameworks to include for model tuning. Possible values are
	// 'Prophet', 'ARIMA', 'DeepAR'. An empty list includes all supported
	// frameworks.
	// Wire name: 'training_frameworks'
	TrainingFrameworks []string `json:"training_frameworks,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateForecastingExperimentRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createForecastingExperimentRequestToPb(&st)
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
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateForecastingExperimentResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createForecastingExperimentResponseToPb(&st)
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
	ExperimentId string `json:"experiment_id"`
	// The type of the model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	// Wire name: 'model_type'
	ModelType string `json:"model_type,omitempty"`
	// The name of the model (optional). If not specified one will be generated.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Parameters attached to the model.
	// Wire name: 'params'
	Params []LoggedModelParameter `json:"params,omitempty"`
	// The ID of the run that created the model.
	// Wire name: 'source_run_id'
	SourceRunId string `json:"source_run_id,omitempty"`
	// Tags attached to the model.
	// Wire name: 'tags'
	Tags []LoggedModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateLoggedModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createLoggedModelRequestToPb(&st)
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
	Model *LoggedModel `json:"model,omitempty"`
}

func (st CreateLoggedModelResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createLoggedModelResponseToPb(&st)
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
	Description string `json:"description,omitempty"`
	// Register models under this name
	// Wire name: 'name'
	Name string `json:"name"`
	// Additional metadata for registered model.
	// Wire name: 'tags'
	Tags []ModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createModelRequestToPb(&st)
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
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

func (st CreateModelResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createModelResponseToPb(&st)
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
	Description string `json:"description,omitempty"`
	// Register model under this name
	// Wire name: 'name'
	Name string `json:"name"`
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	// Wire name: 'run_link'
	RunLink string `json:"run_link,omitempty"`
	// URI indicating the location of the model artifacts.
	// Wire name: 'source'
	Source string `json:"source"`
	// Additional metadata for model version.
	// Wire name: 'tags'
	Tags []ModelVersionTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateModelVersionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createModelVersionRequestToPb(&st)
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
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

func (st CreateModelVersionResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createModelVersionResponseToPb(&st)
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

type CreateOnlineStoreRequest struct {
	// Online store to create.
	// Wire name: 'online_store'
	OnlineStore OnlineStore `json:"online_store"`
}

func (st CreateOnlineStoreRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createOnlineStoreRequestToPb(&st)
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

func (st *CreateOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createOnlineStoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createOnlineStoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateOnlineStoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := createOnlineStoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to create a registry webhook.
type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	// Wire name: 'description'
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
	// Wire name: 'events'
	Events []RegistryWebhookEvent `json:"events"`
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`
	// ID of the job that the webhook runs.
	// Wire name: 'job_spec'
	JobSpec *JobSpec `json:"job_spec,omitempty"`
	// If model name is not specified, a registry-wide webhook is created that
	// listens for the specified events across all versions of all registered
	// models.
	// Wire name: 'model_name'
	ModelName string `json:"model_name,omitempty"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	// Wire name: 'status'
	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateRegistryWebhook) EncodeValues(key string, v *url.Values) error {
	pb, err := createRegistryWebhookToPb(&st)
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
	ExperimentId string `json:"experiment_id,omitempty"`
	// The name of the run.
	// Wire name: 'run_name'
	RunName string `json:"run_name,omitempty"`
	// Unix timestamp in milliseconds of when the run started.
	// Wire name: 'start_time'
	StartTime int64 `json:"start_time,omitempty"`
	// Additional metadata for run.
	// Wire name: 'tags'
	Tags []RunTag `json:"tags,omitempty"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateRun) EncodeValues(key string, v *url.Values) error {
	pb, err := createRunToPb(&st)
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
	Run *Run `json:"run,omitempty"`
}

func (st CreateRunResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createRunResponseToPb(&st)
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

// Details required to create a model version stage transition request.
type CreateTransitionRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	// Wire name: 'name'
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
	// Wire name: 'stage'
	Stage string `json:"stage"`
	// Version of the model.
	// Wire name: 'version'
	Version string `json:"version"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateTransitionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createTransitionRequestToPb(&st)
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
	// New activity generated for stage transition request.
	// Wire name: 'request'
	Request *TransitionRequest `json:"request,omitempty"`
}

func (st CreateTransitionRequestResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createTransitionRequestResponseToPb(&st)
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
	Webhook *RegistryWebhook `json:"webhook,omitempty"`
}

func (st CreateWebhookResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createWebhookResponseToPb(&st)
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
	Digest string `json:"digest"`
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	// Wire name: 'name'
	Name string `json:"name"`
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	// Wire name: 'profile'
	Profile string `json:"profile,omitempty"`
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	// Wire name: 'source'
	Source string `json:"source"`
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	// Wire name: 'source_type'
	SourceType string `json:"source_type"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Dataset) EncodeValues(key string, v *url.Values) error {
	pb, err := datasetToPb(&st)
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
	Dataset Dataset `json:"dataset"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	// Wire name: 'tags'
	Tags []InputTag `json:"tags,omitempty"`
}

func (st DatasetInput) EncodeValues(key string, v *url.Values) error {
	pb, err := datasetInputToPb(&st)
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

type DeleteCommentRequest struct {
	// Unique identifier of an activity
	Id string `json:"-" tf:"-"`
}

func (st DeleteCommentRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteCommentRequestToPb(&st)
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

type DeleteExperiment struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string `json:"experiment_id"`
}

func (st DeleteExperiment) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteExperimentToPb(&st)
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

type DeleteFeatureTagRequest struct {
	// The name of the feature within the feature table.
	FeatureName string `json:"-" tf:"-"`
	// The key of the tag to delete.
	Key string `json:"-" tf:"-"`
	// The name of the feature table.
	TableName string `json:"-" tf:"-"`
}

func (st DeleteFeatureTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteFeatureTagRequestToPb(&st)
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

func (st *DeleteFeatureTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteFeatureTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFeatureTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteFeatureTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteFeatureTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteFeatureTagResponse struct {
}

func (st DeleteFeatureTagResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteFeatureTagResponseToPb(&st)
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

func (st *DeleteFeatureTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteFeatureTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFeatureTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteFeatureTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteFeatureTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteLoggedModelRequest struct {
	// The ID of the logged model to delete.
	ModelId string `json:"-" tf:"-"`
}

func (st DeleteLoggedModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteLoggedModelRequestToPb(&st)
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

type DeleteLoggedModelTagRequest struct {
	// The ID of the logged model to delete the tag from.
	ModelId string `json:"-" tf:"-"`
	// The tag key.
	TagKey string `json:"-" tf:"-"`
}

func (st DeleteLoggedModelTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteLoggedModelTagRequestToPb(&st)
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

type DeleteModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"-" tf:"-"`
}

func (st DeleteModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteModelRequestToPb(&st)
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

type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"-" tf:"-"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"-" tf:"-"`
}

func (st DeleteModelTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteModelTagRequestToPb(&st)
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

type DeleteModelVersionRequest struct {
	// Name of the registered model
	Name string `json:"-" tf:"-"`
	// Model version number
	Version string `json:"-" tf:"-"`
}

func (st DeleteModelVersionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteModelVersionRequestToPb(&st)
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

type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"-" tf:"-"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"-" tf:"-"`
	// Model version number that the tag was logged under.
	Version string `json:"-" tf:"-"`
}

func (st DeleteModelVersionTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteModelVersionTagRequestToPb(&st)
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

type DeleteOnlineStoreRequest struct {
	// Name of the online store to delete.
	Name string `json:"-" tf:"-"`
}

func (st DeleteOnlineStoreRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteOnlineStoreRequestToPb(&st)
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

func (st *DeleteOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteOnlineStoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteOnlineStoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteOnlineStoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteOnlineStoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteOnlineStoreResponse struct {
}

func (st DeleteOnlineStoreResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteOnlineStoreResponseToPb(&st)
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

func (st *DeleteOnlineStoreResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteOnlineStoreResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteOnlineStoreResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteOnlineStoreResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteOnlineStoreResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteRun struct {
	// ID of the run to delete.
	// Wire name: 'run_id'
	RunId string `json:"run_id"`
}

func (st DeleteRun) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRunToPb(&st)
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

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	// Wire name: 'experiment_id'
	ExperimentId string `json:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	// Wire name: 'max_runs'
	MaxRuns int `json:"max_runs,omitempty"`
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	// Wire name: 'max_timestamp_millis'
	MaxTimestampMillis int64 `json:"max_timestamp_millis"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteRuns) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRunsToPb(&st)
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
	RunsDeleted int `json:"runs_deleted,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteRunsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRunsResponseToPb(&st)
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
	Key string `json:"key"`
	// ID of the run that the tag was logged under. Must be provided.
	// Wire name: 'run_id'
	RunId string `json:"run_id"`
}

func (st DeleteTag) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteTagToPb(&st)
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

type DeleteTransitionRequestRequest struct {
	// User-provided comment on the action.
	Comment string `json:"-" tf:"-"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	Creator string `json:"-" tf:"-"`
	// Name of the model.
	Name string `json:"-" tf:"-"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"-" tf:"-"`
	// Version of the model.
	Version string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteTransitionRequestRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteTransitionRequestRequestToPb(&st)
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
	// New activity generated as a result of this operation.
	// Wire name: 'activity'
	Activity *Activity `json:"activity,omitempty"`
}

func (st DeleteTransitionRequestResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteTransitionRequestResponseToPb(&st)
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

type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	Id string `json:"-" tf:"-"`
}

func (st DeleteWebhookRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteWebhookRequestToPb(&st)
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

// An experiment and its metadata.
type Experiment struct {
	// Location where artifacts for the experiment are stored.
	// Wire name: 'artifact_location'
	ArtifactLocation string `json:"artifact_location,omitempty"`
	// Creation time
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// Unique identifier for the experiment.
	// Wire name: 'experiment_id'
	ExperimentId string `json:"experiment_id,omitempty"`
	// Last update time
	// Wire name: 'last_update_time'
	LastUpdateTime int64 `json:"last_update_time,omitempty"`
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	// Wire name: 'lifecycle_stage'
	LifecycleStage string `json:"lifecycle_stage,omitempty"`
	// Human readable name that identifies the experiment.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Tags: Additional metadata key-value pairs.
	// Wire name: 'tags'
	Tags []ExperimentTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Experiment) EncodeValues(key string, v *url.Values) error {
	pb, err := experimentToPb(&st)
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
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExperimentAccessControlRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := experimentAccessControlRequestToPb(&st)
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
	AllPermissions []ExperimentPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExperimentAccessControlResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := experimentAccessControlResponseToPb(&st)
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
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExperimentPermission) EncodeValues(key string, v *url.Values) error {
	pb, err := experimentPermissionToPb(&st)
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

	// Wire name: 'access_control_list'
	AccessControlList []ExperimentAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExperimentPermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := experimentPermissionsToPb(&st)
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
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExperimentPermissionsDescription) EncodeValues(key string, v *url.Values) error {
	pb, err := experimentPermissionsDescriptionToPb(&st)
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
	AccessControlList []ExperimentAccessControlRequest `json:"access_control_list,omitempty"`
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" tf:"-"`
}

func (st ExperimentPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := experimentPermissionsRequestToPb(&st)
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
	Key string `json:"key,omitempty"`
	// The tag value.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExperimentTag) EncodeValues(key string, v *url.Values) error {
	pb, err := experimentTagToPb(&st)
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

// Feature for model version.
type Feature struct {
	// Feature name
	// Wire name: 'feature_name'
	FeatureName string `json:"feature_name,omitempty"`
	// Feature table id
	// Wire name: 'feature_table_id'
	FeatureTableId string `json:"feature_table_id,omitempty"`
	// Feature table name
	// Wire name: 'feature_table_name'
	FeatureTableName string `json:"feature_table_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Feature) EncodeValues(key string, v *url.Values) error {
	pb, err := featureToPb(&st)
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

func (st *Feature) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &featurePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := featureFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Feature) MarshalJSON() ([]byte, error) {
	pb, err := featureToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FeatureLineage struct {
	// List of feature specs that contain this feature.
	// Wire name: 'feature_specs'
	FeatureSpecs []FeatureLineageFeatureSpec `json:"feature_specs,omitempty"`
	// List of Unity Catalog models that were trained on this feature.
	// Wire name: 'models'
	Models []FeatureLineageModel `json:"models,omitempty"`
	// List of online features that use this feature as source.
	// Wire name: 'online_features'
	OnlineFeatures []FeatureLineageOnlineFeature `json:"online_features,omitempty"`
}

func (st FeatureLineage) EncodeValues(key string, v *url.Values) error {
	pb, err := featureLineageToPb(&st)
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

func (st *FeatureLineage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &featureLineagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := featureLineageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FeatureLineage) MarshalJSON() ([]byte, error) {
	pb, err := featureLineageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FeatureLineageFeatureSpec struct {
	// The full name of the feature spec in Unity Catalog.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FeatureLineageFeatureSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := featureLineageFeatureSpecToPb(&st)
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

func (st *FeatureLineageFeatureSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &featureLineageFeatureSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := featureLineageFeatureSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FeatureLineageFeatureSpec) MarshalJSON() ([]byte, error) {
	pb, err := featureLineageFeatureSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FeatureLineageModel struct {
	// The full name of the model in Unity Catalog.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The version of the model.
	// Wire name: 'version'
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FeatureLineageModel) EncodeValues(key string, v *url.Values) error {
	pb, err := featureLineageModelToPb(&st)
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

func (st *FeatureLineageModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &featureLineageModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := featureLineageModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FeatureLineageModel) MarshalJSON() ([]byte, error) {
	pb, err := featureLineageModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FeatureLineageOnlineFeature struct {
	// The name of the online feature (column name).
	// Wire name: 'feature_name'
	FeatureName string `json:"feature_name,omitempty"`
	// The full name of the online table in Unity Catalog.
	// Wire name: 'table_name'
	TableName string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FeatureLineageOnlineFeature) EncodeValues(key string, v *url.Values) error {
	pb, err := featureLineageOnlineFeatureToPb(&st)
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

func (st *FeatureLineageOnlineFeature) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &featureLineageOnlineFeaturePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := featureLineageOnlineFeatureFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FeatureLineageOnlineFeature) MarshalJSON() ([]byte, error) {
	pb, err := featureLineageOnlineFeatureToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Feature list wrap all the features for a model version
type FeatureList struct {

	// Wire name: 'features'
	Features []Feature `json:"features,omitempty"`
}

func (st FeatureList) EncodeValues(key string, v *url.Values) error {
	pb, err := featureListToPb(&st)
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

func (st *FeatureList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &featureListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := featureListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FeatureList) MarshalJSON() ([]byte, error) {
	pb, err := featureListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents a tag on a feature in a feature table.
type FeatureTag struct {

	// Wire name: 'key'
	Key string `json:"key"`

	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FeatureTag) EncodeValues(key string, v *url.Values) error {
	pb, err := featureTagToPb(&st)
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

func (st *FeatureTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &featureTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := featureTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FeatureTag) MarshalJSON() ([]byte, error) {
	pb, err := featureTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Metadata of a single artifact file or directory.
type FileInfo struct {
	// The size in bytes of the file. Unset for directories.
	// Wire name: 'file_size'
	FileSize int64 `json:"file_size,omitempty"`
	// Whether the path is a directory.
	// Wire name: 'is_dir'
	IsDir bool `json:"is_dir,omitempty"`
	// The path relative to the root artifact directory run.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FileInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := fileInfoToPb(&st)
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
	ModelId string `json:"-" tf:"-"`
	// Whether or not the model is ready for use.
	// ``"LOGGED_MODEL_UPLOAD_FAILED"`` indicates that something went wrong when
	// logging the model weights / agent code.
	// Wire name: 'status'
	Status LoggedModelStatus `json:"status"`
}

func (st FinalizeLoggedModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := finalizeLoggedModelRequestToPb(&st)
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
	Model *LoggedModel `json:"model,omitempty"`
}

func (st FinalizeLoggedModelResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := finalizeLoggedModelResponseToPb(&st)
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
	ExperimentId string `json:"experiment_id,omitempty"`
	// The URL to the forecasting experiment page.
	// Wire name: 'experiment_page_url'
	ExperimentPageUrl string `json:"experiment_page_url,omitempty"`
	// The current state of the forecasting experiment.
	// Wire name: 'state'
	State ForecastingExperimentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ForecastingExperiment) EncodeValues(key string, v *url.Values) error {
	pb, err := forecastingExperimentToPb(&st)
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

type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName string `json:"-" tf:"-"`
}

func (st GetByNameRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getByNameRequestToPb(&st)
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

type GetExperimentByNameResponse struct {
	// Experiment details.
	// Wire name: 'experiment'
	Experiment *Experiment `json:"experiment,omitempty"`
}

func (st GetExperimentByNameResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getExperimentByNameResponseToPb(&st)
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

type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" tf:"-"`
}

func (st GetExperimentPermissionLevelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getExperimentPermissionLevelsRequestToPb(&st)
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
	PermissionLevels []ExperimentPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st GetExperimentPermissionLevelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getExperimentPermissionLevelsResponseToPb(&st)
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

type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" tf:"-"`
}

func (st GetExperimentPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getExperimentPermissionsRequestToPb(&st)
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

type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId string `json:"-" tf:"-"`
}

func (st GetExperimentRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getExperimentRequestToPb(&st)
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
	Experiment *Experiment `json:"experiment,omitempty"`
}

func (st GetExperimentResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getExperimentResponseToPb(&st)
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

type GetFeatureLineageRequest struct {
	// The name of the feature.
	FeatureName string `json:"-" tf:"-"`
	// The full name of the feature table in Unity Catalog.
	TableName string `json:"-" tf:"-"`
}

func (st GetFeatureLineageRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getFeatureLineageRequestToPb(&st)
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

func (st *GetFeatureLineageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getFeatureLineageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getFeatureLineageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetFeatureLineageRequest) MarshalJSON() ([]byte, error) {
	pb, err := getFeatureLineageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetFeatureTagRequest struct {
	FeatureName string `json:"-" tf:"-"`

	Key string `json:"-" tf:"-"`

	TableName string `json:"-" tf:"-"`
}

func (st GetFeatureTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getFeatureTagRequestToPb(&st)
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

func (st *GetFeatureTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getFeatureTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getFeatureTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetFeatureTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := getFeatureTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetForecastingExperimentRequest struct {
	// The unique ID of a forecasting experiment
	ExperimentId string `json:"-" tf:"-"`
}

func (st GetForecastingExperimentRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getForecastingExperimentRequestToPb(&st)
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

type GetHistoryRequest struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	MaxResults int `json:"-" tf:"-"`
	// Name of the metric.
	MetricKey string `json:"-" tf:"-"`
	// Token indicating the page of metric histories to fetch.
	PageToken string `json:"-" tf:"-"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId string `json:"-" tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run from which to fetch
	// metric values. This field will be removed in a future MLflow version.
	RunUuid string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetHistoryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getHistoryRequestToPb(&st)
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
	Name string `json:"name"`
	// List of stages.
	// Wire name: 'stages'
	Stages []string `json:"stages,omitempty"`
}

func (st GetLatestVersionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getLatestVersionsRequestToPb(&st)
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
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
}

func (st GetLatestVersionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getLatestVersionsResponseToPb(&st)
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

type GetLoggedModelRequest struct {
	// The ID of the logged model to retrieve.
	ModelId string `json:"-" tf:"-"`
}

func (st GetLoggedModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getLoggedModelRequestToPb(&st)
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
	Model *LoggedModel `json:"model,omitempty"`
}

func (st GetLoggedModelResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getLoggedModelResponseToPb(&st)
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
	Metrics []Metric `json:"metrics,omitempty"`
	// A token that can be used to issue a query for the next page of metric
	// history values. A missing token indicates that no additional metrics are
	// available to fetch.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetMetricHistoryResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getMetricHistoryResponseToPb(&st)
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

type GetModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"-" tf:"-"`
}

func (st GetModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getModelRequestToPb(&st)
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
	RegisteredModelDatabricks *ModelDatabricks `json:"registered_model_databricks,omitempty"`
}

func (st GetModelResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getModelResponseToPb(&st)
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

type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name string `json:"-" tf:"-"`
	// Model version number
	Version string `json:"-" tf:"-"`
}

func (st GetModelVersionDownloadUriRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getModelVersionDownloadUriRequestToPb(&st)
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
	ArtifactUri string `json:"artifact_uri,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetModelVersionDownloadUriResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getModelVersionDownloadUriResponseToPb(&st)
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

type GetModelVersionRequest struct {
	// Name of the registered model
	Name string `json:"-" tf:"-"`
	// Model version number
	Version string `json:"-" tf:"-"`
}

func (st GetModelVersionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getModelVersionRequestToPb(&st)
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
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

func (st GetModelVersionResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getModelVersionResponseToPb(&st)
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

type GetOnlineStoreRequest struct {
	// Name of the online store to get.
	Name string `json:"-" tf:"-"`
}

func (st GetOnlineStoreRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getOnlineStoreRequestToPb(&st)
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

func (st *GetOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getOnlineStoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getOnlineStoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetOnlineStoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := getOnlineStoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" tf:"-"`
}

func (st GetRegisteredModelPermissionLevelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRegisteredModelPermissionLevelsRequestToPb(&st)
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
	PermissionLevels []RegisteredModelPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st GetRegisteredModelPermissionLevelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getRegisteredModelPermissionLevelsResponseToPb(&st)
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

type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" tf:"-"`
}

func (st GetRegisteredModelPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRegisteredModelPermissionsRequestToPb(&st)
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

type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId string `json:"-" tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetRunRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRunRequestToPb(&st)
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
	Run *Run `json:"run,omitempty"`
}

func (st GetRunResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getRunResponseToPb(&st)
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
	Authorization string `json:"authorization,omitempty"`
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	// Wire name: 'enable_ssl_verification'
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	// Wire name: 'secret'
	Secret string `json:"secret,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'url'
	Url string `json:"url"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st HttpUrlSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := httpUrlSpecToPb(&st)
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
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st HttpUrlSpecWithoutSecret) EncodeValues(key string, v *url.Values) error {
	pb, err := httpUrlSpecWithoutSecretToPb(&st)
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
	Key string `json:"key"`
	// The tag value.
	// Wire name: 'value'
	Value string `json:"value"`
}

func (st InputTag) EncodeValues(key string, v *url.Values) error {
	pb, err := inputTagToPb(&st)
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
	AccessToken string `json:"access_token"`
	// ID of the job that the webhook runs.
	// Wire name: 'job_id'
	JobId string `json:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	// Wire name: 'workspace_url'
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st JobSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := jobSpecToPb(&st)
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
	JobId string `json:"job_id,omitempty"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	// Wire name: 'workspace_url'
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st JobSpecWithoutSecret) EncodeValues(key string, v *url.Values) error {
	pb, err := jobSpecWithoutSecretToPb(&st)
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

type ListArtifactsRequest struct {
	// The token indicating the page of artifact results to fetch. `page_token`
	// is not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	PageToken string `json:"-" tf:"-"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path string `json:"-" tf:"-"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId string `json:"-" tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListArtifactsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listArtifactsRequestToPb(&st)
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
	Files []FileInfo `json:"files,omitempty"`
	// The token that can be used to retrieve the next page of artifact results.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// The root artifact directory for the run.
	// Wire name: 'root_uri'
	RootUri string `json:"root_uri,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListArtifactsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listArtifactsResponseToPb(&st)
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

type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults int64 `json:"-" tf:"-"`
	// Token indicating the page of experiments to fetch
	PageToken string `json:"-" tf:"-"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType ViewType `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListExperimentsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listExperimentsRequestToPb(&st)
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
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListExperimentsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listExperimentsResponseToPb(&st)
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

type ListFeatureTagsRequest struct {
	FeatureName string `json:"-" tf:"-"`
	// The maximum number of results to return.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" tf:"-"`

	TableName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListFeatureTagsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listFeatureTagsRequestToPb(&st)
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

func (st *ListFeatureTagsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFeatureTagsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFeatureTagsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFeatureTagsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listFeatureTagsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Response message for ListFeatureTag.
type ListFeatureTagsResponse struct {

	// Wire name: 'feature_tags'
	FeatureTags []FeatureTag `json:"feature_tags,omitempty"`
	// Pagination token to request the next page of results for this query.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListFeatureTagsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listFeatureTagsResponseToPb(&st)
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

func (st *ListFeatureTagsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFeatureTagsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFeatureTagsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFeatureTagsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listFeatureTagsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults int64 `json:"-" tf:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListModelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listModelsRequestToPb(&st)
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
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'registered_models'
	RegisteredModels []Model `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListModelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listModelsResponseToPb(&st)
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

type ListOnlineStoresRequest struct {
	// The maximum number of results to return. Defaults to 100 if not
	// specified.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListOnlineStoresRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listOnlineStoresRequestToPb(&st)
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

func (st *ListOnlineStoresRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listOnlineStoresRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listOnlineStoresRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListOnlineStoresRequest) MarshalJSON() ([]byte, error) {
	pb, err := listOnlineStoresRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListOnlineStoresResponse struct {
	// Pagination token to request the next page of results for this query.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of online stores.
	// Wire name: 'online_stores'
	OnlineStores []OnlineStore `json:"online_stores,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListOnlineStoresResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listOnlineStoresResponseToPb(&st)
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

func (st *ListOnlineStoresResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listOnlineStoresResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listOnlineStoresResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListOnlineStoresResponse) MarshalJSON() ([]byte, error) {
	pb, err := listOnlineStoresResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// Array of registry webhooks.
	// Wire name: 'webhooks'
	Webhooks []RegistryWebhook `json:"webhooks,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListRegistryWebhooks) EncodeValues(key string, v *url.Values) error {
	pb, err := listRegistryWebhooksToPb(&st)
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

type ListTransitionRequestsRequest struct {
	// Name of the registered model.
	Name string `json:"-" tf:"-"`
	// Version of the model.
	Version string `json:"-" tf:"-"`
}

func (st ListTransitionRequestsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listTransitionRequestsRequestToPb(&st)
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
	Requests []Activity `json:"requests,omitempty"`
}

func (st ListTransitionRequestsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listTransitionRequestsResponseToPb(&st)
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
	Events []RegistryWebhookEvent `json:"-" tf:"-"`

	MaxResults int64 `json:"-" tf:"-"`
	// Registered model name If not specified, all webhooks associated with the
	// specified events are listed, regardless of their associated model.
	ModelName string `json:"-" tf:"-"`
	// Token indicating the page of artifact results to fetch
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListWebhooksRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listWebhooksRequestToPb(&st)
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
	Metrics []Metric `json:"metrics,omitempty"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	// Wire name: 'params'
	Params []Param `json:"params,omitempty"`
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	// Wire name: 'tags'
	Tags []RunTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LogBatch) EncodeValues(key string, v *url.Values) error {
	pb, err := logBatchToPb(&st)
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

type LogInputs struct {
	// Dataset inputs
	// Wire name: 'datasets'
	Datasets []DatasetInput `json:"datasets,omitempty"`
	// Model inputs
	// Wire name: 'models'
	Models []ModelInput `json:"models,omitempty"`
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string `json:"run_id"`
}

func (st LogInputs) EncodeValues(key string, v *url.Values) error {
	pb, err := logInputsToPb(&st)
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

type LogLoggedModelParamsRequest struct {
	// The ID of the logged model to log params for.
	ModelId string `json:"-" tf:"-"`
	// Parameters to attach to the model.
	// Wire name: 'params'
	Params []LoggedModelParameter `json:"params,omitempty"`
}

func (st LogLoggedModelParamsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := logLoggedModelParamsRequestToPb(&st)
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

type LogMetric struct {
	// Dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	// Wire name: 'dataset_digest'
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	// Wire name: 'dataset_name'
	DatasetName string `json:"dataset_name,omitempty"`
	// Name of the metric.
	// Wire name: 'key'
	Key string `json:"key"`
	// ID of the logged model associated with the metric, if applicable
	// Wire name: 'model_id'
	ModelId string `json:"model_id,omitempty"`
	// ID of the run under which to log the metric. Must be provided.
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `json:"run_uuid,omitempty"`
	// Step at which to log the metric
	// Wire name: 'step'
	Step int64 `json:"step,omitempty"`
	// Unix timestamp in milliseconds at the time metric was logged.
	// Wire name: 'timestamp'
	Timestamp int64 `json:"timestamp"`
	// Double value of the metric being logged.
	// Wire name: 'value'
	Value float64 `json:"value"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LogMetric) EncodeValues(key string, v *url.Values) error {
	pb, err := logMetricToPb(&st)
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

type LogModel struct {
	// MLmodel file in json format.
	// Wire name: 'model_json'
	ModelJson string `json:"model_json,omitempty"`
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LogModel) EncodeValues(key string, v *url.Values) error {
	pb, err := logModelToPb(&st)
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

type LogOutputsRequest struct {
	// The model outputs from the Run.
	// Wire name: 'models'
	Models []ModelOutput `json:"models,omitempty"`
	// The ID of the Run from which to log outputs.
	// Wire name: 'run_id'
	RunId string `json:"run_id"`
}

func (st LogOutputsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := logOutputsRequestToPb(&st)
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

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	// Wire name: 'key'
	Key string `json:"key"`
	// ID of the run under which to log the param. Must be provided.
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the param being logged. Maximum size is 500 bytes.
	// Wire name: 'value'
	Value string `json:"value"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LogParam) EncodeValues(key string, v *url.Values) error {
	pb, err := logParamToPb(&st)
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

// A logged model message includes logged model attributes, tags, registration
// info, params, and linked run metrics.
type LoggedModel struct {
	// The params and metrics attached to the logged model.
	// Wire name: 'data'
	Data *LoggedModelData `json:"data,omitempty"`
	// The logged model attributes such as model ID, status, tags, etc.
	// Wire name: 'info'
	Info *LoggedModelInfo `json:"info,omitempty"`
}

func (st LoggedModel) EncodeValues(key string, v *url.Values) error {
	pb, err := loggedModelToPb(&st)
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
	Metrics []Metric `json:"metrics,omitempty"`
	// Immutable string key-value pairs of the model.
	// Wire name: 'params'
	Params []LoggedModelParameter `json:"params,omitempty"`
}

func (st LoggedModelData) EncodeValues(key string, v *url.Values) error {
	pb, err := loggedModelDataToPb(&st)
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
	ArtifactUri string `json:"artifact_uri,omitempty"`
	// The timestamp when the model was created in milliseconds since the UNIX
	// epoch.
	// Wire name: 'creation_timestamp_ms'
	CreationTimestampMs int64 `json:"creation_timestamp_ms,omitempty"`
	// The ID of the user or principal that created the model.
	// Wire name: 'creator_id'
	CreatorId int64 `json:"creator_id,omitempty"`
	// The ID of the experiment that owns the model.
	// Wire name: 'experiment_id'
	ExperimentId string `json:"experiment_id,omitempty"`
	// The timestamp when the model was last updated in milliseconds since the
	// UNIX epoch.
	// Wire name: 'last_updated_timestamp_ms'
	LastUpdatedTimestampMs int64 `json:"last_updated_timestamp_ms,omitempty"`
	// The unique identifier for the logged model.
	// Wire name: 'model_id'
	ModelId string `json:"model_id,omitempty"`
	// The type of model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	// Wire name: 'model_type'
	ModelType string `json:"model_type,omitempty"`
	// The name of the model.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The ID of the run that created the model.
	// Wire name: 'source_run_id'
	SourceRunId string `json:"source_run_id,omitempty"`
	// The status of whether or not the model is ready for use.
	// Wire name: 'status'
	Status LoggedModelStatus `json:"status,omitempty"`
	// Details on the current model status.
	// Wire name: 'status_message'
	StatusMessage string `json:"status_message,omitempty"`
	// Mutable string key-value pairs set on the model.
	// Wire name: 'tags'
	Tags []LoggedModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LoggedModelInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := loggedModelInfoToPb(&st)
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
	Key string `json:"key,omitempty"`
	// The value of this param.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LoggedModelParameter) EncodeValues(key string, v *url.Values) error {
	pb, err := loggedModelParameterToPb(&st)
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
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// The tag value.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LoggedModelTag) EncodeValues(key string, v *url.Values) error {
	pb, err := loggedModelTagToPb(&st)
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
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	// Wire name: 'dataset_name'
	DatasetName string `json:"dataset_name,omitempty"`
	// The key identifying the metric.
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// The ID of the logged model or registered model version associated with
	// the metric, if applicable.
	// Wire name: 'model_id'
	ModelId string `json:"model_id,omitempty"`
	// The ID of the run containing the metric.
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// The step at which the metric was logged.
	// Wire name: 'step'
	Step int64 `json:"step,omitempty"`
	// The timestamp at which the metric was recorded.
	// Wire name: 'timestamp'
	Timestamp int64 `json:"timestamp,omitempty"`
	// The value of the metric.
	// Wire name: 'value'
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Metric) EncodeValues(key string, v *url.Values) error {
	pb, err := metricToPb(&st)
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
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Description of this `registered_model`.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	// Wire name: 'latest_versions'
	LatestVersions []ModelVersion `json:"latest_versions,omitempty"`
	// Unique name for the model.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	// Wire name: 'tags'
	Tags []ModelTag `json:"tags,omitempty"`
	// User that created this `registered_model`
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Model) EncodeValues(key string, v *url.Values) error {
	pb, err := modelToPb(&st)
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
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User-specified description for the object.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Last update time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Array of model versions, each the latest version for its stage.
	// Wire name: 'latest_versions'
	LatestVersions []ModelVersion `json:"latest_versions,omitempty"`
	// Name of the model.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Permission level granted for the requesting user on this registered model
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Array of tags associated with the model.
	// Wire name: 'tags'
	Tags []ModelTag `json:"tags,omitempty"`
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ModelDatabricks) EncodeValues(key string, v *url.Values) error {
	pb, err := modelDatabricksToPb(&st)
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
	ModelId string `json:"model_id"`
}

func (st ModelInput) EncodeValues(key string, v *url.Values) error {
	pb, err := modelInputToPb(&st)
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
	ModelId string `json:"model_id"`
	// The step at which the model was produced.
	// Wire name: 'step'
	Step int64 `json:"step"`
}

func (st ModelOutput) EncodeValues(key string, v *url.Values) error {
	pb, err := modelOutputToPb(&st)
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

// Tag for a registered model
type ModelTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// The tag value.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ModelTag) EncodeValues(key string, v *url.Values) error {
	pb, err := modelTagToPb(&st)
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
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Current stage for this `model_version`.
	// Wire name: 'current_stage'
	CurrentStage string `json:"current_stage,omitempty"`
	// Description of this `model_version`.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Unique name of the model
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// Run Link: Direct link to the run that generated this version
	// Wire name: 'run_link'
	RunLink string `json:"run_link,omitempty"`
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	// Wire name: 'source'
	Source string `json:"source,omitempty"`
	// Current status of `model_version`
	// Wire name: 'status'
	Status ModelVersionStatus `json:"status,omitempty"`
	// Details on current `status`, if it is pending or failed.
	// Wire name: 'status_message'
	StatusMessage string `json:"status_message,omitempty"`
	// Tags: Additional metadata key-value pairs for this `model_version`.
	// Wire name: 'tags'
	Tags []ModelVersionTag `json:"tags,omitempty"`
	// User that created this `model_version`.
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`
	// Model's version number.
	// Wire name: 'version'
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ModelVersion) EncodeValues(key string, v *url.Values) error {
	pb, err := modelVersionToPb(&st)
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
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	// Wire name: 'current_stage'
	CurrentStage string `json:"current_stage,omitempty"`
	// User-specified description for the object.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Email Subscription Status: This is the subscription status of the user to
	// the model version Users get subscribed by interacting with the model
	// version.
	// Wire name: 'email_subscription_status'
	EmailSubscriptionStatus RegistryEmailSubscriptionType `json:"email_subscription_status,omitempty"`
	// Feature lineage of `model_version`.
	// Wire name: 'feature_list'
	FeatureList *FeatureList `json:"feature_list,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the model.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Open requests for this `model_versions`. Gap in sequence number is
	// intentional and is done in order to match field sequence numbers of
	// `ModelVersion` proto message
	// Wire name: 'open_requests'
	OpenRequests []Activity `json:"open_requests,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	// Wire name: 'run_link'
	RunLink string `json:"run_link,omitempty"`
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	// Wire name: 'source'
	Source string `json:"source,omitempty"`

	// Wire name: 'status'
	Status Status `json:"status,omitempty"`
	// Details on the current status, for example why registration failed.
	// Wire name: 'status_message'
	StatusMessage string `json:"status_message,omitempty"`
	// Array of tags that are associated with the model version.
	// Wire name: 'tags'
	Tags []ModelVersionTag `json:"tags,omitempty"`
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`
	// Version of the model.
	// Wire name: 'version'
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ModelVersionDatabricks) EncodeValues(key string, v *url.Values) error {
	pb, err := modelVersionDatabricksToPb(&st)
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
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// The tag value.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ModelVersionTag) EncodeValues(key string, v *url.Values) error {
	pb, err := modelVersionTagToPb(&st)
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

// An OnlineStore is a logical database instance that stores and serves features
// online.
type OnlineStore struct {
	// The capacity of the online store. Valid values are "CU_1", "CU_2",
	// "CU_4", "CU_8".
	// Wire name: 'capacity'
	Capacity string `json:"capacity"`
	// The timestamp when the online store was created.
	// Wire name: 'creation_time'
	CreationTime string `json:"creation_time,omitempty"`
	// The email of the creator of the online store.
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// The name of the online store. This is the unique identifier for the
	// online store.
	// Wire name: 'name'
	Name string `json:"name"`
	// The current state of the online store.
	// Wire name: 'state'
	State OnlineStoreState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OnlineStore) EncodeValues(key string, v *url.Values) error {
	pb, err := onlineStoreToPb(&st)
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

func (st *OnlineStore) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &onlineStorePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := onlineStoreFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OnlineStore) MarshalJSON() ([]byte, error) {
	pb, err := onlineStoreToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// Value associated with this param.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Param) EncodeValues(key string, v *url.Values) error {
	pb, err := paramToPb(&st)
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
	// Wire name: 'online_store'
	OnlineStore string `json:"online_store"`
	// The full three-part (catalog, schema, table) name of the online table.
	// Wire name: 'online_table_name'
	OnlineTableName string `json:"online_table_name"`
	// The publish mode of the pipeline that syncs the online table with the
	// source table. Defaults to TRIGGERED if not specified. All publish modes
	// require the source table to have Change Data Feed (CDF) enabled.
	// Wire name: 'publish_mode'
	PublishMode PublishSpecPublishMode `json:"publish_mode,omitempty"`
}

func (st PublishSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := publishSpecToPb(&st)
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

func (st *PublishSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &publishSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := publishSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PublishSpec) MarshalJSON() ([]byte, error) {
	pb, err := publishSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PublishSpecPublishMode string

const PublishSpecPublishModeContinuous PublishSpecPublishMode = `CONTINUOUS`

const PublishSpecPublishModeTriggered PublishSpecPublishMode = `TRIGGERED`

// String representation for [fmt.Print]
func (f *PublishSpecPublishMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PublishSpecPublishMode) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `TRIGGERED`:
		*f = PublishSpecPublishMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "TRIGGERED"`, v)
	}
}

// Values returns all possible values for PublishSpecPublishMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *PublishSpecPublishMode) Values() []PublishSpecPublishMode {
	return []PublishSpecPublishMode{
		PublishSpecPublishModeContinuous,
		PublishSpecPublishModeTriggered,
	}
}

// Type always returns PublishSpecPublishMode to satisfy [pflag.Value] interface
func (f *PublishSpecPublishMode) Type() string {
	return "PublishSpecPublishMode"
}

type PublishTableRequest struct {
	// The specification for publishing the online table from the source table.
	// Wire name: 'publish_spec'
	PublishSpec PublishSpec `json:"publish_spec"`
	// The full three-part (catalog, schema, table) name of the source table.
	SourceTableName string `json:"-" tf:"-"`
}

func (st PublishTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := publishTableRequestToPb(&st)
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

func (st *PublishTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &publishTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := publishTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PublishTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := publishTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PublishTableResponse struct {
	// The full three-part (catalog, schema, table) name of the online table.
	// Wire name: 'online_table_name'
	OnlineTableName string `json:"online_table_name,omitempty"`
	// The ID of the pipeline that syncs the online table with the source table.
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PublishTableResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := publishTableResponseToPb(&st)
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

func (st *PublishTableResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &publishTableResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := publishTableResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PublishTableResponse) MarshalJSON() ([]byte, error) {
	pb, err := publishTableResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RegisteredModelAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegisteredModelAccessControlRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelAccessControlRequestToPb(&st)
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
	AllPermissions []RegisteredModelPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegisteredModelAccessControlResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelAccessControlResponseToPb(&st)
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
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegisteredModelPermission) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelPermissionToPb(&st)
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

	// Wire name: 'access_control_list'
	AccessControlList []RegisteredModelAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegisteredModelPermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelPermissionsToPb(&st)
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
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegisteredModelPermissionsDescription) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelPermissionsDescriptionToPb(&st)
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
	AccessControlList []RegisteredModelAccessControlRequest `json:"access_control_list,omitempty"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" tf:"-"`
}

func (st RegisteredModelPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelPermissionsRequestToPb(&st)
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
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User-specified description for the webhook.
	// Wire name: 'description'
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
	// Wire name: 'events'
	Events []RegistryWebhookEvent `json:"events,omitempty"`

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpecWithoutSecret `json:"http_url_spec,omitempty"`
	// Webhook ID
	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'job_spec'
	JobSpec *JobSpecWithoutSecret `json:"job_spec,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the model whose events would trigger this webhook.
	// Wire name: 'model_name'
	ModelName string `json:"model_name,omitempty"`

	// Wire name: 'status'
	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegistryWebhook) EncodeValues(key string, v *url.Values) error {
	pb, err := registryWebhookToPb(&st)
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
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	// Wire name: 'name'
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
	// Wire name: 'stage'
	Stage string `json:"stage"`
	// Version of the model.
	// Wire name: 'version'
	Version string `json:"version"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RejectTransitionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := rejectTransitionRequestToPb(&st)
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
	// New activity generated as a result of this operation.
	// Wire name: 'activity'
	Activity *Activity `json:"activity,omitempty"`
}

func (st RejectTransitionRequestResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := rejectTransitionRequestResponseToPb(&st)
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
	Name string `json:"name"`
	// If provided, updates the name for this `registered_model`.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RenameModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := renameModelRequestToPb(&st)
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
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

func (st RenameModelResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := renameModelResponseToPb(&st)
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
	ExperimentId string `json:"experiment_id"`
}

func (st RestoreExperiment) EncodeValues(key string, v *url.Values) error {
	pb, err := restoreExperimentToPb(&st)
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

type RestoreRun struct {
	// ID of the run to restore.
	// Wire name: 'run_id'
	RunId string `json:"run_id"`
}

func (st RestoreRun) EncodeValues(key string, v *url.Values) error {
	pb, err := restoreRunToPb(&st)
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

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	// Wire name: 'experiment_id'
	ExperimentId string `json:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	// Wire name: 'max_runs'
	MaxRuns int `json:"max_runs,omitempty"`
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	// Wire name: 'min_timestamp_millis'
	MinTimestampMillis int64 `json:"min_timestamp_millis"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RestoreRuns) EncodeValues(key string, v *url.Values) error {
	pb, err := restoreRunsToPb(&st)
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
	RunsRestored int `json:"runs_restored,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RestoreRunsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := restoreRunsResponseToPb(&st)
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
	Data *RunData `json:"data,omitempty"`
	// Run metadata.
	// Wire name: 'info'
	Info *RunInfo `json:"info,omitempty"`
	// Run inputs.
	// Wire name: 'inputs'
	Inputs *RunInputs `json:"inputs,omitempty"`
}

func (st Run) EncodeValues(key string, v *url.Values) error {
	pb, err := runToPb(&st)
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
	Metrics []Metric `json:"metrics,omitempty"`
	// Run parameters.
	// Wire name: 'params'
	Params []Param `json:"params,omitempty"`
	// Additional metadata key-value pairs.
	// Wire name: 'tags'
	Tags []RunTag `json:"tags,omitempty"`
}

func (st RunData) EncodeValues(key string, v *url.Values) error {
	pb, err := runDataToPb(&st)
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
	ArtifactUri string `json:"artifact_uri,omitempty"`
	// Unix timestamp of when the run ended in milliseconds.
	// Wire name: 'end_time'
	EndTime int64 `json:"end_time,omitempty"`
	// The experiment ID.
	// Wire name: 'experiment_id'
	ExperimentId string `json:"experiment_id,omitempty"`
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	// Wire name: 'lifecycle_stage'
	LifecycleStage string `json:"lifecycle_stage,omitempty"`
	// Unique identifier for the run.
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// The name of the run.
	// Wire name: 'run_name'
	RunName string `json:"run_name,omitempty"`
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `json:"run_uuid,omitempty"`
	// Unix timestamp of when the run started in milliseconds.
	// Wire name: 'start_time'
	StartTime int64 `json:"start_time,omitempty"`
	// Current status of the run.
	// Wire name: 'status'
	Status RunInfoStatus `json:"status,omitempty"`
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RunInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := runInfoToPb(&st)
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
	// Wire name: 'dataset_inputs'
	DatasetInputs []DatasetInput `json:"dataset_inputs,omitempty"`
	// Model inputs to the Run.
	// Wire name: 'model_inputs'
	ModelInputs []ModelInput `json:"model_inputs,omitempty"`
}

func (st RunInputs) EncodeValues(key string, v *url.Values) error {
	pb, err := runInputsToPb(&st)
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
	Key string `json:"key,omitempty"`
	// The tag value.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RunTag) EncodeValues(key string, v *url.Values) error {
	pb, err := runTagToPb(&st)
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
	Filter string `json:"filter,omitempty"`
	// Maximum number of experiments desired. Max threshold is 3000.
	// Wire name: 'max_results'
	MaxResults int64 `json:"max_results,omitempty"`
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	// Wire name: 'order_by'
	OrderBy []string `json:"order_by,omitempty"`
	// Token indicating the page of experiments to fetch
	// Wire name: 'page_token'
	PageToken string `json:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	// Wire name: 'view_type'
	ViewType ViewType `json:"view_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchExperiments) EncodeValues(key string, v *url.Values) error {
	pb, err := searchExperimentsToPb(&st)
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
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchExperimentsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := searchExperimentsResponseToPb(&st)
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
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset.
	// Wire name: 'dataset_name'
	DatasetName string `json:"dataset_name"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchLoggedModelsDataset) EncodeValues(key string, v *url.Values) error {
	pb, err := searchLoggedModelsDatasetToPb(&st)
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
	Ascending bool `json:"ascending,omitempty"`
	// If ``field_name`` refers to a metric, this field specifies the digest of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name and digest will be considered for ordering. This
	// field may only be set if ``dataset_name`` is also set.
	// Wire name: 'dataset_digest'
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// If ``field_name`` refers to a metric, this field specifies the name of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name will be considered for ordering. This field may
	// only be set if ``field_name`` refers to a metric.
	// Wire name: 'dataset_name'
	DatasetName string `json:"dataset_name,omitempty"`
	// The name of the field to order by, e.g. "metrics.accuracy".
	// Wire name: 'field_name'
	FieldName string `json:"field_name"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchLoggedModelsOrderBy) EncodeValues(key string, v *url.Values) error {
	pb, err := searchLoggedModelsOrderByToPb(&st)
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
	Datasets []SearchLoggedModelsDataset `json:"datasets,omitempty"`
	// The IDs of the experiments in which to search for logged models.
	// Wire name: 'experiment_ids'
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	// A filter expression over logged model info and data that allows returning
	// a subset of logged models. The syntax is a subset of SQL that supports
	// AND'ing together binary operations.
	//
	// Example: ``params.alpha < 0.3 AND metrics.accuracy > 0.9``.
	// Wire name: 'filter'
	Filter string `json:"filter,omitempty"`
	// The maximum number of Logged Models to return. The maximum limit is 50.
	// Wire name: 'max_results'
	MaxResults int `json:"max_results,omitempty"`
	// The list of columns for ordering the results, with additional fields for
	// sorting criteria.
	// Wire name: 'order_by'
	OrderBy []SearchLoggedModelsOrderBy `json:"order_by,omitempty"`
	// The token indicating the page of logged models to fetch.
	// Wire name: 'page_token'
	PageToken string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchLoggedModelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := searchLoggedModelsRequestToPb(&st)
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
	Models []LoggedModel `json:"models,omitempty"`
	// The token that can be used to retrieve the next page of logged models.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchLoggedModelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := searchLoggedModelsResponseToPb(&st)
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

type SearchModelVersionsRequest struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" tf:"-"`
	// Maximum number of models desired. Max threshold is 10K.
	MaxResults int64 `json:"-" tf:"-"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	OrderBy []string `json:"-" tf:"-"`
	// Pagination token to go to next page based on previous search query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchModelVersionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := searchModelVersionsRequestToPb(&st)
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
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
	// Pagination token to request next page of models for the same search
	// query.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchModelVersionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := searchModelVersionsResponseToPb(&st)
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

type SearchModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" tf:"-"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	MaxResults int64 `json:"-" tf:"-"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	OrderBy []string `json:"-" tf:"-"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchModelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := searchModelsRequestToPb(&st)
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
	NextPageToken string `json:"next_page_token,omitempty"`
	// Registered Models that match the search criteria.
	// Wire name: 'registered_models'
	RegisteredModels []Model `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchModelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := searchModelsResponseToPb(&st)
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
	// Wire name: 'filter'
	Filter string `json:"filter,omitempty"`
	// Maximum number of runs desired. Max threshold is 50000
	// Wire name: 'max_results'
	MaxResults int `json:"max_results,omitempty"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional `"DESC"` or `"ASC"` annotation, where `"ASC"`
	// is the default. Example: `["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"]`. Tiebreaks are done by start_time `DESC` followed by
	// `run_id` for runs with the same start time (and this is the default
	// ordering criterion if order_by is not provided).
	// Wire name: 'order_by'
	OrderBy []string `json:"order_by,omitempty"`
	// Token for the current page of runs.
	// Wire name: 'page_token'
	PageToken string `json:"page_token,omitempty"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	// Wire name: 'run_view_type'
	RunViewType ViewType `json:"run_view_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchRuns) EncodeValues(key string, v *url.Values) error {
	pb, err := searchRunsToPb(&st)
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
	NextPageToken string `json:"next_page_token,omitempty"`
	// Runs that match the search criteria.
	// Wire name: 'runs'
	Runs []Run `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SearchRunsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := searchRunsResponseToPb(&st)
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
	ExperimentId string `json:"experiment_id"`
	// Name of the tag. Keys up to 250 bytes in size are supported.
	// Wire name: 'key'
	Key string `json:"key"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	// Wire name: 'value'
	Value string `json:"value"`
}

func (st SetExperimentTag) EncodeValues(key string, v *url.Values) error {
	pb, err := setExperimentTagToPb(&st)
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

type SetLoggedModelTagsRequest struct {
	// The ID of the logged model to set the tags on.
	ModelId string `json:"-" tf:"-"`
	// The tags to set on the logged model.
	// Wire name: 'tags'
	Tags []LoggedModelTag `json:"tags,omitempty"`
}

func (st SetLoggedModelTagsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := setLoggedModelTagsRequestToPb(&st)
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

type SetModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	// Wire name: 'key'
	Key string `json:"key"`
	// Unique name of the model.
	// Wire name: 'name'
	Name string `json:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	// Wire name: 'value'
	Value string `json:"value"`
}

func (st SetModelTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := setModelTagRequestToPb(&st)
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

type SetModelVersionTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	// Wire name: 'key'
	Key string `json:"key"`
	// Unique name of the model.
	// Wire name: 'name'
	Name string `json:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	// Wire name: 'value'
	Value string `json:"value"`
	// Model version number.
	// Wire name: 'version'
	Version string `json:"version"`
}

func (st SetModelVersionTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := setModelVersionTagRequestToPb(&st)
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

type SetTag struct {
	// Name of the tag. Keys up to 250 bytes in size are supported.
	// Wire name: 'key'
	Key string `json:"key"`
	// ID of the run under which to log the tag. Must be provided.
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	// Wire name: 'value'
	Value string `json:"value"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SetTag) EncodeValues(key string, v *url.Values) error {
	pb, err := setTagToPb(&st)
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
	// Wire name: 'event'
	Event RegistryWebhookEvent `json:"event,omitempty"`
	// Webhook ID
	// Wire name: 'id'
	Id string `json:"id"`
}

func (st TestRegistryWebhookRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := testRegistryWebhookRequestToPb(&st)
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
	// Body of the response from the webhook URL
	// Wire name: 'body'
	Body string `json:"body,omitempty"`
	// Status code returned by the webhook URL
	// Wire name: 'status_code'
	StatusCode int `json:"status_code,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TestRegistryWebhookResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := testRegistryWebhookResponseToPb(&st)
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

// Details required to transition a model version's stage.
type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	// Wire name: 'archive_existing_versions'
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	// Wire name: 'name'
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
	// Wire name: 'stage'
	Stage string `json:"stage"`
	// Version of the model.
	// Wire name: 'version'
	Version string `json:"version"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TransitionModelVersionStageDatabricks) EncodeValues(key string, v *url.Values) error {
	pb, err := transitionModelVersionStageDatabricksToPb(&st)
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

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	// Wire name: 'available_actions'
	AvailableActions []ActivityAction `json:"available_actions,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
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
	// Wire name: 'to_stage'
	ToStage string `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TransitionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := transitionRequestToPb(&st)
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
	// Updated model version
	// Wire name: 'model_version_databricks'
	ModelVersionDatabricks *ModelVersionDatabricks `json:"model_version_databricks,omitempty"`
}

func (st TransitionStageResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := transitionStageResponseToPb(&st)
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

// Details required to edit a comment on a model version.
type UpdateComment struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string `json:"comment"`
	// Unique identifier of an activity
	// Wire name: 'id'
	Id string `json:"id"`
}

func (st UpdateComment) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCommentToPb(&st)
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
	// Updated comment object
	// Wire name: 'comment'
	Comment *CommentObject `json:"comment,omitempty"`
}

func (st UpdateCommentResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCommentResponseToPb(&st)
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
	ExperimentId string `json:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateExperiment) EncodeValues(key string, v *url.Values) error {
	pb, err := updateExperimentToPb(&st)
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

type UpdateFeatureTagRequest struct {
	FeatureName string `json:"-" tf:"-"`

	// Wire name: 'feature_tag'
	FeatureTag FeatureTag `json:"feature_tag"`

	Key string `json:"-" tf:"-"`

	TableName string `json:"-" tf:"-"`
	// The list of fields to update.
	UpdateMask string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateFeatureTagRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateFeatureTagRequestToPb(&st)
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

func (st *UpdateFeatureTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateFeatureTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateFeatureTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateFeatureTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateFeatureTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string `json:"name"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateModelRequestToPb(&st)
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

	// Wire name: 'registered_model'
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

func (st UpdateModelResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateModelResponseToPb(&st)
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
	Description string `json:"description,omitempty"`
	// Name of the registered model
	// Wire name: 'name'
	Name string `json:"name"`
	// Model version number
	// Wire name: 'version'
	Version string `json:"version"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateModelVersionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateModelVersionRequestToPb(&st)
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
	// Return new version number generated for this model in registry.
	// Wire name: 'model_version'
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

func (st UpdateModelVersionResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateModelVersionResponseToPb(&st)
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

type UpdateOnlineStoreRequest struct {
	// The name of the online store. This is the unique identifier for the
	// online store.
	Name string `json:"-" tf:"-"`
	// Online store to update.
	// Wire name: 'online_store'
	OnlineStore OnlineStore `json:"online_store"`
	// The list of fields to update.
	UpdateMask string `json:"-" tf:"-"`
}

func (st UpdateOnlineStoreRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateOnlineStoreRequestToPb(&st)
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

func (st *UpdateOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateOnlineStoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateOnlineStoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateOnlineStoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateOnlineStoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a registry webhook. Only the fields that need to
// be updated should be specified, and both `http_url_spec` and `job_spec`
// should not be specified in the same request.
type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	// Wire name: 'description'
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
	// Wire name: 'events'
	Events []RegistryWebhookEvent `json:"events,omitempty"`

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`
	// Webhook ID
	// Wire name: 'id'
	Id string `json:"id"`

	// Wire name: 'job_spec'
	JobSpec *JobSpec `json:"job_spec,omitempty"`

	// Wire name: 'status'
	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateRegistryWebhook) EncodeValues(key string, v *url.Values) error {
	pb, err := updateRegistryWebhookToPb(&st)
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
	EndTime int64 `json:"end_time,omitempty"`
	// ID of the run to update. Must be provided.
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// Updated name of the run.
	// Wire name: 'run_name'
	RunName string `json:"run_name,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run to update. This field
	// will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `json:"run_uuid,omitempty"`
	// Updated status of the run.
	// Wire name: 'status'
	Status UpdateRunStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateRun) EncodeValues(key string, v *url.Values) error {
	pb, err := updateRunToPb(&st)
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
	RunInfo *RunInfo `json:"run_info,omitempty"`
}

func (st UpdateRunResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateRunResponseToPb(&st)
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

	// Wire name: 'webhook'
	Webhook *RegistryWebhook `json:"webhook,omitempty"`
}

func (st UpdateWebhookResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateWebhookResponseToPb(&st)
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
