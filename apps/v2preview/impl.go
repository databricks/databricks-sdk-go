// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package appspreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just AppsPreview API methods
type appsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *appsPreviewImpl) Create(ctx context.Context, request CreateAppRequest) (*App, error) {
	var app App
	path := "/api/2.0preview/apps"
	queryParams := make(map[string]any)
	if request.NoCompute != false || slices.Contains(request.ForceSendFields, "NoCompute") {
		queryParams["no_compute"] = request.NoCompute
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.App, &app)
	return &app, err
}

func (a *appsPreviewImpl) Delete(ctx context.Context, request DeleteAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0preview/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsPreviewImpl) Deploy(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0preview/apps/%v/deployments", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.AppDeployment, &appDeployment)
	return &appDeployment, err
}

func (a *appsPreviewImpl) Get(ctx context.Context, request GetAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0preview/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsPreviewImpl) GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0preview/apps/%v/deployments/%v", request.AppName, request.DeploymentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appDeployment)
	return &appDeployment, err
}

func (a *appsPreviewImpl) GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error) {
	var getAppPermissionLevelsResponse GetAppPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0preview/permissions/apps/%v/permissionLevels", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAppPermissionLevelsResponse)
	return &getAppPermissionLevelsResponse, err
}

func (a *appsPreviewImpl) GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}

// List apps.
//
// Lists all apps in the workspace.
func (a *appsPreviewImpl) List(ctx context.Context, request ListAppsRequest) listing.Iterator[App] {

	getNextPage := func(ctx context.Context, req ListAppsRequest) (*ListAppsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAppsResponse) []App {
		return resp.Apps
	}
	getNextReq := func(resp *ListAppsResponse) *ListAppsRequest {
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

// List apps.
//
// Lists all apps in the workspace.
func (a *appsPreviewImpl) ListAll(ctx context.Context, request ListAppsRequest) ([]App, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[App](ctx, iterator)
}
func (a *appsPreviewImpl) internalList(ctx context.Context, request ListAppsRequest) (*ListAppsResponse, error) {
	var listAppsResponse ListAppsResponse
	path := "/api/2.0preview/apps"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAppsResponse)
	return &listAppsResponse, err
}

// List app deployments.
//
// Lists all app deployments for the app with the supplied name.
func (a *appsPreviewImpl) ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) listing.Iterator[AppDeployment] {

	getNextPage := func(ctx context.Context, req ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDeployments(ctx, req)
	}
	getItems := func(resp *ListAppDeploymentsResponse) []AppDeployment {
		return resp.AppDeployments
	}
	getNextReq := func(resp *ListAppDeploymentsResponse) *ListAppDeploymentsRequest {
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

// List app deployments.
//
// Lists all app deployments for the app with the supplied name.
func (a *appsPreviewImpl) ListDeploymentsAll(ctx context.Context, request ListAppDeploymentsRequest) ([]AppDeployment, error) {
	iterator := a.ListDeployments(ctx, request)
	return listing.ToSlice[AppDeployment](ctx, iterator)
}
func (a *appsPreviewImpl) internalListDeployments(ctx context.Context, request ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {
	var listAppDeploymentsResponse ListAppDeploymentsResponse
	path := fmt.Sprintf("/api/2.0preview/apps/%v/deployments", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAppDeploymentsResponse)
	return &listAppDeploymentsResponse, err
}

func (a *appsPreviewImpl) SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}

func (a *appsPreviewImpl) Start(ctx context.Context, request StartAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0preview/apps/%v/start", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsPreviewImpl) Stop(ctx context.Context, request StopAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0preview/apps/%v/stop", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsPreviewImpl) Update(ctx context.Context, request UpdateAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0preview/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.App, &app)
	return &app, err
}

func (a *appsPreviewImpl) UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}
