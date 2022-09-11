package code

import (
	"regexp"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
)

// Package represents a service package, which contains entities and interfaces
// that are relevant to a single service
type Package struct {
	Named
	Components *openapi.Components
	services   map[string]*Service
	types      map[string]*Entity
	emptyTypes []*Named
}

func (pkg *Package) Services() (types []*Service) {
	for _, v := range pkg.services {
		types = append(types, v)
	}
	slices.SortFunc(types, func(a, b *Service) bool {
		return a.Name < b.Name
	})
	return types
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

func (pkg *Package) EmptyTypes() (types []*Named) {
	types = append(types, pkg.emptyTypes...)
	slices.SortFunc(types, func(a, b *Named) bool {
		return a.Name < b.Name
	})
	return types
}

func (pkg *Package) schemaToEntity(s *openapi.Schema, path []string, hasName bool) *Entity {
	if s.IsRef() {
		// if schema is src, load it to this package
		src := pkg.Components.Schemas.Resolve(s)
		if src == nil {
			return nil
		}
		component := pkg.localComponent(&s.Node)
		return pkg.definedEntity(component, *src)
	}
	e := &Entity{
		Named: Named{
			Description: s.Description,
		},
		enum: map[string]EnumEntry{},
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
		e.ArrayValue = pkg.schemaToEntity(s.ArrayValue, append(path, "Item"), false)
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
	e.IsFloat64 = s.Type == "number" && s.Format == "double"
	e.IsInt = s.Type == "integer" || s.Type == "int"
	e.IsEmpty = s.IsEmpty()
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
		e.enum[content] = EnumEntry{
			Named:   Named{name, s.EnumDescriptions[content]},
			Entity:  e,
			Content: content,
		}
	}
	return e
}

func (pkg *Package) localComponent(n *openapi.Node) string {
	component := n.Component()
	if strings.HasPrefix(component, pkg.Name+".") {
		component = strings.ReplaceAll(component, pkg.Name+".", "")
	}
	return component
}

func (pkg *Package) definedEntity(name string, s *openapi.Schema) *Entity {
	if s == nil || s.IsEmpty() {
		entity := &Entity{
			Named: Named{
				Name:        name,
				Description: "",
			},
			IsEmpty: true,
		}
		return pkg.define(entity)
	}

	component := pkg.localComponent(&s.Node)
	if s.IsRef() && pkg.types[component] != nil {
		// entity is defined, return from cache
		return pkg.types[component]
	}
	e := pkg.schemaToEntity(s, []string{name}, true)
	if e == nil {
		// gets here when responses are objects with no properties
		return nil
	}
	if e.Name == "" {
		e.Named = Named{name, s.Description}
	}
	return pkg.define(e)
}

func (pkg *Package) define(entity *Entity) *Entity {
	if entity.IsEmpty {
		if slices.Contains(pkg.emptyTypes, &entity.Named) {
			//panic(fmt.Sprintf("%s is already defined", entity.Name))
			return entity
		}
		pkg.emptyTypes = append(pkg.emptyTypes, &entity.Named)
		return entity
	}
	_, defined := pkg.types[entity.Name]
	if defined {
		//panic(fmt.Sprintf("%s is already defined", entity.Name))
		return entity
	}
	pkg.types[entity.Name] = entity
	return entity
}

func (pkg *Package) HasPathParams() bool {
	for _, s := range pkg.services {
		for _, m := range s.methods {
			if len(m.PathParts) > 0 {
				return true
			}
		}
	}
	return false
}

func (pkg *Package) HasWaits() bool {
	for _, s := range pkg.services {
		for _, m := range s.methods {
			if m.wait != nil {
				return true
			}
		}
	}
	return false
}

func (pkg *Package) Load(spec *openapi.Specification, tag *openapi.Tag) error {
	for prefix, path := range spec.Paths {
		for verb, op := range path.Verbs() {
			if !op.HasTag(tag.Name) {
				continue
			}
			svc, ok := pkg.services[tag.Service]
			if !ok {
				svc = &Service{
					Package:    pkg,
					IsRpcStyle: tag.PathStyle == "rpc",
					methods:    map[string]*Method{},
					Named: Named{
						Name:        tag.Service,
						Description: tag.Description,
					},
				}
				pkg.services[tag.Service] = svc
			}
			params := []openapi.Parameter{}
			seenParams := map[string]bool{}
			for _, list := range [][]openapi.Parameter{path.Parameters, op.Parameters} {
				for _, v := range list {
					param := *pkg.Components.Parameters.Resolve(&v)
					if param == nil {
						return nil
					}
					if seenParams[param.Name] {
						continue
					}
					params = append(params, *param)
					seenParams[param.Name] = true
				}
			}
			method := svc.newMethod(verb, prefix, params, op)
			svc.methods[method.Name] = method
		}
	}
	return nil
}
