# Metastore External Access Support
This change adds support for the external_access_enabled parameter in the Databricks Metastore API.

## Changes Made

1. Added the ExternalAccessEnabled field to the UpdateMetastore struct in service/catalog/model.go
2. Updated the test in service/catalog/metastores_usage_test.go to test the new field

## API Documentation

This implementation supports the external_access_enabled parameter as documented in the Databricks API documentation:
https://docs.databricks.com/api/workspace/metastores/update#external_access_enabled

This parameter allows or prevents non-Databricks Runtime clients from directly accessing entities under the metastore.

## Terraform Usage

With this change, users can add the external_access_enabled parameter to Terraform resources that use the Metastore API:

```hcl
resource "databricks_metastore" "this" {
  name                   = "my-metastore"
  storage_root           = "s3://my-bucket/my-path"
  external_access_enabled = true
}
```

Or when updating an existing metastore:

```hcl
resource "databricks_metastore" "this" {
  id                     = "metastore-id"
  external_access_enabled = true
}
```
