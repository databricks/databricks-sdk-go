# NEXT CHANGELOG

## Release v0.64.0

### New Features and Improvements
* Enabled asynchronous token refreshes by default ([#1208](https://github.com/databricks/databricks-sdk-go/pull/1208)).

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added `UpdateEndpointBudgetPolicy` and `UpdateEndpointCustomTags` methods for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Added `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* Added `PageSize` and `PageToken` fields for [compute.GetEvents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetEvents).
* Added `NextPageToken` and `PrevPageToken` fields for [compute.GetEventsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetEventsResponse).
* Added `NodeTypeFlexibility` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* Added `NodeTypeFlexibility` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* Added `EffectivePerformanceTarget` field for [jobs.RepairHistoryItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairHistoryItem).
* Added `PerformanceTarget` field for [jobs.RepairRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairRun).
* Added `BudgetPolicyId` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* Added `CustomTags` and `EffectiveBudgetPolicyId` fields for [vectorsearch.EndpointInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointInfo).
* Added `Disabled` enum value for [jobs.TerminationCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationCodeCode).
* [Breaking] Changed `CreateIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service to return [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).
* [Breaking] Changed `DeleteDataVectorIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service . HTTP method/verb has changed.
* [Breaking] Changed `DeleteDataVectorIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service with new required argument order.
* [Breaking] Changed [vectorsearch.List](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#List) to.
* [Breaking] Changed `DataArray` field for [vectorsearch.ResultData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ResultData) to type [vectorsearch.ListValueList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ListValueList).
* [Breaking] Changed waiter for [VectorSearchEndpointsAPI.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI.CreateEndpoint).
* [Breaking] Removed `NullValue` field for [vectorsearch.Value](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#Value).
