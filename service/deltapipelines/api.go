// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package deltapipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/retries"
)

func NewDeltaPipelines(client *client.DatabricksClient) DeltaPipelinesService {
	return &DeltaPipelinesAPI{
		client: client,
	}
}

type DeltaPipelinesAPI struct {
	client *client.DatabricksClient
}

func (a *DeltaPipelinesAPI) CreatePipeline(ctx context.Context, request CreatePipelineRequest) (*CreatePipelineResponse, error) {
	var createPipelineResponse CreatePipelineResponse
	path := "/api/2.0/pipelines"
	err := a.client.Post(ctx, path, request, &createPipelineResponse)
	return &createPipelineResponse, err
}

// CreatePipeline and wait to reach RUNNING state
func (a *DeltaPipelinesAPI) CreatePipelineAndWait(ctx context.Context, request CreatePipelineRequest, timeout ...time.Duration) (*CreatePipelineResponse, error) {
	createPipelineResponse, err := a.CreatePipeline(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return createPipelineResponse, retries.Wait(ctx, timeout[0], func() *retries.Err {
		getPipelineResponse, err := a.GetPipeline(ctx, GetPipelineRequest{
			PipelineId: createPipelineResponse.PipelineId,
		})
		if err != nil {
			return retries.Halt(err)
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case GetPipelineResponseStateRunning: // target state
			return nil
		case GetPipelineResponseStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetPipelineResponseStateRunning, status, statusMessage)
			return retries.Halt(err)
		default:
			return retries.Continues(statusMessage)
		}
	})
}

func (a *DeltaPipelinesAPI) DeletePipeline(ctx context.Context, request DeletePipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *DeltaPipelinesAPI) DeletePipelineByPipelineId(ctx context.Context, pipelineId string) error {
	return a.DeletePipeline(ctx, DeletePipelineRequest{
		PipelineId: pipelineId,
	})
}

func (a *DeltaPipelinesAPI) EditPipeline(ctx context.Context, request EditPipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Put(ctx, path, request)
	return err
}

func (a *DeltaPipelinesAPI) GetPipeline(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	var getPipelineResponse GetPipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Get(ctx, path, request, &getPipelineResponse)
	return &getPipelineResponse, err
}

// GetPipeline and wait to reach RUNNING state
func (a *DeltaPipelinesAPI) GetPipelineAndWait(ctx context.Context, request GetPipelineRequest, timeout ...time.Duration) (*GetPipelineResponse, error) {
	getPipelineResponse, err := a.GetPipeline(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return getPipelineResponse, retries.Wait(ctx, timeout[0], func() *retries.Err {
		getPipelineResponse, err := a.GetPipeline(ctx, GetPipelineRequest{
			PipelineId: getPipelineResponse.PipelineId,
		})
		if err != nil {
			return retries.Halt(err)
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case GetPipelineResponseStateRunning: // target state
			return nil
		case GetPipelineResponseStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetPipelineResponseStateRunning, status, statusMessage)
			return retries.Halt(err)
		default:
			return retries.Continues(statusMessage)
		}
	})
}

func (a *DeltaPipelinesAPI) GetPipelineByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error) {
	return a.GetPipeline(ctx, GetPipelineRequest{
		PipelineId: pipelineId,
	})
}

func (a *DeltaPipelinesAPI) GetPipelineByPipelineIdAndWait(ctx context.Context, pipelineId string, timeout ...time.Duration) (*GetPipelineResponse, error) {
	return a.GetPipelineAndWait(ctx, GetPipelineRequest{
		PipelineId: pipelineId,
	}, timeout...)
}

func (a *DeltaPipelinesAPI) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {
	var getUpdateResponse GetUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates/%v", request.PipelineId, request.UpdateId)
	err := a.client.Get(ctx, path, request, &getUpdateResponse)
	return &getUpdateResponse, err
}

func (a *DeltaPipelinesAPI) GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error) {
	return a.GetUpdate(ctx, GetUpdateRequest{
		PipelineId: pipelineId,
		UpdateId:   updateId,
	})
}

func (a *DeltaPipelinesAPI) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	var listUpdatesResponse ListUpdatesResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	err := a.client.Get(ctx, path, request, &listUpdatesResponse)
	return &listUpdatesResponse, err
}

func (a *DeltaPipelinesAPI) ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error) {
	return a.ListUpdates(ctx, ListUpdatesRequest{
		PipelineId: pipelineId,
	})
}

func (a *DeltaPipelinesAPI) ResetPipeline(ctx context.Context, request ResetPipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v/reset", request.PipelineId)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// ResetPipeline and wait to reach RUNNING state
func (a *DeltaPipelinesAPI) ResetPipelineAndWait(ctx context.Context, request ResetPipelineRequest, timeout ...time.Duration) error {
	err := a.ResetPipeline(ctx, request)
	if err != nil {
		return err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Wait(ctx, timeout[0], func() *retries.Err {
		getPipelineResponse, err := a.GetPipeline(ctx, GetPipelineRequest{
			PipelineId: request.PipelineId,
		})
		if err != nil {
			return retries.Halt(err)
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case GetPipelineResponseStateRunning: // target state
			return nil
		case GetPipelineResponseStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetPipelineResponseStateRunning, status, statusMessage)
			return retries.Halt(err)
		default:
			return retries.Continues(statusMessage)
		}
	})
}

// * Starts or queues a pipeline update.
func (a *DeltaPipelinesAPI) StartUpdate(ctx context.Context, request StartUpdateRequest) (*StartUpdateResponse, error) {
	var startUpdateResponse StartUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	err := a.client.Post(ctx, path, request, &startUpdateResponse)
	return &startUpdateResponse, err
}

func (a *DeltaPipelinesAPI) StopPipeline(ctx context.Context, request StopPipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v/stop", request.PipelineId)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// StopPipeline and wait to reach IDLE state
func (a *DeltaPipelinesAPI) StopPipelineAndWait(ctx context.Context, request StopPipelineRequest, timeout ...time.Duration) error {
	err := a.StopPipeline(ctx, request)
	if err != nil {
		return err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Wait(ctx, timeout[0], func() *retries.Err {
		getPipelineResponse, err := a.GetPipeline(ctx, GetPipelineRequest{
			PipelineId: request.PipelineId,
		})
		if err != nil {
			return retries.Halt(err)
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case GetPipelineResponseStateIdle: // target state
			return nil
		case GetPipelineResponseStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetPipelineResponseStateIdle, status, statusMessage)
			return retries.Halt(err)
		default:
			return retries.Continues(statusMessage)
		}
	})
}
