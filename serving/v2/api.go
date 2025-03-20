// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Serving Endpoints, Serving Endpoints Data Plane, etc.
package serving

import (
	"context"
)

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
type ServingEndpointsAPI struct {
	servingEndpointsImpl
}

// Get build logs for a served model.
//
// Retrieves the build logs associated with the provided served model.
func (a *ServingEndpointsAPI) BuildLogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*BuildLogsResponse, error) {
	return a.servingEndpointsImpl.BuildLogs(ctx, BuildLogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Delete a serving endpoint.
func (a *ServingEndpointsAPI) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.servingEndpointsImpl.Delete(ctx, DeleteServingEndpointRequest{
		Name: name,
	})
}

// Get metrics of a serving endpoint.
//
// Retrieves the metrics associated with the provided serving endpoint in either
// Prometheus or OpenMetrics exposition format.
func (a *ServingEndpointsAPI) ExportMetricsByName(ctx context.Context, name string) (*ExportMetricsResponse, error) {
	return a.servingEndpointsImpl.ExportMetrics(ctx, ExportMetricsRequest{
		Name: name,
	})
}

// Get a single serving endpoint.
//
// Retrieves the details for a single serving endpoint.
func (a *ServingEndpointsAPI) GetByName(ctx context.Context, name string) (*ServingEndpointDetailed, error) {
	return a.servingEndpointsImpl.Get(ctx, GetServingEndpointRequest{
		Name: name,
	})
}

// Get the schema for a serving endpoint.
//
// Get the query schema of the serving endpoint in OpenAPI format. The schema
// contains information for the supported paths, input and output format and
// datatypes.
func (a *ServingEndpointsAPI) GetOpenApiByName(ctx context.Context, name string) (*GetOpenApiResponse, error) {
	return a.servingEndpointsImpl.GetOpenApi(ctx, GetOpenApiRequest{
		Name: name,
	})
}

// Get serving endpoint permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ServingEndpointsAPI) GetPermissionLevelsByServingEndpointId(ctx context.Context, servingEndpointId string) (*GetServingEndpointPermissionLevelsResponse, error) {
	return a.servingEndpointsImpl.GetPermissionLevels(ctx, GetServingEndpointPermissionLevelsRequest{
		ServingEndpointId: servingEndpointId,
	})
}

// Get serving endpoint permissions.
//
// Gets the permissions of a serving endpoint. Serving endpoints can inherit
// permissions from their root object.
func (a *ServingEndpointsAPI) GetPermissionsByServingEndpointId(ctx context.Context, servingEndpointId string) (*ServingEndpointPermissions, error) {
	return a.servingEndpointsImpl.GetPermissions(ctx, GetServingEndpointPermissionsRequest{
		ServingEndpointId: servingEndpointId,
	})
}

// Get the latest logs for a served model.
//
// Retrieves the service logs associated with the provided served model.
func (a *ServingEndpointsAPI) LogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*ServerLogsResponse, error) {
	return a.servingEndpointsImpl.Logs(ctx, LogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Serving endpoints DataPlane provides a set of operations to interact with
// data plane endpoints for Serving endpoints service.
type ServingEndpointsDataPlaneAPI struct {
	servingEndpointsDataPlaneImpl
}
