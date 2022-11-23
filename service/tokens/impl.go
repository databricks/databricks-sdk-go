// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

import (
	"context"
	"net/http"
)

type databricksClient interface {
	Do(ctx context.Context, method string, path string, request any, response any) error
	ConfiguredAccountID() string
	IsAws() bool
}

// unexported type that holds implementations of just Tokens API methods
type tokensImpl struct {
	client databricksClient
}

func (a *tokensImpl) Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error) {
	var createTokenResponse CreateTokenResponse
	path := "/api/2.0/token/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createTokenResponse)
	return &createTokenResponse, err
}

func (a *tokensImpl) Delete(ctx context.Context, request RevokeTokenRequest) error {
	path := "/api/2.0/token/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *tokensImpl) List(ctx context.Context) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token/list"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listTokensResponse)
	return &listTokensResponse, err
}
