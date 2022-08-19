package code

import (
	"strings"

	"golang.org/x/exp/slices"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
)

type Named struct {
	Name        string
	Description string
}

func (n *Named) PascalName() string {
	return strings.ReplaceAll(
		strings.Title(
			strings.ReplaceAll(
				n.Name, "_", " ")), " ", "")
}

func (n *Named) CamelName() string {
	cc := n.PascalName()
	return strings.ToLower(cc[0:1]) + cc[1:]
}

// Package represents a service package, which contains entities and interfaces
// that are relevant to a single service
type Package struct {
	Named
	Components *openapi.Components
	Methods    []Method
	types      map[string]*Entity
}

func (pkg *Package) Types() (types []*Entity) {
	for _, v := range pkg.types {
		types = append(types, v)
	}
	slices.SortFunc(types, func(a, b *Entity) bool {
		return a.Name < b.Name
	})
	return types
}

type Method struct {
	Named
	Service  *Package
	Verb     string
	Path     string
	Request  *Entity
	Response *Entity
}

func (m *Method) TitleVerb() string {
	return strings.Title(strings.ToLower(m.Verb))
}

type Field struct {
	Named
	Required bool
	Entity   *Entity
	IsJson   bool
	IsPath   bool
	IsQuery  bool
}

type EnumEntry struct {
	Named
	Entity *Entity
	// SCIM API has schema specifiers
	Content string
}

type Entity struct {
	Named
	Enum       []EnumEntry
	ArrayValue *Entity
	MapValue   *Entity
	IsInt      bool
	IsInt64    bool
	IsBool     bool
	IsString   bool
	fields     map[string]Field
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

func (pkg *Package) schemaToEntity(s *openapi.Schema) *Entity {
	if s.IsRef() {
		// if schema is src, load it to this package
		src := pkg.Components.Schemas.Resolve(s)
		if src == nil {
			return nil
		}
		return pkg.definedEntity(s.Component(), *src)
	}
	e := &Entity{
		Named: Named{
			Description: s.Description,
		},
	}
	// enum
	if len(s.Enum) != 0 && len(s.AliasEnum) == len(s.Enum) {
		for idx, v := range s.Enum {
			e.Enum = append(e.Enum, EnumEntry{
				Named:   Named{s.AliasEnum[idx], s.EnumDescriptions[v]},
				Entity:  e,
				Content: v,
			})
		}
		return e
	}
	if len(s.Enum) != 0 {
		for _, v := range s.Enum {
			e.Enum = append(e.Enum, EnumEntry{
				Named:   Named{v, s.EnumDescriptions[v]},
				Entity:  e,
				Content: v,
			})
		}
		return e
	}
	// object
	if len(s.Properties) != 0 {
		e.fields = map[string]Field{}
		required := map[string]bool{}
		for _, v := range s.Required {
			required[v] = true
		}
		for k, v := range s.Properties {
			e.fields[k] = Field{
				Named:    Named{k, v.Description},
				Entity:   pkg.schemaToEntity(v),
				Required: required[k],
				IsJson:   true,
			}
		}
		return e
	}
	// array
	if s.ArrayValue != nil {
		e.ArrayValue = pkg.schemaToEntity(s.ArrayValue)
		return e
	}
	// map
	if s.MapValue != nil {
		e.MapValue = pkg.schemaToEntity(s.MapValue)
	}
	e.IsBool = s.Type == "boolean" || s.Type == "bool"
	e.IsString = s.Type == "string"
	e.IsInt64 = s.Type == "integer" && s.Format == "int64"
	e.IsInt = s.Type == "integer" || s.Type == "int"
	return e
}

func (pkg *Package) definedEntity(name string, s *openapi.Schema) *Entity {
	if s == nil {
		return nil
	}
	if s.IsEmpty() {
		return nil
	}
	if s.IsRef() && pkg.types[s.Component()] != nil {
		// entity is defined, return from cache
		return pkg.types[s.Component()]
	}
	e := pkg.schemaToEntity(s)
	e.Named = Named{name, s.Description}
	pkg.types[name] = e
	return pkg.types[name]
}

func (pkg *Package) newRequest(params []openapi.Parameter, op *openapi.Operation) *Entity {
	if op.RequestBody == nil && len(params) == 0 {
		return nil
	}
	var hasParams bool
	request := &Entity{fields: map[string]Field{}}
	if op.RequestBody != nil {
		request = pkg.schemaToEntity(op.RequestBody.Schema())
	}
	for _, v := range params {
		param := *pkg.Components.Parameters.Resolve(&v)
		if param == nil {
			continue
		}
		request.fields[param.Name] = Field{
			Named:    Named{param.Name, param.Description},
			Required: param.Required,
			IsPath:   param.In == "path",
			IsQuery:  param.In == "query",
			Entity:   pkg.schemaToEntity(param.Schema),
		}
		hasParams = true
	}
	if !hasParams || request.Name == "" {
		// when there was a merge of params with a request or new entity was made
		request.Name = op.OperationId + "Request"
		pkg.types[request.Name] = request
	}
	return request
}

func (pkg *Package) newMethod(verb, path string, params []openapi.Parameter, op *openapi.Operation) Method {
	return Method{
		Named:   Named{op.OperationId, op.Description},
		Service: pkg,
		Verb:    verb,
		Path:    path,
		Request: pkg.newRequest(params, op),
		Response: pkg.definedEntity(op.OperationId+"Response",
			op.SuccessResponseSchema(pkg.Components)),
	}
}

func (pkg *Package) Load(spec *openapi.Specification, tag *openapi.Tag) error {
	for prefix, path := range spec.Paths {
		for verb, op := range path.Verbs() {
			if !op.HasTag(tag.Name) {
				continue
			}
			params := append(path.Parameters, op.Parameters...)
			method := pkg.newMethod(verb, prefix, params, op)
			pkg.Methods = append(pkg.Methods, method)
		}
		slices.SortFunc(pkg.Methods, func(a, b Method) bool {
			return a.Name < b.Name
		})
	}
	return nil
}
