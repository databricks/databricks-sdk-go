// Usage: go run databricks/openapi/gen/main.go -spec /path/to/spec.json
package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/openapi/code"
)

var ctx Context

func main() {
	flag.StringVar(&ctx.Spec, "spec", "", "location of the spec file")
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
	DryRun bool
}

func (c *Context) Generate() error {
	batch, err := code.NewFromFile(c.Spec)
	if err != nil {
		return err
	}
	return (&Render{
		ctx:   c,
		batch: batch,
		tmpl:  template.Must(template.ParseGlob("databricks/openapi/gen/*.tmpl")),
		fileset: map[string]string{
			"api.go.tmpl":   "service/{{.Name}}/api.go",
			"model.go.tmpl": "service/{{.Name}}/model.go",
		},
	}).Run()
}

type Render struct {
	ctx     *Context
	tmpl    *template.Template
	batch   *code.Batch
	fileset map[string]string
}

func (r *Render) Run() error {
	for _, pkg := range r.batch.Packages {
		for k, v := range r.fileset {
			err := r.File(pkg, k, v)
			if err != nil {
				return fmt.Errorf("%s: %w", pkg.Name, err)
			}
		}
	}
	return nil
}

func (r *Render) File(pkg code.Package, contentTRef, nameT string) error {
	nt, err := template.New("filename").Parse(nameT)
	if err != nil {
		return fmt.Errorf("parse %s: %w", nameT, err)
	}
	var filename, contents strings.Builder
	err = nt.Execute(&filename, pkg)
	if err != nil {
		return fmt.Errorf("exec %s: %w", nameT, err)
	}
	err = r.tmpl.ExecuteTemplate(&contents, contentTRef, &pkg)
	if err != nil {
		return fmt.Errorf("exec %s: %w", contentTRef, err)
	}
	if r.ctx.DryRun {
		fmt.Printf("\n---\nDRY RUN: %s\n---\n%s", &filename, &contents)
		return nil
	}
	if nameT == "stdout" {
		// print something, usually instructions for any manual work
		println(contents.String())
		return nil
	}
	_, err = os.Stat(filename.String())
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(path.Dir(filename.String()), 0o755)
		if err != nil {
			return fmt.Errorf("cannot create parent folders: %w", err)
		}
	}
	file, err := os.OpenFile(filename.String(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("open %s: %w", &filename, err)
	}
	_, err = file.WriteString(contents.String())
	if err != nil {
		return fmt.Errorf("write %s: %w", &filename, err)
	}
	return file.Close()
}
