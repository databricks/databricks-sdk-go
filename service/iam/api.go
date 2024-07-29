// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Access Control, Account Access Control Proxy, Account Groups, Account Service Principals, Account Users, Current User, Groups, Permission Migration, Permissions, Service Principals, Users, Workspace Assignment, etc.
package iam

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type AccountAccessControlInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockAccountAccessControlInterface instead.
	WithImpl(impl AccountAccessControlService) AccountAccessControlInterface

	// Impl returns low-level AccountAccessControl API implementation
	// Deprecated: use MockAccountAccessControlInterface instead.
	Impl() AccountAccessControlService

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

func NewAccountAccessControl(client *client.DatabricksClient) *AccountAccessControlAPI {
	return &AccountAccessControlAPI{
		AccountAccessControlService: &accountAccessControlImpl{
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
	AccountAccessControlService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockAccountAccessControlInterface instead.
func (a *AccountAccessControlAPI) WithImpl(impl AccountAccessControlService) AccountAccessControlInterface {
	return &AccountAccessControlAPI{
		AccountAccessControlService: impl,
	}
}

// Impl returns low-level AccountAccessControl API implementation
// Deprecated: use MockAccountAccessControlInterface instead.
func (a *AccountAccessControlAPI) Impl() AccountAccessControlService {
	return a.AccountAccessControlService
}

type AccountAccessControlProxyInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockAccountAccessControlProxyInterface instead.
	WithImpl(impl AccountAccessControlProxyService) AccountAccessControlProxyInterface

	// Impl returns low-level AccountAccessControlProxy API implementation
	// Deprecated: use MockAccountAccessControlProxyInterface instead.
	Impl() AccountAccessControlProxyService

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

func NewAccountAccessControlProxy(client *client.DatabricksClient) *AccountAccessControlProxyAPI {
	return &AccountAccessControlProxyAPI{
		AccountAccessControlProxyService: &accountAccessControlProxyImpl{
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
	AccountAccessControlProxyService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockAccountAccessControlProxyInterface instead.
func (a *AccountAccessControlProxyAPI) WithImpl(impl AccountAccessControlProxyService) AccountAccessControlProxyInterface {
	return &AccountAccessControlProxyAPI{
		AccountAccessControlProxyService: impl,
	}
}

// Impl returns low-level AccountAccessControlProxy API implementation
// Deprecated: use MockAccountAccessControlProxyInterface instead.
func (a *AccountAccessControlProxyAPI) Impl() AccountAccessControlProxyService {
	return a.AccountAccessControlProxyService
}

type AccountGroupsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockAccountGroupsInterface instead.
	WithImpl(impl AccountGroupsService) AccountGroupsInterface

	// Impl returns low-level AccountGroups API implementation
	// Deprecated: use MockAccountGroupsInterface instead.
	Impl() AccountGroupsService

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

	// GroupDisplayNameToIdMap calls [AccountGroupsAPI.ListAll] and creates a map of results with [Group].DisplayName as key and [Group].Id as value.
	//
	// Returns an error if there's more than one [Group] with the same .DisplayName.
	//
	// Note: All [Group] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	GroupDisplayNameToIdMap(ctx context.Context, request ListAccountGroupsRequest) (map[string]string, error)

	// GetByDisplayName calls [AccountGroupsAPI.GroupDisplayNameToIdMap] and returns a single [Group].
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

func NewAccountGroups(client *client.DatabricksClient) *AccountGroupsAPI {
	return &AccountGroupsAPI{
		AccountGroupsService: &accountGroupsImpl{
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
	AccountGroupsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockAccountGroupsInterface instead.
func (a *AccountGroupsAPI) WithImpl(impl AccountGroupsService) AccountGroupsInterface {
	return &AccountGroupsAPI{
		AccountGroupsService: impl,
	}
}

// Impl returns low-level AccountGroups API implementation
// Deprecated: use MockAccountGroupsInterface instead.
func (a *AccountGroupsAPI) Impl() AccountGroupsService {
	return a.AccountGroupsService
}

// Delete a group.
//
// Deletes a group from the Databricks account.
func (a *AccountGroupsAPI) DeleteById(ctx context.Context, id string) error {
	return a.AccountGroupsService.Delete(ctx, DeleteAccountGroupRequest{
		Id: id,
	})
}

// Get group details.
//
// Gets the information for a specific group in the Databricks account.
func (a *AccountGroupsAPI) GetById(ctx context.Context, id string) (*Group, error) {
	return a.AccountGroupsService.Get(ctx, GetAccountGroupRequest{
		Id: id,
	})
}

// List group details.
//
// Gets all details of the groups associated with the Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountGroupsAPI) List(ctx context.Context, request ListAccountGroupsRequest) listing.Iterator[Group] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListAccountGroupsRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.AccountGroupsService.List(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Resources
	}
	getNextReq := func(resp *ListGroupsResponse) *ListAccountGroupsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[Group, string](
		iterator,
		func(item Group) string {
			return item.Id
		})
	return dedupedIterator
}

// List group details.
//
// Gets all details of the groups associated with the Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountGroupsAPI) ListAll(ctx context.Context, request ListAccountGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

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

type AccountServicePrincipalsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockAccountServicePrincipalsInterface instead.
	WithImpl(impl AccountServicePrincipalsService) AccountServicePrincipalsInterface

	// Impl returns low-level AccountServicePrincipals API implementation
	// Deprecated: use MockAccountServicePrincipalsInterface instead.
	Impl() AccountServicePrincipalsService

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

	// ServicePrincipalDisplayNameToIdMap calls [AccountServicePrincipalsAPI.ListAll] and creates a map of results with [ServicePrincipal].DisplayName as key and [ServicePrincipal].Id as value.
	//
	// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
	//
	// Note: All [ServicePrincipal] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListAccountServicePrincipalsRequest) (map[string]string, error)

	// GetByDisplayName calls [AccountServicePrincipalsAPI.ServicePrincipalDisplayNameToIdMap] and returns a single [ServicePrincipal].
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

func NewAccountServicePrincipals(client *client.DatabricksClient) *AccountServicePrincipalsAPI {
	return &AccountServicePrincipalsAPI{
		AccountServicePrincipalsService: &accountServicePrincipalsImpl{
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
	AccountServicePrincipalsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockAccountServicePrincipalsInterface instead.
func (a *AccountServicePrincipalsAPI) WithImpl(impl AccountServicePrincipalsService) AccountServicePrincipalsInterface {
	return &AccountServicePrincipalsAPI{
		AccountServicePrincipalsService: impl,
	}
}

// Impl returns low-level AccountServicePrincipals API implementation
// Deprecated: use MockAccountServicePrincipalsInterface instead.
func (a *AccountServicePrincipalsAPI) Impl() AccountServicePrincipalsService {
	return a.AccountServicePrincipalsService
}

// Delete a service principal.
//
// Delete a single service principal in the Databricks account.
func (a *AccountServicePrincipalsAPI) DeleteById(ctx context.Context, id string) error {
	return a.AccountServicePrincipalsService.Delete(ctx, DeleteAccountServicePrincipalRequest{
		Id: id,
	})
}

// Get service principal details.
//
// Gets the details for a single service principal define in the Databricks
// account.
func (a *AccountServicePrincipalsAPI) GetById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.AccountServicePrincipalsService.Get(ctx, GetAccountServicePrincipalRequest{
		Id: id,
	})
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountServicePrincipalsAPI) List(ctx context.Context, request ListAccountServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.AccountServicePrincipalsService.List(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalResponse) []ServicePrincipal {
		return resp.Resources
	}
	getNextReq := func(resp *ListServicePrincipalResponse) *ListAccountServicePrincipalsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[ServicePrincipal, string](
		iterator,
		func(item ServicePrincipal) string {
			return item.Id
		})
	return dedupedIterator
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountServicePrincipalsAPI) ListAll(ctx context.Context, request ListAccountServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

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

type AccountUsersInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockAccountUsersInterface instead.
	WithImpl(impl AccountUsersService) AccountUsersInterface

	// Impl returns low-level AccountUsers API implementation
	// Deprecated: use MockAccountUsersInterface instead.
	Impl() AccountUsersService

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

	// UserUserNameToIdMap calls [AccountUsersAPI.ListAll] and creates a map of results with [User].UserName as key and [User].Id as value.
	//
	// Returns an error if there's more than one [User] with the same .UserName.
	//
	// Note: All [User] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	UserUserNameToIdMap(ctx context.Context, request ListAccountUsersRequest) (map[string]string, error)

	// GetByUserName calls [AccountUsersAPI.UserUserNameToIdMap] and returns a single [User].
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

func NewAccountUsers(client *client.DatabricksClient) *AccountUsersAPI {
	return &AccountUsersAPI{
		AccountUsersService: &accountUsersImpl{
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
	AccountUsersService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockAccountUsersInterface instead.
func (a *AccountUsersAPI) WithImpl(impl AccountUsersService) AccountUsersInterface {
	return &AccountUsersAPI{
		AccountUsersService: impl,
	}
}

// Impl returns low-level AccountUsers API implementation
// Deprecated: use MockAccountUsersInterface instead.
func (a *AccountUsersAPI) Impl() AccountUsersService {
	return a.AccountUsersService
}

// Delete a user.
//
// Deletes a user. Deleting a user from a Databricks account also removes
// objects associated with the user.
func (a *AccountUsersAPI) DeleteById(ctx context.Context, id string) error {
	return a.AccountUsersService.Delete(ctx, DeleteAccountUserRequest{
		Id: id,
	})
}

// Get user details.
//
// Gets information for a specific user in Databricks account.
func (a *AccountUsersAPI) GetById(ctx context.Context, id string) (*User, error) {
	return a.AccountUsersService.Get(ctx, GetAccountUserRequest{
		Id: id,
	})
}

// List users.
//
// Gets details for all the users associated with a Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountUsersAPI) List(ctx context.Context, request ListAccountUsersRequest) listing.Iterator[User] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListAccountUsersRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.AccountUsersService.List(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Resources
	}
	getNextReq := func(resp *ListUsersResponse) *ListAccountUsersRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[User, string](
		iterator,
		func(item User) string {
			return item.Id
		})
	return dedupedIterator
}

// List users.
//
// Gets details for all the users associated with a Databricks account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AccountUsersAPI) ListAll(ctx context.Context, request ListAccountUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

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

type CurrentUserInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockCurrentUserInterface instead.
	WithImpl(impl CurrentUserService) CurrentUserInterface

	// Impl returns low-level CurrentUser API implementation
	// Deprecated: use MockCurrentUserInterface instead.
	Impl() CurrentUserService

	// Get current user info.
	//
	// Get details about the current method caller's identity.
	Me(ctx context.Context) (*User, error)
}

func NewCurrentUser(client *client.DatabricksClient) *CurrentUserAPI {
	return &CurrentUserAPI{
		CurrentUserService: &currentUserImpl{
			client: client,
		},
	}
}

// This API allows retrieving information about currently authenticated user or
// service principal.
type CurrentUserAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CurrentUserService)
	CurrentUserService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockCurrentUserInterface instead.
func (a *CurrentUserAPI) WithImpl(impl CurrentUserService) CurrentUserInterface {
	return &CurrentUserAPI{
		CurrentUserService: impl,
	}
}

// Impl returns low-level CurrentUser API implementation
// Deprecated: use MockCurrentUserInterface instead.
func (a *CurrentUserAPI) Impl() CurrentUserService {
	return a.CurrentUserService
}

type GroupsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockGroupsInterface instead.
	WithImpl(impl GroupsService) GroupsInterface

	// Impl returns low-level Groups API implementation
	// Deprecated: use MockGroupsInterface instead.
	Impl() GroupsService

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

	// GroupDisplayNameToIdMap calls [GroupsAPI.ListAll] and creates a map of results with [Group].DisplayName as key and [Group].Id as value.
	//
	// Returns an error if there's more than one [Group] with the same .DisplayName.
	//
	// Note: All [Group] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	GroupDisplayNameToIdMap(ctx context.Context, request ListGroupsRequest) (map[string]string, error)

	// GetByDisplayName calls [GroupsAPI.GroupDisplayNameToIdMap] and returns a single [Group].
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

func NewGroups(client *client.DatabricksClient) *GroupsAPI {
	return &GroupsAPI{
		GroupsService: &groupsImpl{
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
	GroupsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockGroupsInterface instead.
func (a *GroupsAPI) WithImpl(impl GroupsService) GroupsInterface {
	return &GroupsAPI{
		GroupsService: impl,
	}
}

// Impl returns low-level Groups API implementation
// Deprecated: use MockGroupsInterface instead.
func (a *GroupsAPI) Impl() GroupsService {
	return a.GroupsService
}

// Delete a group.
//
// Deletes a group from the Databricks workspace.
func (a *GroupsAPI) DeleteById(ctx context.Context, id string) error {
	return a.GroupsService.Delete(ctx, DeleteGroupRequest{
		Id: id,
	})
}

// Get group details.
//
// Gets the information for a specific group in the Databricks workspace.
func (a *GroupsAPI) GetById(ctx context.Context, id string) (*Group, error) {
	return a.GroupsService.Get(ctx, GetGroupRequest{
		Id: id,
	})
}

// List group details.
//
// Gets all details of the groups associated with the Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GroupsAPI) List(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListGroupsRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.GroupsService.List(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Resources
	}
	getNextReq := func(resp *ListGroupsResponse) *ListGroupsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[Group, string](
		iterator,
		func(item Group) string {
			return item.Id
		})
	return dedupedIterator
}

// List group details.
//
// Gets all details of the groups associated with the Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GroupsAPI) ListAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

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

type PermissionMigrationInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockPermissionMigrationInterface instead.
	WithImpl(impl PermissionMigrationService) PermissionMigrationInterface

	// Impl returns low-level PermissionMigration API implementation
	// Deprecated: use MockPermissionMigrationInterface instead.
	Impl() PermissionMigrationService

	// Migrate Permissions.
	MigratePermissions(ctx context.Context, request MigratePermissionsRequest) (*MigratePermissionsResponse, error)
}

func NewPermissionMigration(client *client.DatabricksClient) *PermissionMigrationAPI {
	return &PermissionMigrationAPI{
		PermissionMigrationService: &permissionMigrationImpl{
			client: client,
		},
	}
}

// APIs for migrating acl permissions, used only by the ucx tool:
// https://github.com/databrickslabs/ucx
type PermissionMigrationAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionMigrationService)
	PermissionMigrationService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockPermissionMigrationInterface instead.
func (a *PermissionMigrationAPI) WithImpl(impl PermissionMigrationService) PermissionMigrationInterface {
	return &PermissionMigrationAPI{
		PermissionMigrationService: impl,
	}
}

// Impl returns low-level PermissionMigration API implementation
// Deprecated: use MockPermissionMigrationInterface instead.
func (a *PermissionMigrationAPI) Impl() PermissionMigrationService {
	return a.PermissionMigrationService
}

type PermissionsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockPermissionsInterface instead.
	WithImpl(impl PermissionsService) PermissionsInterface

	// Impl returns low-level Permissions API implementation
	// Deprecated: use MockPermissionsInterface instead.
	Impl() PermissionsService

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
	// Sets permissions on an object. Objects can inherit permissions from their
	// parent objects or root object.
	Set(ctx context.Context, request PermissionsRequest) (*ObjectPermissions, error)

	// Update object permissions.
	//
	// Updates the permissions on an object. Objects can inherit permissions from
	// their parent objects or root object.
	Update(ctx context.Context, request PermissionsRequest) (*ObjectPermissions, error)
}

func NewPermissions(client *client.DatabricksClient) *PermissionsAPI {
	return &PermissionsAPI{
		PermissionsService: &permissionsImpl{
			client: client,
		},
	}
}

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints.
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
// users can read, run, edit, or manage directories, files, and notebooks.
//
// For the mapping of the required permissions for specific actions or abilities
// and other important information, see [Access Control].
//
// Note that to manage access control on service principals, use **[Account
// Access Control Proxy](:service:accountaccesscontrolproxy)**.
//
// [Access Control]: https://docs.databricks.com/security/auth-authz/access-control/index.html
type PermissionsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsService)
	PermissionsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockPermissionsInterface instead.
func (a *PermissionsAPI) WithImpl(impl PermissionsService) PermissionsInterface {
	return &PermissionsAPI{
		PermissionsService: impl,
	}
}

// Impl returns low-level Permissions API implementation
// Deprecated: use MockPermissionsInterface instead.
func (a *PermissionsAPI) Impl() PermissionsService {
	return a.PermissionsService
}

// Get object permissions.
//
// Gets the permissions of an object. Objects can inherit permissions from their
// parent objects or root object.
func (a *PermissionsAPI) GetByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*ObjectPermissions, error) {
	return a.PermissionsService.Get(ctx, GetPermissionRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

// Get object permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *PermissionsAPI) GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error) {
	return a.PermissionsService.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

type ServicePrincipalsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockServicePrincipalsInterface instead.
	WithImpl(impl ServicePrincipalsService) ServicePrincipalsInterface

	// Impl returns low-level ServicePrincipals API implementation
	// Deprecated: use MockServicePrincipalsInterface instead.
	Impl() ServicePrincipalsService

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

	// ServicePrincipalDisplayNameToIdMap calls [ServicePrincipalsAPI.ListAll] and creates a map of results with [ServicePrincipal].DisplayName as key and [ServicePrincipal].Id as value.
	//
	// Returns an error if there's more than one [ServicePrincipal] with the same .DisplayName.
	//
	// Note: All [ServicePrincipal] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListServicePrincipalsRequest) (map[string]string, error)

	// GetByDisplayName calls [ServicePrincipalsAPI.ServicePrincipalDisplayNameToIdMap] and returns a single [ServicePrincipal].
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

func NewServicePrincipals(client *client.DatabricksClient) *ServicePrincipalsAPI {
	return &ServicePrincipalsAPI{
		ServicePrincipalsService: &servicePrincipalsImpl{
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
	ServicePrincipalsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockServicePrincipalsInterface instead.
func (a *ServicePrincipalsAPI) WithImpl(impl ServicePrincipalsService) ServicePrincipalsInterface {
	return &ServicePrincipalsAPI{
		ServicePrincipalsService: impl,
	}
}

// Impl returns low-level ServicePrincipals API implementation
// Deprecated: use MockServicePrincipalsInterface instead.
func (a *ServicePrincipalsAPI) Impl() ServicePrincipalsService {
	return a.ServicePrincipalsService
}

// Delete a service principal.
//
// Delete a single service principal in the Databricks workspace.
func (a *ServicePrincipalsAPI) DeleteById(ctx context.Context, id string) error {
	return a.ServicePrincipalsService.Delete(ctx, DeleteServicePrincipalRequest{
		Id: id,
	})
}

// Get service principal details.
//
// Gets the details for a single service principal define in the Databricks
// workspace.
func (a *ServicePrincipalsAPI) GetById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.ServicePrincipalsService.Get(ctx, GetServicePrincipalRequest{
		Id: id,
	})
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalsAPI) List(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.ServicePrincipalsService.List(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalResponse) []ServicePrincipal {
		return resp.Resources
	}
	getNextReq := func(resp *ListServicePrincipalResponse) *ListServicePrincipalsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[ServicePrincipal, string](
		iterator,
		func(item ServicePrincipal) string {
			return item.Id
		})
	return dedupedIterator
}

// List service principals.
//
// Gets the set of service principals associated with a Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalsAPI) ListAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

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

type UsersInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockUsersInterface instead.
	WithImpl(impl UsersService) UsersInterface

	// Impl returns low-level Users API implementation
	// Deprecated: use MockUsersInterface instead.
	Impl() UsersService

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

	// UserUserNameToIdMap calls [UsersAPI.ListAll] and creates a map of results with [User].UserName as key and [User].Id as value.
	//
	// Returns an error if there's more than one [User] with the same .UserName.
	//
	// Note: All [User] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	UserUserNameToIdMap(ctx context.Context, request ListUsersRequest) (map[string]string, error)

	// GetByUserName calls [UsersAPI.UserUserNameToIdMap] and returns a single [User].
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
	// Sets permissions on all passwords. Passwords can inherit permissions from
	// their root object.
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

func NewUsers(client *client.DatabricksClient) *UsersAPI {
	return &UsersAPI{
		UsersService: &usersImpl{
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
	UsersService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockUsersInterface instead.
func (a *UsersAPI) WithImpl(impl UsersService) UsersInterface {
	return &UsersAPI{
		UsersService: impl,
	}
}

// Impl returns low-level Users API implementation
// Deprecated: use MockUsersInterface instead.
func (a *UsersAPI) Impl() UsersService {
	return a.UsersService
}

// Delete a user.
//
// Deletes a user. Deleting a user from a Databricks workspace also removes
// objects associated with the user.
func (a *UsersAPI) DeleteById(ctx context.Context, id string) error {
	return a.UsersService.Delete(ctx, DeleteUserRequest{
		Id: id,
	})
}

// Get user details.
//
// Gets information for a specific user in Databricks workspace.
func (a *UsersAPI) GetById(ctx context.Context, id string) (*User, error) {
	return a.UsersService.Get(ctx, GetUserRequest{
		Id: id,
	})
}

// List users.
//
// Gets details for all the users associated with a Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *UsersAPI) List(ctx context.Context, request ListUsersRequest) listing.Iterator[User] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 100
	}
	getNextPage := func(ctx context.Context, req ListUsersRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.UsersService.List(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Resources
	}
	getNextReq := func(resp *ListUsersResponse) *ListUsersRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[User, string](
		iterator,
		func(item User) string {
			return item.Id
		})
	return dedupedIterator
}

// List users.
//
// Gets details for all the users associated with a Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *UsersAPI) ListAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

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

type WorkspaceAssignmentInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockWorkspaceAssignmentInterface instead.
	WithImpl(impl WorkspaceAssignmentService) WorkspaceAssignmentInterface

	// Impl returns low-level WorkspaceAssignment API implementation
	// Deprecated: use MockWorkspaceAssignmentInterface instead.
	Impl() WorkspaceAssignmentService

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

func NewWorkspaceAssignment(client *client.DatabricksClient) *WorkspaceAssignmentAPI {
	return &WorkspaceAssignmentAPI{
		WorkspaceAssignmentService: &workspaceAssignmentImpl{
			client: client,
		},
	}
}

// The Workspace Permission Assignment API allows you to manage workspace
// permissions for principals in your account.
type WorkspaceAssignmentAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WorkspaceAssignmentService)
	WorkspaceAssignmentService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockWorkspaceAssignmentInterface instead.
func (a *WorkspaceAssignmentAPI) WithImpl(impl WorkspaceAssignmentService) WorkspaceAssignmentInterface {
	return &WorkspaceAssignmentAPI{
		WorkspaceAssignmentService: impl,
	}
}

// Impl returns low-level WorkspaceAssignment API implementation
// Deprecated: use MockWorkspaceAssignmentInterface instead.
func (a *WorkspaceAssignmentAPI) Impl() WorkspaceAssignmentService {
	return a.WorkspaceAssignmentService
}

// Delete permissions assignment.
//
// Deletes the workspace permissions assignment in a given account and workspace
// for the specified principal.
func (a *WorkspaceAssignmentAPI) DeleteByWorkspaceIdAndPrincipalId(ctx context.Context, workspaceId int64, principalId int64) error {
	return a.WorkspaceAssignmentService.Delete(ctx, DeleteWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
		PrincipalId: principalId,
	})
}

// List workspace permissions.
//
// Get an array of workspace permissions for the specified account and
// workspace.
func (a *WorkspaceAssignmentAPI) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*WorkspacePermissions, error) {
	return a.WorkspaceAssignmentService.Get(ctx, GetWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

// Get permission assignments.
//
// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WorkspaceAssignmentAPI) List(ctx context.Context, request ListWorkspaceAssignmentRequest) listing.Iterator[PermissionAssignment] {

	getNextPage := func(ctx context.Context, req ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.WorkspaceAssignmentService.List(ctx, req)
	}
	getItems := func(resp *PermissionAssignments) []PermissionAssignment {
		return resp.PermissionAssignments
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get permission assignments.
//
// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WorkspaceAssignmentAPI) ListAll(ctx context.Context, request ListWorkspaceAssignmentRequest) ([]PermissionAssignment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PermissionAssignment](ctx, iterator)
}

// Get permission assignments.
//
// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
func (a *WorkspaceAssignmentAPI) ListByWorkspaceId(ctx context.Context, workspaceId int64) (*PermissionAssignments, error) {
	return a.WorkspaceAssignmentService.List(ctx, ListWorkspaceAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}
