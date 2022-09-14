package code

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
)

type Method struct {
	Named
	Service           *Service
	Verb              string
	Path              string
	PathParts         []PathPart
	Request           *Entity
	Response          *Entity
	EmptyResponseName Named
	wait              *openapi.Wait
	pagination        *openapi.Pagination
	operation         *openapi.Operation
	shortcut          bool
}

type Shortcut struct {
	Named
	Params []Field
	Method *Method
}

type Pagination struct {
	Offset    *Field
	Limit     *Field
	Results   *Field
	Entity    *Entity
	Increment int
}

// NamedIdMap assumes that Pagination is there
type NamedIdMap struct {
	Id     *Field
	Name   *Field
	Entity *Entity

	// if List method returns []Item directly
	// without generated iteration wrapper
	Direct bool
}

type PathPart struct {
	Prefix string
	Field  *Field
}

var pathPairRE = regexp.MustCompile(`(?m)([^\{]+)(\{(\w+)\})?`)

func (m *Method) pathParams() (params []Field) {
	if len(m.PathParts) == 0 {
		return nil
	}
	if !(m.Verb == "GET" || m.Verb == "DELETE") {
		return nil
	}
	for _, part := range m.PathParts {
		if part.Field == nil {
			continue
		}
		params = append(params, *part.Field)
	}
	return params
}

func (m *Method) allowShortcut() bool {
	if m.shortcut {
		return true
	}
	if m.Service.IsRpcStyle {
		return true
	}
	return false
}

func (m *Method) rpcSingleFields() (params []Field) {
	if !m.allowShortcut() {
		return nil
	}
	if m.Request == nil {
		return nil
	}
	if len(m.Request.fields) != 1 {
		// TODO: explicitly annotate with x-databricks-shortcut
		return nil
	}
	return []Field{m.Request.Fields()[0]}
}

func (m *Method) requestShortcutFields() []Field {
	pathParams := m.pathParams()
	rpcFields := m.rpcSingleFields()
	if len(pathParams) == 0 && len(rpcFields) == 0 {
		return nil
	}
	if len(pathParams) > 0 {
		return pathParams
	}
	return rpcFields
}

func (m *Method) Shortcut() *Shortcut {
	params := m.requestShortcutFields()
	if len(params) == 0 {
		return nil
	}
	nameParts := []string{}
	for _, part := range params {
		nameParts = append(nameParts, part.PascalName())
	}
	name := fmt.Sprintf("%sBy%s", m.PascalName(), strings.Join(nameParts, "And"))
	return &Shortcut{
		Named:  Named{name, ""},
		Method: m,
		Params: params,
	}
}

func (m *Method) Wait() *Wait {
	if m.wait == nil {
		return nil
	}
	return &Wait{
		Method: m,
	}
}

func (m *Method) Pagination() *Pagination {
	if m.pagination == nil {
		return nil
	}
	if m.Response.ArrayValue != nil {
		// we assume that method already returns body-as-array
		return nil
	}
	results := m.Response.Field(m.pagination.Results)
	return &Pagination{
		Results:   results,
		Entity:    results.Entity.ArrayValue,
		Offset:    m.Request.Field(m.pagination.Offset),
		Limit:     m.Request.Field(m.pagination.Limit),
		Increment: m.pagination.Increment,
	}
}

func (m *Method) paginationItem() *Entity {
	if m.pagination == nil {
		return nil
	}
	if m.Response.ArrayValue != nil {
		// we assume that method already returns body-as-array
		return m.Response.ArrayValue
	}
	return m.Pagination().Entity
}

func (m *Method) NamedIdMap() *NamedIdMap {
	entity := m.paginationItem()
	if entity == nil {
		return nil
	}
	var id, name *Field
	for _, f := range entity.fields {
		local := f
		if f.Schema.IsIdentifier {
			id = &local
		}
		if f.Schema.IsName {
			name = &local
		}
	}
	if id == nil || name == nil {
		return nil
	}
	return &NamedIdMap{
		Id:     id,
		Name:   name,
		Entity: entity,
		Direct: m.Response.ArrayValue != nil,
	}
}

// GetByName returns method from the same service with x-databricks-crud:read
func (m *Method) GetByName() *Entity {
	n := m.NamedIdMap()
	if n == nil {
		return nil
	}
	return n.Entity
}

func (m *Method) CanHaveResponseBody() bool {
	return m.TitleVerb() == "Get" || m.TitleVerb() == "Post"
}

func (m *Method) TitleVerb() string {
	return strings.Title(strings.ToLower(m.Verb))
}
