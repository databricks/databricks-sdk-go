package databricks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrentWorkspaceIDShortCircuitsWhenConfigHasWorkspaceID(t *testing.T) {
	// When Config.WorkspaceID is already populated (from profile, ?o= query
	// param, env var, or host metadata), CurrentWorkspaceID returns it without
	// hitting the API. This avoids a round-trip and sidesteps SPOG's routing
	// requirement that requests to /api/2.0/preview/scim/v2/Me carry an
	// X-Databricks-Org-Id header.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.0/preview/scim/v2/Me" {
			t.Errorf("Me() should not be called when Config.WorkspaceID is set")
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	w, err := NewWorkspaceClient(&Config{
		Host:        server.URL,
		Token:       "token",
		WorkspaceID: "7474644166319138",
	})
	require.NoError(t, err)

	got, err := w.CurrentWorkspaceID(t.Context())
	require.NoError(t, err)
	assert.Equal(t, int64(7474644166319138), got)
}

func TestCurrentWorkspaceIDFallsBackToAPIWhenConfigMissingWorkspaceID(t *testing.T) {
	var meCalls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.0/preview/scim/v2/Me" {
			meCalls++
			w.Header().Set("X-Databricks-Org-Id", "7474644166319138")
			w.Write([]byte(`{}`))
			return
		}
		// Other bootstrap calls (e.g. /.well-known/databricks-config) are not
		// relevant to this test; respond with a benign 404.
		http.NotFound(w, r)
	}))
	defer server.Close()

	w, err := NewWorkspaceClient(&Config{
		Host:  server.URL,
		Token: "token",
	})
	require.NoError(t, err)

	got, err := w.CurrentWorkspaceID(t.Context())
	require.NoError(t, err)
	assert.Equal(t, int64(7474644166319138), got)
	assert.Equal(t, 1, meCalls)
}
