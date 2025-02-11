// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobspreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just JobsPreview API methods
type jobsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *jobsPreviewImpl) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	var cancelAllRunsResponse CancelAllRunsResponse
	path := "/api/2.2preview/jobs/runs/cancel-all"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelAllRunsResponse)
	return err
}

func (a *jobsPreviewImpl) CancelRun(ctx context.Context, request CancelRun) error {
	var cancelRunResponse CancelRunResponse
	path := "/api/2.2preview/jobs/runs/cancel"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelRunResponse)
	return err
}

func (a *jobsPreviewImpl) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.2preview/jobs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createResponse)
	return &createResponse, err
}

func (a *jobsPreviewImpl) Delete(ctx context.Context, request DeleteJob) error {
	var deleteResponse DeleteResponse
	path := "/api/2.2preview/jobs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *jobsPreviewImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	var deleteRunResponse DeleteRunResponse
	path := "/api/2.2preview/jobs/runs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteRunResponse)
	return err
}

func (a *jobsPreviewImpl) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	var exportRunOutput ExportRunOutput
	path := "/api/2.2preview/jobs/runs/export"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &exportRunOutput)
	return &exportRunOutput, err
}

func (a *jobsPreviewImpl) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	var job Job
	path := "/api/2.2preview/jobs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &job)
	return &job, err
}

func (a *jobsPreviewImpl) GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error) {
	var getJobPermissionLevelsResponse GetJobPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0preview/permissions/jobs/%v/permissionLevels", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getJobPermissionLevelsResponse)
	return &getJobPermissionLevelsResponse, err
}

func (a *jobsPreviewImpl) GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/jobs/%v", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &jobPermissions)
	return &jobPermissions, err
}

func (a *jobsPreviewImpl) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	var run Run
	path := "/api/2.2preview/jobs/runs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &run)
	return &run, err
}

func (a *jobsPreviewImpl) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	var runOutput RunOutput
	path := "/api/2.2preview/jobs/runs/get-output"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &runOutput)
	return &runOutput, err
}

// List jobs.
//
// Retrieves a list of jobs.
func (a *jobsPreviewImpl) List(ctx context.Context, request ListJobsRequest) listing.Iterator[BaseJob] {

	getNextPage := func(ctx context.Context, req ListJobsRequest) (*ListJobsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListJobsResponse) []BaseJob {
		return resp.Jobs
	}
	getNextReq := func(resp *ListJobsResponse) *ListJobsRequest {
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

// List jobs.
//
// Retrieves a list of jobs.
func (a *jobsPreviewImpl) ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BaseJob](ctx, iterator)
}
func (a *jobsPreviewImpl) internalList(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {
	var listJobsResponse ListJobsResponse
	path := "/api/2.2preview/jobs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listJobsResponse)
	return &listJobsResponse, err
}

// List job runs.
//
// List runs in descending order by start time.
func (a *jobsPreviewImpl) ListRuns(ctx context.Context, request ListRunsRequest) listing.Iterator[BaseRun] {

	getNextPage := func(ctx context.Context, req ListRunsRequest) (*ListRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListRuns(ctx, req)
	}
	getItems := func(resp *ListRunsResponse) []BaseRun {
		return resp.Runs
	}
	getNextReq := func(resp *ListRunsResponse) *ListRunsRequest {
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

// List job runs.
//
// List runs in descending order by start time.
func (a *jobsPreviewImpl) ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error) {
	iterator := a.ListRuns(ctx, request)
	return listing.ToSlice[BaseRun](ctx, iterator)
}
func (a *jobsPreviewImpl) internalListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	var listRunsResponse ListRunsResponse
	path := "/api/2.2preview/jobs/runs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRunsResponse)
	return &listRunsResponse, err
}

func (a *jobsPreviewImpl) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	var repairRunResponse RepairRunResponse
	path := "/api/2.2preview/jobs/runs/repair"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &repairRunResponse)
	return &repairRunResponse, err
}

func (a *jobsPreviewImpl) Reset(ctx context.Context, request ResetJob) error {
	var resetResponse ResetResponse
	path := "/api/2.2preview/jobs/reset"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resetResponse)
	return err
}

func (a *jobsPreviewImpl) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	var runNowResponse RunNowResponse
	path := "/api/2.2preview/jobs/run-now"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &runNowResponse)
	return &runNowResponse, err
}

func (a *jobsPreviewImpl) SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/jobs/%v", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &jobPermissions)
	return &jobPermissions, err
}

func (a *jobsPreviewImpl) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	var submitRunResponse SubmitRunResponse
	path := "/api/2.2preview/jobs/runs/submit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &submitRunResponse)
	return &submitRunResponse, err
}

func (a *jobsPreviewImpl) Update(ctx context.Context, request UpdateJob) error {
	var updateResponse UpdateResponse
	path := "/api/2.2preview/jobs/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &updateResponse)
	return err
}

func (a *jobsPreviewImpl) UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {
	var jobPermissions JobPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/jobs/%v", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &jobPermissions)
	return &jobPermissions, err
}

// unexported type that holds implementations of just PolicyComplianceForJobsPreview API methods
type policyComplianceForJobsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *policyComplianceForJobsPreviewImpl) EnforceCompliance(ctx context.Context, request EnforcePolicyComplianceRequest) (*EnforcePolicyComplianceResponse, error) {
	var enforcePolicyComplianceResponse EnforcePolicyComplianceResponse
	path := "/api/2.0preview/policies/jobs/enforce-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &enforcePolicyComplianceResponse)
	return &enforcePolicyComplianceResponse, err
}

func (a *policyComplianceForJobsPreviewImpl) GetCompliance(ctx context.Context, request GetPolicyComplianceRequest) (*GetPolicyComplianceResponse, error) {
	var getPolicyComplianceResponse GetPolicyComplianceResponse
	path := "/api/2.0preview/policies/jobs/get-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPolicyComplianceResponse)
	return &getPolicyComplianceResponse, err
}

// List job policy compliance.
//
// Returns the policy compliance status of all jobs that use a given policy.
// Jobs could be out of compliance if a cluster policy they use was updated
// after the job was last edited and its job clusters no longer comply with the
// updated policy.
func (a *policyComplianceForJobsPreviewImpl) ListCompliance(ctx context.Context, request ListJobComplianceRequest) listing.Iterator[JobCompliance] {

	getNextPage := func(ctx context.Context, req ListJobComplianceRequest) (*ListJobComplianceForPolicyResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCompliance(ctx, req)
	}
	getItems := func(resp *ListJobComplianceForPolicyResponse) []JobCompliance {
		return resp.Jobs
	}
	getNextReq := func(resp *ListJobComplianceForPolicyResponse) *ListJobComplianceRequest {
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

// List job policy compliance.
//
// Returns the policy compliance status of all jobs that use a given policy.
// Jobs could be out of compliance if a cluster policy they use was updated
// after the job was last edited and its job clusters no longer comply with the
// updated policy.
func (a *policyComplianceForJobsPreviewImpl) ListComplianceAll(ctx context.Context, request ListJobComplianceRequest) ([]JobCompliance, error) {
	iterator := a.ListCompliance(ctx, request)
	return listing.ToSlice[JobCompliance](ctx, iterator)
}
func (a *policyComplianceForJobsPreviewImpl) internalListCompliance(ctx context.Context, request ListJobComplianceRequest) (*ListJobComplianceForPolicyResponse, error) {
	var listJobComplianceForPolicyResponse ListJobComplianceForPolicyResponse
	path := "/api/2.0preview/policies/jobs/list-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listJobComplianceForPolicyResponse)
	return &listJobComplianceForPolicyResponse, err
}
