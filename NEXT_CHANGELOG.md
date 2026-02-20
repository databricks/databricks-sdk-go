# NEXT CHANGELOG

## Release v0.111.0

### Breaking Changes

### New Features and Improvements
* Add `DefaultCredentialStrategyProvider` package-level variable to allow 
  overriding the default authentication strategy for all `Config` instances.

### Bug Fixes

### Documentation

### Internal Changes

* Make clusters creation in AWS not depend on cloud parsed from host. Changed default AWS availability for auto-created utility clusters from SPOT to SPOT_WITH_FALLBACK (API default). 

### API Changes
