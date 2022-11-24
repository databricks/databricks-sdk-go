package config

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

func serviceToServiceVisitor(inner, cloud oauth2.TokenSource, header string) func(r *http.Request) error {
	refreshableInner := oauth2.ReuseTokenSource(nil, inner)
	refreshableCloud := oauth2.ReuseTokenSource(nil, cloud)
	return func(r *http.Request) error {
		inner, err := refreshableInner.Token()
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
		inner, err := refreshableInner.Token()
		if err != nil {
			return fmt.Errorf("inner token: %w", err)
		}
		inner.SetAuthHeader(r)
		return nil
	}
}
