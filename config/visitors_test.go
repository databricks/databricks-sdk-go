package config

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/stretchr/testify/assert"
)

func TestAuthInUserAgentVisitorDefault(t *testing.T) {
	c := &Config{
		AuthType: "PAT",
	}
	request := &http.Request{}
	err := authInUserAgentVisitor(c)(request)
	assert.NoError(t, err)
	// tests may be developed and run on different versions of different things
	uac := strings.Split(useragent.FromContext(request.Context()), " ")
	assert.Contains(t, uac, "auth/PAT")
}

func TestAuthInUserAgentVisitorCustom(t *testing.T) {
	c := &Config{
		AuthType: "PAT",
	}
	request := &http.Request{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, AuthType{}, "oath")
	request = request.WithContext(ctx)
	err := authInUserAgentVisitor(c)(request)
	assert.NoError(t, err)
	// tests may be developed and run on different versions of different things
	uac := strings.Split(useragent.FromContext(request.Context()), " ")
	assert.NotContains(t, uac, "auth/PAT")
	assert.Contains(t, uac, "auth/oath")
}

func TestAgentInUserAgentVisitorDetected(t *testing.T) {
	env.CleanupEnvironment(t)
	useragent.ClearCache()
	t.Setenv("CLAUDECODE", "1")

	request := &http.Request{}
	provider := useragent.AgentProvider()
	assert.Equal(t, "claude-code", provider)

	ctx := useragent.InContext(request.Context(), useragent.AgentKey, provider)
	*request = *request.WithContext(ctx)

	ua := useragent.FromContext(request.Context())
	assert.Contains(t, ua, "agent/claude-code")
	assert.Equal(t, 1, strings.Count(ua, "agent/"))
}

func TestAgentInUserAgentVisitorNoAgent(t *testing.T) {
	env.CleanupEnvironment(t)
	useragent.ClearCache()

	request := &http.Request{}
	provider := useragent.AgentProvider()
	assert.Equal(t, "", provider)

	ua := useragent.FromContext(request.Context())
	assert.NotContains(t, ua, "agent/")
}
