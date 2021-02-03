package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type TypeName struct {
	TypeBase
	Name          string
	Selector      string
	TypeArguments *TypeArguments
}

func (n *TypeName) Type(p *Program) ir.Type {
	qualified, d := p.FindDeclaration(n)
	if d == nil {
		p.Error(n.GetPosition(), "undefined: "+n.Name)
		return ir.Void
	}
	if IsBuiltinType(qualified) {
		return ir.NewPointerType(&ir.StructType{TypeName: qualified})
	}
	return counter
}
