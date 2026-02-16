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

// HostCacheKeyProvider is an interface for OAuthArgument implementations that
// can return a host-based cache key regardless of whether a profile is set.
// This is used by PersistentAuth for dual-write: writing tokens to both the
// profile-based key and the legacy host-based key during migration.
type HostCacheKeyProvider interface {
	GetHostCacheKey() string
}
