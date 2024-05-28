// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Apps API methods
type appsImpl struct {
	client *client.DatabricksClient
}

func (a *appsImpl) Create(ctx context.Context, request CreateAppRequest) (*App, error) {
	var app App
	path := "/api/2.0/preview/apps"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &app)
	return &app, err
}

func (a *appsImpl) CreateDeployment(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0/preview/apps/%v/deployments", request.AppName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &appDeployment)
	return &appDeployment, err
}

func (a *appsImpl) Delete(ctx context.Context, request DeleteAppRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/apps/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *appsImpl) Get(ctx context.Context, request GetAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/preview/apps/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &app)
	return &app, err
}

func (a *appsImpl) GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0/preview/apps/%v/deployments/%v", request.AppName, request.DeploymentId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &appDeployment)
	return &appDeployment, err
}

func (a *appsImpl) GetEnvironment(ctx context.Context, request GetAppEnvironmentRequest) (*AppEnvironment, error) {
	var appEnvironment AppEnvironment
	path := fmt.Sprintf("/api/2.0/preview/apps/%v/environment", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &appEnvironment)
	return &appEnvironment, err
}

func (a *appsImpl) List(ctx context.Context, request ListAppsRequest) (*ListAppsResponse, error) {
	var listAppsResponse ListAppsResponse
	path := "/api/2.0/preview/apps"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAppsResponse)
	return &listAppsResponse, err
}

func (a *appsImpl) ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {
	var listAppDeploymentsResponse ListAppDeploymentsResponse
	path := fmt.Sprintf("/api/2.0/preview/apps/%v/deployments", request.AppName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAppDeploymentsResponse)
	return &listAppDeploymentsResponse, err
}

func (a *appsImpl) Stop(ctx context.Context, request StopAppRequest) error {
	var stopAppResponse StopAppResponse
	path := fmt.Sprintf("/api/2.0/preview/apps/%v/stop", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &stopAppResponse)
	return err
}

func (a *appsImpl) Update(ctx context.Context, request UpdateAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/preview/apps/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &app)
	return &app, err
}

// unexported type that holds implementations of just ServingEndpointDataPlane API methods
type servingEndpointDataPlaneImpl struct {
	client       *client.DatabricksClient
	controlPlane *ServingEndpointsAPI
}

func (a *servingEndpointDataPlaneImpl) Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error) {
	var queryEndpointResponse QueryEndpointResponse
	path := fmt.Sprintf("/serving-endpoints/%v/invocations", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &queryEndpointResponse)
	return &queryEndpointResponse, err
}

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
