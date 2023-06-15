// Usage: openapi-codegen
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"sort"

	"github.com/databricks/databricks-sdk-go/openapi/code"
	"github.com/databricks/databricks-sdk-go/openapi/render"
	"github.com/databricks/databricks-sdk-go/openapi/roll"
)

var ctx Context

func main() {
	cfg, err := render.Config()
	if err != nil {
		fmt.Printf("WARN: %s\n\n", err)
	}
	workDir, _ := os.Getwd()
	flag.StringVar(&ctx.Spec, "spec", cfg.Spec, "location of the spec file")
	flag.StringVar(&ctx.GoSDK, "gosdk", cfg.GoSDK, "location of the Go SDK")
	flag.StringVar(&ctx.Target, "target", workDir, "path to directory with .codegen.json")
	flag.BoolVar(&ctx.DryRun, "dry-run", false, "print to stdout instead of real files")
	flag.Parse()
	if ctx.Spec == "" {
		println("USAGE: go run openapi/gen/main.go -spec /path/to/spec.json")
		flag.PrintDefaults()
		os.Exit(1)
	}
	err = ctx.Generate()
	if err != nil {
		fmt.Printf("ERROR: %s\n\n", err)
		os.Exit(1)
	}
}

type Context struct {
	Spec   string
	GoSDK  string
	Target string
	DryRun bool

	Formatter string `json:"formatter"`

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

	batch *code.Batch
	suite *roll.Suite
}

func (c *Context) Generate() error {
	f, err := os.Open(fmt.Sprintf("%s/.codegen.json", c.Target))
	if err != nil {
		return fmt.Errorf("no .codegen.json file in %s: %w", c.Target, err)
	}
	defer f.Close()
	raw, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read all: %w", err)
	}
	err = json.Unmarshal(raw, c)
	if err != nil {
		return fmt.Errorf(".codegen.json: %w", err)
	}
	c.batch, err = code.NewFromFile(c.Spec)
	if err != nil {
		return fmt.Errorf("spec: %w", err)
	}
	if c.GoSDK != "" {
		c.suite, err = roll.NewSuite(path.Join(c.GoSDK, "internal"))
		if err != nil {
			return fmt.Errorf("examples: %w", err)
		}
		err = c.suite.OptimizeWithApiSpec(c.batch)
		if err != nil {
			return fmt.Errorf("optimize examples: %w", err)
		}
	}
	return c.Run()
}

func (c *Context) Run() error {
	var filenames []string
	if c.Batch != nil {
		pass := render.NewPass(ctx.Target, []render.Named{c.batch}, c.Batch)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("batch: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Packages != nil {
		pass := render.NewPass(ctx.Target, c.batch.Packages(), c.Packages)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("packages: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Services != nil {
		pass := render.NewPass(ctx.Target, c.batch.Services(), c.Services)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("services: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Types != nil {
		pass := render.NewPass(ctx.Target, c.batch.Types(), c.Types)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("types: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Examples != nil && c.suite != nil {
		pass := render.NewPass(ctx.Target, c.suite.ServicesExamples(), c.Examples)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("examples: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if c.Samples != nil && c.suite != nil {
		pass := render.NewPass(ctx.Target, c.suite.Samples(), c.Samples)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("examples: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	render.Fomratter(ctx.Target, filenames, c.Formatter)
	sort.Strings(filenames)
	sb := bytes.NewBuffer([]byte{})
	for _, v := range filenames {
		// service/*/api.go linguist-generated=true
		sb.WriteString(v)
		sb.WriteString(" linguist-generated=true\n")
	}
	genMetaFile := fmt.Sprintf("%s/.gitattributes", c.Target)
	return os.WriteFile(genMetaFile, sb.Bytes(), 0o755)
}
