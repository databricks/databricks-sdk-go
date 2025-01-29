package jobs

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/qa"

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
