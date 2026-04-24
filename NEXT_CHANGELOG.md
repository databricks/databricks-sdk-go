# NEXT CHANGELOG

## Release v0.131.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

 * Alias `apierr.ErrResourceAlreadyExists` to `apierr.ErrAlreadyExists` so `errors.Is` matches both the gRPC-canonical `ALREADY_EXISTS` and the Databricks-specific `RESOURCE_ALREADY_EXISTS` error codes, regardless of which one the backend returns.

### Documentation

### Internal Changes

### API Changes
