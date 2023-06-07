// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Access Control, Account Access Control Shard Proxy, Account Groups, Account Service Principals, Account Users, Current User, Groups, Permissions Cluster Policies, Permissions Clusters, Permissions Delta Live Tables, Permissions Directories, Permissions Jobs, Permissions M Lflow Experiments, Permissions Ml Flow Registered Models, Permissions Notebooks, Permissions Passwords, Permissions Pools, Permissions Repos, Permissions Sql Warehouses, Permissions Tokens, Service Principals, Users, Workspace Assignment, etc.
package iam

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewAccountAccessControl(client *client.DatabricksClient) *AccountAccessControlAPI {
	return &AccountAccessControlAPI{
		impl: &accountAccessControlImpl{
			client: client,
		},
	}
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set.
type AccountAccessControlAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountAccessControlService)
	impl AccountAccessControlService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountAccessControlAPI) WithImpl(impl AccountAccessControlService) *AccountAccessControlAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountAccessControl API implementation
func (a *AccountAccessControlAPI) Impl() AccountAccessControlService {
	return a.impl
}

// List assignable roles for a resource.
//
// Gets all the roles that can be granted on an account level resource. A role
// is grantable if the rule set on the resource can contain an access rule of
// the role.
func (a *AccountAccessControlAPI) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	return a.impl.GetAssignableRolesForResource(ctx, request)
}

// Get a rule set.
//
// Get a rule set by its name. A rule set is always attached to a resource and
// contains a list of access rules on the said resource. Currently only a
// default rule set for each resource is supported.
func (a *AccountAccessControlAPI) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	return a.impl.GetRuleSet(ctx, request)
}

// Update a rule set.
//
// Replace the rules of a rule set. First, use get to read the current version
// of the rule set before modifying it. This pattern helps prevent conflicts
// between concurrent updates.
func (a *AccountAccessControlAPI) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	return a.impl.UpdateRuleSet(ctx, request)
}

func NewAccountAccessControlShardProxy(client *client.DatabricksClient) *AccountAccessControlShardProxyAPI {
	return &AccountAccessControlShardProxyAPI{
		impl: &accountAccessControlShardProxyImpl{
			client: client,
		},
	}
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set. A
// workspace must belong to an account for these APIs to work.
type AccountAccessControlShardProxyAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountAccessControlShardProxyService)
	impl AccountAccessControlShardProxyService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountAccessControlShardProxyAPI) WithImpl(impl AccountAccessControlShardProxyService) *AccountAccessControlShardProxyAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountAccessControlShardProxy API implementation
func (a *AccountAccessControlShardProxyAPI) Impl() AccountAccessControlShardProxyService {
	return a.impl
}

// List assignable roles for a resource.
//
// Gets all the roles that can be granted on an account-level resource. A role
// is grantable if the rule set on the resource can contain an access rule of
// the role.
func (a *AccountAccessControlShardProxyAPI) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	return a.impl.GetAssignableRolesForResource(ctx, request)
}

// Get a rule set.
//
// Get a rule set by its name. A rule set is always attached to a resource and
// contains a list of access rules on the said resource. Currently only a
// default rule set for each resource is supported.
func (a *AccountAccessControlShardProxyAPI) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	return a.impl.GetRuleSet(ctx, request)
}

// Update a rule set.
//
// Replace the rules of a rule set. First, use a GET rule set request to read
// the current version of the rule set before modifying it. This pattern helps
// prevent conflicts between concurrent updates.
func (a *AccountAccessControlShardProxyAPI) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	return a.impl.UpdateRuleSet(ctx, request)
}

func NewAccountGroups(client *client.DatabricksClient) *AccountGroupsAPI {
	return &AccountGroupsAPI{
		impl: &accountGroupsImpl{
			client: client,
		},
	}
}

// Groups simplify identity management, making it easier to assign access to
// Databricks account, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks account identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
type AccountGroupsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountGroupsService)
	impl AccountGroupsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountGroupsAPI) WithImpl(impl AccountGroupsService) *AccountGroupsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountGroups API implementation
func (a *AccountGroupsAPI) Impl() AccountGroupsService {
	return a.impl
}

// Create a new group.
//
// Creates a group in the Databricks account with a unique name, using the
// supplied group details.
func (a *AccountGroupsAPI) Create(ctx context.Context, request Group) (*Group, error) {
	return a.impl.Create(ctx, request)
}

// Delete a group.
//
// Deletes a group from the Databricks account.
func (a *AccountGroupsAPI) Delete(ctx context.Context, request DeleteAccountGroupRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a group.
//
// Deletes a group from the Databricks account.
func (a *AccountGroupsAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteAccountGroupRequest{
		Id: id,
	})
}

// Get group details.
//
// Gets the information for a specific group in the Databricks account.
func (a *AccountGroupsAPI) Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error) {
	return a.impl.Get(ctx, request)
}

// Get group details.
//
// Gets the information for a specific group in the Databricks account.
func (a *AccountGroupsAPI) GetById(ctx context.Context, id string) (*Group, error) {
	return a.impl.Get(ctx, GetAccountGroupRequest{
		Id: id,
	})
}

// List group details.
//
// Gets all details of the groups associated with the Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountGroupsAPI) ListAll(ctx context.Context, request ListAccountGroupsRequest) ([]Group, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

// GroupDisplayNameToIdMap calls [AccountGroupsAPI.ListAll] and creates a map of results with [Group].DisplayName as key and [Group].Id as value.
//
// Returns an error if there's more than one [Group] with the same .DisplayName.
//
// Note: All [Group] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountGroupsAPI) GroupDisplayNameToIdMap(ctx context.Context, request ListAccountGroupsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.DisplayName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .DisplayName: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByDisplayName calls [AccountGroupsAPI.GroupDisplayNameToIdMap] and returns a single [Group].
//
// Returns an error if there's more than one [Group] with the same .DisplayName.
//
// Note: All [Group] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountGroupsAPI) GetByDisplayName(ctx context.Context, name string) (*Group, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListAccountGroupsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]Group{}
	for _, v := range result {
		key := v.DisplayName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("Group named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of Group named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update group details.
//
// Partially updates the details of a group.
func (a *AccountGroupsAPI) Patch(ctx context.Context, request PartialUpdate) error {
	return a.impl.Patch(ctx, request)
}

// Replace a group.
//
// Updates the details of a group by replacing the entire group entity.
func (a *AccountGroupsAPI) Update(ctx context.Context, request Group) error {
	return a.impl.Update(ctx, request)
}

func NewAccountServicePrincipals(client *client.DatabricksClient) *AccountServicePrincipalsAPI {
	return &AccountServicePrincipalsAPI{
		impl: &accountServicePrincipalsImpl{
			client: client,
		},
	}
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
type AccountServicePrincipalsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountServicePrincipalsService)
	impl AccountServicePrincipalsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountServicePrincipalsAPI) WithImpl(impl AccountServicePrincipalsService) *AccountServicePrincipalsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountServicePrincipals API implementation
func (a *AccountServicePrincipalsAPI) Impl() AccountServicePrincipalsService {
	return a.impl
}

// Create a service principal.
//
// Creates a new service principal in the Databricks account.
func (a *AccountServicePrincipalsAPI) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	return a.impl.Create(ctx, request)
}

// Delete a service principal.
//
// Delete a single service principal in the Databricks account.
func (a *AccountServicePrincipalsAPI) Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a service principal.
//
// Delete a single service principal in the Databricks account.
func (a *AccountServicePrincipalsAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteAccountServicePrincipalRequest{
		Id: id,
	})
}

// Get service principal details.
//
// Gets the details for a single service principal define in the Databricks
// account.
func (a *AccountServicePrincipalsAPI) Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error) {
	return a.impl.Get(ctx, request)
}

// Get service principal details.
//
// Gets the details for a single service principal define in the Databricks
// account.
func (a *AccountServicePrincipalsAPI) GetById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.impl.Get(ctx, GetAccountServicePrincipalRequest{
		Id: id,
	})
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountServicePrincipalsAPI) ListAll(ctx context.Context, request ListAccountServicePrincipalsRequest) ([]ServicePrincipal, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

// ServicePrincipalDisplayNameToIdMap calls [AccountServicePrincipalsAPI.ListAll] and creates a map of results with [ServicePrincipal].DisplayName as key and [ServicePrincipal].Id as value.
//
// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
//
// Note: All [ServicePrincipal] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountServicePrincipalsAPI) ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListAccountServicePrincipalsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.DisplayName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .DisplayName: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByDisplayName calls [AccountServicePrincipalsAPI.ServicePrincipalDisplayNameToIdMap] and returns a single [ServicePrincipal].
//
// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
//
// Note: All [ServicePrincipal] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountServicePrincipalsAPI) GetByDisplayName(ctx context.Context, name string) (*ServicePrincipal, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListAccountServicePrincipalsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]ServicePrincipal{}
	for _, v := range result {
		key := v.DisplayName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("ServicePrincipal named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of ServicePrincipal named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update service principal details.
//
// Partially updates the details of a single service principal in the Databricks
// account.
func (a *AccountServicePrincipalsAPI) Patch(ctx context.Context, request PartialUpdate) error {
	return a.impl.Patch(ctx, request)
}

// Replace service principal.
//
// Updates the details of a single service principal.
//
// This action replaces the existing service principal with the same name.
func (a *AccountServicePrincipalsAPI) Update(ctx context.Context, request ServicePrincipal) error {
	return a.impl.Update(ctx, request)
}

func NewAccountUsers(client *client.DatabricksClient) *AccountUsersAPI {
	return &AccountUsersAPI{
		impl: &accountUsersImpl{
			client: client,
		},
	}
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
type AccountUsersAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountUsersService)
	impl AccountUsersService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountUsersAPI) WithImpl(impl AccountUsersService) *AccountUsersAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountUsers API implementation
func (a *AccountUsersAPI) Impl() AccountUsersService {
	return a.impl
}

// Create a new user.
//
// Creates a new user in the Databricks account. This new user will also be
// added to the Databricks account.
func (a *AccountUsersAPI) Create(ctx context.Context, request User) (*User, error) {
	return a.impl.Create(ctx, request)
}

// Delete a user.
//
// Deletes a user. Deleting a user from a Databricks account also removes
// objects associated with the user.
func (a *AccountUsersAPI) Delete(ctx context.Context, request DeleteAccountUserRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a user.
//
// Deletes a user. Deleting a user from a Databricks account also removes
// objects associated with the user.
func (a *AccountUsersAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteAccountUserRequest{
		Id: id,
	})
}

// Get user details.
//
// Gets information for a specific user in Databricks account.
func (a *AccountUsersAPI) Get(ctx context.Context, request GetAccountUserRequest) (*User, error) {
	return a.impl.Get(ctx, request)
}

// Get user details.
//
// Gets information for a specific user in Databricks account.
func (a *AccountUsersAPI) GetById(ctx context.Context, id string) (*User, error) {
	return a.impl.Get(ctx, GetAccountUserRequest{
		Id: id,
	})
}

// List users.
//
// Gets details for all the users associated with a Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountUsersAPI) ListAll(ctx context.Context, request ListAccountUsersRequest) ([]User, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

// UserUserNameToIdMap calls [AccountUsersAPI.ListAll] and creates a map of results with [User].UserName as key and [User].Id as value.
//
// Returns an error if there's more than one [User] with the same .UserName.
//
// Note: All [User] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountUsersAPI) UserUserNameToIdMap(ctx context.Context, request ListAccountUsersRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.UserName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .UserName: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByUserName calls [AccountUsersAPI.UserUserNameToIdMap] and returns a single [User].
//
// Returns an error if there's more than one [User] with the same .UserName.
//
// Note: All [User] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountUsersAPI) GetByUserName(ctx context.Context, name string) (*User, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListAccountUsersRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]User{}
	for _, v := range result {
		key := v.UserName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("User named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of User named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update user details.
//
// Partially updates a user resource by applying the supplied operations on
// specific user attributes.
func (a *AccountUsersAPI) Patch(ctx context.Context, request PartialUpdate) error {
	return a.impl.Patch(ctx, request)
}

// Replace a user.
//
// Replaces a user's information with the data supplied in request.
func (a *AccountUsersAPI) Update(ctx context.Context, request User) error {
	return a.impl.Update(ctx, request)
}

func NewCurrentUser(client *client.DatabricksClient) *CurrentUserAPI {
	return &CurrentUserAPI{
		impl: &currentUserImpl{
			client: client,
		},
	}
}

// This API allows retrieving information about currently authenticated user or
// service principal.
type CurrentUserAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CurrentUserService)
	impl CurrentUserService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *CurrentUserAPI) WithImpl(impl CurrentUserService) *CurrentUserAPI {
	a.impl = impl
	return a
}

// Impl returns low-level CurrentUser API implementation
func (a *CurrentUserAPI) Impl() CurrentUserService {
	return a.impl
}

// Get current user info.
//
// Get details about the current method caller's identity.
func (a *CurrentUserAPI) Me(ctx context.Context) (*User, error) {
	return a.impl.Me(ctx)
}

func NewGroups(client *client.DatabricksClient) *GroupsAPI {
	return &GroupsAPI{
		impl: &groupsImpl{
			client: client,
		},
	}
}

// Groups simplify identity management, making it easier to assign access to
// Databricks workspace, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks workspace identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
type GroupsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(GroupsService)
	impl GroupsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *GroupsAPI) WithImpl(impl GroupsService) *GroupsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Groups API implementation
func (a *GroupsAPI) Impl() GroupsService {
	return a.impl
}

// Create a new group.
//
// Creates a group in the Databricks workspace with a unique name, using the
// supplied group details.
func (a *GroupsAPI) Create(ctx context.Context, request Group) (*Group, error) {
	return a.impl.Create(ctx, request)
}

// Delete a group.
//
// Deletes a group from the Databricks workspace.
func (a *GroupsAPI) Delete(ctx context.Context, request DeleteGroupRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a group.
//
// Deletes a group from the Databricks workspace.
func (a *GroupsAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteGroupRequest{
		Id: id,
	})
}

// Get group details.
//
// Gets the information for a specific group in the Databricks workspace.
func (a *GroupsAPI) Get(ctx context.Context, request GetGroupRequest) (*Group, error) {
	return a.impl.Get(ctx, request)
}

// Get group details.
//
// Gets the information for a specific group in the Databricks workspace.
func (a *GroupsAPI) GetById(ctx context.Context, id string) (*Group, error) {
	return a.impl.Get(ctx, GetGroupRequest{
		Id: id,
	})
}

// List group details.
//
// Gets all details of the groups associated with the Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GroupsAPI) ListAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

// GroupDisplayNameToIdMap calls [GroupsAPI.ListAll] and creates a map of results with [Group].DisplayName as key and [Group].Id as value.
//
// Returns an error if there's more than one [Group] with the same .DisplayName.
//
// Note: All [Group] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GroupsAPI) GroupDisplayNameToIdMap(ctx context.Context, request ListGroupsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.DisplayName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .DisplayName: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByDisplayName calls [GroupsAPI.GroupDisplayNameToIdMap] and returns a single [Group].
//
// Returns an error if there's more than one [Group] with the same .DisplayName.
//
// Note: All [Group] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GroupsAPI) GetByDisplayName(ctx context.Context, name string) (*Group, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListGroupsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]Group{}
	for _, v := range result {
		key := v.DisplayName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("Group named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of Group named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update group details.
//
// Partially updates the details of a group.
func (a *GroupsAPI) Patch(ctx context.Context, request PartialUpdate) error {
	return a.impl.Patch(ctx, request)
}

// Replace a group.
//
// Updates the details of a group by replacing the entire group entity.
func (a *GroupsAPI) Update(ctx context.Context, request Group) error {
	return a.impl.Update(ctx, request)
}

func NewPermissionsClusterPolicies(client *client.DatabricksClient) *PermissionsClusterPoliciesAPI {
	return &PermissionsClusterPoliciesAPI{
		impl: &permissionsClusterPoliciesImpl{
			client: client,
		},
	}
}

// This endpoint enables workspace admins to configure permissions on cluster
// policies and check which permission levels can be set.
//
// There are two permission levels for a cluster policy:
//
// * No Permissions
//
// * Can Use (`CAN_USE`).
type PermissionsClusterPoliciesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsClusterPoliciesService)
	impl PermissionsClusterPoliciesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsClusterPoliciesAPI) WithImpl(impl PermissionsClusterPoliciesService) *PermissionsClusterPoliciesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsClusterPolicies API implementation
func (a *PermissionsClusterPoliciesAPI) Impl() PermissionsClusterPoliciesService {
	return a.impl
}

// Get cluster policy permissions.
//
// Get permissions for a specific cluster policy.
func (a *PermissionsClusterPoliciesAPI) Get(ctx context.Context, request GetPermissionsClusterPolicyRequest) (*ClusterPolicyObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get cluster policy permissions.
//
// Get permissions for a specific cluster policy.
func (a *PermissionsClusterPoliciesAPI) GetByClusterPolicyId(ctx context.Context, clusterPolicyId string) (*ClusterPolicyObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsClusterPolicyRequest{
		ClusterPolicyId: clusterPolicyId,
	})
}

// Get permission levels.
//
// Returns a JSON representation of the possible permissions levels for cluster
// policies.
func (a *PermissionsClusterPoliciesAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*ClusterPolicyGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get permission levels.
//
// Returns a JSON representation of the possible permissions levels for cluster
// policies.
func (a *PermissionsClusterPoliciesAPI) GetPermissionLevelsByClusterPolicyId(ctx context.Context, clusterPolicyId string) (*ClusterPolicyGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		ClusterPolicyId: clusterPolicyId,
	})
}

// Set cluster policy permissions.
//
// Update all permissions for a specific cluster policy, specifying all users,
// groups, or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the cluster policy and replaces it with the new permissions
// specified in the request body.
func (a *PermissionsClusterPoliciesAPI) Set(ctx context.Context, request ClusterPolicyPermissionsRequest) error {
	return a.impl.Set(ctx, request)
}

// Update cluster policy permissions.
//
// Update permissions for a specific cluster policy.
func (a *PermissionsClusterPoliciesAPI) Update(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsClusters(client *client.DatabricksClient) *PermissionsClustersAPI {
	return &PermissionsClustersAPI{
		impl: &permissionsClustersImpl{
			client: client,
		},
	}
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
type PermissionsClustersAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsClustersService)
	impl PermissionsClustersService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsClustersAPI) WithImpl(impl PermissionsClustersService) *PermissionsClustersAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsClusters API implementation
func (a *PermissionsClustersAPI) Impl() PermissionsClustersService {
	return a.impl
}

// Get cluster permissions.
//
// Get permissions for a specific cluster.
func (a *PermissionsClustersAPI) Get(ctx context.Context, request GetPermissionsClusterRequest) (*ClusterObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get cluster permissions.
//
// Get permissions for a specific cluster.
func (a *PermissionsClustersAPI) GetByClusterId(ctx context.Context, clusterId string) (*ClusterObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsClusterRequest{
		ClusterId: clusterId,
	})
}

// Get cluster permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// clusters.
func (a *PermissionsClustersAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*ClusterGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get cluster permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// clusters.
func (a *PermissionsClustersAPI) GetPermissionLevelsByClusterId(ctx context.Context, clusterId string) (*ClusterGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		ClusterId: clusterId,
	})
}

// Set cluster permissions.
//
// Update all permissions for a specific cluster, specifying all users, groups,
// or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the cluster and replaces it with the new permissions specified
// in the request body.
func (a *PermissionsClustersAPI) Set(ctx context.Context, request ClusterPermissionsRequest) (*ClusterObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update cluster permissions.
//
// Update permissions for a specific cluster.
func (a *PermissionsClustersAPI) Update(ctx context.Context, request ClusterPermissionsRequest) (*ClusterObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsDeltaLiveTables(client *client.DatabricksClient) *PermissionsDeltaLiveTablesAPI {
	return &PermissionsDeltaLiveTablesAPI{
		impl: &permissionsDeltaLiveTablesImpl{
			client: client,
		},
	}
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
type PermissionsDeltaLiveTablesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsDeltaLiveTablesService)
	impl PermissionsDeltaLiveTablesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsDeltaLiveTablesAPI) WithImpl(impl PermissionsDeltaLiveTablesService) *PermissionsDeltaLiveTablesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsDeltaLiveTables API implementation
func (a *PermissionsDeltaLiveTablesAPI) Impl() PermissionsDeltaLiveTablesService {
	return a.impl
}

// Get pipeline permissions.
//
// Get permissions for a specific pipeline.
func (a *PermissionsDeltaLiveTablesAPI) Get(ctx context.Context, request GetPermissionsDeltaLiveTableRequest) (*DeltaLiveTableObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get pipeline permissions.
//
// Get permissions for a specific pipeline.
func (a *PermissionsDeltaLiveTablesAPI) GetByPipelineId(ctx context.Context, pipelineId string) (*DeltaLiveTableObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsDeltaLiveTableRequest{
		PipelineId: pipelineId,
	})
}

// Get pipeline permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// pipelines.
func (a *PermissionsDeltaLiveTablesAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*DeltaLiveTableGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get pipeline permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// pipelines.
func (a *PermissionsDeltaLiveTablesAPI) GetPermissionLevelsByPipelineId(ctx context.Context, pipelineId string) (*DeltaLiveTableGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		PipelineId: pipelineId,
	})
}

// Set pipeline permissions.
//
// Update all permissions for a specific pipeline, specifying all users, groups,
// or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the pipeline and replaces it with the new permissions
// specified in the request body.
func (a *PermissionsDeltaLiveTablesAPI) Set(ctx context.Context, request DeltaLiveTablePermissionsRequest) (*DeltaLiveTableObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update pipeline permissions.
//
// Update permissions for a specific pipeline.
func (a *PermissionsDeltaLiveTablesAPI) Update(ctx context.Context, request DeltaLiveTablePermissionsRequest) (*DeltaLiveTableObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsDirectories(client *client.DatabricksClient) *PermissionsDirectoriesAPI {
	return &PermissionsDirectoriesAPI{
		impl: &permissionsDirectoriesImpl{
			client: client,
		},
	}
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
type PermissionsDirectoriesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsDirectoriesService)
	impl PermissionsDirectoriesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsDirectoriesAPI) WithImpl(impl PermissionsDirectoriesService) *PermissionsDirectoriesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsDirectories API implementation
func (a *PermissionsDirectoriesAPI) Impl() PermissionsDirectoriesService {
	return a.impl
}

// Get directory permissions.
//
// Get permissions for a specific directory.
func (a *PermissionsDirectoriesAPI) Get(ctx context.Context, request GetPermissionsDirectoryRequest) (*DirectoryObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get directory permissions.
//
// Get permissions for a specific directory.
func (a *PermissionsDirectoriesAPI) GetByDirectoryId(ctx context.Context, directoryId string) (*DirectoryObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsDirectoryRequest{
		DirectoryId: directoryId,
	})
}

// Get directory permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// directories.
func (a *PermissionsDirectoriesAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*DirectoryGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get directory permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// directories.
func (a *PermissionsDirectoriesAPI) GetPermissionLevelsByDirectoryId(ctx context.Context, directoryId string) (*DirectoryGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		DirectoryId: directoryId,
	})
}

// Set directory permissions.
//
// Update all permissions for a specific directory, specifying all users,
// groups, or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the directory and replaces it with the new permissions
// specified in the request body.
func (a *PermissionsDirectoriesAPI) Set(ctx context.Context, request DirectoryPermissionsRequest) (*DirectoryObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update directory permissions.
//
// Update permissions for a specific directory.
func (a *PermissionsDirectoriesAPI) Update(ctx context.Context, request DirectoryPermissionsRequest) (*DirectoryObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsJobs(client *client.DatabricksClient) *PermissionsJobsAPI {
	return &PermissionsJobsAPI{
		impl: &permissionsJobsImpl{
			client: client,
		},
	}
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
type PermissionsJobsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsJobsService)
	impl PermissionsJobsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsJobsAPI) WithImpl(impl PermissionsJobsService) *PermissionsJobsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsJobs API implementation
func (a *PermissionsJobsAPI) Impl() PermissionsJobsService {
	return a.impl
}

// Get job permissions.
//
// Get permissions for a specific job.
func (a *PermissionsJobsAPI) Get(ctx context.Context, request GetPermissionsJobRequest) (*JobObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get job permissions.
//
// Get permissions for a specific job.
func (a *PermissionsJobsAPI) GetByJobId(ctx context.Context, jobId string) (*JobObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsJobRequest{
		JobId: jobId,
	})
}

// Get job permission levels.
//
// Returns a JSON representation of the possible permissions levels for jobs.
func (a *PermissionsJobsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*JobGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get job permission levels.
//
// Returns a JSON representation of the possible permissions levels for jobs.
func (a *PermissionsJobsAPI) GetPermissionLevelsByJobId(ctx context.Context, jobId string) (*JobGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		JobId: jobId,
	})
}

// Set job permissions.
//
// Update all permissions for a specific job, specifying all users, groups, or
// service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the job and replaces it with the new permissions specified in
// the request body.
func (a *PermissionsJobsAPI) Set(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update job permissions.
//
// Update permissions for a specific job.
func (a *PermissionsJobsAPI) Update(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsMLflowExperiments(client *client.DatabricksClient) *PermissionsMLflowExperimentsAPI {
	return &PermissionsMLflowExperimentsAPI{
		impl: &permissionsMLflowExperimentsImpl{
			client: client,
		},
	}
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
type PermissionsMLflowExperimentsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsMLflowExperimentsService)
	impl PermissionsMLflowExperimentsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsMLflowExperimentsAPI) WithImpl(impl PermissionsMLflowExperimentsService) *PermissionsMLflowExperimentsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsMLflowExperiments API implementation
func (a *PermissionsMLflowExperimentsAPI) Impl() PermissionsMLflowExperimentsService {
	return a.impl
}

// Get experiment permissions.
//
// Get permissions for a specific experiment.
func (a *PermissionsMLflowExperimentsAPI) Get(ctx context.Context, request GetPermissionsMLflowExperimentRequest) (*MLflowExperimentObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get experiment permissions.
//
// Get permissions for a specific experiment.
func (a *PermissionsMLflowExperimentsAPI) GetByExperimentId(ctx context.Context, experimentId string) (*MLflowExperimentObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsMLflowExperimentRequest{
		ExperimentId: experimentId,
	})
}

// Get experiment permission levels.
//
// Returns a JSON representation of the possible permissions levels for MLflow
// experiments.
func (a *PermissionsMLflowExperimentsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*MLflowExperimentGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get experiment permission levels.
//
// Returns a JSON representation of the possible permissions levels for MLflow
// experiments.
func (a *PermissionsMLflowExperimentsAPI) GetPermissionLevelsByExperimentId(ctx context.Context, experimentId string) (*MLflowExperimentGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		ExperimentId: experimentId,
	})
}

// Set experiment permissions.
//
// Update all permissions for a specific experiment, specifying all users,
// groups, or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the experiment and replaces it with the new permissions
// specified in the request body.
func (a *PermissionsMLflowExperimentsAPI) Set(ctx context.Context, request MLflowExperimentPermissionsRequest) (*MLflowExperimentObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update experiment permissions.
//
// Update permissions for a specific experiment.
func (a *PermissionsMLflowExperimentsAPI) Update(ctx context.Context, request MLflowExperimentPermissionsRequest) (*MLflowExperimentObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsMlFlowRegisteredModels(client *client.DatabricksClient) *PermissionsMlFlowRegisteredModelsAPI {
	return &PermissionsMlFlowRegisteredModelsAPI{
		impl: &permissionsMlFlowRegisteredModelsImpl{
			client: client,
		},
	}
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
type PermissionsMlFlowRegisteredModelsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsMlFlowRegisteredModelsService)
	impl PermissionsMlFlowRegisteredModelsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsMlFlowRegisteredModelsAPI) WithImpl(impl PermissionsMlFlowRegisteredModelsService) *PermissionsMlFlowRegisteredModelsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsMlFlowRegisteredModels API implementation
func (a *PermissionsMlFlowRegisteredModelsAPI) Impl() PermissionsMlFlowRegisteredModelsService {
	return a.impl
}

// Get registered model permissions.
//
// Get permissions for a specific registered model.
func (a *PermissionsMlFlowRegisteredModelsAPI) Get(ctx context.Context, request GetPermissionsMlFlowRegisteredModelRequest) (*MLflowRegisteredModelObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get registered model permissions.
//
// Get permissions for a specific registered model.
func (a *PermissionsMLFlowRegisteredModelsAPI) GetByRegisteredModelId(ctx context.Context, registeredModelId string) (*MLflowRegisteredModelObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsMlFlowRegisteredModelRequest{
		RegisteredModelId: registeredModelId,
	})
}

// Get registered model permission levels.
//
// Returns a JSON representation of the possible permissions levels for MLflow
// models.
func (a *PermissionsMlFlowRegisteredModelsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*MLflowRegisteredModelGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get registered model permission levels.
//
// Returns a JSON representation of the possible permissions levels for MLflow
// models.
func (a *PermissionsMLFlowRegisteredModelsAPI) GetPermissionLevelsByRegisteredModelId(ctx context.Context, registeredModelId string) (*MLflowRegisteredModelGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		RegisteredModelId: registeredModelId,
	})
}

// Set registered model permissions.
//
// Update all permissions for a specific registered model, specifying all users,
// groups, or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the registered model and replaces it with the new permissions
// specified in the request body.
func (a *PermissionsMlFlowRegisteredModelsAPI) Set(ctx context.Context, request MLflowRegisteredModelPermissionsRequest) (*MLflowRegisteredModelObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update registered model permissions.
//
// Update permissions for a specific registered model.
func (a *PermissionsMlFlowRegisteredModelsAPI) Update(ctx context.Context, request MLflowRegisteredModelPermissionsRequest) (*MLflowRegisteredModelObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsNotebooks(client *client.DatabricksClient) *PermissionsNotebooksAPI {
	return &PermissionsNotebooksAPI{
		impl: &permissionsNotebooksImpl{
			client: client,
		},
	}
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
type PermissionsNotebooksAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsNotebooksService)
	impl PermissionsNotebooksService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsNotebooksAPI) WithImpl(impl PermissionsNotebooksService) *PermissionsNotebooksAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsNotebooks API implementation
func (a *PermissionsNotebooksAPI) Impl() PermissionsNotebooksService {
	return a.impl
}

// Get notebook permissions.
//
// Get permissions for a specific notebook.
func (a *PermissionsNotebooksAPI) Get(ctx context.Context, request GetPermissionsNotebookRequest) (*JobObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get notebook permissions.
//
// Get permissions for a specific notebook.
func (a *PermissionsNotebooksAPI) GetByNotebookId(ctx context.Context, notebookId string) (*JobObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsNotebookRequest{
		NotebookId: notebookId,
	})
}

// Get notebook permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// notebooks.
func (a *PermissionsNotebooksAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*JobGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get notebook permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// notebooks.
func (a *PermissionsNotebooksAPI) GetPermissionLevelsByNotebookId(ctx context.Context, notebookId string) (*JobGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		NotebookId: notebookId,
	})
}

// Set notebook permissions.
//
// Update all permissions for a specific notebook, specifying all users, groups,
// or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the notebook and replaces it with the new permissions
// specified in the request body.
func (a *PermissionsNotebooksAPI) Set(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update notebook permissions.
//
// Update permissions for a specific notebook.
func (a *PermissionsNotebooksAPI) Update(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsPasswords(client *client.DatabricksClient) *PermissionsPasswordsAPI {
	return &PermissionsPasswordsAPI{
		impl: &permissionsPasswordsImpl{
			client: client,
		},
	}
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
type PermissionsPasswordsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsPasswordsService)
	impl PermissionsPasswordsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsPasswordsAPI) WithImpl(impl PermissionsPasswordsService) *PermissionsPasswordsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsPasswords API implementation
func (a *PermissionsPasswordsAPI) Impl() PermissionsPasswordsService {
	return a.impl
}

// Get password permissions.
//
// Get password permissions for the entire workspace.
func (a *PermissionsPasswordsAPI) Get(ctx context.Context) (*PasswordObjectPermissions, error) {
	return a.impl.Get(ctx)
}

// Get password permission levels.
//
// Returns a JSON representation of the possible permissions levels for
// passwords.
func (a *PermissionsPasswordsAPI) GetPermissionLevels(ctx context.Context) (*PasswordGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx)
}

// Set password permissions.
//
// Update all password permissions for the entire workspace, specifying all
// users, groups, or service principals.
//
// **WARNING:** This request overwrites all existing permissions and replaces
// them with the new permissions specified in the request body.
func (a *PermissionsPasswordsAPI) Set(ctx context.Context, request PasswordPermissionsRequest) (*PasswordObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update password permissions.
//
// Update password permissions for the entire workspace.
func (a *PermissionsPasswordsAPI) Update(ctx context.Context, request PasswordPermissionsRequest) (*PasswordObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsPools(client *client.DatabricksClient) *PermissionsPoolsAPI {
	return &PermissionsPoolsAPI{
		impl: &permissionsPoolsImpl{
			client: client,
		},
	}
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
type PermissionsPoolsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsPoolsService)
	impl PermissionsPoolsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsPoolsAPI) WithImpl(impl PermissionsPoolsService) *PermissionsPoolsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsPools API implementation
func (a *PermissionsPoolsAPI) Impl() PermissionsPoolsService {
	return a.impl
}

// Get instance pool permissions.
//
// Get permissions for a specific instance pool.
func (a *PermissionsPoolsAPI) Get(ctx context.Context, request GetPermissionsPoolRequest) (*PoolObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get instance pool permissions.
//
// Get permissions for a specific instance pool.
func (a *PermissionsPoolsAPI) GetByInstancePoolId(ctx context.Context, instancePoolId string) (*PoolObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsPoolRequest{
		InstancePoolId: instancePoolId,
	})
}

// Get instance pool permission levels.
//
// Returns a JSON representation of the possible permissions levels for pools.
func (a *PermissionsPoolsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*PoolGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get instance pool permission levels.
//
// Returns a JSON representation of the possible permissions levels for pools.
func (a *PermissionsPoolsAPI) GetPermissionLevelsByInstancePoolId(ctx context.Context, instancePoolId string) (*PoolGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		InstancePoolId: instancePoolId,
	})
}

// Set instance pool permissions.
//
// Update all permissions for a specific instance pool, specifying all users,
// groups, or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the instance pool and replaces it with the new permissions
// specified in the request body.
func (a *PermissionsPoolsAPI) Set(ctx context.Context, request PoolPermissionsRequest) (*PoolObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update instance pool permissions.
//
// Update permissions for a specific instance pool.
func (a *PermissionsPoolsAPI) Update(ctx context.Context, request PoolPermissionsRequest) (*PoolObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsRepos(client *client.DatabricksClient) *PermissionsReposAPI {
	return &PermissionsReposAPI{
		impl: &permissionsReposImpl{
			client: client,
		},
	}
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
type PermissionsReposAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsReposService)
	impl PermissionsReposService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsReposAPI) WithImpl(impl PermissionsReposService) *PermissionsReposAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsRepos API implementation
func (a *PermissionsReposAPI) Impl() PermissionsReposService {
	return a.impl
}

// Get repo permissions.
//
// Get permissions for a specific repo.
func (a *PermissionsReposAPI) Get(ctx context.Context, request GetPermissionsRepoRequest) (*RepoObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get repo permissions.
//
// Get permissions for a specific repo.
func (a *PermissionsReposAPI) GetByRepoId(ctx context.Context, repoId string) (*RepoObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsRepoRequest{
		RepoId: repoId,
	})
}

// Get repo permission levels.
//
// Returns a JSON representation of the possible permissions levels for repos.
func (a *PermissionsReposAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*RepoGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get repo permission levels.
//
// Returns a JSON representation of the possible permissions levels for repos.
func (a *PermissionsReposAPI) GetPermissionLevelsByRepoId(ctx context.Context, repoId string) (*RepoGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		RepoId: repoId,
	})
}

// Set repo permissions.
//
// Update all permissions for a specific repo, specifying all users, groups, or
// service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the repo and replaces it with the new permissions specified in
// the request body.
func (a *PermissionsReposAPI) Set(ctx context.Context, request RepoPermissionsRequest) (*RepoObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update repo permissions.
//
// Update permissions for a specific repo.
func (a *PermissionsReposAPI) Update(ctx context.Context, request RepoPermissionsRequest) (*RepoObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsSqlWarehouses(client *client.DatabricksClient) *PermissionsSqlWarehousesAPI {
	return &PermissionsSqlWarehousesAPI{
		impl: &permissionsSqlWarehousesImpl{
			client: client,
		},
	}
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
type PermissionsSqlWarehousesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsSqlWarehousesService)
	impl PermissionsSqlWarehousesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsSqlWarehousesAPI) WithImpl(impl PermissionsSqlWarehousesService) *PermissionsSqlWarehousesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsSqlWarehouses API implementation
func (a *PermissionsSqlWarehousesAPI) Impl() PermissionsSqlWarehousesService {
	return a.impl
}

// Get SQL warehouse permissions.
//
// Get permissions for a specific SQL warehouse.
func (a *PermissionsSqlWarehousesAPI) Get(ctx context.Context, request GetPermissionsSqlWarehousRequest) (*SqlWarehouseObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get SQL warehouse permissions.
//
// Get permissions for a specific SQL warehouse.
func (a *PermissionsSQLWarehousesAPI) GetByWarehouseId(ctx context.Context, warehouseId string) (*SqlWarehouseObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionsSqlWarehousRequest{
		WarehouseId: warehouseId,
	})
}

// Get SQL warehouse permission levels.
//
// Returns a JSON representation of the possible permissions levels for SQL
// warehouses.
func (a *PermissionsSqlWarehousesAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*SqlWarehouseGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get SQL warehouse permission levels.
//
// Returns a JSON representation of the possible permissions levels for SQL
// warehouses.
func (a *PermissionsSQLWarehousesAPI) GetPermissionLevelsByWarehouseId(ctx context.Context, warehouseId string) (*SqlWarehouseGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		WarehouseId: warehouseId,
	})
}

// Set SQL warehouse permissions.
//
// Update all permissions for a specific SQL warehouse, specifying all users,
// groups, or service principals.
//
// **WARNING:** This request overwrites all existing direct (non-inherited)
// permissions on the SQL warehouse and replaces it with the new permissions
// specified in the request body.
func (a *PermissionsSqlWarehousesAPI) Set(ctx context.Context, request SqlWarehousePermissionsRequest) (*SqlWarehouseObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update SQL warehouse permissions.
//
// Update permissions for a specific SQL warehouse.
func (a *PermissionsSqlWarehousesAPI) Update(ctx context.Context, request SqlWarehousePermissionsRequest) (*SqlWarehouseObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewPermissionsTokens(client *client.DatabricksClient) *PermissionsTokensAPI {
	return &PermissionsTokensAPI{
		impl: &permissionsTokensImpl{
			client: client,
		},
	}
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
type PermissionsTokensAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsTokensService)
	impl PermissionsTokensService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsTokensAPI) WithImpl(impl PermissionsTokensService) *PermissionsTokensAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PermissionsTokens API implementation
func (a *PermissionsTokensAPI) Impl() PermissionsTokensService {
	return a.impl
}

// Get token permissions.
//
// Get token permissions for the entire workspace.
func (a *PermissionsTokensAPI) Get(ctx context.Context) (*TokenObjectPermissions, error) {
	return a.impl.Get(ctx)
}

// Get token permission levels.
//
// Returns a JSON representation of the possible permissions levels for tokens.
func (a *PermissionsTokensAPI) GetPermissionLevels(ctx context.Context) (*TokenGetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx)
}

// Set token permissions.
//
// Update all token permissions for all users, groups, and service principals
// for the entire workspace.
//
// At the end of processing your request, all users and service principals that
// do not have either `CAN_USE` or `CAN_MANAGE` permission either explicitly or
// implicitly due to group assignment no longer have any tokens permissions.
// Affected users or service principals immediately have all their tokens
// deleted.
//
// **WARNING:** This request has powerful effects for workspace security
// configuration and on a workspace's users if they already use tokens. Use with
// caution. This request overwrites all existing token permissions with the data
// in the request body. By omitting reference to an entity that previously had
// permissions, access is stripped and existing tokens are permanently deleted.
func (a *PermissionsTokensAPI) Set(ctx context.Context, request TokenPermissionsRequest) (*TokenObjectPermissions, error) {
	return a.impl.Set(ctx, request)
}

// Update token permissions.
//
// Grant token permissions for one or more users, groups, or service principals.
// You can only grant the Can Use (`CAN_USE`) permission. The Can Manage
// (`CAN_MANAGE`) permission level cannot be granted with this API because it is
// tied automatically to membership in the `admins` group.
func (a *PermissionsTokensAPI) Update(ctx context.Context, request TokenPermissionsRequest) (*TokenObjectPermissions, error) {
	return a.impl.Update(ctx, request)
}

func NewServicePrincipals(client *client.DatabricksClient) *ServicePrincipalsAPI {
	return &ServicePrincipalsAPI{
		impl: &servicePrincipalsImpl{
			client: client,
		},
	}
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
type ServicePrincipalsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ServicePrincipalsService)
	impl ServicePrincipalsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ServicePrincipalsAPI) WithImpl(impl ServicePrincipalsService) *ServicePrincipalsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level ServicePrincipals API implementation
func (a *ServicePrincipalsAPI) Impl() ServicePrincipalsService {
	return a.impl
}

// Create a service principal.
//
// Creates a new service principal in the Databricks workspace.
func (a *ServicePrincipalsAPI) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	return a.impl.Create(ctx, request)
}

// Delete a service principal.
//
// Delete a single service principal in the Databricks workspace.
func (a *ServicePrincipalsAPI) Delete(ctx context.Context, request DeleteServicePrincipalRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a service principal.
//
// Delete a single service principal in the Databricks workspace.
func (a *ServicePrincipalsAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteServicePrincipalRequest{
		Id: id,
	})
}

// Get service principal details.
//
// Gets the details for a single service principal define in the Databricks
// workspace.
func (a *ServicePrincipalsAPI) Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	return a.impl.Get(ctx, request)
}

// Get service principal details.
//
// Gets the details for a single service principal define in the Databricks
// workspace.
func (a *ServicePrincipalsAPI) GetById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.impl.Get(ctx, GetServicePrincipalRequest{
		Id: id,
	})
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalsAPI) ListAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

// ServicePrincipalDisplayNameToIdMap calls [ServicePrincipalsAPI.ListAll] and creates a map of results with [ServicePrincipal].DisplayName as key and [ServicePrincipal].Id as value.
//
// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
//
// Note: All [ServicePrincipal] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalsAPI) ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListServicePrincipalsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.DisplayName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .DisplayName: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByDisplayName calls [ServicePrincipalsAPI.ServicePrincipalDisplayNameToIdMap] and returns a single [ServicePrincipal].
//
// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
//
// Note: All [ServicePrincipal] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalsAPI) GetByDisplayName(ctx context.Context, name string) (*ServicePrincipal, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListServicePrincipalsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]ServicePrincipal{}
	for _, v := range result {
		key := v.DisplayName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("ServicePrincipal named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of ServicePrincipal named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update service principal details.
//
// Partially updates the details of a single service principal in the Databricks
// workspace.
func (a *ServicePrincipalsAPI) Patch(ctx context.Context, request PartialUpdate) error {
	return a.impl.Patch(ctx, request)
}

// Replace service principal.
//
// Updates the details of a single service principal.
//
// This action replaces the existing service principal with the same name.
func (a *ServicePrincipalsAPI) Update(ctx context.Context, request ServicePrincipal) error {
	return a.impl.Update(ctx, request)
}

func NewUsers(client *client.DatabricksClient) *UsersAPI {
	return &UsersAPI{
		impl: &usersImpl{
			client: client,
		},
	}
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
type UsersAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(UsersService)
	impl UsersService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *UsersAPI) WithImpl(impl UsersService) *UsersAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Users API implementation
func (a *UsersAPI) Impl() UsersService {
	return a.impl
}

// Create a new user.
//
// Creates a new user in the Databricks workspace. This new user will also be
// added to the Databricks account.
func (a *UsersAPI) Create(ctx context.Context, request User) (*User, error) {
	return a.impl.Create(ctx, request)
}

// Delete a user.
//
// Deletes a user. Deleting a user from a Databricks workspace also removes
// objects associated with the user.
func (a *UsersAPI) Delete(ctx context.Context, request DeleteUserRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a user.
//
// Deletes a user. Deleting a user from a Databricks workspace also removes
// objects associated with the user.
func (a *UsersAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteUserRequest{
		Id: id,
	})
}

// Get user details.
//
// Gets information for a specific user in Databricks workspace.
func (a *UsersAPI) Get(ctx context.Context, request GetUserRequest) (*User, error) {
	return a.impl.Get(ctx, request)
}

// Get user details.
//
// Gets information for a specific user in Databricks workspace.
func (a *UsersAPI) GetById(ctx context.Context, id string) (*User, error) {
	return a.impl.Get(ctx, GetUserRequest{
		Id: id,
	})
}

// List users.
//
// Gets details for all the users associated with a Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *UsersAPI) ListAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

// UserUserNameToIdMap calls [UsersAPI.ListAll] and creates a map of results with [User].UserName as key and [User].Id as value.
//
// Returns an error if there's more than one [User] with the same .UserName.
//
// Note: All [User] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *UsersAPI) UserUserNameToIdMap(ctx context.Context, request ListUsersRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.UserName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .UserName: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByUserName calls [UsersAPI.UserUserNameToIdMap] and returns a single [User].
//
// Returns an error if there's more than one [User] with the same .UserName.
//
// Note: All [User] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *UsersAPI) GetByUserName(ctx context.Context, name string) (*User, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListUsersRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]User{}
	for _, v := range result {
		key := v.UserName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("User named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of User named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update user details.
//
// Partially updates a user resource by applying the supplied operations on
// specific user attributes.
func (a *UsersAPI) Patch(ctx context.Context, request PartialUpdate) error {
	return a.impl.Patch(ctx, request)
}

// Replace a user.
//
// Replaces a user's information with the data supplied in request.
func (a *UsersAPI) Update(ctx context.Context, request User) error {
	return a.impl.Update(ctx, request)
}

func NewWorkspaceAssignment(client *client.DatabricksClient) *WorkspaceAssignmentAPI {
	return &WorkspaceAssignmentAPI{
		impl: &workspaceAssignmentImpl{
			client: client,
		},
	}
}

// The Workspace Permission Assignment API allows you to manage workspace
// permissions for principals in your account.
type WorkspaceAssignmentAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WorkspaceAssignmentService)
	impl WorkspaceAssignmentService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *WorkspaceAssignmentAPI) WithImpl(impl WorkspaceAssignmentService) *WorkspaceAssignmentAPI {
	a.impl = impl
	return a
}

// Impl returns low-level WorkspaceAssignment API implementation
func (a *WorkspaceAssignmentAPI) Impl() WorkspaceAssignmentService {
	return a.impl
}

// Delete permissions assignment.
//
// Deletes the workspace permissions assignment in a given account and workspace
// for the specified principal.
func (a *WorkspaceAssignmentAPI) Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete permissions assignment.
//
// Deletes the workspace permissions assignment in a given account and workspace
// for the specified principal.
func (a *WorkspaceAssignmentAPI) DeleteByWorkspaceIdAndPrincipalId(ctx context.Context, workspaceId int64, principalId int64) error {
	return a.impl.Delete(ctx, DeleteWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
		PrincipalId: principalId,
	})
}

// List workspace permissions.
//
// Get an array of workspace permissions for the specified account and
// workspace.
func (a *WorkspaceAssignmentAPI) Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error) {
	return a.impl.Get(ctx, request)
}

// List workspace permissions.
//
// Get an array of workspace permissions for the specified account and
// workspace.
func (a *WorkspaceAssignmentAPI) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*WorkspacePermissions, error) {
	return a.impl.Get(ctx, GetWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

// Get permission assignments.
//
// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WorkspaceAssignmentAPI) ListAll(ctx context.Context, request ListWorkspaceAssignmentRequest) ([]PermissionAssignment, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.PermissionAssignments, nil
}

// Get permission assignments.
//
// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
func (a *WorkspaceAssignmentAPI) ListByWorkspaceId(ctx context.Context, workspaceId int64) (*PermissionAssignments, error) {
	return a.impl.List(ctx, ListWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

// Create or update permissions assignment.
//
// Creates or updates the workspace permissions assignment in a given account
// and workspace for the specified principal.
func (a *WorkspaceAssignmentAPI) Update(ctx context.Context, request UpdateWorkspaceAssignments) error {
	return a.impl.Update(ctx, request)
}
