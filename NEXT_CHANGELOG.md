# NEXT CHANGELOG

## Release v0.80.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [w.Policies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PoliciesAPI) workspace-level service and [w.TemporaryPathCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryPathCredentialsAPI) workspace-level service.
* Added `Create` method for [w.Tables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI) workspace-level service.
* Added `ListDatabaseCatalogs`, `ListSyncedDatabaseTables`, `UpdateDatabaseCatalog` and `UpdateSyncedDatabaseTable` methods for [w.Database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseAPI) workspace-level service.
* Added `FirstOnDemand` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
* Added `UsagePolicyId` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
* Added `UsagePolicyId` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
* Added `UsagePolicyId` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
* Added `ClientRequestId` and `UsageContext` fields for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
* Added `ChannelId`, `ChannelIdSet`, `OauthToken` and `OauthTokenSet` fields for [settings.SlackConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SlackConfig).
* Added `Snapshot` enum value for [ml.PublishSpecPublishMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PublishSpecPublishMode).
* [Breaking] Changed `PublishMode` field for [ml.PublishSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PublishSpec) to be required.