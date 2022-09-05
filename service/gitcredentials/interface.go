// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package gitcredentials

import (
	"context"
)

type GitCredentialsService interface {
	// Creates a Git credential entry for the user. Only one Git credential per
	// user is supported, so any attempts to create credentials if an entry
	// already exists will fail. Use the PATCH endpoint to update existing
	// credentials, or the DELETE endpoint to delete existing credentials.
	CreateCredential(ctx context.Context, createCredentialRequest CreateCredentialRequest) (*GetCredentialResponse, error)

	// Deletes the specified credential
	DeleteCredential(ctx context.Context, deleteCredentialRequest DeleteCredentialRequest) error

	DeleteCredentialByCredentialId(ctx context.Context, credentialId string) error
	// Returns the credential with the given credential ID.
	GetCredential(ctx context.Context, getCredentialRequest GetCredentialRequest) (*GetCredentialResponse, error)

	GetCredentialByCredentialId(ctx context.Context, credentialId string) (*GetCredentialResponse, error)
	// Returns the calling user&#39;s Git credentials. One credential per user is
	// supported.
	GetCredentials(ctx context.Context) (*GetCredentialsResponse, error)

	// Updates the credential.
	UpdateCredential(ctx context.Context, updateCredentialRequest UpdateCredentialRequest) error
}
