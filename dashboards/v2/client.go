// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type GenieClient struct {
	GenieInterface
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
		GenieInterface: NewGenie(apiClient),
	}, nil
}

type LakeviewClient struct {
	LakeviewInterface
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
		LakeviewInterface: NewLakeview(apiClient),
	}, nil
}

type LakeviewEmbeddedClient struct {
	LakeviewEmbeddedInterface
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
		LakeviewEmbeddedInterface: NewLakeviewEmbedded(apiClient),
	}, nil
}

type QueryExecutionClient struct {
	QueryExecutionInterface
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
		QueryExecutionInterface: NewQueryExecution(apiClient),
	}, nil
}
