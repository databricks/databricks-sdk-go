// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Dbfs API methods
type dbfsImpl struct {
	client *client.DatabricksClient
}

func (a *dbfsImpl) AddBlock(ctx context.Context, request AddBlock) (*AddBlockResponse, error) {
	var addBlockResponse AddBlockResponse
	path := "/api/2.0/dbfs/add-block"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &addBlockResponse)
	return &addBlockResponse, err
}

func (a *dbfsImpl) Close(ctx context.Context, request Close) (*CloseResponse, error) {
	var closeResponse CloseResponse
	path := "/api/2.0/dbfs/close"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &closeResponse)
	return &closeResponse, err
}

func (a *dbfsImpl) Create(ctx context.Context, request Create) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/dbfs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &createResponse)
	return &createResponse, err
}

func (a *dbfsImpl) Delete(ctx context.Context, request Delete) (*DeleteResponse, error) {
	var deleteResponse DeleteResponse
	path := "/api/2.0/dbfs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &deleteResponse)
	return &deleteResponse, err
}

func (a *dbfsImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error) {
	var fileInfo FileInfo
	path := "/api/2.0/dbfs/get-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &fileInfo)
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
func (a *dbfsImpl) List(ctx context.Context, request ListDbfsRequest) listing.Iterator[FileInfo] {

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
func (a *dbfsImpl) ListAll(ctx context.Context, request ListDbfsRequest) ([]FileInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FileInfo](ctx, iterator)
}
func (a *dbfsImpl) internalList(ctx context.Context, request ListDbfsRequest) (*ListStatusResponse, error) {
	var listStatusResponse ListStatusResponse
	path := "/api/2.0/dbfs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listStatusResponse)
	return &listStatusResponse, err
}

func (a *dbfsImpl) Mkdirs(ctx context.Context, request MkDirs) (*MkDirsResponse, error) {
	var mkDirsResponse MkDirsResponse
	path := "/api/2.0/dbfs/mkdirs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &mkDirsResponse)
	return &mkDirsResponse, err
}

func (a *dbfsImpl) Move(ctx context.Context, request Move) (*MoveResponse, error) {
	var moveResponse MoveResponse
	path := "/api/2.0/dbfs/move"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &moveResponse)
	return &moveResponse, err
}

func (a *dbfsImpl) Put(ctx context.Context, request Put) (*PutResponse, error) {
	var putResponse PutResponse
	path := "/api/2.0/dbfs/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &putResponse)
	return &putResponse, err
}

func (a *dbfsImpl) Read(ctx context.Context, request ReadDbfsRequest) (*ReadResponse, error) {
	var readResponse ReadResponse
	path := "/api/2.0/dbfs/read"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &readResponse)
	return &readResponse, err
}

// unexported type that holds implementations of just Files API methods
type filesImpl struct {
	client *client.DatabricksClient
}

func (a *filesImpl) CreateDirectory(ctx context.Context, request CreateDirectoryRequest) (*CreateDirectoryResponse, error) {
	var createDirectoryResponse CreateDirectoryResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := do(a.client, ctx, http.MethodPut, path, headers, queryParams, nil, &createDirectoryResponse)
	return &createDirectoryResponse, err
}

func (a *filesImpl) Delete(ctx context.Context, request DeleteFileRequest) (*DeleteResponse, error) {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := do(a.client, ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return &deleteResponse, err
}

func (a *filesImpl) DeleteDirectory(ctx context.Context, request DeleteDirectoryRequest) (*DeleteDirectoryResponse, error) {
	var deleteDirectoryResponse DeleteDirectoryResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := do(a.client, ctx, http.MethodDelete, path, headers, queryParams, request, &deleteDirectoryResponse)
	return &deleteDirectoryResponse, err
}

func (a *filesImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	var downloadResponse DownloadResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/octet-stream"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &downloadResponse)
	return &downloadResponse, err
}

func (a *filesImpl) GetDirectoryMetadata(ctx context.Context, request GetDirectoryMetadataRequest) (*GetDirectoryMetadataResponse, error) {
	var getDirectoryMetadataResponse GetDirectoryMetadataResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := do(a.client, ctx, http.MethodHead, path, headers, queryParams, request, &getDirectoryMetadataResponse)
	return &getDirectoryMetadataResponse, err
}

func (a *filesImpl) GetMetadata(ctx context.Context, request GetMetadataRequest) (*GetMetadataResponse, error) {
	var getMetadataResponse GetMetadataResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := do(a.client, ctx, http.MethodHead, path, headers, queryParams, request, &getMetadataResponse)
	return &getMetadataResponse, err
}

// List directory contents.
//
// Returns the contents of a directory. If there is no directory at the
// specified path, the API returns a HTTP 404 error.
func (a *filesImpl) ListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) listing.Iterator[DirectoryEntry] {

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
func (a *filesImpl) ListDirectoryContentsAll(ctx context.Context, request ListDirectoryContentsRequest) ([]DirectoryEntry, error) {
	iterator := a.ListDirectoryContents(ctx, request)
	return listing.ToSlice[DirectoryEntry](ctx, iterator)
}
func (a *filesImpl) internalListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) (*ListDirectoryResponse, error) {
	var listDirectoryResponse ListDirectoryResponse
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(request.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listDirectoryResponse)
	return &listDirectoryResponse, err
}

func (a *filesImpl) Upload(ctx context.Context, request UploadRequest) (*UploadResponse, error) {
	var uploadResponse UploadResponse
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath))
	queryParams := make(map[string]any)
	if request.Overwrite != false || slices.Contains(request.ForceSendFields, "Overwrite") {
		queryParams["overwrite"] = request.Overwrite
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	err := do(a.client, ctx, http.MethodPut, path, headers, queryParams, request.Contents, &uploadResponse)
	return &uploadResponse, err
}

func do(
	client *client.DatabricksClient,
	ctx context.Context,
	method string,
	path string,
	headers map[string]string,
	queryParams map[string]any,
	request any,
	response any,
	visitors ...func(*http.Request) error,
) error {
	opts := []httpclient.DoOption{}
	for _, v := range visitors {
		opts = append(opts, httpclient.WithRequestVisitor(v))
	}
	opts = append(opts, httpclient.WithQueryParameters(queryParams))
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	opts = append(opts, httpclient.WithRequestData(request))
	opts = append(opts, httpclient.WithResponseUnmarshal(response))

	// Remove extra `/` from path for files API. Once the OpenAPI spec doesn't
	// include the extra slash, we can remove this
	if strings.HasPrefix(path, "/api/2.0/fs/files//") {
		path = strings.Replace(path, "/api/2.0/fs/files//", "/api/2.0/fs/files/", 1)
	}

	return client.ApiClient().Do(ctx, method, path, opts...)
}
