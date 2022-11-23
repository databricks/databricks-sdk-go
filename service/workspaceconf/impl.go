// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaceconf

import (
	"context"
	"net/http"
)

type databricksClient interface {
	Do(ctx context.Context, method string, path string, request any, response any) error
	ConfiguredAccountID() string
	IsAws() bool
}

// unexported type that holds implementations of just WorkspaceConf API methods
type workspaceConfImpl struct {
	client databricksClient
}

func (a *workspaceConfImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*WorkspaceConf, error) {
	var workspaceConf WorkspaceConf
	path := "/api/2.0/workspace-conf"
	err := a.client.Do(ctx, http.MethodGet, path, request, &workspaceConf)
	return &workspaceConf, err
}

func (a *workspaceConfImpl) SetStatus(ctx context.Context, request WorkspaceConf) error {
	path := "/api/2.0/workspace-conf"
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}
