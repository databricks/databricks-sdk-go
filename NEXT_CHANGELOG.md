# NEXT CHANGELOG

## Release v0.61.0

### New Features and Improvements
* Introduce support for Databricks Workload Identity Federation in GitHub workflows ([1177](https://github.com/databricks/databricks-sdk-go/pull/1177)).
  See README.md for instructions.
  [Breaking] Users running their worklows in GitHub Actions, which use Cloud native authentication and also have a `DATABRICKS_CLIENT_ID` and `DATABRICKS_HOST` 
  environment variables set may see their authentication start failing due to the order in which the SDK tries different authentication methods.
  In such case, the `DATABRICKS_AUTH_TYPE` environment variable must be set to match the previously used authentication method.
  

* Support user-to-machine authentication in the SDK ([#1108](https://github.com/databricks/databricks-sdk-go/pull/1108)).
- Instances of `ApiClient` now share the same connection pool by default ([PR #1190](https://github.com/databricks/databricks-sdk-go/pull/1190)).

### Bug Fixes

### Documentation

### Internal Changes

- Stop recommending users to report an issue when the SDK encounters an unknown
  error ([PR #1189](https://github.com/databricks/databricks-sdk-go/pull/1189)).

### API Changes
