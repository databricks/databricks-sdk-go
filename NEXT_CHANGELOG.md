# NEXT CHANGELOG

## Release v0.61.0

### New Features and Improvements
* Introduce support for Databricks Workload Identity Federation in GitHub workflows ([1177](https://github.com/databricks/databricks-sdk-go/pull/1177)).
  See README.md for instructions.
  [Breaking] Users running their worklows in GitHub Actions, which use Cloud native authentication and also have a `DATABRICKS_CLIENT_ID` and `DATABRICKS_HOST` 
  environment variables set may see their authentication start failing due to the order in which the SDK tries different authentication methods.
  In such case, the `DATABRICKS_AUTH_TYPE` environment variable must be set to match the previously used authentication method.
  

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
