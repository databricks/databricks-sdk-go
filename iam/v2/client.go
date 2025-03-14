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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AccessControlClient{
		apiClient:              databricksClient.ApiClient(),
		AccessControlInterface: NewAccessControl(databricksClient),
	}, nil
}

type AccountAccessControlClient struct {
	AccountAccessControlInterface
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
		AccountAccessControlInterface: NewAccountAccessControl(apiClient),
	}, nil
}

type AccountAccessControlProxyClient struct {
	AccountAccessControlProxyInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AccountAccessControlProxyClient{
		apiClient:                          databricksClient.ApiClient(),
		AccountAccessControlProxyInterface: NewAccountAccessControlProxy(databricksClient),
	}, nil
}

type AccountGroupsClient struct {
	AccountGroupsInterface
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
		AccountGroupsInterface: NewAccountGroups(apiClient),
	}, nil
}

type AccountServicePrincipalsClient struct {
	AccountServicePrincipalsInterface
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
		AccountServicePrincipalsInterface: NewAccountServicePrincipals(apiClient),
	}, nil
}

type AccountUsersClient struct {
	AccountUsersInterface
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
		AccountUsersInterface: NewAccountUsers(apiClient),
	}, nil
}

type CurrentUserClient struct {
	CurrentUserInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CurrentUserClient{
		apiClient:            databricksClient.ApiClient(),
		CurrentUserInterface: NewCurrentUser(databricksClient),
	}, nil
}

type GroupsClient struct {
	GroupsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GroupsClient{
		apiClient:       databricksClient.ApiClient(),
		GroupsInterface: NewGroups(databricksClient),
	}, nil
}

type PermissionMigrationClient struct {
	PermissionMigrationInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PermissionMigrationClient{
		apiClient:                    databricksClient.ApiClient(),
		PermissionMigrationInterface: NewPermissionMigration(databricksClient),
	}, nil
}

type PermissionsClient struct {
	PermissionsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PermissionsClient{
		apiClient:            databricksClient.ApiClient(),
		PermissionsInterface: NewPermissions(databricksClient),
	}, nil
}

type ServicePrincipalsClient struct {
	ServicePrincipalsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipalsClient{
		apiClient:                  databricksClient.ApiClient(),
		ServicePrincipalsInterface: NewServicePrincipals(databricksClient),
	}, nil
}

type UsersClient struct {
	UsersInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &UsersClient{
		apiClient:      databricksClient.ApiClient(),
		UsersInterface: NewUsers(databricksClient),
	}, nil
}

type WorkspaceAssignmentClient struct {
	WorkspaceAssignmentInterface
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
		WorkspaceAssignmentInterface: NewWorkspaceAssignment(apiClient),
	}, nil
}
