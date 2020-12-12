package resolver

import (
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/types"
)

func (r *Resolver) resolveExpression(e expression.Expression, typeParams *types.TypeParameters) {
	switch expr := e.(type) {
	case *expression.Literal:

	case *expression.Identifier:
		if r.currentScope.Find(expr.Name) == nil {
			objects := r.findObjectsFromImportScope(expr.Name)
			if len(objects) == 0 {
				r.error(expr.Position, fmt.Sprintf("%s undefined", expr.Name))
			} else if len(objects) > 1 {
				r.error(expr.Position, fmt.Sprintf("ambiguous variable %s", expr.Name))
			}
		}

	case *expression.This:
		r.error(expr.Position, "error use of keyword this")

	case *expression.Super:
		r.error(expr.Position, "error use of keyword base")

	case *expression.Parentheses:
		r.resolveExpression(expr.Expression, typeParams)

	case *expression.MemberAccess:
		r.resolveMemberAccess(expr, typeParams)

	case *expression.Subscripting:
		r.resolveExpression(expr.Parent, typeParams)
		r.resolveExpression(expr.Element, typeParams)
		//TO-DO check operator overload

	case *expression.Invocation:
		r.resolveExpression(expr.Function, typeParams)
		r.resolveArguments(expr.Arguments, typeParams)
		//TO-DO check type match, check if function exist

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

func (r *Resolver) findObjectsFromImportScope(name string) []*Object {
	objects := []*Object{}
	object := r.packageScopes[ast.Global].Find(name)
	if object != nil {
		objects = append(objects, object)
	}
	for _, importScope := range r.importsScopes {
		if object = importScope.Find(name); object != nil {
			objects = append(objects, object)
		}
	}
	return objects
}

// parent can be interface, enum, class, namespace
func (r *Resolver) resolveMemberAccess(e *expression.MemberAccess, typeParams *types.TypeParameters) {
	//TO-DO ??? // should resolve before here
	// can be parent: []  ()
	accessChain := []*expression.MemberAccess{e}
	if p, ok := e.Parent.(*expression.MemberAccess); ok {
		accessChain = append(accessChain, p)
		e = p
	}

	var object *Object
	first := accessChain[len(accessChain)-1]
	switch p := first.Parent.(type) {
	case *expression.This:
	case *expression.Super:
		//TO-DO check parent scope
	case *expression.Identifier:
		object = r.currentScope.Find(p.Name)
		if object == nil {
			object = r.packageScopes[r.source.Namespace].Find(p.Name)
		}
		if object == nil {
			for _, i := range r.source.Imports {
				if i.Alias == p.Name {
					first.FullNamespace = strings.Replace(i.Namespace, ".", "::", -1)
				}
			}
		}
		// TO-DO continue validate type
	}
}
