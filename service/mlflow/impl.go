// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflow

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Experiments API methods
type experimentsImpl struct {
	client *client.DatabricksClient
}

func (a *experimentsImpl) Create(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {
	var createExperimentResponse CreateExperimentResponse
	path := "/api/2.0/mlflow/experiments/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createExperimentResponse)
	return &createExperimentResponse, err
}

func (a *experimentsImpl) Delete(ctx context.Context, request DeleteExperiment) error {
	path := "/api/2.0/mlflow/experiments/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *experimentsImpl) Get(ctx context.Context, request GetExperimentRequest) (*Experiment, error) {
	var experiment Experiment
	path := "/api/2.0/mlflow/experiments/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &experiment)
	return &experiment, err
}

func (a *experimentsImpl) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error) {
	var getExperimentByNameResponse GetExperimentByNameResponse
	path := "/api/2.0/mlflow/experiments/get-by-name"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getExperimentByNameResponse)
	return &getExperimentByNameResponse, err
}

func (a *experimentsImpl) List(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error) {
	var listExperimentsResponse ListExperimentsResponse
	path := "/api/2.0/mlflow/experiments/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listExperimentsResponse)
	return &listExperimentsResponse, err
}

func (a *experimentsImpl) Restore(ctx context.Context, request RestoreExperiment) error {
	path := "/api/2.0/mlflow/experiments/restore"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *experimentsImpl) Search(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error) {
	var searchExperimentsResponse SearchExperimentsResponse
	path := "/api/2.0/mlflow/experiments/search"
	err := a.client.Do(ctx, http.MethodPost, path, request, &searchExperimentsResponse)
	return &searchExperimentsResponse, err
}

func (a *experimentsImpl) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {
	path := "/api/2.0/mlflow/experiments/set-experiment-tag"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *experimentsImpl) Update(ctx context.Context, request UpdateExperiment) error {
	path := "/api/2.0/mlflow/experiments/update"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

// unexported type that holds implementations of just MLflowArtifacts API methods
type mLflowArtifactsImpl struct {
	client *client.DatabricksClient
}

func (a *mLflowArtifactsImpl) List(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error) {
	var listArtifactsResponse ListArtifactsResponse
	path := "/api/2.0/mlflow/artifacts/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listArtifactsResponse)
	return &listArtifactsResponse, err
}

// unexported type that holds implementations of just MLflowDatabricks API methods
type mLflowDatabricksImpl struct {
	client *client.DatabricksClient
}

func (a *mLflowDatabricksImpl) Get(ctx context.Context, request GetRequest) (*GetResponse, error) {
	var getResponse GetResponse
	path := "/api/2.0/mlflow/databricks/registered-models/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getResponse)
	return &getResponse, err
}

func (a *mLflowDatabricksImpl) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {
	var transitionStageResponse TransitionStageResponse
	path := "/api/2.0/mlflow/databricks/model-versions/transition-stage"
	err := a.client.Do(ctx, http.MethodPost, path, request, &transitionStageResponse)
	return &transitionStageResponse, err
}

// unexported type that holds implementations of just MLflowMetrics API methods
type mLflowMetricsImpl struct {
	client *client.DatabricksClient
}

func (a *mLflowMetricsImpl) GetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {
	var getMetricHistoryResponse GetMetricHistoryResponse
	path := "/api/2.0/mlflow/metrics/get-history"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getMetricHistoryResponse)
	return &getMetricHistoryResponse, err
}

// unexported type that holds implementations of just MLflowRuns API methods
type mLflowRunsImpl struct {
	client *client.DatabricksClient
}

func (a *mLflowRunsImpl) Create(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {
	var createRunResponse CreateRunResponse
	path := "/api/2.0/mlflow/runs/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createRunResponse)
	return &createRunResponse, err
}

func (a *mLflowRunsImpl) Delete(ctx context.Context, request DeleteRun) error {
	path := "/api/2.0/mlflow/runs/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *mLflowRunsImpl) DeleteTag(ctx context.Context, request DeleteTag) error {
	path := "/api/2.0/mlflow/runs/delete-tag"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *mLflowRunsImpl) Get(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	var getRunResponse GetRunResponse
	path := "/api/2.0/mlflow/runs/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getRunResponse)
	return &getRunResponse, err
}

func (a *mLflowRunsImpl) LogBatch(ctx context.Context, request LogBatch) error {
	path := "/api/2.0/mlflow/runs/log-batch"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *mLflowRunsImpl) LogMetric(ctx context.Context, request LogMetric) error {
	path := "/api/2.0/mlflow/runs/log-metric"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *mLflowRunsImpl) LogModel(ctx context.Context, request LogModel) error {
	path := "/api/2.0/mlflow/runs/log-model"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *mLflowRunsImpl) LogParameter(ctx context.Context, request LogParam) error {
	path := "/api/2.0/mlflow/runs/log-parameter"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *mLflowRunsImpl) Restore(ctx context.Context, request RestoreRun) error {
	path := "/api/2.0/mlflow/runs/restore"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *mLflowRunsImpl) Search(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error) {
	var searchRunsResponse SearchRunsResponse
	path := "/api/2.0/mlflow/runs/search"
	err := a.client.Do(ctx, http.MethodPost, path, request, &searchRunsResponse)
	return &searchRunsResponse, err
}

func (a *mLflowRunsImpl) SetTag(ctx context.Context, request SetTag) error {
	path := "/api/2.0/mlflow/runs/set-tag"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *mLflowRunsImpl) Update(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {
	var updateRunResponse UpdateRunResponse
	path := "/api/2.0/mlflow/runs/update"
	err := a.client.Do(ctx, http.MethodPost, path, request, &updateRunResponse)
	return &updateRunResponse, err
}

// unexported type that holds implementations of just ModelVersionComments API methods
type modelVersionCommentsImpl struct {
	client *client.DatabricksClient
}

func (a *modelVersionCommentsImpl) Create(ctx context.Context, request CreateComment) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/mlflow/comments/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createResponse)
	return &createResponse, err
}

func (a *modelVersionCommentsImpl) Delete(ctx context.Context, request DeleteRequest) error {
	path := "/api/2.0/mlflow/comments/delete"
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *modelVersionCommentsImpl) Update(ctx context.Context, request UpdateComment) (*UpdateResponse, error) {
	var updateResponse UpdateResponse
	path := "/api/2.0/mlflow/comments/update"
	err := a.client.Do(ctx, http.MethodPost, path, request, &updateResponse)
	return &updateResponse, err
}

// unexported type that holds implementations of just ModelVersions API methods
type modelVersionsImpl struct {
	client *client.DatabricksClient
}

func (a *modelVersionsImpl) Create(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	var createModelVersionResponse CreateModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createModelVersionResponse)
	return &createModelVersionResponse, err
}

func (a *modelVersionsImpl) Delete(ctx context.Context, request DeleteModelVersionRequest) error {
	path := "/api/2.0/mlflow/model-versions/delete"
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *modelVersionsImpl) DeleteTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	path := "/api/2.0/mlflow/model-versions/delete-tag"
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *modelVersionsImpl) Get(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	var getModelVersionResponse GetModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getModelVersionResponse)
	return &getModelVersionResponse, err
}

func (a *modelVersionsImpl) GetDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	var getModelVersionDownloadUriResponse GetModelVersionDownloadUriResponse
	path := "/api/2.0/mlflow/model-versions/get-download-uri"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getModelVersionDownloadUriResponse)
	return &getModelVersionDownloadUriResponse, err
}

func (a *modelVersionsImpl) Search(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
	var searchModelVersionsResponse SearchModelVersionsResponse
	path := "/api/2.0/mlflow/model-versions/search"
	err := a.client.Do(ctx, http.MethodGet, path, request, &searchModelVersionsResponse)
	return &searchModelVersionsResponse, err
}

func (a *modelVersionsImpl) SetTag(ctx context.Context, request SetModelVersionTagRequest) error {
	path := "/api/2.0/mlflow/model-versions/set-tag"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *modelVersionsImpl) TransitionStage(ctx context.Context, request TransitionModelVersionStage) (*TransitionModelVersionStageResponse, error) {
	var transitionModelVersionStageResponse TransitionModelVersionStageResponse
	path := "/api/2.0/mlflow/model-versions/transition-stage"
	err := a.client.Do(ctx, http.MethodPost, path, request, &transitionModelVersionStageResponse)
	return &transitionModelVersionStageResponse, err
}

func (a *modelVersionsImpl) Update(ctx context.Context, request UpdateModelVersionRequest) error {
	path := "/api/2.0/mlflow/model-versions/update"
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just RegisteredModels API methods
type registeredModelsImpl struct {
	client *client.DatabricksClient
}

func (a *registeredModelsImpl) Create(ctx context.Context, request CreateRegisteredModelRequest) (*CreateRegisteredModelResponse, error) {
	var createRegisteredModelResponse CreateRegisteredModelResponse
	path := "/api/2.0/mlflow/registered-models/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createRegisteredModelResponse)
	return &createRegisteredModelResponse, err
}

func (a *registeredModelsImpl) Delete(ctx context.Context, request DeleteRegisteredModelRequest) error {
	path := "/api/2.0/mlflow/registered-models/delete"
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *registeredModelsImpl) DeleteTag(ctx context.Context, request DeleteRegisteredModelTagRequest) error {
	path := "/api/2.0/mlflow/registered-models/delete-tag"
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *registeredModelsImpl) Get(ctx context.Context, request GetRegisteredModelRequest) (*GetRegisteredModelResponse, error) {
	var getRegisteredModelResponse GetRegisteredModelResponse
	path := "/api/2.0/mlflow/registered-models/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getRegisteredModelResponse)
	return &getRegisteredModelResponse, err
}

func (a *registeredModelsImpl) GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
	var getLatestVersionsResponse GetLatestVersionsResponse
	path := "/api/2.0/mlflow/registered-models/get-latest-versions"
	err := a.client.Do(ctx, http.MethodPost, path, request, &getLatestVersionsResponse)
	return &getLatestVersionsResponse, err
}

func (a *registeredModelsImpl) List(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error) {
	var listRegisteredModelsResponse ListRegisteredModelsResponse
	path := "/api/2.0/mlflow/registered-models/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listRegisteredModelsResponse)
	return &listRegisteredModelsResponse, err
}

func (a *registeredModelsImpl) Rename(ctx context.Context, request RenameRegisteredModelRequest) (*RenameRegisteredModelResponse, error) {
	var renameRegisteredModelResponse RenameRegisteredModelResponse
	path := "/api/2.0/mlflow/registered-models/rename"
	err := a.client.Do(ctx, http.MethodPost, path, request, &renameRegisteredModelResponse)
	return &renameRegisteredModelResponse, err
}

func (a *registeredModelsImpl) Search(ctx context.Context, request SearchRegisteredModelsRequest) (*SearchRegisteredModelsResponse, error) {
	var searchRegisteredModelsResponse SearchRegisteredModelsResponse
	path := "/api/2.0/mlflow/registered-models/search"
	err := a.client.Do(ctx, http.MethodGet, path, request, &searchRegisteredModelsResponse)
	return &searchRegisteredModelsResponse, err
}

func (a *registeredModelsImpl) SetTag(ctx context.Context, request SetRegisteredModelTagRequest) error {
	path := "/api/2.0/mlflow/registered-models/set-tag"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *registeredModelsImpl) Update(ctx context.Context, request UpdateRegisteredModelRequest) error {
	path := "/api/2.0/mlflow/registered-models/update"
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just RegistryWebhooks API methods
type registryWebhooksImpl struct {
	client *client.DatabricksClient
}

func (a *registryWebhooksImpl) Create(ctx context.Context, request CreateRegistryWebhook) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/mlflow/registry-webhooks/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createResponse)
	return &createResponse, err
}

func (a *registryWebhooksImpl) Delete(ctx context.Context, request DeleteRequest) error {
	path := "/api/2.0/mlflow/registry-webhooks/delete"
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *registryWebhooksImpl) List(ctx context.Context, request ListRequest) (*ListRegistryWebhooks, error) {
	var listRegistryWebhooks ListRegistryWebhooks
	path := "/api/2.0/mlflow/registry-webhooks/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listRegistryWebhooks)
	return &listRegistryWebhooks, err
}

func (a *registryWebhooksImpl) Test(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	var testRegistryWebhookResponse TestRegistryWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/test"
	err := a.client.Do(ctx, http.MethodPost, path, request, &testRegistryWebhookResponse)
	return &testRegistryWebhookResponse, err
}

func (a *registryWebhooksImpl) Update(ctx context.Context, request UpdateRegistryWebhook) error {
	path := "/api/2.0/mlflow/registry-webhooks/update"
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just TransitionRequests API methods
type transitionRequestsImpl struct {
	client *client.DatabricksClient
}

func (a *transitionRequestsImpl) Approve(ctx context.Context, request ApproveTransitionRequest) (*ApproveResponse, error) {
	var approveResponse ApproveResponse
	path := "/api/2.0/mlflow/transition-requests/approve"
	err := a.client.Do(ctx, http.MethodPost, path, request, &approveResponse)
	return &approveResponse, err
}

func (a *transitionRequestsImpl) Create(ctx context.Context, request CreateTransitionRequest) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/mlflow/transition-requests/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createResponse)
	return &createResponse, err
}

func (a *transitionRequestsImpl) Delete(ctx context.Context, request DeleteRequest) error {
	path := "/api/2.0/mlflow/transition-requests/delete"
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *transitionRequestsImpl) List(ctx context.Context, request ListRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/mlflow/transition-requests/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listResponse)
	return &listResponse, err
}

func (a *transitionRequestsImpl) Reject(ctx context.Context, request RejectTransitionRequest) (*RejectResponse, error) {
	var rejectResponse RejectResponse
	path := "/api/2.0/mlflow/transition-requests/reject"
	err := a.client.Do(ctx, http.MethodPost, path, request, &rejectResponse)
	return &rejectResponse, err
}
