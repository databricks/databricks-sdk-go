// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"context"
)

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
type JobsService interface {

	// Cancel all runs of a job.
	//
	// Cancels all active runs of a job. The runs are canceled asynchronously,
	// so it doesn't prevent new runs from being started.
	CancelAllRuns(ctx context.Context, request CancelAllRuns) error

	// Cancel a run.
	//
	// Cancels a job run or a task run. The run is canceled asynchronously, so
	// it may still be running when this request completes.
	CancelRun(ctx context.Context, request CancelRun) error

	// Create a new job.
	//
	// Create a new job.
	Create(ctx context.Context, request CreateJob) (*CreateResponse, error)

	// Delete a job.
	//
	// Deletes a job.
	Delete(ctx context.Context, request DeleteJob) error

	// Delete a job run.
	//
	// Deletes a non-active run. Returns an error if the run is active.
	DeleteRun(ctx context.Context, request DeleteRun) error

	// Export and retrieve a job run.
	//
	// Export and retrieve the job run task.
	ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error)

	// Get a single job.
	//
	// Retrieves the details for a single job.
	//
	// In Jobs API 2.2, requests for a single job support pagination of `tasks`
	// and `job_clusters` when either exceeds 100 elements. Use the
	// `next_page_token` field to check for more results and pass its value as
	// the `page_token` in subsequent requests. Arrays with fewer than 100
	// elements in a page will be empty on later pages.
	Get(ctx context.Context, request GetJobRequest) (*Job, error)

	// Get job permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error)

	// Get job permissions.
	//
	// Gets the permissions of a job. Jobs can inherit permissions from their
	// root object.
	GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error)

	// Get a single job run.
	//
	// Retrieves the metadata of a run.
	//
	// In Jobs API 2.2, requests for a single job run support pagination of
	// `tasks` and `job_clusters` when either exceeds 100 elements. Use the
	// `next_page_token` field to check for more results and pass its value as
	// the `page_token` in subsequent requests. Arrays with fewer than 100
	// elements in a page will be empty on later pages.
	GetRun(ctx context.Context, request GetRunRequest) (*Run, error)

	// Get the output for a single run.
	//
	// Retrieve the output and metadata of a single task run. When a notebook
	// task returns a value through the `dbutils.notebook.exit()` call, you can
	// use this endpoint to retrieve that value. Databricks restricts this API
	// to returning the first 5 MB of the output. To return a larger result, you
	// can store job results in a cloud storage service.
	//
	// This endpoint validates that the __run_id__ parameter is valid and
	// returns an HTTP status code 400 if the __run_id__ parameter is invalid.
	// Runs are automatically removed after 60 days. If you to want to reference
	// them beyond 60 days, you must save old run results before they expire.
	GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error)

	// List jobs.
	//
	// Retrieves a list of jobs.
	//
	// Use ListAll() to get all BaseJob instances, which will iterate over every result page.
	List(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error)

	// List job runs.
	//
	// List runs in descending order by start time.
	//
	// Use ListRunsAll() to get all BaseRun instances, which will iterate over every result page.
	ListRuns(ctx context.Context, request ListRunsRequest) (*ListRunsResponse, error)

	// Repair a job run.
	//
	// Re-run one or more tasks. Tasks are re-run as part of the original job
	// run. They use the current job and task settings, and can be viewed in the
	// history for the original job run.
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
	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error)

	// Create and trigger a one-time run.
	//
	// Submit a one-time run. This endpoint allows you to submit a workload
	// directly without creating a job. Runs submitted using this endpoint
	// don’t display in the UI. Use the `jobs/runs/get` API to check the run
	// state after the job is submitted.
	Submit(ctx context.Context, request SubmitRun) (*SubmitRunResponse, error)

	// Update job settings partially.
	//
	// Add, update, or remove specific settings of an existing job. Use the
	// [_Reset_ endpoint](:method:jobs/reset) to overwrite all job settings.
	Update(ctx context.Context, request UpdateJob) error

	// Update job permissions.
	//
	// Updates the permissions on a job. Jobs can inherit permissions from their
	// root object.
	UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error)
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
type PolicyComplianceForJobsService interface {

	// Enforce job policy compliance.
	//
	// Updates a job so the job clusters that are created when running the job
	// (specified in `new_cluster`) are compliant with the current versions of
	// their respective cluster policies. All-purpose clusters used in the job
	// will not be updated.
	EnforceCompliance(ctx context.Context, request EnforcePolicyComplianceRequest) (*EnforcePolicyComplianceResponse, error)

	// Get job policy compliance.
	//
	// Returns the policy compliance status of a job. Jobs could be out of
	// compliance if a cluster policy they use was updated after the job was
	// last edited and some of its job clusters no longer comply with their
	// updated policies.
	GetCompliance(ctx context.Context, request GetPolicyComplianceRequest) (*GetPolicyComplianceResponse, error)

	// List job policy compliance.
	//
	// Returns the policy compliance status of all jobs that use a given policy.
	// Jobs could be out of compliance if a cluster policy they use was updated
	// after the job was last edited and its job clusters no longer comply with
	// the updated policy.
	//
	// Use ListComplianceAll() to get all JobCompliance instances, which will iterate over every result page.
	ListCompliance(ctx context.Context, request ListJobComplianceRequest) (*ListJobComplianceForPolicyResponse, error)
}
