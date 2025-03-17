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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GenieClient{
		apiClient:      databricksClient.ApiClient(),
		GenieInterface: NewGenie(databricksClient),
	}, nil
}

type LakeviewClient struct {
	LakeviewInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &LakeviewClient{
		apiClient:         databricksClient.ApiClient(),
		LakeviewInterface: NewLakeview(databricksClient),
	}, nil
}

type LakeviewEmbeddedClient struct {
	LakeviewEmbeddedInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &LakeviewEmbeddedClient{
		apiClient:                 databricksClient.ApiClient(),
		LakeviewEmbeddedInterface: NewLakeviewEmbedded(databricksClient),
	}, nil
}

type QueryExecutionClient struct {
	QueryExecutionInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryExecutionClient{
		apiClient:               databricksClient.ApiClient(),
		QueryExecutionInterface: NewQueryExecution(databricksClient),
	}, nil
}
