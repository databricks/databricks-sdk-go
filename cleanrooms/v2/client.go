// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type CleanRoomAssetsClient struct {
	CleanRoomAssetsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CleanRoomAssetsClient{
		CleanRoomAssetsAPI: CleanRoomAssetsAPI{
			cleanRoomAssetsImpl: cleanRoomAssetsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type CleanRoomTaskRunsClient struct {
	CleanRoomTaskRunsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CleanRoomTaskRunsClient{
		CleanRoomTaskRunsAPI: CleanRoomTaskRunsAPI{
			cleanRoomTaskRunsImpl: cleanRoomTaskRunsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type CleanRoomsClient struct {
	CleanRoomsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CleanRoomsClient{
		CleanRoomsAPI: CleanRoomsAPI{
			cleanRoomsImpl: cleanRoomsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
