package roll

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi/code"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"golang.org/x/exp/slices"
)

func NewSuite(dirname string) (*Suite, error) {
	fset := token.NewFileSet()
	s := &Suite{
		fset:             fset,
		ServiceToPackage: map[string]string{},
	}
	err := filepath.WalkDir(dirname, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "acceptance_test.go") {
			// not transpilable
			return nil
		}
		if strings.HasSuffix(path, "files_test.go") {
			// not transpilable
			return nil
		}
		if strings.HasSuffix(path, "workspaceconf_test.go") {
			// not transpilable
			return nil
		}
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}
		s.expectExamples(file)
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = s.parsePackages(dirname+"/../workspace_client.go", "WorkspaceClient")
	if err != nil {
		return nil, err
	}
	err = s.parsePackages(dirname+"/../account_client.go", "AccountClient")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type Suite struct {
	fset             *token.FileSet
	ServiceToPackage map[string]string
	examples         []*example
	counter          int
}

func (s *Suite) parsePackages(filename, client string) error {
	file, err := parser.ParseFile(s.fset, filename, nil, 0)
	if err != nil {
		return err
	}
	spec, ok := file.Scope.Objects[client].Decl.(*ast.TypeSpec)
	if !ok {
		s.explainAndPanic("type spec", file.Scope.Objects[client].Decl)
	}
	structType, ok := spec.Type.(*ast.StructType)
	if !ok {
		s.explainAndPanic("struct type", spec.Type)
	}
	for _, f := range structType.Fields.List {
		fieldName := f.Names[0].Name
		starExpr, ok := f.Type.(*ast.StarExpr)
		if !ok {
			continue
		}
		selectorExpr, ok := starExpr.X.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		apiName := selectorExpr.Sel.Name
		if !strings.HasSuffix(apiName, "API") {
			continue
		}
		s.ServiceToPackage[fieldName] = s.expectIdent(selectorExpr.X)
	}
	return nil
}

func (s *Suite) FullName() string {
	return "suite"
}

type methodRef struct {
	Pacakge, Service, Method string
}

func (s *Suite) Methods() []methodRef {
	found := map[methodRef]bool{}
	for _, ex := range s.examples {
		for _, v := range ex.Calls {
			if v.Service == nil {
				continue
			}
			if strings.HasSuffix(v.PascalName(), "IdMap") {
				continue
			}
			found[methodRef{
				Pacakge: s.ServiceToPackage[v.Service.PascalName()],
				Service: v.Service.CamelName(),
				Method:  v.OriginalName(),
			}] = true
		}
	}
	methods := []methodRef{}
	for k := range found {
		methods = append(methods, k)
	}
	slices.SortFunc(methods, func(a, b methodRef) bool {
		return a.Service < b.Service && a.Method < b.Method
	})
	return methods
}

func (s *Suite) Samples() (out []*sample) {
	for _, v := range s.Methods() {
		out = append(out, s.usageSamples(v.Service, v.Method)...)
	}
	return out
}

type serviceExample struct {
	code.Named
	Suite   *Suite
	Package string
	Samples []*sample
}

func (s *serviceExample) FullName() string {
	return fmt.Sprintf("%s.%s", s.Package, s.PascalName())
}

func (s *Suite) ServicesExamples() (out []*serviceExample) {
	samples := s.Samples()
	for svc, pkg := range s.ServiceToPackage {
		se := &serviceExample{
			Named: code.Named{
				Name: svc,
			},
			Package: pkg,
			Suite:   s,
		}
		for _, v := range samples {
			if v.Service.PascalName() != se.PascalName() {
				continue
			}
			se.Samples = append(se.Samples, v)
		}
		slices.SortFunc(se.Samples, func(a, b *sample) bool {
			return a.FullName() < b.FullName()
		})
		out = append(out, se)
	}
	return out
}

type sample struct {
	example
	Package string
	Service *code.Named
	Method  *code.Named
	Suite   *Suite
}

func (sa *sample) FullName() string {
	return fmt.Sprintf("%s.%s", sa.Service.PascalName(), sa.Method.CamelName())
}

func (s *Suite) usageSamples(svc, mth string) (out []*sample) {
	svcName := &code.Named{
		Name: svc,
	}
	methodName := &code.Named{
		Name: mth,
	}
	for _, ex := range s.examples {
		c := ex.findCall(svcName.CamelName(), methodName.CamelName())
		if c == nil {
			continue
		}
		sa := &sample{
			example: example{
				Named:     ex.Named,
				IsAccount: ex.IsAccount,
			},
			Service: svcName,
			Method:  methodName,
			Package: s.ServiceToPackage[svcName.PascalName()],
			Suite:   s,
		}
		out = append(out, sa)
		variablesUsed := []string{}
		queue := []expression{c}
		added := map[string]bool{}
		callIds := map[int]bool{}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			switch x := current.(type) {
			case *call:
				if callIds[x.id] {
					continue
				}
				if x.Assign != nil && x.Assign.Name != "_" {
					variablesUsed = append(variablesUsed, x.Assign.CamelName())
					// call methods that may actually create an entity.
					// executed before we append to sa.Calls, as we reverse
					// the slice in the end
					for _, v := range ex.Calls {
						if v.creates != x.Assign.CamelName() {
							continue
						}
						if callIds[v.id] {
							continue
						}
						// put at the front of the queue
						queue = append([]expression{v}, queue...)
					}
				}
				x.IsAccount = ex.IsAccount
				sa.Calls = append(sa.Calls, x)
				callIds[x.id] = true
				x.Traverse(func(e expression) {
					v, ok := e.(*variable)
					if !ok {
						return
					}
					if added[v.CamelName()] {
						// don't add the same variable twice
						return
					}
					found, ok := ex.scope[v.CamelName()]
					if ok {
						queue = append(queue, found)
						added[v.CamelName()] = true
						return
					}
					for _, iv := range ex.Init {
						if iv.CamelName() != v.CamelName() {
							continue
						}
						sa.Init = append(sa.Init, iv)
						return
					}
				})
			default:
				panic("unsupported")
			}
		}
		for _, v := range variablesUsed {
			// TODO: also include ex.Asserts
			for _, c := range ex.Cleanup {
				if !c.HasVariable(v) {
					continue
				}
				if callIds[c.id] {
					continue
				}
				c.Traverse(func(e expression) {
					v, ok := e.(*variable)
					if !ok {
						return
					}
					if added[v.CamelName()] {
						// don't add the same variable twice
						return
					}
					found, ok := ex.scope[v.CamelName()]
					if !ok {
						return
					}
					assignCall, ok := found.(*call)
					if !ok {
						// ideally, we could do multiple optimization passes
						// to discover used variables also in cleanups, but
						// we're not doing it for now.
						return
					}
					if callIds[assignCall.id] {
						return
					}
					sa.Calls = append(sa.Calls, assignCall)
					callIds[assignCall.id] = true
				})
				c.IsAccount = ex.IsAccount
				sa.Cleanup = append(sa.Cleanup, c)
				callIds[c.id] = true
			}
		}
		slices.SortFunc[*call](sa.Calls, func(a, b *call) bool {
			x := a.IsDependentOn(b)
			return x
		})
		reverse(sa.Calls)
		reverse(sa.Cleanup)
	}
	return out
}

func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (s *Suite) assignedNames(a *ast.AssignStmt) (names []string) {
	for _, v := range a.Lhs {
		ident, ok := v.(*ast.Ident)
		if !ok {
			continue
		}
		names = append(names, ident.Name)
	}
	return
}

func (s *Suite) expectExamples(file *ast.File) {
	for _, v := range file.Decls {
		fn, ok := v.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fnName := fn.Name.Name
		if !strings.HasPrefix(fnName, "TestAcc") &&
			!strings.HasPrefix(fnName, "TestMws") &&
			!strings.HasPrefix(fnName, "TestUc") {
			continue
		}
		if strings.HasSuffix(fnName, "NoTranspile") {
			// Tests with NoTranspile suffix are too expensive to automatically
			// translate, we just ignore them.
			continue
		}
		s.examples = append(s.examples, s.expectFn(fn, file))
	}
}

func (s *Suite) expectFn(fn *ast.FuncDecl, file *ast.File) *example {
	testName := fn.Name.Name
	testName = strings.TrimPrefix(testName, "TestAcc")
	testName = strings.TrimPrefix(testName, "TestUcAcc")
	testName = strings.TrimPrefix(testName, "TestMwsAcc")
	ex := &example{
		Named: code.Named{
			Name: testName,
		},
		scope: map[string]expression{},
	}
	lastPos := fn.Pos()
	hint := ""
	for _, v := range fn.Body.List {
		for _, cmt := range file.Comments {
			if cmt.End() < lastPos {
				continue
			}
			if cmt.Pos() > v.Pos() {
				continue
			}
			// figure out comment hint exactly above the given statement
			hint = strings.TrimSpace(cmt.Text())
			break
		}
		switch stmt := v.(type) {
		case *ast.AssignStmt:
			s.expectAssign(ex, stmt, hint)
		case *ast.DeferStmt:
			ex.Cleanup = append(ex.Cleanup, s.expectCall(stmt.Call))
		case *ast.ExprStmt:
			s.expectExprStmt(ex, stmt)
		}
		// store the end of the last statement to figure out the next
		// statement comment.
		lastPos = v.End()
		hint = ""
	}
	return ex
}

var ignoreFns = map[string]bool{
	"SkipNow": true,
	"NoError": true,
	"Lock":    true,
	"Errorf":  true,
	"Skipf":   true,
	"Log":     true,
}

func (s *Suite) expectCleanup(ex *example, ce *ast.CallExpr) bool {
	se, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok || se.Sel.Name != "Cleanup" {
		return false
	}
	if len(ce.Args) == 0 {
		return false
	}
	inlineFn, ok := ce.Args[0].(*ast.FuncLit)
	if !ok {
		return false
	}
	for _, v := range inlineFn.Body.List {
		assign, ok := v.(*ast.AssignStmt)
		if !ok {
			continue
		}
		c := s.expectAssignCall(assign)
		if c == nil {
			continue
		}
		ex.Cleanup = append(ex.Cleanup, c)
	}
	return true
}

func (s *Suite) expectExprStmt(ex *example, stmt *ast.ExprStmt) {
	ce, ok := stmt.X.(*ast.CallExpr)
	if !ok {
		return
	}
	if s.expectCleanup(ex, ce) {
		return
	}
	c := s.expectCall(stmt.X)
	if ignoreFns[c.Name] {
		return
	}
	assertions := map[string]string{
		"EqualError":     "equalError",
		"Contains":       "contains",
		"GreaterOrEqual": ">=",
		"Greater":        ">",
		"Equal":          "==",
		"NotEqual":       "!=",
	}
	op, ok := assertions[c.Name]
	if ok {
		ex.Asserts = append(ex.Asserts, &binaryExpr{
			Left:  c.Args[0],
			Op:    op,
			Right: c.Args[1],
		})
	} else if c.Name == "NotEmpty" {
		// TODO: replace code occurences with assert.True
		ex.Asserts = append(ex.Asserts, &binaryExpr{
			Left:  c.Args[0],
			Op:    "notEmpty",
			Right: nil,
		})
	} else if c.Name == "True" {
		ex.Asserts = append(ex.Asserts, c.Args[0])
	} else {
		s.explainAndPanic("known assertion", c)
	}
}

func (s *Suite) expectAssign(ex *example, stmt *ast.AssignStmt, hint string) {
	names := s.assignedNames(stmt)
	if len(names) == 2 && names[0] == "ctx" {
		// w - workspace, a - account
		ex.IsAccount = names[1] == "a"
		return
	}
	switch x := stmt.Rhs[0].(type) {
	case *ast.CallExpr:
		c := s.expectAssignCall(stmt)
		if c == nil {
			return
		}
		if c.Assign != nil && c.Assign.Name != "_" {
			ex.scope[c.Assign.CamelName()] = c
		}
		if strings.HasPrefix(hint, "creates ") {
			c.creates = strings.TrimPrefix(hint, "creates ")
		}
		ex.Calls = append(ex.Calls, c)
	case *ast.BasicLit:
		lit := s.expectPrimitive(x)
		ex.Init = append(ex.Init, &initVar{
			Named: code.Named{
				Name: names[0],
			},
			Value: lit,
		})
	}
}

func (s *Suite) expectAssignCall(stmt *ast.AssignStmt) *call {
	names := s.assignedNames(stmt)
	c := s.expectCall(stmt.Rhs[0])
	if len(names) == 1 && names[0] != "err" {
		c.Assign = &code.Named{
			Name: names[0],
		}
	}
	if len(names) == 2 && names[1] == "err" {
		c.Assign = &code.Named{
			Name: names[0],
		}
	}
	return c
}

func (s *Suite) expectIdent(e ast.Expr) string {
	ident, ok := e.(*ast.Ident)
	if !ok {
		s.explainAndPanic("ident", e)
		return ""
	}
	return ident.Name
}

func (s *Suite) expectEntity(t *ast.SelectorExpr, cl *ast.CompositeLit) *entity {
	ent := &entity{}
	ent.Name = t.Sel.Name
	ent.Package = s.expectIdent(t.X)
	ent.FieldValues = s.expectFieldValues(cl.Elts)
	return ent
}

func (s *Suite) expectFieldValues(exprs []ast.Expr) (fvs []*fieldValue) {
	for _, v := range exprs {
		switch kv := v.(type) {
		case *ast.KeyValueExpr:
			fvs = append(fvs, &fieldValue{
				Named: code.Named{
					Name: s.expectIdent(kv.Key),
				},
				Value: s.expectExpr(kv.Value),
			})
		default:
			s.explainAndPanic("field value", v)
			return
		}
	}
	return fvs
}

func (s *Suite) expectArray(t *ast.ArrayType, cl *ast.CompositeLit) *array {
	arr := &array{}
	switch elt := t.Elt.(type) {
	case *ast.SelectorExpr:
		arr.Package = s.expectIdent(elt.X)
		arr.Name = s.expectIdent(elt.Sel)
	case *ast.Ident:
		arr.Name = s.expectIdent(elt)
	default:
		s.explainAndPanic("array element", elt)
		return nil
	}
	for _, v := range cl.Elts {
		switch item := v.(type) {
		case *ast.CompositeLit:
			arr.Values = append(arr.Values, &entity{
				Named: code.Named{
					Name: arr.Name,
				},
				Package:     arr.Package,
				FieldValues: s.expectFieldValues(item.Elts),
			})
		default:
			arr.Values = append(arr.Values, s.expectExpr(v))
		}
	}
	return arr
}

func (s *Suite) expectMap(t *ast.MapType, cl *ast.CompositeLit) *mapLiteral {
	m := &mapLiteral{
		KeyType:   s.expectIdent(t.Key),
		ValueType: s.expectIdent(t.Value),
	}
	for _, v := range cl.Elts {
		kv, ok := v.(*ast.KeyValueExpr)
		if !ok {
			s.explainAndPanic("key value expr", v)
		}
		m.Pairs = append(m.Pairs, mapKV{
			Key:   s.expectExpr(kv.Key),
			Value: s.expectExpr(kv.Value),
		})
	}
	return m
}

func (s *Suite) expectPrimitive(x *ast.BasicLit) expression {
	// we directly translate literal values
	if x.Value[0] == '`' {
		txt := x.Value[1 : len(x.Value)-1]
		return &heredoc{
			Value: compute.TrimLeadingWhitespace(txt),
		}
	}
	return &literal{x.Value}
}

func (s *Suite) expectCompositeLiteral(x *ast.CompositeLit) expression {
	switch t := x.Type.(type) {
	case *ast.SelectorExpr:
		return s.expectEntity(t, x)
	case *ast.ArrayType:
		return s.expectArray(t, x)
	case *ast.MapType:
		return s.expectMap(t, x)
	default:
		s.explainAndPanic("composite lit type", t)
		return nil
	}
}

func (s *Suite) expectExpr(e ast.Expr) expression {
	switch x := e.(type) {
	case *ast.BasicLit:
		return s.expectPrimitive(x)
	case *ast.CompositeLit:
		return s.expectCompositeLiteral(x)
	case *ast.UnaryExpr:
		if x.Op != token.AND {
			s.explainAndPanic("only references to composite literals are supported", x)
		}
		y, ok := x.X.(*ast.CompositeLit)
		if !ok {
			s.explainAndPanic("composite lit", x.X)
		}
		e, ok := s.expectCompositeLiteral(y).(*entity)
		if !ok {
			s.explainAndPanic("entity", x)
		}
		e.IsPointer = true
		return e
	case *ast.SelectorExpr:
		return s.expectLookup(x)
	case *ast.Ident:
		return &variable{
			Named: code.Named{
				Name: x.Name,
			},
		}
	case *ast.CallExpr:
		return s.expectCall(x)
	case *ast.BinaryExpr:
		return &binaryExpr{
			Left:  s.expectExpr(x.X),
			Right: s.expectExpr(x.Y),
			Op:    x.Op.String(),
		}
	case *ast.IndexExpr:
		return &indexExpr{
			Left:  s.expectExpr(x.X),
			Right: s.expectExpr(x.Index),
		}
	default:
		s.explainAndPanic("expr", e)
		return nil
	}
}

func (s *Suite) explainAndPanic(expected string, x any) any {
	ast.Print(s.fset, x)
	panic("expected " + expected)
}

func (s *Suite) expectLookup(se *ast.SelectorExpr) *lookup {
	return &lookup{
		X: s.expectExpr(se.X),
		Field: &code.Named{
			Name: s.expectIdent(se.Sel),
		},
	}
}

func (s *Suite) expectCall(e ast.Expr) *call {
	ce, ok := e.(*ast.CallExpr)
	if !ok {
		s.explainAndPanic("call expr", e)
		return nil
	}
	s.counter++
	c := &call{
		id: s.counter,
	}
	switch t := ce.Fun.(type) {
	case *ast.Ident:
		c.Name = t.Name
	case *ast.SelectorExpr:
		c.Name = t.Sel.Name
		switch se := t.X.(type) {
		case *ast.SelectorExpr:
			c.Service = &code.Named{
				Name: se.Sel.Name,
			}
		case *ast.Ident:
			c.Service = &code.Named{
				Name: se.Name,
			}
		default:
			s.explainAndPanic("selector expr", se)
			return nil
		}
	}
	if c.Service != nil && c.Service.Name == "fmt" {
		// fmt.Sprintf is not a service
		c.Service = nil
	}
	for _, v := range ce.Args {
		arg := s.expectExpr(v)
		if x, ok := arg.(*variable); ok && (x.Name == "ctx" || x.Name == "t") {
			// context.Context is irrelevant in other languages than Go
			continue
		}
		c.Args = append(c.Args, arg)
	}
	return c
}
