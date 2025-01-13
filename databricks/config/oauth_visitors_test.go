package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

type mockTokenSource struct {
	mockedTokenFunc func() (*oauth2.Token, error)
}

func (m mockTokenSource) Token() (*oauth2.Token, error) {
	return m.mockedTokenFunc()
}

func TestAzureReuseTokenSource(t *testing.T) {
	mockSource := mockTokenSource{
		mockedTokenFunc: func() (*oauth2.Token, error) {
			return &oauth2.Token{
				Expiry: time.Now().Add(35 * time.Second),
			}, nil
		},
	}

	// Assert the token is not valid if it expires in 35 seconds.
	adjustedSource := azureReuseTokenSource(nil, mockSource)
	token, err := adjustedSource.Token()
	assert.NoError(t, err)
	assert.False(t, token.Valid())
}
