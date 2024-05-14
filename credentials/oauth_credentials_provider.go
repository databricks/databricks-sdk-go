package credentials

import (
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthCredentialsProvider interface {
	SetHeaders(r *http.Request) error
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
