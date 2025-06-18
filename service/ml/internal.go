// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func activityToPb(st *Activity) (*activityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &activityPb{}
	pb.ActivityType = st.ActivityType
	pb.Comment = st.Comment
	pb.CreationTimestamp = st.CreationTimestamp
	pb.FromStage = st.FromStage
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.SystemComment = st.SystemComment
	pb.ToStage = st.ToStage
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type activityPb struct {
	ActivityType         ActivityType `json:"activity_type,omitempty"`
	Comment              string       `json:"comment,omitempty"`
	CreationTimestamp    int64        `json:"creation_timestamp,omitempty"`
	FromStage            Stage        `json:"from_stage,omitempty"`
	Id                   string       `json:"id,omitempty"`
	LastUpdatedTimestamp int64        `json:"last_updated_timestamp,omitempty"`
	SystemComment        string       `json:"system_comment,omitempty"`
	ToStage              Stage        `json:"to_stage,omitempty"`
	UserId               string       `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func activityFromPb(pb *activityPb) (*Activity, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Activity{}
	st.ActivityType = pb.ActivityType
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.FromStage = pb.FromStage
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.SystemComment = pb.SystemComment
	st.ToStage = pb.ToStage
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *activityPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st activityPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func approveTransitionRequestToPb(st *ApproveTransitionRequest) (*approveTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &approveTransitionRequestPb{}
	pb.ArchiveExistingVersions = st.ArchiveExistingVersions
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type approveTransitionRequestPb struct {
	ArchiveExistingVersions bool   `json:"archive_existing_versions"`
	Comment                 string `json:"comment,omitempty"`
	Name                    string `json:"name"`
	Stage                   Stage  `json:"stage"`
	Version                 string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func approveTransitionRequestFromPb(pb *approveTransitionRequestPb) (*ApproveTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApproveTransitionRequest{}
	st.ArchiveExistingVersions = pb.ArchiveExistingVersions
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *approveTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st approveTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func approveTransitionRequestResponseToPb(st *ApproveTransitionRequestResponse) (*approveTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &approveTransitionRequestResponsePb{}
	pb.Activity = st.Activity

	return pb, nil
}

type approveTransitionRequestResponsePb struct {
	Activity *Activity `json:"activity,omitempty"`
}

func approveTransitionRequestResponseFromPb(pb *approveTransitionRequestResponsePb) (*ApproveTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApproveTransitionRequestResponse{}
	st.Activity = pb.Activity

	return st, nil
}

func commentObjectToPb(st *CommentObject) (*commentObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &commentObjectPb{}
	pb.AvailableActions = st.AvailableActions
	pb.Comment = st.Comment
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type commentObjectPb struct {
	AvailableActions     []CommentActivityAction `json:"available_actions,omitempty"`
	Comment              string                  `json:"comment,omitempty"`
	CreationTimestamp    int64                   `json:"creation_timestamp,omitempty"`
	Id                   string                  `json:"id,omitempty"`
	LastUpdatedTimestamp int64                   `json:"last_updated_timestamp,omitempty"`
	UserId               string                  `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func commentObjectFromPb(pb *commentObjectPb) (*CommentObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CommentObject{}
	st.AvailableActions = pb.AvailableActions
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *commentObjectPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st commentObjectPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createCommentToPb(st *CreateComment) (*createCommentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCommentPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

type createCommentPb struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func createCommentFromPb(pb *createCommentPb) (*CreateComment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateComment{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

func createCommentResponseToPb(st *CreateCommentResponse) (*createCommentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCommentResponsePb{}
	pb.Comment = st.Comment

	return pb, nil
}

type createCommentResponsePb struct {
	Comment *CommentObject `json:"comment,omitempty"`
}

func createCommentResponseFromPb(pb *createCommentResponsePb) (*CreateCommentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCommentResponse{}
	st.Comment = pb.Comment

	return st, nil
}

func createExperimentToPb(st *CreateExperiment) (*createExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExperimentPb{}
	pb.ArtifactLocation = st.ArtifactLocation
	pb.Name = st.Name
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createExperimentPb struct {
	ArtifactLocation string          `json:"artifact_location,omitempty"`
	Name             string          `json:"name"`
	Tags             []ExperimentTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createExperimentFromPb(pb *createExperimentPb) (*CreateExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExperiment{}
	st.ArtifactLocation = pb.ArtifactLocation
	st.Name = pb.Name
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createExperimentResponseToPb(st *CreateExperimentResponse) (*createExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExperimentResponsePb{}
	pb.ExperimentId = st.ExperimentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createExperimentResponsePb struct {
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createExperimentResponseFromPb(pb *createExperimentResponsePb) (*CreateExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExperimentResponse{}
	st.ExperimentId = pb.ExperimentId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createExperimentResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createExperimentResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createForecastingExperimentRequestToPb(st *CreateForecastingExperimentRequest) (*createForecastingExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createForecastingExperimentRequestPb{}
	pb.CustomWeightsColumn = st.CustomWeightsColumn
	pb.ExperimentPath = st.ExperimentPath
	pb.ForecastGranularity = st.ForecastGranularity
	pb.ForecastHorizon = st.ForecastHorizon
	pb.FutureFeatureDataPath = st.FutureFeatureDataPath
	pb.HolidayRegions = st.HolidayRegions
	pb.IncludeFeatures = st.IncludeFeatures
	pb.MaxRuntime = st.MaxRuntime
	pb.PredictionDataPath = st.PredictionDataPath
	pb.PrimaryMetric = st.PrimaryMetric
	pb.RegisterTo = st.RegisterTo
	pb.SplitColumn = st.SplitColumn
	pb.TargetColumn = st.TargetColumn
	pb.TimeColumn = st.TimeColumn
	pb.TimeseriesIdentifierColumns = st.TimeseriesIdentifierColumns
	pb.TrainDataPath = st.TrainDataPath
	pb.TrainingFrameworks = st.TrainingFrameworks

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createForecastingExperimentRequestPb struct {
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

func createForecastingExperimentRequestFromPb(pb *createForecastingExperimentRequestPb) (*CreateForecastingExperimentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateForecastingExperimentRequest{}
	st.CustomWeightsColumn = pb.CustomWeightsColumn
	st.ExperimentPath = pb.ExperimentPath
	st.ForecastGranularity = pb.ForecastGranularity
	st.ForecastHorizon = pb.ForecastHorizon
	st.FutureFeatureDataPath = pb.FutureFeatureDataPath
	st.HolidayRegions = pb.HolidayRegions
	st.IncludeFeatures = pb.IncludeFeatures
	st.MaxRuntime = pb.MaxRuntime
	st.PredictionDataPath = pb.PredictionDataPath
	st.PrimaryMetric = pb.PrimaryMetric
	st.RegisterTo = pb.RegisterTo
	st.SplitColumn = pb.SplitColumn
	st.TargetColumn = pb.TargetColumn
	st.TimeColumn = pb.TimeColumn
	st.TimeseriesIdentifierColumns = pb.TimeseriesIdentifierColumns
	st.TrainDataPath = pb.TrainDataPath
	st.TrainingFrameworks = pb.TrainingFrameworks

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createForecastingExperimentRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createForecastingExperimentRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createForecastingExperimentResponseToPb(st *CreateForecastingExperimentResponse) (*createForecastingExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createForecastingExperimentResponsePb{}
	pb.ExperimentId = st.ExperimentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createForecastingExperimentResponsePb struct {
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createForecastingExperimentResponseFromPb(pb *createForecastingExperimentResponsePb) (*CreateForecastingExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateForecastingExperimentResponse{}
	st.ExperimentId = pb.ExperimentId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createForecastingExperimentResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createForecastingExperimentResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createLoggedModelRequestToPb(st *CreateLoggedModelRequest) (*createLoggedModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createLoggedModelRequestPb{}
	pb.ExperimentId = st.ExperimentId
	pb.ModelType = st.ModelType
	pb.Name = st.Name
	pb.Params = st.Params
	pb.SourceRunId = st.SourceRunId
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createLoggedModelRequestPb struct {
	ExperimentId string                 `json:"experiment_id"`
	ModelType    string                 `json:"model_type,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Params       []LoggedModelParameter `json:"params,omitempty"`
	SourceRunId  string                 `json:"source_run_id,omitempty"`
	Tags         []LoggedModelTag       `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createLoggedModelRequestFromPb(pb *createLoggedModelRequestPb) (*CreateLoggedModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateLoggedModelRequest{}
	st.ExperimentId = pb.ExperimentId
	st.ModelType = pb.ModelType
	st.Name = pb.Name
	st.Params = pb.Params
	st.SourceRunId = pb.SourceRunId
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createLoggedModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createLoggedModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createLoggedModelResponseToPb(st *CreateLoggedModelResponse) (*createLoggedModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createLoggedModelResponsePb{}
	pb.Model = st.Model

	return pb, nil
}

type createLoggedModelResponsePb struct {
	Model *LoggedModel `json:"model,omitempty"`
}

func createLoggedModelResponseFromPb(pb *createLoggedModelResponsePb) (*CreateLoggedModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateLoggedModelResponse{}
	st.Model = pb.Model

	return st, nil
}

func createModelRequestToPb(st *CreateModelRequest) (*createModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createModelRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createModelRequestPb struct {
	Description string     `json:"description,omitempty"`
	Name        string     `json:"name"`
	Tags        []ModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createModelRequestFromPb(pb *createModelRequestPb) (*CreateModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelRequest{}
	st.Description = pb.Description
	st.Name = pb.Name
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createModelResponseToPb(st *CreateModelResponse) (*createModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createModelResponsePb{}
	pb.RegisteredModel = st.RegisteredModel

	return pb, nil
}

type createModelResponsePb struct {
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

func createModelResponseFromPb(pb *createModelResponsePb) (*CreateModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelResponse{}
	st.RegisteredModel = pb.RegisteredModel

	return st, nil
}

func createModelVersionRequestToPb(st *CreateModelVersionRequest) (*createModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createModelVersionRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name
	pb.RunId = st.RunId
	pb.RunLink = st.RunLink
	pb.Source = st.Source
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createModelVersionRequestPb struct {
	Description string            `json:"description,omitempty"`
	Name        string            `json:"name"`
	RunId       string            `json:"run_id,omitempty"`
	RunLink     string            `json:"run_link,omitempty"`
	Source      string            `json:"source"`
	Tags        []ModelVersionTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createModelVersionRequestFromPb(pb *createModelVersionRequestPb) (*CreateModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelVersionRequest{}
	st.Description = pb.Description
	st.Name = pb.Name
	st.RunId = pb.RunId
	st.RunLink = pb.RunLink
	st.Source = pb.Source
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createModelVersionResponseToPb(st *CreateModelVersionResponse) (*createModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createModelVersionResponsePb{}
	pb.ModelVersion = st.ModelVersion

	return pb, nil
}

type createModelVersionResponsePb struct {
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

func createModelVersionResponseFromPb(pb *createModelVersionResponsePb) (*CreateModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelVersionResponse{}
	st.ModelVersion = pb.ModelVersion

	return st, nil
}

func createOnlineStoreRequestToPb(st *CreateOnlineStoreRequest) (*createOnlineStoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createOnlineStoreRequestPb{}
	pb.OnlineStore = st.OnlineStore

	return pb, nil
}

type createOnlineStoreRequestPb struct {
	OnlineStore OnlineStore `json:"online_store"`
}

func createOnlineStoreRequestFromPb(pb *createOnlineStoreRequestPb) (*CreateOnlineStoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOnlineStoreRequest{}
	st.OnlineStore = pb.OnlineStore

	return st, nil
}

func createRegistryWebhookToPb(st *CreateRegistryWebhook) (*createRegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRegistryWebhookPb{}
	pb.Description = st.Description
	pb.Events = st.Events
	pb.HttpUrlSpec = st.HttpUrlSpec
	pb.JobSpec = st.JobSpec
	pb.ModelName = st.ModelName
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createRegistryWebhookPb struct {
	Description string                 `json:"description,omitempty"`
	Events      []RegistryWebhookEvent `json:"events"`
	HttpUrlSpec *HttpUrlSpec           `json:"http_url_spec,omitempty"`
	JobSpec     *JobSpec               `json:"job_spec,omitempty"`
	ModelName   string                 `json:"model_name,omitempty"`
	Status      RegistryWebhookStatus  `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRegistryWebhookFromPb(pb *createRegistryWebhookPb) (*CreateRegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRegistryWebhook{}
	st.Description = pb.Description
	st.Events = pb.Events
	st.HttpUrlSpec = pb.HttpUrlSpec
	st.JobSpec = pb.JobSpec
	st.ModelName = pb.ModelName
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createRegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createRunToPb(st *CreateRun) (*createRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRunPb{}
	pb.ExperimentId = st.ExperimentId
	pb.RunName = st.RunName
	pb.StartTime = st.StartTime
	pb.Tags = st.Tags
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createRunPb struct {
	ExperimentId string   `json:"experiment_id,omitempty"`
	RunName      string   `json:"run_name,omitempty"`
	StartTime    int64    `json:"start_time,omitempty"`
	Tags         []RunTag `json:"tags,omitempty"`
	UserId       string   `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRunFromPb(pb *createRunPb) (*CreateRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRun{}
	st.ExperimentId = pb.ExperimentId
	st.RunName = pb.RunName
	st.StartTime = pb.StartTime
	st.Tags = pb.Tags
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createRunResponseToPb(st *CreateRunResponse) (*createRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRunResponsePb{}
	pb.Run = st.Run

	return pb, nil
}

type createRunResponsePb struct {
	Run *Run `json:"run,omitempty"`
}

func createRunResponseFromPb(pb *createRunResponsePb) (*CreateRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRunResponse{}
	st.Run = pb.Run

	return st, nil
}

func createTransitionRequestToPb(st *CreateTransitionRequest) (*createTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTransitionRequestPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createTransitionRequestPb struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name"`
	Stage   Stage  `json:"stage"`
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createTransitionRequestFromPb(pb *createTransitionRequestPb) (*CreateTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTransitionRequest{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createTransitionRequestResponseToPb(st *CreateTransitionRequestResponse) (*createTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTransitionRequestResponsePb{}
	pb.Request = st.Request

	return pb, nil
}

type createTransitionRequestResponsePb struct {
	Request *TransitionRequest `json:"request,omitempty"`
}

func createTransitionRequestResponseFromPb(pb *createTransitionRequestResponsePb) (*CreateTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTransitionRequestResponse{}
	st.Request = pb.Request

	return st, nil
}

func createWebhookResponseToPb(st *CreateWebhookResponse) (*createWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWebhookResponsePb{}
	pb.Webhook = st.Webhook

	return pb, nil
}

type createWebhookResponsePb struct {
	Webhook *RegistryWebhook `json:"webhook,omitempty"`
}

func createWebhookResponseFromPb(pb *createWebhookResponsePb) (*CreateWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWebhookResponse{}
	st.Webhook = pb.Webhook

	return st, nil
}

func datasetToPb(st *Dataset) (*datasetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &datasetPb{}
	pb.Digest = st.Digest
	pb.Name = st.Name
	pb.Profile = st.Profile
	pb.Schema = st.Schema
	pb.Source = st.Source
	pb.SourceType = st.SourceType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type datasetPb struct {
	Digest     string `json:"digest"`
	Name       string `json:"name"`
	Profile    string `json:"profile,omitempty"`
	Schema     string `json:"schema,omitempty"`
	Source     string `json:"source"`
	SourceType string `json:"source_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func datasetFromPb(pb *datasetPb) (*Dataset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dataset{}
	st.Digest = pb.Digest
	st.Name = pb.Name
	st.Profile = pb.Profile
	st.Schema = pb.Schema
	st.Source = pb.Source
	st.SourceType = pb.SourceType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *datasetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st datasetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func datasetInputToPb(st *DatasetInput) (*datasetInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &datasetInputPb{}
	pb.Dataset = st.Dataset
	pb.Tags = st.Tags

	return pb, nil
}

type datasetInputPb struct {
	Dataset Dataset    `json:"dataset"`
	Tags    []InputTag `json:"tags,omitempty"`
}

func datasetInputFromPb(pb *datasetInputPb) (*DatasetInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatasetInput{}
	st.Dataset = pb.Dataset
	st.Tags = pb.Tags

	return st, nil
}

func deleteCommentRequestToPb(st *DeleteCommentRequest) (*deleteCommentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCommentRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteCommentRequestPb struct {
	Id string `json:"-" url:"id"`
}

func deleteCommentRequestFromPb(pb *deleteCommentRequestPb) (*DeleteCommentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCommentRequest{}
	st.Id = pb.Id

	return st, nil
}

func deleteCommentResponseToPb(st *DeleteCommentResponse) (*deleteCommentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCommentResponsePb{}

	return pb, nil
}

type deleteCommentResponsePb struct {
}

func deleteCommentResponseFromPb(pb *deleteCommentResponsePb) (*DeleteCommentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCommentResponse{}

	return st, nil
}

func deleteExperimentToPb(st *DeleteExperiment) (*deleteExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExperimentPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

type deleteExperimentPb struct {
	ExperimentId string `json:"experiment_id"`
}

func deleteExperimentFromPb(pb *deleteExperimentPb) (*DeleteExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExperiment{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

func deleteExperimentResponseToPb(st *DeleteExperimentResponse) (*deleteExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExperimentResponsePb{}

	return pb, nil
}

type deleteExperimentResponsePb struct {
}

func deleteExperimentResponseFromPb(pb *deleteExperimentResponsePb) (*DeleteExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExperimentResponse{}

	return st, nil
}

func deleteLoggedModelRequestToPb(st *DeleteLoggedModelRequest) (*deleteLoggedModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteLoggedModelRequestPb{}
	pb.ModelId = st.ModelId

	return pb, nil
}

type deleteLoggedModelRequestPb struct {
	ModelId string `json:"-" url:"-"`
}

func deleteLoggedModelRequestFromPb(pb *deleteLoggedModelRequestPb) (*DeleteLoggedModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLoggedModelRequest{}
	st.ModelId = pb.ModelId

	return st, nil
}

func deleteLoggedModelResponseToPb(st *DeleteLoggedModelResponse) (*deleteLoggedModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteLoggedModelResponsePb{}

	return pb, nil
}

type deleteLoggedModelResponsePb struct {
}

func deleteLoggedModelResponseFromPb(pb *deleteLoggedModelResponsePb) (*DeleteLoggedModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLoggedModelResponse{}

	return st, nil
}

func deleteLoggedModelTagRequestToPb(st *DeleteLoggedModelTagRequest) (*deleteLoggedModelTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteLoggedModelTagRequestPb{}
	pb.ModelId = st.ModelId
	pb.TagKey = st.TagKey

	return pb, nil
}

type deleteLoggedModelTagRequestPb struct {
	ModelId string `json:"-" url:"-"`
	TagKey  string `json:"-" url:"-"`
}

func deleteLoggedModelTagRequestFromPb(pb *deleteLoggedModelTagRequestPb) (*DeleteLoggedModelTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLoggedModelTagRequest{}
	st.ModelId = pb.ModelId
	st.TagKey = pb.TagKey

	return st, nil
}

func deleteLoggedModelTagResponseToPb(st *DeleteLoggedModelTagResponse) (*deleteLoggedModelTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteLoggedModelTagResponsePb{}

	return pb, nil
}

type deleteLoggedModelTagResponsePb struct {
}

func deleteLoggedModelTagResponseFromPb(pb *deleteLoggedModelTagResponsePb) (*DeleteLoggedModelTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLoggedModelTagResponse{}

	return st, nil
}

func deleteModelRequestToPb(st *DeleteModelRequest) (*deleteModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteModelRequestPb struct {
	Name string `json:"-" url:"name"`
}

func deleteModelRequestFromPb(pb *deleteModelRequestPb) (*DeleteModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteModelResponseToPb(st *DeleteModelResponse) (*deleteModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelResponsePb{}

	return pb, nil
}

type deleteModelResponsePb struct {
}

func deleteModelResponseFromPb(pb *deleteModelResponsePb) (*DeleteModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelResponse{}

	return st, nil
}

func deleteModelTagRequestToPb(st *DeleteModelTagRequest) (*deleteModelTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelTagRequestPb{}
	pb.Key = st.Key
	pb.Name = st.Name

	return pb, nil
}

type deleteModelTagRequestPb struct {
	Key  string `json:"-" url:"key"`
	Name string `json:"-" url:"name"`
}

func deleteModelTagRequestFromPb(pb *deleteModelTagRequestPb) (*DeleteModelTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name

	return st, nil
}

func deleteModelTagResponseToPb(st *DeleteModelTagResponse) (*deleteModelTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelTagResponsePb{}

	return pb, nil
}

type deleteModelTagResponsePb struct {
}

func deleteModelTagResponseFromPb(pb *deleteModelTagResponsePb) (*DeleteModelTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelTagResponse{}

	return st, nil
}

func deleteModelVersionRequestToPb(st *DeleteModelVersionRequest) (*deleteModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionRequestPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

type deleteModelVersionRequestPb struct {
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

func deleteModelVersionRequestFromPb(pb *deleteModelVersionRequestPb) (*DeleteModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

func deleteModelVersionResponseToPb(st *DeleteModelVersionResponse) (*deleteModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionResponsePb{}

	return pb, nil
}

type deleteModelVersionResponsePb struct {
}

func deleteModelVersionResponseFromPb(pb *deleteModelVersionResponsePb) (*DeleteModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionResponse{}

	return st, nil
}

func deleteModelVersionTagRequestToPb(st *DeleteModelVersionTagRequest) (*deleteModelVersionTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionTagRequestPb{}
	pb.Key = st.Key
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

type deleteModelVersionTagRequestPb struct {
	Key     string `json:"-" url:"key"`
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

func deleteModelVersionTagRequestFromPb(pb *deleteModelVersionTagRequestPb) (*DeleteModelVersionTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

func deleteModelVersionTagResponseToPb(st *DeleteModelVersionTagResponse) (*deleteModelVersionTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionTagResponsePb{}

	return pb, nil
}

type deleteModelVersionTagResponsePb struct {
}

func deleteModelVersionTagResponseFromPb(pb *deleteModelVersionTagResponsePb) (*DeleteModelVersionTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionTagResponse{}

	return st, nil
}

func deleteOnlineStoreRequestToPb(st *DeleteOnlineStoreRequest) (*deleteOnlineStoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteOnlineStoreRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteOnlineStoreRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteOnlineStoreRequestFromPb(pb *deleteOnlineStoreRequestPb) (*DeleteOnlineStoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteOnlineStoreRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteOnlineStoreResponseToPb(st *DeleteOnlineStoreResponse) (*deleteOnlineStoreResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteOnlineStoreResponsePb{}

	return pb, nil
}

type deleteOnlineStoreResponsePb struct {
}

func deleteOnlineStoreResponseFromPb(pb *deleteOnlineStoreResponsePb) (*DeleteOnlineStoreResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteOnlineStoreResponse{}

	return st, nil
}

func deleteRunToPb(st *DeleteRun) (*deleteRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

type deleteRunPb struct {
	RunId string `json:"run_id"`
}

func deleteRunFromPb(pb *deleteRunPb) (*DeleteRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRun{}
	st.RunId = pb.RunId

	return st, nil
}

func deleteRunResponseToPb(st *DeleteRunResponse) (*deleteRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunResponsePb{}

	return pb, nil
}

type deleteRunResponsePb struct {
}

func deleteRunResponseFromPb(pb *deleteRunResponsePb) (*DeleteRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRunResponse{}

	return st, nil
}

func deleteRunsToPb(st *DeleteRuns) (*deleteRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunsPb{}
	pb.ExperimentId = st.ExperimentId
	pb.MaxRuns = st.MaxRuns
	pb.MaxTimestampMillis = st.MaxTimestampMillis

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteRunsPb struct {
	ExperimentId       string `json:"experiment_id"`
	MaxRuns            int    `json:"max_runs,omitempty"`
	MaxTimestampMillis int64  `json:"max_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteRunsFromPb(pb *deleteRunsPb) (*DeleteRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRuns{}
	st.ExperimentId = pb.ExperimentId
	st.MaxRuns = pb.MaxRuns
	st.MaxTimestampMillis = pb.MaxTimestampMillis

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteRunsResponseToPb(st *DeleteRunsResponse) (*deleteRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunsResponsePb{}
	pb.RunsDeleted = st.RunsDeleted

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteRunsResponsePb struct {
	RunsDeleted int `json:"runs_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteRunsResponseFromPb(pb *deleteRunsResponsePb) (*DeleteRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRunsResponse{}
	st.RunsDeleted = pb.RunsDeleted

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteTagToPb(st *DeleteTag) (*deleteTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTagPb{}
	pb.Key = st.Key
	pb.RunId = st.RunId

	return pb, nil
}

type deleteTagPb struct {
	Key   string `json:"key"`
	RunId string `json:"run_id"`
}

func deleteTagFromPb(pb *deleteTagPb) (*DeleteTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTag{}
	st.Key = pb.Key
	st.RunId = pb.RunId

	return st, nil
}

func deleteTagResponseToPb(st *DeleteTagResponse) (*deleteTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTagResponsePb{}

	return pb, nil
}

type deleteTagResponsePb struct {
}

func deleteTagResponseFromPb(pb *deleteTagResponsePb) (*DeleteTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTagResponse{}

	return st, nil
}

func deleteTransitionRequestRequestToPb(st *DeleteTransitionRequestRequest) (*deleteTransitionRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTransitionRequestRequestPb{}
	pb.Comment = st.Comment
	pb.Creator = st.Creator
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteTransitionRequestRequestPb struct {
	Comment string                       `json:"-" url:"comment,omitempty"`
	Creator string                       `json:"-" url:"creator"`
	Name    string                       `json:"-" url:"name"`
	Stage   DeleteTransitionRequestStage `json:"-" url:"stage"`
	Version string                       `json:"-" url:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteTransitionRequestRequestFromPb(pb *deleteTransitionRequestRequestPb) (*DeleteTransitionRequestRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTransitionRequestRequest{}
	st.Comment = pb.Comment
	st.Creator = pb.Creator
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteTransitionRequestRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteTransitionRequestRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteTransitionRequestResponseToPb(st *DeleteTransitionRequestResponse) (*deleteTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTransitionRequestResponsePb{}

	return pb, nil
}

type deleteTransitionRequestResponsePb struct {
}

func deleteTransitionRequestResponseFromPb(pb *deleteTransitionRequestResponsePb) (*DeleteTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTransitionRequestResponse{}

	return st, nil
}

func deleteWebhookRequestToPb(st *DeleteWebhookRequest) (*deleteWebhookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWebhookRequestPb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteWebhookRequestPb struct {
	Id string `json:"-" url:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteWebhookRequestFromPb(pb *deleteWebhookRequestPb) (*DeleteWebhookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWebhookRequest{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteWebhookRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteWebhookRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteWebhookResponseToPb(st *DeleteWebhookResponse) (*deleteWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWebhookResponsePb{}

	return pb, nil
}

type deleteWebhookResponsePb struct {
}

func deleteWebhookResponseFromPb(pb *deleteWebhookResponsePb) (*DeleteWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWebhookResponse{}

	return st, nil
}

func experimentToPb(st *Experiment) (*experimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPb{}
	pb.ArtifactLocation = st.ArtifactLocation
	pb.CreationTime = st.CreationTime
	pb.ExperimentId = st.ExperimentId
	pb.LastUpdateTime = st.LastUpdateTime
	pb.LifecycleStage = st.LifecycleStage
	pb.Name = st.Name
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type experimentPb struct {
	ArtifactLocation string          `json:"artifact_location,omitempty"`
	CreationTime     int64           `json:"creation_time,omitempty"`
	ExperimentId     string          `json:"experiment_id,omitempty"`
	LastUpdateTime   int64           `json:"last_update_time,omitempty"`
	LifecycleStage   string          `json:"lifecycle_stage,omitempty"`
	Name             string          `json:"name,omitempty"`
	Tags             []ExperimentTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentFromPb(pb *experimentPb) (*Experiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Experiment{}
	st.ArtifactLocation = pb.ArtifactLocation
	st.CreationTime = pb.CreationTime
	st.ExperimentId = pb.ExperimentId
	st.LastUpdateTime = pb.LastUpdateTime
	st.LifecycleStage = pb.LifecycleStage
	st.Name = pb.Name
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func experimentAccessControlRequestToPb(st *ExperimentAccessControlRequest) (*experimentAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	pb.PermissionLevel = st.PermissionLevel
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type experimentAccessControlRequestPb struct {
	GroupName            string                    `json:"group_name,omitempty"`
	PermissionLevel      ExperimentPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                    `json:"service_principal_name,omitempty"`
	UserName             string                    `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentAccessControlRequestFromPb(pb *experimentAccessControlRequestPb) (*ExperimentAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func experimentAccessControlResponseToPb(st *ExperimentAccessControlResponse) (*experimentAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type experimentAccessControlResponsePb struct {
	AllPermissions       []ExperimentPermission `json:"all_permissions,omitempty"`
	DisplayName          string                 `json:"display_name,omitempty"`
	GroupName            string                 `json:"group_name,omitempty"`
	ServicePrincipalName string                 `json:"service_principal_name,omitempty"`
	UserName             string                 `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentAccessControlResponseFromPb(pb *experimentAccessControlResponsePb) (*ExperimentAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func experimentPermissionToPb(st *ExperimentPermission) (*experimentPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type experimentPermissionPb struct {
	Inherited           bool                      `json:"inherited,omitempty"`
	InheritedFromObject []string                  `json:"inherited_from_object,omitempty"`
	PermissionLevel     ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentPermissionFromPb(pb *experimentPermissionPb) (*ExperimentPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func experimentPermissionsToPb(st *ExperimentPermissions) (*experimentPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type experimentPermissionsPb struct {
	AccessControlList []ExperimentAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                            `json:"object_id,omitempty"`
	ObjectType        string                            `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentPermissionsFromPb(pb *experimentPermissionsPb) (*ExperimentPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func experimentPermissionsDescriptionToPb(st *ExperimentPermissionsDescription) (*experimentPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPermissionsDescriptionPb{}
	pb.Description = st.Description
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type experimentPermissionsDescriptionPb struct {
	Description     string                    `json:"description,omitempty"`
	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentPermissionsDescriptionFromPb(pb *experimentPermissionsDescriptionPb) (*ExperimentPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func experimentPermissionsRequestToPb(st *ExperimentPermissionsRequest) (*experimentPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

type experimentPermissionsRequestPb struct {
	AccessControlList []ExperimentAccessControlRequest `json:"access_control_list,omitempty"`
	ExperimentId      string                           `json:"-" url:"-"`
}

func experimentPermissionsRequestFromPb(pb *experimentPermissionsRequestPb) (*ExperimentPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

func experimentTagToPb(st *ExperimentTag) (*experimentTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type experimentTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentTagFromPb(pb *experimentTagPb) (*ExperimentTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func fileInfoToPb(st *FileInfo) (*fileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileInfoPb{}
	pb.FileSize = st.FileSize
	pb.IsDir = st.IsDir
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type fileInfoPb struct {
	FileSize int64  `json:"file_size,omitempty"`
	IsDir    bool   `json:"is_dir,omitempty"`
	Path     string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileInfoFromPb(pb *fileInfoPb) (*FileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileInfo{}
	st.FileSize = pb.FileSize
	st.IsDir = pb.IsDir
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func finalizeLoggedModelRequestToPb(st *FinalizeLoggedModelRequest) (*finalizeLoggedModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &finalizeLoggedModelRequestPb{}
	pb.ModelId = st.ModelId
	pb.Status = st.Status

	return pb, nil
}

type finalizeLoggedModelRequestPb struct {
	ModelId string            `json:"-" url:"-"`
	Status  LoggedModelStatus `json:"status"`
}

func finalizeLoggedModelRequestFromPb(pb *finalizeLoggedModelRequestPb) (*FinalizeLoggedModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FinalizeLoggedModelRequest{}
	st.ModelId = pb.ModelId
	st.Status = pb.Status

	return st, nil
}

func finalizeLoggedModelResponseToPb(st *FinalizeLoggedModelResponse) (*finalizeLoggedModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &finalizeLoggedModelResponsePb{}
	pb.Model = st.Model

	return pb, nil
}

type finalizeLoggedModelResponsePb struct {
	Model *LoggedModel `json:"model,omitempty"`
}

func finalizeLoggedModelResponseFromPb(pb *finalizeLoggedModelResponsePb) (*FinalizeLoggedModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FinalizeLoggedModelResponse{}
	st.Model = pb.Model

	return st, nil
}

func forecastingExperimentToPb(st *ForecastingExperiment) (*forecastingExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forecastingExperimentPb{}
	pb.ExperimentId = st.ExperimentId
	pb.ExperimentPageUrl = st.ExperimentPageUrl
	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type forecastingExperimentPb struct {
	ExperimentId      string                     `json:"experiment_id,omitempty"`
	ExperimentPageUrl string                     `json:"experiment_page_url,omitempty"`
	State             ForecastingExperimentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func forecastingExperimentFromPb(pb *forecastingExperimentPb) (*ForecastingExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForecastingExperiment{}
	st.ExperimentId = pb.ExperimentId
	st.ExperimentPageUrl = pb.ExperimentPageUrl
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *forecastingExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st forecastingExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getByNameRequestToPb(st *GetByNameRequest) (*getByNameRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getByNameRequestPb{}
	pb.ExperimentName = st.ExperimentName

	return pb, nil
}

type getByNameRequestPb struct {
	ExperimentName string `json:"-" url:"experiment_name"`
}

func getByNameRequestFromPb(pb *getByNameRequestPb) (*GetByNameRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetByNameRequest{}
	st.ExperimentName = pb.ExperimentName

	return st, nil
}

func getExperimentByNameResponseToPb(st *GetExperimentByNameResponse) (*getExperimentByNameResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentByNameResponsePb{}
	pb.Experiment = st.Experiment

	return pb, nil
}

type getExperimentByNameResponsePb struct {
	Experiment *Experiment `json:"experiment,omitempty"`
}

func getExperimentByNameResponseFromPb(pb *getExperimentByNameResponsePb) (*GetExperimentByNameResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentByNameResponse{}
	st.Experiment = pb.Experiment

	return st, nil
}

func getExperimentPermissionLevelsRequestToPb(st *GetExperimentPermissionLevelsRequest) (*getExperimentPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentPermissionLevelsRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

type getExperimentPermissionLevelsRequestPb struct {
	ExperimentId string `json:"-" url:"-"`
}

func getExperimentPermissionLevelsRequestFromPb(pb *getExperimentPermissionLevelsRequestPb) (*GetExperimentPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentPermissionLevelsRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

func getExperimentPermissionLevelsResponseToPb(st *GetExperimentPermissionLevelsResponse) (*getExperimentPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getExperimentPermissionLevelsResponsePb struct {
	PermissionLevels []ExperimentPermissionsDescription `json:"permission_levels,omitempty"`
}

func getExperimentPermissionLevelsResponseFromPb(pb *getExperimentPermissionLevelsResponsePb) (*GetExperimentPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getExperimentPermissionsRequestToPb(st *GetExperimentPermissionsRequest) (*getExperimentPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentPermissionsRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

type getExperimentPermissionsRequestPb struct {
	ExperimentId string `json:"-" url:"-"`
}

func getExperimentPermissionsRequestFromPb(pb *getExperimentPermissionsRequestPb) (*GetExperimentPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentPermissionsRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

func getExperimentRequestToPb(st *GetExperimentRequest) (*getExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

type getExperimentRequestPb struct {
	ExperimentId string `json:"-" url:"experiment_id"`
}

func getExperimentRequestFromPb(pb *getExperimentRequestPb) (*GetExperimentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

func getExperimentResponseToPb(st *GetExperimentResponse) (*getExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentResponsePb{}
	pb.Experiment = st.Experiment

	return pb, nil
}

type getExperimentResponsePb struct {
	Experiment *Experiment `json:"experiment,omitempty"`
}

func getExperimentResponseFromPb(pb *getExperimentResponsePb) (*GetExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentResponse{}
	st.Experiment = pb.Experiment

	return st, nil
}

func getForecastingExperimentRequestToPb(st *GetForecastingExperimentRequest) (*getForecastingExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getForecastingExperimentRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

type getForecastingExperimentRequestPb struct {
	ExperimentId string `json:"-" url:"-"`
}

func getForecastingExperimentRequestFromPb(pb *getForecastingExperimentRequestPb) (*GetForecastingExperimentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetForecastingExperimentRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

func getHistoryRequestToPb(st *GetHistoryRequest) (*getHistoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getHistoryRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.MetricKey = st.MetricKey
	pb.PageToken = st.PageToken
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getHistoryRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	MetricKey  string `json:"-" url:"metric_key"`
	PageToken  string `json:"-" url:"page_token,omitempty"`
	RunId      string `json:"-" url:"run_id,omitempty"`
	RunUuid    string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getHistoryRequestFromPb(pb *getHistoryRequestPb) (*GetHistoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetHistoryRequest{}
	st.MaxResults = pb.MaxResults
	st.MetricKey = pb.MetricKey
	st.PageToken = pb.PageToken
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getHistoryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getHistoryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getLatestVersionsRequestToPb(st *GetLatestVersionsRequest) (*getLatestVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLatestVersionsRequestPb{}
	pb.Name = st.Name
	pb.Stages = st.Stages

	return pb, nil
}

type getLatestVersionsRequestPb struct {
	Name   string   `json:"name"`
	Stages []string `json:"stages,omitempty"`
}

func getLatestVersionsRequestFromPb(pb *getLatestVersionsRequestPb) (*GetLatestVersionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLatestVersionsRequest{}
	st.Name = pb.Name
	st.Stages = pb.Stages

	return st, nil
}

func getLatestVersionsResponseToPb(st *GetLatestVersionsResponse) (*getLatestVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLatestVersionsResponsePb{}
	pb.ModelVersions = st.ModelVersions

	return pb, nil
}

type getLatestVersionsResponsePb struct {
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
}

func getLatestVersionsResponseFromPb(pb *getLatestVersionsResponsePb) (*GetLatestVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLatestVersionsResponse{}
	st.ModelVersions = pb.ModelVersions

	return st, nil
}

func getLoggedModelRequestToPb(st *GetLoggedModelRequest) (*getLoggedModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLoggedModelRequestPb{}
	pb.ModelId = st.ModelId

	return pb, nil
}

type getLoggedModelRequestPb struct {
	ModelId string `json:"-" url:"-"`
}

func getLoggedModelRequestFromPb(pb *getLoggedModelRequestPb) (*GetLoggedModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLoggedModelRequest{}
	st.ModelId = pb.ModelId

	return st, nil
}

func getLoggedModelResponseToPb(st *GetLoggedModelResponse) (*getLoggedModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLoggedModelResponsePb{}
	pb.Model = st.Model

	return pb, nil
}

type getLoggedModelResponsePb struct {
	Model *LoggedModel `json:"model,omitempty"`
}

func getLoggedModelResponseFromPb(pb *getLoggedModelResponsePb) (*GetLoggedModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLoggedModelResponse{}
	st.Model = pb.Model

	return st, nil
}

func getMetricHistoryResponseToPb(st *GetMetricHistoryResponse) (*getMetricHistoryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetricHistoryResponsePb{}
	pb.Metrics = st.Metrics
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getMetricHistoryResponsePb struct {
	Metrics       []Metric `json:"metrics,omitempty"`
	NextPageToken string   `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getMetricHistoryResponseFromPb(pb *getMetricHistoryResponsePb) (*GetMetricHistoryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetricHistoryResponse{}
	st.Metrics = pb.Metrics
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getMetricHistoryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getMetricHistoryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getModelRequestToPb(st *GetModelRequest) (*getModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getModelRequestPb struct {
	Name string `json:"-" url:"name"`
}

func getModelRequestFromPb(pb *getModelRequestPb) (*GetModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelRequest{}
	st.Name = pb.Name

	return st, nil
}

func getModelResponseToPb(st *GetModelResponse) (*getModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelResponsePb{}
	pb.RegisteredModelDatabricks = st.RegisteredModelDatabricks

	return pb, nil
}

type getModelResponsePb struct {
	RegisteredModelDatabricks *ModelDatabricks `json:"registered_model_databricks,omitempty"`
}

func getModelResponseFromPb(pb *getModelResponsePb) (*GetModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelResponse{}
	st.RegisteredModelDatabricks = pb.RegisteredModelDatabricks

	return st, nil
}

func getModelVersionDownloadUriRequestToPb(st *GetModelVersionDownloadUriRequest) (*getModelVersionDownloadUriRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionDownloadUriRequestPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

type getModelVersionDownloadUriRequestPb struct {
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

func getModelVersionDownloadUriRequestFromPb(pb *getModelVersionDownloadUriRequestPb) (*GetModelVersionDownloadUriRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionDownloadUriRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

func getModelVersionDownloadUriResponseToPb(st *GetModelVersionDownloadUriResponse) (*getModelVersionDownloadUriResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionDownloadUriResponsePb{}
	pb.ArtifactUri = st.ArtifactUri

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getModelVersionDownloadUriResponsePb struct {
	ArtifactUri string `json:"artifact_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getModelVersionDownloadUriResponseFromPb(pb *getModelVersionDownloadUriResponsePb) (*GetModelVersionDownloadUriResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionDownloadUriResponse{}
	st.ArtifactUri = pb.ArtifactUri

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getModelVersionDownloadUriResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getModelVersionDownloadUriResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getModelVersionRequestToPb(st *GetModelVersionRequest) (*getModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionRequestPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

type getModelVersionRequestPb struct {
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

func getModelVersionRequestFromPb(pb *getModelVersionRequestPb) (*GetModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

func getModelVersionResponseToPb(st *GetModelVersionResponse) (*getModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionResponsePb{}
	pb.ModelVersion = st.ModelVersion

	return pb, nil
}

type getModelVersionResponsePb struct {
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

func getModelVersionResponseFromPb(pb *getModelVersionResponsePb) (*GetModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionResponse{}
	st.ModelVersion = pb.ModelVersion

	return st, nil
}

func getOnlineStoreRequestToPb(st *GetOnlineStoreRequest) (*getOnlineStoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getOnlineStoreRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getOnlineStoreRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getOnlineStoreRequestFromPb(pb *getOnlineStoreRequestPb) (*GetOnlineStoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOnlineStoreRequest{}
	st.Name = pb.Name

	return st, nil
}

func getRegisteredModelPermissionLevelsRequestToPb(st *GetRegisteredModelPermissionLevelsRequest) (*getRegisteredModelPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelPermissionLevelsRequestPb{}
	pb.RegisteredModelId = st.RegisteredModelId

	return pb, nil
}

type getRegisteredModelPermissionLevelsRequestPb struct {
	RegisteredModelId string `json:"-" url:"-"`
}

func getRegisteredModelPermissionLevelsRequestFromPb(pb *getRegisteredModelPermissionLevelsRequestPb) (*GetRegisteredModelPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelPermissionLevelsRequest{}
	st.RegisteredModelId = pb.RegisteredModelId

	return st, nil
}

func getRegisteredModelPermissionLevelsResponseToPb(st *GetRegisteredModelPermissionLevelsResponse) (*getRegisteredModelPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getRegisteredModelPermissionLevelsResponsePb struct {
	PermissionLevels []RegisteredModelPermissionsDescription `json:"permission_levels,omitempty"`
}

func getRegisteredModelPermissionLevelsResponseFromPb(pb *getRegisteredModelPermissionLevelsResponsePb) (*GetRegisteredModelPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getRegisteredModelPermissionsRequestToPb(st *GetRegisteredModelPermissionsRequest) (*getRegisteredModelPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelPermissionsRequestPb{}
	pb.RegisteredModelId = st.RegisteredModelId

	return pb, nil
}

type getRegisteredModelPermissionsRequestPb struct {
	RegisteredModelId string `json:"-" url:"-"`
}

func getRegisteredModelPermissionsRequestFromPb(pb *getRegisteredModelPermissionsRequestPb) (*GetRegisteredModelPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelPermissionsRequest{}
	st.RegisteredModelId = pb.RegisteredModelId

	return st, nil
}

func getRunRequestToPb(st *GetRunRequest) (*getRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRunRequestPb{}
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getRunRequestPb struct {
	RunId   string `json:"-" url:"run_id"`
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRunRequestFromPb(pb *getRunRequestPb) (*GetRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunRequest{}
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRunRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRunRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getRunResponseToPb(st *GetRunResponse) (*getRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRunResponsePb{}
	pb.Run = st.Run

	return pb, nil
}

type getRunResponsePb struct {
	Run *Run `json:"run,omitempty"`
}

func getRunResponseFromPb(pb *getRunResponsePb) (*GetRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunResponse{}
	st.Run = pb.Run

	return st, nil
}

func httpUrlSpecToPb(st *HttpUrlSpec) (*httpUrlSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpUrlSpecPb{}
	pb.Authorization = st.Authorization
	pb.EnableSslVerification = st.EnableSslVerification
	pb.Secret = st.Secret
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type httpUrlSpecPb struct {
	Authorization         string `json:"authorization,omitempty"`
	EnableSslVerification bool   `json:"enable_ssl_verification,omitempty"`
	Secret                string `json:"secret,omitempty"`
	Url                   string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func httpUrlSpecFromPb(pb *httpUrlSpecPb) (*HttpUrlSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpUrlSpec{}
	st.Authorization = pb.Authorization
	st.EnableSslVerification = pb.EnableSslVerification
	st.Secret = pb.Secret
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *httpUrlSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st httpUrlSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func httpUrlSpecWithoutSecretToPb(st *HttpUrlSpecWithoutSecret) (*httpUrlSpecWithoutSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpUrlSpecWithoutSecretPb{}
	pb.EnableSslVerification = st.EnableSslVerification
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type httpUrlSpecWithoutSecretPb struct {
	EnableSslVerification bool   `json:"enable_ssl_verification,omitempty"`
	Url                   string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func httpUrlSpecWithoutSecretFromPb(pb *httpUrlSpecWithoutSecretPb) (*HttpUrlSpecWithoutSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpUrlSpecWithoutSecret{}
	st.EnableSslVerification = pb.EnableSslVerification
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *httpUrlSpecWithoutSecretPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st httpUrlSpecWithoutSecretPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func inputTagToPb(st *InputTag) (*inputTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &inputTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	return pb, nil
}

type inputTagPb struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func inputTagFromPb(pb *inputTagPb) (*InputTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InputTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	return st, nil
}

func jobSpecToPb(st *JobSpec) (*jobSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobSpecPb{}
	pb.AccessToken = st.AccessToken
	pb.JobId = st.JobId
	pb.WorkspaceUrl = st.WorkspaceUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobSpecPb struct {
	AccessToken  string `json:"access_token"`
	JobId        string `json:"job_id"`
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobSpecFromPb(pb *jobSpecPb) (*JobSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSpec{}
	st.AccessToken = pb.AccessToken
	st.JobId = pb.JobId
	st.WorkspaceUrl = pb.WorkspaceUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobSpecWithoutSecretToPb(st *JobSpecWithoutSecret) (*jobSpecWithoutSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobSpecWithoutSecretPb{}
	pb.JobId = st.JobId
	pb.WorkspaceUrl = st.WorkspaceUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobSpecWithoutSecretPb struct {
	JobId        string `json:"job_id,omitempty"`
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobSpecWithoutSecretFromPb(pb *jobSpecWithoutSecretPb) (*JobSpecWithoutSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSpecWithoutSecret{}
	st.JobId = pb.JobId
	st.WorkspaceUrl = pb.WorkspaceUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobSpecWithoutSecretPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobSpecWithoutSecretPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listArtifactsRequestToPb(st *ListArtifactsRequest) (*listArtifactsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listArtifactsRequestPb{}
	pb.PageToken = st.PageToken
	pb.Path = st.Path
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listArtifactsRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`
	Path      string `json:"-" url:"path,omitempty"`
	RunId     string `json:"-" url:"run_id,omitempty"`
	RunUuid   string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listArtifactsRequestFromPb(pb *listArtifactsRequestPb) (*ListArtifactsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListArtifactsRequest{}
	st.PageToken = pb.PageToken
	st.Path = pb.Path
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listArtifactsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listArtifactsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listArtifactsResponseToPb(st *ListArtifactsResponse) (*listArtifactsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listArtifactsResponsePb{}
	pb.Files = st.Files
	pb.NextPageToken = st.NextPageToken
	pb.RootUri = st.RootUri

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listArtifactsResponsePb struct {
	Files         []FileInfo `json:"files,omitempty"`
	NextPageToken string     `json:"next_page_token,omitempty"`
	RootUri       string     `json:"root_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listArtifactsResponseFromPb(pb *listArtifactsResponsePb) (*ListArtifactsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListArtifactsResponse{}
	st.Files = pb.Files
	st.NextPageToken = pb.NextPageToken
	st.RootUri = pb.RootUri

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listArtifactsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listArtifactsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExperimentsRequestToPb(st *ListExperimentsRequest) (*listExperimentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExperimentsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.ViewType = st.ViewType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExperimentsRequestPb struct {
	MaxResults int64    `json:"-" url:"max_results,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`
	ViewType   ViewType `json:"-" url:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExperimentsRequestFromPb(pb *listExperimentsRequestPb) (*ListExperimentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExperimentsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.ViewType = pb.ViewType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExperimentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExperimentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExperimentsResponseToPb(st *ListExperimentsResponse) (*listExperimentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExperimentsResponsePb{}
	pb.Experiments = st.Experiments
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExperimentsResponsePb struct {
	Experiments   []Experiment `json:"experiments,omitempty"`
	NextPageToken string       `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExperimentsResponseFromPb(pb *listExperimentsResponsePb) (*ListExperimentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExperimentsResponse{}
	st.Experiments = pb.Experiments
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExperimentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExperimentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listModelsRequestToPb(st *ListModelsRequest) (*listModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listModelsRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listModelsRequestFromPb(pb *listModelsRequestPb) (*ListModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listModelsResponseToPb(st *ListModelsResponse) (*listModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.RegisteredModels = st.RegisteredModels

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listModelsResponsePb struct {
	NextPageToken    string  `json:"next_page_token,omitempty"`
	RegisteredModels []Model `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listModelsResponseFromPb(pb *listModelsResponsePb) (*ListModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.RegisteredModels = pb.RegisteredModels

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listOnlineStoresRequestToPb(st *ListOnlineStoresRequest) (*listOnlineStoresRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listOnlineStoresRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listOnlineStoresRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listOnlineStoresRequestFromPb(pb *listOnlineStoresRequestPb) (*ListOnlineStoresRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListOnlineStoresRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listOnlineStoresRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listOnlineStoresRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listOnlineStoresResponseToPb(st *ListOnlineStoresResponse) (*listOnlineStoresResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listOnlineStoresResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.OnlineStores = st.OnlineStores

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listOnlineStoresResponsePb struct {
	NextPageToken string        `json:"next_page_token,omitempty"`
	OnlineStores  []OnlineStore `json:"online_stores,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listOnlineStoresResponseFromPb(pb *listOnlineStoresResponsePb) (*ListOnlineStoresResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListOnlineStoresResponse{}
	st.NextPageToken = pb.NextPageToken
	st.OnlineStores = pb.OnlineStores

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listOnlineStoresResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listOnlineStoresResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listRegistryWebhooksToPb(st *ListRegistryWebhooks) (*listRegistryWebhooksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRegistryWebhooksPb{}
	pb.NextPageToken = st.NextPageToken
	pb.Webhooks = st.Webhooks

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listRegistryWebhooksPb struct {
	NextPageToken string            `json:"next_page_token,omitempty"`
	Webhooks      []RegistryWebhook `json:"webhooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRegistryWebhooksFromPb(pb *listRegistryWebhooksPb) (*ListRegistryWebhooks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRegistryWebhooks{}
	st.NextPageToken = pb.NextPageToken
	st.Webhooks = pb.Webhooks

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRegistryWebhooksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRegistryWebhooksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listTransitionRequestsRequestToPb(st *ListTransitionRequestsRequest) (*listTransitionRequestsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTransitionRequestsRequestPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

type listTransitionRequestsRequestPb struct {
	Name    string `json:"-" url:"name"`
	Version string `json:"-" url:"version"`
}

func listTransitionRequestsRequestFromPb(pb *listTransitionRequestsRequestPb) (*ListTransitionRequestsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTransitionRequestsRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

func listTransitionRequestsResponseToPb(st *ListTransitionRequestsResponse) (*listTransitionRequestsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTransitionRequestsResponsePb{}
	pb.Requests = st.Requests

	return pb, nil
}

type listTransitionRequestsResponsePb struct {
	Requests []Activity `json:"requests,omitempty"`
}

func listTransitionRequestsResponseFromPb(pb *listTransitionRequestsResponsePb) (*ListTransitionRequestsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTransitionRequestsResponse{}
	st.Requests = pb.Requests

	return st, nil
}

func listWebhooksRequestToPb(st *ListWebhooksRequest) (*listWebhooksRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWebhooksRequestPb{}
	pb.Events = st.Events
	pb.ModelName = st.ModelName
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listWebhooksRequestPb struct {
	Events    []RegistryWebhookEvent `json:"-" url:"events,omitempty"`
	ModelName string                 `json:"-" url:"model_name,omitempty"`
	PageToken string                 `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listWebhooksRequestFromPb(pb *listWebhooksRequestPb) (*ListWebhooksRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWebhooksRequest{}
	st.Events = pb.Events
	st.ModelName = pb.ModelName
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listWebhooksRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listWebhooksRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func logBatchToPb(st *LogBatch) (*logBatchPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logBatchPb{}
	pb.Metrics = st.Metrics
	pb.Params = st.Params
	pb.RunId = st.RunId
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type logBatchPb struct {
	Metrics []Metric `json:"metrics,omitempty"`
	Params  []Param  `json:"params,omitempty"`
	RunId   string   `json:"run_id,omitempty"`
	Tags    []RunTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logBatchFromPb(pb *logBatchPb) (*LogBatch, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogBatch{}
	st.Metrics = pb.Metrics
	st.Params = pb.Params
	st.RunId = pb.RunId
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logBatchPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logBatchPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func logBatchResponseToPb(st *LogBatchResponse) (*logBatchResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logBatchResponsePb{}

	return pb, nil
}

type logBatchResponsePb struct {
}

func logBatchResponseFromPb(pb *logBatchResponsePb) (*LogBatchResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogBatchResponse{}

	return st, nil
}

func logInputsToPb(st *LogInputs) (*logInputsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logInputsPb{}
	pb.Datasets = st.Datasets
	pb.Models = st.Models
	pb.RunId = st.RunId

	return pb, nil
}

type logInputsPb struct {
	Datasets []DatasetInput `json:"datasets,omitempty"`
	Models   []ModelInput   `json:"models,omitempty"`
	RunId    string         `json:"run_id"`
}

func logInputsFromPb(pb *logInputsPb) (*LogInputs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogInputs{}
	st.Datasets = pb.Datasets
	st.Models = pb.Models
	st.RunId = pb.RunId

	return st, nil
}

func logInputsResponseToPb(st *LogInputsResponse) (*logInputsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logInputsResponsePb{}

	return pb, nil
}

type logInputsResponsePb struct {
}

func logInputsResponseFromPb(pb *logInputsResponsePb) (*LogInputsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogInputsResponse{}

	return st, nil
}

func logLoggedModelParamsRequestToPb(st *LogLoggedModelParamsRequest) (*logLoggedModelParamsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logLoggedModelParamsRequestPb{}
	pb.ModelId = st.ModelId
	pb.Params = st.Params

	return pb, nil
}

type logLoggedModelParamsRequestPb struct {
	ModelId string                 `json:"-" url:"-"`
	Params  []LoggedModelParameter `json:"params,omitempty"`
}

func logLoggedModelParamsRequestFromPb(pb *logLoggedModelParamsRequestPb) (*LogLoggedModelParamsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogLoggedModelParamsRequest{}
	st.ModelId = pb.ModelId
	st.Params = pb.Params

	return st, nil
}

func logLoggedModelParamsRequestResponseToPb(st *LogLoggedModelParamsRequestResponse) (*logLoggedModelParamsRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logLoggedModelParamsRequestResponsePb{}

	return pb, nil
}

type logLoggedModelParamsRequestResponsePb struct {
}

func logLoggedModelParamsRequestResponseFromPb(pb *logLoggedModelParamsRequestResponsePb) (*LogLoggedModelParamsRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogLoggedModelParamsRequestResponse{}

	return st, nil
}

func logMetricToPb(st *LogMetric) (*logMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logMetricPb{}
	pb.DatasetDigest = st.DatasetDigest
	pb.DatasetName = st.DatasetName
	pb.Key = st.Key
	pb.ModelId = st.ModelId
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid
	pb.Step = st.Step
	pb.Timestamp = st.Timestamp
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type logMetricPb struct {
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

func logMetricFromPb(pb *logMetricPb) (*LogMetric, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogMetric{}
	st.DatasetDigest = pb.DatasetDigest
	st.DatasetName = pb.DatasetName
	st.Key = pb.Key
	st.ModelId = pb.ModelId
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid
	st.Step = pb.Step
	st.Timestamp = pb.Timestamp
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logMetricPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logMetricPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func logMetricResponseToPb(st *LogMetricResponse) (*logMetricResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logMetricResponsePb{}

	return pb, nil
}

type logMetricResponsePb struct {
}

func logMetricResponseFromPb(pb *logMetricResponsePb) (*LogMetricResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogMetricResponse{}

	return st, nil
}

func logModelToPb(st *LogModel) (*logModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logModelPb{}
	pb.ModelJson = st.ModelJson
	pb.RunId = st.RunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type logModelPb struct {
	ModelJson string `json:"model_json,omitempty"`
	RunId     string `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logModelFromPb(pb *logModelPb) (*LogModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogModel{}
	st.ModelJson = pb.ModelJson
	st.RunId = pb.RunId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func logModelResponseToPb(st *LogModelResponse) (*logModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logModelResponsePb{}

	return pb, nil
}

type logModelResponsePb struct {
}

func logModelResponseFromPb(pb *logModelResponsePb) (*LogModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogModelResponse{}

	return st, nil
}

func logOutputsRequestToPb(st *LogOutputsRequest) (*logOutputsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logOutputsRequestPb{}
	pb.Models = st.Models
	pb.RunId = st.RunId

	return pb, nil
}

type logOutputsRequestPb struct {
	Models []ModelOutput `json:"models,omitempty"`
	RunId  string        `json:"run_id"`
}

func logOutputsRequestFromPb(pb *logOutputsRequestPb) (*LogOutputsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogOutputsRequest{}
	st.Models = pb.Models
	st.RunId = pb.RunId

	return st, nil
}

func logOutputsResponseToPb(st *LogOutputsResponse) (*logOutputsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logOutputsResponsePb{}

	return pb, nil
}

type logOutputsResponsePb struct {
}

func logOutputsResponseFromPb(pb *logOutputsResponsePb) (*LogOutputsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogOutputsResponse{}

	return st, nil
}

func logParamToPb(st *LogParam) (*logParamPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logParamPb{}
	pb.Key = st.Key
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type logParamPb struct {
	Key     string `json:"key"`
	RunId   string `json:"run_id,omitempty"`
	RunUuid string `json:"run_uuid,omitempty"`
	Value   string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logParamFromPb(pb *logParamPb) (*LogParam, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogParam{}
	st.Key = pb.Key
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logParamPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logParamPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func logParamResponseToPb(st *LogParamResponse) (*logParamResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logParamResponsePb{}

	return pb, nil
}

type logParamResponsePb struct {
}

func logParamResponseFromPb(pb *logParamResponsePb) (*LogParamResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogParamResponse{}

	return st, nil
}

func loggedModelToPb(st *LoggedModel) (*loggedModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &loggedModelPb{}
	pb.Data = st.Data
	pb.Info = st.Info

	return pb, nil
}

type loggedModelPb struct {
	Data *LoggedModelData `json:"data,omitempty"`
	Info *LoggedModelInfo `json:"info,omitempty"`
}

func loggedModelFromPb(pb *loggedModelPb) (*LoggedModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModel{}
	st.Data = pb.Data
	st.Info = pb.Info

	return st, nil
}

func loggedModelDataToPb(st *LoggedModelData) (*loggedModelDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &loggedModelDataPb{}
	pb.Metrics = st.Metrics
	pb.Params = st.Params

	return pb, nil
}

type loggedModelDataPb struct {
	Metrics []Metric               `json:"metrics,omitempty"`
	Params  []LoggedModelParameter `json:"params,omitempty"`
}

func loggedModelDataFromPb(pb *loggedModelDataPb) (*LoggedModelData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModelData{}
	st.Metrics = pb.Metrics
	st.Params = pb.Params

	return st, nil
}

func loggedModelInfoToPb(st *LoggedModelInfo) (*loggedModelInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &loggedModelInfoPb{}
	pb.ArtifactUri = st.ArtifactUri
	pb.CreationTimestampMs = st.CreationTimestampMs
	pb.CreatorId = st.CreatorId
	pb.ExperimentId = st.ExperimentId
	pb.LastUpdatedTimestampMs = st.LastUpdatedTimestampMs
	pb.ModelId = st.ModelId
	pb.ModelType = st.ModelType
	pb.Name = st.Name
	pb.SourceRunId = st.SourceRunId
	pb.Status = st.Status
	pb.StatusMessage = st.StatusMessage
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type loggedModelInfoPb struct {
	ArtifactUri            string            `json:"artifact_uri,omitempty"`
	CreationTimestampMs    int64             `json:"creation_timestamp_ms,omitempty"`
	CreatorId              int64             `json:"creator_id,omitempty"`
	ExperimentId           string            `json:"experiment_id,omitempty"`
	LastUpdatedTimestampMs int64             `json:"last_updated_timestamp_ms,omitempty"`
	ModelId                string            `json:"model_id,omitempty"`
	ModelType              string            `json:"model_type,omitempty"`
	Name                   string            `json:"name,omitempty"`
	SourceRunId            string            `json:"source_run_id,omitempty"`
	Status                 LoggedModelStatus `json:"status,omitempty"`
	StatusMessage          string            `json:"status_message,omitempty"`
	Tags                   []LoggedModelTag  `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func loggedModelInfoFromPb(pb *loggedModelInfoPb) (*LoggedModelInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModelInfo{}
	st.ArtifactUri = pb.ArtifactUri
	st.CreationTimestampMs = pb.CreationTimestampMs
	st.CreatorId = pb.CreatorId
	st.ExperimentId = pb.ExperimentId
	st.LastUpdatedTimestampMs = pb.LastUpdatedTimestampMs
	st.ModelId = pb.ModelId
	st.ModelType = pb.ModelType
	st.Name = pb.Name
	st.SourceRunId = pb.SourceRunId
	st.Status = pb.Status
	st.StatusMessage = pb.StatusMessage
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *loggedModelInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st loggedModelInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func loggedModelParameterToPb(st *LoggedModelParameter) (*loggedModelParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &loggedModelParameterPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type loggedModelParameterPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func loggedModelParameterFromPb(pb *loggedModelParameterPb) (*LoggedModelParameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModelParameter{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *loggedModelParameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st loggedModelParameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func loggedModelTagToPb(st *LoggedModelTag) (*loggedModelTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &loggedModelTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type loggedModelTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func loggedModelTagFromPb(pb *loggedModelTagPb) (*LoggedModelTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModelTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *loggedModelTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st loggedModelTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func metricToPb(st *Metric) (*metricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &metricPb{}
	pb.DatasetDigest = st.DatasetDigest
	pb.DatasetName = st.DatasetName
	pb.Key = st.Key
	pb.ModelId = st.ModelId
	pb.RunId = st.RunId
	pb.Step = st.Step
	pb.Timestamp = st.Timestamp
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type metricPb struct {
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

func metricFromPb(pb *metricPb) (*Metric, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Metric{}
	st.DatasetDigest = pb.DatasetDigest
	st.DatasetName = pb.DatasetName
	st.Key = pb.Key
	st.ModelId = pb.ModelId
	st.RunId = pb.RunId
	st.Step = pb.Step
	st.Timestamp = pb.Timestamp
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *metricPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st metricPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func modelToPb(st *Model) (*modelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Description = st.Description
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.LatestVersions = st.LatestVersions
	pb.Name = st.Name
	pb.Tags = st.Tags
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type modelPb struct {
	CreationTimestamp    int64          `json:"creation_timestamp,omitempty"`
	Description          string         `json:"description,omitempty"`
	LastUpdatedTimestamp int64          `json:"last_updated_timestamp,omitempty"`
	LatestVersions       []ModelVersion `json:"latest_versions,omitempty"`
	Name                 string         `json:"name,omitempty"`
	Tags                 []ModelTag     `json:"tags,omitempty"`
	UserId               string         `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelFromPb(pb *modelPb) (*Model, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Model{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Description = pb.Description
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.LatestVersions = pb.LatestVersions
	st.Name = pb.Name
	st.Tags = pb.Tags
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func modelDatabricksToPb(st *ModelDatabricks) (*modelDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelDatabricksPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Description = st.Description
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.LatestVersions = st.LatestVersions
	pb.Name = st.Name
	pb.PermissionLevel = st.PermissionLevel
	pb.Tags = st.Tags
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type modelDatabricksPb struct {
	CreationTimestamp    int64           `json:"creation_timestamp,omitempty"`
	Description          string          `json:"description,omitempty"`
	Id                   string          `json:"id,omitempty"`
	LastUpdatedTimestamp int64           `json:"last_updated_timestamp,omitempty"`
	LatestVersions       []ModelVersion  `json:"latest_versions,omitempty"`
	Name                 string          `json:"name,omitempty"`
	PermissionLevel      PermissionLevel `json:"permission_level,omitempty"`
	Tags                 []ModelTag      `json:"tags,omitempty"`
	UserId               string          `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelDatabricksFromPb(pb *modelDatabricksPb) (*ModelDatabricks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelDatabricks{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Description = pb.Description
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.LatestVersions = pb.LatestVersions
	st.Name = pb.Name
	st.PermissionLevel = pb.PermissionLevel
	st.Tags = pb.Tags
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func modelInputToPb(st *ModelInput) (*modelInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelInputPb{}
	pb.ModelId = st.ModelId

	return pb, nil
}

type modelInputPb struct {
	ModelId string `json:"model_id"`
}

func modelInputFromPb(pb *modelInputPb) (*ModelInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelInput{}
	st.ModelId = pb.ModelId

	return st, nil
}

func modelOutputToPb(st *ModelOutput) (*modelOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelOutputPb{}
	pb.ModelId = st.ModelId
	pb.Step = st.Step

	return pb, nil
}

type modelOutputPb struct {
	ModelId string `json:"model_id"`
	Step    int64  `json:"step"`
}

func modelOutputFromPb(pb *modelOutputPb) (*ModelOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelOutput{}
	st.ModelId = pb.ModelId
	st.Step = pb.Step

	return st, nil
}

func modelTagToPb(st *ModelTag) (*modelTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type modelTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelTagFromPb(pb *modelTagPb) (*ModelTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func modelVersionToPb(st *ModelVersion) (*modelVersionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelVersionPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.CurrentStage = st.CurrentStage
	pb.Description = st.Description
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.Name = st.Name
	pb.RunId = st.RunId
	pb.RunLink = st.RunLink
	pb.Source = st.Source
	pb.Status = st.Status
	pb.StatusMessage = st.StatusMessage
	pb.Tags = st.Tags
	pb.UserId = st.UserId
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type modelVersionPb struct {
	CreationTimestamp    int64              `json:"creation_timestamp,omitempty"`
	CurrentStage         string             `json:"current_stage,omitempty"`
	Description          string             `json:"description,omitempty"`
	LastUpdatedTimestamp int64              `json:"last_updated_timestamp,omitempty"`
	Name                 string             `json:"name,omitempty"`
	RunId                string             `json:"run_id,omitempty"`
	RunLink              string             `json:"run_link,omitempty"`
	Source               string             `json:"source,omitempty"`
	Status               ModelVersionStatus `json:"status,omitempty"`
	StatusMessage        string             `json:"status_message,omitempty"`
	Tags                 []ModelVersionTag  `json:"tags,omitempty"`
	UserId               string             `json:"user_id,omitempty"`
	Version              string             `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelVersionFromPb(pb *modelVersionPb) (*ModelVersion, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersion{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.CurrentStage = pb.CurrentStage
	st.Description = pb.Description
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name
	st.RunId = pb.RunId
	st.RunLink = pb.RunLink
	st.Source = pb.Source
	st.Status = pb.Status
	st.StatusMessage = pb.StatusMessage
	st.Tags = pb.Tags
	st.UserId = pb.UserId
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelVersionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelVersionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func modelVersionDatabricksToPb(st *ModelVersionDatabricks) (*modelVersionDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelVersionDatabricksPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.CurrentStage = st.CurrentStage
	pb.Description = st.Description
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.Name = st.Name
	pb.PermissionLevel = st.PermissionLevel
	pb.RunId = st.RunId
	pb.RunLink = st.RunLink
	pb.Source = st.Source
	pb.Status = st.Status
	pb.StatusMessage = st.StatusMessage
	pb.Tags = st.Tags
	pb.UserId = st.UserId
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type modelVersionDatabricksPb struct {
	CreationTimestamp    int64             `json:"creation_timestamp,omitempty"`
	CurrentStage         Stage             `json:"current_stage,omitempty"`
	Description          string            `json:"description,omitempty"`
	LastUpdatedTimestamp int64             `json:"last_updated_timestamp,omitempty"`
	Name                 string            `json:"name,omitempty"`
	PermissionLevel      PermissionLevel   `json:"permission_level,omitempty"`
	RunId                string            `json:"run_id,omitempty"`
	RunLink              string            `json:"run_link,omitempty"`
	Source               string            `json:"source,omitempty"`
	Status               Status            `json:"status,omitempty"`
	StatusMessage        string            `json:"status_message,omitempty"`
	Tags                 []ModelVersionTag `json:"tags,omitempty"`
	UserId               string            `json:"user_id,omitempty"`
	Version              string            `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelVersionDatabricksFromPb(pb *modelVersionDatabricksPb) (*ModelVersionDatabricks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionDatabricks{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.CurrentStage = pb.CurrentStage
	st.Description = pb.Description
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name
	st.PermissionLevel = pb.PermissionLevel
	st.RunId = pb.RunId
	st.RunLink = pb.RunLink
	st.Source = pb.Source
	st.Status = pb.Status
	st.StatusMessage = pb.StatusMessage
	st.Tags = pb.Tags
	st.UserId = pb.UserId
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelVersionDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelVersionDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func modelVersionTagToPb(st *ModelVersionTag) (*modelVersionTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelVersionTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type modelVersionTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelVersionTagFromPb(pb *modelVersionTagPb) (*ModelVersionTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelVersionTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelVersionTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func onlineStoreToPb(st *OnlineStore) (*onlineStorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineStorePb{}
	pb.Capacity = st.Capacity
	pb.CreationTime = st.CreationTime
	pb.Creator = st.Creator
	pb.Name = st.Name
	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type onlineStorePb struct {
	Capacity     string           `json:"capacity,omitempty"`
	CreationTime string           `json:"creation_time,omitempty"`
	Creator      string           `json:"creator,omitempty"`
	Name         string           `json:"name"`
	State        OnlineStoreState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func onlineStoreFromPb(pb *onlineStorePb) (*OnlineStore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineStore{}
	st.Capacity = pb.Capacity
	st.CreationTime = pb.CreationTime
	st.Creator = pb.Creator
	st.Name = pb.Name
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *onlineStorePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st onlineStorePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func paramToPb(st *Param) (*paramPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &paramPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type paramPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func paramFromPb(pb *paramPb) (*Param, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Param{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *paramPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st paramPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func publishSpecToPb(st *PublishSpec) (*publishSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishSpecPb{}
	pb.OnlineStore = st.OnlineStore
	pb.OnlineTableName = st.OnlineTableName
	pb.PublishMode = st.PublishMode

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type publishSpecPb struct {
	OnlineStore     string                 `json:"online_store"`
	OnlineTableName string                 `json:"online_table_name,omitempty"`
	PublishMode     PublishSpecPublishMode `json:"publish_mode,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func publishSpecFromPb(pb *publishSpecPb) (*PublishSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishSpec{}
	st.OnlineStore = pb.OnlineStore
	st.OnlineTableName = pb.OnlineTableName
	st.PublishMode = pb.PublishMode

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publishSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publishSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func publishTableRequestToPb(st *PublishTableRequest) (*publishTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishTableRequestPb{}
	pb.PublishSpec = st.PublishSpec
	pb.SourceTableName = st.SourceTableName

	return pb, nil
}

type publishTableRequestPb struct {
	PublishSpec     PublishSpec `json:"publish_spec"`
	SourceTableName string      `json:"-" url:"-"`
}

func publishTableRequestFromPb(pb *publishTableRequestPb) (*PublishTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishTableRequest{}
	st.PublishSpec = pb.PublishSpec
	st.SourceTableName = pb.SourceTableName

	return st, nil
}

func publishTableResponseToPb(st *PublishTableResponse) (*publishTableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishTableResponsePb{}
	pb.OnlineTableName = st.OnlineTableName
	pb.PipelineId = st.PipelineId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type publishTableResponsePb struct {
	OnlineTableName string `json:"online_table_name,omitempty"`
	PipelineId      string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func publishTableResponseFromPb(pb *publishTableResponsePb) (*PublishTableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishTableResponse{}
	st.OnlineTableName = pb.OnlineTableName
	st.PipelineId = pb.PipelineId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publishTableResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publishTableResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelAccessControlRequestToPb(st *RegisteredModelAccessControlRequest) (*registeredModelAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	pb.PermissionLevel = st.PermissionLevel
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registeredModelAccessControlRequestPb struct {
	GroupName            string                         `json:"group_name,omitempty"`
	PermissionLevel      RegisteredModelPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                         `json:"service_principal_name,omitempty"`
	UserName             string                         `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelAccessControlRequestFromPb(pb *registeredModelAccessControlRequestPb) (*RegisteredModelAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelAccessControlResponseToPb(st *RegisteredModelAccessControlResponse) (*registeredModelAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registeredModelAccessControlResponsePb struct {
	AllPermissions       []RegisteredModelPermission `json:"all_permissions,omitempty"`
	DisplayName          string                      `json:"display_name,omitempty"`
	GroupName            string                      `json:"group_name,omitempty"`
	ServicePrincipalName string                      `json:"service_principal_name,omitempty"`
	UserName             string                      `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelAccessControlResponseFromPb(pb *registeredModelAccessControlResponsePb) (*RegisteredModelAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelPermissionToPb(st *RegisteredModelPermission) (*registeredModelPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registeredModelPermissionPb struct {
	Inherited           bool                           `json:"inherited,omitempty"`
	InheritedFromObject []string                       `json:"inherited_from_object,omitempty"`
	PermissionLevel     RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelPermissionFromPb(pb *registeredModelPermissionPb) (*RegisteredModelPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelPermissionsToPb(st *RegisteredModelPermissions) (*registeredModelPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registeredModelPermissionsPb struct {
	AccessControlList []RegisteredModelAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                                 `json:"object_id,omitempty"`
	ObjectType        string                                 `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelPermissionsFromPb(pb *registeredModelPermissionsPb) (*RegisteredModelPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelPermissionsDescriptionToPb(st *RegisteredModelPermissionsDescription) (*registeredModelPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelPermissionsDescriptionPb{}
	pb.Description = st.Description
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registeredModelPermissionsDescriptionPb struct {
	Description     string                         `json:"description,omitempty"`
	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelPermissionsDescriptionFromPb(pb *registeredModelPermissionsDescriptionPb) (*RegisteredModelPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelPermissionsRequestToPb(st *RegisteredModelPermissionsRequest) (*registeredModelPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.RegisteredModelId = st.RegisteredModelId

	return pb, nil
}

type registeredModelPermissionsRequestPb struct {
	AccessControlList []RegisteredModelAccessControlRequest `json:"access_control_list,omitempty"`
	RegisteredModelId string                                `json:"-" url:"-"`
}

func registeredModelPermissionsRequestFromPb(pb *registeredModelPermissionsRequestPb) (*RegisteredModelPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.RegisteredModelId = pb.RegisteredModelId

	return st, nil
}

func registryWebhookToPb(st *RegistryWebhook) (*registryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registryWebhookPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Description = st.Description
	pb.Events = st.Events
	pb.HttpUrlSpec = st.HttpUrlSpec
	pb.Id = st.Id
	pb.JobSpec = st.JobSpec
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.ModelName = st.ModelName
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registryWebhookPb struct {
	CreationTimestamp    int64                     `json:"creation_timestamp,omitempty"`
	Description          string                    `json:"description,omitempty"`
	Events               []RegistryWebhookEvent    `json:"events,omitempty"`
	HttpUrlSpec          *HttpUrlSpecWithoutSecret `json:"http_url_spec,omitempty"`
	Id                   string                    `json:"id,omitempty"`
	JobSpec              *JobSpecWithoutSecret     `json:"job_spec,omitempty"`
	LastUpdatedTimestamp int64                     `json:"last_updated_timestamp,omitempty"`
	ModelName            string                    `json:"model_name,omitempty"`
	Status               RegistryWebhookStatus     `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registryWebhookFromPb(pb *registryWebhookPb) (*RegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegistryWebhook{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Description = pb.Description
	st.Events = pb.Events
	st.HttpUrlSpec = pb.HttpUrlSpec
	st.Id = pb.Id
	st.JobSpec = pb.JobSpec
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.ModelName = pb.ModelName
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func rejectTransitionRequestToPb(st *RejectTransitionRequest) (*rejectTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rejectTransitionRequestPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type rejectTransitionRequestPb struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name"`
	Stage   Stage  `json:"stage"`
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func rejectTransitionRequestFromPb(pb *rejectTransitionRequestPb) (*RejectTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RejectTransitionRequest{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *rejectTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st rejectTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func rejectTransitionRequestResponseToPb(st *RejectTransitionRequestResponse) (*rejectTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rejectTransitionRequestResponsePb{}
	pb.Activity = st.Activity

	return pb, nil
}

type rejectTransitionRequestResponsePb struct {
	Activity *Activity `json:"activity,omitempty"`
}

func rejectTransitionRequestResponseFromPb(pb *rejectTransitionRequestResponsePb) (*RejectTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RejectTransitionRequestResponse{}
	st.Activity = pb.Activity

	return st, nil
}

func renameModelRequestToPb(st *RenameModelRequest) (*renameModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &renameModelRequestPb{}
	pb.Name = st.Name
	pb.NewName = st.NewName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type renameModelRequestPb struct {
	Name    string `json:"name"`
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func renameModelRequestFromPb(pb *renameModelRequestPb) (*RenameModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RenameModelRequest{}
	st.Name = pb.Name
	st.NewName = pb.NewName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *renameModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st renameModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func renameModelResponseToPb(st *RenameModelResponse) (*renameModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &renameModelResponsePb{}
	pb.RegisteredModel = st.RegisteredModel

	return pb, nil
}

type renameModelResponsePb struct {
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

func renameModelResponseFromPb(pb *renameModelResponsePb) (*RenameModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RenameModelResponse{}
	st.RegisteredModel = pb.RegisteredModel

	return st, nil
}

func restoreExperimentToPb(st *RestoreExperiment) (*restoreExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreExperimentPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

type restoreExperimentPb struct {
	ExperimentId string `json:"experiment_id"`
}

func restoreExperimentFromPb(pb *restoreExperimentPb) (*RestoreExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreExperiment{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

func restoreExperimentResponseToPb(st *RestoreExperimentResponse) (*restoreExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreExperimentResponsePb{}

	return pb, nil
}

type restoreExperimentResponsePb struct {
}

func restoreExperimentResponseFromPb(pb *restoreExperimentResponsePb) (*RestoreExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreExperimentResponse{}

	return st, nil
}

func restoreRunToPb(st *RestoreRun) (*restoreRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

type restoreRunPb struct {
	RunId string `json:"run_id"`
}

func restoreRunFromPb(pb *restoreRunPb) (*RestoreRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRun{}
	st.RunId = pb.RunId

	return st, nil
}

func restoreRunResponseToPb(st *RestoreRunResponse) (*restoreRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreRunResponsePb{}

	return pb, nil
}

type restoreRunResponsePb struct {
}

func restoreRunResponseFromPb(pb *restoreRunResponsePb) (*RestoreRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRunResponse{}

	return st, nil
}

func restoreRunsToPb(st *RestoreRuns) (*restoreRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreRunsPb{}
	pb.ExperimentId = st.ExperimentId
	pb.MaxRuns = st.MaxRuns
	pb.MinTimestampMillis = st.MinTimestampMillis

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type restoreRunsPb struct {
	ExperimentId       string `json:"experiment_id"`
	MaxRuns            int    `json:"max_runs,omitempty"`
	MinTimestampMillis int64  `json:"min_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restoreRunsFromPb(pb *restoreRunsPb) (*RestoreRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRuns{}
	st.ExperimentId = pb.ExperimentId
	st.MaxRuns = pb.MaxRuns
	st.MinTimestampMillis = pb.MinTimestampMillis

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restoreRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restoreRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func restoreRunsResponseToPb(st *RestoreRunsResponse) (*restoreRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreRunsResponsePb{}
	pb.RunsRestored = st.RunsRestored

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type restoreRunsResponsePb struct {
	RunsRestored int `json:"runs_restored,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restoreRunsResponseFromPb(pb *restoreRunsResponsePb) (*RestoreRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRunsResponse{}
	st.RunsRestored = pb.RunsRestored

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restoreRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restoreRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runToPb(st *Run) (*runPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runPb{}
	pb.Data = st.Data
	pb.Info = st.Info
	pb.Inputs = st.Inputs

	return pb, nil
}

type runPb struct {
	Data   *RunData   `json:"data,omitempty"`
	Info   *RunInfo   `json:"info,omitempty"`
	Inputs *RunInputs `json:"inputs,omitempty"`
}

func runFromPb(pb *runPb) (*Run, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Run{}
	st.Data = pb.Data
	st.Info = pb.Info
	st.Inputs = pb.Inputs

	return st, nil
}

func runDataToPb(st *RunData) (*runDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runDataPb{}
	pb.Metrics = st.Metrics
	pb.Params = st.Params
	pb.Tags = st.Tags

	return pb, nil
}

type runDataPb struct {
	Metrics []Metric `json:"metrics,omitempty"`
	Params  []Param  `json:"params,omitempty"`
	Tags    []RunTag `json:"tags,omitempty"`
}

func runDataFromPb(pb *runDataPb) (*RunData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunData{}
	st.Metrics = pb.Metrics
	st.Params = pb.Params
	st.Tags = pb.Tags

	return st, nil
}

func runInfoToPb(st *RunInfo) (*runInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runInfoPb{}
	pb.ArtifactUri = st.ArtifactUri
	pb.EndTime = st.EndTime
	pb.ExperimentId = st.ExperimentId
	pb.LifecycleStage = st.LifecycleStage
	pb.RunId = st.RunId
	pb.RunName = st.RunName
	pb.RunUuid = st.RunUuid
	pb.StartTime = st.StartTime
	pb.Status = st.Status
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runInfoPb struct {
	ArtifactUri    string        `json:"artifact_uri,omitempty"`
	EndTime        int64         `json:"end_time,omitempty"`
	ExperimentId   string        `json:"experiment_id,omitempty"`
	LifecycleStage string        `json:"lifecycle_stage,omitempty"`
	RunId          string        `json:"run_id,omitempty"`
	RunName        string        `json:"run_name,omitempty"`
	RunUuid        string        `json:"run_uuid,omitempty"`
	StartTime      int64         `json:"start_time,omitempty"`
	Status         RunInfoStatus `json:"status,omitempty"`
	UserId         string        `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runInfoFromPb(pb *runInfoPb) (*RunInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunInfo{}
	st.ArtifactUri = pb.ArtifactUri
	st.EndTime = pb.EndTime
	st.ExperimentId = pb.ExperimentId
	st.LifecycleStage = pb.LifecycleStage
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunUuid = pb.RunUuid
	st.StartTime = pb.StartTime
	st.Status = pb.Status
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runInputsToPb(st *RunInputs) (*runInputsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runInputsPb{}
	pb.DatasetInputs = st.DatasetInputs
	pb.ModelInputs = st.ModelInputs

	return pb, nil
}

type runInputsPb struct {
	DatasetInputs []DatasetInput `json:"dataset_inputs,omitempty"`
	ModelInputs   []ModelInput   `json:"model_inputs,omitempty"`
}

func runInputsFromPb(pb *runInputsPb) (*RunInputs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunInputs{}
	st.DatasetInputs = pb.DatasetInputs
	st.ModelInputs = pb.ModelInputs

	return st, nil
}

func runTagToPb(st *RunTag) (*runTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runTagPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runTagFromPb(pb *runTagPb) (*RunTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchExperimentsToPb(st *SearchExperiments) (*searchExperimentsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchExperimentsPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken
	pb.ViewType = st.ViewType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchExperimentsPb struct {
	Filter     string   `json:"filter,omitempty"`
	MaxResults int64    `json:"max_results,omitempty"`
	OrderBy    []string `json:"order_by,omitempty"`
	PageToken  string   `json:"page_token,omitempty"`
	ViewType   ViewType `json:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchExperimentsFromPb(pb *searchExperimentsPb) (*SearchExperiments, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchExperiments{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken
	st.ViewType = pb.ViewType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchExperimentsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchExperimentsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchExperimentsResponseToPb(st *SearchExperimentsResponse) (*searchExperimentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchExperimentsResponsePb{}
	pb.Experiments = st.Experiments
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchExperimentsResponsePb struct {
	Experiments   []Experiment `json:"experiments,omitempty"`
	NextPageToken string       `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchExperimentsResponseFromPb(pb *searchExperimentsResponsePb) (*SearchExperimentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchExperimentsResponse{}
	st.Experiments = pb.Experiments
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchExperimentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchExperimentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchLoggedModelsDatasetToPb(st *SearchLoggedModelsDataset) (*searchLoggedModelsDatasetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchLoggedModelsDatasetPb{}
	pb.DatasetDigest = st.DatasetDigest
	pb.DatasetName = st.DatasetName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchLoggedModelsDatasetPb struct {
	DatasetDigest string `json:"dataset_digest,omitempty"`
	DatasetName   string `json:"dataset_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchLoggedModelsDatasetFromPb(pb *searchLoggedModelsDatasetPb) (*SearchLoggedModelsDataset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchLoggedModelsDataset{}
	st.DatasetDigest = pb.DatasetDigest
	st.DatasetName = pb.DatasetName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchLoggedModelsDatasetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchLoggedModelsDatasetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchLoggedModelsOrderByToPb(st *SearchLoggedModelsOrderBy) (*searchLoggedModelsOrderByPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchLoggedModelsOrderByPb{}
	pb.Ascending = st.Ascending
	pb.DatasetDigest = st.DatasetDigest
	pb.DatasetName = st.DatasetName
	pb.FieldName = st.FieldName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchLoggedModelsOrderByPb struct {
	Ascending     bool   `json:"ascending,omitempty"`
	DatasetDigest string `json:"dataset_digest,omitempty"`
	DatasetName   string `json:"dataset_name,omitempty"`
	FieldName     string `json:"field_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchLoggedModelsOrderByFromPb(pb *searchLoggedModelsOrderByPb) (*SearchLoggedModelsOrderBy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchLoggedModelsOrderBy{}
	st.Ascending = pb.Ascending
	st.DatasetDigest = pb.DatasetDigest
	st.DatasetName = pb.DatasetName
	st.FieldName = pb.FieldName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchLoggedModelsOrderByPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchLoggedModelsOrderByPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchLoggedModelsRequestToPb(st *SearchLoggedModelsRequest) (*searchLoggedModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchLoggedModelsRequestPb{}
	pb.Datasets = st.Datasets
	pb.ExperimentIds = st.ExperimentIds
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchLoggedModelsRequestPb struct {
	Datasets      []SearchLoggedModelsDataset `json:"datasets,omitempty"`
	ExperimentIds []string                    `json:"experiment_ids,omitempty"`
	Filter        string                      `json:"filter,omitempty"`
	MaxResults    int                         `json:"max_results,omitempty"`
	OrderBy       []SearchLoggedModelsOrderBy `json:"order_by,omitempty"`
	PageToken     string                      `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchLoggedModelsRequestFromPb(pb *searchLoggedModelsRequestPb) (*SearchLoggedModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchLoggedModelsRequest{}
	st.Datasets = pb.Datasets
	st.ExperimentIds = pb.ExperimentIds
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchLoggedModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchLoggedModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchLoggedModelsResponseToPb(st *SearchLoggedModelsResponse) (*searchLoggedModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchLoggedModelsResponsePb{}
	pb.Models = st.Models
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchLoggedModelsResponsePb struct {
	Models        []LoggedModel `json:"models,omitempty"`
	NextPageToken string        `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchLoggedModelsResponseFromPb(pb *searchLoggedModelsResponsePb) (*SearchLoggedModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchLoggedModelsResponse{}
	st.Models = pb.Models
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchLoggedModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchLoggedModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchModelVersionsRequestToPb(st *SearchModelVersionsRequest) (*searchModelVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchModelVersionsRequestPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchModelVersionsRequestPb struct {
	Filter     string   `json:"-" url:"filter,omitempty"`
	MaxResults int      `json:"-" url:"max_results,omitempty"`
	OrderBy    []string `json:"-" url:"order_by,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchModelVersionsRequestFromPb(pb *searchModelVersionsRequestPb) (*SearchModelVersionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelVersionsRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchModelVersionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchModelVersionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchModelVersionsResponseToPb(st *SearchModelVersionsResponse) (*searchModelVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchModelVersionsResponsePb{}
	pb.ModelVersions = st.ModelVersions
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchModelVersionsResponsePb struct {
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
	NextPageToken string         `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchModelVersionsResponseFromPb(pb *searchModelVersionsResponsePb) (*SearchModelVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelVersionsResponse{}
	st.ModelVersions = pb.ModelVersions
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchModelVersionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchModelVersionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchModelsRequestToPb(st *SearchModelsRequest) (*searchModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchModelsRequestPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchModelsRequestPb struct {
	Filter     string   `json:"-" url:"filter,omitempty"`
	MaxResults int      `json:"-" url:"max_results,omitempty"`
	OrderBy    []string `json:"-" url:"order_by,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchModelsRequestFromPb(pb *searchModelsRequestPb) (*SearchModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelsRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchModelsResponseToPb(st *SearchModelsResponse) (*searchModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.RegisteredModels = st.RegisteredModels

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchModelsResponsePb struct {
	NextPageToken    string  `json:"next_page_token,omitempty"`
	RegisteredModels []Model `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchModelsResponseFromPb(pb *searchModelsResponsePb) (*SearchModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.RegisteredModels = pb.RegisteredModels

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchRunsToPb(st *SearchRuns) (*searchRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchRunsPb{}
	pb.ExperimentIds = st.ExperimentIds
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken
	pb.RunViewType = st.RunViewType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchRunsPb struct {
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	Filter        string   `json:"filter,omitempty"`
	MaxResults    int      `json:"max_results,omitempty"`
	OrderBy       []string `json:"order_by,omitempty"`
	PageToken     string   `json:"page_token,omitempty"`
	RunViewType   ViewType `json:"run_view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchRunsFromPb(pb *searchRunsPb) (*SearchRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchRuns{}
	st.ExperimentIds = pb.ExperimentIds
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken
	st.RunViewType = pb.RunViewType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchRunsResponseToPb(st *SearchRunsResponse) (*searchRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchRunsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Runs = st.Runs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchRunsResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`
	Runs          []Run  `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchRunsResponseFromPb(pb *searchRunsResponsePb) (*SearchRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchRunsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Runs = pb.Runs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func setExperimentTagToPb(st *SetExperimentTag) (*setExperimentTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setExperimentTagPb{}
	pb.ExperimentId = st.ExperimentId
	pb.Key = st.Key
	pb.Value = st.Value

	return pb, nil
}

type setExperimentTagPb struct {
	ExperimentId string `json:"experiment_id"`
	Key          string `json:"key"`
	Value        string `json:"value"`
}

func setExperimentTagFromPb(pb *setExperimentTagPb) (*SetExperimentTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetExperimentTag{}
	st.ExperimentId = pb.ExperimentId
	st.Key = pb.Key
	st.Value = pb.Value

	return st, nil
}

func setExperimentTagResponseToPb(st *SetExperimentTagResponse) (*setExperimentTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setExperimentTagResponsePb{}

	return pb, nil
}

type setExperimentTagResponsePb struct {
}

func setExperimentTagResponseFromPb(pb *setExperimentTagResponsePb) (*SetExperimentTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetExperimentTagResponse{}

	return st, nil
}

func setLoggedModelTagsRequestToPb(st *SetLoggedModelTagsRequest) (*setLoggedModelTagsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setLoggedModelTagsRequestPb{}
	pb.ModelId = st.ModelId
	pb.Tags = st.Tags

	return pb, nil
}

type setLoggedModelTagsRequestPb struct {
	ModelId string           `json:"-" url:"-"`
	Tags    []LoggedModelTag `json:"tags,omitempty"`
}

func setLoggedModelTagsRequestFromPb(pb *setLoggedModelTagsRequestPb) (*SetLoggedModelTagsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetLoggedModelTagsRequest{}
	st.ModelId = pb.ModelId
	st.Tags = pb.Tags

	return st, nil
}

func setLoggedModelTagsResponseToPb(st *SetLoggedModelTagsResponse) (*setLoggedModelTagsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setLoggedModelTagsResponsePb{}

	return pb, nil
}

type setLoggedModelTagsResponsePb struct {
}

func setLoggedModelTagsResponseFromPb(pb *setLoggedModelTagsResponsePb) (*SetLoggedModelTagsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetLoggedModelTagsResponse{}

	return st, nil
}

func setModelTagRequestToPb(st *SetModelTagRequest) (*setModelTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setModelTagRequestPb{}
	pb.Key = st.Key
	pb.Name = st.Name
	pb.Value = st.Value

	return pb, nil
}

type setModelTagRequestPb struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func setModelTagRequestFromPb(pb *setModelTagRequestPb) (*SetModelTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name
	st.Value = pb.Value

	return st, nil
}

func setModelTagResponseToPb(st *SetModelTagResponse) (*setModelTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setModelTagResponsePb{}

	return pb, nil
}

type setModelTagResponsePb struct {
}

func setModelTagResponseFromPb(pb *setModelTagResponsePb) (*SetModelTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelTagResponse{}

	return st, nil
}

func setModelVersionTagRequestToPb(st *SetModelVersionTagRequest) (*setModelVersionTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setModelVersionTagRequestPb{}
	pb.Key = st.Key
	pb.Name = st.Name
	pb.Value = st.Value
	pb.Version = st.Version

	return pb, nil
}

type setModelVersionTagRequestPb struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Version string `json:"version"`
}

func setModelVersionTagRequestFromPb(pb *setModelVersionTagRequestPb) (*SetModelVersionTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelVersionTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name
	st.Value = pb.Value
	st.Version = pb.Version

	return st, nil
}

func setModelVersionTagResponseToPb(st *SetModelVersionTagResponse) (*setModelVersionTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setModelVersionTagResponsePb{}

	return pb, nil
}

type setModelVersionTagResponsePb struct {
}

func setModelVersionTagResponseFromPb(pb *setModelVersionTagResponsePb) (*SetModelVersionTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelVersionTagResponse{}

	return st, nil
}

func setTagToPb(st *SetTag) (*setTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setTagPb{}
	pb.Key = st.Key
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type setTagPb struct {
	Key     string `json:"key"`
	RunId   string `json:"run_id,omitempty"`
	RunUuid string `json:"run_uuid,omitempty"`
	Value   string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func setTagFromPb(pb *setTagPb) (*SetTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetTag{}
	st.Key = pb.Key
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *setTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st setTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func setTagResponseToPb(st *SetTagResponse) (*setTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setTagResponsePb{}

	return pb, nil
}

type setTagResponsePb struct {
}

func setTagResponseFromPb(pb *setTagResponsePb) (*SetTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetTagResponse{}

	return st, nil
}

func testRegistryWebhookToPb(st *TestRegistryWebhook) (*testRegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &testRegistryWebhookPb{}
	pb.Body = st.Body
	pb.StatusCode = st.StatusCode

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type testRegistryWebhookPb struct {
	Body       string `json:"body,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func testRegistryWebhookFromPb(pb *testRegistryWebhookPb) (*TestRegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestRegistryWebhook{}
	st.Body = pb.Body
	st.StatusCode = pb.StatusCode

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *testRegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st testRegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func testRegistryWebhookRequestToPb(st *TestRegistryWebhookRequest) (*testRegistryWebhookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &testRegistryWebhookRequestPb{}
	pb.Event = st.Event
	pb.Id = st.Id

	return pb, nil
}

type testRegistryWebhookRequestPb struct {
	Event RegistryWebhookEvent `json:"event,omitempty"`
	Id    string               `json:"id"`
}

func testRegistryWebhookRequestFromPb(pb *testRegistryWebhookRequestPb) (*TestRegistryWebhookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestRegistryWebhookRequest{}
	st.Event = pb.Event
	st.Id = pb.Id

	return st, nil
}

func testRegistryWebhookResponseToPb(st *TestRegistryWebhookResponse) (*testRegistryWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &testRegistryWebhookResponsePb{}
	pb.Webhook = st.Webhook

	return pb, nil
}

type testRegistryWebhookResponsePb struct {
	Webhook *TestRegistryWebhook `json:"webhook,omitempty"`
}

func testRegistryWebhookResponseFromPb(pb *testRegistryWebhookResponsePb) (*TestRegistryWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestRegistryWebhookResponse{}
	st.Webhook = pb.Webhook

	return st, nil
}

func transitionModelVersionStageDatabricksToPb(st *TransitionModelVersionStageDatabricks) (*transitionModelVersionStageDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transitionModelVersionStageDatabricksPb{}
	pb.ArchiveExistingVersions = st.ArchiveExistingVersions
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type transitionModelVersionStageDatabricksPb struct {
	ArchiveExistingVersions bool   `json:"archive_existing_versions"`
	Comment                 string `json:"comment,omitempty"`
	Name                    string `json:"name"`
	Stage                   Stage  `json:"stage"`
	Version                 string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func transitionModelVersionStageDatabricksFromPb(pb *transitionModelVersionStageDatabricksPb) (*TransitionModelVersionStageDatabricks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionModelVersionStageDatabricks{}
	st.ArchiveExistingVersions = pb.ArchiveExistingVersions
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *transitionModelVersionStageDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st transitionModelVersionStageDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func transitionRequestToPb(st *TransitionRequest) (*transitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transitionRequestPb{}
	pb.AvailableActions = st.AvailableActions
	pb.Comment = st.Comment
	pb.CreationTimestamp = st.CreationTimestamp
	pb.ToStage = st.ToStage
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type transitionRequestPb struct {
	AvailableActions  []ActivityAction `json:"available_actions,omitempty"`
	Comment           string           `json:"comment,omitempty"`
	CreationTimestamp int64            `json:"creation_timestamp,omitempty"`
	ToStage           Stage            `json:"to_stage,omitempty"`
	UserId            string           `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func transitionRequestFromPb(pb *transitionRequestPb) (*TransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionRequest{}
	st.AvailableActions = pb.AvailableActions
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.ToStage = pb.ToStage
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *transitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st transitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func transitionStageResponseToPb(st *TransitionStageResponse) (*transitionStageResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transitionStageResponsePb{}
	pb.ModelVersion = st.ModelVersion

	return pb, nil
}

type transitionStageResponsePb struct {
	ModelVersion *ModelVersionDatabricks `json:"model_version,omitempty"`
}

func transitionStageResponseFromPb(pb *transitionStageResponsePb) (*TransitionStageResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionStageResponse{}
	st.ModelVersion = pb.ModelVersion

	return st, nil
}

func updateCommentToPb(st *UpdateComment) (*updateCommentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCommentPb{}
	pb.Comment = st.Comment
	pb.Id = st.Id

	return pb, nil
}

type updateCommentPb struct {
	Comment string `json:"comment"`
	Id      string `json:"id"`
}

func updateCommentFromPb(pb *updateCommentPb) (*UpdateComment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateComment{}
	st.Comment = pb.Comment
	st.Id = pb.Id

	return st, nil
}

func updateCommentResponseToPb(st *UpdateCommentResponse) (*updateCommentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCommentResponsePb{}
	pb.Comment = st.Comment

	return pb, nil
}

type updateCommentResponsePb struct {
	Comment *CommentObject `json:"comment,omitempty"`
}

func updateCommentResponseFromPb(pb *updateCommentResponsePb) (*UpdateCommentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCommentResponse{}
	st.Comment = pb.Comment

	return st, nil
}

func updateExperimentToPb(st *UpdateExperiment) (*updateExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExperimentPb{}
	pb.ExperimentId = st.ExperimentId
	pb.NewName = st.NewName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateExperimentPb struct {
	ExperimentId string `json:"experiment_id"`
	NewName      string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateExperimentFromPb(pb *updateExperimentPb) (*UpdateExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExperiment{}
	st.ExperimentId = pb.ExperimentId
	st.NewName = pb.NewName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateExperimentResponseToPb(st *UpdateExperimentResponse) (*updateExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExperimentResponsePb{}

	return pb, nil
}

type updateExperimentResponsePb struct {
}

func updateExperimentResponseFromPb(pb *updateExperimentResponsePb) (*UpdateExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExperimentResponse{}

	return st, nil
}

func updateModelRequestToPb(st *UpdateModelRequest) (*updateModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateModelRequestPb struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateModelRequestFromPb(pb *updateModelRequestPb) (*UpdateModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelRequest{}
	st.Description = pb.Description
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateModelResponseToPb(st *UpdateModelResponse) (*updateModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelResponsePb{}

	return pb, nil
}

type updateModelResponsePb struct {
}

func updateModelResponseFromPb(pb *updateModelResponsePb) (*UpdateModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelResponse{}

	return st, nil
}

func updateModelVersionRequestToPb(st *UpdateModelVersionRequest) (*updateModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelVersionRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateModelVersionRequestPb struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Version     string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateModelVersionRequestFromPb(pb *updateModelVersionRequestPb) (*UpdateModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelVersionRequest{}
	st.Description = pb.Description
	st.Name = pb.Name
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateModelVersionResponseToPb(st *UpdateModelVersionResponse) (*updateModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelVersionResponsePb{}

	return pb, nil
}

type updateModelVersionResponsePb struct {
}

func updateModelVersionResponseFromPb(pb *updateModelVersionResponsePb) (*UpdateModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelVersionResponse{}

	return st, nil
}

func updateOnlineStoreRequestToPb(st *UpdateOnlineStoreRequest) (*updateOnlineStoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateOnlineStoreRequestPb{}
	pb.Name = st.Name
	pb.OnlineStore = st.OnlineStore
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

type updateOnlineStoreRequestPb struct {
	Name        string      `json:"-" url:"-"`
	OnlineStore OnlineStore `json:"online_store"`
	UpdateMask  string      `json:"-" url:"update_mask"`
}

func updateOnlineStoreRequestFromPb(pb *updateOnlineStoreRequestPb) (*UpdateOnlineStoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateOnlineStoreRequest{}
	st.Name = pb.Name
	st.OnlineStore = pb.OnlineStore
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

func updateRegistryWebhookToPb(st *UpdateRegistryWebhook) (*updateRegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRegistryWebhookPb{}
	pb.Description = st.Description
	pb.Events = st.Events
	pb.HttpUrlSpec = st.HttpUrlSpec
	pb.Id = st.Id
	pb.JobSpec = st.JobSpec
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateRegistryWebhookPb struct {
	Description string                 `json:"description,omitempty"`
	Events      []RegistryWebhookEvent `json:"events,omitempty"`
	HttpUrlSpec *HttpUrlSpec           `json:"http_url_spec,omitempty"`
	Id          string                 `json:"id"`
	JobSpec     *JobSpec               `json:"job_spec,omitempty"`
	Status      RegistryWebhookStatus  `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRegistryWebhookFromPb(pb *updateRegistryWebhookPb) (*UpdateRegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRegistryWebhook{}
	st.Description = pb.Description
	st.Events = pb.Events
	st.HttpUrlSpec = pb.HttpUrlSpec
	st.Id = pb.Id
	st.JobSpec = pb.JobSpec
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateRegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateRunToPb(st *UpdateRun) (*updateRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRunPb{}
	pb.EndTime = st.EndTime
	pb.RunId = st.RunId
	pb.RunName = st.RunName
	pb.RunUuid = st.RunUuid
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateRunPb struct {
	EndTime int64           `json:"end_time,omitempty"`
	RunId   string          `json:"run_id,omitempty"`
	RunName string          `json:"run_name,omitempty"`
	RunUuid string          `json:"run_uuid,omitempty"`
	Status  UpdateRunStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRunFromPb(pb *updateRunPb) (*UpdateRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRun{}
	st.EndTime = pb.EndTime
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunUuid = pb.RunUuid
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateRunResponseToPb(st *UpdateRunResponse) (*updateRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRunResponsePb{}
	pb.RunInfo = st.RunInfo

	return pb, nil
}

type updateRunResponsePb struct {
	RunInfo *RunInfo `json:"run_info,omitempty"`
}

func updateRunResponseFromPb(pb *updateRunResponsePb) (*UpdateRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRunResponse{}
	st.RunInfo = pb.RunInfo

	return st, nil
}

func updateWebhookResponseToPb(st *UpdateWebhookResponse) (*updateWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWebhookResponsePb{}

	return pb, nil
}

type updateWebhookResponsePb struct {
}

func updateWebhookResponseFromPb(pb *updateWebhookResponsePb) (*UpdateWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWebhookResponse{}

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
