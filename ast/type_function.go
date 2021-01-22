package ast

import "github.com/panda-foundation/go-compiler/ir"

type TypeFunction struct {
	TypeBase
	ReturnType Type
	Parameters []Type
}

func (f *TypeFunction) Type(c *Context) ir.Type {
	var types []ir.Type
	for _, p := range typ.Parameters {
		types = append(types, TypeOf(c, declarations, p))
	}
	return ir.NewPointerType(ir.NewFuncType(TypeOf(c, declarations, typ.ReturnType), types...))
}
