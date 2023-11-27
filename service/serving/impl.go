// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Apps API methods
type appsImpl struct {
	client *client.DatabricksClient
}

func (a *appsImpl) Create(ctx context.Context, request DeployAppRequest) (*DeploymentStatus, error) {
	var deploymentStatus DeploymentStatus
	path := "/api/2.0/preview/apps/deployments"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deploymentStatus)
	return &deploymentStatus, err
}

func (a *appsImpl) DeleteApp(ctx context.Context, request DeleteAppRequest) (*DeleteAppResponse, error) {
	var deleteAppResponse DeleteAppResponse
	path := fmt.Sprintf("/api/2.0/preview/apps/instances/%v", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteAppResponse)
	return &deleteAppResponse, err
}

func (a *appsImpl) GetApp(ctx context.Context, request GetAppRequest) (*GetAppResponse, error) {
	var getAppResponse GetAppResponse
	path := fmt.Sprintf("/api/2.0/preview/apps/instances/%v", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getAppResponse)
	return &getAppResponse, err
}

func (a *appsImpl) GetAppDeploymentStatus(ctx context.Context, request GetAppDeploymentStatusRequest) (*DeploymentStatus, error) {
	var deploymentStatus DeploymentStatus
	path := fmt.Sprintf("/api/2.0/preview/apps/deployments/%v", strings.TrimPrefix(fmt.Sprint(request.DeploymentId), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &deploymentStatus)
	return &deploymentStatus, err
}

func (a *appsImpl) GetApps(ctx context.Context) (*ListAppsResponse, error) {
	var listAppsResponse ListAppsResponse
	path := "/api/2.0/preview/apps/instances"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listAppsResponse)
	return &listAppsResponse, err
}

func (a *appsImpl) GetEvents(ctx context.Context, request GetEventsRequest) (*ListAppEventsResponse, error) {
	var listAppEventsResponse ListAppEventsResponse
	path := fmt.Sprintf("/api/2.0/preview/apps/%v/events", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAppEventsResponse)
	return &listAppEventsResponse, err
}

// unexported type that holds implementations of just ServingEndpoints API methods
type servingEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *servingEndpointsImpl) BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error) {
	var buildLogsResponse BuildLogsResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/build-logs", strings.TrimPrefix(fmt.Sprint(request.Name), "/"), strings.TrimPrefix(fmt.Sprint(request.ServedModelName), "/"))
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
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *servingEndpointsImpl) ExportMetrics(ctx context.Context, request ExportMetricsRequest) error {
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/metrics", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, nil)
	return err
}

func (a *servingEndpointsImpl) Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) GetPermissionLevels(ctx context.Context, request GetServingEndpointPermissionLevelsRequest) (*GetServingEndpointPermissionLevelsResponse, error) {
	var getServingEndpointPermissionLevelsResponse GetServingEndpointPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v/permissionLevels", strings.TrimPrefix(fmt.Sprint(request.ServingEndpointId), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getServingEndpointPermissionLevelsResponse)
	return &getServingEndpointPermissionLevelsResponse, err
}

func (a *servingEndpointsImpl) GetPermissions(ctx context.Context, request GetServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", strings.TrimPrefix(fmt.Sprint(request.ServingEndpointId), "/"))
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
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/logs", strings.TrimPrefix(fmt.Sprint(request.Name), "/"), strings.TrimPrefix(fmt.Sprint(request.ServedModelName), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &serverLogsResponse)
	return &serverLogsResponse, err
}

func (a *servingEndpointsImpl) Patch(ctx context.Context, request PatchServingEndpointTags) ([]EndpointTag, error) {
	var endpointTagList []EndpointTag
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/tags", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &endpointTagList)
	return endpointTagList, err
}

func (a *servingEndpointsImpl) Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error) {
	var queryEndpointResponse QueryEndpointResponse
	path := fmt.Sprintf("/serving-endpoints/%v/invocations", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &queryEndpointResponse)
	return &queryEndpointResponse, err
}

func (a *servingEndpointsImpl) SetPermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", strings.TrimPrefix(fmt.Sprint(request.ServingEndpointId), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &servingEndpointPermissions)
	return &servingEndpointPermissions, err
}

func (a *servingEndpointsImpl) UpdateConfig(ctx context.Context, request EndpointCoreConfigInput) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/config", strings.TrimPrefix(fmt.Sprint(request.Name), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) UpdatePermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", strings.TrimPrefix(fmt.Sprint(request.ServingEndpointId), "/"))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &servingEndpointPermissions)
	return &servingEndpointPermissions, err
}
