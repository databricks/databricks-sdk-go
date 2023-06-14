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
	return retries.Poll[oauth2.Token](ctx, 1*time.Minute, func() (*oauth2.Token, *retries.Err) {
		token, err := ts.Token()

		if strings.Contains(err.Error(), "throttled") {
			return nil, retries.Continue(err)
		}

		return token, nil
	})
}

func serviceToServiceVisitor(inner, cloud oauth2.TokenSource, header string) func(r *http.Request) error {
	refreshableInner := oauth2.ReuseTokenSource(nil, inner)
	refreshableCloud := oauth2.ReuseTokenSource(nil, cloud)
	return func(r *http.Request) error {
		inner, err := retriableTokenSource(context.Background(), refreshableInner)
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)
		cloud, err := refreshableCloud.Token()
		if err != nil {
			return fmt.Errorf("cloud token: %w", err)
		}
		r.Header.Set(header, cloud.AccessToken)
		return nil
	}
}

func refreshableVisitor(inner oauth2.TokenSource) func(r *http.Request) error {
	refreshableInner := oauth2.ReuseTokenSource(nil, inner)
	return func(r *http.Request) error {
		inner, err := retriableTokenSource(context.Background(), refreshableInner)
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)
		return nil
	}
}
