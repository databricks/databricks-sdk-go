# NEXT CHANGELOG

## Release v0.181.0

### Breaking Changes

### New Features and Improvements

- Add `Config.Headers`, a hook called on every request made by the client to set custom HTTP headers, on top of normal authentication. Use the `StaticHeaders` helper to set a fixed set of headers.

### Bug Fixes

### Documentation

### Internal Changes

- Reduce integration test cluster usage by replacing waiter coverage with HTTP fixtures and removing redundant live tests.

### API Changes
