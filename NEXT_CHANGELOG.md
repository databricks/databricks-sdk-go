# NEXT CHANGELOG

## Release v0.129.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

 * Fix CLI token source `--profile` fallback: `--profile` is a global Cobra flag that old CLIs accept silently instead of reporting "unknown flag", making the previous error-based detection dead code. Now uses `databricks version` to detect CLI capabilities at init time ([#1605](https://github.com/databricks/databricks-sdk-go/pull/1605)).

### Documentation

### Internal Changes

### API Changes
