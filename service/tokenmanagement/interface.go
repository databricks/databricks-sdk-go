// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

import (
	"context"
)

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
type TokenManagementService interface {
	// Delete a token, specified by its ID.
	DeleteToken(ctx context.Context, deleteTokenRequest DeleteTokenRequest) error

	DeleteTokenByTokenId(ctx context.Context, tokenId string) error
	// Get a token, specified by its ID.
	GetTokenInfo(ctx context.Context, getTokenInfoRequest GetTokenInfoRequest) (*TokenInfo, error)

	GetTokenInfoByTokenId(ctx context.Context, tokenId string) (*TokenInfo, error)
	// List all tokens belonging to a workspace or a user.
	ListAllTokens(ctx context.Context, listAllTokensRequest ListAllTokensRequest) (*ListTokensResponse, error)
}
