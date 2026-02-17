# NEXT CHANGELOG

## Release v0.109.0

### Breaking Changes

### New Features and Improvements
* Use profile name as OAuth token cache key instead of host URL, so two profiles on the same host store separate tokens.

### Bug Fixes

### Documentation
* Implement dynamic auth token stale period based on initial token lifetime. Increased up to 20 mins for standard OAuth with proportionally shorter periods for short-lived tokens.

### Internal Changes

### API Changes
