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
// [more info]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
type GitCredentialsService interface {

	// Create a credential entry.
	//
	// Creates a Git credential entry for the user. Only one Git credential per
	// user is supported, so any attempts to create credentials if an entry
	// already exists will fail. Use the PATCH endpoint to update existing
	// credentials, or the DELETE endpoint to delete existing credentials.
	Create(ctx context.Context, request CreateCredentials) (*CreateCredentialsResponse, error)

	// Delete a credential.
	//
	// Deletes the specified Git credential.
	Delete(ctx context.Context, request DeleteGitCredentialRequest) error

	// Get a credential entry.
	//
	// Gets the Git credential with the specified credential ID.
	Get(ctx context.Context, request GetGitCredentialRequest) (*CredentialInfo, error)

	// Get Git credentials.
	//
	// Lists the calling user's Git credentials. One credential per user is
	// supported.
	//
	// Use ListAll() to get all CredentialInfo instances
	List(ctx context.Context) (*GetCredentialsResponse, error)

	// Update a credential.
	//
	// Updates the specified Git credential.
	Update(ctx context.Context, request UpdateCredentials) error
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
type ReposService interface {

	// Create a repo.
	//
	// Creates a repo in the workspace and links it to the remote Git repo
	// specified. Note that repos created programmatically must be linked to a
	// remote Git repo, unlike repos created in the browser.
	Create(ctx context.Context, request CreateRepo) (*RepoInfo, error)

	// Delete a repo.
	//
	// Deletes the specified repo.
	Delete(ctx context.Context, request DeleteRepoRequest) error

	// Get a repo.
	//
	// Returns the repo with the given repo ID.
	Get(ctx context.Context, request GetRepoRequest) (*RepoInfo, error)

	// Get repos.
	//
	// Returns repos that the calling user has Manage permissions on. Results
	// are paginated with each page containing twenty repos.
	//
	// Use ListAll() to get all RepoInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListReposRequest) (*ListReposResponse, error)

	// Update a repo.
	//
	// Updates the repo to a different branch or tag, or updates the repo to the
	// latest commit on the same branch.
	Update(ctx context.Context, request UpdateRepo) error
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
type SecretsService interface {

	// Create a new secret scope.
	//
	// The scope name must consist of alphanumeric characters, dashes,
	// underscores, and periods, and may not exceed 128 characters. The maximum
	// number of scopes in a workspace is 100.
	CreateScope(ctx context.Context, request CreateScope) error

	// Delete an ACL.
	//
	// Deletes the given ACL on the given scope.
	//
	// Users must have the `MANAGE` permission to invoke this API. Throws
	// `RESOURCE_DOES_NOT_EXIST` if no such secret scope, principal, or ACL
	// exists. Throws `PERMISSION_DENIED` if the user does not have permission
	// to make this API call.
	DeleteAcl(ctx context.Context, request DeleteAcl) error

	// Delete a secret scope.
	//
	// Deletes a secret scope.
	//
	// Throws `RESOURCE_DOES_NOT_EXIST` if the scope does not exist. Throws
	// `PERMISSION_DENIED` if the user does not have permission to make this API
	// call.
	DeleteScope(ctx context.Context, request DeleteScope) error

	// Delete a secret.
	//
	// Deletes the secret stored in this secret scope. You must have `WRITE` or
	// `MANAGE` permission on the secret scope.
	//
	// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope or secret
	// exists. Throws `PERMISSION_DENIED` if the user does not have permission
	// to make this API call.
	DeleteSecret(ctx context.Context, request DeleteSecret) error

	// Get secret ACL details.
	//
	// Gets the details about the given ACL, such as the group and permission.
	// Users must have the `MANAGE` permission to invoke this API.
	//
	// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
	// `PERMISSION_DENIED` if the user does not have permission to make this API
	// call.
	GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error)

	// Lists ACLs.
	//
	// List the ACLs for a given secret scope. Users must have the `MANAGE`
	// permission to invoke this API.
	//
	// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
	// `PERMISSION_DENIED` if the user does not have permission to make this API
	// call.
	//
	// Use ListAclsAll() to get all AclItem instances
	ListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error)

	// List all scopes.
	//
	// Lists all secret scopes available in the workspace.
	//
	// Throws `PERMISSION_DENIED` if the user does not have permission to make
	// this API call.
	//
	// Use ListScopesAll() to get all SecretScope instances
	ListScopes(ctx context.Context) (*ListScopesResponse, error)

	// List secret keys.
	//
	// Lists the secret keys that are stored at this scope. This is a
	// metadata-only operation; secret data cannot be retrieved using this API.
	// Users need the READ permission to make this call.
	//
	// The lastUpdatedTimestamp returned is in milliseconds since epoch. Throws
	// `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
	// `PERMISSION_DENIED` if the user does not have permission to make this API
	// call.
	//
	// Use ListSecretsAll() to get all SecretMetadata instances
	ListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error)

	// Create/update an ACL.
	//
	// Creates or overwrites the Access Control List (ACL) associated with the
	// given principal (user or group) on the specified scope point.
	//
	// In general, a user or group will use the most powerful permission
	// available to them, and permissions are ordered as follows:
	//
	// * `MANAGE` - Allowed to change ACLs, and read and write to this secret
	// scope. * `WRITE` - Allowed to read and write to this secret scope. *
	// `READ` - Allowed to read this secret scope and list what secrets are
	// available.
	//
	// Note that in general, secret values can only be read from within a
	// command on a cluster (for example, through a notebook). There is no API
	// to read the actual secret value material outside of a cluster. However,
	// the user's permission will be applied based on who is executing the
	// command, and they must have at least READ permission.
	//
	// Users must have the `MANAGE` permission to invoke this API.
	//
	// The principal is a user or group name corresponding to an existing
	// Databricks principal to be granted or revoked access.
	//
	// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
	// `RESOURCE_ALREADY_EXISTS` if a permission for the principal already
	// exists. Throws `INVALID_PARAMETER_VALUE` if the permission is invalid.
	// Throws `PERMISSION_DENIED` if the user does not have permission to make
	// this API call.
	PutAcl(ctx context.Context, request PutAcl) error

	// Add a secret.
	//
	// Inserts a secret under the provided scope with the given name. If a
	// secret already exists with the same name, this command overwrites the
	// existing secret's value. The server encrypts the secret using the secret
	// scope's encryption settings before storing it.
	//
	// You must have `WRITE` or `MANAGE` permission on the secret scope. The
	// secret key must consist of alphanumeric characters, dashes, underscores,
	// and periods, and cannot exceed 128 characters. The maximum allowed secret
	// value size is 128 KB. The maximum number of secrets in a given scope is
	// 1000.
	//
	// The input fields "string_value" or "bytes_value" specify the type of the
	// secret, which will determine the value returned when the secret value is
	// requested. Exactly one must be specified.
	//
	// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
	// `RESOURCE_LIMIT_EXCEEDED` if maximum number of secrets in scope is
	// exceeded. Throws `INVALID_PARAMETER_VALUE` if the key name or value
	// length is invalid. Throws `PERMISSION_DENIED` if the user does not have
	// permission to make this API call.
	PutSecret(ctx context.Context, request PutSecret) error
}

// The Workspace API allows you to list, import, export, and delete notebooks
// and folders.
//
// A notebook is a web-based interface to a document that contains runnable
// code, visualizations, and explanatory text.
type WorkspaceService interface {

	// Delete a workspace object.
	//
	// Deletes an object or a directory (and optionally recursively deletes all
	// objects in the directory). * If `path` does not exist, this call returns
	// an error `RESOURCE_DOES_NOT_EXIST`. * If `path` is a non-empty directory
	// and `recursive` is set to `false`, this call returns an error
	// `DIRECTORY_NOT_EMPTY`.
	//
	// Object deletion cannot be undone and deleting a directory recursively is
	// not atomic.
	Delete(ctx context.Context, request Delete) error

	// Export a workspace object.
	//
	// Exports an object or the contents of an entire directory.
	//
	// If `path` does not exist, this call returns an error
	// `RESOURCE_DOES_NOT_EXIST`.
	//
	// If the exported data would exceed size limit, this call returns
	// `MAX_NOTEBOOK_SIZE_EXCEEDED`. Currently, this API does not support
	// exporting a library.
	Export(ctx context.Context, request ExportRequest) (*ExportResponse, error)

	// Get status.
	//
	// Gets the status of an object or a directory. If `path` does not exist,
	// this call returns an error `RESOURCE_DOES_NOT_EXIST`.
	GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error)

	// Import a workspace object.
	//
	// Imports a workspace object (for example, a notebook or file) or the
	// contents of an entire directory. If `path` already exists and `overwrite`
	// is set to `false`, this call returns an error `RESOURCE_ALREADY_EXISTS`.
	// One can only use `DBC` format to import a directory.
	Import(ctx context.Context, request Import) error

	// List contents.
	//
	// Lists the contents of a directory, or the object if it is not a
	// directory. If the input path does not exist, this call returns an error
	// `RESOURCE_DOES_NOT_EXIST`.
	//
	// Use ListAll() to get all ObjectInfo instances
	List(ctx context.Context, request ListWorkspaceRequest) (*ListResponse, error)

	// Create a directory.
	//
	// Creates the specified directory (and necessary parent directories if they
	// do not exist). If there is an object (not a directory) at any prefix of
	// the input path, this call returns an error `RESOURCE_ALREADY_EXISTS`.
	//
	// Note that if this operation fails it may have succeeded in creating some
	// of the necessary parent directories.
	Mkdirs(ctx context.Context, request Mkdirs) error
}
