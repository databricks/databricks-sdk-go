// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

func NewPipelines(client *client.DatabricksClient) PipelinesService {
	return &PipelinesAPI{
		client: client,
	}
}

type PipelinesAPI struct {
	client *client.DatabricksClient
}

// Create a pipeline
//
// Creates a new data processing pipeline based on the requested configuration.
// If successful, this method returns the ID of the new pipeline.
func (a *PipelinesAPI) CreatePipeline(ctx context.Context, request CreatePipelineRequest) (*CreatePipelineResponse, error) {
	var createPipelineResponse CreatePipelineResponse
	path := "/api/2.0/pipelines"
	err := a.client.Post(ctx, path, request, &createPipelineResponse)
	return &createPipelineResponse, err
}

// CreatePipeline and wait to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
func (a *PipelinesAPI) CreatePipelineAndWait(ctx context.Context, createPipelineRequest CreatePipelineRequest, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	createPipelineResponse, err := a.CreatePipeline(ctx, createPipelineRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetPipelineResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetPipelineResponse](ctx, i.Timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := a.GetPipeline(ctx, GetPipelineRequest{
			PipelineId: createPipelineResponse.PipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetPipelineResponse]{
				Info:    *getPipelineResponse,
				Timeout: i.Timeout,
			})
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case GetPipelineResponseStateRunning: // target state
			return getPipelineResponse, nil
		case GetPipelineResponseStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetPipelineResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Delete a pipeline
//
// Deletes a pipeline.
func (a *PipelinesAPI) DeletePipeline(ctx context.Context, request DeletePipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a pipeline
//
// Deletes a pipeline.
func (a *PipelinesAPI) DeletePipelineByPipelineId(ctx context.Context, pipelineId string) error {
	return a.DeletePipeline(ctx, DeletePipelineRequest{
		PipelineId: pipelineId,
	})
}

// Get a pipeline
func (a *PipelinesAPI) GetPipeline(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	var getPipelineResponse GetPipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Get(ctx, path, request, &getPipelineResponse)
	return &getPipelineResponse, err
}

// GetPipeline and wait to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
func (a *PipelinesAPI) GetPipelineAndWait(ctx context.Context, getPipelineRequest GetPipelineRequest, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	getPipelineResponse, err := a.GetPipeline(ctx, getPipelineRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetPipelineResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetPipelineResponse](ctx, i.Timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := a.GetPipeline(ctx, GetPipelineRequest{
			PipelineId: getPipelineResponse.PipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetPipelineResponse]{
				Info:    *getPipelineResponse,
				Timeout: i.Timeout,
			})
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case GetPipelineResponseStateRunning: // target state
			return getPipelineResponse, nil
		case GetPipelineResponseStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetPipelineResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get a pipeline
func (a *PipelinesAPI) GetPipelineByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error) {
	return a.GetPipeline(ctx, GetPipelineRequest{
		PipelineId: pipelineId,
	})
}

func (a *PipelinesAPI) GetPipelineByPipelineIdAndWait(ctx context.Context, pipelineId string, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	return a.GetPipelineAndWait(ctx, GetPipelineRequest{
		PipelineId: pipelineId,
	}, options...)
}

// Get a pipeline update
//
// Gets an update from an active pipeline.
func (a *PipelinesAPI) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {
	var getUpdateResponse GetUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates/%v", request.PipelineId, request.UpdateId)
	err := a.client.Get(ctx, path, request, &getUpdateResponse)
	return &getUpdateResponse, err
}

// Get a pipeline update
//
// Gets an update from an active pipeline.
func (a *PipelinesAPI) GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error) {
	return a.GetUpdate(ctx, GetUpdateRequest{
		PipelineId: pipelineId,
		UpdateId:   updateId,
	})
}

// List pipeline updates
//
// List updates for an active pipeline.
func (a *PipelinesAPI) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	var listUpdatesResponse ListUpdatesResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	err := a.client.Get(ctx, path, request, &listUpdatesResponse)
	return &listUpdatesResponse, err
}

// List pipeline updates
//
// List updates for an active pipeline.
func (a *PipelinesAPI) ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error) {
	return a.ListUpdates(ctx, ListUpdatesRequest{
		PipelineId: pipelineId,
	})
}

// Reset a pipeline
//
// Resets a pipeline.
func (a *PipelinesAPI) ResetPipeline(ctx context.Context, request ResetPipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v/reset", request.PipelineId)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// ResetPipeline and wait to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
func (a *PipelinesAPI) ResetPipelineAndWait(ctx context.Context, resetPipelineRequest ResetPipelineRequest, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.ResetPipeline(ctx, resetPipelineRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetPipelineResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetPipelineResponse](ctx, i.Timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := a.GetPipeline(ctx, GetPipelineRequest{
			PipelineId: resetPipelineRequest.PipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetPipelineResponse]{
				Info:    *getPipelineResponse,
				Timeout: i.Timeout,
			})
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case GetPipelineResponseStateRunning: // target state
			return getPipelineResponse, nil
		case GetPipelineResponseStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetPipelineResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Queue a pipeline update
//
// Starts or queues a pipeline update.
func (a *PipelinesAPI) StartUpdate(ctx context.Context, request StartUpdateRequest) (*StartUpdateResponse, error) {
	var startUpdateResponse StartUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	err := a.client.Post(ctx, path, request, &startUpdateResponse)
	return &startUpdateResponse, err
}

// Stop a pipeline
//
// Stops a pipeline.
func (a *PipelinesAPI) StopPipeline(ctx context.Context, request StopPipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v/stop", request.PipelineId)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// StopPipeline and wait to reach IDLE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
func (a *PipelinesAPI) StopPipelineAndWait(ctx context.Context, stopPipelineRequest StopPipelineRequest, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.StopPipeline(ctx, stopPipelineRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetPipelineResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetPipelineResponse](ctx, i.Timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := a.GetPipeline(ctx, GetPipelineRequest{
			PipelineId: stopPipelineRequest.PipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetPipelineResponse]{
				Info:    *getPipelineResponse,
				Timeout: i.Timeout,
			})
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case GetPipelineResponseStateIdle: // target state
			return getPipelineResponse, nil
		case GetPipelineResponseStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetPipelineResponseStateIdle, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Edit a pipeline
//
// Updates a pipeline with the supplied configuration.
func (a *PipelinesAPI) UpdatePipeline(ctx context.Context, request EditPipelineRequest) error {
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	err := a.client.Put(ctx, path, request)
	return err
}
