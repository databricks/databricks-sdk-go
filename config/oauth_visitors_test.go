package config

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

type staticTokenSource struct {
	token *oauth2.Token
	err   error
}

func (s *staticTokenSource) Token() (*oauth2.Token, error) {
	return s.token, s.err
}

func TestServiceToServiceVisitorWithFallback_BothSucceed(t *testing.T) {
	primary := &staticTokenSource{token: &oauth2.Token{AccessToken: "primary-token"}}
	secondary := &staticTokenSource{token: &oauth2.Token{AccessToken: "secondary-token"}}
	visitor := serviceToServiceVisitor(primary, secondary, "X-Secondary", true)

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	err = visitor(req)
	require.NoError(t, err)
	assert.Equal(t, "Bearer primary-token", req.Header.Get("Authorization"))
	assert.Equal(t, "secondary-token", req.Header.Get("X-Secondary"))
}

func TestServiceToServiceVisitorWithFallback_SecondaryFails_SkipsHeader(t *testing.T) {
	primary := &staticTokenSource{token: &oauth2.Token{AccessToken: "primary-token"}}
	secondary := &staticTokenSource{err: fmt.Errorf("secondary failed")}
	visitor := serviceToServiceVisitor(primary, secondary, "X-Secondary", true)

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	err = visitor(req)
	require.NoError(t, err)
	assert.Equal(t, "Bearer primary-token", req.Header.Get("Authorization"))
	assert.Empty(t, req.Header.Get("X-Secondary"))
}

func TestServiceToServiceVisitorWithFallback_PrimaryFails_ReturnsError(t *testing.T) {
	primary := &staticTokenSource{err: fmt.Errorf("primary failed")}
	secondary := &staticTokenSource{token: &oauth2.Token{AccessToken: "secondary-token"}}
	visitor := serviceToServiceVisitor(primary, secondary, "X-Secondary", true)

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	err = visitor(req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "inner token")
}
