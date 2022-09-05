// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflow

// all definitions in this file are in alphabetical order

type CreateExperimentRequest struct {
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

type CreateRunRequest struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime int64 `json:"start_time,omitempty"`
	// Additional metadata for run.
	Tags []RunTag `json:"tags,omitempty"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use &#39;mlflow.user&#39;
	// tag instead.
	UserId string `json:"user_id,omitempty"`
}

type CreateRunResponse struct {
	// The newly created run.
	Run *Run `json:"run,omitempty"`
}

type DeleteExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

type DeleteRunRequest struct {
	// ID of the run to delete.
	RunId string `json:"run_id"`
}

type DeleteTagRequest struct {
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
	// Current life cycle stage of the experiment: &#34;active&#34; or &#34;deleted&#34;.
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

type GetExperimentByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName string ` url:"experiment_name,omitempty"`
}

type GetExperimentByNameResponse struct {
	// Experiment details.
	Experiment *Experiment `json:"experiment,omitempty"`
}

type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId string ` url:"experiment_id,omitempty"`
}

type GetExperimentResponse struct {
	// Experiment details.
	Experiment *Experiment `json:"experiment,omitempty"`
	// A collection of active runs in the experiment. Note: this may not contain
	// all of the experiment&#39;s active runs. This field is deprecated. Please use
	// the &#34;Search Runs&#34; API to fetch runs within an experiment.
	Runs []RunInfo `json:"runs,omitempty"`
}

type GetMetricHistoryRequest struct {
	// Name of the metric.
	MetricKey string ` url:"metric_key,omitempty"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId string ` url:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run from which to fetch metric
	// values. This field will be removed in a future MLflow version.
	RunUuid string ` url:"run_uuid,omitempty"`
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric.
	Metrics []Metric `json:"metrics,omitempty"`
}

type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId string ` url:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run to fetch. This field will
	// be removed in a future MLflow version.
	RunUuid string ` url:"run_uuid,omitempty"`
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	Run *Run `json:"run,omitempty"`
}

type ListArtifactsRequest struct {
	// Token indicating the page of artifact results to fetch
	PageToken string ` url:"page_token,omitempty"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path string ` url:"path,omitempty"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId string ` url:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid string ` url:"run_uuid,omitempty"`
}

type ListArtifactsResponse struct {
	// File location and metadata for artifacts.
	Files []FileInfo `json:"files,omitempty"`
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken string `json:"next_page_token,omitempty"`
	// Root artifact directory for the run.
	RootUri string `json:"root_uri,omitempty"`
}

type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it&#39;ll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults int ` url:"max_results,omitempty"`
	// Token indicating the page of experiments to fetch
	PageToken string ` url:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType string ` url:"view_type,omitempty"`
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`
}

type LogBatchRequest struct {
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

type LogMetricRequest struct {
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

type LogModelRequest struct {
	// MLmodel file in json format.
	ModelJson string `json:"model_json,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id,omitempty"`
}

type LogParamRequest struct {
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

type Param struct {
	// Key identifying this param.
	Key string `json:"key,omitempty"`
	// Value associated with this param.
	Value string `json:"value,omitempty"`
}

type RestoreExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

type RestoreRunRequest struct {
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
	// local path (starting with &#34;/&#34;), or a distributed file system (DFS) path,
	// like ``s3://bucket/directory`` or ``dbfs:/my/directory``. If not set, the
	// local ``./mlruns`` directory is chosen.
	ArtifactUri string `json:"artifact_uri,omitempty"`
	// Unix timestamp of when the run ended in milliseconds.
	EndTime int64 `json:"end_time,omitempty"`
	// The experiment ID.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Current life cycle stage of the experiment : OneOf(&#34;active&#34;, &#34;deleted&#34;)
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
	// and will be removed in a future MLflow release. Use &#39;mlflow.user&#39; tag
	// instead.
	UserId string `json:"user_id,omitempty"`
}

// Current status of the run.
type RunInfoStatus string

const RunInfoStatusRunning RunInfoStatus = `RUNNING`

const RunInfoStatusScheduled RunInfoStatus = `SCHEDULED`

const RunInfoStatusFinished RunInfoStatus = `FINISHED`

const RunInfoStatusFailed RunInfoStatus = `FAILED`

const RunInfoStatusKilled RunInfoStatus = `KILLED`

type RunTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`
}

type SearchExperimentsRequest struct {
	// String representing a SQL filter condition (e.g. &#34;name ILIKE
	// &#39;my-experiment%&#39;&#34;)
	Filter string `json:"filter,omitempty"`
	// Maximum number of experiments desired. Max threshold is 3000.
	MaxResults int64 `json:"max_results,omitempty"`
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional &#34;DESC&#34; or &#34;ASC&#34;
	// annotation, where &#34;ASC&#34; is the default. Tiebreaks are done by experiment
	// id DESC.
	OrderBy []string `json:"order_by,omitempty"`
	// Token indicating the page of experiments to fetch
	PageToken string `json:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType SearchExperimentsRequestViewType `json:"view_type,omitempty"`
}

// Qualifier for type of experiments to be returned. If unspecified, return only
// active experiments.
type SearchExperimentsRequestViewType string

const SearchExperimentsRequestViewTypeActiveOnly SearchExperimentsRequestViewType = `ACTIVE_ONLY`

const SearchExperimentsRequestViewTypeDeletedOnly SearchExperimentsRequestViewType = `DELETED_ONLY`

const SearchExperimentsRequestViewTypeAll SearchExperimentsRequestViewType = `ALL`

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`
}

type SearchRunsRequest struct {
	// List of experiment IDs to search over.
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	// A filter expression over params, metrics, and tags, that allows returning
	// a subset of runs. The syntax is a subset of SQL that supports ANDing
	// together binary operations between a param, metric, or tag and a
	// constant. Example: ``metrics.rmse &lt; 1 and params.model_class =
	// &#39;LogisticRegression&#39;`` You can select columns with special characters
	// (hyphen, space, period, etc.) by using double quotes: ``metrics.&#34;model
	// class&#34; = &#39;LinearRegression&#39; and tags.&#34;user-name&#34; = &#39;Tomas&#39;`` Supported
	// operators are ``=``, ``!=``, ``&gt;``, ``&gt;=``, ``&lt;``, and ``&lt;=``.
	Filter string `json:"filter,omitempty"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults int `json:"max_results,omitempty"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional &#34;DESC&#34; or &#34;ASC&#34; annotation, where &#34;ASC&#34; is the
	// default. Example: [&#34;params.input DESC&#34;, &#34;metrics.alpha ASC&#34;,
	// &#34;metrics.rmse&#34;] Tiebreaks are done by start_time DESC followed by run_id
	// for runs with the same start time (and this is the default ordering
	// criterion if order_by is not provided).
	OrderBy []string `json:"order_by,omitempty"`

	PageToken string `json:"page_token,omitempty"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType SearchRunsRequestRunViewType `json:"run_view_type,omitempty"`
}

// Whether to display only active, only deleted, or all runs. Defaults to only
// active runs.
type SearchRunsRequestRunViewType string

const SearchRunsRequestRunViewTypeActiveOnly SearchRunsRequestRunViewType = `ACTIVE_ONLY`

const SearchRunsRequestRunViewTypeDeletedOnly SearchRunsRequestRunViewType = `DELETED_ONLY`

const SearchRunsRequestRunViewTypeAll SearchRunsRequestRunViewType = `ALL`

type SearchRunsResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`
	// Runs that match the search criteria.
	Runs []Run `json:"runs,omitempty"`
}

type SetExperimentTagRequest struct {
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

type SetTagRequest struct {
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

type UpdateExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
	// If provided, the experiment&#39;s name is changed to the new name. The new
	// name must be unique.
	NewName string `json:"new_name,omitempty"`
}

type UpdateRunRequest struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime int64 `json:"end_time,omitempty"`
	// ID of the run to update. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run to update.. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Updated status of the run.
	Status UpdateRunRequestStatus `json:"status,omitempty"`
}

// Updated status of the run.
type UpdateRunRequestStatus string

const UpdateRunRequestStatusRunning UpdateRunRequestStatus = `RUNNING`

const UpdateRunRequestStatusScheduled UpdateRunRequestStatus = `SCHEDULED`

const UpdateRunRequestStatusFinished UpdateRunRequestStatus = `FINISHED`

const UpdateRunRequestStatusFailed UpdateRunRequestStatus = `FAILED`

const UpdateRunRequestStatusKilled UpdateRunRequestStatus = `KILLED`

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo *RunInfo `json:"run_info,omitempty"`
}
