package u2m_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

type tokenCacheMock struct {
	store  func(key string, t *oauth2.Token) error
	lookup func(key string) (*oauth2.Token, error)
}

func (m *tokenCacheMock) Store(key string, t *oauth2.Token) error {
	if m.store == nil {
		panic("no store mock")
	}
	return m.store(key, t)
}

func (m *tokenCacheMock) Lookup(key string) (*oauth2.Token, error) {
	if m.lookup == nil {
		panic("no lookup mock")
	}
	return m.lookup(key)
}

func TestLoad(t *testing.T) {
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			assert.Equal(t, "https://abc/oidc/accounts/xyz", key)
			return &oauth2.Token{
				AccessToken: "bcd",
				Expiry:      time.Now().Add(1 * time.Minute),
			}, nil
		},
	}
	arg, err := u2m.NewBasicAccountOAuthArgument("https://abc", "xyz")
	assert.NoError(t, err)
	p, err := u2m.NewPersistentAuth(context.Background(), u2m.WithTokenCache(cache), u2m.WithOAuthArgument(arg))
	require.NoError(t, err)
	defer p.Close()
	tok, err := p.Token()
	assert.NoError(t, err)
	assert.Equal(t, "bcd", tok.AccessToken)
	assert.Equal(t, "", tok.RefreshToken)
}

type MockOAuthEndpointSupplier struct{}

func (m MockOAuthEndpointSupplier) GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*u2m.OAuthAuthorizationServer, error) {
	return &u2m.OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/accounts/%s/v1/authorize", accountHost, accountId),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/accounts/%s/v1/token", accountHost, accountId),
	}, nil
}

func (m MockOAuthEndpointSupplier) GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*u2m.OAuthAuthorizationServer, error) {
	return &u2m.OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/v1/authorize", workspaceHost),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/v1/token", workspaceHost),
	}, nil
}

func TestLoadRefresh(t *testing.T) {
	ctx := context.Background()
	expectedKey := "https://accounts.cloud.databricks.com/oidc/accounts/xyz"
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			assert.Equal(t, expectedKey, key)
			return &oauth2.Token{
				AccessToken:  "expired",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(-1 * time.Minute),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			assert.Equal(t, expectedKey, key)
			assert.Equal(t, "def", tok.RefreshToken)
			return nil
		},
	}
	arg, err := u2m.NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	assert.NoError(t, err)
	p, err := u2m.NewPersistentAuth(
		ctx,
		u2m.WithTokenCache(cache),
		u2m.WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=refreshed&refresh_token=def`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		u2m.WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		u2m.WithOAuthArgument(arg),
	)
	require.NoError(t, err)
	defer p.Close()
	tok, err := p.Token()
	assert.NoError(t, err)
	assert.Equal(t, "refreshed", tok.AccessToken)
	assert.Equal(t, "", tok.RefreshToken)
}

func TestChallenge(t *testing.T) {
	ctx := context.Background()

	browserOpened := make(chan string)
	browser := func(redirect string) error {
		u, err := url.ParseRequestURI(redirect)
		if err != nil {
			return err
		}
		assert.Equal(t, "/oidc/accounts/xyz/v1/authorize", u.Path)
		// for now we're ignoring asserting the fields of the redirect
		query := u.Query()
		browserOpened <- query.Get("state")
		return nil
	}
	cache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			assert.Equal(t, "https://accounts.cloud.databricks.com/oidc/accounts/xyz", key)
			assert.Equal(t, "__THAT__", tok.AccessToken)
			assert.Equal(t, "__SOMETHING__", tok.RefreshToken)
			return nil
		},
	}
	arg, err := u2m.NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	assert.NoError(t, err)
	p, err := u2m.NewPersistentAuth(
		ctx,
		u2m.WithTokenCache(cache),
		u2m.WithBrowser(browser),
		u2m.WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=__THAT__&refresh_token=__SOMETHING__`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		u2m.WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		u2m.WithOAuthArgument(arg),
	)
	require.NoError(t, err)
	defer p.Close()

	errc := make(chan error)
	go func() {
		err := p.Challenge()
		errc <- err
		close(errc)
	}()

	state := <-browserOpened
	resp, err := http.Get(fmt.Sprintf("http://localhost:8020?code=__THIS__&state=%s", state))
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	err = <-errc
	assert.NoError(t, err)
}

func TestChallengeFailed(t *testing.T) {
	ctx := context.Background()
	browserOpened := make(chan string)
	browser := func(redirect string) error {
		u, err := url.ParseRequestURI(redirect)
		if err != nil {
			return err
		}
		assert.Equal(t, "/oidc/accounts/xyz/v1/authorize", u.Path)
		// for now we're ignoring asserting the fields of the redirect
		query := u.Query()
		browserOpened <- query.Get("state")
		return nil
	}
	arg, err := u2m.NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	assert.NoError(t, err)
	p, err := u2m.NewPersistentAuth(ctx, u2m.WithBrowser(browser), u2m.WithOAuthArgument(arg))
	require.NoError(t, err)
	defer p.Close()

	errc := make(chan error)
	go func() {
		err := p.Challenge()
		errc <- err
		close(errc)
	}()

	<-browserOpened
	resp, err := http.Get(
		"http://localhost:8020?error=access_denied&error_description=Policy%20evaluation%20failed%20for%20this%20request")
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)

	err = <-errc
	assert.EqualError(t, err, "authorize: access_denied: Policy evaluation failed for this request")
}
