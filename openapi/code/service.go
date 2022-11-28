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
	IsRpcStyle          bool
	IsAccounts          bool
	Package             *Package
	methods             map[string]*Method
	ByPathParamsMethods []*Shortcut
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
		signularServiceName := svc.Singular().PascalName()
		notExplicit := !strings.Contains(op.Name(), signularServiceName)
		if svc.MatchesPackageName() {
			request.Name = op.Name()
		} else if op.Name() == "list" && notExplicit {
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
	return request
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
		parts = append(parts, PathPart{prefix, &field, false})
	}
	return
}

func (svc *Service) newMethod(verb, path string, params []openapi.Parameter, op *openapi.Operation) *Method {
	request := svc.newRequest(params, op)
	respSchema := op.SuccessResponseSchema(svc.Package.Components)
	name := op.Name()
	response := svc.Package.definedEntity(name+"Response", respSchema)
	var emptyResponse Named
	if response != nil && response.IsEmpty {
		emptyResponse = response.Named
		response = nil
	}
	if svc.IsRpcStyle {
		name = filepath.Base(path)
	}
	description := op.Description
	if op.Summary != "" {
		description = fmt.Sprintf("%s\n\n%s", op.Summary, description)
	}
	return &Method{
		Named:             Named{name, description},
		Service:           svc,
		Verb:              strings.ToUpper(verb),
		Path:              path,
		Request:           request,
		PathParts:         svc.paramPath(path, request, params),
		Response:          response,
		EmptyResponseName: emptyResponse,
		wait:              op.Wait,
		operation:         op,
		pagination:        op.Pagination,
		shortcut:          op.Shortcut,
	}
}
