// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type CredentialsClient struct {
	CredentialsInterface

	cfg *config.Config
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
		cfg:                  cfg,
		CredentialsInterface: NewCredentials(apiClient),
	}, nil
}

type EncryptionKeysClient struct {
	EncryptionKeysInterface

	cfg *config.Config
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
		cfg:                     cfg,
		EncryptionKeysInterface: NewEncryptionKeys(apiClient),
	}, nil
}

type NetworksClient struct {
	NetworksInterface

	cfg *config.Config
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
		cfg:               cfg,
		NetworksInterface: NewNetworks(apiClient),
	}, nil
}

type PrivateAccessClient struct {
	PrivateAccessInterface

	cfg *config.Config
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
		cfg:                    cfg,
		PrivateAccessInterface: NewPrivateAccess(apiClient),
	}, nil
}

type StorageClient struct {
	StorageInterface

	cfg *config.Config
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
		cfg:              cfg,
		StorageInterface: NewStorage(apiClient),
	}, nil
}

type VpcEndpointsClient struct {
	VpcEndpointsInterface

	cfg *config.Config
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
		cfg:                   cfg,
		VpcEndpointsInterface: NewVpcEndpoints(apiClient),
	}, nil
}

type WorkspacesClient struct {
	WorkspacesInterface

	cfg *config.Config
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
		cfg:                 cfg,
		WorkspacesInterface: NewWorkspaces(apiClient),
	}, nil
}
