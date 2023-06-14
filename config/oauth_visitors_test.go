package config

import (
	"context"
	"errors"
	"testing"

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
	mockSource := mockTokenSource{
		mockedTokenFunc: func() (*oauth2.Token, error) {
			return nil, errors.New("throttled") // this goes to retry but we need to mock poll or keep call count else test would timeout
		},
	}
	token, err := retriableTokenSource(context.Background(), mockSource)
	assert.Nil(t, token)
	assert.Contains(t, err.Error(), "throttled")
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
