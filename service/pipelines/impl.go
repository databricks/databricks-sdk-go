// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/pipelines/pipelinespb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Pipelines API methods
type pipelinesImpl struct {
	client *client.DatabricksClient
}

func (a *pipelinesImpl) Create(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error) {
	requestPb, pbErr := CreatePipelineToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createPipelineResponsePb pipelinespb.CreatePipelineResponsePb
	path := "/api/2.0/pipelines"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createPipelineResponsePb,
	)
	resp, err := CreatePipelineResponseFromPb(&createPipelineResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) Delete(ctx context.Context, request DeletePipelineRequest) error {
	requestPb, pbErr := DeletePipelineRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *pipelinesImpl) Get(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	requestPb, pbErr := GetPipelineRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPipelineResponsePb pipelinespb.GetPipelineResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getPipelineResponsePb,
	)
	resp, err := GetPipelineResponseFromPb(&getPipelineResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) GetPermissionLevels(ctx context.Context, request GetPipelinePermissionLevelsRequest) (*GetPipelinePermissionLevelsResponse, error) {
	requestPb, pbErr := GetPipelinePermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPipelinePermissionLevelsResponsePb pipelinespb.GetPipelinePermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v/permissionLevels", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getPipelinePermissionLevelsResponsePb,
	)
	resp, err := GetPipelinePermissionLevelsResponseFromPb(&getPipelinePermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) GetPermissions(ctx context.Context, request GetPipelinePermissionsRequest) (*PipelinePermissions, error) {
	requestPb, pbErr := GetPipelinePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var pipelinePermissionsPb pipelinespb.PipelinePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&pipelinePermissionsPb,
	)
	resp, err := PipelinePermissionsFromPb(&pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {
	requestPb, pbErr := GetUpdateRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getUpdateResponsePb pipelinespb.GetUpdateResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates/%v", request.PipelineId, request.UpdateId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getUpdateResponsePb,
	)
	resp, err := GetUpdateResponseFromPb(&getUpdateResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Retrieves events for a pipeline.
func (a *pipelinesImpl) ListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) listing.Iterator[PipelineEvent] {

	getNextPage := func(ctx context.Context, req ListPipelineEventsRequest) (*ListPipelineEventsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListPipelineEvents(ctx, req)
	}
	getItems := func(resp *ListPipelineEventsResponse) []PipelineEvent {
		return resp.Events
	}
	getNextReq := func(resp *ListPipelineEventsResponse) *ListPipelineEventsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Retrieves events for a pipeline.
func (a *pipelinesImpl) ListPipelineEventsAll(ctx context.Context, request ListPipelineEventsRequest) ([]PipelineEvent, error) {
	iterator := a.ListPipelineEvents(ctx, request)
	return listing.ToSlice[PipelineEvent](ctx, iterator)
}

func (a *pipelinesImpl) internalListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) (*ListPipelineEventsResponse, error) {
	requestPb, pbErr := ListPipelineEventsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listPipelineEventsResponsePb pipelinespb.ListPipelineEventsResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/events", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listPipelineEventsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListPipelineEventsResponseFromPb(&listPipelineEventsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists pipelines defined in the Delta Live Tables system.
func (a *pipelinesImpl) ListPipelines(ctx context.Context, request ListPipelinesRequest) listing.Iterator[PipelineStateInfo] {

	getNextPage := func(ctx context.Context, req ListPipelinesRequest) (*ListPipelinesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListPipelines(ctx, req)
	}
	getItems := func(resp *ListPipelinesResponse) []PipelineStateInfo {
		return resp.Statuses
	}
	getNextReq := func(resp *ListPipelinesResponse) *ListPipelinesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists pipelines defined in the Delta Live Tables system.
func (a *pipelinesImpl) ListPipelinesAll(ctx context.Context, request ListPipelinesRequest) ([]PipelineStateInfo, error) {
	iterator := a.ListPipelines(ctx, request)
	return listing.ToSlice[PipelineStateInfo](ctx, iterator)
}

func (a *pipelinesImpl) internalListPipelines(ctx context.Context, request ListPipelinesRequest) (*ListPipelinesResponse, error) {
	requestPb, pbErr := ListPipelinesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listPipelinesResponsePb pipelinespb.ListPipelinesResponsePb
	path := "/api/2.0/pipelines"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listPipelinesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListPipelinesResponseFromPb(&listPipelinesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	requestPb, pbErr := ListUpdatesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listUpdatesResponsePb pipelinespb.ListUpdatesResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listUpdatesResponsePb,
	)
	resp, err := ListUpdatesResponseFromPb(&listUpdatesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) SetPermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {
	requestPb, pbErr := PipelinePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var pipelinePermissionsPb pipelinespb.PipelinePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&pipelinePermissionsPb,
	)
	resp, err := PipelinePermissionsFromPb(&pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error) {
	requestPb, pbErr := StartUpdateToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var startUpdateResponsePb pipelinespb.StartUpdateResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&startUpdateResponsePb,
	)
	resp, err := StartUpdateResponseFromPb(&startUpdateResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) Stop(ctx context.Context, request StopRequest) error {

	path := fmt.Sprintf("/api/2.0/pipelines/%v/stop", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)

	return err
}

func (a *pipelinesImpl) Update(ctx context.Context, request EditPipeline) error {
	requestPb, pbErr := EditPipelineToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *pipelinesImpl) UpdatePermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {
	requestPb, pbErr := PipelinePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var pipelinePermissionsPb pipelinespb.PipelinePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&pipelinePermissionsPb,
	)
	resp, err := PipelinePermissionsFromPb(&pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
