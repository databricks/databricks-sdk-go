package jobs

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"

	"github.com/stretchr/testify/assert"
)

func TestGetRun(t *testing.T) {
	ctx := context.Background()

	t.Run("run with no pagination", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/get?run_id=514594995218126",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId:   123,
							TaskKey: "task1",
						},
						{
							RunId:   1234,
							TaskKey: "task2",
						},
					},
					NextPageToken: "",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/runs/get?run_id=514594995218126",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId:   123,
							TaskKey: "task1",
						},
						{
							RunId:   1234,
							TaskKey: "task2",
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetRunRequest{RunId: 514594995218126}
		run, err := api.GetRun(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(run.Tasks))
		assert.Empty(t, run.Iterations)
		assert.EqualValues(t, 123, run.Tasks[0].RunId)
		assert.EqualValues(t, 1234, run.Tasks[1].RunId)
	})

	t.Run("run with two tasks pages", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?page_token=token1&run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/runs/get?run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetRunRequest{RunId: 111222333}
		run, err := api.GetRun(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 4, len(run.Tasks))
		assert.Empty(t, run.Iterations)
		assert.Empty(t, run.NextPageToken)
		expected := []RunTask{
			{RunId: 123, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 1234, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 222, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 333, ForceSendFields: []string{"RunId", "TaskKey"}},
		}
		assert.Equal(t, expected, run.Tasks)
	})

	t.Run("clusters array is not increased when paginated", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?page_token=token1&run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/runs/get?run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetRunRequest{RunId: 111222333}
		run, err := api.GetRun(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(run.JobClusters))
		assert.Equal(t, "cluster1", run.JobClusters[0].JobClusterKey)
		assert.Equal(t, "cluster2", run.JobClusters[1].JobClusterKey)
	})

	t.Run("run with two iterations pages", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?run_id=4444",
				Response: Run{
					Iterations: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
					},
					Tasks: []RunTask{
						{
							RunId: 999,
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?page_token=token1&run_id=4444",
				Response: Run{
					Iterations: []RunTask{
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					Tasks: []RunTask{
						{
							RunId: 999,
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/runs/get?run_id=4444",
				Response: Run{
					Iterations: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					Tasks: []RunTask{
						{
							RunId: 999,
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetRunRequest{RunId: 4444}
		run, err := api.GetRun(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 4, len(run.Iterations))
		assert.Equal(t, 1, len(run.Tasks))
		assert.Empty(t, run.NextPageToken)
		expected := []RunTask{
			{RunId: 123, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 1234, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 222, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 333, ForceSendFields: []string{"RunId", "TaskKey"}},
		}
		assert.Equal(t, expected, run.Iterations)
		assert.EqualValues(t, 999, run.Tasks[0].RunId)
	})
}

func TestGetJob(t *testing.T) {
	ctx := context.Background()

	t.Run("job with no pagination", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=514594995218126",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task1",
							},
							{
								TaskKey: "task2",
							},
							{
								TaskKey: "task3",
							},
							{
								TaskKey: "task4",
							},
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/get?run_id=514594995218126",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task1",
							},
							{
								TaskKey: "task2",
							},
							{
								TaskKey: "task3",
							},
							{
								TaskKey: "task4",
							},
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetJobRequest{JobId: 514594995218126}
		job, err := api.Get(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 4, len(job.Settings.Tasks))
		assert.EqualValues(t, "task1", job.Settings.Tasks[0].TaskKey)
		assert.EqualValues(t, "task4", job.Settings.Tasks[3].TaskKey)
	})

	t.Run("job with multiple pages", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=514594995218126",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task1",
							},
							{
								TaskKey: "task2",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "cluster1",
							},
							{
								JobClusterKey: "cluster2",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param1",
								Default: "default1",
							},
							{
								Name:    "param2",
								Default: "default2",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env1",
							},
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=514594995218126&page_token=token1",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task3",
							},
							{
								TaskKey: "task4",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "cluster3",
							},
							{
								JobClusterKey: "cluster4",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param3",
								Default: "default3",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env2",
							},
						},
					},
					NextPageToken: "token2",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=514594995218126&page_token=token2",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task5",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env3",
							},
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/get?run_id=514594995218126",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task1",
							},
							{
								TaskKey: "task2",
							},
							{
								TaskKey: "task3",
							},
							{
								TaskKey: "task4",
							},
							{
								TaskKey: "task5",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "cluster1",
							},
							{
								JobClusterKey: "cluster2",
							},
							{
								JobClusterKey: "cluster3",
							},
							{
								JobClusterKey: "cluster4",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param1",
								Default: "default1",
							},
							{
								Name:    "param2",
								Default: "default2",
							},
							{
								Name:    "param3",
								Default: "default3",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env1",
							},
							{
								EnvironmentKey: "env2",
							},
							{
								EnvironmentKey: "env3",
							},
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetJobRequest{JobId: 514594995218126}
		job, err := api.Get(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 5, len(job.Settings.Tasks))
		assert.Equal(t, 4, len(job.Settings.JobClusters))
		assert.Equal(t, 3, len(job.Settings.Parameters))
		assert.Equal(t, 3, len(job.Settings.Environments))
		assert.EqualValues(t, "task1", job.Settings.Tasks[0].TaskKey)
		assert.EqualValues(t, "task4", job.Settings.Tasks[3].TaskKey)
		assert.EqualValues(t, "task5", job.Settings.Tasks[4].TaskKey)
		assert.EqualValues(t, "cluster3", job.Settings.JobClusters[2].JobClusterKey)
		assert.EqualValues(t, "param3", job.Settings.Parameters[2].Name)
		assert.EqualValues(t, "env3", job.Settings.Environments[2].EnvironmentKey)
	})
}
