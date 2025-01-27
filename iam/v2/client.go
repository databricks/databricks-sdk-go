// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccessControlClient struct {
	AccessControlInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewAccessControlClient(cfg *config.Config) (*AccessControlClient, error) {
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

	return &AccessControlClient{
		cfg:                    cfg,
		apiClient:              apiClient,
		AccessControlInterface: NewAccessControl(databricksClient),
	}, nil
}

type AccountAccessControlClient struct {
	AccountAccessControlInterface

	cfg *config.Config
}

func NewAccountAccessControlClient(cfg *config.Config) (*AccountAccessControlClient, error) {
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

	return &AccountAccessControlClient{
		cfg:                           cfg,
		AccountAccessControlInterface: NewAccountAccessControl(apiClient),
	}, nil
}

type AccountAccessControlProxyClient struct {
	AccountAccessControlProxyInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewAccountAccessControlProxyClient(cfg *config.Config) (*AccountAccessControlProxyClient, error) {
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

	return &AccountAccessControlProxyClient{
		cfg:                                cfg,
		apiClient:                          apiClient,
		AccountAccessControlProxyInterface: NewAccountAccessControlProxy(databricksClient),
	}, nil
}

type AccountGroupsClient struct {
	AccountGroupsInterface

	cfg *config.Config
}

func NewAccountGroupsClient(cfg *config.Config) (*AccountGroupsClient, error) {
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

	return &AccountGroupsClient{
		cfg:                    cfg,
		AccountGroupsInterface: NewAccountGroups(apiClient),
	}, nil
}

type AccountServicePrincipalsClient struct {
	AccountServicePrincipalsInterface

	cfg *config.Config
}

func NewAccountServicePrincipalsClient(cfg *config.Config) (*AccountServicePrincipalsClient, error) {
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

	return &AccountServicePrincipalsClient{
		cfg:                               cfg,
		AccountServicePrincipalsInterface: NewAccountServicePrincipals(apiClient),
	}, nil
}

type AccountUsersClient struct {
	AccountUsersInterface

	cfg *config.Config
}

func NewAccountUsersClient(cfg *config.Config) (*AccountUsersClient, error) {
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

	return &AccountUsersClient{
		cfg:                   cfg,
		AccountUsersInterface: NewAccountUsers(apiClient),
	}, nil
}

type CurrentUserClient struct {
	CurrentUserInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewCurrentUserClient(cfg *config.Config) (*CurrentUserClient, error) {
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

	return &CurrentUserClient{
		cfg:                  cfg,
		apiClient:            apiClient,
		CurrentUserInterface: NewCurrentUser(databricksClient),
	}, nil
}

type GroupsClient struct {
	GroupsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewGroupsClient(cfg *config.Config) (*GroupsClient, error) {
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

	return &GroupsClient{
		cfg:             cfg,
		apiClient:       apiClient,
		GroupsInterface: NewGroups(databricksClient),
	}, nil
}

type PermissionMigrationClient struct {
	PermissionMigrationInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewPermissionMigrationClient(cfg *config.Config) (*PermissionMigrationClient, error) {
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

	return &PermissionMigrationClient{
		cfg:                          cfg,
		apiClient:                    apiClient,
		PermissionMigrationInterface: NewPermissionMigration(databricksClient),
	}, nil
}

type PermissionsClient struct {
	PermissionsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewPermissionsClient(cfg *config.Config) (*PermissionsClient, error) {
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

	return &PermissionsClient{
		cfg:                  cfg,
		apiClient:            apiClient,
		PermissionsInterface: NewPermissions(databricksClient),
	}, nil
}

type ServicePrincipalsClient struct {
	ServicePrincipalsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewServicePrincipalsClient(cfg *config.Config) (*ServicePrincipalsClient, error) {
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

	return &ServicePrincipalsClient{
		cfg:                        cfg,
		apiClient:                  apiClient,
		ServicePrincipalsInterface: NewServicePrincipals(databricksClient),
	}, nil
}

type UsersClient struct {
	UsersInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
}

func NewUsersClient(cfg *config.Config) (*UsersClient, error) {
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

	return &UsersClient{
		cfg:            cfg,
		apiClient:      apiClient,
		UsersInterface: NewUsers(databricksClient),
	}, nil
}

type WorkspaceAssignmentClient struct {
	WorkspaceAssignmentInterface

	cfg *config.Config
}

func NewWorkspaceAssignmentClient(cfg *config.Config) (*WorkspaceAssignmentClient, error) {
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

	return &WorkspaceAssignmentClient{
		cfg:                          cfg,
		WorkspaceAssignmentInterface: NewWorkspaceAssignment(apiClient),
	}, nil
}
