// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Jobs API allows you to create, edit, and delete jobs.
package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewJobs(client *client.DatabricksClient) *JobsAPI {
	return &JobsAPI{
		impl: &jobsImpl{
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
// :service:secrets to manage secrets in the [Databricks CLI]. Use the [Secrets
// utility] to reference secrets in notebooks and jobs.
//
// [Databricks CLI]: https://docs.databricks.com/dev-tools/cli/index.html
// [Secrets utility]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-secrets
type JobsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(JobsService)
	impl JobsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *JobsAPI) WithImpl(impl JobsService) *JobsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Jobs API implementation
func (a *JobsAPI) Impl() JobsService {
	return a.impl
}

// Cancel all runs of a job.
//
// Cancels all active runs of a job. The runs are canceled asynchronously, so it
// doesn't prevent new runs from being started.
func (a *JobsAPI) CancelAllRuns(ctx context.Context, request CancelAllRuns) error {
	return a.impl.CancelAllRuns(ctx, request)
}

// Cancel all runs of a job.
//
// Cancels all active runs of a job. The runs are canceled asynchronously, so it
// doesn't prevent new runs from being started.
func (a *JobsAPI) CancelAllRunsByJobId(ctx context.Context, jobId int64) error {
	return a.impl.CancelAllRuns(ctx, CancelAllRuns{
		JobId: jobId,
	})
}

// Cancel a job run.
//
// Cancels a job run. The run is canceled asynchronously, so it may still be
// running when this request completes.
func (a *JobsAPI) CancelRun(ctx context.Context, request CancelRun) error {
	return a.impl.CancelRun(ctx, request)
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
				Info:    run,
				Timeout: i.Timeout,
			})
		}
		status := run.State.LifeCycleState
		statusMessage := fmt.Sprintf("current status: %s", status)
		if run.State != nil {
			statusMessage = run.State.StateMessage
		}
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

// Cancel a job run.
//
// Cancels a job run. The run is canceled asynchronously, so it may still be
// running when this request completes.
func (a *JobsAPI) CancelRunByRunId(ctx context.Context, runId int64) error {
	return a.impl.CancelRun(ctx, CancelRun{
		RunId: runId,
	})
}

func (a *JobsAPI) CancelRunByRunIdAndWait(ctx context.Context, runId int64, options ...retries.Option[Run]) (*Run, error) {
	return a.CancelRunAndWait(ctx, CancelRun{
		RunId: runId,
	}, options...)
}

// Create a new job.
//
// Create a new job.
func (a *JobsAPI) Create(ctx context.Context, request CreateJob) (*CreateResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a job.
//
// Deletes a job.
func (a *JobsAPI) Delete(ctx context.Context, request DeleteJob) error {
	return a.impl.Delete(ctx, request)
}

// Delete a job.
//
// Deletes a job.
func (a *JobsAPI) DeleteByJobId(ctx context.Context, jobId int64) error {
	return a.impl.Delete(ctx, DeleteJob{
		JobId: jobId,
	})
}

// Delete a job run.
//
// Deletes a non-active run. Returns an error if the run is active.
func (a *JobsAPI) DeleteRun(ctx context.Context, request DeleteRun) error {
	return a.impl.DeleteRun(ctx, request)
}

// Delete a job run.
//
// Deletes a non-active run. Returns an error if the run is active.
func (a *JobsAPI) DeleteRunByRunId(ctx context.Context, runId int64) error {
	return a.impl.DeleteRun(ctx, DeleteRun{
		RunId: runId,
	})
}

// Export and retrieve a job run.
//
// Export and retrieve the job run task.
func (a *JobsAPI) ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error) {
	return a.impl.ExportRun(ctx, request)
}

// Get a single job.
//
// Retrieves the details for a single job.
func (a *JobsAPI) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	return a.impl.Get(ctx, request)
}

// Get a single job.
//
// Retrieves the details for a single job.
func (a *JobsAPI) GetByJobId(ctx context.Context, jobId int64) (*Job, error) {
	return a.impl.Get(ctx, GetJobRequest{
		JobId: jobId,
	})
}

// Get a single job run.
//
// Retrieve the metadata of a run.
func (a *JobsAPI) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	return a.impl.GetRun(ctx, request)
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
				Info:    run,
				Timeout: i.Timeout,
			})
		}
		status := run.State.LifeCycleState
		statusMessage := fmt.Sprintf("current status: %s", status)
		if run.State != nil {
			statusMessage = run.State.StateMessage
		}
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
func (a *JobsAPI) GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error) {
	return a.impl.GetRunOutput(ctx, request)
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
	return a.impl.GetRunOutput(ctx, GetRunOutputRequest{
		RunId: runId,
	})
}

// List all jobs.
//
// Retrieves a list of jobs.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error) {
	var results []BaseJob
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	// deduplicate items that may have been added during iteration
	seen := map[int64]bool{}
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Jobs) == 0 {
			break
		}
		for _, v := range response.Jobs {
			id := v.JobId
			if seen[id] {
				// item was added during iteration
				continue
			}
			seen[id] = true
			results = append(results, v)
		}
		request.Offset += int(len(response.Jobs))
	}
	return results, nil
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

// List runs for a job.
//
// List runs in descending order by start time.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error) {
	var results []BaseRun
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	// deduplicate items that may have been added during iteration
	seen := map[int64]bool{}
	for {
		response, err := a.impl.ListRuns(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Runs) == 0 {
			break
		}
		for _, v := range response.Runs {
			id := v.RunId
			if seen[id] {
				// item was added during iteration
				continue
			}
			seen[id] = true
			results = append(results, v)
		}
		request.Offset += int(len(response.Runs))
	}
	return results, nil
}

// Repair a job run.
//
// Re-run one or more tasks. Tasks are re-run as part of the original job run.
// They use the current job and task settings, and can be viewed in the history
// for the original job run.
func (a *JobsAPI) RepairRun(ctx context.Context, request RepairRun) (*RepairRunResponse, error) {
	return a.impl.RepairRun(ctx, request)
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
				Info:    run,
				Timeout: i.Timeout,
			})
		}
		status := run.State.LifeCycleState
		statusMessage := fmt.Sprintf("current status: %s", status)
		if run.State != nil {
			statusMessage = run.State.StateMessage
		}
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

// Overwrites all settings for a job.
//
// Overwrites all the settings for a specific job. Use the Update endpoint to
// update job settings partially.
func (a *JobsAPI) Reset(ctx context.Context, request ResetJob) error {
	return a.impl.Reset(ctx, request)
}

// Trigger a new job run.
//
// Run a job and return the `run_id` of the triggered run.
func (a *JobsAPI) RunNow(ctx context.Context, request RunNow) (*RunNowResponse, error) {
	return a.impl.RunNow(ctx, request)
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
				Info:    run,
				Timeout: i.Timeout,
			})
		}
		status := run.State.LifeCycleState
		statusMessage := fmt.Sprintf("current status: %s", status)
		if run.State != nil {
			statusMessage = run.State.StateMessage
		}
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

// Create and trigger a one-time run.
//
// Submit a one-time run. This endpoint allows you to submit a workload directly
// without creating a job. Runs submitted using this endpoint don’t display in
// the UI. Use the `jobs/runs/get` API to check the run state after the job is
// submitted.
func (a *JobsAPI) Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error) {
	return a.impl.Submit(ctx, request)
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
				Info:    run,
				Timeout: i.Timeout,
			})
		}
		status := run.State.LifeCycleState
		statusMessage := fmt.Sprintf("current status: %s", status)
		if run.State != nil {
			statusMessage = run.State.StateMessage
		}
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

// Partially updates a job.
//
// Add, update, or remove specific settings of an existing job. Use the ResetJob
// to overwrite all job settings.
func (a *JobsAPI) Update(ctx context.Context, request UpdateJob) error {
	return a.impl.Update(ctx, request)
}
