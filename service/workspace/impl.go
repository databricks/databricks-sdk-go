// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just GitCredentials API methods
type gitCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *gitCredentialsImpl) Create(ctx context.Context, request CreateCredentialsRequest) (*CreateCredentialsResponse, error) {
	var createCredentialsResponse CreateCredentialsResponse
	path := "/api/2.0/git-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createCredentialsResponse)
	return &createCredentialsResponse, err
}

func (a *gitCredentialsImpl) Delete(ctx context.Context, request DeleteCredentialsRequest) error {
	var deleteCredentialsResponse DeleteCredentialsResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteCredentialsResponse)
	return err
}

func (a *gitCredentialsImpl) Get(ctx context.Context, request GetCredentialsRequest) (*GetCredentialsResponse, error) {
	var getCredentialsResponse GetCredentialsResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getCredentialsResponse)
	return &getCredentialsResponse, err
}

// Get Git credentials.
//
// Lists the calling user's Git credentials. One credential per user is
// supported.
func (a *gitCredentialsImpl) List(ctx context.Context) listing.Iterator[CredentialInfo] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListCredentialsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListCredentialsResponse) []CredentialInfo {
		return resp.Credentials
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get Git credentials.
//
// Lists the calling user's Git credentials. One credential per user is
// supported.
func (a *gitCredentialsImpl) ListAll(ctx context.Context) ([]CredentialInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[CredentialInfo](ctx, iterator)
}

func (a *gitCredentialsImpl) internalList(ctx context.Context) (*ListCredentialsResponse, error) {
	var listCredentialsResponse ListCredentialsResponse
	path := "/api/2.0/git-credentials"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listCredentialsResponse)
	return &listCredentialsResponse, err
}

func (a *gitCredentialsImpl) Update(ctx context.Context, request UpdateCredentialsRequest) error {
	var updateCredentialsResponse UpdateCredentialsResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateCredentialsResponse)
	return err
}

// unexported type that holds implementations of just Repos API methods
type reposImpl struct {
	client *client.DatabricksClient
}

func (a *reposImpl) Create(ctx context.Context, request CreateRepoRequest) (*CreateRepoResponse, error) {
	var createRepoResponse CreateRepoResponse
	path := "/api/2.0/repos"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createRepoResponse)
	return &createRepoResponse, err
}

func (a *reposImpl) Delete(ctx context.Context, request DeleteRepoRequest) error {
	var deleteRepoResponse DeleteRepoResponse
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteRepoResponse)
	return err
}

func (a *reposImpl) Get(ctx context.Context, request GetRepoRequest) (*GetRepoResponse, error) {
	var getRepoResponse GetRepoResponse
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRepoResponse)
	return &getRepoResponse, err
}

func (a *reposImpl) GetPermissionLevels(ctx context.Context, request GetRepoPermissionLevelsRequest) (*GetRepoPermissionLevelsResponse, error) {
	var getRepoPermissionLevelsResponse GetRepoPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v/permissionLevels", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRepoPermissionLevelsResponse)
	return &getRepoPermissionLevelsResponse, err
}

func (a *reposImpl) GetPermissions(ctx context.Context, request GetRepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &repoPermissions)
	return &repoPermissions, err
}

// Get repos.
//
// Returns repos that the calling user has Manage permissions on. Use
// `next_page_token` to iterate through additional pages.
func (a *reposImpl) List(ctx context.Context, request ListReposRequest) listing.Iterator[RepoInfo] {

	getNextPage := func(ctx context.Context, req ListReposRequest) (*ListReposResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListReposResponse) []RepoInfo {
		return resp.Repos
	}
	getNextReq := func(resp *ListReposResponse) *ListReposRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.NextPageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Get repos.
//
// Returns repos that the calling user has Manage permissions on. Use
// `next_page_token` to iterate through additional pages.
func (a *reposImpl) ListAll(ctx context.Context, request ListReposRequest) ([]RepoInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[RepoInfo](ctx, iterator)
}

func (a *reposImpl) internalList(ctx context.Context, request ListReposRequest) (*ListReposResponse, error) {
	var listReposResponse ListReposResponse
	path := "/api/2.0/repos"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listReposResponse)
	return &listReposResponse, err
}

func (a *reposImpl) SetPermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &repoPermissions)
	return &repoPermissions, err
}

func (a *reposImpl) Update(ctx context.Context, request UpdateRepoRequest) error {
	var updateRepoResponse UpdateRepoResponse
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateRepoResponse)
	return err
}

func (a *reposImpl) UpdatePermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &repoPermissions)
	return &repoPermissions, err
}

// unexported type that holds implementations of just Secrets API methods
type secretsImpl struct {
	client *client.DatabricksClient
}

func (a *secretsImpl) CreateScope(ctx context.Context, request CreateScope) error {
	var createScopeResponse CreateScopeResponse
	path := "/api/2.0/secrets/scopes/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createScopeResponse)
	return err
}

func (a *secretsImpl) DeleteAcl(ctx context.Context, request DeleteAcl) error {
	var deleteAclResponse DeleteAclResponse
	path := "/api/2.0/secrets/acls/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteAclResponse)
	return err
}

func (a *secretsImpl) DeleteScope(ctx context.Context, request DeleteScope) error {
	var deleteScopeResponse DeleteScopeResponse
	path := "/api/2.0/secrets/scopes/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteScopeResponse)
	return err
}

func (a *secretsImpl) DeleteSecret(ctx context.Context, request DeleteSecret) error {
	var deleteSecretResponse DeleteSecretResponse
	path := "/api/2.0/secrets/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteSecretResponse)
	return err
}

func (a *secretsImpl) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {
	var aclItem AclItem
	path := "/api/2.0/secrets/acls/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &aclItem)
	return &aclItem, err
}

func (a *secretsImpl) GetSecret(ctx context.Context, request GetSecretRequest) (*GetSecretResponse, error) {
	var getSecretResponse GetSecretResponse
	path := "/api/2.0/secrets/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getSecretResponse)
	return &getSecretResponse, err
}

// Lists ACLs.
//
// List the ACLs for a given secret scope. Users must have the `MANAGE`
// permission to invoke this API.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListAcls(ctx context.Context, request ListAclsRequest) listing.Iterator[AclItem] {

	getNextPage := func(ctx context.Context, req ListAclsRequest) (*ListAclsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListAcls(ctx, req)
	}
	getItems := func(resp *ListAclsResponse) []AclItem {
		return resp.Items
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Lists ACLs.
//
// List the ACLs for a given secret scope. Users must have the `MANAGE`
// permission to invoke this API.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListAclsAll(ctx context.Context, request ListAclsRequest) ([]AclItem, error) {
	iterator := a.ListAcls(ctx, request)
	return listing.ToSlice[AclItem](ctx, iterator)
}

func (a *secretsImpl) internalListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error) {
	var listAclsResponse ListAclsResponse
	path := "/api/2.0/secrets/acls/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAclsResponse)
	return &listAclsResponse, err
}

// List all scopes.
//
// Lists all secret scopes available in the workspace.
//
// Throws `PERMISSION_DENIED` if the user does not have permission to make this
// API call.
func (a *secretsImpl) ListScopes(ctx context.Context) listing.Iterator[SecretScope] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListScopesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListScopes(ctx)
	}
	getItems := func(resp *ListScopesResponse) []SecretScope {
		return resp.Scopes
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List all scopes.
//
// Lists all secret scopes available in the workspace.
//
// Throws `PERMISSION_DENIED` if the user does not have permission to make this
// API call.
func (a *secretsImpl) ListScopesAll(ctx context.Context) ([]SecretScope, error) {
	iterator := a.ListScopes(ctx)
	return listing.ToSlice[SecretScope](ctx, iterator)
}

func (a *secretsImpl) internalListScopes(ctx context.Context) (*ListScopesResponse, error) {
	var listScopesResponse ListScopesResponse
	path := "/api/2.0/secrets/scopes/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listScopesResponse)
	return &listScopesResponse, err
}

// List secret keys.
//
// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
//
// The lastUpdatedTimestamp returned is in milliseconds since epoch. Throws
// `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListSecrets(ctx context.Context, request ListSecretsRequest) listing.Iterator[SecretMetadata] {

	getNextPage := func(ctx context.Context, req ListSecretsRequest) (*ListSecretsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSecrets(ctx, req)
	}
	getItems := func(resp *ListSecretsResponse) []SecretMetadata {
		return resp.Secrets
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List secret keys.
//
// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
//
// The lastUpdatedTimestamp returned is in milliseconds since epoch. Throws
// `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListSecretsAll(ctx context.Context, request ListSecretsRequest) ([]SecretMetadata, error) {
	iterator := a.ListSecrets(ctx, request)
	return listing.ToSlice[SecretMetadata](ctx, iterator)
}

func (a *secretsImpl) internalListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error) {
	var listSecretsResponse ListSecretsResponse
	path := "/api/2.0/secrets/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSecretsResponse)
	return &listSecretsResponse, err
}

func (a *secretsImpl) PutAcl(ctx context.Context, request PutAcl) error {
	var putAclResponse PutAclResponse
	path := "/api/2.0/secrets/acls/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &putAclResponse)
	return err
}

func (a *secretsImpl) PutSecret(ctx context.Context, request PutSecret) error {
	var putSecretResponse PutSecretResponse
	path := "/api/2.0/secrets/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &putSecretResponse)
	return err
}

// unexported type that holds implementations of just Workspace API methods
type workspaceImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceImpl) Delete(ctx context.Context, request Delete) error {
	var deleteResponse DeleteResponse
	path := "/api/2.0/workspace/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *workspaceImpl) Export(ctx context.Context, request ExportRequest) (*ExportResponse, error) {
	var exportResponse ExportResponse
	path := "/api/2.0/workspace/export"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &exportResponse)
	return &exportResponse, err
}

func (a *workspaceImpl) GetPermissionLevels(ctx context.Context, request GetWorkspaceObjectPermissionLevelsRequest) (*GetWorkspaceObjectPermissionLevelsResponse, error) {
	var getWorkspaceObjectPermissionLevelsResponse GetWorkspaceObjectPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWorkspaceObjectPermissionLevelsResponse)
	return &getWorkspaceObjectPermissionLevelsResponse, err
}

func (a *workspaceImpl) GetPermissions(ctx context.Context, request GetWorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}

func (a *workspaceImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error) {
	var objectInfo ObjectInfo
	path := "/api/2.0/workspace/get-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &objectInfo)
	return &objectInfo, err
}

func (a *workspaceImpl) Import(ctx context.Context, request Import) error {
	var importResponse ImportResponse
	path := "/api/2.0/workspace/import"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &importResponse)
	return err
}

// List contents.
//
// Lists the contents of a directory, or the object if it is not a directory. If
// the input path does not exist, this call returns an error
// `RESOURCE_DOES_NOT_EXIST`.
func (a *workspaceImpl) List(ctx context.Context, request ListWorkspaceRequest) listing.Iterator[ObjectInfo] {

	getNextPage := func(ctx context.Context, req ListWorkspaceRequest) (*ListResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListResponse) []ObjectInfo {
		return resp.Objects
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List contents.
//
// Lists the contents of a directory, or the object if it is not a directory. If
// the input path does not exist, this call returns an error
// `RESOURCE_DOES_NOT_EXIST`.
func (a *workspaceImpl) ListAll(ctx context.Context, request ListWorkspaceRequest) ([]ObjectInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ObjectInfo](ctx, iterator)
}

func (a *workspaceImpl) internalList(ctx context.Context, request ListWorkspaceRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/workspace/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listResponse)
	return &listResponse, err
}

func (a *workspaceImpl) Mkdirs(ctx context.Context, request Mkdirs) error {
	var mkdirsResponse MkdirsResponse
	path := "/api/2.0/workspace/mkdirs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &mkdirsResponse)
	return err
}

func (a *workspaceImpl) SetPermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}

func (a *workspaceImpl) UpdatePermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}
