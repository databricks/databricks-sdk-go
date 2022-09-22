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

// Create a Git credential entry
//
// Creates a Git credential entry for the user. Only one Git credential per user
// is supported, so any attempts to create credentials if an entry already
// exists will fail. Use the PATCH endpoint to update existing credentials, or
// the DELETE endpoint to delete existing credentials.
func (a *GitCredentialsAPI) Create(ctx context.Context, request CreateCredentials) (*CreateCredentialsResponse, error) {
	var createCredentialsResponse CreateCredentialsResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Post(ctx, path, request, &createCredentialsResponse)
	return &createCredentialsResponse, err
}

// Deletes the credential
//
// Deletes the specified credential
func (a *GitCredentialsAPI) Delete(ctx context.Context, request DeleteRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Deletes the credential
//
// Deletes the specified credential
func (a *GitCredentialsAPI) DeleteByCredentialId(ctx context.Context, credentialId int64) error {
	return a.Delete(ctx, DeleteRequest{
		CredentialId: credentialId,
	})
}

// Get a credential entry
//
// Returns the credential with the given credential ID.
func (a *GitCredentialsAPI) Get(ctx context.Context, request GetRequest) (*GetCredentialResponse, error) {
	var getCredentialResponse GetCredentialResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Get(ctx, path, request, &getCredentialResponse)
	return &getCredentialResponse, err
}

// Get a credential entry
//
// Returns the credential with the given credential ID.
func (a *GitCredentialsAPI) GetByCredentialId(ctx context.Context, credentialId int64) (*GetCredentialResponse, error) {
	return a.Get(ctx, GetRequest{
		CredentialId: credentialId,
	})
}

// Get Git credentials
//
// Returns the calling user's Git credentials. One credential per user is
// supported.
func (a *GitCredentialsAPI) List(ctx context.Context) (*GetCredentialsResponse, error) {
	var getCredentialsResponse GetCredentialsResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Get(ctx, path, nil, &getCredentialsResponse)
	return &getCredentialsResponse, err
}

func (a *GitCredentialsAPI) ListAll(ctx context.Context) ([]GetCredentialResponse, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Credentials, nil
}

// Updates the credential
//
// Updates the credential.
func (a *GitCredentialsAPI) Update(ctx context.Context, request UpdateCredentials) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Patch(ctx, path, request)
	return err
}
