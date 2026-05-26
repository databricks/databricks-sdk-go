# NEXT CHANGELOG

## Release v0.138.0

### Breaking Changes

### New Features and Improvements

* Detect VS Code agent-initiated terminal commands via the `VSCODE_AGENT` environment variable (VS Code 1.121+). The User-Agent header now reports `agent/vscode-agent` when set. The previous `COPILOT_MODEL` heuristic (which reported `agent/copilot-vscode`) has been removed; it produced false positives for Copilot CLI BYOK users and never reliably identified VS Code.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
