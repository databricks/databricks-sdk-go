// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

import (
	"context"
)

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
type TokenManagementService interface {

	// Create on-behalf token.
	//
	// Creates a token on behalf of a service principal.
	CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error)

	// Delete a token.
	//
	// Deletes a token, specified by its ID.
	Delete(ctx context.Context, request Delete) error

	// Get token info.
	//
	// Gets information about a token, specified by its ID.
	Get(ctx context.Context, request Get) (*TokenInfo, error)

	// List all tokens.
	//
	// Lists all tokens associated with the specified workspace or user.
	//
	// Use ListAll() to get all TokenInfo instances
	List(ctx context.Context, request List) (*ListTokensResponse, error)
}
