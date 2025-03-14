package config

import (
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type AuthType struct{}

func authInUserAgentVisitor(authType string) func(*http.Request) error {
	return func(r *http.Request) error {
		if t := r.Context().Value(AuthType{}); t != nil {
			authType = t.(string)
		}
		ctx := useragent.InContext(r.Context(), useragent.AuthKey, authType)
		*r = *r.WithContext(ctx) // replace request
		return nil
	}
}
