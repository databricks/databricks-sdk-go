package mocking

import (
	"context"
	"testing"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	// "github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/testkit/mocks"
)

func TestListJobs(t *testing.T) {
	client := mocks.NewMockWorkspaceClient(t)
	client.GetMockJobsAPI().On("ListAll", mock.Anything, jobs.ListJobsRequest{Limit: 10}).Return(
		[]jobs.BaseJob{
			{
				JobId: 64,
			},
			{
				JobId: 65,
			},
		}, nil,
	)
	jobs, err := client.WorkspaceClient.Jobs.ListAll(context.Background(), jobs.ListJobsRequest{Limit: 10})
	require.NoError(t, err)
	require.Equal(t, jobs[0].JobId, int64(64))
	require.Equal(t, jobs[1].JobId, int64(65))
}
