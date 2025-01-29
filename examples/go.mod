module github.com/databricks/databricks-sdk-go/examples

go 1.18

replace github.com/databricks/databricks-sdk-go/databricks => ../databricks

replace github.com/databricks/databricks-sdk-go/compute => ../compute

replace github.com/databricks/databricks-sdk-go/jobs => ../jobs

require golang.org/x/net v0.34.0

require golang.org/x/text v0.21.0 // indirect
