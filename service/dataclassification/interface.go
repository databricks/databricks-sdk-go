// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataclassification

import (
	"context"
)

// Manage data classification for Unity Catalog catalogs. Data classification
// automatically identifies and tags sensitive data (PII) in Unity Catalog
// tables. Each catalog can have at most one configuration resource that
// controls scanning behavior and auto-tagging rules.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DataClassificationService interface {

	// Create Data Classification configuration for a catalog.
	//
	// Creates a new config resource, which enables Data Classification for the
	// specified catalog. - The config must not already exist for the catalog.
	CreateCatalogConfig(ctx context.Context, request CreateCatalogConfigRequest) (*CatalogConfig, error)

	// Delete Data Classification configuration for a catalog.
	DeleteCatalogConfig(ctx context.Context, request DeleteCatalogConfigRequest) error

	// Get the Data Classification configuration for a catalog.
	GetCatalogConfig(ctx context.Context, request GetCatalogConfigRequest) (*CatalogConfig, error)

	// Update the Data Classification configuration for a catalog. - The config
	// must already exist for the catalog. - Updates fields specified in the
	// update_mask. Use update_mask field to perform partial updates of the
	// configuration.
	UpdateCatalogConfig(ctx context.Context, request UpdateCatalogConfigRequest) (*CatalogConfig, error)
}
