// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type GitCredentialsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	GitCredentials GitCredentialsInterface
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
		cfg:            cfg,
		apiClient:      apiClient,
		GitCredentials: NewGitCredentials(databricksClient),
	}, nil
}

type ReposClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Repos ReposInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Repos:     NewRepos(databricksClient),
	}, nil
}

type SecretsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Secrets SecretsInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Secrets:   NewSecrets(databricksClient),
	}, nil
}

type WorkspaceClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Workspace WorkspaceInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Workspace: NewWorkspace(databricksClient),
	}, nil
}
