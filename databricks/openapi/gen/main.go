// Usage: go run databricks/openapi/gen/main.go -spec /private/tmp/processed-databricks-workspace-all.json -target service
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/databricks/databricks-sdk-go/databricks/openapi/code"
)

var ctx Context

func main() {
	flag.StringVar(&ctx.Spec, "spec", "", "location of the spec file")
	flag.StringVar(&ctx.Target, "target", "", "path to directory with .codegen.json")
	flag.BoolVar(&ctx.DryRun, "dry-run", false, "print to stdout instead of real files")

	flag.Parse()

	if ctx.Spec == "" {
		println("USAGE: go run databricks/openapi/gen/main.go -spec /path/to/spec.json")
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := ctx.Generate()
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
	batch, err := code.NewFromFile(c.Spec)
	if err != nil {
		return err
	}
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
}

func (r *Render) Run() error {
	if r.Packages != nil {
		err := newPass(r.batch.Pkgs(), r.Packages).Run()
		if err != nil {
			return fmt.Errorf("packages: %w", err)
		}
	}
	if r.Services != nil {
		err := newPass(r.batch.Services(), r.Services).Run()
		if err != nil {
			return fmt.Errorf("services: %w", err)
		}
	}
	if r.Types != nil {
		err := newPass(r.batch.Types(), r.Types).Run()
		if err != nil {
			return fmt.Errorf("types: %w", err)
		}
	}
	split := strings.Split(r.Formatter, " ")
	cmd := exec.Command(split[0], split[1:]...)
	cmd.Dir = r.ctx.Target
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %w", r.Formatter, err)
	}
	return nil
}

func newPass[T named](items []T, fileset map[string]string) *Pass[T] {
	var tmpls []string
	newFileset := map[string]string{}
	for filename, v := range fileset {
		filename = filepath.Join(ctx.Target, filename)
		tmpls = append(tmpls, filename)
		newFileset[filepath.Base(filename)] = v
	}
	return &Pass[T]{
		Items:   items,
		ctx:     &ctx,
		fileset: newFileset,
		tmpl:    template.Must(template.ParseFiles(tmpls...)),
	}
}

type named interface {
	FullName() string
}

type Pass[T named] struct {
	Items   []T
	ctx     *Context
	tmpl    *template.Template
	fileset map[string]string
}

func (p *Pass[T]) Run() error {
	for _, item := range p.Items {
		fmt.Printf("Processing: %s\n", item.FullName())
		for k, v := range p.fileset {
			err := p.File(item, k, v)
			if err != nil {
				return fmt.Errorf("%s: %w", item.FullName(), err)
			}
		}
	}
	return nil
}

func (p *Pass[T]) File(item T, contentTRef, nameT string) error {
	nt, err := template.New("filename").Parse(nameT)
	if err != nil {
		return fmt.Errorf("parse %s: %w", nameT, err)
	}
	var childFilename, contents strings.Builder
	err = nt.Execute(&childFilename, item)
	if err != nil {
		return fmt.Errorf("exec %s: %w", nameT, err)
	}
	err = p.tmpl.ExecuteTemplate(&contents, contentTRef, &item)
	if err != nil {
		return fmt.Errorf("exec %s: %w", contentTRef, err)
	}
	if p.ctx.DryRun {
		fmt.Printf("\n---\nDRY RUN: %s\n---\n%s", &childFilename, &contents)
		return nil
	}
	if nameT == "stdout" {
		// print something, usually instructions for any manual work
		println(contents.String())
		return nil
	}
	targetFilename := filepath.Join(p.ctx.Target, childFilename.String())
	_, err = os.Stat(targetFilename)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(path.Dir(targetFilename), 0o755)
		if err != nil {
			return fmt.Errorf("cannot create parent folders: %w", err)
		}
	}
	file, err := os.OpenFile(targetFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("open %s: %w", targetFilename, err)
	}
	_, err = file.WriteString(contents.String())
	if err != nil {
		return fmt.Errorf("write %s: %w", targetFilename, err)
	}
	return file.Close()
}
