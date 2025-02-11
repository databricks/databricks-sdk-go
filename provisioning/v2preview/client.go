// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioningpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type CredentialsClient struct {
	CredentialsInterface

	Config *config.Config
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
		Config:               cfg,
		CredentialsInterface: NewCredentials(apiClient),
	}, nil
}

type EncryptionKeysClient struct {
	EncryptionKeysInterface

	Config *config.Config
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
		Config:                  cfg,
		EncryptionKeysInterface: NewEncryptionKeys(apiClient),
	}, nil
}

type NetworksClient struct {
	NetworksInterface

	Config *config.Config
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
		Config:            cfg,
		NetworksInterface: NewNetworks(apiClient),
	}, nil
}

type PrivateAccessClient struct {
	PrivateAccessInterface

	Config *config.Config
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
		Config:                 cfg,
		PrivateAccessInterface: NewPrivateAccess(apiClient),
	}, nil
}

type StorageClient struct {
	StorageInterface

	Config *config.Config
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
		Config:           cfg,
		StorageInterface: NewStorage(apiClient),
	}, nil
}

type VpcEndpointsClient struct {
	VpcEndpointsInterface

	Config *config.Config
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
		Config:                cfg,
		VpcEndpointsInterface: NewVpcEndpoints(apiClient),
	}, nil
}

type WorkspacesClient struct {
	WorkspacesInterface

	Config *config.Config
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
		Config:              cfg,
		WorkspacesInterface: NewWorkspaces(apiClient),
	}, nil
}
