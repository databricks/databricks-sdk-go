package internal

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccJobsApiFullIntegration(t *testing.T) {
	return
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.Background()
	wsc := workspaces.New()

	clusterId := GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID")
	clusterInfo, err := wsc.Clusters.GetByClusterId(ctx, clusterId)
	require.NoError(t, err)

	if !clusterInfo.IsRunningOrResizing() {
		_, err := wsc.Clusters.StartByClusterIdAndWait(ctx, clusterId)
		require.NoError(t, err)
	}

	tmpPath := RandomName("/tmp/jobs-test")
	err = wsc.Workspace.MkdirsByPath(ctx, tmpPath)
	require.NoError(t, err)
	defer wsc.Workspace.Delete(ctx, workspace.DeleteRequest{
		Path: tmpPath,
	})

	filePath := filepath.Join(tmpPath, RandomName("slow-"))
	fileContent := "import time; time.sleep(10); dbutils.notebook.exit('hello')"
	err = wsc.Workspace.Import(ctx, workspace.ImportRequest{
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
