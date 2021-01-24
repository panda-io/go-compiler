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

func (i *Interface) ResolveParents(p *Program) {
	for _, parent := range i.Parents {
		_, d := p.FindDeclaration(parent)
		if d == nil {
			p.Error(parent.Position, fmt.Sprintf("%s undefined", parent.Name))
		} else {
			switch t := d.(type) {
			case *Interface:
				i.Interfaces = append(i.Interfaces, t)

			default:
				p.Error(parent.Position, fmt.Sprintf("invalid parent type: %s", t.Identifier()))
			}
		}
	}
}
