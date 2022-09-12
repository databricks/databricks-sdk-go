package code

import (
	"golang.org/x/exp/slices"
)

type Field struct {
	Named
	Required bool
	Entity   *Entity
	IsJson   bool
	IsPath   bool
	IsQuery  bool
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
	enum       map[string]EnumEntry // TODO: sort
	ArrayValue *Entity
	MapValue   *Entity
	IsInt      bool
	IsInt64    bool
	IsFloat64  bool
	IsBool     bool
	IsString   bool
	IsEmpty    bool
	fields     map[string]Field
}

func (e *Entity) Field(name string) *Field {
	field, ok := e.fields[name]
	if !ok {
		return nil
	}
	return &field
}

func (e *Entity) IsNumber() bool {
	return e.IsInt64 || e.IsInt
}

func (e *Entity) IsObject() bool {
	return len(e.fields) > 0
}

func (e *Entity) Fields() (fields []Field) {
	for _, v := range e.fields {
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
