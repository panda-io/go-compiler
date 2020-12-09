package resolver

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
)

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
