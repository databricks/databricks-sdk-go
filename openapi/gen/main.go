// Usage: openapi-codegen
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/databricks/databricks-sdk-go/openapi/code"
	"github.com/databricks/databricks-sdk-go/openapi/render"
)

var ctx Context

func main() {
	cfg, err := render.Config()
	if err != nil {
		fmt.Printf("WARN: %s\n\n", err)
	}
	workDir, _ := os.Getwd()
	flag.StringVar(&ctx.Spec, "spec", cfg.Spec, "location of the spec file")
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
	Target string
	DryRun bool
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
	var render Render
	err = json.Unmarshal(raw, &render)
	if err != nil {
		return fmt.Errorf(".codegen.json: %w", err)
	}
	render.ctx = c
	batch, err := code.NewFromFile(c.Spec)
	if err != nil {
		return err
	}
	render.batch = batch
	return render.Run()
}

type Render struct {
	ctx       *Context
	batch     *code.Batch
	Formatter string `json:"formatter"`

	// We can generate SDKs in three modes: Packages, Types, Services
	// E.g. Go is Package-focused and Java is Types+Services
	Packages map[string]string `json:"packages,omitempty"`
	Types    map[string]string `json:"types,omitempty"`
	Services map[string]string `json:"services,omitempty"`
	Batch    map[string]string `json:"batch,omitempty"`
}

func (r *Render) Run() error {
	var filenames []string
	if r.Batch != nil {
		pass := render.NewPass(ctx.Target, []render.Named{r.batch}, r.Batch)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("batch: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if r.Packages != nil {
		pass := render.NewPass(ctx.Target, r.batch.Packages(), r.Packages)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("packages: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if r.Services != nil {
		pass := render.NewPass(ctx.Target, r.batch.Services(), r.Services)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("services: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	if r.Types != nil {
		pass := render.NewPass(ctx.Target, r.batch.Types(), r.Types)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("types: %w", err)
		}
		filenames = append(filenames, pass.Filenames...)
	}
	render.Fomratter(ctx.Target, filenames, r.Formatter)
	sort.Strings(filenames)
	sb := bytes.NewBuffer([]byte{})
	for _, v := range filenames {
		// service/*/api.go linguist-generated=true
		sb.WriteString(v)
		sb.WriteString(" linguist-generated=true\n")
	}
	genMetaFile := fmt.Sprintf("%s/.gitattributes", r.ctx.Target)
	return os.WriteFile(genMetaFile, sb.Bytes(), 0o755)
}
