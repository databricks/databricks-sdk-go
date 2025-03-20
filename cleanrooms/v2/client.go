// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// Clean room assets are data and code objects — Tables, volumes, and
// notebooks that are shared with the clean room.
type CleanRoomAssetsClient struct {
	cleanRoomAssetsBaseClient
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
		cleanRoomAssetsBaseClient: cleanRoomAssetsBaseClient{
			cleanRoomAssetsImpl: cleanRoomAssetsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Clean room task runs are the executions of notebooks in a clean room.
type CleanRoomTaskRunsClient struct {
	cleanRoomTaskRunsBaseClient
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
		cleanRoomTaskRunsBaseClient: cleanRoomTaskRunsBaseClient{
			cleanRoomTaskRunsImpl: cleanRoomTaskRunsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// A clean room uses Delta Sharing and serverless compute to provide a secure
// and privacy-protecting environment where multiple parties can work together
// on sensitive enterprise data without direct access to each other’s data.
type CleanRoomsClient struct {
	cleanRoomsBaseClient
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
		cleanRoomsBaseClient: cleanRoomsBaseClient{
			cleanRoomsImpl: cleanRoomsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
