// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

// all definitions in this file are in alphabetical order

type CreateTokenRequest struct {
    // Optional description to attach to the token. 
    Comment string `json:"comment,omitempty"`
    // The lifetime of the token, in seconds. If no lifetime is specified, then 
    // this token remains valid indefinitely. If the token in interal, the 
    // lifetime is at most autoInjectedTokenExpirationTime (default 2 days) in 
    // WebappConf. 
    LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`
}


type CreateTokenResponse struct {
    // The information of the newly-created token. 
    TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
    // The value of the newly-created token. 
    TokenValue string `json:"token_value,omitempty"`
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


type RevokeTokenRequest struct {
    // The ID of the token to be revoked. 
    TokenId string `json:"token_id"`
}

