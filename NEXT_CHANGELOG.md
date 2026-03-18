# NEXT CHANGELOG

## Release v0.125.0

### Breaking Changes

### New Features and Improvements

* Add automatic detection of AI coding agents (Claude Code, Cursor, Cline, Codex, Copilot CLI, Gemini CLI, OpenCode, Antigravity, Openclaw) in the user-agent string. The SDK now appends `agent/<name>` to HTTP request headers when running inside a known AI agent environment.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `DisableGovTagCreation` field for [settings.RestrictWorkspaceAdminsMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsMessage).
* Add `DisableGovTagCreation` field for [settingsv2.RestrictWorkspaceAdminsMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#RestrictWorkspaceAdminsMessage).