package beep

import "github.com/databricks/databricks-sdk-go/openapi/code"

type expression interface {
	Type() string
}

type binaryExpr struct {
	Left, Right expression
	Op          string
}

func (b *binaryExpr) Type() string {
	return "binary"
}

type indexExpr struct {
	Left, Right expression
}

func (i *indexExpr) Type() string {
	return "index"
}

type boolean struct {
	Value bool
}

func (b *boolean) Type() string {
	return "boolean"
}

type literal struct {
	Value string
}

func (l *literal) Type() string {
	return "literal"
}

type lookup struct {
	X     expression
	Field expression
}

func (l *lookup) Type() string {
	return "lookup"
}

type variable struct {
	code.Named
}

func (v *variable) Type() string {
	return "variable"
}

type entity struct {
	code.Named
	Package     string
	FieldValues []*fieldValue
}

func (e *entity) Type() string {
	return "entity"
}

type array struct {
	code.Named
	Package string
	Values  []expression
}

func (a *array) Type() string {
	return "array"
}

type fieldValue struct {
	code.Named
	Value expression
}

type example struct {
	code.Named
	IsAccount bool
	Calls     []*call
	Cleanup   []*call
	Asserts   []*binaryExpr
	scope     map[string]expression
}

type call struct {
	code.Named
	Service *code.Named
	Assign  *code.Named
	Args    []expression
}

func (c *call) Type() string {
	return "call"
}
