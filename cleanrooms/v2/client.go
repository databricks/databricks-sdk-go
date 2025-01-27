// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type CleanRoomAssetsClient struct {
	CleanRoomAssetsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewCleanRoomAssetsClient(cfg *config.Config) (*CleanRoomAssetsClient, error) {
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

	return &CleanRoomAssetsClient{
		cfg:                      cfg,
		apiClient:                apiClient,
		CleanRoomAssetsInterface: NewCleanRoomAssets(databricksClient),
	}, nil
}

type CleanRoomTaskRunsClient struct {
	CleanRoomTaskRunsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewCleanRoomTaskRunsClient(cfg *config.Config) (*CleanRoomTaskRunsClient, error) {
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

	return &CleanRoomTaskRunsClient{
		cfg:                        cfg,
		apiClient:                  apiClient,
		CleanRoomTaskRunsInterface: NewCleanRoomTaskRuns(databricksClient),
	}, nil
}

type CleanRoomsClient struct {
	CleanRoomsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewCleanRoomsClient(cfg *config.Config) (*CleanRoomsClient, error) {
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

	return &CleanRoomsClient{
		cfg:                 cfg,
		apiClient:           apiClient,
		CleanRoomsInterface: NewCleanRooms(databricksClient),
	}, nil
}
