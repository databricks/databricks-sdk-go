package code

import (
	"fmt"
	"os"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
	"golang.org/x/exp/slices"
)

type Batch struct {
	Packages map[string]*Package
}

// NewFromFile loads OpenAPI specification from file
func NewFromFile(name string, includeTags []string) (*Batch, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("no %s file: %w", name, err)
	}
	defer f.Close()
	spec, err := openapi.NewFromReader(f)
	if err != nil {
		return nil, fmt.Errorf("spec from %s: %w", name, err)
	}
	batch := Batch{
		Packages: map[string]*Package{},
	}
	for _, tag := range spec.Tags {
		if len(includeTags) != 0 && !slices.Contains(includeTags, tag.Name) {
			continue
		}
		pkg, ok := batch.Packages[tag.Package]
		if !ok {
			pkg = &Package{
				Named:      Named{tag.Package, tag.Description},
				Components: spec.Components,
				services:   map[string]*Service{},
				types:      map[string]*Entity{},
			}
			batch.Packages[tag.Package] = pkg
		}
		err := pkg.Load(spec, &tag)
		if err != nil {
			return nil, fmt.Errorf("fail to load %s: %w", tag.Name, err)
		}
	}
	return &batch, nil
}

// Pkgs returns sorted slice of packages
func (b *Batch) Pkgs() (pkgs []*Package) {
	for _, pkg := range b.Packages {
		pkgs = append(pkgs, pkg)
	}
	// add some determinism into code generation
	slices.SortFunc(pkgs, func(a, b *Package) bool {
		return a.FullName() < b.FullName()
	})
	return pkgs
}

// Pkgs returns sorted slice of packages
func (b *Batch) Types() (types []*Entity) {
	for _, pkg := range b.Packages {
		types = append(types, pkg.Types()...)
	}
	// add some determinism into code generation
	slices.SortFunc(types, func(a, b *Entity) bool {
		return a.FullName() < b.FullName()
	})
	return types
}

// Pkgs returns sorted slice of packages
func (b *Batch) Services() (services []*Service) {
	for _, pkg := range b.Packages {
		services = append(services, pkg.Services()...)
	}
	// add some determinism into code generation
	slices.SortFunc(services, func(a, b *Service) bool {
		return a.FullName() < b.FullName()
	})
	return services
}
