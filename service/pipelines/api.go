// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Delta Live Tables API allows you to create, edit, delete, start, and view details about pipelines.
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

// WaitGetPipelineIdle repeatedly calls [PipelinesAPI.Get] and waits to reach IDLE state
func (a *PipelinesAPI) WaitGetPipelineIdle(ctx context.Context, pipelineId string,
	timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[GetPipelineResponse](ctx, timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := a.Get(ctx, GetPipelineRequest{
			PipelineId: pipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(getPipelineResponse)
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

// WaitGetPipelineIdle is a wrapper that calls [PipelinesAPI.WaitGetPipelineIdle] and waits to reach IDLE state.
type WaitGetPipelineIdle[R any] struct {
	Response   *R
	PipelineId string `json:"pipeline_id"`
	poll       func(time.Duration, func(*GetPipelineResponse)) (*GetPipelineResponse, error)
	callback   func(*GetPipelineResponse)
	timeout    time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetPipelineIdle[R]) OnProgress(callback func(*GetPipelineResponse)) *WaitGetPipelineIdle[R] {
	w.callback = callback
	return w
}

// Get the GetPipelineResponse with the default timeout of 20 minutes.
func (w *WaitGetPipelineIdle[R]) Get() (*GetPipelineResponse, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the GetPipelineResponse with custom timeout.
func (w *WaitGetPipelineIdle[R]) GetWithTimeout(timeout time.Duration) (*GetPipelineResponse, error) {
	return w.poll(timeout, w.callback)
}

// WaitGetPipelineRunning repeatedly calls [PipelinesAPI.Get] and waits to reach RUNNING state
func (a *PipelinesAPI) WaitGetPipelineRunning(ctx context.Context, pipelineId string,
	timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[GetPipelineResponse](ctx, timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := a.Get(ctx, GetPipelineRequest{
			PipelineId: pipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(getPipelineResponse)
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

// WaitGetPipelineRunning is a wrapper that calls [PipelinesAPI.WaitGetPipelineRunning] and waits to reach RUNNING state.
type WaitGetPipelineRunning[R any] struct {
	Response   *R
	PipelineId string `json:"pipeline_id"`
	poll       func(time.Duration, func(*GetPipelineResponse)) (*GetPipelineResponse, error)
	callback   func(*GetPipelineResponse)
	timeout    time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetPipelineRunning[R]) OnProgress(callback func(*GetPipelineResponse)) *WaitGetPipelineRunning[R] {
	w.callback = callback
	return w
}

// Get the GetPipelineResponse with the default timeout of 20 minutes.
func (w *WaitGetPipelineRunning[R]) Get() (*GetPipelineResponse, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the GetPipelineResponse with custom timeout.
func (w *WaitGetPipelineRunning[R]) GetWithTimeout(timeout time.Duration) (*GetPipelineResponse, error) {
	return w.poll(timeout, w.callback)
}

// Create a pipeline.
//
// Creates a new data processing pipeline based on the requested configuration.
// If successful, this method returns the ID of the new pipeline.
func (a *PipelinesAPI) Create(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a pipeline.
//
// Deletes a pipeline.
func (a *PipelinesAPI) Delete(ctx context.Context, request DeletePipelineRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a pipeline.
//
// Deletes a pipeline.
func (a *PipelinesAPI) DeleteByPipelineId(ctx context.Context, pipelineId string) error {
	return a.impl.Delete(ctx, DeletePipelineRequest{
		PipelineId: pipelineId,
	})
}

// Get a pipeline.
func (a *PipelinesAPI) Get(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	return a.impl.Get(ctx, request)
}

// Get a pipeline.
func (a *PipelinesAPI) GetByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error) {
	return a.impl.Get(ctx, GetPipelineRequest{
		PipelineId: pipelineId,
	})
}

// Get pipeline permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *PipelinesAPI) GetPermissionLevels(ctx context.Context, request GetPipelinePermissionLevelsRequest) (*GetPipelinePermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get pipeline permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *PipelinesAPI) GetPermissionLevelsByPipelineId(ctx context.Context, pipelineId string) (*GetPipelinePermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPipelinePermissionLevelsRequest{
		PipelineId: pipelineId,
	})
}

// Get pipeline permissions.
//
// Gets the permissions of a pipeline. Pipelines can inherit permissions from
// their root object.
func (a *PipelinesAPI) GetPermissions(ctx context.Context, request GetPipelinePermissionsRequest) (*PipelinePermissions, error) {
	return a.impl.GetPermissions(ctx, request)
}

// Get pipeline permissions.
//
// Gets the permissions of a pipeline. Pipelines can inherit permissions from
// their root object.
func (a *PipelinesAPI) GetPermissionsByPipelineId(ctx context.Context, pipelineId string) (*PipelinePermissions, error) {
	return a.impl.GetPermissions(ctx, GetPipelinePermissionsRequest{
		PipelineId: pipelineId,
	})
}

// Get a pipeline update.
//
// Gets an update from an active pipeline.
func (a *PipelinesAPI) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {
	return a.impl.GetUpdate(ctx, request)
}

// Get a pipeline update.
//
// Gets an update from an active pipeline.
func (a *PipelinesAPI) GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error) {
	return a.impl.GetUpdate(ctx, GetUpdateRequest{
		PipelineId: pipelineId,
		UpdateId:   updateId,
	})
}

// List pipeline events.
//
// Retrieves events for a pipeline.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) ListPipelineEventsAll(ctx context.Context, request ListPipelineEventsRequest) ([]PipelineEvent, error) {
	var results []PipelineEvent
	var totalCount int = 0
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.ListPipelineEvents(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Events) == 0 {
			break
		}
		for _, v := range response.Events {
			results = append(results, v)
		}
		count := int(len(response.Events))
		totalCount += count
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
		limit := request.MaxResults
		if limit > 0 && totalCount >= limit {
			break
		}
	}
	return results, nil
}

// List pipeline events.
//
// Retrieves events for a pipeline.
func (a *PipelinesAPI) ListPipelineEventsByPipelineId(ctx context.Context, pipelineId string) (*ListPipelineEventsResponse, error) {
	return a.impl.ListPipelineEvents(ctx, ListPipelineEventsRequest{
		PipelineId: pipelineId,
	})
}

// List pipelines.
//
// Lists pipelines defined in the Delta Live Tables system.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) ListPipelinesAll(ctx context.Context, request ListPipelinesRequest) ([]PipelineStateInfo, error) {
	var results []PipelineStateInfo
	var totalCount int = 0
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
		count := int(len(response.Statuses))
		totalCount += count
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
		limit := request.MaxResults
		if limit > 0 && totalCount >= limit {
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

// GetByName calls [PipelinesAPI.PipelineStateInfoNameToPipelineIdMap] and returns a single [PipelineStateInfo].
//
// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
//
// Note: All [PipelineStateInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) GetByName(ctx context.Context, name string) (*PipelineStateInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListPipelinesAll(ctx, ListPipelinesRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]PipelineStateInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("PipelineStateInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of PipelineStateInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// List pipeline updates.
//
// List updates for an active pipeline.
func (a *PipelinesAPI) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	return a.impl.ListUpdates(ctx, request)
}

// List pipeline updates.
//
// List updates for an active pipeline.
func (a *PipelinesAPI) ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error) {
	return a.impl.ListUpdates(ctx, ListUpdatesRequest{
		PipelineId: pipelineId,
	})
}

// Reset a pipeline.
//
// Resets a pipeline.
func (a *PipelinesAPI) Reset(ctx context.Context, resetRequest ResetRequest) (*WaitGetPipelineRunning[any], error) {
	err := a.impl.Reset(ctx, resetRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetPipelineRunning[any]{

		PipelineId: resetRequest.PipelineId,
		poll: func(timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error) {
			return a.WaitGetPipelineRunning(ctx, resetRequest.PipelineId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [PipelinesAPI.Reset] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
//
// Deprecated: use [PipelinesAPI.Reset].Get() or [PipelinesAPI.WaitGetPipelineRunning]
func (a *PipelinesAPI) ResetAndWait(ctx context.Context, resetRequest ResetRequest, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	wait, err := a.Reset(ctx, resetRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GetPipelineResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GetPipelineResponse) {
		for _, o := range options {
			o(&retries.Info[GetPipelineResponse]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Set pipeline permissions.
//
// Sets permissions on a pipeline. Pipelines can inherit permissions from their
// root object.
func (a *PipelinesAPI) SetPermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {
	return a.impl.SetPermissions(ctx, request)
}

// Start a pipeline.
//
// Starts a new update for the pipeline. If there is already an active update
// for the pipeline, the request will fail and the active update will remain
// running.
func (a *PipelinesAPI) StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error) {
	return a.impl.StartUpdate(ctx, request)
}

// Stop a pipeline.
//
// Stops the pipeline by canceling the active update. If there is no active
// update for the pipeline, this request is a no-op.
func (a *PipelinesAPI) Stop(ctx context.Context, stopRequest StopRequest) (*WaitGetPipelineIdle[any], error) {
	err := a.impl.Stop(ctx, stopRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetPipelineIdle[any]{

		PipelineId: stopRequest.PipelineId,
		poll: func(timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error) {
			return a.WaitGetPipelineIdle(ctx, stopRequest.PipelineId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [PipelinesAPI.Stop] and waits to reach IDLE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
//
// Deprecated: use [PipelinesAPI.Stop].Get() or [PipelinesAPI.WaitGetPipelineIdle]
func (a *PipelinesAPI) StopAndWait(ctx context.Context, stopRequest StopRequest, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	wait, err := a.Stop(ctx, stopRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GetPipelineResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GetPipelineResponse) {
		for _, o := range options {
			o(&retries.Info[GetPipelineResponse]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Edit a pipeline.
//
// Updates a pipeline with the supplied configuration.
func (a *PipelinesAPI) Update(ctx context.Context, request EditPipeline) error {
	return a.impl.Update(ctx, request)
}

// Update pipeline permissions.
//
// Updates the permissions on a pipeline. Pipelines can inherit permissions from
// their root object.
func (a *PipelinesAPI) UpdatePermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {
	return a.impl.UpdatePermissions(ctx, request)
}
