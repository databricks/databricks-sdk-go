# NEXT CHANGELOG

## Release v0.129.0

### Breaking Changes

### New Features and Improvements

 * Add `u2m.WithDiscoveryHost` option to override the default `https://login.databricks.com` host used by the discovery login flow. Intended for testing and development against non-production environments.

### Bug Fixes

### Documentation

### Internal Changes

 * Remove `Experimental_IsUnifiedHost` flag from `HostType()` resolution. The method is deprecated and the Terraform provider will use its own host type detection.
 * Expanded AI agent detection: added Goose, Amp, Augment, Copilot (VS Code), Kiro, Windsurf. Honors the `AGENT=<name>` standard and falls back to `unknown` for unrecognized values. When multiple agent env vars are present (e.g. a Cursor CLI subagent invoked from Claude Code), the user-agent reports `agent/multiple`.
 * Use resolved host type from host metadata in `HostType()` method, falling back to URL pattern matching when metadata is unavailable.

### API Changes
* Add [supervisoragents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents) package.
* Add [w.SecretsUc](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecretsUcAPI) workspace-level service.
* Add [w.SupervisorAgents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#SupervisorAgentsAPI) workspace-level service.
* Add `Update` method for [w.Tokens](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokensAPI) workspace-level service.
* Add `Etag` field for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).
* Add `Etag` field for [dashboards.GenieUpdateSpaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieUpdateSpaceRequest).
* Add `BranchId` field for [postgres.BranchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatus).
* Add `CatalogId` field for [postgres.CatalogCatalogStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CatalogCatalogStatus).
* Add `DatabaseId` field for [postgres.DatabaseDatabaseStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DatabaseDatabaseStatus).
* Add `EndpointId` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `ProjectId` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `RoleId` field for [postgres.RoleRoleStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleStatus).
* Add `Project` field for [postgres.SyncedTableSyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableStatus).
* Add `Manual` field for [provisioning.CreateGcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateGcpKeyInfo).
* Add `Manual` field for [provisioning.GcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#GcpKeyInfo).
* Add `AppsRuntime` and `LakebaseRuntime` fields for [settings.CustomerFacingIngressNetworkPolicyRequestDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicyRequestDestination).
* Add `BlockedInternetDestinations` field for [settings.EgressNetworkPolicyNetworkAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyNetworkAccessPolicy).
* Add `ColumnsToSync` field for [vectorsearch.DeltaSyncVectorIndexSpecResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecResponse).
* Add `BreakingChange` enum value for [jobs.TerminationCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationCodeCode).
* [Breaking] Change `UpdateCatalogConfig` method for [w.DataClassification](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification#DataClassificationAPI) workspace-level service. Method path has changed.
* [Breaking] Change `UpdateDefaultWorkspaceBaseEnvironment` method for [w.Environments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/environments#EnvironmentsAPI) workspace-level service. Method path has changed.
* [Breaking] Change `UpdateKnowledgeAssistant` method for [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service. Method path has changed.
* [Breaking] Change `UpdateBranch`, `UpdateDatabase`, `UpdateEndpoint`, `UpdateProject` and `UpdateRole` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service. Method path has changed.
* [Breaking] Change `UpdateDefaultWarehouseOverride` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service. Method path has changed.