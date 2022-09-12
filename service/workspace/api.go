// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewWorkspace(client *client.DatabricksClient) WorkspaceService {
	return &WorkspaceAPI{
		client: client,
	}
}

type WorkspaceAPI struct {
	client *client.DatabricksClient
}

// Deletes an object or a directory (and optionally recursively deletes all
// objects in the directory). If “path“ does not exist, this call returns an
// error “RESOURCE_DOES_NOT_EXIST“. If “path“ is a non-empty directory and
// “recursive“ is set to “false“, this call returns an error
// “DIRECTORY_NOT_EMPTY“. Object deletion cannot be undone and deleting a
// directory recursively is not atomic. Example of request: .. code :: json {
// "path": "/Users/user@example.com/project", "recursive": true }
func (a *WorkspaceAPI) Delete(ctx context.Context, request DeleteRequest) error {
	path := "/api/2.0/workspace/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Exports a notebook or contents of an entire directory. If “path“ does not
// exist, this call returns an error “RESOURCE_DOES_NOT_EXIST“. One can only
// export a directory in “DBC“ format. If the exported data would exceed size
// limit, this call returns an error “MAX_NOTEBOOK_SIZE_EXCEEDED“. Currently,
// this API does not support exporting a library. Example of request: .. code ::
// json { "path": "/Users/user@example.com/project/ScalaExampleNotebook",
// "format": "SOURCE" } Example of response, where “content“ is
// base64-encoded: .. code :: json { "content":
// "Ly8gRGF0YWJyaWNrcyBub3RlYm9vayBzb3VyY2UKMSsx", } Alternaitvely, one can
// download the exported file by enabling “direct_download“: .. code :: shell
// curl -n -o example.scala \
// 'https://XX.cloud.databricks.com/api/2.0/workspace/export?path=/Users/user@example.com/ScalaExampleNotebook&direct_download=true'
func (a *WorkspaceAPI) Export(ctx context.Context, request ExportRequest) (*ExportResponse, error) {
	var exportResponse ExportResponse
	path := "/api/2.0/workspace/export"
	err := a.client.Get(ctx, path, request, &exportResponse)
	return &exportResponse, err
}

// Gets the status of an object or a directory. If “path“ does not exist, this
// call returns an error “RESOURCE_DOES_NOT_EXIST“. Example of request: ..
// code :: json { "path": "/Users/user@example.com/project/ScaleExampleNotebook"
// } Example of response: .. code :: json { "path":
// "/Users/user@example.com/project/ScalaExampleNotebook", "language": "SCALA",
// "object_type": "NOTEBOOK", "object_id": 789 }
func (a *WorkspaceAPI) GetStatus(ctx context.Context, request GetStatusRequest) (*GetStatusResponse, error) {
	var getStatusResponse GetStatusResponse
	path := "/api/2.0/workspace/get-status"
	err := a.client.Get(ctx, path, request, &getStatusResponse)
	return &getStatusResponse, err
}

// Gets the status of an object or a directory. If “path“ does not exist, this
// call returns an error “RESOURCE_DOES_NOT_EXIST“. Example of request: ..
// code :: json { "path": "/Users/user@example.com/project/ScaleExampleNotebook"
// } Example of response: .. code :: json { "path":
// "/Users/user@example.com/project/ScalaExampleNotebook", "language": "SCALA",
// "object_type": "NOTEBOOK", "object_id": 789 }
func (a *WorkspaceAPI) GetStatusByPath(ctx context.Context, path string) (*GetStatusResponse, error) {
	return a.GetStatus(ctx, GetStatusRequest{
		Path: path,
	})
}

// Imports a notebook or the contents of an entire directory. If “path“
// already exists and “overwrite“ is set to “false“, this call returns an
// error “RESOURCE_ALREADY_EXISTS“. One can only use “DBC“ format to import
// a directory. Example of request, where “content“ is the base64-encoded
// string of “1+1“: .. code :: json { "content": "MSsx\n", "path":
// "/Users/user@example.com/project/ScalaExampleNotebook", "language": "SCALA",
// "overwrite": true, "format": "SOURCE" } Alternatively, one can import a local
// file directly: .. code :: shell curl -n -F
// path=/Users/user@example.com/project/ScalaExampleNotebook -F language=SCALA \
// -F content=@example.scala \
// https://XX.cloud.databricks.com/api/2.0/workspace/import
func (a *WorkspaceAPI) Import(ctx context.Context, request ImportRequest) error {
	path := "/api/2.0/workspace/import"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Lists the contents of a directory, or the object if it is not a directory. If
// the input path does not exist, this call returns an error
// “RESOURCE_DOES_NOT_EXIST“. Example of request: .. code :: json { "path":
// "/Users/user@example.com/" } Example of response: .. code :: json {
// "objects": [ { "path": "/Users/user@example.com/project", "object_type":
// "DIRECTORY", "object_id": 123 }, { "path":
// "/Users/user@example.com/PythonExampleNotebook", "language": "PYTHON",
// "object_type": "NOTEBOOK", "object_id": 456 } ] }
func (a *WorkspaceAPI) List(ctx context.Context, request ListRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/workspace/list"
	err := a.client.Get(ctx, path, request, &listResponse)
	return &listResponse, err
}

// Creates the given directory and necessary parent directories if they do not
// exists. If there exists an object (not a directory) at any prefix of the
// input path, this call returns an error “RESOURCE_ALREADY_EXISTS“. Note that
// if this operation fails it may have succeeded in creating some of the
// necessary parrent directories. Example of request: .. code:: json { "path":
// "/Users/user@example.com/project" }
func (a *WorkspaceAPI) Mkdirs(ctx context.Context, request MkdirsRequest) error {
	path := "/api/2.0/workspace/mkdirs"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Creates the given directory and necessary parent directories if they do not
// exists. If there exists an object (not a directory) at any prefix of the
// input path, this call returns an error “RESOURCE_ALREADY_EXISTS“. Note that
// if this operation fails it may have succeeded in creating some of the
// necessary parrent directories. Example of request: .. code:: json { "path":
// "/Users/user@example.com/project" }
func (a *WorkspaceAPI) MkdirsByPath(ctx context.Context, path string) error {
	return a.Mkdirs(ctx, MkdirsRequest{
		Path: path,
	})
}
