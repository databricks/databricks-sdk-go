// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
)


// The Jobs Runs API allows you to submit, repair, and cancel running jobs.
type JobRunsService interface {
    // Cancels a job run. The run is canceled asynchronously, so it may still be
    // running when this request completes.
    Cancel(ctx context.Context, cancelRun CancelRun) error
    // Cancels all active runs of a job. The runs are canceled asynchronously,
    // so it doesn&#39;t prevent new runs from being started.
    CancelAll(ctx context.Context, cancelAllRuns CancelAllRuns) error
    // Deletes a non-active run. Returns an error if the run is active.
    Delete(ctx context.Context, deleteRun DeleteRun) error
    // Export and retrieve the job run task.
    Export(ctx context.Context, exportRunRequest ExportRunRequest) (*ExportRunOutput, error)
    // Retrieve the metadata of a run.
    Get(ctx context.Context, getRunRequest GetRunRequest) (*Run, error)
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
    GetOutput(ctx context.Context, getRunOutputRequest GetRunOutputRequest) (*RunOutput, error)
    // List runs in descending order by start time.
    List(ctx context.Context, listRunsRequest ListRunsRequest) (*ListRunsResponse, error)
    // Re-run one or more tasks. Tasks are re-run as part of the original job
    // run, use the current job and task settings, and can be viewed in the
    // history for the original job run.
    Repair(ctx context.Context, repairRun RepairRun) (*RepairRunResponse, error)
    // Submit a one-time run. This endpoint allows you to submit a workload
    // directly without creating a job. Runs submitted using this endpoint don?t
    // display in the UI. Use the `jobs/runs/get` API to check the run state
    // after the job is submitted.
    Submit(ctx context.Context, submitRun SubmitRun) (*SubmitRunResponse, error)
}

// The Jobs API allows you to create, edit, and delete jobs.
type JobsService interface {
    // Create a new job.
    Create(ctx context.Context, createJob CreateJob) (*CreateJobResponse, error)
    // Deletes a job.
    Delete(ctx context.Context, deleteJob DeleteJob) error
    // Retrieves the details for a single job.
    Get(ctx context.Context, getJobRequest GetJobRequest) (*Job, error)
    // Retrieves a list of jobs.
    List(ctx context.Context, listJobsRequest ListJobsRequest) (*ListJobsResponse, error)
    // Overwrites all the settings for a specific job. Use the Update endpoint
    // to update job settings partially.
    Reset(ctx context.Context, resetJob ResetJob) error
    // Run a job and return the `run_id` of the triggered run.
    RunNow(ctx context.Context, runNow RunNow) (*RunNowResponse, error)
    // Add, update, or remove specific settings of an existing job. Use the
    // ResetJob to overwrite all job settings.
    Update(ctx context.Context, updateJob UpdateJob) error
}
