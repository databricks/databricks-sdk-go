// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package deltapipelines

import (
	"context"
	"time"
)



type DeltaPipelinesService interface {
    
    CreatePipeline(ctx context.Context, createPipelineRequest CreatePipelineRequest) (*CreatePipelineResponse, error)// CreatePipeline and wait to reach RUNNING state
	CreatePipelineAndWait(ctx context.Context, request CreatePipelineRequest, timeout ...time.Duration) (*CreatePipelineResponse, error)
    
    DeletePipeline(ctx context.Context, deletePipelineRequest DeletePipelineRequest) error
	DeletePipelineByPipelineId(ctx context.Context, pipelineId string) error
    
    EditPipeline(ctx context.Context, editPipelineRequest EditPipelineRequest) error
    
    GetPipeline(ctx context.Context, getPipelineRequest GetPipelineRequest) (*GetPipelineResponse, error)// GetPipeline and wait to reach RUNNING state
	GetPipelineAndWait(ctx context.Context, request GetPipelineRequest, timeout ...time.Duration) (*GetPipelineResponse, error)
	GetPipelineByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error)
	GetPipelineByPipelineIdAndWait(ctx context.Context, pipelineId string, timeout ...time.Duration) (*GetPipelineResponse, error)
    
    GetUpdate(ctx context.Context, getUpdateRequest GetUpdateRequest) (*GetUpdateResponse, error)
	GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error)
    
    ListUpdates(ctx context.Context, listUpdatesRequest ListUpdatesRequest) (*ListUpdatesResponse, error)
	ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error)
    
    ResetPipeline(ctx context.Context, resetPipelineRequest ResetPipelineRequest) error// ResetPipeline and wait to reach RUNNING state
	ResetPipelineAndWait(ctx context.Context, request ResetPipelineRequest, timeout ...time.Duration) error
    // * Starts or queues a pipeline update.
    StartUpdate(ctx context.Context, startUpdateRequest StartUpdateRequest) (*StartUpdateResponse, error)
    
    StopPipeline(ctx context.Context, stopPipelineRequest StopPipelineRequest) error// StopPipeline and wait to reach IDLE state
	StopPipelineAndWait(ctx context.Context, request StopPipelineRequest, timeout ...time.Duration) error
}
