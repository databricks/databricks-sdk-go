package code

import (
	"fmt"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi"
)

func newPolymorphism(components *openapi.Components) *polymorphism {
	return &polymorphism{
		components: components,
		types:      map[string]ChildTypes{},
	}
}

type polymorphism struct {
	components *openapi.Components
	types      map[string]ChildTypes
}

type TypeLookup []*Field

func (d TypeLookup) IsConstant() bool {
	for _, v := range d {
		if v.Schema.Const != nil {
			return true
		}
	}
	return false
}

func (d TypeLookup) Sort() {
	sort.Slice(d, func(i, j int) bool {
		return d[i].Name < d[j].Name
	})
}

func (d TypeLookup) String() string {
	var tmp []string
	for _, v := range d {
		tmp = append(tmp, fmt.Sprintf("%s=%v", v.Name, v.Schema.Const))
	}
	return strings.Join(tmp, ", ")
}

type TypeDiscriminator struct {
	*Entity
	TypeLookup
}

type ChildTypes []*TypeDiscriminator

func (d ChildTypes) IsConstant() bool {
	for _, v := range d {
		if v.IsConstant() {
			return true
		}
	}
	return false
}

func (d ChildTypes) Sort() {
	if d.IsConstant() {
		sort.Slice(d, func(i, j int) bool {
			return d[i].String() < d[j].String()
		})
		return
	}
	sort.Slice(d, func(i, j int) bool {
		return len(d[i].TypeLookup) < len(d[j].TypeLookup)
	})
}

func (p *polymorphism) unresolvedChildren() (map[string]ChildTypes, error) {
	out := map[string]ChildTypes{}
	for typeName, typeSchema := range p.components.Schemas {
		if typeSchema == nil {
			continue
		}
		s := *typeSchema
		if len(s.OneOf) == 0 {
			continue
		}
		if s.Discriminator != nil && len(s.DiscriminatorProperties) > 0 {
			return nil, fmt.Errorf("both x-databricks-discriminator-properties and discriminator there: %s", typeName)
		}
		if s.Discriminator != nil {
			s.DiscriminatorProperties = append(s.DiscriminatorProperties, s.Discriminator.PropertyName)
		}
		if len(s.DiscriminatorProperties) == 0 {
			return nil, fmt.Errorf("oneOf needs discriminators: %s", typeName)
		}
		children := ChildTypes{}
		for _, oneOf := range s.OneOf {
			if oneOf.Ref == "" {
				return nil, fmt.Errorf("oneOf can point only to a type: %s", typeName)
			}
			otherType := p.components.Schemas.Resolve(&openapi.Schema{Node: oneOf})
			if otherType == nil {
				return nil, fmt.Errorf("not found: %s", oneOf.Ref)
			}
			lookup := TypeLookup{}
			for _, propertyName := range s.DiscriminatorProperties {
				unresolved, ok := (*otherType).Properties[propertyName]
				if !ok {
					return nil, fmt.Errorf("%s has no %s", oneOf.Ref, propertyName)
				}
				propertySchema := p.components.Schemas.Resolve(unresolved)
				if propertySchema == nil {
					return nil, fmt.Errorf("cannot resolve property: %s", propertyName)
				}
				lookup = append(lookup, &Field{
					Schema: *propertySchema,
					Named: Named{
						Name: propertyName,
					},
				})
			}
			lookup.Sort()
			otherPackage, otherName, ok := strings.Cut(oneOf.Component(), ".")
			if !ok {
				return nil, fmt.Errorf("no package: %s", oneOf.Ref)
			}
			children = append(children, &TypeDiscriminator{
				Entity: &Entity{
					Named: Named{
						Name: otherName,
					},
					Package: &Package{
						Named: Named{
							Name: otherPackage,
						},
					},
				},
				TypeLookup: lookup,
			})
		}
		children.Sort()
		out[typeName] = children
	}
	return out, nil
}

type Subtype struct {
	*Entity
	Constant string
}

func (p *polymorphism) Load() error {
	types, err := p.unresolvedChildren()
	if err != nil {
		return fmt.Errorf("types: %w", err)
	}
	p.types = types
	return nil
}

func (p *polymorphism) Link(batch *Batch) error {
	for _, pkg := range batch.packages {
		for _, abstractType := range pkg.types {
			if abstractType.ChildTypes == nil {
				continue
			}
			abstractType.Package = pkg
			for _, subType := range abstractType.ChildTypes {
				linkedPackage, ok := batch.packages[subType.Package.Name]
				if !ok {
					return fmt.Errorf("missing package: %s", subType.FullName())
				}
				resolvedType, ok := linkedPackage.types[subType.Name]
				if !ok {
					return fmt.Errorf("missing type: %s", subType.FullName())
				}
				subType.Entity = resolvedType
				resolvedType.AbstractType = abstractType
				for _, field := range subType.TypeLookup {
					field.Entity = resolvedType.fields[field.Name].Entity
				}
			}
		}
	}
	return nil
}

func (p *polymorphism) Resolve(pkgName, typeName string) (*Entity, error) {
	key := fmt.Sprintf("%s.%s", pkgName, typeName)
	discriminators, ok := p.types[key]
	if !ok {
		return nil, fmt.Errorf("not found: %s", key)
	}
	fields := map[string]*Field{}
	for _, v := range discriminators[0].TypeLookup {
		fields[v.Name] = v
		v.IsJson = true
	}
	return &Entity{
		Named: Named{
			Name: typeName,
		},
		ChildTypes: discriminators,
		fields:     fields,
	}, nil
}
