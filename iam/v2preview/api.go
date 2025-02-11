// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Access Control Preview, Account Access Control Preview, Account Access Control Proxy Preview, Account Groups Preview, Account Service Principals Preview, Account Users Preview, Current User Preview, Groups Preview, Permission Migration Preview, Permissions Preview, Service Principals Preview, Users Preview, Workspace Assignment Preview, etc.
package iampreview

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type AccessControlPreviewInterface interface {

	// Check access policy to a resource.
	CheckPolicy(ctx context.Context, request CheckPolicyRequest) (*CheckPolicyResponse, error)
}

func NewAccessControlPreview(client *client.DatabricksClient) *AccessControlPreviewAPI {
	return &AccessControlPreviewAPI{
		accessControlPreviewImpl: accessControlPreviewImpl{
			client: client,
		},
	}
}

// Rule based Access Control for Databricks Resources.
type AccessControlPreviewAPI struct {
	accessControlPreviewImpl
}

type AccountAccessControlPreviewInterface interface {

	// Get assignable roles for a resource.
	//
	// Gets all the roles that can be granted on an account level resource. A role
	// is grantable if the rule set on the resource can contain an access rule of
	// the role.
	GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error)

	// Get a rule set.
	//
	// Get a rule set by its name. A rule set is always attached to a resource and
	// contains a list of access rules on the said resource. Currently only a
	// default rule set for each resource is supported.
	GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error)

	// Update a rule set.
	//
	// Replace the rules of a rule set. First, use get to read the current version
	// of the rule set before modifying it. This pattern helps prevent conflicts
	// between concurrent updates.
	UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error)
}

func NewAccountAccessControlPreview(client *client.DatabricksClient) *AccountAccessControlPreviewAPI {
	return &AccountAccessControlPreviewAPI{
		accountAccessControlPreviewImpl: accountAccessControlPreviewImpl{
			client: client,
		},
	}
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set.
type AccountAccessControlPreviewAPI struct {
	accountAccessControlPreviewImpl
}

type AccountAccessControlProxyPreviewInterface interface {

	// Get assignable roles for a resource.
	//
	// Gets all the roles that can be granted on an account-level resource. A role
	// is grantable if the rule set on the resource can contain an access rule of
	// the role.
	GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error)

	// Get a rule set.
	//
	// Get a rule set by its name. A rule set is always attached to a resource and
	// contains a list of access rules on the said resource. Currently only a
	// default rule set for each resource is supported.
	GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error)

	// Update a rule set.
	//
	// Replace the rules of a rule set. First, use a GET rule set request to read
	// the current version of the rule set before modifying it. This pattern helps
	// prevent conflicts between concurrent updates.
	UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error)
}

func NewAccountAccessControlProxyPreview(client *client.DatabricksClient) *AccountAccessControlProxyPreviewAPI {
	return &AccountAccessControlProxyPreviewAPI{
		accountAccessControlProxyPreviewImpl: accountAccessControlProxyPreviewImpl{
			client: client,
		},
	}
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set. A
// workspace must belong to an account for these APIs to work.
type AccountAccessControlProxyPreviewAPI struct {
	accountAccessControlProxyPreviewImpl
}

type AccountGroupsPreviewInterface interface {

	// Create a new group.
	//
	// Creates a group in the Databricks account with a unique name, using the
	// supplied group details.
	Create(ctx context.Context, request Group) (*Group, error)

	// Delete a group.
	//
	// Deletes a group from the Databricks account.
	Delete(ctx context.Context, request DeleteAccountGroupRequest) error

	// Delete a group.
	//
	// Deletes a group from the Databricks account.
	DeleteById(ctx context.Context, id string) error

	// Get group details.
	//
	// Gets the information for a specific group in the Databricks account.
	Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error)

	// Get group details.
	//
	// Gets the information for a specific group in the Databricks account.
	GetById(ctx context.Context, id string) (*Group, error)

	// List group details.
	//
	// Gets all details of the groups associated with the Databricks account.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListAccountGroupsRequest) listing.Iterator[Group]

	// List group details.
	//
	// Gets all details of the groups associated with the Databricks account.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListAccountGroupsRequest) ([]Group, error)

	// GroupDisplayNameToIdMap calls [AccountGroupsPreviewAPI.ListAll] and creates a map of results with [Group].DisplayName as key and [Group].Id as value.
	//
	// Returns an error if there's more than one [Group] with the same .DisplayName.
	//
	// Note: All [Group] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	GroupDisplayNameToIdMap(ctx context.Context, request ListAccountGroupsRequest) (map[string]string, error)

	// GetByDisplayName calls [AccountGroupsPreviewAPI.GroupDisplayNameToIdMap] and returns a single [Group].
	//
	// Returns an error if there's more than one [Group] with the same .DisplayName.
	//
	// Note: All [Group] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByDisplayName(ctx context.Context, name string) (*Group, error)

	// Update group details.
	//
	// Partially updates the details of a group.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace a group.
	//
	// Updates the details of a group by replacing the entire group entity.
	Update(ctx context.Context, request Group) error
}

func NewAccountGroupsPreview(client *client.DatabricksClient) *AccountGroupsPreviewAPI {
	return &AccountGroupsPreviewAPI{
		accountGroupsPreviewImpl: accountGroupsPreviewImpl{
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
type AccountGroupsPreviewAPI struct {
	accountGroupsPreviewImpl
}

// Delete a group.
//
// Deletes a group from the Databricks account.
func (a *AccountGroupsPreviewAPI) DeleteById(ctx context.Context, id string) error {
	return a.accountGroupsPreviewImpl.Delete(ctx, DeleteAccountGroupRequest{
		Id: id,
	})
}

// Get group details.
//
// Gets the information for a specific group in the Databricks account.
func (a *AccountGroupsPreviewAPI) GetById(ctx context.Context, id string) (*Group, error) {
	return a.accountGroupsPreviewImpl.Get(ctx, GetAccountGroupRequest{
		Id: id,
	})
}

// GroupDisplayNameToIdMap calls [AccountGroupsPreviewAPI.ListAll] and creates a map of results with [Group].DisplayName as key and [Group].Id as value.
//
// Returns an error if there's more than one [Group] with the same .DisplayName.
//
// Note: All [Group] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountGroupsPreviewAPI) GroupDisplayNameToIdMap(ctx context.Context, request ListAccountGroupsRequest) (map[string]string, error) {
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

// GetByDisplayName calls [AccountGroupsPreviewAPI.GroupDisplayNameToIdMap] and returns a single [Group].
//
// Returns an error if there's more than one [Group] with the same .DisplayName.
//
// Note: All [Group] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountGroupsPreviewAPI) GetByDisplayName(ctx context.Context, name string) (*Group, error) {
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

type AccountServicePrincipalsPreviewInterface interface {

	// Create a service principal.
	//
	// Creates a new service principal in the Databricks account.
	Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error)

	// Delete a service principal.
	//
	// Delete a single service principal in the Databricks account.
	Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error

	// Delete a service principal.
	//
	// Delete a single service principal in the Databricks account.
	DeleteById(ctx context.Context, id string) error

	// Get service principal details.
	//
	// Gets the details for a single service principal define in the Databricks
	// account.
	Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error)

	// Get service principal details.
	//
	// Gets the details for a single service principal define in the Databricks
	// account.
	GetById(ctx context.Context, id string) (*ServicePrincipal, error)

	// List service principals.
	//
	// Gets the set of service principals associated with a Databricks account.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListAccountServicePrincipalsRequest) listing.Iterator[ServicePrincipal]

	// List service principals.
	//
	// Gets the set of service principals associated with a Databricks account.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListAccountServicePrincipalsRequest) ([]ServicePrincipal, error)

	// ServicePrincipalDisplayNameToIdMap calls [AccountServicePrincipalsPreviewAPI.ListAll] and creates a map of results with [ServicePrincipal].DisplayName as key and [ServicePrincipal].Id as value.
	//
	// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
	//
	// Note: All [ServicePrincipal] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListAccountServicePrincipalsRequest) (map[string]string, error)

	// GetByDisplayName calls [AccountServicePrincipalsPreviewAPI.ServicePrincipalDisplayNameToIdMap] and returns a single [ServicePrincipal].
	//
	// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
	//
	// Note: All [ServicePrincipal] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByDisplayName(ctx context.Context, name string) (*ServicePrincipal, error)

	// Update service principal details.
	//
	// Partially updates the details of a single service principal in the Databricks
	// account.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace service principal.
	//
	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	Update(ctx context.Context, request ServicePrincipal) error
}

func NewAccountServicePrincipalsPreview(client *client.DatabricksClient) *AccountServicePrincipalsPreviewAPI {
	return &AccountServicePrincipalsPreviewAPI{
		accountServicePrincipalsPreviewImpl: accountServicePrincipalsPreviewImpl{
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
type AccountServicePrincipalsPreviewAPI struct {
	accountServicePrincipalsPreviewImpl
}

// Delete a service principal.
//
// Delete a single service principal in the Databricks account.
func (a *AccountServicePrincipalsPreviewAPI) DeleteById(ctx context.Context, id string) error {
	return a.accountServicePrincipalsPreviewImpl.Delete(ctx, DeleteAccountServicePrincipalRequest{
		Id: id,
	})
}

// Get service principal details.
//
// Gets the details for a single service principal define in the Databricks
// account.
func (a *AccountServicePrincipalsPreviewAPI) GetById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.accountServicePrincipalsPreviewImpl.Get(ctx, GetAccountServicePrincipalRequest{
		Id: id,
	})
}

// ServicePrincipalDisplayNameToIdMap calls [AccountServicePrincipalsPreviewAPI.ListAll] and creates a map of results with [ServicePrincipal].DisplayName as key and [ServicePrincipal].Id as value.
//
// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
//
// Note: All [ServicePrincipal] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountServicePrincipalsPreviewAPI) ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListAccountServicePrincipalsRequest) (map[string]string, error) {
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

// GetByDisplayName calls [AccountServicePrincipalsPreviewAPI.ServicePrincipalDisplayNameToIdMap] and returns a single [ServicePrincipal].
//
// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
//
// Note: All [ServicePrincipal] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountServicePrincipalsPreviewAPI) GetByDisplayName(ctx context.Context, name string) (*ServicePrincipal, error) {
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

type AccountUsersPreviewInterface interface {

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

	// Delete a user.
	//
	// Deletes a user. Deleting a user from a Databricks account also removes
	// objects associated with the user.
	DeleteById(ctx context.Context, id string) error

	// Get user details.
	//
	// Gets information for a specific user in Databricks account.
	Get(ctx context.Context, request GetAccountUserRequest) (*User, error)

	// Get user details.
	//
	// Gets information for a specific user in Databricks account.
	GetById(ctx context.Context, id string) (*User, error)

	// List users.
	//
	// Gets details for all the users associated with a Databricks account.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListAccountUsersRequest) listing.Iterator[User]

	// List users.
	//
	// Gets details for all the users associated with a Databricks account.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListAccountUsersRequest) ([]User, error)

	// UserUserNameToIdMap calls [AccountUsersPreviewAPI.ListAll] and creates a map of results with [User].UserName as key and [User].Id as value.
	//
	// Returns an error if there's more than one [User] with the same .UserName.
	//
	// Note: All [User] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	UserUserNameToIdMap(ctx context.Context, request ListAccountUsersRequest) (map[string]string, error)

	// GetByUserName calls [AccountUsersPreviewAPI.UserUserNameToIdMap] and returns a single [User].
	//
	// Returns an error if there's more than one [User] with the same .UserName.
	//
	// Note: All [User] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByUserName(ctx context.Context, name string) (*User, error)

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

func NewAccountUsersPreview(client *client.DatabricksClient) *AccountUsersPreviewAPI {
	return &AccountUsersPreviewAPI{
		accountUsersPreviewImpl: accountUsersPreviewImpl{
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
type AccountUsersPreviewAPI struct {
	accountUsersPreviewImpl
}

// Delete a user.
//
// Deletes a user. Deleting a user from a Databricks account also removes
// objects associated with the user.
func (a *AccountUsersPreviewAPI) DeleteById(ctx context.Context, id string) error {
	return a.accountUsersPreviewImpl.Delete(ctx, DeleteAccountUserRequest{
		Id: id,
	})
}

// Get user details.
//
// Gets information for a specific user in Databricks account.
func (a *AccountUsersPreviewAPI) GetById(ctx context.Context, id string) (*User, error) {
	return a.accountUsersPreviewImpl.Get(ctx, GetAccountUserRequest{
		Id: id,
	})
}

// UserUserNameToIdMap calls [AccountUsersPreviewAPI.ListAll] and creates a map of results with [User].UserName as key and [User].Id as value.
//
// Returns an error if there's more than one [User] with the same .UserName.
//
// Note: All [User] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountUsersPreviewAPI) UserUserNameToIdMap(ctx context.Context, request ListAccountUsersRequest) (map[string]string, error) {
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

// GetByUserName calls [AccountUsersPreviewAPI.UserUserNameToIdMap] and returns a single [User].
//
// Returns an error if there's more than one [User] with the same .UserName.
//
// Note: All [User] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountUsersPreviewAPI) GetByUserName(ctx context.Context, name string) (*User, error) {
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

type CurrentUserPreviewInterface interface {

	// Get current user info.
	//
	// Get details about the current method caller's identity.
	Me(ctx context.Context) (*User, error)
}

func NewCurrentUserPreview(client *client.DatabricksClient) *CurrentUserPreviewAPI {
	return &CurrentUserPreviewAPI{
		currentUserPreviewImpl: currentUserPreviewImpl{
			client: client,
		},
	}
}

// This API allows retrieving information about currently authenticated user or
// service principal.
type CurrentUserPreviewAPI struct {
	currentUserPreviewImpl
}

type GroupsPreviewInterface interface {

	// Create a new group.
	//
	// Creates a group in the Databricks workspace with a unique name, using the
	// supplied group details.
	Create(ctx context.Context, request Group) (*Group, error)

	// Delete a group.
	//
	// Deletes a group from the Databricks workspace.
	Delete(ctx context.Context, request DeleteGroupRequest) error

	// Delete a group.
	//
	// Deletes a group from the Databricks workspace.
	DeleteById(ctx context.Context, id string) error

	// Get group details.
	//
	// Gets the information for a specific group in the Databricks workspace.
	Get(ctx context.Context, request GetGroupRequest) (*Group, error)

	// Get group details.
	//
	// Gets the information for a specific group in the Databricks workspace.
	GetById(ctx context.Context, id string) (*Group, error)

	// List group details.
	//
	// Gets all details of the groups associated with the Databricks workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group]

	// List group details.
	//
	// Gets all details of the groups associated with the Databricks workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListGroupsRequest) ([]Group, error)

	// GroupDisplayNameToIdMap calls [GroupsPreviewAPI.ListAll] and creates a map of results with [Group].DisplayName as key and [Group].Id as value.
	//
	// Returns an error if there's more than one [Group] with the same .DisplayName.
	//
	// Note: All [Group] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	GroupDisplayNameToIdMap(ctx context.Context, request ListGroupsRequest) (map[string]string, error)

	// GetByDisplayName calls [GroupsPreviewAPI.GroupDisplayNameToIdMap] and returns a single [Group].
	//
	// Returns an error if there's more than one [Group] with the same .DisplayName.
	//
	// Note: All [Group] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByDisplayName(ctx context.Context, name string) (*Group, error)

	// Update group details.
	//
	// Partially updates the details of a group.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace a group.
	//
	// Updates the details of a group by replacing the entire group entity.
	Update(ctx context.Context, request Group) error
}

func NewGroupsPreview(client *client.DatabricksClient) *GroupsPreviewAPI {
	return &GroupsPreviewAPI{
		groupsPreviewImpl: groupsPreviewImpl{
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
type GroupsPreviewAPI struct {
	groupsPreviewImpl
}

// Delete a group.
//
// Deletes a group from the Databricks workspace.
func (a *GroupsPreviewAPI) DeleteById(ctx context.Context, id string) error {
	return a.groupsPreviewImpl.Delete(ctx, DeleteGroupRequest{
		Id: id,
	})
}

// Get group details.
//
// Gets the information for a specific group in the Databricks workspace.
func (a *GroupsPreviewAPI) GetById(ctx context.Context, id string) (*Group, error) {
	return a.groupsPreviewImpl.Get(ctx, GetGroupRequest{
		Id: id,
	})
}

// GroupDisplayNameToIdMap calls [GroupsPreviewAPI.ListAll] and creates a map of results with [Group].DisplayName as key and [Group].Id as value.
//
// Returns an error if there's more than one [Group] with the same .DisplayName.
//
// Note: All [Group] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GroupsPreviewAPI) GroupDisplayNameToIdMap(ctx context.Context, request ListGroupsRequest) (map[string]string, error) {
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

// GetByDisplayName calls [GroupsPreviewAPI.GroupDisplayNameToIdMap] and returns a single [Group].
//
// Returns an error if there's more than one [Group] with the same .DisplayName.
//
// Note: All [Group] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GroupsPreviewAPI) GetByDisplayName(ctx context.Context, name string) (*Group, error) {
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

type PermissionMigrationPreviewInterface interface {

	// Migrate Permissions.
	MigratePermissions(ctx context.Context, request MigratePermissionsRequest) (*MigratePermissionsResponse, error)
}

func NewPermissionMigrationPreview(client *client.DatabricksClient) *PermissionMigrationPreviewAPI {
	return &PermissionMigrationPreviewAPI{
		permissionMigrationPreviewImpl: permissionMigrationPreviewImpl{
			client: client,
		},
	}
}

// APIs for migrating acl permissions, used only by the ucx tool:
// https://github.com/databrickslabs/ucx
type PermissionMigrationPreviewAPI struct {
	permissionMigrationPreviewImpl
}

type PermissionsPreviewInterface interface {

	// Get object permissions.
	//
	// Gets the permissions of an object. Objects can inherit permissions from their
	// parent objects or root object.
	Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error)

	// Get object permissions.
	//
	// Gets the permissions of an object. Objects can inherit permissions from their
	// parent objects or root object.
	GetByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*ObjectPermissions, error)

	// Get object permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error)

	// Get object permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error)

	// Set object permissions.
	//
	// Sets permissions on an object, replacing existing permissions if they exist.
	// Deletes all direct permissions if none are specified. Objects can inherit
	// permissions from their parent objects or root object.
	Set(ctx context.Context, request PermissionsRequest) (*ObjectPermissions, error)

	// Update object permissions.
	//
	// Updates the permissions on an object. Objects can inherit permissions from
	// their parent objects or root object.
	Update(ctx context.Context, request PermissionsRequest) (*ObjectPermissions, error)
}

func NewPermissionsPreview(client *client.DatabricksClient) *PermissionsPreviewAPI {
	return &PermissionsPreviewAPI{
		permissionsPreviewImpl: permissionsPreviewImpl{
			client: client,
		},
	}
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
type PermissionsPreviewAPI struct {
	permissionsPreviewImpl
}

// Get object permissions.
//
// Gets the permissions of an object. Objects can inherit permissions from their
// parent objects or root object.
func (a *PermissionsPreviewAPI) GetByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*ObjectPermissions, error) {
	return a.permissionsPreviewImpl.Get(ctx, GetPermissionRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

// Get object permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *PermissionsPreviewAPI) GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error) {
	return a.permissionsPreviewImpl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

type ServicePrincipalsPreviewInterface interface {

	// Create a service principal.
	//
	// Creates a new service principal in the Databricks workspace.
	Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error)

	// Delete a service principal.
	//
	// Delete a single service principal in the Databricks workspace.
	Delete(ctx context.Context, request DeleteServicePrincipalRequest) error

	// Delete a service principal.
	//
	// Delete a single service principal in the Databricks workspace.
	DeleteById(ctx context.Context, id string) error

	// Get service principal details.
	//
	// Gets the details for a single service principal define in the Databricks
	// workspace.
	Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error)

	// Get service principal details.
	//
	// Gets the details for a single service principal define in the Databricks
	// workspace.
	GetById(ctx context.Context, id string) (*ServicePrincipal, error)

	// List service principals.
	//
	// Gets the set of service principals associated with a Databricks workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal]

	// List service principals.
	//
	// Gets the set of service principals associated with a Databricks workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error)

	// ServicePrincipalDisplayNameToIdMap calls [ServicePrincipalsPreviewAPI.ListAll] and creates a map of results with [ServicePrincipal].DisplayName as key and [ServicePrincipal].Id as value.
	//
	// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
	//
	// Note: All [ServicePrincipal] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListServicePrincipalsRequest) (map[string]string, error)

	// GetByDisplayName calls [ServicePrincipalsPreviewAPI.ServicePrincipalDisplayNameToIdMap] and returns a single [ServicePrincipal].
	//
	// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
	//
	// Note: All [ServicePrincipal] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByDisplayName(ctx context.Context, name string) (*ServicePrincipal, error)

	// Update service principal details.
	//
	// Partially updates the details of a single service principal in the Databricks
	// workspace.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replace service principal.
	//
	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	Update(ctx context.Context, request ServicePrincipal) error
}

func NewServicePrincipalsPreview(client *client.DatabricksClient) *ServicePrincipalsPreviewAPI {
	return &ServicePrincipalsPreviewAPI{
		servicePrincipalsPreviewImpl: servicePrincipalsPreviewImpl{
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
type ServicePrincipalsPreviewAPI struct {
	servicePrincipalsPreviewImpl
}

// Delete a service principal.
//
// Delete a single service principal in the Databricks workspace.
func (a *ServicePrincipalsPreviewAPI) DeleteById(ctx context.Context, id string) error {
	return a.servicePrincipalsPreviewImpl.Delete(ctx, DeleteServicePrincipalRequest{
		Id: id,
	})
}

// Get service principal details.
//
// Gets the details for a single service principal define in the Databricks
// workspace.
func (a *ServicePrincipalsPreviewAPI) GetById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.servicePrincipalsPreviewImpl.Get(ctx, GetServicePrincipalRequest{
		Id: id,
	})
}

// ServicePrincipalDisplayNameToIdMap calls [ServicePrincipalsPreviewAPI.ListAll] and creates a map of results with [ServicePrincipal].DisplayName as key and [ServicePrincipal].Id as value.
//
// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
//
// Note: All [ServicePrincipal] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalsPreviewAPI) ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListServicePrincipalsRequest) (map[string]string, error) {
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

// GetByDisplayName calls [ServicePrincipalsPreviewAPI.ServicePrincipalDisplayNameToIdMap] and returns a single [ServicePrincipal].
//
// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
//
// Note: All [ServicePrincipal] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalsPreviewAPI) GetByDisplayName(ctx context.Context, name string) (*ServicePrincipal, error) {
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

type UsersPreviewInterface interface {

	// Create a new user.
	//
	// Creates a new user in the Databricks workspace. This new user will also be
	// added to the Databricks account.
	Create(ctx context.Context, request User) (*User, error)

	// Delete a user.
	//
	// Deletes a user. Deleting a user from a Databricks workspace also removes
	// objects associated with the user.
	Delete(ctx context.Context, request DeleteUserRequest) error

	// Delete a user.
	//
	// Deletes a user. Deleting a user from a Databricks workspace also removes
	// objects associated with the user.
	DeleteById(ctx context.Context, id string) error

	// Get user details.
	//
	// Gets information for a specific user in Databricks workspace.
	Get(ctx context.Context, request GetUserRequest) (*User, error)

	// Get user details.
	//
	// Gets information for a specific user in Databricks workspace.
	GetById(ctx context.Context, id string) (*User, error)

	// Get password permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context) (*GetPasswordPermissionLevelsResponse, error)

	// Get password permissions.
	//
	// Gets the permissions of all passwords. Passwords can inherit permissions from
	// their root object.
	GetPermissions(ctx context.Context) (*PasswordPermissions, error)

	// List users.
	//
	// Gets details for all the users associated with a Databricks workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListUsersRequest) listing.Iterator[User]

	// List users.
	//
	// Gets details for all the users associated with a Databricks workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListUsersRequest) ([]User, error)

	// UserUserNameToIdMap calls [UsersPreviewAPI.ListAll] and creates a map of results with [User].UserName as key and [User].Id as value.
	//
	// Returns an error if there's more than one [User] with the same .UserName.
	//
	// Note: All [User] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	UserUserNameToIdMap(ctx context.Context, request ListUsersRequest) (map[string]string, error)

	// GetByUserName calls [UsersPreviewAPI.UserUserNameToIdMap] and returns a single [User].
	//
	// Returns an error if there's more than one [User] with the same .UserName.
	//
	// Note: All [User] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByUserName(ctx context.Context, name string) (*User, error)

	// Update user details.
	//
	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes.
	Patch(ctx context.Context, request PartialUpdate) error

	// Set password permissions.
	//
	// Sets permissions on an object, replacing existing permissions if they exist.
	// Deletes all direct permissions if none are specified. Objects can inherit
	// permissions from their root object.
	SetPermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error)

	// Replace a user.
	//
	// Replaces a user's information with the data supplied in request.
	Update(ctx context.Context, request User) error

	// Update password permissions.
	//
	// Updates the permissions on all passwords. Passwords can inherit permissions
	// from their root object.
	UpdatePermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error)
}

func NewUsersPreview(client *client.DatabricksClient) *UsersPreviewAPI {
	return &UsersPreviewAPI{
		usersPreviewImpl: usersPreviewImpl{
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
type UsersPreviewAPI struct {
	usersPreviewImpl
}

// Delete a user.
//
// Deletes a user. Deleting a user from a Databricks workspace also removes
// objects associated with the user.
func (a *UsersPreviewAPI) DeleteById(ctx context.Context, id string) error {
	return a.usersPreviewImpl.Delete(ctx, DeleteUserRequest{
		Id: id,
	})
}

// Get user details.
//
// Gets information for a specific user in Databricks workspace.
func (a *UsersPreviewAPI) GetById(ctx context.Context, id string) (*User, error) {
	return a.usersPreviewImpl.Get(ctx, GetUserRequest{
		Id: id,
	})
}

// UserUserNameToIdMap calls [UsersPreviewAPI.ListAll] and creates a map of results with [User].UserName as key and [User].Id as value.
//
// Returns an error if there's more than one [User] with the same .UserName.
//
// Note: All [User] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *UsersPreviewAPI) UserUserNameToIdMap(ctx context.Context, request ListUsersRequest) (map[string]string, error) {
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

// GetByUserName calls [UsersPreviewAPI.UserUserNameToIdMap] and returns a single [User].
//
// Returns an error if there's more than one [User] with the same .UserName.
//
// Note: All [User] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *UsersPreviewAPI) GetByUserName(ctx context.Context, name string) (*User, error) {
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

type WorkspaceAssignmentPreviewInterface interface {

	// Delete permissions assignment.
	//
	// Deletes the workspace permissions assignment in a given account and workspace
	// for the specified principal.
	Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error

	// Delete permissions assignment.
	//
	// Deletes the workspace permissions assignment in a given account and workspace
	// for the specified principal.
	DeleteByWorkspaceIdAndPrincipalId(ctx context.Context, workspaceId int64, principalId int64) error

	// List workspace permissions.
	//
	// Get an array of workspace permissions for the specified account and
	// workspace.
	Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error)

	// List workspace permissions.
	//
	// Get an array of workspace permissions for the specified account and
	// workspace.
	GetByWorkspaceId(ctx context.Context, workspaceId int64) (*WorkspacePermissions, error)

	// Get permission assignments.
	//
	// Get the permission assignments for the specified Databricks account and
	// Databricks workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListWorkspaceAssignmentRequest) listing.Iterator[PermissionAssignment]

	// Get permission assignments.
	//
	// Get the permission assignments for the specified Databricks account and
	// Databricks workspace.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListWorkspaceAssignmentRequest) ([]PermissionAssignment, error)

	// Get permission assignments.
	//
	// Get the permission assignments for the specified Databricks account and
	// Databricks workspace.
	ListByWorkspaceId(ctx context.Context, workspaceId int64) (*PermissionAssignments, error)

	// Create or update permissions assignment.
	//
	// Creates or updates the workspace permissions assignment in a given account
	// and workspace for the specified principal.
	Update(ctx context.Context, request UpdateWorkspaceAssignments) (*PermissionAssignment, error)
}

func NewWorkspaceAssignmentPreview(client *client.DatabricksClient) *WorkspaceAssignmentPreviewAPI {
	return &WorkspaceAssignmentPreviewAPI{
		workspaceAssignmentPreviewImpl: workspaceAssignmentPreviewImpl{
			client: client,
		},
	}
}

// The Workspace Permission Assignment API allows you to manage workspace
// permissions for principals in your account.
type WorkspaceAssignmentPreviewAPI struct {
	workspaceAssignmentPreviewImpl
}

// Delete permissions assignment.
//
// Deletes the workspace permissions assignment in a given account and workspace
// for the specified principal.
func (a *WorkspaceAssignmentPreviewAPI) DeleteByWorkspaceIdAndPrincipalId(ctx context.Context, workspaceId int64, principalId int64) error {
	return a.workspaceAssignmentPreviewImpl.Delete(ctx, DeleteWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
		PrincipalId: principalId,
	})
}

// List workspace permissions.
//
// Get an array of workspace permissions for the specified account and
// workspace.
func (a *WorkspaceAssignmentPreviewAPI) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*WorkspacePermissions, error) {
	return a.workspaceAssignmentPreviewImpl.Get(ctx, GetWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

// Get permission assignments.
//
// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
func (a *WorkspaceAssignmentPreviewAPI) ListByWorkspaceId(ctx context.Context, workspaceId int64) (*PermissionAssignments, error) {
	return a.workspaceAssignmentPreviewImpl.internalList(ctx, ListWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}
