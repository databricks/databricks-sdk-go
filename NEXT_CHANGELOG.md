# NEXT CHANGELOG

## Release v0.133.0

### Breaking Changes

### New Features and Improvements

* Honor the Vercel `AI_AGENT=<name>` env var as a secondary fallback for
  AI agent detection in the User-Agent header (after the agents.md
  `AGENT=<name>` standard). Unrecognized fallback values now pass through
  the User-Agent sanitized and length-capped at 64 chars instead of being
  coerced to `agent/unknown`, so versioned variants such as
  `claude-code_2-1-141_agent` surface as-is.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
