# NEXT CHANGELOG

## Release v0.81.0

### New Features and Improvements

### Bug Fixes

- Update the `U2M` port selection mechanism to try port `8020` first and fall
  back incrementally (to port `8021`, `8022`, and so on...) if a port is not 
  available. This fixes an issue with the Databricks CLI which is only 
  allowlisted on port `8020`. 

### Documentation

### Internal Changes

### API Changes
* Added [settingsv2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2) and [tags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags) packages.
* Added [w.AppsSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsSettingsAPI) workspace-level service.
* Added [w.EntityTagAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EntityTagAssignmentsAPI) workspace-level service and [w.Rfa](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RfaAPI) workspace-level service.
* Added [a.AccountSettingsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#AccountSettingsV2API) account-level service and [w.WorkspaceSettingsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#WorkspaceSettingsV2API) workspace-level service.
* Added [w.TagPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#TagPoliciesAPI) workspace-level service.
* Added `DeleteConversationMessage`, `ListConversationMessages` and `SendMessageFeedback` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `IncludeAll` field for [dashboards.GenieListConversationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieListConversationsRequest).
* Added `EffectiveUsagePolicyId` field for [jobs.BaseJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseJob).
* Added `EffectiveUsagePolicyId` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
* Added `EffectiveUsagePolicyId` field for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
* Added `EffectiveUsagePolicyId` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
* Added `Tokens` field for [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit).
* Added `UsagePolicyId` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
* Added `EffectiveRunAs` and `RunAs` fields for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2).
* Added `CacheQueryId` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Added `ModelEndpointNameForQuery` field for [vectorsearch.EmbeddingSourceColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EmbeddingSourceColumn).
* [Breaking] Removed `EnvironmentSettings` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
* [Breaking] Removed `EnvironmentSettings` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* [Breaking] Removed `EnvironmentSettings` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
* [Breaking] Removed `Comment`, `DisplayName` and `Tags` fields for [sharing.Share](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Share).