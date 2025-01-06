package oauth

import (
	"context"
	"fmt"
	"strings"
)

// OAuthArgument is an interface that provides the necessary information to
// authenticate with PersistentAuth. Implementations of this interface must
// implement either the WorkspaceOAuthArgument or AccountOAuthArgument
// interface.
type OAuthArgument interface {
	// GetCacheKey returns a unique key for the OAuthArgument. This key is used
	// to store and retrieve the token from the token cache.
	GetCacheKey(ctx context.Context) string
}

// WorkspaceOAuthArgument is an interface that provides the necessary information
// to authenticate using OAuth to a specific workspace.
type WorkspaceOAuthArgument interface {
	OAuthArgument

	// GetWorkspaceHost returns the host of the workspace to authenticate to.
	GetWorkspaceHost(ctx context.Context) string
}

// BasicWorkspaceOAuthArgument is a basic implementation of the WorkspaceOAuthArgument
// interface that links each host with exactly one OAuth token.
type BasicWorkspaceOAuthArgument struct {
	// host is the host of the workspace to authenticate to. This must start
	// with "https://" and must not have a trailing slash.
	host string
}

// NewBasicWorkspaceOAuthArgument creates a new BasicWorkspaceOAuthArgument.
func NewBasicWorkspaceOAuthArgument(host string) (BasicWorkspaceOAuthArgument, error) {
	if !strings.HasPrefix(host, "https://") {
		return BasicWorkspaceOAuthArgument{}, fmt.Errorf("host must start with 'https://': %s", host)
	}
	if strings.HasSuffix(host, "/") {
		return BasicWorkspaceOAuthArgument{}, fmt.Errorf("host must not have a trailing slash: %s", host)
	}
	return BasicWorkspaceOAuthArgument{host: host}, nil
}

// GetWorkspaceHost returns the host of the workspace to authenticate to.
func (a BasicWorkspaceOAuthArgument) GetWorkspaceHost(ctx context.Context) string {
	return a.host
}

// GetCacheKey returns a unique key for caching the OAuth token for the workspace.
func (a BasicWorkspaceOAuthArgument) GetCacheKey(ctx context.Context) string {
	a.host = strings.TrimSuffix(a.host, "/")
	if !strings.HasPrefix(a.host, "http") {
		a.host = fmt.Sprintf("https://%s", a.host)
	}
	return a.host
}

var _ WorkspaceOAuthArgument = BasicWorkspaceOAuthArgument{}

// AccountOAuthArgument is an interface that provides the necessary information
// to authenticate using OAuth to a specific account.
type AccountOAuthArgument interface {
	OAuthArgument

	// GetAccountHost returns the host of the account to authenticate to.
	GetAccountHost(ctx context.Context) string

	// GetAccountId returns the account ID of the account to authenticate to.
	GetAccountId(ctx context.Context) string
}

// BasicAccountOAuthArgument is a basic implementation of the AccountOAuthArgument
// interface that links each account with exactly one OAuth token.
type BasicAccountOAuthArgument struct {
	accountHost string
	accountID   string
}

var _ AccountOAuthArgument = BasicAccountOAuthArgument{}

// NewBasicAccountOAuthArgument creates a new BasicAccountOAuthArgument.
func NewBasicAccountOAuthArgument(accountsHost, accountID string) (BasicAccountOAuthArgument, error) {
	if !strings.HasPrefix(accountsHost, "https://") {
		return BasicAccountOAuthArgument{}, fmt.Errorf("accountsHost must start with 'https://': %s", accountsHost)
	}
	if strings.HasSuffix(accountsHost, "/") {
		return BasicAccountOAuthArgument{}, fmt.Errorf("accountsHost must not have a trailing slash: %s", accountsHost)
	}
	return BasicAccountOAuthArgument{accountHost: accountsHost, accountID: accountID}, nil
}

// GetAccountHost returns the host of the account to authenticate to.
func (a BasicAccountOAuthArgument) GetAccountHost(ctx context.Context) string {
	return a.accountHost
}

// GetAccountId returns the account ID of the account to authenticate to.
func (a BasicAccountOAuthArgument) GetAccountId(ctx context.Context) string {
	return a.accountID
}

// GetCacheKey returns a unique key for caching the OAuth token for the account.
func (a BasicAccountOAuthArgument) GetCacheKey(ctx context.Context) string {
	return fmt.Sprintf("%s/oidc/accounts/%s", a.accountHost, a.accountID)
}
