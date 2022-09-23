// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"github.com/databricks/databricks-sdk-go/retries"
)

func NewJobs(client *client.DatabricksClient) JobsService {
	return &JobsAPI{
		client: client,
	}
}

type JobsAPI struct {
	client *client.DatabricksClient
}

// Cancel all runs of a job
//
// Cancels all active runs of a job. The runs are canceled asynchronously, so it
// doesn't prevent new runs from being started.
func (a *JobsAPI) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	path := "/api/2.1/jobs/runs/cancel-all"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Cancel all runs of a job
//
// Cancels all active runs of a job. The runs are canceled asynchronously, so it
// doesn't prevent new runs from being started.
func (a *JobsAPI) CancelAllRunsByJobId(ctx context.Context, jobId int64) error {
	return a.CancelAllRuns(ctx, CancelAllRuns{
		JobId: jobId,
	})
}

// Cancel a job run
//
// Cancels a job run. The run is canceled asynchronously, so it may still be
// running when this request completes.
func (a *JobsAPI) CancelRun(ctx context.Context, request CancelRun) error {
	path := "/api/2.1/jobs/runs/cancel"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// CancelRun and wait to reach TERMINATED or SKIPPED state
func (a *JobsAPI) CancelRunAndWait(ctx context.Context, cancelRun CancelRun, timeout ...time.Duration) (*Run, error) {
	err := a.CancelRun(ctx, cancelRun)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[Run](ctx, timeout[0], func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: cancelRun.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := run.State.LifeCycleState
		statusMessage := run.State.StateMessage
		switch status {
		case RunLifeCycleStateTerminated, RunLifeCycleStateSkipped: // target state
			return run, nil
		case RunLifeCycleStateInternalError:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				RunLifeCycleStateTerminated, RunLifeCycleStateSkipped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Cancel a job run
//
// Cancels a job run. The run is canceled asynchronously, so it may still be
// running when this request completes.
func (a *JobsAPI) CancelRunByRunId(ctx context.Context, runId int64) error {
	return a.CancelRun(ctx, CancelRun{
		RunId: runId,
	})
}

func (a *JobsAPI) CancelRunByRunIdAndWait(ctx context.Context, runId int64, timeout ...time.Duration) (*Run, error) {
	return a.CancelRunAndWait(ctx, CancelRun{
		RunId: runId,
	}, timeout...)
}

// Create a new job
//
// Create a new job.
func (a *JobsAPI) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.1/jobs/create"
	err := a.client.Post(ctx, path, request, &createResponse)
	return &createResponse, err
}

// Delete a job
//
// Deletes a job.
func (a *JobsAPI) Delete(ctx context.Context, request DeleteJob) error {
	path := "/api/2.1/jobs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Delete a job
//
// Deletes a job.
func (a *JobsAPI) DeleteByJobId(ctx context.Context, jobId int64) error {
	return a.Delete(ctx, DeleteJob{
		JobId: jobId,
	})
}

// Delete a job run
//
// Deletes a non-active run. Returns an error if the run is active.
func (a *JobsAPI) DeleteRun(ctx context.Context, request DeleteRun) error {
	path := "/api/2.1/jobs/runs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Delete a job run
//
// Deletes a non-active run. Returns an error if the run is active.
func (a *JobsAPI) DeleteRunByRunId(ctx context.Context, runId int64) error {
	return a.DeleteRun(ctx, DeleteRun{
		RunId: runId,
	})
}

// Export and retrieve a job run
//
// Export and retrieve the job run task.
func (a *JobsAPI) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	var exportRunOutput ExportRunOutput
	path := "/api/2.1/jobs/runs/export"
	err := a.client.Get(ctx, path, request, &exportRunOutput)
	return &exportRunOutput, err
}

// Get a single job
//
// Retrieves the details for a single job.
func (a *JobsAPI) Get(ctx context.Context, request GetRequest) (*Job, error) {
	var job Job
	path := "/api/2.1/jobs/get"
	err := a.client.Get(ctx, path, request, &job)
	return &job, err
}

// Get a single job
//
// Retrieves the details for a single job.
func (a *JobsAPI) GetByJobId(ctx context.Context, jobId int64) (*Job, error) {
	return a.Get(ctx, GetRequest{
		JobId: jobId,
	})
}

// Get a single job run
//
// Retrieve the metadata of a run.
func (a *JobsAPI) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	var run Run
	path := "/api/2.1/jobs/runs/get"
	err := a.client.Get(ctx, path, request, &run)
	return &run, err
}

// GetRun and wait to reach TERMINATED or SKIPPED state
func (a *JobsAPI) GetRunAndWait(ctx context.Context, getRunRequest GetRunRequest, timeout ...time.Duration) (*Run, error) {
	run, err := a.GetRun(ctx, getRunRequest)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[Run](ctx, timeout[0], func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: run.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := run.State.LifeCycleState
		statusMessage := run.State.StateMessage
		switch status {
		case RunLifeCycleStateTerminated, RunLifeCycleStateSkipped: // target state
			return run, nil
		case RunLifeCycleStateInternalError:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				RunLifeCycleStateTerminated, RunLifeCycleStateSkipped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get the output for a single run
//
// Retrieve the output and metadata of a single task run. When a notebook task
// returns a value through the `dbutils.notebook.exit()` call, you can use this
// endpoint to retrieve that value. " + serviceName + " restricts this API to
// returning the first 5 MB of the output. To return a larger result, you can
// store job results in a cloud storage service.
//
// This endpoint validates that the __run_id__ parameter is valid and returns an
// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
// automatically removed after 60 days. If you to want to reference them beyond
// 60 days, you must save old run results before they expire.
func (a *JobsAPI) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	var runOutput RunOutput
	path := "/api/2.1/jobs/runs/get-output"
	err := a.client.Get(ctx, path, request, &runOutput)
	return &runOutput, err
}

// Get the output for a single run
//
// Retrieve the output and metadata of a single task run. When a notebook task
// returns a value through the `dbutils.notebook.exit()` call, you can use this
// endpoint to retrieve that value. " + serviceName + " restricts this API to
// returning the first 5 MB of the output. To return a larger result, you can
// store job results in a cloud storage service.
//
// This endpoint validates that the __run_id__ parameter is valid and returns an
// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
// automatically removed after 60 days. If you to want to reference them beyond
// 60 days, you must save old run results before they expire.
func (a *JobsAPI) GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error) {
	return a.GetRunOutput(ctx, GetRunOutputRequest{
		RunId: runId,
	})
}

// List all jobs
//
// Retrieves a list of jobs.
//
// Use ListAll to get all Job instances, which will iterate over every result page.
func (a *JobsAPI) List(ctx context.Context, request ListRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.1/jobs/list"
	err := a.client.Get(ctx, path, request, &listResponse)
	return &listResponse, err
}

// ListAll returns all Job instances by calling List for every result page
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListAll(ctx context.Context, request ListRequest) ([]Job, error) {
	var results []Job
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Jobs) == 0 {
			break
		}
		for _, v := range response.Jobs {
			results = append(results, v)
		}
		request.Offset += int(len(response.Jobs)) // TODO: check for duplicates
	}
	return results, nil
}

// List runs for a job
//
// List runs in descending order by start time.
//
// Use ListRunsAll to get all Run instances, which will iterate over every result page.
func (a *JobsAPI) ListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	var listRunsResponse ListRunsResponse
	path := "/api/2.1/jobs/runs/list"
	err := a.client.Get(ctx, path, request, &listRunsResponse)
	return &listRunsResponse, err
}

// ListRunsAll returns all Run instances by calling ListRuns for every result page
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListRunsAll(ctx context.Context, request ListRunsRequest) ([]Run, error) {
	var results []Run
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.ListRuns(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Runs) == 0 {
			break
		}
		for _, v := range response.Runs {
			results = append(results, v)
		}
		request.Offset += int(len(response.Runs)) // TODO: check for duplicates
	}
	return results, nil
}

// Repair a job run
//
// Re-run one or more tasks. Tasks are re-run as part of the original job run.
// They use the current job and task settings, and can be viewed in the history
// for the original job run.
func (a *JobsAPI) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	var repairRunResponse RepairRunResponse
	path := "/api/2.1/jobs/runs/repair"
	err := a.client.Post(ctx, path, request, &repairRunResponse)
	return &repairRunResponse, err
}

// RepairRun and wait to reach TERMINATED or SKIPPED state
func (a *JobsAPI) RepairRunAndWait(ctx context.Context, repairRun RepairRun, timeout ...time.Duration) (*Run, error) {
	_, err := a.RepairRun(ctx, repairRun)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[Run](ctx, timeout[0], func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: repairRun.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := run.State.LifeCycleState
		statusMessage := run.State.StateMessage
		switch status {
		case RunLifeCycleStateTerminated, RunLifeCycleStateSkipped: // target state
			return run, nil
		case RunLifeCycleStateInternalError:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				RunLifeCycleStateTerminated, RunLifeCycleStateSkipped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Overwrites all settings for a job
//
// Overwrites all the settings for a specific job. Use the Update endpoint to
// update job settings partially.
func (a *JobsAPI) Reset(ctx context.Context, request ResetJob) error {
	path := "/api/2.1/jobs/reset"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Trigger a new job run
//
// Run a job and return the `run_id` of the triggered run.
func (a *JobsAPI) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	var runNowResponse RunNowResponse
	path := "/api/2.1/jobs/run-now"
	err := a.client.Post(ctx, path, request, &runNowResponse)
	return &runNowResponse, err
}

// RunNow and wait to reach TERMINATED or SKIPPED state
func (a *JobsAPI) RunNowAndWait(ctx context.Context, runNow RunNow, timeout ...time.Duration) (*Run, error) {
	runNowResponse, err := a.RunNow(ctx, runNow)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[Run](ctx, timeout[0], func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: runNowResponse.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := run.State.LifeCycleState
		statusMessage := run.State.StateMessage
		switch status {
		case RunLifeCycleStateTerminated, RunLifeCycleStateSkipped: // target state
			return run, nil
		case RunLifeCycleStateInternalError:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				RunLifeCycleStateTerminated, RunLifeCycleStateSkipped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Create and trigger a one-time run
//
// Submit a one-time run. This endpoint allows you to submit a workload directly
// without creating a job. Runs submitted using this endpoint don?t display in
// the UI. Use the `jobs/runs/get` API to check the run state after the job is
// submitted.
func (a *JobsAPI) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	var submitRunResponse SubmitRunResponse
	path := "/api/2.1/jobs/runs/submit"
	err := a.client.Post(ctx, path, request, &submitRunResponse)
	return &submitRunResponse, err
}

// Submit and wait to reach TERMINATED or SKIPPED state
func (a *JobsAPI) SubmitAndWait(ctx context.Context, submitRun SubmitRun, timeout ...time.Duration) (*Run, error) {
	submitRunResponse, err := a.Submit(ctx, submitRun)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[Run](ctx, timeout[0], func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: submitRunResponse.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := run.State.LifeCycleState
		statusMessage := run.State.StateMessage
		switch status {
		case RunLifeCycleStateTerminated, RunLifeCycleStateSkipped: // target state
			return run, nil
		case RunLifeCycleStateInternalError:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				RunLifeCycleStateTerminated, RunLifeCycleStateSkipped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Partially updates a job
//
// Add, update, or remove specific settings of an existing job. Use the ResetJob
// to overwrite all job settings.
func (a *JobsAPI) Update(ctx context.Context, request UpdateJob) error {
	path := "/api/2.1/jobs/update"
	err := a.client.Post(ctx, path, request, nil)
	return err
}
