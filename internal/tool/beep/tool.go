package beep

import (
	"fmt"
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

var ignoreFns = map[string]bool{
	"SkipNow": true,
	"NoError": true,
}

func (s *suite) expectExamples() {
	find(s.file, func(fn *ast.FuncDecl) {
		ex := &example{
			Named: code.Named{
				Name: fn.Name.Name,
			},
		}
		s.examples = append(s.examples, ex)
		find(fn, func(as *ast.AssignStmt) {
			names := s.assignedNames(as)
			if len(names) == 2 && names[0] == "ctx" {
				// w - workspace, a - account
				ex.IsAccount = names[1] == "a"
			}
			if len(names) == 1 && names[0] == "err" {
				ex.Calls = append(ex.Calls, s.expectCall(as.Rhs[0]))
			}
			if len(names) == 2 && names[1] == "err" {
				c := s.expectCall(as.Rhs[0])
				c.Assign = &code.Named{
					Name: names[0],
				}
				ex.Calls = append(ex.Calls, c)
			}
		})
		find(fn, func(dfr *ast.DeferStmt) {
			ex.Cleanup = append(ex.Cleanup, s.expectCall(dfr.Call))
		})
		find(fn, func(es *ast.ExprStmt) {
			c := s.expectCall(es.X)
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
		})
		// we ignore test skips for different clouds
	})
}

func (s *suite) expectIdent(e ast.Expr) string {
	ident, ok := e.(*ast.Ident)
	if !ok {
		panic(fmt.Sprintf("%v is not an ident", e))
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
			panic(fmt.Sprintf("unknown field value expr: %v", v))
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
		panic(fmt.Sprintf("unknown array element type: %v", elt))
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
			panic(fmt.Sprintf("unknown composite lit type: %v", t))
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
		panic(fmt.Sprintf("unknown expr: %v", e))
	}
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
		panic(fmt.Sprintf("expected simple call to have *ast.Ident for name, got %v", ce.Fun))
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
