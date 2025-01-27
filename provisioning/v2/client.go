// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type CredentialsClient struct {
	cfg *config.Config

	Credentials CredentialsInterface
}

func NewCredentialsClient(cfg *config.Config) (*CredentialsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CredentialsClient{
		cfg:         cfg,
		Credentials: NewCredentials(apiClient),
	}, nil
}

type EncryptionKeysClient struct {
	cfg *config.Config

	EncryptionKeys EncryptionKeysInterface
}

func NewEncryptionKeysClient(cfg *config.Config) (*EncryptionKeysClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &EncryptionKeysClient{
		cfg:            cfg,
		EncryptionKeys: NewEncryptionKeys(apiClient),
	}, nil
}

type NetworksClient struct {
	cfg *config.Config

	Networks NetworksInterface
}

func NewNetworksClient(cfg *config.Config) (*NetworksClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &NetworksClient{
		cfg:      cfg,
		Networks: NewNetworks(apiClient),
	}, nil
}

type PrivateAccessClient struct {
	cfg *config.Config

	PrivateAccess PrivateAccessInterface
}

func NewPrivateAccessClient(cfg *config.Config) (*PrivateAccessClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PrivateAccessClient{
		cfg:           cfg,
		PrivateAccess: NewPrivateAccess(apiClient),
	}, nil
}

type StorageClient struct {
	cfg *config.Config

	Storage StorageInterface
}

func NewStorageClient(cfg *config.Config) (*StorageClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &StorageClient{
		cfg:     cfg,
		Storage: NewStorage(apiClient),
	}, nil
}

type VpcEndpointsClient struct {
	cfg *config.Config

	VpcEndpoints VpcEndpointsInterface
}

func NewVpcEndpointsClient(cfg *config.Config) (*VpcEndpointsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &VpcEndpointsClient{
		cfg:          cfg,
		VpcEndpoints: NewVpcEndpoints(apiClient),
	}, nil
}

type WorkspacesClient struct {
	cfg *config.Config

	Workspaces WorkspacesInterface
}

func NewWorkspacesClient(cfg *config.Config) (*WorkspacesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WorkspacesClient{
		cfg:        cfg,
		Workspaces: NewWorkspaces(apiClient),
	}, nil
}
