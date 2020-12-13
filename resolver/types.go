package resolver

import (
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/types"
)

func (r *Resolver) resolveType(typ types.Type, typeParams *types.TypeParameters) {
	if typeName, ok := typ.(*types.TypeName); ok {
		r.resolveTypeName(typeName, typeParams)
		if typeName.TypeArguments != nil {
			r.resolveTypeArguments(typeName.TypeArguments, typeParams)
		}
	}
}

func (r *Resolver) resolveParameters(params *types.Parameters, typeParams *types.TypeParameters) {
	for _, arg := range params.Parameters {
		if typeParameter, ok := arg.Type.(*types.TypeName); ok {
			r.resolveTypeName(typeParameter, typeParams)
		}
		r.currentScope.Insert(&Object{
			Kind: VariableObject,
			Name: arg.Name,
			Type: arg.Type,
		})
	}
}

func (r *Resolver) resolveTypeParameters(typeParams *types.TypeParameters, parent *types.TypeParameters) {
	for _, arg := range typeParams.Parameters {
		if parent != nil {
			for _, p := range parent.Parameters {
				if p.Name == arg.Name {
					r.error(arg.GetPosition(), fmt.Sprintf("generic type %s conflict with class template", p.Name))
				}
			}
		}
		if typeParameter, ok := arg.Type.(*types.TypeName); ok {
			r.resolveTypeName(typeParameter, nil)
		}
	}
}

func (r *Resolver) resolveTypeArguments(args *types.TypeArguments, typParams *types.TypeParameters) {
	for _, arg := range args.Arguments {
		if typeArgument, ok := arg.(*types.TypeName); ok {
			r.resolveTypeName(typeArgument, typParams)
		}
	}
}

func (r *Resolver) resolveTypeName(typeName *types.TypeName, typeParams *types.TypeParameters) {
	names := r.findQualifiedName(typeName, typeParams)
	if len(names) == 0 {
		r.error(typeName.GetPosition(), fmt.Sprintf("%s undefined", typeName.Name))
	} else if len(names) > 1 {
		r.error(typeName.GetPosition(), fmt.Sprintf("ambiguous type %s", typeName.Name))
	} else {
		typeName.QualifiedName = names[0]
	}
}

func (r *Resolver) resolveArguments(args *types.Arguments, typeParams *types.TypeParameters) {
	for _, arg := range args.Arguments {
		if expr, ok := arg.(expression.Expression); ok {
			r.resolveExpression(expr, typeParams)
		} else {
			r.error(arg.GetPosition(), "invalid argument")
		}
	}
}

func (r *Resolver) findQualifiedName(typeName *types.TypeName, typeParams *types.TypeParameters) []string {
	names := []string{}
	// search in type parameters
	if typeParams != nil {
		for _, p := range typeParams.Parameters {
			if typeName.Name == p.Name {
				names = append(names, typeName.Name)
			}
		}
	}
	// search global
	names = r.checkQualifiedName(typeName, names, typeName.Name)
	// search by imports
	for _, u := range r.source.Imports {
		if strings.HasPrefix(typeName.Name, u.Alias+".") {
			n := strings.Replace(typeName.Name, u.Alias, u.Namespace, 1)
			names = r.checkQualifiedName(typeName, names, n)
		}
	}
	// search same package
	if !strings.Contains(typeName.Name, ".") {
		for _, m := range r.source.Members {
			if m.Identifier() == typeName.Name {
				n := typeName.Name
				if r.source.Namespace != "" {
					n = r.source.Namespace + "." + n
				}
				names = r.checkQualifiedName(typeName, names, n)
			}
		}
	}
	return names
}

func (r *Resolver) checkQualifiedName(typeName *types.TypeName, names []string, qualifiedName string) []string {
	if _, ok := r.declarations[qualifiedName]; ok {
		found := false
		for _, n := range names {
			if n == qualifiedName {
				found = true
				break
			}
		}
		if !found {
			names = append(names, qualifiedName)
		}
	}
	return names
}
