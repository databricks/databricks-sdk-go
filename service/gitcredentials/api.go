// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package gitcredentials

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type GitcredentialsService interface {
    // Creates a Git credential entry for the user. Only one Git credential per
    // user is supported, so any attempts to create credentials if an entry
    // already exists will fail. Use the PATCH endpoint to update existing
    // credentials, or the DELETE endpoint to delete existing credentials.
    CreateCredential(ctx context.Context, createCredentialRequest CreateCredentialRequest) (*GetCredentialResponse, error)
    // Deletes the specified credential
    DeleteCredential(ctx context.Context, deleteCredentialRequest DeleteCredentialRequest) error
    // Returns the credential with the given credential ID.
    GetCredential(ctx context.Context, getCredentialRequest GetCredentialRequest) (*GetCredentialResponse, error)
    // Returns the calling user&#39;s Git credentials. One credential per user is
    // supported.
    GetCredentials(ctx context.Context) (*GetCredentialsResponse, error)
    // Updates the credential.
    UpdateCredential(ctx context.Context, updateCredentialRequest UpdateCredentialRequest) error
	GetCredentialByCredentialId(ctx context.Context, credentialId string) (*GetCredentialResponse, error)
	DeleteCredentialByCredentialId(ctx context.Context, credentialId string) error
}

func New(client *client.DatabricksClient) GitcredentialsService {
	return &GitcredentialsAPI{
		client: client,
	}
}

type GitcredentialsAPI struct {
	client *client.DatabricksClient
}

// Creates a Git credential entry for the user. Only one Git credential per user
// is supported, so any attempts to create credentials if an entry already
// exists will fail. Use the PATCH endpoint to update existing credentials, or
// the DELETE endpoint to delete existing credentials.
func (a *GitcredentialsAPI) CreateCredential(ctx context.Context, request CreateCredentialRequest) (*GetCredentialResponse, error) {
	var getCredentialResponse GetCredentialResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Post(ctx, path, request, &getCredentialResponse)
	return &getCredentialResponse, err
}

// Deletes the specified credential
func (a *GitcredentialsAPI) DeleteCredential(ctx context.Context, request DeleteCredentialRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Returns the credential with the given credential ID.
func (a *GitcredentialsAPI) GetCredential(ctx context.Context, request GetCredentialRequest) (*GetCredentialResponse, error) {
	var getCredentialResponse GetCredentialResponse
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Get(ctx, path, request, &getCredentialResponse)
	return &getCredentialResponse, err
}

// Returns the calling user&#39;s Git credentials. One credential per user is
// supported.
func (a *GitcredentialsAPI) GetCredentials(ctx context.Context) (*GetCredentialsResponse, error) {
	var getCredentialsResponse GetCredentialsResponse
	path := "/api/2.0/git-credentials"
	err := a.client.Get(ctx, path, nil, &getCredentialsResponse)
	return &getCredentialsResponse, err
}

// Updates the credential.
func (a *GitcredentialsAPI) UpdateCredential(ctx context.Context, request UpdateCredentialRequest) error {
	path := fmt.Sprintf("/api/2.0/git-credentials/%v", request.CredentialId)
	err := a.client.Patch(ctx, path, request)
	return err
}


func (a *GitcredentialsAPI) GetCredentialByCredentialId(ctx context.Context, credentialId string) (*GetCredentialResponse, error) {
	return a.GetCredential(ctx, GetCredentialRequest{
		CredentialId: credentialId,
	})
}

func (a *GitcredentialsAPI) DeleteCredentialByCredentialId(ctx context.Context, credentialId string) error {
	return a.DeleteCredential(ctx, DeleteCredentialRequest{
		CredentialId: credentialId,
	})
}
