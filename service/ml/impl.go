// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Experiments API methods
type experimentsImpl struct {
	client *client.DatabricksClient
}

func (a *experimentsImpl) Client() client.DatabricksClientInterface {
	return a.client
}

func (a *experimentsImpl) CreateExperiment(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {
	var createExperimentResponse CreateExperimentResponse
	path := "/api/2.0/mlflow/experiments/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createExperimentResponse)
	return &createExperimentResponse, err
}

func (a *experimentsImpl) CreateRun(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {
	var createRunResponse CreateRunResponse
	path := "/api/2.0/mlflow/runs/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createRunResponse)
	return &createRunResponse, err
}

func (a *experimentsImpl) DeleteExperiment(ctx context.Context, request DeleteExperiment) error {
	path := "/api/2.0/mlflow/experiments/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	path := "/api/2.0/mlflow/runs/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) DeleteRuns(ctx context.Context, request DeleteRuns) (*DeleteRunsResponse, error) {
	var deleteRunsResponse DeleteRunsResponse
	path := "/api/2.0/mlflow/databricks/runs/delete-runs"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deleteRunsResponse)
	return &deleteRunsResponse, err
}

func (a *experimentsImpl) DeleteTag(ctx context.Context, request DeleteTag) error {
	path := "/api/2.0/mlflow/runs/delete-tag"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentResponse, error) {
	var getExperimentResponse GetExperimentResponse
	path := "/api/2.0/mlflow/experiments/get-by-name"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getExperimentResponse)
	return &getExperimentResponse, err
}

func (a *experimentsImpl) GetExperiment(ctx context.Context, request GetExperimentRequest) (*GetExperimentResponse, error) {
	var getExperimentResponse GetExperimentResponse
	path := "/api/2.0/mlflow/experiments/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getExperimentResponse)
	return &getExperimentResponse, err
}

func (a *experimentsImpl) GetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {
	var getMetricHistoryResponse GetMetricHistoryResponse
	path := "/api/2.0/mlflow/metrics/get-history"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getMetricHistoryResponse)
	return &getMetricHistoryResponse, err
}

func (a *experimentsImpl) GetPermissionLevels(ctx context.Context, request GetExperimentPermissionLevelsRequest) (*GetExperimentPermissionLevelsResponse, error) {
	var getExperimentPermissionLevelsResponse GetExperimentPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v/permissionLevels", request.ExperimentId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getExperimentPermissionLevelsResponse)
	return &getExperimentPermissionLevelsResponse, err
}

func (a *experimentsImpl) GetPermissions(ctx context.Context, request GetExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsImpl) GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	var getRunResponse GetRunResponse
	path := "/api/2.0/mlflow/runs/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getRunResponse)
	return &getRunResponse, err
}

func (a *experimentsImpl) ListArtifacts(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error) {
	var listArtifactsResponse ListArtifactsResponse
	path := "/api/2.0/mlflow/artifacts/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listArtifactsResponse)
	return &listArtifactsResponse, err
}

func (a *experimentsImpl) ListExperiments(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error) {
	var listExperimentsResponse ListExperimentsResponse
	path := "/api/2.0/mlflow/experiments/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listExperimentsResponse)
	return &listExperimentsResponse, err
}

func (a *experimentsImpl) LogBatch(ctx context.Context, request LogBatch) error {
	path := "/api/2.0/mlflow/runs/log-batch"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) LogInputs(ctx context.Context, request LogInputs) error {
	path := "/api/2.0/mlflow/runs/log-inputs"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) LogMetric(ctx context.Context, request LogMetric) error {
	path := "/api/2.0/mlflow/runs/log-metric"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) LogModel(ctx context.Context, request LogModel) error {
	path := "/api/2.0/mlflow/runs/log-model"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) LogParam(ctx context.Context, request LogParam) error {
	path := "/api/2.0/mlflow/runs/log-parameter"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) RestoreExperiment(ctx context.Context, request RestoreExperiment) error {
	path := "/api/2.0/mlflow/experiments/restore"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) RestoreRun(ctx context.Context, request RestoreRun) error {
	path := "/api/2.0/mlflow/runs/restore"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) RestoreRuns(ctx context.Context, request RestoreRuns) (*RestoreRunsResponse, error) {
	var restoreRunsResponse RestoreRunsResponse
	path := "/api/2.0/mlflow/databricks/runs/restore-runs"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &restoreRunsResponse)
	return &restoreRunsResponse, err
}

func (a *experimentsImpl) SearchExperiments(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error) {
	var searchExperimentsResponse SearchExperimentsResponse
	path := "/api/2.0/mlflow/experiments/search"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &searchExperimentsResponse)
	return &searchExperimentsResponse, err
}

func (a *experimentsImpl) SearchRuns(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error) {
	var searchRunsResponse SearchRunsResponse
	path := "/api/2.0/mlflow/runs/search"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &searchRunsResponse)
	return &searchRunsResponse, err
}

func (a *experimentsImpl) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {
	path := "/api/2.0/mlflow/experiments/set-experiment-tag"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) SetPermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsImpl) SetTag(ctx context.Context, request SetTag) error {
	path := "/api/2.0/mlflow/runs/set-tag"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) UpdateExperiment(ctx context.Context, request UpdateExperiment) error {
	path := "/api/2.0/mlflow/experiments/update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *experimentsImpl) UpdatePermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsImpl) UpdateRun(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {
	var updateRunResponse UpdateRunResponse
	path := "/api/2.0/mlflow/runs/update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &updateRunResponse)
	return &updateRunResponse, err
}

// unexported type that holds implementations of just ModelRegistry API methods
type modelRegistryImpl struct {
	client *client.DatabricksClient
}

func (a *modelRegistryImpl) Client() client.DatabricksClientInterface {
	return a.client
}

func (a *modelRegistryImpl) ApproveTransitionRequest(ctx context.Context, request ApproveTransitionRequest) (*ApproveTransitionRequestResponse, error) {
	var approveTransitionRequestResponse ApproveTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/approve"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &approveTransitionRequestResponse)
	return &approveTransitionRequestResponse, err
}

func (a *modelRegistryImpl) CreateComment(ctx context.Context, request CreateComment) (*CreateCommentResponse, error) {
	var createCommentResponse CreateCommentResponse
	path := "/api/2.0/mlflow/comments/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createCommentResponse)
	return &createCommentResponse, err
}

func (a *modelRegistryImpl) CreateModel(ctx context.Context, request CreateModelRequest) (*CreateModelResponse, error) {
	var createModelResponse CreateModelResponse
	path := "/api/2.0/mlflow/registered-models/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createModelResponse)
	return &createModelResponse, err
}

func (a *modelRegistryImpl) CreateModelVersion(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	var createModelVersionResponse CreateModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createModelVersionResponse)
	return &createModelVersionResponse, err
}

func (a *modelRegistryImpl) CreateTransitionRequest(ctx context.Context, request CreateTransitionRequest) (*CreateTransitionRequestResponse, error) {
	var createTransitionRequestResponse CreateTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createTransitionRequestResponse)
	return &createTransitionRequestResponse, err
}

func (a *modelRegistryImpl) CreateWebhook(ctx context.Context, request CreateRegistryWebhook) (*CreateWebhookResponse, error) {
	var createWebhookResponse CreateWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createWebhookResponse)
	return &createWebhookResponse, err
}

func (a *modelRegistryImpl) DeleteComment(ctx context.Context, request DeleteCommentRequest) error {
	path := "/api/2.0/mlflow/comments/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) DeleteModel(ctx context.Context, request DeleteModelRequest) error {
	path := "/api/2.0/mlflow/registered-models/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) DeleteModelTag(ctx context.Context, request DeleteModelTagRequest) error {
	path := "/api/2.0/mlflow/registered-models/delete-tag"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) DeleteModelVersion(ctx context.Context, request DeleteModelVersionRequest) error {
	path := "/api/2.0/mlflow/model-versions/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) DeleteModelVersionTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	path := "/api/2.0/mlflow/model-versions/delete-tag"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) DeleteTransitionRequest(ctx context.Context, request DeleteTransitionRequestRequest) error {
	path := "/api/2.0/mlflow/transition-requests/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) DeleteWebhook(ctx context.Context, request DeleteWebhookRequest) error {
	path := "/api/2.0/mlflow/registry-webhooks/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
	var getLatestVersionsResponse GetLatestVersionsResponse
	path := "/api/2.0/mlflow/registered-models/get-latest-versions"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &getLatestVersionsResponse)
	return &getLatestVersionsResponse, err
}

func (a *modelRegistryImpl) GetModel(ctx context.Context, request GetModelRequest) (*GetModelResponse, error) {
	var getModelResponse GetModelResponse
	path := "/api/2.0/mlflow/databricks/registered-models/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getModelResponse)
	return &getModelResponse, err
}

func (a *modelRegistryImpl) GetModelVersion(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	var getModelVersionResponse GetModelVersionResponse
	path := "/api/2.0/mlflow/model-versions/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getModelVersionResponse)
	return &getModelVersionResponse, err
}

func (a *modelRegistryImpl) GetModelVersionDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	var getModelVersionDownloadUriResponse GetModelVersionDownloadUriResponse
	path := "/api/2.0/mlflow/model-versions/get-download-uri"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getModelVersionDownloadUriResponse)
	return &getModelVersionDownloadUriResponse, err
}

func (a *modelRegistryImpl) GetPermissionLevels(ctx context.Context, request GetRegisteredModelPermissionLevelsRequest) (*GetRegisteredModelPermissionLevelsResponse, error) {
	var getRegisteredModelPermissionLevelsResponse GetRegisteredModelPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v/permissionLevels", request.RegisteredModelId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getRegisteredModelPermissionLevelsResponse)
	return &getRegisteredModelPermissionLevelsResponse, err
}

func (a *modelRegistryImpl) GetPermissions(ctx context.Context, request GetRegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

func (a *modelRegistryImpl) ListModels(ctx context.Context, request ListModelsRequest) (*ListModelsResponse, error) {
	var listModelsResponse ListModelsResponse
	path := "/api/2.0/mlflow/registered-models/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listModelsResponse)
	return &listModelsResponse, err
}

func (a *modelRegistryImpl) ListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) (*ListTransitionRequestsResponse, error) {
	var listTransitionRequestsResponse ListTransitionRequestsResponse
	path := "/api/2.0/mlflow/transition-requests/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listTransitionRequestsResponse)
	return &listTransitionRequestsResponse, err
}

func (a *modelRegistryImpl) ListWebhooks(ctx context.Context, request ListWebhooksRequest) (*ListRegistryWebhooks, error) {
	var listRegistryWebhooks ListRegistryWebhooks
	path := "/api/2.0/mlflow/registry-webhooks/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listRegistryWebhooks)
	return &listRegistryWebhooks, err
}

func (a *modelRegistryImpl) RejectTransitionRequest(ctx context.Context, request RejectTransitionRequest) (*RejectTransitionRequestResponse, error) {
	var rejectTransitionRequestResponse RejectTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/reject"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &rejectTransitionRequestResponse)
	return &rejectTransitionRequestResponse, err
}

func (a *modelRegistryImpl) RenameModel(ctx context.Context, request RenameModelRequest) (*RenameModelResponse, error) {
	var renameModelResponse RenameModelResponse
	path := "/api/2.0/mlflow/registered-models/rename"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &renameModelResponse)
	return &renameModelResponse, err
}

func (a *modelRegistryImpl) SearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
	var searchModelVersionsResponse SearchModelVersionsResponse
	path := "/api/2.0/mlflow/model-versions/search"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &searchModelVersionsResponse)
	return &searchModelVersionsResponse, err
}

func (a *modelRegistryImpl) SearchModels(ctx context.Context, request SearchModelsRequest) (*SearchModelsResponse, error) {
	var searchModelsResponse SearchModelsResponse
	path := "/api/2.0/mlflow/registered-models/search"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &searchModelsResponse)
	return &searchModelsResponse, err
}

func (a *modelRegistryImpl) SetModelTag(ctx context.Context, request SetModelTagRequest) error {
	path := "/api/2.0/mlflow/registered-models/set-tag"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) SetModelVersionTag(ctx context.Context, request SetModelVersionTagRequest) error {
	path := "/api/2.0/mlflow/model-versions/set-tag"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) SetPermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

func (a *modelRegistryImpl) TestRegistryWebhook(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	var testRegistryWebhookResponse TestRegistryWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/test"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &testRegistryWebhookResponse)
	return &testRegistryWebhookResponse, err
}

func (a *modelRegistryImpl) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {
	var transitionStageResponse TransitionStageResponse
	path := "/api/2.0/mlflow/databricks/model-versions/transition-stage"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &transitionStageResponse)
	return &transitionStageResponse, err
}

func (a *modelRegistryImpl) UpdateComment(ctx context.Context, request UpdateComment) (*UpdateCommentResponse, error) {
	var updateCommentResponse UpdateCommentResponse
	path := "/api/2.0/mlflow/comments/update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateCommentResponse)
	return &updateCommentResponse, err
}

func (a *modelRegistryImpl) UpdateModel(ctx context.Context, request UpdateModelRequest) error {
	path := "/api/2.0/mlflow/registered-models/update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) UpdateModelVersion(ctx context.Context, request UpdateModelVersionRequest) error {
	path := "/api/2.0/mlflow/model-versions/update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, nil)
	return err
}

func (a *modelRegistryImpl) UpdatePermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

func (a *modelRegistryImpl) UpdateWebhook(ctx context.Context, request UpdateRegistryWebhook) error {
	path := "/api/2.0/mlflow/registry-webhooks/update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, nil)
	return err
}
