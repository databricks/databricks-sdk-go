// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspacepreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type GitCredentialsClient struct {
	GitCredentialsInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewGitCredentialsClient(cfg *config.Config) (*GitCredentialsClient, error) {
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

	return &GitCredentialsClient{
		Config:                  cfg,
		apiClient:               apiClient,
		GitCredentialsInterface: NewGitCredentials(databricksClient),
	}, nil
}

type ReposClient struct {
	ReposInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewReposClient(cfg *config.Config) (*ReposClient, error) {
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

	return &ReposClient{
		Config:         cfg,
		apiClient:      apiClient,
		ReposInterface: NewRepos(databricksClient),
	}, nil
}

type SecretsClient struct {
	SecretsInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewSecretsClient(cfg *config.Config) (*SecretsClient, error) {
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

	return &SecretsClient{
		Config:           cfg,
		apiClient:        apiClient,
		SecretsInterface: NewSecrets(databricksClient),
	}, nil
}

type WorkspaceClient struct {
	WorkspaceInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewWorkspaceClient(cfg *config.Config) (*WorkspaceClient, error) {
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

	return &WorkspaceClient{
		Config:             cfg,
		apiClient:          apiClient,
		WorkspaceInterface: NewWorkspace(databricksClient),
	}, nil
}
