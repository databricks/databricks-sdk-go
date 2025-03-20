// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Serving Endpoints, Serving Endpoints Data Plane, etc.
package serving

import (
	"context"
)

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

type ServingEndpointsDataPlaneAPI struct {
	servingEndpointsDataPlaneImpl
}
