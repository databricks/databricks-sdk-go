// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewTokens(client *client.DatabricksClient) *TokensAPI {
	return &TokensAPI{
		TokensService: &tokensAPI{
			client: client,
		},
	}
}

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
type TokensAPI struct {
	// TokensService contains low-level REST API interface.
	TokensService
}

// Create a user token
//
// Creates and returns a token for a user.
//
// If this call is made through token authentication, it creates a token with
// the same client ID as the authenticated token. If the user's token quota is
// exceeded, this call returns an error **QUOTA_EXCEEDED**.
func (a *TokensAPI) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {
	return a.TokensService.Create(ctx, request)
}

// Revoke token
//
// Revokes an access token.
//
// If a token with the specified ID is not valid, this call returns an error
// **RESOURCE_DOES_NOT_EXIST**.
func (a *TokensAPI) Delete(ctx context.Context, request RevokeTokenRequest) error {
	return a.TokensService.Delete(ctx, request)
}

// Revoke token
//
// Revokes an access token.
//
// If a token with the specified ID is not valid, this call returns an error
// **RESOURCE_DOES_NOT_EXIST**.
func (a *TokensAPI) DeleteByTokenId(ctx context.Context, tokenId string) error {
	return a.Delete(ctx, RevokeTokenRequest{
		TokenId: tokenId,
	})
}

// List tokens
//
// Lists all the valid tokens for a user-workspace pair.
func (a *TokensAPI) List(ctx context.Context) (*ListTokensResponse, error) {
	return a.TokensService.List(ctx)
}

// unexported type that holds implementations of just Tokens API methods
type tokensAPI struct {
	client *client.DatabricksClient
}

func (a *tokensAPI) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {
	var createTokenResponse CreateTokenResponse
	path := "/api/2.0/token/create"
	err := a.client.Post(ctx, path, request, &createTokenResponse)
	return &createTokenResponse, err
}

func (a *tokensAPI) Delete(ctx context.Context, request RevokeTokenRequest) error {
	path := "/api/2.0/token/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *tokensAPI) List(ctx context.Context) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token/list"
	err := a.client.Get(ctx, path, nil, &listTokensResponse)
	return &listTokensResponse, err
}
