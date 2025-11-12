// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Postgres API methods
type postgresImpl struct {
	client *client.DatabricksClient
}

func (a *postgresImpl) CreateDatabaseBranch(ctx context.Context, request CreateDatabaseBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/branches", request.Parent)
	queryParams := make(map[string]any)

	if request.DatabaseBranchId != "" || slices.Contains(request.ForceSendFields, "DatabaseBranchId") {
		queryParams["database_branch_id"] = request.DatabaseBranchId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DatabaseBranch, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateDatabaseEndpoint(ctx context.Context, request CreateDatabaseEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)

	if request.DatabaseEndpointId != "" || slices.Contains(request.ForceSendFields, "DatabaseEndpointId") {
		queryParams["database_endpoint_id"] = request.DatabaseEndpointId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DatabaseEndpoint, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateDatabaseProject(ctx context.Context, request CreateDatabaseProjectRequest) (*Operation, error) {
	var operation Operation
	path := "/api/2.0/postgres/projects"
	queryParams := make(map[string]any)

	if request.DatabaseProjectId != "" || slices.Contains(request.ForceSendFields, "DatabaseProjectId") {
		queryParams["database_project_id"] = request.DatabaseProjectId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DatabaseProject, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteDatabaseBranch(ctx context.Context, request DeleteDatabaseBranchRequest) error {
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *postgresImpl) DeleteDatabaseEndpoint(ctx context.Context, request DeleteDatabaseEndpointRequest) error {
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *postgresImpl) DeleteDatabaseProject(ctx context.Context, request DeleteDatabaseProjectRequest) error {
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *postgresImpl) GetDatabaseBranch(ctx context.Context, request GetDatabaseBranchRequest) (*DatabaseBranch, error) {
	var databaseBranch DatabaseBranch
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &databaseBranch)
	return &databaseBranch, err
}

func (a *postgresImpl) GetDatabaseEndpoint(ctx context.Context, request GetDatabaseEndpointRequest) (*DatabaseEndpoint, error) {
	var databaseEndpoint DatabaseEndpoint
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &databaseEndpoint)
	return &databaseEndpoint, err
}

func (a *postgresImpl) GetDatabaseOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) GetDatabaseProject(ctx context.Context, request GetDatabaseProjectRequest) (*DatabaseProject, error) {
	var databaseProject DatabaseProject
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &databaseProject)
	return &databaseProject, err
}

// List Database Branches.
func (a *postgresImpl) ListDatabaseBranches(ctx context.Context, request ListDatabaseBranchesRequest) listing.Iterator[DatabaseBranch] {

	getNextPage := func(ctx context.Context, req ListDatabaseBranchesRequest) (*ListDatabaseBranchesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDatabaseBranches(ctx, req)
	}
	getItems := func(resp *ListDatabaseBranchesResponse) []DatabaseBranch {
		return resp.DatabaseBranches
	}
	getNextReq := func(resp *ListDatabaseBranchesResponse) *ListDatabaseBranchesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List Database Branches.
func (a *postgresImpl) ListDatabaseBranchesAll(ctx context.Context, request ListDatabaseBranchesRequest) ([]DatabaseBranch, error) {
	iterator := a.ListDatabaseBranches(ctx, request)
	return listing.ToSlice[DatabaseBranch](ctx, iterator)
}

func (a *postgresImpl) internalListDatabaseBranches(ctx context.Context, request ListDatabaseBranchesRequest) (*ListDatabaseBranchesResponse, error) {
	var listDatabaseBranchesResponse ListDatabaseBranchesResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/branches", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDatabaseBranchesResponse)
	return &listDatabaseBranchesResponse, err
}

// List Database Endpoints.
func (a *postgresImpl) ListDatabaseEndpoints(ctx context.Context, request ListDatabaseEndpointsRequest) listing.Iterator[DatabaseEndpoint] {

	getNextPage := func(ctx context.Context, req ListDatabaseEndpointsRequest) (*ListDatabaseEndpointsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDatabaseEndpoints(ctx, req)
	}
	getItems := func(resp *ListDatabaseEndpointsResponse) []DatabaseEndpoint {
		return resp.DatabaseEndpoints
	}
	getNextReq := func(resp *ListDatabaseEndpointsResponse) *ListDatabaseEndpointsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List Database Endpoints.
func (a *postgresImpl) ListDatabaseEndpointsAll(ctx context.Context, request ListDatabaseEndpointsRequest) ([]DatabaseEndpoint, error) {
	iterator := a.ListDatabaseEndpoints(ctx, request)
	return listing.ToSlice[DatabaseEndpoint](ctx, iterator)
}

func (a *postgresImpl) internalListDatabaseEndpoints(ctx context.Context, request ListDatabaseEndpointsRequest) (*ListDatabaseEndpointsResponse, error) {
	var listDatabaseEndpointsResponse ListDatabaseEndpointsResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDatabaseEndpointsResponse)
	return &listDatabaseEndpointsResponse, err
}

// List Database Projects.
func (a *postgresImpl) ListDatabaseProjects(ctx context.Context, request ListDatabaseProjectsRequest) listing.Iterator[DatabaseProject] {

	getNextPage := func(ctx context.Context, req ListDatabaseProjectsRequest) (*ListDatabaseProjectsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDatabaseProjects(ctx, req)
	}
	getItems := func(resp *ListDatabaseProjectsResponse) []DatabaseProject {
		return resp.DatabaseProjects
	}
	getNextReq := func(resp *ListDatabaseProjectsResponse) *ListDatabaseProjectsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List Database Projects.
func (a *postgresImpl) ListDatabaseProjectsAll(ctx context.Context, request ListDatabaseProjectsRequest) ([]DatabaseProject, error) {
	iterator := a.ListDatabaseProjects(ctx, request)
	return listing.ToSlice[DatabaseProject](ctx, iterator)
}

func (a *postgresImpl) internalListDatabaseProjects(ctx context.Context, request ListDatabaseProjectsRequest) (*ListDatabaseProjectsResponse, error) {
	var listDatabaseProjectsResponse ListDatabaseProjectsResponse
	path := "/api/2.0/postgres/projects"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDatabaseProjectsResponse)
	return &listDatabaseProjectsResponse, err
}

func (a *postgresImpl) RestartDatabaseEndpoint(ctx context.Context, request RestartDatabaseEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateDatabaseBranch(ctx context.Context, request UpdateDatabaseBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DatabaseBranch, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateDatabaseEndpoint(ctx context.Context, request UpdateDatabaseEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DatabaseEndpoint, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateDatabaseProject(ctx context.Context, request UpdateDatabaseProjectRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DatabaseProject, &operation)
	return &operation, err
}
