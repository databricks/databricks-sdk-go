// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// The Jobs API allows you to create, edit, and delete jobs.
type JobsService interface {
    // Cancels all active runs of a job. The runs are canceled asynchronously,
    // so it doesn&#39;t prevent new runs from being started.
    CancelAllRuns(ctx context.Context, cancelAllRunsRequest CancelAllRunsRequest) error
    // Cancels a job run. The run is canceled asynchronously, so it may still be
    // running when this request completes.
    CancelRun(ctx context.Context, cancelRunRequest CancelRunRequest) error
    // Create a new job.
    CreateJob(ctx context.Context, createJobRequest CreateJobRequest) (*CreateJobResponse, error)
    // Deletes a job.
    DeleteJob(ctx context.Context, deleteJobRequest DeleteJobRequest) error
    // Deletes a non-active run. Returns an error if the run is active.
    DeleteRun(ctx context.Context, deleteRunRequest DeleteRunRequest) error
    // Export and retrieve the job run task.
    ExportRun(ctx context.Context, exportRunRequest ExportRunRequest) (*ExportRunResponse, error)
    // Retrieves the details for a single job.
    GetJob(ctx context.Context, getJobRequest GetJobRequest) (*GetJobResponse, error)
    // Retrieve the metadata of a run.
    GetRun(ctx context.Context, getRunRequest GetRunRequest) (*GetRunResponse, error)
    // Retrieve the output and metadata of a single task run. When a notebook
    // task returns a value through the dbutils.notebook.exit() call, you can
    // use this endpoint to retrieve that value. jobs restricts this API to
    // return the first 5 MB of the output. To return a larger result, you can
    // store job results in a cloud storage service. This endpoint validates
    // that the run_id parameter is valid and returns an HTTP status code 400 if
    // the run_id parameter is invalid. Runs are automatically removed after 60
    // days. If you to want to reference them beyond 60 days, you must save old
    // run results before they expire. To export using the UI, see Export job
    // run results. To export using the Jobs API, see Runs export.
    GetRunOutput(ctx context.Context, getRunOutputRequest GetRunOutputRequest) (*GetRunOutputResponse, error)
    // Retrieves a list of jobs.
    ListJobs(ctx context.Context, listJobsRequest ListJobsRequest) (*ListJobsResponse, error)
    // List runs in descending order by start time.
    ListRuns(ctx context.Context, listRunsRequest ListRunsRequest) (*ListRunsResponse, error)
    // Re-run one or more tasks. Tasks are re-run as part of the original job
    // run, use the current job and task settings, and can be viewed in the
    // history for the original job run.
    RepairRun(ctx context.Context, repairRunRequest RepairRunRequest) (*RepairRunResponse, error)
    // Overwrites all the settings for a specific job. Use the Update endpoint
    // to update job settings partially.
    ResetJob(ctx context.Context, resetJobRequest ResetJobRequest) error
    // Run a job and return the `run_id` of the triggered run.
    RunNow(ctx context.Context, runNowRequest RunNowRequest) (*RunNowResponse, error)
    // Submit a one-time run. This endpoint allows you to submit a workload
    // directly without creating a job. Runs submitted using this endpoint don?t
    // display in the UI. Use the `jobs/runs/get` API to check the run state
    // after the job is submitted.
    SubmitRun(ctx context.Context, submitRunRequest SubmitRunRequest) (*SubmitRunResponse, error)
    // Add, update, or remove specific settings of an existing job. Use the
    // ResetJob to overwrite all job settings.
    UpdateJob(ctx context.Context, updateJobRequest UpdateJobRequest) error
}

func New(client *client.DatabricksClient) JobsService {
	return &JobsAPI{
		client: client,
	}
}

type JobsAPI struct {
	client *client.DatabricksClient
}

// Cancels all active runs of a job. The runs are canceled asynchronously, so it
// doesn&#39;t prevent new runs from being started.
func (a *JobsAPI) CancelAllRuns(ctx context.Context, request CancelAllRunsRequest) error {
	path := "/api/2.1/jobs/runs/cancel-all"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Cancels a job run. The run is canceled asynchronously, so it may still be
// running when this request completes.
func (a *JobsAPI) CancelRun(ctx context.Context, request CancelRunRequest) error {
	path := "/api/2.1/jobs/runs/cancel"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Create a new job.
func (a *JobsAPI) CreateJob(ctx context.Context, request CreateJobRequest) (*CreateJobResponse, error) {
	var createJobResponse CreateJobResponse
	path := "/api/2.1/jobs/create"
	err := a.client.Post(ctx, path, request, &createJobResponse)
	return &createJobResponse, err
}

// Deletes a job.
func (a *JobsAPI) DeleteJob(ctx context.Context, request DeleteJobRequest) error {
	path := "/api/2.1/jobs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Deletes a non-active run. Returns an error if the run is active.
func (a *JobsAPI) DeleteRun(ctx context.Context, request DeleteRunRequest) error {
	path := "/api/2.1/jobs/runs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Export and retrieve the job run task.
func (a *JobsAPI) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunResponse, error) {
	var exportRunResponse ExportRunResponse
	path := "/api/2.1/jobs/runs/export"
	err := a.client.Get(ctx, path, request, &exportRunResponse)
	return &exportRunResponse, err
}

// Retrieves the details for a single job.
func (a *JobsAPI) GetJob(ctx context.Context, request GetJobRequest) (*GetJobResponse, error) {
	var getJobResponse GetJobResponse
	path := "/api/2.1/jobs/get"
	err := a.client.Get(ctx, path, request, &getJobResponse)
	return &getJobResponse, err
}

// Retrieve the metadata of a run.
func (a *JobsAPI) GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error) {
	var getRunResponse GetRunResponse
	path := "/api/2.1/jobs/runs/get"
	err := a.client.Get(ctx, path, request, &getRunResponse)
	return &getRunResponse, err
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
func (a *JobsAPI) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*GetRunOutputResponse, error) {
	var getRunOutputResponse GetRunOutputResponse
	path := "/api/2.1/jobs/runs/get-output"
	err := a.client.Get(ctx, path, request, &getRunOutputResponse)
	return &getRunOutputResponse, err
}

// Retrieves a list of jobs.
func (a *JobsAPI) ListJobs(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {
	var listJobsResponse ListJobsResponse
	path := "/api/2.1/jobs/list"
	err := a.client.Get(ctx, path, request, &listJobsResponse)
	return &listJobsResponse, err
}

// List runs in descending order by start time.
func (a *JobsAPI) ListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	var listRunsResponse ListRunsResponse
	path := "/api/2.1/jobs/runs/list"
	err := a.client.Get(ctx, path, request, &listRunsResponse)
	return &listRunsResponse, err
}

// Re-run one or more tasks. Tasks are re-run as part of the original job run,
// use the current job and task settings, and can be viewed in the history for
// the original job run.
func (a *JobsAPI) RepairRun(ctx context.Context, request RepairRunRequest) (*RepairRunResponse, error) {
	var repairRunResponse RepairRunResponse
	path := "/api/2.1/jobs/runs/repair"
	err := a.client.Post(ctx, path, request, &repairRunResponse)
	return &repairRunResponse, err
}

// Overwrites all the settings for a specific job. Use the Update endpoint to
// update job settings partially.
func (a *JobsAPI) ResetJob(ctx context.Context, request ResetJobRequest) error {
	path := "/api/2.1/jobs/reset"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Run a job and return the `run_id` of the triggered run.
func (a *JobsAPI) RunNow(ctx context.Context, request RunNowRequest) (*RunNowResponse, error) {
	var runNowResponse RunNowResponse
	path := "/api/2.1/jobs/run-now"
	err := a.client.Post(ctx, path, request, &runNowResponse)
	return &runNowResponse, err
}

// Submit a one-time run. This endpoint allows you to submit a workload directly
// without creating a job. Runs submitted using this endpoint don?t display in
// the UI. Use the `jobs/runs/get` API to check the run state after the job is
// submitted.
func (a *JobsAPI) SubmitRun(ctx context.Context, request SubmitRunRequest) (*SubmitRunResponse, error) {
	var submitRunResponse SubmitRunResponse
	path := "/api/2.1/jobs/runs/submit"
	err := a.client.Post(ctx, path, request, &submitRunResponse)
	return &submitRunResponse, err
}

// Add, update, or remove specific settings of an existing job. Use the ResetJob
// to overwrite all job settings.
func (a *JobsAPI) UpdateJob(ctx context.Context, request UpdateJobRequest) error {
	path := "/api/2.1/jobs/update"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

