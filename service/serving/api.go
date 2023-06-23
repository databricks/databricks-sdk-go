// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Serving Endpoints API allows you to create, update, and delete model serving endpoints.
package serving

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewServingEndpoints(client *client.DatabricksClient) *ServingEndpointsAPI {
	return &ServingEndpointsAPI{
		impl: &servingEndpointsImpl{
			client: client,
		},
	}
}

// The Serving Endpoints API allows you to create, update, and delete model
// serving endpoints.
//
// You can use a serving endpoint to serve models from the Databricks Model
// Registry. Endpoints expose the underlying models as scalable REST API
// endpoints using serverless compute. This means the endpoints and associated
// compute resources are fully managed by Databricks and will not appear in your
// cloud account. A serving endpoint can consist of one or more MLflow models
// from the Databricks Model Registry, called served models. A serving endpoint
// can have at most ten served models. You can configure traffic settings to
// define how requests should be routed to your served models behind an
// endpoint. Additionally, you can configure the scale of resources that should
// be applied to each served model.
type ServingEndpointsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ServingEndpointsService)
	impl ServingEndpointsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ServingEndpointsAPI) WithImpl(impl ServingEndpointsService) *ServingEndpointsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level ServingEndpoints API implementation
func (a *ServingEndpointsAPI) Impl() ServingEndpointsService {
	return a.impl
}

// WaitGetServingEndpointNotUpdating repeatedly calls [ServingEndpointsAPI.Get] and waits to reach NOT_UPDATING state
func (a *ServingEndpointsAPI) WaitGetServingEndpointNotUpdating(ctx context.Context, name string,
	timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[ServingEndpointDetailed](ctx, timeout, func() (*ServingEndpointDetailed, *retries.Err) {
		servingEndpointDetailed, err := a.Get(ctx, GetServingEndpointRequest{
			Name: name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(servingEndpointDetailed)
		}
		status := servingEndpointDetailed.State.ConfigUpdate
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case EndpointStateConfigUpdateNotUpdating: // target state
			return servingEndpointDetailed, nil
		case EndpointStateConfigUpdateUpdateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				EndpointStateConfigUpdateNotUpdating, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetServingEndpointNotUpdating is a wrapper that calls [ServingEndpointsAPI.WaitGetServingEndpointNotUpdating] and waits to reach NOT_UPDATING state.
type WaitGetServingEndpointNotUpdating[R any] struct {
	Response *R
	Name     string `json:"name"`
	poll     func(time.Duration, func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error)
	callback func(*ServingEndpointDetailed)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetServingEndpointNotUpdating[R]) OnProgress(callback func(*ServingEndpointDetailed)) *WaitGetServingEndpointNotUpdating[R] {
	w.callback = callback
	return w
}

// Get the ServingEndpointDetailed with the default timeout of 20 minutes.
func (w *WaitGetServingEndpointNotUpdating[R]) Get() (*ServingEndpointDetailed, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the ServingEndpointDetailed with custom timeout.
func (w *WaitGetServingEndpointNotUpdating[R]) GetWithTimeout(timeout time.Duration) (*ServingEndpointDetailed, error) {
	return w.poll(timeout, w.callback)
}

// Retrieve the logs associated with building the model's environment for a
// given serving endpoint's served model.
//
// Retrieves the build logs associated with the provided served model.
func (a *ServingEndpointsAPI) BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error) {
	return a.impl.BuildLogs(ctx, request)
}

// Retrieve the logs associated with building the model's environment for a
// given serving endpoint's served model.
//
// Retrieves the build logs associated with the provided served model.
func (a *ServingEndpointsAPI) BuildLogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*BuildLogsResponse, error) {
	return a.impl.BuildLogs(ctx, BuildLogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Create a new serving endpoint.
func (a *ServingEndpointsAPI) Create(ctx context.Context, createServingEndpoint CreateServingEndpoint) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error) {
	servingEndpointDetailed, err := a.impl.Create(ctx, createServingEndpoint)
	if err != nil {
		return nil, err
	}
	return &WaitGetServingEndpointNotUpdating[ServingEndpointDetailed]{
		Response: servingEndpointDetailed,
		Name:     servingEndpointDetailed.Name,
		poll: func(timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error) {
			return a.WaitGetServingEndpointNotUpdating(ctx, servingEndpointDetailed.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ServingEndpointsAPI.Create] and waits to reach NOT_UPDATING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
//
// Deprecated: use [ServingEndpointsAPI.Create].Get() or [ServingEndpointsAPI.WaitGetServingEndpointNotUpdating]
func (a *ServingEndpointsAPI) CreateAndWait(ctx context.Context, createServingEndpoint CreateServingEndpoint, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error) {
	wait, err := a.Create(ctx, createServingEndpoint)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ServingEndpointDetailed]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ServingEndpointDetailed) {
		for _, o := range options {
			o(&retries.Info[ServingEndpointDetailed]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete a serving endpoint.
func (a *ServingEndpointsAPI) Delete(ctx context.Context, request DeleteServingEndpointRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a serving endpoint.
func (a *ServingEndpointsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.impl.Delete(ctx, DeleteServingEndpointRequest{
		Name: name,
	})
}

// Retrieve the metrics associated with a serving endpoint.
//
// Retrieves the metrics associated with the provided serving endpoint in either
// Prometheus or OpenMetrics exposition format.
func (a *ServingEndpointsAPI) ExportMetrics(ctx context.Context, request ExportMetricsRequest) error {
	return a.impl.ExportMetrics(ctx, request)
}

// Retrieve the metrics associated with a serving endpoint.
//
// Retrieves the metrics associated with the provided serving endpoint in either
// Prometheus or OpenMetrics exposition format.
func (a *ServingEndpointsAPI) ExportMetricsByName(ctx context.Context, name string) error {
	return a.impl.ExportMetrics(ctx, ExportMetricsRequest{
		Name: name,
	})
}

// Get a single serving endpoint.
//
// Retrieves the details for a single serving endpoint.
func (a *ServingEndpointsAPI) Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error) {
	return a.impl.Get(ctx, request)
}

// Get a single serving endpoint.
//
// Retrieves the details for a single serving endpoint.
func (a *ServingEndpointsAPI) GetByName(ctx context.Context, name string) (*ServingEndpointDetailed, error) {
	return a.impl.Get(ctx, GetServingEndpointRequest{
		Name: name,
	})
}

// Retrieve all serving endpoints.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServingEndpointsAPI) ListAll(ctx context.Context) ([]ServingEndpoint, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Endpoints, nil
}

// Retrieve the most recent log lines associated with a given serving endpoint's
// served model.
//
// Retrieves the service logs associated with the provided served model.
func (a *ServingEndpointsAPI) Logs(ctx context.Context, request LogsRequest) (*ServerLogsResponse, error) {
	return a.impl.Logs(ctx, request)
}

// Retrieve the most recent log lines associated with a given serving endpoint's
// served model.
//
// Retrieves the service logs associated with the provided served model.
func (a *ServingEndpointsAPI) LogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*ServerLogsResponse, error) {
	return a.impl.Logs(ctx, LogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Query a serving endpoint with provided model input.
func (a *ServingEndpointsAPI) Query(ctx context.Context, request QueryRequest) (*QueryEndpointResponse, error) {
	return a.impl.Query(ctx, request)
}

// Update a serving endpoint with a new config.
//
// Updates any combination of the serving endpoint's served models, the compute
// configuration of those served models, and the endpoint's traffic config. An
// endpoint that already has an update in progress can not be updated until the
// current update completes or fails.
func (a *ServingEndpointsAPI) UpdateConfig(ctx context.Context, endpointCoreConfigInput EndpointCoreConfigInput) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error) {
	servingEndpointDetailed, err := a.impl.UpdateConfig(ctx, endpointCoreConfigInput)
	if err != nil {
		return nil, err
	}
	return &WaitGetServingEndpointNotUpdating[ServingEndpointDetailed]{
		Response: servingEndpointDetailed,
		Name:     servingEndpointDetailed.Name,
		poll: func(timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error) {
			return a.WaitGetServingEndpointNotUpdating(ctx, servingEndpointDetailed.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ServingEndpointsAPI.UpdateConfig] and waits to reach NOT_UPDATING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
//
// Deprecated: use [ServingEndpointsAPI.UpdateConfig].Get() or [ServingEndpointsAPI.WaitGetServingEndpointNotUpdating]
func (a *ServingEndpointsAPI) UpdateConfigAndWait(ctx context.Context, endpointCoreConfigInput EndpointCoreConfigInput, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error) {
	wait, err := a.UpdateConfig(ctx, endpointCoreConfigInput)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ServingEndpointDetailed]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ServingEndpointDetailed) {
		for _, o := range options {
			o(&retries.Info[ServingEndpointDetailed]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}
