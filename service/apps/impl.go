// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Apps API methods
type appsImpl struct {
	client *client.DatabricksClient
}

func (a *appsImpl) Create(ctx context.Context, request CreateAppRequest) (*App, error) {
	var app App
	path := "/api/2.0/apps"
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

func (a *appsImpl) CreateUpdate(ctx context.Context, request AsyncUpdateAppRequest) (*AppUpdate, error) {
	var appUpdate AppUpdate
	path := fmt.Sprintf("/api/2.0/apps/%v/update", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &appUpdate)
	return &appUpdate, err
}

func (a *appsImpl) Delete(ctx context.Context, request DeleteAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsImpl) Deploy(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.AppDeployment, &appDeployment)
	return &appDeployment, err
}

func (a *appsImpl) Get(ctx context.Context, request GetAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsImpl) GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error) {
	var appDeployment AppDeployment
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments/%v", request.AppName, request.DeploymentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appDeployment)
	return &appDeployment, err
}

func (a *appsImpl) GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error) {
	var getAppPermissionLevelsResponse GetAppPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v/permissionLevels", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAppPermissionLevelsResponse)
	return &getAppPermissionLevelsResponse, err
}

func (a *appsImpl) GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}

func (a *appsImpl) GetUpdate(ctx context.Context, request GetAppUpdateRequest) (*AppUpdate, error) {
	var appUpdate AppUpdate
	path := fmt.Sprintf("/api/2.0/apps/%v/update", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &appUpdate)
	return &appUpdate, err
}

// Lists all apps in the workspace.
func (a *appsImpl) List(ctx context.Context, request ListAppsRequest) listing.Iterator[App] {

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

// Lists all apps in the workspace.
func (a *appsImpl) ListAll(ctx context.Context, request ListAppsRequest) ([]App, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[App](ctx, iterator)
}

func (a *appsImpl) internalList(ctx context.Context, request ListAppsRequest) (*ListAppsResponse, error) {
	var listAppsResponse ListAppsResponse
	path := "/api/2.0/apps"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAppsResponse)
	return &listAppsResponse, err
}

// Lists all app deployments for the app with the supplied name.
func (a *appsImpl) ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) listing.Iterator[AppDeployment] {

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

// Lists all app deployments for the app with the supplied name.
func (a *appsImpl) ListDeploymentsAll(ctx context.Context, request ListAppDeploymentsRequest) ([]AppDeployment, error) {
	iterator := a.ListDeployments(ctx, request)
	return listing.ToSlice[AppDeployment](ctx, iterator)
}

func (a *appsImpl) internalListDeployments(ctx context.Context, request ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {
	var listAppDeploymentsResponse ListAppDeploymentsResponse
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAppDeploymentsResponse)
	return &listAppDeploymentsResponse, err
}

func (a *appsImpl) SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}

func (a *appsImpl) Start(ctx context.Context, request StartAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v/start", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsImpl) Stop(ctx context.Context, request StopAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v/stop", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &app)
	return &app, err
}

func (a *appsImpl) Update(ctx context.Context, request UpdateAppRequest) (*App, error) {
	var app App
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.App, &app)
	return &app, err
}

func (a *appsImpl) UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	var appPermissions AppPermissions
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &appPermissions)
	return &appPermissions, err
}

// unexported type that holds implementations of just AppsSettings API methods
type appsSettingsImpl struct {
	client *client.DatabricksClient
}

func (a *appsSettingsImpl) CreateCustomTemplate(ctx context.Context, request CreateCustomTemplateRequest) (*CustomTemplate, error) {
	var customTemplate CustomTemplate
	path := "/api/2.0/apps-settings/templates"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Template, &customTemplate)
	return &customTemplate, err
}

func (a *appsSettingsImpl) DeleteCustomTemplate(ctx context.Context, request DeleteCustomTemplateRequest) (*CustomTemplate, error) {
	var customTemplate CustomTemplate
	path := fmt.Sprintf("/api/2.0/apps-settings/templates/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &customTemplate)
	return &customTemplate, err
}

func (a *appsSettingsImpl) GetCustomTemplate(ctx context.Context, request GetCustomTemplateRequest) (*CustomTemplate, error) {
	var customTemplate CustomTemplate
	path := fmt.Sprintf("/api/2.0/apps-settings/templates/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &customTemplate)
	return &customTemplate, err
}

// Lists all custom templates in the workspace.
func (a *appsSettingsImpl) ListCustomTemplates(ctx context.Context, request ListCustomTemplatesRequest) listing.Iterator[CustomTemplate] {

	getNextPage := func(ctx context.Context, req ListCustomTemplatesRequest) (*ListCustomTemplatesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCustomTemplates(ctx, req)
	}
	getItems := func(resp *ListCustomTemplatesResponse) []CustomTemplate {
		return resp.Templates
	}
	getNextReq := func(resp *ListCustomTemplatesResponse) *ListCustomTemplatesRequest {
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

// Lists all custom templates in the workspace.
func (a *appsSettingsImpl) ListCustomTemplatesAll(ctx context.Context, request ListCustomTemplatesRequest) ([]CustomTemplate, error) {
	iterator := a.ListCustomTemplates(ctx, request)
	return listing.ToSlice[CustomTemplate](ctx, iterator)
}

func (a *appsSettingsImpl) internalListCustomTemplates(ctx context.Context, request ListCustomTemplatesRequest) (*ListCustomTemplatesResponse, error) {
	var listCustomTemplatesResponse ListCustomTemplatesResponse
	path := "/api/2.0/apps-settings/templates"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCustomTemplatesResponse)
	return &listCustomTemplatesResponse, err
}

func (a *appsSettingsImpl) UpdateCustomTemplate(ctx context.Context, request UpdateCustomTemplateRequest) (*CustomTemplate, error) {
	var customTemplate CustomTemplate
	path := fmt.Sprintf("/api/2.0/apps-settings/templates/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request.Template, &customTemplate)
	return &customTemplate, err
}
