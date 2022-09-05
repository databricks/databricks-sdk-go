// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package deltapipelines

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type DeltapipelinesService interface {
	CreatePipeline(ctx context.Context, createPipelineRequest CreatePipelineRequest) (*CreatePipelineResponse, error)

	DeletePipeline(ctx context.Context, deletePipelineRequest DeletePipelineRequest) error

	EditPipeline(ctx context.Context, editPipelineRequest EditPipelineRequest) error

	GetPipeline(ctx context.Context, getPipelineRequest GetPipelineRequest) (*GetPipelineResponse, error)

	GetUpdate(ctx context.Context, getUpdateRequest GetUpdateRequest) (*GetUpdateResponse, error)

	ListUpdates(ctx context.Context, listUpdatesRequest ListUpdatesRequest) (*ListUpdatesResponse, error)

	ResetPipeline(ctx context.Context, resetPipelineRequest ResetPipelineRequest) error
	// * Starts or queues a pipeline update.
	StartUpdate(ctx context.Context, startUpdateRequest StartUpdateRequest) (*StartUpdateResponse, error)

	StopPipeline(ctx context.Context, stopPipelineRequest StopPipelineRequest) error
	GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error)
	GetPipelineByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error)
	DeletePipelineByPipelineId(ctx context.Context, pipelineId string) error
	ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error)
}

func New(client *client.DatabricksClient) DeltapipelinesService {
	return &DeltapipelinesAPI{
		client: client,
	}
}

type DeltapipelinesAPI struct {
	client *client.DatabricksClient
}

func (a *DeltapipelinesAPI) CreatePipeline(ctx context.Context, request CreatePipelineRequest) (*CreatePipelineResponse, error) {
	var createPipelineResponse CreatePipelineResponse
	path := "/api/2.0/pipelines"
	err := a.client.Post(ctx, path, request, &createPipelineResponse)
	return &createPipelineResponse, err
}

func (a *DeltapipelinesAPI) DeletePipeline(ctx context.Context, request DeletePipelineRequest) error {
	path := "/api/2.0/pipelines/" + request.PipelineId
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *DeltapipelinesAPI) EditPipeline(ctx context.Context, request EditPipelineRequest) error {
	path := "/api/2.0/pipelines/" + request.PipelineId
	err := a.client.Put(ctx, path, request)
	return err
}

func (a *DeltapipelinesAPI) GetPipeline(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	var getPipelineResponse GetPipelineResponse
	path := "/api/2.0/pipelines/" + request.PipelineId
	err := a.client.Get(ctx, path, request, &getPipelineResponse)
	return &getPipelineResponse, err
}

func (a *DeltapipelinesAPI) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {
	var getUpdateResponse GetUpdateResponse
	path := "/api/2.0/pipelines/" + request.PipelineId + "/updates/" + request.UpdateId
	err := a.client.Get(ctx, path, request, &getUpdateResponse)
	return &getUpdateResponse, err
}

func (a *DeltapipelinesAPI) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	var listUpdatesResponse ListUpdatesResponse
	path := "/api/2.0/pipelines/" + request.PipelineId
	err := a.client.Get(ctx, path, request, &listUpdatesResponse)
	return &listUpdatesResponse, err
}

func (a *DeltapipelinesAPI) ResetPipeline(ctx context.Context, request ResetPipelineRequest) error {
	path := "/api/2.0/pipelines/" + request.PipelineId
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// * Starts or queues a pipeline update.
func (a *DeltapipelinesAPI) StartUpdate(ctx context.Context, request StartUpdateRequest) (*StartUpdateResponse, error) {
	var startUpdateResponse StartUpdateResponse
	path := "/api/2.0/pipelines/" + request.PipelineId
	err := a.client.Post(ctx, path, request, &startUpdateResponse)
	return &startUpdateResponse, err
}

func (a *DeltapipelinesAPI) StopPipeline(ctx context.Context, request StopPipelineRequest) error {
	path := "/api/2.0/pipelines/" + request.PipelineId
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *DeltapipelinesAPI) GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error) {
	return a.GetUpdate(ctx, GetUpdateRequest{
		PipelineId: pipelineId,
		UpdateId:   updateId,
	})
}

func (a *DeltapipelinesAPI) GetPipelineByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error) {
	return a.GetPipeline(ctx, GetPipelineRequest{
		PipelineId: pipelineId,
	})
}

func (a *DeltapipelinesAPI) DeletePipelineByPipelineId(ctx context.Context, pipelineId string) error {
	return a.DeletePipeline(ctx, DeletePipelineRequest{
		PipelineId: pipelineId,
	})
}

func (a *DeltapipelinesAPI) ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error) {
	return a.ListUpdates(ctx, ListUpdatesRequest{
		PipelineId: pipelineId,
	})
}
