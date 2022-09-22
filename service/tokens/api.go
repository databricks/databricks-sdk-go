// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewTokens(client *client.DatabricksClient) TokensService {
	return &TokensAPI{
		client: client,
	}
}

type TokensAPI struct {
	client *client.DatabricksClient
}

// Create a user token
//
// Creates and returns a token for a user.
//
// If this call is made through token authentication, it creates a token with
// the same client ID as the authenticated token. If the user's token quota is
// exceeded, this call returns an error **QUOTA_EXCEEDED**.
func (a *TokensAPI) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {
	var createTokenResponse CreateTokenResponse
	path := "/api/2.0/token/create"
	err := a.client.Post(ctx, path, request, &createTokenResponse)
	return &createTokenResponse, err
}

// Revoke token
//
// Revokes an access token.
//
// If a token with the specified ID is not valid, this call returns an error
// **RESOURCE_DOES_NOT_EXIST**.
func (a *TokensAPI) Delete(ctx context.Context, request RevokeTokenRequest) error {
	path := "/api/2.0/token/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
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
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token/list"
	err := a.client.Get(ctx, path, nil, &listTokensResponse)
	return &listTokensResponse, err
}
