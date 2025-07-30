// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Lorem Ipsum
package httpcallv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type HttpCallV2Interface interface {

	// This mimics "old" style post requests which have the resource inlined.
	CreateResource(ctx context.Context, request CreateResourceRequest) (*Resource, error)

	GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error)

	// This mimics "new" style post requests which have a body field.
	UpdateResource(ctx context.Context, request UpdateResourceRequest) (*Resource, error)
}

func NewHttpCallV2(client *client.DatabricksClient) *HttpCallV2API {
	return &HttpCallV2API{
		httpCallV2Impl: httpCallV2Impl{
			client: client,
		},
	}
}

// Lorem Ipsum
type HttpCallV2API struct {
	httpCallV2Impl
}
