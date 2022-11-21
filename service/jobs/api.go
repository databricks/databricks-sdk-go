// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

func NewJobs(client *client.DatabricksClient) *JobsAPI {
	return &JobsAPI{
		JobsService: &jobsAPI{
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
type JobsAPI struct {
	// JobsService contains low-level REST API interface.
	JobsService
}

// Cancel all runs of a job
//
// Cancels all active runs of a job. The runs are canceled asynchronously, so it
// doesn't prevent new runs from being started.
func (a *JobsAPI) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	return a.JobsService.CancelAllRuns(ctx, request)
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
	return a.JobsService.CancelRun(ctx, request)
}

// Calls [JobsAPI.CancelRun] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
func (a *JobsAPI) CancelRunAndWait(ctx context.Context, cancelRun CancelRun, options ...retries.Option[Run]) (*Run, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.CancelRun(ctx, cancelRun)
	if err != nil {
		return nil, err
	}
	i := retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[Run](ctx, i.Timeout, func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: cancelRun.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    *run,
				Timeout: i.Timeout,
			})
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

func (a *JobsAPI) CancelRunByRunIdAndWait(ctx context.Context, runId int64, options ...retries.Option[Run]) (*Run, error) {
	return a.CancelRunAndWait(ctx, CancelRun{
		RunId: runId,
	}, options...)
}

// Create a new job
//
// Create a new job.
func (a *JobsAPI) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	return a.JobsService.Create(ctx, request)
}

// Delete a job
//
// Deletes a job.
func (a *JobsAPI) Delete(ctx context.Context, request DeleteJob) error {
	return a.JobsService.Delete(ctx, request)
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
	return a.JobsService.DeleteRun(ctx, request)
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
	return a.JobsService.ExportRun(ctx, request)
}

// Get a single job
//
// Retrieves the details for a single job.
func (a *JobsAPI) Get(ctx context.Context, request GetRequest) (*Job, error) {
	return a.JobsService.Get(ctx, request)
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
	return a.JobsService.GetRun(ctx, request)
}

// Calls [JobsAPI.GetRun] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
func (a *JobsAPI) GetRunAndWait(ctx context.Context, getRunRequest GetRunRequest, options ...retries.Option[Run]) (*Run, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	run, err := a.GetRun(ctx, getRunRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[Run](ctx, i.Timeout, func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: run.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    *run,
				Timeout: i.Timeout,
			})
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
	return a.JobsService.GetRunOutput(ctx, request)
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
		request.Offset += int(len(response.Jobs))
	}
	return results, nil
}

// List runs for a job
//
// List runs in descending order by start time.
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
		request.Offset += int(len(response.Runs))
	}
	return results, nil
}

// Repair a job run
//
// Re-run one or more tasks. Tasks are re-run as part of the original job run.
// They use the current job and task settings, and can be viewed in the history
// for the original job run.
func (a *JobsAPI) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	return a.JobsService.RepairRun(ctx, request)
}

// Calls [JobsAPI.RepairRun] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
func (a *JobsAPI) RepairRunAndWait(ctx context.Context, repairRun RepairRun, options ...retries.Option[Run]) (*Run, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	_, err := a.RepairRun(ctx, repairRun)
	if err != nil {
		return nil, err
	}
	i := retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[Run](ctx, i.Timeout, func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: repairRun.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    *run,
				Timeout: i.Timeout,
			})
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
	return a.JobsService.Reset(ctx, request)
}

// Trigger a new job run
//
// Run a job and return the `run_id` of the triggered run.
func (a *JobsAPI) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	return a.JobsService.RunNow(ctx, request)
}

// Calls [JobsAPI.RunNow] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
func (a *JobsAPI) RunNowAndWait(ctx context.Context, runNow RunNow, options ...retries.Option[Run]) (*Run, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	runNowResponse, err := a.RunNow(ctx, runNow)
	if err != nil {
		return nil, err
	}
	i := retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[Run](ctx, i.Timeout, func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: runNowResponse.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    *run,
				Timeout: i.Timeout,
			})
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
// without creating a job. Runs submitted using this endpoint don’t display in
// the UI. Use the `jobs/runs/get` API to check the run state after the job is
// submitted.
func (a *JobsAPI) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	return a.JobsService.Submit(ctx, request)
}

// Calls [JobsAPI.Submit] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
func (a *JobsAPI) SubmitAndWait(ctx context.Context, submitRun SubmitRun, options ...retries.Option[Run]) (*Run, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	submitRunResponse, err := a.Submit(ctx, submitRun)
	if err != nil {
		return nil, err
	}
	i := retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[Run](ctx, i.Timeout, func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: submitRunResponse.RunId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    *run,
				Timeout: i.Timeout,
			})
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
	return a.JobsService.Update(ctx, request)
}

// unexported type that holds implementations of just Jobs API methods
type jobsAPI struct {
	client *client.DatabricksClient
}

func (a *jobsAPI) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	path := "/api/2.1/jobs/runs/cancel-all"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *jobsAPI) CancelRun(ctx context.Context, request CancelRun) error {
	path := "/api/2.1/jobs/runs/cancel"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *jobsAPI) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.1/jobs/create"
	err := a.client.Post(ctx, path, request, &createResponse)
	return &createResponse, err
}

func (a *jobsAPI) Delete(ctx context.Context, request DeleteJob) error {
	path := "/api/2.1/jobs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *jobsAPI) DeleteRun(ctx context.Context, request DeleteRun) error {
	path := "/api/2.1/jobs/runs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *jobsAPI) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	var exportRunOutput ExportRunOutput
	path := "/api/2.1/jobs/runs/export"
	err := a.client.Get(ctx, path, request, &exportRunOutput)
	return &exportRunOutput, err
}

func (a *jobsAPI) Get(ctx context.Context, request GetRequest) (*Job, error) {
	var job Job
	path := "/api/2.1/jobs/get"
	err := a.client.Get(ctx, path, request, &job)
	return &job, err
}

func (a *jobsAPI) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	var run Run
	path := "/api/2.1/jobs/runs/get"
	err := a.client.Get(ctx, path, request, &run)
	return &run, err
}

func (a *jobsAPI) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	var runOutput RunOutput
	path := "/api/2.1/jobs/runs/get-output"
	err := a.client.Get(ctx, path, request, &runOutput)
	return &runOutput, err
}

func (a *jobsAPI) List(ctx context.Context, request ListRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.1/jobs/list"
	err := a.client.Get(ctx, path, request, &listResponse)
	return &listResponse, err
}

func (a *jobsAPI) ListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error) {
	var listRunsResponse ListRunsResponse
	path := "/api/2.1/jobs/runs/list"
	err := a.client.Get(ctx, path, request, &listRunsResponse)
	return &listRunsResponse, err
}

func (a *jobsAPI) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	var repairRunResponse RepairRunResponse
	path := "/api/2.1/jobs/runs/repair"
	err := a.client.Post(ctx, path, request, &repairRunResponse)
	return &repairRunResponse, err
}

func (a *jobsAPI) Reset(ctx context.Context, request ResetJob) error {
	path := "/api/2.1/jobs/reset"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *jobsAPI) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	var runNowResponse RunNowResponse
	path := "/api/2.1/jobs/run-now"
	err := a.client.Post(ctx, path, request, &runNowResponse)
	return &runNowResponse, err
}

func (a *jobsAPI) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	var submitRunResponse SubmitRunResponse
	path := "/api/2.1/jobs/runs/submit"
	err := a.client.Post(ctx, path, request, &submitRunResponse)
	return &submitRunResponse, err
}

func (a *jobsAPI) Update(ctx context.Context, request UpdateJob) error {
	path := "/api/2.1/jobs/update"
	err := a.client.Post(ctx, path, request, nil)
	return err
}
