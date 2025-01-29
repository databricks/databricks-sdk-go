// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type GenieClient struct {
	GenieInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &GenieClient{
		Config:         cfg,
		apiClient:      apiClient,
		GenieInterface: NewGenie(databricksClient),
	}, nil
}

type LakeviewClient struct {
	LakeviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &LakeviewClient{
		Config:            cfg,
		apiClient:         apiClient,
		LakeviewInterface: NewLakeview(databricksClient),
	}, nil
}
