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
//
// This is used in two places:
//  1. dualWrite: writing tokens to both the profile key and the legacy host
//     key so that older Python/Java SDKs (which only know host keys) can
//     still find the token.
//  2. Token() fallback: when a profile key miss occurs, the host key is tried
//     as a read-only fallback for tokens written before profile-based keys
//     existed.
//
// This interface (and the host-key write path) can be removed once all SDKs
// sharing the token cache have migrated to profile-based keys.
type HostCacheKeyProvider interface {
	GetHostCacheKey() string
}
