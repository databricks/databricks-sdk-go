package code

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/openapi"
	"golang.org/x/exp/slices"
)

type Service struct {
	Named
	IsRpcStyle          bool
	Package             *Package
	methods             map[string]*Method
	ByPathParamsMethods []*Shortcut
}

func (svc *Service) FullName() string {
	return fmt.Sprintf("%s.%s", svc.Name, svc.Package.Name)
}

func (svc *Service) Methods() (methods []*Method) {
	for _, v := range svc.methods {
		methods = append(methods, v)
	}
	slices.SortFunc(methods, func(a, b *Method) bool {
		return a.Name < b.Name
	})
	return methods
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

func (svc *Service) newRequest(params []openapi.Parameter, op *openapi.Operation) *Entity {
	if op.RequestBody == nil && len(params) == 0 {
		return nil
	}
	request := &Entity{fields: map[string]Field{}}
	if op.RequestBody != nil {
		request = svc.Package.schemaToEntity(op.RequestBody.Schema(), []string{op.Name()}, true)
	}
	if request == nil {
		panic(fmt.Errorf("%s request body is nil", op.OperationId))
	}
	if request.fields == nil && request.MapValue == nil {
		panic(fmt.Errorf("%s request schema has no fields", op.OperationId))
	}
	for _, v := range params {
		param := svc.paramToField(op, v)
		if param == nil {
			continue
		}
		field, exists := request.fields[param.Name]
		if !exists {
			field = *param
		}
		field.IsPath = param.IsPath
		field.IsQuery = param.IsQuery
		request.fields[param.Name] = field
	}
	if request.Name == "" {
		// when there was a merge of params with a request or new entity was made
		request.Name = op.Name() + "Request"
		svc.Package.define(request)
	}
	return request
}

func (svc *Service) paramPath(path string, request *Entity, params []openapi.Parameter) (parts []PathPart) {
	var pathParams int
	for _, v := range params {
		if v.In == "path" {
			pathParams++
		}
	}
	if pathParams == 0 {
		return
	}
	for _, v := range pathPairRE.FindAllStringSubmatch(path, -1) {
		field, ok := request.fields[v[3]]
		if !ok {
			parts = append(parts, PathPart{v[1], nil})
			continue
		}
		parts = append(parts, PathPart{v[1], &field})
	}
	return
}

func (svc *Service) newMethod(verb, path string, params []openapi.Parameter, op *openapi.Operation) *Method {
	request := svc.newRequest(params, op)
	respSchema := op.SuccessResponseSchema(svc.Package.Components)
	name := op.Name()
	response := svc.Package.definedEntity(name+"Response", respSchema)
	if svc.IsRpcStyle {
		name = filepath.Base(path)
	}
	return &Method{
		Named:      Named{name, op.Description},
		Service:    svc,
		Verb:       strings.ToUpper(verb),
		Path:       path,
		Request:    request,
		PathParts:  svc.paramPath(path, request, params),
		Response:   response,
		wait:       op.Wait,
		operation:  op,
		pagination: op.Pagination,
		shortcut:   op.Shortcut,
	}
}
