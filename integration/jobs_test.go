package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/jobs/v2"
	"github.com/stretchr/testify/assert"
)

func TestAccJobsGetCorrectErrorNoTranspile(t *testing.T) {
	ctx := workspaceTest(t)
	JobsApi, err := jobs.NewJobsClient(nil)
	assert.NoError(t, err)
	assert.NoError(t, err)
	resp, err := JobsApi.Create(ctx, jobs.CreateJob{
		Name: "test-job",
	})
	assert.NoError(t, err)
	job, err := JobsApi.GetByJobId(ctx, resp.JobId)
	assert.NoError(t, err)
	assert.Equal(t, job.JobId, resp.JobId)
}
