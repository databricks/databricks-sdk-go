package useragent

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/stretchr/testify/assert"
)

func TestAgentProviderNoAgent(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	assert.Equal(t, "", lookupAgentProvider())
}

func TestAgentProviderClaudeCode(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CLAUDECODE", "1")
	assert.Equal(t, "claude-code", lookupAgentProvider())
}

func TestAgentProviderAntigravity(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("ANTIGRAVITY_AGENT", "1")
	assert.Equal(t, "antigravity", lookupAgentProvider())
}

func TestAgentProviderCline(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CLINE_ACTIVE", "1")
	assert.Equal(t, "cline", lookupAgentProvider())
}

func TestAgentProviderCodex(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CODEX_CI", "1")
	assert.Equal(t, "codex", lookupAgentProvider())
}

func TestAgentProviderCopilotCLI(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("COPILOT_CLI", "1")
	assert.Equal(t, "copilot-cli", lookupAgentProvider())
}

func TestAgentProviderCursor(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CURSOR_AGENT", "1")
	assert.Equal(t, "cursor", lookupAgentProvider())
}

func TestAgentProviderGeminiCLI(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("GEMINI_CLI", "1")
	assert.Equal(t, "gemini-cli", lookupAgentProvider())
}

func TestAgentProviderOpenCode(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("OPENCODE", "1")
	assert.Equal(t, "opencode", lookupAgentProvider())
}

func TestAgentProviderMultipleAgents(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("CURSOR_AGENT", "1")
	assert.Equal(t, "", lookupAgentProvider())
}

func TestAgentProviderEmptyValue(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CLAUDECODE", "")
	assert.Equal(t, "", lookupAgentProvider())
}

func TestAgentProviderCached(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CURSOR_AGENT", "1")
	assert.Equal(t, "cursor", AgentProvider())
	// Change env after caching. Cached result should persist.
	t.Setenv("CURSOR_AGENT", "")
	t.Setenv("CLAUDECODE", "1")
	assert.Equal(t, "cursor", AgentProvider())
}
