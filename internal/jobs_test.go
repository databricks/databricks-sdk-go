package internal

import (
	"encoding/base64"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)

	run, err := w.Jobs.SubmitAndWait(ctx, jobs.SubmitRun{
		RunName: RandomName("go-sdk-SubmitAndWait-"),
		Tasks: []jobs.RunSubmitTaskSettings{{
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
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
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	require.NoError(t, err)
	defer w.Jobs.DeleteByJobId(ctx, createdJob.JobId)

	runById, err := w.Jobs.RunNowAndWait(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, runById.Tasks)

	exportedView, err := w.Jobs.ExportRun(ctx, jobs.ExportRunRequest{
		RunId:         runById.Tasks[0].RunId,
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

	cancelledRun, err := w.Jobs.CancelRunAndWait(ctx, jobs.CancelRun{
		RunId: runNowResponse.Response.RunId,
	})
	require.NoError(t, err)

	repairedRun, err := w.Jobs.RepairRunAndWait(ctx, jobs.RepairRun{
		RerunTasks: []string{cancelledRun.Tasks[0].TaskKey},
		RunId:      runNowResponse.Response.RunId,
	})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(repairedRun.Tasks), 1)

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
		JobId: byId.JobId,
		NewSettings: jobs.JobSettings{
			Name:  newName,
			Tasks: byId.Settings.Tasks,
		},
	})
	require.NoError(t, err)

	byId, err = w.Jobs.GetByJobId(ctx, createdJob.JobId)
	require.NoError(t, err)
	assert.Equal(t, byId.Settings.Name, newName)

	byName, err := w.Jobs.GetBySettingsName(ctx, newName)
	require.NoError(t, err)
	assert.Equal(t, byId.JobId, byName.JobId)

	jobList, err := w.Jobs.ListAll(ctx, jobs.ListJobsRequest{
		ExpandTasks: false,
	})
	require.NoError(t, err)
	assert.True(t, len(jobList) >= 1)
}

func TestAccJobsListAllNoDuplicatesNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latestLTS, err := sparkVersions.Select(compute.SparkVersionRequest{
		Latest: true,
	})
	require.NoError(t, err)

	// Select the smallest node type id
	smallestWithDisk, err := w.Clusters.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	require.NoError(t, err)

	for i := 0; i < 34; i++ {
		createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
			Name: RandomName(t.Name()),
			Tasks: []jobs.JobTaskSettings{{
				Description: "test",
				NewCluster: &compute.BaseClusterInfo{
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
	all, err := w.Jobs.ListAll(ctx, jobs.ListJobsRequest{})
	require.NoError(t, err)
	ids := map[int64]bool{}
	for _, v := range all {
		ids[v.JobId] = true
	}
	assert.Equal(t, len(all), len(ids), "Listing produced duplicate results")
}

func TestAccJobsListWithLimitNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)

	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latestLTS, err := sparkVersions.Select(compute.SparkVersionRequest{
		Latest: true,
	})
	require.NoError(t, err)

	// Select the smallest node type id
	smallestWithDisk, err := w.Clusters.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	require.NoError(t, err)

	for i := 0; i < 11; i++ {
		createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
			Name: RandomName(t.Name()),
			Tasks: []jobs.JobTaskSettings{{
				Description: "test",
				NewCluster: &compute.BaseClusterInfo{
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

	jobList, err := w.Jobs.ListAll(ctx, jobs.ListJobsRequest{
		Limit: 5,
	})

	require.NoError(t, err)
	assert.Equal(t, 5, len(jobList), "List returned correct amount of jobs (5)")

	jobList, err = w.Jobs.ListAll(ctx, jobs.ListJobsRequest{
		Limit: 10,
	})

	require.NoError(t, err)
	assert.Equal(t, 10, len(jobList), "List returned correct amount of jobs (10)")
}
