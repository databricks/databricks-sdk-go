// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewPipelines(client *client.DatabricksClient) *PipelinesAPI {
	return &PipelinesAPI{
		impl: &pipelinesImpl{
			client: client,
		},
	}
}

// The Delta Live Tables API allows you to create, edit, delete, start, and view
// details about pipelines.
//
// Delta Live Tables is a framework for building reliable, maintainable, and
// testable data processing pipelines. You define the transformations to perform
// on your data, and Delta Live Tables manages task orchestration, cluster
// management, monitoring, data quality, and error handling.
//
// Instead of defining your data pipelines using a series of separate Apache
// Spark tasks, Delta Live Tables manages how your data is transformed based on
// a target schema you define for each processing step. You can also enforce
// data quality with Delta Live Tables expectations. Expectations allow you to
// define expected data quality and specify how to handle records that fail
// those expectations.
type PipelinesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PipelinesService)
	impl PipelinesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PipelinesAPI) WithImpl(impl PipelinesService) *PipelinesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Pipelines API implementation
func (a *PipelinesAPI) Impl() PipelinesService {
	return a.impl
}

// Create a pipeline
//
// Creates a new data processing pipeline based on the requested configuration.
// If successful, this method returns the ID of the new pipeline.
func (a *PipelinesAPI) CreatePipeline(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error) {
	return a.impl.CreatePipeline(ctx, request)
}

// Calls [PipelinesAPI.CreatePipeline] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
func (a *PipelinesAPI) CreatePipelineAndWait(ctx context.Context, createPipeline CreatePipeline, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	createPipelineResponse, err := a.CreatePipeline(ctx, createPipeline)
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
		case PipelineStateRunning: // target state
			return getPipelineResponse, nil
		case PipelineStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				PipelineStateRunning, status, statusMessage)
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
	return a.impl.DeletePipeline(ctx, request)
}

// Delete a pipeline
//
// Deletes a pipeline.
func (a *PipelinesAPI) DeletePipelineByPipelineId(ctx context.Context, pipelineId string) error {
	return a.impl.DeletePipeline(ctx, DeletePipelineRequest{
		PipelineId: pipelineId,
	})
}

// Get a pipeline
func (a *PipelinesAPI) GetPipeline(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	return a.impl.GetPipeline(ctx, request)
}

// Calls [PipelinesAPI.GetPipeline] and waits to reach RUNNING state
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
		case PipelineStateRunning: // target state
			return getPipelineResponse, nil
		case PipelineStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				PipelineStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get a pipeline
func (a *PipelinesAPI) GetPipelineByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error) {
	return a.impl.GetPipeline(ctx, GetPipelineRequest{
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
	return a.impl.GetUpdate(ctx, request)
}

// Get a pipeline update
//
// Gets an update from an active pipeline.
func (a *PipelinesAPI) GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error) {
	return a.impl.GetUpdate(ctx, GetUpdateRequest{
		PipelineId: pipelineId,
		UpdateId:   updateId,
	})
}

// List pipelines
//
// Lists pipelines defined in the Delta Live Tables system.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) ListPipelinesAll(ctx context.Context, request ListPipelinesRequest) ([]PipelineStateInfo, error) {
	var results []PipelineStateInfo
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.ListPipelines(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Statuses) == 0 {
			break
		}
		for _, v := range response.Statuses {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// PipelineStateInfoNameToPipelineIdMap calls [PipelinesAPI.ListPipelinesAll] and creates a map of results with [PipelineStateInfo].Name as key and [PipelineStateInfo].PipelineId as value.
//
// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
//
// Note: All [PipelineStateInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) PipelineStateInfoNameToPipelineIdMap(ctx context.Context, request ListPipelinesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListPipelinesAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.PipelineId
	}
	return mapping, nil
}

// GetPipelineStateInfoByName calls [PipelinesAPI.PipelineStateInfoNameToPipelineIdMap] and returns a single [PipelineStateInfo].
//
// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
//
// Note: All [PipelineStateInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) GetPipelineStateInfoByName(ctx context.Context, name string) (*PipelineStateInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListPipelinesAll(ctx, ListPipelinesRequest{})
	if err != nil {
		return nil, err
	}
	duplicates := map[string]bool{}
	for _, v := range result {
		key := v.Name
		if duplicates[key] {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		if key != name {
			continue
		}
		duplicates[key] = true
		return &v, nil
	}
	return nil, fmt.Errorf("PipelineStateInfo named '%s' does not exist", name)
}

// List pipeline updates
//
// List updates for an active pipeline.
func (a *PipelinesAPI) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	return a.impl.ListUpdates(ctx, request)
}

// List pipeline updates
//
// List updates for an active pipeline.
func (a *PipelinesAPI) ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error) {
	return a.impl.ListUpdates(ctx, ListUpdatesRequest{
		PipelineId: pipelineId,
	})
}

// Reset a pipeline
//
// Resets a pipeline.
func (a *PipelinesAPI) ResetPipeline(ctx context.Context, request ResetPipelineRequest) error {
	return a.impl.ResetPipeline(ctx, request)
}

// Calls [PipelinesAPI.ResetPipeline] and waits to reach RUNNING state
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
		case PipelineStateRunning: // target state
			return getPipelineResponse, nil
		case PipelineStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				PipelineStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Queue a pipeline update
//
// Starts or queues a pipeline update.
func (a *PipelinesAPI) StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error) {
	return a.impl.StartUpdate(ctx, request)
}

// Stop a pipeline
//
// Stops a pipeline.
func (a *PipelinesAPI) StopPipeline(ctx context.Context, request StopPipelineRequest) error {
	return a.impl.StopPipeline(ctx, request)
}

// Calls [PipelinesAPI.StopPipeline] and waits to reach IDLE state
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
		case PipelineStateIdle: // target state
			return getPipelineResponse, nil
		case PipelineStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				PipelineStateIdle, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Edit a pipeline
//
// Updates a pipeline with the supplied configuration.
func (a *PipelinesAPI) UpdatePipeline(ctx context.Context, request EditPipeline) error {
	return a.impl.UpdatePipeline(ctx, request)
}
