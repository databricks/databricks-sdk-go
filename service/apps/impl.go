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

	requestPb, pbErr := createAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appPb
	path := "/api/2.0/apps"
	queryParams := make(map[string]any)
	if requestPb.NoCompute != false || slices.Contains(requestPb.ForceSendFields, "NoCompute") {
		queryParams["no_compute"] = requestPb.NoCompute
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).App,
		&appPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Delete(ctx context.Context, request DeleteAppRequest) (*App, error) {

	requestPb, pbErr := deleteAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appPb
	path := fmt.Sprintf("/api/2.0/apps/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&appPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Deploy(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error) {

	requestPb, pbErr := createAppDeploymentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appDeploymentPb appDeploymentPb
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", requestPb.AppName)
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
		(*requestPb).AppDeployment,
		&appDeploymentPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appDeploymentFromPb(&appDeploymentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Get(ctx context.Context, request GetAppRequest) (*App, error) {

	requestPb, pbErr := getAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appPb
	path := fmt.Sprintf("/api/2.0/apps/%v", requestPb.Name)
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
		&appPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error) {

	requestPb, pbErr := getAppDeploymentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appDeploymentPb appDeploymentPb
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments/%v", requestPb.AppName, requestPb.DeploymentId)
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
		&appDeploymentPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appDeploymentFromPb(&appDeploymentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error) {

	requestPb, pbErr := getAppPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getAppPermissionLevelsResponsePb getAppPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v/permissionLevels", requestPb.AppName)
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
		&getAppPermissionLevelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getAppPermissionLevelsResponseFromPb(&getAppPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error) {

	requestPb, pbErr := getAppPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPermissionsPb appPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", requestPb.AppName)
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
		&appPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appPermissionsFromPb(&appPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List apps.
//
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

// List apps.
//
// Lists all apps in the workspace.
func (a *appsImpl) ListAll(ctx context.Context, request ListAppsRequest) ([]App, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[App](ctx, iterator)
}

func (a *appsImpl) internalList(ctx context.Context, request ListAppsRequest) (*ListAppsResponse, error) {

	requestPb, pbErr := listAppsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAppsResponsePb listAppsResponsePb
	path := "/api/2.0/apps"
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
		&listAppsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listAppsResponseFromPb(&listAppsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List app deployments.
//
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

// List app deployments.
//
// Lists all app deployments for the app with the supplied name.
func (a *appsImpl) ListDeploymentsAll(ctx context.Context, request ListAppDeploymentsRequest) ([]AppDeployment, error) {
	iterator := a.ListDeployments(ctx, request)
	return listing.ToSlice[AppDeployment](ctx, iterator)
}

func (a *appsImpl) internalListDeployments(ctx context.Context, request ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {

	requestPb, pbErr := listAppDeploymentsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAppDeploymentsResponsePb listAppDeploymentsResponsePb
	path := fmt.Sprintf("/api/2.0/apps/%v/deployments", requestPb.AppName)
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
		&listAppDeploymentsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listAppDeploymentsResponseFromPb(&listAppDeploymentsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {

	requestPb, pbErr := appPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPermissionsPb appPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", requestPb.AppName)
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
		&appPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appPermissionsFromPb(&appPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Start(ctx context.Context, request StartAppRequest) (*App, error) {

	requestPb, pbErr := startAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appPb
	path := fmt.Sprintf("/api/2.0/apps/%v/start", requestPb.Name)
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
		&appPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Stop(ctx context.Context, request StopAppRequest) (*App, error) {

	requestPb, pbErr := stopAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appPb
	path := fmt.Sprintf("/api/2.0/apps/%v/stop", requestPb.Name)
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
		&appPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) Update(ctx context.Context, request UpdateAppRequest) (*App, error) {

	requestPb, pbErr := updateAppRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPb appPb
	path := fmt.Sprintf("/api/2.0/apps/%v", requestPb.Name)
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
		(*requestPb).App,
		&appPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appFromPb(&appPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *appsImpl) UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error) {

	requestPb, pbErr := appPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var appPermissionsPb appPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/apps/%v", requestPb.AppName)
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
		&appPermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := appPermissionsFromPb(&appPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
