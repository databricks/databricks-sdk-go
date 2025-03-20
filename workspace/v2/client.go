// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type GitCredentialsClient struct {
	GitCredentialsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GitCredentialsClient{
		GitCredentialsAPI: GitCredentialsAPI{
			gitCredentialsImpl: gitCredentialsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ReposClient struct {
	ReposAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ReposClient{
		ReposAPI: ReposAPI{
			reposImpl: reposImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type SecretsClient struct {
	SecretsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &SecretsClient{
		SecretsAPI: SecretsAPI{
			secretsImpl: secretsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type WorkspaceClient struct {
	WorkspaceAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WorkspaceClient{
		WorkspaceAPI: WorkspaceAPI{
			workspaceImpl: workspaceImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
