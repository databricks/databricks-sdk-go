# NEXT CHANGELOG

## Release v0.172.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateDirectGroupMember`, `CreateGroup`, `CreateServicePrincipal`, `CreateUser`, `CreateWorkspaceAssignment`, `DeleteDirectGroupMember`, `DeleteGroup`, `DeleteServicePrincipal`, `DeleteUser`, `DeleteWorkspaceAssignment`, `GetDirectGroupMember`, `GetGroup`, `GetServicePrincipal`, `GetUser`, `GetWorkspaceAssignment`, `ListDirectGroupMembers`, `ListGroups`, `ListServicePrincipals`, `ListTransitiveParentGroups`, `ListUsers`, `ListWorkspaceAssignments`, `UpdateGroup`, `UpdateServicePrincipal`, `UpdateUser` and `UpdateWorkspaceAssignment` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service.
* Add `CreateDirectGroupMemberProxy`, `CreateGroupProxy`, `CreateServicePrincipalProxy`, `CreateUserProxy`, `CreateWorkspaceAssignmentProxy`, `DeleteDirectGroupMemberProxy`, `DeleteGroupProxy`, `DeleteServicePrincipalProxy`, `DeleteUserProxy`, `DeleteWorkspaceAssignmentProxy`, `GetDirectGroupMemberProxy`, `GetGroupProxy`, `GetServicePrincipalProxy`, `GetUserProxy`, `GetWorkspaceAssignmentProxy`, `GetWorkspaceIdentityDetail`, `ListDirectGroupMembersProxy`, `ListGroupsProxy`, `ListServicePrincipalsProxy`, `ListTransitiveParentGroupsProxy`, `ListUsersProxy`, `ListWorkspaceAssignmentsProxy`, `UpdateGroupProxy`, `UpdateServicePrincipalProxy`, `UpdateUserProxy`, `UpdateWorkspaceAssignmentProxy` and `UpdateWorkspaceIdentityDetail` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service.
* Add `CreateTime` and `UpdateTime` fields for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).
* Change `NewCluster` field for [jobs.JobCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCluster) to no longer be required.
* [Breaking] Change `NewCluster` field for [jobs.JobCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCluster) to no longer be required.