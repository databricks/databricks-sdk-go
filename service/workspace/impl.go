// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just GitCredentials API methods
type gitCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *gitCredentialsImpl) Create(ctx context.Context, request CreateCredentials) (*CreateCredentialsResponse, error) {
	var createCredentialsResponse CreateCredentialsResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createCredentialsResponse)
	return &createCredentialsResponse, err
}

func (a *gitCredentialsImpl) Delete(ctx context.Context, request DeleteGitCredentialRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *gitCredentialsImpl) Get(ctx context.Context, request GetGitCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *gitCredentialsImpl) List(ctx context.Context) (*GetCredentialsResponse, error) {
	var getCredentialsResponse GetCredentialsResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getCredentialsResponse)
	return &getCredentialsResponse, err
}

func (a *gitCredentialsImpl) Update(ctx context.Context, request UpdateCredentials) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Repos API methods
type reposImpl struct {
	client *client.DatabricksClient
}

func (a *reposImpl) Create(ctx context.Context, request CreateRepo) (*RepoInfo, error) {
	var repoInfo RepoInfo
	path := "/api/2.0/repos"
	err := a.client.Do(ctx, http.MethodPost, path, request, &repoInfo)
	return &repoInfo, err
}

func (a *reposImpl) Delete(ctx context.Context, request DeleteRepoRequest) error {
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *reposImpl) Get(ctx context.Context, request GetRepoRequest) (*RepoInfo, error) {
	var repoInfo RepoInfo
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &repoInfo)
	return &repoInfo, err
}

func (a *reposImpl) List(ctx context.Context, request ListReposRequest) (*ListReposResponse, error) {
	var listReposResponse ListReposResponse
	path := "/api/2.0/repos"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listReposResponse)
	return &listReposResponse, err
}

func (a *reposImpl) Update(ctx context.Context, request UpdateRepo) error {
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Secrets API methods
type secretsImpl struct {
	client *client.DatabricksClient
}

func (a *secretsImpl) CreateScope(ctx context.Context, request CreateScope) error {
	path := "/api/2.0/secrets/scopes/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *secretsImpl) DeleteAcl(ctx context.Context, request DeleteAcl) error {
	path := "/api/2.0/secrets/acls/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *secretsImpl) DeleteScope(ctx context.Context, request DeleteScope) error {
	path := "/api/2.0/secrets/scopes/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *secretsImpl) DeleteSecret(ctx context.Context, request DeleteSecret) error {
	path := "/api/2.0/secrets/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *secretsImpl) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {
	var aclItem AclItem
	path := "/api/2.0/secrets/acls/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &aclItem)
	return &aclItem, err
}

func (a *secretsImpl) ListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error) {
	var listAclsResponse ListAclsResponse
	path := "/api/2.0/secrets/acls/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listAclsResponse)
	return &listAclsResponse, err
}

func (a *secretsImpl) ListScopes(ctx context.Context) (*ListScopesResponse, error) {
	var listScopesResponse ListScopesResponse
	path := "/api/2.0/secrets/scopes/list"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listScopesResponse)
	return &listScopesResponse, err
}

func (a *secretsImpl) ListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error) {
	var listSecretsResponse ListSecretsResponse
	path := "/api/2.0/secrets/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listSecretsResponse)
	return &listSecretsResponse, err
}

func (a *secretsImpl) PutAcl(ctx context.Context, request PutAcl) error {
	path := "/api/2.0/secrets/acls/put"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *secretsImpl) PutSecret(ctx context.Context, request PutSecret) error {
	path := "/api/2.0/secrets/put"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

// unexported type that holds implementations of just Workspace API methods
type workspaceImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceImpl) Delete(ctx context.Context, request Delete) error {
	path := "/api/2.0/workspace/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *workspaceImpl) Export(ctx context.Context, request ExportRequest) (*ExportResponse, error) {
	var exportResponse ExportResponse
	path := "/api/2.0/workspace/export"
	err := a.client.Do(ctx, http.MethodGet, path, request, &exportResponse)
	return &exportResponse, err
}

func (a *workspaceImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error) {
	var objectInfo ObjectInfo
	path := "/api/2.0/workspace/get-status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &objectInfo)
	return &objectInfo, err
}

func (a *workspaceImpl) Import(ctx context.Context, request Import) error {
	path := "/api/2.0/workspace/import"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *workspaceImpl) List(ctx context.Context, request ListWorkspaceRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/workspace/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listResponse)
	return &listResponse, err
}

func (a *workspaceImpl) Mkdirs(ctx context.Context, request Mkdirs) error {
	path := "/api/2.0/workspace/mkdirs"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}
