// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"context"
)

// Lakehouse Apps run directly on a customer’s Databricks instance, integrate
// with their data, use and extend Databricks services, and enable users to
// interact through single sign-on.
type AppsService interface {

	// Create and deploy an application.
	//
	// Creates and deploys an application.
	Create(ctx context.Context, request DeployAppRequest) (*DeploymentStatus, error)

	// Delete an application.
	//
	// Delete an application definition
	DeleteApp(ctx context.Context, request DeleteAppRequest) (*DeleteAppResponse, error)

	// Get definition for an application.
	//
	// Get an application definition
	GetApp(ctx context.Context, request GetAppRequest) (*GetAppResponse, error)

	// Get deployment status for an application.
	//
	// Get deployment status for an application
	GetAppDeploymentStatus(ctx context.Context, request GetAppDeploymentStatusRequest) (*DeploymentStatus, error)

	// List all applications.
	//
	// List all available applications
	GetApps(ctx context.Context) (*ListAppsResponse, error)

	// Get deployment events for an application.
	//
	// Get deployment events for an application
	GetEvents(ctx context.Context, request GetEventsRequest) (*ListAppEventsResponse, error)
}

// The Serving Endpoints API allows you to create, update, and delete model
// serving endpoints.
//
// You can use a serving endpoint to serve models from the Databricks Model
// Registry or from Unity Catalog. Endpoints expose the underlying models as
// scalable REST API endpoints using serverless compute. This means the
// endpoints and associated compute resources are fully managed by Databricks
// and will not appear in your cloud account. A serving endpoint can consist of
// one or more MLflow models from the Databricks Model Registry, called served
// entities. A serving endpoint can have at most ten served entities. You can
// configure traffic settings to define how requests should be routed to your
// served entities behind an endpoint. Additionally, you can configure the scale
// of resources that should be applied to each served entity.
type ServingEndpointsService interface {

	// Get build logs for a served model.
	//
	// Retrieves the build logs associated with the provided served model.
	BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error)

	// Create a new serving endpoint.
	Create(ctx context.Context, request CreateServingEndpoint) (*ServingEndpointDetailed, error)

	// Delete a serving endpoint.
	Delete(ctx context.Context, request DeleteServingEndpointRequest) error

	// Get metrics of a serving endpoint.
	//
	// Retrieves the metrics associated with the provided serving endpoint in
	// either Prometheus or OpenMetrics exposition format.
	ExportMetrics(ctx context.Context, request ExportMetricsRequest) error

	// Get a single serving endpoint.
	//
	// Retrieves the details for a single serving endpoint.
	Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error)

	// Get serving endpoint permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetServingEndpointPermissionLevelsRequest) (*GetServingEndpointPermissionLevelsResponse, error)

	// Get serving endpoint permissions.
	//
	// Gets the permissions of a serving endpoint. Serving endpoints can inherit
	// permissions from their root object.
	GetPermissions(ctx context.Context, request GetServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)

	// Get all serving endpoints.
	//
	// Use ListAll() to get all ServingEndpoint instances
	List(ctx context.Context) (*ListEndpointsResponse, error)

	// Get the latest logs for a served model.
	//
	// Retrieves the service logs associated with the provided served model.
	Logs(ctx context.Context, request LogsRequest) (*ServerLogsResponse, error)

	// Update tags of a serving endpoint.
	//
	// Used to batch add and delete tags from a serving endpoint with a single
	// API call.
	Patch(ctx context.Context, request PatchServingEndpointTags) ([]EndpointTag, error)

	// Update rate limits of a serving endpoint.
	//
	// Used to update the rate limits of a serving endpoint. NOTE: only external
	// and foundation model endpoints are supported as of now.
	Put(ctx context.Context, request PutRequest) (*PutResponse, error)

	// Query a serving endpoint.
	Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error)

	// Set serving endpoint permissions.
	//
	// Sets permissions on a serving endpoint. Serving endpoints can inherit
	// permissions from their root object.
	SetPermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)

	// Update config of a serving endpoint.
	//
	// Updates any combination of the serving endpoint's served entities, the
	// compute configuration of those served entities, and the endpoint's
	// traffic config. An endpoint that already has an update in progress can
	// not be updated until the current update completes or fails.
	UpdateConfig(ctx context.Context, request EndpointCoreConfigInput) (*ServingEndpointDetailed, error)

	// Update serving endpoint permissions.
	//
	// Updates the permissions on a serving endpoint. Serving endpoints can
	// inherit permissions from their root object.
	UpdatePermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)
}
