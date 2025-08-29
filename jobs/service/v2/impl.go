// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v2

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	jobsv2 "github.com/databricks/databricks-sdk-go/protos/jobs/v2"
)

// unexported type that holds implementations of just JobsService API methods
type jobsServiceImpl struct {
	client *client.DatabricksClient
}

func (a *jobsServiceImpl) CreateJob(ctx context.Context, request *jobsv2.CreateJobRequest) (*jobsv2.Job, error) {
	path := "/jobs/v2/jobs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var job jobsv2.Job
	err := a.client.DoProto(ctx, http.MethodPost, path, headers, queryParams, request, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (a *jobsServiceImpl) DeleteJob(ctx context.Context, request *jobsv2.DeleteJobRequest) error {
	path := "/jobs/v2/jobs/" + *request.Name
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	return a.client.DoProto(ctx, http.MethodDelete, path, headers, queryParams, nil, nil)
}

func (a *jobsServiceImpl) GetJob(ctx context.Context, request *jobsv2.GetJobRequest) (*jobsv2.Job, error) {
	path := "/jobs/v2/js/" + *request.Name
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var job jobsv2.Job
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, nil, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (a *jobsServiceImpl) ListJobs(ctx context.Context, request *jobsv2.ListJobsRequest) (*jobsv2.ListJobsResponse, error) {
	path := "/jobs/v2/jobs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	// Add query parameters from request
	if request.PageSize != nil {
		queryParams["page_size"] = *request.PageSize
	}
	if request.PageToken != nil {
		queryParams["page_token"] = *request.PageToken
	}

	var listJobsResponse jobsv2.ListJobsResponse
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, nil, &listJobsResponse)
	if err != nil {
		return nil, err
	}

	return &listJobsResponse, nil
}
