// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Dbfs, Files, etc.
package files

import (
	"context"
)

type dbfsBaseClient struct {
	dbfsImpl
}

// Close the stream.
//
// Closes the stream specified by the input handle. If the handle does not
// exist, this call throws an exception with “RESOURCE_DOES_NOT_EXIST“.
func (a *dbfsBaseClient) CloseByHandle(ctx context.Context, handle int64) (*CloseResponse, error) {
	return a.dbfsImpl.Close(ctx, Close{
		Handle: handle,
	})
}

// Get the information of a file or directory.
//
// Gets the file information for a file or directory. If the file or directory
// does not exist, this call throws an exception with `RESOURCE_DOES_NOT_EXIST`.
func (a *dbfsBaseClient) GetStatusByPath(ctx context.Context, path string) (*FileInfo, error) {
	return a.dbfsImpl.GetStatus(ctx, GetStatusRequest{
		Path: path,
	})
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
func (a *dbfsBaseClient) ListByPath(ctx context.Context, path string) (*ListStatusResponse, error) {
	return a.dbfsImpl.internalList(ctx, ListDbfsRequest{
		Path: path,
	})
}

// Create a directory.
//
// Creates the given directory and necessary parent directories if they do not
// exist. If a file (not a directory) exists at any prefix of the input path,
// this call throws an exception with `RESOURCE_ALREADY_EXISTS`. **Note**: If
// this operation fails, it might have succeeded in creating some of the
// necessary parent directories.
func (a *dbfsBaseClient) MkdirsByPath(ctx context.Context, path string) (*MkDirsResponse, error) {
	return a.dbfsImpl.Mkdirs(ctx, MkDirs{
		Path: path,
	})
}

type filesBaseClient struct {
	filesImpl
}

// Delete a file.
//
// Deletes a file. If the request is successful, there is no response body.
func (a *filesBaseClient) DeleteByFilePath(ctx context.Context, filePath string) (*DeleteResponse, error) {
	return a.filesImpl.Delete(ctx, DeleteFileRequest{
		FilePath: filePath,
	})
}

// Delete a directory.
//
// Deletes an empty directory.
//
// To delete a non-empty directory, first delete all of its contents. This can
// be done by listing the directory contents and deleting each file and
// subdirectory recursively.
func (a *filesBaseClient) DeleteDirectoryByDirectoryPath(ctx context.Context, directoryPath string) (*DeleteDirectoryResponse, error) {
	return a.filesImpl.DeleteDirectory(ctx, DeleteDirectoryRequest{
		DirectoryPath: directoryPath,
	})
}

// Download a file.
//
// Downloads a file. The file contents are the response body. This is a standard
// HTTP file download, not a JSON RPC. It supports the Range and
// If-Unmodified-Since HTTP headers.
func (a *filesBaseClient) DownloadByFilePath(ctx context.Context, filePath string) (*DownloadResponse, error) {
	return a.filesImpl.Download(ctx, DownloadRequest{
		FilePath: filePath,
	})
}

// Get directory metadata.
//
// Get the metadata of a directory. The response HTTP headers contain the
// metadata. There is no response body.
//
// This method is useful to check if a directory exists and the caller has
// access to it.
//
// If you wish to ensure the directory exists, you can instead use `PUT`, which
// will create the directory if it does not exist, and is idempotent (it will
// succeed if the directory already exists).
func (a *filesBaseClient) GetDirectoryMetadataByDirectoryPath(ctx context.Context, directoryPath string) (*GetDirectoryMetadataResponse, error) {
	return a.filesImpl.GetDirectoryMetadata(ctx, GetDirectoryMetadataRequest{
		DirectoryPath: directoryPath,
	})
}

// Get file metadata.
//
// Get the metadata of a file. The response HTTP headers contain the metadata.
// There is no response body.
func (a *filesBaseClient) GetMetadataByFilePath(ctx context.Context, filePath string) (*GetMetadataResponse, error) {
	return a.filesImpl.GetMetadata(ctx, GetMetadataRequest{
		FilePath: filePath,
	})
}

// List directory contents.
//
// Returns the contents of a directory. If there is no directory at the
// specified path, the API returns a HTTP 404 error.
func (a *filesBaseClient) ListDirectoryContentsByDirectoryPath(ctx context.Context, directoryPath string) (*ListDirectoryResponse, error) {
	return a.filesImpl.internalListDirectoryContents(ctx, ListDirectoryContentsRequest{
		DirectoryPath: directoryPath,
	})
}
