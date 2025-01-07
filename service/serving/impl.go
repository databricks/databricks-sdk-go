// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/httpclient"
	goauth "golang.org/x/oauth2"
)

// unexported type that holds implementations of just ServingEndpoints API methods
type servingEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *servingEndpointsImpl) BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error) {
	var buildLogsResponse BuildLogsResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/build-logs", request.Name, request.ServedModelName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &buildLogsResponse)
	return &buildLogsResponse, err
}

func (a *servingEndpointsImpl) Create(ctx context.Context, request CreateServingEndpoint) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := "/api/2.0/serving-endpoints"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) Delete(ctx context.Context, request DeleteServingEndpointRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *servingEndpointsImpl) ExportMetrics(ctx context.Context, request ExportMetricsRequest) (*ExportMetricsResponse, error) {
	var exportMetricsResponse ExportMetricsResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/metrics", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "text/plain"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &exportMetricsResponse)
	return &exportMetricsResponse, err
}

func (a *servingEndpointsImpl) Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) GetOpenApi(ctx context.Context, request GetOpenApiRequest) error {
	var getOpenApiResponse GetOpenApiResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/openapi", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getOpenApiResponse)
	return err
}

func (a *servingEndpointsImpl) GetPermissionLevels(ctx context.Context, request GetServingEndpointPermissionLevelsRequest) (*GetServingEndpointPermissionLevelsResponse, error) {
	var getServingEndpointPermissionLevelsResponse GetServingEndpointPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v/permissionLevels", request.ServingEndpointId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getServingEndpointPermissionLevelsResponse)
	return &getServingEndpointPermissionLevelsResponse, err
}

func (a *servingEndpointsImpl) GetPermissions(ctx context.Context, request GetServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", request.ServingEndpointId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &servingEndpointPermissions)
	return &servingEndpointPermissions, err
}

func (a *servingEndpointsImpl) List(ctx context.Context) (*ListEndpointsResponse, error) {
	var listEndpointsResponse ListEndpointsResponse
	path := "/api/2.0/serving-endpoints"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listEndpointsResponse)
	return &listEndpointsResponse, err
}

func (a *servingEndpointsImpl) Logs(ctx context.Context, request LogsRequest) (*ServerLogsResponse, error) {
	var serverLogsResponse ServerLogsResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/logs", request.Name, request.ServedModelName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &serverLogsResponse)
	return &serverLogsResponse, err
}

func (a *servingEndpointsImpl) Patch(ctx context.Context, request PatchServingEndpointTags) ([]EndpointTag, error) {
	var endpointTagList []EndpointTag
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/tags", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &endpointTagList)
	return endpointTagList, err
}

func (a *servingEndpointsImpl) Put(ctx context.Context, request PutRequest) (*PutResponse, error) {
	var putResponse PutResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/rate-limits", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &putResponse)
	return &putResponse, err
}

func (a *servingEndpointsImpl) PutAiGateway(ctx context.Context, request PutAiGatewayRequest) (*PutAiGatewayResponse, error) {
	var putAiGatewayResponse PutAiGatewayResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/ai-gateway", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &putAiGatewayResponse)
	return &putAiGatewayResponse, err
}

func (a *servingEndpointsImpl) Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error) {
	var queryEndpointResponse QueryEndpointResponse
	path := fmt.Sprintf("/serving-endpoints/%v/invocations", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &queryEndpointResponse)
	return &queryEndpointResponse, err
}

func (a *servingEndpointsImpl) SetPermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", request.ServingEndpointId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &servingEndpointPermissions)
	return &servingEndpointPermissions, err
}

func (a *servingEndpointsImpl) UpdateConfig(ctx context.Context, request EndpointCoreConfigInput) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/config", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) UpdatePermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", request.ServingEndpointId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &servingEndpointPermissions)
	return &servingEndpointPermissions, err
}

// unexported type that holds implementations of just ServingEndpointsDataPlane API methods
type servingEndpointsDataPlaneImpl struct {
	dataPlaneService DataPlaneService
	controlPlane     *ServingEndpointsAPI
	client           *client.DatabricksClient
}

func (a *servingEndpointsDataPlaneImpl) Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error) {
	getRequest := GetServingEndpointRequest{
		Name: request.Name,
	}
	token, err := a.client.Config.GetToken()
	if err != nil {
		return nil, err
	}
	infoGetter := func() (*DataPlaneInfo, error) {
		response, err := a.controlPlane.Get(ctx, getRequest)
		if err != nil {
			return nil, err
		}
		if response.DataPlaneInfo == nil {
			return nil, errors.New("resource does not support direct Data Plane access")
		}
		return response.DataPlaneInfo.QueryInfo, nil
	}
	refresh := func(info *DataPlaneInfo) (*goauth.Token, error) {
		return a.client.GetOAuthToken(ctx, info.AuthorizationDetails, token)
	}
	getParams := []string{
		request.Name,
	}
	endpointUrl, dataPlaneToken, err := a.dataPlaneService.GetDataPlaneDetails("Query", getParams, refresh, infoGetter)
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	opts := []httpclient.DoOption{}
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	var queryEndpointResponse QueryEndpointResponse
	opts = append(opts, httpclient.WithRequestData(request))
	opts = append(opts, httpclient.WithResponseUnmarshal(&queryEndpointResponse))
	opts = append(opts, httpclient.WithToken(dataPlaneToken))
	err = a.client.ApiClient().Do(ctx, http.MethodPost, endpointUrl, opts...)
	return &queryEndpointResponse, err
}
