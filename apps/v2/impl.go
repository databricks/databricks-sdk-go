// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just apps API methods
type appsImpl struct {
	client *client.DatabricksClient
}

func (a *appsImpl) Create(ctx context.Context, request CreateAppRequest) (*App, error) {
	var app App
	path := "/api/2.0/apps"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request.App, &app)
	return &app, err
}

func (a *appsImpl) Delete(ctx context.Context, request DeleteAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &app)
	return &app, err
}

func (a *appsImpl) Deploy(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", request.AppName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request.AppDeployment, &appDeployment)
	return &appDeployment, err
}

func (a *appsImpl) Get(ctx context.Context, request GetAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &app)
	return &app, err
}

func (a *appsImpl) GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments/%v", request.AppName, request.DeploymentId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &appDeployment)
	return &appDeployment, err
}

func (a *appsImpl) GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error) {
	var getAppPermissionLevelsResponse GetAppPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v/permissionLevels", request.AppName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getAppPermissionLevelsResponse)
	return &getAppPermissionLevelsResponse, err
}

func (a *appsImpl) GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &appPermissions)
	return &appPermissions, err
}

func (a *appsImpl) List(ctx context.Context, request ListAppsRequest) (*ListAppsResponse, error) {
	var listAppsResponse ListAppsResponse
	path := "/api/2.0/apps"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAppsResponse)
	return &listAppsResponse, err
}

func (a *appsImpl) ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {
	var listAppDeploymentsResponse ListAppDeploymentsResponse
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", request.AppName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAppDeploymentsResponse)
	return &listAppDeploymentsResponse, err
}

func (a *appsImpl) SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &appPermissions)
	return &appPermissions, err
}

func (a *appsImpl) Start(ctx context.Context, request StartAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v/start", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &app)
	return &app, err
}

func (a *appsImpl) Stop(ctx context.Context, request StopAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v/stop", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &app)
	return &app, err
}

func (a *appsImpl) Update(ctx context.Context, request UpdateAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request.App, &app)
	return &app, err
}

func (a *appsImpl) UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &appPermissions)
	return &appPermissions, err
}
