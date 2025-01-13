package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccJobsGetCorrectErrorNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)
	_, err := w.Jobs.GetByJobId(ctx, 123456789)
	assert.ErrorIs(t, err, apierr.ErrResourceDoesNotExist)
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
