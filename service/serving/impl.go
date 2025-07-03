// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/dataplane"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just ServingEndpoints API methods
type servingEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *servingEndpointsImpl) BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error) {

	requestPb, pbErr := buildLogsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var buildLogsResponsePb buildLogsResponsePb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/build-logs", requestPb.Name, requestPb.ServedModelName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&buildLogsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := buildLogsResponseFromPb(&buildLogsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) Create(ctx context.Context, request CreateServingEndpoint) (*ServingEndpointDetailed, error) {

	requestPb, pbErr := createServingEndpointToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servingEndpointDetailedPb servingEndpointDetailedPb
	path := "/api/2.0/serving-endpoints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&servingEndpointDetailedPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := servingEndpointDetailedFromPb(&servingEndpointDetailedPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) CreateProvisionedThroughputEndpoint(ctx context.Context, request CreatePtEndpointRequest) (*ServingEndpointDetailed, error) {

	requestPb, pbErr := createPtEndpointRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servingEndpointDetailedPb servingEndpointDetailedPb
	path := "/api/2.0/serving-endpoints/pt"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&servingEndpointDetailedPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := servingEndpointDetailedFromPb(&servingEndpointDetailedPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) Delete(ctx context.Context, request DeleteServingEndpointRequest) error {

	requestPb, pbErr := deleteServingEndpointRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *servingEndpointsImpl) ExportMetrics(ctx context.Context, request ExportMetricsRequest) (*ExportMetricsResponse, error) {

	requestPb, pbErr := exportMetricsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var exportMetricsResponsePb exportMetricsResponsePb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/metrics", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "text/plain"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&exportMetricsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := exportMetricsResponseFromPb(&exportMetricsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error) {

	requestPb, pbErr := getServingEndpointRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servingEndpointDetailedPb servingEndpointDetailedPb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&servingEndpointDetailedPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := servingEndpointDetailedFromPb(&servingEndpointDetailedPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) GetOpenApi(ctx context.Context, request GetOpenApiRequest) (*GetOpenApiResponse, error) {

	requestPb, pbErr := getOpenApiRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getOpenApiResponsePb getOpenApiResponsePb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/openapi", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "text/plain"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getOpenApiResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getOpenApiResponseFromPb(&getOpenApiResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) GetPermissionLevels(ctx context.Context, request GetServingEndpointPermissionLevelsRequest) (*GetServingEndpointPermissionLevelsResponse, error) {

	requestPb, pbErr := getServingEndpointPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getServingEndpointPermissionLevelsResponsePb getServingEndpointPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v/permissionLevels", requestPb.ServingEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getServingEndpointPermissionLevelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getServingEndpointPermissionLevelsResponseFromPb(&getServingEndpointPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) GetPermissions(ctx context.Context, request GetServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {

	requestPb, pbErr := getServingEndpointPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servingEndpointPermissionsPb servingEndpointPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", requestPb.ServingEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&servingEndpointPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := servingEndpointPermissionsFromPb(&servingEndpointPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) HttpRequest(ctx context.Context, request ExternalFunctionRequest) (*HttpRequestResponse, error) {

	requestPb, pbErr := externalFunctionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var httpRequestResponsePb httpRequestResponsePb
	path := "/api/2.0/external-function"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "text/plain"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&httpRequestResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := httpRequestResponseFromPb(&httpRequestResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Get all serving endpoints.
func (a *servingEndpointsImpl) List(ctx context.Context) listing.Iterator[ServingEndpoint] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListEndpointsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListEndpointsResponse) []ServingEndpoint {
		return resp.Endpoints
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get all serving endpoints.
func (a *servingEndpointsImpl) ListAll(ctx context.Context) ([]ServingEndpoint, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[ServingEndpoint](ctx, iterator)
}

func (a *servingEndpointsImpl) internalList(ctx context.Context) (*ListEndpointsResponse, error) {

	var listEndpointsResponsePb listEndpointsResponsePb
	path := "/api/2.0/serving-endpoints"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&listEndpointsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listEndpointsResponseFromPb(&listEndpointsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) Logs(ctx context.Context, request LogsRequest) (*ServerLogsResponse, error) {

	requestPb, pbErr := logsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var serverLogsResponsePb serverLogsResponsePb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/logs", requestPb.Name, requestPb.ServedModelName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&serverLogsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := serverLogsResponseFromPb(&serverLogsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) Patch(ctx context.Context, request PatchServingEndpointTags) (*EndpointTags, error) {

	requestPb, pbErr := patchServingEndpointTagsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var endpointTagsPb endpointTagsPb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/tags", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&endpointTagsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := endpointTagsFromPb(&endpointTagsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) Put(ctx context.Context, request PutRequest) (*PutResponse, error) {

	requestPb, pbErr := putRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var putResponsePb putResponsePb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/rate-limits", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&putResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := putResponseFromPb(&putResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) PutAiGateway(ctx context.Context, request PutAiGatewayRequest) (*PutAiGatewayResponse, error) {

	requestPb, pbErr := putAiGatewayRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var putAiGatewayResponsePb putAiGatewayResponsePb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/ai-gateway", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&putAiGatewayResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := putAiGatewayResponseFromPb(&putAiGatewayResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error) {

	requestPb, pbErr := queryEndpointInputToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryEndpointResponsePb queryEndpointResponsePb
	path := fmt.Sprintf("/serving-endpoints/%v/invocations", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&queryEndpointResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := queryEndpointResponseFromPb(&queryEndpointResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) SetPermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {

	requestPb, pbErr := servingEndpointPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servingEndpointPermissionsPb servingEndpointPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", requestPb.ServingEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&servingEndpointPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := servingEndpointPermissionsFromPb(&servingEndpointPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) UpdateConfig(ctx context.Context, request EndpointCoreConfigInput) (*ServingEndpointDetailed, error) {

	requestPb, pbErr := endpointCoreConfigInputToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servingEndpointDetailedPb servingEndpointDetailedPb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/config", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&servingEndpointDetailedPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := servingEndpointDetailedFromPb(&servingEndpointDetailedPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) UpdatePermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {

	requestPb, pbErr := servingEndpointPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servingEndpointPermissionsPb servingEndpointPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", requestPb.ServingEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&servingEndpointPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := servingEndpointPermissionsFromPb(&servingEndpointPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *servingEndpointsImpl) UpdateProvisionedThroughputEndpointConfig(ctx context.Context, request UpdateProvisionedThroughputEndpointConfigRequest) (*ServingEndpointDetailed, error) {

	requestPb, pbErr := updateProvisionedThroughputEndpointConfigRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var servingEndpointDetailedPb servingEndpointDetailedPb
	path := fmt.Sprintf("/api/2.0/serving-endpoints/pt/%v/config", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&servingEndpointDetailedPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := servingEndpointDetailedFromPb(&servingEndpointDetailedPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ServingEndpointsDataPlane API methods
type servingEndpointsDataPlaneImpl struct {
	client       *client.DatabricksClient
	controlPlane *ServingEndpointsAPI
	dpts         dataplane.EndpointTokenSource
	infos        sync.Map
}

func (a *servingEndpointsDataPlaneImpl) dataPlaneInfoQuery(ctx context.Context, request QueryEndpointInput) (*DataPlaneInfo, error) {
	key := "Query/" + strings.Join([]string{
		fmt.Sprintf("%v", request.Name),
	}, "/")

	info, ok := a.infos.Load(key)
	if !ok {
		response, err := a.controlPlane.Get(ctx, GetServingEndpointRequest{
			Name: request.Name,
		})
		if err != nil {
			return nil, err
		}
		info = response.DataPlaneInfo.QueryInfo
		a.infos.Store(key, info)
	}
	return info.(*DataPlaneInfo), nil
}

func (a *servingEndpointsDataPlaneImpl) Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error) {

	requestPb, pbErr := queryEndpointInputToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	dpi, err := a.dataPlaneInfoQuery(ctx, request)
	if err != nil {
		return nil, err
	}
	dpt, err := a.dpts.Token(ctx, dpi.EndpointUrl, dpi.AuthorizationDetails)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	opts := []httpclient.DoOption{}
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	queryParams := make(map[string]any)
	opts = append(opts, httpclient.WithQueryParameters(queryParams))

	var queryEndpointResponsePb queryEndpointResponsePb
	opts = append(opts, httpclient.WithRequestData(*requestPb))
	opts = append(opts, httpclient.WithResponseUnmarshal(&queryEndpointResponsePb))
	opts = append(opts, httpclient.WithToken(dpt))
	err = a.client.ApiClient().Do(ctx, http.MethodPost, dpi.EndpointUrl, opts...)
	if err != nil {
		return nil, err
	}
	resp, err := queryEndpointResponseFromPb(&queryEndpointResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
