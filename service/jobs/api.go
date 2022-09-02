// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewJobRuns(client *client.DatabricksClient) JobRunsService {
	return &JobRunsAPI{
		client: client,
	}
}

type JobRunsAPI struct {
	client *client.DatabricksClient
}

// Cancels a job run. The run is canceled asynchronously, so it may still be
// running when this request completes.
func (a *JobRunsAPI) Cancel(ctx context.Context, request CancelRun) error {
	path := "/api/2.1/jobs/runs/cancel"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Cancels all active runs of a job. The runs are canceled asynchronously, so it
// doesn&#39;t prevent new runs from being started.
func (a *JobRunsAPI) CancelAll(ctx context.Context, request CancelAllRuns) error {
	path := "/api/2.1/jobs/runs/cancel-all"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Deletes a non-active run. Returns an error if the run is active.
func (a *JobRunsAPI) Delete(ctx context.Context, request DeleteRun) error {
	path := "/api/2.1/jobs/runs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Export and retrieve the job run task.
func (a *JobRunsAPI) Export(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	var exportRunOutput ExportRunOutput
	path := "/api/2.1/jobs/runs/export"
	err := a.client.Get(ctx, path, request, &exportRunOutput)
	return &exportRunOutput, err
}

// Retrieve the metadata of a run.
func (a *JobRunsAPI) Get(ctx context.Context, request GetRunRequest) (*Run, error) {
	var run Run
	path := "/api/2.1/jobs/runs/get"
	err := a.client.Get(ctx, path, request, &run)
	return &run, err
}

// Retrieve the output and metadata of a single task run. When a notebook task
// returns a value through the dbutils.notebook.exit() call, you can use this
// endpoint to retrieve that value. jobs restricts this API to return the first
// 5 MB of the output. To return a larger result, you can store job results in a
// cloud storage service. This endpoint validates that the run_id parameter is
// valid and returns an HTTP status code 400 if the run_id parameter is invalid.
// Runs are automatically removed after 60 days. If you to want to reference
// them beyond 60 days, you must save old run results before they expire. To
// export using the UI, see Export job run results. To export using the Jobs
// API, see Runs export.
func (a *JobRunsAPI) GetOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	var runOutput RunOutput
	path := "/api/2.1/jobs/runs/get-output"
	err := a.client.Get(ctx, path, request, &runOutput)
	return &runOutput, err
}

// List runs in descending order by start time.
func (a *JobRunsAPI) List(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	var listRunsResponse ListRunsResponse
	path := "/api/2.1/jobs/runs/list"
	err := a.client.Get(ctx, path, request, &listRunsResponse)
	return &listRunsResponse, err
}

// Re-run one or more tasks. Tasks are re-run as part of the original job run,
// use the current job and task settings, and can be viewed in the history for
// the original job run.
func (a *JobRunsAPI) Repair(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	var repairRunResponse RepairRunResponse
	path := "/api/2.1/jobs/runs/repair"
	err := a.client.Post(ctx, path, request, &repairRunResponse)
	return &repairRunResponse, err
}

// Submit a one-time run. This endpoint allows you to submit a workload directly
// without creating a job. Runs submitted using this endpoint don?t display in
// the UI. Use the `jobs/runs/get` API to check the run state after the job is
// submitted.
func (a *JobRunsAPI) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	var submitRunResponse SubmitRunResponse
	path := "/api/2.1/jobs/runs/submit"
	err := a.client.Post(ctx, path, request, &submitRunResponse)
	return &submitRunResponse, err
}


func NewJobs(client *client.DatabricksClient) JobsService {
	return &JobsAPI{
		client: client,
	}
}

type JobsAPI struct {
	client *client.DatabricksClient
}

// Create a new job.
func (a *JobsAPI) Create(ctx context.Context, request CreateJob) (*CreateJobResponse, error) {
	var createJobResponse CreateJobResponse
	path := "/api/2.1/jobs/create"
	err := a.client.Post(ctx, path, request, &createJobResponse)
	return &createJobResponse, err
}

// Deletes a job.
func (a *JobsAPI) Delete(ctx context.Context, request DeleteJob) error {
	path := "/api/2.1/jobs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Retrieves the details for a single job.
func (a *JobsAPI) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	var job Job
	path := "/api/2.1/jobs/get"
	err := a.client.Get(ctx, path, request, &job)
	return &job, err
}

// Retrieves a list of jobs.
func (a *JobsAPI) List(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {
	var listJobsResponse ListJobsResponse
	path := "/api/2.1/jobs/list"
	err := a.client.Get(ctx, path, request, &listJobsResponse)
	return &listJobsResponse, err
}

// Overwrites all the settings for a specific job. Use the Update endpoint to
// update job settings partially.
func (a *JobsAPI) Reset(ctx context.Context, request ResetJob) error {
	path := "/api/2.1/jobs/reset"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Run a job and return the `run_id` of the triggered run.
func (a *JobsAPI) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	var runNowResponse RunNowResponse
	path := "/api/2.1/jobs/run-now"
	err := a.client.Post(ctx, path, request, &runNowResponse)
	return &runNowResponse, err
}

// Add, update, or remove specific settings of an existing job. Use the ResetJob
// to overwrite all job settings.
func (a *JobsAPI) Update(ctx context.Context, request UpdateJob) error {
	path := "/api/2.1/jobs/update"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

