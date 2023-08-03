package code

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi"
)

// Method represents service RPC
type Method struct {
	Named
	Service *Service
	// HTTP method name
	Verb string
	// Full API Path, including /api/2.x prefix
	Path string
	// Slice of path params, e.g. permissions/{type}/{id}
	PathParts []PathPart
	// Request type representation
	Request *Entity
	// Response type representation
	Response          *Entity
	EmptyResponseName Named

	// The style of the request, either "rpc" or "rest". See the documentation on
	// Operation for more details.
	PathStyle openapi.PathStyle

	// For list APIs, the path of fields in the response entity to follow to get
	// the resource ID.
	IdFieldPath []*Field

	// For list APIs, the path of fields in the response entity to follow to get
	// the user-friendly name of the resource.
	NameFieldPath []*Field

	wait       *openapi.Wait
	pagination *openapi.Pagination
	operation  *openapi.Operation
	shortcut   bool
}

// Shortcut holds definition of "shortcut" methods, that are generated for
// methods with request entities only with required fields.
type Shortcut struct {
	Named
	Params []Field
	Method *Method
}

// Pagination holds definition of result iteration type per specific RPC.
// Databricks as of now has a couple different types of pagination:
//   - next_token/next_page_token + repeated field
//   - offset/limit with zero-based offsets + repeated field
//   - page/limit with 1-based pages + repeated field
//   - repeated inline field
//   - repeated field
type Pagination struct {
	Offset    *Field
	Limit     *Field
	Results   *Field
	Entity    *Entity
	Token     *Binding
	Increment int
}

// NamedIdMap depends on Pagination and is generated, when paginated item
// entity has Identifier and Name fields. End-users usually use this method for
// drop-downs or any other selectors.
type NamedIdMap struct {
	Named
	IdPath   []*Field
	NamePath []*Field
	Entity   *Entity

	// if List method returns []Item directly
	// without generated iteration wrapper
	Direct bool
}

// PathPart represents required field, that is always part of the path
type PathPart struct {
	Prefix      string
	Field       *Field
	IsAccountId bool
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
	if m.PathStyle == openapi.PathStyleRpc {
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
	return []Field{*m.Request.Fields()[0]}
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

// Shortcut creates definition from path params and single-field request entities
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

func (m *Method) IsCrudRead() bool {
	return m.operation.Crud == "read"
}

func (m *Method) IsCrudCreate() bool {
	return m.operation.Crud == "create"
}

func (m *Method) IsJsonOnly() bool {
	return m.operation.JsonOnly
}

func (m *Method) HasIdentifierField() bool {
	return len(m.IdFieldPath) > 0
}

func (m *Method) IdentifierField() *Field {
	if !m.HasIdentifierField() {
		return nil
	}
	return m.IdFieldPath[len(m.IdFieldPath)-1]
}

func (m *Method) HasNameField() bool {
	return len(m.NameFieldPath) > 0
}

// Wait returns definition for long-running operation
func (m *Method) Wait() *Wait {
	if m.wait == nil {
		return nil
	}
	// here it's easier to follow the snake_case, as success states are already
	// in the CONSTANT_CASE and it's easier to convert from constant to snake,
	// than from constant to camel or pascal.
	name := m.Service.Singular().SnakeName()
	success := strings.ToLower(strings.Join(m.wait.Success, "_or_"))
	getStatus, ok := m.Service.methods[m.wait.Poll]
	if !ok {
		panic(fmt.Errorf("cannot find %s.%s", m.Service.Name, m.wait.Poll))
	}
	name = fmt.Sprintf("wait_%s_%s_%s", getStatus.SnakeName(), name, success)
	return &Wait{
		Named: Named{
			Name: name,
		},
		Method: m,
	}
}

// Pagination returns definition for possibly multi-request result iterator
func (m *Method) Pagination() *Pagination {
	if m.pagination == nil {
		return nil
	}
	if m.Response.ArrayValue != nil {
		// we assume that method already returns body-as-array
		return nil
	}
	var token *Binding
	if m.pagination.Token != nil {
		token = &Binding{ // reuse the same datastructure as for waiters
			PollField: m.Request.Field(m.pagination.Token.Request),
			Bind:      m.Response.Field(m.pagination.Token.Response),
		}
	}
	offset := m.Request.Field(m.pagination.Offset)
	limit := m.Request.Field(m.pagination.Limit)
	results := m.Response.Field(m.pagination.Results)
	if results == nil {
		panic(fmt.Errorf("invalid results field '%v': %s",
			m.pagination.Results, m.operation.OperationId))
	}
	entity := results.Entity.ArrayValue
	increment := m.pagination.Increment
	return &Pagination{
		Results:   results,
		Token:     token,
		Entity:    entity,
		Offset:    offset,
		Limit:     limit,
		Increment: increment,
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
	p := m.Pagination()
	return p.Entity
}

func (m *Method) NeedsOffsetDedupe() bool {
	p := m.Pagination()
	return p.Offset != nil && m.HasIdentifierField()
}

func (p *Pagination) MultiRequest() bool {
	return p.Offset != nil || p.Token != nil
}

// NamedIdMap returns name-to-id mapping retrieval definition for all
// entities of a type
func (m *Method) NamedIdMap() *NamedIdMap {
	entity := m.paginationItem()
	if entity == nil || !m.HasIdentifierField() || !m.HasNameField() {
		return nil
	}
	namePath := m.NameFieldPath
	nameParts := []string{entity.PascalName()}
	for _, v := range namePath {
		nameParts = append(nameParts, v.PascalName())
	}
	nameParts = append(nameParts, "To")
	nameParts = append(nameParts, m.IdentifierField().PascalName())
	nameParts = append(nameParts, "Map")
	return &NamedIdMap{
		Named: Named{
			Name: strings.Join(nameParts, ""),
		},
		IdPath:   m.IdFieldPath,
		NamePath: namePath,
		Entity:   entity,
		Direct:   m.Response.ArrayValue != nil,
	}
}

func (n *NamedIdMap) Id() *Field {
	return n.IdPath[len(n.IdPath)-1]
}

// GetByName returns entity from the same service with x-databricks-crud:read
func (m *Method) GetByName() *Entity {
	n := m.NamedIdMap()
	if n == nil {
		return nil
	}
	potentialName := "GetBy"
	for _, v := range n.NamePath {
		potentialName += v.PascalName()
	}
	for _, other := range m.Service.methods {
		shortcut := other.Shortcut()
		if shortcut == nil {
			continue
		}
		if shortcut.PascalName() == potentialName {
			// we already have the shortcut
			return nil
		}
	}
	return n.Entity
}

func (m *Method) CanHaveResponseBody() bool {
	return m.TitleVerb() == "Get" || m.TitleVerb() == "Post"
}

func (m *Method) TitleVerb() string {
	return strings.Title(strings.ToLower(m.Verb))
}

// IsPrivatePreview flags object being in private preview.
func (m *Method) IsPrivatePreview() bool {
	return isPrivatePreview(&m.operation.Node)
}

// IsPublicPreview flags object being in public preview.
func (m *Method) IsPublicPreview() bool {
	return isPublicPreview(&m.operation.Node)
}

func (m *Method) AsFlat() *Named {
	if m.PascalName() == "CreateOboToken" {
		return &m.Named
	}
	methodWords := m.Named.splitASCII()
	svc := m.Service.Named

	remap := map[string]string{
		"ModelRegistry":   "Models",
		"Libraries":       "ClusterLibraries",
		"PolicyFamilies":  "ClusterPolicyFamilies",
		"Workspace":       "Notebooks", // or WorkspaceObjects
		"OAuthEnrollment": "Oauth",
		"CurrentUser":     "",
	}
	if replace, ok := remap[svc.PascalName()]; ok {
		svc = Named{
			Name: replace,
		}
	}

	serviceWords := svc.splitASCII()
	serviceSingularWords := svc.Singular().splitASCII()

	words := []string{}
	if len(methodWords) == 1 && strings.ToLower(methodWords[0]) == "list" {
		words = append(words, "list")
		words = append(words, serviceWords...)
	} else if methodWords[0] == "execute" {
		words = append(words, methodWords[0])
		// command_execution.execute -> execute-command-execution
		words = append(words, serviceWords[0])
	} else {
		words = append(words, methodWords[0])
		words = append(words, serviceSingularWords...)
	}
	words = append(words, methodWords[1:]...)
	// warehouses.get_workspace_warehouse_config -> get-warehouse-workspace-config
	seen := map[string]bool{}
	tmp := []string{}
	for _, w := range words {
		if seen[w] {
			continue
		}
		tmp = append(tmp, w)
		seen[w] = true
	}

	return &Named{
		Name: strings.Join(tmp, "_"),
	}
}

func (m *Method) CmdletName() string {
	words := m.AsFlat().splitASCII()
	noun := &Named{
		Name: strings.Join(words[1:], "_"),
	}
	return fmt.Sprintf("%s-Databricks%s", strings.Title(words[0]), noun.PascalName())
}
