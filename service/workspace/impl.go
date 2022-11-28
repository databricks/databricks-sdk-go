// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Workspace API methods
type workspaceImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceImpl) Delete(ctx context.Context, request Delete) error {
	path := "/api/2.0/workspace/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *workspaceImpl) Export(ctx context.Context, request Export) (*ExportResponse, error) {
	var exportResponse ExportResponse
	path := "/api/2.0/workspace/export"
	err := a.client.Do(ctx, http.MethodGet, path, request, &exportResponse)
	return &exportResponse, err
}

func (a *workspaceImpl) GetStatus(ctx context.Context, request GetStatus) (*ObjectInfo, error) {
	var objectInfo ObjectInfo
	path := "/api/2.0/workspace/get-status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &objectInfo)
	return &objectInfo, err
}

func (a *workspaceImpl) Import(ctx context.Context, request Import) error {
	path := "/api/2.0/workspace/import"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *workspaceImpl) List(ctx context.Context, request List) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/workspace/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listResponse)
	return &listResponse, err
}

func (a *workspaceImpl) Mkdirs(ctx context.Context, request Mkdirs) error {
	path := "/api/2.0/workspace/mkdirs"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}
