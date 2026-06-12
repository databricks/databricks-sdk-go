# NEXT CHANGELOG

## Release v0.145.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Encode the governed-tag key as a single path segment when it is sent as a URL path parameter, so hierarchical keys containing `/` route correctly instead of being split into extra path segments and resolving to no endpoint (`404` / `ENDPOINT_NOT_FOUND`). Affects `GetTagPolicy`/`DeleteTagPolicy`/`UpdateTagPolicy` ([tags.TagPoliciesAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#TagPoliciesAPI)), `GetTagAssignment`/`DeleteTagAssignment`/`UpdateTagAssignment` ([tags.WorkspaceEntityTagAssignmentsAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#WorkspaceEntityTagAssignmentsAPI)), and `Get`/`Delete`/`Update` ([catalog.EntityTagAssignmentsAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EntityTagAssignmentsAPI)).
* Deduplicate identical key/value pairs in the user agent builder ([useragent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/useragent)), so repeatedly injecting the same dimension onto a reused `context.Context` no longer grows the `User-Agent` header without bound. Distinct values for the same key are still preserved.

### Documentation

### Internal Changes

### API Changes
* Change `ResourceId` field for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to no longer be required.
* [Breaking] Change `ResourceId` field for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to no longer be required.