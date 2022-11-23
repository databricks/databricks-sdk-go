// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewTokens(client *client.DatabricksClient) *TokensAPI {
	return &TokensAPI{
		impl: &tokensImpl{
			client: client,
		},
	}
}

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
type TokensAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(TokensService)
	impl TokensService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *TokensAPI) WithImpl(impl TokensService) *TokensAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Tokens API implementation
func (a *TokensAPI) Impl() TokensService {
	return a.impl
}

// Create a user token
//
// Creates and returns a token for a user.
//
// If this call is made through token authentication, it creates a token with
// the same client ID as the authenticated token. If the user's token quota is
// exceeded, this call returns an error **QUOTA_EXCEEDED**.
func (a *TokensAPI) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {
	return a.impl.Create(ctx, request)
}

// Revoke token
//
// Revokes an access token.
//
// If a token with the specified ID is not valid, this call returns an error
// **RESOURCE_DOES_NOT_EXIST**.
func (a *TokensAPI) Delete(ctx context.Context, request RevokeTokenRequest) error {
	return a.impl.Delete(ctx, request)
}

// Revoke token
//
// Revokes an access token.
//
// If a token with the specified ID is not valid, this call returns an error
// **RESOURCE_DOES_NOT_EXIST**.
func (a *TokensAPI) DeleteByTokenId(ctx context.Context, tokenId string) error {
	return a.impl.Delete(ctx, RevokeTokenRequest{
		TokenId: tokenId,
	})
}

// List tokens
//
// Lists all the valid tokens for a user-workspace pair.
func (a *TokensAPI) List(ctx context.Context) (*ListTokensResponse, error) {
	return a.impl.List(ctx)
}
