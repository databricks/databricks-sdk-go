// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Dbfs API methods
type dbfsImpl struct {
	client *client.DatabricksClient
}

func (a *dbfsImpl) AddBlock(ctx context.Context, request AddBlock) error {
	path := "/api/2.0/dbfs/add-block"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *dbfsImpl) Close(ctx context.Context, request Close) error {
	path := "/api/2.0/dbfs/close"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *dbfsImpl) Create(ctx context.Context, request Create) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/dbfs/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createResponse)
	return &createResponse, err
}

func (a *dbfsImpl) Delete(ctx context.Context, request Delete) error {
	path := "/api/2.0/dbfs/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *dbfsImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error) {
	var fileInfo FileInfo
	path := "/api/2.0/dbfs/get-status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &fileInfo)
	return &fileInfo, err
}

func (a *dbfsImpl) List(ctx context.Context, request ListDbfsRequest) (*ListStatusResponse, error) {
	var listStatusResponse ListStatusResponse
	path := "/api/2.0/dbfs/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listStatusResponse)
	return &listStatusResponse, err
}

func (a *dbfsImpl) Mkdirs(ctx context.Context, request MkDirs) error {
	path := "/api/2.0/dbfs/mkdirs"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *dbfsImpl) Move(ctx context.Context, request Move) error {
	path := "/api/2.0/dbfs/move"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *dbfsImpl) Put(ctx context.Context, request Put) error {
	path := "/api/2.0/dbfs/put"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *dbfsImpl) Read(ctx context.Context, request ReadDbfsRequest) (*ReadResponse, error) {
	var readResponse ReadResponse
	path := "/api/2.0/dbfs/read"
	err := a.client.Do(ctx, http.MethodGet, path, request, &readResponse)
	return &readResponse, err
}
