package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type This struct {
	ExpressionBase
}

func (t *This) Type(c *Context, expected ir.Type) ir.Type {
	if c.Function.Class != nil {
		return c.Function.Class.IRStruct
	}
	return nil
}

func (t *This) GenerateIR(c *Context, expected ir.Type) ir.Value {
	if c.Function.Class != nil {
		return c.FindObject(ClassThis)
	}

	c.Program.Error(t.Position, "'this' undefined")
	return nil
}

func (*This) IsConstant(p *Program) bool {
	return false
}

func (*This) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	return nil
}
