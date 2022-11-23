// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package gitcredentials

import (
	"context"
)

// Registers personal access token for Databricks to do operations on behalf of
// the user.
//
// See [more
// info](https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html).
type GitCredentialsService interface {

	// Create a credential entry
	//
	// Creates a Git credential entry for the user. Only one Git credential per
	// user is supported, so any attempts to create credentials if an entry
	// already exists will fail. Use the PATCH endpoint to update existing
	// credentials, or the DELETE endpoint to delete existing credentials.
	Create(ctx context.Context, request CreateCredentials) (*CreateCredentialsResponse, error)

	// Delete a credential
	//
	// Deletes the specified Git credential.
	Delete(ctx context.Context, request DeleteRequest) error

	// Get a credential entry
	//
	// Gets the Git credential with the specified credential ID.
	Get(ctx context.Context, request GetRequest) (*CredentialInfo, error)

	// Get Git credentials
	//
	// Lists the calling user's Git credentials. One credential per user is
	// supported.
	//
	// Use ListAll() to get all CredentialInfo instances
	List(ctx context.Context) (*GetCredentialsResponse, error)

	// Update a credential
	//
	// Updates the specified Git credential.
	Update(ctx context.Context, request UpdateCredentials) error
}
