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
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ctx := context.Background()
	wsc := workspaces.New()

	clusterId := GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID")
	err := wsc.Clusters.StartByClusterIdAndWait(ctx, clusterId)
	require.NoError(t, err)

	tmpPath := RandomName("/tmp/jobs-test")
	err = wsc.Workspace.Mkdirs(ctx, workspace.MkdirsRequest{
		Path: tmpPath,
	})
	require.NoError(t, err)

	filePath := filepath.Join(tmpPath, RandomName())
	fileContent := "import time; time.sleep(10); dbutils.notebook.exit('hello')"
	err = wsc.Workspace.Import(ctx, workspace.ImportRequest{
		Content:   base64.StdEncoding.Strict().EncodeToString([]byte(fileContent)),
		Format:    "SOURCE",
		Language:  "PYTHON",
		Overwrite: true,
		Path:      filePath,
	})
	require.NoError(t, err)

	res, err := wsc.Workspace.GetStatus(ctx, workspace.GetStatusRequest{
		Path: filePath,
	})
	require.NoError(t, err)

	submitResp, err := wsc.Jobs.SubmitAndWait(ctx, jobs.SubmitRun{
		IdempotencyToken: RandomName(),
		RunName:          RandomName("go-sdk-"),
		Tasks: []jobs.RunSubmitTaskSettings{{
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: filePath,
			},
			TaskKey: RandomName(),
		}},
	})
	assert.NoError(t, err)
	output, err := wsc.Jobs.GetRunOutputByRunId(ctx, submitResp.RunId)
	assert.NoError(t, err)
	assert.Equal(t, output.NotebookOutput.Result, "hello")

	createdJob, err := wsc.Jobs.Create(ctx, jobs.CreateJob{
		Name: RandomName(),
		Tasks: []jobs.JobTaskSettings{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: res.Path,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	require.NoError(t, err)

	runResp, err := wsc.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	assert.NoError(t, err)

	runDetails, err := wsc.Jobs.GetRun(ctx, jobs.GetRunRequest{
		RunId: runResp.RunId,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, runDetails.Tasks)

	_, err = wsc.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	assert.NoError(t, err)

	err = wsc.Jobs.CancelAllRunsByJobId(ctx, createdJob.JobId)
	assert.NoError(t, err)

	err = wsc.Jobs.CancelRunAndWait(ctx, jobs.CancelRun{
		RunId: runDetails.Tasks[0].RunId,
	})
	assert.NoError(t, err)

	runDetails, err = wsc.Jobs.GetRun(ctx, jobs.GetRunRequest{
		RunId: runDetails.Tasks[0].RunId,
	})
	assert.NoError(t, err)
	assert.Equal(t, runDetails.State.LifeCycleState, jobs.RunLifeCycleStateTerminated)
	assert.Equal(t, runDetails.State.ResultState, jobs.RunResultStateCanceled)
	assert.True(t, runDetails.State.UserCancelledOrTimedout)

	_, err = wsc.Jobs.RepairRunAndWait(ctx, jobs.RepairRun{
		RerunTasks: []string{runDetails.Tasks[0].TaskKey},
		RunId:      runDetails.RunId,
	})
	assert.NoError(t, err)

	runDetails, err = wsc.Jobs.GetRun(ctx, jobs.GetRunRequest{
		RunId: runDetails.RunId,
	})
	assert.NoError(t, err)

	output, err = wsc.Jobs.GetRunOutputByRunId(ctx, runDetails.FirstTask().RunId)
	assert.NoError(t, err)
	assert.Equal(t, output.NotebookOutput.Result, "hello")

	runDetails, err = wsc.Jobs.GetRun(ctx, jobs.GetRunRequest{
		RunId: runDetails.RunId,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, runDetails.Tasks)

	exportedView, err := wsc.Jobs.ExportRun(ctx, jobs.ExportRunRequest{
		RunId:         runDetails.FirstTask().RunId,
		ViewsToExport: "CODE",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, exportedView.Views)
	assert.Equal(t, exportedView.Views[0].Type, jobs.ViewTypeNotebook)
	assert.NotEmpty(t, exportedView.Views[0].Content)

	newName := RandomName("updated")
	err = wsc.Jobs.Update(ctx, jobs.UpdateJob{
		JobId: createdJob.JobId,
		NewSettings: &jobs.JobSettings{
			Name:              newName,
			MaxConcurrentRuns: 5,
		},
	})
	assert.NoError(t, err)

	byId, err := wsc.Jobs.GetByJobId(ctx, createdJob.JobId)
	assert.NoError(t, err)

	assert.Equal(t, byId.Settings.Name, newName)
	assert.Equal(t, byId.Settings.MaxConcurrentRuns, 5)

	/** reset the job to default settings
	- max concurrent run : 1
	- Name : from default job
	- Tasks : from default job
	*/
	newName = RandomName("updated-for-reset")
	err = wsc.Jobs.Reset(ctx, jobs.ResetJob{
		JobId: createdJob.JobId,
		NewSettings: &jobs.JobSettings{
			Name:  newName,
			Tasks: byId.Settings.Tasks,
		},
	})
	assert.NoError(t, err)

	byId, err = wsc.Jobs.GetByJobId(ctx, createdJob.JobId)
	assert.NoError(t, err)
	assert.Equal(t, byId.Settings.Name, newName)

	jobList, err := wsc.Jobs.List(ctx, jobs.ListRequest{
		ExpandTasks: false,
	})
	assert.NoError(t, err)

	var jobsIdList []int64
	for _, job := range jobList.Jobs {
		jobsIdList = append(jobsIdList, job.JobId)
	}
	assert.Contains(t, jobsIdList, createdJob.JobId)

	err = wsc.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	assert.NoError(t, err)

	jobList, err = wsc.Jobs.List(ctx, jobs.ListRequest{})
	assert.NoError(t, err)
	for _, job := range jobList.Jobs {
		// TODO: change assertion to by-name
		assert.NotEqual(t, job.JobId, createdJob.JobId)
	}
}
