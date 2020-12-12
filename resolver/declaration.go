package resolver

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/types"
)

func (r *Resolver) resolveDeclaration(d declaration.Declaration, typeParams *types.TypeParameters) {
	switch m := d.(type) {
	case *declaration.Variable:
		r.resolveVariable(m, typeParams)

	case *declaration.Function:
		r.resolveFunction(m, typeParams)

	case *declaration.Enum:
		// TO-DO validate const expr

	case *declaration.Interface:
		if m.TypeParameters != nil {
			r.resolveTypeParameters(m.TypeParameters, typeParams)
		}
		r.resolveParents(m.Parents, typeParams)
		for _, f := range m.Members {
			r.resolveFunction(f.(*declaration.Function), m.TypeParameters)
		}

	case *declaration.Class:
		if m.TypeParameters != nil {
			r.resolveTypeParameters(m.TypeParameters, typeParams)
		}
		r.resolveParents(m.Parents, typeParams)
		r.currentScope = r.classScopes[m.QualifinedName]
		for _, member := range m.Members {
			switch memberType := member.(type) {
			case *declaration.Variable:
				r.resolveVariable(memberType, typeParams)

			case *declaration.Function:
				r.resolveFunction(memberType, typeParams)
			}
		}
		r.currentScope = nil
	}
}

func (r *Resolver) resolveVariable(variable *declaration.Variable, typeParams *types.TypeParameters) {
	r.resolveType(variable.Type, typeParams)
	if variable.Value != nil {
		r.resolveExpression(variable.Value, typeParams)
	}
	// TO-DO validate const expr
}

func (r *Resolver) resolveFunction(function *declaration.Function, typeParams *types.TypeParameters) {
	if function.TypeParameters != nil {
		r.resolveTypeParameters(function.TypeParameters, typeParams)
	}
	if function.ReturnType != nil {
		r.resolveType(function.ReturnType, typeParams)
	}

	r.currentScope = r.currentScope.OpenScope()
	if function.Parameters != nil {
		tp := typeParams
		if tp == nil {
			tp = function.TypeParameters
		} else if function.TypeParameters != nil {
			tp = &types.TypeParameters{}
			tp.Parameters = append(typeParams.Parameters, function.TypeParameters.Parameters...)
		}
		r.resolveParameters(function.Parameters, tp)
	}
	if function.Body != nil {
		r.resolveStatement(function.Body, typeParams)
	}
	r.currentScope = r.currentScope.CloseScope()
}

func (r *Resolver) resolveParents(parents []*types.TypeName, typeParams *types.TypeParameters) {
	total := 0
	for _, p := range parents {
		r.resolveTypeName(p, typeParams)
		if r.declarations[p.QualifiedName] != nil && r.declarations[p.QualifiedName].Kind == ClassObject {
			total++
		}
	}
	if total > 1 {
		r.error(parents[0].GetPosition(), "class can not inherit more than 1 class")
	}
}
