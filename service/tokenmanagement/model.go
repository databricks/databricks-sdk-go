// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

// all definitions in this file are in alphabetical order

type DeleteTokenRequest struct {
	// The ID of the token to get.
	TokenId string ` path:"token_id"`
}

type GetTokenInfoRequest struct {
	// The ID of the token to get.
	TokenId string ` path:"token_id"`
}

type ListAllTokensRequest struct {
	// User ID of the user that created the token.
	CreatedById string ` url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string ` url:"created_by_username,omitempty"`
}

type ListTokensResponse struct {
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment string `json:"comment,omitempty"`
	// User ID of the user that created the token.
	CreatedById int64 `json:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"created_by_username,omitempty"`
	// Timestamp when the token was created
	CreationTime int64 `json:"creation_time,omitempty"`
	// Timestamp when the token expires
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// User ID of the user that owns the token.
	OwnerId int64 `json:"owner_id,omitempty"`
	// ID of the token
	TokenId string `json:"token_id,omitempty"`
}
