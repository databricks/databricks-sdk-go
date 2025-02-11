// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspacepreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type GitCredentialsPreviewClient struct {
	GitCredentialsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewGitCredentialsPreviewClient(cfg *config.Config) (*GitCredentialsPreviewClient, error) {
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

	return &GitCredentialsPreviewClient{
		Config:                         cfg,
		apiClient:                      apiClient,
		GitCredentialsPreviewInterface: NewGitCredentialsPreview(databricksClient),
	}, nil
}

type ReposPreviewClient struct {
	ReposPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewReposPreviewClient(cfg *config.Config) (*ReposPreviewClient, error) {
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

	return &ReposPreviewClient{
		Config:                cfg,
		apiClient:             apiClient,
		ReposPreviewInterface: NewReposPreview(databricksClient),
	}, nil
}

type SecretsPreviewClient struct {
	SecretsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewSecretsPreviewClient(cfg *config.Config) (*SecretsPreviewClient, error) {
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

	return &SecretsPreviewClient{
		Config:                  cfg,
		apiClient:               apiClient,
		SecretsPreviewInterface: NewSecretsPreview(databricksClient),
	}, nil
}

type WorkspacePreviewClient struct {
	WorkspacePreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewWorkspacePreviewClient(cfg *config.Config) (*WorkspacePreviewClient, error) {
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

	return &WorkspacePreviewClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		WorkspacePreviewInterface: NewWorkspacePreview(databricksClient),
	}, nil
}
