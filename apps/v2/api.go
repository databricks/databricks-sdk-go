// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Apps run directly on a customer’s Databricks instance, integrate with their data, use and extend Databricks services, and enable users to interact through single sign-on.
package apps

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type appsBaseClient struct {
	appsImpl
}

// Create an app.
//
// Creates a new app.
func (a *appsBaseClient) Create(ctx context.Context, createAppRequest CreateAppRequest) (*AppsCreateWaiter, error) {
	app, err := a.appsImpl.Create(ctx, createAppRequest)
	if err != nil {
		return nil, err
	}
	return &AppsCreateWaiter{
		RawResponse: app,
		name:        app.Name,
		service:     a,
	}, nil
}

type AppsCreateWaiter struct {
	// RawResponse is the raw response of the Create call.
	RawResponse *App
	service     *appsBaseClient
	name        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *AppsCreateWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*App, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[App](ctx, opts.Timeout, func() (*App, *retries.Err) {
		app, err := w.service.Get(ctx, GetAppRequest{
			Name: w.name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := app.ComputeStatus.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if app.ComputeStatus != nil {
			statusMessage = app.ComputeStatus.Message
		}
		switch status {
		case ComputeStateActive: // target state
			return app, nil
		case ComputeStateError, ComputeStateStopped:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ComputeStateActive, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})

}

// Delete an app.
//
// Deletes an app.
func (a *appsBaseClient) DeleteByName(ctx context.Context, name string) (*App, error) {
	return a.appsImpl.Delete(ctx, DeleteAppRequest{
		Name: name,
	})
}

// Create an app deployment.
//
// Creates an app deployment for the app with the supplied name.
func (a *appsBaseClient) Deploy(ctx context.Context, createAppDeploymentRequest CreateAppDeploymentRequest) (*AppsDeployWaiter, error) {
	appDeployment, err := a.appsImpl.Deploy(ctx, createAppDeploymentRequest)
	if err != nil {
		return nil, err
	}
	return &AppsDeployWaiter{
		RawResponse:  appDeployment,
		appName:      createAppDeploymentRequest.AppName,
		deploymentId: appDeployment.DeploymentId,
		service:      a,
	}, nil
}

type AppsDeployWaiter struct {
	// RawResponse is the raw response of the Deploy call.
	RawResponse  *AppDeployment
	service      *appsBaseClient
	appName      string
	deploymentId string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *AppsDeployWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*AppDeployment, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[AppDeployment](ctx, opts.Timeout, func() (*AppDeployment, *retries.Err) {
		appDeployment, err := w.service.GetDeployment(ctx, GetAppDeploymentRequest{
			AppName:      w.appName,
			DeploymentId: w.deploymentId,
		})
		if err != nil {
			return nil, retries.Halt(err)
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

// Get an app.
//
// Retrieves information for the app with the supplied name.
func (a *appsBaseClient) GetByName(ctx context.Context, name string) (*App, error) {
	return a.appsImpl.Get(ctx, GetAppRequest{
		Name: name,
	})
}

// Get an app deployment.
//
// Retrieves information for the app deployment with the supplied name and
// deployment id.
func (a *appsBaseClient) GetDeploymentByAppNameAndDeploymentId(ctx context.Context, appName string, deploymentId string) (*AppDeployment, error) {
	return a.appsImpl.GetDeployment(ctx, GetAppDeploymentRequest{
		AppName:      appName,
		DeploymentId: deploymentId,
	})
}

// Get app permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *appsBaseClient) GetPermissionLevelsByAppName(ctx context.Context, appName string) (*GetAppPermissionLevelsResponse, error) {
	return a.appsImpl.GetPermissionLevels(ctx, GetAppPermissionLevelsRequest{
		AppName: appName,
	})
}

// Get app permissions.
//
// Gets the permissions of an app. Apps can inherit permissions from their root
// object.
func (a *appsBaseClient) GetPermissionsByAppName(ctx context.Context, appName string) (*AppPermissions, error) {
	return a.appsImpl.GetPermissions(ctx, GetAppPermissionsRequest{
		AppName: appName,
	})
}

// List app deployments.
//
// Lists all app deployments for the app with the supplied name.
func (a *appsBaseClient) ListDeploymentsByAppName(ctx context.Context, appName string) (*ListAppDeploymentsResponse, error) {
	return a.appsImpl.internalListDeployments(ctx, ListAppDeploymentsRequest{
		AppName: appName,
	})
}

// Start an app.
//
// Start the last active deployment of the app in the workspace.
func (a *appsBaseClient) Start(ctx context.Context, startAppRequest StartAppRequest) (*AppsStartWaiter, error) {
	app, err := a.appsImpl.Start(ctx, startAppRequest)
	if err != nil {
		return nil, err
	}
	return &AppsStartWaiter{
		RawResponse: app,
		name:        app.Name,
		service:     a,
	}, nil
}

type AppsStartWaiter struct {
	// RawResponse is the raw response of the Start call.
	RawResponse *App
	service     *appsBaseClient
	name        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *AppsStartWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*App, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[App](ctx, opts.Timeout, func() (*App, *retries.Err) {
		app, err := w.service.Get(ctx, GetAppRequest{
			Name: w.name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := app.ComputeStatus.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if app.ComputeStatus != nil {
			statusMessage = app.ComputeStatus.Message
		}
		switch status {
		case ComputeStateActive: // target state
			return app, nil
		case ComputeStateError, ComputeStateStopped:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ComputeStateActive, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})

}

// Stop an app.
//
// Stops the active deployment of the app in the workspace.
func (a *appsBaseClient) Stop(ctx context.Context, stopAppRequest StopAppRequest) (*AppsStopWaiter, error) {
	app, err := a.appsImpl.Stop(ctx, stopAppRequest)
	if err != nil {
		return nil, err
	}
	return &AppsStopWaiter{
		RawResponse: app,
		name:        app.Name,
		service:     a,
	}, nil
}

type AppsStopWaiter struct {
	// RawResponse is the raw response of the Stop call.
	RawResponse *App
	service     *appsBaseClient
	name        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *AppsStopWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*App, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[App](ctx, opts.Timeout, func() (*App, *retries.Err) {
		app, err := w.service.Get(ctx, GetAppRequest{
			Name: w.name,
		})
		if err != nil {
			return nil, retries.Halt(err)
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
