package code

import (
	"fmt"
	"os"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
)

type Batch struct {
	Packages []Package
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
	batch := Batch{}
	for _, tag := range spec.Tags {
		pkg := &Package{
			Named:      Named{tag.Package, tag.Description},
			Components: spec.Components,
			types:      map[string]*Entity{},
		}
		err := pkg.Load(spec, &tag)
		if err != nil {
			return nil, fmt.Errorf("fail to load %s: %w", tag.Name, err)
		}
		batch.Packages = append(batch.Packages, *pkg)
	}
	return &batch, nil
}
