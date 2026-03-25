package config

import (
	"net/http"

	"github.com/databricks/databricks-sdk-go/useragent"
)

type AuthType struct{}

// workspaceIDHeaderVisitor sets the X-Databricks-Org-Id header on every
// request when a workspace ID is configured. SPOG (Single Pane of Glass)
// hosts serve multiple workspaces behind one URL, so the server needs this
// header to route requests to the correct workspace.
func workspaceIDHeaderVisitor(c *Config) func(*http.Request) error {
	return func(r *http.Request) error {
		if c.WorkspaceID != "" {
			r.Header.Set("X-Databricks-Org-Id", c.WorkspaceID)
		}
		return nil
	}
}

func authInUserAgentVisitor(c *Config) func(*http.Request) error {
	return func(r *http.Request) error {
		authType := c.AuthType
		if t := r.Context().Value(AuthType{}); t != nil {
			authType = t.(string)
		}
		ctx := useragent.InContext(r.Context(), useragent.AuthKey, authType)
		*r = *r.WithContext(ctx) // replace request
		return nil
	}
}
