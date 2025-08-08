// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/workspace/workspacepb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just GitCredentials API methods
type gitCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *gitCredentialsImpl) Create(ctx context.Context, request CreateCredentialsRequest) (*CreateCredentialsResponse, error) {
	requestPb, pbErr := CreateCredentialsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createCredentialsResponsePb workspacepb.CreateCredentialsResponsePb
	path := "/api/2.0/git-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createCredentialsResponsePb,
	)
	resp, err := CreateCredentialsResponseFromPb(&createCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *gitCredentialsImpl) Delete(ctx context.Context, request DeleteCredentialsRequest) error {
	requestPb, pbErr := DeleteCredentialsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *gitCredentialsImpl) Get(ctx context.Context, request GetCredentialsRequest) (*GetCredentialsResponse, error) {
	requestPb, pbErr := GetCredentialsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getCredentialsResponsePb workspacepb.GetCredentialsResponsePb
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getCredentialsResponsePb,
	)
	resp, err := GetCredentialsResponseFromPb(&getCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Lists the calling user's Git credentials. One credential per user is
// supported.
func (a *gitCredentialsImpl) ListAll(ctx context.Context) ([]CredentialInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[CredentialInfo](ctx, iterator)
}

func (a *gitCredentialsImpl) internalList(ctx context.Context) (*ListCredentialsResponse, error) {

	var listCredentialsResponsePb workspacepb.ListCredentialsResponsePb
	path := "/api/2.0/git-credentials"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&listCredentialsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListCredentialsResponseFromPb(&listCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *gitCredentialsImpl) Update(ctx context.Context, request UpdateCredentialsRequest) error {
	requestPb, pbErr := UpdateCredentialsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just Repos API methods
type reposImpl struct {
	client *client.DatabricksClient
}

func (a *reposImpl) Create(ctx context.Context, request CreateRepoRequest) (*CreateRepoResponse, error) {
	requestPb, pbErr := CreateRepoRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createRepoResponsePb workspacepb.CreateRepoResponsePb
	path := "/api/2.0/repos"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createRepoResponsePb,
	)
	resp, err := CreateRepoResponseFromPb(&createRepoResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) Delete(ctx context.Context, request DeleteRepoRequest) error {
	requestPb, pbErr := DeleteRepoRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *reposImpl) Get(ctx context.Context, request GetRepoRequest) (*GetRepoResponse, error) {
	requestPb, pbErr := GetRepoRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRepoResponsePb workspacepb.GetRepoResponsePb
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getRepoResponsePb,
	)
	resp, err := GetRepoResponseFromPb(&getRepoResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) GetPermissionLevels(ctx context.Context, request GetRepoPermissionLevelsRequest) (*GetRepoPermissionLevelsResponse, error) {
	requestPb, pbErr := GetRepoPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getRepoPermissionLevelsResponsePb workspacepb.GetRepoPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v/permissionLevels", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getRepoPermissionLevelsResponsePb,
	)
	resp, err := GetRepoPermissionLevelsResponseFromPb(&getRepoPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) GetPermissions(ctx context.Context, request GetRepoPermissionsRequest) (*RepoPermissions, error) {
	requestPb, pbErr := GetRepoPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var repoPermissionsPb workspacepb.RepoPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&repoPermissionsPb,
	)
	resp, err := RepoPermissionsFromPb(&repoPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Returns repos that the calling user has Manage permissions on. Use
// `next_page_token` to iterate through additional pages.
func (a *reposImpl) ListAll(ctx context.Context, request ListReposRequest) ([]RepoInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[RepoInfo](ctx, iterator)
}

func (a *reposImpl) internalList(ctx context.Context, request ListReposRequest) (*ListReposResponse, error) {
	requestPb, pbErr := ListReposRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listReposResponsePb workspacepb.ListReposResponsePb
	path := "/api/2.0/repos"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listReposResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListReposResponseFromPb(&listReposResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) SetPermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {
	requestPb, pbErr := RepoPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var repoPermissionsPb workspacepb.RepoPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&repoPermissionsPb,
	)
	resp, err := RepoPermissionsFromPb(&repoPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *reposImpl) Update(ctx context.Context, request UpdateRepoRequest) error {
	requestPb, pbErr := UpdateRepoRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *reposImpl) UpdatePermissions(ctx context.Context, request RepoPermissionsRequest) (*RepoPermissions, error) {
	requestPb, pbErr := RepoPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var repoPermissionsPb workspacepb.RepoPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&repoPermissionsPb,
	)
	resp, err := RepoPermissionsFromPb(&repoPermissionsPb)
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
	requestPb, pbErr := CreateScopeToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/secrets/scopes/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *secretsImpl) DeleteAcl(ctx context.Context, request DeleteAcl) error {
	requestPb, pbErr := DeleteAclToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/secrets/acls/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *secretsImpl) DeleteScope(ctx context.Context, request DeleteScope) error {
	requestPb, pbErr := DeleteScopeToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/secrets/scopes/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *secretsImpl) DeleteSecret(ctx context.Context, request DeleteSecret) error {
	requestPb, pbErr := DeleteSecretToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/secrets/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *secretsImpl) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {
	requestPb, pbErr := GetAclRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var aclItemPb workspacepb.AclItemPb
	path := "/api/2.0/secrets/acls/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&aclItemPb,
	)
	resp, err := AclItemFromPb(&aclItemPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *secretsImpl) GetSecret(ctx context.Context, request GetSecretRequest) (*GetSecretResponse, error) {
	requestPb, pbErr := GetSecretRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getSecretResponsePb workspacepb.GetSecretResponsePb
	path := "/api/2.0/secrets/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getSecretResponsePb,
	)
	resp, err := GetSecretResponseFromPb(&getSecretResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists the ACLs set on the given scope.
//
// Users must have the “MANAGE“ permission to invoke this API.
//
// Example response:
//
// .. code::
//
// { "acls": [{ "principal": "admins", "permission": "MANAGE" },{ "principal":
// "data-scientists", "permission": "READ" }] }
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
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

// Lists the ACLs set on the given scope.
//
// Users must have the “MANAGE“ permission to invoke this API.
//
// Example response:
//
// .. code::
//
// { "acls": [{ "principal": "admins", "permission": "MANAGE" },{ "principal":
// "data-scientists", "permission": "READ" }] }
//
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListAclsAll(ctx context.Context, request ListAclsRequest) ([]AclItem, error) {
	iterator := a.ListAcls(ctx, request)
	return listing.ToSlice[AclItem](ctx, iterator)
}

func (a *secretsImpl) internalListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error) {
	requestPb, pbErr := ListAclsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAclsResponsePb workspacepb.ListAclsResponsePb
	path := "/api/2.0/secrets/acls/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listAclsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListAclsResponseFromPb(&listAclsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists all secret scopes available in the workspace.
//
// Example response:
//
// .. code::
//
// { "scopes": [{ "name": "my-databricks-scope", "backend_type": "DATABRICKS"
// },{ "name": "mount-points", "backend_type": "DATABRICKS" }] }
//
// Throws “PERMISSION_DENIED“ if the user does not have permission to make
// this API call.
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

// Lists all secret scopes available in the workspace.
//
// Example response:
//
// .. code::
//
// { "scopes": [{ "name": "my-databricks-scope", "backend_type": "DATABRICKS"
// },{ "name": "mount-points", "backend_type": "DATABRICKS" }] }
//
// Throws “PERMISSION_DENIED“ if the user does not have permission to make
// this API call.
func (a *secretsImpl) ListScopesAll(ctx context.Context) ([]SecretScope, error) {
	iterator := a.ListScopes(ctx)
	return listing.ToSlice[SecretScope](ctx, iterator)
}

func (a *secretsImpl) internalListScopes(ctx context.Context) (*ListScopesResponse, error) {

	var listScopesResponsePb workspacepb.ListScopesResponsePb
	path := "/api/2.0/secrets/scopes/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&listScopesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListScopesResponseFromPb(&listScopesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
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
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
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

// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
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
// Throws “RESOURCE_DOES_NOT_EXIST“ if no such secret scope exists. Throws
// “PERMISSION_DENIED“ if the user does not have permission to make this API
// call.
func (a *secretsImpl) ListSecretsAll(ctx context.Context, request ListSecretsRequest) ([]SecretMetadata, error) {
	iterator := a.ListSecrets(ctx, request)
	return listing.ToSlice[SecretMetadata](ctx, iterator)
}

func (a *secretsImpl) internalListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error) {
	requestPb, pbErr := ListSecretsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSecretsResponsePb workspacepb.ListSecretsResponsePb
	path := "/api/2.0/secrets/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listSecretsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListSecretsResponseFromPb(&listSecretsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *secretsImpl) PutAcl(ctx context.Context, request PutAcl) error {
	requestPb, pbErr := PutAclToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/secrets/acls/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *secretsImpl) PutSecret(ctx context.Context, request PutSecret) error {
	requestPb, pbErr := PutSecretToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/secrets/put"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just Workspace API methods
type workspaceImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceImpl) Delete(ctx context.Context, request Delete) error {
	requestPb, pbErr := DeleteToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/workspace/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *workspaceImpl) Export(ctx context.Context, request ExportRequest) (*ExportResponse, error) {
	requestPb, pbErr := ExportRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var exportResponsePb workspacepb.ExportResponsePb
	path := "/api/2.0/workspace/export"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&exportResponsePb,
	)
	resp, err := ExportResponseFromPb(&exportResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) GetPermissionLevels(ctx context.Context, request GetWorkspaceObjectPermissionLevelsRequest) (*GetWorkspaceObjectPermissionLevelsResponse, error) {
	requestPb, pbErr := GetWorkspaceObjectPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getWorkspaceObjectPermissionLevelsResponsePb workspacepb.GetWorkspaceObjectPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getWorkspaceObjectPermissionLevelsResponsePb,
	)
	resp, err := GetWorkspaceObjectPermissionLevelsResponseFromPb(&getWorkspaceObjectPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) GetPermissions(ctx context.Context, request GetWorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	requestPb, pbErr := GetWorkspaceObjectPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceObjectPermissionsPb workspacepb.WorkspaceObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&workspaceObjectPermissionsPb,
	)
	resp, err := WorkspaceObjectPermissionsFromPb(&workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error) {
	requestPb, pbErr := GetStatusRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var objectInfoPb workspacepb.ObjectInfoPb
	path := "/api/2.0/workspace/get-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&objectInfoPb,
	)
	resp, err := ObjectInfoFromPb(&objectInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) Import(ctx context.Context, request Import) error {
	requestPb, pbErr := ImportToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/workspace/import"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

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

// Lists the contents of a directory, or the object if it is not a directory. If
// the input path does not exist, this call returns an error
// `RESOURCE_DOES_NOT_EXIST`.
func (a *workspaceImpl) ListAll(ctx context.Context, request ListWorkspaceRequest) ([]ObjectInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ObjectInfo](ctx, iterator)
}

func (a *workspaceImpl) internalList(ctx context.Context, request ListWorkspaceRequest) (*ListResponse, error) {
	requestPb, pbErr := ListWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listResponsePb workspacepb.ListResponsePb
	path := "/api/2.0/workspace/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListResponseFromPb(&listResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) Mkdirs(ctx context.Context, request Mkdirs) error {
	requestPb, pbErr := MkdirsToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/workspace/mkdirs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *workspaceImpl) SetPermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	requestPb, pbErr := WorkspaceObjectPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceObjectPermissionsPb workspacepb.WorkspaceObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&workspaceObjectPermissionsPb,
	)
	resp, err := WorkspaceObjectPermissionsFromPb(&workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceImpl) UpdatePermissions(ctx context.Context, request WorkspaceObjectPermissionsRequest) (*WorkspaceObjectPermissions, error) {
	requestPb, pbErr := WorkspaceObjectPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspaceObjectPermissionsPb workspacepb.WorkspaceObjectPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.WorkspaceObjectType, request.WorkspaceObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&workspaceObjectPermissionsPb,
	)
	resp, err := WorkspaceObjectPermissionsFromPb(&workspaceObjectPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
