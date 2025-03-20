// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccessControlClient struct {
	AccessControlAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AccessControlClient{
		AccessControlAPI: AccessControlAPI{
			accessControlImpl: accessControlImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type AccountAccessControlClient struct {
	AccountAccessControlAPI
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
		AccountAccessControlAPI: AccountAccessControlAPI{
			accountAccessControlImpl: accountAccessControlImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type AccountAccessControlProxyClient struct {
	AccountAccessControlProxyAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AccountAccessControlProxyClient{
		AccountAccessControlProxyAPI: AccountAccessControlProxyAPI{
			accountAccessControlProxyImpl: accountAccessControlProxyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type AccountGroupsClient struct {
	AccountGroupsAPI
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
		AccountGroupsAPI: AccountGroupsAPI{
			accountGroupsImpl: accountGroupsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type AccountServicePrincipalsClient struct {
	AccountServicePrincipalsAPI
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
		AccountServicePrincipalsAPI: AccountServicePrincipalsAPI{
			accountServicePrincipalsImpl: accountServicePrincipalsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type AccountUsersClient struct {
	AccountUsersAPI
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
		AccountUsersAPI: AccountUsersAPI{
			accountUsersImpl: accountUsersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type CurrentUserClient struct {
	CurrentUserAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CurrentUserClient{
		CurrentUserAPI: CurrentUserAPI{
			currentUserImpl: currentUserImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type GroupsClient struct {
	GroupsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GroupsClient{
		GroupsAPI: GroupsAPI{
			groupsImpl: groupsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type PermissionMigrationClient struct {
	PermissionMigrationAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PermissionMigrationClient{
		PermissionMigrationAPI: PermissionMigrationAPI{
			permissionMigrationImpl: permissionMigrationImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type PermissionsClient struct {
	PermissionsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PermissionsClient{
		PermissionsAPI: PermissionsAPI{
			permissionsImpl: permissionsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ServicePrincipalsClient struct {
	ServicePrincipalsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipalsClient{
		ServicePrincipalsAPI: ServicePrincipalsAPI{
			servicePrincipalsImpl: servicePrincipalsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type UsersClient struct {
	UsersAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &UsersClient{
		UsersAPI: UsersAPI{
			usersImpl: usersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type WorkspaceAssignmentClient struct {
	WorkspaceAssignmentAPI
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
		WorkspaceAssignmentAPI: WorkspaceAssignmentAPI{
			workspaceAssignmentImpl: workspaceAssignmentImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
