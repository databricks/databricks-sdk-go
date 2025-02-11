// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboardspreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type GeniePreviewClient struct {
	GeniePreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewGeniePreviewClient(cfg *config.Config) (*GeniePreviewClient, error) {
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

	return &GeniePreviewClient{
		Config:                cfg,
		apiClient:             apiClient,
		GeniePreviewInterface: NewGeniePreview(databricksClient),
	}, nil
}

type LakeviewEmbeddedPreviewClient struct {
	LakeviewEmbeddedPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewLakeviewEmbeddedPreviewClient(cfg *config.Config) (*LakeviewEmbeddedPreviewClient, error) {
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

	return &LakeviewEmbeddedPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		LakeviewEmbeddedPreviewInterface: NewLakeviewEmbeddedPreview(databricksClient),
	}, nil
}

type LakeviewPreviewClient struct {
	LakeviewPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewLakeviewPreviewClient(cfg *config.Config) (*LakeviewPreviewClient, error) {
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

	return &LakeviewPreviewClient{
		Config:                   cfg,
		apiClient:                apiClient,
		LakeviewPreviewInterface: NewLakeviewPreview(databricksClient),
	}, nil
}

type QueryExecutionPreviewClient struct {
	QueryExecutionPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewQueryExecutionPreviewClient(cfg *config.Config) (*QueryExecutionPreviewClient, error) {
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

	return &QueryExecutionPreviewClient{
		Config:                         cfg,
		apiClient:                      apiClient,
		QueryExecutionPreviewInterface: NewQueryExecutionPreview(databricksClient),
	}, nil
}
