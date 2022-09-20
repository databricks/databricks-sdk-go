package code

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
	"golang.org/x/exp/slices"
)

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
	return f.Entity != nil && !f.Required && f.Entity.IsObject()
}

type EnumEntry struct {
	Named
	Entity *Entity
	// SCIM API has schema specifiers
	Content string
}

type Entity struct {
	Named
	Package    *Package
	enum       map[string]EnumEntry // TODO: sort
	ArrayValue *Entity
	MapValue   *Entity
	IsInt      bool
	IsInt64    bool
	IsFloat64  bool
	IsBool     bool
	IsString   bool
	fields     map[string]Field
}

func (e *Entity) FullName() string {
	return fmt.Sprintf("%s.%s", e.Package.Name, e.Name)
}

func (e *Entity) PascalName() string {
	if e.Name == "" && e.ArrayValue != nil {
		return e.ArrayValue.PascalName() + "List"
	}
	return e.Named.PascalName()
}

func (e *Entity) CamelName() string {
	if e.Name == "" && e.ArrayValue != nil {
		return e.ArrayValue.CamelName() + "List"
	}
	return e.Named.CamelName()
}

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

func (e *Entity) IsNumber() bool {
	return e.IsInt64 || e.IsInt
}

func (e *Entity) IsObject() bool {
	return e.MapValue == nil && len(e.fields) > 0
}

func (e *Entity) Fields() (fields []Field) {
	for _, v := range e.fields {
		v.Of = e
		fields = append(fields, v)
	}
	slices.SortFunc(fields, func(a, b Field) bool {
		return a.Name < b.Name
	})
	return fields
}

func (e *Entity) Enum() (enum []EnumEntry) {
	for _, v := range e.enum {
		enum = append(enum, v)
	}
	slices.SortFunc(enum, func(a, b EnumEntry) bool {
		return a.Name < b.Name
	})
	return enum
}
