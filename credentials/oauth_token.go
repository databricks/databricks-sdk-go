package credentials

type OAuthToken struct {
	AccessToken string `json:"access_token"  auth:",sensitive"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
