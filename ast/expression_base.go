package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Base struct {
	ExpressionBase
}

func (*Base) Type(c *Context) ir.Type {
	return nil
}

func (*Base) GenerateIR(c *Context) ir.Value {
	return nil
}

func (*Base) IsConstant(p *Program) bool {
	return false
}

func (*Base) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	return nil
}
