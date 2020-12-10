package resolver

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/types"
)

func (r *Resolver) resolveExpression(e expression.Expression, typeParams *types.TypeParameters) {
	switch expr := e.(type) {
	case *expression.Literal:

	case *expression.Identifier:
		//TO-DO ???

	case *expression.This:
		//TO-DO ???

	case *expression.Super:
		//TO-DO ???

	case *expression.Parentheses:
		r.resolveExpression(expr.Expression, typeParams)

	case *expression.MemberAccess:
		//TO-DO ???

	case *expression.Subscripting:
		r.resolveExpression(expr.Parent, typeParams)
		r.resolveExpression(expr.Element, typeParams)

	case *expression.Invocation:
		r.resolveExpression(expr.Function, typeParams)
		r.resolveArguments(expr.Arguments, typeParams)

	case *expression.New:
		r.resolveType(expr.Type, typeParams)
		r.resolveArguments(expr.Arguments, typeParams)

	case *expression.Increment:
		r.resolveExpression(expr.Expression, typeParams)

	case *expression.Decrement:
		r.resolveExpression(expr.Expression, typeParams)

	case *expression.Unary:
		r.resolveExpression(expr.Expression, typeParams)

	case *expression.Binary:
		r.resolveExpression(expr.Left, typeParams)
		r.resolveExpression(expr.Right, typeParams)

	case *expression.Conditional:
		r.resolveExpression(expr.Condition, typeParams)
		r.resolveExpression(expr.First, typeParams)
		r.resolveExpression(expr.Second, typeParams)
	}
}
