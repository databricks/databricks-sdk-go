// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanroomspreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type CleanRoomAssetsPreviewClient struct {
	CleanRoomAssetsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewCleanRoomAssetsPreviewClient(cfg *config.Config) (*CleanRoomAssetsPreviewClient, error) {
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

	return &CleanRoomAssetsPreviewClient{
		Config:                          cfg,
		apiClient:                       apiClient,
		CleanRoomAssetsPreviewInterface: NewCleanRoomAssetsPreview(databricksClient),
	}, nil
}

type CleanRoomTaskRunsPreviewClient struct {
	CleanRoomTaskRunsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewCleanRoomTaskRunsPreviewClient(cfg *config.Config) (*CleanRoomTaskRunsPreviewClient, error) {
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

	return &CleanRoomTaskRunsPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		CleanRoomTaskRunsPreviewInterface: NewCleanRoomTaskRunsPreview(databricksClient),
	}, nil
}

type CleanRoomsPreviewClient struct {
	CleanRoomsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewCleanRoomsPreviewClient(cfg *config.Config) (*CleanRoomsPreviewClient, error) {
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

	return &CleanRoomsPreviewClient{
		Config:                     cfg,
		apiClient:                  apiClient,
		CleanRoomsPreviewInterface: NewCleanRoomsPreview(databricksClient),
	}, nil
}
