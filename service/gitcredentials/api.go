// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package gitcredentials

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewGitCredentials(client *client.DatabricksClient) GitCredentialsService {
	return &GitCredentialsAPI{
		client: client,
	}
}

type GitCredentialsAPI struct {
	client *client.DatabricksClient
}

// Creates a Git credential entry for the user. Only one Git credential per user
// is supported, so any attempts to create credentials if an entry already
// exists will fail. Use the PATCH endpoint to update existing credentials, or
// the DELETE endpoint to delete existing credentials.
func (a *GitCredentialsAPI) CreateCredential(ctx context.Context, request CreateCredentialRequest) (*GetCredentialResponse, error) {
	var getCredentialResponse GetCredentialResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Post(ctx, path, request, &getCredentialResponse)
	return &getCredentialResponse, err
}

// Deletes the specified credential
func (a *GitCredentialsAPI) DeleteCredential(ctx context.Context, request DeleteCredentialRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Deletes the specified credential
func (a *GitCredentialsAPI) DeleteCredentialByCredentialId(ctx context.Context, credentialId string) error {
	return a.DeleteCredential(ctx, DeleteCredentialRequest{
		CredentialId: credentialId,
	})
}

// Returns the credential with the given credential ID.
func (a *GitCredentialsAPI) GetCredential(ctx context.Context, request GetCredentialRequest) (*GetCredentialResponse, error) {
	var getCredentialResponse GetCredentialResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Get(ctx, path, request, &getCredentialResponse)
	return &getCredentialResponse, err
}

// Returns the credential with the given credential ID.
func (a *GitCredentialsAPI) GetCredentialByCredentialId(ctx context.Context, credentialId string) (*GetCredentialResponse, error) {
	return a.GetCredential(ctx, GetCredentialRequest{
		CredentialId: credentialId,
	})
}

// Returns the calling user&#39;s Git credentials. One credential per user is
// supported.
func (a *GitCredentialsAPI) GetCredentials(ctx context.Context) (*GetCredentialsResponse, error) {
	var getCredentialsResponse GetCredentialsResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Get(ctx, path, nil, &getCredentialsResponse)
	return &getCredentialsResponse, err
}

// Updates the credential.
func (a *GitCredentialsAPI) UpdateCredential(ctx context.Context, request UpdateCredentialRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Patch(ctx, path, request)
	return err
}
