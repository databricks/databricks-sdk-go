// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package filespreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type DbfsClient struct {
	DbfsInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &DbfsClient{
		Config:        cfg,
		apiClient:     apiClient,
		DbfsInterface: NewDbfs(databricksClient),
	}, nil
}

type FilesClient struct {
	FilesInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &FilesClient{
		Config:         cfg,
		apiClient:      apiClient,
		FilesInterface: NewFiles(databricksClient),
	}, nil
}
