// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/jobs/jobspb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Jobs API methods
type jobsImpl struct {
	client *client.DatabricksClient
}

func (a *jobsImpl) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	requestPb, pbErr := CancelAllRunsToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.2/jobs/runs/cancel-all"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *jobsImpl) CancelRun(ctx context.Context, request CancelRun) error {
	requestPb, pbErr := CancelRunToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.2/jobs/runs/cancel"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *jobsImpl) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	requestPb, pbErr := CreateJobToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createResponsePb jobspb.CreateResponsePb
	path := "/api/2.2/jobs/create"
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
		&createResponsePb,
	)
	resp, err := CreateResponseFromPb(&createResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Delete(ctx context.Context, request DeleteJob) error {
	requestPb, pbErr := DeleteJobToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.2/jobs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *jobsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {
	requestPb, pbErr := DeleteRunToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.2/jobs/runs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *jobsImpl) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	requestPb, pbErr := ExportRunRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var exportRunOutputPb jobspb.ExportRunOutputPb
	path := "/api/2.2/jobs/runs/export"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&exportRunOutputPb,
	)
	resp, err := ExportRunOutputFromPb(&exportRunOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	requestPb, pbErr := GetJobRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var jobPb jobspb.JobPb
	path := "/api/2.2/jobs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&jobPb,
	)
	resp, err := JobFromPb(&jobPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error) {
	requestPb, pbErr := GetJobPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getJobPermissionLevelsResponsePb jobspb.GetJobPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v/permissionLevels", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getJobPermissionLevelsResponsePb,
	)
	resp, err := GetJobPermissionLevelsResponseFromPb(&getJobPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error) {
	requestPb, pbErr := GetJobPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var jobPermissionsPb jobspb.JobPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&jobPermissionsPb,
	)
	resp, err := JobPermissionsFromPb(&jobPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	requestPb, pbErr := GetRunRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var runPb jobspb.RunPb
	path := "/api/2.2/jobs/runs/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&runPb,
	)
	resp, err := RunFromPb(&runPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	requestPb, pbErr := GetRunOutputRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var runOutputPb jobspb.RunOutputPb
	path := "/api/2.2/jobs/runs/get-output"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&runOutputPb,
	)
	resp, err := RunOutputFromPb(&runOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Retrieves a list of jobs.
func (a *jobsImpl) List(ctx context.Context, request ListJobsRequest) listing.Iterator[BaseJob] {

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

// Retrieves a list of jobs.
func (a *jobsImpl) ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BaseJob](ctx, iterator)
}

func (a *jobsImpl) internalList(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {
	requestPb, pbErr := ListJobsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listJobsResponsePb jobspb.ListJobsResponsePb
	path := "/api/2.2/jobs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listJobsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListJobsResponseFromPb(&listJobsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List runs in descending order by start time.
func (a *jobsImpl) ListRuns(ctx context.Context, request ListRunsRequest) listing.Iterator[BaseRun] {

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

// List runs in descending order by start time.
func (a *jobsImpl) ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error) {
	iterator := a.ListRuns(ctx, request)
	return listing.ToSlice[BaseRun](ctx, iterator)
}

func (a *jobsImpl) internalListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	requestPb, pbErr := ListRunsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listRunsResponsePb jobspb.ListRunsResponsePb
	path := "/api/2.2/jobs/runs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listRunsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListRunsResponseFromPb(&listRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	requestPb, pbErr := RepairRunToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var repairRunResponsePb jobspb.RepairRunResponsePb
	path := "/api/2.2/jobs/runs/repair"
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
		&repairRunResponsePb,
	)
	resp, err := RepairRunResponseFromPb(&repairRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Reset(ctx context.Context, request ResetJob) error {
	requestPb, pbErr := ResetJobToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.2/jobs/reset"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *jobsImpl) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	requestPb, pbErr := RunNowToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var runNowResponsePb jobspb.RunNowResponsePb
	path := "/api/2.2/jobs/run-now"
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
		&runNowResponsePb,
	)
	resp, err := RunNowResponseFromPb(&runNowResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {
	requestPb, pbErr := JobPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var jobPermissionsPb jobspb.JobPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
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
		&jobPermissionsPb,
	)
	resp, err := JobPermissionsFromPb(&jobPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	requestPb, pbErr := SubmitRunToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var submitRunResponsePb jobspb.SubmitRunResponsePb
	path := "/api/2.2/jobs/runs/submit"
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
		&submitRunResponsePb,
	)
	resp, err := SubmitRunResponseFromPb(&submitRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Update(ctx context.Context, request UpdateJob) error {
	requestPb, pbErr := UpdateJobToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.2/jobs/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *jobsImpl) UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {
	requestPb, pbErr := JobPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var jobPermissionsPb jobspb.JobPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
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
		&jobPermissionsPb,
	)
	resp, err := JobPermissionsFromPb(&jobPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just PolicyComplianceForJobs API methods
type policyComplianceForJobsImpl struct {
	client *client.DatabricksClient
}

func (a *policyComplianceForJobsImpl) EnforceCompliance(ctx context.Context, request EnforcePolicyComplianceRequest) (*EnforcePolicyComplianceResponse, error) {
	requestPb, pbErr := EnforcePolicyComplianceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enforcePolicyComplianceResponsePb jobspb.EnforcePolicyComplianceResponsePb
	path := "/api/2.0/policies/jobs/enforce-compliance"
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
		&enforcePolicyComplianceResponsePb,
	)
	resp, err := EnforcePolicyComplianceResponseFromPb(&enforcePolicyComplianceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *policyComplianceForJobsImpl) GetCompliance(ctx context.Context, request GetPolicyComplianceRequest) (*GetPolicyComplianceResponse, error) {
	requestPb, pbErr := GetPolicyComplianceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPolicyComplianceResponsePb jobspb.GetPolicyComplianceResponsePb
	path := "/api/2.0/policies/jobs/get-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getPolicyComplianceResponsePb,
	)
	resp, err := GetPolicyComplianceResponseFromPb(&getPolicyComplianceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Returns the policy compliance status of all jobs that use a given policy.
// Jobs could be out of compliance if a cluster policy they use was updated
// after the job was last edited and its job clusters no longer comply with the
// updated policy.
func (a *policyComplianceForJobsImpl) ListCompliance(ctx context.Context, request ListJobComplianceRequest) listing.Iterator[JobCompliance] {

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

// Returns the policy compliance status of all jobs that use a given policy.
// Jobs could be out of compliance if a cluster policy they use was updated
// after the job was last edited and its job clusters no longer comply with the
// updated policy.
func (a *policyComplianceForJobsImpl) ListComplianceAll(ctx context.Context, request ListJobComplianceRequest) ([]JobCompliance, error) {
	iterator := a.ListCompliance(ctx, request)
	return listing.ToSlice[JobCompliance](ctx, iterator)
}

func (a *policyComplianceForJobsImpl) internalListCompliance(ctx context.Context, request ListJobComplianceRequest) (*ListJobComplianceForPolicyResponse, error) {
	requestPb, pbErr := ListJobComplianceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listJobComplianceForPolicyResponsePb jobspb.ListJobComplianceForPolicyResponsePb
	path := "/api/2.0/policies/jobs/list-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listJobComplianceForPolicyResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListJobComplianceForPolicyResponseFromPb(&listJobComplianceForPolicyResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
