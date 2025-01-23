// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/httpclient"
)

// unexported type that holds implementations of just Dbfs API methods
type dbfsImpl struct {
	client *client.DatabricksClient
}

func (a *dbfsImpl) AddBlock(ctx context.Context, request AddBlock) error {
	var addBlockResponse AddBlockResponse
	path := "/api/2.0/dbfs/add-block"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &addBlockResponse)
	return err
}

func (a *dbfsImpl) Close(ctx context.Context, request Close) error {
	var closeResponse CloseResponse
	path := "/api/2.0/dbfs/close"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &closeResponse)
	return err
}

func (a *dbfsImpl) Create(ctx context.Context, request Create) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/dbfs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createResponse)
	return &createResponse, err
}

func (a *dbfsImpl) Delete(ctx context.Context, request Delete) error {
	var deleteResponse DeleteResponse
	path := "/api/2.0/dbfs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *dbfsImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error) {
	var fileInfo FileInfo
	path := "/api/2.0/dbfs/get-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &fileInfo)
	return &fileInfo, err
}

func (a *dbfsImpl) List(ctx context.Context, request ListDbfsRequest) (*ListStatusResponse, error) {
	var listStatusResponse ListStatusResponse
	path := "/api/2.0/dbfs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listStatusResponse)
	return &listStatusResponse, err
}

func (a *dbfsImpl) Mkdirs(ctx context.Context, request MkDirs) error {
	var mkDirsResponse MkDirsResponse
	path := "/api/2.0/dbfs/mkdirs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &mkDirsResponse)
	return err
}

func (a *dbfsImpl) Move(ctx context.Context, request Move) error {
	var moveResponse MoveResponse
	path := "/api/2.0/dbfs/move"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &moveResponse)
	return err
}

func (a *dbfsImpl) Put(ctx context.Context, request Put) error {
	var putResponse PutResponse
	path := "/api/2.0/dbfs/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &putResponse)
	return err
}

func (a *dbfsImpl) Read(ctx context.Context, request ReadDbfsRequest) (*ReadResponse, error) {
	var readResponse ReadResponse
	path := "/api/2.0/dbfs/read"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &readResponse)
	return &readResponse, err
}

// unexported type that holds implementations of just Files API methods
type filesImpl struct {
	client *client.DatabricksClient
}

func (a *filesImpl) CreateDirectory(ctx context.Context, request CreateDirectoryRequest) error {
	var createDirectoryResponse CreateDirectoryResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, nil, &createDirectoryResponse)
	return err
}

func (a *filesImpl) Delete(ctx context.Context, request DeleteFileRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *filesImpl) DeleteDirectory(ctx context.Context, request DeleteDirectoryRequest) error {
	var deleteDirectoryResponse DeleteDirectoryResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteDirectoryResponse)
	return err
}

func (a *filesImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	var downloadResponse DownloadResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/octet-stream"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &downloadResponse)
	return &downloadResponse, err
}

func (a *filesImpl) GetDirectoryMetadata(ctx context.Context, request GetDirectoryMetadataRequest) error {
	var getDirectoryMetadataResponse GetDirectoryMetadataResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodHead, path, headers, queryParams, request, &getDirectoryMetadataResponse)
	return err
}

func (a *filesImpl) GetMetadata(ctx context.Context, request GetMetadataRequest) (*GetMetadataResponse, error) {
	var getMetadataResponse GetMetadataResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodHead, path, headers, queryParams, request, &getMetadataResponse)
	return &getMetadataResponse, err
}

func (a *filesImpl) ListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) (*ListDirectoryResponse, error) {
	var listDirectoryResponse ListDirectoryResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDirectoryResponse)
	return &listDirectoryResponse, err
}

func (a *filesImpl) Upload(ctx context.Context, request UploadRequest) error {
	var uploadResponse UploadResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	if request.Overwrite != false || slices.Contains(request.ForceSendFields, "Overwrite") {
		queryParams["overwrite"] = request.Overwrite
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request.Contents, &uploadResponse)
	return err
}
