package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type This struct {
	ExpressionBase
}

func (t *This) Type(c *Context) ir.Type {
	if c.Function.Class != nil {
		return c.Function.Class.IRStruct
	}
	c.Program.Error(t.Position, "'this' is not defined")
	return nil
}

func (t *This) GenerateIR(c *Context) ir.Value {
	if c.Function.Class != nil {
		return c.FindObject(ClassThis)
	}
	c.Program.Error(t.Position, "'this' is not defined")
	return nil
}

func (*This) IsConstant() bool {
	return false
}

func (*This) GenerateConstIR(p *Program, exprect ir.Type) ir.Value {
	return nil
}
