// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobruns

import (
	"context"
	

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// The Jobs Runs API allows you to submit, repair, and cancel running jobs.
type JobrunsService interface {
}

func New(client *client.DatabricksClient) JobrunsService {
	return &JobrunsAPI{
		client: client,
	}
}

type JobrunsAPI struct {
	client *client.DatabricksClient
}

