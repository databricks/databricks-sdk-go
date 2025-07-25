// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Lorem Ipsum
package basicv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
)

type BasicV2Interface interface {
	CreateBasic(ctx context.Context, request CreateBasicRequest) (*Basic, error)

	DeleteBasic(ctx context.Context, request DeleteBasicRequest) error

	GetBasic(ctx context.Context, request GetBasicRequest) (*Basic, error)

	//
	// This method is generated by Databricks SDK Code Generator.
	ListBasics(ctx context.Context, request ListBasicsRequest) listing.Iterator[Basic]

	//
	// This method is generated by Databricks SDK Code Generator.
	ListBasicsAll(ctx context.Context, request ListBasicsRequest) ([]Basic, error)

	UpdateBasic(ctx context.Context, request UpdateBasicRequest) (*Basic, error)
}

func NewBasicV2(client *client.DatabricksClient) *BasicV2API {
	return &BasicV2API{
		basicV2Impl: basicV2Impl{
			client: client,
		},
	}
}

// Lorem Ipsum
type BasicV2API struct {
	basicV2Impl
}
