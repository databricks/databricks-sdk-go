package credentials

import (
	"net/http"
)

type CredentialsProvider interface {
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
