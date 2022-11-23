// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
	"fmt"
	"net/http"
)

type databricksClient interface {
	Do(ctx context.Context, method string, path string, request any, response any) error
	ConfiguredAccountID() string
	IsAws() bool
}

// unexported type that holds implementations of just Pipelines API methods
type pipelinesImpl struct {
	client databricksClient
}

func (a *pipelinesImpl) CreatePipeline(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error) {
	var createPipelineResponse CreatePipelineResponse
	path := "/api/2.0/pipelines"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createPipelineResponse)
	return &createPipelineResponse, err
}

func (a *pipelinesImpl) DeletePipeline(ctx context.Context, request DeletePipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *pipelinesImpl) GetPipeline(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	var getPipelineResponse GetPipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getPipelineResponse)
	return &getPipelineResponse, err
}

func (a *pipelinesImpl) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {
	var getUpdateResponse GetUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates/%v", request.PipelineId, request.UpdateId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getUpdateResponse)
	return &getUpdateResponse, err
}

func (a *pipelinesImpl) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	var listUpdatesResponse ListUpdatesResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &listUpdatesResponse)
	return &listUpdatesResponse, err
}

func (a *pipelinesImpl) ResetPipeline(ctx context.Context, request ResetPipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v/reset", request.PipelineId)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *pipelinesImpl) StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error) {
	var startUpdateResponse StartUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	err := a.client.Do(ctx, http.MethodPost, path, request, &startUpdateResponse)
	return &startUpdateResponse, err
}

func (a *pipelinesImpl) StopPipeline(ctx context.Context, request StopPipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v/stop", request.PipelineId)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *pipelinesImpl) UpdatePipeline(ctx context.Context, request EditPipeline) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}
