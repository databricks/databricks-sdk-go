package useragent

import (
	"os"
	"sync"
)

// envMatcher matches an environment variable. If value is empty, matching is
// presence-only (any value, including empty string, counts as a match). If
// value is non-empty, the env var must be set to exactly that value.
type envMatcher struct {
	envVar string
	value  string
}

// knownAgent describes a single AI coding agent and the environment matchers
// that identify it. The agent is detected if ANY matcher in matchAny fires.
type knownAgent struct {
	product  string
	matchAny []envMatcher
}

// agentEnvVar is the agents.md standard env var. When set to a value we
// don't specifically recognize, detection falls back to "unknown".
const agentEnvVar = "AGENT"

// listKnownAgents returns the canonical list of AI coding agents.
// Keep this list in sync with databricks-sdk-py and databricks-sdk-java.
// Agents are listed alphabetically by product name.
func listKnownAgents() []knownAgent {
	return []knownAgent{
		{
			product: "amp",
			matchAny: []envMatcher{
				{envVar: "AMP_CURRENT_THREAD_ID"},
				{envVar: agentEnvVar, value: "amp"},
			},
		},
		{
			product: "antigravity",
			matchAny: []envMatcher{
				{envVar: "ANTIGRAVITY_AGENT"}, // Closed source (Google)
			},
		},
		{
			product: "augment",
			matchAny: []envMatcher{
				{envVar: "AUGMENT_AGENT"},
			},
		},
		{
			product: "claude-code",
			matchAny: []envMatcher{
				{envVar: "CLAUDECODE"}, // https://github.com/anthropics/claude-code
			},
		},
		{
			product: "cline",
			matchAny: []envMatcher{
				{envVar: "CLINE_ACTIVE"}, // https://github.com/cline/cline (v3.24.0+)
			},
		},
		{
			product: "codex",
			matchAny: []envMatcher{
				{envVar: "CODEX_CI"}, // https://github.com/openai/codex
			},
		},
		{
			product: "copilot-cli",
			matchAny: []envMatcher{
				{envVar: "COPILOT_CLI"}, // https://github.com/features/copilot
			},
		},
		{
			product: "copilot-vscode",
			matchAny: []envMatcher{
				{envVar: "COPILOT_MODEL"}, // VS Code Copilot
			},
		},
		{
			product: "cursor",
			matchAny: []envMatcher{
				{envVar: "CURSOR_AGENT"}, // Closed source
			},
		},
		{
			product: "gemini-cli",
			matchAny: []envMatcher{
				{envVar: "GEMINI_CLI"}, // https://google-gemini.github.io/gemini-cli
			},
		},
		{
			product: "goose",
			matchAny: []envMatcher{
				{envVar: "GOOSE_TERMINAL"},
				{envVar: agentEnvVar, value: "goose"},
			},
		},
		{
			product: "kiro",
			matchAny: []envMatcher{
				{envVar: "KIRO"},
			},
		},
		{
			product: "opencode",
			matchAny: []envMatcher{
				{envVar: "OPENCODE"}, // https://github.com/opencode-ai/opencode
			},
		},
		{
			product: "openclaw",
			matchAny: []envMatcher{
				{envVar: "OPENCLAW_SHELL"}, // https://github.com/anthropics/openclaw
			},
		},
		{
			product: "windsurf",
			matchAny: []envMatcher{
				{envVar: "WINDSURF_AGENT"},
			},
		},
	}
}

// matcherFires returns true if the matcher's env var is set (for presence
// checks) or set to the exact expected value (for value checks).
func matcherFires(m envMatcher) bool {
	v, ok := os.LookupEnv(m.envVar)
	if !ok {
		return false
	}
	if m.value == "" {
		return true
	}
	return v == m.value
}

// lookupAgentProvider checks environment variables for known AI agents.
//
// For each agent, it fires if ANY of its matchers fires. The function counts
// how many distinct agents matched:
//   - Exactly one agent matched: return its product name.
//   - More than one agent matched: return "" (ambiguity).
//   - Zero agents matched: if the agents.md standard AGENT env var is set to
//     any non-empty value, return "unknown". Otherwise return "".
//
// Unlike CI/CD detection (which returns the first match), agent detection
// uses an ambiguity guard because agent env vars can be stacked (e.g., running
// Cline inside Cursor).
func lookupAgentProvider() string {
	var detected string
	count := 0
	for _, a := range listKnownAgents() {
		fired := false
		for _, m := range a.matchAny {
			if matcherFires(m) {
				fired = true
				break
			}
		}
		if fired {
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
	if count == 0 {
		if v, ok := os.LookupEnv(agentEnvVar); ok && v != "" {
			return "unknown"
		}
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
