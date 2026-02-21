package config

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/authconv"
	"golang.org/x/oauth2"
)

// cacheOptions returns the auth.Option values derived from the given Config.
func cacheOptions(cfg *Config) []auth.Option {
	return []auth.Option{
		auth.WithAsyncRefresh(!cfg.DisableAsyncTokenRefresh),
	}
}

// serviceToServiceVisitor returns a visitor that sets the Authorization header
// to the token from the auth token source and the provided secondary header to
// the token from the secondary token source.
func serviceToServiceVisitor(primary, secondary oauth2.TokenSource, secondaryHeader string, opts ...auth.Option) func(r *http.Request) error {
	refreshableAuth := auth.NewCachedTokenSource(authconv.AuthTokenSource(primary), opts...)
	refreshableSecondary := auth.NewCachedTokenSource(authconv.AuthTokenSource(secondary), opts...)
	return func(r *http.Request) error {
		inner, err := refreshableAuth.Token(context.Background())
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)

		cloud, err := refreshableSecondary.Token(context.Background())
		if err != nil {
			return fmt.Errorf("cloud token: %w", err)
		}
		r.Header.Set(secondaryHeader, cloud.AccessToken)
		return nil
	}
}

// The same as serviceToServiceVisitor, but without a secondary token source.
func refreshableVisitor(inner oauth2.TokenSource, opts ...auth.Option) func(r *http.Request) error {
	return refreshableAuthVisitor(authconv.AuthTokenSource(inner), opts...)
}

// The same as serviceToServiceVisitor, but without a secondary token source.
func refreshableAuthVisitor(inner auth.TokenSource, opts ...auth.Option) func(r *http.Request) error {
	cts := auth.NewCachedTokenSource(inner, opts...)
	return func(r *http.Request) error {
		inner, err := cts.Token(context.Background())
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)
		return nil
	}
}

func azureVisitor(cfg *Config, inner func(*http.Request) error) func(*http.Request) error {
	return func(r *http.Request) error {
		if cfg.AzureResourceID != "" {
			r.Header.Set(xDatabricksAzureWorkspaceResourceId, cfg.AzureResourceID)
		}
		return inner(r)
	}
}

// azureReuseTokenSource returns a cached token source that refreshes token 40
// seconds before they expire. The reason for this is that Azure Databricks
// rejects tokens that expire in 30 seconds or less and we want to give a 10
// second buffer.
func azureReuseTokenSource(t *oauth2.Token, ts oauth2.TokenSource, opts ...auth.Option) oauth2.TokenSource {
	early := wrap(ts, func(t *oauth2.Token) *oauth2.Token {
		t.Expiry = t.Expiry.Add(-40 * time.Second)
		return t
	})

	allOpts := append([]auth.Option{auth.WithCachedToken(t)}, opts...)
	return authconv.OAuth2TokenSource(auth.NewCachedTokenSource(
		authconv.AuthTokenSource(early),
		allOpts...,
	))
}

func wrap(ts oauth2.TokenSource, fn func(*oauth2.Token) *oauth2.Token) oauth2.TokenSource {
	return &tokenSourceWrapper{fn: fn, inner: ts}
}

type tokenSourceWrapper struct {
	fn    func(*oauth2.Token) *oauth2.Token
	inner oauth2.TokenSource
}

func (w *tokenSourceWrapper) Token() (*oauth2.Token, error) {
	t, err := w.inner.Token()
	if err != nil {
		return nil, err
	}
	return w.fn(t), nil
}
