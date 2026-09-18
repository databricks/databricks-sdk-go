# NEXT CHANGELOG

## Release v0.182.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateMcpServiceUserMappedCredential`, `DeleteMcpServiceUserMappedCredential` and `GetMcpServiceUserMappedCredential` methods for [w.AiGateway](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AiGatewayAPI) workspace-level service.
* Add `TelemetryExportDestinations` field for [apps.AppUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppUpdate).
* Add `Options` field for [catalog.McpServiceConfigSourceConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#McpServiceConfigSourceConnection).
* Add `HeaderAuth` field for [catalog.ModelProviderServiceConfigCustomProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigCustomProviderDirectConfig).
* Add `TimezoneId` field for [ml.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CronSchedule).
* Add `AvroOptions` and `ProtobufOptions` fields for [pipelines.Transformer](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Transformer).
* Add `TiktokAds` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `AgentService` and `Skill` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* Add `TiktokAds` and `Smartsheet` enum values for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Add `Avro` and `Protobuf` enum values for [pipelines.TransformerFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TransformerFormat).