// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just Pipelines API methods
type pipelinesImpl struct {
	client *httpclient.ApiClient
}

func (a *pipelinesImpl) Create(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error) {
	var createPipelineResponse CreatePipelineResponse
	path := "/api/2.0/pipelines"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &createPipelineResponse)
	return &createPipelineResponse, err
}

func (a *pipelinesImpl) Delete(ctx context.Context, request DeletePipelineRequest) (*DeletePipelineResponse, error) {
	var deletePipelineResponse DeletePipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodDelete, path, headers, queryParams, request, &deletePipelineResponse)
	return &deletePipelineResponse, err
}

func (a *pipelinesImpl) Get(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error) {
	var getPipelineResponse GetPipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &getPipelineResponse)
	return &getPipelineResponse, err
}

func (a *pipelinesImpl) GetPermissionLevels(ctx context.Context, request GetPipelinePermissionLevelsRequest) (*GetPipelinePermissionLevelsResponse, error) {
	var getPipelinePermissionLevelsResponse GetPipelinePermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v/permissionLevels", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &getPipelinePermissionLevelsResponse)
	return &getPipelinePermissionLevelsResponse, err
}

func (a *pipelinesImpl) GetPermissions(ctx context.Context, request GetPipelinePermissionsRequest) (*PipelinePermissions, error) {
	var pipelinePermissions PipelinePermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &pipelinePermissions)
	return &pipelinePermissions, err
}

func (a *pipelinesImpl) GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error) {
	var getUpdateResponse GetUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates/%v", request.PipelineId, request.UpdateId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &getUpdateResponse)
	return &getUpdateResponse, err
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
	return listing.ToSliceN[PipelineEvent, int](ctx, iterator, request.MaxResults)

}
func (a *pipelinesImpl) internalListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) (*ListPipelineEventsResponse, error) {
	var listPipelineEventsResponse ListPipelineEventsResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/events", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listPipelineEventsResponse)
	return &listPipelineEventsResponse, err
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
	return listing.ToSliceN[PipelineStateInfo, int](ctx, iterator, request.MaxResults)

}
func (a *pipelinesImpl) internalListPipelines(ctx context.Context, request ListPipelinesRequest) (*ListPipelinesResponse, error) {
	var listPipelinesResponse ListPipelinesResponse
	path := "/api/2.0/pipelines"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listPipelinesResponse)
	return &listPipelinesResponse, err
}

func (a *pipelinesImpl) ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error) {
	var listUpdatesResponse ListUpdatesResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listUpdatesResponse)
	return &listUpdatesResponse, err
}

func (a *pipelinesImpl) SetPermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {
	var pipelinePermissions PipelinePermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPut, path, headers, queryParams, request, &pipelinePermissions)
	return &pipelinePermissions, err
}

func (a *pipelinesImpl) StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error) {
	var startUpdateResponse StartUpdateResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/updates", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &startUpdateResponse)
	return &startUpdateResponse, err
}

func (a *pipelinesImpl) Stop(ctx context.Context, request StopRequest) (*StopPipelineResponse, error) {
	var stopPipelineResponse StopPipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v/stop", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, nil, &stopPipelineResponse)
	return &stopPipelineResponse, err
}

func (a *pipelinesImpl) Update(ctx context.Context, request EditPipeline) (*EditPipelineResponse, error) {
	var editPipelineResponse EditPipelineResponse
	path := fmt.Sprintf("/api/2.0/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPut, path, headers, queryParams, request, &editPipelineResponse)
	return &editPipelineResponse, err
}

func (a *pipelinesImpl) UpdatePermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error) {
	var pipelinePermissions PipelinePermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPatch, path, headers, queryParams, request, &pipelinePermissions)
	return &pipelinePermissions, err
}

func do(
	client *httpclient.ApiClient,
	ctx context.Context,
	method string,
	path string,
	headers map[string]string,
	queryParams map[string]any,
	request any,
	response any,
	visitors ...func(*http.Request) error,
) error {
	opts := []httpclient.DoOption{}
	for _, v := range visitors {
		opts = append(opts, httpclient.WithRequestVisitor(v))
	}
	opts = append(opts, httpclient.WithQueryParameters(queryParams))
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	opts = append(opts, httpclient.WithRequestData(request))
	opts = append(opts, httpclient.WithResponseUnmarshal(response))

	// Remove extra `/` from path for files API. Once the OpenAPI spec doesn't
	// include the extra slash, we can remove this
	if strings.HasPrefix(path, "/api/2.0/fs/files//") {
		path = strings.Replace(path, "/api/2.0/fs/files//", "/api/2.0/fs/files/", 1)
	}

	return client.Do(ctx, method, path, opts...)
}
