package beep

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi/code"
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

func (s *suite) ImprotedServices() (svcs []string) {
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
	ex := &example{
		Named: code.Named{
			Name: fn.Name.Name,
		},
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
	c := s.expectAssignCall(stmt)
	if c == nil {
		return
	}
	ex.Calls = append(ex.Calls, c)
}

func (s *suite) expectAssignCall(stmt *ast.AssignStmt) *call {
	names := s.assignedNames(stmt)
	if len(names) == 1 && names[0] == "err" {
		return s.expectCall(stmt.Rhs[0])
	}
	if len(names) == 2 && names[1] == "err" {
		c := s.expectCall(stmt.Rhs[0])
		c.Assign = &code.Named{
			Name: names[0],
		}
		return c
	}
	return nil
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

func (s *suite) expectExpr(e ast.Expr) expression {
	switch x := e.(type) {
	case *ast.BasicLit:
		// we directly translate literal values
		return &literal{x.Value}
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
		return s.expectSimpleCall(x)
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
		X:     s.expectExpr(se.X),
		Field: s.expectExpr(se.Sel),
	}
}

func (s *suite) expectSimpleCall(ce *ast.CallExpr) *call {
	ident, ok := ce.Fun.(*ast.Ident)
	if !ok {
		s.explainAndPanic("ident", ce.Fun)
		return nil
	}
	c := &call{
		Named: code.Named{
			Name: ident.Name,
		},
	}
	for _, v := range ce.Args {
		arg := s.expectExpr(v)
		c.Args = append(c.Args, arg)
	}
	return c
}

func (s *suite) expectCall(e ast.Expr) *call {
	c := &call{}
	find(e, func(ce *ast.CallExpr) {
		// parse path to function to service/method name
		visit(ce.Fun, func(se *ast.SelectorExpr) ast.Visitor {
			c.Name = se.Sel.Name
			return visit(se.X, func(se *ast.SelectorExpr) ast.Visitor {
				c.Service = &code.Named{
					Name: se.Sel.Name,
				}
				// we don't parse X: *ast.Ident as we know the w/a context,
				// though we may add it later
				return nil
			})
		})
		for _, v := range ce.Args {
			arg := s.expectExpr(v)
			if x, ok := arg.(*variable); ok && (x.Name == "ctx" || x.Name == "t") {
				// context.Context is irrelevant in other languages than Go
				continue
			}
			c.Args = append(c.Args, arg)
		}
	})
	return c
}
