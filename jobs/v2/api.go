// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Jobs, Policy Compliance For Jobs, etc.
package jobs

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type jobsBaseClient struct {
	jobsImpl
}

// Cancel a run.
//
// Cancels a job run or a task run. The run is canceled asynchronously, so it
// may still be running when this request completes.
func (a *jobsBaseClient) CancelRunByRunId(ctx context.Context, runId int64) (*CancelRunResponse, error) {
	return a.jobsImpl.CancelRun(ctx, CancelRun{
		RunId: runId,
	})
}

// Delete a job.
//
// Deletes a job.
func (a *jobsBaseClient) DeleteByJobId(ctx context.Context, jobId int64) (*DeleteResponse, error) {
	return a.jobsImpl.Delete(ctx, DeleteJob{
		JobId: jobId,
	})
}

// Delete a job run.
//
// Deletes a non-active run. Returns an error if the run is active.
func (a *jobsBaseClient) DeleteRunByRunId(ctx context.Context, runId int64) (*DeleteRunResponse, error) {
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
func (a *jobsBaseClient) GetByJobId(ctx context.Context, jobId int64) (*Job, error) {
	return a.jobsImpl.Get(ctx, GetJobRequest{
		JobId: jobId,
	})
}

// Get job permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *jobsBaseClient) GetPermissionLevelsByJobId(ctx context.Context, jobId string) (*GetJobPermissionLevelsResponse, error) {
	return a.jobsImpl.GetPermissionLevels(ctx, GetJobPermissionLevelsRequest{
		JobId: jobId,
	})
}

// Get job permissions.
//
// Gets the permissions of a job. Jobs can inherit permissions from their root
// object.
func (a *jobsBaseClient) GetPermissionsByJobId(ctx context.Context, jobId string) (*JobPermissions, error) {
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
func (a *jobsBaseClient) GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error) {
	return a.jobsImpl.GetRunOutput(ctx, GetRunOutputRequest{
		RunId: runId,
	})
}

// BaseJobSettingsNameToJobIdMap calls [jobsBaseClient.ListAll] and creates a map of results with [BaseJob].Settings.Name as key and [BaseJob].JobId as value.
//
// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
//
// Note: All [BaseJob] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *jobsBaseClient) BaseJobSettingsNameToJobIdMap(ctx context.Context, request ListJobsRequest) (map[string]int64, error) {
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

// GetBySettingsName calls [jobsBaseClient.BaseJobSettingsNameToJobIdMap] and returns a single [BaseJob].
//
// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
//
// Note: All [BaseJob] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *jobsBaseClient) GetBySettingsName(ctx context.Context, name string) (*BaseJob, error) {
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

type policyComplianceForJobsBaseClient struct {
	policyComplianceForJobsImpl
}

// Get job policy compliance.
//
// Returns the policy compliance status of a job. Jobs could be out of
// compliance if a cluster policy they use was updated after the job was last
// edited and some of its job clusters no longer comply with their updated
// policies.
func (a *policyComplianceForJobsBaseClient) GetComplianceByJobId(ctx context.Context, jobId int64) (*GetPolicyComplianceResponse, error) {
	return a.policyComplianceForJobsImpl.GetCompliance(ctx, GetPolicyComplianceRequest{
		JobId: jobId,
	})
}
