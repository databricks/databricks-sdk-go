// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

import (
	"context"
)

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
type TokensService interface {

	// Create a user token
	//
	// Creates and returns a token for a user.
	//
	// If this call is made through token authentication, it creates a token
	// with the same client ID as the authenticated token. If the user's token
	// quota is exceeded, this call returns an error **QUOTA_EXCEEDED**.
	Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error)

	// Revoke token
	//
	// Revokes an access token.
	//
	// If a token with the specified ID is not valid, this call returns an error
	// **RESOURCE_DOES_NOT_EXIST**.
	Delete(ctx context.Context, request RevokeTokenRequest) error

	// List tokens
	//
	// Lists all the valid tokens for a user-workspace pair.
	List(ctx context.Context) (*ListTokensResponse, error)
}
