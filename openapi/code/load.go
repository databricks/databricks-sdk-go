package code

import (
	"fmt"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi"
	"golang.org/x/exp/slices"
)

type Batch struct {
	packages map[string]*Package
}

// NewFromFile loads OpenAPI specification from file
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
		packages: map[string]*Package{},
	}
	for _, tag := range spec.Tags {
		pkg, ok := batch.packages[tag.Package]
		if !ok {
			pkg = &Package{
				Named:      Named{tag.Package, ""},
				Components: spec.Components,
				services:   map[string]*Service{},
				types:      map[string]*Entity{},
			}
			batch.packages[tag.Package] = pkg
		}
		err := pkg.Load(spec, tag)
		if err != nil {
			return nil, fmt.Errorf("fail to load %s: %w", tag.Name, err)
		}
	}
	// add some packages at least some description
	for _, pkg := range batch.packages {
		if len(pkg.services) > 1 {
			continue
		}
		// we know that we have just one service
		for _, svc := range pkg.services {
			pkg.Description = svc.Summary()
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
	// add some determinism into code generation for globally stable order in
	// files like for workspaces/accounts clinets.
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
	// add some determinism into code generation for globally stable order in
	// files like api.go and/or {{.Package.Name}}.py and clients.
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
	// this requires us to do the proper sorting of services. E.g.
	//
	//  - Groups: scim.NewAccountGroups(apiClient),
	//  - ServicePrincipals: scim.NewAccountServicePrincipals(apiClient),
	//  - Users: scim.NewAccountUsers(apiClient),
	//
	// more services may follow this pattern in the future.
	norm := func(name string) string {
		if !strings.HasPrefix(name, "Account") {
			return name
		}
		// sorting-only rename: AccountGroups -> GroupsAccount
		return name[7:] + "Account"
	}
	// add some determinism into code generation for globally stable order in
	// files like api.go and/or {{.Package.Name}}.py and clients.
	slices.SortFunc(services, func(a, b *Service) bool {
		// not using .FullName() here, as in "batch" context
		// services have to be sorted globally, not only within a package.
		// alternatively we may think on adding .ReverseFullName() to sort on.
		return norm(a.Name) < norm(b.Name)
	})
	return services
}
