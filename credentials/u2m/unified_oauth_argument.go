package u2m

import (
	"fmt"
)

// UnifiedOAuthArgument is an interface that provides the necessary information
// to authenticate using OAuth to a host that supports both account and workspace APIs.
type UnifiedOAuthArgument interface {
	OAuthArgument

	// GetHost returns the host to authenticate to.
	GetHost() string

	// GetAccountId returns the account ID of the account to authenticate to.
	GetAccountId() string
}

// BasicUnifiedOAuthArgument is a basic implementation of the UnifiedOAuthArgument
// interface that links each account with exactly one OAuth token.
type BasicUnifiedOAuthArgument struct {
	host      string
	accountID string
}

var _ UnifiedOAuthArgument = BasicUnifiedOAuthArgument{}

// NewBasicUnifiedOAuthArgument creates a new BasicUnifiedOAuthArgument.
func NewBasicUnifiedOAuthArgument(accountsHost, accountID string) (BasicUnifiedOAuthArgument, error) {
	if err := validateHost(accountsHost); err != nil {
		return BasicUnifiedOAuthArgument{}, err
	}
	return BasicUnifiedOAuthArgument{host: accountsHost, accountID: accountID}, nil
}

// GetAccountHost returns the host of the account to authenticate to.
func (a BasicUnifiedOAuthArgument) GetHost() string {
	return a.host
}

// GetAccountId returns the account ID of the account to authenticate to.
func (a BasicUnifiedOAuthArgument) GetAccountId() string {
	return a.accountID
}

// GetCacheKey returns a unique key for caching the OAuth token for the account.
// The key is in the format "<hostName>/oidc/accounts/<accountID>".
func (a BasicUnifiedOAuthArgument) GetCacheKey() string {
	return fmt.Sprintf("%s/oidc/accounts/%s", a.host, a.accountID)
}
