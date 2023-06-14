// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Access Control, Account Access Control Proxy, Account Groups, Account Service Principals, Account Users, Current User, Groups, Permissions, Service Principals, Users, Workspace Assignment, etc.
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

// Get assignable roles for a resource.
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

func NewAccountAccessControlProxy(client *client.DatabricksClient) *AccountAccessControlProxyAPI {
	return &AccountAccessControlProxyAPI{
		impl: &accountAccessControlProxyImpl{
			client: client,
		},
	}
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set. A
// workspace must belong to an account for these APIs to work.
type AccountAccessControlProxyAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AccountAccessControlProxyService)
	impl AccountAccessControlProxyService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AccountAccessControlProxyAPI) WithImpl(impl AccountAccessControlProxyService) *AccountAccessControlProxyAPI {
	a.impl = impl
	return a
}

// Impl returns low-level AccountAccessControlProxy API implementation
func (a *AccountAccessControlProxyAPI) Impl() AccountAccessControlProxyService {
	return a.impl
}

// Get assignable roles for a resource.
//
// Gets all the roles that can be granted on an account-level resource. A role
// is grantable if the rule set on the resource can contain an access rule of
// the role.
func (a *AccountAccessControlProxyAPI) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	return a.impl.GetAssignableRolesForResource(ctx, request)
}

// Get a rule set.
//
// Get a rule set by its name. A rule set is always attached to a resource and
// contains a list of access rules on the said resource. Currently only a
// default rule set for each resource is supported.
func (a *AccountAccessControlProxyAPI) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	return a.impl.GetRuleSet(ctx, request)
}

// Update a rule set.
//
// Replace the rules of a rule set. First, use a GET rule set request to read
// the current version of the rule set before modifying it. This pattern helps
// prevent conflicts between concurrent updates.
func (a *AccountAccessControlProxyAPI) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
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

func NewPermissions(client *client.DatabricksClient) *PermissionsAPI {
	return &PermissionsAPI{
		impl: &permissionsImpl{
			client: client,
		},
	}
}

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints.
type PermissionsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsService)
	impl PermissionsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsAPI) WithImpl(impl PermissionsService) *PermissionsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Permissions API implementation
func (a *PermissionsAPI) Impl() PermissionsService {
	return a.impl
}

// Get object permissions.
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error) {
	return a.impl.Get(ctx, request)
}

// Get object permissions.
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) GetByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*ObjectPermissions, error) {
	return a.impl.Get(ctx, GetPermissionRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

// Get permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *PermissionsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *PermissionsAPI) GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

// Set permissions.
//
// Sets permissions on object. Objects can inherit permissions from their parent
// objects and root objects.
func (a *PermissionsAPI) Set(ctx context.Context, request PermissionsRequest) error {
	return a.impl.Set(ctx, request)
}

// Update permission.
//
// Updates the permissions on an object.
func (a *PermissionsAPI) Update(ctx context.Context, request PermissionsRequest) error {
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
