// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dbfs

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewDbfs(client *client.DatabricksClient) *DbfsAPI {
	return &DbfsAPI{
		impl: &dbfsImpl{
			client: client,
		},
	}
}

// DBFS API makes it simple to interact with various data sources without having
// to include a users credentials every time to read a file.
type DbfsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(DbfsService)
	impl DbfsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *DbfsAPI) WithImpl(impl DbfsService) *DbfsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Dbfs API implementation
func (a *DbfsAPI) Impl() DbfsService {
	return a.impl
}

// Append data block
//
// Appends a block of data to the stream specified by the input handle. If the
// handle does not exist, this call will throw an exception with
// “RESOURCE_DOES_NOT_EXIST“.
//
// If the block of data exceeds 1 MB, this call will throw an exception with
// “MAX_BLOCK_SIZE_EXCEEDED“.
//
// Example of request: ``` { "data": "ZGF0YWJyaWNrcwo=", "handle": 7904256 } ```
func (a *DbfsAPI) AddBlock(ctx context.Context, request AddBlock) error {
	return a.impl.AddBlock(ctx, request)
}

// Close the stream
//
// Closes the stream specified by the input handle. If the handle does not
// exist, this call throws an exception with “RESOURCE_DOES_NOT_EXIST“.
func (a *DbfsAPI) Close(ctx context.Context, request Close) error {
	return a.impl.Close(ctx, request)
}

// Close the stream
//
// Closes the stream specified by the input handle. If the handle does not
// exist, this call throws an exception with “RESOURCE_DOES_NOT_EXIST“.
func (a *DbfsAPI) CloseByHandle(ctx context.Context, handle int64) error {
	return a.impl.Close(ctx, Close{
		Handle: handle,
	})
}

// Open a stream
//
// "Opens a stream to write to a file and returns a handle to this stream. There
// is a 10 minute idle timeout on this handle. If a file or directory already
// exists on the given path and __overwrite__ is set to `false`, this call
// throws an exception with “RESOURCE_ALREADY_EXISTS“.
//
// A typical workflow for file upload would be:
//
// 1. Issue a `create` call and get a handle. 2. Issue one or more `add-block`
// calls with the handle you have. 3. Issue a `close` call with the handle you
// have.
func (a *DbfsAPI) Create(ctx context.Context, request Create) (*CreateResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a file/directory
//
// "Deletes the file or directory (optionally, recursively delete all files in
// the directory).
//
// This all throws an exception with “IO_ERROR“ if the path is a non-empty
// directory and recursive is set to `false` or other similar errors.",
func (a *DbfsAPI) Delete(ctx context.Context, request Delete) error {
	return a.impl.Delete(ctx, request)
}

// Get the information of a file or directory
//
// Gets the file information for a file or directory. If the file or directory
// does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“.
func (a *DbfsAPI) GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error) {
	return a.impl.GetStatus(ctx, request)
}

// Get the information of a file or directory
//
// Gets the file information for a file or directory. If the file or directory
// does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“.
func (a *DbfsAPI) GetStatusByPath(ctx context.Context, path string) (*FileInfo, error) {
	return a.impl.GetStatus(ctx, GetStatusRequest{
		Path: path,
	})
}

// List directory contents or file details
//
// Lists the contents of a directory, or details of a file. If the file or
// directory does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“.
//
// Example of reply:
//
// ``` { "files": [ { "path": "/a.cpp", "is_dir": false, "file_size\": 261 }, {
// "path": "/databricks-results", "is_dir": true, "file_size\": 0 } ] } ```
//
// This method is generated by Databricks SDK Code Generator.
func (a *DbfsAPI) ListAll(ctx context.Context, request ListRequest) ([]FileInfo, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Files, nil
}

// List directory contents or file details
//
// Lists the contents of a directory, or details of a file. If the file or
// directory does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“.
//
// Example of reply:
//
// ``` { "files": [ { "path": "/a.cpp", "is_dir": false, "file_size\": 261 }, {
// "path": "/databricks-results", "is_dir": true, "file_size\": 0 } ] } ```
func (a *DbfsAPI) ListByPath(ctx context.Context, path string) (*ListStatusResponse, error) {
	return a.impl.List(ctx, ListRequest{
		Path: path,
	})
}

// Create a directory
//
// Creates the given directory and necessary parent directories if they do not
// exist. If a file (not a directory) exists at any prefix of the input path,
// this call throws an exception with “RESOURCE_ALREADY_EXISTS“. **Note**: If
// this operation fails, it might have succeeded in creating some of the
// necessary parent directories.",
func (a *DbfsAPI) Mkdirs(ctx context.Context, request MkDirs) error {
	return a.impl.Mkdirs(ctx, request)
}

// Create a directory
//
// Creates the given directory and necessary parent directories if they do not
// exist. If a file (not a directory) exists at any prefix of the input path,
// this call throws an exception with “RESOURCE_ALREADY_EXISTS“. **Note**: If
// this operation fails, it might have succeeded in creating some of the
// necessary parent directories.",
func (a *DbfsAPI) MkdirsByPath(ctx context.Context, path string) error {
	return a.impl.Mkdirs(ctx, MkDirs{
		Path: path,
	})
}

// Move a file
//
// Moves a file from one location to another location within DBFS. If the source
// file does not exist, this call throws an exception with
// “RESOURCE_DOES_NOT_EXIST“. If a file already exists in the destination
// path, this call throws an exception with “RESOURCE_ALREADY_EXISTS“. If the
// given source path is a directory, this call always recursively moves all
// files.",
func (a *DbfsAPI) Move(ctx context.Context, request Move) error {
	return a.impl.Move(ctx, request)
}

// Upload a file
//
// Uploads a file through the use of multipart form post. It is mainly used for
// streaming uploads, but can also be used as a convenient single call for data
// upload.
//
// Example usage:
//
// ``` curl -u USER:PASS -F contents=@localsrc -F
// path="https://XX.cloud.databricks.com/api/2.0/dbfs/put" ```
//
// Please note that “localsrc“ is the path to a local file to upload and this
// usage is only supported with multipart form post (such as using -F or --form
// with curl).
//
// Alternatively you can pass contents as base64 string.
//
// The amount of data that can be passed (when not streaming) using the
// __contents__ parameter is limited to 1 MB. “MAX_BLOCK_SIZE_EXCEEDED“ will
// be thrown if this limit is exceeded.
//
// If you want to upload large files, use the streaming upload. For details, see
// :method:create, :method:addBlock, :method:close.
func (a *DbfsAPI) Put(ctx context.Context, request Put) error {
	return a.impl.Put(ctx, request)
}

// Get the contents of a file
//
// "Returns the contents of a file. If the file does not exist, this call throws
// an exception with “RESOURCE_DOES_NOT_EXIST“. If the path is a directory,
// the read length is negative, or if the offset is negative, this call throws
// an exception with “INVALID_PARAMETER_VALUE“. If the read length exceeds 1
// MB, this call throws an\nexception with “MAX_READ_SIZE_EXCEEDED“.
//
// If “offset + length“ exceeds the number of bytes in a file, it reads the
// contents until the end of file.",
func (a *DbfsAPI) Read(ctx context.Context, request ReadRequest) (*ReadResponse, error) {
	return a.impl.Read(ctx, request)
}
