# NEXT CHANGELOG

## Release v0.145.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Encode the governed-tag key as a single path segment when it is sent as a URL path parameter, so hierarchical keys containing `/` route correctly instead of being split into extra path segments and resolving to no endpoint (`404` / `ENDPOINT_NOT_FOUND`). Affects `GetTagPolicy`/`DeleteTagPolicy`/`UpdateTagPolicy` ([tags.TagPoliciesAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#TagPoliciesAPI)), `GetTagAssignment`/`DeleteTagAssignment`/`UpdateTagAssignment` ([tags.WorkspaceEntityTagAssignmentsAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#WorkspaceEntityTagAssignmentsAPI)), and `Get`/`Delete`/`Update` ([catalog.EntityTagAssignmentsAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EntityTagAssignmentsAPI)).

### Documentation

### Internal Changes

### API Changes
* Change `ResourceId` field for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to no longer be required.
* [Breaking] Change `ResourceId` field for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to no longer be required.
* Add `Dynamics365` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `EndpointId` field for [vectorsearch.MiniVectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#MiniVectorIndex).
* Add `EndpointId` field for [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).