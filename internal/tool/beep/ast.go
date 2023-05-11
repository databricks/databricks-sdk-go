package beep

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
	name string
}

func (v *variable) Type() string {
	return "variable"
}

type entity struct {
	pkg         string
	name        string
	fieldValues []*fieldValue
}

func (e *entity) Type() string {
	return "entity"
}

type array struct {
	pkg    string
	name   string
	Values []expression
}

func (a *array) Type() string {
	return "array"
}

type fieldValue struct {
	name  string
	value expression
}

type example struct {
	name      string
	isAccount bool
	calls     []*call
	cleanup   []*call
	asserts   []*binaryExpr
	scope     map[string]expression
}

type call struct {
	svc    string
	method string
	assign string
	args   []expression
}

func (c *call) Type() string {
	return "call"
}
