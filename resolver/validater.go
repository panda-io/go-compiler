package resolver

//TO-DO validate
//invalid recursive type // invalid inheritance
//validate if access able
//match type
//template function, not virtual
//return abstract class is not allowed

/********************************
- constant expressions
  - constant expressions used in declarations of:
    - array bounds
    - selectors in case statements
    - bit-field length specification
    - enumeration initializers
  - only operands that are legal constant expressions:
    - literals
    - enumeration constants
	- values declared as const
*********************************/

/*
import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/token"
)


//TO-DO match type (int, float, string ... with unary, binary)
// match number, string...

type ValueType int

const (
	InvalidValue ValueType = iota
	IntegerValue
	FloatValue
	CharValue
	StringValue
	BoolValue
	NullValue
	VoidValue
	ObjectValue
)

type ExpressionValue struct {
	Type          ValueType
	QualifiedName string
}

func (e *ExpressionValue) IsNumber() bool {
	return e.Type == IntegerValue || e.Type == FloatValue
}

func (r *Resolver) expressionValue(e expression.Expression) *ExpressionValue {
	switch t := e.(type) {
	case *expression.Literal:
		switch t.Type {
		case token.BOOL:
			return &ExpressionValue{
				Type: BoolValue,
			}
		case token.INT:
			return &ExpressionValue{
				Type: IntegerValue,
			}
		case token.FLOAT:
			return &ExpressionValue{
				Type: FloatValue,
			}
		case token.CHAR:
			return &ExpressionValue{
				Type: CharValue,
			}
		case token.STRING:
			return &ExpressionValue{
				Type: StringValue,
			}
		case token.NULL:
			return &ExpressionValue{
				Type: NullValue,
			}
		default:
			return &ExpressionValue{
				Type: InvalidValue,
			}
		}

	case *expression.Identifier:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.This:
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Super:
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Parentheses:
		return r.expressionValue(t.Expression)

	case *expression.MemberAccess:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Subscripting:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Invocation:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.New:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Increment:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Decrement:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Unary:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Binary:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	case *expression.Conditional:
		//TO-DO
		return &ExpressionValue{
			Type: InvalidValue,
		}

	default:
		return &ExpressionValue{
			Type: InvalidValue,
		}
	}
}

func (r *Resolver) constExpression(e expression.Expression) bool {
	switch t := e.(type) {

	case *expression.Literal:
		return true

	case *expression.Identifier:
		//TO-DO
		return false

	case *expression.This:
		return false

	case *expression.Super:
		return false

	case *expression.Parentheses:
		return r.constExpression(t.Expression)

	case *expression.MemberAccess:
		//TO-DO
		return false

	case *expression.Subscripting:
		return false

	case *expression.Invocation:
		return false

	case *expression.New:
		return false

	case *expression.Increment:
		return false

	case *expression.Decrement:
		return false

	case *expression.Unary:
		return r.constExpression(t.Expression)

	case *expression.Binary:
		//TO-DO check token // not assign then check children
		return false

	case *expression.Conditional:
		return r.constExpression(t.Condition) &&
			r.constExpression(t.First) &&
			r.constExpression(t.Second)

	default:
		return false
	}
}
*/
