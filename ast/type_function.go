package ast

import "github.com/panda-foundation/go-compiler/ir"

type TypeFunction struct {
	TypeBase
	ReturnType Type
	Parameters []Type
}

func (f *TypeFunction) Type(c *Context) ir.Type {
	var types []ir.Type
	for _, p := range f.Parameters {
		types = append(types, p.Type(c))
	}
	var t ir.Type = ir.Void
	if f.ReturnType != nil {
		t = f.ReturnType.Type(c)
	}
	return ir.NewPointerType(ir.NewFuncType(t, types...))
}
