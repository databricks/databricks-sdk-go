// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Experiments API methods
type experimentsImpl struct {
	client *client.DatabricksClient
}

func (a *experimentsImpl) CreateExperiment(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {

	requestPb, pbErr := createExperimentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createExperimentResponsePb createExperimentResponsePb
	path := "/api/2.0/mlflow/experiments/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createExperimentResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createExperimentResponseFromPb(&createExperimentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) CreateLoggedModel(ctx context.Context, request CreateLoggedModelRequest) (*CreateLoggedModelResponse, error) {

	requestPb, pbErr := createLoggedModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createLoggedModelResponsePb createLoggedModelResponsePb
	path := "/api/2.0/mlflow/logged-models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createLoggedModelResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createLoggedModelResponseFromPb(&createLoggedModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) CreateRun(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {

	requestPb, pbErr := createRunToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createRunResponsePb createRunResponsePb
	path := "/api/2.0/mlflow/runs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createRunResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createRunResponseFromPb(&createRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) DeleteExperiment(ctx context.Context, request DeleteExperiment) error {

	requestPb, pbErr := deleteExperimentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteExperimentResponsePb deleteExperimentResponsePb
	path := "/api/2.0/mlflow/experiments/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteExperimentResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) DeleteLoggedModel(ctx context.Context, request DeleteLoggedModelRequest) error {

	requestPb, pbErr := deleteLoggedModelRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteLoggedModelResponsePb deleteLoggedModelResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", requestPb.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteLoggedModelResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) DeleteLoggedModelTag(ctx context.Context, request DeleteLoggedModelTagRequest) error {

	requestPb, pbErr := deleteLoggedModelTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteLoggedModelTagResponsePb deleteLoggedModelTagResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/tags/%v", requestPb.ModelId, requestPb.TagKey)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteLoggedModelTagResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {

	requestPb, pbErr := deleteRunToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteRunResponsePb deleteRunResponsePb
	path := "/api/2.0/mlflow/runs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteRunResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) DeleteRuns(ctx context.Context, request DeleteRuns) (*DeleteRunsResponse, error) {

	requestPb, pbErr := deleteRunsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteRunsResponsePb deleteRunsResponsePb
	path := "/api/2.0/mlflow/databricks/runs/delete-runs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteRunsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteRunsResponseFromPb(&deleteRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) DeleteTag(ctx context.Context, request DeleteTag) error {

	requestPb, pbErr := deleteTagToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteTagResponsePb deleteTagResponsePb
	path := "/api/2.0/mlflow/runs/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteTagResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) FinalizeLoggedModel(ctx context.Context, request FinalizeLoggedModelRequest) (*FinalizeLoggedModelResponse, error) {

	requestPb, pbErr := finalizeLoggedModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var finalizeLoggedModelResponsePb finalizeLoggedModelResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", requestPb.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&finalizeLoggedModelResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := finalizeLoggedModelResponseFromPb(&finalizeLoggedModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error) {

	requestPb, pbErr := getByNameRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getExperimentByNameResponsePb getExperimentByNameResponsePb
	path := "/api/2.0/mlflow/experiments/get-by-name"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getExperimentByNameResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getExperimentByNameResponseFromPb(&getExperimentByNameResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetCredentialsForTraceDataDownload(ctx context.Context, request GetCredentialsForTraceDataDownloadRequest) (*GetCredentialsForTraceDataDownloadResponse, error) {

	requestPb, pbErr := getCredentialsForTraceDataDownloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getCredentialsForTraceDataDownloadResponsePb getCredentialsForTraceDataDownloadResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/traces/%v/credentials-for-data-download", requestPb.RequestId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getCredentialsForTraceDataDownloadResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getCredentialsForTraceDataDownloadResponseFromPb(&getCredentialsForTraceDataDownloadResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetCredentialsForTraceDataUpload(ctx context.Context, request GetCredentialsForTraceDataUploadRequest) (*GetCredentialsForTraceDataUploadResponse, error) {

	requestPb, pbErr := getCredentialsForTraceDataUploadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getCredentialsForTraceDataUploadResponsePb getCredentialsForTraceDataUploadResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/traces/%v/credentials-for-data-upload", requestPb.RequestId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getCredentialsForTraceDataUploadResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getCredentialsForTraceDataUploadResponseFromPb(&getCredentialsForTraceDataUploadResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetExperiment(ctx context.Context, request GetExperimentRequest) (*GetExperimentResponse, error) {

	requestPb, pbErr := getExperimentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getExperimentResponsePb getExperimentResponsePb
	path := "/api/2.0/mlflow/experiments/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getExperimentResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getExperimentResponseFromPb(&getExperimentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Get metric history for a run.
//
// Gets a list of all values for the specified metric for a given run.
func (a *experimentsImpl) GetHistory(ctx context.Context, request GetHistoryRequest) listing.Iterator[Metric] {

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

// Get metric history for a run.
//
// Gets a list of all values for the specified metric for a given run.
func (a *experimentsImpl) GetHistoryAll(ctx context.Context, request GetHistoryRequest) ([]Metric, error) {
	iterator := a.GetHistory(ctx, request)
	return listing.ToSlice[Metric](ctx, iterator)
}

func (a *experimentsImpl) internalGetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {

	requestPb, pbErr := getHistoryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getMetricHistoryResponsePb getMetricHistoryResponsePb
	path := "/api/2.0/mlflow/metrics/get-history"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getMetricHistoryResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getMetricHistoryResponseFromPb(&getMetricHistoryResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetLoggedModel(ctx context.Context, request GetLoggedModelRequest) (*GetLoggedModelResponse, error) {

	requestPb, pbErr := getLoggedModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getLoggedModelResponsePb getLoggedModelResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", requestPb.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getLoggedModelResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getLoggedModelResponseFromPb(&getLoggedModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetPermissionLevels(ctx context.Context, request GetExperimentPermissionLevelsRequest) (*GetExperimentPermissionLevelsResponse, error) {

	requestPb, pbErr := getExperimentPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getExperimentPermissionLevelsResponsePb getExperimentPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v/permissionLevels", requestPb.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getExperimentPermissionLevelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getExperimentPermissionLevelsResponseFromPb(&getExperimentPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetPermissions(ctx context.Context, request GetExperimentPermissionsRequest) (*ExperimentPermissions, error) {

	requestPb, pbErr := getExperimentPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var experimentPermissionsPb experimentPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", requestPb.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&experimentPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := experimentPermissionsFromPb(&experimentPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {

	requestPb, pbErr := getRunRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRunResponsePb getRunResponsePb
	path := "/api/2.0/mlflow/runs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getRunResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getRunResponseFromPb(&getRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List artifacts.
//
// List artifacts for a run. Takes an optional `artifact_path` prefix which if
// specified, the response contains only artifacts with the specified prefix. A
// maximum of 1000 artifacts will be retrieved for UC Volumes. Please call
// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
// Volumes, which supports pagination. See [List directory contents | Files
// API](/api/workspace/files/listdirectorycontents).
func (a *experimentsImpl) ListArtifacts(ctx context.Context, request ListArtifactsRequest) listing.Iterator[FileInfo] {

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

// List artifacts.
//
// List artifacts for a run. Takes an optional `artifact_path` prefix which if
// specified, the response contains only artifacts with the specified prefix. A
// maximum of 1000 artifacts will be retrieved for UC Volumes. Please call
// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
// Volumes, which supports pagination. See [List directory contents | Files
// API](/api/workspace/files/listdirectorycontents).
func (a *experimentsImpl) ListArtifactsAll(ctx context.Context, request ListArtifactsRequest) ([]FileInfo, error) {
	iterator := a.ListArtifacts(ctx, request)
	return listing.ToSlice[FileInfo](ctx, iterator)
}

func (a *experimentsImpl) internalListArtifacts(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error) {

	requestPb, pbErr := listArtifactsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listArtifactsResponsePb listArtifactsResponsePb
	path := "/api/2.0/mlflow/artifacts/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listArtifactsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listArtifactsResponseFromPb(&listArtifactsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List experiments.
//
// Gets a list of all experiments.
func (a *experimentsImpl) ListExperiments(ctx context.Context, request ListExperimentsRequest) listing.Iterator[Experiment] {

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
func (a *experimentsImpl) ListExperimentsAll(ctx context.Context, request ListExperimentsRequest) ([]Experiment, error) {
	iterator := a.ListExperiments(ctx, request)
	return listing.ToSlice[Experiment](ctx, iterator)
}

func (a *experimentsImpl) internalListExperiments(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error) {

	requestPb, pbErr := listExperimentsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listExperimentsResponsePb listExperimentsResponsePb
	path := "/api/2.0/mlflow/experiments/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listExperimentsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listExperimentsResponseFromPb(&listExperimentsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) ListLoggedModelArtifacts(ctx context.Context, request ListLoggedModelArtifactsRequest) (*ListLoggedModelArtifactsResponse, error) {

	requestPb, pbErr := listLoggedModelArtifactsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listLoggedModelArtifactsResponsePb listLoggedModelArtifactsResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/artifacts/directories", requestPb.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listLoggedModelArtifactsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listLoggedModelArtifactsResponseFromPb(&listLoggedModelArtifactsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) LogBatch(ctx context.Context, request LogBatch) error {

	requestPb, pbErr := logBatchToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var logBatchResponsePb logBatchResponsePb
	path := "/api/2.0/mlflow/runs/log-batch"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&logBatchResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) LogInputs(ctx context.Context, request LogInputs) error {

	requestPb, pbErr := logInputsToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var logInputsResponsePb logInputsResponsePb
	path := "/api/2.0/mlflow/runs/log-inputs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&logInputsResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) LogLoggedModelParams(ctx context.Context, request LogLoggedModelParamsRequest) error {

	requestPb, pbErr := logLoggedModelParamsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var logLoggedModelParamsRequestResponsePb logLoggedModelParamsRequestResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/params", requestPb.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&logLoggedModelParamsRequestResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) LogMetric(ctx context.Context, request LogMetric) error {

	requestPb, pbErr := logMetricToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var logMetricResponsePb logMetricResponsePb
	path := "/api/2.0/mlflow/runs/log-metric"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&logMetricResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) LogModel(ctx context.Context, request LogModel) error {

	requestPb, pbErr := logModelToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var logModelResponsePb logModelResponsePb
	path := "/api/2.0/mlflow/runs/log-model"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&logModelResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) LogOutputs(ctx context.Context, request LogOutputsRequest) error {

	requestPb, pbErr := logOutputsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var logOutputsResponsePb logOutputsResponsePb
	path := "/api/2.0/mlflow/runs/outputs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&logOutputsResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) LogParam(ctx context.Context, request LogParam) error {

	requestPb, pbErr := logParamToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var logParamResponsePb logParamResponsePb
	path := "/api/2.0/mlflow/runs/log-parameter"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&logParamResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) RestoreExperiment(ctx context.Context, request RestoreExperiment) error {

	requestPb, pbErr := restoreExperimentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var restoreExperimentResponsePb restoreExperimentResponsePb
	path := "/api/2.0/mlflow/experiments/restore"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&restoreExperimentResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) RestoreRun(ctx context.Context, request RestoreRun) error {

	requestPb, pbErr := restoreRunToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var restoreRunResponsePb restoreRunResponsePb
	path := "/api/2.0/mlflow/runs/restore"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&restoreRunResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) RestoreRuns(ctx context.Context, request RestoreRuns) (*RestoreRunsResponse, error) {

	requestPb, pbErr := restoreRunsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var restoreRunsResponsePb restoreRunsResponsePb
	path := "/api/2.0/mlflow/databricks/runs/restore-runs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&restoreRunsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := restoreRunsResponseFromPb(&restoreRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Search experiments.
//
// Searches for experiments that satisfy specified search criteria.
func (a *experimentsImpl) SearchExperiments(ctx context.Context, request SearchExperiments) listing.Iterator[Experiment] {

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
func (a *experimentsImpl) SearchExperimentsAll(ctx context.Context, request SearchExperiments) ([]Experiment, error) {
	iterator := a.SearchExperiments(ctx, request)
	return listing.ToSlice[Experiment](ctx, iterator)
}

func (a *experimentsImpl) internalSearchExperiments(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error) {

	requestPb, pbErr := searchExperimentsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchExperimentsResponsePb searchExperimentsResponsePb
	path := "/api/2.0/mlflow/experiments/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&searchExperimentsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := searchExperimentsResponseFromPb(&searchExperimentsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) SearchLoggedModels(ctx context.Context, request SearchLoggedModelsRequest) (*SearchLoggedModelsResponse, error) {

	requestPb, pbErr := searchLoggedModelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchLoggedModelsResponsePb searchLoggedModelsResponsePb
	path := "/api/2.0/mlflow/logged-models/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&searchLoggedModelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := searchLoggedModelsResponseFromPb(&searchLoggedModelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Search for runs.
//
// Searches for runs that satisfy expressions.
//
// Search expressions can use `mlflowMetric` and `mlflowParam` keys.
func (a *experimentsImpl) SearchRuns(ctx context.Context, request SearchRuns) listing.Iterator[Run] {

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
// Search expressions can use `mlflowMetric` and `mlflowParam` keys.
func (a *experimentsImpl) SearchRunsAll(ctx context.Context, request SearchRuns) ([]Run, error) {
	iterator := a.SearchRuns(ctx, request)
	return listing.ToSlice[Run](ctx, iterator)
}

func (a *experimentsImpl) internalSearchRuns(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error) {

	requestPb, pbErr := searchRunsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchRunsResponsePb searchRunsResponsePb
	path := "/api/2.0/mlflow/runs/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&searchRunsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := searchRunsResponseFromPb(&searchRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {

	requestPb, pbErr := setExperimentTagToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var setExperimentTagResponsePb setExperimentTagResponsePb
	path := "/api/2.0/mlflow/experiments/set-experiment-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&setExperimentTagResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) SetLoggedModelTags(ctx context.Context, request SetLoggedModelTagsRequest) error {

	requestPb, pbErr := setLoggedModelTagsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var setLoggedModelTagsResponsePb setLoggedModelTagsResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/tags", requestPb.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&setLoggedModelTagsResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) SetPermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {

	requestPb, pbErr := experimentPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var experimentPermissionsPb experimentPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", requestPb.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&experimentPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := experimentPermissionsFromPb(&experimentPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) SetTag(ctx context.Context, request SetTag) error {

	requestPb, pbErr := setTagToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var setTagResponsePb setTagResponsePb
	path := "/api/2.0/mlflow/runs/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&setTagResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) UpdateExperiment(ctx context.Context, request UpdateExperiment) error {

	requestPb, pbErr := updateExperimentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateExperimentResponsePb updateExperimentResponsePb
	path := "/api/2.0/mlflow/experiments/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateExperimentResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *experimentsImpl) UpdatePermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {

	requestPb, pbErr := experimentPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var experimentPermissionsPb experimentPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", requestPb.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&experimentPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := experimentPermissionsFromPb(&experimentPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) UpdateRun(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {

	requestPb, pbErr := updateRunToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateRunResponsePb updateRunResponsePb
	path := "/api/2.0/mlflow/runs/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateRunResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := updateRunResponseFromPb(&updateRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just forecasting API methods
type forecastingImpl struct {
	client *client.DatabricksClient
}

func (a *forecastingImpl) CreateExperiment(ctx context.Context, request CreateForecastingExperimentRequest) (*CreateForecastingExperimentResponse, error) {

	requestPb, pbErr := createForecastingExperimentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createForecastingExperimentResponsePb createForecastingExperimentResponsePb
	path := "/api/2.0/automl/create-forecasting-experiment"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createForecastingExperimentResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createForecastingExperimentResponseFromPb(&createForecastingExperimentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *forecastingImpl) GetExperiment(ctx context.Context, request GetForecastingExperimentRequest) (*ForecastingExperiment, error) {

	requestPb, pbErr := getForecastingExperimentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var forecastingExperimentPb forecastingExperimentPb
	path := fmt.Sprintf("/api/2.0/automl/get-forecasting-experiment/%v", requestPb.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&forecastingExperimentPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := forecastingExperimentFromPb(&forecastingExperimentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ModelRegistry API methods
type modelRegistryImpl struct {
	client *client.DatabricksClient
}

func (a *modelRegistryImpl) ApproveTransitionRequest(ctx context.Context, request ApproveTransitionRequest) (*ApproveTransitionRequestResponse, error) {

	requestPb, pbErr := approveTransitionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var approveTransitionRequestResponsePb approveTransitionRequestResponsePb
	path := "/api/2.0/mlflow/transition-requests/approve"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&approveTransitionRequestResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := approveTransitionRequestResponseFromPb(&approveTransitionRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateComment(ctx context.Context, request CreateComment) (*CreateCommentResponse, error) {

	requestPb, pbErr := createCommentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createCommentResponsePb createCommentResponsePb
	path := "/api/2.0/mlflow/comments/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createCommentResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createCommentResponseFromPb(&createCommentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateModel(ctx context.Context, request CreateModelRequest) (*CreateModelResponse, error) {

	requestPb, pbErr := createModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createModelResponsePb createModelResponsePb
	path := "/api/2.0/mlflow/registered-models/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createModelResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createModelResponseFromPb(&createModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateModelVersion(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {

	requestPb, pbErr := createModelVersionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createModelVersionResponsePb createModelVersionResponsePb
	path := "/api/2.0/mlflow/model-versions/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createModelVersionResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createModelVersionResponseFromPb(&createModelVersionResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateTransitionRequest(ctx context.Context, request CreateTransitionRequest) (*CreateTransitionRequestResponse, error) {

	requestPb, pbErr := createTransitionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createTransitionRequestResponsePb createTransitionRequestResponsePb
	path := "/api/2.0/mlflow/transition-requests/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createTransitionRequestResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createTransitionRequestResponseFromPb(&createTransitionRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateWebhook(ctx context.Context, request CreateRegistryWebhook) (*CreateWebhookResponse, error) {

	requestPb, pbErr := createRegistryWebhookToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createWebhookResponsePb createWebhookResponsePb
	path := "/api/2.0/mlflow/registry-webhooks/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createWebhookResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createWebhookResponseFromPb(&createWebhookResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) DeleteComment(ctx context.Context, request DeleteCommentRequest) error {

	requestPb, pbErr := deleteCommentRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteCommentResponsePb deleteCommentResponsePb
	path := "/api/2.0/mlflow/comments/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteCommentResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) DeleteModel(ctx context.Context, request DeleteModelRequest) error {

	requestPb, pbErr := deleteModelRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteModelResponsePb deleteModelResponsePb
	path := "/api/2.0/mlflow/registered-models/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteModelResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) DeleteModelTag(ctx context.Context, request DeleteModelTagRequest) error {

	requestPb, pbErr := deleteModelTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteModelTagResponsePb deleteModelTagResponsePb
	path := "/api/2.0/mlflow/registered-models/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteModelTagResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) DeleteModelVersion(ctx context.Context, request DeleteModelVersionRequest) error {

	requestPb, pbErr := deleteModelVersionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteModelVersionResponsePb deleteModelVersionResponsePb
	path := "/api/2.0/mlflow/model-versions/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteModelVersionResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) DeleteModelVersionTag(ctx context.Context, request DeleteModelVersionTagRequest) error {

	requestPb, pbErr := deleteModelVersionTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteModelVersionTagResponsePb deleteModelVersionTagResponsePb
	path := "/api/2.0/mlflow/model-versions/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteModelVersionTagResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) DeleteTransitionRequest(ctx context.Context, request DeleteTransitionRequestRequest) error {

	requestPb, pbErr := deleteTransitionRequestRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteTransitionRequestResponsePb deleteTransitionRequestResponsePb
	path := "/api/2.0/mlflow/transition-requests/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteTransitionRequestResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) DeleteWebhook(ctx context.Context, request DeleteWebhookRequest) error {

	requestPb, pbErr := deleteWebhookRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteWebhookResponsePb deleteWebhookResponsePb
	path := "/api/2.0/mlflow/registry-webhooks/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteWebhookResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

// Get the latest version.
//
// Gets the latest version of a registered model.
func (a *modelRegistryImpl) GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) listing.Iterator[ModelVersion] {

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
func (a *modelRegistryImpl) GetLatestVersionsAll(ctx context.Context, request GetLatestVersionsRequest) ([]ModelVersion, error) {
	iterator := a.GetLatestVersions(ctx, request)
	return listing.ToSlice[ModelVersion](ctx, iterator)
}

func (a *modelRegistryImpl) internalGetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {

	requestPb, pbErr := getLatestVersionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getLatestVersionsResponsePb getLatestVersionsResponsePb
	path := "/api/2.0/mlflow/registered-models/get-latest-versions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getLatestVersionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getLatestVersionsResponseFromPb(&getLatestVersionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetModel(ctx context.Context, request GetModelRequest) (*GetModelResponse, error) {

	requestPb, pbErr := getModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getModelResponsePb getModelResponsePb
	path := "/api/2.0/mlflow/databricks/registered-models/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getModelResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getModelResponseFromPb(&getModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetModelVersion(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {

	requestPb, pbErr := getModelVersionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getModelVersionResponsePb getModelVersionResponsePb
	path := "/api/2.0/mlflow/model-versions/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getModelVersionResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getModelVersionResponseFromPb(&getModelVersionResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetModelVersionDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {

	requestPb, pbErr := getModelVersionDownloadUriRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getModelVersionDownloadUriResponsePb getModelVersionDownloadUriResponsePb
	path := "/api/2.0/mlflow/model-versions/get-download-uri"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getModelVersionDownloadUriResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getModelVersionDownloadUriResponseFromPb(&getModelVersionDownloadUriResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetPermissionLevels(ctx context.Context, request GetRegisteredModelPermissionLevelsRequest) (*GetRegisteredModelPermissionLevelsResponse, error) {

	requestPb, pbErr := getRegisteredModelPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRegisteredModelPermissionLevelsResponsePb getRegisteredModelPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v/permissionLevels", requestPb.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getRegisteredModelPermissionLevelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getRegisteredModelPermissionLevelsResponseFromPb(&getRegisteredModelPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetPermissions(ctx context.Context, request GetRegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {

	requestPb, pbErr := getRegisteredModelPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelPermissionsPb registeredModelPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", requestPb.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&registeredModelPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := registeredModelPermissionsFromPb(&registeredModelPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List models.
//
// Lists all available registered models, up to the limit specified in
// __max_results__.
func (a *modelRegistryImpl) ListModels(ctx context.Context, request ListModelsRequest) listing.Iterator[Model] {

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
func (a *modelRegistryImpl) ListModelsAll(ctx context.Context, request ListModelsRequest) ([]Model, error) {
	iterator := a.ListModels(ctx, request)
	return listing.ToSliceN[Model, int](ctx, iterator, request.MaxResults)

}

func (a *modelRegistryImpl) internalListModels(ctx context.Context, request ListModelsRequest) (*ListModelsResponse, error) {

	requestPb, pbErr := listModelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listModelsResponsePb listModelsResponsePb
	path := "/api/2.0/mlflow/registered-models/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listModelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listModelsResponseFromPb(&listModelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List transition requests.
//
// Gets a list of all open stage transition requests for the model version.
func (a *modelRegistryImpl) ListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) listing.Iterator[Activity] {

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
func (a *modelRegistryImpl) ListTransitionRequestsAll(ctx context.Context, request ListTransitionRequestsRequest) ([]Activity, error) {
	iterator := a.ListTransitionRequests(ctx, request)
	return listing.ToSlice[Activity](ctx, iterator)
}

func (a *modelRegistryImpl) internalListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) (*ListTransitionRequestsResponse, error) {

	requestPb, pbErr := listTransitionRequestsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listTransitionRequestsResponsePb listTransitionRequestsResponsePb
	path := "/api/2.0/mlflow/transition-requests/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listTransitionRequestsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listTransitionRequestsResponseFromPb(&listTransitionRequestsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List registry webhooks.
//
// **NOTE:** This endpoint is in Public Preview.
//
// Lists all registry webhooks.
func (a *modelRegistryImpl) ListWebhooks(ctx context.Context, request ListWebhooksRequest) listing.Iterator[RegistryWebhook] {

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
func (a *modelRegistryImpl) ListWebhooksAll(ctx context.Context, request ListWebhooksRequest) ([]RegistryWebhook, error) {
	iterator := a.ListWebhooks(ctx, request)
	return listing.ToSlice[RegistryWebhook](ctx, iterator)
}

func (a *modelRegistryImpl) internalListWebhooks(ctx context.Context, request ListWebhooksRequest) (*ListRegistryWebhooks, error) {

	requestPb, pbErr := listWebhooksRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listRegistryWebhooksPb listRegistryWebhooksPb
	path := "/api/2.0/mlflow/registry-webhooks/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listRegistryWebhooksPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listRegistryWebhooksFromPb(&listRegistryWebhooksPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) RejectTransitionRequest(ctx context.Context, request RejectTransitionRequest) (*RejectTransitionRequestResponse, error) {

	requestPb, pbErr := rejectTransitionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var rejectTransitionRequestResponsePb rejectTransitionRequestResponsePb
	path := "/api/2.0/mlflow/transition-requests/reject"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&rejectTransitionRequestResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := rejectTransitionRequestResponseFromPb(&rejectTransitionRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) RenameModel(ctx context.Context, request RenameModelRequest) (*RenameModelResponse, error) {

	requestPb, pbErr := renameModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var renameModelResponsePb renameModelResponsePb
	path := "/api/2.0/mlflow/registered-models/rename"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&renameModelResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := renameModelResponseFromPb(&renameModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Searches model versions.
//
// Searches for specific model versions based on the supplied __filter__.
func (a *modelRegistryImpl) SearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) listing.Iterator[ModelVersion] {

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
func (a *modelRegistryImpl) SearchModelVersionsAll(ctx context.Context, request SearchModelVersionsRequest) ([]ModelVersion, error) {
	iterator := a.SearchModelVersions(ctx, request)
	return listing.ToSliceN[ModelVersion, int](ctx, iterator, request.MaxResults)

}

func (a *modelRegistryImpl) internalSearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {

	requestPb, pbErr := searchModelVersionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchModelVersionsResponsePb searchModelVersionsResponsePb
	path := "/api/2.0/mlflow/model-versions/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&searchModelVersionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := searchModelVersionsResponseFromPb(&searchModelVersionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Search models.
//
// Search for registered models based on the specified __filter__.
func (a *modelRegistryImpl) SearchModels(ctx context.Context, request SearchModelsRequest) listing.Iterator[Model] {

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
func (a *modelRegistryImpl) SearchModelsAll(ctx context.Context, request SearchModelsRequest) ([]Model, error) {
	iterator := a.SearchModels(ctx, request)
	return listing.ToSliceN[Model, int](ctx, iterator, request.MaxResults)

}

func (a *modelRegistryImpl) internalSearchModels(ctx context.Context, request SearchModelsRequest) (*SearchModelsResponse, error) {

	requestPb, pbErr := searchModelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchModelsResponsePb searchModelsResponsePb
	path := "/api/2.0/mlflow/registered-models/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&searchModelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := searchModelsResponseFromPb(&searchModelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) SetModelTag(ctx context.Context, request SetModelTagRequest) error {

	requestPb, pbErr := setModelTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var setModelTagResponsePb setModelTagResponsePb
	path := "/api/2.0/mlflow/registered-models/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&setModelTagResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) SetModelVersionTag(ctx context.Context, request SetModelVersionTagRequest) error {

	requestPb, pbErr := setModelVersionTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var setModelVersionTagResponsePb setModelVersionTagResponsePb
	path := "/api/2.0/mlflow/model-versions/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&setModelVersionTagResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) SetPermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {

	requestPb, pbErr := registeredModelPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelPermissionsPb registeredModelPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", requestPb.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&registeredModelPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := registeredModelPermissionsFromPb(&registeredModelPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) TestRegistryWebhook(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {

	requestPb, pbErr := testRegistryWebhookRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var testRegistryWebhookResponsePb testRegistryWebhookResponsePb
	path := "/api/2.0/mlflow/registry-webhooks/test"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&testRegistryWebhookResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := testRegistryWebhookResponseFromPb(&testRegistryWebhookResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {

	requestPb, pbErr := transitionModelVersionStageDatabricksToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var transitionStageResponsePb transitionStageResponsePb
	path := "/api/2.0/mlflow/databricks/model-versions/transition-stage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&transitionStageResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := transitionStageResponseFromPb(&transitionStageResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) UpdateComment(ctx context.Context, request UpdateComment) (*UpdateCommentResponse, error) {

	requestPb, pbErr := updateCommentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateCommentResponsePb updateCommentResponsePb
	path := "/api/2.0/mlflow/comments/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateCommentResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := updateCommentResponseFromPb(&updateCommentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) UpdateModel(ctx context.Context, request UpdateModelRequest) error {

	requestPb, pbErr := updateModelRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateModelResponsePb updateModelResponsePb
	path := "/api/2.0/mlflow/registered-models/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateModelResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) UpdateModelVersion(ctx context.Context, request UpdateModelVersionRequest) error {

	requestPb, pbErr := updateModelVersionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateModelVersionResponsePb updateModelVersionResponsePb
	path := "/api/2.0/mlflow/model-versions/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateModelVersionResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *modelRegistryImpl) UpdatePermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {

	requestPb, pbErr := registeredModelPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelPermissionsPb registeredModelPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", requestPb.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&registeredModelPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := registeredModelPermissionsFromPb(&registeredModelPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) UpdateWebhook(ctx context.Context, request UpdateRegistryWebhook) error {

	requestPb, pbErr := updateRegistryWebhookToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateWebhookResponsePb updateWebhookResponsePb
	path := "/api/2.0/mlflow/registry-webhooks/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateWebhookResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}
