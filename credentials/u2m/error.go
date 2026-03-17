package u2m

import "errors"

// ErrMissingRefreshToken is returned when a token refresh is requested but the
// cached OAuth token does not include a refresh token.
var ErrMissingRefreshToken = errors.New("cached token has no refresh token")

// InvalidRefreshTokenError is returned from PersistentAuth's Load() method
// if the access token has expired and the refresh token in the token cache
// is invalid.
type InvalidRefreshTokenError struct {
	error
}
