# NEXT CHANGELOG

## Release v0.173.0

### Breaking Changes

 * Raise the minimum supported Go version from 1.24 to 1.25

### New Features and Improvements

### Bug Fixes

- Recover from concurrent OAuth token cache writes by using the fresh token stored by another process.
- Fix Spark runtime selection for major-only runtime versions.

### Documentation

### Internal Changes

### API Changes
* Add `AgentType` field for [dashboards.GenieConversationSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversationSummary).
* Add `GroupId` field for [iamv2.DirectGroupMember](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DirectGroupMember).
* Add `EffectiveServerlessComputeId` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `PipelineChannel` field for [postgres.NewPipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#NewPipelineSpec).
* Add `TextAttachmentPurposeAnswer` enum value for [dashboards.TextAttachmentPurpose](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TextAttachmentPurpose).
* [Breaking] Add pagination for [AccountIamV2API.ListDirectGroupMembers](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListDirectGroupMembers).
* [Breaking] Add pagination for [AccountIamV2API.ListWorkspaceAssignmentDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListWorkspaceAssignmentDetails).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListDirectGroupMembersProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListDirectGroupMembersProxy).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListWorkspaceAssignmentDetailsProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListWorkspaceAssignmentDetailsProxy).
* Add `PgSpecificTypeHalfvec` and `PgSpecificTypeVarchar` enum values for [database.SyncedTableSpecPgSpecificType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpecPgSpecificType).
* Add `PgSpecificTypeHalfvec` and `PgSpecificTypeVarchar` enum values for [postgres.SyncedTableSyncedTableSpecPgSpecificType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableSpecPgSpecificType).
* [Breaking] Add pagination for [AccountIamV2API.ListWorkspaceAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListWorkspaceAssignments).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListWorkspaceAssignmentsProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListWorkspaceAssignmentsProxy).