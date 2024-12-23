// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"context"
)

// Apps run directly on a customer’s Databricks instance, integrate with their
// data, use and extend Databricks services, and enable users to interact
// through single sign-on.
type AppsService interface {

	// Create an app.
	//
	// Creates a new app.
	Create(ctx context.Context, request CreateAppRequest) (*App, error)

	// Delete an app.
	//
	// Deletes an app.
	Delete(ctx context.Context, request DeleteAppRequest) (*App, error)

	// Create an app deployment.
	//
	// Creates an app deployment for the app with the supplied name.
	Deploy(ctx context.Context, request CreateAppDeploymentRequest) (*AppDeployment, error)

	// Get an app.
	//
	// Retrieves information for the app with the supplied name.
	Get(ctx context.Context, request GetAppRequest) (*App, error)

	// Get an app deployment.
	//
	// Retrieves information for the app deployment with the supplied name and
	// deployment id.
	GetDeployment(ctx context.Context, request GetAppDeploymentRequest) (*AppDeployment, error)

	// Get app permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetAppPermissionLevelsRequest) (*GetAppPermissionLevelsResponse, error)

	// Get app permissions.
	//
	// Gets the permissions of an app. Apps can inherit permissions from their
	// root object.
	GetPermissions(ctx context.Context, request GetAppPermissionsRequest) (*AppPermissions, error)

	// List apps.
	//
	// Lists all apps in the workspace.
	//
	// Use ListAll() to get all App instances, which will iterate over every result page.
	List(ctx context.Context, request ListAppsRequest) (*ListAppsResponse, error)

	// List app deployments.
	//
	// Lists all app deployments for the app with the supplied name.
	//
	// Use ListDeploymentsAll() to get all AppDeployment instances, which will iterate over every result page.
	ListDeployments(ctx context.Context, request ListAppDeploymentsRequest) (*ListAppDeploymentsResponse, error)

	// Set app permissions.
	//
	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error)

	// Start an app.
	//
	// Start the last active deployment of the app in the workspace.
	Start(ctx context.Context, request StartAppRequest) (*App, error)

	// Stop an app.
	//
	// Stops the active deployment of the app in the workspace.
	Stop(ctx context.Context, request StopAppRequest) (*App, error)

	// Update an app.
	//
	// Updates the app with the supplied name.
	Update(ctx context.Context, request UpdateAppRequest) (*App, error)

	// Update app permissions.
	//
	// Updates the permissions on an app. Apps can inherit permissions from
	// their root object.
	UpdatePermissions(ctx context.Context, request AppPermissionsRequest) (*AppPermissions, error)
}
