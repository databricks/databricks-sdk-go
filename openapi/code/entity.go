package code

import (
	"fmt"
	"sort"

	"github.com/databricks/databricks-sdk-go/openapi"
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
	IsHeader bool
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

// IsRequestBodyField indicates a field which can only be set as part of request body
// There are some fields, such as PipelineId for example, which can be both used in JSON and
// as path parameters. In code generation we handle path and request body parameters separately
// by making path parameters always required. Thus, we don't need to consider such fields
// as request body fields anymore.
func (f *Field) IsRequestBodyField() bool {
	return f.IsJson && !f.IsPath && !f.IsQuery && !f.IsHeader
}

// Call the provided callback on this field and any nested fields in its entity,
// recursively.
func (f *Field) Traverse(fn func(*Field)) {
	fn(f)
	if f.Entity != nil && len(f.Entity.fields) > 0 {
		for _, field := range f.Entity.fields {
			field.Traverse(fn)
		}
	}
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
	Package      *Package
	enum         map[string]EnumEntry
	ArrayValue   *Entity
	MapValue     *Entity
	IsInt        bool
	IsInt64      bool
	IsFloat64    bool
	IsBool       bool
	IsString     bool
	IsByteStream bool
	IsEmpty      bool

	// this field does not have a concrete type
	IsAny bool

	// this field is computed on the platform side
	IsComputed bool

	// if entity has required fields, this is the order of them
	RequiredOrder []string
	fields        map[string]*Field

	// Schema references the OpenAPI schema this entity was created from.
	Schema *openapi.Schema
}

// Whether the Entity contains a basic GoLang type which is not required
func (e *Entity) ShouldIncludeForceSendFields() bool {
	for _, field := range e.fields {
		fieldType := field.Entity
		if !field.Required && fieldType.IsBasicGoLangType() {
			return true
		}
	}
	return false
}

// Whether this entity represents a basic GoLang type
func (e *Entity) IsBasicGoLangType() bool {
	return e.IsBool || e.IsInt64 || e.IsInt || e.IsFloat64 || e.IsString
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
	return field
}

// Given a list of field names, return the list of *Field objects which result
// from following the path of fields in the entity.
func (e *Entity) GetUnderlyingFields(path []string) ([]*Field, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("empty path is not allowed (entity: %s)", e.FullName())
	}
	if len(path) == 1 {
		return []*Field{e.Field(path[0])}, nil
	}
	field := e.Field(path[0])
	if field == nil {
		return nil, fmt.Errorf("field %s not found in entity %s", path[0], e.FullName())
	}
	rest, err := field.Entity.GetUnderlyingFields(path[1:])
	if err != nil {
		return nil, err
	}
	return append([]*Field{field}, rest...), nil
}

// IsObject returns true if entity is not a Mpa and has more than zero fields
func (e *Entity) IsObject() bool {
	return e.MapValue == nil && len(e.fields) > 0
}

func (e *Entity) IsMap() bool {
	return e.MapValue != nil
}

// IsExternal returns true if entity is declared in external package and
// has to be imported from it
func (e *Entity) IsExternal() bool {
	return e.Package != nil && len(e.Package.types) == 0
}

func (e *Entity) RequiredFields() (fields []*Field) {
	for _, r := range e.RequiredOrder {
		v := e.fields[r]
		v.Of = e
		fields = append(fields, v)
	}

	// Path fields should always be first in required arguments order.
	// We use stable sort to male sure we preserve the path arguments order
	sort.SliceStable(fields, func(a, b int) bool {
		return fields[a].IsPath && !fields[b].IsPath
	})
	return
}

func (e *Entity) RequiredPathFields() (fields []*Field) {
	for _, r := range e.RequiredOrder {
		v := e.fields[r]
		if !v.IsPath {
			continue
		}
		v.Of = e
		fields = append(fields, v)
	}
	return
}

func (e *Entity) RequiredRequestBodyFields() (fields []*Field) {
	for _, r := range e.RequiredOrder {
		v := e.fields[r]
		if !v.IsRequestBodyField() {
			continue
		}
		v.Of = e
		fields = append(fields, v)
	}
	return
}

func (e *Entity) NonRequiredFields() (fields []*Field) {
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
	pascalNameSort(fields)
	return
}

// Fields returns sorted slice of field representations
func (e *Entity) Fields() (fields []*Field) {
	for _, v := range e.fields {
		v.Of = e
		fields = append(fields, v)
	}
	pascalNameSort(fields)
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

// HasByteStreamField returns true if any of the fields is a ByteStream
func (e *Entity) HasByteStreamField() bool {
	for _, v := range e.fields {
		if v.Entity.IsByteStream {
			return true
		}
	}
	return false
}

// HasHeaderField returns true if any of the fields is from header
func (e *Entity) HasHeaderField() bool {
	for _, v := range e.fields {
		if v.IsHeader {
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

// Enum returns all entries for enum entities
func (e *Entity) Enum() (enum []EnumEntry) {
	for _, v := range e.enum {
		enum = append(enum, v)
	}
	sort.Slice(enum, func(i, j int) bool {
		return enum[i].Name < enum[j].Name
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

func (e *Entity) HasRequiredPathFields() bool {
	return len(e.RequiredPathFields()) > 0
}

func (e *Entity) HasRequiredRequestBodyFields() bool {
	return len(e.RequiredRequestBodyFields()) > 0
}

// IsPrivatePreview flags object being in private preview.
func (e *Entity) IsPrivatePreview() bool {
	return e.Schema != nil && isPrivatePreview(&e.Schema.Node)
}

// IsPublicPreview flags object being in public preview.
func (e *Entity) IsPublicPreview() bool {
	return e.Schema != nil && isPublicPreview(&e.Schema.Node)
}

func (e *Entity) IsRequest() bool {
	for _, svc := range e.Package.services {
		for _, m := range svc.methods {
			if m.Request == e {
				return true
			}
		}
	}
	return false
}

func (e *Entity) IsResponse() bool {
	for _, svc := range e.Package.services {
		for _, m := range svc.methods {
			if m.Response == e {
				return true
			}
		}
	}
	return false
}

func (e *Entity) IsReferred() bool {
	for _, t := range e.Package.types {
		for _, f := range t.fields {
			if f.Entity == e {
				return true
			}
		}
	}
	return false
}

func (e *Entity) Traverse(fn func(*Entity)) {
	fn(e)
	for _, f := range e.fields {
		f.Entity.Traverse(fn)
	}
}
