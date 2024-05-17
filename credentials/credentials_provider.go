package credentials

import (
	"net/http"
)

// CredentialsProvider is an interface for providing credentials to the client.
// Implementations of this interface should set the necessary headers on the request.
type CredentialsProvider interface {
	// SetHeaders sets the necessary headers on the request.
	SetHeaders(r *http.Request) error
}

type credentialsProvider struct {
	setHeaders func(r *http.Request) error
}

func (c *credentialsProvider) SetHeaders(r *http.Request) error {
	return c.setHeaders(r)
}

func NewCredentialsProvider(visitor func(r *http.Request) error) CredentialsProvider {
	return &credentialsProvider{
		setHeaders: visitor,
	}
}
