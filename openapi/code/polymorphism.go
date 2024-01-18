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
		superTypes: map[string]*Entity{},
		types:      map[string]*Entity{},
	}
}

type polymorphism struct {
	components *openapi.Components
	superTypes map[string]*Entity
	types      map[string]*Entity
}

type discriminatorSpec struct {
	Property string
	Constant string
}

func (p *polymorphism) loadDiscriminators() (map[discriminatorSpec]string, error) {
	discriminators := map[discriminatorSpec]string{}
	for typeName, typeSchema := range p.components.Schemas {
		if typeSchema == nil {
			continue
		}
		properties := (*typeSchema).Properties
		if len(properties) == 0 {
			// not at object
			continue
		}
		for fieldName, fieldSchema := range properties {
			resolvedFieldSchema := p.components.Schemas.Resolve(fieldSchema)
			if resolvedFieldSchema == nil {
				return nil, fmt.Errorf("cannot resolve: %s.%s", typeName, fieldName)
			}
			constant := (*resolvedFieldSchema).Const
			if constant == nil {
				continue
			}
			constantValue, ok := constant.(string)
			if !ok {
				// non-string constants cannot be used as discriminators
				// see https://swagger.io/docs/specification/data-models/inheritance-and-polymorphism/
				continue
			}
			discriminator := discriminatorSpec{
				Property: fieldName,
				Constant: constantValue,
			}
			_, exists := discriminators[discriminator]
			if exists {
				return nil, fmt.Errorf("duplicate oneOf discriminator: %s=%s", discriminator.Property, discriminator.Constant)
			}
			discriminators[discriminator] = fmt.Sprintf("#/components/schemas/%s", typeName)
		}
	}
	return discriminators, nil
}

type Subtype struct {
	*Entity
	Constant string
}

func (p *polymorphism) loadSubtypes(discriminators map[discriminatorSpec]string) (map[string]*Entity, error) {
	subTypes := map[string]*Entity{}
	for typeName, typeSchema := range p.components.Schemas {
		if typeSchema == nil {
			continue
		}
		s := *typeSchema
		if len(s.OneOf) == 0 {
			continue
		}
		if s.Discriminator == nil {
			return nil, fmt.Errorf("missing discriminator: %s", typeName)
		}
		propertyName := s.Discriminator.PropertyName
		if propertyName == "" {
			return nil, fmt.Errorf("missing discriminator.propetyName: %s", typeName)
		}
		packageName, thisName, ok := strings.Cut(typeName, ".")
		if !ok {
			return nil, fmt.Errorf("corrupt type name: %s", typeName)
		}
		thisEntity := &Entity{
			Named: Named{
				Name: thisName,
			},
			Package: &Package{
				Named: Named{
					Name: packageName,
				},
			},
			Discriminator: &Field{
				Required: true,
				Named: Named{
					Name: propertyName,
				},
			},
		}
		for _, oneOf := range s.OneOf {
			if oneOf.Ref == "" {
				return nil, fmt.Errorf("oneOf can point only to a type: %s", typeName)
			}
			otherType := p.components.Schemas.Resolve(&openapi.Schema{Node: oneOf})
			if otherType == nil {
				return nil, fmt.Errorf("not found: %s", oneOf.Ref)
			}
			propertySchema, ok := (*otherType).Properties[propertyName]
			if !ok {
				return nil, fmt.Errorf("%s has no %s", oneOf.Ref, propertyName)
			}
			constantValue, ok := propertySchema.Const.(string)
			if !ok {
				return nil, fmt.Errorf("%s is not a string: %v", propertyName, propertySchema.Const)
			}
			spec := discriminatorSpec{propertyName, constantValue}
			backRef, ok := discriminators[spec]
			if !ok {
				return nil, fmt.Errorf("discriminator not found: %s", spec)
			}
			if oneOf.Ref != backRef {
				return nil, fmt.Errorf("discriminator conflict: %s and %s", oneOf.Ref, backRef)
			}
			subtypePackage, subtypeName, ok := strings.Cut(oneOf.Ref, ".")
			if !ok {
				return nil, fmt.Errorf("invalid ref: %s", oneOf.Ref)
			}
			some, exists := p.superTypes[backRef]
			if exists && !(some.Name == thisName && some.Package.Name == packageName) {
				return nil, fmt.Errorf("supertype conflict: %s already has %s", backRef, some.FullName())
			}
			p.superTypes[backRef] = thisEntity
			thisEntity.SubTypes = append(thisEntity.SubTypes, Subtype{
				// later to be linked to a real entity
				Entity: &Entity{
					Named: Named{
						Name: subtypeName,
					},
					Package: &Package{
						Named: Named{
							Name: subtypePackage,
						},
					},
				},
				Constant: constantValue,
			})
		}
		sort.Slice(thisEntity.SubTypes, func(i, j int) bool {
			return strings.Compare(thisEntity.SubTypes[i].Constant, thisEntity.SubTypes[j].Constant) > 0
		})
		subTypes[typeName] = thisEntity
	}
	return subTypes, nil
}

func (p *polymorphism) Load() error {
	discriminators, err := p.loadDiscriminators()
	if err != nil {
		return fmt.Errorf("disriminators: %w", err)
	}
	types, err := p.loadSubtypes(discriminators)
	if err != nil {
		return fmt.Errorf("subtypes: %w", err)
	}
	p.types = types
	return nil
}

func (p *polymorphism) Link(batch *Batch) error {
	for _, pkg := range batch.packages {
		for _, t := range pkg.types {
			if !t.Schema.IsOneOf() {
				continue
			}
			for _, st := range t.SubTypes {
				linkedPackage, ok := batch.packages[st.Package.Name]
				if !ok {
					return fmt.Errorf("missing package: %s", st.FullName())
				}
				linkedType, ok := linkedPackage.types[st.Name]
				if !ok {
					return fmt.Errorf("missing tyoe: %s", st.FullName())
				}
				st.Entity = linkedType
				linkedType.SuperType = t
			}
		}
	}
	return nil
}

func (p *polymorphism) Resolve(path []string) (*Entity, error) {
	return nil, nil
}
