package u2m

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultOAuthScopes_NormalizationAndDefault(t *testing.T) {
	s := newOAuthScopes()
	got := s.GetScopes()
	want := []string{"all-apis", "offline_access"}
	require.Equal(t, want, got, "GetScopes() should return normalized default scopes")
	require.True(t, s.isDefault(), "IsDefault() should be true for default scopes")
}

func TestNewOAuthScopes_Normalization(t *testing.T) {
	// duplicates, mixed case, empty, and whitespace
	s := newOAuthScopes("All-APIs", "offline_access", "", "  x:y ", "x:y")
	got := s.GetScopes()
	// normalized, sorted
	want := []string{"all-apis", "offline_access", "x:y"}
	require.Equal(t, want, got, "GetScopes() should return normalized, sorted scopes")
	require.False(t, s.isDefault(), "IsDefault() should be false for non-default scopes")
}

func TestComputeScopedCacheKey_DefaultKeepsBase(t *testing.T) {
	base := "https://abc"
	key := computeScopedCacheKey(base, newOAuthScopes())
	require.Equal(t, base, key)
}

func TestComputeScopedCacheKey_NonDefaultIncludesReadableScopes(t *testing.T) {
	base := "https://abc"
	key := computeScopedCacheKey(base, newOAuthScopes("foo", "offline_access"))
	require.True(t, strings.HasPrefix(key, base+"#scopes="))
	require.Contains(t, key, "foo")
	// ensure normalization kept offline_access
	require.Contains(t, key, "offline_access")
	// stability check: same scopes order-insensitive yields same key
	key2 := computeScopedCacheKey(base, newOAuthScopes("OFFLINE_ACCESS", "foo"))
	require.Equal(t, key, key2)
}
