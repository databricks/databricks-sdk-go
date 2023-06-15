package roll

import (
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi/code"
)

type expression interface {
	Type() string
}

type traversable interface {
	Traverse(func(expression))
}

type binaryExpr struct {
	Left, Right expression
	Op          string
}

func (b *binaryExpr) Traverse(cb func(expression)) {
	cb(b.Left)
	if t, ok := b.Left.(traversable); ok {
		t.Traverse(cb)
	}
	cb(b.Right)
	if t, ok := b.Right.(traversable); ok {
		t.Traverse(cb)
	}
}

func (b *binaryExpr) Type() string {
	return "binary"
}

type indexExpr struct {
	Left, Right expression
}

func (i *indexExpr) Traverse(cb func(expression)) {
	cb(i.Left)
	if t, ok := i.Left.(traversable); ok {
		t.Traverse(cb)
	}
	cb(i.Right)
	if t, ok := i.Right.(traversable); ok {
		t.Traverse(cb)
	}
}

func (i *indexExpr) Type() string {
	return "index"
}

// type boolean struct {
// 	Value bool
// }

// func (b *boolean) Type() string {
// 	return "boolean"
// }

type literal struct {
	Value string
}

func (l *literal) Type() string {
	return "literal"
}

type heredoc struct {
	Value string
}

func (l *heredoc) Type() string {
	return "heredoc"
}

type lookup struct {
	X     expression
	Field *code.Named
}

func (l *lookup) Variable() string {
	switch x := l.X.(type) {
	case *variable:
		return x.Name
	default:
		return ""
	}
}

func (l *lookup) Traverse(cb func(expression)) {
	cb(l.X)
	if t, ok := l.X.(traversable); ok {
		t.Traverse(cb)
	}
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
	IsPointer   bool
}

func (e *entity) Traverse(cb func(expression)) {
	for _, v := range e.FieldValues {
		cb(v.Value)
		if t, ok := v.Value.(traversable); ok {
			t.Traverse(cb)
		}
	}
}

func (e *entity) Type() string {
	return "entity"
}

type mapKV struct {
	Key   expression
	Value expression
}

type mapLiteral struct {
	KeyType   string
	ValueType string
	Pairs     []mapKV
}

func (e *mapLiteral) Type() string {
	return "map"
}

func (e *mapLiteral) Traverse(cb func(expression)) {
	for _, v := range e.Pairs {
		if t, ok := v.Key.(traversable); ok {
			t.Traverse(cb)
		}
		if t, ok := v.Value.(traversable); ok {
			t.Traverse(cb)
		}
	}
}

type array struct {
	code.Named
	Package string
	Values  []expression
}

func (a *array) Traverse(cb func(expression)) {
	for _, v := range a.Values {
		cb(v)
		if t, ok := v.(traversable); ok {
			t.Traverse(cb)
		}
	}
}

func (a *array) Type() string {
	return "array"
}

type fieldValue struct {
	code.Named
	Value expression
}

type call struct {
	code.Named
	IsAccount bool
	Service   *code.Named
	Assign    *code.Named
	Args      []expression

	// ID to avoid duplicates. alternative could be hashing,
	// but implementation would grow more complex than needed.
	id int

	// hint about the call creating an entity behind the variable
	creates string
}

func (c *call) IsDependentOn(other *call) bool {
	if other.Assign == nil {
		return false
	}
	result := []bool{false}
	c.Traverse(func(e expression) {
		v, ok := e.(*variable)
		if !ok {
			return
		}
		if other.Assign.CamelName() == v.CamelName() {
			result[0] = true
			return
		}
	})
	return result[0]
}

func (c *call) IsWait() bool {
	return strings.HasSuffix(c.Name, "AndWait")
}

func (c *call) Request() (fv []*fieldValue) {
	if strings.Contains(c.Name, "By") {
		fields := strings.Split(strings.Split(c.Name, "By")[0], "And")
		for i, name := range fields {
			fv = append(fv, &fieldValue{
				Named: code.Named{
					Name: name,
				},
				Value: c.Args[i],
			})
		}
		return fv
	}
	if len(c.Args) == 0 {
		return fv
	}
	e, ok := c.Args[0].(*entity)
	if !ok {
		return fv
	}
	return e.FieldValues
}

func (c *call) Original() *code.Named {
	name := c.CamelName()
	name = strings.Split(name, "By")[0]
	name = strings.TrimSuffix(name, "AndWait")
	name = strings.TrimSuffix(name, "All")
	return &code.Named{
		Name: name,
	}
}

func (c *call) OriginalName() string {
	camel := c.CamelName()
	camel = strings.Split(camel, "By")[0]
	camel = strings.TrimSuffix(camel, "AndWait")
	camel = strings.TrimSuffix(camel, "All")
	return camel
}

func (c *call) HasVariable(camelName string) bool {
	// this assumes that v is already camelCased
	tmp := []int{0}
	c.Traverse(func(e expression) {
		v, ok := e.(*variable)
		if !ok {
			return
		}
		if v.CamelName() == camelName {
			tmp[0]++
		}
	})
	return tmp[0] > 0
}

func (c *call) Traverse(cb func(expression)) {
	for _, v := range c.Args {
		cb(v)
		if t, ok := v.(traversable); ok {
			t.Traverse(cb)
		}
	}
}

func (c *call) Type() string {
	return "call"
}

type initVar struct {
	code.Named
	Value expression
}

type example struct {
	code.Named
	// TODO: add Method and Service
	IsAccount bool
	Calls     []*call
	Cleanup   []*call
	Asserts   []expression
	Init      []*initVar
	scope     map[string]expression
}

func (ex *example) FullName() string {
	return ex.Name
}

func (ex *example) findCall(svcCamelName, methodCamelName string) *call {
	for _, v := range ex.Calls {
		if v.Service == nil {
			continue
		}
		if v.Service.CamelName() != svcCamelName {
			continue
		}
		if v.OriginalName() != methodCamelName {
			continue
		}
		return v
	}
	for _, v := range ex.Cleanup {
		if v.Service == nil {
			continue
		}
		if v.Service.CamelName() != svcCamelName {
			continue
		}
		if v.OriginalName() != methodCamelName {
			continue
		}
		return v
	}
	return nil
}
