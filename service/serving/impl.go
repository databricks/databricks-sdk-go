// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just ServingEndpoints API methods
type servingEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *servingEndpointsImpl) BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error) {
	var buildLogsResponse BuildLogsResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/build-logs", request.Name, request.ServedModelName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &buildLogsResponse)
	return &buildLogsResponse, err
}

func (a *servingEndpointsImpl) Create(ctx context.Context, request CreateServingEndpoint) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := "/api/2.0/serving-endpoints"
	err := a.client.Do(ctx, http.MethodPost, path, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) Delete(ctx context.Context, request DeleteServingEndpointRequest) error {
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *servingEndpointsImpl) ExportMetrics(ctx context.Context, request ExportMetricsRequest) error {
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/metrics", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, nil)
	return err
}

func (a *servingEndpointsImpl) Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) List(ctx context.Context) (*ListEndpointsResponse, error) {
	var listEndpointsResponse ListEndpointsResponse
	path := "/api/2.0/serving-endpoints"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listEndpointsResponse)
	return &listEndpointsResponse, err
}

func (a *servingEndpointsImpl) Logs(ctx context.Context, request LogsRequest) (*ServerLogsResponse, error) {
	var serverLogsResponse ServerLogsResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/logs", request.Name, request.ServedModelName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &serverLogsResponse)
	return &serverLogsResponse, err
}

func (a *servingEndpointsImpl) Query(ctx context.Context, request QueryRequest) (*QueryEndpointResponse, error) {
	var queryEndpointResponse QueryEndpointResponse
	path := fmt.Sprintf("/serving-endpoints/%v/invocations", request.Name)
	err := a.client.Do(ctx, http.MethodPost, path, request, &queryEndpointResponse)
	return &queryEndpointResponse, err
}

func (a *servingEndpointsImpl) UpdateConfig(ctx context.Context, request EndpointCoreConfigInput) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/config", request.Name)
	err := a.client.Do(ctx, http.MethodPut, path, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}
