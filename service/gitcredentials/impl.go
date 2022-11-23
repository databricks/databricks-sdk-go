// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package gitcredentials

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just GitCredentials API methods
type gitCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *gitCredentialsImpl) Create(ctx context.Context, request CreateCredentials) (*CreateCredentialsResponse, error) {
	var createCredentialsResponse CreateCredentialsResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Post(ctx, path, request, &createCredentialsResponse)
	return &createCredentialsResponse, err
}

func (a *gitCredentialsImpl) Delete(ctx context.Context, request DeleteRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *gitCredentialsImpl) Get(ctx context.Context, request GetRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Get(ctx, path, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *gitCredentialsImpl) List(ctx context.Context) (*GetCredentialsResponse, error) {
	var getCredentialsResponse GetCredentialsResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Get(ctx, path, nil, &getCredentialsResponse)
	return &getCredentialsResponse, err
}

func (a *gitCredentialsImpl) Update(ctx context.Context, request UpdateCredentials) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Patch(ctx, path, request)
	return err
}
