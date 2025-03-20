// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Apps run directly on a customer’s Databricks instance, integrate with their data, use and extend Databricks services, and enable users to interact through single sign-on.
package apps

import (
	"context"
)

type appsBaseClient struct {
	appsImpl
}

// Delete an app.
//
// Deletes an app.
func (a *appsBaseClient) DeleteByName(ctx context.Context, name string) (*App, error) {
	return a.appsImpl.Delete(ctx, DeleteAppRequest{
		Name: name,
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
