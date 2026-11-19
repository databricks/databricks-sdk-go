// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Manage data classification for Unity Catalog catalogs. Data classification
// automatically identifies and tags sensitive data (PII) in Unity Catalog
// tables. Each catalog can have at most one configuration resource that
// controls scanning behavior and auto-tagging rules.
package dataclassification

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type DataClassificationInterface interface {

	// Create Data Classification configuration for a catalog.
	//
	// Creates a new config resource, which enables Data Classification for the
	// specified catalog. The catalog must not already have a config. Creating the
	// config starts a scan of the selected schemas and schedules recurring scans.
	CreateCatalogConfig(ctx context.Context, request CreateCatalogConfigRequest) (*CatalogConfig, error)

	// Delete Data Classification configuration for a catalog.
	//
	// Deleting the config disables Data Classification for the catalog and stops
	// its scheduled scans. Tags that auto-tagging already applied to columns are
	// not removed. Deleting a config that does not exist succeeds.
	DeleteCatalogConfig(ctx context.Context, request DeleteCatalogConfigRequest) error

	// Get the Data Classification configuration for a catalog.
	GetCatalogConfig(ctx context.Context, request GetCatalogConfigRequest) (*CatalogConfig, error)

	// Update the Data Classification configuration for a catalog.
	//
	// The config must already exist for the catalog. Only the fields named in
	// `update_mask` change; the other fields keep their current values.
	UpdateCatalogConfig(ctx context.Context, request UpdateCatalogConfigRequest) (*CatalogConfig, error)
}

func NewDataClassification(client *client.DatabricksClient) *DataClassificationAPI {
	return &DataClassificationAPI{
		dataClassificationImpl: dataClassificationImpl{
			client: client,
		},
	}
}

// Manage data classification for Unity Catalog catalogs. Data classification
// automatically identifies and tags sensitive data (PII) in Unity Catalog
// tables. Each catalog can have at most one configuration resource that
// controls scanning behavior and auto-tagging rules.
type DataClassificationAPI struct {
	dataClassificationImpl
}
