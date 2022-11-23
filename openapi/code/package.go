// Package holds higher-level abstractions on top of OpenAPI that are used
// to generate code via text/template for Databricks SDK in different languages.
package code

import (
	"regexp"
	"sort"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/databricks/databricks-sdk-go/openapi"
)

// Package represents a service package, which contains entities and interfaces
// that are relevant to a single service
type Package struct {
	Named
	Components *openapi.Components
	services   map[string]*Service
	types      map[string]*Entity
	emptyTypes []*Named
	extImports map[string]*Entity
}

// FullName just returns pacakge name
func (pkg *Package) FullName() string {
	return pkg.CamelName()
}

// Services returns sorted slice of services
func (pkg *Package) Services() (types []*Service) {
	for _, v := range pkg.services {
		types = append(types, v)
	}
	slices.SortFunc(types, func(a, b *Service) bool {
		return a.PascalName() < b.PascalName()
	})
	return types
}

// Types returns sorted slice of types
func (pkg *Package) Types() (types []*Entity) {
	for _, v := range pkg.types {
		types = append(types, v)
	}
	slices.SortFunc(types, func(a, b *Entity) bool {
		return a.PascalName() < b.PascalName()
	})
	return types
}

// EmptyTypes returns sorted list of types without fields
func (pkg *Package) EmptyTypes() (types []*Named) {
	types = append(types, pkg.emptyTypes...)
	slices.SortFunc(types, func(a, b *Named) bool {
		return a.PascalName() < b.PascalName()
	})
	return types
}

// HasPagination returns try if any service within this package has result
// iteration
func (pkg *Package) HasPagination() bool {
	for _, v := range pkg.services {
		if v.HasPagination() {
			return true
		}
	}
	return false
}

func (pkg *Package) ImportedEntities() (res []*Entity) {
	for _, e := range pkg.extImports {
		res = append(res, e)
	}
	slices.SortFunc(res, func(a, b *Entity) bool {
		return a.FullName() < b.FullName()
	})
	return
}

func (pkg *Package) ImportedPackages() (res []string) {
	tmp := map[string]bool{}
	for _, e := range pkg.extImports {
		tmp[e.Package.Name] = true
	}
	for name := range tmp {
		res = append(res, name)
	}
	sort.Strings(res)
	return
}

func (pkg *Package) schemaToEntity(s *openapi.Schema, path []string, hasName bool) *Entity {
	if s.IsRef() {
		pair := strings.Split(s.Component(), ".")
		if len(pair) == 2 && pair[0] != pkg.Name {
			schemaPackage := pair[0]
			schemaType := pair[1]
			if pkg.extImports == nil {
				pkg.extImports = map[string]*Entity{}
			}
			// referred entity is declared in another package
			pkg.extImports[s.Component()] = &Entity{
				Named: Named{
					Name: schemaType,
				},
				Package: &Package{
					Named: Named{
						Name: schemaPackage,
					},
				},
			}
			return pkg.extImports[s.Component()]
		}
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
	e.IsAny = s.IsAny || s.Type == "object" && s.IsEmpty()
	e.IsIdentifier = s.IsIdentifier
	e.IsName = s.IsName
	e.IsComputed = s.IsComputed
	e.RequiredOrder = s.Required
	return e
}

// makeObject converts OpenAPI Schema into type representation
func (pkg *Package) makeObject(e *Entity, s *openapi.Schema, path []string) *Entity {
	e.fields = map[string]Field{}
	required := map[string]bool{}
	for _, v := range s.Required {
		required[v] = true
	}
	for k, v := range s.Properties {
		if v.Description == "" && v.IsRef() {
			vv := pkg.Components.Schemas.Resolve(v)
			if vv != nil {
				v.Description = (*vv).Description
			}
		}
		named := Named{k, v.Description}
		e.fields[k] = Field{
			Named:    named,
			Entity:   pkg.schemaToEntity(v, append(path, named.PascalName()), false),
			Required: required[k],
			Schema:   v,
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
		description := s.EnumDescriptions[content]
		e.enum[content] = EnumEntry{
			Named:   Named{name, description},
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
	if e.ArrayValue != nil {
		return e
	}
	if e.Name == "" {
		e.Named = Named{name, s.Description}
	}
	pkg.define(e)
	return pkg.types[e.Name]
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
	entity.Package = pkg
	pkg.types[entity.Name] = entity
	return entity
}

// HasPathParams returns true if any service has methods that rely on path params
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

// HasWaits returns true if any service has methods with long-running operations
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

// Load takes OpenAPI specification and loads a service model
func (pkg *Package) Load(spec *openapi.Specification, tag *openapi.Tag) error {
	accountServices := map[string]bool{}
	var accountsRE = regexp.MustCompile(`/accounts/`)
	for prefix, path := range spec.Paths {
		for _, op := range path.Verbs() {
			if !op.HasTag(tag.Name) {
				continue
			}
			if !accountsRE.MatchString(prefix) {
				continue
			}
			accountServices[tag.Service] = true
		}
	}
	for prefix, path := range spec.Paths {
		for verb, op := range path.Verbs() {
			if !op.HasTag(tag.Name) {
				continue
			}
			svc, ok := pkg.services[tag.Service]
			if !ok {
				svc = &Service{
					Package:    pkg,
					IsAccounts: accountServices[tag.Service],
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
					// do not propagate common path parameter to account-level APIs
					if svc.IsAccounts && param.In == "path" && param.Name == "account_id" {
						continue
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
