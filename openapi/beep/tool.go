package beep

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi/code"
	"golang.org/x/exp/slices"
)

func NewSuite(filename string) (*suite, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, err
	}
	s := &suite{
		fset: fset,
		file: file,
	}
	s.expectExamples()
	return s, nil
}

type suite struct {
	fset     *token.FileSet
	file     *ast.File
	examples []*example
}

func (s *suite) FullName() string {
	return "xxx"
}

type methodRef struct {
	service, method string
}

func (s *suite) methods() []methodRef {
	found := map[methodRef]bool{}
	for _, ex := range s.examples {
		for _, v := range ex.Calls {
			if v.Service == nil {
				continue
			}
			found[methodRef{
				service: v.Service.CamelName(),
				method:  v.OriginalName(),
			}] = true
		}
	}
	methods := []methodRef{}
	for k := range found {
		methods = append(methods, k)
	}
	slices.SortFunc(methods, func(a, b methodRef) bool {
		return a.service < b.service && a.method < b.method
	})
	return methods
}

func (s *suite) Samples() (out []*sample) {
	for _, v := range s.methods() {
		out = append(out, s.usageSamples(v.service, v.method)...)
	}
	return out
}

type sample struct {
	example
	Service *code.Named
	Method  *code.Named
}

func (sa *sample) FullName() string {
	return fmt.Sprintf("%s.%s", sa.Service.PascalName(), sa.Method.CamelName())
}

func (s *suite) usageSamples(svc, mth string) (out []*sample) {
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
				Named: ex.Named,
			},
			Service: svcName,
			Method:  methodName,
		}
		out = append(out, sa)
		variablesUsed := []string{}
		queue := []expression{c}
		added := map[string]bool{}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			switch x := current.(type) {
			case *call:
				sa.Calls = append(sa.Calls, x)
				if x.Assign != nil && x.Assign.Name != "_" {
					variablesUsed = append(variablesUsed, x.Assign.CamelName())
				}
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
		for i, j := 0, len(sa.Calls)-1; i < j; i, j = i+1, j-1 {
			sa.Calls[i], sa.Calls[j] = sa.Calls[j], sa.Calls[i]
		}
		for _, v := range variablesUsed {
			// for _, c := range ex.Asserts {
			// 	if !c.HasVariable(v) {
			// 		continue
			// 	}
			// 	sample.Cleanup = append(sample.Cleanup, c)
			// }
			for _, c := range ex.Cleanup {
				if !c.HasVariable(v) {
					continue
				}
				sa.Cleanup = append(sa.Cleanup, c)
			}
		}
	}
	return out
}

func (s *suite) ImprotedPackages() (svcs []string) {
	ast.Inspect(s.file, func(raw ast.Node) bool {
		switch n := raw.(type) {
		case *ast.ImportSpec:
			name := s.expectString(n.Path)
			prefix := "github.com/databricks/databricks-sdk-go/service/"
			if !strings.HasPrefix(name, prefix) {
				return true
			}
			svcs = append(svcs, strings.TrimPrefix(name, prefix))
		case *ast.FuncDecl:
			// save cycles
			return false
		}
		return true
	})
	return svcs
}

func (s *suite) assignedNames(a *ast.AssignStmt) (names []string) {
	for _, v := range a.Lhs {
		ident, ok := v.(*ast.Ident)
		if !ok {
			continue
		}
		names = append(names, ident.Name)
	}
	return
}

func (s *suite) expectString(l *ast.BasicLit) string {
	if l.Kind != token.STRING {
		return ""
	}
	return strings.Trim(l.Value, `"`)
}

func (s *suite) expectExamples() {
	for _, v := range s.file.Decls {
		fn, ok := v.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if strings.HasSuffix(fn.Name.Name, "NoTranspile") {
			// Tests with NoTranspile suffix are too expensive to automatically
			// translate, we just ignore them.
			continue
		}
		s.examples = append(s.examples, s.expectFn(fn))
	}
}

func (s *suite) expectFn(fn *ast.FuncDecl) *example {
	testName := fn.Name.Name
	testName = strings.TrimPrefix(testName, "TestAcc")
	ex := &example{
		Named: code.Named{
			Name: testName,
		},
		scope: map[string]expression{},
	}
	for _, v := range fn.Body.List {
		switch stmt := v.(type) {
		case *ast.AssignStmt:
			s.expectAssign(ex, stmt)
		case *ast.DeferStmt:
			ex.Cleanup = append(ex.Cleanup, s.expectCall(stmt.Call))
		case *ast.ExprStmt:
			s.expectExprStmt(ex, stmt)
		}
		// we ignore test skips for different clouds
	}
	return ex
}

var ignoreFns = map[string]bool{
	"SkipNow": true,
	"NoError": true,
	"Lock":    true,
	"Errorf":  true,
	"Skipf":   true,
}

func (s *suite) expectCleanup(ex *example, ce *ast.CallExpr) bool {
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

func (s *suite) expectExprStmt(ex *example, stmt *ast.ExprStmt) {
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
	switch c.Name {
	case "Equal":
		ex.Asserts = append(ex.Asserts, &binaryExpr{
			Left:  c.Args[0],
			Op:    "==",
			Right: c.Args[1],
		})
	case "NotEqual":
		ex.Asserts = append(ex.Asserts, &binaryExpr{
			Left:  c.Args[0],
			Op:    "!=",
			Right: c.Args[1],
		})
	case "True":
		ex.Asserts = append(ex.Asserts, &binaryExpr{
			Left:  c.Args[0],
			Op:    "==",
			Right: &boolean{true},
		})
	default:
		print(c)
	}
}

func (s *suite) expectAssign(ex *example, stmt *ast.AssignStmt) {
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

func (s *suite) expectAssignCall(stmt *ast.AssignStmt) *call {
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

func (s *suite) expectIdent(e ast.Expr) string {
	ident, ok := e.(*ast.Ident)
	if !ok {
		s.explainAndPanic("ident", e)
		return ""
	}
	return ident.Name
}

func (s *suite) expectEntity(t *ast.SelectorExpr, cl *ast.CompositeLit) *entity {
	ent := &entity{}
	ent.Name = t.Sel.Name
	ent.Package = s.expectIdent(t.X)
	ent.FieldValues = s.expectFieldValues(cl.Elts)
	return ent
}

func (s *suite) expectFieldValues(exprs []ast.Expr) (fvs []*fieldValue) {
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

func (s *suite) expectArray(t *ast.ArrayType, cl *ast.CompositeLit) *array {
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

func (s *suite) expectPrimitive(x *ast.BasicLit) *literal {
	// we directly translate literal values
	return &literal{x.Value}
}

func (s *suite) expectExpr(e ast.Expr) expression {
	switch x := e.(type) {
	case *ast.BasicLit:
		return s.expectPrimitive(x)
	case *ast.CompositeLit:
		switch t := x.Type.(type) {
		case *ast.SelectorExpr:
			return s.expectEntity(t, x)
		case *ast.ArrayType:
			return s.expectArray(t, x)
		default:
			s.explainAndPanic("composite lit type", t)
			return nil
		}
	case *ast.UnaryExpr:
		if x.Op != token.AND {
			panic("only references to composite literals are supported")
		}
		return s.expectExpr(x.X)
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

func (s *suite) explainAndPanic(expected string, x any) any {
	ast.Print(s.fset, x)
	panic("expected " + expected)
}

func (s *suite) expectLookup(se *ast.SelectorExpr) *lookup {
	return &lookup{
		X: s.expectExpr(se.X),
		Field: &code.Named{
			Name: s.expectIdent(se.Sel),
		},
	}
}

func (s *suite) expectCall(e ast.Expr) *call {
	ce, ok := e.(*ast.CallExpr)
	if !ok {
		s.explainAndPanic("call expr", e)
		return nil
	}
	c := &call{}
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
