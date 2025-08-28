// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v2

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just JobsService API methods
type jobsServiceImpl struct {
	client *client.DatabricksClient
}

func (a *Impl) CreateJob(ctx context.Context, request CreateJobRequest) (*Job, error) {
	var job Job
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.Method, path, headers, queryParams, nil, &job)
	return &job, err
}

func (a *Impl) DeleteJob(ctx context.Context, request DeleteJobRequest) error {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.Method, path, headers, queryParams, nil, nil)
	return err
}

func (a *Impl) GetJob(ctx context.Context, request GetJobRequest) (*Job, error) {
	var job Job
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.Method, path, headers, queryParams, nil, &job)
	return &job, err
}

func (a *Impl) ListJobs(ctx context.Context, request ListJobsRequest) (*ListJobsResponse, error) {
	var listJobsResponse ListJobsResponse
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.Method, path, headers, queryParams, nil, &listJobsResponse)
	return &listJobsResponse, err
}
