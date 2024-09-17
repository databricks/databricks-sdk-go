// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Apps run directly on a customer’s Databricks instance, integrate with their data, use and extend Databricks services, and enable users to interact through single sign-on.
package apps

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type AppsInterface interface {

	// WaitGetAppActive repeatedly calls [AppsAPI.Get] and waits to reach ACTIVE state
	WaitGetAppActive(ctx context.Context, name string,
		timeout time.Duration, callback func(*App)) (*App, error)

	// WaitGetAppStopped repeatedly calls [AppsAPI.Get] and waits to reach STOPPED state
	WaitGetAppStopped(ctx context.Context, name string,
		timeout time.Duration, callback func(*App)) (*App, error)

	// WaitGetDeploymentAppSucceeded repeatedly calls [AppsAPI.GetDeployment] and waits to reach SUCCEEDED state
	WaitGetDeploymentAppSucceeded(ctx context.Context, appName string, deploymentId string,
		timeout time.Duration, callback func(*AppDeployment)) (*AppDeployment, error)

	// Create an app.
	//
	// Creates a new app.
	Create(ctx context.Context, createAppRequest CreateAppRequest) (*WaitGetAppActive[App], error)

	// Calls [AppsAPIInterface.Create] and waits to reach ACTIVE state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[App](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.Create].Get() or [AppsAPIInterface.WaitGetAppActive]
	CreateAndWait(ctx context.Context, createAppRequest CreateAppRequest, options ...retries.Option[App]) (*App, error)

	// Delete an app.
	//
	// Deletes an app.
	Delete(ctx context.Context, request DeleteAppRequest) (*App, error)

	// Delete an app.
	//
	// Deletes an app.
	DeleteByName(ctx context.Context, name string) (*App, error)

	// Create an app deployment.
	//
	// Creates an app deployment for the app with the supplied name.
	Deploy(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest) (*WaitGetDeploymentAppSucceeded[AppDeployment], error)

	// Calls [AppsAPIInterface.Deploy] and waits to reach SUCCEEDED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[AppDeployment](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.Deploy].Get() or [AppsAPIInterface.WaitGetDeploymentAppSucceeded]
	DeployAndWait(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest, options ...retries.Option[AppDeployment]) (*AppDeployment, error)

	// Get an app.
	//
	// Retrieves information for the app with the supplied name.
	Get(ctx context.Context, request GetAppRequest) (*App, error)

	// Get an app.
	//
	// Retrieves information for the app with the supplied name.
	GetByName(ctx context.Context, name string) (*App, error)

	// Get an app deployment.
	//
	// Retrieves information for the app deployment with the supplied name and
	// deployment id.
	GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error)

	// Get an app deployment.
	//
	// Retrieves information for the app deployment with the supplied name and
	// deployment id.
	GetDeploymentByAppNameAndDeploymentId(ctx context.Context, appName string, deploymentId string) (*AppDeployment, error)

	// Get app permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error)

	// Get app permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByAppName(ctx context.Context, appName string) (*GetAppPermissionLevelsResponse, error)

	// Get app permissions.
	//
	// Gets the permissions of an app. Apps can inherit permissions from their root
	// object.
	GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error)

	// Get app permissions.
	//
	// Gets the permissions of an app. Apps can inherit permissions from their root
	// object.
	GetPermissionsByAppName(ctx context.Context, appName string) (*AppPermissions, error)

	// List apps.
	//
	// Lists all apps in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListAppsRequest) listing.Iterator[App]

	// List apps.
	//
	// Lists all apps in the workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListAppsRequest) ([]App, error)

	// List app deployments.
	//
	// Lists all app deployments for the app with the supplied name.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) listing.Iterator[AppDeployment]

	// List app deployments.
	//
	// Lists all app deployments for the app with the supplied name.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDeploymentsAll(ctx context.Context, request ListAppDeploymentsRequest) ([]AppDeployment, error)

	// List app deployments.
	//
	// Lists all app deployments for the app with the supplied name.
	ListDeploymentsByAppName(ctx context.Context, appName string) (*ListAppDeploymentsResponse, error)

	// Set app permissions.
	//
	// Sets permissions on an app. Apps can inherit permissions from their root
	// object.
	SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error)

	// Start an app.
	//
	// Start the last active deployment of the app in the workspace.
	Start(ctx context.Context, startAppRequest StartAppRequest) (*WaitGetAppActive[App], error)

	// Calls [AppsAPIInterface.Start] and waits to reach ACTIVE state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[App](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.Start].Get() or [AppsAPIInterface.WaitGetAppActive]
	StartAndWait(ctx context.Context, startAppRequest StartAppRequest, options ...retries.Option[App]) (*App, error)

	// Stop an app.
	//
	// Stops the active deployment of the app in the workspace.
	Stop(ctx context.Context, stopAppRequest StopAppRequest) (*WaitGetAppStopped[App], error)

	// Calls [AppsAPIInterface.Stop] and waits to reach STOPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[App](60*time.Minute) functional option.
	//
	// Deprecated: use [AppsAPIInterface.Stop].Get() or [AppsAPIInterface.WaitGetAppStopped]
	StopAndWait(ctx context.Context, stopAppRequest StopAppRequest, options ...retries.Option[App]) (*App, error)

	// Update an app.
	//
	// Updates the app with the supplied name.
	Update(ctx context.Context, request UpdateAppRequest) (*App, error)

	// Update app permissions.
	//
	// Updates the permissions on an app. Apps can inherit permissions from their
	// root object.
	UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error)
}

func NewApps(client *client.DatabricksClient) *AppsAPI {
	return &AppsAPI{
		appsImpl: appsImpl{
			client: client,
		},
	}
}

// Apps run directly on a customer’s Databricks instance, integrate with their
// data, use and extend Databricks services, and enable users to interact
// through single sign-on.
type AppsAPI struct {
	appsImpl
}

// WaitGetAppActive repeatedly calls [AppsAPI.Get] and waits to reach ACTIVE state
func (a *AppsAPI) WaitGetAppActive(ctx context.Context, name string,
	timeout time.Duration, callback func(*App)) (*App, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[App](ctx, timeout, func() (*App, *retries.Err) {
		app, err := a.Get(ctx, GetAppRequest{
			Name: name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(app)
		}
		status := app.ComputeStatus.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if app.ComputeStatus != nil {
			statusMessage = app.ComputeStatus.Message
		}
		switch status {
		case ComputeStateActive: // target state
			return app, nil
		case ComputeStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ComputeStateActive, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetAppActive is a wrapper that calls [AppsAPI.WaitGetAppActive] and waits to reach ACTIVE state.
type WaitGetAppActive[R any] struct {
	Response *R
	Name     string `json:"name"`
	Poll     func(time.Duration, func(*App)) (*App, error)
	callback func(*App)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetAppActive[R]) OnProgress(callback func(*App)) *WaitGetAppActive[R] {
	w.callback = callback
	return w
}

// Get the App with the default timeout of 20 minutes.
func (w *WaitGetAppActive[R]) Get() (*App, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the App with custom timeout.
func (w *WaitGetAppActive[R]) GetWithTimeout(timeout time.Duration) (*App, error) {
	return w.Poll(timeout, w.callback)
}

// WaitGetAppStopped repeatedly calls [AppsAPI.Get] and waits to reach STOPPED state
func (a *AppsAPI) WaitGetAppStopped(ctx context.Context, name string,
	timeout time.Duration, callback func(*App)) (*App, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[App](ctx, timeout, func() (*App, *retries.Err) {
		app, err := a.Get(ctx, GetAppRequest{
			Name: name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(app)
		}
		status := app.ComputeStatus.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if app.ComputeStatus != nil {
			statusMessage = app.ComputeStatus.Message
		}
		switch status {
		case ComputeStateStopped: // target state
			return app, nil
		case ComputeStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ComputeStateStopped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetAppStopped is a wrapper that calls [AppsAPI.WaitGetAppStopped] and waits to reach STOPPED state.
type WaitGetAppStopped[R any] struct {
	Response *R
	Name     string `json:"name"`
	Poll     func(time.Duration, func(*App)) (*App, error)
	callback func(*App)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetAppStopped[R]) OnProgress(callback func(*App)) *WaitGetAppStopped[R] {
	w.callback = callback
	return w
}

// Get the App with the default timeout of 20 minutes.
func (w *WaitGetAppStopped[R]) Get() (*App, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the App with custom timeout.
func (w *WaitGetAppStopped[R]) GetWithTimeout(timeout time.Duration) (*App, error) {
	return w.Poll(timeout, w.callback)
}

// WaitGetDeploymentAppSucceeded repeatedly calls [AppsAPI.GetDeployment] and waits to reach SUCCEEDED state
func (a *AppsAPI) WaitGetDeploymentAppSucceeded(ctx context.Context, appName string, deploymentId string,
	timeout time.Duration, callback func(*AppDeployment)) (*AppDeployment, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[AppDeployment](ctx, timeout, func() (*AppDeployment, *retries.Err) {
		appDeployment, err := a.GetDeployment(ctx, GetAppDeploymentRequest{
			AppName:      appName,
			DeploymentId: deploymentId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(appDeployment)
		}
		status := appDeployment.Status.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if appDeployment.Status != nil {
			statusMessage = appDeployment.Status.Message
		}
		switch status {
		case AppDeploymentStateSucceeded: // target state
			return appDeployment, nil
		case AppDeploymentStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				AppDeploymentStateSucceeded, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetDeploymentAppSucceeded is a wrapper that calls [AppsAPI.WaitGetDeploymentAppSucceeded] and waits to reach SUCCEEDED state.
type WaitGetDeploymentAppSucceeded[R any] struct {
	Response     *R
	AppName      string `json:"app_name"`
	DeploymentId string `json:"deployment_id"`
	Poll         func(time.Duration, func(*AppDeployment)) (*AppDeployment, error)
	callback     func(*AppDeployment)
	timeout      time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetDeploymentAppSucceeded[R]) OnProgress(callback func(*AppDeployment)) *WaitGetDeploymentAppSucceeded[R] {
	w.callback = callback
	return w
}

// Get the AppDeployment with the default timeout of 20 minutes.
func (w *WaitGetDeploymentAppSucceeded[R]) Get() (*AppDeployment, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the AppDeployment with custom timeout.
func (w *WaitGetDeploymentAppSucceeded[R]) GetWithTimeout(timeout time.Duration) (*AppDeployment, error) {
	return w.Poll(timeout, w.callback)
}

// Create an app.
//
// Creates a new app.
func (a *AppsAPI) Create(ctx context.Context, createAppRequest CreateAppRequest) (*WaitGetAppActive[App], error) {
	app, err := a.appsImpl.Create(ctx, createAppRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetAppActive[App]{
		Response: app,
		Name:     app.Name,
		Poll: func(timeout time.Duration, callback func(*App)) (*App, error) {
			return a.WaitGetAppActive(ctx, app.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.Create] and waits to reach ACTIVE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[App](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.Create].Get() or [AppsAPI.WaitGetAppActive]
func (a *AppsAPI) CreateAndWait(ctx context.Context, createAppRequest CreateAppRequest, options ...retries.Option[App]) (*App, error) {
	wait, err := a.Create(ctx, createAppRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[App]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *App) {
		for _, o := range options {
			o(&retries.Info[App]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete an app.
//
// Deletes an app.
func (a *AppsAPI) DeleteByName(ctx context.Context, name string) (*App, error) {
	return a.appsImpl.Delete(ctx, DeleteAppRequest{
		Name: name,
	})
}

// Create an app deployment.
//
// Creates an app deployment for the app with the supplied name.
func (a *AppsAPI) Deploy(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest) (*WaitGetDeploymentAppSucceeded[AppDeployment], error) {
	appDeployment, err := a.appsImpl.Deploy(ctx, createAppDeploymentRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetDeploymentAppSucceeded[AppDeployment]{
		Response:     appDeployment,
		AppName:      createAppDeploymentRequest.AppName,
		DeploymentId: appDeployment.DeploymentId,
		Poll: func(timeout time.Duration, callback func(*AppDeployment)) (*AppDeployment, error) {
			return a.WaitGetDeploymentAppSucceeded(ctx, createAppDeploymentRequest.AppName, appDeployment.DeploymentId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.Deploy] and waits to reach SUCCEEDED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[AppDeployment](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.Deploy].Get() or [AppsAPI.WaitGetDeploymentAppSucceeded]
func (a *AppsAPI) DeployAndWait(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest, options ...retries.Option[AppDeployment]) (*AppDeployment, error) {
	wait, err := a.Deploy(ctx, createAppDeploymentRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[AppDeployment]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *AppDeployment) {
		for _, o := range options {
			o(&retries.Info[AppDeployment]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Get an app.
//
// Retrieves information for the app with the supplied name.
func (a *AppsAPI) GetByName(ctx context.Context, name string) (*App, error) {
	return a.appsImpl.Get(ctx, GetAppRequest{
		Name: name,
	})
}

// Get an app deployment.
//
// Retrieves information for the app deployment with the supplied name and
// deployment id.
func (a *AppsAPI) GetDeploymentByAppNameAndDeploymentId(ctx context.Context, appName string, deploymentId string) (*AppDeployment, error) {
	return a.appsImpl.GetDeployment(ctx, GetAppDeploymentRequest{
		AppName:      appName,
		DeploymentId: deploymentId,
	})
}

// Get app permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *AppsAPI) GetPermissionLevelsByAppName(ctx context.Context, appName string) (*GetAppPermissionLevelsResponse, error) {
	return a.appsImpl.GetPermissionLevels(ctx, GetAppPermissionLevelsRequest{
		AppName: appName,
	})
}

// Get app permissions.
//
// Gets the permissions of an app. Apps can inherit permissions from their root
// object.
func (a *AppsAPI) GetPermissionsByAppName(ctx context.Context, appName string) (*AppPermissions, error) {
	return a.appsImpl.GetPermissions(ctx, GetAppPermissionsRequest{
		AppName: appName,
	})
}

// List apps.
//
// Lists all apps in the workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AppsAPI) List(ctx context.Context, request ListAppsRequest) listing.Iterator[App] {

	getNextPage := func(ctx context.Context, req ListAppsRequest) (*ListAppsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.appsImpl.List(ctx, req)
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
//
// This method is generated by Databricks SDK Code Generator.
func (a *AppsAPI) ListAll(ctx context.Context, request ListAppsRequest) ([]App, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[App](ctx, iterator)
}

// List app deployments.
//
// Lists all app deployments for the app with the supplied name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AppsAPI) ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) listing.Iterator[AppDeployment] {

	getNextPage := func(ctx context.Context, req ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.appsImpl.ListDeployments(ctx, req)
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
//
// This method is generated by Databricks SDK Code Generator.
func (a *AppsAPI) ListDeploymentsAll(ctx context.Context, request ListAppDeploymentsRequest) ([]AppDeployment, error) {
	iterator := a.ListDeployments(ctx, request)
	return listing.ToSlice[AppDeployment](ctx, iterator)
}

// List app deployments.
//
// Lists all app deployments for the app with the supplied name.
func (a *AppsAPI) ListDeploymentsByAppName(ctx context.Context, appName string) (*ListAppDeploymentsResponse, error) {
	return a.appsImpl.ListDeployments(ctx, ListAppDeploymentsRequest{
		AppName: appName,
	})
}

// Start an app.
//
// Start the last active deployment of the app in the workspace.
func (a *AppsAPI) Start(ctx context.Context, startAppRequest StartAppRequest) (*WaitGetAppActive[App], error) {
	app, err := a.appsImpl.Start(ctx, startAppRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetAppActive[App]{
		Response: app,
		Name:     app.Name,
		Poll: func(timeout time.Duration, callback func(*App)) (*App, error) {
			return a.WaitGetAppActive(ctx, app.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.Start] and waits to reach ACTIVE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[App](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.Start].Get() or [AppsAPI.WaitGetAppActive]
func (a *AppsAPI) StartAndWait(ctx context.Context, startAppRequest StartAppRequest, options ...retries.Option[App]) (*App, error) {
	wait, err := a.Start(ctx, startAppRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[App]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *App) {
		for _, o := range options {
			o(&retries.Info[App]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Stop an app.
//
// Stops the active deployment of the app in the workspace.
func (a *AppsAPI) Stop(ctx context.Context, stopAppRequest StopAppRequest) (*WaitGetAppStopped[App], error) {
	app, err := a.appsImpl.Stop(ctx, stopAppRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetAppStopped[App]{
		Response: app,
		Name:     app.Name,
		Poll: func(timeout time.Duration, callback func(*App)) (*App, error) {
			return a.WaitGetAppStopped(ctx, app.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [AppsAPI.Stop] and waits to reach STOPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[App](60*time.Minute) functional option.
//
// Deprecated: use [AppsAPI.Stop].Get() or [AppsAPI.WaitGetAppStopped]
func (a *AppsAPI) StopAndWait(ctx context.Context, stopAppRequest StopAppRequest, options ...retries.Option[App]) (*App, error) {
	wait, err := a.Stop(ctx, stopAppRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[App]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *App) {
		for _, o := range options {
			o(&retries.Info[App]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}
