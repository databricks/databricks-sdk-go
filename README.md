Databricks SDK for Go
---

Initial commit includes porting of Core Authentication layer.

TODO:
---

- [ ] Try pulling up packages for Azure and Google
- [ ] Pass tests for CommonEnvironmentClient
- [ ] CommandFactory should be done better
- [ ] Propagate optional key-value pairs to user agent
- [ ] Propagate OS as default key-value pair to user agent
- [ ] Get rid of Terraform SDK leftovers
- [ ] Mention contributors from Terraform provider side
- [ ] figure out if we should replace go-retryablehttp (MPL 2.0) with own impl, or something close to vendor/cloud.google.com/go/compute/metadata/retry.go
