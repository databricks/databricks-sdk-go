package code

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/openapi"
	"golang.org/x/exp/slices"
)

// Field of a Type (Entity)
type Field struct {
	Named
	Required bool
	Entity   *Entity
	Of       *Entity
	IsJson   bool
	IsPath   bool
	IsQuery  bool
	Schema   *openapi.Schema
}

func (f *Field) IsOptionalObject() bool {
	return f.Entity != nil && !f.Required && (f.Entity.IsObject() || f.Entity.IsExternal())
}

// IsPrivatePreview flags object being in private preview.
func (f *Field) IsPrivatePreview() bool {
	return f.Schema != nil && isPrivatePreview(&f.Schema.Node)
}

// IsPublicPreview flags object being in public preview.
func (f *Field) IsPublicPreview() bool {
	return f.Schema != nil && isPublicPreview(&f.Schema.Node)
}

type EnumEntry struct {
	Named
	Entity *Entity
	// SCIM API has schema specifiers
	Content string
}

// Entity represents a Type
type Entity struct {
	Named
	Package    *Package
	enum       map[string]EnumEntry
	ArrayValue *Entity
	MapValue   *Entity
	IsInt      bool
	IsInt64    bool
	IsFloat64  bool
	IsBool     bool
	IsString   bool
	IsEmpty    bool

	// this field does not have a concrete type
	IsAny bool

	// this field holds the identifier of the entity
	IsIdentifier bool

	// this field holds a name of the entity
	IsName bool

	// this field is computed on the platform side
	IsComputed bool

	// if entity has required fields, this is the order of them
	RequiredOrder []string
	fields        map[string]Field

	// Schema references the OpenAPI schema this entity was created from.
	Schema *openapi.Schema
}

// FullName includes package name and untransformed name of the entity
func (e *Entity) FullName() string {
	return fmt.Sprintf("%s.%s", e.Package.FullName(), e.PascalName())
}

// PascalName overrides parent implementation by appending List
// suffix for unnamed list types
func (e *Entity) PascalName() string {
	if e.Name == "" && e.ArrayValue != nil {
		return e.ArrayValue.PascalName() + "List"
	}
	return e.Named.PascalName()
}

// CamelName overrides parent implementation by appending List
// suffix for unnamed list types
func (e *Entity) CamelName() string {
	if e.Name == "" && e.ArrayValue != nil {
		return e.ArrayValue.CamelName() + "List"
	}
	return e.Named.CamelName()
}

// Field gets field representation by name or nil
func (e *Entity) Field(name string) *Field {
	if e == nil {
		// nil received: entity is not present
		return nil
	}
	field, ok := e.fields[name]
	if !ok {
		return nil
	}
	field.Of = e
	return &field
}

// IsObject returns true if entity is not a Mpa and has more than zero fields
func (e *Entity) IsObject() bool {
	return e.MapValue == nil && len(e.fields) > 0
}

// IsExternal returns true if entity is declared in external package and
// has to be imported from it
func (e *Entity) IsExternal() bool {
	return e.Package != nil && len(e.Package.types) == 0
}

func (e *Entity) RequiredFields() (fields []Field) {
	for _, r := range e.RequiredOrder {
		v := e.fields[r]
		v.Of = e
		fields = append(fields, v)
	}
	return
}

func (e *Entity) NonRequiredFields() (fields []Field) {
	required := map[string]bool{}
	for _, r := range e.RequiredOrder {
		required[r] = true
	}
	for k, v := range e.fields {
		if required[k] {
			// handled in [Entity.RequiredFields]
			continue
		}
		v.Of = e
		fields = append(fields, v)
	}
	slices.SortFunc(fields, func(a, b Field) bool {
		return a.CamelName() < b.CamelName()
	})
	return
}

// Fields returns sorted slice of field representations
func (e *Entity) Fields() (fields []Field) {
	for _, v := range e.fields {
		v.Of = e
		fields = append(fields, v)
	}
	slices.SortFunc(fields, func(a, b Field) bool {
		return a.CamelName() < b.CamelName()
	})
	return fields
}

// HasQueryField returns true if any of the fields is from query
func (e *Entity) HasQueryField() bool {
	for _, v := range e.fields {
		if v.IsQuery {
			return true
		}
	}
	return false
}

// HasJsonField returns true if any of the fields is in the body
func (e *Entity) HasJsonField() bool {
	for _, v := range e.fields {
		if v.IsJson {
			return true
		}
	}
	return false
}

// Does this type have x-databricks-id field?
func (e *Entity) HasIdentifierField() bool {
	return e.IdentifierField() != nil
}

// Return field with x-databricks-id
func (e *Entity) IdentifierField() *Field {
	for _, v := range e.fields {
		if v.Entity.IsIdentifier {
			return &v
		}
	}
	return nil
}

// Does this type have x-databricks-name field?
func (e *Entity) HasNameField() bool {
	for _, v := range e.fields {
		if v.Entity.IsName {
			return true
		}
	}
	return false
}

// Does this type have a single x-databricks-name field?
func (e *Entity) HasSingleNameField() bool {
	count := 0
	for _, v := range e.fields {
		if v.Entity.IsName {
			count++
		}
	}
	return (count == 1)
}

// Enum returns all entries for enum entities
func (e *Entity) Enum() (enum []EnumEntry) {
	for _, v := range e.enum {
		enum = append(enum, v)
	}
	slices.SortFunc(enum, func(a, b EnumEntry) bool {
		return a.Name < b.Name
	})
	return enum
}

func (e *Entity) IsPrimitive() bool {
	return e.IsNumber() || e.IsBool || e.IsString || len(e.enum) > 0
}

// IsNumber returns true if field is numeric
func (e *Entity) IsNumber() bool {
	return e.IsInt64 || e.IsInt || e.IsFloat64
}

func (e *Entity) IsOnlyPrimitiveFields() bool {
	for _, v := range e.fields {
		if !v.Entity.IsPrimitive() {
			return false
		}
	}
	return true
}

func (e *Entity) IsAllRequiredFieldsPrimitive() bool {
	for _, v := range e.RequiredFields() {
		if !v.Entity.IsPrimitive() {
			return false
		}
	}
	return true
}

// IsPrivatePreview flags object being in private preview.
func (e *Entity) IsPrivatePreview() bool {
	return e.Schema != nil && isPrivatePreview(&e.Schema.Node)
}

// IsPublicPreview flags object being in public preview.
func (e *Entity) IsPublicPreview() bool {
	return e.Schema != nil && isPublicPreview(&e.Schema.Node)
}
