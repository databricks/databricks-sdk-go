// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Jobs API methods
type jobsImpl struct {
	client *client.DatabricksClient
}

func (a *jobsImpl) Client() client.DatabricksClientInterface {
	return a.client
}

func (a *jobsImpl) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	path := "/api/2.1/jobs/runs/cancel-all"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *jobsImpl) CancelRun(ctx context.Context, request CancelRun) error {
	path := "/api/2.1/jobs/runs/cancel"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *jobsImpl) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.1/jobs/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createResponse)
	return &createResponse, err
}

func (a *jobsImpl) Delete(ctx context.Context, request DeleteJob) error {
	path := "/api/2.1/jobs/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *jobsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	path := "/api/2.1/jobs/runs/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *jobsImpl) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	var exportRunOutput ExportRunOutput
	path := "/api/2.1/jobs/runs/export"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &exportRunOutput)
	return &exportRunOutput, err
}

func (a *jobsImpl) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	var job Job
	path := "/api/2.1/jobs/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &job)
	return &job, err
}

func (a *jobsImpl) GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error) {
	var getJobPermissionLevelsResponse GetJobPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v/permissionLevels", request.JobId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getJobPermissionLevelsResponse)
	return &getJobPermissionLevelsResponse, err
}

func (a *jobsImpl) GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &jobPermissions)
	return &jobPermissions, err
}

func (a *jobsImpl) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	var run Run
	path := "/api/2.1/jobs/runs/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &run)
	return &run, err
}

func (a *jobsImpl) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	var runOutput RunOutput
	path := "/api/2.1/jobs/runs/get-output"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &runOutput)
	return &runOutput, err
}

func (a *jobsImpl) List(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {
	var listJobsResponse ListJobsResponse
	path := "/api/2.1/jobs/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listJobsResponse)
	return &listJobsResponse, err
}

func (a *jobsImpl) ListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	var listRunsResponse ListRunsResponse
	path := "/api/2.1/jobs/runs/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listRunsResponse)
	return &listRunsResponse, err
}

func (a *jobsImpl) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	var repairRunResponse RepairRunResponse
	path := "/api/2.1/jobs/runs/repair"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &repairRunResponse)
	return &repairRunResponse, err
}

func (a *jobsImpl) Reset(ctx context.Context, request ResetJob) error {
	path := "/api/2.1/jobs/reset"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *jobsImpl) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	var runNowResponse RunNowResponse
	path := "/api/2.1/jobs/run-now"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &runNowResponse)
	return &runNowResponse, err
}

func (a *jobsImpl) SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &jobPermissions)
	return &jobPermissions, err
}

func (a *jobsImpl) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	var submitRunResponse SubmitRunResponse
	path := "/api/2.1/jobs/runs/submit"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &submitRunResponse)
	return &submitRunResponse, err
}

func (a *jobsImpl) Update(ctx context.Context, request UpdateJob) error {
	path := "/api/2.1/jobs/update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *jobsImpl) UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &jobPermissions)
	return &jobPermissions, err
}
