// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlpreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just ExperimentsPreview API methods
type experimentsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *experimentsPreviewImpl) CreateExperiment(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {
	var createExperimentResponse CreateExperimentResponse
	path := "/api/2.0preview/mlflow/experiments/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createExperimentResponse)
	return &createExperimentResponse, err
}

func (a *experimentsPreviewImpl) CreateRun(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {
	var createRunResponse CreateRunResponse
	path := "/api/2.0preview/mlflow/runs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createRunResponse)
	return &createRunResponse, err
}

func (a *experimentsPreviewImpl) DeleteExperiment(ctx context.Context, request DeleteExperiment) error {
	var deleteExperimentResponse DeleteExperimentResponse
	path := "/api/2.0preview/mlflow/experiments/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteExperimentResponse)
	return err
}

func (a *experimentsPreviewImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	var deleteRunResponse DeleteRunResponse
	path := "/api/2.0preview/mlflow/runs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteRunResponse)
	return err
}

func (a *experimentsPreviewImpl) DeleteRuns(ctx context.Context, request DeleteRuns) (*DeleteRunsResponse, error) {
	var deleteRunsResponse DeleteRunsResponse
	path := "/api/2.0preview/mlflow/databricks/runs/delete-runs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteRunsResponse)
	return &deleteRunsResponse, err
}

func (a *experimentsPreviewImpl) DeleteTag(ctx context.Context, request DeleteTag) error {
	var deleteTagResponse DeleteTagResponse
	path := "/api/2.0preview/mlflow/runs/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteTagResponse)
	return err
}

func (a *experimentsPreviewImpl) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentResponse, error) {
	var getExperimentResponse GetExperimentResponse
	path := "/api/2.0preview/mlflow/experiments/get-by-name"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExperimentResponse)
	return &getExperimentResponse, err
}

func (a *experimentsPreviewImpl) GetExperiment(ctx context.Context, request GetExperimentRequest) (*GetExperimentResponse, error) {
	var getExperimentResponse GetExperimentResponse
	path := "/api/2.0preview/mlflow/experiments/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExperimentResponse)
	return &getExperimentResponse, err
}

// Get history of a given metric within a run.
//
// Gets a list of all values for the specified metric for a given run.
func (a *experimentsPreviewImpl) GetHistory(ctx context.Context, request GetHistoryRequest) listing.Iterator[Metric] {

	getNextPage := func(ctx context.Context, req GetHistoryRequest) (*GetMetricHistoryResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGetHistory(ctx, req)
	}
	getItems := func(resp *GetMetricHistoryResponse) []Metric {
		return resp.Metrics
	}
	getNextReq := func(resp *GetMetricHistoryResponse) *GetHistoryRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Get history of a given metric within a run.
//
// Gets a list of all values for the specified metric for a given run.
func (a *experimentsPreviewImpl) GetHistoryAll(ctx context.Context, request GetHistoryRequest) ([]Metric, error) {
	iterator := a.GetHistory(ctx, request)
	return listing.ToSliceN[Metric, int](ctx, iterator, request.MaxResults)

}
func (a *experimentsPreviewImpl) internalGetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {
	var getMetricHistoryResponse GetMetricHistoryResponse
	path := "/api/2.0preview/mlflow/metrics/get-history"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getMetricHistoryResponse)
	return &getMetricHistoryResponse, err
}

func (a *experimentsPreviewImpl) GetPermissionLevels(ctx context.Context, request GetExperimentPermissionLevelsRequest) (*GetExperimentPermissionLevelsResponse, error) {
	var getExperimentPermissionLevelsResponse GetExperimentPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0preview/permissions/experiments/%v/permissionLevels", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExperimentPermissionLevelsResponse)
	return &getExperimentPermissionLevelsResponse, err
}

func (a *experimentsPreviewImpl) GetPermissions(ctx context.Context, request GetExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsPreviewImpl) GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	var getRunResponse GetRunResponse
	path := "/api/2.0preview/mlflow/runs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRunResponse)
	return &getRunResponse, err
}

// Get all artifacts.
//
// List artifacts for a run. Takes an optional `artifact_path` prefix. If it is
// specified, the response contains only artifacts with the specified prefix.
// This API does not support pagination when listing artifacts in UC Volumes. A
// maximum of 1000 artifacts will be retrieved for UC Volumes. Please call
// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
// Volumes, which supports pagination. See [List directory contents | Files
// API](/api/workspace/files/listdirectorycontents).
func (a *experimentsPreviewImpl) ListArtifacts(ctx context.Context, request ListArtifactsRequest) listing.Iterator[FileInfo] {

	getNextPage := func(ctx context.Context, req ListArtifactsRequest) (*ListArtifactsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListArtifacts(ctx, req)
	}
	getItems := func(resp *ListArtifactsResponse) []FileInfo {
		return resp.Files
	}
	getNextReq := func(resp *ListArtifactsResponse) *ListArtifactsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Get all artifacts.
//
// List artifacts for a run. Takes an optional `artifact_path` prefix. If it is
// specified, the response contains only artifacts with the specified prefix.
// This API does not support pagination when listing artifacts in UC Volumes. A
// maximum of 1000 artifacts will be retrieved for UC Volumes. Please call
// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
// Volumes, which supports pagination. See [List directory contents | Files
// API](/api/workspace/files/listdirectorycontents).
func (a *experimentsPreviewImpl) ListArtifactsAll(ctx context.Context, request ListArtifactsRequest) ([]FileInfo, error) {
	iterator := a.ListArtifacts(ctx, request)
	return listing.ToSlice[FileInfo](ctx, iterator)
}
func (a *experimentsPreviewImpl) internalListArtifacts(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error) {
	var listArtifactsResponse ListArtifactsResponse
	path := "/api/2.0preview/mlflow/artifacts/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listArtifactsResponse)
	return &listArtifactsResponse, err
}

// List experiments.
//
// Gets a list of all experiments.
func (a *experimentsPreviewImpl) ListExperiments(ctx context.Context, request ListExperimentsRequest) listing.Iterator[Experiment] {

	getNextPage := func(ctx context.Context, req ListExperimentsRequest) (*ListExperimentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExperiments(ctx, req)
	}
	getItems := func(resp *ListExperimentsResponse) []Experiment {
		return resp.Experiments
	}
	getNextReq := func(resp *ListExperimentsResponse) *ListExperimentsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List experiments.
//
// Gets a list of all experiments.
func (a *experimentsPreviewImpl) ListExperimentsAll(ctx context.Context, request ListExperimentsRequest) ([]Experiment, error) {
	iterator := a.ListExperiments(ctx, request)
	return listing.ToSliceN[Experiment, int](ctx, iterator, request.MaxResults)

}
func (a *experimentsPreviewImpl) internalListExperiments(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error) {
	var listExperimentsResponse ListExperimentsResponse
	path := "/api/2.0preview/mlflow/experiments/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExperimentsResponse)
	return &listExperimentsResponse, err
}

func (a *experimentsPreviewImpl) LogBatch(ctx context.Context, request LogBatch) error {
	var logBatchResponse LogBatchResponse
	path := "/api/2.0preview/mlflow/runs/log-batch"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &logBatchResponse)
	return err
}

func (a *experimentsPreviewImpl) LogInputs(ctx context.Context, request LogInputs) error {
	var logInputsResponse LogInputsResponse
	path := "/api/2.0preview/mlflow/runs/log-inputs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &logInputsResponse)
	return err
}

func (a *experimentsPreviewImpl) LogMetric(ctx context.Context, request LogMetric) error {
	var logMetricResponse LogMetricResponse
	path := "/api/2.0preview/mlflow/runs/log-metric"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &logMetricResponse)
	return err
}

func (a *experimentsPreviewImpl) LogModel(ctx context.Context, request LogModel) error {
	var logModelResponse LogModelResponse
	path := "/api/2.0preview/mlflow/runs/log-model"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &logModelResponse)
	return err
}

func (a *experimentsPreviewImpl) LogParam(ctx context.Context, request LogParam) error {
	var logParamResponse LogParamResponse
	path := "/api/2.0preview/mlflow/runs/log-parameter"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &logParamResponse)
	return err
}

func (a *experimentsPreviewImpl) RestoreExperiment(ctx context.Context, request RestoreExperiment) error {
	var restoreExperimentResponse RestoreExperimentResponse
	path := "/api/2.0preview/mlflow/experiments/restore"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &restoreExperimentResponse)
	return err
}

func (a *experimentsPreviewImpl) RestoreRun(ctx context.Context, request RestoreRun) error {
	var restoreRunResponse RestoreRunResponse
	path := "/api/2.0preview/mlflow/runs/restore"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &restoreRunResponse)
	return err
}

func (a *experimentsPreviewImpl) RestoreRuns(ctx context.Context, request RestoreRuns) (*RestoreRunsResponse, error) {
	var restoreRunsResponse RestoreRunsResponse
	path := "/api/2.0preview/mlflow/databricks/runs/restore-runs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &restoreRunsResponse)
	return &restoreRunsResponse, err
}

// Search experiments.
//
// Searches for experiments that satisfy specified search criteria.
func (a *experimentsPreviewImpl) SearchExperiments(ctx context.Context, request SearchExperiments) listing.Iterator[Experiment] {

	getNextPage := func(ctx context.Context, req SearchExperiments) (*SearchExperimentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchExperiments(ctx, req)
	}
	getItems := func(resp *SearchExperimentsResponse) []Experiment {
		return resp.Experiments
	}
	getNextReq := func(resp *SearchExperimentsResponse) *SearchExperiments {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Search experiments.
//
// Searches for experiments that satisfy specified search criteria.
func (a *experimentsPreviewImpl) SearchExperimentsAll(ctx context.Context, request SearchExperiments) ([]Experiment, error) {
	iterator := a.SearchExperiments(ctx, request)
	return listing.ToSlice[Experiment](ctx, iterator)
}
func (a *experimentsPreviewImpl) internalSearchExperiments(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error) {
	var searchExperimentsResponse SearchExperimentsResponse
	path := "/api/2.0preview/mlflow/experiments/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &searchExperimentsResponse)
	return &searchExperimentsResponse, err
}

// Search for runs.
//
// Searches for runs that satisfy expressions.
//
// Search expressions can use `mlflowMetric` and `mlflowParam` keys.",
func (a *experimentsPreviewImpl) SearchRuns(ctx context.Context, request SearchRuns) listing.Iterator[Run] {

	getNextPage := func(ctx context.Context, req SearchRuns) (*SearchRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchRuns(ctx, req)
	}
	getItems := func(resp *SearchRunsResponse) []Run {
		return resp.Runs
	}
	getNextReq := func(resp *SearchRunsResponse) *SearchRuns {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Search for runs.
//
// Searches for runs that satisfy expressions.
//
// Search expressions can use `mlflowMetric` and `mlflowParam` keys.",
func (a *experimentsPreviewImpl) SearchRunsAll(ctx context.Context, request SearchRuns) ([]Run, error) {
	iterator := a.SearchRuns(ctx, request)
	return listing.ToSlice[Run](ctx, iterator)
}
func (a *experimentsPreviewImpl) internalSearchRuns(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error) {
	var searchRunsResponse SearchRunsResponse
	path := "/api/2.0preview/mlflow/runs/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &searchRunsResponse)
	return &searchRunsResponse, err
}

func (a *experimentsPreviewImpl) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {
	var setExperimentTagResponse SetExperimentTagResponse
	path := "/api/2.0preview/mlflow/experiments/set-experiment-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &setExperimentTagResponse)
	return err
}

func (a *experimentsPreviewImpl) SetPermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsPreviewImpl) SetTag(ctx context.Context, request SetTag) error {
	var setTagResponse SetTagResponse
	path := "/api/2.0preview/mlflow/runs/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &setTagResponse)
	return err
}

func (a *experimentsPreviewImpl) UpdateExperiment(ctx context.Context, request UpdateExperiment) error {
	var updateExperimentResponse UpdateExperimentResponse
	path := "/api/2.0preview/mlflow/experiments/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &updateExperimentResponse)
	return err
}

func (a *experimentsPreviewImpl) UpdatePermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	var experimentPermissions ExperimentPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &experimentPermissions)
	return &experimentPermissions, err
}

func (a *experimentsPreviewImpl) UpdateRun(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {
	var updateRunResponse UpdateRunResponse
	path := "/api/2.0preview/mlflow/runs/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &updateRunResponse)
	return &updateRunResponse, err
}

// unexported type that holds implementations of just ModelRegistryPreview API methods
type modelRegistryPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *modelRegistryPreviewImpl) ApproveTransitionRequest(ctx context.Context, request ApproveTransitionRequest) (*ApproveTransitionRequestResponse, error) {
	var approveTransitionRequestResponse ApproveTransitionRequestResponse
	path := "/api/2.0preview/mlflow/transition-requests/approve"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &approveTransitionRequestResponse)
	return &approveTransitionRequestResponse, err
}

func (a *modelRegistryPreviewImpl) CreateComment(ctx context.Context, request CreateComment) (*CreateCommentResponse, error) {
	var createCommentResponse CreateCommentResponse
	path := "/api/2.0preview/mlflow/comments/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createCommentResponse)
	return &createCommentResponse, err
}

func (a *modelRegistryPreviewImpl) CreateModel(ctx context.Context, request CreateModelRequest) (*CreateModelResponse, error) {
	var createModelResponse CreateModelResponse
	path := "/api/2.0preview/mlflow/registered-models/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createModelResponse)
	return &createModelResponse, err
}

func (a *modelRegistryPreviewImpl) CreateModelVersion(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	var createModelVersionResponse CreateModelVersionResponse
	path := "/api/2.0preview/mlflow/model-versions/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createModelVersionResponse)
	return &createModelVersionResponse, err
}

func (a *modelRegistryPreviewImpl) CreateTransitionRequest(ctx context.Context, request CreateTransitionRequest) (*CreateTransitionRequestResponse, error) {
	var createTransitionRequestResponse CreateTransitionRequestResponse
	path := "/api/2.0preview/mlflow/transition-requests/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createTransitionRequestResponse)
	return &createTransitionRequestResponse, err
}

func (a *modelRegistryPreviewImpl) CreateWebhook(ctx context.Context, request CreateRegistryWebhook) (*CreateWebhookResponse, error) {
	var createWebhookResponse CreateWebhookResponse
	path := "/api/2.0preview/mlflow/registry-webhooks/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createWebhookResponse)
	return &createWebhookResponse, err
}

func (a *modelRegistryPreviewImpl) DeleteComment(ctx context.Context, request DeleteCommentRequest) error {
	var deleteCommentResponse DeleteCommentResponse
	path := "/api/2.0preview/mlflow/comments/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteCommentResponse)
	return err
}

func (a *modelRegistryPreviewImpl) DeleteModel(ctx context.Context, request DeleteModelRequest) error {
	var deleteModelResponse DeleteModelResponse
	path := "/api/2.0preview/mlflow/registered-models/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteModelResponse)
	return err
}

func (a *modelRegistryPreviewImpl) DeleteModelTag(ctx context.Context, request DeleteModelTagRequest) error {
	var deleteModelTagResponse DeleteModelTagResponse
	path := "/api/2.0preview/mlflow/registered-models/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteModelTagResponse)
	return err
}

func (a *modelRegistryPreviewImpl) DeleteModelVersion(ctx context.Context, request DeleteModelVersionRequest) error {
	var deleteModelVersionResponse DeleteModelVersionResponse
	path := "/api/2.0preview/mlflow/model-versions/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteModelVersionResponse)
	return err
}

func (a *modelRegistryPreviewImpl) DeleteModelVersionTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	var deleteModelVersionTagResponse DeleteModelVersionTagResponse
	path := "/api/2.0preview/mlflow/model-versions/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteModelVersionTagResponse)
	return err
}

func (a *modelRegistryPreviewImpl) DeleteTransitionRequest(ctx context.Context, request DeleteTransitionRequestRequest) error {
	var deleteTransitionRequestResponse DeleteTransitionRequestResponse
	path := "/api/2.0preview/mlflow/transition-requests/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteTransitionRequestResponse)
	return err
}

func (a *modelRegistryPreviewImpl) DeleteWebhook(ctx context.Context, request DeleteWebhookRequest) error {
	var deleteWebhookResponse DeleteWebhookResponse
	path := "/api/2.0preview/mlflow/registry-webhooks/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteWebhookResponse)
	return err
}

// Get the latest version.
//
// Gets the latest version of a registered model.
func (a *modelRegistryPreviewImpl) GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) listing.Iterator[ModelVersion] {

	getNextPage := func(ctx context.Context, req GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGetLatestVersions(ctx, req)
	}
	getItems := func(resp *GetLatestVersionsResponse) []ModelVersion {
		return resp.ModelVersions
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get the latest version.
//
// Gets the latest version of a registered model.
func (a *modelRegistryPreviewImpl) GetLatestVersionsAll(ctx context.Context, request GetLatestVersionsRequest) ([]ModelVersion, error) {
	iterator := a.GetLatestVersions(ctx, request)
	return listing.ToSlice[ModelVersion](ctx, iterator)
}
func (a *modelRegistryPreviewImpl) internalGetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
	var getLatestVersionsResponse GetLatestVersionsResponse
	path := "/api/2.0preview/mlflow/registered-models/get-latest-versions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &getLatestVersionsResponse)
	return &getLatestVersionsResponse, err
}

func (a *modelRegistryPreviewImpl) GetModel(ctx context.Context, request GetModelRequest) (*GetModelResponse, error) {
	var getModelResponse GetModelResponse
	path := "/api/2.0preview/mlflow/databricks/registered-models/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getModelResponse)
	return &getModelResponse, err
}

func (a *modelRegistryPreviewImpl) GetModelVersion(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	var getModelVersionResponse GetModelVersionResponse
	path := "/api/2.0preview/mlflow/model-versions/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getModelVersionResponse)
	return &getModelVersionResponse, err
}

func (a *modelRegistryPreviewImpl) GetModelVersionDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	var getModelVersionDownloadUriResponse GetModelVersionDownloadUriResponse
	path := "/api/2.0preview/mlflow/model-versions/get-download-uri"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getModelVersionDownloadUriResponse)
	return &getModelVersionDownloadUriResponse, err
}

func (a *modelRegistryPreviewImpl) GetPermissionLevels(ctx context.Context, request GetRegisteredModelPermissionLevelsRequest) (*GetRegisteredModelPermissionLevelsResponse, error) {
	var getRegisteredModelPermissionLevelsResponse GetRegisteredModelPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0preview/permissions/registered-models/%v/permissionLevels", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRegisteredModelPermissionLevelsResponse)
	return &getRegisteredModelPermissionLevelsResponse, err
}

func (a *modelRegistryPreviewImpl) GetPermissions(ctx context.Context, request GetRegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

// List models.
//
// Lists all available registered models, up to the limit specified in
// __max_results__.
func (a *modelRegistryPreviewImpl) ListModels(ctx context.Context, request ListModelsRequest) listing.Iterator[Model] {

	getNextPage := func(ctx context.Context, req ListModelsRequest) (*ListModelsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListModels(ctx, req)
	}
	getItems := func(resp *ListModelsResponse) []Model {
		return resp.RegisteredModels
	}
	getNextReq := func(resp *ListModelsResponse) *ListModelsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List models.
//
// Lists all available registered models, up to the limit specified in
// __max_results__.
func (a *modelRegistryPreviewImpl) ListModelsAll(ctx context.Context, request ListModelsRequest) ([]Model, error) {
	iterator := a.ListModels(ctx, request)
	return listing.ToSliceN[Model, int](ctx, iterator, request.MaxResults)

}
func (a *modelRegistryPreviewImpl) internalListModels(ctx context.Context, request ListModelsRequest) (*ListModelsResponse, error) {
	var listModelsResponse ListModelsResponse
	path := "/api/2.0preview/mlflow/registered-models/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listModelsResponse)
	return &listModelsResponse, err
}

// List transition requests.
//
// Gets a list of all open stage transition requests for the model version.
func (a *modelRegistryPreviewImpl) ListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) listing.Iterator[Activity] {

	getNextPage := func(ctx context.Context, req ListTransitionRequestsRequest) (*ListTransitionRequestsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListTransitionRequests(ctx, req)
	}
	getItems := func(resp *ListTransitionRequestsResponse) []Activity {
		return resp.Requests
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List transition requests.
//
// Gets a list of all open stage transition requests for the model version.
func (a *modelRegistryPreviewImpl) ListTransitionRequestsAll(ctx context.Context, request ListTransitionRequestsRequest) ([]Activity, error) {
	iterator := a.ListTransitionRequests(ctx, request)
	return listing.ToSlice[Activity](ctx, iterator)
}
func (a *modelRegistryPreviewImpl) internalListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) (*ListTransitionRequestsResponse, error) {
	var listTransitionRequestsResponse ListTransitionRequestsResponse
	path := "/api/2.0preview/mlflow/transition-requests/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTransitionRequestsResponse)
	return &listTransitionRequestsResponse, err
}

// List registry webhooks.
//
// **NOTE:** This endpoint is in Public Preview.
//
// Lists all registry webhooks.
func (a *modelRegistryPreviewImpl) ListWebhooks(ctx context.Context, request ListWebhooksRequest) listing.Iterator[RegistryWebhook] {

	getNextPage := func(ctx context.Context, req ListWebhooksRequest) (*ListRegistryWebhooks, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListWebhooks(ctx, req)
	}
	getItems := func(resp *ListRegistryWebhooks) []RegistryWebhook {
		return resp.Webhooks
	}
	getNextReq := func(resp *ListRegistryWebhooks) *ListWebhooksRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List registry webhooks.
//
// **NOTE:** This endpoint is in Public Preview.
//
// Lists all registry webhooks.
func (a *modelRegistryPreviewImpl) ListWebhooksAll(ctx context.Context, request ListWebhooksRequest) ([]RegistryWebhook, error) {
	iterator := a.ListWebhooks(ctx, request)
	return listing.ToSlice[RegistryWebhook](ctx, iterator)
}
func (a *modelRegistryPreviewImpl) internalListWebhooks(ctx context.Context, request ListWebhooksRequest) (*ListRegistryWebhooks, error) {
	var listRegistryWebhooks ListRegistryWebhooks
	path := "/api/2.0preview/mlflow/registry-webhooks/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRegistryWebhooks)
	return &listRegistryWebhooks, err
}

func (a *modelRegistryPreviewImpl) RejectTransitionRequest(ctx context.Context, request RejectTransitionRequest) (*RejectTransitionRequestResponse, error) {
	var rejectTransitionRequestResponse RejectTransitionRequestResponse
	path := "/api/2.0preview/mlflow/transition-requests/reject"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &rejectTransitionRequestResponse)
	return &rejectTransitionRequestResponse, err
}

func (a *modelRegistryPreviewImpl) RenameModel(ctx context.Context, request RenameModelRequest) (*RenameModelResponse, error) {
	var renameModelResponse RenameModelResponse
	path := "/api/2.0preview/mlflow/registered-models/rename"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &renameModelResponse)
	return &renameModelResponse, err
}

// Searches model versions.
//
// Searches for specific model versions based on the supplied __filter__.
func (a *modelRegistryPreviewImpl) SearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) listing.Iterator[ModelVersion] {

	getNextPage := func(ctx context.Context, req SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchModelVersions(ctx, req)
	}
	getItems := func(resp *SearchModelVersionsResponse) []ModelVersion {
		return resp.ModelVersions
	}
	getNextReq := func(resp *SearchModelVersionsResponse) *SearchModelVersionsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Searches model versions.
//
// Searches for specific model versions based on the supplied __filter__.
func (a *modelRegistryPreviewImpl) SearchModelVersionsAll(ctx context.Context, request SearchModelVersionsRequest) ([]ModelVersion, error) {
	iterator := a.SearchModelVersions(ctx, request)
	return listing.ToSliceN[ModelVersion, int](ctx, iterator, request.MaxResults)

}
func (a *modelRegistryPreviewImpl) internalSearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
	var searchModelVersionsResponse SearchModelVersionsResponse
	path := "/api/2.0preview/mlflow/model-versions/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &searchModelVersionsResponse)
	return &searchModelVersionsResponse, err
}

// Search models.
//
// Search for registered models based on the specified __filter__.
func (a *modelRegistryPreviewImpl) SearchModels(ctx context.Context, request SearchModelsRequest) listing.Iterator[Model] {

	getNextPage := func(ctx context.Context, req SearchModelsRequest) (*SearchModelsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchModels(ctx, req)
	}
	getItems := func(resp *SearchModelsResponse) []Model {
		return resp.RegisteredModels
	}
	getNextReq := func(resp *SearchModelsResponse) *SearchModelsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Search models.
//
// Search for registered models based on the specified __filter__.
func (a *modelRegistryPreviewImpl) SearchModelsAll(ctx context.Context, request SearchModelsRequest) ([]Model, error) {
	iterator := a.SearchModels(ctx, request)
	return listing.ToSliceN[Model, int](ctx, iterator, request.MaxResults)

}
func (a *modelRegistryPreviewImpl) internalSearchModels(ctx context.Context, request SearchModelsRequest) (*SearchModelsResponse, error) {
	var searchModelsResponse SearchModelsResponse
	path := "/api/2.0preview/mlflow/registered-models/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &searchModelsResponse)
	return &searchModelsResponse, err
}

func (a *modelRegistryPreviewImpl) SetModelTag(ctx context.Context, request SetModelTagRequest) error {
	var setModelTagResponse SetModelTagResponse
	path := "/api/2.0preview/mlflow/registered-models/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &setModelTagResponse)
	return err
}

func (a *modelRegistryPreviewImpl) SetModelVersionTag(ctx context.Context, request SetModelVersionTagRequest) error {
	var setModelVersionTagResponse SetModelVersionTagResponse
	path := "/api/2.0preview/mlflow/model-versions/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &setModelVersionTagResponse)
	return err
}

func (a *modelRegistryPreviewImpl) SetPermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

func (a *modelRegistryPreviewImpl) TestRegistryWebhook(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	var testRegistryWebhookResponse TestRegistryWebhookResponse
	path := "/api/2.0preview/mlflow/registry-webhooks/test"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &testRegistryWebhookResponse)
	return &testRegistryWebhookResponse, err
}

func (a *modelRegistryPreviewImpl) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {
	var transitionStageResponse TransitionStageResponse
	path := "/api/2.0preview/mlflow/databricks/model-versions/transition-stage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &transitionStageResponse)
	return &transitionStageResponse, err
}

func (a *modelRegistryPreviewImpl) UpdateComment(ctx context.Context, request UpdateComment) (*UpdateCommentResponse, error) {
	var updateCommentResponse UpdateCommentResponse
	path := "/api/2.0preview/mlflow/comments/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateCommentResponse)
	return &updateCommentResponse, err
}

func (a *modelRegistryPreviewImpl) UpdateModel(ctx context.Context, request UpdateModelRequest) error {
	var updateModelResponse UpdateModelResponse
	path := "/api/2.0preview/mlflow/registered-models/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateModelResponse)
	return err
}

func (a *modelRegistryPreviewImpl) UpdateModelVersion(ctx context.Context, request UpdateModelVersionRequest) error {
	var updateModelVersionResponse UpdateModelVersionResponse
	path := "/api/2.0preview/mlflow/model-versions/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateModelVersionResponse)
	return err
}

func (a *modelRegistryPreviewImpl) UpdatePermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	var registeredModelPermissions RegisteredModelPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &registeredModelPermissions)
	return &registeredModelPermissions, err
}

func (a *modelRegistryPreviewImpl) UpdateWebhook(ctx context.Context, request UpdateRegistryWebhook) error {
	var updateWebhookResponse UpdateWebhookResponse
	path := "/api/2.0preview/mlflow/registry-webhooks/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateWebhookResponse)
	return err
}
