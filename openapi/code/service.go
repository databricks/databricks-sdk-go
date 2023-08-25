package code

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi"
	"golang.org/x/exp/slices"
)

// Service represents specific Databricks API
type Service struct {
	Named
	PathStyle           openapi.PathStyle
	IsAccounts          bool
	Package             *Package
	methods             map[string]*Method
	ByPathParamsMethods []*Shortcut
	tag                 *openapi.Tag
}

// FullName holds package name and service name
func (svc *Service) FullName() string {
	return fmt.Sprintf("%s.%s", svc.Package.FullName(), svc.PascalName())
}

// MatchesPackageName if package name and service name are the same,
// e.g. `clusters` package & `Clusters` service
func (svc *Service) MatchesPackageName() bool {
	return strings.ToLower(svc.Name) == svc.Package.Name
}

// Methods returns sorted slice of methods
func (svc *Service) Methods() (methods []*Method) {
	for _, v := range svc.methods {
		methods = append(methods, v)
	}
	slices.SortFunc(methods, func(a, b *Method) bool {
		return a.CamelName() < b.CamelName()
	})
	return methods
}

// List returns a method annotated with x-databricks-crud:list
func (svc *Service) List() *Method {
	for _, v := range svc.methods {
		if v.Operation.Crud == "list" {
			return v
		}
	}
	return nil
}

// List returns a method annotated with x-databricks-crud:create
func (svc *Service) Create() *Method {
	for _, v := range svc.methods {
		if v.Operation.Crud == "create" {
			return v
		}
	}
	return nil
}

// List returns a method annotated with x-databricks-crud:read
func (svc *Service) Read() *Method {
	for _, v := range svc.methods {
		if v.Operation.Crud == "read" {
			return v
		}
	}
	return nil
}

// List returns a method annotated with x-databricks-crud:update
func (svc *Service) Update() *Method {
	for _, v := range svc.methods {
		if v.Operation.Crud == "update" {
			return v
		}
	}
	return nil
}

// List returns a method annotated with x-databricks-crud:delete
func (svc *Service) Delete() *Method {
	for _, v := range svc.methods {
		if v.Operation.Crud == "delete" {
			return v
		}
	}
	return nil
}

// HasPagination returns true if any method has result iteration
func (svc *Service) HasPagination() bool {
	for _, v := range svc.methods {
		p := v.pagination
		if p == nil {
			continue
		}
		if p.Offset != "" || p.Token != nil {
			return true
		}
	}
	return false
}

func (svc *Service) paramToField(op *openapi.Operation, param openapi.Parameter) *Field {
	named := Named{param.Name, param.Description}
	return &Field{
		Named:    named,
		Required: param.Required,
		IsPath:   param.In == "path",
		IsQuery:  param.In == "query",
		Entity: svc.Package.schemaToEntity(param.Schema, []string{
			op.Name(),
			named.PascalName(),
		}, false),
	}
}

var crudNames = map[string]bool{
	"create":  true,
	"read":    true,
	"get":     true,
	"update":  true,
	"replace": true,
	"delete":  true,
	"list":    true,
	"restore": true,
}

// Returns the schema representing a request body for a given operation, along with the mime type.
// For requests whose mime type is not application/json, the request body is always a byte stream.
func (svc *Service) getBaseSchemaAndMimeType(body *openapi.Body) (*openapi.Schema, openapi.MimeType) {
	mimeType, mediaType := body.MimeTypeAndMediaType()
	schema := mediaType.GetSchema()
	if mimeType.IsByteStream() {
		schema = &openapi.Schema{
			Type: "object",
			Properties: map[string]*openapi.Schema{
				openapi.MediaTypeNonJsonBodyFieldName: schema,
			},
		}
	}
	return schema, mimeType
}

func (svc *Service) updateEntityTypeFromMimeType(entity *Entity, mimeType openapi.MimeType) {
	if !mimeType.IsByteStream() {
		return
	}
	// For request or response bodies that are not application/json, the body
	// is modeled by a byte stream.
	entity.IsByteStream = true
	entity.IsEmpty = false
	entity.IsAny = false
}

// Construct the base request entity for a given operation. For requests whose
// mime type is not application/json, the request body is always a byte stream.
// For requests whose mime type is application/json, the request body consists
// of the top-level fields of the request object as defined in the OpenAPI spec.
// Additionally, for non-application/json requests, the request body is nested
// into a field named "contents".
func (svc *Service) newMethodEntity(op *openapi.Operation) (*Entity, openapi.MimeType, *Field) {
	if op.RequestBody == nil {
		return &Entity{fields: map[string]*Field{}}, "", nil
	}
	requestSchema, mimeType := svc.getBaseSchemaAndMimeType(op.RequestBody)
	res := svc.Package.schemaToEntity(requestSchema, []string{op.Name()}, true)
	if res == nil {
		panic(fmt.Errorf("%s request body is nil", op.OperationId))
	}

	var bodyField *Field
	if mimeType.IsByteStream() {
		bodyField = res.fields[openapi.MediaTypeNonJsonBodyFieldName]
	}

	// This next block of code is needed to make up for shortcomings in
	// schemaToEntity. That function (and the Entity structure) assumes that all
	// entities are modeled by JSON objects. Later, we should change Entity
	// and schemaToEntity to be more tolerant of non-JSON entities; for now, we
	// put this hack in place to make things work.
	if mimeType.IsByteStream() {
		for _, v := range res.fields {
			v.IsJson = false
		}
		svc.updateEntityTypeFromMimeType(bodyField.Entity, mimeType)
	}

	return res, mimeType, bodyField
}

func (svc *Service) addParams(request *Entity, op *openapi.Operation, params []openapi.Parameter) {
	for _, v := range params {
		if v.In == "header" {
			continue
		}
		param := svc.paramToField(op, v)
		if param == nil {
			continue
		}
		field, exists := request.fields[param.Name]
		if !exists {
			field = param
		}
		field.IsPath = param.IsPath
		field.IsQuery = param.IsQuery
		request.fields[param.Name] = field
		if param.Required {
			var alreadyRequired bool
			for _, v := range request.RequiredOrder {
				if v == param.Name {
					alreadyRequired = true
					break
				}
			}
			if !alreadyRequired {
				// TODO: figure out what to do with entity+param requests
				request.RequiredOrder = append(request.RequiredOrder, param.Name)
			}
		}
		if field.IsQuery {
			// recursively update field entity and sub entities with IsQuery = true
			// this should be safe as paramToField() should recursively create
			// all needed sub-entities
			field.Traverse(
				func(f *Field) {
					f.IsQuery = true
				})
		}
	}
	// IsQuery may have been set on some fields, so the request entity and
	// sub-entities need to be updated
	request.Traverse(
		func(e *Entity) {
			svc.Package.updateType(e)
		})
}

// The body param must be added after all other params so that it appears in the
// correct position in shortcut methods.
func (svc *Service) addBodyParamIfNeeded(request *Entity, mimeType openapi.MimeType) {
	if mimeType.IsByteStream() {
		request.RequiredOrder = append(request.RequiredOrder, openapi.MediaTypeNonJsonBodyFieldName)
	}
}

// Use heuristics to construct a "good" name for the request entity, as the name
// was not provided by the original OpenAPI spec.
func (svc *Service) nameAndDefineRequest(request *Entity, op *openapi.Operation) {
	if request.Name != "" {
		panic(fmt.Sprintf("request entity already has a name: %s", request.Name))
	}
	// when there was a merge of params with a request or new entity was made
	signularServiceName := svc.Singular().PascalName()
	notExplicit := !strings.Contains(op.Name(), signularServiceName)
	if op.Name() == "list" && notExplicit {
		request.Name = op.Name() + svc.Name + "Request"
	} else if crudNames[strings.ToLower(op.Name())] {
		request.Name = op.Name() + signularServiceName + "Request"
	} else {
		request.Name = op.Name() + "Request"
	}
	if svc.Package.Name == "scim" {
		request.Name = strings.ReplaceAll(request.Name, "Account", "")
	}
	request.Description = op.Summary
	svc.Package.define(request)
}

// Constructs the request object metadata for a method. This consists of
//
//  1. the request entity (i.e. the parameters and/or body)
//  2. the request MIME type
//  3. the field pointing to the request body (for non-JSON requests)
//
// The request entity includes fields for every parameter in the request (path,
// query, and body). If the request is defined anonymously (i.e. it is not
// refactored into a named type), the name for the request is constructed from
// the operation name and service name.
func (svc *Service) newRequest(params []openapi.Parameter, op *openapi.Operation) (*Entity, openapi.MimeType, *Field) {
	if op.RequestBody == nil && len(params) == 0 {
		return nil, "", nil
	}
	request, mimeType, bodyField := svc.newMethodEntity(op)
	if request.fields == nil && request.MapValue == nil {
		return nil, "", nil
	}
	svc.addParams(request, op, params)
	svc.addBodyParamIfNeeded(request, mimeType)
	if request.Name == "" {
		svc.nameAndDefineRequest(request, op)
	}
	return request, mimeType, bodyField
}

func (svc *Service) newResponse(op *openapi.Operation) (*Entity, openapi.MimeType, *Field, Named) {
	body := op.SuccessResponseBody(svc.Package.Components)
	schema, mimeType := svc.getBaseSchemaAndMimeType(body)
	name := op.Name()
	response := svc.Package.definedEntity(name+"Response", schema)
	var bodyField *Field
	if mimeType.IsByteStream() {
		bodyField = response.fields[openapi.MediaTypeNonJsonBodyFieldName]
	}

	// This next block of code is needed to make up for shortcomings in
	// schemaToEntity. That function (and the Entity structure) assumes that all
	// entities are modeled by JSON objects. Later, we should change Entity
	// and schemaToEntity to be more tolerant of non-JSON entities; for now, we
	// put this hack in place to make things work.
	if mimeType.IsByteStream() {
		svc.updateEntityTypeFromMimeType(bodyField.Entity, mimeType)
		for _, v := range response.fields {
			v.IsJson = false
		}
	}

	var emptyResponse Named
	if response != nil && response.IsEmpty {
		emptyResponse = response.Named
		response = nil
	}
	return response, mimeType, bodyField, emptyResponse
}

func (svc *Service) paramPath(path string, request *Entity, params []openapi.Parameter) (parts []PathPart) {
	var pathParams int
	for _, v := range params {
		if v.In == "path" {
			pathParams++
		}
	}
	if svc.IsAccounts && pathParams == 0 {
		// account-level services do always have `/accounts/2.0` in path
		pathParams++
	}
	if pathParams == 0 {
		return
	}
	for _, v := range pathPairRE.FindAllStringSubmatch(path, -1) {
		prefix := v[1]
		name := v[3]
		if svc.IsAccounts && name == "account_id" {
			parts = append(parts, PathPart{prefix, nil, true})
			continue
		}
		if request == nil {
			// e.g. POST /api/2.0/accounts/{account_id}/budget
			parts = append(parts, PathPart{prefix, nil, false})
			continue
		}
		field, ok := request.fields[name]
		if !ok {
			parts = append(parts, PathPart{prefix, nil, false})
			continue
		}
		parts = append(parts, PathPart{prefix, field, false})
	}
	return
}

func (svc *Service) getPathStyle(op *openapi.Operation) openapi.PathStyle {
	if op.PathStyle != "" {
		return op.PathStyle
	}
	if svc.PathStyle != "" {
		return svc.PathStyle
	}
	return openapi.PathStyleRest
}

func (svc *Service) newMethod(verb, path string, params []openapi.Parameter, op *openapi.Operation) *Method {
	methodName := op.Name()
	request, reqMimeType, reqBodyField := svc.newRequest(params, op)
	response, respMimeType, respBodyField, emptyResponse := svc.newResponse(op)
	requestStyle := svc.getPathStyle(op)
	if requestStyle == openapi.PathStyleRpc {
		methodName = filepath.Base(path)
	}
	description := op.Description
	summary := strings.TrimSpace(op.Summary)
	// merge summary into description
	if summary != "" {
		// add a dot to the end of the summary, so that it could be extracted
		// in templated with [*Named.Summary], separating from the rest of
		// description.
		if !strings.HasSuffix(summary, ".") {
			summary += "."
		}
		description = fmt.Sprintf("%s\n\n%s", summary, description)
	}

	var nameFieldPath, idFieldPath []*Field
	respEntity := getPaginationEntity(response, op.Pagination)
	if op.HasNameField() && respEntity != nil {
		nameField, err := respEntity.GetUnderlyingFields(op.NameField)
		if err != nil {
			panic(fmt.Errorf("[%s] could not find name field %q: %w", op.OperationId, op.NameField, err))
		}
		nameFieldPath = nameField
	}
	if op.HasIdentifierField() && respEntity != nil {
		idField, err := respEntity.GetUnderlyingFields(op.IdField)
		if err != nil {
			panic(fmt.Errorf("[%s] could not find id field %q: %w", op.OperationId, op.IdField, err))
		}
		idFieldPath = idField
	}
	headers := map[string]string{}
	if reqMimeType != "" {
		headers["Content-Type"] = string(reqMimeType)
	}
	if respMimeType != "" {
		headers["Accept"] = string(respMimeType)
	}
	return &Method{
		Named:               Named{methodName, description},
		Service:             svc,
		Verb:                strings.ToUpper(verb),
		Path:                path,
		Request:             request,
		PathParts:           svc.paramPath(path, request, params),
		Response:            response,
		EmptyResponseName:   emptyResponse,
		PathStyle:           requestStyle,
		NameFieldPath:       nameFieldPath,
		IdFieldPath:         idFieldPath,
		RequestBodyField:    reqBodyField,
		ResponseBodyField:   respBodyField,
		FixedRequestHeaders: headers,
		wait:                op.Wait,
		Operation:           op,
		pagination:          op.Pagination,
		shortcut:            op.Shortcut,
	}
}

func (svc *Service) HasWaits() bool {
	for _, v := range svc.methods {
		if v.wait != nil {
			return true
		}
	}
	return false
}

func (svc *Service) Waits() (waits []*Wait) {
	seen := map[string]bool{}
	for _, m := range svc.methods {
		if m.wait == nil {
			continue
		}
		wait := m.Wait()
		if seen[wait.Name] {
			continue
		}
		waits = append(waits, wait)
		seen[wait.Name] = true
	}
	slices.SortFunc(waits, func(a, b *Wait) bool {
		return a.Name < b.Name
	})
	return waits
}

// IsPrivatePreview flags object being in private preview.
func (svc *Service) IsPrivatePreview() bool {
	return isPrivatePreview(&svc.tag.Node)
}

// IsPublicPreview flags object being in public preview.
func (svc *Service) IsPublicPreview() bool {
	return isPublicPreview(&svc.tag.Node)
}

func getPaginationEntity(entity *Entity, pagination *openapi.Pagination) *Entity {
	if pagination == nil {
		return nil
	}
	if pagination.Inline {
		return entity.ArrayValue
	}
	return entity.Field(pagination.Results).Entity.ArrayValue
}
