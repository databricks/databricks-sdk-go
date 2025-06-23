package u2m

import (
	"fmt"
)

// AccountOAuthArgument is an interface that provides the necessary information
// to authenticate using OAuth to a specific account.
type AccountOAuthArgument interface {
	OAuthArgument

	// GetAccountHost returns the host of the account to authenticate to.
	GetAccountHost() string

	// GetAccountId returns the account ID of the account to authenticate to.
	GetAccountId() string
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
	if err := validateHost(accountsHost); err != nil {
		return BasicAccountOAuthArgument{}, err
	}
	return BasicAccountOAuthArgument{accountHost: accountsHost, accountID: accountID}, nil
}

// GetAccountHost returns the host of the account to authenticate to.
func (a BasicAccountOAuthArgument) GetAccountHost() string {
	return a.accountHost
}

// GetAccountId returns the account ID of the account to authenticate to.
func (a BasicAccountOAuthArgument) GetAccountId() string {
	return a.accountID
}

// GetCacheKey returns a unique key for caching the OAuth token for the account.
// The key is in the format "<accountHost>/oidc/accounts/<accountID>".
func (a BasicAccountOAuthArgument) GetCacheKey() string {
	return fmt.Sprintf("%s/oidc/accounts/%s", a.accountHost, a.accountID)
}
