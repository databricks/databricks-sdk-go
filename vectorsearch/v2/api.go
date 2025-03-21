// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Vector Search Endpoints, Vector Search Indexes, etc.
package vectorsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type vectorSearchEndpointsBaseClient struct {
	vectorSearchEndpointsImpl
}

// Create an endpoint.
//
// Create a new endpoint.
func (a *vectorSearchEndpointsBaseClient) CreateEndpoint(ctx context.Context, createEndpoint CreateEndpoint) (*VectorSearchEndpointsCreateEndpointWaiter, error) {
	endpointInfo, err := a.vectorSearchEndpointsImpl.CreateEndpoint(ctx, createEndpoint)
	if err != nil {
		return nil, err
	}
	return &VectorSearchEndpointsCreateEndpointWaiter{
		RawResponse:  endpointInfo,
		endpointName: endpointInfo.Name,
		service:      a,
	}, nil
}

type VectorSearchEndpointsCreateEndpointWaiter struct {
	// RawResponse is the raw response of the CreateEndpoint call.
	RawResponse  *EndpointInfo
	service      *vectorSearchEndpointsBaseClient
	endpointName string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *VectorSearchEndpointsCreateEndpointWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*EndpointInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[EndpointInfo](ctx, opts.Timeout, func() (*EndpointInfo, *retries.Err) {
		endpointInfo, err := w.service.GetEndpoint(ctx, GetEndpointRequest{
			EndpointName: w.endpointName,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := endpointInfo.EndpointStatus.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if endpointInfo.EndpointStatus != nil {
			statusMessage = endpointInfo.EndpointStatus.Message
		}
		switch status {
		case EndpointStatusStateOnline: // target state
			return endpointInfo, nil
		case EndpointStatusStateOffline:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				EndpointStatusStateOnline, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})

}

// Delete an endpoint.
func (a *vectorSearchEndpointsBaseClient) DeleteEndpointByEndpointName(ctx context.Context, endpointName string) (*DeleteEndpointResponse, error) {
	return a.vectorSearchEndpointsImpl.DeleteEndpoint(ctx, DeleteEndpointRequest{
		EndpointName: endpointName,
	})
}

// Get an endpoint.
func (a *vectorSearchEndpointsBaseClient) GetEndpointByEndpointName(ctx context.Context, endpointName string) (*EndpointInfo, error) {
	return a.vectorSearchEndpointsImpl.GetEndpoint(ctx, GetEndpointRequest{
		EndpointName: endpointName,
	})
}

type vectorSearchIndexesBaseClient struct {
	vectorSearchIndexesImpl
}

// Delete an index.
//
// Delete an index.
func (a *vectorSearchIndexesBaseClient) DeleteIndexByIndexName(ctx context.Context, indexName string) (*DeleteIndexResponse, error) {
	return a.vectorSearchIndexesImpl.DeleteIndex(ctx, DeleteIndexRequest{
		IndexName: indexName,
	})
}

// Get an index.
//
// Get an index.
func (a *vectorSearchIndexesBaseClient) GetIndexByIndexName(ctx context.Context, indexName string) (*VectorIndex, error) {
	return a.vectorSearchIndexesImpl.GetIndex(ctx, GetIndexRequest{
		IndexName: indexName,
	})
}
