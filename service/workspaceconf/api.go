// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaceconf

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

func NewWorkspaceConf(client *client.DatabricksClient) *WorkspaceConfAPI {
	return &WorkspaceConfAPI{
		impl: &workspaceConfImpl{
			client: client,
		},
	}
}

// This API allows updating known workspace settings for advanced users.
type WorkspaceConfAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WorkspaceConfService)
	impl WorkspaceConfService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *WorkspaceConfAPI) WithImpl(impl WorkspaceConfService) *WorkspaceConfAPI {
	a.impl = impl
	return a
}

// Impl returns low-level WorkspaceConf API implementation
func (a *WorkspaceConfAPI) Impl() WorkspaceConfService {
	return a.impl
}

// Check configuration status
//
// Gets the configuration status for a workspace.
func (a *WorkspaceConfAPI) GetStatus(ctx context.Context, request GetStatus) (*WorkspaceConf, error) {
	return a.impl.GetStatus(ctx, request)
}

// Enable/disable features
//
// Sets the configuration status for a workspace, including enabling or
// disabling it.
func (a *WorkspaceConfAPI) SetStatus(ctx context.Context, request WorkspaceConf) error {
	return a.impl.SetStatus(ctx, request)
}
