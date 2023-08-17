package config

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
	"golang.org/x/oauth2"
)

func retriableTokenSource(ctx context.Context, ts oauth2.TokenSource) (*oauth2.Token, error) {
	return retries.Poll(ctx, 1*time.Minute, func() (*oauth2.Token, *retries.Err) {
		token, err := ts.Token()
		if err == nil {
			return token, nil
		}
		var retryKeywords = []string{"throttled", "too many requests", "429", "request limit exceeded", "rate limit"}
		for _, retryKeyword := range retryKeywords {
			if strings.Contains(err.Error(), retryKeyword) {
				return nil, retries.Continue(err)
			}
		}
		return nil, retries.Halt(err)
	})
}

// serviceToServiceVisitor returns a visitor that sets the Authorization header to the token from the auth token source
// and the provided secondary header to the token from the secondary token source. If secondary is nil, the secondary
// header is not set. If tolerateSecondaryFailure is true, the visitor will not return an error if the cloud token source fails.
func serviceToServiceVisitor(auth, secondary oauth2.TokenSource, secondaryHeader string, tolerateSecondaryFailure bool) func(r *http.Request) error {
	refreshableAuth := oauth2.ReuseTokenSource(nil, auth)
	var refreshableSecondary oauth2.TokenSource
	if secondary != nil {
		refreshableSecondary = oauth2.ReuseTokenSource(nil, secondary)
	}
	return func(r *http.Request) error {
		inner, err := retriableTokenSource(r.Context(), refreshableAuth)
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)

		if refreshableSecondary == nil {
			return nil
		}
		cloud, err := retriableTokenSource(r.Context(), refreshableSecondary)
		if err == nil {
			r.Header.Set(secondaryHeader, cloud.AccessToken)
		} else if !tolerateSecondaryFailure {
			return fmt.Errorf("cloud token: %w", err)
		}
		return nil
	}
}

func refreshableVisitor(inner oauth2.TokenSource) func(r *http.Request) error {
	return serviceToServiceVisitor(inner, nil, "", false)
}

func azureVisitor(workspaceResourceId string, inner func(*http.Request) error) func(*http.Request) error {
	return func(r *http.Request) error {
		if workspaceResourceId != "" {
			r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", workspaceResourceId)
		}
		return inner(r)
	}
}
