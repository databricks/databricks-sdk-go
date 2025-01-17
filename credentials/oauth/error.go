package oauth

// InvalidRefreshTokenError is returned from PersistentAuth's Load() method
// if the access token has expired and the refresh token in the token cache
// is invalid.
type InvalidRefreshTokenError struct {
	err error
}

func (e *InvalidRefreshTokenError) Error() string {
	return e.err.Error()
}

func (e *InvalidRefreshTokenError) Unwrap() error {
	return e.err
}

var _ error = &InvalidRefreshTokenError{}
