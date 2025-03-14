// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccessControlClient struct {
	AccessControlInterface
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
		CurrentUserInterface: NewCurrentUser(databricksClient),
	}, nil
}

type GroupsClient struct {
	GroupsInterface
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
		GroupsInterface: NewGroups(databricksClient),
	}, nil
}

type PermissionMigrationClient struct {
	PermissionMigrationInterface
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
		PermissionMigrationInterface: NewPermissionMigration(databricksClient),
	}, nil
}

type PermissionsClient struct {
	PermissionsInterface
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
		PermissionsInterface: NewPermissions(databricksClient),
	}, nil
}

type ServicePrincipalsClient struct {
	ServicePrincipalsInterface
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
		ServicePrincipalsInterface: NewServicePrincipals(databricksClient),
	}, nil
}

type UsersClient struct {
	UsersInterface
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
