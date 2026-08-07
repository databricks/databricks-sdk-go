# NEXT CHANGELOG

## Release v0.171.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `IncludeValue` field for [catalog.GetSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSecretRequest).
* Add `InputColumn` and `OutputColumn` fields for [pipelines.Transformer](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Transformer).
* Add `Netsuite` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `GpuXlarge8` enum value for [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
* Add `GpuXlarge8` enum value for [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetMcpServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetMcpServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetModelProviderServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelProviderServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetModelServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListMcpServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListMcpServicesRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListModelProviderServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelProviderServicesRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListModelServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelServicesRequest).
* [Breaking] Remove `BrowseOnly` field for [catalog.McpService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#McpService).
* [Breaking] Remove `BrowseOnly` field for [catalog.ModelProviderService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderService).
* [Breaking] Remove `BrowseOnly` field for [catalog.ModelService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelService).