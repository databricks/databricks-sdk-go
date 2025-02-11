// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelinespreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type PipelinesClient struct {
	PipelinesInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &PipelinesClient{
		Config:             cfg,
		apiClient:          apiClient,
		PipelinesInterface: NewPipelines(databricksClient),
	}, nil
}
