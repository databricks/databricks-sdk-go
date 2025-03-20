// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// DBFS API makes it simple to interact with various data sources without having
// to include a users credentials every time to read a file.
type DbfsClient struct {
	dbfsBaseClient
}

func NewDbfsClient(cfg *config.Config) (*DbfsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DbfsClient{
		dbfsBaseClient: dbfsBaseClient{
			dbfsImpl: dbfsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The Files API is a standard HTTP API that allows you to read, write, list,
// and delete files and directories by referring to their URI. The API makes
// working with file content as raw bytes easier and more efficient.
//
// The API supports [Unity Catalog volumes], where files and directories to
// operate on are specified using their volume URI path, which follows the
// format
// /Volumes/&lt;catalog_name&gt;/&lt;schema_name&gt;/&lt;volume_name&gt;/&lt;path_to_file&gt;.
//
// The Files API has two distinct endpoints, one for working with files
// (`/fs/files`) and another one for working with directories
// (`/fs/directories`). Both endpoints use the standard HTTP methods GET, HEAD,
// PUT, and DELETE to manage files and directories specified using their URI
// path. The path is always absolute.
//
// Some Files API client features are currently experimental. To enable them,
// set `enable_experimental_files_api_client = True` in your configuration
// profile or use the environment variable
// `DATABRICKS_ENABLE_EXPERIMENTAL_FILES_API_CLIENT=True`.
//
// [Unity Catalog volumes]: https://docs.databricks.com/en/connect/unity-catalog/volumes.html
type FilesClient struct {
	filesBaseClient
}

func NewFilesClient(cfg *config.Config) (*FilesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &FilesClient{
		filesBaseClient: filesBaseClient{
			filesImpl: filesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
