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

const (
	// agentEnvVar is the agents.md standard env var.
	agentEnvVar = "AGENT"

	// aiAgentEnvVar is the Vercel @vercel/detect-agent convention. It
	// serves the same purpose as agentEnvVar; agentEnvFallback consults it
	// only when agentEnvVar is unset or empty.
	aiAgentEnvVar = "AI_AGENT"

	// maxAgentFallbackLen caps the length of a value surfaced via the
	// AGENT / AI_AGENT fallbacks. Explicit-matcher products are already
	// short, known strings; only the fallback path can carry arbitrary
	// user-supplied lengths into the User-Agent.
	maxAgentFallbackLen = 64
)

// listKnownAgents returns the canonical list of AI coding agents.
// Keep this list, and the AGENT / AI_AGENT fallback handling in
// agentEnvFallback, in sync with databricks-sdk-py and databricks-sdk-java.
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
// AGENT and AI_AGENT env vars, so that an explicit signal (e.g.
// CLAUDECODE=1) always wins over a conflicting AGENT=<name> value.
//
// The function counts how many distinct agents matched via explicit env vars:
//   - Exactly one agent matched: return its product name.
//   - More than one agent matched: return "multiple". Agent env vars can be
//     stacked when one agent invokes another as a subagent (e.g. Claude Code
//     spawning a Cursor CLI subprocess), so the child process inherits env
//     vars from multiple layers.
//   - Zero agents matched: see agentEnvFallback.
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
		return agentEnvFallback()
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

// agentEnvFallback consults the generic AGENT and AI_AGENT env vars and
// returns a sanitized name suitable for the User-Agent. AGENT (agents.md) is
// checked first; AI_AGENT (Vercel @vercel/detect-agent) is checked only when
// AGENT is unset or empty. Empty is treated as unset for both.
//
// The value is passed through, not matched against the known-agents list:
// arbitrary tool names (including versioned variants like
// "claude-code_2-1-141") are valuable telemetry that bucketing belongs
// downstream, not in the SDK. The value is sanitized so it satisfies the
// User-Agent allowlist and capped at maxAgentFallbackLen to keep the header
// bounded.
func agentEnvFallback() string {
	v := os.Getenv(agentEnvVar)
	if v == "" {
		v = os.Getenv(aiAgentEnvVar)
	}
	if v == "" {
		return ""
	}
	v = Sanitize(v)
	if len(v) > maxAgentFallbackLen {
		v = v[:maxAgentFallbackLen]
	}
	return v
}

var (
	agentOnce sync.Once
	agentName string
)

// AgentProvider returns the detected AI agent name, cached for the process lifetime.
// Returns one of:
//   - the known product name when exactly one agent is detected via explicit
//     env matchers;
//   - "multiple" when multiple explicit matchers fire for different agents
//     (typically nested agents, e.g. Cursor CLI running as a Claude Code
//     subagent);
//   - a sanitized, length-capped value from AGENT or AI_AGENT when no
//     explicit matcher fired (see agentEnvFallback);
//   - "" when no agent is detected.
func AgentProvider() string {
	agentOnce.Do(func() {
		agentName = lookupAgentProvider()
	})
	return agentName
}
