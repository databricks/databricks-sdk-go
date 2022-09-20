package code

import (
	"fmt"
	"os"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
)

type Batch struct {
	Packages map[string]*Package
}

func NewFromFile(name string) (*Batch, error) {
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

func (b *Batch) Pkgs() (pkgs []*Package) {
	for _, pkg := range b.Packages {
		pkgs = append(pkgs, pkg)
	}
	return pkgs
}

func (b *Batch) Types() (types []*Entity) {
	for _, pkg := range b.Packages {
		types = append(types, pkg.Types()...)
	}
	return types
}

func (b *Batch) Services() (services []*Service) {
	for _, pkg := range b.Packages {
		services = append(services, pkg.Services()...)
	}
	return services
}
