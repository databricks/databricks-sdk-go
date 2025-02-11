// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iampreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccessControlPreviewClient struct {
	AccessControlPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewAccessControlPreviewClient(cfg *config.Config) (*AccessControlPreviewClient, error) {
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

	return &AccessControlPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		AccessControlPreviewInterface: NewAccessControlPreview(databricksClient),
	}, nil
}

type AccountAccessControlPreviewClient struct {
	AccountAccessControlPreviewInterface

	Config *config.Config
}

func NewAccountAccessControlPreviewClient(cfg *config.Config) (*AccountAccessControlPreviewClient, error) {
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

	return &AccountAccessControlPreviewClient{
		Config:                               cfg,
		AccountAccessControlPreviewInterface: NewAccountAccessControlPreview(apiClient),
	}, nil
}

type AccountAccessControlProxyPreviewClient struct {
	AccountAccessControlProxyPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewAccountAccessControlProxyPreviewClient(cfg *config.Config) (*AccountAccessControlProxyPreviewClient, error) {
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

	return &AccountAccessControlProxyPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		AccountAccessControlProxyPreviewInterface: NewAccountAccessControlProxyPreview(databricksClient),
	}, nil
}

type AccountGroupsPreviewClient struct {
	AccountGroupsPreviewInterface

	Config *config.Config
}

func NewAccountGroupsPreviewClient(cfg *config.Config) (*AccountGroupsPreviewClient, error) {
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

	return &AccountGroupsPreviewClient{
		Config:                        cfg,
		AccountGroupsPreviewInterface: NewAccountGroupsPreview(apiClient),
	}, nil
}

type AccountServicePrincipalsPreviewClient struct {
	AccountServicePrincipalsPreviewInterface

	Config *config.Config
}

func NewAccountServicePrincipalsPreviewClient(cfg *config.Config) (*AccountServicePrincipalsPreviewClient, error) {
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

	return &AccountServicePrincipalsPreviewClient{
		Config:                                   cfg,
		AccountServicePrincipalsPreviewInterface: NewAccountServicePrincipalsPreview(apiClient),
	}, nil
}

type AccountUsersPreviewClient struct {
	AccountUsersPreviewInterface

	Config *config.Config
}

func NewAccountUsersPreviewClient(cfg *config.Config) (*AccountUsersPreviewClient, error) {
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

	return &AccountUsersPreviewClient{
		Config:                       cfg,
		AccountUsersPreviewInterface: NewAccountUsersPreview(apiClient),
	}, nil
}

type CurrentUserPreviewClient struct {
	CurrentUserPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewCurrentUserPreviewClient(cfg *config.Config) (*CurrentUserPreviewClient, error) {
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

	return &CurrentUserPreviewClient{
		Config:                      cfg,
		apiClient:                   apiClient,
		CurrentUserPreviewInterface: NewCurrentUserPreview(databricksClient),
	}, nil
}

type GroupsPreviewClient struct {
	GroupsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewGroupsPreviewClient(cfg *config.Config) (*GroupsPreviewClient, error) {
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

	return &GroupsPreviewClient{
		Config:                 cfg,
		apiClient:              apiClient,
		GroupsPreviewInterface: NewGroupsPreview(databricksClient),
	}, nil
}

type PermissionMigrationPreviewClient struct {
	PermissionMigrationPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewPermissionMigrationPreviewClient(cfg *config.Config) (*PermissionMigrationPreviewClient, error) {
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

	return &PermissionMigrationPreviewClient{
		Config:                              cfg,
		apiClient:                           apiClient,
		PermissionMigrationPreviewInterface: NewPermissionMigrationPreview(databricksClient),
	}, nil
}

type PermissionsPreviewClient struct {
	PermissionsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewPermissionsPreviewClient(cfg *config.Config) (*PermissionsPreviewClient, error) {
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

	return &PermissionsPreviewClient{
		Config:                      cfg,
		apiClient:                   apiClient,
		PermissionsPreviewInterface: NewPermissionsPreview(databricksClient),
	}, nil
}

type ServicePrincipalsPreviewClient struct {
	ServicePrincipalsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewServicePrincipalsPreviewClient(cfg *config.Config) (*ServicePrincipalsPreviewClient, error) {
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

	return &ServicePrincipalsPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		ServicePrincipalsPreviewInterface: NewServicePrincipalsPreview(databricksClient),
	}, nil
}

type UsersPreviewClient struct {
	UsersPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewUsersPreviewClient(cfg *config.Config) (*UsersPreviewClient, error) {
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

	return &UsersPreviewClient{
		Config:                cfg,
		apiClient:             apiClient,
		UsersPreviewInterface: NewUsersPreview(databricksClient),
	}, nil
}

type WorkspaceAssignmentPreviewClient struct {
	WorkspaceAssignmentPreviewInterface

	Config *config.Config
}

func NewWorkspaceAssignmentPreviewClient(cfg *config.Config) (*WorkspaceAssignmentPreviewClient, error) {
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

	return &WorkspaceAssignmentPreviewClient{
		Config:                              cfg,
		WorkspaceAssignmentPreviewInterface: NewWorkspaceAssignmentPreview(apiClient),
	}, nil
}
