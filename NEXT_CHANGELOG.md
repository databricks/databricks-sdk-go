# NEXT CHANGELOG

## Release v0.59.0

### New Features and Improvements

### Bug Fixes

* Fix unlikely issue due to conflicting error details in `APIError`.

### Documentation

### Internal Changes

* Add `poll.SimpleError` to mock waiter objects returning errors  ([#1155](https://github.com/databricks/databricks-sdk-go/pull/1155))
* Refactor `APIError` to expose different types of error details ([#1153](https://github.com/databricks/databricks-sdk-go/pull/1153)). 
* Update Jobs ListJobs API to support paginated responses ([#1150](https://github.com/databricks/databricks-sdk-go/pull/1150))
* Introduce automated tagging ([#1148](https://github.com/databricks/databricks-sdk-go/pull/1148)).
* Update Jobs GetJob API to support paginated responses  ([#1133](https://github.com/databricks/databricks-sdk-go/pull/1133)).
* Update Jobs GetRun API to support paginated responses  ([#1132](https://github.com/databricks/databricks-sdk-go/pull/1132)).

### API Changes
