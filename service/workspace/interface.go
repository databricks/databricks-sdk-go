// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"
)

// Registers personal access token for Databricks to do operations on behalf of
// the user.
//
// See [more info].
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [more info]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
type GitCredentialsService interface {

	// Creates a Git credential entry for the user. Only one Git credential per
	// user is supported, so any attempts to create credentials if an entry
	// already exists will fail. Use the PATCH endpoint to update existing
	// credentials, or the DELETE endpoint to delete existing credentials.
	Create(ctx context.Context, request CreateCredentialsRequest) (*CreateCredentialsResponse, error)

	// Deletes the specified Git credential.
	Delete(ctx context.Context, request DeleteCredentialsRequest) error

	// Gets the Git credential with the specified credential ID.
	Get(ctx context.Context, request GetCredentialsRequest) (*GetCredentialsResponse, error)

	// Lists the calling user's Git credentials. One credential per user is
	// supported.
	List(ctx context.Context) (*ListCredentialsResponse, error)

	// Updates the specified Git credential.
	Update(ctx context.Context, request UpdateCredentialsRequest) error
}

// The Repos API allows users to manage their git repos. Users can use the API
// to access all repos that they have manage permissions on.
//
// Databricks Repos is a visual Git client in Databricks. It supports common Git
// operations such a cloning a repository, committing and pushing, pulling,
// branch management, and visual comparison of diffs when committing.
//
// Within Repos you can develop code in notebooks or other files and follow data
// science and engineering code development best practices using Git for version
// control, collaboration, and CI/CD.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ReposService interface {

	// Creates a repo in the workspace and links it to the remote Git repo
	// specified. Note that repos created programmatically must be linked to a
	// remote Git repo, unlike repos created in the browser.
	Create(ctx context.Context, request CreateRepoRequest) (*CreateRepoResponse, error)

	// Deletes the specified repo.
	Delete(ctx context.Context, request DeleteRepoRequest) error

	// Returns the repo with the given repo ID.
	Get(ctx context.Context, request GetRepoRequest) (*GetRepoResponse, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetRepoPermissionLevelsRequest) (*GetRepoPermissionLevelsResponse, error)

	// Gets the permissions of a repo. Repos can inherit permissions from their
	// root object.
	GetPermissions(ctx context.Context, request GetRepoPermissionsRequest) (*RepoPermissions, error)

	// Returns repos that the calling user has Manage permissions on. Use
	// `next_page_token` to iterate through additional pages.
	List(ctx context.Context, request ListReposRequest) (*ListReposResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error)

	// Updates the repo to a different branch or tag, or updates the repo to the
	// latest commit on the same branch.
	Update(ctx context.Context, request UpdateRepoRequest) error

	// Updates the permissions on a repo. Repos can inherit permissions from
	// their root object.
	UpdatePermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error)
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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SecretsService interface {

	// Creates a new secret scope.
	//
	// The scope name must consist of alphanumeric characters, dashes,
	// underscores, and periods, and may not exceed 128 characters.
	//
	// Example request:
	//
	// .. code::
	//
	// { "scope": "my-simple-databricks-scope", "initial_manage_principal":
	// "users" "scope_backend_type": "databricks|azure_keyvault", # below is
	// only required if scope type is azure_keyvault "backend_azure_keyvault": {
	// "resource_id":
	// "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/xxxx/providers/Microsoft.KeyVault/vaults/xxxx",
	// "tenant_id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", "dns_name":
	// "https://xxxx.vault.azure.net/", } }
	//
	// If ``initial_manage_principal`` is specified, the initial ACL applied to
	// the scope is applied to the supplied principal (user or group) with
	// ``MANAGE`` permissions. The only supported principal for this option is
	// the group ``users``, which contains all users in the workspace. If
	// ``initial_manage_principal`` is not specified, the initial ACL with
	// ``MANAGE`` permission applied to the scope is assigned to the API request
	// issuer's user identity.
	//
	// If ``scope_backend_type`` is ``azure_keyvault``, a secret scope is
	// created with secrets from a given Azure KeyVault. The caller must provide
	// the keyvault_resource_id and the tenant_id for the key vault. If
	// ``scope_backend_type`` is ``databricks`` or is unspecified, an empty
	// secret scope is created and stored in Databricks's own storage.
	//
	// Throws ``RESOURCE_ALREADY_EXISTS`` if a scope with the given name already
	// exists. Throws ``RESOURCE_LIMIT_EXCEEDED`` if maximum number of scopes in
	// the workspace is exceeded. Throws ``INVALID_PARAMETER_VALUE`` if the
	// scope name is invalid. Throws ``BAD_REQUEST`` if request violated
	// constraints. Throws ``CUSTOMER_UNAUTHORIZED`` if normal user attempts to
	// create a scope with name reserved for databricks internal usage. Throws
	// ``UNAUTHENTICATED`` if unable to verify user access permission on Azure
	// KeyVault
	CreateScope(ctx context.Context, request CreateScope) error

	// Deletes the given ACL on the given scope.
	//
	// Users must have the ``MANAGE`` permission to invoke this API.
	//
	// Example request:
	//
	// .. code::
	//
	// { "scope": "my-secret-scope", "principal": "data-scientists" }
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope, principal, or
	// ACL exists. Throws ``PERMISSION_DENIED`` if the user does not have
	// permission to make this API call. Throws ``INVALID_PARAMETER_VALUE`` if
	// the permission or principal is invalid.
	DeleteAcl(ctx context.Context, request DeleteAcl) error

	// Deletes a secret scope.
	//
	// Example request:
	//
	// .. code::
	//
	// { "scope": "my-secret-scope" }
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if the scope does not exist. Throws
	// ``PERMISSION_DENIED`` if the user does not have permission to make this
	// API call. Throws ``BAD_REQUEST`` if system user attempts to delete
	// internal secret scope.
	DeleteScope(ctx context.Context, request DeleteScope) error

	// Deletes the secret stored in this secret scope. You must have ``WRITE``
	// or ``MANAGE`` permission on the Secret Scope.
	//
	// Example request:
	//
	// .. code::
	//
	// { "scope": "my-secret-scope", "key": "my-secret-key" }
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope or secret
	// exists. Throws ``PERMISSION_DENIED`` if the user does not have permission
	// to make this API call. Throws ``BAD_REQUEST`` if system user attempts to
	// delete an internal secret, or request is made against Azure KeyVault
	// backed scope.
	DeleteSecret(ctx context.Context, request DeleteSecret) error

	// Describes the details about the given ACL, such as the group and
	// permission.
	//
	// Users must have the ``MANAGE`` permission to invoke this API.
	//
	// Example response:
	//
	// .. code::
	//
	// { "principal": "data-scientists", "permission": "READ" }
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
	// ``PERMISSION_DENIED`` if the user does not have permission to make this
	// API call. Throws ``INVALID_PARAMETER_VALUE`` if the permission or
	// principal is invalid.
	GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error)

	// Gets a secret for a given key and scope. This API can only be called from
	// the DBUtils interface. Users need the READ permission to make this call.
	//
	// Example response:
	//
	// .. code::
	//
	// { "key": "my-string-key", "value": <bytes of the secret value> }
	//
	// Note that the secret value returned is in bytes. The interpretation of
	// the bytes is determined by the caller in DBUtils and the type the data is
	// decoded into.
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret or secret scope
	// exists. Throws ``PERMISSION_DENIED`` if the user does not have permission
	// to make this API call.
	//
	// Note: This is explicitly an undocumented API. It also doesn't need to be
	// supported for the /preview prefix, because it's not a customer-facing API
	// (i.e. only used for DBUtils SecretUtils to fetch secrets).
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope or secret
	// exists. Throws ``BAD_REQUEST`` if normal user calls get secret outside of
	// a notebook. AKV specific errors: Throws ``INVALID_PARAMETER_VALUE`` if
	// secret name is not alphanumeric or too long. Throws ``PERMISSION_DENIED``
	// if secret manager cannot access AKV with 403 error Throws
	// ``MALFORMED_REQUEST`` if secret manager cannot access AKV with any other
	// 4xx error
	GetSecret(ctx context.Context, request GetSecretRequest) (*GetSecretResponse, error)

	// Lists the ACLs set on the given scope.
	//
	// Users must have the ``MANAGE`` permission to invoke this API.
	//
	// Example response:
	//
	// .. code::
	//
	// { "acls": [{ "principal": "admins", "permission": "MANAGE" },{
	// "principal": "data-scientists", "permission": "READ" }] }
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
	// ``PERMISSION_DENIED`` if the user does not have permission to make this
	// API call.
	ListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error)

	// Lists all secret scopes available in the workspace.
	//
	// Example response:
	//
	// .. code::
	//
	// { "scopes": [{ "name": "my-databricks-scope", "backend_type":
	// "DATABRICKS" },{ "name": "mount-points", "backend_type": "DATABRICKS" }]
	// }
	//
	// Throws ``PERMISSION_DENIED`` if the user does not have permission to make
	// this API call.
	ListScopes(ctx context.Context) (*ListScopesResponse, error)

	// Lists the secret keys that are stored at this scope. This is a
	// metadata-only operation; secret data cannot be retrieved using this API.
	// Users need the READ permission to make this call.
	//
	// Example response:
	//
	// .. code::
	//
	// { "secrets": [ { "key": "my-string-key"", "last_updated_timestamp":
	// "1520467595000" }, { "key": "my-byte-key", "last_updated_timestamp":
	// "1520467595000" }, ] }
	//
	// The lastUpdatedTimestamp returned is in milliseconds since epoch.
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
	// ``PERMISSION_DENIED`` if the user does not have permission to make this
	// API call.
	ListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error)

	// Creates or overwrites the ACL associated with the given principal (user
	// or group) on the specified scope point. In general, a user or group will
	// use the most powerful permission available to them, and permissions are
	// ordered as follows:
	//
	// * ``MANAGE`` - Allowed to change ACLs, and read and write to this secret
	// scope. * ``WRITE`` - Allowed to read and write to this secret scope. *
	// ``READ`` - Allowed to read this secret scope and list what secrets are
	// available.
	//
	// Note that in general, secret values can only be read from within a
	// command on a cluster (for example, through a notebook). There is no API
	// to read the actual secret value material outside of a cluster. However,
	// the user's permission will be applied based on who is executing the
	// command, and they must have at least READ permission.
	//
	// Users must have the ``MANAGE`` permission to invoke this API.
	//
	// Example request:
	//
	// .. code::
	//
	// { "scope": "my-secret-scope", "principal": "data-scientists",
	// "permission": "READ" }
	//
	// The principal is a user or group name corresponding to an existing
	// Databricks principal to be granted or revoked access.
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
	// ``RESOURCE_ALREADY_EXISTS`` if a permission for the principal already
	// exists. Throws ``INVALID_PARAMETER_VALUE`` if the permission or principal
	// is invalid. Throws ``PERMISSION_DENIED`` if the user does not have
	// permission to make this API call.
	PutAcl(ctx context.Context, request PutAcl) error

	// Inserts a secret under the provided scope with the given name. If a
	// secret already exists with the same name, this command overwrites the
	// existing secret's value. The server encrypts the secret using the secret
	// scope's encryption settings before storing it. You must have ``WRITE`` or
	// ``MANAGE`` permission on the secret scope.
	//
	// The secret key must consist of alphanumeric characters, dashes,
	// underscores, and periods, and cannot exceed 128 characters. The maximum
	// allowed secret value size is 128 KB. The maximum number of secrets in a
	// given scope is 1000.
	//
	// Example request:
	//
	// .. code::
	//
	// { "scope": "my-databricks-scope", "key": "my-string-key", "string_value":
	// "foobar" }
	//
	// The input fields "string_value" or "bytes_value" specify the type of the
	// secret, which will determine the value returned when the secret value is
	// requested. Exactly one must be specified.
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no such secret scope exists. Throws
	// ``RESOURCE_LIMIT_EXCEEDED`` if maximum number of secrets in scope is
	// exceeded. Throws ``INVALID_PARAMETER_VALUE`` if the request parameters
	// are invalid. Throws ``PERMISSION_DENIED`` if the user does not have
	// permission to make this API call. Throws ``MALFORMED_REQUEST`` if request
	// is incorrectly formatted or conflicting. Throws ``BAD_REQUEST`` if
	// request is made against Azure KeyVault backed scope.
	PutSecret(ctx context.Context, request PutSecret) error
}

// The Workspace API allows you to list, import, export, and delete notebooks
// and folders.
//
// A notebook is a web-based interface to a document that contains runnable
// code, visualizations, and explanatory text.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceService interface {

	// Deletes an object or a directory (and optionally recursively deletes all
	// objects in the directory). * If `path` does not exist, this call returns
	// an error `RESOURCE_DOES_NOT_EXIST`. * If `path` is a non-empty directory
	// and `recursive` is set to `false`, this call returns an error
	// `DIRECTORY_NOT_EMPTY`.
	//
	// Object deletion cannot be undone and deleting a directory recursively is
	// not atomic.
	Delete(ctx context.Context, request Delete) error

	// Exports an object or the contents of an entire directory.
	//
	// If `path` does not exist, this call returns an error
	// `RESOURCE_DOES_NOT_EXIST`.
	//
	// If the exported data would exceed size limit, this call returns
	// `MAX_NOTEBOOK_SIZE_EXCEEDED`. Currently, this API does not support
	// exporting a library.
	Export(ctx context.Context, request ExportRequest) (*ExportResponse, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetWorkspaceObjectPermissionLevelsRequest) (*GetWorkspaceObjectPermissionLevelsResponse, error)

	// Gets the permissions of a workspace object. Workspace objects can inherit
	// permissions from their parent objects or root object.
	GetPermissions(ctx context.Context, request GetWorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error)

	// Gets the status of an object or a directory. If `path` does not exist,
	// this call returns an error `RESOURCE_DOES_NOT_EXIST`.
	GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error)

	// Imports a workspace object (for example, a notebook or file) or the
	// contents of an entire directory. If `path` already exists and `overwrite`
	// is set to `false`, this call returns an error `RESOURCE_ALREADY_EXISTS`.
	// To import a directory, you can use either the `DBC` format or the
	// `SOURCE` format with the `language` field unset. To import a single file
	// as `SOURCE`, you must set the `language` field.
	Import(ctx context.Context, request Import) error

	// Lists the contents of a directory, or the object if it is not a
	// directory. If the input path does not exist, this call returns an error
	// `RESOURCE_DOES_NOT_EXIST`.
	List(ctx context.Context, request ListWorkspaceRequest) (*ListResponse, error)

	// Creates the specified directory (and necessary parent directories if they
	// do not exist). If there is an object (not a directory) at any prefix of
	// the input path, this call returns an error `RESOURCE_ALREADY_EXISTS`.
	//
	// Note that if this operation fails it may have succeeded in creating some
	// of the necessary parent directories.
	Mkdirs(ctx context.Context, request Mkdirs) error

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their parent objects or root object.
	SetPermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error)

	// Updates the permissions on a workspace object. Workspace objects can
	// inherit permissions from their parent objects or root object.
	UpdatePermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error)
}
