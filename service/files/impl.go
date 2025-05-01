// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Dbfs API methods
type dbfsImpl struct {
	client *client.DatabricksClient
}

func (a *dbfsImpl) AddBlock(ctx context.Context, request AddBlock) error {

	requestPb, pbErr := addBlockToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var addBlockResponsePb addBlockResponsePb
	path := "/api/2.0/dbfs/add-block"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &addBlockResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *dbfsImpl) Close(ctx context.Context, request Close) error {

	requestPb, pbErr := closeToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var closeResponsePb closeResponsePb
	path := "/api/2.0/dbfs/close"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &closeResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *dbfsImpl) Create(ctx context.Context, request Create) (*CreateResponse, error) {

	requestPb, pbErr := createToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createResponsePb createResponsePb
	path := "/api/2.0/dbfs/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := createResponseFromPb(&createResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dbfsImpl) Delete(ctx context.Context, request Delete) error {

	requestPb, pbErr := deleteToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := "/api/2.0/dbfs/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &deleteResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *dbfsImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error) {

	requestPb, pbErr := getStatusRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var fileInfoPb fileInfoPb
	path := "/api/2.0/dbfs/get-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &fileInfoPb)
	if err != nil {
		return nil, err
	}
	resp, err := fileInfoFromPb(&fileInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	requestPb, pbErr := listDbfsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listStatusResponsePb listStatusResponsePb
	path := "/api/2.0/dbfs/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listStatusResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listStatusResponseFromPb(&listStatusResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dbfsImpl) Mkdirs(ctx context.Context, request MkDirs) error {

	requestPb, pbErr := mkDirsToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var mkDirsResponsePb mkDirsResponsePb
	path := "/api/2.0/dbfs/mkdirs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &mkDirsResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *dbfsImpl) Move(ctx context.Context, request Move) error {

	requestPb, pbErr := moveToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var moveResponsePb moveResponsePb
	path := "/api/2.0/dbfs/move"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &moveResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *dbfsImpl) Put(ctx context.Context, request Put) error {

	requestPb, pbErr := putToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var putResponsePb putResponsePb
	path := "/api/2.0/dbfs/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &putResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *dbfsImpl) Read(ctx context.Context, request ReadDbfsRequest) (*ReadResponse, error) {

	requestPb, pbErr := readDbfsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var readResponsePb readResponsePb
	path := "/api/2.0/dbfs/read"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &readResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := readResponseFromPb(&readResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Files API methods
type filesImpl struct {
	client *client.DatabricksClient
}

func (a *filesImpl) CreateDirectory(ctx context.Context, request CreateDirectoryRequest) error {

	requestPb, pbErr := createDirectoryRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var createDirectoryResponsePb createDirectoryResponsePb
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(requestPb.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, nil, &createDirectoryResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *filesImpl) Delete(ctx context.Context, request DeleteFileRequest) error {

	requestPb, pbErr := deleteFileRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(requestPb.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, (*requestPb), &deleteResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *filesImpl) DeleteDirectory(ctx context.Context, request DeleteDirectoryRequest) error {

	requestPb, pbErr := deleteDirectoryRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteDirectoryResponsePb deleteDirectoryResponsePb
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(requestPb.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, (*requestPb), &deleteDirectoryResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *filesImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {

	requestPb, pbErr := downloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var downloadResponsePb downloadResponsePb
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(requestPb.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/octet-stream"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &downloadResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := downloadResponseFromPb(&downloadResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *filesImpl) GetDirectoryMetadata(ctx context.Context, request GetDirectoryMetadataRequest) error {

	requestPb, pbErr := getDirectoryMetadataRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var getDirectoryMetadataResponsePb getDirectoryMetadataResponsePb
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(requestPb.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodHead, path, headers, queryParams, (*requestPb), &getDirectoryMetadataResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *filesImpl) GetMetadata(ctx context.Context, request GetMetadataRequest) (*GetMetadataResponse, error) {

	requestPb, pbErr := getMetadataRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getMetadataResponsePb getMetadataResponsePb
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(requestPb.FilePath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodHead, path, headers, queryParams, (*requestPb), &getMetadataResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getMetadataResponseFromPb(&getMetadataResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	requestPb, pbErr := listDirectoryContentsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listDirectoryResponsePb listDirectoryResponsePb
	path := fmt.Sprintf("/api/2.0/fs/directories%v", httpclient.EncodeMultiSegmentPathParameter(requestPb.DirectoryPath))
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listDirectoryResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listDirectoryResponseFromPb(&listDirectoryResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *filesImpl) Upload(ctx context.Context, request UploadRequest) error {

	requestPb, pbErr := uploadRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var uploadResponsePb uploadResponsePb
	path := fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(requestPb.FilePath))
	queryParams := make(map[string]any)
	if requestPb.Overwrite != false || slices.Contains(requestPb.ForceSendFields, "Overwrite") {
		queryParams["overwrite"] = requestPb.Overwrite
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, (*requestPb).Contents, &uploadResponsePb)
	if err != nil {
		return err
	}

	return err
}
