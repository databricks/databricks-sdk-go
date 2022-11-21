// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaceconf

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewWorkspaceConf(client *client.DatabricksClient) *WorkspaceConfAPI {
	return &WorkspaceConfAPI{
		WorkspaceConfService: &workspaceConfAPI{
			client: client,
		},
	}
}

// This API allows updating known workspace settings for advanced users.
type WorkspaceConfAPI struct {
	// WorkspaceConfService contains low-level REST API interface.
	WorkspaceConfService
}

// Check configuration status
//
// Gets the configuration status for a workspace.
func (a *WorkspaceConfAPI) GetStatus(ctx context.Context, request GetStatusRequest) (*WorkspaceConf, error) {
	return a.WorkspaceConfService.GetStatus(ctx, request)
}

// Enable/disable features
//
// Sets the configuration status for a workspace, including enabling or
// disabling it.
func (a *WorkspaceConfAPI) SetStatus(ctx context.Context, request WorkspaceConf) error {
	return a.WorkspaceConfService.SetStatus(ctx, request)
}

// unexported type that holds implementations of just WorkspaceConf API methods
type workspaceConfAPI struct {
	client *client.DatabricksClient
}

func (a *workspaceConfAPI) GetStatus(ctx context.Context, request GetStatusRequest) (*WorkspaceConf, error) {
	var workspaceConf WorkspaceConf
	path := "/api/2.0/workspace-conf"
	err := a.client.Get(ctx, path, request, &workspaceConf)
	return &workspaceConf, err
}

func (a *workspaceConfAPI) SetStatus(ctx context.Context, request WorkspaceConf) error {
	path := "/api/2.0/workspace-conf"
	err := a.client.Patch(ctx, path, request)
	return err
}
