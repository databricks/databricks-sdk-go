package jobs

import "context"

func (a *JobsAPI) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	run, err := a.jobsImpl.GetRun(ctx, request)
	if err != nil {
		return nil, err
	}

	isPaginatingIterations := run.Iterations != nil && len(run.Iterations) > 0

	for len(run.NextPageToken) != 0 {
		request.PageToken = run.NextPageToken
		nextRun, err := a.jobsImpl.GetRun(ctx, request)
		if err != nil {
			return nil, err
		}

		if isPaginatingIterations {
			run.Iterations = append(run.Iterations, nextRun.Iterations...)
		} else {
			run.Tasks = append(run.Tasks, nextRun.Tasks...)
		}
		run.NextPageToken = nextRun.NextPageToken
	}

	run.PrevPageToken = ""
	return run, nil
}
