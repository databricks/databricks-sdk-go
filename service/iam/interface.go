// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"context"
)

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set.
type AccountAccessControlService interface {

	// List assignable roles for a resource.
	//
	// Gets all the roles that can be granted on an account level resource. A
	// role is grantable if the rule set on the resource can contain an access
	// rule of the role.
	GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error)

	// Get a rule set.
	//
	// Get a rule set by its name. A rule set is always attached to a resource
	// and contains a list of access rules on the said resource. Currently only
	// a default rule set for each resource is supported.
	GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error)

	// Update a rule set.
	//
	// Replace the rules of a rule set. First, use get to read the current
	// version of the rule set before modifying it. This pattern helps prevent
	// conflicts between concurrent updates.
	UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error)
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set. A
// workspace must belong to an account for these APIs to work.
type AccountAccessControlShardProxyService interface {

	// List assignable roles for a resource.
	//
	// Gets all the roles that can be granted on an account-level resource. A
	// role is grantable if the rule set on the resource can contain an access
	// rule of the role.
	GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error)

	// Get a rule set.
	//
	// Get a rule set by its name. A rule set is always attached to a resource
	// and contains a list of access rules on the said resource. Currently only
	// a default rule set for each resource is supported.
	GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error)

	// Update a rule set.
	//
	// Replace the rules of a rule set. First, use a GET rule set request to
	// read the current version of the rule set before modifying it. This
	// pattern helps prevent conflicts between concurrent updates.
	UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error)
}

// Groups simplify identity management, making it easier to assign access to
// Databricks account, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks account identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
type AccountGroupsService interface {

	// Create a new group.
	//
	// Creates a group in the Databricks account with a unique name, using the
	// supplied group details.
	Create(ctx context.Context, request Group) (*Group, error)

	// Delete a group.
	//
	// Deletes a group from the Databricks account.
	Delete(ctx context.Context, request DeleteAccountGroupRequest) error

	// Get group details.
	//
	// Gets the information for a specific group in the Databricks account.
	Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error)

	// List group details.
	//
	// Gets all details of the groups associated with the Databricks account.
	//
	// Use ListAll() to get all Group instances
	List(ctx context.Context, request ListAccountGroupsRequest) (*ListGroupsResponse, error)

	// Update group details.
	//
	// Partially updates the details of a group.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace a group.
	//
	// Updates the details of a group by replacing the entire group entity.
	Update(ctx context.Context, request Group) error
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
type AccountServicePrincipalsService interface {

	// Create a service principal.
	//
	// Creates a new service principal in the Databricks account.
	Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error)

	// Delete a service principal.
	//
	// Delete a single service principal in the Databricks account.
	Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error

	// Get service principal details.
	//
	// Gets the details for a single service principal define in the Databricks
	// account.
	Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error)

	// List service principals.
	//
	// Gets the set of service principals associated with a Databricks account.
	//
	// Use ListAll() to get all ServicePrincipal instances
	List(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error)

	// Update service principal details.
	//
	// Partially updates the details of a single service principal in the
	// Databricks account.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace service principal.
	//
	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	Update(ctx context.Context, request ServicePrincipal) error
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
type AccountUsersService interface {

	// Create a new user.
	//
	// Creates a new user in the Databricks account. This new user will also be
	// added to the Databricks account.
	Create(ctx context.Context, request User) (*User, error)

	// Delete a user.
	//
	// Deletes a user. Deleting a user from a Databricks account also removes
	// objects associated with the user.
	Delete(ctx context.Context, request DeleteAccountUserRequest) error

	// Get user details.
	//
	// Gets information for a specific user in Databricks account.
	Get(ctx context.Context, request GetAccountUserRequest) (*User, error)

	// List users.
	//
	// Gets details for all the users associated with a Databricks account.
	//
	// Use ListAll() to get all User instances
	List(ctx context.Context, request ListAccountUsersRequest) (*ListUsersResponse, error)

	// Update user details.
	//
	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace a user.
	//
	// Replaces a user's information with the data supplied in request.
	Update(ctx context.Context, request User) error
}

// This API allows retrieving information about currently authenticated user or
// service principal.
type CurrentUserService interface {

	// Get current user info.
	//
	// Get details about the current method caller's identity.
	Me(ctx context.Context) (*User, error)
}

// Groups simplify identity management, making it easier to assign access to
// Databricks workspace, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks workspace identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
type GroupsService interface {

	// Create a new group.
	//
	// Creates a group in the Databricks workspace with a unique name, using the
	// supplied group details.
	Create(ctx context.Context, request Group) (*Group, error)

	// Delete a group.
	//
	// Deletes a group from the Databricks workspace.
	Delete(ctx context.Context, request DeleteGroupRequest) error

	// Get group details.
	//
	// Gets the information for a specific group in the Databricks workspace.
	Get(ctx context.Context, request GetGroupRequest) (*Group, error)

	// List group details.
	//
	// Gets all details of the groups associated with the Databricks workspace.
	//
	// Use ListAll() to get all Group instances
	List(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error)

	// Update group details.
	//
	// Partially updates the details of a group.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace a group.
	//
	// Updates the details of a group by replacing the entire group entity.
	Update(ctx context.Context, request Group) error
}

// This endpoint enables workspace admins to configure permissions on cluster
// policies and check which permission levels can be set.
//
// There are two permission levels for a cluster policy:
//
// * No Permissions
//
// * Can Use (`CAN_USE`).
type PermissionsClusterPoliciesService interface {

	// Get cluster policy permissions.
	//
	// Get permissions for a specific cluster policy.
	Get(ctx context.Context, request GetPermissionsClusterPolicyRequest) (*ClusterPolicyObjectPermissions, error)

	// Get permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// cluster policies.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*ClusterPolicyGetPermissionLevelsResponse, error)

	// Set cluster policy permissions.
	//
	// Update all permissions for a specific cluster policy, specifying all
	// users, groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the cluster policy and replaces it with the new
	// permissions specified in the request body.
	Set(ctx context.Context, request ClusterPolicyPermissionsRequest) error

	// Update cluster policy permissions.
	//
	// Update permissions for a specific cluster policy.
	Update(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyObjectPermissions, error)
}

// This endpoint enables users to configure permissions on clusters and check
// which permission levels can be set.
//
// There are four permission levels for a cluster:
//
// * No Permissions
//
// * Can Attach To (`CAN_ATTACH_TO`)
//
// * Can Restart (`CAN_RESTART`)
//
// * Can Manage (`CAN_MANAGE`)
//
// For the mapping of the required permissions for specific actions or
// abilities, see [Cluster access control].
//
// If the cluster is created from a job, the cluster inherits permissions from
// the job.
//
// [Cluster access control]: https://docs.databricks.com/security/access-control/cluster-acl.html
type PermissionsClustersService interface {

	// Get cluster permissions.
	//
	// Get permissions for a specific cluster.
	Get(ctx context.Context, request GetPermissionsClusterRequest) (*ClusterObjectPermissions, error)

	// Get cluster permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// clusters.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*ClusterGetPermissionLevelsResponse, error)

	// Set cluster permissions.
	//
	// Update all permissions for a specific cluster, specifying all users,
	// groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the cluster and replaces it with the new permissions
	// specified in the request body.
	Set(ctx context.Context, request ClusterPermissionsRequest) (*ClusterObjectPermissions, error)

	// Update cluster permissions.
	//
	// Update permissions for a specific cluster.
	Update(ctx context.Context, request ClusterPermissionsRequest) (*ClusterObjectPermissions, error)
}

// This endpoint enables users to configure permissions on Delta Live
// Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html)
// pipelines and check which permission levels can be set.
//
// There are five permission levels for pipelines:
//
// * No Permissions
//
// * Can View (`CAN_VIEW`) — User can view this pipeline.
//
// * Can Run (`CAN_RUN`) — User can run this pipeline.
//
// * Can Manage (`CAN_MANAGE`) — User can manage this pipeline. Workspace
// admins are granted the Can Manage permission by default.
//
// * Is Owner (`IS_OWNER`) — User is the owner of this pipeline. Only
// workspace admins can change this permission. Only one user or one service
// principal can be granted `IS_OWNER` permission on a pipeline at a given time,
// and this permission cannot be granted to groups. The API enforces these
// rules.
//
// For the mapping of the required permissions for specific actions or
// abilities, see [Delta Live Tables access control].
//
// [Delta Live Tables access control]: https://docs.databricks.com/security/access-control/dlt-acl.html
type PermissionsDeltaLiveTablesService interface {

	// Get pipeline permissions.
	//
	// Get permissions for a specific pipeline.
	Get(ctx context.Context, request GetPermissionsDeltaLiveTableRequest) (*DeltaLiveTableObjectPermissions, error)

	// Get pipeline permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// pipelines.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*DeltaLiveTableGetPermissionLevelsResponse, error)

	// Set pipeline permissions.
	//
	// Update all permissions for a specific pipeline, specifying all users,
	// groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the pipeline and replaces it with the new permissions
	// specified in the request body.
	Set(ctx context.Context, request DeltaLiveTablePermissionsRequest) (*DeltaLiveTableObjectPermissions, error)

	// Update pipeline permissions.
	//
	// Update permissions for a specific pipeline.
	Update(ctx context.Context, request DeltaLiveTablePermissionsRequest) (*DeltaLiveTableObjectPermissions, error)
}

// This endpoint enables users to configure permissions on directories and check
// which permission levels can be set. Note that in the web application user
// interface and in some other documentation, directories are referred to as
// _folders_.
//
// There are five permission levels for directories:
//
// * No Permissions
//
// * Can Read (`CAN_READ`) — User can read items this directory
//
// * Can Run (`CAN_RUN`) — Can run items in this directory.
//
// * Can Edit (`CAN_EDIT`) — Can edit items in this directory.
//
// * Can Manage (`CAN_MANAGE`) — Can manage this directory.
//
// **IMPORTANT:** Notebooks and experiments in a folder inherit all permissions
// settings of that folder. For example, a user that has Run permission on a
// folder has Run permission on the notebooks in that folder.
//
// For the mapping of the required permissions for specific actions or
// abilities, see [Folder permissions].
//
// Permissions on directories are inherited from its parent directories. You can
// set permissions directly on the root directory, which has no parent, so the
// root directory never inherits permissions.
//
// [Folder permissions]: https://docs.databricks.com/security/access-control/workspace-acl.html#folder-permissions
type PermissionsDirectoriesService interface {

	// Get directory permissions.
	//
	// Get permissions for a specific directory.
	Get(ctx context.Context, request GetPermissionsDirectoryRequest) (*DirectoryObjectPermissions, error)

	// Get directory permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// directories.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*DirectoryGetPermissionLevelsResponse, error)

	// Set directory permissions.
	//
	// Update all permissions for a specific directory, specifying all users,
	// groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the directory and replaces it with the new permissions
	// specified in the request body.
	Set(ctx context.Context, request DirectoryPermissionsRequest) (*DirectoryObjectPermissions, error)

	// Update directory permissions.
	//
	// Update permissions for a specific directory.
	Update(ctx context.Context, request DirectoryPermissionsRequest) (*DirectoryObjectPermissions, error)
}

// This endpoint enables users to configure permissions on jobs and check which
// permission levels can be set.
//
// There are five permission levels for jobs:
//
// * No Permissions
//
// * Can View (`CAN_VIEW`) — User can view this job
//
// * Can Manage Run (`CAN_MANAGE_RUN`) — User can manage or run this job.
//
// * Is Owner (`IS_OWNER`) — User is the owner of this job.
//
// * Can Manage (`CAN_MANAGE`) — User can manage this job. Workspace admins
// are granted the Can Manage permission by default and can assign that
// permission to non-admin users.
//
// For the mapping of the required permissions for specific actions or
// abilities, see [Jobs access control].
//
// [Jobs access control]: https://docs.databricks.com/security/access-control/jobs-acl.html
type PermissionsJobsService interface {

	// Get job permissions.
	//
	// Get permissions for a specific job.
	Get(ctx context.Context, request GetPermissionsJobRequest) (*JobObjectPermissions, error)

	// Get job permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// jobs.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*JobGetPermissionLevelsResponse, error)

	// Set job permissions.
	//
	// Update all permissions for a specific job, specifying all users, groups,
	// or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the job and replaces it with the new permissions specified
	// in the request body.
	Set(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error)

	// Update job permissions.
	//
	// Update permissions for a specific job.
	Update(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error)
}

// This endpoint enables users to configure permissions on MLflow experiments
// and check which permission levels can be set.
//
// You can assign four permission levels to experiments:
//
// * No Permissions
//
// * Can Read (`CAN_READ`)
//
// * Can Edit (`CAN_EDIT`)
//
// * Can Manage (`CAN_MANAGE`)
//
// For the mapping of the required permissions for specific actions or
// abilities, see [MLflow experiment permissions].
//
// For auto-generated experiments (for example, when a user runs a notebook
// without calling `mlflow.set_experiment()` explicitly), permissions can only
// be changed by using notebook permissions.
//
// For more information, see [notebook experiments].
//
// [MLflow experiment permissions]: https://docs.databricks.com/security/access-control/workspace-acl.html#mlflow-experiment-permissions-1
// [notebook experiments]: https://docs.databricks.com/applications/mlflow/tracking.html#notebook-experiments
type PermissionsMLflowExperimentsService interface {

	// Get experiment permissions.
	//
	// Get permissions for a specific experiment.
	Get(ctx context.Context, request GetPermissionsMLflowExperimentRequest) (*MLflowExperimentObjectPermissions, error)

	// Get experiment permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// MLflow experiments.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*MLflowExperimentGetPermissionLevelsResponse, error)

	// Set experiment permissions.
	//
	// Update all permissions for a specific experiment, specifying all users,
	// groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the experiment and replaces it with the new permissions
	// specified in the request body.
	Set(ctx context.Context, request MLflowExperimentPermissionsRequest) (*MLflowExperimentObjectPermissions, error)

	// Update experiment permissions.
	//
	// Update permissions for a specific experiment.
	Update(ctx context.Context, request MLflowExperimentPermissionsRequest) (*MLflowExperimentObjectPermissions, error)
}

// This endpoint enables users to configure permissions on MLflow models and
// check which permission levels can be set.
//
// You can assign six permission levels to registered models:
//
// * No Permissions
//
// * Can Read (`CAN_READ`)
//
// * Can Edit (`CAN_EDIT`)
//
// * Can Manage Staging Versions (`CAN_MANAGE_STAGING_VERSIONS`)
//
// * Can Manage Production Versions (`CAN_MANAGE_PRODUCTION_VERSIONS`)
//
// * Can Manage (`CAN_MANAGE`).
//
// For the mapping of the required permissions for specific actions or
// abilities, see [MLflow model permissions].
//
// MLflow registered models inherit permissions from their root object.
//
// [MLflow model permissions]: https://docs.databricks.com/security/access-control/workspace-acl.html#mlflow-model-permissions
type PermissionsMlFlowRegisteredModelsService interface {

	// Get registered model permissions.
	//
	// Get permissions for a specific registered model.
	Get(ctx context.Context, request GetPermissionsMlFlowRegisteredModelRequest) (*MLflowRegisteredModelObjectPermissions, error)

	// Get registered model permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// MLflow models.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*MLflowRegisteredModelGetPermissionLevelsResponse, error)

	// Set registered model permissions.
	//
	// Update all permissions for a specific registered model, specifying all
	// users, groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the registered model and replaces it with the new
	// permissions specified in the request body.
	Set(ctx context.Context, request MLflowRegisteredModelPermissionsRequest) (*MLflowRegisteredModelObjectPermissions, error)

	// Update registered model permissions.
	//
	// Update permissions for a specific registered model.
	Update(ctx context.Context, request MLflowRegisteredModelPermissionsRequest) (*MLflowRegisteredModelObjectPermissions, error)
}

// This endpoint enables users to configure permissions on notebooks and check
// which permission levels can be set.
//
// There are five permission levels for notebooks:
//
// * No Permissions
//
// * Can Read (`CAN_READ`) — User can read this notebook. The user can also
// run the notebook via %run or notebook workflows.
//
// * Can Run (`CAN_RUN`) — User can run this notebook. The user can not only
// use %run and notebook workflows, but also run commands and can attach or
// detach notebooks.
//
// * Can Edit (`CAN_EDIT`) — User can edit this notebook.
//
// * Can Manage (`CAN_MANAGE`) — User can manage this job. Workspace admins
// are granted the Can Manage permission by default and can assign that
// permission to non-admin users.
//
// For the mapping of the required permissions for specific actions or
// abilities, see [Notebook access control].
//
// Notebooks inherit permissions from their root object. Additionally, notebooks
// inherit permissions from parent directories, similar to how directories
// inherit permissions from their parent directories. For example, a notebook
// with path `/Users/jsmith@example.com/myNotebook` can inherit permissions from
// **all** of the following objects: `/` (root directory), `/Users`, and
// `/Users/jsmith@example.com`.
//
// [Notebook access control]: https://docs.databricks.com/security/access-control/workspace-acl.html#notebook-permissions
type PermissionsNotebooksService interface {

	// Get notebook permissions.
	//
	// Get permissions for a specific notebook.
	Get(ctx context.Context, request GetPermissionsNotebookRequest) (*JobObjectPermissions, error)

	// Get notebook permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// notebooks.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*JobGetPermissionLevelsResponse, error)

	// Set notebook permissions.
	//
	// Update all permissions for a specific notebook, specifying all users,
	// groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the notebook and replaces it with the new permissions
	// specified in the request body.
	Set(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error)

	// Update notebook permissions.
	//
	// Update permissions for a specific notebook.
	Update(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error)
}

// This endpoint enables admins to configure permissions on passwords and check
// which permission levels can be set. Password permissions control which users
// can use native authentication with username/password when single sign-on
// (SSO) is enabled for the workspace and Unified Login is disabled. The only
// supported permission to add is `CAN_USE`, which specifies login is permitted
// even if SSO is enabled.
//
// **IMPORTANT:** If SSO is not enabled, these permissions have no effect. All
// users with locally-stored (native authentication) passwords can sign in with
// the web application and authenticate with the REST API.
//
// By default, the built-in group for all users (`users`) has this permission.
// Workspace admins can replace all permissions for the workspace to remove that
// group permission and enable the permission for the `admins` group or only for
// specific admin users.
//
// It is important to understand the following authentication differences:
//
// * **From the web application user interface**, if SSO is enabled, there are
// two tabs for login. The non-SSO login is for admin users only. You can login
// with a native authentication password only if all of the following are true:
// (a) You are in the `admins` group. (b) You have the password permission
// `CAN_USE`. (c) You have a locally-stored (native authentication) password,
// independent of whether your user was originally created using SSO/SCIM.
// Although uncommon, users created using SSO/SCIM can create a native
// authentication password using the password recovery user interface.
//
// * **From the REST API**, if SSO is enabled, you can login only if you have
// the password permission `CAN_USE`. You do not need to be in the `admins`
// group.
type PermissionsPasswordsService interface {

	// Get password permissions.
	//
	// Get password permissions for the entire workspace.
	Get(ctx context.Context) (*PasswordObjectPermissions, error)

	// Get password permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// passwords.
	GetPermissionLevels(ctx context.Context) (*PasswordGetPermissionLevelsResponse, error)

	// Set password permissions.
	//
	// Update all password permissions for the entire workspace, specifying all
	// users, groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing permissions and
	// replaces them with the new permissions specified in the request body.
	Set(ctx context.Context, request PasswordPermissionsRequest) (*PasswordObjectPermissions, error)

	// Update password permissions.
	//
	// Update password permissions for the entire workspace.
	Update(ctx context.Context, request PasswordPermissionsRequest) (*PasswordObjectPermissions, error)
}

// This endpoint enables users to configure permissions on pools and check which
// permission levels can be set.
//
// There are three permission levels for a pool:
//
// * No Permissions
//
// * Can Attach To (`CAN_ATTACH_TO`)
//
// * Can Manage (`CAN_MANAGE`)
//
// For the mapping of the required permissions for specific actions or
// abilities, see [Pool access control].
//
// [Pool access control]: https://docs.databricks.com/security/access-control/pool-acl.html
type PermissionsPoolsService interface {

	// Get instance pool permissions.
	//
	// Get permissions for a specific instance pool.
	Get(ctx context.Context, request GetPermissionsPoolRequest) (*PoolObjectPermissions, error)

	// Get instance pool permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// pools.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*PoolGetPermissionLevelsResponse, error)

	// Set instance pool permissions.
	//
	// Update all permissions for a specific instance pool, specifying all
	// users, groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the instance pool and replaces it with the new permissions
	// specified in the request body.
	Set(ctx context.Context, request PoolPermissionsRequest) (*PoolObjectPermissions, error)

	// Update instance pool permissions.
	//
	// Update permissions for a specific instance pool.
	Update(ctx context.Context, request PoolPermissionsRequest) (*PoolObjectPermissions, error)
}

// This endpoint enables users to configure permissions on repos and check which
// permission levels can be set.
//
// There are five permission levels for repos:
//
// * No Permissions
//
// * Can Read (`CAN_READ`) — Can read items in this repo.
//
// * Can Run (`CAN_RUN`) — Can run items in this repo.
//
// * Can Edit (`CAN_EDIT`) — Can edit items in this repo.
//
// * Can Manage (`CAN_MANAGE`) — Can manage this repo.
type PermissionsReposService interface {

	// Get repo permissions.
	//
	// Get permissions for a specific repo.
	Get(ctx context.Context, request GetPermissionsRepoRequest) (*RepoObjectPermissions, error)

	// Get repo permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// repos.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*RepoGetPermissionLevelsResponse, error)

	// Set repo permissions.
	//
	// Update all permissions for a specific repo, specifying all users, groups,
	// or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the repo and replaces it with the new permissions
	// specified in the request body.
	Set(ctx context.Context, request RepoPermissionsRequest) (*RepoObjectPermissions, error)

	// Update repo permissions.
	//
	// Update permissions for a specific repo.
	Update(ctx context.Context, request RepoPermissionsRequest) (*RepoObjectPermissions, error)
}

// This endpoint enables users to configure permissions on SQL warehouses and
// check which permission levels can be set.
//
// You can assign four permission levels to SQL warehouses:
//
// * No Permissions
//
// * Can Use (`CAN_USE`)
//
// * Can Manage (`CAN_MANAGE`)
//
// * Is Owner (`IS_OWNER`)
//
// For the mapping of the required permissions for specific actions or
// abilities, see [SQL warehouse permissions].
//
// [SQL warehouse permissions]: https://docs.databricks.com/sql/user/security/access-control/sql-endpoint-acl.html#sql-warehouse-permissions
type PermissionsSqlWarehousesService interface {

	// Get SQL warehouse permissions.
	//
	// Get permissions for a specific SQL warehouse.
	Get(ctx context.Context, request GetPermissionsSqlWarehousRequest) (*SqlWarehouseObjectPermissions, error)

	// Get SQL warehouse permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for SQL
	// warehouses.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*SqlWarehouseGetPermissionLevelsResponse, error)

	// Set SQL warehouse permissions.
	//
	// Update all permissions for a specific SQL warehouse, specifying all
	// users, groups, or service principals.
	//
	// **WARNING:** This request overwrites all existing direct (non-inherited)
	// permissions on the SQL warehouse and replaces it with the new permissions
	// specified in the request body.
	Set(ctx context.Context, request SqlWarehousePermissionsRequest) (*SqlWarehouseObjectPermissions, error)

	// Update SQL warehouse permissions.
	//
	// Update permissions for a specific SQL warehouse.
	Update(ctx context.Context, request SqlWarehousePermissionsRequest) (*SqlWarehouseObjectPermissions, error)
}

// This endpoint enables admins to configure permissions on tokens and check
// which permission levels can be set.
//
// There are several levels of token permissions that a user can have:
//
// * No Permissions
//
// * Can Use (`CAN_USE`) — The default is for no users to have the Can Use
// permission. Admins must explicitly grant those permissions, whether to the
// entire `users` group or on a user-by-user or group-by-group basis.
//
// * Can Manage (`CAN_MANAGE`) — Workspace admins only. The `admins` group
// gets this permission and it cannot be changed. No other groups or users can
// be granted this permission. The API enforces these rules.
//
// **WARNING:** If you revoke the Can Use permission from a group and a user
// does not have the Can Use permission directly or indirectly through another
// group, that user’s tokens are immediately deleted. Deleted tokens cannot be
// retrieved.
type PermissionsTokensService interface {

	// Get token permissions.
	//
	// Get token permissions for the entire workspace.
	Get(ctx context.Context) (*TokenObjectPermissions, error)

	// Get token permission levels.
	//
	// Returns a JSON representation of the possible permissions levels for
	// tokens.
	GetPermissionLevels(ctx context.Context) (*TokenGetPermissionLevelsResponse, error)

	// Set token permissions.
	//
	// Update all token permissions for all users, groups, and service
	// principals for the entire workspace.
	//
	// At the end of processing your request, all users and service principals
	// that do not have either `CAN_USE` or `CAN_MANAGE` permission either
	// explicitly or implicitly due to group assignment no longer have any
	// tokens permissions. Affected users or service principals immediately have
	// all their tokens deleted.
	//
	// **WARNING:** This request has powerful effects for workspace security
	// configuration and on a workspace's users if they already use tokens. Use
	// with caution. This request overwrites all existing token permissions with
	// the data in the request body. By omitting reference to an entity that
	// previously had permissions, access is stripped and existing tokens are
	// permanently deleted.
	Set(ctx context.Context, request TokenPermissionsRequest) (*TokenObjectPermissions, error)

	// Update token permissions.
	//
	// Grant token permissions for one or more users, groups, or service
	// principals. You can only grant the Can Use (`CAN_USE`) permission. The
	// Can Manage (`CAN_MANAGE`) permission level cannot be granted with this
	// API because it is tied automatically to membership in the `admins` group.
	Update(ctx context.Context, request TokenPermissionsRequest) (*TokenObjectPermissions, error)
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
type ServicePrincipalsService interface {

	// Create a service principal.
	//
	// Creates a new service principal in the Databricks workspace.
	Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error)

	// Delete a service principal.
	//
	// Delete a single service principal in the Databricks workspace.
	Delete(ctx context.Context, request DeleteServicePrincipalRequest) error

	// Get service principal details.
	//
	// Gets the details for a single service principal define in the Databricks
	// workspace.
	Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error)

	// List service principals.
	//
	// Gets the set of service principals associated with a Databricks
	// workspace.
	//
	// Use ListAll() to get all ServicePrincipal instances
	List(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error)

	// Update service principal details.
	//
	// Partially updates the details of a single service principal in the
	// Databricks workspace.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace service principal.
	//
	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	Update(ctx context.Context, request ServicePrincipal) error
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
type UsersService interface {

	// Create a new user.
	//
	// Creates a new user in the Databricks workspace. This new user will also
	// be added to the Databricks account.
	Create(ctx context.Context, request User) (*User, error)

	// Delete a user.
	//
	// Deletes a user. Deleting a user from a Databricks workspace also removes
	// objects associated with the user.
	Delete(ctx context.Context, request DeleteUserRequest) error

	// Get user details.
	//
	// Gets information for a specific user in Databricks workspace.
	Get(ctx context.Context, request GetUserRequest) (*User, error)

	// List users.
	//
	// Gets details for all the users associated with a Databricks workspace.
	//
	// Use ListAll() to get all User instances
	List(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error)

	// Update user details.
	//
	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace a user.
	//
	// Replaces a user's information with the data supplied in request.
	Update(ctx context.Context, request User) error
}

// The Workspace Permission Assignment API allows you to manage workspace
// permissions for principals in your account.
type WorkspaceAssignmentService interface {

	// Delete permissions assignment.
	//
	// Deletes the workspace permissions assignment in a given account and
	// workspace for the specified principal.
	Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error

	// List workspace permissions.
	//
	// Get an array of workspace permissions for the specified account and
	// workspace.
	Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error)

	// Get permission assignments.
	//
	// Get the permission assignments for the specified Databricks account and
	// Databricks workspace.
	//
	// Use ListAll() to get all PermissionAssignment instances
	List(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error)

	// Create or update permissions assignment.
	//
	// Creates or updates the workspace permissions assignment in a given
	// account and workspace for the specified principal.
	Update(ctx context.Context, request UpdateWorkspaceAssignments) error
}
