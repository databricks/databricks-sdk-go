// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/ml/mlpb"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Experiments API methods
type experimentsImpl struct {
	client *client.DatabricksClient
}

func (a *experimentsImpl) CreateExperiment(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error) {
	requestPb, pbErr := CreateExperimentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createExperimentResponsePb mlpb.CreateExperimentResponsePb
	path := "/api/2.0/mlflow/experiments/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createExperimentResponsePb,
	)
	resp, err := CreateExperimentResponseFromPb(&createExperimentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) CreateLoggedModel(ctx context.Context, request CreateLoggedModelRequest) (*CreateLoggedModelResponse, error) {
	requestPb, pbErr := CreateLoggedModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createLoggedModelResponsePb mlpb.CreateLoggedModelResponsePb
	path := "/api/2.0/mlflow/logged-models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createLoggedModelResponsePb,
	)
	resp, err := CreateLoggedModelResponseFromPb(&createLoggedModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) CreateRun(ctx context.Context, request CreateRun) (*CreateRunResponse, error) {
	requestPb, pbErr := CreateRunToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createRunResponsePb mlpb.CreateRunResponsePb
	path := "/api/2.0/mlflow/runs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createRunResponsePb,
	)
	resp, err := CreateRunResponseFromPb(&createRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) DeleteExperiment(ctx context.Context, request DeleteExperiment) error {
	requestPb, pbErr := DeleteExperimentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/experiments/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) DeleteLoggedModel(ctx context.Context, request DeleteLoggedModelRequest) error {
	requestPb, pbErr := DeleteLoggedModelRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", request.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) DeleteLoggedModelTag(ctx context.Context, request DeleteLoggedModelTagRequest) error {
	requestPb, pbErr := DeleteLoggedModelTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/tags/%v", request.ModelId, request.TagKey)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	requestPb, pbErr := DeleteRunToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) DeleteRuns(ctx context.Context, request DeleteRuns) (*DeleteRunsResponse, error) {
	requestPb, pbErr := DeleteRunsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteRunsResponsePb mlpb.DeleteRunsResponsePb
	path := "/api/2.0/mlflow/databricks/runs/delete-runs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteRunsResponsePb,
	)
	resp, err := DeleteRunsResponseFromPb(&deleteRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) DeleteTag(ctx context.Context, request DeleteTag) error {
	requestPb, pbErr := DeleteTagToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) FinalizeLoggedModel(ctx context.Context, request FinalizeLoggedModelRequest) (*FinalizeLoggedModelResponse, error) {
	requestPb, pbErr := FinalizeLoggedModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var finalizeLoggedModelResponsePb mlpb.FinalizeLoggedModelResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", request.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&finalizeLoggedModelResponsePb,
	)
	resp, err := FinalizeLoggedModelResponseFromPb(&finalizeLoggedModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error) {
	requestPb, pbErr := GetByNameRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getExperimentByNameResponsePb mlpb.GetExperimentByNameResponsePb
	path := "/api/2.0/mlflow/experiments/get-by-name"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getExperimentByNameResponsePb,
	)
	resp, err := GetExperimentByNameResponseFromPb(&getExperimentByNameResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetExperiment(ctx context.Context, request GetExperimentRequest) (*GetExperimentResponse, error) {
	requestPb, pbErr := GetExperimentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getExperimentResponsePb mlpb.GetExperimentResponsePb
	path := "/api/2.0/mlflow/experiments/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getExperimentResponsePb,
	)
	resp, err := GetExperimentResponseFromPb(&getExperimentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets a list of all values for the specified metric for a given run.
func (a *experimentsImpl) GetHistoryAll(ctx context.Context, request GetHistoryRequest) ([]Metric, error) {
	iterator := a.GetHistory(ctx, request)
	return listing.ToSlice[Metric](ctx, iterator)
}

func (a *experimentsImpl) internalGetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error) {
	requestPb, pbErr := GetHistoryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getMetricHistoryResponsePb mlpb.GetMetricHistoryResponsePb
	path := "/api/2.0/mlflow/metrics/get-history"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getMetricHistoryResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := GetMetricHistoryResponseFromPb(&getMetricHistoryResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetLoggedModel(ctx context.Context, request GetLoggedModelRequest) (*GetLoggedModelResponse, error) {
	requestPb, pbErr := GetLoggedModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getLoggedModelResponsePb mlpb.GetLoggedModelResponsePb
	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v", request.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getLoggedModelResponsePb,
	)
	resp, err := GetLoggedModelResponseFromPb(&getLoggedModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetPermissionLevels(ctx context.Context, request GetExperimentPermissionLevelsRequest) (*GetExperimentPermissionLevelsResponse, error) {
	requestPb, pbErr := GetExperimentPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getExperimentPermissionLevelsResponsePb mlpb.GetExperimentPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v/permissionLevels", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getExperimentPermissionLevelsResponsePb,
	)
	resp, err := GetExperimentPermissionLevelsResponseFromPb(&getExperimentPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetPermissions(ctx context.Context, request GetExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	requestPb, pbErr := GetExperimentPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var experimentPermissionsPb mlpb.ExperimentPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&experimentPermissionsPb,
	)
	resp, err := ExperimentPermissionsFromPb(&experimentPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	requestPb, pbErr := GetRunRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRunResponsePb mlpb.GetRunResponsePb
	path := "/api/2.0/mlflow/runs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getRunResponsePb,
	)
	resp, err := GetRunResponseFromPb(&getRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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
	requestPb, pbErr := ListArtifactsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listArtifactsResponsePb mlpb.ListArtifactsResponsePb
	path := "/api/2.0/mlflow/artifacts/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listArtifactsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListArtifactsResponseFromPb(&listArtifactsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets a list of all experiments.
func (a *experimentsImpl) ListExperimentsAll(ctx context.Context, request ListExperimentsRequest) ([]Experiment, error) {
	iterator := a.ListExperiments(ctx, request)
	return listing.ToSlice[Experiment](ctx, iterator)
}

func (a *experimentsImpl) internalListExperiments(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error) {
	requestPb, pbErr := ListExperimentsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listExperimentsResponsePb mlpb.ListExperimentsResponsePb
	path := "/api/2.0/mlflow/experiments/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listExperimentsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListExperimentsResponseFromPb(&listExperimentsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) LogBatch(ctx context.Context, request LogBatch) error {
	requestPb, pbErr := LogBatchToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/log-batch"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) LogInputs(ctx context.Context, request LogInputs) error {
	requestPb, pbErr := LogInputsToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/log-inputs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) LogLoggedModelParams(ctx context.Context, request LogLoggedModelParamsRequest) error {
	requestPb, pbErr := LogLoggedModelParamsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/params", request.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) LogMetric(ctx context.Context, request LogMetric) error {
	requestPb, pbErr := LogMetricToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/log-metric"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) LogModel(ctx context.Context, request LogModel) error {
	requestPb, pbErr := LogModelToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/log-model"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) LogOutputs(ctx context.Context, request LogOutputsRequest) error {
	requestPb, pbErr := LogOutputsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/outputs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) LogParam(ctx context.Context, request LogParam) error {
	requestPb, pbErr := LogParamToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/log-parameter"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) RestoreExperiment(ctx context.Context, request RestoreExperiment) error {
	requestPb, pbErr := RestoreExperimentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/experiments/restore"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) RestoreRun(ctx context.Context, request RestoreRun) error {
	requestPb, pbErr := RestoreRunToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/restore"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) RestoreRuns(ctx context.Context, request RestoreRuns) (*RestoreRunsResponse, error) {
	requestPb, pbErr := RestoreRunsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var restoreRunsResponsePb mlpb.RestoreRunsResponsePb
	path := "/api/2.0/mlflow/databricks/runs/restore-runs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&restoreRunsResponsePb,
	)
	resp, err := RestoreRunsResponseFromPb(&restoreRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Searches for experiments that satisfy specified search criteria.
func (a *experimentsImpl) SearchExperimentsAll(ctx context.Context, request SearchExperiments) ([]Experiment, error) {
	iterator := a.SearchExperiments(ctx, request)
	return listing.ToSlice[Experiment](ctx, iterator)
}

func (a *experimentsImpl) internalSearchExperiments(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error) {
	requestPb, pbErr := SearchExperimentsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchExperimentsResponsePb mlpb.SearchExperimentsResponsePb
	path := "/api/2.0/mlflow/experiments/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&searchExperimentsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := SearchExperimentsResponseFromPb(&searchExperimentsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) SearchLoggedModels(ctx context.Context, request SearchLoggedModelsRequest) (*SearchLoggedModelsResponse, error) {
	requestPb, pbErr := SearchLoggedModelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchLoggedModelsResponsePb mlpb.SearchLoggedModelsResponsePb
	path := "/api/2.0/mlflow/logged-models/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&searchLoggedModelsResponsePb,
	)
	resp, err := SearchLoggedModelsResponseFromPb(&searchLoggedModelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Searches for runs that satisfy expressions.
//
// Search expressions can use `mlflowMetric` and `mlflowParam` keys.
func (a *experimentsImpl) SearchRunsAll(ctx context.Context, request SearchRuns) ([]Run, error) {
	iterator := a.SearchRuns(ctx, request)
	return listing.ToSlice[Run](ctx, iterator)
}

func (a *experimentsImpl) internalSearchRuns(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error) {
	requestPb, pbErr := SearchRunsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchRunsResponsePb mlpb.SearchRunsResponsePb
	path := "/api/2.0/mlflow/runs/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&searchRunsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := SearchRunsResponseFromPb(&searchRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) SetExperimentTag(ctx context.Context, request SetExperimentTag) error {
	requestPb, pbErr := SetExperimentTagToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/experiments/set-experiment-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) SetLoggedModelTags(ctx context.Context, request SetLoggedModelTagsRequest) error {
	requestPb, pbErr := SetLoggedModelTagsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/mlflow/logged-models/%v/tags", request.ModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) SetPermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	requestPb, pbErr := ExperimentPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var experimentPermissionsPb mlpb.ExperimentPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&experimentPermissionsPb,
	)
	resp, err := ExperimentPermissionsFromPb(&experimentPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) SetTag(ctx context.Context, request SetTag) error {
	requestPb, pbErr := SetTagToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/runs/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) UpdateExperiment(ctx context.Context, request UpdateExperiment) error {
	requestPb, pbErr := UpdateExperimentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/experiments/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *experimentsImpl) UpdatePermissions(ctx context.Context, request ExperimentPermissionsRequest) (*ExperimentPermissions, error) {
	requestPb, pbErr := ExperimentPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var experimentPermissionsPb mlpb.ExperimentPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&experimentPermissionsPb,
	)
	resp, err := ExperimentPermissionsFromPb(&experimentPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *experimentsImpl) UpdateRun(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error) {
	requestPb, pbErr := UpdateRunToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateRunResponsePb mlpb.UpdateRunResponsePb
	path := "/api/2.0/mlflow/runs/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&updateRunResponsePb,
	)
	resp, err := UpdateRunResponseFromPb(&updateRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just FeatureStore API methods
type featureStoreImpl struct {
	client *client.DatabricksClient
}

func (a *featureStoreImpl) CreateOnlineStore(ctx context.Context, request CreateOnlineStoreRequest) (*OnlineStore, error) {
	requestPb, pbErr := CreateOnlineStoreRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var onlineStorePb mlpb.OnlineStorePb
	path := "/api/2.0/feature-store/online-stores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.OnlineStore,
		&onlineStorePb,
	)
	resp, err := OnlineStoreFromPb(&onlineStorePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *featureStoreImpl) DeleteOnlineStore(ctx context.Context, request DeleteOnlineStoreRequest) error {
	requestPb, pbErr := DeleteOnlineStoreRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/feature-store/online-stores/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *featureStoreImpl) GetOnlineStore(ctx context.Context, request GetOnlineStoreRequest) (*OnlineStore, error) {
	requestPb, pbErr := GetOnlineStoreRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var onlineStorePb mlpb.OnlineStorePb
	path := fmt.Sprintf("/api/2.0/feature-store/online-stores/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&onlineStorePb,
	)
	resp, err := OnlineStoreFromPb(&onlineStorePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List Online Feature Stores.
func (a *featureStoreImpl) ListOnlineStores(ctx context.Context, request ListOnlineStoresRequest) listing.Iterator[OnlineStore] {

	getNextPage := func(ctx context.Context, req ListOnlineStoresRequest) (*ListOnlineStoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListOnlineStores(ctx, req)
	}
	getItems := func(resp *ListOnlineStoresResponse) []OnlineStore {
		return resp.OnlineStores
	}
	getNextReq := func(resp *ListOnlineStoresResponse) *ListOnlineStoresRequest {
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

// List Online Feature Stores.
func (a *featureStoreImpl) ListOnlineStoresAll(ctx context.Context, request ListOnlineStoresRequest) ([]OnlineStore, error) {
	iterator := a.ListOnlineStores(ctx, request)
	return listing.ToSlice[OnlineStore](ctx, iterator)
}

func (a *featureStoreImpl) internalListOnlineStores(ctx context.Context, request ListOnlineStoresRequest) (*ListOnlineStoresResponse, error) {
	requestPb, pbErr := ListOnlineStoresRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listOnlineStoresResponsePb mlpb.ListOnlineStoresResponsePb
	path := "/api/2.0/feature-store/online-stores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listOnlineStoresResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListOnlineStoresResponseFromPb(&listOnlineStoresResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *featureStoreImpl) PublishTable(ctx context.Context, request PublishTableRequest) (*PublishTableResponse, error) {
	requestPb, pbErr := PublishTableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var publishTableResponsePb mlpb.PublishTableResponsePb
	path := fmt.Sprintf("/api/2.0/feature-store/tables/%v/publish", request.SourceTableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&publishTableResponsePb,
	)
	resp, err := PublishTableResponseFromPb(&publishTableResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *featureStoreImpl) UpdateOnlineStore(ctx context.Context, request UpdateOnlineStoreRequest) (*OnlineStore, error) {
	requestPb, pbErr := UpdateOnlineStoreRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var onlineStorePb mlpb.OnlineStorePb
	path := fmt.Sprintf("/api/2.0/feature-store/online-stores/%v", request.Name)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.OnlineStore,
		&onlineStorePb,
	)
	resp, err := OnlineStoreFromPb(&onlineStorePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Forecasting API methods
type forecastingImpl struct {
	client *client.DatabricksClient
}

func (a *forecastingImpl) CreateExperiment(ctx context.Context, request CreateForecastingExperimentRequest) (*CreateForecastingExperimentResponse, error) {
	requestPb, pbErr := CreateForecastingExperimentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createForecastingExperimentResponsePb mlpb.CreateForecastingExperimentResponsePb
	path := "/api/2.0/automl/create-forecasting-experiment"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createForecastingExperimentResponsePb,
	)
	resp, err := CreateForecastingExperimentResponseFromPb(&createForecastingExperimentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *forecastingImpl) GetExperiment(ctx context.Context, request GetForecastingExperimentRequest) (*ForecastingExperiment, error) {
	requestPb, pbErr := GetForecastingExperimentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var forecastingExperimentPb mlpb.ForecastingExperimentPb
	path := fmt.Sprintf("/api/2.0/automl/get-forecasting-experiment/%v", request.ExperimentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&forecastingExperimentPb,
	)
	resp, err := ForecastingExperimentFromPb(&forecastingExperimentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just MaterializedFeatures API methods
type materializedFeaturesImpl struct {
	client *client.DatabricksClient
}

func (a *materializedFeaturesImpl) CreateFeatureTag(ctx context.Context, request CreateFeatureTagRequest) (*FeatureTag, error) {
	requestPb, pbErr := CreateFeatureTagRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var featureTagPb mlpb.FeatureTagPb
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags", request.TableName, request.FeatureName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.FeatureTag,
		&featureTagPb,
	)
	resp, err := FeatureTagFromPb(&featureTagPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *materializedFeaturesImpl) DeleteFeatureTag(ctx context.Context, request DeleteFeatureTagRequest) error {
	requestPb, pbErr := DeleteFeatureTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags/%v", request.TableName, request.FeatureName, request.Key)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *materializedFeaturesImpl) GetFeatureLineage(ctx context.Context, request GetFeatureLineageRequest) (*FeatureLineage, error) {
	requestPb, pbErr := GetFeatureLineageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var featureLineagePb mlpb.FeatureLineagePb
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/lineage", request.TableName, request.FeatureName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&featureLineagePb,
	)
	resp, err := FeatureLineageFromPb(&featureLineagePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *materializedFeaturesImpl) GetFeatureTag(ctx context.Context, request GetFeatureTagRequest) (*FeatureTag, error) {
	requestPb, pbErr := GetFeatureTagRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var featureTagPb mlpb.FeatureTagPb
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags/%v", request.TableName, request.FeatureName, request.Key)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&featureTagPb,
	)
	resp, err := FeatureTagFromPb(&featureTagPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists FeatureTags.
func (a *materializedFeaturesImpl) ListFeatureTags(ctx context.Context, request ListFeatureTagsRequest) listing.Iterator[FeatureTag] {

	getNextPage := func(ctx context.Context, req ListFeatureTagsRequest) (*ListFeatureTagsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListFeatureTags(ctx, req)
	}
	getItems := func(resp *ListFeatureTagsResponse) []FeatureTag {
		return resp.FeatureTags
	}
	getNextReq := func(resp *ListFeatureTagsResponse) *ListFeatureTagsRequest {
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

// Lists FeatureTags.
func (a *materializedFeaturesImpl) ListFeatureTagsAll(ctx context.Context, request ListFeatureTagsRequest) ([]FeatureTag, error) {
	iterator := a.ListFeatureTags(ctx, request)
	return listing.ToSlice[FeatureTag](ctx, iterator)
}

func (a *materializedFeaturesImpl) internalListFeatureTags(ctx context.Context, request ListFeatureTagsRequest) (*ListFeatureTagsResponse, error) {
	requestPb, pbErr := ListFeatureTagsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listFeatureTagsResponsePb mlpb.ListFeatureTagsResponsePb
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags", request.TableName, request.FeatureName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listFeatureTagsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListFeatureTagsResponseFromPb(&listFeatureTagsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *materializedFeaturesImpl) UpdateFeatureTag(ctx context.Context, request UpdateFeatureTagRequest) (*FeatureTag, error) {
	requestPb, pbErr := UpdateFeatureTagRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var featureTagPb mlpb.FeatureTagPb
	path := fmt.Sprintf("/api/2.0/feature-store/feature-tables/%v/features/%v/tags/%v", request.TableName, request.FeatureName, request.Key)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" || slices.Contains(requestPb.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.FeatureTag,
		&featureTagPb,
	)
	resp, err := FeatureTagFromPb(&featureTagPb)
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
	requestPb, pbErr := ApproveTransitionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var approveTransitionRequestResponsePb mlpb.ApproveTransitionRequestResponsePb
	path := "/api/2.0/mlflow/transition-requests/approve"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&approveTransitionRequestResponsePb,
	)
	resp, err := ApproveTransitionRequestResponseFromPb(&approveTransitionRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateComment(ctx context.Context, request CreateComment) (*CreateCommentResponse, error) {
	requestPb, pbErr := CreateCommentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createCommentResponsePb mlpb.CreateCommentResponsePb
	path := "/api/2.0/mlflow/comments/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createCommentResponsePb,
	)
	resp, err := CreateCommentResponseFromPb(&createCommentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateModel(ctx context.Context, request CreateModelRequest) (*CreateModelResponse, error) {
	requestPb, pbErr := CreateModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createModelResponsePb mlpb.CreateModelResponsePb
	path := "/api/2.0/mlflow/registered-models/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createModelResponsePb,
	)
	resp, err := CreateModelResponseFromPb(&createModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateModelVersion(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	requestPb, pbErr := CreateModelVersionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createModelVersionResponsePb mlpb.CreateModelVersionResponsePb
	path := "/api/2.0/mlflow/model-versions/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createModelVersionResponsePb,
	)
	resp, err := CreateModelVersionResponseFromPb(&createModelVersionResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateTransitionRequest(ctx context.Context, request CreateTransitionRequest) (*CreateTransitionRequestResponse, error) {
	requestPb, pbErr := CreateTransitionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createTransitionRequestResponsePb mlpb.CreateTransitionRequestResponsePb
	path := "/api/2.0/mlflow/transition-requests/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createTransitionRequestResponsePb,
	)
	resp, err := CreateTransitionRequestResponseFromPb(&createTransitionRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) CreateWebhook(ctx context.Context, request CreateRegistryWebhook) (*CreateWebhookResponse, error) {
	requestPb, pbErr := CreateRegistryWebhookToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createWebhookResponsePb mlpb.CreateWebhookResponsePb
	path := "/api/2.0/mlflow/registry-webhooks/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createWebhookResponsePb,
	)
	resp, err := CreateWebhookResponseFromPb(&createWebhookResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) DeleteComment(ctx context.Context, request DeleteCommentRequest) error {
	requestPb, pbErr := DeleteCommentRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/comments/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *modelRegistryImpl) DeleteModel(ctx context.Context, request DeleteModelRequest) error {
	requestPb, pbErr := DeleteModelRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/registered-models/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *modelRegistryImpl) DeleteModelTag(ctx context.Context, request DeleteModelTagRequest) error {
	requestPb, pbErr := DeleteModelTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/registered-models/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *modelRegistryImpl) DeleteModelVersion(ctx context.Context, request DeleteModelVersionRequest) error {
	requestPb, pbErr := DeleteModelVersionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/model-versions/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *modelRegistryImpl) DeleteModelVersionTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	requestPb, pbErr := DeleteModelVersionTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/model-versions/delete-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *modelRegistryImpl) DeleteTransitionRequest(ctx context.Context, request DeleteTransitionRequestRequest) (*DeleteTransitionRequestResponse, error) {
	requestPb, pbErr := DeleteTransitionRequestRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteTransitionRequestResponsePb mlpb.DeleteTransitionRequestResponsePb
	path := "/api/2.0/mlflow/transition-requests/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteTransitionRequestResponsePb,
	)
	resp, err := DeleteTransitionRequestResponseFromPb(&deleteTransitionRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) DeleteWebhook(ctx context.Context, request DeleteWebhookRequest) error {
	requestPb, pbErr := DeleteWebhookRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/registry-webhooks/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

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

// Gets the latest version of a registered model.
func (a *modelRegistryImpl) GetLatestVersionsAll(ctx context.Context, request GetLatestVersionsRequest) ([]ModelVersion, error) {
	iterator := a.GetLatestVersions(ctx, request)
	return listing.ToSlice[ModelVersion](ctx, iterator)
}

func (a *modelRegistryImpl) internalGetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
	requestPb, pbErr := GetLatestVersionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getLatestVersionsResponsePb mlpb.GetLatestVersionsResponsePb
	path := "/api/2.0/mlflow/registered-models/get-latest-versions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&getLatestVersionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := GetLatestVersionsResponseFromPb(&getLatestVersionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetModel(ctx context.Context, request GetModelRequest) (*GetModelResponse, error) {
	requestPb, pbErr := GetModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getModelResponsePb mlpb.GetModelResponsePb
	path := "/api/2.0/mlflow/databricks/registered-models/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getModelResponsePb,
	)
	resp, err := GetModelResponseFromPb(&getModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetModelVersion(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	requestPb, pbErr := GetModelVersionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getModelVersionResponsePb mlpb.GetModelVersionResponsePb
	path := "/api/2.0/mlflow/model-versions/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getModelVersionResponsePb,
	)
	resp, err := GetModelVersionResponseFromPb(&getModelVersionResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetModelVersionDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	requestPb, pbErr := GetModelVersionDownloadUriRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getModelVersionDownloadUriResponsePb mlpb.GetModelVersionDownloadUriResponsePb
	path := "/api/2.0/mlflow/model-versions/get-download-uri"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getModelVersionDownloadUriResponsePb,
	)
	resp, err := GetModelVersionDownloadUriResponseFromPb(&getModelVersionDownloadUriResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetPermissionLevels(ctx context.Context, request GetRegisteredModelPermissionLevelsRequest) (*GetRegisteredModelPermissionLevelsResponse, error) {
	requestPb, pbErr := GetRegisteredModelPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRegisteredModelPermissionLevelsResponsePb mlpb.GetRegisteredModelPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v/permissionLevels", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getRegisteredModelPermissionLevelsResponsePb,
	)
	resp, err := GetRegisteredModelPermissionLevelsResponseFromPb(&getRegisteredModelPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) GetPermissions(ctx context.Context, request GetRegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	requestPb, pbErr := GetRegisteredModelPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelPermissionsPb mlpb.RegisteredModelPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&registeredModelPermissionsPb,
	)
	resp, err := RegisteredModelPermissionsFromPb(&registeredModelPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Lists all available registered models, up to the limit specified in
// __max_results__.
func (a *modelRegistryImpl) ListModelsAll(ctx context.Context, request ListModelsRequest) ([]Model, error) {
	iterator := a.ListModels(ctx, request)
	return listing.ToSlice[Model](ctx, iterator)
}

func (a *modelRegistryImpl) internalListModels(ctx context.Context, request ListModelsRequest) (*ListModelsResponse, error) {
	requestPb, pbErr := ListModelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listModelsResponsePb mlpb.ListModelsResponsePb
	path := "/api/2.0/mlflow/registered-models/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listModelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListModelsResponseFromPb(&listModelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets a list of all open stage transition requests for the model version.
func (a *modelRegistryImpl) ListTransitionRequestsAll(ctx context.Context, request ListTransitionRequestsRequest) ([]Activity, error) {
	iterator := a.ListTransitionRequests(ctx, request)
	return listing.ToSlice[Activity](ctx, iterator)
}

func (a *modelRegistryImpl) internalListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) (*ListTransitionRequestsResponse, error) {
	requestPb, pbErr := ListTransitionRequestsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listTransitionRequestsResponsePb mlpb.ListTransitionRequestsResponsePb
	path := "/api/2.0/mlflow/transition-requests/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listTransitionRequestsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListTransitionRequestsResponseFromPb(&listTransitionRequestsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// **NOTE:** This endpoint is in Public Preview. Lists all registry webhooks.
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

// **NOTE:** This endpoint is in Public Preview. Lists all registry webhooks.
func (a *modelRegistryImpl) ListWebhooksAll(ctx context.Context, request ListWebhooksRequest) ([]RegistryWebhook, error) {
	iterator := a.ListWebhooks(ctx, request)
	return listing.ToSlice[RegistryWebhook](ctx, iterator)
}

func (a *modelRegistryImpl) internalListWebhooks(ctx context.Context, request ListWebhooksRequest) (*ListRegistryWebhooks, error) {
	requestPb, pbErr := ListWebhooksRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listRegistryWebhooksPb mlpb.ListRegistryWebhooksPb
	path := "/api/2.0/mlflow/registry-webhooks/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listRegistryWebhooksPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListRegistryWebhooksFromPb(&listRegistryWebhooksPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) RejectTransitionRequest(ctx context.Context, request RejectTransitionRequest) (*RejectTransitionRequestResponse, error) {
	requestPb, pbErr := RejectTransitionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var rejectTransitionRequestResponsePb mlpb.RejectTransitionRequestResponsePb
	path := "/api/2.0/mlflow/transition-requests/reject"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&rejectTransitionRequestResponsePb,
	)
	resp, err := RejectTransitionRequestResponseFromPb(&rejectTransitionRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) RenameModel(ctx context.Context, request RenameModelRequest) (*RenameModelResponse, error) {
	requestPb, pbErr := RenameModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var renameModelResponsePb mlpb.RenameModelResponsePb
	path := "/api/2.0/mlflow/registered-models/rename"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&renameModelResponsePb,
	)
	resp, err := RenameModelResponseFromPb(&renameModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Searches for specific model versions based on the supplied __filter__.
func (a *modelRegistryImpl) SearchModelVersionsAll(ctx context.Context, request SearchModelVersionsRequest) ([]ModelVersion, error) {
	iterator := a.SearchModelVersions(ctx, request)
	return listing.ToSlice[ModelVersion](ctx, iterator)
}

func (a *modelRegistryImpl) internalSearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
	requestPb, pbErr := SearchModelVersionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchModelVersionsResponsePb mlpb.SearchModelVersionsResponsePb
	path := "/api/2.0/mlflow/model-versions/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&searchModelVersionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := SearchModelVersionsResponseFromPb(&searchModelVersionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Search for registered models based on the specified __filter__.
func (a *modelRegistryImpl) SearchModelsAll(ctx context.Context, request SearchModelsRequest) ([]Model, error) {
	iterator := a.SearchModels(ctx, request)
	return listing.ToSlice[Model](ctx, iterator)
}

func (a *modelRegistryImpl) internalSearchModels(ctx context.Context, request SearchModelsRequest) (*SearchModelsResponse, error) {
	requestPb, pbErr := SearchModelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var searchModelsResponsePb mlpb.SearchModelsResponsePb
	path := "/api/2.0/mlflow/registered-models/search"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&searchModelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := SearchModelsResponseFromPb(&searchModelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) SetModelTag(ctx context.Context, request SetModelTagRequest) error {
	requestPb, pbErr := SetModelTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/registered-models/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *modelRegistryImpl) SetModelVersionTag(ctx context.Context, request SetModelVersionTagRequest) error {
	requestPb, pbErr := SetModelVersionTagRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/mlflow/model-versions/set-tag"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *modelRegistryImpl) SetPermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	requestPb, pbErr := RegisteredModelPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelPermissionsPb mlpb.RegisteredModelPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&registeredModelPermissionsPb,
	)
	resp, err := RegisteredModelPermissionsFromPb(&registeredModelPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) TestRegistryWebhook(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	requestPb, pbErr := TestRegistryWebhookRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var testRegistryWebhookResponsePb mlpb.TestRegistryWebhookResponsePb
	path := "/api/2.0/mlflow/registry-webhooks/test"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&testRegistryWebhookResponsePb,
	)
	resp, err := TestRegistryWebhookResponseFromPb(&testRegistryWebhookResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error) {
	requestPb, pbErr := TransitionModelVersionStageDatabricksToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var transitionStageResponsePb mlpb.TransitionStageResponsePb
	path := "/api/2.0/mlflow/databricks/model-versions/transition-stage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&transitionStageResponsePb,
	)
	resp, err := TransitionStageResponseFromPb(&transitionStageResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) UpdateComment(ctx context.Context, request UpdateComment) (*UpdateCommentResponse, error) {
	requestPb, pbErr := UpdateCommentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateCommentResponsePb mlpb.UpdateCommentResponsePb
	path := "/api/2.0/mlflow/comments/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updateCommentResponsePb,
	)
	resp, err := UpdateCommentResponseFromPb(&updateCommentResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) UpdateModel(ctx context.Context, request UpdateModelRequest) (*UpdateModelResponse, error) {
	requestPb, pbErr := UpdateModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateModelResponsePb mlpb.UpdateModelResponsePb
	path := "/api/2.0/mlflow/registered-models/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updateModelResponsePb,
	)
	resp, err := UpdateModelResponseFromPb(&updateModelResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) UpdateModelVersion(ctx context.Context, request UpdateModelVersionRequest) (*UpdateModelVersionResponse, error) {
	requestPb, pbErr := UpdateModelVersionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateModelVersionResponsePb mlpb.UpdateModelVersionResponsePb
	path := "/api/2.0/mlflow/model-versions/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updateModelVersionResponsePb,
	)
	resp, err := UpdateModelVersionResponseFromPb(&updateModelVersionResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) UpdatePermissions(ctx context.Context, request RegisteredModelPermissionsRequest) (*RegisteredModelPermissions, error) {
	requestPb, pbErr := RegisteredModelPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelPermissionsPb mlpb.RegisteredModelPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&registeredModelPermissionsPb,
	)
	resp, err := RegisteredModelPermissionsFromPb(&registeredModelPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelRegistryImpl) UpdateWebhook(ctx context.Context, request UpdateRegistryWebhook) (*UpdateWebhookResponse, error) {
	requestPb, pbErr := UpdateRegistryWebhookToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateWebhookResponsePb mlpb.UpdateWebhookResponsePb
	path := "/api/2.0/mlflow/registry-webhooks/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updateWebhookResponsePb,
	)
	resp, err := UpdateWebhookResponseFromPb(&updateWebhookResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
