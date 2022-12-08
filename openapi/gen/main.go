// Usage: openapi-codegen
package main

import (
	"bytes"
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

	"github.com/databricks/databricks-sdk-go/openapi/code"
)

var ctx Context

func main() {
	cfg, err := getConfig()
	if err != nil {
		cfg = &Config{}
		fmt.Printf("WARN: %s\n\n", err)
	}
	workDir, _ := os.Getwd()
	flag.StringVar(&ctx.Spec, "spec", cfg.Spec, "location of the spec file")
	flag.StringVar(&ctx.Target, "target", workDir, "path to directory with .codegen.json")
	flag.BoolVar(&ctx.DryRun, "dry-run", false, "print to stdout instead of real files")
	flag.Parse()
	if ctx.Spec == "" {
		println("USAGE: go run databricks/openapi/gen/main.go -spec /path/to/spec.json")
		flag.PrintDefaults()
		os.Exit(1)
	}
	err = ctx.Generate()
	if err != nil {
		fmt.Printf("ERROR: %s\n\n", err)
		os.Exit(1)
	}
}

type Config struct {
	Spec string `json:"spec"`
}

func getConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("home: %w", err)
	}
	loc := filepath.Join(home, ".openapi-codegen.json")
	f, err := os.Open(loc)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", loc, err)
	}
	raw, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read all: %w", err)
	}
	var config Config
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", loc, err)
	}
	return &config, nil
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
	batch, err := code.NewFromFile(c.Spec, render.IncludeTags)
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
	Packages    map[string]string `json:"packages,omitempty"`
	Types       map[string]string `json:"types,omitempty"`
	Services    map[string]string `json:"services,omitempty"`
	Batch       map[string]string `json:"batch,omitempty"`
	IncludeTags []string          `json:"includeTags,omitempty"`
}

func (r *Render) Run() error {
	var filenames []string
	if r.Batch != nil {
		pass := newPass([]named{r.batch}, r.Batch)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("batch: %w", err)
		}
		filenames = append(filenames, pass.filenames...)
	}
	if r.Packages != nil {
		pass := newPass(r.batch.Packages(), r.Packages)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("packages: %w", err)
		}
		filenames = append(filenames, pass.filenames...)
	}
	if r.Services != nil {
		pass := newPass(r.batch.Services(), r.Services)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("services: %w", err)
		}
		filenames = append(filenames, pass.filenames...)
	}
	if r.Types != nil {
		pass := newPass(r.batch.Types(), r.Types)
		err := pass.Run()
		if err != nil {
			return fmt.Errorf("types: %w", err)
		}
		filenames = append(filenames, pass.filenames...)
	}
	for _, formatter := range strings.Split(r.Formatter, "&&") {
		formatter = strings.TrimSpace(formatter)
		formatter = strings.ReplaceAll(formatter, "$FILENAMES",
			strings.Join(filenames, " "))
		split := strings.Split(formatter, " ")

		// create pipe to forward stdout and stderr to same fd,
		// so that it's clear why formatter failed.
		reader, writer := io.Pipe()
		out := bytes.NewBuffer([]byte{})
		go io.Copy(out, reader)
		defer reader.Close()
		defer writer.Close()

		cmd := exec.Command(split[0], split[1:]...)
		cmd.Dir = r.ctx.Target
		cmd.Stdout = writer
		cmd.Stderr = writer
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("%s:\n%s", formatter, out.Bytes())
		}
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
		tmpl:    template.Must(template.New("codegen").Funcs(code.HelperFuncs).ParseFiles(tmpls...)),
	}
}

type named interface {
	FullName() string
}

type Pass[T named] struct {
	Items     []T
	ctx       *Context
	tmpl      *template.Template
	fileset   map[string]string
	filenames []string
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
	p.filenames = append(p.filenames, childFilename.String())
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
