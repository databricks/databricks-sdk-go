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

	requestPb, pbErr := createCredentialsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createCredentialsResponsePb createCredentialsResponsePb
	path := "/api/2.0/git-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createCredentialsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := createCredentialsResponseFromPb(&createCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *gitCredentialsImpl) Delete(ctx context.Context, request DeleteCredentialsRequest) error {

	requestPb, pbErr := deleteCredentialsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteCredentialsResponsePb deleteCredentialsResponsePb
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", requestPb.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, (*requestPb), &deleteCredentialsResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *gitCredentialsImpl) Get(ctx context.Context, request GetCredentialsRequest) (*GetCredentialsResponse, error) {

	requestPb, pbErr := getCredentialsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getCredentialsResponsePb getCredentialsResponsePb
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", requestPb.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getCredentialsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getCredentialsResponseFromPb(&getCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	var listCredentialsResponsePb listCredentialsResponsePb
	path := "/api/2.0/git-credentials"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listCredentialsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listCredentialsResponseFromPb(&listCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *gitCredentialsImpl) Update(ctx context.Context, request UpdateCredentialsRequest) error {

	requestPb, pbErr := updateCredentialsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateCredentialsResponsePb updateCredentialsResponsePb
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", requestPb.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &updateCredentialsResponsePb)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just Repos API methods
type reposImpl struct {
	client *client.DatabricksClient
}

func (a *reposImpl) Create(ctx context.Context, request CreateRepoRequest) (*CreateRepoResponse, error) {

	requestPb, pbErr := createRepoRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createRepoResponsePb createRepoResponsePb
	path := "/api/2.0/repos"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createRepoResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := createRepoResponseFromPb(&createRepoResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) Delete(ctx context.Context, request DeleteRepoRequest) error {

	requestPb, pbErr := deleteRepoRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteRepoResponsePb deleteRepoResponsePb
	path := fmt.Sprintf("/api/2.0/repos/%v", requestPb.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, (*requestPb), &deleteRepoResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *reposImpl) Get(ctx context.Context, request GetRepoRequest) (*GetRepoResponse, error) {

	requestPb, pbErr := getRepoRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRepoResponsePb getRepoResponsePb
	path := fmt.Sprintf("/api/2.0/repos/%v", requestPb.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getRepoResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getRepoResponseFromPb(&getRepoResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) GetPermissionLevels(ctx context.Context, request GetRepoPermissionLevelsRequest) (*GetRepoPermissionLevelsResponse, error) {

	requestPb, pbErr := getRepoPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRepoPermissionLevelsResponsePb getRepoPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v/permissionLevels", requestPb.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getRepoPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getRepoPermissionLevelsResponseFromPb(&getRepoPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) GetPermissions(ctx context.Context, request GetRepoPermissionsRequest) (*RepoPermissions, error) {

	requestPb, pbErr := getRepoPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var repoPermissionsPb repoPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", requestPb.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &repoPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := repoPermissionsFromPb(&repoPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	requestPb, pbErr := listReposRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listReposResponsePb listReposResponsePb
	path := "/api/2.0/repos"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listReposResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listReposResponseFromPb(&listReposResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) SetPermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {

	requestPb, pbErr := repoPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var repoPermissionsPb repoPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", requestPb.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, (*requestPb), &repoPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := repoPermissionsFromPb(&repoPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) Update(ctx context.Context, request UpdateRepoRequest) error {

	requestPb, pbErr := updateRepoRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateRepoResponsePb updateRepoResponsePb
	path := fmt.Sprintf("/api/2.0/repos/%v", requestPb.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &updateRepoResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *reposImpl) UpdatePermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {

	requestPb, pbErr := repoPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var repoPermissionsPb repoPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", requestPb.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &repoPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := repoPermissionsFromPb(&repoPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Secrets API methods
type secretsImpl struct {
	client *client.DatabricksClient
}

func (a *secretsImpl) CreateScope(ctx context.Context, request CreateScope) error {

	requestPb, pbErr := createScopeToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var createScopeResponsePb createScopeResponsePb
	path := "/api/2.0/secrets/scopes/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createScopeResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *secretsImpl) DeleteAcl(ctx context.Context, request DeleteAcl) error {

	requestPb, pbErr := deleteAclToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteAclResponsePb deleteAclResponsePb
	path := "/api/2.0/secrets/acls/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &deleteAclResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *secretsImpl) DeleteScope(ctx context.Context, request DeleteScope) error {

	requestPb, pbErr := deleteScopeToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteScopeResponsePb deleteScopeResponsePb
	path := "/api/2.0/secrets/scopes/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &deleteScopeResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *secretsImpl) DeleteSecret(ctx context.Context, request DeleteSecret) error {

	requestPb, pbErr := deleteSecretToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteSecretResponsePb deleteSecretResponsePb
	path := "/api/2.0/secrets/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &deleteSecretResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *secretsImpl) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {

	requestPb, pbErr := getAclRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aclItemPb aclItemPb
	path := "/api/2.0/secrets/acls/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &aclItemPb)
	if err != nil {
		return nil, err
	}
	resp, err := aclItemFromPb(&aclItemPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *secretsImpl) GetSecret(ctx context.Context, request GetSecretRequest) (*GetSecretResponse, error) {

	requestPb, pbErr := getSecretRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getSecretResponsePb getSecretResponsePb
	path := "/api/2.0/secrets/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getSecretResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getSecretResponseFromPb(&getSecretResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	requestPb, pbErr := listAclsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAclsResponsePb listAclsResponsePb
	path := "/api/2.0/secrets/acls/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listAclsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listAclsResponseFromPb(&listAclsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	var listScopesResponsePb listScopesResponsePb
	path := "/api/2.0/secrets/scopes/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listScopesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listScopesResponseFromPb(&listScopesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	requestPb, pbErr := listSecretsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSecretsResponsePb listSecretsResponsePb
	path := "/api/2.0/secrets/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listSecretsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listSecretsResponseFromPb(&listSecretsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *secretsImpl) PutAcl(ctx context.Context, request PutAcl) error {

	requestPb, pbErr := putAclToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var putAclResponsePb putAclResponsePb
	path := "/api/2.0/secrets/acls/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &putAclResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *secretsImpl) PutSecret(ctx context.Context, request PutSecret) error {

	requestPb, pbErr := putSecretToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var putSecretResponsePb putSecretResponsePb
	path := "/api/2.0/secrets/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &putSecretResponsePb)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just Workspace API methods
type workspaceImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceImpl) Delete(ctx context.Context, request Delete) error {

	requestPb, pbErr := deleteToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := "/api/2.0/workspace/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &deleteResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *workspaceImpl) Export(ctx context.Context, request ExportRequest) (*ExportResponse, error) {

	requestPb, pbErr := exportRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var exportResponsePb exportResponsePb
	path := "/api/2.0/workspace/export"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &exportResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := exportResponseFromPb(&exportResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) GetPermissionLevels(ctx context.Context, request GetWorkspaceObjectPermissionLevelsRequest) (*GetWorkspaceObjectPermissionLevelsResponse, error) {

	requestPb, pbErr := getWorkspaceObjectPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getWorkspaceObjectPermissionLevelsResponsePb getWorkspaceObjectPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", requestPb.WorkspaceObjectType, requestPb.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getWorkspaceObjectPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getWorkspaceObjectPermissionLevelsResponseFromPb(&getWorkspaceObjectPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) GetPermissions(ctx context.Context, request GetWorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {

	requestPb, pbErr := getWorkspaceObjectPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceObjectPermissionsPb workspaceObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", requestPb.WorkspaceObjectType, requestPb.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := workspaceObjectPermissionsFromPb(&workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error) {

	requestPb, pbErr := getStatusRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var objectInfoPb objectInfoPb
	path := "/api/2.0/workspace/get-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &objectInfoPb)
	if err != nil {
		return nil, err
	}
	resp, err := objectInfoFromPb(&objectInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) Import(ctx context.Context, request Import) error {

	requestPb, pbErr := importToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var importResponsePb importResponsePb
	path := "/api/2.0/workspace/import"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &importResponsePb)
	if err != nil {
		return err
	}

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

	requestPb, pbErr := listWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listResponsePb listResponsePb
	path := "/api/2.0/workspace/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listResponseFromPb(&listResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) Mkdirs(ctx context.Context, request Mkdirs) error {

	requestPb, pbErr := mkdirsToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var mkdirsResponsePb mkdirsResponsePb
	path := "/api/2.0/workspace/mkdirs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &mkdirsResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *workspaceImpl) SetPermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {

	requestPb, pbErr := workspaceObjectPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceObjectPermissionsPb workspaceObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", requestPb.WorkspaceObjectType, requestPb.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, (*requestPb), &workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := workspaceObjectPermissionsFromPb(&workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) UpdatePermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {

	requestPb, pbErr := workspaceObjectPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceObjectPermissionsPb workspaceObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", requestPb.WorkspaceObjectType, requestPb.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := workspaceObjectPermissionsFromPb(&workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
