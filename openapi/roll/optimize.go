package roll

import (
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi/code"
)

type enum struct {
	code.Named
	Package string
	Entity  *code.Named
	Content string
}

func (e *enum) Type() string {
	return "enum"
}

func (s *Suite) OptimizeWithApiSpec(b *code.Batch) error {
	o := &suiteOptimizer{b}
	for _, e := range s.examples {
		err := o.optimizeExample(e)
		if err != nil {
			return err
		}
	}
	return nil
}

type suiteOptimizer struct {
	*code.Batch
}

func (s *suiteOptimizer) optimizeExample(e *example) error {
	for _, c := range e.Calls {
		err := s.optimizeCall(c)
		if err != nil {
			return err
		}
	}
	for _, c := range e.Cleanup {
		err := s.optimizeCall(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *suiteOptimizer) optimizeCall(c *call) error {
	for i := range c.Args {
		err := s.optimizeExpression(&c.Args[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *suiteOptimizer) enumConstant(l *lookup) expression {
	potentialPackage := l.Variable()
	potentialEnumValue := l.Field.PascalName()
	for _, pkg := range s.Packages() {
		if pkg.Name != potentialPackage {
			continue
		}
		for _, v := range pkg.Types() {
			if v.Schema != nil && !v.Schema.IsEnum() {
				continue
			}
			prefix := v.PascalName()
			if !strings.HasPrefix(potentialEnumValue, prefix) {
				continue
			}
			value := strings.TrimPrefix(potentialEnumValue, prefix)
			for _, e := range v.Enum() {
				if e.PascalName() != value {
					continue
				}
				return &enum{
					Package: potentialPackage,
					Entity: &code.Named{
						Name: prefix,
					},
					Named: code.Named{
						Name: value,
					},
					Content: e.Content,
				}
			}
		}
	}
	return nil
}

func (s *suiteOptimizer) optimizeExpression(e *expression) (err error) {
	switch x := (*e).(type) {
	case *call:
		return s.optimizeCall(x)
	case *lookup:
		enumConstant := s.enumConstant(x)
		if enumConstant != nil {
			*e = enumConstant
			return
		}
		return s.optimizeExpression(&x.X)
	case *binaryExpr:
		err = s.optimizeExpression(&x.Left)
		if err != nil {
			return err
		}
		err = s.optimizeExpression(&x.Right)
		if err != nil {
			return err
		}
	case *indexExpr:
		err = s.optimizeExpression(&x.Left)
		if err != nil {
			return err
		}
		err = s.optimizeExpression(&x.Right)
		if err != nil {
			return err
		}
	case *entity:
		for i := range x.FieldValues {
			err = s.optimizeExpression(&x.FieldValues[i].Value)
			if err != nil {
				return err
			}
		}
	case *mapLiteral:
		for i := range x.Pairs {
			err = s.optimizeExpression(&x.Pairs[i].Key)
			if err != nil {
				return err
			}
			err = s.optimizeExpression(&x.Pairs[i].Value)
			if err != nil {
				return err
			}
		}
	case *array:
		for i := range x.Values {
			err = s.optimizeExpression(&x.Values[i])
			if err != nil {
				return err
			}
		}
	default:
		return nil
	}
	return nil
}
