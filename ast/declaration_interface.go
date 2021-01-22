package ast

import "fmt"

type Interface struct {
	DeclarationBase
	TypeParameters *TypeParameters
	Parents        []*TypeName
	Functions      []*Function

	Interfaces []*Interface
}

func (i *Interface) AddFunction(f *Function) error {
	for _, v := range i.Functions {
		if v.Name.Name == f.Name.Name {
			return fmt.Errorf("%s redeclared", f.Name.Name)
		}
	}
	i.Functions = append(i.Functions, f)
	return nil
}

func (i *Interface) ResolveParents(ctx *Context) {
	for _, p := range i.Parents {
		_, d := ctx.FindDeclaration(p)
		if d == nil {
			ctx.Error(p.Position, fmt.Sprintf("%s undefined", p.Name))
		} else {
			switch t := d.(type) {
			case *Interface:
				i.Interfaces = append(i.Interfaces, t)

			default:
				ctx.Error(p.Position, fmt.Sprintf("invalid parent type: %s", t.Identifier()))
			}
		}
	}
}
