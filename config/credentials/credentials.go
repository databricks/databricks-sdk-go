package credentials

import (
	"net/http"

	"golang.org/x/oauth2"
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

// OAuthCredentialsProvider is a specialized CredentialsProvider uses and provides an OAuth token.
type OAuthCredentialsProvider interface {
	CredentialsProvider
	// Token returns the OAuth token generated by the provider.
	Token() (*oauth2.Token, error)
}

type oauthCredentialsProvider struct {
	setHeaders func(r *http.Request) error
	token      func() (*oauth2.Token, error)
}

func (c *oauthCredentialsProvider) SetHeaders(r *http.Request) error {
	return c.setHeaders(r)
}

func (c *oauthCredentialsProvider) Token() (*oauth2.Token, error) {
	return c.token()
}

func NewOAuthCredentialsProvider(visitor func(r *http.Request) error, tokenProvider func() (*oauth2.Token, error)) OAuthCredentialsProvider {
	return &oauthCredentialsProvider{
		setHeaders: visitor,
		token:      tokenProvider,
	}
}
