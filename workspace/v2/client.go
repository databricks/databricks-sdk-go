// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type GitCredentialsClient struct {
	GitCredentialsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GitCredentialsClient{
		apiClient:               databricksClient.ApiClient(),
		GitCredentialsInterface: NewGitCredentials(databricksClient),
	}, nil
}

type ReposClient struct {
	ReposInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ReposClient{
		apiClient:      databricksClient.ApiClient(),
		ReposInterface: NewRepos(databricksClient),
	}, nil
}

type SecretsClient struct {
	SecretsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &SecretsClient{
		apiClient:        databricksClient.ApiClient(),
		SecretsInterface: NewSecrets(databricksClient),
	}, nil
}

type WorkspaceClient struct {
	WorkspaceInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WorkspaceClient{
		apiClient:          databricksClient.ApiClient(),
		WorkspaceInterface: NewWorkspace(databricksClient),
	}, nil
}
