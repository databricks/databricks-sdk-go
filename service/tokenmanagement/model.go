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
    // The information for each token.
    TokenInfos []PublicTokenInfo `json:"token_infos,omitempty"`
}


type PublicTokenInfo struct {
    // Comment the token was created with, if applicable.
    Comment string `json:"comment,omitempty"`
    // Server time (in epoch milliseconds) when the token was created.
    CreationTime int64 `json:"creation_time,omitempty"`
    // Server time (in epoch milliseconds) when the token will expire, or -1 if
    // not applicable.
    ExpiryTime int64 `json:"expiry_time,omitempty"`
    // The ID of this token.
    TokenId string `json:"token_id,omitempty"`
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

// Comment that describes the purpose of the token, specified by the token
// creator.

// User ID of the user that created the token.

// Username of the user that created the token.

// Timestamp when the token was created

// Timestamp when the token expires

// User ID of the user that owns the token.

// ID of the token

