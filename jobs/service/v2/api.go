// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Jobs Service, etc.
package v2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	jobsv2 "github.com/databricks/databricks-sdk-go/protos/jobs/v2"
)

type JobsServiceInterface interface {
	CreateJob(ctx context.Context, request *jobsv2.CreateJobRequest) (*jobsv2.Job, error)

	DeleteJob(ctx context.Context, request *jobsv2.DeleteJobRequest) error

	GetJob(ctx context.Context, request *jobsv2.GetJobRequest) (*jobsv2.Job, error)

	ListJobs(ctx context.Context, request *jobsv2.ListJobsRequest) (*jobsv2.ListJobsResponse, error)
}

func NewJobsService(client *client.DatabricksClient) *JobsServiceAPI {
	return &JobsServiceAPI{
		jobsServiceImpl: jobsServiceImpl{
			client: client,
		},
	}
}

type JobsServiceAPI struct {
	jobsServiceImpl
}
