package config

import (
	"context"
	"errors"
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

func TestOAuthWithRetry(t *testing.T) {
	triedOnce := false
	mockSource := mockTokenSource{
		mockedTokenFunc: func() (*oauth2.Token, error) {
			if triedOnce == true {
				return nil, errors.New("retried")
			}
			triedOnce = true
			return nil, errors.New("throttled")
		},
	}
	token, err := retriableTokenSource(context.Background(), mockSource)
	assert.Nil(t, token)
	assert.Contains(t, err.Error(), "retried")
}

func TestOAuthWithoutRetry(t *testing.T) {
	mockSource := mockTokenSource{
		mockedTokenFunc: func() (*oauth2.Token, error) {
			return nil, errors.New("halt")
		},
	}
	token, err := retriableTokenSource(context.Background(), mockSource)
	assert.Nil(t, token)
	assert.Contains(t, err.Error(), "halt")
}

func TestOAuthWithValidToken(t *testing.T) {
	mockSource := mockTokenSource{
		mockedTokenFunc: func() (*oauth2.Token, error) {
			testToken := oauth2.Token{}
			testToken.TokenType = "test"
			return &testToken, nil
		},
	}
	token, err := retriableTokenSource(context.Background(), mockSource)
	assert.Equal(t, "test", token.TokenType)
	assert.NoError(t, err)
}

func TestAzureAdjustExpiry(t *testing.T) {
	mockSource := mockTokenSource{
		mockedTokenFunc: func() (*oauth2.Token, error) {
			return &oauth2.Token{
				Expiry: time.Now().Add(35 * time.Second),
			}, nil
		},
	}

	// Assert the token is not valid if it expires in 35 seconds.
	adjustedSource := azureAdjustExpiry(mockSource)
	token, err := adjustedSource.Token()
	assert.NoError(t, err)
	assert.False(t, token.Valid())
}
