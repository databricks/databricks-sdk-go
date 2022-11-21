// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewTokenManagement(client *client.DatabricksClient) *TokenManagementAPI {
	return &TokenManagementAPI{
		TokenManagementService: &tokenManagementAPI{
			client: client,
		},
	}
}

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
type TokenManagementAPI struct {
	// TokenManagementService contains low-level REST API interface.
	TokenManagementService
}

// Create on-behalf token
//
// Creates a token on behalf of a service principal.
func (a *TokenManagementAPI) CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error) {
	return a.TokenManagementService.CreateOboToken(ctx, request)
}

// Delete a token
//
// Deletes a token, specified by its ID.
func (a *TokenManagementAPI) DeleteToken(ctx context.Context, request DeleteTokenRequest) error {
	return a.TokenManagementService.DeleteToken(ctx, request)
}

// Delete a token
//
// Deletes a token, specified by its ID.
func (a *TokenManagementAPI) DeleteTokenByTokenId(ctx context.Context, tokenId string) error {
	return a.DeleteToken(ctx, DeleteTokenRequest{
		TokenId: tokenId,
	})
}

// Get token info
//
// Gets information about a token, specified by its ID.
func (a *TokenManagementAPI) GetTokenInfo(ctx context.Context, request GetTokenInfoRequest) (*TokenInfo, error) {
	return a.TokenManagementService.GetTokenInfo(ctx, request)
}

// Get token info
//
// Gets information about a token, specified by its ID.
func (a *TokenManagementAPI) GetTokenInfoByTokenId(ctx context.Context, tokenId string) (*TokenInfo, error) {
	return a.GetTokenInfo(ctx, GetTokenInfoRequest{
		TokenId: tokenId,
	})
}

// List all tokens
//
// Lists all tokens associated with the specified workspace or user.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TokenManagementAPI) ListTokensAll(ctx context.Context, request ListTokensRequest) ([]TokenInfo, error) {
	response, err := a.ListTokens(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.TokenInfos, nil
}

// unexported type that holds implementations of just TokenManagement API methods
type tokenManagementAPI struct {
	client *client.DatabricksClient
}

func (a *tokenManagementAPI) CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error) {
	var createOboTokenResponse CreateOboTokenResponse
	path := "/api/2.0/token-management/on-behalf-of/tokens"
	err := a.client.Post(ctx, path, request, &createOboTokenResponse)
	return &createOboTokenResponse, err
}

func (a *tokenManagementAPI) DeleteToken(ctx context.Context, request DeleteTokenRequest) error {
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *tokenManagementAPI) GetTokenInfo(ctx context.Context, request GetTokenInfoRequest) (*TokenInfo, error) {
	var tokenInfo TokenInfo
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Get(ctx, path, request, &tokenInfo)
	return &tokenInfo, err
}

func (a *tokenManagementAPI) ListTokens(ctx context.Context, request ListTokensRequest) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token-management/tokens"
	err := a.client.Get(ctx, path, request, &listTokensResponse)
	return &listTokensResponse, err
}
