// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type GenieClient struct {
	GenieAPI
}

func NewGenieClient(cfg *config.Config) (*GenieClient, error) {
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

	return &GenieClient{
		GenieAPI: GenieAPI{
			genieImpl: genieImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type LakeviewClient struct {
	LakeviewAPI
}

func NewLakeviewClient(cfg *config.Config) (*LakeviewClient, error) {
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

	return &LakeviewClient{
		LakeviewAPI: LakeviewAPI{
			lakeviewImpl: lakeviewImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type LakeviewEmbeddedClient struct {
	LakeviewEmbeddedAPI
}

func NewLakeviewEmbeddedClient(cfg *config.Config) (*LakeviewEmbeddedClient, error) {
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

	return &LakeviewEmbeddedClient{
		LakeviewEmbeddedAPI: LakeviewEmbeddedAPI{
			lakeviewEmbeddedImpl: lakeviewEmbeddedImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type QueryExecutionClient struct {
	QueryExecutionAPI
}

func NewQueryExecutionClient(cfg *config.Config) (*QueryExecutionClient, error) {
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

	return &QueryExecutionClient{
		QueryExecutionAPI: QueryExecutionAPI{
			queryExecutionImpl: queryExecutionImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
