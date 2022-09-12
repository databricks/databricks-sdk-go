// Usage: go run databricks/openapi/gen/main.go -spec /private/tmp/processed-databricks-workspace-all.json -target service
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"text/template"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

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
	var tmpls []string
	fileset := map[string]string{}
	for filename, v := range render.Fileset {
		filename = filepath.Join(c.Target, filename)
		tmpls = append(tmpls, filename)
		fileset[filepath.Base(filename)] = v
	}
	render.Fileset = fileset
	render.tmpl = template.Must(template.ParseFiles(tmpls...))
	return render.Run()
}

type Render struct {
	ctx       *Context
	tmpl      *template.Template
	batch     *code.Batch
	Formatter string            `json:"formatter"`
	Fileset   map[string]string `json:"fileset"`
}

func (r *Render) Run() error {
	for _, pkg := range r.batch.Packages {
		fmt.Printf("Processing: %s\n", pkg.Name)
		for k, v := range r.Fileset {
			err := r.File(pkg, k, v)
			if err != nil {
				return fmt.Errorf("%s: %w", pkg.Name, err)
			}
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

func (r *Render) File(pkg *code.Package, contentTRef, nameT string) error {
	nt, err := template.New("filename").Parse(nameT)
	if err != nil {
		return fmt.Errorf("parse %s: %w", nameT, err)
	}
	var childFilename, contents strings.Builder
	err = nt.Execute(&childFilename, pkg)
	if err != nil {
		return fmt.Errorf("exec %s: %w", nameT, err)
	}
	err = r.tmpl.ExecuteTemplate(&contents, contentTRef, &pkg)
	if err != nil {
		return fmt.Errorf("exec %s: %w", contentTRef, err)
	}
	if r.ctx.DryRun {
		fmt.Printf("\n---\nDRY RUN: %s\n---\n%s", &childFilename, &contents)
		return nil
	}
	if nameT == "stdout" {
		// print something, usually instructions for any manual work
		println(contents.String())
		return nil
	}
	targetFilename := filepath.Join(r.ctx.Target, childFilename.String())
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
