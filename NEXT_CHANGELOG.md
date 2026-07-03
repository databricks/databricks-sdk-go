# NEXT CHANGELOG

## Release v0.153.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `OmitPermissionsInResponse` field for [catalog.UpdatePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePermissions).
* Add `ExcludedSchemas` field for [dataclassification.CatalogConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification#CatalogConfig).
* Add `FullName` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User).
* Add `StatusMessage` field for [jobs.AiRuntimeTaskOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTaskOutput).
* Add `SqlCondition` field for [jobs.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CronSchedule).
* Add `SqlCondition` field for [jobs.TriggerInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerInfo).
* Add `SqlCondition` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
* Add `SqlCondition` field for [jobs.TriggerStateProto](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerStateProto).
* Add `FirstDistinct`, `FirstN`, `LastDistinct` and `LastN` fields for [ml.AggregationFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#AggregationFunction).
* Add `AvroSchema` and `ProtoSchema` fields for [ml.SchemaConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SchemaConfig).
* Add `FreshnessTarget` field for [ml.StreamingMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamingMode).
* Add `LongRolling` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `ReadOnlyPooledHost` and `ReadWritePooledHost` fields for [postgres.EndpointHosts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointHosts).
* Add `LastActiveTime` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `ExpireTime`, `GroupName` and `Ttl` fields for [postgres.GenerateDatabaseCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#GenerateDatabaseCredentialRequest).
* Add `ReadMetadata` enum value for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* [Breaking] Remove `Name` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User).