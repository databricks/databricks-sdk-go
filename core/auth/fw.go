package auth

import (
	"net/http"

	"golang.org/x/oauth2"
)

type RequestVisitor interface {
	Visit(*http.Request) error
}

type Credentials interface {
	Name() string
	Configure() (RequestVisitor, error)
}

type refreshableVisitor struct {
	Header string
	oauth2.TokenSource
}

func (rv refreshableVisitor) Visit(r *http.Request) error {
	token, err := rv.Token() // this also refreshes
	if err != nil {
		return err
	}
	if rv.Header == "" {
		rv.Header = "Authorization" // todo
	}
	r.Header.Set(rv.Header, token.Type()+" "+token.AccessToken)
	return nil
}
