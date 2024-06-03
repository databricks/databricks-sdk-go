// Package holds higher-level abstractions on top of OpenAPI that are used
// to generate code via text/template for Databricks SDK in different languages.
package code

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/openapi"
)

// Package represents a service package, which contains entities and interfaces
// that are relevant to a single service
type Package struct {
	Named
	Components *openapi.Components
	services   map[string]*Service
	types      map[string]*Entity
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
	pascalNameSort(types)
	return types
}

func (pkg *Package) addRecursively(service *Service, result []*Service) []*Service {
	result = append(result, service)
	subservices := make([]*Service, 0, len(service.subservices))
	for _, v := range service.subservices {
		subservices = append(subservices, v)
	}
	pascalNameSort(subservices)
	for _, svc := range subservices {
		result = pkg.addRecursively(svc, result)
	}
	return result
}

// Returns the Services sorted such has parents always come before subservices.
func (pkg *Package) ServicesSortedByParent() []*Service {
	allServices := pkg.Services()
	resultServices := []*Service{}
	for _, svc := range allServices {
		if svc.ParentService != nil {
			continue
		}
		resultServices = pkg.addRecursively(svc, resultServices)
	}

	return resultServices
}

// MainService returns a Service that matches Package name
func (pkg *Package) MainService() *Service {
	for _, svc := range pkg.services {
		if !svc.MatchesPackageName() {
			continue
		}
		return svc
	}
	return nil
}

// Types returns sorted slice of types
func (pkg *Package) Types() (types []*Entity) {
	for _, v := range pkg.types {
		types = append(types, v)
	}
	pascalNameSort(types)
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
	fullNameSort(res)
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

// schemaToEntity converts a schema into an Entity
// processedEntities keeps track of the entities that are being generated to avoid infinite recursion.
func (pkg *Package) schemaToEntity(s *openapi.Schema, path []string, hasName bool, processedEntities map[string]*Entity) *Entity {
	if s.IsRef() {
		pair := strings.Split(s.Component(), ".")
		if len(pair) == 2 && pair[0] != pkg.Name {
			schemaPackage := pair[0]
			schemaType := pair[1]
			if pkg.extImports == nil {
				pkg.extImports = map[string]*Entity{}
			}
			known, ok := pkg.extImports[s.Component()]
			if ok {
				return known
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
		return pkg.definedEntity(component, *src, processedEntities)
	}
	e := &Entity{
		Named: Named{
			Description: s.Description,
		},
		Schema:  s,
		Package: pkg,
		enum:    map[string]EnumEntry{},
	}
	if s.JsonPath != "" {
		processedEntities[s.JsonPath] = e
	}
	// pull embedded types up, if they can be defined at package level
	if s.IsDefinable() && !hasName {
		// TODO: log message or panic when overrides a type
		e.Named.Name = strings.Join(path, "")
		pkg.define(e)
	}
	e.fields = map[string]*Field{}
	e.IsAny = s.IsAny
	e.IsComputed = s.IsComputed
	e.RequiredOrder = s.Required
	// enum
	if len(s.Enum) != 0 {
		return pkg.makeEnum(e, s, path)
	}
	// object
	if len(s.Properties) != 0 {
		return pkg.makeObject(e, s, path, processedEntities)
	}
	// array
	if s.ArrayValue != nil {
		e.ArrayValue = pkg.schemaToEntity(s.ArrayValue, append(path, "Item"), false, processedEntities)
		return e
	}
	// map
	if s.MapValue != nil {
		e.MapValue = pkg.schemaToEntity(s.MapValue, path, hasName, processedEntities)
		return e
	}
	e.IsBool = s.Type == "boolean" || s.Type == "bool"
	e.IsString = s.Type == "string"
	e.IsInt64 = s.Type == "integer" && s.Format == "int64"
	e.IsFloat64 = s.Type == "number" && s.Format == "double"
	e.IsInt = s.Type == "integer" || s.Type == "int"
	return e
}

// makeObject converts OpenAPI Schema into type representation
// processedEntities keeps track of the entities that are being generated to avoid infinite recursion.
func (pkg *Package) makeObject(e *Entity, s *openapi.Schema, path []string, processedEntities map[string]*Entity) *Entity {
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
		e.fields[k] = &Field{
			Named:    named,
			Entity:   pkg.schemaToEntity(v, append(path, named.PascalName()), false, processedEntities),
			Required: required[k],
			Schema:   v,
			IsJson:   true,
		}
	}
	pkg.updateType(e)
	return e
}

var whitespace = regexp.MustCompile(`\s+`)

func (pkg *Package) makeEnum(e *Entity, s *openapi.Schema, path []string) *Entity {
	for idx, content := range s.Enum {
		name := content
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

// definedEntity defines and returns the requested entity based on the schema.
// processedEntities keeps track of the entities that are being generated to avoid infinite recursion.
func (pkg *Package) definedEntity(name string, s *openapi.Schema, processedEntities map[string]*Entity) *Entity {
	if s == nil {
		entity := &Entity{
			Named: Named{
				Name:        name,
				Description: "",
			},
			fields: map[string]*Field{},
		}
		return pkg.define(entity)
	}
	if entity, ok := processedEntities[s.JsonPath]; ok {
		// Return existing entity if it's already being generated.
		return entity
	}

	e := pkg.schemaToEntity(s, []string{name}, true, processedEntities)
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
	return pkg.define(e)
}

func (pkg *Package) define(entity *Entity) *Entity {
	k := entity.PascalName()
	_, defined := pkg.types[k]
	if defined {
		//panic(fmt.Sprintf("%s is already defined", entity.Name))
		return entity
	}
	if entity.Package == nil {
		entity.Package = pkg
	}
	pkg.types[k] = entity
	return entity
}

func (pkg *Package) updateType(entity *Entity) {
	e, defined := pkg.types[entity.PascalName()]
	if !defined {
		return
	}
	for k, v := range entity.fields {
		e.fields[k] = v
	}
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

func (pkg *Package) getService(tag *openapi.Tag) *Service {
	svc, ok := pkg.services[tag.Service]
	if !ok {
		svc = &Service{
			Package:     pkg,
			IsAccounts:  tag.IsAccounts,
			PathStyle:   tag.PathStyle,
			methods:     map[string]*Method{},
			subservices: map[string]*Service{},
			Named: Named{
				Name:        tag.Service,
				Description: tag.Description,
			},
			tag: tag,
		}
		pkg.services[tag.Service] = svc
	}
	return svc
}

// Load takes OpenAPI specification and loads a service model
func (pkg *Package) Load(ctx context.Context, spec *openapi.Specification, tag openapi.Tag) error {
	svc := pkg.getService(&tag)
	for k, v := range spec.Components.Schemas {
		split := strings.Split(k, ".")
		if split[0] != pkg.Name {
			continue
		}
		pkg.definedEntity(split[1], *v, map[string]*Entity{})
	}
	// Fill in subservice information
	if tag.ParentService != "" {
		parentTag, err := spec.GetTagByServiceName(tag.ParentService)
		if err != nil {
			return err
		}
		parentSvc := pkg.getService(parentTag)
		parentSvc.subservices[svc.Name] = svc
		svc.ParentService = parentSvc
	}

	// Link ControlPlane and DataPlane services
	if tag.ControlPlaneService != "" {
		controlPlaneTag, err := spec.GetTagByServiceName(tag.ControlPlaneService)
		if err != nil {
			return err
		}
		controlPlaneService := pkg.getService(controlPlaneTag)
		svc.ControlPlaneService = controlPlaneService
	}

	for prefix, path := range spec.Paths {
		for verb, op := range path.Verbs() {
			if op.OperationId == "Files.getStatusHead" {
				// skip this method, it needs to be removed from the spec
				continue
			}
			if !op.HasTag(tag.Name) {
				continue
			}
			logger.Infof(ctx, "pkg.Load %s %s", verb, prefix)
			params := []openapi.Parameter{}
			seenParams := map[string]bool{}
			for _, list := range [][]openapi.Parameter{path.Parameters, op.Parameters} {
				for _, v := range list {
					param, err := pkg.resolveParam(&v)
					if err != nil {
						return fmt.Errorf("could not resolve parameter %s for %s %s. This could be due to a problem in the definition of this parameter. If using $ref, ensure that $ref is used inside the 'schema' keyword", v.Name, verb, prefix)
					}
					if param == nil {
						return nil
					}
					// We don't support headers in requests.
					if param.In == "header" {
						continue
					}
					// do not propagate common path parameter to account-level APIs
					if svc.IsAccounts && param.In == "path" && param.Name == "account_id" {
						continue
					}
					if seenParams[param.Name] {
						continue
					}
					if prefix == "/api/2.0/workspace/export" && param.Name == "direct_download" {
						// prevent changing the response content type via request parameter
						// https://github.com/databricks/databricks-sdk-py/issues/104
						continue
					}
					params = append(params, *param)
					seenParams[param.Name] = true
				}
			}
			method, err := svc.newMethod(verb, prefix, params, op)
			if err != nil {
				return err
			}
			svc.methods[method.Name] = method
		}
	}
	return nil
}

func (pkg *Package) resolveParam(v *openapi.Parameter) (param *openapi.Parameter, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	param = *pkg.Components.Parameters.Resolve(v)
	return
}
