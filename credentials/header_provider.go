package credentials

import "golang.org/x/oauth2"

type HeaderProvider interface {
	Token() (*oauth2.Token, error)
}
