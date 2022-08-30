// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflow

import (
	"context"
	

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type MlflowService interface {
    // Create an experiment with a name. Returns the ID of the newly created
    // experiment. Validates that another experiment with the same name does not
    // already exist and fails if another experiment with the same name already
    // exists. Throws ``RESOURCE_ALREADY_EXISTS`` if a experiment with the given
    // name exists.
    CreateExperiment(ctx context.Context, createExperimentRequest CreateExperimentRequest) (*CreateExperimentResponse, error)
    // Create a new run within an experiment. A run is usually a single
    // execution of a machine learning or data ETL pipeline. MLflow uses runs to
    // track :ref:`mlflowParam`, :ref:`mlflowMetric`, and :ref:`mlflowRunTag`
    // associated with a single execution.
    CreateRun(ctx context.Context, createRunRequest CreateRunRequest) (*CreateRunResponse, error)
    // Mark an experiment and associated metadata, runs, metrics, params, and
    // tags for deletion. If the experiment uses FileStore, artifacts associated
    // with experiment are also deleted.
    DeleteExperiment(ctx context.Context, deleteExperimentRequest DeleteExperimentRequest) error
    // Mark a run for deletion.
    DeleteRun(ctx context.Context, deleteRunRequest DeleteRunRequest) error
    // Delete a tag on a run. Tags are run metadata that can be updated during a
    // run and after a run completes.
    DeleteTag(ctx context.Context, deleteTagRequest DeleteTagRequest) error
    // Get metadata for an experiment. This method works on deleted experiments.
    GetExperiment(ctx context.Context, getExperimentRequest GetExperimentRequest) (*GetExperimentResponse, error)
    // Get metadata for an experiment. This endpoint will return deleted
    // experiments, but prefers the active experiment if an active and deleted
    // experiment share the same name. If multiple deleted experiments share the
    // same name, the API will return one of them. Throws
    // ``RESOURCE_DOES_NOT_EXIST`` if no experiment with the specified name
    // exists.
    GetExperimentByName(ctx context.Context, getExperimentByNameRequest GetExperimentByNameRequest) (*GetExperimentByNameResponse, error)
    // Get a list of all values for the specified metric for a given run.
    GetMetricHistory(ctx context.Context, getMetricHistoryRequest GetMetricHistoryRequest) (*GetMetricHistoryResponse, error)
    // Get metadata, metrics, params, and tags for a run. In the case where
    // multiple metrics with the same key are logged for a run, return only the
    // value with the latest timestamp. If there are multiple values with the
    // latest timestamp, return the maximum of these values.
    GetRun(ctx context.Context, getRunRequest GetRunRequest) (*GetRunResponse, error)
    // List artifacts for a run. Takes an optional ``artifact_path`` prefix
    // which if specified, the response contains only artifacts with the
    // specified prefix.
    ListArtifacts(ctx context.Context, listArtifactsRequest ListArtifactsRequest) (*ListArtifactsResponse, error)
    // Get a list of all experiments.
    ListExperiments(ctx context.Context, listExperimentsRequest ListExperimentsRequest) (*ListExperimentsResponse, error)
    // Log a batch of metrics, params, and tags for a run. If any data failed to
    // be persisted, the server will respond with an error (non-200 status
    // code). In case of error (due to internal server error or an invalid
    // request), partial data may be written. You can write metrics, params, and
    // tags in interleaving fashion, but within a given entity type are
    // guaranteed to follow the order specified in the request body. That is,
    // for an API request like .. code-block:: json { &#34;run_id&#34;:
    // &#34;2a14ed5c6a87499199e0106c3501eab8&#34;, &#34;metrics&#34;: [ {&#34;key&#34;: &#34;mae&#34;, &#34;value&#34;:
    // 2.5, &#34;timestamp&#34;: 1552550804}, {&#34;key&#34;: &#34;rmse&#34;, &#34;value&#34;: 2.7, &#34;timestamp&#34;:
    // 1552550804}, ], &#34;params&#34;: [ {&#34;key&#34;: &#34;model_class&#34;, &#34;value&#34;:
    // &#34;LogisticRegression&#34;}, ] } the server is guaranteed to write metric
    // &#34;rmse&#34; after &#34;mae&#34;, though it may write param &#34;model_class&#34; before both
    // metrics, after &#34;mae&#34;, or after both metrics. The overwrite behavior for
    // metrics, params, and tags is as follows: - Metrics: metric values are
    // never overwritten. Logging a metric (key, value, timestamp) appends to
    // the set of values for the metric with the provided key. - Tags: tag
    // values can be overwritten by successive writes to the same tag key. That
    // is, if multiple tag values with the same key are provided in the same API
    // request, the last-provided tag value is written. Logging the same tag
    // (key, value) is permitted - that is, logging a tag is idempotent. -
    // Params: once written, param values cannot be changed (attempting to
    // overwrite a param value will result in an error). However, logging the
    // same param (key, value) is permitted - that is, logging a param is
    // idempotent. Request Limits -------------- A single JSON-serialized API
    // request may be up to 1 MB in size and contain: - No more than 1000
    // metrics, params, and tags in total - Up to 1000 metrics - Up to 100
    // params - Up to 100 tags For example, a valid request might contain 900
    // metrics, 50 params, and 50 tags, but logging 900 metrics, 50 params, and
    // 51 tags is invalid. The following limits also apply to metric, param, and
    // tag keys and values: - Metric, param, and tag keys can be up to 250
    // characters in length - Param and tag values can be up to 250 characters
    // in length
    LogBatch(ctx context.Context, logBatchRequest LogBatchRequest) error
    // Log a metric for a run. A metric is a key-value pair (string key, float
    // value) with an associated timestamp. Examples include the various metrics
    // that represent ML model accuracy. A metric can be logged multiple times.
    LogMetric(ctx context.Context, logMetricRequest LogMetricRequest) error
    // .. note:: Experimental: This API may change or be removed in a future
    // release without warning.
    LogModel(ctx context.Context, logModelRequest LogModelRequest) error
    // Log a param used for a run. A param is a key-value pair (string key,
    // string value). Examples include hyperparameters used for ML model
    // training and constant dates and values used in an ETL pipeline. A param
    // can be logged only once for a run.
    LogParam(ctx context.Context, logParamRequest LogParamRequest) error
    // Restore an experiment marked for deletion. This also restores associated
    // metadata, runs, metrics, params, and tags. If experiment uses FileStore,
    // underlying artifacts associated with experiment are also restored. Throws
    // ``RESOURCE_DOES_NOT_EXIST`` if experiment was never created or was
    // permanently deleted.
    RestoreExperiment(ctx context.Context, restoreExperimentRequest RestoreExperimentRequest) error
    // Restore a deleted run.
    RestoreRun(ctx context.Context, restoreRunRequest RestoreRunRequest) error
    // Search for experiments that satisfy specified search criteria.
    SearchExperiments(ctx context.Context, searchExperimentsRequest SearchExperimentsRequest) (*SearchExperimentsResponse, error)
    // Search for runs that satisfy expressions. Search expressions can use
    // :ref:`mlflowMetric` and :ref:`mlflowParam` keys.
    SearchRuns(ctx context.Context, searchRunsRequest SearchRunsRequest) (*SearchRunsResponse, error)
    // Set a tag on an experiment. Experiment tags are metadata that can be
    // updated.
    SetExperimentTag(ctx context.Context, setExperimentTagRequest SetExperimentTagRequest) error
    // Set a tag on a run. Tags are run metadata that can be updated during a
    // run and after a run completes.
    SetTag(ctx context.Context, setTagRequest SetTagRequest) error
    // Update experiment metadata.
    UpdateExperiment(ctx context.Context, updateExperimentRequest UpdateExperimentRequest) error
    // Update run metadata.
    UpdateRun(ctx context.Context, updateRunRequest UpdateRunRequest) (*UpdateRunResponse, error)
}

func New(client *client.DatabricksClient) MlflowService {
	return &MlflowAPI{
		client: client,
	}
}

type MlflowAPI struct {
	client *client.DatabricksClient
}

// Create an experiment with a name. Returns the ID of the newly created
// experiment. Validates that another experiment with the same name does not
// already exist and fails if another experiment with the same name already
// exists. Throws ``RESOURCE_ALREADY_EXISTS`` if a experiment with the given
// name exists.
func (a *MlflowAPI) CreateExperiment(ctx context.Context, request CreateExperimentRequest) (*CreateExperimentResponse, error) {
	var createExperimentResponse CreateExperimentResponse
	path := "/api/2.0/mlflow/experiments/create"
	err := a.client.Post(ctx, path, request, &createExperimentResponse)
	return &createExperimentResponse, err
}

// Create a new run within an experiment. A run is usually a single execution of
// a machine learning or data ETL pipeline. MLflow uses runs to track
// :ref:`mlflowParam`, :ref:`mlflowMetric`, and :ref:`mlflowRunTag` associated
// with a single execution.
func (a *MlflowAPI) CreateRun(ctx context.Context, request CreateRunRequest) (*CreateRunResponse, error) {
	var createRunResponse CreateRunResponse
	path := "/api/2.0/mlflow/runs/create"
	err := a.client.Post(ctx, path, request, &createRunResponse)
	return &createRunResponse, err
}

// Mark an experiment and associated metadata, runs, metrics, params, and tags
// for deletion. If the experiment uses FileStore, artifacts associated with
// experiment are also deleted.
func (a *MlflowAPI) DeleteExperiment(ctx context.Context, request DeleteExperimentRequest) error {
	path := "/api/2.0/mlflow/experiments/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Mark a run for deletion.
func (a *MlflowAPI) DeleteRun(ctx context.Context, request DeleteRunRequest) error {
	path := "/api/2.0/mlflow/runs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Delete a tag on a run. Tags are run metadata that can be updated during a run
// and after a run completes.
func (a *MlflowAPI) DeleteTag(ctx context.Context, request DeleteTagRequest) error {
	path := "/api/2.0/mlflow/runs/delete-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Get metadata for an experiment. This method works on deleted experiments.
func (a *MlflowAPI) GetExperiment(ctx context.Context, request GetExperimentRequest) (*GetExperimentResponse, error) {
	var getExperimentResponse GetExperimentResponse
	path := "/api/2.0/mlflow/experiments/get"
	err := a.client.Get(ctx, path, request, &getExperimentResponse)
	return &getExperimentResponse, err
}

// Get metadata for an experiment. This endpoint will return deleted
// experiments, but prefers the active experiment if an active and deleted
// experiment share the same name. If multiple deleted experiments share the
// same name, the API will return one of them. Throws
// ``RESOURCE_DOES_NOT_EXIST`` if no experiment with the specified name exists.
func (a *MlflowAPI) GetExperimentByName(ctx context.Context, request GetExperimentByNameRequest) (*GetExperimentByNameResponse, error) {
	var getExperimentByNameResponse GetExperimentByNameResponse
	path := "/api/2.0/mlflow/experiments/get-by-name"
	err := a.client.Get(ctx, path, request, &getExperimentByNameResponse)
	return &getExperimentByNameResponse, err
}

// Get a list of all values for the specified metric for a given run.
func (a *MlflowAPI) GetMetricHistory(ctx context.Context, request GetMetricHistoryRequest) (*GetMetricHistoryResponse, error) {
	var getMetricHistoryResponse GetMetricHistoryResponse
	path := "/api/2.0/mlflow/metrics/get-history"
	err := a.client.Get(ctx, path, request, &getMetricHistoryResponse)
	return &getMetricHistoryResponse, err
}

// Get metadata, metrics, params, and tags for a run. In the case where multiple
// metrics with the same key are logged for a run, return only the value with
// the latest timestamp. If there are multiple values with the latest timestamp,
// return the maximum of these values.
func (a *MlflowAPI) GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	var getRunResponse GetRunResponse
	path := "/api/2.0/mlflow/runs/get"
	err := a.client.Get(ctx, path, request, &getRunResponse)
	return &getRunResponse, err
}

// List artifacts for a run. Takes an optional ``artifact_path`` prefix which if
// specified, the response contains only artifacts with the specified prefix.
func (a *MlflowAPI) ListArtifacts(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error) {
	var listArtifactsResponse ListArtifactsResponse
	path := "/api/2.0/mlflow/artifacts/list"
	err := a.client.Get(ctx, path, request, &listArtifactsResponse)
	return &listArtifactsResponse, err
}

// Get a list of all experiments.
func (a *MlflowAPI) ListExperiments(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error) {
	var listExperimentsResponse ListExperimentsResponse
	path := "/api/2.0/mlflow/experiments/list"
	err := a.client.Get(ctx, path, request, &listExperimentsResponse)
	return &listExperimentsResponse, err
}

// Log a batch of metrics, params, and tags for a run. If any data failed to be
// persisted, the server will respond with an error (non-200 status code). In
// case of error (due to internal server error or an invalid request), partial
// data may be written. You can write metrics, params, and tags in interleaving
// fashion, but within a given entity type are guaranteed to follow the order
// specified in the request body. That is, for an API request like ..
// code-block:: json { &#34;run_id&#34;: &#34;2a14ed5c6a87499199e0106c3501eab8&#34;, &#34;metrics&#34;:
// [ {&#34;key&#34;: &#34;mae&#34;, &#34;value&#34;: 2.5, &#34;timestamp&#34;: 1552550804}, {&#34;key&#34;: &#34;rmse&#34;,
// &#34;value&#34;: 2.7, &#34;timestamp&#34;: 1552550804}, ], &#34;params&#34;: [ {&#34;key&#34;: &#34;model_class&#34;,
// &#34;value&#34;: &#34;LogisticRegression&#34;}, ] } the server is guaranteed to write metric
// &#34;rmse&#34; after &#34;mae&#34;, though it may write param &#34;model_class&#34; before both
// metrics, after &#34;mae&#34;, or after both metrics. The overwrite behavior for
// metrics, params, and tags is as follows: - Metrics: metric values are never
// overwritten. Logging a metric (key, value, timestamp) appends to the set of
// values for the metric with the provided key. - Tags: tag values can be
// overwritten by successive writes to the same tag key. That is, if multiple
// tag values with the same key are provided in the same API request, the
// last-provided tag value is written. Logging the same tag (key, value) is
// permitted - that is, logging a tag is idempotent. - Params: once written,
// param values cannot be changed (attempting to overwrite a param value will
// result in an error). However, logging the same param (key, value) is
// permitted - that is, logging a param is idempotent. Request Limits
// -------------- A single JSON-serialized API request may be up to 1 MB in size
// and contain: - No more than 1000 metrics, params, and tags in total - Up to
// 1000 metrics - Up to 100 params - Up to 100 tags For example, a valid request
// might contain 900 metrics, 50 params, and 50 tags, but logging 900 metrics,
// 50 params, and 51 tags is invalid. The following limits also apply to metric,
// param, and tag keys and values: - Metric, param, and tag keys can be up to
// 250 characters in length - Param and tag values can be up to 250 characters
// in length
func (a *MlflowAPI) LogBatch(ctx context.Context, request LogBatchRequest) error {
	path := "/api/2.0/mlflow/runs/log-batch"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Log a metric for a run. A metric is a key-value pair (string key, float
// value) with an associated timestamp. Examples include the various metrics
// that represent ML model accuracy. A metric can be logged multiple times.
func (a *MlflowAPI) LogMetric(ctx context.Context, request LogMetricRequest) error {
	path := "/api/2.0/mlflow/runs/log-metric"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// .. note:: Experimental: This API may change or be removed in a future release
// without warning.
func (a *MlflowAPI) LogModel(ctx context.Context, request LogModelRequest) error {
	path := "/api/2.0/mlflow/runs/log-model"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Log a param used for a run. A param is a key-value pair (string key, string
// value). Examples include hyperparameters used for ML model training and
// constant dates and values used in an ETL pipeline. A param can be logged only
// once for a run.
func (a *MlflowAPI) LogParam(ctx context.Context, request LogParamRequest) error {
	path := "/api/2.0/mlflow/runs/log-parameter"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Restore an experiment marked for deletion. This also restores associated
// metadata, runs, metrics, params, and tags. If experiment uses FileStore,
// underlying artifacts associated with experiment are also restored. Throws
// ``RESOURCE_DOES_NOT_EXIST`` if experiment was never created or was
// permanently deleted.
func (a *MlflowAPI) RestoreExperiment(ctx context.Context, request RestoreExperimentRequest) error {
	path := "/api/2.0/mlflow/experiments/restore"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Restore a deleted run.
func (a *MlflowAPI) RestoreRun(ctx context.Context, request RestoreRunRequest) error {
	path := "/api/2.0/mlflow/runs/restore"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Search for experiments that satisfy specified search criteria.
func (a *MlflowAPI) SearchExperiments(ctx context.Context, request SearchExperimentsRequest) (*SearchExperimentsResponse, error) {
	var searchExperimentsResponse SearchExperimentsResponse
	path := "/api/2.0/mlflow/experiments/search"
	err := a.client.Post(ctx, path, request, &searchExperimentsResponse)
	return &searchExperimentsResponse, err
}

// Search for runs that satisfy expressions. Search expressions can use
// :ref:`mlflowMetric` and :ref:`mlflowParam` keys.
func (a *MlflowAPI) SearchRuns(ctx context.Context, request SearchRunsRequest) (*SearchRunsResponse, error) {
	var searchRunsResponse SearchRunsResponse
	path := "/api/2.0/mlflow/runs/search"
	err := a.client.Post(ctx, path, request, &searchRunsResponse)
	return &searchRunsResponse, err
}

// Set a tag on an experiment. Experiment tags are metadata that can be updated.
func (a *MlflowAPI) SetExperimentTag(ctx context.Context, request SetExperimentTagRequest) error {
	path := "/api/2.0/mlflow/experiments/set-experiment-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Set a tag on a run. Tags are run metadata that can be updated during a run
// and after a run completes.
func (a *MlflowAPI) SetTag(ctx context.Context, request SetTagRequest) error {
	path := "/api/2.0/mlflow/runs/set-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Update experiment metadata.
func (a *MlflowAPI) UpdateExperiment(ctx context.Context, request UpdateExperimentRequest) error {
	path := "/api/2.0/mlflow/experiments/update"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Update run metadata.
func (a *MlflowAPI) UpdateRun(ctx context.Context, request UpdateRunRequest) (*UpdateRunResponse, error) {
	var updateRunResponse UpdateRunResponse
	path := "/api/2.0/mlflow/runs/update"
	err := a.client.Post(ctx, path, request, &updateRunResponse)
	return &updateRunResponse, err
}

