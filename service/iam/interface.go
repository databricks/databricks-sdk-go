// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"context"
)

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set.
type AccountAccessControlService interface {

	// Get assignable roles for a resource.
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
type AccountAccessControlProxyService interface {

	// Get assignable roles for a resource.
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

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints.
type PermissionsService interface {

	// Get object permissions.
	//
	// Gets the permission of an object. Objects can inherit permissions from
	// their parent objects or root objects.
	Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error)

	// Get permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error)

	// Set permissions.
	//
	// Sets permissions on object. Objects can inherit permissions from their
	// parent objects and root objects.
	Set(ctx context.Context, request PermissionsRequest) error

	// Update permission.
	//
	// Updates the permissions on an object.
	Update(ctx context.Context, request PermissionsRequest) error
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
