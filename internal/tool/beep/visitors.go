package beep

import "go/ast"

type walker func(ast.Node) ast.Visitor

func (f walker) Visit(node ast.Node) ast.Visitor {
	return f(node)
}

func visit[X any](n ast.Node, visitor func(X) ast.Visitor) ast.Visitor {
	x, ok := n.(X)
	if !ok {
		return walker(func(n ast.Node) ast.Visitor {
			return visit(n, visitor)
		})
	}
	return visitor(x)
}

func find[X any](n ast.Node, visitor func(X)) {
	ast.Walk(walker(func(n ast.Node) ast.Visitor {
		return visit(n, func(x X) ast.Visitor {
			visitor(x)
			return nil
		})
	}), n)
}