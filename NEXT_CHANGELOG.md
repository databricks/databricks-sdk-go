# NEXT CHANGELOG

## Release v0.152.0

### Breaking Changes

* Query parameters listed in `ForceSendFields` are now sent on the wire when they hold a zero value. Previously `ForceSendFields` had no effect on query parameters (only JSON body fields were honored), so such a parameter was silently omitted; it is now serialized with its explicit value (for example `cascade=false`). Callers that unknowingly relied on the previous no-op behavior will now send the parameter.

### New Features and Improvements

### Bug Fixes

* Honor `ForceSendFields` for query parameters, including fields nested inside request sub-structs. A zero-valued scalar field (for example a `false` bool such as `cascade` on the pipelines Delete request) listed in `ForceSendFields` is now sent on the wire instead of being dropped by the `url` tag's `omitempty`, matching the existing behavior for JSON request bodies.

### Documentation

### Internal Changes

### API Changes
* Add `Etag` field for [disasterrecovery.UpdateFailoverGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#UpdateFailoverGroupRequest).