// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Dbfs API methods
type dbfsImpl struct {
	client *client.DatabricksClient
}

func (a *dbfsImpl) AddBlock(ctx context.Context, request AddBlock) error {
	path := "/api/2.0/dbfs/add-block"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *dbfsImpl) Close(ctx context.Context, request Close) error {
	path := "/api/2.0/dbfs/close"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *dbfsImpl) Create(ctx context.Context, request Create) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/dbfs/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createResponse)
	return &createResponse, err
}

func (a *dbfsImpl) Delete(ctx context.Context, request Delete) error {
	path := "/api/2.0/dbfs/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *dbfsImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error) {
	var fileInfo FileInfo
	path := "/api/2.0/dbfs/get-status"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &fileInfo)
	return &fileInfo, err
}

func (a *dbfsImpl) List(ctx context.Context, request ListDbfsRequest) (*ListStatusResponse, error) {
	var listStatusResponse ListStatusResponse
	path := "/api/2.0/dbfs/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listStatusResponse)
	return &listStatusResponse, err
}

func (a *dbfsImpl) Mkdirs(ctx context.Context, request MkDirs) error {
	path := "/api/2.0/dbfs/mkdirs"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *dbfsImpl) Move(ctx context.Context, request Move) error {
	path := "/api/2.0/dbfs/move"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *dbfsImpl) Put(ctx context.Context, request Put) error {
	path := "/api/2.0/dbfs/put"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, nil)
	return err
}

func (a *dbfsImpl) Read(ctx context.Context, request ReadDbfsRequest) (*ReadResponse, error) {
	var readResponse ReadResponse
	path := "/api/2.0/dbfs/read"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &readResponse)
	return &readResponse, err
}

// unexported type that holds implementations of just Files API methods
type filesImpl struct {
	client *client.DatabricksClient
}

func (a *filesImpl) CreateDirectory(ctx context.Context, request CreateDirectoryRequest) error {
	path := fmt.Sprintf("/api/2.0/fs/directories%v", request.DirectoryPath)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPut, path, headers, nil, nil)
	return err
}

func (a *filesImpl) Delete(ctx context.Context, request DeleteFileRequest) error {
	path := fmt.Sprintf("/api/2.0/fs/files%v", request.FilePath)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *filesImpl) DeleteDirectory(ctx context.Context, request DeleteDirectoryRequest) error {
	path := fmt.Sprintf("/api/2.0/fs/directories%v", request.DirectoryPath)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *filesImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	var downloadResponse DownloadResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", request.FilePath)
	headers := make(map[string]string)
	headers["Accept"] = "application/octet-stream"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &downloadResponse)
	return &downloadResponse, err
}

func (a *filesImpl) GetDirectoryMetadata(ctx context.Context, request GetDirectoryMetadataRequest) error {
	path := fmt.Sprintf("/api/2.0/fs/directories%v", request.DirectoryPath)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodHead, path, headers, request, nil)
	return err
}

func (a *filesImpl) GetMetadata(ctx context.Context, request GetMetadataRequest) (*GetMetadataResponse, error) {
	var getMetadataResponse GetMetadataResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", request.FilePath)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodHead, path, headers, request, &getMetadataResponse)
	return &getMetadataResponse, err
}

func (a *filesImpl) ListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) (*ListDirectoryResponse, error) {
	var listDirectoryResponse ListDirectoryResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", request.DirectoryPath)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listDirectoryResponse)
	return &listDirectoryResponse, err
}

func (a *filesImpl) Upload(ctx context.Context, request UploadRequest) error {
	path := fmt.Sprintf("/api/2.0/fs/files%v", request.FilePath)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request.Contents, nil)
	return err
}
