// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewTokenManagement(client *client.DatabricksClient) TokenManagementService {
	return &TokenManagementAPI{
		client: client,
	}
}

type TokenManagementAPI struct {
	client *client.DatabricksClient
}

// Create a token on behalf of a service principal.
func (a *TokenManagementAPI) CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error) {
	var createOboTokenResponse CreateOboTokenResponse
	path := "/api/2.0/token-management/on-behalf-of/tokens"
	err := a.client.Post(ctx, path, request, &createOboTokenResponse)
	return &createOboTokenResponse, err
}

// Delete a token, specified by its ID.
func (a *TokenManagementAPI) DeleteToken(ctx context.Context, request DeleteTokenRequest) error {
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a token, specified by its ID.
func (a *TokenManagementAPI) DeleteTokenByTokenId(ctx context.Context, tokenId string) error {
	return a.DeleteToken(ctx, DeleteTokenRequest{
		TokenId: tokenId,
	})
}

// Get a token, specified by its ID.
func (a *TokenManagementAPI) GetTokenInfo(ctx context.Context, request GetTokenInfoRequest) (*TokenInfo, error) {
	var tokenInfo TokenInfo
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Get(ctx, path, request, &tokenInfo)
	return &tokenInfo, err
}

// Get a token, specified by its ID.
func (a *TokenManagementAPI) GetTokenInfoByTokenId(ctx context.Context, tokenId string) (*TokenInfo, error) {
	return a.GetTokenInfo(ctx, GetTokenInfoRequest{
		TokenId: tokenId,
	})
}

// List all tokens belonging to a workspace or a user.
func (a *TokenManagementAPI) ListAllTokens(ctx context.Context, request ListAllTokensRequest) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token-management/tokens"
	err := a.client.Get(ctx, path, request, &listTokensResponse)
	return &listTokensResponse, err
}
