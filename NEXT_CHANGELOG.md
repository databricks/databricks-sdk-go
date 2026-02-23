# NEXT CHANGELOG

## Release v0.111.0

### Breaking Changes

### New Features and Improvements
* Add `DefaultCredentialStrategyProvider` package-level variable to allow 
  overriding the default authentication strategy for all `Config` instances.

### Bug Fixes

### Documentation
* Added "Retries" section to README.

### Internal Changes

* Add `NewCredentialsChain` to allow internal tools to configure a custom
  credentials chain that differs from the SDK's default chain. This is not
  intended for end-user consumption.
* Export OIDC credential types (`GitHubOIDCCredentials`, `AzureDevOpsOIDCCredentials`,
  `EnvOIDCCredentials`, `FileOIDCCredentials`) so they can be used as building
  blocks in custom credentials chains.
* Make clusters creation in AWS not depend on cloud parsed from host. Changed default AWS availability for auto-created utility clusters from SPOT to SPOT_WITH_FALLBACK (API default).

### API Changes
* Add `PatchEndpoint` method for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Add `App` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Add `SecurableKind` field for [apps.AppResourceUcSecurable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurable).
* Add `ReplaceWhereOverrides` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `ReadOnlyHost` field for [postgres.EndpointHosts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointHosts).
* Add `Group` field for [postgres.EndpointSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointSpec).
* Add `Group` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `InitialEndpointSpec` field for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `MinQps` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* Add `ScalingInfo` field for [vectorsearch.EndpointInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointInfo).
* Add `Modify` enum value for [apps.AppResourceUcSecurableUcSecurablePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurableUcSecurablePermission).
* Add `HivemetastoreConnectivityFailure` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `CouldNotGetDashboardSchemaException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Add `Degraded` enum value for [postgres.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatusState).
* Add `HivemetastoreConnectivityFailure` enum value for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
