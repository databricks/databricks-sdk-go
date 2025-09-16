package u2m

import (
	"slices"
	"sort"
	"strings"
)

// OAuthScopes encapsulates OAuth scopes configuration and normalization logic.
//
// It guarantees that the resulting scope list:
// - always includes "offline_access" (to receive refresh tokens)
// - is lower-cased, de-duplicated, and stable-sorted
//
// The default value equals the SDK's historical behavior: offline_access + all-apis.
type oauthScopes struct {
	values []string
}

var defaultScopeValues = []string{
	"offline_access", // Ensures OAuth token includes refresh token.
	"all-apis",       // Ensures OAuth token has access to all control-plane APIs.
}

// newOAuthScopes constructs oauthScopes from the provided scope values.
// If no scopes are provided, it returns the default scopes.
func newOAuthScopes(scopes ...string) oauthScopes {
	if len(scopes) == 0 {
		scopes = defaultScopeValues
	}
	vals := make([]string, len(scopes))
	copy(vals, scopes)
	return oauthScopes{values: vals}
}

// GetScopes returns the normalized, de-duplicated, stable-sorted list of scopes,
// guaranteeing that "offline_access" is included.
// Exported because it implements the OAuthArgument interface.
func (s oauthScopes) GetScopes() []string {
	return normalizeScopes(s.values)
}

// isDefault reports whether the scopes are equivalent to the default scopes
// (order-insensitive, case-insensitive, duplicates ignored).
func (s oauthScopes) isDefault() bool {
	left := normalizeScopes(s.values)
	right := normalizeScopes(defaultScopeValues)
	return slices.Equal(left, right)
}

// ComputeScopedCacheKey produces a backward-compatible cache key:
//   - For default scopes, it returns the baseKey unchanged
//   - For non-default scopes, it appends a stable hash suffix of the scopes
//     to avoid very long keys while ensuring uniqueness
//
// Result format for non-default scopes: "<baseKey>#scopes=<scope1,scope2,...>"
// where scopes are normalized and comma-separated for readability.
func computeScopedCacheKey(baseKey string, scopes oauthScopes) string {
	if scopes.isDefault() {
		return baseKey
	}
	joined := strings.Join(normalizeScopes(scopes.values), ",")
	return baseKey + "#scopes=" + joined
}

// normalizeScopes lowercases, ensures offline_access presence, removes duplicates,
// and returns a stable-sorted slice.
func normalizeScopes(in []string) []string {
	set := map[string]struct{}{}
	for _, s := range in {
		v := strings.TrimSpace(strings.ToLower(s))
		if v == "" {
			continue
		}
		set[v] = struct{}{}
	}
	// Always include offline_access to guarantee refresh tokens
	set["offline_access"] = struct{}{}

	out := make([]string, 0, len(set))
	for v := range set {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}
