// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	"github.com/databricks/databricks-sdk-go/retries"
)

// The Jobs API allows you to create, edit, and delete jobs.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type JobsService interface {

	// Cancel all runs of a job
	//
	// Cancels all active runs of a job. The runs are canceled asynchronously,
	// so it doesn't prevent new runs from being started.
	CancelAllRuns(ctx context.Context, request CancelAllRuns) error

	// CancelAllRunsByJobId calls CancelAllRuns, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	CancelAllRunsByJobId(ctx context.Context, jobId int64) error

	// Cancel a job run
	//
	// Cancels a job run. The run is canceled asynchronously, so it may still be
	// running when this request completes.
	CancelRun(ctx context.Context, request CancelRun) error

	// CancelRunAndWait calls CancelRun() and waits to reach TERMINATED or SKIPPED state
	//
	// This method is generated by Databricks SDK Code Generator.
	CancelRunAndWait(ctx context.Context, request CancelRun, options ...retries.Option[Run]) (*Run, error)
	// CancelRunByRunId calls CancelRun, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	CancelRunByRunId(ctx context.Context, runId int64) error

	// CancelRunByRunIdAndWait calls CancelRunByRunId and waits until Run is in desired state.
	//
	// This method is generated by Databricks SDK Code Generator.
	CancelRunByRunIdAndWait(ctx context.Context, runId int64, options ...retries.Option[Run]) (*Run, error)

	// Create a new job
	//
	// Create a new job.
	Create(ctx context.Context, request CreateJob) (*CreateResponse, error)

	// Delete a job
	//
	// Deletes a job.
	Delete(ctx context.Context, request DeleteJob) error

	// DeleteByJobId calls Delete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteByJobId(ctx context.Context, jobId int64) error

	// Delete a job run
	//
	// Deletes a non-active run. Returns an error if the run is active.
	DeleteRun(ctx context.Context, request DeleteRun) error

	// DeleteRunByRunId calls DeleteRun, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteRunByRunId(ctx context.Context, runId int64) error

	// Export and retrieve a job run
	//
	// Export and retrieve the job run task.
	ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error)

	// Get a single job
	//
	// Retrieves the details for a single job.
	Get(ctx context.Context, request GetRequest) (*Job, error)

	// GetByJobId calls Get, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByJobId(ctx context.Context, jobId int64) (*Job, error)

	// Get a single job run
	//
	// Retrieve the metadata of a run.
	GetRun(ctx context.Context, request GetRunRequest) (*Run, error)

	// GetRunAndWait calls GetRun() and waits to reach TERMINATED or SKIPPED state
	//
	// This method is generated by Databricks SDK Code Generator.
	GetRunAndWait(ctx context.Context, request GetRunRequest, options ...retries.Option[Run]) (*Run, error)

	// Get the output for a single run
	//
	// Retrieve the output and metadata of a single task run. When a notebook
	// task returns a value through the `dbutils.notebook.exit()` call, you can
	// use this endpoint to retrieve that value. " + serviceName + " restricts
	// this API to returning the first 5 MB of the output. To return a larger
	// result, you can store job results in a cloud storage service.
	//
	// This endpoint validates that the __run_id__ parameter is valid and
	// returns an HTTP status code 400 if the __run_id__ parameter is invalid.
	// Runs are automatically removed after 60 days. If you to want to reference
	// them beyond 60 days, you must save old run results before they expire.
	GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error)

	// GetRunOutputByRunId calls GetRunOutput, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error)

	// List all jobs
	//
	// Retrieves a list of jobs.
	//
	// Use ListAll() to get all Job instances, which will iterate over every result page.
	List(ctx context.Context, request ListRequest) (*ListResponse, error)

	// ListAll calls List() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListRequest) ([]Job, error)

	// List runs for a job
	//
	// List runs in descending order by start time.
	//
	// Use ListRunsAll() to get all Run instances, which will iterate over every result page.
	ListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error)

	// ListRunsAll calls ListRuns() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRunsAll(ctx context.Context, request ListRunsRequest) ([]Run, error)

	// Repair a job run
	//
	// Re-run one or more tasks. Tasks are re-run as part of the original job
	// run. They use the current job and task settings, and can be viewed in the
	// history for the original job run.
	RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error)

	// RepairRunAndWait calls RepairRun() and waits to reach TERMINATED or SKIPPED state
	//
	// This method is generated by Databricks SDK Code Generator.
	RepairRunAndWait(ctx context.Context, request RepairRun, options ...retries.Option[Run]) (*Run, error)

	// Overwrites all settings for a job
	//
	// Overwrites all the settings for a specific job. Use the Update endpoint
	// to update job settings partially.
	Reset(ctx context.Context, request ResetJob) error

	// Trigger a new job run
	//
	// Run a job and return the `run_id` of the triggered run.
	RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error)

	// RunNowAndWait calls RunNow() and waits to reach TERMINATED or SKIPPED state
	//
	// This method is generated by Databricks SDK Code Generator.
	RunNowAndWait(ctx context.Context, request RunNow, options ...retries.Option[Run]) (*Run, error)

	// Create and trigger a one-time run
	//
	// Submit a one-time run. This endpoint allows you to submit a workload
	// directly without creating a job. Runs submitted using this endpoint don?t
	// display in the UI. Use the `jobs/runs/get` API to check the run state
	// after the job is submitted.
	Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error)

	// SubmitAndWait calls Submit() and waits to reach TERMINATED or SKIPPED state
	//
	// This method is generated by Databricks SDK Code Generator.
	SubmitAndWait(ctx context.Context, request SubmitRun, options ...retries.Option[Run]) (*Run, error)

	// Partially updates a job
	//
	// Add, update, or remove specific settings of an existing job. Use the
	// ResetJob to overwrite all job settings.
	Update(ctx context.Context, request UpdateJob) error
}
