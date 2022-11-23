package internal

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestCluster(ctx context.Context, w *databricks.WorkspaceClient, t *testing.T) string {
	clusterName := RandomName(t.Name())

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latest, err := sparkVersions.Select(clusters.SparkVersionRequest{
		Latest: true,
	})
	require.NoError(t, err)

	// Create cluster and wait for it to start properly
	clstr, err := w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	require.NoError(t, err)
	return clstr.ClusterId
}

func TestAccJobsApiFullIntegration(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountClient() {
		t.SkipNow()
	}

	clusterId := createTestCluster(ctx, w, t)
	defer w.Clusters.PermanentDeleteByClusterId(ctx, clusterId)

	tmpPath := RandomName("/tmp/jobs-test")
	err := w.Workspace.MkdirsByPath(ctx, tmpPath)
	require.NoError(t, err)
	defer w.Workspace.Delete(ctx, workspace.Delete{
		Path: tmpPath,
	})

	filePath := filepath.Join(tmpPath, RandomName("slow-"))
	err = w.Workspace.Import(ctx, workspace.PythonNotebookOverwrite(filePath, `
		import time
		time.sleep(10)
		dbutils.notebook.exit('hello')`))
	require.NoError(t, err)

	run, err := w.Jobs.SubmitAndWait(ctx, jobs.SubmitRun{
		RunName: RandomName("go-sdk-SubmitAndWait-"),
		Tasks: []jobs.RunSubmitTaskSettings{{
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: filePath,
			},
			TaskKey: RandomName(),
		}},
	})
	require.NoError(t, err)
	defer w.Jobs.DeleteRunByRunId(ctx, run.RunId)

	output, err := w.Jobs.GetRunOutputByRunId(ctx, run.Tasks[0].RunId)
	require.NoError(t, err)
	assert.Equal(t, output.NotebookOutput.Result, "hello")

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: RandomName("go-sdk-Create-"),
		Tasks: []jobs.JobTaskSettings{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: filePath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	require.NoError(t, err)
	defer w.Jobs.DeleteByJobId(ctx, createdJob.JobId)

	run, err = w.Jobs.RunNowAndWait(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, run.Tasks)

	exportedView, err := w.Jobs.ExportRun(ctx, jobs.ExportRunRequest{
		RunId:         run.Tasks[0].RunId,
		ViewsToExport: "CODE",
	})
	require.NoError(t, err)
	assert.NotEmpty(t, exportedView.Views)
	assert.Equal(t, exportedView.Views[0].Type, jobs.ViewTypeNotebook)
	assert.NotEmpty(t, exportedView.Views[0].Content)

	_, err = w.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	require.NoError(t, err)

	err = w.Jobs.CancelAllRunsByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)

	runNowResponse, err := w.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	require.NoError(t, err)

	run, err = w.Jobs.CancelRunAndWait(ctx, jobs.CancelRun{
		RunId: runNowResponse.RunId,
	})
	require.NoError(t, err)

	run, err = w.Jobs.RepairRunAndWait(ctx, jobs.RepairRun{
		RerunTasks: []string{run.Tasks[0].TaskKey},
		RunId:      runNowResponse.RunId,
	})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(run.Tasks), 1)

	newName := RandomName("updated")
	err = w.Jobs.Update(ctx, jobs.UpdateJob{
		JobId: createdJob.JobId,
		NewSettings: &jobs.JobSettings{
			Name:              newName,
			MaxConcurrentRuns: 5,
		},
	})
	require.NoError(t, err)

	byId, err := w.Jobs.GetByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)

	assert.Equal(t, byId.Settings.Name, newName)
	assert.Equal(t, byId.Settings.MaxConcurrentRuns, 5)

	newName = RandomName("updated-for-reset")
	err = w.Jobs.Reset(ctx, jobs.ResetJob{
		JobId: createdJob.JobId,
		NewSettings: &jobs.JobSettings{
			Name:  newName,
			Tasks: byId.Settings.Tasks,
		},
	})
	require.NoError(t, err)

	byId, err = w.Jobs.GetByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)
	assert.Equal(t, byId.Settings.Name, newName)

	jobList, err := w.Jobs.ListAll(ctx, jobs.ListRequest{
		ExpandTasks: false,
	})
	require.NoError(t, err)

	var jobsIdList []int64
	for _, job := range jobList {
		jobsIdList = append(jobsIdList, job.JobId)
	}
	assert.Contains(t, jobsIdList, createdJob.JobId)

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)

	jobList, err = w.Jobs.ListAll(ctx, jobs.ListRequest{})
	require.NoError(t, err)
	for _, job := range jobList {
		// TODO: change assertion to by-name
		assert.NotEqual(t, job.JobId, createdJob.JobId)
	}
}

func TestAccJobsListAllNoDuplicates(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountClient() {
		t.SkipNow()
	}

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latestLTS, err := sparkVersions.Select(clusters.SparkVersionRequest{
		Latest: true,
	})
	require.NoError(t, err)

	// Fetch list of available node types
	nodeTypes, err := w.Clusters.ListNodeTypes(ctx)
	require.NoError(t, err)

	// Select the smallest node type id
	smallestWithDisk, err := nodeTypes.Smallest(clusters.NodeTypeRequest{
		LocalDisk: true,
	})
	require.NoError(t, err)

	for i := 0; i < 34; i++ {
		createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
			Name: RandomName(t.Name()),
			Tasks: []jobs.JobTaskSettings{{
				Description: "test",
				NewCluster: &clusters.CreateCluster{
					SparkVersion: latestLTS,
					NodeTypeId:   smallestWithDisk,
					NumWorkers:   1,
				},
				TaskKey:        "test",
				TimeoutSeconds: 0,
				NotebookTask: &jobs.NotebookTask{
					NotebookPath: "/foo/bar",
				},
			}},
		})
		require.NoError(t, err)
		t.Cleanup(func() {
			err := w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
			require.NoError(t, err)
		})
	}
	all, err := w.Jobs.ListAll(ctx, jobs.ListRequest{})
	require.NoError(t, err)
	ids := map[int64]bool{}
	for _, v := range all {
		ids[v.JobId] = true
	}
	assert.Equal(t, len(all), len(ids), "Listing produced duplicate results")
}
