// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

import (
	"context"
)



type TokensService interface {
    // Creates and returns a token for a user. If this call is made through
    // token authentication, it will create the token that has the same client
    // id with the authenticated token. This call returns an error
    // ``QUOTA_EXCEEDED`` if over the token quota for the user.
    CreateToken(ctx context.Context, createTokenRequest CreateTokenRequest) (*CreateTokenResponse, error)
    // Lists all the valid tokens for a user-workspace pair.
    ListTokens(ctx context.Context) (*ListTokensResponse, error)
    // Revokes an access token. This call returns an error
    // ``RESOURCE_DOES_NOT_EXIST`` if a token with the given ID is not valid.
    RevokeToken(ctx context.Context, revokeTokenRequest RevokeTokenRequest) error
}
