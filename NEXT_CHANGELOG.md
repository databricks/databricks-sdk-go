# NEXT CHANGELOG

## Release v0.148.0

### Breaking Changes

### New Features and Improvements

* Added a `meta-harness` user-agent dimension that reports the omnigent meta-harness (detected via the `OMNIGENT` environment variable) independently of agent detection.

### Bug Fixes
* Fix potential panic in `shouldRetry` function by using comma-ok idiom for type assertion ([#1438](https://github.com/databricks/databricks-sdk-go/pull/1438)).

### Documentation

### Internal Changes
* Remove deprecated `rand.Seed` call in retry backoff logic ([#1438](https://github.com/databricks/databricks-sdk-go/pull/1438)).

### API Changes
* Add `Xlarge` enum value for [apps.ComputeSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ComputeSize).