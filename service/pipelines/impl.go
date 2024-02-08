// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Pipelines API methods
type pipelinesImpl struct {
	client *client.DatabricksClient
}

func (a *pipelinesImpl) Create(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error) {
	var createPipelineResponse CreatePipelineResponse
	path := "/api/2.0/pipelines"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createPipelineResponse)
	return &createPipelineResponse, err
}

func (a *pipelinesImpl) Delete(ctx context.Context, request DeletePipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *pipelinesImpl) Get(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	var getPipelineResponse GetPipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getPipelineResponse)
	return &getPipelineResponse, err
}

func (a *pipelinesImpl) GetPermissionLevels(ctx context.Context, request GetPipelinePermissionLevelsRequest) (*GetPipelinePermissionLevelsResponse, error) {
	var getPipelinePermissionLevelsResponse GetPipelinePermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v/permissionLevels", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getPipelinePermissionLevelsResponse)
	return &getPipelinePermissionLevelsResponse, err
}

func (a *pipelinesImpl) GetPermissions(ctx context.Context, request GetPipelinePermissionsRequest) (*PipelinePermissions, error) {
	var pipelinePermissions PipelinePermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &pipelinePermissions)
	return &pipelinePermissions, err
}

func (a *pipelinesImpl) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {
	var getUpdateResponse GetUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates/%v", request.PipelineId, request.UpdateId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getUpdateResponse)
	return &getUpdateResponse, err
}

func (a *pipelinesImpl) ListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) (*ListPipelineEventsResponse, error) {
	var listPipelineEventsResponse ListPipelineEventsResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/events", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listPipelineEventsResponse)
	return &listPipelineEventsResponse, err
}

func (a *pipelinesImpl) ListPipelines(ctx context.Context, request ListPipelinesRequest) (*ListPipelinesResponse, error) {
	var listPipelinesResponse ListPipelinesResponse
	path := "/api/2.0/pipelines"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listPipelinesResponse)
	return &listPipelinesResponse, err
}

func (a *pipelinesImpl) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	var listUpdatesResponse ListUpdatesResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listUpdatesResponse)
	return &listUpdatesResponse, err
}

func (a *pipelinesImpl) SetPermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {
	var pipelinePermissions PipelinePermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &pipelinePermissions)
	return &pipelinePermissions, err
}

func (a *pipelinesImpl) StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error) {
	var startUpdateResponse StartUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &startUpdateResponse)
	return &startUpdateResponse, err
}

func (a *pipelinesImpl) Stop(ctx context.Context, request StopRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v/stop", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, nil)
	return err
}

func (a *pipelinesImpl) Update(ctx context.Context, request EditPipeline) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, nil)
	return err
}

func (a *pipelinesImpl) UpdatePermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {
	var pipelinePermissions PipelinePermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &pipelinePermissions)
	return &pipelinePermissions, err
}
