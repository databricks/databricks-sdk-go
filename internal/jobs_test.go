package internal

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestCluster(ctx context.Context, wsc *workspaces.WorkspacesClient, t *testing.T) string {
	clusterName := RandomName("sdk-go-jobs-test-")

	// Fetch list of spark runtime versions
	sparkVersions, err := wsc.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latestLTS, err := sparkVersions.Select(clusters.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	require.NoError(t, err)

	// Fetch list of available node types
	nodeTypes, err := wsc.Clusters.ListNodeTypes(ctx)
	require.NoError(t, err)

	// Select the smallest node type id
	smallestWithDisk, err := nodeTypes.Smallest(clusters.NodeTypeRequest{
		LocalDisk: true,
	})
	require.NoError(t, err)

	// Create cluster and wait for it to start properly
	clstr, err := wsc.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latestLTS,
		NodeTypeId:             smallestWithDisk,
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
	wsc := workspaces.New()

	clusterId := createTestCluster(ctx, wsc, t)
	defer wsc.Clusters.PermanentDeleteByClusterId(ctx, clusterId)

	tmpPath := RandomName("/tmp/jobs-test")
	err := wsc.Workspace.MkdirsByPath(ctx, tmpPath)
	require.NoError(t, err)
	defer wsc.Workspace.Delete(ctx, workspace.Delete{
		Path: tmpPath,
	})

	filePath := filepath.Join(tmpPath, RandomName("slow-"))
	fileContent := "import time; time.sleep(10); dbutils.notebook.exit('hello')"
	err = wsc.Workspace.Import(ctx, workspace.Import{
		Content:   base64.StdEncoding.Strict().EncodeToString([]byte(fileContent)),
		Format:    "SOURCE",
		Language:  "PYTHON",
		Overwrite: true,
		Path:      filePath,
	})
	require.NoError(t, err)

	run, err := wsc.Jobs.SubmitAndWait(ctx, jobs.SubmitRun{
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
	defer wsc.Jobs.DeleteRunByRunId(ctx, run.RunId)

	output, err := wsc.Jobs.GetRunOutputByRunId(ctx, run.Tasks[0].RunId)
	require.NoError(t, err)
	assert.Equal(t, output.NotebookOutput.Result, "hello")

	createdJob, err := wsc.Jobs.Create(ctx, jobs.CreateJob{
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
	defer wsc.Jobs.DeleteByJobId(ctx, createdJob.JobId)

	run, err = wsc.Jobs.RunNowAndWait(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, run.Tasks)

	exportedView, err := wsc.Jobs.ExportRun(ctx, jobs.ExportRunRequest{
		RunId:         run.Tasks[0].RunId,
		ViewsToExport: "CODE",
	})
	require.NoError(t, err)
	assert.NotEmpty(t, exportedView.Views)
	assert.Equal(t, exportedView.Views[0].Type, jobs.ViewTypeNotebook)
	assert.NotEmpty(t, exportedView.Views[0].Content)

	_, err = wsc.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	require.NoError(t, err)

	err = wsc.Jobs.CancelAllRunsByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)

	runNowResponse, err := wsc.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	require.NoError(t, err)

	run, err = wsc.Jobs.CancelRunAndWait(ctx, jobs.CancelRun{
		RunId: runNowResponse.RunId,
	})
	require.NoError(t, err)

	run, err = wsc.Jobs.RepairRunAndWait(ctx, jobs.RepairRun{
		RerunTasks: []string{run.Tasks[0].TaskKey},
		RunId:      runNowResponse.RunId,
	})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(run.Tasks), 1)

	newName := RandomName("updated")
	err = wsc.Jobs.Update(ctx, jobs.UpdateJob{
		JobId: createdJob.JobId,
		NewSettings: &jobs.JobSettings{
			Name:              newName,
			MaxConcurrentRuns: 5,
		},
	})
	require.NoError(t, err)

	byId, err := wsc.Jobs.GetByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)

	assert.Equal(t, byId.Settings.Name, newName)
	assert.Equal(t, byId.Settings.MaxConcurrentRuns, 5)

	newName = RandomName("updated-for-reset")
	err = wsc.Jobs.Reset(ctx, jobs.ResetJob{
		JobId: createdJob.JobId,
		NewSettings: &jobs.JobSettings{
			Name:  newName,
			Tasks: byId.Settings.Tasks,
		},
	})
	require.NoError(t, err)

	byId, err = wsc.Jobs.GetByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)
	assert.Equal(t, byId.Settings.Name, newName)

	jobList, err := wsc.Jobs.List(ctx, jobs.ListRequest{
		ExpandTasks: false,
	})
	require.NoError(t, err)

	var jobsIdList []int64
	for _, job := range jobList.Jobs {
		jobsIdList = append(jobsIdList, job.JobId)
	}
	assert.Contains(t, jobsIdList, createdJob.JobId)

	err = wsc.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)

	jobList, err = wsc.Jobs.List(ctx, jobs.ListRequest{})
	require.NoError(t, err)
	for _, job := range jobList.Jobs {
		// TODO: change assertion to by-name
		assert.NotEqual(t, job.JobId, createdJob.JobId)
	}
}
