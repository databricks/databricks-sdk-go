package internal

import (
	"encoding/base64"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
)

func TestAccJobsApiFullIntegration(t *testing.T) {
	ctx, w := workspaceTest(t)
	clusterId := sharedRunningCluster(t, ctx, w)
	notebookPath := myNotebookPath(t, w)

	err := w.Workspace.Import(ctx, workspace.Import{
		Path:      notebookPath,
		Overwrite: true,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content: base64.StdEncoding.EncodeToString([]byte(`
			import time
			time.sleep(10)
			dbutils.notebook.exit('hello')`)),
	})
	assert.NoError(t, err)

	run, err := w.Jobs.SubmitAndWait(ctx, jobs.SubmitRun{
		RunName: RandomName("go-sdk-SubmitAndWait-"),
		Tasks: []jobs.SubmitTask{{
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey: RandomName(),
		}},
	})
	assert.NoError(t, err)
	defer w.Jobs.DeleteRunByRunId(ctx, run.RunId)

	output, err := w.Jobs.GetRunOutputByRunId(ctx, run.Tasks[0].RunId)
	assert.NoError(t, err)
	assert.Equal(t, output.NotebookOutput.Result, "hello")

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: RandomName("go-sdk-Create-"),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	assert.NoError(t, err)
	defer w.Jobs.DeleteByJobId(ctx, createdJob.JobId)

	runById, err := w.Jobs.RunNowAndWait(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, runById.Tasks)

	exportedView, err := w.Jobs.ExportRun(ctx, jobs.ExportRunRequest{
		RunId:         runById.Tasks[0].RunId,
		ViewsToExport: "CODE",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, exportedView.Views)
	assert.Equal(t, exportedView.Views[0].Type, jobs.ViewTypeNotebook)
	assert.NotEmpty(t, exportedView.Views[0].Content)

	_, err = w.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	assert.NoError(t, err)

	runList, err := w.Jobs.ListRunsAll(ctx, jobs.ListRunsRequest{
		JobId: createdJob.JobId,
	})
	assert.NoError(t, err)
	assert.Equal(t, createdJob.JobId, runList[0].JobId)

	err = w.Jobs.CancelAllRuns(ctx, jobs.CancelAllRuns{
		JobId: createdJob.JobId,
	})

	assert.NoError(t, err)

	runNowResponse, err := w.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	assert.NoError(t, err)

	cancelledRun, err := w.Jobs.CancelRunAndWait(ctx, jobs.CancelRun{
		RunId: runNowResponse.Response.RunId,
	})
	assert.NoError(t, err)

	repairedRun, err := w.Jobs.RepairRunAndWait(ctx, jobs.RepairRun{
		RerunTasks: []string{cancelledRun.Tasks[0].TaskKey},
		RunId:      runNowResponse.Response.RunId,
	})
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(repairedRun.Tasks), 1)

	newName := RandomName("updated")
	err = w.Jobs.Update(ctx, jobs.UpdateJob{
		JobId: createdJob.JobId,
		NewSettings: &jobs.JobSettings{
			Name:              newName,
			MaxConcurrentRuns: 5,
		},
	})
	assert.NoError(t, err)

	byId, err := w.Jobs.GetByJobId(ctx, createdJob.JobId)
	assert.NoError(t, err)

	assert.Equal(t, byId.Settings.Name, newName)
	assert.Equal(t, byId.Settings.MaxConcurrentRuns, 5)

	newName = RandomName("updated-for-reset")
	err = w.Jobs.Reset(ctx, jobs.ResetJob{
		JobId: byId.JobId,
		NewSettings: jobs.JobSettings{
			Name:  newName,
			Tasks: byId.Settings.Tasks,
		},
	})
	assert.NoError(t, err)

	byId, err = w.Jobs.GetByJobId(ctx, createdJob.JobId)
	assert.NoError(t, err)
	assert.Equal(t, byId.Settings.Name, newName)

	byName, err := w.Jobs.GetBySettingsName(ctx, newName)
	assert.NoError(t, err)
	assert.Equal(t, byId.JobId, byName.JobId)

	jobList, err := w.Jobs.ListAll(ctx, jobs.ListJobsRequest{
		ExpandTasks: false,
	})
	assert.NoError(t, err)
	assert.True(t, len(jobList) >= 1)
}

func TestAccJobsGetCorrectErrorNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)
	_, err := w.Jobs.GetByJobId(ctx, 123456789)
	assert.ErrorIs(t, err, apierr.ErrResourceDoesNotExist)
}

func TestAccJobsListAllNoDuplicatesNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	assert.NoError(t, err)

	// Select the latest LTS version
	latestLTS, err := sparkVersions.Select(compute.SparkVersionRequest{
		Latest: true,
	})
	assert.NoError(t, err)

	// Select the smallest node type id
	smallestWithDisk, err := w.Clusters.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	assert.NoError(t, err)

	for i := 0; i < 34; i++ {
		createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
			Name: RandomName(t.Name()),
			Tasks: []jobs.Task{{
				Description: "test",
				NewCluster: &compute.ClusterSpec{
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
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
			assert.NoError(t, err)
		})
	}
	all, err := w.Jobs.ListAll(ctx, jobs.ListJobsRequest{})
	assert.NoError(t, err)
	ids := map[int64]bool{}
	for _, v := range all {
		ids[v.JobId] = true
	}
	assert.Equal(t, len(all), len(ids), "Listing produced duplicate results")
}
