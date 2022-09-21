// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dbfs

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewDbfs(client *client.DatabricksClient) DbfsService {
	return &DbfsAPI{
		client: client,
	}
}

type DbfsAPI struct {
	client *client.DatabricksClient
}

// Appends a block of data to the stream specified by the input handle. If the
// handle does not exist, this call will throw an exception with
// “RESOURCE_DOES_NOT_EXIST“. If the block of data exceeds 1 MB, this call
// will throw an exception with “MAX_BLOCK_SIZE_EXCEEDED“. Example of request:
// .. code:: { "data": "ZGF0YWJyaWNrcwo=", "handle": 7904256 }
func (a *DbfsAPI) AddBlock(ctx context.Context, request AddBlock) error {
	path := "/api/2.0/dbfs/add-block"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Closes the stream specified by the input handle. If the handle does not
// exist, this call throws an exception with “RESOURCE_DOES_NOT_EXIST“.
func (a *DbfsAPI) Close(ctx context.Context, request Close) error {
	path := "/api/2.0/dbfs/close"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Closes the stream specified by the input handle. If the handle does not
// exist, this call throws an exception with “RESOURCE_DOES_NOT_EXIST“.
func (a *DbfsAPI) CloseByHandle(ctx context.Context, handle int64) error {
	return a.Close(ctx, Close{
		Handle: handle,
	})
}

// Opens a stream to write to a file and returns a handle to this stream. There
// is a 10 minute idle timeout on this handle. If a file or directory already
// exists on the given path and overwrite is set to false, this call throws an
// exception with “RESOURCE_ALREADY_EXISTS“. A typical workflow for file
// upload would be: 1) Issue a “create“ call and get a handle. 2) Issue one or
// more “add-block“ calls with the handle you have. 3) Issue a “close“ call
// with the handle you have.
func (a *DbfsAPI) Create(ctx context.Context, request Create) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/dbfs/create"
	err := a.client.Post(ctx, path, request, &createResponse)
	return &createResponse, err
}

// Deletes the file or directory (optionally, recursively delete all files in
// the directory). This call throws an exception with “IO_ERROR“ if the path
// is a non-empty directory and recursive is set to false or other similar
// errors.
func (a *DbfsAPI) Delete(ctx context.Context, request Delete) error {
	path := "/api/2.0/dbfs/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Gets the file information of a file or directory. If the file or directory
// does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“.
func (a *DbfsAPI) GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error) {
	var fileInfo FileInfo
	path := "/api/2.0/dbfs/get-status"
	err := a.client.Get(ctx, path, request, &fileInfo)
	return &fileInfo, err
}

// Gets the file information of a file or directory. If the file or directory
// does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“.
func (a *DbfsAPI) GetStatusByPath(ctx context.Context, path string) (*FileInfo, error) {
	return a.GetStatus(ctx, GetStatusRequest{
		Path: path,
	})
}

// Lists the contents of a directory, or details of the file. If the file or
// directory does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“. Example of reply: .. code:: { "files": [ {
// "path": "/a.cpp", "is_dir": false, "file_size": 261 }, { "path":
// "/databricks-results", "is_dir": true, "file_size": 0 } ] }
func (a *DbfsAPI) List(ctx context.Context, request ListRequest) (*ListStatusResponse, error) {
	var listStatusResponse ListStatusResponse
	path := "/api/2.0/dbfs/list"
	err := a.client.Get(ctx, path, request, &listStatusResponse)
	return &listStatusResponse, err
}

// Lists the contents of a directory, or details of the file. If the file or
// directory does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“. Example of reply: .. code:: { "files": [ {
// "path": "/a.cpp", "is_dir": false, "file_size": 261 }, { "path":
// "/databricks-results", "is_dir": true, "file_size": 0 } ] }
func (a *DbfsAPI) ListByPath(ctx context.Context, path string) (*ListStatusResponse, error) {
	return a.List(ctx, ListRequest{
		Path: path,
	})
}

// Creates the given directory and necessary parent directories if they do not
// exist. If a file exists (not a directory) at any prefix of the input path,
// this call throws an exception with “RESOURCE_ALREADY_EXISTS“. **Note**: If
// this operation fails, it might have succeeded in creating some of the
// necessary parent directories.
func (a *DbfsAPI) Mkdirs(ctx context.Context, request MkDirs) error {
	path := "/api/2.0/dbfs/mkdirs"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Creates the given directory and necessary parent directories if they do not
// exist. If a file exists (not a directory) at any prefix of the input path,
// this call throws an exception with “RESOURCE_ALREADY_EXISTS“. **Note**: If
// this operation fails, it might have succeeded in creating some of the
// necessary parent directories.
func (a *DbfsAPI) MkdirsByPath(ctx context.Context, path string) error {
	return a.Mkdirs(ctx, MkDirs{
		Path: path,
	})
}

// Moves a file from one location to another location within DBFS. If the source
// file does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“. If a file already exists in the destination
// path, this call throws an exception with “RESOURCE_ALREADY_EXISTS“. If the
// given source path is a directory, this call always recursively moves all
// files.
func (a *DbfsAPI) Move(ctx context.Context, request Move) error {
	path := "/api/2.0/dbfs/move"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Uploads a file through the use of multipart form post. It is mainly used for
// streaming uploads, but can also be used as a convenient single call for data
// upload. Example usage: .. code:: curl -u USER:PASS -F contents=@localsrc -F
// path="PATH" https://XX.cloud.databricks.com/api/2.0/dbfs/put Please note that
// “localsrc“ is the path to a local file to upload and this usage is only
// supported with multipart form post (i.e. using -F or --form with curl).
// Alternatively you can pass contents as base64 string. Examples: .. code::
// curl -u USER:PASS -F contents="BASE64" -F path="PATH"
// https://XX.cloud.databricks.com/api/2.0/dbfs/put .. code:: curl -u USER:PASS
// -H "Content-Type: application/json" -d '{"path":"PATH","contents":"BASE64"}'
// https://XX.cloud.databricks.com/api/2.0/dbfs/put“ Amount of data that can be
// passed using contents (i.e. not streaming) parameter is limited to 1 MB,
// “MAX_BLOCK_SIZE_EXCEEDED“ will be thrown if exceeded. If you want to upload
// large files, use streaming upload. For details, see
// :ref:`dbfsDbfsServicecreate`, :ref:`dbfsDbfsServiceaddBlock` and
// :ref:`dbfsDbfsServiceclose`.
func (a *DbfsAPI) Put(ctx context.Context, request Put) error {
	path := "/api/2.0/dbfs/put"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Returns the contents of a file. If the file does not exist, this call throws
// an exception with “RESOURCE_DOES_NOT_EXIST“. If the path is a directory,
// the read length is negative, or if the offset is negative, this call throws
// an exception with “INVALID_PARAMETER_VALUE“. If the read length exceeds 1
// MB, this call throws an exception with “MAX_READ_SIZE_EXCEEDED“. If
// “offset + length“ exceeds the number of bytes in a file, it reads contents
// until the end of file.
func (a *DbfsAPI) Read(ctx context.Context, request ReadRequest) (*ReadResponse, error) {
	var readResponse ReadResponse
	path := "/api/2.0/dbfs/read"
	err := a.client.Get(ctx, path, request, &readResponse)
	return &readResponse, err
}
