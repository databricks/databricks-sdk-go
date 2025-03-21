// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Serving Endpoints, Serving Endpoints Data Plane, etc.
package serving

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type servingEndpointsBaseClient struct {
	servingEndpointsImpl
}

// Get build logs for a served model.
//
// Retrieves the build logs associated with the provided served model.
func (a *servingEndpointsBaseClient) BuildLogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*BuildLogsResponse, error) {
	return a.servingEndpointsImpl.BuildLogs(ctx, BuildLogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Create a new serving endpoint.
func (a *servingEndpointsBaseClient) Create(ctx context.Context, createServingEndpoint CreateServingEndpoint) (*ServingEndpointsCreateWaiter, error) {
	servingEndpointDetailed, err := a.servingEndpointsImpl.Create(ctx, createServingEndpoint)
	if err != nil {
		return nil, err
	}
	return &ServingEndpointsCreateWaiter{
		RawResponse: servingEndpointDetailed,
		name:        servingEndpointDetailed.Name,
		service:     a,
	}, nil
}

type ServingEndpointsCreateWaiter struct {
	// RawResponse is the raw response of the Create call.
	RawResponse *ServingEndpointDetailed
	service     *servingEndpointsBaseClient
	name        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *ServingEndpointsCreateWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*ServingEndpointDetailed, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[ServingEndpointDetailed](ctx, opts.Timeout, func() (*ServingEndpointDetailed, *retries.Err) {
		servingEndpointDetailed, err := w.service.Get(ctx, GetServingEndpointRequest{
			Name: w.name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := servingEndpointDetailed.State.ConfigUpdate
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case EndpointStateConfigUpdateNotUpdating: // target state
			return servingEndpointDetailed, nil
		case EndpointStateConfigUpdateUpdateFailed, EndpointStateConfigUpdateUpdateCanceled:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				EndpointStateConfigUpdateNotUpdating, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})

}

// Delete a serving endpoint.
func (a *servingEndpointsBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.servingEndpointsImpl.Delete(ctx, DeleteServingEndpointRequest{
		Name: name,
	})
}

// Get metrics of a serving endpoint.
//
// Retrieves the metrics associated with the provided serving endpoint in either
// Prometheus or OpenMetrics exposition format.
func (a *servingEndpointsBaseClient) ExportMetricsByName(ctx context.Context, name string) (*ExportMetricsResponse, error) {
	return a.servingEndpointsImpl.ExportMetrics(ctx, ExportMetricsRequest{
		Name: name,
	})
}

// Get a single serving endpoint.
//
// Retrieves the details for a single serving endpoint.
func (a *servingEndpointsBaseClient) GetByName(ctx context.Context, name string) (*ServingEndpointDetailed, error) {
	return a.servingEndpointsImpl.Get(ctx, GetServingEndpointRequest{
		Name: name,
	})
}

// Get the schema for a serving endpoint.
//
// Get the query schema of the serving endpoint in OpenAPI format. The schema
// contains information for the supported paths, input and output format and
// datatypes.
func (a *servingEndpointsBaseClient) GetOpenApiByName(ctx context.Context, name string) (*GetOpenApiResponse, error) {
	return a.servingEndpointsImpl.GetOpenApi(ctx, GetOpenApiRequest{
		Name: name,
	})
}

// Get serving endpoint permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *servingEndpointsBaseClient) GetPermissionLevelsByServingEndpointId(ctx context.Context, servingEndpointId string) (*GetServingEndpointPermissionLevelsResponse, error) {
	return a.servingEndpointsImpl.GetPermissionLevels(ctx, GetServingEndpointPermissionLevelsRequest{
		ServingEndpointId: servingEndpointId,
	})
}

// Get serving endpoint permissions.
//
// Gets the permissions of a serving endpoint. Serving endpoints can inherit
// permissions from their root object.
func (a *servingEndpointsBaseClient) GetPermissionsByServingEndpointId(ctx context.Context, servingEndpointId string) (*ServingEndpointPermissions, error) {
	return a.servingEndpointsImpl.GetPermissions(ctx, GetServingEndpointPermissionsRequest{
		ServingEndpointId: servingEndpointId,
	})
}

// Get the latest logs for a served model.
//
// Retrieves the service logs associated with the provided served model.
func (a *servingEndpointsBaseClient) LogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*ServerLogsResponse, error) {
	return a.servingEndpointsImpl.Logs(ctx, LogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Update config of a serving endpoint.
//
// Updates any combination of the serving endpoint's served entities, the
// compute configuration of those served entities, and the endpoint's traffic
// config. An endpoint that already has an update in progress can not be updated
// until the current update completes or fails.
func (a *servingEndpointsBaseClient) UpdateConfig(ctx context.Context, endpointCoreConfigInput EndpointCoreConfigInput) (*ServingEndpointsUpdateConfigWaiter, error) {
	servingEndpointDetailed, err := a.servingEndpointsImpl.UpdateConfig(ctx, endpointCoreConfigInput)
	if err != nil {
		return nil, err
	}
	return &ServingEndpointsUpdateConfigWaiter{
		RawResponse: servingEndpointDetailed,
		name:        servingEndpointDetailed.Name,
		service:     a,
	}, nil
}

type ServingEndpointsUpdateConfigWaiter struct {
	// RawResponse is the raw response of the UpdateConfig call.
	RawResponse *ServingEndpointDetailed
	service     *servingEndpointsBaseClient
	name        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *ServingEndpointsUpdateConfigWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*ServingEndpointDetailed, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[ServingEndpointDetailed](ctx, opts.Timeout, func() (*ServingEndpointDetailed, *retries.Err) {
		servingEndpointDetailed, err := w.service.Get(ctx, GetServingEndpointRequest{
			Name: w.name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := servingEndpointDetailed.State.ConfigUpdate
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case EndpointStateConfigUpdateNotUpdating: // target state
			return servingEndpointDetailed, nil
		case EndpointStateConfigUpdateUpdateFailed, EndpointStateConfigUpdateUpdateCanceled:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				EndpointStateConfigUpdateNotUpdating, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})

}

type servingEndpointsDataPlaneBaseClient struct {
	servingEndpointsDataPlaneImpl
}
