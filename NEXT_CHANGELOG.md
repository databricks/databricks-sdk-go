# NEXT CHANGELOG

## Release v0.90.0

### New Features and Improvements
* Add support for unified hosts, i.e. hosts that support both workspace-level and account-level operations
* Deprecate Config.IsAccountClient, which will not work for unified hosts, and replace it with Config.HostType and Config.ConfigType methods.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `GoogleAds`, `TiktokAds`, `SalesforceMarketingCloud`, `Hubspot`, `WorkdayHcm`, `Guidewire` and `Zendesk` enum values for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).