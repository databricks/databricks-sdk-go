# NEXT CHANGELOG

## Release v0.171.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `ForwardUserAccessToken` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `ForwardUserAccessToken` field for [apps.AppUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppUpdate).
* Add `IncludeValue` field for [catalog.GetSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSecretRequest).
* Add `DockerImageUrl` field for [jobs.AiRuntimeTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTask).
* Add `InputColumn` and `OutputColumn` fields for [pipelines.Transformer](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Transformer).
* Add `Netsuite` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `GpuXlarge8` enum value for [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
* Add `GpuXlarge8` enum value for [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
* Add `LegacyMode` enum value for [settings.CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode).
* Change `AccountSpStatus` and `DisplayName` fields for [iamv2.ServicePrincipal](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ServicePrincipal) to be required.
* Change `AccountUserStatus`, `FullName` and `Username` fields for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User) to be required.
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetMcpServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetMcpServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetModelProviderServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelProviderServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetModelServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListMcpServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListMcpServicesRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListModelProviderServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelProviderServicesRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListModelServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelServicesRequest).
* [Breaking] Remove `BrowseOnly` field for [catalog.McpService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#McpService).
* [Breaking] Remove `BrowseOnly` field for [catalog.ModelProviderService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderService).
* [Breaking] Remove `BrowseOnly` field for [catalog.ModelService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelService).