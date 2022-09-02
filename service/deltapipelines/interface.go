// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package deltapipelines

import (
	"context"
)



type DeltaPipelinesService interface {
    
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
	DeletePipelineByPipelineId(ctx context.Context, pipelineId string) error
	GetPipelineByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error)
	GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error)
	ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error)
}
