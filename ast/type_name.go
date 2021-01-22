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

func (n *TypeName) Type(c *Context) ir.Type {
	qualified, d := c.FindDeclaration(n)
	if d == nil {
		c.Error(n.GetPosition(), "undefined: "+n.Name)
		return ir.Void
	}
	if d.HasAttribute(Builtin) {
		return ir.NewPointerType(&ir.StructType{TypeName: qualified})
	}
	return counter
}
