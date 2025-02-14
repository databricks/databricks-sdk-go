package u2m

// InvalidRefreshTokenError is returned from PersistentAuth's Load() method
// if the access token has expired and the refresh token in the token cache
// is invalid.
type InvalidRefreshTokenError struct {
	error
}
