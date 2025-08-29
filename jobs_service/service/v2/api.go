// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage JobsService Service, etc.
package v2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	jobsv2 "github.com/databricks/databricks-sdk-go/protos/jobs/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type JobsServiceInterface interface {
	CreateJob(ctx context.Context, request *jobsv2.CreateJobRequest) (*jobsv2.Job, error)

	DeleteJob(ctx context.Context, request *jobsv2.DeleteJobRequest) (*emptypb.Empty, error)

	GetJob(ctx context.Context, request *jobsv2.GetJobRequest) (*jobsv2.Job, error)

	ListJobs(ctx context.Context, request *jobsv2.ListJobsRequest) (*jobsv2.ListJobsResponse, error)

	UpdateJob(ctx context.Context, request *jobsv2.UpdateJobRequest) (*jobsv2.Job, error)
}

func NewJobsServiceService(client *client.DatabricksClient) *JobsServiceAPI {
	return &JobsServiceAPI{
		jobs_serviceServiceImpl: jobs_serviceServiceImpl{
			client: client,
		},
	}
}

type JobsServiceAPI struct {
	jobs_serviceServiceImpl
}
