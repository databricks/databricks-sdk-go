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
	// Currently it is only defined for top level schemas
	JsonPath string `json:"-"`
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
	setJsonPaths(spec)
	return &spec, nil
}

func setJsonPaths(spec Specification) {
	for name, schema := range spec.Components.Schemas {
		deref := *schema
		deref.JsonPath = fmt.Sprintf("%s,%s", "#/components/schemas/", name)
	}
}

type Specification struct {
	Node
	Paths      map[string]Path `json:"paths"`
	Components *Components     `json:"components"`
	Tags       []Tag           `json:"tags"`
}

type PathStyle string

const (
	// PathStyleRpc indicates that the endpoint is an RPC-style endpoint.
	// The endpoint path is an action, and the entity to act on is specified
	// in the request body.
	PathStyleRpc PathStyle = "rpc"

	// PathStyleRest indicates that the endpoint is a REST-style endpoint.
	// The endpoint path is a resource, and the operation to perform on the
	// resource is specified in the HTTP method.
	PathStyleRest PathStyle = "rest"
)

func (r *PathStyle) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return fmt.Errorf("cannot unmarshal RequestStyle: %w", err)
	}
	switch s {
	case "rpc", "rest":
		*r = PathStyle(s)
	default:
		return fmt.Errorf("invalid RequestStyle: %s", s)
	}
	return nil
}

type Tag struct {
	Node
	Package    string    `json:"x-databricks-package"`
	PathStyle  PathStyle `json:"x-databricks-path-style"`
	Service    string    `json:"x-databricks-service"`
	IsAccounts bool      `json:"x-databricks-is-accounts"`
	Name       string    `json:"name"`
}

type Path struct {
	Node
	Parameters []Parameter `json:"parameters,omitempty"`
	Get        *Operation  `json:"get,omitempty"`
	Head       *Operation  `json:"head,omitempty"`
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
	if path.Head != nil {
		m["HEAD"] = path.Head
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

type fieldPath []string

func (fp fieldPath) String() string {
	return strings.Join(fp, ".")
}

// Operation is the equivalent of method
type Operation struct {
	Node
	Wait       *Wait       `json:"x-databricks-wait,omitempty"`
	Pagination *Pagination `json:"x-databricks-pagination,omitempty"`
	Shortcut   bool        `json:"x-databricks-shortcut,omitempty"`
	Crud       string      `json:"x-databricks-crud,omitempty"`
	JsonOnly   bool        `json:"x-databricks-cli-json-only,omitempty"`

	// The x-databricks-path-style field indicates whether the operation has a
	// RESTful path style or a RPC style. When specified, this overrides the
	// service-level setting. Valid values are "rest" and "rpc". "rest" means
	// that the operation has a RESTful path style, i.e. the path represents
	// a resource and the HTTP method represents an action on the resource.
	// "rpc" means that the operation has a RPC style, i.e. the path represents
	// an action and the request body represents the resource.
	PathStyle PathStyle `json:"x-databricks-path-style,omitempty"`

	// The x-databricks-request-type-name field defines the name to use for
	// the request type in the generated client. This may be specified only
	// if the operation does NOT have a request body, thus only uses a request
	// type to encapsulate path and query parameters.
	RequestTypeName string `json:"x-databricks-request-type-name,omitempty"`

	// For list APIs, the path to the field in the response entity that contains
	// the resource ID.
	IdField fieldPath `json:"x-databricks-id,omitempty"`

	// For list APIs, the path to the field in the response entity that contains
	// the user-friendly name of the resource.
	NameField fieldPath `json:"x-databricks-name,omitempty"`

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

func (o *Operation) SuccessResponseBody(c *Components) *Body {
	for _, v := range []string{"200", "201"} {
		response, ok := o.Responses[v]
		if ok {
			return (*c.Responses.Resolve(response))
		}
	}
	return nil
}

func (o *Operation) HasNameField() bool {
	return len(o.NameField) > 0
}

func (o *Operation) HasIdentifierField() bool {
	return len(o.IdField) > 0
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
	Required     bool    `json:"required,omitempty"`
	In           string  `json:"in,omitempty"`
	Name         string  `json:"name,omitempty"`
	MultiSegment bool    `json:"x-databricks-multi-segment,omitempty"`
	Schema       *Schema `json:"schema,omitempty"`
}

type Body struct {
	Node
	Required bool                  `json:"required,omitempty"`
	Content  map[string]MediaType  `json:"content,omitempty"`
	Headers  map[string]*Parameter `json:"headers,omitempty"`
}

type MimeType string

const (
	MimeTypeJson        MimeType = "application/json"
	MimeTypeOctetStream MimeType = "application/octet-stream"
	MimeTypeTextPlain   MimeType = "text/plain"
)

// IsByteStream returns true if the body should be modeled as a byte stream.
// Today, we only support application/json and application/octet-stream, and non
// application/json entities are all modeled as byte streams.
func (m MimeType) IsByteStream() bool {
	return m != "" && m != MimeTypeJson
}

var allowedMimeTypes = []MimeType{
	MimeTypeJson,
	MimeTypeOctetStream,
	MimeTypeTextPlain,
}

func (b *Body) MimeTypeAndMediaType() (MimeType, *MediaType) {
	if b == nil || b.Content == nil {
		return "", nil
	}
	for _, m := range allowedMimeTypes {
		if mediaType, ok := b.Content[string(m)]; ok {
			return m, &mediaType
		}
	}
	return "", nil
}

type MediaType struct {
	Node
	Schema *Schema `json:"schema,omitempty"`
}

const MediaTypeNonJsonBodyFieldName = "contents"

func (m *MediaType) GetSchema() *Schema {
	if m == nil {
		return nil
	}
	return m.Schema
}
