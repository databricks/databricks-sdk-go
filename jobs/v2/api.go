// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Jobs , Policy Compliance For Jobs , etc.
package jobs

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type JobsInterface interface {

	// Cancel all runs of a job.
	//
	// Cancels all active runs of a job. The runs are canceled asynchronously, so it
	// doesn't prevent new runs from being started.
	CancelAllRuns(ctx context.Context, request CancelAllRuns) error

	// Cancel a run.
	//
	// Cancels a job run or a task run. The run is canceled asynchronously, so it
	// may still be running when this request completes.
	CancelRun(ctx context.Context, request CancelRun) error

	// Cancel a run.
	//
	// Cancels a job run or a task run. The run is canceled asynchronously, so it
	// may still be running when this request completes.
	CancelRunByRunId(ctx context.Context, runId int64) error

	// Create a new job.
	//
	// Create a new job.
	Create(ctx context.Context, request CreateJob) (*CreateResponse, error)

	// Delete a job.
	//
	// Deletes a job.
	Delete(ctx context.Context, request DeleteJob) error

	// Delete a job.
	//
	// Deletes a job.
	DeleteByJobId(ctx context.Context, jobId int64) error

	// Delete a job run.
	//
	// Deletes a non-active run. Returns an error if the run is active.
	DeleteRun(ctx context.Context, request DeleteRun) error

	// Delete a job run.
	//
	// Deletes a non-active run. Returns an error if the run is active.
	DeleteRunByRunId(ctx context.Context, runId int64) error

	// Export and retrieve a job run.
	//
	// Export and retrieve the job run task.
	ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error)

	// Get a single job.
	//
	// Retrieves the details for a single job.
	//
	// In Jobs API 2.2, requests for a single job support pagination of `tasks` and
	// `job_clusters` when either exceeds 100 elements. Use the `next_page_token`
	// field to check for more results and pass its value as the `page_token` in
	// subsequent requests. Arrays with fewer than 100 elements in a page will be
	// empty on later pages.
	Get(ctx context.Context, request GetJobRequest) (*Job, error)

	// Get a single job.
	//
	// Retrieves the details for a single job.
	//
	// In Jobs API 2.2, requests for a single job support pagination of `tasks` and
	// `job_clusters` when either exceeds 100 elements. Use the `next_page_token`
	// field to check for more results and pass its value as the `page_token` in
	// subsequent requests. Arrays with fewer than 100 elements in a page will be
	// empty on later pages.
	GetByJobId(ctx context.Context, jobId int64) (*Job, error)

	// Get job permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error)

	// Get job permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByJobId(ctx context.Context, jobId string) (*GetJobPermissionLevelsResponse, error)

	// Get job permissions.
	//
	// Gets the permissions of a job. Jobs can inherit permissions from their root
	// object.
	GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error)

	// Get job permissions.
	//
	// Gets the permissions of a job. Jobs can inherit permissions from their root
	// object.
	GetPermissionsByJobId(ctx context.Context, jobId string) (*JobPermissions, error)

	// Get a single job run.
	//
	// Retrieves the metadata of a run.
	//
	// In Jobs API 2.2, requests for a single job run support pagination of `tasks`
	// and `job_clusters` when either exceeds 100 elements. Use the
	// `next_page_token` field to check for more results and pass its value as the
	// `page_token` in subsequent requests. Arrays with fewer than 100 elements in a
	// page will be empty on later pages.
	GetRun(ctx context.Context, request GetRunRequest) (*Run, error)

	// Get the output for a single run.
	//
	// Retrieve the output and metadata of a single task run. When a notebook task
	// returns a value through the `dbutils.notebook.exit()` call, you can use this
	// endpoint to retrieve that value. Databricks restricts this API to returning
	// the first 5 MB of the output. To return a larger result, you can store job
	// results in a cloud storage service.
	//
	// This endpoint validates that the __run_id__ parameter is valid and returns an
	// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
	// automatically removed after 60 days. If you to want to reference them beyond
	// 60 days, you must save old run results before they expire.
	GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error)

	// Get the output for a single run.
	//
	// Retrieve the output and metadata of a single task run. When a notebook task
	// returns a value through the `dbutils.notebook.exit()` call, you can use this
	// endpoint to retrieve that value. Databricks restricts this API to returning
	// the first 5 MB of the output. To return a larger result, you can store job
	// results in a cloud storage service.
	//
	// This endpoint validates that the __run_id__ parameter is valid and returns an
	// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
	// automatically removed after 60 days. If you to want to reference them beyond
	// 60 days, you must save old run results before they expire.
	GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error)

	// List jobs.
	//
	// Retrieves a list of jobs.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListJobsRequest) listing.Iterator[BaseJob]

	// List jobs.
	//
	// Retrieves a list of jobs.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error)

	// BaseJobSettingsNameToJobIdMap calls [JobsAPI.ListAll] and creates a map of results with [BaseJob].Settings.Name as key and [BaseJob].JobId as value.
	//
	// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
	//
	// Note: All [BaseJob] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	BaseJobSettingsNameToJobIdMap(ctx context.Context, request ListJobsRequest) (map[string]int64, error)

	// GetBySettingsName calls [JobsAPI.BaseJobSettingsNameToJobIdMap] and returns a single [BaseJob].
	//
	// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
	//
	// Note: All [BaseJob] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetBySettingsName(ctx context.Context, name string) (*BaseJob, error)

	// List job runs.
	//
	// List runs in descending order by start time.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRuns(ctx context.Context, request ListRunsRequest) listing.Iterator[BaseRun]

	// List job runs.
	//
	// List runs in descending order by start time.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error)

	// Repair a job run.
	//
	// Re-run one or more tasks. Tasks are re-run as part of the original job run.
	// They use the current job and task settings, and can be viewed in the history
	// for the original job run.
	RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error)

	// Update all job settings (reset).
	//
	// Overwrite all settings for the given job. Use the [_Update_
	// endpoint](:method:jobs/update) to update job settings partially.
	Reset(ctx context.Context, request ResetJob) error

	// Trigger a new job run.
	//
	// Run a job and return the `run_id` of the triggered run.
	RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error)

	// Set job permissions.
	//
	// Sets permissions on an object, replacing existing permissions if they exist.
	// Deletes all direct permissions if none are specified. Objects can inherit
	// permissions from their root object.
	SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error)

	// Create and trigger a one-time run.
	//
	// Submit a one-time run. This endpoint allows you to submit a workload directly
	// without creating a job. Runs submitted using this endpoint don’t display in
	// the UI. Use the `jobs/runs/get` API to check the run state after the job is
	// submitted.
	Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error)

	// Update job settings partially.
	//
	// Add, update, or remove specific settings of an existing job. Use the [_Reset_
	// endpoint](:method:jobs/reset) to overwrite all job settings.
	Update(ctx context.Context, request UpdateJob) error

	// Update job permissions.
	//
	// Updates the permissions on a job. Jobs can inherit permissions from their
	// root object.
	UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error)
}

func NewJobs(client *client.DatabricksClient) *JobsAPI {
	return &JobsAPI{
		jobsImpl: jobsImpl{
			client: client,
		},
	}
}

// The Jobs API allows you to create, edit, and delete jobs.
//
// You can use a Databricks job to run a data processing or data analysis task
// in a Databricks cluster with scalable resources. Your job can consist of a
// single task or can be a large, multi-task workflow with complex dependencies.
// Databricks manages the task orchestration, cluster management, monitoring,
// and error reporting for all of your jobs. You can run your jobs immediately
// or periodically through an easy-to-use scheduling system. You can implement
// job tasks using notebooks, JARS, Delta Live Tables pipelines, or Python,
// Scala, Spark submit, and Java applications.
//
// You should never hard code secrets or store them in plain text. Use the
// [Secrets CLI] to manage secrets in the [Databricks CLI]. Use the [Secrets
// utility] to reference secrets in notebooks and jobs.
//
// [Databricks CLI]: https://docs.databricks.com/dev-tools/cli/index.html
// [Secrets CLI]: https://docs.databricks.com/dev-tools/cli/secrets-cli.html
// [Secrets utility]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-secrets
type JobsAPI struct {
	jobsImpl
}

// Cancel a run.
//
// Cancels a job run or a task run. The run is canceled asynchronously, so it
// may still be running when this request completes.
func (a *JobsAPI) CancelRunByRunId(ctx context.Context, runId int64) error {
	return a.jobsImpl.CancelRun(ctx, CancelRun{
		RunId: runId,
	})
}

// Delete a job.
//
// Deletes a job.
func (a *JobsAPI) DeleteByJobId(ctx context.Context, jobId int64) error {
	return a.jobsImpl.Delete(ctx, DeleteJob{
		JobId: jobId,
	})
}

// Delete a job run.
//
// Deletes a non-active run. Returns an error if the run is active.
func (a *JobsAPI) DeleteRunByRunId(ctx context.Context, runId int64) error {
	return a.jobsImpl.DeleteRun(ctx, DeleteRun{
		RunId: runId,
	})
}

// Get a single job.
//
// Retrieves the details for a single job.
//
// In Jobs API 2.2, requests for a single job support pagination of `tasks` and
// `job_clusters` when either exceeds 100 elements. Use the `next_page_token`
// field to check for more results and pass its value as the `page_token` in
// subsequent requests. Arrays with fewer than 100 elements in a page will be
// empty on later pages.
func (a *JobsAPI) GetByJobId(ctx context.Context, jobId int64) (*Job, error) {
	return a.jobsImpl.Get(ctx, GetJobRequest{
		JobId: jobId,
	})
}

// Get job permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *JobsAPI) GetPermissionLevelsByJobId(ctx context.Context, jobId string) (*GetJobPermissionLevelsResponse, error) {
	return a.jobsImpl.GetPermissionLevels(ctx, GetJobPermissionLevelsRequest{
		JobId: jobId,
	})
}

// Get job permissions.
//
// Gets the permissions of a job. Jobs can inherit permissions from their root
// object.
func (a *JobsAPI) GetPermissionsByJobId(ctx context.Context, jobId string) (*JobPermissions, error) {
	return a.jobsImpl.GetPermissions(ctx, GetJobPermissionsRequest{
		JobId: jobId,
	})
}

// Get the output for a single run.
//
// Retrieve the output and metadata of a single task run. When a notebook task
// returns a value through the `dbutils.notebook.exit()` call, you can use this
// endpoint to retrieve that value. Databricks restricts this API to returning
// the first 5 MB of the output. To return a larger result, you can store job
// results in a cloud storage service.
//
// This endpoint validates that the __run_id__ parameter is valid and returns an
// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
// automatically removed after 60 days. If you to want to reference them beyond
// 60 days, you must save old run results before they expire.
func (a *JobsAPI) GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error) {
	return a.jobsImpl.GetRunOutput(ctx, GetRunOutputRequest{
		RunId: runId,
	})
}

// List jobs.
//
// Retrieves a list of jobs.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) List(ctx context.Context, request ListJobsRequest) listing.Iterator[BaseJob] {

	getNextPage := func(ctx context.Context, req ListJobsRequest) (*ListJobsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.jobsImpl.List(ctx, req)
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
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BaseJob](ctx, iterator)
}

// BaseJobSettingsNameToJobIdMap calls [JobsAPI.ListAll] and creates a map of results with [BaseJob].Settings.Name as key and [BaseJob].JobId as value.
//
// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
//
// Note: All [BaseJob] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) BaseJobSettingsNameToJobIdMap(ctx context.Context, request ListJobsRequest) (map[string]int64, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]int64{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Settings.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Settings.Name: %s", key)
		}
		mapping[key] = v.JobId
	}
	return mapping, nil
}

// GetBySettingsName calls [JobsAPI.BaseJobSettingsNameToJobIdMap] and returns a single [BaseJob].
//
// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
//
// Note: All [BaseJob] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) GetBySettingsName(ctx context.Context, name string) (*BaseJob, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListJobsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]BaseJob{}
	for _, v := range result {
		key := v.Settings.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("BaseJob named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of BaseJob named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// List job runs.
//
// List runs in descending order by start time.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListRuns(ctx context.Context, request ListRunsRequest) listing.Iterator[BaseRun] {

	getNextPage := func(ctx context.Context, req ListRunsRequest) (*ListRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.jobsImpl.ListRuns(ctx, req)
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
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error) {
	iterator := a.ListRuns(ctx, request)
	return listing.ToSlice[BaseRun](ctx, iterator)
}

type PolicyComplianceForJobsInterface interface {

	// Enforce job policy compliance.
	//
	// Updates a job so the job clusters that are created when running the job
	// (specified in `new_cluster`) are compliant with the current versions of their
	// respective cluster policies. All-purpose clusters used in the job will not be
	// updated.
	EnforceCompliance(ctx context.Context, request EnforcePolicyComplianceRequest) (*EnforcePolicyComplianceResponse, error)

	// Get job policy compliance.
	//
	// Returns the policy compliance status of a job. Jobs could be out of
	// compliance if a cluster policy they use was updated after the job was last
	// edited and some of its job clusters no longer comply with their updated
	// policies.
	GetCompliance(ctx context.Context, request GetPolicyComplianceRequest) (*GetPolicyComplianceResponse, error)

	// Get job policy compliance.
	//
	// Returns the policy compliance status of a job. Jobs could be out of
	// compliance if a cluster policy they use was updated after the job was last
	// edited and some of its job clusters no longer comply with their updated
	// policies.
	GetComplianceByJobId(ctx context.Context, jobId int64) (*GetPolicyComplianceResponse, error)

	// List job policy compliance.
	//
	// Returns the policy compliance status of all jobs that use a given policy.
	// Jobs could be out of compliance if a cluster policy they use was updated
	// after the job was last edited and its job clusters no longer comply with the
	// updated policy.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListCompliance(ctx context.Context, request ListJobComplianceRequest) listing.Iterator[JobCompliance]

	// List job policy compliance.
	//
	// Returns the policy compliance status of all jobs that use a given policy.
	// Jobs could be out of compliance if a cluster policy they use was updated
	// after the job was last edited and its job clusters no longer comply with the
	// updated policy.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListComplianceAll(ctx context.Context, request ListJobComplianceRequest) ([]JobCompliance, error)
}

func NewPolicyComplianceForJobs(client *client.DatabricksClient) *PolicyComplianceForJobsAPI {
	return &PolicyComplianceForJobsAPI{
		policyComplianceForJobsImpl: policyComplianceForJobsImpl{
			client: client,
		},
	}
}

// The compliance APIs allow you to view and manage the policy compliance status
// of jobs in your workspace. This API currently only supports compliance
// controls for cluster policies.
//
// A job is in compliance if its cluster configurations satisfy the rules of all
// their respective cluster policies. A job could be out of compliance if a
// cluster policy it uses was updated after the job was last edited. The job is
// considered out of compliance if any of its clusters no longer comply with
// their updated policies.
//
// The get and list compliance APIs allow you to view the policy compliance
// status of a job. The enforce compliance API allows you to update a job so
// that it becomes compliant with all of its policies.
type PolicyComplianceForJobsAPI struct {
	policyComplianceForJobsImpl
}

// Get job policy compliance.
//
// Returns the policy compliance status of a job. Jobs could be out of
// compliance if a cluster policy they use was updated after the job was last
// edited and some of its job clusters no longer comply with their updated
// policies.
func (a *PolicyComplianceForJobsAPI) GetComplianceByJobId(ctx context.Context, jobId int64) (*GetPolicyComplianceResponse, error) {
	return a.policyComplianceForJobsImpl.GetCompliance(ctx, GetPolicyComplianceRequest{
		JobId: jobId,
	})
}

// List job policy compliance.
//
// Returns the policy compliance status of all jobs that use a given policy.
// Jobs could be out of compliance if a cluster policy they use was updated
// after the job was last edited and its job clusters no longer comply with the
// updated policy.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PolicyComplianceForJobsAPI) ListCompliance(ctx context.Context, request ListJobComplianceRequest) listing.Iterator[JobCompliance] {

	getNextPage := func(ctx context.Context, req ListJobComplianceRequest) (*ListJobComplianceForPolicyResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.policyComplianceForJobsImpl.ListCompliance(ctx, req)
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
//
// This method is generated by Databricks SDK Code Generator.
func (a *PolicyComplianceForJobsAPI) ListComplianceAll(ctx context.Context, request ListJobComplianceRequest) ([]JobCompliance, error) {
	iterator := a.ListCompliance(ctx, request)
	return listing.ToSlice[JobCompliance](ctx, iterator)
}
