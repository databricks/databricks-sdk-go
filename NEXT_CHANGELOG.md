# NEXT CHANGELOG

## Release v0.129.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

 * Alias `apierr.ErrResourceAlreadyExists` to `apierr.ErrAlreadyExists` so `errors.Is` matches both the gRPC-canonical `ALREADY_EXISTS` and the Databricks-specific `RESOURCE_ALREADY_EXISTS` error codes, regardless of which one the backend returns.

### Documentation

### Internal Changes

 * Expanded AI agent detection: added Goose, Amp, Augment, Copilot (VS Code), Kiro, Windsurf. Honors the `AGENT=<name>` standard and falls back to `unknown` for unrecognized values. When multiple agent env vars are present (e.g. a Cursor CLI subagent invoked from Claude Code), the user-agent reports `agent/multiple`.

### API Changes
