package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/databricks/databricks-sdk-go/openapi/code"
	"github.com/databricks/databricks-sdk-go/openapi/render"
	"github.com/databricks/databricks-sdk-go/openapi/roll"
)

type Toolchain struct {
	Required     []string `json:"required"`
	PreSetup     []string `json:"pre_setup,omitempty"`
	PrependPath  string   `json:"prepend_path,omitempty"`
	Setup        []string `json:"setup,omitempty"`
	PostGenerate []string `json:"post_generate,omitempty"`
}

type Generator struct {
	Formatter string `json:"formatter"`

	// TemplateLibraries is a list of files containing go template definitions
	// that are reused in different stages of codegen. E.g. the "type" template
	// is needed in both the "types" and "services" stages.
	TemplateLibraries []string `json:"template_libraries,omitempty"`

	// We can generate SDKs in three modes: Packages, Types, Services
	// E.g. Go is Package-focused and Java is Types+Services
	Packages map[string]string `json:"packages,omitempty"`
	Types    map[string]string `json:"types,omitempty"`
	Services map[string]string `json:"services,omitempty"`
	Batch    map[string]string `json:"batch,omitempty"`

	// special case for usage example templates, that are generated
	// from Go SDK integration tests
	Examples map[string]string `json:"examples,omitempty"`
	Samples  map[string]string `json:"samples,omitempty"`

	// version bumps
	Version map[string]string `json:"version,omitempty"`

	// code generation toolchain configuration
	Toolchain *Toolchain `json:"toolchain,omitempty"`

	dir string
}

func NewGenerator(target string) (*Generator, error) {
	f, err := os.Open(fmt.Sprintf("%s/.codegen.json", target))
	if err != nil {
		return nil, fmt.Errorf("no .codegen.json file in %s: %w", target, err)
	}
	defer f.Close()
	raw, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read all: %w", err)
	}
	var c Generator
	err = json.Unmarshal(raw, &c)
	if err != nil {
		return nil, fmt.Errorf(".codegen.json: %w", err)
	}
	c.dir = target
	return &c, nil
}

func (c *Generator) Apply(ctx context.Context, batch *code.Batch, suite *roll.Suite) error {
	if suite != nil {
		err := suite.OptimizeWithApiSpec(batch)
		if err != nil {
			return fmt.Errorf("optimize examples: %w", err)
		}
	}
	var filenames []string
	if c.Batch != nil {
		pass := render.NewPass(c.dir, []render.Named{batch}, c.Batch, c.TemplateLibraries)
		err := pass.Run(ctx)
		if err != nil {
			return fmt.Errorf("batch: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Packages != nil {
		pass := render.NewPass(c.dir, batch.Packages(), c.Packages, c.TemplateLibraries)
		err := pass.Run(ctx)
		if err != nil {
			return fmt.Errorf("packages: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Services != nil {
		pass := render.NewPass(c.dir, batch.Services(), c.Services, c.TemplateLibraries)
		err := pass.Run(ctx)
		if err != nil {
			return fmt.Errorf("services: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Types != nil {
		pass := render.NewPass(c.dir, batch.Types(), c.Types, c.TemplateLibraries)
		err := pass.Run(ctx)
		if err != nil {
			return fmt.Errorf("types: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Examples != nil && suite != nil {
		pass := render.NewPass(c.dir, suite.ServicesExamples(), c.Examples, c.TemplateLibraries)
		err := pass.Run(ctx)
		if err != nil {
			return fmt.Errorf("examples: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Samples != nil && suite != nil {
		pass := render.NewPass(c.dir, suite.Samples(), c.Samples, c.TemplateLibraries)
		err := pass.Run(ctx)
		if err != nil {
			return fmt.Errorf("examples: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	render.Formatter(ctx, c.dir, filenames, c.Formatter)
	sort.Strings(filenames)
	sb := bytes.NewBuffer([]byte{})
	for _, v := range filenames {
		// service/*/api.go linguist-generated=true
		sb.WriteString(v)
		sb.WriteString(" linguist-generated=true\n")
	}
	genMetaFile := fmt.Sprintf("%s/.gitattributes", c.dir)
	return os.WriteFile(genMetaFile, sb.Bytes(), 0o755)
}
