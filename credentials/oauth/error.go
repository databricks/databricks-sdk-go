package oauth

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
