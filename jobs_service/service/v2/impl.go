// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v2

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	jobsv2 "github.com/databricks/databricks-sdk-go/protos/jobs/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

// unexported type that holds implementations of just JobsService API methods
type jobs_serviceServiceImpl struct {
	client *client.DatabricksClient
}

func (a *jobs_serviceServiceImpl) CreateJob(ctx context.Context, request *jobsv2.CreateJobRequest) (*jobsv2.Job, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var job jobsv2.Job
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil

}

func (a *jobs_serviceServiceImpl) DeleteJob(ctx context.Context, request *jobsv2.DeleteJobRequest) (*emptypb.Empty, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, nil)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil

}

func (a *jobs_serviceServiceImpl) GetJob(ctx context.Context, request *jobsv2.GetJobRequest) (*jobsv2.Job, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var job jobsv2.Job
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil

}

func (a *jobs_serviceServiceImpl) ListJobs(ctx context.Context, request *jobsv2.ListJobsRequest) (*jobsv2.ListJobsResponse, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var listJobsResponse jobsv2.ListJobsResponse
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, &listJobsResponse)
	if err != nil {
		return nil, err
	}

	return &listJobsResponse, nil

}

func (a *jobs_serviceServiceImpl) UpdateJob(ctx context.Context, request *jobsv2.UpdateJobRequest) (*jobsv2.Job, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var job jobsv2.Job
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil

}
