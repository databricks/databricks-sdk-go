// Usage: openapi-codegen
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/databricks/databricks-sdk-go/openapi/code"
	"github.com/databricks/databricks-sdk-go/openapi/generator"
	"github.com/databricks/databricks-sdk-go/openapi/render"
	"github.com/databricks/databricks-sdk-go/openapi/roll"
)

var c Context

func main() {
	ctx := context.Background()
	cfg, err := render.Config()
	if err != nil {
		fmt.Printf("WARN: %s\n\n", err)
	}
	workDir, _ := os.Getwd()
	flag.StringVar(&c.Spec, "spec", cfg.Spec, "location of the spec file")
	flag.StringVar(&c.GoSDK, "gosdk", cfg.GoSDK, "location of the Go SDK")
	flag.StringVar(&c.Target, "target", workDir, "path to directory with .codegen.json")
	flag.BoolVar(&c.DryRun, "dry-run", false, "print to stdout instead of real files")
	flag.Parse()
	if c.Spec == "" {
		println("USAGE: go run openapi/gen/main.go -spec /path/to/spec.json")
		flag.PrintDefaults()
		os.Exit(1)
	}
	err = c.Run(ctx)
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
}

func (c *Context) Run(ctx context.Context) error {
	spec, err := code.NewFromFile(ctx, c.Spec)
	if err != nil {
		return fmt.Errorf("spec: %w", err)
	}
	var suite *roll.Suite
	if c.GoSDK != "" {
		suite, err = roll.NewSuite(path.Join(c.GoSDK, "internal"))
		if err != nil {
			return fmt.Errorf("examples: %w", err)
		}
	}
	gen, err := generator.NewGenerator(c.Target)
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	return gen.Apply(ctx, spec, suite)
}
