# NEXT CHANGELOG

## Release v0.112.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Pass `--profile` to CLI token source when profile is set, and add read-fallback to migrate legacy host-keyed tokens to profile keys ([#1497](https://github.com/databricks/databricks-sdk-go/pull/1497)).

### Documentation

### Internal Changes

### API Changes
* Add `Parameters` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `Parameters` field for [pipelines.UpdateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfo).
* [Breaking] Change `GetDownloadFullQueryResult` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service with new required argument order.
* [Breaking] Change `Name` field for [apps.Space](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#Space) to be required.
* Change `Name` field for [apps.Space](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#Space) to be required.
* [Breaking] Change `Id` and `UserId` fields for [dashboards.GenieConversation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversation) to no longer be required.
* [Breaking] Change `CreatedTimestamp` and `Title` fields for [dashboards.GenieConversationSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversationSummary) to no longer be required.
* [Breaking] Change `DownloadIdSignature` field for [dashboards.GenieGetDownloadFullQueryResultRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGetDownloadFullQueryResultRequest) to be required.
* [Breaking] Change `Id` field for [dashboards.GenieMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieMessage) to no longer be required.