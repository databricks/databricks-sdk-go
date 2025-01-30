// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just Jobs API methods
type jobsImpl struct {
	client *client.DatabricksClient
}

func (a *jobsImpl) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	var cancelAllRunsResponse CancelAllRunsResponse
	path := "/api/2.1/jobs/runs/cancel-all"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelAllRunsResponse)
	return err
}

func (a *jobsImpl) CancelRun(ctx context.Context, request CancelRun) error {
	var cancelRunResponse CancelRunResponse
	path := "/api/2.1/jobs/runs/cancel"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelRunResponse)
	return err
}

func (a *jobsImpl) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.1/jobs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createResponse)
	return &createResponse, err
}

func (a *jobsImpl) Delete(ctx context.Context, request DeleteJob) error {
	var deleteResponse DeleteResponse
	path := "/api/2.1/jobs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *jobsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	var deleteRunResponse DeleteRunResponse
	path := "/api/2.1/jobs/runs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteRunResponse)
	return err
}

func (a *jobsImpl) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	var exportRunOutput ExportRunOutput
	path := "/api/2.1/jobs/runs/export"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &exportRunOutput)
	return &exportRunOutput, err
}

func (a *jobsImpl) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	var job Job
	path := "/api/2.1/jobs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &job)
	return &job, err
}

func (a *jobsImpl) GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error) {
	var getJobPermissionLevelsResponse GetJobPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v/permissionLevels", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getJobPermissionLevelsResponse)
	return &getJobPermissionLevelsResponse, err
}

func (a *jobsImpl) GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &jobPermissions)
	return &jobPermissions, err
}

func (a *jobsImpl) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	var run Run
	path := "/api/2.1/jobs/runs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &run)
	return &run, err
}

func (a *jobsImpl) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	var runOutput RunOutput
	path := "/api/2.1/jobs/runs/get-output"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &runOutput)
	return &runOutput, err
}

func (a *jobsImpl) List(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {
	var listJobsResponse ListJobsResponse
	path := "/api/2.1/jobs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listJobsResponse)
	return &listJobsResponse, err
}

func (a *jobsImpl) ListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	var listRunsResponse ListRunsResponse
	path := "/api/2.1/jobs/runs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRunsResponse)
	return &listRunsResponse, err
}

func (a *jobsImpl) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	var repairRunResponse RepairRunResponse
	path := "/api/2.1/jobs/runs/repair"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &repairRunResponse)
	return &repairRunResponse, err
}

func (a *jobsImpl) Reset(ctx context.Context, request ResetJob) error {
	var resetResponse ResetResponse
	path := "/api/2.1/jobs/reset"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resetResponse)
	return err
}

func (a *jobsImpl) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	var runNowResponse RunNowResponse
	path := "/api/2.1/jobs/run-now"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &runNowResponse)
	return &runNowResponse, err
}

func (a *jobsImpl) SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &jobPermissions)
	return &jobPermissions, err
}

func (a *jobsImpl) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	var submitRunResponse SubmitRunResponse
	path := "/api/2.1/jobs/runs/submit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &submitRunResponse)
	return &submitRunResponse, err
}

func (a *jobsImpl) Update(ctx context.Context, request UpdateJob) error {
	var updateResponse UpdateResponse
	path := "/api/2.1/jobs/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &updateResponse)
	return err
}

func (a *jobsImpl) UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &jobPermissions)
	return &jobPermissions, err
}

// unexported type that holds implementations of just PolicyComplianceForJobs API methods
type policyComplianceForJobsImpl struct {
	client *client.DatabricksClient
}

func (a *policyComplianceForJobsImpl) EnforceCompliance(ctx context.Context, request EnforcePolicyComplianceRequest) (*EnforcePolicyComplianceResponse, error) {
	var enforcePolicyComplianceResponse EnforcePolicyComplianceResponse
	path := "/api/2.0/policies/jobs/enforce-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &enforcePolicyComplianceResponse)
	return &enforcePolicyComplianceResponse, err
}

func (a *policyComplianceForJobsImpl) GetCompliance(ctx context.Context, request GetPolicyComplianceRequest) (*GetPolicyComplianceResponse, error) {
	var getPolicyComplianceResponse GetPolicyComplianceResponse
	path := "/api/2.0/policies/jobs/get-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPolicyComplianceResponse)
	return &getPolicyComplianceResponse, err
}

func (a *policyComplianceForJobsImpl) ListCompliance(ctx context.Context, request ListJobComplianceRequest) (*ListJobComplianceForPolicyResponse, error) {
	var listJobComplianceForPolicyResponse ListJobComplianceForPolicyResponse
	path := "/api/2.0/policies/jobs/list-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listJobComplianceForPolicyResponse)
	return &listJobComplianceForPolicyResponse, err
}
