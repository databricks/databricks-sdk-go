// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflow

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewExperiments(client *client.DatabricksClient) ExperimentsService {
	return &ExperimentsAPI{
		client: client,
	}
}

type ExperimentsAPI struct {
	client *client.DatabricksClient
}

// Create an experiment with a name. Returns the ID of the newly created
// experiment. Validates that another experiment with the same name does not
// already exist and fails if another experiment with the same name already
// exists. Throws “RESOURCE_ALREADY_EXISTS“ if a experiment with the given
// name exists.
func (a *ExperimentsAPI) Create(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {
	var createExperimentResponse CreateExperimentResponse
	path := "/api/2.0/mlflow/experiments/create"
	err := a.client.Post(ctx, path, request, &createExperimentResponse)
	return &createExperimentResponse, err
}

// Mark an experiment and associated metadata, runs, metrics, params, and tags
// for deletion. If the experiment uses FileStore, artifacts associated with
// experiment are also deleted.
func (a *ExperimentsAPI) Delete(ctx context.Context, request DeleteExperiment) error {
	path := "/api/2.0/mlflow/experiments/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Mark an experiment and associated metadata, runs, metrics, params, and tags
// for deletion. If the experiment uses FileStore, artifacts associated with
// experiment are also deleted.
func (a *ExperimentsAPI) DeleteByExperimentId(ctx context.Context, experimentId string) error {
	return a.Delete(ctx, DeleteExperiment{
		ExperimentId: experimentId,
	})
}

// Get metadata for an experiment. This method works on deleted experiments.
func (a *ExperimentsAPI) Get(ctx context.Context, request GetExperimentRequest) (*Experiment, error) {
	var experiment Experiment
	path := "/api/2.0/mlflow/experiments/get"
	err := a.client.Get(ctx, path, request, &experiment)
	return &experiment, err
}

// Get metadata for an experiment. This method works on deleted experiments.
func (a *ExperimentsAPI) GetByExperimentId(ctx context.Context, experimentId string) (*Experiment, error) {
	return a.Get(ctx, GetExperimentRequest{
		ExperimentId: experimentId,
	})
}

// Get metadata for an experiment. This endpoint will return deleted
// experiments, but prefers the active experiment if an active and deleted
// experiment share the same name. If multiple deleted experiments share the
// same name, the API will return one of them. Throws
// “RESOURCE_DOES_NOT_EXIST“ if no experiment with the specified name exists.
func (a *ExperimentsAPI) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error) {
	var getExperimentByNameResponse GetExperimentByNameResponse
	path := "/api/2.0/mlflow/experiments/get-by-name"
	err := a.client.Get(ctx, path, request, &getExperimentByNameResponse)
	return &getExperimentByNameResponse, err
}

// Get metadata for an experiment. This endpoint will return deleted
// experiments, but prefers the active experiment if an active and deleted
// experiment share the same name. If multiple deleted experiments share the
// same name, the API will return one of them. Throws
// “RESOURCE_DOES_NOT_EXIST“ if no experiment with the specified name exists.
func (a *ExperimentsAPI) GetByNameByExperimentName(ctx context.Context, experimentName string) (*GetExperimentByNameResponse, error) {
	return a.GetByName(ctx, GetByNameRequest{
		ExperimentName: experimentName,
	})
}

// Get a list of all experiments.
func (a *ExperimentsAPI) List(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error) {
	var listExperimentsResponse ListExperimentsResponse
	path := "/api/2.0/mlflow/experiments/list"
	err := a.client.Get(ctx, path, request, &listExperimentsResponse)
	return &listExperimentsResponse, err
}

// Restore an experiment marked for deletion. This also restores associated
// metadata, runs, metrics, params, and tags. If experiment uses FileStore,
// underlying artifacts associated with experiment are also restored. Throws
// “RESOURCE_DOES_NOT_EXIST“ if experiment was never created or was
// permanently deleted.
func (a *ExperimentsAPI) Restore(ctx context.Context, request RestoreExperiment) error {
	path := "/api/2.0/mlflow/experiments/restore"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Restore an experiment marked for deletion. This also restores associated
// metadata, runs, metrics, params, and tags. If experiment uses FileStore,
// underlying artifacts associated with experiment are also restored. Throws
// “RESOURCE_DOES_NOT_EXIST“ if experiment was never created or was
// permanently deleted.
func (a *ExperimentsAPI) RestoreByExperimentId(ctx context.Context, experimentId string) error {
	return a.Restore(ctx, RestoreExperiment{
		ExperimentId: experimentId,
	})
}

// Search for experiments that satisfy specified search criteria.
func (a *ExperimentsAPI) Search(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error) {
	var searchExperimentsResponse SearchExperimentsResponse
	path := "/api/2.0/mlflow/experiments/search"
	err := a.client.Post(ctx, path, request, &searchExperimentsResponse)
	return &searchExperimentsResponse, err
}

// Set a tag on an experiment. Experiment tags are metadata that can be updated.
func (a *ExperimentsAPI) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {
	path := "/api/2.0/mlflow/experiments/set-experiment-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Update experiment metadata.
func (a *ExperimentsAPI) Update(ctx context.Context, request UpdateExperiment) error {
	path := "/api/2.0/mlflow/experiments/update"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func NewMLflowArtifacts(client *client.DatabricksClient) MLflowArtifactsService {
	return &MLflowArtifactsAPI{
		client: client,
	}
}

type MLflowArtifactsAPI struct {
	client *client.DatabricksClient
}

// List artifacts for a run. Takes an optional “artifact_path“ prefix which if
// specified, the response contains only artifacts with the specified prefix.
func (a *MLflowArtifactsAPI) List(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error) {
	var listArtifactsResponse ListArtifactsResponse
	path := "/api/2.0/mlflow/artifacts/list"
	err := a.client.Get(ctx, path, request, &listArtifactsResponse)
	return &listArtifactsResponse, err
}

func NewMLflowDatabricks(client *client.DatabricksClient) MLflowDatabricksService {
	return &MLflowDatabricksAPI{
		client: client,
	}
}

type MLflowDatabricksAPI struct {
	client *client.DatabricksClient
}

func (a *MLflowDatabricksAPI) Get(ctx context.Context, request GetRequest) (*GetResponse, error) {
	var getResponse GetResponse
	path := "/api/2.0/mlflow/databricks/registered-models/get"
	err := a.client.Get(ctx, path, request, &getResponse)
	return &getResponse, err
}

func (a *MLflowDatabricksAPI) GetByName(ctx context.Context, name string) (*GetResponse, error) {
	return a.Get(ctx, GetRequest{
		Name: name,
	})
}

// Transition a model version's stage. This is a <Workspace> version of the
// [MLflow
// endpoint](https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage)
// that also accepts a comment associated with the transition to be recorded.
func (a *MLflowDatabricksAPI) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {
	var transitionStageResponse TransitionStageResponse
	path := "/api/2.0/mlflow/databricks/model-versions/transition-stage"
	err := a.client.Post(ctx, path, request, &transitionStageResponse)
	return &transitionStageResponse, err
}

func NewMLflowMetrics(client *client.DatabricksClient) MLflowMetricsService {
	return &MLflowMetricsAPI{
		client: client,
	}
}

type MLflowMetricsAPI struct {
	client *client.DatabricksClient
}

// Get a list of all values for the specified metric for a given run.
func (a *MLflowMetricsAPI) GetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {
	var getMetricHistoryResponse GetMetricHistoryResponse
	path := "/api/2.0/mlflow/metrics/get-history"
	err := a.client.Get(ctx, path, request, &getMetricHistoryResponse)
	return &getMetricHistoryResponse, err
}

func NewMLflowRuns(client *client.DatabricksClient) MLflowRunsService {
	return &MLflowRunsAPI{
		client: client,
	}
}

type MLflowRunsAPI struct {
	client *client.DatabricksClient
}

// Create a new run within an experiment. A run is usually a single execution of
// a machine learning or data ETL pipeline. MLflow uses runs to track
// :ref:`mlflowParam`, :ref:`mlflowMetric`, and :ref:`mlflowRunTag` associated
// with a single execution.
func (a *MLflowRunsAPI) Create(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {
	var createRunResponse CreateRunResponse
	path := "/api/2.0/mlflow/runs/create"
	err := a.client.Post(ctx, path, request, &createRunResponse)
	return &createRunResponse, err
}

// Mark a run for deletion.
func (a *MLflowRunsAPI) Delete(ctx context.Context, request DeleteRun) error {
	path := "/api/2.0/mlflow/runs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Mark a run for deletion.
func (a *MLflowRunsAPI) DeleteByRunId(ctx context.Context, runId string) error {
	return a.Delete(ctx, DeleteRun{
		RunId: runId,
	})
}

// Delete a tag on a run. Tags are run metadata that can be updated during a run
// and after a run completes.
func (a *MLflowRunsAPI) DeleteTag(ctx context.Context, request DeleteTag) error {
	path := "/api/2.0/mlflow/runs/delete-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Get metadata, metrics, params, and tags for a run. In the case where multiple
// metrics with the same key are logged for a run, return only the value with
// the latest timestamp. If there are multiple values with the latest timestamp,
// return the maximum of these values.
func (a *MLflowRunsAPI) Get(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	var getRunResponse GetRunResponse
	path := "/api/2.0/mlflow/runs/get"
	err := a.client.Get(ctx, path, request, &getRunResponse)
	return &getRunResponse, err
}

// Log a batch of metrics, params, and tags for a run. If any data failed to be
// persisted, the server will respond with an error (non-200 status code). In
// case of error (due to internal server error or an invalid request), partial
// data may be written. You can write metrics, params, and tags in interleaving
// fashion, but within a given entity type are guaranteed to follow the order
// specified in the request body. That is, for an API request like ..
// code-block:: json { "run_id": "2a14ed5c6a87499199e0106c3501eab8", "metrics":
// [ {"key": "mae", "value": 2.5, "timestamp": 1552550804}, {"key": "rmse",
// "value": 2.7, "timestamp": 1552550804}, ], "params": [ {"key": "model_class",
// "value": "LogisticRegression"}, ] } the server is guaranteed to write metric
// "rmse" after "mae", though it may write param "model_class" before both
// metrics, after "mae", or after both metrics. The overwrite behavior for
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
func (a *MLflowRunsAPI) LogBatch(ctx context.Context, request LogBatch) error {
	path := "/api/2.0/mlflow/runs/log-batch"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Log a metric for a run. A metric is a key-value pair (string key, float
// value) with an associated timestamp. Examples include the various metrics
// that represent ML model accuracy. A metric can be logged multiple times.
func (a *MLflowRunsAPI) LogMetric(ctx context.Context, request LogMetric) error {
	path := "/api/2.0/mlflow/runs/log-metric"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// .. note:: Experimental: This API may change or be removed in a future release
// without warning.
func (a *MLflowRunsAPI) LogModel(ctx context.Context, request LogModel) error {
	path := "/api/2.0/mlflow/runs/log-model"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Log a param used for a run. A param is a key-value pair (string key, string
// value). Examples include hyperparameters used for ML model training and
// constant dates and values used in an ETL pipeline. A param can be logged only
// once for a run.
func (a *MLflowRunsAPI) LogParameter(ctx context.Context, request LogParam) error {
	path := "/api/2.0/mlflow/runs/log-parameter"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Restore a deleted run.
func (a *MLflowRunsAPI) Restore(ctx context.Context, request RestoreRun) error {
	path := "/api/2.0/mlflow/runs/restore"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Restore a deleted run.
func (a *MLflowRunsAPI) RestoreByRunId(ctx context.Context, runId string) error {
	return a.Restore(ctx, RestoreRun{
		RunId: runId,
	})
}

// Search for runs that satisfy expressions. Search expressions can use
// :ref:`mlflowMetric` and :ref:`mlflowParam` keys.
func (a *MLflowRunsAPI) Search(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error) {
	var searchRunsResponse SearchRunsResponse
	path := "/api/2.0/mlflow/runs/search"
	err := a.client.Post(ctx, path, request, &searchRunsResponse)
	return &searchRunsResponse, err
}

// Set a tag on a run. Tags are run metadata that can be updated during a run
// and after a run completes.
func (a *MLflowRunsAPI) SetTag(ctx context.Context, request SetTag) error {
	path := "/api/2.0/mlflow/runs/set-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Update run metadata.
func (a *MLflowRunsAPI) Update(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {
	var updateRunResponse UpdateRunResponse
	path := "/api/2.0/mlflow/runs/update"
	err := a.client.Post(ctx, path, request, &updateRunResponse)
	return &updateRunResponse, err
}

func NewModelVersionComments(client *client.DatabricksClient) ModelVersionCommentsService {
	return &ModelVersionCommentsAPI{
		client: client,
	}
}

type ModelVersionCommentsAPI struct {
	client *client.DatabricksClient
}

// Make a comment on a model version. A comment can be submitted either by a
// user or programmatically to display relevant information about the model. For
// example, test results or deployment errors.
func (a *ModelVersionCommentsAPI) Create(ctx context.Context, request CreateComment) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/mlflow/comments/create"
	err := a.client.Post(ctx, path, request, &createResponse)
	return &createResponse, err
}

// Delete a comment on a model version.
func (a *ModelVersionCommentsAPI) Delete(ctx context.Context, request DeleteComment) error {
	path := "/api/2.0/mlflow/comments/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a comment on a model version.
func (a *ModelVersionCommentsAPI) DeleteById(ctx context.Context, id string) error {
	return a.Delete(ctx, DeleteComment{
		Id: id,
	})
}

// Edit a comment on a model version.
func (a *ModelVersionCommentsAPI) Update(ctx context.Context, request UpdateComment) (*UpdateResponse, error) {
	var updateResponse UpdateResponse
	path := "/api/2.0/mlflow/comments/update"
	err := a.client.Post(ctx, path, request, &updateResponse)
	return &updateResponse, err
}

func NewModelVersions(client *client.DatabricksClient) ModelVersionsService {
	return &ModelVersionsAPI{
		client: client,
	}
}

type ModelVersionsAPI struct {
	client *client.DatabricksClient
}

func (a *ModelVersionsAPI) Create(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	var createModelVersionResponse CreateModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/create"
	err := a.client.Post(ctx, path, request, &createModelVersionResponse)
	return &createModelVersionResponse, err
}

func (a *ModelVersionsAPI) Delete(ctx context.Context, request DeleteModelVersionRequest) error {
	path := "/api/2.0/mlflow/model-versions/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *ModelVersionsAPI) DeleteTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	path := "/api/2.0/mlflow/model-versions/delete-tag"
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *ModelVersionsAPI) Get(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	var getModelVersionResponse GetModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/get"
	err := a.client.Get(ctx, path, request, &getModelVersionResponse)
	return &getModelVersionResponse, err
}

func (a *ModelVersionsAPI) GetDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	var getModelVersionDownloadUriResponse GetModelVersionDownloadUriResponse
	path := "/api/2.0/mlflow/model-versions/get-download-uri"
	err := a.client.Get(ctx, path, request, &getModelVersionDownloadUriResponse)
	return &getModelVersionDownloadUriResponse, err
}

func (a *ModelVersionsAPI) Search(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
	var searchModelVersionsResponse SearchModelVersionsResponse
	path := "/api/2.0/mlflow/model-versions/search"
	err := a.client.Get(ctx, path, request, &searchModelVersionsResponse)
	return &searchModelVersionsResponse, err
}

func (a *ModelVersionsAPI) SetTag(ctx context.Context, request SetModelVersionTagRequest) error {
	path := "/api/2.0/mlflow/model-versions/set-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *ModelVersionsAPI) TransitionStage(ctx context.Context, request TransitionModelVersionStage) (*TransitionModelVersionStageResponse, error) {
	var transitionModelVersionStageResponse TransitionModelVersionStageResponse
	path := "/api/2.0/mlflow/model-versions/transition-stage"
	err := a.client.Post(ctx, path, request, &transitionModelVersionStageResponse)
	return &transitionModelVersionStageResponse, err
}

func (a *ModelVersionsAPI) Update(ctx context.Context, request UpdateModelVersionRequest) error {
	path := "/api/2.0/mlflow/model-versions/update"
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewRegisteredModels(client *client.DatabricksClient) RegisteredModelsService {
	return &RegisteredModelsAPI{
		client: client,
	}
}

type RegisteredModelsAPI struct {
	client *client.DatabricksClient
}

// Throws “RESOURCE_ALREADY_EXISTS“ if a registered model with the given name
// exists.
func (a *RegisteredModelsAPI) Create(ctx context.Context, request CreateRegisteredModelRequest) (*CreateRegisteredModelResponse, error) {
	var createRegisteredModelResponse CreateRegisteredModelResponse
	path := "/api/2.0/mlflow/registered-models/create"
	err := a.client.Post(ctx, path, request, &createRegisteredModelResponse)
	return &createRegisteredModelResponse, err
}

func (a *RegisteredModelsAPI) Delete(ctx context.Context, request DeleteRegisteredModelRequest) error {
	path := "/api/2.0/mlflow/registered-models/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *RegisteredModelsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.Delete(ctx, DeleteRegisteredModelRequest{
		Name: name,
	})
}

func (a *RegisteredModelsAPI) DeleteTag(ctx context.Context, request DeleteRegisteredModelTagRequest) error {
	path := "/api/2.0/mlflow/registered-models/delete-tag"
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *RegisteredModelsAPI) Get(ctx context.Context, request GetRegisteredModelRequest) (*GetRegisteredModelResponse, error) {
	var getRegisteredModelResponse GetRegisteredModelResponse
	path := "/api/2.0/mlflow/registered-models/get"
	err := a.client.Get(ctx, path, request, &getRegisteredModelResponse)
	return &getRegisteredModelResponse, err
}

func (a *RegisteredModelsAPI) GetByName(ctx context.Context, name string) (*GetRegisteredModelResponse, error) {
	return a.Get(ctx, GetRegisteredModelRequest{
		Name: name,
	})
}

func (a *RegisteredModelsAPI) GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
	var getLatestVersionsResponse GetLatestVersionsResponse
	path := "/api/2.0/mlflow/registered-models/get-latest-versions"
	err := a.client.Post(ctx, path, request, &getLatestVersionsResponse)
	return &getLatestVersionsResponse, err
}

func (a *RegisteredModelsAPI) List(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error) {
	var listRegisteredModelsResponse ListRegisteredModelsResponse
	path := "/api/2.0/mlflow/registered-models/list"
	err := a.client.Get(ctx, path, request, &listRegisteredModelsResponse)
	return &listRegisteredModelsResponse, err
}

func (a *RegisteredModelsAPI) Rename(ctx context.Context, request RenameRegisteredModelRequest) (*RenameRegisteredModelResponse, error) {
	var renameRegisteredModelResponse RenameRegisteredModelResponse
	path := "/api/2.0/mlflow/registered-models/rename"
	err := a.client.Post(ctx, path, request, &renameRegisteredModelResponse)
	return &renameRegisteredModelResponse, err
}

func (a *RegisteredModelsAPI) Search(ctx context.Context, request SearchRegisteredModelsRequest) (*SearchRegisteredModelsResponse, error) {
	var searchRegisteredModelsResponse SearchRegisteredModelsResponse
	path := "/api/2.0/mlflow/registered-models/search"
	err := a.client.Get(ctx, path, request, &searchRegisteredModelsResponse)
	return &searchRegisteredModelsResponse, err
}

func (a *RegisteredModelsAPI) SetTag(ctx context.Context, request SetRegisteredModelTagRequest) error {
	path := "/api/2.0/mlflow/registered-models/set-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *RegisteredModelsAPI) Update(ctx context.Context, request UpdateRegisteredModelRequest) error {
	path := "/api/2.0/mlflow/registered-models/update"
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewRegistryWebhooks(client *client.DatabricksClient) RegistryWebhooksService {
	return &RegistryWebhooksAPI{
		client: client,
	}
}

type RegistryWebhooksAPI struct {
	client *client.DatabricksClient
}

// This endpoint is in Public Preview. Create a registry webhook.
func (a *RegistryWebhooksAPI) Create(ctx context.Context, request CreateRegistryWebhook) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/mlflow/registry-webhooks/create"
	err := a.client.Post(ctx, path, request, &createResponse)
	return &createResponse, err
}

// This endpoint is in Public Preview. Delete a registry webhook.
func (a *RegistryWebhooksAPI) Delete(ctx context.Context, request DeleteRequest) error {
	path := "/api/2.0/mlflow/registry-webhooks/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

// This endpoint is in Public Preview. Delete a registry webhook.
func (a *RegistryWebhooksAPI) DeleteById(ctx context.Context, id string) error {
	return a.Delete(ctx, DeleteRequest{
		Id: id,
	})
}

// This endpoint is in Public Preview. List registry webhooks.
func (a *RegistryWebhooksAPI) List(ctx context.Context, request ListRegistryWebhooks) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/mlflow/registry-webhooks/list"
	err := a.client.Get(ctx, path, request, &listResponse)
	return &listResponse, err
}

// This endpoint is in Public Preview. Test a registry webhook.
func (a *RegistryWebhooksAPI) Test(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	var testRegistryWebhookResponse TestRegistryWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/test"
	err := a.client.Post(ctx, path, request, &testRegistryWebhookResponse)
	return &testRegistryWebhookResponse, err
}

// This endpoint is in Public Preview. Update a registry webhook.
func (a *RegistryWebhooksAPI) Update(ctx context.Context, request UpdateRegistryWebhook) error {
	path := "/api/2.0/mlflow/registry-webhooks/update"
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewTransitionRequests(client *client.DatabricksClient) TransitionRequestsService {
	return &TransitionRequestsAPI{
		client: client,
	}
}

type TransitionRequestsAPI struct {
	client *client.DatabricksClient
}

// Approve model version stage transition request.
func (a *TransitionRequestsAPI) Approve(ctx context.Context, request ApproveTransitionRequest) (*ApproveResponse, error) {
	var approveResponse ApproveResponse
	path := "/api/2.0/mlflow/transition-requests/approve"
	err := a.client.Post(ctx, path, request, &approveResponse)
	return &approveResponse, err
}

// Make a model version stage transition request.
func (a *TransitionRequestsAPI) Create(ctx context.Context, request CreateTransitionRequest) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/mlflow/transition-requests/create"
	err := a.client.Post(ctx, path, request, &createResponse)
	return &createResponse, err
}

// Cancel model version stage transition request.
func (a *TransitionRequestsAPI) Delete(ctx context.Context, request DeleteRequest) error {
	path := "/api/2.0/mlflow/transition-requests/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get all open stage transition requests for the model version.
func (a *TransitionRequestsAPI) List(ctx context.Context, request ListRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/mlflow/transition-requests/list"
	err := a.client.Get(ctx, path, request, &listResponse)
	return &listResponse, err
}

// Reject model version stage transition request.
func (a *TransitionRequestsAPI) Reject(ctx context.Context, request RejectTransitionRequest) (*RejectResponse, error) {
	var rejectResponse RejectResponse
	path := "/api/2.0/mlflow/transition-requests/reject"
	err := a.client.Post(ctx, path, request, &rejectResponse)
	return &rejectResponse, err
}
