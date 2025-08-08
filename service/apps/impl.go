// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/apps/appspb"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Apps API methods
type appsImpl struct {
	client *client.DatabricksClient
}

func (a *appsImpl) Create(ctx context.Context, request CreateAppRequest) (*App, error) {
	requestPb, pbErr := CreateAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appspb.AppPb
	path := "/api/2.0/apps"
	queryParams := make(map[string]any)
	if requestPb.NoCompute != false || slices.Contains(requestPb.ForceSendFields, "NoCompute") {
		queryParams["no_compute"] = requestPb.NoCompute
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.App,
		&appPb,
	)
	resp, err := AppFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Delete(ctx context.Context, request DeleteAppRequest) (*App, error) {
	requestPb, pbErr := DeleteAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appspb.AppPb
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&appPb,
	)
	resp, err := AppFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Deploy(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error) {
	requestPb, pbErr := CreateAppDeploymentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appDeploymentPb appspb.AppDeploymentPb
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.AppDeployment,
		&appDeploymentPb,
	)
	resp, err := AppDeploymentFromPb(&appDeploymentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Get(ctx context.Context, request GetAppRequest) (*App, error) {
	requestPb, pbErr := GetAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appspb.AppPb
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&appPb,
	)
	resp, err := AppFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error) {
	requestPb, pbErr := GetAppDeploymentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appDeploymentPb appspb.AppDeploymentPb
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments/%v", request.AppName, request.DeploymentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&appDeploymentPb,
	)
	resp, err := AppDeploymentFromPb(&appDeploymentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error) {
	requestPb, pbErr := GetAppPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getAppPermissionLevelsResponsePb appspb.GetAppPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v/permissionLevels", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getAppPermissionLevelsResponsePb,
	)
	resp, err := GetAppPermissionLevelsResponseFromPb(&getAppPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error) {
	requestPb, pbErr := GetAppPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPermissionsPb appspb.AppPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&appPermissionsPb,
	)
	resp, err := AppPermissionsFromPb(&appPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
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
	requestPb, pbErr := ListAppsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAppsResponsePb appspb.ListAppsResponsePb
	path := "/api/2.0/apps"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listAppsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListAppsResponseFromPb(&listAppsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
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
	requestPb, pbErr := ListAppDeploymentsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAppDeploymentsResponsePb appspb.ListAppDeploymentsResponsePb
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listAppDeploymentsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListAppDeploymentsResponseFromPb(&listAppDeploymentsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	requestPb, pbErr := AppPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPermissionsPb appspb.AppPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&appPermissionsPb,
	)
	resp, err := AppPermissionsFromPb(&appPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Start(ctx context.Context, request StartAppRequest) (*App, error) {
	requestPb, pbErr := StartAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appspb.AppPb
	path := fmt.Sprintf("/api/2.0/apps/%v/start", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&appPb,
	)
	resp, err := AppFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Stop(ctx context.Context, request StopAppRequest) (*App, error) {
	requestPb, pbErr := StopAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appspb.AppPb
	path := fmt.Sprintf("/api/2.0/apps/%v/stop", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&appPb,
	)
	resp, err := AppFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Update(ctx context.Context, request UpdateAppRequest) (*App, error) {
	requestPb, pbErr := UpdateAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appspb.AppPb
	path := fmt.Sprintf("/api/2.0/apps/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.App,
		&appPb,
	)
	resp, err := AppFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {
	requestPb, pbErr := AppPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPermissionsPb appspb.AppPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", request.AppName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&appPermissionsPb,
	)
	resp, err := AppPermissionsFromPb(&appPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
