// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type PipelinesClient struct {
	PipelinesAPI
}

func NewPipelinesClient(cfg *config.Config) (*PipelinesClient, error) {
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

	return &PipelinesClient{
		PipelinesAPI: PipelinesAPI{
			pipelinesImpl: pipelinesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
