// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaceconf

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewWorkspaceConf(client *client.DatabricksClient) WorkspaceConfService {
	return &WorkspaceConfAPI{
		client: client,
	}
}

type WorkspaceConfAPI struct {
	client *client.DatabricksClient
}

func (a *WorkspaceConfAPI) GetStatus(ctx context.Context, request GetStatusRequest) (*WorkspaceConf, error) {
	var workspaceConf WorkspaceConf
	path := "/api/2.0/workspace-conf"
	err := a.client.Get(ctx, path, request, &workspaceConf)
	return &workspaceConf, err
}

func (a *WorkspaceConfAPI) SetStatus(ctx context.Context, request WorkspaceConf) error {
	path := "/api/2.0/workspace-conf"
	err := a.client.Patch(ctx, path, request)
	return err
}
