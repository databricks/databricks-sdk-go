package code

import (
	"regexp"
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
				strings.ReplaceAll(n.Name, "$", ""),
				"_", " ")),
		" ", "")
}

func (n *Named) CamelName() string {
	cc := n.PascalName()
	return strings.ToLower(cc[0:1]) + cc[1:]
}

func (n *Named) HasComment() bool {
	return n.Description != ""
}

func (n *Named) Comment(prefix string, maxLen int) string {
	if n.Description == "" {
		return ""
	}
	var lines []string
	currentLine := strings.Builder{}
	for _, v := range whitespace.Split(n.Description, -1) {
		if len(prefix)+currentLine.Len()+len(v)+1 > maxLen {
			lines = append(lines, currentLine.String())
			currentLine.Reset()
		}
		currentLine.WriteString(v)
		currentLine.WriteRune(' ')
	}
	if currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
		currentLine.Reset()
	}
	return strings.TrimLeft(prefix, "\t ") + strings.Join(lines, "\n"+prefix)
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

func (pkg *Package) schemaToEntity(s *openapi.Schema, path []string, hasName bool) *Entity {
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
	// pull embedded types up, if they can be defined at package level
	if s.IsDefinable() && !hasName {
		// TODO: log message or panic when overrides a type
		e.Named.Name = strings.Join(path, "")
		pkg.define(e)
	}
	// enum
	if len(s.Enum) != 0 {
		return pkg.makeEnum(e, s, path)
	}
	// object
	if len(s.Properties) != 0 {
		return pkg.makeObject(e, s, path)
	}
	// array
	if s.ArrayValue != nil {
		e.ArrayValue = pkg.schemaToEntity(s.ArrayValue, path, hasName)
		return e
	}
	// map
	if s.MapValue != nil {
		e.MapValue = pkg.schemaToEntity(s.MapValue, path, hasName)
		return e
	}
	e.IsBool = s.Type == "boolean" || s.Type == "bool"
	e.IsString = s.Type == "string"
	e.IsInt64 = s.Type == "integer" && s.Format == "int64"
	e.IsInt = s.Type == "integer" || s.Type == "int"
	return e
}

func (pkg *Package) makeObject(e *Entity, s *openapi.Schema, path []string) *Entity {
	e.fields = map[string]Field{}
	required := map[string]bool{}
	for _, v := range s.Required {
		required[v] = true
	}
	for k, v := range s.Properties {
		named := Named{k, v.Description}
		e.fields[k] = Field{
			Named:    named,
			Entity:   pkg.schemaToEntity(v, append(path, named.PascalName()), false),
			Required: required[k],
			IsJson:   true,
		}
	}
	return e
}

var nonAlphanum = regexp.MustCompile(`[^a-zA-Z]`)
var whitespace = regexp.MustCompile(`\s+`)

func (pkg *Package) makeEnum(e *Entity, s *openapi.Schema, path []string) *Entity {
	for idx, content := range s.Enum {
		name := content
		name = nonAlphanum.ReplaceAllString(name, " ")
		var splits []string
		for _, v := range whitespace.Split(name, -1) {
			splits = append(splits, strings.Title(strings.ToLower(v)))
		}
		name = strings.Join(splits, "")
		if len(s.AliasEnum) == len(s.Enum) {
			name = s.AliasEnum[idx]
		}
		e.Enum = append(e.Enum, EnumEntry{
			Named:   Named{name, s.EnumDescriptions[content]},
			Entity:  e,
			Content: content,
		})
	}
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
	e := pkg.schemaToEntity(s, []string{name}, true)
	if e.Name == "" {
		e.Named = Named{name, s.Description}
	}
	pkg.define(e)
	return pkg.types[e.Name]
}

func (pkg *Package) newRequest(params []openapi.Parameter, op *openapi.Operation) *Entity {
	if op.RequestBody == nil && len(params) == 0 {
		return nil
	}
	//var hasParams bool
	request := &Entity{fields: map[string]Field{}}
	schemaPath := []string{op.OperationId}
	if op.RequestBody != nil {
		request = pkg.schemaToEntity(op.RequestBody.Schema(), schemaPath, true)
	}
	for _, v := range params {
		param := *pkg.Components.Parameters.Resolve(&v)
		if param == nil {
			continue
		}
		named := Named{param.Name, param.Description}
		request.fields[param.Name] = Field{
			Named:    named,
			Required: param.Required,
			IsPath:   param.In == "path",
			IsQuery:  param.In == "query",
			Entity:   pkg.schemaToEntity(param.Schema, append(schemaPath, named.PascalName()), false),
		}
		//hasParams = true
	}
	if request.Name == "" {
		// when there was a merge of params with a request or new entity was made
		request.Name = op.OperationId + "Request"
		pkg.define(request)
	}
	return request
}

func (pkg *Package) define(entity *Entity) {
	pkg.types[entity.Name] = entity
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
