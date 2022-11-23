// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package secrets

import (
	"context"
)

func NewSecrets(client databricksClient) *SecretsAPI {
	return &SecretsAPI{
		impl: &secretsImpl{
			client: client,
		},
	}
}

// The Secrets API allows you to manage secrets, secret scopes, and access
// permissions.
//
// Sometimes accessing data requires that you authenticate to external data
// sources through JDBC. Instead of directly entering your credentials into a
// notebook, use Databricks secrets to store your credentials and reference them
// in notebooks and jobs.
//
// Administrators, secret creators, and users granted permission can read
// Databricks secrets. While Databricks makes an effort to redact secret values
// that might be displayed in notebooks, it is not possible to prevent such
// users from reading secrets.
type SecretsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(SecretsService)
	impl SecretsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *SecretsAPI) WithImpl(impl SecretsService) *SecretsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Secrets API implementation
func (a *SecretsAPI) Impl() SecretsService {
	return a.impl
}

// Create a new secret scope
//
// The scope name must consist of alphanumeric characters, dashes, underscores,
// and periods, and may not exceed 128 characters. The maximum number of scopes
// in a workspace is 100.
func (a *SecretsAPI) CreateScope(ctx context.Context, request CreateScope) error {
	return a.impl.CreateScope(ctx, request)
}

// Delete an ACL
//
// Deletes the given ACL on the given scope.
//
// Users must have the “MANAGE“ permission to invoke this API. Throws
// “RESOURCE_DOES_NOT_EXIST“ if no such secret scope, principal, or ACL
// exists. Throws “PERMISSION_DENIED“ if the user does not have permission to
// make this API call.
func (a *SecretsAPI) DeleteAcl(ctx context.Context, request DeleteAcl) error {
	return a.impl.DeleteAcl(ctx, request)
}

// Delete a secret scope
//
// Deletes a secret scope.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if the scope does not exist. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *SecretsAPI) DeleteScope(ctx context.Context, request DeleteScope) error {
	return a.impl.DeleteScope(ctx, request)
}

// Delete a secret scope
//
// Deletes a secret scope.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if the scope does not exist. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *SecretsAPI) DeleteScopeByScope(ctx context.Context, scope string) error {
	return a.impl.DeleteScope(ctx, DeleteScope{
		Scope: scope,
	})
}

// Delete a secret
//
// Deletes the secret stored in this secret scope. You must have “WRITE“ or
// “MANAGE“ permission on the secret scope.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope or secret exists.
// Throws “PERMISSION_DENIED“ if the user does not have permission to make
// this API call.
func (a *SecretsAPI) DeleteSecret(ctx context.Context, request DeleteSecret) error {
	return a.impl.DeleteSecret(ctx, request)
}

// Get secret ACL details
//
// Gets the details about the given ACL, such as the group and permission. Users
// must have the “MANAGE“ permission to invoke this API.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *SecretsAPI) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {
	return a.impl.GetAcl(ctx, request)
}

// Lists ACLs
//
// List the ACLs for a given secret scope. Users must have the “MANAGE“
// permission to invoke this API.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SecretsAPI) ListAclsAll(ctx context.Context, request ListAclsRequest) ([]AclItem, error) {
	response, err := a.impl.ListAcls(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}

// Lists ACLs
//
// List the ACLs for a given secret scope. Users must have the “MANAGE“
// permission to invoke this API.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *SecretsAPI) ListAclsByScope(ctx context.Context, scope string) (*ListAclsResponse, error) {
	return a.impl.ListAcls(ctx, ListAclsRequest{
		Scope: scope,
	})
}

// List all scopes
//
// Lists all secret scopes available in the workspace.
//
// Throws “PERMISSION_DENIED“ if the user does not have permission to make
// this API call.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SecretsAPI) ListScopesAll(ctx context.Context) ([]SecretScope, error) {
	response, err := a.impl.ListScopes(ctx)
	if err != nil {
		return nil, err
	}
	return response.Scopes, nil
}

// List secret keys
//
// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
//
// The lastUpdatedTimestamp returned is in milliseconds since epoch. Throws
// “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SecretsAPI) ListSecretsAll(ctx context.Context, request ListSecretsRequest) ([]SecretMetadata, error) {
	response, err := a.impl.ListSecrets(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Secrets, nil
}

// List secret keys
//
// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
//
// The lastUpdatedTimestamp returned is in milliseconds since epoch. Throws
// “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *SecretsAPI) ListSecretsByScope(ctx context.Context, scope string) (*ListSecretsResponse, error) {
	return a.impl.ListSecrets(ctx, ListSecretsRequest{
		Scope: scope,
	})
}

// Create/update an ACL
//
// Creates or overwrites the Access Control List (ACL) associated with the given
// principal (user or group) on the specified scope point.
//
// In general, a user or group will use the most powerful permission available
// to them, and permissions are ordered as follows:
//
// * “MANAGE“ - Allowed to change ACLs, and read and write to this secret
// scope. * “WRITE“ - Allowed to read and write to this secret scope. *
// “READ“ - Allowed to read this secret scope and list what secrets are
// available.
//
// Note that in general, secret values can only be read from within a
// command\non a cluster (for example, through a notebook). There is no API to
// read the actual secret value material outside of a cluster. However, the
// user's permission will be applied based on who is executing the command, and
// they must have at least READ permission.
//
// Users must have the “MANAGE“ permission to invoke this API.
//
// The principal is a user or group name corresponding to an existing Databricks
// principal to be granted or revoked access.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “RESOURCE_ALREADY_EXISTS“ if a permission for the principal already exists.
// Throws “INVALID_PARAMETER_VALUE“ if the permission is invalid. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *SecretsAPI) PutAcl(ctx context.Context, request PutAcl) error {
	return a.impl.PutAcl(ctx, request)
}

// Add a secret
//
// Inserts a secret under the provided scope with the given name. If a secret
// already exists with the same name, this command overwrites the existing
// secret's value. The server encrypts the secret using the secret scope's
// encryption settings before storing it.
//
// You must have “WRITE“ or “MANAGE“ permission on the secret scope. The
// secret key must consist of alphanumeric characters, dashes, underscores, and
// periods, and cannot exceed 128 characters. The maximum allowed secret value
// size is 128 KB. The maximum number of secrets in a given scope is 1000.
//
// The input fields "string_value" or "bytes_value" specify the type of the
// secret, which will determine the value returned when the secret value is
// requested. Exactly one must be specified.
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “RESOURCE_LIMIT_EXCEEDED“ if maximum number of secrets in scope is
// exceeded. Throws “INVALID_PARAMETER_VALUE“ if the key name or value length
// is invalid. Throws “PERMISSION_DENIED“ if the user does not have permission
// to make this API call.
func (a *SecretsAPI) PutSecret(ctx context.Context, request PutSecret) error {
	return a.impl.PutSecret(ctx, request)
}
