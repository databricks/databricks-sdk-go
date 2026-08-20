# NEXT CHANGELOG

## Release v0.171.0

### Breaking Changes

### New Features and Improvements

- Added group role assumption to built-in SDK OAuth M2M and workload identity
  federation authentication, with safeguards that prevent other built-in
  credential strategies from silently obtaining normal-access credentials.

### Bug Fixes

### Documentation

### Internal Changes

- Added group ID configuration and U2M OAuth support for CLI role assumption.
  Account authentication is rejected, while unified authentication passes the
  group ID to the OAuth server for compatibility with future support.

### API Changes
