// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Jobs API methods
type jobsImpl struct {
	client *client.DatabricksClient
}

func (a *jobsImpl) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	path := "/api/2.1/jobs/runs/cancel-all"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *jobsImpl) CancelRun(ctx context.Context, request CancelRun) error {
	path := "/api/2.1/jobs/runs/cancel"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *jobsImpl) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.1/jobs/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createResponse)
	return &createResponse, err
}

func (a *jobsImpl) Delete(ctx context.Context, request DeleteJob) error {
	path := "/api/2.1/jobs/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *jobsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	path := "/api/2.1/jobs/runs/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *jobsImpl) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	var exportRunOutput ExportRunOutput
	path := "/api/2.1/jobs/runs/export"
	err := a.client.Do(ctx, http.MethodGet, path, request, &exportRunOutput)
	return &exportRunOutput, err
}

func (a *jobsImpl) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	var job Job
	path := "/api/2.1/jobs/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &job)
	return &job, err
}

func (a *jobsImpl) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	var run Run
	path := "/api/2.1/jobs/runs/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &run)
	return &run, err
}

func (a *jobsImpl) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	var runOutput RunOutput
	path := "/api/2.1/jobs/runs/get-output"
	err := a.client.Do(ctx, http.MethodGet, path, request, &runOutput)
	return &runOutput, err
}

func (a *jobsImpl) List(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {
	var listJobsResponse ListJobsResponse
	path := "/api/2.1/jobs/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listJobsResponse)
	return &listJobsResponse, err
}

func (a *jobsImpl) ListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	var listRunsResponse ListRunsResponse
	path := "/api/2.1/jobs/runs/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listRunsResponse)
	return &listRunsResponse, err
}

func (a *jobsImpl) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	var repairRunResponse RepairRunResponse
	path := "/api/2.1/jobs/runs/repair"
	err := a.client.Do(ctx, http.MethodPost, path, request, &repairRunResponse)
	return &repairRunResponse, err
}

func (a *jobsImpl) Reset(ctx context.Context, request ResetJob) error {
	path := "/api/2.1/jobs/reset"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *jobsImpl) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	var runNowResponse RunNowResponse
	path := "/api/2.1/jobs/run-now"
	err := a.client.Do(ctx, http.MethodPost, path, request, &runNowResponse)
	return &runNowResponse, err
}

func (a *jobsImpl) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	var submitRunResponse SubmitRunResponse
	path := "/api/2.1/jobs/runs/submit"
	err := a.client.Do(ctx, http.MethodPost, path, request, &submitRunResponse)
	return &submitRunResponse, err
}

func (a *jobsImpl) Update(ctx context.Context, request UpdateJob) error {
	path := "/api/2.1/jobs/update"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}
