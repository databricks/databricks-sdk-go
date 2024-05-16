package credentials

// OAuthToken represents an OAuth token as defined by the OAuth 2.0 Authorization Framework.
// https://datatracker.ietf.org/doc/html/rfc6749
type OAuthToken struct {
	// The access token issued by the authorization server. This is the token that will be used to authenticate requests.
	AccessToken string `json:"access_token"  auth:",sensitive"`
	// Time in seconds until the token expires.
	ExpiresIn int `json:"expires_in"`
	// The scope of the token. This is a space-separated list of strings that represent the permissions granted by the token.
	Scope string `json:"scope"`
	// The type of token that was issued.
	TokenType string `json:"token_type"`
}
