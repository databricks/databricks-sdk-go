// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package filespreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type DbfsPreviewClient struct {
	DbfsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewDbfsPreviewClient(cfg *config.Config) (*DbfsPreviewClient, error) {
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

	return &DbfsPreviewClient{
		Config:               cfg,
		apiClient:            apiClient,
		DbfsPreviewInterface: NewDbfsPreview(databricksClient),
	}, nil
}

type FilesPreviewClient struct {
	FilesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewFilesPreviewClient(cfg *config.Config) (*FilesPreviewClient, error) {
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

	return &FilesPreviewClient{
		Config:                cfg,
		apiClient:             apiClient,
		FilesPreviewInterface: NewFilesPreview(databricksClient),
	}, nil
}
