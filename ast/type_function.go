package ast

import "github.com/panda-foundation/go-compiler/ir"

type TypeFunction struct {
	TypeBase
	ReturnType Type
	Parameters []Type
}

func (f *TypeFunction) Type(p *Program) ir.Type {
	var types []ir.Type
	for _, param := range f.Parameters {
		types = append(types, param.Type(p))
	}
	var t ir.Type = ir.Void
	if f.ReturnType != nil {
		t = f.ReturnType.Type(p)
	}
	return ir.NewPointerType(ir.NewFuncType(t, types...))
}
