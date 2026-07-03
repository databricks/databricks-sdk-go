# NEXT CHANGELOG

## Release v0.153.0

### Breaking Changes

### New Features and Improvements

- Added support for assuming a Databricks group during U2M (user-to-machine) OAuth login.
  A new `group_id` configuration field (also settable via the `DATABRICKS_GROUP_ID` environment
  variable or a `.databrickscfg` profile) and a `u2m.WithAssumeGroup` option send the numeric
  group ID as the `assume_group` query parameter on the OAuth authorize request, scoping the
  minted token to that group. Assuming a group is only supported for workspace-level logins.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
