// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

// all definitions in this file are in alphabetical order

type CreateOboTokenRequest struct {
	ApplicationId string `json:"application_id"`

	Comment string `json:"comment"`

	LifetimeSeconds any/* MISSING TYPE */ `json:"lifetime_seconds"`
}

type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`

	TokenValue string `json:"token_value,omitempty"`
}

type DeleteTokenRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" path:"token_id"`
}

type GetTokenInfoRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" path:"token_id"`
}

type ListAllTokensRequest struct {
	// User ID of the user that created the token.
	CreatedById string `json:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`
}

type ListTokensResponse struct {
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

type TokenInfo struct {
	Comment string `json:"comment,omitempty"`

	CreatedById int64 `json:"created_by_id,omitempty"`

	CreatedByUsername string `json:"created_by_username,omitempty"`

	CreationTime int64 `json:"creation_time,omitempty"`

	ExpiryTime int64 `json:"expiry_time,omitempty"`

	OwnerId int64 `json:"owner_id,omitempty"`

	TokenId string `json:"token_id,omitempty"`
}

// Application ID of the service principal.

// Comment that describes the purpose of the token, specified by the token
// creator.

// Comment that describes the purpose of the token.

// User ID of the user that created the token.

// Username of the user that created the token.

// Timestamp when the token was created.

// Timestamp when the token expires.

// The number of seconds before the token expires.

// User ID of the user that owns the token.

// ID of the token.

// Value of the token.
