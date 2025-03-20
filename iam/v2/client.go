// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// Rule based Access Control for Databricks Resources.
type AccessControlClient struct {
	accessControlBaseClient
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
		accessControlBaseClient: accessControlBaseClient{
			accessControlImpl: accessControlImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set.
type AccountAccessControlClient struct {
	accountAccessControlBaseClient
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
		accountAccessControlBaseClient: accountAccessControlBaseClient{
			accountAccessControlImpl: accountAccessControlImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set. A
// workspace must belong to an account for these APIs to work.
type AccountAccessControlProxyClient struct {
	accountAccessControlProxyBaseClient
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
		accountAccessControlProxyBaseClient: accountAccessControlProxyBaseClient{
			accountAccessControlProxyImpl: accountAccessControlProxyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Groups simplify identity management, making it easier to assign access to
// Databricks account, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks account identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
type AccountGroupsClient struct {
	accountGroupsBaseClient
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
		accountGroupsBaseClient: accountGroupsBaseClient{
			accountGroupsImpl: accountGroupsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
type AccountServicePrincipalsClient struct {
	accountServicePrincipalsBaseClient
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
		accountServicePrincipalsBaseClient: accountServicePrincipalsBaseClient{
			accountServicePrincipalsImpl: accountServicePrincipalsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// User identities recognized by Databricks and represented by email addresses.
//
// Databricks recommends using SCIM provisioning to sync users and groups
// automatically from your identity provider to your Databricks account. SCIM
// streamlines onboarding a new employee or team by using your identity provider
// to create users and groups in Databricks account and give them the proper
// level of access. When a user leaves your organization or no longer needs
// access to Databricks account, admins can terminate the user in your identity
// provider and that user’s account will also be removed from Databricks
// account. This ensures a consistent offboarding process and prevents
// unauthorized users from accessing sensitive data.
type AccountUsersClient struct {
	accountUsersBaseClient
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
		accountUsersBaseClient: accountUsersBaseClient{
			accountUsersImpl: accountUsersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// This API allows retrieving information about currently authenticated user or
// service principal.
type CurrentUserClient struct {
	currentUserBaseClient
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
		currentUserBaseClient: currentUserBaseClient{
			currentUserImpl: currentUserImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Groups simplify identity management, making it easier to assign access to
// Databricks workspace, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks workspace identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
type GroupsClient struct {
	groupsBaseClient
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
		groupsBaseClient: groupsBaseClient{
			groupsImpl: groupsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// APIs for migrating acl permissions, used only by the ucx tool:
// https://github.com/databrickslabs/ucx
type PermissionMigrationClient struct {
	permissionMigrationBaseClient
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
		permissionMigrationBaseClient: permissionMigrationBaseClient{
			permissionMigrationImpl: permissionMigrationImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints.
//
// * **[Apps permissions](:service:apps)** — Manage which users can manage or
// use apps.
//
// * **[Cluster permissions](:service:clusters)** — Manage which users can
// manage, restart, or attach to clusters.
//
// * **[Cluster policy permissions](:service:clusterpolicies)** — Manage which
// users can use cluster policies.
//
// * **[Delta Live Tables pipeline permissions](:service:pipelines)** — Manage
// which users can view, manage, run, cancel, or own a Delta Live Tables
// pipeline.
//
// * **[Job permissions](:service:jobs)** — Manage which users can view,
// manage, trigger, cancel, or own a job.
//
// * **[MLflow experiment permissions](:service:experiments)** — Manage which
// users can read, edit, or manage MLflow experiments.
//
// * **[MLflow registered model permissions](:service:modelregistry)** —
// Manage which users can read, edit, or manage MLflow registered models.
//
// * **[Password permissions](:service:users)** — Manage which users can use
// password login when SSO is enabled.
//
// * **[Instance Pool permissions](:service:instancepools)** — Manage which
// users can manage or attach to pools.
//
// * **[Repo permissions](repos)** — Manage which users can read, run, edit,
// or manage a repo.
//
// * **[Serving endpoint permissions](:service:servingendpoints)** — Manage
// which users can view, query, or manage a serving endpoint.
//
// * **[SQL warehouse permissions](:service:warehouses)** — Manage which users
// can use or manage SQL warehouses.
//
// * **[Token permissions](:service:tokenmanagement)** — Manage which users
// can create or use tokens.
//
// * **[Workspace object permissions](:service:workspace)** — Manage which
// users can read, run, edit, or manage alerts, dbsql-dashboards, directories,
// files, notebooks and queries.
//
// For the mapping of the required permissions for specific actions or abilities
// and other important information, see [Access Control].
//
// Note that to manage access control on service principals, use **[Account
// Access Control Proxy](:service:accountaccesscontrolproxy)**.
//
// [Access Control]: https://docs.databricks.com/security/auth-authz/access-control/index.html
type PermissionsClient struct {
	permissionsBaseClient
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
		permissionsBaseClient: permissionsBaseClient{
			permissionsImpl: permissionsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
type ServicePrincipalsClient struct {
	servicePrincipalsBaseClient
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
		servicePrincipalsBaseClient: servicePrincipalsBaseClient{
			servicePrincipalsImpl: servicePrincipalsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// User identities recognized by Databricks and represented by email addresses.
//
// Databricks recommends using SCIM provisioning to sync users and groups
// automatically from your identity provider to your Databricks workspace. SCIM
// streamlines onboarding a new employee or team by using your identity provider
// to create users and groups in Databricks workspace and give them the proper
// level of access. When a user leaves your organization or no longer needs
// access to Databricks workspace, admins can terminate the user in your
// identity provider and that user’s account will also be removed from
// Databricks workspace. This ensures a consistent offboarding process and
// prevents unauthorized users from accessing sensitive data.
type UsersClient struct {
	usersBaseClient
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
		usersBaseClient: usersBaseClient{
			usersImpl: usersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The Workspace Permission Assignment API allows you to manage workspace
// permissions for principals in your account.
type WorkspaceAssignmentClient struct {
	workspaceAssignmentBaseClient
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
		workspaceAssignmentBaseClient: workspaceAssignmentBaseClient{
			workspaceAssignmentImpl: workspaceAssignmentImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
