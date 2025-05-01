// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Pipelines API methods
type pipelinesImpl struct {
	client *client.DatabricksClient
}

func (a *pipelinesImpl) Create(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error) {

	requestPb, pbErr := createPipelineToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createPipelineResponsePb createPipelineResponsePb
	path := "/api/2.0/pipelines"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createPipelineResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := createPipelineResponseFromPb(&createPipelineResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) Delete(ctx context.Context, request DeletePipelineRequest) error {

	requestPb, pbErr := deletePipelineRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deletePipelineResponsePb deletePipelineResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, (*requestPb), &deletePipelineResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *pipelinesImpl) Get(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {

	requestPb, pbErr := getPipelineRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPipelineResponsePb getPipelineResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getPipelineResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getPipelineResponseFromPb(&getPipelineResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) GetPermissionLevels(ctx context.Context, request GetPipelinePermissionLevelsRequest) (*GetPipelinePermissionLevelsResponse, error) {

	requestPb, pbErr := getPipelinePermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPipelinePermissionLevelsResponsePb getPipelinePermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v/permissionLevels", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getPipelinePermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getPipelinePermissionLevelsResponseFromPb(&getPipelinePermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) GetPermissions(ctx context.Context, request GetPipelinePermissionsRequest) (*PipelinePermissions, error) {

	requestPb, pbErr := getPipelinePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var pipelinePermissionsPb pipelinePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := pipelinePermissionsFromPb(&pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {

	requestPb, pbErr := getUpdateRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getUpdateResponsePb getUpdateResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates/%v", requestPb.PipelineId, requestPb.UpdateId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getUpdateResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getUpdateResponseFromPb(&getUpdateResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List pipeline events.
//
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

// List pipeline events.
//
// Retrieves events for a pipeline.
func (a *pipelinesImpl) ListPipelineEventsAll(ctx context.Context, request ListPipelineEventsRequest) ([]PipelineEvent, error) {
	iterator := a.ListPipelineEvents(ctx, request)
	return listing.ToSlice[PipelineEvent](ctx, iterator)
}

func (a *pipelinesImpl) internalListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) (*ListPipelineEventsResponse, error) {

	requestPb, pbErr := listPipelineEventsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listPipelineEventsResponsePb listPipelineEventsResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/events", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listPipelineEventsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listPipelineEventsResponseFromPb(&listPipelineEventsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List pipelines.
//
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

// List pipelines.
//
// Lists pipelines defined in the Delta Live Tables system.
func (a *pipelinesImpl) ListPipelinesAll(ctx context.Context, request ListPipelinesRequest) ([]PipelineStateInfo, error) {
	iterator := a.ListPipelines(ctx, request)
	return listing.ToSlice[PipelineStateInfo](ctx, iterator)
}

func (a *pipelinesImpl) internalListPipelines(ctx context.Context, request ListPipelinesRequest) (*ListPipelinesResponse, error) {

	requestPb, pbErr := listPipelinesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listPipelinesResponsePb listPipelinesResponsePb
	path := "/api/2.0/pipelines"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listPipelinesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listPipelinesResponseFromPb(&listPipelinesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {

	requestPb, pbErr := listUpdatesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listUpdatesResponsePb listUpdatesResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listUpdatesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listUpdatesResponseFromPb(&listUpdatesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) SetPermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {

	requestPb, pbErr := pipelinePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var pipelinePermissionsPb pipelinePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, (*requestPb), &pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := pipelinePermissionsFromPb(&pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error) {

	requestPb, pbErr := startUpdateToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var startUpdateResponsePb startUpdateResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &startUpdateResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := startUpdateResponseFromPb(&startUpdateResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *pipelinesImpl) Stop(ctx context.Context, request StopRequest) error {

	requestPb, pbErr := stopRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var stopPipelineResponsePb stopPipelineResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v/stop", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &stopPipelineResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *pipelinesImpl) Update(ctx context.Context, request EditPipeline) error {

	requestPb, pbErr := editPipelineToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var editPipelineResponsePb editPipelineResponsePb
	path := fmt.Sprintf("/api/2.0/pipelines/%v", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, (*requestPb), &editPipelineResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *pipelinesImpl) UpdatePermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {

	requestPb, pbErr := pipelinePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var pipelinePermissionsPb pipelinePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", requestPb.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := pipelinePermissionsFromPb(&pipelinePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
