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
				// https://ampcode.com/ (also sets AGENT=amp, handled by the
				// central fallback in lookupAgentProvider).
				{envVar: "AMP_CURRENT_THREAD_ID"},
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
				{envVar: "AUGMENT_AGENT"}, // https://www.augmentcode.com/
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
				// VS Code Copilot terminal, best-effort heuristic, not officially identified
				{envVar: "COPILOT_MODEL"},
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
				// https://block.github.io/goose/ (also sets AGENT=goose, handled
				// by the central fallback in lookupAgentProvider).
				{envVar: "GOOSE_TERMINAL"},
			},
		},
		{
			product: "kiro",
			matchAny: []envMatcher{
				{envVar: "KIRO"}, // https://kiro.dev/ (Amazon)
			},
		},
		{
			product: "openclaw",
			matchAny: []envMatcher{
				{envVar: "OPENCLAW_SHELL"}, // https://github.com/anthropics/openclaw
			},
		},
		{
			product: "opencode",
			matchAny: []envMatcher{
				{envVar: "OPENCODE"}, // https://github.com/opencode-ai/opencode
			},
		},
		{
			product: "windsurf",
			matchAny: []envMatcher{
				{envVar: "WINDSURF_AGENT"}, // https://codeium.com/windsurf (Codeium)
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
// Explicit product-specific env vars always take precedence over the generic
// agents.md AGENT env var. AGENT is consulted only as a fallback when no
// explicit matcher fires, so that an explicit signal (e.g. CLAUDECODE=1)
// always wins over a conflicting AGENT=<name> value.
//
// For each agent, it fires if ANY of its matchers fires. The function counts
// how many distinct agents matched via explicit matchers:
//   - Exactly one agent matched: return its product name.
//   - More than one agent matched: return "" (ambiguity).
//   - Zero agents matched: if the agents.md standard AGENT env var is set to
//     a non-empty value, return that value if it matches a known product name,
//     or "unknown" otherwise. If AGENT is not set, return "".
//
// Unlike CI/CD detection (which returns the first match), agent detection
// uses an ambiguity guard because agent env vars can be stacked (e.g., running
// Cline inside Cursor).
func lookupAgentProvider() string {
	agents := listKnownAgents()
	var detected string
	count := 0
	for _, a := range agents {
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
			// Honor the agents.md standard: if AGENT is set to a known product
			// name that wasn't caught by any explicit matcher, report it directly.
			for _, a := range agents {
				if a.product == v {
					return v
				}
			}
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
// Returns one of:
//   - the known product name when exactly one agent is detected via explicit
//     env matchers, or when AGENT is set to a known product name and no
//     explicit matcher fired;
//   - "unknown" when no explicit matcher fired and AGENT is set to a value
//     that is not a known product name;
//   - "" when no agent is detected, or when multiple explicit matchers fire
//     for different agents (ambiguity).
func AgentProvider() string {
	agentOnce.Do(func() {
		agentName = lookupAgentProvider()
	})
	return agentName
}
