# NEXT CHANGELOG

## Release v0.69.0

### New Features and Improvements

- Add support to authenticate with Account-wide token federation from the 
  following auth methods: `env-oidc`, `file-oidc`, and `github-oidc`.  

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added `ExternalAccessEnabled` field to `catalog.UpdateMetastore` struct to support controlling non-DBR client access to metastore entities.
