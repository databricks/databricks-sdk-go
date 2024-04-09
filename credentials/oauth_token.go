package credentials

type OAuthToken struct {
	AccessToken          string                 `json:"access_token"  auth:",sensitive"`
	AuthorizationDetails []AuthorizationDetails `json:"authorization_details"`
	ExpiresIn            int                    `json:"expires_in"`
	Scope                string                 `json:"scope"`
	TokenType            string                 `json:"token_type"`
}
