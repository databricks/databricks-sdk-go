package render

import (
	"bytes"
	"encoding/json"
	"errors"
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

type toolConfig struct {
	Spec  string `json:"spec"`
	GoSDK string `json:"gosdk,omitempty"`
}

func Config() (toolConfig, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return toolConfig{}, fmt.Errorf("home: %w", err)
	}
	loc := filepath.Join(home, ".openapi-codegen.json")
	f, err := os.Open(loc)
	if err != nil {
		return toolConfig{}, fmt.Errorf("open %s: %w", loc, err)
	}
	raw, err := io.ReadAll(f)
	if err != nil {
		return toolConfig{}, fmt.Errorf("read all: %w", err)
	}
	var cfg toolConfig
	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("parse %s: %w", loc, err)
	}
	return cfg, nil
}

func NewPass[T Named](target string, items []T, fileset map[string]string) *pass[T] {
	var tmpls []string
	newFileset := map[string]string{}
	for filename, v := range fileset {
		filename = filepath.Join(target, filename)
		tmpls = append(tmpls, filename)
		newFileset[filepath.Base(filename)] = v
	}
	t := template.New("codegen").Funcs(code.HelperFuncs)
	t = t.Funcs(template.FuncMap{
		"load": func(tmpl string) (string, error) {
			_, err := t.ParseFiles(tmpl)
			return "", err
		},
	})
	return &pass[T]{
		Items:   items,
		target:  target,
		fileset: newFileset,
		tmpl:    template.Must(t.ParseFiles(tmpls...)),
	}
}

type Named interface {
	FullName() string
}

type pass[T Named] struct {
	Items     []T
	target    string
	tmpl      *template.Template
	fileset   map[string]string
	Filenames []string
}

func (p *pass[T]) Run() error {
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

func (p *pass[T]) File(item T, contentTRef, nameT string) error {
	nt, err := template.New("filename").Parse(nameT)
	if err != nil {
		return fmt.Errorf("parse %s: %w", nameT, err)
	}
	var contents strings.Builder
	err = p.tmpl.ExecuteTemplate(&contents, contentTRef, &item)
	if errors.Is(err, code.ErrSkipThisFile) {
		// special case for CLI generation with `{{skipThisFile}}`
		return nil
	}
	if err != nil {
		return fmt.Errorf("exec %s: %w", contentTRef, err)
	}
	var childFilename strings.Builder
	err = nt.Execute(&childFilename, item)
	if err != nil {
		return fmt.Errorf("exec %s: %w", nameT, err)
	}
	p.Filenames = append(p.Filenames, childFilename.String())
	if nameT == "stdout" {
		// print something, usually instructions for any manual work
		println(contents.String())
		return nil
	}
	targetFilename := filepath.Join(p.target, childFilename.String())
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

func Fomratter(target string, filenames []string, formatSpec string) error {
	for _, formatter := range strings.Split(formatSpec, "&&") {
		formatter = strings.TrimSpace(formatter)
		fmt.Printf("Formatting: %s\n", formatter)

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
		cmd.Dir = target
		cmd.Stdout = writer
		cmd.Stderr = writer
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("%s:\n%s", formatter, out.Bytes())
		}
	}
	return nil
}
