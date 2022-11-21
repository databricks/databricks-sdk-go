package code

import (
	"fmt"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
	"golang.org/x/exp/slices"
)

type Batch struct {
	packages map[string]*Package
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
		packages: map[string]*Package{},
	}
	for _, tag := range spec.Tags {
		if len(includeTags) != 0 && !slices.Contains(includeTags, tag.Name) {
			continue
		}
		pkg, ok := batch.packages[tag.Package]
		if !ok {
			pkg = &Package{
				Named:      Named{tag.Package, tag.Description},
				Components: spec.Components,
				services:   map[string]*Service{},
				types:      map[string]*Entity{},
			}
			batch.packages[tag.Package] = pkg
		}
		err := pkg.Load(spec, &tag)
		if err != nil {
			return nil, fmt.Errorf("fail to load %s: %w", tag.Name, err)
		}
	}
	return &batch, nil
}

func (b *Batch) FullName() string {
	return "all"
}

// Packages returns sorted slice of packages
func (b *Batch) Packages() (pkgs []*Package) {
	for _, pkg := range b.packages {
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
	for _, pkg := range b.packages {
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
	for _, pkg := range b.packages {
		services = append(services, pkg.Services()...)
	}
	// we'll have more and more account level equivalents of APIs that are
	// currently workspace-level. In the AccountsClient we're striping
	// the `Account` prefix, so that naming and ordering is more logical.
	// this requires us to do the proper sorting of services.
	norm := func(name string) string {
		if !strings.HasPrefix(name, "Account") {
			return name
		}
		return name[7:] + "Account"
	}
	// add some determinism into code generation
	slices.SortFunc(services, func(a, b *Service) bool {
		// not using .FullName() here, as in "batch" context
		// services have to be sorted globally, not only within a package.
		// alternatively we may think on adding .ReverseFullName() to sort on.
		return norm(a.Name) < norm(b.Name)
	})
	return services
}
