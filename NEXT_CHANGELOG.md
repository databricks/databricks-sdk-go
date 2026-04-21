# NEXT CHANGELOG

## Release v0.129.0

### Breaking Changes

### New Features and Improvements

 * Add `u2m.WithDiscoveryHost` option to override the default `https://login.databricks.com` host used by the discovery login flow. Intended for testing and development against non-production environments.

### Bug Fixes

### Documentation

### Internal Changes

 * Expanded AI agent detection: added Goose, Amp, Augment, Copilot (VS Code), Kiro, Windsurf. Honors the `AGENT=<name>` standard and falls back to `unknown` for unrecognized values. When multiple agent env vars are present (e.g. a Cursor CLI subagent invoked from Claude Code), the user-agent reports `agent/multiple`.

### API Changes
