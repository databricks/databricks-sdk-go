package useragent

import (
	"os"
	"sync"
)

// knownAgent describes a single AI coding agent and the environment variable
// that identifies it. The agent is detected if envVar is set (any value,
// including the empty string, counts as a match).
type knownAgent struct {
	envVar  string
	product string
}

// agentEnvVar is the agents.md standard env var. When set to a value we
// don't specifically recognize, detection falls back to "unknown".
const agentEnvVar = "AGENT"

// listKnownAgents returns the canonical list of AI coding agents.
// Keep this list in sync with databricks-sdk-py and databricks-sdk-java.
// Agents are listed alphabetically by product name.
func listKnownAgents() []knownAgent {
	return []knownAgent{
		{envVar: "AMP_CURRENT_THREAD_ID", product: "amp"},     // https://ampcode.com/ (also sets AGENT=amp, handled by the central fallback in lookupAgentProvider)
		{envVar: "ANTIGRAVITY_AGENT", product: "antigravity"}, // Closed source (Google)
		{envVar: "AUGMENT_AGENT", product: "augment"},         // https://www.augmentcode.com/
		{envVar: "CLAUDECODE", product: "claude-code"},        // https://github.com/anthropics/claude-code
		{envVar: "CLINE_ACTIVE", product: "cline"},            // https://github.com/cline/cline (v3.24.0+)
		{envVar: "CODEX_CI", product: "codex"},                // https://github.com/openai/codex
		{envVar: "COPILOT_CLI", product: "copilot-cli"},       // https://github.com/features/copilot
		{envVar: "COPILOT_MODEL", product: "copilot-vscode"},  // VS Code Copilot terminal, best-effort heuristic, not officially identified
		{envVar: "CURSOR_AGENT", product: "cursor"},           // Closed source
		{envVar: "GEMINI_CLI", product: "gemini-cli"},         // https://google-gemini.github.io/gemini-cli
		{envVar: "GOOSE_TERMINAL", product: "goose"},          // https://block.github.io/goose/ (also sets AGENT=goose, handled by the central fallback in lookupAgentProvider)
		{envVar: "KIRO", product: "kiro"},                     // https://kiro.dev/ (Amazon)
		{envVar: "OPENCLAW_SHELL", product: "openclaw"},       // https://github.com/anthropics/openclaw
		{envVar: "OPENCODE", product: "opencode"},             // https://github.com/opencode-ai/opencode
		{envVar: "WINDSURF_AGENT", product: "windsurf"},       // https://codeium.com/windsurf (Codeium)
	}
}

// lookupAgentProvider checks environment variables for known AI agents.
//
// Explicit product-specific env vars always take precedence over the generic
// agents.md AGENT env var. AGENT is consulted only as a fallback when no
// explicit matcher fires, so that an explicit signal (e.g. CLAUDECODE=1)
// always wins over a conflicting AGENT=<name> value.
//
// The function counts how many distinct agents matched via explicit env vars:
//   - Exactly one agent matched: return its product name.
//   - More than one agent matched: return "multiple". Agent env vars can be
//     stacked when one agent invokes another as a subagent (e.g. Claude Code
//     spawning a Cursor CLI subprocess), so the child process inherits env
//     vars from multiple layers.
//   - Zero agents matched: if the agents.md standard AGENT env var is set to
//     a non-empty value, return that value if it matches a known product name,
//     or "unknown" otherwise. If AGENT is not set, return "".
func lookupAgentProvider() string {
	agents := listKnownAgents()

	var matches []string
	for _, a := range agents {
		if _, ok := os.LookupEnv(a.envVar); ok {
			matches = append(matches, a.product)
		}
	}

	// Known BYOK false positive: Copilot CLI users often set COPILOT_MODEL
	// alongside COPILOT_CLI. That is a single copilot-cli signal, not a
	// stacked multi-agent setup, so drop the copilot-vscode match.
	matches = collapseCopilotBYOK(matches)

	switch len(matches) {
	case 1:
		return matches[0]
	case 0:
		return agentEnvFallback(agents)
	default:
		return "multiple"
	}
}

func collapseCopilotBYOK(matches []string) []string {
	hasCLI, hasVSCode := false, false
	for _, m := range matches {
		if m == "copilot-cli" {
			hasCLI = true
		}
		if m == "copilot-vscode" {
			hasVSCode = true
		}
	}
	if !hasCLI || !hasVSCode {
		return matches
	}
	filtered := make([]string, 0, len(matches)-1)
	for _, m := range matches {
		if m == "copilot-vscode" {
			continue
		}
		filtered = append(filtered, m)
	}
	return filtered
}

// agentEnvFallback honors the agents.md AGENT=<name> standard.
// Returns the value if it matches a known product name, "unknown" if AGENT
// is set to any other non-empty value, and "" if AGENT is unset or empty.
func agentEnvFallback(agents []knownAgent) string {
	v, ok := os.LookupEnv(agentEnvVar)
	if !ok || v == "" {
		return ""
	}
	for _, a := range agents {
		if a.product == v {
			return v
		}
	}
	return "unknown"
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
//   - "multiple" when multiple explicit matchers fire for different agents
//     (typically nested agents, e.g. Cursor CLI running as a Claude Code
//     subagent);
//   - "unknown" when no explicit matcher fired and AGENT is set to a value
//     that is not a known product name;
//   - "" when no agent is detected.
func AgentProvider() string {
	agentOnce.Do(func() {
		agentName = lookupAgentProvider()
	})
	return agentName
}
