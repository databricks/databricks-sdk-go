// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// Genie provides a no-code experience for business users, powered by AI/BI.
// Analysts set up spaces that business users can use to ask questions using
// natural language. Genie uses data registered to Unity Catalog and requires at
// least CAN USE permission on a Pro or Serverless SQL warehouse. Also,
// Databricks Assistant must be enabled.
type GenieClient struct {
	GenieAPI
}

func NewGenieClient(cfg *config.Config) (*GenieClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GenieClient{
		GenieAPI: GenieAPI{
			genieImpl: genieImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
type LakeviewClient struct {
	LakeviewAPI
}

func NewLakeviewClient(cfg *config.Config) (*LakeviewClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &LakeviewClient{
		LakeviewAPI: LakeviewAPI{
			lakeviewImpl: lakeviewImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Token-based Lakeview APIs for embedding dashboards in external applications.
type LakeviewEmbeddedClient struct {
	LakeviewEmbeddedAPI
}

func NewLakeviewEmbeddedClient(cfg *config.Config) (*LakeviewEmbeddedClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &LakeviewEmbeddedClient{
		LakeviewEmbeddedAPI: LakeviewEmbeddedAPI{
			lakeviewEmbeddedImpl: lakeviewEmbeddedImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Query execution APIs for AI / BI Dashboards
type QueryExecutionClient struct {
	QueryExecutionAPI
}

func NewQueryExecutionClient(cfg *config.Config) (*QueryExecutionClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryExecutionClient{
		QueryExecutionAPI: QueryExecutionAPI{
			queryExecutionImpl: queryExecutionImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
