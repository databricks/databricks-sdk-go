# NEXT CHANGELOG

## Release v0.86.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `UpdateNotifications` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
* Add `Parameters` field for [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Add `DatabaseInstanceName` field for [database.CreateDatabaseInstanceRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#CreateDatabaseInstanceRoleRequest).
* Add `CustomTags`, `EffectiveCustomTags`, `EffectiveUsagePolicyId` and `UsagePolicyId` fields for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Add `EffectiveAttributes` and `InstanceName` fields for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole).
* Add `StreamNative` enum value for [catalog.SystemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemType).
* Add `ExceededMaxTokenLengthException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Change `Name` field for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole) to be required.
* [Breaking] Change `Name` field for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole) to be required.