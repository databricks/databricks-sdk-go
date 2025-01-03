package oauth

import (
	"context"
	"fmt"
	"strings"
)

// OAuthArgument is an interface that provides the necessary information to
// authenticate with PersistentAuth.
type OAuthArgument interface {
	// GetHost returns the host of the account or workspace to authenticate to.
	GetHost(ctx context.Context) string

	// GetAccountId returns the account ID of the account to authenticate to.
	GetAccountId(ctx context.Context) string

	// GetCacheKey returns the key to use for caching the token. On Challenge,
	// this key is used to store the token. On Load, this key is used to lookup
	// the token.
	GetCacheKey(ctx context.Context) string
}

type BasicOAuthArgument struct {
	Host      string
	AccountID string
}

var _ OAuthArgument = BasicOAuthArgument{}

func (a BasicOAuthArgument) GetHost(ctx context.Context) string {
	return a.Host
}

func (a BasicOAuthArgument) GetAccountId(ctx context.Context) string {
	return a.AccountID
}

// key is currently used for two purposes: OIDC URL prefix and token cache key.
// once we decide to start storing scopes in the token cache, we should change
// this approach.
func (a BasicOAuthArgument) GetCacheKey(ctx context.Context) string {
	a.Host = strings.TrimSuffix(a.Host, "/")
	if !strings.HasPrefix(a.Host, "http") {
		a.Host = fmt.Sprintf("https://%s", a.Host)
	}
	if a.AccountID != "" {
		return fmt.Sprintf("%s/oidc/accounts/%s", a.Host, a.AccountID)
	}
	return a.Host
}
