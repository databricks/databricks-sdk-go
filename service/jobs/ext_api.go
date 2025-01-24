package jobs

import "context"

// GetRun retrieves a run based on the provided request.
// It handles pagination if the run contains multiple iterations or tasks.
func (a *JobsAPI) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	run, err := a.jobsImpl.GetRun(ctx, request)
	if err != nil {
		return nil, err
	}

	// When querying a Job run, a page token is returned when there are more than 100 tasks. No iterations are defined for a Job run. Therefore, the next page in the response only includes the next page of tasks.
	// When querying a ForEach task run, a page token is returned when there are more than 100 iterations. Only a single task is returned, corresponding to the ForEach task itself. Therefore, the client only reads the iterations from the next page and not the tasks.
	isPaginatingIterations := len(run.Iterations) > 0

	pageToken := run.NextPageToken
	for pageToken != "" {
		request.PageToken = pageToken
		nextRun, err := a.jobsImpl.GetRun(ctx, request)
		if err != nil {
			return nil, err
		}

		if isPaginatingIterations {
			run.Iterations = append(run.Iterations, nextRun.Iterations...)
		} else {
			run.Tasks = append(run.Tasks, nextRun.Tasks...)
		}
		pageToken = nextRun.NextPageToken
	}

	return run, nil
}

// Get retrieves a job based on the provided request.
// It handles pagination if the job contains multiple tasks, job_clusters, job_parameters or environments.
func (a *JobsAPI) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	job, err := a.jobsImpl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	for job.NextPageToken != "" {
		request.PageToken = job.NextPageToken
		nextJob, err := a.jobsImpl.Get(ctx, request)
		if err != nil {
			return nil, err
		}

		job.Settings.Tasks = append(job.Settings.Tasks, nextJob.Settings.Tasks...)
		job.Settings.JobClusters = append(job.Settings.JobClusters, nextJob.Settings.JobClusters...)
		job.Settings.Parameters = append(job.Settings.Parameters, nextJob.Settings.Parameters...)
		job.Settings.Environments = append(job.Settings.Environments, nextJob.Settings.Environments...)
		job.NextPageToken = nextJob.NextPageToken
	}

	return job, nil
}
