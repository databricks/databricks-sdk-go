# NEXT CHANGELOG

## Release v0.181.0

### Breaking Changes

### New Features and Improvements

- Add `Config.LiteswapTarget` (env `DATABRICKS_LITESWAP_TARGET`, profile field `liteswap_target`) to route API requests to a liteswap test instance of a service in shared pre-production. Dev/test only; empty (the default) is a no-op.

### Bug Fixes

### Documentation

### Internal Changes

- Reduce integration test cluster usage by replacing waiter coverage with HTTP fixtures and removing redundant live tests.

### API Changes
