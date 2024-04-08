package config

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// serviceToServiceVisitor returns a visitor that sets the Authorization header to the token from the auth token source
// and the provided secondary header to the token from the secondary token source.
func serviceToServiceVisitor(auth, secondary oauth2.TokenSource, secondaryHeader string) func(r *http.Request) error {
	refreshableAuth := oauth2.ReuseTokenSource(nil, auth)
	refreshableSecondary := oauth2.ReuseTokenSource(nil, secondary)
	return func(r *http.Request) error {
		inner, err := refreshableAuth.Token()
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)

		cloud, err := refreshableSecondary.Token()
		if err != nil {
			return fmt.Errorf("cloud token: %w", err)
		}
		r.Header.Set(secondaryHeader, cloud.AccessToken)
		return nil
	}
}

// type HeaderFactory func() http.Header

// func (h *HeaderFactory) Token() (*oauth2.Token, error) {
// 	return nil, nil
// }

// The same as serviceToServiceVisitor, but without a secondary token source.
func refreshableVisitor(inner oauth2.TokenSource) func(r *http.Request) error {
	refreshableAuth := oauth2.ReuseTokenSource(nil, inner)
	return func(r *http.Request) error {
		inner, err := refreshableAuth.Token()
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

// azureReuseTokenSource calls into oauth2.ReuseTokenSourceWithExpiry with a 40 second expiry window.
// By default, the oauth2 library refreshes a token 10 seconds before it expires.
// Azure Databricks rejects tokens that expire in 30 seconds or less.
// We combine these and refresh the token 40 seconds before it expires.
func azureReuseTokenSource(t *oauth2.Token, ts oauth2.TokenSource) oauth2.TokenSource {
	return oauth2.ReuseTokenSourceWithExpiry(t, ts, 40*time.Second)
}
