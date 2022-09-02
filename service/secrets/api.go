// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package secrets

import (
	"context"
	

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewSecrets(client *client.DatabricksClient) SecretsService {
	return &SecretsAPI{
		client: client,
	}
}

type SecretsAPI struct {
	client *client.DatabricksClient
}

// Creates a new secret scope. The scope name must consist of alphanumeric
// characters, dashes, underscores, and periods, and may not exceed 128
// characters. The maximum number of scopes in a workspace is 100. Example
// request: .. code:: { &#34;scope&#34;: &#34;my-simple-databricks-scope&#34;,
// &#34;initial_manage_principal&#34;: &#34;users&#34; &#34;scope_backend_type&#34;:
// &#34;databricks|azure_keyvault&#34;, # below is only required if scope type is
// azure_keyvault &#34;backend_azure_keyvault&#34;: { &#34;resource_id&#34;:
// &#34;/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/xxxx/providers/Microsoft.KeyVault/vaults/xxxx&#34;,
// &#34;tenant_id&#34;: &#34;xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx&#34;, &#34;dns_name&#34;:
// &#34;https://xxxx.vault.azure.net/&#34;, } } If ``initial_manage_principal`` is
// specified, the initial ACL applied to the scope is applied to the supplied
// principal (user or group) with ``MANAGE`` permissions. The only supported
// principal for this option is the group ``users``, which contains all users in
// the workspace. If ``initial_manage_principal`` is not specified, the initial
// ACL with ``MANAGE`` permission applied to the scope is assigned to the API
// request issuer&#39;s user identity. If ``scope_backend_type`` is
// ``azure_keyvault``, a secret scope is created with secrets from a given Azure
// KeyVault. The caller must provide the keyvault_resource_id and the tenant_id
// for the key vault. If ``scope_backend_type`` is ``databricks`` or is
// unspecified, an empty secret scope is created and stored in Databricks&#39;s own
// storage. Throws ``RESOURCE_ALREADY_EXISTS`` if a scope with the given name
// already exists. Throws ``RESOURCE_LIMIT_EXCEEDED`` if maximum number of
// scopes in the workspace is exceeded. Throws ``INVALID_PARAMETER_VALUE`` if
// the scope name is invalid.
func (a *SecretsAPI) CreateScope(ctx context.Context, request CreateScopeRequest) error {
	path := "/api/2.0/secrets/scopes/create"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Deletes the given ACL on the given scope. Users must have the ``MANAGE``
// permission to invoke this API. Example request: .. code:: { &#34;scope&#34;:
// &#34;my-secret-scope&#34;, &#34;principal&#34;: &#34;data-scientists&#34; } Throws
// ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope, principal, or ACL
// exists. Throws ``PERMISSION_DENIED`` if the user does not have permission to
// make this API call.
func (a *SecretsAPI) DeleteAcl(ctx context.Context, request DeleteAclRequest) error {
	path := "/api/2.0/secrets/acls/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Deletes a secret scope. Example request: .. code:: { &#34;scope&#34;:
// &#34;my-secret-scope&#34; } Throws ``RESOURCE_DOES_NOT_EXIST`` if the scope does not
// exist. Throws ``PERMISSION_DENIED`` if the user does not have permission to
// make this API call.
func (a *SecretsAPI) DeleteScope(ctx context.Context, request DeleteScopeRequest) error {
	path := "/api/2.0/secrets/scopes/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Deletes the secret stored in this secret scope. You must have ``WRITE`` or
// ``MANAGE`` permission on the Secret Scope. Example request: .. code:: {
// &#34;scope&#34;: &#34;my-secret-scope&#34;, &#34;key&#34;: &#34;my-secret-key&#34; } Throws
// ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope or secret exists. Throws
// ``PERMISSION_DENIED`` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) DeleteSecret(ctx context.Context, request DeleteSecretRequest) error {
	path := "/api/2.0/secrets/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Describes the details about the given ACL, such as the group and permission.
// Users must have the ``MANAGE`` permission to invoke this API. Example
// response: .. code:: { &#34;principal&#34;: &#34;data-scientists&#34;, &#34;permission&#34;: &#34;READ&#34; }
// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
// ``PERMISSION_DENIED`` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) GetAcl(ctx context.Context, request GetAclRequest) (*GetAclResponse, error) {
	var getAclResponse GetAclResponse
	path := "/api/2.0/secrets/acls/get"
	err := a.client.Get(ctx, path, request, &getAclResponse)
	return &getAclResponse, err
}

// Lists the ACLs set on the given scope. Users must have the ``MANAGE``
// permission to invoke this API. Example response: .. code:: { &#34;acls&#34;: [{
// &#34;principal&#34;: &#34;admins&#34;, &#34;permission&#34;: &#34;MANAGE&#34; },{ &#34;principal&#34;:
// &#34;data-scientists&#34;, &#34;permission&#34;: &#34;READ&#34; }] } Throws
// ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
// ``PERMISSION_DENIED`` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) ListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error) {
	var listAclsResponse ListAclsResponse
	path := "/api/2.0/secrets/acls/list"
	err := a.client.Get(ctx, path, request, &listAclsResponse)
	return &listAclsResponse, err
}

// Lists all secret scopes available in the workspace. Example response: ..
// code:: { &#34;scopes&#34;: [{ &#34;name&#34;: &#34;my-databricks-scope&#34;, &#34;backend_type&#34;:
// &#34;DATABRICKS&#34; },{ &#34;name&#34;: &#34;mount-points&#34;, &#34;backend_type&#34;: &#34;DATABRICKS&#34; }] }
// Throws ``PERMISSION_DENIED`` if the user does not have permission to make
// this API call.
func (a *SecretsAPI) ListScopes(ctx context.Context) (*ListScopesResponse, error) {
	var listScopesResponse ListScopesResponse
	path := "/api/2.0/secrets/scopes/list"
	err := a.client.Get(ctx, path, nil, &listScopesResponse)
	return &listScopesResponse, err
}

// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call. Example response: .. code:: { &#34;secrets&#34;: [
// { &#34;key&#34;: &#34;my-string-key&#34;&#34;, &#34;last_updated_timestamp&#34;: &#34;1520467595000&#34; }, {
// &#34;key&#34;: &#34;my-byte-key&#34;, &#34;last_updated_timestamp&#34;: &#34;1520467595000&#34; }, ] } The
// lastUpdatedTimestamp returned is in milliseconds since epoch. Throws
// ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
// ``PERMISSION_DENIED`` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) ListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error) {
	var listSecretsResponse ListSecretsResponse
	path := "/api/2.0/secrets/list"
	err := a.client.Get(ctx, path, request, &listSecretsResponse)
	return &listSecretsResponse, err
}

// Creates or overwrites the ACL associated with the given principal (user or
// group) on the specified scope point. In general, a user or group will use the
// most powerful permission available to them, and permissions are ordered as
// follows: * ``MANAGE`` - Allowed to change ACLs, and read and write to this
// secret scope. * ``WRITE`` - Allowed to read and write to this secret scope. *
// ``READ`` - Allowed to read this secret scope and list what secrets are
// available. Note that in general, secret values can only be read from within a
// command on a cluster (for example, through a notebook). There is no API to
// read the actual secret value material outside of a cluster. However, the
// user&#39;s permission will be applied based on who is executing the command, and
// they must have at least READ permission. Users must have the ``MANAGE``
// permission to invoke this API. Example request: .. code:: { &#34;scope&#34;:
// &#34;my-secret-scope&#34;, &#34;principal&#34;: &#34;data-scientists&#34;, &#34;permission&#34;: &#34;READ&#34; } The
// principal is a user or group name corresponding to an existing Databricks
// principal to be granted or revoked access. Throws ``RESOURCE_DOES_NOT_EXIST``
// if no such secret scope exists. Throws ``RESOURCE_ALREADY_EXISTS`` if a
// permission for the principal already exists. Throws
// ``INVALID_PARAMETER_VALUE`` if the permission is invalid. Throws
// ``PERMISSION_DENIED`` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) PutAcl(ctx context.Context, request PutAclRequest) error {
	path := "/api/2.0/secrets/acls/put"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Inserts a secret under the provided scope with the given name. If a secret
// already exists with the same name, this command overwrites the existing
// secret&#39;s value. The server encrypts the secret using the secret scope&#39;s
// encryption settings before storing it. You must have ``WRITE`` or ``MANAGE``
// permission on the secret scope. The secret key must consist of alphanumeric
// characters, dashes, underscores, and periods, and cannot exceed 128
// characters. The maximum allowed secret value size is 128 KB. The maximum
// number of secrets in a given scope is 1000. Example request: .. code:: {
// &#34;scope&#34;: &#34;my-databricks-scope&#34;, &#34;key&#34;: &#34;my-string-key&#34;, &#34;string_value&#34;:
// &#34;foobar&#34; } The input fields &#34;string_value&#34; or &#34;bytes_value&#34; specify the type
// of the secret, which will determine the value returned when the secret value
// is requested. Exactly one must be specified. Throws
// ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
// ``RESOURCE_LIMIT_EXCEEDED`` if maximum number of secrets in scope is
// exceeded. Throws ``INVALID_PARAMETER_VALUE`` if the key name or value length
// is invalid. Throws ``PERMISSION_DENIED`` if the user does not have permission
// to make this API call.
func (a *SecretsAPI) PutSecret(ctx context.Context, request PutSecretRequest) error {
	path := "/api/2.0/secrets/put"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

