package resolver

/*
import (
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
)

func (r *Resolver) resolveStatement(stmt statement.Statement, typeParams *types.TypeParameters) {
	if stmt == nil {
		return
	}

	// open scope when start compound, if, for, switch, case, try, function
	switch current := stmt.(type) {
	case *statement.Empty:

	case *statement.Raw:

	case *statement.Expression:
		r.resolveExpression(current.Expression, typeParams)

	case *statement.Declaration:
		r.resolveType(current.Type, typeParams)
		r.resolveExpression(current.Value, typeParams)
		r.currentScope.Insert(&Object{
			Kind: VariableObject,
			Name: current.Name.Name,
			Type: current.Type,
		})

	case *statement.Return:
		r.resolveExpression(current.Expression, typeParams)

	case *statement.Continue:

	case *statement.Break:

	case *statement.Try:
		r.currentScope = r.currentScope.OpenScope()
		r.resolveStatement(current.Try, typeParams)
		r.currentScope = r.currentScope.CloseScope()

		r.currentScope = r.currentScope.OpenScope()
		r.resolveParameters(current.Operand, typeParams)
		r.resolveStatement(current.Catch, typeParams)
		r.currentScope = r.currentScope.CloseScope()

		r.currentScope = r.currentScope.OpenScope()
		r.resolveStatement(current.Finally, typeParams)
		r.currentScope = r.currentScope.CloseScope()

	case *statement.Throw:
		r.resolveExpression(current.Expression, typeParams)

	case *statement.If:
		r.currentScope = r.currentScope.OpenScope()
		r.resolveStatement(current.Initialization, typeParams)
		r.resolveStatement(current.Condition, typeParams)
		r.resolveStatement(current.Body, typeParams)
		r.resolveStatement(current.Else, typeParams)
		r.currentScope = r.currentScope.CloseScope()

	case *statement.Switch:
		r.currentScope = r.currentScope.OpenScope()
		r.resolveStatement(current.Initialization, typeParams)
		r.resolveStatement(current.Operand, typeParams)
		for _, c := range current.Body {
			r.currentScope = r.currentScope.OpenScope()
			r.resolveExpression(c.Case, typeParams)
			r.resolveStatement(c.Body, typeParams)
			r.currentScope = r.currentScope.CloseScope()
		}
		r.currentScope = r.currentScope.CloseScope()

	case *statement.For:
		r.currentScope = r.currentScope.OpenScope()
		r.resolveStatement(current.Initialization, typeParams)
		r.resolveStatement(current.Condition, typeParams)
		r.resolveStatement(current.Post, typeParams)
		r.resolveStatement(current.Body, typeParams)
		r.currentScope = r.currentScope.CloseScope()

	case *statement.Foreach:
		r.currentScope = r.currentScope.OpenScope()
		r.resolveStatement(current.Key, typeParams)
		r.resolveStatement(current.Item, typeParams)
		r.resolveExpression(current.Iterator, typeParams)
		r.resolveStatement(current.Body, typeParams)
		r.currentScope = r.currentScope.CloseScope()

	case *statement.Block:
		r.currentScope = r.currentScope.OpenScope()
		for _, s := range current.Statements {
			r.resolveStatement(s, typeParams)
		}
		r.currentScope = r.currentScope.CloseScope()
	}
}
*/
