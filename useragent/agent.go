package useragent

import (
	"os"
	"sync"
)

// knownAgent maps an environment variable to an agent product name.
type knownAgent struct {
	envVar  string
	product string
}

// listKnownAgents returns the canonical list of AI coding agents.
// Keep this list in sync with databricks-sdk-py and databricks-sdk-java.
func listKnownAgents() []knownAgent {
	return []knownAgent{
		{"ANTIGRAVITY_AGENT", "antigravity"}, // Closed source (Google)
		{"CLAUDECODE", "claude-code"},        // https://github.com/anthropics/claude-code
		{"CLINE_ACTIVE", "cline"},            // https://github.com/cline/cline (v3.24.0+)
		{"CODEX_CI", "codex"},                // https://github.com/openai/codex
		{"CURSOR_AGENT", "cursor"},           // Closed source
		{"GEMINI_CLI", "gemini-cli"},         // https://google-gemini.github.io/gemini-cli
		{"OPENCODE", "opencode"},             // https://github.com/opencode-ai/opencode
	}
}

// lookupAgentProvider checks environment for known agent env vars.
// Returns the product name if exactly one is set (non-empty).
// Returns empty string if zero or multiple are set.
//
// Unlike CI/CD detection (which returns the first match), agent detection
// uses an ambiguity guard: multiple matches return empty. This is because
// CI/CD providers are mutually exclusive in practice, but agent env vars
// can be stacked (e.g., running Cline inside Cursor).
func lookupAgentProvider() string {
	var detected string
	count := 0
	for _, a := range listKnownAgents() {
		if v := os.Getenv(a.envVar); v != "" {
			detected = a.product
			count++
			if count > 1 {
				break
			}
		}
	}
	if count == 1 {
		return detected
	}
	return ""
}

var (
	agentOnce sync.Once
	agentName string
)

// AgentProvider returns the detected AI agent name, cached for the process lifetime.
// Returns empty string if no agent or multiple agents detected.
func AgentProvider() string {
	agentOnce.Do(func() {
		agentName = lookupAgentProvider()
	})
	return agentName
}
