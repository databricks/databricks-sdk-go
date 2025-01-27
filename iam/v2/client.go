// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccessControlClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	AccessControl AccessControlInterface
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
		cfg:           cfg,
		apiClient:     apiClient,
		AccessControl: NewAccessControl(databricksClient),
	}, nil
}

type AccountAccessControlClient struct {
	cfg *config.Config

	AccountAccessControl AccountAccessControlInterface
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
		cfg:                  cfg,
		AccountAccessControl: NewAccountAccessControl(apiClient),
	}, nil
}

type AccountAccessControlProxyClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	AccountAccessControlProxy AccountAccessControlProxyInterface
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		AccountAccessControlProxy: NewAccountAccessControlProxy(databricksClient),
	}, nil
}

type AccountGroupsClient struct {
	cfg *config.Config

	AccountGroups AccountGroupsInterface
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
		cfg:           cfg,
		AccountGroups: NewAccountGroups(apiClient),
	}, nil
}

type AccountServicePrincipalsClient struct {
	cfg *config.Config

	AccountServicePrincipals AccountServicePrincipalsInterface
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
		cfg:                      cfg,
		AccountServicePrincipals: NewAccountServicePrincipals(apiClient),
	}, nil
}

type AccountUsersClient struct {
	cfg *config.Config

	AccountUsers AccountUsersInterface
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
		cfg:          cfg,
		AccountUsers: NewAccountUsers(apiClient),
	}, nil
}

type CurrentUserClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	CurrentUser CurrentUserInterface
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
		cfg:         cfg,
		apiClient:   apiClient,
		CurrentUser: NewCurrentUser(databricksClient),
	}, nil
}

type GroupsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Groups GroupsInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Groups:    NewGroups(databricksClient),
	}, nil
}

type PermissionMigrationClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	PermissionMigration PermissionMigrationInterface
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
		cfg:                 cfg,
		apiClient:           apiClient,
		PermissionMigration: NewPermissionMigration(databricksClient),
	}, nil
}

type PermissionsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Permissions PermissionsInterface
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
		cfg:         cfg,
		apiClient:   apiClient,
		Permissions: NewPermissions(databricksClient),
	}, nil
}

type ServicePrincipalsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ServicePrincipals ServicePrincipalsInterface
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
		cfg:               cfg,
		apiClient:         apiClient,
		ServicePrincipals: NewServicePrincipals(databricksClient),
	}, nil
}

type UsersClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Users UsersInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Users:     NewUsers(databricksClient),
	}, nil
}

type WorkspaceAssignmentClient struct {
	cfg *config.Config

	WorkspaceAssignment WorkspaceAssignmentInterface
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
		cfg:                 cfg,
		WorkspaceAssignment: NewWorkspaceAssignment(apiClient),
	}, nil
}
