// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package filespreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just DbfsPreview API methods
type dbfsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *dbfsPreviewImpl) AddBlock(ctx context.Context, request AddBlock) error {
	var addBlockResponse AddBlockResponse
	path := "/api/2.0preview/dbfs/add-block"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &addBlockResponse)
	return err
}

func (a *dbfsPreviewImpl) Close(ctx context.Context, request Close) error {
	var closeResponse CloseResponse
	path := "/api/2.0preview/dbfs/close"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &closeResponse)
	return err
}

func (a *dbfsPreviewImpl) Create(ctx context.Context, request Create) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0preview/dbfs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createResponse)
	return &createResponse, err
}

func (a *dbfsPreviewImpl) Delete(ctx context.Context, request Delete) error {
	var deleteResponse DeleteResponse
	path := "/api/2.0preview/dbfs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *dbfsPreviewImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error) {
	var fileInfo FileInfo
	path := "/api/2.0preview/dbfs/get-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &fileInfo)
	return &fileInfo, err
}

// List directory contents or file details.
//
// List the contents of a directory, or details of the file. If the file or
// directory does not exist, this call throws an exception with
// `RESOURCE_DOES_NOT_EXIST`.
//
// When calling list on a large directory, the list operation will time out
// after approximately 60 seconds. We strongly recommend using list only on
// directories containing less than 10K files and discourage using the DBFS REST
// API for operations that list more than 10K files. Instead, we recommend that
// you perform such operations in the context of a cluster, using the [File
// system utility (dbutils.fs)](/dev-tools/databricks-utils.html#dbutils-fs),
// which provides the same functionality without timing out.
func (a *dbfsPreviewImpl) List(ctx context.Context, request ListDbfsRequest) listing.Iterator[FileInfo] {

	getNextPage := func(ctx context.Context, req ListDbfsRequest) (*ListStatusResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListStatusResponse) []FileInfo {
		return resp.Files
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List directory contents or file details.
//
// List the contents of a directory, or details of the file. If the file or
// directory does not exist, this call throws an exception with
// `RESOURCE_DOES_NOT_EXIST`.
//
// When calling list on a large directory, the list operation will time out
// after approximately 60 seconds. We strongly recommend using list only on
// directories containing less than 10K files and discourage using the DBFS REST
// API for operations that list more than 10K files. Instead, we recommend that
// you perform such operations in the context of a cluster, using the [File
// system utility (dbutils.fs)](/dev-tools/databricks-utils.html#dbutils-fs),
// which provides the same functionality without timing out.
func (a *dbfsPreviewImpl) ListAll(ctx context.Context, request ListDbfsRequest) ([]FileInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FileInfo](ctx, iterator)
}
func (a *dbfsPreviewImpl) internalList(ctx context.Context, request ListDbfsRequest) (*ListStatusResponse, error) {
	var listStatusResponse ListStatusResponse
	path := "/api/2.0preview/dbfs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listStatusResponse)
	return &listStatusResponse, err
}

func (a *dbfsPreviewImpl) Mkdirs(ctx context.Context, request MkDirs) error {
	var mkDirsResponse MkDirsResponse
	path := "/api/2.0preview/dbfs/mkdirs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &mkDirsResponse)
	return err
}

func (a *dbfsPreviewImpl) Move(ctx context.Context, request Move) error {
	var moveResponse MoveResponse
	path := "/api/2.0preview/dbfs/move"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &moveResponse)
	return err
}

func (a *dbfsPreviewImpl) Put(ctx context.Context, request Put) error {
	var putResponse PutResponse
	path := "/api/2.0preview/dbfs/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &putResponse)
	return err
}

func (a *dbfsPreviewImpl) Read(ctx context.Context, request ReadDbfsRequest) (*ReadResponse, error) {
	var readResponse ReadResponse
	path := "/api/2.0preview/dbfs/read"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &readResponse)
	return &readResponse, err
}

// unexported type that holds implementations of just FilesPreview API methods
type filesPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *filesPreviewImpl) CreateDirectory(ctx context.Context, request CreateDirectoryRequest) error {
	var createDirectoryResponse CreateDirectoryResponse
	path := fmt.Sprintf("/api/2.0preview/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, nil, &createDirectoryResponse)
	return err
}

func (a *filesPreviewImpl) Delete(ctx context.Context, request DeleteFileRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *filesPreviewImpl) DeleteDirectory(ctx context.Context, request DeleteDirectoryRequest) error {
	var deleteDirectoryResponse DeleteDirectoryResponse
	path := fmt.Sprintf("/api/2.0preview/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteDirectoryResponse)
	return err
}

func (a *filesPreviewImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	var downloadResponse DownloadResponse
	path := fmt.Sprintf("/api/2.0preview/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/octet-stream"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &downloadResponse)
	return &downloadResponse, err
}

func (a *filesPreviewImpl) GetDirectoryMetadata(ctx context.Context, request GetDirectoryMetadataRequest) error {
	var getDirectoryMetadataResponse GetDirectoryMetadataResponse
	path := fmt.Sprintf("/api/2.0preview/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodHead, path, headers, queryParams, request, &getDirectoryMetadataResponse)
	return err
}

func (a *filesPreviewImpl) GetMetadata(ctx context.Context, request GetMetadataRequest) (*GetMetadataResponse, error) {
	var getMetadataResponse GetMetadataResponse
	path := fmt.Sprintf("/api/2.0preview/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodHead, path, headers, queryParams, request, &getMetadataResponse)
	return &getMetadataResponse, err
}

// List directory contents.
//
// Returns the contents of a directory. If there is no directory at the
// specified path, the API returns a HTTP 404 error.
func (a *filesPreviewImpl) ListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) listing.Iterator[DirectoryEntry] {

	getNextPage := func(ctx context.Context, req ListDirectoryContentsRequest) (*ListDirectoryResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDirectoryContents(ctx, req)
	}
	getItems := func(resp *ListDirectoryResponse) []DirectoryEntry {
		return resp.Contents
	}
	getNextReq := func(resp *ListDirectoryResponse) *ListDirectoryContentsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List directory contents.
//
// Returns the contents of a directory. If there is no directory at the
// specified path, the API returns a HTTP 404 error.
func (a *filesPreviewImpl) ListDirectoryContentsAll(ctx context.Context, request ListDirectoryContentsRequest) ([]DirectoryEntry, error) {
	iterator := a.ListDirectoryContents(ctx, request)
	return listing.ToSliceN[DirectoryEntry, int64](ctx, iterator, request.PageSize)

}
func (a *filesPreviewImpl) internalListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) (*ListDirectoryResponse, error) {
	var listDirectoryResponse ListDirectoryResponse
	path := fmt.Sprintf("/api/2.0preview/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDirectoryResponse)
	return &listDirectoryResponse, err
}

func (a *filesPreviewImpl) Upload(ctx context.Context, request UploadRequest) error {
	var uploadResponse UploadResponse
	path := fmt.Sprintf("/api/2.0preview/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	if request.Overwrite != false || slices.Contains(request.ForceSendFields, "Overwrite") {
		queryParams["overwrite"] = request.Overwrite
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request.Contents, &uploadResponse)
	return err
}
