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

func (a *pipelinesImpl) CreatePipeline(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error) {
	var createPipelineResponse CreatePipelineResponse
	path := "/api/2.0/pipelines"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createPipelineResponse)
	return &createPipelineResponse, err
}

func (a *pipelinesImpl) DeletePipeline(ctx context.Context, request DeletePipeline) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *pipelinesImpl) GetPipeline(ctx context.Context, request GetPipeline) (*GetPipelineResponse, error) {
	var getPipelineResponse GetPipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getPipelineResponse)
	return &getPipelineResponse, err
}

func (a *pipelinesImpl) GetUpdate(ctx context.Context, request GetUpdate) (*GetUpdateResponse, error) {
	var getUpdateResponse GetUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates/%v", request.PipelineId, request.UpdateId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getUpdateResponse)
	return &getUpdateResponse, err
}

func (a *pipelinesImpl) ListPipelines(ctx context.Context, request ListPipelines) (*ListPipelinesResponse, error) {
	var listPipelinesResponse ListPipelinesResponse
	path := "/api/2.0/pipelines"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listPipelinesResponse)
	return &listPipelinesResponse, err
}

func (a *pipelinesImpl) ListUpdates(ctx context.Context, request ListUpdates) (*ListUpdatesResponse, error) {
	var listUpdatesResponse ListUpdatesResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &listUpdatesResponse)
	return &listUpdatesResponse, err
}

func (a *pipelinesImpl) ResetPipeline(ctx context.Context, request ResetPipeline) error {
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

func (a *pipelinesImpl) StopPipeline(ctx context.Context, request StopPipeline) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v/stop", request.PipelineId)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *pipelinesImpl) UpdatePipeline(ctx context.Context, request EditPipeline) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}
