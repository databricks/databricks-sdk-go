// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs provide specific management operations for Lakeview dashboards.
package dashboards

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type LakeviewInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockLakeviewInterface instead.
	WithImpl(impl LakeviewService) LakeviewInterface

	// Impl returns low-level Lakeview API implementation
	// Deprecated: use MockLakeviewInterface instead.
	Impl() LakeviewService

	// Publish dashboard.
	//
	// Publish the current draft dashboard.
	Publish(ctx context.Context, request PublishRequest) error
}

func NewLakeview(client *client.DatabricksClient) *LakeviewAPI {
	return &LakeviewAPI{
		impl: &lakeviewImpl{
			client: client,
		},
	}
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
type LakeviewAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(LakeviewService)
	impl LakeviewService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockLakeviewInterface instead.
func (a *LakeviewAPI) WithImpl(impl LakeviewService) LakeviewInterface {
	a.impl = impl
	return a
}

// Impl returns low-level Lakeview API implementation
// Deprecated: use MockLakeviewInterface instead.
func (a *LakeviewAPI) Impl() LakeviewService {
	return a.impl
}

// Publish dashboard.
//
// Publish the current draft dashboard.
func (a *LakeviewAPI) Publish(ctx context.Context, request PublishRequest) error {
	return a.impl.Publish(ctx, request)
}
