// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Jobs API methods
type jobsImpl struct {
	client *client.DatabricksClient
}

func (a *jobsImpl) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {

	requestPb, pbErr := cancelAllRunsToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var cancelAllRunsResponsePb cancelAllRunsResponsePb
	path := "/api/2.2/jobs/runs/cancel-all"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&cancelAllRunsResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *jobsImpl) CancelRun(ctx context.Context, request CancelRun) error {

	requestPb, pbErr := cancelRunToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var cancelRunResponsePb cancelRunResponsePb
	path := "/api/2.2/jobs/runs/cancel"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&cancelRunResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *jobsImpl) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {

	requestPb, pbErr := createJobToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createResponsePb createResponsePb
	path := "/api/2.2/jobs/create"
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
		&createResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createResponseFromPb(&createResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Delete(ctx context.Context, request DeleteJob) error {

	requestPb, pbErr := deleteJobToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := "/api/2.2/jobs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *jobsImpl) DeleteRun(ctx context.Context, request DeleteRun) error {

	requestPb, pbErr := deleteRunToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteRunResponsePb deleteRunResponsePb
	path := "/api/2.2/jobs/runs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
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

func (a *jobsImpl) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {

	requestPb, pbErr := exportRunRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var exportRunOutputPb exportRunOutputPb
	path := "/api/2.2/jobs/runs/export"
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
		&exportRunOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := exportRunOutputFromPb(&exportRunOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Get(ctx context.Context, request GetJobRequest) (*Job, error) {

	requestPb, pbErr := getJobRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var jobPb jobPb
	path := "/api/2.2/jobs/get"
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
		&jobPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := jobFromPb(&jobPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error) {

	requestPb, pbErr := getJobPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getJobPermissionLevelsResponsePb getJobPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v/permissionLevels", requestPb.JobId)
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
		&getJobPermissionLevelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getJobPermissionLevelsResponseFromPb(&getJobPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error) {

	requestPb, pbErr := getJobPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var jobPermissionsPb jobPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", requestPb.JobId)
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
		&jobPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := jobPermissionsFromPb(&jobPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {

	requestPb, pbErr := getRunRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var runPb runPb
	path := "/api/2.2/jobs/runs/get"
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
		&runPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := runFromPb(&runPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {

	requestPb, pbErr := getRunOutputRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var runOutputPb runOutputPb
	path := "/api/2.2/jobs/runs/get-output"
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
		&runOutputPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := runOutputFromPb(&runOutputPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List jobs.
//
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

// List jobs.
//
// Retrieves a list of jobs.
func (a *jobsImpl) ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BaseJob](ctx, iterator)
}

func (a *jobsImpl) internalList(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {

	requestPb, pbErr := listJobsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listJobsResponsePb listJobsResponsePb
	path := "/api/2.2/jobs/list"
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
		&listJobsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listJobsResponseFromPb(&listJobsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List job runs.
//
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

// List job runs.
//
// List runs in descending order by start time.
func (a *jobsImpl) ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error) {
	iterator := a.ListRuns(ctx, request)
	return listing.ToSlice[BaseRun](ctx, iterator)
}

func (a *jobsImpl) internalListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {

	requestPb, pbErr := listRunsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listRunsResponsePb listRunsResponsePb
	path := "/api/2.2/jobs/runs/list"
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
		&listRunsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listRunsResponseFromPb(&listRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {

	requestPb, pbErr := repairRunToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var repairRunResponsePb repairRunResponsePb
	path := "/api/2.2/jobs/runs/repair"
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
		&repairRunResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := repairRunResponseFromPb(&repairRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Reset(ctx context.Context, request ResetJob) error {

	requestPb, pbErr := resetJobToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var resetResponsePb resetResponsePb
	path := "/api/2.2/jobs/reset"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&resetResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *jobsImpl) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {

	requestPb, pbErr := runNowToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var runNowResponsePb runNowResponsePb
	path := "/api/2.2/jobs/run-now"
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
		&runNowResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := runNowResponseFromPb(&runNowResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {

	requestPb, pbErr := jobPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var jobPermissionsPb jobPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", requestPb.JobId)
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
		&jobPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := jobPermissionsFromPb(&jobPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {

	requestPb, pbErr := submitRunToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var submitRunResponsePb submitRunResponsePb
	path := "/api/2.2/jobs/runs/submit"
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
		&submitRunResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := submitRunResponseFromPb(&submitRunResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *jobsImpl) Update(ctx context.Context, request UpdateJob) error {

	requestPb, pbErr := updateJobToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateResponsePb updateResponsePb
	path := "/api/2.2/jobs/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *jobsImpl) UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error) {

	requestPb, pbErr := jobPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var jobPermissionsPb jobPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", requestPb.JobId)
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
		&jobPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := jobPermissionsFromPb(&jobPermissionsPb)
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

	requestPb, pbErr := enforcePolicyComplianceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enforcePolicyComplianceResponsePb enforcePolicyComplianceResponsePb
	path := "/api/2.0/policies/jobs/enforce-compliance"
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
		&enforcePolicyComplianceResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := enforcePolicyComplianceResponseFromPb(&enforcePolicyComplianceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *policyComplianceForJobsImpl) GetCompliance(ctx context.Context, request GetPolicyComplianceRequest) (*GetPolicyComplianceResponse, error) {

	requestPb, pbErr := getPolicyComplianceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPolicyComplianceResponsePb getPolicyComplianceResponsePb
	path := "/api/2.0/policies/jobs/get-compliance"
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
		&getPolicyComplianceResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getPolicyComplianceResponseFromPb(&getPolicyComplianceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List job policy compliance.
//
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

// List job policy compliance.
//
// Returns the policy compliance status of all jobs that use a given policy.
// Jobs could be out of compliance if a cluster policy they use was updated
// after the job was last edited and its job clusters no longer comply with the
// updated policy.
func (a *policyComplianceForJobsImpl) ListComplianceAll(ctx context.Context, request ListJobComplianceRequest) ([]JobCompliance, error) {
	iterator := a.ListCompliance(ctx, request)
	return listing.ToSlice[JobCompliance](ctx, iterator)
}

func (a *policyComplianceForJobsImpl) internalListCompliance(ctx context.Context, request ListJobComplianceRequest) (*ListJobComplianceForPolicyResponse, error) {

	requestPb, pbErr := listJobComplianceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listJobComplianceForPolicyResponsePb listJobComplianceForPolicyResponsePb
	path := "/api/2.0/policies/jobs/list-compliance"
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
		&listJobComplianceForPolicyResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listJobComplianceForPolicyResponseFromPb(&listJobComplianceForPolicyResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
