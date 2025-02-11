// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioningpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type CredentialsPreviewClient struct {
	CredentialsPreviewInterface

	Config *config.Config
}

func NewCredentialsPreviewClient(cfg *config.Config) (*CredentialsPreviewClient, error) {
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

	return &CredentialsPreviewClient{
		Config:                      cfg,
		CredentialsPreviewInterface: NewCredentialsPreview(apiClient),
	}, nil
}

type EncryptionKeysPreviewClient struct {
	EncryptionKeysPreviewInterface

	Config *config.Config
}

func NewEncryptionKeysPreviewClient(cfg *config.Config) (*EncryptionKeysPreviewClient, error) {
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

	return &EncryptionKeysPreviewClient{
		Config:                         cfg,
		EncryptionKeysPreviewInterface: NewEncryptionKeysPreview(apiClient),
	}, nil
}

type NetworksPreviewClient struct {
	NetworksPreviewInterface

	Config *config.Config
}

func NewNetworksPreviewClient(cfg *config.Config) (*NetworksPreviewClient, error) {
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

	return &NetworksPreviewClient{
		Config:                   cfg,
		NetworksPreviewInterface: NewNetworksPreview(apiClient),
	}, nil
}

type PrivateAccessPreviewClient struct {
	PrivateAccessPreviewInterface

	Config *config.Config
}

func NewPrivateAccessPreviewClient(cfg *config.Config) (*PrivateAccessPreviewClient, error) {
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

	return &PrivateAccessPreviewClient{
		Config:                        cfg,
		PrivateAccessPreviewInterface: NewPrivateAccessPreview(apiClient),
	}, nil
}

type StoragePreviewClient struct {
	StoragePreviewInterface

	Config *config.Config
}

func NewStoragePreviewClient(cfg *config.Config) (*StoragePreviewClient, error) {
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

	return &StoragePreviewClient{
		Config:                  cfg,
		StoragePreviewInterface: NewStoragePreview(apiClient),
	}, nil
}

type VpcEndpointsPreviewClient struct {
	VpcEndpointsPreviewInterface

	Config *config.Config
}

func NewVpcEndpointsPreviewClient(cfg *config.Config) (*VpcEndpointsPreviewClient, error) {
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

	return &VpcEndpointsPreviewClient{
		Config:                       cfg,
		VpcEndpointsPreviewInterface: NewVpcEndpointsPreview(apiClient),
	}, nil
}

type WorkspacesPreviewClient struct {
	WorkspacesPreviewInterface

	Config *config.Config
}

func NewWorkspacesPreviewClient(cfg *config.Config) (*WorkspacesPreviewClient, error) {
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

	return &WorkspacesPreviewClient{
		Config:                     cfg,
		WorkspacesPreviewInterface: NewWorkspacesPreview(apiClient),
	}, nil
}
