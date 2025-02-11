// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Serving Endpoints, Serving Endpoints Data Plane Preview, Serving Endpoints Preview, etc.
package servingpreview

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
)

type ServingEndpointsInterface interface {
}

func NewServingEndpoints(client *client.DatabricksClient) *ServingEndpointsAPI {
	return &ServingEndpointsAPI{
		servingEndpointsImpl: servingEndpointsImpl{
			client: client,
		},
	}
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
type ServingEndpointsAPI struct {
	servingEndpointsImpl
}

type ServingEndpointsDataPlanePreviewInterface interface {

	// Query a serving endpoint.
	Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error)
}

func NewServingEndpointsDataPlanePreview(client *client.DatabricksClient) *ServingEndpointsDataPlanePreviewAPI {
	return &ServingEndpointsDataPlanePreviewAPI{
		servingEndpointsDataPlanePreviewImpl: servingEndpointsDataPlanePreviewImpl{
			client: client,
		},
	}
}

// Serving endpoints DataPlane provides a set of operations to interact with
// data plane endpoints for Serving endpoints service.
type ServingEndpointsDataPlanePreviewAPI struct {
	servingEndpointsDataPlanePreviewImpl
}

type ServingEndpointsPreviewInterface interface {

	// Get build logs for a served model.
	//
	// Retrieves the build logs associated with the provided served model.
	BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error)

	// Get build logs for a served model.
	//
	// Retrieves the build logs associated with the provided served model.
	BuildLogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*BuildLogsResponse, error)

	// Create a new serving endpoint.
	Create(ctx context.Context, request CreateServingEndpoint) (*ServingEndpointDetailed, error)

	// Delete a serving endpoint.
	Delete(ctx context.Context, request DeleteServingEndpointRequest) error

	// Delete a serving endpoint.
	DeleteByName(ctx context.Context, name string) error

	// Get metrics of a serving endpoint.
	//
	// Retrieves the metrics associated with the provided serving endpoint in either
	// Prometheus or OpenMetrics exposition format.
	ExportMetrics(ctx context.Context, request ExportMetricsRequest) (*ExportMetricsResponse, error)

	// Get metrics of a serving endpoint.
	//
	// Retrieves the metrics associated with the provided serving endpoint in either
	// Prometheus or OpenMetrics exposition format.
	ExportMetricsByName(ctx context.Context, name string) (*ExportMetricsResponse, error)

	// Get a single serving endpoint.
	//
	// Retrieves the details for a single serving endpoint.
	Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error)

	// Get a single serving endpoint.
	//
	// Retrieves the details for a single serving endpoint.
	GetByName(ctx context.Context, name string) (*ServingEndpointDetailed, error)

	// Get the schema for a serving endpoint.
	//
	// Get the query schema of the serving endpoint in OpenAPI format. The schema
	// contains information for the supported paths, input and output format and
	// datatypes.
	GetOpenApi(ctx context.Context, request GetOpenApiRequest) (*GetOpenApiResponse, error)

	// Get the schema for a serving endpoint.
	//
	// Get the query schema of the serving endpoint in OpenAPI format. The schema
	// contains information for the supported paths, input and output format and
	// datatypes.
	GetOpenApiByName(ctx context.Context, name string) (*GetOpenApiResponse, error)

	// Get serving endpoint permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetServingEndpointPermissionLevelsRequest) (*GetServingEndpointPermissionLevelsResponse, error)

	// Get serving endpoint permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByServingEndpointId(ctx context.Context, servingEndpointId string) (*GetServingEndpointPermissionLevelsResponse, error)

	// Get serving endpoint permissions.
	//
	// Gets the permissions of a serving endpoint. Serving endpoints can inherit
	// permissions from their root object.
	GetPermissions(ctx context.Context, request GetServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)

	// Get serving endpoint permissions.
	//
	// Gets the permissions of a serving endpoint. Serving endpoints can inherit
	// permissions from their root object.
	GetPermissionsByServingEndpointId(ctx context.Context, servingEndpointId string) (*ServingEndpointPermissions, error)

	// Make external services call using the credentials stored in UC Connection.
	HttpRequest(ctx context.Context, request ExternalFunctionRequest) (*HttpRequestResponse, error)

	// Get all serving endpoints.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context) listing.Iterator[ServingEndpoint]

	// Get all serving endpoints.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context) ([]ServingEndpoint, error)

	// Get the latest logs for a served model.
	//
	// Retrieves the service logs associated with the provided served model.
	Logs(ctx context.Context, request LogsRequest) (*ServerLogsResponse, error)

	// Get the latest logs for a served model.
	//
	// Retrieves the service logs associated with the provided served model.
	LogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*ServerLogsResponse, error)

	// Update tags of a serving endpoint.
	//
	// Used to batch add and delete tags from a serving endpoint with a single API
	// call.
	Patch(ctx context.Context, request PatchServingEndpointTags) (*EndpointTags, error)

	// Update rate limits of a serving endpoint.
	//
	// Used to update the rate limits of a serving endpoint. NOTE: Only foundation
	// model endpoints are currently supported. For external models, use AI Gateway
	// to manage rate limits.
	Put(ctx context.Context, request PutRequest) (*PutResponse, error)

	// Update AI Gateway of a serving endpoint.
	//
	// Used to update the AI Gateway of a serving endpoint. NOTE: Only external
	// model and provisioned throughput endpoints are currently supported.
	PutAiGateway(ctx context.Context, request PutAiGatewayRequest) (*PutAiGatewayResponse, error)

	// Query a serving endpoint.
	Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error)

	// Set serving endpoint permissions.
	//
	// Sets permissions on an object, replacing existing permissions if they exist.
	// Deletes all direct permissions if none are specified. Objects can inherit
	// permissions from their root object.
	SetPermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)

	// Update config of a serving endpoint.
	//
	// Updates any combination of the serving endpoint's served entities, the
	// compute configuration of those served entities, and the endpoint's traffic
	// config. An endpoint that already has an update in progress can not be updated
	// until the current update completes or fails.
	UpdateConfig(ctx context.Context, request EndpointCoreConfigInput) (*ServingEndpointDetailed, error)

	// Update serving endpoint permissions.
	//
	// Updates the permissions on a serving endpoint. Serving endpoints can inherit
	// permissions from their root object.
	UpdatePermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)
}

func NewServingEndpointsPreview(client *client.DatabricksClient) *ServingEndpointsPreviewAPI {
	return &ServingEndpointsPreviewAPI{
		servingEndpointsPreviewImpl: servingEndpointsPreviewImpl{
			client: client,
		},
	}
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
type ServingEndpointsPreviewAPI struct {
	servingEndpointsPreviewImpl
}

// Get build logs for a served model.
//
// Retrieves the build logs associated with the provided served model.
func (a *ServingEndpointsPreviewAPI) BuildLogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*BuildLogsResponse, error) {
	return a.servingEndpointsPreviewImpl.BuildLogs(ctx, BuildLogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Delete a serving endpoint.
func (a *ServingEndpointsPreviewAPI) DeleteByName(ctx context.Context, name string) error {
	return a.servingEndpointsPreviewImpl.Delete(ctx, DeleteServingEndpointRequest{
		Name: name,
	})
}

// Get metrics of a serving endpoint.
//
// Retrieves the metrics associated with the provided serving endpoint in either
// Prometheus or OpenMetrics exposition format.
func (a *ServingEndpointsPreviewAPI) ExportMetricsByName(ctx context.Context, name string) (*ExportMetricsResponse, error) {
	return a.servingEndpointsPreviewImpl.ExportMetrics(ctx, ExportMetricsRequest{
		Name: name,
	})
}

// Get a single serving endpoint.
//
// Retrieves the details for a single serving endpoint.
func (a *ServingEndpointsPreviewAPI) GetByName(ctx context.Context, name string) (*ServingEndpointDetailed, error) {
	return a.servingEndpointsPreviewImpl.Get(ctx, GetServingEndpointRequest{
		Name: name,
	})
}

// Get the schema for a serving endpoint.
//
// Get the query schema of the serving endpoint in OpenAPI format. The schema
// contains information for the supported paths, input and output format and
// datatypes.
func (a *ServingEndpointsPreviewAPI) GetOpenApiByName(ctx context.Context, name string) (*GetOpenApiResponse, error) {
	return a.servingEndpointsPreviewImpl.GetOpenApi(ctx, GetOpenApiRequest{
		Name: name,
	})
}

// Get serving endpoint permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ServingEndpointsPreviewAPI) GetPermissionLevelsByServingEndpointId(ctx context.Context, servingEndpointId string) (*GetServingEndpointPermissionLevelsResponse, error) {
	return a.servingEndpointsPreviewImpl.GetPermissionLevels(ctx, GetServingEndpointPermissionLevelsRequest{
		ServingEndpointId: servingEndpointId,
	})
}

// Get serving endpoint permissions.
//
// Gets the permissions of a serving endpoint. Serving endpoints can inherit
// permissions from their root object.
func (a *ServingEndpointsPreviewAPI) GetPermissionsByServingEndpointId(ctx context.Context, servingEndpointId string) (*ServingEndpointPermissions, error) {
	return a.servingEndpointsPreviewImpl.GetPermissions(ctx, GetServingEndpointPermissionsRequest{
		ServingEndpointId: servingEndpointId,
	})
}

// Get the latest logs for a served model.
//
// Retrieves the service logs associated with the provided served model.
func (a *ServingEndpointsPreviewAPI) LogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*ServerLogsResponse, error) {
	return a.servingEndpointsPreviewImpl.Logs(ctx, LogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}
