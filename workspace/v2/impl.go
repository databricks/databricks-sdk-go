// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just GitCredentials API methods
type gitCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *gitCredentialsImpl) Create(ctx context.Context, request CreateCredentialsRequest) (*CreateCredentialsResponse, error) {
	var createCredentialsResponse CreateCredentialsResponse
	path := "/api/2.0/git-credentials"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createCredentialsResponse)
	return &createCredentialsResponse, err
}

func (a *gitCredentialsImpl) Delete(ctx context.Context, request DeleteCredentialsRequest) error {
	var deleteCredentialsResponse DeleteCredentialsResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteCredentialsResponse)
	return err
}

func (a *gitCredentialsImpl) Get(ctx context.Context, request GetCredentialsRequest) (*GetCredentialsResponse, error) {
	var getCredentialsResponse GetCredentialsResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getCredentialsResponse)
	return &getCredentialsResponse, err
}

func (a *gitCredentialsImpl) List(ctx context.Context) (*ListCredentialsResponse, error) {
	var listCredentialsResponse ListCredentialsResponse
	path := "/api/2.0/git-credentials"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listCredentialsResponse)
	return &listCredentialsResponse, err
}

func (a *gitCredentialsImpl) Update(ctx context.Context, request UpdateCredentialsRequest) error {
	var updateCredentialsResponse UpdateCredentialsResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateCredentialsResponse)
	return err
}

// unexported type that holds implementations of just Repos API methods
type reposImpl struct {
	client *client.DatabricksClient
}

func (a *reposImpl) Create(ctx context.Context, request CreateRepoRequest) (*CreateRepoResponse, error) {
	var createRepoResponse CreateRepoResponse
	path := "/api/2.0/repos"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createRepoResponse)
	return &createRepoResponse, err
}

func (a *reposImpl) Delete(ctx context.Context, request DeleteRepoRequest) error {
	var deleteRepoResponse DeleteRepoResponse
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteRepoResponse)
	return err
}

func (a *reposImpl) Get(ctx context.Context, request GetRepoRequest) (*GetRepoResponse, error) {
	var getRepoResponse GetRepoResponse
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getRepoResponse)
	return &getRepoResponse, err
}

func (a *reposImpl) GetPermissionLevels(ctx context.Context, request GetRepoPermissionLevelsRequest) (*GetRepoPermissionLevelsResponse, error) {
	var getRepoPermissionLevelsResponse GetRepoPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v/permissionLevels", request.RepoId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getRepoPermissionLevelsResponse)
	return &getRepoPermissionLevelsResponse, err
}

func (a *reposImpl) GetPermissions(ctx context.Context, request GetRepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &repoPermissions)
	return &repoPermissions, err
}

func (a *reposImpl) List(ctx context.Context, request ListReposRequest) (*ListReposResponse, error) {
	var listReposResponse ListReposResponse
	path := "/api/2.0/repos"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listReposResponse)
	return &listReposResponse, err
}

func (a *reposImpl) SetPermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &repoPermissions)
	return &repoPermissions, err
}

func (a *reposImpl) Update(ctx context.Context, request UpdateRepoRequest) error {
	var updateRepoResponse UpdateRepoResponse
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateRepoResponse)
	return err
}

func (a *reposImpl) UpdatePermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {
	var repoPermissions RepoPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &repoPermissions)
	return &repoPermissions, err
}

// unexported type that holds implementations of just Secrets API methods
type secretsImpl struct {
	client *client.DatabricksClient
}

func (a *secretsImpl) CreateScope(ctx context.Context, request CreateScope) error {
	var createScopeResponse CreateScopeResponse
	path := "/api/2.0/secrets/scopes/create"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createScopeResponse)
	return err
}

func (a *secretsImpl) DeleteAcl(ctx context.Context, request DeleteAcl) error {
	var deleteAclResponse DeleteAclResponse
	path := "/api/2.0/secrets/acls/delete"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deleteAclResponse)
	return err
}

func (a *secretsImpl) DeleteScope(ctx context.Context, request DeleteScope) error {
	var deleteScopeResponse DeleteScopeResponse
	path := "/api/2.0/secrets/scopes/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deleteScopeResponse)
	return err
}

func (a *secretsImpl) DeleteSecret(ctx context.Context, request DeleteSecret) error {
	var deleteSecretResponse DeleteSecretResponse
	path := "/api/2.0/secrets/delete"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deleteSecretResponse)
	return err
}

func (a *secretsImpl) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {
	var aclItem AclItem
	path := "/api/2.0/secrets/acls/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &aclItem)
	return &aclItem, err
}

func (a *secretsImpl) GetSecret(ctx context.Context, request GetSecretRequest) (*GetSecretResponse, error) {
	var getSecretResponse GetSecretResponse
	path := "/api/2.0/secrets/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getSecretResponse)
	return &getSecretResponse, err
}

func (a *secretsImpl) ListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error) {
	var listAclsResponse ListAclsResponse
	path := "/api/2.0/secrets/acls/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAclsResponse)
	return &listAclsResponse, err
}

func (a *secretsImpl) ListScopes(ctx context.Context) (*ListScopesResponse, error) {
	var listScopesResponse ListScopesResponse
	path := "/api/2.0/secrets/scopes/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listScopesResponse)
	return &listScopesResponse, err
}

func (a *secretsImpl) ListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error) {
	var listSecretsResponse ListSecretsResponse
	path := "/api/2.0/secrets/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listSecretsResponse)
	return &listSecretsResponse, err
}

func (a *secretsImpl) PutAcl(ctx context.Context, request PutAcl) error {
	var putAclResponse PutAclResponse
	path := "/api/2.0/secrets/acls/put"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &putAclResponse)
	return err
}

func (a *secretsImpl) PutSecret(ctx context.Context, request PutSecret) error {
	var putSecretResponse PutSecretResponse
	path := "/api/2.0/secrets/put"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &putSecretResponse)
	return err
}

// unexported type that holds implementations of just Workspace API methods
type workspaceImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceImpl) Delete(ctx context.Context, request Delete) error {
	var deleteResponse DeleteResponse
	path := "/api/2.0/workspace/delete"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deleteResponse)
	return err
}

func (a *workspaceImpl) Export(ctx context.Context, request ExportRequest) (*ExportResponse, error) {
	var exportResponse ExportResponse
	path := "/api/2.0/workspace/export"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &exportResponse)
	return &exportResponse, err
}

func (a *workspaceImpl) GetPermissionLevels(ctx context.Context, request GetWorkspaceObjectPermissionLevelsRequest) (*GetWorkspaceObjectPermissionLevelsResponse, error) {
	var getWorkspaceObjectPermissionLevelsResponse GetWorkspaceObjectPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.WorkspaceObjectType, request.WorkspaceObjectId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getWorkspaceObjectPermissionLevelsResponse)
	return &getWorkspaceObjectPermissionLevelsResponse, err
}

func (a *workspaceImpl) GetPermissions(ctx context.Context, request GetWorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}

func (a *workspaceImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error) {
	var objectInfo ObjectInfo
	path := "/api/2.0/workspace/get-status"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &objectInfo)
	return &objectInfo, err
}

func (a *workspaceImpl) Import(ctx context.Context, request Import) error {
	var importResponse ImportResponse
	path := "/api/2.0/workspace/import"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &importResponse)
	return err
}

func (a *workspaceImpl) List(ctx context.Context, request ListWorkspaceRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/workspace/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listResponse)
	return &listResponse, err
}

func (a *workspaceImpl) Mkdirs(ctx context.Context, request Mkdirs) error {
	var mkdirsResponse MkdirsResponse
	path := "/api/2.0/workspace/mkdirs"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &mkdirsResponse)
	return err
}

func (a *workspaceImpl) SetPermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}

func (a *workspaceImpl) UpdatePermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	var workspaceObjectPermissions WorkspaceObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &workspaceObjectPermissions)
	return &workspaceObjectPermissions, err
}
