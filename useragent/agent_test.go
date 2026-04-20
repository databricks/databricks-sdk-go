package useragent

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
)

func TestLookupAgentProvider(t *testing.T) {
	tests := []struct {
		name   string
		envs   map[string]string
		expect string
	}{
		{
			name:   "no agent",
			envs:   nil,
			expect: "",
		},
		{
			name:   "claude code",
			envs:   map[string]string{"CLAUDECODE": "1"},
			expect: "claude-code",
		},
		{
			name:   "antigravity",
			envs:   map[string]string{"ANTIGRAVITY_AGENT": "1"},
			expect: "antigravity",
		},
		{
			name:   "cline",
			envs:   map[string]string{"CLINE_ACTIVE": "1"},
			expect: "cline",
		},
		{
			name:   "codex",
			envs:   map[string]string{"CODEX_CI": "1"},
			expect: "codex",
		},
		{
			name:   "copilot cli",
			envs:   map[string]string{"COPILOT_CLI": "1"},
			expect: "copilot-cli",
		},
		{
			name:   "cursor",
			envs:   map[string]string{"CURSOR_AGENT": "1"},
			expect: "cursor",
		},
		{
			name:   "gemini cli",
			envs:   map[string]string{"GEMINI_CLI": "1"},
			expect: "gemini-cli",
		},
		{
			name:   "opencode",
			envs:   map[string]string{"OPENCODE": "1"},
			expect: "opencode",
		},
		{
			name:   "openclaw",
			envs:   map[string]string{"OPENCLAW_SHELL": "exec"},
			expect: "openclaw",
		},
		{
			name:   "multiple agents",
			envs:   map[string]string{"CLAUDECODE": "1", "CURSOR_AGENT": "1"},
			expect: "",
		},
		{
			name:   "empty value still counts as set",
			envs:   map[string]string{"CLAUDECODE": ""},
			expect: "claude-code",
		},
		// New agent detections.
		{
			name:   "goose via GOOSE_TERMINAL",
			envs:   map[string]string{"GOOSE_TERMINAL": "1"},
			expect: "goose",
		},
		{
			name:   "goose via AGENT",
			envs:   map[string]string{"AGENT": "goose"},
			expect: "goose",
		},
		{
			name:   "goose via both GOOSE_TERMINAL and AGENT is not ambiguous",
			envs:   map[string]string{"GOOSE_TERMINAL": "1", "AGENT": "goose"},
			expect: "goose",
		},
		{
			name:   "amp via AMP_CURRENT_THREAD_ID",
			envs:   map[string]string{"AMP_CURRENT_THREAD_ID": "abc123"},
			expect: "amp",
		},
		{
			name:   "amp via AGENT",
			envs:   map[string]string{"AGENT": "amp"},
			expect: "amp",
		},
		{
			name:   "amp via both AMP_CURRENT_THREAD_ID and AGENT is not ambiguous",
			envs:   map[string]string{"AMP_CURRENT_THREAD_ID": "abc123", "AGENT": "amp"},
			expect: "amp",
		},
		{
			name:   "augment",
			envs:   map[string]string{"AUGMENT_AGENT": "1"},
			expect: "augment",
		},
		{
			name:   "copilot vscode",
			envs:   map[string]string{"COPILOT_MODEL": "gpt-4"},
			expect: "copilot-vscode",
		},
		{
			name:   "kiro",
			envs:   map[string]string{"KIRO": "1"},
			expect: "kiro",
		},
		{
			name:   "windsurf",
			envs:   map[string]string{"WINDSURF_AGENT": "1"},
			expect: "windsurf",
		},
		// AGENT fallback behavior.
		{
			name:   "AGENT with unknown value falls back to unknown",
			envs:   map[string]string{"AGENT": "someweirdthing"},
			expect: "unknown",
		},
		{
			name:   "AGENT empty string does not trigger fallback",
			envs:   map[string]string{"AGENT": ""},
			expect: "",
		},
		// Cross-agent ambiguity cases.
		{
			name:   "AGENT=goose and CLAUDECODE is ambiguous",
			envs:   map[string]string{"AGENT": "goose", "CLAUDECODE": "1"},
			expect: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env.CleanupEnvironment(t)
			ClearCache()
			for k, v := range tt.envs {
				t.Setenv(k, v)
			}
			got := lookupAgentProvider()
			if got != tt.expect {
				t.Errorf("lookupAgentProvider() = %q, want %q", got, tt.expect)
			}
		})
	}
}

func TestAgentProviderCached(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CURSOR_AGENT", "1")
	got := AgentProvider()
	if got != "cursor" {
		t.Errorf("AgentProvider() = %q, want %q", got, "cursor")
	}
	// Change env after caching. Cached result should persist.
	t.Setenv("CURSOR_AGENT", "")
	t.Setenv("CLAUDECODE", "1")
	got = AgentProvider()
	if got != "cursor" {
		t.Errorf("AgentProvider() after cache = %q, want %q", got, "cursor")
	}
}
