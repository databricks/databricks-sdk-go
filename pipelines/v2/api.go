// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Delta Live Tables API allows you to create, edit, delete, start, and view details about pipelines.
package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type pipelinesBaseClient struct {
	pipelinesImpl
}

// Delete a pipeline.
//
// Deletes a pipeline.
func (a *pipelinesBaseClient) DeleteByPipelineId(ctx context.Context, pipelineId string) (*DeletePipelineResponse, error) {
	return a.pipelinesImpl.Delete(ctx, DeletePipelineRequest{
		PipelineId: pipelineId,
	})
}

// Get a pipeline.
func (a *pipelinesBaseClient) GetByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error) {
	return a.pipelinesImpl.Get(ctx, GetPipelineRequest{
		PipelineId: pipelineId,
	})
}

// Get pipeline permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *pipelinesBaseClient) GetPermissionLevelsByPipelineId(ctx context.Context, pipelineId string) (*GetPipelinePermissionLevelsResponse, error) {
	return a.pipelinesImpl.GetPermissionLevels(ctx, GetPipelinePermissionLevelsRequest{
		PipelineId: pipelineId,
	})
}

// Get pipeline permissions.
//
// Gets the permissions of a pipeline. Pipelines can inherit permissions from
// their root object.
func (a *pipelinesBaseClient) GetPermissionsByPipelineId(ctx context.Context, pipelineId string) (*PipelinePermissions, error) {
	return a.pipelinesImpl.GetPermissions(ctx, GetPipelinePermissionsRequest{
		PipelineId: pipelineId,
	})
}

// Get a pipeline update.
//
// Gets an update from an active pipeline.
func (a *pipelinesBaseClient) GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error) {
	return a.pipelinesImpl.GetUpdate(ctx, GetUpdateRequest{
		PipelineId: pipelineId,
		UpdateId:   updateId,
	})
}

// List pipeline events.
//
// Retrieves events for a pipeline.
func (a *pipelinesBaseClient) ListPipelineEventsByPipelineId(ctx context.Context, pipelineId string) (*ListPipelineEventsResponse, error) {
	return a.pipelinesImpl.internalListPipelineEvents(ctx, ListPipelineEventsRequest{
		PipelineId: pipelineId,
	})
}

// PipelineStateInfoNameToPipelineIdMap calls [pipelinesBaseClient.ListPipelinesAll] and creates a map of results with [PipelineStateInfo].Name as key and [PipelineStateInfo].PipelineId as value.
//
// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
//
// Note: All [PipelineStateInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *pipelinesBaseClient) PipelineStateInfoNameToPipelineIdMap(ctx context.Context, request ListPipelinesRequest) (map[string]string, error) {
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

// GetByName calls [pipelinesBaseClient.PipelineStateInfoNameToPipelineIdMap] and returns a single [PipelineStateInfo].
//
// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
//
// Note: All [PipelineStateInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *pipelinesBaseClient) GetByName(ctx context.Context, name string) (*PipelineStateInfo, error) {
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
func (a *pipelinesBaseClient) ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error) {
	return a.pipelinesImpl.ListUpdates(ctx, ListUpdatesRequest{
		PipelineId: pipelineId,
	})
}

// Stop a pipeline.
//
// Stops the pipeline by canceling the active update. If there is no active
// update for the pipeline, this request is a no-op.
func (a *pipelinesBaseClient) Stop(ctx context.Context, stopRequest StopRequest) (*PipelinesStopWaiter, error) {
	stopPipelineResponse, err := a.pipelinesImpl.Stop(ctx, stopRequest)
	if err != nil {
		return nil, err
	}
	return &PipelinesStopWaiter{
		RawResponse: stopPipelineResponse,
		pipelineId:  stopRequest.PipelineId,
		service:     a,
	}, nil
}

type PipelinesStopWaiter struct {
	// RawResponse is the raw response of the Stop call.
	RawResponse *StopPipelineResponse
	service     *pipelinesBaseClient
	pipelineId  string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *PipelinesStopWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[GetPipelineResponse](ctx, opts.Timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := w.service.Get(ctx, GetPipelineRequest{
			PipelineId: w.pipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := getPipelineResponse.State
		statusMessage := fmt.Sprintf("current status: %s", status)
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
