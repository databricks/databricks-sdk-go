// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package scim

import (
	"context"
)

// Groups simplify identity management, making it easier to assign access to
// Databricks Account, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks Account identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type AccountGroupsService interface {

	// Create a new group
	//
	// Creates a group in the Databricks Account with a unique name, using the
	// supplied group details.
	CreateGroup(ctx context.Context, request Group) (*Group, error)

	// Delete a group
	//
	// Deletes a group from the Databricks Account.
	DeleteGroup(ctx context.Context, request DeleteGroupRequest) error

	// DeleteGroupById calls DeleteGroup, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteGroupById(ctx context.Context, id int64) error

	// Get group details
	//
	// Gets the information for a specific group in the Databricks Account.
	GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error)

	// GetGroupById calls GetGroup, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetGroupById(ctx context.Context, id int64) (*Group, error)

	// List group details
	//
	// Gets all details of the groups associated with the Databricks Account.
	//
	// Use ListGroupsAll() to get all Group instances
	ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error)

	// ListGroupsAll calls ListGroups() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListGroupsAll(ctx context.Context, request ListGroupsRequest) ([]Group, error)

	// Update group details
	//
	// Partially updates the details of a group.
	PatchGroup(ctx context.Context, request PartialUpdate) error

	// Replace a group
	//
	// Updates the details of a group by replacing the entire group entity.
	UpdateGroup(ctx context.Context, request Group) error
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type AccountServicePrincipalsService interface {

	// Create a service principal
	//
	// Creates a new service principal in the Databricks Account.
	CreateServicePrincipal(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error)

	// Delete a service principal
	//
	// Delete a single service principal in the Databricks Account.
	DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error

	// DeleteServicePrincipalById calls DeleteServicePrincipal, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteServicePrincipalById(ctx context.Context, id int64) error

	// Get service principal details
	//
	// Gets the details for a single service principal define in the Databricks
	// Account.
	GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error)

	// GetServicePrincipalById calls GetServicePrincipal, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetServicePrincipalById(ctx context.Context, id int64) (*ServicePrincipal, error)

	// List service principals
	//
	// Gets the set of service principals associated with a Databricks Account.
	//
	// Use ListServicePrincipalsAll() to get all ServicePrincipal instances
	ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error)

	// ListServicePrincipalsAll calls ListServicePrincipals() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListServicePrincipalsAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error)

	// Update service principal details
	//
	// Partially updates the details of a single service principal in the
	// Databricks Account.
	PatchServicePrincipal(ctx context.Context, request PartialUpdate) error

	// Replace service principal
	//
	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	UpdateServicePrincipal(ctx context.Context, request ServicePrincipal) error
}

// User identities recognized by Databricks and represented by email addresses.
//
// Databricks recommends using SCIM provisioning to sync users and groups
// automatically from your identity provider to your Databricks Account. SCIM
// streamlines onboarding a new employee or team by using your identity provider
// to create users and groups in Databricks Account and give them the proper
// level of access. When a user leaves your organization or no longer needs
// access to Databricks Account, admins can terminate the user in your identity
// provider and that user’s account will also be removed from Databricks
// Account. This ensures a consistent offboarding process and prevents
// unauthorized users from accessing sensitive data.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type AccountUsersService interface {

	// Create a new user
	//
	// Creates a new user in the Databricks Account. This new user will also be
	// added to the Databricks account.
	CreateUser(ctx context.Context, request User) (*User, error)

	// Delete a user
	//
	// Deletes a user. Deleting a user from a Databricks Account also removes
	// objects associated with the user.
	DeleteUser(ctx context.Context, request DeleteUserRequest) error

	// DeleteUserById calls DeleteUser, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteUserById(ctx context.Context, id int64) error

	// Get user details
	//
	// Gets information for a specific user in Databricks Account.
	GetUser(ctx context.Context, request GetUserRequest) (*User, error)

	// GetUserById calls GetUser, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetUserById(ctx context.Context, id int64) (*User, error)

	// List users
	//
	// Gets details for all the users associated with a Databricks Account.
	//
	// Use ListUsersAll() to get all User instances
	ListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error)

	// ListUsersAll calls ListUsers() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListUsersAll(ctx context.Context, request ListUsersRequest) ([]User, error)

	// Update user details
	//
	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes.
	PatchUser(ctx context.Context, request PartialUpdate) error

	// Replace a user
	//
	// Replaces a user's information with the data supplied in request.
	UpdateUser(ctx context.Context, request User) error
}

// This API allows retrieving information about currently authenticated user or
// service principal.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type CurrentUserService interface {

	// Get current user info
	//
	// Get details about the current method caller's identity.
	Me(ctx context.Context) (*User, error)
}

// Groups simplify identity management, making it easier to assign access to
// Databricks Workspace, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks Workspace identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type GroupsService interface {

	// Create a new group
	//
	// Creates a group in the Databricks Workspace with a unique name, using the
	// supplied group details.
	CreateGroup(ctx context.Context, request Group) (*Group, error)

	// Delete a group
	//
	// Deletes a group from the Databricks Workspace.
	DeleteGroup(ctx context.Context, request DeleteGroupRequest) error

	// DeleteGroupById calls DeleteGroup, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteGroupById(ctx context.Context, id int64) error

	// Get group details
	//
	// Gets the information for a specific group in the Databricks Workspace.
	GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error)

	// GetGroupById calls GetGroup, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetGroupById(ctx context.Context, id int64) (*Group, error)

	// List group details
	//
	// Gets all details of the groups associated with the Databricks Workspace.
	//
	// Use ListGroupsAll() to get all Group instances
	ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error)

	// ListGroupsAll calls ListGroups() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListGroupsAll(ctx context.Context, request ListGroupsRequest) ([]Group, error)

	// Update group details
	//
	// Partially updates the details of a group.
	PatchGroup(ctx context.Context, request PartialUpdate) error

	// Replace a group
	//
	// Updates the details of a group by replacing the entire group entity.
	UpdateGroup(ctx context.Context, request Group) error
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type ServicePrincipalsService interface {

	// Create a service principal
	//
	// Creates a new service principal in the Databricks Workspace.
	CreateServicePrincipal(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error)

	// Delete a service principal
	//
	// Delete a single service principal in the Databricks Workspace.
	DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error

	// DeleteServicePrincipalById calls DeleteServicePrincipal, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteServicePrincipalById(ctx context.Context, id int64) error

	// Get service principal details
	//
	// Gets the details for a single service principal define in the Databricks
	// Workspace.
	GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error)

	// GetServicePrincipalById calls GetServicePrincipal, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetServicePrincipalById(ctx context.Context, id int64) (*ServicePrincipal, error)

	// List service principals
	//
	// Gets the set of service principals associated with a Databricks
	// Workspace.
	//
	// Use ListServicePrincipalsAll() to get all ServicePrincipal instances
	ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error)

	// ListServicePrincipalsAll calls ListServicePrincipals() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListServicePrincipalsAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error)

	// Update service principal details
	//
	// Partially updates the details of a single service principal in the
	// Databricks Workspace.
	PatchServicePrincipal(ctx context.Context, request PartialUpdate) error

	// Replace service principal
	//
	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	UpdateServicePrincipal(ctx context.Context, request ServicePrincipal) error
}

// User identities recognized by Databricks and represented by email addresses.
//
// Databricks recommends using SCIM provisioning to sync users and groups
// automatically from your identity provider to your Databricks Workspace. SCIM
// streamlines onboarding a new employee or team by using your identity provider
// to create users and groups in Databricks Workspace and give them the proper
// level of access. When a user leaves your organization or no longer needs
// access to Databricks Workspace, admins can terminate the user in your
// identity provider and that user’s account will also be removed from
// Databricks Workspace. This ensures a consistent offboarding process and
// prevents unauthorized users from accessing sensitive data.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type UsersService interface {

	// Create a new user
	//
	// Creates a new user in the Databricks Workspace. This new user will also
	// be added to the Databricks account.
	CreateUser(ctx context.Context, request User) (*User, error)

	// Delete a user
	//
	// Deletes a user. Deleting a user from a Databricks Workspace also removes
	// objects associated with the user.
	DeleteUser(ctx context.Context, request DeleteUserRequest) error

	// DeleteUserById calls DeleteUser, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteUserById(ctx context.Context, id int64) error

	// Get user details
	//
	// Gets information for a specific user in Databricks Workspace.
	GetUser(ctx context.Context, request GetUserRequest) (*User, error)

	// GetUserById calls GetUser, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetUserById(ctx context.Context, id int64) (*User, error)

	// List users
	//
	// Gets details for all the users associated with a Databricks Workspace.
	//
	// Use ListUsersAll() to get all User instances
	ListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error)

	// ListUsersAll calls ListUsers() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListUsersAll(ctx context.Context, request ListUsersRequest) ([]User, error)

	// Update user details
	//
	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes.
	PatchUser(ctx context.Context, request PartialUpdate) error

	// Replace a user
	//
	// Replaces a user's information with the data supplied in request.
	UpdateUser(ctx context.Context, request User) error
}
