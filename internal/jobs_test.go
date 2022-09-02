package internal

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func waitForRunToFinish(t *testing.T, ctx context.Context, jobsService jobs.JobsService, waitTimeout time.Duration, runId int64) {
	err := retries.Wait(ctx, waitTimeout, func() *retries.Err {
		runDetails, err := jobsService.GetRun(ctx, jobs.GetRunRequest{
			IncludeHistory: false,
			RunId:          runId,
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, runDetails.Tasks)

		if runDetails.State.LifeCycleState == jobs.RunLifeCycleStateTerminated {
			return nil
		}
		if slices.Contains([]jobs.RunLifeCycleState{jobs.RunLifeCycleStatePending, jobs.RunLifeCycleStateRunning, jobs.RunLifeCycleStateTerminating}, runDetails.State.LifeCycleState) {
			return retries.Continues("pending")
		}
		return retries.Halt(fmt.Errorf("Unexpected run stated"))
	})
	assert.NoError(t, err)
}

func getLastTaskOutput(t *testing.T, ctx context.Context, jobsService jobs.JobsService, runId int64) *jobs.RunOutput {
	runDetails, err := jobsService.GetRun(ctx, jobs.GetRunRequest{
		RunId: runId,
	})
	assert.NoError(t, err)

	output, err := jobsService.GetRunOutput(ctx, jobs.GetRunOutputRequest{
		RunId: runDetails.Tasks[len(runDetails.Tasks)-1].RunId,
	})
	assert.NoError(t, err)

	return output
}

func createTestJob(t *testing.T, ctx context.Context, jobsService jobs.JobsService, testPath string, testId string) *jobs.CreateJobResponse {
	createResp, err := jobsService.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("Test job %s", testId),
		Tasks: []jobs.JobTaskSettings{{
			Description:       "test",
			ExistingClusterId: GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID"),
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: testPath,
			},
			TaskKey:        fmt.Sprintf("Testjob%s", testId),
			TimeoutSeconds: 0,
		}},
		TimeoutSeconds: 0,
	})
	assert.NoError(t, err)
	return createResp
}

func assertRunCanceled(t *testing.T, ctx context.Context, jobsService jobs.JobsService, runId int64) {
	waitForRunToFinish(t, ctx, jobsService, 60*time.Second, runId)
	runDetails, err := jobsService.GetRun(ctx, jobs.GetRunRequest{
		RunId: runId,
	})
	assert.NoError(t, err)
	assert.Equal(t, runDetails.State.LifeCycleState, jobs.RunLifeCycleStateTerminated)
	assert.Equal(t, runDetails.State.ResultState, jobs.RunResultStateCanceled)
	assert.True(t, runDetails.State.UserCancelledOrTimedout)
}

func setup(t *testing.T) (ctx context.Context, apiClient *client.DatabricksClient) {
	ctx = context.Background()
	apiClient = client.New(&databricks.Config{})
	startDefaultTestCluster(t, ctx, apiClient)
	return ctx, apiClient
}

func TestAccCreateAndRunJob(t *testing.T) {
	ctx, apiClient := setup(t)
	jobsService := jobs.New(apiClient)
	testId := fmt.Sprint(rand.Int())

	pythonPath := createTestPythonFile(t, ctx, apiClient, "/tmp/go-sdk-jobs-tests", testId, "dbutils.notebook.exit('hello')")
	createResp := createTestJob(t, ctx, jobsService, pythonPath, testId)

	runResp, err := jobsService.RunNow(ctx, jobs.RunNow{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)

	runList, err := jobsService.ListRuns(ctx, jobs.ListRunsRequest{
		CompletedOnly: false,
		JobId:         createResp.JobId,
		Limit:         1,
	})
	assert.NoError(t, err)
	var runIdList []int64
	for _, run := range runList.Runs {
		runIdList = append(runIdList, run.RunId)
	}
	assert.Contains(t, runIdList, runResp.RunId)

	waitForRunToFinish(t, ctx, jobsService, 60*time.Second, runResp.RunId)
	output := getLastTaskOutput(t, ctx, jobsService, runResp.RunId)
	assert.Equal(t, output.NotebookOutput.Result, "hello")

	jobsService.DeleteRun(ctx, jobs.DeleteRun{
		RunId: runResp.RunId,
	})
	runList, err = jobsService.ListRuns(ctx, jobs.ListRunsRequest{
		CompletedOnly: false,
		JobId:         createResp.JobId,
		Limit:         1,
	})
	assert.NoError(t, err)
	for _, run := range runList.Runs {
		assert.NotEqual(t, run.RunId, runResp.RunId)
	}
}

func TestAccSubmitOneTimeRun(t *testing.T) {
	ctx, apiClient := setup(t)
	jobsService := jobs.New(apiClient)
	testId := fmt.Sprint(rand.Int())
	testPath := createTestPythonFile(t, ctx, apiClient, "/tmp/go-sdk-jobs-tests", testId, "dbutils.notebook.exit('hello')")

	submitResp, err := jobsService.SubmitRun(ctx, jobs.SubmitRun{
		IdempotencyToken: fmt.Sprintf("test-%s", testId),
		RunName:          fmt.Sprintf("test-%s", testId),
		Tasks: []jobs.RunSubmitTaskSettings{{
			ExistingClusterId: GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID"),
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: testPath,
			},
			TaskKey: fmt.Sprintf("Testjob%s", testId),
		}},
	})
	assert.NoError(t, err)

	waitForRunToFinish(t, ctx, jobsService, 60*time.Second, submitResp.RunId)
	output := getLastTaskOutput(t, ctx, jobsService, submitResp.RunId)
	assert.Equal(t, output.NotebookOutput.Result, "hello")
}

func TestAccCreateAndCancelRun(t *testing.T) {
	ctx, apiClient := setup(t)
	jobsService := jobs.New(apiClient)
	testId := fmt.Sprint(rand.Int())

	fileContent := "import time; time.sleep(10); dbutils.notebook.exit('hello')"
	pythonPath := createTestPythonFile(t, ctx, apiClient, "/tmp/go-sdk-jobs-tests", testId, fileContent)
	createResp := createTestJob(t, ctx, jobsService, pythonPath, testId)

	//Cancel single run
	runResp, err := jobsService.RunNow(ctx, jobs.RunNow{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)

	jobsService.CancelRun(ctx, jobs.CancelRun{
		RunId: runResp.RunId,
	})
	assertRunCanceled(t, ctx, jobsService, runResp.RunId)

	//Cancel all runs
	runResp, err = jobsService.RunNow(ctx, jobs.RunNow{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)
	runResp2, err := jobsService.RunNow(ctx, jobs.RunNow{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)

	err = jobsService.CancelAllRuns(ctx, jobs.CancelAllRuns{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)

	assertRunCanceled(t, ctx, jobsService, runResp.RunId)
	assertRunCanceled(t, ctx, jobsService, runResp2.RunId)
}

func TestAccCreateAndDeleteJob(t *testing.T) {
	ctx, apiClient := setup(t)
	jobsService := jobs.New(apiClient)
	testId := fmt.Sprint(rand.Int())

	fileContent := "import time; time.sleep(10); dbutils.notebook.exit('hello')"
	pythonPath := createTestPythonFile(t, ctx, apiClient, "/tmp/go-sdk-jobs-tests", testId, fileContent)
	createResp := createTestJob(t, ctx, jobsService, pythonPath, testId)

	jobList, err := jobsService.ListJobs(ctx, jobs.ListJobsRequest{
		ExpandTasks: false,
	})
	assert.NoError(t, err)

	var jobsIdList []int64
	for _, job := range jobList.Jobs {
		jobsIdList = append(jobsIdList, job.JobId)
	}
	assert.Contains(t, jobsIdList, createResp.JobId)

	err = jobsService.DeleteJob(ctx, jobs.DeleteJob{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)

	jobList, err = jobsService.ListJobs(ctx, jobs.ListJobsRequest{
		ExpandTasks: false,
	})
	assert.NoError(t, err)
	for _, job := range jobList.Jobs {
		assert.NotEqual(t, job.JobId, createResp.JobId)
	}
}

func TestAccResetAndUpdateJob(t *testing.T) {
	ctx, apiClient := setup(t)
	jobsService := jobs.New(apiClient)
	testId := fmt.Sprint(rand.Int())

	fileContent := "import time; time.sleep(10); dbutils.notebook.exit('hello')"
	pythonPath := createTestPythonFile(t, ctx, apiClient, "/tmp/go-sdk-jobs-tests", testId, fileContent)
	createResp := createTestJob(t, ctx, jobsService, pythonPath, testId)

	defaultJobDetails, err := jobsService.GetJob(ctx, jobs.GetJobRequest{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)

	err = jobsService.UpdateJob(ctx, jobs.UpdateJobRequest{
		JobId: createResp.JobId,
		NewSettings: &jobs.JobSettings{
			Name:              fmt.Sprintf("New name %s", testId),
			MaxConcurrentRuns: 5,
		},
	})
	assert.NoError(t, err)
	jobDetails, err := jobsService.GetJob(ctx, jobs.GetJobRequest{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)
	assert.Equal(t, jobDetails.Settings.Name, fmt.Sprintf("New name %s", testId))
	assert.Equal(t, jobDetails.Settings.MaxConcurrentRuns, 5)
	//Update operation should not change other fields
	jobDetails.Settings.Name = defaultJobDetails.Settings.Name
	jobDetails.Settings.MaxConcurrentRuns = defaultJobDetails.Settings.MaxConcurrentRuns
	assert.Equal(t, jobDetails, defaultJobDetails)

	/** reset the job to default settings
	- max concurrent run : 1
	- Name : from default job
	- Tasks : from default job
	*/
	err = jobsService.ResetJob(ctx, jobs.ResetJobRequest{
		JobId: createResp.JobId,
		NewSettings: &jobs.JobSettings{
			Name:  defaultJobDetails.Settings.Name,
			Tasks: defaultJobDetails.Settings.Tasks,
		},
	})
	assert.NoError(t, err)
	jobDetails, err = jobsService.GetJob(ctx, jobs.GetJobRequest{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)
	assert.Equal(t, jobDetails, defaultJobDetails)
}

func TestAccRepairAndExportRun(t *testing.T) {
	ctx, apiClient := setup(t)
	jobsService := jobs.NewJobs(apiClient)
	testId := fmt.Sprint(rand.Int())

	fileContent := "import time; time.sleep(10); dbutils.notebook.exit('hello')"
	pythonPath := createTestPythonFile(t, ctx, apiClient, "/tmp/go-sdk-jobs-tests", testId, fileContent)
	createResp := createTestJob(t, ctx, jobsService, pythonPath, testId)

	runResp, err := jobsService.RunNow(ctx, jobs.RunNow{
		JobId: createResp.JobId,
	})
	assert.NoError(t, err)
	runDetails, err := jobsService.GetRun(ctx, jobs.GetRunRequest{
		RunId: runResp.RunId,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, runDetails.Tasks)
	err = jobsService.CancelRun(ctx, jobs.CancelRun{
		RunId: runDetails.Tasks[0].RunId,
	})
	assert.NoError(t, err)
	assertRunCanceled(t, ctx, jobsService, runDetails.Tasks[0].RunId)
	waitForRunToFinish(t, ctx, jobsService, 1*time.Minute, runDetails.RunId)

	_, err = jobsService.RepairRun(ctx, jobs.RepairRun{
		RerunTasks: []string{runDetails.Tasks[0].TaskKey},
		RunId:      runDetails.RunId,
	})
	assert.NoError(t, err)

	waitForRunToFinish(t, ctx, jobsService, 1*time.Minute, runDetails.RunId)
	output := getLastTaskOutput(t, ctx, jobsService, runDetails.RunId)
	assert.Equal(t, output.NotebookOutput.Result, "hello")

	runDetails, err = jobsService.GetRun(ctx, jobs.GetRunRequest{
		RunId: runDetails.RunId,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, runDetails.Tasks)
	exportedView, err := jobsService.ExportRun(ctx, jobs.ExportRunRequest{
		RunId:         runDetails.Tasks[len(runDetails.Tasks)-1].RunId,
		ViewsToExport: "CODE",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, exportedView.Views)
	assert.Equal(t, exportedView.Views[0].Type, jobs.ViewTypeNotebook)
	assert.NotEmpty(t, exportedView.Views[0].Content)
}
