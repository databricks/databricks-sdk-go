// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataclassification

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Auto-tagging configuration for a classification tag. When enabled, detected
// columns are automatically tagged with Unity Catalog tags.
type AutoTaggingConfig struct {
	// Whether auto-tagging is enabled or disabled for this classification tag.
	AutoTaggingMode AutoTaggingConfigAutoTaggingMode `json:"auto_tagging_mode"`
	// The Classification Tag (e.g., "class.name", "class.location")
	ClassificationTag string `json:"classification_tag"`
}

// Auto-tagging mode.
type AutoTaggingConfigAutoTaggingMode string

const AutoTaggingConfigAutoTaggingModeAutoTaggingDisabled AutoTaggingConfigAutoTaggingMode = `AUTO_TAGGING_DISABLED`

const AutoTaggingConfigAutoTaggingModeAutoTaggingEnabled AutoTaggingConfigAutoTaggingMode = `AUTO_TAGGING_ENABLED`

// String representation for [fmt.Print]
func (f *AutoTaggingConfigAutoTaggingMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AutoTaggingConfigAutoTaggingMode) Set(v string) error {
	switch v {
	case `AUTO_TAGGING_DISABLED`, `AUTO_TAGGING_ENABLED`:
		*f = AutoTaggingConfigAutoTaggingMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO_TAGGING_DISABLED", "AUTO_TAGGING_ENABLED"`, v)
	}
}

// Values returns all possible values for AutoTaggingConfigAutoTaggingMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *AutoTaggingConfigAutoTaggingMode) Values() []AutoTaggingConfigAutoTaggingMode {
	return []AutoTaggingConfigAutoTaggingMode{
		AutoTaggingConfigAutoTaggingModeAutoTaggingDisabled,
		AutoTaggingConfigAutoTaggingModeAutoTaggingEnabled,
	}
}

// Type always returns AutoTaggingConfigAutoTaggingMode to satisfy [pflag.Value] interface
func (f *AutoTaggingConfigAutoTaggingMode) Type() string {
	return "AutoTaggingConfigAutoTaggingMode"
}

// Data Classification configuration for a Unity Catalog catalog. This message
// follows the "At Most One Resource" pattern: at most one CatalogConfig exists
// per catalog. - Full CRUD operations are supported: Create enables Data
// Classification, Delete disables it - It has no unique identifier of its own
// and uses its parent catalog's identifier (catalog_name)
type CatalogConfig struct {
	// List of auto-tagging configurations for this catalog. Empty list means no
	// auto-tagging is enabled.
	AutoTagConfigs []AutoTaggingConfig `json:"auto_tag_configs,omitempty"`
	// Schemas to include in the scan. Empty list is not supported as it results
	// in a no-op scan. If `included_schemas` is not set, all schemas are
	// scanned.
	IncludedSchemas *CatalogConfigSchemaNames `json:"included_schemas,omitempty"`
	// Resource name in the format: catalogs/{catalog_name}/config.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CatalogConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Wrapper message for a list of schema names.
type CatalogConfigSchemaNames struct {
	Names []string `json:"names"`
}

type CreateCatalogConfigRequest struct {
	// The configuration to create.
	CatalogConfig CatalogConfig `json:"catalog_config"`
	// Parent resource in the format: catalogs/{catalog_name}
	Parent string `json:"-" url:"-"`
}

type DeleteCatalogConfigRequest struct {
	// Resource name in the format: catalogs/{catalog_name}/config
	Name string `json:"-" url:"-"`
}

type GetCatalogConfigRequest struct {
	// Resource name in the format: catalogs/{catalog_name}/config
	Name string `json:"-" url:"-"`
}

type UpdateCatalogConfigRequest struct {
	// The configuration to apply to the catalog. The name field in
	// catalog_config identifies which resource to update.
	CatalogConfig CatalogConfig `json:"catalog_config"`
	// Resource name in the format: catalogs/{catalog_name}/config.
	Name string `json:"-" url:"-"`
	// Field mask specifying which fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}
