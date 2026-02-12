// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package networking

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Endpoints API methods
type endpointsImpl struct {
	client *client.DatabricksClient
}

func (a *endpointsImpl) CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/networking/v1/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Endpoint, &endpoint)
	return &endpoint, err
}

func (a *endpointsImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error {
	path := fmt.Sprintf("/api/networking/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *endpointsImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/networking/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &endpoint)
	return &endpoint, err
}

// Lists all network connectivity endpoints for the account.
func (a *endpointsImpl) ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[Endpoint] {

	getNextPage := func(ctx context.Context, req ListEndpointsRequest) (*ListEndpointsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListEndpoints(ctx, req)
	}
	getItems := func(resp *ListEndpointsResponse) []Endpoint {
		return resp.Items
	}
	getNextReq := func(resp *ListEndpointsResponse) *ListEndpointsRequest {
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

// Lists all network connectivity endpoints for the account.
func (a *endpointsImpl) ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]Endpoint, error) {
	iterator := a.ListEndpoints(ctx, request)
	return listing.ToSlice[Endpoint](ctx, iterator)
}

func (a *endpointsImpl) internalListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error) {
	var listEndpointsResponse ListEndpointsResponse
	path := fmt.Sprintf("/api/networking/v1/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listEndpointsResponse)
	return &listEndpointsResponse, err
}
