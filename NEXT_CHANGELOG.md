# NEXT CHANGELOG

## Release v0.181.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

- Reduce integration test cluster usage by replacing waiter coverage with HTTP fixtures and removing redundant live tests.

### API Changes
* Add `CreateMcpServiceUserMappedCredential`, `DeleteMcpServiceUserMappedCredential` and `GetMcpServiceUserMappedCredential` methods for [w.AiGateway](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AiGatewayAPI) workspace-level service.
* Add `AvroOptions` and `ProtobufOptions` fields for [pipelines.Transformer](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Transformer).
* Add `Avro` and `Protobuf` enum values for [pipelines.TransformerFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TransformerFormat).