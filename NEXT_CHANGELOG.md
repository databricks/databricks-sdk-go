# NEXT CHANGELOG

## Release v0.132.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `ConfluenceOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `Confluence` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Change `Description` field for [supervisoragents.SupervisorAgent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#SupervisorAgent) to no longer be required.
* [Breaking] Change `Description` field for [supervisoragents.SupervisorAgent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#SupervisorAgent) to no longer be required.
* Add `CreateExample`, `DeleteExample`, `GetExample`, `GetPermissionLevels`, `GetPermissions`, `ListExamples`, `SetPermissions`, `UpdateExample` and `UpdatePermissions` methods for [w.SupervisorAgents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#SupervisorAgentsAPI) workspace-level service.
* Add `MetaAdsOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `MetaMarketing` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `MetaMarketing` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Change `Guidelines` field for [knowledgeassistants.Example](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#Example) to no longer be required.
* [Breaking] Change `Guidelines` field for [knowledgeassistants.Example](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#Example) to no longer be required.
* Change `Description` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool) to no longer be required.
* [Breaking] Change `Description` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool) to no longer be required.
* Add `Zendesk` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `R2TempCredentials` field for [catalog.TemporaryCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryCredentials).
* Add `ZendeskSupportOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `AzureKeyInfo` field for [provisioning.CreateCustomerManagedKeyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateCustomerManagedKeyRequest).
* Add `TargetQps` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* Add `ColumnsToIndex` field for [vectorsearch.DeltaSyncVectorIndexSpecRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecRequest).
* Add `ColumnsToIndex` field for [vectorsearch.DeltaSyncVectorIndexSpecResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecResponse).
* Add `RequestedTargetQps` field for [vectorsearch.EndpointScalingInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointScalingInfo).
* Add `TargetQps` field for [vectorsearch.PatchEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#PatchEndpointRequest).
* Add `Jira` and `Zendesk` enum values for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* [Breaking] Remove `MinQps` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* [Breaking] Remove `RequestedMinQps` field for [vectorsearch.EndpointScalingInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointScalingInfo).
* [Breaking] Remove `MinQps` field for [vectorsearch.PatchEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#PatchEndpointRequest).