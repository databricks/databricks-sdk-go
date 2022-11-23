// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package secrets

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just Secrets API methods
type secretsImpl struct {
	client *client.DatabricksClient
}

func (a *secretsImpl) CreateScope(ctx context.Context, request CreateScope) error {
	path := "/api/2.0/secrets/scopes/create"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *secretsImpl) DeleteAcl(ctx context.Context, request DeleteAcl) error {
	path := "/api/2.0/secrets/acls/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *secretsImpl) DeleteScope(ctx context.Context, request DeleteScope) error {
	path := "/api/2.0/secrets/scopes/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *secretsImpl) DeleteSecret(ctx context.Context, request DeleteSecret) error {
	path := "/api/2.0/secrets/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *secretsImpl) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {
	var aclItem AclItem
	path := "/api/2.0/secrets/acls/get"
	err := a.client.Get(ctx, path, request, &aclItem)
	return &aclItem, err
}

func (a *secretsImpl) ListAcls(ctx context.Context, request ListAclsRequest) (*ListAclsResponse, error) {
	var listAclsResponse ListAclsResponse
	path := "/api/2.0/secrets/acls/list"
	err := a.client.Get(ctx, path, request, &listAclsResponse)
	return &listAclsResponse, err
}

func (a *secretsImpl) ListScopes(ctx context.Context) (*ListScopesResponse, error) {
	var listScopesResponse ListScopesResponse
	path := "/api/2.0/secrets/scopes/list"
	err := a.client.Get(ctx, path, nil, &listScopesResponse)
	return &listScopesResponse, err
}

func (a *secretsImpl) ListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error) {
	var listSecretsResponse ListSecretsResponse
	path := "/api/2.0/secrets/list"
	err := a.client.Get(ctx, path, request, &listSecretsResponse)
	return &listSecretsResponse, err
}

func (a *secretsImpl) PutAcl(ctx context.Context, request PutAcl) error {
	path := "/api/2.0/secrets/acls/put"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *secretsImpl) PutSecret(ctx context.Context, request PutSecret) error {
	path := "/api/2.0/secrets/put"
	err := a.client.Post(ctx, path, request, nil)
	return err
}
