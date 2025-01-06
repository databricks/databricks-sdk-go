package oauth

import (
	"context"
	"fmt"
	"strings"
)

// OAuthArgument is an interface that provides the necessary information to
// authenticate with PersistentAuth.
type OAuthArgument interface {
	// GetCacheKey returns a unique key for the OAuthArgument. This key is used
	// to store and retrieve the token from the token cache.
	GetCacheKey(ctx context.Context) string
}

type WorkspaceOAuthArgument interface {
	OAuthArgument

	// GetWorkspaceHost returns the host of the workspace to authenticate to.
	GetWorkspaceHost(ctx context.Context) string
}

type BasicWorkspaceOAuthArgument struct {
	// host is the host of the workspace to authenticate to. This must start
	// with "https://" and must not have a trailing slash.
	host string
}

func NewBasicWorkspaceOAuthArgument(host string) (BasicWorkspaceOAuthArgument, error) {
	if !strings.HasPrefix(host, "https://") {
		return BasicWorkspaceOAuthArgument{}, fmt.Errorf("host must start with 'https://': %s", host)
	}
	if strings.HasSuffix(host, "/") {
		return BasicWorkspaceOAuthArgument{}, fmt.Errorf("host must not have a trailing slash: %s", host)
	}
	return BasicWorkspaceOAuthArgument{host: host}, nil
}

func (a BasicWorkspaceOAuthArgument) GetWorkspaceHost(ctx context.Context) string {
	return a.host
}

// key is currently used for two purposes: OIDC URL prefix and token cache key.
// once we decide to start storing scopes in the token cache, we should change
// this approach.
func (a BasicWorkspaceOAuthArgument) GetCacheKey(ctx context.Context) string {
	a.host = strings.TrimSuffix(a.host, "/")
	if !strings.HasPrefix(a.host, "http") {
		a.host = fmt.Sprintf("https://%s", a.host)
	}
	return a.host
}

var _ WorkspaceOAuthArgument = BasicWorkspaceOAuthArgument{}

type AccountOAuthArgument interface {
	OAuthArgument

	// GetAccountHost returns the host of the account to authenticate to.
	GetAccountHost(ctx context.Context) string

	// GetAccountId returns the account ID of the account to authenticate to.
	GetAccountId(ctx context.Context) string
}

type BasicAccountOAuthArgument struct {
	accountHost string
	accountID   string
}

var _ AccountOAuthArgument = BasicAccountOAuthArgument{}

func NewBasicAccountOAuthArgument(accountsHost, accountID string) (BasicAccountOAuthArgument, error) {
	if !strings.HasPrefix(accountsHost, "https://") {
		return BasicAccountOAuthArgument{}, fmt.Errorf("accountsHost must start with 'https://': %s", accountsHost)
	}
	if strings.HasSuffix(accountsHost, "/") {
		return BasicAccountOAuthArgument{}, fmt.Errorf("accountsHost must not have a trailing slash: %s", accountsHost)
	}
	return BasicAccountOAuthArgument{accountHost: accountsHost, accountID: accountID}, nil
}

func (a BasicAccountOAuthArgument) GetAccountHost(ctx context.Context) string {
	return a.accountHost
}

func (a BasicAccountOAuthArgument) GetAccountId(ctx context.Context) string {
	return a.accountID
}

// key is currently used for two purposes: OIDC URL prefix and token cache key.
// once we decide to start storing scopes in the token cache, we should change
// this approach.
func (a BasicAccountOAuthArgument) GetCacheKey(ctx context.Context) string {
	return fmt.Sprintf("%s/oidc/accounts/%s", a.accountHost, a.accountID)
}
