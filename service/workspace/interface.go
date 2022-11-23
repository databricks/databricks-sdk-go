// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"
)

// The Workspace API allows you to list, import, export, and delete notebooks
// and folders.
//
// A notebook is a web-based interface to a document that contains runnable
// code, visualizations, and explanatory text.
type WorkspaceService interface {

	// Delete a workspace object
	//
	// Deletes an object or a directory (and optionally recursively deletes all
	// objects in the directory). * If ``path`` does not exist, this call
	// returns an error ``RESOURCE_DOES_NOT_EXIST``. * If ``path`` is a
	// non-empty directory and ``recursive`` is set to ``false``, this call
	// returns an error ``DIRECTORY_NOT_EMPTY``.
	//
	// Object deletion cannot be undone and deleting a directory recursively is
	// not atomic.
	Delete(ctx context.Context, request Delete) error

	// Export a notebook
	//
	// Exports a notebook or the contents of an entire directory.
	//
	// If ``path`` does not exist, this call returns an error
	// ``RESOURCE_DOES_NOT_EXIST``.
	//
	// One can only export a directory in ``DBC`` format. If the exported data
	// would exceed size limit, this call returns
	// ``MAX_NOTEBOOK_SIZE_EXCEEDED``. Currently, this API does not support
	// exporting a library.
	//
	// Alternatively, one can download the exported file by enabling
	// ``direct_download``. Example: `curl -n -o example.scala
	// 'https://XX.cloud.databricks.com/api/2.0/workspace/export?path=/Users/user@example.com/ScalaExampleNotebook&direct_download=true'`
	Export(ctx context.Context, request ExportRequest) (*ExportResponse, error)

	// Get status
	//
	// Gets the status of an object or a directory. If ``path`` does not exist,
	// this call returns an error ``RESOURCE_DOES_NOT_EXIST``.
	GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error)

	// Import a notebook
	//
	// Imports a notebook or the contents of an entire directory. If ``path``
	// already exists and ``overwrite`` is set to ``false``, this call returns
	// an error ``RESOURCE_ALREADY_EXISTS``. One can only use ``DBC`` format to
	// import a directory.
	Import(ctx context.Context, request Import) error

	// List contents
	//
	// Lists the contents of a directory, or the object if it is not a
	// directory.If the input path does not exist, this call returns an error
	// ``RESOURCE_DOES_NOT_EXIST``.
	//
	// Use ListAll() to get all ObjectInfo instances
	List(ctx context.Context, request ListRequest) (*ListResponse, error)

	// Create a directory
	//
	// Creates the specified directory (and necessary parent directories if they
	// do not exist). If there is an object (not a directory) at any prefix of
	// the input path, this call returns an error ``RESOURCE_ALREADY_EXISTS``.
	//
	// Note that if this operation fails it may have succeeded in creating some
	// of the necessary\nparrent directories.
	Mkdirs(ctx context.Context, request Mkdirs) error
}
