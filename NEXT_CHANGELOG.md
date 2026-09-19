# NEXT CHANGELOG

## Release v0.182.0

### Breaking Changes

### New Features and Improvements

* Added the `oauth-m2m-gcp` authentication type, which authenticates with a Databricks OAuth service-principal token and passes a Google Cloud access token through the `X-Databricks-GCP-SA-Access-Token` header. This lets GCP account-level provisioning APIs be called with a Databricks-governed identity when SSO is enabled and Google ID token auth is disabled. It must be selected explicitly via `AuthType: "oauth-m2m-gcp"`.

### Bug Fixes

### Documentation

* Documented the `oauth-m2m-gcp` authentication type in the GCP section of the README.

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