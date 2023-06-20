package openapi

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type Node struct {
	Description string `json:"description,omitempty"`
	Preview     string `json:"x-databricks-preview,omitempty"`
	Ref         string `json:"$ref,omitempty"`
}

// IsRef flags object being a reference to a component
func (n *Node) IsRef() bool {
	return n.Ref != ""
}

// Components is the basename of the reference path. Usually a class name
func (n *Node) Component() string {
	s := strings.Split(n.Ref, "/")
	return s[len(s)-1]
}

func NewFromReader(r io.Reader) (*Specification, error) {
	raw, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("cannot read openapi spec: %w", err)
	}
	var spec Specification
	err = json.Unmarshal(raw, &spec)
	if err != nil {
		return nil, fmt.Errorf("cannot parse openapi spec: %w", err)
	}
	return &spec, nil
}

type Specification struct {
	Node
	Paths      map[string]Path `json:"paths"`
	Components *Components     `json:"components"`
	Tags       []Tag           `json:"tags"`
}

type Tag struct {
	Node
	Package    string `json:"x-databricks-package"`
	PathStyle  string `json:"x-databricks-path-style"`
	Service    string `json:"x-databricks-service"`
	IsAccounts bool   `json:"x-databricks-is-accounts"`
	Name       string `json:"name"`
}

type Path struct {
	Node
	Parameters []Parameter `json:"parameters,omitempty"`
	Get        *Operation  `json:"get,omitempty"`
	Post       *Operation  `json:"post,omitempty"`
	Put        *Operation  `json:"put,omitempty"`
	Patch      *Operation  `json:"patch,omitempty"`
	Delete     *Operation  `json:"delete,omitempty"`
}

// Verbs returns a map of HTTP methods for a Path
func (path *Path) Verbs() map[string]*Operation {
	m := map[string]*Operation{}
	if path.Get != nil {
		m["GET"] = path.Get
	}
	if path.Post != nil {
		m["POST"] = path.Post
	}
	if path.Put != nil {
		m["PUT"] = path.Put
	}
	if path.Patch != nil {
		m["PATCH"] = path.Patch
	}
	if path.Delete != nil {
		m["DELETE"] = path.Delete
	}
	return m
}

// Operation is the equivalent of method
type Operation struct {
	Node
	Wait        *Wait            `json:"x-databricks-wait,omitempty"`
	Pagination  *Pagination      `json:"x-databricks-pagination,omitempty"`
	Shortcut    bool             `json:"x-databricks-shortcut,omitempty"`
	Crud        string           `json:"x-databricks-crud,omitempty"`
	JsonOnly    bool             `json:"x-databricks-cli-json-only,omitempty"`
	Summary     string           `json:"summary,omitempty"`
	OperationId string           `json:"operationId"`
	Tags        []string         `json:"tags"`
	Parameters  []Parameter      `json:"parameters,omitempty"`
	Responses   map[string]*Body `json:"responses"`
	RequestBody *Body            `json:"requestBody,omitempty"`
}

// Name is picking the last element of <ServiceName>.<method> string,
// that is coming in as part of Databricks OpenAPI spec.
func (o *Operation) Name() string {
	split := strings.Split(o.OperationId, ".")
	if len(split) == 2 {
		return split[1]
	}
	return o.OperationId
}

func (o *Operation) HasTag(tag string) bool {
	for _, v := range o.Tags {
		if v == tag {
			return true
		}
	}
	return false
}

func (o *Operation) SuccessResponseSchema(c *Components) *Schema {
	httpOk, ok := o.Responses["200"]
	if ok {
		return (*c.Responses.Resolve(httpOk)).Schema()
	}
	httpCreated, ok := o.Responses["201"]
	if ok {
		return (*c.Responses.Resolve(httpCreated)).Schema()
	}
	return nil
}

type node interface {
	IsRef() bool
	Component() string
}

type refs[T node] map[string]*T

func (c refs[T]) Resolve(item T) *T {
	if reflect.ValueOf(item).IsNil() {
		return nil
	}
	if !item.IsRef() {
		return &item
	}
	return c[item.Component()]
}

type Components struct {
	Node
	Parameters refs[*Parameter] `json:"parameters,omitempty"`
	Responses  refs[*Body]      `json:"responses,omitempty"`
	Schemas    refs[*Schema]    `json:"schemas,omitempty"`
}

type Schema struct {
	Node
	IsIdentifier     bool               `json:"x-databricks-id,omitempty"`
	IsName           bool               `json:"x-databricks-name,omitempty"`
	IsComputed       bool               `json:"x-databricks-computed,omitempty"`
	IsAny            bool               `json:"x-databricks-any,omitempty"`
	Type             string             `json:"type,omitempty"`
	Enum             []string           `json:"enum,omitempty"`
	AliasEnum        []string           `json:"x-databricks-alias-enum,omitempty"`
	EnumDescriptions map[string]string  `json:"x-databricks-enum-descriptions,omitempty"`
	Default          any                `json:"default,omitempty"`
	Example          any                `json:"example,omitempty"`
	Format           string             `json:"format,omitempty"`
	Required         []string           `json:"required,omitempty"`
	Properties       map[string]*Schema `json:"properties,omitempty"`
	ArrayValue       *Schema            `json:"items,omitempty"`
	MapValue         *Schema            `json:"additionalProperties,omitempty"`
}

func (s *Schema) IsEnum() bool {
	return len(s.Enum) != 0
}

func (s *Schema) IsObject() bool {
	return len(s.Properties) != 0
}

// IsDefinable states that type could be translated into a valid top-level type
// in Go, Python, Java, Scala, and JavaScript
func (s *Schema) IsDefinable() bool {
	return s.IsObject() || s.IsEnum()
}

func (s *Schema) IsMap() bool {
	return s.MapValue != nil
}

func (s *Schema) IsArray() bool {
	return s.ArrayValue != nil
}

func (s *Schema) IsEmpty() bool {
	if s.IsMap() {
		return false
	}
	if s.IsArray() {
		return false
	}
	if s.IsObject() {
		return false
	}
	if s.IsRef() {
		return false
	}
	if s.Type == "object" || s.Type == "" {
		return true
	}
	return false
}

type Parameter struct {
	Node
	Required bool    `json:"required,omitempty"`
	In       string  `json:"in,omitempty"`
	Name     string  `json:"name,omitempty"`
	Schema   *Schema `json:"schema,omitempty"`
}

type Body struct {
	Node
	Required bool               `json:"required,omitempty"`
	Content  map[string]Example `json:"content,omitempty"`
}

func (b *Body) Schema() *Schema {
	if b.Content == nil {
		return nil
	}
	j, ok := b.Content["application/json"]
	if !ok {
		return nil
	}
	return j.Schema
}

type Example struct {
	Node
	Schema *Schema `json:"schema,omitempty"`
}
