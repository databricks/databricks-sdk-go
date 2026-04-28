# NEXT CHANGELOG

## Release v0.131.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `UndeleteProject` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `ConfidentialComputeType` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
* Add `Purge` field for [postgres.DeleteProjectRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DeleteProjectRequest).
* Add `ShowDeleted` field for [postgres.ListProjectsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ListProjectsRequest).
* Add `DeleteTime` and `PurgeTime` fields for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `UcConnection` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool).
* Change `Name` field for [supervisoragents.Connection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Connection) to no longer be required.
* [Breaking] Change `Name` field for [supervisoragents.Connection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Connection) to no longer be required.
* Add [disasterrecovery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery) package.
* Add [a.DisasterRecovery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#DisasterRecoveryAPI) account-level service.
* Add `CreateExample`, `DeleteExample`, `GetExample`, `ListExamples` and `UpdateExample` methods for [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service.
* Add `DeltaTableName` field for [ml.BackfillSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#BackfillSource).
* Add `Confluence` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* [Breaking] Remove `Connection` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool).