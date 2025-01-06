package oauth_test

import (
	"context"
	"crypto/tls"
	_ "embed"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/credentials/oauth"
	"github.com/databricks/databricks-sdk-go/qa"
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
	p, err := oauth.NewPersistentAuth(context.Background(), oauth.WithTokenCache(cache))
	require.NoError(t, err)
	defer p.Close()
	arg, err := oauth.NewBasicAccountOAuthArgument("https://abc", "xyz")
	assert.NoError(t, err)
	tok, err := p.Load(context.Background(), arg)
	assert.NoError(t, err)
	assert.Equal(t, "bcd", tok.AccessToken)
	assert.Equal(t, "", tok.RefreshToken)
}

func useInsecureOAuthHttpClientForTests(ctx context.Context) context.Context {
	return context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
}

func TestLoadRefresh(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "POST",
			Resource: "/oidc/accounts/xyz/v1/token",
			Response: `access_token=refreshed&refresh_token=def`,
		},
	}.ApplyClient(t, func(ctx context.Context, c *client.DatabricksClient) {
		ctx = useInsecureOAuthHttpClientForTests(ctx)
		expectedKey := fmt.Sprintf("%s/oidc/accounts/xyz", c.Config.Host)
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
		p, err := oauth.NewPersistentAuth(context.Background(), oauth.WithTokenCache(cache))
		require.NoError(t, err)
		defer p.Close()
		arg, err := oauth.NewBasicAccountOAuthArgument(c.Config.Host, "xyz")
		assert.NoError(t, err)
		tok, err := p.Load(ctx, arg)
		assert.NoError(t, err)
		assert.Equal(t, "refreshed", tok.AccessToken)
		assert.Equal(t, "", tok.RefreshToken)
	})
}

func TestChallenge(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "POST",
			Resource: "/oidc/accounts/xyz/v1/token",
			Response: `access_token=__THAT__&refresh_token=__SOMETHING__`,
		},
	}.ApplyClient(t, func(ctx context.Context, c *client.DatabricksClient) {
		ctx = useInsecureOAuthHttpClientForTests(ctx)
		expectedKey := fmt.Sprintf("%s/oidc/accounts/xyz", c.Config.Host)

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
				assert.Equal(t, expectedKey, key)
				assert.Equal(t, "__SOMETHING__", tok.RefreshToken)
				return nil
			},
		}
		p, err := oauth.NewPersistentAuth(context.Background(), oauth.WithTokenCache(cache), oauth.WithBrowser(browser))
		require.NoError(t, err)
		defer p.Close()
		arg, err := oauth.NewBasicAccountOAuthArgument(c.Config.Host, "xyz")
		assert.NoError(t, err)

		errc := make(chan error)
		go func() {
			errc <- p.Challenge(ctx, arg)
		}()

		state := <-browserOpened
		resp, err := http.Get(fmt.Sprintf("http://localhost:8020?code=__THIS__&state=%s", state))
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, 200, resp.StatusCode)

		err = <-errc
		assert.NoError(t, err)
	})
}

func TestChallengeFailed(t *testing.T) {
	qa.HTTPFixtures{}.ApplyClient(t, func(ctx context.Context, c *client.DatabricksClient) {
		ctx = useInsecureOAuthHttpClientForTests(ctx)

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
		p, err := oauth.NewPersistentAuth(context.Background(), oauth.WithBrowser(browser))
		require.NoError(t, err)
		defer p.Close()
		arg, err := oauth.NewBasicAccountOAuthArgument(c.Config.Host, "xyz")
		assert.NoError(t, err)

		errc := make(chan error)
		go func() {
			errc <- p.Challenge(ctx, arg)
		}()

		<-browserOpened
		resp, err := http.Get(
			"http://localhost:8020?error=access_denied&error_description=Policy%20evaluation%20failed%20for%20this%20request")
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, 400, resp.StatusCode)

		err = <-errc
		assert.EqualError(t, err, "authorize: access_denied: Policy evaluation failed for this request")
	})
}
