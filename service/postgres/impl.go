// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Postgres API methods
type postgresImpl struct {
	client *client.DatabricksClient
}

func (a *postgresImpl) CreateBranch(ctx context.Context, request CreateBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/branches", request.Parent)
	queryParams := make(map[string]any)

	if request.BranchId != "" {
		queryParams["branch_id"] = request.BranchId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Branch, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)

	if request.EndpointId != "" {
		queryParams["endpoint_id"] = request.EndpointId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Endpoint, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateProject(ctx context.Context, request CreateProjectRequest) (*Operation, error) {
	var operation Operation
	path := "/api/2.0/postgres/projects"
	queryParams := make(map[string]any)

	if request.ProjectId != "" {
		queryParams["project_id"] = request.ProjectId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Project, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateRole(ctx context.Context, request CreateRoleRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/roles", request.Parent)
	queryParams := make(map[string]any)

	if request.RoleId != "" || slices.Contains(request.ForceSendFields, "RoleId") {
		queryParams["role_id"] = request.RoleId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Role, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteBranch(ctx context.Context, request DeleteBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteProject(ctx context.Context, request DeleteProjectRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteRole(ctx context.Context, request DeleteRoleRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error) {
	var databaseCredential DatabaseCredential
	path := "/api/2.0/postgres/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &databaseCredential)
	return &databaseCredential, err
}

func (a *postgresImpl) GetBranch(ctx context.Context, request GetBranchRequest) (*Branch, error) {
	var branch Branch
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &branch)
	return &branch, err
}

func (a *postgresImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &endpoint)
	return &endpoint, err
}

func (a *postgresImpl) GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) GetProject(ctx context.Context, request GetProjectRequest) (*Project, error) {
	var project Project
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &project)
	return &project, err
}

func (a *postgresImpl) GetRole(ctx context.Context, request GetRoleRequest) (*Role, error) {
	var role Role
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &role)
	return &role, err
}

// Returns a paginated list of database branches in the project.
func (a *postgresImpl) ListBranches(ctx context.Context, request ListBranchesRequest) listing.Iterator[Branch] {

	getNextPage := func(ctx context.Context, req ListBranchesRequest) (*ListBranchesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListBranches(ctx, req)
	}
	getItems := func(resp *ListBranchesResponse) []Branch {
		return resp.Branches
	}
	getNextReq := func(resp *ListBranchesResponse) *ListBranchesRequest {
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

// Returns a paginated list of database branches in the project.
func (a *postgresImpl) ListBranchesAll(ctx context.Context, request ListBranchesRequest) ([]Branch, error) {
	iterator := a.ListBranches(ctx, request)
	return listing.ToSlice[Branch](ctx, iterator)
}

func (a *postgresImpl) internalListBranches(ctx context.Context, request ListBranchesRequest) (*ListBranchesResponse, error) {
	var listBranchesResponse ListBranchesResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/branches", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBranchesResponse)
	return &listBranchesResponse, err
}

// Returns a paginated list of compute endpoints in the branch.
func (a *postgresImpl) ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[Endpoint] {

	getNextPage := func(ctx context.Context, req ListEndpointsRequest) (*ListEndpointsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListEndpoints(ctx, req)
	}
	getItems := func(resp *ListEndpointsResponse) []Endpoint {
		return resp.Endpoints
	}
	getNextReq := func(resp *ListEndpointsResponse) *ListEndpointsRequest {
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

// Returns a paginated list of compute endpoints in the branch.
func (a *postgresImpl) ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]Endpoint, error) {
	iterator := a.ListEndpoints(ctx, request)
	return listing.ToSlice[Endpoint](ctx, iterator)
}

func (a *postgresImpl) internalListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error) {
	var listEndpointsResponse ListEndpointsResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listEndpointsResponse)
	return &listEndpointsResponse, err
}

// Returns a paginated list of database projects in the workspace that the user
// has permission to access.
func (a *postgresImpl) ListProjects(ctx context.Context, request ListProjectsRequest) listing.Iterator[Project] {

	getNextPage := func(ctx context.Context, req ListProjectsRequest) (*ListProjectsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListProjects(ctx, req)
	}
	getItems := func(resp *ListProjectsResponse) []Project {
		return resp.Projects
	}
	getNextReq := func(resp *ListProjectsResponse) *ListProjectsRequest {
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

// Returns a paginated list of database projects in the workspace that the user
// has permission to access.
func (a *postgresImpl) ListProjectsAll(ctx context.Context, request ListProjectsRequest) ([]Project, error) {
	iterator := a.ListProjects(ctx, request)
	return listing.ToSlice[Project](ctx, iterator)
}

func (a *postgresImpl) internalListProjects(ctx context.Context, request ListProjectsRequest) (*ListProjectsResponse, error) {
	var listProjectsResponse ListProjectsResponse
	path := "/api/2.0/postgres/projects"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProjectsResponse)
	return &listProjectsResponse, err
}

// Returns a paginated list of Postgres roles in the branch.
func (a *postgresImpl) ListRoles(ctx context.Context, request ListRolesRequest) listing.Iterator[Role] {

	getNextPage := func(ctx context.Context, req ListRolesRequest) (*ListRolesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListRoles(ctx, req)
	}
	getItems := func(resp *ListRolesResponse) []Role {
		return resp.Roles
	}
	getNextReq := func(resp *ListRolesResponse) *ListRolesRequest {
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

// Returns a paginated list of Postgres roles in the branch.
func (a *postgresImpl) ListRolesAll(ctx context.Context, request ListRolesRequest) ([]Role, error) {
	iterator := a.ListRoles(ctx, request)
	return listing.ToSlice[Role](ctx, iterator)
}

func (a *postgresImpl) internalListRoles(ctx context.Context, request ListRolesRequest) (*ListRolesResponse, error) {
	var listRolesResponse ListRolesResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/roles", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRolesResponse)
	return &listRolesResponse, err
}

func (a *postgresImpl) UpdateBranch(ctx context.Context, request UpdateBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Branch, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Endpoint, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateProject(ctx context.Context, request UpdateProjectRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.HostType() == config.UnifiedHost && cfg.WorkspaceID != "" {
		headers["X-Databricks-Org-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Project, &operation)
	return &operation, err
}
