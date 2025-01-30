package u2m

// OAuthArgument is an interface that provides the necessary information to
// authenticate with PersistentAuth. Implementations of this interface must
// implement either the WorkspaceOAuthArgument or AccountOAuthArgument
// interface.
type OAuthArgument interface {
	// GetCacheKey returns a unique key for the OAuthArgument. This key is used
	// to store and retrieve the token from the token cache.
	GetCacheKey() string
}
