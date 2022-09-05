// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
type TokenmanagementService interface {
	// Delete a token, specified by its ID.
	DeleteToken(ctx context.Context, deleteTokenRequest DeleteTokenRequest) error
	// Get a token, specified by its ID.
	GetTokenInfo(ctx context.Context, getTokenInfoRequest GetTokenInfoRequest) (*TokenInfo, error)
	// List all tokens belonging to a workspace or a user.
	ListAllTokens(ctx context.Context, listAllTokensRequest ListAllTokensRequest) (*ListTokensResponse, error)
	GetTokenInfoByTokenId(ctx context.Context, tokenId string) (*TokenInfo, error)
	DeleteTokenByTokenId(ctx context.Context, tokenId string) error
}

func New(client *client.DatabricksClient) TokenmanagementService {
	return &TokenmanagementAPI{
		client: client,
	}
}

type TokenmanagementAPI struct {
	client *client.DatabricksClient
}

// Delete a token, specified by its ID.
func (a *TokenmanagementAPI) DeleteToken(ctx context.Context, request DeleteTokenRequest) error {
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get a token, specified by its ID.
func (a *TokenmanagementAPI) GetTokenInfo(ctx context.Context, request GetTokenInfoRequest) (*TokenInfo, error) {
	var tokenInfo TokenInfo
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Get(ctx, path, request, &tokenInfo)
	return &tokenInfo, err
}

// List all tokens belonging to a workspace or a user.
func (a *TokenmanagementAPI) ListAllTokens(ctx context.Context, request ListAllTokensRequest) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token-management/tokens"
	err := a.client.Get(ctx, path, request, &listTokensResponse)
	return &listTokensResponse, err
}

func (a *TokenmanagementAPI) GetTokenInfoByTokenId(ctx context.Context, tokenId string) (*TokenInfo, error) {
	return a.GetTokenInfo(ctx, GetTokenInfoRequest{
		TokenId: tokenId,
	})
}

func (a *TokenmanagementAPI) DeleteTokenByTokenId(ctx context.Context, tokenId string) error {
	return a.DeleteToken(ctx, DeleteTokenRequest{
		TokenId: tokenId,
	})
}
